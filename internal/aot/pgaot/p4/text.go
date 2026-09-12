package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replace_text(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	v10 = m.G0
	v12 = v10 - int32(1088)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum_packed(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = F_pg_detoast_datum_packed(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v25 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v56 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v28 = int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v30&int32(254) == int32(2) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v43 = int32(1)
	if v25&v43 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v43)%32)) - v43
		goto L5
	} else {
		goto L15
	}
L9:
	;
	v39 = v28
	goto L11
L10:
	;
	v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
	goto L11
L11:
	;
	if v30 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v28
	goto L14
L13:
	;
	v42 = v39
	goto L14
L14:
	;
	v55 = v42
	goto L5
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L5
L16:
	;
	if v55 <= int32(0) {
		v226 = v15
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v59 = int32(4)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v61&int32(254) == int32(2) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v74 = int32(1)
	if v56&v74 != 0 {
		v86 = int32(base.Ui32(v56)>>(uint(v74)%32)) - v74
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v70 = v59
	goto L22
L21:
	;
	v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
	goto L22
L22:
	;
	if v61 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v73 = v59
	goto L25
L24:
	;
	v73 = v70
	goto L25
L25:
	;
	v86 = v73
	goto L16
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	m.G0 = v12 + int32(1088)
	return v226
L28:
	;
	if v86 <= int32(0) {
		v226 = v15
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_text_position_setup(m, v15, v20, v91, v12+int32(16))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v98 = F_text_position_next(m, v12+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v98 == int32(0) {
		v226 = v15
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v102 = int32(1)
	v103 = v15 + v102
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v106&v102 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v109 = v103
	goto L35
L34:
	;
	v109 = v15 + int32(4)
	goto L35
L35:
	;
	v113 = v23 + int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1068))
	F_initStringInfo(m, v12)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v117 = v109
	v118 = v114
	goto L37
L37:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v127 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v174 == int32(1) {
		goto L64
	} else {
		goto L65
	}
L39:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_appendBinaryStringInfo(m, v12, v117, v118-v117)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v135 = v133 & int32(1)
	if v135 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v136 = v113
	goto L46
L45:
	;
	v136 = v23 + int32(4)
	goto L46
L46:
	;
	if v133 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	F_appendBinaryStringInfo(m, v12, v136, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L58
	}
L48:
	;
	v139 = int32(4)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v141&int32(254) == int32(2) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v154 = int32(1)
	if v135 != 0 {
		v164 = int32(base.Ui32(v133)>>(uint(v154)%32)) - v154
		goto L47
	} else {
		goto L57
	}
L51:
	;
	v150 = v139
	goto L53
L52:
	;
	v150 = base.B2i32(v141 == int32(18)) << (uint(v139) % 32)
	goto L53
L53:
	;
	if v141 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v153 = v139
	goto L56
L55:
	;
	v153 = v150
	goto L56
L56:
	;
	v164 = v153
	goto L47
L57:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v164 = int32(base.Ui32(v158)>>(uint(int32(2))%32)) - int32(4)
	goto L47
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1072))
	v168 = v118 + v167
	v171 = F_text_position_next(m, v12+int32(16))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	if v171 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+1068))
	v117 = v168
	v118 = v173
	goto L37
L61:
	;
	goto L62
L62:
	;
	goto L38
L63:
	;
	F_appendBinaryStringInfo(m, v12, v168, v201+v15-v168)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L77
	}
L64:
	;
	v177 = int32(6)
	v179 = int32(18)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v181 == v179 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v193 = int32(1)
	if v174&v193 != 0 {
		v201 = int32(base.Ui32(v174) >> (uint(v193) % 32))
		goto L63
	} else {
		goto L76
	}
L67:
	;
	v184 = v179
	goto L69
L68:
	;
	v184 = int32(2)
	goto L69
L69:
	;
	if v181&int32(254) == int32(2) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v189 = v177
	goto L72
L71:
	;
	v189 = v184
	goto L72
L72:
	;
	if v181 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v192 = v177
	goto L75
L74:
	;
	v192 = v189
	goto L75
L75:
	;
	v201 = v192
	goto L63
L76:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v201 = int32(base.Ui32(v197) >> (uint(int32(2)) % 32))
	goto L63
L77:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v209 = v207 + int32(4)
	v210 = F_palloc(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v209 << (uint(int32(2)) % 32)
	if v207 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	F_pfree(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L83
	}
L80:
	;
	v217 = F__emscripten_memcpy_bulkmem(m, v210+int32(4), v206, v207)
	mBase = m.M
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	v226 = v210
	goto L27
}
func F_text_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = v10 + int32(1)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v25 = v23 & int32(1)
			if v25 != 0 {
				v26 = v15
			} else {
				v26 = v10 + int32(4)
			}
			if v23 == int32(1) {
				v29 = int32(4)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v31&int32(254) == int32(2) {
					v40 = v29
				} else {
					v40 = base.B2i32(v31 == int32(18)) << (uint(v29) % 32)
				}
				if v31 == int32(1) {
					v43 = v29
				} else {
					v43 = v40
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v25 != 0 {
					v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v55 = int32(1)
			v56 = v17 + v55
			if v19&v55 != 0 {
				v61 = v56
			} else {
				v61 = v17 + int32(4)
			}
			if v19 == int32(1) {
				v64 = int32(4)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
				if v66&int32(254) == int32(2) {
					v75 = v64
				} else {
					v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
				}
				if v66 == int32(1) {
					v78 = v64
				} else {
					v78 = v75
				}
				v91 = v78
			} else {
				v79 = int32(1)
				if v19&v79 != 0 {
					v91 = int32(base.Ui32(v19)>>(uint(v79)%32)) - v79
				} else {
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v92 = F_varstr_cmp(m, v26, v54, v61, v91, v20)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int32(0)
			} else {
				v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v94 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v98 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v92) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v92) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v98 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v92) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v92) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_text_reverse(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = int32(1)
	v14 = v9 + v13
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v19 = v17 & v13
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = v14
	goto L5
L4:
	;
	v20 = v9 + int32(4)
	goto L5
L5:
	;
	if v17 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v50 = v48 + int32(4)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v23 = int32(4)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v25&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v38 = int32(1)
	if v19 != 0 {
		v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v34 = v23
	goto L12
L11:
	;
	v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
	goto L12
L12:
	;
	if v25 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v37 = v23
	goto L15
L14:
	;
	v37 = v34
	goto L15
L15:
	;
	v48 = v37
	goto L6
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
	v56 = v50 + v51
	v57 = v20 + v48
	v59 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60*int32(28))+uint32(_consts[1304])))
	goto L19
L18:
	;
	return v51
L19:
	;
	if v65 <= int32(1) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui32(v57) <= base.Ui32(v20) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if base.Ui32(v57) <= base.Ui32(v20) {
		goto L18
	} else {
		goto L34
	}
L23:
	;
	v72 = v48 & int32(7)
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = v20
	v75 = v56
	v76 = int32(0)
	goto L27
L25:
	;
	v90 = v20
	v91 = v56
	goto L26
L26:
	;
	if base.Ui32(v48-int32(1)) < base.Ui32(int32(7)) {
		goto L18
	} else {
		goto L30
	}
L27:
	;
	v81 = int32(1)
	v82 = v75 - v81
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	*(*uint8)(unsafe.Add(mBase, uint32(v82))) = uint8(v83)
	v86 = v74 + v81
	v88 = v76 + v81
	if v88 != v72 {
		v74 = v86
		v75 = v82
		v76 = v88
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v90 = v86
	v91 = v82
	goto L26
L29:
	;
	goto L28
L30:
	;
	v99 = v90
	v100 = v91
	goto L31
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(1)))) = uint8(v108)
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(2)))) = uint8(v112)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(3)))) = uint8(v116)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(4)))) = uint8(v120)
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(5)))) = uint8(v124)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(6)))) = uint8(v128)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v100-int32(7)))) = uint8(v132)
	v134 = int32(8)
	v135 = v100 - v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v136)
	v139 = v99 + v134
	if v139 != v57 {
		v99 = v139
		v100 = v135
		goto L31
	} else {
		goto L33
	}
L32:
	;
	goto L18
L33:
	;
	goto L32
L34:
	;
	v142 = v20
	v143 = v56
	goto L35
L35:
	;
	v149 = F_pg_mblen_range(m, v142, v57)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L18
L37:
	;
	v151 = v143 - v149
	if v149 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v154 = v142 + v149
	if base.Ui32(v154) < base.Ui32(v57) {
		v142 = v154
		v143 = v151
		goto L35
	} else {
		goto L42
	}
L39:
	;
	v152 = F__emscripten_memcpy_bulkmem(m, v151, v142, v149)
	mBase = m.M
	goto L41
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L36
}
func F_text_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v24 = v22 & int32(1)
			if v24 != 0 {
				v25 = v14
			} else {
				v25 = v9 + int32(4)
			}
			if v22 == int32(1) {
				v28 = int32(4)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v30&int32(254) == int32(2) {
					v39 = v28
				} else {
					v39 = base.B2i32(v30 == int32(18)) << (uint(v28) % 32)
				}
				if v30 == int32(1) {
					v42 = v28
				} else {
					v42 = v39
				}
				v53 = v42
			} else {
				v43 = int32(1)
				if v24 != 0 {
					v53 = int32(base.Ui32(v22)>>(uint(v43)%32)) - v43
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = int32(1)
			v55 = v16 + v54
			if v18&v54 != 0 {
				v60 = v55
			} else {
				v60 = v16 + int32(4)
			}
			if v18 == int32(1) {
				v63 = int32(4)
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
				if v65&int32(254) == int32(2) {
					v74 = v63
				} else {
					v74 = base.B2i32(v65 == int32(18)) << (uint(v63) % 32)
				}
				if v65 == int32(1) {
					v77 = v63
				} else {
					v77 = v74
				}
				v90 = v77
			} else {
				v78 = int32(1)
				if v18&v78 != 0 {
					v90 = int32(base.Ui32(v18)>>(uint(v78)%32)) - v78
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v91 = F_varstr_cmp(m, v25, v53, v60, v90, v19)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				if v91 < int32(0) {
					v95 = v9
				} else {
					v95 = v16
				}
				return v95
			}
		}
	}
}
func F_text_to_cstring(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	v4 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		if v8 == int32(1) {
			v11 = int32(4)
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+1)))
			if v13&int32(254) == int32(2) {
				v22 = v11
			} else {
				v22 = base.B2i32(v13 == int32(18)) << (uint(v11) % 32)
			}
			if v13 == int32(1) {
				v25 = v11
			} else {
				v25 = v22
			}
			v38 = v25
		} else {
			v26 = int32(1)
			if v8&v26 != 0 {
				v38 = int32(base.Ui32(v8)>>(uint(v26)%32)) - v26
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
				v38 = int32(base.Ui32(v32)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v41 = F_palloc(m, v38+int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = int32(1)
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
			if v45&v43 != 0 {
				v48 = v43
			} else {
				v48 = int32(4)
			}
			if v38 != 0 {
				v50 = F__emscripten_memcpy_bulkmem(m, v41, v4+v48, v38)
				mBase = m.M
				v51 = v50
			} else {
				v51 = v41
			}
			v53 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v38+v51))) = uint8(v53)
			if l0 != v4 {
				F_pfree(m, v4)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					return v51
				}
			} else {
				return v51
			}
		}
	}
}
func F_text_to_stavalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[900]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v14
	v17 = *(*int64)(unsafe.Add(mBase, _consts[901]))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v17
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(v19)
	v22 = F_text_to_cstring(m, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v26)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+36)) = int64(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+68)) = uint8(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l3
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v22
		v39 = int32(3)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+46)) = uint16(v39)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v11 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l1
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v48 = m.T0[v47].(func(*base.Module, int32) int32)(m, v11+int32(28))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v22)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
				if v52 == int32(1) {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v55))) = int32(19)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
					F_ThrowErrorData(m, v58)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v89 = int32(0)
						v91 = v89
						v92 = v89
						*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
						m.G0 = v11 + int32(112)
						return v91
					}
				} else {
					v61 = F_pg_detoast_datum(m, v48)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = F_array_contains_nulls(m, v61)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							if v63 == int32(0) {
								v91 = v48
								v92 = v19
								*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
								m.G0 = v11 + int32(112)
								return v91
							} else {
								v67 = int32(0)
								v71 = F_errstart(m, int32(19), v67)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									if v71 == int32(0) {
										v91 = v67
										v92 = v67
										*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
										m.G0 = v11 + int32(112)
										return v91
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
											F_errmsg(m, int32(158085), v11)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493065), int32(738), int32(157546))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v89 = int32(0)
													v91 = v89
													v92 = v89
													*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v92)
													m.G0 = v11 + int32(112)
													return v91
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
}
