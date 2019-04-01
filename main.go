package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv/parser"
	"strings"
)

type treePrintListener struct {
	*parser.BaseLangListener
	b         *strings.Builder
	ruleNames []string
	level     int
}

func newTreePrintListener(ruleNames []string) *treePrintListener {
	return &treePrintListener{b: &strings.Builder{}, ruleNames: ruleNames}
}

func (s *treePrintListener) String() string {
	return s.b.String()
}

func (s *treePrintListener) print(str string) {
	s.b.WriteString(str)
}

func (s *treePrintListener) printLineBreak() {
	s.b.WriteString("\n")
}

func (s *treePrintListener) println(str string) {
	s.print(str)
	s.printLineBreak()
}

func (s *treePrintListener) getIndent() string {
	indent := "  "
	return strings.Repeat(indent, s.level)
}

func (s *treePrintListener) printIndent() {
	s.print(s.getIndent())
}

func (s *treePrintListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	s.printIndent()
	s.println(s.ruleNames[ctx.GetRuleIndex()])
	s.level++
}

func (s *treePrintListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	s.level--
}

func (s *treePrintListener) VisitTerminal(node antlr.TerminalNode) {
	s.printIndent()
	s.print(node.GetText())
	if node.GetSymbol().GetTokenType() < parser.LangParserIDENTIFIER {
		s.printLineBreak()
		return
	}
	if token := node.GetSymbol(); token != nil {
		s.print(fmt.Sprintf("  [line %d, offset: %d]\n", token.GetLine(), token.GetColumn()))
	}
}

func (s *treePrintListener) ExitTerminal(node antlr.TerminalNode) {
	s.level--
}

func main() {
	input := `
		def testing() of bool end
		
		def func(first of byte, second) of bool array[10] end
	`
	is := antlr.NewInputStream(input)
	lexer := parser.NewLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewLangParser(stream)
	listener := newTreePrintListener(p.GetRuleNames())
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	fmt.Print(listener)
}
