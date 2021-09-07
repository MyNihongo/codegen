package codegen

import "strings"

type stmt interface {
	write(sb *strings.Builder)
}

type value interface {
	writeValue(sb *strings.Builder)
}
