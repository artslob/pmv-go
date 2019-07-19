// Code generated from Lang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lang

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 57, 243,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 3, 7, 3, 44, 10, 3, 12,
	3, 14, 3, 47, 11, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 7, 5, 56,
	10, 5, 12, 5, 14, 5, 59, 11, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	6, 7, 6, 68, 10, 6, 12, 6, 14, 6, 71, 11, 6, 5, 6, 73, 10, 6, 3, 6, 3,
	6, 3, 6, 5, 6, 78, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 83, 10, 7, 3, 8, 3, 8,
	3, 8, 5, 8, 88, 10, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 95, 10, 8, 12,
	8, 14, 8, 98, 11, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 105, 10, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 126,
	10, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 14, 3, 15, 7, 15, 141, 10, 15, 12, 15, 14, 15, 144, 11,
	15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 7, 18, 153, 10, 18,
	12, 18, 14, 18, 156, 11, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 172, 10,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 7, 19, 212, 10, 19, 12,
	19, 14, 19, 215, 11, 19, 5, 19, 217, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 7, 19, 225, 10, 19, 12, 19, 14, 19, 228, 11, 19, 5, 19, 230,
	10, 19, 3, 19, 7, 19, 233, 10, 19, 12, 19, 14, 19, 236, 11, 19, 3, 20,
	3, 20, 3, 20, 5, 20, 241, 10, 20, 3, 20, 2, 4, 14, 36, 21, 2, 4, 6, 8,
	10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 2, 11, 3, 2,
	3, 4, 3, 2, 14, 21, 3, 2, 24, 25, 4, 2, 6, 6, 26, 26, 3, 2, 30, 31, 3,
	2, 33, 36, 3, 2, 37, 39, 3, 2, 33, 34, 3, 2, 40, 41, 2, 265, 2, 40, 3,
	2, 2, 2, 4, 45, 3, 2, 2, 2, 6, 50, 3, 2, 2, 2, 8, 52, 3, 2, 2, 2, 10, 62,
	3, 2, 2, 2, 12, 79, 3, 2, 2, 2, 14, 87, 3, 2, 2, 2, 16, 99, 3, 2, 2, 2,
	18, 125, 3, 2, 2, 2, 20, 127, 3, 2, 2, 2, 22, 130, 3, 2, 2, 2, 24, 133,
	3, 2, 2, 2, 26, 136, 3, 2, 2, 2, 28, 142, 3, 2, 2, 2, 30, 145, 3, 2, 2,
	2, 32, 148, 3, 2, 2, 2, 34, 154, 3, 2, 2, 2, 36, 171, 3, 2, 2, 2, 38, 237,
	3, 2, 2, 2, 40, 41, 9, 2, 2, 2, 41, 3, 3, 2, 2, 2, 42, 44, 5, 6, 4, 2,
	43, 42, 3, 2, 2, 2, 44, 47, 3, 2, 2, 2, 45, 43, 3, 2, 2, 2, 45, 46, 3,
	2, 2, 2, 46, 48, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 48, 49, 7, 2, 2, 3, 49,
	5, 3, 2, 2, 2, 50, 51, 5, 8, 5, 2, 51, 7, 3, 2, 2, 2, 52, 53, 7, 5, 2,
	2, 53, 57, 5, 10, 6, 2, 54, 56, 5, 18, 10, 2, 55, 54, 3, 2, 2, 2, 56, 59,
	3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2,
	59, 57, 3, 2, 2, 2, 60, 61, 7, 6, 2, 2, 61, 9, 3, 2, 2, 2, 62, 63, 7, 50,
	2, 2, 63, 72, 7, 7, 2, 2, 64, 69, 5, 12, 7, 2, 65, 66, 7, 8, 2, 2, 66,
	68, 5, 12, 7, 2, 67, 65, 3, 2, 2, 2, 68, 71, 3, 2, 2, 2, 69, 67, 3, 2,
	2, 2, 69, 70, 3, 2, 2, 2, 70, 73, 3, 2, 2, 2, 71, 69, 3, 2, 2, 2, 72, 64,
	3, 2, 2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 77, 7, 9, 2, 2,
	75, 76, 7, 10, 2, 2, 76, 78, 5, 14, 8, 2, 77, 75, 3, 2, 2, 2, 77, 78, 3,
	2, 2, 2, 78, 11, 3, 2, 2, 2, 79, 82, 7, 50, 2, 2, 80, 81, 7, 10, 2, 2,
	81, 83, 5, 14, 8, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 13, 3,
	2, 2, 2, 84, 85, 8, 8, 1, 2, 85, 88, 5, 16, 9, 2, 86, 88, 7, 50, 2, 2,
	87, 84, 3, 2, 2, 2, 87, 86, 3, 2, 2, 2, 88, 96, 3, 2, 2, 2, 89, 90, 12,
	3, 2, 2, 90, 91, 7, 11, 2, 2, 91, 92, 7, 12, 2, 2, 92, 93, 7, 55, 2, 2,
	93, 95, 7, 13, 2, 2, 94, 89, 3, 2, 2, 2, 95, 98, 3, 2, 2, 2, 96, 94, 3,
	2, 2, 2, 96, 97, 3, 2, 2, 2, 97, 15, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 99,
	100, 9, 3, 2, 2, 100, 17, 3, 2, 2, 2, 101, 102, 5, 20, 11, 2, 102, 104,
	5, 22, 12, 2, 103, 105, 5, 24, 13, 2, 104, 103, 3, 2, 2, 2, 104, 105, 3,
	2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 7, 6, 2, 2, 107, 126, 3, 2, 2,
	2, 108, 109, 5, 26, 14, 2, 109, 110, 5, 28, 15, 2, 110, 111, 7, 6, 2, 2,
	111, 126, 3, 2, 2, 2, 112, 113, 5, 30, 16, 2, 113, 114, 5, 32, 17, 2, 114,
	115, 7, 22, 2, 2, 115, 126, 3, 2, 2, 2, 116, 117, 7, 23, 2, 2, 117, 126,
	7, 22, 2, 2, 118, 119, 5, 36, 19, 2, 119, 120, 7, 22, 2, 2, 120, 126, 3,
	2, 2, 2, 121, 122, 9, 4, 2, 2, 122, 123, 5, 34, 18, 2, 123, 124, 9, 5,
	2, 2, 124, 126, 3, 2, 2, 2, 125, 101, 3, 2, 2, 2, 125, 108, 3, 2, 2, 2,
	125, 112, 3, 2, 2, 2, 125, 116, 3, 2, 2, 2, 125, 118, 3, 2, 2, 2, 125,
	121, 3, 2, 2, 2, 126, 19, 3, 2, 2, 2, 127, 128, 7, 27, 2, 2, 128, 129,
	5, 36, 19, 2, 129, 21, 3, 2, 2, 2, 130, 131, 7, 28, 2, 2, 131, 132, 5,
	18, 10, 2, 132, 23, 3, 2, 2, 2, 133, 134, 7, 29, 2, 2, 134, 135, 5, 18,
	10, 2, 135, 25, 3, 2, 2, 2, 136, 137, 9, 6, 2, 2, 137, 138, 5, 36, 19,
	2, 138, 27, 3, 2, 2, 2, 139, 141, 5, 18, 10, 2, 140, 139, 3, 2, 2, 2, 141,
	144, 3, 2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 29, 3,
	2, 2, 2, 144, 142, 3, 2, 2, 2, 145, 146, 7, 32, 2, 2, 146, 147, 5, 18,
	10, 2, 147, 31, 3, 2, 2, 2, 148, 149, 9, 6, 2, 2, 149, 150, 5, 36, 19,
	2, 150, 33, 3, 2, 2, 2, 151, 153, 5, 18, 10, 2, 152, 151, 3, 2, 2, 2, 153,
	156, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 35, 3,
	2, 2, 2, 156, 154, 3, 2, 2, 2, 157, 158, 8, 19, 1, 2, 158, 159, 9, 7, 2,
	2, 159, 172, 5, 36, 19, 22, 160, 161, 7, 7, 2, 2, 161, 162, 5, 36, 19,
	2, 162, 163, 7, 9, 2, 2, 163, 172, 3, 2, 2, 2, 164, 172, 5, 2, 2, 2, 165,
	172, 7, 51, 2, 2, 166, 172, 7, 52, 2, 2, 167, 172, 7, 53, 2, 2, 168, 172,
	7, 54, 2, 2, 169, 172, 7, 55, 2, 2, 170, 172, 7, 50, 2, 2, 171, 157, 3,
	2, 2, 2, 171, 160, 3, 2, 2, 2, 171, 164, 3, 2, 2, 2, 171, 165, 3, 2, 2,
	2, 171, 166, 3, 2, 2, 2, 171, 167, 3, 2, 2, 2, 171, 168, 3, 2, 2, 2, 171,
	169, 3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172, 234, 3, 2, 2, 2, 173, 174,
	12, 21, 2, 2, 174, 175, 9, 8, 2, 2, 175, 233, 5, 36, 19, 22, 176, 177,
	12, 20, 2, 2, 177, 178, 9, 9, 2, 2, 178, 233, 5, 36, 19, 21, 179, 180,
	12, 19, 2, 2, 180, 181, 7, 56, 2, 2, 181, 233, 5, 36, 19, 20, 182, 183,
	12, 18, 2, 2, 183, 184, 7, 57, 2, 2, 184, 233, 5, 36, 19, 19, 185, 186,
	12, 17, 2, 2, 186, 187, 9, 10, 2, 2, 187, 233, 5, 36, 19, 18, 188, 189,
	12, 16, 2, 2, 189, 190, 7, 42, 2, 2, 190, 233, 5, 36, 19, 17, 191, 192,
	12, 15, 2, 2, 192, 193, 7, 43, 2, 2, 193, 233, 5, 36, 19, 16, 194, 195,
	12, 14, 2, 2, 195, 196, 7, 44, 2, 2, 196, 233, 5, 36, 19, 15, 197, 198,
	12, 13, 2, 2, 198, 199, 7, 45, 2, 2, 199, 233, 5, 36, 19, 14, 200, 201,
	12, 12, 2, 2, 201, 202, 7, 46, 2, 2, 202, 233, 5, 36, 19, 13, 203, 204,
	12, 11, 2, 2, 204, 205, 7, 47, 2, 2, 205, 233, 5, 36, 19, 12, 206, 207,
	12, 24, 2, 2, 207, 216, 7, 7, 2, 2, 208, 213, 5, 36, 19, 2, 209, 210, 7,
	8, 2, 2, 210, 212, 5, 36, 19, 2, 211, 209, 3, 2, 2, 2, 212, 215, 3, 2,
	2, 2, 213, 211, 3, 2, 2, 2, 213, 214, 3, 2, 2, 2, 214, 217, 3, 2, 2, 2,
	215, 213, 3, 2, 2, 2, 216, 208, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217,
	218, 3, 2, 2, 2, 218, 233, 7, 9, 2, 2, 219, 220, 12, 23, 2, 2, 220, 229,
	7, 12, 2, 2, 221, 226, 5, 38, 20, 2, 222, 223, 7, 8, 2, 2, 223, 225, 5,
	38, 20, 2, 224, 222, 3, 2, 2, 2, 225, 228, 3, 2, 2, 2, 226, 224, 3, 2,
	2, 2, 226, 227, 3, 2, 2, 2, 227, 230, 3, 2, 2, 2, 228, 226, 3, 2, 2, 2,
	229, 221, 3, 2, 2, 2, 229, 230, 3, 2, 2, 2, 230, 231, 3, 2, 2, 2, 231,
	233, 7, 13, 2, 2, 232, 173, 3, 2, 2, 2, 232, 176, 3, 2, 2, 2, 232, 179,
	3, 2, 2, 2, 232, 182, 3, 2, 2, 2, 232, 185, 3, 2, 2, 2, 232, 188, 3, 2,
	2, 2, 232, 191, 3, 2, 2, 2, 232, 194, 3, 2, 2, 2, 232, 197, 3, 2, 2, 2,
	232, 200, 3, 2, 2, 2, 232, 203, 3, 2, 2, 2, 232, 206, 3, 2, 2, 2, 232,
	219, 3, 2, 2, 2, 233, 236, 3, 2, 2, 2, 234, 232, 3, 2, 2, 2, 234, 235,
	3, 2, 2, 2, 235, 37, 3, 2, 2, 2, 236, 234, 3, 2, 2, 2, 237, 240, 5, 36,
	19, 2, 238, 239, 7, 48, 2, 2, 239, 241, 5, 36, 19, 2, 240, 238, 3, 2, 2,
	2, 240, 241, 3, 2, 2, 2, 241, 39, 3, 2, 2, 2, 22, 45, 57, 69, 72, 77, 82,
	87, 96, 104, 125, 142, 154, 171, 213, 216, 226, 229, 232, 234, 240,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'true'", "'false'", "'def'", "'end'", "'('", "','", "')'", "'of'",
	"'array'", "'['", "']'", "'bool'", "'byte'", "'int'", "'uint'", "'long'",
	"'ulong'", "'char'", "'string'", "';'", "'break'", "'begin'", "'{'", "'}'",
	"'if'", "'then'", "'else'", "'while'", "'until'", "'do'", "'+'", "'-'",
	"'~'", "'!'", "'*'", "'/'", "'%'", "'=='", "'!='", "'&'", "'^'", "'|'",
	"'&&'", "'||'", "'='", "'..'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "EMPTY", "IDENTIFIER", "STR",
	"CHAR", "HEX", "BITS", "DEC", "SHIFT_OP", "CMP_ARR_OP",
}

var ruleNames = []string{
	"boolRule", "source", "sourceItem", "funcDef", "funcSignature", "arg",
	"typeRef", "built", "statement", "ifExpr", "ifThen", "ifElse", "whileExpr",
	"whileBody", "untilStatement", "untilExpr", "blockBody", "expr", "ranges",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LangParser struct {
	*antlr.BaseParser
}

func NewLangParser(input antlr.TokenStream) *LangParser {
	this := new(LangParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Lang.g4"

	return this
}

// LangParser tokens.
const (
	LangParserEOF        = antlr.TokenEOF
	LangParserT__0       = 1
	LangParserT__1       = 2
	LangParserT__2       = 3
	LangParserT__3       = 4
	LangParserT__4       = 5
	LangParserT__5       = 6
	LangParserT__6       = 7
	LangParserT__7       = 8
	LangParserT__8       = 9
	LangParserT__9       = 10
	LangParserT__10      = 11
	LangParserT__11      = 12
	LangParserT__12      = 13
	LangParserT__13      = 14
	LangParserT__14      = 15
	LangParserT__15      = 16
	LangParserT__16      = 17
	LangParserT__17      = 18
	LangParserT__18      = 19
	LangParserT__19      = 20
	LangParserT__20      = 21
	LangParserT__21      = 22
	LangParserT__22      = 23
	LangParserT__23      = 24
	LangParserT__24      = 25
	LangParserT__25      = 26
	LangParserT__26      = 27
	LangParserT__27      = 28
	LangParserT__28      = 29
	LangParserT__29      = 30
	LangParserT__30      = 31
	LangParserT__31      = 32
	LangParserT__32      = 33
	LangParserT__33      = 34
	LangParserT__34      = 35
	LangParserT__35      = 36
	LangParserT__36      = 37
	LangParserT__37      = 38
	LangParserT__38      = 39
	LangParserT__39      = 40
	LangParserT__40      = 41
	LangParserT__41      = 42
	LangParserT__42      = 43
	LangParserT__43      = 44
	LangParserT__44      = 45
	LangParserT__45      = 46
	LangParserEMPTY      = 47
	LangParserIDENTIFIER = 48
	LangParserSTR        = 49
	LangParserCHAR       = 50
	LangParserHEX        = 51
	LangParserBITS       = 52
	LangParserDEC        = 53
	LangParserSHIFT_OP   = 54
	LangParserCMP_ARR_OP = 55
)

// LangParser rules.
const (
	LangParserRULE_boolRule       = 0
	LangParserRULE_source         = 1
	LangParserRULE_sourceItem     = 2
	LangParserRULE_funcDef        = 3
	LangParserRULE_funcSignature  = 4
	LangParserRULE_arg            = 5
	LangParserRULE_typeRef        = 6
	LangParserRULE_built          = 7
	LangParserRULE_statement      = 8
	LangParserRULE_ifExpr         = 9
	LangParserRULE_ifThen         = 10
	LangParserRULE_ifElse         = 11
	LangParserRULE_whileExpr      = 12
	LangParserRULE_whileBody      = 13
	LangParserRULE_untilStatement = 14
	LangParserRULE_untilExpr      = 15
	LangParserRULE_blockBody      = 16
	LangParserRULE_expr           = 17
	LangParserRULE_ranges         = 18
)

// IBoolRuleContext is an interface to support dynamic dispatch.
type IBoolRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolRuleContext differentiates from other interfaces.
	IsBoolRuleContext()
}

type BoolRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolRuleContext() *BoolRuleContext {
	var p = new(BoolRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_boolRule
	return p
}

func (*BoolRuleContext) IsBoolRuleContext() {}

func NewBoolRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolRuleContext {
	var p = new(BoolRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_boolRule

	return p
}

func (s *BoolRuleContext) GetParser() antlr.Parser { return s.parser }
func (s *BoolRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBoolRule(s)
	}
}

func (s *BoolRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBoolRule(s)
	}
}

func (p *LangParser) BoolRule() (localctx IBoolRuleContext) {
	localctx = NewBoolRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LangParserRULE_boolRule)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__0 || _la == LangParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) EOF() antlr.TerminalNode {
	return s.GetToken(LangParserEOF, 0)
}

func (s *SourceContext) AllSourceItem() []ISourceItemContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceItemContext)(nil)).Elem())
	var tst = make([]ISourceItemContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceItemContext)
		}
	}

	return tst
}

func (s *SourceContext) SourceItem(i int) ISourceItemContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceItemContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceItemContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterSource(s)
	}
}

func (s *SourceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitSource(s)
	}
}

func (p *LangParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LangParserRULE_source)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(43)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__2 {
		{
			p.SetState(40)
			p.SourceItem()
		}

		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(46)
		p.Match(LangParserEOF)
	}

	return localctx
}

// ISourceItemContext is an interface to support dynamic dispatch.
type ISourceItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceItemContext differentiates from other interfaces.
	IsSourceItemContext()
}

type SourceItemContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceItemContext() *SourceItemContext {
	var p = new(SourceItemContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_sourceItem
	return p
}

func (*SourceItemContext) IsSourceItemContext() {}

func NewSourceItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceItemContext {
	var p = new(SourceItemContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_sourceItem

	return p
}

func (s *SourceItemContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceItemContext) FuncDef() IFuncDefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncDefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncDefContext)
}

func (s *SourceItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterSourceItem(s)
	}
}

func (s *SourceItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitSourceItem(s)
	}
}

func (p *LangParser) SourceItem() (localctx ISourceItemContext) {
	localctx = NewSourceItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LangParserRULE_sourceItem)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.FuncDef()
	}

	return localctx
}

// IFuncDefContext is an interface to support dynamic dispatch.
type IFuncDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncDefContext differentiates from other interfaces.
	IsFuncDefContext()
}

type FuncDefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncDefContext() *FuncDefContext {
	var p = new(FuncDefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_funcDef
	return p
}

func (*FuncDefContext) IsFuncDefContext() {}

func NewFuncDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncDefContext {
	var p = new(FuncDefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_funcDef

	return p
}

func (s *FuncDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncDefContext) FuncSignature() IFuncSignatureContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFuncSignatureContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFuncSignatureContext)
}

func (s *FuncDefContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *FuncDefContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *FuncDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterFuncDef(s)
	}
}

func (s *FuncDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitFuncDef(s)
	}
}

func (p *LangParser) FuncDef() (localctx IFuncDefContext) {
	localctx = NewFuncDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LangParserRULE_funcDef)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(LangParserT__2)
	}
	{
		p.SetState(51)
		p.FuncSignature()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__0)|(1<<LangParserT__1)|(1<<LangParserT__4)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__22)|(1<<LangParserT__24)|(1<<LangParserT__27)|(1<<LangParserT__28)|(1<<LangParserT__29)|(1<<LangParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserT__31-32))|(1<<(LangParserT__32-32))|(1<<(LangParserT__33-32))|(1<<(LangParserIDENTIFIER-32))|(1<<(LangParserSTR-32))|(1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32)))) != 0) {
		{
			p.SetState(52)
			p.Statement()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(LangParserT__3)
	}

	return localctx
}

// IFuncSignatureContext is an interface to support dynamic dispatch.
type IFuncSignatureContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFuncSignatureContext differentiates from other interfaces.
	IsFuncSignatureContext()
}

type FuncSignatureContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncSignatureContext() *FuncSignatureContext {
	var p = new(FuncSignatureContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_funcSignature
	return p
}

func (*FuncSignatureContext) IsFuncSignatureContext() {}

func NewFuncSignatureContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncSignatureContext {
	var p = new(FuncSignatureContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_funcSignature

	return p
}

func (s *FuncSignatureContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncSignatureContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LangParserIDENTIFIER, 0)
}

func (s *FuncSignatureContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *FuncSignatureContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *FuncSignatureContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *FuncSignatureContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncSignatureContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncSignatureContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterFuncSignature(s)
	}
}

func (s *FuncSignatureContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitFuncSignature(s)
	}
}

func (p *LangParser) FuncSignature() (localctx IFuncSignatureContext) {
	localctx = NewFuncSignatureContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LangParserRULE_funcSignature)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(LangParserIDENTIFIER)
	}
	{
		p.SetState(61)
		p.Match(LangParserT__4)
	}
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserIDENTIFIER {
		{
			p.SetState(62)
			p.Arg()
		}
		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LangParserT__5 {
			{
				p.SetState(63)
				p.Match(LangParserT__5)
			}
			{
				p.SetState(64)
				p.Arg()
			}

			p.SetState(69)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(72)
		p.Match(LangParserT__6)
	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 {
		{
			p.SetState(73)
			p.Match(LangParserT__7)
		}
		{
			p.SetState(74)
			p.typeRef(0)
		}

	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LangParserIDENTIFIER, 0)
}

func (s *ArgContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitArg(s)
	}
}

func (p *LangParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LangParserRULE_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(LangParserIDENTIFIER)
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__7 {
		{
			p.SetState(78)
			p.Match(LangParserT__7)
		}
		{
			p.SetState(79)
			p.typeRef(0)
		}

	}

	return localctx
}

// ITypeRefContext is an interface to support dynamic dispatch.
type ITypeRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeRefContext differentiates from other interfaces.
	IsTypeRefContext()
}

type TypeRefContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeRefContext() *TypeRefContext {
	var p = new(TypeRefContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_typeRef
	return p
}

func (*TypeRefContext) IsTypeRefContext() {}

func NewTypeRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeRefContext {
	var p = new(TypeRefContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_typeRef

	return p
}

func (s *TypeRefContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeRefContext) CopyFrom(ctx *TypeRefContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ArrayContext struct {
	*TypeRefContext
}

func NewArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayContext {
	var p = new(ArrayContext)

	p.TypeRefContext = NewEmptyTypeRefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRefContext))

	return p
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) TypeRef() ITypeRefContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeRefContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeRefContext)
}

func (s *ArrayContext) DEC() antlr.TerminalNode {
	return s.GetToken(LangParserDEC, 0)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitArray(s)
	}
}

type BuiltinContext struct {
	*TypeRefContext
}

func NewBuiltinContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BuiltinContext {
	var p = new(BuiltinContext)

	p.TypeRefContext = NewEmptyTypeRefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRefContext))

	return p
}

func (s *BuiltinContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltinContext) Built() IBuiltContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBuiltContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBuiltContext)
}

func (s *BuiltinContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBuiltin(s)
	}
}

func (s *BuiltinContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBuiltin(s)
	}
}

type CustomContext struct {
	*TypeRefContext
}

func NewCustomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CustomContext {
	var p = new(CustomContext)

	p.TypeRefContext = NewEmptyTypeRefContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeRefContext))

	return p
}

func (s *CustomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CustomContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LangParserIDENTIFIER, 0)
}

func (s *CustomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterCustom(s)
	}
}

func (s *CustomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitCustom(s)
	}
}

func (p *LangParser) TypeRef() (localctx ITypeRefContext) {
	return p.typeRef(0)
}

func (p *LangParser) typeRef(_p int) (localctx ITypeRefContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTypeRefContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITypeRefContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, LangParserRULE_typeRef, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(85)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__11, LangParserT__12, LangParserT__13, LangParserT__14, LangParserT__15, LangParserT__16, LangParserT__17, LangParserT__18:
		localctx = NewBuiltinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(83)
			p.Built()
		}

	case LangParserIDENTIFIER:
		localctx = NewCustomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(84)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(94)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewArrayContext(p, NewTypeRefContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LangParserRULE_typeRef)
			p.SetState(87)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(88)
				p.Match(LangParserT__8)
			}
			{
				p.SetState(89)
				p.Match(LangParserT__9)
			}
			{
				p.SetState(90)
				p.Match(LangParserDEC)
			}
			{
				p.SetState(91)
				p.Match(LangParserT__10)
			}

		}
		p.SetState(96)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
	}

	return localctx
}

// IBuiltContext is an interface to support dynamic dispatch.
type IBuiltContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBuiltContext differentiates from other interfaces.
	IsBuiltContext()
}

type BuiltContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltContext() *BuiltContext {
	var p = new(BuiltContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_built
	return p
}

func (*BuiltContext) IsBuiltContext() {}

func NewBuiltContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltContext {
	var p = new(BuiltContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_built

	return p
}

func (s *BuiltContext) GetParser() antlr.Parser { return s.parser }
func (s *BuiltContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBuilt(s)
	}
}

func (s *BuiltContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBuilt(s)
	}
}

func (p *LangParser) Built() (localctx IBuiltContext) {
	localctx = NewBuiltContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LangParserRULE_built)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__11)|(1<<LangParserT__12)|(1<<LangParserT__13)|(1<<LangParserT__14)|(1<<LangParserT__15)|(1<<LangParserT__16)|(1<<LangParserT__17)|(1<<LangParserT__18))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionContext struct {
	*StatementContext
}

func NewExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionContext {
	var p = new(ExpressionContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitExpression(s)
	}
}

type BreakContext struct {
	*StatementContext
}

func NewBreakContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakContext {
	var p = new(BreakContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BreakContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBreak(s)
	}
}

func (s *BreakContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBreak(s)
	}
}

type LoopContext struct {
	*StatementContext
}

func NewLoopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LoopContext {
	var p = new(LoopContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) WhileExpr() IWhileExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileExprContext)
}

func (s *LoopContext) WhileBody() IWhileBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhileBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhileBodyContext)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLoop(s)
	}
}

type RepeatContext struct {
	*StatementContext
}

func NewRepeatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RepeatContext {
	var p = new(RepeatContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *RepeatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RepeatContext) UntilStatement() IUntilStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntilStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntilStatementContext)
}

func (s *RepeatContext) UntilExpr() IUntilExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUntilExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUntilExprContext)
}

func (s *RepeatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterRepeat(s)
	}
}

func (s *RepeatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitRepeat(s)
	}
}

type BlockContext struct {
	*StatementContext
}

func NewBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockContext {
	var p = new(BlockContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) BlockBody() IBlockBodyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockBodyContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBlock(s)
	}
}

type IfContext struct {
	*StatementContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) IfExpr() IIfExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfExprContext)
}

func (s *IfContext) IfThen() IIfThenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfThenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfThenContext)
}

func (s *IfContext) IfElse() IIfElseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfElseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfElseContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitIf(s)
	}
}

func (p *LangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LangParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__24:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(99)
			p.IfExpr()
		}
		{
			p.SetState(100)
			p.IfThen()
		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LangParserT__26 {
			{
				p.SetState(101)
				p.IfElse()
			}

		}
		{
			p.SetState(104)
			p.Match(LangParserT__3)
		}

	case LangParserT__27, LangParserT__28:
		localctx = NewLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.WhileExpr()
		}
		{
			p.SetState(107)
			p.WhileBody()
		}
		{
			p.SetState(108)
			p.Match(LangParserT__3)
		}

	case LangParserT__29:
		localctx = NewRepeatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.UntilStatement()
		}
		{
			p.SetState(111)
			p.UntilExpr()
		}
		{
			p.SetState(112)
			p.Match(LangParserT__19)
		}

	case LangParserT__20:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Match(LangParserT__20)
		}
		{
			p.SetState(115)
			p.Match(LangParserT__19)
		}

	case LangParserT__0, LangParserT__1, LangParserT__4, LangParserT__30, LangParserT__31, LangParserT__32, LangParserT__33, LangParserIDENTIFIER, LangParserSTR, LangParserCHAR, LangParserHEX, LangParserBITS, LangParserDEC:
		localctx = NewExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)
			p.expr(0)
		}
		{
			p.SetState(117)
			p.Match(LangParserT__19)
		}

	case LangParserT__21, LangParserT__22:
		localctx = NewBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__21 || _la == LangParserT__22) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(120)
			p.BlockBody()
		}
		{
			p.SetState(121)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__3 || _la == LangParserT__23) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfExprContext is an interface to support dynamic dispatch.
type IIfExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfExprContext differentiates from other interfaces.
	IsIfExprContext()
}

type IfExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfExprContext() *IfExprContext {
	var p = new(IfExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_ifExpr
	return p
}

func (*IfExprContext) IsIfExprContext() {}

func NewIfExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfExprContext {
	var p = new(IfExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_ifExpr

	return p
}

func (s *IfExprContext) GetParser() antlr.Parser { return s.parser }

func (s *IfExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (p *LangParser) IfExpr() (localctx IIfExprContext) {
	localctx = NewIfExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LangParserRULE_ifExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.Match(LangParserT__24)
	}
	{
		p.SetState(126)
		p.expr(0)
	}

	return localctx
}

// IIfThenContext is an interface to support dynamic dispatch.
type IIfThenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfThenContext differentiates from other interfaces.
	IsIfThenContext()
}

type IfThenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfThenContext() *IfThenContext {
	var p = new(IfThenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_ifThen
	return p
}

func (*IfThenContext) IsIfThenContext() {}

func NewIfThenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfThenContext {
	var p = new(IfThenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_ifThen

	return p
}

func (s *IfThenContext) GetParser() antlr.Parser { return s.parser }

func (s *IfThenContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfThenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfThenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfThenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterIfThen(s)
	}
}

func (s *IfThenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitIfThen(s)
	}
}

func (p *LangParser) IfThen() (localctx IIfThenContext) {
	localctx = NewIfThenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LangParserRULE_ifThen)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(128)
		p.Match(LangParserT__25)
	}
	{
		p.SetState(129)
		p.Statement()
	}

	return localctx
}

// IIfElseContext is an interface to support dynamic dispatch.
type IIfElseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfElseContext differentiates from other interfaces.
	IsIfElseContext()
}

type IfElseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfElseContext() *IfElseContext {
	var p = new(IfElseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_ifElse
	return p
}

func (*IfElseContext) IsIfElseContext() {}

func NewIfElseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfElseContext {
	var p = new(IfElseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_ifElse

	return p
}

func (s *IfElseContext) GetParser() antlr.Parser { return s.parser }

func (s *IfElseContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfElseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfElseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfElseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterIfElse(s)
	}
}

func (s *IfElseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitIfElse(s)
	}
}

func (p *LangParser) IfElse() (localctx IIfElseContext) {
	localctx = NewIfElseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LangParserRULE_ifElse)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(LangParserT__26)
	}
	{
		p.SetState(132)
		p.Statement()
	}

	return localctx
}

// IWhileExprContext is an interface to support dynamic dispatch.
type IWhileExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileExprContext differentiates from other interfaces.
	IsWhileExprContext()
}

type WhileExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileExprContext() *WhileExprContext {
	var p = new(WhileExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_whileExpr
	return p
}

func (*WhileExprContext) IsWhileExprContext() {}

func NewWhileExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileExprContext {
	var p = new(WhileExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_whileExpr

	return p
}

func (s *WhileExprContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *WhileExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterWhileExpr(s)
	}
}

func (s *WhileExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitWhileExpr(s)
	}
}

func (p *LangParser) WhileExpr() (localctx IWhileExprContext) {
	localctx = NewWhileExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LangParserRULE_whileExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__27 || _la == LangParserT__28) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(135)
		p.expr(0)
	}

	return localctx
}

// IWhileBodyContext is an interface to support dynamic dispatch.
type IWhileBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhileBodyContext differentiates from other interfaces.
	IsWhileBodyContext()
}

type WhileBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileBodyContext() *WhileBodyContext {
	var p = new(WhileBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_whileBody
	return p
}

func (*WhileBodyContext) IsWhileBodyContext() {}

func NewWhileBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileBodyContext {
	var p = new(WhileBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_whileBody

	return p
}

func (s *WhileBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileBodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *WhileBodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterWhileBody(s)
	}
}

func (s *WhileBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitWhileBody(s)
	}
}

func (p *LangParser) WhileBody() (localctx IWhileBodyContext) {
	localctx = NewWhileBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LangParserRULE_whileBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__0)|(1<<LangParserT__1)|(1<<LangParserT__4)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__22)|(1<<LangParserT__24)|(1<<LangParserT__27)|(1<<LangParserT__28)|(1<<LangParserT__29)|(1<<LangParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserT__31-32))|(1<<(LangParserT__32-32))|(1<<(LangParserT__33-32))|(1<<(LangParserIDENTIFIER-32))|(1<<(LangParserSTR-32))|(1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32)))) != 0) {
		{
			p.SetState(137)
			p.Statement()
		}

		p.SetState(142)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUntilStatementContext is an interface to support dynamic dispatch.
type IUntilStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUntilStatementContext differentiates from other interfaces.
	IsUntilStatementContext()
}

type UntilStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntilStatementContext() *UntilStatementContext {
	var p = new(UntilStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_untilStatement
	return p
}

func (*UntilStatementContext) IsUntilStatementContext() {}

func NewUntilStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UntilStatementContext {
	var p = new(UntilStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_untilStatement

	return p
}

func (s *UntilStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UntilStatementContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *UntilStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntilStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UntilStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterUntilStatement(s)
	}
}

func (s *UntilStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitUntilStatement(s)
	}
}

func (p *LangParser) UntilStatement() (localctx IUntilStatementContext) {
	localctx = NewUntilStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LangParserRULE_untilStatement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(LangParserT__29)
	}
	{
		p.SetState(144)
		p.Statement()
	}

	return localctx
}

// IUntilExprContext is an interface to support dynamic dispatch.
type IUntilExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUntilExprContext differentiates from other interfaces.
	IsUntilExprContext()
}

type UntilExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUntilExprContext() *UntilExprContext {
	var p = new(UntilExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_untilExpr
	return p
}

func (*UntilExprContext) IsUntilExprContext() {}

func NewUntilExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UntilExprContext {
	var p = new(UntilExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_untilExpr

	return p
}

func (s *UntilExprContext) GetParser() antlr.Parser { return s.parser }

func (s *UntilExprContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UntilExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UntilExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UntilExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterUntilExpr(s)
	}
}

func (s *UntilExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitUntilExpr(s)
	}
}

func (p *LangParser) UntilExpr() (localctx IUntilExprContext) {
	localctx = NewUntilExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LangParserRULE_untilExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__27 || _la == LangParserT__28) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(147)
		p.expr(0)
	}

	return localctx
}

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_blockBody
	return p
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockBodyContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (p *LangParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LangParserRULE_blockBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__0)|(1<<LangParserT__1)|(1<<LangParserT__4)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__22)|(1<<LangParserT__24)|(1<<LangParserT__27)|(1<<LangParserT__28)|(1<<LangParserT__29)|(1<<LangParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserT__31-32))|(1<<(LangParserT__32-32))|(1<<(LangParserT__33-32))|(1<<(LangParserIDENTIFIER-32))|(1<<(LangParserSTR-32))|(1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32)))) != 0) {
		{
			p.SetState(149)
			p.Statement()
		}

		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivModContext struct {
	*ExprContext
	op antlr.Token
}

func NewMulDivModContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivModContext {
	var p = new(MulDivModContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MulDivModContext) GetOp() antlr.Token { return s.op }

func (s *MulDivModContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivModContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivModContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MulDivModContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivModContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterMulDivMod(s)
	}
}

func (s *MulDivModContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitMulDivMod(s)
	}
}

type OrContext struct {
	*ExprContext
	op antlr.Token
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetOp() antlr.Token { return s.op }

func (s *OrContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitOr(s)
	}
}

type AndLogicalContext struct {
	*ExprContext
	op antlr.Token
}

func NewAndLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndLogicalContext {
	var p = new(AndLogicalContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndLogicalContext) GetOp() antlr.Token { return s.op }

func (s *AndLogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndLogicalContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndLogicalContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAndLogical(s)
	}
}

func (s *AndLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAndLogical(s)
	}
}

type LiteralBitsContext struct {
	*ExprContext
}

func NewLiteralBitsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralBitsContext {
	var p = new(LiteralBitsContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralBitsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralBitsContext) BITS() antlr.TerminalNode {
	return s.GetToken(LangParserBITS, 0)
}

func (s *LiteralBitsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralBits(s)
	}
}

func (s *LiteralBitsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralBits(s)
	}
}

type ShiftContext struct {
	*ExprContext
	op antlr.Token
}

func NewShiftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ShiftContext {
	var p = new(ShiftContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ShiftContext) GetOp() antlr.Token { return s.op }

func (s *ShiftContext) SetOp(v antlr.Token) { s.op = v }

func (s *ShiftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ShiftContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ShiftContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ShiftContext) SHIFT_OP() antlr.TerminalNode {
	return s.GetToken(LangParserSHIFT_OP, 0)
}

func (s *ShiftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterShift(s)
	}
}

func (s *ShiftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitShift(s)
	}
}

type LiteralBoolContext struct {
	*ExprContext
}

func NewLiteralBoolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralBoolContext {
	var p = new(LiteralBoolContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralBoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralBoolContext) BoolRule() IBoolRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolRuleContext)
}

func (s *LiteralBoolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralBool(s)
	}
}

func (s *LiteralBoolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralBool(s)
	}
}

type CompareArrowContext struct {
	*ExprContext
	op antlr.Token
}

func NewCompareArrowContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareArrowContext {
	var p = new(CompareArrowContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompareArrowContext) GetOp() antlr.Token { return s.op }

func (s *CompareArrowContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareArrowContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareArrowContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompareArrowContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareArrowContext) CMP_ARR_OP() antlr.TerminalNode {
	return s.GetToken(LangParserCMP_ARR_OP, 0)
}

func (s *CompareArrowContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterCompareArrow(s)
	}
}

func (s *CompareArrowContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitCompareArrow(s)
	}
}

type AddSubContext struct {
	*ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type CompareEqualContext struct {
	*ExprContext
	op antlr.Token
}

func NewCompareEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareEqualContext {
	var p = new(CompareEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CompareEqualContext) GetOp() antlr.Token { return s.op }

func (s *CompareEqualContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CompareEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CompareEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterCompareEqual(s)
	}
}

func (s *CompareEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitCompareEqual(s)
	}
}

type UnaryContext struct {
	*ExprContext
	op antlr.Token
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryContext) GetOp() antlr.Token { return s.op }

func (s *UnaryContext) SetOp(v antlr.Token) { s.op = v }

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *UnaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterUnary(s)
	}
}

func (s *UnaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitUnary(s)
	}
}

type BracesContext struct {
	*ExprContext
}

func NewBracesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracesContext {
	var p = new(BracesContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BracesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracesContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BracesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBraces(s)
	}
}

func (s *BracesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBraces(s)
	}
}

type LiteralDecContext struct {
	*ExprContext
}

func NewLiteralDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralDecContext {
	var p = new(LiteralDecContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralDecContext) DEC() antlr.TerminalNode {
	return s.GetToken(LangParserDEC, 0)
}

func (s *LiteralDecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralDec(s)
	}
}

func (s *LiteralDecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralDec(s)
	}
}

type LiteralHexContext struct {
	*ExprContext
}

func NewLiteralHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralHexContext {
	var p = new(LiteralHexContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralHexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralHexContext) HEX() antlr.TerminalNode {
	return s.GetToken(LangParserHEX, 0)
}

func (s *LiteralHexContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralHex(s)
	}
}

func (s *LiteralHexContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralHex(s)
	}
}

type CallContext struct {
	*ExprContext
}

func NewCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CallContext {
	var p = new(CallContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *CallContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterCall(s)
	}
}

func (s *CallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitCall(s)
	}
}

type OrLogicalContext struct {
	*ExprContext
	op antlr.Token
}

func NewOrLogicalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrLogicalContext {
	var p = new(OrLogicalContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrLogicalContext) GetOp() antlr.Token { return s.op }

func (s *OrLogicalContext) SetOp(v antlr.Token) { s.op = v }

func (s *OrLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrLogicalContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrLogicalContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterOrLogical(s)
	}
}

func (s *OrLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitOrLogical(s)
	}
}

type SliceContext struct {
	*ExprContext
}

func NewSliceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SliceContext {
	var p = new(SliceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SliceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SliceContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SliceContext) AllRanges() []IRangesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRangesContext)(nil)).Elem())
	var tst = make([]IRangesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRangesContext)
		}
	}

	return tst
}

func (s *SliceContext) Ranges(i int) IRangesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRangesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *SliceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterSlice(s)
	}
}

func (s *SliceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitSlice(s)
	}
}

type AndContext struct {
	*ExprContext
	op antlr.Token
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetOp() antlr.Token { return s.op }

func (s *AndContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAnd(s)
	}
}

type LiteralCharContext struct {
	*ExprContext
}

func NewLiteralCharContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralCharContext {
	var p = new(LiteralCharContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralCharContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralCharContext) CHAR() antlr.TerminalNode {
	return s.GetToken(LangParserCHAR, 0)
}

func (s *LiteralCharContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralChar(s)
	}
}

func (s *LiteralCharContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralChar(s)
	}
}

type XorContext struct {
	*ExprContext
	op antlr.Token
}

func NewXorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *XorContext {
	var p = new(XorContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *XorContext) GetOp() antlr.Token { return s.op }

func (s *XorContext) SetOp(v antlr.Token) { s.op = v }

func (s *XorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *XorContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *XorContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *XorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterXor(s)
	}
}

func (s *XorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitXor(s)
	}
}

type PlaceContext struct {
	*ExprContext
}

func NewPlaceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PlaceContext {
	var p = new(PlaceContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PlaceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PlaceContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LangParserIDENTIFIER, 0)
}

func (s *PlaceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterPlace(s)
	}
}

func (s *PlaceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitPlace(s)
	}
}

type LiteralStrContext struct {
	*ExprContext
}

func NewLiteralStrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralStrContext {
	var p = new(LiteralStrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralStrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralStrContext) STR() antlr.TerminalNode {
	return s.GetToken(LangParserSTR, 0)
}

func (s *LiteralStrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteralStr(s)
	}
}

func (s *LiteralStrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteralStr(s)
	}
}

type AssignContext struct {
	*ExprContext
	op antlr.Token
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AssignContext) GetOp() antlr.Token { return s.op }

func (s *AssignContext) SetOp(v antlr.Token) { s.op = v }

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AssignContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (p *LangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 34
	p.EnterRecursionRule(localctx, 34, LangParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(169)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__30, LangParserT__31, LangParserT__32, LangParserT__33:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(156)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*UnaryContext).op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(LangParserT__30-31))|(1<<(LangParserT__31-31))|(1<<(LangParserT__32-31))|(1<<(LangParserT__33-31)))) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*UnaryContext).op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(157)
			p.expr(20)
		}

	case LangParserT__4:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(158)
			p.Match(LangParserT__4)
		}
		{
			p.SetState(159)
			p.expr(0)
		}
		{
			p.SetState(160)
			p.Match(LangParserT__6)
		}

	case LangParserT__0, LangParserT__1:
		localctx = NewLiteralBoolContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(162)
			p.BoolRule()
		}

	case LangParserSTR:
		localctx = NewLiteralStrContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(163)
			p.Match(LangParserSTR)
		}

	case LangParserCHAR:
		localctx = NewLiteralCharContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(164)
			p.Match(LangParserCHAR)
		}

	case LangParserHEX:
		localctx = NewLiteralHexContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(LangParserHEX)
		}

	case LangParserBITS:
		localctx = NewLiteralBitsContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(166)
			p.Match(LangParserBITS)
		}

	case LangParserDEC:
		localctx = NewLiteralDecContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(167)
			p.Match(LangParserDEC)
		}

	case LangParserIDENTIFIER:
		localctx = NewPlaceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(168)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(232)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 19)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 19)", ""))
				}
				{
					p.SetState(172)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la-35)&-(0x1f+1)) == 0 && ((1<<uint((_la-35)))&((1<<(LangParserT__34-35))|(1<<(LangParserT__35-35))|(1<<(LangParserT__36-35)))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(173)
					p.expr(20)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 18)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 18)", ""))
				}
				{
					p.SetState(175)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LangParserT__30 || _la == LangParserT__31) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(176)
					p.expr(19)
				}

			case 3:
				localctx = NewShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(178)

					var _m = p.Match(LangParserSHIFT_OP)

					localctx.(*ShiftContext).op = _m
				}
				{
					p.SetState(179)
					p.expr(18)
				}

			case 4:
				localctx = NewCompareArrowContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(181)

					var _m = p.Match(LangParserCMP_ARR_OP)

					localctx.(*CompareArrowContext).op = _m
				}
				{
					p.SetState(182)
					p.expr(17)
				}

			case 5:
				localctx = NewCompareEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(184)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareEqualContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LangParserT__37 || _la == LangParserT__38) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareEqualContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(185)
					p.expr(16)
				}

			case 6:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(187)

					var _m = p.Match(LangParserT__39)

					localctx.(*AndContext).op = _m
				}
				{
					p.SetState(188)
					p.expr(15)
				}

			case 7:
				localctx = NewXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(190)

					var _m = p.Match(LangParserT__40)

					localctx.(*XorContext).op = _m
				}
				{
					p.SetState(191)
					p.expr(14)
				}

			case 8:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(193)

					var _m = p.Match(LangParserT__41)

					localctx.(*OrContext).op = _m
				}
				{
					p.SetState(194)
					p.expr(13)
				}

			case 9:
				localctx = NewAndLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(196)

					var _m = p.Match(LangParserT__42)

					localctx.(*AndLogicalContext).op = _m
				}
				{
					p.SetState(197)
					p.expr(12)
				}

			case 10:
				localctx = NewOrLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(198)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(199)

					var _m = p.Match(LangParserT__43)

					localctx.(*OrLogicalContext).op = _m
				}
				{
					p.SetState(200)
					p.expr(11)
				}

			case 11:
				localctx = NewAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(201)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(202)

					var _m = p.Match(LangParserT__44)

					localctx.(*AssignContext).op = _m
				}
				{
					p.SetState(203)
					p.expr(10)
				}

			case 12:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(204)

				if !(p.Precpred(p.GetParserRuleContext(), 22)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 22)", ""))
				}
				{
					p.SetState(205)
					p.Match(LangParserT__4)
				}
				p.SetState(214)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__0)|(1<<LangParserT__1)|(1<<LangParserT__4)|(1<<LangParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserT__31-32))|(1<<(LangParserT__32-32))|(1<<(LangParserT__33-32))|(1<<(LangParserIDENTIFIER-32))|(1<<(LangParserSTR-32))|(1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32)))) != 0) {
					{
						p.SetState(206)
						p.expr(0)
					}
					p.SetState(211)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__5 {
						{
							p.SetState(207)
							p.Match(LangParserT__5)
						}
						{
							p.SetState(208)
							p.expr(0)
						}

						p.SetState(213)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(216)
					p.Match(LangParserT__6)
				}

			case 13:
				localctx = NewSliceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(217)

				if !(p.Precpred(p.GetParserRuleContext(), 21)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 21)", ""))
				}
				{
					p.SetState(218)
					p.Match(LangParserT__9)
				}
				p.SetState(227)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__0)|(1<<LangParserT__1)|(1<<LangParserT__4)|(1<<LangParserT__30))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserT__31-32))|(1<<(LangParserT__32-32))|(1<<(LangParserT__33-32))|(1<<(LangParserIDENTIFIER-32))|(1<<(LangParserSTR-32))|(1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32)))) != 0) {
					{
						p.SetState(219)
						p.Ranges()
					}
					p.SetState(224)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__5 {
						{
							p.SetState(220)
							p.Match(LangParserT__5)
						}
						{
							p.SetState(221)
							p.Ranges()
						}

						p.SetState(226)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(229)
					p.Match(LangParserT__10)
				}

			}

		}
		p.SetState(234)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LangParserRULE_ranges
	return p
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RangesContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterRanges(s)
	}
}

func (s *RangesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitRanges(s)
	}
}

func (p *LangParser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LangParserRULE_ranges)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)
		p.expr(0)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__45 {
		{
			p.SetState(236)
			p.Match(LangParserT__45)
		}
		{
			p.SetState(237)
			p.expr(0)
		}

	}

	return localctx
}

func (p *LangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *TypeRefContext = nil
		if localctx != nil {
			t = localctx.(*TypeRefContext)
		}
		return p.TypeRef_Sempred(t, predIndex)

	case 17:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LangParser) TypeRef_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 19)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 18)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 15)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 22)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 21)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
