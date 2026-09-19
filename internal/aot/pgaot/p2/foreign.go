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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v7
	v19 = v12 + int32(32)
	F_ScanKeyInit(m, v19, int32(11), int32(3), int32(184), l1)
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
	v29 = F_systable_beginscan(m, l0, int32(2699), v26, int32(0), v26, v19)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v90 != 0 {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	v31 = F_systable_getnext(m, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v31 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = v31
	goto L7
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+22)))
	v46 = v44 + v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	if v47 != l2 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L3
L9:
	;
	v79 = F_systable_getnext(m, v29)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 != l3 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	v53 = v51 - int32(1644)
	if base.Ui32(v53) <= base.Ui32(int32(11)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v60 != int32(2) {
		goto L9
	} else {
		goto L16
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(2))%32))+uint32(_c_F_GetForeignKeyCheckTriggers[0])))
	v60 = v58
	goto L15
L14:
	;
	v60 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+80)))
	if v63&int32(4) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v73 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v66
	v73 = v66
	goto L17
L19:
	;
	goto L20
L20:
	;
	if v63&int32(16) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v70
	goto L23
L22:
	;
	goto L23
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v73 = v72
	goto L17
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v76 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	goto L9
L26:
	;
	if v79 != 0 {
		v35 = v79
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
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v91 == int32(0) {
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
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	F_systable_endscan(m, v29)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_GetForeignKeyCheckTriggers_0), v12)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_GetForeignKeyCheckTriggers_1), int32(_a_F_GetForeignKeyCheckTriggers_2), int32(_a_F_GetForeignKeyCheckTriggers_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_GetForeignKeyCheckTriggers_4), v12+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_GetForeignKeyCheckTriggers_1), int32(_a_F_GetForeignKeyCheckTriggers_5), int32(_a_F_GetForeignKeyCheckTriggers_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
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
	var v49 int32
	_ = v49
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
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
		*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(_a_F_createForeignKeyCheckTriggers_0)
		v24 = int32(256)
		*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v24)
		*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(181)
		v29 = F_SystemFuncName(m, int32(_a_F_createForeignKeyCheckTriggers_1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v29
			v32 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v32
			*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+26)) = int32(_a_F_createForeignKeyCheckTriggers_2)
			v38 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)) = uint8(v38)
			*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v32
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+44)) = uint8(v42)
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
			*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v32
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+45)) = uint8(v44)
			v49 = v15 + int32(4)
			F_CreateTrigger(m, v49, v18, v32, l0, l1, l3, l4, l5, v38)
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
						*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = int32(_a_F_createForeignKeyCheckTriggers_0)
						v65 = int32(256)
						*(*uint16)(unsafe.Add(mBase, uint32(v59)+4)) = uint16(v65)
						*(*int32)(unsafe.Add(mBase, uint32(v59))) = int32(181)
						v70 = F_SystemFuncName(m, int32(_a_F_createForeignKeyCheckTriggers_3))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v70
							v73 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v59)+32)) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+26)) = int32(_a_F_createForeignKeyCheckTriggers_4)
							v79 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+24)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v73
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+44)) = uint8(v83)
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+13)))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+48)) = v73
							*(*uint8)(unsafe.Add(mBase, uint32(v59)+45)) = uint8(v85)
							F_CreateTrigger(m, v49, v59, v73, l0, l1, l3, l4, l6, v79)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
								F_CommandCounterIncrement(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l8))) = v93
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13935(m, l0, int32(_a_F_has_foreign_data_wrapper_privilege_id_0), int32(2328))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_has_foreign_data_wrapper_privilege_name[0]))
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
					v21 = F_convert_any_priv_string(m, v10, int32(_a_F_has_foreign_data_wrapper_privilege_name_0))
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
						v22 = F_convert_any_priv_string(m, v11, int32(_a_F_has_foreign_data_wrapper_privilege_name_name_0))
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
