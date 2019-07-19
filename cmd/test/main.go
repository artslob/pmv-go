package main

import (
	"fmt"
	"github.com/artslob/pmv-go/cfg"
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
	functionBlocks := cfg.ParseInputToCFG(input)
	fmt.Print(cfg.NewCfgPrinter().Print(functionBlocks, true).String())
}
