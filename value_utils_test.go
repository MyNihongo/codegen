package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteAlias(t *testing.T) {
	const want = `alias.`

	var sb strings.Builder
	writeAlias(&sb, "alias")

	assert.Equal(t, want, sb.String())
}

func TestWriteAliasEmpty(t *testing.T) {
	var sb strings.Builder
	writeAlias(&sb, "")

	assert.Empty(t, sb.String())
}

func TestWriteValuesEmpty(t *testing.T) {
	var sb strings.Builder
	writeValues(&sb, make([]value, 0))

	assert.Empty(t, sb.String())
}

func TestWriteValuesSingle(t *testing.T) {
	const want = `name.field`

	vals := []value{
		Identifier("name").Field("field"),
	}

	var sb strings.Builder
	writeValues(&sb, vals)

	assert.Equal(t, want, sb.String())
}

func TestWriteValuesMultiple(t *testing.T) {
	const want = `name.field,alias.myFunc(a,b)`

	vals := []value{
		Identifier("name").Field("field"),
		QualFuncCall("alias", "myFunc").Args(Identifier("a"), Identifier("b")),
	}

	var sb strings.Builder
	writeValues(&sb, vals)

	assert.Equal(t, want, sb.String())
}
