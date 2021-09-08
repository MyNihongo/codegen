package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdentifier(t *testing.T) {
	const want = `obj`

	var sb strings.Builder
	Identifier("obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierPointer(t *testing.T) {
	const want = `*obj`

	var sb strings.Builder
	Identifier("obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifier(t *testing.T) {
	const want = `alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierPointer(t *testing.T) {
	const want = `*alias.obj`

	var sb strings.Builder
	QualIdentifier("alias", "obj").Pointer().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierErr(t *testing.T) {
	const want = `err`

	var sb strings.Builder
	Err().writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNil(t *testing.T) {
	const want = `nil`

	var sb strings.Builder
	Nil().writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierString(t *testing.T) {
	const want = `"my string value"`

	var sb strings.Builder
	String("my string value").writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierInt(t *testing.T) {
	const want = `123`

	var sb strings.Builder
	Int(123).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierField(t *testing.T) {
	const want = `val.field`

	var sb strings.Builder
	Identifier("val").Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldPointer(t *testing.T) {
	const want = `(*val).field`

	var sb strings.Builder
	Identifier("val").Pointer().
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierField(t *testing.T) {
	const want = `alias.val.field`

	var sb strings.Builder
	QualIdentifier("alias", "val").Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierFieldPointer(t *testing.T) {
	const want = `(*alias.val).field`

	var sb strings.Builder
	QualIdentifier("alias", "val").Pointer().
		Field("field").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNestedField(t *testing.T) {
	const want = `obj.field1.field2`

	var sb strings.Builder
	Identifier("obj").
		Field("field1").Field("field2").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierFieldAssign(t *testing.T) {
	const want = `obj.field1=myFunc(a,b)`

	var sb strings.Builder
	Identifier("obj").
		Field("field1").Assign(FuncCall("myFunc").Args(Identifier("a"), Identifier("b"))).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierCallEmpty(t *testing.T) {
	const want = `obj.field.myFunc()`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierCallArgSingle(t *testing.T) {
	const want = `obj.field.myFunc(a)`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").Args(Identifier("a")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierCallArgs(t *testing.T) {
	const want = `obj.field.myFunc(a,anotherFunc(b))`

	var sb strings.Builder
	Identifier("obj").
		Field("field").Call("myFunc").Args(Identifier("a"), FuncCall("anotherFunc").Args(Identifier("b"))).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierAssign(t *testing.T) {
	const want = `a=alias.GetMyValue()`

	var sb strings.Builder
	Identifier("a").Assign(QualFuncCall("alias", "GetMyValue")).
		writeStmt(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierEquals(t *testing.T) {
	const want = `a==myFunc()`

	var sb strings.Builder
	Identifier("a").Equals(FuncCall("myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestQualIdentifierEquals(t *testing.T) {
	const want = `alias.Var==b`

	var sb strings.Builder
	QualIdentifier("alias", "Var").Equals(Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierNotEquals(t *testing.T) {
	const want = `a!=myFunc()`

	var sb strings.Builder
	Identifier("a").NotEquals(FuncCall("myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierIsNil(t *testing.T) {
	const want = `a==nil`

	var sb strings.Builder
	Identifier("a").IsNil().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierIsNotNil(t *testing.T) {
	const want = `a!=nil`

	var sb strings.Builder
	Identifier("a").IsNotNil().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestIdentifierIsNotEmpty(t *testing.T) {
	const want = `len(a)!=0`

	var sb strings.Builder
	Identifier("a").IsNotEmpty().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
