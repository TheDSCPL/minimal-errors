package errors_test

import (
	"testing"

	"github.com/TheDSCPL/minimal-errors"

	"github.com/stretchr/testify/assert"
)

type iMeta1 = map[string]interface{}
type iMeta2 struct {
	Name string
}
type iMeta3 struct {
	Name string
}

func TestMetadata(t *testing.T) {
	assert.Nil(t, errors.Metadata(nil))

	err := errors.New("foo")
	assert.Equal(t, err, errors.WithMetadata(err, nil))
	assert.Nil(t, errors.Metadata(err))

	meta1 := iMeta1{"foo": "bar"}
	meta2 := iMeta2{"BAR_2"}
	meta3 := iMeta3{"BAR_3"}

	var (
		ret1 iMeta1
		ret2 iMeta2
		ret3 iMeta3
		ok1  bool
		ok2  bool
		ok3  bool
	)

	err = errors.WithMetadata(err, meta1)
	assert.Equal(t, meta1, errors.Metadata(err))
	// TypedMetadata
	ret1, ok1 = errors.TypedMetadata[iMeta1](err)
	assert.Equal(t, meta1, ret1)
	assert.True(t, ok1)
	ret2, ok2 = errors.TypedMetadata[iMeta2](err)
	assert.Zero(t, ret2)
	assert.False(t, ok2)
	ret3, ok3 = errors.TypedMetadata[iMeta3](err)
	assert.Zero(t, ret3)
	assert.False(t, ok3)
	// TypedDeepMetadata
	ret1, ok1 = errors.TypedDeepMetadata[iMeta1](err)
	assert.Equal(t, meta1, ret1)
	assert.True(t, ok1)
	ret2, ok2 = errors.TypedDeepMetadata[iMeta2](err)
	assert.Zero(t, ret2)
	assert.False(t, ok2)
	ret3, ok3 = errors.TypedDeepMetadata[iMeta3](err)
	assert.Zero(t, ret3)
	assert.False(t, ok3)

	err = errors.WithMetadata(err, meta2)
	assert.Equal(t, meta2, errors.Metadata(err))
	// TypedMetadata
	ret1, ok1 = errors.TypedMetadata[iMeta1](err)
	assert.Zero(t, ret1)
	assert.False(t, ok1)
	ret2, ok2 = errors.TypedMetadata[iMeta2](err)
	assert.Equal(t, meta2, ret2)
	assert.True(t, ok2)
	ret3, ok3 = errors.TypedMetadata[iMeta3](err)
	assert.Zero(t, ret3)
	assert.False(t, ok3)
	// TypedDeepMetadata
	ret1, ok1 = errors.TypedDeepMetadata[iMeta1](err)
	assert.Equal(t, meta1, ret1)
	assert.True(t, ok1)
	ret2, ok2 = errors.TypedDeepMetadata[iMeta2](err)
	assert.Equal(t, meta2, ret2)
	assert.True(t, ok2)
	ret3, ok3 = errors.TypedDeepMetadata[iMeta3](err)
	assert.Zero(t, ret3)
	assert.False(t, ok3)

	err = errors.WithMetadata(err, meta3)
	assert.Equal(t, meta3, errors.Metadata(err))
	// TypedMetadata
	ret1, ok1 = errors.TypedMetadata[iMeta1](err)
	assert.Zero(t, ret1)
	assert.False(t, ok1)
	ret2, ok2 = errors.TypedMetadata[iMeta2](err)
	assert.Zero(t, ret2)
	assert.False(t, ok2)
	ret3, ok3 = errors.TypedMetadata[iMeta3](err)
	assert.Equal(t, meta3, ret3)
	assert.True(t, ok3)
	// TypedDeepMetadata
	ret1, ok1 = errors.TypedDeepMetadata[iMeta1](err)
	assert.Equal(t, meta1, ret1)
	assert.True(t, ok1)
	ret2, ok2 = errors.TypedDeepMetadata[iMeta2](err)
	assert.Equal(t, meta2, ret2)
	assert.True(t, ok2)
	ret3, ok3 = errors.TypedDeepMetadata[iMeta3](err)
	assert.Equal(t, meta3, ret3)
	assert.True(t, ok3)
}
