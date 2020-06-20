// +build linux darwin

package greuseport

import (
	"syscall"
	"unsafe"
)

type setsockopt struct{}

func (setsockopt) Call(fd int, level int, opt int, v int32) error {
	_, _, e := syscall.Syscall6(syscall.SYS_SETSOCKOPT, uintptr(fd), uintptr(level), uintptr(opt), uintptr(unsafe.Pointer(&v)), uintptr(4), 0)
	if e != 0 {
		return e
	}
	return nil
}

var proc setsockopt

func control(_ string, _ string, c syscall.RawConn) (err error) {
	return c.Control(func(fd uintptr) {
		if err = proc.Call(int(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
			return
		}

		if err = proc.Call(int(fd), syscall.SOL_SOCKET, 0xf /* SO_REUSEPORT */, 1); err != nil {
			return
		}
	})
}
