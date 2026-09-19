package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__brin_parallel_build_main(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
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
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int64
	_ = v167
	var v176 int64
	_ = v176
	var v178 int64
	_ = v178
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	v14 = F_shm_toc_lookup(m, l1, int64(-5764607523034234877), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[0])) = v14
		F_pgstat_report_activity(m, int32(3), v14)
		mBase = m.M
		v21 = F_shm_toc_lookup(m, l1, int64(-5764607523034234879), int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(v21)+24))
			v28 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[1]))
			if v28 == int32(0) {
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[2])))
				if v32&int32(1) == int32(0) {
				} else {
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v28)+392))
					if int32(1)&base.B2i32(v39 != int64(0)) != 0 {
					} else {
						v43 = int32(_a_F__brin_parallel_build_main_0)
						v45 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3]))
						v46 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3])) = v45 + v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v49 + v46
						*(*int64)(unsafe.Add(mBase, uint32(v28)+392)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v49 + int32(2)
						v60 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3]))
						*(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3])) = v60 - v46
					}
				}
			}
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v23 != 0 {
				v67 = int32(4)
			} else {
				v67 = int32(5)
			}
			v68 = F_table_open(m, v64, v67)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				if v23 != 0 {
					v73 = int32(3)
				} else {
					v73 = int32(8)
				}
				v74 = F_index_open(m, v70, v73)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v78 = F_palloc(m, int32(80))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						v80 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v80
						*(*int32)(unsafe.Add(mBase, uint32(v78))) = v74
						*(*int64)(unsafe.Add(mBase, uint32(v78)+16)) = v80
						v85 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v85
						*(*int32)(unsafe.Add(mBase, uint32(v78)+40)) = v85
						*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v85
						*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v76
						v92 = F_brin_build_desc(m, v74)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v78)+44)) = v92
							v95 = F_brin_new_memtuple(m, v92)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								v97 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v78)+72)) = v97
								v99 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v78)+64)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v78)+48)) = v95
								v103 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[4]))
								*(*int64)(unsafe.Add(mBase, uint32(v78)+52)) = v99
								*(*int32)(unsafe.Add(mBase, uint32(v78)+60)) = v103
								v107 = *(*int32)(unsafe.Add(mBase, uint32(v78)+28))
								v109 = base.I32_rem_u_s(int32(-2), v76)
								*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = v107 - v109 - int32(2)
								v116 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), v97)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_tuplesort_attach_shared(m, v116, l0)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return
									} else {
										base.MemoryCopy(m, int32(_a_F__brin_parallel_build_main_1), int32(_a_F__brin_parallel_build_main_2), int32(128))
										v126 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[5]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[6])) = v126
										v130 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[7]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[8])) = v130
										v134 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[9]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[10])) = v134
										v138 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[11]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[12])) = v138
										v141 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[13]))
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										v143 = base.I32_div_s(v141, v142)
										F__brin_parallel_scan_and_build(m, v78, v21, v116, v68, v74, v143)
										mBase = m.M
										v145 = m.ExcPending
										if v145 != 0 {
											return
										} else {
											v148 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												v152 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
												mBase = m.M
												v153 = m.ExcPending
												if v153 != 0 {
													return
												} else {
													v155 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[14]))
													v158 = v148 + v155<<(uint(int32(7))%32)
													v161 = v152 + v155<<(uint(int32(5))%32)
													base.MemoryFill(m, v158, int32(0), int32(128))
													F_BufferUsageAccumDiff(m, v158, int32(_a_F__brin_parallel_build_main_1))
													mBase = m.M
													v167 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v161)+24)) = v167
													*(*int64)(unsafe.Add(mBase, uint32(v161)+16)) = v167
													*(*int64)(unsafe.Add(mBase, uint32(v161)+8)) = v167
													*(*int64)(unsafe.Add(mBase, uint32(v161))) = v167
													v176 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[7]))
													v178 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[8]))
													*(*int64)(unsafe.Add(mBase, uint32(v161)+16)) = v176 - v178
													v182 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[11]))
													v184 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[12]))
													*(*int64)(unsafe.Add(mBase, uint32(v161))) = v182 - v184
													v188 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[9]))
													v190 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[10]))
													*(*int64)(unsafe.Add(mBase, uint32(v161)+8)) = v188 - v190
													v194 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[5]))
													v196 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[6]))
													*(*int64)(unsafe.Add(mBase, uint32(v161)+24)) = v194 - v196
													F_relation_close(m, v74, v73)
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return
													} else {
														F_relation_close(m, v68, v67)
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return
														} else {
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
func F_brin_bloom_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v112 int32
	_ = v112
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v204 int32
	_ = v204
	var v209 int64
	_ = v209
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v256 int64
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int64
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int64
	_ = v325
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v357 int64
	_ = v357
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	v2 = int32(0)
	v16 = int64(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		v27 = F_pg_detoast_datum(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
			v31 = int32(base.Ui32(v29) >> (uint(int32(3)) % 32))
			if v31 != 0 {
				v33 = v31 & int32(3)
				v34 = int32(16)
				v35 = v21 + v34
				v37 = v27 + v34
				v38 = int32(0)
				if base.Ui32(int32(32)) <= base.Ui32(v29) {
					v43 = v38
					v56 = v2
					for {
						v59 = v43 + v35
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v37))))
						v63 = v60 | v62
						*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
						v66 = v43 | int32(1)
						v67 = v35 + v66
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v37))))
						v71 = v68 | v70
						*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v71)
						v74 = v43 | int32(2)
						v75 = v35 + v74
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+v37))))
						v79 = v76 | v78
						*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v79)
						v82 = v43 | int32(3)
						v83 = v35 + v82
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v37))))
						v87 = v84 | v86
						*(*uint8)(unsafe.Add(mBase, uint32(v83))) = uint8(v87)
						v89 = int32(4)
						v90 = v43 + v89
						v92 = v56 + v89
						if v92 != v31&int32(536870908) {
							v43 = v90
							v56 = v92
							continue
						} else {
							break
						}
						break
					}
					if v33 == int32(0) {
					} else {
						v96 = v90
						v112 = v96
						v126 = v2
						for {
							v128 = v112 + v35
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
							v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v37))))
							v132 = v129 | v131
							*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v132)
							v134 = int32(1)
							v137 = v126 + v134
							if v137 != v33 {
								v112 = v112 + v134
								v126 = v137
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v96 = v38
					v112 = v96
					v126 = v2
					for {
						v128 = v112 + v35
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
						v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v37))))
						v132 = v129 | v131
						*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v132)
						v134 = int32(1)
						v137 = v126 + v134
						if v137 != v33 {
							v112 = v112 + v134
							v126 = v137
							continue
						} else {
							break
						}
						break
					}
				}
				v156 = v21 + int32(16)
				if base.Ui32(int32(32)) <= base.Ui32(v29) {
					v204 = v156
					v209 = int64(0)
					if base.B2i32(v204 != (v204+int32(3))&int32(-4))|base.B2i32(v31 < int32(4)) != 0 {
						v288 = v204
						v289 = v31
						v294 = v209
					} else {
						v219 = v31 - int32(4)
						v223 = int32(base.Ui32(v219)>>(uint(int32(2))%32)) + int32(1)
						v225 = v223 & int32(3)
						if base.Ui32(int32(12)) <= base.Ui32(v219) {
							v230 = v204
							v231 = v31
							v234 = int32(0)
							v236 = v209
							for {
								v237 = int32(16)
								v238 = v231 - v237
								v240 = v230 + v237
								v241 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
								v250 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
								v256 = base.I64_extend_i32_u(base.I32_popcnt(v241)) + (base.I64_extend_i32_u(base.I32_popcnt(v244)) + (base.I64_extend_i32_u(base.I32_popcnt(v247)) + (v236 + base.I64_extend_i32_u(base.I32_popcnt(v250)))))
								v258 = v234 + int32(4)
								if v258 != v223&int32(2147483644) {
									v230 = v240
									v231 = v238
									v234 = v258
									v236 = v256
									continue
								} else {
									break
								}
								break
							}
							if v225 == int32(0) {
								v288 = v240
								v289 = v238
								v294 = v256
							} else {
								v262 = v240
								v263 = v238
								v268 = v256
								v270 = v262
								v271 = v263
								v272 = int32(0)
								v276 = v268
								for {
									v277 = int32(4)
									v278 = v271 - v277
									v280 = v270 + v277
									v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
									v284 = v276 + base.I64_extend_i32_u(base.I32_popcnt(v281))
									v286 = v272 + int32(1)
									if v286 != v225 {
										v270 = v280
										v271 = v278
										v272 = v286
										v276 = v284
										continue
									} else {
										break
									}
									break
								}
								v288 = v280
								v289 = v278
								v294 = v284
							}
						} else {
							v262 = v204
							v263 = v31
							v268 = v209
							v270 = v262
							v271 = v263
							v272 = int32(0)
							v276 = v268
							for {
								v277 = int32(4)
								v278 = v271 - v277
								v280 = v270 + v277
								v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
								v284 = v276 + base.I64_extend_i32_u(base.I32_popcnt(v281))
								v286 = v272 + int32(1)
								if v286 != v225 {
									v270 = v280
									v271 = v278
									v272 = v286
									v276 = v284
									continue
								} else {
									break
								}
								break
							}
							v288 = v280
							v289 = v278
							v294 = v284
						}
					}
					if v289 == int32(0) {
						v357 = v294
					} else {
						v298 = v289 & int32(3)
						if v298 == int32(0) {
							v319 = v288
							v321 = v289
							v325 = v294
						} else {
							v302 = v288
							v304 = v289
							v306 = int32(0)
							v308 = v294
							for {
								v309 = int32(1)
								v310 = v302 + v309
								v312 = v304 - v309
								v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
								v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_brin_bloom_union[0]))))
								v315 = v308 + v314
								v317 = v306 + v309
								if v317 != v298 {
									v302 = v310
									v304 = v312
									v306 = v317
									v308 = v315
									continue
								} else {
									break
								}
								break
							}
							v319 = v310
							v321 = v312
							v325 = v315
						}
						if base.Ui32(v289) < base.Ui32(int32(4)) {
							v357 = v325
						} else {
							v328 = v319
							v330 = v321
							v334 = v325
							for {
								v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+3)))
								v336 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_c_F_brin_bloom_union[0]))))
								v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+2)))
								v338 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v337)+uint32(_c_F_brin_bloom_union[0]))))
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
								v340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_brin_bloom_union[0]))))
								v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
								v342 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v341)+uint32(_c_F_brin_bloom_union[0]))))
								v346 = v336 + (v338 + (v340 + (v334 + v342)))
								v347 = int32(4)
								v350 = v330 - v347
								if v350 != 0 {
									v328 = v328 + v347
									v330 = v350
									v334 = v346
									continue
								} else {
									break
								}
								break
							}
							v357 = v346
						}
					}
					v373 = v357
				} else {
					v160 = v156
					v161 = int32(0)
					v175 = v16
					for {
						v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
						v177 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v176)+uint32(_c_F_brin_bloom_union[0]))))
						v178 = v175 + v177
						v179 = int32(1)
						v182 = v161 + v179
						if v182 != v31 {
							v160 = v160 + v179
							v161 = v182
							v175 = v178
							continue
						} else {
							break
						}
						break
					}
					v373 = v178
				}
			} else {
				if base.Ui32(v29) < base.Ui32(int32(32)) {
					v373 = v16
				} else {
					v204 = v21 + int32(16)
					v209 = int64(0)
					if base.B2i32(v204 != (v204+int32(3))&int32(-4))|base.B2i32(v31 < int32(4)) != 0 {
						v288 = v204
						v289 = v31
						v294 = v209
					} else {
						v219 = v31 - int32(4)
						v223 = int32(base.Ui32(v219)>>(uint(int32(2))%32)) + int32(1)
						v225 = v223 & int32(3)
						if base.Ui32(int32(12)) <= base.Ui32(v219) {
							v230 = v204
							v231 = v31
							v234 = int32(0)
							v236 = v209
							for {
								v237 = int32(16)
								v238 = v231 - v237
								v240 = v230 + v237
								v241 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
								v247 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
								v250 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
								v256 = base.I64_extend_i32_u(base.I32_popcnt(v241)) + (base.I64_extend_i32_u(base.I32_popcnt(v244)) + (base.I64_extend_i32_u(base.I32_popcnt(v247)) + (v236 + base.I64_extend_i32_u(base.I32_popcnt(v250)))))
								v258 = v234 + int32(4)
								if v258 != v223&int32(2147483644) {
									v230 = v240
									v231 = v238
									v234 = v258
									v236 = v256
									continue
								} else {
									break
								}
								break
							}
							if v225 == int32(0) {
								v288 = v240
								v289 = v238
								v294 = v256
							} else {
								v262 = v240
								v263 = v238
								v268 = v256
								v270 = v262
								v271 = v263
								v272 = int32(0)
								v276 = v268
								for {
									v277 = int32(4)
									v278 = v271 - v277
									v280 = v270 + v277
									v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
									v284 = v276 + base.I64_extend_i32_u(base.I32_popcnt(v281))
									v286 = v272 + int32(1)
									if v286 != v225 {
										v270 = v280
										v271 = v278
										v272 = v286
										v276 = v284
										continue
									} else {
										break
									}
									break
								}
								v288 = v280
								v289 = v278
								v294 = v284
							}
						} else {
							v262 = v204
							v263 = v31
							v268 = v209
							v270 = v262
							v271 = v263
							v272 = int32(0)
							v276 = v268
							for {
								v277 = int32(4)
								v278 = v271 - v277
								v280 = v270 + v277
								v281 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
								v284 = v276 + base.I64_extend_i32_u(base.I32_popcnt(v281))
								v286 = v272 + int32(1)
								if v286 != v225 {
									v270 = v280
									v271 = v278
									v272 = v286
									v276 = v284
									continue
								} else {
									break
								}
								break
							}
							v288 = v280
							v289 = v278
							v294 = v284
						}
					}
					if v289 == int32(0) {
						v357 = v294
					} else {
						v298 = v289 & int32(3)
						if v298 == int32(0) {
							v319 = v288
							v321 = v289
							v325 = v294
						} else {
							v302 = v288
							v304 = v289
							v306 = int32(0)
							v308 = v294
							for {
								v309 = int32(1)
								v310 = v302 + v309
								v312 = v304 - v309
								v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
								v314 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v313)+uint32(_c_F_brin_bloom_union[0]))))
								v315 = v308 + v314
								v317 = v306 + v309
								if v317 != v298 {
									v302 = v310
									v304 = v312
									v306 = v317
									v308 = v315
									continue
								} else {
									break
								}
								break
							}
							v319 = v310
							v321 = v312
							v325 = v315
						}
						if base.Ui32(v289) < base.Ui32(int32(4)) {
							v357 = v325
						} else {
							v328 = v319
							v330 = v321
							v334 = v325
							for {
								v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+3)))
								v336 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_c_F_brin_bloom_union[0]))))
								v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+2)))
								v338 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v337)+uint32(_c_F_brin_bloom_union[0]))))
								v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328)+1)))
								v340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_brin_bloom_union[0]))))
								v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
								v342 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v341)+uint32(_c_F_brin_bloom_union[0]))))
								v346 = v336 + (v338 + (v340 + (v334 + v342)))
								v347 = int32(4)
								v350 = v330 - v347
								if v350 != 0 {
									v328 = v328 + v347
									v330 = v350
									v334 = v346
									continue
								} else {
									break
								}
								break
							}
							v357 = v346
						}
					}
					v373 = v357
				}
			}
			*(*uint32)(unsafe.Add(mBase, uint32(v21)+12)) = uint32(v373)
			v375 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
			if v21 != v376 {
				F_pfree(m, v376)
				mBase = m.M
				v379 = m.ExcPending
				if v379 != 0 {
					return int32(0)
				} else {
					v380 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v380))) = v21
					v382 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
					v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
					if v27 != v383 {
						F_pfree(m, v27)
						mBase = m.M
						v386 = m.ExcPending
						if v386 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						return int32(0)
					}
				}
			} else {
				v382 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
				if v27 != v383 {
					F_pfree(m, v27)
					mBase = m.M
					v386 = m.ExcPending
					if v386 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_brin_free_tuple(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_pfree(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_brin_metapage_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
		v15 = F_superuser(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v19 = F_verify_brin_page(m, v9, int32(_a_F_brin_metapage_info_0), int32(_a_F_brin_metapage_info_1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+14)))
					if v21 == int32(0) {
						v24 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
						v64 = int32(0)
						m.G0 = v6 + int32(48)
						return v64
					} else {
						v30 = F_get_call_result_type(m, l0, int32(0), v6+int32(44))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							if v30 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_brin_metapage_info_2), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_brin_metapage_info_3), int32(365), int32(_a_F_brin_metapage_info_4))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v34 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								v35 = F_BlessTupleDesc(m, v34)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v35
									v38 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v38
									v41 = F_psprintf(m, int32(_a_F_brin_metapage_info_5), v6)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = F_cstring_to_text(m, v41)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v43
											v46 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
											*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v46
											v48 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
											*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v48
											v50 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v19)+36)))
											v51 = F_Int64GetDatum(m, v50)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v51
												v54 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
												v59 = F_heap_form_tuple(m, v54, v6+int32(16), v6+int32(12))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
													v62 = F_HeapTupleHeaderGetDatum(m, v61)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return int32(0)
													} else {
														v64 = v62
														m.G0 = v6 + int32(48)
														return v64
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
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_brin_metapage_info_6), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_brin_metapage_info_3), int32(356), int32(_a_F_brin_metapage_info_4))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
func F_brin_minmax_consistent(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+6)))
	switch v21 - int32(1) {
	case 0, 1:
		v63 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, v21)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v67 = F_FunctionCall2Coll(m, v63, v14, v66, v18)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				v69 = v67
				m.G0 = v12 + int32(16)
				return v69
			}
		}
	case 2:
		v26 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, int32(2))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v32 = F_FunctionCall2Coll(m, v26, v14, v31, v18)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v32 == int32(0) {
					v69 = int32(0)
					m.G0 = v12 + int32(16)
					return v69
				} else {
					v37 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, int32(4))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v41 = F_FunctionCall2Coll(m, v37, v14, v40, v18)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v69 = v41
							m.G0 = v12 + int32(16)
							return v69
						}
					}
				}
			}
		}
	case 3, 4:
		v43 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, v21)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = F_FunctionCall2Coll(m, v43, v14, v46, v18)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v69 = v47
				m.G0 = v12 + int32(16)
				return v69
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+6)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v53
			F_errmsg_internal(m, int32(_a_F_brin_minmax_consistent_0), v12)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_brin_minmax_consistent_1), int32(195), int32(_a_F_brin_minmax_consistent_2))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
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
func F_brin_minmax_multi_opcinfo(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13853(m, l0, int32(_a_F_brin_minmax_multi_opcinfo_0), int32(188))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_brin_minmax_multi_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_brin_minmax_multi_options_0), int32(_a_F_brin_minmax_multi_options_1), int32(32), int32(8), int32(256))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_brin_minmax_multi_summary_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_brin_minmax_multi_summary_in_0), int32(2984), int32(_a_F_brin_minmax_multi_summary_in_1), int32(_a_F_brin_minmax_multi_summary_in_2), int32(_a_F_brin_minmax_multi_summary_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brin_range_deserialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = F_palloc0(m, l0<<(uint(int32(2))%32)+int32(36))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v40
	v43 = l1 + int32(20)
	v44 = F_get_typbyval(m, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v47 = F_get_typlen(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v51 = v21 + v22<<(uint(int32(1))%32)
	v53 = base.B2i32(v51 <= int32(0))
	if v53|v44 != 0 {
		v153 = v3
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v53 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L6:
	;
	v59 = int32(0)
	v67 = v59
	v69 = v3
	v70 = v43
	goto L7
L7:
	;
	if base.B2i32(v47 <= v59) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	if v138 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L9:
	;
	v143 = v67 + int32(1)
	if v143 < v51 {
		v67 = v143
		v69 = v138
		v70 = v139
		goto L7
	} else {
		goto L33
	}
L10:
	;
	v138 = v69 + (v47+int32(7))&int32(_a_F_brin_range_deserialize_0)
	v139 = v70
	goto L9
L11:
	;
	goto L12
L12:
	;
	switch v47&int32(_a_F_brin_range_deserialize_1) - int32(_a_F_brin_range_deserialize_2) {
	case 0:
		goto L15
	case 1:
		goto L16
	default:
		v138 = v69
		v139 = v70
		goto L9
	}
L13:
	;
	v138 = v69 + v134
	v139 = v131 + v70
	goto L9
L14:
	;
	v113 = int32(18)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	v117 = base.B2i32(v115 == v113)
	if v115 == v113 {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v103 = F_strlen(m, v70)
	mBase = m.M
	v138 = v69 + v103&int32(-8) + int32(8)
	v139 = v70 + v103 + int32(1)
	goto L9
L16:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v85 == int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if v85&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = int32(base.Ui32(v85) >> (uint(int32(1)) % 32))
	v131 = v91
	v134 = (v91 + int32(7)) & int32(248)
	goto L13
L19:
	;
	goto L20
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v98 = int32(base.Ui32(v96) >> (uint(int32(2)) % 32))
	v131 = v98
	v134 = (v98 + int32(7)) & int32(2147483640)
	goto L13
L21:
	;
	v118 = v113
	goto L23
L22:
	;
	v118 = int32(2)
	goto L23
L23:
	;
	v124 = base.B2i32(base.Ui32((v115-int32(1))&int32(255)) < base.Ui32(int32(3)))
	if base.Ui32((v115-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v125 = int32(6)
	goto L26
L25:
	;
	v125 = v118
	goto L26
L26:
	;
	v126 = int32(8)
	if v115 == v113 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v129 = int32(24)
	goto L29
L28:
	;
	v129 = v126
	goto L29
L29:
	;
	if base.Ui32((v115-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v130 = v126
	goto L32
L31:
	;
	v130 = v129
	goto L32
L32:
	;
	v131 = v125
	v134 = v130
	goto L13
L33:
	;
	goto L8
L34:
	;
	v153 = int32(0)
	goto L5
L35:
	;
	goto L36
L36:
	;
	v148 = F_palloc(m, v138)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v153 = v148
	goto L5
L38:
	;
	v169 = v27 + int32(36)
	v175 = v47 & int32(_a_F_brin_range_deserialize_1)
	v179 = v43
	v180 = int32(0)
	v182 = v153
	goto L41
L39:
	;
	goto L40
L40:
	;
	m.G0 = v19 + int32(16)
	return v27
L41:
	;
	if v44 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L40
L43:
	;
	v326 = v180 + int32(1)
	if v326 != v51 {
		v179 = v321
		v180 = v326
		v182 = v323
		goto L41
	} else {
		goto L100
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(0)
	if v47 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	if int32(0) < v47 {
		goto L58
	} else {
		goto L59
	}
L47:
	;
	base.MemoryCopy(m, v19+int32(12), v179, v47)
	goto L49
L48:
	;
	goto L49
L49:
	;
	switch v175 - int32(1) {
	case 0:
		goto L51
	case 1:
		goto L54
	default:
		goto L52
	case 3:
		goto L53
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169+v180<<(uint(int32(2))%32)))) = v219
	v321 = v179 + v47
	v323 = v182
	goto L43
L51:
	;
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+12)))
	v219 = v218
	goto L50
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v219 = v204
	goto L50
L54:
	;
	v203 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+12)))
	v219 = v203
	goto L50
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v47
	F_errmsg_internal(m, int32(_a_F_brin_range_deserialize_3), v19)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_brin_range_deserialize_4), int32(70), int32(_a_F_brin_range_deserialize_5))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169+v180<<(uint(int32(2))%32)))) = v182
	if v47 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	switch v175 - int32(_a_F_brin_range_deserialize_2) {
	case 0:
		goto L65
	case 1:
		goto L66
	default:
		v321 = v179
		v323 = v182
		goto L43
	}
L61:
	;
	base.MemoryCopy(m, v182, v179, v47)
	goto L63
L62:
	;
	goto L63
L63:
	;
	v321 = v179 + v47
	v323 = v182 + (v47+int32(7))&int32(_a_F_brin_range_deserialize_0)
	goto L43
L64:
	;
	v300 = int32(8)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v305 = base.B2i32(v303 == int32(18))
	if v303 == int32(18) {
		goto L88
	} else {
		goto L89
	}
L65:
	;
	v286 = F_strlen(m, v179)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v169+v180<<(uint(int32(2))%32)))) = v182
	v292 = v286 + int32(1)
	if v292 != 0 {
		goto L85
	} else {
		goto L86
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169+v180<<(uint(int32(2))%32)))) = v182
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v237 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v262 != 0 {
		goto L78
	} else {
		goto L79
	}
L68:
	;
	v241 = int32(18)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	if v243 == v241 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	v254 = int32(1)
	if v237&v254 != 0 {
		v262 = int32(base.Ui32(v237) >> (uint(v254) % 32))
		goto L67
	} else {
		goto L77
	}
L71:
	;
	v246 = v241
	goto L73
L72:
	;
	v246 = int32(2)
	goto L73
L73:
	;
	if base.Ui32((v243-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v253 = int32(6)
	goto L76
L75:
	;
	v253 = v246
	goto L76
L76:
	;
	v262 = v253
	goto L67
L77:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v262 = int32(base.Ui32(v258) >> (uint(int32(2)) % 32))
	goto L67
L78:
	;
	base.MemoryCopy(m, v182, v179, v262)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v264 == int32(1) {
		goto L64
	} else {
		goto L81
	}
L81:
	;
	if v264&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v270 = int32(base.Ui32(v264) >> (uint(int32(1)) % 32))
	v321 = v179 + v270
	v323 = v182 + (v270+int32(7))&int32(248)
	goto L43
L83:
	;
	goto L84
L84:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v279 = int32(base.Ui32(v277) >> (uint(int32(2)) % 32))
	v321 = v179 + v279
	v323 = v182 + (v279+int32(7))&int32(2147483640)
	goto L43
L85:
	;
	base.MemoryCopy(m, v182, v179, v292)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v321 = v179 + v292
	v323 = v182 + v286&int32(-8) + int32(8)
	goto L43
L88:
	;
	v306 = int32(24)
	goto L90
L89:
	;
	v306 = v300
	goto L90
L90:
	;
	v312 = base.B2i32(base.Ui32((v303-int32(1))&int32(255)) < base.Ui32(int32(3)))
	if base.Ui32((v303-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v313 = v300
	goto L93
L92:
	;
	v313 = v306
	goto L93
L93:
	;
	if v303 == int32(18) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v318 = int32(18)
	goto L96
L95:
	;
	v318 = int32(2)
	goto L96
L96:
	;
	if base.Ui32((v303-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v319 = int32(6)
	goto L99
L98:
	;
	v319 = v318
	goto L99
L99:
	;
	v321 = v179 + v319
	v323 = v182 + v313
	goto L43
L100:
	;
	goto L42
}
