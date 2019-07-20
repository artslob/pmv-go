package cfg

import (
	"fmt"
	"github.com/artslob/pmv-go/cfg/blocks"
	"strings"
)

type cfgPrinter struct {
	visitedIds    map[int]struct{}
	builder       *strings.Builder
	functionNames map[string]int
}

func NewCfgPrinter() *cfgPrinter {
	return &cfgPrinter{
		visitedIds:    map[int]struct{}{},
		builder:       &strings.Builder{},
		functionNames: map[string]int{},
	}
}

func (p cfgPrinter) String() string {
	return p.builder.String()
}

func (p *cfgPrinter) Print(functionBlocks []blocks.Block, drawFuncCalls bool) *cfgPrinter {
	if p.builder.Len() == 0 {
		p.builder.WriteString("digraph G {\n")
		defer p.builder.WriteString("}\n")
	}
	if drawFuncCalls {
		for _, b := range functionBlocks {
			p.functionNames[b.(*blocks.Function).Text] = b.GetId()
		}
	}
	for _, b := range functionBlocks {
		p.print(b)
	}
	return p
}

func (p *cfgPrinter) print(block blocks.Block) {
	if _, visited := p.visitedIds[block.GetId()]; visited {
		return
	}
	p.visitedIds[block.GetId()] = struct{}{}
	p.builder.WriteString(block.String())
	if defaultBlock, ok := block.(*blocks.DefaultBlock); ok && len(defaultBlock.FunctionCalls) != 0 {
		for name := range defaultBlock.FunctionCalls {
			if id, exist := p.functionNames[name]; exist {
				p.builder.WriteString(fmt.Sprintf("%2d -> %2d [style=dotted]\n", block.GetId(), id))
			}
		}
	}
	if block.GetNext() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d\n", block.GetId(), block.GetNext().GetId()))
		p.print(block.GetNext())
	}
	if block.GetBranch() != nil {
		p.builder.WriteString(fmt.Sprintf("%2d -> %2d [color=blue1]\n", block.GetId(), block.GetBranch().GetId()))
		p.print(block.GetBranch())
	}
}
