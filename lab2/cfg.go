package lab2

import (
	"github.com/artslob/pmv-go/lab2/block"
	"github.com/artslob/pmv-go/parser"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId int
	start     block.Block
	end       block.Block
	blocks    block.Stack
	level     int
}

func NewCFGListener() *CFGListener {
	l := &CFGListener{}
	l.start = &block.SimpleBlock{Id: -1, Text: "START"}
	l.end = &block.SimpleBlock{Id: -2, Text: "END"}
	l.blocks.Push(l.start)
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) ExitExpression(ctx *parser.ExpressionContext) {
	s.blocks.Push(&block.SimpleBlock{Id: s.nextId(), Text: ctx.GetText()})
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Push(s.end)
	for s.blocks.Size() > 1 {
		last := s.blocks.Pop()
		second := s.blocks.Pop()
		second.Append(last)
		s.blocks.Push(second.GetAddress())
	}
}

func (s *CFGListener) ExitIfExpr(ctx *parser.IfExprContext) {
	s.blocks.Push(&block.IfExpr{Id: s.nextId(), Text: ctx.Expr().GetText()})
}

func (s *CFGListener) ExitIf(ctx *parser.IfContext) {
	var else_ block.Block
	if ctx.IfElse() == nil {
		// stack with only then:   [n]: then, [n-1]: if
		else_ = &block.SimpleBlock{Id: s.nextId()}

	} else {
		// with else: [n]: else, [n-1]: then, [n-2]: if
		else_ = s.blocks.Pop()
	}
	then := s.blocks.Pop()
	expr := s.blocks.Pop()
	end := &block.SimpleBlock{Id: s.nextId()}
	expr.Append(then)
	expr.SetBranch(else_)
	then.Append(end)
	else_.Append(end)
	ifBlock := block.IfBlock{
		Id:    s.nextId(),
		Expr:  expr,
		Then:  then,
		Else_: else_,
		End:   end,
	}
	s.blocks.Push(&ifBlock)
}
