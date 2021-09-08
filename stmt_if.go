package codegen

import "strings"

type ifStmt struct {
	declaration *declarationStmt
	value       value
	stmts       []stmt
	prev        *ifStmt
	isFinal     bool
}

func If(val value) *ifStmt {
	return newIf(nil, nil, val)
}

func IfDeclr(declare *declarationStmt, val value) *ifStmt {
	return newIf(nil, declare, val)
}

func (i *ifStmt) ElseIfDeclr(declare *declarationStmt, val value) *ifStmt {
	return newIf(i, declare, val)
}

func (i *ifStmt) Else(stmts ...stmt) *ifStmt {
	stmt := newIf(i, nil, nil)
	stmt.stmts = stmts

	return stmt
}

func (i *ifStmt) Block(stmts ...stmt) *ifStmt {
	i.stmts = stmts
	return i
}

func newIf(prev *ifStmt, declare *declarationStmt, val value) *ifStmt {
	if prev != nil {
		prev.isFinal = false
	}

	return &ifStmt{
		declaration: declare,
		value:       val,
		prev:        prev,
		isFinal:     true,
	}
}

func (i *ifStmt) writeStmt(sb *strings.Builder) bool {
	if i.prev != nil {
		i.prev.writeStmt(sb)
		sb.WriteString(" else ")
	}

	if i.value != nil {
		sb.WriteString("if ")

		if i.declaration != nil {
			i.declaration.writeStmt(sb)
			sb.WriteByte(';')
		}

		i.value.writeValue(sb)
	}

	writeStmts(sb, i.stmts, i.isFinal)
	return !i.isFinal
}
