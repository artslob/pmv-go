package commands

import (
	"fmt"
	"github.com/artslob/pmv-go/codegen/types"
	"strconv"
)

type Command interface {
	GetOpCode() Code
	// Length returns size of command in bytes. Every command takes 1 byte minimum - operation code byte.
	// Then command can have optional argument, that specify for example index. So if argument is 1 byte, then
	// length is 2 bytes; if argument is int, then length is 5 bytes.
	Length() int
	Type() types.Type
	fmt.Stringer
}

type BaseCommand struct {
	code    Code
	cmdType types.Type
}

func (cmd BaseCommand) GetOpCode() Code {
	return cmd.code
}

func (cmd BaseCommand) Length() int {
	return 1
}

func (cmd BaseCommand) Type() types.Type {
	return cmd.cmdType
}

func (cmd BaseCommand) String() string {
	return cmd.GetOpCode().String()
}

func NewByteCommand(code Code, cmdType types.Type, arg byte) *ByteCommand {
	return &ByteCommand{BaseCommand: BaseCommand{code: code, cmdType: cmdType}, Arg: arg}
}

type ByteCommand struct {
	BaseCommand
	Arg byte
}

func (cmd ByteCommand) Length() int {
	return 2
}

func (cmd ByteCommand) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), strconv.Itoa(int(cmd.Arg)))
}

func NewInt32Command(code Code, cmdType types.Type, arg int32) *Int32Command {
	return &Int32Command{BaseCommand: BaseCommand{code: code, cmdType: cmdType}, Arg: arg}
}

type Int32Command struct {
	BaseCommand
	Arg int32
}

func (cmd Int32Command) Length() int {
	return 5
}

func (cmd Int32Command) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), strconv.Itoa(int(cmd.Arg)))
}

func NewInt64Command(code Code, cmdType types.Type, arg int64) *Int64Command {
	return &Int64Command{BaseCommand: BaseCommand{code: code, cmdType: cmdType}, Arg: arg}
}

type Int64Command struct {
	BaseCommand
	Arg int64
}

func (cmd Int64Command) Length() int {
	return 9
}

func (cmd Int64Command) String() string {
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), strconv.FormatInt(cmd.Arg, 10))
}

type LeftRightCommands struct {
	Left, Right Command
}

func NewLeftRightCommands(left Command, right Command) LeftRightCommands {
	return LeftRightCommands{Left: left, Right: right}
}

func (cmd LeftRightCommands) String() string {
	return fmt.Sprintf("%s\n%s", cmd.Left.String(), cmd.Right.String())
}

type TwoChildCommand struct {
	BaseCommand
	children LeftRightCommands
}

func NewTwoChildCommand(code Code, cmdType types.Type, left Command, right Command) TwoChildCommand {
	return TwoChildCommand{BaseCommand: BaseCommand{code: code, cmdType: cmdType}, children: NewLeftRightCommands(left, right)}
}

func (cmd TwoChildCommand) String() string {
	return fmt.Sprintf("%s\n%s", cmd.children.String(), cmd.BaseCommand.String())
}
