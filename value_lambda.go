package codegen

import "strings"

type lambdaValue struct {
	block *funcBlock
}

// Lambda creates a new lambda function value
func Lambda() *lambdaValue {
	block := newFunc("")
	block.newLine = false

	return &lambdaValue{
		block: block,
	}
}

// Params appends parameters to the lambda definition
func (l *lambdaValue) Params(params ...*ParamDecl) *lambdaValue {
	l.block.Params(params...)
	return l
}

// ReturnTypes appends return parameters to the lambda definition
func (l *lambdaValue) ReturnTypes(returnTypes ...*ReturnTypeDecl) *lambdaValue {
	l.block.ReturnTypes(returnTypes...)
	return l
}

// Block appends the block to the lambda definition
func (l *lambdaValue) Block(stmts ...Stmt) *lambdaValue {
	l.block.Block(stmts...)
	return l
}

// Call calls the lambda definition
func (l *lambdaValue) Call() *callValue {
	return newCallValue(l, "")
}

func (l *lambdaValue) writeValue(sb *strings.Builder) {
	l.block.write(sb)
}

func (l *lambdaValue) isPointer() bool {
	return false
}
