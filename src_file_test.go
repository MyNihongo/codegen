package codegen

import (
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
	got := NewFile(packageName, codeGen).generate()

	assert.Equal(t, want, got)
}

func TestFileImport(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
import (
"strings"
)
`
	got := NewFile(packageName, codeGen).
		Import("strings").
		generate()

	assert.Equal(t, want, got)
}

func TestFileImportAlias(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
import (
str "strings"
)
`
	got := NewFile(packageName, codeGen).
		ImportAlias("strings", "str").
		generate()

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
	got := NewFile(packageName, codeGen).
		Import("os").ImportAlias("strings", "str").
		generate()

	assert.Equal(t, want, got)
}

func TestFileCommentF(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
// this is a file comment
`
	got := NewFile(packageName, codeGen).
		CommentF("this is a file %s", "comment").
		generate()

	assert.Equal(t, want, got)
}

func TestFileFunc(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
func funcName(param1 *alias.typeName,param2 string){
}
`
	param1, param2 := NewQualParam("param1", "alias", "typeName").Pointer(), NewParam("param2", "string")
	f := NewFile(packageName, codeGen)
	f.NewFunc("funcName").Params(param1, param2)

	got := f.generate()

	assert.Equal(t, want, got)
}

func TestFileMethodGetter(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
func (m *MyType) funcName()int{
return m.field
}
`
	this, retType := NewThis("MyType").Pointer(), NewReturnType("int")

	f := NewFile(packageName, codeGen)
	f.NewMethod(this, "funcName").
		ReturnTypes(retType).
		Block(
			NewReturn(
				NewIdentifier("m").Field("field"),
			),
		)

	got := f.generate()

	assert.Equal(t, want, got)
}
