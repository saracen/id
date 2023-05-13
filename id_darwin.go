//go:build darwin || ios

package id

import (
	"fmt"
	"syscall"
	"unsafe"
)

//go:linkname syscall_syscall syscall.syscall
func syscall_syscall(fn, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)

//go:cgo_import_dynamic libc_gethostuuid gethostuuid "/usr/lib/libSystem.B.dylib"
var libc_gethostuuid_trampoline_addr uintptr

func id() (string, error) {
	var uuid [16]byte
	pUUID := unsafe.Pointer(&uuid[0])
	pWait := &syscall.Timespec{Sec: 10}

	_, _, err := syscall_syscall(libc_gethostuuid_trampoline_addr, uintptr(pUUID), uintptr(unsafe.Pointer(pWait)), 0)
	if err != 0 {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", uuid[:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), nil
}
