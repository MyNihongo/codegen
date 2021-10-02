package codegen

import "strings"

type castValue struct {
	val     Value
	name    *nameHelper
	pointer bool
}

// Pointer turns the cast value into a pointer
func (c *castValue) Pointer() *castValue {
	c.pointer = true
	return c
}

func newCastValue(val Value, alias, typeName string, isPointer bool) *castValue {
	return &castValue{
		val: val,
		name: &nameHelper{
			alias:      alias,
			identifier: typeName,
			isPointer:  isPointer,
		},
	}
}

func (c *castValue) writeValue(sb *strings.Builder) {
	writePointerValueAccess(sb, c.val)
	sb.WriteString(".(")

	c.name.writeValue(sb)

	sb.WriteByte(')')
}

func (c *castValue) isPointer() bool {
	return c.pointer
}
