package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeSystemUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
	v11 = F_psprintf(m, int32(167372), v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[36]))
		v16 = F_MemoryContextStrdup(m, v15, v11)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1417])) = v16
			F_pfree(m, v11)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_SystemFuncName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = F_makeString(m, int32(312227))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v8
		v13 = F_makeString(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v13
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v17
			v21 = F_list_make2_impl(m, v5+int32(4), v5)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				m.G0 = v5 + int32(16)
				return v21
			}
		}
	}
}
func F_system_nextsampleblock(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int64
	_ = v297
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v15
	v17 = int32(-1)
	if base.Ui32(l1) <= base.Ui32(v14) {
		v305 = v17
		v306 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v306
	m.G0 = v11 + int32(16)
	return v305
L2:
	;
	v19 = v14
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v19
	v29 = v19 + int32(1)
	v31 = v11 + int32(8)
	v38 = int32(-1636608424)
	if v31&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v305 = v17
	v306 = v3
	goto L1
L5:
	;
	v297 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	if base.Ui64(base.I64_extend_i32_u(v292^v284-base.I32_rotl(v292, int32(24)))) < base.Ui64(v297) {
		goto L45
	} else {
		goto L46
	}
L6:
	;
	v270 = int32(14)
	v272 = v266 ^ v267 - base.I32_rotl(v266, v270)
	v276 = v272 ^ v265 - base.I32_rotl(v272, int32(11))
	v280 = v276 ^ v266 - base.I32_rotl(v276, int32(25))
	v284 = v280 ^ v272 - base.I32_rotl(v280, int32(16))
	v288 = v284 ^ v276 - base.I32_rotl(v284, int32(4))
	v292 = v288 ^ v280 - base.I32_rotl(v288, v270)
	goto L5
L7:
	;
	switch int32(7) {
	case 0:
		v258 = v38
		v259 = v38
		v260 = v38
		goto L34
	case 1:
		v251 = v38
		v252 = v38
		v253 = v38
		goto L35
	case 2:
		v244 = v38
		v245 = v38
		v246 = v38
		goto L36
	case 3:
		v238 = v38
		v239 = v38
		goto L37
	case 4:
		v234 = v38
		v235 = v38
		goto L38
	case 5:
		v228 = v38
		v229 = v38
		goto L39
	case 6:
		v222 = v38
		v223 = v38
		goto L40
	case 7:
		v217 = v38
		goto L41
	case 8:
		v212 = v38
		goto L42
	case 9:
		v207 = v38
		goto L43
	case 10:
		goto L44
	default:
		v265 = v38
		v266 = v38
		v267 = v38
		goto L6
	}
L9:
	;
	goto L12
L10:
	;
	goto L11
L11:
	;
	goto L14
L12:
	;
	goto L7
L13:
	;
	switch int32(7) {
	case 0:
		v144 = v38
		goto L20
	case 1:
		v139 = v38
		goto L21
	case 2:
		goto L22
	case 3:
		v132 = v38
		goto L23
	case 4:
		v129 = v38
		goto L24
	case 5:
		v124 = v38
		goto L25
	case 6:
		goto L26
	case 7:
		v115 = v38
		goto L27
	case 8:
		v110 = v38
		goto L28
	case 9:
		v105 = v38
		goto L29
	case 10:
		goto L30
	default:
		v265 = v38
		v266 = v38
		v267 = v38
		goto L6
	}
L14:
	;
	goto L13
L20:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v265 = v144 + v145
	v266 = v38
	v267 = v38
	goto L6
L21:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v144 = v140<<(uint(int32(8))%32) + v139
	goto L20
L22:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
	v139 = v135<<(uint(int32(16))%32) + v38
	goto L21
L23:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v265 = v133 + v38
	v266 = v132
	v267 = v38
	goto L6
L24:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	v132 = v129 + v130
	goto L23
L25:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
	v129 = v125<<(uint(int32(8))%32) + v124
	goto L24
L26:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	v124 = v120<<(uint(int32(16))%32) + v38
	goto L25
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v265 = v116 + v38
	v266 = v118 + v38
	v267 = v115
	goto L6
L28:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+8)))
	v115 = v111<<(uint(int32(8))%32) + v110
	goto L27
L29:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
	v110 = v106<<(uint(int32(16))%32) + v105
	goto L28
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+10)))
	v105 = v101<<(uint(int32(24))%32) + v38
	goto L29
L34:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v265 = v258 + v261
	v266 = v259
	v267 = v260
	goto L6
L35:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v258 = v254<<(uint(int32(8))%32) + v251
	v259 = v252
	v260 = v253
	goto L34
L36:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
	v251 = v247<<(uint(int32(16))%32) + v244
	v252 = v245
	v253 = v246
	goto L35
L37:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
	v244 = v240<<(uint(int32(24))%32) + v38
	v245 = v238
	v246 = v239
	goto L36
L38:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	v238 = v234 + v236
	v239 = v235
	goto L37
L39:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+5)))
	v234 = v230<<(uint(int32(8))%32) + v228
	v235 = v229
	goto L38
L40:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+6)))
	v228 = v224<<(uint(int32(16))%32) + v222
	v229 = v223
	goto L39
L41:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+7)))
	v222 = v218<<(uint(int32(24))%32) + v38
	v223 = v217
	goto L40
L42:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+8)))
	v217 = v213<<(uint(int32(8))%32) + v212
	goto L41
L43:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+9)))
	v212 = v208<<(uint(int32(16))%32) + v207
	goto L42
L44:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+10)))
	v207 = v203<<(uint(int32(24))%32) + v38
	goto L43
L45:
	;
	v305 = v19
	v306 = v29
	goto L1
L46:
	;
	goto L47
L47:
	;
	if v29 != l1 {
		v19 = v29
		goto L3
	} else {
		goto L48
	}
L48:
	;
	goto L4
}
func F_system_samplescangetsamplesize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 float32
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 float32
	_ = v19
	var v35 float32
	_ = v35
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v50 float64
	_ = v50
	var v54 float64
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	v9 = float32(0.1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_estimate_expression_value(m, l0, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v14 != int32(7) {
			v35 = v9
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v17 != 0 {
				v35 = v9
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v19 = base.F32_reinterpret_i32(v18)
				if base.F32_ge(v19, float32(0)) == int32(0) {
					v35 = v9
				} else {
					if base.F32_le(v19, float32(100)) == int32(0) {
						v35 = v9
					} else {
						if base.Ui32(int32(2139095040)) < base.Ui32(v18&int32(2147483647)) {
							v35 = v9
						} else {
							v35 = base.F32_div(v19, float32(100))
						}
					}
				}
			}
		}
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		v40 = base.F64_promote_f32(base.F32_mul(v35, base.F32_convert_i32_u(v37)))
		v42 = float64(1e+100)
		if base.F64_gt(v40, v42) != 0 {
			v54 = v42
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)) {
				v54 = v42
			} else {
				v50 = float64(1)
				if base.F64_le(v40, v50) != 0 {
					v54 = v50
				} else {
					v54 = base.F64_nearest(v40)
				}
			}
		}
		if base.F64_lt(v54, float64(4.294967296e+09))&base.F64_ge(v54, float64(0)) != 0 {
			v60 = base.I32_trunc_f64_u(v54)
			v62 = v60
		} else {
			v62 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l3))) = v62
		v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v66 = base.F64_mul(v64, base.F64_promote_f32(v35))
		v68 = float64(1e+100)
		if base.F64_gt(v66, v68) != 0 {
			v80 = v68
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
				v80 = v68
			} else {
				v76 = float64(1)
				if base.F64_le(v66, v76) != 0 {
					v80 = v76
				} else {
					v80 = base.F64_nearest(v66)
				}
			}
		}
		*(*float64)(unsafe.Add(mBase, uint32(l4))) = v80
		return
	}
}
