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

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterRepeat is called when entering the repeat production.
	EnterRepeat(c *RepeatContext)

	// EnterBreak is called when entering the break production.
	EnterBreak(c *BreakContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterIfExpr is called when entering the ifExpr production.
	EnterIfExpr(c *IfExprContext)

	// EnterIfThen is called when entering the ifThen production.
	EnterIfThen(c *IfThenContext)

	// EnterIfElse is called when entering the ifElse production.
	EnterIfElse(c *IfElseContext)

	// EnterWhileExpr is called when entering the whileExpr production.
	EnterWhileExpr(c *WhileExprContext)

	// EnterWhileBody is called when entering the whileBody production.
	EnterWhileBody(c *WhileBodyContext)

	// EnterUntilStatement is called when entering the untilStatement production.
	EnterUntilStatement(c *UntilStatementContext)

	// EnterUntilExpr is called when entering the untilExpr production.
	EnterUntilExpr(c *UntilExprContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterSlice is called when entering the slice production.
	EnterSlice(c *SliceContext)

	// EnterBinary is called when entering the binary production.
	EnterBinary(c *BinaryContext)

	// EnterUnary is called when entering the unary production.
	EnterUnary(c *UnaryContext)

	// EnterPlace is called when entering the place production.
	EnterPlace(c *PlaceContext)

	// EnterBraces is called when entering the braces production.
	EnterBraces(c *BracesContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterRanges is called when entering the ranges production.
	EnterRanges(c *RangesContext)

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

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitRepeat is called when exiting the repeat production.
	ExitRepeat(c *RepeatContext)

	// ExitBreak is called when exiting the break production.
	ExitBreak(c *BreakContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitIfExpr is called when exiting the ifExpr production.
	ExitIfExpr(c *IfExprContext)

	// ExitIfThen is called when exiting the ifThen production.
	ExitIfThen(c *IfThenContext)

	// ExitIfElse is called when exiting the ifElse production.
	ExitIfElse(c *IfElseContext)

	// ExitWhileExpr is called when exiting the whileExpr production.
	ExitWhileExpr(c *WhileExprContext)

	// ExitWhileBody is called when exiting the whileBody production.
	ExitWhileBody(c *WhileBodyContext)

	// ExitUntilStatement is called when exiting the untilStatement production.
	ExitUntilStatement(c *UntilStatementContext)

	// ExitUntilExpr is called when exiting the untilExpr production.
	ExitUntilExpr(c *UntilExprContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitSlice is called when exiting the slice production.
	ExitSlice(c *SliceContext)

	// ExitBinary is called when exiting the binary production.
	ExitBinary(c *BinaryContext)

	// ExitUnary is called when exiting the unary production.
	ExitUnary(c *UnaryContext)

	// ExitPlace is called when exiting the place production.
	ExitPlace(c *PlaceContext)

	// ExitBraces is called when exiting the braces production.
	ExitBraces(c *BracesContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitRanges is called when exiting the ranges production.
	ExitRanges(c *RangesContext)
}
