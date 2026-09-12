package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MarkPostmasterChildInactive(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1152]))
	*(*int32)(unsafe.Add(mBase, uint32(v4+v6<<(uint(int32(2))%32))+44)) = int32(1)
	return
}
func F_PostmasterIsAliveInternal(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1153]))
	v12 = F_read(m, v8, v5+int32(15), int32(1))
	mBase = m.M
	if v12 < int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		if v16 == int32(6) {
			m.G0 = v5 + int32(16)
			return int32(base.Ui32(v12) >> (uint(int32(31)) % 32))
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(282008), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(475899), int32(382), int32(298529))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		if v12 == int32(0) {
			m.G0 = v5 + int32(16)
			return int32(base.Ui32(v12) >> (uint(int32(31)) % 32))
		} else {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(355370), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(475899), int32(384), int32(298529))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_SendPostmasterSignal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[418])))
	if v3 == int32(1) {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[773]))
		*(*int32)(unsafe.Add(mBase, uint32(v7+l0<<(uint(int32(2))%32)))) = int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[1151]))
		v16 = F_kill(m, v14, int32(10))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
