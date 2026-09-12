package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_restoreTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v4 = m.G0
	v6 = v4 - int32(1104)
	m.G0 = v6
	if base.Ui32(l0) < base.Ui32(l1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = l0
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v6 + int32(1104)
	return
L4:
	;
	if v9 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v39 = v9 + int32(1)
	if v39 != l1 {
		v9 = v39
		goto L4
	} else {
		goto L13
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v9
	v19 = F_pg_snprintf(m, v6+int32(16), int32(64), int32(_a_F_restoreTimeLineHistoryFiles_0), v6)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	v28 = F_RestoreArchivedFile(m, v6+int32(80), v6+int32(16), int32(_a_F_restoreTimeLineHistoryFiles_1), int64(0), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v28 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_KeepFileRestoredFromArchive(m, v6+int32(80), v6+int32(16))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2724), int32(0), v4, v5)
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v8 != int32(-2147483648) {
		if v8 == int32(2147483647) {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			if v21 != int32(2147483647) {
				v45 = v20
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v45+v6, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				v54 = F_Int64GetDatum(m, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v54
				}
			} else {
				if v20 != int64(9223372036854775807) {
					v45 = v20
					v47 = int64(86400000000)
					v48 = base.I64_rem_s(v45+v6, v47)
					if v48 < int64(0) {
						v53 = v48 + v47
					} else {
						v53 = v48
					}
					v54 = F_Int64GetDatum(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						return v54
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2126), int32(_a_F_time_pl_interval_2))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
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
		} else {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
			v45 = v13
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v45+v6, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			v54 = F_Int64GetDatum(m, v53)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				return v54
			}
		}
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
		if v15 != int32(-2147483648) {
			v45 = v14
			v47 = int64(86400000000)
			v48 = base.I64_rem_s(v45+v6, v47)
			if v48 < int64(0) {
				v53 = v48 + v47
			} else {
				v53 = v48
			}
			v54 = F_Int64GetDatum(m, v53)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				return v54
			}
		} else {
			if v14 == int64(-9223372036854775807-1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_time_pl_interval_0), int32(0))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_time_pl_interval_1), int32(2126), int32(_a_F_time_pl_interval_2))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
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
				v45 = v14
				v47 = int64(86400000000)
				v48 = base.I64_rem_s(v45+v6, v47)
				if v48 < int64(0) {
					v53 = v48 + v47
				} else {
					v53 = v48
				}
				v54 = F_Int64GetDatum(m, v53)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v54
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
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.Ui32(v7) <= base.Ui32(int32(6)) {
		v11 = v7 << (uint(int32(3)) % 32)
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_time_scale[0])))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_time_scale[1])))
		if int64(0) <= v6 {
			v20 = v6 + v17
			v21 = base.I64_rem_s(v20, v14)
			v23 = F_Int64GetDatum(m, v20-v21)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v23
			}
		} else {
			v28 = v17 - v6
			v29 = base.I64_rem_s(v28, v14)
			v35 = v29 - v28
			v36 = F_Int64GetDatum(m, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				return v36
			}
		}
	} else {
		v35 = v6
		v36 = F_Int64GetDatum(m, v35)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			return v36
		}
	}
}
