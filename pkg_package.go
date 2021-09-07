package codegen

import (
	"strings"
)

type packageBlock struct {
	pkgName string
}

func (p *packageBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "package %s", p.pkgName)
}

// pkg creates a package declaration statement
func pkg(pkgName string) *block {
	ptr := new(block)
	*ptr = &packageBlock{pkgName: pkgName}
	return ptr
}
