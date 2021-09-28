package codegen

import (
	"sort"
	"strings"
)

type importLine struct {
	alias string
	path  string
}

type importsBlock struct {
	lines map[string]*importLine
}

// Imports creates a new imports block
func (f *File) Imports(imports ...*importLine) *importsBlock {
	for _, value := range imports {
		f.imports.addImport(value)
	}

	return f.imports
}

// AddImport adds a new import statement to the import block
func (i *importsBlock) AddImport(path string) *importLine {
	return i.addImport(Import(path))
}

// AddImportAlias adds a new import statement with its package alias to the import block
func (i *importsBlock) AddImportAlias(path, alias string) *importLine {
	return i.addImport(ImportAlias(path, alias))
}

// Import creates a new import statemet without an alias
func Import(path string) *importLine {
	return &importLine{path: path}
}

// ImportAlias creates a new import statement with an alias
func ImportAlias(path, alias string) *importLine {
	return &importLine{path: path, alias: alias}
}

func (i *importsBlock) addImport(line *importLine) *importLine {
	if existing, ok := i.lines[line.path]; ok {
		return existing
	} else {
		i.lines[line.path] = line
		return line
	}
}

func newImportsBlock() *importsBlock {
	return &importsBlock{
		lines: make(map[string]*importLine),
	}
}

func (i *importsBlock) write(sb *strings.Builder) {
	if count := len(i.lines); count == 0 {
		return
	} else {
		keys := make([]string, 0, count)
		for k := range i.lines {
			keys = append(keys, k)
		}

		if count == 1 {
			writeF(sb, "import ")
			i.lines[keys[0]].wr(sb)
		} else {
			sort.Strings(keys)
			writeNewLine(sb, "import (")

			for _, key := range keys {
				i.lines[key].wr(sb)
			}

			writeByteNewLine(sb, ')')
		}
	}
}

func (i *importLine) wr(sb *strings.Builder) {
	if len(i.alias) != 0 {
		writeF(sb, "%s ", i.alias)
	}

	writeNewLineF(sb, "\"%s\"", i.path)
}
