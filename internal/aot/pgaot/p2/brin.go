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
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_brinLockRevmapPageForUpdate[0]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v18^int32(-1))<<(uint(int32(6))%32))+16))
				v39 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_brinLockRevmapPageForUpdate[1]))
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
			F_errmsg_internal(m, int32(_a_F_brinLockRevmapPageForUpdate_0), v8)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_brinLockRevmapPageForUpdate_1), int32(471), int32(_a_F_brinLockRevmapPageForUpdate_2))
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v42 int32
	_ = v42
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v299 int32
	_ = v299
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 int64
	_ = v433
	var v438 int64
	_ = v438
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = base.I32_div_u_s(l1, v20)
	v23 = base.I32_div_u_s(v21, int32(1360))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v24) <= base.Ui32(v23) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L10
	} else {
		goto L108
	}
L2:
	;
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v18 + int32(48)
	return
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[0]))
	if v42 != 0 {
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
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v46, int32(2))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
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
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v50 < int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v483) <= base.Ui32(v23) {
		goto L5
	} else {
		goto L107
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+36))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v69 != v70 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(v50^int32(-1))<<(uint(int32(2))%32))))
	v68 = v60
	goto L14
L16:
	;
	goto L17
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v68 = v62 + v50<<(uint(int32(13))%32) + int32(-8192)
	goto L14
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v69
	F_LockBuffer(m, v50, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v77 = v69 + int32(1)
	v79 = F_RelationGetNumberOfBlocksInFork(m, v45, int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L10
	} else {
		goto L24
	}
L21:
	;
	goto L13
L22:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v462, int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L10
	} else {
		goto L105
	}
L23:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+14)))
	if v157 != 0 {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	if base.Ui32(v77) < base.Ui32(v79) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = F_ReadBuffer(m, v45, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v45
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v107
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v109
	v113 = int32(0)
	v116 = F_ExtendBufferedRel(m, v18+int32(16), v113, v113, int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L33
	}
L28:
	;
	F_LockBuffer(m, v82, int32(2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v82 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v82^int32(-1))<<(uint(int32(2))%32))))
	v155 = v82
	v156 = v96
	goto L23
L31:
	;
	goto L32
L32:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v155 = v82
	v156 = v98 + v82<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L33:
	;
	if v116 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v136 != v77 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[3]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121+(v116^int32(-1))<<(uint(int32(6))%32))+16))
	v136 = v127
	goto L34
L36:
	;
	goto L37
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[4]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129+v116<<(uint(int32(6))%32)+int32(-64))+16))
	v136 = v135
	goto L34
L38:
	;
	v448 = v116
	goto L22
L39:
	;
	goto L40
L40:
	;
	if v116 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141+(v116^int32(-1))<<(uint(int32(2))%32))))
	v155 = v116
	v156 = v147
	goto L23
L42:
	;
	goto L43
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v155 = v116
	v156 = v149 + v116<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L44:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+v158)+6)))
	if v160 != int32(_a_F_brinRevmapExtend_0) {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v163 = int32(0)
	if v155 < v163 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L46
L48:
	;
	if v248 != 0 {
		goto L63
	} else {
		goto L64
	}
L49:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+14)))
	if v182 == int32(0) {
		v248 = v163
		goto L48
	} else {
		goto L53
	}
L50:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+(v155^int32(-1))<<(uint(int32(2))%32))))
	v181 = v173
	goto L49
L51:
	;
	goto L52
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v181 = v175 + v155<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+12)))
	if base.Ui32(v185) < base.Ui32(int32(25)) {
		v248 = v163
		goto L48
	} else {
		goto L54
	}
L54:
	;
	v193 = int32(base.Ui32(v185+int32(_a_F_brinRevmapExtend_1))>>(uint(int32(2))%32)) & int32(_a_F_brinRevmapExtend_2)
	if v193 == int32(0) {
		v248 = v163
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v201 = int32(1)
	goto L56
L56:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181+int32(20)+v201&int32(_a_F_brinRevmapExtend_2)<<(uint(int32(2))%32))+1)))
	if v219&int32(384) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v229 = int32(1)
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v231 = v181 + v230
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+4)))
	v234 = v232 | v229
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+4)) = uint16(v234)
	F_MarkBufferDirtyHint(m, v155, v229)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L62
	}
L58:
	;
	v225 = v201 + int32(1)
	if base.Ui32(v225&int32(_a_F_brinRevmapExtend_2)) <= base.Ui32(v193) {
		v201 = v225
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
	v248 = v163
	goto L48
L62:
	;
	v248 = v229
	goto L48
L63:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_LockBuffer(m, v254, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v380 = int32(_a_F_brinRevmapExtend_3)
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5])) = v382 + int32(1)
	v386 = int32(_a_F_brinRevmapExtend_4)
	F_PageInit(m, v156, int32(_a_F_brinRevmapExtend_5), int32(8))
	mBase = m.M
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v156+v390)+6)) = uint16(v386)
	goto L90
L66:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v259 = m.G0
	v261 = v259 - int32(16)
	m.G0 = v261
	v263 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v261)+12)) = v263
	if v155 < v263 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	F_UnlockReleaseBuffer(m, v155)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L10
	} else {
		goto L89
	}
L68:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+12)))
	if base.Ui32(v283) < base.Ui32(int32(25)) {
		goto L67
	} else {
		goto L72
	}
L69:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v268+(v155^int32(-1))<<(uint(int32(2))%32))))
	v282 = v274
	goto L68
L70:
	;
	goto L71
L71:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v282 = v276 + v155<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	v291 = int32(base.Ui32(v283+int32(_a_F_brinRevmapExtend_1))>>(uint(int32(2))%32)) & int32(_a_F_brinRevmapExtend_2)
	if v291 == int32(0) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v299 = int32(1)
	goto L74
L74:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[0]))
	if v313 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L67
L76:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v317 = v299 & int32(_a_F_brinRevmapExtend_2)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v282+int32(20)+v317<<(uint(int32(2))%32))))
	if v321&int32(_a_F_brinRevmapExtend_6) != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	v328 = int32(base.Ui32(v321) >> (uint(int32(17)) % 32))
	v332 = F_brin_copy_tuple(m, v282+v321&int32(_a_F_brinRevmapExtend_7), v328, int32(0), v261+int32(12))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	v354 = v299
	goto L82
L82:
	;
	v356 = v354 + int32(1)
	if base.Ui32(v356&int32(_a_F_brinRevmapExtend_2)) <= base.Ui32(v291) {
		v299 = v356
		goto L74
	} else {
		goto L88
	}
L83:
	;
	F_LockBuffer(m, v155, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v339 = F_brin_doupdate(m, v45, v258, l0, v337, v155, v317, v332, v328, v332, v328, int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_LockBuffer(m, v155, int32(1))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282+v344)+6)))
	if v346 != int32(_a_F_brinRevmapExtend_0) {
		goto L67
	} else {
		goto L87
	}
L87:
	;
	v354 = v299 - (v339 ^ int32(1))
	goto L82
L88:
	;
	goto L75
L89:
	;
	m.G0 = v261 + int32(16)
	goto L13
L90:
	;
	F_MarkBufferDirty(m, v155)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v395 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)) = uint16(v395)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+36)) = v77
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MarkBufferDirty(m, v398)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+48))
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402)+118)))
	if v403 != int32(112) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v441 = int32(_a_F_brinRevmapExtend_3)
	v443 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5])) = v443 - int32(1)
	v448 = v155
	goto L22
L94:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[6]))
	if v407 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v401)+32))
	if v410 != 0 {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v77
	F_XLogBeginInsert(m)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L10
	} else {
		goto L100
	}
L98:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v401)+40))
	if v411 != 0 {
		goto L93
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_XLogRegisterData(m, v18+int32(32), int32(4))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_XLogRegisterBuffer(m, int32(0), v421, int32(8))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	F_XLogRegisterBuffer(m, int32(1), v155, int32(6))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v431 = F_XLogInsert(m, int32(17), int32(64))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v433 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = base.I64_rotr(v431, v433)
	*(*uint32)(unsafe.Add(mBase, uint32(v156)+4)) = uint32(v431)
	v438 = int64(base.Ui64(v431) >> (uint(v433) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v156))) = uint32(v438)
	goto L93
L105:
	;
	F_UnlockReleaseBuffer(m, v448)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	goto L13
L107:
	;
	goto L6
L108:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	v510 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156+v510)+6)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if v155 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v532
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v513 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v512
	F_errmsg(m, int32(_a_F_brinRevmapExtend_8), v18)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L10
	} else {
		goto L114
	}
L111:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[3]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v517+(v155^int32(-1))<<(uint(int32(6))%32))+16))
	v532 = v523
	goto L110
L112:
	;
	goto L113
L113:
	;
	v525 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[4]))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v525+v155<<(uint(int32(6))%32)+int32(-64))+16))
	v532 = v531
	goto L110
L114:
	;
	F_errfinish(m, int32(_a_F_brinRevmapExtend_9), int32(586), int32(_a_F_brinRevmapExtend_10))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brin_bloom_opcinfo(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13853(m, l0, int32(_a_F_brin_bloom_opcinfo_0), int32(44))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	F_add_local_real_reloption(m, v2, int32(_a_F_brin_bloom_options_0), int32(_a_F_brin_bloom_options_1), float64(-0.1), float64(-1), float64(2.147483647e+09), int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_add_local_real_reloption(m, v2, int32(_a_F_brin_bloom_options_2), int32(_a_F_brin_bloom_options_3), float64(0.01), float64(0.0001), float64(0.25), int32(16))
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	if l3 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
		if v6 != 0 {
			if base.Ui32(l1) <= base.Ui32(v6) {
				v16 = l2
				if l1 != 0 {
					base.MemoryCopy(m, v16, l0, l1)
				} else {
				}
				return v16
			} else {
				v13 = F_repalloc(m, l2, l1)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = l1
					v16 = v13
					if l1 != 0 {
						base.MemoryCopy(m, v16, l0, l1)
					} else {
					}
					return v16
				}
			}
		} else {
			v8 = F_palloc(m, l1)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v16 = v8
				if l1 != 0 {
					base.MemoryCopy(m, v16, l0, l1)
				} else {
				}
				return v16
			}
		}
	} else {
		v8 = F_palloc(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v16 = v8
			if l1 != 0 {
				base.MemoryCopy(m, v16, l0, l1)
			} else {
			}
			return v16
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
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	if base.Ui32(l6) < base.Ui32(int32(_a_F_brin_doinsert_0)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L5
	} else {
		goto L109
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
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L105
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
	if v53 != int32(_a_F_brin_doinsert_1) {
		v68 = v8
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[0]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v32^int32(-1))<<(uint(int32(2))%32))))
	v50 = v42
	goto L13
L15:
	;
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[1]))
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
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(v120^int32(-1))<<(uint(int32(2))%32))))
	v138 = v130
	goto L33
L35:
	;
	goto L36
L36:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[1]))
	v138 = v132 + v120<<(uint(int32(13))%32) + int32(-8192)
	goto L33
L37:
	;
	v158 = int32(_a_F_brin_doinsert_2)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2]))
	v161 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2])) = v160 + v161
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+41)))
	if v164 == v161 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[3]))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v120^int32(-1))<<(uint(int32(6))%32))+16))
	v157 = v148
	goto L37
L39:
	;
	goto L40
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[4]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+v120<<(uint(int32(6))%32)+int32(-64))+16))
	v157 = v156
	goto L37
L41:
	;
	v167 = int32(_a_F_brin_doinsert_3)
	v169 = int32(0)
	if v169|(v138&int32(3)|int32(1)) == v169 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	goto L43
L43:
	;
	v221 = int32(0)
	v223 = F_PageAddItemExtended(m, v138, l5, l6, v221, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L55
	}
L44:
	;
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)))
	v219 = int32(_a_F_brin_doinsert_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v138+v217)+6)) = uint16(v219)
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+10)) = int32(_a_F_brin_doinsert_4)
	v208 = int32(_a_F_brin_doinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+18)) = uint16(v208)
	v214 = int32(_a_F_brin_doinsert_6)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)) = uint16(v214)
	*(*uint16)(unsafe.Add(mBase, uint32(v138)+14)) = uint16(v214)
	goto L44
L46:
	;
	goto L49
L47:
	;
	goto L48
L48:
	;
	goto L54
L49:
	;
	v185 = v138 + v167
	v187 = v138 + int32(4)
	if base.Ui32(v187) < base.Ui32(v185) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v189 = v185
	goto L52
L51:
	;
	v189 = v187
	goto L52
L52:
	;
	v194 = (v138^int32(-1)+v189)&int32(-4) + int32(4)
	if v194 == int32(0) {
		goto L45
	} else {
		goto L53
	}
L53:
	;
	base.MemoryFill(m, v138, int32(0), v194)
	goto L45
L54:
	;
	base.MemoryFill(m, v138, int32(0), v167)
	goto L45
L55:
	;
	if v223 == int32(0) {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_MarkBufferDirty(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	if v164 == int32(0) {
		v250 = v8
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+46)) = uint16(v223)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+24)) = uint16(v223)
	v254 = base.I32_rotr(v157, int32(16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+42)) = v254
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v254
	v258 = v18 + int32(20)
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258))))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+2)))
	if v118 < int32(0) {
		goto L68
	} else {
		goto L69
	}
L59:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)))
	v233 = v138 + v232
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v233)+6)))
	if v234 != int32(_a_F_brin_doinsert_1) {
		v250 = v8
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+4)))
	if v237&int32(1) != 0 {
		v250 = v8
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v240 = int32(4)
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+14)))
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v243 = v241 - v242
	if v243 <= v240 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v250 = v246 - int32(4)
	goto L58
L63:
	;
	v246 = v240
	goto L65
L64:
	;
	v246 = v243
	goto L65
L65:
	;
	goto L62
L66:
	;
	F_MarkBufferDirty(m, v118)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L77
	}
L67:
	;
	v281 = base.I32_div_u_s(l4, l1)
	v283 = base.I32_rem_u_s(v281, int32(1360))
	v286 = v280 + v283*int32(6)
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+28)) = uint16(v287)
	if v287 != 0 {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[0]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266+(v118^int32(-1))<<(uint(int32(2))%32))))
	v280 = v272
	goto L67
L69:
	;
	goto L70
L70:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[1]))
	v280 = v274 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L67
L71:
	;
	v290 = v262
	goto L73
L72:
	;
	v290 = int32(-1)
	goto L73
L73:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+26)) = uint16(v290)
	if v287 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v293 = v261
	goto L76
L75:
	;
	v293 = int32(-1)
	goto L76
L76:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v286)+24)) = uint16(v293)
	goto L66
L77:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+118)))
	if v298 != int32(112) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v369 = int32(_a_F_brin_doinsert_2)
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2])) = v371 - int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_LockBuffer(m, v375, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L98
	}
L79:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[5]))
	if v302 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v305 != 0 {
		goto L78
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+36)) = uint16(v223)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = l4
	F_XLogBeginInsert(m)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v306 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	F_XLogRegisterData(m, v18+int32(28), int32(10))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v164 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v321 = int32(14)
	goto L89
L88:
	;
	v321 = int32(8)
	goto L89
L89:
	;
	F_XLogRegisterBuffer(m, int32(0), v318, v321)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_XLogRegisterBufData(m, int32(0), l5, l6)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	F_XLogRegisterBuffer(m, int32(1), v118, int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	v336 = F_XLogInsert(m, int32(17), v164<<(uint(int32(7))%32)|int32(16))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	v338 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = base.I64_rotr(v336, v338)
	if v118 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362)+4)) = base.I32_wrap_i64(v336)
	*(*int32)(unsafe.Add(mBase, uint32(v362))) = base.I32_wrap_i64(int64(base.Ui64(v336) >> (uint(v338) % 64)))
	goto L78
L95:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[0]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348+(v118^int32(-1))<<(uint(int32(2))%32))))
	v362 = v354
	goto L94
L96:
	;
	goto L97
L97:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[1]))
	v362 = v356 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L94
L98:
	;
	F_LockBuffer(m, v118, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	if v164 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_RecordPageWithFreeSpace(m, l0, v157, v250)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	m.G0 = v18 + int32(48)
	return v223
L103:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v157, v157+int32(1))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(_a_F_brin_doinsert_7)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v399 + int32(4)
	F_errmsg(m, int32(_a_F_brin_doinsert_8), v18)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_brin_doinsert_9), int32(362), int32(_a_F_brin_doinsert_10))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	F_errmsg_internal(m, int32(_a_F_brin_doinsert_11), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_brin_doinsert_9), int32(414), int32(_a_F_brin_doinsert_10))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
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
	v32 = int32(_a_F_brin_fill_empty_ranges_0)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_brin_fill_empty_ranges[0]))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_fill_empty_ranges[0])) = v35
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
	*(*int32)(unsafe.Add(mBase, _c_F_brin_fill_empty_ranges[0])) = v33
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
func F_brin_free_desc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
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
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
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
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v16 = base.I32_extend16_s(v15)
	v17 = int32(4)
	v21 = v13 + v16<<(uint(v17)%32) + v17
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
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
	v29 = F_datumCopy(m, v23, v27, v28)
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
		v243 = v44
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
	return v243
L7:
	;
	v51 = v12 + v15<<(uint(int32(2))%32) + int32(16)
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
		goto L21
	} else {
		goto L22
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
	v105 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+115)) = uint8(v105)
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
	v95 = F_FunctionCall1Coll(m, v56, v22, v23)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
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
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+24)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v56)+20)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v56)+16)) = int32(0)
	goto L17
L17:
	;
	goto L13
L18:
	;
	if v95 == int32(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	if v100 != 0 {
		v243 = v44
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v101 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v101
	return v101
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+114)))
	if v112 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v243 = int32(1)
	goto L6
L24:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+113)))
	if v162 != 0 {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v114 = v111 + int32(56)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+60))
	if v115 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+114)) = uint8(v157)
	goto L24
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v118)+216))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v118)+204))
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+6)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v120+v122*(v16-int32(1))<<(uint(int32(2))%32)+int32(52)-int32(4))))
	goto L30
L28:
	;
	goto L29
L29:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = F_FunctionCall2Coll(m, v114, v22, v154, v23)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L34
	}
L30:
	;
	if v134 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v139 = F_index_getprocinfo(m, v137, v16, int32(13))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v139)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v114)+16)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v114)+8)) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v114)+16)) = int32(0)
	goto L33
L33:
	;
	goto L29
L34:
	;
	if v155 != 0 {
		v243 = v44
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	v217 = F_inclusion_get_procinfo(m, v12, v16&int32(_a_F_brin_inclusion_add_value_0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L48
	}
L37:
	;
	v164 = v161 + int32(28)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161)+32))
	if v165 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+113)) = uint8(v212)
	goto L36
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v168)+216))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)+204))
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+6)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v170+v172*(v16-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
	goto L42
L40:
	;
	goto L41
L41:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v205 = F_FunctionCall2Coll(m, v164, v22, v204, v23)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L46
	}
L42:
	;
	if v184 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L43:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v189 = F_index_getprocinfo(m, v187, v16, int32(12))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v189)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v164)+16)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v164)+24)) = v194
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v189)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v164)+8)) = v196
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v189)))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v164)+20)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v164)+16)) = int32(0)
	goto L45
L45:
	;
	goto L41
L46:
	;
	if v205 != 0 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v208 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v208
	return v208
L48:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v221 = F_FunctionCall2Coll(m, v217, v22, v220, v23)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
	if v223 != 0 {
		v234 = v221
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v236))) = v234
	goto L23
L51:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v221 == v225 {
		v234 = v221
		goto L50
	} else {
		goto L52
	}
L52:
	;
	F_pfree(m, v225)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	if v221 != v23 {
		v234 = v221
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+6)))
	v231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v232 = F_datumCopy(m, v23, v230, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v234 = v232
	goto L50
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v11 = v9 & int32(_a_F_brin_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v11)
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v13)+6)))
	switch v15 - int32(_a_F_brin_mask_1) {
	case 0:
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if base.Ui32(v18) < base.Ui32(int32(25)) {
			v24 = v13
			v25 = l0 + v24
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
			v28 = v26 & int32(_a_F_brin_mask_2)
			*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v28)
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
				v25 = l0 + v24
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
				v28 = v26 & int32(_a_F_brin_mask_2)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v28)
				return
			}
		}
	default:
		v24 = v13
		v25 = l0 + v24
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
		v28 = v26 & int32(_a_F_brin_mask_2)
		*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v28)
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
			v25 = l0 + v24
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
			v28 = v26 & int32(_a_F_brin_mask_2)
			*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v28)
			return
		}
	}
}
func F_brin_memtuple_initialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
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
	var v72 int32
	_ = v72
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_MemoryContextReset(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if int32(0) < v14 {
			v17 = int32(20)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = l0 + (v14*v17+int32(31))&int32(-8)
			v34 = int32(0)
			for {
				v40 = l0 + int32(24) + v34*int32(20)
				v42 = v34 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v42)
				v44 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v33
				v49 = int32(256)
				*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)) = uint16(v49)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v28
				v52 = int32(2)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1+v17+v34<<(uint(v52)%32))))
				v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55))))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v42 < v61 {
					v33 = v33 + v56<<(uint(v52)%32)
					v34 = v42
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		v72 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v72)
		return
	}
}
func F_brin_minmax_multi_add_value(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 float64
	_ = v259
	var v262 int32
	_ = v262
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v451 int32
	_ = v451
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v610 int32
	_ = v610
	var v618 int32
	_ = v618
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v924 int32
	_ = v924
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = F_get_fn_opclass_options(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30))))
	v46 = v36 + v37<<(uint(int32(4))%32) + v41*int32(100) - int32(80)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+3)))
	if v48 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v166 = v41 & int32(_a_F_brin_minmax_multi_add_value_0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v169 = F_minmax_multi_get_strategy_procinfo(m, v29, v166, v167, int32(1))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+180))
	if v52 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if v110 != 0 {
		v159 = v110
		v164 = v2
		goto L3
	} else {
		goto L30
	}
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v57 = v53 * int32(291)
	goto L9
L8:
	;
	v57 = int32(_a_F_brin_minmax_multi_add_value_1)
	goto L9
L9:
	;
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v63 = int32(32)
	goto L12
L12:
	;
	v64 = int32(_a_F_brin_minmax_multi_add_value_2)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v67
	v72 = v63 * int32(10)
	if base.Ui32(v72) < base.Ui32(v57) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v60 = v58
	goto L15
L14:
	;
	v60 = int32(32)
	goto L15
L15:
	;
	v63 = v60
	goto L12
L16:
	;
	v74 = v72
	goto L18
L17:
	;
	v74 = v57
	goto L18
L18:
	;
	if v63 < v74 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v76 = v74
	goto L21
L20:
	;
	v76 = v63
	goto L21
L21:
	;
	if v76 <= int32(256) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = int32(256)
	goto L24
L23:
	;
	v79 = v76
	goto L24
L24:
	;
	if int32(_a_F_brin_minmax_multi_add_value_3) <= v79 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = int32(_a_F_brin_minmax_multi_add_value_3)
	goto L27
L26:
	;
	v82 = v79
	goto L27
L27:
	;
	v87 = F_palloc0(m, v82<<(uint(int32(2))%32)+int32(36))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v87)+8)) = uint16(v41)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v47
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92
	v95 = int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v100 = F_minmax_multi_get_strategy_procinfo(m, v29, v41&int32(_a_F_brin_minmax_multi_add_value_0), v98, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = v100
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v87
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+3)) = uint8(v108)
	v159 = v87
	v164 = v95
	goto L3
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+180))
	if v112 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v117 = v113 * int32(291)
	goto L33
L32:
	;
	v117 = int32(_a_F_brin_minmax_multi_add_value_1)
	goto L33
L33:
	;
	v118 = int32(_a_F_brin_minmax_multi_add_value_2)
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v121
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	v127 = F_pg_detoast_datum(m, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
	v131 = v129 * int32(10)
	if base.Ui32(v131) < base.Ui32(v117) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v133 = v131
	goto L37
L36:
	;
	v133 = v117
	goto L37
L37:
	;
	if v129 < v133 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = v133
	goto L40
L39:
	;
	v135 = v129
	goto L40
L40:
	;
	if v135 <= int32(256) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v138 = int32(256)
	goto L43
L42:
	;
	v138 = v135
	goto L43
L43:
	;
	if int32(_a_F_brin_minmax_multi_add_value_3) <= v138 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v141 = int32(_a_F_brin_minmax_multi_add_value_3)
	goto L46
L45:
	;
	v141 = v138
	goto L46
L46:
	;
	v142 = F_brin_range_deserialize(m, v141, v127)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v47
	*(*uint16)(unsafe.Add(mBase, uint32(v142)+8)) = uint16(v41)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v146
	v151 = F_minmax_multi_get_strategy_procinfo(m, v29, v41&int32(_a_F_brin_minmax_multi_add_value_0), v146, int32(1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v142
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v119
	v159 = v142
	v164 = v2
	goto L3
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v175 = v171 + v172<<(uint(int32(1))%32)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	if v175 < v176 {
		v642 = v172
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v28
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v642 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L51:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	v180 = F_minmax_multi_get_strategy_procinfo(m, v29, v166, v178, int32(1))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_range_deduplicate_values(m, v159)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	if base.F64_le(base.F64_convert_i32_s(v184+v185<<(uint(int32(1))%32)), base.F64_mul(base.F64_convert_i32_s(v190), float64(0.5))) != 0 {
		v642 = v185
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v201 = F_AllocSetContextCreateInternal(m, v196, int32(_a_F_brin_minmax_multi_add_value_4), int32(0), int32(_a_F_brin_minmax_multi_add_value_3), int32(_a_F_brin_minmax_multi_add_value_5))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v203 = int32(_a_F_brin_minmax_multi_add_value_2)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v201
	v209 = F_build_expanded_ranges(m, v180, v47, v159, v26+int32(4))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v211 = F_minmax_multi_get_procinfo(m, v29, v166)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v213 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v217 = v213 - int32(1)
	v220 = F_palloc0(m, v217<<(uint(int32(4))%32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v295 = v2
	goto L60
L60:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v319 = F_reduce_expanded_ranges(m, v209, v213, v295, base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_s(v314), float64(0.5))), v180, v47)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L70
	}
L61:
	;
	if int32(0) < v217 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v226 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	F_pg_qsort(m, v220, v217, int32(16), int32(20))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L69
	}
L65:
	;
	v250 = v209 + v226*int32(12)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+4))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+12))
	v253 = F_FunctionCall2Coll(m, v211, v47, v251, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	v257 = v220 + v226<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v257))) = v226
	v259 = *(*float64)(unsafe.Add(mBase, uint32(v253)))
	*(*float64)(unsafe.Add(mBase, uint32(v257)+8)) = v259
	v262 = v226 + int32(1)
	if v262 != v217 {
		v226 = v262
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v295 = v220
	goto L60
L70:
	;
	v321 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v321
	if v319 <= v321 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = v618
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v204
	F_MemoryContextDelete(m, v201)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L106
	}
L72:
	;
	v610 = int32(0)
	goto L74
L73:
	;
	v327 = v159 + int32(36)
	v329 = v319 - int32(1)
	if v329 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v610
	v618 = v610
	goto L71
L75:
	;
	v473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v473
	if v329 == v473 {
		goto L92
	} else {
		goto L93
	}
L76:
	;
	v435 = v209 + v414*int32(12)
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+8)))
	if v436 != 0 {
		v451 = v411
		goto L75
	} else {
		goto L90
	}
L77:
	;
	v332 = int32(0)
	v411 = v332
	v414 = v332
	goto L76
L78:
	;
	goto L79
L79:
	;
	v338 = int32(0)
	v342 = v338
	v345 = v338
	v351 = v338
	goto L80
L80:
	;
	v366 = v209 + v345*int32(12)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+8)))
	if v367 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v319&int32(1) == int32(0) {
		v451 = v401
		goto L75
	} else {
		goto L89
	}
L82:
	;
	v370 = int32(2)
	v372 = v327 + v342<<(uint(v370)%32)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v366)))
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+4)) = v375
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v377 + int32(1)
	v383 = v342 + v370
	goto L84
L83:
	;
	v383 = v342
	goto L84
L84:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+20)))
	if v385 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v388 = int32(2)
	v390 = v327 + v383<<(uint(v388)%32)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v366)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = v391
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v366)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v390)+4)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v395 + int32(1)
	v401 = v383 + v388
	goto L87
L86:
	;
	v401 = v383
	goto L87
L87:
	;
	v403 = int32(2)
	v404 = v345 + v403
	v406 = v351 + v403
	if v406 != v319&int32(2147483646) {
		v342 = v401
		v345 = v404
		v351 = v406
		goto L80
	} else {
		goto L88
	}
L88:
	;
	goto L81
L89:
	;
	v411 = v401
	v414 = v404
	goto L76
L90:
	;
	v437 = int32(2)
	v439 = v327 + v411<<(uint(v437)%32)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v440
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v435)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+4)) = v442
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+16)) = v444 + int32(1)
	v451 = v411 + v437
	goto L75
L91:
	;
	v575 = v209 + v554*int32(12)
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+8)))
	if v576 != int32(1) {
		v618 = v556
		goto L71
	} else {
		goto L105
	}
L92:
	;
	v551 = v451
	v554 = int32(0)
	v556 = v473
	goto L91
L93:
	;
	goto L94
L94:
	;
	v483 = int32(0)
	v486 = v451
	v489 = v483
	v491 = v473
	v495 = v483
	goto L95
L95:
	;
	v510 = v209 + v489*int32(12)
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+8)))
	if v511 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v319&int32(1) == int32(0) {
		v618 = v542
		goto L71
	} else {
		goto L104
	}
L97:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	*(*int32)(unsafe.Add(mBase, uint32(v327+v486<<(uint(int32(2))%32)))) = v517
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v520 = int32(1)
	v521 = v519 + v520
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v521
	v525 = v486 + v520
	v526 = v521
	goto L99
L98:
	;
	v525 = v486
	v526 = v491
	goto L99
L99:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+20)))
	if v527 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v327+v525<<(uint(int32(2))%32)))) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v536 = int32(1)
	v537 = v535 + v536
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v537
	v541 = v525 + v536
	v542 = v537
	goto L102
L101:
	;
	v541 = v525
	v542 = v526
	goto L102
L102:
	;
	v543 = int32(2)
	v544 = v489 + v543
	v546 = v495 + v543
	if v546 != v319&int32(2147483646) {
		v486 = v541
		v489 = v544
		v491 = v542
		v495 = v546
		goto L95
	} else {
		goto L103
	}
L103:
	;
	goto L96
L104:
	;
	v551 = v541
	v554 = v544
	v556 = v542
	goto L91
L105:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	*(*int32)(unsafe.Add(mBase, uint32(v327+v551<<(uint(int32(2))%32)))) = v582
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v610 = v584 + int32(1)
	goto L74
L106:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v642 = v640
	goto L50
L107:
	;
	m.G0 = v26 + int32(16)
	return v924
L108:
	;
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+82)))
	v881 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+72)))
	v882 = F_datumCopy(m, v28, v880, v881)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L141
	}
L109:
	;
	v924 = v164 | base.B2i32(v176 <= v175)
	goto L107
L110:
	;
	v765 = F_minmax_multi_get_strategy_procinfo(m, v29, v41&int32(_a_F_brin_minmax_multi_add_value_0), v665, int32(3))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L129
	}
L111:
	;
	v669 = v159 + int32(36)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v669+v642<<(uint(int32(3))%32)-int32(4))))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v159)+36))
	v678 = v41 & int32(_a_F_brin_minmax_multi_add_value_0)
	v680 = F_minmax_multi_get_strategy_procinfo(m, v29, v678, v665, int32(1))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v682 = F_FunctionCall2Coll(m, v680, v47, v28, v676)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v682 != 0 {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v685 = F_minmax_multi_get_strategy_procinfo(m, v29, v678, v665, int32(5))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v687 = F_FunctionCall2Coll(m, v685, v47, v28, v675)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v687 != 0 {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v689 = int32(0)
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v692 = v690 - int32(1)
	if v692 < v689 {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	v696 = v692
	v697 = v692
	v699 = v689
	goto L119
L119:
	;
	v719 = base.I32_div_s(v697, int32(2))
	v722 = v669 + v719<<(uint(int32(3))%32)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+4))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	v725 = F_FunctionCall2Coll(m, v680, v47, v28, v724)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L110
L121:
	;
	if v736 <= v735 {
		v696 = v735
		v697 = v735 + v736
		v699 = v736
		goto L119
	} else {
		goto L128
	}
L122:
	;
	if v725 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v735 = v719 - int32(1)
	v736 = v699
	goto L121
L124:
	;
	goto L125
L125:
	;
	v729 = F_FunctionCall2Coll(m, v685, v47, v28, v723)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v729 == int32(0) {
		goto L109
	} else {
		goto L127
	}
L127:
	;
	v735 = v696
	v736 = v719 + int32(1)
	goto L121
L128:
	;
	goto L120
L129:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	if int32(16) <= v767 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v770
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v772
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v782 = int32(4)
	v786 = F_bsearch_arg(m, v26+int32(12), v159+v776<<(uint(int32(3))%32)+int32(36), v767, v782, int32(21), v26+v782)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v790 = v788 << (uint(int32(1)) % 32)
	if v790+v767 <= v790 {
		goto L108
	} else {
		goto L135
	}
L133:
	;
	if v786 != 0 {
		goto L109
	} else {
		goto L134
	}
L134:
	;
	goto L108
L135:
	;
	v796 = v790
	goto L136
L136:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v159+int32(36)+v796<<(uint(int32(2))%32))))
	v822 = F_FunctionCall2Coll(m, v765, v47, v28, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	goto L108
L138:
	;
	if v822 != 0 {
		goto L109
	} else {
		goto L139
	}
L139:
	;
	v824 = int32(1)
	v825 = v796 + v824
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v159)+20))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	if v825 < v826+v827<<(uint(v824)%32) {
		v796 = v825
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v159)+16))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v884<<(uint(int32(3))%32)+v888<<(uint(int32(2))%32))+36)) = v882
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v159)+24))
	v894 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+24)) = v893 + v894
	if v893 != 0 {
		v924 = v894
		goto L107
	} else {
		goto L142
	}
L142:
	;
	v898 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = v898
	v924 = v898
	goto L107
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v16 = v13 + int32(128)
	F_initStringInfo(m, v16)
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
	F_appendStringInfoChar(m, v16, int32(123))
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
	F_getTypeOutputInfo(m, v27, v13+int32(120), v13+int32(127))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
	F_fmgr_info(m, v34, v13+int32(92))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v42
	F_appendStringInfo(m, v16, int32(_a_F_brin_minmax_multi_summary_out_0), v13+int32(48))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v53 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v141 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v131 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v59 = int32(0)
	v61 = v59
	v63 = v59
	v66 = int32(0)
	goto L13
L13:
	;
	v72 = v13 + int32(76)
	F_initStringInfo(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v107 <= int32(0) {
		v131 = v94
		goto L9
	} else {
		goto L22
	}
L15:
	;
	v76 = v13 + int32(92)
	v79 = v40 + int32(36) + v61<<(uint(int32(2))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = F_OutputFunctionCall(m, v76, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v84 = F_OutputFunctionCall(m, v76, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v81
	F_appendStringInfo(m, v72, int32(_a_F_brin_minmax_multi_summary_out_1), v13+int32(32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v94 = v61 + int32(2)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v97 = F_cstring_to_text_with_len(m, v95, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v103 = F_accumArrayResult(m, v63, v97, int32(0), int32(25), v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v106 = v66 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	if v106 < v107 {
		v61 = v94
		v63 = v103
		v66 = v106
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	F_getTypeOutputInfo(m, int32(2277), v72, v13+int32(75))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v118 = F_makeArrayResult(m, v103, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v121 = F_OidOutputFunctionCall(m, v120, v118)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v121
	F_appendStringInfo(m, v13+int32(128), int32(_a_F_brin_minmax_multi_summary_out_2), v13+int32(16))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v131 = v94
	goto L9
L27:
	;
	F_appendStringInfoChar(m, v13+int32(128), int32(125))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L40
	}
L28:
	;
	v146 = int32(0)
	v148 = v131
	v150 = v146
	v153 = v146
	goto L29
L29:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(36)+v148<<(uint(int32(2))%32))))
	v165 = F_FunctionCall1Coll(m, v13+int32(92), int32(0), v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	if v179 <= int32(0) {
		goto L27
	} else {
		goto L35
	}
L31:
	;
	v167 = F_cstring_to_text(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v173 = F_accumArrayResult(m, v150, v167, int32(0), int32(25), v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v175 = int32(1)
	v178 = v153 + v175
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v40)+24))
	if v178 < v179 {
		v148 = v148 + v175
		v150 = v173
		v153 = v178
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	F_getTypeOutputInfo(m, int32(2277), v13+int32(76), v13+int32(75))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v192 = F_makeArrayResult(m, v173, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v13)+76))
	v195 = F_OidOutputFunctionCall(m, v194, v192)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v195
	F_appendStringInfo(m, v13+int32(128), int32(_a_F_brin_minmax_multi_summary_out_3), v13)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L27
L40:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v13)+128))
	m.G0 = v13 + int32(144)
	return v218
}
func F_brin_minmax_multi_summary_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_brin_minmax_multi_summary_recv_0), int32(3121), int32(_a_F_brin_minmax_multi_summary_recv_1), int32(_a_F_brin_minmax_multi_summary_recv_2), int32(_a_F_brin_minmax_multi_summary_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_new_memtuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = F_palloc0(m, (v11*int32(20)+int32(31))&int32(-8)+v18<<(uint(int32(2))%32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v29 = F_palloc(m, v26<<(uint(int32(2))%32))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v29
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			v34 = F_palloc(m, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v34
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				v39 = F_palloc(m, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v41)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v39
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_brin_new_memtuple[0]))
					v50 = F_AllocSetContextCreateInternal(m, v45, int32(_a_F_brin_new_memtuple_0), int32(0), int32(_a_F_brin_new_memtuple_1), int32(_a_F_brin_new_memtuple_2))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v50
						F_MemoryContextReset(m, v50)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
							if int32(0) < v56 {
								v59 = int32(20)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								v75 = v22 + (v56*v59+int32(31))&int32(-8)
								v76 = int32(0)
								for {
									v82 = v22 + int32(24) + v76*int32(20)
									v84 = v76 + int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v82))) = uint16(v84)
									v86 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v86
									*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v86
									*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v75
									v91 = int32(256)
									*(*uint16)(unsafe.Add(mBase, uint32(v82)+2)) = uint16(v91)
									*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v70
									v94 = int32(2)
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0+v59+v76<<(uint(v94)%32))))
									v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v97))))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
									if v84 < v103 {
										v75 = v75 + v98<<(uint(v94)%32)
										v76 = v84
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v114 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v114)
							return v22
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	v2 = l1
	v3 = int32(_a_F_brin_page_init_0)
	v5 = int32(0)
	if v5|(l0&int32(3)|int32(1)) == v5 {
		v21 = l0 + v3
		v23 = l0 + int32(4)
		if base.Ui32(v23) < base.Ui32(v21) {
			v25 = v21
		} else {
			v25 = v23
		}
		v30 = (l0^int32(-1)+v25)&int32(-4) + int32(4)
		if v30 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v30)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v3)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F_brin_page_init_1)
	v44 = int32(_a_F_brin_page_init_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v44)
	v50 = int32(_a_F_brin_page_init_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v50)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v50)
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v53)+6)) = uint16(v2)
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
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
	v131 = F_palloc0(m, v121)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
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
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v121 = int32(20)
	goto L4
L10:
	;
	goto L11
L11:
	;
	v40 = v2
	v41 = int32(20)
	goto L12
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)+v40<<(uint(int32(2))%32))))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v55 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v121 = v81
	goto L4
L14:
	;
	v81 = v80 + v41
	v83 = v40 + int32(1)
	if v83 != v25 {
		v40 = v83
		v41 = v81
		goto L12
	} else {
		goto L25
	}
L15:
	;
	v59 = int32(18)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v61 == v59 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v72 = int32(1)
	if v55&v72 != 0 {
		v80 = int32(base.Ui32(v55) >> (uint(v72) % 32))
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v64 = v59
	goto L20
L19:
	;
	v64 = int32(2)
	goto L20
L20:
	;
	if base.Ui32((v61-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = int32(6)
	goto L23
L22:
	;
	v71 = v64
	goto L23
L23:
	;
	v80 = v71
	goto L14
L24:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v80 = int32(base.Ui32(v76) >> (uint(int32(2)) % 32))
	goto L14
L25:
	;
	goto L13
L26:
	;
	if v25 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v121 = v29*v25 + int32(20)
	goto L4
L29:
	;
	v121 = int32(20)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v94 = v2
	v95 = int32(20)
	goto L32
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)+v94<<(uint(int32(2))%32))))
	v109 = F_strlen(m, v108)
	mBase = m.M
	v111 = int32(1)
	v112 = v109 + v95 + v111
	v114 = v94 + v111
	if v114 != v25 {
		v94 = v114
		v95 = v112
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v121 = v112
	goto L4
L34:
	;
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v121 << (uint(int32(2)) % 32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v137
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+12)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v141
	if int32(0) < v25 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v146 = l0 + int32(36)
	v157 = v131 + int32(20)
	v158 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	m.G0 = v15 + int32(16)
	return v131
L39:
	;
	if v27 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	v246 = v158 + int32(1)
	if v246 != v25 {
		v157 = v242
		v158 = v246
		goto L39
	} else {
		goto L81
	}
L42:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v146+v158<<(uint(int32(2))%32))))
	switch v29&int32(_a_F_brin_range_serialize_0) - int32(1) {
	case 0:
		goto L46
	case 1:
		goto L49
	default:
		goto L47
	case 3:
		goto L48
	}
L43:
	;
	goto L44
L44:
	;
	if int32(0) < v29 {
		goto L56
	} else {
		goto L57
	}
L45:
	;
	if v29 != 0 {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)) = uint8(v171)
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v171
	goto L45
L49:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)) = uint16(v171)
	goto L45
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v29
	F_errmsg_internal(m, int32(_a_F_brin_range_serialize_1), v15)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_brin_range_serialize_2), int32(230), int32(_a_F_brin_range_serialize_3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	base.MemoryCopy(m, v157, v15+int32(12), v29)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v242 = v157 + v29
	goto L41
L56:
	;
	if v29 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	goto L58
L58:
	;
	if v29 == int32(-1) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v146+v158<<(uint(int32(2))%32))))
	base.MemoryCopy(m, v157, v197, v29)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v242 = v157 + v29
	goto L41
L62:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v146+v158<<(uint(int32(2))%32))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
	if v206 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	goto L64
L64:
	;
	if v29 != int32(-2) {
		v242 = v157
		goto L41
	} else {
		goto L77
	}
L65:
	;
	if v230 != 0 {
		goto L74
	} else {
		goto L75
	}
L66:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	if base.Ui32((v210-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v230 = int32(6)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v222 = int32(1)
	if v206&v222 != 0 {
		v230 = int32(base.Ui32(v206) >> (uint(v222) % 32))
		goto L65
	} else {
		goto L73
	}
L69:
	;
	v217 = int32(18)
	if v210 == v217 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v221 = v217
	goto L72
L71:
	;
	v221 = int32(2)
	goto L72
L72:
	;
	v230 = v221
	goto L65
L73:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v230 = int32(base.Ui32(v226) >> (uint(int32(2)) % 32))
	goto L65
L74:
	;
	base.MemoryCopy(m, v157, v205, v230)
	goto L76
L75:
	;
	goto L76
L76:
	;
	v242 = v157 + v230
	goto L41
L77:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v146+v158<<(uint(int32(2))%32))))
	v237 = F_strlen(m, v236)
	mBase = m.M
	v239 = v237 + int32(1)
	if v239 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	base.MemoryCopy(m, v157, v236, v239)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v242 = v157 + v239
	goto L41
L81:
	;
	goto L40
}
func F_brin_xlog_insert_update(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
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
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
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
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(v15)+48)))
	if v16 < int32(0) {
		v20 = F_XLogInitBufferForRedo(m, l0, int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v20
			if v20 < int32(0) {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v20^int32(-1))<<(uint(int32(2))%32))))
				v40 = v32
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
				v40 = v34 + v20<<(uint(int32(13))%32) + int32(-8192)
			}
			v41 = int32(_a_F_brin_xlog_insert_update_0)
			F_PageInit(m, v40, int32(_a_F_brin_xlog_insert_update_1), int32(8))
			mBase = m.M
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
			*(*uint16)(unsafe.Add(mBase, uint32(v40+v45)+6)) = uint16(v41)
			if v20 < int32(0) {
				v51 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[2]))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v20^int32(-1))<<(uint(int32(6))%32))+16))
				v66 = v57
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[3]))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+v20<<(uint(int32(6))%32)+int32(-64))+16))
				v66 = v65
			}
			v92 = v66
			v93 = int32(0)
			v95 = v12 + int32(20)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+72))
			if v98 < v93 {
				v120 = v93
				v123 = v120
			} else {
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+int32(0))+76)))
				if v103 != int32(1) {
					v120 = v93
					v123 = v120
				} else {
					v107 = v97 + int32(76)
					v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+43)))
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
							v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+48)))
							*(*int32)(unsafe.Add(mBase, uint32(v95))) = v116
						} else {
						}
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
						v120 = v118
						v123 = v120
					}
				}
			}
			v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
			v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			if v126 < int32(0) {
				v130 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+(v126^int32(-1))<<(uint(int32(2))%32))))
				v144 = v136
			} else {
				v138 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
				v144 = v138 + v126<<(uint(int32(13))%32) + int32(-8192)
			}
			v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+12)))
			if base.Ui32(v145) < base.Ui32(int32(25)) {
				v156 = int32(1)
			} else {
				v156 = int32(base.Ui32(v145+int32(_a_F_brin_xlog_insert_update_2))>>(uint(int32(2))%32))&int32(_a_F_brin_xlog_insert_update_3) + int32(1)
			}
			if base.Ui32(v156) < base.Ui32(v124) {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v272 = m.ExcPending
				if v272 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_brin_xlog_insert_update_4), int32(0))
					mBase = m.M
					v276 = m.ExcPending
					if v276 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_brin_xlog_insert_update_5), int32(88), int32(_a_F_brin_xlog_insert_update_6))
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
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v160 = F_PageAddItemExtended(m, v144, v123, v158, v124, int32(1))
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
							F_errmsg_internal(m, int32(_a_F_brin_xlog_insert_update_7), int32(0))
							mBase = m.M
							v289 = m.ExcPending
							if v289 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_brin_xlog_insert_update_5), int32(92), int32(_a_F_brin_xlog_insert_update_6))
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
						*(*int64)(unsafe.Add(mBase, uint32(v144))) = base.I64_rotr(v14, int64(32))
						v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						F_MarkBufferDirty(m, v167)
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return
						} else {
							v171 = v92
							v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							if v175 != 0 {
								F_UnlockReleaseBuffer(m, v175)
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
										return
									} else {
										if v181 == int32(0) {
											v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
											v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if v191 < int32(0) {
												v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
												v209 = v201
											} else {
												v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
												v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
											}
											v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
											v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
											v217 = v12 + int32(12)
											v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
											v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
											if v191 < int32(0) {
												v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
												v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
												v239 = v231
											} else {
												v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
												v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
											}
											v240 = base.I32_div_u_s(v210, v211)
											v242 = base.I32_rem_u_s(v240, int32(1360))
											v245 = v239 + v242*int32(6)
											v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
											if v246 != 0 {
												v249 = v221
											} else {
												v249 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
											if v246 != 0 {
												v252 = v220
											} else {
												v252 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
											*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											F_MarkBufferDirty(m, v257)
											mBase = m.M
											v259 = m.ExcPending
											if v259 != 0 {
												return
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											}
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											} else {
												m.G0 = v12 + int32(32)
												return
											}
										}
									}
								}
							} else {
								v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
								mBase = m.M
								v182 = m.ExcPending
								if v182 != 0 {
									return
								} else {
									if v181 == int32(0) {
										v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
										v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
										if v191 < int32(0) {
											v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
											v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
											v209 = v201
										} else {
											v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
											v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
										}
										v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
										v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
										v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
										v217 = v12 + int32(12)
										v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
										v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
										if v191 < int32(0) {
											v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
											v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
											v239 = v231
										} else {
											v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
											v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
										}
										v240 = base.I32_div_u_s(v210, v211)
										v242 = base.I32_rem_u_s(v240, int32(1360))
										v245 = v239 + v242*int32(6)
										v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
										if v246 != 0 {
											v249 = v221
										} else {
											v249 = int32(-1)
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
										if v246 != 0 {
											v252 = v220
										} else {
											v252 = int32(-1)
										}
										*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
										*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
										v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
										F_MarkBufferDirty(m, v257)
										mBase = m.M
										v259 = m.ExcPending
										if v259 != 0 {
											return
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											} else {
												m.G0 = v12 + int32(32)
												return
											}
										}
									} else {
										v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
										if v263 != 0 {
											F_UnlockReleaseBuffer(m, v263)
											mBase = m.M
											v265 = m.ExcPending
											if v265 != 0 {
												return
											} else {
												m.G0 = v12 + int32(32)
												return
											}
										} else {
											m.G0 = v12 + int32(32)
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
		v70 = F_XLogReadBufferForRedo(m, l0, int32(0), v12+int32(28))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
			if v72 < int32(0) {
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[2]))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v72^int32(-1))<<(uint(int32(6))%32))+16))
				v91 = v82
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[3]))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+v72<<(uint(int32(6))%32)+int32(-64))+16))
				v91 = v90
			}
			if v70 != 0 {
				v171 = v91
				v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
				if v175 != 0 {
					F_UnlockReleaseBuffer(m, v175)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return
					} else {
						v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return
						} else {
							if v181 == int32(0) {
								v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
								*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
								v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
								*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
								v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								if v191 < int32(0) {
									v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
									v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
									v209 = v201
								} else {
									v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
									v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
								}
								v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
								*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
								v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
								v217 = v12 + int32(12)
								v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
								v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
								if v191 < int32(0) {
									v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
									v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
									v239 = v231
								} else {
									v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
									v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
								}
								v240 = base.I32_div_u_s(v210, v211)
								v242 = base.I32_rem_u_s(v240, int32(1360))
								v245 = v239 + v242*int32(6)
								v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
								if v246 != 0 {
									v249 = v221
								} else {
									v249 = int32(-1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
								if v246 != 0 {
									v252 = v220
								} else {
									v252 = int32(-1)
								}
								*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
								*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
								v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								F_MarkBufferDirty(m, v257)
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return
								} else {
									v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
									if v263 != 0 {
										F_UnlockReleaseBuffer(m, v263)
										mBase = m.M
										v265 = m.ExcPending
										if v265 != 0 {
											return
										} else {
											m.G0 = v12 + int32(32)
											return
										}
									} else {
										m.G0 = v12 + int32(32)
										return
									}
								}
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								if v263 != 0 {
									F_UnlockReleaseBuffer(m, v263)
									mBase = m.M
									v265 = m.ExcPending
									if v265 != 0 {
										return
									} else {
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									m.G0 = v12 + int32(32)
									return
								}
							}
						}
					}
				} else {
					v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return
					} else {
						if v181 == int32(0) {
							v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
							v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							if v191 < int32(0) {
								v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
								v209 = v201
							} else {
								v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
								v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
							}
							v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
							v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
							v217 = v12 + int32(12)
							v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
							v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
							if v191 < int32(0) {
								v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
								v239 = v231
							} else {
								v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
								v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
							}
							v240 = base.I32_div_u_s(v210, v211)
							v242 = base.I32_rem_u_s(v240, int32(1360))
							v245 = v239 + v242*int32(6)
							v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
							if v246 != 0 {
								v249 = v221
							} else {
								v249 = int32(-1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
							if v246 != 0 {
								v252 = v220
							} else {
								v252 = int32(-1)
							}
							*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
							*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
							v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							F_MarkBufferDirty(m, v257)
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return
							} else {
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								if v263 != 0 {
									F_UnlockReleaseBuffer(m, v263)
									mBase = m.M
									v265 = m.ExcPending
									if v265 != 0 {
										return
									} else {
										m.G0 = v12 + int32(32)
										return
									}
								} else {
									m.G0 = v12 + int32(32)
									return
								}
							}
						} else {
							v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							if v263 != 0 {
								F_UnlockReleaseBuffer(m, v263)
								mBase = m.M
								v265 = m.ExcPending
								if v265 != 0 {
									return
								} else {
									m.G0 = v12 + int32(32)
									return
								}
							} else {
								m.G0 = v12 + int32(32)
								return
							}
						}
					}
				}
			} else {
				v92 = v91
				v93 = int32(0)
				v95 = v12 + int32(20)
				v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+72))
				if v98 < v93 {
					v120 = v93
					v123 = v120
				} else {
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97+int32(0))+76)))
					if v103 != int32(1) {
						v120 = v93
						v123 = v120
					} else {
						v107 = v97 + int32(76)
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+43)))
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
								v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+48)))
								*(*int32)(unsafe.Add(mBase, uint32(v95))) = v116
							} else {
							}
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v107)+44))
							v120 = v118
							v123 = v120
						}
					}
				}
				v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
				if v126 < int32(0) {
					v130 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+(v126^int32(-1))<<(uint(int32(2))%32))))
					v144 = v136
				} else {
					v138 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
					v144 = v138 + v126<<(uint(int32(13))%32) + int32(-8192)
				}
				v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+12)))
				if base.Ui32(v145) < base.Ui32(int32(25)) {
					v156 = int32(1)
				} else {
					v156 = int32(base.Ui32(v145+int32(_a_F_brin_xlog_insert_update_2))>>(uint(int32(2))%32))&int32(_a_F_brin_xlog_insert_update_3) + int32(1)
				}
				if base.Ui32(v156) < base.Ui32(v124) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v272 = m.ExcPending
					if v272 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F_brin_xlog_insert_update_4), int32(0))
						mBase = m.M
						v276 = m.ExcPending
						if v276 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_brin_xlog_insert_update_5), int32(88), int32(_a_F_brin_xlog_insert_update_6))
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
					v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					v160 = F_PageAddItemExtended(m, v144, v123, v158, v124, int32(1))
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
								F_errmsg_internal(m, int32(_a_F_brin_xlog_insert_update_7), int32(0))
								mBase = m.M
								v289 = m.ExcPending
								if v289 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_brin_xlog_insert_update_5), int32(92), int32(_a_F_brin_xlog_insert_update_6))
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
							*(*int64)(unsafe.Add(mBase, uint32(v144))) = base.I64_rotr(v14, int64(32))
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							F_MarkBufferDirty(m, v167)
							mBase = m.M
							v169 = m.ExcPending
							if v169 != 0 {
								return
							} else {
								v171 = v92
								v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
								if v175 != 0 {
									F_UnlockReleaseBuffer(m, v175)
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
											return
										} else {
											if v181 == int32(0) {
												v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
												*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
												*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
												v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
												v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
												if v191 < int32(0) {
													v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
													v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
													v209 = v201
												} else {
													v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
													v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
												}
												v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
												v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
												*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
												*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
												v217 = v12 + int32(12)
												v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
												v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
												if v191 < int32(0) {
													v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
													v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
													v239 = v231
												} else {
													v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
													v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
												}
												v240 = base.I32_div_u_s(v210, v211)
												v242 = base.I32_rem_u_s(v240, int32(1360))
												v245 = v239 + v242*int32(6)
												v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
												*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
												if v246 != 0 {
													v249 = v221
												} else {
													v249 = int32(-1)
												}
												*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
												if v246 != 0 {
													v252 = v220
												} else {
													v252 = int32(-1)
												}
												*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
												*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
												v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
												F_MarkBufferDirty(m, v257)
												mBase = m.M
												v259 = m.ExcPending
												if v259 != 0 {
													return
												} else {
													v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
													if v263 != 0 {
														F_UnlockReleaseBuffer(m, v263)
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return
														} else {
															m.G0 = v12 + int32(32)
															return
														}
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												}
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											}
										}
									}
								} else {
									v181 = F_XLogReadBufferForRedo(m, l0, int32(1), v12+int32(28))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
										return
									} else {
										if v181 == int32(0) {
											v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)) = uint16(v185)
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+22)) = uint16(v171)
											v189 = int32(base.Ui32(v171) >> (uint(int32(16)) % 32))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+20)) = uint16(v189)
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if v191 < int32(0) {
												v195 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
												v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+(v191^int32(-1))<<(uint(int32(2))%32))))
												v209 = v201
											} else {
												v203 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
												v209 = v203 + v191<<(uint(int32(13))%32) + int32(-8192)
											}
											v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
											v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+24)))
											*(*uint16)(unsafe.Add(mBase, uint32(v12)+16)) = uint16(v212)
											v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v214
											v217 = v12 + int32(12)
											v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217))))
											v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+2)))
											if v191 < int32(0) {
												v225 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[0]))
												v231 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v191^int32(-1))<<(uint(int32(2))%32))))
												v239 = v231
											} else {
												v233 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[1]))
												v239 = v233 + v191<<(uint(int32(13))%32) + int32(-8192)
											}
											v240 = base.I32_div_u_s(v210, v211)
											v242 = base.I32_rem_u_s(v240, int32(1360))
											v245 = v239 + v242*int32(6)
											v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)))
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+28)) = uint16(v246)
											if v246 != 0 {
												v249 = v221
											} else {
												v249 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+26)) = uint16(v249)
											if v246 != 0 {
												v252 = v220
											} else {
												v252 = int32(-1)
											}
											*(*uint16)(unsafe.Add(mBase, uint32(v245)+24)) = uint16(v252)
											*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotr(v14, int64(32))
											v257 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											F_MarkBufferDirty(m, v257)
											mBase = m.M
											v259 = m.ExcPending
											if v259 != 0 {
												return
											} else {
												v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
												if v263 != 0 {
													F_UnlockReleaseBuffer(m, v263)
													mBase = m.M
													v265 = m.ExcPending
													if v265 != 0 {
														return
													} else {
														m.G0 = v12 + int32(32)
														return
													}
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											}
										} else {
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
											if v263 != 0 {
												F_UnlockReleaseBuffer(m, v263)
												mBase = m.M
												v265 = m.ExcPending
												if v265 != 0 {
													return
												} else {
													m.G0 = v12 + int32(32)
													return
												}
											} else {
												m.G0 = v12 + int32(32)
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
