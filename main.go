package main

import (
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/lab1"
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"log"
	"os"
)

var (
	// to disable warning about unused functions
	_ = lab1Main
	_ = lab2Main
)

func checkError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(2)
	}
}

func lab1Main() {
	CommandLine := flag.NewFlagSet("lab1", flag.ExitOnError)
	inputFilename := CommandLine.String("i", "input.txt", "Input file with text of a program to parse.")
	outputFilename := CommandLine.String("o", "", "Output file where parsed tree will be written. Default - stdout.")
	_ = CommandLine.Parse(os.Args[1:])
	data, err := ioutil.ReadFile(*inputFilename)
	checkError(err)
	input := string(data)
	output := lab1.ParseAstToString(input)
	if *outputFilename == "" {
		fmt.Print(output)
	} else {
		err := ioutil.WriteFile(*outputFilename, []byte(output), 0644)
		checkError(err)
	}
}

func lab2Main() {
	var input = `
def func()
	do {
		a = a + b;
		if b > 10 then
			break;
		end
		do {
			b = b + a;
			if b == 10 then
				break;
			end
		}
		until b < 42;
		b = b * 2;
	} until a < 3;
	print(t);
end
`
	head := lab2.ParseInputToCFG(input)
	printer := lab2.NewCfgPrinter()
	printer.Print(head)
	fmt.Print(printer.String())
}

func main() {
	lab2Main()
}
