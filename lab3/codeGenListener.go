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
