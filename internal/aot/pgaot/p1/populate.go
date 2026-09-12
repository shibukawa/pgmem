package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_recordset_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v9 != 0 {
		if v9 <= int32(1) {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = int64(309237645376)
			v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = v15
			v22 = F_hash_create(m, int32(410857), int32(100), v4+int32(-48), int32(1048))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v22
				m.G0 = v6 - int32(-64)
				return int32(0)
			}
		} else {
			m.G0 = v6 - int32(-64)
			return int32(0)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
				F_errmsg(m, int32(118517), v6)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518949), int32(4226), int32(88562))
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
