package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncReturnType(t *testing.T) {
	const want = `type`

	var sb strings.Builder
	NewReturnType("type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypePointer(t *testing.T) {
	const want = `*type`

	var sb strings.Builder
	NewReturnType("type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnType(t *testing.T) {
	const want = `alias.type`

	var sb strings.Builder
	NewQualReturnType("alias", "type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnTypePointer(t *testing.T) {
	const want = `*alias.type`

	var sb strings.Builder
	NewQualReturnType("alias", "type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypesEmpty(t *testing.T) {
	var sb strings.Builder
	params := make([]*returnType, 0)
	writeReturnTypes(&sb, params)

	assert.Empty(t, sb.String())
}

func TestFuncReturnTypes(t *testing.T) {
	const want = `(type1,alias.type2,*type3,*alias.type4)`

	var sb strings.Builder
	params := []*returnType{
		NewReturnType("type1"),
		NewQualReturnType("alias", "type2"),
		NewReturnType("type3").Pointer(),
		NewQualReturnType("alias", "type4").Pointer(),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}
