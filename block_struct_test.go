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
		decl.AddProp(fmt.Sprintf("prop%d", i+1), "string")
	}

	var sb strings.Builder
	decl.write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestStructQualAddProp(t *testing.T) {
	const want = `type myStruct struct {
	prop1 alias.MyType1
	prop2 alias.MyType2
	prop3 alias.MyType3
	prop4 alias.MyType4
	prop5 alias.MyType5
}
`
	decl := newStruct("myStruct")

	for i := 0; i < 5; i++ {
		name, typ := fmt.Sprintf("prop%d", i+1), fmt.Sprintf("MyType%d", i+1)
		decl.AddQualProp(name, "alias", typ)
	}

	var sb strings.Builder
	decl.write(&sb)

	assert.Equal(t, want, formatSb(sb))
}
