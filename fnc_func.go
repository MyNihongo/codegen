package codegen

type funcStmt struct {
	name     string
	params   []*paramStmt
	retTypes []*returnType
}

func (f *file) NewFunc(name string) *funcStmt {
	return newFunc(name)
}

func newFunc(name string) *funcStmt {
	return &funcStmt{
		name:     name,
		params:   make([]*paramStmt, 0),
		retTypes: make([]*returnType, 0),
	}
}
