package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v14 != v16 {
		v82 = int32(0)
		m.G0 = v10 + int32(32)
		return v82
	} else {
		v18 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v18)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v20
		if l1 != 0 {
			v22 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v22
			v34 = v22
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v25 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v25
				v34 = int32(-1)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v28
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
				v34 = v30 - v28
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v34
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
		v40 = int32(0)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
		v49 = F_populate_record_field(m, v37, v38, v39, v40, v41, int64(0), v10+int32(12), v10+int32(31), v47, v40)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
			if v53 == int32(0) {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v66 = F_accumArrayResult(m, v61, v49, v62, v64, v65)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
					v73 = v68 + v14<<(uint(int32(2))%32) - int32(4)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
					*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 + int32(1)
					v82 = int32(0)
					m.G0 = v10 + int32(32)
					return v82
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
				if v56 != int32(453) {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					v66 = F_accumArrayResult(m, v61, v49, v62, v64, v65)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
						v73 = v68 + v14<<(uint(int32(2))%32) - int32(4)
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
						*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 + int32(1)
						v82 = int32(0)
						m.G0 = v10 + int32(32)
						return v82
					}
				} else {
					v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+4)))
					if v60 != 0 {
						v82 = int32(23)
						m.G0 = v10 + int32(32)
						return v82
					} else {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+31)))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						v66 = F_accumArrayResult(m, v61, v49, v62, v64, v65)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
							v73 = v68 + v14<<(uint(int32(2))%32) - int32(4)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
							*(*int32)(unsafe.Add(mBase, uint32(v73))) = v74 + int32(1)
							v82 = int32(0)
							m.G0 = v10 + int32(32)
							return v82
						}
					}
				}
			}
		}
	}
}
func F_populate_array_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v11 <= v2 {
		if v9 <= int32(0) {
			F_populate_array_report_expected_array(m, v10, v9)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v104 = int32(23)
				return v104
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v9
			v18 = F_palloc_mul(m, int32(4), v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v18
				v24 = F_palloc0_mul(m, int32(4), v9)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v24
					v28 = v9 & int32(3)
					v29 = int32(0)
					if base.Ui32(int32(4)) <= base.Ui32(v9) {
						v34 = v29
						v40 = v2
						for {
							v42 = v34 << (uint(int32(2)) % 32)
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							v45 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v42+v43))) = v45
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v47+v42)+4)) = v45
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v51+v42)+8)) = v45
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v55+v42)+12)) = v45
							v59 = int32(4)
							v60 = v34 + v59
							v62 = v40 + v59
							if v62 != v9&int32(2147483644) {
								v34 = v60
								v40 = v62
								continue
							} else {
								break
							}
							break
						}
						if v28 == int32(0) {
						} else {
							v66 = v60
							v73 = v66
							v76 = v2
							for {
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32)))) = int32(-1)
								v86 = int32(1)
								v89 = v76 + v86
								if v89 != v28 {
									v73 = v73 + v86
									v76 = v89
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v66 = v29
						v73 = v66
						v76 = v2
						for {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32)))) = int32(-1)
							v86 = int32(1)
							v89 = v76 + v86
							if v89 != v28 {
								v73 = v73 + v86
								v76 = v89
								continue
							} else {
								break
							}
							break
						}
					}
					return int32(0)
				}
			}
		}
	} else {
		if v11 <= v9 {
			v104 = v2
			return v104
		} else {
			F_populate_array_report_expected_array(m, v10, v9)
			mBase = m.M
			v102 = m.ExcPending
			if v102 != 0 {
				return int32(0)
			} else {
				v104 = int32(23)
				return v104
			}
		}
	}
}
func F_populate_record_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int64, l6 int32, l7 int32, l8 int32, l9 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int64
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v370 int64
	_ = v370
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int64
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int64
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int64
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int64
	_ = v432
	v16 = m.G0
	v18 = v16 - int32(128)
	m.G0 = v18
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == v24 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v31 = int32(1)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v33 == v31 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 == l2 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_prepare_column_cache(m, l0, l1, l2, l4, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	goto L3
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v46)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v49 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	if v32 == int32(0) {
		v46 = v31
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v32 == int32(0) {
		v46 = v31
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v46 = base.B2i32(v38 == int32(11))
	goto L9
L14:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v46 = base.B2i32(v43 == int32(0))
	goto L9
L15:
	;
	if v46 != 0 {
		goto L35
	} else {
		goto L36
	}
L16:
	;
	v62 = int32(115)
	if v48&int32(-3) == int32(97) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v52 == int32(1) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v55 == int32(0) {
		v72 = v48
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v72 = v48
	goto L15
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v58 != int32(1) {
		v72 = v48
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	v68 = v62
	goto L25
L24:
	;
	v68 = v48
	goto L25
L25:
	;
	if v48 == int32(67) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v71 = v62
	goto L28
L27:
	;
	v71 = v68
	goto L28
L28:
	;
	v72 = v71
	goto L15
L29:
	;
	m.G0 = v18 + int32(128)
	return v432
L30:
	;
	v424 = F_JsonbValueToJsonb(m, v112)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L141
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L138
	}
L32:
	;
	v406 = F_domain_check_safe(m, v400, v401&int32(1), l1, l0+int32(56), l4, l8)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L136
	}
L33:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v397 = F_populate_record_field(m, v393, v394, v395, l3, l4, int64(0), l6, l7, l8, l9)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L135
	}
L34:
	;
	v382 = l0 + int32(44)
	v383 = base.I32_wrap_i64(l5)
	if v383 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L35:
	;
	if v72 == int32(67) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	switch v72 - int32(67) {
	case 0, 32:
		goto L34
	default:
		goto L31
	case 30:
		goto L40
	case 33:
		goto L33
	case 48:
		goto L41
	}
L38:
	;
	v77 = int64(0)
	if v72 == int32(100) {
		v400 = v77
		v401 = int32(1)
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v432 = v77
	goto L29
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l0 + int32(44)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_populate_record_field[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v216
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v221 = F_initArrayResult(m, v219, v216, v218)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L96
	}
L41:
	;
	if v49 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v199 = F_InputFunctionCallSafe(m, l0+int32(16), v192, v196, l2, l8, v18+int32(32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L88
	}
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v192 = v188
	v193 = v82
	goto L42
L44:
	;
	F_escape_json(m, v18+int32(68), v82)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L87
	}
L45:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if base.B2i32(l1 != int32(3802))&base.B2i32(l1 != int32(114)) != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v110 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if base.B2i32(l9 == v110)|base.B2i32(v113 != int32(1)) == v110 {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	if v83 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v89 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v93 = v18 + int32(68)
	F_initStringInfo(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v83 < int32(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_escape_json_with_len(m, v93, v82, v83)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L43
L54:
	;
	v192 = v82
	v193 = v82
	goto L42
L55:
	;
	goto L56
L56:
	;
	v104 = F_palloc(m, v83+int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v83 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	base.MemoryCopy(m, v104, v82, v83)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v108 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v104))) = uint8(v108)
	v192 = v104
	v193 = v82
	goto L42
L61:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v122 = F_pnstrdup(m, v120, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if l1 == int32(3802) {
		goto L30
	} else {
		goto L65
	}
L64:
	;
	v192 = v122
	v193 = int32(0)
	goto L42
L65:
	;
	if l1 == int32(114) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L84
	}
L67:
	;
	v162 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v166 = F_JsonbToCString(m, v162, v164, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L83
	}
L68:
	;
	if v113 == int32(18) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	switch v113 - int32(1) {
	case 0:
		goto L76
	case 1:
		goto L74
	case 2:
		goto L75
	default:
		goto L66
	case 17:
		goto L67
	}
L71:
	;
	v130 = int32(0)
	v132 = F_JsonbValueToJsonb(m, v112)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v139 = F_JsonbToCString(m, v130, v132+int32(4), int32(base.Ui32(v136)>>(uint(int32(2))%32)))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v192 = v139
	v193 = v130
	goto L42
L74:
	;
	v155 = int32(0)
	v158 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v112)+8)))
	v159 = F_DirectFunctionCall1Coll(m, int32(664), v155, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L82
	}
L75:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+8)))
	if v151 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v146 = F_pnstrdup(m, v144, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v192 = v146
	v193 = int32(0)
	goto L42
L78:
	;
	v152 = int32(_a_F_populate_record_field_3)
	goto L80
L79:
	;
	v152 = int32(_a_F_populate_record_field_4)
	goto L80
L80:
	;
	v153 = F_pstrdup(m, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v192 = v153
	v193 = int32(0)
	goto L42
L82:
	;
	v192 = base.I32_wrap_i64(v159)
	v193 = v155
	goto L42
L83:
	;
	v192 = v166
	v193 = v162
	goto L42
L84:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v172
	F_errmsg_internal(m, int32(_a_F_populate_record_field_5), v18+int32(16))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_populate_record_field_1), int32(3199), int32(_a_F_populate_record_field_6))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	goto L43
L88:
	;
	if v199 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = int64(0)
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v205)
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v192 != v193 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_pfree(m, v192)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	v432 = v210
	goto L29
L95:
	;
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = l8
	v228 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+52)) = v228
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v232 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v376)
	v432 = v370
	goto L29
L98:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v305 = F_palloc_mul(m, int32(4), v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L118
	}
L99:
	;
	v235 = int32(0)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v236 < v235 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v293 = F_populate_array_dim_jsonb(m, v18+int32(32), v231, int32(1))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L116
	}
L102:
	;
	v239 = F_strlen(m, v231)
	mBase = m.M
	v240 = v239
	goto L104
L103:
	;
	v240 = v236
	goto L104
L104:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_populate_record_field[1]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	goto L105
L105:
	;
	v245 = F_makeJsonLexContextCstringLen(m, v235, v231, v240, v243, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v247 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+84)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v245
	*(*int64)(unsafe.Add(mBase, uint32(v18)+76)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = int32(1499)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(1500)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = int32(1501)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(1502)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(1503)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v18 + int32(108)
	v272 = F_pg_parse_json(m, v245, v18+int32(68))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v272 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_json_errsave_error(m, v272, v245, l8)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_freeJsonLexContext(m, v276)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	if v279 == int32(0) {
		goto L98
	} else {
		goto L113
	}
L113:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	if v282 != int32(453) {
		goto L98
	} else {
		goto L114
	}
L114:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+4)))
	if v285 == int32(0) {
		goto L98
	} else {
		goto L115
	}
L115:
	;
	v288 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v288)
	v432 = v228
	goto L29
L116:
	;
	if v293 == int32(0) {
		v370 = v228
		v376 = v218
		goto L97
	} else {
		goto L117
	}
L117:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)))
	*(*int32)(unsafe.Add(mBase, uint32(v297))) = v299
	goto L98
L118:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(0) < v307 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v317 = int32(0)
	goto L122
L120:
	;
	v346 = v307
	goto L121
L121:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v354 = F_makeMdArrayResult(m, v350, v346, v351, v305, v352, int32(1))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L125
	}
L122:
	;
	v329 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v305+v317<<(uint(int32(2))%32)))) = v329
	v332 = v317 + v329
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v332 < v333 {
		v317 = v332
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v346 = v333
	goto L121
L124:
	;
	goto L123
L125:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	F_pfree(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	F_pfree(m, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_pfree(m, v305)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v370 = v354
	v376 = int32(0)
	goto L97
L129:
	;
	v387 = F_populate_composite(m, v382, l1, l4, int32(0), l6, l7, l8)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v389 = F_pg_detoast_datum(m, v383)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	v432 = v387
	goto L29
L133:
	;
	v391 = F_populate_composite(m, v382, l1, l4, v389, l6, l7, l8)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v432 = v391
	goto L29
L135:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7))))
	v400 = v397
	v401 = v399
	goto L32
L136:
	;
	if v406 != 0 {
		v432 = v400
		goto L29
	} else {
		goto L137
	}
L137:
	;
	v408 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v408)
	v432 = int64(0)
	goto L29
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v72
	F_errmsg_internal(m, int32(_a_F_populate_record_field_0), v18)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_populate_record_field_1), int32(3470), int32(_a_F_populate_record_field_2))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	v432 = base.I64_extend_i32_u(v424)
	goto L29
}
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int64
	_ = v139
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	v3 = l2
	v6 = int32(0)
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v10
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v21 == v6 {
		v25 = F_MemoryContextAllocZero(m, v20, int32(72))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int64(0)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v20
			if l3 != 0 {
				F_get_record_type_from_argument(m, l0, l1, v25)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					v38 = v25
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v39 == int32(0) {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v43 = F_pg_detoast_datum(m, v42)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
							if v45 != int32(2249) {
								v57 = v38
								v58 = v43
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v48
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v50
								v57 = v38
								v58 = v43
							}
							v61 = l0 + l3<<(uint(int32(4))%32)
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
							if v62 == int32(1) {
								if v58 != 0 {
									v144 = base.I64_extend_i32_u(v58)
								} else {
									v66 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
									v144 = int64(0)
								}
								m.G0 = v13 - int32(-64)
								return v144
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
								if v3 != 0 {
									v71 = F_pg_detoast_datum_packed(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										v73 = int32(1)
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
										v77 = v75 & v73
										if v77 != 0 {
											v78 = v73
										} else {
											v78 = int32(4)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
										if v75 == int32(1) {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
											if v86 == int32(18) {
												v89 = int32(16)
											} else {
												v89 = int32(0)
											}
											if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v96 = int32(4)
											} else {
												v96 = v89
											}
											v107 = v96
										} else {
											v97 = int32(1)
											if v77 != 0 {
												v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
												v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int64(0)
										} else {
											v144 = v139
											m.G0 = v13 - int32(-64)
											return v144
										}
									}
								} else {
									v111 = F_pg_detoast_datum(m, v70)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
										v115 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int64(0)
										} else {
											v144 = v139
											m.G0 = v13 - int32(-64)
											return v144
										}
									}
								}
							}
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
						if v52 != int32(2249) {
							v57 = v38
							v58 = v6
							v61 = l0 + l3<<(uint(int32(4))%32)
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
							if v62 == int32(1) {
								if v58 != 0 {
									v144 = base.I64_extend_i32_u(v58)
								} else {
									v66 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
									v144 = int64(0)
								}
								m.G0 = v13 - int32(-64)
								return v144
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
								if v3 != 0 {
									v71 = F_pg_detoast_datum_packed(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return int64(0)
									} else {
										v73 = int32(1)
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
										v77 = v75 & v73
										if v77 != 0 {
											v78 = v73
										} else {
											v78 = int32(4)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
										if v75 == int32(1) {
											v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
											if v86 == int32(18) {
												v89 = int32(16)
											} else {
												v89 = int32(0)
											}
											if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v96 = int32(4)
											} else {
												v96 = v89
											}
											v107 = v96
										} else {
											v97 = int32(1)
											if v77 != 0 {
												v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
												v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int64(0)
										} else {
											v144 = v139
											m.G0 = v13 - int32(-64)
											return v144
										}
									}
								} else {
									v111 = F_pg_detoast_datum(m, v70)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
										v115 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
										*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
										v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int64(0)
										} else {
											v144 = v139
											m.G0 = v13 - int32(-64)
											return v144
										}
									}
								}
							}
						} else {
							F_get_record_type_from_query(m, l0, l1, v38)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int64(0)
							} else {
								v57 = v38
								v58 = v6
								v61 = l0 + l3<<(uint(int32(4))%32)
								v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
								if v62 == int32(1) {
									if v58 != 0 {
										v144 = base.I64_extend_i32_u(v58)
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
										v144 = int64(0)
									}
									m.G0 = v13 - int32(-64)
									return v144
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
									if v3 != 0 {
										v71 = F_pg_detoast_datum_packed(m, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int64(0)
										} else {
											v73 = int32(1)
											v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
											v77 = v75 & v73
											if v77 != 0 {
												v78 = v73
											} else {
												v78 = int32(4)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
											if v75 == int32(1) {
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
												if v86 == int32(18) {
													v89 = int32(16)
												} else {
													v89 = int32(0)
												}
												if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
													v96 = int32(4)
												} else {
													v96 = v89
												}
												v107 = v96
											} else {
												v97 = int32(1)
												if v77 != 0 {
													v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
													v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
											v130 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int64(0)
											} else {
												v144 = v139
												m.G0 = v13 - int32(-64)
												return v144
											}
										}
									} else {
										v111 = F_pg_detoast_datum(m, v70)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
											v115 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
											*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
											v130 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
											v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int64(0)
											} else {
												v144 = v139
												m.G0 = v13 - int32(-64)
												return v144
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_get_record_type_from_query(m, l0, l1, v25)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					v57 = v25
					v58 = v6
					v61 = l0 + l3<<(uint(int32(4))%32)
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
					if v62 == int32(1) {
						if v58 != 0 {
							v144 = base.I64_extend_i32_u(v58)
						} else {
							v66 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
							v144 = int64(0)
						}
						m.G0 = v13 - int32(-64)
						return v144
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
						if v3 != 0 {
							v71 = F_pg_detoast_datum_packed(m, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								v73 = int32(1)
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
								v77 = v75 & v73
								if v77 != 0 {
									v78 = v73
								} else {
									v78 = int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
								if v75 == int32(1) {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
									if v86 == int32(18) {
										v89 = int32(16)
									} else {
										v89 = int32(0)
									}
									if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v96 = int32(4)
									} else {
										v96 = v89
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v70)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						}
					}
				}
			}
		}
	} else {
		if l3 == int32(0) {
			v57 = v21
			v58 = v6
			v61 = l0 + l3<<(uint(int32(4))%32)
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
			if v62 == int32(1) {
				if v58 != 0 {
					v144 = base.I64_extend_i32_u(v58)
				} else {
					v66 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
					v144 = int64(0)
				}
				m.G0 = v13 - int32(-64)
				return v144
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
				if v3 != 0 {
					v71 = F_pg_detoast_datum_packed(m, v70)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int64(0)
					} else {
						v73 = int32(1)
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
						v77 = v75 & v73
						if v77 != 0 {
							v78 = v73
						} else {
							v78 = int32(4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
						if v75 == int32(1) {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
							if v86 == int32(18) {
								v89 = int32(16)
							} else {
								v89 = int32(0)
							}
							if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v96 = int32(4)
							} else {
								v96 = v89
							}
							v107 = v96
						} else {
							v97 = int32(1)
							if v77 != 0 {
								v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
								v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
						v130 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							v144 = v139
							m.G0 = v13 - int32(-64)
							return v144
						}
					}
				} else {
					v111 = F_pg_detoast_datum(m, v70)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
						v115 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
						*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
						v130 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							v144 = v139
							m.G0 = v13 - int32(-64)
							return v144
						}
					}
				}
			}
		} else {
			v38 = v21
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v39 == int32(0) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v43 = F_pg_detoast_datum(m, v42)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
					if v45 != int32(2249) {
						v57 = v38
						v58 = v43
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v48
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v50
						v57 = v38
						v58 = v43
					}
					v61 = l0 + l3<<(uint(int32(4))%32)
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
					if v62 == int32(1) {
						if v58 != 0 {
							v144 = base.I64_extend_i32_u(v58)
						} else {
							v66 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
							v144 = int64(0)
						}
						m.G0 = v13 - int32(-64)
						return v144
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
						if v3 != 0 {
							v71 = F_pg_detoast_datum_packed(m, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								v73 = int32(1)
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
								v77 = v75 & v73
								if v77 != 0 {
									v78 = v73
								} else {
									v78 = int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
								if v75 == int32(1) {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
									if v86 == int32(18) {
										v89 = int32(16)
									} else {
										v89 = int32(0)
									}
									if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v96 = int32(4)
									} else {
										v96 = v89
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v70)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						}
					}
				}
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
				if v52 != int32(2249) {
					v57 = v38
					v58 = v6
					v61 = l0 + l3<<(uint(int32(4))%32)
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
					if v62 == int32(1) {
						if v58 != 0 {
							v144 = base.I64_extend_i32_u(v58)
						} else {
							v66 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
							v144 = int64(0)
						}
						m.G0 = v13 - int32(-64)
						return v144
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
						if v3 != 0 {
							v71 = F_pg_detoast_datum_packed(m, v70)
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return int64(0)
							} else {
								v73 = int32(1)
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
								v77 = v75 & v73
								if v77 != 0 {
									v78 = v73
								} else {
									v78 = int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
								if v75 == int32(1) {
									v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
									if v86 == int32(18) {
										v89 = int32(16)
									} else {
										v89 = int32(0)
									}
									if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v96 = int32(4)
									} else {
										v96 = v89
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v70)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									v144 = v139
									m.G0 = v13 - int32(-64)
									return v144
								}
							}
						}
					}
				} else {
					F_get_record_type_from_query(m, l0, l1, v38)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int64(0)
					} else {
						v57 = v38
						v58 = v6
						v61 = l0 + l3<<(uint(int32(4))%32)
						v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+32)))
						if v62 == int32(1) {
							if v58 != 0 {
								v144 = base.I64_extend_i32_u(v58)
							} else {
								v66 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v66)
								v144 = int64(0)
							}
							m.G0 = v13 - int32(-64)
							return v144
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v13)+48)) = uint8(v3)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
							if v3 != 0 {
								v71 = F_pg_detoast_datum_packed(m, v70)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return int64(0)
								} else {
									v73 = int32(1)
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
									v77 = v75 & v73
									if v77 != 0 {
										v78 = v73
									} else {
										v78 = int32(4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v71 + v78
									if v75 == int32(1) {
										v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
										if v86 == int32(18) {
											v89 = int32(16)
										} else {
											v89 = int32(0)
										}
										if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v96 = int32(4)
										} else {
											v96 = v89
										}
										v107 = v96
									} else {
										v97 = int32(1)
										if v77 != 0 {
											v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
											v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v107
									v130 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										v144 = v139
										m.G0 = v13 - int32(-64)
										return v144
									}
								}
							} else {
								v111 = F_pg_detoast_datum(m, v70)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(18)
									v115 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v111 + v115
									*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v11 + int32(-56)
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
									v130 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v13)+47)) = uint8(v130)
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v139 = F_populate_composite(m, v57+int32(48), v134, v20, v58, v11+int32(-16), v11+int32(-17), l4)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										v144 = v139
										m.G0 = v13 - int32(-64)
										return v144
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
func F_populate_recordset_object_field_end(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14357(m, l0, l1, l2, int32(2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_populate_recordset_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v413 int32
	_ = v413
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 == v5 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	v413 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v413)
	m.G0 = v16 + int32(112)
	return
L2:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v394
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	v397 = F_CreateTupleDescCopy(m, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L113
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L17
	} else {
		goto L109
	}
L4:
	;
	v306 = v298
	v309 = v297
	goto L97
L5:
	;
	v297 = v5
	v298 = int32(1)
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L17
	} else {
		goto L93
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L17
	} else {
		goto L89
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L17
	} else {
		goto L85
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(389) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+12)))
	if v24&int32(2) == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(2)
	if v30 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v74 = l0 + l3<<(uint(int32(4))%32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+32)))
	if v75 != 0 {
		goto L1
	} else {
		goto L32
	}
L13:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v52 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v38 = F_MemoryContextAllocZero(m, v36, int32(72))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if l3 == int32(0) {
		v70 = v30
		v71 = v5
		goto L12
	} else {
		goto L24
	}
L17:
	;
	return
L18:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v38
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v43
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_get_record_type_from_argument(m, l0, l1, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_get_record_type_from_query(m, l0, l1, v38)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v51 = v38
	goto L13
L23:
	;
	v70 = v38
	v71 = v5
	goto L12
L24:
	;
	v51 = v30
	goto L13
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v56 = F_pg_detoast_datum(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L17
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v65 != int32(2249) {
		v70 = v51
		v71 = v5
		goto L12
	} else {
		goto L30
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v58 != int32(2249) {
		v70 = v51
		v71 = v56
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+56)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+60)) = v63
	v70 = v51
	v71 = v56
	goto L12
L30:
	;
	F_get_record_type_from_query(m, l0, l1, v51)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v70 = v51
	v71 = v5
	goto L12
L32:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	if v79 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v114 = F_palloc0(m, int32(36))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L48
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v70)+60))
	v92 = F_lookup_rowtype_tupdesc(m, v89, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L40
	}
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	v89 = v82
	goto L34
L36:
	;
	goto L37
L37:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	if v83 != v84 {
		v89 = v84
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v70)+60))
	if v86 == v87 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v89 = v83
	goto L34
L40:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	if v94 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_FreeTupleDesc(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v97 = int32(_a_F_populate_recordset_worker_0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0])) = v78
	v101 = F_CreateTupleDescCopy(m, v92)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L17
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+52)) = v101
	*(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0])) = v98
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	if v106 < int32(0) {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	F_DecrTupleDescRefCount(m, v92)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v116 = int32(_a_F_populate_recordset_worker_0)
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0])) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[1]))
	v130 = F_tuplestore_begin_heap(m, int32(base.Ui32(v122&int32(4))>>(uint(int32(2))%32)), int32(0), v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = v130
	*(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[0])) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v114)+32)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v114)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v71
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(24))))
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v139 = F_pg_detoast_datum_packed(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L17
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v217 = F_pg_detoast_datum(m, v138)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L17
	} else {
		goto L78
	}
L53:
	;
	v142 = F_palloc0(m, int32(40))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v144 = F_pg_detoast_datum_packed(m, v139)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	v177 = v16 + int32(40)
	v178 = int32(1)
	if v146&v178 != 0 {
		goto L67
	} else {
		goto L68
	}
L56:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v146 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	if v152 == int32(18) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v163 = int32(1)
	if v146&v163 != 0 {
		v175 = int32(base.Ui32(v146)>>(uint(v163)%32)) - v163
		goto L55
	} else {
		goto L66
	}
L60:
	;
	v155 = int32(16)
	goto L62
L61:
	;
	v155 = int32(0)
	goto L62
L62:
	;
	if base.Ui32((v152-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v162 = int32(4)
	goto L65
L64:
	;
	v162 = v155
	goto L65
L65:
	;
	v175 = v162
	goto L55
L66:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v175 = int32(base.Ui32(v169)>>(uint(int32(2))%32)) - int32(4)
	goto L55
L67:
	;
	v182 = v178
	goto L69
L68:
	;
	v182 = int32(4)
	goto L69
L69:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[2]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	goto L70
L70:
	;
	v188 = F_makeJsonLexContextCstringLen(m, v177, v144+v182, v175, v186, int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+36)) = int32(1504)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+28)) = int32(1505)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = int32(1506)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = int32(1507)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1508)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = int32(1509)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = int32(1510)
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v177
	v206 = F_pg_parse_json(m, v177, v142)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	if v206 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_json_errsave_error(m, v206, v177, int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L17
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_freeJsonLexContext(m, v16+int32(40))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L17
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(0)
	goto L2
L78:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	if v219&int32(1342177280) != int32(1073741824) {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v226 = F_JsonbIteratorInit(m, v217+int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v226
	v234 = F_JsonbIteratorNext(m, v16+int32(36), v16+int32(40), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L17
	} else {
		goto L82
	}
L81:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v236 != int32(18) {
		goto L3
	} else {
		goto L83
	}
L82:
	;
	switch v234 {
	case 0:
		goto L2
	default:
		goto L5
	case 3:
		goto L81
	}
L83:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+3)))
	if v240&int32(32) == int32(0) {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v297 = v239
	v298 = int32(0)
	goto L4
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_populate_recordset_worker_1), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(4056), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L17
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_populate_recordset_worker_4), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(4061), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L17
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
	F_errmsg(m, int32(_a_F_populate_recordset_worker_5), v16+int32(16))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(_a_F_populate_recordset_worker_6), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L17
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	if v306 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v309
	v315 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+28)) = uint8(v315)
	F_populate_recordset_record(m, v114, v16+int32(28))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L17
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L103
L102:
	;
	v306 = int32(1)
	goto L97
L103:
	;
	v340 = F_JsonbIteratorNext(m, v16+int32(36), v16+int32(40), int32(1))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L17
	} else {
		goto L106
	}
L104:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v342 != int32(18) {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	switch v340 {
	case 0:
		goto L2
	default:
		goto L103
	case 3:
		goto L105
	}
L107:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+3)))
	if v346&int32(32) == int32(0) {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v306 = int32(0)
	v309 = v345
	goto L97
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg(m, int32(_a_F_populate_recordset_worker_7), v16)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(_a_F_populate_recordset_worker_8), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L17
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v397
	goto L1
}
