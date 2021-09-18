package codegen

import (
	"fmt"
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

func TestStructAddProp(t *testing.T) {
	const want = `type myStruct struct {
prop1 string
prop2 string
prop3 string
prop4 string
prop5 string
}
`
	decl := newStruct("myStruct")

	for i := 0; i < 5; i++ {
		newProp := Property(fmt.Sprintf("prop%d", i+1), "string")
		decl.AddProp(newProp)
	}

	var sb strings.Builder
	decl.write(&sb)

	assert.Equal(t, want, sb.String())
}
