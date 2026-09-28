package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogLogicalInvalidations(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[0]))
	if v10 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v19 = v13 - v14 + (v16 - v17)
		if v19 <= int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v19
			F_XLogBeginInsert(m)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_XLogRegisterData(m, v7+int32(12), int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v32 = v30 - v31
					if int32(0) < v32 {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[1]))
						v37 = int32(4)
						F_XLogRegisterData(m, v36+v31<<(uint(v37)%32), v32<<(uint(v37)%32))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v46 = v44 - v45
							if int32(0) < v46 {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[2]))
								v51 = int32(4)
								F_XLogRegisterData(m, v50+v45<<(uint(v51)%32), v46<<(uint(v51)%32))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v60 = F_XLogInsert(m, int32(1), int32(96))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								v60 = F_XLogInsert(m, int32(1), int32(96))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v46 = v44 - v45
						if int32(0) < v46 {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[2]))
							v51 = int32(4)
							F_XLogRegisterData(m, v50+v45<<(uint(v51)%32), v46<<(uint(v51)%32))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v60 = F_XLogInsert(m, int32(1), int32(96))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v60 = F_XLogInsert(m, int32(1), int32(96))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_check_log_extension_options(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int64
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v329 float64
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v655 int32
	_ = v655
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v674 int32
	_ = v674
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v725 int32
	_ = v725
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v17 + int32(32)
	return v725
L2:
	;
	v725 = int32(1)
	goto L1
L3:
	;
	v23 = F_strlen(m, v19)
	mBase = m.M
	v26 = F_guc_malloc(m, v23+int32(101))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v20 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L2
L7:
	;
	goto L6
L8:
	;
	return int32(0)
L9:
	;
	if v26 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v725 = int32(0)
	goto L1
L11:
	;
	goto L12
L12:
	;
	v35 = v23 + int32(1)
	v43 = v26
	v45 = int32(8)
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
	goto L2
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
	goto L13
L15:
	;
	v52 = v43 + int32(4)
	v55 = v52 + v45*int32(12)
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v560 = int32(0)
	goto L130
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v55, v56, v35)
	goto L19
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v55
	v59 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55))))
	goto L20
L20:
	;
	if base.B2i32(v59 == int32(32))|base.B2i32(base.Ui32((v59-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = v55
	goto L24
L22:
	;
	v99 = v55
	goto L23
L23:
	;
	v110 = int32(0)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v111 == v110 {
		goto L14
	} else {
		goto L28
	}
L24:
	;
	v84 = v72 + int32(1)
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72)+1)))
	goto L26
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v84
	v99 = v84
	goto L23
L26:
	;
	if base.B2i32(v85 == int32(32))|base.B2i32(base.Ui32((v85-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v72 = v84
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v116 = v110
	goto L32
L29:
	;
	goto L16
L30:
	;
	F_bms_free(m, v43)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L8
	} else {
		goto L127
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(-1)
	v539 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[1])) = v539
	goto L124
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(0)
	v130 = int32(_a_F_check_log_extension_options_0)
	v136 = F_scan_identifier(m, v17+int32(24), v17+int32(28), int32(44))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v515
	if v45 <= v116 {
		goto L30
	} else {
		goto L123
	}
L34:
	;
	if v136 == int32(0) {
		v523 = v130
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v140 == v136 {
		v523 = v130
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v143 = int32(*(*int8)(unsafe.Add(mBase, uint32(v142))))
	goto L37
L37:
	;
	if base.B2i32(v143 == int32(32))|base.B2i32(base.Ui32((v143-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	goto L41
L39:
	;
	goto L40
L40:
	;
	v195 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	switch v198 - int32(39) {
	case 0:
		goto L48
	case 1, 2, 3, 4:
		goto L46
	case 5:
		v356 = v195
		v361 = v195
		goto L45
	default:
		goto L47
	}
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v167 + int32(1)
	v171 = int32(*(*int8)(unsafe.Add(mBase, uint32(v167)+1)))
	goto L43
L42:
	;
	goto L40
L43:
	;
	if base.B2i32(v171 == int32(32))|base.B2i32(base.Ui32((v171-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v369 = int32(*(*int8)(unsafe.Add(mBase, uint32(v368))))
	goto L94
L46:
	;
	if base.Ui32(int32(10)) <= base.Ui32((v198-int32(48))&int32(255)) {
		goto L70
	} else {
		goto L71
	}
L47:
	;
	if v198 == int32(0) {
		v356 = v195
		v361 = v195
		goto L45
	} else {
		goto L68
	}
L48:
	;
	v202 = v197 + int32(1)
	v203 = int32(39)
	v204 = F___strchrnul(m, v202, v203)
	mBase = m.M
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204))))
	if v206 == v203 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v210
	if v210 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v210 = v204
	goto L52
L51:
	;
	v210 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v523 = int32(_a_F_check_log_extension_options_1)
	goto L31
L54:
	;
	goto L55
L55:
	;
	v218 = v210
	goto L57
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v230
	v356 = int32(476)
	v361 = v202
	goto L45
L57:
	;
	v230 = v218 + int32(1)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v231 != int32(39) {
		goto L56
	} else {
		goto L59
	}
L58:
	;
	v523 = int32(_a_F_check_log_extension_options_1)
	goto L31
L59:
	;
	v234 = F_strlen(m, v218)
	mBase = m.M
	if v234 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	base.MemoryCopy(m, v218, v230, v234)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v236
	v240 = int32(39)
	v241 = F___strchrnul(m, v236+int32(1), v240)
	mBase = m.M
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v243 == v240 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v247
	if v247 != 0 {
		v218 = v247
		goto L57
	} else {
		goto L67
	}
L64:
	;
	v247 = v241
	goto L66
L65:
	;
	v247 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	goto L58
L68:
	;
	goto L46
L69:
	;
	v351 = F_scan_identifier(m, v17+int32(20), v17+int32(28), int32(44))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L8
	} else {
		goto L92
	}
L70:
	;
	switch v198 - int32(43) {
	case 0, 2:
		goto L73
	default:
		goto L69
	}
L71:
	;
	goto L72
L72:
	;
	v271 = v197
	v272 = v198
	goto L75
L73:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	if base.Ui32(int32(9)) < base.Ui32((v262-int32(48))&int32(255)) {
		goto L69
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v284 = v272 & int32(255)
	v285 = int32(0)
	if base.B2i32(v284 == v285)|base.B2i32(v284 == int32(44)) == v285 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v342 = v302 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v342
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+1)))
	v271 = v342
	v272 = v344
	goto L75
L78:
	;
	v292 = base.I32_extend8_s(v272)
	goto L81
L79:
	;
	v305 = v271
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v305
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	v309 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v305))) = uint8(v309)
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[0])) = v309
	v318 = F_strtox_2(m, v197, v17+int32(16), v309, int64(2147483648))
	mBase = m.M
	goto L83
L81:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if base.B2i32(v292 == int32(32))|base.B2i32(base.Ui32((v292-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	v305 = v302
	goto L80
L83:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[0]))
	if v321 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v308)
	v523 = int32(_a_F_check_log_extension_options_2)
	goto L31
L85:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v336))) = uint8(v308)
	v356 = v335
	v361 = v197
	goto L45
L86:
	;
	v329 = F_strtod(m, v197, v17+int32(16))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L8
	} else {
		goto L90
	}
L87:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v322 == v197 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v322))))
	if v324 != 0 {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v335 = int32(473)
	goto L85
L90:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v332 != 0 {
		goto L84
	} else {
		goto L91
	}
L91:
	;
	v335 = int32(474)
	goto L85
L92:
	;
	if v351 != 0 {
		v356 = int32(476)
		v361 = v351
		goto L45
	} else {
		goto L93
	}
L93:
	;
	v523 = int32(_a_F_check_log_extension_options_3)
	goto L31
L94:
	;
	if base.B2i32(v369 == int32(32))|base.B2i32(base.Ui32((v369-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	goto L98
L96:
	;
	goto L97
L97:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421))))
	if v422 == int32(44) {
		goto L103
	} else {
		goto L104
	}
L98:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v393 + int32(1)
	v397 = int32(*(*int8)(unsafe.Add(mBase, uint32(v393)+1)))
	goto L100
L99:
	;
	goto L97
L100:
	;
	if base.B2i32(v397 == int32(32))|base.B2i32(base.Ui32((v397-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v501 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v500))) = uint8(v501)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v503 != 0 {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v421 + int32(1)
	v428 = int32(*(*int8)(unsafe.Add(mBase, uint32(v421)+1)))
	goto L106
L104:
	;
	goto L105
L105:
	;
	if v422 == int32(0) {
		goto L102
	} else {
		goto L115
	}
L106:
	;
	if base.B2i32(v428 == int32(32))|base.B2i32(base.Ui32((v428-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	goto L110
L108:
	;
	goto L109
L109:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v481 != 0 {
		goto L102
	} else {
		goto L114
	}
L110:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v452 + int32(1)
	v456 = int32(*(*int8)(unsafe.Add(mBase, uint32(v452)+1)))
	goto L112
L111:
	;
	goto L109
L112:
	;
	if base.B2i32(v456 == int32(32))|base.B2i32(base.Ui32((v456-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v523 = int32(_a_F_check_log_extension_options_4)
	goto L31
L115:
	;
	v523 = int32(_a_F_check_log_extension_options_5)
	goto L31
L116:
	;
	v504 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v503))) = uint8(v504)
	goto L118
L117:
	;
	goto L118
L118:
	;
	if base.Ui32(v116) < base.Ui32(v45) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v509 = v52 + v116*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v509)+8)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v509))) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v509)+4)) = v361
	goto L121
L120:
	;
	goto L121
L121:
	;
	v515 = v116 + int32(1)
	if v422 == int32(44) {
		v116 = v515
		goto L32
	} else {
		goto L122
	}
L122:
	;
	goto L33
L123:
	;
	goto L29
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v523
	v545 = F_format_elog_string(m, int32(_a_F_check_log_extension_options_6), v17)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[2])) = v545
	F_bms_free(m, v43)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	v725 = int32(0)
	goto L1
L127:
	;
	v556 = F_guc_malloc(m, v23+int32(5)+v515*int32(12))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	if v556 != 0 {
		v43 = v556
		v45 = v515
		goto L15
	} else {
		goto L129
	}
L129:
	;
	v725 = int32(0)
	goto L1
L130:
	;
	v574 = v52 + v560*int32(12)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v574)+4))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v574)+8))
	v578 = int32(0)
	v580 = m.G0
	v582 = v580 - int32(16)
	m.G0 = v582
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[3]))
	if v585 <= v578 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	F_bms_free(m, v43)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L8
	} else {
		goto L155
	}
L132:
	;
	m.G0 = v582 + int32(16)
	if v674 != 0 {
		goto L151
	} else {
		goto L152
	}
L133:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[1])) = v655
	goto L149
L134:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[4]))
	v593 = v578
	goto L135
L135:
	;
	v606 = v589 + v593*int32(12)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607))))
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
	if base.B2i32(v610 == int32(0))|base.B2i32(v610 != v613) != 0 {
		v631 = v610
		v632 = v613
		goto L138
	} else {
		goto L139
	}
L136:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v606)+8))
	v638 = m.T0[v637].(func(*base.Module, int32, int32, int32) int32)(m, v575, v576, v577)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L8
	} else {
		goto L148
	}
L137:
	;
	if v631-v632 != 0 {
		goto L144
	} else {
		goto L145
	}
L138:
	;
	goto L137
L139:
	;
	v616 = v607
	v617 = v575
	goto L140
L140:
	;
	v620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v617)+1)))
	v621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v616)+1)))
	if v621 == int32(0) {
		v631 = v621
		v632 = v620
		goto L138
	} else {
		goto L142
	}
L141:
	;
	v631 = v621
	v632 = v620
	goto L138
L142:
	;
	v624 = int32(1)
	if v621 == v620 {
		v616 = v616 + v624
		v617 = v617 + v624
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v635 = v593 + int32(1)
	if v585 != v635 {
		v593 = v635
		goto L135
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	goto L136
L147:
	;
	goto L133
L148:
	;
	v674 = v638
	goto L132
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+4)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v582))) = int32(_a_F_check_log_extension_options_7)
	v663 = F_format_elog_string(m, int32(_a_F_check_log_extension_options_8), v582)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L8
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_extension_options[5])) = v663
	v674 = v578
	goto L132
L151:
	;
	v684 = v560 + int32(1)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v684 < v685 {
		v560 = v684
		goto L130
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	goto L131
L154:
	;
	goto L13
L155:
	;
	v725 = int32(0)
	goto L1
}
func F_log(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v85 int32
	_ = v85
	var v92 float64
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v113 float64
	_ = v113
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v156 float64
	_ = v156
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	var v179 float64
	_ = v179
	v11 = base.I64_reinterpret_f64(l0)
	if base.Ui64(v11-int64(4606619468846596096)) <= base.Ui64(int64(854320534781951)) {
		if v11 == int64(4607182418800017408) {
			return float64(0)
		} else {
			v21 = base.F64_add(l0, float64(-1))
			v23 = base.F64_mul(v21, float64(1.34217728e+08))
			v25 = base.F64_sub(base.F64_add(v21, v23), v23)
			v28 = *(*float64)(unsafe.Add(mBase, _c_F_log[0]))
			v29 = base.F64_mul(base.F64_mul(v25, v25), v28)
			v30 = base.F64_add(v21, v29)
			v31 = base.F64_mul(v21, v21)
			v32 = base.F64_mul(v21, v31)
			v34 = *(*float64)(unsafe.Add(mBase, _c_F_log[1]))
			v37 = *(*float64)(unsafe.Add(mBase, _c_F_log[2]))
			v40 = *(*float64)(unsafe.Add(mBase, _c_F_log[3]))
			v43 = *(*float64)(unsafe.Add(mBase, _c_F_log[4]))
			v49 = *(*float64)(unsafe.Add(mBase, _c_F_log[5]))
			v52 = *(*float64)(unsafe.Add(mBase, _c_F_log[6]))
			v55 = *(*float64)(unsafe.Add(mBase, _c_F_log[7]))
			v61 = *(*float64)(unsafe.Add(mBase, _c_F_log[8]))
			v64 = *(*float64)(unsafe.Add(mBase, _c_F_log[9]))
			v67 = *(*float64)(unsafe.Add(mBase, _c_F_log[10]))
			return base.F64_add(v30, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, v34), base.F64_add(base.F64_mul(v31, v37), base.F64_add(base.F64_mul(v21, v40), v43)))), base.F64_add(base.F64_mul(v31, v49), base.F64_add(base.F64_mul(v21, v52), v55)))), base.F64_add(base.F64_mul(v31, v61), base.F64_add(base.F64_mul(v21, v64), v67)))), base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(v21, v25), v28), base.F64_add(v21, v25)), base.F64_add(v29, base.F64_sub(v21, v30)))))
		}
	} else {
		v85 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(48)) % 64)))
		if base.Ui32(v85-int32(_a_F_log_0)) <= base.Ui32(int32(-32737)) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v92 = float64(-1)
				v94 = m.G0
				*(*float64)(unsafe.Add(mBase, uint32(v94-int32(16))+8)) = v92
				return base.F64_div(v92, float64(0))
			} else {
				if v11 == int64(9218868437227405312) {
					v179 = l0
					return v179
				} else {
					v104 = int32(_a_F_log_0)
					if base.B2i32(v85&v104 != v104)&base.B2i32(base.Ui32(v85) <= base.Ui32(int32(_a_F_log_1))) == int32(0) {
						v113 = base.F64_sub(l0, l0)
						return base.F64_div(v113, v113)
					} else {
						v121 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15))) - int64(234187180623265792)
						v123 = v121 - int64(4604367669032910848)
						v126 = base.F64_convert_i64_s(v123 >> (uint(int64(52)) % 64))
						v128 = *(*float64)(unsafe.Add(mBase, _c_F_log[11]))
						v136 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
						v137 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[12])))
						v138 = base.F64_add(base.F64_mul(v126, v128), v137)
						v139 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[13])))
						v144 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[14])))
						v146 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[15])))
						v148 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v144), v146))
						v149 = base.F64_add(v138, v148)
						v150 = base.F64_mul(v148, v148)
						v153 = *(*float64)(unsafe.Add(mBase, _c_F_log[16]))
						v156 = *(*float64)(unsafe.Add(mBase, _c_F_log[17]))
						v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[18]))
						v163 = *(*float64)(unsafe.Add(mBase, _c_F_log[19]))
						v168 = *(*float64)(unsafe.Add(mBase, _c_F_log[20]))
						v171 = *(*float64)(unsafe.Add(mBase, _c_F_log[21]))
						v179 = base.F64_add(v149, base.F64_add(base.F64_mul(base.F64_mul(v148, v150), base.F64_add(base.F64_mul(v150, base.F64_add(base.F64_mul(v148, v153), v156)), base.F64_add(base.F64_mul(v148, v160), v163))), base.F64_add(base.F64_mul(v150, v168), base.F64_add(base.F64_mul(v126, v171), base.F64_add(v148, base.F64_sub(v138, v149))))))
						return v179
					}
				}
			}
		} else {
			v121 = v11
			v123 = v121 - int64(4604367669032910848)
			v126 = base.F64_convert_i64_s(v123 >> (uint(int64(52)) % 64))
			v128 = *(*float64)(unsafe.Add(mBase, _c_F_log[11]))
			v136 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
			v137 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[12])))
			v138 = base.F64_add(base.F64_mul(v126, v128), v137)
			v139 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[13])))
			v144 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[14])))
			v146 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[15])))
			v148 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v144), v146))
			v149 = base.F64_add(v138, v148)
			v150 = base.F64_mul(v148, v148)
			v153 = *(*float64)(unsafe.Add(mBase, _c_F_log[16]))
			v156 = *(*float64)(unsafe.Add(mBase, _c_F_log[17]))
			v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[18]))
			v163 = *(*float64)(unsafe.Add(mBase, _c_F_log[19]))
			v168 = *(*float64)(unsafe.Add(mBase, _c_F_log[20]))
			v171 = *(*float64)(unsafe.Add(mBase, _c_F_log[21]))
			v179 = base.F64_add(v149, base.F64_add(base.F64_mul(base.F64_mul(v148, v150), base.F64_add(base.F64_mul(v150, base.F64_add(base.F64_mul(v148, v153), v156)), base.F64_add(base.F64_mul(v148, v160), v163))), base.F64_add(base.F64_mul(v150, v168), base.F64_add(base.F64_mul(v126, v171), base.F64_add(v148, base.F64_sub(v138, v149))))))
			return v179
		}
	}
}
func F_log_disconnections(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_log_disconnections[0]))
	v18 = *(*int64)(unsafe.Add(mBase, _c_F_log_disconnections[1]))
	v22 = m.G0
	v23 = int32(16)
	v24 = v22 - v23
	m.G0 = v24
	F_gettimeofday(m, v24)
	mBase = m.M
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
	m.G0 = v24 + v23
	v43 = v28 + v27*int64(1000000) - int64(946684800000000) - v18
	if v43 <= int64(0) {
		v55 = int32(0)
		v56 = int32(0)
	} else {
		v47 = int64(1000000)
		v48 = base.I64_div_u_s(v43, v47)
		v55 = base.I32_wrap_i64(v48)
		v56 = base.I32_wrap_i64(v43 - v48*v47)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(44)))) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)))) = v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v60 = int32(3600)
	v61 = base.I32_div_s(v59, v60)
	v64 = v59 - v61*v60
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v69 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		return
	} else {
		if v69 != 0 {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)+292))
			v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v16)+364))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v16)+360))
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v16)+276))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v71
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v75
			*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v74
			*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v73
			if v72 != 0 {
				v82 = int32(_a_F_log_disconnections_0)
			} else {
				v82 = int32(_a_F_log_disconnections_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v82
			v85 = base.I32_div_s(v66, int32(1000))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v85
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v61
			v89 = int32(60)
			v90 = base.I32_div_s(base.I32_extend16_s(v64), v89)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = base.I32_extend16_s(v90)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = base.I32_extend16_s(v64 - v90*v89)
			F_errmsg(m, int32(_a_F_log_disconnections_2), v13)
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_log_disconnections_3), int32(_a_F_log_disconnections_4), int32(_a_F_log_disconnections_5))
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return
				} else {
					m.G0 = v13 + int32(48)
					return
				}
			}
		} else {
			m.G0 = v13 + int32(48)
			return
		}
	}
}
func F_log_newpage_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0]))
	if v18 <= int32(31) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(_a_F_log_newpage_range_0)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[1]))
	v25 = F_repalloc(m, v23, int32(_a_F_log_newpage_range_1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[2]))
	if v78 <= int32(19) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[1])) = v25
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0]))
	v32 = int32(_a_F_log_newpage_range_2)
	v33 = (int32(32) - v30) * v32
	v35 = v30 * v32
	v36 = v25 + v35
	if v36&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v33)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0])) = int32(32)
	goto L3
L7:
	;
	if v30 == int32(32) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v61 = v33
	goto L9
L9:
	;
	if v61 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v50 = v25 + int32(_a_F_log_newpage_range_1)
	v53 = v35 + v25 + int32(4)
	if base.Ui32(v53) < base.Ui32(v50) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v50
	goto L13
L12:
	;
	v55 = v53
	goto L13
L13:
	;
	v61 = (v25^int32(-1)-v35+v55)&int32(-4) + int32(4)
	goto L9
L14:
	;
	base.MemoryFill(m, v36, int32(0), v61)
	goto L6
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[3]))
	v84 = F_repalloc(m, v82, int32(240))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[2])) = int32(20)
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[3])) = v84
	goto L17
L19:
	;
	m.G0 = v15 + int32(128)
	return
L20:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v96 = int32(9)
	goto L23
L22:
	;
	v96 = int32(1)
	goto L23
L23:
	;
	v105 = int32(0)
	goto L24
L24:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[4]))
	if v110 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(l2) <= base.Ui32(v105) {
		goto L19
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v119 = int32(0)
	v123 = v105
	goto L31
L31:
	;
	v127 = int32(0)
	v129 = F_ReadBufferExtended(m, l0, l1, v123, v127, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	if v161 == int32(0) {
		goto L19
	} else {
		goto L45
	}
L33:
	;
	F_LockBufferInternal(m, v129, int32(3))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v129 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v165 = v123 + int32(1)
	if base.B2i32(v161 <= int32(31))&base.B2i32(base.Ui32(v165) < base.Ui32(l2)) != 0 {
		v119 = v161
		v123 = v165
		goto L31
	} else {
		goto L44
	}
L36:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v151)+14)))
	if v152 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[5]))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137+(v129^int32(-1))<<(uint(int32(2))%32))))
	v151 = v143
	goto L36
L38:
	;
	goto L39
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[6]))
	v151 = v145 + v129<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v119<<(uint(int32(2))%32)))) = v129
	v161 = v119 + int32(1)
	goto L35
L41:
	;
	goto L42
L42:
	;
	F_UnlockReleaseBuffer(m, v129)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v161 = v119
	goto L35
L44:
	;
	goto L32
L45:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v172 = int32(0)
	v173 = int32(_a_F_log_newpage_range_3)
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7])) = v175 + int32(1)
	if v172 < v161 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if base.Ui32(v165) < base.Ui32(l2) {
		v105 = v165
		goto L24
	} else {
		goto L69
	}
L48:
	;
	v184 = v172
	goto L51
L49:
	;
	goto L50
L50:
	;
	v281 = F_XLogInsert(m, int32(0), int32(176))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L4
	} else {
		goto L68
	}
L51:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v15+v184<<(uint(int32(2))%32))))
	F_MarkBufferDirty(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v206 = int32(0)
	v209 = F_XLogInsert(m, v206, int32(176))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L56
	}
L53:
	;
	F_XLogRegisterBuffer(m, v184&int32(255), v196, v96)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v204 = v184 + int32(1)
	if v204 != v161 {
		v184 = v204
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v216 = v206
	goto L57
L57:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v15+v216<<(uint(int32(2))%32))))
	if v228 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v252 = int32(_a_F_log_newpage_range_3)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7])) = v254 - int32(1)
	v261 = int32(0)
	goto L64
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v246))) = base.I64_rotl(v209, int64(32))
	v249 = v216 + int32(1)
	if v249 != v161 {
		v216 = v249
		goto L57
	} else {
		goto L63
	}
L60:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[5]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232+(v228^int32(-1))<<(uint(int32(2))%32))))
	v246 = v238
	goto L59
L61:
	;
	goto L62
L62:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[6]))
	v246 = v240 + v228<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	goto L58
L64:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v15+v261<<(uint(int32(2))%32))))
	F_UnlockReleaseBuffer(m, v273)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L47
L66:
	;
	v277 = v261 + int32(1)
	if v277 != v161 {
		v261 = v277
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v283 = int32(_a_F_log_newpage_range_3)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7])) = v285 - int32(1)
	goto L47
L69:
	;
	goto L25
}
