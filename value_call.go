package codegen

import "strings"

type callValue struct {
	val  Value
	name string
	args []Value
}

// Args creates a new argument value for a function call
func (c *callValue) Args(args ...Value) *callValue {
	c.args = args
	return c
}

// Field appends a new field getter after the function call
func (c *callValue) Field(fieldName string) *fieldValue {
	return newField(c, fieldName)
}

// Call appends a new function call after the function call
func (c *callValue) Call(funcName string) *callValue {
	return newCallValue(c, funcName)
}

func newCallValue(val Value, funcName string) *callValue {
	return &callValue{
		val:  val,
		name: funcName,
		args: make([]Value, 0),
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
