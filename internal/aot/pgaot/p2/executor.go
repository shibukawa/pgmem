package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecutorFinish(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _consts[439]))
	if v4 != 0 {
		m.T0[v4].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	} else {
		F_standard_ExecutorFinish(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_executor_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	if l0 == int32(0) {
		return
	} else {
		if l1 < int32(0) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v7 == int32(0) {
				return
			} else {
				v10 = F_pg_mbstrlen_with_len(m, v7, l1)
				mBase = m.M
				v11 = m.ExcPending
				if v11 != 0 {
					return
				} else {
					v14 = F_errposition(m, v10+int32(1))
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
