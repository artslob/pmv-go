package main

import (
	"fmt"
	"github.com/artslob/pmv-go/codegen"
)

func main() {
	var input = `
def func()
	a = (5 + 10) - (4 + 2);
end
`
	listener := codegen.GenerateCode(input)
	fmt.Println(codegen.GetBytecodeString(*listener))
}
