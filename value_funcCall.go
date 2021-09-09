package codegen

import "strings"

type qualNameFuncValue struct {
	alias string
	name  string
	isPtr bool
	args  []Value
}

// FuncCall creates a new function call
func FuncCall(name string) *qualNameFuncValue {
	return newFuncCall("", name)
}

// QualFuncCall creates a new function call with a package alias
func QualFuncCall(alias, name string) *qualNameFuncValue {
	return newFuncCall(alias, name)
}

// Pointer turns the value to a pointer type
func (q *qualNameFuncValue) Pointer() *qualNameFuncValue {
	q.isPtr = true
	return q
}

// Args appeng argument for the function call
func (q *qualNameFuncValue) Args(args ...Value) *qualNameFuncValue {
	q.args = args
	return q
}

// Field appends a new field getter after the function call
func (q *qualNameFuncValue) Field(fieldName string) *fieldValue {
	return newField(q, fieldName)
}

// Call appends a new function call after the function call
func (q *qualNameFuncValue) Call(funcName string) *callValue {
	return newCallValue(q, funcName)
}

func newFuncCall(alias, name string) *qualNameFuncValue {
	return &qualNameFuncValue{
		name:  name,
		alias: alias,
		args:  make([]Value, 0),
	}
}

func (q *qualNameFuncValue) writeStmt(sb *strings.Builder) bool {
	q.writeValue(sb)
	return true
}

func (q *qualNameFuncValue) writeValue(sb *strings.Builder) {
	if q.isPtr {
		sb.WriteByte('*')
	}

	writeAlias(sb, q.alias)
	writeFuncCall(sb, q.name, q.args)
}

func (q *qualNameFuncValue) isPointer() bool {
	return q.isPtr
}
