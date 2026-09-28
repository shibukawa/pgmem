package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_find_base_rel_ignore_join(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v10) <= base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(_a_F_find_base_rel_ignore_join_0), v8)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_find_base_rel_ignore_join_1), int32(606), int32(_a_F_find_base_rel_ignore_join_2))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v13 = l1 << (uint(int32(2)) % 32)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)))
		if v16 != 0 {
			m.G0 = v8 + int32(16)
			return v16
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17+v13)))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
					F_errmsg_internal(m, int32(_a_F_find_base_rel_ignore_join_0), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_find_base_rel_ignore_join_1), int32(606), int32(_a_F_find_base_rel_ignore_join_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
				if v22 != int32(2) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
						F_errmsg_internal(m, int32(_a_F_find_base_rel_ignore_join_0), v8)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_find_base_rel_ignore_join_1), int32(606), int32(_a_F_find_base_rel_ignore_join_2))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
					if v25 != 0 {
						m.G0 = v8 + int32(16)
						return v16
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
							F_errmsg_internal(m, int32(_a_F_find_base_rel_ignore_join_0), v8)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_find_base_rel_ignore_join_1), int32(606), int32(_a_F_find_base_rel_ignore_join_2))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
		}
	}
}
func F_getBaseType(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = F_getBaseTypeAndTypmod(m, l0, v5+int32(12))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v9
	}
}
func F_getBaseTypeAndTypmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_ReleaseCatCache(m, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L3
	} else {
		goto L16
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L3
	} else {
		goto L13
	}
L3:
	;
	return int32(0)
L4:
	;
	if v12 == int32(0) {
		v46 = l0
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
	v20 = v18 + v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+79)))
	if v21 != int32(100) {
		v64 = l0
		v67 = v12
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v26 = v20
	v27 = v12
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+132))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v30
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v64 = v29
	v67 = v36
	goto L1
L9:
	;
	v36 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v29))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v36 == int32(0) {
		v46 = v29
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v42 = v40 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+79)))
	if v43 == int32(100) {
		v26 = v42
		v27 = v36
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v46
	F_errmsg_internal(m, int32(_a_F_getBaseTypeAndTypmod_0), v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_getBaseTypeAndTypmod_1), int32(2864), int32(_a_F_getBaseTypeAndTypmod_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	m.G0 = v8 + int32(16)
	return v64
}
