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
	var v37 int64
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
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
	var v200 int64
	_ = v200
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	v14 = F_shm_toc_lookup(m, l1, int64(-5764607523034234877), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[5])) = v14
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
			v28 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			if v28 == int32(0) {
			} else {
				v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[10])))
				if v32 != int32(1) {
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(v28)+392))
					if int32(1)&base.B2i32(v37 != int64(0)) != 0 {
					} else {
						v41 = int32(4459156)
						v43 = *(*int32)(unsafe.Add(mBase, _consts[11]))
						v44 = int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[11])) = v43 + v44
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v47 + v44
						*(*int64)(unsafe.Add(mBase, uint32(v28)+392)) = v24
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v47 + int32(2)
						v58 = *(*int32)(unsafe.Add(mBase, _consts[11]))
						*(*int32)(unsafe.Add(mBase, _consts[11])) = v58 - v44
					}
				}
			}
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v23 != 0 {
				v65 = int32(4)
			} else {
				v65 = int32(5)
			}
			v66 = F_table_open(m, v62, v65)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				if v23 != 0 {
					v71 = int32(3)
				} else {
					v71 = int32(8)
				}
				v72 = F_index_open(m, v68, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return
				} else {
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v76 = F_palloc(m, int32(80))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						v78 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v76))) = v72
						v81 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v76)+40)) = v81
						*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v81
						*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = v74
						*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = v81
						v90 = F_brin_build_desc(m, v72)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v76)+44)) = v90
							v93 = F_brin_new_memtuple(m, v90)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v95 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v76)+72)) = v95
								v97 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v76)+64)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v76)+48)) = v93
								v101 = *(*int32)(unsafe.Add(mBase, _consts[3]))
								*(*int64)(unsafe.Add(mBase, uint32(v76)+52)) = v97
								*(*int32)(unsafe.Add(mBase, uint32(v76)+60)) = v101
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
								v107 = base.I32_rem_u_s(int32(-2), v74)
								*(*int32)(unsafe.Add(mBase, uint32(v76)+36)) = v105 - v107 - int32(2)
								v114 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), v95)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return
								} else {
									F_tuplesort_attach_shared(m, v114, l0)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										v121 = F___memcpy(m, int32(4362920), int32(4362760), int32(128))
										mBase = m.M
										v124 = *(*int64)(unsafe.Add(mBase, _consts[12]))
										*(*int64)(unsafe.Add(mBase, _consts[13])) = v124
										v128 = *(*int64)(unsafe.Add(mBase, _consts[14]))
										*(*int64)(unsafe.Add(mBase, _consts[15])) = v128
										v132 = *(*int64)(unsafe.Add(mBase, _consts[16]))
										*(*int64)(unsafe.Add(mBase, _consts[17])) = v132
										v136 = *(*int64)(unsafe.Add(mBase, _consts[18]))
										*(*int64)(unsafe.Add(mBase, _consts[19])) = v136
										v139 = *(*int32)(unsafe.Add(mBase, _consts[7]))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										v141 = base.I32_div_s(v139, v140)
										F__brin_parallel_scan_and_build(m, v76, v21, v114, v66, v72, v141)
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return
										} else {
											v146 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												v150 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
												mBase = m.M
												v151 = m.ExcPending
												if v151 != 0 {
													return
												} else {
													v153 = *(*int32)(unsafe.Add(mBase, _consts[20]))
													v159 = v150 + v153<<(uint(int32(5))%32)
													v164 = F___memset(m, v146+v153<<(uint(int32(7))%32), int32(0), int32(128))
													mBase = m.M
													F_BufferUsageAccumDiff(m, v164, int32(4362920))
													mBase = m.M
													v168 = v159 + int32(24)
													v169 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v168))) = v169
													v172 = v159 + int32(16)
													*(*int64)(unsafe.Add(mBase, uint32(v172))) = v169
													v176 = v159 + int32(8)
													*(*int64)(unsafe.Add(mBase, uint32(v176))) = v169
													*(*int64)(unsafe.Add(mBase, uint32(v159))) = v169
													v182 = *(*int64)(unsafe.Add(mBase, _consts[14]))
													v184 = *(*int64)(unsafe.Add(mBase, _consts[15]))
													*(*int64)(unsafe.Add(mBase, uint32(v172))) = v182 - v184
													v188 = *(*int64)(unsafe.Add(mBase, _consts[18]))
													v190 = *(*int64)(unsafe.Add(mBase, _consts[19]))
													*(*int64)(unsafe.Add(mBase, uint32(v159))) = v188 - v190
													v194 = *(*int64)(unsafe.Add(mBase, _consts[16]))
													v196 = *(*int64)(unsafe.Add(mBase, _consts[17]))
													*(*int64)(unsafe.Add(mBase, uint32(v176))) = v194 - v196
													v200 = *(*int64)(unsafe.Add(mBase, _consts[12]))
													v202 = *(*int64)(unsafe.Add(mBase, _consts[13]))
													*(*int64)(unsafe.Add(mBase, uint32(v168))) = v200 - v202
													F_relation_close(m, v72, v71)
													mBase = m.M
													v206 = m.ExcPending
													if v206 != 0 {
														return
													} else {
														F_sequence_close(m, v66, v65)
														mBase = m.M
														v208 = m.ExcPending
														if v208 != 0 {
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
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int64
	_ = v259
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v302 int64
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int64
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int64
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int64
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v398 int64
	_ = v398
	var v399 int32
	_ = v399
	var v402 int64
	_ = v402
	var v406 int64
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v417 int64
	_ = v417
	var v434 int64
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	v2 = int32(0)
	v17 = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v28 = F_pg_detoast_datum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			if base.Ui32(v30) < base.Ui32(int32(8)) {
				v434 = v17
			} else {
				v33 = int32(3)
				v34 = int32(base.Ui32(v30) >> (uint(v33) % 32))
				v36 = v34 & v33
				v37 = int32(16)
				v38 = v22 + v37
				v40 = v28 + v37
				v41 = int32(0)
				v43 = v34 - int32(1)
				if base.Ui32(v33) <= base.Ui32(v43) {
					v48 = v41
					v55 = v2
					for {
						v65 = v48 + v38
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48+v40))))
						v69 = v66 | v68
						*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v69)
						v72 = v48 | int32(1)
						v73 = v38 + v72
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v40))))
						v77 = v74 | v76
						*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v77)
						v80 = v48 | int32(2)
						v81 = v38 + v80
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v40))))
						v85 = v82 | v84
						*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v85)
						v88 = v48 | int32(3)
						v89 = v38 + v88
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
						v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88+v40))))
						v93 = v90 | v92
						*(*uint8)(unsafe.Add(mBase, uint32(v89))) = uint8(v93)
						v95 = int32(4)
						v96 = v48 + v95
						v98 = v55 + v95
						if v98 != v34&int32(536870908) {
							v48 = v96
							v55 = v98
							continue
						} else {
							break
						}
						break
					}
					v100 = v96
				} else {
					v100 = v41
				}
				if v36 != 0 {
					v117 = v100
					v130 = v2
					for {
						v134 = v117 + v38
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
						v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117+v40))))
						v138 = v135 | v137
						*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v138)
						v140 = int32(1)
						v143 = v130 + v140
						if v143 != v36 {
							v117 = v117 + v140
							v130 = v143
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v163 = v22 + int32(16)
				if base.Ui32(v30) <= base.Ui32(int32(31)) {
					if base.Ui32(int32(3)) <= base.Ui32(v43) {
						v169 = v163
						v170 = int32(0)
						v185 = v17
						for {
							v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+3)))
							v189 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v186)+uint32(_consts[21]))))
							v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+2)))
							v193 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_consts[21]))))
							v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
							v197 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v194)+uint32(_consts[21]))))
							v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
							v201 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v198)+uint32(_consts[21]))))
							v205 = v189 + (v193 + (v197 + (v185 + v201)))
							v206 = int32(4)
							v207 = v169 + v206
							v209 = v170 + v206
							if v209 != 0 {
								v169 = v207
								v170 = v209
								v185 = v205
								continue
							} else {
								break
							}
							break
						}
						v210 = v207
						v226 = v205
					} else {
						v210 = v163
						v226 = v17
					}
					v228 = v210
					v229 = int32(0)
					v244 = v226
					for {
						v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
						v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_consts[21]))))
						v249 = v244 + v248
						v250 = int32(1)
						v253 = v229 + v250
						if v253 != v34 {
							v228 = v228 + v250
							v229 = v253
							v244 = v249
							continue
						} else {
							break
						}
						break
					}
					v434 = v249
				} else {
					v259 = int64(0)
					if v34 < int32(4) {
						v338 = v163
						v339 = v34
						v344 = v259
					} else {
						if v163 != (v22+int32(19))&int32(-4) {
							v338 = v163
							v339 = v34
							v344 = v259
						} else {
							v268 = v34 - int32(4)
							v272 = int32(base.Ui32(v268)>>(uint(int32(2))%32)) + int32(1)
							v274 = v272 & int32(3)
							if base.Ui32(v268) < base.Ui32(int32(12)) {
								v310 = v163
								v311 = v34
								v316 = v259
							} else {
								v280 = v163
								v281 = v34
								v282 = int32(0)
								v286 = v259
								for {
									v287 = *(*int32)(unsafe.Add(mBase, uint32(v280)+12))
									v290 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
									v293 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
									v296 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
									v302 = base.I64_extend_i32_u(base.I32_popcnt(v287)) + (base.I64_extend_i32_u(base.I32_popcnt(v290)) + (base.I64_extend_i32_u(base.I32_popcnt(v293)) + (v286 + base.I64_extend_i32_u(base.I32_popcnt(v296)))))
									v303 = int32(16)
									v304 = v281 - v303
									v306 = v280 + v303
									v308 = v282 + int32(4)
									if v308 != v272&int32(2147483644) {
										v280 = v306
										v281 = v304
										v282 = v308
										v286 = v302
										continue
									} else {
										break
									}
									break
								}
								v310 = v306
								v311 = v304
								v316 = v302
							}
							if v274 == int32(0) {
								v338 = v310
								v339 = v311
								v344 = v316
							} else {
								v321 = v311
								v322 = v310
								v323 = int32(0)
								v326 = v316
								for {
									v327 = int32(4)
									v328 = v321 - v327
									v329 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
									v332 = v326 + base.I64_extend_i32_u(base.I32_popcnt(v329))
									v334 = v322 + v327
									v336 = v323 + int32(1)
									if v336 != v274 {
										v321 = v328
										v322 = v334
										v323 = v336
										v326 = v332
										continue
									} else {
										break
									}
									break
								}
								v338 = v334
								v339 = v328
								v344 = v332
							}
						}
					}
					if v339 == int32(0) {
						v417 = v344
					} else {
						v348 = v339 & int32(3)
						if v348 == int32(0) {
							v371 = v338
							v373 = v339
							v377 = v344
						} else {
							v354 = v339
							v355 = v338
							v356 = int32(0)
							v358 = v344
							for {
								v359 = int32(1)
								v360 = v354 - v359
								v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v355))))
								v364 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v361)+uint32(_consts[21]))))
								v365 = v358 + v364
								v367 = v355 + v359
								v369 = v356 + v359
								if v369 != v348 {
									v354 = v360
									v355 = v367
									v356 = v369
									v358 = v365
									continue
								} else {
									break
								}
								break
							}
							v371 = v367
							v373 = v360
							v377 = v365
						}
						if base.Ui32(v339) < base.Ui32(int32(4)) {
							v417 = v377
						} else {
							v380 = v371
							v382 = v373
							v386 = v377
							for {
								v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+3)))
								v390 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v387)+uint32(_consts[21]))))
								v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+2)))
								v394 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v391)+uint32(_consts[21]))))
								v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+1)))
								v398 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v395)+uint32(_consts[21]))))
								v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380))))
								v402 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v399)+uint32(_consts[21]))))
								v406 = v390 + (v394 + (v398 + (v386 + v402)))
								v407 = int32(4)
								v410 = v382 - v407
								if v410 != 0 {
									v380 = v380 + v407
									v382 = v410
									v386 = v406
									continue
								} else {
									break
								}
								break
							}
							v417 = v406
						}
					}
					v434 = v417
				}
			}
			*(*uint32)(unsafe.Add(mBase, uint32(v22)+12)) = uint32(v434)
			v436 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)))
			if v22 != v437 {
				F_pfree(m, v437)
				mBase = m.M
				v440 = m.ExcPending
				if v440 != 0 {
					return int32(0)
				} else {
					v441 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v441))) = v22
					v443 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
					if v28 != v444 {
						F_pfree(m, v28)
						mBase = m.M
						v447 = m.ExcPending
						if v447 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					} else {
						return int32(0)
					}
				}
			} else {
				v443 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
				if v28 != v444 {
					F_pfree(m, v28)
					mBase = m.M
					v447 = m.ExcPending
					if v447 != 0 {
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
	var v70 int32
	_ = v70
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
	switch v21 - int32(1) {
	case 0, 1:
		v63 = F_minmax_get_strategy_procinfo(m, v20, v18, v17, v21)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			v67 = F_FunctionCall2Coll(m, v63, v19, v66, v16)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				v70 = v67
				m.G0 = v12 + int32(16)
				return v70
			}
		}
	case 2:
		v26 = F_minmax_get_strategy_procinfo(m, v20, v18, v17, int32(2))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
			v32 = F_FunctionCall2Coll(m, v26, v19, v31, v16)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v32 == int32(0) {
					v70 = int32(0)
					m.G0 = v12 + int32(16)
					return v70
				} else {
					v37 = F_minmax_get_strategy_procinfo(m, v20, v18, v17, int32(4))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
						v41 = F_FunctionCall2Coll(m, v37, v19, v40, v16)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v70 = v41
							m.G0 = v12 + int32(16)
							return v70
						}
					}
				}
			}
		}
	case 3, 4:
		v43 = F_minmax_get_strategy_procinfo(m, v20, v18, v17, v21)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
			v47 = F_FunctionCall2Coll(m, v43, v19, v46, v16)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v70 = v47
				m.G0 = v12 + int32(16)
				return v70
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v53
			F_errmsg_internal(m, int32(460999), v12)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(481461), int32(195), int32(89969))
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v3 = F_palloc0(m, int32(188))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+2)) = uint8(v7)
		*(*uint16)(unsafe.Add(mBase, uint32(v3))) = uint16(v7)
		*(*int32)(unsafe.Add(mBase, uint32(v3)+4)) = (v3 + int32(19)) & int32(-8)
		v18 = F_lookup_type_cache(m, int32(4601), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = v18
			return v3
		}
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
	F_add_local_int_reloption(m, v2, int32(392659), int32(477730), int32(32), int32(8), int32(256))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_brin_minmax_multi_summary_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(17430)
			F_errmsg(m, int32(188657), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(486510), int32(2984), int32(273028))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_brin_range_deserialize(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v28 = F_palloc0(m, l0<<(uint(int32(2))%32)+int32(36))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l0
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v35
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v41
	v44 = l1 + int32(20)
	v45 = F_get_typbyval(m, v41)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v48 = F_get_typlen(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v52 = v22 + v23<<(uint(int32(1))%32)
	v54 = base.B2i32(v52 <= int32(0))
	if v52 <= int32(0) {
		v212 = v3
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v54 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L6:
	;
	if v45 != 0 {
		v212 = v3
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v59 = int32(0)
	v67 = v59
	v68 = v3
	v71 = v44
	goto L8
L8:
	;
	if base.B2i32(v48 <= v59) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	if v196 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L10:
	;
	v203 = v67 + int32(1)
	if v203 < v52 {
		v67 = v203
		v68 = v196
		v71 = v197
		goto L8
	} else {
		goto L54
	}
L11:
	;
	v196 = v68 + (v48+int32(7))&int32(65528)
	v197 = v71
	goto L10
L12:
	;
	goto L13
L13:
	;
	switch v48&int32(65535) - int32(65534) {
	case 0:
		goto L16
	case 1:
		goto L17
	default:
		v196 = v68
		v197 = v71
		goto L10
	}
L14:
	;
	v196 = v68 + v191
	v197 = v71 + v190
	goto L10
L15:
	;
	v169 = int32(8)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v174 = base.B2i32(v172 == int32(18))
	if v172 == int32(18) {
		goto L39
	} else {
		goto L40
	}
L16:
	;
	if v71&int32(3) == int32(0) {
		v127 = v71
		goto L24
	} else {
		goto L25
	}
L17:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v86 == int32(1) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	if v86&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v92 = int32(base.Ui32(v86) >> (uint(int32(1)) % 32))
	v190 = v92
	v191 = (v92 + int32(7)) & int32(248)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v99 = int32(base.Ui32(v97) >> (uint(int32(2)) % 32))
	v190 = v99
	v191 = (v99 + int32(7)) & int32(2147483640)
	goto L14
L22:
	;
	v196 = v68 + v160&int32(-8) + int32(8)
	v197 = v71 + v160 + int32(1)
	goto L10
L23:
	;
	v160 = v152 - v71
	goto L22
L24:
	;
	v131 = v127
	goto L33
L25:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v111 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v160 = int32(0)
	goto L22
L27:
	;
	goto L28
L28:
	;
	v116 = v71
	goto L29
L29:
	;
	v120 = v116 + int32(1)
	if v120&int32(3) == int32(0) {
		v127 = v120
		goto L24
	} else {
		goto L31
	}
L30:
	;
	v152 = v120
	goto L23
L31:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v125 != 0 {
		v116 = v120
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v140 = int32(-2139062144)
	if (int32(16843008)-v137|v137)&v140 == v140 {
		v131 = v131 + int32(4)
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v146 = v131
	goto L36
L35:
	;
	goto L34
L36:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 != 0 {
		v146 = v146 + int32(1)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v152 = v146
	goto L23
L38:
	;
	goto L37
L39:
	;
	v175 = int32(24)
	goto L41
L40:
	;
	v175 = v169
	goto L41
L41:
	;
	v179 = base.B2i32(v172&int32(254) == int32(2))
	if v172&int32(254) == int32(2) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v180 = v169
	goto L44
L43:
	;
	v180 = v175
	goto L44
L44:
	;
	v181 = int32(6)
	if v172 == int32(18) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v185 = int32(18)
	goto L47
L46:
	;
	v185 = int32(2)
	goto L47
L47:
	;
	if v172&int32(254) == int32(2) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v186 = v181
	goto L50
L49:
	;
	v186 = v185
	goto L50
L50:
	;
	if v172 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v189 = v181
	goto L53
L52:
	;
	v189 = v186
	goto L53
L53:
	;
	v190 = v189
	v191 = v180
	goto L14
L54:
	;
	goto L9
L55:
	;
	v212 = int32(0)
	goto L5
L56:
	;
	goto L57
L57:
	;
	v208 = F_palloc(m, v196)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v212 = v208
	goto L5
L59:
	;
	v230 = v28 + int32(36)
	v236 = v48 & int32(65535)
	v240 = v44
	v241 = int32(0)
	v242 = v212
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v20 + int32(16)
	return v28
L62:
	;
	if v45 != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	v453 = v241 + int32(1)
	if v453 != v52 {
		v240 = v447
		v241 = v453
		v242 = v448
		goto L62
	} else {
		goto L148
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(0)
	if v48 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	goto L67
L67:
	;
	if int32(0) < v48 {
		goto L80
	} else {
		goto L81
	}
L68:
	;
	switch v236 - int32(1) {
	case 0:
		goto L73
	case 1:
		goto L76
	default:
		goto L74
	case 3:
		goto L75
	}
L69:
	;
	v261 = F__emscripten_memcpy_bulkmem(m, v20+int32(12), v240, v48)
	mBase = m.M
	goto L71
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v241<<(uint(int32(2))%32)))) = v282
	v447 = v240 + v48
	v448 = v242
	goto L64
L73:
	;
	v281 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+12)))
	v282 = v281
	goto L72
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v282 = v267
	goto L72
L76:
	;
	v266 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+12)))
	v282 = v266
	goto L72
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v48
	F_errmsg_internal(m, int32(472527), v20)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(319286), int32(70), int32(66336))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v241<<(uint(int32(2))%32)))) = v242
	if v48 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	switch v236 - int32(65534) {
	case 0:
		goto L88
	case 1:
		goto L89
	default:
		v447 = v240
		v448 = v242
		goto L64
	}
L83:
	;
	v447 = v240 + v48
	v448 = v292 + (v48+int32(7))&int32(65528)
	goto L64
L84:
	;
	v291 = F__emscripten_memcpy_bulkmem(m, v242, v240, v48)
	mBase = m.M
	v292 = v291
	goto L86
L85:
	;
	v292 = v242
	goto L86
L86:
	;
	goto L83
L87:
	;
	v424 = int32(8)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	v429 = base.B2i32(v427 == int32(18))
	if v427 == int32(18) {
		goto L133
	} else {
		goto L134
	}
L88:
	;
	if v240&int32(3) == int32(0) {
		v376 = v240
		goto L114
	} else {
		goto L115
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v241<<(uint(int32(2))%32)))) = v242
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v301 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v328 != 0 {
		goto L105
	} else {
		goto L106
	}
L91:
	;
	v304 = int32(6)
	v306 = int32(18)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	if v308 == v306 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v320 = int32(1)
	if v301&v320 != 0 {
		v328 = int32(base.Ui32(v301) >> (uint(v320) % 32))
		goto L90
	} else {
		goto L103
	}
L94:
	;
	v311 = v306
	goto L96
L95:
	;
	v311 = int32(2)
	goto L96
L96:
	;
	if v308&int32(254) == int32(2) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v316 = v304
	goto L99
L98:
	;
	v316 = v311
	goto L99
L99:
	;
	if v308 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v319 = v304
	goto L102
L101:
	;
	v319 = v316
	goto L102
L102:
	;
	v328 = v319
	goto L90
L103:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v328 = int32(base.Ui32(v324) >> (uint(int32(2)) % 32))
	goto L90
L104:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v331 == int32(1) {
		goto L87
	} else {
		goto L108
	}
L105:
	;
	v329 = F__emscripten_memcpy_bulkmem(m, v242, v240, v328)
	mBase = m.M
	v330 = v329
	goto L107
L106:
	;
	v330 = v242
	goto L107
L107:
	;
	goto L104
L108:
	;
	if v331&int32(1) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v337 = int32(base.Ui32(v331) >> (uint(int32(1)) % 32))
	v447 = v240 + v337
	v448 = v330 + (v337+int32(7))&int32(248)
	goto L64
L110:
	;
	goto L111
L111:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	v346 = int32(base.Ui32(v344) >> (uint(int32(2)) % 32))
	v447 = v240 + v346
	v448 = v330 + (v346+int32(7))&int32(2147483640)
	goto L64
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v241<<(uint(int32(2))%32)))) = v242
	v415 = v409 + int32(1)
	if v415 != 0 {
		goto L130
	} else {
		goto L131
	}
L113:
	;
	v409 = v401 - v240
	goto L112
L114:
	;
	v380 = v376
	goto L123
L115:
	;
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	if v360 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v409 = int32(0)
	goto L112
L117:
	;
	goto L118
L118:
	;
	v365 = v240
	goto L119
L119:
	;
	v369 = v365 + int32(1)
	if v369&int32(3) == int32(0) {
		v376 = v369
		goto L114
	} else {
		goto L121
	}
L120:
	;
	v401 = v369
	goto L113
L121:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	if v374 != 0 {
		v365 = v369
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v389 = int32(-2139062144)
	if (int32(16843008)-v386|v386)&v389 == v389 {
		v380 = v380 + int32(4)
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v395 = v380
	goto L126
L125:
	;
	goto L124
L126:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v399 != 0 {
		v395 = v395 + int32(1)
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v401 = v395
	goto L113
L128:
	;
	goto L127
L129:
	;
	v447 = v240 + v415
	v448 = v417 + v409&int32(-8) + int32(8)
	goto L64
L130:
	;
	v416 = F__emscripten_memcpy_bulkmem(m, v242, v240, v415)
	mBase = m.M
	v417 = v416
	goto L132
L131:
	;
	v417 = v242
	goto L132
L132:
	;
	goto L129
L133:
	;
	v430 = int32(24)
	goto L135
L134:
	;
	v430 = v424
	goto L135
L135:
	;
	v434 = base.B2i32(v427&int32(254) == int32(2))
	if v427&int32(254) == int32(2) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v435 = v424
	goto L138
L137:
	;
	v435 = v430
	goto L138
L138:
	;
	v437 = int32(6)
	if v427 == int32(18) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v441 = int32(18)
	goto L141
L140:
	;
	v441 = int32(2)
	goto L141
L141:
	;
	if v427&int32(254) == int32(2) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v442 = v437
	goto L144
L143:
	;
	v442 = v441
	goto L144
L144:
	;
	if v427 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v445 = v437
	goto L147
L146:
	;
	v445 = v442
	goto L147
L147:
	;
	v447 = v240 + v445
	v448 = v330 + v435
	goto L64
L148:
	;
	goto L63
}
