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
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerCreate[0]))
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
	var v53 int32
	_ = v53
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
	var v163 int32
	_ = v163
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
					v163 = int32(0)
					for {
						v172 = l0 + int32(24) + v163<<(uint(int32(3))%32)
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
						v228 = v163 + v224
						v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
						if base.Ui32(v228) < base.Ui32(v229) {
							v163 = v228
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
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerEnlarge[0]))
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
							v163 = int32(0)
							for {
								v172 = l0 + int32(24) + v163<<(uint(int32(3))%32)
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
								v228 = v163 + v224
								v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
								if base.Ui32(v228) < base.Ui32(v229) {
									v163 = v228
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
							v53 = int32(0)
							for {
								v62 = v22 + v53<<(uint(int32(3))%32)
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
								v129 = v53 + int32(1)
								if v129 != v25 {
									v53 = v129
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
								v163 = int32(0)
								for {
									v172 = l0 + int32(24) + v163<<(uint(int32(3))%32)
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
									v228 = v163 + v224
									v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
									if base.Ui32(v228) < base.Ui32(v229) {
										v163 = v228
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
			F_errmsg_internal(m, int32(_a_F_ResourceOwnerEnlarge_0), int32(0))
			mBase = m.M
			v262 = m.ExcPending
			if v262 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ResourceOwnerEnlarge_1), int32(459), int32(_a_F_ResourceOwnerEnlarge_2))
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v16+v12<<(uint(int32(2))%32)-int32(4))))
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
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerForgetLock_0), v10)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerForgetLock_1), int32(1100), int32(_a_F_ResourceOwnerForgetLock_2))
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
	var v91 int32
	_ = v91
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
	var v162 int32
	_ = v162
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
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v439 int32
	_ = v439
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
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
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int64
	_ = v606
	var v607 int64
	_ = v607
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v618 int64
	_ = v618
	var v620 int64
	_ = v620
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v810 int32
	_ = v810
	var v815 int32
	_ = v815
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int64
	_ = v830
	var v831 int64
	_ = v831
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int64
	_ = v842
	var v844 int64
	_ = v844
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v1007 int32
	_ = v1007
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
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
	v270 = int32(_a_F_ResourceOwnerReleaseInternal_0)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0])) = l0
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
	F_pg_qsort(m, v241, v223, int32(8), int32(1817))
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
	v91 = int32(0)
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	v110 = v107 + v91<<(uint(int32(3))%32)
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
	if v89 != v91 {
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
	v125 = v91 + int32(1)
	if base.Ui32(v125) < base.Ui32(v123) {
		v89 = v122
		v90 = v123
		v91 = v125
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
	v162 = int32(0)
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
	v185 = *(*int64)(unsafe.Add(mBase, uint32(l0+int32(24)+v162<<(uint(v179)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v178+v160<<(uint(v179)%32)))) = v185
	v187 = int32(1)
	v188 = v160 + v187
	v190 = v162 + v187
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if base.Ui32(v190) < base.Ui32(v191) {
		v160 = v188
		v162 = v190
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
	v1031 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[1]))
	if v1031 != 0 {
		goto L233
	} else {
		goto L234
	}
L40:
	;
	F_ResourceOwnerReleaseAll(m, l0, int32(3), l2)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L6
	} else {
		goto L232
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
	v309 = int32(_a_F_ResourceOwnerReleaseInternal_1)
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2])) = v311 + int32(1)
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
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v291)+44))
	if v383 != 0 {
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
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[3]))
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
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_2), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(284), int32(_a_F_ResourceOwnerReleaseInternal_4))
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
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_5), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(293), int32(_a_F_ResourceOwnerReleaseInternal_4))
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
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_6), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(301), int32(_a_F_ResourceOwnerReleaseInternal_4))
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
	*(*int32)(unsafe.Add(mBase, uint32(v291)+44)) = int32(0)
	goto L73
L72:
	;
	goto L73
L73:
	;
	v386 = int32(_a_F_ResourceOwnerReleaseInternal_1)
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2])) = v388 - int32(1)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+356))
	if v392 == int32(0) {
		goto L39
	} else {
		goto L74
	}
L74:
	;
	if v392 != v283 {
		v291 = v392
		goto L46
	} else {
		goto L75
	}
L75:
	;
	goto L47
L76:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[4]))
	if l0 != v397 {
		goto L39
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v420 = base.B2i32(base.Ui32(int32(15)) < base.Ui32(v418))
	if base.Ui32(int32(15)) < base.Ui32(v418) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[5]))
	if v400 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
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
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L86
	}
L83:
	;
	v403 = int32(1)
	F_LockReleaseAll(m, v403, l2^v403)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_LockReleaseAll(m, int32(2), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
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
	v421 = int32(0)
	goto L89
L88:
	;
	v421 = l0 + int32(292)
	goto L89
L89:
	;
	if base.Ui32(int32(15)) < base.Ui32(v418) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v423 = int32(0)
	goto L92
L91:
	;
	v423 = v418
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
	v424 = m.G0
	v426 = v424 - int32(32)
	m.G0 = v426
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v421 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v902 = m.G0
	v904 = v902 - int32(32)
	m.G0 = v904
	if v421 != 0 {
		goto L216
	} else {
		goto L217
	}
L96:
	;
	m.G0 = v426 + int32(32)
	goto L39
L97:
	;
	v432 = v423 - int32(1)
	if v432 < int32(0) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v654 = v426 + int32(12)
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[6]))
	F_hash_seq_init(m, v654, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L6
	} else {
		goto L156
	}
L100:
	;
	v439 = v432
	goto L101
L101:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v421+v439<<(uint(int32(2))%32))))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)+40))
	v463 = v461 - int32(1)
	if v463 < int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	if int32(0) < v439 {
		v439 = v439 - int32(1)
		goto L101
	} else {
		goto L155
	}
L104:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v460)+48))
	v468 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	if v463 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if v566 < int32(0) {
		goto L103
	} else {
		goto L141
	}
L106:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v466+v533<<(uint(int32(4))%32))))
	if v554 == v430 {
		goto L132
	} else {
		goto L133
	}
L107:
	;
	v471 = int32(-1)
	v533 = v463
	v534 = v471
	v535 = v471
	goto L106
L108:
	;
	goto L109
L109:
	;
	v478 = int32(-1)
	v484 = v463
	v485 = v478
	v486 = v478
	v491 = int32(0)
	goto L110
L110:
	;
	v503 = v484 - int32(1)
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v466+v484<<(uint(int32(4))%32))))
	v508 = base.B2i32(v507 == v468)
	if v507 == v468 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v461&int32(1) == int32(0) {
		v565 = v521
		v566 = v515
		goto L105
	} else {
		goto L131
	}
L112:
	;
	v509 = v484
	goto L114
L113:
	;
	v509 = v486
	goto L114
L114:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v466+v503<<(uint(int32(4))%32))))
	v514 = base.B2i32(v513 == v468)
	if v513 == v468 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v515 = v503
	goto L117
L116:
	;
	v515 = v509
	goto L117
L117:
	;
	if v430 == v507 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v517 = v484
	goto L120
L119:
	;
	v517 = v485
	goto L120
L120:
	;
	if v507 == v468 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v518 = v485
	goto L123
L122:
	;
	v518 = v517
	goto L123
L123:
	;
	if v430 == v513 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v520 = v503
	goto L126
L125:
	;
	v520 = v518
	goto L126
L126:
	;
	if v513 == v468 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v521 = v518
	goto L129
L128:
	;
	v521 = v520
	goto L129
L129:
	;
	v522 = int32(2)
	v523 = v484 - v522
	v525 = v491 + v522
	if v525 != v461&int32(-2) {
		v484 = v523
		v485 = v521
		v486 = v515
		v491 = v525
		goto L110
	} else {
		goto L130
	}
L130:
	;
	goto L111
L131:
	;
	v533 = v523
	v534 = v521
	v535 = v515
	goto L106
L132:
	;
	v556 = v533
	goto L134
L133:
	;
	v556 = v534
	goto L134
L134:
	;
	v557 = base.B2i32(v554 == v468)
	if v554 == v468 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v558 = v534
	goto L137
L136:
	;
	v558 = v556
	goto L137
L137:
	;
	if v554 == v468 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v559 = v533
	goto L140
L139:
	;
	v559 = v535
	goto L140
L140:
	;
	v565 = v558
	v566 = v559
	goto L105
L141:
	;
	v586 = v466 + v566<<(uint(int32(4))%32)
	if v565 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v624 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	F_ResourceOwnerForgetLock(m, v624, v460)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L6
	} else {
		goto L154
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v586))) = v430
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+18)))
	if base.Ui32(v591) <= base.Ui32(int32(15)) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	v605 = v466 + v565<<(uint(int32(4))%32)
	v606 = *(*int64)(unsafe.Add(mBase, uint32(v605)+8))
	v607 = *(*int64)(unsafe.Add(mBase, uint32(v586)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v605)+8)) = v606 + v607
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v460)+40))
	v612 = v610 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v460)+40)) = v612
	if v612 <= v566 {
		goto L142
	} else {
		goto L153
	}
L146:
	;
	goto L142
L147:
	;
	if v591 != int32(15) {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	goto L149
L149:
	;
	goto L146
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430+v591<<(uint(int32(2))%32))+292)) = v460
	goto L152
L151:
	;
	goto L152
L152:
	;
	v601 = v591 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+18)) = uint8(v601)
	goto L149
L153:
	;
	v617 = v466 + v612<<(uint(int32(4))%32)
	v618 = *(*int64)(unsafe.Add(mBase, uint32(v617)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v586)+8)) = v618
	v620 = *(*int64)(unsafe.Add(mBase, uint32(v617)))
	*(*int64)(unsafe.Add(mBase, uint32(v586))) = v620
	goto L142
L154:
	;
	goto L103
L155:
	;
	goto L102
L156:
	;
	v659 = F_hash_seq_search(m, v654)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	if v659 == int32(0) {
		goto L96
	} else {
		goto L158
	}
L158:
	;
	v663 = v659
	goto L159
L159:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v663)+40))
	v687 = v685 - int32(1)
	if v687 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L96
L161:
	;
	v875 = F_hash_seq_search(m, v426+int32(12))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L6
	} else {
		goto L213
	}
L162:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v663)+48))
	v692 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	if v687 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if v790 < int32(0) {
		goto L161
	} else {
		goto L199
	}
L164:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v690+v757<<(uint(int32(4))%32))))
	if v778 == v430 {
		goto L190
	} else {
		goto L191
	}
L165:
	;
	v695 = int32(-1)
	v757 = v687
	v758 = v695
	v759 = v695
	goto L164
L166:
	;
	goto L167
L167:
	;
	v702 = int32(-1)
	v708 = v687
	v709 = v702
	v710 = v702
	v715 = int32(0)
	goto L168
L168:
	;
	v727 = v708 - int32(1)
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v690+v708<<(uint(int32(4))%32))))
	v732 = base.B2i32(v731 == v692)
	if v731 == v692 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v685&int32(1) == int32(0) {
		v789 = v745
		v790 = v739
		goto L163
	} else {
		goto L189
	}
L170:
	;
	v733 = v708
	goto L172
L171:
	;
	v733 = v710
	goto L172
L172:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v690+v727<<(uint(int32(4))%32))))
	v738 = base.B2i32(v737 == v692)
	if v737 == v692 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v739 = v727
	goto L175
L174:
	;
	v739 = v733
	goto L175
L175:
	;
	if v430 == v731 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v741 = v708
	goto L178
L177:
	;
	v741 = v709
	goto L178
L178:
	;
	if v731 == v692 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v742 = v709
	goto L181
L180:
	;
	v742 = v741
	goto L181
L181:
	;
	if v430 == v737 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v744 = v727
	goto L184
L183:
	;
	v744 = v742
	goto L184
L184:
	;
	if v737 == v692 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v745 = v742
	goto L187
L186:
	;
	v745 = v744
	goto L187
L187:
	;
	v746 = int32(2)
	v747 = v708 - v746
	v749 = v715 + v746
	if v749 != v685&int32(-2) {
		v708 = v747
		v709 = v745
		v710 = v739
		v715 = v749
		goto L168
	} else {
		goto L188
	}
L188:
	;
	goto L169
L189:
	;
	v757 = v747
	v758 = v745
	v759 = v739
	goto L164
L190:
	;
	v780 = v757
	goto L192
L191:
	;
	v780 = v758
	goto L192
L192:
	;
	v781 = base.B2i32(v778 == v692)
	if v778 == v692 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v782 = v758
	goto L195
L194:
	;
	v782 = v780
	goto L195
L195:
	;
	if v778 == v692 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v783 = v757
	goto L198
L197:
	;
	v783 = v759
	goto L198
L198:
	;
	v789 = v782
	v790 = v783
	goto L163
L199:
	;
	v810 = v690 + v790<<(uint(int32(4))%32)
	if v789 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v848 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	F_ResourceOwnerForgetLock(m, v848, v663)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L6
	} else {
		goto L212
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v810))) = v430
	v815 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v430)+18)))
	if base.Ui32(v815) <= base.Ui32(int32(15)) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	goto L203
L203:
	;
	v829 = v690 + v789<<(uint(int32(4))%32)
	v830 = *(*int64)(unsafe.Add(mBase, uint32(v829)+8))
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v810)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v829)+8)) = v830 + v831
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v663)+40))
	v836 = v834 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v663)+40)) = v836
	if v836 <= v790 {
		goto L200
	} else {
		goto L211
	}
L204:
	;
	goto L200
L205:
	;
	if v815 != int32(15) {
		goto L208
	} else {
		goto L209
	}
L206:
	;
	goto L207
L207:
	;
	goto L204
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v430+v815<<(uint(int32(2))%32))+292)) = v663
	goto L210
L209:
	;
	goto L210
L210:
	;
	v825 = v815 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v430)+18)) = uint8(v825)
	goto L207
L211:
	;
	v841 = v690 + v836<<(uint(int32(4))%32)
	v842 = *(*int64)(unsafe.Add(mBase, uint32(v841)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v810)+8)) = v842
	v844 = *(*int64)(unsafe.Add(mBase, uint32(v841)))
	*(*int64)(unsafe.Add(mBase, uint32(v810))) = v844
	goto L200
L212:
	;
	goto L161
L213:
	;
	if v875 != 0 {
		v663 = v875
		goto L159
	} else {
		goto L214
	}
L214:
	;
	goto L160
L215:
	;
	m.G0 = v904 + int32(32)
	goto L39
L216:
	;
	v907 = v423 - int32(1)
	if v907 < int32(0) {
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v942 = v904 + int32(12)
	v944 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[6]))
	F_hash_seq_init(m, v942, v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L6
	} else {
		goto L224
	}
L219:
	;
	v910 = v907
	goto L220
L220:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v421+v910<<(uint(int32(2))%32))))
	F_ReleaseLockIfHeld(m, v935, int32(0))
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L6
	} else {
		goto L222
	}
L221:
	;
	goto L215
L222:
	;
	if v910 != 0 {
		v910 = v910 - int32(1)
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v947 = F_hash_seq_search(m, v942)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	if v947 == int32(0) {
		goto L215
	} else {
		goto L226
	}
L226:
	;
	v955 = v947
	goto L227
L227:
	;
	F_ReleaseLockIfHeld(m, v955, int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L6
	} else {
		goto L229
	}
L228:
	;
	goto L215
L229:
	;
	v978 = F_hash_seq_search(m, v904+int32(12))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	if v978 != 0 {
		v955 = v978
		goto L227
	} else {
		goto L231
	}
L231:
	;
	goto L228
L232:
	;
	goto L39
L233:
	;
	v1036 = v1031
	goto L236
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0])) = v271
	return
L236:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1036)))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+8))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	m.T0[v1056].(func(*base.Module, int32, int32, int32, int32))(m, l1, l2, l3, v1055)
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L6
	} else {
		goto L238
	}
L237:
	;
	goto L235
L238:
	;
	if v1054 != 0 {
		v1036 = v1054
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
}
