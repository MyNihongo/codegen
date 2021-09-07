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
	stmts := []stmt{
		NewReturn(NewIdentifier("val")),
	}

	var sb strings.Builder
	writeStmts(&sb, stmts, false)

	assert.Equal(t, want, sb.String())
}

func TestStatementOneNewLine(t *testing.T) {
	const want = `{
return val
}
`
	stmts := []stmt{
		NewReturn(NewIdentifier("val")),
	}

	var sb strings.Builder
	writeStmts(&sb, stmts, true)

	assert.Equal(t, want, sb.String())
}
