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
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
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
					F_errmsg_internal(m, int32(30502), v6)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520780), int32(3629), int32(387029))
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
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		v61 = int32(0)
		m.G0 = v6 + int32(32)
		return v61
	case 1:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = F_cstring_to_text_with_len(m, v18, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v61 = v20
			m.G0 = v6 + int32(32)
			return v61
		}
	case 2:
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v25 = F_DirectFunctionCall1Coll(m, int32(618), int32(0), v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = F_cstring_to_text(m, v25)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v61 = v27
				m.G0 = v6 + int32(32)
				return v61
			}
		}
	case 3:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v9 != int32(1) {
			v59 = F_cstring_to_text_with_len(m, int32(376448), int32(5))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = v59
				m.G0 = v6 + int32(32)
				return v61
			}
		} else {
			v14 = F_cstring_to_text_with_len(m, int32(358953), int32(4))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v61 = v14
				m.G0 = v6 + int32(32)
				return v61
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v47
			F_errmsg_internal(m, int32(504746), v6)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(514510), int32(1843), int32(69184))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 18:
		F_initStringInfo(m, v6+int32(16))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v37 = F_JsonbToCString(m, v6+int32(16), v35, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)+20))
				v41 = F_cstring_to_text_with_len(m, v39, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v61 = v41
					m.G0 = v6 + int32(32)
					return v61
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
	var v94 int32
	_ = v94
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
	var v304 int32
	_ = v304
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
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
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
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
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
	v94 = v46
	goto L20
L20:
	;
	if v45 == int32(0) {
		goto L16
	} else {
		goto L24
	}
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
	v94 = v86
	goto L20
L23:
	;
	goto L22
L24:
	;
	v109 = v94
	v110 = int32(0)
	goto L25
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
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_convertJsonbValue(m, l0, v15+int32(76), v189+v179*int32(20), l3+int32(1))
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
	v304 = v256
	goto L50
L50:
	;
	if v255 == int32(0) {
		goto L46
	} else {
		goto L54
	}
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
	v304 = v296
	goto L50
L53:
	;
	goto L52
L54:
	;
	v319 = v304
	v320 = int32(0)
	goto L55
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
	v370 = v369 + v366
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
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_convertJsonbScalar(m, l0, v15+int32(76), v394+v385*int32(44))
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
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_convertJsonbValue(m, l0, v15+int32(76), v439+v430*int32(44)+int32(20), l3+int32(1))
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
	F_errmsg_internal(m, int32(87051), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(517863), int32(1624), int32(362300))
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
	F_errmsg(m, int32(167682), v15)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(517863), int32(1685), int32(26722))
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
	F_errmsg(m, int32(167682), v15+int32(16))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(517863), int32(1705), int32(26722))
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
	F_errmsg(m, int32(167749), v15+int32(32))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(517863), int32(1766), int32(118285))
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
	F_errmsg(m, int32(167749), v15+int32(48))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(517863), int32(1801), int32(118285))
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
	F_errmsg(m, int32(167749), v15-int32(-64))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(517863), int32(1821), int32(118285))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
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
	var v122 int32
	_ = v122
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v122
L2:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v9) {
		goto L22
	} else {
		goto L23
	}
L3:
	;
	switch v5 {
	case 0:
		v122 = int32(1)
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
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	return base.B2i32(v25 == v26)
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = F_DirectFunctionCall2Coll(m, int32(1344), int32(0), v16, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
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
	return base.B2i32(v18 != int32(0))
L13:
	;
	F_errmsg_internal(m, int32(383690), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(517863), int32(1432), int32(361991))
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
	F_errmsg_internal(m, int32(338174), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(517863), int32(1435), int32(361991))
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
	v122 = base.B2i32(v118 == int32(0))
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
func F_jsonb_array_elements_text(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	F_elements_worker_jsonb(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_build_object_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
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
	var v19 int64
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L28
	}
L2:
	;
	v19 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v19
	v27 = F_pushJsonbValue(m, v13+int32(32), int32(6), int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
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
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L23
	}
L5:
	;
	return int32(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v27
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+29)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+28)) = uint8(v6)
	if int32(0) < l0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v44 = v7
	goto L10
L8:
	;
	goto L9
L9:
	;
	v102 = F_pushJsonbValue(m, v13+int32(32), int32(7), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L21
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v44))))
	if v48 == int32(1) {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v52 = v44 | int32(1)
	if v5 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v86 = v44 + int32(2)
	if v86 < l0 {
		v44 = v86
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v60 = v44 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l1+v60)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3+v60)))
	F_add_jsonb(m, v62, int32(0), v13+int32(32), v67, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L18
	}
L15:
	;
	if v6 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v52))))
	if v56&int32(1) != 0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v72 = v52 << (uint(int32(2)) % 32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1+v72)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v52))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3+v72)))
	F_add_jsonb(m, v74, v76, v13+int32(32), v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v102
	v105 = F_JsonbValueToJsonb(m, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v13 + int32(48)
	return v105
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(130213), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(709019)
	F_errhint(m, int32(619106), v13+int32(16))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(521213), int32(1137), int32(229475))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v44 | int32(1)
	F_errmsg(m, int32(314878), v13)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(521213), int32(1153), int32(229475))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
func F_jsonb_delete_idx(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
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
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v20&int32(268435456) == v18 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L4:
	;
	if v20&int32(536870912) != 0 {
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
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L35
	}
L7:
	;
	if v20&int32(268435455) == int32(0) {
		v101 = v13
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v10 + int32(32)
	return v101
L9:
	;
	v33 = F_JsonbIteratorInit(m, v13+int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v33
	v41 = F_JsonbIteratorNext(m, v10+int32(24), v10+int32(4), int32(0))
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if base.Ui32(v43-v17) <= base.Ui32(v44) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v48 = v17
	goto L14
L13:
	;
	v48 = v43
	goto L14
L14:
	;
	if v17 < int32(0) {
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
	v52 = v17
	goto L17
L17:
	;
	if base.Ui32(v44) <= base.Ui32(v52) {
		v101 = v13
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v57 = F_pushJsonbValue(m, v10+int32(28), v41, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v62 = int32(0)
	v65 = int32(0)
	goto L20
L20:
	;
	v72 = F_JsonbIteratorNext(m, v10+int32(24), v10+int32(4), int32(1))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	v97 = F_JsonbValueToJsonb(m, v65)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v95 = F_pushJsonbValue(m, v10+int32(28), v72, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L24:
	;
	if v72 != int32(3) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v72 == int32(0) {
		goto L22
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v86 = v62 + int32(1)
	if v62 == v52 {
		v62 = v86
		goto L20
	} else {
		goto L32
	}
L28:
	;
	v78 = int32(4)
	if base.Ui32(v72) < base.Ui32(v78) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = v10 + v78
	goto L31
L30:
	;
	v83 = int32(0)
	goto L31
L31:
	;
	v90 = v62
	v92 = v83
	goto L23
L32:
	;
	v90 = v86
	v92 = v10 + int32(4)
	goto L23
L33:
	;
	v62 = v90
	v65 = v95
	goto L20
L34:
	;
	v101 = v97
	goto L8
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(239483), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(514510), int32(4819), int32(30374))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errmsg(m, int32(29174), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(514510), int32(4824), int32(30374))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_delete_path(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v19 < int32(2) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v22&int32(268435456) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(239454), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(514510), int32(5003), int32(334727))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
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
					if v22&int32(268435455) == int32(0) {
						v60 = v10
						m.G0 = v7 + int32(32)
						return v60
					} else {
						F_deconstruct_array_builtin(m, v15, int32(25), v7+int32(28), v7+int32(24), v7+int32(20))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
							if v38 == int32(0) {
								v60 = v10
								m.G0 = v7 + int32(32)
								return v60
							} else {
								v43 = F_JsonbIteratorInit(m, v10+int32(4))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v43
									v48 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
									v53 = int32(0)
									v56 = F_setPath(m, v7+int32(16), v48, v49, v50, v7+int32(12), v53, v53, int32(2))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										v58 = F_JsonbValueToJsonb(m, v56)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											v60 = v58
											m.G0 = v7 + int32(32)
											return v60
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
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(125349), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(514510), int32(4998), int32(334727))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
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
func F_jsonb_each(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_each_worker_jsonb(m, l0, int32(339930), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_extract_path_text(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_get_jsonb_path_all(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v15 int32
	_ = v15
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(128)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_strlen(m, v11)
	mBase = m.M
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = v9 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v2
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v20
	v27 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = F_makeJsonLexContextCstringLen(m, v9+int32(60), v11, v12, v28, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int32(0)
	} else {
		v34 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v34)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v13
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(1326)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1327)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = int32(1328)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(1329)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1330)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(1331)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(40)
		v54 = F_pg_parse_json_or_errsave(m, v9+int32(60), v9, v13)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			if v54 != 0 {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
				v57 = F_JsonbValueToJsonb(m, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					v59 = v57
					m.G0 = v9 + int32(128)
					return v59
				}
			} else {
				v59 = v2
				m.G0 = v9 + int32(128)
				return v59
			}
		}
	}
}
func F_jsonb_in_object_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_pushJsonbValue(m, l0, int32(7), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
		return int32(0)
	}
}
func F_jsonb_insert(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(0)
				v25 = v8 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(18)
				v28 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v19 + v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - v28
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v37 < int32(2) {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
					if v40&int32(16) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(239428), int32(0))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(514510), int32(5051), int32(87557))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
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
						F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(24), v8+int32(20), v8+int32(16))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
							if v52 != 0 {
								v55 = F_JsonbIteratorInit(m, v11+int32(4))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v55
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
									v63 = int32(8)
									if v21 != 0 {
										v70 = int32(16)
									} else {
										v70 = v63
									}
									v71 = F_setPath(m, v8+int32(12), v60, v61, v62, v8+v63, int32(0), v8+int32(28), v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int32(0)
									} else {
										v73 = F_JsonbValueToJsonb(m, v71)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											v75 = v73
											m.G0 = v8 + int32(48)
											return v75
										}
									}
								}
							} else {
								v75 = v11
								m.G0 = v8 + int32(48)
								return v75
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(352845954))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(125349), int32(0))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(514510), int32(5046), int32(87557))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
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
}
func F_jsonb_int2(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
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
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v41 = int32(0)
							m.G0 = v8 + int32(32)
							return v41
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v41 = int32(0)
						m.G0 = v8 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(94972))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v35 = F_DirectFunctionCall1Coll(m, int32(1333), int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v37 {
							v41 = v35
							m.G0 = v8 + int32(32)
							return v41
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v8 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(94972))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
func F_jsonb_int8(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
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
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v41 = int32(0)
							m.G0 = v8 + int32(32)
							return v41
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v41 = int32(0)
						m.G0 = v8 + int32(32)
						return v41
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(95157))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 2:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
					v35 = F_DirectFunctionCall1Coll(m, int32(1277), int32(0), v34)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v11 == v37 {
							v41 = v35
							m.G0 = v8 + int32(32)
							return v41
						} else {
							F_pfree(m, v11)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = v35
								m.G0 = v8 + int32(32)
								return v41
							}
						}
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(95157))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
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
func F_jsonb_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v20 != int32(0))
							}
						} else {
							return base.B2i32(v20 != int32(0))
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v20 != int32(0))
						}
					} else {
						return base.B2i32(v20 != int32(0))
					}
				}
			}
		}
	}
}
func F_jsonb_object_agg_transfn_worker(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	v2 = l1
	v3 = l2
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = v12 + int32(76)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 == v4 {
		goto L10
	} else {
		goto L11
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v104
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v157
	goto L1
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L28
	} else {
		goto L114
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L28
	} else {
		goto L110
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L28
	} else {
		goto L106
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L28
	} else {
		goto L102
	}
L7:
	;
	if v45 != 0 {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v45 = v42
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
	v42 = v38
	goto L8
L10:
	;
	v34 = int32(0)
	if v15 == v34 {
		v42 = v34
		goto L8
	} else {
		goto L20
	}
L11:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	switch v20 - int32(429) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L10
	}
L12:
	;
	if v15 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	if v15 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = int32(1)
	goto L7
L15:
	;
	goto L16
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v37 = v27
	v38 = int32(1)
	goto L9
L17:
	;
	v45 = int32(2)
	goto L7
L18:
	;
	goto L19
L19:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
	v37 = v32
	v38 = int32(2)
	goto L9
L20:
	;
	v37 = v34
	v38 = v4
	goto L9
L21:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v46 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L28
	} else {
		goto L99
	}
L24:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v105 == int32(1) {
		goto L4
	} else {
		goto L38
	}
L25:
	;
	v49 = int32(4549024)
	v50 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	v55 = F_palloc(m, int32(20))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v103 = v101
	v104 = v100
	goto L24
L28:
	;
	return int32(0)
L29:
	;
	v60 = F_palloc0(m, int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v60
	v65 = F_pushJsonbValue(m, v60, int32(6), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v65
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+28)) = uint8(v3)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*uint8)(unsafe.Add(mBase, uint32(v70)+29)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v50
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v76 = F_get_fn_expr_argtype(m, v74, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	if v76 == int32(0) {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_json_categorize_type(m, v76, int32(1), v55+int32(4), v55+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v89 = F_get_fn_expr_argtype(m, v87, int32(2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v89 == int32(0) {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	F_json_categorize_type(m, v89, int32(1), v55+int32(12), v55+int32(16))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v103 = v60
	v104 = v55
	goto L24
L38:
	;
	v108 = int32(0)
	if v2 == v108 {
		v114 = v108
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v117 = v12 - int32(-64)
	v118 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v118
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v118
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	F_datum_to_jsonb_internal(m, v115, int32(0), v12+int32(56), v125, v126, int32(1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L28
	} else {
		goto L43
	}
L40:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v3 != 0 {
		v114 = v111
		goto L39
	} else {
		goto L41
	}
L41:
	;
	if v111&int32(1) != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v114 = v111
	goto L39
L43:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v131 = F_JsonbValueToJsonb(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L28
	} else {
		goto L44
	}
L44:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v133 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v137 = v136
	goto L47
L46:
	;
	v137 = v4
	goto L47
L47:
	;
	v138 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v138
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v138
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v104)+16))
	F_datum_to_jsonb_internal(m, v137, v133, v12+int32(56), v144, v145, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L28
	} else {
		goto L48
	}
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v150 = F_JsonbValueToJsonb(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L28
	} else {
		goto L49
	}
L49:
	;
	v154 = F_JsonbIteratorInit(m, v131+int32(4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	v156 = int32(4549024)
	v157 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v154
	goto L53
L51:
	;
	v244 = F_JsonbIteratorInit(m, v150+int32(4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L28
	} else {
		goto L71
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L28
	} else {
		goto L68
	}
L53:
	;
	v178 = F_JsonbIteratorNext(m, v12+int32(52), v12+int32(32), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L28
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
	v226 = F_pushJsonbValue(m, v103, int32(2), v12+int32(32))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L28
	} else {
		goto L67
	}
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v194 != int32(1) {
		goto L3
	} else {
		goto L62
	}
L56:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	if v180 != 0 {
		goto L53
	} else {
		goto L58
	}
L57:
	;
	switch v178 {
	case 0:
		goto L51
	default:
		goto L52
	case 3:
		goto L55
	case 4:
		goto L56
	case 5:
		goto L53
	}
L58:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L28
	} else {
		goto L59
	}
L59:
	;
	F_errmsg_internal(m, int32(21820), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L28
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(521213), int32(1791), int32(229561))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L28
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v200 = F_palloc(m, v197+int32(1))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L28
	} else {
		goto L63
	}
L63:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v210 = F_pg_snprintf(m, v200, v204+int32(1), int32(215163), v12+int32(16))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L28
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v200
	v216 = F_pushJsonbValue(m, v103, int32(1), v12+int32(32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L28
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v216
	if v114&int32(1) == int32(0) {
		goto L53
	} else {
		goto L66
	}
L66:
	;
	goto L54
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v226
	goto L2
L68:
	;
	F_errmsg_internal(m, int32(21820), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L28
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(521213), int32(1824), int32(229561))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L28
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v244
	v252 = int32(0)
	goto L72
L72:
	;
	v262 = F_JsonbIteratorNext(m, v12+int32(52), v12+int32(32), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L28
	} else {
		goto L79
	}
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L28
	} else {
		goto L96
	}
L74:
	;
	goto L73
L75:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	switch v284 - int32(1) {
	case 0:
		goto L87
	case 1:
		goto L86
	default:
		goto L85
	}
L76:
	;
	v281 = F_pushJsonbValue(m, v103, v262, int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L28
	} else {
		goto L84
	}
L77:
	;
	v271 = int32(1)
	if v252&v271 != 0 {
		v252 = v271
		goto L72
	} else {
		goto L82
	}
L78:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+44)))
	if v265 != 0 {
		v252 = int32(1)
		goto L72
	} else {
		goto L80
	}
L79:
	;
	switch v262 {
	case 0:
		goto L2
	case 1, 2, 3:
		goto L75
	case 4:
		goto L78
	case 5:
		goto L77
	case 6, 7:
		goto L76
	default:
		goto L74
	}
L80:
	;
	v268 = F_pushJsonbValue(m, v103, int32(4), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L28
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v268
	goto L72
L82:
	;
	v274 = int32(0)
	v277 = F_pushJsonbValue(m, v103, int32(5), v274)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L28
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v277
	v252 = v274
	goto L72
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v281
	goto L72
L85:
	;
	if v252&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v304 = F_DirectFunctionCall1Coll(m, int32(1332), int32(0), v303)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L28
	} else {
		goto L90
	}
L87:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v290 = F_palloc(m, v287+int32(1))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L28
	} else {
		goto L88
	}
L88:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	v298 = F_pg_snprintf(m, v290, v294+int32(1), int32(215163), v12)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L28
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v290
	goto L85
L90:
	;
	v306 = F_pg_detoast_datum(m, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L28
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v306
	goto L85
L92:
	;
	v313 = int32(2)
	goto L94
L93:
	;
	v313 = v262
	goto L94
L94:
	;
	v316 = F_pushJsonbValue(m, v103, v313, v12+int32(32))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L28
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v316
	goto L72
L96:
	;
	F_errmsg_internal(m, int32(384197), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L28
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(521213), int32(1883), int32(229561))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L28
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errmsg_internal(m, int32(66649), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L28
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(521213), int32(1693), int32(229561))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L28
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L28
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(386007), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L28
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(521213), int32(1718), int32(229561))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L28
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L28
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(386007), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L28
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(521213), int32(1728), int32(229561))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L28
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L28
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(315458), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L28
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(521213), int32(1744), int32(229561))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L28
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L28
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(164462), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L28
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(521213), int32(1806), int32(229561))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L28
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_object_two_arg(m *base.Module, l0 int32) int32 {
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
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v20
	v28 = F_pushJsonbValue(m, v6+int32(-40), int32(6), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if int32(1) < v19 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L44
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L40
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L36
	}
L8:
	;
	if v19 != v18 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_deconstruct_array_builtin(m, v11, int32(25), v6+int32(-4), v6+int32(-12), v6+int32(-20))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v140 = F_pushJsonbValue(m, v6+int32(-40), int32(7), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L34
	}
L13:
	;
	F_deconstruct_array_builtin(m, v16, int32(25), v6+int32(-8), v6+int32(-16), v6+int32(-24))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v51 != v52 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	if int32(0) < v51 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v57 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	F_pfree(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v57))))
	if v64 == int32(1) {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	goto L18
L21:
	;
	v68 = v57 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68+v69)))
	v72 = F_text_to_cstring(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v74 = F_strlen(m, v72)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v74
	v77 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77
	v84 = F_pushJsonbValue(m, v6+int32(-40), v77, v6+int32(-60))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87+v57))))
	if v89 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v101 = int32(0)
	goto L26
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90+v68)))
	v93 = F_text_to_cstring(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v101
	v108 = F_pushJsonbValue(m, v6+int32(-40), int32(2), v6+int32(-60))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v95 = F_strlen(m, v93)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v95
	v101 = int32(1)
	goto L26
L28:
	;
	v111 = v57 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	if v111 < v112 {
		v57 = v111
		goto L19
	} else {
		goto L29
	}
L29:
	;
	goto L20
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	F_pfree(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
	F_pfree(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
	F_pfree(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L12
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v140
	v143 = F_JsonbValueToJsonb(m, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	m.G0 = v8 - int32(-64)
	return v143
L36:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(125349), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(521213), int32(1401), int32(341022))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(154787), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(521213), int32(1412), int32(341022))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
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
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(21782), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(521213), int32(1423), int32(341022))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
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
func F_jsonb_path_exists_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_exists_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_path_match_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
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
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v23 == int32(4) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v32 = int32(1)
	v33 = int32(0)
	goto L6
L6:
	;
	v40 = F_executeJsonPath(m, v18, v33, int32(1410), int32(1411), v13, base.B2i32(v32 == int32(0)), v10+int32(8), l1)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v32 = base.B2i32(v29 != int32(0))
	v33 = v27
	goto L6
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v42 != v13 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v13)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v46 != v18 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	F_pfree(m, v18)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v50 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L29
	}
L18:
	;
	m.G0 = v10 + int32(16)
	return v74
L19:
	;
	v74 = int32(0)
	goto L18
L20:
	;
	if v32 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v53 == int32(0) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v61 = v50
	goto L23
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	switch v62 {
	case 0:
		goto L26
	default:
		goto L20
	case 3:
		goto L27
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v56 != int32(1) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v61 = v60
	goto L23
L26:
	;
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
	goto L19
L27:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+4)))
	v74 = v63
	goto L18
L28:
	;
	v69 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v69)
	goto L19
L29:
	;
	F_errcode(m, int32(135004290))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(464722), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(520780), int32(489), int32(324592))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_path_ops__extract_nodes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l2 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v9
		F_JsonbHashScalarValue(m, l2, v7+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v19 = F_palloc(m, int32(8))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(2)
				v24 = F_lappend(m, l3, v19)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v28 = v24
					m.G0 = v7 + int32(16)
					return v28
				}
			}
		}
	} else {
		v28 = l3
		m.G0 = v7 + int32(16)
		return v28
	}
}
func F_jsonb_path_query_array_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_array_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_first(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_first_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_subscript_check_subscripts(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
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
	var v31 int32
	_ = v31
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	v4 = int32(0)
	v7 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v9 <= v4 {
		v83 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L21
	} else {
		goto L25
	}
L2:
	;
	return v83
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v14 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v31 = v4
	goto L9
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v18 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != int32(23) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v25 <= int32(0) {
		v83 = v7
		goto L2
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34+v31))))
	if v36 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v83 = v76
	goto L2
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v31))))
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
	v76 = int32(1)
	v78 = v31 + v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v78 < v79 {
		v31 = v78
		goto L9
	} else {
		goto L24
	}
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
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
	v53 = v31 << (uint(int32(2)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57+v53)))
	if v59 == int32(23) {
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
	v64 = F_DirectFunctionCall1Coll(m, int32(1314), int32(0), v56)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v70 = v56
	goto L20
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v71+v53))) = v70
	goto L13
L21:
	;
	return int32(0)
L22:
	;
	v68 = F_cstring_to_text(m, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v70 = v68
	goto L20
L24:
	;
	goto L10
L25:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(315088), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(514851), int32(204), int32(125255))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
func F_jsonb_to_record(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = int32(0)
	v6 = F_populate_record_worker(m, l0, int32(437327), v3, v3, v3)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_jsonb_to_tsvector_byid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = l0 + int32(28)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = l0 + int32(36)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		v23 = F_pg_detoast_datum(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_parse_jsonb_index_flags(m, v23)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v12
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v28
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v10 + int32(8)
				F_iterate_jsonb_values(m, v16, v25, v10+int32(24))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v41 = F_make_tsvector(m, v10+int32(8))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v43 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
								if v47 != v23 {
									F_pfree(m, v23)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(32)
										return v41
									}
								} else {
									m.G0 = v10 + int32(32)
									return v41
								}
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
							if v47 != v23 {
								F_pfree(m, v23)
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									m.G0 = v10 + int32(32)
									return v41
								}
							} else {
								m.G0 = v10 + int32(32)
								return v41
							}
						}
					}
				}
			}
		}
	}
}
