//go:build zos

package id

import (
        "runtime"
        "unsafe"
)

func id() (string, error) {
    cvt := *(*int32)(unsafe.Pointer(uintptr(16)))
    pccaavt := *(*int32)(unsafe.Pointer(uintptr(cvt) + 764))
    pcca := *(*int32)(unsafe.Pointer(uintptr(pccaavt))) + 4
    
    var res [12]byte
    copy(res[:], (*[12]byte)(unsafe.Pointer(uintptr(pcca)))[:])
    
    runtime.CallLeFuncByPtr(runtime.XplinkLibvec+0x6e3<<4,
        []uintptr{uintptr(unsafe.Pointer(&res[0])), uintptr(12)})
    
    return string(res[:]), nil
}
