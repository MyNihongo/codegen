package codegen

import "strings"

type qualNameFunc struct {
	alias string
	name  string
	args  []value
}

func FuncCall(name string) *qualNameFunc {
	return newFuncCall("", name)
}

func QualFuncCall(alias, name string) *qualNameFunc {
	return newFuncCall(alias, name)
}

func (q *qualNameFunc) Args(args ...value) *qualNameFunc {
	q.args = args
	return q
}

func newFuncCall(alias, name string) *qualNameFunc {
	return &qualNameFunc{
		name:  name,
		alias: alias,
		args:  make([]value, 0),
	}
}

func (q *qualNameFunc) writeValue(sb *strings.Builder) {
	writeAlias(sb, q.alias)
	writeF(sb, "%s(", q.name)
	writeValues(sb, q.args)
	sb.WriteByte(')')
}

func (q *qualNameFunc) isPointer() bool {
	return false
}

func (q *qualNameFunc) Field(fieldName string) value {
	return newField(q, fieldName)
}
