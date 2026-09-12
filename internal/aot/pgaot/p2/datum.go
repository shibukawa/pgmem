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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
					if v10 == int32(1) {
						v33 = int32(6)
					} else {
						v33 = v30
					}
					v42 = v33
					v43 = F_palloc(m, v42)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						if v42 != 0 {
							v45 = F__emscripten_memcpy_bulkmem(m, v43, l0, v42)
							mBase = m.M
							v46 = v45
						} else {
							v46 = v43
						}
						return v46
					}
				}
			} else {
				v34 = int32(1)
				if v7&v34 != 0 {
					v42 = int32(base.Ui32(v7) >> (uint(v34) % 32))
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = int32(base.Ui32(v38) >> (uint(int32(2)) % 32))
				}
				v43 = F_palloc(m, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					if v42 != 0 {
						v45 = F__emscripten_memcpy_bulkmem(m, v43, l0, v42)
						mBase = m.M
						v46 = v45
					} else {
						v46 = v43
					}
					return v46
				}
			}
		} else {
			v49 = F_datumGetSize(m, l0, int32(0), l2)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v51 = F_palloc(m, v49)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					if v49 != 0 {
						v53 = F__emscripten_memcpy_bulkmem(m, v51, l0, v49)
						mBase = m.M
						v54 = v53
					} else {
						v54 = v51
					}
					return v54
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
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
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
	return v373
L2:
	;
	v373 = base.B2i32(l0 == l1)
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
	v373 = base.B2i32(v75 == int32(0))
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
	v308 = v234 + int32(1)
	if base.Ui32(int32(4)) <= base.Ui32(v308) {
		goto L109
	} else {
		goto L110
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L30
	} else {
		goto L103
	}
L28:
	;
	if l0&int32(3) == int32(0) {
		v201 = l0
		goto L70
	} else {
		goto L71
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
	v373 = int32(0)
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
		v373 = v174
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
	v373 = v174
	goto L1
L68:
	;
	if l1&int32(3) == int32(0) {
		v258 = l1
		goto L87
	} else {
		goto L88
	}
L69:
	;
	v234 = v226 - l0
	goto L68
L70:
	;
	v205 = v201
	goto L79
L71:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v185 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v234 = int32(0)
	goto L68
L73:
	;
	goto L74
L74:
	;
	v190 = l0
	goto L75
L75:
	;
	v194 = v190 + int32(1)
	if v194&int32(3) == int32(0) {
		v201 = v194
		goto L70
	} else {
		goto L77
	}
L76:
	;
	v226 = v194
	goto L69
L77:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	if v199 != 0 {
		v190 = v194
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v214 = int32(-2139062144)
	if (int32(16843008)-v211|v211)&v214 == v214 {
		v205 = v205 + int32(4)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v220 = v205
	goto L82
L81:
	;
	goto L80
L82:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v224 != 0 {
		v220 = v220 + int32(1)
		goto L82
	} else {
		goto L84
	}
L83:
	;
	v226 = v220
	goto L69
L84:
	;
	goto L83
L85:
	;
	if v234 == v291 {
		goto L26
	} else {
		goto L102
	}
L86:
	;
	v291 = v283 - l1
	goto L85
L87:
	;
	v262 = v258
	goto L96
L88:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v242 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v291 = int32(0)
	goto L85
L90:
	;
	goto L91
L91:
	;
	v247 = l1
	goto L92
L92:
	;
	v251 = v247 + int32(1)
	if v251&int32(3) == int32(0) {
		v258 = v251
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v283 = v251
	goto L86
L94:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	if v256 != 0 {
		v247 = v251
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	v271 = int32(-2139062144)
	if (int32(16843008)-v268|v268)&v271 == v271 {
		v262 = v262 + int32(4)
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v277 = v262
	goto L99
L98:
	;
	goto L97
L99:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v281 != 0 {
		v277 = v277 + int32(1)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v283 = v277
	goto L86
L101:
	;
	goto L100
L102:
	;
	v373 = int32(0)
	goto L1
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l3
	F_errmsg_internal(m, int32(460890), v9)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L30
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(474028), int32(323), int32(219872))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L30
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v373 = base.B2i32(v370 == int32(0))
	goto L1
L107:
	;
	v370 = int32(0)
	goto L106
L108:
	;
	v344 = v339
	v345 = v340
	v346 = v341
	goto L118
L109:
	;
	if (l0|l1)&int32(3) != 0 {
		v339 = l0
		v340 = l1
		v341 = v308
		goto L108
	} else {
		goto L112
	}
L110:
	;
	v332 = l0
	v333 = l1
	v334 = v308
	goto L111
L111:
	;
	if v334 == int32(0) {
		goto L107
	} else {
		goto L117
	}
L112:
	;
	v316 = l0
	v317 = l1
	v318 = v308
	goto L113
L113:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v321 != v322 {
		v339 = v316
		v340 = v317
		v341 = v318
		goto L108
	} else {
		goto L115
	}
L114:
	;
	v332 = v327
	v333 = v325
	v334 = v329
	goto L111
L115:
	;
	v324 = int32(4)
	v325 = v317 + v324
	v327 = v316 + v324
	v329 = v318 - v324
	if base.Ui32(int32(3)) < base.Ui32(v329) {
		v316 = v327
		v317 = v325
		v318 = v329
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v339 = v332
	v340 = v333
	v341 = v334
	goto L108
L118:
	;
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345))))
	if v349 == v350 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v370 = v349 - v350
	goto L106
L120:
	;
	v352 = int32(1)
	v357 = v346 - v352
	if v357 != 0 {
		v344 = v344 + v352
		v345 = v345 + v352
		v346 = v357
		goto L118
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	goto L119
L123:
	;
	goto L107
}
