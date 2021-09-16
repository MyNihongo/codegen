package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterfaceFuncEmpty(t *testing.T) {
	const want = `myFunc()`

	var sb strings.Builder
	FuncDecl("myFunc").wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncParamOne(t *testing.T) {
	const want = `myFunc(param myType)`

	var sb strings.Builder
	FuncDecl("myFunc").Params(Param("param", "myType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncQualParamOne(t *testing.T) {
	const want = `myFunc(param alias.MyType)`

	var sb strings.Builder
	FuncDecl("myFunc").Params(QualParam("param", "alias", "MyType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncParamMultiple(t *testing.T) {
	const want = `myFunc(param myType,anotherParam myType2)`

	var sb strings.Builder
	FuncDecl("myFunc").Params(Param("param", "myType"), Param("anotherParam", "myType2")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncQualParamMultiple(t *testing.T) {
	const want = `myFunc(param alias.MyType,anotherParam alias.MyType2)`

	var sb strings.Builder
	FuncDecl("myFunc").Params(QualParam("param", "alias", "MyType"), QualParam("anotherParam", "alias", "MyType2")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncReturnTypeOne(t *testing.T) {
	const want = `myFunc()myType`

	var sb strings.Builder
	FuncDecl("myFunc").ReturnTypes(ReturnType("myType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncQualReturnTypeOne(t *testing.T) {
	const want = `myFunc()alias.MyType`

	var sb strings.Builder
	FuncDecl("myFunc").ReturnTypes(QualReturnType("alias", "MyType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncReturnTypeMultiple(t *testing.T) {
	const want = `myFunc()(myType,anotherType)`

	var sb strings.Builder
	FuncDecl("myFunc").ReturnTypes(ReturnType("myType"), ReturnType("anotherType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}

func TestInterfaceFuncQualReturnTypeMultiple(t *testing.T) {
	const want = `myFunc()(alias.MyType,alias.AnotherType)`

	var sb strings.Builder
	FuncDecl("myFunc").ReturnTypes(QualReturnType("alias", "MyType"), QualReturnType("alias", "AnotherType")).
		wr(&sb)

	assert.Equal(t, want, sb.String())
}
