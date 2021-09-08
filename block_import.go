package codegen

import "strings"

type importBlock struct {
	alias string
	path  string
}

// Import creates a new import statemet without an alias
func (f *file) Import(path string) *file {
	f.imports.append(&importBlock{path: path})
	return f
}

// ImportAlias creates a new import statement with an alias
func (f *file) ImportAlias(path, alias string) *file {
	f.imports.append(&importBlock{path: path, alias: alias})
	return f
}

func (i *importBlock) write(sb *strings.Builder) {
	if len(i.alias) != 0 {
		writeF(sb, "%s ", i.alias)
	}

	writeNewLineF(sb, "\"%s\"", i.path)
}
