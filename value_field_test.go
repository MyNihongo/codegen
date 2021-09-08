package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifierNestedField(t *testing.T) {
	const want = `obj.field1.field2`

	var sb strings.Builder
	Identifier("obj").
		Field("field1").Field("field2").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldAssign(t *testing.T) {
	const want = `obj.field1=myFunc(a,b)`

	var sb strings.Builder
	Identifier("obj").
		Field("field1").Assign(FuncCall("myFunc").Args(Identifier("a"), Identifier("b"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldCallEmpty(t *testing.T) {
	const want = `obj.field.myFunc()`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldCallArgSingle(t *testing.T) {
	const want = `obj.field.myFunc(a)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").Args(Identifier("a")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldCallArgs(t *testing.T) {
	const want = `obj.field.myFunc(a,anotherFunc(b))`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").Args(Identifier("a"), FuncCall("anotherFunc").Args(Identifier("b"))).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
