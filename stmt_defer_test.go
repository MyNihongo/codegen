package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeferFuncCall(t *testing.T) {
	const want = `defer myFunc()
`
	var sb strings.Builder
	Defer(FuncCall("myFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestDeferQualFuncCall(t *testing.T) {
	const want = `defer alias.MyFunc()
`

	var sb strings.Builder
	Defer(QualFuncCall("alias", "MyFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
