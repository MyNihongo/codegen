package codegen

import "strings"

type castValue struct {
	val  Value
	name *nameHelper
}

// SetIsPointer sets whether or not the casting type is a pointer or not
func (c *castValue) SetIsPointer(isPointer bool) *castValue {
	c.name.setIsPointer(isPointer)
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

	c.name.wr(sb)

	sb.WriteByte(')')
}

func (c *castValue) isPointer() bool {
	return false
}
