package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructNoProperties(t *testing.T) {
	const want = `myStruct{}`

	var sb strings.Builder
	InitStruct("myStruct").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructNoPropertiesAddress(t *testing.T) {
	const want = `&myStruct{}`

	var sb strings.Builder
	InitStruct("myStruct").Address().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructOneProp(t *testing.T) {
	const want = `myStruct{prop:val}`

	var sb strings.Builder
	InitStruct("myStruct").Props(
		PropValue("prop", Identifier("val")),
	).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructMultipleProps(t *testing.T) {
	const want = `myStruct{prop1:val1,prop2:val2}`

	var sb strings.Builder
	InitStruct("myStruct").Props(
		PropValue("prop1", Identifier("val1")),
		PropValue("prop2", Identifier("val2")),
	).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
