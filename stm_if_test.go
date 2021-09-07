package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfDeclaration(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
}
`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().Nil(),
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
		Err().NotNil(),
	).Block(
		Return(Nil(), Err()),
	).ElseIfDeclr(
		Declare("val", "err").Values(FuncCall("anotherFunc").Args(Identifier("val"))),
		Err().NotNil(),
	).Block(
		Identifier("abc").Assign(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
