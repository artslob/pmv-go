package lab2

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"strings"
)

func ParseInputToCFG(input string) *Block {
	p := lab1.GetParser(input)
	listener := newCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.head
}

func PrintCFG(head *Block, builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("digraph G {\n")
		defer func() { builder.WriteString("}\n") }()
	}
	builder.WriteString(fmt.Sprintf("%d[label=\"%s\"]\n", head.id, head.text))
	if head.next != nil {
		builder.WriteString(fmt.Sprintf("%d->%d\n", head.id, head.next.id))
		PrintCFG(head.next, builder)
	}
	if head.branch != nil {
		builder.WriteString(fmt.Sprintf("%d->%d [style=dotted]\n", head.id, head.branch.id))
		PrintCFG(head.branch, builder)
	}
}
