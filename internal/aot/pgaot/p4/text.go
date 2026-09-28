package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replace_text(m *base.Module, l0 int32) int64 {
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	v7 = m.G0
	v9 = v7 - int32(1088)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v22 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	m.G0 = v9 + int32(1088)
	return base.I64_extend_i32_u(v212)
L6:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v54 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v28 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v39 = int32(1)
	if v22&v39 != 0 {
		v51 = int32(base.Ui32(v22)>>(uint(v39)%32)) - v39
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v31 = int32(16)
	goto L12
L11:
	;
	v31 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v28-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(4)
	goto L15
L14:
	;
	v38 = v31
	goto L15
L15:
	;
	v51 = v38
	goto L6
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if base.B2i32(v51 <= int32(0))|base.B2i32(v83 <= int32(0)) != 0 {
		v212 = v12
		goto L5
	} else {
		goto L28
	}
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v60 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v71 = int32(1)
	if v54&v71 != 0 {
		v83 = int32(base.Ui32(v54)>>(uint(v71)%32)) - v71
		goto L17
	} else {
		goto L27
	}
L21:
	;
	v63 = int32(16)
	goto L23
L22:
	;
	v63 = int32(0)
	goto L23
L23:
	;
	if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = int32(4)
	goto L26
L25:
	;
	v70 = v63
	goto L26
L26:
	;
	v83 = v70
	goto L17
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v83 = int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(4)
	goto L17
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v89 = v9 + int32(16)
	F_text_position_setup(m, v12, v17, v87, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v92 = F_text_position_next(m, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v92 == int32(0) {
		v212 = v12
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v96 = int32(1)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v98&v96 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v101 = v96
	goto L34
L33:
	;
	v101 = int32(4)
	goto L34
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
	F_initStringInfo(m, v9)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v106 = v12 + v101
	v107 = v103
	goto L36
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_replace_text[0]))
	if v113 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v164 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_appendBinaryStringInfo(m, v9, v106, v107-v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v119 == int32(1) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v149 = int32(1)
	if v119&v149 != 0 {
		goto L54
	} else {
		goto L55
	}
L44:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v125 == int32(18) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v136 = int32(1)
	if v119&v136 != 0 {
		v148 = int32(base.Ui32(v119)>>(uint(v136)%32)) - v136
		goto L43
	} else {
		goto L53
	}
L47:
	;
	v128 = int32(16)
	goto L49
L48:
	;
	v128 = int32(0)
	goto L49
L49:
	;
	if base.Ui32((v125-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v135 = int32(4)
	goto L52
L51:
	;
	v135 = v128
	goto L52
L52:
	;
	v148 = v135
	goto L43
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v148 = int32(base.Ui32(v142)>>(uint(int32(2))%32)) - int32(4)
	goto L43
L54:
	;
	v153 = v149
	goto L56
L55:
	;
	v153 = int32(4)
	goto L56
L56:
	;
	F_appendBinaryStringInfo(m, v9, v20+v153, v148)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1072))
	v158 = v107 + v157
	v161 = F_text_position_next(m, v9+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v161 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v9)+1068))
	v106 = v158
	v107 = v163
	goto L36
L60:
	;
	goto L61
L61:
	;
	goto L37
L62:
	;
	F_appendBinaryStringInfo(m, v9, v158, v189+v12-v158)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v168 = int32(18)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v170 == v168 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	v181 = int32(1)
	if v164&v181 != 0 {
		v189 = int32(base.Ui32(v164) >> (uint(v181) % 32))
		goto L62
	} else {
		goto L72
	}
L66:
	;
	v173 = v168
	goto L68
L67:
	;
	v173 = int32(2)
	goto L68
L68:
	;
	if base.Ui32((v170-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v180 = int32(6)
	goto L71
L70:
	;
	v180 = v173
	goto L71
L71:
	;
	v189 = v180
	goto L62
L72:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v189 = int32(base.Ui32(v185) >> (uint(int32(2)) % 32))
	goto L62
L73:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v197 = v195 + int32(4)
	v198 = F_palloc(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198))) = v197 << (uint(int32(2)) % 32)
	if v195 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	base.MemoryCopy(m, v198+int32(4), v194, v195)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	F_pfree(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v212 = v198
	goto L5
}
func F_text_lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v48 = int32(1)
			if v17&v48 != 0 {
				v52 = v48
			} else {
				v52 = int32(4)
			}
			v54 = int32(1)
			if v16&v54 != 0 {
				v58 = v54
			} else {
				v58 = int32(4)
			}
			if v16 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v88 = v75
			} else {
				v76 = int32(1)
				if v16&v76 != 0 {
					v88 = int32(base.Ui32(v16)>>(uint(v76)%32)) - v76
				} else {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v89 = F_varstr_cmp(m, v9+v52, v46, v14+v58, v88, v47)
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return int64(0)
			} else {
				v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v91 != v9 {
					F_pfree(m, v9)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v95 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(int32(base.Ui32(v89) >> (uint(int32(31)) % 32)))
							}
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v89) >> (uint(int32(31)) % 32)))
						}
					}
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v95 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v89) >> (uint(int32(31)) % 32)))
						}
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v89) >> (uint(int32(31)) % 32)))
					}
				}
			}
		}
	}
}
func F_text_reverse(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v11 = int32(1)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v15 = v13 & v11
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = v11
	goto L5
L4:
	;
	v16 = int32(4)
	goto L5
L5:
	;
	v17 = v7 + v16
	if v13 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v46 = v44 + int32(4)
	v47 = F_palloc(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
	if v23 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v34 = int32(1)
	if v15 != 0 {
		v44 = int32(base.Ui32(v13)>>(uint(v34)%32)) - v34
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v26 = int32(16)
	goto L12
L11:
	;
	v26 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = int32(4)
	goto L15
L14:
	;
	v33 = v26
	goto L15
L15:
	;
	v44 = v33
	goto L6
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46 << (uint(int32(2)) % 32)
	v52 = v46 + v47
	v53 = v17 + v44
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_text_reverse[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56*int32(28))+uint32(_c_F_text_reverse[1])))
	goto L19
L18:
	;
	return base.I64_extend_i32_u(v47)
L19:
	;
	if v61 <= int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v44 <= int32(0) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v44 <= int32(0) {
		goto L18
	} else {
		goto L27
	}
L23:
	;
	v66 = v17
	v67 = v52
	goto L24
L24:
	;
	v71 = int32(1)
	v72 = v67 - v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v73)
	v76 = v66 + v71
	if base.Ui32(v76) < base.Ui32(v53) {
		v66 = v76
		v67 = v72
		goto L24
	} else {
		goto L26
	}
L25:
	;
	goto L18
L26:
	;
	goto L25
L27:
	;
	v80 = v17
	v81 = v52
	goto L28
L28:
	;
	v85 = F_pg_mblen_range(m, v80, v53)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L18
L30:
	;
	v87 = v81 - v85
	if v85 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	base.MemoryCopy(m, v87, v80, v85)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v89 = v80 + v85
	if base.Ui32(v89) < base.Ui32(v53) {
		v80 = v89
		v81 = v87
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
}
func F_text_smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v16 == int32(1) {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v22 == int32(18) {
					v25 = int32(16)
				} else {
					v25 = int32(0)
				}
				if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v32 = int32(4)
				} else {
					v32 = v25
				}
				v45 = v32
			} else {
				v33 = int32(1)
				if v16&v33 != 0 {
					v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v47 = int32(1)
			if v16&v47 != 0 {
				v51 = v47
			} else {
				v51 = int32(4)
			}
			v53 = int32(1)
			if v15&v53 != 0 {
				v57 = v53
			} else {
				v57 = int32(4)
			}
			if v15 == int32(1) {
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
				if v64 == int32(18) {
					v67 = int32(16)
				} else {
					v67 = int32(0)
				}
				if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v74 = int32(4)
				} else {
					v74 = v67
				}
				v87 = v74
			} else {
				v75 = int32(1)
				if v15&v75 != 0 {
					v87 = int32(base.Ui32(v15)>>(uint(v75)%32)) - v75
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v88 = F_varstr_cmp(m, v8+v51, v45, v13+v57, v87, v46)
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				if v88 < int32(0) {
					v92 = v8
				} else {
					v92 = v13
				}
				return base.I64_extend_i32_u(v92)
			}
		}
	}
}
func F_text_to_cstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v5 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
		if v9 == int32(1) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+1)))
			if v15 == int32(18) {
				v18 = int32(16)
			} else {
				v18 = int32(0)
			}
			if base.Ui32((v15-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v25 = int32(4)
			} else {
				v25 = v18
			}
			v38 = v25
		} else {
			v26 = int32(1)
			if v9&v26 != 0 {
				v38 = int32(base.Ui32(v9)>>(uint(v26)%32)) - v26
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v41 = F_palloc(m, v38+int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			if v38 != 0 {
				v43 = int32(1)
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
				if v45&v43 != 0 {
					v48 = v43
				} else {
					v48 = int32(4)
				}
				base.MemoryCopy(m, v41, v5+v48, v38)
			} else {
			}
			v52 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v38+v41))) = uint8(v52)
			if l0 != v5 {
				F_pfree(m, v5)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					return v41
				}
			} else {
				return v41
			}
		}
	}
}
