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
func For(decl *declarationStmt, check Value, post Stmt) *forBlock {
	if check == nil {
		panic("check cannot be null")
	}

	return &forBlock{
		decl:  decl,
		check: check,
		post:  post,
		stmts: make([]Stmt, 0),
	}
}

// Block appends statements to the for-loop block
func (f *forBlock) Block(statements ...Stmt) *forBlock {
	f.stmts = statements
	return f
}

func (f *forBlock) write(sb *strings.Builder) {

}
