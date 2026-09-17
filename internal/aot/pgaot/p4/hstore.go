package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreArrayToPairs(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	F_deconstruct_array_builtin(m, l0, int32(25), v12+int32(28), v12+int32(24), v12+int32(20))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
		if v25 != 0 {
			if base.Ui32(int32(53687092)) <= base.Ui32(v25) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = int32(53687091)
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v121
						F_errmsg(m, int32(_a_F_hstoreArrayToPairs_0), v12)
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hstoreArrayToPairs_1), int32(102), int32(_a_F_hstoreArrayToPairs_2))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
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
				v30 = F_palloc(m, v25*int32(20))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if v32 <= int32(0) {
						v84 = int32(0)
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
						v38 = int32(0)
						v40 = v38
						v41 = v38
						for {
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40+v37))))
							if v50 == int32(0) {
								v55 = v30 + v41*int32(20)
								v56 = int32(2)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v36+v40<<(uint(v56)%32))))
								v60 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v55))) = v59 + v60
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								v64 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v64
								v68 = int32(1)
								*(*uint16)(unsafe.Add(mBase, uint32(v55)+16)) = uint16(v68)
								*(*int32)(unsafe.Add(mBase, uint32(v55)+8)) = int32(base.Ui32(v63)>>(uint(v56)%32)) - v60
								v77 = v41 + v68
							} else {
								v77 = v41
							}
							v81 = v40 + int32(1)
							if v81 != v32 {
								v40 = v81
								v41 = v77
								continue
							} else {
								break
							}
							break
						}
						v84 = v77
					}
					v94 = F_hstoreUniquePairs(m, v30, v84, v12+int32(16))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v101 = v30
						v106 = v94
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v106
						m.G0 = v12 + int32(32)
						return v101
					}
				}
			}
		} else {
			v101 = int32(0)
			v106 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v106
			m.G0 = v12 + int32(32)
			return v101
		}
	}
}
func F_hstoreValidNewFormat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	v2 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = v13 & int32(268435455)
	if base.B2i32(v15 == v2)|base.B2i32(v13 < v2) != 0 {
		v147 = int32(2)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v147
L2:
	;
	v22 = l0 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if int32(0) <= v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v28 = int32(0)
	v30 = v15 << (uint(int32(3)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v22-int32(4))))
	v39 = v30 + v34&int32(1073741823) + int32(8)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(base.Ui32(v40) >> (uint(int32(2)) % 32))
	if base.Ui32(v42) < base.Ui32(v39) {
		v147 = v28
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v44 = int32(1)
	v48 = v44
	goto L7
L7:
	;
	v60 = v22 + v48<<(uint(int32(2))%32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v61 < int32(0) {
		v147 = v28
		goto L1
	} else {
		goto L9
	}
L8:
	;
	if v15 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v64 = int32(1073741823)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60-int32(4))))
	if base.Ui32(v61&v64) < base.Ui32(v68&v64) {
		v147 = v28
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v73 = v48 + int32(1)
	if v73 != v15<<(uint(v44)%32) {
		v48 = v73
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v77 = int32(2)
	if base.Ui32(v15) <= base.Ui32(v77) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v39 == v42 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v80 = v77
	goto L17
L16:
	;
	v80 = v15
	goto L17
L17:
	;
	v84 = int32(1)
	goto L18
L18:
	;
	v94 = v84 << (uint(int32(3)) % 32)
	v95 = v22 + v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v98 = v96 & int32(1073741823)
	if int32(0) <= v96 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L14
L20:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95-int32(4))))
	v107 = v98 - v103&int32(1073741823)
	goto L22
L21:
	;
	v107 = v98
	goto L22
L22:
	;
	v108 = l0 + v94
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v111 = v109 & int32(1073741823)
	if int32(0) <= v109 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108-int32(4))))
	v120 = v111 - v116&int32(1073741823)
	goto L25
L24:
	;
	v120 = v111
	goto L25
L25:
	;
	if v96&int32(1073741824)|base.B2i32(base.Ui32(v107) < base.Ui32(v120)) != 0 {
		v147 = int32(0)
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v127 = v84 + int32(1)
	if v127 != v80 {
		v84 = v127
		goto L18
	} else {
		goto L27
	}
L27:
	;
	goto L19
L28:
	;
	v143 = int32(2)
	goto L30
L29:
	;
	v143 = int32(1)
	goto L30
L30:
	;
	v147 = v143
	goto L1
}
func F_hstore_concat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	v2 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_hstoreUpgrade(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = F_hstoreUpgrade(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v33 = int32(2)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v39 = F_palloc(m, int32(base.Ui32(v32)>>(uint(v33)%32))+int32(base.Ui32(v35)>>(uint(v33)%32)))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v44 = int32(268435455)
	v45 = v43 & v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v48 = v46 & v44
	v49 = v45 + v48
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v49 | int32(-2147483648)
	v53 = int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = (v41+v42&v53)&v53 - int32(32)
	v62 = v39 + int32(4)
	if v48 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v67 = int32(base.Ui32(v65) >> (uint(int32(2)) % 32))
	if v67 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	if v45 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	base.MemoryCopy(m, v39, v30, v67)
	goto L10
L9:
	;
	goto L10
L10:
	;
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v70 = v45 << (uint(int32(3)) % 32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v62)))
	v80 = (v70+v72)<<(uint(int32(2))%32) + int32(32)
	goto L13
L12:
	;
	v80 = int32(32)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v45 | int32(-2147483648)
	return v39
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v90 = int32(base.Ui32(v88) >> (uint(int32(2)) % 32))
	if v90 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v106 = int32(8)
	v107 = v30 + v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v109 = int32(3)
	v111 = int32(2147483640)
	v113 = v107 + v108<<(uint(v109)%32)&v111
	v115 = v25 + v106
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v121 = v115 + v116<<(uint(v109)%32)&v111
	v122 = int32(1)
	v125 = v39 + v106
	v130 = v125 + v49<<(uint(v109)%32)&v111
	v132 = int32(0)
	v134 = v122
	v135 = v122
	v136 = v125
	v140 = v130
	v144 = v2
	v150 = v2
	goto L20
L17:
	;
	base.MemoryCopy(m, v39, v25, v90)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v93 = v48 << (uint(int32(3)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v62+v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v48 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = (v93+v95)<<(uint(int32(2))%32) + int32(32)
	return v39
L20:
	;
	if v135&int32(1) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v452 | int32(-2147483648)
	v456 = v439 - v130
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v446 == v457&int32(268435455) {
		goto L97
	} else {
		goto L98
	}
L22:
	;
	v444 = base.B2i32(v433 < v45)
	v446 = v144 + int32(1)
	v449 = v443 + v150
	v450 = base.B2i32(base.Ui32(v449) < base.Ui32(v48))
	if v444|v450 != 0 {
		v132 = v433
		v134 = v444
		v135 = v450
		v136 = v136 + int32(8)
		v140 = v439
		v144 = v446
		v150 = v449
		goto L20
	} else {
		goto L96
	}
L23:
	;
	v364 = v358 + v107
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v365 < int32(0) {
		goto L81
	} else {
		goto L82
	}
L24:
	;
	v358 = v132 << (uint(int32(3)) % 32)
	v359 = int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v163 = v150 << (uint(int32(3)) % 32)
	if v134&int32(1) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v290 = v163 + v115
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	if v291 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L28:
	;
	v168 = v163 + v115
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v170 = int32(1073741823)
	v171 = v169 & v170
	v173 = v132 << (uint(int32(3)) % 32)
	v174 = v107 + v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v177 = v175 & v170
	if int32(0) <= v169 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v168-int32(4))))
	v186 = v171 - v182&int32(1073741823)
	goto L31
L30:
	;
	v186 = v171
	goto L31
L31:
	;
	if int32(0) <= v175 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174-int32(4))))
	v195 = v177 - v191&int32(1073741823)
	goto L34
L33:
	;
	v195 = v177
	goto L34
L34:
	;
	if v195 == v186 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if int32(0) <= v169 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v195 < v186 {
		v358 = v173
		v359 = int32(1)
		goto L23
	} else {
		goto L63
	}
L38:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v168-int32(4))))
	v205 = v201 & int32(1073741823)
	goto L40
L39:
	;
	v205 = int32(0)
	goto L40
L40:
	;
	v206 = v205 + v121
	if int32(0) <= v175 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v174-int32(4))))
	v215 = v211 & int32(1073741823)
	goto L43
L42:
	;
	v215 = int32(0)
	goto L43
L43:
	;
	v216 = v215 + v113
	if base.Ui32(int32(4)) <= base.Ui32(v186) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	if v278 < int32(0) {
		goto L27
	} else {
		goto L62
	}
L45:
	;
	v278 = int32(0)
	goto L44
L46:
	;
	v252 = v247
	v253 = v248
	v254 = v249
	goto L56
L47:
	;
	if (v206|v216)&int32(3) != 0 {
		v247 = v206
		v248 = v216
		v249 = v186
		goto L46
	} else {
		goto L50
	}
L48:
	;
	v240 = v206
	v241 = v216
	v242 = v186
	goto L49
L49:
	;
	if v242 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L50:
	;
	v224 = v206
	v225 = v216
	v226 = v186
	goto L51
L51:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v229 != v230 {
		v247 = v224
		v248 = v225
		v249 = v226
		goto L46
	} else {
		goto L53
	}
L52:
	;
	v240 = v235
	v241 = v233
	v242 = v237
	goto L49
L53:
	;
	v232 = int32(4)
	v233 = v225 + v232
	v235 = v224 + v232
	v237 = v226 - v232
	if base.Ui32(int32(3)) < base.Ui32(v237) {
		v224 = v235
		v225 = v233
		v226 = v237
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v247 = v240
	v248 = v241
	v249 = v242
	goto L46
L56:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v257 == v258 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v278 = v257 - v258
	goto L44
L58:
	;
	v260 = int32(1)
	v265 = v254 - v260
	if v265 != 0 {
		v252 = v252 + v260
		v253 = v253 + v260
		v254 = v265
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
	goto L45
L62:
	;
	v358 = v173
	v359 = v278
	goto L23
L63:
	;
	goto L27
L64:
	;
	v310 = v290 + int32(4)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v314 = int32(0)
	if v314 <= v311 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v295 = v291 & int32(1073741823)
	v305 = v295
	v306 = v295
	v308 = v121
	goto L64
L66:
	;
	goto L67
L67:
	;
	v296 = int32(1073741823)
	v297 = v291 & v296
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v290-int32(4))))
	v302 = v300 & v296
	v305 = v297 - v302
	v306 = v297
	v308 = v302 + v121
	goto L64
L68:
	;
	v317 = v306
	goto L70
L69:
	;
	v317 = v314
	goto L70
L70:
	;
	v319 = v305 + (v311&int32(1073741823) - v317)
	if v319 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	base.MemoryCopy(m, v140, v308, v319)
	goto L73
L72:
	;
	goto L73
L73:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v323 = v321 & int32(1073741823)
	if int32(0) <= v321 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v290-int32(4))))
	v332 = v323 - v328&int32(1073741823)
	goto L76
L75:
	;
	v332 = v323
	goto L76
L76:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	v336 = int32(0)
	if v336 <= v333 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v339 = v323
	goto L79
L78:
	;
	v339 = v336
	goto L79
L79:
	;
	v340 = v333&int32(1073741823) - v339
	v342 = v340 + (v332 + v140)
	v343 = v342 - v130
	v345 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = (v343 - v340) & v345
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v348&int32(1073741824) | v343&v345
	v433 = v132
	v439 = v342
	v443 = int32(1)
	goto L22
L80:
	;
	v384 = v364 + int32(4)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v388 = int32(0)
	if v388 <= v385 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v369 = v365 & int32(1073741823)
	v379 = v369
	v380 = v369
	v382 = v113
	goto L80
L82:
	;
	goto L83
L83:
	;
	v370 = int32(1073741823)
	v371 = v365 & v370
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v364-int32(4))))
	v376 = v374 & v370
	v379 = v371 - v376
	v380 = v371
	v382 = v376 + v113
	goto L80
L84:
	;
	v391 = v380
	goto L86
L85:
	;
	v391 = v388
	goto L86
L86:
	;
	v393 = v379 + (v385&int32(1073741823) - v391)
	if v393 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	base.MemoryCopy(m, v140, v382, v393)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	v397 = v395 & int32(1073741823)
	if int32(0) <= v395 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v364-int32(4))))
	v406 = v397 - v402&int32(1073741823)
	goto L92
L91:
	;
	v406 = v397
	goto L92
L92:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v410 = int32(0)
	if v410 <= v407 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v413 = v397
	goto L95
L94:
	;
	v413 = v410
	goto L95
L95:
	;
	v414 = v407&int32(1073741823) - v413
	v416 = v414 + (v406 + v140)
	v417 = v416 - v130
	v419 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = (v417 - v414) & v419
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v422&int32(1073741824) | v417&v419
	v433 = v132 + int32(1)
	v439 = v416
	v443 = base.B2i32(v359 == int32(0))
	goto L22
L96:
	;
	goto L21
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = (v446<<(uint(int32(3))%32)+v456)<<(uint(int32(2))%32) + int32(32)
	return v39
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v446 | int32(-2147483648)
	if v456 == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	base.MemoryCopy(m, v125+v446<<(uint(int32(3))%32)&int32(2147483640), v130, v456)
	goto L97
}
func F_hstore_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_eq_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int32(0))
	}
}
func F_hstore_exists_all(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_hstoreUpgrade(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_hstoreArrayToPairs(m, v27, v18+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v33 <= int32(0) {
		v194 = int32(1)
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v18 + int32(16)
	return v194
L6:
	;
	v37 = v21 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v40 = v38 & int32(268435455)
	v45 = int32(0)
	v53 = int32(0)
	goto L7
L7:
	;
	if v40 <= v45 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v194 = v185
	goto L5
L9:
	;
	v194 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	v64 = v31 + v53*int32(20)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v67 = v45
	v70 = v40
	goto L12
L12:
	;
	v85 = int32(base.Ui32(v70-v67)>>(uint(int32(1))%32)) + v67
	v88 = v37 + v85<<(uint(int32(3))%32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v91 = v89 & int32(1073741823)
	if int32(0) <= v89 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L8
L14:
	;
	v185 = int32(0)
	v189 = base.B2i32(v184 < v185)
	if v184 < v185 {
		goto L46
	} else {
		goto L47
	}
L15:
	;
	v111 = v109 + (v37 + v40<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	if base.Ui32(v65) < base.Ui32(v104) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88-int32(4))))
	v98 = v96 & int32(1073741823)
	v99 = v91 - v98
	if v99 != v65 {
		v104 = v99
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v91 == v65 {
		v109 = int32(0)
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v109 = v98
	goto L15
L21:
	;
	v104 = v91
	goto L16
L22:
	;
	v108 = int32(1)
	goto L24
L23:
	;
	v108 = int32(-1)
	goto L24
L24:
	;
	v184 = v108
	goto L14
L25:
	;
	if v173 != 0 {
		v184 = v173
		goto L14
	} else {
		goto L43
	}
L26:
	;
	v173 = int32(0)
	goto L25
L27:
	;
	v147 = v142
	v148 = v143
	v149 = v144
	goto L37
L28:
	;
	if (v111|v66)&int32(3) != 0 {
		v142 = v111
		v143 = v66
		v144 = v65
		goto L27
	} else {
		goto L31
	}
L29:
	;
	v135 = v111
	v136 = v66
	v137 = v65
	goto L30
L30:
	;
	if v137 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L31:
	;
	v119 = v111
	v120 = v66
	v121 = v65
	goto L32
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v124 != v125 {
		v142 = v119
		v143 = v120
		v144 = v121
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v135 = v130
	v136 = v128
	v137 = v132
	goto L30
L34:
	;
	v127 = int32(4)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v142 = v135
	v143 = v136
	v144 = v137
	goto L27
L37:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 == v153 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v173 = v152 - v153
	goto L25
L39:
	;
	v155 = int32(1)
	v160 = v149 - v155
	if v160 != 0 {
		v147 = v147 + v155
		v148 = v148 + v155
		v149 = v160
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L26
L43:
	;
	v174 = int32(0)
	if v85 < v174 {
		v194 = v174
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v177 = int32(1)
	v181 = v53 + v177
	if v181 != v33 {
		v45 = v85 + v177
		v53 = v181
		goto L7
	} else {
		goto L45
	}
L45:
	;
	v194 = v177
	goto L5
L46:
	;
	v190 = v85 + int32(1)
	goto L48
L47:
	;
	v190 = v67
	goto L48
L48:
	;
	if v184 < v185 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v191 = v70
	goto L51
L50:
	;
	v191 = v85
	goto L51
L51:
	;
	if v190 < v191 {
		v67 = v190
		v70 = v191
		goto L12
	} else {
		goto L52
	}
L52:
	;
	goto L13
}
func F_hstore_from_record(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
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
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v36 = F_lookup_rowtype_tupdesc_domain(m, v35, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L8
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = F_get_fn_expr_argtype(m, v22, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v33 = v2
	v34 = int32(-1)
	v35 = v24
	goto L1
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v33 = v29
	v34 = v31
	v35 = v32
	goto L1
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v57 == v35 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v51 = F_MemoryContextAlloc(m, v46, v38*int32(40)+int32(16))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	if v43 != v38 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v57 = v45
	v58 = v40
	goto L9
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = int64(0)
	v57 = v2
	v58 = v51
	goto L9
L14:
	;
	v104 = int32(0)
	v107 = F_palloc(m, v38*int32(20))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L29
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v60 == v34 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v65 = v38 * int32(40)
	v67 = v65 + int32(16)
	if v58&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v67)) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v35
	goto L14
L20:
	;
	if v67 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v67 == int32(0) {
		goto L19
	} else {
		goto L28
	}
L23:
	;
	v79 = v58 + v65 + int32(16)
	v81 = v58 + int32(4)
	if base.Ui32(v81) < base.Ui32(v79) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = v79
	goto L26
L25:
	;
	v83 = v81
	goto L26
L26:
	;
	v88 = (v58^int32(-1)+v83)&int32(-4) + int32(4)
	if v88 == int32(0) {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	base.MemoryFill(m, v58, int32(0), v88)
	goto L19
L28:
	;
	base.MemoryFill(m, v58, int32(0), v67)
	goto L19
L29:
	;
	if v33 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v33
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v112
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)) = uint16(v112)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(-1)
	v118 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(base.Ui32(v110) >> (uint(v118) % 32))
	v125 = F_palloc(m, v38<<(uint(v118)%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L33
	}
L31:
	;
	v132 = int32(0)
	v133 = v2
	goto L32
L32:
	;
	if int32(0) < v38 {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v127 = F_palloc(m, v38)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v16+int32(8), v36, v125, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v132 = v127
	v133 = v125
	goto L32
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L71
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L67
	}
L38:
	;
	v141 = int32(0)
	v144 = v104
	goto L41
L39:
	;
	v233 = v104
	goto L40
L40:
	;
	v243 = F_hstoreUniquePairs(m, v107, v233, v16+int32(28))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L5
	} else {
		goto L61
	}
L41:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v158 = v36 + v152<<(uint(int32(4))%32) + v141*int32(100)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+111)))
	if v159 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v233 = v223
	goto L40
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+88))
	v165 = v107 + v144*int32(20)
	v167 = v158 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v167
	v169 = F_strlen(m, v167)
	mBase = m.M
	if base.Ui32(int32(1073741824)) <= base.Ui32(v169) {
		goto L37
	} else {
		goto L46
	}
L44:
	;
	v223 = v144
	goto L45
L45:
	;
	v226 = v141 + int32(1)
	if v226 != v38 {
		v141 = v226
		v144 = v223
		goto L41
	} else {
		goto L60
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v169
	if v132 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+17)) = uint8(v215)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+16)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v213
	v223 = v144 + int32(1)
	goto L45
L48:
	;
	v183 = v58 + int32(16) + v141*int32(40)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v162 != v184 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141+v132))))
	if v174 != int32(1) {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = int32(0)
	v213 = int32(4)
	v214 = int32(1)
	goto L47
L52:
	;
	goto L51
L53:
	;
	F_getTypeOutputInfo(m, v162, v183+int32(4), v16+int32(28))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v133+v141<<(uint(int32(2))%32))))
	v206 = F_OutputFunctionCall(m, v183+int32(12), v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L58
	}
L56:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+20))
	F_fmgr_info_cxt(m, v192, v183+int32(12), v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v162
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v206
	v210 = F_strlen(m, v206)
	mBase = m.M
	if base.Ui32(int32(1073741824)) <= base.Ui32(v210) {
		goto L36
	} else {
		goto L59
	}
L59:
	;
	v213 = v210
	v214 = int32(0)
	goto L47
L60:
	;
	goto L42
L61:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v246 = F_hstorePairs(m, v107, v243, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if int32(0) <= v248 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	F_DecrTupleDescRefCount(m, v36)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L5
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	m.G0 = v16 + int32(32)
	return v246
L66:
	;
	goto L65
L67:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_hstore_from_record_0), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_hstore_from_record_1), int32(413), int32(_a_F_hstore_from_record_2))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
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
	F_errcode(m, int32(16777346))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(_a_F_hstore_from_record_3), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_hstore_from_record_1), int32(433), int32(_a_F_hstore_from_record_4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_hstoreUpgrade(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int32(4)
		v10 = v5 + v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v15 = int32(base.Ui32(v11)>>(uint(int32(2))%32)) - v9
		v21 = v15 - int32(1636608432)
		if v10&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v15) {
				v130 = v10
				v131 = v15
				v132 = v21
				v133 = v21
				v134 = v21
				for {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
					v137 = v136 + v133
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					v141 = v140 + v134
					v143 = int32(4)
					v145 = v138 + v132 - v141 ^ base.I32_rotl(v141, v143)
					v149 = v137 - v145 ^ base.I32_rotl(v145, int32(6))
					v150 = v141 + v137
					v151 = v145 + v150
					v152 = v149 + v151
					v156 = v150 - v149 ^ base.I32_rotl(v149, int32(8))
					v160 = v151 - v156 ^ base.I32_rotl(v156, int32(16))
					v164 = v152 - v160 ^ base.I32_rotl(v160, int32(19))
					v165 = v156 + v152
					v166 = v160 + v165
					v167 = v164 + v166
					v171 = v165 - v164 ^ base.I32_rotl(v164, v143)
					v172 = int32(12)
					v173 = v130 + v172
					v175 = v131 - v172
					if base.Ui32(int32(11)) < base.Ui32(v175) {
						v130 = v173
						v131 = v175
						v132 = v166
						v133 = v167
						v134 = v171
						continue
					} else {
						break
					}
					break
				}
				v178 = v173
				v179 = v175
				v180 = v166
				v181 = v167
				v182 = v171
			} else {
				v178 = v10
				v179 = v15
				v180 = v21
				v181 = v21
				v182 = v21
			}
			switch v179 - int32(1) {
			case 0:
				v241 = v180
				v242 = v181
				v243 = v182
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 1:
				v234 = v180
				v235 = v181
				v236 = v182
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 2:
				v227 = v180
				v228 = v181
				v229 = v182
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 3:
				v221 = v181
				v222 = v182
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 4:
				v217 = v181
				v218 = v182
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 5:
				v211 = v181
				v212 = v182
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 6:
				v205 = v181
				v206 = v182
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 7:
				v200 = v182
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 8:
				v195 = v182
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 9:
				v190 = v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			case 10:
				v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+10)))
				v190 = v186<<(uint(int32(24))%32) + v182
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+9)))
				v195 = v191<<(uint(int32(16))%32) + v190
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
				v200 = v196<<(uint(int32(8))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
				v205 = v201<<(uint(int32(24))%32) + v181
				v206 = v200
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
				v211 = v207<<(uint(int32(16))%32) + v205
				v212 = v206
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
				v217 = v213<<(uint(int32(8))%32) + v211
				v218 = v212
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
				v221 = v217 + v219
				v222 = v218
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
				v227 = v223<<(uint(int32(24))%32) + v180
				v228 = v221
				v229 = v222
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
				v234 = v230<<(uint(int32(16))%32) + v227
				v235 = v228
				v236 = v229
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
				v241 = v237<<(uint(int32(8))%32) + v234
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
				v248 = v241 + v244
				v249 = v242
				v250 = v243
			default:
				v248 = v180
				v249 = v181
				v250 = v182
			}
		} else {
			if base.Ui32(v15) < base.Ui32(int32(12)) {
				v76 = v10
				v77 = v15
				v78 = v21
				v79 = v21
				v80 = v21
			} else {
				v28 = v10
				v29 = v15
				v30 = v21
				v31 = v21
				v32 = v21
				for {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
					v35 = v34 + v31
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
					v39 = v38 + v32
					v41 = int32(4)
					v43 = v36 + v30 - v39 ^ base.I32_rotl(v39, v41)
					v47 = v35 - v43 ^ base.I32_rotl(v43, int32(6))
					v48 = v39 + v35
					v49 = v43 + v48
					v50 = v47 + v49
					v54 = v48 - v47 ^ base.I32_rotl(v47, int32(8))
					v58 = v49 - v54 ^ base.I32_rotl(v54, int32(16))
					v62 = v50 - v58 ^ base.I32_rotl(v58, int32(19))
					v63 = v54 + v50
					v64 = v58 + v63
					v65 = v62 + v64
					v69 = v63 - v62 ^ base.I32_rotl(v62, v41)
					v70 = int32(12)
					v71 = v28 + v70
					v73 = v29 - v70
					if base.Ui32(int32(11)) < base.Ui32(v73) {
						v28 = v71
						v29 = v73
						v30 = v64
						v31 = v65
						v32 = v69
						continue
					} else {
						break
					}
					break
				}
				v76 = v71
				v77 = v73
				v78 = v64
				v79 = v65
				v80 = v69
			}
			switch v77 - int32(1) {
			case 0:
				v127 = v78
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 1:
				v122 = v78
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 2:
				v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)))
				v122 = v118<<(uint(int32(16))%32) + v78
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
				v248 = v127 + v128
				v249 = v79
				v250 = v80
			case 3:
				v115 = v79
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 4:
				v112 = v79
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 5:
				v107 = v79
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 6:
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+6)))
				v107 = v103<<(uint(int32(16))%32) + v79
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+5)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
				v115 = v112 + v113
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v248 = v116 + v78
				v249 = v115
				v250 = v80
			case 7:
				v98 = v80
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 8:
				v93 = v80
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 9:
				v88 = v80
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
				v93 = v89<<(uint(int32(16))%32) + v88
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			case 10:
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+10)))
				v88 = v84<<(uint(int32(24))%32) + v80
				v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+9)))
				v93 = v89<<(uint(int32(16))%32) + v88
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
				v98 = v94<<(uint(int32(8))%32) + v93
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
				v248 = v99 + v78
				v249 = v101 + v79
				v250 = v98
			default:
				v248 = v78
				v249 = v79
				v250 = v80
			}
		}
		v253 = int32(14)
		v255 = v249 ^ v250 - base.I32_rotl(v249, v253)
		v259 = v255 ^ v248 - base.I32_rotl(v255, int32(11))
		v263 = v259 ^ v249 - base.I32_rotl(v259, int32(25))
		v267 = v263 ^ v255 - base.I32_rotl(v263, int32(16))
		v271 = v267 ^ v259 - base.I32_rotl(v267, int32(4))
		v275 = v271 ^ v263 - base.I32_rotl(v271, v253)
		v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v280 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v283 = m.ExcPending
			if v283 != 0 {
				return int32(0)
			} else {
				return v275 ^ v267 - base.I32_rotl(v275, int32(24))
			}
		} else {
			return v275 ^ v267 - base.I32_rotl(v275, int32(24))
		}
	}
}
func F_hstore_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_lt_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v6) >> (uint(int32(31)) % 32))
	}
}
func F_hstore_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
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
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v454 int32
	_ = v454
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	v2 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_hstoreUpgrade(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v21 = v19 & int32(268435455)
		if v21 != 0 {
			v23 = v15 + int32(8)
			v26 = v23 + v21<<(uint(int32(3))%32)
			v28 = v2
			v30 = v2
			for {
				v42 = v23 + v28<<(uint(int32(3))%32)
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
				v45 = v43 & int32(1073741823)
				if int32(0) <= v43 {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v42-int32(4))))
					v54 = v45 - v50&int32(1073741823)
				} else {
					v54 = v45
				}
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				if v57&int32(1073741824) != 0 {
					v73 = int32(4)
				} else {
					v61 = int32(1073741823)
					if v57 < int32(0) {
						v68 = v57 & v61
					} else {
						v68 = v57 - v43&v61
					}
					v73 = v68<<(uint(int32(1))%32) + int32(2)
				}
				v77 = v73 + (v54<<(uint(int32(1))%32) + v30) + int32(6)
				v79 = v28 + int32(1)
				if v79 != v21 {
					v28 = v79
					v30 = v77
					continue
				} else {
					break
				}
				break
			}
			v81 = F_palloc(m, v77)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v83 = v81
				v92 = v2
				for {
					v96 = int32(34)
					*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v96)
					v100 = v23 + v92<<(uint(int32(3))%32)
					v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
					if v101 < int32(0) {
						v115 = v101 & int32(1073741823)
						v117 = v26
					} else {
						v106 = int32(1073741823)
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v100-int32(4))))
						v112 = v110 & v106
						v115 = v101&v106 - v112
						v117 = v112 + v26
					}
					v119 = v83 + int32(1)
					if v115 <= int32(0) {
						v260 = v119
					} else {
						v124 = v115 & int32(3)
						if v124 != 0 {
							v125 = v117
							v126 = v119
							v131 = int32(0)
							for {
								v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
								if base.B2i32(v138 != int32(92))&base.B2i32(v138 != int32(34)) == int32(0) {
									v146 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v146)
									v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
									v151 = v126 + int32(1)
									v152 = v148
								} else {
									v151 = v126
									v152 = v138
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v152)
								v154 = int32(1)
								v155 = v151 + v154
								v157 = v125 + v154
								v159 = v131 + v154
								if v159 != v124 {
									v125 = v157
									v126 = v155
									v131 = v159
									continue
								} else {
									break
								}
								break
							}
							v161 = v157
							v162 = v155
						} else {
							v161 = v117
							v162 = v119
						}
						if base.Ui32(v115) < base.Ui32(int32(4)) {
							v260 = v162
						} else {
							v176 = v161
							v177 = v162
							for {
								v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
								if base.B2i32(v189 != int32(92))&base.B2i32(v189 != int32(34)) == int32(0) {
									v197 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v197)
									v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
									v202 = v177 + int32(1)
									v203 = v199
								} else {
									v202 = v177
									v203 = v189
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v203)
								v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
								if base.B2i32(v205 == int32(92))|base.B2i32(v205 == int32(34)) != 0 {
									v211 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)) = uint8(v211)
									v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+1)))
									v218 = v213
									v219 = v202 + int32(2)
								} else {
									v218 = v205
									v219 = v202 + int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v219))) = uint8(v218)
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+2)))
								if base.B2i32(v221 == int32(92))|base.B2i32(v221 == int32(34)) != 0 {
									v227 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)) = uint8(v227)
									v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+2)))
									v234 = v229
									v235 = v219 + int32(2)
								} else {
									v234 = v221
									v235 = v219 + int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v235))) = uint8(v234)
								v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+3)))
								if base.B2i32(v237 == int32(92))|base.B2i32(v237 == int32(34)) != 0 {
									v243 = int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)) = uint8(v243)
									v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+3)))
									v250 = v245
									v251 = v235 + int32(2)
								} else {
									v250 = v237
									v251 = v235 + int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v251))) = uint8(v250)
								v254 = v251 + int32(1)
								v256 = v176 + int32(4)
								if v256-v117 < v115 {
									v176 = v256
									v177 = v254
									continue
								} else {
									break
								}
								break
							}
							v260 = v254
						}
					}
					v272 = int32(62)
					*(*uint8)(unsafe.Add(mBase, uint32(v260)+2)) = uint8(v272)
					v274 = int32(_a_F_hstore_out_0)
					*(*uint16)(unsafe.Add(mBase, uint32(v260))) = uint16(v274)
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
					if v276&int32(64) != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(v260)+3)) = int32(1280070990)
						v471 = v260 + int32(7)
					} else {
						v283 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v260)+3)) = uint8(v283)
						v285 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
						if v285 < int32(0) {
							v298 = v285 & int32(1073741823)
							v299 = v26
						} else {
							v290 = int32(1073741823)
							v292 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
							v294 = v292 & v290
							v298 = v285&v290 - v294
							v299 = v294 + v26
						}
						v301 = v260 + int32(4)
						if v298 <= int32(0) {
							v442 = v301
						} else {
							v306 = v298 & int32(3)
							if v306 != 0 {
								v307 = v299
								v308 = v301
								v313 = int32(0)
								for {
									v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
									if base.B2i32(v320 != int32(92))&base.B2i32(v320 != int32(34)) == int32(0) {
										v328 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v308))) = uint8(v328)
										v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
										v333 = v308 + int32(1)
										v334 = v330
									} else {
										v333 = v308
										v334 = v320
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v334)
									v336 = int32(1)
									v337 = v333 + v336
									v339 = v307 + v336
									v341 = v313 + v336
									if v341 != v306 {
										v307 = v339
										v308 = v337
										v313 = v341
										continue
									} else {
										break
									}
									break
								}
								v343 = v339
								v344 = v337
							} else {
								v343 = v299
								v344 = v301
							}
							if base.Ui32(v298) < base.Ui32(int32(4)) {
								v442 = v344
							} else {
								v358 = v343
								v359 = v344
								for {
									v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
									if base.B2i32(v371 != int32(92))&base.B2i32(v371 != int32(34)) == int32(0) {
										v379 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v379)
										v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358))))
										v384 = v359 + int32(1)
										v385 = v381
									} else {
										v384 = v359
										v385 = v371
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v384))) = uint8(v385)
									v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
									if base.B2i32(v387 == int32(92))|base.B2i32(v387 == int32(34)) != 0 {
										v393 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v384)+1)) = uint8(v393)
										v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+1)))
										v400 = v395
										v401 = v384 + int32(2)
									} else {
										v400 = v387
										v401 = v384 + int32(1)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v400)
									v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+2)))
									if base.B2i32(v403 == int32(92))|base.B2i32(v403 == int32(34)) != 0 {
										v409 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v401)+1)) = uint8(v409)
										v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+2)))
										v416 = v411
										v417 = v401 + int32(2)
									} else {
										v416 = v403
										v417 = v401 + int32(1)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v417))) = uint8(v416)
									v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+3)))
									if base.B2i32(v419 == int32(92))|base.B2i32(v419 == int32(34)) != 0 {
										v425 = int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v417)+1)) = uint8(v425)
										v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+3)))
										v432 = v427
										v433 = v417 + int32(2)
									} else {
										v432 = v419
										v433 = v417 + int32(1)
									}
									*(*uint8)(unsafe.Add(mBase, uint32(v433))) = uint8(v432)
									v436 = v433 + int32(1)
									v438 = v358 + int32(4)
									if v438-v299 < v298 {
										v358 = v438
										v359 = v436
										continue
									} else {
										break
									}
									break
								}
								v442 = v436
							}
						}
						v454 = int32(34)
						*(*uint8)(unsafe.Add(mBase, uint32(v442))) = uint8(v454)
						v471 = v442 + int32(1)
					}
					v473 = v92 + int32(1)
					if v21 != v473 {
						v475 = int32(_a_F_hstore_out_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v471))) = uint16(v475)
						v479 = v471 + int32(2)
					} else {
						v479 = v471
					}
					if v21 != v473 {
						v83 = v479
						v92 = v473
						continue
					} else {
						break
					}
					break
				}
				v481 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v479))) = uint8(v481)
				return v81
			}
		} else {
			v485 = F_pstrdup(m, int32(_a_F_hstore_out_2))
			mBase = m.M
			v486 = m.ExcPending
			if v486 != 0 {
				return int32(0)
			} else {
				return v485
			}
		}
	}
}
func F_hstore_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
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
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L6
	} else {
		goto L42
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L34
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L6
	} else {
		goto L30
	}
L5:
	;
	m.G0 = v11 + int32(16)
	return v91
L6:
	;
	return int32(0)
L7:
	;
	if v15 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v22 = F_palloc(m, int32(8))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(53687092)) <= base.Ui32(v15) {
		goto L4
	} else {
		goto L12
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22))) = int64(-9223372036854775776)
	v91 = v22
	goto L5
L12:
	;
	v30 = F_palloc(m, v15*int32(20))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v38 = int32(0)
	goto L14
L14:
	;
	v41 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L16
	}
L15:
	;
	v86 = F_hstoreUniquePairs(m, v30, v15, v11+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L28
	}
L16:
	;
	if v41 < int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v47 = v30 + v38*int32(20)
	v50 = F_pq_getmsgtext(m, v13, v41, v11+int32(12))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v53) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v56 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+17)) = uint8(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v53
	v60 = F_pq_getmsgint(m, v13, int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L21
	}
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+16)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v77
	v82 = v38 + int32(1)
	if v82 != v15 {
		v38 = v82
		goto L14
	} else {
		goto L27
	}
L21:
	;
	if v60 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v64
	v77 = v64
	v78 = int32(1)
	goto L20
L23:
	;
	goto L24
L24:
	;
	v70 = F_pq_getmsgtext(m, v13, v60, v11+int32(12))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v70
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if base.Ui32(int32(1073741824)) <= base.Ui32(v74) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v77 = v74
	v78 = int32(0)
	goto L20
L27:
	;
	goto L15
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v89 = F_hstorePairs(m, v30, v86, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v91 = v89
	goto L5
L30:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(53687091)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v15
	F_errmsg(m, int32(_a_F_hstore_recv_0), v11)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_hstore_recv_1), int32(525), int32(_a_F_hstore_recv_2))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(_a_F_hstore_recv_3), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_hstore_recv_1), int32(536), int32(_a_F_hstore_recv_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_hstore_recv_4), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_hstore_recv_1), int32(413), int32(_a_F_hstore_recv_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(16777346))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(_a_F_hstore_recv_6), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_hstore_recv_1), int32(433), int32(_a_F_hstore_recv_7))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_slice_to_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v272 int32
	_ = v272
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v354 int32
	_ = v354
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_hstoreUpgrade(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_deconstruct_array_builtin(m, v32, int32(25), v23+int32(12), v23+int32(8), v23+int32(4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v43 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v23 + int32(16)
	return v354
L6:
	;
	v47 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v51 = F_palloc(m, v43<<(uint(int32(2))%32))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v354 = v47
	goto L5
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v54 = F_palloc(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if int32(0) < v56 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v60 = v26 + int32(8)
	v69 = int32(0)
	v78 = v56
	goto L15
L13:
	;
	goto L14
L14:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v324 = v32 + int32(16)
	v332 = F_construct_md_array(m, v51, v54, v322, v324, v324+v322<<(uint(int32(2))%32), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L69
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v69))))
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L14
L17:
	;
	v300 = v69 + int32(1)
	if v300 < v291 {
		v69 = v300
		v78 = v291
		goto L15
	} else {
		goto L68
	}
L18:
	;
	v272 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v69+v54))) = uint8(v272)
	*(*int32)(unsafe.Add(mBase, uint32(v51+v69<<(uint(int32(2))%32)))) = int32(0)
	v291 = v78
	goto L17
L19:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v91 = v89 & int32(268435455)
	if v91 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v97 = int32(2)
	v98 = v69 << (uint(v97) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98+v99)))
	v102 = int32(4)
	v103 = v101 + v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v108 = int32(base.Ui32(v104)>>(uint(v97)%32)) - v102
	v110 = int32(0)
	v116 = v91
	goto L21
L21:
	;
	v133 = int32(base.Ui32(v116-v110)>>(uint(int32(1))%32)) + v110
	v136 = v60 + v133<<(uint(int32(3))%32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v139 = v137 & int32(1073741823)
	if int32(0) <= v137 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L18
L23:
	;
	v247 = base.B2i32(v243 < int32(0))
	if v243 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L24:
	;
	v159 = v158 + (v60 + v91<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v108) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if base.Ui32(v108) < base.Ui32(v151) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v136-int32(4))))
	v146 = v144 & int32(1073741823)
	v147 = v139 - v146
	if v147 != v108 {
		v151 = v147
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v139 == v108 {
		v158 = int32(0)
		goto L24
	} else {
		goto L30
	}
L29:
	;
	v158 = v146
	goto L24
L30:
	;
	v151 = v139
	goto L25
L31:
	;
	v156 = int32(1)
	goto L33
L32:
	;
	v156 = int32(-1)
	goto L33
L33:
	;
	v243 = v156
	goto L23
L34:
	;
	if v221 != 0 {
		v243 = v221
		goto L23
	} else {
		goto L52
	}
L35:
	;
	v221 = int32(0)
	goto L34
L36:
	;
	v195 = v190
	v196 = v191
	v197 = v192
	goto L46
L37:
	;
	if (v159|v103)&int32(3) != 0 {
		v190 = v159
		v191 = v103
		v192 = v108
		goto L36
	} else {
		goto L40
	}
L38:
	;
	v183 = v159
	v184 = v103
	v185 = v108
	goto L39
L39:
	;
	if v185 == int32(0) {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v167 = v159
	v168 = v103
	v169 = v108
	goto L41
L41:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	if v172 != v173 {
		v190 = v167
		v191 = v168
		v192 = v169
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v183 = v178
	v184 = v176
	v185 = v180
	goto L39
L43:
	;
	v175 = int32(4)
	v176 = v168 + v175
	v178 = v167 + v175
	v180 = v169 - v175
	if base.Ui32(int32(3)) < base.Ui32(v180) {
		v167 = v178
		v168 = v176
		v169 = v180
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v190 = v183
	v191 = v184
	v192 = v185
	goto L36
L46:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v200 == v201 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v221 = v200 - v201
	goto L34
L48:
	;
	v203 = int32(1)
	v208 = v197 - v203
	if v208 != 0 {
		v195 = v195 + v203
		v196 = v196 + v203
		v197 = v208
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L35
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v222&int32(1073741824) != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v226 = int32(0)
	v228 = base.B2i32(v226 <= v222)
	if v226 <= v222 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v229 = v139
	goto L56
L55:
	;
	v229 = v226
	goto L56
L56:
	;
	if v226 <= v222 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v234 = v222 - v139
	goto L59
L58:
	;
	v234 = v222 & int32(1073741823)
	goto L59
L59:
	;
	v235 = F_cstring_to_text_with_len(m, v60+v30<<(uint(int32(3))%32)&int32(2147483640)+v229, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51+v98))) = v235
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69+v54))) = uint8(v239)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v291 = v241
	goto L17
L61:
	;
	v248 = v133 + int32(1)
	goto L63
L62:
	;
	v248 = v110
	goto L63
L63:
	;
	if v243 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v249 = v116
	goto L66
L65:
	;
	v249 = v133
	goto L66
L66:
	;
	if v248 < v249 {
		v110 = v248
		v116 = v249
		goto L21
	} else {
		goto L67
	}
L67:
	;
	goto L22
L68:
	;
	goto L16
L69:
	;
	v354 = v332
	goto L5
}
func F_hstore_slice_to_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
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
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v308 int32
	_ = v308
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_hstoreUpgrade(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v32 = F_pg_detoast_datum(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = F_hstoreArrayToPairs(m, v32, v23+int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v23 + int32(16)
	return v308
L6:
	;
	v41 = int32(0)
	v44 = F_hstorePairs(m, v41, v41, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v48 = F_palloc(m, v38*int32(20))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v308 = v44
	goto L5
L10:
	;
	if int32(0) < v38 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = v26 + int32(8)
	v54 = int32(3)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v61 = v59 & int32(268435455)
	v66 = int32(0)
	v75 = v2
	v79 = v2
	v82 = v2
	goto L14
L12:
	;
	v275 = v2
	v282 = v2
	goto L13
L13:
	;
	v286 = F_hstorePairs(m, v48, v275, v282)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L64
	}
L14:
	;
	if v61 <= v66 {
		v243 = v66
		v252 = v75
		v259 = v82
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v275 = v252
	v282 = v259
	goto L13
L16:
	;
	v264 = v79 + int32(1)
	if v264 != v38 {
		v66 = v243
		v75 = v252
		v79 = v264
		v82 = v259
		goto L14
	} else {
		goto L63
	}
L17:
	;
	v89 = v36 + v79*int32(20)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v92 = v66
	v94 = v61
	goto L18
L18:
	;
	v115 = int32(base.Ui32(v94-v92)>>(uint(int32(1))%32)) + v92
	v118 = v53 + v115<<(uint(int32(3))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v121 = v119 & int32(1073741823)
	if int32(0) <= v119 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v243 = v240
	v252 = v75
	v259 = v82
	goto L16
L20:
	;
	v239 = base.B2i32(v235 < int32(0))
	if v235 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L21:
	;
	v141 = v140 + (v53 + v61<<(uint(v54)%32))
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	if base.Ui32(v90) < base.Ui32(v133) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118-int32(4))))
	v128 = v126 & int32(1073741823)
	v129 = v121 - v128
	if v129 != v90 {
		v133 = v129
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v121 == v90 {
		v140 = int32(0)
		goto L21
	} else {
		goto L27
	}
L26:
	;
	v140 = v128
	goto L21
L27:
	;
	v133 = v121
	goto L22
L28:
	;
	v138 = int32(1)
	goto L30
L29:
	;
	v138 = int32(-1)
	goto L30
L30:
	;
	v235 = v138
	goto L20
L31:
	;
	if v203 != 0 {
		v235 = v203
		goto L20
	} else {
		goto L49
	}
L32:
	;
	v203 = int32(0)
	goto L31
L33:
	;
	v177 = v172
	v178 = v173
	v179 = v174
	goto L43
L34:
	;
	if (v141|v91)&int32(3) != 0 {
		v172 = v141
		v173 = v91
		v174 = v90
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v165 = v141
	v166 = v91
	v167 = v90
	goto L36
L36:
	;
	if v167 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v149 = v141
	v150 = v91
	v151 = v90
	goto L38
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v154 != v155 {
		v172 = v149
		v173 = v150
		v174 = v151
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v165 = v160
	v166 = v158
	v167 = v162
	goto L36
L40:
	;
	v157 = int32(4)
	v158 = v150 + v157
	v160 = v149 + v157
	v162 = v151 - v157
	if base.Ui32(int32(3)) < base.Ui32(v162) {
		v149 = v160
		v150 = v158
		v151 = v162
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v172 = v165
	v173 = v166
	v174 = v167
	goto L33
L43:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v182 == v183 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v203 = v182 - v183
	goto L31
L45:
	;
	v185 = int32(1)
	v190 = v179 - v185
	if v190 != 0 {
		v177 = v177 + v185
		v178 = v178 + v185
		v179 = v190
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	v206 = v48 + v75*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v206)+8)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = v91
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+17)) = uint8(v210)
	v215 = int32(base.Ui32(v209)>>(uint(int32(30))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+16)) = uint8(v215)
	v218 = v209 & int32(1073741823)
	v221 = base.B2i32(v210 <= v209)
	if v210 <= v209 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v222 = v218 - v121
	goto L52
L51:
	;
	v222 = v218
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+12)) = v222
	if v210 <= v209 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v225 = v121
	goto L55
L54:
	;
	v225 = int32(0)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v206)+4)) = v53 + v30<<(uint(v54)%32)&int32(2147483640) + v225
	v230 = int32(1)
	v243 = v115 + v230
	v252 = v75 + v230
	v259 = v90 + v82 + v222
	goto L16
L56:
	;
	v240 = v115 + int32(1)
	goto L58
L57:
	;
	v240 = v92
	goto L58
L58:
	;
	if v235 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v241 = v94
	goto L61
L60:
	;
	v241 = v115
	goto L61
L61:
	;
	if v240 < v241 {
		v92 = v240
		v94 = v241
		goto L18
	} else {
		goto L62
	}
L62:
	;
	goto L19
L63:
	;
	goto L15
L64:
	;
	v308 = v286
	goto L5
}
func F_hstore_subscript_fetch(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if v9 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v13)
	return
L2:
	;
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_hstoreUpgrade(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v193 = v173 + v188<<(uint(int32(3))%32)&int32(2147483640)
	if v177 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	return
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v24 = int32(1)
	v25 = v22 + v24
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v30 = v28 & v24
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v31 = v25
	goto L10
L9:
	;
	v31 = v22 + int32(4)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L26
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v37 == int32(18) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v48 = int32(1)
	if v30 != 0 {
		v58 = int32(base.Ui32(v28)>>(uint(v48)%32)) - v48
		goto L11
	} else {
		goto L21
	}
L15:
	;
	v40 = int32(16)
	goto L17
L16:
	;
	v40 = int32(0)
	goto L17
L17:
	;
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v47 = int32(4)
	goto L20
L19:
	;
	v47 = v40
	goto L20
L20:
	;
	v58 = v47
	goto L11
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L11
L22:
	;
	if int32(0) <= v159 {
		goto L54
	} else {
		goto L55
	}
L23:
	;
	goto L22
L26:
	;
	v67 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v71 = v69 & int32(268435455)
	if v67 < v71 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = v17 + int32(8)
	v83 = v67
	v84 = v71
	goto L31
L29:
	;
	goto L30
L30:
	;
	v159 = int32(-1)
	goto L23
L31:
	;
	v92 = int32(base.Ui32(v84-v83)>>(uint(int32(1))%32)) + v83
	v95 = v74 + v92<<(uint(int32(3))%32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v98 = v96 & int32(1073741823)
	if int32(0) <= v96 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L30
L33:
	;
	v129 = base.B2i32(v124 < int32(0))
	if v124 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v119 = F_memcmp(m, v117+(v74+v71<<(uint(int32(3))%32)), v31, v58)
	mBase = m.M
	if v119 != 0 {
		v124 = v119
		goto L33
	} else {
		goto L44
	}
L35:
	;
	if base.Ui32(v58) < base.Ui32(v110) {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95-int32(4))))
	v105 = v103 & int32(1073741823)
	v106 = v98 - v105
	if v106 != v58 {
		v110 = v106
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v58 == v98 {
		v117 = int32(0)
		goto L34
	} else {
		goto L40
	}
L39:
	;
	v117 = v105
	goto L34
L40:
	;
	v110 = v98
	goto L35
L41:
	;
	v115 = int32(1)
	goto L43
L42:
	;
	v115 = int32(-1)
	goto L43
L43:
	;
	v124 = v115
	goto L33
L44:
	;
	v159 = v92
	goto L23
L46:
	;
	v130 = v92 + int32(1)
	goto L48
L47:
	;
	v130 = v83
	goto L48
L48:
	;
	if v124 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v131 = v84
	goto L51
L50:
	;
	v131 = v92
	goto L51
L51:
	;
	if v130 < v131 {
		v83 = v130
		v84 = v131
		goto L31
	} else {
		goto L52
	}
L52:
	;
	goto L32
L54:
	;
	v173 = v17 + int32(8)
	v176 = v173 + v159<<(uint(int32(3))%32)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v177&int32(1073741824) == int32(0) {
		goto L4
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v186)
	return
L57:
	;
	goto L56
L58:
	;
	v206 = F_cstring_to_text_with_len(m, v204, v203)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L5
	} else {
		goto L62
	}
L59:
	;
	v203 = v177 & int32(1073741823)
	v204 = v193
	goto L58
L60:
	;
	goto L61
L61:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v200 = v198 & int32(1073741823)
	v203 = v177 - v200
	v204 = v193 + v200
	goto L58
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v208))) = v206
	return
}
func F_hstore_to_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_hstoreUpgrade(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v184
L2:
	;
	return int32(0)
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v21 = v19 & int32(268435455)
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = F_cstring_to_text_with_len(m, int32(_a_F_hstore_to_json_0), int32(2))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v29 = v15 + int32(8)
	v32 = v29 + v21<<(uint(int32(3))%32)
	F_initStringInfo(m, v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v184 = v26
	goto L1
L8:
	;
	F_appendStringInfoChar(m, v12, int32(123))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v21 != int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(0)
	goto L13
L11:
	;
	v108 = int32(0)
	goto L12
L12:
	;
	v116 = v29 + v108<<(uint(int32(3))%32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v54 = v29 + v43<<(uint(int32(3))%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v55 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v108 = v104
	goto L12
L15:
	;
	F_escape_json_with_len(m, v12, v71, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v69 = v55 & int32(1073741823)
	v71 = v32
	goto L15
L17:
	;
	goto L18
L18:
	;
	v60 = int32(1073741823)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54-int32(4))))
	v66 = v64 & v60
	v69 = v55&v60 - v66
	v71 = v32 + v66
	goto L15
L19:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v77&int32(1073741824) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L2
	} else {
		goto L31
	}
L22:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_3))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v77 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L21
L26:
	;
	F_escape_json_with_len(m, v12, v94, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L2
	} else {
		goto L30
	}
L27:
	;
	v92 = v77 & int32(1073741823)
	v94 = v32
	goto L26
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v89 = v87 & int32(1073741823)
	v92 = v77 - v89
	v94 = v89 + v32
	goto L26
L30:
	;
	goto L21
L31:
	;
	v104 = v43 + int32(1)
	if v43 != v21-int32(2) {
		v43 = v104
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	F_escape_json_with_len(m, v12, v133, v131)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L2
	} else {
		goto L37
	}
L34:
	;
	v131 = v117 & int32(1073741823)
	v133 = v32
	goto L33
L35:
	;
	goto L36
L36:
	;
	v122 = int32(1073741823)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	v128 = v126 & v122
	v131 = v117&v122 - v128
	v133 = v32 + v128
	goto L33
L37:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v139&int32(1073741824) != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v21 != v108+int32(1) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L2
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v139 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L39
L44:
	;
	F_escape_json_with_len(m, v12, v156, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L48
	}
L45:
	;
	v154 = v139 & int32(1073741823)
	v156 = v32
	goto L44
L46:
	;
	goto L47
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v151 = v149 & int32(1073741823)
	v154 = v139 - v151
	v156 = v151 + v32
	goto L44
L48:
	;
	goto L39
L49:
	;
	F_appendStringInfoString(m, v12, int32(_a_F_hstore_to_json_2))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	F_appendStringInfoChar(m, v12, int32(125))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v173 = F_cstring_to_text_with_len(m, v171, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v184 = v173
	goto L1
}
func F_hstore_to_jsonb(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_hstoreUpgrade(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v19
	v25 = F_pushJsonbValue(m, v11+int32(44), int32(6), v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = v18 & int32(268435455)
	if v28 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v14 + int32(8)
	v33 = v30 + v28<<(uint(int32(3))%32)
	v38 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v122 = F_pushJsonbValue(m, v11+int32(44), int32(7), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L24
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = int32(1)
	v47 = v30 + v38<<(uint(int32(3))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = v48 & int32(1073741823)
	if v48 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v66 + v33
	v74 = F_pushJsonbValue(m, v11+int32(44), int32(1), v11+int32(24))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50
	v66 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v56 = v47 - int32(4)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v50 - v57&v58
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v66 = v62 & v58
	goto L9
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v76&int32(1073741824) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v105 = F_pushJsonbValue(m, v11+int32(44), int32(2), v11+int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(1)
	if v76 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v97 + v33
	goto L14
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v76 & int32(1073741823)
	v97 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v90 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v76 - v89&v90
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v97 = v94 & v90
	goto L18
L22:
	;
	v108 = v38 + int32(1)
	if v108 != v28 {
		v38 = v108
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L8
L24:
	;
	v124 = F_JsonbValueToJsonb(m, v122)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	m.G0 = v11 + int32(48)
	return v124
}
