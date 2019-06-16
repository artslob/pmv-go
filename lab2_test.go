package main

import (
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"path/filepath"
	"testing"
)

// TODO check 24 test

func TestCfgListener(t *testing.T) {
	cfgInputDir := filepath.Join("testdata", "cfg-input")
	tests, err := ioutil.ReadDir(cfgInputDir)
	if err != nil {
		t.Fatalf("failed to read files in dir %q: %q", cfgInputDir, err)
	}
	for _, testFile := range tests {
		if testFile.IsDir() {
			continue
		}
		testName := testFile.Name()
		input, err := ioutil.ReadFile(filepath.Join(cfgInputDir, testName))
		if err != nil {
			t.Fatalf("failed to read test %q: %s", testName, err)
		}
		head := lab2.ParseInputToCFG(string(input))
		printer := lab2.NewCfgPrinter()
		printer.Print(head)
		parsed := printer.String()
		file := filepath.Join("testdata", "cfg-output", testName)
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
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
		}
	}
}
