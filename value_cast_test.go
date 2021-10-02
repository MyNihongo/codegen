package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCastType(t *testing.T) {
	const want = `a.(string)`

	var sb strings.Builder
	newCastValue(Identifier("a"), "", "string", false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCastTypePointer(t *testing.T) {
	const want = `a.(*string)`

	var sb strings.Builder
	newCastValue(Identifier("a"), "", "string", true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
