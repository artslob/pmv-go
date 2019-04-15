package main

import (
	"fmt"
	"github.com/artslob/pmv-go/lab2"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestCfgListener(t *testing.T) {
	for i := 1; i <= 11; i++ {
		inputFile := filepath.Join("testdata", "cfg-input", fmt.Sprintf("%d.txt", i))
		input, err := ioutil.ReadFile(inputFile)
		if err != nil {
			t.Fatalf("failed to read test %d input: %s", i, err)
		}
		head := lab2.ParseInputToCFG(string(input))
		var builder strings.Builder
		printer := lab2.NewCfgPrinter()
		printer.Print(head, &builder)
		parsed := builder.String()
		file := filepath.Join("testdata", "cfg-output", fmt.Sprintf("%d.txt", i))
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
			t.Errorf("parsed input not equal to expected in test %d", i)
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
		}
	}
}
