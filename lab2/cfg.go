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
	l.start = &Block{id: -1}
	l.end = &Block{id: -2}
	l.blocks.Push(l.start)
	return l
}

func (s *CFGListener) nextId() int {
	s.currentId++
	return s.currentId
}
