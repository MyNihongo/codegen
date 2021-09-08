package codegen

import "strings"

type importGroupBlock struct {
	imports []*importBlock
}

// importGroup creates an empty group with imports
func importGroup() *importGroupBlock {
	return &importGroupBlock{
		imports: make([]*importBlock, 0),
	}
}

// append adds a new import statement to the group
func (i *importGroupBlock) append(stmt *importBlock) {
	i.imports = append(i.imports, stmt)
}

func (i *importGroupBlock) write(sb *strings.Builder) {
	if len(i.imports) == 0 {
		return
	}

	writeNewLine(sb, "import (")

	for _, stmt := range i.imports {
		stmt.write(sb)
	}

	writeByteNewLine(sb, ')')
}
