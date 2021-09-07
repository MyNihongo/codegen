package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnEmpty(t *testing.T) {
	const want = `return`

	var sb strings.Builder
	NewReturn().writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnOne(t *testing.T) {
	const want = `return variable`
	id := NewIdentifier("variable")

	var sb strings.Builder
	NewReturn(id).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnMultiple(t *testing.T) {
	const want = `return variable,alias.val.field`
	id1, id2 := NewIdentifier("variable"), NewQualIdentifier("alias", "val").Field("field")

	var sb strings.Builder
	NewReturn(id1, id2).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
