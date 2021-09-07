package codegen

import "strings"

type returnStmt struct {
	values []value
}

func Return(values ...value) stmt {
	return &returnStmt{
		values: values,
	}
}

func (r *returnStmt) writeStmt(sb *strings.Builder) {
	sb.WriteString("return")

	if len(r.values) != 0 {
		sb.WriteByte(' ')
		writeValues(sb, r.values)
	}
}
