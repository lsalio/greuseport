// +build windows

package greuseport

/*
#cgo CFLAGS: -Wall -std=c99
#cgo LDFLAGS: -lWs2_32
#include <stdbool.h>
#include <winsock.h>

bool reuse(int fd) {
	char value = 1;

	return setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &value, sizeof(value)) != SOCKET_ERROR;
}
*/
import "C"
import (
	"syscall"
)

func control(_ string, _ string, c syscall.RawConn) error {
	return c.Control(func(fd uintptr) {
		C.reuse(C.int(fd))
	})
}
