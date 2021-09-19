package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallHelperEmpty(t *testing.T) {
	const want = `myFunc()`

	var sb strings.Builder
	newCallHelper("myFunc", []Value{}).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallHelperSingle(t *testing.T) {
	const want = `myFunc(a)`
	val := Identifier("a")

	var sb strings.Builder
	newCallHelper("myFunc", []Value{val}).wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestCallHelperMultiple(t *testing.T) {
	const want = `myFunc(a,anotherFunc(b))`
	val1, val2 := Identifier("a"), FuncCall("anotherFunc").Args(Identifier("b"))

	var sb strings.Builder
	newCallHelper("myFunc", []Value{val1, val2}).wr(&sb)

	assert.Equal(t, want, sb.String())
}
