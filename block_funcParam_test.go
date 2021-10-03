package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamName(t *testing.T) {
	const want = `name`
	got := Param("name", "typeName").Name()

	assert.Equal(t, want, got)
}

func TestQualParamFullName(t *testing.T) {
	const want = `name`
	got := QualParam("name", "alias", "typeName").Name()

	assert.Equal(t, want, got)
}

func TestParamFullType(t *testing.T) {
	const want = `typeName`
	got := Param("name", "typeName").FullType()

	assert.Equal(t, want, got)
}

func TestQualParamFullType(t *testing.T) {
	const want = `alias.typeName`
	got := QualParam("name", "alias", "typeName").FullType()

	assert.Equal(t, want, got)
}

func TestParamFullTypePointer(t *testing.T) {
	const want = `*typeName`
	got := Param("name", "typeName").Pointer().FullType()

	assert.Equal(t, want, got)
}

func TestQualParamTypePointer(t *testing.T) {
	const want = `*alias.typeName`
	got := QualParam("name", "alias", "typeName").Pointer().FullType()

	assert.Equal(t, want, got)
}

func TestFuncParam(t *testing.T) {
	const want = `name type`

	var sb strings.Builder
	Param("name", "type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamPointer(t *testing.T) {
	const want = `name *type`

	var sb strings.Builder
	Param("name", "type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParam(t *testing.T) {
	const want = `name alias.type`

	var sb strings.Builder
	QualParam("name", "alias", "type").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncQualParamPointer(t *testing.T) {
	const want = `name *alias.type`

	var sb strings.Builder
	QualParam("name", "alias", "type").Pointer().
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamsEmpty(t *testing.T) {
	const want = `()`

	var sb strings.Builder
	params := make([]*ParamDecl, 0)
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamsOne(t *testing.T) {
	const want = `(name1 type)`

	var sb strings.Builder
	params := []*ParamDecl{
		Param("name1", "type"),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParams(t *testing.T) {
	const want = `(name1 type,name2 alias.type,name3 *type,name4 *alias.type)`

	var sb strings.Builder
	params := []*ParamDecl{
		Param("name1", "type"),
		QualParam("name2", "alias", "type"),
		Param("name3", "type").Pointer(),
		QualParam("name4", "alias", "type").Pointer(),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamSetIsPointerTrue(t *testing.T) {
	const want = `(name1 *type)`

	var sb strings.Builder
	params := []*ParamDecl{
		Param("name1", "type").SetIsPointer(true),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncParamSetIsPointerFalse(t *testing.T) {
	const want = `(name1 type)`

	var sb strings.Builder
	params := []*ParamDecl{
		Param("name1", "type").SetIsPointer(false),
	}
	writeParams(&sb, params)

	assert.Equal(t, want, sb.String())
}

func TestFuncGetters(t *testing.T) {
	fixture := QualParam("param", "alias", "MyType")

	assert.Equal(t, fixture.Name(), "param")
	assert.Equal(t, fixture.TypeAlias(), "alias")
	assert.Equal(t, fixture.TypeName(), "MyType")
}
