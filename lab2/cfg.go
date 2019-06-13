package lab2

import (
	"github.com/artslob/pmv-go/lab2/blocks"
	"github.com/artslob/pmv-go/parser"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId int
	start     blocks.Block
	end       blocks.Block
	blocks    blocks.Stack
	level     int
}

func NewCFGListener() *CFGListener {
	l := &CFGListener{}
	l.start = &blocks.DefaultBlock{Id: -1, Text: "START"}
	l.end = &blocks.DefaultBlock{Id: -2, Text: "END"}
	l.blocks.Push(l.start)
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) ExitExpression(ctx *parser.ExpressionContext) {
	s.blocks.Push(&blocks.DefaultBlock{Id: s.nextId(), Text: ctx.GetText()})
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Push(s.end)
	for s.blocks.Size() > 1 {
		last := s.blocks.Pop()
		second := s.blocks.Pop()
		second.SetNext(last)
		s.blocks.Push(second)
	}
}

func (s *CFGListener) ExitIfExpr(ctx *parser.IfExprContext) {
	s.blocks.Push(&blocks.BranchBlock{
		DefaultBlock: blocks.DefaultBlock{
			Id:   s.nextId(),
			Text: ctx.Expr().GetText(),
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
	expr.SetNext(then)
	expr.SetBranch(else_)
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
