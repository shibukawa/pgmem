package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regc_wc_isalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isalnum[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	if v6 == int32(1) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v23 = int32(0)
			return v23
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_regc_wc_isalnum[1]))))
			return base.B2i32(v11&int32(3) != int32(0))
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
		v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l0, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v19
			return v23
		}
	}
}
func F_regc_wc_isdigit(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isdigit[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
	if v5 == int32(1) {
		return base.B2i32(base.Ui32(l0-int32(48)) < base.Ui32(int32(10)))
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
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
func F_regc_wc_isprint(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isprint[0]))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
	if v5 == int32(1) {
		return base.B2i32(base.Ui32(l0-int32(32)) < base.Ui32(int32(95)))
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
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
