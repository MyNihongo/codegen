package codegen

import "strings"

type ifStmt struct {
	declaration *declarationStmt
	value       value
}

func IfDeclr(declare *declarationStmt, val value) *ifStmt {
	return &ifStmt{
		declaration: declare,
		value:       val,
	}
}

func (i *ifStmt) writeStmt(sb *strings.Builder) {

}
