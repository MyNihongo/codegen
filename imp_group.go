package codegen

import "strings"

type importGroupStmt struct {
	imports []importStmt
}

// importGroup creates an empty group with imports
func importGroup() *importGroupStmt {
	return &importGroupStmt{
		imports: make([]importStmt, 0),
	}
}

// append adds a new import statement to the group
func (i *importGroupStmt) append(stmt *importStmt) {
	i.imports = append(i.imports, *stmt)
}

func (i *importGroupStmt) write(sb *strings.Builder) {
	if len(i.imports) == 0 {
		return
	}

	writeNewLine(sb, "import (")

	for _, stmt := range i.imports {
		stmt.write(sb)
	}

	writeByteNewLine(sb, ')')
}
