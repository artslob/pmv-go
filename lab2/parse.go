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
	visitedIds map[int]struct{}
	builder    *strings.Builder
}

func (p cfgPrinter) String() string {
	return p.builder.String()
}

func NewCfgPrinter() *cfgPrinter {
	return &cfgPrinter{visitedIds: map[int]struct{}{}, builder: &strings.Builder{}}
}

func (p *cfgPrinter) Print(block blocks.Block) {
	if p.builder.Len() == 0 {
		p.builder.WriteString("digraph G {\n")
		defer p.builder.WriteString("}\n")
	}
	if _, visited := p.visitedIds[block.GetId()]; visited {
		return
	}
	p.visitedIds[block.GetId()] = struct{}{}
	p.builder.WriteString(block.String())
	if block.GetNext() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.GetId(), block.GetNext().GetId()))
		p.Print(block.GetNext())
	}
	if block.GetBranch() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d [style=dotted]\n", block.GetId(), block.GetBranch().GetId()))
		p.Print(block.GetBranch())
	}
}
