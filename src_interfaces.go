package codegen

import "strings"

type stmt interface {
	write(sb *strings.Builder)
}
