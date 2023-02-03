package errors_test

import (
	"testing"

	"github.com/TheDSCPL/minimal-errors"

	"github.com/stretchr/testify/assert"
)

func TestCustomFunctionsWithStdlib(t *testing.T) {
	// This test is to ensure that the custom error factories
	// are compatible the stdlib errors API (errors.Is,
	// errors.As, errors.Unwrap)
	baseErr1 := &customError1{"base"}
	baseErr2 := errors.Wrap(baseErr1, "foo")
	baseErrIndependent := errors.New("independent")

	// Wrap
	{
		err := errors.Wrap(baseErr2, "bar")

		assert.Equal(t, baseErr2, errors.Unwrap(err))
		assert.True(t, errors.Is(err, baseErr2))
		assert.True(t, errors.Is(err, baseErr1))
		var asPtrResult *customError1
		var asReturn bool
		assert.NotPanics(t, func() {
			asReturn = errors.As(err, &asPtrResult)
		})
		assert.True(t, asReturn)
		assert.Panics(t, func() {
			asReturn = errors.As(err, &customNonError{})
		})
	}

	// Redefine
	{
		err := errors.Redefine(baseErrIndependent, baseErr1)

		assert.Equal(t, baseErrIndependent, errors.Unwrap(err))
		assert.True(t, errors.Is(err, baseErrIndependent))
		assert.False(t, errors.Is(err, baseErr1))
		assert.False(t, errors.Is(err, baseErr2))
		var asPtrResult *customError1
		var asReturn bool
		assert.NotPanics(t, func() {
			asReturn = errors.As(err, &asPtrResult)
		})
		assert.False(t, asReturn)
		assert.NotPanics(t, func() {
			// will have the same underlying type as baseErrIndependent
			tmp := errors.New("tmp")
			asReturn = errors.As(err, &tmp)
		})
		assert.True(t, asReturn)
		assert.Panics(t, func() {
			asReturn = errors.As(err, &customNonError{})
		})
	}

	// WithMetadata
	{
		err := errors.WithMetadata(baseErr2, iMeta2{})

		assert.Equal(t, baseErr2, errors.Unwrap(err))
		assert.True(t, errors.Is(err, baseErr1))
		assert.True(t, errors.Is(err, baseErr2))
		var asPtrResult *customError1
		var asReturn bool
		assert.NotPanics(t, func() {
			asReturn = errors.As(err, &asPtrResult)
		})
		assert.True(t, asReturn)
		assert.Panics(t, func() {
			asReturn = errors.As(err, &customNonError{})
		})
	}

	// WithCause
	{
		err := errors.WithCause(baseErr2, baseErrIndependent)

		assert.Equal(t, baseErr2, errors.Unwrap(err))
		assert.True(t, errors.Is(err, baseErr1))
		assert.True(t, errors.Is(err, baseErr2))
		assert.False(t, errors.Is(err, baseErrIndependent))
		var asPtrResult *customError1
		var asReturn bool
		assert.NotPanics(t, func() {
			asReturn = errors.As(err, &asPtrResult)
		})
		assert.True(t, asReturn)
		assert.Panics(t, func() {
			errors.As(err, &customNonError{})
		})
	}
}

type customError1 struct {
	msg string
}

func (e *customError1) Error() string {
	return e.msg
}

type customNonError struct {
	_ string
}
