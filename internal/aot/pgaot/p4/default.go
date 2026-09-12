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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
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
	F_ScanKeyInit(m, v12+int32(-48), int32(2), int32(3), int32(184), l1)
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
	v39 = F_systable_beginscan(m, v24, int32(2686), v34, int32(0), v34, v12+int32(-48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v39)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L7:
	;
	v41 = F_systable_getnext(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v41 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v99 = v3
	v100 = v3
	v102 = v3
	v108 = int32(0)
	goto L6
L10:
	;
	goto L11
L11:
	;
	v47 = v41
	v48 = v3
	v49 = v3
	v51 = v3
	v53 = v3
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v59 = v57 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+88)))
	if v60 != int32(1) {
		v88 = v48
		v89 = v49
		v90 = v51
		v91 = v53
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v99 = v88
	v100 = v89
	v102 = v90
	v108 = base.B2i32(v91 == int32(1))
	goto L6
L14:
	;
	v93 = F_systable_getnext(m, v39)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+84))
	if v16 == v63 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v84 = v49
	v85 = v53
	v86 = v48 + int32(1)
	goto L18
L17:
	;
	if v48 != 0 {
		v88 = v48
		v89 = v49
		v90 = v51
		v91 = v53
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v88 = v86
	v89 = v84
	v90 = v87
	v91 = v85
	goto L14
L19:
	;
	v67 = int32(0)
	v68 = F_IsBinaryCoercible(m, v16, v63)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	if v68 == int32(0) {
		v88 = v67
		v89 = v49
		v90 = v51
		v91 = v53
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v59)+84))
	v73 = F_IsPreferredType(m, v20, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v73 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v80 = v53
	v81 = v49 + int32(1)
	goto L25
L24:
	;
	if v49 != 0 {
		v88 = v67
		v89 = v49
		v90 = v51
		v91 = v53
		goto L14
	} else {
		goto L26
	}
L25:
	;
	v84 = v81
	v85 = v80
	v86 = int32(0)
	goto L18
L26:
	;
	v80 = v53 + int32(1)
	v81 = int32(0)
	goto L25
L27:
	;
	if v93 != 0 {
		v47 = v93
		v48 = v88
		v49 = v89
		v51 = v90
		v53 = v91
		goto L12
	} else {
		goto L28
	}
L28:
	;
	goto L13
L29:
	;
	F_sequence_close(m, v24, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if int32(2) <= v99 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	m.G0 = v14 - int32(-64)
	v137 = int32(0)
	if v108&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	F_errcode(m, int32(290948))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v123 = F_format_type_be(m, v16)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v123
	F_errmsg(m, int32(193742), v14)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(495411), int32(2421), int32(131318))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	v141 = v102
	goto L41
L40:
	;
	v141 = v137
	goto L41
L41:
	;
	if v100 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v142 = v137
	goto L44
L43:
	;
	v142 = v141
	goto L44
L44:
	;
	if v100 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v145 = v102
	goto L47
L46:
	;
	v145 = v142
	goto L47
L47:
	;
	if v99 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v148 = v102
	goto L50
L49:
	;
	v148 = v145
	goto L50
L50:
	;
	return v148
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
		v15 = *(*int32)(unsafe.Add(mBase, _consts[140]))
		*(*int32)(unsafe.Add(mBase, _consts[141])) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(423436)
		v22 = F_format_elog_string(m, int32(572258), v7)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[142])) = v22
			v99 = int32(0)
			m.G0 = v7 - int32(-64)
			return v99
		}
	} else {
		v27 = F_strlen(m, v9)
		mBase = m.M
		if base.Ui32(int32(64)) <= base.Ui32(v27) {
			v32 = *(*int32)(unsafe.Add(mBase, _consts[140]))
			*(*int32)(unsafe.Add(mBase, _consts[141])) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(63)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(423436)
			v43 = F_format_elog_string(m, int32(662103), v5+int32(-48))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[142])) = v43
				v99 = int32(0)
				m.G0 = v7 - int32(-64)
				return v99
			}
		} else {
			v46 = int32(1)
			v48 = *(*int32)(unsafe.Add(mBase, _consts[25]))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
			if base.B2i32(v49 == int32(2)) == int32(0) {
				v99 = v46
				m.G0 = v7 - int32(-64)
				return v99
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
											F_errmsg(m, int32(72971), v5+int32(-32))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(498889), int32(137), int32(423430))
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
								v87 = *(*int32)(unsafe.Add(mBase, _consts[140]))
								*(*int32)(unsafe.Add(mBase, _consts[141])) = v87
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v90
								v96 = F_format_elog_string(m, int32(579761), v5+int32(-16))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, _consts[142])) = v96
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
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(45), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+8))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
