package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/ast"
	"github.com/artslob/pmv-go/cfg"
	"github.com/artslob/pmv-go/codegen"
)

func main() {
	var input = `
def func()
	a = (5 + 10) - (4 + 2);
end
`
	listener := cfg.NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, ast.GetParser(input).Source())
	fmt.Println(codegen.GetBytecodeStringFromCfg(listener))
}
