package codegen

import "strings"

type callValue struct {
	val Value
	*callHelper
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

// Cast casts the return value of the call to the specified type
func (c *callValue) Cast(typeName string) *castValue {
	return newCastValue(c, "", typeName, false)
}

// CastPointer casts the return value of the call to a pointer of the specified type
func (c *callValue) CastPointer(typeName string) *castValue {
	return newCastValue(c, "", typeName, true)
}

// CastQual casts the return value of the call to the specified type with an alias
func (c *callValue) CastQual(alias, typeName string) *castValue {
	return newCastValue(c, alias, typeName, false)
}

// CastQualPointer casts the return value of the call to a pointer of the specified type with an alias
func (c *callValue) CastQualPointer(alias, typeName string) *castValue {
	return newCastValue(c, alias, typeName, true)
}

func newCallValue(val Value, funcName string) *callValue {
	return &callValue{
		val:        val,
		callHelper: newCallHelper(funcName, make([]Value, 0)),
	}
}

func (c *callValue) getCall() Value {
	return c
}

func (c *callValue) writeStmt(sb *strings.Builder) bool {
	c.writeValue(sb)
	newLine(sb)
	return false
}

func (c *callValue) writeValue(sb *strings.Builder) {
	writePointerValueAccess(sb, c.val)

	if len(c.callHelper.name) != 0 {
		sb.WriteByte('.')
	}

	c.callHelper.wr(sb)
}

func (c *callValue) isPointer() bool {
	return false
}
