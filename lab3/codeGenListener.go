package lab3

import (
	"github.com/artslob/pmv-go/parser"
)

type CodeGeneratorListener struct {
	*parser.BaseLangListener
}

func NewCodeGeneratorListener() *CodeGeneratorListener {
	return new(CodeGeneratorListener)
}
