package lab2

import (
	"fmt"
	"github.com/artslob/pmv-go/parser"
)

type BlockClass int

const (
	DEFAULT BlockClass = iota
	START
	END
)

type Block struct {
	id    int
	text  string
	next  *Block
	level int
	class BlockClass
}

func (b *Block) String() string {
	label := ""
	switch b.class {
	case START:
		label = "START"
	case END:
		label = "END"
	default:
		label = b.text
	}
	return fmt.Sprintf("%2d [label=\"%s\"]\n", b.id, label)
}

type CFGListener struct {
	*parser.BaseLangListener
	currentId    int
	level        int
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
	block := &Block{id: s.nextId(), text: ctx.GetText(), level: s.level}
	if s.currentBlock.level == block.level && s.currentBlock.next == nil {
		s.currentBlock.next = block
	}
	s.currentBlock = block
}

func (s *CFGListener) ExitFuncDef(ctx *parser.FuncDefContext) {
	s.currentBlock.next = s.end
}