package codegen

import "strings"

type identifierValue struct {
	declaration *qualNameVal
}

func Identifier(name string) *identifierValue {
	return &identifierValue{
		declaration: QualName("", name),
	}
}

func QualIdentifier(alias, name string) *identifierValue {
	return &identifierValue{
		declaration: QualName(alias, name),
	}
}

func (i *identifierValue) Pointer() *identifierValue {
	i.declaration.pointer()
	return i
}

func (i *identifierValue) Field(fieldName string) *fieldValue {
	return newField(i, fieldName)
}

func (i *identifierValue) writeValue(sb *strings.Builder) {
	i.declaration.writeValue(sb)
}

func (i *identifierValue) isPointer() bool {
	return i.declaration.isPointer
}
