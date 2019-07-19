package tests

import (
	"github.com/artslob/pmv-go/cfg"
	"path/filepath"
	"testing"
)

type CfgTestParser struct {
}

func (CfgTestParser) Parse(got string) string {
	head := cfg.ParseInputToCFG(string(got))
	return cfg.NewCfgPrinter().Print(head, false).String()
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
