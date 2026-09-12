package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_wal_flush_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[111])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[112]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[111])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(136236), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(600796), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(519185), int32(324), int32(257358))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
	} else {
		v40 = int32(4457632)
		v41 = *(*int32)(unsafe.Add(mBase, _consts[112]))
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v41)+280))
		*(*int64)(unsafe.Add(mBase, uint32(v41)+280)) = v42
		*(*int64)(unsafe.Add(mBase, _consts[117])) = v42
		v47 = *(*int32)(unsafe.Add(mBase, _consts[112]))
		v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)+272))
		*(*int64)(unsafe.Add(mBase, uint32(v47)+272)) = v48
		*(*int64)(unsafe.Add(mBase, _consts[116])) = v48
		v57 = *(*int64)(unsafe.Add(mBase, _consts[117]))
		v58 = F_Int64GetDatum(m, v57)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int32(0)
		} else {
			return v58
		}
	}
}
