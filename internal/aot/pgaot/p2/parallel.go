package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ExecParallelCreateReaders(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if v2 < v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_palloc(m, v8<<(uint(int32(2))%32))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v13
	v17 = v2
	goto L6
L6:
	;
	v23 = v17 << (uint(int32(2)) % 32)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23+v24)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v17<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34+v23)))
	v38 = F_palloc0(m, int32(4))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v36
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v41+v23))) = v38
	v45 = v17 + int32(1)
	if v45 != v8 {
		v17 = v45
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_ExecParallelEstimate(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
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
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	if l0 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v9 + int32(1)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v13 - int32(403) {
		case 0:
			v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+36)))
			if v186 != int32(1) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v192 = F_add_size(m, int32(20), v191)
				mBase = m.M
				v193 = m.ExcPending
				if v193 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v192
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)+36))
					v200 = F_add_size(m, v195, (v192+int32(31))&int32(-32))
					mBase = m.M
					v201 = m.ExcPending
					if v201 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v189)+36)) = v200
						v203 = *(*int32)(unsafe.Add(mBase, uint32(v189)+40))
						v205 = F_add_size(m, v203, int32(1))
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v189)+40)) = v205
							v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
							mBase = m.M
							v429 = m.ExcPending
							if v429 != 0 {
								return int32(0)
							} else {
								return v428
							}
						}
					}
				}
			}
		default:
			v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
			mBase = m.M
			v429 = m.ExcPending
			if v429 != 0 {
				return int32(0)
			} else {
				return v428
			}
		case 6:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
			if v17 == int32(1) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v24 = F_table_parallelscan_estimate(m, v21, v23)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v24
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
					v34 = F_add_size(m, v29, (v24+int32(31))&int32(-32))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v34
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
						v39 = F_add_size(m, v37, int32(1))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v39
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_ExecSeqScanInstrumentEstimate(m, l0, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				}
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ExecSeqScanInstrumentEstimate(m, l0, v44)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 8:
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+36)))
			if v48 == int32(1) {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
				v57 = F_index_parallelscan_estimate(m, v52, v53, v54, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v57
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
					v65 = F_add_size(m, v60, (v57+int32(31))&int32(-32))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v51)+36)) = v65
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
						v70 = F_add_size(m, v68, int32(1))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v51)+40)) = v70
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_ExecIndexOnlyScanInstrumentEstimate(m, l0, v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				}
			} else {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ExecIndexOnlyScanInstrumentEstimate(m, l0, v75)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 9:
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+36)))
			if v79 == int32(1) {
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
				v88 = F_index_parallelscan_estimate(m, v83, v84, v85, v87)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v88
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+36))
					v96 = F_add_size(m, v91, (v88+int32(31))&int32(-32))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v96
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
						v101 = F_add_size(m, v99, int32(1))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v101
							v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_ExecIndexOnlyScanInstrumentEstimate(m, l0, v106)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				}
			} else {
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ExecIndexOnlyScanInstrumentEstimate(m, l0, v106)
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 10:
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v110 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
				if v113 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(v109)+36))
					v123 = F_add_size(m, v116, (v113<<(uint(int32(3))%32)+int32(39))&int32(-32))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = v123
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v109)+40))
						v128 = F_add_size(m, v126, int32(1))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v109)+40)) = v128
							v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
							mBase = m.M
							v429 = m.ExcPending
							if v429 != 0 {
								return int32(0)
							} else {
								return v428
							}
						}
					}
				}
			}
		case 11:
			v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+36)))
			if v233 == int32(1) {
				v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+36))
				v239 = F_add_size(m, v237, int32(32))
				mBase = m.M
				v240 = m.ExcPending
				if v240 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v236)+36)) = v239
					v242 = *(*int32)(unsafe.Add(mBase, uint32(v236)+40))
					v244 = F_add_size(m, v242, int32(1))
					mBase = m.M
					v245 = m.ExcPending
					if v245 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v236)+40)) = v244
						v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v249 == int32(0) {
							v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
							mBase = m.M
							v429 = m.ExcPending
							if v429 != 0 {
								return int32(0)
							} else {
								return v428
							}
						} else {
							v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
							if v252 == int32(0) {
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							} else {
								v257 = F_mul_size(m, v252, int32(72))
								mBase = m.M
								v258 = m.ExcPending
								if v258 != 0 {
									return int32(0)
								} else {
									v259 = F_add_size(m, int32(8), v257)
									mBase = m.M
									v260 = m.ExcPending
									if v260 != 0 {
										return int32(0)
									} else {
										v261 = *(*int32)(unsafe.Add(mBase, uint32(v248)+36))
										v266 = F_add_size(m, v261, (v259+int32(31))&int32(-32))
										mBase = m.M
										v267 = m.ExcPending
										if v267 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v248)+36)) = v266
											v269 = *(*int32)(unsafe.Add(mBase, uint32(v248)+40))
											v271 = F_add_size(m, v269, int32(1))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v248)+40)) = v271
												v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
												mBase = m.M
												v429 = m.ExcPending
												if v429 != 0 {
													return int32(0)
												} else {
													return v428
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
				v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v249 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v252 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
					if v252 == int32(0) {
						v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
						mBase = m.M
						v429 = m.ExcPending
						if v429 != 0 {
							return int32(0)
						} else {
							return v428
						}
					} else {
						v257 = F_mul_size(m, v252, int32(72))
						mBase = m.M
						v258 = m.ExcPending
						if v258 != 0 {
							return int32(0)
						} else {
							v259 = F_add_size(m, int32(8), v257)
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return int32(0)
							} else {
								v261 = *(*int32)(unsafe.Add(mBase, uint32(v248)+36))
								v266 = F_add_size(m, v261, (v259+int32(31))&int32(-32))
								mBase = m.M
								v267 = m.ExcPending
								if v267 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v248)+36)) = v266
									v269 = *(*int32)(unsafe.Add(mBase, uint32(v248)+40))
									v271 = F_add_size(m, v269, int32(1))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v248)+40)) = v271
										v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
										mBase = m.M
										v429 = m.ExcPending
										if v429 != 0 {
											return int32(0)
										} else {
											return v428
										}
									}
								}
							}
						}
					}
				}
			}
		case 13:
			v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+36)))
			if v157 == int32(1) {
				v160 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
				v164 = F_table_parallelscan_estimate(m, v161, v163)
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v164
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v160)+36))
					v172 = F_add_size(m, v167, (v164+int32(31))&int32(-32))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v160)+36)) = v172
						v175 = *(*int32)(unsafe.Add(mBase, uint32(v160)+40))
						v177 = F_add_size(m, v175, int32(1))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v160)+40)) = v177
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							F_ExecSeqScanInstrumentEstimate(m, l0, v182)
							mBase = m.M
							v184 = m.ExcPending
							if v184 != 0 {
								return int32(0)
							} else {
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				}
			} else {
				v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				F_ExecSeqScanInstrumentEstimate(m, l0, v182)
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
					return int32(0)
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 21:
			v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+36)))
			if v133 != int32(1) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+148))
				if v138 != 0 {
					v139 = m.T0[v138].(func(*base.Module, int32, int32) int32)(m, l0, v136)
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v139
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v136)+36))
						v147 = F_add_size(m, v142, (v139+int32(31))&int32(-32))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v136)+36)) = v147
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)+40))
							v152 = F_add_size(m, v150, int32(1))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v136)+40)) = v152
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 22:
			v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+36)))
			if v209 != int32(1) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)+28))
				if v214 != 0 {
					v215 = m.T0[v214].(func(*base.Module, int32, int32) int32)(m, l0, v212)
					mBase = m.M
					v216 = m.ExcPending
					if v216 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v215
						v218 = *(*int32)(unsafe.Add(mBase, uint32(v212)+36))
						v223 = F_add_size(m, v218, (v215+int32(31))&int32(-32))
						mBase = m.M
						v224 = m.ExcPending
						if v224 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v212)+36)) = v223
							v226 = *(*int32)(unsafe.Add(mBase, uint32(v212)+40))
							v228 = F_add_size(m, v226, int32(1))
							mBase = m.M
							v229 = m.ExcPending
							if v229 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v212)+40)) = v228
								v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
								mBase = m.M
								v429 = m.ExcPending
								if v429 != 0 {
									return int32(0)
								} else {
									return v428
								}
							}
						}
					}
				} else {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				}
			}
		case 26:
			v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+36)))
			if v276 != int32(1) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v280 = *(*int32)(unsafe.Add(mBase, uint32(v279)+36))
				v282 = F_add_size(m, v280, int32(224))
				mBase = m.M
				v283 = m.ExcPending
				if v283 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v279)+36)) = v282
					v285 = *(*int32)(unsafe.Add(mBase, uint32(v279)+40))
					v287 = F_add_size(m, v285, int32(1))
					mBase = m.M
					v288 = m.ExcPending
					if v288 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v279)+40)) = v287
						v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
						mBase = m.M
						v429 = m.ExcPending
						if v429 != 0 {
							return int32(0)
						} else {
							return v428
						}
					}
				}
			}
		case 28:
			v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v399 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v402 = *(*int32)(unsafe.Add(mBase, uint32(v398)+12))
				if v402 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v406 = F_mul_size(m, v402, int32(40))
					mBase = m.M
					v407 = m.ExcPending
					if v407 != 0 {
						return int32(0)
					} else {
						v409 = F_add_size(m, v406, int32(8))
						mBase = m.M
						v410 = m.ExcPending
						if v410 != 0 {
							return int32(0)
						} else {
							v411 = *(*int32)(unsafe.Add(mBase, uint32(v398)+36))
							v416 = F_add_size(m, v411, (v409+int32(31))&int32(-32))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v398)+36)) = v416
								v419 = *(*int32)(unsafe.Add(mBase, uint32(v398)+40))
								v421 = F_add_size(m, v419, int32(1))
								mBase = m.M
								v422 = m.ExcPending
								if v422 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v398)+40)) = v421
									v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										return v428
									}
								}
							}
						}
					}
				}
			}
		case 29:
			v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v318 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v321 = *(*int32)(unsafe.Add(mBase, uint32(v317)+12))
				if v321 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v325 = F_mul_size(m, v321, int32(16))
					mBase = m.M
					v326 = m.ExcPending
					if v326 != 0 {
						return int32(0)
					} else {
						v328 = F_add_size(m, v325, int32(8))
						mBase = m.M
						v329 = m.ExcPending
						if v329 != 0 {
							return int32(0)
						} else {
							v330 = *(*int32)(unsafe.Add(mBase, uint32(v317)+36))
							v335 = F_add_size(m, v330, (v328+int32(31))&int32(-32))
							mBase = m.M
							v336 = m.ExcPending
							if v336 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v317)+36)) = v335
								v338 = *(*int32)(unsafe.Add(mBase, uint32(v317)+40))
								v340 = F_add_size(m, v338, int32(1))
								mBase = m.M
								v341 = m.ExcPending
								if v341 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v317)+40)) = v340
									v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										return v428
									}
								}
							}
						}
					}
				}
			}
		case 30:
			v344 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v345 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
				if v348 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v352 = F_mul_size(m, v348, int32(96))
					mBase = m.M
					v353 = m.ExcPending
					if v353 != 0 {
						return int32(0)
					} else {
						v355 = F_add_size(m, v352, int32(8))
						mBase = m.M
						v356 = m.ExcPending
						if v356 != 0 {
							return int32(0)
						} else {
							v357 = *(*int32)(unsafe.Add(mBase, uint32(v344)+36))
							v362 = F_add_size(m, v357, (v355+int32(31))&int32(-32))
							mBase = m.M
							v363 = m.ExcPending
							if v363 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v344)+36)) = v362
								v365 = *(*int32)(unsafe.Add(mBase, uint32(v344)+40))
								v367 = F_add_size(m, v365, int32(1))
								mBase = m.M
								v368 = m.ExcPending
								if v368 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v344)+40)) = v367
									v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										return v428
									}
								}
							}
						}
					}
				}
			}
		case 32:
			v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v372 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v375 = *(*int32)(unsafe.Add(mBase, uint32(v371)+12))
				if v375 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v379 = F_mul_size(m, v375, int32(24))
					mBase = m.M
					v380 = m.ExcPending
					if v380 != 0 {
						return int32(0)
					} else {
						v382 = F_add_size(m, v379, int32(8))
						mBase = m.M
						v383 = m.ExcPending
						if v383 != 0 {
							return int32(0)
						} else {
							v384 = *(*int32)(unsafe.Add(mBase, uint32(v371)+36))
							v389 = F_add_size(m, v384, (v382+int32(31))&int32(-32))
							mBase = m.M
							v390 = m.ExcPending
							if v390 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v371)+36)) = v389
								v392 = *(*int32)(unsafe.Add(mBase, uint32(v371)+40))
								v394 = F_add_size(m, v392, int32(1))
								mBase = m.M
								v395 = m.ExcPending
								if v395 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v371)+40)) = v394
									v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										return v428
									}
								}
							}
						}
					}
				}
			}
		case 37:
			v290 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v291 == int32(0) {
				v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
				mBase = m.M
				v429 = m.ExcPending
				if v429 != 0 {
					return int32(0)
				} else {
					return v428
				}
			} else {
				v294 = *(*int32)(unsafe.Add(mBase, uint32(v290)+12))
				if v294 == int32(0) {
					v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
					mBase = m.M
					v429 = m.ExcPending
					if v429 != 0 {
						return int32(0)
					} else {
						return v428
					}
				} else {
					v298 = F_mul_size(m, v294, int32(20))
					mBase = m.M
					v299 = m.ExcPending
					if v299 != 0 {
						return int32(0)
					} else {
						v301 = F_add_size(m, v298, int32(4))
						mBase = m.M
						v302 = m.ExcPending
						if v302 != 0 {
							return int32(0)
						} else {
							v303 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
							v308 = F_add_size(m, v303, (v301+int32(31))&int32(-32))
							mBase = m.M
							v309 = m.ExcPending
							if v309 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v290)+36)) = v308
								v311 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
								v313 = F_add_size(m, v311, int32(1))
								mBase = m.M
								v314 = m.ExcPending
								if v314 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v290)+40)) = v313
									v428 = F_planstate_tree_walker_impl(m, l0, int32(672), l1)
									mBase = m.M
									v429 = m.ExcPending
									if v429 != 0 {
										return int32(0)
									} else {
										return v428
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
func F_ExecParallelRetrieveInstrumentation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v187 int64
	_ = v187
	var v188 int64
	_ = v188
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v3 < v14 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v62 = l1 + v56 + v58*v21*int32(440)
	if int32(0) < v58 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v21 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(16)+v21<<(uint(int32(2))%32))))
	if v29 == v13 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v32 = v21 + int32(1)
	if v32 != v14 {
		v21 = v32
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v13
	F_errmsg_internal(m, int32(_a_F_ExecParallelRetrieveInstrumentation_0), v10)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_ExecParallelRetrieveInstrumentation_1), int32(1109), int32(_a_F_ExecParallelRetrieveInstrumentation_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v68 = int32(0)
	goto L16
L14:
	;
	v202 = v58
	goto L15
L15:
	;
	v206 = int32(_a_F_ExecParallelRetrieveInstrumentation_3)
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0])) = v210
	v213 = F_mul_size(m, v202, int32(440))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L26
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = v62 + v68*int32(440)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v73)+392))
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)+392))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+392)) = v77 + v78
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v73)+184))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)+184))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+184)) = v81 + v82
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v76)+400))
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v73)+400))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+400)) = base.F64_add(v85, v86)
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v76)+408))
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v73)+408))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+408)) = base.F64_add(v89, v90)
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v76)+416))
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v73)+416))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+416)) = base.F64_add(v93, v94)
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v76)+424))
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v73)+424))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+424)) = base.F64_add(v97, v98)
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v76)+432))
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v73)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+432)) = base.F64_add(v101, v102)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v105 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v202 = v197
	goto L15
L18:
	;
	v196 = v68 + int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v196 < v197 {
		v68 = v196
		goto L16
	} else {
		goto L25
	}
L19:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v73)+192))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v76)+192))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+192)) = v108 + v109
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v73)+200))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v76)+200))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+200)) = v112 + v113
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v73)+208))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v76)+208))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+208)) = v116 + v117
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v73)+216))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v76)+216))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+216)) = v120 + v121
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v73)+224))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v76)+224))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+224)) = v124 + v125
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v73)+232))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v76)+232))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+232)) = v128 + v129
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v73)+240))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v76)+240))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+240)) = v132 + v133
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v73)+248))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v76)+248))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+248)) = v136 + v137
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v73)+256))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v76)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+256)) = v140 + v141
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v73)+264))
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v76)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+264)) = v144 + v145
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v73)+272))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v76)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+272)) = v148 + v149
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v73)+280))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v76)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+280)) = v152 + v153
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v73)+288))
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v76)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+288)) = v156 + v157
	v160 = *(*int64)(unsafe.Add(mBase, uint32(v73)+296))
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v76)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+296)) = v160 + v161
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v73)+304))
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v76)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+304)) = v164 + v165
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v73)+312))
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v76)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+312)) = v168 + v169
	goto L21
L20:
	;
	goto L21
L21:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)))
	if v172 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v73)+336))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v76)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+336)) = v175 + v176
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v73)+320))
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v76)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+320)) = v179 + v180
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v73)+328))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v76)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+328)) = v183 + v184
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v73)+344))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v76)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+344)) = v187 + v188
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v73)+352))
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v76)+352))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+352)) = v191 + v192
	goto L24
L23:
	;
	goto L24
L24:
	;
	goto L18
L25:
	;
	goto L17
L26:
	;
	v217 = F_palloc(m, v213+int32(8))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v217
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0])) = v207
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v222
	if v213 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v224+int32(8), v62, v213)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v228 - int32(409) {
	case 0:
		goto L33
	default:
		goto L31
	case 2:
		goto L42
	case 3:
		goto L41
	case 4:
		goto L40
	case 5:
		goto L34
	case 7:
		goto L32
	case 22:
		goto L35
	case 23:
		goto L39
	case 24:
		goto L38
	case 26:
		goto L36
	case 31:
		goto L37
	}
L31:
	;
	v371 = F_planstate_tree_walker_impl(m, l0, int32(675), l1)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L9
	} else {
		goto L93
	}
L32:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v351 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L33:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v335 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L34:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	if v319 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L35:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v306 != 0 {
		goto L74
	} else {
		goto L75
	}
L36:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v293 != 0 {
		goto L67
	} else {
		goto L68
	}
L37:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v277 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L38:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v264 != 0 {
		goto L56
	} else {
		goto L57
	}
L39:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v251 != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	F_ExecBitmapIndexScanRetrieveInstrumentation(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L48
	}
L41:
	;
	F_ExecBitmapIndexScanRetrieveInstrumentation(m, l0)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L47
	}
L42:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v231 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	goto L31
L44:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	v238 = v234<<(uint(int32(3))%32) + int32(8)
	v239 = F_palloc(m, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v239
	if v238 == int32(0) {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	base.MemoryCopy(m, v239, v231, v238)
	goto L43
L47:
	;
	goto L31
L48:
	;
	goto L31
L49:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v256 = v252<<(uint(int32(4))%32) | int32(8)
	v257 = F_palloc(m, v256)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L31
L52:
	;
	if v256 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	base.MemoryCopy(m, v257, v259, v256)
	goto L55
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v257
	goto L51
L56:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v269 = v265*int32(96) | int32(8)
	v270 = F_palloc(m, v269)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L31
L59:
	;
	if v269 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	base.MemoryCopy(m, v270, v272, v269)
	goto L62
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v270
	goto L58
L63:
	;
	goto L31
L64:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v284 = v280*int32(20) + int32(4)
	v285 = F_palloc(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v285
	if v284 == int32(0) {
		goto L63
	} else {
		goto L66
	}
L66:
	;
	base.MemoryCopy(m, v285, v277, v284)
	goto L63
L67:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v298 = v294*int32(24) + int32(8)
	v299 = F_palloc(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L31
L70:
	;
	if v298 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	base.MemoryCopy(m, v299, v301, v298)
	goto L73
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v299
	goto L69
L74:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v311 = v307*int32(40) + int32(8)
	v312 = F_palloc(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L9
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L31
L77:
	;
	if v311 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v312, v314, v311)
	goto L80
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v312
	goto L76
L81:
	;
	goto L31
L82:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v326 = v322*int32(72) + int32(8)
	v327 = F_palloc(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v327
	if v326 == int32(0) {
		goto L81
	} else {
		goto L84
	}
L84:
	;
	base.MemoryCopy(m, v327, v319, v326)
	goto L81
L85:
	;
	goto L31
L86:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v342 = v338*int32(56) + int32(8)
	v343 = F_palloc(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v343
	if v342 == int32(0) {
		goto L85
	} else {
		goto L88
	}
L88:
	;
	base.MemoryCopy(m, v343, v335, v342)
	goto L85
L89:
	;
	goto L31
L90:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)))
	v358 = v354*int32(56) + int32(8)
	v359 = F_palloc(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v359
	if v358 == int32(0) {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	base.MemoryCopy(m, v359, v351, v358)
	goto L89
L93:
	;
	m.G0 = v10 + int32(16)
	return v371
}
func F_ParallelApplyWorkerMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
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
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	v9 = m.G0
	v11 = v9 - int32(144)
	m.G0 = v11
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[0])) = uint8(v14)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v25 = int32(967)
	switch v25 {
	case 0, 2:
		goto L2
	default:
		goto L3
	}
L1:
	;
	v62 = m.G0
	v64 = v62 - int32(32)
	m.G0 = v64
	v67 = int32(969)
	switch v67 {
	case 0, 2:
		goto L12
	default:
		goto L13
	}
L2:
	;
	F_sigemptyset(m, v22+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(268435456)
	switch v25 {
	case 0:
		goto L7
	default:
		goto L5
	case 2:
		goto L6
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[1])) = int32(965)
	goto L2
L4:
	;
	goto L9
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(_a_F_ParallelApplyWorkerMain_0)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
	goto L4
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(-2)
	goto L4
L9:
	;
	goto L10
L10:
	;
	v54 = F___sigaction(m, v14, v22+int32(12), int32(0))
	mBase = m.M
	m.G0 = v22 + int32(32)
	goto L1
L11:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	F_sigemptyset(m, v64+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(268435456)
	switch v67 {
	case 0:
		goto L17
	default:
		goto L15
	case 2:
		goto L16
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[2])) = int32(967)
	goto L12
L14:
	;
	goto L19
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(_a_F_ParallelApplyWorkerMain_0)
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(0)
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(-2)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v96 = F___sigaction(m, int32(12), v64+int32(12), int32(0))
	mBase = m.M
	m.G0 = v64 + int32(32)
	goto L11
L21:
	;
	return
L22:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[3]))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+1336))
	v105 = F_dsm_attach(m, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L21
	} else {
		goto L27
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L21
	} else {
		goto L133
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L21
	} else {
		goto L130
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L21
	} else {
		goto L127
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L21
	} else {
		goto L123
	}
L27:
	;
	if v105 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v108)))
	if v110 == int64(2021433447) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L21
	} else {
		goto L119
	}
L31:
	;
	if v112 == int32(0) {
		goto L26
	} else {
		goto L35
	}
L32:
	;
	v112 = v108
	goto L34
L33:
	;
	v112 = int32(0)
	goto L34
L34:
	;
	goto L31
L35:
	;
	v118 = F_shm_toc_lookup(m, v112, int64(1), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4])) = v118
	v123 = F_shm_toc_lookup(m, v112, int64(2), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[5]))
	F_shm_mq_set_receiver(m, v123, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L21
	} else {
		goto L38
	}
L38:
	;
	v129 = F_shm_mq_attach(m, v123, v105)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L21
	} else {
		goto L39
	}
L39:
	;
	v131 = base.I32_wrap_i64(l0)
	F_logicalrep_worker_attach(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L21
	} else {
		goto L40
	}
L40:
	;
	F_before_shmem_exit(m, int32(1050), base.I64_extend_i32_u(v105))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L41
	}
L41:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v142 = base.AtomicRmwXchg32(m, v139, int32(0), int32(1))
	if v142 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_s_lock(m, v139, int32(_a_F_ParallelApplyWorkerMain_1))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L21
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+18)))
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+16)) = v131
	*(*uint16)(unsafe.Add(mBase, uint32(v150)+12)) = uint16(v148)
	v153 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v150))), uint32(v153))
	v158 = F_shm_toc_lookup(m, v112, int64(3), v153)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L21
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[5]))
	F_shm_mq_set_sender(m, v158, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L21
	} else {
		goto L47
	}
L47:
	;
	v164 = F_shm_mq_attach(m, v158, v105)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_pq_redirect_to_shm_mq(m, v105, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[7])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[8])) = v170
	goto L50
L50:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v178 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v177)+88)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v177)+112)) = v178
	*(*int64)(unsafe.Add(mBase, uint32(v177)+96)) = v178
	F_InitializeLogRepWorker(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	v187 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[0])) = uint8(v187)
	F_StartTransactionCommand(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[9]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v196 = v11 + int32(32)
	F_ReplicationOriginNameForLogicalRep(m, v193, int32(0), v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	v200 = F_replorigin_by_name(m, v196, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L21
	} else {
		goto L54
	}
L54:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+64))
	F_replorigin_session_setup(m, v200, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[10])) = uint16(v200)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	F_CacheRegisterSyscacheCallback(m, int32(68), int32(1051), int64(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	F_set_apply_error_context_origin(m, v196)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[11]))
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[12]))
	v227 = F_AllocSetContextCreateInternal(m, v222, int32(_a_F_ParallelApplyWorkerMain_2), int32(0), int32(_a_F_ParallelApplyWorkerMain_3), int32(_a_F_ParallelApplyWorkerMain_4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[13])) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v11)+136)) = int32(1052)
	v232 = int32(_a_F_ParallelApplyWorkerMain_5)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14])) = v11 + int32(132)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v233
	goto L60
L60:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[15]))
	if v248 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L21
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	if v252 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L64
L66:
	;
	v255 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[17]))
	if v275 != 0 {
		goto L76
	} else {
		goto L77
	}
L69:
	;
	if v255 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[9]))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v259
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_6), v11+int32(16))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L21
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L21
	} else {
		goto L75
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(724), int32(_a_F_ParallelApplyWorkerMain_8))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L21
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[17])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L21
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[11])) = v284
	v291 = F_shm_mq_receive(m, v129, v11+int32(124), v11+int32(128), int32(1))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L21
	} else {
		goto L83
	}
L79:
	;
	goto L78
L80:
	;
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[13]))
	F_MemoryContextReset(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L21
	} else {
		goto L118
	}
L81:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v317 = base.AtomicRmwXchg32(m, v314, int32(0), int32(1))
	if v317 != 0 {
		goto L88
	} else {
		goto L89
	}
L82:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	if v293 == int32(0) {
		goto L25
	} else {
		goto L84
	}
L83:
	;
	switch v291 {
	case 0:
		goto L82
	case 1:
		goto L81
	default:
		goto L23
	}
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+116)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v293
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+108)) = v299
	v302 = v11 + int32(108)
	v303 = F_pq_getmsgbyte(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L21
	} else {
		goto L85
	}
L85:
	;
	if v303 != int32(119) {
		goto L24
	} else {
		goto L86
	}
L86:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+120)) = v307 + int32(24)
	F_apply_dispatch(m, v302)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L21
	} else {
		goto L87
	}
L87:
	;
	goto L80
L88:
	;
	F_s_lock(m, v314, int32(_a_F_ParallelApplyWorkerMain_1))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L21
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+32))
	v324 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v322))), uint32(v324))
	switch v323 {
	case 0:
		goto L92
	case 1:
		goto L94
	default:
		v361 = v323
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[18]))
	v403 = F_WaitLatch(m, v399, int32(41), int32(1000), int32(83886089))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L21
	} else {
		goto L109
	}
L93:
	;
	switch v361 - int32(2) {
	case 0:
		goto L104
	case 1:
		goto L103
	default:
		goto L80
	}
L94:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	F_LockApplyTransactionForSession(m, v329, v332, int32(0), int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L21
	} else {
		goto L95
	}
L95:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6]))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+32))
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	F_UnlockApplyTransactionForSession(m, v339, v342, int32(0), int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L21
	} else {
		goto L96
	}
L96:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v351 = base.AtomicRmwXchg32(m, v348, int32(0), int32(1))
	if v351 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	F_s_lock(m, v348, int32(_a_F_ParallelApplyWorkerMain_1))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L21
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)+32))
	v358 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v356))), uint32(v358))
	v361 = v357
	goto L93
L100:
	;
	goto L99
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v392)+32)) = v393
	v395 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v392))), uint32(v395))
	goto L80
L102:
	;
	F_s_lock(m, v387, int32(_a_F_ParallelApplyWorkerMain_1))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L21
	} else {
		goto L108
	}
L103:
	;
	v371 = int32(0)
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	F_apply_spooled_messages(m, v373+int32(36), v376, int64(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L21
	} else {
		goto L106
	}
L104:
	;
	v365 = int32(3)
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v370 = base.AtomicRmwXchg32(m, v367, int32(0), int32(1))
	if v370 != 0 {
		v387 = v367
		v388 = v365
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v392 = v367
	v393 = v365
	goto L101
L106:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4]))
	v383 = int32(0)
	v384 = base.AtomicRmwXchg32(m, v381, v383, int32(1))
	if v384 == v383 {
		v392 = v381
		v393 = v371
		goto L101
	} else {
		goto L107
	}
L107:
	;
	v387 = v381
	v388 = v371
	goto L102
L108:
	;
	v392 = v387
	v393 = v388
	goto L101
L109:
	;
	if v403&int32(1) != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[18]))
	v409 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = v409
	v414 = base.AtomicRmwOr32(m, v409, int32(_a_F_ParallelApplyWorkerMain_9), v409)
	goto L113
L111:
	;
	goto L112
L112:
	;
	if v403&int32(8) == int32(0) {
		goto L80
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[19]))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+20))
	goto L115
L115:
	;
	if v421 == int32(2) {
		goto L80
	} else {
		goto L116
	}
L116:
	;
	v425 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L21
	} else {
		goto L117
	}
L117:
	;
	goto L80
L118:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[11])) = v219
	goto L60
L119:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_10), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(909), int32(_a_F_ParallelApplyWorkerMain_11))
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L21
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L21
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_12), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L21
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(915), int32(_a_F_ParallelApplyWorkerMain_11))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L21
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelApplyWorkerMain_13), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L21
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(778), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L21
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v303
	F_errmsg_internal(m, int32(_a_F_ParallelApplyWorkerMain_15), v11)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L21
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(788), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L21
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L21
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_16), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L21
	} else {
		goto L135
	}
L135:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_7), int32(835), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_parallel_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+96)))
	if v12 != int32(115) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v74
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	v19 = int32(_a_F_is_parallel_safe_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+8)) = uint16(v19)
	v21 = l0
	v26 = int32(0)
	goto L5
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+72))
	if v15 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v74 = int32(1)
	goto L1
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	if v27 == int32(0) {
		v60 = v26
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v64 = F_max_parallel_hazard_walker(m, l1, v9+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L16
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v61 != 0 {
		v21 = v61
		v26 = v60
		goto L5
	} else {
		goto L15
	}
L8:
	;
	v30 = int32(0)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v31 <= v30 {
		v60 = v26
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v38 = v30
	v39 = v26
	goto L10
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v38<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+40))
	v46 = F_list_concat(m, v39, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v60 = v46
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v46
	v52 = v38 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v52 < v53 {
		v38 = v52
		v39 = v46
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	goto L6
L16:
	;
	v74 = v64 ^ int32(1)
	goto L1
}
func F_parallel_vacuum_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
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
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	switch v11 - int32(1) {
	case 0:
		v15 = int32(_a_F_parallel_vacuum_error_callback_0)
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
			F_errcontext_msg(m, v15, v8)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	case 1:
		v15 = int32(_a_F_parallel_vacuum_error_callback_1)
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v19
			F_errcontext_msg(m, v15, v8)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	default:
		m.G0 = v8 + int32(16)
		return
	}
}
func F_parallel_vacuum_update_shared_delay_params(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[0]))
	if v9 == int32(0) {
		m.G0 = v6 + int32(32)
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[1]))
		if v12 == v14 {
			m.G0 = v6 + int32(32)
			return
		} else {
			v18 = base.AtomicRmwXchg32(m, v9, int32(4), int32(1))
			if v18 != 0 {
				F_s_lock(m, v9+int32(4), int32(_a_F_parallel_vacuum_update_shared_delay_params_0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[0]))
					v27 = *(*float64)(unsafe.Add(mBase, uint32(v26)+8))
					*(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[2])) = v27
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
					*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[3])) = v30
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[4])) = v33
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[5])) = v36
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
					*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[6])) = v39
					v41 = int32(0)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+4)), uint32(v41))
					F_VacuumUpdateCosts(m)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[1])) = v12
						v50 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							if v50 == int32(0) {
								m.G0 = v6 + int32(32)
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[6]))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v55
								v58 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[4]))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v58
								v61 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[5]))
								*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v61
								v64 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[7]))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v64
								v67 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[8]))
								*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v67
								F_errmsg_internal(m, int32(_a_F_parallel_vacuum_update_shared_delay_params_1), v6)
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_parallel_vacuum_update_shared_delay_params_2), int32(686), int32(_a_F_parallel_vacuum_update_shared_delay_params_3))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return
									} else {
										m.G0 = v6 + int32(32)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[0]))
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v26)+8))
				*(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[2])) = v27
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[3])) = v30
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
				*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[4])) = v33
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[5])) = v36
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
				*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[6])) = v39
				v41 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+4)), uint32(v41))
				F_VacuumUpdateCosts(m)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[1])) = v12
					v50 = F_errstart(m, int32(13), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						if v50 == int32(0) {
							m.G0 = v6 + int32(32)
							return
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[6]))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v55
							v58 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[4]))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v58
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[5]))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v61
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v64
							v67 = *(*float64)(unsafe.Add(mBase, _c_F_parallel_vacuum_update_shared_delay_params[8]))
							*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v67
							F_errmsg_internal(m, int32(_a_F_parallel_vacuum_update_shared_delay_params_1), v6)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_parallel_vacuum_update_shared_delay_params_2), int32(686), int32(_a_F_parallel_vacuum_update_shared_delay_params_3))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									m.G0 = v6 + int32(32)
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
