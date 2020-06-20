// +build windows

package greuseport

import (
	"sync"
	"syscall"
	"unsafe"
)

type setsockopt struct {
	dll  *syscall.DLL
	proc *syscall.Proc
	mu   sync.Mutex
}

func (s *setsockopt) Call(fd uintptr, level int, opt int, val int) error {
	if err := s.FindProc(); err != nil {
		return err
	}

	v := int32(val)
	args := []uintptr{
		fd,
		uintptr(int32(level)),
		uintptr(int32(opt)),
		uintptr(unsafe.Pointer((*byte)(unsafe.Pointer(&v)))),
		uintptr(int32(unsafe.Sizeof(v))),
	}

	r, _, err := s.proc.Call(args...)
	if err != nil {
		if errno, ok := err.(syscall.Errno); ok {
			if errno != 0 {
				return errno
			}
		} else {
			return err
		}
	}

	if r == uintptr(^uint32(0)) {
		return syscall.Errno(r)
	}
	return nil
}

func (s *setsockopt) FindProc() (err error) {
	if s.proc == nil {
		s.mu.Lock()
		defer s.mu.Unlock()

		if s.dll, err = syscall.LoadDLL("ws2_32.dll"); err == nil {
			s.proc, err = s.dll.FindProc("setsockopt")
		}
	}
	return
}

var proc setsockopt

func control(_ string, _ string, c syscall.RawConn) (err error) {
	return c.Control(func(fd uintptr) {
		err = proc.Call(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1)
	})
}
