package tests

import (
	"github.com/artslob/pmv-go/lab3"
	"path/filepath"
	"testing"
)

type BytecodeGeneratorTest struct {
}

func (BytecodeGeneratorTest) Parse(got string) string {
	listener := lab3.GenerateCode(got)
	return lab3.GetBytecodeString(*listener)
}

func (BytecodeGeneratorTest) InputDir() string {
	return filepath.Join("testdata", "bytecode-gen-input")
}

func (BytecodeGeneratorTest) OutputDir() string {
	return filepath.Join("testdata", "bytecode-gen-output")
}

func TestBytecodeGenerator(t *testing.T) {
	tester := CmpFileContentTester{BytecodeGeneratorTest{}}
	tester.RunTest(t)
}
