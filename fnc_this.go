package codegen

import "strings"

type thisVal struct {
	*paramVal
}

func This(typeName string) *thisVal {
	name := createThisName(typeName)
	return &thisVal{
		paramVal: Param(name, typeName),
	}
}

func QualThis(alias, typeName string) *thisVal {
	name := createThisName(typeName)

	return &thisVal{
		paramVal: QualParam(name, alias, typeName),
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
