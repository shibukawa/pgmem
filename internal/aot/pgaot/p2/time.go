package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_restoreTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(1104)
	m.G0 = v8
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v8 + int32(1104)
	return
L4:
	;
	if v11 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v39 = v11 + int32(1)
	if v39 != l1 {
		v11 = v39
		goto L4
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	v20 = v8 + int32(16)
	v23 = F_pg_snprintf(m, v20, int32(64), int32(_a_F_restoreTimeLineHistoryFiles_0), v8)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v26 = v8 + int32(80)
	v30 = F_RestoreArchivedFile(m, v26, v20, int32(_a_F_restoreTimeLineHistoryFiles_1), int64(0), int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v30 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_KeepFileRestoredFromArchive(m, v26, v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	goto L5
}
func F_time_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(2919), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v11 = F_abs_interval(m, base.I32_wrap_i64(v6))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_time_hash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v11 int32
	_ = v11
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_hash_bytes_uint32(m, base.I32_wrap_i64(v3>>(uint(int64(63))%64)^int64(base.Ui64(v3)>>(uint(int64(32))%64))^v3))
	mBase = m.M
	return base.I64_extend_i32_u(v11)
}
func F_time_pl_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 != int32(-2147483648) {
		if v7 == int32(2147483647) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			if base.B2i32(v19 != int32(2147483647))|base.B2i32(v22 != int64(9223372036854775807)) != 0 {
				v45 = v22
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v45+v5, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				return v53
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2177), int32(_a_F_time_pl_interval_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
			v12 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			v45 = v12
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v45+v5, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			return v53
		}
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v14 != int32(-2147483648) {
			v45 = v13
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v45+v5, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			return v53
		} else {
			if v13 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2177), int32(_a_F_time_pl_interval_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
				v45 = v13
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v45+v5, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				return v53
			}
		}
	}
}
func F_time_scale(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v26 int64
	_ = v26
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v6) <= base.Ui32(int32(6)) {
		v10 = v6 << (uint(int32(3)) % 32)
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_time_scale[0])))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_time_scale[1])))
		if int64(0) <= v5 {
			v15 = v5 + v12
			v16 = base.I64_rem_s(v15, v11)
			return v15 - v16
		} else {
			v19 = v12 - v5
			v20 = base.I64_rem_s(v19, v11)
			v26 = v20 - v19
			return v26
		}
	} else {
		v26 = v5
		return v26
	}
}
