package commands

import (
	"strconv"
)

type PushByteCommand struct {
	BaseCommand
	Arg byte
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{BaseCommand: BaseCommand{code: PushByte, length: 2, argString: strconv.Itoa(int(arg))}, Arg: arg}
}

type PushIntCommand struct {
	BaseCommand
	Arg int32
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{BaseCommand: BaseCommand{code: PushInt, length: 5, argString: strconv.Itoa(int(arg))}, Arg: arg}
}

type PushLongCommand struct {
	BaseCommand
	Arg int64
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{BaseCommand: BaseCommand{code: PushLong, length: 9, argString: strconv.FormatInt(arg, 10)}, Arg: arg}
}

type PushReferenceCommand struct {
	BaseCommand
	Arg int32
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{BaseCommand: BaseCommand{code: PushReference, length: 5, argString: strconv.Itoa(int(arg))}, Arg: arg}
}
