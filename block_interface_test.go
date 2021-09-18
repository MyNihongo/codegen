package codegen

import (
	"fmt"
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

func TestInterfaceAddFunc(t *testing.T) {
	const want = `type myInterface interface {
myFunc1()
myFunc2(a1 string)
myFunc3(a1 string,a2 string)
myFunc4(a1 string,a2 string,a3 string)
myFunc5(a1 string,a2 string,a3 string,a4 string)
}
`
	decl := newInterface("myInterface")

	for i := 0; i < 5; i++ {
		params := make([]*ParamDecl, i)

		for j := 0; j < i; j++ {
			params[j] = Param(fmt.Sprintf("a%d", j+1), "string")
		}

		newFunc := FuncDecl(fmt.Sprintf("myFunc%d", i+1)).Params(params...)
		decl.AddFunc(newFunc)
	}

	var sb strings.Builder
	decl.write(&sb)

	assert.Equal(t, want, sb.String())
}
