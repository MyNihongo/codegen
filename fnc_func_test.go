package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncEmpty(t *testing.T) {
	const want = `func funcName(){
}
`
	var sb strings.Builder
	newFunc("funcName").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncWithParams(t *testing.T) {
	const want = `func funcName(name typeName,name2 *alias.typeName){
}
`
	param1, param2 := NewParam("name", "typeName"),
		NewQualParam("name2", NewQualName("alias", "typeName")).Pointer()

	var sb strings.Builder
	newFunc("funcName").
		Params(param1, param2).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncWithReturnType(t *testing.T) {
	const want = `func funcName()type{
}
`
	retType := NewReturnType("type")

	var sb strings.Builder
	newFunc("funcName").
		ReturnTypes(retType).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncWithReturnTypes(t *testing.T) {
	const want = `func funcName()(type,*alias.type){
}
`
	retType1, retType2 := NewReturnType("type"),
		NewQualReturnType("alias", "type").Pointer()

	var sb strings.Builder
	newFunc("funcName").
		ReturnTypes(retType1, retType2).
		write(&sb)

	assert.Equal(t, want, sb.String())
}
