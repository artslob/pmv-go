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
	var tables = []string{
		`
		def func()
			t = 1;
			c = 3;
			if t < 2 then
				t += 2;
			end
			print(t);
		end
		`,
	}
	for i, input := range tables {
		head := lab2.ParseInputToCFG(input)
		var builder strings.Builder
		lab2.PrintCFG(head, &builder)
		parsed := builder.String()
		file := filepath.Join("testdata", "cfg", fmt.Sprintf("%d.txt", i+1))
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
			t.Errorf("parsed input not equal to expected in test %d", i+1)
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
		}
	}
}
