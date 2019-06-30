package operations

type Code byte

const (
	Nop  Code = iota
	Halt      = iota

	Neg = iota
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

	// push 1 byte to stack from command argument
	PushByte = iota
	// push 4 bytes to stack from command argument
	PushInt = iota
	// push 8 bytes to stack from command argument
	PushLong = iota
	// push 4 bytes to stack from command argument
	PushReference = iota

	// store 1 byte from stack to local variable with #index
	StoreByte = iota
	// store 4 bytes from stack to local variable with #index
	StoreInt = iota
	// store 8 bytes from stack to local variable with #index
	StoreLong = iota
	// store 4 bytes from stack to local variable with #index
	StoreReference = iota

	// fetch 1 byte from local variable with #index onto the stack
	FetchByte = iota
	// fetch 4 bytes from local variable with #index onto the stack
	FetchInt = iota
	// fetch 8 bytes from local variable with #index onto the stack
	FetchLong = iota
	// fetch 4 bytes from local variable with #index onto the stack
	FetchReference = iota
)
