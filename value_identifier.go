package codegen

import (
	"fmt"
	"strings"
)

type identifierValue struct {
	declaration *nameHelper
	isAddress   bool
}

// Identifier creates a new identifier (variable, value, etc.)
func Identifier(name string) *identifierValue {
	return &identifierValue{
		declaration: newNameHelper("", name),
	}
}

// Identifier creates a new identifier with a package alias (variable, value, etc.)
func QualIdentifier(alias, name string) *identifierValue {
	return &identifierValue{
		declaration: newNameHelper(alias, name),
	}
}

// Pointer dereferences the identifier
func (i *identifierValue) Pointer() *identifierValue {
	i.SetIsPointer(true)
	return i
}

// SetIsPointer sets whether or not an identifier is a pointer
func (i *identifierValue) SetIsPointer(isPointer bool) *identifierValue {
	i.declaration.setIsPointer(isPointer)
	return i
}

// Address turns the identifier into an address type (pointer to the identifier)
func (i *identifierValue) Address() *identifierValue {
	i.isAddress = true
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

// Cast casts the identifier to the specified type
func (i *identifierValue) Cast(typeName string) *castValue {
	return newCastValue(i, "", typeName, false)
}

// CastPointer casts the identifier to a pointer of the specified type
func (i *identifierValue) CastPointer(typeName string) *castValue {
	return newCastValue(i, "", typeName, true)
}

// CastQual casts the identifier to the specified type with an alias
func (i *identifierValue) CastQual(alias, typeName string) *castValue {
	return newCastValue(i, alias, typeName, false)
}

// CastPointer casts the identifier to a pointer of the specified type with an alias
func (i *identifierValue) CastQualPointer(alias, typeName string) *castValue {
	return newCastValue(i, alias, typeName, true)
}

// Assign assigns a value to the identifier
func (i *identifierValue) Assign(val Value) *assignStmt {
	return newAssignment(i, val)
}

// Equals compares a value of the identifier for equality
func (i *identifierValue) Equals(value Value) *comparisonValue {
	return newEquals(i, value, cmpType_Equals)
}

// NotEquals compares a value of the identifier for not being equal
func (i *identifierValue) NotEquals(value Value) *comparisonValue {
	return newEquals(i, value, cmpType_NotEquals)
}

// LessThan compares a value of the identifier for being less
func (i *identifierValue) LessThan(value Value) *comparisonValue {
	return newEquals(i, value, cmpType_LessThan)
}

// IsNil checks whether or not the identifier is nil
func (i *identifierValue) IsNil() *comparisonValue {
	return newEquals(i, Nil(), cmpType_Equals)
}

// IsNil checks whether or not the identifier is not nil
func (i *identifierValue) IsNotNil() *comparisonValue {
	return newEquals(i, Nil(), cmpType_NotEquals)
}

// IsNotEmpty checks whether or not the identifier's length is empty
func (i *identifierValue) IsNotEmpty() *comparisonValue {
	return Len(i).NotEquals(Int(0))
}

// Increment increments the identifier (++ syntax)
func (i *identifierValue) Increment() *incrementStmt {
	return newIncrement(i)
}

// AtIndex access a value at the specified index
func (i *identifierValue) AtIndex(value Value) *squireBracketsValue {
	return newSquireBrackets(i, value)
}

func (i *identifierValue) writeValue(sb *strings.Builder) {
	writeAddressValueAccess(sb, i.declaration.wr, i.isAddress)
}

func (i *identifierValue) isPointer() bool {
	return i.declaration.isPointer
}
