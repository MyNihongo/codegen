package codegen

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaEmpty(t *testing.T) {
	const want = `func (){
}`

	var sb strings.Builder
	Lambda().writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaNotPointer(t *testing.T) {
	got := Lambda().isPointer()
	assert.False(t, got)
}

func TestLambdaParamSingle(t *testing.T) {
	const want = `func (param string){
}`

	var sb strings.Builder
	Lambda().Params(Param("param", "string")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaParamMultiple(t *testing.T) {
	const want = `func (param string,param2 alias.MyType){
}`

	var sb strings.Builder
	Lambda().Params(
		Param("param", "string"),
		QualParam("param2", "alias", "MyType"),
	).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaReturnTypeSingle(t *testing.T) {
	const want = `func ()string{
return "str value"
}`

	var sb strings.Builder
	Lambda().ReturnTypes(
		Type("string"),
	).Block(
		Return(String("str value")),
	).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaReturnTypeMultiple(t *testing.T) {
	const want = `func ()(string,error){
return "str value",nil
}`

	var sb strings.Builder
	Lambda().ReturnTypes(
		Type("string"),
		ReturnTypeError(),
	).Block(
		Return(String("str value"), Nil()),
	).writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaCallArgsEmpty(t *testing.T) {
	const want = `func (){
}()`

	var sb strings.Builder
	Lambda().Call().
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaCallArgsSingle(t *testing.T) {
	const want = `func (){
}(a)`

	var sb strings.Builder
	Lambda().Call().Args(Identifier("a")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}

func TestLambdaCallArgsMultiple(t *testing.T) {
	const want = `func (){
}(a,b)`

	var sb strings.Builder
	Lambda().Call().Args(Identifier("a"), Identifier("b")).
		writeValue(&sb)

	assert.Equal(t, want, sb.String())
}
