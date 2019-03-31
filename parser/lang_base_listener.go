// Code generated from Lang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLangListener is a complete listener for a parse tree produced by LangParser.
type BaseLangListener struct{}

var _ LangListener = &BaseLangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterSource is called when production source is entered.
func (s *BaseLangListener) EnterSource(ctx *SourceContext) {}

// ExitSource is called when production source is exited.
func (s *BaseLangListener) ExitSource(ctx *SourceContext) {}

// EnterSourceItem is called when production sourceItem is entered.
func (s *BaseLangListener) EnterSourceItem(ctx *SourceItemContext) {}

// ExitSourceItem is called when production sourceItem is exited.
func (s *BaseLangListener) ExitSourceItem(ctx *SourceItemContext) {}

// EnterFuncDef is called when production funcDef is entered.
func (s *BaseLangListener) EnterFuncDef(ctx *FuncDefContext) {}

// ExitFuncDef is called when production funcDef is exited.
func (s *BaseLangListener) ExitFuncDef(ctx *FuncDefContext) {}

// EnterFuncSignature is called when production funcSignature is entered.
func (s *BaseLangListener) EnterFuncSignature(ctx *FuncSignatureContext) {}

// ExitFuncSignature is called when production funcSignature is exited.
func (s *BaseLangListener) ExitFuncSignature(ctx *FuncSignatureContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseLangListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseLangListener) ExitArg(ctx *ArgContext) {}

// EnterArray is called when production array is entered.
func (s *BaseLangListener) EnterArray(ctx *ArrayContext) {}

// ExitArray is called when production array is exited.
func (s *BaseLangListener) ExitArray(ctx *ArrayContext) {}

// EnterBuiltin is called when production builtin is entered.
func (s *BaseLangListener) EnterBuiltin(ctx *BuiltinContext) {}

// ExitBuiltin is called when production builtin is exited.
func (s *BaseLangListener) ExitBuiltin(ctx *BuiltinContext) {}

// EnterCustom is called when production custom is entered.
func (s *BaseLangListener) EnterCustom(ctx *CustomContext) {}

// ExitCustom is called when production custom is exited.
func (s *BaseLangListener) ExitCustom(ctx *CustomContext) {}

// EnterBuilt is called when production built is entered.
func (s *BaseLangListener) EnterBuilt(ctx *BuiltContext) {}

// ExitBuilt is called when production built is exited.
func (s *BaseLangListener) ExitBuilt(ctx *BuiltContext) {}
