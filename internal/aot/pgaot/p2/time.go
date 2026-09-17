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
func F_time_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(2705), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_abs_interval(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_time_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v12 = F_hash_bytes_uint32(m, base.I32_wrap_i64(v4>>(uint(int64(63))%64)^int64(base.Ui64(v4)>>(uint(int64(32))%64))^v4))
	mBase = m.M
	return v12
}
func F_time_pl_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v8 != int32(-2147483648) {
		if v8 == int32(2147483647) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			if base.B2i32(v20 != int32(2147483647))|base.B2i32(v23 != int64(9223372036854775807)) != 0 {
				v46 = v23
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v46+v6, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				v55 = F_Int64GetDatum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v55
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2126), int32(_a_F_time_pl_interval_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v46 = v13
			v48 = int64(86400000000)
			v49 = base.I64_rem_s(v46+v6, v48)
			if v49 < int64(0) {
				v54 = v49 + v48
			} else {
				v54 = v49
			}
			v55 = F_Int64GetDatum(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return v55
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v15 != int32(-2147483648) {
			v46 = v14
			v48 = int64(86400000000)
			v49 = base.I64_rem_s(v46+v6, v48)
			if v49 < int64(0) {
				v54 = v49 + v48
			} else {
				v54 = v49
			}
			v55 = F_Int64GetDatum(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				return v55
			}
		} else {
			if v14 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2126), int32(_a_F_time_pl_interval_2))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
				v46 = v14
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v46+v6, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				v55 = F_Int64GetDatum(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v55
				}
			}
		}
	}
}
func F_time_scale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v7) <= base.Ui32(int32(6)) {
		v11 = v7 << (uint(int32(3)) % 32)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_time_scale[0])))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_time_scale[1])))
		if int64(0) <= v6 {
			v16 = v6 + v13
			v17 = base.I64_rem_s(v16, v12)
			v19 = F_Int64GetDatum(m, v16-v17)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return v19
			}
		} else {
			v24 = v13 - v6
			v25 = base.I64_rem_s(v24, v12)
			v31 = v25 - v24
			v32 = F_Int64GetDatum(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				return v32
			}
		}
	} else {
		v31 = v6
		v32 = F_Int64GetDatum(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			return v32
		}
	}
}
