package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PageGetItemIdCareful_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v16 = l2 + l3<<(uint(int32(2))%32) + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v21 = int32(base.Ui32(v17) >> (uint(int32(17)) % 32))
	if base.Ui32(v17&int32(_a_F_PageGetItemIdCareful_2_0)+v21) < base.Ui32(int32(_a_F_PageGetItemIdCareful_2_1)) {
		switch int32(base.Ui32(v17)>>(uint(int32(15))%32)) & int32(3) {
		case 0, 2:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v84 + int32(4)
					F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_2), v10+int32(80))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = int32(base.Ui32(v95)>>(uint(int32(15))%32)) & int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(base.Ui32(v95) >> (uint(int32(17)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v95 & int32(_a_F_PageGetItemIdCareful_2_0)
						F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10+int32(48))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3524), int32(_a_F_PageGetItemIdCareful_2_5))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
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
		default:
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v84 + int32(4)
						F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_2), v10+int32(80))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = int32(base.Ui32(v95)>>(uint(int32(15))%32)) & int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(base.Ui32(v95) >> (uint(int32(17)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v95 & int32(_a_F_PageGetItemIdCareful_2_0)
							F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10+int32(48))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3524), int32(_a_F_PageGetItemIdCareful_2_5))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
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
				m.G0 = v10 + int32(96)
				return v16
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v45 + int32(4)
				F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_6), v10+int32(32))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(base.Ui32(v54)>>(uint(int32(15))%32)) & int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(base.Ui32(v54) >> (uint(int32(17)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v54 & int32(_a_F_PageGetItemIdCareful_2_0)
					F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3508), int32(_a_F_PageGetItemIdCareful_2_5))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
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
func F_ParseComplexProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v7 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	if v31 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v29 = F_get_expr_result_tupdesc(m, l2, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L13
	}
L4:
	;
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v15 = F_GetNSItemByRangeTablePosn(m, l0, v13, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v23 != int32(2249) {
		goto L3
	} else {
		goto L11
	}
L8:
	;
	return int32(0)
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v20 = F_scanNSItemForColumn(m, l0, v15, v19, l1, l3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	return v20
L11:
	;
	v26 = F_expandRecordVariable(m, l0, l2)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v31 = v26
	goto L2
L13:
	;
	v31 = v29
	goto L2
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 <= int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v43 = int32(0)
	goto L16
L16:
	;
	v51 = v31 + v34<<(uint(int32(4))%32) + int32(20) + v43*int32(100)
	v53 = v51 + int32(4)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if base.B2i32(v56 == int32(0))|base.B2i32(v56 != v59) != 0 {
		v77 = v56
		v78 = v59
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L1
L18:
	;
	v98 = v43 + int32(1)
	if v98 != v34 {
		v43 = v98
		goto L16
	} else {
		goto L29
	}
L19:
	;
	if v77-v78 != 0 {
		goto L18
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v62 = l1
	v63 = v53
	goto L22
L22:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v67 == int32(0) {
		v77 = v67
		v78 = v66
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v77 = v67
	v78 = v66
	goto L20
L24:
	;
	v70 = int32(1)
	if v67 == v66 {
		v62 = v62 + v70
		v63 = v63 + v70
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+91)))
	if v80 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v82 = F_palloc0(m, int32(24))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v85 = v43 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v82)+8)) = uint16(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = int32(25)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v51)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = v94
	return v82
L29:
	;
	goto L17
}
func F_ParseLongOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	v5 = int32(_a_F_ParseLongOption_0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*int8)(unsafe.Add(mBase, _c_F_ParseLongOption[0])))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v211 = v210
	goto L58
L2:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
	if v73 == int32(61) {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	m.G0 = v11 + int32(32)
	v71 = v64 - l0
	goto L2
L4:
	;
	F___memset(m, v11, int32(0), int32(32))
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseLongOption[0])))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseLongOption[1])))
	if v14 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v15 = F___strchrnul(m, l0, v13)
	mBase = m.M
	v64 = v15
	goto L3
L8:
	;
	goto L7
L9:
	;
	v21 = v5
	v22 = v19
	goto L12
L10:
	;
	goto L11
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v43 == int32(0) {
		v64 = l0
		goto L3
	} else {
		goto L15
	}
L12:
	;
	v29 = v11 + int32(base.Ui32(v22)>>(uint(int32(3))%32))&int32(28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30 | v31<<(uint(v22)%32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v35 != 0 {
		v21 = v21 + v31
		v22 = v35
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	v47 = l0
	v48 = v43
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v48)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v56)>>(uint(v48)%32))&int32(1) != 0 {
		v64 = v47
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v64 = v62
	goto L3
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v62 = v47 + int32(1)
	if v60 != 0 {
		v47 = v62
		v48 = v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v77 = v71 + int32(1)
	v78 = F_palloc(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v202 = F_pstrdup(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L23
	} else {
		goto L57
	}
L23:
	;
	return
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
	if v77 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v200 = F_pstrdup(m, l0+v77)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L56
	}
L26:
	;
	v196 = F_strlen(m, v192)
	mBase = m.M
	goto L25
L27:
	;
	v192 = l0
	goto L26
L28:
	;
	goto L29
L29:
	;
	v86 = v77 - int32(1)
	if (v78^l0)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v189)
	v192 = v185
	goto L26
L31:
	;
	v170 = v165
	v171 = v166
	v172 = v167
	goto L52
L32:
	;
	if v160 == int32(0) {
		v185 = v158
		v186 = v159
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v158 = l0
	v159 = v78
	v160 = v86
	goto L32
L34:
	;
	goto L35
L35:
	;
	v90 = int32(0)
	if base.B2i32(l0&int32(3) == v90)|base.B2i32(v86 == v90) == v90 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v126 == int32(0) {
		v185 = v123
		v186 = v124
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v102 = l0
	v103 = v78
	v104 = v86
	goto L40
L38:
	;
	goto L39
L39:
	;
	v123 = l0
	v124 = v78
	v125 = v86
	v126 = base.B2i32(v86 != v90)
	goto L36
L40:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v106)
	if v106 == int32(0) {
		v165 = v102
		v166 = v103
		v167 = v104
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v123 = v117
	v124 = v111
	v125 = v113
	v126 = v115
	goto L36
L42:
	;
	v110 = int32(1)
	v111 = v103 + v110
	v113 = v104 - v110
	v114 = int32(0)
	v115 = base.B2i32(v113 != v114)
	v117 = v102 + v110
	if v117&int32(3) == v114 {
		v123 = v117
		v124 = v111
		v125 = v113
		v126 = v115
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v113 != 0 {
		v102 = v117
		v103 = v111
		v104 = v113
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.B2i32(v129 == int32(0))|base.B2i32(base.Ui32(v125) < base.Ui32(int32(4))) != 0 {
		v158 = v123
		v159 = v124
		v160 = v125
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v136 = v123
	v137 = v124
	v138 = v125
	goto L47
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v144 = int32(-2139062144)
	if (int32(16843008)-v141|v141)&v144 != v144 {
		v165 = v136
		v166 = v137
		v167 = v138
		goto L31
	} else {
		goto L49
	}
L48:
	;
	v158 = v152
	v159 = v150
	v160 = v154
	goto L32
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v141
	v149 = int32(4)
	v150 = v137 + v149
	v152 = v136 + v149
	v154 = v138 - v149
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v136 = v152
		v137 = v150
		v138 = v154
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v165 = v158
	v166 = v159
	v167 = v160
	goto L31
L52:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	if v174 == int32(0) {
		v185 = v170
		v186 = v171
		goto L30
	} else {
		goto L54
	}
L53:
	;
	v185 = v181
	v186 = v179
	goto L30
L54:
	;
	v178 = int32(1)
	v179 = v171 + v178
	v181 = v170 + v178
	v183 = v172 - v178
	if v183 != 0 {
		v170 = v181
		v171 = v179
		v172 = v183
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v208 = v200
	goto L1
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v202
	v208 = int32(0)
	goto L1
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != int32(45) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v211 = v211 + int32(1)
	goto L58
L61:
	;
	if v215 != 0 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v218 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v218)
	goto L60
L64:
	;
	return
}
func F_PreCommit_CheckForSerializationFailure(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v149+int32(3584))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v119+int32(3584))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L29
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	v17 = F_LWLockAcquire(m, v13+int32(3584), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	return
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v21&int32(2056) == int32(8) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[2]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+32))
	v96 = v94 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v99 | int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v104+int32(3584))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v30 = v20 + int32(40)
	if v26 == v30 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = v26
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	if v42&int32(9) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v81 != v30 {
		v33 = v81
		goto L12
	} else {
		goto L27
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = v41 + int32(40)
	if v45 == v49 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v53 = v45
	goto L18
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v20 != v60 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L14
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v70 != v49 {
		v53 = v70
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+108)))
	if v62&int32(41) != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v42&int32(2) != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = v42 | int32(8)
	goto L14
L26:
	;
	goto L19
L27:
	;
	goto L13
L28:
	;
	goto L5
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_PreCommit_CheckForSerializationFailure_0), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errdetail_internal(m, int32(_a_F_PreCommit_CheckForSerializationFailure_1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errhint(m, int32(_a_F_PreCommit_CheckForSerializationFailure_2), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_PreCommit_CheckForSerializationFailure_3), int32(_a_F_PreCommit_CheckForSerializationFailure_4), int32(_a_F_PreCommit_CheckForSerializationFailure_5))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_PreCommit_CheckForSerializationFailure_0), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errdetail_internal(m, int32(_a_F_PreCommit_CheckForSerializationFailure_6), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errhint(m, int32(_a_F_PreCommit_CheckForSerializationFailure_2), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_PreCommit_CheckForSerializationFailure_3), int32(_a_F_PreCommit_CheckForSerializationFailure_7), int32(_a_F_PreCommit_CheckForSerializationFailure_5))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PreCommit_Portals(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_Portals[0]))
	F_hash_seq_init(m, v12, v14)
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
	v19 = F_hash_seq_search(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L38
	}
L5:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v22 = v19
	v26 = v2
	goto L9
L7:
	;
	v91 = v2
	goto L8
L8:
	;
	m.G0 = v9 + int32(32)
	return v91
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+84)))
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v91 = v81
	goto L8
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+85)))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v34 == int32(3) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v84 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L36
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	if v37 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v49 = int32(0)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+76)))
	if base.B2i32(v48 == v49)|(base.B2i32(v51&int32(32) == v49)|base.B2i32(v34 != int32(2))) == v49 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v44
	v81 = v26
	goto L15
L22:
	;
	F_UnregisterSnapshotFromOwner(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = int32(0)
	goto L21
L25:
	;
	goto L24
L26:
	;
	v70 = v9 + int32(12)
	F_hash_seq_term(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L34
	}
L27:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v48 == int32(0) {
		v81 = v26
		goto L15
	} else {
		goto L32
	}
L30:
	;
	F_HoldPortal(m, v27)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	F_PortalDrop(m, v27, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_Portals[0]))
	F_hash_seq_init(m, v70, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v81 = int32(1)
	goto L15
L36:
	;
	if v84 != 0 {
		v22 = v84
		v26 = v81
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L10
L38:
	;
	F_errmsg_internal(m, int32(_a_F_PreCommit_Portals_0), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_PreCommit_Portals_1), int32(695), int32(_a_F_PreCommit_Portals_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_PreCommit_Portals_3), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_PreCommit_Portals_1), int32(738), int32(_a_F_PreCommit_Portals_2))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PreCommit_on_commit_actions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v1 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_on_commit_actions[0]))
	if v12 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = v1
	v20 = v1
	v23 = v1
	goto L6
L4:
	;
	v53 = v1
	v56 = v1
	goto L5
L5:
	;
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v19<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v29 != 0 {
		v45 = v20
		v46 = v23
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = v45
	v56 = v46
	goto L5
L8:
	;
	v48 = v19 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v48 < v49 {
		v19 = v48
		v20 = v45
		v23 = v46
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	switch v30 - int32(2) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		v45 = v20
		v46 = v23
		goto L8
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v43 = F_lappend_oid(m, v23, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L15
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreCommit_on_commit_actions[1])))
	if v34&int32(1) == int32(0) {
		v45 = v20
		v46 = v23
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v40 = F_lappend_oid(m, v20, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v45 = v40
	v46 = v23
	goto L8
L15:
	;
	v45 = v20
	v46 = v43
	goto L8
L16:
	;
	goto L7
L17:
	;
	v57 = int32(0)
	if v53 == v57 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L20:
	;
	goto L19
L21:
	;
	F_heap_truncate_check_FKs(m, int32(0), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if int32(0) < v65 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v68 = v57
	v69 = v57
	goto L28
L26:
	;
	v89 = v57
	goto L27
L27:
	;
	F_heap_truncate_check_FKs(m, v89, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v68<<(uint(int32(2))%32))))
	v80 = F_table_open(m, v78, int32(8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	v89 = v82
	goto L27
L30:
	;
	v82 = F_lappend(m, v69, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v85 = v68 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v85 < v86 {
		v68 = v85
		v69 = v82
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v89 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	goto L20
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v99 <= int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v105 = int32(0)
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v105<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+119)))
	if v115 == int32(112) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	F_relation_close(m, v113, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L48
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+188))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
	m.T0[v119].(func(*base.Module, int32))(m, v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	F_RelationTruncateIndexes(m, v113)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+112))
	if v125 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v129 = F_table_open(m, v125, int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+188))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+116))
	m.T0[v132].(func(*base.Module, int32))(m, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_RelationTruncateIndexes(m, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_relation_close(m, v129, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	v145 = v105 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v145 < v146 {
		v105 = v145
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v168 = F_new_object_addresses(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if int32(0) < v170 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	v204 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L59
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1259)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180+v175<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v186
	F_add_exact_object_address(m, v9+int32(4), v168)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v195 = v175 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v195 < v196 {
		v175 = v195
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_PushActiveSnapshot(m, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_performMultipleDeletions(m, v168, int32(1), int32(5))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	goto L1
}
func F_PrefetchSharedBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v19
	v24 = v11 + int32(12)
	v25 = F_BufTableHashCode(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchSharedBuffer[0]))
		v35 = v28 + v25&int32(127)<<(uint(int32(7))%32) + int32(_a_F_PrefetchSharedBuffer_0)
		v37 = F_LWLockAcquire(m, v35, int32(1))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = F_BufTableLookup(m, v24, v25)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_LWLockRelease(m, v35)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					if v39 < int32(0) {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchSharedBuffer[1])))
						if v46&int32(1) != 0 {
							m.G0 = v11 + int32(32)
							return
						} else {
							v50 = F_smgrprefetch(m, l1, l2, l3, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if v50 == int32(0) {
								} else {
									v54 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v54)
								}
								m.G0 = v11 + int32(32)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v39 + int32(1)
						m.G0 = v11 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_PrepareInplaceInvalidationState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[0]))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v14
		} else {
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[1])) = v18
			*(*int64)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[2])) = v18
		}
		*(*int32)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[3])) = v5
		return v5
	}
}
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = F_MemoryContextAlloc(m, v4, int32(64))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_fmgr_info_cxt(m, l0, v6, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v6
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+60)) = uint8(v15)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+52)) = uint8(v15)
			v19 = int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+46)) = uint16(v19)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+44)) = uint8(v15)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(1819)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v6
			return
		}
	}
}
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PreventCommandIfParallelMode[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	if v10 != 0 {
		v13 = int32(1)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)))
		v13 = v12
	}
	if v13&int32(1) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandIfParallelMode_0), v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandIfParallelMode_1), int32(429), int32(_a_F_PreventCommandIfParallelMode_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandIfReadOnly[0])))
	if v8 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandIfReadOnly_0), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandIfReadOnly_1), int32(411), int32(_a_F_PreventCommandIfReadOnly_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v11) < base.Ui32(int32(2)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		if int32(2) <= v14 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errcode(m, int32(16777538))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
					F_errmsg(m, int32(_a_F_PreventInTransactionBlock_0), v7+int32(16))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3668), int32(_a_F_PreventInTransactionBlock_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
			if l0 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16777538))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
						F_errmsg(m, int32(_a_F_PreventInTransactionBlock_3), v7+int32(32))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3677), int32(_a_F_PreventInTransactionBlock_2))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
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
				v19 = int32(_a_F_PreventInTransactionBlock_4)
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[1])) = v21 | int32(4)
				m.G0 = v7 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errcode(m, int32(16777538))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(_a_F_PreventInTransactionBlock_5), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3658), int32(_a_F_PreventInTransactionBlock_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
func F_ProcessCheckpointerInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[0]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1]))
			if v6 == int32(0) {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
				if v35 != 0 {
					F_ProcessLogMemoryContextInterrupt(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_SyncRepUpdateSyncStandbysDefined(m)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_UpdateFullPageWrites(m)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v21 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								if v21 == int32(0) {
									v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
									if v35 != 0 {
										F_ProcessLogMemoryContextInterrupt(m)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											return
										}
									} else {
										return
									}
								} else {
									F_errmsg_internal(m, int32(_a_F_ProcessCheckpointerInterrupts_0), int32(0))
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ProcessCheckpointerInterrupts_1), int32(1402), int32(_a_F_ProcessCheckpointerInterrupts_2))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
											if v35 != 0 {
												F_ProcessLogMemoryContextInterrupt(m)
												mBase = m.M
												v37 = m.ExcPending
												if v37 != 0 {
													return
												} else {
													return
												}
											} else {
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
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1]))
		if v6 == int32(0) {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
			if v35 != 0 {
				F_ProcessLogMemoryContextInterrupt(m)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_SyncRepUpdateSyncStandbysDefined(m)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_UpdateFullPageWrites(m)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v21 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 == int32(0) {
								v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
								if v35 != 0 {
									F_ProcessLogMemoryContextInterrupt(m)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							} else {
								F_errmsg_internal(m, int32(_a_F_ProcessCheckpointerInterrupts_0), int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ProcessCheckpointerInterrupts_1), int32(1402), int32(_a_F_ProcessCheckpointerInterrupts_2))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
										if v35 != 0 {
											F_ProcessLogMemoryContextInterrupt(m)
											mBase = m.M
											v37 = m.ExcPending
											if v37 != 0 {
												return
											} else {
												return
											}
										} else {
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
func F_ProcessCopyOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
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
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v894 int32
	_ = v894
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1002 int32
	_ = v1002
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1142 int32
	_ = v1142
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1184 int32
	_ = v1184
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1207 int32
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1217 int32
	_ = v1217
	var v1220 int32
	_ = v1220
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1315 int32
	_ = v1315
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1338 int32
	_ = v1338
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1360 int32
	_ = v1360
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1383 int32
	_ = v1383
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1404 int32
	_ = v1404
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1448 int32
	_ = v1448
	var v1453 int32
	_ = v1453
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1473 int32
	_ = v1473
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1495 int64
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int64
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int64
	_ = v1499
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1516 int32
	_ = v1516
	var v1521 int32
	_ = v1521
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1545 int32
	_ = v1545
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1577 int32
	_ = v1577
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1652 int32
	_ = v1652
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1786 int32
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1808 int32
	_ = v1808
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1830 int32
	_ = v1830
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1862 int32
	_ = v1862
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1884 int32
	_ = v1884
	var v1889 int32
	_ = v1889
	var v1890 int32
	_ = v1890
	var v1896 int32
	_ = v1896
	var v1899 int32
	_ = v1899
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1926 int32
	_ = v1926
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1944 int32
	_ = v1944
	var v1951 int32
	_ = v1951
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1979 int32
	_ = v1979
	var v1986 int32
	_ = v1986
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2060 int32
	_ = v2060
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2086 int32
	_ = v2086
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2127 int32
	_ = v2127
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2139 int32
	_ = v2139
	var v2143 int32
	_ = v2143
	var v2146 int32
	_ = v2146
	var v2150 int32
	_ = v2150
	var v2155 int32
	_ = v2155
	var v2159 int32
	_ = v2159
	var v2162 int32
	_ = v2162
	var v2163 int32
	_ = v2163
	var v2169 int32
	_ = v2169
	var v2174 int32
	_ = v2174
	var v2178 int32
	_ = v2178
	var v2181 int32
	_ = v2181
	var v2188 int32
	_ = v2188
	var v2193 int32
	_ = v2193
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2204 int32
	_ = v2204
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2223 int32
	_ = v2223
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2239 int32
	_ = v2239
	var v2242 int32
	_ = v2242
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2275 int32
	_ = v2275
	var v2278 int32
	_ = v2278
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2307 int32
	_ = v2307
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2311 int32
	_ = v2311
	var v2313 int32
	_ = v2313
	var v2319 int32
	_ = v2319
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2329 int32
	_ = v2329
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2350 int64
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2362 int32
	_ = v2362
	var v2365 int32
	_ = v2365
	var v2372 int32
	_ = v2372
	var v2377 int32
	_ = v2377
	var v2381 int32
	_ = v2381
	var v2384 int32
	_ = v2384
	var v2391 int32
	_ = v2391
	var v2396 int32
	_ = v2396
	var v2400 int32
	_ = v2400
	var v2403 int32
	_ = v2403
	var v2412 int32
	_ = v2412
	var v2417 int32
	_ = v2417
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2433 int32
	_ = v2433
	var v2438 int32
	_ = v2438
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2471 int32
	_ = v2471
	var v2476 int32
	_ = v2476
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2487 int32
	_ = v2487
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2503 int32
	_ = v2503
	var v2508 int32
	_ = v2508
	var v2512 int32
	_ = v2512
	var v2515 int32
	_ = v2515
	var v2526 int32
	_ = v2526
	var v2531 int32
	_ = v2531
	var v2535 int32
	_ = v2535
	var v2538 int32
	_ = v2538
	var v2547 int32
	_ = v2547
	var v2552 int32
	_ = v2552
	var v2559 int32
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2569 int32
	_ = v2569
	var v2574 int32
	_ = v2574
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(464)
	m.G0 = v20
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_palloc0(m, int32(112))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = int32(-1)
	if l3 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v27 = v25
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L4
	} else {
		goto L867
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L4
	} else {
		goto L863
	}
L8:
	;
	v2258 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1632))))
	v2259 = F___strchrnul(m, v1675, v2258)
	mBase = m.M
	v2261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259))))
	if v2261 == v2258&int32(255) {
		goto L770
	} else {
		goto L771
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L4
	} else {
		goto L757
	}
L10:
	;
	if l2 != 0 {
		v2257 = v2229
		goto L8
	} else {
		goto L755
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2213 = m.ExcPending
	if v2213 != 0 {
		goto L4
	} else {
		goto L751
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2197 = m.ExcPending
	if v2197 != 0 {
		goto L4
	} else {
		goto L747
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2178 = m.ExcPending
	if v2178 != 0 {
		goto L4
	} else {
		goto L743
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L4
	} else {
		goto L739
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2143 = m.ExcPending
	if v2143 != 0 {
		goto L4
	} else {
		goto L735
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L4
	} else {
		goto L731
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L4
	} else {
		goto L727
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2095 = m.ExcPending
	if v2095 != 0 {
		goto L4
	} else {
		goto L723
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L4
	} else {
		goto L719
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L4
	} else {
		goto L715
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L4
	} else {
		goto L711
	}
L22:
	;
	F_errorConflictingDefElem(m, v56, l0)
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L4
	} else {
		goto L710
	}
L23:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v1597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)))
	if v1597 != 0 {
		goto L528
	} else {
		goto L529
	}
L24:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v32 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v43 = v5
	v44 = v5
	v45 = v5
	v46 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	goto L26
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v44<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = int32(_a_F_ProcessCopyOptions_0)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[0])))
	if base.B2i32(v61 == int32(0))|base.B2i32(v61 != v64) != 0 {
		v82 = v61
		v83 = v64
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L23
L28:
	;
	v1576 = v44 + int32(1)
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1576 < v1577 {
		v43 = v1567
		v44 = v1576
		v45 = v1568
		v46 = v1569
		v48 = v1571
		v49 = v1572
		v50 = v1573
		goto L26
	} else {
		goto L525
	}
L29:
	;
	if v82-v83 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v67 = v57
	v68 = v58
	goto L32
L32:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v72
		v83 = v71
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v82 = v72
	v83 = v71
	goto L30
L34:
	;
	v75 = int32(1)
	if v72 == v71 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v87 = F_defGetString(m, v56)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v206 = int32(_a_F_ProcessCopyOptions_1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[1])))
	if base.B2i32(v209 == int32(0))|base.B2i32(v209 != v212) != 0 {
		v230 = v209
		v231 = v212
		goto L77
	} else {
		goto L78
	}
L39:
	;
	if v43&int32(1) != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v91 = int32(_a_F_ProcessCopyOptions_2)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[2])))
	if base.B2i32(v94 == int32(0))|base.B2i32(v94 != v97) != 0 {
		v115 = v94
		v116 = v97
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v115-v116 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v100 = v87
	v101 = v91
	goto L44
L44:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v105
		v116 = v104
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v115 = v105
	v116 = v104
	goto L42
L46:
	;
	v108 = int32(1)
	if v105 == v104 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v1567 = int32(1)
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L49:
	;
	goto L50
L50:
	;
	v121 = int32(_a_F_ProcessCopyOptions_3)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[3])))
	if base.B2i32(v124 == int32(0))|base.B2i32(v124 != v127) != 0 {
		v145 = v124
		v146 = v127
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v145-v146 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	v130 = v87
	v131 = v121
	goto L54
L54:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v135 == int32(0) {
		v145 = v135
		v146 = v134
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v145 = v135
	v146 = v134
	goto L52
L56:
	;
	v138 = int32(1)
	if v135 == v134 {
		v130 = v130 + v138
		v131 = v131 + v138
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)) = uint8(v150)
	v1567 = v150
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L59:
	;
	goto L60
L60:
	;
	v153 = int32(_a_F_ProcessCopyOptions_4)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[4])))
	if base.B2i32(v156 == int32(0))|base.B2i32(v156 != v159) != 0 {
		v177 = v156
		v178 = v159
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v177-v178 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	v162 = v87
	v163 = v153
	goto L64
L64:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v167 == int32(0) {
		v177 = v167
		v178 = v166
		goto L62
	} else {
		goto L66
	}
L65:
	;
	v177 = v167
	v178 = v166
	goto L62
L66:
	;
	v170 = int32(1)
	if v167 == v166 {
		v162 = v162 + v170
		v163 = v163 + v170
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+4)) = uint8(v182)
	v1567 = v182
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L69:
	;
	goto L70
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+320)) = v87
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_5), v20+int32(320))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(576), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	if v230-v231 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	v215 = v57
	v216 = v206
	goto L79
L79:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v220 == int32(0) {
		v230 = v220
		v231 = v219
		goto L77
	} else {
		goto L81
	}
L80:
	;
	v230 = v220
	v231 = v219
	goto L77
L81:
	;
	v223 = int32(1)
	if v220 == v219 {
		v215 = v215 + v223
		v216 = v216 + v223
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if v46 != 0 {
		goto L22
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v239 = int32(_a_F_ProcessCopyOptions_8)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[5])))
	if base.B2i32(v242 == int32(0))|base.B2i32(v242 != v245) != 0 {
		v263 = v242
		v264 = v245
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v235 = F_defGetBoolean(m, v56)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)) = uint8(v235)
	v1567 = v43
	v1568 = v45
	v1569 = int32(1)
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L88:
	;
	if v263-v264 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	goto L88
L90:
	;
	v248 = v57
	v249 = v239
	goto L91
L91:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	if v253 == int32(0) {
		v263 = v253
		v264 = v252
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v263 = v253
	v264 = v252
	goto L89
L93:
	;
	v256 = int32(1)
	if v253 == v252 {
		v248 = v248 + v256
		v249 = v249 + v256
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v268 != 0 {
		goto L22
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v272 = int32(_a_F_ProcessCopyOptions_9)
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[6])))
	if base.B2i32(v275 == int32(0))|base.B2i32(v275 != v278) != 0 {
		v296 = v275
		v297 = v278
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v269 = F_defGetString(m, v56)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v269
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L100:
	;
	if v296-v297 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	goto L100
L102:
	;
	v281 = v57
	v282 = v272
	goto L103
L103:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281)+1)))
	if v286 == int32(0) {
		v296 = v286
		v297 = v285
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v296 = v286
	v297 = v285
	goto L101
L105:
	;
	v289 = int32(1)
	if v286 == v285 {
		v281 = v281 + v289
		v282 = v282 + v289
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v301 != 0 {
		goto L22
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v305 = int32(_a_F_ProcessCopyOptions_10)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[7])))
	if base.B2i32(v308 == int32(0))|base.B2i32(v308 != v311) != 0 {
		v329 = v308
		v330 = v311
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v302 = F_defGetString(m, v56)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v302
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L112:
	;
	if v329-v330 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	goto L112
L114:
	;
	v314 = v57
	v315 = v305
	goto L115
L115:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	if v319 == int32(0) {
		v329 = v319
		v330 = v318
		goto L113
	} else {
		goto L117
	}
L116:
	;
	v329 = v319
	v330 = v318
	goto L113
L117:
	;
	v322 = int32(1)
	if v319 == v318 {
		v314 = v314 + v322
		v315 = v315 + v322
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v334 != 0 {
		goto L22
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v338 = int32(_a_F_ProcessCopyOptions_11)
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[8])))
	if base.B2i32(v341 == int32(0))|base.B2i32(v341 != v344) != 0 {
		v362 = v341
		v363 = v344
		goto L125
	} else {
		goto L126
	}
L122:
	;
	v335 = F_defGetString(m, v56)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L4
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v335
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L124:
	;
	if v362-v363 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	goto L124
L126:
	;
	v347 = v57
	v348 = v338
	goto L127
L127:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+1)))
	if v352 == int32(0) {
		v362 = v352
		v363 = v351
		goto L125
	} else {
		goto L129
	}
L128:
	;
	v362 = v352
	v363 = v351
	goto L125
L129:
	;
	v355 = int32(1)
	if v352 == v351 {
		v347 = v347 + v355
		v348 = v348 + v355
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	if v45&int32(1) != 0 {
		goto L22
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v611 = int32(_a_F_ProcessCopyOptions_12)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v617 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[9])))
	if base.B2i32(v614 == int32(0))|base.B2i32(v614 != v617) != 0 {
		v635 = v614
		v636 = v617
		goto L214
	} else {
		goto L215
	}
L134:
	;
	v369 = int32(1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v371 == int32(0) {
		v607 = v369
		goto L135
	} else {
		goto L136
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v607
	v1567 = v43
	v1568 = v369
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L136:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
	if v374 == int32(465) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	if base.Ui32(v377) < base.Ui32(int32(2)) {
		v607 = v377
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v380 = F_defGetString(m, v56)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L4
	} else {
		goto L141
	}
L140:
	;
	goto L6
L141:
	;
	v385 = v380
	v386 = int32(_a_F_ProcessCopyOptions_13)
	goto L143
L142:
	;
	if v423 == int32(0) {
		v607 = v369
		goto L135
	} else {
		goto L155
	}
L143:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385))))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	if v389 == v390 {
		v412 = v389
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v423 = int32(0)
	goto L142
L145:
	;
	v414 = int32(1)
	if v412 != 0 {
		v385 = v385 + v414
		v386 = v386 + v414
		goto L143
	} else {
		goto L154
	}
L146:
	;
	if base.Ui32((v389-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v400 = v389 | int32(32)
	goto L149
L148:
	;
	v400 = v389
	goto L149
L149:
	;
	if base.Ui32((v390-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v409 = v390 | int32(32)
	goto L152
L151:
	;
	v409 = v390
	goto L152
L152:
	;
	if v400 == v409 {
		v412 = v400
		goto L145
	} else {
		goto L153
	}
L153:
	;
	v423 = v400 - v409
	goto L142
L154:
	;
	goto L144
L155:
	;
	v430 = v380
	v431 = int32(_a_F_ProcessCopyOptions_14)
	goto L157
L156:
	;
	if v468 == int32(0) {
		v607 = int32(0)
		goto L135
	} else {
		goto L169
	}
L157:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430))))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v431))))
	if v434 == v435 {
		v457 = v434
		goto L159
	} else {
		goto L160
	}
L158:
	;
	v468 = int32(0)
	goto L156
L159:
	;
	v459 = int32(1)
	if v457 != 0 {
		v430 = v430 + v459
		v431 = v431 + v459
		goto L157
	} else {
		goto L168
	}
L160:
	;
	if base.Ui32((v434-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v445 = v434 | int32(32)
	goto L163
L162:
	;
	v445 = v434
	goto L163
L163:
	;
	if base.Ui32((v435-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v454 = v435 | int32(32)
	goto L166
L165:
	;
	v454 = v435
	goto L166
L166:
	;
	if v445 == v454 {
		v457 = v445
		goto L159
	} else {
		goto L167
	}
L167:
	;
	v468 = v445 - v454
	goto L156
L168:
	;
	goto L158
L169:
	;
	v475 = v380
	v476 = int32(_a_F_ProcessCopyOptions_15)
	goto L171
L170:
	;
	if v513 == int32(0) {
		v607 = int32(1)
		goto L135
	} else {
		goto L183
	}
L171:
	;
	v479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v479 == v480 {
		v502 = v479
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v513 = int32(0)
	goto L170
L173:
	;
	v504 = int32(1)
	if v502 != 0 {
		v475 = v475 + v504
		v476 = v476 + v504
		goto L171
	} else {
		goto L182
	}
L174:
	;
	if base.Ui32((v479-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v490 = v479 | int32(32)
	goto L177
L176:
	;
	v490 = v479
	goto L177
L177:
	;
	if base.Ui32((v480-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v499 = v480 | int32(32)
	goto L180
L179:
	;
	v499 = v480
	goto L180
L180:
	;
	if v490 == v499 {
		v502 = v490
		goto L173
	} else {
		goto L181
	}
L181:
	;
	v513 = v490 - v499
	goto L170
L182:
	;
	goto L172
L183:
	;
	v520 = v380
	v521 = int32(_a_F_ProcessCopyOptions_16)
	goto L185
L184:
	;
	if v558 == int32(0) {
		v607 = int32(0)
		goto L135
	} else {
		goto L197
	}
L185:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v520))))
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521))))
	if v524 == v525 {
		v547 = v524
		goto L187
	} else {
		goto L188
	}
L186:
	;
	v558 = int32(0)
	goto L184
L187:
	;
	v549 = int32(1)
	if v547 != 0 {
		v520 = v520 + v549
		v521 = v521 + v549
		goto L185
	} else {
		goto L196
	}
L188:
	;
	if base.Ui32((v524-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v535 = v524 | int32(32)
	goto L191
L190:
	;
	v535 = v524
	goto L191
L191:
	;
	if base.Ui32((v525-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v544 = v525 | int32(32)
	goto L194
L193:
	;
	v544 = v525
	goto L194
L194:
	;
	if v535 == v544 {
		v547 = v535
		goto L187
	} else {
		goto L195
	}
L195:
	;
	v558 = v535 - v544
	goto L184
L196:
	;
	goto L186
L197:
	;
	v564 = v380
	v565 = int32(_a_F_ProcessCopyOptions_17)
	goto L199
L198:
	;
	if l2|v602 == int32(0) {
		goto L21
	} else {
		goto L211
	}
L199:
	;
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564))))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565))))
	if v568 == v569 {
		v591 = v568
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v602 = int32(0)
	goto L198
L201:
	;
	v593 = int32(1)
	if v591 != 0 {
		v564 = v564 + v593
		v565 = v565 + v593
		goto L199
	} else {
		goto L210
	}
L202:
	;
	if base.Ui32((v568-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v579 = v568 | int32(32)
	goto L205
L204:
	;
	v579 = v568
	goto L205
L205:
	;
	if base.Ui32((v569-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v588 = v569 | int32(32)
	goto L208
L207:
	;
	v588 = v569
	goto L208
L208:
	;
	if v579 == v588 {
		v591 = v579
		goto L201
	} else {
		goto L209
	}
L209:
	;
	v602 = v579 - v588
	goto L198
L210:
	;
	goto L200
L211:
	;
	if v602 != 0 {
		goto L6
	} else {
		goto L212
	}
L212:
	;
	v607 = int32(2)
	goto L135
L213:
	;
	if v635-v636 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L214:
	;
	goto L213
L215:
	;
	v620 = v57
	v621 = v611
	goto L216
L216:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+1)))
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620)+1)))
	if v625 == int32(0) {
		v635 = v625
		v636 = v624
		goto L214
	} else {
		goto L218
	}
L217:
	;
	v635 = v625
	v636 = v624
	goto L214
L218:
	;
	v628 = int32(1)
	if v625 == v624 {
		v620 = v620 + v628
		v621 = v621 + v628
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v640 != 0 {
		goto L22
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v644 = int32(_a_F_ProcessCopyOptions_18)
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v650 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[10])))
	if base.B2i32(v647 == int32(0))|base.B2i32(v647 != v650) != 0 {
		v668 = v647
		v669 = v650
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v641 = F_defGetString(m, v56)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v641
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L225:
	;
	if v668-v669 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L226:
	;
	goto L225
L227:
	;
	v653 = v57
	v654 = v644
	goto L228
L228:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654)+1)))
	v658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v653)+1)))
	if v658 == int32(0) {
		v668 = v658
		v669 = v657
		goto L226
	} else {
		goto L230
	}
L229:
	;
	v668 = v658
	v669 = v657
	goto L226
L230:
	;
	v661 = int32(1)
	if v658 == v657 {
		v653 = v653 + v661
		v654 = v654 + v661
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v673 != 0 {
		goto L22
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v677 = int32(_a_F_ProcessCopyOptions_19)
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v683 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[11])))
	if base.B2i32(v680 == int32(0))|base.B2i32(v680 != v683) != 0 {
		v701 = v680
		v702 = v683
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v674 = F_defGetString(m, v56)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v674
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v710
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L238:
	;
	if v701-v702 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L239:
	;
	goto L238
L240:
	;
	v686 = v57
	v687 = v677
	goto L241
L241:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v687)+1)))
	v691 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+1)))
	if v691 == int32(0) {
		v701 = v691
		v702 = v690
		goto L239
	} else {
		goto L243
	}
L242:
	;
	v701 = v691
	v702 = v690
	goto L239
L243:
	;
	v694 = int32(1)
	if v691 == v690 {
		v686 = v686 + v694
		v687 = v687 + v694
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v706 != 0 {
		goto L22
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v743 = int32(_a_F_ProcessCopyOptions_20)
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v749 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[12])))
	if base.B2i32(v746 == int32(0))|base.B2i32(v746 != v749) != 0 {
		v767 = v746
		v768 = v749
		goto L260
	} else {
		goto L261
	}
L248:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v707 == int32(1) {
		goto L22
	} else {
		goto L249
	}
L249:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v710 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L254
	}
L251:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	if v713 == int32(1) {
		goto L237
	} else {
		goto L252
	}
L252:
	;
	if v713 != int32(77) {
		goto L250
	} else {
		goto L253
	}
L253:
	;
	v718 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)) = uint8(v718)
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L254:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+368)) = v728
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_21), v20+int32(368))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L256
	}
L256:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v735)
	mBase = m.M
	v737 = m.ExcPending
	if v737 != 0 {
		goto L4
	} else {
		goto L257
	}
L257:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(635), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L4
	} else {
		goto L258
	}
L258:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L259:
	;
	if v767-v768 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L260:
	;
	goto L259
L261:
	;
	v752 = v57
	v753 = v743
	goto L262
L262:
	;
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753)+1)))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+1)))
	if v757 == int32(0) {
		v767 = v757
		v768 = v756
		goto L260
	} else {
		goto L264
	}
L263:
	;
	v767 = v757
	v768 = v756
	goto L260
L264:
	;
	v760 = int32(1)
	if v757 == v756 {
		v752 = v752 + v760
		v753 = v753 + v760
		goto L262
	} else {
		goto L265
	}
L265:
	;
	goto L263
L266:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v772 != 0 {
		goto L22
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v810 = int32(_a_F_ProcessCopyOptions_22)
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v816 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[13])))
	if base.B2i32(v813 == int32(0))|base.B2i32(v813 != v816) != 0 {
		v834 = v813
		v835 = v816
		goto L283
	} else {
		goto L284
	}
L269:
	;
	v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v773 == int32(1) {
		goto L22
	} else {
		goto L270
	}
L270:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v776 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L4
	} else {
		goto L277
	}
L272:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v776)))
	if v779 != int32(1) {
		goto L273
	} else {
		goto L274
	}
L273:
	;
	if v779 != int32(77) {
		goto L271
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = v776
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L276:
	;
	v784 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)) = uint8(v784)
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L277:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L4
	} else {
		goto L278
	}
L278:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+384)) = v795
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_21), v20+int32(384))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L4
	} else {
		goto L279
	}
L279:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v802)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(650), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L4
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	if v834-v835 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L283:
	;
	goto L282
L284:
	;
	v819 = v57
	v820 = v810
	goto L285
L285:
	;
	v823 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v820)+1)))
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v819)+1)))
	if v824 == int32(0) {
		v834 = v824
		v835 = v823
		goto L283
	} else {
		goto L287
	}
L286:
	;
	v834 = v824
	v835 = v823
	goto L283
L287:
	;
	v827 = int32(1)
	if v824 == v823 {
		v819 = v819 + v827
		v820 = v820 + v827
		goto L285
	} else {
		goto L288
	}
L288:
	;
	goto L286
L289:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v839 != 0 {
		goto L22
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v877 = int32(_a_F_ProcessCopyOptions_23)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[14])))
	if base.B2i32(v880 == int32(0))|base.B2i32(v880 != v883) != 0 {
		v901 = v880
		v902 = v883
		goto L306
	} else {
		goto L307
	}
L292:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v840 == int32(1) {
		goto L22
	} else {
		goto L293
	}
L293:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v843 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L4
	} else {
		goto L300
	}
L295:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v843)))
	if v846 != int32(1) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	if v846 != int32(77) {
		goto L294
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v843
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L299:
	;
	v851 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)) = uint8(v851)
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L300:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+400)) = v862
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_21), v20+int32(400))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(665), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L4
	} else {
		goto L304
	}
L304:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L305:
	;
	if v901-v902 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L306:
	;
	goto L305
L307:
	;
	v886 = v57
	v887 = v877
	goto L308
L308:
	;
	v890 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887)+1)))
	v891 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v886)+1)))
	if v891 == int32(0) {
		v901 = v891
		v902 = v890
		goto L306
	} else {
		goto L310
	}
L309:
	;
	v901 = v891
	v902 = v890
	goto L306
L310:
	;
	v894 = int32(1)
	if v891 == v890 {
		v886 = v886 + v894
		v887 = v887 + v894
		goto L308
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	v906 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)))
	if v906 == int32(1) {
		goto L22
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v938 = int32(_a_F_ProcessCopyOptions_24)
	v941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v944 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[15])))
	if base.B2i32(v941 == int32(0))|base.B2i32(v941 != v944) != 0 {
		v962 = v941
		v963 = v944
		goto L327
	} else {
		goto L328
	}
L315:
	;
	v909 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)) = uint8(v909)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v911 != 0 {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L4
	} else {
		goto L321
	}
L317:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)))
	if v912 != int32(1) {
		goto L316
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v911
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = v50
	goto L28
L320:
	;
	goto L319
L321:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+416)) = v923
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_21), v20+int32(416))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L4
	} else {
		goto L323
	}
L323:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(684), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L4
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	if v962-v963 == int32(0) {
		goto L333
	} else {
		goto L334
	}
L327:
	;
	goto L326
L328:
	;
	v947 = v57
	v948 = v938
	goto L329
L329:
	;
	v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+1)))
	v952 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v947)+1)))
	if v952 == int32(0) {
		v962 = v952
		v963 = v951
		goto L327
	} else {
		goto L331
	}
L330:
	;
	v962 = v952
	v963 = v951
	goto L327
L331:
	;
	v955 = int32(1)
	if v952 == v951 {
		v947 = v947 + v955
		v948 = v948 + v955
		goto L329
	} else {
		goto L332
	}
L332:
	;
	goto L330
L333:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= v967 {
		goto L22
	} else {
		goto L336
	}
L334:
	;
	goto L335
L335:
	;
	v1089 = int32(_a_F_ProcessCopyOptions_25)
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[16])))
	if base.B2i32(v1092 == int32(0))|base.B2i32(v1092 != v1095) != 0 {
		v1113 = v1092
		v1114 = v1095
		goto L371
	} else {
		goto L372
	}
L336:
	;
	v970 = F_defGetString(m, v56)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L4
	} else {
		goto L337
	}
L337:
	;
	v979 = m.G0
	v981 = v979 + int32(-64)
	m.G0 = v981
	v983 = int32(-1)
	if v970 == int32(0) {
		v1058 = v983
		goto L339
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1058
	if int32(0) <= v1058 {
		v1567 = v43
		v1568 = v45
		v1569 = v46
		v1571 = v48
		v1572 = v49
		v1573 = v50
		goto L28
	} else {
		goto L364
	}
L339:
	;
	m.G0 = v981 - int32(-64)
	goto L338
L340:
	;
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v970))))
	if v986 == int32(0) {
		v1058 = v983
		goto L339
	} else {
		goto L341
	}
L341:
	;
	v989 = F_strlen(m, v970)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v989) {
		v1058 = v983
		goto L339
	} else {
		goto L342
	}
L342:
	;
	v992 = v970
	v993 = v986
	v994 = v981
	goto L343
L343:
	;
	v1002 = F_isalnum(m, v993&int32(255))
	mBase = m.M
	if v1002 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v1019 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1015))) = uint8(v1019)
	v1023 = int32(*(*int8)(unsafe.Add(mBase, uint32(v981))))
	v1024 = int32(_a_F_ProcessCopyOptions_26)
	v1025 = int32(_a_F_ProcessCopyOptions_27)
	goto L352
L345:
	;
	if base.Ui32((v993-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	v1015 = v994
	goto L347
L347:
	;
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+1)))
	if v1016 != 0 {
		v992 = v992 + int32(1)
		v993 = v1016
		v994 = v1015
		goto L343
	} else {
		goto L351
	}
L348:
	;
	v1011 = v993 | int32(32)
	goto L350
L349:
	;
	v1011 = v993
	goto L350
L350:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v994))) = uint8(v1011)
	v1015 = v994 + int32(1)
	goto L347
L351:
	;
	goto L344
L352:
	;
	v1037 = v1025 + (v1024-v1025)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)))
	v1039 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1038))))
	v1040 = v1023 - v1039
	if v1040 != 0 {
		v1043 = v1040
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1058 = v983
	goto L339
L354:
	;
	v1047 = base.B2i32(v1043 < int32(0))
	if v1043 < int32(0) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	v1041 = F_strcmp(m, v981, v1038)
	mBase = m.M
	if v1041 != 0 {
		v1043 = v1041
		goto L354
	} else {
		goto L356
	}
L356:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	v1058 = v1042
	goto L339
L357:
	;
	v1048 = v1037 - int32(8)
	goto L359
L358:
	;
	v1048 = v1024
	goto L359
L359:
	;
	if v1043 < int32(0) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1051 = v1025
	goto L362
L361:
	;
	v1051 = v1037 + int32(8)
	goto L362
L362:
	;
	if base.Ui32(v1051) <= base.Ui32(v1048) {
		v1024 = v1048
		v1025 = v1051
		goto L352
	} else {
		goto L363
	}
L363:
	;
	goto L353
L364:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L4
	} else {
		goto L365
	}
L365:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L4
	} else {
		goto L366
	}
L366:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+432)) = v1074
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_28), v20+int32(432))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L4
	} else {
		goto L367
	}
L367:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1081)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L4
	} else {
		goto L368
	}
L368:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(696), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L4
	} else {
		goto L369
	}
L369:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L370:
	;
	if v1113-v1114 == int32(0) {
		goto L377
	} else {
		goto L378
	}
L371:
	;
	goto L370
L372:
	;
	v1098 = v57
	v1099 = v1089
	goto L373
L373:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099)+1)))
	v1103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1098)+1)))
	if v1103 == int32(0) {
		v1113 = v1103
		v1114 = v1102
		goto L371
	} else {
		goto L375
	}
L374:
	;
	v1113 = v1103
	v1114 = v1102
	goto L371
L375:
	;
	v1106 = int32(1)
	if v1103 == v1102 {
		v1098 = v1098 + v1106
		v1099 = v1099 + v1106
		goto L373
	} else {
		goto L376
	}
L376:
	;
	goto L374
L377:
	;
	if v48 != 0 {
		goto L22
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	v1261 = int32(_a_F_ProcessCopyOptions_29)
	v1264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1267 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[17])))
	if base.B2i32(v1264 == int32(0))|base.B2i32(v1264 != v1267) != 0 {
		v1285 = v1264
		v1286 = v1267
		goto L428
	} else {
		goto L429
	}
L380:
	;
	v1118 = m.G0
	v1120 = v1118 - int32(32)
	m.G0 = v1120
	v1122 = F_defGetString(m, v56)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L4
	} else {
		goto L382
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+84)) = v1210
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = int32(1)
	v1572 = v49
	v1573 = v50
	goto L28
L382:
	;
	if l2 != 0 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L4
	} else {
		goto L422
	}
L384:
	;
	v1127 = v1122
	v1128 = int32(_a_F_ProcessCopyOptions_30)
	goto L388
L385:
	;
	goto L386
L386:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L4
	} else {
		goto L417
	}
L387:
	;
	if v1165 != 0 {
		goto L400
	} else {
		goto L401
	}
L388:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1127))))
	v1132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1128))))
	if v1131 == v1132 {
		v1154 = v1131
		goto L390
	} else {
		goto L391
	}
L389:
	;
	v1165 = int32(0)
	goto L387
L390:
	;
	v1156 = int32(1)
	if v1154 != 0 {
		v1127 = v1127 + v1156
		v1128 = v1128 + v1156
		goto L388
	} else {
		goto L399
	}
L391:
	;
	if base.Ui32((v1131-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1142 = v1131 | int32(32)
	goto L394
L393:
	;
	v1142 = v1131
	goto L394
L394:
	;
	if base.Ui32((v1132-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1151 = v1132 | int32(32)
	goto L397
L396:
	;
	v1151 = v1132
	goto L397
L397:
	;
	if v1142 == v1151 {
		v1154 = v1142
		goto L390
	} else {
		goto L398
	}
L398:
	;
	v1165 = v1142 - v1151
	goto L387
L399:
	;
	goto L389
L400:
	;
	v1169 = v1122
	v1170 = int32(_a_F_ProcessCopyOptions_31)
	goto L404
L401:
	;
	v1210 = int32(0)
	goto L402
L402:
	;
	m.G0 = v1120 + int32(32)
	goto L381
L403:
	;
	if v1207 != 0 {
		goto L383
	} else {
		goto L416
	}
L404:
	;
	v1173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1169))))
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170))))
	if v1173 == v1174 {
		v1196 = v1173
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1207 = int32(0)
	goto L403
L406:
	;
	v1198 = int32(1)
	if v1196 != 0 {
		v1169 = v1169 + v1198
		v1170 = v1170 + v1198
		goto L404
	} else {
		goto L415
	}
L407:
	;
	if base.Ui32((v1173-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1184 = v1173 | int32(32)
	goto L410
L409:
	;
	v1184 = v1173
	goto L410
L410:
	;
	if base.Ui32((v1174-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v1193 = v1174 | int32(32)
	goto L413
L412:
	;
	v1193 = v1174
	goto L413
L413:
	;
	if v1184 == v1193 {
		v1196 = v1184
		goto L406
	} else {
		goto L414
	}
L414:
	;
	v1207 = v1184 - v1193
	goto L403
L415:
	;
	goto L405
L416:
	;
	v1210 = int32(1)
	goto L402
L417:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L4
	} else {
		goto L418
	}
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1120)+20)) = int32(_a_F_ProcessCopyOptions_32)
	*(*int32)(unsafe.Add(mBase, uint32(v1120)+16)) = int32(_a_F_ProcessCopyOptions_33)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v1120+int32(16))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L4
	} else {
		goto L419
	}
L419:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1230)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L4
	} else {
		goto L420
	}
L420:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(442), int32(_a_F_ProcessCopyOptions_35))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L4
	} else {
		goto L421
	}
L421:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L422:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L4
	} else {
		goto L423
	}
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1120)+4)) = v1122
	*(*int32)(unsafe.Add(mBase, uint32(v1120))) = int32(_a_F_ProcessCopyOptions_33)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_36), v1120)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L4
	} else {
		goto L424
	}
L424:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1251)
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L4
	} else {
		goto L425
	}
L425:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(456), int32(_a_F_ProcessCopyOptions_35))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L4
	} else {
		goto L426
	}
L426:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L427:
	;
	if v1285-v1286 == int32(0) {
		goto L434
	} else {
		goto L435
	}
L428:
	;
	goto L427
L429:
	;
	v1270 = v57
	v1271 = v1261
	goto L430
L430:
	;
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271)+1)))
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1270)+1)))
	if v1275 == int32(0) {
		v1285 = v1275
		v1286 = v1274
		goto L428
	} else {
		goto L432
	}
L431:
	;
	v1285 = v1275
	v1286 = v1274
	goto L428
L432:
	;
	v1278 = int32(1)
	if v1275 == v1274 {
		v1270 = v1270 + v1278
		v1271 = v1271 + v1278
		goto L430
	} else {
		goto L433
	}
L433:
	;
	goto L431
L434:
	;
	if v49 != 0 {
		goto L22
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	v1456 = int32(_a_F_ProcessCopyOptions_37)
	v1459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1462 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[18])))
	if base.B2i32(v1459 == int32(0))|base.B2i32(v1459 != v1462) != 0 {
		v1480 = v1459
		v1481 = v1462
		goto L490
	} else {
		goto L491
	}
L437:
	;
	v1290 = m.G0
	v1292 = v1290 - int32(16)
	m.G0 = v1292
	v1295 = F_defGetString(m, v56)
	mBase = m.M
	v1296 = m.ExcPending
	if v1296 != 0 {
		goto L4
	} else {
		goto L441
	}
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v1429
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = int32(1)
	v1573 = v50
	goto L28
L439:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L4
	} else {
		goto L484
	}
L440:
	;
	m.G0 = v1292 + int32(16)
	goto L438
L441:
	;
	v1300 = v1295
	v1301 = int32(_a_F_ProcessCopyOptions_38)
	goto L443
L442:
	;
	if v1338 == int32(0) {
		v1429 = int32(-1)
		goto L440
	} else {
		goto L455
	}
L443:
	;
	v1304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1300))))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301))))
	if v1304 == v1305 {
		v1327 = v1304
		goto L445
	} else {
		goto L446
	}
L444:
	;
	v1338 = int32(0)
	goto L442
L445:
	;
	v1329 = int32(1)
	if v1327 != 0 {
		v1300 = v1300 + v1329
		v1301 = v1301 + v1329
		goto L443
	} else {
		goto L454
	}
L446:
	;
	if base.Ui32((v1304-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v1315 = v1304 | int32(32)
	goto L449
L448:
	;
	v1315 = v1304
	goto L449
L449:
	;
	if base.Ui32((v1305-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v1324 = v1305 | int32(32)
	goto L452
L451:
	;
	v1324 = v1305
	goto L452
L452:
	;
	if v1315 == v1324 {
		v1327 = v1315
		goto L445
	} else {
		goto L453
	}
L453:
	;
	v1338 = v1315 - v1324
	goto L442
L454:
	;
	goto L444
L455:
	;
	v1345 = v1295
	v1346 = int32(_a_F_ProcessCopyOptions_10)
	goto L457
L456:
	;
	if v1383 == int32(0) {
		v1429 = int32(0)
		goto L440
	} else {
		goto L469
	}
L457:
	;
	v1349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1345))))
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346))))
	if v1349 == v1350 {
		v1372 = v1349
		goto L459
	} else {
		goto L460
	}
L458:
	;
	v1383 = int32(0)
	goto L456
L459:
	;
	v1374 = int32(1)
	if v1372 != 0 {
		v1345 = v1345 + v1374
		v1346 = v1346 + v1374
		goto L457
	} else {
		goto L468
	}
L460:
	;
	if base.Ui32((v1349-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v1360 = v1349 | int32(32)
	goto L463
L462:
	;
	v1360 = v1349
	goto L463
L463:
	;
	if base.Ui32((v1350-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L464
	} else {
		goto L465
	}
L464:
	;
	v1369 = v1350 | int32(32)
	goto L466
L465:
	;
	v1369 = v1350
	goto L466
L466:
	;
	if v1360 == v1369 {
		v1372 = v1360
		goto L459
	} else {
		goto L467
	}
L467:
	;
	v1383 = v1360 - v1369
	goto L456
L468:
	;
	goto L458
L469:
	;
	v1389 = v1295
	v1390 = int32(_a_F_ProcessCopyOptions_39)
	goto L471
L470:
	;
	if v1427 != 0 {
		goto L439
	} else {
		goto L483
	}
L471:
	;
	v1393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1389))))
	v1394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1390))))
	if v1393 == v1394 {
		v1416 = v1393
		goto L473
	} else {
		goto L474
	}
L472:
	;
	v1427 = int32(0)
	goto L470
L473:
	;
	v1418 = int32(1)
	if v1416 != 0 {
		v1389 = v1389 + v1418
		v1390 = v1390 + v1418
		goto L471
	} else {
		goto L482
	}
L474:
	;
	if base.Ui32((v1393-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v1404 = v1393 | int32(32)
	goto L477
L476:
	;
	v1404 = v1393
	goto L477
L477:
	;
	if base.Ui32((v1394-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1413 = v1394 | int32(32)
	goto L480
L479:
	;
	v1413 = v1394
	goto L480
L480:
	;
	if v1404 == v1413 {
		v1416 = v1404
		goto L473
	} else {
		goto L481
	}
L481:
	;
	v1427 = v1404 - v1413
	goto L470
L482:
	;
	goto L472
L483:
	;
	v1429 = int32(1)
	goto L440
L484:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L4
	} else {
		goto L485
	}
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+4)) = v1295
	*(*int32)(unsafe.Add(mBase, uint32(v1292))) = int32(_a_F_ProcessCopyOptions_40)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_36), v1292)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L4
	} else {
		goto L486
	}
L486:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1446)
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L4
	} else {
		goto L487
	}
L487:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(514), int32(_a_F_ProcessCopyOptions_41))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L4
	} else {
		goto L488
	}
L488:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L489:
	;
	if v1480-v1481 == int32(0) {
		goto L496
	} else {
		goto L497
	}
L490:
	;
	goto L489
L491:
	;
	v1465 = v57
	v1466 = v1456
	goto L492
L492:
	;
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1466)+1)))
	v1470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465)+1)))
	if v1470 == int32(0) {
		v1480 = v1470
		v1481 = v1469
		goto L490
	} else {
		goto L494
	}
L493:
	;
	v1480 = v1470
	v1481 = v1469
	goto L490
L494:
	;
	v1473 = int32(1)
	if v1470 == v1469 {
		v1465 = v1465 + v1473
		v1466 = v1466 + v1473
		goto L492
	} else {
		goto L495
	}
L495:
	;
	goto L493
L496:
	;
	if v50 != 0 {
		goto L22
	} else {
		goto L499
	}
L497:
	;
	goto L498
L498:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L4
	} else {
		goto L520
	}
L499:
	;
	v1485 = m.G0
	v1487 = v1485 - int32(32)
	m.G0 = v1487
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v1489 != 0 {
		goto L502
	} else {
		goto L503
	}
L500:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+96)) = v1499
	v1567 = v43
	v1568 = v45
	v1569 = v46
	v1571 = v48
	v1572 = v49
	v1573 = int32(1)
	goto L28
L501:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L4
	} else {
		goto L516
	}
L502:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1489)))
	if v1490 == int32(468) {
		goto L506
	} else {
		goto L507
	}
L503:
	;
	goto L504
L504:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1508 = m.ExcPending
	if v1508 != 0 {
		goto L4
	} else {
		goto L512
	}
L505:
	;
	if v1499 <= int64(0) {
		goto L501
	} else {
		goto L511
	}
L506:
	;
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+4))
	v1495 = F_pg_strtoint64_safe(m, v1493, int32(0))
	mBase = m.M
	v1496 = m.ExcPending
	if v1496 != 0 {
		goto L4
	} else {
		goto L509
	}
L507:
	;
	goto L508
L508:
	;
	v1497 = F_defGetInt64(m, v56)
	mBase = m.M
	v1498 = m.ExcPending
	if v1498 != 0 {
		goto L4
	} else {
		goto L510
	}
L509:
	;
	v1499 = v1495
	goto L505
L510:
	;
	v1499 = v1497
	goto L505
L511:
	;
	m.G0 = v1487 + int32(32)
	goto L500
L512:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L4
	} else {
		goto L513
	}
L513:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1487))) = v1512
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_42), v1487)
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L4
	} else {
		goto L514
	}
L514:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(476), int32(_a_F_ProcessCopyOptions_43))
	mBase = m.M
	v1521 = m.ExcPending
	if v1521 != 0 {
		goto L4
	} else {
		goto L515
	}
L515:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L516:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1528 = m.ExcPending
	if v1528 != 0 {
		goto L4
	} else {
		goto L517
	}
L517:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1487)+16)) = v1499
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_44), v1487+int32(16))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L4
	} else {
		goto L518
	}
L518:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(486), int32(_a_F_ProcessCopyOptions_43))
	mBase = m.M
	v1539 = m.ExcPending
	if v1539 != 0 {
		goto L4
	} else {
		goto L519
	}
L519:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L520:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1548 = m.ExcPending
	if v1548 != 0 {
		goto L4
	} else {
		goto L521
	}
L521:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+448)) = v1549
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_45), v20+int32(448))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L4
	} else {
		goto L522
	}
L522:
	;
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1556)
	mBase = m.M
	v1558 = m.ExcPending
	if v1558 != 0 {
		goto L4
	} else {
		goto L523
	}
L523:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(724), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L4
	} else {
		goto L524
	}
L524:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L525:
	;
	goto L27
L526:
	;
	v1633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v1634 == int32(0) {
		goto L546
	} else {
		goto L547
	}
L527:
	;
	v1628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+6)))
	if v1628 != 0 {
		goto L541
	} else {
		goto L542
	}
L528:
	;
	if v1596 != 0 {
		goto L20
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	v1624 = v27 + int32(32)
	if v1596 != 0 {
		v1631 = v1624
		v1632 = v1596
		goto L526
	} else {
		goto L540
	}
L531:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v1598 != 0 {
		goto L19
	} else {
		goto L532
	}
L532:
	;
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v1599 == int32(0) {
		goto L533
	} else {
		goto L534
	}
L533:
	;
	v1625 = v27 + int32(32)
	goto L527
L534:
	;
	goto L535
L535:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L4
	} else {
		goto L536
	}
L536:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L4
	} else {
		goto L537
	}
L537:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(_a_F_ProcessCopyOptions_46)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_47), v20+int32(272))
	mBase = m.M
	v1617 = m.ExcPending
	if v1617 != 0 {
		goto L4
	} else {
		goto L538
	}
L538:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(745), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L4
	} else {
		goto L539
	}
L539:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L540:
	;
	v1625 = v1624
	goto L527
L541:
	;
	v1629 = int32(_a_F_ProcessCopyOptions_48)
	goto L543
L542:
	;
	v1629 = int32(_a_F_ProcessCopyOptions_49)
	goto L543
L543:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1629
	v1631 = v1625
	v1632 = v1629
	goto L526
L544:
	;
	v1679 = F_strlen(m, v1632)
	mBase = m.M
	if v1679 != int32(1) {
		goto L18
	} else {
		goto L563
	}
L545:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v1662 == int32(0) {
		goto L557
	} else {
		goto L558
	}
L546:
	;
	v1637 = int32(0)
	v1641 = v1633 & int32(1)
	if v1641 != 0 {
		goto L549
	} else {
		goto L550
	}
L547:
	;
	goto L548
L548:
	;
	v1650 = F_strlen(m, v1634)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1650
	v1652 = int32(0)
	if v1633&int32(1) == v1652 {
		v1674 = v1652
		v1675 = v1634
		v1678 = v1650
		goto L544
	} else {
		goto L556
	}
L549:
	;
	v1642 = v1637
	goto L551
L550:
	;
	v1642 = int32(2)
	goto L551
L551:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1642
	v1644 = int32(_a_F_ProcessCopyOptions_50)
	v1645 = int32(_a_F_ProcessCopyOptions_51)
	if v1641 != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v1648 = v1645
	goto L554
L553:
	;
	v1648 = v1644
	goto L554
L554:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v1648
	if v1641 != 0 {
		v1659 = v1645
		v1661 = v1642
		goto L545
	} else {
		goto L555
	}
L555:
	;
	v1674 = v1637
	v1675 = v1644
	v1678 = v1642
	goto L544
L556:
	;
	v1659 = v1634
	v1661 = v1650
	goto L545
L557:
	;
	v1665 = int32(_a_F_ProcessCopyOptions_52)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v1665
	v1668 = v1665
	goto L559
L558:
	;
	v1668 = v1662
	goto L559
L559:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1670 == int32(0) {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1668
	goto L562
L561:
	;
	goto L562
L562:
	;
	v1674 = int32(1)
	v1675 = v1659
	v1678 = v1661
	goto L544
L563:
	;
	v1682 = int32(13)
	v1683 = F___strchrnul(m, v1632, v1682)
	mBase = m.M
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1683))))
	if v1685 == v1682 {
		goto L565
	} else {
		goto L566
	}
L564:
	;
	if v1689 != 0 {
		goto L17
	} else {
		goto L568
	}
L565:
	;
	v1689 = v1683
	goto L567
L566:
	;
	v1689 = int32(0)
	goto L567
L567:
	;
	goto L564
L568:
	;
	v1690 = int32(10)
	v1691 = F___strchrnul(m, v1632, v1690)
	mBase = m.M
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1691))))
	if v1693 == v1690 {
		goto L570
	} else {
		goto L571
	}
L569:
	;
	if v1697 != 0 {
		goto L17
	} else {
		goto L573
	}
L570:
	;
	v1697 = v1691
	goto L572
L571:
	;
	v1697 = int32(0)
	goto L572
L572:
	;
	goto L569
L573:
	;
	v1698 = int32(13)
	v1699 = F___strchrnul(m, v1675, v1698)
	mBase = m.M
	v1701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1699))))
	if v1701 == v1698 {
		goto L575
	} else {
		goto L576
	}
L574:
	;
	if v1705 != 0 {
		goto L16
	} else {
		goto L578
	}
L575:
	;
	v1705 = v1699
	goto L577
L576:
	;
	v1705 = int32(0)
	goto L577
L577:
	;
	goto L574
L578:
	;
	v1706 = int32(10)
	v1707 = F___strchrnul(m, v1675, v1706)
	mBase = m.M
	v1709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1707))))
	if v1709 == v1706 {
		goto L580
	} else {
		goto L581
	}
L579:
	;
	if v1713 != 0 {
		goto L16
	} else {
		goto L583
	}
L580:
	;
	v1713 = v1707
	goto L582
L581:
	;
	v1713 = int32(0)
	goto L582
L582:
	;
	goto L579
L583:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v1714 != 0 {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v1715 = F_strlen(m, v1714)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v1715
	v1717 = int32(13)
	v1718 = F___strchrnul(m, v1714, v1717)
	mBase = m.M
	v1720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1718))))
	if v1720 == v1717 {
		goto L588
	} else {
		goto L589
	}
L585:
	;
	goto L586
L586:
	;
	if v1674 == int32(0) {
		goto L597
	} else {
		goto L598
	}
L587:
	;
	if v1724 != 0 {
		goto L15
	} else {
		goto L591
	}
L588:
	;
	v1724 = v1718
	goto L590
L589:
	;
	v1724 = int32(0)
	goto L590
L590:
	;
	goto L587
L591:
	;
	v1725 = int32(10)
	v1726 = F___strchrnul(m, v1714, v1725)
	mBase = m.M
	v1728 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1726))))
	if v1728 == v1725 {
		goto L593
	} else {
		goto L594
	}
L592:
	;
	if v1732 != 0 {
		goto L15
	} else {
		goto L596
	}
L593:
	;
	v1732 = v1726
	goto L595
L594:
	;
	v1732 = int32(0)
	goto L595
L595:
	;
	goto L592
L596:
	;
	goto L586
L597:
	;
	v1735 = int32(_a_F_ProcessCopyOptions_53)
	v1736 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1632))))
	v1737 = int32(39)
	goto L603
L598:
	;
	goto L599
L599:
	;
	if v1597 != 0 {
		goto L626
	} else {
		goto L627
	}
L600:
	;
	if v1842 != 0 {
		goto L14
	} else {
		goto L625
	}
L601:
	;
	v1842 = int32(0)
	goto L600
L602:
	;
	v1820 = v1813
	v1822 = v1815
	goto L619
L603:
	;
	goto L610
L610:
	;
	v1776 = v1736 & int32(255)
	v1777 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[19])))
	if base.B2i32(v1776 == v1777)|int32(0) == int32(0) {
		goto L611
	} else {
		goto L612
	}
L611:
	;
	v1786 = v1735
	v1788 = v1737
	goto L614
L612:
	;
	v1806 = v1735
	v1808 = v1737
	goto L613
L613:
	;
	if v1808 == int32(0) {
		goto L601
	} else {
		goto L618
	}
L614:
	;
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v1786)))
	v1793 = v1792 ^ v1776*int32(16843009)
	v1796 = int32(-2139062144)
	if (int32(16843008)-v1793|v1793)&v1796 != v1796 {
		v1813 = v1786
		v1815 = v1788
		goto L602
	} else {
		goto L616
	}
L615:
	;
	v1806 = v1801
	v1808 = v1803
	goto L613
L616:
	;
	v1800 = int32(4)
	v1801 = v1786 + v1800
	v1803 = v1788 - v1800
	if base.Ui32(int32(3)) < base.Ui32(v1803) {
		v1786 = v1801
		v1788 = v1803
		goto L614
	} else {
		goto L617
	}
L617:
	;
	goto L615
L618:
	;
	v1813 = v1806
	v1815 = v1808
	goto L602
L619:
	;
	v1825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1820))))
	if v1736&int32(255) == v1825 {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	goto L601
L621:
	;
	v1842 = v1820
	goto L600
L622:
	;
	goto L623
L623:
	;
	v1827 = int32(1)
	v1830 = v1822 - v1827
	if v1830 != 0 {
		v1820 = v1820 + v1827
		v1822 = v1830
		goto L619
	} else {
		goto L624
	}
L624:
	;
	goto L620
L625:
	;
	goto L599
L626:
	;
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if v1843 != 0 {
		goto L13
	} else {
		goto L629
	}
L627:
	;
	goto L628
L628:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v1674 == int32(0) {
		goto L639
	} else {
		goto L640
	}
L629:
	;
	goto L628
L630:
	;
	if l2 == int32(0) {
		goto L9
	} else {
		goto L709
	}
L631:
	;
	if l2 == int32(0) {
		goto L697
	} else {
		goto L698
	}
L632:
	;
	if l2 == int32(0) {
		goto L692
	} else {
		goto L693
	}
L633:
	;
	if l2 != 0 {
		goto L11
	} else {
		goto L690
	}
L634:
	;
	if l2 != 0 {
		goto L678
	} else {
		goto L679
	}
L635:
	;
	v1957 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v1957 != 0 {
		goto L633
	} else {
		goto L671
	}
L636:
	;
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v1932 == int32(0) {
		goto L663
	} else {
		goto L664
	}
L637:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	v1913 = F_strlen(m, v1912)
	mBase = m.M
	if v1913 == int32(1) {
		goto L635
	} else {
		goto L658
	}
L638:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1890 == int32(0) {
		goto L636
	} else {
		goto L653
	}
L639:
	;
	if v1844 == int32(0) {
		goto L638
	} else {
		goto L642
	}
L640:
	;
	goto L641
L641:
	;
	v1868 = F_strlen(m, v1844)
	mBase = m.M
	if v1868 != int32(1) {
		goto L12
	} else {
		goto L647
	}
L642:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L4
	} else {
		goto L643
	}
L643:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L4
	} else {
		goto L644
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = int32(_a_F_ProcessCopyOptions_54)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_55), v20+int32(224))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L4
	} else {
		goto L645
	}
L645:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(822), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L4
	} else {
		goto L646
	}
L646:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L647:
	;
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1632))))
	v1872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1844))))
	if v1871 != v1872 {
		goto L637
	} else {
		goto L648
	}
L648:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L4
	} else {
		goto L649
	}
L649:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1880 = m.ExcPending
	if v1880 != 0 {
		goto L4
	} else {
		goto L650
	}
L650:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_56), int32(0))
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L4
	} else {
		goto L651
	}
L651:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(832), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L4
	} else {
		goto L652
	}
L652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L653:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L4
	} else {
		goto L654
	}
L654:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L4
	} else {
		goto L655
	}
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = int32(_a_F_ProcessCopyOptions_57)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_55), v20+int32(208))
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L4
	} else {
		goto L656
	}
L656:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(839), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L4
	} else {
		goto L657
	}
L657:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L658:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L4
	} else {
		goto L659
	}
L659:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L4
	} else {
		goto L660
	}
L660:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_58), int32(0))
	mBase = m.M
	v1926 = m.ExcPending
	if v1926 != 0 {
		goto L4
	} else {
		goto L661
	}
L661:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(844), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L4
	} else {
		goto L662
	}
L662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L663:
	;
	v1935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v1935 != int32(1) {
		goto L634
	} else {
		goto L666
	}
L664:
	;
	goto L665
L665:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L4
	} else {
		goto L667
	}
L666:
	;
	goto L665
L667:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L4
	} else {
		goto L668
	}
L668:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = int32(_a_F_ProcessCopyOptions_59)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_55), v20+int32(160))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L4
	} else {
		goto L669
	}
L669:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(851), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L4
	} else {
		goto L670
	}
L670:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L671:
	;
	if l2 != 0 {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	v1958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v1958&int32(1) != 0 {
		goto L11
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v1961 == int32(0) {
		goto L632
	} else {
		goto L676
	}
L675:
	;
	goto L674
L676:
	;
	if l2 != 0 {
		v2257 = v1844
		goto L8
	} else {
		goto L677
	}
L677:
	;
	goto L7
L678:
	;
	v1964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+48)))
	if v1964&int32(1) != 0 {
		goto L11
	} else {
		goto L681
	}
L679:
	;
	goto L680
L680:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v1967 == int32(0) {
		goto L682
	} else {
		goto L683
	}
L681:
	;
	goto L680
L682:
	;
	v1970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v1970 != int32(1) {
		goto L631
	} else {
		goto L685
	}
L683:
	;
	goto L684
L684:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L4
	} else {
		goto L686
	}
L685:
	;
	goto L684
L686:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L4
	} else {
		goto L687
	}
L687:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = int32(_a_F_ProcessCopyOptions_60)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_55), v20+int32(176))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L4
	} else {
		goto L688
	}
L688:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(866), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v1991 = m.ExcPending
	if v1991 != 0 {
		goto L4
	} else {
		goto L689
	}
L689:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L690:
	;
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v1992 != 0 {
		goto L7
	} else {
		goto L691
	}
L691:
	;
	goto L632
L692:
	;
	v1995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v1995&int32(1) != 0 {
		goto L7
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v1998 != 0 {
		goto L630
	} else {
		goto L696
	}
L695:
	;
	goto L694
L696:
	;
	v2229 = v1844
	goto L10
L697:
	;
	v2001 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+60)))
	if v2001&int32(1) != 0 {
		goto L7
	} else {
		goto L700
	}
L698:
	;
	goto L699
L699:
	;
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	if v2004 == int32(0) {
		goto L701
	} else {
		goto L702
	}
L700:
	;
	goto L699
L701:
	;
	v2008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v2008 != int32(1) {
		v2229 = int32(0)
		goto L10
	} else {
		goto L704
	}
L702:
	;
	goto L703
L703:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L4
	} else {
		goto L705
	}
L704:
	;
	goto L703
L705:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L4
	} else {
		goto L706
	}
L706:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = int32(_a_F_ProcessCopyOptions_61)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_55), v20+int32(192))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L4
	} else {
		goto L707
	}
L707:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(882), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L4
	} else {
		goto L708
	}
L708:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L709:
	;
	v2257 = v1844
	goto L8
L710:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L711:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L4
	} else {
		goto L712
	}
L712:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+352)) = v380
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_62), v20+int32(352))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L4
	} else {
		goto L713
	}
L713:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(415), int32(_a_F_ProcessCopyOptions_63))
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L4
	} else {
		goto L714
	}
L714:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L715:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2060 = m.ExcPending
	if v2060 != 0 {
		goto L4
	} else {
		goto L716
	}
L716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+304)) = int32(_a_F_ProcessCopyOptions_64)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_47), v20+int32(304))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L4
	} else {
		goto L717
	}
L717:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(735), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L4
	} else {
		goto L718
	}
L718:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L719:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L4
	} else {
		goto L720
	}
L720:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = int32(_a_F_ProcessCopyOptions_65)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_47), v20+int32(288))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L4
	} else {
		goto L721
	}
L721:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(740), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L4
	} else {
		goto L722
	}
L722:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L723:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L4
	} else {
		goto L724
	}
L724:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_66), int32(0))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L4
	} else {
		goto L725
	}
L725:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(767), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L4
	} else {
		goto L726
	}
L726:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L727:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2114 = m.ExcPending
	if v2114 != 0 {
		goto L4
	} else {
		goto L728
	}
L728:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_67), int32(0))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L4
	} else {
		goto L729
	}
L729:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(774), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L4
	} else {
		goto L730
	}
L730:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L731:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L4
	} else {
		goto L732
	}
L732:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_68), int32(0))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L4
	} else {
		goto L733
	}
L733:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(780), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L4
	} else {
		goto L734
	}
L734:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L735:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L4
	} else {
		goto L736
	}
L736:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_69), int32(0))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L4
	} else {
		goto L737
	}
L737:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(790), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L4
	} else {
		goto L738
	}
L738:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L739:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2162 = m.ExcPending
	if v2162 != 0 {
		goto L4
	} else {
		goto L740
	}
L740:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(v1631)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v2163
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_70), v20+int32(256))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L4
	} else {
		goto L741
	}
L741:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(808), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L4
	} else {
		goto L742
	}
L742:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L743:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2181 = m.ExcPending
	if v2181 != 0 {
		goto L4
	} else {
		goto L744
	}
L744:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = int32(_a_F_ProcessCopyOptions_71)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_47), v20+int32(240))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L4
	} else {
		goto L745
	}
L745:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(815), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L4
	} else {
		goto L746
	}
L746:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L747:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L4
	} else {
		goto L748
	}
L748:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_72), int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L4
	} else {
		goto L749
	}
L749:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(827), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L4
	} else {
		goto L750
	}
L750:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L751:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L4
	} else {
		goto L752
	}
L752:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(_a_F_ProcessCopyOptions_73)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(_a_F_ProcessCopyOptions_59)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v20)
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L4
	} else {
		goto L753
	}
L753:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(858), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L4
	} else {
		goto L754
	}
L754:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L755:
	;
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v2230&int32(1) == int32(0) {
		v2257 = v2229
		goto L8
	} else {
		goto L756
	}
L756:
	;
	goto L9
L757:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L4
	} else {
		goto L758
	}
L758:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = int32(_a_F_ProcessCopyOptions_32)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = int32(_a_F_ProcessCopyOptions_61)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v20+int32(128))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L4
	} else {
		goto L759
	}
L759:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(891), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2256 = m.ExcPending
	if v2256 != 0 {
		goto L4
	} else {
		goto L760
	}
L760:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L761:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L4
	} else {
		goto L859
	}
L762:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2496 = m.ExcPending
	if v2496 != 0 {
		goto L4
	} else {
		goto L855
	}
L763:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L4
	} else {
		goto L851
	}
L764:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2461 = m.ExcPending
	if v2461 != 0 {
		goto L4
	} else {
		goto L847
	}
L765:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L4
	} else {
		goto L843
	}
L766:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2421 = m.ExcPending
	if v2421 != 0 {
		goto L4
	} else {
		goto L839
	}
L767:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2400 = m.ExcPending
	if v2400 != 0 {
		goto L4
	} else {
		goto L835
	}
L768:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L4
	} else {
		goto L831
	}
L769:
	;
	if v2265 == int32(0) {
		goto L773
	} else {
		goto L774
	}
L770:
	;
	v2265 = v2259
	goto L772
L771:
	;
	v2265 = int32(0)
	goto L772
L772:
	;
	goto L769
L773:
	;
	if v1674 != 0 {
		goto L776
	} else {
		goto L777
	}
L774:
	;
	goto L775
L775:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L4
	} else {
		goto L827
	}
L776:
	;
	v2268 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2257))))
	v2269 = F___strchrnul(m, v1675, v2268)
	mBase = m.M
	v2271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2269))))
	if v2271 == v2268&int32(255) {
		goto L780
	} else {
		goto L781
	}
L777:
	;
	goto L778
L778:
	;
	if l2 == int32(0) {
		goto L784
	} else {
		goto L785
	}
L779:
	;
	if v2275 != 0 {
		goto L768
	} else {
		goto L783
	}
L780:
	;
	v2275 = v2269
	goto L782
L781:
	;
	v2275 = int32(0)
	goto L782
L782:
	;
	goto L779
L783:
	;
	goto L778
L784:
	;
	v2278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+5)))
	if v2278&int32(1) != 0 {
		goto L767
	} else {
		goto L787
	}
L785:
	;
	goto L786
L786:
	;
	if v1714 == int32(0) {
		goto L788
	} else {
		goto L789
	}
L787:
	;
	goto L786
L788:
	;
	if v1597 != 0 {
		goto L819
	} else {
		goto L820
	}
L789:
	;
	if l2 == int32(0) {
		goto L766
	} else {
		goto L790
	}
L790:
	;
	v2285 = F___strchrnul(m, v1714, v2258)
	mBase = m.M
	v2287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2285))))
	if v2287 == v2258&int32(255) {
		goto L792
	} else {
		goto L793
	}
L791:
	;
	if v2291 != 0 {
		goto L765
	} else {
		goto L795
	}
L792:
	;
	v2291 = v2285
	goto L794
L793:
	;
	v2291 = int32(0)
	goto L794
L794:
	;
	goto L791
L795:
	;
	if v1674 != 0 {
		goto L796
	} else {
		goto L797
	}
L796:
	;
	v2292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2257))))
	v2293 = F___strchrnul(m, v1714, v2292)
	mBase = m.M
	v2295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2293))))
	if v2295 == v2292&int32(255) {
		goto L800
	} else {
		goto L801
	}
L797:
	;
	goto L798
L798:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v1678 != v2300 {
		goto L788
	} else {
		goto L804
	}
L799:
	;
	if v2299 != 0 {
		goto L764
	} else {
		goto L803
	}
L800:
	;
	v2299 = v2293
	goto L802
L801:
	;
	v2299 = int32(0)
	goto L802
L802:
	;
	goto L799
L803:
	;
	goto L798
L804:
	;
	if v1678 == int32(0) {
		goto L806
	} else {
		goto L807
	}
L805:
	;
	if v2346 == int32(0) {
		goto L763
	} else {
		goto L818
	}
L806:
	;
	v2346 = int32(0)
	goto L805
L807:
	;
	goto L808
L808:
	;
	v2307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1675))))
	if v2307 != 0 {
		goto L809
	} else {
		goto L810
	}
L809:
	;
	v2308 = v1675
	v2309 = v1714
	v2310 = v1678
	v2311 = v2307
	goto L813
L810:
	;
	v2334 = v1714
	v2338 = int32(0)
	goto L811
L811:
	;
	v2339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2334))))
	v2346 = v2338 - v2339
	goto L805
L812:
	;
	v2334 = v2329
	v2338 = v2331
	goto L811
L813:
	;
	v2313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2309))))
	if base.B2i32(v2311 != v2313)|base.B2i32(v2313 == int32(0)) != 0 {
		v2329 = v2309
		v2331 = v2311
		goto L812
	} else {
		goto L815
	}
L814:
	;
	v2329 = v2323
	v2331 = int32(0)
	goto L812
L815:
	;
	v2319 = v2310 - int32(1)
	if v2319 == int32(0) {
		v2329 = v2309
		v2331 = v2311
		goto L812
	} else {
		goto L816
	}
L816:
	;
	v2322 = int32(1)
	v2323 = v2309 + v2322
	v2324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2308)+1)))
	if v2324 != 0 {
		v2308 = v2308 + v2322
		v2309 = v2323
		v2310 = v2319
		v2311 = v2324
		goto L813
	} else {
		goto L817
	}
L817:
	;
	goto L814
L818:
	;
	goto L788
L819:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v2349 != 0 {
		goto L762
	} else {
		goto L822
	}
L820:
	;
	goto L821
L821:
	;
	v2350 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	if v2350 != int64(0) {
		goto L823
	} else {
		goto L824
	}
L822:
	;
	goto L821
L823:
	;
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	if v2353 == int32(0) {
		goto L761
	} else {
		goto L826
	}
L824:
	;
	goto L825
L825:
	;
	m.G0 = v20 + int32(464)
	return
L826:
	;
	goto L825
L827:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2365 = m.ExcPending
	if v2365 != 0 {
		goto L4
	} else {
		goto L828
	}
L828:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = int32(_a_F_ProcessCopyOptions_65)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_74), v20+int32(112))
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L4
	} else {
		goto L829
	}
L829:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(899), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L4
	} else {
		goto L830
	}
L830:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L831:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2384 = m.ExcPending
	if v2384 != 0 {
		goto L4
	} else {
		goto L832
	}
L832:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(_a_F_ProcessCopyOptions_65)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_75), v20+int32(96))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L4
	} else {
		goto L833
	}
L833:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(908), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2396 = m.ExcPending
	if v2396 != 0 {
		goto L4
	} else {
		goto L834
	}
L834:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L835:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L4
	} else {
		goto L836
	}
L836:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_ProcessCopyOptions_32)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(_a_F_ProcessCopyOptions_76)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v20+int32(80))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L4
	} else {
		goto L837
	}
L837:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(917), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L4
	} else {
		goto L838
	}
L838:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L839:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L4
	} else {
		goto L840
	}
L840:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = int32(_a_F_ProcessCopyOptions_32)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(_a_F_ProcessCopyOptions_46)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v20-int32(-64))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L4
	} else {
		goto L841
	}
L841:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(927), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2438 = m.ExcPending
	if v2438 != 0 {
		goto L4
	} else {
		goto L842
	}
L842:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L843:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L4
	} else {
		goto L844
	}
L844:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(_a_F_ProcessCopyOptions_46)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_74), v20+int32(48))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L4
	} else {
		goto L845
	}
L845:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(935), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L4
	} else {
		goto L846
	}
L846:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L847:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L4
	} else {
		goto L848
	}
L848:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(_a_F_ProcessCopyOptions_46)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_75), v20+int32(32))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L4
	} else {
		goto L849
	}
L849:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(944), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2476 = m.ExcPending
	if v2476 != 0 {
		goto L4
	} else {
		goto L850
	}
L850:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L851:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2483 = m.ExcPending
	if v2483 != 0 {
		goto L4
	} else {
		goto L852
	}
L852:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_77), int32(0))
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L4
	} else {
		goto L853
	}
L853:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(952), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2492 = m.ExcPending
	if v2492 != 0 {
		goto L4
	} else {
		goto L854
	}
L854:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L855:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2499 = m.ExcPending
	if v2499 != 0 {
		goto L4
	} else {
		goto L856
	}
L856:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_78), int32(0))
	mBase = m.M
	v2503 = m.ExcPending
	if v2503 != 0 {
		goto L4
	} else {
		goto L857
	}
L857:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(958), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2508 = m.ExcPending
	if v2508 != 0 {
		goto L4
	} else {
		goto L858
	}
L858:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L859:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2515 = m.ExcPending
	if v2515 != 0 {
		goto L4
	} else {
		goto L860
	}
L860:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(_a_F_ProcessCopyOptions_79)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_ProcessCopyOptions_33)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(_a_F_ProcessCopyOptions_80)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_81), v20+int32(16))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L4
	} else {
		goto L861
	}
L861:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(966), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L4
	} else {
		goto L862
	}
L862:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L863:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L4
	} else {
		goto L864
	}
L864:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+148)) = int32(_a_F_ProcessCopyOptions_32)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = int32(_a_F_ProcessCopyOptions_60)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_34), v20+int32(144))
	mBase = m.M
	v2547 = m.ExcPending
	if v2547 != 0 {
		goto L4
	} else {
		goto L865
	}
L865:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(874), int32(_a_F_ProcessCopyOptions_7))
	mBase = m.M
	v2552 = m.ExcPending
	if v2552 != 0 {
		goto L4
	} else {
		goto L866
	}
L866:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L867:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L4
	} else {
		goto L868
	}
L868:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+336)) = v2563
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_82), v20+int32(336))
	mBase = m.M
	v2569 = m.ExcPending
	if v2569 != 0 {
		goto L4
	} else {
		goto L869
	}
L869:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_6), int32(424), int32(_a_F_ProcessCopyOptions_63))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L4
	} else {
		goto L870
	}
L870:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessPendingWrites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v55 int32
	_ = v55
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v104 int64
	_ = v104
	var v110 int64
	_ = v110
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
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
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	goto L2
L1:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L60
	}
L2:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[0]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if v156 != 0 {
		goto L47
	} else {
		goto L48
	}
L4:
	;
	return
L5:
	;
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[1]))
	if v12 <= int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v65 = m.T0[v64].(func(*base.Module) int32)(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L21
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[3]))
	if v16 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[4]))
	if base.I64_extend_i32_u(v16)*int64(1000)+v12 <= v20 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[5])))
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L12:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_errmsg(m, int32(_a_F_ProcessPendingWrites_0), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_ProcessPendingWrites_1), int32(2789), int32(_a_F_ProcessPendingWrites_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L1
L16:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[4]))
	if v44 < base.I64_extend_i32_u(int32(base.Ui32(v16)>>(uint(int32(1))%32)))*int64(1000)+v12 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = m.T0[v58].(func(*base.Module) int32)(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L6
L21:
	;
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = m.G0
	v71 = int32(16)
	v72 = v70 - v71
	m.G0 = v72
	F_gettimeofday(m, v72)
	mBase = m.M
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	v76 = int64(*(*int32)(unsafe.Add(mBase, uint32(v72)+8)))
	m.G0 = v72 + v71
	v84 = v76 + v75*int64(1000000) - int64(946684800000000)
	goto L25
L23:
	;
	goto L24
L24:
	;
	goto L3
L25:
	;
	v85 = int32(_a_F_ProcessPendingWrites_3)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[3]))
	if v87 <= int32(0) {
		v123 = v85
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_WalSndWait(m, int32(6), v123, int32(100663304))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L33
	}
L27:
	;
	v91 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[1]))
	if v91 <= int64(0) {
		v123 = v85
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[5])))
	v104 = base.I64_extend_i32_u(int32(base.Ui32(v87)>>(uint((v95^int32(-1))&int32(1))%32)))*int64(1000) + v91
	if v104 <= v84 {
		v122 = int32(0)
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v123 = v122
	goto L26
L30:
	;
	goto L29
L31:
	;
	v110 = v104 - v84
	if base.B2i32(int64(0) < v84)^base.B2i32(v110 < v104)|base.B2i32(int64(2147483646000) < v110) != 0 {
		v122 = int32(2147483647)
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v119 = base.I64_div_s(v110+int64(999), int64(1000))
	v122 = base.I32_wrap_i64(v119)
	goto L30
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = int32(0)
	goto L34
L34:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[6]))
	if v134 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[7]))
	if v138 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[7])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v150 = m.T0[v149].(func(*base.Module) int32)(m)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	if v150 == int32(0) {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	return
L47:
	;
	goto L46
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(1)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v159 == int32(0) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	if v162 == int32(0) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[8]))
	if v166 == v162 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v168 = m.G0
	v170 = v168 - int32(16)
	m.G0 = v170
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[9]))
	if v173 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v196 = F_pgmem_kill(m, v162, int32(23))
	mBase = m.M
	goto L47
L54:
	;
	m.G0 = v170 + int32(16)
	goto L46
L55:
	;
	v176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v170)+15)) = uint8(v176)
	goto L56
L56:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[10]))
	v184 = F_write(m, v180, v170+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v184 {
		goto L54
	} else {
		goto L58
	}
L57:
	;
	goto L54
L58:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[11]))
	if v188 == int32(27) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PushCopiedSnapshot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PushCopiedSnapshot[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(2)
	v13 = int32(72)
	v18 = v9<<(uint(v11)%32) + v13
	if int32(0) < v8 {
		v21 = (v8+v9)<<(uint(v11)%32) + v13
	} else {
		v21 = v18
	}
	v22 = F_MemoryContextAlloc(m, v7, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v26
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v28
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v30
		v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v32
		v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v34
		v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v36
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v22))) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = int64(0)
		v42 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v42
		*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v42
		v46 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v22)+30)) = uint8(v46)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v48 != 0 {
			v50 = v22 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v50
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v54 = v52 << (uint(int32(2)) % 32)
			if v54 == int32(0) {
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				base.MemoryCopy(m, v50, v57, v54)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v63 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
		} else {
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v66 == int32(1) {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				if v69 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
				} else {
					v72 = v22 + v18
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v76 = v74 << (uint(int32(2)) % 32)
					if v76 == int32(0) {
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v72, v79, v76)
					}
				}
			} else {
				v72 = v22 + v18
				*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v72
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v76 = v74 << (uint(int32(2)) % 32)
				if v76 == int32(0) {
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v72, v79, v76)
				}
			}
		}
		v86 = *(*int32)(unsafe.Add(mBase, _c_F_PushCopiedSnapshot[1]))
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
		F_PushActiveSnapshotWithLevel(m, v22, v87)
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			return
		}
	}
}
func F_p_isalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v3 == int32(1) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v6+v9<<(uint(int32(2))%32))))
			if base.Ui32(int32(127)) < base.Ui32(v13) {
				v63 = int32(1)
				return v63
			} else {
				return base.B2i32(base.Ui32(v13-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v13|int32(32)-int32(97)) < base.Ui32(int32(26)))
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+v30<<(uint(int32(2))%32))))
			if base.Ui32(int32(10)) <= base.Ui32(v34-int32(48)) {
				v39 = F_iswalpha(m, v34)
				mBase = m.M
				v43 = base.B2i32(v39 != int32(0))
			} else {
				v43 = int32(1)
			}
			return v43
		}
	} else {
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v47))))
		v63 = base.B2i32(base.Ui32(v49-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v49|int32(32)-int32(97)) < base.Ui32(int32(26)))
		return v63
	}
}
func F_p_isspecial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_p_isspecial[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12*int32(28))+uint32(_c_F_p_isspecial[1])))
	v18 = m.T0[v17].(func(*base.Module, int32) int32)(m, v6+v8)
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
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(1)
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_p_isspecial[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	goto L7
L6:
	;
	return int32(0)
L7:
	;
	if v28 != int32(6) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v31 != int32(1) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = v34
	goto L12
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = v35
	goto L12
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+v38<<(uint(int32(2))%32))))
	v45 = int32(_a_F_p_isspecial_0)
	v46 = int32(_a_F_p_isspecial_1)
	goto L13
L13:
	;
	v55 = v45 + (v46-v45)>>(uint(int32(3))%32)<<(uint(int32(2))%32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v42 == v56 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L6
L15:
	;
	return int32(1)
L16:
	;
	goto L17
L17:
	;
	v62 = base.B2i32(base.Ui32(v56) < base.Ui32(v42))
	if base.Ui32(v56) < base.Ui32(v42) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = v55 + int32(4)
	goto L20
L19:
	;
	v63 = v45
	goto L20
L20:
	;
	if base.Ui32(v56) < base.Ui32(v42) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v64 = v46
	goto L23
L22:
	;
	v64 = v55
	goto L23
L23:
	;
	if base.Ui32(v63) < base.Ui32(v64) {
		v45 = v63
		v46 = v64
		goto L13
	} else {
		goto L24
	}
L24:
	;
	goto L14
}
func F_p_isurlchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != int32(1) {
		v24 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9))))
		if base.Ui32((v11-int32(127))&int32(255)) < base.Ui32(int32(162)) {
			v24 = v2
		} else {
			switch v11 - int32(60) {
			case 0, 2, 32, 34, 36, 63, 64, 65:
				v24 = v2
			case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62:
				v24 = int32(1)
			default:
				if v11 == int32(34) {
					v24 = v2
				} else {
					v24 = int32(1)
				}
			}
		}
	}
	return v24
}
func F_pairingheap_GISTSearchItem_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v4 < v10 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v101
L2:
	;
	v101 = int32(0) - v69
	goto L1
L3:
	;
	v13 = int32(32)
	v20 = v10
	v21 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	v84 = int32(-1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v86 == v84 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	v27 = v21 << (uint(int32(4)) % 32)
	v28 = l0 + v13 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v29 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v73 = v21 + int32(1)
	if v73 < v71 {
		v20 = v71
		v21 = v73
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v27)+40)))
	if v33 != 0 {
		v71 = v20
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v36 = v27 + (l1 + v13)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(-1)
L13:
	;
	return int32(1)
L14:
	;
	goto L15
L15:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v46 = int64(9223372036854775807)
	v47 = base.I64_reinterpret_f64(v40) & v46
	v50 = base.I64_reinterpret_f64(v41) & v46
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v50) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v69 != 0 {
		goto L2
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v69 = int32(0) - v60&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v50))|base.F64_lt(v40, v41))
	goto L17
L19:
	;
	v60 = base.B2i32(base.Ui64(v47) < base.Ui64(int64(9218868437227405313)))
	goto L18
L20:
	;
	goto L21
L21:
	;
	v55 = int32(1)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v47))|base.F64_gt(v40, v41) != 0 {
		v69 = v55
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v60 = v55
	goto L18
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v71 = v70
	goto L8
L24:
	;
	goto L7
L25:
	;
	return int32(0)
L26:
	;
	if v85 == int32(-1) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v85 == int32(-1) {
		v101 = v84
		goto L1
	} else {
		goto L30
	}
L29:
	;
	return int32(1)
L30:
	;
	goto L25
}
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, v8, l1, v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = base.B2i32(v11 < int32(0))
			if v11 < int32(0) {
				v15 = v8
			} else {
				v15 = l1
			}
			if v11 < int32(0) {
				v16 = l1
			} else {
				v16 = v8
			}
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v15
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v16
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15
			v23 = v16
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
			return
		}
	} else {
		v23 = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
		v28 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
		return
	}
}
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_palloc_aligned[0]))
	if base.Ui32(l1) <= base.Ui32(int32(8)) {
		v9 = F_MemoryContextAllocExtended(m, v6, l0, l2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	} else {
		v15 = F_MemoryContextAllocExtended(m, v6, l0+l1, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v22 = (v15 + l1 + int32(7)) & (int32(0) - l1)
			v24 = v22 - int32(8)
			*(*int64)(unsafe.Add(mBase, uint32(v24))) = base.I64_extend_i32_u(l1)<<(uint(int64(5))%64) | base.I64_extend_i32_u(v24-v15)<<(uint(int64(34))%64) | int64(6)
			return v22
		}
	}
}
func F_parseNameAndArgTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = F_pstrdup(m, l0)
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
	v22 = v18
	v30 = v7
	goto L4
L3:
	;
	m.G0 = v16 + int32(16)
	return v406
L4:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v35 != int32(34) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v18)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L109
	}
L6:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v394 = int32(1)
	v22 = v22 + v394
	v30 = v30 ^ v394
	goto L4
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v63)
	v65 = F_stringToQualifiedNameList(m, v18, l5)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	if (base.B2i32(v35 != int32(40))|v30)&int32(1) == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v47 = F_errsave_start(m, l5)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v22 = v22 + int32(1)
	goto L4
L15:
	;
	if v47 == int32(0) {
		v406 = v7
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_0), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(1924), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v406 = v7
	goto L3
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
	if v65 == int32(0) {
		v406 = v7
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(1)
	v71 = v22 + v70
	v72 = F_strlen(m, v71)
	mBase = m.M
	v81 = v72 + v70
	goto L22
L22:
	;
	v89 = v81 - int32(1)
	v90 = v22 + v89
	if v81 < int32(3) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v104 != int32(41) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	goto L23
L25:
	;
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90))))
	goto L26
L26:
	;
	if base.B2i32(v93 == int32(32))|base.B2i32(base.Ui32((v93-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v81 = v89
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v107 = int32(0)
	v108 = F_errsave_start(m, l5)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v124)
	v127 = v124
	v129 = v71
	v139 = v7
	goto L36
L31:
	;
	if v108 == int32(0) {
		v406 = v107
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_3), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(1942), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v406 = v107
	goto L3
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v127
	v143 = v129
	goto L38
L38:
	;
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143))))
	goto L40
L39:
	;
	v166 = int32(0)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v168 == v166 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if base.B2i32(v156 == int32(32))|base.B2i32(base.Ui32((v156-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v143 = v143 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v139 == int32(0) {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v191 = v168
	v193 = v143
	v197 = v166
	v201 = v166
	goto L51
L45:
	;
	v174 = int32(0)
	v175 = F_errsave_start(m, l5)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v175 == int32(0) {
		v406 = v174
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_4), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(1961), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v406 = v174
	goto L3
L51:
	;
	if v191 != int32(34) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v191 = v393
	v193 = v193 + int32(1)
	v197 = v389
	v201 = v390
	goto L51
L54:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(100) <= v361 {
		goto L101
	} else {
		goto L102
	}
L55:
	;
	if v191 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	goto L57
L57:
	;
	v389 = v197 ^ int32(1)
	v390 = v201
	goto L53
L58:
	;
	v259 = v193 - int32(1)
	if base.Ui32(v259) < base.Ui32(v143) {
		goto L76
	} else {
		goto L77
	}
L59:
	;
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v251)
	v253 = int32(1)
	v256 = v253
	v257 = v193 + v253
	goto L58
L60:
	;
	if (base.B2i32(v191 != int32(44))|v197)&int32(1)|v201 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v226 = int32(0)
	if (v197|base.B2i32(v201 != v226))&int32(1) == v226 {
		v256 = v226
		v257 = v193
		goto L58
	} else {
		goto L70
	}
L63:
	;
	if v197&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v389 = int32(1)
	v390 = v201
	goto L53
L65:
	;
	goto L66
L66:
	;
	v217 = int32(0)
	switch v191 - int32(40) {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L69
	}
L67:
	;
	v389 = v217
	v390 = v201 - int32(1)
	goto L53
L68:
	;
	v389 = v217
	v390 = v201 + int32(1)
	goto L53
L69:
	;
	switch v191 - int32(91) {
	case 0:
		goto L68
	default:
		v389 = v217
		v390 = v201
		goto L53
	case 2:
		goto L67
	}
L70:
	;
	v234 = int32(0)
	v235 = F_errsave_start(m, l5)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v235 == int32(0) {
		v406 = v234
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_5), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(1993), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v406 = v234
	goto L3
L76:
	;
	if l1 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v261 = v259
	goto L78
L78:
	;
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v261))))
	goto L80
L79:
	;
	goto L76
L80:
	;
	if base.B2i32(v274 == int32(32))|base.B2i32(base.Ui32((v274-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v286)
	v289 = v261 - int32(1)
	if base.Ui32(v143) <= base.Ui32(v289) {
		v261 = v289
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v356 = F_parseTypeString(m, v143, v16+int32(12), v16+int32(8), l5)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L99
	}
L84:
	;
	v309 = v143
	v310 = int32(_a_F_parseNameAndArgTypes_6)
	goto L86
L85:
	;
	if v347 != 0 {
		goto L83
	} else {
		goto L98
	}
L86:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v313 == v314 {
		v336 = v313
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v347 = int32(0)
	goto L85
L88:
	;
	v338 = int32(1)
	if v336 != 0 {
		v309 = v309 + v338
		v310 = v310 + v338
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v313-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v324 = v313 | int32(32)
	goto L92
L91:
	;
	v324 = v313
	goto L92
L92:
	;
	if base.Ui32((v314-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v333 = v314 | int32(32)
	goto L95
L94:
	;
	v333 = v314
	goto L95
L95:
	;
	if v324 == v333 {
		v336 = v324
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v347 = v324 - v333
	goto L85
L97:
	;
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	goto L54
L99:
	;
	if v356 != 0 {
		goto L54
	} else {
		goto L100
	}
L100:
	;
	v406 = int32(0)
	goto L3
L101:
	;
	v364 = int32(0)
	v365 = F_errsave_start(m, l5)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v361<<(uint(int32(2))%32)))) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v127 = v386 + int32(1)
	v129 = v257
	v139 = v256
	goto L36
L104:
	;
	if v365 == int32(0) {
		v406 = v364
		goto L3
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_7), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2029), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v406 = v364
	goto L3
L109:
	;
	v406 = int32(1)
	goto L3
}
func F_parse_filename_for_nontemp_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v20-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v141 = v5
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
		v34 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(4294967295))
		mBase = m.M
		v35 = base.I32_wrap_i64(v34)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
		if v37 != 0 {
			v141 = v5
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if base.B2i32(v35 == int32(0))|base.B2i32(l0 == v40) != 0 {
				v141 = v5
			} else {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
				if v43 != int32(95) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
					v94 = v43
					v95 = v40
					if v94&int32(255) == int32(46) {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
						if base.Ui32((v100-int32(58))&int32(255)) < base.Ui32(int32(247)) {
							v141 = v5
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
							v111 = v95 + int32(1)
							v116 = F_strtox_2(m, v111, v12+int32(8), int32(10), int64(4294967295))
							mBase = m.M
							v117 = base.I32_wrap_i64(v116)
							v119 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
							if v119 != 0 {
								v141 = v5
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if base.B2i32(v117 == int32(0))|base.B2i32(v111 == v122) != 0 {
									v141 = v5
								} else {
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
									v128 = v117
									v129 = v125
									if v129&int32(255) != 0 {
										v141 = v5
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
										v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
										v141 = int32(1)
									}
								}
							}
						}
					} else {
						v128 = v5
						v129 = v94
						if v129&int32(255) != 0 {
							v141 = v5
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
							v141 = int32(1)
						}
					}
				} else {
					v49 = v40 + int32(1)
					v51 = v12 + int32(12)
					v54 = int32(3)
					v57 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_0), v49, v54)
					mBase = m.M
					if v57 == int32(0) {
						v79 = v54
						v80 = int32(1)
						if v51 == int32(0) {
							v86 = v79
						} else {
							v83 = v79
							v84 = v80
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
							v86 = v83
						}
					} else {
						v61 = int32(2)
						v65 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_1), v49, v61)
						mBase = m.M
						if v65 == int32(0) {
							v79 = v61
							v80 = v61
							if v51 == int32(0) {
								v86 = v79
							} else {
								v83 = v79
								v84 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
								v86 = v83
							}
						} else {
							v68 = int32(4)
							v71 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_2), v49, v68)
							mBase = m.M
							if v71 == int32(0) {
								if v51 != 0 {
									v83 = v68
									v84 = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
									v86 = v83
								} else {
									v86 = v68
								}
							} else {
								v76 = int32(0)
								if v51 == v76 {
									v86 = v76
								} else {
									v83 = v76
									v84 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
									v86 = v83
								}
							}
						}
					}
					if v86 <= int32(0) {
						v141 = v5
					} else {
						v92 = v86 + v40 + int32(1)
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v94 = v93
						v95 = v92
						if v94&int32(255) == int32(46) {
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
							if base.Ui32((v100-int32(58))&int32(255)) < base.Ui32(int32(247)) {
								v141 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
								v111 = v95 + int32(1)
								v116 = F_strtox_2(m, v111, v12+int32(8), int32(10), int64(4294967295))
								mBase = m.M
								v117 = base.I32_wrap_i64(v116)
								v119 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
								if v119 != 0 {
									v141 = v5
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if base.B2i32(v117 == int32(0))|base.B2i32(v111 == v122) != 0 {
										v141 = v5
									} else {
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
										v128 = v117
										v129 = v125
										if v129&int32(255) != 0 {
											v141 = v5
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
											v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
											v141 = int32(1)
										}
									}
								}
							}
						} else {
							v128 = v5
							v129 = v94
							if v129&int32(255) != 0 {
								v141 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
								v141 = int32(1)
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	return v141
}
func F_parse_new_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v106 = v15
			m.G0 = v10 + int32(16)
			return v106
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if base.Ui32(v22) < base.Ui32(int32(192)) {
				v100 = v22
				v101 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
				v106 = v101
				m.G0 = v10 + int32(16)
				return v106
			} else {
				if base.Ui32(v22) <= base.Ui32(int32(223)) {
					v30 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(14))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 < int32(0) {
							v106 = v30
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
							v100 = v34 | v22<<(uint(int32(8))%32) - int32(_a_F_parse_new_len_0)
							v101 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
							v106 = v101
						}
						m.G0 = v10 + int32(16)
						return v106
					}
				} else {
					if v22 == int32(255) {
						v46 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(13))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v46 < int32(0) {
								v106 = v46
								m.G0 = v10 + int32(16)
								return v106
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
								v54 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(12))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 < int32(0) {
										v106 = v54
										m.G0 = v10 + int32(16)
										return v106
									} else {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
										v62 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(11))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											if v62 < int32(0) {
												v106 = v62
												m.G0 = v10 + int32(16)
												return v106
											} else {
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
												v70 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(10))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													if v70 < int32(0) {
														v106 = v70
														m.G0 = v10 + int32(16)
														return v106
													} else {
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
														v75 = int32(8)
														v88 = v74 | (v58<<(uint(v75)%32)|v50<<(uint(int32(16))%32)|v66)<<(uint(v75)%32)
														v92 = int32(1)
														if base.Ui32(v88) < base.Ui32(int32(16777217)) {
															v100 = v88
															v101 = v92
															*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
															v106 = v101
															m.G0 = v10 + int32(16)
															return v106
														} else {
															F_px_debug(m, int32(_a_F_parse_new_len_1), int32(0))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																v106 = int32(-100)
																m.G0 = v10 + int32(16)
																return v106
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
						v88 = int32(1) << (uint(v22) % 32)
						v92 = int32(2)
						if base.Ui32(v88) < base.Ui32(int32(16777217)) {
							v100 = v88
							v101 = v92
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
							v106 = v101
							m.G0 = v10 + int32(16)
							return v106
						} else {
							F_px_debug(m, int32(_a_F_parse_new_len_1), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								v106 = int32(-100)
								m.G0 = v10 + int32(16)
								return v106
							}
						}
					}
				}
			}
		}
	}
}
func F_parser_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 < v3) != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 == int32(0) {
			return
		} else {
			v11 = F_pg_mbstrlen_with_len(m, v8, l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v15 = F_errposition(m, v11+int32(1))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var __phi133 int32
	_ = __phi133
	var v135 int32
	_ = v135
	var __phi135 int32
	_ = __phi135
	var v136 int32
	_ = v136
	var __phi136 int32
	_ = __phi136
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	v5 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v5
	v20 = F_lookup_ts_config_cache(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v23 = F_lookup_ts_parser_cache(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = F_FunctionCall2Coll(m, v23+int32(28), int32(0), l2, l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+28)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v14)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v14)+36)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v14)+44)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = int32(0)
	goto L5
L5:
	;
	v59 = F_FunctionCall3Coll(m, v23+int32(56), int32(0), v28, v12+int32(-8), v12+int32(-4))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v254 = F_FunctionCall1Coll(m, v23+int32(84), int32(0), v28)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L47
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v64 = int32(0)
	if base.B2i32(v61 < int32(2047))|base.B2i32(v59 <= v64) == v64 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if int32(0) < v59 {
		goto L5
	} else {
		goto L46
	}
L9:
	;
	v71 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v94 = F_palloc(m, int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	if v71 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_parsetext_0), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(2047)
	F_errdetail(m, int32(_a_F_parsetext_1), v14)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_parsetext_2), int32(389), int32(_a_F_parsetext_3))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v59
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	if v99 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v94
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v105
	v110 = F_LexizeExec(m, v12+int32(-56), v103)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v94
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v94
	goto L19
L23:
	;
	if v110 == int32(0) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v119 = v110
	goto L25
L25:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v125 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v129 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L8
L27:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	__phi133 = v119
	__phi135 = v132
	__phi136 = v119 + int32(4)
	v133 = __phi133
	v135 = __phi135
	v136 = __phi136
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v119)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L43
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v144 == v135 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L29
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v135 << (uint(int32(1)) % 32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v152 = F_repalloc(m, v149, v135<<(uint(int32(5))%32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+2)))
	if v155&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v152
	goto L34
L36:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v158 + int32(1)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v164 = int32(4)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v168 = F_strlen(m, v167)
	mBase = m.M
	*(*uint16)(unsafe.Add(mBase, uint32(v162+v163<<(uint(v164)%32))+2)) = uint16(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	*(*int32)(unsafe.Add(mBase, uint32(v170+v171<<(uint(v164)%32))+12)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133))))
	*(*uint16)(unsafe.Add(mBase, uint32(v177+v178<<(uint(v164)%32))+4)) = uint16(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+2)))
	v191 = v189 & int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v184+v185<<(uint(v164)%32)))) = uint16(v191)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v198 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v164)%32))+6)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v205 = int32(_a_F_parsetext_4)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v205 <= v206 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v209 = v205
	goto L41
L40:
	;
	v209 = v206
	goto L41
L41:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v200+v201<<(uint(v164)%32))+8)) = uint16(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v213 = v211 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v213
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	if v217 != 0 {
		__phi133 = v133 + int32(8)
		__phi135 = v213
		__phi136 = v133 + int32(12)
		v133 = __phi133
		v135 = __phi135
		v136 = __phi136
		goto L30
	} else {
		goto L42
	}
L42:
	;
	goto L31
L43:
	;
	v236 = F_LexizeExec(m, v12+int32(-56), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v236 != 0 {
		v119 = v236
		goto L25
	} else {
		goto L45
	}
L45:
	;
	goto L26
L46:
	;
	goto L6
L47:
	;
	m.G0 = v14 - int32(-64)
	return
}
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v85 float32
	_ = v85
	var v89 float64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
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
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v169 float64
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v175 int32
	_ = v175
	var v177 float64
	_ = v177
	var v179 float64
	_ = v179
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v198 float64
	_ = v198
	var v207 float64
	_ = v207
	var v210 float64
	_ = v210
	var v218 float64
	_ = v218
	var v224 float64
	_ = v224
	var v225 int32
	_ = v225
	var v228 float64
	_ = v228
	var v239 float64
	_ = v239
	var v242 float64
	_ = v242
	var v250 float64
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v268 float64
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v289 float64
	_ = v289
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = int64(0)
	if l7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = float64(0.995)
	goto L3
L2:
	;
	v29 = float64(0.005)
	goto L3
L3:
	;
	v36 = F_get_restriction_variable(m, l0, l3, l4, v21-int32(-64), v21+int32(60), v21+int32(59))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v21 + int32(96)
	return v289
L5:
	;
	return float64(0)
L6:
	;
	if v36 == int32(0) {
		v289 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+59)))
	if v42 != int32(1) {
		v268 = v29
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v272 == int32(0) {
		v289 = v268
		goto L4
	} else {
		goto L76
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(7) {
		v268 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+24)))
	if v49 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v268 = float64(0)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v53&int32(-9) != int32(17) {
		v268 = v29
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	switch v62 - int32(17) {
	case 0:
		goto L18
	case 1, 3, 4, 5, 6, 7:
		v268 = v29
		goto L8
	case 2:
		goto L16
	case 8:
		v77 = v62
		v78 = int32(98)
		v79 = int32(664)
		v80 = int32(667)
		goto L15
	default:
		goto L17
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v81 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v77 = int32(25)
	v78 = int32(254)
	v79 = int32(255)
	v80 = int32(257)
	goto L15
L17:
	;
	if v62 != int32(1042) {
		v268 = v29
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v77 = v62
	v78 = int32(1955)
	v79 = int32(1957)
	v80 = int32(1960)
	goto L15
L19:
	;
	v77 = v62
	v78 = int32(1054)
	v79 = int32(1058)
	v80 = int32(1061)
	goto L15
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v85 = *(*float32)(unsafe.Add(mBase, uint32(v82+v83)+8))
	v89 = base.F64_promote_f32(v85)
	goto L22
L21:
	;
	v89 = float64(0)
	goto L22
L22:
	;
	v94 = F_pattern_fixed_prefix(m, v45, l6, l5, v21+int32(52), v21+int32(40))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v96 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v94 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 == v77 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v77
	goto L24
L27:
	;
	if v96 == int32(0) {
		v268 = v250
		goto L8
	} else {
		goto L73
	}
L28:
	;
	if l7 != 0 {
		goto L68
	} else {
		goto L69
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v109 = int32(0)
	v112 = F_var_eq_const(m, v21-int32(-64), v78, l5, v108, v109, int32(1), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if l2 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v239 = v112
	goto L28
L33:
	;
	v116 = l2
	goto L35
L34:
	;
	v114 = F_get_opcode(m, l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v118 = v21 + int32(8)
	F_fmgr_info(m, v116, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v116 = v114
	goto L35
L37:
	;
	v122 = v21 - int32(-64)
	v126 = F_histogram_selectivity(m, v122, v118, l5, v58, int32(1), v21+int32(36))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if int32(99) < v128 {
		v207 = v126
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v210 = float64(0.0001)
	if base.F64_lt(v207, v210) != 0 {
		v218 = v210
		goto L64
	} else {
		goto L65
	}
L40:
	;
	if v94 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = m.G0
	v135 = v133 - int32(32)
	m.G0 = v135
	v137 = F_get_opcode(m, v80)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	v190 = float64(1)
	goto L43
L43:
	;
	v191 = *(*float64)(unsafe.Add(mBase, uint32(v21)+40))
	v192 = base.F64_mul(v190, v191)
	if base.F64_lt(v126, float64(0)) != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v140 = v135 + int32(4)
	F_fmgr_info(m, v137, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v144 = int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v148 = F_ineq_histogram_selectivity(m, l0, v122, v80, v140, v144, v144, l5, v146, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if base.F64_lt(v148, float64(0)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v154 = F_get_opcode(m, v79)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v179 = float64(0.005)
	goto L49
L49:
	;
	m.G0 = v135 + int32(32)
	v190 = v179
	goto L43
L50:
	;
	F_fmgr_info(m, v154, v140)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v158 = F_make_greater_string(m, v96, v140, l5)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v158 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v164 = F_ineq_histogram_selectivity(m, l0, v122, v79, v140, v160, v160, l5, v162, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	v169 = v148
	goto L55
L55:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	v171 = int32(0)
	v174 = F_var_eq_const(m, v122, v78, l5, v170, v171, int32(1), v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	v169 = base.F64_add(base.F64_add(v148, v164), float64(-1))
	goto L55
L57:
	;
	if base.F64_lt(v174, v169) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v177 = v169
	goto L60
L59:
	;
	v177 = v174
	goto L60
L60:
	;
	v179 = v177
	goto L49
L61:
	;
	v207 = v192
	goto L39
L62:
	;
	goto L63
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v198 = base.F64_div(base.F64_convert_i32_s(v195), float64(100))
	v207 = base.F64_add(base.F64_mul(v126, v198), base.F64_mul(v192, base.F64_sub(float64(1), v198)))
	goto L39
L64:
	;
	v224 = F_mcv_selectivity(m, v21-int32(-64), v21+int32(8), l5, v58, int32(1), v21)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	if base.F64_gt(v207, float64(0.9999)) == int32(0) {
		v218 = v207
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v218 = float64(0.9999)
	goto L64
L67:
	;
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v239 = base.F64_add(v224, base.F64_mul(v218, base.F64_sub(base.F64_sub(float64(1), v89), v228)))
	goto L28
L68:
	;
	v242 = base.F64_sub(base.F64_sub(float64(1), v239), v89)
	goto L70
L69:
	;
	v242 = v239
	goto L70
L70:
	;
	if base.F64_lt(v242, float64(0)) != 0 {
		v250 = float64(0)
		goto L27
	} else {
		goto L71
	}
L71:
	;
	if base.F64_gt(v242, float64(1)) == int32(0) {
		v250 = v242
		goto L27
	} else {
		goto L72
	}
L72:
	;
	v250 = float64(1)
	goto L27
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	F_pfree(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v96)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v268 = v250
	goto L8
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	m.T0[v275].(func(*base.Module, int32))(m, v272)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v289 = v268
	goto L4
}
func F_pcb_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = F_geterrcode(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(67371461) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v7 == int32(0) {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v10 < int32(0) {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					if v13 == int32(0) {
						return
					} else {
						v16 = F_pg_mbstrlen_with_len(m, v13, v10)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							v20 = F_errposition(m, v16+int32(1))
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
			}
		}
	}
}
func F_pfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_c_F_pfree[0])))
	m.T0[v9].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_pgsql_version(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text(m, int32(_a_F_pgsql_version_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgstattuple_approx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstattuple_approx_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstattuple_approx_1), int32(221), int32(_a_F_pgstattuple_approx_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v26 = F_pgstattuple_approx_internal(m, v3, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		}
	}
}
func F_pgstattuple_approx_v1_5(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pgstattuple_approx_internal(m, v2, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_plainto_tsquery_byid(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14028(m, l0, int32(1), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_portuguese_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v603 int32
	_ = v603
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v660 int32
	_ = v660
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v717 int32
	_ = v717
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v774 int32
	_ = v774
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1261 int32
	_ = v1261
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1275 int32
	_ = v1275
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v7
	goto L2
L1:
	;
	return v1275
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v10
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= v10 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v56
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v81 < v71 {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	goto L3
L5:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = v64
	goto L2
L6:
	;
	if v56 <= v55 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v10
	v55 = v10
	v56 = v15
	goto L6
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v10))))
	v21 = v19 - int32(227)
	v22 = int32(0)
	if base.B2i32(v21 == v22)|base.B2i32(v21 == int32(18)) == v22 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v31 = F_find_among(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_0), int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v35
	switch v31 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L5
	}
L12:
	;
	v48 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v42 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_2))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = v35
	v56 = v39
	goto L6
L15:
	;
	if int32(0) <= v42 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v1275 = v42
	goto L1
L17:
	;
	if int32(0) <= v48 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v1275 = v48
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55 + int32(1)
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v571 < v71 {
		goto L168
	} else {
		goto L169
	}
L21:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v558)+8)) = v557
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v351 < v71 {
		goto L102
	} else {
		goto L103
	}
L23:
	;
	if v124 != 0 {
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v83 = v71
	goto L26
L25:
	;
	v83 = v81
	goto L26
L26:
	;
	goto L28
L27:
	;
	v124 = v119
	goto L23
L28:
	;
	if v71 == v83 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v119 = int32(0)
	goto L27
L30:
	;
	v124 = int32(-1)
	goto L23
L31:
	;
	goto L32
L32:
	;
	v95 = int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v71))))
	if int32(250) < v98 {
		v119 = v95
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v100 = v98 - int32(97)
	if v100 < int32(0) {
		v119 = v95
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v100)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v106)>>(uint(v100&int32(7))%32))&int32(1) == int32(0) {
		v119 = v95
		goto L27
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71 + int32(1)
	goto L36
L36:
	;
	goto L29
L37:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v134 < v125 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v239 < v125 {
		goto L72
	} else {
		goto L73
	}
L39:
	;
	if v174 != 0 {
		goto L38
	} else {
		goto L54
	}
L40:
	;
	v136 = v125
	goto L42
L41:
	;
	v136 = v134
	goto L42
L42:
	;
	goto L44
L43:
	;
	v174 = v171
	goto L39
L44:
	;
	if v125 == v136 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v171 = int32(0)
	goto L43
L46:
	;
	v174 = int32(-1)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147+v125))))
	if int32(250) < v149 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125 + int32(1)
	goto L53
L50:
	;
	v151 = v149 - int32(97)
	if v151 < int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v154 = int32(1)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v151)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v158)>>(uint(v151&int32(7))%32))&v154 != 0 {
		v171 = v154
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	goto L45
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v183 < v182 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v223 < int32(0) {
		goto L38
	} else {
		goto L70
	}
L56:
	;
	v185 = v182
	goto L58
L57:
	;
	v185 = v183
	goto L58
L58:
	;
	v192 = v182
	goto L60
L59:
	;
	v223 = v203
	goto L55
L60:
	;
	if v192 == v185 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v223 = int32(-1)
	goto L55
L63:
	;
	goto L64
L64:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196+v192))))
	if int32(250) < v198 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v215 = v192 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v215
	v192 = v215
	goto L60
L66:
	;
	v200 = v198 - int32(97)
	if v200 < int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v203 = int32(1)
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v200)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v207)>>(uint(v200&int32(7))%32))&v203 != 0 {
		goto L59
	} else {
		goto L68
	}
L68:
	;
	goto L65
L70:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v557 = v226 + v223
	goto L21
L71:
	;
	if v282 != 0 {
		goto L22
	} else {
		goto L85
	}
L72:
	;
	v241 = v125
	goto L74
L73:
	;
	v241 = v239
	goto L74
L74:
	;
	goto L76
L75:
	;
	v282 = v277
	goto L71
L76:
	;
	if v125 == v241 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v277 = int32(0)
	goto L75
L78:
	;
	v282 = int32(-1)
	goto L71
L79:
	;
	goto L80
L80:
	;
	v253 = int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254+v125))))
	if int32(250) < v256 {
		v277 = v253
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v258 = v256 - int32(97)
	if v258 < int32(0) {
		v277 = v253
		goto L75
	} else {
		goto L82
	}
L82:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v258)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v264)>>(uint(v258&int32(7))%32))&int32(1) == int32(0) {
		v277 = v253
		goto L75
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v125 + int32(1)
	goto L84
L84:
	;
	goto L77
L85:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v292 < v291 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v335 < int32(0) {
		goto L22
	} else {
		goto L100
	}
L87:
	;
	v294 = v291
	goto L89
L88:
	;
	v294 = v292
	goto L89
L89:
	;
	v300 = v291
	goto L91
L90:
	;
	v335 = int32(1)
	goto L86
L91:
	;
	if v300 == v294 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v335 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307+v300))))
	if int32(250) < v309 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v311 = v309 - int32(97)
	if v311 < int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v311)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v317)>>(uint(v311&int32(7))%32))&int32(1) == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v326 = v300 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v326
	v300 = v326
	goto L91
L100:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v557 = v338 + v335
	goto L21
L101:
	;
	if v391 != 0 {
		goto L20
	} else {
		goto L116
	}
L102:
	;
	v353 = v71
	goto L104
L103:
	;
	v353 = v351
	goto L104
L104:
	;
	goto L106
L105:
	;
	v391 = v388
	goto L101
L106:
	;
	if v71 == v353 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v388 = int32(0)
	goto L105
L108:
	;
	v391 = int32(-1)
	goto L101
L109:
	;
	goto L110
L110:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364+v71))))
	if int32(250) < v366 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71 + int32(1)
	goto L115
L112:
	;
	v368 = v366 - int32(97)
	if v368 < int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v371 = int32(1)
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v368)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v375)>>(uint(v368&int32(7))%32))&v371 != 0 {
		v388 = v371
		goto L105
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	goto L107
L116:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v401 < v392 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v506 < v392 {
		goto L151
	} else {
		goto L152
	}
L118:
	;
	if v441 != 0 {
		goto L117
	} else {
		goto L133
	}
L119:
	;
	v403 = v392
	goto L121
L120:
	;
	v403 = v401
	goto L121
L121:
	;
	goto L123
L122:
	;
	v441 = v438
	goto L118
L123:
	;
	if v392 == v403 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v438 = int32(0)
	goto L122
L125:
	;
	v441 = int32(-1)
	goto L118
L126:
	;
	goto L127
L127:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v414+v392))))
	if int32(250) < v416 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392 + int32(1)
	goto L132
L129:
	;
	v418 = v416 - int32(97)
	if v418 < int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v421 = int32(1)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v418)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v425)>>(uint(v418&int32(7))%32))&v421 != 0 {
		v438 = v421
		goto L122
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	goto L124
L133:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v450 < v449 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v490 < int32(0) {
		goto L117
	} else {
		goto L149
	}
L135:
	;
	v452 = v449
	goto L137
L136:
	;
	v452 = v450
	goto L137
L137:
	;
	v459 = v449
	goto L139
L138:
	;
	v490 = v470
	goto L134
L139:
	;
	if v459 == v452 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v490 = int32(-1)
	goto L134
L142:
	;
	goto L143
L143:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463+v459))))
	if int32(250) < v465 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v482 = v459 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v482
	v459 = v482
	goto L139
L145:
	;
	v467 = v465 - int32(97)
	if v467 < int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v470 = int32(1)
	v474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v467)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v474)>>(uint(v467&int32(7))%32))&v470 != 0 {
		goto L138
	} else {
		goto L147
	}
L147:
	;
	goto L144
L149:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v557 = v493 + v490
	goto L21
L150:
	;
	if v549 != 0 {
		goto L20
	} else {
		goto L164
	}
L151:
	;
	v508 = v392
	goto L153
L152:
	;
	v508 = v506
	goto L153
L153:
	;
	goto L155
L154:
	;
	v549 = v544
	goto L150
L155:
	;
	if v392 == v508 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v544 = int32(0)
	goto L154
L157:
	;
	v549 = int32(-1)
	goto L150
L158:
	;
	goto L159
L159:
	;
	v520 = int32(1)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521+v392))))
	if int32(250) < v523 {
		v544 = v520
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v525 = v523 - int32(97)
	if v525 < int32(0) {
		v544 = v520
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v525)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v531)>>(uint(v525&int32(7))%32))&int32(1) == int32(0) {
		v544 = v520
		goto L154
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v392 + int32(1)
	goto L163
L163:
	;
	goto L156
L164:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v551 <= v550 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v557 = v550 + int32(1)
	goto L21
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v71
	v792 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v792
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v792
	if v792-int32(2) <= v71 {
		goto L231
	} else {
		goto L232
	}
L167:
	;
	if v611 < int32(0) {
		goto L166
	} else {
		goto L182
	}
L168:
	;
	v573 = v71
	goto L170
L169:
	;
	v573 = v571
	goto L170
L170:
	;
	v580 = v71
	goto L172
L171:
	;
	v611 = v591
	goto L167
L172:
	;
	if v580 == v573 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v611 = int32(-1)
	goto L167
L175:
	;
	goto L176
L176:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v580))))
	if int32(250) < v586 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v603 = v580 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v603
	v580 = v603
	goto L172
L178:
	;
	v588 = v586 - int32(97)
	if v588 < int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v591 = int32(1)
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v588)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v595)>>(uint(v588&int32(7))%32))&v591 != 0 {
		goto L171
	} else {
		goto L180
	}
L180:
	;
	goto L177
L182:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v615 = v614 + v611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v615
	v626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v626 < v615 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v669 < int32(0) {
		goto L166
	} else {
		goto L197
	}
L184:
	;
	v628 = v615
	goto L186
L185:
	;
	v628 = v626
	goto L186
L186:
	;
	v634 = v615
	goto L188
L187:
	;
	v669 = int32(1)
	goto L183
L188:
	;
	if v634 == v628 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v669 = int32(-1)
	goto L183
L191:
	;
	goto L192
L192:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641+v634))))
	if int32(250) < v643 {
		goto L187
	} else {
		goto L193
	}
L193:
	;
	v645 = v643 - int32(97)
	if v645 < int32(0) {
		goto L187
	} else {
		goto L194
	}
L194:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v645)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v651)>>(uint(v645&int32(7))%32))&int32(1) == int32(0) {
		goto L187
	} else {
		goto L195
	}
L195:
	;
	v660 = v634 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v660
	v634 = v660
	goto L188
L197:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v673 = v672 + v669
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v673
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v675)+4)) = v673
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v685 < v684 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v725 < int32(0) {
		goto L166
	} else {
		goto L213
	}
L199:
	;
	v687 = v684
	goto L201
L200:
	;
	v687 = v685
	goto L201
L201:
	;
	v694 = v684
	goto L203
L202:
	;
	v725 = v705
	goto L198
L203:
	;
	if v694 == v687 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v725 = int32(-1)
	goto L198
L206:
	;
	goto L207
L207:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698+v694))))
	if int32(250) < v700 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v717 = v694 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v717
	v694 = v717
	goto L203
L209:
	;
	v702 = v700 - int32(97)
	if v702 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v705 = int32(1)
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v702)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v709)>>(uint(v702&int32(7))%32))&v705 != 0 {
		goto L202
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v729 = v728 + v725
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v729
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v740 < v729 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v783 < int32(0) {
		goto L166
	} else {
		goto L228
	}
L215:
	;
	v742 = v729
	goto L217
L216:
	;
	v742 = v740
	goto L217
L217:
	;
	v748 = v729
	goto L219
L218:
	;
	v783 = int32(1)
	goto L214
L219:
	;
	if v748 == v742 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v783 = int32(-1)
	goto L214
L222:
	;
	goto L223
L223:
	;
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v755+v748))))
	if int32(250) < v757 {
		goto L218
	} else {
		goto L224
	}
L224:
	;
	v759 = v757 - int32(97)
	if v759 < int32(0) {
		goto L218
	} else {
		goto L225
	}
L225:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v759)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v765)>>(uint(v759&int32(7))%32))&int32(1) == int32(0) {
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v774 = v748 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v774
	v748 = v774
	goto L219
L228:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v786))) = v787 + v783
	goto L166
L229:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1143
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1143
	v1148 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_3), int32(4))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L10
	} else {
		goto L333
	}
L230:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1110
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1110
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1110 <= v1113 {
		goto L229
	} else {
		goto L325
	}
L231:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1067
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1069)+8))
	if v1070 <= v1067 {
		goto L313
	} else {
		goto L314
	}
L232:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v800 = int32(1)
	v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798+v792-v800))))
	if base.B2i32(v802&int32(224) != int32(96))|base.B2i32(v800<<(uint(v802)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_4) == int32(0)) != 0 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v816 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_5), int32(45))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	if v816 == int32(0) {
		goto L231
	} else {
		goto L235
	}
L235:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v820
	switch v816 - int32(1) {
	case 0:
		goto L244
	case 1:
		goto L243
	case 2:
		goto L242
	case 3:
		goto L241
	case 4:
		goto L240
	case 5:
		goto L239
	case 6:
		goto L238
	case 7:
		goto L237
	case 8:
		goto L236
	default:
		goto L230
	}
L236:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v1044)+8))
	if v820 < v1045 {
		goto L231
	} else {
		goto L307
	}
L237:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v1008)))
	if v820 < v1009 {
		goto L231
	} else {
		goto L296
	}
L238:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v966)))
	if v820 < v967 {
		goto L231
	} else {
		goto L286
	}
L239:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v931)))
	if v820 < v932 {
		goto L231
	} else {
		goto L276
	}
L240:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v820 < v859 {
		goto L231
	} else {
		goto L257
	}
L241:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v849)))
	if v820 < v850 {
		goto L231
	} else {
		goto L254
	}
L242:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	if v820 < v841 {
		goto L231
	} else {
		goto L251
	}
L243:
	;
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)))
	if v820 < v832 {
		goto L231
	} else {
		goto L248
	}
L244:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)))
	if v820 < v825 {
		goto L231
	} else {
		goto L245
	}
L245:
	;
	v827 = F_slice_del(m, l0)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	if int32(0) <= v827 {
		goto L230
	} else {
		goto L247
	}
L247:
	;
	v1275 = v827
	goto L1
L248:
	;
	v836 = F_slice_from_s(m, l0, int32(3), int32(_a_F_portuguese_ISO_8859_1_stem_6))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	if int32(0) <= v836 {
		goto L230
	} else {
		goto L250
	}
L250:
	;
	v1275 = v836
	goto L1
L251:
	;
	v845 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_7))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L252
	}
L252:
	;
	if int32(0) <= v845 {
		goto L230
	} else {
		goto L253
	}
L253:
	;
	v1275 = v845
	goto L1
L254:
	;
	v854 = F_slice_from_s(m, l0, int32(4), int32(_a_F_portuguese_ISO_8859_1_stem_8))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L255
	}
L255:
	;
	if int32(0) <= v854 {
		goto L230
	} else {
		goto L256
	}
L256:
	;
	v1275 = v854
	goto L1
L257:
	;
	v861 = F_slice_del(m, l0)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L10
	} else {
		goto L258
	}
L258:
	;
	if v861 < int32(0) {
		v1275 = v861
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v865
	v868 = v865 - int32(1)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v868 <= v869 {
		goto L230
	} else {
		goto L260
	}
L260:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v873 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v871+v868))))
	if base.B2i32(v873&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v873)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_9) == int32(0)) != 0 {
		goto L230
	} else {
		goto L261
	}
L261:
	;
	v887 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_10), int32(4))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L10
	} else {
		goto L262
	}
L262:
	;
	if v887 == int32(0) {
		goto L230
	} else {
		goto L263
	}
L263:
	;
	v891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v891
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v893)))
	if v891 < v894 {
		goto L230
	} else {
		goto L264
	}
L264:
	;
	v896 = F_slice_del(m, l0)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L265
	}
L265:
	;
	if v896 < int32(0) {
		v1275 = v896
		goto L1
	} else {
		goto L266
	}
L266:
	;
	if v887 != int32(1) {
		goto L230
	} else {
		goto L267
	}
L267:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v902
	v904 = int32(2)
	v906 = int32(0)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v902-v909 < v904 {
		v919 = v906
		goto L269
	} else {
		goto L270
	}
L268:
	;
	if v919 == int32(0) {
		goto L230
	} else {
		goto L272
	}
L269:
	;
	goto L268
L270:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v915 = F_memcmp(m, v912+v902-v904, int32(_a_F_portuguese_ISO_8859_1_stem_11), v904)
	mBase = m.M
	if v915 != 0 {
		v919 = v906
		goto L269
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v902 - v904
	v919 = int32(1)
	goto L269
L272:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v922
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v924)))
	if v922 < v925 {
		goto L230
	} else {
		goto L273
	}
L273:
	;
	v927 = F_slice_del(m, l0)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L10
	} else {
		goto L274
	}
L274:
	;
	if int32(0) <= v927 {
		goto L230
	} else {
		goto L275
	}
L275:
	;
	v1275 = v927
	goto L1
L276:
	;
	v934 = F_slice_del(m, l0)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L10
	} else {
		goto L277
	}
L277:
	;
	if v934 < int32(0) {
		v1275 = v934
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v938
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v938-int32(3) <= v940 {
		goto L230
	} else {
		goto L279
	}
L279:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v944+v938-int32(1)))))
	switch v948 - int32(101) {
	case 0, 7:
		goto L280
	default:
		goto L230
	}
L280:
	;
	v953 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_12), int32(3))
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L10
	} else {
		goto L281
	}
L281:
	;
	if v953 == int32(0) {
		goto L230
	} else {
		goto L282
	}
L282:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v957
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v959)))
	if v957 < v960 {
		goto L230
	} else {
		goto L283
	}
L283:
	;
	v962 = F_slice_del(m, l0)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L10
	} else {
		goto L284
	}
L284:
	;
	if int32(0) <= v962 {
		goto L230
	} else {
		goto L285
	}
L285:
	;
	v1275 = v962
	goto L1
L286:
	;
	v969 = F_slice_del(m, l0)
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L10
	} else {
		goto L287
	}
L287:
	;
	if v969 < int32(0) {
		v1275 = v969
		goto L1
	} else {
		goto L288
	}
L288:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v973
	v976 = v973 - int32(1)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v976 <= v977 {
		goto L230
	} else {
		goto L289
	}
L289:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v979+v976))))
	if base.B2i32(v981&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v981)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_13) == int32(0)) != 0 {
		goto L230
	} else {
		goto L290
	}
L290:
	;
	v995 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_14), int32(3))
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L10
	} else {
		goto L291
	}
L291:
	;
	if v995 == int32(0) {
		goto L230
	} else {
		goto L292
	}
L292:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v999
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v1001)))
	if v999 < v1002 {
		goto L230
	} else {
		goto L293
	}
L293:
	;
	v1004 = F_slice_del(m, l0)
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L10
	} else {
		goto L294
	}
L294:
	;
	if int32(0) <= v1004 {
		goto L230
	} else {
		goto L295
	}
L295:
	;
	v1275 = v1004
	goto L1
L296:
	;
	v1011 = F_slice_del(m, l0)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L10
	} else {
		goto L297
	}
L297:
	;
	if v1011 < int32(0) {
		v1275 = v1011
		goto L1
	} else {
		goto L298
	}
L298:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1015
	v1017 = int32(2)
	v1019 = int32(0)
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1015-v1022 < v1017 {
		v1032 = v1019
		goto L300
	} else {
		goto L301
	}
L299:
	;
	if v1032 == int32(0) {
		goto L230
	} else {
		goto L303
	}
L300:
	;
	goto L299
L301:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1028 = F_memcmp(m, v1025+v1015-v1017, int32(_a_F_portuguese_ISO_8859_1_stem_15), v1017)
	mBase = m.M
	if v1028 != 0 {
		v1032 = v1019
		goto L300
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1015 - v1017
	v1032 = int32(1)
	goto L300
L303:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1035
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)))
	if v1035 < v1038 {
		goto L230
	} else {
		goto L304
	}
L304:
	;
	v1040 = F_slice_del(m, l0)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L10
	} else {
		goto L305
	}
L305:
	;
	if int32(0) <= v1040 {
		goto L230
	} else {
		goto L306
	}
L306:
	;
	v1275 = v1040
	goto L1
L307:
	;
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v820 <= v1047 {
		goto L231
	} else {
		goto L308
	}
L308:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1049+v820-int32(1)))))
	if v1053 != int32(101) {
		goto L231
	} else {
		goto L309
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v820 - int32(1)
	v1061 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_16))
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L10
	} else {
		goto L310
	}
L310:
	;
	if int32(0) <= v1061 {
		goto L230
	} else {
		goto L311
	}
L311:
	;
	v1275 = v1061
	goto L1
L312:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1100
	v1102 = F_slice_del(m, l0)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L10
	} else {
		goto L323
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1067
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1070
	v1077 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_17), int32(120))
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L10
	} else {
		goto L316
	}
L314:
	;
	v1081 = v1067
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1081
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1081
	v1087 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_18), int32(7))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L10
	} else {
		goto L318
	}
L316:
	;
	if v1077 != 0 {
		goto L312
	} else {
		goto L317
	}
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1073
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1081 = v1080
	goto L315
L318:
	;
	if v1087 == int32(0) {
		goto L229
	} else {
		goto L319
	}
L319:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1091
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+8))
	if v1091 < v1094 {
		goto L229
	} else {
		goto L320
	}
L320:
	;
	v1096 = F_slice_del(m, l0)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L10
	} else {
		goto L321
	}
L321:
	;
	if int32(0) <= v1096 {
		goto L229
	} else {
		goto L322
	}
L322:
	;
	v1275 = v1096
	goto L1
L323:
	;
	if v1102 < int32(0) {
		v1275 = v1102
		goto L1
	} else {
		goto L324
	}
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1073
	goto L230
L325:
	;
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1116 = v1115 + v1110
	v1119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116-int32(1)))))
	if v1119 != int32(105) {
		goto L229
	} else {
		goto L326
	}
L326:
	;
	v1123 = v1110 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1123
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1123
	if v1123 <= v1113 {
		goto L229
	} else {
		goto L327
	}
L327:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116-int32(2)))))
	if v1129 != int32(99) {
		goto L229
	} else {
		goto L328
	}
L328:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1132)+8))
	if v1110 <= v1133 {
		goto L229
	} else {
		goto L329
	}
L329:
	;
	v1135 = F_slice_del(m, l0)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L10
	} else {
		goto L330
	}
L330:
	;
	if v1135 < int32(0) {
		v1275 = v1135
		goto L1
	} else {
		goto L331
	}
L331:
	;
	goto L229
L332:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1219
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1224 = v1219
	v1226 = v1221
	goto L354
L333:
	;
	if v1148 == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1152
	switch v1148 - int32(1) {
	case 0:
		goto L336
	case 1:
		goto L335
	default:
		goto L332
	}
L335:
	;
	v1210 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_19))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L10
	} else {
		goto L352
	}
L336:
	;
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1156)+8))
	if v1152 < v1157 {
		goto L332
	} else {
		goto L337
	}
L337:
	;
	v1159 = F_slice_del(m, l0)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L10
	} else {
		goto L338
	}
L338:
	;
	if v1159 < int32(0) {
		v1275 = v1159
		goto L1
	} else {
		goto L339
	}
L339:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1163
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1163 <= v1165 {
		goto L332
	} else {
		goto L340
	}
L340:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1168 = v1167 + v1163
	v1170 = v1168 - int32(1)
	v1171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170))))
	if v1171 != int32(117) {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1199
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+8))
	if v1199 < v1202 {
		goto L332
	} else {
		goto L349
	}
L342:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1163
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170))))
	if v1186 != int32(105) {
		goto L332
	} else {
		goto L346
	}
L343:
	;
	v1175 = v1163 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1175
	if v1175 <= v1165 {
		goto L342
	} else {
		goto L344
	}
L344:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168-int32(2)))))
	if v1181 == int32(103) {
		v1199 = v1175
		goto L341
	} else {
		goto L345
	}
L345:
	;
	goto L342
L346:
	;
	v1190 = v1163 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1190
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1190
	if v1190 <= v1165 {
		goto L332
	} else {
		goto L347
	}
L347:
	;
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168-int32(2)))))
	if v1196 != int32(99) {
		goto L332
	} else {
		goto L348
	}
L348:
	;
	v1199 = v1190
	goto L341
L349:
	;
	v1204 = F_slice_del(m, l0)
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L10
	} else {
		goto L350
	}
L350:
	;
	if int32(0) <= v1204 {
		goto L332
	} else {
		goto L351
	}
L351:
	;
	v1275 = v1204
	goto L1
L352:
	;
	if v1210 < int32(0) {
		v1275 = v1210
		goto L1
	} else {
		goto L353
	}
L353:
	;
	goto L332
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1224
	v1230 = v1224 + int32(1)
	if v1230 < v1226 {
		goto L360
	} else {
		goto L361
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1219
	v1275 = int32(1)
	goto L1
L356:
	;
	goto L355
L357:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1224 = v1270
	v1226 = v1269
	goto L354
L358:
	;
	if v1261 <= v1259 {
		goto L356
	} else {
		goto L372
	}
L359:
	;
	v1240 = F_find_among(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_20), int32(3))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L10
	} else {
		goto L364
	}
L360:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232+v1230))))
	if v1234 == int32(126) {
		goto L359
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1224
	v1259 = v1224
	v1261 = v1226
	goto L358
L363:
	;
	goto L362
L364:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1242
	switch v1240 - int32(1) {
	case 0:
		goto L367
	case 1:
		goto L366
	case 2:
		goto L365
	default:
		goto L357
	}
L365:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1259 = v1242
	v1261 = v1258
	goto L358
L366:
	;
	v1254 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_21))
	mBase = m.M
	v1255 = m.ExcPending
	if v1255 != 0 {
		goto L10
	} else {
		goto L370
	}
L367:
	;
	v1248 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_22))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L10
	} else {
		goto L368
	}
L368:
	;
	if int32(0) <= v1248 {
		goto L357
	} else {
		goto L369
	}
L369:
	;
	v1275 = v1248
	goto L1
L370:
	;
	if int32(0) <= v1254 {
		goto L357
	} else {
		goto L371
	}
L371:
	;
	v1275 = v1254
	goto L1
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1259 + int32(1)
	goto L357
}
func F_pow(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v92 int32
	_ = v92
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 float64
	_ = v167
	var v173 int32
	_ = v173
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 float64
	_ = v192
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 float64
	_ = v218
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v233 int32
	_ = v233
	var v234 float64
	_ = v234
	var v238 int64
	_ = v238
	var v243 float64
	_ = v243
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v250 float64
	_ = v250
	var v251 float64
	_ = v251
	var v253 float64
	_ = v253
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v273 float64
	_ = v273
	var v276 float64
	_ = v276
	var v279 float64
	_ = v279
	var v283 float64
	_ = v283
	var v286 float64
	_ = v286
	var v291 float64
	_ = v291
	var v294 float64
	_ = v294
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
	var v301 float64
	_ = v301
	var v306 float64
	_ = v306
	var v307 float64
	_ = v307
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v336 float64
	_ = v336
	var v338 float64
	_ = v338
	var v350 float64
	_ = v350
	var v352 float64
	_ = v352
	var v353 int32
	_ = v353
	var v355 float64
	_ = v355
	var v358 float64
	_ = v358
	var v359 float64
	_ = v359
	var v360 float64
	_ = v360
	var v362 float64
	_ = v362
	var v365 float64
	_ = v365
	var v369 float64
	_ = v369
	var v370 float64
	_ = v370
	var v373 float64
	_ = v373
	var v376 float64
	_ = v376
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v386 int64
	_ = v386
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v395 float64
	_ = v395
	var v396 int64
	_ = v396
	var v401 int64
	_ = v401
	var v410 float64
	_ = v410
	var v416 int64
	_ = v416
	var v417 float64
	_ = v417
	var v418 float64
	_ = v418
	var v419 float64
	_ = v419
	var v423 float64
	_ = v423
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v443 float64
	_ = v443
	var v444 float64
	_ = v444
	var v451 float64
	_ = v451
	var v454 float64
	_ = v454
	var v458 float64
	_ = v458
	var v467 float64
	_ = v467
	var v468 float64
	_ = v468
	var v480 float64
	_ = v480
	var v483 float64
	_ = v483
	v11 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v24 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
	v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l1)) >> (uint(int64(52)) % 64)))
	v29 = int32(2047)
	v30 = v28 & v29
	v32 = v30 - int32(1086)
	v33 = base.I64_reinterpret_f64(l1)
	v34 = base.I64_reinterpret_f64(l0)
	if base.B2i32(base.Ui32(int32(-129)) < base.Ui32(v32))&base.B2i32(base.Ui32(int32(-2046)) <= base.Ui32(v24-v29)) != 0 {
		v213 = v34
		v215 = v11
		v216 = int64(-134217728)
		v218 = base.F64_reinterpret_i64(v33 & v216)
		v220 = v213 - int64(4604531861337669632)
		v221 = int64(52)
		v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
		v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
		v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
		v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
		v238 = v213 - v220&int64(-4503599627370496)
		v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
		v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
		v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
		v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
		v251 = base.F64_add(v247, v250)
		v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
		v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
		v256 = base.F64_add(base.F64_mul(v223, v253), v255)
		v257 = base.F64_add(v251, v256)
		v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
		v263 = base.F64_mul(v251, v262)
		v264 = base.F64_mul(v247, v262)
		v268 = base.F64_mul(v247, v264)
		v269 = base.F64_add(v257, v268)
		v273 = base.F64_mul(v251, v263)
		v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
		v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
		v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
		v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
		v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
		v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
		v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
		v299 = base.F64_add(v269, v298)
		v301 = base.F64_add(v298, base.F64_sub(v269, v299))
		*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
		v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
		v307 = base.F64_mul(v218, v306)
		v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
		v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
		if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
			v353 = v320
			v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
			v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
			v359 = base.F64_add(base.F64_mul(v307, v355), v358)
			v360 = base.F64_sub(v359, v358)
			v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
			v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
			v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
			v370 = base.F64_mul(v369, v369)
			v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
			v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
			v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
			v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
			v386 = base.I64_reinterpret_f64(v359)
			v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
			v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
			v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
			v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
			v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
			if v353 == int32(0) {
				if v386&int64(2147483648) == int64(0) {
					v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
					v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
				} else {
					v416 = v401 + int64(4602678819172646912)
					v417 = base.F64_reinterpret_i64(v416)
					v418 = base.F64_mul(v417, v395)
					v419 = base.F64_add(v418, v417)
					if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
						v423 = float64(2.2250738585072014e-308)
						v425 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
						v432 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
						if base.F64_lt(v419, float64(0)) != 0 {
							v443 = float64(-1)
						} else {
							v443 = float64(1)
						}
						v444 = base.F64_add(v419, v443)
						v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
						if base.F64_eq(v451, float64(0)) != 0 {
							v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
						} else {
							v454 = v451
						}
						v458 = v454
					} else {
						v458 = v419
					}
					v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
				}
				v480 = v467
			} else {
				v468 = base.F64_reinterpret_i64(v401)
				v480 = base.F64_add(base.F64_mul(v468, v395), v468)
			}
		} else {
			if base.Ui32(v320) < base.Ui32(v325) {
				v336 = base.F64_add(v307, float64(1))
				if v215 != 0 {
					v338 = base.F64_neg(v336)
				} else {
					v338 = v336
				}
				v480 = v338
			} else {
				if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
					v353 = int32(0)
					v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
					v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
					v359 = base.F64_add(base.F64_mul(v307, v355), v358)
					v360 = base.F64_sub(v359, v358)
					v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
					v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
					v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
					v370 = base.F64_mul(v369, v369)
					v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
					v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
					v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
					v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
					v386 = base.I64_reinterpret_f64(v359)
					v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
					v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
					v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
					v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
					v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
					if v353 == int32(0) {
						if v386&int64(2147483648) == int64(0) {
							v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
							v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
						} else {
							v416 = v401 + int64(4602678819172646912)
							v417 = base.F64_reinterpret_i64(v416)
							v418 = base.F64_mul(v417, v395)
							v419 = base.F64_add(v418, v417)
							if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
								v423 = float64(2.2250738585072014e-308)
								v425 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
								v432 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
								if base.F64_lt(v419, float64(0)) != 0 {
									v443 = float64(-1)
								} else {
									v443 = float64(1)
								}
								v444 = base.F64_add(v419, v443)
								v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
								if base.F64_eq(v451, float64(0)) != 0 {
									v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
								} else {
									v454 = v451
								}
								v458 = v454
							} else {
								v458 = v419
							}
							v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
						}
						v480 = v467
					} else {
						v468 = base.F64_reinterpret_i64(v401)
						v480 = base.F64_add(base.F64_mul(v468, v395), v468)
					}
				} else {
					if base.I64_reinterpret_f64(v307) < int64(0) {
						v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
						mBase = m.M
						v480 = v350
					} else {
						v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
						mBase = m.M
						v480 = v352
					}
				}
			}
		}
		v483 = v480
	} else {
		if base.Ui64(v33<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
			v48 = float64(1)
			if v34 == int64(4607182418800017408) {
				v483 = v48
			} else {
				v52 = v33 << (uint(int64(1)) % 64)
				if v52 == int64(0) {
					v483 = v48
				} else {
					v58 = v34 << (uint(int64(1)) % 64)
					if base.B2i32(base.Ui64(v52) < base.Ui64(int64(-9007199254740991)))&base.B2i32(base.Ui64(v58) <= base.Ui64(int64(-9007199254740992))) == int32(0) {
						v483 = base.F64_add(l0, l1)
					} else {
						if v58 == int64(9214364837600034816) {
							v483 = v48
						} else {
							if base.B2i32(v33 < int64(0))^base.B2i32(base.Ui64(v58) < base.Ui64(int64(9214364837600034816))) != 0 {
								v74 = float64(0)
							} else {
								v74 = base.F64_mul(l1, l1)
							}
							v483 = v74
						}
					}
				}
			}
		} else {
			if base.Ui64(v34<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
				v81 = base.F64_mul(l0, l0)
				if v34 < int64(0) {
					v92 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v92) < base.Ui32(int32(1023)) {
						v116 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v92) {
							v116 = int32(2)
						} else {
							v99 = int64(1)
							v103 = v99 << (uint(base.I64_extend_i32_u(int32(1075)-v92)) % 64)
							if (v103-v99)&v33 != int64(0) {
								v116 = int32(0)
							} else {
								if v33&v103 == int64(0) {
									v114 = int32(2)
								} else {
									v114 = int32(1)
								}
								v116 = v114
							}
						}
					}
					if v116 == int32(1) {
						v119 = base.F64_neg(v81)
					} else {
						v119 = v81
					}
					v120 = v119
				} else {
					v120 = v81
				}
				if int64(0) <= v33 {
					v483 = v120
				} else {
					v124 = base.F64_div(float64(1), v120)
					v126 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v126-int32(16))+8)) = v124
					v483 = v124
				}
			} else {
				if v34 < int64(0) {
					v140 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v140) < base.Ui32(int32(1023)) {
						v164 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v140) {
							v164 = int32(2)
						} else {
							v147 = int64(1)
							v151 = v147 << (uint(base.I64_extend_i32_u(int32(1075)-v140)) % 64)
							if (v151-v147)&v33 != int64(0) {
								v164 = int32(0)
							} else {
								if v33&v151 == int64(0) {
									v162 = int32(2)
								} else {
									v162 = int32(1)
								}
								v164 = v162
							}
						}
					}
					if v164 == int32(0) {
						v167 = base.F64_sub(l0, l0)
						v483 = base.F64_div(v167, v167)
					} else {
						if v164 == int32(1) {
							v173 = int32(_a_F_pow_0)
						} else {
							v173 = int32(0)
						}
						v179 = base.I64_reinterpret_f64(l0) & int64(9223372036854775807)
						v180 = v24 & int32(2047)
						v181 = v173
						if base.Ui32(v32) <= base.Ui32(int32(-129)) {
							if v179 == int64(4607182418800017408) {
								v483 = float64(1)
							} else {
								if base.Ui32(v30) <= base.Ui32(int32(957)) {
									if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
										v192 = l1
									} else {
										v192 = base.F64_neg(l1)
									}
									v483 = base.F64_add(v192, float64(1))
								} else {
									if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
										v202 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
										mBase = m.M
										v483 = v202
									} else {
										v205 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
										mBase = m.M
										v483 = v205
									}
								}
							}
						} else {
							if v180 != 0 {
								v213 = v179
								v215 = v181
							} else {
								v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
								v215 = v181
							}
							v216 = int64(-134217728)
							v218 = base.F64_reinterpret_i64(v33 & v216)
							v220 = v213 - int64(4604531861337669632)
							v221 = int64(52)
							v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
							v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
							v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
							v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
							v238 = v213 - v220&int64(-4503599627370496)
							v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
							v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
							v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
							v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
							v251 = base.F64_add(v247, v250)
							v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
							v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
							v256 = base.F64_add(base.F64_mul(v223, v253), v255)
							v257 = base.F64_add(v251, v256)
							v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
							v263 = base.F64_mul(v251, v262)
							v264 = base.F64_mul(v247, v262)
							v268 = base.F64_mul(v247, v264)
							v269 = base.F64_add(v257, v268)
							v273 = base.F64_mul(v251, v263)
							v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
							v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
							v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
							v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
							v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
							v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
							v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
							v299 = base.F64_add(v269, v298)
							v301 = base.F64_add(v298, base.F64_sub(v269, v299))
							*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
							v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
							v307 = base.F64_mul(v218, v306)
							v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
							v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
							if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
								v353 = v320
								v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
								v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
								v359 = base.F64_add(base.F64_mul(v307, v355), v358)
								v360 = base.F64_sub(v359, v358)
								v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
								v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
								v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
								v370 = base.F64_mul(v369, v369)
								v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
								v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
								v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
								v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
								v386 = base.I64_reinterpret_f64(v359)
								v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
								v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
								v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
								v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
								v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
								if v353 == int32(0) {
									if v386&int64(2147483648) == int64(0) {
										v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
										v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
									} else {
										v416 = v401 + int64(4602678819172646912)
										v417 = base.F64_reinterpret_i64(v416)
										v418 = base.F64_mul(v417, v395)
										v419 = base.F64_add(v418, v417)
										if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
											v423 = float64(2.2250738585072014e-308)
											v425 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
											v432 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
											if base.F64_lt(v419, float64(0)) != 0 {
												v443 = float64(-1)
											} else {
												v443 = float64(1)
											}
											v444 = base.F64_add(v419, v443)
											v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
											if base.F64_eq(v451, float64(0)) != 0 {
												v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
											} else {
												v454 = v451
											}
											v458 = v454
										} else {
											v458 = v419
										}
										v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
									}
									v480 = v467
								} else {
									v468 = base.F64_reinterpret_i64(v401)
									v480 = base.F64_add(base.F64_mul(v468, v395), v468)
								}
							} else {
								if base.Ui32(v320) < base.Ui32(v325) {
									v336 = base.F64_add(v307, float64(1))
									if v215 != 0 {
										v338 = base.F64_neg(v336)
									} else {
										v338 = v336
									}
									v480 = v338
								} else {
									if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
										v353 = int32(0)
										v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
										v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
										v359 = base.F64_add(base.F64_mul(v307, v355), v358)
										v360 = base.F64_sub(v359, v358)
										v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
										v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
										v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
										v370 = base.F64_mul(v369, v369)
										v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
										v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
										v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
										v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
										v386 = base.I64_reinterpret_f64(v359)
										v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
										v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
										v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
										v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
										v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
										if v353 == int32(0) {
											if v386&int64(2147483648) == int64(0) {
												v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
												v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
											} else {
												v416 = v401 + int64(4602678819172646912)
												v417 = base.F64_reinterpret_i64(v416)
												v418 = base.F64_mul(v417, v395)
												v419 = base.F64_add(v418, v417)
												if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
													v423 = float64(2.2250738585072014e-308)
													v425 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
													v432 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
													if base.F64_lt(v419, float64(0)) != 0 {
														v443 = float64(-1)
													} else {
														v443 = float64(1)
													}
													v444 = base.F64_add(v419, v443)
													v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
													if base.F64_eq(v451, float64(0)) != 0 {
														v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
													} else {
														v454 = v451
													}
													v458 = v454
												} else {
													v458 = v419
												}
												v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
											}
											v480 = v467
										} else {
											v468 = base.F64_reinterpret_i64(v401)
											v480 = base.F64_add(base.F64_mul(v468, v395), v468)
										}
									} else {
										if base.I64_reinterpret_f64(v307) < int64(0) {
											v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
											mBase = m.M
											v480 = v350
										} else {
											v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
											mBase = m.M
											v480 = v352
										}
									}
								}
							}
							v483 = v480
						}
					}
				} else {
					v179 = v34
					v180 = v24
					v181 = v11
					if base.Ui32(v32) <= base.Ui32(int32(-129)) {
						if v179 == int64(4607182418800017408) {
							v483 = float64(1)
						} else {
							if base.Ui32(v30) <= base.Ui32(int32(957)) {
								if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
									v192 = l1
								} else {
									v192 = base.F64_neg(l1)
								}
								v483 = base.F64_add(v192, float64(1))
							} else {
								if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
									v202 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
									mBase = m.M
									v483 = v202
								} else {
									v205 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
									mBase = m.M
									v483 = v205
								}
							}
						}
					} else {
						if v180 != 0 {
							v213 = v179
							v215 = v181
						} else {
							v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
							v215 = v181
						}
						v216 = int64(-134217728)
						v218 = base.F64_reinterpret_i64(v33 & v216)
						v220 = v213 - int64(4604531861337669632)
						v221 = int64(52)
						v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
						v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
						v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
						v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
						v238 = v213 - v220&int64(-4503599627370496)
						v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
						v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
						v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
						v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
						v251 = base.F64_add(v247, v250)
						v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
						v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
						v256 = base.F64_add(base.F64_mul(v223, v253), v255)
						v257 = base.F64_add(v251, v256)
						v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
						v263 = base.F64_mul(v251, v262)
						v264 = base.F64_mul(v247, v262)
						v268 = base.F64_mul(v247, v264)
						v269 = base.F64_add(v257, v268)
						v273 = base.F64_mul(v251, v263)
						v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
						v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
						v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
						v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
						v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
						v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
						v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
						v299 = base.F64_add(v269, v298)
						v301 = base.F64_add(v298, base.F64_sub(v269, v299))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
						v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
						v307 = base.F64_mul(v218, v306)
						v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
						v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
						if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
							v353 = v320
							v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
							v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
							v359 = base.F64_add(base.F64_mul(v307, v355), v358)
							v360 = base.F64_sub(v359, v358)
							v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
							v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
							v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
							v370 = base.F64_mul(v369, v369)
							v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
							v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
							v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
							v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
							v386 = base.I64_reinterpret_f64(v359)
							v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
							v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
							v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
							v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
							v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
							if v353 == int32(0) {
								if v386&int64(2147483648) == int64(0) {
									v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
									v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
								} else {
									v416 = v401 + int64(4602678819172646912)
									v417 = base.F64_reinterpret_i64(v416)
									v418 = base.F64_mul(v417, v395)
									v419 = base.F64_add(v418, v417)
									if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
										v423 = float64(2.2250738585072014e-308)
										v425 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
										v432 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
										if base.F64_lt(v419, float64(0)) != 0 {
											v443 = float64(-1)
										} else {
											v443 = float64(1)
										}
										v444 = base.F64_add(v419, v443)
										v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
										if base.F64_eq(v451, float64(0)) != 0 {
											v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
										} else {
											v454 = v451
										}
										v458 = v454
									} else {
										v458 = v419
									}
									v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
								}
								v480 = v467
							} else {
								v468 = base.F64_reinterpret_i64(v401)
								v480 = base.F64_add(base.F64_mul(v468, v395), v468)
							}
						} else {
							if base.Ui32(v320) < base.Ui32(v325) {
								v336 = base.F64_add(v307, float64(1))
								if v215 != 0 {
									v338 = base.F64_neg(v336)
								} else {
									v338 = v336
								}
								v480 = v338
							} else {
								if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
									v353 = int32(0)
									v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
									v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
									v359 = base.F64_add(base.F64_mul(v307, v355), v358)
									v360 = base.F64_sub(v359, v358)
									v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
									v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
									v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
									v370 = base.F64_mul(v369, v369)
									v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
									v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
									v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
									v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
									v386 = base.I64_reinterpret_f64(v359)
									v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
									v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
									v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
									v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
									v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
									if v353 == int32(0) {
										if v386&int64(2147483648) == int64(0) {
											v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
											v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
										} else {
											v416 = v401 + int64(4602678819172646912)
											v417 = base.F64_reinterpret_i64(v416)
											v418 = base.F64_mul(v417, v395)
											v419 = base.F64_add(v418, v417)
											if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
												v423 = float64(2.2250738585072014e-308)
												v425 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
												v432 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
												if base.F64_lt(v419, float64(0)) != 0 {
													v443 = float64(-1)
												} else {
													v443 = float64(1)
												}
												v444 = base.F64_add(v419, v443)
												v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
												if base.F64_eq(v451, float64(0)) != 0 {
													v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
												} else {
													v454 = v451
												}
												v458 = v454
											} else {
												v458 = v419
											}
											v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
										}
										v480 = v467
									} else {
										v468 = base.F64_reinterpret_i64(v401)
										v480 = base.F64_add(base.F64_mul(v468, v395), v468)
									}
								} else {
									if base.I64_reinterpret_f64(v307) < int64(0) {
										v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
										mBase = m.M
										v480 = v350
									} else {
										v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
										mBase = m.M
										v480 = v352
									}
								}
							}
						}
						v483 = v480
					}
				}
			}
		}
	}
	m.G0 = v19 + int32(16)
	return v483
}
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v17 = m.Wasi_snapshot_preview1.Fd_pread(m, l0, v8+int32(8), int32(1), l3, v8+int32(4))
	mBase = m.M
	if v17 == int32(0) {
		v24 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pread[0])) = v17
		v24 = int32(-1)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	m.G0 = v8 + int32(16)
	if v24 != 0 {
		v30 = int32(-1)
	} else {
		v30 = v25
	}
	return v30
}
func F_printtup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v15 == v17 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v129 < v128 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v19 == v16 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v23 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	F_pfree(m, v23)
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v15
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v30
	if v16 <= v30 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	v36 = F_palloc0(m, v16*int32(40))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v36
	v44 = int32(0)
	goto L13
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v52 = v49 + v44*int32(40)
	if v22 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L1
L15:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	F_fmgr_info(m, v109, v52+int32(12))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L28
	}
L16:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+v15+v44*int32(100))+88))
	F_getTypeOutputInfo(m, v102, v52, v52+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L27
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v56 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+10)) = uint16(v56)
	v96 = v55 << (uint(int32(4)) % 32)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22+v44<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v52)+10)) = uint16(v64)
	v67 = v60 << (uint(int32(4)) % 32)
	switch v64 {
	case 0:
		v96 = v67
		goto L16
	case 1:
		goto L21
	default:
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67+v15+v44*int32(100))+88))
	v74 = v52 + int32(4)
	F_getTypeBinaryOutputInfo(m, v72, v74, v52+int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v107 = v74
	goto L15
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = base.I32_extend16_s(v64)
	F_errmsg(m, int32(_a_F_printtup_0), v13)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_printtup_1), int32(292), int32(_a_F_printtup_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
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
	v107 = v52
	goto L15
L28:
	;
	v115 = v44 + int32(1)
	if v115 != v16 {
		v44 = v115
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	F_slot_getsomeattrs_int(m, l0, v128)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v133 = int32(_a_F_printtup_3)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_printtup[0]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_printtup[0])) = v136
	v139 = l1 + int32(40)
	F_resetStringInfo(m, v139)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = int32(68)
	goto L34
L33:
	;
	goto L32
L34:
	;
	F_enlargeStringInfo(m, v139, int32(2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v149 = int32(8)
	v155 = v16<<(uint(v149)%32) | int32(base.Ui32(v16&int32(_a_F_printtup_4))>>(uint(v149)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v146+v147))) = uint16(v155)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v146 + int32(2)
	if int32(0) < v16 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v166 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_pq_endmessage_reuse(m, v139)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L9
	} else {
		goto L55
	}
L39:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v166))))
	if v175 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	v247 = v166 + int32(1)
	if v247 != v16 {
		v166 = v247
		goto L39
	} else {
		goto L54
	}
L42:
	;
	F_enlargeStringInfo(m, v139, int32(4))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189+v166<<(uint(int32(2))%32))))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v197 = v194 + v166*int32(40)
	v199 = v197 + int32(12)
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v197)+10)))
	if v200 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v181+v182))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v181 + int32(4)
	goto L41
L46:
	;
	v203 = F_OutputFunctionCall(m, v199, v193)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v208 = F_SendFunctionCall(m, v199, v193)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	v205 = F_strlen(m, v203)
	mBase = m.M
	F_pq_sendcountedtext(m, v139, v203, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L41
L51:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	F_enlargeStringInfo(m, v139, int32(4))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v217 = int32(2)
	v219 = int32(4)
	v220 = int32(base.Ui32(v210)>>(uint(v217)%32)) - v219
	v221 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v214+v215))) = base.I32_rotr(v220&v221, int32(8)) | base.I32_rotr(v220, int32(24))&v221
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v214 + v219
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	F_appendBinaryStringInfo(m, v139, v208+v219, int32(base.Ui32(v236)>>(uint(v217)%32))-v219)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	goto L41
L54:
	;
	goto L40
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_printtup[0])) = v134
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_MemoryContextReset(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	m.G0 = v13 + int32(16)
	return int32(1)
}
func F_printtup_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = l0 + int32(40)
	F_initStringInfo(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_printtup_startup[0]))
		v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_printtup_startup_0), int32(0), int32(_a_F_printtup_startup_1), int32(_a_F_printtup_startup_2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v16
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v19 == int32(1) {
				v22 = F_FetchPortalTargetList(m, v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
					F_SendRowDescriptionMessage(m, v7, l2, v22, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				return
			}
		}
	}
}
func F_proclock_hash(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_proclock_hash[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_get_hash_value(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6 ^ v10<<(uint(int32(4))%32)
	}
}
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_psprintf[0]))
	v14 = F_palloc(m, int32(128))
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
	*(*int32)(unsafe.Add(mBase, _c_F_psprintf[0])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v22 = F_pvsnprintf(m, v14, int32(128), l0, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v22) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v14
	v29 = v22
	goto L7
L5:
	;
	v44 = v14
	goto L6
L6:
	;
	m.G0 = v9 + int32(16)
	return v44
L7:
	;
	F_pfree(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v44 = v34
	goto L6
L9:
	;
	v34 = F_palloc(m, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_psprintf[0])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v39 = F_pvsnprintf(m, v34, v29, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(v29) <= base.Ui32(v39) {
		v28 = v34
		v29 = v39
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_pstrdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v2 = int32(0)
	v4 = F_strlen(m, l0)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pstrdup[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v2)
	v10 = v4 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v6, v10, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			base.MemoryCopy(m, v14, l0, v10)
		} else {
		}
		return v14
	}
}
func F_pull_exec_paramids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	v3 = int32(0)
	if l0 == v3 {
		v24 = v3
		return v24
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(8) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v9 != int32(1) {
				v24 = v3
				return v24
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v14 = F_bms_add_member(m, v12, v13)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					return int32(0)
				}
			}
		} else {
			v22 = F_expression_tree_walker_impl(m, l0, int32(909), l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				return v24
			}
		}
	}
}
func F_pull_up_sublinks_jointree_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(32)
	return v181
L4:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v181 = v19
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v22 - int32(63) {
	case 0:
		goto L7
	case 1:
		goto L9
	case 2:
		goto L10
	default:
		goto L8
	}
L7:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v177 = F_bms_make_singleton(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L42
	}
L9:
	;
	v80 = F_palloc(m, int32(40))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 == int32(0) {
		v62 = v4
		v63 = v4
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v66 = F_makeFromExpr(m, v63, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		v62 = v4
		v63 = v4
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = v4
	v37 = v4
	v38 = v4
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v38<<(uint(int32(2))%32))))
	v46 = F_pull_up_sublinks_jointree_recurse(m, l0, v43, v11+int32(28))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v62 = v51
	v63 = v48
	goto L11
L16:
	;
	v48 = F_lappend(m, v37, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v51 = F_bms_join(m, v36, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v54 = v38 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v54 < v55 {
		v36 = v51
		v37 = v48
		v38 = v54
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v72 = int32(0)
	v74 = F_pull_up_sublinks_qual_recurse(m, l0, v69, v11+int32(28), v62, v72, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v62
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v181 = v78
	goto L3
L22:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+32)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+24)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+16)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v80
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v96 = F_pull_up_sublinks_jointree_recurse(m, l0, v93, v11+int32(28))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v102 = F_pull_up_sublinks_jointree_recurse(m, l0, v99, v11+int32(24))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	switch v105 {
	case 0:
		goto L26
	case 1:
		goto L29
	case 2:
		goto L25
	case 3:
		goto L28
	default:
		goto L27
	}
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v154 = F_bms_join(m, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L26:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v145 = F_bms_union(m, v143, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v119 = int32(0)
	v121 = F_pull_up_sublinks_qual_recurse(m, l0, v115, v80+int32(12), v118, v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v110 = int32(0)
	v112 = F_pull_up_sublinks_qual_recurse(m, l0, v106, v80+int32(16), v109, v110, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v112
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v121
	goto L25
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v128
	F_errmsg_internal(m, int32(_a_F_pull_up_sublinks_jointree_recurse_0), v11+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_pull_up_sublinks_jointree_recurse_1), int32(613), int32(_a_F_pull_up_sublinks_jointree_recurse_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
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
	v147 = int32(0)
	v149 = F_pull_up_sublinks_qual_recurse(m, l0, v140, v11+int32(20), v145, v147, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v149
	goto L25
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v80)+36))
	if v157 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v158 = F_bms_add_member(m, v154, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v181 = v161
	goto L3
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v158
	goto L40
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v166
	F_errmsg_internal(m, int32(_a_F_pull_up_sublinks_jointree_recurse_3), v11)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pull_up_sublinks_jointree_recurse_1), int32(632), int32(_a_F_pull_up_sublinks_jointree_recurse_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v177
	v181 = l1
	goto L3
}
func F_pull_up_subqueries_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v656 int32
	_ = v656
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v766 int32
	_ = v766
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1027 int32
	_ = v1027
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	F_check_stack_depth(m)
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
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pull_up_subqueries_recurse[0]))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v26 - int32(63) {
	case 0:
		goto L13
	case 1:
		goto L9
	case 2:
		goto L12
	default:
		goto L11
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v16 + int32(80)
	return v1067
L8:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v571 = F_copyObjectImpl(m, v47)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L133
	}
L9:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v499 {
	case 0:
		goto L117
	case 1, 4, 5:
		goto L118
	case 2:
		goto L119
	case 3:
		goto L120
	default:
		goto L121
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L114
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L111
	}
L12:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v440 == int32(0) {
		v1067 = l1
		goto L7
	} else {
		goto L105
	}
L13:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31+v32<<(uint(int32(2))%32)-int32(4))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v39 != int32(1) {
		v124 = v39
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v363 != int32(3) {
		v1067 = l1
		goto L7
	} else {
		goto L86
	}
L15:
	;
	v322 = F_palloc0(m, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L81
	}
L16:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v321 = v316<<(uint(int32(2))%32) + int32(4)
	goto L15
L17:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+52))
	if v231 != 0 {
		goto L68
	} else {
		goto L69
	}
L18:
	;
	if l2|l3|base.B2i32(v124 != int32(5)) != 0 {
		goto L14
	} else {
		goto L48
	}
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v43 = F_is_simple_subquery(m, l0, v42, v38, l2)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v97 != int32(1) {
		v124 = v97
		goto L18
	} else {
		goto L37
	}
L21:
	;
	if v43 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	if l3 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+60))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v51 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v54 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v62 = v50
	goto L28
L27:
	;
	goto L26
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v70 != int32(65) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v70 != int32(63) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v75 != 0 {
		goto L20
	} else {
		goto L34
	}
L33:
	;
	goto L8
L34:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 != int32(1) {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v62 = v83
	goto L28
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v101 != int32(67) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v104 != int32(1) {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+144))
	if v107 == int32(0) {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v100)+124))
	if v110 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+128))
	if v111 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+132))
	if v112 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v100)+140))
	if v113 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)+48))
	if v114 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+20))
	v116 = F_is_simple_union_all_recurse(m, v107, v100, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v116 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v124 = v118
	goto L18
L48:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	if v136 == int32(0) {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v139 != int32(1) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v142 = F_expression_returns_set(m, v136)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v142 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v145 = F_contain_volatile_functions(m, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v145 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+52))
	if v148 == int32(0) {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v151 != int32(1) {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v38 != v155 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v158 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v162 = F_copyObjectImpl(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v215 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v215
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v147 + int32(39)
	if v204 != 0 {
		goto L16
	} else {
		goto L67
	}
L59:
	;
	if v162 == int32(0) {
		v204 = v158
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v167 <= int32(0) {
		v204 = v158
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v173 = v158
	v176 = int32(1)
	v178 = int32(0)
	goto L62
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v178<<(uint(int32(2))%32))))
	v190 = int32(0)
	v192 = F_makeTargetEntry(m, v188, base.I32_extend16_s(v176), v190, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v204 = v194
	goto L58
L64:
	;
	v194 = F_lappend(m, v173, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v196 = int32(1)
	v199 = v178 + v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v199 < v200 {
		v173 = v194
		v176 = v176 + v196
		v178 = v199
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v321 = int32(4)
	goto L15
L68:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v234 = v232
	goto L70
L69:
	;
	v234 = int32(0)
	goto L70
L70:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v229)+52))
	v237 = F_copyObjectImpl(m, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v239 = m.G0
	v240 = int32(16)
	v241 = v239 - v240
	m.G0 = v241
	*(*int32)(unsafe.Add(mBase, uint32(v241)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v241)+8)) = int32(-1)
	v251 = F_range_table_walker_impl(m, v237, int32(1050), v241+int32(8), v240)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	m.G0 = v241 + int32(16)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if base.B2i32(v237 == int32(0))|base.B2i32(v258 != int32(1)) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v229)+56))
	F_CombineRangeTables(m, v303+int32(52), v303+int32(56), v237, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L79
	}
L74:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v262 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v271 = int32(0)
	goto L76
L76:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v271<<(uint(int32(2))%32))))
	v284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+124)) = uint8(v284)
	v287 = v271 + v284
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v287 < v288 {
		v271 = v287
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L73
L78:
	;
	goto L77
L79:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v229)+144))
	F_pull_up_union_leaf_queries(m, v311, l0, v235, v229, v234)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v314 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+20)) = uint8(v314)
	v1067 = l1
	goto L7
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v322
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v331 = F_palloc0(m, int32(136))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+12)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = int32(101)
	v339 = F_makeAlias(m, int32(_a_F_pull_up_subqueries_recurse_0), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331)+8)) = v339
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v331
	v347 = F_list_make1_impl(m, int32(1), v16+int32(12))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+52)) = v347
	v1067 = l1
	goto L7
L86:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+72)))
	if v366 != 0 {
		v1067 = l1
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	if v367 == int32(0) {
		v1067 = l1
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v370 != int32(1) {
		v1067 = l1
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v367)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	if v376 != int32(7) {
		v1067 = l1
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v374)+8))
	if v379 != int32(1) {
		v1067 = l1
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v374)+12))
	if v382 != 0 {
		v1067 = l1
		goto L7
	} else {
		goto L92
	}
L92:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v388 = F_get_expr_result_type(m, v375, v16+int32(68), v16-int32(-64))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v388 != 0 {
		v1067 = l1
		goto L7
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v393 = int32(0)
	v395 = F_makeTargetEntry(m, v391, int32(1), v393, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v395
	v402 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v383 + int32(39)
	v407 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v407
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v402
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v413
	if v402 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v402)+4))
	v423 = v417<<(uint(int32(2))%32) + int32(4)
	goto L99
L98:
	;
	v423 = int32(4)
	goto L99
L99:
	;
	v424 = F_palloc0(m, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v424
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v383)+108))
	if v427 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	goto L103
L102:
	;
	goto L103
L103:
	;
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), l3)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)) = uint8(v434)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = int32(8)
	v1067 = l1
	goto L7
L105:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v443 <= int32(0) {
		v1067 = l1
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v451 = int32(0)
	goto L107
L107:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v462 = v459 + v451<<(uint(int32(2))%32)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)))
	v465 = F_pull_up_subqueries_recurse(m, l0, v463, l2, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v1067 = l1
	goto L7
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v465
	v469 = v451 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v469 < v470 {
		v451 = v469
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v476
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_1), v16)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(1255), int32(_a_F_pull_up_subqueries_recurse_3))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_4), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(2222), int32(_a_F_pull_up_subqueries_recurse_5))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v548 = F_pull_up_subqueries_recurse(m, l0, v546, l2, int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L131
	}
L118:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v538 = F_pull_up_subqueries_recurse(m, l0, v536, l1, int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L129
	}
L119:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v528 = F_pull_up_subqueries_recurse(m, l0, v526, l1, int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L127
	}
L120:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v518 = F_pull_up_subqueries_recurse(m, l0, v516, l1, int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L125
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v504
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_6), v16+int32(16))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(1249), int32(_a_F_pull_up_subqueries_recurse_3))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v518
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v523 = F_pull_up_subqueries_recurse(m, l0, v521, l1, int32(0))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v523
	v1067 = l1
	goto L7
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v528
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v533 = F_pull_up_subqueries_recurse(m, l0, v531, l1, int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v533
	v1067 = l1
	goto L7
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v538
	v541 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v543 = F_pull_up_subqueries_recurse(m, l0, v541, l1, int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v543
	v1067 = l1
	goto L7
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v548
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v553 = F_pull_up_subqueries_recurse(m, l0, v551, l2, int32(0))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v553
	v1067 = l1
	goto L7
L133:
	;
	v574 = F_palloc0(m, int32(384))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+4)) = v571
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = int32(267)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+8)) = v579
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v574)+12)) = v581
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v584 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v574)+20)) = v584
	*(*int32)(unsafe.Add(mBase, uint32(v574)+16)) = v583
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_pull_up_subqueries_recurse[1]))
	v589 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v574)+321)) = uint16(v589)
	*(*int32)(unsafe.Add(mBase, uint32(v574)+312)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v574)+280)) = v588
	*(*int64)(unsafe.Add(mBase, uint32(v574)+72)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v574)+80)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v574)+85)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v574)+116)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v574)+124)) = v584
	*(*int64)(unsafe.Add(mBase, uint32(v574)+132)) = v584
	base.MemoryFill(m, v574+int32(192), v589, int32(88))
	*(*int64)(unsafe.Add(mBase, uint32(v574)+344)) = int64(4294967295)
	F_replace_empty_jointree(m, v571)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571)+39)))
	if v615 == int32(1) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v618)+60))
	v622 = F_pull_up_sublinks_jointree_recurse(m, v574, v619, v16+int32(24))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v641)+52))
	if v642 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	if v624 != int32(65) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v622
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v622
	v632 = F_list_make1_impl(m, int32(1), v16+int32(4))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	v637 = v622
	goto L142
L142:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v638)+60)) = v637
	goto L138
L143:
	;
	v635 = F_makeFromExpr(m, v632, int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v637 = v635
	goto L142
L145:
	;
	v703 = F_expand_virtual_generated_columns(m, v574)
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L1
	} else {
		goto L156
	}
L146:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	if v645 <= int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v656 = int32(0)
	goto L148
L148:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v642)+12))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v662+v656<<(uint(int32(2))%32))))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v666)+12))
	if v667 != int32(3) {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	goto L145
L150:
	;
	v687 = v656 + int32(1)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
	if v687 < v688 {
		v656 = v687
		goto L148
	} else {
		goto L155
	}
L151:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666)+68))
	v671 = F_eval_const_expressions(m, v574, v670)
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v666)+68)) = v671
	v674 = F_inline_set_returning_function(m, v574, v666)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	if v674 == int32(0) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v678 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v666)+72)) = uint8(v678)
	*(*uint8)(unsafe.Add(mBase, uint32(v666)+40)) = uint8(v678)
	*(*int32)(unsafe.Add(mBase, uint32(v666)+36)) = v674
	*(*int32)(unsafe.Add(mBase, uint32(v666)+12)) = int32(1)
	goto L150
L155:
	;
	goto L149
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+4)) = v703
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v703)+60))
	v707 = int32(0)
	v709 = F_pull_up_subqueries_recurse(m, v574, v706, v707, v707)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v711)+60)) = v709
	v713 = F_is_simple_subquery(m, l0, v703, v38, l2)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	if v713 == int32(0) {
		v1067 = l1
		goto L7
	} else {
		goto L159
	}
L159:
	;
	if l3 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v703)+60))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+4))
	if v718 != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	goto L162
L162:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v703)+76))
	v784 = F_flatten_join_alias_vars(m, v574, v782, v783)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L1
	} else {
		goto L179
	}
L163:
	;
	if v766 == int32(0) {
		v1067 = l1
		goto L7
	} else {
		goto L178
	}
L164:
	;
	v724 = v717
	goto L168
L165:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v717)+8))
	if v719 != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v766 = int32(1)
	goto L163
L167:
	;
	v766 = v752
	goto L163
L168:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v724)))
	if v735 != int32(65) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v752 = int32(0)
	goto L167
L170:
	;
	goto L169
L171:
	;
	if v735 == int32(63) {
		v752 = int32(1)
		goto L167
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v724)+8))
	if v740 != 0 {
		goto L170
	} else {
		goto L175
	}
L174:
	;
	goto L170
L175:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v724)+4))
	if v741 == int32(0) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	if v744 != int32(1) {
		goto L170
	} else {
		goto L177
	}
L177:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v741)+12))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v747)))
	v724 = v748
	goto L168
L178:
	;
	goto L162
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v703)+76)) = v784
	v787 = int32(0)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v570)+52))
	if v789 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v789)+4))
	v791 = v790
	goto L182
L181:
	;
	v791 = v787
	goto L182
L182:
	;
	F_OffsetVarNodes(m, v703, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v574)+128))
	F_OffsetVarNodes(m, v794, v791)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_IncrementVarSublevelsUp(m, v703, int32(-1), int32(1))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v574)+128))
	F_IncrementVarSublevelsUp(m, v801, int32(-1), int32(1))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = l0
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v703)+76))
	v808 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v808
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v807
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v38
	v813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if v813 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v703)+60))
	v817 = int32(1)
	v819 = F_get_relids_in_jointree(m, v816, v817, v817)
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	v841 = v787
	v842 = v808
	v843 = v807
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v569
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v570 + int32(39)
	if v843 != 0 {
		goto L197
	} else {
		goto L198
	}
L190:
	;
	v823 = F_palloc(m, int32(8))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v570)+52))
	if v825 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v825)+4))
	v827 = v826
	goto L194
L193:
	;
	v827 = int32(0)
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823)+4)) = v827
	v833 = F_palloc0(m, v827<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v823))) = v833
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v570)+60))
	F_get_nullingrels_recurse(m, v836, int32(0), v823)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v703)+76))
	v841 = v819
	v842 = v823
	v843 = v840
	goto L189
L197:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v843)+4))
	v859 = v853<<(uint(int32(2))%32) + int32(4)
	goto L199
L198:
	;
	v859 = int32(4)
	goto L199
L199:
	;
	v860 = F_palloc0(m, v859)
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v860
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v570)+108))
	if v863 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	goto L203
L202:
	;
	goto L203
L203:
	;
	F_perform_pullup_replace_vars(m, l0, v16+int32(24), l3)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+124)))
	if v872 != int32(1) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v703)+52))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v703)+56))
	F_CombineRangeTables(m, v570+int32(52), v570+int32(56), v925, v926)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L216
	}
L206:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v703)+52))
	if v875 == int32(0) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v878 <= int32(0) {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	v889 = int32(0)
	goto L209
L209:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v875)+12))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895+v889<<(uint(int32(2))%32))))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v899)+12))
	switch v900 {
	case 0:
		goto L213
	case 1, 3, 4, 5:
		goto L212
	default:
		goto L211
	}
L210:
	;
	goto L205
L211:
	;
	v907 = v889 + int32(1)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v907 < v908 {
		v889 = v907
		goto L209
	} else {
		goto L215
	}
L212:
	;
	v904 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v899)+124)) = uint8(v904)
	goto L211
L213:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v899)+32))
	if v901 == int32(0) {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	goto L210
L216:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v570)+140))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v703)+140))
	v931 = F_list_concat(m, v929, v930)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v570)+140)) = v931
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v934)+68))
	if v935 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v574)+128))
	v1043 = F_list_concat(m, v1041, v1042)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L1
	} else {
		goto L244
	}
L219:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v703)+60))
	v941 = F_get_relids_in_jointree(m, v938, int32(1), int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v936 != 0 {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1041 = int32(0)
	goto L218
L222:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+68))
	if v944 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v941
	v946 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v946
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v569
	v953 = F_query_or_expression_tree_walker_impl(m, v570, int32(852), v16+int32(68), v946)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v955 = int32(0)
	v956 = m.G0
	v958 = v956 - int32(16)
	m.G0 = v958
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v960 == v955 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L225
L227:
	;
	m.G0 = v958 + int32(16)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1041 = v1027
	goto L218
L228:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v960)+4))
	if v963 <= int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v970 = int32(-1)
	v974 = v955
	goto L230
L230:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v960)+12))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v980+v974<<(uint(int32(2))%32))))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v984)+8))
	if v569 == v985 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L227
L232:
	;
	if v970 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v993 = v970
	goto L234
L234:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v994)+68))
	if v995 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v989 = F_bms_singleton_member(m, v941)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	v991 = v970
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v984)+8)) = v991
	v993 = v991
	goto L234
L238:
	;
	v991 = v989
	goto L237
L239:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v984)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v958)+12)) = v941
	v998 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v958)+8)) = v998
	*(*int32)(unsafe.Add(mBase, uint32(v958)+4)) = v569
	v1005 = F_query_or_expression_tree_walker_impl(m, v996, int32(852), v958+int32(4), v998)
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1008 = v974 + int32(1)
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v960)+4))
	if v1008 < v1009 {
		v970 = v993
		v974 = v1008
		goto L230
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	goto L231
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v1043
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = int32(0)
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+39)))
	v1049 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+39)))
	v1050 = v1048 | v1049
	*(*uint8)(unsafe.Add(mBase, uint32(v570)+39)) = uint8(v1050)
	v1052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570)+44)))
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+44)))
	v1054 = v1052 | v1053
	*(*uint8)(unsafe.Add(mBase, uint32(v570)+44)) = uint8(v1054)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v703)+60))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+8))
	if v1057 != 0 {
		v1067 = v1056
		goto L7
	} else {
		goto L245
	}
L245:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+4))
	if v1058 == int32(0) {
		v1067 = v1056
		goto L7
	} else {
		goto L246
	}
L246:
	;
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+4))
	if v1061 != int32(1) {
		v1067 = v1056
		goto L7
	} else {
		goto L247
	}
L247:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1058)+12))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1064)))
	v1067 = v1065
	goto L7
}
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		v14 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(12), l2, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v14 {
				v22 = v14
				v24 = F_palloc0(m, int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
					if v22 != 0 {
						v32 = F_palloc(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
							v41 = v35
							m.G0 = v9 + int32(16)
							return v41
						}
					} else {
						v34 = int32(0)
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
						v41 = v35
						m.G0 = v9 + int32(16)
						return v41
					}
				}
			} else {
				v41 = v14
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		v22 = int32(0)
		v24 = F_palloc0(m, int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
			if v22 != 0 {
				v32 = F_palloc(m, v22)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v41 = v35
					m.G0 = v9 + int32(16)
					return v41
				}
			} else {
				v34 = int32(0)
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v41 = v35
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	}
}
func F_push_into_mbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	if int32(0) < l3 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
		if v9 == int32(1) {
			F_px_debug(m, int32(_a_F_push_into_mbuf_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(-12)
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v20) < base.Ui32(v21+l3) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = v20 - v24 + (l3+int32(_a_F_push_into_mbuf_1))&int32(2147467264)
				v31 = F_repalloc(m, v24, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31 + v30
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v39 = v31 + (v37 - v35)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31 + (v41 - v35)
					v45 = v39
					if l3 != 0 {
						base.MemoryCopy(m, v45, l2, l3)
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v49 + l3
					return int32(0)
				}
			} else {
				v45 = v21
				if l3 != 0 {
					base.MemoryCopy(m, v45, l2, l3)
				} else {
				}
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v49 + l3
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_pushval_morph(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4)
	v23 = F_palloc(m, int32(64))
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
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_parsetext(m, v26, v16, l2, l3)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if int32(0) < v29 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v16 + int32(16)
	return
L5:
	;
	v36 = int32(0)
	v40 = v29
	v41 = v7
	v44 = v7
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v41 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pfree(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	if v95 <= v36 {
		v223 = v36
		v227 = v95
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v36<<(uint(int32(4))%32))+8)))
	v95 = v40
	v96 = v52
	v99 = v44
	goto L10
L12:
	;
	goto L13
L13:
	;
	v54 = v41 + int32(1)
	v56 = v36 << (uint(int32(4)) % 32)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46+v56)+8)))
	if base.Ui32(v58) <= base.Ui32(v54) {
		v95 = v40
		v96 = v58
		v99 = v44
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v62 = v54
	v71 = v44
	goto L15
L15:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v95 = v87
	v96 = v85
	v99 = v80
	goto L10
L17:
	;
	if v71 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v75, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v79 = int32(1)
	v80 = v71 + v79
	v82 = v62 + v79
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83+v56)+8)))
	if base.Ui32(v82) < base.Ui32(v85) {
		v62 = v82
		v71 = v80
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L16
L23:
	;
	if v99 != 0 {
		goto L50
	} else {
		goto L51
	}
L24:
	;
	v106 = v36
	v110 = v95
	v113 = int32(0)
	goto L25
L25:
	;
	v117 = v106 << (uint(int32(4)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v119 = v117 + v118
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+8)))
	if v96 != v120 {
		v223 = v106
		v227 = v110
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v223 = v201
	v227 = v216
	goto L23
L27:
	;
	if v110 <= v106 {
		v201 = v106
		v205 = v110
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v113 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+4)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+2)))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119))))
	F_pushValue(m, l1, v124, v125, l4, l5|int32(base.Ui32(v126&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v134+v117)+12))
	F_pfree(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v140 = v106 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v141 <= v140 {
		v201 = v140
		v205 = v141
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v147 = v144 + v140<<(uint(int32(4))%32)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+8)))
	if v96 != v148 {
		v201 = v140
		v205 = v141
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v152 = v147
	v153 = v140
	v157 = v141
	v159 = int32(1)
	goto L34
L34:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+4)))
	if v123 != v163 {
		v201 = v153
		v205 = v157
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v201 = v187
	v205 = v188
	goto L28
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+2)))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152))))
	F_pushValue(m, l1, v165, v166, l4, l5|int32(base.Ui32(v167&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v153<<(uint(int32(4))%32))+12))
	F_pfree(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pushOperator(m, l1, int32(2), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v187 = v153 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v188 <= v187 {
		v201 = v187
		v205 = v188
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v195 = v192 + v187<<(uint(int32(4))%32)
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+8)))
	if v96 == v196 {
		v152 = v195
		v153 = v187
		v157 = v188
		v159 = v159 + int32(1)
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	F_pushOperator(m, l1, int32(3), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v216 = v205
	goto L47
L47:
	;
	if v201 < v216 {
		v106 = v201
		v110 = v216
		v113 = v113 + int32(1)
		goto L25
	} else {
		goto L49
	}
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v216 = v215
	goto L47
L49:
	;
	goto L26
L50:
	;
	v233 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v233, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v238 = v227
	goto L52
L52:
	;
	if v223 < v238 {
		v36 = v223
		v40 = v238
		v41 = v96
		v44 = v99 + int32(1)
		goto L8
	} else {
		goto L54
	}
L53:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v238 = v237
	goto L52
L54:
	;
	goto L9
L55:
	;
	goto L4
L56:
	;
	goto L4
}
func F_pvsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_pg_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v10 {
			if base.Ui32(l1) <= base.Ui32(v10) {
				if base.Ui32(int32(1073741823)) <= base.Ui32(v10) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_pvsnprintf_0), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pvsnprintf_1), int32(140), int32(_a_F_pvsnprintf_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v21 = v10 + int32(1)
					m.G0 = v8 + int32(16)
					return v21
				}
			} else {
				v21 = v10
				m.G0 = v8 + int32(16)
				return v21
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
				F_errmsg_internal(m, int32(_a_F_pvsnprintf_3), v8)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pvsnprintf_1), int32(113), int32(_a_F_pvsnprintf_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
