package commands

import "fmt"

type ModBaseCommand struct {
	BaseCommand
	Left  Command
	Right Command
}

func (cmd ModBaseCommand) String() string {
	return fmt.Sprintf("%s\n%s\n%s", cmd.Left.String(), cmd.Right.String(), cmd.BaseCommand.String())
}

type ModByteCommand struct {
	ModBaseCommand
}

func NewModByteCommand(right, left Command) *ModByteCommand {
	return &ModByteCommand{ModBaseCommand{
		BaseCommand: BaseCommand{code: ModByte},
		Left:        left,
		Right:       right,
	}}
}

type ModIntCommand struct {
	ModBaseCommand
}

func NewModIntCommand(right, left Command) *ModIntCommand {
	return &ModIntCommand{ModBaseCommand{
		BaseCommand: BaseCommand{code: ModInt},
		Left:        left,
		Right:       right,
	}}
}

type ModLongCommand struct {
	ModBaseCommand
}

func NewModLongCommand(right, left Command) *ModLongCommand {
	return &ModLongCommand{ModBaseCommand{
		BaseCommand: BaseCommand{code: ModLong},
		Left:        left,
		Right:       right,
	}}
}
