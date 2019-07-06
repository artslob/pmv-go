package lab3

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/artslob/pmv-go/lab1"
	"strings"
)

func GenerateCode(input string) *CodeGeneratorListener {
	p := lab1.GetParser(input)
	listener := NewCodeGeneratorListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Source())
	return listener
}

func GetBytecodeString(listener CodeGeneratorListener) string {
	var builder strings.Builder
	for _, cmd := range listener.CommandStack {
		builder.WriteString(cmd.String())
		builder.WriteString("\n")
	}
	return builder.String()
}
