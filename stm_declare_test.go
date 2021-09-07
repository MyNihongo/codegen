package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclareOneVariable(t *testing.T) {
	const want = `newVar:=alias.ConstVal`

	var sb strings.Builder
	Declare("newVar").Values(QualIdentifier("alias", "ConstVal")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestDeclareMultipleVariables(t *testing.T) {
	const want = `var1,var2:=alias.ConstVal,alias.Value2`

	var sb strings.Builder
	Declare("var1", "var2").Values(QualIdentifier("alias", "ConstVal"), QualIdentifier("alias", "Value2")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
