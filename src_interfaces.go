package codegen

import "strings"

type block interface {
	write(sb *strings.Builder)
}

type value interface {
	writeValue(sb *strings.Builder)
}
