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

func checkError(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(2)
	}
}

func lab1_main() {
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

var input = `
def func()
    do begin
		if a > 10 then {
			t *= 2;
			a += a;
		} end
	end until a < 5;
	print(a);
end
`

func main() {
	head := lab2.ParseInputToCFG(input)
	printer := lab2.NewCfgPrinter()
	printer.Print(head)
	fmt.Print(printer.String())
}
