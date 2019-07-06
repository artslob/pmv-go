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
	code        Code
	argAsString string
}

func (cmd BaseCommand) GetOpCode() Code {
	return cmd.code
}

func (cmd BaseCommand) Length() int {
	return 1
}

func (cmd BaseCommand) String() string {
	if cmd.argAsString == "" {
		return cmd.GetOpCode().String()
	}
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), cmd.argAsString)
}

type ByteCommand struct {
	BaseCommand
	Arg byte
}

func (cmd ByteCommand) Length() int {
	return 2
}

type Int32Command struct {
	BaseCommand
	Arg int32
}

func (cmd Int32Command) Length() int {
	return 5
}

type Int64Command struct {
	BaseCommand
	Arg int64
}

func (cmd Int64Command) Length() int {
	return 9
}
