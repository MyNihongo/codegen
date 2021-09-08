package codegen

import "strings"

type fieldValue struct {
	val  value
	name string
}

func (f *fieldValue) Field(fieldName string) *fieldValue {
	return newField(f, fieldName)
}

func (f *fieldValue) Call(name string) *callValue {
	return newCallValue(f, name)
}

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
