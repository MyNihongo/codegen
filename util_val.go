package codegen

import "strings"

func writeValues(sb *strings.Builder, vals []value) {
	for i, v := range vals {
		if i != 0 {
			sb.WriteByte(',')
		}

		v.writeValue(sb)
	}
}
