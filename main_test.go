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
		{
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
			`source
  sourceItem
	funcDef
	  def
	  funcSignature
		f  [line 2, offset: 7]
		(
		arg
		  first  [line 2, offset: 9]
		  of
		  typeRef
			built
			  int
		,
		arg
		  arr  [line 2, offset: 23]
		  of
		  typeRef
			typeRef
			  File  [line 2, offset: 30]
			array
			[
			5  [line 2, offset: 42]
			]
		)
		of
		typeRef
		  built
			string
	  statement
		if
		expr
		  expr
			a  [line 3, offset: 7]
		  <  [line 3, offset: 9]
		  expr
			3  [line 3, offset: 11]
		then
		statement
		  {
		  statement
			expr
			  expr
				expr
				  expr
					expr
					  expr
						expr
						  a  [line 4, offset: 5]
						+=  [line 4, offset: 7]
						expr
						  3  [line 4, offset: 10]
					  *  [line 4, offset: 12]
					  expr
						arr  [line 4, offset: 14]
					[
					ranges
					  expr
						0  [line 4, offset: 18]
					  ..
					  expr
						2  [line 4, offset: 21]
					,
					ranges
					  expr
						4  [line 4, offset: 24]
					]
				  +  [line 4, offset: 27]
				  expr
					~  [line 4, offset: 29]
					expr
					  default  [line 4, offset: 30]
				(
				)
			  &  [line 4, offset: 40]
			  expr
				0x0F  [line 4, offset: 42]
			;
		  }
		else
		statement
		  {
		  statement
			while
			expr
			  expr
				a  [line 6, offset: 11]
			  >  [line 6, offset: 13]
			  expr
				0  [line 6, offset: 15]
			statement
			  {
			  statement
				expr
				  expr
					expr
					  expr
						a  [line 7, offset: 6]
					  =  [line 7, offset: 8]
					  expr
						a  [line 7, offset: 10]
					-  [line 7, offset: 12]
					expr
					  !  [line 7, offset: 14]
					  expr
						b  [line 7, offset: 16]
				  (
				  )
				;
			  statement
				if
				expr
				  expr
					q  [line 8, offset: 9]
				  ==  [line 8, offset: 11]
				  expr
					3  [line 8, offset: 14]
				then
				statement
				  break
				  ;
				end
			  }
			end
		  }
		end
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
