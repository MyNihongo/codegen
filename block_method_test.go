package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMethodThis(t *testing.T) {
	const want = `func (t type) funcName(){
}
`
	var sb strings.Builder
	this := This("type")

	newMethod(this, "funcName").write(&sb)
	assert.Equal(t, want, sb.String())
}

func TestMethodThisPointer(t *testing.T) {
	const want = `func (t *type) funcName(){
}
`
	var sb strings.Builder
	this := This("type").Pointer()

	newMethod(this, "funcName").write(&sb)
	assert.Equal(t, want, sb.String())
}

func TestMethodThisAlias(t *testing.T) {
	const want = `func (t alias.type) funcName(){
}
`
	var sb strings.Builder
	this := QualThis("alias", "type")

	newMethod(this, "funcName").write(&sb)
	assert.Equal(t, want, sb.String())
}

func TestMethodThisAliasPointer(t *testing.T) {
	const want = `func (t *alias.type) funcName(){
}
`
	var sb strings.Builder
	this := QualThis("alias", "type").Pointer()

	newMethod(this, "funcName").write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMethodWithParams(t *testing.T) {
	const want = `func (t type) funcName(name typeName,name2 *alias.typeName){
}
`
	this := This("type")
	param1, param2 := Param("name", "typeName"),
		QualParam("name2", "alias", "typeName").Pointer()

	var sb strings.Builder
	newMethod(this, "funcName").
		Params(param1, param2).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMethodWithReturnType(t *testing.T) {
	const want = `func (t type) funcName()type{
}
`
	this, retType := This("type"), Type("type")

	var sb strings.Builder
	newMethod(this, "funcName").
		ReturnTypes(retType).
		write(&sb)

	assert.Equal(t, want, sb.String())
}

func TestMethodWithReturnTypes(t *testing.T) {
	const want = `func (t type) funcName()(type,*alias.type){
}
`
	this := This("type")
	retType1, retType2 := Type("type"),
		QualType("alias", "type").Pointer()

	var sb strings.Builder
	newMethod(this, "funcName").
		ReturnTypes(retType1, retType2).
		write(&sb)

	assert.Equal(t, want, sb.String())
}
