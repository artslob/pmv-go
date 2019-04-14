package lab2

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"strings"
)

func ParseInputToCFG(input string) Block {
	p := lab1.GetParser(input)
	listener := NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.start
}

func PrintCFG(block Block, builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("digraph G {\n")
		defer func() { builder.WriteString("}\n") }()
	}
	builder.WriteString(block.String())
	if block.getNext() != nil {
		builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.getId(), block.getNext().getId()))
		PrintCFG(block.getNext(), builder)
	}
	//if block.branch != nil {
	//	builder.WriteString(fmt.Sprintf("%2d -> %2d [style=dotted]\n", block.id, block.branch.id))
	//	PrintCFG(block.branch, builder)
	//}
}
