package commands

import "fmt"

type StoreByteCommand struct {
	ByteCommand
	LeftRightCommands
}

func NewStoreByteCommand(arg byte, right, left Command) *StoreByteCommand {
	return &StoreByteCommand{
		ByteCommand:       *NewByteCommand(StoreByte, arg),
		LeftRightCommands: NewLeftRightCommands(left, right),
	}
}

func (cmd StoreByteCommand) String() string {
	return fmt.Sprintf("%s\n%s", cmd.LeftRightCommands.String(), cmd.ByteCommand.String())
}

type StoreIntCommand struct {
	Int32Command
	LeftRightCommands
}

func NewStoreIntCommand(arg int32, right, left Command) *StoreIntCommand {
	return &StoreIntCommand{
		Int32Command:      *NewInt32Command(StoreInt, arg),
		LeftRightCommands: NewLeftRightCommands(left, right),
	}
}

func (cmd StoreIntCommand) String() string {
	return fmt.Sprintf("%s\n%s", cmd.LeftRightCommands.String(), cmd.Int32Command.String())
}

type StoreLongCommand struct {
	Int64Command
	LeftRightCommands
}

func NewStoreLongCommand(arg int64, right, left Command) *StoreLongCommand {
	return &StoreLongCommand{
		Int64Command:      *NewInt64Command(StoreLong, arg),
		LeftRightCommands: NewLeftRightCommands(left, right),
	}
}

func (cmd StoreLongCommand) String() string {
	return fmt.Sprintf("%s\n%s", cmd.LeftRightCommands.String(), cmd.Int64Command.String())
}

type StoreReferenceCommand struct {
	Int32Command
	LeftRightCommands
}

func NewStoreReferenceCommand(arg int32, right, left Command) *StoreReferenceCommand {
	return &StoreReferenceCommand{
		Int32Command:      *NewInt32Command(StoreReference, arg),
		LeftRightCommands: NewLeftRightCommands(left, right),
	}
}

func (cmd StoreReferenceCommand) String() string {
	return fmt.Sprintf("%s\n%s", cmd.LeftRightCommands.String(), cmd.Int32Command.String())
}
