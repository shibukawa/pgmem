package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChooseConstraintName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	v14 = m.G0
	v16 = v14 - int32(192)
	m.G0 = v16
	v18 = int32(1)
	v21 = F_table_open(m, int32(2606), v18)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v166 = v155
	goto L40
L4:
	;
	v27 = v16 + int32(128)
	goto L10
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	v153 = F_pg_snprintf(m, v16+int32(128), int32(64), int32(466248), v16+int32(16))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L39
	}
L7:
	;
	v155 = int32(0)
	goto L3
L8:
	;
	v140 = F_strlen(m, v129)
	mBase = m.M
	goto L7
L10:
	;
	goto L11
L11:
	;
	v34 = int32(63)
	if (v27^l2)&int32(3) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v133)
	goto L8
L13:
	;
	v114 = v109
	v115 = v110
	v116 = v111
	goto L35
L14:
	;
	if v104 == int32(0) {
		v129 = v102
		v130 = v103
		goto L12
	} else {
		goto L34
	}
L15:
	;
	v102 = l2
	v103 = v27
	v104 = v34
	goto L14
L16:
	;
	goto L17
L17:
	;
	if l2&int32(3) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v71 == int32(0) {
		v129 = v68
		v130 = v69
		goto L12
	} else {
		goto L27
	}
L19:
	;
	v68 = l2
	v69 = v27
	v70 = v34
	v71 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v47 = l2
	v48 = v27
	v49 = v34
	goto L22
L22:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v51)
	if v51 == int32(0) {
		v109 = v47
		v110 = v48
		v111 = v49
		goto L13
	} else {
		goto L24
	}
L23:
	;
	v68 = v62
	v69 = v56
	v70 = v58
	v71 = v60
	goto L18
L24:
	;
	v55 = int32(1)
	v56 = v48 + v55
	v58 = v49 - v55
	v59 = int32(0)
	v60 = base.B2i32(v58 != v59)
	v62 = v47 + v55
	if v62&int32(3) == v59 {
		v68 = v62
		v69 = v56
		v70 = v58
		v71 = v60
		goto L18
	} else {
		goto L25
	}
L25:
	;
	if v58 != 0 {
		v47 = v62
		v48 = v56
		v49 = v58
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v74 == int32(0) {
		v102 = v68
		v103 = v69
		v104 = v70
		goto L14
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32(v70) < base.Ui32(int32(4)) {
		v102 = v68
		v103 = v69
		v104 = v70
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v80 = v68
	v81 = v69
	v82 = v70
	goto L30
L30:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v88 = int32(-2139062144)
	if (int32(16843008)-v85|v85)&v88 != v88 {
		v109 = v80
		v110 = v81
		v111 = v82
		goto L13
	} else {
		goto L32
	}
L31:
	;
	v102 = v96
	v103 = v94
	v104 = v98
	goto L14
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v85
	v93 = int32(4)
	v94 = v81 + v93
	v96 = v80 + v93
	v98 = v82 - v93
	if base.Ui32(int32(3)) < base.Ui32(v98) {
		v80 = v96
		v81 = v94
		v82 = v98
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v109 = v102
	v110 = v103
	v111 = v104
	goto L13
L35:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	*(*uint8)(unsafe.Add(mBase, uint32(v115))) = uint8(v118)
	if v118 == int32(0) {
		v129 = v114
		v130 = v115
		goto L12
	} else {
		goto L37
	}
L36:
	;
	v129 = v125
	v130 = v123
	goto L12
L37:
	;
	v122 = int32(1)
	v123 = v115 + v122
	v125 = v114 + v122
	v127 = v116 - v122
	if v127 != 0 {
		v114 = v125
		v115 = v123
		v116 = v127
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v155 = v18
	goto L3
L40:
	;
	v173 = F_makeObjectName(m, l0, l1, v16+int32(128))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	F_sequence_close(m, v21, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L71
	}
L42:
	;
	goto L41
L43:
	;
	if l4 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	F_pfree(m, v173)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L69
	}
L45:
	;
	F_ScanKeyInit(m, v16+int32(32), int32(2), int32(3), int32(62), v173)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L63
	}
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v177 <= int32(0) {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v180 = int32(0)
	if v180 < v177 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v183 = v177
	goto L50
L49:
	;
	v183 = v180
	goto L50
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v192 = int32(0)
	goto L51
L51:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v184+v192<<(uint(int32(2))%32))))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v206 == int32(0) {
		v225 = v205
		v226 = v206
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L45
L53:
	;
	if v226-v225 == int32(0) {
		goto L44
	} else {
		goto L61
	}
L54:
	;
	goto L53
L55:
	;
	if v205 != v206 {
		v225 = v205
		v226 = v206
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v210 = v202
	v211 = v173
	goto L57
L57:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	if v215 == int32(0) {
		v225 = v214
		v226 = v215
		goto L54
	} else {
		goto L59
	}
L58:
	;
	v225 = v214
	v226 = v215
	goto L54
L59:
	;
	v218 = int32(1)
	if v214 == v215 {
		v210 = v210 + v218
		v211 = v211 + v218
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v231 = v192 + int32(1)
	if v231 != v183 {
		v192 = v231
		goto L51
	} else {
		goto L62
	}
L62:
	;
	goto L52
L63:
	;
	v253 = int32(3)
	F_ScanKeyInit(m, v16+int32(80), v253, v253, int32(184), l3)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v264 = F_systable_beginscan(m, v21, int32(2664), int32(1), int32(0), int32(2), v16+int32(32))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v266 = F_systable_getnext(m, v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_systable_endscan(m, v264)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v266 == int32(0) {
		goto L42
	} else {
		goto L68
	}
L68:
	;
	goto L44
L69:
	;
	v288 = v166 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	v295 = F_pg_snprintf(m, v16+int32(128), int32(64), int32(466248), v16)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v166 = v288
	goto L40
L71:
	;
	m.G0 = v16 + int32(192)
	return v173
}
func F_ConstraintNameIsUsed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v12 = F_table_open(m, int32(2606), int32(1))
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l0 != 0 {
			v20 = int32(0)
		} else {
			v20 = l1
		}
		F_ScanKeyInit(m, v8, int32(9), int32(3), int32(184), v20)
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if l0 == int32(1) {
				v31 = l1
			} else {
				v31 = int32(0)
			}
			F_ScanKeyInit(m, v8+int32(48), int32(10), int32(3), int32(184), v31)
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v8+int32(96), int32(2), int32(3), int32(62), l2)
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v45 = F_systable_beginscan(m, v12, int32(2665), int32(1), int32(0), int32(3), v8)
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = F_systable_getnext(m, v45)
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_systable_endscan(m, v45)
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_sequence_close(m, v12, int32(1))
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(144)
									return base.B2i32(v47 != int32(0))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_get_constraint_type(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = F_SearchSysCache1(m, int32(19), l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(41055), v6)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(499632), int32(1242), int32(366630))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+22)))
			v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v28+v29)+72)))
			F_ReleaseCatCache(m, v9)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v31
			}
		}
	}
}
