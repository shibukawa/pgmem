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
	v7 = F_MemoryContextAllocZero(m, v5, int32(616))
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
		v17 = v7 + int32(608)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+612)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+608)) = v17
		return v7
	}
}
func F_ResourceOwnerEnlarge(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v72 int64
	_ = v72
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v132 int32
	_ = v132
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v191 int64
	_ = v191
	var v194 int64
	_ = v194
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v13 != int32(1) {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
		if base.Ui32(int32(32)) <= base.Ui32(v16) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+544))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v20+v16) < base.Ui32(v19) {
				v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
				if v160 != 0 {
					v167 = int32(0)
					for {
						v178 = l0 + int32(24) + v167<<(uint(int32(4))%32)
						v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
						v181 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
						v182 = int64(33)
						v186 = (int64(base.Ui64(v181)>>(uint(v182)%64)) ^ v181) * int64(-49064778989728563)
						v191 = (int64(base.Ui64(v186)>>(uint(v182)%64)) ^ v186) * int64(-4265267296055464877)
						v194 = int64(base.Ui64(v191)>>(uint(v182)%64)) ^ v191
						v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
						v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
						v207 = base.I32_wrap_i64(base.I64_extend_i32_u(v179) + int64(base.Ui64(v194)>>(uint(int64(7))%64)) + int64(367372515) ^ v194)
						for {
							v218 = v207 & (v202 - int32(1))
							v222 = v218 << (uint(int32(4)) % 32)
							v223 = v205 + v222
							v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
							if v224 != 0 {
								v207 = v218 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						*(*int64)(unsafe.Add(mBase, uint32(v223))) = v181
						v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
						*(*int32)(unsafe.Add(mBase, uint32(v226+v222)+8)) = v179
						v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v230 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v229 + v230
						v234 = v167 + v230
						v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
						if base.Ui32(v234) < base.Ui32(v235) {
							v167 = v234
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v249 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v249)
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerEnlarge[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
				if v26 != 0 {
					v30 = v26 << (uint(int32(1)) % 32)
				} else {
					v30 = int32(64)
				}
				v33 = F_MemoryContextAllocZero(m, v25, v30<<(uint(int32(4))%32))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+540)) = v30
					*(*int32)(unsafe.Add(mBase, uint32(l0)+536)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
					v40 = v30 - int32(32)
					v44 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) * int32(3)
					if base.Ui32(v40) < base.Ui32(v44) {
						v46 = v40
					} else {
						v46 = v44
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+544)) = v46
					if v23 == int32(0) {
						v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
						if v160 != 0 {
							v167 = int32(0)
							for {
								v178 = l0 + int32(24) + v167<<(uint(int32(4))%32)
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
								v181 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
								v182 = int64(33)
								v186 = (int64(base.Ui64(v181)>>(uint(v182)%64)) ^ v181) * int64(-49064778989728563)
								v191 = (int64(base.Ui64(v186)>>(uint(v182)%64)) ^ v186) * int64(-4265267296055464877)
								v194 = int64(base.Ui64(v191)>>(uint(v182)%64)) ^ v191
								v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
								v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
								v207 = base.I32_wrap_i64(base.I64_extend_i32_u(v179) + int64(base.Ui64(v194)>>(uint(int64(7))%64)) + int64(367372515) ^ v194)
								for {
									v218 = v207 & (v202 - int32(1))
									v222 = v218 << (uint(int32(4)) % 32)
									v223 = v205 + v222
									v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
									if v224 != 0 {
										v207 = v218 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								*(*int64)(unsafe.Add(mBase, uint32(v223))) = v181
								v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
								*(*int32)(unsafe.Add(mBase, uint32(v226+v222)+8)) = v179
								v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v230 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v229 + v230
								v234 = v167 + v230
								v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
								if base.Ui32(v234) < base.Ui32(v235) {
									v167 = v234
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v249 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v249)
						return
					} else {
						if v26 != 0 {
							v53 = int32(0)
							for {
								v64 = v23 + v53<<(uint(int32(4))%32)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
								if v65 != 0 {
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
									v68 = int64(33)
									v72 = (int64(base.Ui64(v67)>>(uint(v68)%64)) ^ v67) * int64(-49064778989728563)
									v77 = (int64(base.Ui64(v72)>>(uint(v68)%64)) ^ v72) * int64(-4265267296055464877)
									v80 = int64(base.Ui64(v77)>>(uint(v68)%64)) ^ v77
									v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
									v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
									v93 = base.I32_wrap_i64(base.I64_extend_i32_u(v65) + int64(base.Ui64(v80)>>(uint(int64(7))%64)) + int64(367372515) ^ v80)
									for {
										v104 = v93 & (v88 - int32(1))
										v108 = v104 << (uint(int32(4)) % 32)
										v109 = v91 + v108
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
										if v110 != 0 {
											v93 = v104 + int32(1)
											continue
										} else {
											break
										}
										break
									}
									*(*int64)(unsafe.Add(mBase, uint32(v109))) = v67
									v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
									*(*int32)(unsafe.Add(mBase, uint32(v112+v108)+8)) = v65
									v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v115 + int32(1)
								} else {
								}
								v132 = v53 + int32(1)
								if v132 != v26 {
									v53 = v132
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						F_pfree(m, v23)
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return
						} else {
							v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
							if v160 != 0 {
								v167 = int32(0)
								for {
									v178 = l0 + int32(24) + v167<<(uint(int32(4))%32)
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
									v181 = *(*int64)(unsafe.Add(mBase, uint32(v178)))
									v182 = int64(33)
									v186 = (int64(base.Ui64(v181)>>(uint(v182)%64)) ^ v181) * int64(-49064778989728563)
									v191 = (int64(base.Ui64(v186)>>(uint(v182)%64)) ^ v186) * int64(-4265267296055464877)
									v194 = int64(base.Ui64(v191)>>(uint(v182)%64)) ^ v191
									v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
									v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
									v207 = base.I32_wrap_i64(base.I64_extend_i32_u(v179) + int64(base.Ui64(v194)>>(uint(int64(7))%64)) + int64(367372515) ^ v194)
									for {
										v218 = v207 & (v202 - int32(1))
										v222 = v218 << (uint(int32(4)) % 32)
										v223 = v205 + v222
										v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
										if v224 != 0 {
											v207 = v218 + int32(1)
											continue
										} else {
											break
										}
										break
									}
									*(*int64)(unsafe.Add(mBase, uint32(v223))) = v181
									v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
									*(*int32)(unsafe.Add(mBase, uint32(v226+v222)+8)) = v179
									v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v230 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v229 + v230
									v234 = v167 + v230
									v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
									if base.Ui32(v234) < base.Ui32(v235) {
										v167 = v234
										continue
									} else {
										break
									}
									break
								}
							} else {
							}
							v249 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v249)
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
		v266 = m.ExcPending
		if v266 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ResourceOwnerEnlarge_0), int32(0))
			mBase = m.M
			v270 = m.ExcPending
			if v270 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ResourceOwnerEnlarge_1), int32(466), int32(_a_F_ResourceOwnerEnlarge_2))
				mBase = m.M
				v275 = m.ExcPending
				if v275 != 0 {
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
	v16 = l0 + int32(548)
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
	F_errfinish(m, int32(_a_F_ResourceOwnerForgetLock_1), int32(1107), int32(_a_F_ResourceOwnerForgetLock_2))
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
	var v115 int32
	_ = v115
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
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
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v610 int64
	_ = v610
	var v611 int64
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v624 int64
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v719 int32
	_ = v719
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v834 int64
	_ = v834
	var v835 int64
	_ = v835
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v846 int64
	_ = v846
	var v848 int64
	_ = v848
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v1011 int32
	_ = v1011
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
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
	v274 = int32(_a_F_ResourceOwnerReleaseInternal_0)
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
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
	F_pg_qsort(m, v245, v227, int32(16), int32(2040))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L38
	}
L16:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if v154 != 0 {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
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
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v227 = v129
	v245 = l0 + int32(24)
	goto L15
L20:
	;
	v136 = int32(0)
	goto L16
L21:
	;
	goto L22
L22:
	;
	v89 = int32(0)
	v90 = int32(0)
	v91 = v80
	goto L23
L23:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v110 = v107 + v90<<(uint(int32(4))%32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+8))
	if v111 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v136 = v124
	goto L16
L25:
	;
	if v89 != v90 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v124 = v89
	v125 = v91
	goto L27
L27:
	;
	v127 = v90 + int32(1)
	if base.Ui32(v127) < base.Ui32(v125) {
		v89 = v124
		v90 = v127
		v91 = v125
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v115 = v107 + v89<<(uint(int32(4))%32)
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v110)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+8)) = v116
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v110)))
	*(*int64)(unsafe.Add(mBase, uint32(v115))) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+540))
	v121 = v120
	goto L30
L29:
	;
	v121 = v91
	goto L30
L30:
	;
	v124 = v89 + int32(1)
	v125 = v121
	goto L27
L31:
	;
	goto L24
L32:
	;
	v162 = v136
	v163 = int32(0)
	goto L35
L33:
	;
	v201 = v136
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v201
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)) = uint8(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v227 = v201
	v245 = v222
	goto L15
L35:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+536))
	v181 = int32(4)
	v183 = v180 + v162<<(uint(v181)%32)
	v186 = l0 + int32(24) + v163<<(uint(v181)%32)
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v183)+8)) = v187
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
	*(*int64)(unsafe.Add(mBase, uint32(v183))) = v189
	v191 = int32(1)
	v192 = v162 + v191
	v194 = v163 + v191
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	if base.Ui32(v194) < base.Ui32(v195) {
		v162 = v192
		v163 = v194
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v201 = v192
	goto L34
L37:
	;
	goto L36
L38:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v250)
	goto L14
L39:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[1]))
	if v1035 != 0 {
		goto L233
	} else {
		goto L234
	}
L40:
	;
	F_ResourceOwnerReleaseAll(m, l0, int32(3), l2)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
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
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+612))
	if v283 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v287 = l0 + int32(608)
	if v283 == v287 {
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v290 = l2 ^ int32(1)
	v295 = v283
	goto L46
L46:
	;
	v313 = int32(_a_F_ResourceOwnerReleaseInternal_1)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2])) = v315 + int32(1)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+4)) = v324
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	*(*int32)(unsafe.Add(mBase, uint32(v324))) = v326
	goto L48
L47:
	;
	goto L39
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v295-int32(4)))) = int32(0)
	v331 = v295 - int32(36)
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	switch v332 {
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
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v295)+44))
	if v387 != 0 {
		goto L71
	} else {
		goto L72
	}
L50:
	;
	if v290 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[3]))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+16))
	if v331 != v348 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_2), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(284), int32(_a_F_ResourceOwnerReleaseInternal_4))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
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
	F_pgaio_io_reclaim(m, v331)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L63
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v347)+16)) = int32(0)
	if v290 != 0 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v354 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v354 == int32(0) {
		goto L56
	} else {
		goto L60
	}
L60:
	;
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_5), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(293), int32(_a_F_ResourceOwnerReleaseInternal_4))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
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
	v385 = m.ExcPending
	if v385 != 0 {
		goto L6
	} else {
		goto L70
	}
L65:
	;
	v371 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v371 == int32(0) {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	F_errmsg_internal(m, int32(_a_F_ResourceOwnerReleaseInternal_6), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_ResourceOwnerReleaseInternal_3), int32(301), int32(_a_F_ResourceOwnerReleaseInternal_4))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v295)+44)) = int32(0)
	goto L73
L72:
	;
	goto L73
L73:
	;
	v390 = int32(_a_F_ResourceOwnerReleaseInternal_1)
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[2])) = v392 - int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+612))
	if v396 == int32(0) {
		goto L39
	} else {
		goto L74
	}
L74:
	;
	if v396 != v287 {
		v295 = v396
		goto L46
	} else {
		goto L75
	}
L75:
	;
	goto L47
L76:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[4]))
	if l0 != v401 {
		goto L39
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+18)))
	v424 = base.B2i32(base.Ui32(int32(15)) < base.Ui32(v422))
	if base.Ui32(int32(15)) < base.Ui32(v422) {
		goto L87
	} else {
		goto L88
	}
L79:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[5]))
	if v404 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_LockErrorCleanup(m)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
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
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L86
	}
L83:
	;
	v407 = int32(1)
	F_LockReleaseAll(m, v407, l2^v407)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_LockReleaseAll(m, int32(2), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
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
	v425 = int32(0)
	goto L89
L88:
	;
	v425 = l0 + int32(548)
	goto L89
L89:
	;
	if base.Ui32(int32(15)) < base.Ui32(v422) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v427 = int32(0)
	goto L92
L91:
	;
	v427 = v422
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
	v428 = m.G0
	v430 = v428 - int32(32)
	m.G0 = v430
	v433 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	if v425 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v906 = m.G0
	v908 = v906 - int32(32)
	m.G0 = v908
	if v425 != 0 {
		goto L216
	} else {
		goto L217
	}
L96:
	;
	m.G0 = v430 + int32(32)
	goto L39
L97:
	;
	v436 = v427 - int32(1)
	if v436 < int32(0) {
		goto L96
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v658 = v430 + int32(12)
	v660 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[6]))
	F_hash_seq_init(m, v658, v660)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L6
	} else {
		goto L156
	}
L100:
	;
	v443 = v436
	goto L101
L101:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v425+v443<<(uint(int32(2))%32))))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v464)+40))
	v467 = v465 - int32(1)
	if v467 < int32(0) {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	goto L96
L103:
	;
	if int32(0) < v443 {
		v443 = v443 - int32(1)
		goto L101
	} else {
		goto L155
	}
L104:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v464)+48))
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	if v467 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	if v570 < int32(0) {
		goto L103
	} else {
		goto L141
	}
L106:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v470+v537<<(uint(int32(4))%32))))
	if v558 == v434 {
		goto L132
	} else {
		goto L133
	}
L107:
	;
	v475 = int32(-1)
	v537 = v467
	v538 = v475
	v539 = v475
	goto L106
L108:
	;
	goto L109
L109:
	;
	v482 = int32(-1)
	v488 = v467
	v489 = v482
	v490 = v482
	v495 = int32(0)
	goto L110
L110:
	;
	v507 = v488 - int32(1)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v470+v488<<(uint(int32(4))%32))))
	v512 = base.B2i32(v511 == v472)
	if v511 == v472 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v465&int32(1) == int32(0) {
		v569 = v525
		v570 = v519
		goto L105
	} else {
		goto L131
	}
L112:
	;
	v513 = v488
	goto L114
L113:
	;
	v513 = v490
	goto L114
L114:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v470+v507<<(uint(int32(4))%32))))
	v518 = base.B2i32(v517 == v472)
	if v517 == v472 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v519 = v507
	goto L117
L116:
	;
	v519 = v513
	goto L117
L117:
	;
	if v434 == v511 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v521 = v488
	goto L120
L119:
	;
	v521 = v489
	goto L120
L120:
	;
	if v511 == v472 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v522 = v489
	goto L123
L122:
	;
	v522 = v521
	goto L123
L123:
	;
	if v434 == v517 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v524 = v507
	goto L126
L125:
	;
	v524 = v522
	goto L126
L126:
	;
	if v517 == v472 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v525 = v522
	goto L129
L128:
	;
	v525 = v524
	goto L129
L129:
	;
	v526 = int32(2)
	v527 = v488 - v526
	v529 = v495 + v526
	if v529 != v465&int32(-2) {
		v488 = v527
		v489 = v525
		v490 = v519
		v495 = v529
		goto L110
	} else {
		goto L130
	}
L130:
	;
	goto L111
L131:
	;
	v537 = v527
	v538 = v525
	v539 = v519
	goto L106
L132:
	;
	v560 = v537
	goto L134
L133:
	;
	v560 = v538
	goto L134
L134:
	;
	v561 = base.B2i32(v558 == v472)
	if v558 == v472 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v562 = v538
	goto L137
L136:
	;
	v562 = v560
	goto L137
L137:
	;
	if v558 == v472 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v563 = v537
	goto L140
L139:
	;
	v563 = v539
	goto L140
L140:
	;
	v569 = v562
	v570 = v563
	goto L105
L141:
	;
	v590 = v470 + v570<<(uint(int32(4))%32)
	if v569 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	v628 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	F_ResourceOwnerForgetLock(m, v628, v464)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L6
	} else {
		goto L154
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = v434
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+18)))
	if base.Ui32(v595) <= base.Ui32(int32(15)) {
		goto L147
	} else {
		goto L148
	}
L144:
	;
	goto L145
L145:
	;
	v609 = v470 + v569<<(uint(int32(4))%32)
	v610 = *(*int64)(unsafe.Add(mBase, uint32(v609)+8))
	v611 = *(*int64)(unsafe.Add(mBase, uint32(v590)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v609)+8)) = v610 + v611
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v464)+40))
	v616 = v614 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v464)+40)) = v616
	if v616 <= v570 {
		goto L142
	} else {
		goto L153
	}
L146:
	;
	goto L142
L147:
	;
	if v595 != int32(15) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v434+v595<<(uint(int32(2))%32))+548)) = v464
	goto L152
L151:
	;
	goto L152
L152:
	;
	v605 = v595 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v434)+18)) = uint8(v605)
	goto L149
L153:
	;
	v621 = v470 + v616<<(uint(int32(4))%32)
	v622 = *(*int64)(unsafe.Add(mBase, uint32(v621)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v590)+8)) = v622
	v624 = *(*int64)(unsafe.Add(mBase, uint32(v621)))
	*(*int64)(unsafe.Add(mBase, uint32(v590))) = v624
	goto L142
L154:
	;
	goto L103
L155:
	;
	goto L102
L156:
	;
	v663 = F_hash_seq_search(m, v658)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	if v663 == int32(0) {
		goto L96
	} else {
		goto L158
	}
L158:
	;
	v667 = v663
	goto L159
L159:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v667)+40))
	v691 = v689 - int32(1)
	if v691 < int32(0) {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L96
L161:
	;
	v879 = F_hash_seq_search(m, v430+int32(12))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L6
	} else {
		goto L213
	}
L162:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v667)+48))
	v696 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	if v691 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	if v794 < int32(0) {
		goto L161
	} else {
		goto L199
	}
L164:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v694+v761<<(uint(int32(4))%32))))
	if v782 == v434 {
		goto L190
	} else {
		goto L191
	}
L165:
	;
	v699 = int32(-1)
	v761 = v691
	v762 = v699
	v763 = v699
	goto L164
L166:
	;
	goto L167
L167:
	;
	v706 = int32(-1)
	v712 = v691
	v713 = v706
	v714 = v706
	v719 = int32(0)
	goto L168
L168:
	;
	v731 = v712 - int32(1)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v694+v712<<(uint(int32(4))%32))))
	v736 = base.B2i32(v735 == v696)
	if v735 == v696 {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	if v689&int32(1) == int32(0) {
		v793 = v749
		v794 = v743
		goto L163
	} else {
		goto L189
	}
L170:
	;
	v737 = v712
	goto L172
L171:
	;
	v737 = v714
	goto L172
L172:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v694+v731<<(uint(int32(4))%32))))
	v742 = base.B2i32(v741 == v696)
	if v741 == v696 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v743 = v731
	goto L175
L174:
	;
	v743 = v737
	goto L175
L175:
	;
	if v434 == v735 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v745 = v712
	goto L178
L177:
	;
	v745 = v713
	goto L178
L178:
	;
	if v735 == v696 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v746 = v713
	goto L181
L180:
	;
	v746 = v745
	goto L181
L181:
	;
	if v434 == v741 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v748 = v731
	goto L184
L183:
	;
	v748 = v746
	goto L184
L184:
	;
	if v741 == v696 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v749 = v746
	goto L187
L186:
	;
	v749 = v748
	goto L187
L187:
	;
	v750 = int32(2)
	v751 = v712 - v750
	v753 = v719 + v750
	if v753 != v689&int32(-2) {
		v712 = v751
		v713 = v749
		v714 = v743
		v719 = v753
		goto L168
	} else {
		goto L188
	}
L188:
	;
	goto L169
L189:
	;
	v761 = v751
	v762 = v749
	v763 = v743
	goto L164
L190:
	;
	v784 = v761
	goto L192
L191:
	;
	v784 = v762
	goto L192
L192:
	;
	v785 = base.B2i32(v782 == v696)
	if v782 == v696 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v786 = v762
	goto L195
L194:
	;
	v786 = v784
	goto L195
L195:
	;
	if v782 == v696 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v787 = v761
	goto L198
L197:
	;
	v787 = v763
	goto L198
L198:
	;
	v793 = v786
	v794 = v787
	goto L163
L199:
	;
	v814 = v694 + v794<<(uint(int32(4))%32)
	if v793 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v852 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0]))
	F_ResourceOwnerForgetLock(m, v852, v667)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L6
	} else {
		goto L212
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v814))) = v434
	v819 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+18)))
	if base.Ui32(v819) <= base.Ui32(int32(15)) {
		goto L205
	} else {
		goto L206
	}
L202:
	;
	goto L203
L203:
	;
	v833 = v694 + v793<<(uint(int32(4))%32)
	v834 = *(*int64)(unsafe.Add(mBase, uint32(v833)+8))
	v835 = *(*int64)(unsafe.Add(mBase, uint32(v814)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v833)+8)) = v834 + v835
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v667)+40))
	v840 = v838 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v667)+40)) = v840
	if v840 <= v794 {
		goto L200
	} else {
		goto L211
	}
L204:
	;
	goto L200
L205:
	;
	if v819 != int32(15) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v434+v819<<(uint(int32(2))%32))+548)) = v667
	goto L210
L209:
	;
	goto L210
L210:
	;
	v829 = v819 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v434)+18)) = uint8(v829)
	goto L207
L211:
	;
	v845 = v694 + v840<<(uint(int32(4))%32)
	v846 = *(*int64)(unsafe.Add(mBase, uint32(v845)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v814)+8)) = v846
	v848 = *(*int64)(unsafe.Add(mBase, uint32(v845)))
	*(*int64)(unsafe.Add(mBase, uint32(v814))) = v848
	goto L200
L212:
	;
	goto L161
L213:
	;
	if v879 != 0 {
		v667 = v879
		goto L159
	} else {
		goto L214
	}
L214:
	;
	goto L160
L215:
	;
	m.G0 = v908 + int32(32)
	goto L39
L216:
	;
	v911 = v427 - int32(1)
	if v911 < int32(0) {
		goto L215
	} else {
		goto L219
	}
L217:
	;
	goto L218
L218:
	;
	v946 = v908 + int32(12)
	v948 = *(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[6]))
	F_hash_seq_init(m, v946, v948)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L6
	} else {
		goto L224
	}
L219:
	;
	v914 = v911
	goto L220
L220:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v425+v914<<(uint(int32(2))%32))))
	F_ReleaseLockIfHeld(m, v939, int32(0))
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L6
	} else {
		goto L222
	}
L221:
	;
	goto L215
L222:
	;
	if v914 != 0 {
		v914 = v914 - int32(1)
		goto L220
	} else {
		goto L223
	}
L223:
	;
	goto L221
L224:
	;
	v951 = F_hash_seq_search(m, v946)
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	if v951 == int32(0) {
		goto L215
	} else {
		goto L226
	}
L226:
	;
	v959 = v951
	goto L227
L227:
	;
	F_ReleaseLockIfHeld(m, v959, int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L6
	} else {
		goto L229
	}
L228:
	;
	goto L215
L229:
	;
	v982 = F_hash_seq_search(m, v908+int32(12))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L6
	} else {
		goto L230
	}
L230:
	;
	if v982 != 0 {
		v959 = v982
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
	v1040 = v1035
	goto L236
L234:
	;
	goto L235
L235:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ResourceOwnerReleaseInternal[0])) = v275
	return
L236:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1040)))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+8))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1040)+4))
	m.T0[v1060].(func(*base.Module, int32, int32, int32, int32))(m, l1, l2, l3, v1059)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L6
	} else {
		goto L238
	}
L237:
	;
	goto L235
L238:
	;
	if v1058 != 0 {
		v1040 = v1058
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
}
