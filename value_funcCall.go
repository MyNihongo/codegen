package codegen

import "strings"

type qualNameFuncValue struct {
	alias string
	name  string
	isPtr bool
	args  []value
}

func FuncCall(name string) *qualNameFuncValue {
	return newFuncCall("", name)
}

func QualFuncCall(alias, name string) *qualNameFuncValue {
	return newFuncCall(alias, name)
}

func (q *qualNameFuncValue) Pointer() *qualNameFuncValue {
	q.isPtr = true
	return q
}

func (q *qualNameFuncValue) Args(args ...value) *qualNameFuncValue {
	q.args = args
	return q
}

func (q *qualNameFuncValue) Field(fieldName string) *fieldValue {
	return newField(q, fieldName)
}

func newFuncCall(alias, name string) *qualNameFuncValue {
	return &qualNameFuncValue{
		name:  name,
		alias: alias,
		args:  make([]value, 0),
	}
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
