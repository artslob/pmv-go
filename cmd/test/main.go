package main

import (
	"fmt"
	"github.com/artslob/pmv-go/lab2"
)

func main() {
	var input = `
def func()
	a = a + qwe() + qwe();
	print(t);
end

def qwe()
	a = 1;
	func();
	z = func() + test() + func() + t();
end
`
	functionBlocks := lab2.ParseInputToCFG(input)
	fmt.Print(lab2.NewCfgPrinter().Print(functionBlocks, true).String())
}
