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

func main() {
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
		functionBlocks := lab2.ParseInputToCFG(string(content))
		output := lab2.NewCfgPrinter().Print(functionBlocks, false).String()
		err = ioutil.WriteFile(filepath.Join(*dir, fmt.Sprintf("%s.pmv", file)), []byte(output), 0644)
		checkError(err)
	}

	// TODO all content file with links between function calls
}
