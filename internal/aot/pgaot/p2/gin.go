package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"sync/atomic"
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
	var v78 int64
	_ = v78
	var v80 int32
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = l0
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v80
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
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_GinNewBuffer[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v19^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L15
L17:
	;
	goto L18
L18:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_GinNewBuffer[1]))
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
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int64
	_ = v132
	var v141 int64
	_ = v141
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	v10 = m.G0
	v12 = v10 - int32(_a_F__gin_parallel_build_main_0)
	m.G0 = v12
	v17 = F_shm_toc_lookup(m, l1, int64(-5764607523034234877), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[0])) = v17
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
						v41 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[1]))) = uint16(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[2]))) = v41
						v45 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[3]))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[4]))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[5]))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[6]))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[7]))) = v45
						v56 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[8]))
						v61 = F_AllocSetContextCreateInternal(m, v56, int32(_a_F__gin_parallel_build_main_1), v41, int32(_a_F__gin_parallel_build_main_2), int32(_a_F__gin_parallel_build_main_3))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[9]))) = v61
							v65 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[8]))
							v70 = F_AllocSetContextCreateInternal(m, v65, int32(_a_F__gin_parallel_build_main_4), int32(0), int32(_a_F__gin_parallel_build_main_2), int32(_a_F__gin_parallel_build_main_3))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[10]))) = v70
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[11]))) = v12
								F_ginInitBA(m, v12+int32(_a_F__gin_parallel_build_main_5))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									v80 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), int32(0))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										F_tuplesort_attach_shared(m, v80, l0)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											base.MemoryCopy(m, int32(_a_F__gin_parallel_build_main_6), int32(_a_F__gin_parallel_build_main_7), int32(128))
											v90 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[12]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[13])) = v90
											v94 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[14]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[15])) = v94
											v98 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[16]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[17])) = v98
											v102 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[18]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[19])) = v102
											v105 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[20]))
											v106 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
											v107 = base.I32_div_s(v105, v106)
											F__gin_parallel_scan_and_build(m, v12, v24, v80, v31, v37, v107, int32(0))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return
											} else {
												v113 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return
												} else {
													v117 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return
													} else {
														v120 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[21]))
														v123 = v113 + v120<<(uint(int32(7))%32)
														v126 = v117 + v120<<(uint(int32(5))%32)
														base.MemoryFill(m, v123, int32(0), int32(128))
														F_BufferUsageAccumDiff(m, v123, int32(_a_F__gin_parallel_build_main_6))
														mBase = m.M
														v132 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v126)+24)) = v132
														*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v132
														*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v132
														*(*int64)(unsafe.Add(mBase, uint32(v126))) = v132
														v141 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[14]))
														v143 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[15]))
														*(*int64)(unsafe.Add(mBase, uint32(v126)+16)) = v141 - v143
														v147 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[18]))
														v149 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[19]))
														*(*int64)(unsafe.Add(mBase, uint32(v126))) = v147 - v149
														v153 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[16]))
														v155 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[17]))
														*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v153 - v155
														v159 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[12]))
														v161 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[13]))
														*(*int64)(unsafe.Add(mBase, uint32(v126)+24)) = v159 - v161
														F_relation_close(m, v37, v36)
														mBase = m.M
														v165 = m.ExcPending
														if v165 != 0 {
															return
														} else {
															F_relation_close(m, v31, v30)
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return
															} else {
																m.G0 = v12 + int32(_a_F__gin_parallel_build_main_0)
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
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
	var v225 int32
	_ = v225
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
	var v234 float64
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
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
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v351 int32
	_ = v351
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 float64
	_ = v380
	var v385 int32
	_ = v385
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 float64
	_ = v397
	var v398 float64
	_ = v398
	var v401 float64
	_ = v401
	var v402 float64
	_ = v402
	var v405 int32
	_ = v405
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v18 = F_palloc0(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(-1)
	v23 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(v23)
	v26 = base.I32_div_s(l5, int32(2))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[0]))) = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[1]))) = v28
	v30 = F_tuplesort_begin_index_gin(m, l4, v26, v18)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2]))) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[0])))
	v35 = F_tuplesort_begin_index_gin(m, l4, v33, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3]))) = v35
	v38 = F_BuildIndexInfo(m, l4)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+121)) = uint8(v40)
	v43 = int32(0)
	v49 = F_table_beginscan_parallel(m, l3, l1-int32(-64))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+188))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+140))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l3, l4, v38, int32(1), v43, l6, v43, int32(-1), int32(54), l0, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_ginFlushBuildState(m, l0, l4)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = F_GinBufferInit(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	v149 = F_tuplesort_getgintuple(m, v57, v15+int32(12))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
	F_tuplesort_performsort(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[4]))
	if v72 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5]))) = int64(0)
	goto L10
L15:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
	F_tuplesort_performsort(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])))
	if v76&int32(1) == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v81 = int32(_a_F__gin_parallel_scan_and_build_0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7]))
	v84 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7])) = v83 + v84
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v87 + v84
	*(*int64)(unsafe.Add(mBase, uint32(v72+int32(80))+232)) = int64(3)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v95 + v84
	v101 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7]))
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7])) = v101 - v84
	goto L16
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5]))) = int64(0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[4]))
	if v114 == int32(0) {
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
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])))
	if v118&int32(1) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v123 = int32(_a_F__gin_parallel_scan_and_build_0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7]))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7])) = v125 + v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v129 + v126
	*(*int64)(unsafe.Add(mBase, uint32(v114+int32(80))+232)) = int64(4)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v137 + v126
	v143 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7]))
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7])) = v143 - v126
	goto L21
L24:
	;
	if v149 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v152 = v59 + int32(4)
	v156 = v149
	goto L28
L26:
	;
	goto L27
L27:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v321 != 0 {
		goto L72
	} else {
		goto L73
	}
L28:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[8]))
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
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
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
	F_GinBufferStoreTuple(m, v59, v156)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L69
	}
L35:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	if v258 < int32(1024) {
		goto L34
	} else {
		goto L61
	}
L36:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+4)))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	if v173 != v174 {
		v214 = v174
		v216 = v172
		v217 = v169
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+12)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	v228 = F__gin_build_tuple(m, v214&int32(_a_F__gin_parallel_scan_and_build_1), v216&int32(255), v222, v223, v224, v225, v217, v15+int32(8))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L54
	}
L38:
	;
	v177 = v172 & int32(255)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+11)))
	if v177 != v178 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v214 = v173
	v216 = v172
	v217 = v169
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
	v181 = v156 + int32(16)
	v182 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v59)+28))
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
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v214 = v213
	v216 = v212
	v217 = v211
	goto L37
L54:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_tuplesort_putgintuple(m, v230, v228, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v234 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5]))) = base.F64_add(v234, float64(1))
	F_pfree(m, v228)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	if v240 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v245 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v59)+20)) = v245
	v247 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)) = uint8(v247)
	*(*uint16)(unsafe.Add(mBase, uint32(v59))) = uint16(v247)
	*(*int32)(unsafe.Add(mBase, uint32(v152)+7)) = v247
	*(*int64)(unsafe.Add(mBase, uint32(v152))) = v245
	goto L34
L58:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	if v241 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	F_pfree(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v262+v263 < v261 {
		goto L34
	} else {
		goto L62
	}
L62:
	;
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+12)))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	v274 = F__gin_build_tuple(m, v266, v267, v268, v269, v270, v271, v258, v15+int32(8))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
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
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v286 = (v282 - v283) * int32(6)
	if v286 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	base.MemoryCopy(m, v287, v287+v283*int32(6), v286)
	goto L68
L67:
	;
	goto L68
L68:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v296 - v293
	goto L34
L69:
	;
	v307 = F_tuplesort_getgintuple(m, v57, v15+int32(12))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v307 != 0 {
		v156 = v307
		goto L28
	} else {
		goto L71
	}
L71:
	;
	goto L29
L72:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59))))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+12)))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	v330 = F__gin_build_tuple(m, v322, v323, v324, v325, v326, v327, v321, v15+int32(8))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v59)+32))
	if v360 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L75:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_tuplesort_putgintuple(m, v332, v330, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v336 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5]))) = base.F64_add(v336, float64(1))
	F_pfree(m, v330)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v343 = v59 + int32(4)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	if v344 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v349 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v59)+20)) = v349
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)) = uint8(v351)
	*(*uint16)(unsafe.Add(mBase, uint32(v59))) = uint16(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v343)+7)) = v351
	*(*int64)(unsafe.Add(mBase, uint32(v343))) = v349
	goto L74
L79:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	if v345 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	F_pfree(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	F_pfree(m, v59)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L89
	}
L83:
	;
	F_pfree(m, v360)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v365 == int32(0) {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+2)))
	if v368 != 0 {
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+14)))
	if v369 != 0 {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	F_pfree(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	goto L82
L89:
	;
	F_tuplesort_end(m, v57)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	F_tuplesort_performsort(m, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v380 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9]))) = base.F64_add(v53, v380)
	v385 = base.AtomicRmwXchg32(m, l1, int32(28), int32(1))
	if v385 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_s_lock(m, l1+int32(28), int32(_a_F__gin_parallel_scan_and_build_2), int32(2086), int32(_a_F__gin_parallel_scan_and_build_3))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v393 + int32(1)
	v397 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9])))
	v398 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v397, v398)
	v401 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[5])))
	v402 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = base.F64_add(v401, v402)
	v405 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+28)), uint32(v405))
	F_ConditionVariableSignal(m, l1+int32(16))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	F_tuplesort_end(m, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	m.G0 = v15 + int32(16)
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
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l1<<(uint(int32(2))%32)+l0)+uint32(_c_F_ginCompareAttEntries[0])))
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertValue[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(v11^int32(-1))<<(uint(int32(2))%32))))
	v29 = v21
	goto L1
L3:
	;
	goto L4
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_ginInsertValue[1]))
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
	F_errmsg_internal(m, int32(_a_F_ginInsertValue_0), v9)
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
	F_errfinish(m, int32(_a_F_ginInsertValue_1), int32(783), int32(_a_F_ginInsertValue_2))
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
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l0^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[1]))
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
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[0]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v28^int32(-1))<<(uint(int32(2))%32))))
					v53 = v45
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[1]))
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
						F_errmsg_internal(m, int32(_a_F_ginStepRight_0), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_ginStepRight_1), int32(192), int32(_a_F_ginStepRight_2))
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
								v67 = m.T0[l2].(func(*base.Module) int32)(m)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v22))) = v67
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v70)
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
									F_errmsg_internal(m, int32(_a_F_gin_btree_extract_query_0), v14)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_gin_btree_extract_query_1), int32(97), int32(_a_F_gin_btree_extract_query_2))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
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
							v67 = m.T0[l2].(func(*base.Module) int32)(m)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22))) = v67
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v70)
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
								F_errmsg_internal(m, int32(_a_F_gin_btree_extract_query_0), v14)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_gin_btree_extract_query_1), int32(97), int32(_a_F_gin_btree_extract_query_2))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
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
	if base.Ui32((v15-int32(15))&int32(_a_F_gin_consistent_jsonb_0)) <= base.Ui32(v56) {
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
	F_errmsg_internal(m, int32(_a_F_gin_consistent_jsonb_1), v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_jsonb_2), int32(1007), int32(_a_F_gin_consistent_jsonb_3))
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
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
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
	v74 = int32(0)
	if base.B2i32(v62 == v74)|base.B2i32(v62 <= v74) == v74 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	base.MemoryCopy(m, v67+int32(5), v64, v62)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v84 = int32(1)
	v85 = v35 << (uint(v84) % 32)
	v86 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v33+v85<<(uint(v86)%32)))) = v67
	v93 = (v85 | v84) << (uint(v86) % 32)
	v94 = v28 + v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95&int32(1073741824) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33+v93))) = v136
	v143 = v35 + int32(1)
	if v143 != v19 {
		v35 = v143
		goto L7
	} else {
		goto L28
	}
L18:
	;
	v99 = F_palloc(m, int32(5))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v95 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v101 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+4)) = uint8(v101)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = int32(20)
	v136 = v99
	goto L17
L22:
	;
	v120 = v117 + int32(5)
	v121 = F_palloc(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v117 = v95 & int32(1073741823)
	v118 = v31
	goto L22
L24:
	;
	goto L25
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v94-int32(4))))
	v113 = v111 & int32(1073741823)
	v117 = v95 - v113
	v118 = v113 + v31
	goto L22
L26:
	;
	v123 = int32(86)
	*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v120 << (uint(int32(2)) % 32)
	v128 = int32(0)
	if base.B2i32(v117 == v128)|base.B2i32(v117 <= v128) != 0 {
		v136 = v121
		goto L17
	} else {
		goto L27
	}
L27:
	;
	base.MemoryCopy(m, v121+int32(5), v118, v117)
	v136 = v121
	goto L17
L28:
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
		v18 = F_DirectFunctionCall2Coll(m, int32(1325), int32(0), v17, v11)
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
		if base.Ui32(int32(1)) < base.Ui32((v12-int32(15))&int32(_a_F_gin_extract_jsonb_query_path_0)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
				F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_query_path_1), v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_extract_jsonb_query_path_2), int32(1212), int32(_a_F_gin_extract_jsonb_query_path_3))
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(0), int32(2708))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_char(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(0), int32(2102))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_float4(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13907(m, l0, int32(-8388608), int32(2098))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(2099)
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
					v46 = v14 & int32(_a_F_gin_extract_query_float8_0)
					switch v46 - int32(1) {
					case 0, 1:
						v66 = F_Float8GetDatum(m, math.Float64frombits(uint64(0xfff0000000000000)))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v66
							v69 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v69)
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
							F_errmsg_internal(m, int32(_a_F_gin_extract_query_float8_1), v11)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_extract_query_float8_2), int32(97), int32(_a_F_gin_extract_query_float8_3))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13909(m, l0, int32(8), int32(_a_F_gin_extract_query_macaddr8_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_extract_query_timestamp(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13908(m, l0, int64(-9223372036854775807-1), int32(1430))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_gin_metapage_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v5 = m.G0
	v7 = v5 - int32(96)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_superuser(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				v16 = F_get_page_from_raw(m, v10)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+14)))
					if v18 == int32(0) {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
						v95 = int32(0)
						m.G0 = v7 + int32(96)
						return v95
					} else {
						v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
						v25 = int32(8)
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)))
						if (v24<<(uint(v25)%32)-v27)&int32(_a_F_gin_metapage_info_0) != v25 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_gin_metapage_info_1), int32(0))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+16)))
										v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
										v129 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v129
										*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = (v128<<(uint(v129)%32) - v127) & int32(_a_F_gin_metapage_info_0)
										F_errdetail(m, int32(_a_F_gin_metapage_info_2), v7+int32(16))
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(55), int32(_a_F_gin_metapage_info_4))
											mBase = m.M
											v146 = m.ExcPending
											if v146 != 0 {
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
						} else {
							v33 = v16 + v27
							v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
							if v34 != int32(8) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_gin_metapage_info_5), int32(0))
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v158
											F_errdetail(m, int32(_a_F_gin_metapage_info_6), v7)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(64), int32(_a_F_gin_metapage_info_4))
												mBase = m.M
												v169 = m.ExcPending
												if v169 != 0 {
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
							} else {
								v40 = F_get_call_result_type(m, l0, int32(0), v7+int32(92))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									if v40 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_gin_metapage_info_7), int32(0))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(68), int32(_a_F_gin_metapage_info_4))
												mBase = m.M
												v182 = m.ExcPending
												if v182 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v44 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v7)+40)) = uint16(v44)
										*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(0)
										v48 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+24)))
										v49 = F_Int64GetDatum(m, v48)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v49
											v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+28)))
											v53 = F_Int64GetDatum(m, v52)
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v53
												v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v56
												v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+36)))
												v59 = F_Int64GetDatum(m, v58)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = v59
													v62 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
													v63 = F_Int64GetDatum(m, v62)
													mBase = m.M
													v64 = m.ExcPending
													if v64 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = v63
														v66 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+48)))
														v67 = F_Int64GetDatum(m, v66)
														mBase = m.M
														v68 = m.ExcPending
														if v68 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = v67
															v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+52)))
															v71 = F_Int64GetDatum(m, v70)
															mBase = m.M
															v72 = m.ExcPending
															if v72 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v71
																v74 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v16)+56)))
																v75 = F_Int64GetDatum(m, v74)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v75
																	v78 = *(*int64)(unsafe.Add(mBase, uint32(v16)+64))
																	v79 = F_Int64GetDatum(m, v78)
																	mBase = m.M
																	v80 = m.ExcPending
																	if v80 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+80)) = v79
																		v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
																		*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = v82
																		v84 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
																		v89 = F_heap_form_tuple(m, v84, v7+int32(48), v7+int32(32))
																		mBase = m.M
																		v90 = m.ExcPending
																		if v90 != 0 {
																			return int32(0)
																		} else {
																			v91 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
																			v92 = F_HeapTupleHeaderGetDatum(m, v91)
																			mBase = m.M
																			v93 = m.ExcPending
																			if v93 != 0 {
																				return int32(0)
																			} else {
																				v95 = v92
																				m.G0 = v7 + int32(96)
																				return v95
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v106 = m.ExcPending
					if v106 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_gin_metapage_info_8), int32(0))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(42), int32(_a_F_gin_metapage_info_4))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
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
