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
	var v39 float32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 float32
	_ = v48
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
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
					v39 = *(*float32)(unsafe.Add(mBase, uint32(l2+v26<<(uint(int32(2))%32))))
					v40 = base.I32_reinterpret_f32(v39)
					v42 = int32(base.Ui32(v40) >> (uint(int32(16)) % 32))
					v48 = base.F32_abs(v39)
					if base.F32_eq(v48, math.Float32frombits(uint32(0x7f800000))) != 0 {
						v127 = v42 & int32(_a_F_HalfvecUpdateCenter_0)
					} else {
						v52 = v40 & int32(_a_F_HalfvecUpdateCenter_1)
						if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v48)) {
							v127 = v42&int32(_a_F_HalfvecUpdateCenter_2) | int32(base.Ui32(v52)>>(uint(int32(13))%32)) | int32(_a_F_HalfvecUpdateCenter_3)
						} else {
							v64 = v42 & int32(_a_F_HalfvecUpdateCenter_2)
							v68 = int32(base.Ui32(v40)>>(uint(int32(23))%32)) & int32(255)
							if base.Ui32(v68) < base.Ui32(int32(99)) {
								v127 = v64
							} else {
								if base.Ui32(v68) <= base.Ui32(int32(112)) {
									v80 = int32(1)<<(uint(v68-int32(90))%32) + int32(base.Ui32(v52)>>(uint(int32(113)-v68)%32))
									v82 = v80 | v40
									v83 = v80
								} else {
									v82 = v40
									v83 = v52
								}
								v85 = int32(base.Ui32(v83) >> (uint(int32(13)) % 32))
								v88 = int32(3)
								v89 = int32(base.Ui32(v83)>>(uint(int32(12))%32)) & v88
								if v89 != v88 {
									if v89 != int32(1) {
										v100 = v85
									} else {
										if v82&int32(4095) == int32(0) {
											v100 = v85
										} else {
											v100 = v85 + int32(1)
										}
									}
								} else {
									v100 = v85 + int32(1)
								}
								v106 = base.B2i32(v100 == int32(1024))
								if v100 == int32(1024) {
									v107 = int32(-126)
								} else {
									v107 = int32(-127)
								}
								v108 = v107 + v68
								if int32(16) <= v108 {
									v127 = v64 | int32(_a_F_HalfvecUpdateCenter_4)
								} else {
									if int32(-15) < v108 {
										v118 = v108<<(uint(int32(10))%32) + int32(_a_F_HalfvecUpdateCenter_5) | v64
									} else {
										v118 = v64
									}
									if v100 == int32(1024) {
										v120 = int32(0)
									} else {
										v120 = v100
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
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
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
		if v14 == int32(-1) {
			m.G0 = v7 + int32(16)
			return v10
		} else {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+4)))
			if v14 == v17 {
				m.G0 = v7 + int32(16)
				return v10
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v14
						F_errmsg(m, int32(_a_F_halfvec_0), v7)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_1), int32(92), int32(_a_F_halfvec_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 float32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 float32
	_ = v262
	var v266 int32
	_ = v266
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
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
	v365 = int32(0)
	v366 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v366 <= v365 {
		goto L4
	} else {
		goto L129
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
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L125
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
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68+(v25+v47)))))
	v168 = v163 & int32(1023)
	v172 = v163 << (uint(int32(16)) % 32) & int32(-2147483648)
	v175 = int32(31)
	v176 = int32(base.Ui32(v163)>>(uint(int32(10))%32)) & v175
	if v176 != v175 {
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
	v255 = base.F32_add(base.F32_reinterpret_i32(v156|v155<<(uint(int32(13))%32)), base.F32_reinterpret_i32(v249|v248<<(uint(int32(13))%32)))
	v256 = base.I32_reinterpret_f32(v255)
	v258 = int32(base.Ui32(v256) >> (uint(int32(16)) % 32))
	v262 = base.F32_abs(v255)
	if base.F32_eq(v262, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v341 = v258 & int32(_a_F_halfvec_add_0)
		goto L101
	} else {
		goto L102
	}
L59:
	;
	goto L58
L60:
	;
	v248 = v168
	v249 = v176<<(uint(int32(23))%32) + v172 + int32(939524096)
	goto L59
L61:
	;
	if v163&int32(512) != 0 {
		goto L71
	} else {
		goto L72
	}
L62:
	;
	if v176 != 0 {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	if v168 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	if v168 != 0 {
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v248 = int32(0)
	v249 = v172
	goto L59
L67:
	;
	v248 = int32(0)
	v249 = v172 | int32(2139095040)
	goto L59
L68:
	;
	goto L69
L69:
	;
	v248 = v168
	v249 = v172 | int32(2143289344)
	goto L59
L70:
	;
	v248 = v236 & int32(1022)
	v249 = v238 | v172
	goto L59
L71:
	;
	v236 = v168 << (uint(int32(1)) % 32)
	v238 = int32(939524096)
	goto L70
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(255)) < base.Ui32(v168) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v236 = v168 << (uint(int32(2)) % 32)
	v238 = int32(931135488)
	goto L70
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(127)) < base.Ui32(v168) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v236 = v168 << (uint(int32(3)) % 32)
	v238 = int32(922746880)
	goto L70
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(63)) < base.Ui32(v168) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v236 = v168 << (uint(int32(4)) % 32)
	v238 = int32(914358272)
	goto L70
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(31)) < base.Ui32(v168) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v236 = v168 << (uint(int32(5)) % 32)
	v238 = int32(905969664)
	goto L70
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(15)) < base.Ui32(v168) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v236 = v168 << (uint(int32(6)) % 32)
	v238 = int32(897581056)
	goto L70
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(7)) < base.Ui32(v168) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v236 = v168 << (uint(int32(7)) % 32)
	v238 = int32(889192448)
	goto L70
L90:
	;
	goto L91
L91:
	;
	if base.Ui32(int32(3)) < base.Ui32(v168) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v236 = v168 << (uint(int32(8)) % 32)
	v238 = int32(880803840)
	goto L70
L93:
	;
	goto L94
L94:
	;
	v231 = base.B2i32(v168 == int32(1))
	if v168 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v232 = int32(1024)
	goto L97
L96:
	;
	v232 = v168 << (uint(int32(9)) % 32)
	goto L97
L97:
	;
	if v168 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v235 = int32(864026624)
	goto L100
L99:
	;
	v235 = int32(872415232)
	goto L100
L100:
	;
	v236 = v232
	v238 = v235
	goto L70
L101:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v68+v52))) = uint16(v341)
	v344 = v53 + int32(1)
	if v44 != v344 {
		v53 = v344
		goto L13
	} else {
		goto L124
	}
L102:
	;
	v266 = v256 & int32(_a_F_halfvec_add_1)
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v262)) {
		v341 = v258&int32(_a_F_halfvec_add_2) | int32(base.Ui32(v266)>>(uint(int32(13))%32)) | int32(_a_F_halfvec_add_3)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v278 = v258 & int32(_a_F_halfvec_add_2)
	v282 = int32(base.Ui32(v256)>>(uint(int32(23))%32)) & int32(255)
	if base.Ui32(v282) < base.Ui32(int32(99)) {
		v341 = v278
		goto L101
	} else {
		goto L104
	}
L104:
	;
	if base.Ui32(v282) <= base.Ui32(int32(112)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v294 = int32(1)<<(uint(v282-int32(90))%32) + int32(base.Ui32(v266)>>(uint(int32(113)-v282)%32))
	v296 = v294
	v297 = v294 | v256
	goto L107
L106:
	;
	v296 = v266
	v297 = v256
	goto L107
L107:
	;
	v299 = int32(base.Ui32(v296) >> (uint(int32(13)) % 32))
	v302 = int32(3)
	v303 = int32(base.Ui32(v296)>>(uint(int32(12))%32)) & v302
	if v303 != v302 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v320 = base.B2i32(v314 == int32(1024))
	if v314 == int32(1024) {
		goto L114
	} else {
		goto L115
	}
L109:
	;
	if v303 != int32(1) {
		v314 = v299
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v314 = v299 + int32(1)
	goto L108
L112:
	;
	if v297&int32(4095) == int32(0) {
		v314 = v299
		goto L108
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v321 = int32(-126)
	goto L116
L115:
	;
	v321 = int32(-127)
	goto L116
L116:
	;
	v322 = v321 + v282
	if int32(16) <= v322 {
		v341 = v278 | int32(_a_F_halfvec_add_4)
		goto L101
	} else {
		goto L117
	}
L117:
	;
	if int32(-15) < v322 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v332 = v322<<(uint(int32(10))%32) + int32(_a_F_halfvec_add_5) | v278
	goto L120
L119:
	;
	v332 = v278
	goto L120
L120:
	;
	if v314 == int32(1024) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v334 = int32(0)
	goto L123
L122:
	;
	v334 = v314
	goto L123
L123:
	;
	v341 = v332 | v334
	goto L101
L124:
	;
	goto L14
L125:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+4)))
	v354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v354
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v353
	F_errmsg(m, int32(_a_F_halfvec_add_6), v17)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_halfvec_add_7), int32(80), int32(_a_F_halfvec_add_8))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v369 = v365
	goto L130
L130:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52+v369<<(uint(int32(1))%32)))))
	if v386&int32(_a_F_halfvec_add_9) != int32(_a_F_halfvec_add_4) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L136
	}
L132:
	;
	v392 = v369 + int32(1)
	if v366 != v392 {
		v369 = v392
		goto L130
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	goto L131
L135:
	;
	goto L4
L136:
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
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
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v843 int32
	_ = v843
	var v848 int32
	_ = v848
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
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
					v41 = v14 + v29<<(uint(int32(1))%32)
					v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
					v47 = v42 & int32(1023)
					v51 = v42 << (uint(int32(16)) % 32) & int32(-2147483648)
					v54 = int32(31)
					v55 = int32(base.Ui32(v42)>>(uint(int32(10))%32)) & v54
					if v55 != v54 {
						if v55 != 0 {
							v127 = v47
							v128 = v55<<(uint(int32(23))%32) + v51 + int32(939524096)
						} else {
							if v47 != 0 {
								if v42&int32(512) != 0 {
									v115 = v47 << (uint(int32(1)) % 32)
									v117 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v47) {
										v115 = v47 << (uint(int32(2)) % 32)
										v117 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v47) {
											v115 = v47 << (uint(int32(3)) % 32)
											v117 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v47) {
												v115 = v47 << (uint(int32(4)) % 32)
												v117 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v47) {
													v115 = v47 << (uint(int32(5)) % 32)
													v117 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v47) {
														v115 = v47 << (uint(int32(6)) % 32)
														v117 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v47) {
															v115 = v47 << (uint(int32(7)) % 32)
															v117 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v47) {
																v115 = v47 << (uint(int32(8)) % 32)
																v117 = int32(880803840)
															} else {
																v110 = base.B2i32(v47 == int32(1))
																if v47 == int32(1) {
																	v111 = int32(1024)
																} else {
																	v111 = v47 << (uint(int32(9)) % 32)
																}
																if v47 == int32(1) {
																	v114 = int32(864026624)
																} else {
																	v114 = int32(872415232)
																}
																v115 = v111
																v117 = v114
															}
														}
													}
												}
											}
										}
									}
								}
								v127 = v115 & int32(1022)
								v128 = v117 | v51
							} else {
								v127 = int32(0)
								v128 = v51
							}
						}
					} else {
						if v47 == int32(0) {
							v127 = int32(0)
							v128 = v51 | int32(2139095040)
						} else {
							v127 = v47
							v128 = v51 | int32(2143289344)
						}
					}
					v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
					v143 = v138 & int32(1023)
					v147 = v138 << (uint(int32(16)) % 32) & int32(-2147483648)
					v150 = int32(31)
					v151 = int32(base.Ui32(v138)>>(uint(int32(10))%32)) & v150
					if v151 != v150 {
						if v151 != 0 {
							v223 = v143
							v224 = v151<<(uint(int32(23))%32) + v147 + int32(939524096)
						} else {
							if v143 != 0 {
								if v138&int32(512) != 0 {
									v211 = v143 << (uint(int32(1)) % 32)
									v213 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v143) {
										v211 = v143 << (uint(int32(2)) % 32)
										v213 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v143) {
											v211 = v143 << (uint(int32(3)) % 32)
											v213 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v143) {
												v211 = v143 << (uint(int32(4)) % 32)
												v213 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v143) {
													v211 = v143 << (uint(int32(5)) % 32)
													v213 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v143) {
														v211 = v143 << (uint(int32(6)) % 32)
														v213 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v143) {
															v211 = v143 << (uint(int32(7)) % 32)
															v213 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v143) {
																v211 = v143 << (uint(int32(8)) % 32)
																v213 = int32(880803840)
															} else {
																v206 = base.B2i32(v143 == int32(1))
																if v143 == int32(1) {
																	v207 = int32(1024)
																} else {
																	v207 = v143 << (uint(int32(9)) % 32)
																}
																if v143 == int32(1) {
																	v210 = int32(864026624)
																} else {
																	v210 = int32(872415232)
																}
																v211 = v207
																v213 = v210
															}
														}
													}
												}
											}
										}
									}
								}
								v223 = v211 & int32(1022)
								v224 = v213 | v147
							} else {
								v223 = int32(0)
								v224 = v147
							}
						}
					} else {
						if v143 == int32(0) {
							v223 = int32(0)
							v224 = v147 | int32(2139095040)
						} else {
							v223 = v143
							v224 = v147 | int32(2143289344)
						}
					}
					v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
					v240 = v235 & int32(1023)
					v244 = v235 << (uint(int32(16)) % 32) & int32(-2147483648)
					v247 = int32(31)
					v248 = int32(base.Ui32(v235)>>(uint(int32(10))%32)) & v247
					if v248 != v247 {
						if v248 != 0 {
							v320 = v240
							v321 = v248<<(uint(int32(23))%32) + v244 + int32(939524096)
						} else {
							if v240 != 0 {
								if v235&int32(512) != 0 {
									v308 = v240 << (uint(int32(1)) % 32)
									v310 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v240) {
										v308 = v240 << (uint(int32(2)) % 32)
										v310 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v240) {
											v308 = v240 << (uint(int32(3)) % 32)
											v310 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v240) {
												v308 = v240 << (uint(int32(4)) % 32)
												v310 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v240) {
													v308 = v240 << (uint(int32(5)) % 32)
													v310 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v240) {
														v308 = v240 << (uint(int32(6)) % 32)
														v310 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v240) {
															v308 = v240 << (uint(int32(7)) % 32)
															v310 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v240) {
																v308 = v240 << (uint(int32(8)) % 32)
																v310 = int32(880803840)
															} else {
																v303 = base.B2i32(v240 == int32(1))
																if v240 == int32(1) {
																	v304 = int32(1024)
																} else {
																	v304 = v240 << (uint(int32(9)) % 32)
																}
																if v240 == int32(1) {
																	v307 = int32(864026624)
																} else {
																	v307 = int32(872415232)
																}
																v308 = v304
																v310 = v307
															}
														}
													}
												}
											}
										}
									}
								}
								v320 = v308 & int32(1022)
								v321 = v310 | v244
							} else {
								v320 = int32(0)
								v321 = v244
							}
						}
					} else {
						if v240 == int32(0) {
							v320 = int32(0)
							v321 = v244 | int32(2139095040)
						} else {
							v320 = v240
							v321 = v244 | int32(2143289344)
						}
					}
					v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)))
					v337 = v332 & int32(1023)
					v341 = v332 << (uint(int32(16)) % 32) & int32(-2147483648)
					v344 = int32(31)
					v345 = int32(base.Ui32(v332)>>(uint(int32(10))%32)) & v344
					if v345 != v344 {
						if v345 != 0 {
							v417 = v337
							v418 = v345<<(uint(int32(23))%32) + v341 + int32(939524096)
						} else {
							if v337 != 0 {
								if v332&int32(512) != 0 {
									v405 = v337 << (uint(int32(1)) % 32)
									v407 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v337) {
										v405 = v337 << (uint(int32(2)) % 32)
										v407 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v337) {
											v405 = v337 << (uint(int32(3)) % 32)
											v407 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v337) {
												v405 = v337 << (uint(int32(4)) % 32)
												v407 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v337) {
													v405 = v337 << (uint(int32(5)) % 32)
													v407 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v337) {
														v405 = v337 << (uint(int32(6)) % 32)
														v407 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v337) {
															v405 = v337 << (uint(int32(7)) % 32)
															v407 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v337) {
																v405 = v337 << (uint(int32(8)) % 32)
																v407 = int32(880803840)
															} else {
																v400 = base.B2i32(v337 == int32(1))
																if v337 == int32(1) {
																	v401 = int32(1024)
																} else {
																	v401 = v337 << (uint(int32(9)) % 32)
																}
																if v337 == int32(1) {
																	v404 = int32(864026624)
																} else {
																	v404 = int32(872415232)
																}
																v405 = v401
																v407 = v404
															}
														}
													}
												}
											}
										}
									}
								}
								v417 = v405 & int32(1022)
								v418 = v407 | v341
							} else {
								v417 = int32(0)
								v418 = v341
							}
						}
					} else {
						if v337 == int32(0) {
							v417 = int32(0)
							v418 = v341 | int32(2139095040)
						} else {
							v417 = v337
							v418 = v341 | int32(2143289344)
						}
					}
					v429 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+8)))
					v434 = v429 & int32(1023)
					v438 = v429 << (uint(int32(16)) % 32) & int32(-2147483648)
					v441 = int32(31)
					v442 = int32(base.Ui32(v429)>>(uint(int32(10))%32)) & v441
					if v442 != v441 {
						if v442 != 0 {
							v514 = v434
							v515 = v442<<(uint(int32(23))%32) + v438 + int32(939524096)
						} else {
							if v434 != 0 {
								if v429&int32(512) != 0 {
									v502 = v434 << (uint(int32(1)) % 32)
									v504 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v434) {
										v502 = v434 << (uint(int32(2)) % 32)
										v504 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v434) {
											v502 = v434 << (uint(int32(3)) % 32)
											v504 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v434) {
												v502 = v434 << (uint(int32(4)) % 32)
												v504 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v434) {
													v502 = v434 << (uint(int32(5)) % 32)
													v504 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v434) {
														v502 = v434 << (uint(int32(6)) % 32)
														v504 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v434) {
															v502 = v434 << (uint(int32(7)) % 32)
															v504 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v434) {
																v502 = v434 << (uint(int32(8)) % 32)
																v504 = int32(880803840)
															} else {
																v497 = base.B2i32(v434 == int32(1))
																if v434 == int32(1) {
																	v498 = int32(1024)
																} else {
																	v498 = v434 << (uint(int32(9)) % 32)
																}
																if v434 == int32(1) {
																	v501 = int32(864026624)
																} else {
																	v501 = int32(872415232)
																}
																v502 = v498
																v504 = v501
															}
														}
													}
												}
											}
										}
									}
								}
								v514 = v502 & int32(1022)
								v515 = v504 | v438
							} else {
								v514 = int32(0)
								v515 = v438
							}
						}
					} else {
						if v434 == int32(0) {
							v514 = int32(0)
							v515 = v438 | int32(2139095040)
						} else {
							v514 = v434
							v515 = v438 | int32(2143289344)
						}
					}
					v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+10)))
					v531 = v526 & int32(1023)
					v535 = v526 << (uint(int32(16)) % 32) & int32(-2147483648)
					v538 = int32(31)
					v539 = int32(base.Ui32(v526)>>(uint(int32(10))%32)) & v538
					if v539 != v538 {
						if v539 != 0 {
							v611 = v531
							v612 = v539<<(uint(int32(23))%32) + v535 + int32(939524096)
						} else {
							if v531 != 0 {
								if v526&int32(512) != 0 {
									v599 = v531 << (uint(int32(1)) % 32)
									v601 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v531) {
										v599 = v531 << (uint(int32(2)) % 32)
										v601 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v531) {
											v599 = v531 << (uint(int32(3)) % 32)
											v601 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v531) {
												v599 = v531 << (uint(int32(4)) % 32)
												v601 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v531) {
													v599 = v531 << (uint(int32(5)) % 32)
													v601 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v531) {
														v599 = v531 << (uint(int32(6)) % 32)
														v601 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v531) {
															v599 = v531 << (uint(int32(7)) % 32)
															v601 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v531) {
																v599 = v531 << (uint(int32(8)) % 32)
																v601 = int32(880803840)
															} else {
																v594 = base.B2i32(v531 == int32(1))
																if v531 == int32(1) {
																	v595 = int32(1024)
																} else {
																	v595 = v531 << (uint(int32(9)) % 32)
																}
																if v531 == int32(1) {
																	v598 = int32(864026624)
																} else {
																	v598 = int32(872415232)
																}
																v599 = v595
																v601 = v598
															}
														}
													}
												}
											}
										}
									}
								}
								v611 = v599 & int32(1022)
								v612 = v601 | v535
							} else {
								v611 = int32(0)
								v612 = v535
							}
						}
					} else {
						if v531 == int32(0) {
							v611 = int32(0)
							v612 = v535 | int32(2139095040)
						} else {
							v611 = v531
							v612 = v535 | int32(2143289344)
						}
					}
					v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+12)))
					v628 = v623 & int32(1023)
					v632 = v623 << (uint(int32(16)) % 32) & int32(-2147483648)
					v635 = int32(31)
					v636 = int32(base.Ui32(v623)>>(uint(int32(10))%32)) & v635
					if v636 != v635 {
						if v636 != 0 {
							v708 = v628
							v709 = v636<<(uint(int32(23))%32) + v632 + int32(939524096)
						} else {
							if v628 != 0 {
								if v623&int32(512) != 0 {
									v696 = v628 << (uint(int32(1)) % 32)
									v698 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v628) {
										v696 = v628 << (uint(int32(2)) % 32)
										v698 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v628) {
											v696 = v628 << (uint(int32(3)) % 32)
											v698 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v628) {
												v696 = v628 << (uint(int32(4)) % 32)
												v698 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v628) {
													v696 = v628 << (uint(int32(5)) % 32)
													v698 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v628) {
														v696 = v628 << (uint(int32(6)) % 32)
														v698 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v628) {
															v696 = v628 << (uint(int32(7)) % 32)
															v698 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v628) {
																v696 = v628 << (uint(int32(8)) % 32)
																v698 = int32(880803840)
															} else {
																v691 = base.B2i32(v628 == int32(1))
																if v628 == int32(1) {
																	v692 = int32(1024)
																} else {
																	v692 = v628 << (uint(int32(9)) % 32)
																}
																if v628 == int32(1) {
																	v695 = int32(864026624)
																} else {
																	v695 = int32(872415232)
																}
																v696 = v692
																v698 = v695
															}
														}
													}
												}
											}
										}
									}
								}
								v708 = v696 & int32(1022)
								v709 = v698 | v632
							} else {
								v708 = int32(0)
								v709 = v632
							}
						}
					} else {
						if v628 == int32(0) {
							v708 = int32(0)
							v709 = v632 | int32(2139095040)
						} else {
							v708 = v628
							v709 = v632 | int32(2143289344)
						}
					}
					v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+14)))
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
					v814 = base.F32_gt(base.F32_reinterpret_i32(v128|v127<<(uint(int32(13))%32)), float32(0))<<(uint(int32(7))%32) | base.F32_gt(base.F32_reinterpret_i32(v224|v223<<(uint(int32(13))%32)), float32(0))<<(uint(int32(6))%32) | base.F32_gt(base.F32_reinterpret_i32(v321|v320<<(uint(int32(13))%32)), float32(0))<<(uint(int32(5))%32) | base.F32_gt(base.F32_reinterpret_i32(v418|v417<<(uint(int32(13))%32)), float32(0))<<(uint(int32(4))%32) | base.F32_gt(base.F32_reinterpret_i32(v515|v514<<(uint(int32(13))%32)), float32(0))<<(uint(int32(3))%32) | base.F32_gt(base.F32_reinterpret_i32(v612|v611<<(uint(int32(13))%32)), float32(0))<<(uint(int32(2))%32) | base.F32_gt(base.F32_reinterpret_i32(v709|v708<<(uint(int32(13))%32)), float32(0))<<(uint(int32(1))%32) | base.F32_gt(base.F32_reinterpret_i32(v806|v805<<(uint(int32(13))%32)), float32(0))
					*(*uint8)(unsafe.Add(mBase, uint32(v19+int32(base.Ui32(v29)>>(uint(int32(3))%32))))) = uint8(v814)
					v817 = v29 + int32(8)
					if v817 < base.I32_extend16_s(v22)<<(uint(int32(3))%32) {
						v29 = v817
						continue
					} else {
						break
					}
					break
				}
				v819 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
				v820 = v817
				v821 = v819
			} else {
				v820 = v23
				v821 = v20
			}
			if v820 < base.I32_extend16_s(v821) {
				v829 = v820
				for {
					v838 = v19 + int32(base.Ui32(v829)>>(uint(int32(3))%32))
					v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838))))
					v843 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14+v829<<(uint(int32(1))%32)))))
					v848 = v843 & int32(1023)
					v852 = v843 << (uint(int32(16)) % 32) & int32(-2147483648)
					v855 = int32(31)
					v856 = int32(base.Ui32(v843)>>(uint(int32(10))%32)) & v855
					if v856 != v855 {
						if v856 != 0 {
							v928 = v848
							v929 = v856<<(uint(int32(23))%32) + v852 + int32(939524096)
						} else {
							if v848 != 0 {
								if v843&int32(512) != 0 {
									v916 = v848 << (uint(int32(1)) % 32)
									v918 = int32(939524096)
								} else {
									if base.Ui32(int32(255)) < base.Ui32(v848) {
										v916 = v848 << (uint(int32(2)) % 32)
										v918 = int32(931135488)
									} else {
										if base.Ui32(int32(127)) < base.Ui32(v848) {
											v916 = v848 << (uint(int32(3)) % 32)
											v918 = int32(922746880)
										} else {
											if base.Ui32(int32(63)) < base.Ui32(v848) {
												v916 = v848 << (uint(int32(4)) % 32)
												v918 = int32(914358272)
											} else {
												if base.Ui32(int32(31)) < base.Ui32(v848) {
													v916 = v848 << (uint(int32(5)) % 32)
													v918 = int32(905969664)
												} else {
													if base.Ui32(int32(15)) < base.Ui32(v848) {
														v916 = v848 << (uint(int32(6)) % 32)
														v918 = int32(897581056)
													} else {
														if base.Ui32(int32(7)) < base.Ui32(v848) {
															v916 = v848 << (uint(int32(7)) % 32)
															v918 = int32(889192448)
														} else {
															if base.Ui32(int32(3)) < base.Ui32(v848) {
																v916 = v848 << (uint(int32(8)) % 32)
																v918 = int32(880803840)
															} else {
																v911 = base.B2i32(v848 == int32(1))
																if v848 == int32(1) {
																	v912 = int32(1024)
																} else {
																	v912 = v848 << (uint(int32(9)) % 32)
																}
																if v848 == int32(1) {
																	v915 = int32(864026624)
																} else {
																	v915 = int32(872415232)
																}
																v916 = v912
																v918 = v915
															}
														}
													}
												}
											}
										}
									}
								}
								v928 = v916 & int32(1022)
								v929 = v918 | v852
							} else {
								v928 = int32(0)
								v929 = v852
							}
						}
					} else {
						if v848 == int32(0) {
							v928 = int32(0)
							v929 = v852 | int32(2139095040)
						} else {
							v928 = v848
							v929 = v852 | int32(2143289344)
						}
					}
					v942 = v839 | base.F32_gt(base.F32_reinterpret_i32(v929|v928<<(uint(int32(13))%32)), float32(0))<<(uint((v829^int32(-1))&int32(7))%32)
					*(*uint8)(unsafe.Add(mBase, uint32(v838))) = uint8(v942)
					v945 = v829 + int32(1)
					v946 = int32(*(*int16)(unsafe.Add(mBase, uint32(v9)+4)))
					if v945 < v946 {
						v829 = v945
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
	var v136 float32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 float32
	_ = v229
	var v235 int32
	_ = v235
	var v257 int32
	_ = v257
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
	return v257
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
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43+(v18+v26)))))
	v143 = v138 & int32(1023)
	v147 = v138 << (uint(int32(16)) % 32) & int32(-2147483648)
	v150 = int32(31)
	v151 = int32(base.Ui32(v138)>>(uint(int32(10))%32)) & v150
	if v151 != v150 {
		goto L59
	} else {
		goto L60
	}
L13:
	;
	v136 = base.F32_reinterpret_i32(v131 | v130<<(uint(int32(13))%32))
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
	if base.F32_lt(v136, v229) != 0 {
		v257 = int32(0)
		goto L4
	} else {
		goto L98
	}
L56:
	;
	v229 = base.F32_reinterpret_i32(v224 | v223<<(uint(int32(13))%32))
	goto L55
L57:
	;
	v223 = v143
	v224 = v151<<(uint(int32(23))%32) + v147 + int32(939524096)
	goto L56
L58:
	;
	if v138&int32(512) != 0 {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	if v151 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v143 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v143 != 0 {
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v223 = int32(0)
	v224 = v147
	goto L56
L64:
	;
	v223 = int32(0)
	v224 = v147 | int32(2139095040)
	goto L56
L65:
	;
	goto L66
L66:
	;
	v223 = v143
	v224 = v147 | int32(2143289344)
	goto L56
L67:
	;
	v223 = v211 & int32(1022)
	v224 = v213 | v147
	goto L56
L68:
	;
	v211 = v143 << (uint(int32(1)) % 32)
	v213 = int32(939524096)
	goto L67
L69:
	;
	goto L70
L70:
	;
	if base.Ui32(int32(255)) < base.Ui32(v143) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v211 = v143 << (uint(int32(2)) % 32)
	v213 = int32(931135488)
	goto L67
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(127)) < base.Ui32(v143) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v211 = v143 << (uint(int32(3)) % 32)
	v213 = int32(922746880)
	goto L67
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(63)) < base.Ui32(v143) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v211 = v143 << (uint(int32(4)) % 32)
	v213 = int32(914358272)
	goto L67
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(31)) < base.Ui32(v143) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v211 = v143 << (uint(int32(5)) % 32)
	v213 = int32(905969664)
	goto L67
L81:
	;
	goto L82
L82:
	;
	if base.Ui32(int32(15)) < base.Ui32(v143) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v211 = v143 << (uint(int32(6)) % 32)
	v213 = int32(897581056)
	goto L67
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(7)) < base.Ui32(v143) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v211 = v143 << (uint(int32(7)) % 32)
	v213 = int32(889192448)
	goto L67
L87:
	;
	goto L88
L88:
	;
	if base.Ui32(int32(3)) < base.Ui32(v143) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v211 = v143 << (uint(int32(8)) % 32)
	v213 = int32(880803840)
	goto L67
L90:
	;
	goto L91
L91:
	;
	v206 = base.B2i32(v143 == int32(1))
	if v143 == int32(1) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v207 = int32(1024)
	goto L94
L93:
	;
	v207 = v143 << (uint(int32(9)) % 32)
	goto L94
L94:
	;
	if v143 == int32(1) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v210 = int32(864026624)
	goto L97
L96:
	;
	v210 = int32(872415232)
	goto L97
L97:
	;
	v211 = v207
	v213 = v210
	goto L67
L98:
	;
	if base.F32_gt(v136, v229) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v235 = v31 + int32(1)
	if v235 == v23 {
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
	v31 = v235
	goto L10
L103:
	;
	return int32(0)
L104:
	;
	goto L105
L105:
	;
	v257 = base.B2i32(v21 < v20)
	goto L4
}
func F_halfvec_le(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 float32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 float32
	_ = v230
	var v238 int32
	_ = v238
	var v254 int32
	_ = v254
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
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+4)))
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18)+4)))
	v23 = base.B2i32(v21 < v22)
	if v21 < v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v21 < v22 {
		goto L104
	} else {
		goto L105
	}
L5:
	;
	v24 = v21
	goto L7
L6:
	;
	v24 = v22
	goto L7
L7:
	;
	if v24 <= int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L9
L9:
	;
	v44 = v32 << (uint(int32(1)) % 32)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v27+v44))))
	v51 = v46 & int32(1023)
	v55 = v46 << (uint(int32(16)) % 32) & int32(-2147483648)
	v58 = int32(31)
	v59 = int32(base.Ui32(v46)>>(uint(int32(10))%32)) & v58
	if v59 != v58 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int32(0)
L11:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44+(v18+v27)))))
	v144 = v139 & int32(1023)
	v148 = v139 << (uint(int32(16)) % 32) & int32(-2147483648)
	v151 = int32(31)
	v152 = int32(base.Ui32(v139)>>(uint(int32(10))%32)) & v151
	if v152 != v151 {
		goto L58
	} else {
		goto L59
	}
L12:
	;
	v137 = base.F32_reinterpret_i32(v132 | v131<<(uint(int32(13))%32))
	goto L11
L13:
	;
	v131 = v51
	v132 = v59<<(uint(int32(23))%32) + v55 + int32(939524096)
	goto L12
L14:
	;
	if v46&int32(512) != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	if v59 != 0 {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v51 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v51 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v131 = int32(0)
	v132 = v55
	goto L12
L20:
	;
	v131 = int32(0)
	v132 = v55 | int32(2139095040)
	goto L12
L21:
	;
	goto L22
L22:
	;
	v131 = v51
	v132 = v55 | int32(2143289344)
	goto L12
L23:
	;
	v131 = v119 & int32(1022)
	v132 = v121 | v55
	goto L12
L24:
	;
	v119 = v51 << (uint(int32(1)) % 32)
	v121 = int32(939524096)
	goto L23
L25:
	;
	goto L26
L26:
	;
	if base.Ui32(int32(255)) < base.Ui32(v51) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v119 = v51 << (uint(int32(2)) % 32)
	v121 = int32(931135488)
	goto L23
L28:
	;
	goto L29
L29:
	;
	if base.Ui32(int32(127)) < base.Ui32(v51) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v119 = v51 << (uint(int32(3)) % 32)
	v121 = int32(922746880)
	goto L23
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(63)) < base.Ui32(v51) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v119 = v51 << (uint(int32(4)) % 32)
	v121 = int32(914358272)
	goto L23
L34:
	;
	goto L35
L35:
	;
	if base.Ui32(int32(31)) < base.Ui32(v51) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v119 = v51 << (uint(int32(5)) % 32)
	v121 = int32(905969664)
	goto L23
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(15)) < base.Ui32(v51) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v119 = v51 << (uint(int32(6)) % 32)
	v121 = int32(897581056)
	goto L23
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(int32(7)) < base.Ui32(v51) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v119 = v51 << (uint(int32(7)) % 32)
	v121 = int32(889192448)
	goto L23
L43:
	;
	goto L44
L44:
	;
	if base.Ui32(int32(3)) < base.Ui32(v51) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v119 = v51 << (uint(int32(8)) % 32)
	v121 = int32(880803840)
	goto L23
L46:
	;
	goto L47
L47:
	;
	v114 = base.B2i32(v51 == int32(1))
	if v51 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v115 = int32(1024)
	goto L50
L49:
	;
	v115 = v51 << (uint(int32(9)) % 32)
	goto L50
L50:
	;
	if v51 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v118 = int32(864026624)
	goto L53
L52:
	;
	v118 = int32(872415232)
	goto L53
L53:
	;
	v119 = v115
	v121 = v118
	goto L23
L54:
	;
	if base.F32_lt(v137, v230) != 0 {
		goto L97
	} else {
		goto L98
	}
L55:
	;
	v230 = base.F32_reinterpret_i32(v225 | v224<<(uint(int32(13))%32))
	goto L54
L56:
	;
	v224 = v144
	v225 = v152<<(uint(int32(23))%32) + v148 + int32(939524096)
	goto L55
L57:
	;
	if v139&int32(512) != 0 {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	if v152 != 0 {
		goto L56
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v144 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	if v144 != 0 {
		goto L57
	} else {
		goto L62
	}
L62:
	;
	v224 = int32(0)
	v225 = v148
	goto L55
L63:
	;
	v224 = int32(0)
	v225 = v148 | int32(2139095040)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v224 = v144
	v225 = v148 | int32(2143289344)
	goto L55
L66:
	;
	v224 = v212 & int32(1022)
	v225 = v214 | v148
	goto L55
L67:
	;
	v212 = v144 << (uint(int32(1)) % 32)
	v214 = int32(939524096)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui32(int32(255)) < base.Ui32(v144) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v212 = v144 << (uint(int32(2)) % 32)
	v214 = int32(931135488)
	goto L66
L71:
	;
	goto L72
L72:
	;
	if base.Ui32(int32(127)) < base.Ui32(v144) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v212 = v144 << (uint(int32(3)) % 32)
	v214 = int32(922746880)
	goto L66
L74:
	;
	goto L75
L75:
	;
	if base.Ui32(int32(63)) < base.Ui32(v144) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v212 = v144 << (uint(int32(4)) % 32)
	v214 = int32(914358272)
	goto L66
L77:
	;
	goto L78
L78:
	;
	if base.Ui32(int32(31)) < base.Ui32(v144) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v212 = v144 << (uint(int32(5)) % 32)
	v214 = int32(905969664)
	goto L66
L80:
	;
	goto L81
L81:
	;
	if base.Ui32(int32(15)) < base.Ui32(v144) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v212 = v144 << (uint(int32(6)) % 32)
	v214 = int32(897581056)
	goto L66
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(7)) < base.Ui32(v144) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v212 = v144 << (uint(int32(7)) % 32)
	v214 = int32(889192448)
	goto L66
L86:
	;
	goto L87
L87:
	;
	if base.Ui32(int32(3)) < base.Ui32(v144) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v212 = v144 << (uint(int32(8)) % 32)
	v214 = int32(880803840)
	goto L66
L89:
	;
	goto L90
L90:
	;
	v207 = base.B2i32(v144 == int32(1))
	if v144 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v208 = int32(1024)
	goto L93
L92:
	;
	v208 = v144 << (uint(int32(9)) % 32)
	goto L93
L93:
	;
	if v144 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v211 = int32(864026624)
	goto L96
L95:
	;
	v211 = int32(872415232)
	goto L96
L96:
	;
	v212 = v208
	v214 = v211
	goto L66
L97:
	;
	return int32(1)
L98:
	;
	goto L99
L99:
	;
	if base.F32_gt(v137, v230) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v238 = v32 + int32(1)
	if v238 == v24 {
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
	v32 = v238
	goto L9
L104:
	;
	v254 = int32(1)
	goto L106
L105:
	;
	v254 = base.B2i32(v21 <= v22)
	goto L106
L106:
	;
	return v254
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
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v63 float64
	_ = v63
	var v74 float64
	_ = v74
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v105 float64
	_ = v105
	var v111 float64
	_ = v111
	var v116 float64
	_ = v116
	var v120 float64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
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
					v45 = base.I64_reinterpret_f64(v40)
					v50 = base.I32_wrap_i64(int64(base.Ui64(v45)>>(uint(int64(32))%64))) & int32(2147483647)
					if base.Ui32(int32(1072693248)) <= base.Ui32(v50) {
						if base.I32_wrap_i64(v45)|(v50-int32(1072693248)) == int32(0) {
							if int64(0) <= v45 {
								v63 = float64(0)
							} else {
								v63 = float64(3.141592653589793)
							}
							v120 = v63
						} else {
							v120 = base.F64_div(float64(0), base.F64_sub(v40, v40))
						}
					} else {
						if base.Ui32(v50) <= base.Ui32(int32(1071644671)) {
							if base.Ui32(v50) < base.Ui32(int32(1012924417)) {
								v116 = float64(1.5707963267948966)
								v120 = v116
							} else {
								v74 = F_R(m, base.F64_mul(v40, v40))
								mBase = m.M
								v120 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v40, v74)), v40), float64(1.5707963267948966))
							}
						} else {
							if v45 < int64(0) {
								v86 = base.F64_mul(base.F64_add(v40, float64(1)), float64(0.5))
								v87 = base.F64_sqrt(v86)
								v88 = F_R(m, v86)
								mBase = m.M
								v93 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v87, base.F64_add(base.F64_mul(v87, v88), float64(-6.123233995736766e-17))))
								v120 = base.F64_add(v93, v93)
							} else {
								v98 = base.F64_mul(base.F64_sub(float64(1), v40), float64(0.5))
								v99 = base.F64_sqrt(v98)
								v100 = F_R(m, v98)
								mBase = m.M
								v105 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v99) & int64(-4294967296))
								v111 = base.F64_add(base.F64_add(base.F64_mul(v99, v100), base.F64_div(base.F64_sub(v98, base.F64_mul(v105, v105)), base.F64_add(v99, v105))), v105)
								v116 = base.F64_add(v111, v111)
								v120 = v116
							}
						}
					}
					v123 = F_Float8GetDatum(m, base.F64_div(v120, float64(3.141592653589793)))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v123
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						v136 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
						v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v137
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v136
						F_errmsg(m, int32(_a_F_halfvec_spherical_distance_0), v9)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_halfvec_spherical_distance_1), int32(80), int32(_a_F_halfvec_spherical_distance_2))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
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
