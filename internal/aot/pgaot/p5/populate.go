package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_array_element_end(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v13 != v15 {
		v81 = int32(0)
		m.G0 = v9 + int32(32)
		return v81
	} else {
		v17 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v17)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v19
		if l1 != 0 {
			v21 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v21
			v33 = v21
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v24 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v24
				v33 = int32(-1)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				v33 = v29 - v27
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v33
		v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
		v39 = int32(0)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
		v48 = F_populate_record_field(m, v36, v37, v38, v39, v40, v39, v9+int32(12), v9+int32(31), v46, v39)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
			if v52 == int32(0) {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
				v65 = F_accumArrayResult(m, v60, v48, v61, v63, v64)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
					v72 = v67 + v13<<(uint(int32(2))%32) - int32(4)
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
					*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 + int32(1)
					v81 = int32(0)
					m.G0 = v9 + int32(32)
					return v81
				}
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
				if v55 != int32(447) {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					v65 = F_accumArrayResult(m, v60, v48, v61, v63, v64)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
						v72 = v67 + v13<<(uint(int32(2))%32) - int32(4)
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 + int32(1)
						v81 = int32(0)
						m.G0 = v9 + int32(32)
						return v81
					}
				} else {
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
					if v59 != 0 {
						v81 = int32(23)
						m.G0 = v9 + int32(32)
						return v81
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+31)))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
						v65 = F_accumArrayResult(m, v60, v48, v61, v63, v64)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							v72 = v67 + v13<<(uint(int32(2))%32) - int32(4)
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							*(*int32)(unsafe.Add(mBase, uint32(v72))) = v73 + int32(1)
							v81 = int32(0)
							m.G0 = v9 + int32(32)
							return v81
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v39 int32
	_ = v39
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
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v11 <= v2 {
		if v9 <= int32(0) {
			F_populate_array_report_expected_array(m, v10, v9)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v99 = int32(23)
				return v99
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v9
			v18 = v9 << (uint(int32(2)) % 32)
			v19 = F_palloc(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v19
				v24 = F_palloc0(m, v18)
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
						v39 = v2
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
							v62 = v39 + v59
							if v62 != v9&int32(2147483644) {
								v34 = v60
								v39 = v62
								continue
							} else {
								break
							}
							break
						}
						v64 = v60
					} else {
						v64 = v29
					}
					if v28 == int32(0) {
						v99 = v2
					} else {
						v74 = v64
						v75 = int32(0)
						for {
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v81+v74<<(uint(int32(2))%32)))) = int32(-1)
							v87 = int32(1)
							v90 = v75 + v87
							if v90 != v28 {
								v74 = v74 + v87
								v75 = v90
								continue
							} else {
								break
							}
							break
						}
						v99 = v2
					}
					return v99
				}
			}
		}
	} else {
		if v11 <= v9 {
			v99 = v2
			return v99
		} else {
			F_populate_array_report_expected_array(m, v10, v9)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v99 = int32(23)
				return v99
			}
		}
	}
}
func F_populate_record_field(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v441 int32
	_ = v441
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
	return int32(0)
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
	v47 = int32(1)
	v48 = v46 & v47
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v51 == v47 {
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
	v76 = int32(1)
	if v46&v76 != 0 {
		goto L35
	} else {
		goto L36
	}
L16:
	;
	v64 = int32(115)
	if v50&int32(-3) == int32(97) {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v54 == int32(1) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v57 == int32(0) {
		v74 = v50
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v74 = v50
	goto L15
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if v60 != int32(1) {
		v74 = v50
		goto L15
	} else {
		goto L22
	}
L22:
	;
	goto L16
L23:
	;
	v70 = v64
	goto L25
L24:
	;
	v70 = v50
	goto L25
L25:
	;
	if v50 == int32(67) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v73 = v64
	goto L28
L27:
	;
	v73 = v70
	goto L28
L28:
	;
	v74 = v73
	goto L15
L29:
	;
	m.G0 = v18 + int32(128)
	return v441
L30:
	;
	v427 = F_JsonbValueToJsonb(m, v117)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L142
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L139
	}
L32:
	;
	v409 = F_domain_check_safe(m, v403, v404&int32(1), l1, l0+int32(56), l4, l8)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L137
	}
L33:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v400 = F_populate_record_field(m, v396, v397, v398, l3, l4, int32(0), l6, l7, l8, l9)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L136
	}
L34:
	;
	v386 = l0 + int32(44)
	if l5 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L35:
	;
	if v74 == int32(67) {
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	switch v74 - int32(67) {
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
	v81 = int32(0)
	if v74 == int32(100) {
		v403 = v81
		v404 = v76
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v441 = v81
	goto L29
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = l0 + int32(44)
	v216 = int32(0)
	v218 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v218
	v220 = int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v223 = F_initArrayResult(m, v221, v218, v220)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L97
	}
L41:
	;
	if v51 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v200 = F_InputFunctionCallSafe(m, l0+int32(16), v192, v197, l2, l8, v18+int32(32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L89
	}
L43:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v192 = v191
	v193 = v86
	goto L42
L44:
	;
	F_escape_json(m, v18+int32(68), v86)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L88
	}
L45:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
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
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if l9 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L48:
	;
	if v87 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v93 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	F_initStringInfo(m, v18+int32(68))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v87 < int32(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	F_escape_json_with_len(m, v18+int32(68), v86, v87)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L43
L54:
	;
	v192 = v86
	v193 = v86
	goto L42
L55:
	;
	goto L56
L56:
	;
	v110 = F_palloc(m, v87+int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v87 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v115 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+v87))) = uint8(v115)
	v192 = v110
	v193 = v86
	goto L42
L59:
	;
	v112 = F__emscripten_memcpy_bulkmem(m, v110, v86, v87)
	mBase = m.M
	v113 = v112
	goto L61
L60:
	;
	v113 = v110
	goto L61
L61:
	;
	goto L58
L62:
	;
	if l1 == int32(3802) {
		goto L30
	} else {
		goto L66
	}
L63:
	;
	if v118 != int32(1) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v126 = F_pnstrdup(m, v124, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v192 = v126
	v193 = int32(0)
	goto L42
L66:
	;
	if l1 == int32(114) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L85
	}
L68:
	;
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v169 = F_JsonbToCString(m, v165, v167, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L84
	}
L69:
	;
	if v118 == int32(18) {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	switch v118 - int32(1) {
	case 0:
		goto L77
	case 1:
		goto L75
	case 2:
		goto L76
	default:
		goto L67
	case 17:
		goto L68
	}
L72:
	;
	v134 = int32(0)
	v136 = F_JsonbValueToJsonb(m, v117)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v143 = F_JsonbToCString(m, v134, v136+int32(4), int32(base.Ui32(v140)>>(uint(int32(2))%32)))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v192 = v143
	v193 = v134
	goto L42
L75:
	;
	v159 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v163 = F_DirectFunctionCall1Coll(m, int32(618), v159, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L83
	}
L76:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	if v155 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	v150 = F_pnstrdup(m, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v192 = v150
	v193 = int32(0)
	goto L42
L79:
	;
	v156 = int32(363021)
	goto L81
L80:
	;
	v156 = int32(380590)
	goto L81
L81:
	;
	v157 = F_pstrdup(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v192 = v157
	v193 = int32(0)
	goto L42
L83:
	;
	v192 = v163
	v193 = v159
	goto L42
L84:
	;
	v192 = v169
	v193 = v165
	goto L42
L85:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v175
	F_errmsg_internal(m, int32(510411), v18+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(520272), int32(3199), int32(241939))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	goto L43
L89:
	;
	if v200 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v206)
	goto L92
L91:
	;
	goto L92
L92:
	;
	if v192 != v193 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_pfree(m, v192)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v441 = v211
	goto L29
L96:
	;
	goto L95
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = l8
	*(*int64)(unsafe.Add(mBase, uint32(v18)+52)) = int64(0)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v233 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v380)
	v441 = v381
	goto L29
L99:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v309 = F_palloc(m, v306<<(uint(int32(2))%32))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L119
	}
L100:
	;
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v237 < v236 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v296 = F_populate_array_dim_jsonb(m, v18+int32(32), v232, int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L117
	}
L103:
	;
	v240 = F_strlen(m, v232)
	mBase = m.M
	v241 = v240
	goto L105
L104:
	;
	v241 = v237
	goto L105
L105:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+4))
	goto L106
L106:
	;
	v246 = F_makeJsonLexContextCstringLen(m, v236, v232, v241, v244, int32(1))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v249 = v18 + int32(84)
	v250 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v249))) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v249))) = int32(1378)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v18)+76)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = int32(1379)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(1380)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = int32(1381)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(1382)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v18 + int32(108)
	v275 = F_pg_parse_json(m, v246, v18+int32(68))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v275 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_json_errsave_error(m, v275, v246, l8)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_freeJsonLexContext(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	goto L111
L113:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	if v282 == int32(0) {
		goto L99
	} else {
		goto L114
	}
L114:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v285 != int32(447) {
		goto L99
	} else {
		goto L115
	}
L115:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+4)))
	if v288 == int32(0) {
		goto L99
	} else {
		goto L116
	}
L116:
	;
	v291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v291)
	v441 = v216
	goto L29
L117:
	;
	if v296 == int32(0) {
		v380 = v220
		v381 = v216
		goto L98
	} else {
		goto L118
	}
L118:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = v302
	goto L99
L119:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(0) < v311 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v321 = int32(0)
	goto L123
L121:
	;
	v350 = v311
	goto L122
L122:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v358 = F_makeMdArrayResult(m, v354, v350, v355, v309, v356, int32(1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L126
	}
L123:
	;
	v333 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v309+v321<<(uint(int32(2))%32)))) = v333
	v336 = v321 + v333
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v336 < v337 {
		v321 = v336
		goto L123
	} else {
		goto L125
	}
L124:
	;
	v350 = v337
	goto L122
L125:
	;
	goto L124
L126:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	F_pfree(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	F_pfree(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_pfree(m, v309)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v380 = int32(0)
	v381 = v358
	goto L98
L130:
	;
	v390 = F_populate_composite(m, v386, l1, l4, int32(0), l6, l7, l8)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v392 = F_pg_detoast_datum(m, l5)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	v441 = v390
	goto L29
L134:
	;
	v394 = F_populate_composite(m, v386, l1, l4, v392, l6, l7, l8)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v441 = v394
	goto L29
L136:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7))))
	v403 = v400
	v404 = v402
	goto L32
L137:
	;
	if v409 != 0 {
		v441 = v403
		goto L29
	} else {
		goto L138
	}
L138:
	;
	v411 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v411)
	v441 = int32(0)
	goto L29
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v74
	F_errmsg_internal(m, int32(723838), v18)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(520272), int32(3470), int32(453757))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L142:
	;
	v441 = v427
	goto L29
}
func F_populate_record_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
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
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v3 = l2
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v14
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v20 == v6 {
		v24 = F_MemoryContextAllocZero(m, v19, int32(72))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v24
			*(*int32)(unsafe.Add(mBase, uint32(v24)+68)) = v19
			if l3 != 0 {
				F_get_record_type_from_argument(m, l0, l1, v24)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v37 = v24
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v38 == int32(0) {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v42 = F_pg_detoast_datum(m, v41)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
							if v44 != int32(2249) {
								v56 = v37
								v57 = v42
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v47
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v49
								v56 = v37
								v57 = v42
							}
							v60 = l0 + l3<<(uint(int32(3))%32)
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
							if v61 == int32(1) {
								if v57 != 0 {
									v146 = v57
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									v146 = int32(0)
								}
								m.G0 = v12 + int32(48)
								return v146
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
								if v3 != 0 {
									v69 = F_pg_detoast_datum_packed(m, v68)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = int32(1)
										v72 = v69 + v71
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
										v77 = v75 & v71
										if v77 != 0 {
											v78 = v72
										} else {
											v78 = v69 + int32(4)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
										if v75 == int32(1) {
											v82 = int32(4)
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
											if v84&int32(254) == int32(2) {
												v93 = v82
											} else {
												v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
											}
											if v84 == int32(1) {
												v96 = v82
											} else {
												v96 = v93
											}
											v107 = v96
										} else {
											v97 = int32(1)
											if v77 != 0 {
												v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
												v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
										v131 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v146 = v140
											m.G0 = v12 + int32(48)
											return v146
										}
									}
								} else {
									v111 = F_pg_detoast_datum(m, v68)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
										v115 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
										v131 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v146 = v140
											m.G0 = v12 + int32(48)
											return v146
										}
									}
								}
							}
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						if v51 != int32(2249) {
							v56 = v37
							v57 = v6
							v60 = l0 + l3<<(uint(int32(3))%32)
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
							if v61 == int32(1) {
								if v57 != 0 {
									v146 = v57
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									v146 = int32(0)
								}
								m.G0 = v12 + int32(48)
								return v146
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
								if v3 != 0 {
									v69 = F_pg_detoast_datum_packed(m, v68)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v71 = int32(1)
										v72 = v69 + v71
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
										v77 = v75 & v71
										if v77 != 0 {
											v78 = v72
										} else {
											v78 = v69 + int32(4)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
										if v75 == int32(1) {
											v82 = int32(4)
											v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
											if v84&int32(254) == int32(2) {
												v93 = v82
											} else {
												v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
											}
											if v84 == int32(1) {
												v96 = v82
											} else {
												v96 = v93
											}
											v107 = v96
										} else {
											v97 = int32(1)
											if v77 != 0 {
												v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
											} else {
												v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
												v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
										v131 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v146 = v140
											m.G0 = v12 + int32(48)
											return v146
										}
									}
								} else {
									v111 = F_pg_detoast_datum(m, v68)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
										v115 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
										v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
										v131 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											v146 = v140
											m.G0 = v12 + int32(48)
											return v146
										}
									}
								}
							}
						} else {
							F_get_record_type_from_query(m, l0, l1, v37)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v56 = v37
								v57 = v6
								v60 = l0 + l3<<(uint(int32(3))%32)
								v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
								if v61 == int32(1) {
									if v57 != 0 {
										v146 = v57
									} else {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
										v146 = int32(0)
									}
									m.G0 = v12 + int32(48)
									return v146
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
									if v3 != 0 {
										v69 = F_pg_detoast_datum_packed(m, v68)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int32(0)
										} else {
											v71 = int32(1)
											v72 = v69 + v71
											v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
											v77 = v75 & v71
											if v77 != 0 {
												v78 = v72
											} else {
												v78 = v69 + int32(4)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
											if v75 == int32(1) {
												v82 = int32(4)
												v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
												if v84&int32(254) == int32(2) {
													v93 = v82
												} else {
													v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
												}
												if v84 == int32(1) {
													v96 = v82
												} else {
													v96 = v93
												}
												v107 = v96
											} else {
												v97 = int32(1)
												if v77 != 0 {
													v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
												} else {
													v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
													v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
											v131 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v146 = v140
												m.G0 = v12 + int32(48)
												return v146
											}
										}
									} else {
										v111 = F_pg_detoast_datum(m, v68)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
											v115 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
											v131 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
											v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												v146 = v140
												m.G0 = v12 + int32(48)
												return v146
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_get_record_type_from_query(m, l0, l1, v24)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v56 = v24
					v57 = v6
					v60 = l0 + l3<<(uint(int32(3))%32)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
					if v61 == int32(1) {
						if v57 != 0 {
							v146 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v146 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v146
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
						if v3 != 0 {
							v69 = F_pg_detoast_datum_packed(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(1)
								v72 = v69 + v71
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								v77 = v75 & v71
								if v77 != 0 {
									v78 = v72
								} else {
									v78 = v69 + int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
								if v75 == int32(1) {
									v82 = int32(4)
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v84&int32(254) == int32(2) {
										v93 = v82
									} else {
										v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
									}
									if v84 == int32(1) {
										v96 = v82
									} else {
										v96 = v93
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						}
					}
				}
			}
		}
	} else {
		if l3 == int32(0) {
			v56 = v20
			v57 = v6
			v60 = l0 + l3<<(uint(int32(3))%32)
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
			if v61 == int32(1) {
				if v57 != 0 {
					v146 = v57
				} else {
					v64 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
					v146 = int32(0)
				}
				m.G0 = v12 + int32(48)
				return v146
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
				if v3 != 0 {
					v69 = F_pg_detoast_datum_packed(m, v68)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = int32(1)
						v72 = v69 + v71
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
						v77 = v75 & v71
						if v77 != 0 {
							v78 = v72
						} else {
							v78 = v69 + int32(4)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
						if v75 == int32(1) {
							v82 = int32(4)
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
							if v84&int32(254) == int32(2) {
								v93 = v82
							} else {
								v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
							}
							if v84 == int32(1) {
								v96 = v82
							} else {
								v96 = v93
							}
							v107 = v96
						} else {
							v97 = int32(1)
							if v77 != 0 {
								v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
								v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
						v131 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							v146 = v140
							m.G0 = v12 + int32(48)
							return v146
						}
					}
				} else {
					v111 = F_pg_detoast_datum(m, v68)
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
						v115 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
						v131 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return int32(0)
						} else {
							v146 = v140
							m.G0 = v12 + int32(48)
							return v146
						}
					}
				}
			}
		} else {
			v37 = v20
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v38 == int32(0) {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v42 = F_pg_detoast_datum(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v44 != int32(2249) {
						v56 = v37
						v57 = v42
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v47
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v37)+60)) = v49
						v56 = v37
						v57 = v42
					}
					v60 = l0 + l3<<(uint(int32(3))%32)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
					if v61 == int32(1) {
						if v57 != 0 {
							v146 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v146 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v146
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
						if v3 != 0 {
							v69 = F_pg_detoast_datum_packed(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(1)
								v72 = v69 + v71
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								v77 = v75 & v71
								if v77 != 0 {
									v78 = v72
								} else {
									v78 = v69 + int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
								if v75 == int32(1) {
									v82 = int32(4)
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v84&int32(254) == int32(2) {
										v93 = v82
									} else {
										v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
									}
									if v84 == int32(1) {
										v96 = v82
									} else {
										v96 = v93
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						}
					}
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				if v51 != int32(2249) {
					v56 = v37
					v57 = v6
					v60 = l0 + l3<<(uint(int32(3))%32)
					v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
					if v61 == int32(1) {
						if v57 != 0 {
							v146 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v146 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v146
					} else {
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
						if v3 != 0 {
							v69 = F_pg_detoast_datum_packed(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v71 = int32(1)
								v72 = v69 + v71
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
								v77 = v75 & v71
								if v77 != 0 {
									v78 = v72
								} else {
									v78 = v69 + int32(4)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
								if v75 == int32(1) {
									v82 = int32(4)
									v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v84&int32(254) == int32(2) {
										v93 = v82
									} else {
										v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
									}
									if v84 == int32(1) {
										v96 = v82
									} else {
										v96 = v93
									}
									v107 = v96
								} else {
									v97 = int32(1)
									if v77 != 0 {
										v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
									} else {
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						} else {
							v111 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v115 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
								v131 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
								v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int32(0)
								} else {
									v146 = v140
									m.G0 = v12 + int32(48)
									return v146
								}
							}
						}
					}
				} else {
					F_get_record_type_from_query(m, l0, l1, v37)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = v37
						v57 = v6
						v60 = l0 + l3<<(uint(int32(3))%32)
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+24)))
						if v61 == int32(1) {
							if v57 != 0 {
								v146 = v57
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
								v146 = int32(0)
							}
							m.G0 = v12 + int32(48)
							return v146
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v3)
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
							if v3 != 0 {
								v69 = F_pg_detoast_datum_packed(m, v68)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v71 = int32(1)
									v72 = v69 + v71
									v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
									v77 = v75 & v71
									if v77 != 0 {
										v78 = v72
									} else {
										v78 = v69 + int32(4)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v78
									if v75 == int32(1) {
										v82 = int32(4)
										v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
										if v84&int32(254) == int32(2) {
											v93 = v82
										} else {
											v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
										}
										if v84 == int32(1) {
											v96 = v82
										} else {
											v96 = v93
										}
										v107 = v96
									} else {
										v97 = int32(1)
										if v77 != 0 {
											v107 = int32(base.Ui32(v75)>>(uint(v97)%32)) - v97
										} else {
											v101 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
											v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v107
									v131 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										v146 = v140
										m.G0 = v12 + int32(48)
										return v146
									}
								}
							} else {
								v111 = F_pg_detoast_datum(m, v68)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
									v115 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v111 + v115
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v121)>>(uint(int32(2))%32)) - v115
									v131 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v131)
									v135 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v140 = F_populate_composite(m, v56+int32(48), v135, v19, v57, v12+int32(32), v12+int32(31), l4)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int32(0)
									} else {
										v146 = v140
										m.G0 = v12 + int32(48)
										return v146
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	if int32(2) < v11 {
		m.G0 = v8 + int32(16)
		return int32(0)
	} else {
		v14 = F_strlen(m, l1)
		mBase = m.M
		if base.Ui32(int32(63)) < base.Ui32(v14) {
			m.G0 = v8 + int32(16)
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v21 = F_hash_search(m, v17, l1, int32(1), v8+int32(15))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v25
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v30 = v29 - v27
					v33 = F_palloc(m, v30+int32(1))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						if v30 != 0 {
							v36 = F__emscripten_memcpy_bulkmem(m, v33, v35, v30)
							mBase = m.M
							v37 = v36
						} else {
							v37 = v33
						}
						v39 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v37+v30))) = uint8(v39)
						v42 = v33
						*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v42
						m.G0 = v8 + int32(16)
						return int32(0)
					}
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v42 = v41
					*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v42
					m.G0 = v8 + int32(16)
					return int32(0)
				}
			}
		}
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v420 int32
	_ = v420
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
	v420 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v420)
	m.G0 = v16 + int32(112)
	return
L2:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v115)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	v404 = F_CreateTupleDescCopy(m, v403)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L17
	} else {
		goto L113
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L17
	} else {
		goto L109
	}
L4:
	;
	v310 = v305
	v316 = v304
	goto L97
L5:
	;
	v304 = v5
	v305 = int32(1)
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L17
	} else {
		goto L93
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L17
	} else {
		goto L89
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L17
	} else {
		goto L85
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v21 != int32(383) {
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
	v74 = l0 + l3<<(uint(int32(3))%32)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+24)))
	if v75 != 0 {
		goto L1
	} else {
		goto L32
	}
L13:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
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
		v70 = v5
		v71 = v30
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
	v70 = v5
	v71 = v38
	goto L12
L24:
	;
	v51 = v30
	goto L13
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
		v70 = v5
		v71 = v51
		goto L12
	} else {
		goto L30
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v58 != int32(2249) {
		v70 = v56
		v71 = v51
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
	v70 = v56
	v71 = v51
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
	v70 = v5
	v71 = v51
	goto L12
L32:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71)+68))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
	if v79 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v115 = F_palloc0(m, int32(36))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L48
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v71)+60))
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	v89 = v82
	goto L34
L36:
	;
	goto L37
L37:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if v83 != v84 {
		v89 = v84
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71)+60))
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
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v71)+52))
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
	v97 = int32(4562096)
	v98 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v78
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
	*(*int32)(unsafe.Add(mBase, uint32(v71)+52)) = v101
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v98
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
	v117 = int32(4562096)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[21]))
	v131 = F_tuplestore_begin_heap(m, int32(base.Ui32(v123&int32(4))>>(uint(int32(2))%32)), int32(0), v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L17
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+24)) = v131
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v115)+32)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v115)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v115)+28)) = v70
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(20))))
	if l2 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v140 = F_pg_detoast_datum_packed(m, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L17
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v224 = F_pg_detoast_datum(m, v139)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L17
	} else {
		goto L78
	}
L53:
	;
	v143 = F_palloc0(m, int32(40))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v147 = F_pg_detoast_datum_packed(m, v140)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v149 = int32(1)
	v150 = v147 + v149
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v155 = v153 & v149
	if v155 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v156 = v150
	goto L58
L57:
	;
	v156 = v147 + int32(4)
	goto L58
L58:
	;
	if v153 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, _consts[247]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	goto L70
L60:
	;
	v159 = int32(4)
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	if v161&int32(254) == int32(2) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v174 = int32(1)
	if v155 != 0 {
		v184 = int32(base.Ui32(v153)>>(uint(v174)%32)) - v174
		goto L59
	} else {
		goto L69
	}
L63:
	;
	v170 = v159
	goto L65
L64:
	;
	v170 = base.B2i32(v161 == int32(18)) << (uint(v159) % 32)
	goto L65
L65:
	;
	if v161 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v173 = v159
	goto L68
L67:
	;
	v173 = v170
	goto L68
L68:
	;
	v184 = v173
	goto L59
L69:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v184 = int32(base.Ui32(v178)>>(uint(int32(2))%32)) - int32(4)
	goto L59
L70:
	;
	v189 = F_makeJsonLexContextCstringLen(m, v16+int32(44), v156, v184, v187, int32(1))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v143)+36)) = int32(1383)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+28)) = int32(1384)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+12)) = int32(1385)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = int32(1386)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = int32(1387)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = int32(1388)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = int32(1389)
	v207 = v16 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = v207
	v211 = F_pg_parse_json(m, v207, v143)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	if v211 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_json_errsave_error(m, v211, v16+int32(44), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L17
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_freeJsonLexContext(m, v16+int32(44))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L17
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115))) = int32(0)
	goto L2
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	if v226&int32(1342177280) != int32(1073741824) {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v233 = F_JsonbIteratorInit(m, v224+int32(4))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v233
	v241 = F_JsonbIteratorNext(m, v16+int32(40), v16+int32(44), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L17
	} else {
		goto L82
	}
L81:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v243 != int32(18) {
		goto L3
	} else {
		goto L83
	}
L82:
	;
	switch v241 {
	case 0:
		goto L2
	default:
		goto L5
	case 3:
		goto L81
	}
L83:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+3)))
	if v247&int32(32) == int32(0) {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v304 = v246
	v305 = int32(0)
	goto L4
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(114122), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(520272), int32(4056), int32(231805))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(66279), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(520272), int32(4061), int32(231805))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	v291 = m.ExcPending
	if v291 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
	F_errmsg(m, int32(25234), v16+int32(16))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(520272), int32(4176), int32(231805))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
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
	if v310 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v316
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)) = uint8(v322)
	F_populate_recordset_record(m, v115, v16+int32(32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
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
	v310 = int32(1)
	goto L97
L103:
	;
	v347 = F_JsonbIteratorNext(m, v16+int32(40), v16+int32(44), int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L17
	} else {
		goto L106
	}
L104:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v349 != int32(18) {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	switch v347 {
	case 0:
		goto L2
	default:
		goto L103
	case 3:
		goto L105
	}
L107:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v352)+3)))
	if v353&int32(32) == int32(0) {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v310 = int32(0)
	v316 = v352
	goto L97
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg(m, int32(133881), v16)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(520272), int32(4193), int32(231805))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v404
	goto L1
}
