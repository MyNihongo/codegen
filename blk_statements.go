package codegen

import "strings"

func writeStmts(sb *strings.Builder, stmts []*stmt, appendNewLine bool) {
	writeByteNewLine(sb, '{')
	sb.WriteByte('}')

	if appendNewLine {
		newLine(sb)
	}
}
