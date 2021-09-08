package codegen

import "strings"

type funcBlock struct {
	name     string
	params   []*paramValue
	retTypes []*returnType
	stmts    []stmt
}

// Func creates a new function code block
func (f *file) Func(name string) *funcBlock {
	fnc := newFunc(name)
	f.append(fnc)

	return fnc
}

// Params appends function parameters
func (f *funcBlock) Params(params ...*paramValue) *funcBlock {
	f.params = params
	return f
}

// ReturnTypes appends function return parameters
func (f *funcBlock) ReturnTypes(returnTypes ...*returnType) *funcBlock {
	f.retTypes = returnTypes
	return f
}

// Block appends the function block
func (f *funcBlock) Block(stmts ...stmt) {
	f.stmts = stmts
}

func newFunc(name string) *funcBlock {
	return &funcBlock{
		name:     name,
		params:   make([]*paramValue, 0),
		retTypes: make([]*returnType, 0),
		stmts:    make([]stmt, 0),
	}
}

func (f *funcBlock) write(sb *strings.Builder) {
	writeF(sb, "func %s", f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.retTypes)
	writeStmts(sb, f.stmts, true)
}
