// Code generated by "stringer -type Code"; DO NOT EDIT.

package commands

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Nop-0]
	_ = x[Halt-1]
	_ = x[NegByte-2]
	_ = x[NegInt-3]
	_ = x[NegLong-4]
	_ = x[Not-5]
	_ = x[Mul-6]
	_ = x[Div-7]
	_ = x[Mod-8]
	_ = x[Add-9]
	_ = x[Sub-10]
	_ = x[IfEqZero-11]
	_ = x[IfNotEqZero-12]
	_ = x[IfLessZero-13]
	_ = x[IfLessEqZero-14]
	_ = x[IfGtZero-15]
	_ = x[IfGtEqZero-16]
	_ = x[IfEq-17]
	_ = x[IfNotEq-18]
	_ = x[IfLess-19]
	_ = x[IfLessEq-20]
	_ = x[IfGt-21]
	_ = x[IfGtEq-22]
	_ = x[And-23]
	_ = x[Or-24]
	_ = x[Xor-25]
	_ = x[PushByte-26]
	_ = x[PushInt-27]
	_ = x[PushLong-28]
	_ = x[PushReference-29]
	_ = x[StoreByte-30]
	_ = x[StoreInt-31]
	_ = x[StoreLong-32]
	_ = x[StoreReference-33]
	_ = x[FetchByte-34]
	_ = x[FetchInt-35]
	_ = x[FetchLong-36]
	_ = x[FetchReference-37]
}

const _Code_name = "NopHaltNegByteNegIntNegLongNotMulDivModAddSubIfEqZeroIfNotEqZeroIfLessZeroIfLessEqZeroIfGtZeroIfGtEqZeroIfEqIfNotEqIfLessIfLessEqIfGtIfGtEqAndOrXorPushBytePushIntPushLongPushReferenceStoreByteStoreIntStoreLongStoreReferenceFetchByteFetchIntFetchLongFetchReference"

var _Code_index = [...]uint16{0, 3, 7, 14, 20, 27, 30, 33, 36, 39, 42, 45, 53, 64, 74, 86, 94, 104, 108, 115, 121, 129, 133, 139, 142, 144, 147, 155, 162, 170, 183, 192, 200, 209, 223, 232, 240, 249, 263}

func (i Code) String() string {
	if i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}