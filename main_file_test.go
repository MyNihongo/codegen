package codegen

import (
	"fmt"
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
	f := NewFile(packageName, codeGen).Imports(
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
	f := NewFile(packageName, codeGen).Imports(
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
	f := NewFile(packageName, codeGen).Imports(
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
	f := NewFile(packageName, codeGen).
		CommentF("this is a file %s", "comment")

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
