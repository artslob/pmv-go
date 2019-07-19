package tests

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/ast"
	"github.com/artslob/pmv-go/cfg"
	"github.com/artslob/pmv-go/codegen"
	"path/filepath"
	"testing"
)

type BytecodeGeneratorTest struct {
}

func (BytecodeGeneratorTest) Parse(got string) string {
	listener := cfg.NewCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, ast.GetParser(got).Source())
	return codegen.GetBytecodeStringFromCfg(listener)
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
