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
			v11 = *(*int32)(unsafe.Add(mBase, _consts[286]))
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
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v10 != int32(1) {
		if l2 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v13 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
						F_errmsg(m, int32(378814), v6+int32(-48))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(491199), int32(1085), int32(246561))
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
			} else {
				F_initStringInfo(m, v6+int32(-16))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
					F_appendStringInfo(m, v6+int32(-16), int32(657140), v6+int32(-32))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
						F_PreventInTransactionBlock(m, l3, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
							F_pfree(m, v31)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
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
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
				F_errmsg(m, int32(246024), v8)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(491199), int32(1071), int32(246561))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
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
	var v50 int32
	_ = v50
	var v52 int64
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
	var v70 int64
	_ = v70
	var v72 int32
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
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
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
	v47 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v52
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
	F_sequence_close(m, v57, int32(0))
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
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v12)+20))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v72
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
	v87 = v85 - int32(96)
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
	F_ScanKeyInit(m, v87, int32(1), int32(3), int32(184), v75)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_ScanKeyInit(m, v87+int32(48), int32(2), int32(3), int32(184), v74)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v109 = F_systable_beginscan(m, v91, int32(2673), int32(1), int32(0), int32(2), v87)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v111 = F_systable_getnext(m, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	if v111 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = v111
	v120 = v84
	goto L37
L35:
	;
	v144 = v84
	goto L36
L36:
	;
	F_systable_endscan(m, v109)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L45
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+22)))
	v124 = v122 + v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	if v125 != int32(3079) {
		v134 = v120
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v144 = v134
	goto L36
L39:
	;
	v135 = F_systable_getnext(m, v109)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L43
	}
L40:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+24)))
	if v128 != int32(120) {
		v134 = v120
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v132 = F_lappend_oid(m, v120, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v134 = v132
	goto L39
L43:
	;
	if v135 != 0 {
		v116 = v135
		v120 = v134
		goto L37
	} else {
		goto L44
	}
L44:
	;
	goto L38
L45:
	;
	F_sequence_close(m, v91, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	m.G0 = v87 + int32(96)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v155 = int32(0)
	if v144 == v155 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v193 != 0 {
		goto L24
	} else {
		goto L60
	}
L48:
	;
	v193 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v161 <= int32(0) {
		v186 = v155
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v193 = v186
	goto L47
L52:
	;
	v164 = int32(0)
	if v164 < v161 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v167 = v161
	goto L55
L54:
	;
	v167 = v164
	goto L55
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v170 = int32(0)
	goto L56
L56:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168+v170<<(uint(int32(2))%32))))
	v179 = base.B2i32(v178 == v154)
	if v178 == v154 {
		v186 = v179
		goto L51
	} else {
		goto L58
	}
L57:
	;
	v186 = v179
	goto L51
L58:
	;
	v181 = v170 + int32(1)
	if v181 != v167 {
		v170 = v181
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
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	goto L24
}
