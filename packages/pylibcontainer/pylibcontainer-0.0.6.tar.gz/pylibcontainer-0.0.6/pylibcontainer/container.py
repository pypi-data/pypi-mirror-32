from __future__ import print_function
import os
import pwd
import grp
from uuid import uuid4
from os.path import exists, join
from cgroupspy import trees
from tmsyscall.unshare import unshare, CLONE_NEWNS, CLONE_NEWUTS, CLONE_NEWIPC, CLONE_NEWPID, CLONE_NEWNET
from tmsyscall.mount import mount, unmount, mount_procfs
from tmsyscall.mount import MS_BIND, MS_PRIVATE, MS_REC, MNT_DETACH, MS_REMOUNT, MS_RDONLY
from tmsyscall.pivot_root import pivot_root
from pylibcontainer.utils import HumanSize
from pylibcontainer.colorhelper import print_info, print_list

DEFAULT_limit_in_bytes = 1024*1024

def drop_privileges(uid_name='nobody', gid_name='nogroup'):

    # Get the uid/gid from the name
    running_uid = pwd.getpwnam(uid_name).pw_uid
    running_gid = grp.getgrnam(gid_name).gr_gid

    # Remove group privileges
    os.setgroups([])

    # Try setting the new uid/gid
    os.setgid(running_gid)
    os.setuid(running_uid)

    # Ensure a very conservative umask
    _ = os.umask(0o77)

def setup_memory_cgroup(container_id, pid):
    """ Create cgroup for the container id """
    cgroup_tree = trees.Tree()
    cgroup_set = cgroup_tree.get_node_by_path('/memory/pylibcontainer')
    if not cgroup_set:
        cgroup_set = cgroup_tree.get_node_by_path('/memory/')
        pylibcontainer_cgroup = cgroup_set.create_cgroup('pylibcontainer')
        cgroup_set = pylibcontainer_cgroup

    # container_mem_cgroup
    new_group = cgroup_set.create_cgroup(container_id)
    new_group.controller.tasks = pid
    new_group.controller.limit_in_bytes = DEFAULT_limit_in_bytes

def delete_memory_cgroup(container_id):
    cgroup_tree = trees.Tree()
    cgroup_set = cgroup_tree.get_node_by_path('/memory/pylibcontainer')
    cgroup_set.delete_cgroup(container_id)

def print_memory_stats(container_id):
    cgroup_tree = trees.Tree()
    cgroup_set = cgroup_tree.get_node_by_path('/memory/pylibcontainer/' + container_id)
    controler = cgroup_set.controller
    max_usage_in_bytes = controler.max_usage_in_bytes
    memsw_max_usage_in_bytes = controler.memsw_max_usage_in_bytes
    info_list = {
        "mem used" : "{0:.2S}".format(HumanSize(max_usage_in_bytes)),
        "mem+swap used": "{0:.2S}".format(HumanSize(memsw_max_usage_in_bytes))
    }
    print_list("Max Mem Info", info_list)


def setup_process_isolation(rootfs_path):
    # Detach from parent's mount, hostname, ipc and net  namespaces
    unshare(CLONE_NEWNS| CLONE_NEWUTS | CLONE_NEWIPC| CLONE_NEWNET)

    # Set mount propagation to private recursively. Hopefully equivalent to
    #    mount --make-rprivate /
    # This is needed to prevent mounts in this container leaking to the parent.
    mount('none', '/', None, MS_REC|MS_PRIVATE, "")

    root_fs = rootfs_path
    os.chmod(rootfs_path, 0o755)

    # This bind mount call is needed to satisfy a requirement of the `pivotroot` system call
    #   "new_root and put_old must not be on the same file system as the current root"
    # It is achieved by mounting "new_root" as a bind mount to "new root"
    mount(root_fs, root_fs, "", MS_BIND|MS_REC, "")

    old_root = join(root_fs, ".old_root")
    if not exists(old_root):
        os.makedirs(old_root, 0o700)

    # Remount it as readonly
    mount(root_fs, root_fs, "", MS_BIND|MS_REC|MS_REMOUNT|MS_RDONLY, "")
    pivot_root(root_fs, old_root)

    # We don't want the host root to be available to the container
    unmount("/.old_root", MNT_DETACH)

    os.chdir("/")

    # Mount /proc for apps that need it
    if not exists("proc"):
        os.makedirs("proc", 0o700)
    mount_procfs('.')  # py


def child(rootfs_path, cmd):
    setup_process_isolation(rootfs_path)
    drop_privileges()
    if  len(cmd) > 1 or cmd[0] != '-':
        os.execvp(cmd[0], cmd)

def parent(child_pid):
    container_id = str(uuid4())
    setup_memory_cgroup(container_id, child_pid)
    result = os.waitpid(child_pid, 0)
    print_memory_stats(container_id)
    delete_memory_cgroup(container_id)
    return result


def runc(rootfs_path, command):
    # Detach from pid namespace so that our child get's a clean /proc with the new namespace
    print_info("Memory limit:", "{0:.2S}".format(HumanSize(DEFAULT_limit_in_bytes)))
    unshare(CLONE_NEWPID)
    pid = os.fork()
    if pid == 0:
        child(rootfs_path, command)
        exit(0)
    return parent(pid)
