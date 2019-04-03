package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/parser"
)

func getParser(input string) *parser.LangParser {
	is := antlr.NewInputStream(input)
	lexer := parser.NewLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewLangParser(stream)
}

func parseToStringAst(input string) string {
	p := getParser(input)
	listener := newTreePrintListener(p.GetRuleNames())
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.String()
}

func main() {
	input := `
		def func(first of byte, second of File) of bool array[10]
			if a < 3 then a = 3; end
		end
	`
	fmt.Print(parseToStringAst(input))
}
