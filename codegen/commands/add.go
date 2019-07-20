package commands

import "fmt"

type AddBaseCommand struct {
	BaseCommand
	Left  Command
	Right Command
}

func (cmd AddBaseCommand) String() string {
	return fmt.Sprintf("%s\n%s\n%s", cmd.Left.String(), cmd.Right.String(), cmd.BaseCommand.String())
}

type AddByteCommand struct {
	AddBaseCommand
}

func NewAddByteCommand(right, left Command) *AddByteCommand {
	return &AddByteCommand{AddBaseCommand{
		BaseCommand: BaseCommand{code: AddByte},
		Left:        left,
		Right:       right,
	}}
}

type AddIntCommand struct {
	AddBaseCommand
}

func NewAddIntCommand(right, left Command) *AddIntCommand {
	return &AddIntCommand{AddBaseCommand{
		BaseCommand: BaseCommand{code: AddInt},
		Left:        left,
		Right:       right,
	}}
}

type AddLongCommand struct {
	AddBaseCommand
}

func NewAddLongCommand(right, left Command) *AddLongCommand {
	return &AddLongCommand{AddBaseCommand{
		BaseCommand: BaseCommand{code: AddLong},
		Left:        left,
		Right:       right,
	}}
}
