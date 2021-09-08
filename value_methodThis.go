package codegen

import "strings"

type thisValue struct {
	*paramValue
}

func This(typeName string) *thisValue {
	name := createThisName(typeName)
	return &thisValue{
		paramValue: Param(name, typeName),
	}
}

func QualThis(alias, typeName string) *thisValue {
	name := createThisName(typeName)

	return &thisValue{
		paramValue: QualParam(name, alias, typeName),
	}
}

func (t *thisValue) Pointer() *thisValue {
	t.paramValue.Pointer()
	return t
}

func createThisName(typeName string) string {
	if len(typeName) == 0 {
		panic("typeName must not be empty")
	}

	return strings.ToLower(string(typeName[0]))
}
