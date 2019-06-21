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

// EnterIf is called when production if is entered.
func (s *BaseLangListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseLangListener) ExitIf(ctx *IfContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseLangListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseLangListener) ExitLoop(ctx *LoopContext) {}

// EnterRepeat is called when production repeat is entered.
func (s *BaseLangListener) EnterRepeat(ctx *RepeatContext) {}

// ExitRepeat is called when production repeat is exited.
func (s *BaseLangListener) ExitRepeat(ctx *RepeatContext) {}

// EnterBreak is called when production break is entered.
func (s *BaseLangListener) EnterBreak(ctx *BreakContext) {}

// ExitBreak is called when production break is exited.
func (s *BaseLangListener) ExitBreak(ctx *BreakContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseLangListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseLangListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseLangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseLangListener) ExitBlock(ctx *BlockContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseLangListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseLangListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterIfThen is called when production ifThen is entered.
func (s *BaseLangListener) EnterIfThen(ctx *IfThenContext) {}

// ExitIfThen is called when production ifThen is exited.
func (s *BaseLangListener) ExitIfThen(ctx *IfThenContext) {}

// EnterIfElse is called when production ifElse is entered.
func (s *BaseLangListener) EnterIfElse(ctx *IfElseContext) {}

// ExitIfElse is called when production ifElse is exited.
func (s *BaseLangListener) ExitIfElse(ctx *IfElseContext) {}

// EnterWhileExpr is called when production whileExpr is entered.
func (s *BaseLangListener) EnterWhileExpr(ctx *WhileExprContext) {}

// ExitWhileExpr is called when production whileExpr is exited.
func (s *BaseLangListener) ExitWhileExpr(ctx *WhileExprContext) {}

// EnterWhileBody is called when production whileBody is entered.
func (s *BaseLangListener) EnterWhileBody(ctx *WhileBodyContext) {}

// ExitWhileBody is called when production whileBody is exited.
func (s *BaseLangListener) ExitWhileBody(ctx *WhileBodyContext) {}

// EnterUntilStatement is called when production untilStatement is entered.
func (s *BaseLangListener) EnterUntilStatement(ctx *UntilStatementContext) {}

// ExitUntilStatement is called when production untilStatement is exited.
func (s *BaseLangListener) ExitUntilStatement(ctx *UntilStatementContext) {}

// EnterUntilExpr is called when production untilExpr is entered.
func (s *BaseLangListener) EnterUntilExpr(ctx *UntilExprContext) {}

// ExitUntilExpr is called when production untilExpr is exited.
func (s *BaseLangListener) ExitUntilExpr(ctx *UntilExprContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseLangListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseLangListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterMulDivMod is called when production mulDivMod is entered.
func (s *BaseLangListener) EnterMulDivMod(ctx *MulDivModContext) {}

// ExitMulDivMod is called when production mulDivMod is exited.
func (s *BaseLangListener) ExitMulDivMod(ctx *MulDivModContext) {}

// EnterOr is called when production or is entered.
func (s *BaseLangListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseLangListener) ExitOr(ctx *OrContext) {}

// EnterAndLogical is called when production andLogical is entered.
func (s *BaseLangListener) EnterAndLogical(ctx *AndLogicalContext) {}

// ExitAndLogical is called when production andLogical is exited.
func (s *BaseLangListener) ExitAndLogical(ctx *AndLogicalContext) {}

// EnterShift is called when production shift is entered.
func (s *BaseLangListener) EnterShift(ctx *ShiftContext) {}

// ExitShift is called when production shift is exited.
func (s *BaseLangListener) ExitShift(ctx *ShiftContext) {}

// EnterCompareArrow is called when production compareArrow is entered.
func (s *BaseLangListener) EnterCompareArrow(ctx *CompareArrowContext) {}

// ExitCompareArrow is called when production compareArrow is exited.
func (s *BaseLangListener) ExitCompareArrow(ctx *CompareArrowContext) {}

// EnterAddSub is called when production addSub is entered.
func (s *BaseLangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production addSub is exited.
func (s *BaseLangListener) ExitAddSub(ctx *AddSubContext) {}

// EnterCompareEqual is called when production compareEqual is entered.
func (s *BaseLangListener) EnterCompareEqual(ctx *CompareEqualContext) {}

// ExitCompareEqual is called when production compareEqual is exited.
func (s *BaseLangListener) ExitCompareEqual(ctx *CompareEqualContext) {}

// EnterUnary is called when production unary is entered.
func (s *BaseLangListener) EnterUnary(ctx *UnaryContext) {}

// ExitUnary is called when production unary is exited.
func (s *BaseLangListener) ExitUnary(ctx *UnaryContext) {}

// EnterBraces is called when production braces is entered.
func (s *BaseLangListener) EnterBraces(ctx *BracesContext) {}

// ExitBraces is called when production braces is exited.
func (s *BaseLangListener) ExitBraces(ctx *BracesContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseLangListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseLangListener) ExitLiteral(ctx *LiteralContext) {}

// EnterCall is called when production call is entered.
func (s *BaseLangListener) EnterCall(ctx *CallContext) {}

// ExitCall is called when production call is exited.
func (s *BaseLangListener) ExitCall(ctx *CallContext) {}

// EnterOrLogical is called when production orLogical is entered.
func (s *BaseLangListener) EnterOrLogical(ctx *OrLogicalContext) {}

// ExitOrLogical is called when production orLogical is exited.
func (s *BaseLangListener) ExitOrLogical(ctx *OrLogicalContext) {}

// EnterSlice is called when production slice is entered.
func (s *BaseLangListener) EnterSlice(ctx *SliceContext) {}

// ExitSlice is called when production slice is exited.
func (s *BaseLangListener) ExitSlice(ctx *SliceContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseLangListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseLangListener) ExitAnd(ctx *AndContext) {}

// EnterXor is called when production xor is entered.
func (s *BaseLangListener) EnterXor(ctx *XorContext) {}

// ExitXor is called when production xor is exited.
func (s *BaseLangListener) ExitXor(ctx *XorContext) {}

// EnterPlace is called when production place is entered.
func (s *BaseLangListener) EnterPlace(ctx *PlaceContext) {}

// ExitPlace is called when production place is exited.
func (s *BaseLangListener) ExitPlace(ctx *PlaceContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseLangListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseLangListener) ExitAssign(ctx *AssignContext) {}

// EnterRanges is called when production ranges is entered.
func (s *BaseLangListener) EnterRanges(ctx *RangesContext) {}

// ExitRanges is called when production ranges is exited.
func (s *BaseLangListener) ExitRanges(ctx *RangesContext) {}
