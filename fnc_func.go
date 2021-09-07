package codegen

import "strings"

type funcStmt struct {
	name     string
	params   []*paramVal
	retTypes []*returnType
	stmts    []*stmt
}

// NewFunc creates a new function code block
func (f *file) NewFunc(name string) *funcStmt {
	ptr := new(block)
	fnc := newFunc(name)
	*ptr = fnc

	f.append(ptr)
	return fnc
}

// Params appends function parameters
func (f *funcStmt) Params(params ...*paramVal) *funcStmt {
	f.params = params
	return f
}

// ReturnTypes appends function return parameters
func (f *funcStmt) ReturnTypes(returnTypes ...*returnType) *funcStmt {
	f.retTypes = returnTypes
	return f
}

func newFunc(name string) *funcStmt {
	return &funcStmt{
		name:     name,
		params:   make([]*paramVal, 0),
		retTypes: make([]*returnType, 0),
		stmts:    make([]*stmt, 0),
	}
}

func (f *funcStmt) write(sb *strings.Builder) {
	writeF(sb, "func %s", f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.retTypes)
	writeStmts(sb, f.stmts, true)
}
