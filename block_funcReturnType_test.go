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

func TestFuncReturnTypeSetIsPointerTrue(t *testing.T) {
	const want = `*type1`

	var sb strings.Builder
	params := []*ReturnTypeDecl{
		ReturnType("type1").SetIsPointer(true),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypeSetIsPointerFalse(t *testing.T) {
	const want = `type1`

	var sb strings.Builder
	params := []*ReturnTypeDecl{
		ReturnType("type1").SetIsPointer(false),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestReturnParamGetters(t *testing.T) {
	fixture := QualReturnType("alias", "MyType").Pointer()

	assert.Equal(t, "alias", fixture.GetTypeAlias())
	assert.Equal(t, "MyType", fixture.GetTypeName())
	assert.Equal(t, true, fixture.GetIsPointer())
}

func TestReturnIsValidTrue(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			got := QualReturnType(alias, "type").SetIsPointer(pointer).isValid()
			assert.True(t, got)
		}
	}
}

func TestReturnIsValidFalse(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			got := QualReturnType(alias, "").SetIsPointer(pointer).isValid()
			assert.False(t, got)
		}
	}
}
