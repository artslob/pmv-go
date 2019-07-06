package lab3

import (
	"github.com/artslob/pmv-go/lab3/commands"
	"github.com/artslob/pmv-go/parser"
	"strconv"
)

type CodeGeneratorListener struct {
	*parser.BaseLangListener
	CommandStack commands.Stack
}

func NewCodeGeneratorListener() *CodeGeneratorListener {
	return new(CodeGeneratorListener)
}

func (s *CodeGeneratorListener) ExitLiteralDec(ctx *parser.LiteralDecContext) {
	i, err := strconv.ParseInt(ctx.GetText(), 10, 32)
	if err != nil {
		panic(err)
	}
	s.CommandStack.Push(commands.NewPushIntCommand(int32(i)))
}

func (s *CodeGeneratorListener) ExitAddSub(ctx *parser.AddSubContext) {
	op := ctx.GetOp().GetText()
	switch op {
	case "+":
		s.CommandStack.Push(commands.NewAddIntCommand())
	case "-":
		s.CommandStack.Push(commands.NewSubIntCommand())
	default:
		panic("unknown operand at add sub command: " + op)
	}
}
