package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv/parser"
)

func parseToStringAst(input string) string {
	is := antlr.NewInputStream(input)
	lexer := parser.NewLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewLangParser(stream)
	listener := newTreePrintListener(p.GetRuleNames())
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.String()
}

func main() {
	input := `
		def testing() of bool end

		def func(first of byte, second) of bool array[10] end
	`
	fmt.Print(parseToStringAst(input))
}
