package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestType(t *testing.T) {
	const want = `type myType int64
`
	var sb strings.Builder
	newType("myType", "", "int64").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestTypePointer(t *testing.T) {
	const want = `type myType *int64
`
	var sb strings.Builder
	newType("myType", "", "int64").Pointer().
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualType(t *testing.T) {
	const want = `type myType alias.MyType
`
	var sb strings.Builder
	newType("myType", "alias", "MyType").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualTypePointer(t *testing.T) {
	const want = `type myType *alias.MyType
`
	var sb strings.Builder
	newType("myType", "alias", "MyType").Pointer().
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestTypeSetIsPointerTrue(t *testing.T) {
	const want = `type myVar *myType
`
	var sb strings.Builder
	newType("myVar", "", "myType").SetIsPointer(true).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestTypeSetIsPointerFalse(t *testing.T) {
	const want = `type myVar myType
`
	var sb strings.Builder
	newType("myVar", "", "myType").SetIsPointer(false).
		write(&sb)

	assert.Equal(t, want, sb.String())
}
