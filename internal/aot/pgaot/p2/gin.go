package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GinNewBuffer(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return v90
L2:
	;
	return int32(0)
L3:
	;
	if v9 != int32(-1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v9
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v80
	v82 = int32(8)
	v84 = int32(0)
	v87 = F_ExtendBufferedRel(m, v7+v82, v84, v84, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L2
	} else {
		goto L28
	}
L7:
	;
	v19 = F_ReadBuffer(m, l0, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v21 = F_ConditionalLockBuffer(m, v19)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v21 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v19 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	goto L13
L13:
	;
	F_ReleaseBuffer(m, v19)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L2
	} else {
		goto L25
	}
L14:
	;
	if v60 != 0 {
		v90 = v19
		goto L1
	} else {
		goto L23
	}
L15:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+14)))
	if v42 == int32(0) {
		v60 = int32(1)
		goto L14
	} else {
		goto L19
	}
L16:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v19^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L15
L17:
	;
	goto L18
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v41 = v35 + v19<<(uint(int32(13))%32) + int32(-8192)
	goto L15
L19:
	;
	v45 = int32(0)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+16)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v46)+6)))
	if v48&int32(4) == v45 {
		v60 = v45
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	if v54 == int32(0) {
		v60 = int32(1)
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v57 = F_GlobalVisCheckRemovableXid(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v60 = v57
	goto L14
L23:
	;
	F_LockBuffer(m, v19, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	v67 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v67 != int32(-1) {
		v16 = v67
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v90 = v87
	goto L1
}
func F__gin_parallel_build_main(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	v10 = m.G0
	v12 = v10 - int32(5808)
	m.G0 = v12
	v17 = F_shm_toc_lookup(m, l1, int64(-5764607523034234877), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[30])) = v17
		F_pgstat_report_activity(m, int32(3), v17)
		mBase = m.M
		v24 = F_shm_toc_lookup(m, l1, int64(-5764607523034234879), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
			if v29 != 0 {
				v30 = int32(4)
			} else {
				v30 = int32(5)
			}
			v31 = F_table_open(m, v26, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				if v29 != 0 {
					v36 = int32(3)
				} else {
					v36 = int32(8)
				}
				v37 = F_index_open(m, v33, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_initGinState(m, v12, v37)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v43 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[31]))) = uint16(v43)
						v47 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[32]))) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[33]))) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[34]))) = v47
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[35]))) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[36]))) = v43
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[37]))) = v47
						v66 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v71 = F_AllocSetContextCreateInternal(m, v66, int32(65007), v43, int32(8192), int32(8388608))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[38]))) = v71
							v75 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							v80 = F_AllocSetContextCreateInternal(m, v75, int32(264373), int32(0), int32(8192), int32(8388608))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[39]))) = v80
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_consts[40]))) = v12
								F_ginInitBA(m, v12+int32(5728))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v90 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), int32(0))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_tuplesort_attach_shared(m, v90, l0)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return
										} else {
											v97 = F___memcpy(m, int32(4447192), int32(4447032), int32(128))
											mBase = m.M
											v100 = *(*int64)(unsafe.Add(mBase, _consts[41]))
											*(*int64)(unsafe.Add(mBase, _consts[42])) = v100
											v104 = *(*int64)(unsafe.Add(mBase, _consts[43]))
											*(*int64)(unsafe.Add(mBase, _consts[44])) = v104
											v108 = *(*int64)(unsafe.Add(mBase, _consts[45]))
											*(*int64)(unsafe.Add(mBase, _consts[46])) = v108
											v112 = *(*int64)(unsafe.Add(mBase, _consts[47]))
											*(*int64)(unsafe.Add(mBase, _consts[48])) = v112
											v115 = *(*int32)(unsafe.Add(mBase, _consts[49]))
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
											v117 = base.I32_div_s(v115, v116)
											F__gin_parallel_scan_and_build(m, v12, v24, v90, v31, v37, v117, int32(0))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v123 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return
												} else {
													v127 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return
													} else {
														v130 = *(*int32)(unsafe.Add(mBase, _consts[50]))
														v136 = v127 + v130<<(uint(int32(5))%32)
														v141 = F___memset(m, v123+v130<<(uint(int32(7))%32), int32(0), int32(128))
														mBase = m.M
														F_BufferUsageAccumDiff(m, v141, int32(4447192))
														mBase = m.M
														v145 = v136 + int32(24)
														v146 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v145))) = v146
														v149 = v136 + int32(16)
														*(*int64)(unsafe.Add(mBase, uint32(v149))) = v146
														v153 = v136 + int32(8)
														*(*int64)(unsafe.Add(mBase, uint32(v153))) = v146
														*(*int64)(unsafe.Add(mBase, uint32(v136))) = v146
														v159 = *(*int64)(unsafe.Add(mBase, _consts[43]))
														v161 = *(*int64)(unsafe.Add(mBase, _consts[44]))
														*(*int64)(unsafe.Add(mBase, uint32(v149))) = v159 - v161
														v165 = *(*int64)(unsafe.Add(mBase, _consts[47]))
														v167 = *(*int64)(unsafe.Add(mBase, _consts[48]))
														*(*int64)(unsafe.Add(mBase, uint32(v136))) = v165 - v167
														v171 = *(*int64)(unsafe.Add(mBase, _consts[45]))
														v173 = *(*int64)(unsafe.Add(mBase, _consts[46]))
														*(*int64)(unsafe.Add(mBase, uint32(v153))) = v171 - v173
														v177 = *(*int64)(unsafe.Add(mBase, _consts[41]))
														v179 = *(*int64)(unsafe.Add(mBase, _consts[42]))
														*(*int64)(unsafe.Add(mBase, uint32(v145))) = v177 - v179
														F_relation_close(m, v37, v36)
														mBase = m.M
														v183 = m.ExcPending
														if v183 != 0 {
															return
														} else {
															F_sequence_close(m, v31, v30)
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v12 + int32(5808)
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
func F__gin_parallel_scan_and_build(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v54 float64
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v166 int32
	_ = v166
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 float64
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v494 int32
	_ = v494
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 float64
	_ = v523
	var v526 int32
	_ = v526
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 float64
	_ = v540
	var v541 float64
	_ = v541
	var v544 float64
	_ = v544
	var v547 float64
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v19 = F_palloc0(m, int32(12))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(-1)
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v24)
	v27 = base.I32_div_s(l5, int32(2))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[22]))) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[23]))) = v29
	v31 = F_tuplesort_begin_index_gin(m, l4, v27, v19)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24]))) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[22])))
	v36 = F_tuplesort_begin_index_gin(m, l4, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25]))) = v36
	v39 = F_BuildIndexInfo(m, l4)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+121)) = uint8(v41)
	v44 = int32(0)
	v50 = F_table_beginscan_parallel(m, l3, l1-int32(-64))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+188))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+140))
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l3, l4, v39, int32(1), v44, l6, v44, int32(-1), int32(54), l0, v50)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_ginFlushBuildState(m, l0, l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25])))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v60 = F_GinBufferInit(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if l6 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v146 = F_tuplesort_getgintuple(m, v58, v16+int32(12))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25])))
	F_tuplesort_performsort(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v73 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27]))) = int64(0)
	goto L10
L15:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[25])))
	F_tuplesort_performsort(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v77 != int32(1) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v80 = int32(4543428)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v83 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 + v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v86 + v83
	*(*int64)(unsafe.Add(mBase, uint32(v73+int32(80))+232)) = int64(3)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v94 + v83
	v100 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v100 - v83
	goto L16
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27]))) = int64(0)
	v113 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	if v113 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L10
L21:
	;
	goto L20
L22:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, _consts[28])))
	if v117 != int32(1) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v120 = int32(4543428)
	v122 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v122 + v123
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v126 + v123
	*(*int64)(unsafe.Add(mBase, uint32(v113+int32(80))+232)) = int64(4)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v134 + v123
	v140 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v140 - v123
	goto L21
L24:
	;
	if v146 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v149 = v60 + int32(4)
	v155 = v146
	goto L28
L26:
	;
	goto L27
L27:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v464 != 0 {
		goto L115
	} else {
		goto L116
	}
L28:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v166 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v169 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L32
L34:
	;
	F_GinBufferStoreTuple(m, v60, v155)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L112
	}
L35:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	if v258 < int32(1024) {
		goto L34
	} else {
		goto L61
	}
L36:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+4)))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	if v173 != v174 {
		v214 = v174
		v215 = v172
		v216 = v169
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+12)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v227 = F__gin_build_tuple(m, v214&int32(65535), v215&int32(255), v221, v222, v223, v224, v216, v16+int32(8))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L54
	}
L38:
	;
	v177 = v172 & int32(255)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+11)))
	if v177 != v178 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v214 = v173
	v215 = v172
	v216 = v169
	goto L37
L40:
	;
	goto L41
L41:
	;
	if v177 != 0 {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v181 = v155 + int32(16)
	v182 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	if v185 == v182 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v189 = v188
	goto L45
L44:
	;
	v189 = v181
	goto L45
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	v191 = int32(36)
	v193 = v190 + v173*v191
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193-int32(20))))
	v199 = m.T0[v198].(func(*base.Module, int32, int32, int32) int32)(m, v184, v189, v193-v191)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v199 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v204 = v182
	goto L49
L48:
	;
	v204 = int32(0) - v199
	goto L49
L49:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193-int32(28)))))
	if v207 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v208 = v204
	goto L52
L51:
	;
	v208 = v199
	goto L52
L52:
	;
	if v208 == int32(0) {
		goto L35
	} else {
		goto L53
	}
L53:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v214 = v213
	v215 = v212
	v216 = v211
	goto L37
L54:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_tuplesort_putgintuple(m, v229, v227, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v233 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27]))) = base.F64_add(v233, float64(1))
	F_pfree(m, v227)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v239 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v244 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+20)) = v244
	v246 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v246)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v246)
	*(*int32)(unsafe.Add(mBase, uint32(v60+int32(11)))) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v149))) = v244
	goto L34
L58:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	if v240 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	F_pfree(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v262+v263 < v261 {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+12)))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v274 = F__gin_build_tuple(m, v266, v267, v268, v269, v270, v271, v258, v16+int32(8))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_tuplesort_putgintuple(m, v276, v274, v277)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_pfree(m, v274)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v284 = int32(6)
	v286 = v282 + v283*v284
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v290 = (v287 - v283) * v284
	if v282 == v286 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v438 - v435
	goto L34
L67:
	;
	goto L66
L68:
	;
	v294 = v282 + v290
	if base.Ui32(v286-v294) <= base.Ui32(int32(0)-v290<<(uint(int32(1))%32)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v301 = F___memcpy(m, v282, v286, v290)
	mBase = m.M
	goto L66
L70:
	;
	goto L71
L71:
	;
	v304 = (v282 ^ v286) & int32(3)
	if base.Ui32(v282) < base.Ui32(v286) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v406 == int32(0) {
		goto L67
	} else {
		goto L108
	}
L73:
	;
	if base.Ui32(v384) <= base.Ui32(int32(3)) {
		v405 = v383
		v406 = v384
		v407 = v385
		goto L72
	} else {
		goto L104
	}
L74:
	;
	if v304 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v304 != 0 {
		v366 = v290
		goto L87
	} else {
		goto L88
	}
L77:
	;
	v405 = v286
	v406 = v290
	v407 = v282
	goto L72
L78:
	;
	goto L79
L79:
	;
	if v282&int32(3) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v383 = v286
	v384 = v290
	v385 = v282
	goto L73
L81:
	;
	goto L82
L82:
	;
	v311 = v286
	v312 = v290
	v313 = v282
	goto L83
L83:
	;
	if v312 == int32(0) {
		goto L67
	} else {
		goto L85
	}
L84:
	;
	v383 = v320
	v384 = v322
	v385 = v324
	goto L73
L85:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311))))
	*(*uint8)(unsafe.Add(mBase, uint32(v313))) = uint8(v317)
	v319 = int32(1)
	v320 = v311 + v319
	v322 = v312 - v319
	v324 = v313 + v319
	if v324&int32(3) != 0 {
		v311 = v320
		v312 = v322
		v313 = v324
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	if v366 == int32(0) {
		goto L67
	} else {
		goto L100
	}
L88:
	;
	if v294&int32(3) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v331 = v290
	goto L92
L90:
	;
	v346 = v290
	goto L91
L91:
	;
	if base.Ui32(v346) <= base.Ui32(int32(3)) {
		v366 = v346
		goto L87
	} else {
		goto L96
	}
L92:
	;
	if v331 == int32(0) {
		goto L67
	} else {
		goto L94
	}
L93:
	;
	v346 = v337
	goto L91
L94:
	;
	v337 = v331 - int32(1)
	v338 = v282 + v337
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v337))))
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v340)
	if v338&int32(3) != 0 {
		v331 = v337
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v353 = v346
	goto L97
L97:
	;
	v357 = v353 - int32(4)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v286+v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v282+v357))) = v360
	if base.Ui32(int32(3)) < base.Ui32(v357) {
		v353 = v357
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v366 = v357
	goto L87
L99:
	;
	goto L98
L100:
	;
	v373 = v366
	goto L101
L101:
	;
	v377 = v373 - int32(1)
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v377))))
	*(*uint8)(unsafe.Add(mBase, uint32(v282+v377))) = uint8(v380)
	if v377 != 0 {
		v373 = v377
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L67
L103:
	;
	goto L102
L104:
	;
	v390 = v383
	v391 = v384
	v392 = v385
	goto L105
L105:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	*(*int32)(unsafe.Add(mBase, uint32(v392))) = v394
	v396 = int32(4)
	v397 = v390 + v396
	v399 = v392 + v396
	v401 = v391 - v396
	if base.Ui32(int32(3)) < base.Ui32(v401) {
		v390 = v397
		v391 = v401
		v392 = v399
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v405 = v397
	v406 = v401
	v407 = v399
	goto L72
L107:
	;
	goto L106
L108:
	;
	v412 = v405
	v413 = v406
	v414 = v407
	goto L109
L109:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412))))
	*(*uint8)(unsafe.Add(mBase, uint32(v414))) = uint8(v416)
	v418 = int32(1)
	v423 = v413 - v418
	if v423 != 0 {
		v412 = v412 + v418
		v413 = v423
		v414 = v414 + v418
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L67
L111:
	;
	goto L110
L112:
	;
	v449 = F_tuplesort_getgintuple(m, v58, v16+int32(12))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v449 != 0 {
		v155 = v449
		goto L28
	} else {
		goto L114
	}
L114:
	;
	goto L29
L115:
	;
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v468 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+12)))
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v473 = F__gin_build_tuple(m, v465, v466, v467, v468, v469, v470, v464, v16+int32(8))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	if v503 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L118:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	F_tuplesort_putgintuple(m, v475, v473, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v479 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27]))) = base.F64_add(v479, float64(1))
	F_pfree(m, v473)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v486 = v60 + int32(4)
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v487 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v492 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+20)) = v492
	v494 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v494)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v494)
	*(*int32)(unsafe.Add(mBase, uint32(v486)+7)) = v494
	*(*int64)(unsafe.Add(mBase, uint32(v486))) = v492
	goto L117
L122:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	if v488 != 0 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v486)))
	F_pfree(m, v489)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	F_pfree(m, v60)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L132
	}
L126:
	;
	F_pfree(m, v503)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v508 == int32(0) {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v511 != 0 {
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+14)))
	if v512 != 0 {
		goto L125
	} else {
		goto L130
	}
L130:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	F_pfree(m, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L125
L132:
	;
	F_tuplesort_end(m, v58)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	F_tuplesort_performsort(m, v520)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v523 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[29])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[29]))) = base.F64_add(v54, v523)
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(1)
	if v526 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	F_s_lock(m, l1+int32(28), int32(511906), int32(2086), int32(447276))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v536 + int32(1)
	v540 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[29])))
	v541 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v540, v541)
	v544 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[27])))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(0)
	v547 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = base.F64_add(v544, v547)
	F_ConditionVariableSignal(m, l1+int32(16))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L139
	}
L138:
	;
	goto L137
L139:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[24])))
	F_tuplesort_end(m, v554)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	m.G0 = v16 + int32(16)
	return
}
func F_ginCompareAttEntries(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	if l1 != l4 {
		if base.Ui32(l1) < base.Ui32(l4) {
			v12 = int32(-1)
		} else {
			v12 = int32(1)
		}
		return v12
	} else {
		if l3 != l6 {
			if l3 < l6 {
				v18 = int32(-1)
			} else {
				v18 = int32(1)
			}
			return v18
		} else {
			if l3 != 0 {
				v36 = int32(0)
				return v36
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+uint32(_consts[53])))
				v32 = F_FunctionCall2Coll(m, l1*int32(28)+l0+int32(112), v31, l2, l5)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = v32
					return v36
				}
			}
		}
	}
}
func F_ginInsertValue(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v29)+6)))
	if v32&int32(64) != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(v11^int32(-1))<<(uint(int32(2))%32))))
	v29 = v21
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v29 = v23 + v11<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v37 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v61 = F_ginPlaceToPage(m, l0, l1, l2, int32(-1), int32(0), l3)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return
L9:
	;
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v40 + int32(4)
	F_errmsg_internal(m, int32(719155), v9)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_ginFinishSplit(m, l0, l1, int32(0), l3)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	F_errfinish(m, int32(519531), int32(783), int32(108824))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L7
L16:
	;
	m.G0 = v9 + int32(16)
	return
L17:
	;
	if v61 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_LockBuffer(m, v63, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_ginFinishSplit(m, l0, l1, int32(1), l3)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L30
	}
L21:
	;
	v68 = l1
	goto L22
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)+20))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L16
L24:
	;
	F_ReleaseBuffer(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_pfree(m, v68)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	if v73 != 0 {
		v68 = v73
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	goto L16
}
func F_ginStepRight(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	if l0 < int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l0^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[6]))
		v23 = v17 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
	v25 = v24 + v23
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = F_ReadBuffer(m, l1, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v28, l2)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_UnlockReleaseBuffer(m, l0)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				if v28 < int32(0) {
					v39 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v28^int32(-1))<<(uint(int32(2))%32))))
					v53 = v45
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v53 = v47 + v28<<(uint(int32(13))%32) + int32(-8192)
				}
				v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
				v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54+v53)+6)))
				if (v56^v26)&int32(3) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(382836), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(519531), int32(192), int32(110809))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					return v28
				}
			}
		}
	}
}
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_palloc(m, int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v27 = F_palloc(m, int32(16))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
			v32 = F_palloc(m, v29)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v32
				v35 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v35)
				if v2 != 0 {
					v37 = F_pg_detoast_datum(m, v18)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v37
						*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = l3
						*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v2)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v39
						*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v17)
						v45 = F_palloc(m, int32(4))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v45))) = v27
							switch v17 - int32(1) {
							case 0, 1:
								v71 = m.T0[l2].(func(*base.Module) int32)(m)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22))) = v71
									v74 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v74)
									m.G0 = v14 + int32(16)
									return v22
								}
							case 2:
								*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
								m.G0 = v14 + int32(16)
								return v22
							case 3, 4:
								v51 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v51)
								*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
								m.G0 = v14 + int32(16)
								return v22
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
									F_errmsg_internal(m, int32(500693), v14)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516266), int32(97), int32(15809))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
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
					}
				} else {
					v39 = v18
					*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = l3
					*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v2)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v39
					*(*uint16)(unsafe.Add(mBase, uint32(v27))) = uint16(v17)
					v45 = F_palloc(m, int32(4))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v45))) = v27
						switch v17 - int32(1) {
						case 0, 1:
							v71 = m.T0[l2].(func(*base.Module) int32)(m)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22))) = v71
								v74 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v74)
								m.G0 = v14 + int32(16)
								return v22
							}
						case 2:
							*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
							m.G0 = v14 + int32(16)
							return v22
						case 3, 4:
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v51)
							*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
							m.G0 = v14 + int32(16)
							return v22
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v17
								F_errmsg_internal(m, int32(500693), v14)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(516266), int32(97), int32(15809))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
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
				}
			}
		}
	}
}
func F_gin_consistent_jsonb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	switch v15 - int32(7) {
	case 0:
		goto L5
	default:
		goto L3
	case 2, 3:
		goto L2
	case 4:
		goto L4
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v92
L2:
	;
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v88)
	v92 = v88
	goto L1
L3:
	;
	v56 = int32(1)
	if base.Ui32((v15-int32(15))&int32(65535)) <= base.Ui32(v56) {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v37)
	v40 = int32(0)
	if v13 <= v40 {
		v92 = v37
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v18 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v18)
	v21 = int32(0)
	if v13 <= v21 {
		v92 = v18
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v24 = v21
	goto L7
L7:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v14))))
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v92 = int32(0)
	goto L1
L9:
	;
	v34 = v24 + int32(1)
	if v13 != v34 {
		v24 = v34
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v92 = v18
	goto L1
L13:
	;
	v43 = v40
	goto L14
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v14))))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v92 = int32(0)
	goto L1
L16:
	;
	v53 = v43 + int32(1)
	if v13 != v53 {
		v43 = v53
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	v92 = v37
	goto L1
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v64)
	if v13 <= int32(0) {
		v92 = v56
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L24
	} else {
		goto L26
	}
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v69 = F_execute_jsp_gin_node(m, v68, v14)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	v92 = base.B2i32(v69 != int32(0))
	goto L1
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
	F_errmsg_internal(m, int32(500693), v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(516278), int32(1007), int32(523523))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_hstore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_hstoreUpgrade(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v19 = v17 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19 << (uint(int32(1)) % 32)
	if v19 == int32(0) {
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
	v28 = v12 + int32(8)
	v30 = v19 << (uint(int32(3)) % 32)
	v31 = v28 + v30
	v33 = F_palloc(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v35 = int32(0)
	goto L7
L7:
	;
	v47 = v28 + v35<<(uint(int32(3))%32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v48 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v33
L9:
	;
	v66 = v62 + int32(5)
	v67 = F_palloc(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	v62 = v48 & int32(1073741823)
	v64 = v31
	goto L9
L11:
	;
	goto L12
L12:
	;
	v53 = int32(1073741823)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47-int32(4))))
	v59 = v57 & v53
	v62 = v48&v53 - v59
	v64 = v59 + v31
	goto L9
L13:
	;
	v69 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v66 << (uint(int32(2)) % 32)
	if int32(0) < v62 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v62 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	goto L16
L16:
	;
	v80 = int32(1)
	v81 = v35 << (uint(v80) % 32)
	v82 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v33+v81<<(uint(v82)%32)))) = v67
	v89 = (v81 | v80) << (uint(v82) % 32)
	v90 = v28 + v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91&int32(1073741824) != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v78 = F__emscripten_memcpy_bulkmem(m, v67+int32(5), v64, v62)
	mBase = m.M
	goto L20
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89+v33))) = v130
	v137 = v35 + int32(1)
	if v137 != v19 {
		v35 = v137
		goto L7
	} else {
		goto L36
	}
L22:
	;
	v95 = F_palloc(m, int32(5))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v91 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v97 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+4)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(20)
	v130 = v95
	goto L21
L26:
	;
	v116 = v112 + int32(5)
	v117 = F_palloc(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v112 = v91 & int32(1073741823)
	v114 = v31
	goto L26
L28:
	;
	goto L29
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v90-int32(4))))
	v109 = v107 & int32(1073741823)
	v112 = v91 - v109
	v114 = v109 + v31
	goto L26
L30:
	;
	v119 = int32(86)
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)) = uint8(v119)
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v116 << (uint(int32(2)) % 32)
	if v112 <= int32(0) {
		v130 = v117
		goto L21
	} else {
		goto L31
	}
L31:
	;
	if v112 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v130 = v117
	goto L21
L33:
	;
	v128 = F__emscripten_memcpy_bulkmem(m, v117+int32(5), v114, v112)
	mBase = m.M
	goto L35
L34:
	;
	goto L35
L35:
	;
	goto L32
L36:
	;
	goto L8
}
func F_gin_extract_jsonb_query_path(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if v12 == int32(7) {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v18 = F_DirectFunctionCall2Coll(m, int32(1341), int32(0), v17, v11)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v22 == int32(0) {
				v39 = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
				v42 = v39
			} else {
				v42 = v18
			}
			m.G0 = v8 + int32(16)
			return v42
		}
	} else {
		if base.Ui32(int32(1)) < base.Ui32((v12-int32(15))&int32(65535)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
				F_errmsg_internal(m, int32(500693), v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(516278), int32(1212), int32(334211))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v32 = F_pg_detoast_datum(m, v31)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v36 = F_extract_jsp_query(m, v32, v12, int32(1), v11, v35)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != 0 {
						v42 = v36
					} else {
						v39 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(2)
						v42 = v39
					}
					m.G0 = v8 + int32(16)
					return v42
				}
			}
		}
	}
}
func F_gin_extract_query_bool(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2724)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
func F_gin_extract_query_char(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2118)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
func F_gin_extract_query_float4(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2114)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(-8388608)
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v71)
						m.G0 = v11 + int32(16)
						return v19
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
func F_gin_extract_query_float8(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2115)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_Float8GetDatum(m, math.Float64frombits(uint64(0xfff0000000000000)))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
func F_gin_extract_query_macaddr8(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(4194)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_palloc0(m, int32(8))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
func F_gin_extract_query_timestamp(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_palloc(m, int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v24 = F_palloc(m, int32(16))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v17))) = v26
			v29 = F_palloc(m, v26)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v29
				v32 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(1446)
				*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)) = uint8(v32)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v15
				*(*uint16)(unsafe.Add(mBase, uint32(v24))) = uint16(v14)
				v41 = F_palloc(m, int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v24
					v46 = v14 & int32(65535)
					switch v46 - int32(1) {
					case 0, 1:
						v70 = F_Int64GetDatum(m, int64(-9223372036854775807-1))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v70
							v73 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v73)
							m.G0 = v11 + int32(16)
							return v19
						}
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					case 3, 4:
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v49)
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = v15
						m.G0 = v11 + int32(16)
						return v19
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v46
							F_errmsg_internal(m, int32(500693), v11)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(516266), int32(97), int32(15809))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			}
		}
	}
}
