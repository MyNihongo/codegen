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

func newCallValue(val value, name string) *callValue {
	return &callValue{
		val:  val,
		name: name,
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
