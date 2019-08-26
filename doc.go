// Package openshim2 provides shims for the OpenBSD pledge and unveil syscalls.
//
// Unlike plain golang.org/x/sys/unix, openshim2 provides no-op function
// equivalents to allow compiling programs on OpenBSD and non-OpenBSD systems.
//
// See openshim2_openbsd.go for function documentation.
package openshim2
