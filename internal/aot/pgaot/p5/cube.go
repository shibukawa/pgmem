package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_cmp_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v52 float64
	_ = v52
	var v54 int32
	_ = v54
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v123 float64
	_ = v123
	var v125 int32
	_ = v125
	var v130 float64
	_ = v130
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v137 int32
	_ = v137
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v165 int32
	_ = v165
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 float64
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 float64
	_ = v222
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v228 float64
	_ = v228
	var v230 float64
	_ = v230
	var v233 float64
	_ = v233
	var v241 int32
	_ = v241
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 float64
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v277 float64
	_ = v277
	var v279 int32
	_ = v279
	var v280 float64
	_ = v280
	var v283 float64
	_ = v283
	var v285 float64
	_ = v285
	var v288 float64
	_ = v288
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 float64
	_ = v329
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v340 float64
	_ = v340
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v356 float64
	_ = v356
	var v358 float64
	_ = v358
	var v359 float64
	_ = v359
	var v365 int32
	_ = v365
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 float64
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v401 float64
	_ = v401
	var v403 int32
	_ = v403
	var v404 float64
	_ = v404
	var v417 float64
	_ = v417
	var v419 float64
	_ = v419
	var v420 float64
	_ = v420
	var v427 int32
	_ = v427
	var v462 int32
	_ = v462
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = int32(2147483647)
	v21 = v19 & v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = v22 & v20
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v462
L2:
	;
	v462 = int32(1)
	goto L1
L3:
	;
	v26 = v21
	goto L5
L4:
	;
	v26 = v24
	goto L5
L5:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = int32(8)
	v28 = l1 + v27
	v30 = l0 + v27
	v38 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(v24) < base.Ui32(v21) {
		goto L43
	} else {
		goto L44
	}
L9:
	;
	v50 = v38 << (uint(int32(3)) % 32)
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v30+v50)))
	v54 = base.B2i32(v19 < int32(0))
	if v19 < int32(0) {
		v62 = v52
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v97 = int32(8)
	v98 = l1 + v97
	v100 = l0 + v97
	v109 = int32(0)
	goto L26
L11:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v28+v50)))
	v66 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		v73 = v64
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v30+(v38+v19)<<(uint(int32(3))%32))))
	if base.F64_lt(v52, v59) != 0 {
		v62 = v52
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v62 = v59
	goto L11
L14:
	;
	if base.F64_gt(v62, v73) != 0 {
		goto L2
	} else {
		goto L17
	}
L15:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v28+(v38+v22)<<(uint(int32(3))%32))))
	if base.F64_lt(v64, v71) != 0 {
		v73 = v64
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v73 = v71
	goto L14
L17:
	;
	if v19 < int32(0) {
		v82 = v52
		goto L18
	} else {
		goto L19
	}
L18:
	;
	if v22 < int32(0) {
		v90 = v64
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v80 = *(*float64)(unsafe.Add(mBase, uint32(v30+(v38+v19)<<(uint(int32(3))%32))))
	if base.F64_lt(v52, v80) != 0 {
		v82 = v52
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v82 = v80
	goto L18
L21:
	;
	v92 = int32(-1)
	if base.F64_lt(v82, v90) != 0 {
		v462 = v92
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v88 = *(*float64)(unsafe.Add(mBase, uint32(v28+(v38+v22)<<(uint(int32(3))%32))))
	if base.F64_lt(v64, v88) != 0 {
		v90 = v64
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v90 = v88
	goto L21
L24:
	;
	v95 = v38 + int32(1)
	if v95 != v26 {
		v38 = v95
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	v121 = v109 << (uint(int32(3)) % 32)
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v100+v121)))
	v125 = base.B2i32(v19 < int32(0))
	if v19 < int32(0) {
		v133 = v123
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L8
L28:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v98+v121)))
	v137 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		v144 = v135
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v100+(v109+v19)<<(uint(int32(3))%32))))
	if base.F64_gt(v123, v130) != 0 {
		v133 = v123
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v133 = v130
	goto L28
L31:
	;
	if base.F64_gt(v133, v144) != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v98+(v109+v22)<<(uint(int32(3))%32))))
	if base.F64_gt(v135, v142) != 0 {
		v144 = v135
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v144 = v142
	goto L31
L34:
	;
	if v19 < int32(0) {
		v153 = v123
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if v22 < int32(0) {
		v161 = v135
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v100+(v109+v19)<<(uint(int32(3))%32))))
	if base.F64_gt(v123, v151) != 0 {
		v153 = v123
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v153 = v151
	goto L35
L38:
	;
	if base.F64_lt(v153, v161) != 0 {
		v462 = v92
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v98+(v109+v22)<<(uint(int32(3))%32))))
	if base.F64_gt(v135, v159) != 0 {
		v161 = v135
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v161 = v159
	goto L38
L41:
	;
	v165 = v109 + int32(1)
	if v165 != v26 {
		v109 = v165
		goto L26
	} else {
		goto L42
	}
L42:
	;
	goto L27
L43:
	;
	v187 = l0 + int32(8)
	v197 = v26
	goto L46
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(v24) <= base.Ui32(v21) {
		goto L82
	} else {
		goto L83
	}
L46:
	;
	v210 = v187 + v197<<(uint(int32(3))%32)
	v211 = *(*float64)(unsafe.Add(mBase, uint32(v210)))
	if base.B2i32(v19 < int32(0)) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v257 = v26
	goto L64
L48:
	;
	if base.F64_lt(v233, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L49:
	;
	v215 = int32(3)
	v217 = v187 + (v197+v19)<<(uint(v215)%32)
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v187+(v197+v21)<<(uint(v215)%32))))
	if base.F64_lt(v211, v222) != 0 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	if base.F64_gt(v211, float64(0)) != 0 {
		goto L2
	} else {
		goto L59
	}
L52:
	;
	v224 = v210
	goto L54
L53:
	;
	v224 = v217
	goto L54
L54:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
	if base.F64_gt(v225, float64(0)) != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v217)))
	if base.F64_lt(v211, v228) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v230 = v211
	goto L58
L57:
	;
	v230 = v228
	goto L58
L58:
	;
	v233 = v230
	goto L48
L59:
	;
	v233 = v211
	goto L48
L60:
	;
	return int32(-1)
L61:
	;
	goto L62
L62:
	;
	v241 = v197 + int32(1)
	if v241 != v21 {
		v197 = v241
		goto L46
	} else {
		goto L63
	}
L63:
	;
	goto L47
L64:
	;
	v265 = v187 + v257<<(uint(int32(3))%32)
	v266 = *(*float64)(unsafe.Add(mBase, uint32(v265)))
	if base.B2i32(v19 < int32(0)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	return int32(-1)
L66:
	;
	if base.F64_lt(v288, float64(0)) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L67:
	;
	v270 = int32(3)
	v272 = v187 + (v19+v257)<<(uint(v270)%32)
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v187+(v257+v21)<<(uint(v270)%32))))
	if base.F64_gt(v266, v277) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	if base.F64_gt(v266, float64(0)) != 0 {
		goto L2
	} else {
		goto L77
	}
L70:
	;
	v279 = v265
	goto L72
L71:
	;
	v279 = v272
	goto L72
L72:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v279)))
	if base.F64_gt(v280, float64(0)) != 0 {
		goto L2
	} else {
		goto L73
	}
L73:
	;
	v283 = *(*float64)(unsafe.Add(mBase, uint32(v272)))
	if base.F64_gt(v266, v283) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v285 = v266
	goto L76
L75:
	;
	v285 = v283
	goto L76
L76:
	;
	v288 = v285
	goto L66
L77:
	;
	v288 = v266
	goto L66
L78:
	;
	v295 = int32(1)
	v297 = v257 + v295
	if v297 == v21 {
		v462 = v295
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	goto L65
L81:
	;
	v257 = v297
	goto L64
L82:
	;
	return int32(0)
L83:
	;
	goto L84
L84:
	;
	v305 = l1 + int32(8)
	v321 = v21
	goto L85
L85:
	;
	v328 = v305 + v321<<(uint(int32(3))%32)
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
	if base.B2i32(v22 < int32(0)) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v381 = v26
	goto L102
L87:
	;
	if base.F64_lt(v359, float64(0)) != 0 {
		goto L2
	} else {
		goto L100
	}
L88:
	;
	v356 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.F64_lt(v329, v356) != 0 {
		goto L97
	} else {
		goto L98
	}
L89:
	;
	v333 = int32(3)
	v335 = v305 + (v22+v321)<<(uint(v333)%32)
	v340 = *(*float64)(unsafe.Add(mBase, uint32(v305+(v321+v24)<<(uint(v333)%32))))
	if base.F64_lt(v329, v340) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	if base.F64_gt(v329, float64(0)) == int32(0) {
		v359 = v329
		goto L87
	} else {
		goto L96
	}
L92:
	;
	v342 = v328
	goto L94
L93:
	;
	v342 = v335
	goto L94
L94:
	;
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_gt(v343, float64(0)) == int32(0) {
		goto L88
	} else {
		goto L95
	}
L95:
	;
	return int32(-1)
L96:
	;
	return int32(-1)
L97:
	;
	v358 = v329
	goto L99
L98:
	;
	v358 = v356
	goto L99
L99:
	;
	v359 = v358
	goto L87
L100:
	;
	v365 = v321 + int32(1)
	if v365 != v24 {
		v321 = v365
		goto L85
	} else {
		goto L101
	}
L101:
	;
	goto L86
L102:
	;
	v389 = v305 + v381<<(uint(int32(3))%32)
	v390 = *(*float64)(unsafe.Add(mBase, uint32(v389)))
	if base.B2i32(v22 < int32(0)) == int32(0) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v462 = int32(-1)
	goto L1
L104:
	;
	if base.F64_lt(v420, float64(0)) != 0 {
		goto L2
	} else {
		goto L117
	}
L105:
	;
	v417 = *(*float64)(unsafe.Add(mBase, uint32(v396)))
	if base.F64_gt(v390, v417) != 0 {
		goto L114
	} else {
		goto L115
	}
L106:
	;
	v394 = int32(3)
	v396 = v305 + (v22+v381)<<(uint(v394)%32)
	v401 = *(*float64)(unsafe.Add(mBase, uint32(v305+(v381+v24)<<(uint(v394)%32))))
	if base.F64_gt(v390, v401) != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	goto L108
L108:
	;
	if base.F64_gt(v390, float64(0)) == int32(0) {
		v420 = v390
		goto L104
	} else {
		goto L113
	}
L109:
	;
	v403 = v389
	goto L111
L110:
	;
	v403 = v396
	goto L111
L111:
	;
	v404 = *(*float64)(unsafe.Add(mBase, uint32(v403)))
	if base.F64_gt(v404, float64(0)) == int32(0) {
		goto L105
	} else {
		goto L112
	}
L112:
	;
	return int32(-1)
L113:
	;
	return int32(-1)
L114:
	;
	v419 = v390
	goto L116
L115:
	;
	v419 = v417
	goto L116
L116:
	;
	v420 = v419
	goto L104
L117:
	;
	v427 = v381 + int32(1)
	if v24 != v427 {
		v381 = v427
		goto L102
	} else {
		goto L118
	}
L118:
	;
	goto L103
}
func F_cube_coord(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
					F_errmsg(m, int32(172345), v7)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(499390), int32(1615), int32(420268))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			if base.Ui32(v17<<(uint(int32(1))%32)) < base.Ui32(v14) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(172345), v7)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499390), int32(1615), int32(420268))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
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
				v22 = v14 - int32(1)
				if v17 < int32(0) {
					v27 = base.I32_rem_u_s(v22, v17&int32(2147483647))
					v28 = v27
				} else {
					v28 = v22
				}
				v32 = *(*float64)(unsafe.Add(mBase, uint32(v10+v28<<(uint(int32(3))%32))+8))
				v33 = F_Float8GetDatum(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v33
				}
			}
		}
	}
}
func F_cube_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 float64
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v103 float64
	_ = v103
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v121 float64
	_ = v121
	var v123 float64
	_ = v123
	var v125 float64
	_ = v125
	var v127 float64
	_ = v127
	var v129 int32
	_ = v129
	var v136 float64
	_ = v136
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 float64
	_ = v162
	var v167 int32
	_ = v167
	var v179 float64
	_ = v179
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v198 float64
	_ = v198
	var v200 float64
	_ = v200
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v215 int32
	_ = v215
	var v222 float64
	_ = v222
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v29 = int32(2147483647)
			v30 = v28 & v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v33 = v31 & v29
			v34 = base.B2i32(base.Ui32(v30) < base.Ui32(v33))
			if base.Ui32(v30) < base.Ui32(v33) {
				v35 = v26
			} else {
				v35 = v21
			}
			if base.Ui32(v30) < base.Ui32(v33) {
				v36 = v21
			} else {
				v36 = v26
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v39 = v37 & int32(2147483647)
			if v39 == int32(0) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v136 = v2
				v140 = v42
			} else {
				v43 = int32(8)
				v44 = v36 + v43
				v46 = v35 + v43
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v54 = v2
				v56 = int32(0)
				for {
					v69 = v56 << (uint(int32(3)) % 32)
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v46+v69)))
					v73 = *(*float64)(unsafe.Add(mBase, uint32(v69+v44)))
					if int32(0) <= v47 {
						v80 = *(*float64)(unsafe.Add(mBase, uint32(v46+(v56+v47)<<(uint(int32(3))%32))))
						v81 = v80
					} else {
						v81 = v71
					}
					if int32(0) <= v37 {
						v88 = *(*float64)(unsafe.Add(mBase, uint32(v44+(v56+v37)<<(uint(int32(3))%32))))
						v89 = v88
					} else {
						v89 = v73
					}
					if base.F64_le(v81, v89) == int32(0) {
						v107 = float64(0)
						if base.F64_gt(v81, v89) == int32(0) {
							v125 = v107
						} else {
							if base.F64_lt(v73, v71) == int32(0) {
								v125 = v107
							} else {
								if base.F64_gt(v81, v73) == int32(0) {
									v125 = v107
								} else {
									if base.F64_lt(v89, v71) == int32(0) {
										v125 = v107
									} else {
										if base.F64_gt(v81, v71) != 0 {
											v121 = v71
										} else {
											v121 = v81
										}
										if base.F64_lt(v89, v73) != 0 {
											v123 = v73
										} else {
											v123 = v89
										}
										v125 = base.F64_sub(v121, v123)
									}
								}
							}
						}
					} else {
						if base.F64_ge(v73, v71) == int32(0) {
							v107 = float64(0)
							if base.F64_gt(v81, v89) == int32(0) {
								v125 = v107
							} else {
								if base.F64_lt(v73, v71) == int32(0) {
									v125 = v107
								} else {
									if base.F64_gt(v81, v73) == int32(0) {
										v125 = v107
									} else {
										if base.F64_lt(v89, v71) == int32(0) {
											v125 = v107
										} else {
											if base.F64_gt(v81, v71) != 0 {
												v121 = v71
											} else {
												v121 = v81
											}
											if base.F64_lt(v89, v73) != 0 {
												v123 = v73
											} else {
												v123 = v89
											}
											v125 = base.F64_sub(v121, v123)
										}
									}
								}
							}
						} else {
							if base.F64_le(v81, v73) == int32(0) {
								v107 = float64(0)
								if base.F64_gt(v81, v89) == int32(0) {
									v125 = v107
								} else {
									if base.F64_lt(v73, v71) == int32(0) {
										v125 = v107
									} else {
										if base.F64_gt(v81, v73) == int32(0) {
											v125 = v107
										} else {
											if base.F64_lt(v89, v71) == int32(0) {
												v125 = v107
											} else {
												if base.F64_gt(v81, v71) != 0 {
													v121 = v71
												} else {
													v121 = v81
												}
												if base.F64_lt(v89, v73) != 0 {
													v123 = v73
												} else {
													v123 = v89
												}
												v125 = base.F64_sub(v121, v123)
											}
										}
									}
								}
							} else {
								if base.F64_ge(v89, v71) == int32(0) {
									v107 = float64(0)
									if base.F64_gt(v81, v89) == int32(0) {
										v125 = v107
									} else {
										if base.F64_lt(v73, v71) == int32(0) {
											v125 = v107
										} else {
											if base.F64_gt(v81, v73) == int32(0) {
												v125 = v107
											} else {
												if base.F64_lt(v89, v71) == int32(0) {
													v125 = v107
												} else {
													if base.F64_gt(v81, v71) != 0 {
														v121 = v71
													} else {
														v121 = v81
													}
													if base.F64_lt(v89, v73) != 0 {
														v123 = v73
													} else {
														v123 = v89
													}
													v125 = base.F64_sub(v121, v123)
												}
											}
										}
									}
								} else {
									if base.F64_gt(v89, v73) != 0 {
										v103 = v73
									} else {
										v103 = v89
									}
									if base.F64_lt(v81, v71) != 0 {
										v105 = v71
									} else {
										v105 = v81
									}
									v125 = base.F64_sub(v103, v105)
								}
							}
						}
					}
					v127 = base.F64_add(base.F64_mul(v125, v125), v54)
					v129 = v56 + int32(1)
					if v129 != v39 {
						v54 = v127
						v56 = v129
						continue
					} else {
						break
					}
					break
				}
				v136 = v127
				v140 = v47
			}
			v151 = v140 & int32(2147483647)
			if base.Ui32(v39) < base.Ui32(v151) {
				v154 = v35 + int32(8)
				v162 = v136
				v167 = v39
				for {
					v179 = *(*float64)(unsafe.Add(mBase, uint32(v154+v167<<(uint(int32(3))%32))))
					if base.B2i32(v140 < int32(0)) == int32(0) {
						v186 = *(*float64)(unsafe.Add(mBase, uint32(v154+(v151+v167)<<(uint(int32(3))%32))))
						v187 = v186
					} else {
						v187 = v179
					}
					if base.F64_le(v179, float64(0)) == int32(0) {
						v200 = float64(0)
						if base.F64_gt(v179, v200) == int32(0) {
							v211 = v200
						} else {
							if base.F64_gt(v187, float64(0)) == int32(0) {
								v211 = v200
							} else {
								if base.F64_gt(v187, v179) != 0 {
									v210 = v179
								} else {
									v210 = v187
								}
								v211 = v210
							}
						}
					} else {
						if base.F64_le(v187, float64(0)) == int32(0) {
							v200 = float64(0)
							if base.F64_gt(v179, v200) == int32(0) {
								v211 = v200
							} else {
								if base.F64_gt(v187, float64(0)) == int32(0) {
									v211 = v200
								} else {
									if base.F64_gt(v187, v179) != 0 {
										v210 = v179
									} else {
										v210 = v187
									}
									v211 = v210
								}
							}
						} else {
							if base.F64_lt(v187, v179) != 0 {
								v198 = v179
							} else {
								v198 = v187
							}
							v211 = base.F64_sub(float64(0), v198)
						}
					}
					v213 = base.F64_add(base.F64_mul(v211, v211), v162)
					v215 = v167 + int32(1)
					if v215 != v151 {
						v162 = v213
						v167 = v215
						continue
					} else {
						break
					}
					break
				}
				v222 = v213
			} else {
				v222 = v136
			}
			v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v236 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v240 = m.ExcPending
					if v240 != 0 {
						return int32(0)
					} else {
						v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 != v241 {
							F_pfree(m, v26)
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									return v251
								}
							}
						} else {
							v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								return v251
							}
						}
					}
				} else {
					v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 != v241 {
						F_pfree(m, v26)
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return int32(0)
						} else {
							v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								return v251
							}
						}
					} else {
						v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							return v251
						}
					}
				}
			} else {
				if v236 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v245 = m.ExcPending
					if v245 != 0 {
						return int32(0)
					} else {
						v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 == v246 {
							v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								return v251
							}
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									return v251
								}
							}
						}
					}
				} else {
					v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 == v246 {
						v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
						mBase = m.M
						v252 = m.ExcPending
						if v252 != 0 {
							return int32(0)
						} else {
							return v251
						}
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return int32(0)
						} else {
							v251 = F_Float8GetDatum(m, base.F64_sqrt(v222))
							mBase = m.M
							v252 = m.ExcPending
							if v252 != 0 {
								return int32(0)
							} else {
								return v251
							}
						}
					}
				}
			}
		}
	}
}
func F_cube_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_cube_union_v0(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_cube_ur_coord(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v35 float64
	_ = v35
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v8 = float64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v14 <= int32(0) {
			v48 = v8
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v19 = v17 & int32(2147483647)
			if base.Ui32(v19) < base.Ui32(v14) {
				v48 = v8
			} else {
				v22 = v10 + int32(8)
				v24 = v14 - int32(1)
				v27 = v22 + v24<<(uint(int32(3))%32)
				if v17 < int32(0) {
					v41 = v27
				} else {
					v30 = *(*float64)(unsafe.Add(mBase, uint32(v27)))
					v35 = *(*float64)(unsafe.Add(mBase, uint32(v22+(v24+v19)<<(uint(int32(3))%32))))
					if base.F64_gt(v30, v35) != 0 {
						v41 = v27
					} else {
						v41 = v22 + (v17+v24)<<(uint(int32(3))%32)
					}
				}
				v42 = *(*float64)(unsafe.Add(mBase, uint32(v41)))
				v48 = v42
			}
		}
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v49 != v10 {
			F_pfree(m, v10)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = F_Float8GetDatum(m, v48)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					return v53
				}
			}
		} else {
			v53 = F_Float8GetDatum(m, v48)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v53
			}
		}
	}
}
func F_cube_yylex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v580 int32
	_ = v580
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
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
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1072 int32
	_ = v1072
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1121 int32
	_ = v1121
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1139 int32
	_ = v1139
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = l0
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v70 = l1
	v74 = v68
	v81 = int32(0)
	goto L22
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v68 = v16
	goto L1
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v19 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = int32(1)
	goto L7
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v24 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[265]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v28
	goto L10
L9:
	;
	goto L10
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v34
	goto L13
L12:
	;
	goto L13
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v36 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v60
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v63
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v65)
	v68 = v60
	goto L1
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36+v37<<(uint(int32(2))%32))))
	if v41 != 0 {
		v57 = v41
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_cube_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v49 = F_cube_yy_create_buffer(m, v47, int32(16384), l1)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v51+v52<<(uint(int32(2))%32)))) = v49
	v57 = v49
	goto L14
L22:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	v86 = v70
	v87 = v84
	v89 = v74
	v90 = v74
	v97 = v81
	goto L24
L24:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_consts[1461]))))
	if int64(1)<<(uint(base.I64_extend_i32_u(v87))%64)&int64(101125980167) == int64(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v87
	goto L28
L27:
	;
	goto L28
L28:
	;
	v111 = int32(1)
	v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87<<(uint(v111)%32))+uint32(_consts[1462]))))
	v116 = v115 + v101
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v116<<(uint(v111)%32))+uint32(_consts[1463]))))
	if v121 != v87 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v125 = v87
	goto L32
L30:
	;
	v154 = v116
	goto L31
L31:
	;
	v167 = int32(1)
	v173 = int32(*(*int16)(unsafe.Add(mBase, uint32(v154<<(uint(v167)%32))+uint32(_consts[1464]))))
	if v173 != int32(36) {
		v87 = v173
		v89 = v89 + v167
		goto L24
	} else {
		goto L35
	}
L32:
	;
	v136 = int32(1)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v125<<(uint(v136)%32))+uint32(_consts[1465]))))
	v141 = base.I32_extend16_s(v140)
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v141<<(uint(v136)%32))+uint32(_consts[1462]))))
	v147 = v146 + v101
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147<<(uint(v136)%32))+uint32(_consts[1463]))))
	if v140 != v152 {
		v125 = v141
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v154 = v147
	goto L31
L34:
	;
	goto L33
L35:
	;
	v184 = v90
	goto L36
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v86)+64))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	v193 = v189
	v197 = v190
	v199 = v184
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v197 - v199
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)) = uint8(v207)
	v209 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v197
	v216 = int32(*(*int16)(unsafe.Add(mBase, uint32(v193<<(uint(int32(1))%32))+uint32(_consts[1466]))))
	v217 = v216
	v222 = v197
	goto L40
L40:
	;
	if v217 != int32(12) {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1188
	*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = int32(0)
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	v1203 = base.I32_div_s(v1199-int32(1), int32(2))
	v217 = v1203 + int32(13)
	v222 = v1188
	goto L40
L43:
	;
	F_yy_fatal_error_6(m, int32(31348))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L19
	} else {
		goto L244
	}
L44:
	;
	F_yy_fatal_error_6(m, int32(683508))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L19
	} else {
		goto L243
	}
L45:
	;
	F_yy_fatal_error_6(m, int32(453961))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L19
	} else {
		goto L242
	}
L46:
	;
	F_yy_fatal_error_6(m, int32(449724))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L19
	} else {
		goto L241
	}
L47:
	;
	return v1167
L48:
	;
	v1167 = int32(258)
	goto L47
L49:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1163))) = v1164
	goto L48
L50:
	;
	F_yy_fatal_error_6(m, int32(423886))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L19
	} else {
		goto L240
	}
L51:
	;
	switch v217 {
	case 0:
		goto L63
	case 1:
		goto L49
	case 2:
		goto L62
	case 3:
		goto L61
	case 4:
		goto L60
	case 5:
		goto L59
	case 6:
		goto L58
	case 7:
		goto L57
	case 8:
		goto L56
	case 9:
		v70 = v86
		v74 = v222
		v81 = v97
		goto L22
	case 10:
		goto L55
	case 11:
		goto L54
	default:
		goto L50
	case 13:
		v1167 = v97
		goto L47
	}
L52:
	;
	goto L53
L53:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v272)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v274+v275<<(uint(int32(2))%32))))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+44))
	if v280 != 0 {
		goto L66
	} else {
		goto L67
	}
L54:
	;
	F_yy_fatal_error_6(m, int32(453313))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L19
	} else {
		goto L64
	}
L55:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v266 = int32(*(*int8)(unsafe.Add(mBase, uint32(v265))))
	return v266
L56:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = int32(669711)
	return int32(263)
L57:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = int32(685142)
	return int32(260)
L58:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v250))) = int32(686625)
	return int32(259)
L59:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v245))) = int32(685142)
	return int32(262)
L60:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = int32(686625)
	return int32(261)
L61:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v238
	goto L48
L62:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v234))) = v235
	goto L48
L63:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v232)
	v184 = v199
	goto L36
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v294 = v293 + v291
	if base.Ui32(v292) <= base.Ui32(v294) {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v290 = v280
	v291 = v281
	goto L65
L67:
	;
	goto L68
L68:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v285 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v279)+44)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = v284
	v290 = v285
	v291 = v282
	goto L65
L69:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v299 = v271 ^ int32(-1) + v197
	v300 = v296 + v299
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v299 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if base.Ui32(v294+int32(1)) < base.Ui32(v292) {
		goto L46
	} else {
		goto L101
	}
L72:
	;
	v305 = v302
	v307 = v296
	goto L75
L73:
	;
	v397 = v302
	goto L74
L74:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v397))%64)&int64(101125980167) == int64(0) {
		goto L90
	} else {
		goto L91
	}
L75:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	if v316 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v397 = v391
	goto L74
L77:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+uint32(_consts[1461]))))
	v321 = v319
	goto L79
L78:
	;
	v321 = int32(1)
	goto L79
L79:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v305))%64)&int64(101125980167) == int64(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v307
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v305
	goto L82
L81:
	;
	goto L82
L82:
	;
	v331 = int32(1)
	v335 = int32(*(*int16)(unsafe.Add(mBase, uint32(v305<<(uint(v331)%32))+uint32(_consts[1462]))))
	v336 = v321 + v335
	v341 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336<<(uint(v331)%32))+uint32(_consts[1463]))))
	if v341 != v305 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v345 = v305
	goto L86
L84:
	;
	v374 = v336
	goto L85
L85:
	;
	v387 = int32(1)
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v374<<(uint(v387)%32))+uint32(_consts[1464]))))
	v393 = v307 + v387
	if v393 != v300 {
		v305 = v391
		v307 = v393
		goto L75
	} else {
		goto L89
	}
L86:
	;
	v356 = int32(1)
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v345<<(uint(v356)%32))+uint32(_consts[1465]))))
	v361 = base.I32_extend16_s(v360)
	v366 = int32(*(*int16)(unsafe.Add(mBase, uint32(v361<<(uint(v356)%32))+uint32(_consts[1462]))))
	v367 = v321 + v366
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v367<<(uint(v356)%32))+uint32(_consts[1463]))))
	if v360 != v372 {
		v345 = v361
		goto L86
	} else {
		goto L88
	}
L87:
	;
	v374 = v367
	goto L85
L88:
	;
	goto L87
L89:
	;
	goto L76
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v300
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v397
	goto L92
L91:
	;
	goto L92
L92:
	;
	v417 = int32(1)
	v421 = int32(*(*int16)(unsafe.Add(mBase, uint32(v397<<(uint(v417)%32))+uint32(_consts[1462]))))
	v423 = v421 + v417
	v428 = int32(*(*int16)(unsafe.Add(mBase, uint32(v423<<(uint(v417)%32))+uint32(_consts[1463]))))
	if v428 != v397 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v432 = v397
	goto L96
L94:
	;
	v462 = v423
	goto L95
L95:
	;
	if v462 == int32(0) {
		v184 = v296
		goto L36
	} else {
		goto L99
	}
L96:
	;
	v443 = int32(1)
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432<<(uint(v443)%32))+uint32(_consts[1465]))))
	v448 = base.I32_extend16_s(v447)
	v453 = int32(*(*int16)(unsafe.Add(mBase, uint32(v448<<(uint(v443)%32))+uint32(_consts[1462]))))
	v455 = v453 + v443
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v455<<(uint(v443)%32))+uint32(_consts[1463]))))
	if v447 != v460 {
		v432 = v448
		goto L96
	} else {
		goto L98
	}
L97:
	;
	v462 = v455
	goto L95
L98:
	;
	goto L97
L99:
	;
	v481 = int32(*(*int16)(unsafe.Add(mBase, uint32(v462<<(uint(int32(1))%32))+uint32(_consts[1464]))))
	if v481 == int32(36) {
		v184 = v296
		goto L36
	} else {
		goto L100
	}
L100:
	;
	v485 = v300 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v485
	v87 = v481
	v89 = v485
	v90 = v296
	goto L24
L101:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v279)+40))
	if v491 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v1064 = v1051 + v1055
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v1064
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if base.Ui32(v1064) <= base.Ui32(v1056) {
		v193 = v1066
		v197 = v1064
		v199 = v1056
		goto L38
	} else {
		goto L224
	}
L104:
	;
	if v292-v490 != int32(1) {
		v1051 = v293
		v1055 = v291
		v1056 = v490
		goto L103
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v499 = v490 ^ int32(-1) + v292
	if v499 != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v1188 = v490
	goto L42
L108:
	;
	v893 = v879 + v499
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v881)+12))
	if base.Ui32(v893) <= base.Ui32(v894) {
		goto L195
	} else {
		goto L196
	}
L109:
	;
	if v499 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L110:
	;
	v500 = int32(7)
	v501 = v499 & v500
	if base.Ui32(v292-v490-int32(2)) < base.Ui32(v500) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v604 = v279
	v615 = v290
	goto L112
L112:
	;
	if v615 == int32(2) {
		goto L126
	} else {
		goto L127
	}
L113:
	;
	if v501 != 0 {
		goto L120
	} else {
		goto L121
	}
L114:
	;
	v546 = v293
	v548 = v490
	goto L113
L115:
	;
	goto L116
L116:
	;
	v510 = v293
	v512 = v490
	v513 = int32(0)
	goto L117
L117:
	;
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	*(*uint8)(unsafe.Add(mBase, uint32(v510))) = uint8(v523)
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+1)) = uint8(v525)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+2)) = uint8(v527)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+3)) = uint8(v529)
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+4)) = uint8(v531)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+5)) = uint8(v533)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+6)) = uint8(v535)
	v537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v510)+7)) = uint8(v537)
	v539 = int32(8)
	v540 = v510 + v539
	v542 = v512 + v539
	v544 = v513 + v539
	if v544 != v499&int32(-8) {
		v510 = v540
		v512 = v542
		v513 = v544
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v546 = v540
	v548 = v542
	goto L113
L119:
	;
	goto L118
L120:
	;
	v560 = v546
	v562 = v548
	v563 = int32(0)
	goto L123
L121:
	;
	goto L122
L122:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v595+v596<<(uint(int32(2))%32))))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v600)+44))
	v604 = v600
	v615 = v601
	goto L112
L123:
	;
	v573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v562))))
	*(*uint8)(unsafe.Add(mBase, uint32(v560))) = uint8(v573)
	v575 = int32(1)
	v580 = v563 + v575
	if v580 != v501 {
		v560 = v560 + v575
		v562 = v562 + v575
		v563 = v580
		goto L123
	} else {
		goto L125
	}
L124:
	;
	goto L122
L125:
	;
	goto L124
L126:
	;
	v618 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v618
	*(*int32)(unsafe.Add(mBase, uint32(v604)+16)) = v618
	v851 = v604
	goto L109
L127:
	;
	goto L128
L128:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v604)+12))
	v623 = v490 - v292
	v624 = v622 + v623
	if v624 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v86)+36))
	v628 = v622
	v630 = v604
	v632 = v627
	goto L132
L130:
	;
	v679 = v604
	v680 = v624
	goto L131
L131:
	;
	v690 = int32(16777216)
	if base.Ui32(v690) <= base.Ui32(v680) {
		goto L148
	} else {
		goto L149
	}
L132:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v630)+20))
	if v641 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v679 = v672
	v680 = v674
	goto L131
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = int32(0)
	goto L43
L135:
	;
	goto L136
L136:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v630)+4))
	if base.Ui32(int32(2147483646)) <= base.Ui32(v628) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v652 = int32(-3)
	goto L139
L138:
	;
	v652 = v628 << (uint(int32(1)) % 32)
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+12)) = v652
	v655 = v652 + int32(2)
	if v646 != 0 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v630)+4)) = v660
	if v660 == int32(0) {
		goto L43
	} else {
		goto L146
	}
L141:
	;
	v656 = F_repalloc(m, v646, v655)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L19
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v658 = F_palloc(m, v655)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L19
	} else {
		goto L145
	}
L144:
	;
	v660 = v656
	goto L140
L145:
	;
	v660 = v658
	goto L140
L146:
	;
	v665 = v660 + (v632 - v646)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v665
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v667+v668<<(uint(int32(2))%32))))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v674 = v673 + v623
	if v674 == int32(0) {
		v628 = v673
		v630 = v672
		v632 = v665
		goto L132
	} else {
		goto L147
	}
L147:
	;
	goto L133
L148:
	;
	v693 = v690
	goto L150
L149:
	;
	v693 = v680
	goto L150
L150:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v679)+24))
	if v695 != 0 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v839
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v841+v842<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v846)+16)) = v839
	if v839 != 0 {
		v879 = v839
		v881 = v846
		v892 = int32(0)
		goto L108
	} else {
		goto L189
	}
L152:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v826+v827<<(uint(int32(2))%32))))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+4))
	v835 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v832+v499+v727))) = uint8(v835)
	v839 = v727 + int32(1)
	goto L151
L153:
	;
	v696 = int32(0)
	goto L157
L154:
	;
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(0)
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v751 = F_fread(m, v747+v499, int32(1), v693, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L19
	} else {
		goto L170
	}
L156:
	;
	switch v713 {
	case 0:
		goto L162
	default:
		v839 = v727
		goto L151
	case 11:
		goto L152
	}
L157:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v710 = F_do_getc(m, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L19
	} else {
		goto L160
	}
L158:
	;
	v727 = v693
	goto L156
L159:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v714+v715<<(uint(int32(2))%32))))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v719)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v720+v499+v696))) = uint8(v710)
	v725 = v696 + int32(1)
	if v725 != v693 {
		v696 = v725
		goto L157
	} else {
		goto L161
	}
L160:
	;
	v713 = v710 + int32(1)
	switch v713 {
	case 0, 11:
		v727 = v696
		goto L156
	default:
		goto L159
	}
L161:
	;
	goto L158
L162:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v728)+76))
	if v729 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if int32(base.Ui32(v734)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		v839 = v727
		goto L151
	} else {
		goto L168
	}
L164:
	;
	goto L163
L165:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v734 = v732
	goto L164
L166:
	;
	goto L167
L167:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v728)))
	v734 = v733
	goto L164
L168:
	;
	F_yy_fatal_error_6(m, int32(453961))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L19
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	v753 = v751
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v753
	if v753 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v818+v819<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v823)+16)) = v753
	v879 = v753
	v881 = v823
	v892 = int32(0)
	goto L108
L173:
	;
	goto L172
L174:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+76))
	if v768 < int32(0) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	if int32(base.Ui32(v773)>>(uint(int32(5))%32))&int32(1) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L176:
	;
	goto L175
L177:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v773 = v771
	goto L176
L178:
	;
	goto L179
L179:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	v773 = v772
	goto L176
L180:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v780+v781<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v785)+16)) = int32(0)
	v851 = v785
	goto L109
L181:
	;
	goto L182
L182:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v789 != int32(27) {
		goto L45
	} else {
		goto L183
	}
L183:
	;
	v793 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[86])) = v793
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v767)+76))
	if v793 <= v795 {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v807 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v806+v807<<(uint(int32(2))%32))))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	v816 = F_fread(m, v812+v499, int32(1), v693, v815)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L19
	} else {
		goto L188
	}
L185:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v798 & int32(-49)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v767)))
	*(*int32)(unsafe.Add(mBase, uint32(v767))) = v802 & int32(-49)
	goto L184
L188:
	;
	v753 = v816
	goto L171
L189:
	;
	v851 = v846
	goto L109
L190:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	F_cube_yyrestart(m, v864, v86)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L19
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	v875 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v851)+44)) = v875
	v879 = int32(0)
	v881 = v851
	v892 = v875
	goto L108
L193:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v867+v868<<(uint(int32(2))%32))))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v879 = v873
	v881 = v872
	v892 = int32(1)
	goto L108
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v922
	v925 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v920+v922))) = uint8(v925)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v929 = int32(2)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v927+v928<<(uint(v929)%32))))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v932)+4))
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v933+v934)+1)) = uint8(v925)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v938+v939<<(uint(v929)%32))))
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+80)) = v944
	if v892 == int32(1) {
		v1188 = v944
		goto L42
	} else {
		goto L205
	}
L195:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v881)+4))
	v920 = v896
	v922 = v893
	goto L194
L196:
	;
	goto L197
L197:
	;
	v899 = v893 + int32(base.Ui32(v879)>>(uint(int32(1))%32))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v881)+4))
	if v900 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v906+v907<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v911)+4)) = v905
	if v905 == int32(0) {
		goto L44
	} else {
		goto L204
	}
L199:
	;
	v901 = F_repalloc(m, v900, v899)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L19
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v903 = F_palloc(m, v899)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L19
	} else {
		goto L203
	}
L202:
	;
	v905 = v901
	goto L198
L203:
	;
	v905 = v903
	goto L198
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v911)+12)) = v899 - int32(2)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v920 = v905
	v922 = v918 + v499
	goto L194
L205:
	;
	switch v892 - int32(1) {
	case 0:
		goto L102
	case 1:
		goto L206
	default:
		goto L207
	}
L206:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
	v1051 = v944
	v1055 = v1050
	v1056 = v944
	goto L103
L207:
	;
	v952 = v271 ^ int32(-1) + v197
	v953 = v944 + v952
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v953
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v86)+44))
	if v952 == int32(0) {
		v87 = v955
		v89 = v953
		v90 = v944
		goto L24
	} else {
		goto L208
	}
L208:
	;
	v960 = v955
	v967 = v944
	goto L209
L209:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	if v971 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	v87 = v1046
	v89 = v953
	v90 = v944
	goto L24
L211:
	;
	v974 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v971)+uint32(_consts[1461]))))
	v976 = v974
	goto L213
L212:
	;
	v976 = int32(1)
	goto L213
L213:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v960))%64)&int64(101125980167) == int64(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v967
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v960
	goto L216
L215:
	;
	goto L216
L216:
	;
	v986 = int32(1)
	v990 = int32(*(*int16)(unsafe.Add(mBase, uint32(v960<<(uint(v986)%32))+uint32(_consts[1462]))))
	v991 = v976 + v990
	v996 = int32(*(*int16)(unsafe.Add(mBase, uint32(v991<<(uint(v986)%32))+uint32(_consts[1463]))))
	if v996 != v960 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v1000 = v960
	goto L220
L218:
	;
	v1029 = v991
	goto L219
L219:
	;
	v1042 = int32(1)
	v1046 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1029<<(uint(v1042)%32))+uint32(_consts[1464]))))
	v1048 = v967 + v1042
	if v953 != v1048 {
		v960 = v1046
		v967 = v1048
		goto L209
	} else {
		goto L223
	}
L220:
	;
	v1011 = int32(1)
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1000<<(uint(v1011)%32))+uint32(_consts[1465]))))
	v1016 = base.I32_extend16_s(v1015)
	v1021 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1016<<(uint(v1011)%32))+uint32(_consts[1462]))))
	v1022 = v976 + v1021
	v1027 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1022<<(uint(v1011)%32))+uint32(_consts[1463]))))
	if v1015 != v1027 {
		v1000 = v1016
		goto L220
	} else {
		goto L222
	}
L221:
	;
	v1029 = v1022
	goto L219
L222:
	;
	goto L221
L223:
	;
	goto L210
L224:
	;
	v1070 = v1066
	v1072 = v1056
	goto L225
L225:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072))))
	if v1081 != 0 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v193 = v1156
	v197 = v1064
	v199 = v1056
	goto L38
L227:
	;
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1081)+uint32(_consts[1461]))))
	v1086 = v1084
	goto L229
L228:
	;
	v1086 = int32(1)
	goto L229
L229:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v1070))%64)&int64(101125980167) == int64(0) {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+68)) = v1072
	*(*int32)(unsafe.Add(mBase, uint32(v86)+64)) = v1070
	goto L232
L231:
	;
	goto L232
L232:
	;
	v1096 = int32(1)
	v1100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1070<<(uint(v1096)%32))+uint32(_consts[1462]))))
	v1101 = v1086 + v1100
	v1106 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1101<<(uint(v1096)%32))+uint32(_consts[1463]))))
	if v1106 != v1070 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1110 = v1070
	goto L236
L234:
	;
	v1139 = v1101
	goto L235
L235:
	;
	v1152 = int32(1)
	v1156 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1139<<(uint(v1152)%32))+uint32(_consts[1464]))))
	v1158 = v1072 + v1152
	if v1158 != v1064 {
		v1070 = v1156
		v1072 = v1158
		goto L225
	} else {
		goto L239
	}
L236:
	;
	v1121 = int32(1)
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110<<(uint(v1121)%32))+uint32(_consts[1465]))))
	v1126 = base.I32_extend16_s(v1125)
	v1131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1126<<(uint(v1121)%32))+uint32(_consts[1462]))))
	v1132 = v1086 + v1131
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1132<<(uint(v1121)%32))+uint32(_consts[1463]))))
	if v1125 != v1137 {
		v1110 = v1126
		goto L236
	} else {
		goto L238
	}
L237:
	;
	v1139 = v1132
	goto L235
L238:
	;
	goto L237
L239:
	;
	goto L226
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_yylex_init_extra(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	if l1 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(28)
		return int32(1)
	} else {
		v12 = F_palloc(m, int32(96))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
			if v12 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _consts[86])) = int32(48)
				return int32(1)
			} else {
				v27 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), int32(96))
				mBase = m.M
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				*(*int32)(unsafe.Add(mBase, uint32(v28))) = l0
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v31 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+60)) = v31
				v33 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v30)+52)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v30)+44)) = v31
				*(*int64)(unsafe.Add(mBase, uint32(v30)+36)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v30)+4)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v30)+12)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v31
				return v31
			}
		}
	}
}
func F_cube_yyset_extra(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = l0
	return
}
