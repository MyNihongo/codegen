package codegen

import "strings"

type assignStmt struct {
	valRight Value
	valLeft  Value
}

func newAssignment(valRight Value, valLeft Value) *assignStmt {
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
