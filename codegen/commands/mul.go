package commands

import "fmt"

type MulBaseCommand struct {
	BaseCommand
	Left  Command
	Right Command
}

func (cmd MulBaseCommand) String() string {
	return fmt.Sprintf("%s\n%s\n%s", cmd.Left.String(), cmd.Right.String(), cmd.BaseCommand.String())
}

type MulByteCommand struct {
	MulBaseCommand
}

func NewMulByteCommand(right, left Command) *MulByteCommand {
	return &MulByteCommand{MulBaseCommand{
		BaseCommand: BaseCommand{code: MulByte},
		Left:        left,
		Right:       right,
	}}
}

type MulIntCommand struct {
	MulBaseCommand
}

func NewMulIntCommand(right, left Command) *MulIntCommand {
	return &MulIntCommand{MulBaseCommand{
		BaseCommand: BaseCommand{code: MulInt},
		Left:        left,
		Right:       right,
	}}
}

type MulLongCommand struct {
	MulBaseCommand
}

func NewMulLongCommand(right, left Command) *MulLongCommand {
	return &MulLongCommand{MulBaseCommand{
		BaseCommand: BaseCommand{code: MulLong},
		Left:        left,
		Right:       right,
	}}
}
