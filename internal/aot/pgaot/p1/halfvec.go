package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HalfvecItemSize(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = F_mul_size(m, int32(2), l0)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_add_size(m, int32(8), v4)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_HalfvecUpdateCenter(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v42 float32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 float32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v2 = l1
	v13 = F_mul_size(m, int32(2), v2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_add_size(m, int32(8), v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v2)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15 << (uint(int32(2)) % 32)
			if int32(0) < v2 {
				v26 = int32(0)
				for {
					v42 = *(*float32)(unsafe.Add(mBase, uint32(l2+v26<<(uint(int32(2))%32))))
					v43 = base.I32_reinterpret_f32(v42)
					v45 = int32(base.Ui32(v43) >> (uint(int32(16)) % 32))
					v48 = base.F32_abs(v42)
					if base.F32_eq(v48, math.Float32frombits(uint32(0x7f800000))) != 0 {
						v127 = v45 & int32(_a_F_HalfvecUpdateCenter_0)
					} else {
						v52 = v45 & int32(_a_F_HalfvecUpdateCenter_1)
						v54 = v43 & int32(_a_F_HalfvecUpdateCenter_2)
						if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v48)) {
							v127 = v52 | int32(base.Ui32(v54)>>(uint(int32(13))%32)) | int32(_a_F_HalfvecUpdateCenter_3)
						} else {
							v66 = int32(base.Ui32(v43)>>(uint(int32(23))%32)) & int32(255)
							if base.Ui32(v66) < base.Ui32(int32(99)) {
								v127 = v52
							} else {
								if base.Ui32(v66) <= base.Ui32(int32(112)) {
									v78 = int32(1)<<(uint(v66-int32(90))%32) + int32(base.Ui32(v54)>>(uint(int32(113)-v66)%32))
									v80 = v78 | v43
									v81 = v78
								} else {
									v80 = v43
									v81 = v54
								}
								v83 = int32(base.Ui32(v81) >> (uint(int32(13)) % 32))
								v88 = int32(1)
								v92 = int32(3)
								v93 = int32(base.Ui32(v81)>>(uint(int32(12))%32)) & v92
								if base.B2i32(v93 != v92)&(base.B2i32(v80&int32(4095) == int32(0))|base.B2i32(v93 != v88)) != 0 {
									v104 = v83
								} else {
									v104 = v83 + v88
								}
								v106 = base.B2i32(v104 == int32(1024))
								if v104 == int32(1024) {
									v107 = int32(-126)
								} else {
									v107 = int32(-127)
								}
								v108 = v107 + v66
								if int32(16) <= v108 {
									v127 = v52 | int32(_a_F_HalfvecUpdateCenter_4)
								} else {
									if int32(-15) < v108 {
										v118 = v108<<(uint(int32(10))%32) + int32(_a_F_HalfvecUpdateCenter_5) | v52
									} else {
										v118 = v52
									}
									if v104 == int32(1024) {
										v120 = int32(0)
									} else {
										v120 = v104
									}
									v127 = v118 | v120
								}
							}
						}
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(8)+v26<<(uint(int32(1))%32)))) = uint16(v127)
					v130 = v26 + int32(1)
					if v130 != v2 {
						v26 = v130
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return
		}
	}
}
func F_halfvec(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13912(m, l0, int32(92), int32(_a_F_halfvec_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_halfvec_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 float32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 float32
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v384 int32
	_ = v384
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	if v27 == v28 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v17 + int32(16)
	return v37
L5:
	;
	v363 = int32(0)
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v364 <= v363 {
		goto L4
	} else {
		goto L126
	}
L6:
	;
	v33 = F_mul_size(m, int32(2), base.I32_extend16_s(v27))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L122
	}
L9:
	;
	v35 = F_add_size(m, int32(8), v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v37 = F_palloc0(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v35 << (uint(int32(2)) % 32)
	v43 = int32(0)
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v44 <= v43 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v47 = int32(8)
	v52 = v37 + v47
	v53 = v43
	goto L13
L13:
	;
	v68 = v53 << (uint(int32(1)) % 32)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v47+v68))))
	v75 = v70 & int32(1023)
	v79 = v70 << (uint(int32(16)) % 32) & int32(-2147483648)
	v82 = int32(31)
	v83 = int32(base.Ui32(v70)>>(uint(int32(10))%32)) & v82
	if v83 != v82 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	goto L5
L15:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+(v25+v47)))))
	v167 = v162 & int32(1023)
	v171 = v162 << (uint(int32(16)) % 32) & int32(-2147483648)
	v174 = int32(31)
	v175 = int32(base.Ui32(v162)>>(uint(int32(10))%32)) & v174
	if v175 != v174 {
		goto L62
	} else {
		goto L63
	}
L16:
	;
	goto L15
L17:
	;
	v155 = v75
	v156 = v83<<(uint(int32(23))%32) + v79 + int32(939524096)
	goto L16
L18:
	;
	if v70&int32(512) != 0 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	if v83 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v75 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v75 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v155 = int32(0)
	v156 = v79
	goto L16
L24:
	;
	v155 = int32(0)
	v156 = v79 | int32(2139095040)
	goto L16
L25:
	;
	goto L26
L26:
	;
	v155 = v75
	v156 = v79 | int32(2143289344)
	goto L16
L27:
	;
	v155 = v143 & int32(1022)
	v156 = v145 | v79
	goto L16
L28:
	;
	v143 = v75 << (uint(int32(1)) % 32)
	v145 = int32(939524096)
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(255)) < base.Ui32(v75) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v143 = v75 << (uint(int32(2)) % 32)
	v145 = int32(931135488)
	goto L27
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(int32(127)) < base.Ui32(v75) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v143 = v75 << (uint(int32(3)) % 32)
	v145 = int32(922746880)
	goto L27
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(63)) < base.Ui32(v75) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v143 = v75 << (uint(int32(4)) % 32)
	v145 = int32(914358272)
	goto L27
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(31)) < base.Ui32(v75) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v143 = v75 << (uint(int32(5)) % 32)
	v145 = int32(905969664)
	goto L27
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(15)) < base.Ui32(v75) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v143 = v75 << (uint(int32(6)) % 32)
	v145 = int32(897581056)
	goto L27
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(7)) < base.Ui32(v75) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v143 = v75 << (uint(int32(7)) % 32)
	v145 = int32(889192448)
	goto L27
L47:
	;
	goto L48
L48:
	;
	if base.Ui32(int32(3)) < base.Ui32(v75) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v143 = v75 << (uint(int32(8)) % 32)
	v145 = int32(880803840)
	goto L27
L50:
	;
	goto L51
L51:
	;
	v138 = base.B2i32(v75 == int32(1))
	if v75 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v139 = int32(1024)
	goto L54
L53:
	;
	v139 = v75 << (uint(int32(9)) % 32)
	goto L54
L54:
	;
	if v75 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v142 = int32(864026624)
	goto L57
L56:
	;
	v142 = int32(872415232)
	goto L57
L57:
	;
	v143 = v139
	v145 = v142
	goto L27
L58:
	;
	v253 = base.F32_add(base.F32_reinterpret_i32(v156|v155<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v248|v247<<(uint(int32(13))%32)))
	v254 = base.I32_reinterpret_f32(v253)
	v256 = int32(base.Ui32(v254) >> (uint(int32(16)) % 32))
	v260 = base.F32_abs(v253)
	if base.F32_eq(v260, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v339 = v256 & int32(_a_F_halfvec_add_0)
		goto L101
	} else {
		goto L102
	}
L59:
	;
	goto L58
L60:
	;
	v247 = v167
	v248 = v175<<(uint(int32(23))%32) + v171 + int32(939524096)
	goto L59
L61:
	;
	if v162&int32(512) != 0 {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	if v175 != 0 {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v167 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v167 != 0 {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v247 = int32(0)
	v248 = v171
	goto L59
L67:
	;
	v247 = int32(0)
	v248 = v171 | int32(2139095040)
	goto L59
L68:
	;
	goto L69
L69:
	;
	v247 = v167
	v248 = v171 | int32(2143289344)
	goto L59
L70:
	;
	v247 = v235 & int32(1022)
	v248 = v237 | v171
	goto L59
L71:
	;
	v235 = v167 << (uint(int32(1)) % 32)
	v237 = int32(939524096)
	goto L70
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(255)) < base.Ui32(v167) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v235 = v167 << (uint(int32(2)) % 32)
	v237 = int32(931135488)
	goto L70
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(127)) < base.Ui32(v167) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v235 = v167 << (uint(int32(3)) % 32)
	v237 = int32(922746880)
	goto L70
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(63)) < base.Ui32(v167) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v235 = v167 << (uint(int32(4)) % 32)
	v237 = int32(914358272)
	goto L70
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(31)) < base.Ui32(v167) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v235 = v167 << (uint(int32(5)) % 32)
	v237 = int32(905969664)
	goto L70
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(15)) < base.Ui32(v167) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v235 = v167 << (uint(int32(6)) % 32)
	v237 = int32(897581056)
	goto L70
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(7)) < base.Ui32(v167) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v235 = v167 << (uint(int32(7)) % 32)
	v237 = int32(889192448)
	goto L70
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(int32(3)) < base.Ui32(v167) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v235 = v167 << (uint(int32(8)) % 32)
	v237 = int32(880803840)
	goto L70
L93:
	;
	goto L94
L94:
	;
	v230 = base.B2i32(v167 == int32(1))
	if v167 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v231 = int32(1024)
	goto L97
L96:
	;
	v231 = v167 << (uint(int32(9)) % 32)
	goto L97
L97:
	;
	if v167 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v234 = int32(864026624)
	goto L100
L99:
	;
	v234 = int32(872415232)
	goto L100
L100:
	;
	v235 = v231
	v237 = v234
	goto L70
L101:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v68+v52))) = uint16(v339)
	v342 = v53 + int32(1)
	if v44 != v342 {
		v53 = v342
		goto L13
	} else {
		goto L121
	}
L102:
	;
	v264 = v256 & int32(_a_F_halfvec_add_1)
	v266 = v254 & int32(_a_F_halfvec_add_2)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v260)) {
		v339 = v264 | int32(base.Ui32(v266)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_add_3)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v278 = int32(base.Ui32(v254)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v278) < base.Ui32(int32(99)) {
		v339 = v264
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if base.Ui32(v278) <= base.Ui32(int32(112)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v290 = int32(1)<<(uint(v278-int32(90))%32) + int32(base.Ui32(v266)>>(uint(int32(113)-v278)%32))
	v292 = v290 | v254
	v293 = v290
	goto L107
L106:
	;
	v292 = v254
	v293 = v266
	goto L107
L107:
	;
	v295 = int32(base.Ui32(v293) >> (uint(int32(13)) % 32))
	v300 = int32(1)
	v304 = int32(3)
	v305 = int32(base.Ui32(v293)>>(uint(int32(12))%32)) & v304
	if base.B2i32(v305 != v304)&(base.B2i32(v292&int32(4095) == int32(0))|base.B2i32(v305 != v300)) != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v316 = v295
	goto L110
L109:
	;
	v316 = v295 + v300
	goto L110
L110:
	;
	v318 = base.B2i32(v316 == int32(1024))
	if v316 == int32(1024) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v319 = int32(-126)
	goto L113
L112:
	;
	v319 = int32(-127)
	goto L113
L113:
	;
	v320 = v319 + v278
	if int32(16) <= v320 {
		v339 = v264 | int32(_a_F_halfvec_add_4)
		goto L101
	} else {
		goto L114
	}
L114:
	;
	if int32(-15) < v320 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v330 = v320<<(uint(int32(10))%32) + int32(_a_F_halfvec_add_5) | v264
	goto L117
L116:
	;
	v330 = v264
	goto L117
L117:
	;
	if v316 == int32(1024) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v332 = int32(0)
	goto L120
L119:
	;
	v332 = v316
	goto L120
L120:
	;
	v339 = v330 | v332
	goto L101
L121:
	;
	goto L14
L122:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v351 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	v352 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v352
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v351
	F_errmsg(m, int32(_a_F_halfvec_add_6), v17)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_halfvec_add_7), int32(80), int32(_a_F_halfvec_add_8))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v367 = v363
	goto L127
L127:
	;
	v384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v367<<(uint(int32(1))%32)))))
	if v384&int32(_a_F_halfvec_add_9) != int32(_a_F_halfvec_add_4) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L133
	}
L129:
	;
	v390 = v367 + int32(1)
	if v364 != v390 {
		v367 = v390
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	goto L4
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_halfvec_binary_quantize(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
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
	var v250 int32
	_ = v250
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v828 int32
	_ = v828
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(8)
		v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
		v16 = F_InitBitVector(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(8)
			v19 = v16 + v18
			v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
			v22 = base.I32_div_s(v20, v18)
			v23 = int32(0)
			if v18 <= v20 {
				v29 = v23
				for {
					v43 = v14 + v29<<(uint(int32(1))%32)
					v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
					v49 = v44 & int32(1023)
					v53 = v44 << (uint(int32(16)) % 32) & int32(-2147483648)
					v56 = int32(31)
					v57 = int32(base.Ui32(v44)>>(uint(int32(10))%32)) & v56
					if v57 != v56 {
						if v57 != 0 {
							v129 = v49
							v130 = v57<<(uint(int32(23))%32) + v53 + int32(939524096)
						} else {
							if v49 != 0 {
								if v44&int32(512) != 0 {
									v117 = v49 << (uint(int32(1)) % 32)
									v119 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v49) {
										v117 = v49 << (uint(int32(2)) % 32)
										v119 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v49) {
											v117 = v49 << (uint(int32(3)) % 32)
											v119 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v49) {
												v117 = v49 << (uint(int32(4)) % 32)
												v119 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v49) {
													v117 = v49 << (uint(int32(5)) % 32)
													v119 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v49) {
														v117 = v49 << (uint(int32(6)) % 32)
														v119 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v49) {
															v117 = v49 << (uint(int32(7)) % 32)
															v119 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v49) {
																v117 = v49 << (uint(int32(8)) % 32)
																v119 = int32(880803840)
															} else {
																v112 = base.B2i32(v49 == int32(1))
																if v49 == int32(1) {
																	v113 = int32(1024)
																} else {
																	v113 = v49 << (uint(int32(9)) % 32)
																}
																if v49 == int32(1) {
																	v116 = int32(864026624)
																} else {
																	v116 = int32(872415232)
																}
																v117 = v113
																v119 = v116
															}
														}
													}
												}
											}
										}
									}
								}
								v129 = v117 & int32(1022)
								v130 = v119 | v53
							} else {
								v129 = int32(0)
								v130 = v53
							}
						}
					} else {
						if v49 == int32(0) {
							v129 = int32(0)
							v130 = v53 | int32(2139095040)
						} else {
							v129 = v49
							v130 = v53 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v130|v129<<(uint(int32(13))%32)), float32(0)) != 0 {
						v137 = int32(-128)
					} else {
						v137 = int32(0)
					}
					v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+2)))
					v145 = v140 & int32(1023)
					v149 = v140 << (uint(int32(16)) % 32) & int32(-2147483648)
					v152 = int32(31)
					v153 = int32(base.Ui32(v140)>>(uint(int32(10))%32)) & v152
					if v153 != v152 {
						if v153 != 0 {
							v225 = v145
							v226 = v153<<(uint(int32(23))%32) + v149 + int32(939524096)
						} else {
							if v145 != 0 {
								if v140&int32(512) != 0 {
									v213 = v145 << (uint(int32(1)) % 32)
									v215 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v145) {
										v213 = v145 << (uint(int32(2)) % 32)
										v215 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v145) {
											v213 = v145 << (uint(int32(3)) % 32)
											v215 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v145) {
												v213 = v145 << (uint(int32(4)) % 32)
												v215 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v145) {
													v213 = v145 << (uint(int32(5)) % 32)
													v215 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v145) {
														v213 = v145 << (uint(int32(6)) % 32)
														v215 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v145) {
															v213 = v145 << (uint(int32(7)) % 32)
															v215 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v145) {
																v213 = v145 << (uint(int32(8)) % 32)
																v215 = int32(880803840)
															} else {
																v208 = base.B2i32(v145 == int32(1))
																if v145 == int32(1) {
																	v209 = int32(1024)
																} else {
																	v209 = v145 << (uint(int32(9)) % 32)
																}
																if v145 == int32(1) {
																	v212 = int32(864026624)
																} else {
																	v212 = int32(872415232)
																}
																v213 = v209
																v215 = v212
															}
														}
													}
												}
											}
										}
									}
								}
								v225 = v213 & int32(1022)
								v226 = v215 | v149
							} else {
								v225 = int32(0)
								v226 = v149
							}
						}
					} else {
						if v145 == int32(0) {
							v225 = int32(0)
							v226 = v149 | int32(2139095040)
						} else {
							v225 = v145
							v226 = v149 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v226|v225<<(uint(int32(13))%32)), float32(0)) != 0 {
						v233 = int32(64)
					} else {
						v233 = int32(0)
					}
					v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+4)))
					v242 = v237 & int32(1023)
					v246 = v237 << (uint(int32(16)) % 32) & int32(-2147483648)
					v249 = int32(31)
					v250 = int32(base.Ui32(v237)>>(uint(int32(10))%32)) & v249
					if v250 != v249 {
						if v250 != 0 {
							v322 = v242
							v323 = v250<<(uint(int32(23))%32) + v246 + int32(939524096)
						} else {
							if v242 != 0 {
								if v237&int32(512) != 0 {
									v310 = v242 << (uint(int32(1)) % 32)
									v312 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v242) {
										v310 = v242 << (uint(int32(2)) % 32)
										v312 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v242) {
											v310 = v242 << (uint(int32(3)) % 32)
											v312 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v242) {
												v310 = v242 << (uint(int32(4)) % 32)
												v312 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v242) {
													v310 = v242 << (uint(int32(5)) % 32)
													v312 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v242) {
														v310 = v242 << (uint(int32(6)) % 32)
														v312 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v242) {
															v310 = v242 << (uint(int32(7)) % 32)
															v312 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v242) {
																v310 = v242 << (uint(int32(8)) % 32)
																v312 = int32(880803840)
															} else {
																v305 = base.B2i32(v242 == int32(1))
																if v242 == int32(1) {
																	v306 = int32(1024)
																} else {
																	v306 = v242 << (uint(int32(9)) % 32)
																}
																if v242 == int32(1) {
																	v309 = int32(864026624)
																} else {
																	v309 = int32(872415232)
																}
																v310 = v306
																v312 = v309
															}
														}
													}
												}
											}
										}
									}
								}
								v322 = v310 & int32(1022)
								v323 = v312 | v246
							} else {
								v322 = int32(0)
								v323 = v246
							}
						}
					} else {
						if v242 == int32(0) {
							v322 = int32(0)
							v323 = v246 | int32(2139095040)
						} else {
							v322 = v242
							v323 = v246 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v323|v322<<(uint(int32(13))%32)), float32(0)) != 0 {
						v330 = int32(32)
					} else {
						v330 = int32(0)
					}
					v332 = int32(16)
					v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+6)))
					v339 = v334 & int32(1023)
					v343 = v334 << (uint(v332) % 32) & int32(-2147483648)
					v346 = int32(31)
					v347 = int32(base.Ui32(v334)>>(uint(int32(10))%32)) & v346
					if v347 != v346 {
						if v347 != 0 {
							v419 = v339
							v420 = v347<<(uint(int32(23))%32) + v343 + int32(939524096)
						} else {
							if v339 != 0 {
								if v334&int32(512) != 0 {
									v407 = v339 << (uint(int32(1)) % 32)
									v409 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v339) {
										v407 = v339 << (uint(int32(2)) % 32)
										v409 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v339) {
											v407 = v339 << (uint(int32(3)) % 32)
											v409 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v339) {
												v407 = v339 << (uint(int32(4)) % 32)
												v409 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v339) {
													v407 = v339 << (uint(int32(5)) % 32)
													v409 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v339) {
														v407 = v339 << (uint(int32(6)) % 32)
														v409 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v339) {
															v407 = v339 << (uint(int32(7)) % 32)
															v409 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v339) {
																v407 = v339 << (uint(int32(8)) % 32)
																v409 = int32(880803840)
															} else {
																v402 = base.B2i32(v339 == int32(1))
																if v339 == int32(1) {
																	v403 = int32(1024)
																} else {
																	v403 = v339 << (uint(int32(9)) % 32)
																}
																if v339 == int32(1) {
																	v406 = int32(864026624)
																} else {
																	v406 = int32(872415232)
																}
																v407 = v403
																v409 = v406
															}
														}
													}
												}
											}
										}
									}
								}
								v419 = v407 & int32(1022)
								v420 = v409 | v343
							} else {
								v419 = int32(0)
								v420 = v343
							}
						}
					} else {
						if v339 == int32(0) {
							v419 = int32(0)
							v420 = v343 | int32(2139095040)
						} else {
							v419 = v339
							v420 = v343 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v420|v419<<(uint(int32(13))%32)), float32(0)) != 0 {
						v427 = v332
					} else {
						v427 = int32(0)
					}
					v431 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
					v436 = v431 & int32(1023)
					v440 = v431 << (uint(int32(16)) % 32) & int32(-2147483648)
					v443 = int32(31)
					v444 = int32(base.Ui32(v431)>>(uint(int32(10))%32)) & v443
					if v444 != v443 {
						if v444 != 0 {
							v516 = v436
							v517 = v444<<(uint(int32(23))%32) + v440 + int32(939524096)
						} else {
							if v436 != 0 {
								if v431&int32(512) != 0 {
									v504 = v436 << (uint(int32(1)) % 32)
									v506 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v436) {
										v504 = v436 << (uint(int32(2)) % 32)
										v506 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v436) {
											v504 = v436 << (uint(int32(3)) % 32)
											v506 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v436) {
												v504 = v436 << (uint(int32(4)) % 32)
												v506 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v436) {
													v504 = v436 << (uint(int32(5)) % 32)
													v506 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v436) {
														v504 = v436 << (uint(int32(6)) % 32)
														v506 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v436) {
															v504 = v436 << (uint(int32(7)) % 32)
															v506 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v436) {
																v504 = v436 << (uint(int32(8)) % 32)
																v506 = int32(880803840)
															} else {
																v499 = base.B2i32(v436 == int32(1))
																if v436 == int32(1) {
																	v500 = int32(1024)
																} else {
																	v500 = v436 << (uint(int32(9)) % 32)
																}
																if v436 == int32(1) {
																	v503 = int32(864026624)
																} else {
																	v503 = int32(872415232)
																}
																v504 = v500
																v506 = v503
															}
														}
													}
												}
											}
										}
									}
								}
								v516 = v504 & int32(1022)
								v517 = v506 | v440
							} else {
								v516 = int32(0)
								v517 = v440
							}
						}
					} else {
						if v436 == int32(0) {
							v516 = int32(0)
							v517 = v440 | int32(2139095040)
						} else {
							v516 = v436
							v517 = v440 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v517|v516<<(uint(int32(13))%32)), float32(0)) != 0 {
						v524 = int32(8)
					} else {
						v524 = int32(0)
					}
					v528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+10)))
					v533 = v528 & int32(1023)
					v537 = v528 << (uint(int32(16)) % 32) & int32(-2147483648)
					v540 = int32(31)
					v541 = int32(base.Ui32(v528)>>(uint(int32(10))%32)) & v540
					if v541 != v540 {
						if v541 != 0 {
							v613 = v533
							v614 = v541<<(uint(int32(23))%32) + v537 + int32(939524096)
						} else {
							if v533 != 0 {
								if v528&int32(512) != 0 {
									v601 = v533 << (uint(int32(1)) % 32)
									v603 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v533) {
										v601 = v533 << (uint(int32(2)) % 32)
										v603 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v533) {
											v601 = v533 << (uint(int32(3)) % 32)
											v603 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v533) {
												v601 = v533 << (uint(int32(4)) % 32)
												v603 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v533) {
													v601 = v533 << (uint(int32(5)) % 32)
													v603 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v533) {
														v601 = v533 << (uint(int32(6)) % 32)
														v603 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v533) {
															v601 = v533 << (uint(int32(7)) % 32)
															v603 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v533) {
																v601 = v533 << (uint(int32(8)) % 32)
																v603 = int32(880803840)
															} else {
																v596 = base.B2i32(v533 == int32(1))
																if v533 == int32(1) {
																	v597 = int32(1024)
																} else {
																	v597 = v533 << (uint(int32(9)) % 32)
																}
																if v533 == int32(1) {
																	v600 = int32(864026624)
																} else {
																	v600 = int32(872415232)
																}
																v601 = v597
																v603 = v600
															}
														}
													}
												}
											}
										}
									}
								}
								v613 = v601 & int32(1022)
								v614 = v603 | v537
							} else {
								v613 = int32(0)
								v614 = v537
							}
						}
					} else {
						if v533 == int32(0) {
							v613 = int32(0)
							v614 = v537 | int32(2139095040)
						} else {
							v613 = v533
							v614 = v537 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v614|v613<<(uint(int32(13))%32)), float32(0)) != 0 {
						v621 = int32(4)
					} else {
						v621 = int32(0)
					}
					v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+12)))
					v630 = v625 & int32(1023)
					v634 = v625 << (uint(int32(16)) % 32) & int32(-2147483648)
					v637 = int32(31)
					v638 = int32(base.Ui32(v625)>>(uint(int32(10))%32)) & v637
					if v638 != v637 {
						if v638 != 0 {
							v710 = v630
							v711 = v638<<(uint(int32(23))%32) + v634 + int32(939524096)
						} else {
							if v630 != 0 {
								if v625&int32(512) != 0 {
									v698 = v630 << (uint(int32(1)) % 32)
									v700 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v630) {
										v698 = v630 << (uint(int32(2)) % 32)
										v700 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v630) {
											v698 = v630 << (uint(int32(3)) % 32)
											v700 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v630) {
												v698 = v630 << (uint(int32(4)) % 32)
												v700 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v630) {
													v698 = v630 << (uint(int32(5)) % 32)
													v700 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v630) {
														v698 = v630 << (uint(int32(6)) % 32)
														v700 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v630) {
															v698 = v630 << (uint(int32(7)) % 32)
															v700 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v630) {
																v698 = v630 << (uint(int32(8)) % 32)
																v700 = int32(880803840)
															} else {
																v693 = base.B2i32(v630 == int32(1))
																if v630 == int32(1) {
																	v694 = int32(1024)
																} else {
																	v694 = v630 << (uint(int32(9)) % 32)
																}
																if v630 == int32(1) {
																	v697 = int32(864026624)
																} else {
																	v697 = int32(872415232)
																}
																v698 = v694
																v700 = v697
															}
														}
													}
												}
											}
										}
									}
								}
								v710 = v698 & int32(1022)
								v711 = v700 | v634
							} else {
								v710 = int32(0)
								v711 = v634
							}
						}
					} else {
						if v630 == int32(0) {
							v710 = int32(0)
							v711 = v634 | int32(2139095040)
						} else {
							v710 = v630
							v711 = v634 | int32(2143289344)
						}
					}
					if base.F32_gt(base.F32_reinterpret_i32(v711|v710<<(uint(int32(13))%32)), float32(0)) != 0 {
						v718 = int32(2)
					} else {
						v718 = int32(0)
					}
					v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+14)))
					v725 = v720 & int32(1023)
					v729 = v720 << (uint(int32(16)) % 32) & int32(-2147483648)
					v732 = int32(31)
					v733 = int32(base.Ui32(v720)>>(uint(int32(10))%32)) & v732
					if v733 != v732 {
						if v733 != 0 {
							v805 = v725
							v806 = v733<<(uint(int32(23))%32) + v729 + int32(939524096)
						} else {
							if v725 != 0 {
								if v720&int32(512) != 0 {
									v793 = v725 << (uint(int32(1)) % 32)
									v795 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v725) {
										v793 = v725 << (uint(int32(2)) % 32)
										v795 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v725) {
											v793 = v725 << (uint(int32(3)) % 32)
											v795 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v725) {
												v793 = v725 << (uint(int32(4)) % 32)
												v795 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v725) {
													v793 = v725 << (uint(int32(5)) % 32)
													v795 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v725) {
														v793 = v725 << (uint(int32(6)) % 32)
														v795 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v725) {
															v793 = v725 << (uint(int32(7)) % 32)
															v795 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v725) {
																v793 = v725 << (uint(int32(8)) % 32)
																v795 = int32(880803840)
															} else {
																v788 = base.B2i32(v725 == int32(1))
																if v725 == int32(1) {
																	v789 = int32(1024)
																} else {
																	v789 = v725 << (uint(int32(9)) % 32)
																}
																if v725 == int32(1) {
																	v792 = int32(864026624)
																} else {
																	v792 = int32(872415232)
																}
																v793 = v789
																v795 = v792
															}
														}
													}
												}
											}
										}
									}
								}
								v805 = v793 & int32(1022)
								v806 = v795 | v729
							} else {
								v805 = int32(0)
								v806 = v729
							}
						}
					} else {
						if v725 == int32(0) {
							v805 = int32(0)
							v806 = v729 | int32(2139095040)
						} else {
							v805 = v725
							v806 = v729 | int32(2143289344)
						}
					}
					v813 = v137 | v233 | v330 | v427 | v524 | v621 | v718 | base.F32_gt(base.F32_reinterpret_i32(v806|v805<<(uint(int32(13))%32)), float32(0))
					*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v29)>>(uint(int32(3))%32))))) = uint8(v813)
					v816 = v29 + int32(8)
					if v816 < base.I32_extend16_s(v22)<<(uint(int32(3))%32) {
						v29 = v816
						continue
					} else {
						break
					}
					break
				}
				v818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				v819 = v816
				v820 = v818
			} else {
				v819 = v23
				v820 = v20
			}
			if v819 < base.I32_extend16_s(v820) {
				v828 = v819
				for {
					v837 = v19 + int32(base.Ui32(v828)>>(uint(int32(3))%32))
					v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837))))
					v842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v828<<(uint(int32(1))%32)))))
					v847 = v842 & int32(1023)
					v851 = v842 << (uint(int32(16)) % 32) & int32(-2147483648)
					v854 = int32(31)
					v855 = int32(base.Ui32(v842)>>(uint(int32(10))%32)) & v854
					if v855 != v854 {
						if v855 != 0 {
							v927 = v847
							v928 = v855<<(uint(int32(23))%32) + v851 + int32(939524096)
						} else {
							if v847 != 0 {
								if v842&int32(512) != 0 {
									v915 = v847 << (uint(int32(1)) % 32)
									v917 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v847) {
										v915 = v847 << (uint(int32(2)) % 32)
										v917 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v847) {
											v915 = v847 << (uint(int32(3)) % 32)
											v917 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v847) {
												v915 = v847 << (uint(int32(4)) % 32)
												v917 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v847) {
													v915 = v847 << (uint(int32(5)) % 32)
													v917 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v847) {
														v915 = v847 << (uint(int32(6)) % 32)
														v917 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v847) {
															v915 = v847 << (uint(int32(7)) % 32)
															v917 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v847) {
																v915 = v847 << (uint(int32(8)) % 32)
																v917 = int32(880803840)
															} else {
																v910 = base.B2i32(v847 == int32(1))
																if v847 == int32(1) {
																	v911 = int32(1024)
																} else {
																	v911 = v847 << (uint(int32(9)) % 32)
																}
																if v847 == int32(1) {
																	v914 = int32(864026624)
																} else {
																	v914 = int32(872415232)
																}
																v915 = v911
																v917 = v914
															}
														}
													}
												}
											}
										}
									}
								}
								v927 = v915 & int32(1022)
								v928 = v917 | v851
							} else {
								v927 = int32(0)
								v928 = v851
							}
						}
					} else {
						if v847 == int32(0) {
							v927 = int32(0)
							v928 = v851 | int32(2139095040)
						} else {
							v927 = v847
							v928 = v851 | int32(2143289344)
						}
					}
					v940 = v838 | base.F32_gt(base.F32_reinterpret_i32(v928|v927<<(uint(int32(13))%32)), float32(0))<<(uint((v828^int32(-1))&int32(7))%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v837))) = uint8(v940)
					v943 = v828 + int32(1)
					v944 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
					if v943 < v944 {
						v828 = v943
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			return v16
		}
	}
}
func F_halfvec_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 float32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 float32
	_ = v227
	var v233 int32
	_ = v233
	var v259 int32
	_ = v259
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v22 = base.B2i32(v20 < v21)
	if v20 < v21 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v259
L5:
	;
	if v20 < v21 {
		goto L103
	} else {
		goto L104
	}
L6:
	;
	v23 = v20
	goto L8
L7:
	;
	v23 = v21
	goto L8
L8:
	;
	if v23 <= int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v26 = int32(8)
	v31 = int32(0)
	goto L10
L10:
	;
	v43 = v31 << (uint(int32(1)) % 32)
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v26+v43))))
	v50 = v45 & int32(1023)
	v54 = v45 << (uint(int32(16)) % 32) & int32(-2147483648)
	v57 = int32(31)
	v58 = int32(base.Ui32(v45)>>(uint(int32(10))%32)) & v57
	if v58 != v57 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	return int32(1)
L12:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+v26+v43))))
	v142 = v137 & int32(1023)
	v146 = v137 << (uint(int32(16)) % 32) & int32(-2147483648)
	v149 = int32(31)
	v150 = int32(base.Ui32(v137)>>(uint(int32(10))%32)) & v149
	if v150 != v149 {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v135 = base.F32_reinterpret_i32(v131 | v130<<(uint(int32(13))%32))
	goto L12
L14:
	;
	v130 = v50
	v131 = v58<<(uint(int32(23))%32) + v54 + int32(939524096)
	goto L13
L15:
	;
	if v45&int32(512) != 0 {
		goto L25
	} else {
		goto L26
	}
L16:
	;
	if v58 != 0 {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v50 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v50 != 0 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v130 = int32(0)
	v131 = v54
	goto L13
L21:
	;
	v130 = int32(0)
	v131 = v54 | int32(2139095040)
	goto L13
L22:
	;
	goto L23
L23:
	;
	v130 = v50
	v131 = v54 | int32(2143289344)
	goto L13
L24:
	;
	v130 = v118 & int32(1022)
	v131 = v120 | v54
	goto L13
L25:
	;
	v118 = v50 << (uint(int32(1)) % 32)
	v120 = int32(939524096)
	goto L24
L26:
	;
	goto L27
L27:
	;
	if base.Ui32(int32(255)) < base.Ui32(v50) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v118 = v50 << (uint(int32(2)) % 32)
	v120 = int32(931135488)
	goto L24
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(127)) < base.Ui32(v50) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v118 = v50 << (uint(int32(3)) % 32)
	v120 = int32(922746880)
	goto L24
L32:
	;
	goto L33
L33:
	;
	if base.Ui32(int32(63)) < base.Ui32(v50) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v118 = v50 << (uint(int32(4)) % 32)
	v120 = int32(914358272)
	goto L24
L35:
	;
	goto L36
L36:
	;
	if base.Ui32(int32(31)) < base.Ui32(v50) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v118 = v50 << (uint(int32(5)) % 32)
	v120 = int32(905969664)
	goto L24
L38:
	;
	goto L39
L39:
	;
	if base.Ui32(int32(15)) < base.Ui32(v50) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v118 = v50 << (uint(int32(6)) % 32)
	v120 = int32(897581056)
	goto L24
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(7)) < base.Ui32(v50) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v118 = v50 << (uint(int32(7)) % 32)
	v120 = int32(889192448)
	goto L24
L44:
	;
	goto L45
L45:
	;
	if base.Ui32(int32(3)) < base.Ui32(v50) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v118 = v50 << (uint(int32(8)) % 32)
	v120 = int32(880803840)
	goto L24
L47:
	;
	goto L48
L48:
	;
	v113 = base.B2i32(v50 == int32(1))
	if v50 == int32(1) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v114 = int32(1024)
	goto L51
L50:
	;
	v114 = v50 << (uint(int32(9)) % 32)
	goto L51
L51:
	;
	if v50 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v117 = int32(864026624)
	goto L54
L53:
	;
	v117 = int32(872415232)
	goto L54
L54:
	;
	v118 = v114
	v120 = v117
	goto L24
L55:
	;
	if base.F32_lt(v135, v227) != 0 {
		v259 = int32(0)
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v227 = base.F32_reinterpret_i32(v223 | v222<<(uint(int32(13))%32))
	goto L55
L57:
	;
	v222 = v142
	v223 = v150<<(uint(int32(23))%32) + v146 + int32(939524096)
	goto L56
L58:
	;
	if v137&int32(512) != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	if v150 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v142 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v142 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v222 = int32(0)
	v223 = v146
	goto L56
L64:
	;
	v222 = int32(0)
	v223 = v146 | int32(2139095040)
	goto L56
L65:
	;
	goto L66
L66:
	;
	v222 = v142
	v223 = v146 | int32(2143289344)
	goto L56
L67:
	;
	v222 = v210 & int32(1022)
	v223 = v212 | v146
	goto L56
L68:
	;
	v210 = v142 << (uint(int32(1)) % 32)
	v212 = int32(939524096)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(255)) < base.Ui32(v142) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v210 = v142 << (uint(int32(2)) % 32)
	v212 = int32(931135488)
	goto L67
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(127)) < base.Ui32(v142) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v210 = v142 << (uint(int32(3)) % 32)
	v212 = int32(922746880)
	goto L67
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(63)) < base.Ui32(v142) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v210 = v142 << (uint(int32(4)) % 32)
	v212 = int32(914358272)
	goto L67
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(31)) < base.Ui32(v142) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v210 = v142 << (uint(int32(5)) % 32)
	v212 = int32(905969664)
	goto L67
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v142) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v210 = v142 << (uint(int32(6)) % 32)
	v212 = int32(897581056)
	goto L67
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(7)) < base.Ui32(v142) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v210 = v142 << (uint(int32(7)) % 32)
	v212 = int32(889192448)
	goto L67
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(3)) < base.Ui32(v142) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v210 = v142 << (uint(int32(8)) % 32)
	v212 = int32(880803840)
	goto L67
L90:
	;
	goto L91
L91:
	;
	v205 = base.B2i32(v142 == int32(1))
	if v142 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v206 = int32(1024)
	goto L94
L93:
	;
	v206 = v142 << (uint(int32(9)) % 32)
	goto L94
L94:
	;
	if v142 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v209 = int32(864026624)
	goto L97
L96:
	;
	v209 = int32(872415232)
	goto L97
L97:
	;
	v210 = v206
	v212 = v209
	goto L67
L98:
	;
	if base.F32_gt(v135, v227) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v233 = v31 + int32(1)
	if v233 == v23 {
		goto L5
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L11
L102:
	;
	v31 = v233
	goto L10
L103:
	;
	return int32(0)
L104:
	;
	goto L105
L105:
	;
	v259 = base.B2i32(v21 < v20)
	goto L4
}
func F_halfvec_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 float32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 float32
	_ = v225
	var v233 int32
	_ = v233
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = base.B2i32(v19 < v20)
	if v19 < v20 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v21 | base.B2i32(v19 <= v20)
L5:
	;
	v22 = v19
	goto L7
L6:
	;
	v22 = v20
	goto L7
L7:
	;
	if v22 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v25 = int32(8)
	v30 = int32(0)
	goto L9
L9:
	;
	v41 = v30 << (uint(int32(1)) % 32)
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12+v25+v41))))
	v48 = v43 & int32(1023)
	v52 = v43 << (uint(int32(16)) % 32) & int32(-2147483648)
	v55 = int32(31)
	v56 = int32(base.Ui32(v43)>>(uint(int32(10))%32)) & v55
	if v56 != v55 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(0)
L11:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+v25+v41))))
	v140 = v135 & int32(1023)
	v144 = v135 << (uint(int32(16)) % 32) & int32(-2147483648)
	v147 = int32(31)
	v148 = int32(base.Ui32(v135)>>(uint(int32(10))%32)) & v147
	if v148 != v147 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v133 = base.F32_reinterpret_i32(v129 | v128<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v128 = v48
	v129 = v56<<(uint(int32(23))%32) + v52 + int32(939524096)
	goto L12
L14:
	;
	if v43&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v56 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v48 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v128 = int32(0)
	v129 = v52
	goto L12
L20:
	;
	v128 = int32(0)
	v129 = v52 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v128 = v48
	v129 = v52 | int32(2143289344)
	goto L12
L23:
	;
	v128 = v116 & int32(1022)
	v129 = v118 | v52
	goto L12
L24:
	;
	v116 = v48 << (uint(int32(1)) % 32)
	v118 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v48) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v48 << (uint(int32(2)) % 32)
	v118 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v48) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v116 = v48 << (uint(int32(3)) % 32)
	v118 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v48) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v116 = v48 << (uint(int32(4)) % 32)
	v118 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v48) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v116 = v48 << (uint(int32(5)) % 32)
	v118 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v48) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v116 = v48 << (uint(int32(6)) % 32)
	v118 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v48) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v116 = v48 << (uint(int32(7)) % 32)
	v118 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v48) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v116 = v48 << (uint(int32(8)) % 32)
	v118 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v111 = base.B2i32(v48 == int32(1))
	if v48 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v112 = int32(1024)
	goto L50
L49:
	;
	v112 = v48 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v48 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v115 = int32(864026624)
	goto L53
L52:
	;
	v115 = int32(872415232)
	goto L53
L53:
	;
	v116 = v112
	v118 = v115
	goto L23
L54:
	;
	if base.F32_lt(v133, v225) != 0 {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v225 = base.F32_reinterpret_i32(v221 | v220<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v220 = v140
	v221 = v148<<(uint(int32(23))%32) + v144 + int32(939524096)
	goto L55
L57:
	;
	if v135&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v148 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v140 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v140 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v220 = int32(0)
	v221 = v144
	goto L55
L63:
	;
	v220 = int32(0)
	v221 = v144 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v220 = v140
	v221 = v144 | int32(2143289344)
	goto L55
L66:
	;
	v220 = v208 & int32(1022)
	v221 = v210 | v144
	goto L55
L67:
	;
	v208 = v140 << (uint(int32(1)) % 32)
	v210 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v140) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v208 = v140 << (uint(int32(2)) % 32)
	v210 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v140) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v208 = v140 << (uint(int32(3)) % 32)
	v210 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v140) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v208 = v140 << (uint(int32(4)) % 32)
	v210 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v140) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v208 = v140 << (uint(int32(5)) % 32)
	v210 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v140) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v208 = v140 << (uint(int32(6)) % 32)
	v210 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v140) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v208 = v140 << (uint(int32(7)) % 32)
	v210 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v140) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v208 = v140 << (uint(int32(8)) % 32)
	v210 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v203 = base.B2i32(v140 == int32(1))
	if v140 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v204 = int32(1024)
	goto L93
L92:
	;
	v204 = v140 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v140 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v207 = int32(864026624)
	goto L96
L95:
	;
	v207 = int32(872415232)
	goto L96
L96:
	;
	v208 = v204
	v210 = v207
	goto L66
L97:
	;
	return int32(1)
L98:
	;
	goto L99
L99:
	;
	if base.F32_gt(v133, v225) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v233 = v30 + int32(1)
	if v233 == v22 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L10
L103:
	;
	v30 = v233
	goto L9
}
func F_halfvec_spherical_distance(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 float32
	_ = v30
	var v31 int32
	_ = v31
	var v40 float64
	_ = v40
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v62 float64
	_ = v62
	var v73 float64
	_ = v73
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v92 float64
	_ = v92
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v114 float64
	_ = v114
	var v117 float64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
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
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
			if v19 == v20 {
				v24 = int32(8)
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_halfvec_spherical_distance[0]))
				v30 = m.T0[v29].(func(*base.Module, int32, int32, int32) float32)(m, base.I32_extend16_s(v19), v12+v24, v17+v24)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					if base.F32_gt(v30, float32(1)) != 0 {
						v40 = float64(1)
					} else {
						if base.F32_lt(v30, float32(-1)) == int32(0) {
							v40 = base.F64_promote_f32(v30)
						} else {
							v40 = float64(-1)
						}
					}
					v44 = base.I64_reinterpret_f64(v40)
					v49 = base.I32_wrap_i64(int64(base.Ui64(v44)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v49) {
						if base.I32_wrap_i64(v44)|(v49-int32(1072693248)) == int32(0) {
							if int64(0) <= v44 {
								v62 = float64(0)
							} else {
								v62 = float64(3.141592653589793)
							}
							v117 = v62
						} else {
							v117 = base.F64_div(float64(0), base.F64_sub(v40, v40))
						}
					} else {
						if base.Ui32(v49) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v49) < base.Ui32(int32(1012924417)) {
								v114 = float64(1.5707963267948966)
								v117 = v114
							} else {
								v73 = F_R(m, base.F64_mul(v40, v40))
								mBase = m.M
								v117 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v40, v73)), v40), float64(1.5707963267948966))
							}
						} else {
							if v44 < int64(0) {
								v85 = base.F64_mul(base.F64_add(v40, float64(1)), float64(0.5))
								v86 = base.F64_sqrt(v85)
								v87 = F_R(m, v85)
								mBase = m.M
								v92 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v86, base.F64_add(base.F64_mul(v86, v87), float64(-6.123233995736766e-17))))
								v117 = base.F64_add(v92, v92)
							} else {
								v97 = base.F64_mul(base.F64_sub(float64(1), v40), float64(0.5))
								v98 = base.F64_sqrt(v97)
								v99 = F_R(m, v97)
								mBase = m.M
								v104 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v98) & int64(-4294967296))
								v110 = base.F64_add(base.F64_add(base.F64_mul(v98, v99), base.F64_div(base.F64_sub(v97, base.F64_mul(v104, v104)), base.F64_add(v98, v104))), v104)
								v114 = base.F64_add(v110, v110)
								v117 = v114
							}
						}
					}
					v120 = F_Float8GetDatum(m, base.F64_div(v117, float64(3.141592653589793)))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v120
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v129 = m.ExcPending
				if v129 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return int32(0)
					} else {
						v133 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v133
						F_errmsg(m, int32(_a_F_halfvec_spherical_distance_0), v9)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_spherical_distance_1), int32(80), int32(_a_F_halfvec_spherical_distance_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
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
