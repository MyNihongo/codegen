package codegen

import "strings"

type methodBlock struct {
	this *thisValue
	*funcBlock
}

// Method creates a new method block
func (f *File) Method(this *thisValue, name string) *methodBlock {
	method := newMethod(this, name)
	f.append(method)

	return method
}

// Params appends method parameters
func (m *methodBlock) Params(params ...*ParamValue) *methodBlock {
	m.funcBlock.Params(params...)
	return m
}

// ReturnTypes appends function return parameters
func (m *methodBlock) ReturnTypes(returnTypes ...*ReturnTypeDecl) *methodBlock {
	m.funcBlock.ReturnTypes(returnTypes...)
	return m
}

func newMethod(this *thisValue, name string) *methodBlock {
	return &methodBlock{
		this:      this,
		funcBlock: newFunc(name),
	}
}

func (m *methodBlock) write(sb *strings.Builder) {
	sb.WriteString("func (")
	m.this.writeValue(sb)
	writeF(sb, ") %s", m.name)

	writeParams(sb, m.params)
	writeReturnTypes(sb, m.retTypes)
	writeStmts(sb, m.stmts, true)
}
