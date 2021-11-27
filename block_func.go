package codegen

import "strings"

type funcBlock struct {
	name     string
	params   []*ParamDecl
	retTypes []*TypeDecl
	stmts    []Stmt
	newLine  bool
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
func (f *funcBlock) ReturnTypes(returnTypes ...*TypeDecl) *funcBlock {
	f.retTypes = returnTypes
	return f
}

// Block appends the function block
func (f *funcBlock) Block(stmts ...Stmt) *funcBlock {
	f.stmts = stmts
	return f
}

// AddStatement adds a new statement to the function block
func (f *funcBlock) AddStatement(stmt Stmt) {
	f.stmts = append(f.stmts, stmt)
}

func newFunc(name string) *funcBlock {
	return &funcBlock{
		name:     name,
		params:   make([]*ParamDecl, 0),
		retTypes: make([]*TypeDecl, 0),
		stmts:    make([]Stmt, 0),
		newLine:  true,
	}
}

func (f *funcBlock) write(sb *strings.Builder) {
	writeF(sb, "func %s", f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.retTypes)
	writeStmtsBlock(sb, f.stmts, f.newLine)
}
