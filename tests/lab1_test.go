package tests

import (
	"github.com/artslob/pmv-go/lab1"
	"path/filepath"
	"testing"
)

type TreePrintParser struct {
}

func (TreePrintParser) Parse(got string) string {
	return lab1.ParseAstToString(got)
}

func (TreePrintParser) InputDir() string {
	return filepath.Join("testdata", "print-tree-input")
}

func (TreePrintParser) OutputDir() string {
	return filepath.Join("testdata", "print-tree-output")
}

func TestTreePrintListener(t *testing.T) {
	tester := CmpFileContentTester{TreePrintParser{}}
	tester.RunTest(t)
}
