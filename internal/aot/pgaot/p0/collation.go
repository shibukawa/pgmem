package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CollationGetCollid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CollationGetCollid[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	goto L1
L1:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return int32(0)
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_CollationGetCollid[1]))
	if v15 == int32(0) {
		v80 = v2
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v80
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) < v18 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = v2
	goto L9
L7:
	;
	goto L8
L8:
	;
	v80 = int32(0)
	goto L4
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v26<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_CollationGetCollid[2]))
	if v31 == v33 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v69 = v26 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v69 < v70 {
		v26 = v69
		goto L9
	} else {
		goto L24
	}
L12:
	;
	v37 = F_GetSysCacheOid(m, int32(15), l0, v9, v31, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if v37 != 0 {
		v80 = v37
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v41 = F_SearchSysCache3(m, int32(15), l0, int32(-1), v31)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v41 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v47 = v45 + v46
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+76)))
	if v48 != int32(105) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	F_ReleaseCatCache(m, v41)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	if base.B2i32(base.Ui32(v9) < base.Ui32(int32(35)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v9))%64)&int64(34357509982) != int64(0)) != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	F_ReleaseCatCache(m, v41)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L11
L22:
	;
	if v63 != 0 {
		v80 = v63
		goto L4
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	goto L10
}
func F_get_collation_actual_version_builtin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == int32(67) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v5 + int32(16)
	return int32(_a_F_get_collation_actual_version_builtin_0)
L2:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v13 = int32(_a_F_get_collation_actual_version_builtin_1)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_collation_actual_version_builtin[0])))
	if base.B2i32(v16 == int32(0))|base.B2i32(v16 != v19) != 0 {
		v37 = v16
		v38 = v19
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L4
L6:
	;
	if v37-v38 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L7:
	;
	goto L6
L8:
	;
	v22 = l0
	v23 = v13
	goto L9
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v27 == int32(0) {
		v37 = v27
		v38 = v26
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v37 = v27
	v38 = v26
	goto L7
L11:
	;
	v30 = int32(1)
	if v27 == v26 {
		v22 = v22 + v30
		v23 = v23 + v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(_a_F_get_collation_actual_version_builtin_2)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_get_collation_actual_version_builtin[1])))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v66-v67 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L15:
	;
	goto L14
L16:
	;
	v51 = l0
	v52 = v42
	goto L17
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v66 = v56
	v67 = v55
	goto L15
L19:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
	F_errmsg(m, int32(_a_F_get_collation_actual_version_builtin_3), v5)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_get_collation_actual_version_builtin_4), int32(189), int32(_a_F_get_collation_actual_version_builtin_5))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_merge_collation_state(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	if base.Ui32(v13) < base.Ui32(l1) {
		*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = l0
		if l1 != int32(2) {
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = l3
			v40 = int32(20)
			v41 = l4
			*(*int32)(unsafe.Add(mBase, uint32(v40+l5))) = v41
		}
		m.G0 = v11 + int32(16)
		return
	} else {
		if l1 != v13 {
			m.G0 = v11 + int32(16)
			return
		} else {
			switch l1 - int32(1) {
			case 0:
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				if l0 == v25 {
				} else {
					if v25 == int32(100) {
						*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = l0
						v39 = int32(12)
						v40 = v39
						v41 = l2
						*(*int32)(unsafe.Add(mBase, uint32(v40+l5))) = v41
					} else {
						if l0 == int32(100) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(l5)+8)) = int32(2)
							v39 = int32(20)
							v40 = v39
							v41 = l2
							*(*int32)(unsafe.Add(mBase, uint32(v40+l5))) = v41
						}
					}
				}
				m.G0 = v11 + int32(16)
				return
			default:
				m.G0 = v11 + int32(16)
				return
			case 2:
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
				if l0 == v44 {
					m.G0 = v11 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						F_errcode(m, int32(17432708))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
							v54 = F_get_collation_name(m, v53)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v56 = F_get_collation_name(m, l0)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
									F_errmsg(m, int32(_a_F_merge_collation_state_0), v11)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
										F_parser_errposition(m, v63, l2)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_merge_collation_state_1), int32(858), int32(_a_F_merge_collation_state_2))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
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
		}
	}
}
