package errors

import (
	pkgErrors "github.com/pkg/errors"
)

// ---- STACK ---- //

var (
	// WithStack annotates err with a stack trace at the point WithStack was called.
	// If err is nil, WithStack returns nil.
	//
	// ** Docs copied from https://pkg.go.dev/github.com/pkg/errors#WithStack
	WithStack = pkgErrors.WithStack
)

type (
	// Frame represents a program counter inside a stack frame.
	// For historical reasons if Frame is interpreted as a uintptr
	// its value represents the program counter + 1.
	//
	// ** Docs copied from https://pkg.go.dev/github.com/pkg/errors#Frame
	Frame = pkgErrors.Frame
	// StackTraceT is stack of Frames from innermost (newest) to outermost (oldest).
	//
	// ** Docs copied from https://pkg.go.dev/github.com/pkg/errors#StackTrace
	StackTraceT = pkgErrors.StackTrace
)

// StackTrace returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//	type causer interface {
//	       Cause() error
//	}
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func StackTrace(err error) StackTraceT {
	var stack StackTraceT

	for stack == nil && err != nil {
		tracer, ok := err.(stackTracer)
		if !ok {
			break
		}

		stack = tracer.StackTrace()
		err = Unwrap(err)
	}

	return stack
}

type stackTracer interface {
	StackTrace() StackTraceT
}

// ---- CAUSE ---- //

var (
	// Cause returns the underlying cause of the error, if possible.
	// An error value has a cause if it implements the following
	// interface:
	//
	//     type causer interface {
	//            Cause() error
	//     }
	//
	// If the error does not implement Cause, the original error will
	// be returned. If the error is nil, nil will be returned without further
	// investigation.
	//
	// ** Docs copied from https://pkg.go.dev/github.com/pkg/errors#Cause
	Cause = pkgErrors.Cause
)

// WithCause returns an error that annotates err with the provided cause.
func WithCause(err error, cause error) error {
	return &causedError{err, cause}
}

type iCausedError interface {
	Cause() error
	Unwrap() error
}

type causedError struct {
	error
	cause error
}

var _ iCausedError = (*causedError)(nil)

func (c causedError) Cause() error {
	return c.cause
}

func (c causedError) Unwrap() error {
	return c.error
}
