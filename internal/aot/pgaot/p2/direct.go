package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DirectFunctionCall5Coll(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v28 int32
	_ = v28
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v8 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+104)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = l2
	v28 = int32(5)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+26)) = uint16(v28)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
	v39 = m.T0[l0].(func(*base.Module, int32) int64)(m, v11+int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int64(0)
	} else {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
		if v43 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
				F_errmsg_internal(m, int32(_a_F_DirectFunctionCall5Coll_0), v11)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_DirectFunctionCall5Coll_1), int32(911), int32(_a_F_DirectFunctionCall5Coll_2))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(112)
			return v39
		}
	}
}
