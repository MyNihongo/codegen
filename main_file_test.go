package codegen

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	packageName = "packageName"
	codeGen     = "my-nihongo-codegen"
)

func TestFileDeclaration(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
`
	f := NewFile(packageName, codeGen)
	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileImport(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
import "strings"
`
	f := NewFile(packageName, codeGen)
	f.Imports(
		Import("strings"),
	)

	got := fmt.Sprintf("%#v", f)
	assert.Equal(t, want, got)
}

func TestFileImportAlias(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
import str "strings"
`
	f := NewFile(packageName, codeGen)
	f.Imports(
		ImportAlias("strings", "str"),
	)

	got := fmt.Sprintf("%#v", f)
	assert.Equal(t, want, got)
}

func TestFileImports(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
import (
"os"
str "strings"
)
`
	f := NewFile(packageName, codeGen)
	f.Imports(
		Import("os"),
		ImportAlias("strings", "str"),
	)

	got := fmt.Sprintf("%#v", f)
	assert.Equal(t, want, got)
}

func TestFileCommentF(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
// this is a file comment
`
	f := NewFile(packageName, codeGen)
	f.CommentF("this is a file %s", "comment")

	got := fmt.Sprintf("%#v", f)
	assert.Equal(t, want, got)
}

func TestFileFunc(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
func funcName(param1 *alias.typeName,param2 string)(string,error){
if val,err:=os.Getenv("env_var");err!=nil{
return nil,errors.Errorf("this is a format %d",123)
} else {
if len(val)!=0{
return val.myField,nil
} else {
return nil,err
}
}
}
`
	f := NewFile(packageName, codeGen)

	param1, param2 := QualParam("param1", "alias", "typeName").Pointer(), Param("param2", "string")
	f.Func("funcName").Params(param1, param2).ReturnTypes(ReturnType("string"), ReturnTypeError()).Block(
		IfDecl(
			Declare("val", "err").Values(QualFuncCall("os", "Getenv").Args(String("env_var"))),
			Err().IsNotNil(),
		).Block(
			Return(Nil(), QualFuncCall("errors", "Errorf").Args(String("this is a format %d"), Int(123))),
		).Else(
			If(Identifier("val").IsNotEmpty()).Block(
				Return(Identifier("val").Field("myField"), Nil()),
			).Else(
				Return(Nil(), Err()),
			),
		),
	)

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileMethodGetter(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
func (m *MyType) funcName()int{
return m.field
}
`
	this, retType := This("MyType").Pointer(), ReturnType("int")

	f := NewFile(packageName, codeGen)
	f.Method(this, "funcName").
		ReturnTypes(retType).
		Block(
			Return(
				Identifier("m").Field("field"),
			),
		)

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareInterface(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myInterface interface {
myFunc()string
}
`
	f := NewFile(packageName, codeGen)
	f.Interface("myInterface").Funcs(
		FuncDecl("myFunc").ReturnTypes(ReturnType("string")),
	)

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareInterfaceDynamic(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myInterface interface {
MyFunc1()
MyFunc2()
MyFunc3()
}
// MyFunc1 does something
func (m *myInterface) MyFunc1(){
}
// MyFunc2 does something
func (m *myInterface) MyFunc2(){
}
// MyFunc3 does something
func (m *myInterface) MyFunc3(){
}
`

	f := NewFile(packageName, codeGen)
	decl := f.Interface("myInterface")

	for i := 0; i < 3; i++ {
		methodName := fmt.Sprintf("MyFunc%d", i+1)
		decl.AddFunc(FuncDecl(methodName))

		f.CommentF("%s does something", methodName)
		f.Method(
			This("myInterface").Pointer(),
			methodName,
		)
	}

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareStruct(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myStruct struct {
prop string
}
`
	f := NewFile(packageName, codeGen)
	f.Struct("myStruct").Props(
		Property("prop", "string"),
	)

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareStructDynamic(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myStruct struct {
myProp1 string
myProp2 string
myProp3 string
}
// MyProp1 does something
func (m *myStruct) MyProp1(){
}
// MyProp2 does something
func (m *myStruct) MyProp2(){
}
// MyProp3 does something
func (m *myStruct) MyProp3(){
}
`
	f := NewFile(packageName, codeGen)
	decl := f.Struct("myStruct")

	for i := 0; i < 3; i++ {
		propName := fmt.Sprintf("myProp%d", i+1)
		methodName := strings.Title(propName)

		decl.AddProp(Property(propName, "string"))

		f.CommentF("%s does something", methodName)
		f.Method(
			This("myStruct").Pointer(),
			methodName,
		)
	}

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareType(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myType int
`
	f := NewFile(packageName, codeGen)
	f.Type("myType", "int")

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileDeclareQualType(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
type myType alias.MyType
`
	f := NewFile(packageName, codeGen)
	f.QualType("myType", "alias", "MyType")

	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}

func TestFileVariableDeclaration(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
var var1,var2 string
var var3 alias.MyType
`
	f := NewFile(packageName, codeGen)
	f.DeclareVars(
		Var("var1", "string"),
		Var("var2", "string"),
		QualVar("var3", "alias", "MyType"),
	)

	got := fmt.Sprintf("%#v", f)
	assert.Equal(t, want, got)
}
