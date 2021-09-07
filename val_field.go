package codegen

import "strings"

type fieldValue struct {
	val  value
	name string
}

func (f *fieldValue) Field(fieldName string) *fieldValue {
	return newField(f, fieldName)
}

func newField(val value, name string) *fieldValue {
	return &fieldValue{
		val:  val,
		name: name,
	}
}

func (f *fieldValue) writeValue(sb *strings.Builder) {
	isPointer := f.val.isPointer()
	if isPointer {
		sb.WriteByte('(')
	}

	f.val.writeValue(sb)

	if isPointer {
		sb.WriteByte(')')
	}

	writeF(sb, ".%s", f.name)
}

func (f *fieldValue) isPointer() bool {
	return false
}
