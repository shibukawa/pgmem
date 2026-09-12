package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_replication_origin_drop(m *base.Module, l0 int32) int32 {
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v12)
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
			F_errcode(m, int32(100663618))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(14386), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493370), int32(200), int32(159111))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v34 = F_text_to_cstring(m, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			F_replorigin_drop_by_name(m, v34, int32(0), int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v34)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		}
	}
}
