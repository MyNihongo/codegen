package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	const want = `if val!=nil{
return val
}
`
	var sb strings.Builder

	If(
		Identifier("val").IsNotNil(),
	).Block(
		Return(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfDeclarationIf(t *testing.T) {
	const want = `if val!=nil{
return val
} else if val,varr:=alias.Func();len(varr)!=0{
val=varr
}
`
	var sb strings.Builder

	If(
		Identifier("val").IsNotNil(),
	).Block(
		Return(Identifier("val")),
	).ElseIfDeclr(
		Declare("val", "varr").Values(QualFuncCall("alias", "Func")),
		Identifier("varr").IsNotEmpty(),
	).Block(
		Identifier("val").Assign(Identifier("varr")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfElse(t *testing.T) {
	const want = `if val!=nil{
return val
} else {
return nil
}
`
	var sb strings.Builder

	If(
		Identifier("val").IsNotNil(),
	).Block(
		Return(Identifier("val")),
	).Else(
		Return(Nil()),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfDeclaration(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
}
`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().IsNil(),
	).Block(
		Identifier("config").Field("myVar").Assign(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfDeclarationElseIfDeclaration(t *testing.T) {
	const want = `if val,err:=alias.myFunc();err!=nil{
return nil,err
} else if val,err:=anotherFunc(val);err!=nil{
abc=val
}
`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("alias", "myFunc")),
		Err().IsNotNil(),
	).Block(
		Return(Nil(), Err()),
	).ElseIfDeclr(
		Declare("val", "err").Values(FuncCall("anotherFunc").Args(Identifier("val"))),
		Err().IsNotNil(),
	).Block(
		Identifier("abc").Assign(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfDeclarationElse(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
} else {
return nil,err
}
`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().IsNil(),
	).Block(
		Identifier("config").Field("myVar").Assign(Identifier("val")),
	).Else(
		Return(Nil(), Err()),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIfDeclarationElseIfDeclarationElse(t *testing.T) {
	const want = `if val,err:=alias.myFunc();err!=nil{
return nil,err
} else if val,err:=anotherFunc(val);err!=nil{
abc=val
} else {
return val,nil
}
`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("alias", "myFunc")),
		Err().IsNotNil(),
	).Block(
		Return(Nil(), Err()),
	).ElseIfDeclr(
		Declare("val", "err").Values(FuncCall("anotherFunc").Args(Identifier("val"))),
		Err().IsNotNil(),
	).Block(
		Identifier("abc").Assign(Identifier("val")),
	).Else(
		Return(Identifier("val"), Nil()),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
