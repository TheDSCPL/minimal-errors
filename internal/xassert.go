package internal

import (
	"reflect"

	"github.com/stretchr/testify/assert"
)

func AssertEqualFn(t assert.TestingT, expected, actual interface{}, msgAndArgs ...interface{}) bool {
	// assert.Equal doesn't allow to compare function pointers
	expected, actual = reflect.ValueOf(expected).Pointer(), reflect.ValueOf(actual).Pointer()
	return assert.Equal(t, expected, actual, msgAndArgs...)
}
