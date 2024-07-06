package errors

import "fmt"

// Wrapf is a convenience function for wrapping an err with a message.
func Wrapf(err error, msg string, args ...any) error {
	args = append(args, err)
	return fmt.Errorf(msg+": %w", args...)
}

// Errorf creates a new error with a message.
func Errorf(msg string, args ...any) error {
	return fmt.Errorf(msg, args...)
}
