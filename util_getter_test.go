package codegen

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateGetter(t *testing.T) {
	const want = `// Code generated by my-nihongo-codegen. DO NOT EDIT.
package packageName
func (t *TypeName) MyField()int{
return t.myField
}
`
	f := NewFile(packageName, codeGen)
	f.GenerateGetter(This("TypeName").Pointer(), "myField", ReturnType("int"))
	got := fmt.Sprintf("%#v", f)

	assert.Equal(t, want, got)
}
