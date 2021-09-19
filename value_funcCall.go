package codegen

import "strings"

type funcCallValue struct {
	alias string
	isPtr bool
	*callHelper
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
		alias:      alias,
		callHelper: newCallHelper(name, make([]Value, 0)),
	}
}

func (f *funcCallValue) getCall() Value {
	return f
}

func (f *funcCallValue) writeStmt(sb *strings.Builder) bool {
	f.writeValue(sb)
	return true
}

func (f *funcCallValue) writeValue(sb *strings.Builder) {
	if f.isPtr {
		sb.WriteByte('*')
	}

	writeAlias(sb, f.alias)
	f.callHelper.wr(sb)
}

func (f *funcCallValue) isPointer() bool {
	return f.isPtr
}
