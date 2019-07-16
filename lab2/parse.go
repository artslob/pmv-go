package lab2

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"github.com/artslob/pmv-go/lab2/blocks"
	"strings"
)

func ParseInputToCFG(input string) []blocks.Block {
	p := lab1.GetParser(input)
	listener := NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.functionBlocks
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

func (p *cfgPrinter) Print(functionBlocks []blocks.Block) {
	if p.builder.Len() == 0 {
		p.builder.WriteString("digraph G {\n")
		defer p.builder.WriteString("}\n")
	}
	for _, b := range functionBlocks {
		p.print(b)
	}
}

func (p *cfgPrinter) print(block blocks.Block) {
	if _, visited := p.visitedIds[block.GetId()]; visited {
		return
	}
	p.visitedIds[block.GetId()] = struct{}{}
	p.builder.WriteString(block.String())
	if block.GetNext() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.GetId(), block.GetNext().GetId()))
		p.print(block.GetNext())
	}
	if block.GetBranch() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d [color=blue1]\n", block.GetId(), block.GetBranch().GetId()))
		p.print(block.GetBranch())
	}
}
