package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetDefaultOpClass(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = F_getBaseType(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = F_TypeCategory(m, v16)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_table_open(m, int32(2616), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = v14 + int32(16)
	F_ScanKeyInit(m, v27, int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = int32(1)
	v38 = F_systable_beginscan(m, v24, int32(2686), v35, int32(0), v35, v27)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = F_systable_getnext(m, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = v40
	v44 = v3
	v47 = v3
	v48 = v3
	v50 = v3
	goto L11
L9:
	;
	v96 = v3
	v99 = v3
	v100 = v3
	v105 = int32(0)
	goto L10
L10:
	;
	F_systable_endscan(m, v38)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L28
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+88)))
	if v56 != int32(1) {
		v84 = v44
		v85 = v47
		v86 = v48
		v87 = v50
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v96 = v84
	v99 = v85
	v100 = v86
	v105 = base.B2i32(v87 == int32(1))
	goto L10
L13:
	;
	v89 = F_systable_getnext(m, v38)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L26
	}
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+84))
	if v16 == v59 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v80 = v47
	v81 = v50
	v82 = v44 + int32(1)
	goto L17
L16:
	;
	if v44 != 0 {
		v84 = v44
		v85 = v47
		v86 = v48
		v87 = v50
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v84 = v82
	v85 = v80
	v86 = v83
	v87 = v81
	goto L13
L18:
	;
	v63 = int32(0)
	v64 = F_IsBinaryCoercible(m, v16, v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v64 == int32(0) {
		v84 = v63
		v85 = v47
		v86 = v48
		v87 = v50
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v55)+84))
	v69 = F_IsPreferredType(m, v20, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v69 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = v50
	v77 = v47 + int32(1)
	goto L24
L23:
	;
	if v47 != 0 {
		v84 = v63
		v85 = v47
		v86 = v48
		v87 = v50
		goto L13
	} else {
		goto L25
	}
L24:
	;
	v80 = v77
	v81 = v76
	v82 = int32(0)
	goto L17
L25:
	;
	v76 = v50 + int32(1)
	v77 = int32(0)
	goto L24
L26:
	;
	if v89 != 0 {
		v43 = v89
		v44 = v84
		v47 = v85
		v48 = v86
		v50 = v87
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(2) <= v96 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v14 + int32(80)
	v134 = int32(0)
	if v105&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	F_errcode(m, int32(_a_F_GetDefaultOpClass_0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v120 = F_format_type_be(m, v16)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v120
	F_errmsg(m, int32(_a_F_GetDefaultOpClass_1), v14)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_GetDefaultOpClass_2), int32(2463), int32(_a_F_GetDefaultOpClass_3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	v138 = v100
	goto L40
L39:
	;
	v138 = v134
	goto L40
L40:
	;
	if v99 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v139 = v134
	goto L43
L42:
	;
	v139 = v138
	goto L43
L43:
	;
	if v99 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v142 = v100
	goto L46
L45:
	;
	v142 = v139
	goto L46
L46:
	;
	if v96 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v145 = v100
	goto L49
L48:
	;
	v145 = v142
	goto L49
L49:
	;
	return v145
}
func F_check_default_table_access_method(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v10 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[1])) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_check_default_table_access_method_0)
		v22 = F_format_elog_string(m, int32(_a_F_check_default_table_access_method_1), v7)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[2])) = v22
			v99 = int32(0)
			m.G0 = v7 - int32(-64)
			return v99
		}
	} else {
		v27 = F_strlen(m, v9)
		mBase = m.M
		if base.Ui32(int32(64)) <= base.Ui32(v27) {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[1])) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(63)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(_a_F_check_default_table_access_method_0)
			v43 = F_format_elog_string(m, int32(_a_F_check_default_table_access_method_2), v5+int32(-48))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[2])) = v43
				v99 = int32(0)
				m.G0 = v7 - int32(-64)
				return v99
			}
		} else {
			v46 = int32(1)
			v48 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[3]))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
			if base.B2i32(v49 == int32(2)) == int32(0) {
				v99 = v46
				m.G0 = v7 - int32(-64)
				return v99
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[4]))
				if v55 == int32(0) {
					v99 = v46
					m.G0 = v7 - int32(-64)
					return v99
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v60 = F_get_table_am_oid(m, v58, int32(1))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						if v60 != 0 {
							v99 = v46
							m.G0 = v7 - int32(-64)
							return v99
						} else {
							if l2 == int32(12) {
								v66 = F_errstart(m, int32(18), int32(0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 == int32(0) {
										v99 = v46
										m.G0 = v7 - int32(-64)
										return v99
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v73
											F_errmsg(m, int32(_a_F_check_default_table_access_method_3), v5+int32(-32))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_check_default_table_access_method_4), int32(136), int32(_a_F_check_default_table_access_method_5))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int32(0)
												} else {
													v99 = v46
													m.G0 = v7 - int32(-64)
													return v99
												}
											}
										}
									}
								}
							} else {
								v87 = *(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[1])) = v87
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v90
								v96 = F_format_elog_string(m, int32(_a_F_check_default_table_access_method_6), v5+int32(-16))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _c_F_check_default_table_access_method[2])) = v96
									v99 = int32(0)
									m.G0 = v7 - int32(-64)
									return v99
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_default_partition_oid(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14288(m, l0, int32(45))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
