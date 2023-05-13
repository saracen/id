//go:build aix

package id

import (
	"golang.org/x/sys/unix"
)

func id() (string, error) {
	var name unix.Utsname
	if err := unix.Uname(&name); err != nil {
		return "", err
	}

	return unix.ByteSliceToString(name.Machine[:]), nil
}
