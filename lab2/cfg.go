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
	block := &Block{id: s.nextId(), text: ctx.GetText()}
	topBlock := s.blocks.Pop()
	if topBlock.next == nil {
		topBlock.next = block
	}
	s.blocks.Push(block)
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.blocks.Pop().next = s.end
}
