package main

import (
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/lab1"
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"log"
	"os"
	"strings"
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
	t = 1;
	if a == 3 then
		if b > 1 then
			a += b;
		else
			a = 2;
		end
	else
		if q == 3 then
			q = 2;
		end
	end
	print(a);
end
`

func main() {
	head := lab2.ParseInputToCFG(input)
	var builder strings.Builder
	printer := lab2.NewCfgPrinter()
	printer.Print(head, &builder)
	fmt.Print(builder.String())
}
