package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallArgsEmpty(t *testing.T) {
	const want = `a.myFunc()`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallArgsSingle(t *testing.T) {
	const want = `a.myFunc(b)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").Args(Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallArgsMultiple(t *testing.T) {
	const want = `a.myFunc(b,"str value")`

	var sb strings.Builder
	Identifier("a").Call("myFunc").Args(Identifier("b"), String("str value")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallPointerArgs(t *testing.T) {
	const want = `(*a).myFunc()`

	var sb strings.Builder
	Identifier("a").Pointer().Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallField(t *testing.T) {
	const want = `a.myFunc().field`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallCall(t *testing.T) {
	const want = `a.myFunc(a).someFunc(a)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").Args(Identifier("a")).
		Call("someFunc").Args(Identifier("a")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallAsStmt(t *testing.T) {
	const want = `a.myFunc()
`
	var sb strings.Builder
	got := Identifier("a").Call("myFunc").
		writeStmt(&sb)

	assert.False(t, got)
	assert.Equal(t, want, sb.String())
}

func TestCallCast(t *testing.T) {
	const want = `a.myFunc().(string)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		Cast("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallCastPointer(t *testing.T) {
	const want = `a.myFunc().(*string)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		CastPointer("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallCastQual(t *testing.T) {
	const want = `a.myFunc().(alias.MyType)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		CastQual("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallCastQualPointer(t *testing.T) {
	const want = `a.myFunc().(*alias.MyType)`

	var sb strings.Builder
	Identifier("a").Call("myFunc").
		CastQualPointer("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
