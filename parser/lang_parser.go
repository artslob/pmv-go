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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 178,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 7, 2, 24, 10,
	2, 12, 2, 14, 2, 27, 11, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 7,
	4, 36, 10, 4, 12, 4, 14, 4, 39, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 7, 5, 48, 10, 5, 12, 5, 14, 5, 51, 11, 5, 5, 5, 53, 10, 5, 3,
	5, 3, 5, 3, 5, 5, 5, 58, 10, 5, 3, 6, 3, 6, 3, 6, 5, 6, 63, 10, 6, 3, 7,
	3, 7, 3, 7, 5, 7, 68, 10, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 75, 10,
	7, 12, 7, 14, 7, 78, 11, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 5, 9, 89, 10, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 96, 10, 9,
	12, 9, 14, 9, 99, 11, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 7, 9, 110, 10, 9, 12, 9, 14, 9, 113, 11, 9, 3, 9, 5, 9, 116, 10,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 123, 10, 9, 12, 9, 14, 9, 126, 11,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	137, 10, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7,
	10, 147, 10, 10, 12, 10, 14, 10, 150, 11, 10, 5, 10, 152, 10, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 7, 10, 160, 10, 10, 12, 10, 14, 10,
	163, 11, 10, 5, 10, 165, 10, 10, 3, 10, 7, 10, 168, 10, 10, 12, 10, 14,
	10, 171, 11, 10, 3, 11, 3, 11, 3, 11, 5, 11, 176, 10, 11, 3, 11, 2, 5,
	12, 16, 18, 12, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 2, 7, 3, 2, 12, 19,
	3, 2, 23, 24, 3, 2, 27, 28, 4, 2, 4, 4, 29, 29, 3, 2, 33, 38, 2, 194, 2,
	25, 3, 2, 2, 2, 4, 30, 3, 2, 2, 2, 6, 32, 3, 2, 2, 2, 8, 42, 3, 2, 2, 2,
	10, 59, 3, 2, 2, 2, 12, 67, 3, 2, 2, 2, 14, 79, 3, 2, 2, 2, 16, 115, 3,
	2, 2, 2, 18, 136, 3, 2, 2, 2, 20, 172, 3, 2, 2, 2, 22, 24, 5, 4, 3, 2,
	23, 22, 3, 2, 2, 2, 24, 27, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 26, 3,
	2, 2, 2, 26, 28, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 28, 29, 7, 2, 2, 3, 29,
	3, 3, 2, 2, 2, 30, 31, 5, 6, 4, 2, 31, 5, 3, 2, 2, 2, 32, 33, 7, 3, 2,
	2, 33, 37, 5, 8, 5, 2, 34, 36, 5, 16, 9, 2, 35, 34, 3, 2, 2, 2, 36, 39,
	3, 2, 2, 2, 37, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 40, 3, 2, 2, 2,
	39, 37, 3, 2, 2, 2, 40, 41, 7, 4, 2, 2, 41, 7, 3, 2, 2, 2, 42, 43, 7, 32,
	2, 2, 43, 52, 7, 5, 2, 2, 44, 49, 5, 10, 6, 2, 45, 46, 7, 6, 2, 2, 46,
	48, 5, 10, 6, 2, 47, 45, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2,
	2, 2, 49, 50, 3, 2, 2, 2, 50, 53, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 44,
	3, 2, 2, 2, 52, 53, 3, 2, 2, 2, 53, 54, 3, 2, 2, 2, 54, 57, 7, 7, 2, 2,
	55, 56, 7, 8, 2, 2, 56, 58, 5, 12, 7, 2, 57, 55, 3, 2, 2, 2, 57, 58, 3,
	2, 2, 2, 58, 9, 3, 2, 2, 2, 59, 62, 7, 32, 2, 2, 60, 61, 7, 8, 2, 2, 61,
	63, 5, 12, 7, 2, 62, 60, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 11, 3, 2,
	2, 2, 64, 65, 8, 7, 1, 2, 65, 68, 5, 14, 8, 2, 66, 68, 7, 32, 2, 2, 67,
	64, 3, 2, 2, 2, 67, 66, 3, 2, 2, 2, 68, 76, 3, 2, 2, 2, 69, 70, 12, 3,
	2, 2, 70, 71, 7, 9, 2, 2, 71, 72, 7, 10, 2, 2, 72, 73, 7, 37, 2, 2, 73,
	75, 7, 11, 2, 2, 74, 69, 3, 2, 2, 2, 75, 78, 3, 2, 2, 2, 76, 74, 3, 2,
	2, 2, 76, 77, 3, 2, 2, 2, 77, 13, 3, 2, 2, 2, 78, 76, 3, 2, 2, 2, 79, 80,
	9, 2, 2, 2, 80, 15, 3, 2, 2, 2, 81, 82, 8, 9, 1, 2, 82, 83, 7, 20, 2, 2,
	83, 84, 5, 18, 10, 2, 84, 85, 7, 21, 2, 2, 85, 88, 5, 16, 9, 2, 86, 87,
	7, 22, 2, 2, 87, 89, 5, 16, 9, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2,
	2, 89, 90, 3, 2, 2, 2, 90, 91, 7, 4, 2, 2, 91, 116, 3, 2, 2, 2, 92, 93,
	9, 3, 2, 2, 93, 97, 5, 18, 10, 2, 94, 96, 5, 16, 9, 2, 95, 94, 3, 2, 2,
	2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2, 98, 100,
	3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 7, 4, 2, 2, 101, 116, 3, 2, 2,
	2, 102, 103, 7, 26, 2, 2, 103, 116, 7, 25, 2, 2, 104, 105, 5, 18, 10, 2,
	105, 106, 7, 25, 2, 2, 106, 116, 3, 2, 2, 2, 107, 111, 9, 4, 2, 2, 108,
	110, 5, 16, 9, 2, 109, 108, 3, 2, 2, 2, 110, 113, 3, 2, 2, 2, 111, 109,
	3, 2, 2, 2, 111, 112, 3, 2, 2, 2, 112, 114, 3, 2, 2, 2, 113, 111, 3, 2,
	2, 2, 114, 116, 9, 5, 2, 2, 115, 81, 3, 2, 2, 2, 115, 92, 3, 2, 2, 2, 115,
	102, 3, 2, 2, 2, 115, 104, 3, 2, 2, 2, 115, 107, 3, 2, 2, 2, 116, 124,
	3, 2, 2, 2, 117, 118, 12, 6, 2, 2, 118, 119, 9, 3, 2, 2, 119, 120, 5, 18,
	10, 2, 120, 121, 7, 25, 2, 2, 121, 123, 3, 2, 2, 2, 122, 117, 3, 2, 2,
	2, 123, 126, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 124, 125, 3, 2, 2, 2, 125,
	17, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 127, 128, 8, 10, 1, 2, 128, 129,
	7, 40, 2, 2, 129, 137, 5, 18, 10, 8, 130, 131, 7, 5, 2, 2, 131, 132, 5,
	18, 10, 2, 132, 133, 7, 7, 2, 2, 133, 137, 3, 2, 2, 2, 134, 137, 9, 6,
	2, 2, 135, 137, 7, 32, 2, 2, 136, 127, 3, 2, 2, 2, 136, 130, 3, 2, 2, 2,
	136, 134, 3, 2, 2, 2, 136, 135, 3, 2, 2, 2, 137, 169, 3, 2, 2, 2, 138,
	139, 12, 9, 2, 2, 139, 140, 7, 39, 2, 2, 140, 168, 5, 18, 10, 10, 141,
	142, 12, 6, 2, 2, 142, 151, 7, 5, 2, 2, 143, 148, 5, 18, 10, 2, 144, 145,
	7, 6, 2, 2, 145, 147, 5, 18, 10, 2, 146, 144, 3, 2, 2, 2, 147, 150, 3,
	2, 2, 2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 152, 3, 2, 2,
	2, 150, 148, 3, 2, 2, 2, 151, 143, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 168, 7, 7, 2, 2, 154, 155, 12, 5, 2, 2, 155, 164,
	7, 10, 2, 2, 156, 161, 5, 20, 11, 2, 157, 158, 7, 6, 2, 2, 158, 160, 5,
	20, 11, 2, 159, 157, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2,
	2, 2, 161, 162, 3, 2, 2, 2, 162, 165, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2,
	164, 156, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2, 166,
	168, 7, 11, 2, 2, 167, 138, 3, 2, 2, 2, 167, 141, 3, 2, 2, 2, 167, 154,
	3, 2, 2, 2, 168, 171, 3, 2, 2, 2, 169, 167, 3, 2, 2, 2, 169, 170, 3, 2,
	2, 2, 170, 19, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 172, 175, 5, 18, 10, 2,
	173, 174, 7, 30, 2, 2, 174, 176, 5, 18, 10, 2, 175, 173, 3, 2, 2, 2, 175,
	176, 3, 2, 2, 2, 176, 21, 3, 2, 2, 2, 23, 25, 37, 49, 52, 57, 62, 67, 76,
	88, 97, 111, 115, 124, 136, 148, 151, 161, 164, 167, 169, 175,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'def'", "'end'", "'('", "','", "')'", "'of'", "'array'", "'['", "']'",
	"'bool'", "'byte'", "'int'", "'uint'", "'long'", "'ulong'", "'char'", "'string'",
	"'if'", "'then'", "'else'", "'while'", "'until'", "';'", "'break'", "'begin'",
	"'{'", "'}'", "'..'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "EMPTY", "IDENTIFIER", "STR",
	"CHAR", "HEX", "BITS", "DEC", "BOOL", "BIN_OP", "UN_OP",
}

var ruleNames = []string{
	"source", "sourceItem", "funcDef", "funcSignature", "arg", "typeRef", "built",
	"statement", "expr", "ranges",
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
	LangParserEMPTY      = 29
	LangParserIDENTIFIER = 30
	LangParserSTR        = 31
	LangParserCHAR       = 32
	LangParserHEX        = 33
	LangParserBITS       = 34
	LangParserDEC        = 35
	LangParserBOOL       = 36
	LangParserBIN_OP     = 37
	LangParserUN_OP      = 38
)

// LangParser rules.
const (
	LangParserRULE_source        = 0
	LangParserRULE_sourceItem    = 1
	LangParserRULE_funcDef       = 2
	LangParserRULE_funcSignature = 3
	LangParserRULE_arg           = 4
	LangParserRULE_typeRef       = 5
	LangParserRULE_built         = 6
	LangParserRULE_statement     = 7
	LangParserRULE_expr          = 8
	LangParserRULE_ranges        = 9
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
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__0 {
		{
			p.SetState(20)
			p.SourceItem()
		}

		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(26)
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
		p.SetState(28)
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
		p.SetState(30)
		p.Match(LangParserT__0)
	}
	{
		p.SetState(31)
		p.FuncSignature()
	}
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__17)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__23)|(1<<LangParserT__24)|(1<<LangParserT__25)|(1<<LangParserIDENTIFIER)|(1<<LangParserSTR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32))|(1<<(LangParserBOOL-32))|(1<<(LangParserUN_OP-32)))) != 0) {
		{
			p.SetState(32)
			p.statement(0)
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
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
		p.SetState(40)
		p.Match(LangParserIDENTIFIER)
	}
	{
		p.SetState(41)
		p.Match(LangParserT__2)
	}
	p.SetState(50)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserIDENTIFIER {
		{
			p.SetState(42)
			p.Arg()
		}
		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LangParserT__3 {
			{
				p.SetState(43)
				p.Match(LangParserT__3)
			}
			{
				p.SetState(44)
				p.Arg()
			}

			p.SetState(49)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(52)
		p.Match(LangParserT__4)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__5 {
		{
			p.SetState(53)
			p.Match(LangParserT__5)
		}
		{
			p.SetState(54)
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
		p.SetState(57)
		p.Match(LangParserIDENTIFIER)
	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__5 {
		{
			p.SetState(58)
			p.Match(LangParserT__5)
		}
		{
			p.SetState(59)
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
	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__9, LangParserT__10, LangParserT__11, LangParserT__12, LangParserT__13, LangParserT__14, LangParserT__15, LangParserT__16:
		localctx = NewBuiltinContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(63)
			p.Built()
		}

	case LangParserIDENTIFIER:
		localctx = NewCustomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(64)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(74)
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
			p.SetState(67)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(68)
				p.Match(LangParserT__6)
			}
			{
				p.SetState(69)
				p.Match(LangParserT__7)
			}
			{
				p.SetState(70)
				p.Match(LangParserDEC)
			}
			{
				p.SetState(71)
				p.Match(LangParserT__8)
			}

		}
		p.SetState(76)
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
		p.SetState(77)
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

func (s *LoopContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LoopContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *LoopContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
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

func (s *RepeatContext) Statement() IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *RepeatContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
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

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
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

func (s *IfContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IfContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *IfContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
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
	return p.statement(0)
}

func (p *LangParser) statement(_p int) (localctx IStatementContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewStatementContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IStatementContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, LangParserRULE_statement, _p)
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
	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserT__17:
		localctx = NewIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(80)
			p.Match(LangParserT__17)
		}
		{
			p.SetState(81)
			p.expr(0)
		}
		{
			p.SetState(82)
			p.Match(LangParserT__18)
		}
		{
			p.SetState(83)
			p.statement(0)
		}
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LangParserT__19 {
			{
				p.SetState(84)
				p.Match(LangParserT__19)
			}
			{
				p.SetState(85)
				p.statement(0)
			}

		}
		{
			p.SetState(88)
			p.Match(LangParserT__1)
		}

	case LangParserT__20, LangParserT__21:
		localctx = NewLoopContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(90)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__20 || _la == LangParserT__21) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(91)
			p.expr(0)
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__17)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__23)|(1<<LangParserT__24)|(1<<LangParserT__25)|(1<<LangParserIDENTIFIER)|(1<<LangParserSTR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32))|(1<<(LangParserBOOL-32))|(1<<(LangParserUN_OP-32)))) != 0) {
			{
				p.SetState(92)
				p.statement(0)
			}

			p.SetState(97)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(98)
			p.Match(LangParserT__1)
		}

	case LangParserT__23:
		localctx = NewBreakContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(100)
			p.Match(LangParserT__23)
		}
		{
			p.SetState(101)
			p.Match(LangParserT__22)
		}

	case LangParserT__2, LangParserIDENTIFIER, LangParserSTR, LangParserCHAR, LangParserHEX, LangParserBITS, LangParserDEC, LangParserBOOL, LangParserUN_OP:
		localctx = NewExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.expr(0)
		}
		{
			p.SetState(103)
			p.Match(LangParserT__22)
		}

	case LangParserT__24, LangParserT__25:
		localctx = NewBlockContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(105)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__24 || _la == LangParserT__25) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(109)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserT__17)|(1<<LangParserT__20)|(1<<LangParserT__21)|(1<<LangParserT__23)|(1<<LangParserT__24)|(1<<LangParserT__25)|(1<<LangParserIDENTIFIER)|(1<<LangParserSTR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32))|(1<<(LangParserBOOL-32))|(1<<(LangParserUN_OP-32)))) != 0) {
			{
				p.SetState(106)
				p.statement(0)
			}

			p.SetState(111)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(112)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LangParserT__1 || _la == LangParserT__26) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRepeatContext(p, NewStatementContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, LangParserRULE_statement)
			p.SetState(115)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(116)
				_la = p.GetTokenStream().LA(1)

				if !(_la == LangParserT__20 || _la == LangParserT__21) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(117)
				p.expr(0)
			}
			{
				p.SetState(118)
				p.Match(LangParserT__22)
			}

		}
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext())
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

type BinaryContext struct {
	*ExprContext
}

func NewBinaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryContext {
	var p = new(BinaryContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BinaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *BinaryContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *BinaryContext) BIN_OP() antlr.TerminalNode {
	return s.GetToken(LangParserBIN_OP, 0)
}

func (s *BinaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBinary(s)
	}
}

func (s *BinaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBinary(s)
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

func (p *LangParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *LangParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, LangParserRULE_expr, _p)
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
	p.SetState(134)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LangParserUN_OP:
		localctx = NewUnaryContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(126)
			p.Match(LangParserUN_OP)
		}
		{
			p.SetState(127)
			p.expr(6)
		}

	case LangParserT__2:
		localctx = NewBracesContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(128)
			p.Match(LangParserT__2)
		}
		{
			p.SetState(129)
			p.expr(0)
		}
		{
			p.SetState(130)
			p.Match(LangParserT__4)
		}

	case LangParserSTR, LangParserCHAR, LangParserHEX, LangParserBITS, LangParserDEC, LangParserBOOL:
		localctx = NewLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(132)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(LangParserSTR-31))|(1<<(LangParserCHAR-31))|(1<<(LangParserHEX-31))|(1<<(LangParserBITS-31))|(1<<(LangParserDEC-31))|(1<<(LangParserBOOL-31)))) != 0) {
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
			p.SetState(133)
			p.Match(LangParserIDENTIFIER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(167)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(165)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinaryContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(136)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(137)
					p.Match(LangParserBIN_OP)
				}
				{
					p.SetState(138)
					p.expr(8)
				}

			case 2:
				localctx = NewCallContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(139)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(140)
					p.Match(LangParserT__2)
				}
				p.SetState(149)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserIDENTIFIER)|(1<<LangParserSTR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32))|(1<<(LangParserBOOL-32))|(1<<(LangParserUN_OP-32)))) != 0) {
					{
						p.SetState(141)
						p.expr(0)
					}
					p.SetState(146)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__3 {
						{
							p.SetState(142)
							p.Match(LangParserT__3)
						}
						{
							p.SetState(143)
							p.expr(0)
						}

						p.SetState(148)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(151)
					p.Match(LangParserT__4)
				}

			case 3:
				localctx = NewSliceContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, LangParserRULE_expr)
				p.SetState(152)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(153)
					p.Match(LangParserT__7)
				}
				p.SetState(162)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LangParserT__2)|(1<<LangParserIDENTIFIER)|(1<<LangParserSTR))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LangParserCHAR-32))|(1<<(LangParserHEX-32))|(1<<(LangParserBITS-32))|(1<<(LangParserDEC-32))|(1<<(LangParserBOOL-32))|(1<<(LangParserUN_OP-32)))) != 0) {
					{
						p.SetState(154)
						p.Ranges()
					}
					p.SetState(159)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == LangParserT__3 {
						{
							p.SetState(155)
							p.Match(LangParserT__3)
						}
						{
							p.SetState(156)
							p.Ranges()
						}

						p.SetState(161)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(164)
					p.Match(LangParserT__8)
				}

			}

		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 18, LangParserRULE_ranges)
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
		p.SetState(170)
		p.expr(0)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LangParserT__27 {
		{
			p.SetState(171)
			p.Match(LangParserT__27)
		}
		{
			p.SetState(172)
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

	case 7:
		var t *StatementContext = nil
		if localctx != nil {
			t = localctx.(*StatementContext)
		}
		return p.Statement_Sempred(t, predIndex)

	case 8:
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

func (p *LangParser) Statement_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *LangParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
