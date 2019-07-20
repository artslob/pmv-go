package cfg

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/ast"
	"github.com/artslob/pmv-go/cfg/blocks"
)

func ParseInputToCFG(input string) []blocks.Block {
	p := ast.GetParser(input)
	listener := NewCFGListener(false)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.functionBlocks
}
