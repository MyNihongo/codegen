package codegen

import (
	"strings"
)

type interfaceBlock struct {
	name  string
	funcs []*FuncDeclaration
}

// Interface creates a new interface declaration block
func (f *File) Interface(interfaceName string) *interfaceBlock {
	i := &interfaceBlock{
		name:  interfaceName,
		funcs: make([]*FuncDeclaration, 0),
	}

	f.append(i)
	return i
}

// Funcs adds function declarations to the interface
func (i *interfaceBlock) Funcs(funcs ...*FuncDeclaration) {
	i.funcs = funcs
}

func (i *interfaceBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "type %s interface {", i.name)

	for _, function := range i.funcs {
		function.wr(sb)
	}

	writeByteNewLine(sb, '}')
}
