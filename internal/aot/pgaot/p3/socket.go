package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_socket_putmessage_noblock(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v7 = *(*int32)(unsafe.Add(mBase, _consts[582]))
	v10 = l2 + v7 + int32(5)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[583]))
	if v12 < v10 {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[584]))
		v16 = F_repalloc(m, v15, v10)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[583])) = v10
			*(*int32)(unsafe.Add(mBase, _consts[584])) = v16
			v24 = *(*int32)(unsafe.Add(mBase, _consts[465]))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
			v26 = m.T0[v25].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, _consts[465]))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
		v26 = m.T0[v25].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			return
		}
	}
}
