package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualsIdentifierIdentifier(t *testing.T) {
	const want = `a==b.field`

	var sb strings.Builder
	newEquals(Identifier("a"), Identifier("b").Field("field"), true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNotEqualsIdentifierIdentifier(t *testing.T) {
	const want = `a!=b.field`

	var sb strings.Builder
	newEquals(Identifier("a"), Identifier("b").Field("field"), false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEqualsIdentifierFunc(t *testing.T) {
	const want = `a==myFunc()`

	var sb strings.Builder
	newEquals(Identifier("a"), FuncCall("myFunc"), true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNotEqualsIdentifierFunc(t *testing.T) {
	const want = `a!=myFunc()`

	var sb strings.Builder
	newEquals(Identifier("a"), FuncCall("myFunc"), false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEqualsFuncIdentifier(t *testing.T) {
	const want = `myFunc()==a`

	var sb strings.Builder
	newEquals(FuncCall("myFunc"), Identifier("a"), true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNotEqualsFuncIdentifier(t *testing.T) {
	const want = `myFunc()!=b.field`

	var sb strings.Builder
	newEquals(FuncCall("myFunc"), Identifier("b").Field("field"), false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEqualsFuncFunc(t *testing.T) {
	const want = `alias.MyFunc()==myFunc()`

	var sb strings.Builder
	newEquals(QualFuncCall("alias", "MyFunc"), FuncCall("myFunc"), true).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestNotEqualsFuncFunc(t *testing.T) {
	const want = `alias.MyFunc()!=myFunc()`

	var sb strings.Builder
	newEquals(QualFuncCall("alias", "MyFunc"), FuncCall("myFunc"), false).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestEqualsNoPointer(t *testing.T) {
	got := newEquals(Identifier("a"), Identifier("b"), true).
		isPointer()

	assert.False(t, got)
}
