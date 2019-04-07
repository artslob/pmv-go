package lab2

import (
	"github.com/artslob/pmv-go/parser"
)

type CFGListener struct {
	*parser.BaseLangListener
	currentId    int
	currentBlock *Block
	start        *Block
	end          *Block
}

func NewCFGListener() *CFGListener {
	l := &CFGListener{}
	l.start = &Block{class: START, id: -1}
	l.end = &Block{class: END, id: -2}
	l.currentBlock = l.start
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}

func (s *CFGListener) EnterExpression(ctx *parser.ExpressionContext) {
	block := &Block{id: s.nextId(), text: ctx.GetText()}
	if s.currentBlock.next == nil {
		s.currentBlock.next = block
	}
	s.currentBlock = block
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.currentBlock.next = s.end
}
