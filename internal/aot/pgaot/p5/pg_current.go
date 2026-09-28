package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_wal_flush_lsn(m *base.Module, l0 int32) int64 {
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	var v63 int64
	_ = v63
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+308))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[0])) = uint8(v12)
		v14 = v12
	} else {
		v14 = int32(0)
	}
	if v14 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_current_wal_flush_lsn_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					F_errhint(m, int32(_a_F_pg_current_wal_flush_lsn_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_current_wal_flush_lsn_2), int32(354), int32(_a_F_pg_current_wal_flush_lsn_3))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
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
		v37 = int32(0)
		v39 = int32(_a_F_pg_current_wal_flush_lsn_4)
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[1]))
		v41 = int64(0)
		v44 = base.AtomicRmwCmpxchg64(m, v40, int32(272), v41, v41)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[2])) = v44
		v49 = base.AtomicRmwOr32(m, v37, int32(_a_F_pg_current_wal_flush_lsn_5), v37)
		v52 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[1]))
		v56 = base.AtomicRmwCmpxchg64(m, v52, int32(264), v41, v41)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[3])) = v56
		v63 = *(*int64)(unsafe.Add(mBase, _c_F_pg_current_wal_flush_lsn[2]))
		return v63
	}
}
