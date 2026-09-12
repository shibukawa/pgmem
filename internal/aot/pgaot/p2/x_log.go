package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetXLogBuffer(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v45 int64
	_ = v45
	var v48 int64
	_ = v48
	var v57 int32
	_ = v57
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v1 = l0
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = int64(base.Ui64(v1) >> (uint(int64(13)) % 64))
	v18 = *(*int64)(unsafe.Add(mBase, _consts[173]))
	if v16 == v18 {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[174]))
		v134 = v21 + base.I32_wrap_i64(v1)&int32(8191)
		m.G0 = v13 + int32(16)
		return v134
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, _consts[175]))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+300))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+304))
		v33 = base.I64_rem_u_s(v16, base.I64_extend_i32_s(v29+int32(1)))
		v34 = base.I32_wrap_i64(v33)
		v37 = v28 + v34<<(uint(int32(3))%32)
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
		*(*int64)(unsafe.Add(mBase, uint32(v37))) = v38
		v41 = v1 & int64(8191)
		v42 = int64(-8192)
		v45 = v1&v42 - v42
		if v45 != v38 {
			v48 = v41 - int64(24)
			if base.Ui64(int64(16)) < base.Ui64(v48) {
				v76 = v1
			} else {
				switch base.I32_wrap_i64(v48) - int32(1) {
				case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
					v76 = v1
				case 15:
					v68 = *(*int32)(unsafe.Add(mBase, _consts[176]))
					if base.Ui64(v1&base.I64_extend_i32_s(v68-int32(1))) < base.Ui64(int64(8192)) {
						v75 = v1 - int64(40)
					} else {
						v75 = v1
					}
					v76 = v75
				default:
					v57 = *(*int32)(unsafe.Add(mBase, _consts[176]))
					if base.Ui64(int64(8192)) < base.Ui64(v1&base.I64_extend_i32_s(v57-int32(1))) {
						v64 = v1 - int64(24)
					} else {
						v64 = v1
					}
					v76 = v64
				}
			}
			v78 = *(*int32)(unsafe.Add(mBase, _consts[177]))
			v80 = int32(*(*uint8)(unsafe.Add(mBase, _consts[178])))
			if v80 != 0 {
				F_LWLockUpdateVar(m, v78+int32(896), v78+int32(912), v76)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_AdvanceXLInsertBuffer(m, v1, l1, int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _consts[175]))
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+300))
						v107 = v104 + v34<<(uint(int32(3))%32)
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
						*(*int64)(unsafe.Add(mBase, uint32(v107))) = v108
						if v108 != v45 {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v1)
								v145 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v145)
								F_errmsg_internal(m, int32(504899), v13)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(489152), int32(1712), int32(222884))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, _consts[173])) = v16
							v119 = *(*int32)(unsafe.Add(mBase, _consts[175]))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+296))
							v123 = v120 + v34<<(uint(int32(13))%32)
							*(*int32)(unsafe.Add(mBase, _consts[174])) = v123
							v134 = v123 + base.I32_wrap_i64(v41)
							m.G0 = v13 + int32(16)
							return v134
						}
					}
				}
			} else {
				v90 = *(*int32)(unsafe.Add(mBase, _consts[179]))
				v93 = v78 + v90<<(uint(int32(7))%32)
				F_LWLockUpdateVar(m, v93, v93+int32(16), v76)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_AdvanceXLInsertBuffer(m, v1, l1, int32(0))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						v103 = *(*int32)(unsafe.Add(mBase, _consts[175]))
						v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+300))
						v107 = v104 + v34<<(uint(int32(3))%32)
						v108 = *(*int64)(unsafe.Add(mBase, uint32(v107)))
						*(*int64)(unsafe.Add(mBase, uint32(v107))) = v108
						if v108 != v45 {
							F_errstart_cold(m, int32(23), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								*(*uint32)(unsafe.Add(mBase, uint32(v13)+4)) = uint32(v1)
								v145 = int64(base.Ui64(v1) >> (uint(int64(32)) % 64))
								*(*uint32)(unsafe.Add(mBase, uint32(v13))) = uint32(v145)
								F_errmsg_internal(m, int32(504899), v13)
								mBase = m.M
								v149 = m.ExcPending
								if v149 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(489152), int32(1712), int32(222884))
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, _consts[173])) = v16
							v119 = *(*int32)(unsafe.Add(mBase, _consts[175]))
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+296))
							v123 = v120 + v34<<(uint(int32(13))%32)
							*(*int32)(unsafe.Add(mBase, _consts[174])) = v123
							v134 = v123 + base.I32_wrap_i64(v41)
							m.G0 = v13 + int32(16)
							return v134
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, _consts[173])) = v16
			v119 = *(*int32)(unsafe.Add(mBase, _consts[175]))
			v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+296))
			v123 = v120 + v34<<(uint(int32(13))%32)
			*(*int32)(unsafe.Add(mBase, _consts[174])) = v123
			v134 = v123 + base.I32_wrap_i64(v41)
			m.G0 = v13 + int32(16)
			return v134
		}
	}
}
func F_XLogArchiveCleanup(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	v3 = m.G0
	v5 = v3 - int32(1056)
	m.G0 = v5
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(366132)
	v16 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(172701), v5+int32(16))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = v5 + int32(32)
		v20 = F_unlink(m, v19)
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(22602)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		v28 = F_pg_snprintf(m, v19, int32(1024), int32(172701), v5)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v32 = F_unlink(m, v5+int32(32))
			mBase = m.M
			m.G0 = v5 + int32(1056)
			return
		}
	}
}
func F_XLogFlush(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v1 int64
	_ = v1
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v321 int64
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int64
	_ = v421
	var v422 int32
	_ = v422
	var v434 int64
	_ = v434
	var v444 int32
	_ = v444
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v513 int64
	_ = v513
	var v518 int32
	_ = v518
	var v520 int64
	_ = v520
	var v523 int64
	_ = v523
	var v524 int64
	_ = v524
	var v527 int64
	_ = v527
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	v1 = l0
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+308))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[180]))
	if v22 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v16 + int32(48)
	return
L2:
	;
	v44 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	if base.Ui64(v1) <= base.Ui64(v44) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	F_UpdateMinRecoveryPoint(m, v1, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	if v22 != 0 {
		goto L2
	} else {
		goto L11
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+316))
	v32 = base.B2i32(v30 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v32)
	if v30 != int32(2) {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[180])) = int32(1)
	goto L2
L10:
	;
	goto L9
L11:
	;
	goto L3
L12:
	;
	return
L13:
	;
	goto L1
L14:
	;
	v46 = int32(4465060)
	v48 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v48 + int32(1)
	v64 = v1
	goto L16
L15:
	;
	v477 = int32(4465060)
	v479 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v480 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v479 - v480
	v485 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v485 == v480 {
		goto L99
	} else {
		goto L100
	}
L16:
	;
	v65 = int32(4365968)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+280)) = v67
	v69 = int32(4366008)
	*(*int64)(unsafe.Add(mBase, _consts[181])) = v67
	v72 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v72)+272)) = v73
	*(*int64)(unsafe.Add(mBase, _consts[183])) = v73
	v78 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	if base.Ui64(v1) <= base.Ui64(v78) {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v308 = int32(4365968)
	v309 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v309)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v309)+280)) = v310
	v312 = int32(4366008)
	*(*int64)(unsafe.Add(mBase, _consts[181])) = v310
	v315 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+272)) = v316
	*(*int64)(unsafe.Add(mBase, _consts[183])) = v316
	v321 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	if base.Ui64(v321) < base.Ui64(v1) {
		goto L74
	} else {
		goto L75
	}
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+440))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+440)) = int32(1)
	if v82 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	F_s_lock(m, v86+int32(440), int32(489152), int32(2843), int32(316686))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+440)) = int32(0)
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v95)+184))
	if base.Ui64(v98) < base.Ui64(v64) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v100 = v64
	goto L25
L24:
	;
	v100 = v98
	goto L25
L25:
	;
	v101 = F_WaitXLogInsertionsToFinish(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v106 = v104 + int32(1024)
	v108 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	if v108 < int32(200) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v305 == int32(0) {
		v64 = v100
		goto L16
	} else {
		goto L73
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v113 = int32(4465052)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v115 + int32(1)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v121 = v119
	goto L31
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L12
	} else {
		goto L70
	}
L31:
	;
	v136 = v121 & int32(524287)
	if v136 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(v136 == int32(0)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v137 = v121
	goto L35
L34:
	;
	v137 = v121 | int32(262144)
	goto L35
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v141 = base.B2i32(v121 == v140)
	if v121 == v140 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v142 = v137
	goto L38
L37:
	;
	v142 = v140
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v142
	if v141 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v121 = v140
	goto L31
L40:
	;
	goto L41
L41:
	;
	goto L32
L42:
	;
	F_LWLockQueueSelf(m, v106, int32(2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v263 = int32(4392060)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[184]))
	v266 = v264 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_consts[187]))) = v106
	v271 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[184])) = v264 + v271
	*(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_consts[188]))) = int32(0)
	v305 = v271
	goto L27
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v153 = v151
	goto L46
L46:
	;
	v168 = v153 & int32(524287)
	if v168 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if base.B2i32(v168 == int32(0)) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L48:
	;
	v169 = v153
	goto L50
L49:
	;
	v169 = v153 | int32(262144)
	goto L50
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v173 = base.B2i32(v153 == v172)
	if v153 == v172 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v174 = v169
	goto L53
L52:
	;
	v174 = v172
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+4)) = v174
	if v173 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v153 = v172
	goto L46
L55:
	;
	goto L56
L56:
	;
	goto L47
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106))))
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v183 | int32(16777216)
	v189 = int32(0)
	goto L60
L58:
	;
	goto L59
L59:
	;
	F_LWLockDequeueSelf(m, v106)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L69
	}
L60:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+74)))
	if v203 != 0 {
		v189 = v189 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[189]))
	v206 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v206
	if v206 < v189 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	v211 = v189
	goto L66
L64:
	;
	goto L65
L65:
	;
	v241 = int32(4465052)
	v243 = *(*int32)(unsafe.Add(mBase, _consts[186]))
	*(*int32)(unsafe.Add(mBase, _consts[186])) = v243 - int32(1)
	v305 = int32(0)
	goto L27
L66:
	;
	v224 = int32(1)
	if base.Ui32(v224) < base.Ui32(v211) {
		v211 = v211 - v224
		goto L66
	} else {
		goto L68
	}
L67:
	;
	goto L65
L68:
	;
	goto L67
L69:
	;
	goto L44
L70:
	;
	F_errmsg_internal(m, int32(277596), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(488713), int32(1425), int32(101772))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L12
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	goto L17
L74:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v324 <= int32(0) {
		v434 = v101
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v459+int32(1024))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L12
	} else {
		goto L98
	}
L77:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v434
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v434
	F_XLogWrite(m, v16+int32(16), v20, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L12
	} else {
		goto L97
	}
L78:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, _consts[191])))
	if v328 != int32(1) {
		v434 = v101
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v331 = int32(0)
	v334 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	if v334 == v331 {
		v414 = int32(1)
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v414 == int32(0) {
		v434 = v101
		goto L77
	} else {
		goto L94
	}
L81:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _consts[193]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	if v339 <= int32(0) {
		v390 = v331
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v414 = base.B2i32(v334 <= v390)
	goto L80
L83:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _consts[185]))
	v347 = *(*int32)(unsafe.Add(mBase, _consts[194]))
	v352 = v331
	v357 = int32(0)
	goto L84
L84:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v338+int32(36)+v357<<(uint(int32(2))%32))))
	if v365 == int32(-1) {
		v382 = v352
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v390 = v382
	goto L82
L86:
	;
	v385 = v357 + int32(1)
	if v385 != v339 {
		v352 = v382
		v357 = v385
		goto L84
	} else {
		goto L93
	}
L87:
	;
	v370 = v347 + v365*int32(640)
	if v370 == v345 {
		v382 = v352
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v370)+36))
	if v372 == int32(0) {
		v382 = v352
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v370)+44))
	if v375 == int32(0) {
		v382 = v352
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v370)+92))
	if v378 != 0 {
		v382 = v352
		goto L86
	} else {
		goto L91
	}
L91:
	;
	v380 = v352 + int32(1)
	if v334 <= v380 {
		v390 = v380
		goto L82
	} else {
		goto L92
	}
L92:
	;
	v382 = v380
	goto L86
L93:
	;
	goto L85
L94:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	F_pg_usleep(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	v421 = F_WaitXLogInsertionsToFinish(m, v101)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	v434 = v421
	goto L77
L97:
	;
	goto L76
L98:
	;
	goto L15
L99:
	;
	v490 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+316))
	v492 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(base.B2i32(v491 != v492))
	v497 = base.B2i32(v491 == v492)
	goto L101
L100:
	;
	v497 = v480
	goto L101
L101:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, _consts[195])))
	if v499 != int32(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v513 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	if base.Ui64(v1) <= base.Ui64(v513) {
		goto L1
	} else {
		goto L106
	}
L103:
	;
	v503 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[195])) = uint8(v503)
	v506 = *(*int32)(unsafe.Add(mBase, _consts[196]))
	if v506 <= v503 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	F_WalSndWakeup(m, int32(1), v497)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	v520 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+12)) = uint32(v520)
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+4)) = uint32(v1)
	v523 = int64(32)
	v524 = int64(base.Ui64(v1) >> (uint(v523) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16))) = uint32(v524)
	v527 = int64(base.Ui64(v520) >> (uint(v523) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v16)+8)) = uint32(v527)
	F_errmsg_internal(m, int32(505137), v16)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(489152), int32(2942), int32(316686))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_XLogPrefetcherComputeStats(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+124))
	if v7 != 0 {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)+16))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+120))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)+16))
		v13 = base.I32_wrap_i64(v8 - v10)
	} else {
		v13 = int32(0)
	}
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v18 + v15
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v24 - int64(-8192)
	return
}
func F_XLogReadDetermineTimeline(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int64
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
	var v73 int64
	_ = v73
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+1192)))
	v14 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	v17 = v13 + v14*v15
	if v17 != l1 {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
		if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
			m.G0 = v11 + int32(16)
			return
		} else {
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
			if v36 == int64(0) {
				v47 = F_readTimeLineHistory(m, l3)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
					v50 = base.I64_div_u_s(l1, v49)
					v51 = int64(1)
					v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
						v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
							F_list_free_deep(m, v47)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								v68 = F_errstart(m, int32(12), int32(0))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									if v68 == int32(0) {
										m.G0 = v11 + int32(16)
										return
									} else {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
										v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
										*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
										v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
										*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
										F_errmsg_internal(m, int32(505987), v11)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											F_errfinish(m, int32(485062), int32(800), int32(366877))
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												m.G0 = v11 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if l3 == v32 {
					v47 = F_readTimeLineHistory(m, l3)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
						v50 = base.I64_div_u_s(l1, v49)
						v51 = int64(1)
						v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
							v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
								F_list_free_deep(m, v47)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v68 = F_errstart(m, int32(12), int32(0))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										if v68 == int32(0) {
											m.G0 = v11 + int32(16)
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
											v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
											v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
											F_errmsg_internal(m, int32(505987), v11)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_errfinish(m, int32(485062), int32(800), int32(366877))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if v32 == int32(0) {
						v47 = F_readTimeLineHistory(m, l3)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
							v50 = base.I64_div_u_s(l1, v49)
							v51 = int64(1)
							v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
								v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
									F_list_free_deep(m, v47)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v68 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if v68 == int32(0) {
												m.G0 = v11 + int32(16)
												return
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
												v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
												v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
												F_errmsg_internal(m, int32(505987), v11)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													F_errfinish(m, int32(485062), int32(800), int32(366877))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v44 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
						v45 = base.I64_div_u_s(v36, v14)
						if base.Ui64(v44) < base.Ui64(v45) {
							m.G0 = v11 + int32(16)
							return
						} else {
							v47 = F_readTimeLineHistory(m, l3)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
								v50 = base.I64_div_u_s(l1, v49)
								v51 = int64(1)
								v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
									v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
										F_list_free_deep(m, v47)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v68 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												if v68 == int32(0) {
													m.G0 = v11 + int32(16)
													return
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
													v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
													v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
													F_errmsg_internal(m, int32(505987), v11)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_errfinish(m, int32(485062), int32(800), int32(366877))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16)
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
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v19 == int32(0) {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
			if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
				m.G0 = v11 + int32(16)
				return
			} else {
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
				if v36 == int64(0) {
					v47 = F_readTimeLineHistory(m, l3)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
						v50 = base.I64_div_u_s(l1, v49)
						v51 = int64(1)
						v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
							v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
								F_list_free_deep(m, v47)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									v68 = F_errstart(m, int32(12), int32(0))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										if v68 == int32(0) {
											m.G0 = v11 + int32(16)
											return
										} else {
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
											v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
											*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
											v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
											*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
											F_errmsg_internal(m, int32(505987), v11)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return
											} else {
												F_errfinish(m, int32(485062), int32(800), int32(366877))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l3 == v32 {
						v47 = F_readTimeLineHistory(m, l3)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
							v50 = base.I64_div_u_s(l1, v49)
							v51 = int64(1)
							v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
								v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
									F_list_free_deep(m, v47)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v68 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if v68 == int32(0) {
												m.G0 = v11 + int32(16)
												return
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
												v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
												v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
												F_errmsg_internal(m, int32(505987), v11)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													F_errfinish(m, int32(485062), int32(800), int32(366877))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v32 == int32(0) {
							v47 = F_readTimeLineHistory(m, l3)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
								v50 = base.I64_div_u_s(l1, v49)
								v51 = int64(1)
								v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
									v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
										F_list_free_deep(m, v47)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v68 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												if v68 == int32(0) {
													m.G0 = v11 + int32(16)
													return
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
													v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
													v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
													F_errmsg_internal(m, int32(505987), v11)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_errfinish(m, int32(485062), int32(800), int32(366877))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v44 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
							v45 = base.I64_div_u_s(v36, v14)
							if base.Ui64(v44) < base.Ui64(v45) {
								m.G0 = v11 + int32(16)
								return
							} else {
								v47 = F_readTimeLineHistory(m, l3)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
									v50 = base.I64_div_u_s(l1, v49)
									v51 = int64(1)
									v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
										v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
											F_list_free_deep(m, v47)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return
											} else {
												v68 = F_errstart(m, int32(12), int32(0))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													if v68 == int32(0) {
														m.G0 = v11 + int32(16)
														return
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
														v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
														*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
														v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
														F_errmsg_internal(m, int32(505987), v11)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return
														} else {
															F_errfinish(m, int32(485062), int32(800), int32(366877))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16)
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
			}
		} else {
			v24 = int32(8191)
			if base.Ui32(v24) <= base.Ui32(l2) {
				v27 = v24
			} else {
				v27 = l2
			}
			if base.Ui64(l1+base.I64_extend_i32_u(v27)) <= base.Ui64(l1+base.I64_extend_i32_u(v19)) {
				m.G0 = v11 + int32(16)
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
				if base.B2i32(v32 == l3)&base.B2i32(base.Ui64(v17) <= base.Ui64(l1)) != 0 {
					m.G0 = v11 + int32(16)
					return
				} else {
					v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
					if v36 == int64(0) {
						v47 = F_readTimeLineHistory(m, l3)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
							v50 = base.I64_div_u_s(l1, v49)
							v51 = int64(1)
							v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
								v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
									F_list_free_deep(m, v47)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v68 = F_errstart(m, int32(12), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											if v68 == int32(0) {
												m.G0 = v11 + int32(16)
												return
											} else {
												v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
												v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
												v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
												*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
												F_errmsg_internal(m, int32(505987), v11)
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return
												} else {
													F_errfinish(m, int32(485062), int32(800), int32(366877))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						if l3 == v32 {
							v47 = F_readTimeLineHistory(m, l3)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
								v50 = base.I64_div_u_s(l1, v49)
								v51 = int64(1)
								v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
									v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
										F_list_free_deep(m, v47)
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											v68 = F_errstart(m, int32(12), int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												if v68 == int32(0) {
													m.G0 = v11 + int32(16)
													return
												} else {
													v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
													v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
													v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
													*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
													F_errmsg_internal(m, int32(505987), v11)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														F_errfinish(m, int32(485062), int32(800), int32(366877))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															m.G0 = v11 + int32(16)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v32 == int32(0) {
								v47 = F_readTimeLineHistory(m, l3)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return
								} else {
									v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
									v50 = base.I64_div_u_s(l1, v49)
									v51 = int64(1)
									v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
										v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
											F_list_free_deep(m, v47)
											mBase = m.M
											v65 = m.ExcPending
											if v65 != 0 {
												return
											} else {
												v68 = F_errstart(m, int32(12), int32(0))
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													if v68 == int32(0) {
														m.G0 = v11 + int32(16)
														return
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
														v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
														*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
														*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
														v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
														*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
														F_errmsg_internal(m, int32(505987), v11)
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return
														} else {
															F_errfinish(m, int32(485062), int32(800), int32(366877))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return
															} else {
																m.G0 = v11 + int32(16)
																return
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v44 = base.I64_div_u_s(l1+base.I64_extend_i32_u(l2), v14)
								v45 = base.I64_div_u_s(v36, v14)
								if base.Ui64(v44) < base.Ui64(v45) {
									m.G0 = v11 + int32(16)
									return
								} else {
									v47 = F_readTimeLineHistory(m, l3)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										v49 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
										v50 = base.I64_div_u_s(l1, v49)
										v51 = int64(1)
										v56 = F_tliOfPointInHistory(m, (v50+v51)*v49-v51, v47)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1224)) = v56
											v61 = F_tliSwitchPoint(m, v56, v47, l0+int32(1240))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(l0)+1232)) = v61
												F_list_free_deep(m, v47)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													v68 = F_errstart(m, int32(12), int32(0))
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														if v68 == int32(0) {
															m.G0 = v11 + int32(16)
															return
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
															v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+8)) = uint32(v73)
															*(*int32)(unsafe.Add(mBase, uint32(v11))) = v72
															v77 = int64(base.Ui64(v73) >> (uint(int64(32)) % 64))
															*(*uint32)(unsafe.Add(mBase, uint32(v11)+4)) = uint32(v77)
															F_errmsg_internal(m, int32(505987), v11)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return
															} else {
																F_errfinish(m, int32(485062), int32(800), int32(366877))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
																	return
																} else {
																	m.G0 = v11 + int32(16)
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
				}
			}
		}
	}
}
func F_XLogRecordPageWithFreeSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v13
	v17 = base.I32_div_u_s(l1, int32(4069))
	v19 = base.I32_div_u_s(l1, int32(16556761))
	v25 = F_XLogReadBufferExtended(m, v9, int32(1), v17+v19+int32(2), int32(3), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_LockBuffer(m, v25, int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(8159)) < base.Ui32(l2) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = int32(-1)
	goto L6
L5:
	;
	v35 = int32(base.Ui32(l2) >> (uint(int32(5)) % 32))
	goto L6
L6:
	;
	if v25 < int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+14)))
	if v57 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42+(v25^int32(-1))<<(uint(int32(2))%32))))
	v56 = v48
	goto L7
L9:
	;
	goto L10
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v56 = v50 + v25<<(uint(int32(13))%32) + int32(-8192)
	goto L7
L11:
	;
	if v56&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L13
L13:
	;
	v102 = v35 & int32(255)
	v107 = v56 + int32(28)
	v109 = l1 - v17*int32(4069) + int32(4095)
	v110 = v107 + v109
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v111 != v102 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+10)) = int32(1572864)
	v92 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+18)) = uint16(v92)
	v98 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+16)) = uint16(v98)
	*(*uint16)(unsafe.Add(mBase, uint32(v56)+14)) = uint16(v98)
	goto L14
L16:
	;
	v86 = F___memset(m, v56, int32(0), int32(8192))
	mBase = m.M
	goto L15
L17:
	;
	goto L16
L24:
	;
	if v201 != 0 {
		goto L55
	} else {
		goto L56
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v102)
	v121 = v109
	goto L28
L26:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if base.Ui32(v113) < base.Ui32(v102) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v201 = int32(0)
	goto L24
L28:
	;
	v123 = int32(1)
	v124 = v121 - v123
	v125 = int32(2)
	v126 = base.I32_div_s(v124, v125)
	v128 = v126 << (uint(v123) % 32)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v128)+1)))
	v132 = v128 + v125
	if base.Ui32(v132) <= base.Ui32(int32(8163)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if base.Ui32(v151) < base.Ui32(v102) {
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v136 = v130 & int32(255)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v132))))
	if base.Ui32(v138) < base.Ui32(v136) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v141 = v130
	goto L32
L32:
	;
	v143 = v107 + v126
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v144 != v141&int32(255) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v140 = v136
	goto L35
L34:
	;
	v140 = v138
	goto L35
L35:
	;
	v141 = v140
	goto L32
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v141)
	if int32(1) < v124 {
		v121 = v126
		goto L28
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L29
L39:
	;
	goto L38
L40:
	;
	v157 = int32(4094)
	goto L43
L41:
	;
	goto L42
L42:
	;
	v201 = int32(1)
	goto L24
L43:
	;
	if base.Ui32(int32(4081)) < base.Ui32(v157) {
		v179 = int32(0)
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L42
L45:
	;
	v180 = v107 + v157
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != v179&int32(255) {
		goto L51
	} else {
		goto L52
	}
L46:
	;
	v164 = v157 << (uint(int32(1)) % 32)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+v164)+1)))
	if v157 == int32(4081) {
		v179 = v166
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v170 = v166 & int32(255)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107+(v164+int32(2))))))
	if base.Ui32(v174) < base.Ui32(v170) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v176 = v170
	goto L50
L49:
	;
	v176 = v174
	goto L50
L50:
	;
	v179 = v176
	goto L45
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v179)
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v157 != 0 {
		v157 = v157 - int32(1)
		goto L43
	} else {
		goto L54
	}
L54:
	;
	goto L44
L55:
	;
	F_MarkBufferDirtyHint(m, v25, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_UnlockReleaseBuffer(m, v25)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	m.G0 = v9 + int32(16)
	return
}
func F_XLogRegisterBufData(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _consts[233]))
	v15 = v12 + l0*int32(8260)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v16 != 0 {
		v18 = *(*int32)(unsafe.Add(mBase, _consts[234]))
		v20 = *(*int32)(unsafe.Add(mBase, _consts[235]))
		if v20 <= v18 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(496058), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, _consts[234]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v71
					v74 = *(*int32)(unsafe.Add(mBase, _consts[235]))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v74
					F_errdetail_internal(m, int32(606554), v9+int32(16))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return
					} else {
						F_errfinish(m, int32(484174), int32(428), int32(496247))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
			if base.Ui32(int32(65536)) <= base.Ui32(v22+l2|l2) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(496058), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v94
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(65535)
						F_errdetail_internal(m, int32(571800), v9)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return
						} else {
							F_errfinish(m, int32(484174), int32(433), int32(496247))
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[234])) = v18 + int32(1)
				v32 = *(*int32)(unsafe.Add(mBase, _consts[236]))
				v35 = v32 + v18*int32(12)
				*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l1
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v35
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v41 + l2
				m.G0 = v9 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
			F_errmsg_internal(m, int32(242350), v9+int32(32))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				F_errfinish(m, int32(484174), int32(416), int32(496247))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
