package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceEmpty(t *testing.T) {
	const want = `type myInterface interface {
}
`
	var sb strings.Builder
	newInterface("myInterface").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceOne(t *testing.T) {
	const want = `type myInterface interface {
myFunc(param string)string
}
`
	var sb strings.Builder
	newInterface("myInterface").Funcs(
		FuncDecl("myFunc").Params(Param("param", "string")).ReturnTypes(ReturnType("string")),
	).write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceMultiple(t *testing.T) {
	const want = `type myInterface interface {
myFunc(param string)string
anotherFunc(param int64)alias.MyType
}
`
	var sb strings.Builder
	newInterface("myInterface").Funcs(
		FuncDecl("myFunc").Params(Param("param", "string")).ReturnTypes(ReturnType("string")),
		FuncDecl("anotherFunc").Params(Param("param", "int64")).ReturnTypes(QualReturnType("alias", "MyType")),
	).write(&sb)

	assert.Equal(t, want, sb.String())
}
