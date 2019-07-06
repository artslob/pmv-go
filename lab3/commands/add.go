package commands

import "strconv"

type AddByteCommand struct {
	ByteCommand
}

func NewAddByteCommand(arg byte) *AddByteCommand {
	return &AddByteCommand{
		ByteCommand: ByteCommand{
			BaseCommand: BaseCommand{
				code:        AddByte,
				argAsString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type AddIntCommand struct {
	Int32Command
}

func NewAddIntCommand(arg int32) *AddIntCommand {
	return &AddIntCommand{
		Int32Command: Int32Command{
			BaseCommand: BaseCommand{
				code:        AddInt,
				argAsString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type AddLongCommand struct {
	Int64Command
}

func NewAddLongCommand(arg int64) *AddLongCommand {
	return &AddLongCommand{
		Int64Command: Int64Command{
			BaseCommand: BaseCommand{
				code:        AddLong,
				argAsString: strconv.FormatInt(arg, 10),
			},
			Arg: arg,
		},
	}
}
