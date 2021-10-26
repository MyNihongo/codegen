package codegen

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	packageName = "packageName"
	codeGen     = "my-nihongo-codegen"
)

func formatFile(file *File) string {
	return formatString(file.GoString())
}

func formatSb(sb strings.Builder) string {
	return formatString(sb.String())
}

func formatString(strVal string) string {
	if bytes, err := format.Source([]byte(strVal)); err != nil {
		panic(err)
	} else {
		return string(bytes)
	}
}

func TestFileSave(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

import "strings"

func myFunc(param1 string, sb strings.Builder) {
	if len(sb) != 0 {
		return
	} else {
		sb.WriteString("str")
		return
	}
}
`

	path, err := os.Getwd()
	assert.Nil(t, err)

	path = filepath.Join(path, "_save_test.go")

	f := NewFile(packageName, codeGen)
	f.AddImport("strings")
	f.Func("myFunc").Params(
		Param("param1", "string"),
		QualParam("sb", "strings", "Builder"),
	).Block(
		If(Identifier("sb").IsNotEmpty()).Block(
			Return(),
		).Else(
			Identifier("sb").Call("WriteString").Args(String("str")),
			Return(),
		),
	)

	err = f.Save(path)
	assert.Nil(t, err)

	defer os.Remove(path)

	content, err := ioutil.ReadFile(path)
	assert.Nil(t, err)

	got := string(content)
	assert.Equal(t, want, got)
}

func TestFileAddImport(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

import "strings"
`
	f := NewFile(packageName, codeGen)
	f.AddImport("strings")
	got := formatFile(f)

	assert.Equal(t, want, got)
}

func TestFileAddImportAlias(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

import alias "strings"
`
	f := NewFile(packageName, codeGen)
	f.AddImportAlias("strings", "alias")
	got := formatFile(f)

	assert.Equal(t, want, got)
}

func TestFileDeclaration(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
`
	f := NewFile(packageName, codeGen)
	got := formatFile(f)

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

	got := formatFile(f)
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

	got := formatFile(f)
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

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileCommentF(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

// this is a file comment
`
	f := NewFile(packageName, codeGen)
	f.CommentF("this is a file %s", "comment")

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileFunc(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

func funcName(param1 *alias.typeName, param2 string) (string, error) {
	if val, err := os.Getenv("env_var"); err != nil {
		return nil, errors.Errorf("this is a format %d", 123)
	} else {
		if len(val) != 0 {
			return val.myField, nil
		} else {
			return nil, err
		}
	}
}
`
	f := NewFile(packageName, codeGen)

	param1, param2 := QualParam("param1", "alias", "typeName").Pointer(), Param("param2", "string")
	f.Func("funcName").Params(param1, param2).ReturnTypes(Type("string"), ReturnTypeError()).Block(
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

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileMethodGetter(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

func (m *MyType) funcName() int {
	return m.field
}
`
	this, retType := This("MyType").Pointer(), Type("int")

	f := NewFile(packageName, codeGen)
	f.Method(this, "funcName").
		ReturnTypes(retType).
		Block(
			Return(
				Identifier("m").Field("field"),
			),
		)

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileDeclareInterface(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

type myInterface interface {
	myFunc() string
}
`
	f := NewFile(packageName, codeGen)
	f.Interface("myInterface").Funcs(
		FuncDecl("myFunc").ReturnTypes(Type("string")),
	)

	got := formatFile(f)
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
func (m *myInterface) MyFunc1() {
}

// MyFunc2 does something
func (m *myInterface) MyFunc2() {
}

// MyFunc3 does something
func (m *myInterface) MyFunc3() {
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

	got := formatFile(f)
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

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileDeclareStructDynamic(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

type myStruct struct {
	myProp1     string
	myProp1_NEW alias.MyType
	myProp2     string
	myProp2_NEW alias.MyType
	myProp3     string
	myProp3_NEW alias.MyType
}

// MyProp1 does something
func (m *myStruct) MyProp1() {
}

// MyProp2 does something
func (m *myStruct) MyProp2() {
}

// MyProp3 does something
func (m *myStruct) MyProp3() {
}
`
	f := NewFile(packageName, codeGen)
	decl := f.Struct("myStruct")

	for i := 0; i < 3; i++ {
		propName := fmt.Sprintf("myProp%d", i+1)
		methodName := strings.Title(propName)

		decl.AddProp(propName, "string")
		decl.AddQualProp(fmt.Sprintf("%s_NEW", propName), "alias", "MyType")

		f.CommentF("%s does something", methodName)
		f.Method(
			This("myStruct").Pointer(),
			methodName,
		)
	}

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileDeclareType(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

type myType int
`
	f := NewFile(packageName, codeGen)
	f.Type("myType", "int")

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileDeclareQualType(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

type myType alias.MyType
`
	f := NewFile(packageName, codeGen)
	f.QualType("myType", "alias", "MyType")

	got := formatFile(f)
	assert.Equal(t, want, got)
}

func TestFileVariableDeclaration(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName

var var1, var2 string
var var3 alias.MyType
`
	f := NewFile(packageName, codeGen)
	f.DeclareVars(
		Var("var1", "string"),
		Var("var2", "string"),
		QualVar("var3", "alias", "MyType"),
	)

	got := formatFile(f)
	assert.Equal(t, want, got)
}
