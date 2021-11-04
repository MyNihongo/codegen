package codegen

import "strings"

type MapTypeDecl struct {
	keyType   *TypeDecl
	valueType *TypeDecl
}

func MapType(keyType *TypeDecl, valueType *TypeDecl) *MapTypeDecl {
	if keyType.name.isArray {
		panic("key type cannot be an array")
	}

	return &MapTypeDecl{
		keyType:   keyType,
		valueType: valueType,
	}
}

func (m *MapTypeDecl) String() string {
	var sb strings.Builder
	m.wr(&sb)

	return sb.String()
}

func (m *MapTypeDecl) wr(sb *strings.Builder) {
	sb.WriteString("map[")
	m.keyType.wr(sb)
	sb.WriteByte(']')

	m.valueType.wr(sb)
}
