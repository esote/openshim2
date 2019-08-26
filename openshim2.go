// +build !openbsd

package openshim2

// LazySysctls is a no-op for non-OpenBSD systems.
func LazySysctls() error {
	return nil
}

// Pledge is a no-op for non-OpenBSD systems.
func Pledge(promises, execpromises string) error {
	return nil
}

// PledgePromises is a no-op for non-OpenBSD systems.
func PledgePromises(promises string) error {
	return nil
}

// PledgeExecpromises is a no-op for non-OpenBSD systems.
func PledgeExecpromises(execpromises string) error {
	return nil
}

// Unveil is a no-op for non-OpenBSD systems.
func Unveil(path, flags string) error {
	return nil
}

// UnveilBlock is a no-op for non-OpenBSD systems.
func UnveilBlock() error {
	return nil
}
