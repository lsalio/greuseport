//+build !windows,!linux,!darwin

package greuseport

import (
	"errors"
	"syscall"
)

func control(_ string, _ string, c syscall.RawConn) error {
	return errors.New("not supports environment")
}
