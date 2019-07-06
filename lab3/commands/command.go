package commands

import (
	"fmt"
	"strconv"
)

type Command interface {
	GetOpCode() Code
	// Length returns size of command in bytes. Every command takes 1 byte minimum - operation code byte.
	// Then command can have optional argument, that specify for example index. So if argument is 1 byte, then
	// length is 2 bytes; if argument is int, then length is 5 bytes.
	Length() int
	fmt.Stringer
}

type BaseCommand struct {
	code Code
	// If command length > 1 (it has optional arguments, like index) then this field contains string
	// representation of this argument.
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

func NewByteCommand(code Code, arg byte) *ByteCommand {
	return &ByteCommand{
		BaseCommand: BaseCommand{
			code:        code,
			argAsString: strconv.Itoa(int(arg)),
		},
		Arg: arg,
	}
}

type Int32Command struct {
	BaseCommand
	Arg int32
}

func (cmd Int32Command) Length() int {
	return 5
}

func NewInt32Command(code Code, arg int32) *Int32Command {
	return &Int32Command{
		BaseCommand: BaseCommand{
			code:        code,
			argAsString: strconv.Itoa(int(arg)),
		},
		Arg: arg,
	}
}

type Int64Command struct {
	BaseCommand
	Arg int64
}

func (cmd Int64Command) Length() int {
	return 9
}

func NewInt64Command(code Code, arg int64) *Int64Command {
	return &Int64Command{
		BaseCommand: BaseCommand{
			code:        code,
			argAsString: strconv.FormatInt(arg, 10),
		},
		Arg: arg,
	}
}
