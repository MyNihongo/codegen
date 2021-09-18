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
	i := newInterface(interfaceName)
	f.append(i)

	return i
}

// Funcs adds function declarations to the interface
func (i *interfaceBlock) Funcs(functions ...*FuncDeclaration) *interfaceBlock {
	i.funcs = functions
	return i
}

// AddFunc adds a new function declaration to the interface
func (i *interfaceBlock) AddFunc(function *FuncDeclaration) {
	i.funcs = append(i.funcs, function)
}

func newInterface(name string) *interfaceBlock {
	return &interfaceBlock{
		name:  name,
		funcs: make([]*FuncDeclaration, 0),
	}
}

func (i *interfaceBlock) write(sb *strings.Builder) {
	writeNewLineF(sb, "type %s interface {", i.name)

	for _, function := range i.funcs {
		function.wr(sb)
		newLine(sb)
	}

	writeByteNewLine(sb, '}')
}
