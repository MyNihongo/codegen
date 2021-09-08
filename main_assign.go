package codegen

import "strings"

type assignStmt struct {
	valRight value
	valLeft  value
}

func newAssignment(valRight value, valLeft value) *assignStmt {
	return &assignStmt{
		valRight: valRight,
		valLeft:  valLeft,
	}
}

func (a *assignStmt) writeStmt(sb *strings.Builder) bool {
	a.valRight.writeValue(sb)
	sb.WriteByte('=')
	a.valLeft.writeValue(sb)

	return true
}
