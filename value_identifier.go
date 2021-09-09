package codegen

import (
	"fmt"
	"strings"
)

type identifierValue struct {
	declaration *nameValue
}

// Identifier creates a new identifier (variable, value, etc.)
func Identifier(name string) *identifierValue {
	return &identifierValue{
		declaration: qualName("", name),
	}
}

// Identifier creates a new identifier with a package alias (variable, value, etc.)
func QualIdentifier(alias, name string) *identifierValue {
	return &identifierValue{
		declaration: qualName(alias, name),
	}
}

// Pointer turns the identifier into a pointer type
func (i *identifierValue) Pointer() *identifierValue {
	i.declaration.pointer()
	return i
}

// Err creates a new identifier named `err`
func Err() *identifierValue {
	return Identifier("err")
}

// Nil creates a new identifier named `nil`
func Nil() *identifierValue {
	return Identifier("nil")
}

// String creates a new stiring value identifier
func String(strValue string) *identifierValue {
	return Identifier(fmt.Sprintf("\"%s\"", strValue))
}

// Int creates a new integer (int32) value identifier
func Int(intVal int) *identifierValue {
	return Identifier(fmt.Sprintf("%d", intVal))
}

// Field appends a new field getter after the identifier
func (i *identifierValue) Field(fieldName string) *fieldValue {
	return newField(i, fieldName)
}

// Call appends a new function call after the identifier
func (i *identifierValue) Call(funcName string) *callValue {
	return newCallValue(i, funcName)
}

// Assign assigns a value to the identifier
func (i *identifierValue) Assign(val value) *assignStmt {
	return newAssignment(i, val)
}

// Equals compares a value of the identifier for equality
func (i *identifierValue) Equals(value value) *equalsValue {
	return newEquals(i, value, true)
}

// Equals compares a value of the identifier for not being equal
func (i *identifierValue) NotEquals(value value) *equalsValue {
	return newEquals(i, value, false)
}

// IsNil checks whether or not the identifier is nil
func (i *identifierValue) IsNil() *equalsValue {
	return newEquals(i, Nil(), true)
}

// IsNil checks whether or not the identifier is not nil
func (i *identifierValue) IsNotNil() *equalsValue {
	return newEquals(i, Nil(), false)
}

// IsNotEmpty checks whether or not the identifier's length is empty
func (i *identifierValue) IsNotEmpty() *equalsValue {
	return Len(i).NotEquals(Int(0))
}

func (i *identifierValue) writeValue(sb *strings.Builder) {
	i.declaration.writeValue(sb)
}

func (i *identifierValue) isPointer() bool {
	return i.declaration.isPointer
}
