package codegen

import "strings"

type identifierValue struct {
	declaration *qualNameVal
	isPointer   bool
}

func NewIdentifier(name string) *identifierValue {
	return &identifierValue{
		declaration: NewQualName("", name),
	}
}

func NewQualIdentifier(alias, name string) *identifierValue {
	return &identifierValue{
		declaration: NewQualName(alias, name),
	}
}

func (i *identifierValue) Pointer() *identifierValue {
	i.isPointer = true
	return i
}

func (i *identifierValue) writeValue(sb *strings.Builder) {
	i.declaration.writeValue(sb)
}

func (i *identifierValue) Field(fieldName string) value {
	return newField(i, fieldName)
}
