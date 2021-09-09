package codegen

import "strings"

type returnStmt struct {
	values []Value
}

// Return creates a new return statement
func Return(values ...Value) Stmt {
	return &returnStmt{
		values: values,
	}
}

func (r *returnStmt) writeStmt(sb *strings.Builder) bool {
	sb.WriteString("return")

	if len(r.values) != 0 {
		sb.WriteByte(' ')
		writeValues(sb, r.values)
	}

	return true
}
