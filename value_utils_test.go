package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteValuesEmpty(t *testing.T) {
	var sb strings.Builder
	writeValues(&sb, make([]Value, 0))

	assert.Empty(t, sb.String())
}

func TestWriteValuesSingle(t *testing.T) {
	const want = `name.field`

	vals := []Value{
		Identifier("name").Field("field"),
	}

	var sb strings.Builder
	writeValues(&sb, vals)

	assert.Equal(t, want, sb.String())
}

func TestWriteValuesMultiple(t *testing.T) {
	const want = `name.field,alias.myFunc(a,b)`

	vals := []Value{
		Identifier("name").Field("field"),
		QualFuncCall("alias", "myFunc").Args(Identifier("a"), Identifier("b")),
	}

	var sb strings.Builder
	writeValues(&sb, vals)

	assert.Equal(t, want, sb.String())
}

func TestWritePointerAccessNoPointer(t *testing.T) {
	const want = `myFunc(a)`

	val := FuncCall("myFunc").Args(Identifier("a"))

	var sb strings.Builder
	writePointerValueAccess(&sb, val)

	assert.Equal(t, want, sb.String())
}

func TestWritePointerAccessPointer(t *testing.T) {
	const want = `(*myFunc(a))`

	val := FuncCall("myFunc").Args(Identifier("a")).Pointer()

	var sb strings.Builder
	writePointerValueAccess(&sb, val)

	assert.Equal(t, want, sb.String())
}

func TestWritePointerAccessStructInit(t *testing.T) {
	const want = `myStruct{prop:a.someFunc()}`

	val := InitStruct("myStruct").Props(
		PropValue("prop", Identifier("a").Call("someFunc")),
	)

	var sb strings.Builder
	writePointerValueAccess(&sb, val)

	assert.Equal(t, want, sb.String())
}

func TestWritePointerAccessStructAddressInit(t *testing.T) {
	const want = `&myStruct{prop:a.someFunc()}`

	val := InitStruct("myStruct").Props(
		PropValue("prop", Identifier("a").Call("someFunc")),
	).Address()

	var sb strings.Builder
	writePointerValueAccess(&sb, val)

	assert.Equal(t, want, sb.String())
}

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
