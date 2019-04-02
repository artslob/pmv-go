package main

import (
	"strings"
	"testing"
)

func TestTreePrintListener(t *testing.T) {
	var tables = []struct {
		input, expected string
	}{
		{`
			def testing() of bool end

			def func(first of byte, second) of bool array[10] end
			`,
			`source
  sourceItem
	funcDef
	  def
	  funcSignature
		testing  [line 2, offset: 7]
		(
		)
		of
		typeRef
		  built
			bool
	  end
  sourceItem
	funcDef
	  def
	  funcSignature
		func  [line 4, offset: 7]
		(
		arg
		  first  [line 4, offset: 12]
		  of
		  typeRef
			built
			  byte
		,
		arg
		  second  [line 4, offset: 27]
		)
		of
		typeRef
		  typeRef
			built
			  bool
		  array
		  [
		  10  [line 4, offset: 49]
		  ]
	  end
  <EOF>
`,
		},
	}
	for _, testCase := range tables {
		parsed := parseToStringAst(testCase.input)
		expected := strings.ReplaceAll(testCase.expected, "\t", "    ")
		if parsed != expected {
			t.Logf("%q\n", expected)
			t.Logf("%q\n", parsed)
			t.Error("parsed input not equal to expected")
		}
	}
}
