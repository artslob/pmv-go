package main

import (
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/parser"
	"io/ioutil"
	"log"
	"os"
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

func checkError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(2)
	}
}

func lab1() {
	CommandLine := flag.NewFlagSet("lab1", flag.ExitOnError)
	inputFilename := CommandLine.String("i", "input.txt", "Input file with text of a program to parse.")
	outputFilename := CommandLine.String("o", "", "Output file where parsed tree will be written. Default - stdout.")
	_ = CommandLine.Parse(os.Args[1:])
	data, err := ioutil.ReadFile(*inputFilename)
	checkError(err)
	input := string(data)
	output := parseToStringAst(input)
	if *outputFilename == "" {
		fmt.Print(output)
	} else {
		err := ioutil.WriteFile(*outputFilename, []byte(output), 0644)
		checkError(err)
	}
}

func main() {
	input := `
		def func(first of byte, second of File) of bool array[10]
			if a < 3 then a = 3; end
		end
	`
	fmt.Print(parseToStringAst(input))
}
