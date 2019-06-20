package tests

import (
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type TestParser interface {
	// use method under test to process input data
	Parse(got string) string
	// dir that contains files with input content for test
	InputDir() string
	// dir that contains files with expected output
	OutputDir() string
}

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

type CmpFileContentTester struct {
	TestParser
}

func (tester CmpFileContentTester) RunTest(t *testing.T) {
	inputDir := tester.InputDir()
	tests, err := ioutil.ReadDir(inputDir)
	if err != nil {
		t.Fatalf("failed to read files in dir %q: %q", inputDir, err)
	}
	for _, testFile := range tests {
		if testFile.IsDir() {
			continue
		}
		testName := testFile.Name()
		input, err := ioutil.ReadFile(filepath.Join(inputDir, testName))
		if err != nil {
			t.Fatalf("failed to read test %q: %s", testName, err)
		}
		parsed := tester.Parse(string(input))
		file := filepath.Join(tester.OutputDir(), testName)
		if *update {
			t.Logf("update golden file %s", file)
			if err := ioutil.WriteFile(file, []byte(parsed), 0644); err != nil {
				t.Fatalf("failed to update golden file: %s", err)
			}
		}
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("failed reading file: %s", file)
		}
		expected := string(content)
		if parsed != expected {
			t.Errorf("parsed input not equal to expected in test %q", testName)
			t.Logf("parsed:   %q\n", parsed)
			t.Logf("expected: %q\n", expected)
		}
	}
}

func TestCfgListener(t *testing.T) {
	tester := CmpFileContentTester{
		TestParser: CfgTestParser{},
	}
	tester.RunTest(t)
}
