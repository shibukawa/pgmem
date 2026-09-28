package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleSetHintBits(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	if l3 == int32(0) {
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
		F_BufferSetHintBits16(m, l0+int32(20), v42|l2, l1)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			return
		}
	} else {
		if int32(0) <= l1 {
			v11 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleSetHintBits[0]))
			v17 = int64(0)
			v20 = base.AtomicRmwCmpxchg64(m, v11+l1*int32(56)-int32(32), int32(0), v17, v17)
			v27 = base.I32_wrap_i64(int64(base.Ui64(v20&int64(2147483648)) >> (uint(int64(31)) % 64)))
		} else {
			v27 = int32(0)
		}
		if v27 == int32(0) {
			v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
			F_BufferSetHintBits16(m, l0+int32(20), v42|l2, l1)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				return
			}
		} else {
			v30 = F_TransactionIdGetCommitLSN(m, l3)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = F_XLogNeedsFlush(m, v30)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
						F_BufferSetHintBits16(m, l0+int32(20), v42|l2, l1)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							return
						}
					} else {
						v36 = F_BufferGetLSNAtomic(m, l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							if base.Ui64(v36) < base.Ui64(v30) {
								return
							} else {
								v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
								F_BufferSetHintBits16(m, l0+int32(20), v42|l2, l1)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_heap_deform_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v299 int64
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v323 int32
	_ = v323
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v384 int64
	_ = v384
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v387 int64
	_ = v387
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v456 int64
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v479 int32
	_ = v479
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v502 int32
	_ = v502
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+18)))
	v26 = v24 & int32(2047)
	if v22 < v26 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v22
	goto L3
L2:
	;
	v28 = v26
	goto L3
L3:
	;
	if v21 < v28 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v21
	goto L6
L5:
	;
	v30 = v28
	goto L6
L6:
	;
	v32 = v23 + int32(23)
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+20)))
	if v33&int32(1) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
	v110 = v23 + v109
	if v93 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L8:
	;
	v93 = v30
	v98 = v28
	goto L7
L9:
	;
	goto L10
L10:
	;
	v39 = v28 >> (uint(int32(3)) % 32)
	if v39 <= int32(0) {
		v69 = v5
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+v32))))
	v88 = base.I32_ctz(v82^int32(-1)) + v69<<(uint(int32(3))%32)
	if v88 < v28 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v46 = v5
	goto L13
L13:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v32))))
	if v59 != int32(255) {
		v69 = v46
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v69 = v39
	goto L11
L15:
	;
	v63 = v46 + int32(1)
	if v63 != v39 {
		v46 = v63
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v90 = v88
	goto L19
L18:
	;
	v90 = v28
	goto L19
L19:
	;
	if v30 < v90 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = v30
	goto L22
L21:
	;
	v92 = v90
	goto L22
L22:
	;
	v93 = v92
	v98 = v90
	goto L7
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v180
	if v164 < v98 {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	v113 = int32(0)
	v164 = v113
	v180 = v113
	goto L23
L25:
	;
	goto L26
L26:
	;
	v122 = int32(0)
	goto L27
L27:
	;
	v135 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3+v122))) = uint8(v135)
	v138 = v122 << (uint(int32(3)) % 32)
	v139 = l1 + int32(28) + v138
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v139))))
	v141 = v110 + v140
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+4)))
	if v143 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v162 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1+v93<<(uint(int32(3))%32))+22)))
	v164 = v93
	v180 = v162 + v140
	goto L23
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v138))) = v154
	v157 = v122 + int32(1)
	if v157 != v93 {
		v122 = v157
		goto L27
	} else {
		goto L37
	}
L30:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+2)))
	switch v146 - int32(1) {
	case 0:
		goto L34
	case 1:
		goto L35
	default:
		goto L33
	case 3:
		goto L36
	}
L31:
	;
	goto L32
L32:
	;
	v154 = base.I64_extend_i32_u(v141)
	goto L29
L33:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	v154 = v152
	goto L29
L34:
	;
	v151 = int64(*(*int8)(unsafe.Add(mBase, uint32(v141))))
	v154 = v151
	goto L29
L35:
	;
	v150 = int64(*(*int16)(unsafe.Add(mBase, uint32(v141))))
	v154 = v150
	goto L29
L36:
	;
	v149 = int64(*(*int32)(unsafe.Add(mBase, uint32(v141))))
	v154 = v149
	goto L29
L37:
	;
	goto L28
L38:
	;
	v185 = v164
	goto L41
L39:
	;
	v304 = v164
	goto L40
L40:
	;
	if v304 < v28 {
		goto L72
	} else {
		goto L73
	}
L41:
	;
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185+l3))) = uint8(v202)
	v205 = v185 << (uint(int32(3)) % 32)
	v208 = v19 + int32(12)
	v209 = v205 + (l1 + int32(28))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
	v211 = int32(*(*int16)(unsafe.Add(mBase, uint32(v209)+2)))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
	if v202 < v211 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v304 = v98
	goto L40
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v205))) = v299
	v302 = v185 + int32(1)
	if v302 != v98 {
		v185 = v302
		goto L41
	} else {
		goto L71
	}
L44:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v221 = (v212 + v215 - int32(1)) & (int32(0) - v212)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v221 + v211
	v224 = v110 + v221
	if v210 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v211 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	switch v211 - int32(1) {
	case 0:
		goto L53
	case 1:
		goto L52
	default:
		goto L50
	case 3:
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v299 = base.I64_extend_i32_u(v224)
	goto L43
L50:
	;
	v230 = *(*int64)(unsafe.Add(mBase, uint32(v224)))
	v299 = v230
	goto L43
L51:
	;
	v229 = int64(*(*int32)(unsafe.Add(mBase, uint32(v224))))
	v299 = v229
	goto L43
L52:
	;
	v228 = int64(*(*int16)(unsafe.Add(mBase, uint32(v224))))
	v299 = v228
	goto L43
L53:
	;
	v227 = int64(*(*int8)(unsafe.Add(mBase, uint32(v224))))
	v299 = v227
	goto L43
L54:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v232))))
	if v236&int32(1) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v282 = int32(1)
	v286 = (v232 + v212 - v282) & (int32(0) - v212)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v286
	v288 = v110 + v286
	v289 = F_strlen(m, v288)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v289 + v286 + v282
	v299 = base.I64_extend_i32_u(v288)
	goto L43
L57:
	;
	v246 = (v232 + v212 - int32(1)) & (int32(0) - v212)
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v246
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v246))))
	v250 = v246
	v251 = v249
	goto L59
L58:
	;
	v250 = v232
	v251 = v236
	goto L59
L59:
	;
	v252 = v110 + v250
	if v251 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v277 + v250
	v299 = base.I64_extend_i32_u(v252)
	goto L43
L61:
	;
	v256 = int32(18)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+1)))
	if v258 == v256 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v269 = int32(1)
	if v251&v269 != 0 {
		v277 = int32(base.Ui32(v251) >> (uint(v269) % 32))
		goto L60
	} else {
		goto L70
	}
L64:
	;
	v261 = v256
	goto L66
L65:
	;
	v261 = int32(2)
	goto L66
L66:
	;
	if base.Ui32((v258-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v268 = int32(6)
	goto L69
L68:
	;
	v268 = v261
	goto L69
L69:
	;
	v277 = v268
	goto L60
L70:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v277 = int32(base.Ui32(v273) >> (uint(int32(2)) % 32))
	goto L60
L71:
	;
	goto L42
L72:
	;
	v323 = v304
	goto L75
L73:
	;
	v462 = v304
	goto L74
L74:
	;
	if v462 < v22 {
		goto L110
	} else {
		goto L111
	}
L75:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v323>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v342)>>(uint(v323&int32(7))%32))&int32(1) == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v462 = v28
	goto L74
L77:
	;
	v460 = v323 + int32(1)
	if v460 != v28 {
		v323 = v460
		goto L75
	} else {
		goto L109
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v323<<(uint(int32(3))%32)))) = int64(0)
	v356 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+l3))) = uint8(v356)
	goto L77
L79:
	;
	goto L80
L80:
	;
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+l3))) = uint8(v359)
	v362 = v323 << (uint(int32(3)) % 32)
	v365 = v19 + int32(12)
	v366 = l1 + int32(28) + v362
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+4)))
	v368 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366)+2)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+5)))
	if v359 < v368 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v362))) = v456
	goto L77
L82:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	v378 = (v369 + v372 - int32(1)) & (int32(0) - v369)
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v378 + v368
	v381 = v110 + v378
	if v367 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	if v368 == int32(-1) {
		goto L92
	} else {
		goto L93
	}
L85:
	;
	switch v368 - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	default:
		goto L88
	case 3:
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v456 = base.I64_extend_i32_u(v381)
	goto L81
L88:
	;
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v381)))
	v456 = v387
	goto L81
L89:
	;
	v386 = int64(*(*int32)(unsafe.Add(mBase, uint32(v381))))
	v456 = v386
	goto L81
L90:
	;
	v385 = int64(*(*int16)(unsafe.Add(mBase, uint32(v381))))
	v456 = v385
	goto L81
L91:
	;
	v384 = int64(*(*int8)(unsafe.Add(mBase, uint32(v381))))
	v456 = v384
	goto L81
L92:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v389))))
	if v393&int32(1) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v439 = int32(1)
	v443 = (v389 + v369 - v439) & (int32(0) - v369)
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v443
	v445 = v110 + v443
	v446 = F_strlen(m, v445)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v446 + v443 + v439
	v456 = base.I64_extend_i32_u(v445)
	goto L81
L95:
	;
	v403 = (v389 + v369 - int32(1)) & (int32(0) - v369)
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v403
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v403))))
	v407 = v403
	v408 = v406
	goto L97
L96:
	;
	v407 = v389
	v408 = v393
	goto L97
L97:
	;
	v409 = v110 + v407
	if v408 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v365))) = v434 + v407
	v456 = base.I64_extend_i32_u(v409)
	goto L81
L99:
	;
	v413 = int32(18)
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409)+1)))
	if v415 == v413 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v426 = int32(1)
	if v408&v426 != 0 {
		v434 = int32(base.Ui32(v408) >> (uint(v426) % 32))
		goto L98
	} else {
		goto L108
	}
L102:
	;
	v418 = v413
	goto L104
L103:
	;
	v418 = int32(2)
	goto L104
L104:
	;
	if base.Ui32((v415-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v425 = int32(6)
	goto L107
L106:
	;
	v425 = v418
	goto L107
L107:
	;
	v434 = v425
	goto L98
L108:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v434 = int32(base.Ui32(v430) >> (uint(int32(2)) % 32))
	goto L98
L109:
	;
	goto L76
L110:
	;
	v479 = v462
	goto L113
L111:
	;
	goto L112
L112:
	;
	m.G0 = v19 + int32(16)
	return
L113:
	;
	v500 = v479 + int32(1)
	v501 = F_getmissingattr(m, l1, v500, v479+l3)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L112
L115:
	;
	return
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2+v479<<(uint(int32(3))%32)))) = v501
	if v500 != v22 {
		v479 = v500
		goto L113
	} else {
		goto L117
	}
L117:
	;
	goto L114
}
func F_heap_getattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14311(m, l0, l1, l2, l3, int32(_a_F_heap_getattr_2_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_heap_getattr_4(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)))
	if base.Ui32(v12&int32(2047)) <= base.Ui32(int32(2)) {
		v18 = F_getmissingattr(m, l1, int32(3), l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v80 = v18
			m.G0 = v9 + int32(16)
			return v80
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+44)))
			if int32(0) <= v30 {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v35 = v24 + v33 + v30
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
				if v36 == int32(1) {
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+46)))
					if base.I32_popcnt(v39) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
							F_errmsg_internal(m, int32(_a_F_heap_getattr_4_0), v9)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_4_1), int32(123), int32(_a_F_heap_getattr_4_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v39) {
						case 0:
							v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35))))
							v80 = v44
							m.G0 = v9 + int32(16)
							return v80
						case 1:
							v45 = int64(*(*int16)(unsafe.Add(mBase, uint32(v35))))
							v80 = v45
							m.G0 = v9 + int32(16)
							return v80
						case 2:
							v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35))))
							v80 = v46
							m.G0 = v9 + int32(16)
							return v80
						case 3:
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
							v80 = v47
							m.G0 = v9 + int32(16)
							return v80
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
								F_errmsg_internal(m, int32(_a_F_heap_getattr_4_0), v9)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_4_1), int32(123), int32(_a_F_heap_getattr_4_2))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
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
				} else {
					v80 = base.I64_extend_i32_u(v35)
					m.G0 = v9 + int32(16)
					return v80
				}
			} else {
				v63 = F_nocachegetattr(m, l0, int32(3), l1)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					v80 = v63
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		} else {
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)))
			if v65&int32(4) == int32(0) {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v70)
				v80 = int64(0)
				m.G0 = v9 + int32(16)
				return v80
			} else {
				v74 = F_nocachegetattr(m, l0, int32(3), l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					v80 = v74
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		}
	}
}
func F_heap_getattr_8(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v80 int64
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+18)))
	if base.Ui32(v12&int32(2047)) <= base.Ui32(int32(6)) {
		v18 = F_getmissingattr(m, l1, int32(7), l2)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v80 = v18
			m.G0 = v9 + int32(16)
			return v80
		}
	} else {
		v22 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v22)
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+20)))
		if v25&int32(1) == v22 {
			v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+76)))
			if int32(0) <= v30 {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v35 = v24 + v33 + v30
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+80)))
				if v36 == int32(1) {
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+78)))
					if base.I32_popcnt(v39) != int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
							F_errmsg_internal(m, int32(_a_F_heap_getattr_8_0), v9)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_8_1), int32(123), int32(_a_F_heap_getattr_8_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						switch base.I32_ctz(v39) {
						case 0:
							v44 = int64(*(*int8)(unsafe.Add(mBase, uint32(v35))))
							v80 = v44
							m.G0 = v9 + int32(16)
							return v80
						case 1:
							v45 = int64(*(*int16)(unsafe.Add(mBase, uint32(v35))))
							v80 = v45
							m.G0 = v9 + int32(16)
							return v80
						case 2:
							v46 = int64(*(*int32)(unsafe.Add(mBase, uint32(v35))))
							v80 = v46
							m.G0 = v9 + int32(16)
							return v80
						case 3:
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
							v80 = v47
							m.G0 = v9 + int32(16)
							return v80
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v39
								F_errmsg_internal(m, int32(_a_F_heap_getattr_8_0), v9)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_heap_getattr_8_1), int32(123), int32(_a_F_heap_getattr_8_2))
									mBase = m.M
									v60 = m.ExcPending
									if v60 != 0 {
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
				} else {
					v80 = base.I64_extend_i32_u(v35)
					m.G0 = v9 + int32(16)
					return v80
				}
			} else {
				v63 = F_nocachegetattr(m, l0, int32(7), l1)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int64(0)
				} else {
					v80 = v63
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		} else {
			v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+23)))
			if v65&int32(64) == int32(0) {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v70)
				v80 = int64(0)
				m.G0 = v9 + int32(16)
				return v80
			} else {
				v74 = F_nocachegetattr(m, l0, int32(7), l1)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					v80 = v74
					m.G0 = v9 + int32(16)
					return v80
				}
			}
		}
	}
}
func F_heap_getnextslot_tidrange(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	v10 = l0 + int32(72)
	v12 = l0 + int32(22)
	v14 = l0 + int32(16)
	goto L2
L1:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+272))
	if v114 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v25&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	return int32(0)
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v34 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_heapgettup_pagemode(m, l0, l1, v24, v23)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_heapgettup(m, l0, l1, v24, v23)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L4
L10:
	;
	goto L4
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	m.T0[v38].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	v48 = int32(16)
	v50 = v46 | v47<<(uint(v48)%32)
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v55 = v51 | v52<<(uint(v48)%32)
	if base.Ui32(v50) < base.Ui32(v55) {
		v66 = int32(-1)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return int32(0)
L15:
	;
	if v66 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	goto L15
L17:
	;
	if base.Ui32(v55) < base.Ui32(v50) {
		v66 = int32(1)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if base.Ui32(v60) < base.Ui32(v61) {
		v66 = int32(-1)
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v66 = base.B2i32(base.Ui32(v61) < base.Ui32(v60))
	goto L16
L20:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	m.T0[v70].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)))
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10))))
	v82 = int32(16)
	v84 = v80 | v81<<(uint(v82)%32)
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+2)))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12))))
	v89 = v85 | v86<<(uint(v82)%32)
	if base.Ui32(v84) < base.Ui32(v89) {
		v100 = int32(-1)
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if l1 != int32(-1) {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	return int32(0)
L25:
	;
	if v100 <= int32(0) {
		goto L1
	} else {
		goto L30
	}
L26:
	;
	goto L25
L27:
	;
	if base.Ui32(v89) < base.Ui32(v84) {
		v100 = int32(1)
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	if base.Ui32(v94) < base.Ui32(v95) {
		v100 = int32(-1)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v100 = base.B2i32(base.Ui32(v95) < base.Ui32(v94))
	goto L26
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	m.T0[v104].(func(*base.Module, int32))(m, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	if l1 != int32(1) {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	goto L3
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_ExecStoreBufferHeapTuple(m, l0+int32(68), l2, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L8
	} else {
		goto L39
	}
L34:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+268)))
	if v117 != int32(1) {
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v124 = v114
	goto L36
L36:
	;
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v124)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v124)+24)) = v125 + int64(1)
	goto L33
L37:
	;
	F_pgstat_assoc_relation(m, v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+272))
	v124 = v123
	goto L36
L39:
	;
	return int32(1)
}
func F_heap_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(207)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_heap_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
}
func F_heap_lock_updated_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int64
	_ = v526
	var v527 int32
	_ = v527
	var v529 int64
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v553 int64
	_ = v553
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v797 int32
	_ = v797
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	if v23 != int32(_a_F_heap_lock_updated_tuple_0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(48)
	return v840
L2:
	;
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	if v26&v27 != int32(_a_F_heap_lock_updated_tuple_1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v840 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	if l1&int32(_a_F_heap_lock_updated_tuple_2) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v100
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3))))
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+2)))
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v100
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v100)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v105)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+34)) = uint16(v104)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v103)
	v119 = F_heap_fetch(m, l0, int32(_a_F_heap_lock_updated_tuple_3), v21+int32(28), v21+int32(24), v100)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L21
	}
L8:
	;
	v90 = l2
	goto L7
L9:
	;
	goto L10
L10:
	;
	v43 = F_GetMultiXactIdMembers(m, l2, v21+int32(28), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	if v43 <= int32(0) {
		v90 = int32(0)
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v51 = int32(0)
	goto L15
L13:
	;
	F_pfree(m, v48)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L19
	}
L14:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v79 = v77
	goto L13
L15:
	;
	v69 = v48 + v51<<(uint(int32(3))%32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v70) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v79 = int32(0)
	goto L13
L17:
	;
	v74 = v51 + int32(1)
	if v74 != v43 {
		v51 = v74
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v90 = v79
	goto L7
L20:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v834 == int32(0) {
		v840 = v817
		goto L1
	} else {
		goto L180
	}
L21:
	;
	if v119 == int32(0) {
		v817 = v100
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v124 = v21 + int32(32)
	v136 = v90
	v140 = v103<<(uint(int32(16))%32) | v104
	goto L25
L23:
	;
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L5
	} else {
		goto L179
	}
L24:
	;
	v797 = int32(0)
	goto L23
L25:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	goto L29
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	goto L26
L28:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615)+21)))
	if v616&int32(8) != 0 {
		goto L24
	} else {
		goto L147
	}
L29:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[0]))
	if v172 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	switch v577 - int32(9) {
	case 0, 1, 2:
		goto L27
	case 3:
		goto L28
	default:
		v797 = v576
		goto L23
	}
L31:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v146 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L33
L35:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209)+20)))
	v212 = v210 & int32(768)
	if v136 != 0 {
		goto L48
	} else {
		goto L49
	}
L36:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+10)))
	if v187&int32(4) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v178+(v146^int32(-1))<<(uint(int32(2))%32))))
	v186 = v180
	goto L36
L38:
	;
	goto L39
L39:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	v186 = v182 + v146<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	F_LockBufferInternal(m, v146, int32(3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_visibilitymap_pin(m, l0, v140, v21+int32(12))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L46
	}
L43:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+10)))
	if v195&int32(4) == int32(0) {
		goto L35
	} else {
		goto L44
	}
L44:
	;
	F_UnlockBuffer(m, v146)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	F_LockBufferInternal(m, v146, int32(3))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	goto L35
L48:
	;
	if v212 != int32(768) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v212 != int32(768) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v217 = v215
	goto L53
L52:
	;
	v217 = int32(2)
	goto L53
L53:
	;
	if v217 != v136 {
		goto L24
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v223 = v221
	goto L57
L56:
	;
	v223 = int32(2)
	goto L57
L57:
	;
	v224 = F_TransactionIdDidAbort(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	if v224 != 0 {
		goto L24
	} else {
		goto L59
	}
L59:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+18)))
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+20)))
	if v229&int32(2048) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	if v577 == int32(5) {
		goto L29
	} else {
		goto L146
	}
L61:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_pfree(m, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L5
	} else {
		goto L145
	}
L62:
	;
	if v229&int32(_a_F_heap_lock_updated_tuple_2) != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v398 = v228
	goto L64
L64:
	;
	F_compute_new_xmax_infomask(m, v227, v229, v398&int32(_a_F_heap_lock_updated_tuple_1), l4, l5, int32(0), v21+int32(16), v21+int32(22), v21+int32(20))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L5
	} else {
		goto L114
	}
L65:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v395)+18)))
	v398 = v396
	goto L64
L66:
	;
	v241 = F_GetMultiXactIdMembers(m, v227, v21, int32(base.Ui32(v229&int32(128))>>(uint(int32(7))%32)))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v319 = int32(0)
	if base.B2i32(v229&int32(128) == v319)&base.B2i32(v229&int32(80) != int32(64)) == v319 {
		goto L92
	} else {
		goto L93
	}
L69:
	;
	if int32(0) < v241 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v247 = int32(0)
	goto L73
L71:
	;
	goto L72
L72:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v312 == int32(0) {
		goto L65
	} else {
		goto L88
	}
L73:
	;
	v264 = v247 << (uint(int32(3)) % 32)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v266 = v264 + v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v273 = F_test_lockmode_for_conflict(m, v267, v268, l5, v21+int32(28), v21+int32(11))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L75
	}
L74:
	;
	goto L72
L75:
	;
	if v273 == int32(2) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v571 = int32(12)
	goto L61
L77:
	;
	goto L78
L78:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+11)))
	if v278 == int32(1) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_UnlockBuffer(m, v146)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if v273 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v283+v264)))
	F_XactLockTableWait(m, v285, l0, v124, int32(4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v571 = int32(5)
	goto L61
L84:
	;
	v571 = int32(8)
	goto L61
L85:
	;
	goto L86
L86:
	;
	v292 = v247 + int32(1)
	if v292 != v241 {
		v247 = v292
		goto L73
	} else {
		goto L87
	}
L87:
	;
	goto L74
L88:
	;
	F_pfree(m, v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	goto L65
L90:
	;
	v364 = F_test_lockmode_for_conflict(m, v359, v227, l5, v21+int32(28), v21+int32(11))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L5
	} else {
		goto L106
	}
L91:
	;
	v359 = int32(1)
	goto L90
L92:
	;
	switch int32(base.Ui32(v229)>>(uint(int32(4))%32))&int32(5) - int32(1) {
	case 0:
		v359 = int32(0)
		goto L90
	case 1, 2:
		goto L27
	case 3:
		goto L96
	case 4:
		goto L91
	default:
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v228&int32(_a_F_heap_lock_updated_tuple_4) != 0 {
		goto L103
	} else {
		goto L104
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L5
	} else {
		goto L100
	}
L96:
	;
	if v228&int32(_a_F_heap_lock_updated_tuple_4) != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v339 = int32(3)
	goto L99
L98:
	;
	v339 = int32(2)
	goto L99
L99:
	;
	v359 = v339
	goto L90
L100:
	;
	F_errmsg_internal(m, int32(_a_F_heap_lock_updated_tuple_5), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_heap_lock_updated_tuple_6), int32(_a_F_heap_lock_updated_tuple_7), int32(_a_F_heap_lock_updated_tuple_8))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v357 = int32(5)
	goto L105
L104:
	;
	v357 = int32(4)
	goto L105
L105:
	;
	v359 = v357
	goto L90
L106:
	;
	if v364 == int32(2) {
		goto L28
	} else {
		goto L107
	}
L107:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+11)))
	if v368 == int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_UnlockBuffer(m, v146)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L5
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	if v364 != 0 {
		v797 = v364
		goto L23
	} else {
		goto L113
	}
L111:
	;
	F_XactLockTableWait(m, v227, l0, v124, int32(4))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	v576 = v364
	v577 = int32(5)
	goto L60
L113:
	;
	goto L65
L114:
	;
	v426 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+10)))
	v428 = v426 & int32(4)
	if v428 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_LockBufferInternal(m, v429, int32(3))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L5
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v434 = int32(_a_F_heap_lock_updated_tuple_9)
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3])) = v436 + int32(1)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+4)) = v441
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v443)+20)))
	v446 = v444 & int32(_a_F_heap_lock_updated_tuple_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v443)+20)) = uint16(v446)
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v443)+18)))
	v450 = v448 & int32(_a_F_heap_lock_updated_tuple_11)
	*(*uint16)(unsafe.Add(mBase, uint32(v443)+18)) = uint16(v450)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)))
	v454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)))
	v455 = v453 | v454
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+20)) = uint16(v455)
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)))
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)))
	v459 = v457 | v458
	*(*uint16)(unsafe.Add(mBase, uint32(v452)+18)) = uint16(v459)
	F_MarkBufferDirty(m, v146)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L5
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+10)))
	if v463&int32(4) != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v468 = F_visibilitymap_clear(m, v140, v466, int32(2))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L5
	} else {
		goto L123
	}
L121:
	;
	v470 = int32(0)
	goto L122
L122:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v471)+118)))
	if v472 != int32(112) {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v470 = v468
	goto L122
L124:
	;
	v560 = int32(_a_F_heap_lock_updated_tuple_9)
	v562 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[3])) = v562 - int32(1)
	if v428 == int32(0) {
		goto L28
	} else {
		goto L143
	}
L125:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[4]))
	if v476 <= int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v479 != 0 {
		goto L124
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L5
	} else {
		goto L131
	}
L129:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v480 != 0 {
		goto L124
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	F_XLogRegisterBuffer(m, int32(0), v146, int32(8))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L132
	}
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+7)) = uint8(v470)
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v441
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v489)
	v495 = int32(1)
	v497 = int32(8)
	v499 = int32(4)
	v514 = int32(base.Ui32(v457)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v453)>>(uint(v495)%32))&v497 | (int32(base.Ui32(v453)>>(uint(v499)%32))&v499 | (int32(base.Ui32(v453)>>(uint(int32(12))%32))&v495 | int32(base.Ui32(v453)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)) = uint8(v514)
	F_XLogRegisterData(m, v21, v497)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L5
	} else {
		goto L133
	}
L133:
	;
	if v470 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_XLogRegisterBuffer(m, int32(1), v520, int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L5
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v553 = F_XLogInsert(m, int32(9), int32(96))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L5
	} else {
		goto L142
	}
L137:
	;
	v526 = F_XLogInsert(m, int32(9), int32(96))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v529 = base.I64_rotl(v526, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v186))) = v529
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v531 < int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[1]))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v535+(v531^int32(-1))<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v541))) = v529
	goto L124
L140:
	;
	goto L141
L141:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _c_F_heap_lock_updated_tuple[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v544+v531<<(uint(int32(13))%32))+uint32(_c_F_heap_lock_updated_tuple[5]))) = v529
	goto L124
L142:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v186))) = base.I64_rotl(v553, int64(32))
	goto L124
L143:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_UnlockBuffer(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L5
	} else {
		goto L144
	}
L144:
	;
	goto L28
L145:
	;
	v576 = v273
	v577 = v571
	goto L60
L146:
	;
	goto L30
L147:
	;
	v619 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v615)+16)))
	if v619 == int32(_a_F_heap_lock_updated_tuple_0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v622 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v615)+12)))
	v623 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v615)+14)))
	if v622&v623 == int32(_a_F_heap_lock_updated_tuple_1) {
		goto L24
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v628 = v615 + int32(12)
	v629 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	v631 = int32(16)
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+2)))
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628))))
	if v629|v630<<(uint(v631)%32) == v634|v635<<(uint(v631)%32) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	goto L150
L152:
	;
	if v645 != 0 {
		goto L24
	} else {
		goto L158
	}
L153:
	;
	goto L152
L154:
	;
	v641 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	v642 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v628)+4)))
	if v641 == v642 {
		v645 = int32(1)
		goto L153
	} else {
		goto L157
	}
L155:
	;
	goto L156
L156:
	;
	v645 = int32(0)
	goto L153
L157:
	;
	goto L156
L158:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v648 = F_HeapTupleHeaderIsOnlyLocked(m, v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L5
	} else {
		goto L159
	}
L159:
	;
	if v648 != 0 {
		v797 = int32(0)
		goto L23
	} else {
		goto L160
	}
L160:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v650)+4))
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v650)+20)))
	if v652&int32(_a_F_heap_lock_updated_tuple_12) != int32(_a_F_heap_lock_updated_tuple_2) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v735 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v719)+12)))
	v736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v719)+14)))
	v737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v719)+16)))
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L5
	} else {
		goto L176
	}
L162:
	;
	v719 = v650
	v725 = v651
	goto L161
L163:
	;
	goto L164
L164:
	;
	v657 = int32(0)
	v659 = F_GetMultiXactIdMembers(m, v651, v21, v657)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L5
	} else {
		goto L165
	}
L165:
	;
	if int32(0) < v659 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v667 = int32(0)
	goto L171
L167:
	;
	v706 = v657
	goto L168
L168:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	v719 = v716
	v725 = v706
	goto L161
L169:
	;
	F_pfree(m, v664)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L175
	}
L170:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	v695 = v693
	goto L169
L171:
	;
	v685 = v664 + v667<<(uint(int32(3))%32)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v686) {
		goto L170
	} else {
		goto L173
	}
L172:
	;
	v695 = int32(0)
	goto L169
L173:
	;
	v690 = v667 + int32(1)
	if v690 != v659 {
		v667 = v690
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v706 = v695
	goto L168
L176:
	;
	v740 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v740
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+22)) = uint16(v740)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v737)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+34)) = uint16(v736)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+32)) = uint16(v735)
	v757 = F_heap_fetch(m, l0, int32(_a_F_heap_lock_updated_tuple_3), v21+int32(28), v21+int32(24), v740)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L5
	} else {
		goto L177
	}
L177:
	;
	if v757 != 0 {
		v136 = v725
		v140 = v736 | v735<<(uint(int32(16))%32)
		goto L25
	} else {
		goto L178
	}
L178:
	;
	v817 = v740
	goto L20
L179:
	;
	v817 = v797
	goto L20
L180:
	;
	F_ReleaseBuffer(m, v834)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L5
	} else {
		goto L181
	}
L181:
	;
	v840 = v817
	goto L1
}
func F_heap_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	v2 = l1
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v17 = v15 & int32(_a_F_heap_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v17)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		v28 = int32(0)
		if base.B2i32(base.Ui32(v21) < base.Ui32(int32(25)))|base.B2i32((v21+int32(_a_F_heap_mask_1))&int32(_a_F_heap_mask_2) == v28) == v28 {
			v34 = int32(base.Ui32(v2) >> (uint(int32(16)) % 32))
			v43 = int32(1)
			for {
				v50 = l0 + int32(20) + v43&int32(_a_F_heap_mask_3)<<(uint(int32(2))%32)
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v54 = l0 + v51&int32(_a_F_heap_mask_4)
				if v51&int32(_a_F_heap_mask_5) != int32(_a_F_heap_mask_6) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = int32(0)
					v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+20)))
					v64 = int32(768)
					if v61&v64 == v64 {
						v68 = int32(-3073)
					} else {
						v68 = int32(15)
					}
					v69 = v61 & v68
					*(*uint16)(unsafe.Add(mBase, uint32(v54)+20)) = uint16(v69)
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
					if v71 != int32(_a_F_heap_mask_7) {
					} else {
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)) = uint16(v43)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+14)) = uint16(v2)
						*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)) = uint16(v34)
					}
				}
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				v80 = int32(base.Ui32(v78) >> (uint(int32(17)) % 32))
				if v80 == int32(0) {
				} else {
					v87 = (v80+int32(7))&int32(_a_F_heap_mask_0) - v80
					v88 = int32(0)
					if base.B2i32(v87 <= v88)|base.B2i32(v87 == v88) != 0 {
					} else {
						base.MemoryFill(m, v54+v80, int32(0), v87)
					}
				}
				v98 = v43 + int32(1)
				v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
				if base.Ui32(int32(25)) <= base.Ui32(v101) {
					v109 = int32(base.Ui32(v101+int32(_a_F_heap_mask_1)) >> (uint(int32(2)) % 32))
				} else {
					v109 = int32(0)
				}
				if base.Ui32(v98&int32(_a_F_heap_mask_3)) <= base.Ui32(v109&int32(_a_F_heap_mask_3)) {
					v43 = v98
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
func F_heap_page_prune_opt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	v9 = m.G0
	v11 = v9 - int32(656)
	m.G0 = v11
	if l1 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[0])))
	if v33 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(l1^int32(-1))<<(uint(int32(2))%32))))
	v30 = v22
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[2]))
	v30 = v24 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	m.G0 = v11 + int32(656)
	return
L6:
	;
	if v43 != 0 {
		goto L5
	} else {
		goto L10
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[3]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+308))
	v41 = base.B2i32(v39 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[0])) = uint8(v41)
	v43 = v41
	goto L9
L8:
	;
	v43 = int32(0)
	goto L9
L9:
	;
	goto L6
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v44 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v47 = F_GlobalVisHorizonKindForRel(m, l0)
	mBase = m.M
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_c_F_heap_page_prune_opt[4])))
	goto L12
L12:
	;
	v51 = F_GlobalVisTestIsRemovableXid(m, v50, v44)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	if v51 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = int32(819)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v63 = base.I32_div_s(int32(_a_F_heap_page_prune_opt_0)-v58<<(uint(int32(13))%32), int32(100))
	if base.Ui32(v63) <= base.Ui32(v56) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v69 = int32(819)
	goto L18
L18:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v70&int32(2) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v66 = v56
	goto L21
L20:
	;
	v66 = v63
	goto L21
L21:
	;
	v69 = v66
	goto L18
L22:
	;
	v78 = int32(4)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v81 = v79 - v80
	if v81 <= v78 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L24
L24:
	;
	if l1 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L25:
	;
	if base.Ui32(v69) <= base.Ui32(v141) {
		goto L5
	} else {
		goto L44
	}
L26:
	;
	v84 = v78
	goto L28
L27:
	;
	v84 = v81
	goto L28
L28:
	;
	v86 = v84 - int32(4)
	if v86 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v141 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v80) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v141 = v86
	goto L25
L33:
	;
	v97 = int32(base.Ui32(v80+int32(_a_F_heap_page_prune_opt_1)) >> (uint(int32(2)) % 32))
	goto L35
L34:
	;
	v97 = int32(0)
	goto L35
L35:
	;
	if base.Ui32(v97&int32(_a_F_heap_page_prune_opt_2)) < base.Ui32(int32(291)) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v102&int32(1) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = int32(0)
	goto L25
L38:
	;
	goto L39
L39:
	;
	v111 = int32(1)
	goto L40
L40:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(20)+v111&int32(_a_F_heap_page_prune_opt_2)<<(uint(int32(2))%32))+1)))
	if v120&int32(384) == int32(0) {
		goto L32
	} else {
		goto L42
	}
L41:
	;
	v141 = int32(0)
	goto L25
L42:
	;
	v126 = v111 + int32(1)
	v127 = int32(_a_F_heap_page_prune_opt_2)
	if base.Ui32(v126&v127) <= base.Ui32(v97&v127) {
		v111 = v126
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L24
L45:
	;
	F_visibilitymap_pin(m, l0, v161, l2)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L49
	}
L46:
	;
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[5]))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146+(l1^int32(-1))*int32(56))+16))
	v161 = v152
	goto L45
L47:
	;
	goto L48
L48:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[6]))
	v155 = int32(56)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+l1*v155-v155)+16))
	v161 = v160
	goto L45
L49:
	;
	v164 = F_ConditionalLockBufferForCleanup(m, l1)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	if v164 == int32(0) {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v168&int32(2) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	F_UnlockBuffer(m, l1)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L13
	} else {
		goto L115
	}
L53:
	;
	v176 = int32(4)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v179 = v177 - v178
	if v179 <= v176 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v244 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v244
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v243
	if l3 != 0 {
		goto L76
	} else {
		goto L77
	}
L56:
	;
	if base.Ui32(v69) <= base.Ui32(v239) {
		goto L52
	} else {
		goto L75
	}
L57:
	;
	v182 = v176
	goto L59
L58:
	;
	v182 = v179
	goto L59
L59:
	;
	v184 = v182 - int32(4)
	if v184 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v239 = int32(0)
	goto L56
L61:
	;
	goto L62
L62:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v178) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v239 = v184
	goto L56
L64:
	;
	v195 = int32(base.Ui32(v178+int32(_a_F_heap_page_prune_opt_1)) >> (uint(int32(2)) % 32))
	goto L66
L65:
	;
	v195 = int32(0)
	goto L66
L66:
	;
	if base.Ui32(v195&int32(_a_F_heap_page_prune_opt_2)) < base.Ui32(int32(291)) {
		goto L63
	} else {
		goto L67
	}
L67:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v200&int32(1) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v239 = int32(0)
	goto L56
L69:
	;
	goto L70
L70:
	;
	v209 = int32(1)
	goto L71
L71:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(20)+v209&int32(_a_F_heap_page_prune_opt_2)<<(uint(int32(2))%32))+1)))
	if v218&int32(384) == int32(0) {
		goto L63
	} else {
		goto L73
	}
L72:
	;
	v239 = int32(0)
	goto L56
L73:
	;
	v224 = v209 + int32(1)
	v225 = int32(_a_F_heap_page_prune_opt_2)
	if base.Ui32(v224&v225) <= base.Ui32(v195&v225) {
		v209 = v224
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L55
L76:
	;
	v252 = int32(12)
	goto L78
L77:
	;
	v252 = int32(4)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v252
	v260 = int32(0)
	F_heap_page_prune_and_freeze(m, v11+int32(12), v11+int32(40), v11+int32(654), v260, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L13
	} else {
		goto L79
	}
L79:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v265 < v264 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v268 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)))
	if v285 != int32(1) {
		goto L52
	} else {
		goto L89
	}
L83:
	;
	goto L82
L84:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v271 != int32(1) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	v277 = v268
	goto L86
L86:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v277)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v277)+96)) = v278 - base.I64_extend_i32_s(v264-v265)
	goto L83
L87:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v277 = v276
	goto L86
L89:
	;
	v291 = int32(4)
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+14)))
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+12)))
	v294 = v292 - v293
	if v294 <= v291 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	F_UnlockBuffer(m, l1)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L13
	} else {
		goto L109
	}
L91:
	;
	v297 = v291
	goto L93
L92:
	;
	v297 = v294
	goto L93
L93:
	;
	v299 = v297 - int32(4)
	if v299 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v354 = int32(0)
	goto L90
L95:
	;
	goto L96
L96:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v293) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v354 = v299
	goto L90
L98:
	;
	v310 = int32(base.Ui32(v293+int32(_a_F_heap_page_prune_opt_1)) >> (uint(int32(2)) % 32))
	goto L100
L99:
	;
	v310 = int32(0)
	goto L100
L100:
	;
	if base.Ui32(v310&int32(_a_F_heap_page_prune_opt_2)) < base.Ui32(int32(291)) {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+10)))
	if v315&int32(1) == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v354 = int32(0)
	goto L90
L103:
	;
	goto L104
L104:
	;
	v324 = int32(1)
	goto L105
L105:
	;
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+int32(20)+v324&int32(_a_F_heap_page_prune_opt_2)<<(uint(int32(2))%32))+1)))
	if v333&int32(384) == int32(0) {
		goto L97
	} else {
		goto L107
	}
L106:
	;
	v354 = int32(0)
	goto L90
L107:
	;
	v339 = v324 + int32(1)
	v340 = int32(_a_F_heap_page_prune_opt_2)
	if base.Ui32(v339&v340) <= base.Ui32(v310&v340) {
		v324 = v339
		goto L105
	} else {
		goto L108
	}
L108:
	;
	goto L106
L109:
	;
	if l1 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	F_RecordPageWithFreeSpace(m, l0, v375, v354)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L13
	} else {
		goto L114
	}
L111:
	;
	v360 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[5]))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v360+(l1^int32(-1))*int32(56))+16))
	v375 = v366
	goto L110
L112:
	;
	goto L113
L113:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_heap_page_prune_opt[6]))
	v369 = int32(56)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368+l1*v369-v369)+16))
	v375 = v374
	goto L110
L114:
	;
	goto L5
L115:
	;
	goto L5
}
