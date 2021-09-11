package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclareVarsNone(t *testing.T) {
	var sb strings.Builder
	got := DeclareVars().writeStmt(&sb)

	assert.Empty(t, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsVar(t *testing.T) {
	const want = `var val string
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val", "string"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualVar(t *testing.T) {
	const want = `var val alias.MyType
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val", "alias", "MyType"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsVarPointer(t *testing.T) {
	const want = `var val *string
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val", "string").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualVarPointer(t *testing.T) {
	const want = `var val *alias.MyType
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val", "alias", "MyType").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsDifferentType(t *testing.T) {
	const want = `var val1 string
var val2 int
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val1", "string"),
		Var("val2", "int"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsSameType(t *testing.T) {
	const want = `var val1,val2 string
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val1", "string"),
		Var("val2", "string"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsDifferentTypePointer(t *testing.T) {
	const want = `var val1 *string
var val2 *int
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val1", "string").Pointer(),
		Var("val2", "int").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsSameTypePointer(t *testing.T) {
	const want = `var val1,val2 *string
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val1", "string").Pointer(),
		Var("val2", "string").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualDifferentType(t *testing.T) {
	const want = `var val1 alias.MyType1
var val2 alias.MyType2
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val1", "alias", "MyType1"),
		QualVar("val2", "alias", "MyType2"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualSameType(t *testing.T) {
	const want = `var val1,val2 alias.MyType
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val1", "alias", "MyType"),
		QualVar("val2", "alias", "MyType"),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualDifferentTypePointer(t *testing.T) {
	const want = `var val1 *alias.MyType1
var val2 *alias.MyType2
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val1", "alias", "MyType1").Pointer(),
		QualVar("val2", "alias", "MyType2").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsQualSameTypePointer(t *testing.T) {
	const want = `var val1,val2 *alias.MyType
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val1", "alias", "MyType").Pointer(),
		QualVar("val2", "alias", "MyType").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareVarsPointerNotPointer(t *testing.T) {
	const want = `var val1 string
var val2 *string
`
	var sb strings.Builder
	got := DeclareVars(
		Var("val1", "string"),
		Var("val2", "string").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestDeclareQualVarsPointerNotPointer(t *testing.T) {
	const want = `var val1 alias.MyType
var val2 *alias.MyType
`
	var sb strings.Builder
	got := DeclareVars(
		QualVar("val1", "alias", "MyType"),
		QualVar("val2", "alias", "MyType").Pointer(),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}
