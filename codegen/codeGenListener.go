package codegen

import (
	"github.com/artslob/pmv-go/codegen/commands"
	"github.com/artslob/pmv-go/parser"
)

type CodeGeneratorListener struct {
	*parser.BaseLangListener
	CommandStack commands.Stack
}

func NewCodeGeneratorListener() *CodeGeneratorListener {
	return new(CodeGeneratorListener)
}
