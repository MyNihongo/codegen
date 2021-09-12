package codegen

import "strings"

type funcBlock struct {
	name     string
	params   []*ParamDecl
	retTypes []*ReturnTypeDecl
	stmts    []Stmt
}

// Func creates a new function code block
func (f *File) Func(name string) *funcBlock {
	fnc := newFunc(name)
	f.append(fnc)

	return fnc
}

// Params appends function parameters
func (f *funcBlock) Params(params ...*ParamDecl) *funcBlock {
	f.params = params
	return f
}

// ReturnTypes appends function return parameters
func (f *funcBlock) ReturnTypes(returnTypes ...*ReturnTypeDecl) *funcBlock {
	f.retTypes = returnTypes
	return f
}

// Block appends the function block
func (f *funcBlock) Block(stmts ...Stmt) {
	f.stmts = stmts
}

func newFunc(name string) *funcBlock {
	return &funcBlock{
		name:     name,
		params:   make([]*ParamDecl, 0),
		retTypes: make([]*ReturnTypeDecl, 0),
		stmts:    make([]Stmt, 0),
	}
}

func (f *funcBlock) write(sb *strings.Builder) {
	writeF(sb, "func %s", f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.retTypes)
	writeStmts(sb, f.stmts, true)
}
