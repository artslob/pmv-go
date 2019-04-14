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
	blocks    block.BlockStack
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
	b := block.SimpleBlock{Id: s.nextId(), Text: ctx.GetText()}
	s.blocks.Push(&b)
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Push(s.end)
	for s.blocks.Size() > 1 {
		last := s.blocks.Pop()
		second := s.blocks.Pop()
		second.Append(last)
		s.blocks.Push(second)
	}
}
