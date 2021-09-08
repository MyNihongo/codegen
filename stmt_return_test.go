package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnEmpty(t *testing.T) {
	const want = `return`

	var sb strings.Builder
	Return().writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnOne(t *testing.T) {
	const want = `return variable`
	id := Identifier("variable")

	var sb strings.Builder
	Return(id).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestReturnMultiple(t *testing.T) {
	const want = `return variable,alias.val.field`
	id1, id2 := Identifier("variable"), QualIdentifier("alias", "val").Field("field")

	var sb strings.Builder
	Return(id1, id2).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
