package codegen

import "strings"

type FuncDeclaration struct {
	name        string
	params      []*ParamDecl
	returnTypes []*ReturnTypeDecl
}

// FuncDecl creates a new function declaration
func FuncDecl(funcName string) *FuncDeclaration {
	return &FuncDeclaration{
		name:        funcName,
		params:      make([]*ParamDecl, 0),
		returnTypes: make([]*ReturnTypeDecl, 0),
	}
}

// Params adds parameters to the function declaration
func (f *FuncDeclaration) Params(params ...*ParamDecl) *FuncDeclaration {
	f.params = params
	return f
}

// ReturnTypes adds return types to the function declaration
func (f *FuncDeclaration) ReturnTypes(returnTypes ...*ReturnTypeDecl) *FuncDeclaration {
	f.returnTypes = returnTypes
	return f
}

func (f *FuncDeclaration) wr(sb *strings.Builder) {
	sb.WriteString(f.name)
	writeParams(sb, f.params)
	writeReturnTypes(sb, f.returnTypes)
}
