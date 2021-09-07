package codegen

import "strings"

type qualNameFuncValue struct {
	alias string
	name  string
	args  []value
}

func FuncCall(name string) *qualNameFuncValue {
	return newFuncCall("", name)
}

func QualFuncCall(alias, name string) *qualNameFuncValue {
	return newFuncCall(alias, name)
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
	writeAlias(sb, q.alias)
	writeF(sb, "%s(", q.name)
	writeValues(sb, q.args)
	sb.WriteByte(')')
}

func (q *qualNameFuncValue) isPointer() bool {
	return false
}
