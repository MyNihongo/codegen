package codegen

import "strings"

type methodBlock struct {
	this *thisDecl
	*funcBlock
}

// Method creates a new method block
func (f *File) Method(this *thisDecl, name string) *methodBlock {
	method := newMethod(this, name)
	f.append(method)

	return method
}

// Params appends method parameters
func (m *methodBlock) Params(params ...*ParamDecl) *methodBlock {
	m.funcBlock.Params(params...)
	return m
}

// ReturnTypes appends function return parameters
func (m *methodBlock) ReturnTypes(returnTypes ...*TypeDecl) *methodBlock {
	m.funcBlock.ReturnTypes(returnTypes...)
	return m
}

func newMethod(this *thisDecl, name string) *methodBlock {
	return &methodBlock{
		this:      this,
		funcBlock: newFunc(name),
	}
}

func (m *methodBlock) write(sb *strings.Builder) {
	sb.WriteString("func (")
	m.this.wr(sb)
	writeF(sb, ") %s", m.name)

	writeParams(sb, m.params)
	writeReturnTypes(sb, m.retTypes)
	writeStmtsBlock(sb, m.stmts, true)
}
