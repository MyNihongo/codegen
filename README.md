## Basic sample
This is a code generator for Go.
```go
package main

import (
	gen "github.com/MyNihongo/codegen"
	"fmt"
)

func main() {
	f := gen.NewFile("cool", "my-generator").Imports(
		gen.Import("fmt"),
		gen.ImportAlias("strings", "str"),
	)

	f.Func("main").Block(
		gen.Declare("val").Values(gen.QualFuncCall("str", "Title").Args(gen.String("hello, 世界!"))),
		gen.QualFuncCall("fmt", "Println").Args(gen.Identifier("val")),
	)

	if err := f.Save(`path to the file.go`); err != nil {
		fmt.Println(err)
	}
}

```
This code will produce the following output
```go
// Code generated by my-generator. DO NOT EDIT.
package cool

import (
	"fmt"
	str "strings"
)

func main() {
	val := str.Title("hello, 世界!")
	fmt.Println(val)
}

```

## Motivation
I have used [Jennifer](https://github.com/dave/jennifer) for code generation (don't get me wrong, the library is awesome :rocket:), but I thought that its API is a bit too low-level.

With this library I aim to provide high-level functions with will resemble the actual Go code and will not require to write low-level statements.

- `gofmt` is used for formatting
- auto-generation comment is added automatically

## Documentation
#### Identifiers
```go
gen.Identifier("a")
// a

gen.Identifier("a").Pointer()
// *a

gen.Identifier("a").Address()
// &(a)
```
#### Useful identifier methods
```go
gen.Identifier("a").Equals(gen.Identifier("b"))
// a == b

gen.Identifier("a").NotEquals(gen.Identifier("b"))
// a != b

gen.Err().IsNotNil()
// err != nil

gen.Err().IsNil()
// err == nil

gen.String("my string").IsNotEmpty()
// len("my string") != 0
```
#### Declare / assign variables
```go
gen.Declare("val").Values(gen.String("my string"))
// val := "my string"

gen.Declare("val", "err").Values(FuncCall("myFunc").Args(Identifier("a")))
// val, err := myFunc(a)

gen.DeclareVars(
	gen.Var("val", "string"),
	gen.QualVar("sb", "strings", "Builder"),
)
// var val string
// var sb strings.Builder

gen.Identifier("myVar").Assign(gen.Identifier("val"))
// myVar = val
```
#### Call functions
```go
gen.FuncCall("myFunc")
// myFunc()

gen.FuncCall("myFunc").Args(
	gen.Identifier("a"),
	gen.FuncCall("anotherFunc"),
)
// myFunc(a, anotherFunc())

gen.QualFuncCall("fmt", "Println").Args(
	gen.String("string value")
)
// fmt.Println("string value")
```
#### Call go functions
```go
gen.Len(Identifier("str"))
// len(str)
```
#### Access fields, call methods
```go
gen.Identifier("a").
	Field("field").
	Call("myFunc").Args(String("str")).
	Field("field2")
// a.field.myFunc("str").field2

gen.FuncCall("myFunc").
	Call("anotherFunc").Args(Identifier("a")).
	Field("field")
// myFunc().anotherFunc(a).field
```
#### Pointers
```go
gen.Identifier("a").Pointer().
	Field("field")
// (*a).field

gen.FuncCall("myFunc").Pointer().
	Field("field")
// (*myFunc()).field
```
#### Functions and methods
For methods the first argument is formatted according to Go conventions (first letter of the type in lowercase)
```go
f := gen.NewFile("cool", "my-generator")
f.Func("myFunc")
// func myFunc() {}

f.Func("myFunc").
	Params(
		gen.Param("val", "string"),
		gen.QualParam("sb", "strings", "Builder").Pointer(),
	)
// func myFunc(val string, sb *strings.Builder) {
// }

f.Func("myFunc").ReturnTypes(
		gen.ReturnType("myType").Pointer(),
		gen.ReturnTypeError(),
	).Block(
		gen.Return(gen.Identifier("a"), gen.Nil()),
	)
// func myFunc() (*myType, error) {
//	return a, nil
// }

f.Func("myFunc").Block(
	gen.Return(),
)
// func myFunc() {
//	return
// }
```
Function API is available for methods as well
```go
f := gen.NewFile("cool", "my-generator")
f.Method(
	gen.This("MyTypeName"),
	"coolMethod",
).ReturnTypes(
	gen.ReturnType("string"),
).Block(
	gen.Return(gen.Identifier("m").Field("field")),
)
// func (m MyTypeName) coolMethod() string {
//	return m.field
// }
```
#### If statements
If, else-if and else statements can be chained
```go
gen.If(
	gen.Identifier("val").IsNotNil(),
).Block(
	gen.Return(Identifier("val")),
).Else(
	gen.Return(Nil()),
)
// if val != nil{
//	return val
// } else {
//	return nil
// }

IfDecl(
	gen.Declare("val", "err").Values(gen.QualFuncCall("strconv", "Atoi").Args(gen.QualFuncCall("os", "Getenv").Args(gen.String("ENV_VAR")))),
	gen.Err().IsNil(),
).Block(
	gen.Identifier("config").Field("myVar").Assign(gen.Identifier("val")),
)
// if val, err := strconv.Atoi(os.Getenv("ENV_VAR")); err == nil {
//	config.myVar = val
// }
```
### Utility methods
#### Generate a getter
```go
f := gen.NewFile("cool", "my-generator").
f.GenerateGetter(gen.This("TypeName").Pointer(), "myField", gen.ReturnType("int"))
// func (t *TypeName) MyField() int {
//	return t.myField
// }
```