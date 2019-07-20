package cfg

import (
	"github.com/artslob/pmv-go/cfg/blocks"
	"github.com/artslob/pmv-go/codegen/commands"
	"github.com/artslob/pmv-go/parser"
	"github.com/artslob/pmv-go/utils"
	"strconv"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId      int
	blocks         blocks.Stack
	loopId         int
	loopIdStack    utils.IntStack
	breaks         map[int][]blocks.Block
	functionBlocks blocks.Stack
	calls          map[string]struct{}
	CommandStack   commands.Stack
}

func NewCFGListener() *CFGListener {
	return &CFGListener{breaks: map[int][]blocks.Block{}}
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) nextLoopId() int {
	s.loopId++
	return s.loopId
}

func (s *CFGListener) getCalls() map[string]struct{} {
	result := s.calls
	s.calls = nil
	return result
}

func (s *CFGListener) ExitLiteralDec(ctx *parser.LiteralDecContext) {
	i, err := strconv.ParseInt(ctx.GetText(), 10, 32)
	if err != nil {
		panic(err)
	}
	s.CommandStack.Push(commands.NewPushIntCommand(int32(i)))
}

func (s *CFGListener) ExitAddSub(ctx *parser.AddSubContext) {
	if s.CommandStack.Size() < 2 {
		return // TODO remove
	}
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		s.CommandStack.Push(commands.NewAddIntCommand(s.CommandStack.Pop(), s.CommandStack.Pop()))
	case "-":
		s.CommandStack.Push(commands.NewSubIntCommand(s.CommandStack.Pop(), s.CommandStack.Pop()))
	default:
		panic("unknown operand at add sub command: " + op)
	}
}

func (s *CFGListener) ExitMulDivMod(ctx *parser.MulDivModContext) {
	if s.CommandStack.Size() < 2 {
		return // TODO remove
	}
	op := ctx.GetOp().GetText()
	switch op {
	case "*":
		s.CommandStack.Push(commands.NewMulIntCommand(s.CommandStack.Pop(), s.CommandStack.Pop()))
	case "/":
		s.CommandStack.Push(commands.NewDivIntCommand(s.CommandStack.Pop(), s.CommandStack.Pop()))
	case "%":
		s.CommandStack.Push(commands.NewModIntCommand(s.CommandStack.Pop(), s.CommandStack.Pop()))
	default:
		panic("unknown operand at add sub command: " + op)
	}
}

func (s *CFGListener) ExitExpression(ctx *parser.ExpressionContext) {
	s.blocks.Push(&blocks.DefaultBlock{Id: s.nextId(), Text: ctx.GetText(), FunctionCalls: s.getCalls()})
}

func (s *CFGListener) ExitFuncSignature(ctx *parser.FuncSignatureContext) {
	s.blocks.Push(blocks.NewFunction(s.nextId(), ctx.IDENTIFIER().GetText()))
}

func (s *CFGListener) ExitCall(ctx *parser.CallContext) {
	if s.calls == nil {
		s.calls = map[string]struct{}{}
	}
	s.calls[ctx.Expr(0).GetText()] = struct{}{}
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Push(&blocks.DefaultBlock{Id: s.nextId(), Text: "END"})
	for s.blocks.Size() > 1 {
		last := s.blocks.Pop()
		second := s.blocks.Pop()
		second.SetNext(last)
		s.blocks.Push(second)
	}
	s.functionBlocks.Push(s.blocks.Pop())
}

func (s *CFGListener) ExitIfExpr(ctx *parser.IfExprContext) {
	s.blocks.Push(&blocks.BranchBlock{
		DefaultBlock: blocks.DefaultBlock{
			Id:            s.nextId(),
			Text:          ctx.Expr().GetText(),
			FunctionCalls: s.getCalls(),
		},
	})
}

func (s *CFGListener) ExitIf(ctx *parser.IfContext) {
	var else_ blocks.Block
	if ctx.IfElse() == nil {
		// stack with only then:   [n]: then, [n-1]: if
		else_ = blocks.NewEmptyBlock(s.nextId())
	} else {
		// with else: [n]: else, [n-1]: then, [n-2]: if
		else_ = s.blocks.Pop()
	}
	then := s.blocks.Pop()
	expr := s.blocks.Pop()
	end := blocks.NewEmptyBlock(s.nextId())
	expr.SetBranch(then)
	expr.SetNext(else_)
	then.SetNext(end)
	else_.SetNext(end)
	ifBlock := blocks.IfBlock{
		DefaultBlock: blocks.DefaultBlock{Id: s.nextId()},
		Expr:         expr,
		Then:         then,
		Else_:        else_,
		End:          end,
	}
	s.blocks.Push(&ifBlock)
}

func (s *CFGListener) ExitBlockBody(ctx *parser.BlockBodyContext) {
	if ctx.AllStatement() == nil || len(ctx.AllStatement()) == 0 {
		s.blocks.Push(blocks.NewEmptyBlock(s.nextId()))
		return
	}
	group := blocks.Group{}
	for range ctx.AllStatement() {
		group.AddBlock(s.blocks.Pop())
	}
	s.blocks.Push(&group)
}

func (s *CFGListener) ExitWhileBody(ctx *parser.WhileBodyContext) {
	if ctx.AllStatement() == nil || len(ctx.AllStatement()) == 0 {
		s.blocks.Push(blocks.NewEmptyBlock(s.nextId()))
		return
	}
	group := blocks.Group{}
	for range ctx.AllStatement() {
		group.AddBlock(s.blocks.Pop())
	}
	s.blocks.Push(&group)
}

func (s *CFGListener) ExitWhileExpr(ctx *parser.WhileExprContext) {
	s.blocks.Push(&blocks.BranchBlock{
		DefaultBlock: blocks.DefaultBlock{
			Id:            s.nextId(),
			Text:          ctx.Expr().GetText(),
			FunctionCalls: s.getCalls(),
		},
	})
}

func (s *CFGListener) ExitLoop(ctx *parser.LoopContext) {
	currentLoopId := s.loopIdStack.Pop()
	body := s.blocks.Pop()
	expr := s.blocks.Pop()
	expr.SetBranch(body)
	body.SetNext(expr)

	while := &blocks.While{
		DefaultBlock: blocks.DefaultBlock{Id: s.nextId()},
		Expression:   expr,
		Body:         body,
		Breaks:       s.breaks[currentLoopId],
	}
	s.blocks.Push(while)
	delete(s.breaks, currentLoopId)
}

func (s *CFGListener) ExitUntilExpr(ctx *parser.UntilExprContext) {
	s.blocks.Push(&blocks.BranchBlock{
		DefaultBlock: blocks.DefaultBlock{
			Id:            s.nextId(),
			Text:          ctx.Expr().GetText(),
			FunctionCalls: s.getCalls(),
		},
	})
}

func (s *CFGListener) ExitRepeat(ctx *parser.RepeatContext) {
	currentLoopId := s.loopIdStack.Pop()
	expression := s.blocks.Pop()
	statement := s.blocks.Pop()
	statement.SetNext(expression)
	expression.SetBranch(statement)

	until := &blocks.Until{
		DefaultBlock: blocks.DefaultBlock{Id: s.nextId()},
		Statement:    statement,
		Expression:   expression,
		Breaks:       s.breaks[currentLoopId],
	}
	s.blocks.Push(until)
	delete(s.breaks, currentLoopId)
}

func (s *CFGListener) ExitBreak(ctx *parser.BreakContext) {
	if s.loopIdStack.Empty() {
		panic("encounter break while being outside of loop")
	}
	breakBlock := &blocks.DefaultBlock{Id: s.nextId(), Text: ctx.GetText()}
	s.blocks.Push(breakBlock)
	currentLoopId := *s.loopIdStack.Peek()
	s.breaks[currentLoopId] = append(s.breaks[currentLoopId], breakBlock)
}

func (s *CFGListener) EnterLoop(ctx *parser.LoopContext) {
	s.loopIdStack.Push(s.nextLoopId())
}

func (s *CFGListener) EnterRepeat(ctx *parser.RepeatContext) {
	s.loopIdStack.Push(s.nextLoopId())
}
