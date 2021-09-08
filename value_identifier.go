package codegen

import (
	"fmt"
	"strings"
)

type identifierValue struct {
	declaration *nameValue
}

func Identifier(name string) *identifierValue {
	return &identifierValue{
		declaration: qualName("", name),
	}
}

func QualIdentifier(alias, name string) *identifierValue {
	return &identifierValue{
		declaration: qualName(alias, name),
	}
}

func Err() *identifierValue {
	return Identifier("err")
}

func Nil() *identifierValue {
	return Identifier("nil")
}

func String(strValue string) *identifierValue {
	return Identifier(fmt.Sprintf("\"%s\"", strValue))
}

func Int(intVal int) *identifierValue {
	return Identifier(fmt.Sprintf("%d", intVal))
}

func (i *identifierValue) Pointer() *identifierValue {
	i.declaration.pointer()
	return i
}

func (i *identifierValue) Field(fieldName string) *fieldValue {
	return newField(i, fieldName)
}

func (i *identifierValue) Assign(val value) *assignStmt {
	return newAssignment(i, val)
}

func (i *identifierValue) Equals(value value) *equalsValue {
	return newEquals(i, value, true)
}

func (i *identifierValue) NotEquals(value value) *equalsValue {
	return newEquals(i, value, false)
}

func (i *identifierValue) IsNil() *equalsValue {
	return newEquals(i, Nil(), true)
}

func (i *identifierValue) IsNotNil() *equalsValue {
	return newEquals(i, Nil(), false)
}

func (i *identifierValue) IsNotEmpty() *equalsValue {
	return Len(i).NotEquals(Int(0))
}

func (i *identifierValue) writeValue(sb *strings.Builder) {
	i.declaration.writeValue(sb)
}

func (i *identifierValue) isPointer() bool {
	return i.declaration.isPointer
}
