package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_current_logfile_1arg(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_pg_current_logfile(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_current_wal_lsn(m *base.Module, l0 int32) int32 {
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
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[0])))
	if v4 == int32(1) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[1]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+316))
		v12 = base.B2i32(v10 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[0])) = uint8(v12)
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
				F_errmsg(m, int32(_a_F_pg_current_wal_lsn_0), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errhint(m, int32(_a_F_pg_current_wal_lsn_1), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_current_wal_lsn_2), int32(282), int32(_a_F_pg_current_wal_lsn_3))
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
		v37 = int64(0)
		v39 = int32(_a_F_pg_current_wal_lsn_4)
		v40 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[1]))
		v44 = base.AtomicRmwCmpxchg64(m, v40, int32(280), v37, v37)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[2])) = v44
		v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[1]))
		v52 = base.AtomicRmwCmpxchg64(m, v48, int32(272), v37, v37)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_current_wal_lsn[3])) = v52
		v54 = F_Int64GetDatum(m, v52)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			return v54
		}
	}
}
