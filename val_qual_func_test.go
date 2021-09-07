package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncCallEmpty(t *testing.T) {
	const want = `myFunc()`

	var sb strings.Builder
	FuncCall("myFunc").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallEmpty(t *testing.T) {
	const want = `alias.MyFunc()`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallOneArg(t *testing.T) {
	const want = `myFunc(someFunc())`

	var sb strings.Builder
	FuncCall("myFunc").Args(FuncCall("someFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallEmptyOneArg(t *testing.T) {
	const want = `alias.MyFunc(alias2.someFunc())`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(QualFuncCall("alias2", "someFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallArgs(t *testing.T) {
	const want = `myFunc(a,b)`

	var sb strings.Builder
	FuncCall("myFunc").Args(Identifier("a"), Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallArgs(t *testing.T) {
	const want = `alias.MyFunc(alias.a,b)`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
