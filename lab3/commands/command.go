package commands

import "fmt"

type Command interface {
	GetOpCode() Code
	// Length returns size of command in bytes. Every command takes 1 byte minimum - operation code byte.
	// Then command can have optional argument, that specify for example index. So if argument is 1 byte, then
	// length is 2 bytes; if argument is int, then length is 5 bytes.
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
