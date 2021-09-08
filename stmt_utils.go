package codegen

import "strings"

func writeStmts(sb *strings.Builder, stmts []stmt, appendNewLine bool) {
	writeByteNewLine(sb, '{')

	for _, stmt := range stmts {
		stmt.writeStmt(sb)
		newLine(sb)
	}

	sb.WriteByte('}')

	if appendNewLine {
		newLine(sb)
	}
}
