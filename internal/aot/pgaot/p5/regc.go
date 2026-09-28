package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regc_wc_isgraph(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isgraph[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
	if v5 == int32(1) {
		return base.B2i32(base.Ui32(l0-int32(33)) < base.Ui32(int32(94)))
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
		v15 = m.T0[v14].(func(*base.Module, int32, int32) int32)(m, l0, v4)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
func F_regc_wc_isupper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isupper[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
	if v5 == int32(1) {
		return base.B2i32(base.Ui32(l0-int32(65)) < base.Ui32(int32(26)))
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
		v15 = m.T0[v14].(func(*base.Module, int32, int32) int32)(m, l0, v4)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	}
}
