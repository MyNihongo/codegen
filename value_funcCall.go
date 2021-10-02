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

// Pointer dereferences the return value of the function
func (f *funcCallValue) Pointer() *funcCallValue {
	f.isPtr = true
	return f
}

// Args appeng argument for the function call
func (f *funcCallValue) Args(args ...Value) *funcCallValue {
	f.args = args
	return f
}

// Field appends a new field getter after the function call
func (f *funcCallValue) Field(fieldName string) *fieldValue {
	return newField(f, fieldName)
}

// Call appends a new function call after the function call
func (f *funcCallValue) Call(funcName string) *callValue {
	return newCallValue(f, funcName)
}

// Cast casts the function return value to the specified type
func (f *funcCallValue) Cast(typeName string) *castValue {
	return newCastValue(f, "", typeName, false)
}

// CastPointer casts thefunction return value to a pointer of the specified type
func (f *funcCallValue) CastPointer(typeName string) *castValue {
	return newCastValue(f, "", typeName, true)
}

// CastQual casts the function return value to the specified type with an alias
func (f *funcCallValue) CastQual(alias, typeName string) *castValue {
	return newCastValue(f, alias, typeName, false)
}

// CastQualPointer casts thefunction return value to a pointer of the specified type with an alias
func (f *funcCallValue) CastQualPointer(alias, typeName string) *castValue {
	return newCastValue(f, alias, typeName, true)
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
