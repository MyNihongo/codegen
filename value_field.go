package codegen

import "strings"

type fieldValue struct {
	val       Value
	name      string
	isAddress bool
}

// Field appends a new field getter after the field
func (f *fieldValue) Field(fieldName string) *fieldValue {
	return newField(f, fieldName)
}

// Call appends a new function call after the field
func (f *fieldValue) Call(name string) *callValue {
	return newCallValue(f, name)
}

// Cast casts the field to the specified type
func (f *fieldValue) Cast(typeName string) *castValue {
	return newCastValue(f, "", typeName, false)
}

// CastPointer casts the field to a pointer of the specified type
func (f *fieldValue) CastPointer(typeName string) *castValue {
	return newCastValue(f, "", typeName, true)
}

// CastQual casts the field to the specified type with an alias
func (f *fieldValue) CastQual(alias, typeName string) *castValue {
	return newCastValue(f, alias, typeName, false)
}

// CastQualPointer casts the field to a pointer of the specified type with an alias
func (f *fieldValue) CastQualPointer(alias, typeName string) *castValue {
	return newCastValue(f, alias, typeName, true)
}

// Assign assigns a value to the field
func (f *fieldValue) Assign(val Value) *assignStmt {
	return newAssignment(f, val)
}

// Equals compares a value of the identifier for equality
func (f *fieldValue) Equals(value Value) *comparisonValue {
	return newEquals(f, value, cmpType_Equals)
}

// Equals compares a value of the identifier for not being equal
func (f *fieldValue) NotEquals(value Value) *comparisonValue {
	return newEquals(f, value, cmpType_NotEquals)
}

// Address turns the field into an address type (pointer to the field)
func (f *fieldValue) Address() *fieldValue {
	f.isAddress = true
	return f
}

func newField(val Value, name string) *fieldValue {
	return &fieldValue{
		val:  val,
		name: name,
	}
}

func (f *fieldValue) writeValue(sb *strings.Builder) {
	write := func(b *strings.Builder) {
		writePointerValueAccess(b, f.val)
		writeF(b, ".%s", f.name)
	}

	writeAddressValueAccess(sb, write, f.isAddress)
}

func (f *fieldValue) isPointer() bool {
	return false
}
