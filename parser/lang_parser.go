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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 234,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 7, 2, 40, 10, 2, 12, 2, 14, 2, 43, 11, 2, 3, 2,
	3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7, 4, 52, 10, 4, 12, 4, 14, 4, 55,
	11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 7, 5, 64, 10, 5, 12, 5,
	14, 5, 67, 11, 5, 5, 5, 69, 10, 5, 3, 5, 3, 5, 3, 5, 5, 5, 74, 10, 5, 3,
	6, 3, 6, 3, 6, 5, 6, 79, 10, 6, 3, 7, 3, 7, 3, 7, 5, 7, 84, 10, 7, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 91, 10, 7, 12, 7, 14, 7, 94, 11, 7, 3, 8,
	3, 8, 3, 9, 3, 9, 3, 9, 5, 9, 101, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 5, 9, 122, 10, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 7, 14, 137, 10, 14, 12,
	14, 14, 14, 140, 11, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	7, 17, 149, 10, 17, 12, 17, 14, 17, 152, 11, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 163, 10, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 7, 18, 203, 10, 18, 12, 18, 14, 18, 206,
	11, 18, 5, 18, 208, 10, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 7,
	18, 216, 10, 18, 12, 18, 14, 18, 219, 11, 18, 5, 18, 221, 10, 18, 3, 18,
	7, 18, 224, 10, 18, 12, 18, 14, 18, 227, 11, 18, 3, 19, 3, 19, 3, 19, 5,
	19, 232, 10, 19, 3, 19, 2, 4, 12, 34, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18,
	20, 22, 24, 26, 28, 30, 32, 34, 36, 2, 10, 3, 2, 12, 19, 3, 2, 22, 23,
	4, 2, 4, 4, 24, 24, 3, 2, 28, 29, 3, 2, 47, 52, 3, 2, 31, 33, 3, 2, 34,
	35, 3, 2, 36, 37, 2, 252, 2, 41, 3, 2, 2, 2, 4, 46, 3, 2, 2, 2, 6, 48,
	3, 2, 2, 2, 8, 58, 3, 2, 2, 2, 10, 75, 3, 2, 2, 2, 12, 83, 3, 2, 2, 2,
	14, 95, 3, 2, 2, 2, 16, 121, 3, 2, 2, 2, 18, 123, 3, 2, 2, 2, 20, 126,
	3, 2, 2, 2, 22, 129, 3, 2, 2, 2, 24, 132, 3, 2, 2, 2, 26, 138, 3, 2, 2,
	2, 28, 141, 3, 2, 2, 2, 30, 144, 3, 2, 2, 2, 32, 150, 3, 2, 2, 2, 34, 162,
	3, 2, 2, 2, 36, 228, 3, 2, 2, 2, 38, 40, 5, 4, 3, 2, 39, 38, 3, 2, 2, 2,
	40, 43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 3,
	2, 2, 2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 2, 2, 3, 45, 3, 3, 2, 2, 2, 46,
	47, 5, 6, 4, 2, 47, 5, 3, 2, 2, 2, 48, 49, 7, 3, 2, 2, 49, 53, 5, 8, 5,
	2, 50, 52, 5, 16, 9, 2, 51, 50, 3, 2, 2, 2, 52, 55, 3, 2, 2, 2, 53, 51,
	3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 56, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2,
	56, 57, 7, 4, 2, 2, 57, 7, 3, 2, 2, 2, 58, 59, 7, 46, 2, 2, 59, 68, 7,
	5, 2, 2, 60, 65, 5, 10, 6, 2, 61, 62, 7, 6, 2, 2, 62, 64, 5, 10, 6, 2,
	63, 61, 3, 2, 2, 2, 64, 67, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 65, 66, 3,
	2, 2, 2, 66, 69, 3, 2, 2, 2, 67, 65, 3, 2, 2, 2, 68, 60, 3, 2, 2, 2, 68,
	69, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 73, 7, 7, 2, 2, 71, 72, 7, 8, 2,
	2, 72, 74, 5, 12, 7, 2, 73, 71, 3, 2, 2, 2, 73, 74, 3, 2, 2, 2, 74, 9,
	3, 2, 2, 2, 75, 78, 7, 46, 2, 2, 76, 77, 7, 8, 2, 2, 77, 79, 5, 12, 7,
	2, 78, 76, 3, 2, 2, 2, 78, 79, 3, 2, 2, 2, 79, 11, 3, 2, 2, 2, 80, 81,
	8, 7, 1, 2, 81, 84, 5, 14, 8, 2, 82, 84, 7, 46, 2, 2, 83, 80, 3, 2, 2,
	2, 83, 82, 3, 2, 2, 2, 84, 92, 3, 2, 2, 2, 85, 86, 12, 3, 2, 2, 86, 87,
	7, 9, 2, 2, 87, 88, 7, 10, 2, 2, 88, 89, 7, 51, 2, 2, 89, 91, 7, 11, 2,
	2, 90, 85, 3, 2, 2, 2, 91, 94, 3, 2, 2, 2, 92, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 13, 3, 2, 2, 2, 94, 92, 3, 2, 2, 2, 95, 96, 9, 2, 2, 2,
	96, 15, 3, 2, 2, 2, 97, 98, 5, 18, 10, 2, 98, 100, 5, 20, 11, 2, 99, 101,
	5, 22, 12, 2, 100, 99, 3, 2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 3, 2,
	2, 2, 102, 103, 7, 4, 2, 2, 103, 122, 3, 2, 2, 2, 104, 105, 5, 24, 13,
	2, 105, 106, 5, 26, 14, 2, 106, 107, 7, 4, 2, 2, 107, 122, 3, 2, 2, 2,
	108, 109, 5, 28, 15, 2, 109, 110, 5, 30, 16, 2, 110, 111, 7, 20, 2, 2,
	111, 122, 3, 2, 2, 2, 112, 113, 7, 21, 2, 2, 113, 122, 7, 20, 2, 2, 114,
	115, 5, 34, 18, 2, 115, 116, 7, 20, 2, 2, 116, 122, 3, 2, 2, 2, 117, 118,
	9, 3, 2, 2, 118, 119, 5, 32, 17, 2, 119, 120, 9, 4, 2, 2, 120, 122, 3,
	2, 2, 2, 121, 97, 3, 2, 2, 2, 121, 104, 3, 2, 2, 2, 121, 108, 3, 2, 2,
	2, 121, 112, 3, 2, 2, 2, 121, 114, 3, 2, 2, 2, 121, 117, 3, 2, 2, 2, 122,
	17, 3, 2, 2, 2, 123, 124, 7, 25, 2, 2, 124, 125, 5, 34, 18, 2, 125, 19,
	3, 2, 2, 2, 126, 127, 7, 26, 2, 2, 127, 128, 5, 16, 9, 2, 128, 21, 3, 2,
	2, 2, 129, 130, 7, 27, 2, 2, 130, 131, 5, 16, 9, 2, 131, 23, 3, 2, 2, 2,
	132, 133, 9, 5, 2, 2, 133, 134, 5, 34, 18, 2, 134, 25, 3, 2, 2, 2, 135,
	137, 5, 16, 9, 2, 136, 135, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136,
	3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 27, 3, 2, 2, 2, 140, 138, 3, 2,
	2, 2, 141, 142, 7, 30, 2, 2, 142, 143, 5, 16, 9, 2, 143, 29, 3, 2, 2, 2,
	144, 145, 9, 5, 2, 2, 145, 146, 5, 34, 18, 2, 146, 31, 3, 2, 2, 2, 147,
	149, 5, 16, 9, 2, 148, 147, 3, 2, 2, 2, 149, 152, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 150, 151, 3, 2, 2, 2, 151, 33, 3, 2, 2, 2, 152, 150, 3, 2,
	2, 2, 153, 154, 8, 18, 1, 2, 154, 155, 7, 55, 2, 2, 155, 163, 5, 34, 18,
	17, 156, 157, 7, 5, 2, 2, 157, 158, 5, 34, 18, 2, 158, 159, 7, 7, 2, 2,
	159, 163, 3, 2, 2, 2, 160, 163, 9, 6, 2, 2, 161, 163, 7, 46, 2, 2, 162,
	153, 3, 2, 2, 2, 162, 156, 3, 2, 2, 2, 162, 160, 3, 2, 2, 2, 162, 161,
	3, 2, 2, 2, 163, 225, 3, 2, 2, 2, 164, 165, 12, 16, 2, 2, 165, 166, 9,
	7, 2, 2, 166, 224, 5, 34, 18, 17, 167, 168, 12, 15, 2, 2, 168, 169, 9,
	8, 2, 2, 169, 224, 5, 34, 18, 16, 170, 171, 12, 14, 2, 2, 171, 172, 7,
	53, 2, 2, 172, 224, 5, 34, 18, 15, 173, 174, 12, 13, 2, 2, 174, 175, 7,
	54, 2, 2, 175, 224, 5, 34, 18, 14, 176, 177, 12, 12, 2, 2, 177, 178, 9,
	9, 2, 2, 178, 224, 5, 34, 18, 13, 179, 180, 12, 11, 2, 2, 180, 181, 7,
	38, 2, 2, 181, 224, 5, 34, 18, 12, 182, 183, 12, 10, 2, 2, 183, 184, 7,
	39, 2, 2, 184, 224, 5, 34, 18, 11, 185, 186, 12, 9, 2, 2, 186, 187, 7,
	40, 2, 2, 187, 224, 5, 34, 18, 10, 188, 189, 12, 8, 2, 2, 189, 190, 7,
	41, 2, 2, 190, 224, 5, 34, 18, 9, 191, 192, 12, 7, 2, 2, 192, 193, 7, 42,
	2, 2, 193, 224, 5, 34, 18, 8, 194, 195, 12, 6, 2, 2, 195, 196, 7, 43, 2,
	2, 196, 224, 5, 34, 18, 7, 197, 198, 12, 19, 2, 2, 198, 207, 7, 5, 2, 2,
	199, 204, 5, 34, 18, 2, 200, 201, 7, 6, 2, 2, 201, 203, 5, 34, 18, 2, 202,
	200, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 205,
	3, 2, 2, 2, 205, 208, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207, 199, 3, 2,
	2, 2, 207, 208, 3, 2, 2, 2, 208, 209, 3, 2, 2, 2, 209, 224, 7, 7, 2, 2,
	210, 211, 12, 18, 2, 2, 211, 220, 7, 10, 2, 2, 212, 217, 5, 36, 19, 2,
	213, 214, 7, 6, 2, 2, 214, 216, 5, 36, 19, 2, 215, 213, 3, 2, 2, 2, 216,
	219, 3, 2, 2, 2, 217, 215, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 221,
	3, 2, 2, 2, 219, 217, 3, 2, 2, 2, 220, 212, 3, 2, 2, 2, 220, 221, 3, 2,
	2, 2, 221, 222, 3, 2, 2, 2, 222, 224, 7, 11, 2, 2, 223, 164, 3, 2, 2, 2,
	223, 167, 3, 2, 2, 2, 223, 170, 3, 2, 2, 2, 223, 173, 3, 2, 2, 2, 223,
	176, 3, 2, 2, 2, 223, 179, 3, 2, 2, 2, 223, 182, 3, 2, 2, 2, 223, 185,
	3, 2, 2, 2, 223, 188, 3, 2, 2, 2, 223, 191, 3, 2, 2, 2, 223, 194, 3, 2,
	2, 2, 223, 197, 3, 2, 2, 2, 223, 210, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2,
	225, 223, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 35, 3, 2, 2, 2, 227, 225,
	3, 2, 2, 2, 228, 231, 5, 34, 18, 2, 229, 230, 7, 44, 2, 2, 230, 232, 5,
	34, 18, 2, 231, 229, 3, 2, 2, 2, 231, 232, 3, 2, 2, 2, 232, 37, 3, 2, 2,
	2, 22, 41, 53, 65, 68, 73, 78, 83, 92, 100, 121, 138, 150, 162, 204, 207,
	217, 220, 223, 225, 231,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'def'", "'end'", "'('", "','", "')'", "'of'", "'array'", "'['", "']'",
	"'bool'", "'byte'", "'int'", "'uint'", "'long'", "'ulong'", "'char'", "'string'",
	"';'", "'break'", "'begin'", "'{'", "'}'", "'if'", "'then'", "'else'",
	"'while'", "'until'", "'do'", "'*'", "'/'", "'%'", "'+'", "'-'", "'=='",
	"'!='", "'&'", "'^'", "'|'", "'&&'", "'||'", "'='", "'..'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "EMPTY", "IDENTIFIER", "STR", "CHAR", "HEX",
	"BITS", "DEC", "BOOL", "SHIFT_OP", "CMP_ARR_OP", "UN_OP",
}

var ruleNames = []string{
	"source", "sourceItem", "funcDef", "funcSignature", "arg", "typeRef", "built",
	"statement", "ifExpr", "ifThen", "ifElse", "whileExpr", "whileBody", "untilStatement",
	"untilExpr", "blockBody", "expr", "ranges",
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
	LangParserEMPTY      = 43
	LangParserIDENTIFIER = 44
	LangParserSTR        = 45
	LangParserCHAR       = 46
	LangParserHEX        = 47
	LangParserBITS       = 48
	LangParserDEC        = 49
	LangParserBOOL       = 50
	LangParserSHIFT_OP   = 51
	LangParserCMP_ARR_OP = 52
	LangParserUN_OP      = 53
)

// LangParser rules.
const (
	LangParserRULE_source         = 0
	LangParserRULE_sourceItem     = 1
	LangParserRULE_funcDef        = 2
	LangParserRULE_funcSignature  = 3
	LangParserRULE_arg            = 4
	LangParserRULE_typeRef        = 5
	LangParserRULE_built          = 6
	LangParserRULE_statement      = 7
	LangParserRULE_ifExpr         = 8
	LangParserRULE_ifThen         = 9
	LangParserRULE_ifElse         = 10
	LangParserRULE_whileExpr      = 11
	LangParserRULE_whileBody      = 12
	LangParserRULE_untilStatement = 13
	LangParserRULE_untilExpr      = 14
	LangParserRULE_blockBody      = 15
	LangParserRULE_expr           = 16
	LangParserRULE_ranges         = 17
)

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
	p.EnterRule(localctx, 0, LangParserRULE_source)
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
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__0 {
		{
			p.SetState(36)
			p.SourceItem()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
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
	p.EnterRule(localctx, 2, LangParserRULE_sourceItem)

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
		p.SetState(44)
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
	p.EnterRule(localctx, 4, LangParserRULE_funcDef)
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
		p.SetState(46)
		p.Match(LangParserT__0)
	}
	{
		p.SetState(47)
		p.FuncSignature()
	}
	p.SetState(51)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__18)|(1<<LangParserT__19)|(1<<LangParserT__20)|(1<<LangParserT__22)|(1<<LangParserT__25)|(1<<LangParserT__26)|(1<<LangParserT__27))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LangParserIDENTIFIER-44))|(1<<(LangParserSTR-44))|(1<<(LangParserCHAR-44))|(1<<(LangParserHEX-44))|(1<<(LangParserBITS-44))|(1<<(LangParserDEC-44))|(1<<(LangParserBOOL-44))|(1<<(LangParserUN_OP-44)))) != 0) {
		{
			p.SetState(48)
			p.Statement()
		}

		p.SetState(53)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(54)
		p.Match(LangParserT__1)
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
	p.EnterRule(localctx, 6, LangParserRULE_funcSignature)
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
		p.SetState(56)
		p.Match(LangParserIDENTIFIER)
	}
	{
		p.SetState(57)
		p.Match(LangParserT__2)
	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserIDENTIFIER {
		{
			p.SetState(58)
			p.Arg()
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LangParserT__3 {
			{
				p.SetState(59)
				p.Match(LangParserT__3)
			}
			{
				p.SetState(60)
				p.Arg()
			}

			p.SetState(65)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(68)
		p.Match(LangParserT__4)
	}
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__5 {
		{
			p.SetState(69)
			p.Match(LangParserT__5)
		}
		{
			p.SetState(70)
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
	p.EnterRule(localctx, 8, LangParserRULE_arg)
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
		p.SetState(73)
		p.Match(LangParserIDENTIFIER)
	}
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__5 {
		{
			p.SetState(74)
			p.Match(LangParserT__5)
		}
		{
			p.SetState(75)
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
	_startState := 10
	p.EnterRecursionRule(localctx, 10, LangParserRULE_typeRef, _p)

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
	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__9, LangParserT__10, LangParserT__11, LangParserT__12, LangParserT__13, LangParserT__14, LangParserT__15, LangParserT__16:
		localctx = NewBuiltinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(79)
			p.Built()
		}

	case LangParserIDENTIFIER:
		localctx = NewCustomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
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
			p.SetState(83)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(84)
				p.Match(LangParserT__6)
			}
			{
				p.SetState(85)
				p.Match(LangParserT__7)
			}
			{
				p.SetState(86)
				p.Match(LangParserDEC)
			}
			{
				p.SetState(87)
				p.Match(LangParserT__8)
			}

		}
		p.SetState(92)
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
	p.EnterRule(localctx, 12, LangParserRULE_built)
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
		p.SetState(93)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__9)|(1<<LangParserT__10)|(1<<LangParserT__11)|(1<<LangParserT__12)|(1<<LangParserT__13)|(1<<LangParserT__14)|(1<<LangParserT__15)|(1<<LangParserT__16))) != 0) {
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
	p.EnterRule(localctx, 14, LangParserRULE_statement)
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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__22:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.IfExpr()
		}
		{
			p.SetState(96)
			p.IfThen()
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LangParserT__24 {
			{
				p.SetState(97)
				p.IfElse()
			}

		}
		{
			p.SetState(100)
			p.Match(LangParserT__1)
		}

	case LangParserT__25, LangParserT__26:
		localctx = NewLoopContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(102)
			p.WhileExpr()
		}
		{
			p.SetState(103)
			p.WhileBody()
		}
		{
			p.SetState(104)
			p.Match(LangParserT__1)
		}

	case LangParserT__27:
		localctx = NewRepeatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
			p.UntilStatement()
		}
		{
			p.SetState(107)
			p.UntilExpr()
		}
		{
			p.SetState(108)
			p.Match(LangParserT__17)
		}

	case LangParserT__18:
		localctx = NewBreakContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.Match(LangParserT__18)
		}
		{
			p.SetState(111)
			p.Match(LangParserT__17)
		}

	case LangParserT__2, LangParserIDENTIFIER, LangParserSTR, LangParserCHAR, LangParserHEX, LangParserBITS, LangParserDEC, LangParserBOOL, LangParserUN_OP:
		localctx = NewExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.expr(0)
		}
		{
			p.SetState(113)
			p.Match(LangParserT__17)
		}

	case LangParserT__19, LangParserT__20:
		localctx = NewBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(115)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__19 || _la == LangParserT__20) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(116)
			p.BlockBody()
		}
		{
			p.SetState(117)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__1 || _la == LangParserT__21) {
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
	p.EnterRule(localctx, 16, LangParserRULE_ifExpr)

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
		p.SetState(121)
		p.Match(LangParserT__22)
	}
	{
		p.SetState(122)
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
	p.EnterRule(localctx, 18, LangParserRULE_ifThen)

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
		p.SetState(124)
		p.Match(LangParserT__23)
	}
	{
		p.SetState(125)
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
	p.EnterRule(localctx, 20, LangParserRULE_ifElse)

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
		p.SetState(127)
		p.Match(LangParserT__24)
	}
	{
		p.SetState(128)
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
	p.EnterRule(localctx, 22, LangParserRULE_whileExpr)
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
		p.SetState(130)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__25 || _la == LangParserT__26) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(131)
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
	p.EnterRule(localctx, 24, LangParserRULE_whileBody)
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
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__18)|(1<<LangParserT__19)|(1<<LangParserT__20)|(1<<LangParserT__22)|(1<<LangParserT__25)|(1<<LangParserT__26)|(1<<LangParserT__27))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LangParserIDENTIFIER-44))|(1<<(LangParserSTR-44))|(1<<(LangParserCHAR-44))|(1<<(LangParserHEX-44))|(1<<(LangParserBITS-44))|(1<<(LangParserDEC-44))|(1<<(LangParserBOOL-44))|(1<<(LangParserUN_OP-44)))) != 0) {
		{
			p.SetState(133)
			p.Statement()
		}

		p.SetState(138)
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
	p.EnterRule(localctx, 26, LangParserRULE_untilStatement)

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
		p.SetState(139)
		p.Match(LangParserT__27)
	}
	{
		p.SetState(140)
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
	p.EnterRule(localctx, 28, LangParserRULE_untilExpr)
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
		p.SetState(142)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__25 || _la == LangParserT__26) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(143)
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
	p.EnterRule(localctx, 30, LangParserRULE_blockBody)
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
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__18)|(1<<LangParserT__19)|(1<<LangParserT__20)|(1<<LangParserT__22)|(1<<LangParserT__25)|(1<<LangParserT__26)|(1<<LangParserT__27))) != 0) || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LangParserIDENTIFIER-44))|(1<<(LangParserSTR-44))|(1<<(LangParserCHAR-44))|(1<<(LangParserHEX-44))|(1<<(LangParserBITS-44))|(1<<(LangParserDEC-44))|(1<<(LangParserBOOL-44))|(1<<(LangParserUN_OP-44)))) != 0) {
		{
			p.SetState(145)
			p.Statement()
		}

		p.SetState(150)
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
}

func NewUnaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryContext {
	var p = new(UnaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *UnaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryContext) UN_OP() antlr.TerminalNode {
	return s.GetToken(LangParserUN_OP, 0)
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

type LiteralContext struct {
	*ExprContext
}

func NewLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralContext {
	var p = new(LiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) BOOL() antlr.TerminalNode {
	return s.GetToken(LangParserBOOL, 0)
}

func (s *LiteralContext) STR() antlr.TerminalNode {
	return s.GetToken(LangParserSTR, 0)
}

func (s *LiteralContext) CHAR() antlr.TerminalNode {
	return s.GetToken(LangParserCHAR, 0)
}

func (s *LiteralContext) HEX() antlr.TerminalNode {
	return s.GetToken(LangParserHEX, 0)
}

func (s *LiteralContext) BITS() antlr.TerminalNode {
	return s.GetToken(LangParserBITS, 0)
}

func (s *LiteralContext) DEC() antlr.TerminalNode {
	return s.GetToken(LangParserDEC, 0)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitLiteral(s)
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
	_startState := 32
	p.EnterRecursionRule(localctx, 32, LangParserRULE_expr, _p)
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
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserUN_OP:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(152)
			p.Match(LangParserUN_OP)
		}
		{
			p.SetState(153)
			p.expr(15)
		}

	case LangParserT__2:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(154)
			p.Match(LangParserT__2)
		}
		{
			p.SetState(155)
			p.expr(0)
		}
		{
			p.SetState(156)
			p.Match(LangParserT__4)
		}

	case LangParserSTR, LangParserCHAR, LangParserHEX, LangParserBITS, LangParserDEC, LangParserBOOL:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(158)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(LangParserSTR-45))|(1<<(LangParserCHAR-45))|(1<<(LangParserHEX-45))|(1<<(LangParserBITS-45))|(1<<(LangParserDEC-45))|(1<<(LangParserBOOL-45)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case LangParserIDENTIFIER:
		localctx = NewPlaceContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(159)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivModContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(162)

				if !(p.Precpred(p.GetParserRuleContext(), 14)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 14)", ""))
				}
				{
					p.SetState(163)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivModContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__28)|(1<<LangParserT__29)|(1<<LangParserT__30))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivModContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(164)
					p.expr(15)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(165)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
				}
				{
					p.SetState(166)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LangParserT__31 || _la == LangParserT__32) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(167)
					p.expr(14)
				}

			case 3:
				localctx = NewShiftContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(168)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(169)

					var _m = p.Match(LangParserSHIFT_OP)

					localctx.(*ShiftContext).op = _m
				}
				{
					p.SetState(170)
					p.expr(13)
				}

			case 4:
				localctx = NewCompareArrowContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(171)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(172)

					var _m = p.Match(LangParserCMP_ARR_OP)

					localctx.(*CompareArrowContext).op = _m
				}
				{
					p.SetState(173)
					p.expr(12)
				}

			case 5:
				localctx = NewCompareEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(174)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(175)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareEqualContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == LangParserT__33 || _la == LangParserT__34) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareEqualContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(176)
					p.expr(11)
				}

			case 6:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(177)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(178)

					var _m = p.Match(LangParserT__35)

					localctx.(*AndContext).op = _m
				}
				{
					p.SetState(179)
					p.expr(10)
				}

			case 7:
				localctx = NewXorContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(180)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(181)

					var _m = p.Match(LangParserT__36)

					localctx.(*XorContext).op = _m
				}
				{
					p.SetState(182)
					p.expr(9)
				}

			case 8:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(183)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(184)

					var _m = p.Match(LangParserT__37)

					localctx.(*OrContext).op = _m
				}
				{
					p.SetState(185)
					p.expr(8)
				}

			case 9:
				localctx = NewAndLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(186)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(187)

					var _m = p.Match(LangParserT__38)

					localctx.(*AndLogicalContext).op = _m
				}
				{
					p.SetState(188)
					p.expr(7)
				}

			case 10:
				localctx = NewOrLogicalContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(189)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(190)

					var _m = p.Match(LangParserT__39)

					localctx.(*OrLogicalContext).op = _m
				}
				{
					p.SetState(191)
					p.expr(6)
				}

			case 11:
				localctx = NewAssignContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(192)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(193)

					var _m = p.Match(LangParserT__40)

					localctx.(*AssignContext).op = _m
				}
				{
					p.SetState(194)
					p.expr(5)
				}

			case 12:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(195)

				if !(p.Precpred(p.GetParserRuleContext(), 17)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 17)", ""))
				}
				{
					p.SetState(196)
					p.Match(LangParserT__2)
				}
				p.SetState(205)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == LangParserT__2 || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LangParserIDENTIFIER-44))|(1<<(LangParserSTR-44))|(1<<(LangParserCHAR-44))|(1<<(LangParserHEX-44))|(1<<(LangParserBITS-44))|(1<<(LangParserDEC-44))|(1<<(LangParserBOOL-44))|(1<<(LangParserUN_OP-44)))) != 0) {
					{
						p.SetState(197)
						p.expr(0)
					}
					p.SetState(202)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__3 {
						{
							p.SetState(198)
							p.Match(LangParserT__3)
						}
						{
							p.SetState(199)
							p.expr(0)
						}

						p.SetState(204)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(207)
					p.Match(LangParserT__4)
				}

			case 13:
				localctx = NewSliceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(208)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(209)
					p.Match(LangParserT__7)
				}
				p.SetState(218)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if _la == LangParserT__2 || (((_la-44)&-(0x1f+1)) == 0 && ((1<<uint((_la-44)))&((1<<(LangParserIDENTIFIER-44))|(1<<(LangParserSTR-44))|(1<<(LangParserCHAR-44))|(1<<(LangParserHEX-44))|(1<<(LangParserBITS-44))|(1<<(LangParserDEC-44))|(1<<(LangParserBOOL-44))|(1<<(LangParserUN_OP-44)))) != 0) {
					{
						p.SetState(210)
						p.Ranges()
					}
					p.SetState(215)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__3 {
						{
							p.SetState(211)
							p.Match(LangParserT__3)
						}
						{
							p.SetState(212)
							p.Ranges()
						}

						p.SetState(217)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(220)
					p.Match(LangParserT__8)
				}

			}

		}
		p.SetState(225)
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
	p.EnterRule(localctx, 34, LangParserRULE_ranges)
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
		p.SetState(226)
		p.expr(0)
	}
	p.SetState(229)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__41 {
		{
			p.SetState(227)
			p.Match(LangParserT__41)
		}
		{
			p.SetState(228)
			p.expr(0)
		}

	}

	return localctx
}

func (p *LangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *TypeRefContext = nil
		if localctx != nil {
			t = localctx.(*TypeRefContext)
		}
		return p.TypeRef_Sempred(t, predIndex)

	case 16:
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
		return p.Precpred(p.GetParserRuleContext(), 14)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 17)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 16)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
