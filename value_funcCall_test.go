package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFuncCallEmpty(t *testing.T) {
	const want = `myFunc()`

	var sb strings.Builder
	got := FuncCall("myFunc").writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestQualFuncCallEmpty(t *testing.T) {
	const want = `alias.MyFunc()`

	var sb strings.Builder
	got := QualFuncCall("alias", "MyFunc").writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestFuncCallPointer(t *testing.T) {
	const want = `*myFunc()`

	var sb strings.Builder
	got := FuncCall("myFunc").Pointer().
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestQualFuncCallPointer(t *testing.T) {
	const want = `*alias.MyFunc()`

	var sb strings.Builder
	got := QualFuncCall("alias", "MyFunc").Pointer().
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestFuncCallOneArg(t *testing.T) {
	const want = `myFunc(someFunc())`

	var sb strings.Builder
	got := FuncCall("myFunc").Args(FuncCall("someFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestQualFuncCallEmptyOneArg(t *testing.T) {
	const want = `alias.MyFunc(alias2.someFunc())`

	var sb strings.Builder
	got := QualFuncCall("alias", "MyFunc").Args(QualFuncCall("alias2", "someFunc")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestFuncCallArgs(t *testing.T) {
	const want = `myFunc(a,b)`

	var sb strings.Builder
	got := FuncCall("myFunc").Args(Identifier("a"), Identifier("b")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
}

func TestQualFuncCallArgs(t *testing.T) {
	const want = `alias.MyFunc(alias.a,b)`

	var sb strings.Builder
	got := QualFuncCall("alias", "MyFunc").Args(QualIdentifier("alias", "a"), Identifier("b")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
	assert.True(t, got)
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

func TestFuncCallCast(t *testing.T) {
	const want = `myFunc().(string)`

	var sb strings.Builder
	FuncCall("myFunc").Cast("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallPointerCast(t *testing.T) {
	const want = `(*myFunc()).(string)`

	var sb strings.Builder
	FuncCall("myFunc").Pointer().Cast("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallCastPointer(t *testing.T) {
	const want = `myFunc().(*string)`

	var sb strings.Builder
	FuncCall("myFunc").CastPointer("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallPointerCastPointer(t *testing.T) {
	const want = `(*myFunc()).(*string)`

	var sb strings.Builder
	FuncCall("myFunc").Pointer().CastPointer("string").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallCastQual(t *testing.T) {
	const want = `myFunc().(alias.MyType)`

	var sb strings.Builder
	FuncCall("myFunc").CastQual("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallPointerCastQual(t *testing.T) {
	const want = `(*myFunc()).(alias.MyType)`

	var sb strings.Builder
	FuncCall("myFunc").Pointer().CastQual("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallCastQualPointer(t *testing.T) {
	const want = `myFunc().(*alias.MyType)`

	var sb strings.Builder
	FuncCall("myFunc").CastQualPointer("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFuncCallPointerCastQualPointer(t *testing.T) {
	const want = `(*myFunc()).(*alias.MyType)`

	var sb strings.Builder
	FuncCall("myFunc").Pointer().CastQualPointer("alias", "MyType").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualFuncCallCast(t *testing.T) {
	const want = `alias.MyFunc(a,b).(string)`

	var sb strings.Builder
	QualFuncCall("alias", "MyFunc").Args(Identifier("a"), Identifier("b")).
		Cast("string").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
