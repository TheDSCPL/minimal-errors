package errors_test

import (
	stdErrors "errors"
	"fmt"
	"testing"

	errors "github.com/TheDSCPL/minimal-errors"
	"github.com/TheDSCPL/minimal-errors/internal"

	"github.com/stretchr/testify/assert"
)

func TestStdlibAliases(t *testing.T) {
	internal.AssertEqualFn(t, stdErrors.New, errors.New)
	internal.AssertEqualFn(t, stdErrors.As, errors.As)
	internal.AssertEqualFn(t, stdErrors.Is, errors.Is)
	internal.AssertEqualFn(t, stdErrors.Unwrap, errors.Unwrap)
	internal.AssertEqualFn(t, fmt.Errorf, errors.Errorf)
}

func TestWrap(t *testing.T) {
	baseErr := stdErrors.New("base")

	err := errors.Wrap(baseErr, "wrap")
	assert.Equal(t, "wrap: base", err.Error())
	assert.Equal(t, baseErr, errors.Unwrap(err))

	err = errors.Wrap(baseErr, "")
	assert.Equal(t, "base", err.Error())
	assert.Equal(t, baseErr, errors.Unwrap(err))

	err = errors.Wrap(nil, "foo")
	assert.Nil(t, err)
}

func TestRedefine(t *testing.T) {
	oldErr := stdErrors.New("old")
	baseErr := stdErrors.New("base")

	err := errors.Redefine(baseErr, oldErr)
	assert.Equal(t, "base: old", err.Error())
	assert.Equal(t, baseErr, errors.Unwrap(err))

	err = errors.Redefine(baseErr, nil)
	assert.Equal(t, "base", err.Error())
	assert.Equal(t, baseErr, err)

	err = errors.Redefine(nil, oldErr)
	assert.Equal(t, "old", err.Error())
	assert.Nil(t, errors.Unwrap(err))

	err = errors.Redefine(nil, nil)
	assert.Nil(t, err)
}
