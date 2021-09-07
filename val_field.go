package codegen

import "strings"

type fieldValue struct {
	val  value
	name string
}

func newField(val value, name string) *fieldValue {
	return &fieldValue{
		val:  val,
		name: name,
	}
}

func (f *fieldValue) writeValue(sb *strings.Builder) {
	f.val.writeValue(sb)
	writeF(sb, ".%s", f.name)
}

func (f *fieldValue) Field(fieldName string) value {
	return newField(f, fieldName)
}
