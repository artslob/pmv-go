package lab3

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
)

func GenerateCode(input string) *CodeGeneratorListener {
	p := lab1.GetParser(input)
	listener := NewCodeGeneratorListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener
}
