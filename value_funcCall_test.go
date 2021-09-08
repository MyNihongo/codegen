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

func TestFuncCallPointer(t *testing.T) {
	const want = `*myFunc()`

	var sb strings.Builder
	FuncCall("myFunc").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallPointer(t *testing.T) {
	const want = `*alias.MyFunc()`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Pointer().
		writeValue(&sb)

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

func TestFuncCallField(t *testing.T) {
	const want = `myFunc(alias.a,b).field`

	var sb strings.Builder
	FuncCall("myFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallField(t *testing.T) {
	const want = `alias.MyFunc(alias.a,b).field`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallCallEmpty(t *testing.T) {
	const want = `myFunc(alias.a,b).myFunc()`

	var sb strings.Builder
	FuncCall("myFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallCallEmpty(t *testing.T) {
	const want = `alias.MyFunc(alias.a,b).myFunc()`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallCallArgs(t *testing.T) {
	const want = `myFunc(alias.a,b).myFunc(this)`

	var sb strings.Builder
	FuncCall("myFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Call("myFunc").Args(Identifier("this")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallCallArgs(t *testing.T) {
	const want = `alias.MyFunc(alias.a,b).myFunc(alias.AnotherFunc())`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		Call("myFunc").Args(QualFuncCall("alias", "AnotherFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
