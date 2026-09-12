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
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v270 int64
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int64
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v345 int64
	_ = v345
	var v347 int64
	_ = v347
	var v349 int64
	_ = v349
	var v350 int32
	_ = v350
	var v352 int64
	_ = v352
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v398 int64
	_ = v398
	var v399 int64
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	v3 = int32(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v10 + int64(1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v14 {
	case 0:
		goto L9
	case 1:
		goto L8
	case 2:
		goto L7
	default:
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v292 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v301+v292<<(uint(int32(2))%32)))) = l1
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v306 < v307 {
		goto L82
	} else {
		goto L83
	}
L2:
	;
	v290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v290)
	v292 = v53
	goto L1
L3:
	;
	if v62 < base.I64_extend_i32_u((v245-v54)<<(uint(int32(2))%32)) {
		goto L2
	} else {
		goto L71
	}
L4:
	;
	v242 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v242)
	v245 = int32(536870911)
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L46
	} else {
		goto L67
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L46
	} else {
		goto L64
	}
L7:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v146 = v142 + v143*int32(24)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+4)))
	if v147 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v92 {
		goto L36
	} else {
		goto L37
	}
L9:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v15 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v21 = v18
	v22 = v3
	v23 = v15
	goto L13
L11:
	;
	goto L12
L12:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v53 < v54-int32(1) {
		v292 = v53
		goto L1
	} else {
		goto L19
	}
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)))
	if v28 != int32(1) {
		v38 = v23
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L12
L15:
	;
	v42 = v22 + int32(1)
	if v42 < v38 {
		v21 = v21 + int32(24)
		v22 = v42
		v23 = v38
		goto L13
	} else {
		goto L18
	}
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v22 == v31 {
		v38 = v23
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v33 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+4)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v38 = v37
	goto L15
L18:
	;
	goto L14
L19:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)))
	if v58 != int32(1) {
		v292 = v53
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v63 = v61 - v62
	if v63 <= v62 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v54 != int32(2147483647) {
		goto L4
	} else {
		goto L35
	}
L22:
	;
	if v86 <= v54 {
		goto L2
	} else {
		goto L33
	}
L23:
	;
	if int32(1073741822) < v54 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+88)) = uint8(v69)
	v75 = base.F64_mul(base.F64_div(base.F64_convert_i64_s(v61), base.F64_convert_i64_s(v63)), base.F64_convert_i32_s(v54))
	v76 = float64(2.147483647e+09)
	if base.F64_lt(v75, v76) != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v86 = v54 << (uint(int32(1)) % 32)
	goto L22
L27:
	;
	v79 = v75
	goto L29
L28:
	;
	v79 = v76
	goto L29
L29:
	;
	if base.F64_lt(base.F64_abs(v79), float64(2.147483648e+09)) != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = base.I32_trunc_f64_s(v79)
	v86 = v83
	goto L22
L31:
	;
	goto L32
L32:
	;
	v86 = int32(-2147483648)
	goto L22
L33:
	;
	if base.Ui32(v86) < base.Ui32(int32(536870911)) {
		v245 = v86
		goto L3
	} else {
		goto L34
	}
L34:
	;
	goto L4
L35:
	;
	goto L2
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v98 = v95
	v99 = v3
	v100 = v92
	goto L39
L37:
	;
	goto L38
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	m.T0[v139].(func(*base.Module, int32, int32))(m, l0, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L46
	} else {
		goto L47
	}
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)))
	if v105 != int32(1) {
		v124 = v100
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v128 = v99 + int32(1)
	if v128 < v124 {
		v98 = v98 + int32(24)
		v99 = v128
		v100 = v124
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v99 == v108 {
		v124 = v100
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v110 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+4)) = uint8(v110)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v98+int32(12)))) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v112)+32))
	v120 = int64(*(*int32)(unsafe.Add(mBase, uint32(v112)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v98+int32(16)))) = v119 + v120
	goto L44
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v124 = v123
	goto L41
L45:
	;
	goto L40
L46:
	;
	return
L47:
	;
	return
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v146+int32(12)))) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v150)+32))
	v158 = int64(*(*int32)(unsafe.Add(mBase, uint32(v150)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v146+int32(16)))) = v157 + v158
	goto L51
L49:
	;
	goto L50
L50:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v165 = F_BufFileSeek(m, v161, v162, v163, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L46
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	if v165 != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v169 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v175 = v172
	v176 = v3
	v177 = v169
	goto L57
L55:
	;
	goto L56
L56:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	m.T0[v209].(func(*base.Module, int32, int32))(m, l0, l1)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L46
	} else {
		goto L63
	}
L57:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
	if v182 != int32(1) {
		v194 = v177
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v198 = v176 + int32(1)
	if v198 < v194 {
		v175 = v175 + int32(24)
		v176 = v198
		v177 = v194
		goto L57
	} else {
		goto L62
	}
L60:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v176 == v185 {
		v194 = v177
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)) = uint8(v187)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v189
	v191 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+16)) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v194 = v193
	goto L59
L62:
	;
	goto L58
L63:
	;
	return
L64:
	;
	F_errmsg_internal(m, int32(352754), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L46
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(498595), int32(941), int32(245779))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L46
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
	F_errcode_for_file_access(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L46
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(386924), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L46
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(498595), int32(921), int32(245779))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L46
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
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v253 = F_GetMemoryChunkSpace(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L46
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v245
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v256 + base.I64_extend_i32_u(v253)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v263 = F_repalloc_huge(m, v260, v245<<(uint(int32(2))%32))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L46
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v263
	v266 = F_GetMemoryChunkSpace(m, v263)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L46
	} else {
		goto L74
	}
L74:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v270 = v268 - base.I64_extend_i32_u(v266)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v270
	if int64(0) <= v270 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v292 = v274
	goto L1
L76:
	;
	goto L77
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L46
	} else {
		goto L78
	}
L78:
	;
	F_errmsg_internal(m, int32(364116), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L46
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(498595), int32(716), int32(164031))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L46
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	return
L82:
	;
	v309 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) <= v309 {
		goto L81
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	F_PrepareTempTablespaces(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L46
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	v314 = int32(4515300)
	v315 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v317
	v319 = int32(4515248)
	v320 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+16))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v323
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v326 = F_BufFileCreateTemp(m, v325)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L46
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v326
	*(*int32)(unsafe.Add(mBase, _consts[179])) = v315
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v320
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v337 = int32(base.Ui32(v333)>>(uint(int32(3))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v337)
	v339 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v340 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v366 = v364
	goto L101
L89:
	;
	v343 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v344 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v345 = v343 - v344
	if v345 < v339 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v349 = F_BufFileSize(m, v326)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L46
	} else {
		goto L96
	}
L92:
	;
	v347 = v339
	goto L94
L93:
	;
	v347 = v345
	goto L94
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v347
	goto L88
L95:
	;
	v357 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)) = uint8(v357)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v356
	goto L88
L96:
	;
	if v349 < v339 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v352 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v356 = v352
	goto L95
L98:
	;
	goto L99
L99:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v354 = F_BufFileSize(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v356 = v354
	goto L95
L101:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if int32(0) < v374 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = int64(0)
	goto L81
L103:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v381 = v377
	v382 = int32(0)
	v383 = v374
	goto L106
L104:
	;
	goto L105
L105:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v366 < v418 {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v381)+8))
	if v366 != v388 {
		v403 = v383
		goto L108
	} else {
		goto L109
	}
L107:
	;
	goto L105
L108:
	;
	v407 = v382 + int32(1)
	if v407 < v403 {
		v381 = v381 + int32(24)
		v382 = v407
		v383 = v403
		goto L106
	} else {
		goto L112
	}
L109:
	;
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+4)))
	if v390 != 0 {
		v403 = v383
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v391)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v381+int32(12)))) = v396
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v391)+32))
	v399 = int64(*(*int32)(unsafe.Add(mBase, uint32(v391)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v381+int32(16)))) = v398 + v399
	goto L111
L111:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v403 = v402
	goto L108
L112:
	;
	goto L107
L113:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v366<<(uint(int32(2))%32))))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	m.T0[v425].(func(*base.Module, int32, int32))(m, l0, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L46
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L102
L116:
	;
	v366 = v366 + int32(1)
	goto L101
}
