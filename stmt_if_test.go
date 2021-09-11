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

	got := If(
		Identifier("val").IsNotNil(),
	).Block(
		Return(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestIfDeclarationIf(t *testing.T) {
	const want = `if val!=nil{
return val
} else if val,varr:=alias.Func();len(varr)!=0{
val=varr
}
`
	var sb strings.Builder

	got := If(
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
	assert.False(t, got)
}

func TestIfElse(t *testing.T) {
	const want = `if val!=nil{
return val
} else {
return nil
}
`
	var sb strings.Builder

	got := If(
		Identifier("val").IsNotNil(),
	).Block(
		Return(Identifier("val")),
	).Else(
		Return(Nil()),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestIfDeclaration(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
}
`
	var sb strings.Builder

	got := IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().IsNil(),
	).Block(
		Identifier("config").Field("myVar").Assign(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}

func TestIfDeclarationElseIfDeclaration(t *testing.T) {
	const want = `if val,err:=alias.myFunc();err!=nil{
return nil,err
} else if val,err:=anotherFunc(val);err!=nil{
abc=val
}
`
	var sb strings.Builder

	got := IfDeclr(
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
	assert.False(t, got)
}

func TestIfDeclarationElse(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
} else {
return nil,err
}
`
	var sb strings.Builder

	got := IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().IsNil(),
	).Block(
		Identifier("config").Field("myVar").Assign(Identifier("val")),
	).Else(
		Return(Nil(), Err()),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
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

	got := IfDeclr(
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
	assert.False(t, got)
}

func TestIfElseIf(t *testing.T) {
	const want = `if len(a)!=0{
b.field=a
} else if a==nil{
b.field2=a
} else {
b.field3=a
}
`

	var sb strings.Builder

	got := If(Identifier("a").IsNotEmpty()).Block(
		Identifier("b").Field("field").Assign(Identifier("a")),
	).ElseIf(Identifier("a").IsNil()).Block(
		Identifier("b").Field("field2").Assign(Identifier("a")),
	).Else(
		Identifier("b").Field("field3").Assign(Identifier("a")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.False(t, got)
}
