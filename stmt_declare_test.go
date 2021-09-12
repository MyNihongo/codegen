package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclareOneVariable(t *testing.T) {
	const want = `newVar:=alias.ConstVal`

	var sb strings.Builder
	got := Declare("newVar").Values(QualIdentifier("alias", "ConstVal")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestAssignOneVariable(t *testing.T) {
	const want = `newVar=alias.ConstVal`

	var sb strings.Builder
	got := Assign("newVar").Values(QualIdentifier("alias", "ConstVal")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestDeclareMultipleVariables(t *testing.T) {
	const want = `var1,var2:=alias.ConstVal,alias.Value2`

	var sb strings.Builder
	got := Declare("var1", "var2").Values(QualIdentifier("alias", "ConstVal"), QualIdentifier("alias", "Value2")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestAssignMultipleVariables(t *testing.T) {
	const want = `var1,var2=alias.ConstVal,alias.Value2`

	var sb strings.Builder
	got := Assign("var1", "var2").Values(QualIdentifier("alias", "ConstVal"), QualIdentifier("alias", "Value2")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestDeclareFunc(t *testing.T) {
	const want = `varr,err:=execFunc(anotherVar)`

	var sb strings.Builder
	got := Declare("varr", "err").Values(FuncCall("execFunc").Args(Identifier("anotherVar"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestAssignFunc(t *testing.T) {
	const want = `varr,err=execFunc(anotherVar)`

	var sb strings.Builder
	got := Assign("varr", "err").Values(FuncCall("execFunc").Args(Identifier("anotherVar"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestDeclarationPanicNoVars(t *testing.T) {
	assert.Panics(t, func() {
		Declare()
	})
}

func TestAssignmentPanicNoVars(t *testing.T) {
	assert.Panics(t, func() {
		Assign()
	})
}

func TestPanicNoValues(t *testing.T) {
	assert.Panics(t, func() {
		Declare("varr").Values()
	})
}
