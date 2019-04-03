package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestTreePrintListener(t *testing.T) {
	var tables = []string{
		`
		def testing() of bool end

		def func(first of byte, second) of bool array[10] end
		`,
		`
		def f(first of int, arr of File array [5]) of string
			if a < 3 then {
				a += 3 * arr[0..2, 4] + ~default() & 0x0F;
			} else {
				while a > 0 {
					a = a - ! b();
					if q == 3 then break; end
				} end
			} end
		end
		`,
		`
		def q() of void
			_a = a_b_1;
			a = ident;
			a = "str with \n escapes \t";
			a = 'a';
			a = 0x1FA0;
			a = 0b0110;
			a = 010;
			a = true & false;
		end
		`,
	}
	for i, input := range tables {
		parsed := parseAstToString(input)
		file := filepath.Join("testdata", "print-tree", fmt.Sprintf("%d.txt", i+1))
		content, err := ioutil.ReadFile(file)
		if err != nil {
			t.Fatalf("failed reading file: %s", file)
		}
		expected := string(content)
		if parsed != expected {
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
			t.Error("parsed input not equal to expected")
		}
	}
}
