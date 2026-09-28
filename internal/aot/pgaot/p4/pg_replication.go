package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_replication_origin_advance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		F_replorigin_check_prerequisites(m)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_LockRelationOid(m, int32(_a_F_pg_replication_origin_advance_0), int32(3))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = F_text_to_cstring(m, v5)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					v19 = F_replorigin_by_name(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						v22 = int32(1)
						F_replorigin_advance(m, v19, v9, int64(0), v22, v22)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							F_UnlockRelationOid(m, int32(_a_F_pg_replication_origin_advance_0), int32(3))
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int64(0)
							} else {
								return int64(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_replication_origin_oid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_oid[0])))
	if v7 == int32(1) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_oid[1]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+308))
		v15 = base.B2i32(v13 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_oid[0])) = uint8(v15)
		v17 = v15
	} else {
		v17 = int32(0)
	}
	if v17 == int32(0) {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v22 = F_text_to_cstring(m, v21)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int64(0)
		} else {
			v26 = F_cstring_to_text(m, v22)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int64(0)
			} else {
				v29 = F_SearchSysCache1(m, int32(59), base.I64_extend_i32_u(v26))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					if v29 == int32(0) {
						F_pfree(m, v22)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							return int64(0)
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
						v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+v36))))
						F_ReleaseCatCache(m, v29)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v22)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int64(0)
							} else {
								if v38 == int32(0) {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v38)
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_replication_origin_oid_0), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_replication_origin_oid_1), int32(217), int32(_a_F_pg_replication_origin_oid_2))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
