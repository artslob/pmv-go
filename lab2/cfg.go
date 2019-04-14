package lab2

import (
	"github.com/artslob/pmv-go/parser"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId int
	start     Block
	end       Block
	blocks    BlockStack
	level     int
}

func NewCFGListener() *CFGListener {
	l := &CFGListener{}
	l.start = &SimpleBlock{id: -1, text: "START"}
	l.end = &SimpleBlock{id: -2, text: "END"}
	l.blocks.Push(l.start)
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) ExitExpression(ctx *parser.ExpressionContext) {
	b := SimpleBlock{id: s.nextId(), text: ctx.GetText()}
	s.blocks.Push(&b)
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Push(s.end)
	for s.blocks.Size() > 1 {
		last := s.blocks.Pop()
		second := s.blocks.Pop()
		second.append(last)
		s.blocks.Push(second)
	}
}
