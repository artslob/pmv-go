package operations

type Code byte

const (
	Nop  Code = iota
	Halt      = iota

	// negate byte on the stack
	NegByte = iota
	// negate int on the stack
	NegInt = iota
	// negate long on the stack
	NegLong = iota

	Not = iota

	Mul = iota
	Div = iota
	Mod = iota

	Add = iota
	Sub = iota

	IfEqZero     = iota
	IfNotEqZero  = iota
	IfLessZero   = iota
	IfLessEqZero = iota
	IfGtZero     = iota
	IfGtEqZero   = iota

	IfEq     = iota
	IfNotEq  = iota
	IfLess   = iota
	IfLessEq = iota
	IfGt     = iota
	IfGtEq   = iota

	And = iota
	Or  = iota
	Xor = iota

	// push byte to stack from command argument
	PushByte = iota
	// push int to stack from command argument
	PushInt = iota
	// push long to stack from command argument
	PushLong = iota
	// push reference to stack from command argument
	PushReference = iota

	// store byte from stack to local variable with #index
	StoreByte = iota
	// store int from stack to local variable with #index
	StoreInt = iota
	// store long from stack to local variable with #index
	StoreLong = iota
	// store reference from stack to local variable with #index
	StoreReference = iota

	// fetch byte from local variable with #index onto the stack
	FetchByte = iota
	// fetch int from local variable with #index onto the stack
	FetchInt = iota
	// fetch long from local variable with #index onto the stack
	FetchLong = iota
	// fetch reference from local variable with #index onto the stack
	FetchReference = iota
)
