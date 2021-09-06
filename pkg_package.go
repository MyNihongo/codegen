package codegen

import (
	"strings"
)

type packageStmt struct {
	pkgName string
}

func (p *packageStmt) write(sb *strings.Builder) {
	writeNewLineF(sb, "package %s", p.pkgName)
}

// pkg creates a package declaration statement
func pkg(pkgName string) *stmt {
	ptr := new(stmt)
	*ptr = &packageStmt{pkgName: pkgName}
	return ptr
}
