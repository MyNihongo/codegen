package codegen

import "strings"

func (f *file) GenerateGetter(this *thisValue, fieldName string, returnType *returnType) {
	methodName := strings.Title(fieldName)

	f.Method(this, methodName).ReturnTypes(returnType).Block(
		Return(Identifier(this.name).Field(fieldName)),
	)
}
