package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

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
			a += 1; until a < 10;
			{
				a += 5;
				if a > 100 then break; end
				a ^= xor;
			} until call() == 0;
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
		`
		def q() of void
			while a < 1 end
			{}
			begin end
			begin end until a > 3;
		end
		`,
	}
	for i, input := range tables {
		parsed := parseAstToString(input)
		file := filepath.Join("testdata", "print-tree", fmt.Sprintf("%d.txt", i+1))
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
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
			t.Error("parsed input not equal to expected")
		}
	}
}
