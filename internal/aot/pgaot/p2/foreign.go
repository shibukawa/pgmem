package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetForeignKeyCheckTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	F_ScanKeyInit(m, v12+int32(32), int32(11), int32(3), int32(184), l1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = int32(1)
	v31 = F_systable_beginscan(m, l0, int32(2699), v26, int32(0), v26, v12+int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v94 != 0 {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	v33 = F_systable_getnext(m, v31)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = v33
	goto L7
L7:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46 + v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+84))
	if v49 != l2 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v83 = F_systable_getnext(m, v31)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 != l3 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+76))
	v56 = v53 - int32(1644)
	if base.Ui32(v56) <= base.Ui32(int32(11)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v64 != int32(2) {
		goto L9
	} else {
		goto L16
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56<<(uint(int32(2))%32))+uint32(_consts[403])))
	v64 = v63
	goto L15
L14:
	;
	v64 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+80)))
	if v67&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v77 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v70
	v77 = v70
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v67&int32(16) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v74
	goto L23
L22:
	;
	goto L23
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v77 = v76
	goto L17
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v80 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L9
L26:
	;
	if v83 != 0 {
		v37 = v83
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v95 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	F_systable_endscan(m, v31)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v12 + int32(80)
	return
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
	F_errmsg_internal(m, int32(40534), v12)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(491372), int32(12179), int32(134013))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg_internal(m, int32(40603), v12+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(491372), int32(12182), int32(134013))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_createForeignKeyCheckTriggers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = F_palloc0(m, int32(52))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(489181)
		v24 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(181)
		v29 = F_SystemFuncName(m, int32(149036))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v29
			v32 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+26)) = int32(262144)
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v38)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v32
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+44)) = uint8(v42)
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
			*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v32
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+45)) = uint8(v44)
			F_CreateTrigger(m, v15+int32(4), v18, v32, l0, l1, l3, l4, l5, v38)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				F_CommandCounterIncrement(m)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l7))) = v54
					v59 = F_palloc0(m, int32(52))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = int32(489181)
						v65 = int32(256)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)) = uint16(v65)
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(181)
						v70 = F_SystemFuncName(m, int32(419575))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v70
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+26)) = int32(1048576)
							v79 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v73
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+44)) = uint8(v83)
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+48)) = v73
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+45)) = uint8(v85)
							F_CreateTrigger(m, v15+int32(4), v59, v73, l0, l1, l3, l4, l6, v79)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								F_CommandCounterIncrement(m)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l8))) = v95
									m.G0 = v15 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_has_foreign_data_wrapper_privilege_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v16)
		v20 = *(*int32)(unsafe.Add(mBase, _consts[276]))
		v22 = F_convert_any_priv_string(m, v12, int32(1640672))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v26 = F_object_aclcheck_ext(m, int32(2328), v10, v20, v22, v8+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v28 == int32(1) {
					v31 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v31)
					v35 = int32(0)
				} else {
					v35 = base.B2i32(v26 == int32(0))
				}
				m.G0 = v8 + int32(16)
				return v35
			}
		}
	}
}
func F_has_foreign_data_wrapper_privilege_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_pg_detoast_datum_packed(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[276]))
			v15 = F_text_to_cstring(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v18 = F_get_foreign_data_wrapper_oid(m, v15, int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v21 = F_convert_any_priv_string(m, v10, int32(1640672))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = F_object_aclcheck(m, int32(2328), v18, v13, v21)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v23 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_has_foreign_data_wrapper_privilege_name_name(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_get_role_oid_or_public(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v16 = F_text_to_cstring(m, v6)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v19 = F_get_foreign_data_wrapper_oid(m, v16, int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = F_convert_any_priv_string(m, v11, int32(1640672))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = F_object_aclcheck(m, int32(2328), v19, v13, v22)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v24 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
