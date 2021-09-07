package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldIdentifier(t *testing.T) {
	const want = `val.field`

	var sb strings.Builder
	Identifier("val").Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldIdentifierPointer(t *testing.T) {
	const want = `(*val).field`

	var sb strings.Builder
	Identifier("val").Pointer().
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldQualIdentifier(t *testing.T) {
	const want = `alias.val.field`

	var sb strings.Builder
	QualIdentifier("alias", "val").Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldQualIdentifierPointer(t *testing.T) {
	const want = `(*alias.val).field`

	var sb strings.Builder
	QualIdentifier("alias", "val").Pointer().
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNestedField(t *testing.T) {
	const want = `obj.field1.field2`

	var sb strings.Builder
	Identifier("obj").
		Field("field1").Field("field2").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
