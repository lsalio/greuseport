// +build linux darwin

package greuseport

/*
#cgo CFLAGS: -Wall -std=c99
#include <stdbool.h>
#include <sys/types.h>
#include <sys/socket.h>

bool reuse(int fd) {
	char value = 1;

	if (setsockopt(fd, SOL_SOCKET, SO_REUSEADDR, &value, sizeof(value)) < 0) {
		return false;
	}
	if (setsockopt(fd, SOL_SOCKET, SO_REUSEPORT, &value, sizeof(value)) < 0) {
		return false;
	}
	return true;
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
