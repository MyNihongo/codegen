package codegen

import "strings"

type ifStmt struct {
	declaration *declarationStmt
	value       value
	stmts       []stmt
}

func IfDeclr(declare *declarationStmt, val value) *ifStmt {
	return &ifStmt{
		declaration: declare,
		value:       val,
	}
}

func (i *ifStmt) Block(stmts ...stmt) *ifStmt {
	i.stmts = stmts
	return i
}

func (i *ifStmt) writeStmt(sb *strings.Builder) {
	sb.WriteString("if ")

	if i.declaration != nil {
		i.declaration.writeStmt(sb)
		sb.WriteByte(';')
	}

	i.value.writeValue(sb)
	writeStmts(sb, i.stmts, false)
}
