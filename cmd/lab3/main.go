package main

import (
	"fmt"
	"github.com/artslob/pmv-go/lab3"
)

func main() {
	var input = `
def func()
	a = (5 + 10) - (4 + 2);
end
`
	listener := lab3.GenerateCode(input)
	fmt.Println(lab3.GetBytecodeString(*listener))
}
