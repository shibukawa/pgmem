package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbType(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(18) {
		v22 = v8
		m.G0 = v6 + int32(16)
		return v22
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		if v12&int32(536870912) != 0 {
			v22 = int32(17)
			m.G0 = v6 + int32(16)
			return v22
		} else {
			if v12&int32(1073741824) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v33
					F_errmsg_internal(m, int32(_a_F_JsonbType_0), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_JsonbType_1), int32(3947), int32(_a_F_JsonbType_2))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v22 = int32(16)
				m.G0 = v6 + int32(16)
				return v22
			}
		}
	}
}
func F_JsonbValueAsText(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		v60 = int32(0)
		m.G0 = v6 + int32(32)
		return v60
	case 1:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = F_cstring_to_text_with_len(m, v18, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v60 = v20
			m.G0 = v6 + int32(32)
			return v60
		}
	case 2:
		v24 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
		v25 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v28 = F_cstring_to_text(m, base.I32_wrap_i64(v25))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v60 = v28
				m.G0 = v6 + int32(32)
				return v60
			}
		}
	case 3:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v9 != int32(1) {
			v58 = F_cstring_to_text_with_len(m, int32(_a_F_JsonbValueAsText_0), int32(5))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v60 = v58
				m.G0 = v6 + int32(32)
				return v60
			}
		} else {
			v14 = F_cstring_to_text_with_len(m, int32(_a_F_JsonbValueAsText_1), int32(4))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v60 = v14
				m.G0 = v6 + int32(32)
				return v60
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v46
			F_errmsg_internal(m, int32(_a_F_JsonbValueAsText_2), v6)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_JsonbValueAsText_3), int32(1843), int32(_a_F_JsonbValueAsText_4))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 18:
		v31 = v6 + int32(16)
		F_initStringInfo(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v36 = F_JsonbToCString(m, v31, v34, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				v40 = F_cstring_to_text_with_len(m, v38, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v60 = v40
					m.G0 = v6 + int32(32)
					return v60
				}
			}
		}
	}
}
func F_convertJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v468 int32
	_ = v468
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	v13 = m.G0
	v15 = v13 - int32(80)
	m.G0 = v15
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L96
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L92
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L88
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L84
	}
L8:
	;
	m.G0 = v15 + int32(80)
	return
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if base.Ui32(v21) < base.Ui32(int32(4)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_convertJsonbScalar(m, l0, l1, l2)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L83
	}
L11:
	;
	switch v21 - int32(16) {
	case 0:
		goto L14
	case 1:
		goto L13
	default:
		goto L12
	case 16:
		goto L10
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L80
	}
L13:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v241 = (v237 + int32(3)) & int32(-4)
	v242 = v241 - v237
	F_enlargeStringInfo(m, l0, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L45
	}
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = (v27 + int32(3)) & int32(-4)
	v32 = v31 - v27
	F_enlargeStringInfo(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v32 + v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38+v36))) = uint8(v40)
	if v32 <= v40 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	F_enlargeStringInfo(m, l0, int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L28
	}
L17:
	;
	v45 = v32 & int32(3)
	v46 = int32(0)
	if base.Ui32(v27-v31) <= base.Ui32(int32(-4)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = v46
	v58 = int32(0)
	goto L21
L19:
	;
	v96 = v46
	goto L20
L20:
	;
	v109 = v96
	v110 = int32(0)
	goto L25
L21:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v68 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v65+v35+v57))) = uint8(v68)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70+v35+v57)+1)) = uint8(v68)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v35+v57)+2)) = uint8(v68)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v35+v57)+3)) = uint8(v68)
	v85 = int32(4)
	v86 = v57 + v85
	v88 = v58 + v85
	if v88 != v32&int32(2147483644) {
		v57 = v86
		v58 = v88
		goto L21
	} else {
		goto L23
	}
L22:
	;
	if v45 == int32(0) {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v96 = v86
	goto L20
L25:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v117+v35+v109))) = uint8(v120)
	v122 = int32(1)
	v125 = v110 + v122
	if v125 != v45 {
		v109 = v109 + v122
		v110 = v125
		goto L25
	} else {
		goto L27
	}
L26:
	;
	goto L16
L27:
	;
	goto L26
L28:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v145 = v143 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v145))) = uint8(v149)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v139 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v155 = int32(1342177280)
	goto L31
L30:
	;
	v155 = int32(1073741824)
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143+v151))) = v155 | v26
	v159 = v26 << (uint(int32(2)) % 32)
	F_enlargeStringInfo(m, l0, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v163 = v162 + v159
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165+v163))) = uint8(v167)
	if v167 < v26 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v173 = int32(0)
	v179 = v173
	v180 = v173
	v181 = v162
	goto L36
L34:
	;
	goto L35
L35:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v230 = v229 - v27
	if int32(268435456) <= v230 {
		goto L6
	} else {
		goto L44
	}
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_convertJsonbValue(m, l0, v15+int32(76), v189+v179<<(uint(int32(5))%32), l3+int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	goto L35
L38:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v198 = v195&int32(268435455) + v180
	if base.Ui32(int32(268435456)) <= base.Ui32(v198) {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v179&int32(31) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v210 = v195
	goto L42
L41:
	;
	v210 = v195&int32(1879048192) | v198 | int32(-2147483648)
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201+v181))) = v210
	v215 = v179 + int32(1)
	if v215 != v26 {
		v179 = v215
		v180 = v198
		v181 = v181 + int32(4)
		goto L36
	} else {
		goto L43
	}
L43:
	;
	goto L37
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v230 | int32(1342177280)
	goto L8
L45:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v246 = v242 + v245
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v250 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v248+v246))) = uint8(v250)
	if v242 <= v250 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_enlargeStringInfo(m, l0, int32(4))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L58
	}
L47:
	;
	v255 = v242 & int32(3)
	v256 = int32(0)
	if base.Ui32(v237-v241) <= base.Ui32(int32(-4)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v267 = v256
	v268 = int32(0)
	goto L51
L49:
	;
	v306 = v256
	goto L50
L50:
	;
	v319 = v306
	v320 = int32(0)
	goto L55
L51:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v278 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v275+v245+v267))) = uint8(v278)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v280+v245+v267)+1)) = uint8(v278)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v285+v245+v267)+2)) = uint8(v278)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v290+v245+v267)+3)) = uint8(v278)
	v295 = int32(4)
	v296 = v267 + v295
	v298 = v268 + v295
	if v298 != v242&int32(2147483644) {
		v267 = v296
		v268 = v298
		goto L51
	} else {
		goto L53
	}
L52:
	;
	if v255 == int32(0) {
		goto L46
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v306 = v296
	goto L50
L55:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v330 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v245+v319))) = uint8(v330)
	v332 = int32(1)
	v335 = v320 + v332
	if v335 != v255 {
		v319 = v319 + v332
		v320 = v335
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L46
L57:
	;
	goto L56
L58:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v354 = v352 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v358 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v356+v354))) = uint8(v358)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v352+v360))) = v236 | int32(536870912)
	v366 = v236 << (uint(int32(3)) % 32)
	F_enlargeStringInfo(m, l0, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v370 = v366 + v369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v370
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v372+v370))) = uint8(v374)
	if v374 < v236 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v378 = int32(0)
	v384 = v369
	v385 = v378
	v386 = v378
	goto L63
L61:
	;
	goto L62
L62:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v483 = v482 - v237
	if int32(268435456) <= v483 {
		goto L3
	} else {
		goto L79
	}
L63:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_convertJsonbScalar(m, l0, v15+int32(76), v394+v385*int32(72))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v429 = v418
	v430 = int32(0)
	v431 = v403
	goto L71
L65:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v403 = v400&int32(268435455) + v386
	if base.Ui32(int32(268435456)) <= base.Ui32(v403) {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v385&int32(31) != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v415 = v400
	goto L69
L68:
	;
	v415 = v400&int32(1879048192) | v403 | int32(-2147483648)
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v406+v384))) = v415
	v418 = v384 + int32(4)
	v420 = v385 + int32(1)
	if v420 != v236 {
		v384 = v418
		v385 = v420
		v386 = v403
		goto L63
	} else {
		goto L70
	}
L70:
	;
	goto L64
L71:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_convertJsonbValue(m, l0, v15+int32(76), v439+v430*int32(72)+int32(32), l3+int32(1))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L62
L73:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v15)+76))
	v450 = v447&int32(268435455) + v431
	if base.Ui32(int32(268435456)) <= base.Ui32(v450) {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if (v430+v236)&int32(31) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v463 = v447
	goto L77
L76:
	;
	v463 = v447&int32(1879048192) | v450 | int32(-2147483648)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453+v429))) = v463
	v468 = v430 + int32(1)
	if v468 != v236 {
		v429 = v429 + int32(4)
		v430 = v468
		v431 = v450
		goto L71
	} else {
		goto L78
	}
L78:
	;
	goto L72
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v483 | int32(1342177280)
	goto L8
L80:
	;
	F_errmsg_internal(m, int32(_a_F_convertJsonbValue_0), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1738), int32(_a_F_convertJsonbValue_2))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	goto L8
L84:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(268435455)
	F_errmsg(m, int32(_a_F_convertJsonbValue_3), v15)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1799), int32(_a_F_convertJsonbValue_4))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(268435455)
	F_errmsg(m, int32(_a_F_convertJsonbValue_3), v15+int32(16))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1819), int32(_a_F_convertJsonbValue_4))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = int32(268435455)
	F_errmsg(m, int32(_a_F_convertJsonbValue_5), v15+int32(32))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1880), int32(_a_F_convertJsonbValue_6))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = int32(268435455)
	F_errmsg(m, int32(_a_F_convertJsonbValue_5), v15+int32(48))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1915), int32(_a_F_convertJsonbValue_6))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(268435455)
	F_errmsg(m, int32(_a_F_convertJsonbValue_5), v15-int32(-64))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_convertJsonbValue_1), int32(1935), int32(_a_F_convertJsonbValue_6))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_equalsJsonbScalarValue(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v121
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v9) {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	switch v5 {
	case 0:
		v121 = int32(1)
		goto L1
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	default:
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L16
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L11
	} else {
		goto L13
	}
L7:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	return base.B2i32(v25 == v26)
L8:
	;
	v16 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
	v17 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+8)))
	v18 = F_DirectFunctionCall2Coll(m, int32(1468), int32(0), v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v9 == v10 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	return int32(0)
L11:
	;
	return int32(0)
L12:
	;
	return base.B2i32(v18 != int64(0))
L13:
	;
	F_errmsg_internal(m, int32(_a_F_equalsJsonbScalarValue_0), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_equalsJsonbScalarValue_1), int32(1546), int32(_a_F_equalsJsonbScalarValue_2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	F_errmsg_internal(m, int32(_a_F_equalsJsonbScalarValue_3), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_equalsJsonbScalarValue_1), int32(1549), int32(_a_F_equalsJsonbScalarValue_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v121 = base.B2i32(v118 == int32(0))
	goto L1
L20:
	;
	v118 = int32(0)
	goto L19
L21:
	;
	v92 = v87
	v93 = v88
	v94 = v89
	goto L31
L22:
	;
	if (v55|v56)&int32(3) != 0 {
		v87 = v55
		v88 = v56
		v89 = v9
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v80 = v55
	v81 = v56
	v82 = v9
	goto L24
L24:
	;
	if v82 == int32(0) {
		goto L20
	} else {
		goto L30
	}
L25:
	;
	v64 = v55
	v65 = v56
	v66 = v9
	goto L26
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v69 != v70 {
		v87 = v64
		v88 = v65
		v89 = v66
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v80 = v75
	v81 = v73
	v82 = v77
	goto L24
L28:
	;
	v72 = int32(4)
	v73 = v65 + v72
	v75 = v64 + v72
	v77 = v66 - v72
	if base.Ui32(int32(3)) < base.Ui32(v77) {
		v64 = v75
		v65 = v73
		v66 = v77
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v87 = v80
	v88 = v81
	v89 = v82
	goto L21
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v97 == v98 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v118 = v97 - v98
	goto L19
L33:
	;
	v100 = int32(1)
	v105 = v94 - v100
	if v105 != 0 {
		v92 = v92 + v100
		v93 = v93 + v100
		v94 = v105
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L20
}
func F_jsonb_array_elements_text(m *base.Module, l0 int32) int64 {
	var v6 int32
	_ = v6
	F_elements_worker_jsonb(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_jsonb_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	v5 = l4
	v6 = l5
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if l0&int32(1) == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L5
	} else {
		goto L28
	}
L2:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v19
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v21
	F_pushJsonbValue(m, v13+int32(24), int32(6), v19)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L23
	}
L5:
	;
	return int64(0)
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+41)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+40)) = uint8(v6)
	if int32(0) < l0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v46 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_pushJsonbValue(m, v13+int32(24), int32(7), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L21
	}
L10:
	;
	v48 = l2 + v46
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49 == int32(1) {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v52 = int32(0)
	if base.B2i32(v5 == v52)|v6 == v52 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v92 = v46 + int32(2)
	if v92 < l0 {
		v46 = v92
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)))
	if v57&int32(1) != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l1+v46<<(uint(int32(3))%32))))
	v66 = v13 + int32(24)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3+v46<<(uint(int32(2))%32))))
	F_add_jsonb(m, v63, int32(0), v66, v70, int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v75 = v46 | int32(1)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(l1+v75<<(uint(int32(3))%32))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v75))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l3+v75<<(uint(int32(2))%32))))
	F_add_jsonb(m, v79, v81, v66, v85, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	goto L11
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v111 = F_JsonbValueToJsonb(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v13 + int32(48)
	return base.I64_extend_i32_u(v111)
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(_a_F_jsonb_build_object_worker_0), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_jsonb_build_object_worker_1)
	F_errhint(m, int32(_a_F_jsonb_build_object_worker_2), v13+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_jsonb_build_object_worker_3), int32(1142), int32(_a_F_jsonb_build_object_worker_4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v46 | int32(1)
	F_errmsg(m, int32(_a_F_jsonb_build_object_worker_5), v13)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_jsonb_build_object_worker_3), int32(1158), int32(_a_F_jsonb_build_object_worker_4))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_delete_idx(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v16
	v18 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v22&int32(268435456) == v16 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L36
	}
L4:
	;
	if v22&int32(536870912) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L32
	}
L7:
	;
	if v22&int32(268435455) == int32(0) {
		v87 = v11
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v8 - int32(-64)
	return base.I64_extend_i32_u(v87)
L9:
	;
	v35 = F_JsonbIteratorInit(m, v11+int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v35
	v41 = F_JsonbIteratorNext(m, v6+int32(-28), v8, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if base.Ui32(v43-v15) <= base.Ui32(v44) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v15
	goto L14
L13:
	;
	v48 = v43
	goto L14
L14:
	;
	if v15 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = v48 + v44
	goto L17
L16:
	;
	v52 = v15
	goto L17
L17:
	;
	if base.Ui32(v44) <= base.Ui32(v52) {
		v87 = v11
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_pushJsonbValue(m, v6+int32(-24), v41, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(0)
	goto L20
L20:
	;
	v68 = F_JsonbIteratorNext(m, v6+int32(-28), v8, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L25
	}
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v83 = F_JsonbValueToJsonb(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	goto L21
L23:
	;
	if base.Ui32(v68) < base.Ui32(int32(4)) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v72 = v60 + int32(1)
	if v60 == v52 {
		v60 = v72
		goto L20
	} else {
		goto L26
	}
L25:
	;
	switch v68 {
	case 0:
		goto L22
	default:
		v73 = v60
		goto L23
	case 3:
		goto L24
	}
L26:
	;
	v73 = v72
	goto L23
L27:
	;
	v79 = v8
	goto L29
L28:
	;
	v79 = int32(0)
	goto L29
L29:
	;
	F_pushJsonbValue(m, v6+int32(-24), v68, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v60 = v73
	goto L20
L31:
	;
	v87 = v83
	goto L8
L32:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_idx_0), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_idx_1), int32(_a_F_jsonb_delete_idx_2), int32(_a_F_jsonb_delete_idx_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_jsonb_delete_idx_4), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_jsonb_delete_idx_1), int32(_a_F_jsonb_delete_idx_5), int32(_a_F_jsonb_delete_idx_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_delete_path(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
			v19 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v19
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v19
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v23 < int32(2) {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v26&int32(268435456) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_jsonb_delete_path_0), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_jsonb_delete_path_1), int32(_a_F_jsonb_delete_path_2), int32(_a_F_jsonb_delete_path_3))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v26&int32(268435455) == int32(0) {
						v65 = v10
						m.G0 = v7 + int32(48)
						return base.I64_extend_i32_u(v65)
					} else {
						F_deconstruct_array_builtin(m, v15, int32(25), v7+int32(44), v7+int32(40), v7+int32(36))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
							if v42 == int32(0) {
								v65 = v10
								m.G0 = v7 + int32(48)
								return base.I64_extend_i32_u(v65)
							} else {
								v47 = F_JsonbIteratorInit(m, v10+int32(4))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
									v57 = int32(0)
									F_setPath(m, v7+int32(32), v52, v53, v54, v7+int32(8), v57, v57, int32(2))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int64(0)
									} else {
										v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										v63 = F_JsonbValueToJsonb(m, v62)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int64(0)
										} else {
											v65 = v63
											m.G0 = v7 + int32(48)
											return base.I64_extend_i32_u(v65)
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_jsonb_delete_path_4), int32(0))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_jsonb_delete_path_1), int32(_a_F_jsonb_delete_path_5), int32(_a_F_jsonb_delete_path_3))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int64(0)
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
func F_jsonb_each(m *base.Module, l0 int32) int64 {
	var v7 int32
	_ = v7
	F_each_worker_jsonb(m, l0, int32(_a_F_jsonb_each_0), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_jsonb_extract_path_text(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_jsonb_path_all(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	v2 = int32(0)
	v6 = int64(0)
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_strlen(m, v11)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v2
	v25 = v9 + int32(60)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_in[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = F_makeJsonLexContextCstringLen(m, v25, v11, v12, v28, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v13
		v35 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+56)) = uint8(v35)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1451)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1452)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(1453)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1454)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1455)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(1456)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(40)
		v52 = F_pg_parse_json_or_errsave(m, v25, v9, v13)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			if v52 != 0 {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
				v55 = F_JsonbValueToJsonb(m, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int64(0)
				} else {
					v58 = base.I64_extend_i32_u(v55)
					m.G0 = v9 + int32(128)
					return v58
				}
			} else {
				v58 = v6
				m.G0 = v9 + int32(128)
				return v58
			}
		}
	}
}
func F_jsonb_in_object_end(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pushJsonbValue(m, l0, int32(7), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_insert(m *base.Module, l0 int32) int64 {
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
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v20 = F_pg_detoast_datum(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(0)
				v25 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v25
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v25
				v30 = v9 + int32(48)
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(18)
				v33 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v20 + v33
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - v33
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v42 < int32(2) {
					v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
					if v45&int32(16) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_jsonb_insert_0), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_jsonb_insert_1), int32(_a_F_jsonb_insert_2), int32(_a_F_jsonb_insert_3))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						F_deconstruct_array_builtin(m, v17, int32(25), v9+int32(44), v9+int32(40), v9+int32(36))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
							if v57 != 0 {
								v60 = F_JsonbIteratorInit(m, v12+int32(4))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v60
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v68 = int32(8)
									if v22 == int64(0) {
										v75 = v68
									} else {
										v75 = int32(16)
									}
									F_setPath(m, v9+int32(32), v65, v66, v67, v9+v68, int32(0), v30, v75)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int64(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
										v79 = F_JsonbValueToJsonb(m, v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int64(0)
										} else {
											v81 = v79
											m.G0 = v9 + int32(80)
											return base.I64_extend_i32_u(v81)
										}
									}
								}
							} else {
								v81 = v12
								m.G0 = v9 + int32(80)
								return base.I64_extend_i32_u(v81)
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(352845954))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_jsonb_insert_4), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_jsonb_insert_1), int32(_a_F_jsonb_insert_5), int32(_a_F_jsonb_insert_3))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int64(0)
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
}
func F_jsonb_int2(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14322(m, l0, int32(1457), int32(_a_F_jsonb_int2_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_int8(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14322(m, l0, int32(1409), int32(_a_F_jsonb_int8_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v17 != int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 != int32(0)))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 != int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v17 != int32(0)))
					}
				}
			}
		}
	}
}
func F_jsonb_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v2 = l1
	v3 = l2
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 == v4 {
		v31 = int32(0)
		if v12 == v31 {
			v39 = v31
		} else {
			v34 = v31
			v35 = v4
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
			v39 = v35
		}
		v42 = v39
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		switch v17 - int32(435) {
		case 0:
			if v12 == int32(0) {
				v42 = int32(1)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+168))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
				v34 = v24
				v35 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		case 1:
			if v12 == int32(0) {
				v42 = int32(2)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)+376))
				v34 = v29
				v35 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
				v42 = v39
			}
		default:
			v31 = int32(0)
			if v12 == v31 {
				v39 = v31
			} else {
				v34 = v31
				v35 = v4
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = v34
				v39 = v35
			}
			v42 = v39
		}
	}
	if v42 != 0 {
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v43 == int32(1) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			v48 = F_MemoryContextAllocZero(m, v46, int32(36))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int64(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
				F_pushJsonbValue(m, v48, int32(6), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v58)+40)) = uint8(v3)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
					*(*uint8)(unsafe.Add(mBase, uint32(v60)+41)) = uint8(v2)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v64 = F_get_fn_expr_argtype(m, v62, int32(1))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						if v64 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v143 = m.ExcPending
								if v143 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_jsonb_object_agg_transfn_worker_0), int32(0))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_jsonb_object_agg_transfn_worker_1), int32(1615), int32(_a_F_jsonb_object_agg_transfn_worker_2))
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							F_json_categorize_type(m, v64, int32(1), v48+int32(20), v48+int32(24))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v77 = F_get_fn_expr_argtype(m, v75, int32(2))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int64(0)
								} else {
									if v77 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(50856066))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_jsonb_object_agg_transfn_worker_0), int32(0))
												mBase = m.M
												v163 = m.ExcPending
												if v163 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_jsonb_object_agg_transfn_worker_1), int32(1625), int32(_a_F_jsonb_object_agg_transfn_worker_2))
													mBase = m.M
													v168 = m.ExcPending
													if v168 != 0 {
														return int64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										F_json_categorize_type(m, v77, int32(1), v48+int32(28), v48+int32(32))
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int64(0)
										} else {
											v89 = v48
											v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
											if v91 == int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(50856066))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_jsonb_object_agg_transfn_worker_3), int32(0))
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_jsonb_object_agg_transfn_worker_1), int32(1639), int32(_a_F_jsonb_object_agg_transfn_worker_2))
															mBase = m.M
															v184 = m.ExcPending
															if v184 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v94 = int32(0)
												if base.B2i32(v2 == v94)|v3 == v94 {
													v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
													if v99&int32(1) != 0 {
														m.G0 = v9 + int32(16)
														return base.I64_extend_i32_u(v89)
													} else {
														v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
														v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
														v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
														F_datum_to_jsonb_internal(m, v102, int32(0), v89, v104, v105, int32(1))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return int64(0)
														} else {
															v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
															if v109 != 0 {
																v112 = int64(0)
															} else {
																v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
																v112 = v111
															}
															v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
															F_datum_to_jsonb_internal(m, v112, v109, v89, v113, v114, int32(0))
															mBase = m.M
															v117 = m.ExcPending
															if v117 != 0 {
																return int64(0)
															} else {
																m.G0 = v9 + int32(16)
																return base.I64_extend_i32_u(v89)
															}
														}
													}
												} else {
													v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
													v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
													F_datum_to_jsonb_internal(m, v102, int32(0), v89, v104, v105, int32(1))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return int64(0)
													} else {
														v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
														if v109 != 0 {
															v112 = int64(0)
														} else {
															v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
															v112 = v111
														}
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
														F_datum_to_jsonb_internal(m, v112, v109, v89, v113, v114, int32(0))
														mBase = m.M
														v117 = m.ExcPending
														if v117 != 0 {
															return int64(0)
														} else {
															m.G0 = v9 + int32(16)
															return base.I64_extend_i32_u(v89)
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
		} else {
			v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v89 = v88
			v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v91 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_jsonb_object_agg_transfn_worker_3), int32(0))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_jsonb_object_agg_transfn_worker_1), int32(1639), int32(_a_F_jsonb_object_agg_transfn_worker_2))
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v94 = int32(0)
				if base.B2i32(v2 == v94)|v3 == v94 {
					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
					if v99&int32(1) != 0 {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v89)
					} else {
						v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
						F_datum_to_jsonb_internal(m, v102, int32(0), v89, v104, v105, int32(1))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int64(0)
						} else {
							v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
							if v109 != 0 {
								v112 = int64(0)
							} else {
								v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
								v112 = v111
							}
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
							F_datum_to_jsonb_internal(m, v112, v109, v89, v113, v114, int32(0))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v89)
							}
						}
					}
				} else {
					v102 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
					F_datum_to_jsonb_internal(m, v102, int32(0), v89, v104, v105, int32(1))
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int64(0)
					} else {
						v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
						if v109 != 0 {
							v112 = int64(0)
						} else {
							v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
							v112 = v111
						}
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
						F_datum_to_jsonb_internal(m, v112, v109, v89, v113, v114, int32(0))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v89)
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v127 = m.ExcPending
		if v127 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_jsonb_object_agg_transfn_worker_4), int32(0))
			mBase = m.M
			v131 = m.ExcPending
			if v131 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_jsonb_object_agg_transfn_worker_1), int32(1594), int32(_a_F_jsonb_object_agg_transfn_worker_2))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_jsonb_object_two_arg(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v22 int64
	_ = v22
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v20 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v20
	v22 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v22
	F_pushJsonbValue(m, v8+int32(32), int32(6), v20)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.B2i32(v19 != v18)|base.B2i32(int32(1) < v19) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L44
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L36
	}
L10:
	;
	F_deconstruct_array_builtin(m, v11, int32(25), v8+int32(76), v8+int32(68), v8+int32(60))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_pushJsonbValue(m, v8+int32(32), int32(7), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L34
	}
L13:
	;
	F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(72), v8-int32(-64), v8+int32(56))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	if v56 != v57 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if int32(0) < v56 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	F_pfree(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v62))))
	if v69 == int32(1) {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v73 = v62 << (uint(int32(3)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v73+v74)))
	v77 = F_text_to_cstring(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v79 = F_strlen(m, v77)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v79
	v82 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v82
	F_pushJsonbValue(m, v8+int32(32), v82, v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v62))))
	if v91 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v104 = int32(0)
	goto L26
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v93+v73)))
	v96 = F_text_to_cstring(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v104
	F_pushJsonbValue(m, v8+int32(32), int32(2), v8)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v98 = F_strlen(m, v96)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v98
	v104 = int32(1)
	goto L26
L28:
	;
	v112 = v62 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	if v112 < v113 {
		v62 = v112
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L20
L30:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	F_pfree(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
	F_pfree(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	F_pfree(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v144 = F_JsonbValueToJsonb(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	m.G0 = v8 + int32(80)
	return base.I64_extend_i32_u(v144)
L36:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_two_arg_0), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_two_arg_1), int32(1406), int32(_a_F_jsonb_object_two_arg_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_two_arg_3), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_two_arg_1), int32(1417), int32(_a_F_jsonb_object_two_arg_2))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_jsonb_object_two_arg_4), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_two_arg_1), int32(1428), int32(_a_F_jsonb_object_two_arg_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_path_exists_tz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_exists_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_path_match_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v22 == int32(4) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v31 = int32(1)
	v32 = int32(0)
	goto L6
L6:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(8589934592)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v11
	v42 = F_executeJsonPath(m, v19, v32, int32(1534), int32(1535), v14, base.B2i32(v31 == v33), v11, l1)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v31 = base.B2i32(v28 != int64(0))
	v32 = v26
	goto L6
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v44 != v14 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v14)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v48 != v19 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	F_pfree(m, v19)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v52 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	goto L15
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L25
	}
L18:
	;
	m.G0 = v11 + int32(80)
	return v62
L19:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v59)
	v62 = int64(0)
	goto L18
L20:
	;
	if v31 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	switch v55 {
	case 0:
		goto L22
	default:
		goto L20
	case 3:
		goto L23
	}
L22:
	;
	goto L19
L23:
	;
	v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
	v62 = v56
	goto L18
L24:
	;
	goto L19
L25:
	;
	F_errcode(m, int32(135004290))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_jsonb_path_match_internal_0), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_jsonb_path_match_internal_1), int32(521), int32(_a_F_jsonb_path_match_internal_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_path_ops__extract_nodes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v10
		F_JsonbHashScalarValue(m, l2, v8+int32(12))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)))
			v20 = F_palloc(m, int32(16))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(2)
				v25 = F_lappend(m, l3, v20)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v28 = v25
					m.G0 = v8 + int32(16)
					return v28
				}
			}
		}
	} else {
		v28 = l3
		m.G0 = v8 + int32(16)
		return v28
	}
}
func F_jsonb_path_query_array_tz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_array_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_first(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_first_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_subscript_check_subscripts(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L21
	} else {
		goto L25
	}
L2:
	;
	return int32(1)
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = int32(0)
	goto L9
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v17 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v19 != int32(23) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11))) = uint8(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v24 <= int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v30))))
	if v36 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L2
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v30))))
	if v41 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v81 = v30 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v81 < v82 {
		v30 = v81
		goto L9
	} else {
		goto L24
	}
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v44 == int32(1) {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = v30 << (uint(int32(3)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v53+v54)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v30<<(uint(int32(2))%32))))
	if v61 == int32(23) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v48 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
	return int32(0)
L18:
	;
	v66 = F_DirectFunctionCall1Coll(m, int32(1438), int32(0), v56)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = v56
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v75+v53))) = v74
	goto L13
L21:
	;
	return int32(0)
L22:
	;
	v71 = F_cstring_to_text(m, base.I32_wrap_i64(v66))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v74 = base.I64_extend_i32_u(v71)
	goto L20
L24:
	;
	goto L10
L25:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_jsonb_subscript_check_subscripts_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_jsonb_subscript_check_subscripts_1), int32(205), int32(_a_F_jsonb_subscript_check_subscripts_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_to_record(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(_a_F_jsonb_to_record_0), v3, v3, v3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_jsonb_to_tsvector_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = F_parse_jsonb_index_flags(m, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int64(0)
			} else {
				*(*uint32)(unsafe.Add(mBase, uint32(v10)+28)) = uint32(v12)
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
				v29 = v10 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v29
				F_iterate_jsonb_values(m, v14, v21, v10+int32(24))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = F_make_tsvector(m, v29)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v37 != v14 {
							F_pfree(m, v14)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v41 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int64(0)
									} else {
										m.G0 = v10 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								} else {
									m.G0 = v10 + int32(32)
									return base.I64_extend_i32_u(v35)
								}
							}
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v41 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									m.G0 = v10 + int32(32)
									return base.I64_extend_i32_u(v35)
								}
							} else {
								m.G0 = v10 + int32(32)
								return base.I64_extend_i32_u(v35)
							}
						}
					}
				}
			}
		}
	}
}
