package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DatumGetAnyArrayP(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v2 != int32(1) {
		v12 = F_pg_detoast_datum(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v12
		}
	} else {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v5&int32(254) != int32(2) {
			v12 = F_pg_detoast_datum(m, l0)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				return v12
			}
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
			return v10
		}
	}
}
func F_datumCopy(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	if l1 != 0 {
		return l0
	} else {
		if l2 == int32(-1) {
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v7 == int32(1) {
				v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v10&int32(254) == int32(2) {
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
					v16 = F_EOH_get_flat_size(m, v15)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = F_palloc(m, v16)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							F_EOH_flatten_into(m, v15, v20, v16)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								return v20
							}
						}
					}
				} else {
					v26 = int32(18)
					if v10 == v26 {
						v30 = v26
					} else {
						v30 = int32(2)
					}
					if base.Ui32((v10-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v37 = int32(6)
					} else {
						v37 = v30
					}
					v46 = v37
					v47 = F_palloc(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v46 == int32(0) {
							v61 = v47
							return v61
						} else {
							base.MemoryCopy(m, v47, l0, v46)
							return v47
						}
					}
				}
			} else {
				v38 = int32(1)
				if v7&v38 != 0 {
					v46 = int32(base.Ui32(v7) >> (uint(v38) % 32))
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v46 = int32(base.Ui32(v42) >> (uint(int32(2)) % 32))
				}
				v47 = F_palloc(m, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v46 == int32(0) {
						v61 = v47
						return v61
					} else {
						base.MemoryCopy(m, v47, l0, v46)
						return v47
					}
				}
			}
		} else {
			v54 = F_datumGetSize(m, l0, int32(0), l2)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int32(0)
			} else {
				v56 = F_palloc(m, v54)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v54 == int32(0) {
						v61 = v56
					} else {
						base.MemoryCopy(m, v56, l0, v54)
						v61 = v56
					}
					return v61
				}
			}
		}
	}
}
func F_datum_image_eq(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v261
L2:
	;
	v261 = base.B2i32(l0 == l1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	if int32(0) < l3 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if base.Ui32(int32(4)) <= base.Ui32(l3) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	goto L7
L7:
	;
	switch l3 + int32(2) {
	case 0:
		goto L28
	case 1:
		goto L29
	default:
		goto L27
	}
L8:
	;
	v261 = base.B2i32(v75 == int32(0))
	goto L1
L9:
	;
	v75 = int32(0)
	goto L8
L10:
	;
	v49 = v44
	v50 = v45
	v51 = v46
	goto L20
L11:
	;
	if (l0|l1)&int32(3) != 0 {
		v44 = l0
		v45 = l1
		v46 = l3
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v37 = l0
	v38 = l1
	v39 = l3
	goto L13
L13:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v21 = l0
	v22 = l1
	v23 = l3
	goto L15
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v26 != v27 {
		v44 = v21
		v45 = v22
		v46 = v23
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v37 = v32
	v38 = v30
	v39 = v34
	goto L13
L17:
	;
	v29 = int32(4)
	v30 = v22 + v29
	v32 = v21 + v29
	v34 = v23 - v29
	if base.Ui32(int32(3)) < base.Ui32(v34) {
		v21 = v32
		v22 = v30
		v23 = v34
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v44 = v37
	v45 = v38
	v46 = v39
	goto L10
L20:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v54 == v55 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v75 = v54 - v55
	goto L8
L22:
	;
	v57 = int32(1)
	v62 = v51 - v57
	if v62 != 0 {
		v49 = v49 + v57
		v50 = v50 + v57
		v51 = v62
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v196 = v178 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v196) {
		goto L75
	} else {
		goto L76
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L30
	} else {
		goto L69
	}
L28:
	;
	v178 = F_strlen(m, l0)
	mBase = m.M
	v179 = F_strlen(m, l1)
	mBase = m.M
	if v178 == v179 {
		goto L26
	} else {
		goto L68
	}
L29:
	;
	v80 = F_toast_raw_datum_size(m, l0)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	v84 = F_toast_raw_datum_size(m, l1)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	if v80 != v84 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v261 = int32(0)
	goto L1
L34:
	;
	goto L35
L35:
	;
	v88 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v90 = F_pg_detoast_datum_packed(m, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	v92 = int32(1)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v94&v92 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = v92
	goto L40
L39:
	;
	v97 = int32(4)
	goto L40
L40:
	;
	v98 = v88 + v97
	v99 = int32(1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v101&v99 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v104 = v99
	goto L43
L42:
	;
	v104 = int32(4)
	goto L43
L43:
	;
	v105 = v90 + v104
	v106 = int32(4)
	v107 = v80 - v106
	if base.Ui32(v106) <= base.Ui32(v107) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if l0 != v88 {
		goto L62
	} else {
		goto L63
	}
L45:
	;
	v169 = int32(0)
	goto L44
L46:
	;
	v143 = v138
	v144 = v139
	v145 = v140
	goto L56
L47:
	;
	if (v98|v105)&int32(3) != 0 {
		v138 = v98
		v139 = v105
		v140 = v107
		goto L46
	} else {
		goto L50
	}
L48:
	;
	v131 = v98
	v132 = v105
	v133 = v107
	goto L49
L49:
	;
	if v133 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L50:
	;
	v115 = v98
	v116 = v105
	v117 = v107
	goto L51
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v120 != v121 {
		v138 = v115
		v139 = v116
		v140 = v117
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v131 = v126
	v132 = v124
	v133 = v128
	goto L49
L53:
	;
	v123 = int32(4)
	v124 = v116 + v123
	v126 = v115 + v123
	v128 = v117 - v123
	if base.Ui32(int32(3)) < base.Ui32(v128) {
		v115 = v126
		v116 = v124
		v117 = v128
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v138 = v131
	v139 = v132
	v140 = v133
	goto L46
L56:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 == v149 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v169 = v148 - v149
	goto L44
L58:
	;
	v151 = int32(1)
	v156 = v145 - v151
	if v156 != 0 {
		v143 = v143 + v151
		v144 = v144 + v151
		v145 = v156
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L45
L62:
	;
	F_pfree(m, v88)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L30
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v174 = base.B2i32(v169 == int32(0))
	if l1 == v90 {
		v261 = v174
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	F_pfree(m, v90)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L30
	} else {
		goto L67
	}
L67:
	;
	v261 = v174
	goto L1
L68:
	;
	v261 = int32(0)
	goto L1
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
	F_errmsg_internal(m, int32(_a_F_datum_image_eq_0), v9)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L30
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_datum_image_eq_1), int32(323), int32(_a_F_datum_image_eq_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v261 = base.B2i32(v258 == int32(0))
	goto L1
L73:
	;
	v258 = int32(0)
	goto L72
L74:
	;
	v232 = v227
	v233 = v228
	v234 = v229
	goto L84
L75:
	;
	if (l0|l1)&int32(3) != 0 {
		v227 = l0
		v228 = l1
		v229 = v196
		goto L74
	} else {
		goto L78
	}
L76:
	;
	v220 = l0
	v221 = l1
	v222 = v196
	goto L77
L77:
	;
	if v222 == int32(0) {
		goto L73
	} else {
		goto L83
	}
L78:
	;
	v204 = l0
	v205 = l1
	v206 = v196
	goto L79
L79:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v209 != v210 {
		v227 = v204
		v228 = v205
		v229 = v206
		goto L74
	} else {
		goto L81
	}
L80:
	;
	v220 = v215
	v221 = v213
	v222 = v217
	goto L77
L81:
	;
	v212 = int32(4)
	v213 = v205 + v212
	v215 = v204 + v212
	v217 = v206 - v212
	if base.Ui32(int32(3)) < base.Ui32(v217) {
		v204 = v215
		v205 = v213
		v206 = v217
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v227 = v220
	v228 = v221
	v229 = v222
	goto L74
L84:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	if v237 == v238 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v258 = v237 - v238
	goto L72
L86:
	;
	v240 = int32(1)
	v245 = v234 - v240
	if v245 != 0 {
		v232 = v232 + v240
		v233 = v233 + v240
		v234 = v245
		goto L84
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	goto L73
}
