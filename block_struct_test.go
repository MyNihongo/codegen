package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructEmpty(t *testing.T) {
	const want = `type myStruct struct {
}
`
	var sb strings.Builder
	newStruct("myStruct").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructOne(t *testing.T) {
	const want = `type myStruct struct {
prop string
}
`
	var sb strings.Builder
	newStruct("myStruct").Props(
		Property("prop", "string"),
	).write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructMultiple(t *testing.T) {
	const want = `type myStruct struct {
prop string
prop alias.MyType
}
`
	var sb strings.Builder
	newStruct("myStruct").Props(
		Property("prop", "string"),
		QualProperty("prop", "alias", "MyType"),
	).write(&sb)

	assert.Equal(t, want, sb.String())
}
