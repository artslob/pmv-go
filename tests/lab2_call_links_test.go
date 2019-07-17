package tests

import (
	"github.com/artslob/pmv-go/lab2"
	"path/filepath"
	"testing"
)

type CfgLinksTestParser struct {
}

func (CfgLinksTestParser) Parse(got string) string {
	head := lab2.ParseInputToCFG(string(got))
	return lab2.NewCfgPrinter().Print(head, true).String()
}

func (CfgLinksTestParser) InputDir() string {
	return filepath.Join("testdata", "cfg-links-input")
}

func (CfgLinksTestParser) OutputDir() string {
	return filepath.Join("testdata", "cfg-links-output")
}

func TestCfgLinksListener(t *testing.T) {
	tester := CmpFileContentTester{CfgLinksTestParser{}}
	tester.RunTest(t)
}
