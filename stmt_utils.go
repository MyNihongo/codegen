package codegen

import "strings"

func writeStmts(sb *strings.Builder, stmts []Stmt, appendNewLine bool) {
	writeByteNewLine(sb, '{')

	for _, stmt := range stmts {
		if stmt.writeStmt(sb) {
			newLine(sb)
		}
	}

	sb.WriteByte('}')

	if appendNewLine {
		newLine(sb)
	}
}
