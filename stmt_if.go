package codegen

import "strings"

type ifStmt struct {
	declaration *declarationStmt
	value       value
	stmts       []stmt
	prev        *ifStmt
	isFinal     bool
}

// Creates a new if block
func If(val value) *ifStmt {
	return newIf(nil, nil, val)
}

// Creates a new if block with variable declaration
func IfDeclr(declare *declarationStmt, val value) *ifStmt {
	return newIf(nil, declare, val)
}

// Appends a new else-if block to the existing if block
func (i *ifStmt) ElseIf(val value) *ifStmt {
	return newIf(i, nil, val)
}

// Appends a new else-if block with variable declaration to the existing if block
func (i *ifStmt) ElseIfDeclr(declare *declarationStmt, val value) *ifStmt {
	return newIf(i, declare, val)
}

// Appends the final else block to the existing if block
func (i *ifStmt) Else(stmts ...stmt) stmt {
	stmt := newIf(i, nil, nil)
	stmt.stmts = stmts

	return stmt
}

// Appends block statement to the existing if statement
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
