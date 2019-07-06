package main

import (
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/lab1"
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

func main() {
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
