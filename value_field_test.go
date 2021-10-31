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
	got := Identifier("obj").
		Field("field1").Assign(FuncCall("myFunc").Args(Identifier("a"), Identifier("b"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
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

func TestFieldCast(t *testing.T) {
	const want = `obj.field.(string)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Cast("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldCastPointer(t *testing.T) {
	const want = `obj.field.(*string)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").CastPointer("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldCastQual(t *testing.T) {
	const want = `obj.field.(alias.MyType)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").CastQual("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldCastQualPointer(t *testing.T) {
	const want = `obj.field.(*alias.MyType)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").CastQualPointer("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldEquals(t *testing.T) {
	const want = `obj.field==field`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Equals(Identifier("field")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldNotEquals(t *testing.T) {
	const want = `obj.field!=field`

	var sb strings.Builder
	Identifier("obj").
		Field("field").NotEquals(Identifier("field")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFieldAddress(t *testing.T) {
	const want = `&(obj.field)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Address().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
