package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_brinLockRevmapPageForUpdate(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = base.I32_div_u_s(l1, v10)
	v13 = base.I32_div_u_s(v11, int32(1360))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v13) < base.Ui32(v14) {
		v17 = v13 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v18 == int32(0) {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v51 = F_ReadBuffer(m, v50, v17)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v51
				v54 = v51
				F_LockBuffer(m, v54, int32(2))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v54
				}
			}
		} else {
			if v18 < int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v18^int32(-1))<<(uint(int32(6))%32))+16))
				v39 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v18<<(uint(int32(6))%32)+int32(-64))+16))
				v39 = v38
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v17 == v39 {
				v54 = v40
				F_LockBuffer(m, v54, int32(2))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return v54
				}
			} else {
				if v40 == int32(0) {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v51 = F_ReadBuffer(m, v50, v17)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v51
						v54 = v51
						F_LockBuffer(m, v54, int32(2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v54
						}
					}
				} else {
					F_ReleaseBuffer(m, v40)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v51 = F_ReadBuffer(m, v50, v17)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v51
							v54 = v51
							F_LockBuffer(m, v54, int32(2))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(16)
								return v54
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
			F_errmsg_internal(m, int32(48056), v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497115), int32(471), int32(225857))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
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
func F_brinRevmapExtend(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int64
	_ = v451
	var v452 int32
	_ = v452
	var v453 int64
	_ = v453
	var v458 int64
	_ = v458
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v505 int32
	_ = v505
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = base.I32_div_u_s(l1, v21)
	v24 = base.I32_div_u_s(v22, int32(1360))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v25) <= base.Ui32(v24) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L10
	} else {
		goto L109
	}
L2:
	;
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v19 + int32(48)
	return
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v48, int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v52 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v505) <= base.Ui32(v24) {
		goto L5
	} else {
		goto L108
	}
L14:
	;
	v72 = v70 + int32(36)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 != v74 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v52^int32(-1))<<(uint(int32(2))%32))))
	v70 = v62
	goto L14
L16:
	;
	goto L17
L17:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v70 = v64 + v52<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v73
	F_LockBuffer(m, v52, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v81 = v73 + int32(1)
	v83 = F_RelationGetNumberOfBlocksInFork(m, v47, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L24
	}
L21:
	;
	goto L13
L22:
	;
	F_UnlockReleaseBuffer(m, v473)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L10
	} else {
		goto L107
	}
L23:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+14)))
	if v165 != 0 {
		goto L45
	} else {
		goto L46
	}
L24:
	;
	if base.Ui32(v81) < base.Ui32(v83) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v86 = F_ReadBuffer(m, v47, v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = int64(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v47
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v19)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v113
	v117 = int32(0)
	v120 = F_ExtendBufferedRel(m, v19+int32(16), v117, v117, int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L33
	}
L28:
	;
	F_LockBuffer(m, v86, int32(2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v86 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(v86^int32(-1))<<(uint(int32(2))%32))))
	v163 = v86
	v164 = v100
	goto L23
L31:
	;
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v163 = v86
	v164 = v102 + v86<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L33:
	;
	if v120 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v140 != v81 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v120^int32(-1))<<(uint(int32(6))%32))+16))
	v140 = v131
	goto L34
L36:
	;
	goto L37
L37:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133+v120<<(uint(int32(6))%32)+int32(-64))+16))
	v140 = v139
	goto L34
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v142, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v120 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v473 = v120
	goto L22
L42:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149+(v120^int32(-1))<<(uint(int32(2))%32))))
	v163 = v120
	v164 = v155
	goto L23
L43:
	;
	goto L44
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v163 = v120
	v164 = v157 + v120<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L45:
	;
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+16)))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+v166)+6)))
	if v168 != int32(61587) {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v171 = int32(0)
	if v163 < v171 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L47
L49:
	;
	if v261 != 0 {
		goto L64
	} else {
		goto L65
	}
L50:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+14)))
	if v190 == int32(0) {
		v261 = v171
		goto L49
	} else {
		goto L54
	}
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175+(v163^int32(-1))<<(uint(int32(2))%32))))
	v189 = v181
	goto L50
L52:
	;
	goto L53
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v189 = v183 + v163<<(uint(int32(13))%32) + int32(-8192)
	goto L50
L54:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+12)))
	if base.Ui32(v193) < base.Ui32(int32(25)) {
		v261 = v171
		goto L49
	} else {
		goto L55
	}
L55:
	;
	v201 = int32(base.Ui32(v193+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v201 == int32(0) {
		v261 = v171
		goto L49
	} else {
		goto L56
	}
L56:
	;
	v212 = int32(1)
	goto L57
L57:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v212&int32(65535)<<(uint(int32(2))%32)+(v189+int32(24))-int32(3)))))
	if v230&int32(384) == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v240 = int32(1)
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+16)))
	v244 = v189 + v241 + int32(4)
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v244))))
	v247 = v245 | v240
	*(*uint16)(unsafe.Add(mBase, uint32(v244))) = uint16(v247)
	F_MarkBufferDirtyHint(m, v163, v240)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L10
	} else {
		goto L63
	}
L59:
	;
	v236 = v212 + int32(1)
	if base.Ui32(v236&int32(65535)) <= base.Ui32(v201) {
		v212 = v236
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v261 = v171
	goto L49
L63:
	;
	v261 = v240
	goto L49
L64:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v268, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v400 = int32(4514932)
	v402 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v402 + int32(1)
	v406 = int32(61586)
	F_PageInit(m, v164, int32(8192), int32(8))
	mBase = m.M
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v164+v410)+6)) = uint16(v406)
	goto L91
L67:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v273 = m.G0
	v275 = v273 - int32(16)
	m.G0 = v275
	v277 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v275)+12)) = v277
	if v163 < v277 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	F_UnlockReleaseBuffer(m, v163)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L90
	}
L69:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296)+12)))
	if base.Ui32(v297) < base.Ui32(int32(25)) {
		goto L68
	} else {
		goto L73
	}
L70:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282+(v163^int32(-1))<<(uint(int32(2))%32))))
	v296 = v288
	goto L69
L71:
	;
	goto L72
L72:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v296 = v290 + v163<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	v305 = int32(base.Ui32(v297+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v305 == int32(0) {
		goto L68
	} else {
		goto L74
	}
L74:
	;
	v318 = int32(1)
	goto L75
L75:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v330 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L68
L77:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L10
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v334 = v318 & int32(65535)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v334<<(uint(int32(2))%32)+(v296+int32(24))-int32(4))))
	if v340&int32(98304) != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v347 = int32(base.Ui32(v340) >> (uint(int32(17)) % 32))
	v351 = F_brin_copy_tuple(m, v296+v340&int32(32767), v347, int32(0), v275+int32(12))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L10
	} else {
		goto L84
	}
L82:
	;
	v373 = v318
	goto L83
L83:
	;
	v375 = v373 + int32(1)
	if base.Ui32(v375&int32(65535)) <= base.Ui32(v305) {
		v318 = v375
		goto L75
	} else {
		goto L89
	}
L84:
	;
	F_LockBuffer(m, v163, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v358 = F_brin_doupdate(m, v47, v272, l0, v356, v163, v334, v351, v347, v351, v347, int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_LockBuffer(m, v163, int32(1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296)+16)))
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296+int32(6)+v363))))
	if v365 != int32(61587) {
		goto L68
	} else {
		goto L88
	}
L88:
	;
	v373 = v318 - (v358 ^ int32(1))
	goto L83
L89:
	;
	goto L76
L90:
	;
	m.G0 = v275 + int32(16)
	goto L13
L91:
	;
	F_MarkBufferDirty(m, v163)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v415 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+12)) = uint16(v415)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v81
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MarkBufferDirty(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v421)+48))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v422)+118)))
	if v423 != int32(112) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v461 = int32(4514932)
	v463 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v463 - int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v467, int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L10
	} else {
		goto L106
	}
L95:
	;
	v427 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v427 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v421)+32))
	if v430 != 0 {
		goto L94
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v81
	F_XLogBeginInsert(m)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L101
	}
L99:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v421)+40))
	if v431 != 0 {
		goto L94
	} else {
		goto L100
	}
L100:
	;
	goto L98
L101:
	;
	F_XLogRegisterData(m, v19+int32(32), int32(4))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_XLogRegisterBuffer(m, int32(0), v441, int32(8))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	F_XLogRegisterBuffer(m, int32(1), v163, int32(6))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v451 = F_XLogInsert(m, int32(17), int32(64))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	v453 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = base.I64_rotr(v451, v453)
	*(*uint32)(unsafe.Add(mBase, uint32(v164)+4)) = uint32(v451)
	v458 = int64(base.Ui64(v451) >> (uint(v453) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v164))) = uint32(v458)
	goto L94
L106:
	;
	v473 = v163
	goto L22
L107:
	;
	goto L13
L108:
	;
	goto L6
L109:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	v533 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+16)))
	v535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164+v533)+6)))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v47)+48))
	if v163 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v536 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v535
	F_errmsg(m, int32(48199), v19)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L10
	} else {
		goto L115
	}
L112:
	;
	v540 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v540+(v163^int32(-1))<<(uint(int32(6))%32))+16))
	v555 = v546
	goto L111
L113:
	;
	goto L114
L114:
	;
	v548 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v548+v163<<(uint(int32(6))%32)+int32(-64))+16))
	v555 = v554
	goto L111
L115:
	;
	F_errfinish(m, int32(497115), int32(586), int32(426648))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_bloom_opcinfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = F_palloc0(m, int32(44))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)) = uint8(v7)
		*(*uint16)(unsafe.Add(mBase, uint32(v3))) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = (v3 + int32(19)) & int32(-8)
		v18 = F_lookup_type_cache(m, int32(4600), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v18
			return v3
		}
	}
}
func F_brin_bloom_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_real_reloption(m, v2, int32(402523), int32(403582), float64(-0.1), float64(-1), float64(2.147483647e+09), int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_add_local_real_reloption(m, v2, int32(355207), int32(132485), float64(0.01), float64(0.0001), float64(0.25), int32(16))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return int32(0)
		}
	}
}
func F_brin_copy_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	if l3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v6 != 0 {
			if base.Ui32(v6) < base.Ui32(l1) {
				v16 = F_repalloc(m, l2, l1)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = l1
					v19 = v16
					if l1 != 0 {
						v20 = F__emscripten_memcpy_bulkmem(m, v19, l0, l1)
						mBase = m.M
						v21 = v20
					} else {
						v21 = v19
					}
					return v21
				}
			} else {
				v19 = l2
				if l1 != 0 {
					v20 = F__emscripten_memcpy_bulkmem(m, v19, l0, l1)
					mBase = m.M
					v21 = v20
				} else {
					v21 = v19
				}
				return v21
			}
		} else {
			v8 = F_palloc(m, l1)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				if l1 != 0 {
					v12 = F__emscripten_memcpy_bulkmem(m, v8, l0, l1)
					mBase = m.M
					v13 = v12
				} else {
					v13 = v8
				}
				return v13
			}
		}
	} else {
		v8 = F_palloc(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			if l1 != 0 {
				v12 = F__emscripten_memcpy_bulkmem(m, v8, l0, l1)
				mBase = m.M
				v13 = v12
			} else {
				v13 = v8
			}
			return v13
		}
	}
}
func F_brin_doinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	if base.Ui32(l6) < base.Ui32(int32(8153)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L111
	}
L2:
	;
	F_brinRevmapExtend(m, l2, l4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
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
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L107
	}
L5:
	;
	return int32(0)
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v118 = F_brinLockRevmapPageForUpdate(m, l2, l4)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L32
	}
L8:
	;
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+41)) = uint8(v101)
	goto L7
L9:
	;
	goto L28
L10:
	;
	F_LockBuffer(m, v26, int32(2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v32 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if base.Ui32(v68) < base.Ui32(l6) {
		goto L23
	} else {
		goto L24
	}
L13:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+16)))
	v52 = v51 + v50
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
	if v53 != int32(61587) {
		v68 = v8
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v32^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L13
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v50 = v44 + v32<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L17:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
	if v56&int32(1) != 0 {
		v68 = v8
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v59 = int32(4)
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+14)))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50)+12)))
	v62 = v60 - v61
	if v62 <= v59 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v68 = v65 - int32(4)
	goto L12
L20:
	;
	v65 = v59
	goto L22
L21:
	;
	v65 = v62
	goto L22
L22:
	;
	goto L19
L23:
	;
	F_UnlockReleaseBuffer(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v69 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L9
L27:
	;
	goto L9
L28:
	;
	v96 = F_brin_getinsertbuffer(m, l0, int32(0), l6, v18+int32(41))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L30
	}
L29:
	;
	goto L7
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v96
	if v96 == int32(0) {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v120 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v120 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(v120^int32(-1))<<(uint(int32(2))%32))))
	v138 = v130
	goto L33
L35:
	;
	goto L36
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v138 = v132 + v120<<(uint(int32(13))%32) + int32(-8192)
	goto L33
L37:
	;
	v158 = int32(4514932)
	v160 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v161 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v160 + v161
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+41)))
	if v164 == v161 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v120^int32(-1))<<(uint(int32(6))%32))+16))
	v157 = v148
	goto L37
L39:
	;
	goto L40
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+v120<<(uint(int32(6))%32)+int32(-64))+16))
	v157 = v156
	goto L37
L41:
	;
	if v138&int32(3) != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L43
L43:
	;
	v212 = int32(0)
	v214 = F_PageAddItemExtended(m, v138, l5, l6, v212, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L5
	} else {
		goto L54
	}
L44:
	;
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)))
	v210 = int32(61587)
	*(*uint16)(unsafe.Add(mBase, uint32(v138+v208)+6)) = uint16(v210)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+10)) = int32(1572864)
	v199 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+18)) = uint16(v199)
	v205 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)) = uint16(v205)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+14)) = uint16(v205)
	goto L44
L46:
	;
	v193 = F___memset(m, v138, int32(0), int32(8192))
	mBase = m.M
	goto L45
L47:
	;
	goto L46
L54:
	;
	if v214 == int32(0) {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_MarkBufferDirty(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	if v164 == int32(0) {
		v241 = v8
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+46)) = uint16(v214)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+24)) = uint16(v214)
	v245 = base.I32_rotr(v157, int32(16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+42)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v245
	v249 = v18 + int32(20)
	v252 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249))))
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+2)))
	if v118 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L58:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)))
	v224 = v138 + v223
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+6)))
	if v225 != int32(61587) {
		v241 = v8
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+4)))
	if v228&int32(1) != 0 {
		v241 = v8
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v231 = int32(4)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+14)))
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v234 = v232 - v233
	if v234 <= v231 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v241 = v237 - int32(4)
	goto L57
L62:
	;
	v237 = v231
	goto L64
L63:
	;
	v237 = v234
	goto L64
L64:
	;
	goto L61
L65:
	;
	F_MarkBufferDirty(m, v118)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L76
	}
L66:
	;
	v272 = base.I32_div_u_s(l4, l1)
	v274 = base.I32_rem_u_s(v272, int32(1360))
	v277 = v271 + v274*int32(6)
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v249)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+28)) = uint16(v278)
	if v278 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257+(v118^int32(-1))<<(uint(int32(2))%32))))
	v271 = v263
	goto L66
L68:
	;
	goto L69
L69:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v271 = v265 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L66
L70:
	;
	v281 = v253
	goto L72
L71:
	;
	v281 = int32(-1)
	goto L72
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+26)) = uint16(v281)
	if v278 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v284 = v252
	goto L75
L74:
	;
	v284 = int32(-1)
	goto L75
L75:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v277)+24)) = uint16(v284)
	goto L65
L76:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+118)))
	if v289 != int32(112) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v359 = int32(4514932)
	v361 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v361 - int32(1)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_LockBuffer(m, v365, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L100
	}
L78:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v293 <= int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v296 != 0 {
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+36)) = uint16(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v297 != 0 {
		goto L77
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_XLogRegisterData(m, v18+int32(28), int32(10))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v164 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v312 = int32(14)
	goto L88
L87:
	;
	v312 = int32(8)
	goto L88
L88:
	;
	F_XLogRegisterBuffer(m, int32(0), v309, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	F_XLogRegisterBufData(m, int32(0), l5, l6)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_XLogRegisterBuffer(m, int32(1), v118, int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	if v164 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v325 = int32(144)
	goto L94
L93:
	;
	v325 = int32(16)
	goto L94
L94:
	;
	v326 = F_XLogInsert(m, int32(17), v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v328 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = base.I64_rotr(v326, v328)
	if v118 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v352)+4)) = base.I32_wrap_i64(v326)
	*(*int32)(unsafe.Add(mBase, uint32(v352))) = base.I32_wrap_i64(int64(base.Ui64(v326) >> (uint(v328) % 64)))
	goto L77
L97:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338+(v118^int32(-1))<<(uint(int32(2))%32))))
	v352 = v344
	goto L96
L98:
	;
	goto L99
L99:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v352 = v346 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	F_LockBuffer(m, v118, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v164 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	F_RecordPageWithFreeSpace(m, l0, v157, v241)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v18 + int32(48)
	return v214
L105:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v157, v157+int32(1))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(8152)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v389 + int32(4)
	F_errmsg(m, int32(694372), v18)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(494943), int32(362), int32(81827))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errmsg_internal(m, int32(408477), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(494943), int32(414), int32(81827))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_fill_empty_ranges(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	if l1 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(0)
	goto L3
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = v11 + l1
	goto L3
L3:
	;
	if base.Ui32(v13) < base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v13
	goto L7
L5:
	;
	goto L6
L6:
	;
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v26 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v51 = F_brin_doinsert(m, v47, v48, v49, l0+int32(24), v20, v45, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L13
	} else {
		goto L16
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v30 = F_brin_new_memtuple(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v20
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v45 = v44
	goto L9
L13:
	;
	return
L14:
	;
	v32 = int32(4520272)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v38 = F_brin_form_tuple(m, v37, v20, v30, l0+int32(56))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v38
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
	v45 = v38
	goto L9
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v54 = v53 + v20
	if base.Ui32(v54) < base.Ui32(l2) {
		v20 = v54
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L8
}
func F_brin_inclusion_add_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v16 = base.I32_extend16_s(v15)
	v17 = int32(4)
	v21 = v13 + v16<<(uint(v17)%32) + v17
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
	if v24 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v29 = F_datumCopy(m, v22, v27, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v46 != 0 {
		v260 = v44
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v36)
	goto L3
L6:
	;
	return v260
L7:
	;
	v51 = v15<<(uint(int32(2))%32) + v12 + int32(16)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+115)))
	if v54 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v24 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v56 = v53 + int32(84)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+88))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+115)) = uint8(v112)
	goto L8
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+216))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)+204))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v62+v64*(v16-int32(1))<<(uint(int32(2))%32)+int32(56)-int32(4))))
	goto L14
L12:
	;
	goto L13
L13:
	;
	if v56 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L14:
	;
	if v76 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v81 = F_index_getprocinfo(m, v79, v16, int32(14))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v86 = v53 + int32(100)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(0)
	goto L17
L17:
	;
	goto L13
L18:
	;
	v100 = F_FunctionCall1Coll(m, v56, v23, v22)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v100 == int32(0) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v106 = v104 + int32(8)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v107 != 0 {
		v260 = v44
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v108
	return v108
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+114)))
	if v119 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v260 = int32(1)
	goto L6
L25:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+113)))
	if v174 != 0 {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v121 = v118 + int32(56)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+60))
	if v122 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v118)+114)) = uint8(v169)
	goto L25
L28:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v125)+216))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125)+204))
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+6)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v127+v129*(v16-int32(1))<<(uint(int32(2))%32)+int32(52)-int32(4))))
	goto L31
L29:
	;
	goto L30
L30:
	;
	if v121 == int32(0) {
		goto L25
	} else {
		goto L35
	}
L31:
	;
	if v141 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v146 = F_index_getprocinfo(m, v144, v16, int32(13))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v151 = v118 + int32(72)
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v146)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v151))) = v152
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v146)))
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v146)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+24)) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v146)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v151))) = int32(0)
	goto L34
L34:
	;
	goto L30
L35:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v167 = F_FunctionCall2Coll(m, v121, v23, v166, v22)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v167 != 0 {
		v260 = v44
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L25
L38:
	;
	v234 = F_inclusion_get_procinfo(m, v12, v16&int32(65535))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L51
	}
L39:
	;
	v176 = v173 + int32(28)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+32))
	if v177 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v229 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v173)+113)) = uint8(v229)
	goto L38
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+216))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+204))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+6)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v182+v184*(v16-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
	goto L44
L42:
	;
	goto L43
L43:
	;
	if v176 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L44:
	;
	if v196 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L45:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v201 = F_index_getprocinfo(m, v199, v16, int32(12))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v206 = v173 + int32(44)
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v201)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v206))) = v207
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	*(*int64)(unsafe.Add(mBase, uint32(v176))) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v201)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+24)) = v211
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v176)+8)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v176)+20)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = int32(0)
	goto L47
L47:
	;
	goto L43
L48:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = F_FunctionCall2Coll(m, v176, v23, v221, v22)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v222 != 0 {
		goto L38
	} else {
		goto L50
	}
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v225 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = v225
	return v225
L51:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v238 = F_FunctionCall2Coll(m, v234, v23, v237, v22)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
	if v240 != 0 {
		v251 = v238
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = v251
	goto L24
L54:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v238 == v242 {
		v251 = v238
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_pfree(m, v242)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	if v238 != v22 {
		v251 = v238
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
	v248 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v249 = F_datumCopy(m, v22, v247, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	v251 = v249
	goto L53
}
func F_brin_mask(m *base.Module, l0 int32, l1 int32) {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v11 = v9 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v11)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v13)+6)))
	switch v15 - int32(61585) {
	case 0:
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(v18) < base.Ui32(int32(25)) {
			v24 = v13
			v27 = l0 + v24 + int32(4)
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
			v30 = v28 & int32(65534)
			*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v30)
			return
		} else {
			F_mask_unused_space(m, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
				v24 = v23
				v27 = l0 + v24 + int32(4)
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
				v30 = v28 & int32(65534)
				*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v30)
				return
			}
		}
	default:
		v24 = v13
		v27 = l0 + v24 + int32(4)
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
		v30 = v28 & int32(65534)
		*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v30)
		return
	case 2:
		F_mask_unused_space(m, l0)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			v24 = v23
			v27 = l0 + v24 + int32(4)
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27))))
			v30 = v28 & int32(65534)
			*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v30)
			return
		}
	}
}
func F_brin_memtuple_initialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_MemoryContextReset(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if int32(0) < v13 {
			v16 = int32(20)
			v30 = l0 + (v13*v16+int32(31))&int32(-8)
			v31 = int32(0)
			for {
				v37 = l0 + int32(24) + v31*int32(20)
				v39 = v31 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v37))) = uint16(v39)
				v41 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v37)+16)) = v41
				*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v41
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v30
				v46 = int32(256)
				*(*uint16)(unsafe.Add(mBase, uint32(v37)+2)) = uint16(v46)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v48
				v50 = int32(2)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1+v16+v31<<(uint(v50)%32))))
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53))))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				if v39 < v59 {
					v30 = v30 + v54<<(uint(v50)%32)
					v31 = v39
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v69 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v69)
		return
	}
}
func F_brin_minmax_multi_add_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 float64
	_ = v260
	var v263 int32
	_ = v263
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v314 int32
	_ = v314
	var v317 float64
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
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
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v526 int32
	_ = v526
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v624 int32
	_ = v624
	var v627 int32
	_ = v627
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
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	var v823 int32
	_ = v823
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(16)
	m.G0 = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = F_get_fn_opclass_options(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28))))
	v45 = v35 + v36<<(uint(int32(4))%32) + v40*int32(100) - int32(80)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)))
	if v47 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v165 = v40 & int32(65535)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	v168 = F_minmax_multi_get_strategy_procinfo(m, v29, v165, v166, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+180))
	if v51 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v109 != 0 {
		v158 = v109
		v163 = v2
		goto L3
	} else {
		goto L30
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v56 = v52 * int32(291)
	goto L9
L8:
	;
	v56 = int32(37248)
	goto L9
L9:
	;
	if v31 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v57 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v62 = int32(32)
	goto L12
L12:
	;
	v63 = int32(4520272)
	v64 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v66
	v71 = v62 * int32(10)
	if base.Ui32(v71) < base.Ui32(v56) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v59 = v57
	goto L15
L14:
	;
	v59 = int32(32)
	goto L15
L15:
	;
	v62 = v59
	goto L12
L16:
	;
	v73 = v71
	goto L18
L17:
	;
	v73 = v56
	goto L18
L18:
	;
	if v62 < v73 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v75 = v73
	goto L21
L20:
	;
	v75 = v62
	goto L21
L21:
	;
	if v75 <= int32(256) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = int32(256)
	goto L24
L23:
	;
	v78 = v75
	goto L24
L24:
	;
	if int32(8192) <= v78 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = int32(8192)
	goto L27
L26:
	;
	v81 = v78
	goto L27
L27:
	;
	v86 = F_palloc0(m, v81<<(uint(int32(2))%32)+int32(36))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v86)+8)) = uint16(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v46
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v91
	v94 = int32(1)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	v99 = F_minmax_multi_get_strategy_procinfo(m, v29, v40&int32(65535), v97, v94)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v99
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v86
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+3)) = uint8(v107)
	v158 = v86
	v163 = v94
	goto L3
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+180))
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v116 = v112 * int32(291)
	goto L33
L32:
	;
	v116 = int32(37248)
	goto L33
L33:
	;
	v117 = int32(4520272)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v120
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = F_pg_detoast_datum(m, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v130 = v128 * int32(10)
	if base.Ui32(v130) < base.Ui32(v116) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v132 = v130
	goto L37
L36:
	;
	v132 = v116
	goto L37
L37:
	;
	if v128 < v132 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v134 = v132
	goto L40
L39:
	;
	v134 = v128
	goto L40
L40:
	;
	if v134 <= int32(256) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v137 = int32(256)
	goto L43
L42:
	;
	v137 = v134
	goto L43
L43:
	;
	if int32(8192) <= v137 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v140 = int32(8192)
	goto L46
L45:
	;
	v140 = v137
	goto L46
L46:
	;
	v141 = F_brin_range_deserialize(m, v140, v126)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v46
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+8)) = uint16(v40)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v145
	v150 = F_minmax_multi_get_strategy_procinfo(m, v29, v40&int32(65535), v145, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v141
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v118
	v158 = v141
	v163 = v2
	goto L3
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v174 = v170 + v171<<(uint(int32(1))%32)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v158)+28))
	if v174 < v175 {
		v549 = v171
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v27
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	if v549 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
	v179 = F_minmax_multi_get_strategy_procinfo(m, v29, v165, v177, int32(1))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_range_deduplicate_values(m, v158)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v158)+28))
	if base.F64_le(base.F64_convert_i32_s(v183+v184<<(uint(int32(1))%32)), base.F64_mul(base.F64_convert_i32_s(v189), float64(0.5))) != 0 {
		v549 = v184
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v194 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v201 = F_AllocSetContextCreateInternal(m, v196, int32(61257), v194, int32(8192), int32(8388608))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v203 = int32(4520272)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v201
	v209 = F_build_expanded_ranges(m, v179, v46, v158, v25+int32(4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v213 = F_minmax_multi_get_procinfo(m, v29, v40&int32(65535))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v215 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v219 = v215 - int32(1)
	v222 = F_palloc0(m, v219<<(uint(int32(4))%32))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v296 = v194
	goto L60
L60:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v158)+28))
	v317 = base.F64_mul(base.F64_convert_i32_s(v314), float64(0.5))
	if base.F64_lt(base.F64_abs(v317), float64(2.147483648e+09)) != 0 {
		goto L71
	} else {
		goto L72
	}
L61:
	;
	if int32(0) < v219 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v228 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	F_pg_qsort(m, v222, v219, int32(16), int32(20))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L69
	}
L65:
	;
	v251 = v209 + v228*int32(12)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	v254 = F_FunctionCall2Coll(m, v213, v46, v252, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	v258 = v222 + v228<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v258))) = v228
	v260 = *(*float64)(unsafe.Add(mBase, uint32(v254)))
	*(*float64)(unsafe.Add(mBase, uint32(v258)+8)) = v260
	v263 = v228 + int32(1)
	if v263 != v219 {
		v228 = v263
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v296 = v222
	goto L60
L70:
	;
	v324 = F_reduce_expanded_ranges(m, v209, v215, v296, v323, v179, v46)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L74
	}
L71:
	;
	v321 = base.I32_trunc_f64_s(v317)
	v323 = v321
	goto L70
L72:
	;
	goto L73
L73:
	;
	v323 = int32(-2147483648)
	goto L70
L74:
	;
	v326 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = v326
	if v326 < v324 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v526
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v204
	F_MemoryContextDelete(m, v201)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L100
	}
L76:
	;
	v331 = v158 + int32(36)
	v332 = int32(0)
	v335 = v332
	v336 = v332
	goto L79
L77:
	;
	v503 = int32(0)
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v503
	v526 = v503
	goto L75
L79:
	;
	v358 = v209 + v336*int32(12)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+8)))
	if v359 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v380 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v380
	v383 = int32(1)
	if v324 == v383 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v362 = int32(2)
	v364 = v331 + v335<<(uint(v362)%32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	*(*int32)(unsafe.Add(mBase, uint32(v364))) = v365
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+4)) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = v369 + int32(1)
	v375 = v335 + v362
	goto L83
L82:
	;
	v375 = v335
	goto L83
L83:
	;
	v378 = v336 + int32(1)
	if v378 != v324 {
		v335 = v375
		v336 = v378
		goto L79
	} else {
		goto L84
	}
L84:
	;
	goto L80
L85:
	;
	if v324&v383 == int32(0) {
		v526 = v465
		goto L75
	} else {
		goto L98
	}
L86:
	;
	v460 = v375
	v461 = int32(0)
	v465 = v380
	goto L85
L87:
	;
	goto L88
L88:
	;
	v390 = int32(0)
	v393 = v375
	v394 = v390
	v395 = v390
	v398 = v380
	goto L89
L89:
	;
	v416 = v209 + v394*int32(12)
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+8)))
	if v417 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v460 = v452
	v461 = v455
	v465 = v453
	goto L85
L91:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
	*(*int32)(unsafe.Add(mBase, uint32(v331+v393<<(uint(int32(2))%32)))) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v426 = int32(1)
	v427 = v425 + v426
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v427
	v431 = v393 + v426
	v432 = v427
	goto L93
L92:
	;
	v431 = v393
	v432 = v398
	goto L93
L93:
	;
	v433 = int32(1)
	v437 = v209 + (v394|v433)*int32(12)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+8)))
	if v438 == v433 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	*(*int32)(unsafe.Add(mBase, uint32(v331+v431<<(uint(int32(2))%32)))) = v444
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v447 = int32(1)
	v448 = v446 + v447
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v448
	v452 = v431 + v447
	v453 = v448
	goto L96
L95:
	;
	v452 = v431
	v453 = v432
	goto L96
L96:
	;
	v454 = int32(2)
	v455 = v394 + v454
	v457 = v395 + v454
	if v457 != v324&int32(2147483646) {
		v393 = v452
		v394 = v455
		v395 = v457
		v398 = v453
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v485 = v209 + v461*int32(12)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485)+8)))
	if v486 != int32(1) {
		v526 = v465
		goto L75
	} else {
		goto L99
	}
L99:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v485)))
	*(*int32)(unsafe.Add(mBase, uint32(v331+v460<<(uint(int32(2))%32)))) = v492
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v503 = v494 + int32(1)
	goto L78
L100:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v549 = v547
	goto L50
L101:
	;
	m.G0 = v25 + int32(16)
	return v823 | v163
L102:
	;
	v780 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+82)))
	v781 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+72)))
	v782 = F_datumCopy(m, v27, v780, v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L135
	}
L103:
	;
	v823 = base.B2i32(v175 <= v174)
	goto L101
L104:
	;
	v669 = F_minmax_multi_get_strategy_procinfo(m, v29, v40&int32(65535), v571, int32(3))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L123
	}
L105:
	;
	v575 = v158 + int32(36)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v575+v549<<(uint(int32(3))%32)-int32(4))))
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v158)+36))
	v584 = v40 & int32(65535)
	v586 = F_minmax_multi_get_strategy_procinfo(m, v29, v584, v571, int32(1))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v588 = F_FunctionCall2Coll(m, v586, v46, v27, v582)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v588 != 0 {
		goto L104
	} else {
		goto L108
	}
L108:
	;
	v591 = F_minmax_multi_get_strategy_procinfo(m, v29, v584, v571, int32(5))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v593 = F_FunctionCall2Coll(m, v591, v46, v27, v581)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v593 != 0 {
		goto L104
	} else {
		goto L111
	}
L111:
	;
	v595 = int32(0)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v598 = v596 - int32(1)
	if v598 < v595 {
		goto L104
	} else {
		goto L112
	}
L112:
	;
	v602 = v598
	v603 = v595
	v604 = v598
	goto L113
L113:
	;
	v624 = base.I32_div_s(v604, int32(2))
	v627 = v575 + v624<<(uint(int32(3))%32)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+4))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v627)))
	v630 = F_FunctionCall2Coll(m, v586, v46, v27, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L104
L115:
	;
	if v641 <= v640 {
		v602 = v640
		v603 = v641
		v604 = v640 + v641
		goto L113
	} else {
		goto L122
	}
L116:
	;
	if v630 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v640 = v624 - int32(1)
	v641 = v603
	goto L115
L118:
	;
	goto L119
L119:
	;
	v634 = F_FunctionCall2Coll(m, v591, v46, v27, v628)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	if v634 == int32(0) {
		goto L103
	} else {
		goto L121
	}
L121:
	;
	v640 = v602
	v641 = v624 + int32(1)
	goto L115
L122:
	;
	goto L114
L123:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	if int32(16) <= v671 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v674
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v676
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v686 = int32(4)
	v690 = F_bsearch_arg(m, v25+int32(12), v158+v680<<(uint(int32(3))%32)+int32(36), v671, v686, int32(21), v25+v686)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v694 = v692 << (uint(int32(1)) % 32)
	if v694+v671 <= v694 {
		goto L102
	} else {
		goto L129
	}
L127:
	;
	if v690 != 0 {
		goto L103
	} else {
		goto L128
	}
L128:
	;
	goto L102
L129:
	;
	v700 = v694
	goto L130
L130:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v158+int32(36)+v700<<(uint(int32(2))%32))))
	v725 = F_FunctionCall2Coll(m, v669, v46, v27, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L102
L132:
	;
	if v725 != 0 {
		goto L103
	} else {
		goto L133
	}
L133:
	;
	v727 = int32(1)
	v728 = v700 + v727
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v158)+20))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	if v728 < v729+v730<<(uint(v727)%32) {
		v700 = v728
		goto L130
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v786 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v158+(v784+v785<<(uint(v786)%32))<<(uint(int32(2))%32))+36)) = v782
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+24)) = v793 + v786
	if v793 != 0 {
		v823 = v786
		goto L101
	} else {
		goto L136
	}
L136:
	;
	v798 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = v798
	v823 = v798
	goto L101
}
func F_brin_minmax_multi_distance_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(86400000000)
	v8 = base.I64_div_s(v6, v7)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v15 = base.I64_rem_s(v13, v7)
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
	v22 = base.I64_div_s(v13, int64(-86400000000))
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+12)))
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+12)))
	v35 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i64_s(v6-v8*v7-v15), float64(8.64e+10)), base.F64_convert_i64_s(v20+(v22+v8)-v25+(v27-v28)*int64(30))))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		return v35
	}
}
func F_brin_minmax_multi_distance_timetz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	v16 = F_Float8GetDatum(m, base.F64_convert_i64_s(v4-v6+base.I64_extend_i32_s(v8-v9)*int64(1000000)))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		return v16
	}
}
func F_brin_minmax_multi_summary_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	F_initStringInfo(m, v11+int32(128))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_appendStringInfoChar(m, v11+int32(128), int32(123))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	F_getTypeOutputInfo(m, v27, v11+int32(120), v11+int32(127))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	F_fmgr_info(m, v34, v11+int32(92))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v40 = F_brin_range_deserialize(m, v39, v25)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v42
	F_appendStringInfo(m, v11+int32(128), int32(482994), v11+int32(48))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v55 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v145 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v137 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = int32(0)
	v63 = v61
	v65 = v61
	v67 = int32(0)
	goto L13
L13:
	;
	F_initStringInfo(m, v11+int32(76))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v111 <= int32(0) {
		v137 = v98
		goto L9
	} else {
		goto L22
	}
L15:
	;
	v79 = v40 + int32(36) + v63<<(uint(int32(2))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = F_OutputFunctionCall(m, v11+int32(92), v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v86 = F_OutputFunctionCall(m, v11+int32(92), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v81
	F_appendStringInfo(m, v11+int32(76), int32(206200), v11+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v98 = v63 + int32(2)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v101 = F_cstring_to_text_with_len(m, v99, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v107 = F_accumArrayResult(m, v65, v101, int32(0), int32(25), v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v110 = v67 + int32(1)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v110 < v111 {
		v63 = v98
		v65 = v107
		v67 = v110
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	F_getTypeOutputInfo(m, int32(2277), v11+int32(76), v11+int32(75))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v124 = F_makeArrayResult(m, v107, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v127 = F_OidOutputFunctionCall(m, v126, v124)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v127
	F_appendStringInfo(m, v11+int32(128), int32(200984), v11+int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v137 = v98
	goto L9
L27:
	;
	F_appendStringInfoChar(m, v11+int32(128), int32(125))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L40
	}
L28:
	;
	v150 = int32(0)
	v152 = v137
	v154 = v150
	v156 = v150
	goto L29
L29:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(36)+v152<<(uint(int32(2))%32))))
	v167 = F_FunctionCall1Coll(m, v11+int32(92), int32(0), v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	if v181 <= int32(0) {
		goto L27
	} else {
		goto L35
	}
L31:
	;
	v169 = F_cstring_to_text(m, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v175 = F_accumArrayResult(m, v154, v169, int32(0), int32(25), v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v177 = int32(1)
	v180 = v156 + v177
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v180 < v181 {
		v152 = v152 + v177
		v154 = v175
		v156 = v180
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	F_getTypeOutputInfo(m, int32(2277), v11+int32(76), v11+int32(75))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v194 = F_makeArrayResult(m, v175, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v11)+76))
	v197 = F_OidOutputFunctionCall(m, v196, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v197
	F_appendStringInfo(m, v11+int32(128), int32(200870), v11)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L27
L40:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	m.G0 = v11 + int32(144)
	return v218
}
func F_brin_minmax_multi_summary_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(17916)
			F_errmsg(m, int32(192913), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(498863), int32(3121), int32(36429))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_brin_new_memtuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = F_palloc0(m, (v10*int32(20)+int32(31))&int32(-8)+v17<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v28 = F_palloc(m, v25<<(uint(int32(2))%32))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v28
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
			v33 = F_palloc(m, v32)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v33
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				v38 = F_palloc(m, v37)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v40)
					*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v38
					v44 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v49 = F_AllocSetContextCreateInternal(m, v44, int32(384325), int32(0), int32(8192), int32(8388608))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v49
						F_MemoryContextReset(m, v49)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
							if int32(0) < v55 {
								v58 = int32(20)
								v72 = int32(0)
								v73 = v21 + (v55*v58+int32(31))&int32(-8)
								for {
									v79 = v21 + int32(24) + v72*int32(20)
									v81 = v72 + int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v81)
									v83 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = v83
									*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v83
									*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v73
									v88 = int32(256)
									*(*uint16)(unsafe.Add(mBase, uint32(v79)+2)) = uint16(v88)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v90
									v92 = int32(2)
									v95 = *(*int32)(unsafe.Add(mBase, uint32(l0+v58+v72<<(uint(v92)%32))))
									v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95))))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
									if v81 < v101 {
										v72 = v81
										v73 = v73 + v96<<(uint(v92)%32)
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v111 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v111)
							return v21
						}
					}
				}
			}
		}
	}
}
func F_brin_page_init(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v2 = l1
	if l0&int32(3) != 0 {
	} else {
	}
	v29 = F___memset(m, l0, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v35 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v35)
	v41 = int32(8184)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v41)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v41)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v44)+6)) = uint16(v2)
	return
}
func F_brin_range_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	F_range_deduplicate_values(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v25 = v21 + v22<<(uint(int32(1))%32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = F_get_typbyval(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = F_get_typlen(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v133 = F_palloc0(m, v123)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L38
	}
L5:
	;
	if v29 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v25 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v29 == int32(-2) {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v123 = int32(20)
	goto L4
L10:
	;
	goto L11
L11:
	;
	v41 = int32(20)
	v42 = v2
	goto L12
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)+v42<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v55 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v123 = v83
	goto L4
L14:
	;
	v83 = v82 + v41
	v85 = v42 + int32(1)
	if v85 != v25 {
		v41 = v83
		v42 = v85
		goto L12
	} else {
		goto L28
	}
L15:
	;
	v58 = int32(6)
	v60 = int32(18)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v62 == v60 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v74 = int32(1)
	if v55&v74 != 0 {
		v82 = int32(base.Ui32(v55) >> (uint(v74) % 32))
		goto L14
	} else {
		goto L27
	}
L18:
	;
	v65 = v60
	goto L20
L19:
	;
	v65 = int32(2)
	goto L20
L20:
	;
	if v62&int32(254) == int32(2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = v58
	goto L23
L22:
	;
	v70 = v65
	goto L23
L23:
	;
	if v62 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v73 = v58
	goto L26
L25:
	;
	v73 = v70
	goto L26
L26:
	;
	v82 = v73
	goto L14
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v82 = int32(base.Ui32(v78) >> (uint(int32(2)) % 32))
	goto L14
L28:
	;
	goto L13
L29:
	;
	if v25 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v123 = v29*v25 + int32(20)
	goto L4
L32:
	;
	v123 = int32(20)
	goto L4
L33:
	;
	goto L34
L34:
	;
	v97 = int32(20)
	v98 = v2
	goto L35
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)+v98<<(uint(int32(2))%32))))
	v111 = F_strlen(m, v110)
	mBase = m.M
	v113 = int32(1)
	v114 = v111 + v97 + v113
	v116 = v98 + v113
	if v116 != v25 {
		v97 = v114
		v98 = v116
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v123 = v114
	goto L4
L37:
	;
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v123 << (uint(int32(2)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+12)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = v143
	if int32(0) < v25 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v148 = l0 + int32(36)
	v160 = int32(0)
	v161 = v133 + int32(20)
	goto L42
L40:
	;
	goto L41
L41:
	;
	m.G0 = v15 + int32(16)
	return v133
L42:
	;
	if v27 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L41
L44:
	;
	v252 = v160 + int32(1)
	if v252 != v25 {
		v160 = v252
		v161 = v249
		goto L42
	} else {
		goto L88
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v148+v160<<(uint(int32(2))%32))))
	switch v29&int32(65535) - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L52
	default:
		goto L50
	case 3:
		goto L51
	}
L46:
	;
	goto L47
L47:
	;
	if int32(0) < v29 {
		goto L60
	} else {
		goto L61
	}
L48:
	;
	if v29 != 0 {
		goto L57
	} else {
		goto L58
	}
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v173)
	goto L48
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v173
	goto L48
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v173)
	goto L48
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v29
	F_errmsg_internal(m, int32(484426), v15)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(327351), int32(230), int32(309384))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v249 = v193 + v29
	goto L44
L57:
	;
	v192 = F__emscripten_memcpy_bulkmem(m, v161, v15+int32(12), v29)
	mBase = m.M
	v193 = v192
	goto L59
L58:
	;
	v193 = v161
	goto L59
L59:
	;
	goto L56
L60:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v148+v160<<(uint(int32(2))%32))))
	if v29 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	if v29 == int32(-1) {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v249 = v202 + v29
	goto L44
L64:
	;
	v201 = F__emscripten_memcpy_bulkmem(m, v161, v200, v29)
	mBase = m.M
	v202 = v201
	goto L66
L65:
	;
	v202 = v161
	goto L66
L66:
	;
	goto L63
L67:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v148+v160<<(uint(int32(2))%32))))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
	if v210 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	goto L69
L69:
	;
	if v29 != int32(-2) {
		v249 = v161
		goto L44
	} else {
		goto L83
	}
L70:
	;
	if v234 != 0 {
		goto L80
	} else {
		goto L81
	}
L71:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	if base.Ui32((v214-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v234 = int32(6)
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v226 = int32(1)
	if v210&v226 != 0 {
		v234 = int32(base.Ui32(v210) >> (uint(v226) % 32))
		goto L70
	} else {
		goto L78
	}
L74:
	;
	v221 = int32(18)
	if v214 == v221 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v225 = v221
	goto L77
L76:
	;
	v225 = int32(2)
	goto L77
L77:
	;
	v234 = v225
	goto L70
L78:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v234 = int32(base.Ui32(v230) >> (uint(int32(2)) % 32))
	goto L70
L79:
	;
	v249 = v236 + v234
	goto L44
L80:
	;
	v235 = F__emscripten_memcpy_bulkmem(m, v161, v209, v234)
	mBase = m.M
	v236 = v235
	goto L82
L81:
	;
	v236 = v161
	goto L82
L82:
	;
	goto L79
L83:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v148+v160<<(uint(int32(2))%32))))
	v242 = F_strlen(m, v241)
	mBase = m.M
	v244 = v242 + int32(1)
	if v244 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v249 = v246 + v244
	goto L44
L85:
	;
	v245 = F__emscripten_memcpy_bulkmem(m, v161, v241, v244)
	mBase = m.M
	v246 = v245
	goto L87
L86:
	;
	v246 = v161
	goto L87
L87:
	;
	goto L84
L88:
	;
	goto L43
}
func F_brin_xlog_insert_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+48)))
	if v15 < int32(0) {
		v19 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v19
			if v19 < int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(v19^int32(-1))<<(uint(int32(2))%32))))
				v39 = v31
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v39 = v33 + v19<<(uint(int32(13))%32) + int32(-8192)
			}
			v40 = int32(61587)
			F_PageInit(m, v39, int32(8192), int32(8))
			mBase = m.M
			v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+16)))
			*(*uint16)(unsafe.Add(mBase, uint32(v39+v44)+6)) = uint16(v40)
			if v19 < int32(0) {
				v50 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v50+(v19^int32(-1))<<(uint(int32(6))%32))+16))
				v65 = v56
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v58+v19<<(uint(int32(6))%32)+int32(-64))+16))
				v65 = v64
			}
			v92 = v65
			v93 = int32(0)
			v95 = v11 + int32(20)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+72))
			if v98 < v93 {
				v120 = v93
				v123 = v120
			} else {
				v104 = v97 + int32(76)
				v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
				if v105 != int32(1) {
					v120 = v93
					v123 = v120
				} else {
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+43)))
					if v108 == int32(0) {
						if v95 == int32(0) {
							v120 = v93
							v123 = v120
						} else {
							v113 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v95))) = v113
							v123 = v113
						}
					} else {
						if v95 != 0 {
							v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+48)))
							*(*int32)(unsafe.Add(mBase, uint32(v95))) = v116
						} else {
						}
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v104)+44))
						v120 = v118
						v123 = v120
					}
				}
			}
			v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			if v125 < int32(0) {
				v129 = *(*int32)(unsafe.Add(mBase, _consts[5]))
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v125^int32(-1))<<(uint(int32(2))%32))))
				v143 = v135
			} else {
				v137 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v143 = v137 + v125<<(uint(int32(13))%32) + int32(-8192)
			}
			v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+12)))
			if base.Ui32(v144) < base.Ui32(int32(25)) {
				v155 = int32(1)
			} else {
				v155 = int32(base.Ui32(v144+int32(262120))>>(uint(int32(2))%32))&int32(65535) + int32(1)
			}
			v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
			if base.Ui32(v155) < base.Ui32(v156) {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v272 = m.ExcPending
				if v272 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(228729), int32(0))
					mBase = m.M
					v276 = m.ExcPending
					if v276 != 0 {
						return
					} else {
						F_errfinish(m, int32(499243), int32(88), int32(356736))
						mBase = m.M
						v281 = m.ExcPending
						if v281 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				v160 = F_PageAddItemExtended(m, v143, v123, v158, v156, int32(1))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					if v160 == int32(0) {
						F_errstart_cold(m, int32(23), int32(0))
						mBase = m.M
						v285 = m.ExcPending
						if v285 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(385618), int32(0))
							mBase = m.M
							v289 = m.ExcPending
							if v289 != 0 {
								return
							} else {
								F_errfinish(m, int32(499243), int32(92), int32(356736))
								mBase = m.M
								v294 = m.ExcPending
								if v294 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v143))) = base.I64_rotr(v13, int64(32))
						v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
						F_MarkBufferDirty(m, v167)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							v171 = v92
							v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
							if v174 != 0 {
								F_UnlockReleaseBuffer(m, v174)
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										if v180 == int32(0) {
											v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
											v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
											v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											if v190 < int32(0) {
												v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
												v208 = v200
											} else {
												v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
												v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
											}
											v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
											v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
											v216 = v11 + int32(12)
											v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
											v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
											if v190 < int32(0) {
												v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
												v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
												v238 = v230
											} else {
												v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
												v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
											}
											v239 = base.I32_div_u_s(v209, v210)
											v241 = base.I32_rem_u_s(v239, int32(1360))
											v244 = v238 + v241*int32(6)
											v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
											if v245 != 0 {
												v248 = v220
											} else {
												v248 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
											if v245 != 0 {
												v251 = v219
											} else {
												v251 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
											*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
											v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											F_MarkBufferDirty(m, v256)
											mBase = m.M
											v258 = m.ExcPending
											if v258 != 0 {
												return
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											}
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											} else {
												m.G0 = v11 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return
								} else {
									if v180 == int32(0) {
										v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
										*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
										v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
										if v190 < int32(0) {
											v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
											v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
											v208 = v200
										} else {
											v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
										}
										v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
										*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
										v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
										v216 = v11 + int32(12)
										v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
										v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
										if v190 < int32(0) {
											v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
											v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
											v238 = v230
										} else {
											v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
											v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
										}
										v239 = base.I32_div_u_s(v209, v210)
										v241 = base.I32_rem_u_s(v239, int32(1360))
										v244 = v238 + v241*int32(6)
										v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
										if v245 != 0 {
											v248 = v220
										} else {
											v248 = int32(-1)
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
										if v245 != 0 {
											v251 = v219
										} else {
											v251 = int32(-1)
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
										*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
										v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
										F_MarkBufferDirty(m, v256)
										mBase = m.M
										v258 = m.ExcPending
										if v258 != 0 {
											return
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											} else {
												m.G0 = v11 + int32(32)
												return
											}
										}
									} else {
										v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
										if v263 != 0 {
											F_UnlockReleaseBuffer(m, v263)
											mBase = m.M
											v265 = m.ExcPending
											if v265 != 0 {
												return
											} else {
												m.G0 = v11 + int32(32)
												return
											}
										} else {
											m.G0 = v11 + int32(32)
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
	} else {
		v69 = F_XLogReadBufferForRedo(m, l0, int32(0), v11+int32(28))
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return
		} else {
			v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
			if v71 < int32(0) {
				v75 = *(*int32)(unsafe.Add(mBase, _consts[8]))
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v75+(v71^int32(-1))<<(uint(int32(6))%32))+16))
				v90 = v81
			} else {
				v83 = *(*int32)(unsafe.Add(mBase, _consts[9]))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+v71<<(uint(int32(6))%32)+int32(-64))+16))
				v90 = v89
			}
			if v69 != 0 {
				v171 = v90
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				if v174 != 0 {
					F_UnlockReleaseBuffer(m, v174)
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return
					} else {
						v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return
						} else {
							if v180 == int32(0) {
								v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
								*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
								v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
								*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
								if v190 < int32(0) {
									v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
									v208 = v200
								} else {
									v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
								}
								v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
								*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
								v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
								v216 = v11 + int32(12)
								v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
								v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
								if v190 < int32(0) {
									v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
									v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
									v238 = v230
								} else {
									v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
									v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
								}
								v239 = base.I32_div_u_s(v209, v210)
								v241 = base.I32_rem_u_s(v239, int32(1360))
								v244 = v238 + v241*int32(6)
								v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
								if v245 != 0 {
									v248 = v220
								} else {
									v248 = int32(-1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
								if v245 != 0 {
									v251 = v219
								} else {
									v251 = int32(-1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
								*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
								v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
								F_MarkBufferDirty(m, v256)
								mBase = m.M
								v258 = m.ExcPending
								if v258 != 0 {
									return
								} else {
									v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
									if v263 != 0 {
										F_UnlockReleaseBuffer(m, v263)
										mBase = m.M
										v265 = m.ExcPending
										if v265 != 0 {
											return
										} else {
											m.G0 = v11 + int32(32)
											return
										}
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								}
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
								if v263 != 0 {
									F_UnlockReleaseBuffer(m, v263)
									mBase = m.M
									v265 = m.ExcPending
									if v265 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								} else {
									m.G0 = v11 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						if v180 == int32(0) {
							v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
							*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
							*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
							v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
							*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
							if v190 < int32(0) {
								v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
								v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
								v208 = v200
							} else {
								v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
							}
							v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
							*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
							v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
							v216 = v11 + int32(12)
							v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
							v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
							if v190 < int32(0) {
								v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
								v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
								v238 = v230
							} else {
								v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
								v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
							}
							v239 = base.I32_div_u_s(v209, v210)
							v241 = base.I32_rem_u_s(v239, int32(1360))
							v244 = v238 + v241*int32(6)
							v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
							if v245 != 0 {
								v248 = v220
							} else {
								v248 = int32(-1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
							if v245 != 0 {
								v251 = v219
							} else {
								v251 = int32(-1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
							*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
							v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
							F_MarkBufferDirty(m, v256)
							mBase = m.M
							v258 = m.ExcPending
							if v258 != 0 {
								return
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
								if v263 != 0 {
									F_UnlockReleaseBuffer(m, v263)
									mBase = m.M
									v265 = m.ExcPending
									if v265 != 0 {
										return
									} else {
										m.G0 = v11 + int32(32)
										return
									}
								} else {
									m.G0 = v11 + int32(32)
									return
								}
							}
						} else {
							v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
							if v263 != 0 {
								F_UnlockReleaseBuffer(m, v263)
								mBase = m.M
								v265 = m.ExcPending
								if v265 != 0 {
									return
								} else {
									m.G0 = v11 + int32(32)
									return
								}
							} else {
								m.G0 = v11 + int32(32)
								return
							}
						}
					}
				}
			} else {
				v92 = v90
				v93 = int32(0)
				v95 = v11 + int32(20)
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+72))
				if v98 < v93 {
					v120 = v93
					v123 = v120
				} else {
					v104 = v97 + int32(76)
					v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
					if v105 != int32(1) {
						v120 = v93
						v123 = v120
					} else {
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+43)))
						if v108 == int32(0) {
							if v95 == int32(0) {
								v120 = v93
								v123 = v120
							} else {
								v113 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v95))) = v113
								v123 = v113
							}
						} else {
							if v95 != 0 {
								v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+48)))
								*(*int32)(unsafe.Add(mBase, uint32(v95))) = v116
							} else {
							}
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v104)+44))
							v120 = v118
							v123 = v120
						}
					}
				}
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
				if v125 < int32(0) {
					v129 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+(v125^int32(-1))<<(uint(int32(2))%32))))
					v143 = v135
				} else {
					v137 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v143 = v137 + v125<<(uint(int32(13))%32) + int32(-8192)
				}
				v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+12)))
				if base.Ui32(v144) < base.Ui32(int32(25)) {
					v155 = int32(1)
				} else {
					v155 = int32(base.Ui32(v144+int32(262120))>>(uint(int32(2))%32))&int32(65535) + int32(1)
				}
				v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
				if base.Ui32(v155) < base.Ui32(v156) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(228729), int32(0))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return
						} else {
							F_errfinish(m, int32(499243), int32(88), int32(356736))
							mBase = m.M
							v281 = m.ExcPending
							if v281 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
					v160 = F_PageAddItemExtended(m, v143, v123, v158, v156, int32(1))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						if v160 == int32(0) {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v285 = m.ExcPending
							if v285 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(385618), int32(0))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return
								} else {
									F_errfinish(m, int32(499243), int32(92), int32(356736))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v143))) = base.I64_rotr(v13, int64(32))
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
							F_MarkBufferDirty(m, v167)
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return
							} else {
								v171 = v92
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
								if v174 != 0 {
									F_UnlockReleaseBuffer(m, v174)
									mBase = m.M
									v176 = m.ExcPending
									if v176 != 0 {
										return
									} else {
										v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return
										} else {
											if v180 == int32(0) {
												v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
												v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
												v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												if v190 < int32(0) {
													v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
													v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
													v208 = v200
												} else {
													v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
													v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
												}
												v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
												*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
												v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
												v216 = v11 + int32(12)
												v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
												v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
												if v190 < int32(0) {
													v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
													v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
													v238 = v230
												} else {
													v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
													v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
												}
												v239 = base.I32_div_u_s(v209, v210)
												v241 = base.I32_rem_u_s(v239, int32(1360))
												v244 = v238 + v241*int32(6)
												v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
												*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
												if v245 != 0 {
													v248 = v220
												} else {
													v248 = int32(-1)
												}
												*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
												if v245 != 0 {
													v251 = v219
												} else {
													v251 = int32(-1)
												}
												*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
												*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
												v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												F_MarkBufferDirty(m, v256)
												mBase = m.M
												v258 = m.ExcPending
												if v258 != 0 {
													return
												} else {
													v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
													if v263 != 0 {
														F_UnlockReleaseBuffer(m, v263)
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return
														} else {
															m.G0 = v11 + int32(32)
															return
														}
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												}
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v180 = F_XLogReadBufferForRedo(m, l0, int32(1), v11+int32(28))
									mBase = m.M
									v181 = m.ExcPending
									if v181 != 0 {
										return
									} else {
										if v180 == int32(0) {
											v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)) = uint16(v184)
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v171)
											v188 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v188)
											v190 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											if v190 < int32(0) {
												v194 = *(*int32)(unsafe.Add(mBase, _consts[5]))
												v200 = *(*int32)(unsafe.Add(mBase, uint32(v194+(v190^int32(-1))<<(uint(int32(2))%32))))
												v208 = v200
											} else {
												v202 = *(*int32)(unsafe.Add(mBase, _consts[6]))
												v208 = v202 + v190<<(uint(int32(13))%32) + int32(-8192)
											}
											v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+24)))
											*(*uint16)(unsafe.Add(mBase, uint32(v11)+16)) = uint16(v211)
											v213 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v213
											v216 = v11 + int32(12)
											v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
											v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+2)))
											if v190 < int32(0) {
												v224 = *(*int32)(unsafe.Add(mBase, _consts[5]))
												v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v190^int32(-1))<<(uint(int32(2))%32))))
												v238 = v230
											} else {
												v232 = *(*int32)(unsafe.Add(mBase, _consts[6]))
												v238 = v232 + v190<<(uint(int32(13))%32) + int32(-8192)
											}
											v239 = base.I32_div_u_s(v209, v210)
											v241 = base.I32_rem_u_s(v239, int32(1360))
											v244 = v238 + v241*int32(6)
											v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+4)))
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+28)) = uint16(v245)
											if v245 != 0 {
												v248 = v220
											} else {
												v248 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+26)) = uint16(v248)
											if v245 != 0 {
												v251 = v219
											} else {
												v251 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v244)+24)) = uint16(v251)
											*(*int64)(unsafe.Add(mBase, uint32(v208))) = base.I64_rotr(v13, int64(32))
											v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											F_MarkBufferDirty(m, v256)
											mBase = m.M
											v258 = m.ExcPending
											if v258 != 0 {
												return
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v11 + int32(32)
														return
													}
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											}
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v11 + int32(32)
													return
												}
											} else {
												m.G0 = v11 + int32(32)
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
		}
	}
}
