package commands

import "fmt"

type Command interface {
	GetOpCode() Code
	Length() int
	fmt.Stringer
}

type BaseCommand struct {
	code      Code
	length    int
	argString string
}

func (cmd BaseCommand) GetOpCode() Code {
	return cmd.code
}

func (cmd BaseCommand) Length() int {
	return cmd.length
}

func (cmd BaseCommand) String() string {
	if cmd.argString == "" {
		return cmd.GetOpCode().String()
	}
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), cmd.argString)
}
