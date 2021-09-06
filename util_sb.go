package codegen

import (
	"fmt"
	"strings"
)

func writeNewLine(sb *strings.Builder, str string) {
	sb.WriteString(str)
	newLine(sb)
}

func writeF(sb *strings.Builder, format string, a ...interface{}) {
	sb.WriteString(fmt.Sprintf(format, a...))
}

func writeNewLineF(sb *strings.Builder, format string, a ...interface{}) {
	sb.WriteString(fmt.Sprintf(format, a...))
	newLine(sb)
}

func writeByteNewLine(sb *strings.Builder, c byte) {
	sb.WriteByte(c)
	newLine(sb)
}

func newLine(sb *strings.Builder) {
	sb.WriteByte('\n')
}
