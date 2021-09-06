package codegen

import "strings"

type importStmt struct {
	alias string
	path  string
}

// Import creates a new import statemet without an alias
func (f *file) Import(path string) *file {
	f.imports.append(&importStmt{path: path})
	return f
}

// ImportAlias creates a new import statement with an alias
func (f *file) ImportAlias(path, alias string) *file {
	f.imports.append(&importStmt{path: path, alias: alias})
	return f
}

func (i *importStmt) write(sb *strings.Builder) {
	if len(i.alias) != 0 {
		writeF(sb, "%s ", i.alias)
	}

	writeNewLineF(sb, "\"%s\"", i.path)
}
