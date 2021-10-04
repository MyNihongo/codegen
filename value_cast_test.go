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

func TestCastQualType(t *testing.T) {
	const want = `a.(alias.MyType)`

	var sb strings.Builder
	newCastValue(Identifier("a"), "alias", "MyType", false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCastQualTypePointer(t *testing.T) {
	const want = `a.(*alias.MyType)`

	var sb strings.Builder
	newCastValue(Identifier("a"), "alias", "MyType", true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCastNotPointer(t *testing.T) {
	got := newCastValue(Identifier("a"), "alias", "MyType", true).
		isPointer()

	assert.False(t, got)
}

func TestCastSetIsPointerFalse(t *testing.T) {
	got := newCastValue(Identifier("a"), "", "string", true)
	assert.True(t, got.name.isPointer)
	assert.False(t, got.isPointer())

	got.SetIsPointer(false)
	assert.False(t, got.name.isPointer)
	assert.False(t, got.isPointer())
}

func TestCastSetIsPointerTrue(t *testing.T) {
	got := newCastValue(Identifier("a"), "", "string", false)
	assert.False(t, got.name.isPointer)
	assert.False(t, got.isPointer())

	got.SetIsPointer(true)
	assert.True(t, got.name.isPointer)
	assert.False(t, got.isPointer())
}
