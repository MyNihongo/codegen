package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapTypeKeyArrayPanics(t *testing.T) {
	assert.Panics(t, func() {
		key, value := Type("string").Array(), Type("string")
		MapType(key, value)
	})
}

func TestMapTypeRegular(t *testing.T) {
	const want = `map[string]string`

	var sb strings.Builder
	key, value := Type("string"), Type("string")
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMapTypeQualRegular(t *testing.T) {
	const want = `map[alias.MyType]alias.SecondType`

	var sb strings.Builder
	key, value := QualType("alias", "MyType"), QualType("alias", "SecondType")
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMapKeyPointer(t *testing.T) {
	const want = `map[*string]string`

	var sb strings.Builder
	key, value := Type("string").Pointer(), Type("string")
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMapValuePointer(t *testing.T) {
	const want = `map[string]*string`

	var sb strings.Builder
	key, value := Type("string"), Type("string").Pointer()
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMapValueArray(t *testing.T) {
	const want = `map[string][]string`

	var sb strings.Builder
	key, value := Type("string"), Type("string").Array()
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMapValuePointerArray(t *testing.T) {
	const want = `map[string][]*string`

	var sb strings.Builder
	key, value := Type("string"), Type("string").Pointer().Array()
	MapType(key, value).wr(&sb)

	assert.Equal(t, want, sb.String())
}
