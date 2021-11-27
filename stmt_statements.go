package codegen

import "strings"

type statementsStmt struct {
	stmts []Stmt
}

// Stmts combines multiple statements into a single one
func Stmts(stmts []Stmt) Stmt {
	return &statementsStmt{
		stmts: stmts,
	}
}

func (s *statementsStmt) writeStmt(sb *strings.Builder) bool {
	if len(s.stmts) != 0 {
		writeStmts(sb, s.stmts)
	}

	return false
}
