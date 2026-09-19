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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
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
	v27 = v12 + int32(-48)
	F_ScanKeyInit(m, v27, int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = int32(1)
	v37 = F_systable_beginscan(m, v24, int32(2686), v34, int32(0), v34, v27)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = F_systable_getnext(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v42 = v39
	v43 = v3
	v46 = v3
	v47 = v3
	v49 = v3
	goto L11
L9:
	;
	v95 = v3
	v98 = v3
	v99 = v3
	v104 = int32(0)
	goto L10
L10:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L28
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v54 = v52 + v53
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+88)))
	if v55 != int32(1) {
		v83 = v43
		v84 = v46
		v85 = v47
		v86 = v49
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v95 = v83
	v98 = v84
	v99 = v85
	v104 = base.B2i32(v86 == int32(1))
	goto L10
L13:
	;
	v88 = F_systable_getnext(m, v37)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L26
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	if v16 == v58 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v79 = v46
	v80 = v49
	v81 = v43 + int32(1)
	goto L17
L16:
	;
	if v43 != 0 {
		v83 = v43
		v84 = v46
		v85 = v47
		v86 = v49
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v83 = v81
	v84 = v79
	v85 = v82
	v86 = v80
	goto L13
L18:
	;
	v62 = int32(0)
	v63 = F_IsBinaryCoercible(m, v16, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v63 == int32(0) {
		v83 = v62
		v84 = v46
		v85 = v47
		v86 = v49
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+84))
	v68 = F_IsPreferredType(m, v20, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v75 = v49
	v76 = v46 + int32(1)
	goto L24
L23:
	;
	if v46 != 0 {
		v83 = v62
		v84 = v46
		v85 = v47
		v86 = v49
		goto L13
	} else {
		goto L25
	}
L24:
	;
	v79 = v76
	v80 = v75
	v81 = int32(0)
	goto L17
L25:
	;
	v75 = v49 + int32(1)
	v76 = int32(0)
	goto L24
L26:
	;
	if v88 != 0 {
		v42 = v88
		v43 = v83
		v46 = v84
		v47 = v85
		v49 = v86
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
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if int32(2) <= v95 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	m.G0 = v14 - int32(-64)
	v133 = int32(0)
	if v104&int32(1) != 0 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	F_errcode(m, int32(_a_F_GetDefaultOpClass_0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v119 = F_format_type_be(m, v16)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v119
	F_errmsg(m, int32(_a_F_GetDefaultOpClass_1), v14)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_GetDefaultOpClass_2), int32(2421), int32(_a_F_GetDefaultOpClass_3))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
	v137 = v99
	goto L40
L39:
	;
	v137 = v133
	goto L40
L40:
	;
	if v98 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v138 = v133
	goto L43
L42:
	;
	v138 = v137
	goto L43
L43:
	;
	if v98 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v141 = v99
	goto L46
L45:
	;
	v141 = v138
	goto L46
L46:
	;
	if v95 == int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v144 = v99
	goto L49
L48:
	;
	v144 = v141
	goto L49
L49:
	;
	return v144
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
												F_errfinish(m, int32(_a_F_check_default_table_access_method_4), int32(137), int32(_a_F_check_default_table_access_method_5))
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
	v3 = Fn13899(m, l0, int32(45))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
