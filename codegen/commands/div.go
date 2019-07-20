package commands

import "fmt"

type DivBaseCommand struct {
	BaseCommand
	Left  Command
	Right Command
}

func (cmd DivBaseCommand) String() string {
	return fmt.Sprintf("%s\n%s\n%s", cmd.Left.String(), cmd.Right.String(), cmd.BaseCommand.String())
}

type DivByteCommand struct {
	DivBaseCommand
}

func NewDivByteCommand(right, left Command) *DivByteCommand {
	return &DivByteCommand{DivBaseCommand{
		BaseCommand: BaseCommand{code: DivByte},
		Left:        left,
		Right:       right,
	}}
}

type DivIntCommand struct {
	DivBaseCommand
}

func NewDivIntCommand(right, left Command) *DivIntCommand {
	return &DivIntCommand{DivBaseCommand{
		BaseCommand: BaseCommand{code: DivInt},
		Left:        left,
		Right:       right,
	}}
}

type DivLongCommand struct {
	DivBaseCommand
}

func NewDivLongCommand(right, left Command) *DivLongCommand {
	return &DivLongCommand{DivBaseCommand{
		BaseCommand: BaseCommand{code: DivLong},
		Left:        left,
		Right:       right,
	}}
}
