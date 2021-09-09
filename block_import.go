package codegen

import "strings"

type importLine struct {
	alias string
	path  string
}

type importsBlock struct {
	lines []*importLine
}

// Imports creates a new imports block
func (f *file) Imports(imports ...*importLine) *file {
	f.imports.lines = imports
	return f
}

// Import creates a new import statemet without an alias
func Import(path string) *importLine {
	return &importLine{path: path}
}

// ImportAlias creates a new import statement with an alias
func ImportAlias(path, alias string) *importLine {
	return &importLine{path: path, alias: alias}
}

func newImportsBlock() *importsBlock {
	return &importsBlock{
		lines: make([]*importLine, 0),
	}
}

func (i *importsBlock) write(sb *strings.Builder) {
	if count := len(i.lines); count == 0 {
		Return()
	} else if count == 1 {
		writeF(sb, "import ")
		i.lines[0].wr(sb)
	} else {
		writeNewLine(sb, "import (")

		for _, l := range i.lines {
			l.wr(sb)
		}

		writeByteNewLine(sb, ')')
	}
}

func (i *importLine) wr(sb *strings.Builder) {
	if len(i.alias) != 0 {
		writeF(sb, "%s ", i.alias)
	}

	writeNewLineF(sb, "\"%s\"", i.path)
}
