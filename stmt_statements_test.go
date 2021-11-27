package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStmtsEmpty(t *testing.T) {
	inputs := []Stmt{}

	var sb strings.Builder
	Stmts(inputs).writeStmt(&sb)

	assert.Empty(t, sb.String())
}

func TestStmtsSingle(t *testing.T) {
	const want = `var val string
`

	inputs := []Stmt{
		DeclareVars(Var("val", "string")),
	}

	var sb strings.Builder
	Stmts(inputs).writeStmt(&sb)

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
	Stmts(inputs).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
