package commands

import (
	"strconv"
)

type PushByteCommand struct {
	ByteCommand
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{
		ByteCommand: ByteCommand{
			BaseCommand: BaseCommand{
				code:      PushByte,
				argString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type PushIntCommand struct {
	Int32Command
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{
		Int32Command: Int32Command{
			BaseCommand: BaseCommand{
				code:      PushInt,
				argString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}

type PushLongCommand struct {
	Int64Command
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{
		Int64Command: Int64Command{
			BaseCommand: BaseCommand{
				code:      PushLong,
				argString: strconv.FormatInt(arg, 10),
			},
			Arg: arg,
		},
	}
}

type PushReferenceCommand struct {
	Int32Command
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{
		Int32Command: Int32Command{
			BaseCommand: BaseCommand{
				code:      PushReference,
				argString: strconv.Itoa(int(arg)),
			},
			Arg: arg,
		},
	}
}
