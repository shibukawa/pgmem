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
						F_errmsg(m, int32(_a_F_ExecOpenScanRelation_0), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_ExecOpenScanRelation_1), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ExecOpenScanRelation_2), int32(786), int32(_a_F_ExecOpenScanRelation_3))
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
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
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0]))
	v23 = F_AllocSetContextCreateInternal(m, v18, int32(_a_F_RelationBuildRowSecurity_0), int32(0), int32(1024), int32(_a_F_RelationBuildRowSecurity_1))
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = F_MemoryContextStrdup(m, v23, v25+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v28
	v32 = F_MemoryContextAllocZero(m, v23, int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v23
	v37 = F_table_open(m, int32(3256), int32(1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = v13 + int32(-56)
	v41 = int32(3)
	v44 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v40, v41, v41, int32(184), v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v48 = int32(1)
	v51 = F_systable_beginscan(m, v37, int32(3258), v48, int32(0), v48, v40)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L63
	}
L8:
	;
	v53 = F_systable_getnext(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v53 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = v53
	goto L13
L11:
	;
	goto L12
L12:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L44
	}
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
	v70 = F_MemoryContextAllocZero(m, v23, int32(28))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v72 = v67 + v68
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+72)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+4)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+12)) = uint8(v75)
	v79 = F_MemoryContextStrdup(m, v23, v72+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	v85 = v13 + int32(-57)
	v86 = F_heap_getattr_7(m, v59, int32(6), v83, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v88 == int32(1) {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v23
	v94 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v86))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+8)) = v94
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v18
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	v101 = F_heap_getattr_7(m, v59, int32(7), v100, v85)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v103 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v37)+52))
	v125 = F_heap_getattr_7(m, v59, int32(8), v122, v13+int32(-57))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L28
	}
L22:
	;
	v107 = F_text_to_cstring(m, base.I32_wrap_i64(v101))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = int32(0)
	goto L21
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v23
	v111 = F_stringToNode(m, v107)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+16)) = v111
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v18
	F_pfree(m, v107)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L21
L28:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+7)))
	if v127 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	v146 = F_checkExprHasSubLink(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L36
	}
L30:
	;
	v131 = F_text_to_cstring(m, base.I32_wrap_i64(v125))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = int32(0)
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v23
	v135 = F_stringToNode(m, v131)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v135
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v18
	F_pfree(m, v131)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	if v146 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v152 = int32(1)
	goto L39
L38:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v150 = F_checkExprHasSubLink(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)) = uint8(v152)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v23
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v157 = F_lcons(m, v70, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v152 = v150
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v157
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[0])) = v18
	v162 = F_systable_getnext(m, v51)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v162 != 0 {
		v59 = v162
		goto L13
	} else {
		goto L43
	}
L43:
	;
	goto L14
L44:
	;
	F_relation_close(m, v37, int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRowSecurity[1]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	if v186 != v182 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v32
	m.G0 = v15 - int32(-64)
	return
L47:
	;
	if v186 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	if v182 != 0 {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v23)+28))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	if v191 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v190 == int32(0) {
		goto L50
	} else {
		goto L56
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+28)) = v190
	goto L52
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = v190
	goto L52
L56:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+24)) = v196
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v182
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v203
	if v203 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+24)) = v23
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+20)) = v23
	goto L46
L63:
	;
	F_errmsg_internal(m, int32(_a_F_RelationBuildRowSecurity_2), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_RelationBuildRowSecurity_3), int32(265), int32(_a_F_RelationBuildRowSecurity_4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
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
	var v15 int32
	_ = v15
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
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
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
	v15 = v12
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v9)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L35
	}
L7:
	;
	v19 = v15 + int32(19)
	v20 = int32(_a_F_RelationCacheInitFileRemoveInDir_0)
	v24 = m.G0
	v26 = v24 - int32(32)
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v26)+24)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+16)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v26))) = v27
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitFileRemoveInDir[0])))
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
	v104 = F_strlen(m, v19)
	mBase = m.M
	if v103 == v104 {
		goto L28
	} else {
		goto L29
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
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitFileRemoveInDir[1])))
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
		v93 = v19
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v103 = v93 - v19
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
		v93 = v74
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v93 = v91
	goto L22
L26:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v91 = v74 + int32(1)
	if v89 != 0 {
		v74 = v91
		v75 = v89
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(_a_F_RelationCacheInitFileRemoveInDir_1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
	v111 = v7 + int32(16)
	v114 = F_pg_snprintf(m, v111, int32(2048), int32(_a_F_RelationCacheInitFileRemoveInDir_2), v7)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v121 = F_ReadDirExtended(m, v9, l0, int32(15))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	F_unlink_initfile(m, v111, int32(15))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	if v121 != 0 {
		v15 = v121
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L8
L35:
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
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[0]))
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
	v27 = int32(_a_F_RelationDestroyRelation_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+80)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v16
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[1])) = v16
	goto L7
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[1]))
	v25 = v20
	goto L8
L10:
	;
	goto L11
L11:
	;
	v22 = int32(_a_F_RelationDestroyRelation_0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[0])) = v22
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
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[2]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[3])) = v91 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v92+v91<<(uint(int32(2))%32)))) = v52
	goto L21
L27:
	;
	v58 = int32(_a_F_RelationDestroyRelation_1)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[4]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[4])) = v62
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
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[6]))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[3]))
	if v77 < v75 {
		v91 = v77
		v92 = v55
		goto L26
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[6])) = int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[2])) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[4])) = v59
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
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[6])) = v75 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[2])) = v81
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDestroyRelation[3]))
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
	var v27 int32
	_ = v27
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v16 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v14 - int32(-64)
	return v167
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
	v27 = v12 + int32(-56)
	v31 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v27, int32(2), int32(3), int32(184), v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v167 = v20
	goto L1
L7:
	;
	v36 = F_table_open(m, int32(2610), int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(1)
	v42 = F_systable_beginscan(m, v36, int32(2678), v39, int32(0), v39, v27)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v44 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v48 = v2
	v49 = v44
	v52 = v2
	v55 = v2
	v56 = v2
	goto L14
L12:
	;
	v109 = v2
	v113 = v2
	v116 = v2
	v117 = v2
	goto L13
L13:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L33
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+22)))
	v59 = v57 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+21)))
	if v60 != int32(1) {
		v101 = v48
		v102 = v52
		v103 = v55
		v104 = v56
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v109 = v101
	v113 = v102
	v116 = v103
	v117 = v104
	goto L13
L16:
	;
	v105 = F_systable_getnext(m, v42)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L31
	}
L17:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v64 = F_lappend_oid(m, v48, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	if v66 != int32(1) {
		v101 = v64
		v102 = v52
		v103 = v55
		v104 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v71 = F_heap_attisnull(m, v49, int32(21), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	if v71 == int32(0) {
		v101 = v64
		v102 = v52
		v103 = v55
		v104 = v56
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	if v75 != int32(1) {
		v89 = v52
		v90 = v56
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	if v91 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v55
		v104 = v90
		goto L16
	} else {
		goto L28
	}
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+18)))
	if v78 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
	if v82 != int32(112) {
		v89 = v52
		v90 = v56
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+16)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v89 = v88
	v90 = v85 ^ int32(1)
	goto L22
L27:
	;
	goto L26
L28:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+18)))
	if v94 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v55
		v104 = v90
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	if v97 != int32(1) {
		v101 = v64
		v102 = v89
		v103 = v55
		v104 = v90
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v101 = v64
	v102 = v89
	v103 = v100
	v104 = v90
	goto L16
L31:
	;
	if v105 != 0 {
		v48 = v101
		v49 = v105
		v52 = v102
		v55 = v103
		v56 = v104
		goto L14
	} else {
		goto L32
	}
L32:
	;
	goto L15
L33:
	;
	F_relation_close(m, v36, int32(1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_list_sort(m, v109, int32(502))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v126 = int32(_a_F_RelationGetIndexList_0)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexList[0]))
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexList[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexList[0])) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v133 = F_list_copy(m, v109)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v136 = v117 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(v136)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v133
	v140 = int32(0)
	if v136|(base.B2i32(v113 == v140)|base.B2i32(v25 != int32(100))) == v140 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexList[0])) = v127
	F_list_free(m, v132)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L44
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v113
	goto L37
L39:
	;
	goto L40
L40:
	;
	v149 = int32(0)
	if base.B2i32(v116 == v149)|base.B2i32(v25 != int32(105)) == v149 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v116
	goto L37
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L37
L44:
	;
	v167 = v109
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
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
	v23 = v12 + int32(-56)
	F_ScanKeyInit(m, v23, int32(9), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = int32(1)
	v34 = F_systable_beginscan(m, v18, int32(2665), v31, int32(0), v31, v23)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_systable_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v41 = v36
	v48 = v4
	goto L9
L7:
	;
	v145 = v4
	goto L8
L8:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v51 = v49 + v50
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+72)))
	if v52 != int32(110) {
		v132 = v48
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v145 = v132
	goto L8
L11:
	;
	v133 = F_systable_getnext(m, v34)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+106)))
	if v57&int32(1) != 0 {
		v132 = v48
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v60 = F_extractNotNullColumn(m, v41)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
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
	v127 = F_lappend(m, v48, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L29
	}
L19:
	;
	v63 = F_palloc(m, int32(28))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v90 = F_palloc0(m, int32(108))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v67
	v71 = F_pstrdup(m, v51+int32(4))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+20)) = uint8(v73)
	v75 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v75
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)) = uint16(v60)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v71
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+24)) = uint16(v75)
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)) = uint8(v73)
	v85 = v79 ^ v73
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+21)) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v63)+26)) = uint8(v87)
	v124 = v63
	goto L18
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = int64(4294967457)
	v96 = F_pstrdup(m, v51+int32(4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v90)+104)) = int32(-1)
	v100 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+12)) = uint16(v100)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v96
	v104 = F_get_attname(m, l0, v60, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v106 = F_makeString(m, v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v106
	v111 = F_list_make1_impl(m, int32(1), v14)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v113 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+14)) = uint8(v113)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+32)) = v111
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+76)))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+16)) = uint8(v113)
	v120 = v116 ^ v113
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+15)) = uint8(v120)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+106)))
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+17)) = uint8(v122)
	v124 = v90
	goto L18
L29:
	;
	v132 = v127
	goto L11
L30:
	;
	if v133 != 0 {
		v41 = v133
		v48 = v132
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	F_relation_close(m, v18, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v3
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	switch v11 - int32(73) {
	case 0, 32:
		goto L3
	default:
		goto L1
	case 36, 39, 41, 43, 45:
		v16 = v3
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[0]))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	v16 = v15
	goto L2
L4:
	;
	v21 = int32(_a_F_RelationParseRelOptions_0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[1]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[1])) = v25
	v28 = F_CreateTemplateTupleDesc(m, int32(34))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v151 = v18
	goto L6
L6:
	;
	v156 = F_extractRelOptions(m, l1, v151, v16)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L32
	}
L7:
	;
	return
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+4)) = int64(-4294965047)
	v35 = v3
	goto L9
L9:
	;
	v39 = int32(100)
	v40 = v35 * v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	base.MemoryCopy(m, v40+(v28+v41<<(uint(int32(3))%32))+int32(28), v40+int32(_a_F_RelationParseRelOptions_1), v39)
	F_populate_compact_attribute(m, v28, v35)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v58 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v58 < v67 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v55 = v35 + int32(1)
	if v55 != int32(34) {
		v35 = v55
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[0])) = v28
	*(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[1])) = v22
	v151 = v28
	goto L6
L14:
	;
	v71 = v28 + int32(28)
	v78 = v58
	v79 = v67
	v81 = v58
	goto L18
L15:
	;
	v135 = v58
	v142 = v67
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v135
	goto L13
L17:
	;
	v135 = v129
	v142 = v108
	goto L16
L18:
	;
	v87 = v71 + v67<<(uint(int32(3))%32) + v78*int32(100)
	v90 = v71 + v78<<(uint(int32(3))%32)
	if v67 != v79 {
		v108 = v79
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v129 = v67
	goto L17
L20:
	;
	v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+2)))
	if v109 <= int32(0) {
		v129 = v78
		goto L17
	} else {
		goto L28
	}
L21:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+7)))
	if v92 != int32(118) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v108 = v78
	goto L20
L23:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
	if v95 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+6)))
	if v98&int32(6) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v90)+2)))
	if v101 <= int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+90)))
	if v104 != int32(118) {
		v108 = v67
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+90)))
	if v112 == int32(118) {
		v129 = v78
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
	v121 = (v81 + v115 - int32(1)) & (int32(0) - v115)
	if int32(_a_F_RelationParseRelOptions_2) < v121 {
		v129 = v78
		goto L17
	} else {
		goto L30
	}
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v90))) = uint16(v121)
	v127 = v78 + int32(1)
	if v127 != v67 {
		v78 = v127
		v79 = v108
		v81 = v121 + v109
		goto L18
	} else {
		goto L31
	}
L31:
	;
	goto L19
L32:
	;
	if v156 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_RelationParseRelOptions[2]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v165 = F_MemoryContextAlloc(m, v161, int32(base.Ui32(v162)>>(uint(int32(2))%32)))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v165
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v170 = int32(base.Ui32(v168) >> (uint(int32(2)) % 32))
	if v170 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryCopy(m, v165, v156, v170)
	goto L37
L36:
	;
	goto L37
L37:
	;
	F_pfree(m, v156)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L38
	}
L38:
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
	var v51 int64
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[0])))
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
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[6]))
	if v39 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[5]))
	if v32 == int32(0) {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[5])) = int32(0)
	v45 = v32
	goto L5
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[6])) = int32(0)
	v45 = v39
	goto L5
L15:
	;
	v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
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
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[0])))
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
	F_errmsg_internal(m, int32(_a_F_RelationSetNewRelfilenumber_3), v13+int32(16))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_RelationSetNewRelfilenumber_1), int32(3886), int32(_a_F_RelationSetNewRelfilenumber_2))
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
	F_relation_close(m, v48, int32(3))
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
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[1]))
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
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[2]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[2])) = v191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v191<<(uint(int32(2))%32))+uint32(_c_F_RelationSetNewRelfilenumber[3]))) = v194
	goto L55
L57:
	;
	goto L58
L58:
	;
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RelationSetNewRelfilenumber[4])) = uint8(v205)
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
	F_errmsg(m, int32(_a_F_RelationSetNewRelfilenumber_4), int32(0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_RelationSetNewRelfilenumber_1), int32(3791), int32(_a_F_RelationSetNewRelfilenumber_2))
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
	F_errmsg(m, int32(_a_F_RelationSetNewRelfilenumber_5), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_RelationSetNewRelfilenumber_1), int32(3801), int32(_a_F_RelationSetNewRelfilenumber_2))
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
	F_errmsg(m, int32(_a_F_RelationSetNewRelfilenumber_6), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_RelationSetNewRelfilenumber_1), int32(3809), int32(_a_F_RelationSetNewRelfilenumber_2))
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
	F_errmsg_internal(m, int32(_a_F_RelationSetNewRelfilenumber_0), v13)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_RelationSetNewRelfilenumber_1), int32(3820), int32(_a_F_RelationSetNewRelfilenumber_2))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = v4
	v28 = v4
	v29 = v4
	goto L3
L3:
	;
	if int32(0) < v18 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v23 = v20
	v24 = v21
	goto L6
L5:
	;
	v23 = v4
	v24 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v27 = v24
	v28 = v19
	v29 = v23
	goto L3
L7:
	;
	v35 = v27
	v40 = v4
	v41 = v4
	goto L10
L8:
	;
	v103 = v27
	v109 = v4
	goto L9
L9:
	;
	if v103 != 0 {
		goto L34
	} else {
		goto L35
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = l0 + v45<<(uint(int32(3))%32) + v40*int32(100)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+119)))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v103 = v88
	v109 = v90
	goto L9
L12:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v94 = F_lappend(m, v93, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L16
	} else {
		goto L32
	}
L13:
	;
	v56 = F_pstrdup(m, int32(_a_F_buildRelationAliases_0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v35 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	return
L17:
	;
	v58 = F_makeString(m, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v35 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v61 = F_lappend(m, v60, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v88 = v35
	v89 = v58
	v90 = v41 + int32(1)
	goto L12
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v61
	goto L21
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v70 = F_lappend(m, v68, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v84 = F_pstrdup(m, v51+int32(32))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L16
	} else {
		goto L30
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v70
	v74 = v35 + int32(4)
	if base.Ui32(v74) < base.Ui32(v66+v67<<(uint(int32(2))%32)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v80 = v74
	goto L29
L28:
	;
	v80 = int32(0)
	goto L29
L29:
	;
	v88 = v80
	v89 = v69
	v90 = v41
	goto L12
L30:
	;
	v86 = F_makeString(m, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(0)
	v89 = v86
	v90 = v41
	goto L12
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v94
	v98 = v40 + int32(1)
	if v98 != v18 {
		v35 = v88
		v40 = v98
		v41 = v90
		goto L10
	} else {
		goto L33
	}
L33:
	;
	goto L11
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L16
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	m.G0 = v16 + int32(16)
	return
L37:
	;
	F_errcode(m, int32(_a_F_buildRelationAliases_1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v18 - v109
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v120
	F_errmsg(m, int32(_a_F_buildRelationAliases_2), v16)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_buildRelationAliases_3), int32(1280), int32(_a_F_buildRelationAliases_4))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
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
							F_errmsg(m, int32(_a_F_check_relation_updatable_0), v7+int32(16))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_relation_updatable_1), int32(2798), int32(_a_F_check_relation_updatable_2))
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
							F_errmsg(m, int32(_a_F_check_relation_updatable_3), v7)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_relation_updatable_1), int32(2807), int32(_a_F_check_relation_updatable_2))
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
func F_fetch_relation_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
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
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_fetch_relation_list[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v17
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_fetch_relation_list[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v20
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_relation_list[2]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
	v25 = m.T0[v24].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = v14 + int32(96)
	F_initStringInfo(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = v14 - int32(-64)
	F_initStringInfo(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_GetPublicationsStr(m, l1, v34, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(_a_F_fetch_relation_list_0) < v25 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v44 = int32(4)
	goto L8
L7:
	;
	v44 = int32(3)
	goto L8
L8:
	;
	if int32(_a_F_fetch_relation_list_1) <= v25 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	F_pfree(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v74
	F_appendStringInfo(m, v14+int32(96), v72, v14+int32(32))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = int32(22)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v49
	F_appendStringInfo(m, v30, int32(_a_F_fetch_relation_list_2), v14+int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+92)) = int32(1003)
	v62 = v14 + int32(96)
	F_appendStringInfoString(m, v62, int32(_a_F_fetch_relation_list_3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	if base.Ui32(int32(_a_F_fetch_relation_list_4)) < base.Ui32(v25) {
		v72 = int32(_a_F_fetch_relation_list_5)
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	v66 = int32(_a_F_fetch_relation_list_6)
	if v25 < int32(_a_F_fetch_relation_list_7) {
		v72 = v66
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_appendStringInfoString(m, v62, int32(_a_F_fetch_relation_list_8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v72 = v66
	goto L10
L19:
	;
	goto L9
L20:
	;
	v87 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_relation_list[2]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+60))
	v94 = m.T0[v93].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v88, v44, v14+int32(80))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	F_pfree(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v99 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	v104 = F_MakeSingleTupleTableSlot(m, v102, int32(_a_F_fetch_relation_list_9))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L79
	}
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v109 = F_tuplestore_gettupleslot(m, v106, int32(1), int32(0), v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v109 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v114 = v87
	goto L31
L29:
	;
	v237 = v87
	goto L30
L30:
	;
	F_ExecDropSingleTupleTableSlot(m, v104)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L65
	}
L31:
	;
	v125 = F_palloc(m, int32(8))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v237 = v225
	goto L30
L33:
	;
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+6)))
	if v127 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	m.T0[v132].(func(*base.Module, int32, int32))(m, v104, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = F_text_to_cstring(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+6)))
	if v139 <= int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+16))
	m.T0[v144].(func(*base.Module, int32, int32))(m, v104, int32(2))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	v149 = F_text_to_cstring(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v104)+6)))
	if v151 <= int32(2) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	m.T0[v156].(func(*base.Module, int32, int32))(m, v104, int32(3))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+16)))
	v162 = F_makeRangeVar(m, v137, v149, int32(-1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v162
	if base.B2i32(v160 == int32(83))|base.B2i32(v25 < int32(_a_F_fetch_relation_list_7)) != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v225 = F_lappend(m, v114, v125)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L61
	}
L50:
	;
	v169 = int32(0)
	if v114 == v169 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v175 = v169
	goto L52
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v183 <= v175 {
		goto L49
	} else {
		goto L54
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v175<<(uint(int32(2))%32)+v189)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v193 = F_equal(m, v192, v162)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v193 == int32(0) {
		v175 = v175 + int32(1)
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v137
	F_errmsg(m, int32(_a_F_fetch_relation_list_10), v14)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_fetch_relation_list_11), int32(3448), int32(_a_F_fetch_relation_list_12))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	m.T0[v228].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v234 = F_tuplestore_gettupleslot(m, v231, int32(1), int32(0), v104)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v234 != 0 {
		v114 = v225
		goto L31
	} else {
		goto L64
	}
L64:
	;
	goto L32
L65:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	if v249 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	F_pfree(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v252 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	F_tuplestore_end(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
	if v255 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L72
L74:
	;
	F_FreeTupleDesc(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_pfree(m, v94)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	m.G0 = v14 + int32(112)
	return v237
L79:
	;
	F_errcode(m, int32(100663808))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v271
	F_errmsg(m, int32(_a_F_fetch_relation_list_13), v14+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_fetch_relation_list_11), int32(3420), int32(_a_F_fetch_relation_list_12))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
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
	v19 = v8 + int32(-56)
	F_ScanKeyInit(m, v19, int32(9), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(1)
	v30 = F_systable_beginscan(m, v14, int32(2665), v27, int32(0), v27, v19)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_systable_endscan(m, v30)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v32 = F_systable_getnext(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v32 == int32(0) {
		v69 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = v32
	goto L8
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+72)))
	v48 = v46 - int32(112)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v48))|base.B2i32(int32(1)<<(uint(v48)%32)&int32(289) == int32(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v69 = v3
	goto L4
L10:
	;
	v61 = F_systable_getnext(m, v30)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+88))
	if v58 != l1 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v69 = v60
	goto L4
L13:
	;
	if v61 != 0 {
		v36 = v61
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	F_relation_close(m, v14, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	m.G0 = v10 - int32(-64)
	return v69
}
