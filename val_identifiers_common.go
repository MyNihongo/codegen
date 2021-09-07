package codegen

import "fmt"

func Err() *identifierValue {
	return Identifier("err")
}

func Nil() *identifierValue {
	return Identifier("nil")
}

func String(strValue string) *identifierValue {
	return Identifier(fmt.Sprintf("\"%s\"", strValue))
}
