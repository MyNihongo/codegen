# Basic sample
This is a code generator for Go.
```go
package main

import (
	"fmt"
	gen "github.com/MyNihongo/codegen"
)

func main() {
	f := gen.NewFile("cool", "my-generator").
		Import("fmt").
		ImportAlias("strings", "str")

	f.Func("main").Block(
		gen.QualFuncCall("fmt", "Println").Args(String("Hello, 世界!"))
	)

	if err := f.Save("path to the file.go"); err != nil {
		fmt.Println(err)
	}
}
```

## Motivation
todo...

## Documentation
todo...