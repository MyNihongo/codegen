package codegen

import "strings"

type funcBlock struct {
	name     string
	params   []*paramVal
	retTypes []*returnType
	stmts    []*stmt
}

// NewFunc creates a new function code block
func (f *file) NewFunc(name string) *funcBlock {
	fnc := newFunc(name)
	ptr := new(block)
	*ptr = fnc

	f.append(ptr)
	return fnc
}

// Params appends function parameters
func (f *funcBlock) Params(params ...*paramVal) *funcBlock {
	f.params = params
	return f
}

// ReturnTypes appends function return parameters
func (f *funcBlock) ReturnTypes(returnTypes ...*returnType) *funcBlock {
	f.retTypes = returnTypes
	return f
}

// Block appends the function block
func (f *funcBlock) Block() {
}

func newFunc(name string) *funcBlock {
	return &funcBlock{
		name:     name,
		params:   make([]*paramVal, 0),
		retTypes: make([]*returnType, 0),
		stmts:    make([]*stmt, 0),
	}
}

func (f *funcBlock) write(sb *strings.Builder) {
	writeF(sb, "func %s", f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.retTypes)
	writeStmts(sb, f.stmts, true)
}
