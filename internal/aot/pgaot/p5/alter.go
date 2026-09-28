package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AlterTableInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v5 = F_AlterTableGetLockLevel(m, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v7 = F_relation_open(m, l0, v5)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_AlterTableInternal[0]))
			if v11 == int32(0) {
			} else {
				v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
				if v14 != 0 {
				} else {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l0
				}
			}
			v17 = int32(0)
			F_ATController(m, v17, v7, l1, int32(1), v5, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_CheckAlterSubOption(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v10 != int32(1) {
		if l2 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
						F_errmsg(m, int32(_a_F_CheckAlterSubOption_0), v6+int32(-48))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckAlterSubOption_1), int32(1512), int32(_a_F_CheckAlterSubOption_2))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v17 = v6 + int32(-16)
				F_initStringInfo(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
					F_appendStringInfo(m, v17, int32(_a_F_CheckAlterSubOption_3), v6+int32(-32))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						F_PreventInTransactionBlock(m, l3, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							F_pfree(m, v29)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								m.G0 = v8 - int32(-64)
								return
							}
						}
					}
				}
			}
		} else {
			m.G0 = v8 - int32(-64)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_errmsg(m, int32(_a_F_CheckAlterSubOption_4), v8)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_CheckAlterSubOption_1), int32(1498), int32(_a_F_CheckAlterSubOption_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_ExecAlterObjectDependsStmt(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v18 == int32(0) {
		v39 = v15
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_get_object_address(m, l0, v14, v39, v12+int32(16), int32(8), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L3
	} else {
		goto L14
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v22 = F_makeString(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v24 = F_lcons(m, v22, v15)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = F_makeString(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	v31 = v24
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v32 == int32(0) {
		v39 = v31
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v29 = F_lcons(m, v27, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v31 = v29
	goto L8
L11:
	;
	v35 = F_makeString(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v37 = F_lcons(m, v35, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v39 = v37
	goto L1
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAlterObjectDependsStmt[0]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_check_object_ownership(m, v47, v49, v12, v48, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_relation_close(m, v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v65 = int32(0)
	F_get_object_address(m, v12+int32(20), int32(15), v64, v65, int32(8), v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L3
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v72
	goto L23
L22:
	;
	goto L23
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)))
	if v76 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	m.G0 = v12 + int32(32)
	return
L25:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	F_deleteDependencyRecordsForSpecific(m, v75, v74, int32(120), v80, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v84 = int32(0)
	v85 = m.G0
	v87 = v85 - int32(112)
	m.G0 = v87
	v91 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	F_ScanKeyInit(m, v87, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v75))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_ScanKeyInit(m, v87+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(v74))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v111 = F_systable_beginscan(m, v91, int32(2673), int32(1), int32(0), int32(2), v87)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v113 = F_systable_getnext(m, v111)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v113 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v117 = v84
	v118 = v113
	goto L37
L35:
	;
	v141 = v84
	goto L36
L36:
	;
	F_systable_endscan(m, v111)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L45
	}
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+16))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+22)))
	v126 = v124 + v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	if v127 != int32(3079) {
		v136 = v117
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v141 = v136
	goto L36
L39:
	;
	v137 = F_systable_getnext(m, v111)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L3
	} else {
		goto L43
	}
L40:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
	if v130 != int32(120) {
		v136 = v117
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v134 = F_lappend_oid(m, v117, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v136 = v134
	goto L39
L43:
	;
	if v137 != 0 {
		v117 = v136
		v118 = v137
		goto L37
	} else {
		goto L44
	}
L44:
	;
	goto L38
L45:
	;
	F_relation_close(m, v91, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v87 + int32(112)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v157 = int32(0)
	if v141 == v157 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v195 != 0 {
		goto L24
	} else {
		goto L60
	}
L48:
	;
	v195 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	if v163 <= int32(0) {
		v189 = v157
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v195 = v189
	goto L47
L52:
	;
	v166 = int32(0)
	if v166 < v163 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v169 = v163
	goto L55
L54:
	;
	v169 = v166
	goto L55
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v141)+12))
	v172 = int32(0)
	goto L56
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v170+v172<<(uint(int32(2))%32))))
	v181 = base.B2i32(v180 == v156)
	if v180 == v156 {
		v189 = v181
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v189 = v181
	goto L51
L58:
	;
	v183 = v172 + int32(1)
	if v183 != v169 {
		v172 = v183
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	F_recordDependencyOn(m, l0, v12+int32(20), int32(120))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	goto L24
}
