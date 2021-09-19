package codegen

import "strings"

type funcCallValue struct {
	alias string
	name  string
	args  []Value
	isPtr bool
}

// FuncCall creates a new function call
func FuncCall(name string) *funcCallValue {
	return newFuncCall("", name)
}

// QualFuncCall creates a new function call with a package alias
func QualFuncCall(alias, name string) *funcCallValue {
	return newFuncCall(alias, name)
}

// Pointer turns the value to a pointer type
func (q *funcCallValue) Pointer() *funcCallValue {
	q.isPtr = true
	return q
}

// Args appeng argument for the function call
func (q *funcCallValue) Args(args ...Value) *funcCallValue {
	q.args = args
	return q
}

// Field appends a new field getter after the function call
func (q *funcCallValue) Field(fieldName string) *fieldValue {
	return newField(q, fieldName)
}

// Call appends a new function call after the function call
func (q *funcCallValue) Call(funcName string) *callValue {
	return newCallValue(q, funcName)
}

func newFuncCall(alias, name string) *funcCallValue {
	return &funcCallValue{
		name:  name,
		alias: alias,
		args:  make([]Value, 0),
	}
}

func (q *funcCallValue) writeStmt(sb *strings.Builder) bool {
	q.writeValue(sb)
	return true
}

func (q *funcCallValue) writeValue(sb *strings.Builder) {
	if q.isPtr {
		sb.WriteByte('*')
	}

	writeAlias(sb, q.alias)
	writeFuncCall(sb, q.name, q.args)
}

func (q *funcCallValue) isPointer() bool {
	return q.isPtr
}
