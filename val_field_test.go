package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldIdentifier(t *testing.T) {
	const want = `val.field`

	var sb strings.Builder
	NewIdentifier("val").Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldIdentifierPointer(t *testing.T) {
	const want = `(*val).field`

	var sb strings.Builder
	NewIdentifier("val").Pointer().
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
