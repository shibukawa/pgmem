package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RangeVarCallbackForDropRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
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
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v18 = base.B2i32(l1 == l2)
	if l1 == l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_UnlockRelationOid(m, v19, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v24 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	goto L5
L8:
	;
	F_UnlockRelationOid(m, v24, int32(8))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(0)
	goto L1
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L85
	}
L11:
	;
	m.G0 = v15 + int32(48)
	return
L12:
	;
	v36 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v36 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+22)))
	v42 = v40 + v41
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+131)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+16)) = uint8(v44)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+17)) = uint8(v46)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v49 == int32(112) {
		v56 = int32(114)
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3))))
	if v56 != v57 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v49 == int32(73) {
		v56 = int32(105)
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v56 = base.I32_extend8_s(v49)
	goto L15
L18:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v60 = int32(_a_F_RangeVarCallbackForDropRelation_0)
	v67 = v60
	goto L21
L19:
	;
	goto L20
L20:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForDropRelation[0]))
	v126 = F_object_ownercheck(m, int32(1259), l1, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L42
	}
L21:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v76 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v82 = v60
	goto L27
L23:
	;
	if v57&int32(255) != v76 {
		v67 = v67 + int32(24)
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L25
L27:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v93 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L6
	} else {
		goto L33
	}
L29:
	;
	if v93 != v49 {
		v82 = v82 + int32(24)
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L31
L33:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v59
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	F_errmsg(m, v106, v15+int32(32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if v93 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v111
	F_errhint(m, int32(_a_F_RangeVarCallbackForDropRelation_1), v15+int32(16))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForDropRelation_2), int32(1529), int32(_a_F_RangeVarCallbackForDropRelation_3))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v151 = int32(0)
	v153 = int32(1)
	if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForDropRelation_4)) {
		v161 = v153
		goto L56
	} else {
		goto L57
	}
L42:
	;
	if v126 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForDropRelation[0]))
	v132 = F_object_ownercheck(m, int32(2615), v129, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	if v132 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v135 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+119)))
	switch v135 - int32(73) {
	case 0, 32:
		v145 = int32(20)
		goto L47
	default:
		goto L48
	case 10:
		goto L52
	case 29:
		goto L49
	case 36:
		goto L50
	case 45:
		goto L51
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_aclcheck_error(m, int32(2), v147, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L53
	}
L47:
	;
	v147 = v145
	goto L46
L48:
	;
	v145 = int32(41)
	goto L47
L49:
	;
	v147 = int32(18)
	goto L46
L50:
	;
	v147 = int32(23)
	goto L46
L51:
	;
	v147 = int32(51)
	goto L46
L52:
	;
	v147 = int32(37)
	goto L46
L53:
	;
	goto L41
L54:
	;
	if v182&int32(1) != 0 {
		goto L67
	} else {
		goto L68
	}
L55:
	;
	if v161 == int32(0) {
		v182 = v151
		goto L54
	} else {
		goto L59
	}
L56:
	;
	goto L55
L57:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	if v156 == int32(99) {
		v161 = v153
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v159 = F_isTempToastNamespace(m, v156)
	mBase = m.M
	v161 = v159
	goto L56
L59:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
	if v164 != int32(105) {
		v182 = v151
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v168 = F_SearchSysCache1(m, int32(34), l1)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	if v168 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+22)))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174+v175)+18)))
	F_ReleaseCatCache(m, v168)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L66
	}
L65:
	;
	goto L11
L66:
	;
	v182 = v177 ^ int32(1)
	goto L54
L67:
	;
	F_ReleaseCatCache(m, v36)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L75
	}
L68:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RangeVarCallbackForDropRelation[1])))
	if v187&int32(1) != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v191 = int32(1)
	if base.Ui32(l1) < base.Ui32(int32(_a_F_RangeVarCallbackForDropRelation_4)) {
		v199 = v191
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v199 != 0 {
		goto L10
	} else {
		goto L74
	}
L71:
	;
	goto L70
L72:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v42)+68))
	if v194 == int32(99) {
		v199 = v191
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v197 = F_isTempToastNamespace(m, v194)
	mBase = m.M
	v199 = v197
	goto L71
L74:
	;
	goto L67
L75:
	;
	if v18|base.B2i32(v56 != int32(105)) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if v18|base.B2i32(v43&int32(1) == int32(0)) != 0 {
		goto L11
	} else {
		goto L81
	}
L77:
	;
	v206 = F_IndexGetRelation(m, l1, int32(1))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v206
	if v206 == int32(0) {
		goto L76
	} else {
		goto L79
	}
L79:
	;
	F_LockRelationOid(m, v206, v17)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	v220 = F_get_partition_parent(m, l1, int32(1))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v220
	if v220 == int32(0) {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	F_LockRelationOid(m, v220, int32(8))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	goto L11
L85:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v245
	F_errmsg(m, int32(_a_F_RangeVarCallbackForDropRelation_5), v15)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForDropRelation_2), int32(1811), int32(_a_F_RangeVarCallbackForDropRelation_6))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RangeVarCallbackForLockTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == int32(0) {
		m.G0 = v10 + int32(16)
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		v15 = F_get_rel_relkind(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v18 = v15 & int32(255)
			v20 = v18 - int32(112)
			v27 = int32(0)
			if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v20))|base.B2i32(int32(1)<<(uint(v20)%32)&int32(69) == v27) == v27 {
				v32 = F_get_rel_persistence(m, l1)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(116) {
						v36 = int32(_a_F_RangeVarCallbackForLockTable_0)
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForLockTable[0]))
						*(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForLockTable[0])) = v38 | int32(1)
					} else {
					}
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForLockTable[1]))
					if v14 < int32(2) {
						v51 = int64(16414)
					} else {
						v51 = int64(16412)
					}
					v53 = F_pg_class_aclcheck(m, l1, v43, base.I64_extend_i32_u(base.B2i32(v14 < int32(4)))|v51)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return
					} else {
						if v53 == int32(0) {
							m.G0 = v10 + int32(16)
							return
						} else {
							v57 = F_get_rel_relkind(m, l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								switch v57 - int32(73) {
								case 0, 32:
									v68 = int32(20)
									v70 = v68
								default:
									v68 = int32(41)
									v70 = v68
								case 10:
									v70 = int32(37)
								case 29:
									v70 = int32(18)
								case 36:
									v70 = int32(23)
								case 45:
									v70 = int32(51)
								}
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_aclcheck_error(m, v53, v70, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					}
				}
			} else {
				if v18 == int32(0) {
					m.G0 = v10 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v83
							F_errmsg(m, int32(_a_F_RangeVarCallbackForLockTable_1), v10)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_errdetail_relkind_not_supported(m, v15)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RangeVarCallbackForLockTable_2), int32(93), int32(_a_F_RangeVarCallbackForLockTable_3))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
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
func F_RangeVarCallbackForReindexIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v14&int32(8) != 0 {
		v17 = int32(4)
	} else {
		v17 = int32(5)
	}
	v18 = base.B2i32(l1 == l2)
	v19 = int32(0)
	if v18|base.B2i32(l2 == v19) == v19 {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
		F_UnlockRelationOid(m, v24, v17)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
			if l1 == int32(0) {
				m.G0 = v10 + int32(16)
				return
			} else {
				v31 = F_get_rel_relkind(m, l1)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					if v31 == int32(0) {
						m.G0 = v10 + int32(16)
						return
					} else {
						if v31&int32(-33) != int32(73) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errcode(m, int32(151027844))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v68
									F_errmsg(m, int32(_a_F_RangeVarCallbackForReindexIndex_0), v10)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RangeVarCallbackForReindexIndex_1), int32(3016), int32(_a_F_RangeVarCallbackForReindexIndex_2))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
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
							v40 = F_IndexGetRelation(m, l1, int32(1))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								if v40 == int32(0) {
									m.G0 = v10 + int32(16)
									return
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForReindexIndex[0]))
									v47 = F_pg_class_aclcheck(m, v40, v45, int64(16384))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										if v47 != 0 {
											v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											F_aclcheck_error(m, v47, int32(20), v50)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return
											} else {
												if l1 == l2 {
													m.G0 = v10 + int32(16)
													return
												} else {
													F_LockRelationOid(m, v40, v17)
													mBase = m.M
													v54 = m.ExcPending
													if v54 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v40
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										} else {
											if l1 == l2 {
												m.G0 = v10 + int32(16)
												return
											} else {
												F_LockRelationOid(m, v40, v17)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v40
													m.G0 = v10 + int32(16)
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
			}
		}
	} else {
		if l1 == int32(0) {
			m.G0 = v10 + int32(16)
			return
		} else {
			v31 = F_get_rel_relkind(m, l1)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				if v31 == int32(0) {
					m.G0 = v10 + int32(16)
					return
				} else {
					if v31&int32(-33) != int32(73) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v68
								F_errmsg(m, int32(_a_F_RangeVarCallbackForReindexIndex_0), v10)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RangeVarCallbackForReindexIndex_1), int32(3016), int32(_a_F_RangeVarCallbackForReindexIndex_2))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
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
						v40 = F_IndexGetRelation(m, l1, int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							if v40 == int32(0) {
								m.G0 = v10 + int32(16)
								return
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForReindexIndex[0]))
								v47 = F_pg_class_aclcheck(m, v40, v45, int64(16384))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									if v47 != 0 {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										F_aclcheck_error(m, v47, int32(20), v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											if l1 == l2 {
												m.G0 = v10 + int32(16)
												return
											} else {
												F_LockRelationOid(m, v40, v17)
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v40
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									} else {
										if l1 == l2 {
											m.G0 = v10 + int32(16)
											return
										} else {
											F_LockRelationOid(m, v40, v17)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v40
												m.G0 = v10 + int32(16)
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
		}
	}
}
func F_RangeVarCallbackForStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = base.B2i32(l1 == l2)
	if l1 == l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_UnlockRelationOid(m, v14, int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L1
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L60
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L55
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L52
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L48
	}
L10:
	;
	m.G0 = v11 - int32(-64)
	return
L11:
	;
	v25 = F_get_rel_relkind(m, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v25&int32(-33) == int32(73) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v32 = F_IndexGetRelation(m, l1, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v34 = l1
	goto L15
L15:
	;
	if l1 != l2 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v34 = v32
	goto L15
L17:
	;
	v62 = F_SearchSysCache1(m, int32(57), v34)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L28
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l1 == v34 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v36 == int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v36 != v34 {
		goto L9
	} else {
		goto L27
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v47
	F_errmsg(m, int32(_a_F_RangeVarCallbackForStats_6), v9+int32(-32))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForStats_1), int32(178), int32(_a_F_RangeVarCallbackForStats_2))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	goto L17
L28:
	;
	if v62 == int32(0) {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+22)))
	v68 = v66 + v67
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+119)))
	v71 = v69 - int32(102)
	if base.B2i32(base.Ui32(int32(12)) < base.Ui32(v71))|base.B2i32(int32(1)<<(uint(v71)%32)&int32(_a_F_RangeVarCallbackForStats_3) == int32(0)) != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+117)))
	if v81 == int32(1) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForStats[0]))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForStats[1]))
	v89 = F_object_ownercheck(m, int32(1262), v86, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	F_ReleaseCatCache(m, v62)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L45
	}
L33:
	;
	if v89 != 0 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackForStats[1]))
	v94 = F_pg_class_aclcheck(m, v34, v92, int64(16384))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	if v94 == int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+119)))
	switch v98 - int32(73) {
	case 0, 32:
		v108 = int32(20)
		goto L38
	default:
		goto L39
	case 10:
		goto L43
	case 29:
		goto L40
	case 36:
		goto L41
	case 45:
		goto L42
	}
L37:
	;
	F_aclcheck_error(m, v94, v110, v68+int32(4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L44
	}
L38:
	;
	v110 = v108
	goto L37
L39:
	;
	v108 = int32(41)
	goto L38
L40:
	;
	v110 = int32(18)
	goto L37
L41:
	;
	v110 = int32(23)
	goto L37
L42:
	;
	v110 = int32(51)
	goto L37
L43:
	;
	v110 = int32(37)
	goto L37
L44:
	;
	goto L32
L45:
	;
	if v13|base.B2i32(l1 == v34) != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_LockRelationOid(m, v34, int32(4))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v34
	goto L10
L48:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v138
	F_errmsg(m, int32(_a_F_RangeVarCallbackForStats_7), v9+int32(-16))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForStats_1), int32(193), int32(_a_F_RangeVarCallbackForStats_2))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v34
	F_errmsg_internal(m, int32(_a_F_RangeVarCallbackForStats_0), v11)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForStats_1), int32(198), int32(_a_F_RangeVarCallbackForStats_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v68 + int32(4)
	F_errmsg(m, int32(_a_F_RangeVarCallbackForStats_4), v9+int32(-48))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	v178 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68)+119)))
	F_errdetail_relkind_not_supported(m, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForStats_1), int32(214), int32(_a_F_RangeVarCallbackForStats_2))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_RangeVarCallbackForStats_5), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_RangeVarCallbackForStats_1), int32(220), int32(_a_F_RangeVarCallbackForStats_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_addRangeTableEntryForValues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	v6 = l5
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v15 = F_palloc0(m, int32(136))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(101)
	v22 = F_pstrdup(m, int32(_a_F_addRangeTableEntryForValues_0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+104)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l1
	v28 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(5)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v28
	v35 = F_makeAlias(m, v22, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = v41
	goto L7
L6:
	;
	v42 = v37
	goto L7
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = v44
	goto L10
L9:
	;
	v45 = v37
	goto L10
L10:
	;
	if v45 < v42 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = v45
	goto L14
L12:
	;
	v78 = v45
	goto L13
L13:
	;
	if v78 <= v42 {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v57 = v49 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v57
	v60 = v12 + int32(32)
	v65 = F_pg_snprintf(m, v60, int32(64), int32(_a_F_addRangeTableEntryForValues_1), v12+int32(16))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v78 = v42
	goto L13
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v68 = F_pstrdup(m, v60)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v70 = F_makeString(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v72 = F_lappend(m, v67, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v72
	if v57 != v42 {
		v49 = v57
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+125)) = uint8(v86)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+124)) = uint8(v6)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v35
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v91 = F_lappend(m, v90, v15)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L29
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v91
	if v91 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v96 = v94
	goto L27
L26:
	;
	v96 = int32(0)
	goto L27
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v15)+96))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+100))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v15)+104))
	v100 = F_buildNSItemFromLists(m, v15, v96, v97, v98, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v12 + int32(96)
	return v100
L29:
	;
	F_errcode(m, int32(_a_F_addRangeTableEntryForValues_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v22
	F_errmsg(m, int32(_a_F_addRangeTableEntryForValues_3), v12)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_addRangeTableEntryForValues_4), int32(2195), int32(_a_F_addRangeTableEntryForValues_5))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_range_subtype(m *base.Module, l0 int32) int32 {
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
	v4 = F_SearchSysCache1(m, int32(55), l0)
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
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+4))
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
func F_makeRangeVarFromNameList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v11 = F_makeRangeVar(m, v2, v2, int32(-1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l0 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = F_NameListToString(m, l0)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v45
						F_errmsg(m, int32(_a_F_makeRangeVarFromNameList_0), v6)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_makeRangeVarFromNameList_1), int32(3579), int32(_a_F_makeRangeVarFromNameList_2))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
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
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v17 - int32(1) {
			case 0:
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v57 = v56
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v59
				m.G0 = v6 + int32(16)
				return v11
			case 1:
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v57 = v24 + int32(4)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v59
				m.G0 = v6 + int32(16)
				return v11
			case 2:
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v57 = v35 + int32(8)
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v59
				m.G0 = v6 + int32(16)
				return v11
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = F_NameListToString(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v45
							F_errmsg(m, int32(_a_F_makeRangeVarFromNameList_0), v6)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_makeRangeVarFromNameList_1), int32(3579), int32(_a_F_makeRangeVarFromNameList_2))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
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
func F_range_adjacent_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_range_adjacent_multirange_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_adjacent_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_adjacent_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_adjacent_multirange_2), int32(558), int32(_a_F_range_adjacent_multirange_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_adjacent_multirange_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_adjacent_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_adjacent_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_adjacent_multirange_2), int32(558), int32(_a_F_range_adjacent_multirange_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_adjacent_multirange_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_range_after_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v68 = v34
						m.G0 = v9 + int32(48)
						return v68
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v44 == int32(0) {
							v68 = v34
							m.G0 = v9 + int32(48)
							return v68
						} else {
							v48 = v9 + int32(40)
							F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								v61 = v9 + int32(16)
								F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v64 = F_range_cmp_bounds(m, v33, v48, v61)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										v68 = base.B2i32(int32(0) < v64)
										m.G0 = v9 + int32(48)
										return v68
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_after_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_after_multirange_1), v9)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_after_multirange_2), int32(558), int32(_a_F_range_after_multirange_3))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v68 = v34
								m.G0 = v9 + int32(48)
								return v68
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v44 == int32(0) {
									v68 = v34
									m.G0 = v9 + int32(48)
									return v68
								} else {
									v48 = v9 + int32(40)
									F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
										v61 = v9 + int32(16)
										F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = F_range_cmp_bounds(m, v33, v48, v61)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return int32(0)
											} else {
												v68 = base.B2i32(int32(0) < v64)
												m.G0 = v9 + int32(48)
												return v68
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_after_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_after_multirange_1), v9)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_after_multirange_2), int32(558), int32(_a_F_range_after_multirange_3))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v68 = v34
							m.G0 = v9 + int32(48)
							return v68
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v44 == int32(0) {
								v68 = v34
								m.G0 = v9 + int32(48)
								return v68
							} else {
								v48 = v9 + int32(40)
								F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
									v61 = v9 + int32(16)
									F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v64 = F_range_cmp_bounds(m, v33, v48, v61)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v68 = base.B2i32(int32(0) < v64)
											m.G0 = v9 + int32(48)
											return v68
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
func F_range_after_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v47 = v4
		m.G0 = v8 + int32(48)
		return v47
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v47 = v4
			m.G0 = v8 + int32(48)
			return v47
		} else {
			v23 = v8 + int32(40)
			F_range_deserialize(m, l0, l1, v23, v8+int32(32), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v38 = v8 + int32(16)
				F_multirange_get_bounds(m, l0, l2, v32-int32(1), v8+int32(24), v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = F_range_cmp_bounds(m, l0, v23, v38)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v47 = base.B2i32(int32(0) < v41)
						m.G0 = v8 + int32(48)
						return v47
					}
				}
			}
		}
	}
}
func F_range_before(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_before_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_before_0), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_before_1), int32(1776), int32(_a_F_range_before_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_before_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_before_0), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_before_1), int32(1776), int32(_a_F_range_before_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_before_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_before_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7+int32(40), v7+int32(24), v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7+int32(32), v7+int32(16), v7+int32(14))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v31 != 0 {
					v103 = v30
					m.G0 = v7 + int32(48)
					return v103
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32&int32(1) != 0 {
						v103 = v30
						m.G0 = v7 + int32(48)
						return v103
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
						if v36 == int32(1) {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
							if v35&int32(1) != 0 {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
								if v42 == v39 {
									v97 = int32(0)
								} else {
									v46 = int32(1)
									if v39&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v97 = v49
								}
							} else {
								v51 = int32(1)
								if v39&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v97 = v54
							}
							v103 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
							m.G0 = v7 + int32(48)
							return v103
						} else {
							if v35&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v97 = v60
								v103 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
								m.G0 = v7 + int32(48)
								return v103
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								v66 = F_FunctionCall2Coll(m, l0+int32(212), v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v97 = v66
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
										if v69 == int32(0) {
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v68&int32(1) == int32(0) {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
												if v77 == v72 {
													v97 = int32(0)
												} else {
													v80 = int32(1)
													if v72&v80 != 0 {
														v84 = v80
													} else {
														v84 = int32(-1)
													}
													v97 = v84
												}
											} else {
												v85 = int32(1)
												if v72&v85 != 0 {
													v89 = v85
												} else {
													v89 = int32(-1)
												}
												v97 = v89
											}
										} else {
											if v68&int32(1) != 0 {
												v97 = int32(0)
											} else {
												v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
												if v95 != 0 {
													v96 = int32(-1)
												} else {
													v96 = int32(1)
												}
												v97 = v96
											}
										}
									}
									v103 = int32(base.Ui32(v97) >> (uint(int32(31)) % 32))
									m.G0 = v7 + int32(48)
									return v103
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_range_before_internal_0), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_range_before_internal_1), int32(677), int32(_a_F_range_before_internal_2))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
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
func F_range_contained_by(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_contains_internal(m, v32, v17, v12)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_contained_by_0), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_contained_by_1), int32(1776), int32(_a_F_range_contained_by_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_contains_internal(m, v32, v17, v12)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_contained_by_0), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contained_by_1), int32(1776), int32(_a_F_range_contained_by_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_contains_internal(m, v32, v17, v12)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_contains_elem_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	F_range_deserialize(m, l0, l1, v8+int32(24), v8+int32(16), v8+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v21 != 0 {
			v54 = v20
			m.G0 = v8 + int32(32)
			return v54
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
			if v22 != 0 {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
				if v37 != 0 {
					v54 = int32(1)
					m.G0 = v8 + int32(32)
					return v54
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v42 = F_FunctionCall2Coll(m, l0+int32(212), v40, v41, l2)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 < int32(0) {
							v54 = v20
						} else {
							if v42 != 0 {
								v54 = int32(1)
							} else {
								v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
								if v46&int32(1) == int32(0) {
									v54 = v20
								} else {
									v54 = int32(1)
								}
							}
						}
						m.G0 = v8 + int32(32)
						return v54
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
				v27 = F_FunctionCall2Coll(m, l0+int32(212), v25, v26, l2)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if int32(0) < v27 {
						v54 = v20
						m.G0 = v8 + int32(32)
						return v54
					} else {
						if v27 != 0 {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
							if v37 != 0 {
								v54 = int32(1)
								m.G0 = v8 + int32(32)
								return v54
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
								v42 = F_FunctionCall2Coll(m, l0+int32(212), v40, v41, l2)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									if v42 < int32(0) {
										v54 = v20
									} else {
										if v42 != 0 {
											v54 = int32(1)
										} else {
											v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
											if v46&int32(1) == int32(0) {
												v54 = v20
											} else {
												v54 = int32(1)
											}
										}
									}
									m.G0 = v8 + int32(32)
									return v54
								}
							}
						} else {
							v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+29)))
							if v31&int32(1) == int32(0) {
								v54 = v20
								m.G0 = v8 + int32(32)
								return v54
							} else {
								v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
								if v37 != 0 {
									v54 = int32(1)
									m.G0 = v8 + int32(32)
									return v54
								} else {
									v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
									v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									v42 = F_FunctionCall2Coll(m, l0+int32(212), v40, v41, l2)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										if v42 < int32(0) {
											v54 = v20
										} else {
											if v42 != 0 {
												v54 = int32(1)
											} else {
												v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)))
												if v46&int32(1) == int32(0) {
													v54 = v20
												} else {
													v54 = int32(1)
												}
											}
										}
										m.G0 = v8 + int32(32)
										return v54
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
func F_range_contains_elem_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	if v3 == int32(457) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v12 = F_find_simplified_clause(m, v6, v10, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = v12
			return v18
		}
	} else {
		v18 = int32(0)
		return v18
	}
}
func F_range_gist_fallback_split(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v4 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v4
	if v11&int32(_a_F_range_gist_fallback_split_0) != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_range_gist_fallback_split_0)
	v25 = (v11 + v22) & v22
	v26 = int32(1)
	v29 = base.I32_div_s(v25-v26, int32(2))
	v33 = v26
	v36 = v26
	v37 = v4
	v38 = v4
	goto L4
L2:
	;
	v94 = v4
	v95 = v4
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v95
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v94
	return
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(4)+v33<<(uint(int32(4))%32))))
	v46 = F_pg_detoast_datum(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v94 = v82
	v95 = v83
	goto L3
L6:
	;
	return
L7:
	;
	if base.Ui32(v33) <= base.Ui32(v29) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v85 = v36 + int32(1)
	v87 = v85 & int32(_a_F_range_gist_fallback_split_0)
	if base.Ui32(v87) <= base.Ui32(v25) {
		v33 = v87
		v36 = v85
		v37 = v82
		v38 = v83
		goto L4
	} else {
		goto L22
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v49 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v65 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	v57 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v55 + v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v60+v55<<(uint(v57)%32)))) = uint16(v36)
	v82 = v56
	v83 = v38
	goto L8
L13:
	;
	v55 = v49
	v56 = v46
	goto L12
L14:
	;
	goto L15
L15:
	;
	v52 = F_range_super_union(m, l0, v37, v46)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v55 = v54
	v56 = v52
	goto L12
L17:
	;
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v71 + v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v76+v71<<(uint(v73)%32)))) = uint16(v36)
	v82 = v37
	v83 = v72
	goto L8
L18:
	;
	v71 = v65
	v72 = v46
	goto L17
L19:
	;
	goto L20
L20:
	;
	v68 = F_range_super_union(m, l0, v38, v46)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v71 = v70
	v72 = v68
	goto L17
L22:
	;
	goto L5
}
func F_range_lower(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
		if v17 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18 == v15 {
				v28 = v17
				F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v37 == int32(0) {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
						if v40&int32(1) == int32(0) {
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
							v49 = v48
						} else {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v49 = int32(0)
						}
					} else {
						v45 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
						v49 = int32(0)
					}
					m.G0 = v8 + int32(32)
					return v49
				}
			} else {
				v21 = F_lookup_type_cache(m, v15, int32(2048))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
					if v23 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
							F_errmsg_internal(m, int32(_a_F_range_lower_0), v8)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_lower_1), int32(1776), int32(_a_F_range_lower_2))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v21
						v28 = v21
						F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v37 == int32(0) {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
								if v40&int32(1) == int32(0) {
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
									v49 = v48
								} else {
									v45 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
									v49 = int32(0)
								}
							} else {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v49 = int32(0)
							}
							m.G0 = v8 + int32(32)
							return v49
						}
					}
				}
			}
		} else {
			v21 = F_lookup_type_cache(m, v15, int32(2048))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+200))
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
						F_errmsg_internal(m, int32(_a_F_range_lower_0), v8)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_range_lower_1), int32(1776), int32(_a_F_range_lower_2))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v21
					v28 = v21
					F_range_deserialize(m, v28, v11, v8+int32(24), v8+int32(16), v8+int32(15))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v37 == int32(0) {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+28)))
							if v40&int32(1) == int32(0) {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								v49 = v48
							} else {
								v45 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
								v49 = int32(0)
							}
						} else {
							v45 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v45)
							v49 = int32(0)
						}
						m.G0 = v8 + int32(32)
						return v49
					}
				}
			}
		}
	}
}
func F_range_overlaps_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = F_range_overlaps_multirange_internal(m, v33, v12, v17)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_overlaps_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_overlaps_multirange_1), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_overlaps_multirange_2), int32(558), int32(_a_F_range_overlaps_multirange_3))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = F_range_overlaps_multirange_internal(m, v33, v12, v17)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_overlaps_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_overlaps_multirange_1), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_overlaps_multirange_2), int32(558), int32(_a_F_range_overlaps_multirange_3))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = F_range_overlaps_multirange_internal(m, v33, v12, v17)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_range_overleft(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_overleft_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_overleft_0), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_overleft_1), int32(1776), int32(_a_F_range_overleft_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_overleft_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_overleft_0), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_overleft_1), int32(1776), int32(_a_F_range_overleft_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_overleft_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_overright_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v68 = v34
						m.G0 = v9 + int32(48)
						return v68
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v44 == int32(0) {
							v68 = v34
							m.G0 = v9 + int32(48)
							return v68
						} else {
							v48 = v9 + int32(40)
							F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v57 = v9 + int32(24)
								F_multirange_get_bounds(m, v33, v17, int32(0), v57, v9+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v62 = F_range_cmp_bounds(m, v33, v48, v57)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v68 = int32(base.Ui32(v62^int32(-1)) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v68
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_overright_multirange_0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_overright_multirange_1), v9)
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_range_overright_multirange_2), int32(558), int32(_a_F_range_overright_multirange_3))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v68 = v34
								m.G0 = v9 + int32(48)
								return v68
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v44 == int32(0) {
									v68 = v34
									m.G0 = v9 + int32(48)
									return v68
								} else {
									v48 = v9 + int32(40)
									F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v57 = v9 + int32(24)
										F_multirange_get_bounds(m, v33, v17, int32(0), v57, v9+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v62 = F_range_cmp_bounds(m, v33, v48, v57)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return int32(0)
											} else {
												v68 = int32(base.Ui32(v62^int32(-1)) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v68
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(_a_F_range_overright_multirange_0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_overright_multirange_1), v9)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_range_overright_multirange_2), int32(558), int32(_a_F_range_overright_multirange_3))
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v68 = v34
							m.G0 = v9 + int32(48)
							return v68
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v44 == int32(0) {
								v68 = v34
								m.G0 = v9 + int32(48)
								return v68
							} else {
								v48 = v9 + int32(40)
								F_range_deserialize(m, v33, v12, v48, v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v57 = v9 + int32(24)
									F_multirange_get_bounds(m, v33, v17, int32(0), v57, v9+int32(16))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v62 = F_range_cmp_bounds(m, v33, v48, v57)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v68 = int32(base.Ui32(v62^int32(-1)) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v68
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
func F_range_typanalyze(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v6 = F_getBaseType(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_range_get_typcache(m, l0, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			if v12 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_range_typanalyze[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v16
				v18 = v16
			} else {
				v18 = v12
			}
			*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = v10
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1474)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = v18 * int32(300)
			return int32(1)
		}
	}
}
