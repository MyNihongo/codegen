package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	const want = `obj`

	var sb strings.Builder
	Identifier("obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierPointer(t *testing.T) {
	const want = `*obj`

	var sb strings.Builder
	Identifier("obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifier(t *testing.T) {
	const want = `alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierPointer(t *testing.T) {
	const want = `*alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierEquals(t *testing.T) {
	const want = `a==myFunc()`

	var sb strings.Builder
	Identifier("a").Equals(FuncCall("myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierEquals(t *testing.T) {
	const want = `alias.Var==b`

	var sb strings.Builder
	QualIdentifier("alias", "Var").Equals(Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNotEquals(t *testing.T) {
	const want = `a!=myFunc()`

	var sb strings.Builder
	Identifier("a").NotEquals(FuncCall("myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNil(t *testing.T) {
	const want = `a==nil`

	var sb strings.Builder
	Identifier("a").Nil().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNotNil(t *testing.T) {
	const want = `a!=nil`

	var sb strings.Builder
	Identifier("a").NotNil().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
