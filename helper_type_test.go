package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncReturnType(t *testing.T) {
	const want = `type`

	var sb strings.Builder
	Type("type").wr(&sb)

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
	Type("type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnType(t *testing.T) {
	const want = `alias.type`

	var sb strings.Builder
	QualType("alias", "type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualReturnTypePointer(t *testing.T) {
	const want = `*alias.type`

	var sb strings.Builder
	QualType("alias", "type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypesEmpty(t *testing.T) {
	var sb strings.Builder
	params := make([]*TypeDecl, 0)
	writeReturnTypes(&sb, params)

	assert.Empty(t, sb.String())
}

func TestFuncReturnTypesOne(t *testing.T) {
	const want = `type1`

	var sb strings.Builder
	params := []*TypeDecl{
		Type("type1"),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypes(t *testing.T) {
	const want = `(type1,alias.type2,*type3,*alias.type4)`

	var sb strings.Builder
	params := []*TypeDecl{
		Type("type1"),
		QualType("alias", "type2"),
		Type("type3").Pointer(),
		QualType("alias", "type4").Pointer(),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypeSetIsPointerTrue(t *testing.T) {
	const want = `*type1`

	var sb strings.Builder
	params := []*TypeDecl{
		Type("type1").SetIsPointer(true),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncReturnTypeSetIsPointerFalse(t *testing.T) {
	const want = `type1`

	var sb strings.Builder
	params := []*TypeDecl{
		Type("type1").SetIsPointer(false),
	}
	writeReturnTypes(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestReturnParamGetters(t *testing.T) {
	fixture := QualType("alias", "MyType").Pointer()

	assert.Equal(t, "alias", fixture.GetTypeAlias())
	assert.Equal(t, "MyType", fixture.GetTypeName())
	assert.Equal(t, true, fixture.GetIsPointer())
}

func TestReturnIsValidTrue(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			got := QualType(alias, "type").SetIsPointer(pointer).isValid()
			assert.True(t, got)
		}
	}
}

func TestReturnIsValidFalse(t *testing.T) {
	aliases, pointers := []string{"", "alias"}, []bool{true, false}

	for _, alias := range aliases {
		for _, pointer := range pointers {
			got := QualType(alias, "").SetIsPointer(pointer).isValid()
			assert.False(t, got)
		}
	}
}

func TestReturnTypeOneNotValid(t *testing.T) {
	fixture := []*TypeDecl{
		Type(""),
	}

	var sb strings.Builder
	writeReturnTypes(&sb, fixture)

	assert.Empty(t, sb.String())
}

func TestReturnTypeMultipleNotValid(t *testing.T) {
	fixture := []*TypeDecl{
		Type(""),
		QualType("alias", ""),
	}

	var sb strings.Builder
	writeReturnTypes(&sb, fixture)

	assert.Empty(t, sb.String())
}

func TestReturnTypeMultipleNotValidOneValid(t *testing.T) {
	const want = `(string)`

	fixture := []*TypeDecl{
		Type(""),
		Type("string"),
		QualType("alias", ""),
	}

	var sb strings.Builder
	writeReturnTypes(&sb, fixture)

	assert.Equal(t, want, sb.String())
}

func TestReturnTypeMultipleNotValidMultipleValid(t *testing.T) {
	const want = `(string,alias.MyType)`

	fixture := []*TypeDecl{
		Type(""),
		Type("string"),
		QualType("alias", ""),
		QualType("alias", "MyType"),
	}

	var sb strings.Builder
	writeReturnTypes(&sb, fixture)

	assert.Equal(t, want, sb.String())
}

func TestReturnTypeArray(t *testing.T) {
	const want = `[]string`

	var sb strings.Builder
	Type("string").Array().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnTypePointerArray(t *testing.T) {
	const want = `[]*string`

	var sb strings.Builder
	Type("string").Array().Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualReturnTypeArray(t *testing.T) {
	const want = `[]alias.MyType`

	var sb strings.Builder
	QualType("alias", "MyType").Array().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualReturnTypePointerArray(t *testing.T) {
	const want = `[]*alias.MyType`

	var sb strings.Builder
	QualType("alias", "MyType").Array().Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestTypeSetIsArrayTrue(t *testing.T) {
	const want = `[]string`

	var sb strings.Builder
	Type("string").SetIsArray(true).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestTypeSetIsArrayFalse(t *testing.T) {
	const want = `string`

	var sb strings.Builder
	Type("string").Array().SetIsArray(false).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}
