package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/parser"
)

func GetParser(input string) *parser.LangParser {
	is := antlr.NewInputStream(input)
	lexer := parser.NewLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewLangParser(stream)
}

func ParseAstToString(input string) string {
	p := GetParser(input)
	listener := newTreePrintListener(p.GetRuleNames())
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.String()
}
