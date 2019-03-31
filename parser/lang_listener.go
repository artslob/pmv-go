// Code generated from Lang.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LangListener is a complete listener for a parse tree produced by LangParser.
type LangListener interface {
	antlr.ParseTreeListener

	// EnterSource is called when entering the source production.
	EnterSource(c *SourceContext)

	// EnterSourceItem is called when entering the sourceItem production.
	EnterSourceItem(c *SourceItemContext)

	// EnterFuncDef is called when entering the funcDef production.
	EnterFuncDef(c *FuncDefContext)

	// EnterFuncSignature is called when entering the funcSignature production.
	EnterFuncSignature(c *FuncSignatureContext)

	// EnterArg is called when entering the arg production.
	EnterArg(c *ArgContext)

	// EnterArray is called when entering the array production.
	EnterArray(c *ArrayContext)

	// EnterBuiltin is called when entering the builtin production.
	EnterBuiltin(c *BuiltinContext)

	// EnterCustom is called when entering the custom production.
	EnterCustom(c *CustomContext)

	// EnterBuilt is called when entering the built production.
	EnterBuilt(c *BuiltContext)

	// ExitSource is called when exiting the source production.
	ExitSource(c *SourceContext)

	// ExitSourceItem is called when exiting the sourceItem production.
	ExitSourceItem(c *SourceItemContext)

	// ExitFuncDef is called when exiting the funcDef production.
	ExitFuncDef(c *FuncDefContext)

	// ExitFuncSignature is called when exiting the funcSignature production.
	ExitFuncSignature(c *FuncSignatureContext)

	// ExitArg is called when exiting the arg production.
	ExitArg(c *ArgContext)

	// ExitArray is called when exiting the array production.
	ExitArray(c *ArrayContext)

	// ExitBuiltin is called when exiting the builtin production.
	ExitBuiltin(c *BuiltinContext)

	// ExitCustom is called when exiting the custom production.
	ExitCustom(c *CustomContext)

	// ExitBuilt is called when exiting the built production.
	ExitBuilt(c *BuiltContext)
}
