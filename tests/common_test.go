package tests

import (
	"flag"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

type TestParser interface {
	// use method under test to process input data
	Parse(got string) string
	// dir that contains files with input content for test
	InputDir() string
	// dir that contains files with expected output
	OutputDir() string
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
