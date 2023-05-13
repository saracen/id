//go:build openbsd || freebsd || dragonfly || hurd || illumos || netbsd || solaris

package id

import (
	"syscall"
)

func id() (string, error) {
	return fallback(
		func() (string, error) { return syscall.Sysctl("hw.uuid") },       // openbsd
		func() (string, error) { return syscall.Sysctl("kern.hostuuid") }, // freebsd
		read("/etc/hostid"),
		run("hostid"),
		run("kenv", "-q", "smbios.system.uuid"),
	)
}
