package codegen

import "strings"

type incrementStmt struct {
	val Value
}

func newIncrement(val Value) *incrementStmt {
	return &incrementStmt{
		val: val,
	}
}

func (i *incrementStmt) writeStmt(sb *strings.Builder) bool {
	writePointerValueAccess(sb, i.val)
	sb.WriteString("++")

	return true
}
