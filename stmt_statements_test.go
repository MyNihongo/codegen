package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStmtsEmpty(t *testing.T) {
	inputs := []Stmt{}

	var sb strings.Builder
	got := Stmts(inputs).writeStmt(&sb)

	assert.False(t, got)
	assert.Empty(t, sb.String())
}

func TestStmtsSingle(t *testing.T) {
	const want = `var val string
`
	inputs := []Stmt{
		DeclareVars(Var("val", "string")),
	}

	var sb strings.Builder
	got := Stmts(inputs).writeStmt(&sb)

	assert.False(t, got)
	assert.Equal(t, want, sb.String())
}

func TestStmtsMultiple(t *testing.T) {
	const want = `var val string
if myVar{
return
}
`
	inputs := []Stmt{
		DeclareVars(Var("val", "string")),
		If(Identifier("myVar")).Block(
			Return(),
		),
	}

	var sb strings.Builder
	got := Stmts(inputs).writeStmt(&sb)

	assert.False(t, got)
	assert.Equal(t, want, sb.String())
}
