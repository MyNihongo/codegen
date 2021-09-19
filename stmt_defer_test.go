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

func TestDeferFuncArgs(t *testing.T) {
	const want = `defer myFunc(a)
`
	var sb strings.Builder
	Defer(FuncCall("myFunc").Args(Identifier("a"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestDeferFuncFunc(t *testing.T) {
	const want = `defer myFunc().MyFunc()
`
	var sb strings.Builder
	Defer(FuncCall("myFunc").Call("MyFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestDeferFieldCall(t *testing.T) {
	const want = `defer a.b.myFunc(b)
`
	var sb strings.Builder
	Defer(Identifier("a").Field("b").Call("myFunc").Args(Identifier("b"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestDeferLambdaCall(t *testing.T) {
	const want = `defer func (){
a.MyFunc()
}()
`
	lambda := Lambda().Block(
		Identifier("a").Call("MyFunc"),
	).Call()

	var sb strings.Builder
	Defer(lambda).writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}
