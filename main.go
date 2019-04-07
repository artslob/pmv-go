package main

import (
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/lab1"
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
