// Package errors contains an opinionated error handling style.
//
// It is a drop-in replacement of stdlib and is very lightweight, as
// all functions which are present in stdlib errors are just aliases.
//
// It uses stdlib errors as much as possible and github.com/pkg/errors
// only for stack traces and causes.
//
// Note about github.com/pkg/errors: According to its documentation,
// that library has been archived because its purpose has been made
// irrelevant by Go 1.13's errors, with the exception of stack traces
// and causes, which weren't implemented in stdlib.
package errors

import (
	stdErrors "errors"
	"fmt"
)

var (
	// New returns an error that formats as the given text.
	// Each call to New returns a distinct error value even if the text is identical.
	//
	// ** Docs copied from https://pkg.go.dev/errors#New
	New = stdErrors.New
	// As finds the first error in err's chain that matches target, and if one is found, sets
	// target to that error value and returns true. Otherwise, it returns false.
	//
	// The chain consists of err itself followed by the sequence of errors obtained by
	// repeatedly calling Unwrap.
	//
	// An error matches target if the error's concrete value is assignable to the value
	// pointed to by target, or if the error has a method As(interface{}) bool such that
	// As(target) returns true. In the latter case, the As method is responsible for
	// setting target.
	//
	// An error type might provide an As method so it can be treated as if it were a
	// different error type.
	//
	// As panics if target is not a non-nil pointer to either a type that implements
	// error, or to any interface type.
	//
	// ** Docs copied from https://pkg.go.dev/errors#As
	As = stdErrors.As
	// Is reports whether any error in err's chain matches target.
	//
	// The chain consists of err itself followed by the sequence of errors obtained by
	// repeatedly calling Unwrap.
	//
	// An error is considered to match a target if it is equal to that target or if
	// it implements a method Is(error) bool such that Is(target) returns true.
	//
	// An error type might provide an Is method so it can be treated as equivalent
	// to an existing error. For example, if MyError defines
	//
	//	func (m MyError) Is(target error) bool { return target == fs.ErrExist }
	//
	// then Is(MyError{}, fs.ErrExist) returns true. See syscall.Errno.Is for
	// an example in the standard library. An Is method should only shallowly
	// compare err and the target and not call Unwrap on either.
	//
	// ** Docs copied from https://pkg.go.dev/errors#Is
	Is = stdErrors.Is
	// Unwrap returns the result of calling the Unwrap method on err, if err's
	// type contains an Unwrap method returning error.
	// Otherwise, Unwrap returns nil.
	//
	// ** Docs copied from https://pkg.go.dev/errors#Unwrap
	Unwrap = stdErrors.Unwrap
	// Errorf formats according to a format specifier and returns the string as a
	// value that satisfies error.
	//
	// If the format specifier includes a %w verb with an error operand,
	// the returned error will implement an Unwrap method returning the operand. It is
	// invalid to include more than one %w verb or to supply it with an operand
	// that does not implement the error interface. The %w verb is otherwise
	// a synonym for %v.
	//
	// ** Docs copied from https://pkg.go.dev/fmt#Errorf
	Errorf = fmt.Errorf
)

// Wrap returns a new error that wraps the error passed
// in the err parameter.
//
// Syntax sugar for Errorf("%s: %w", message, err)
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	if message != "" {
		message += ": "
	}

	return fmt.Errorf("%s%w", message, err)
}

// Redefine returns a new error that wraps baseErr
// but keeps the message of oldErr. oldErr will not
// be wrapped by the new error and, thus, will not be
// detectable by errors.Is, errors.As, Cause or any other
// method based on error unwrapping.
//
// This function differs from Wrap because the second
// argument is an error instead of a string.
//
// This function differs from WithCause because this is
// lighter, as it only saves the message of the oldErr
// instead of a reference to it.
//
// If oldErr is nil, returns baseErr.
//
// If baseErr is nil, returns a new error with the
// message of oldErr.
func Redefine(baseErr error, oldErr error) error {
	if oldErr == nil {
		return baseErr
	}

	if baseErr == nil {
		return New(oldErr.Error())
	}

	return fmt.Errorf("%w: %s", baseErr, oldErr.Error())
}
