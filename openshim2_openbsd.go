// +build openbsd

package openshim2

import (
	"net"

	"golang.org/x/sys/unix"
)

// LazySysctls forces the initialization of lazy sysctl caches in the net
// package. This is only needed if you pledge "inet" and use the net package.
// This should be called only once and before any calls to pledge.
//
// See https://github.com/golang/go/issues/31927
func LazySysctls() error {
	l, err := net.Listen("tcp", "localhost:0")

	if err != nil {
		return err
	}

	return l.Close()
}

// Pledge is a wrapper for x/sys/unix Pledge.
func Pledge(promises, execpromises string) error {
	return unix.Pledge(promises, execpromises)
}

// PledgePromises is a wrapper for x/sys/unix PledgePromises.
func PledgePromises(promises string) error {
	return unix.PledgePromises(promises)
}

// PledgeExecpromises is a wrapper for x/sys/unix PledgeExecpromises.
func PledgeExecpromises(execpromises string) error {
	return unix.PledgeExecpromises(execpromises)
}

// Unveil is a wrapper for x/sys/unix Unveil.
func Unveil(path, flags string) error {
	return unix.Unveil(path, flags)
}

// UnveilBlock is a wrapper for x/sys/unix UnveilBlock.
func UnveilBlock() error {
	return unix.UnveilBlock()
}
