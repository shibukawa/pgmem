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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
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
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v150 int64
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int64
	_ = v179
	var v190 int64
	_ = v190
	var v192 int64
	_ = v192
	var v196 int64
	_ = v196
	var v198 int64
	_ = v198
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v208 int64
	_ = v208
	var v210 int64
	_ = v210
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
						v53 = int32(0)
						v55 = int32(_a_F__brin_parallel_build_main_1)
						v56 = base.AtomicRmwOr32(m, v53, v55, v53)
						*(*int64)(unsafe.Add(mBase, uint32(v28)+392)) = v24
						v61 = base.AtomicRmwOr32(m, v53, v55, v53)
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
						*(*int32)(unsafe.Add(mBase, uint32(v28))) = v62 + v46
						v68 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3]))
						*(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[3])) = v68 - v46
					}
				}
			}
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v23 != 0 {
				v75 = int32(4)
			} else {
				v75 = int32(5)
			}
			v76 = F_table_open(m, v72, v75)
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return
			} else {
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
				if v23 != 0 {
					v81 = int32(3)
				} else {
					v81 = int32(8)
				}
				v82 = F_index_open(m, v78, v81)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
					v86 = F_palloc(m, int32(80))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v88 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v86))) = v82
						*(*int64)(unsafe.Add(mBase, uint32(v86)+16)) = v88
						v93 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v86)+24)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(v86)+32)) = v93
						*(*int32)(unsafe.Add(mBase, uint32(v86)+28)) = v84
						v100 = F_brin_build_desc(m, v82)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v86)+44)) = v100
							v103 = F_brin_new_memtuple(m, v100)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								v105 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v86)+72)) = v105
								v107 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v86)+64)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v86)+48)) = v103
								v111 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[4]))
								*(*int64)(unsafe.Add(mBase, uint32(v86)+52)) = v107
								*(*int32)(unsafe.Add(mBase, uint32(v86)+60)) = v111
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
								v117 = base.I32_rem_u_s(int32(-2), v84)
								*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v115 - v117 - int32(2)
								v124 = F_shm_toc_lookup(m, l1, int64(-5764607523034234878), v105)
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return
								} else {
									F_tuplesort_attach_shared(m, v124, l0)
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return
									} else {
										base.MemoryCopy(m, int32(_a_F__brin_parallel_build_main_2), int32(_a_F__brin_parallel_build_main_3), int32(128))
										v134 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[5]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[6])) = v134
										v138 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[7]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[8])) = v138
										v142 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[9]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[10])) = v142
										v146 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[11]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[12])) = v146
										v150 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[13]))
										*(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[14])) = v150
										v153 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[15]))
										v154 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
										v155 = base.I32_div_s(v153, v154)
										F__brin_parallel_scan_and_build(m, v86, v21, v124, v76, v82, v155)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return
										} else {
											v160 = F_shm_toc_lookup(m, l1, int64(-5764607523034234875), int32(0))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												v164 = F_shm_toc_lookup(m, l1, int64(-5764607523034234876), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return
												} else {
													v167 = *(*int32)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[16]))
													v170 = v160 + v167<<(uint(int32(7))%32)
													v173 = v164 + v167*int32(40)
													base.MemoryFill(m, v170, int32(0), int32(128))
													F_BufferUsageAccumDiff(m, v170, int32(_a_F__brin_parallel_build_main_2))
													mBase = m.M
													v179 = int64(0)
													*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v179
													*(*int64)(unsafe.Add(mBase, uint32(v173)+24)) = v179
													*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v179
													*(*int64)(unsafe.Add(mBase, uint32(v173)+8)) = v179
													*(*int64)(unsafe.Add(mBase, uint32(v173))) = v179
													v190 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[9]))
													v192 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[10]))
													*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v190 - v192
													v196 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[13]))
													v198 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[14]))
													*(*int64)(unsafe.Add(mBase, uint32(v173))) = v196 - v198
													v202 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[11]))
													v204 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[12]))
													*(*int64)(unsafe.Add(mBase, uint32(v173)+8)) = v202 - v204
													v208 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[7]))
													v210 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[8]))
													*(*int64)(unsafe.Add(mBase, uint32(v173)+24)) = v208 - v210
													v214 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[5]))
													v216 = *(*int64)(unsafe.Add(mBase, _c_F__brin_parallel_build_main[6]))
													*(*int64)(unsafe.Add(mBase, uint32(v173)+32)) = v214 - v216
													F_relation_close(m, v82, v81)
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return
													} else {
														F_relation_close(m, v76, v75)
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
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
func F_brin_bloom_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int64
	_ = v16
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v186 int64
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v222 int64
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v243 int64
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v271 int32
	_ = v271
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v296 int64
	_ = v296
	var v297 int32
	_ = v297
	var v298 int64
	_ = v298
	var v299 int32
	_ = v299
	var v300 int64
	_ = v300
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int64
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v324 int64
	_ = v324
	var v325 int64
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int64
	_ = v333
	var v343 int64
	_ = v343
	var v359 int64
	_ = v359
	var v362 int32
	_ = v362
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int64
	_ = v372
	var v376 int32
	_ = v376
	v2 = int32(0)
	v16 = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int64(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v28 = F_pg_detoast_datum(m, v27)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
			v32 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
			if v32 != 0 {
				v34 = v32 & int32(3)
				v35 = int32(16)
				v36 = v22 + v35
				v38 = v28 + v35
				v39 = int32(0)
				if base.Ui32(int32(32)) <= base.Ui32(v30) {
					v44 = v39
					v57 = v2
					for {
						v61 = v44 + v36
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
						v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v38))))
						v65 = v62 | v64
						*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v65)
						v68 = v44 | int32(1)
						v69 = v36 + v68
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v68))))
						v73 = v70 | v72
						*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v73)
						v76 = v44 | int32(2)
						v77 = v36 + v76
						v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v76))))
						v81 = v78 | v80
						*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v81)
						v84 = v44 | int32(3)
						v85 = v36 + v84
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
						v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v84))))
						v89 = v86 | v88
						*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v89)
						v91 = int32(4)
						v92 = v44 + v91
						v94 = v57 + v91
						if v94 != v32&int32(536870908) {
							v44 = v92
							v57 = v94
							continue
						} else {
							break
						}
						break
					}
					if v34 == int32(0) {
					} else {
						v98 = v92
						v115 = v98
						v129 = v2
						for {
							v132 = v115 + v36
							v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
							v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v38))))
							v136 = v133 | v135
							*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v136)
							v138 = int32(1)
							v141 = v129 + v138
							if v141 != v34 {
								v115 = v115 + v138
								v129 = v141
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v98 = v39
					v115 = v98
					v129 = v2
					for {
						v132 = v115 + v36
						v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
						v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115+v38))))
						v136 = v133 | v135
						*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v136)
						v138 = int32(1)
						v141 = v129 + v138
						if v141 != v34 {
							v115 = v115 + v138
							v129 = v141
							continue
						} else {
							break
						}
						break
					}
				}
				v161 = v22 + int32(16)
				if base.Ui32(int32(63)) < base.Ui32(v30) {
					v271 = v161
					v272 = int64(0)
					v273 = int32(0)
					if v32 == v273 {
						v343 = int64(0)
					} else {
						v280 = v32 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v285 = v271
							v287 = v272
							v290 = v273
							for {
								v291 = int32(4)
								v292 = v285 + v291
								v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+3)))
								v294 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v293)+uint32(_c_F_brin_bloom_union[0]))))
								v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+2)))
								v296 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v295)+uint32(_c_F_brin_bloom_union[0]))))
								v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
								v298 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v297)+uint32(_c_F_brin_bloom_union[0]))))
								v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
								v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v299)+uint32(_c_F_brin_bloom_union[0]))))
								v304 = v294 + (v296 + (v298 + (v287 + v300)))
								v306 = v290 + v291
								if v306 != v32&int32(-4) {
									v285 = v292
									v287 = v304
									v290 = v306
									continue
								} else {
									break
								}
								break
							}
							if v280 == int32(0) {
								v333 = v304
							} else {
								v310 = v292
								v312 = v304
								v317 = v310
								v318 = int32(0)
								v319 = v312
								for {
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
									v324 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_brin_bloom_union[0]))))
									v325 = v319 + v324
									v326 = int32(1)
									v329 = v318 + v326
									if v329 != v280 {
										v317 = v317 + v326
										v318 = v329
										v319 = v325
										continue
									} else {
										break
									}
									break
								}
								v333 = v325
							}
						} else {
							v310 = v271
							v312 = v272
							v317 = v310
							v318 = int32(0)
							v319 = v312
							for {
								v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
								v324 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_brin_bloom_union[0]))))
								v325 = v319 + v324
								v326 = int32(1)
								v329 = v318 + v326
								if v329 != v280 {
									v317 = v317 + v326
									v318 = v329
									v319 = v325
									continue
								} else {
									break
								}
								break
							}
							v333 = v325
						}
						v343 = v333
					}
					v359 = v343
				} else {
					v165 = v32 & int32(3)
					if base.Ui32(int32(32)) <= base.Ui32(v30) {
						v171 = v161
						v174 = int32(0)
						v186 = v16
						for {
							v188 = int32(4)
							v189 = v171 + v188
							v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+3)))
							v191 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v190)+uint32(_c_F_brin_bloom_union[0]))))
							v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+2)))
							v193 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v192)+uint32(_c_F_brin_bloom_union[0]))))
							v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+1)))
							v195 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v194)+uint32(_c_F_brin_bloom_union[0]))))
							v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
							v197 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v196)+uint32(_c_F_brin_bloom_union[0]))))
							v201 = v191 + (v193 + (v195 + (v186 + v197)))
							v203 = v174 + v188
							if v203 != v32&int32(4) {
								v171 = v189
								v174 = v203
								v186 = v201
								continue
							} else {
								break
							}
							break
						}
						if v165 == int32(0) {
							v359 = v201
						} else {
							v207 = v189
							v222 = v201
							v225 = v207
							v226 = int32(0)
							v240 = v222
							for {
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
								v243 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_brin_bloom_union[0]))))
								v244 = v240 + v243
								v245 = int32(1)
								v248 = v226 + v245
								if v248 != v165 {
									v225 = v225 + v245
									v226 = v248
									v240 = v244
									continue
								} else {
									break
								}
								break
							}
							v359 = v244
						}
					} else {
						v207 = v161
						v222 = v16
						v225 = v207
						v226 = int32(0)
						v240 = v222
						for {
							v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
							v243 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v242)+uint32(_c_F_brin_bloom_union[0]))))
							v244 = v240 + v243
							v245 = int32(1)
							v248 = v226 + v245
							if v248 != v165 {
								v225 = v225 + v245
								v226 = v248
								v240 = v244
								continue
							} else {
								break
							}
							break
						}
						v359 = v244
					}
				}
			} else {
				if base.Ui32(v30) < base.Ui32(int32(64)) {
					v359 = v16
				} else {
					v271 = v22 + int32(16)
					v272 = int64(0)
					v273 = int32(0)
					if v32 == v273 {
						v343 = int64(0)
					} else {
						v280 = v32 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v285 = v271
							v287 = v272
							v290 = v273
							for {
								v291 = int32(4)
								v292 = v285 + v291
								v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+3)))
								v294 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v293)+uint32(_c_F_brin_bloom_union[0]))))
								v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+2)))
								v296 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v295)+uint32(_c_F_brin_bloom_union[0]))))
								v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
								v298 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v297)+uint32(_c_F_brin_bloom_union[0]))))
								v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
								v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v299)+uint32(_c_F_brin_bloom_union[0]))))
								v304 = v294 + (v296 + (v298 + (v287 + v300)))
								v306 = v290 + v291
								if v306 != v32&int32(-4) {
									v285 = v292
									v287 = v304
									v290 = v306
									continue
								} else {
									break
								}
								break
							}
							if v280 == int32(0) {
								v333 = v304
							} else {
								v310 = v292
								v312 = v304
								v317 = v310
								v318 = int32(0)
								v319 = v312
								for {
									v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
									v324 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_brin_bloom_union[0]))))
									v325 = v319 + v324
									v326 = int32(1)
									v329 = v318 + v326
									if v329 != v280 {
										v317 = v317 + v326
										v318 = v329
										v319 = v325
										continue
									} else {
										break
									}
									break
								}
								v333 = v325
							}
						} else {
							v310 = v271
							v312 = v272
							v317 = v310
							v318 = int32(0)
							v319 = v312
							for {
								v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317))))
								v324 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v323)+uint32(_c_F_brin_bloom_union[0]))))
								v325 = v319 + v324
								v326 = int32(1)
								v329 = v318 + v326
								if v329 != v280 {
									v317 = v317 + v326
									v318 = v329
									v319 = v325
									continue
								} else {
									break
								}
								break
							}
							v333 = v325
						}
						v343 = v333
					}
					v359 = v343
				}
			}
			*(*uint32)(unsafe.Add(mBase, uint32(v22)+12)) = uint32(v359)
			v362 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			v363 = *(*int64)(unsafe.Add(mBase, uint32(v362)))
			v364 = base.I64_extend_i32_u(v22)
			if v363 != v364 {
				F_pfree(m, base.I32_wrap_i64(v363))
				mBase = m.M
				v368 = m.ExcPending
				if v368 != 0 {
					return int64(0)
				} else {
					v369 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v369))) = v364
					v371 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					v372 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
					if v372 != base.I64_extend_i32_u(v28) {
						F_pfree(m, v28)
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int64(0)
						} else {
							return int64(0)
						}
					} else {
						return int64(0)
					}
				}
			} else {
				v371 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				v372 = *(*int64)(unsafe.Add(mBase, uint32(v371)))
				if v372 != base.I64_extend_i32_u(v28) {
					F_pfree(m, v28)
					mBase = m.M
					v376 = m.ExcPending
					if v376 != 0 {
						return int64(0)
					} else {
						return int64(0)
					}
				} else {
					return int64(0)
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
func F_brin_metapage_info(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int64
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
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v16 = F_superuser(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v16 != 0 {
				v20 = F_verify_brin_page(m, v10, int32(_a_F_brin_metapage_info_0), int32(_a_F_brin_metapage_info_1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)))
					if v22 == int32(0) {
						v25 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v25)
						v64 = int64(0)
						m.G0 = v7 - int32(-64)
						return v64
					} else {
						v31 = F_get_call_result_type(m, l0, int32(0), v5+int32(-4))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							if v31 != int32(1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_brin_metapage_info_2), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_brin_metapage_info_3), int32(366), int32(_a_F_brin_metapage_info_4))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
								v36 = F_BlessTupleDesc(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = v36
									v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v39
									v42 = F_psprintf(m, int32(_a_F_brin_metapage_info_5), v7)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int64(0)
									} else {
										v44 = F_cstring_to_text(m, v42)
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = base.I64_extend_i32_u(v44)
											v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+28)))
											*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v48
											v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+32)))
											*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v50
											v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v20)+36)))
											*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v52
											v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+60))
											v59 = F_heap_form_tuple(m, v54, v5+int32(-48), v5+int32(-52))
											mBase = m.M
											v60 = m.ExcPending
											if v60 != 0 {
												return int64(0)
											} else {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
												v62 = F_HeapTupleHeaderGetDatum(m, v61)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return int64(0)
												} else {
													v64 = v62
													m.G0 = v7 - int32(-64)
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
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_brin_metapage_info_6), int32(0))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_brin_metapage_info_3), int32(357), int32(_a_F_brin_metapage_info_4))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
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
func F_brin_minmax_consistent(m *base.Module, l0 int32) int64 {
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
	var v18 int64
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
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
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
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+4)))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+6)))
	switch v21 - int32(1) {
	case 0, 1:
		v63 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, v21)
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int64(0)
		} else {
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
			v67 = F_FunctionCall2Coll(m, v63, v14, v66, v18)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
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
			return int64(0)
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
			v32 = F_FunctionCall2Coll(m, v26, v14, v31, v18)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int64(0)
			} else {
				if v32 == int64(0) {
					v69 = int64(0)
					m.G0 = v12 + int32(16)
					return v69
				} else {
					v37 = F_minmax_get_strategy_procinfo(m, v16, v20, v19, int32(4))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+8))
						v41 = F_FunctionCall2Coll(m, v37, v14, v40, v18)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
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
			return int64(0)
		} else {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+8))
			v47 = F_FunctionCall2Coll(m, v43, v14, v46, v18)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int64(0)
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
			return int64(0)
		} else {
			v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+6)))
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v53
			F_errmsg_internal(m, int32(_a_F_brin_minmax_consistent_0), v12)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_brin_minmax_consistent_1), int32(195), int32(_a_F_brin_minmax_consistent_2))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
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
func F_brin_minmax_multi_opcinfo(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14234(m, l0, int32(_a_F_brin_minmax_multi_opcinfo_0), int32(188))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_brin_minmax_multi_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_brin_minmax_multi_options_0), int32(_a_F_brin_minmax_multi_options_1), int32(32), int32(8), int32(256))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_brin_minmax_multi_summary_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_brin_minmax_multi_summary_in_0), int32(2982), int32(_a_F_brin_minmax_multi_summary_in_1), int32(_a_F_brin_minmax_multi_summary_in_2), int32(_a_F_brin_minmax_multi_summary_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v207 int64
	_ = v207
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	v3 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = F_palloc0(m, l0<<(uint(int32(3))%32)+int32(40))
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
	v169 = v27 + int32(40)
	v183 = v43
	v184 = int32(0)
	v186 = v153
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
	v332 = v184 + int32(1)
	if v332 != v51 {
		v183 = v327
		v184 = v332
		v186 = v329
		goto L41
	} else {
		goto L102
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
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
		goto L60
	} else {
		goto L61
	}
L47:
	;
	base.MemoryCopy(m, v19+int32(8), v183, v47)
	goto L49
L48:
	;
	goto L49
L49:
	;
	if base.I32_popcnt(v47) != int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169+v184<<(uint(int32(3))%32)))) = v224
	v327 = v183 + v47
	v329 = v186
	goto L43
L51:
	;
	v223 = int64(*(*int8)(unsafe.Add(mBase, uint32(v19)+8)))
	v224 = v223
	goto L50
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L57
	}
L53:
	;
	switch base.I32_ctz(v47) {
	case 0:
		goto L51
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	default:
		goto L52
	}
L54:
	;
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
	v224 = v209
	goto L50
L55:
	;
	v208 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+8)))
	v224 = v208
	goto L50
L56:
	;
	v207 = int64(*(*int16)(unsafe.Add(mBase, uint32(v19)+8)))
	v224 = v207
	goto L50
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v47
	F_errmsg_internal(m, int32(_a_F_brin_range_deserialize_3), v19)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_brin_range_deserialize_4), int32(123), int32(_a_F_brin_range_deserialize_5))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169+v184<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v186)
	if v47 != 0 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	switch v47&int32(_a_F_brin_range_deserialize_1) - int32(_a_F_brin_range_deserialize_2) {
	case 0:
		goto L67
	case 1:
		goto L68
	default:
		v327 = v183
		v329 = v186
		goto L43
	}
L63:
	;
	base.MemoryCopy(m, v186, v183, v47)
	goto L65
L64:
	;
	goto L65
L65:
	;
	v327 = v183 + v47
	v329 = v186 + (v47+int32(7))&int32(_a_F_brin_range_deserialize_0)
	goto L43
L66:
	;
	v306 = int32(8)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	v311 = base.B2i32(v309 == int32(18))
	if v309 == int32(18) {
		goto L90
	} else {
		goto L91
	}
L67:
	;
	v291 = F_strlen(m, v183)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v169+v184<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v186)
	v298 = v291 + int32(1)
	if v298 != 0 {
		goto L87
	} else {
		goto L88
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v169+v184<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v186)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v242 == int32(1) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v267 != 0 {
		goto L80
	} else {
		goto L81
	}
L70:
	;
	v246 = int32(18)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v248 == v246 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v259 = int32(1)
	if v242&v259 != 0 {
		v267 = int32(base.Ui32(v242) >> (uint(v259) % 32))
		goto L69
	} else {
		goto L79
	}
L73:
	;
	v251 = v246
	goto L75
L74:
	;
	v251 = int32(2)
	goto L75
L75:
	;
	if base.Ui32((v248-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v258 = int32(6)
	goto L78
L77:
	;
	v258 = v251
	goto L78
L78:
	;
	v267 = v258
	goto L69
L79:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v267 = int32(base.Ui32(v263) >> (uint(int32(2)) % 32))
	goto L69
L80:
	;
	base.MemoryCopy(m, v186, v183, v267)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v269 == int32(1) {
		goto L66
	} else {
		goto L83
	}
L83:
	;
	if v269&int32(1) != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v275 = int32(base.Ui32(v269) >> (uint(int32(1)) % 32))
	v327 = v183 + v275
	v329 = v186 + (v275+int32(7))&int32(248)
	goto L43
L85:
	;
	goto L86
L86:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v284 = int32(base.Ui32(v282) >> (uint(int32(2)) % 32))
	v327 = v183 + v284
	v329 = v186 + (v284+int32(7))&int32(2147483640)
	goto L43
L87:
	;
	base.MemoryCopy(m, v186, v183, v298)
	goto L89
L88:
	;
	goto L89
L89:
	;
	v327 = v183 + v298
	v329 = v186 + v291&int32(-8) + int32(8)
	goto L43
L90:
	;
	v312 = int32(24)
	goto L92
L91:
	;
	v312 = v306
	goto L92
L92:
	;
	v318 = base.B2i32(base.Ui32((v309-int32(1))&int32(255)) < base.Ui32(int32(3)))
	if base.Ui32((v309-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v319 = v306
	goto L95
L94:
	;
	v319 = v312
	goto L95
L95:
	;
	if v309 == int32(18) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v324 = int32(18)
	goto L98
L97:
	;
	v324 = int32(2)
	goto L98
L98:
	;
	if base.Ui32((v309-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v325 = int32(6)
	goto L101
L100:
	;
	v325 = v324
	goto L101
L101:
	;
	v327 = v183 + v325
	v329 = v186 + v319
	goto L43
L102:
	;
	goto L42
}
