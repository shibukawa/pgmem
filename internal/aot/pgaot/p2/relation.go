package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecOpenScanRelation(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = F_ExecGetRangeTableRelation(m, l0, l1)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if l2&int32(65) != 0 {
			m.G0 = v7 + int32(16)
			return v9
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+129)))
			if v16 != 0 {
				m.G0 = v7 + int32(16)
				return v9
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v24 + int32(4)
						F_errmsg(m, int32(432728), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(611615), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(477261), int32(760), int32(255012))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
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
func F_GetRelationIdentityOrPK(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = F_RelationGetReplicaIndex(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			v10 = v3
			return v10
		} else {
			v8 = F_RelationGetPrimaryKeyIndex(m, l0, int32(0))
			v9 = m.ExcPending
			if v9 != 0 {
				return int32(0)
			} else {
				v10 = v8
				return v10
			}
		}
	}
}
func F_RelationBuildRowSecurity(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(200092), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v27 = F_MemoryContextStrdup(m, v22, v24+int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v27
	goto L4
L4:
	;
	v31 = F_MemoryContextAllocZero(m, v22, int32(8))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v22
	v36 = F_table_open(m, int32(3256), int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = int32(3)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v12+int32(-48), v40, v40, int32(184), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v47 = int32(1)
	v52 = F_systable_beginscan(m, v36, int32(3258), v47, int32(0), v47, v12+int32(-48))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L64
	}
L9:
	;
	v54 = F_systable_getnext(m, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v57 = v54
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_systable_endscan(m, v52)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v70 = F_MemoryContextAllocZero(m, v22, int32(28))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v72 = v67 + v68
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+12)) = uint8(v75)
	v79 = F_MemoryContextStrdup(m, v22, v72+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v86 = F_heap_getattr_7(m, v57, int32(6), v83, v12+int32(-49))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v88 == int32(1) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v93 = F_pg_detoast_datum_copy(m, v86)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v93
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v102 = F_heap_getattr_7(m, v57, int32(7), v99, v12+int32(-49))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v104 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v125 = F_heap_getattr_7(m, v57, int32(8), v122, v12+int32(-49))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L29
	}
L23:
	;
	v107 = F_text_to_cstring(m, v102)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = int32(0)
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v111 = F_stringToNode(m, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v111
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	F_pfree(m, v107)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v127 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v146 = F_checkExprHasSubLink(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L37
	}
L31:
	;
	v130 = F_text_to_cstring(m, v125)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = int32(0)
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v134 = F_stringToNode(m, v130)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v134
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	F_pfree(m, v130)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	if v146 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v151 = int32(1)
	goto L40
L39:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v149 = F_checkExprHasSubLink(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)) = uint8(v151)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v22
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v156 = F_lcons(m, v70, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v151 = v149
	goto L40
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v156
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
	v161 = F_systable_getnext(m, v52)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v161 != 0 {
		v57 = v161
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L15
L45:
	;
	F_sequence_close(m, v36, int32(1))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v184 != v180 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v31
	m.G0 = v14 - int32(-64)
	return
L48:
	;
	if v184 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	if v180 != 0 {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v189 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v188 == int32(0) {
		goto L51
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+28)) = v188
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184)+20)) = v188
	goto L53
L57:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+24)) = v194
	goto L51
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v180
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v201
	if v201 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
	goto L50
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+24)) = v22
	goto L63
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = v22
	goto L47
L64:
	;
	F_errmsg_internal(m, int32(157539), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(475739), int32(265), int32(10186))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCacheInitFileRemoveInDir(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	v5 = m.G0
	v7 = v5 - int32(2064)
	m.G0 = v7
	v9 = F_AllocateDir(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = F_ReadDirExtended(m, v9, l0, int32(15))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v12
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v9)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L54
	}
L7:
	;
	v19 = v16 + int32(19)
	v20 = int32(529931)
	v24 = m.G0
	v26 = v24 - int32(32)
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1176])))
	if v35 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	if v19&int32(3) == int32(0) {
		v127 = v19
		goto L32
	} else {
		goto L33
	}
L10:
	;
	v103 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1177])))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v43 = v19
	goto L16
L14:
	;
	goto L15
L15:
	;
	v53 = v20
	v54 = v35
	goto L19
L16:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v49 == v35 {
		v43 = v43 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v103 = v43 - v19
	goto L9
L18:
	;
	goto L17
L19:
	;
	v61 = v26 + int32(base.Ui32(v54)>>(uint(int32(3))%32))&int32(28)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v62 | v63<<(uint(v54)%32)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v67 != 0 {
		v53 = v53 + v63
		v54 = v67
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v70 == int32(0) {
		v95 = v19
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v103 = v95 - v19
	goto L9
L23:
	;
	v74 = v19
	v75 = v70
	goto L24
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v26+int32(base.Ui32(v75)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v83)>>(uint(v75)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v95 = v91
	goto L22
L26:
	;
	v95 = v74
	goto L22
L27:
	;
	goto L28
L28:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v91 = v74 + int32(1)
	if v89 != 0 {
		v74 = v91
		v75 = v89
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	if v103 == v160 {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v160 = v152 - v19
	goto L30
L32:
	;
	v131 = v127
	goto L41
L33:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v111 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v160 = int32(0)
	goto L30
L35:
	;
	goto L36
L36:
	;
	v116 = v19
	goto L37
L37:
	;
	v120 = v116 + int32(1)
	if v120&int32(3) == int32(0) {
		v127 = v120
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v152 = v120
	goto L31
L39:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v125 != 0 {
		v116 = v120
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v140 = int32(-2139062144)
	if (int32(16843008)-v137|v137)&v140 == v140 {
		v131 = v131 + int32(4)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v146 = v131
	goto L44
L43:
	;
	goto L42
L44:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 != 0 {
		v146 = v146 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v152 = v146
	goto L31
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(95907)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v170 = F_pg_snprintf(m, v7+int32(16), int32(2048), int32(169616), v7)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v178 = F_ReadDirExtended(m, v9, l0, int32(15))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	F_unlink_initfile(m, v7+int32(16), int32(15))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if v178 != 0 {
		v16 = v178
		goto L7
	} else {
		goto L53
	}
L53:
	;
	goto L8
L54:
	;
	m.G0 = v7 + int32(2064)
	return
}
func F_RelationDestroyRelation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+72))
	v11 = v9 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+72)) = v11
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_smgrclose(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v16 = v6 + int32(76)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+76)) = v25
	v27 = int32(4378696)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v16
	*(*int32)(unsafe.Add(mBase, _consts[1165])) = v16
	goto L7
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[1165]))
	v25 = v20
	goto L8
L10:
	;
	goto L11
L11:
	;
	v22 = int32(4378696)
	*(*int32)(unsafe.Add(mBase, _consts[1164])) = v22
	v25 = v22
	goto L8
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	goto L3
L14:
	;
	v40 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v40
	goto L16
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v50 = v48 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v50
	if v50 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	F_FreeTriggerDesc(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L34
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, _consts[1166]))
	if v55 == v53 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	goto L25
L25:
	;
	F_FreeTupleDesc(m, v52)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L33
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1167])) = v91 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92+v91<<(uint(int32(2))%32)))) = v52
	goto L21
L27:
	;
	v58 = int32(4455216)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v62 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v62
	v65 = F_palloc(m, int32(64))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[1168]))
	v77 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	if v77 < v75 {
		v91 = v77
		v92 = v55
		goto L26
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1168])) = int32(16)
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v65
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
	v91 = v53
	v92 = v65
	goto L26
L31:
	;
	v81 = F_repalloc(m, v55, v75<<(uint(int32(3))%32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1168])) = v75 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _consts[1166])) = v81
	v90 = *(*int32)(unsafe.Add(mBase, _consts[1167]))
	v91 = v90
	v92 = v81
	goto L26
L33:
	;
	goto L21
L34:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	F_list_free_deep(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L35
	}
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_list_free(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	F_list_free(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_bms_free(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L38
	}
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_bms_free(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_bms_free(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	F_bms_free(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_bms_free(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v135 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L12
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v138 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v141 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v144 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L53
L55:
	;
	F_pfree(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+260))
	if v147 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	F_pfree(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L12
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v150 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	F_MemoryContextDelete(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v153 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	F_MemoryContextDelete(m, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v156 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L69
L71:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	F_MemoryContextDelete(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v160 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L73
L75:
	;
	F_MemoryContextDelete(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v163 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	goto L77
L79:
	;
	F_MemoryContextDelete(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L12
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v166 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_MemoryContextDelete(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v169 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_MemoryContextDelete(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	F_pfree(m, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L91
	}
L90:
	;
	goto L89
L91:
	;
	return
}
func F_RelationGetIndexList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 + int32(48)
	return v162
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v20 = F_list_copy(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+130)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v14, int32(2), int32(3), int32(184), v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v162 = v20
	goto L1
L7:
	;
	v34 = F_table_open(m, int32(2610), int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L9
	}
L8:
	;
	F_systable_endscan(m, v40)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L34
	}
L9:
	;
	v37 = int32(1)
	v40 = F_systable_beginscan(m, v34, int32(2678), v37, int32(0), v37, v14)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v42 = F_systable_getnext(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v109 = v2
	v112 = v2
	v113 = v2
	v114 = v2
	goto L8
L13:
	;
	goto L14
L14:
	;
	v48 = v2
	v50 = v42
	v51 = v2
	v52 = v2
	v53 = v2
	goto L15
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v59 = v57 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+21)))
	if v60 != int32(1) {
		v101 = v48
		v102 = v51
		v103 = v52
		v104 = v53
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v109 = v101
	v112 = v102
	v113 = v103
	v114 = v104
	goto L8
L17:
	;
	v105 = F_systable_getnext(m, v40)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L32
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v64 = F_lappend_oid(m, v48, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	if v66 != int32(1) {
		v101 = v64
		v102 = v51
		v103 = v52
		v104 = v53
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v71 = F_heap_attisnull(m, v50, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	if v71 == int32(0) {
		v101 = v64
		v102 = v51
		v103 = v52
		v104 = v53
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	if v75 != int32(1) {
		v89 = v51
		v90 = v53
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	if v91 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v52
		v104 = v90
		goto L17
	} else {
		goto L29
	}
L24:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+18)))
	if v78 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
	if v82 != int32(112) {
		v89 = v51
		v90 = v53
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v89 = v88
	v90 = v85 ^ int32(1)
	goto L23
L28:
	;
	goto L27
L29:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+18)))
	if v94 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v52
		v104 = v90
		goto L17
	} else {
		goto L30
	}
L30:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	if v97 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v52
		v104 = v90
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v101 = v64
	v102 = v89
	v103 = v100
	v104 = v90
	goto L17
L32:
	;
	if v105 != 0 {
		v48 = v101
		v50 = v105
		v51 = v102
		v52 = v103
		v53 = v104
		goto L15
	} else {
		goto L33
	}
L33:
	;
	goto L16
L34:
	;
	F_sequence_close(m, v34, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_list_sort(m, v109, int32(467))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v126 = int32(4455216)
	v127 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v133 = F_list_copy(m, v109)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v136 = v114 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(v136)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v133
	if v25 != int32(100) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v154)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v127
	F_list_free(m, v132)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L46
	}
L39:
	;
	if v25 != int32(105) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if v112 == int32(0) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v114&int32(1) != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v112
	goto L38
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L38
L44:
	;
	if v113 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v113
	goto L38
L46:
	;
	v162 = v109
	goto L1
}
func F_RelationGetNotNullConstraints(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v18 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v12+int32(-48), int32(9), int32(3), int32(184), l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = int32(1)
	v35 = F_systable_beginscan(m, v18, int32(2665), v30, int32(0), v30, v12+int32(-48))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v37 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = v37
	v46 = v4
	goto L9
L7:
	;
	v145 = v4
	goto L8
L8:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+22)))
	v52 = v50 + v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+72)))
	if v53 != int32(110) {
		v134 = v46
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v145 = v134
	goto L8
L11:
	;
	v136 = F_systable_getnext(m, v35)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L30
	}
L12:
	;
	if l2 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+106)))
	if v58&int32(1) != 0 {
		v134 = v46
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v61 = F_extractNotNullColumn(m, v42)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	if l1 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v130 = F_lappend(m, v46, v127)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L29
	}
L19:
	;
	v64 = F_palloc(m, int32(28))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v91 = F_palloc0(m, int32(108))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v68
	v72 = F_pstrdup(m, v52+int32(4))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+20)) = uint8(v74)
	v76 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v76
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+12)) = uint16(v61)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v72
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+24)) = uint16(v76)
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+22)) = uint8(v74)
	v86 = v80 ^ v74
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+21)) = uint8(v86)
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v64)+26)) = uint8(v88)
	v127 = v64
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v91))) = int64(4294967457)
	v97 = F_pstrdup(m, v52+int32(4))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+104)) = int32(-1)
	v101 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+12)) = uint16(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = v97
	v105 = F_get_attname(m, l0, v61, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v107 = F_makeString(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v107
	v114 = F_list_make1_impl(m, int32(1), v12+int32(-56))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v116 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+14)) = uint8(v116)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+32)) = v114
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+16)) = uint8(v116)
	v123 = v119 ^ v116
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+15)) = uint8(v123)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+17)) = uint8(v125)
	v127 = v91
	goto L18
L29:
	;
	v134 = v130
	goto L11
L30:
	;
	if v136 != 0 {
		v42 = v136
		v46 = v134
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	F_sequence_close(m, v18, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v14 - int32(-64)
	return v145
}
func F_RelationIsVisible(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_RelationIsVisibleExt(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_RelationParseRelOptions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v3
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
	switch v12 - int32(73) {
	case 0, 32:
		goto L3
	default:
		goto L1
	case 36, 39, 41, 43, 45:
		v17 = v3
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1163]))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+72))
	v17 = v16
	goto L2
L4:
	;
	v22 = int32(4455216)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v26 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
	v29 = F_CreateTemplateTupleDesc(m, int32(34))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v69 = v19
	goto L6
L6:
	;
	v75 = F_extractRelOptions(m, l1, v69, v17)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L17
	}
L7:
	;
	return
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+4)) = int64(-4294965047)
	v38 = v3
	goto L9
L9:
	;
	v43 = int32(100)
	v44 = v38 * v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	goto L12
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1163])) = v29
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	v69 = v29
	goto L6
L11:
	;
	F_populate_compact_attribute(m, v29, v38)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L15
	}
L12:
	;
	v53 = F__emscripten_memcpy_bulkmem(m, v44+(v29+int32(20)+v45<<(uint(int32(4))%32)), v44+int32(1712384), v43)
	mBase = m.M
	goto L14
L14:
	;
	goto L11
L15:
	;
	v58 = v38 + int32(1)
	if v58 != int32(34) {
		v38 = v58
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v75 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v84 = F_MemoryContextAlloc(m, v80, int32(base.Ui32(v81)>>(uint(int32(2))%32)))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v89 = int32(base.Ui32(v87) >> (uint(int32(2)) % 32))
	if v89 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_pfree(m, v75)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L24
	}
L21:
	;
	v90 = F__emscripten_memcpy_bulkmem(m, v84, v75, v89)
	mBase = m.M
	goto L23
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L1
}
func F_RelationSetNewRelfilenumber(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	v2 = l1
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v3
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[293])))
	if v21 == v3 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L71
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L9
	} else {
		goto L67
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L9
	} else {
		goto L63
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L59
	}
L5:
	;
	v48 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L15
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
	v26 = F_GetNewRelFileNumber(m, v24, int32(0), v2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+119)))
	switch v28 - int32(105) {
	case 0:
		goto L12
	default:
		goto L2
	case 9:
		goto L11
	}
L9:
	;
	return
L10:
	;
	v45 = v26
	goto L5
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[1172]))
	if v39 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[1171]))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1171])) = int32(0)
	v45 = v32
	goto L5
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1172])) = int32(0)
	v45 = v39
	goto L5
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v52 = F_SearchSysCacheLockedCopy1(m, int32(57), v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	if v52 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+92)) = uint16(v56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[293])))
	if v63 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v88
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+119)))
	switch v92 - int32(83) {
	case 0, 22:
		goto L29
	default:
		goto L28
	case 26, 31, 33:
		goto L27
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v69
	v73 = F_smgropen(m, v13+int32(48), v66)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_RelationDropStorage(m, l0)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v73
	F_smgrdounlinkall(m, v13-int32(-64), int32(1), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v13)+64))
	F_smgrclose(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L18
L26:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+119)))
	switch v135 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L38
	default:
		goto L37
	}
L27:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+112))
	m.T0[v131].(func(*base.Module, int32, int32, int32, int32, int32))(m, l0, v13-int32(-64), v2, v13+int32(80), v13+int32(84))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L35
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L32
	}
L29:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v97
	v102 = F_RelationCreateStorage(m, v13+int32(32), v2, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_smgrclose(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v110 + int32(4)
	F_errmsg_internal(m, int32(391384), v13+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(482315), int32(3890), int32(219893))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	goto L26
L36:
	;
	F_UnlockTuple(m, v48, v13+int32(88), int32(7))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L47
	}
L37:
	;
	v149 = v60 + v61
	*(*int32)(unsafe.Add(mBase, uint32(v149)+88)) = v45
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+119)))
	if v152 != int32(83) {
		goto L43
	} else {
		goto L44
	}
L38:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134)+88))
	if v138 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v139 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+117)))
	F_RelationMapUpdateMap(m, v141, v45, v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L41
	}
L41:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	goto L36
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v149)+104)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v149)+96)) = int64(-4647714815446351872)
	goto L45
L44:
	;
	goto L45
L45:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+136)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+84))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+118)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v149)+140)) = v161
	F_CatalogTupleUpdate(m, v48, v13+int32(88), v52)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	goto L36
L47:
	;
	F_pfree(m, v52)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_sequence_close(m, v48, int32(3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+8))
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v186 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v184
	goto L54
L53:
	;
	goto L54
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	if v191 <= int32(31) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	m.G0 = v13 + int32(96)
	return
L56:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _consts[283])) = v191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v191<<(uint(int32(2))%32))+uint32(_consts[284]))) = v194
	goto L55
L57:
	;
	goto L58
L58:
	;
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[285])) = uint8(v205)
	goto L55
L59:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(398197), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(482315), int32(3795), int32(219893))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(398321), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(482315), int32(3805), int32(219893))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(398133), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(482315), int32(3813), int32(219893))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v263
	F_errmsg_internal(m, int32(44601), v13)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(482315), int32(3824), int32(219893))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L9
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_buildRelationAliases(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if int32(0) < v19 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v32 = v4
	v33 = v4
	v34 = v4
	goto L1
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v32 = v29
	v33 = v22
	v34 = v28
	goto L1
L6:
	;
	v28 = v4
	v29 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v28 = v26
	v29 = v27
	goto L5
L9:
	;
	v42 = v32
	v46 = v4
	v47 = v4
	goto L12
L10:
	;
	v111 = v32
	v116 = v4
	goto L11
L11:
	;
	if v111 != 0 {
		goto L36
	} else {
		goto L37
	}
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = l0 + int32(20) + v53<<(uint(int32(4))%32) + v46*int32(100)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+91)))
	if v60 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v111 = v96
	v116 = v98
	goto L11
L14:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v102 = F_lappend(m, v101, v97)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L34
	}
L15:
	;
	v64 = F_pstrdup(m, int32(722455))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v42 != 0 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	return
L19:
	;
	v66 = F_makeString(m, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if v42 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v69 = F_lappend(m, v68, v66)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L18
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v96 = v42
	v97 = v66
	v98 = v47 + int32(1)
	goto L14
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v69
	goto L23
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v78 = F_lappend(m, v76, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L18
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v92 = F_pstrdup(m, v59+int32(4))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L18
	} else {
		goto L32
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v78
	v82 = v42 + int32(4)
	if base.Ui32(v82) < base.Ui32(v74+v75<<(uint(int32(2))%32)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = v82
	goto L31
L30:
	;
	v88 = int32(0)
	goto L31
L31:
	;
	v96 = v88
	v97 = v77
	v98 = v47
	goto L14
L32:
	;
	v94 = F_makeString(m, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	v96 = int32(0)
	v97 = v94
	v98 = v47
	goto L14
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v102
	v106 = v46 + int32(1)
	if v106 != v19 {
		v42 = v96
		v46 = v106
		v47 = v98
		goto L12
	} else {
		goto L35
	}
L35:
	;
	goto L13
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	m.G0 = v17 + int32(16)
	return
L39:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v19 - v116
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v129
	F_errmsg(m, int32(441759), v17)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(479537), int32(1252), int32(154121))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_check_relation_updatable(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	if v11 == int32(112) {
		m.G0 = v7 + int32(32)
		return
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v14 != 0 {
			m.G0 = v7 + int32(32)
			return
		} else {
			v15 = F_GetRelationIdentityOrPK(m, v9)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						if v15 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v24
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v25
							F_errmsg(m, int32(656403), v7+int32(16))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(478658), int32(2527), int32(378001))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v24
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v25
							F_errmsg(m, int32(515687), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(478658), int32(2536), int32(378001))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
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
func F_get_relation_idx_constraint_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v10, int32(9), int32(3), int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = int32(1)
	v27 = F_systable_beginscan(m, v14, int32(2665), v24, int32(0), v24, v10)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	v29 = F_systable_getnext(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v29 == int32(0) {
		v63 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v33 = v29
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v42 = v40 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+72)))
	v45 = v43 - int32(112)
	if base.Ui32(int32(8)) < base.Ui32(v45) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v63 = v3
	goto L4
L10:
	;
	v57 = F_systable_getnext(m, v27)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	if int32(1)<<(uint(v45)%32)&int32(289) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42)+88))
	if v54 != l1 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v63 = v56
	goto L4
L14:
	;
	if v57 != 0 {
		v33 = v57
		goto L8
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	F_sequence_close(m, v14, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	m.G0 = v10 + int32(48)
	return v63
}
func F_relation_has_unique_index_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	v7 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v15 == v7 {
		v427 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v427
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v18 == int32(0) {
		v63 = l2
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v63|l3 == int32(0) {
		v427 = v7
		goto L1
	} else {
		goto L17
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v21 <= int32(0) {
		v63 = l2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v26 = l2
	v31 = v7
	goto L6
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v31<<(uint(int32(2))%32))))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	if v43 == int32(0) {
		v56 = v26
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v63 = v56
	goto L3
L8:
	;
	v58 = v31 + int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v58 < v59 {
		v26 = v56
		v31 = v58
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)+44))
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+48))
	if v47 != 0 {
		v56 = v26
		goto L8
	} else {
		goto L13
	}
L11:
	;
	v50 = int32(1)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v42)+120)) = uint8(v50)
	v52 = F_lappend(m, v26, v42)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v50 = int32(0)
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	v56 = v52
	goto L8
L16:
	;
	goto L7
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v78 == int32(0) {
		v427 = v7
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 <= int32(0) {
		v427 = v7
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v94 = v7
	goto L20
L20:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+v94<<(uint(int32(2))%32))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+101)))
	if v103 != int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v427 = int32(0)
	goto L1
L22:
	;
	v413 = v94 + int32(1)
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v413 < v414 {
		v94 = v413
		goto L20
	} else {
		goto L105
	}
L23:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+103)))
	if v106 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)+88))
	if v109 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
	if v110 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v378 != v383 {
		goto L22
	} else {
		goto L103
	}
L27:
	;
	v113 = int32(0)
	v378 = v113
	v383 = v110
	v388 = v113
	goto L26
L28:
	;
	goto L29
L29:
	;
	v115 = int32(0)
	v118 = v115
	v128 = v115
	goto L30
L30:
	;
	if v63 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v378 = v374
	v383 = v375
	v388 = v370
	goto L26
L32:
	;
	v374 = v118 + int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
	if v374 < v375 {
		v118 = v374
		v128 = v370
		goto L30
	} else {
		goto L102
	}
L33:
	;
	v313 = int32(0)
	goto L86
L34:
	;
	v133 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v134 <= v133 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v143 = v133
	goto L36
L36:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v152 = int32(2)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v143<<(uint(v152)%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+96))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v102)+52))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v118<<(uint(v152)%32))))
	v162 = int32(0)
	if v156 == v162 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L33
L38:
	;
	v289 = v143 + int32(1)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v289 < v290 {
		v143 = v289
		goto L36
	} else {
		goto L85
	}
L39:
	;
	if v200 == int32(0) {
		goto L38
	} else {
		goto L52
	}
L40:
	;
	v200 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v168 <= int32(0) {
		v193 = v162
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v200 = v193
	goto L39
L44:
	;
	v171 = int32(0)
	if v171 < v168 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v174 = v168
	goto L47
L46:
	;
	v174 = v171
	goto L47
L47:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v177 = int32(0)
	goto L48
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v175+v177<<(uint(int32(2))%32))))
	v186 = base.B2i32(v185 == v161)
	if v185 == v161 {
		v193 = v186
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v193 = v186
	goto L43
L50:
	;
	v188 = v177 + int32(1)
	if v188 != v174 {
		v177 = v188
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+28))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+120)))
	if v205 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v222 = F_match_index_to_operand(m, v221, v118, v102)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L14
	} else {
		goto L62
	}
L54:
	;
	v208 = int32(0)
	if v204 == v208 {
		v221 = v208
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v204 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	if v211 < int32(2) {
		v221 = v208
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v221 = v215
	goto L53
L59:
	;
	v221 = int32(0)
	goto L53
L60:
	;
	goto L61
L61:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v221 = v220
	goto L53
L62:
	;
	if v222 == int32(0) {
		goto L38
	} else {
		goto L63
	}
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	if v226 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v273 != int32(1) {
		v370 = v128
		goto L32
	} else {
		goto L80
	}
L65:
	;
	v273 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v235 = int32(1)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v236 <= v235 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v239 = v235
	goto L70
L69:
	;
	v239 = v236
	goto L70
L70:
	;
	v242 = int32(0)
	v244 = v242
	v245 = v242
	goto L71
L71:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v226+int32(8)+v244<<(uint(int32(2))%32))))
	if v253 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v273 = v266
	goto L64
L73:
	;
	goto L72
L74:
	;
	v254 = int32(2)
	if v245 != 0 {
		v266 = v254
		goto L73
	} else {
		goto L77
	}
L75:
	;
	v259 = v245
	goto L76
L76:
	;
	v262 = v244 + int32(1)
	if v262 != v239 {
		v244 = v262
		v245 = v259
		goto L71
	} else {
		goto L79
	}
L77:
	;
	v255 = int32(1)
	if base.Ui32(v255) < base.Ui32(base.I32_popcnt(v253)) {
		v266 = v254
		goto L73
	} else {
		goto L78
	}
L78:
	;
	v259 = v255
	goto L76
L79:
	;
	v266 = v259
	goto L73
L80:
	;
	v276 = int32(4455216)
	v277 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v279
	if l5 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v281 = F_lappend(m, v128, v155)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L14
	} else {
		goto L84
	}
L82:
	;
	v283 = v128
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v277
	v370 = v283
	goto L32
L84:
	;
	v283 = v281
	goto L83
L85:
	;
	goto L37
L86:
	;
	v321 = int32(0)
	if l3 == v321 {
		v331 = v321
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v102)+40))
	v378 = v118
	v383 = v358
	v388 = v128
	goto L26
L88:
	;
	if l4 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v325 <= v313 {
		v331 = int32(0)
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v331 = v327 + v313<<(uint(int32(2))%32)
	goto L88
L91:
	;
	goto L87
L92:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v334 <= v313 {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	if v331 == int32(0) {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v341 = v338 + v313<<(uint(int32(2))%32)
	if v341 == int32(0) {
		goto L91
	} else {
		goto L95
	}
L95:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v346 = F_match_index_to_operand(m, v345, v118, v102)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	if v346 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v102)+52))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348+v118<<(uint(int32(2))%32))))
	v353 = F_op_in_opfamily(m, v344, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L14
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v313 = v313 + int32(1)
	goto L86
L100:
	;
	if v353 != 0 {
		v370 = v128
		goto L32
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	goto L31
L103:
	;
	if l5 == int32(0) {
		v427 = int32(1)
		goto L1
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v388
	return int32(1)
L105:
	;
	goto L21
}
