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
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	if v13 != v15 {
		v80 = int32(0)
		m.G0 = v9 + int32(32)
		return v80
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
					v80 = int32(0)
					m.G0 = v9 + int32(32)
					return v80
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
						v80 = int32(0)
						m.G0 = v9 + int32(32)
						return v80
					}
				} else {
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+4)))
					if v59 != 0 {
						v80 = int32(23)
						m.G0 = v9 + int32(32)
						return v80
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
							v80 = int32(0)
							m.G0 = v9 + int32(32)
							return v80
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
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
	v427 = F_JsonbValueToJsonb(m, v116)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L141
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L138
	}
L32:
	;
	v409 = F_domain_check_safe(m, v403, v404&int32(1), l1, l0+int32(56), l4, l8)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L136
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
		goto L135
	}
L34:
	;
	v386 = l0 + int32(44)
	if l5 == int32(0) {
		goto L129
	} else {
		goto L130
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
	v218 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_populate_record_field[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v220
	v222 = int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v225 = F_initArrayResult(m, v223, v220, v222)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L96
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
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v202 = F_InputFunctionCallSafe(m, l0+int32(16), v195, v199, l2, l8, v18+int32(32))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L88
	}
L43:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+68))
	v195 = v191
	v196 = v86
	goto L42
L44:
	;
	F_escape_json(m, v18+int32(68), v86)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L87
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
	v114 = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if base.B2i32(l9 == v114)|base.B2i32(v117 != int32(1)) == v114 {
		goto L61
	} else {
		goto L62
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
	v97 = v18 + int32(68)
	F_initStringInfo(m, v97)
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
	F_escape_json_with_len(m, v97, v86, v87)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	goto L43
L54:
	;
	v195 = v86
	v196 = v86
	goto L42
L55:
	;
	goto L56
L56:
	;
	v108 = F_palloc(m, v87+int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v87 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	base.MemoryCopy(m, v108, v86, v87)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v108))) = uint8(v112)
	v195 = v108
	v196 = v86
	goto L42
L61:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v126 = F_pnstrdup(m, v124, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	v195 = v126
	v196 = int32(0)
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
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L84
	}
L67:
	;
	v165 = int32(0)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v169 = F_JsonbToCString(m, v165, v167, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L83
	}
L68:
	;
	if v117 == int32(18) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	switch v117 - int32(1) {
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
	v134 = int32(0)
	v136 = F_JsonbValueToJsonb(m, v116)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v143 = F_JsonbToCString(m, v134, v136+int32(4), int32(base.Ui32(v140)>>(uint(int32(2))%32)))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v195 = v143
	v196 = v134
	goto L42
L74:
	;
	v159 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v163 = F_DirectFunctionCall1Coll(m, int32(618), v159, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L82
	}
L75:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+4)))
	if v155 != 0 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v150 = F_pnstrdup(m, v148, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v195 = v150
	v196 = int32(0)
	goto L42
L78:
	;
	v156 = int32(_a_F_populate_record_field_3)
	goto L80
L79:
	;
	v156 = int32(_a_F_populate_record_field_4)
	goto L80
L80:
	;
	v157 = F_pstrdup(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v195 = v157
	v196 = int32(0)
	goto L42
L82:
	;
	v195 = v163
	v196 = v159
	goto L42
L83:
	;
	v195 = v169
	v196 = v165
	goto L42
L84:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v175
	F_errmsg_internal(m, int32(_a_F_populate_record_field_5), v18+int32(16))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_populate_record_field_1), int32(3199), int32(_a_F_populate_record_field_6))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	if v202 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(0)
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v208)
	goto L91
L90:
	;
	goto L91
L91:
	;
	if v195 != v196 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_pfree(m, v195)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v441 = v213
	goto L29
L95:
	;
	goto L94
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v225
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = l8
	*(*int64)(unsafe.Add(mBase, uint32(v18)+52)) = int64(0)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6))))
	if v235 == int32(1) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v380)
	v441 = v381
	goto L29
L98:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	v309 = F_palloc(m, v306<<(uint(int32(2))%32))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L118
	}
L99:
	;
	v238 = int32(0)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v239 < v238 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v296 = F_populate_array_dim_jsonb(m, v18+int32(32), v234, int32(1))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L116
	}
L102:
	;
	v242 = F_strlen(m, v234)
	mBase = m.M
	v243 = v242
	goto L104
L103:
	;
	v243 = v239
	goto L104
L104:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_populate_record_field[1]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	goto L105
L105:
	;
	v248 = F_makeJsonLexContextCstringLen(m, v238, v234, v243, v246, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v250 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+84)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v248
	*(*int64)(unsafe.Add(mBase, uint32(v18)+76)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = int32(1359)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+100)) = int32(1360)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = int32(1361)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(1362)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = int32(1363)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v18 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v18 + int32(108)
	v275 = F_pg_parse_json(m, v248, v18+int32(68))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if v275 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F_json_errsave_error(m, v275, v248, l8)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_freeJsonLexContext(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	if v282 == int32(0) {
		goto L98
	} else {
		goto L113
	}
L113:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v285 != int32(447) {
		goto L98
	} else {
		goto L114
	}
L114:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+4)))
	if v288 == int32(0) {
		goto L98
	} else {
		goto L115
	}
L115:
	;
	v291 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v291)
	v441 = v218
	goto L29
L116:
	;
	if v296 == int32(0) {
		v380 = v222
		v381 = v218
		goto L97
	} else {
		goto L117
	}
L117:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v300))) = v302
	goto L98
L118:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(0) < v311 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v321 = int32(0)
	goto L122
L120:
	;
	v350 = v311
	goto L121
L121:
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
		goto L125
	}
L122:
	;
	v333 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v309+v321<<(uint(int32(2))%32)))) = v333
	v336 = v321 + v333
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if v336 < v337 {
		v321 = v336
		goto L122
	} else {
		goto L124
	}
L123:
	;
	v350 = v337
	goto L121
L124:
	;
	goto L123
L125:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v18)+52))
	F_pfree(m, v360)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v18)+56))
	F_pfree(m, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	F_pfree(m, v309)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	v380 = int32(0)
	v381 = v358
	goto L97
L129:
	;
	v390 = F_populate_composite(m, v386, l1, l4, int32(0), l6, l7, l8)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v392 = F_pg_detoast_datum(m, l5)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L133
	}
L132:
	;
	v441 = v390
	goto L29
L133:
	;
	v394 = F_populate_composite(m, v386, l1, l4, v392, l6, l7, l8)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v441 = v394
	goto L29
L135:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7))))
	v403 = v400
	v404 = v402
	goto L32
L136:
	;
	if v409 != 0 {
		v441 = v403
		goto L29
	} else {
		goto L137
	}
L137:
	;
	v411 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v411)
	v441 = int32(0)
	goto L29
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v74
	F_errmsg_internal(m, int32(_a_F_populate_record_field_0), v18)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_populate_record_field_1), int32(3470), int32(_a_F_populate_record_field_2))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
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
									v145 = v57
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									v145 = int32(0)
								}
								m.G0 = v12 + int32(48)
								return v145
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
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
											if v85 == int32(18) {
												v88 = int32(16)
											} else {
												v88 = int32(0)
											}
											if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v95 = int32(4)
											} else {
												v95 = v88
											}
											v106 = v95
										} else {
											v96 = int32(1)
											if v77 != 0 {
												v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
												v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											v145 = v139
											m.G0 = v12 + int32(48)
											return v145
										}
									}
								} else {
									v110 = F_pg_detoast_datum(m, v68)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
										v114 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											v145 = v139
											m.G0 = v12 + int32(48)
											return v145
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
									v145 = v57
								} else {
									v64 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
									v145 = int32(0)
								}
								m.G0 = v12 + int32(48)
								return v145
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
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
											if v85 == int32(18) {
												v88 = int32(16)
											} else {
												v88 = int32(0)
											}
											if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
												v95 = int32(4)
											} else {
												v95 = v88
											}
											v106 = v95
										} else {
											v96 = int32(1)
											if v77 != 0 {
												v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
												v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
											}
										}
										*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											v145 = v139
											m.G0 = v12 + int32(48)
											return v145
										}
									}
								} else {
									v110 = F_pg_detoast_datum(m, v68)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
										v114 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
										*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
										v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
										v130 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return int32(0)
										} else {
											v145 = v139
											m.G0 = v12 + int32(48)
											return v145
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
										v145 = v57
									} else {
										v64 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
										v145 = int32(0)
									}
									m.G0 = v12 + int32(48)
									return v145
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
												v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
												if v85 == int32(18) {
													v88 = int32(16)
												} else {
													v88 = int32(0)
												}
												if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
													v95 = int32(4)
												} else {
													v95 = v88
												}
												v106 = v95
											} else {
												v96 = int32(1)
												if v77 != 0 {
													v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
													v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
												}
											}
											*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
											v130 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v145 = v139
												m.G0 = v12 + int32(48)
												return v145
											}
										}
									} else {
										v110 = F_pg_detoast_datum(m, v68)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
											v114 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
											*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
											*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
											v130 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
											v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												v145 = v139
												m.G0 = v12 + int32(48)
												return v145
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
							v145 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v145 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v145
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
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v85 == int32(18) {
										v88 = int32(16)
									} else {
										v88 = int32(0)
									}
									if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v95 = int32(4)
									} else {
										v95 = v88
									}
									v106 = v95
								} else {
									v96 = int32(1)
									if v77 != 0 {
										v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
								}
							}
						} else {
							v110 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v114 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
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
					v145 = v57
				} else {
					v64 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
					v145 = int32(0)
				}
				m.G0 = v12 + int32(48)
				return v145
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
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
							if v85 == int32(18) {
								v88 = int32(16)
							} else {
								v88 = int32(0)
							}
							if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v95 = int32(4)
							} else {
								v95 = v88
							}
							v106 = v95
						} else {
							v96 = int32(1)
							if v77 != 0 {
								v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
							} else {
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
								v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
						v130 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							v145 = v139
							m.G0 = v12 + int32(48)
							return v145
						}
					}
				} else {
					v110 = F_pg_detoast_datum(m, v68)
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
						v114 = int32(4)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
						*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
						v130 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
						v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							v145 = v139
							m.G0 = v12 + int32(48)
							return v145
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
							v145 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v145 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v145
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
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v85 == int32(18) {
										v88 = int32(16)
									} else {
										v88 = int32(0)
									}
									if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v95 = int32(4)
									} else {
										v95 = v88
									}
									v106 = v95
								} else {
									v96 = int32(1)
									if v77 != 0 {
										v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
								}
							}
						} else {
							v110 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v114 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
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
							v145 = v57
						} else {
							v64 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
							v145 = int32(0)
						}
						m.G0 = v12 + int32(48)
						return v145
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
									v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
									if v85 == int32(18) {
										v88 = int32(16)
									} else {
										v88 = int32(0)
									}
									if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v95 = int32(4)
									} else {
										v95 = v88
									}
									v106 = v95
								} else {
									v96 = int32(1)
									if v77 != 0 {
										v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
										v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
								}
							}
						} else {
							v110 = F_pg_detoast_datum(m, v68)
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
								v114 = int32(4)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
								v130 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									v145 = v139
									m.G0 = v12 + int32(48)
									return v145
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
								v145 = v57
							} else {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v64)
								v145 = int32(0)
							}
							m.G0 = v12 + int32(48)
							return v145
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
										v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
										if v85 == int32(18) {
											v88 = int32(16)
										} else {
											v88 = int32(0)
										}
										if base.Ui32((v85-int32(1))&int32(255)) < base.Ui32(int32(3)) {
											v95 = int32(4)
										} else {
											v95 = v88
										}
										v106 = v95
									} else {
										v96 = int32(1)
										if v77 != 0 {
											v106 = int32(base.Ui32(v75)>>(uint(v96)%32)) - v96
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
											v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
										}
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v106
									v130 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										v145 = v139
										m.G0 = v12 + int32(48)
										return v145
									}
								}
							} else {
								v110 = F_pg_detoast_datum(m, v68)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(18)
									v114 = int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v110 + v114
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v12 + int32(8)
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(base.Ui32(v120)>>(uint(int32(2))%32)) - v114
									v130 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v12)+31)) = uint8(v130)
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									v139 = F_populate_composite(m, v56+int32(48), v134, v19, v57, v12+int32(32), v12+int32(31), l4)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										v145 = v139
										m.G0 = v12 + int32(48)
										return v145
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
	v5 = Fn13997(m, l0, l1, l2, int32(2))
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
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
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
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v414 int32
	_ = v414
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
	v414 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v414)
	m.G0 = v16 + int32(112)
	return
L2:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v114)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v70)+52))
	v398 = F_CreateTupleDescCopy(m, v397)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L17
	} else {
		goto L113
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L17
	} else {
		goto L109
	}
L4:
	;
	v307 = v299
	v310 = v298
	goto L97
L5:
	;
	v298 = v5
	v299 = int32(1)
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L17
	} else {
		goto L93
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L17
	} else {
		goto L89
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
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
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v74+int32(20))))
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
	v218 = F_pg_detoast_datum(m, v138)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
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
	v146 = F_pg_detoast_datum_packed(m, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L17
	} else {
		goto L55
	}
L55:
	;
	v148 = int32(1)
	v149 = v146 + v148
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	v154 = v152 & v148
	if v154 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v155 = v149
	goto L58
L57:
	;
	v155 = v146 + int32(4)
	goto L58
L58:
	;
	if v152 == int32(1) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_populate_recordset_worker[2]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	goto L70
L60:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v161 == int32(18) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v172 = int32(1)
	if v154 != 0 {
		v182 = int32(base.Ui32(v152)>>(uint(v172)%32)) - v172
		goto L59
	} else {
		goto L69
	}
L63:
	;
	v164 = int32(16)
	goto L65
L64:
	;
	v164 = int32(0)
	goto L65
L65:
	;
	if base.Ui32((v161-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v171 = int32(4)
	goto L68
L67:
	;
	v171 = v164
	goto L68
L68:
	;
	v182 = v171
	goto L59
L69:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v182 = int32(base.Ui32(v176)>>(uint(int32(2))%32)) - int32(4)
	goto L59
L70:
	;
	v187 = F_makeJsonLexContextCstringLen(m, v16+int32(44), v155, v182, v185, int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L17
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+36)) = int32(1364)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+28)) = int32(1365)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+12)) = int32(1366)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v142)+24)) = int32(1367)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1368)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+8)) = int32(1369)
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = int32(1370)
	v205 = v16 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v205
	v207 = F_pg_parse_json(m, v205, v142)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L17
	} else {
		goto L72
	}
L72:
	;
	if v207 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_json_errsave_error(m, v207, v205, int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
	v215 = m.ExcPending
	if v215 != 0 {
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
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v220&int32(1342177280) != int32(1073741824) {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v227 = F_JsonbIteratorInit(m, v218+int32(4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v227
	v235 = F_JsonbIteratorNext(m, v16+int32(40), v16+int32(44), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L17
	} else {
		goto L82
	}
L81:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v237 != int32(18) {
		goto L3
	} else {
		goto L83
	}
L82:
	;
	switch v235 {
	case 0:
		goto L2
	default:
		goto L5
	case 3:
		goto L81
	}
L83:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+3)))
	if v241&int32(32) == int32(0) {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v298 = v240
	v299 = int32(0)
	goto L4
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L17
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_populate_recordset_worker_1), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(4056), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
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
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(_a_F_populate_recordset_worker_4), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L17
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(4061), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
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
	v285 = m.ExcPending
	if v285 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l1
	F_errmsg(m, int32(_a_F_populate_recordset_worker_5), v16+int32(16))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(_a_F_populate_recordset_worker_6), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
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
	if v307 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v310
	v316 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+32)) = uint8(v316)
	F_populate_recordset_record(m, v114, v16+int32(32))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
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
	v307 = int32(1)
	goto L97
L103:
	;
	v341 = F_JsonbIteratorNext(m, v16+int32(40), v16+int32(44), int32(1))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L17
	} else {
		goto L106
	}
L104:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	if v343 != int32(18) {
		goto L3
	} else {
		goto L107
	}
L105:
	;
	goto L104
L106:
	;
	switch v341 {
	case 0:
		goto L2
	default:
		goto L103
	case 3:
		goto L105
	}
L107:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+3)))
	if v347&int32(32) == int32(0) {
		goto L3
	} else {
		goto L108
	}
L108:
	;
	v307 = int32(0)
	v310 = v346
	goto L97
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L17
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg(m, int32(_a_F_populate_recordset_worker_7), v16)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L17
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_populate_recordset_worker_2), int32(_a_F_populate_recordset_worker_8), int32(_a_F_populate_recordset_worker_3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v398
	goto L1
}
