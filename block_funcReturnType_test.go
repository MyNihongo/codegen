package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncReturnType(t *testing.T) {
	const want = `type`

	var sb strings.Builder
	ReturnType("type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnTypeError(t *testing.T) {
	const want = `error`

	var sb strings.Builder
	ReturnTypeError().wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypePointer(t *testing.T) {
	const want = `*type`

	var sb strings.Builder
	ReturnType("type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnType(t *testing.T) {
	const want = `alias.type`

	var sb strings.Builder
	QualReturnType("alias", "type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnTypePointer(t *testing.T) {
	const want = `*alias.type`

	var sb strings.Builder
	QualReturnType("alias", "type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypesEmpty(t *testing.T) {
	var sb strings.Builder
	params := make([]*ReturnTypeDecl, 0)
	writeReturnTypes(&sb, params)

	assert.Empty(t, sb.String())
}

func TestFuncReturnTypesOne(t *testing.T) {
	const want = `type1`

	var sb strings.Builder
	params := []*ReturnTypeDecl{
		ReturnType("type1"),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypes(t *testing.T) {
	const want = `(type1,alias.type2,*type3,*alias.type4)`

	var sb strings.Builder
	params := []*ReturnTypeDecl{
		ReturnType("type1"),
		QualReturnType("alias", "type2"),
		ReturnType("type3").Pointer(),
		QualReturnType("alias", "type4").Pointer(),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}
