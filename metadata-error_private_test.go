package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Testing methods of private structs to ensure 100% coverage

func TestPrivateStructMethods(t *testing.T) {
	assert.Nil(t, (*metadataError)(nil).Metadata())
	assert.Nil(t, (*metadataError)(nil).Unwrap())
}
