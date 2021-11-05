package codegen

import "strings"

type assignStmt struct {
	valLeft  Value
	valRight Value
}

func newAssignment(valLeft Value, valRight Value) *assignStmt {
	return &assignStmt{
		valLeft:  valLeft,
		valRight: valRight,
	}
}

func (a *assignStmt) writeStmt(sb *strings.Builder) bool {
	a.valLeft.writeValue(sb)
	sb.WriteByte('=')
	a.valRight.writeValue(sb)

	return true
}
