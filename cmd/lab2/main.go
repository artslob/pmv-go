package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/artslob/pmv-go/cfg"
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

type ConcatenatedFiles []byte

func (cf *ConcatenatedFiles) append(bytes []byte) {
	if len(bytes) == 0 {
		return
	}
	if len(*cf) != 0 {
		*cf = append(*cf, []byte("\n")...)
	}
	*cf = append(*cf, bytes...)
}

func main() {
	filesFlag := flag.String("files", "", "Set of source filesFlag to be parsed. "+
		`Specify like this: -files "in1 in2 in3".`)
	dir := flag.String("dir", "", "Output dir for cfg representation")
	cfName := flag.String("cfName", "out.cfg", "filename for concatenated cfg")
	flag.Parse()
	flagIsRequired(filesFlag, "files")
	flagIsRequired(dir, "dir")

	exist, err := directoryExist(*dir)
	checkError(err)
	if !exist {
		checkError(errors.New(fmt.Sprintf("specified directory does not exist: %q", *dir)))
	}

	files := strings.Split(strings.TrimSpace(*filesFlag), " ")
	var cf ConcatenatedFiles

	for _, file := range files {
		file = strings.TrimSpace(file)
		if file == "" {
			continue
		}
		content, err := ioutil.ReadFile(file)
		checkError(err)
		cf.append(content)
		functionBlocks := cfg.ParseInputToCFG(string(content))
		output := cfg.NewCfgPrinter().Print(functionBlocks, false).String()
		err = ioutil.WriteFile(filepath.Join(*dir, fmt.Sprintf("%s.cfg", file)), []byte(output), 0644)
		checkError(err)
	}

	output := cfg.NewCfgPrinter().Print(cfg.ParseInputToCFG(string(cf)), true).String()
	err = ioutil.WriteFile(filepath.Join(*dir, *cfName), []byte(output), 0644)
}
