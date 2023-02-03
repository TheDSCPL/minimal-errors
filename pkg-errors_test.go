package errors_test

import (
	"github.com/TheDSCPL/minimal-errors/internal"
	"runtime"
	"testing"

	"github.com/TheDSCPL/minimal-errors"

	pkgErrors "github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestPkgErrorsAliases(t *testing.T) {
	// Test function aliases
	internal.AssertEqualFn(t, pkgErrors.Cause, errors.Cause)
	internal.AssertEqualFn(t, pkgErrors.WithStack, errors.WithStack)

	// Test type aliases (on build time)
	var _ errors.Frame = *new(pkgErrors.Frame)
	var _ errors.StackTraceT = *new(pkgErrors.StackTrace)
}

func TestStackTrace(t *testing.T) {
	err := errors.New("foo")
	stackTrace := errors.StackTrace(err)
	assert.Nil(t, stackTrace)

	err = errors.WithStack(err)
	stackTrace = errors.StackTrace(err)
	assert.NotNil(t, stackTrace)
	latestCallerFunction := runtime.FuncForPC(uintptr(stackTrace[0]) - 1)
	thisFunctionPc, _, _, ok := runtime.Caller(0)
	assert.True(t, ok)
	thisFunction := runtime.FuncForPC(thisFunctionPc)
	assert.Equal(t, thisFunction.Name(), latestCallerFunction.Name())
}

func TestCause(t *testing.T) {
	baseErr := errors.New("base")
	cause := errors.New("cause")

	err := errors.WithCause(baseErr, cause)
	assert.Equal(t, cause, errors.Cause(err))
}
