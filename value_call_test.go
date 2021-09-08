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
