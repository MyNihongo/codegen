package codegen

import "strings"

// GenerateGetter creates a public getter method according to the field
func (f *File) GenerateGetter(this *thisDecl, fieldName string, returnType *ReturnTypeDecl) {
	methodName := strings.Title(fieldName)

	f.Method(this, methodName).ReturnTypes(returnType).Block(
		Return(Identifier(this.name).Field(fieldName)),
	)
}
