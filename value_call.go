package codegen

import "strings"

type callValue struct {
	val  value
	name string
	args []value
}

func (c *callValue) Args(args ...value) *callValue {
	c.args = args
	return c
}

func (c *callValue) Field(fieldName string) *fieldValue {
	return newField(c, fieldName)
}

func (c *callValue) Call(funcName string) *callValue {
	return newCallValue(c, funcName)
}

func newCallValue(val value, funcName string) *callValue {
	return &callValue{
		val:  val,
		name: funcName,
		args: make([]value, 0),
	}
}

func (c *callValue) writeValue(sb *strings.Builder) {
	writePointerValueAccess(sb, c.val)
	sb.WriteByte('.')
	writeFuncCall(sb, c.name, c.args)
}

func (c *callValue) isPointer() bool {
	return false
}
