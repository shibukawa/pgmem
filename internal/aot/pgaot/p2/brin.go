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
	var v33 int32
	_ = v33
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
				F_LockBufferInternal(m, v54, int32(3))
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
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v18^int32(-1))*int32(56))+16))
				v39 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_brinLockRevmapPageForUpdate[1]))
				v33 = int32(56)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v18*v33-v33)+16))
				v39 = v38
			}
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v17 == v39 {
				v54 = v40
				F_LockBufferInternal(m, v54, int32(3))
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
						F_LockBufferInternal(m, v54, int32(3))
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
							F_LockBufferInternal(m, v54, int32(3))
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
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
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
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
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int64
	_ = v428
	var v429 int32
	_ = v429
	var v431 int64
	_ = v431
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v476 int32
	_ = v476
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
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
	v499 = m.ExcPending
	if v499 != 0 {
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
	F_LockBufferInternal(m, v46, int32(3))
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
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v476) <= base.Ui32(v23) {
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
	F_UnlockBuffer(m, v50)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v76 = v69 + int32(1)
	v78 = F_RelationGetNumberOfBlocksInFork(m, v45, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L24
	}
L21:
	;
	goto L13
L22:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_UnlockBuffer(m, v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L10
	} else {
		goto L105
	}
L23:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+14)))
	if v156 != 0 {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	if base.Ui32(v76) < base.Ui32(v78) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = F_ReadBuffer(m, v45, v76)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v108
	v112 = int32(0)
	v115 = F_ExtendBufferedRel(m, v18+int32(16), v112, v112, int32(8))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L33
	}
L28:
	;
	F_LockBufferInternal(m, v81, int32(3))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v81 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(v81^int32(-1))<<(uint(int32(2))%32))))
	v154 = v81
	v155 = v95
	goto L23
L31:
	;
	goto L32
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v154 = v81
	v155 = v97 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L33:
	;
	if v115 < int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v135 != v76 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[3]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+(v115^int32(-1))*int32(56))+16))
	v135 = v126
	goto L34
L36:
	;
	goto L37
L37:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[4]))
	v129 = int32(56)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128+v115*v129-v129)+16))
	v135 = v134
	goto L34
L38:
	;
	v442 = v115
	goto L22
L39:
	;
	goto L40
L40:
	;
	if v115 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140+(v115^int32(-1))<<(uint(int32(2))%32))))
	v154 = v115
	v155 = v146
	goto L23
L42:
	;
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v154 = v115
	v155 = v148 + v115<<(uint(int32(13))%32) + int32(-8192)
	goto L23
L44:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155+v157)+6)))
	if v159 != int32(_a_F_brinRevmapExtend_0) {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v162 = int32(0)
	if v154 < v162 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	goto L46
L48:
	;
	if v247 != 0 {
		goto L63
	} else {
		goto L64
	}
L49:
	;
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+14)))
	if v181 == int32(0) {
		v247 = v162
		goto L48
	} else {
		goto L53
	}
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166+(v154^int32(-1))<<(uint(int32(2))%32))))
	v180 = v172
	goto L49
L51:
	;
	goto L52
L52:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v180 = v174 + v154<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+12)))
	if base.Ui32(v184) < base.Ui32(int32(25)) {
		v247 = v162
		goto L48
	} else {
		goto L54
	}
L54:
	;
	v192 = int32(base.Ui32(v184+int32(_a_F_brinRevmapExtend_1))>>(uint(int32(2))%32)) & int32(_a_F_brinRevmapExtend_2)
	if v192 == int32(0) {
		v247 = v162
		goto L48
	} else {
		goto L55
	}
L55:
	;
	v200 = int32(1)
	goto L56
L56:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180+int32(20)+v200&int32(_a_F_brinRevmapExtend_2)<<(uint(int32(2))%32))+1)))
	if v218&int32(384) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v228 = int32(1)
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+16)))
	v230 = v180 + v229
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)))
	v233 = v231 | v228
	*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)) = uint16(v233)
	F_MarkBufferDirtyHint(m, v154, v228)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L62
	}
L58:
	;
	v224 = v200 + int32(1)
	if base.Ui32(v224&int32(_a_F_brinRevmapExtend_2)) <= base.Ui32(v192) {
		v200 = v224
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
	v247 = v162
	goto L48
L62:
	;
	v247 = v228
	goto L48
L63:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_UnlockBuffer(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L10
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v377 = int32(_a_F_brinRevmapExtend_3)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5])) = v379 + int32(1)
	v383 = int32(_a_F_brinRevmapExtend_4)
	F_PageInit(m, v155, int32(_a_F_brinRevmapExtend_5), int32(8))
	mBase = m.M
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v155+v387)+6)) = uint16(v383)
	goto L90
L66:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v257 = m.G0
	v259 = v257 - int32(16)
	m.G0 = v259
	v261 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+12)) = v261
	if v154 < v261 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	F_UnlockReleaseBuffer(m, v154)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L10
	} else {
		goto L89
	}
L68:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+12)))
	if base.Ui32(v281) < base.Ui32(int32(25)) {
		goto L67
	} else {
		goto L72
	}
L69:
	;
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[1]))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266+(v154^int32(-1))<<(uint(int32(2))%32))))
	v280 = v272
	goto L68
L70:
	;
	goto L71
L71:
	;
	v274 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[2]))
	v280 = v274 + v154<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	v289 = int32(base.Ui32(v281+int32(_a_F_brinRevmapExtend_1))>>(uint(int32(2))%32)) & int32(_a_F_brinRevmapExtend_2)
	if v289 == int32(0) {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v297 = int32(1)
	goto L74
L74:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[0]))
	if v311 != 0 {
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
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v315 = v297 & int32(_a_F_brinRevmapExtend_2)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v280+int32(20)+v315<<(uint(int32(2))%32))))
	if v319&int32(_a_F_brinRevmapExtend_6) != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L78
L80:
	;
	v326 = int32(base.Ui32(v319) >> (uint(int32(17)) % 32))
	v330 = F_brin_copy_tuple(m, v280+v319&int32(_a_F_brinRevmapExtend_7), v326, int32(0), v259+int32(12))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	v351 = v297
	goto L82
L82:
	;
	v353 = v351 + int32(1)
	if base.Ui32(v353&int32(_a_F_brinRevmapExtend_2)) <= base.Ui32(v289) {
		v297 = v353
		goto L74
	} else {
		goto L88
	}
L83:
	;
	F_UnlockBuffer(m, v154)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v336 = F_brin_doupdate(m, v45, v256, l0, v334, v154, v315, v330, v326, v330, v326, int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	F_LockBufferInternal(m, v154, int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280)+16)))
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280+v341)+6)))
	if v343 != int32(_a_F_brinRevmapExtend_0) {
		goto L67
	} else {
		goto L87
	}
L87:
	;
	v351 = v297 - (v336 ^ int32(1))
	goto L82
L88:
	;
	goto L75
L89:
	;
	m.G0 = v259 + int32(16)
	goto L13
L90:
	;
	F_MarkBufferDirty(m, v154)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v392 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)) = uint16(v392)
	*(*int32)(unsafe.Add(mBase, uint32(v68)+36)) = v76
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MarkBufferDirty(m, v395)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v398)+48))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+118)))
	if v400 != int32(112) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v435 = int32(_a_F_brinRevmapExtend_3)
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[5])) = v437 - int32(1)
	v442 = v154
	goto L22
L94:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[6]))
	if v404 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v398)+32))
	if v407 != 0 {
		goto L93
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v76
	F_XLogBeginInsert(m)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L100
	}
L98:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v398)+40))
	if v408 != 0 {
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
	v416 = m.ExcPending
	if v416 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_XLogRegisterBuffer(m, int32(0), v418, int32(8))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L10
	} else {
		goto L102
	}
L102:
	;
	F_XLogRegisterBuffer(m, int32(1), v154, int32(6))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v428 = F_XLogInsert(m, int32(17), int32(64))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	v431 = base.I64_rotl(v428, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v431
	*(*int64)(unsafe.Add(mBase, uint32(v155))) = v431
	goto L93
L105:
	;
	F_UnlockReleaseBuffer(m, v442)
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
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
	v502 = m.ExcPending
	if v502 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	v503 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+16)))
	v505 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155+v503)+6)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
	if v154 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v506 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v505
	F_errmsg(m, int32(_a_F_brinRevmapExtend_8), v18)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L10
	} else {
		goto L114
	}
L111:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[3]))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510+(v154^int32(-1))*int32(56))+16))
	v525 = v516
	goto L110
L112:
	;
	goto L113
L113:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_brinRevmapExtend[4]))
	v519 = int32(56)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v518+v154*v519-v519)+16))
	v525 = v524
	goto L110
L114:
	;
	F_errfinish(m, int32(_a_F_brinRevmapExtend_9), int32(586), int32(_a_F_brinRevmapExtend_10))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
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
func F_brin_bloom_opcinfo(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14234(m, l0, int32(_a_F_brin_bloom_opcinfo_0), int32(44))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_brin_bloom_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(24)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_real_reloption(m, v2, int32(_a_F_brin_bloom_options_0), int32(_a_F_brin_bloom_options_1), float64(-0.1), float64(-1), float64(2.147483647e+09), int32(8))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		F_add_local_real_reloption(m, v2, int32(_a_F_brin_bloom_options_2), int32(_a_F_brin_bloom_options_3), float64(0.01), float64(0.0001), float64(0.25), int32(16))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			return int64(0)
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
	var v151 int32
	_ = v151
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
	var v339 int64
	_ = v339
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
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
	if base.Ui32(l6) < base.Ui32(int32(_a_F_brin_doinsert_0)) {
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
	v385 = m.ExcPending
	if v385 != 0 {
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
	F_LockBufferInternal(m, v26, int32(3))
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v142+(v120^int32(-1))*int32(56))+16))
	v157 = v148
	goto L37
L39:
	;
	goto L40
L40:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[4]))
	v151 = int32(56)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+v120*v151-v151)+16))
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
	v361 = int32(_a_F_brin_doinsert_2)
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[2])) = v363 - int32(1)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_UnlockBuffer(m, v367)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
	v339 = base.I64_rotl(v336, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v339
	if v118 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v358))) = v339
	goto L78
L95:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[0]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v344+(v118^int32(-1))<<(uint(int32(2))%32))))
	v358 = v350
	goto L94
L96:
	;
	goto L97
L97:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_brin_doinsert[1]))
	v358 = v352 + v118<<(uint(int32(13))%32) + int32(-8192)
	goto L94
L98:
	;
	F_UnlockBuffer(m, v118)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
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
	v373 = m.ExcPending
	if v373 != 0 {
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
	v377 = m.ExcPending
	if v377 != 0 {
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
	v388 = m.ExcPending
	if v388 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(_a_F_brin_doinsert_7)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v389 + int32(4)
	F_errmsg(m, int32(_a_F_brin_doinsert_8), v18)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_brin_doinsert_9), int32(361), int32(_a_F_brin_doinsert_10))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
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
	v411 = m.ExcPending
	if v411 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_brin_doinsert_9), int32(412), int32(_a_F_brin_doinsert_10))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
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
func F_brin_inclusion_add_value(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v201 int32
	_ = v201
	var v203 int64
	_ = v203
	var v205 int64
	_ = v205
	var v210 int32
	_ = v210
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v252 int64
	_ = v252
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
	v16 = base.I32_extend16_s(v15)
	v21 = v13 + v16<<(uint(int32(3))%32) + int32(20)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)))
	if v24 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+2)))
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
	v44 = int64(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+8))
	if v46 != v44 {
		v252 = v44
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int64(0)
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v29
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v36 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v36
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v36
	v41 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+3)) = uint8(v41)
	goto L3
L6:
	;
	return v252
L7:
	;
	v53 = v12 + v15<<(uint(int32(2))%32) + int32(16)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+115)))
	if v56 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if v24 != 0 {
		v252 = int64(1)
		goto L6
	} else {
		goto L21
	}
L9:
	;
	v58 = v55 + int32(84)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+88))
	if v59 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v109 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55)+115)) = uint8(v109)
	goto L8
L11:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+216))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+204))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v65)+6)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v64+v66*(v16-int32(1))<<(uint(int32(2))%32)+int32(56)-int32(4))))
	goto L14
L12:
	;
	goto L13
L13:
	;
	v97 = F_FunctionCall1Coll(m, v58, v22, v23)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L18
	}
L14:
	;
	if v78 == int32(0) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v83 = F_index_getprocinfo(m, v81, v16, int32(14))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+16)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v58)+8)) = v90
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	*(*int64)(unsafe.Add(mBase, uint32(v58))) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = int32(0)
	goto L17
L17:
	;
	goto L13
L18:
	;
	if v97 == int64(0) {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v101)+16))
	if v102 != int64(0) {
		v252 = v44
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v105 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v101)+16)) = v105
	return v105
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+114)))
	if v115 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+113)))
	if v169 != 0 {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	v117 = v114 + int32(56)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
	if v118 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+114)) = uint8(v163)
	goto L22
L25:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)+216))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+204))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+6)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v123+v125*(v16-int32(1))<<(uint(int32(2))%32)+int32(52)-int32(4))))
	goto L28
L26:
	;
	goto L27
L27:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v157)))
	v159 = F_FunctionCall2Coll(m, v117, v22, v158, v23)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L32
	}
L28:
	;
	if v137 == int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v142 = F_index_getprocinfo(m, v140, v16, int32(13))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v142)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v117)+16)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = v147
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v142)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v117)+8)) = v149
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v117)+20)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = int32(0)
	goto L31
L31:
	;
	goto L27
L32:
	;
	if v159 != int64(0) {
		v252 = int64(0)
		goto L6
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	v226 = F_inclusion_get_procinfo(m, v12, v16&int32(_a_F_brin_inclusion_add_value_0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L46
	}
L35:
	;
	v171 = v168 + int32(28)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168)+32))
	if v172 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+113)) = uint8(v221)
	goto L34
L37:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v175)+216))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+204))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v178)+6)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v177+v179*(v16-int32(1))<<(uint(int32(2))%32)+int32(48)-int32(4))))
	goto L40
L38:
	;
	goto L39
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v211 = *(*int64)(unsafe.Add(mBase, uint32(v210)))
	v212 = F_FunctionCall2Coll(m, v171, v22, v211, v23)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L44
	}
L40:
	;
	if v191 == int32(0) {
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v196 = F_index_getprocinfo(m, v194, v16, int32(12))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v196)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v171)+16)) = v199
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+24)) = v201
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v203
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v205
	*(*int32)(unsafe.Add(mBase, uint32(v171)+20)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v171)+16)) = int32(0)
	goto L43
L43:
	;
	goto L39
L44:
	;
	if v212 != int64(0) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v217 = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v216)+8)) = v217
	return v217
L46:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v228)))
	v230 = F_FunctionCall2Coll(m, v226, v22, v229, v23)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v232 != 0 {
		v245 = v230
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v246))) = v245
	v252 = int64(1)
	goto L6
L49:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v234 == base.I32_wrap_i64(v230) {
		v245 = v230
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	if v230 != v23 {
		v245 = v230
		goto L48
	} else {
		goto L52
	}
L52:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	v241 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+2)))
	v242 = F_datumCopy(m, v23, v240, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v245 = v242
	goto L48
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
	var v19 int32
	_ = v19
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
	var v49 int32
	_ = v49
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
			v19 = int32(24)
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = l0 + (v14*v19+int32(31))&int32(-8)
			v34 = int32(0)
			for {
				v40 = l0 + v19 + v34*int32(24)
				v42 = v34 + int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v40))) = uint16(v42)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v40)+8)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v33
				v49 = int32(256)
				*(*uint16)(unsafe.Add(mBase, uint32(v40)+2)) = uint16(v49)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v28
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(20)+v34<<(uint(int32(2))%32))))
				v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55))))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
				if v42 < v61 {
					v33 = v33 + v56<<(uint(int32(3))%32)
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
func F_brin_minmax_multi_add_value(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int64
	_ = v401
	var v403 int64
	_ = v403
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int64
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
	var v549 int64
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v600 int64
	_ = v600
	var v602 int32
	_ = v602
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v699 int64
	_ = v699
	var v700 int64
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int64
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int64
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int64
	_ = v753
	var v754 int64
	_ = v754
	var v755 int64
	_ = v755
	var v756 int32
	_ = v756
	var v761 int64
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v855 int64
	_ = v855
	var v856 int64
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int64
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v938 int64
	_ = v938
	var v962 int64
	_ = v962
	v22 = int64(0)
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_get_fn_opclass_options(m, v33)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32))))
	v48 = v38 + v39<<(uint(int32(3))%32) + v43*int32(100) - int32(72)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)))
	if v50 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v170 = v43 & int32(_a_F_brin_minmax_multi_add_value_0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	v173 = F_minmax_multi_get_strategy_procinfo(m, v31, v170, v171, int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L49
	}
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+180))
	if v54 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	if v113 != 0 {
		v163 = v113
		v168 = v22
		goto L3
	} else {
		goto L30
	}
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v59 = v55 * int32(291)
	goto L9
L8:
	;
	v59 = int32(_a_F_brin_minmax_multi_add_value_1)
	goto L9
L9:
	;
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v60 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v65 = int32(32)
	goto L12
L12:
	;
	v66 = int32(_a_F_brin_minmax_multi_add_value_2)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v69
	v74 = v65 * int32(10)
	if base.Ui32(v74) < base.Ui32(v59) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v62 = v60
	goto L15
L14:
	;
	v62 = int32(32)
	goto L15
L15:
	;
	v65 = v62
	goto L12
L16:
	;
	v76 = v74
	goto L18
L17:
	;
	v76 = v59
	goto L18
L18:
	;
	if v65 < v76 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v78 = v76
	goto L21
L20:
	;
	v78 = v65
	goto L21
L21:
	;
	if v78 <= int32(256) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v81 = int32(256)
	goto L24
L23:
	;
	v81 = v78
	goto L24
L24:
	;
	if int32(_a_F_brin_minmax_multi_add_value_3) <= v81 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v84 = int32(_a_F_brin_minmax_multi_add_value_3)
	goto L27
L26:
	;
	v84 = v81
	goto L27
L27:
	;
	v89 = F_palloc0(m, v84<<(uint(int32(3))%32)+int32(40))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v89)+8)) = uint16(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v49
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v94
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	v101 = F_minmax_multi_get_strategy_procinfo(m, v31, v43&int32(_a_F_brin_minmax_multi_add_value_0), v99, int32(1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(20)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = base.I64_extend_i32_u(v89)
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)) = uint8(v110)
	v163 = v89
	v168 = int64(1)
	goto L3
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+180))
	if v115 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v120 = v116 * int32(291)
	goto L33
L32:
	;
	v120 = int32(_a_F_brin_minmax_multi_add_value_1)
	goto L33
L33:
	;
	v121 = int32(_a_F_brin_minmax_multi_add_value_2)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v124
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = F_pg_detoast_datum(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v134 = v132 * int32(10)
	if base.Ui32(v134) < base.Ui32(v120) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v136 = v134
	goto L37
L36:
	;
	v136 = v120
	goto L37
L37:
	;
	if v132 < v136 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v138 = v136
	goto L40
L39:
	;
	v138 = v132
	goto L40
L40:
	;
	if v138 <= int32(256) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = int32(256)
	goto L43
L42:
	;
	v141 = v138
	goto L43
L43:
	;
	if int32(_a_F_brin_minmax_multi_add_value_3) <= v141 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v144 = int32(_a_F_brin_minmax_multi_add_value_3)
	goto L46
L45:
	;
	v144 = v141
	goto L46
L46:
	;
	v145 = F_brin_range_deserialize(m, v144, v130)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v49
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+8)) = uint16(v43)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v149
	v154 = F_minmax_multi_get_strategy_procinfo(m, v31, v43&int32(_a_F_brin_minmax_multi_add_value_0), v149, int32(1))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = int32(20)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = base.I64_extend_i32_u(v145)
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v122
	v163 = v145
	v168 = v22
	goto L3
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v179 = v175 + v176<<(uint(int32(1))%32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	if v179 < v180 {
		v664 = v176
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v30
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	if v664 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L51:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
	v184 = F_minmax_multi_get_strategy_procinfo(m, v31, v170, v182, int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_range_deduplicate_values(m, v163)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	if base.F64_le(base.F64_convert_i32_s(v188+v189<<(uint(int32(1))%32)), base.F64_mul(base.F64_convert_i32_s(v194), float64(0.5))) != 0 {
		v664 = v189
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	v205 = F_AllocSetContextCreateInternal(m, v200, int32(_a_F_brin_minmax_multi_add_value_4), int32(0), int32(_a_F_brin_minmax_multi_add_value_3), int32(_a_F_brin_minmax_multi_add_value_5))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v207 = int32(_a_F_brin_minmax_multi_add_value_2)
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v205
	v213 = F_build_expanded_ranges(m, v184, v49, v163, v28+int32(8))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v215 = F_minmax_multi_get_procinfo(m, v31, v170)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v217 != int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v222 = v217 - int32(1)
	v223 = F_palloc0_mul(m, int32(16), v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v301 = int32(0)
	goto L60
L60:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v163)+28))
	v327 = F_reduce_expanded_ranges(m, v213, v217, v301, base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_s(v322), float64(0.5))), v184, v49)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L70
	}
L61:
	;
	if int32(0) < v222 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v229 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	F_pg_qsort(m, v223, v222, int32(16), int32(21))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L69
	}
L65:
	;
	v255 = v223 + v229<<(uint(int32(4))%32)
	v258 = v213 + v229*int32(24)
	v259 = *(*int64)(unsafe.Add(mBase, uint32(v258)+8))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v258)+24))
	v261 = F_FunctionCall2Coll(m, v215, v49, v259, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v255)+8)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v229
	v266 = v229 + int32(1)
	if v266 != v222 {
		v229 = v266
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v301 = v223
	goto L60
L70:
	;
	v329 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v329
	if v327 <= v329 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = v638
	*(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_add_value[0])) = v208
	F_MemoryContextDelete(m, v205)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L106
	}
L72:
	;
	v630 = int32(0)
	goto L74
L73:
	;
	v335 = v163 + int32(40)
	v337 = v327 - int32(1)
	if v337 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v630
	v638 = v630
	goto L71
L75:
	;
	v487 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v487
	if v337 == v487 {
		goto L92
	} else {
		goto L93
	}
L76:
	;
	v447 = v213 + v424*int32(24)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+16)))
	if v448 != 0 {
		v463 = v421
		goto L75
	} else {
		goto L90
	}
L77:
	;
	v340 = int32(0)
	v421 = v340
	v424 = v340
	goto L76
L78:
	;
	goto L79
L79:
	;
	v346 = int32(0)
	v350 = v346
	v353 = v346
	v359 = v346
	goto L80
L80:
	;
	v376 = v213 + v353*int32(24)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+16)))
	if v377 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	if v327&int32(1) == int32(0) {
		v463 = v411
		goto L75
	} else {
		goto L89
	}
L82:
	;
	v382 = v335 + v350<<(uint(int32(3))%32)
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v376)))
	*(*int64)(unsafe.Add(mBase, uint32(v382))) = v383
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v376)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v382)+8)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v387 + int32(1)
	v393 = v350 + int32(2)
	goto L84
L83:
	;
	v393 = v350
	goto L84
L84:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+40)))
	if v395 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v400 = v335 + v393<<(uint(int32(3))%32)
	v401 = *(*int64)(unsafe.Add(mBase, uint32(v376)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v400))) = v401
	v403 = *(*int64)(unsafe.Add(mBase, uint32(v376)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v400)+8)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v405 + int32(1)
	v411 = v393 + int32(2)
	goto L87
L86:
	;
	v411 = v393
	goto L87
L87:
	;
	v413 = int32(2)
	v414 = v353 + v413
	v416 = v359 + v413
	if v416 != v327&int32(2147483646) {
		v350 = v411
		v353 = v414
		v359 = v416
		goto L80
	} else {
		goto L88
	}
L88:
	;
	goto L81
L89:
	;
	v421 = v411
	v424 = v414
	goto L76
L90:
	;
	v451 = v335 + v421<<(uint(int32(3))%32)
	v452 = *(*int64)(unsafe.Add(mBase, uint32(v447)))
	*(*int64)(unsafe.Add(mBase, uint32(v451))) = v452
	v454 = *(*int64)(unsafe.Add(mBase, uint32(v447)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v451)+8)) = v454
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v456 + int32(1)
	v463 = v421 + int32(2)
	goto L75
L91:
	;
	v593 = v213 + v570*int32(24)
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v593)+16)))
	if v594 != int32(1) {
		v638 = v572
		goto L71
	} else {
		goto L105
	}
L92:
	;
	v567 = v463
	v570 = int32(0)
	v572 = v487
	goto L91
L93:
	;
	goto L94
L94:
	;
	v497 = int32(0)
	v500 = v463
	v503 = v497
	v505 = v487
	v509 = v497
	goto L95
L95:
	;
	v526 = v213 + v503*int32(24)
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526)+16)))
	if v527 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v327&int32(1) == int32(0) {
		v638 = v558
		goto L71
	} else {
		goto L104
	}
L97:
	;
	v533 = *(*int64)(unsafe.Add(mBase, uint32(v526)))
	*(*int64)(unsafe.Add(mBase, uint32(v335+v500<<(uint(int32(3))%32)))) = v533
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v536 = int32(1)
	v537 = v535 + v536
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v537
	v541 = v500 + v536
	v542 = v537
	goto L99
L98:
	;
	v541 = v500
	v542 = v505
	goto L99
L99:
	;
	v543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526)+40)))
	if v543 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v526)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v335+v541<<(uint(int32(3))%32)))) = v549
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v552 = int32(1)
	v553 = v551 + v552
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v553
	v557 = v541 + v552
	v558 = v553
	goto L102
L101:
	;
	v557 = v541
	v558 = v542
	goto L102
L102:
	;
	v559 = int32(2)
	v560 = v503 + v559
	v562 = v509 + v559
	if v562 != v327&int32(2147483646) {
		v500 = v557
		v503 = v560
		v505 = v558
		v509 = v562
		goto L95
	} else {
		goto L103
	}
L103:
	;
	goto L96
L104:
	;
	v567 = v557
	v570 = v560
	v572 = v558
	goto L91
L105:
	;
	v600 = *(*int64)(unsafe.Add(mBase, uint32(v593)))
	*(*int64)(unsafe.Add(mBase, uint32(v335+v567<<(uint(int32(3))%32)))) = v600
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	v630 = v602 + int32(1)
	goto L74
L106:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v664 = v662
	goto L50
L107:
	;
	m.G0 = v28 + int32(16)
	return v962
L108:
	;
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+82)))
	v922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+72)))
	v923 = F_datumCopy(m, v30, v921, v922)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L141
	}
L109:
	;
	v962 = v168 | base.I64_extend_i32_u(base.B2i32(v180 <= v179))
	goto L107
L110:
	;
	v799 = F_minmax_multi_get_strategy_procinfo(m, v31, v43&int32(_a_F_brin_minmax_multi_add_value_0), v689, int32(3))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L129
	}
L111:
	;
	v693 = v163 + int32(40)
	v699 = *(*int64)(unsafe.Add(mBase, uint32(v693+v664<<(uint(int32(4))%32)-int32(8))))
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v163)+40))
	v702 = v43 & int32(_a_F_brin_minmax_multi_add_value_0)
	v704 = F_minmax_multi_get_strategy_procinfo(m, v31, v702, v689, int32(1))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v706 = F_FunctionCall2Coll(m, v704, v49, v30, v700)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v706 != int64(0) {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v711 = F_minmax_multi_get_strategy_procinfo(m, v31, v702, v689, int32(5))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v713 = F_FunctionCall2Coll(m, v711, v49, v30, v699)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v713 != int64(0) {
		goto L110
	} else {
		goto L117
	}
L117:
	;
	v717 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v720 = v718 - int32(1)
	if v720 < v717 {
		goto L110
	} else {
		goto L118
	}
L118:
	;
	v724 = v720
	v725 = v720
	v727 = v717
	goto L119
L119:
	;
	v749 = base.I32_div_s(v725, int32(2))
	v752 = v693 + v749<<(uint(int32(4))%32)
	v753 = *(*int64)(unsafe.Add(mBase, uint32(v752)+8))
	v754 = *(*int64)(unsafe.Add(mBase, uint32(v752)))
	v755 = F_FunctionCall2Coll(m, v704, v49, v30, v754)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L110
L121:
	;
	if v768 <= v767 {
		v724 = v767
		v725 = v767 + v768
		v727 = v768
		goto L119
	} else {
		goto L128
	}
L122:
	;
	if v755 != int64(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v767 = v749 - int32(1)
	v768 = v727
	goto L121
L124:
	;
	goto L125
L125:
	;
	v761 = F_FunctionCall2Coll(m, v711, v49, v30, v753)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v761 == int64(0) {
		goto L109
	} else {
		goto L127
	}
L127:
	;
	v767 = v724
	v768 = v749 + int32(1)
	goto L121
L128:
	;
	goto L120
L129:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	if int32(16) <= v801 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v804
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v806
	v808 = int32(8)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v818 = F_bsearch_arg(m, v28+v808, v163+v810<<(uint(int32(4))%32)+int32(40), v801, v808, int32(22), v28)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v822 = v820 << (uint(int32(1)) % 32)
	if v822+v801 <= v822 {
		goto L108
	} else {
		goto L135
	}
L133:
	;
	if v818 != 0 {
		goto L109
	} else {
		goto L134
	}
L134:
	;
	goto L108
L135:
	;
	v828 = v822
	goto L136
L136:
	;
	v855 = *(*int64)(unsafe.Add(mBase, uint32(v163+int32(40)+v828<<(uint(int32(3))%32))))
	v856 = F_FunctionCall2Coll(m, v799, v49, v30, v855)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	goto L108
L138:
	;
	if v856 != int64(0) {
		goto L109
	} else {
		goto L139
	}
L139:
	;
	v860 = int32(1)
	v861 = v828 + v860
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v163)+20))
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	if v861 < v862+v863<<(uint(v860)%32) {
		v828 = v861
		goto L136
	} else {
		goto L140
	}
L140:
	;
	goto L137
L141:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v163+v925<<(uint(int32(4))%32)+v929<<(uint(int32(3))%32))+40)) = v923
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v163)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+24)) = v934 + int32(1)
	v938 = int64(1)
	if v934 != 0 {
		v962 = v938
		goto L107
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
	v962 = v938
	goto L107
}
func F_brin_minmax_multi_distance_interval(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(86400000000)
	v8 = base.I64_div_s(v6, v7)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	v15 = base.I64_rem_s(v13, v7)
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+8)))
	v22 = base.I64_div_s(v13, int64(-86400000000))
	v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+8)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v5)+12)))
	v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+12)))
	return base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i64_s(v6-v8*v7-v15), float64(8.64e+10)), base.F64_convert_i64_s(v20+(v22+v8)-v25+(v27-v28)*int64(30))))
}
func F_brin_minmax_multi_distance_timetz(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	return base.I64_reinterpret_f64(base.F64_convert_i64_s(v4 - v6 + base.I64_extend_i32_s(v8-v9)*int64(1000000)))
}
func F_brin_minmax_multi_summary_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v17 = v14 + int32(128)
	F_initStringInfo(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_appendStringInfoChar(m, v17, int32(123))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	F_getTypeOutputInfo(m, v28, v14+int32(120), v14+int32(127))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+120))
	F_fmgr_info(m, v35, v14+int32(92))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v41 = F_brin_range_deserialize(m, v40, v26)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v43
	F_appendStringInfo(m, v17, int32(_a_F_brin_minmax_multi_summary_out_0), v14+int32(48))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v54 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v145 <= int32(0) {
		goto L27
	} else {
		goto L28
	}
L10:
	;
	v134 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v60 = int32(0)
	v62 = v60
	v64 = v60
	v67 = int32(0)
	goto L13
L13:
	;
	v74 = v14 + int32(76)
	F_initStringInfo(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	if v110 <= int32(0) {
		v134 = v96
		goto L9
	} else {
		goto L22
	}
L15:
	;
	v78 = v14 + int32(92)
	v81 = v41 + int32(40) + v62<<(uint(int32(3))%32)
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	v83 = F_OutputFunctionCall(m, v78, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
	v86 = F_OutputFunctionCall(m, v78, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v83
	F_appendStringInfo(m, v74, int32(_a_F_brin_minmax_multi_summary_out_1), v14+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = v62 + int32(2)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v14)+80))
	v99 = F_cstring_to_text_with_len(m, v97, v98)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v106 = F_accumArrayResult(m, v64, base.I64_extend_i32_u(v99), int32(0), int32(25), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v109 = v67 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v109 < v110 {
		v62 = v96
		v64 = v106
		v67 = v109
		goto L13
	} else {
		goto L21
	}
L21:
	;
	goto L14
L22:
	;
	F_getTypeOutputInfo(m, int32(2277), v74, v14+int32(75))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v121 = F_makeArrayResult(m, v106, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v124 = F_OidOutputFunctionCall(m, v123, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v124
	F_appendStringInfo(m, v14+int32(128), int32(_a_F_brin_minmax_multi_summary_out_2), v14+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v134 = v96
	goto L9
L27:
	;
	F_appendStringInfoChar(m, v14+int32(128), int32(125))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L40
	}
L28:
	;
	v150 = int32(0)
	v152 = v134
	v154 = v150
	v157 = v150
	goto L29
L29:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v41+int32(40)+v152<<(uint(int32(3))%32))))
	v170 = F_FunctionCall1Coll(m, v14+int32(92), int32(0), v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	if v186 <= int32(0) {
		goto L27
	} else {
		goto L35
	}
L31:
	;
	v173 = F_cstring_to_text(m, base.I32_wrap_i64(v170))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v180 = F_accumArrayResult(m, v154, base.I64_extend_i32_u(v173), int32(0), int32(25), v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v182 = int32(1)
	v185 = v157 + v182
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v185 < v186 {
		v152 = v152 + v182
		v154 = v180
		v157 = v185
		goto L29
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	F_getTypeOutputInfo(m, int32(2277), v14+int32(76), v14+int32(75))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_brin_minmax_multi_summary_out[0]))
	v199 = F_makeArrayResult(m, v180, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
	v202 = F_OidOutputFunctionCall(m, v201, v199)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v202
	F_appendStringInfo(m, v14+int32(128), int32(_a_F_brin_minmax_multi_summary_out_3), v14)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L27
L40:
	;
	v226 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v14)+128)))
	m.G0 = v14 + int32(144)
	return v226
}
func F_brin_minmax_multi_summary_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_brin_minmax_multi_summary_recv_0), int32(3119), int32(_a_F_brin_minmax_multi_summary_recv_1), int32(_a_F_brin_minmax_multi_summary_recv_2), int32(_a_F_brin_minmax_multi_summary_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = F_palloc0(m, (v11*int32(24)+int32(31))&int32(-8)+v18<<(uint(int32(3))%32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v28 = F_palloc_mul(m, int32(8), v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v28
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
			v34 = F_palloc_mul(m, int32(1), v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v34
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
				v40 = F_palloc_mul(m, int32(1), v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v42)
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v40
					v46 = *(*int32)(unsafe.Add(mBase, _c_F_brin_new_memtuple[0]))
					v51 = F_AllocSetContextCreateInternal(m, v46, int32(_a_F_brin_new_memtuple_0), int32(0), int32(_a_F_brin_new_memtuple_1), int32(_a_F_brin_new_memtuple_2))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v51
						F_MemoryContextReset(m, v51)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
							if int32(0) < v57 {
								v62 = int32(24)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
								v76 = v22 + (v57*v62+int32(31))&int32(-8)
								v77 = int32(0)
								for {
									v83 = v22 + v62 + v77*int32(24)
									v85 = v77 + int32(1)
									*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v85)
									*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v83)+8)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v76
									v92 = int32(256)
									*(*uint16)(unsafe.Add(mBase, uint32(v83)+2)) = uint16(v92)
									*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = v71
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v77<<(uint(int32(2))%32))))
									v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98))))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
									if v85 < v104 {
										v76 = v76 + v99<<(uint(int32(3))%32)
										v77 = v85
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v115 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v115)
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v188 int64
	_ = v188
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	F_range_deduplicate_values(m, l0)
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = v24 + v25<<(uint(int32(1))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = F_get_typbyval(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = F_get_typlen(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v143 = F_palloc0(m, v130)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L35
	}
L5:
	;
	if v32 == int32(-1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	if v28 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if v32 == int32(-2) {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v130 = int32(20)
	goto L4
L10:
	;
	goto L11
L11:
	;
	v43 = v2
	v44 = int32(20)
	goto L12
L12:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)+v43<<(uint(int32(3))%32))))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	if v61 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v130 = v87
	goto L4
L14:
	;
	v87 = v86 + v44
	v89 = v43 + int32(1)
	if v89 != v28 {
		v43 = v89
		v44 = v87
		goto L12
	} else {
		goto L25
	}
L15:
	;
	v65 = int32(18)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v67 == v65 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v78 = int32(1)
	if v61&v78 != 0 {
		v86 = int32(base.Ui32(v61) >> (uint(v78) % 32))
		goto L14
	} else {
		goto L24
	}
L18:
	;
	v70 = v65
	goto L20
L19:
	;
	v70 = int32(2)
	goto L20
L20:
	;
	if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v77 = int32(6)
	goto L23
L22:
	;
	v77 = v70
	goto L23
L23:
	;
	v86 = v77
	goto L14
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v86 = int32(base.Ui32(v82) >> (uint(int32(2)) % 32))
	goto L14
L25:
	;
	goto L13
L26:
	;
	if v28 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v130 = v32*v28 + int32(20)
	goto L4
L29:
	;
	v130 = int32(20)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v100 = v2
	v101 = int32(20)
	goto L32
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(40)+v100<<(uint(int32(3))%32))))
	v118 = F_strlen(m, v117)
	mBase = m.M
	v120 = int32(1)
	v121 = v118 + v101 + v120
	v123 = v100 + v120
	if v123 != v28 {
		v100 = v123
		v101 = v121
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v130 = v121
	goto L4
L34:
	;
	goto L33
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v130 << (uint(int32(2)) % 32)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v149
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = v153
	if int32(0) < v28 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v158 = l0 + int32(40)
	v171 = v143 + int32(20)
	v172 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	m.G0 = v18 + int32(16)
	return v143
L39:
	;
	if v30 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	v267 = v172 + int32(1)
	if v267 != v28 {
		v171 = v262
		v172 = v267
		goto L39
	} else {
		goto L85
	}
L42:
	;
	if base.I32_popcnt(v32) != int32(1) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	goto L44
L44:
	;
	if int32(0) < v32 {
		goto L58
	} else {
		goto L59
	}
L45:
	;
	if v32 != 0 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v188)
	goto L45
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L52
	}
L48:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v158+v172<<(uint(int32(3))%32))))
	switch base.I32_ctz(v32) {
	case 0:
		goto L46
	case 1:
		goto L51
	case 2:
		goto L50
	case 3:
		goto L49
	default:
		goto L47
	}
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v188
	goto L45
L50:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v18)+8)) = uint32(v188)
	goto L45
L51:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)) = uint16(v188)
	goto L45
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v32
	F_errmsg_internal(m, int32(_a_F_brin_range_serialize_0), v18)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_brin_range_serialize_1), int32(474), int32(_a_F_brin_range_serialize_2))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	base.MemoryCopy(m, v171, v18+int32(8), v32)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v262 = v171 + v32
	goto L41
L58:
	;
	if v32 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	if base.B2i32(v32 != int32(-1)) == int32(0) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v158+v172<<(uint(int32(3))%32))))
	base.MemoryCopy(m, v171, v216, v32)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v262 = v171 + v32
	goto L41
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v158+v172<<(uint(int32(3))%32))))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
	if v225 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	goto L66
L66:
	;
	if v32 != int32(-2) {
		v262 = v171
		goto L41
	} else {
		goto L81
	}
L67:
	;
	if v250 != 0 {
		goto L78
	} else {
		goto L79
	}
L68:
	;
	v229 = int32(18)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	if v231 == v229 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v242 = int32(1)
	if v225&v242 != 0 {
		v250 = int32(base.Ui32(v225) >> (uint(v242) % 32))
		goto L67
	} else {
		goto L77
	}
L71:
	;
	v234 = v229
	goto L73
L72:
	;
	v234 = int32(2)
	goto L73
L73:
	;
	if base.Ui32((v231-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v241 = int32(6)
	goto L76
L75:
	;
	v241 = v234
	goto L76
L76:
	;
	v250 = v241
	goto L67
L77:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v250 = int32(base.Ui32(v246) >> (uint(int32(2)) % 32))
	goto L67
L78:
	;
	base.MemoryCopy(m, v171, v224, v250)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v262 = v171 + v250
	goto L41
L81:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v158+v172<<(uint(int32(3))%32))))
	v257 = F_strlen(m, v256)
	mBase = m.M
	v259 = v257 + int32(1)
	if v259 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	base.MemoryCopy(m, v171, v256, v259)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v262 = v171 + v259
	goto L41
L85:
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
	var v60 int32
	_ = v60
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
	var v85 int32
	_ = v85
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
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v20^int32(-1))*int32(56))+16))
				v66 = v57
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[3]))
				v60 = int32(56)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+v20*v60-v60)+16))
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
				F_errstart_cold(m, int32(24), int32(0))
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
						F_errstart_cold(m, int32(24), int32(0))
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
						*(*int64)(unsafe.Add(mBase, uint32(v144))) = base.I64_rotl(v14, int64(32))
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
											*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
										*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
				v82 = *(*int32)(unsafe.Add(mBase, uint32(v76+(v72^int32(-1))*int32(56))+16))
				v91 = v82
			} else {
				v84 = *(*int32)(unsafe.Add(mBase, _c_F_brin_xlog_insert_update[3]))
				v85 = int32(56)
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v84+v72*v85-v85)+16))
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
								*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
							*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
					F_errstart_cold(m, int32(24), int32(0))
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
							F_errstart_cold(m, int32(24), int32(0))
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
							*(*int64)(unsafe.Add(mBase, uint32(v144))) = base.I64_rotl(v14, int64(32))
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
												*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
											*(*int64)(unsafe.Add(mBase, uint32(v209))) = base.I64_rotl(v14, int64(32))
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
