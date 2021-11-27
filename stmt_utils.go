package codegen

import "strings"

func writeStmtsBlock(sb *strings.Builder, stmts []Stmt, appendNewLine bool) {
	writeByteNewLine(sb, '{')
	writeStmts(sb, stmts)
	sb.WriteByte('}')

	if appendNewLine {
		newLine(sb)
	}
}

func writeStmts(sb *strings.Builder, stmts []Stmt) {
	for _, stmt := range stmts {
		if stmt.writeStmt(sb) {
			newLine(sb)
		}
	}
}
