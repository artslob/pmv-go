package tests

import (
	"github.com/artslob/pmv-go/lab2"
	"path/filepath"
	"testing"
)

type CfgTestParser struct {
}

func (CfgTestParser) Parse(got string) string {
	head := lab2.ParseInputToCFG(string(got))
	printer := lab2.NewCfgPrinter()
	printer.Print(head)
	return printer.String()
}

func (CfgTestParser) InputDir() string {
	return filepath.Join("testdata", "cfg-input")
}

func (CfgTestParser) OutputDir() string {
	return filepath.Join("testdata", "cfg-output")
}

func TestCfgListener(t *testing.T) {
	tester := CmpFileContentTester{CfgTestParser{}}
	tester.RunTest(t)
}
