package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_replication_origin_session_progress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_replorigin_check_prerequisites(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_progress[0]))
		if v12 != 0 {
			v16 = F_LWLockAcquire(m, v12+int32(44), int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_progress[0]))
				v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
				F_LWLockRelease(m, v19+int32(44))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = int64(0)
					if base.B2i32(v6 == v26)|base.B2i32(v20 == v26) == int32(0) {
						F_XLogFlush(m, v20)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							if v21 == int64(0) {
								v37 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
							} else {
							}
							return v21
						}
					} else {
						if v21 == int64(0) {
							v37 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
						} else {
						}
						return v21
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_replication_origin_session_progress_0), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_replication_origin_session_progress_1), int32(1544), int32(_a_F_pg_replication_origin_session_progress_2))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
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
	}
}
func F_pg_replication_origin_session_setup(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	F_replorigin_check_prerequisites(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v9 = F_text_to_cstring(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v12 = F_replorigin_by_name(m, v9, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				F_replorigin_session_setup(m, v12, v14)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					*(*uint16)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_setup[0])) = uint16(v12)
					F_pfree(m, v9)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						return int64(0)
					}
				}
			}
		}
	}
}
func F_pg_replication_origin_xact_reset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int64
	_ = v7
	F_replorigin_check_prerequisites(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v7 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_xact_reset[0])) = v7
		*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_xact_reset[1])) = v7
		return v7
	}
}
func F_pg_replication_origin_xact_setup(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_replorigin_check_prerequisites(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_xact_setup[0]))
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_replication_origin_xact_setup_0), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_replication_origin_xact_setup_1), int32(1564), int32(_a_F_pg_replication_origin_xact_setup_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_xact_setup[1])) = v3
			v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_xact_setup[2])) = v31
			return int64(0)
		}
	}
}
