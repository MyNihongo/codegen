package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquireBrackets(t *testing.T) {
	const want = `arr[0]`

	var sb strings.Builder
	newSquireBrackets(Identifier("arr"), Int(0)).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestSquireBracketsPointer(t *testing.T) {
	got := newSquireBrackets(Identifier("arr"), Int(0)).
		isPointer()

	assert.False(t, got)
}

func TestSquireBracketsAssign(t *testing.T) {
	const want = `myMap["key"]=myFunc()`

	var sb strings.Builder
	newSquireBrackets(Identifier("myMap"), String("key")).
		Assign(FuncCall("myFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
