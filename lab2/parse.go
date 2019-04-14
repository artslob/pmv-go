package lab2

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"github.com/artslob/pmv-go/lab2/block"
	"strings"
)

func ParseInputToCFG(input string) block.Block {
	p := lab1.GetParser(input)
	listener := NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.start
}

func PrintCFG(block block.Block, builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("digraph G {\n")
		defer builder.WriteString("}\n")
	}
	builder.WriteString(block.String())
	if block.GetNext() != nil {
		builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.GetId(), block.GetNext().GetId()))
		PrintCFG(block.GetNext(), builder)
	}
	if block.GetBranch() != nil {
		builder.WriteString(fmt.Sprintf("%2d -> %2d [style=dotted]\n", block.GetId(), block.GetBranch().GetId()))
		PrintCFG(block.GetBranch(), builder)
	}
}
