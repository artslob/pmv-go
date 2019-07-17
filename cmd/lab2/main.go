package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func simple() {
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

def qwe()
	a = 1;
end
`
	head := lab2.ParseInputToCFG(input)
	printer := lab2.NewCfgPrinter()
	printer.Print(head)
	fmt.Print(printer.String())
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func flagIsRequired(f *string, name string) {
	if strings.TrimSpace(*f) == "" {
		log.Fatalf("%s flag is required", name)
	}
}

func directoryExist(path string) (bool, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if stat.Mode().IsDir() {
		return true, nil
	}
	return false, nil
}

func lab2Main() {
	filesFlag := flag.String("files", "", "Set of source filesFlag to be parsed. "+
		`Specify like this: -files "in1 in2 in3".`)
	dir := flag.String("dir", "", "Output dir for cfg representation")
	flag.Parse()
	flagIsRequired(filesFlag, "files")
	flagIsRequired(dir, "dir")

	exist, err := directoryExist(*dir)
	checkError(err)
	if !exist {
		checkError(errors.New(fmt.Sprintf("specified directory does not exist: %q", *dir)))
	}

	files := strings.Split(strings.TrimSpace(*filesFlag), " ")
	var allContent []byte

	for _, file := range files {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		content, err := ioutil.ReadFile(file)
		checkError(err)
		allContent = append(allContent, content...)
		head := lab2.ParseInputToCFG(string(content))
		printer := lab2.NewCfgPrinter()
		printer.Print(head)
		output := printer.String()
		err = ioutil.WriteFile(filepath.Join(*dir, fmt.Sprintf("%s.pmv", file)), []byte(output), 0644)
		checkError(err)
	}

	// TODO all content file with links between function calls
}

func main() {
	lab2Main()
}
