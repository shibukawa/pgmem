package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResourceOwnerCreate(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v7 = F_MemoryContextAllocZero(m, v5, int32(360))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
		} else {
		}
		v17 = v7 + int32(352)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+356)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+352)) = v17
		return v7
	}
}
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
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
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
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
	var v113 int32
	_ = v113
	var v129 int32
	_ = v129
	var v143 int32
	_ = v143
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v12 != int32(1) {
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
		if base.Ui32(int32(32)) <= base.Ui32(v15) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+288))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v19+v15) < base.Ui32(v18) {
				v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
				if v155 != 0 {
					v162 = int32(0)
					for {
						v172 = l0 + int32(24) + v162<<(uint(int32(3))%32)
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
						v175 = int32(16)
						v179 = (int32(base.Ui32(v174)>>(uint(v175)%32)) ^ v174) * int32(-2048144789)
						v184 = (int32(base.Ui32(v179)>>(uint(int32(13))%32)) ^ v179) * int32(-1028477387)
						v187 = int32(base.Ui32(v184)>>(uint(v175)%32)) ^ v184
						v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
						v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
						v202 = v173 + v187<<(uint(int32(6))%32) + int32(base.Ui32(v187)>>(uint(int32(2))%32)) - int32(1640531527) ^ v187
						for {
							v212 = v202 & (v197 - int32(1))
							v216 = v212 << (uint(int32(3)) % 32)
							v217 = v200 + v216
							v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
							if v218 != 0 {
								v202 = v212 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v217))) = v174
						v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
						*(*int32)(unsafe.Add(mBase, uint32(v220+v216)+4)) = v173
						v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v224 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223 + v224
						v228 = v162 + v224
						v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
						if base.Ui32(v228) < base.Ui32(v229) {
							v162 = v228
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v242 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v242)
				return
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
				v24 = *(*int32)(unsafe.Add(mBase, _consts[12]))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
				if v25 != 0 {
					v29 = v25 << (uint(int32(1)) % 32)
				} else {
					v29 = int32(64)
				}
				v32 = F_MemoryContextAllocZero(m, v24, v29<<(uint(int32(3))%32))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v29
					*(*int32)(unsafe.Add(mBase, uint32(l0)+280)) = v32
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
					v39 = v29 - int32(32)
					v43 = int32(base.Ui32(v29)>>(uint(int32(2))%32)) * int32(3)
					if base.Ui32(v39) < base.Ui32(v43) {
						v45 = v39
					} else {
						v45 = v43
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+288)) = v45
					if v22 == int32(0) {
						v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
						if v155 != 0 {
							v162 = int32(0)
							for {
								v172 = l0 + int32(24) + v162<<(uint(int32(3))%32)
								v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
								v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
								v175 = int32(16)
								v179 = (int32(base.Ui32(v174)>>(uint(v175)%32)) ^ v174) * int32(-2048144789)
								v184 = (int32(base.Ui32(v179)>>(uint(int32(13))%32)) ^ v179) * int32(-1028477387)
								v187 = int32(base.Ui32(v184)>>(uint(v175)%32)) ^ v184
								v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
								v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
								v202 = v173 + v187<<(uint(int32(6))%32) + int32(base.Ui32(v187)>>(uint(int32(2))%32)) - int32(1640531527) ^ v187
								for {
									v212 = v202 & (v197 - int32(1))
									v216 = v212 << (uint(int32(3)) % 32)
									v217 = v200 + v216
									v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
									if v218 != 0 {
										v202 = v212 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v217))) = v174
								v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
								*(*int32)(unsafe.Add(mBase, uint32(v220+v216)+4)) = v173
								v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v224 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223 + v224
								v228 = v162 + v224
								v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
								if base.Ui32(v228) < base.Ui32(v229) {
									v162 = v228
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v242 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v242)
						return
					} else {
						if v25 != 0 {
							v52 = int32(0)
							for {
								v62 = v22 + v52<<(uint(int32(3))%32)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
								if v63 != 0 {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
									v65 = int32(16)
									v69 = (int32(base.Ui32(v64)>>(uint(v65)%32)) ^ v64) * int32(-2048144789)
									v74 = (int32(base.Ui32(v69)>>(uint(int32(13))%32)) ^ v69) * int32(-1028477387)
									v77 = int32(base.Ui32(v74)>>(uint(v65)%32)) ^ v74
									v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
									v92 = v63 + v77<<(uint(int32(6))%32) + int32(base.Ui32(v77)>>(uint(int32(2))%32)) - int32(1640531527) ^ v77
									for {
										v102 = v92 & (v87 - int32(1))
										v106 = v102 << (uint(int32(3)) % 32)
										v107 = v90 + v106
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
										if v108 != 0 {
											v92 = v102 + int32(1)
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v64
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
									*(*int32)(unsafe.Add(mBase, uint32(v110+v106)+4)) = v63
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v113 + int32(1)
								} else {
								}
								v129 = v52 + int32(1)
								if v129 != v25 {
									v52 = v129
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						F_pfree(m, v22)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return
						} else {
							v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
							if v155 != 0 {
								v162 = int32(0)
								for {
									v172 = l0 + int32(24) + v162<<(uint(int32(3))%32)
									v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
									v174 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
									v175 = int32(16)
									v179 = (int32(base.Ui32(v174)>>(uint(v175)%32)) ^ v174) * int32(-2048144789)
									v184 = (int32(base.Ui32(v179)>>(uint(int32(13))%32)) ^ v179) * int32(-1028477387)
									v187 = int32(base.Ui32(v184)>>(uint(v175)%32)) ^ v184
									v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
									v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
									v202 = v173 + v187<<(uint(int32(6))%32) + int32(base.Ui32(v187)>>(uint(int32(2))%32)) - int32(1640531527) ^ v187
									for {
										v212 = v202 & (v197 - int32(1))
										v216 = v212 << (uint(int32(3)) % 32)
										v217 = v200 + v216
										v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
										if v218 != 0 {
											v202 = v212 + int32(1)
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v217))) = v174
									v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
									*(*int32)(unsafe.Add(mBase, uint32(v220+v216)+4)) = v173
									v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v224 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223 + v224
									v228 = v162 + v224
									v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
									if base.Ui32(v228) < base.Ui32(v229) {
										v162 = v228
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v242 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v242)
							return
						}
					}
				}
			}
		} else {
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v258 = m.ExcPending
		if v258 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(466861), int32(0))
			mBase = m.M
			v262 = m.ExcPending
			if v262 != 0 {
				return
			} else {
				F_errfinish(m, int32(520654), int32(459), int32(420119))
				mBase = m.M
				v267 = m.ExcPending
				if v267 != 0 {
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
func F_ResourceOwnerForgetAioHandle(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v6
	return
}
func F_ResourceOwnerForgetLock(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	if base.Ui32(v12) <= base.Ui32(int32(15)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	v16 = l0 + int32(292)
	v20 = v12
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(16)
	return
L5:
	;
	if v20 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32)+v16-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v42 = v40 - int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)) = uint8(v42)
	goto L4
L7:
	;
	v27 = v20 - int32(1)
	v30 = v16 + v27<<(uint(int32(2))%32)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if l1 != v31 {
		v20 = v27
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	F_errmsg_internal(m, int32(193146), v10)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(520654), int32(1100), int32(333724))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ResourceOwnerReleaseInternal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v495 int32
	_ = v495
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v582 int32
	_ = v582
	var v588 int32
	_ = v588
	var v589 int64
	_ = v589
	var v590 int64
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int64
	_ = v601
	var v603 int64
	_ = v603
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v795 int32
	_ = v795
	var v796 int64
	_ = v796
	var v797 int64
	_ = v797
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int64
	_ = v808
	var v810 int64
	_ = v810
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v975 int32
	_ = v975
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v28 = v23
	goto L4
L2:
	;
	goto L3
L3:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v71 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_ResourceOwnerReleaseInternal(m, v28, l1, l2, l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if v48 != 0 {
		v28 = v48
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v74)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	if v76 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	goto L14
L14:
	;
	v270 = int32(4562148)
	v271 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	*(*int32)(unsafe.Add(mBase, _consts[181])) = l0
	switch l1 - int32(1) {
	case 0:
		goto L42
	case 1:
		goto L41
	case 2:
		goto L40
	default:
		goto L39
	}
L15:
	;
	F_pg_qsort(m, v241, v223, int32(8), int32(1836))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L38
	}
L16:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v152 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v80 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v223 = v127
	v241 = l0 + int32(24)
	goto L15
L20:
	;
	v134 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v89 = int32(0)
	v90 = v80
	v92 = int32(0)
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v110 = v107 + v92<<(uint(int32(3))%32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v111 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v134 = v122
	goto L16
L25:
	;
	if v89 != v92 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v122 = v89
	v123 = v90
	goto L27
L27:
	;
	v125 = v92 + int32(1)
	if base.Ui32(v125) < base.Ui32(v123) {
		v89 = v122
		v90 = v123
		v92 = v125
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v107+v89<<(uint(int32(3))%32)))) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	v119 = v118
	goto L30
L29:
	;
	v119 = v90
	goto L30
L30:
	;
	v122 = v89 + int32(1)
	v123 = v119
	goto L27
L31:
	;
	goto L24
L32:
	;
	v160 = v134
	v163 = int32(0)
	goto L35
L33:
	;
	v197 = v134
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v197
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v223 = v197
	v241 = v218
	goto L15
L35:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v179 = int32(3)
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)+v163<<(uint(v179)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v178+v160<<(uint(v179)%32)))) = v185
	v187 = int32(1)
	v188 = v160 + v187
	v190 = v163 + v187
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if base.Ui32(v190) < base.Ui32(v191) {
		v160 = v188
		v163 = v190
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v197 = v188
	goto L34
L37:
	;
	goto L36
L38:
	;
	v246 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v246)
	goto L14
L39:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _consts[1208]))
	if v999 != 0 {
		goto L235
	} else {
		goto L236
	}
L40:
	;
	F_ResourceOwnerReleaseAll(m, l0, int32(3), l2)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L6
	} else {
		goto L234
	}
L41:
	;
	if l3 != 0 {
		goto L76
	} else {
		goto L77
	}
L42:
	;
	F_ResourceOwnerReleaseAll(m, l0, int32(1), l2)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v279 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v283 = l0 + int32(352)
	if v279 == v283 {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v286 = l2 ^ int32(1)
	v291 = v279
	goto L46
L46:
	;
	v309 = int32(4556748)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v311 + int32(1)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v322
	goto L48
L47:
	;
	goto L39
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291-int32(4)))) = int32(0)
	v327 = v291 - int32(36)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	switch v328 {
	case 0:
		goto L52
	case 1:
		goto L51
	case 2, 3:
		goto L50
	default:
		goto L49
	}
L49:
	;
	v384 = v291 + int32(44)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	if v385 != 0 {
		goto L71
	} else {
		goto L72
	}
L50:
	;
	if v286 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[781]))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v343)+16))
	if v327 != v344 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(469186), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(521469), int32(284), int32(229821))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L6
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
	F_pgaio_io_reclaim(m, v327)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L6
	} else {
		goto L63
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v343)+16)) = int32(0)
	if v286 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v350 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v350 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	F_errmsg_internal(m, int32(410034), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(521469), int32(293), int32(229821))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	goto L56
L63:
	;
	goto L49
L64:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L70
	}
L65:
	;
	v367 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v367 == int32(0) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	F_errmsg_internal(m, int32(462932), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(521469), int32(301), int32(229821))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L64
L70:
	;
	goto L49
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v384))) = int32(0)
	goto L73
L72:
	;
	goto L73
L73:
	;
	v388 = int32(4556748)
	v390 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, _consts[117])) = v390 - int32(1)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v394 == int32(0) {
		goto L39
	} else {
		goto L74
	}
L74:
	;
	if v394 != v283 {
		v291 = v394
		goto L46
	} else {
		goto L75
	}
L75:
	;
	goto L47
L76:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	if l0 != v399 {
		goto L39
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v422 = base.B2i32(base.Ui32(int32(15)) < base.Ui32(v420))
	if base.Ui32(int32(15)) < base.Ui32(v420) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	if v402 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_ReleasePredicateLocks(m, l2, int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L6
	} else {
		goto L86
	}
L83:
	;
	v405 = int32(1)
	F_LockReleaseAll(m, v405, l2^v405)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_LockReleaseAll(m, int32(2), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	goto L39
L87:
	;
	v423 = int32(0)
	goto L89
L88:
	;
	v423 = l0 + int32(292)
	goto L89
L89:
	;
	if base.Ui32(int32(15)) < base.Ui32(v420) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v425 = int32(0)
	goto L92
L91:
	;
	v425 = v420
	goto L92
L92:
	;
	if l2 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v426 = m.G0
	v428 = v426 - int32(32)
	m.G0 = v428
	v431 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v423 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v868 = m.G0
	v870 = v868 - int32(32)
	m.G0 = v870
	if v423 != 0 {
		goto L218
	} else {
		goto L219
	}
L96:
	;
	m.G0 = v428 + int32(32)
	goto L39
L97:
	;
	v434 = v425 - int32(1)
	if v434 < int32(0) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	F_hash_seq_init(m, v428+int32(12), v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L6
	} else {
		goto L157
	}
L100:
	;
	v441 = v434
	goto L101
L101:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v423+v441<<(uint(int32(2))%32))))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v462)+40))
	v465 = v463 - int32(1)
	if v465 < int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	if int32(0) < v441 {
		v441 = v441 - int32(1)
		goto L101
	} else {
		goto L156
	}
L104:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v462)+48))
	v472 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v465 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	if v463&int32(1) != 0 {
		goto L130
	} else {
		goto L131
	}
L106:
	;
	v475 = int32(-1)
	v533 = v465
	v534 = v475
	v535 = v475
	goto L105
L107:
	;
	goto L108
L108:
	;
	v480 = int32(-1)
	v486 = v465
	v487 = v480
	v488 = v480
	v495 = int32(0)
	goto L109
L109:
	;
	v505 = v486 - int32(1)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v468+v486<<(uint(int32(4))%32))))
	v510 = base.B2i32(v509 == v472)
	if v509 == v472 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v533 = v525
	v534 = v523
	v535 = v517
	goto L105
L111:
	;
	v511 = v486
	goto L113
L112:
	;
	v511 = v488
	goto L113
L113:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v468+v505<<(uint(int32(4))%32))))
	v516 = base.B2i32(v515 == v472)
	if v515 == v472 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v517 = v505
	goto L116
L115:
	;
	v517 = v511
	goto L116
L116:
	;
	if v432 == v509 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v519 = v486
	goto L119
L118:
	;
	v519 = v487
	goto L119
L119:
	;
	if v509 == v472 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v520 = v487
	goto L122
L121:
	;
	v520 = v519
	goto L122
L122:
	;
	if v432 == v515 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v522 = v505
	goto L125
L124:
	;
	v522 = v520
	goto L125
L125:
	;
	if v515 == v472 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v523 = v520
	goto L128
L127:
	;
	v523 = v522
	goto L128
L128:
	;
	v524 = int32(2)
	v525 = v486 - v524
	v527 = v495 + v524
	if v527 != v463&int32(-2) {
		v486 = v525
		v487 = v523
		v488 = v517
		v495 = v527
		goto L109
	} else {
		goto L129
	}
L129:
	;
	goto L110
L130:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v468+v533<<(uint(int32(4))%32))))
	if v554 == v432 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v560 = v534
	v561 = v535
	goto L132
L132:
	;
	if v561 < int32(0) {
		goto L103
	} else {
		goto L142
	}
L133:
	;
	v556 = v533
	goto L135
L134:
	;
	v556 = v534
	goto L135
L135:
	;
	v557 = base.B2i32(v554 == v472)
	if v554 == v472 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v558 = v534
	goto L138
L137:
	;
	v558 = v556
	goto L138
L138:
	;
	if v554 == v472 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v559 = v533
	goto L141
L140:
	;
	v559 = v535
	goto L141
L141:
	;
	v560 = v558
	v561 = v559
	goto L132
L142:
	;
	v567 = v468 + v561<<(uint(int32(4))%32)
	if v560 < int32(0) {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	v607 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_ResourceOwnerForgetLock(m, v607, v462)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L6
	} else {
		goto L155
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = v432
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+18)))
	if base.Ui32(v572) <= base.Ui32(int32(15)) {
		goto L148
	} else {
		goto L149
	}
L145:
	;
	goto L146
L146:
	;
	v588 = v468 + v560<<(uint(int32(4))%32) + int32(8)
	v589 = *(*int64)(unsafe.Add(mBase, uint32(v588)))
	v590 = *(*int64)(unsafe.Add(mBase, uint32(v567)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v588))) = v589 + v590
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v462)+40))
	v595 = v593 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v462)+40)) = v595
	if v595 <= v561 {
		goto L143
	} else {
		goto L154
	}
L147:
	;
	goto L143
L148:
	;
	if v572 != int32(15) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432+v572<<(uint(int32(2))%32))+292)) = v462
	goto L153
L152:
	;
	goto L153
L153:
	;
	v582 = v572 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+18)) = uint8(v582)
	goto L150
L154:
	;
	v600 = v468 + v595<<(uint(int32(4))%32)
	v601 = *(*int64)(unsafe.Add(mBase, uint32(v600)))
	*(*int64)(unsafe.Add(mBase, uint32(v567))) = v601
	v603 = *(*int64)(unsafe.Add(mBase, uint32(v600)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v567)+8)) = v603
	goto L143
L155:
	;
	goto L103
L156:
	;
	goto L102
L157:
	;
	v644 = F_hash_seq_search(m, v428+int32(12))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L158
	}
L158:
	;
	if v644 == int32(0) {
		goto L96
	} else {
		goto L159
	}
L159:
	;
	v648 = v644
	goto L160
L160:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v648)+40))
	v672 = v670 - int32(1)
	if v672 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	goto L96
L162:
	;
	v841 = F_hash_seq_search(m, v428+int32(12))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L6
	} else {
		goto L215
	}
L163:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v648)+48))
	v679 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	if v672 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v670&int32(1) != 0 {
		goto L189
	} else {
		goto L190
	}
L165:
	;
	v682 = int32(-1)
	v740 = v672
	v741 = v682
	v742 = v682
	goto L164
L166:
	;
	goto L167
L167:
	;
	v687 = int32(-1)
	v693 = v672
	v694 = v687
	v695 = v687
	v702 = int32(0)
	goto L168
L168:
	;
	v712 = v693 - int32(1)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v675+v693<<(uint(int32(4))%32))))
	v717 = base.B2i32(v716 == v679)
	if v716 == v679 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v740 = v732
	v741 = v730
	v742 = v724
	goto L164
L170:
	;
	v718 = v693
	goto L172
L171:
	;
	v718 = v695
	goto L172
L172:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v675+v712<<(uint(int32(4))%32))))
	v723 = base.B2i32(v722 == v679)
	if v722 == v679 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v724 = v712
	goto L175
L174:
	;
	v724 = v718
	goto L175
L175:
	;
	if v432 == v716 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v726 = v693
	goto L178
L177:
	;
	v726 = v694
	goto L178
L178:
	;
	if v716 == v679 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v727 = v694
	goto L181
L180:
	;
	v727 = v726
	goto L181
L181:
	;
	if v432 == v722 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v729 = v712
	goto L184
L183:
	;
	v729 = v727
	goto L184
L184:
	;
	if v722 == v679 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v730 = v727
	goto L187
L186:
	;
	v730 = v729
	goto L187
L187:
	;
	v731 = int32(2)
	v732 = v693 - v731
	v734 = v702 + v731
	if v734 != v670&int32(-2) {
		v693 = v732
		v694 = v730
		v695 = v724
		v702 = v734
		goto L168
	} else {
		goto L188
	}
L188:
	;
	goto L169
L189:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v675+v740<<(uint(int32(4))%32))))
	if v761 == v432 {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v767 = v741
	v768 = v742
	goto L191
L191:
	;
	if v768 < int32(0) {
		goto L162
	} else {
		goto L201
	}
L192:
	;
	v763 = v740
	goto L194
L193:
	;
	v763 = v741
	goto L194
L194:
	;
	v764 = base.B2i32(v761 == v679)
	if v761 == v679 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v765 = v741
	goto L197
L196:
	;
	v765 = v763
	goto L197
L197:
	;
	if v761 == v679 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v766 = v740
	goto L200
L199:
	;
	v766 = v742
	goto L200
L200:
	;
	v767 = v765
	v768 = v766
	goto L191
L201:
	;
	v774 = v675 + v768<<(uint(int32(4))%32)
	if v767 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _consts[181]))
	F_ResourceOwnerForgetLock(m, v814, v648)
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L6
	} else {
		goto L214
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v774))) = v432
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+18)))
	if base.Ui32(v779) <= base.Ui32(int32(15)) {
		goto L207
	} else {
		goto L208
	}
L204:
	;
	goto L205
L205:
	;
	v795 = v675 + v767<<(uint(int32(4))%32) + int32(8)
	v796 = *(*int64)(unsafe.Add(mBase, uint32(v795)))
	v797 = *(*int64)(unsafe.Add(mBase, uint32(v774)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v795))) = v796 + v797
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v648)+40))
	v802 = v800 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v648)+40)) = v802
	if v802 <= v768 {
		goto L202
	} else {
		goto L213
	}
L206:
	;
	goto L202
L207:
	;
	if v779 != int32(15) {
		goto L210
	} else {
		goto L211
	}
L208:
	;
	goto L209
L209:
	;
	goto L206
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432+v779<<(uint(int32(2))%32))+292)) = v648
	goto L212
L211:
	;
	goto L212
L212:
	;
	v789 = v779 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v432)+18)) = uint8(v789)
	goto L209
L213:
	;
	v807 = v675 + v802<<(uint(int32(4))%32)
	v808 = *(*int64)(unsafe.Add(mBase, uint32(v807)))
	*(*int64)(unsafe.Add(mBase, uint32(v774))) = v808
	v810 = *(*int64)(unsafe.Add(mBase, uint32(v807)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v774)+8)) = v810
	goto L202
L214:
	;
	goto L162
L215:
	;
	if v841 != 0 {
		v648 = v841
		goto L160
	} else {
		goto L216
	}
L216:
	;
	goto L161
L217:
	;
	m.G0 = v870 + int32(32)
	goto L39
L218:
	;
	v873 = v425 - int32(1)
	if v873 < int32(0) {
		goto L217
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	v910 = *(*int32)(unsafe.Add(mBase, _consts[1209]))
	F_hash_seq_init(m, v870+int32(12), v910)
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L6
	} else {
		goto L226
	}
L221:
	;
	v876 = v873
	goto L222
L222:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v423+v876<<(uint(int32(2))%32))))
	F_ReleaseLockIfHeld(m, v901, int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L6
	} else {
		goto L224
	}
L223:
	;
	goto L217
L224:
	;
	if v876 != 0 {
		v876 = v876 - int32(1)
		goto L222
	} else {
		goto L225
	}
L225:
	;
	goto L223
L226:
	;
	v915 = F_hash_seq_search(m, v870+int32(12))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	if v915 == int32(0) {
		goto L217
	} else {
		goto L228
	}
L228:
	;
	v923 = v915
	goto L229
L229:
	;
	F_ReleaseLockIfHeld(m, v923, int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L6
	} else {
		goto L231
	}
L230:
	;
	goto L217
L231:
	;
	v946 = F_hash_seq_search(m, v870+int32(12))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L6
	} else {
		goto L232
	}
L232:
	;
	if v946 != 0 {
		v923 = v946
		goto L229
	} else {
		goto L233
	}
L233:
	;
	goto L230
L234:
	;
	goto L39
L235:
	;
	v1004 = v999
	goto L238
L236:
	;
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, _consts[181])) = v271
	return
L238:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	m.T0[v1024].(func(*base.Module, int32, int32, int32, int32))(m, l1, l2, l3, v1023)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L6
	} else {
		goto L240
	}
L239:
	;
	goto L237
L240:
	;
	if v1022 != 0 {
		v1004 = v1022
		goto L238
	} else {
		goto L241
	}
L241:
	;
	goto L239
}
