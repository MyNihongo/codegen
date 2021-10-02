package codegen

import "strings"

type fieldValue struct {
	val  Value
	name string
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

func newField(val Value, name string) *fieldValue {
	return &fieldValue{
		val:  val,
		name: name,
	}
}

func (f *fieldValue) writeValue(sb *strings.Builder) {
	writePointerValueAccess(sb, f.val)
	writeF(sb, ".%s", f.name)
}

func (f *fieldValue) isPointer() bool {
	return false
}
