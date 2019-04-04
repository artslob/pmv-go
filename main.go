package main

import (
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/parser"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func getParser(input string) *parser.LangParser {
	is := antlr.NewInputStream(input)
	lexer := parser.NewLangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	return parser.NewLangParser(stream)
}

func parseAstToString(input string) string {
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
	output := parseAstToString(input)
	if *outputFilename == "" {
		fmt.Print(output)
	} else {
		err := ioutil.WriteFile(*outputFilename, []byte(output), 0644)
		checkError(err)
	}
}

func main() {
	input := `
		def func()
			t = 1;
			c = 3;
			if t < 2 then
				t += 2;
			end
			print(t);
		end
	`
	//fmt.Print(parseAstToString(input))
	head := parseInputToCFG(input)
	var builder strings.Builder
	printCFG(head, &builder)
	fmt.Println("\n ")
	fmt.Print(builder.String())
}
