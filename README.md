# id

`id` returns a stable machine identifier.

Supports:

- AIX (via the machine component of `uname` syscall)
- BSD (via `sysctl hw.uuid`, `sysctl kern.hostuuid`, `/etc/hostid`, `hostid`, or `kenv -q smbios.system.uuid`)
- Darwin (via `gethostuuid`)
- Linux (via `/var/lib/dbus/machine-id`, `/etc/machine-id` or `/sys/class/dmi/id/product_uuid`)
- Windows (via registry `SOFTWARE\Microsoft\Cryptography\MachineGuid`)
- z/OS
