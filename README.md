Learn Go Struct Validator
==========================

```go
package main

import (
	"fmt"

	v "github.com/nipeharefa/learn-go-validator"
)

type ExampleStruct struct {
	Name string `vd:"not_blank"`
}

func main() {

	e := &ExampleStruct{}

	errValidation := v.Validate(e)

	fmt.Println(errValidation)
}
```