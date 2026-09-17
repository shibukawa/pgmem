package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v10 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v14 - int32(397) {
		case 0:
			v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+36)))
			if v159 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v165 = F_add_size(m, int32(20), v164)
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v165
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v162)+36))
					v173 = F_add_size(m, v168, (v165+int32(31))&int32(-32))
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v162)+36)) = v173
						v176 = *(*int32)(unsafe.Add(mBase, uint32(v162)+40))
						v178 = F_add_size(m, v176, int32(1))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v162)+40)) = v178
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				}
			}
		default:
			v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
			mBase = m.M
			v396 = m.ExcPending
			if v396 != 0 {
				return int32(0)
			} else {
				return v395
			}
		case 6:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+36)))
			if v18 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
				v25 = F_table_parallelscan_estimate(m, v22, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v25
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
					v35 = F_add_size(m, v30, (v25+int32(31))&int32(-32))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v35
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
						v40 = F_add_size(m, v38, int32(1))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v40
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				}
			}
		case 8:
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+36)))
			if v44|v46&int32(1) != 0 {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
				v60 = F_index_parallelscan_estimate(m, v50, v51, v52, v54, base.B2i32(v44 != int32(0)), v46&int32(1), v59)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v60
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v43)+36))
					v68 = F_add_size(m, v63, (v60+int32(31))&int32(-32))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v43)+36)) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
						v73 = F_add_size(m, v71, int32(1))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v43)+40)) = v73
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				}
			} else {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			}
		case 9:
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+36)))
			if v78|v80&int32(1) != 0 {
				v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
				v94 = F_index_parallelscan_estimate(m, v84, v85, v86, v88, base.B2i32(v78 != int32(0)), v80&int32(1), v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v94
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v77)+36))
					v102 = F_add_size(m, v97, (v94+int32(31))&int32(-32))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v77)+36)) = v102
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v77)+40))
						v107 = F_add_size(m, v105, int32(1))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v77)+40)) = v107
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				}
			} else {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			}
		case 10:
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v112 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
				if v115 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
					v125 = F_add_size(m, v118, (v115<<(uint(int32(3))%32)+int32(39))&int32(-32))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v111)+36)) = v125
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v111)+40))
						v130 = F_add_size(m, v128, int32(1))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v111)+40)) = v130
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				}
			}
		case 11:
			v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+36)))
			if v206 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v210 = int32(32)
				v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v211 == int32(0) {
					v231 = v210
					v232 = *(*int32)(unsafe.Add(mBase, uint32(v209)+36))
					v233 = F_add_size(m, v232, v231)
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v209)+36)) = v233
						v236 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
						v238 = F_add_size(m, v236, int32(1))
						mBase = m.M
						v239 = m.ExcPending
						if v239 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v209)+40)) = v238
							v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
							mBase = m.M
							v396 = m.ExcPending
							if v396 != 0 {
								return int32(0)
							} else {
								return v395
							}
						}
					}
				} else {
					v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
					if v214 <= int32(0) {
						v231 = v210
						v232 = *(*int32)(unsafe.Add(mBase, uint32(v209)+36))
						v233 = F_add_size(m, v232, v231)
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v209)+36)) = v233
							v236 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
							v238 = F_add_size(m, v236, int32(1))
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v209)+40)) = v238
								v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v396 = m.ExcPending
								if v396 != 0 {
									return int32(0)
								} else {
									return v395
								}
							}
						}
					} else {
						v219 = F_add_size(m, int32(24), int32(8))
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return int32(0)
						} else {
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
							v223 = F_mul_size(m, v221, int32(16))
							mBase = m.M
							v224 = m.ExcPending
							if v224 != 0 {
								return int32(0)
							} else {
								v225 = F_add_size(m, v219, v223)
								mBase = m.M
								v226 = m.ExcPending
								if v226 != 0 {
									return int32(0)
								} else {
									v231 = (v225 + int32(31)) & int32(-32)
									v232 = *(*int32)(unsafe.Add(mBase, uint32(v209)+36))
									v233 = F_add_size(m, v232, v231)
									mBase = m.M
									v234 = m.ExcPending
									if v234 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v209)+36)) = v233
										v236 = *(*int32)(unsafe.Add(mBase, uint32(v209)+40))
										v238 = F_add_size(m, v236, int32(1))
										mBase = m.M
										v239 = m.ExcPending
										if v239 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v209)+40)) = v238
											v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
											mBase = m.M
											v396 = m.ExcPending
											if v396 != 0 {
												return int32(0)
											} else {
												return v395
											}
										}
									}
								}
							}
						}
					}
				}
			}
		case 21:
			v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+36)))
			if v135 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v138 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+144))
				if v140 != 0 {
					v141 = m.T0[v140].(func(*base.Module, int32, int32) int32)(m, l0, v138)
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v141
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+36))
						v149 = F_add_size(m, v144, (v141+int32(31))&int32(-32))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v138)+36)) = v149
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+40))
							v154 = F_add_size(m, v152, int32(1))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v138)+40)) = v154
								v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v396 = m.ExcPending
								if v396 != 0 {
									return int32(0)
								} else {
									return v395
								}
							}
						}
					}
				} else {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				}
			}
		case 22:
			v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+36)))
			if v182 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+28))
				if v187 != 0 {
					v188 = m.T0[v187].(func(*base.Module, int32, int32) int32)(m, l0, v185)
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+124)) = v188
						v191 = *(*int32)(unsafe.Add(mBase, uint32(v185)+36))
						v196 = F_add_size(m, v191, (v188+int32(31))&int32(-32))
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v185)+36)) = v196
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v185)+40))
							v201 = F_add_size(m, v199, int32(1))
							mBase = m.M
							v202 = m.ExcPending
							if v202 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v185)+40)) = v201
								v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
								mBase = m.M
								v396 = m.ExcPending
								if v396 != 0 {
									return int32(0)
								} else {
									return v395
								}
							}
						}
					}
				} else {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				}
			}
		case 26:
			v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+36)))
			if v242 != int32(1) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+36))
				v248 = F_add_size(m, v246, int32(224))
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v245)+36)) = v248
					v251 = *(*int32)(unsafe.Add(mBase, uint32(v245)+40))
					v253 = F_add_size(m, v251, int32(1))
					mBase = m.M
					v254 = m.ExcPending
					if v254 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v245)+40)) = v253
						v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
						mBase = m.M
						v396 = m.ExcPending
						if v396 != 0 {
							return int32(0)
						} else {
							return v395
						}
					}
				}
			}
		case 28:
			v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v365 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)+12))
				if v368 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v372 = F_mul_size(m, v368, int32(40))
					mBase = m.M
					v373 = m.ExcPending
					if v373 != 0 {
						return int32(0)
					} else {
						v375 = F_add_size(m, v372, int32(8))
						mBase = m.M
						v376 = m.ExcPending
						if v376 != 0 {
							return int32(0)
						} else {
							v377 = *(*int32)(unsafe.Add(mBase, uint32(v364)+36))
							v382 = F_add_size(m, v377, (v375+int32(31))&int32(-32))
							mBase = m.M
							v383 = m.ExcPending
							if v383 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v364)+36)) = v382
								v385 = *(*int32)(unsafe.Add(mBase, uint32(v364)+40))
								v387 = F_add_size(m, v385, int32(1))
								mBase = m.M
								v388 = m.ExcPending
								if v388 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v364)+40)) = v387
									v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										return v395
									}
								}
							}
						}
					}
				}
			}
		case 29:
			v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v284 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+12))
				if v287 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v291 = F_mul_size(m, v287, int32(16))
					mBase = m.M
					v292 = m.ExcPending
					if v292 != 0 {
						return int32(0)
					} else {
						v294 = F_add_size(m, v291, int32(8))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int32(0)
						} else {
							v296 = *(*int32)(unsafe.Add(mBase, uint32(v283)+36))
							v301 = F_add_size(m, v296, (v294+int32(31))&int32(-32))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v283)+36)) = v301
								v304 = *(*int32)(unsafe.Add(mBase, uint32(v283)+40))
								v306 = F_add_size(m, v304, int32(1))
								mBase = m.M
								v307 = m.ExcPending
								if v307 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v283)+40)) = v306
									v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										return v395
									}
								}
							}
						}
					}
				}
			}
		case 30:
			v310 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v311 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v314 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
				if v314 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v318 = F_mul_size(m, v314, int32(96))
					mBase = m.M
					v319 = m.ExcPending
					if v319 != 0 {
						return int32(0)
					} else {
						v321 = F_add_size(m, v318, int32(8))
						mBase = m.M
						v322 = m.ExcPending
						if v322 != 0 {
							return int32(0)
						} else {
							v323 = *(*int32)(unsafe.Add(mBase, uint32(v310)+36))
							v328 = F_add_size(m, v323, (v321+int32(31))&int32(-32))
							mBase = m.M
							v329 = m.ExcPending
							if v329 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v310)+36)) = v328
								v331 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
								v333 = F_add_size(m, v331, int32(1))
								mBase = m.M
								v334 = m.ExcPending
								if v334 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v310)+40)) = v333
									v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										return v395
									}
								}
							}
						}
					}
				}
			}
		case 32:
			v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v338 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
				if v341 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v345 = F_mul_size(m, v341, int32(24))
					mBase = m.M
					v346 = m.ExcPending
					if v346 != 0 {
						return int32(0)
					} else {
						v348 = F_add_size(m, v345, int32(8))
						mBase = m.M
						v349 = m.ExcPending
						if v349 != 0 {
							return int32(0)
						} else {
							v350 = *(*int32)(unsafe.Add(mBase, uint32(v337)+36))
							v355 = F_add_size(m, v350, (v348+int32(31))&int32(-32))
							mBase = m.M
							v356 = m.ExcPending
							if v356 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v337)+36)) = v355
								v358 = *(*int32)(unsafe.Add(mBase, uint32(v337)+40))
								v360 = F_add_size(m, v358, int32(1))
								mBase = m.M
								v361 = m.ExcPending
								if v361 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v337)+40)) = v360
									v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										return v395
									}
								}
							}
						}
					}
				}
			}
		case 37:
			v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v257 == int32(0) {
				v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
				mBase = m.M
				v396 = m.ExcPending
				if v396 != 0 {
					return int32(0)
				} else {
					return v395
				}
			} else {
				v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
				if v260 == int32(0) {
					v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
					mBase = m.M
					v396 = m.ExcPending
					if v396 != 0 {
						return int32(0)
					} else {
						return v395
					}
				} else {
					v264 = F_mul_size(m, v260, int32(20))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						v267 = F_add_size(m, v264, int32(4))
						mBase = m.M
						v268 = m.ExcPending
						if v268 != 0 {
							return int32(0)
						} else {
							v269 = *(*int32)(unsafe.Add(mBase, uint32(v256)+36))
							v274 = F_add_size(m, v269, (v267+int32(31))&int32(-32))
							mBase = m.M
							v275 = m.ExcPending
							if v275 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v256)+36)) = v274
								v277 = *(*int32)(unsafe.Add(mBase, uint32(v256)+40))
								v279 = F_add_size(m, v277, int32(1))
								mBase = m.M
								v280 = m.ExcPending
								if v280 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v256)+40)) = v279
									v395 = F_planstate_tree_walker_impl(m, l0, int32(626), l1)
									mBase = m.M
									v396 = m.ExcPending
									if v396 != 0 {
										return int32(0)
									} else {
										return v395
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 float64
	_ = v89
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 float64
	_ = v106
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v122 float64
	_ = v122
	var v123 float64
	_ = v123
	var v126 float64
	_ = v126
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v135 float64
	_ = v135
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v169 int64
	_ = v169
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v177 int64
	_ = v177
	var v178 int64
	_ = v178
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v205 int32
	_ = v205
	var v208 int64
	_ = v208
	var v209 int64
	_ = v209
	var v212 int64
	_ = v212
	var v213 int64
	_ = v213
	var v216 int64
	_ = v216
	var v217 int64
	_ = v217
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
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
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
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
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
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
	v62 = l1 + v56 + v58*v21*int32(416)
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
	F_errfinish(m, int32(_a_F_ExecParallelRetrieveInstrumentation_1), int32(1050), int32(_a_F_ExecParallelRetrieveInstrumentation_2))
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
	v231 = v58
	goto L15
L15:
	;
	v235 = int32(_a_F_ExecParallelRetrieveInstrumentation_3)
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0])) = v239
	v242 = F_mul_size(m, v231, int32(416))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L9
	} else {
		goto L33
	}
L16:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v76 = v62 + v68*int32(416)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+4)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
	if v80 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v231 = v226
	goto L15
L18:
	;
	v225 = v68 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v225 < v226 {
		v68 = v225
		goto L16
	} else {
		goto L32
	}
L19:
	;
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v73)+16))
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+16)) = v102 + v103
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v76)+32))
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v73)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+32)) = base.F64_add(v106, v107)
	v110 = *(*float64)(unsafe.Add(mBase, uint32(v76)+200))
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v73)+200))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+200)) = base.F64_add(v110, v111)
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v76)+208))
	v115 = *(*float64)(unsafe.Add(mBase, uint32(v73)+208))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+208)) = base.F64_add(v114, v115)
	v118 = *(*float64)(unsafe.Add(mBase, uint32(v76)+216))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(v73)+216))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+216)) = base.F64_add(v118, v119)
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v76)+224))
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v73)+224))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+224)) = base.F64_add(v122, v123)
	v126 = *(*float64)(unsafe.Add(mBase, uint32(v76)+232))
	v127 = *(*float64)(unsafe.Add(mBase, uint32(v73)+232))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+232)) = base.F64_add(v126, v127)
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v76)+240))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v73)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+240)) = base.F64_add(v130, v131)
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v76)+248))
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v73)+248))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+248)) = base.F64_add(v134, v135)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v138 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	if v79&int32(1) == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v79&int32(1) == int32(0) {
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)) = uint8(v87)
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v76)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v73)+24)) = v89
	goto L19
L24:
	;
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v76)+24))
	v96 = *(*float64)(unsafe.Add(mBase, uint32(v73)+24))
	if base.F64_lt(v95, v96) == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v73)+24)) = v95
	goto L19
L26:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v73)+256))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v76)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+256)) = v141 + v142
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v73)+264))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v76)+264))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+264)) = v145 + v146
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v73)+272))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v76)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+272)) = v149 + v150
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v73)+280))
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v76)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+280)) = v153 + v154
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v73)+288))
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v76)+288))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+288)) = v157 + v158
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v73)+296))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v76)+296))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+296)) = v161 + v162
	v165 = *(*int64)(unsafe.Add(mBase, uint32(v73)+304))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v76)+304))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+304)) = v165 + v166
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v73)+312))
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v76)+312))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+312)) = v169 + v170
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v73)+320))
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v76)+320))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+320)) = v173 + v174
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v73)+328))
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v76)+328))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+328)) = v177 + v178
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v73)+336))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v76)+336))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+336)) = v181 + v182
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v73)+344))
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v76)+344))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+344)) = v185 + v186
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v73)+352))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v76)+352))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+352)) = v189 + v190
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v73)+360))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v76)+360))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+360)) = v193 + v194
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v73)+368))
	v198 = *(*int64)(unsafe.Add(mBase, uint32(v76)+368))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+368)) = v197 + v198
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v73)+376))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v76)+376))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+376)) = v201 + v202
	goto L28
L27:
	;
	goto L28
L28:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)))
	if v205 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v208 = *(*int64)(unsafe.Add(mBase, uint32(v73)+400))
	v209 = *(*int64)(unsafe.Add(mBase, uint32(v76)+400))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+400)) = v208 + v209
	v212 = *(*int64)(unsafe.Add(mBase, uint32(v73)+384))
	v213 = *(*int64)(unsafe.Add(mBase, uint32(v76)+384))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+384)) = v212 + v213
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v73)+392))
	v217 = *(*int64)(unsafe.Add(mBase, uint32(v76)+392))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+392)) = v216 + v217
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v73)+408))
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v76)+408))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+408)) = v220 + v221
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L18
L32:
	;
	goto L17
L33:
	;
	v246 = F_palloc(m, v242+int32(8))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v246
	*(*int32)(unsafe.Add(mBase, _c_F_ExecParallelRetrieveInstrumentation[0])) = v236
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v251
	if v242 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	base.MemoryCopy(m, v253+int32(8), v62, v242)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v257 - int32(405) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	case 3:
		goto L39
	default:
		goto L38
	case 20:
		goto L40
	case 21:
		goto L44
	case 22:
		goto L43
	case 24:
		goto L41
	case 29:
		goto L42
	}
L38:
	;
	v368 = F_planstate_tree_walker_impl(m, l0, int32(629), l1)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L9
	} else {
		goto L90
	}
L39:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v348 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L40:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	if v335 != 0 {
		goto L79
	} else {
		goto L80
	}
L41:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	if v322 != 0 {
		goto L72
	} else {
		goto L73
	}
L42:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v306 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L43:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	if v293 != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v280 != 0 {
		goto L54
	} else {
		goto L55
	}
L45:
	;
	F_ExecBitmapIndexScanRetrieveInstrumentation(m, l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L9
	} else {
		goto L53
	}
L46:
	;
	F_ExecBitmapIndexScanRetrieveInstrumentation(m, l0)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L9
	} else {
		goto L52
	}
L47:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	if v260 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	goto L38
L49:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v267 = v263<<(uint(int32(3))%32) + int32(8)
	v268 = F_palloc(m, v267)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v268
	if v267 == int32(0) {
		goto L48
	} else {
		goto L51
	}
L51:
	;
	base.MemoryCopy(m, v268, v260, v267)
	goto L48
L52:
	;
	goto L38
L53:
	;
	goto L38
L54:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v285 = v281<<(uint(int32(4))%32) | int32(8)
	v286 = F_palloc(m, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L9
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L38
L57:
	;
	if v285 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	base.MemoryCopy(m, v286, v288, v285)
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v286
	goto L56
L61:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	v298 = v294*int32(96) | int32(8)
	v299 = F_palloc(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L38
L64:
	;
	if v298 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+284))
	base.MemoryCopy(m, v299, v301, v298)
	goto L67
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+284)) = v299
	goto L63
L68:
	;
	goto L38
L69:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v313 = v309*int32(20) + int32(4)
	v314 = F_palloc(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v314
	if v313 == int32(0) {
		goto L68
	} else {
		goto L71
	}
L71:
	;
	base.MemoryCopy(m, v314, v306, v313)
	goto L68
L72:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)))
	v327 = v323*int32(24) + int32(8)
	v328 = F_palloc(m, v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L9
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L38
L75:
	;
	if v327 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+352))
	base.MemoryCopy(m, v328, v330, v327)
	goto L78
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+352)) = v328
	goto L74
L79:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v340 = v336*int32(40) + int32(8)
	v341 = F_palloc(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L9
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L38
L82:
	;
	if v340 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v341, v343, v340)
	goto L85
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v341
	goto L81
L86:
	;
	goto L38
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v355 = v351<<(uint(int32(4))%32) | int32(8)
	v356 = F_palloc(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v356
	if v355 == int32(0) {
		goto L86
	} else {
		goto L89
	}
L89:
	;
	base.MemoryCopy(m, v356, v348, v355)
	goto L86
L90:
	;
	m.G0 = v10 + int32(16)
	return v368
}
func F_ParallelApplyWorkerMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v232 int64
	_ = v232
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v554 int32
	_ = v554
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[0])) = uint8(v13)
	v16 = int32(914)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	switch int32(916) {
	case 0, 2:
		v30 = v16
		goto L2
	default:
		goto L3
	}
L1:
	;
	v62 = int32(295)
	v64 = m.G0
	v66 = v64 - int32(32)
	m.G0 = v66
	switch int32(297) {
	case 0, 2:
		v76 = v62
		goto L15
	default:
		goto L16
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v30
	F_sigemptyset(m, v20+int32(16))
	mBase = m.M
	goto L5
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[1])) = v16
	v30 = int32(_a_F_ParallelApplyWorkerMain_0)
	goto L2
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = int32(268435456)
	v42 = v20 + int32(12)
	goto L9
L7:
	;
	m.G0 = v20 + int32(32)
	goto L1
L9:
	;
	goto L10
L10:
	;
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v48 = int32(20)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[2])) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[3])) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[4])) = v54
	goto L13
L12:
	;
	goto L13
L13:
	;
	goto L7
L14:
	;
	v108 = int32(916)
	v110 = m.G0
	v112 = v110 - int32(32)
	m.G0 = v112
	switch int32(918) {
	case 0, 2:
		v122 = v108
		goto L28
	default:
		goto L29
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+12)) = v76
	F_sigemptyset(m, v66+int32(16))
	mBase = m.M
	goto L18
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[5])) = v62
	v76 = int32(_a_F_ParallelApplyWorkerMain_0)
	goto L15
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+24)) = int32(268435456)
	v88 = v66 + int32(12)
	goto L22
L20:
	;
	m.G0 = v66 + int32(32)
	goto L14
L22:
	;
	goto L23
L23:
	;
	if v88 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v95 = int32(300)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[6])) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v88)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[7])) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v88)))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[8])) = v100
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L20
L27:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v122
	F_sigemptyset(m, v112+int32(16))
	mBase = m.M
	goto L31
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[9])) = v108
	v122 = int32(_a_F_ParallelApplyWorkerMain_0)
	goto L28
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = int32(268435456)
	v134 = v112 + int32(12)
	goto L35
L33:
	;
	m.G0 = v112 + int32(32)
	goto L27
L35:
	;
	goto L36
L36:
	;
	if v134 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = int32(240)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[10])) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v134)+8))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[11])) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v134)))
	*(*int64)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[12])) = v146
	goto L39
L38:
	;
	goto L39
L39:
	;
	goto L33
L40:
	;
	return
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[13]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+1328))
	v158 = F_dsm_attach(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L40
	} else {
		goto L46
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L40
	} else {
		goto L146
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L40
	} else {
		goto L143
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L40
	} else {
		goto L140
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L40
	} else {
		goto L136
	}
L46:
	;
	if v158 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+24))
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v161)))
	if v163 == int64(2021433447) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L40
	} else {
		goto L132
	}
L50:
	;
	if v165 == int32(0) {
		goto L45
	} else {
		goto L54
	}
L51:
	;
	v165 = v161
	goto L53
L52:
	;
	v165 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v171 = F_shm_toc_lookup(m, v165, int64(1), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L40
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14])) = v171
	v176 = F_shm_toc_lookup(m, v165, int64(2), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L40
	} else {
		goto L56
	}
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[15]))
	F_shm_mq_set_receiver(m, v176, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L40
	} else {
		goto L57
	}
L57:
	;
	v182 = F_shm_mq_attach(m, v176, v158)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L40
	} else {
		goto L58
	}
L58:
	;
	F_logicalrep_worker_attach(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L40
	} else {
		goto L59
	}
L59:
	;
	F_before_shmem_exit(m, int32(985), v158)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L40
	} else {
		goto L60
	}
L60:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = int32(1)
	if v191 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	F_s_lock(m, v195, int32(_a_F_ParallelApplyWorkerMain_1), int32(930), int32(_a_F_ParallelApplyWorkerMain_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L40
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+18)))
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+16)) = l0
	*(*uint16)(unsafe.Add(mBase, uint32(v205)+12)) = uint16(v203)
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v208
	v212 = F_shm_toc_lookup(m, v165, int64(3), v208)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L40
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[15]))
	F_shm_mq_set_sender(m, v212, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L40
	} else {
		goto L66
	}
L66:
	;
	v218 = F_shm_mq_attach(m, v212, v158)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L40
	} else {
		goto L67
	}
L67:
	;
	F_pq_redirect_to_shm_mq(m, v158, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L40
	} else {
		goto L68
	}
L68:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[17])) = int32(-1)
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[18])) = v224
	goto L69
L69:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v232 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v231)+80)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v231)+104)) = v232
	*(*int64)(unsafe.Add(mBase, uint32(v231)+88)) = v232
	F_InitializeLogRepWorker(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L40
	} else {
		goto L70
	}
L70:
	;
	v241 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[0])) = uint8(v241)
	F_StartTransactionCommand(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L40
	} else {
		goto L71
	}
L71:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[19]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v250 = v10 + int32(32)
	F_ReplicationOriginNameForLogicalRep(m, v247, int32(0), v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L40
	} else {
		goto L72
	}
L72:
	;
	v254 = F_replorigin_by_name(m, v250, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L40
	} else {
		goto L73
	}
L73:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+64))
	F_replorigin_session_setup(m, v254, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L40
	} else {
		goto L74
	}
L74:
	;
	*(*uint16)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[20])) = uint16(v254)
	F_CommitTransactionCommand(m)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L40
	} else {
		goto L75
	}
L75:
	;
	F_CacheRegisterSyscacheCallback(m, int32(68), int32(986), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L40
	} else {
		goto L76
	}
L76:
	;
	F_set_apply_error_context_origin(m, v250)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L40
	} else {
		goto L77
	}
L77:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[21]))
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[22]))
	v281 = F_AllocSetContextCreateInternal(m, v276, int32(_a_F_ParallelApplyWorkerMain_3), int32(0), int32(_a_F_ParallelApplyWorkerMain_4), int32(_a_F_ParallelApplyWorkerMain_5))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L40
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[23])) = v281
	*(*int32)(unsafe.Add(mBase, uint32(v10)+136)) = int32(987)
	v286 = int32(_a_F_ParallelApplyWorkerMain_6)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[24])) = v10 + int32(132)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+132)) = v287
	goto L79
L79:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[25]))
	if v301 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L40
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[26]))
	if v305 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	v308 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L40
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[27]))
	if v328 != 0 {
		goto L95
	} else {
		goto L96
	}
L88:
	;
	if v308 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[19]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v312
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_7), v10+int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L40
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L40
	} else {
		goto L94
	}
L92:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(720), int32(_a_F_ParallelApplyWorkerMain_8))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L40
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[27])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L40
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[23]))
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[21])) = v337
	v344 = F_shm_mq_receive(m, v182, v10+int32(124), v10+int32(128), int32(1))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L40
	} else {
		goto L102
	}
L98:
	;
	goto L97
L99:
	;
	v476 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[23]))
	F_MemoryContextReset(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L40
	} else {
		goto L131
	}
L100:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	*(*int32)(unsafe.Add(mBase, uint32(v367))) = int32(1)
	if v368 != 0 {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v10)+124))
	if v346 == int32(0) {
		goto L44
	} else {
		goto L103
	}
L102:
	;
	switch v344 {
	case 0:
		goto L101
	case 1:
		goto L100
	default:
		goto L42
	}
L103:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10)+116)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = v346
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v10)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+108)) = v352
	v355 = v10 + int32(108)
	v356 = F_pq_getmsgbyte(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L40
	} else {
		goto L104
	}
L104:
	;
	if v356 != int32(119) {
		goto L43
	} else {
		goto L105
	}
L105:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v10)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+120)) = v360 + int32(24)
	F_apply_dispatch(m, v355)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L40
	} else {
		goto L106
	}
L106:
	;
	goto L99
L107:
	;
	v372 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	F_s_lock(m, v372, int32(_a_F_ParallelApplyWorkerMain_1), int32(1531), int32(_a_F_ParallelApplyWorkerMain_9))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L40
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+32))
	switch v382 {
	case 0:
		goto L111
	case 1:
		goto L113
	default:
		v421 = v382
		goto L112
	}
L110:
	;
	goto L109
L111:
	;
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[28]))
	v463 = F_WaitLatch(m, v459, int32(41), int32(1000), int32(83886089))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L40
	} else {
		goto L128
	}
L112:
	;
	switch v421 - int32(2) {
	case 0:
		goto L123
	case 1:
		goto L122
	default:
		goto L99
	}
L113:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v384)+32))
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	F_LockApplyTransactionForSession(m, v385, v388, int32(0), int32(1))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L40
	} else {
		goto L114
	}
L114:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[16]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)+32))
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	F_UnlockApplyTransactionForSession(m, v395, v398, int32(0), int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L40
	} else {
		goto L115
	}
L115:
	;
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)))
	*(*int32)(unsafe.Add(mBase, uint32(v404))) = int32(1)
	if v405 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	F_s_lock(m, v409, int32(_a_F_ParallelApplyWorkerMain_1), int32(1531), int32(_a_F_ParallelApplyWorkerMain_9))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L40
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v416 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	*(*int32)(unsafe.Add(mBase, uint32(v416))) = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+32))
	v421 = v419
	goto L112
L119:
	;
	goto L118
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v454))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v454)+32)) = v453
	goto L99
L121:
	;
	F_s_lock(m, v447, int32(_a_F_ParallelApplyWorkerMain_1), int32(1508), int32(_a_F_ParallelApplyWorkerMain_10))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L40
	} else {
		goto L127
	}
L122:
	;
	v430 = int32(0)
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v432)+4))
	F_apply_spooled_messages(m, v432+int32(36), v435, int64(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L40
	} else {
		goto L125
	}
L123:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = int32(1)
	v429 = int32(3)
	if v426 != 0 {
		v446 = v429
		v447 = v425
		goto L121
	} else {
		goto L124
	}
L124:
	;
	v453 = v429
	v454 = v425
	goto L120
L125:
	;
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[14]))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v440)))
	*(*int32)(unsafe.Add(mBase, uint32(v440))) = int32(1)
	if v441 == int32(0) {
		v453 = v430
		v454 = v440
		goto L120
	} else {
		goto L126
	}
L126:
	;
	v446 = v430
	v447 = v440
	goto L121
L127:
	;
	v453 = v446
	v454 = v447
	goto L120
L128:
	;
	if v463&int32(1) == int32(0) {
		goto L99
	} else {
		goto L129
	}
L129:
	;
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = int32(0)
	goto L130
L130:
	;
	goto L99
L131:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ParallelApplyWorkerMain[21])) = v273
	goto L79
L132:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L40
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_11), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L40
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(897), int32(_a_F_ParallelApplyWorkerMain_2))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L40
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L40
	} else {
		goto L137
	}
L137:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_12), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L40
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(903), int32(_a_F_ParallelApplyWorkerMain_2))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L40
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errmsg_internal(m, int32(_a_F_ParallelApplyWorkerMain_13), int32(0))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L40
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(774), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L40
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v356
	F_errmsg_internal(m, int32(_a_F_ParallelApplyWorkerMain_15), v10)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L40
	} else {
		goto L144
	}
L144:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(784), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L40
	} else {
		goto L145
	}
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L146:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L40
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_ParallelApplyWorkerMain_16), int32(0))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L40
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_ParallelApplyWorkerMain_1), int32(822), int32(_a_F_ParallelApplyWorkerMain_14))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L40
	} else {
		goto L149
	}
L149:
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
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+84)))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+64))
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
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
