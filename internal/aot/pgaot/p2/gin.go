package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	return v89
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
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v7)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v79
	v81 = int32(8)
	v83 = int32(0)
	v86 = F_ExtendBufferedRel(m, v7+v81, v83, v83, v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
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
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L25
	}
L14:
	;
	if v60 != 0 {
		v89 = v19
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
	F_UnlockBuffer(m, v19)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	v66 = F_GetFreeIndexPage(m, l0)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	if v66 != int32(-1) {
		v16 = v66
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L8
L28:
	;
	v89 = v86
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
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v140 int64
	_ = v140
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
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
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[8]))) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[9]))) = v45
						v60 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[10]))
						v65 = F_AllocSetContextCreateInternal(m, v60, int32(_a_F__gin_parallel_build_main_1), v41, int32(_a_F__gin_parallel_build_main_2), int32(_a_F__gin_parallel_build_main_3))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[11]))) = v65
							v69 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[10]))
							v74 = F_AllocSetContextCreateInternal(m, v69, int32(_a_F__gin_parallel_build_main_4), int32(0), int32(_a_F__gin_parallel_build_main_2), int32(_a_F__gin_parallel_build_main_3))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[12]))) = v74
								*(*int32)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F__gin_parallel_build_main[13]))) = v12
								F_ginInitBA(m, v12+int32(_a_F__gin_parallel_build_main_5))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									v84 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_tuplesort_attach_shared(m, v84, l0)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											base.MemoryCopy(m, int32(_a_F__gin_parallel_build_main_6), int32(_a_F__gin_parallel_build_main_7), int32(128))
											v94 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[14]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[15])) = v94
											v98 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[16]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[17])) = v98
											v102 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[18]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[19])) = v102
											v106 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[20]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[21])) = v106
											v110 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[22]))
											*(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[23])) = v110
											v113 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[24]))
											v114 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
											v115 = base.I32_div_s(v113, v114)
											F__gin_parallel_scan_and_build(m, v12, v24, v84, v31, v37, v115, int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												v121 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return
												} else {
													v125 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[25]))
														v131 = v121 + v128<<(uint(int32(7))%32)
														v134 = v125 + v128*int32(40)
														base.MemoryFill(m, v131, int32(0), int32(128))
														F_BufferUsageAccumDiff(m, v131, int32(_a_F__gin_parallel_build_main_6))
														mBase = m.M
														v140 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v134)+32)) = v140
														*(*int64)(unsafe.Add(mBase, uint32(v134)+24)) = v140
														*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v140
														*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v140
														*(*int64)(unsafe.Add(mBase, uint32(v134))) = v140
														v151 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[18]))
														v153 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[19]))
														*(*int64)(unsafe.Add(mBase, uint32(v134)+16)) = v151 - v153
														v157 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[22]))
														v159 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[23]))
														*(*int64)(unsafe.Add(mBase, uint32(v134))) = v157 - v159
														v163 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[20]))
														v165 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[21]))
														*(*int64)(unsafe.Add(mBase, uint32(v134)+8)) = v163 - v165
														v169 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[16]))
														v171 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[17]))
														*(*int64)(unsafe.Add(mBase, uint32(v134)+24)) = v169 - v171
														v175 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[14]))
														v177 = *(*int64)(unsafe.Add(mBase, _c_F__gin_parallel_build_main[15]))
														*(*int64)(unsafe.Add(mBase, uint32(v134)+32)) = v175 - v177
														F_relation_close(m, v37, v36)
														mBase = m.M
														v181 = m.ExcPending
														if v181 != 0 {
															return
														} else {
															F_relation_close(m, v31, v30)
															mBase = m.M
															v183 = m.ExcPending
															if v183 != 0 {
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v201 int64
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int64
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
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
	var v245 int32
	_ = v245
	var v246 float64
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int64
	_ = v257
	var v259 int32
	_ = v259
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 float64
	_ = v410
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 float64
	_ = v454
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 float64
	_ = v469
	var v470 float64
	_ = v470
	var v473 float64
	_ = v473
	var v474 float64
	_ = v474
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
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
	v50 = F_table_beginscan_parallel(m, l3, l1-int32(-64), v43)
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
	v54 = m.T0[v53].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l3, l4, v38, int32(1), v43, l6, v43, int32(-1), int32(57), l0, v50)
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
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
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
	if l6 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v160 = F_tuplesort_getgintuple(m, v58, v15+int32(12))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L24
	}
L11:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[4]))
	if v66 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
	F_tuplesort_performsort(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L23
	}
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[3])))
	F_tuplesort_performsort(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[5])))
	if v70&int32(1) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v75 = int32(_a_F__gin_parallel_scan_and_build_0)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6]))
	v78 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])) = v77 + v78
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v81 + v78
	v85 = int32(0)
	v87 = int32(_a_F__gin_parallel_scan_and_build_1)
	v88 = base.AtomicRmwOr32(m, v85, v87, v85)
	*(*int64)(unsafe.Add(mBase, uint32(v66+int32(80))+232)) = int64(3)
	v96 = base.AtomicRmwOr32(m, v85, v87, v85)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v97 + v78
	v103 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])) = v103 - v78
	goto L15
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[4]))
	if v114 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L10
L20:
	;
	goto L19
L21:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[5])))
	if v118&int32(1) == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v123 = int32(_a_F__gin_parallel_scan_and_build_0)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6]))
	v126 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])) = v125 + v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v129 + v126
	v133 = int32(0)
	v135 = int32(_a_F__gin_parallel_scan_and_build_1)
	v136 = base.AtomicRmwOr32(m, v133, v135, v133)
	*(*int64)(unsafe.Add(mBase, uint32(v114+int32(80))+232)) = int64(4)
	v144 = base.AtomicRmwOr32(m, v133, v135, v133)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v145 + v126
	v151 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[6])) = v151 - v126
	goto L20
L23:
	;
	goto L10
L24:
	;
	if v160 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v163 = v60 + int32(8)
	v167 = v160
	goto L28
L26:
	;
	goto L27
L27:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v395 != 0 {
		goto L88
	} else {
		goto L89
	}
L28:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F__gin_parallel_scan_and_build[7]))
	if v177 != 0 {
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v180 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v275 <= int32(0) {
		v293 = v275
		goto L62
	} else {
		goto L63
	}
L35:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+4)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	if v184 != v185 {
		v226 = v185
		v228 = v183
		v229 = v180
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+20)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	v240 = F__gin_build_tuple(m, v226&int32(_a_F__gin_parallel_scan_and_build_2), v228&int32(255), v234, v235, v236, v237, v229, v15+int32(8))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L54
	}
L37:
	;
	v188 = v183 & int32(255)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+15)))
	if v188 != v189 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v226 = v184
	v228 = v183
	v229 = v180
	goto L36
L39:
	;
	goto L40
L40:
	;
	if v188 != 0 {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v192 = v167 + int32(24)
	v193 = int32(1)
	v195 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	if v196 == v193 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v60)+36))
	v203 = int32(36)
	v205 = v202 + v184*v203
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205-int32(20))))
	v211 = m.T0[v210].(func(*base.Module, int64, int64, int32) int32)(m, v195, v201, v205-v203)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L46
	}
L43:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v192)))
	v201 = v199
	goto L42
L44:
	;
	goto L45
L45:
	;
	v201 = base.I64_extend_i32_u(v192)
	goto L42
L46:
	;
	if v211 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v216 = v193
	goto L49
L48:
	;
	v216 = int32(0) - v211
	goto L49
L49:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205-int32(28)))))
	if v219 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v220 = v216
	goto L52
L51:
	;
	v220 = v211
	goto L52
L52:
	;
	if v220 == int32(0) {
		goto L34
	} else {
		goto L53
	}
L53:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v226 = v225
	v228 = v224
	v229 = v223
	goto L36
L54:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_tuplesort_putgintuple(m, v242, v240, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[8])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[8]))) = base.F64_add(v246, float64(1))
	F_pfree(m, v240)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v252 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v257 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+28)) = v257
	v259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v259)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v259)
	*(*int64)(unsafe.Add(mBase, uint32(v163)+7)) = v257
	*(*int64)(unsafe.Add(mBase, uint32(v163))) = v257
	goto L34
L58:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	if v253 != 0 {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	F_pfree(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L57
L61:
	;
	if v339 != 0 {
		goto L76
	} else {
		goto L77
	}
L62:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	if v293 <= v294 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	v279 = int32(6)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v290 = F_itemptr_comparator(m, v278+v275*v279-v279, (v167+v284+int32(25))&int32(-2))
	mBase = m.M
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v290 != 0 {
		v293 = v291
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v291
	v293 = v291
	goto L62
L65:
	;
	if int32(1024) <= v327 {
		goto L73
	} else {
		goto L74
	}
L66:
	;
	v327 = v294
	goto L65
L67:
	;
	goto L68
L68:
	;
	v301 = v294
	goto L69
L69:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v314 = F_itemptr_comparator(m, v304+v301*int32(6), (v167+int32(24)+v308+int32(1))&int32(-2))
	mBase = m.M
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	if int32(0) < v314 {
		v327 = v315
		goto L65
	} else {
		goto L71
	}
L70:
	;
	v327 = v319
	goto L65
L71:
	;
	v318 = int32(1)
	v319 = v315 + v318
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v319
	v322 = v301 + v318
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v322 < v323 {
		v301 = v322
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	v339 = base.B2i32(v333 <= v334+v335)
	goto L75
L74:
	;
	v339 = int32(0)
	goto L75
L75:
	;
	goto L61
L76:
	;
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	v343 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+20)))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v349 = F__gin_build_tuple(m, v340, v341, v342, v343, v344, v345, v346, v15+int32(8))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_GinBufferStoreTuple(m, v60, v167)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L85
	}
L79:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_tuplesort_putgintuple(m, v351, v349, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v349)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v361 = (v357 - v358) * int32(6)
	if v361 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	base.MemoryCopy(m, v362, v362+v358*int32(6), v361)
	goto L84
L83:
	;
	goto L84
L84:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = int32(0)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v60)+28)) = v371 - v368
	goto L78
L85:
	;
	v381 = F_tuplesort_getgintuple(m, v58, v15+int32(12))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v381 != 0 {
		v167 = v381
		goto L28
	} else {
		goto L87
	}
L87:
	;
	goto L29
L88:
	;
	v396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60))))
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	v399 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+20)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	v404 = F__gin_build_tuple(m, v396, v397, v398, v399, v400, v401, v395, v15+int32(8))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	if v434 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L91:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_tuplesort_putgintuple(m, v406, v404, v407)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v410 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[8])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[8]))) = base.F64_add(v410, float64(1))
	F_pfree(m, v404)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v417 = v60 + int32(8)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v418 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v423 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+28)) = v423
	v425 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v425)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v425)
	*(*int64)(unsafe.Add(mBase, uint32(v417)+7)) = v423
	*(*int64)(unsafe.Add(mBase, uint32(v417))) = v423
	goto L90
L95:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	if v419 != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	F_pfree(m, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	F_pfree(m, v60)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L105
	}
L99:
	;
	F_pfree(m, v434)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v439 == int32(0) {
		goto L98
	} else {
		goto L101
	}
L101:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
	if v442 != 0 {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	if v443 != 0 {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_pfree(m, v444)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	F_tuplesort_end(m, v58)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	F_tuplesort_performsort(m, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v454 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9])))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9]))) = base.F64_add(v54, v454)
	v459 = base.AtomicRmwXchg32(m, l1, int32(28), int32(1))
	if v459 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_s_lock(m, l1+int32(28), int32(_a_F__gin_parallel_scan_and_build_3))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v465 + int32(1)
	v469 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[9])))
	v470 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+40)) = base.F64_add(v469, v470)
	v473 = *(*float64)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[8])))
	v474 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l1)+48)) = base.F64_add(v473, v474)
	v477 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l1)+28)), uint32(v477))
	F_ConditionVariableSignal(m, l1+int32(16))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F__gin_parallel_scan_and_build[2])))
	F_tuplesort_end(m, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	m.G0 = v15 + int32(16)
	return
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
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
	F_UnlockBuffer(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
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
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L30
	}
L21:
	;
	v67 = l1
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v73 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L16
L24:
	;
	F_ReleaseBuffer(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_pfree(m, v67)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	if v72 != 0 {
		v67 = v72
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
		if l2 == int32(0) {
			F_UnlockBuffer(m, v28)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_UnlockReleaseBuffer(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v28 < int32(0) {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[0]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v28^int32(-1))<<(uint(int32(2))%32))))
						v57 = v49
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[1]))
						v57 = v51 + v28<<(uint(int32(13))%32) + int32(-8192)
					}
					v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
					v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+v57)+6)))
					if (v60^v26)&int32(3) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_ginStepRight_0), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ginStepRight_1), int32(192), int32(_a_F_ginStepRight_2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
		} else {
			F_LockBufferInternal(m, v28, l2)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_UnlockReleaseBuffer(m, l0)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					if v28 < int32(0) {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[0]))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v28^int32(-1))<<(uint(int32(2))%32))))
						v57 = v49
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, _c_F_ginStepRight[1]))
						v57 = v51 + v28<<(uint(int32(13))%32) + int32(-8192)
					}
					v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
					v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58+v57)+6)))
					if (v60^v26)&int32(3) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_ginStepRight_0), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ginStepRight_1), int32(192), int32(_a_F_ginStepRight_2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
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
}
func F_gin_btree_extract_query(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int64
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v22 = base.I32_wrap_i64(v21)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = F_palloc(m, int32(8))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return int64(0)
	} else {
		v35 = F_palloc(m, int32(32))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int64(0)
		} else {
			v38 = F_palloc(m, int32(1))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int64(0)
			} else {
				v41 = v22 & int32(_a_F_gin_btree_extract_query_0)
				v43 = int32(base.Ui32(v41) >> (uint(int32(4)) % 32))
				v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v43))))
				if v45 == int32(1) {
					v49 = F_pg_detoast_datum(m, base.I32_wrap_i64(v28))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int64(0)
					} else {
						v52 = base.I64_extend_i32_u(v49)
						v53 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = v53
						*(*int32)(unsafe.Add(mBase, uint32(v26))) = v38
						v56 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v56)
						switch v22&int32(15) - v53 {
						case 0, 1:
							v90 = m.T0[l1].(func(*base.Module) int64)(m)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v90
								v93 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v93)
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
								*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
								v107 = F_palloc(m, int32(4))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
									m.G0 = v19 + int32(16)
									return base.I64_extend_i32_u(v30)
								}
							}
						case 2:
							if l3 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
								*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
								v107 = F_palloc(m, int32(4))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
									m.G0 = v19 + int32(16)
									return base.I64_extend_i32_u(v30)
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
								if v67 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
									*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
									*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
									*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
									v107 = F_palloc(m, int32(4))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
										m.G0 = v19 + int32(16)
										return base.I64_extend_i32_u(v30)
									}
								} else {
									v70 = m.T0[v67].(func(*base.Module, int64) int64)(m, v52)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v30))) = v70
										v73 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v73)
										*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
										*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
										v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
										*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
										v107 = F_palloc(m, int32(4))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
											*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
											m.G0 = v19 + int32(16)
											return base.I64_extend_i32_u(v30)
										}
									}
								}
							}
						case 3, 4:
							v60 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v60)
							if l3 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
								*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
								v107 = F_palloc(m, int32(4))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
									m.G0 = v19 + int32(16)
									return base.I64_extend_i32_u(v30)
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
								if v67 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
									*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
									*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
									*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
									v107 = F_palloc(m, int32(4))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
										m.G0 = v19 + int32(16)
										return base.I64_extend_i32_u(v30)
									}
								} else {
									v70 = m.T0[v67].(func(*base.Module, int64) int64)(m, v52)
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v30))) = v70
										v73 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v73)
										*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
										*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
										v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
										*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
										v107 = F_palloc(m, int32(4))
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
											*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
											m.G0 = v19 + int32(16)
											return base.I64_extend_i32_u(v30)
										}
									}
								}
							}
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
								F_errmsg_internal(m, int32(_a_F_gin_btree_extract_query_1), v19)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_gin_btree_extract_query_2), int32(135), int32(_a_F_gin_btree_extract_query_3))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int64(0)
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
					v52 = v28
					v53 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = v53
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v38
					v56 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v56)
					switch v22&int32(15) - v53 {
					case 0, 1:
						v90 = m.T0[l1].(func(*base.Module) int64)(m)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int64(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v30))) = v90
							v93 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v93)
							*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
							*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
							*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
							v107 = F_palloc(m, int32(4))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
								m.G0 = v19 + int32(16)
								return base.I64_extend_i32_u(v30)
							}
						}
					case 2:
						if l3 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
							*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
							*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
							*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
							v107 = F_palloc(m, int32(4))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
								m.G0 = v19 + int32(16)
								return base.I64_extend_i32_u(v30)
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
							if v67 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
								*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
								v107 = F_palloc(m, int32(4))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
									m.G0 = v19 + int32(16)
									return base.I64_extend_i32_u(v30)
								}
							} else {
								v70 = m.T0[v67].(func(*base.Module, int64) int64)(m, v52)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v30))) = v70
									v73 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
									*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
									*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
									v107 = F_palloc(m, int32(4))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
										m.G0 = v19 + int32(16)
										return base.I64_extend_i32_u(v30)
									}
								}
							}
						}
					case 3, 4:
						v60 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v60)
						if l3 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
							*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
							*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
							*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
							v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
							v107 = F_palloc(m, int32(4))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
								m.G0 = v19 + int32(16)
								return base.I64_extend_i32_u(v30)
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
							if v67 == int32(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v52
								*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
								*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
								*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
								v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
								v107 = F_palloc(m, int32(4))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
									*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
									m.G0 = v19 + int32(16)
									return base.I64_extend_i32_u(v30)
								}
							} else {
								v70 = m.T0[v67].(func(*base.Module, int64) int64)(m, v52)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v30))) = v70
									v73 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v73)
									*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v52
									*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v21)
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
									*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v98
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l4+v43<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v35)+24)) = v103
									v107 = F_palloc(m, int32(4))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v25)))) = v107
										*(*int32)(unsafe.Add(mBase, uint32(v107))) = v35
										m.G0 = v19 + int32(16)
										return base.I64_extend_i32_u(v30)
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v19))) = v41
							F_errmsg_internal(m, int32(_a_F_gin_btree_extract_query_1), v19)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_gin_btree_extract_query_2), int32(135), int32(_a_F_gin_btree_extract_query_3))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int64(0)
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
func F_gin_consistent_jsonb(m *base.Module, l0 int32) int64 {
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
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = base.I32_wrap_i64(v15)
	switch v16&int32(_a_F_gin_consistent_jsonb_0) - int32(7) {
	case 0:
		goto L7
	default:
		goto L4
	case 2:
		goto L3
	case 3:
		goto L6
	case 4:
		goto L5
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v99
L2:
	;
	v99 = int64(1)
	goto L1
L3:
	;
	v95 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v95)
	goto L2
L4:
	;
	if base.Ui32((v16-int32(15))&int32(_a_F_gin_consistent_jsonb_0)) <= base.Ui32(int32(1)) {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	v42 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v42)
	v44 = int32(0)
	v45 = int64(1)
	if v13 <= v44 {
		v99 = v45
		goto L1
	} else {
		goto L15
	}
L6:
	;
	v40 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v40)
	goto L2
L7:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v21)
	v23 = int32(0)
	v24 = int64(1)
	if v13 <= v23 {
		v99 = v24
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v27 = v23
	goto L9
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27+v14))))
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v99 = int64(0)
	goto L1
L11:
	;
	v37 = v27 + int32(1)
	if v13 != v37 {
		v27 = v37
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L10
L14:
	;
	v99 = v24
	goto L1
L15:
	;
	v48 = v44
	goto L16
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v14))))
	if v56 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = int64(0)
	goto L1
L18:
	;
	v58 = v48 + int32(1)
	if v13 != v58 {
		v48 = v58
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v99 = v45
	goto L1
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v68 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v68)
	if v13 <= int32(0) {
		goto L2
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L26
	} else {
		goto L28
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v73 = F_execute_jsp_gin_node(m, v72, v14)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	return int64(0)
L27:
	;
	v99 = base.I64_extend_i32_u(base.B2i32(v73 != int32(0)))
	goto L1
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16 & int32(_a_F_gin_consistent_jsonb_0)
	F_errmsg_internal(m, int32(_a_F_gin_consistent_jsonb_1), v10)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_gin_consistent_jsonb_2), int32(1008), int32(_a_F_gin_consistent_jsonb_3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gin_extract_hstore(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
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
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
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
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
L4:
	;
	goto L5
L5:
	;
	v28 = v12 + int32(8)
	v31 = v28 + v19<<(uint(int32(3))%32)
	v35 = F_palloc(m, v19<<(uint(int32(4))%32))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v37 = int32(0)
	goto L7
L7:
	;
	v49 = v28 + v37<<(uint(int32(3))%32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return base.I64_extend_i32_u(v35)
L9:
	;
	v68 = v64 + int32(5)
	v69 = F_palloc(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L10:
	;
	v64 = v50 & int32(1073741823)
	v66 = v31
	goto L9
L11:
	;
	goto L12
L12:
	;
	v55 = int32(1073741823)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49-int32(4))))
	v61 = v59 & v55
	v64 = v50&v55 - v61
	v66 = v61 + v31
	goto L9
L13:
	;
	v71 = int32(75)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+4)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v68 << (uint(int32(2)) % 32)
	v76 = int32(0)
	if base.B2i32(v64 == v76)|base.B2i32(v64 <= v76) == v76 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	base.MemoryCopy(m, v69+int32(5), v66, v64)
	goto L16
L15:
	;
	goto L16
L16:
	;
	v86 = int32(1)
	v87 = v37 << (uint(v86) % 32)
	*(*int64)(unsafe.Add(mBase, uint32(v35+v87<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v69)
	v94 = v87 | v86
	v97 = v28 + v94<<(uint(int32(2))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98&int32(1073741824) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35+v94<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v139)
	v149 = v37 + int32(1)
	if v149 != v19 {
		v37 = v149
		goto L7
	} else {
		goto L28
	}
L18:
	;
	v102 = F_palloc(m, int32(5))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v98 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v104 = int32(78)
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+4)) = uint8(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = int32(20)
	v139 = v102
	goto L17
L22:
	;
	v123 = v120 + int32(5)
	v124 = F_palloc(m, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v120 = v98 & int32(1073741823)
	v121 = v31
	goto L22
L24:
	;
	goto L25
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v97-int32(4))))
	v116 = v114 & int32(1073741823)
	v120 = v98 - v116
	v121 = v116 + v31
	goto L22
L26:
	;
	v126 = int32(86)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+4)) = uint8(v126)
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v123 << (uint(int32(2)) % 32)
	v131 = int32(0)
	if base.B2i32(v120 == v131)|base.B2i32(v120 <= v131) != 0 {
		v139 = v124
		goto L17
	} else {
		goto L27
	}
L27:
	;
	base.MemoryCopy(m, v124+int32(5), v121, v120)
	v139 = v124
	goto L17
L28:
	;
	goto L8
}
func F_gin_extract_jsonb_query_path(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = base.I32_wrap_i64(v12)
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = base.I32_wrap_i64(v14)
	if v15&int32(_a_F_gin_extract_jsonb_query_path_0) == int32(7) {
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v25 = F_DirectFunctionCall2Coll(m, int32(1465), int32(0), v22, v12&int64(4294967295))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = base.I32_wrap_i64(v25)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			if v30 == int32(0) {
				v49 = v29
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(2)
				v52 = v49
			} else {
				v52 = v29
			}
			m.G0 = v9 + int32(16)
			return base.I64_extend_i32_u(v52)
		}
	} else {
		if base.Ui32(int32(1)) < base.Ui32((v15-int32(15))&int32(_a_F_gin_extract_jsonb_query_path_0)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v15 & int32(_a_F_gin_extract_jsonb_query_path_0)
				F_errmsg_internal(m, int32(_a_F_gin_extract_jsonb_query_path_1), v9)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_gin_extract_jsonb_query_path_2), int32(1213), int32(_a_F_gin_extract_jsonb_query_path_3))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v40 = F_pg_detoast_datum(m, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				v46 = F_extract_jsp_query(m, v40, v15&int32(_a_F_gin_extract_jsonb_query_path_0), int32(1), v13, v45)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int64(0)
				} else {
					if v46 != 0 {
						v52 = v46
					} else {
						v49 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(2)
						v52 = v49
					}
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v52)
				}
			}
		}
	}
}
func F_gin_extract_query_bool(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_bool_0), int32(_a_F_gin_extract_query_bool_1), int32(0), int32(_a_F_gin_extract_query_bool_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_char(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_char_0), int32(_a_F_gin_extract_query_char_1), int32(0), int32(_a_F_gin_extract_query_char_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_float4(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_float4_0), int32(_a_F_gin_extract_query_float4_1), int32(_a_F_gin_extract_query_float4_2), int32(_a_F_gin_extract_query_float4_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_float8(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_float8_0), int32(_a_F_gin_extract_query_float8_1), int32(_a_F_gin_extract_query_float8_2), int32(_a_F_gin_extract_query_float8_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_macaddr8(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_macaddr8_0), int32(_a_F_gin_extract_query_macaddr8_1), int32(0), int32(_a_F_gin_extract_query_macaddr8_2))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_extract_query_timestamp(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_gin_btree_extract_query(m, l0, int32(_a_F_gin_extract_query_timestamp_0), int32(_a_F_gin_extract_query_timestamp_1), int32(_a_F_gin_extract_query_timestamp_2), int32(_a_F_gin_extract_query_timestamp_3))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_gin_metapage_info(m *base.Module, l0 int32) int64 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = F_superuser(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if v15 != 0 {
				v17 = F_get_page_from_raw(m, v11)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+14)))
					if v19 == int32(0) {
						v22 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
						v80 = int64(0)
						m.G0 = v8 + int32(144)
						return v80
					} else {
						v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
						v26 = int32(8)
						v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
						if (v25<<(uint(v26)%32)-v28)&int32(_a_F_gin_metapage_info_0) != v26 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_gin_metapage_info_1), int32(0))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
										v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+19)))
										v114 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v114
										*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = (v113<<(uint(v114)%32) - v112) & int32(_a_F_gin_metapage_info_0)
										v125 = F_errdetail(m, int32(_a_F_gin_metapage_info_2), v8+int32(16))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(55), int32(_a_F_gin_metapage_info_4))
											mBase = m.M
											v131 = m.ExcPending
											if v131 != 0 {
												return int64(0)
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
							v34 = v17 + v28
							v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
							if v35 != int32(8) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_gin_metapage_info_5), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34)+6)))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v143
											v148 = F_errdetail(m, int32(_a_F_gin_metapage_info_6), v8)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(64), int32(_a_F_gin_metapage_info_4))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int64(0)
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
								v41 = F_get_call_result_type(m, l0, int32(0), v8+int32(140))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int64(0)
								} else {
									if v41 != int32(1) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return int64(0)
										} else {
											F_errmsg_internal(m, int32(_a_F_gin_metapage_info_7), int32(0))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(68), int32(_a_F_gin_metapage_info_4))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									} else {
										v45 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v8)+40)) = uint16(v45)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(0)
										v49 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+24)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v49
										v51 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+28)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v51
										v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+32)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v53
										v55 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+36)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+72)) = v55
										v57 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v57
										v59 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+48)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+88)) = v59
										v61 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+52)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v61
										v63 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+56)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+104)) = v63
										v65 = *(*int64)(unsafe.Add(mBase, uint32(v17)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+112)) = v65
										v67 = int64(*(*int32)(unsafe.Add(mBase, uint32(v17)+72)))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+120)) = v67
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+140))
										v74 = F_heap_form_tuple(m, v69, v8+int32(48), v8+int32(32))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
											v77 = F_HeapTupleHeaderGetDatum(m, v76)
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int64(0)
											} else {
												v80 = v77
												m.G0 = v8 + int32(144)
												return v80
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
				v88 = m.ExcPending
				if v88 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_gin_metapage_info_8), int32(0))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_gin_metapage_info_3), int32(42), int32(_a_F_gin_metapage_info_4))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int64(0)
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
