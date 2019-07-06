package commands

type Code byte

const (
	Nop Code = iota
	Halt

	// negate byte on the stack
	NegByte
	// negate int on the stack
	NegInt
	// negate long on the stack
	NegLong

	Not

	Mul
	Div
	Mod

	AddByte
	AddInt
	AddLong

	SubByte
	SubInt
	SubLong

	IfEqZero
	IfNotEqZero
	IfLessZero
	IfLessEqZero
	IfGtZero
	IfGtEqZero

	IfEq
	IfNotEq
	IfLess
	IfLessEq
	IfGt
	IfGtEq

	And
	Or
	Xor

	// push byte to stack from command argument
	PushByte
	// push int to stack from command argument
	PushInt
	// push long to stack from command argument
	PushLong
	// push reference to stack from command argument
	PushReference

	// store byte from stack to local variable with #index
	StoreByte
	// store int from stack to local variable with #index
	StoreInt
	// store long from stack to local variable with #index
	StoreLong
	// store reference from stack to local variable with #index
	StoreReference

	// fetch byte from local variable with #index onto the stack
	FetchByte
	// fetch int from local variable with #index onto the stack
	FetchInt
	// fetch long from local variable with #index onto the stack
	FetchLong
	// fetch reference from local variable with #index onto the stack
	FetchReference
)
