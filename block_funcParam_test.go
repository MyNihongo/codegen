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

func TestQualParamName(t *testing.T) {
	const want = `name`
	got := QualParam("name", "alias", "typeName").Name()

	assert.Equal(t, want, got)
}

func TestParamType(t *testing.T) {
	const want = `typeName`
	got := Param("name", "typeName").TypeName()

	assert.Equal(t, want, got)
}

func TestQualParamType(t *testing.T) {
	const want = `alias.typeName`
	got := QualParam("name", "alias", "typeName").TypeName()

	assert.Equal(t, want, got)
}

func TestParamTypePointer(t *testing.T) {
	const want = `*typeName`
	got := Param("name", "typeName").Pointer().TypeName()

	assert.Equal(t, want, got)
}

func TestQualParamTypePointer(t *testing.T) {
	const want = `*alias.typeName`
	got := QualParam("name", "alias", "typeName").Pointer().TypeName()

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
