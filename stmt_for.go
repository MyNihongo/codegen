package codegen

import (
	"strings"
)

type forBlock struct {
	decl  *declarationStmt
	check Value
	post  Stmt
	stmts []Stmt
}

// For creates a new for-loop block
func For(decl *declarationStmt, check Value, postStatement Stmt) *forBlock {
	if check == nil {
		panic("check cannot be null")
	}

	return &forBlock{
		decl:  decl,
		check: check,
		post:  postStatement,
		stmts: make([]Stmt, 0),
	}
}

// Block appends statements to the for-loop block
func (f *forBlock) Block(statements ...Stmt) *forBlock {
	f.stmts = statements
	return f
}

func (f *forBlock) writeStmt(sb *strings.Builder) bool {
	sb.WriteString("for ")
	if f.decl != nil {
		f.decl.writeStmt(sb)
	}

	sb.WriteByte(';')
	f.check.writeValue(sb)
	sb.WriteByte(';')

	if f.post != nil {
		f.post.writeStmt(sb)
	}

	sb.WriteByte(' ')
	writeStmtsBlock(sb, f.stmts, true)
	return false
}
