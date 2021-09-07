package codegen

import "strings"

type thisVal struct {
	*paramVal
}

func NewThis(typeName string) *thisVal {
	name := createThisName(typeName)
	return &thisVal{
		paramVal: NewParam(name, typeName),
	}
}

func NewQualThis(alias, typeName string) *thisVal {
	name := createThisName(typeName)

	return &thisVal{
		paramVal: NewQualParam(name, alias, typeName),
	}
}

func (t *thisVal) Pointer() *thisVal {
	t.paramVal.Pointer()
	return t
}

func createThisName(typeName string) string {
	if len(typeName) == 0 {
		panic("typeName must not be empty")
	}

	return strings.ToLower(string(typeName[0]))
}
