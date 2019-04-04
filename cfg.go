package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/parser"
	"strings"
)

func parseInputToCFG(input string) *Block {
	p := getParser(input)
	listener := newCFGListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener.head
}

func printCFG(head *Block, builder *strings.Builder) {
	if builder.Len() == 0 {
		builder.WriteString("digraph G {\n")
		defer func() { builder.WriteString("}\n") }()
	}
	builder.WriteString(fmt.Sprintf("%d[label=\"%s\"]\n", head.id, head.text))
	if head.next != nil {
		builder.WriteString(fmt.Sprintf("%d->%d\n", head.id, head.next.id))
		printCFG(head.next, builder)
	}
	if head.branch != nil {
		builder.WriteString(fmt.Sprintf("%d->%d [style=dotted]\n", head.id, head.branch.id))
		printCFG(head.branch, builder)
	}
}

type BlockType int

const (
	DEFAULT BlockType = iota
	IF_EXPR
)

type Block struct {
	next, branch *Block
	id           int
	text         string
	blkType      BlockType
	level        int
}

type CFGListener struct {
	*parser.BaseLangListener
	head         *Block
	block        *Block
	currentId    int
	currentLevel int
	blocks       []*Block
}

func newCFGListener() *CFGListener {
	return &CFGListener{}
}

func (s *CFGListener) increaseId() {
	s.currentId++
}

func (s *CFGListener) pushBlock(block *Block) {
	s.blocks = append(s.blocks, block)
}

func (s *CFGListener) popBlock() *Block {
	res := s.blocks[len(s.blocks)-1]
	s.blocks = s.blocks[:len(s.blocks)-1]
	return res
}

func (s *CFGListener) removeBlock(i int) {
	s.blocks = append(s.blocks[:i], s.blocks[i+1:]...)
}

func (s *CFGListener) lastBlock() *Block {
	if len(s.blocks) == 0 {
		return nil
	}
	return s.blocks[len(s.blocks)-1]
}

func (s *CFGListener) emptyBlocks() bool {
	return len(s.blocks) == 0
}

func (s *CFGListener) currentBlock(block *Block) {
	if s.head == nil {
		s.head = block
	}
	s.block = block
}

func (s *CFGListener) EnterIfExpr(ctx *parser.IfExprContext) {
	s.increaseId()
	text := ctx.Expr().GetText()
	block := &Block{id: s.currentId, text: text, level: s.currentLevel, blkType: IF_EXPR}
	if s.lastBlock() != nil {
		s.lastBlock().next = block
		s.pushBlock(block)
	}
	fmt.Println(text)
	s.currentBlock(block)
}

func (s *CFGListener) EnterIfThen(ctx *parser.IfThenContext) {
	s.currentLevel++
}

func (s *CFGListener) ExitIfThen(ctx *parser.IfThenContext) {
	s.currentLevel--
}

func (s *CFGListener) EnterExpression(ctx *parser.ExpressionContext) {
	s.increaseId()
	block := &Block{id: s.currentId, text: ctx.GetText(), level: s.currentLevel}
	fmt.Print("expr ", s.currentId, " ")
	fmt.Println(ctx.GetText())
	s.currentBlock(block)
	if s.lastBlock() == nil {
		s.pushBlock(block)
		return
	}
	for i := len(s.blocks) - 1; i >= 0; i-- {
		lastBlock := s.blocks[i]
		if lastBlock.level == s.currentLevel {
			s.removeBlock(i)
		}
		if lastBlock.next == nil {
			lastBlock.next = block
		}
		if lastBlock.blkType == IF_EXPR && lastBlock.branch == nil && lastBlock.level == s.currentLevel {
			lastBlock.branch = block
		}
	}
	s.pushBlock(block)
}
