package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatementOne(t *testing.T) {
	const want = `{
return val
}`
	stmts := []Stmt{
		Return(Identifier("val")),
	}

	var sb strings.Builder
	writeStmtsBlock(&sb, stmts, false)

	assert.Equal(t, want, sb.String())
}

func TestStatementOneNewLine(t *testing.T) {
	const want = `{
return val
}
`
	stmts := []Stmt{
		Return(Identifier("val")),
	}

	var sb strings.Builder
	writeStmtsBlock(&sb, stmts, true)

	assert.Equal(t, want, sb.String())
}
