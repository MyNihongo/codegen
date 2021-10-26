package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoFuncLenField(t *testing.T) {
	const want = `len(a.field)`

	var sb strings.Builder
	Len(Identifier("a").Field("field")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncLenFuncCall(t *testing.T) {
	const want = `len(alias.myFunc())`

	var sb strings.Builder
	Len(QualFuncCall("alias", "myFunc")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncEquals(t *testing.T) {
	const want = `len(a)==0`

	var sb strings.Builder
	Len(Identifier("a")).Equals(Int(0)).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncNotEquals(t *testing.T) {
	const want = `len(a)!=0`

	var sb strings.Builder
	Len(Identifier("a")).NotEquals(Int(0)).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncNoPointer(t *testing.T) {
	got := newGoFunc("len").isPointer()

	assert.False(t, got)
}

func TestGoFuncMake(t *testing.T) {
	const want = `make([]string,0)`

	var sb strings.Builder
	Make(Type("string")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeCount(t *testing.T) {
	const want = `make([]string,12)`

	var sb strings.Builder
	MakeWithCount(Type("string"), 12).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakePointer(t *testing.T) {
	const want = `make([]*string,0)`

	var sb strings.Builder
	Make(Type("string").Pointer()).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeCountPointer(t *testing.T) {
	const want = `make([]*string,12)`

	var sb strings.Builder
	MakeWithCount(Type("string").Pointer(), 12).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeQual(t *testing.T) {
	const want = `make([]alias.MyType,0)`

	var sb strings.Builder
	Make(QualType("alias", "MyType")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeCountQual(t *testing.T) {
	const want = `make([]alias.MyType,12)`

	var sb strings.Builder
	MakeWithCount(QualType("alias", "MyType"), 12).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeQualPointer(t *testing.T) {
	const want = `make([]*alias.MyType,0)`

	var sb strings.Builder
	Make(QualType("alias", "MyType").Pointer()).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestGoFuncMakeCountQualPointer(t *testing.T) {
	const want = `make([]*alias.MyType,12)`

	var sb strings.Builder
	MakeWithCount(QualType("alias", "MyType").Pointer(), 12).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFoFuncAppendEmpty(t *testing.T) {
	const want = `append(arr)`

	var sb strings.Builder
	Append(Identifier("arr")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFoFuncAppendSingle(t *testing.T) {
	const want = `append(arr,createElement())`

	var sb strings.Builder
	Append(Identifier("arr"), FuncCall("createElement")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestFoFuncAppendMultiple(t *testing.T) {
	const want = `append(arr,createElement(),obj.field)`

	var sb strings.Builder
	Append(Identifier("arr"), FuncCall("createElement"), Identifier("obj").Field("field")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
