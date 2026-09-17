package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplestore_ateof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2+v3*int32(24))+4)))
	return v7
}
func F_tuplestore_clear(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	v2 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 == v2 {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = v10 - v11
		if v12 < v6 {
			v14 = v6
		} else {
			v14 = v12
		}
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v14
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		if v30 != 0 {
			F_BufFileClose(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				F_MemoryContextReset(m, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v41 = F_GetMemoryChunkSpace(m, v40)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
						v45 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
						v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						if v45 < v55 {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
							v60 = v58
							v61 = v2
							for {
								v64 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
								*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
								v71 = v61 + int32(1)
								v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
								if v71 < v72 {
									v60 = v60 + int32(24)
									v61 = v71
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
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
			F_MemoryContextReset(m, v35)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v41 = F_GetMemoryChunkSpace(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
					v45 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
					*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
					v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
					if v45 < v55 {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
						v60 = v58
						v61 = v2
						for {
							v64 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
							v71 = v61 + int32(1)
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							if v71 < v72 {
								v60 = v60 + int32(24)
								v61 = v71
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
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v17 = F_BufFileSize(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 < v6 {
				v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
				v24 = v20
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v25)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v24
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				if v30 != 0 {
					F_BufFileClose(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						F_MemoryContextReset(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v41 = F_GetMemoryChunkSpace(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
								v45 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
								*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
								v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
								if v45 < v55 {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									v60 = v58
									v61 = v2
									for {
										v64 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
										*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
										v71 = v61 + int32(1)
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
										if v71 < v72 {
											v60 = v60 + int32(24)
											v61 = v71
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
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					F_MemoryContextReset(m, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
						v41 = F_GetMemoryChunkSpace(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
							v45 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
							*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
							v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							if v45 < v55 {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
								v60 = v58
								v61 = v2
								for {
									v64 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
									*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
									v71 = v61 + int32(1)
									v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
									if v71 < v72 {
										v60 = v60 + int32(24)
										v61 = v71
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
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				v22 = F_BufFileSize(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = v22
					v25 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v25)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v24
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if v30 != 0 {
						F_BufFileClose(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							F_MemoryContextReset(m, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v41 = F_GetMemoryChunkSpace(m, v40)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
									v45 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
									*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
									v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
									if v45 < v55 {
										v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
										v60 = v58
										v61 = v2
										for {
											v64 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
											*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
											v71 = v61 + int32(1)
											v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
											if v71 < v72 {
												v60 = v60 + int32(24)
												v61 = v71
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
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						F_MemoryContextReset(m, v35)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v38
							v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v41 = F_GetMemoryChunkSpace(m, v40)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v43
								v45 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)) = uint8(v45)
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v45
								*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v43
								v51 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v51 - base.I64_extend_i32_u(v41)
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
								if v45 < v55 {
									v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
									v60 = v58
									v61 = v2
									for {
										v64 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v64
										*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
										v71 = v61 + int32(1)
										v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
										if v71 < v72 {
											v60 = v60 + int32(24)
											v61 = v71
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
					}
				}
			}
		}
	}
}
func F_tuplestore_puttuple_common(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int64
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v330 int32
	_ = v330
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v388 int64
	_ = v388
	var v389 int64
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	v3 = int32(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v10 + int64(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	default:
		goto L7
	}
L1:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	m.T0[v440].(func(*base.Module, int32, int32))(m, l0, l1)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L48
	} else {
		goto L113
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v282 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v291+v282<<(uint(int32(2))%32)))) = l1
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v296 < v297 {
		goto L78
	} else {
		goto L79
	}
L3:
	;
	v280 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v280)
	v282 = v53
	goto L2
L4:
	;
	if v62 < base.I64_extend_i32_u((v235-v54)<<(uint(int32(2))%32)) {
		goto L3
	} else {
		goto L67
	}
L5:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v232)
	v235 = int32(536870911)
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L48
	} else {
		goto L63
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L48
	} else {
		goto L60
	}
L8:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v139 = v135 + v136*int32(24)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+4)))
	if v140 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L9:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v88 {
		goto L34
	} else {
		goto L35
	}
L10:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v15 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v21 = v18
	v22 = v15
	v23 = v3
	goto L14
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v53 < v54-int32(1) {
		v282 = v53
		goto L2
	} else {
		goto L20
	}
L14:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v28 != int32(1) {
		v38 = v22
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	v42 = v23 + int32(1)
	if v42 < v38 {
		v21 = v21 + int32(24)
		v22 = v38
		v23 = v42
		goto L14
	} else {
		goto L19
	}
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v23 == v31 {
		v38 = v22
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v38 = v37
	goto L16
L19:
	;
	goto L15
L20:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v58 != int32(1) {
		v282 = v53
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v63 = v61 - v62
	if v63 <= v62 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	if v54 != int32(2147483647) {
		goto L5
	} else {
		goto L33
	}
L23:
	;
	if v82 <= v54 {
		goto L3
	} else {
		goto L31
	}
L24:
	;
	if int32(1073741822) < v54 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v69)
	v75 = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v61), base.F64_convert_i64_s(v63)), base.F64_convert_i32_s(v54))
	v76 = float64(2.147483647e+09)
	if base.F64_lt(v75, v76) != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v82 = v54 << (uint(int32(1)) % 32)
	goto L23
L28:
	;
	v79 = v75
	goto L30
L29:
	;
	v79 = v76
	goto L30
L30:
	;
	v82 = base.I32_trunc_sat_f64_s(v79)
	goto L23
L31:
	;
	if base.Ui32(v82) < base.Ui32(int32(536870911)) {
		v235 = v82
		goto L4
	} else {
		goto L32
	}
L32:
	;
	goto L5
L33:
	;
	goto L3
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v94 = v91
	v95 = v88
	v96 = v3
	goto L37
L35:
	;
	goto L36
L36:
	;
	goto L1
L37:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v101 != int32(1) {
		v120 = v95
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v124 = v96 + int32(1)
	if v124 < v120 {
		v94 = v94 + int32(24)
		v95 = v120
		v96 = v124
		goto L37
	} else {
		goto L43
	}
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v96 == v104 {
		v120 = v95
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v106 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)) = uint8(v106)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v94+int32(12)))) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v108)+32))
	v116 = int64(*(*int32)(unsafe.Add(mBase, uint32(v108)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v94+int32(16)))) = v115 + v116
	goto L42
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v120 = v119
	goto L39
L43:
	;
	goto L38
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v139+int32(12)))) = v148
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v143)+32))
	v151 = int64(*(*int32)(unsafe.Add(mBase, uint32(v143)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v139+int32(16)))) = v150 + v151
	goto L47
L45:
	;
	goto L46
L46:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v156 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v158 = F_BufFileSeek(m, v154, v155, v156, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	return
L49:
	;
	if v158 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v162 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v168 = v165
	v169 = v162
	v170 = v3
	goto L54
L52:
	;
	goto L53
L53:
	;
	goto L1
L54:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+4)))
	if v175 != int32(1) {
		v187 = v169
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v191 = v170 + int32(1)
	if v191 < v187 {
		v168 = v168 + int32(24)
		v169 = v187
		v170 = v191
		goto L54
	} else {
		goto L59
	}
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v170 == v178 {
		v187 = v169
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168)+4)) = uint8(v180)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+12)) = v182
	v184 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v168)+16)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v187 = v186
	goto L56
L59:
	;
	goto L55
L60:
	;
	F_errmsg_internal(m, int32(_a_F_tuplestore_puttuple_common_0), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L48
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_tuplestore_puttuple_common_1), int32(941), int32(_a_F_tuplestore_puttuple_common_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L48
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L48
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_tuplestore_puttuple_common_3), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L48
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_tuplestore_puttuple_common_1), int32(921), int32(_a_F_tuplestore_puttuple_common_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L48
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v243 = F_GetMemoryChunkSpace(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L48
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v235
	v246 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v246 + base.I64_extend_i32_u(v243)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v253 = F_repalloc_huge(m, v250, v235<<(uint(int32(2))%32))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L48
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v253
	v256 = F_GetMemoryChunkSpace(m, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L48
	} else {
		goto L70
	}
L70:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v260 = v258 - base.I64_extend_i32_u(v256)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v260
	if int64(0) <= v260 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v282 = v264
	goto L2
L72:
	;
	goto L73
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L48
	} else {
		goto L74
	}
L74:
	;
	F_errmsg_internal(m, int32(_a_F_tuplestore_puttuple_common_4), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L48
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_tuplestore_puttuple_common_1), int32(716), int32(_a_F_tuplestore_puttuple_common_5))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L48
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	return
L78:
	;
	v299 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) <= v299 {
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L48
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v304 = int32(_a_F_tuplestore_puttuple_common_6)
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[0]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[0])) = v307
	v309 = int32(_a_F_tuplestore_puttuple_common_7)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[1]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[1])) = v313
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v316 = F_BufFileCreateTemp(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L48
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v316
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[0])) = v305
	*(*int32)(unsafe.Add(mBase, _c_F_tuplestore_puttuple_common[1])) = v310
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v327 = int32(base.Ui32(v323)>>(uint(int32(3))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v327)
	v329 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v330 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v356 = v354
	goto L97
L85:
	;
	v333 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v334 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v335 = v333 - v334
	if v335 < v329 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v339 = F_BufFileSize(m, v316)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L48
	} else {
		goto L92
	}
L88:
	;
	v337 = v329
	goto L90
L89:
	;
	v337 = v335
	goto L90
L90:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v337
	goto L84
L91:
	;
	v347 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v347)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v346
	goto L84
L92:
	;
	if v339 < v329 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v342 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v346 = v342
	goto L91
L94:
	;
	goto L95
L95:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v344 = F_BufFileSize(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L48
	} else {
		goto L96
	}
L96:
	;
	v346 = v344
	goto L91
L97:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v364 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = int64(0)
	goto L77
L99:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v371 = v367
	v372 = v364
	v373 = int32(0)
	goto L102
L100:
	;
	goto L101
L101:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v356 < v408 {
		goto L109
	} else {
		goto L110
	}
L102:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v371)+8))
	if v356 != v378 {
		v393 = v372
		goto L104
	} else {
		goto L105
	}
L103:
	;
	goto L101
L104:
	;
	v397 = v373 + int32(1)
	if v397 < v393 {
		v371 = v371 + int32(24)
		v372 = v393
		v373 = v397
		goto L102
	} else {
		goto L108
	}
L105:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v371)+4)))
	if v380 != 0 {
		v393 = v372
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v381)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v371+int32(12)))) = v386
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v381)+32))
	v389 = int64(*(*int32)(unsafe.Add(mBase, uint32(v381)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v371+int32(16)))) = v388 + v389
	goto L107
L107:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v393 = v392
	goto L104
L108:
	;
	goto L103
L109:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v410+v356<<(uint(int32(2))%32))))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	m.T0[v415].(func(*base.Module, int32, int32))(m, l0, v414)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L48
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	goto L98
L112:
	;
	v356 = v356 + int32(1)
	goto L97
L113:
	;
	return
}
