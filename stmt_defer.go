package codegen

import "strings"

type deferStmt struct {
	call Value
}

func Defer(call caller) Stmt {
	return &deferStmt{
		call: call.getCall(),
	}
}

func (d *deferStmt) writeStmt(sb *strings.Builder) bool {
	sb.WriteString("defer ")
	d.call.writeValue(sb)
	newLine(sb)

	return false
}
