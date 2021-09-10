package codegen

import "strings"

type thisValue struct {
	*ParamValue
}

// This creates a new method this-parameter for a
func This(typeName string) *thisValue {
	name := createThisName(typeName)
	return &thisValue{
		ParamValue: Param(name, typeName),
	}
}

// QualThis creates a new this-parameter for a method with a package alias
func QualThis(alias, typeName string) *thisValue {
	name := createThisName(typeName)

	return &thisValue{
		ParamValue: QualParam(name, alias, typeName),
	}
}

// Pointer turns the this-parameter to a pointer type
func (t *thisValue) Pointer() *thisValue {
	t.ParamValue.Pointer()
	return t
}

func createThisName(typeName string) string {
	if len(typeName) == 0 {
		panic("typeName must not be empty")
	}

	return strings.ToLower(string(typeName[0]))
}
