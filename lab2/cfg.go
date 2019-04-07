package lab2

import (
	"github.com/artslob/pmv-go/parser"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId int
	start     *Block
	end       *Block
	blocks    BlockStack
	level     int
}

func NewCFGListener() *CFGListener {
	l := &CFGListener{}
	l.start = &Block{class: START, id: -1}
	l.end = &Block{class: END, id: -2}
	l.blocks.Push(l.start)
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) ExitExpression(ctx *parser.ExpressionContext) {
	block := &Block{id: s.nextId(), text: ctx.GetText(), level: s.level}
	topBlock := s.blocks.Peek()
	defer s.blocks.Push(block)
	if topBlock.level != s.level {
		return
	}
	if topBlock.next == nil {
		topBlock.next = block
	} else if topBlock.class == IF {
		topBlock.next.next = block
		topBlock.branch = block
	}
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Pop().next = s.end
}

func (s *CFGListener) ExitIfExpr(ctx *parser.IfExprContext) {
	ifExpr := &Block{id: s.nextId(), text: ctx.Expr().GetText(), class: IF, level: s.level}
	s.blocks.Push(ifExpr)
}

func (s *CFGListener) EnterIfThen(ctx *parser.IfThenContext) {
	s.level++
}

func (s *CFGListener) ExitIfThen(ctx *parser.IfThenContext) {
	s.level--
}

func (s *CFGListener) ExitIf(ctx *parser.IfContext) {
	withElse := ctx.IfElse() != nil
	// on stack: n: else, n-1: then, n-2: if
	if withElse == true {
		// TODO
		return
	}
	// on stack: n: then, n-1: if
	then := s.blocks.Pop()
	ifExpr := s.blocks.Pop()
	ifExpr.next = then
	s.blocks.Peek().next = ifExpr
	s.blocks.Push(ifExpr)
}
