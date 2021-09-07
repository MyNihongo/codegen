package codegen

import "strings"

type returnStmt struct {
	values []value
}

func NewReturn(values ...value) stmt {
	return &returnStmt{
		values: values,
	}
}

func (r *returnStmt) writeStmt(sb *strings.Builder) {
	sb.WriteString("return")

	if len(r.values) != 0 {
		sb.WriteByte(' ')

		for i, v := range r.values {
			if i != 0 {
				sb.WriteByte(',')
			}

			v.writeValue(sb)
		}
	}
}
