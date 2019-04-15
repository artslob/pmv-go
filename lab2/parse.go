package lab2

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"github.com/artslob/pmv-go/lab2/blocks"
	"strings"
)

func ParseInputToCFG(input string) blocks.Block {
	p := lab1.GetParser(input)
	listener := NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.start
}

type cfgPrinter struct {
	visitedIds map[int]bool
}

func NewCfgPrinter() *cfgPrinter {
	return &cfgPrinter{visitedIds: map[int]bool{}}
}

func (printer *cfgPrinter) Print(block blocks.Block, builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("digraph G {\n")
		defer builder.WriteString("}\n")
	}
	if visited := printer.visitedIds[block.GetId()]; visited {
		return
	} else {
		printer.visitedIds[block.GetId()] = true
	}
	builder.WriteString(block.String())
	if block.GetNext() != nil {
		builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.GetId(), block.GetNext().GetId()))
		printer.Print(block.GetNext(), builder)
	}
	if block.GetBranch() != nil {
		builder.WriteString(fmt.Sprintf("%2d -> %2d [style=dotted]\n", block.GetId(), block.GetBranch().GetId()))
		printer.Print(block.GetBranch(), builder)
	}
}
