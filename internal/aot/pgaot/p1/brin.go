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
						v41 = int32(4484100)
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
										v121 = F___memcpy(m, int32(4387864), int32(4387704), int32(128))
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
													F_BufferUsageAccumDiff(m, v164, int32(4387864))
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
			F_errmsg_internal(m, int32(470788), v12)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(491908), int32(195), int32(91572))
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
	F_add_local_int_reloption(m, v2, int32(401041), int32(487940), int32(32), int32(8), int32(256))
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
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(17900)
			F_errmsg(m, int32(192260), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(497082), int32(2984), int32(278655))
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
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
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
		v156 = v3
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v54 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	if v45 != 0 {
		v156 = v3
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
	if v140 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L10:
	;
	v147 = v67 + int32(1)
	if v147 < v52 {
		v67 = v147
		v68 = v140
		v71 = v141
		goto L8
	} else {
		goto L37
	}
L11:
	;
	v140 = v68 + (v48+int32(7))&int32(65528)
	v141 = v71
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
		v140 = v68
		v141 = v71
		goto L10
	}
L14:
	;
	v140 = v68 + v135
	v141 = v71 + v134
	goto L10
L15:
	;
	v113 = int32(8)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v118 = base.B2i32(v116 == int32(18))
	if v116 == int32(18) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v104 = F_strlen(m, v71)
	mBase = m.M
	v140 = v68 + v104&int32(-8) + int32(8)
	v141 = v71 + v104 + int32(1)
	goto L10
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
	v134 = v92
	v135 = (v92 + int32(7)) & int32(248)
	goto L14
L20:
	;
	goto L21
L21:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v99 = int32(base.Ui32(v97) >> (uint(int32(2)) % 32))
	v134 = v99
	v135 = (v99 + int32(7)) & int32(2147483640)
	goto L14
L22:
	;
	v119 = int32(24)
	goto L24
L23:
	;
	v119 = v113
	goto L24
L24:
	;
	v123 = base.B2i32(v116&int32(254) == int32(2))
	if v116&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v124 = v113
	goto L27
L26:
	;
	v124 = v119
	goto L27
L27:
	;
	v125 = int32(6)
	if v116 == int32(18) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v129 = int32(18)
	goto L30
L29:
	;
	v129 = int32(2)
	goto L30
L30:
	;
	if v116&int32(254) == int32(2) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v130 = v125
	goto L33
L32:
	;
	v130 = v129
	goto L33
L33:
	;
	if v116 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v133 = v125
	goto L36
L35:
	;
	v133 = v130
	goto L36
L36:
	;
	v134 = v133
	v135 = v124
	goto L14
L37:
	;
	goto L9
L38:
	;
	v156 = int32(0)
	goto L5
L39:
	;
	goto L40
L40:
	;
	v152 = F_palloc(m, v140)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v156 = v152
	goto L5
L42:
	;
	v174 = v28 + int32(36)
	v180 = v48 & int32(65535)
	v184 = v44
	v185 = int32(0)
	v186 = v156
	goto L45
L43:
	;
	goto L44
L44:
	;
	m.G0 = v20 + int32(16)
	return v28
L45:
	;
	if v45 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L44
L47:
	;
	v341 = v185 + int32(1)
	if v341 != v52 {
		v184 = v335
		v185 = v341
		v186 = v336
		goto L45
	} else {
		goto L114
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(0)
	if v48 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) < v48 {
		goto L63
	} else {
		goto L64
	}
L51:
	;
	switch v180 - int32(1) {
	case 0:
		goto L56
	case 1:
		goto L59
	default:
		goto L57
	case 3:
		goto L58
	}
L52:
	;
	v205 = F__emscripten_memcpy_bulkmem(m, v20+int32(12), v184, v48)
	mBase = m.M
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174+v185<<(uint(int32(2))%32)))) = v226
	v335 = v184 + v48
	v336 = v186
	goto L47
L56:
	;
	v225 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20)+12)))
	v226 = v225
	goto L55
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v226 = v211
	goto L55
L59:
	;
	v210 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+12)))
	v226 = v210
	goto L55
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v48
	F_errmsg_internal(m, int32(482718), v20)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(326157), int32(70), int32(67716))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v174+v185<<(uint(int32(2))%32)))) = v186
	if v48 != 0 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	switch v180 - int32(65534) {
	case 0:
		goto L71
	case 1:
		goto L72
	default:
		v335 = v184
		v336 = v186
		goto L47
	}
L66:
	;
	v335 = v184 + v48
	v336 = v236 + (v48+int32(7))&int32(65528)
	goto L47
L67:
	;
	v235 = F__emscripten_memcpy_bulkmem(m, v186, v184, v48)
	mBase = m.M
	v236 = v235
	goto L69
L68:
	;
	v236 = v186
	goto L69
L69:
	;
	goto L66
L70:
	;
	v312 = int32(8)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	v317 = base.B2i32(v315 == int32(18))
	if v315 == int32(18) {
		goto L99
	} else {
		goto L100
	}
L71:
	;
	v297 = F_strlen(m, v184)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v174+v185<<(uint(int32(2))%32)))) = v186
	v303 = v297 + int32(1)
	if v303 != 0 {
		goto L96
	} else {
		goto L97
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174+v185<<(uint(int32(2))%32)))) = v186
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v245 == int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v272 != 0 {
		goto L88
	} else {
		goto L89
	}
L74:
	;
	v248 = int32(6)
	v250 = int32(18)
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	if v252 == v250 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v264 = int32(1)
	if v245&v264 != 0 {
		v272 = int32(base.Ui32(v245) >> (uint(v264) % 32))
		goto L73
	} else {
		goto L86
	}
L77:
	;
	v255 = v250
	goto L79
L78:
	;
	v255 = int32(2)
	goto L79
L79:
	;
	if v252&int32(254) == int32(2) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v260 = v248
	goto L82
L81:
	;
	v260 = v255
	goto L82
L82:
	;
	if v252 == int32(1) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v263 = v248
	goto L85
L84:
	;
	v263 = v260
	goto L85
L85:
	;
	v272 = v263
	goto L73
L86:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v272 = int32(base.Ui32(v268) >> (uint(int32(2)) % 32))
	goto L73
L87:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v275 == int32(1) {
		goto L70
	} else {
		goto L91
	}
L88:
	;
	v273 = F__emscripten_memcpy_bulkmem(m, v186, v184, v272)
	mBase = m.M
	v274 = v273
	goto L90
L89:
	;
	v274 = v186
	goto L90
L90:
	;
	goto L87
L91:
	;
	if v275&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v281 = int32(base.Ui32(v275) >> (uint(int32(1)) % 32))
	v335 = v184 + v281
	v336 = v274 + (v281+int32(7))&int32(248)
	goto L47
L93:
	;
	goto L94
L94:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v290 = int32(base.Ui32(v288) >> (uint(int32(2)) % 32))
	v335 = v184 + v290
	v336 = v274 + (v290+int32(7))&int32(2147483640)
	goto L47
L95:
	;
	v335 = v184 + v303
	v336 = v305 + v297&int32(-8) + int32(8)
	goto L47
L96:
	;
	v304 = F__emscripten_memcpy_bulkmem(m, v186, v184, v303)
	mBase = m.M
	v305 = v304
	goto L98
L97:
	;
	v305 = v186
	goto L98
L98:
	;
	goto L95
L99:
	;
	v318 = int32(24)
	goto L101
L100:
	;
	v318 = v312
	goto L101
L101:
	;
	v322 = base.B2i32(v315&int32(254) == int32(2))
	if v315&int32(254) == int32(2) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v323 = v312
	goto L104
L103:
	;
	v323 = v318
	goto L104
L104:
	;
	v325 = int32(6)
	if v315 == int32(18) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v329 = int32(18)
	goto L107
L106:
	;
	v329 = int32(2)
	goto L107
L107:
	;
	if v315&int32(254) == int32(2) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v330 = v325
	goto L110
L109:
	;
	v330 = v329
	goto L110
L110:
	;
	if v315 == int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v333 = v325
	goto L113
L112:
	;
	v333 = v330
	goto L113
L113:
	;
	v335 = v184 + v333
	v336 = v274 + v323
	goto L47
L114:
	;
	goto L46
}
