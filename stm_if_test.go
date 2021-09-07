package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfDeclaration(t *testing.T) {
	const want = `if val,err:=strconv.Atoi(os.Getenv("ENV_VAR"));err==nil{
config.myVar=val
}`
	var sb strings.Builder

	IfDeclr(
		Declare("val", "err").Values(QualFuncCall("strconv", "Atoi").Args(QualFuncCall("os", "Getenv").Args(String("ENV_VAR")))),
		Err().Nil(),
	).Block(
		Identifier("config").Field("myVar").Assign(Identifier("val")),
	).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
