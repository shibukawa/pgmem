package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_assign_stats_fetch_consistency(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_stats_fetch_consistency[0]))
	if l0 != v4 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_assign_stats_fetch_consistency[1])) = uint8(v7)
	} else {
	}
	return
}
func F_stats_check_required_arg(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l2<<(uint(int32(4))%32))+32)))
	if v12 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1+l2<<(uint(int32(3))%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v25
				F_errmsg(m, int32(_a_F_stats_check_required_arg_0), v7)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_stats_check_required_arg_1), int32(61), int32(_a_F_stats_check_required_arg_2))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
