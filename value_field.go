package codegen

import "strings"

type fieldValue struct {
	val  value
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

// Assign assigns a value to the field
func (f *fieldValue) Assign(val value) *assignStmt {
	return newAssignment(f, val)
}

func newField(val value, name string) *fieldValue {
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
