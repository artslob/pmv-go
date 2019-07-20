package commands

import "fmt"

type SubBaseCommand struct {
	BaseCommand
	Left  Command
	Right Command
}

func (cmd SubBaseCommand) String() string {
	return fmt.Sprintf("%s\n%s\n%s", cmd.Left.String(), cmd.Right.String(), cmd.BaseCommand.String())
}

type SubByteCommand struct {
	SubBaseCommand
}

func NewSubByteCommand(right, left Command) *SubByteCommand {
	return &SubByteCommand{SubBaseCommand{
		BaseCommand: BaseCommand{code: SubByte},
		Left:        left,
		Right:       right,
	}}
}

type SubIntCommand struct {
	SubBaseCommand
}

func NewSubIntCommand(right, left Command) *SubIntCommand {
	return &SubIntCommand{SubBaseCommand{
		BaseCommand: BaseCommand{code: SubInt},
		Left:        left,
		Right:       right,
	}}
}

type SubLongCommand struct {
	SubBaseCommand
}

func NewSubLongCommand(right, left Command) *SubLongCommand {
	return &SubLongCommand{SubBaseCommand{
		BaseCommand: BaseCommand{code: SubLong},
		Left:        left,
		Right:       right,
	}}
}
