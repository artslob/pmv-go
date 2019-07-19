package tests

import (
	"github.com/artslob/pmv-go/cfg"
	"path/filepath"
	"testing"
)

type CfgLinksTestParser struct {
}

func (CfgLinksTestParser) Parse(got string) string {
	head := cfg.ParseInputToCFG(string(got))
	return cfg.NewCfgPrinter().Print(head, true).String()
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
