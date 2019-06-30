package commands

import (
	"fmt"
	"strconv"
)

type PushCommand struct {
	code      Code
	length    int
	argString string
}

func (cmd PushCommand) GetOpCode() Code {
	return cmd.code
}

func (cmd PushCommand) Length() int {
	return cmd.length
}

func (cmd PushCommand) String() string {
	if cmd.argString == "" {
		return cmd.GetOpCode().String()
	}
	return fmt.Sprintf("%s %s", cmd.GetOpCode().String(), cmd.argString)
}

type PushByteCommand struct {
	PushCommand
	Arg byte
}

func NewPushByteCommand(arg byte) *PushByteCommand {
	return &PushByteCommand{PushCommand: PushCommand{code: PushByte, length: 2, argString: strconv.Itoa(int(arg))}, Arg: arg}
}

type PushIntCommand struct {
	PushCommand
	Arg int32
}

func NewPushIntCommand(arg int32) *PushIntCommand {
	return &PushIntCommand{PushCommand: PushCommand{code: PushInt, length: 5, argString: strconv.Itoa(int(arg))}, Arg: arg}
}

type PushLongCommand struct {
	PushCommand
	Arg int64
}

func NewPushLongCommand(arg int64) *PushLongCommand {
	return &PushLongCommand{PushCommand: PushCommand{code: PushLong, length: 9, argString: strconv.FormatInt(arg, 10)}, Arg: arg}
}

type PushReferenceCommand struct {
	PushCommand
	Arg int32
}

func NewPushReferenceCommand(arg int32) *PushReferenceCommand {
	return &PushReferenceCommand{PushCommand: PushCommand{code: PushReference, length: 5, argString: strconv.Itoa(int(arg))}, Arg: arg}
}
