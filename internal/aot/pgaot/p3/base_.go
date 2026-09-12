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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if base.Ui32(v10) <= base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(501382), v8)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(524449), int32(476), int32(289140))
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
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13+v14)))
		if v16 != 0 {
			m.G0 = v8 + int32(16)
			return v16
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17+v13)))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
					F_errmsg_internal(m, int32(501382), v8)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524449), int32(476), int32(289140))
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
						F_errmsg_internal(m, int32(501382), v8)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(524449), int32(476), int32(289140))
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
							F_errmsg_internal(m, int32(501382), v8)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524449), int32(476), int32(289140))
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_ReleaseCatCache(m, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L16
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L3
	} else {
		goto L13
	}
L3:
	;
	return int32(0)
L4:
	;
	if v11 == int32(0) {
		v44 = l0
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
	v19 = v17 + v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+79)))
	if v20 != int32(100) {
		v62 = l0
		v65 = v11
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v25 = v19
	v26 = v11
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+132))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+136))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v62 = v28
	v65 = v34
	goto L1
L9:
	;
	v34 = F_SearchSysCache1(m, int32(82), v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v34 == int32(0) {
		v44 = v28
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	v40 = v38 + v39
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+79)))
	if v41 == int32(100) {
		v25 = v40
		v26 = v34
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v44
	F_errmsg_internal(m, int32(55200), v8)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(524184), int32(2690), int32(442915))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
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
	return v62
}
