package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v231 float64
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l0 < int32(_a_F_uuid_abbrev_abort_0) {
		v326 = v3
		m.G0 = v10 + int32(96)
		return v326
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		if v15 < int64(10000) {
			v326 = v3
			m.G0 = v10 + int32(96)
			return v326
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)))
			if v18 != int32(1) {
				v326 = v3
				m.G0 = v10 + int32(96)
				return v326
			} else {
				v22 = v14 + int32(16)
				v23 = float64(0)
				v25 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v32 != 0 {
					v33 = int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					if v32 == v33 {
						v69 = v23
						v71 = v25
					} else {
						v41 = v23
						v43 = v25
						v45 = v25
						for {
							v50 = float64(1)
							v51 = v43 + v35
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
							v53 = F_ldexp(m, v50, v52)
							mBase = m.M
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v57 = F_ldexp(m, v50, v56)
							mBase = m.M
							v62 = base.F64_add(base.F64_add(v41, base.F64_div(v50, v57)), base.F64_div(v50, v53))
							v63 = int32(2)
							v64 = v43 + v63
							v66 = v45 + v63
							if v66 != v32&int32(-2) {
								v41 = v62
								v43 = v64
								v45 = v66
								continue
							} else {
								break
							}
							break
						}
						v69 = v62
						v71 = v64
					}
					if v32&v33 != 0 {
						v78 = float64(1)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v35))))
						v82 = F_ldexp(m, v78, v81)
						mBase = m.M
						v85 = base.F64_add(v69, base.F64_div(v78, v82))
					} else {
						v85 = v69
					}
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v87 = base.F64_div(v86, v85)
					v88 = base.F64_convert_i32_u(v32)
					if base.F64_le(v87, base.F64_mul(v88, float64(2.5))) == int32(0) {
						v191 = v87
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
						v231 = v212
					} else {
						v95 = v32 & int32(3)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						v97 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v108 = v97
							v109 = int32(0)
							v110 = v97
							for {
								v115 = v108 + v96
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
								v117 = int32(0)
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
								v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
								v131 = v110 + base.B2i32(v116 == v117) + base.B2i32(v120 == v117) + base.B2i32(v124 == v117) + base.B2i32(v128 == v117)
								v132 = int32(4)
								v133 = v108 + v132
								v135 = v109 + v132
								if v135 != v32&int32(-4) {
									v108 = v133
									v109 = v135
									v110 = v131
									continue
								} else {
									break
								}
								break
							}
							v140 = v133
							v142 = v131
						} else {
							v140 = v97
							v142 = v97
						}
						if v95 != 0 {
							v150 = v140
							v152 = v142
							v153 = v97
							for {
								v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v96))))
								v161 = v152 + base.B2i32(v158 == int32(0))
								v162 = int32(1)
								v165 = v153 + v162
								if v165 != v95 {
									v150 = v150 + v162
									v152 = v161
									v153 = v165
									continue
								} else {
									break
								}
								break
							}
							v172 = v161
						} else {
							v172 = v142
						}
						if v172 == int32(0) {
							v212 = v87
							v231 = v212
						} else {
							v181 = F_log(m, base.F64_div(v88, base.F64_convert_i32_s(v172)))
							mBase = m.M
							v231 = base.F64_mul(v181, v88)
						}
					}
				} else {
					v183 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v185 = base.F64_div(v183, float64(0))
					if base.F64_le(v185, base.F64_mul(base.F64_convert_i32_u(v32), float64(2.5))) != 0 {
						v212 = v185
					} else {
						v191 = v185
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
					}
					v231 = v212
				}
				if base.F64_gt(v231, float64(100000)) != 0 {
					v235 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_uuid_abbrev_abort[0])))
					if v235 != int32(1) {
						v259 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
						v326 = v3
						m.G0 = v10 + int32(96)
						return v326
					} else {
						v240 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							if v240 == int32(0) {
								v259 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
								v326 = v3
								m.G0 = v10 + int32(96)
								return v326
							} else {
								v246 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v246
								*(*float64)(unsafe.Add(mBase, uint32(v10))) = v231
								F_errmsg_internal(m, int32(_a_F_uuid_abbrev_abort_1), v10)
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_uuid_abbrev_abort_2), int32(356), int32(_a_F_uuid_abbrev_abort_3))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										v259 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+8)) = uint8(v259)
										v326 = v3
										m.G0 = v10 + int32(96)
										return v326
									}
								}
							}
						}
					}
				} else {
					v262 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_uuid_abbrev_abort[0])))
					v263 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v263), float64(2000)), float64(0.5)), v231) != 0 {
						v270 = int32(1)
						if v262&v270 == int32(0) {
							v326 = v270
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v277 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v278 = m.ExcPending
							if v278 != 0 {
								return int32(0)
							} else {
								if v277 == int32(0) {
									v326 = v270
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v281 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v281
									*(*float64)(unsafe.Add(mBase, uint32(v10)+32)) = v231
									*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v281), float64(2000)), float64(0.5))
									F_errmsg_internal(m, int32(_a_F_uuid_abbrev_abort_4), v10+int32(32))
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_uuid_abbrev_abort_2), int32(374), int32(_a_F_uuid_abbrev_abort_3))
										mBase = m.M
										v300 = m.ExcPending
										if v300 != 0 {
											return int32(0)
										} else {
											v326 = v270
											m.G0 = v10 + int32(96)
											return v326
										}
									}
								}
							}
						}
					} else {
						if v262&int32(1) == int32(0) {
							v326 = v3
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v307 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v308 = m.ExcPending
							if v308 != 0 {
								return int32(0)
							} else {
								if v307 == int32(0) {
									v326 = v3
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v311 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v311
									*(*float64)(unsafe.Add(mBase, uint32(v10)+64)) = v231
									F_errmsg_internal(m, int32(_a_F_uuid_abbrev_abort_5), v10-int32(-64))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_uuid_abbrev_abort_2), int32(381), int32(_a_F_uuid_abbrev_abort_3))
										mBase = m.M
										v324 = m.ExcPending
										if v324 != 0 {
											return int32(0)
										} else {
											v326 = v3
											m.G0 = v10 + int32(96)
											return v326
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
func F_uuid_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc(m, v10)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = int32(0)
	v26 = v13 + base.B2i32(v19 == int32(123))
	goto L5
L3:
	;
	m.G0 = v11 + int32(16)
	return v15
L4:
	;
	v114 = F_errsave_start(m, v23)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v33 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L6:
	;
	if v19 == int32(123) {
		goto L25
	} else {
		goto L26
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26))))
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v39)
	v42 = v39 & int32(255)
	goto L9
L9:
	;
	if base.B2i32(base.Ui32(v42-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v42|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v57 = int32(base.Ui32(v39) >> (uint(int32(8)) % 32))
	goto L11
L11:
	;
	if base.B2i32(base.Ui32(v57-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v57|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)) = uint8(v71)
	v79 = F_strtox_2(m, v11+int32(12), v71, int32(16), int64(4294967295))
	mBase = m.M
	v80 = base.I32_wrap_i64(v79)
	goto L13
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v25+v15))) = uint8(v80)
	v83 = v26 + int32(2)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != int32(45) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v97 = v25 + int32(1)
	if v97 != int32(16) {
		v25 = v97
		v26 = v95
		goto L5
	} else {
		goto L24
	}
L15:
	;
	v95 = v83
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v25&int32(1) != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v91 = v26 + int32(3)
	goto L20
L19:
	;
	v91 = v83
	goto L20
L20:
	;
	if v25 != int32(15) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v94 = v91
	goto L23
L22:
	;
	v94 = v83
	goto L23
L23:
	;
	v95 = v94
	goto L14
L24:
	;
	goto L6
L25:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v102 != int32(125) {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	v107 = v95
	goto L27
L27:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v108 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L28:
	;
	v107 = v95 + int32(1)
	goto L27
L29:
	;
	goto L4
L30:
	;
	if v114 == int32(0) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_uuid_in_0)
	F_errmsg(m, int32(_a_F_uuid_in_1), v11)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, v23, int32(_a_F_uuid_in_2), int32(183), int32(_a_F_uuid_in_3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L3
}
func F_uuid_increment(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = v12
		v15 = v6 + int32(15)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
		if v16 != int32(255) {
			v130 = v16
			v131 = v15
			v133 = v130 + int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
			v135 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
			return v6
		} else {
			v19 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v19)
			v22 = v6 + int32(14)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			if v23 != int32(255) {
				v130 = v23
				v131 = v22
				v133 = v130 + int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
				v135 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
				return v6
			} else {
				v26 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v26)
				v29 = v6 + int32(13)
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				if v30 != int32(255) {
					v130 = v30
					v131 = v29
					v133 = v130 + int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
					v135 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
					return v6
				} else {
					v33 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v33)
					v36 = v6 + int32(12)
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
					if v37 != int32(255) {
						v130 = v37
						v131 = v36
						v133 = v130 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
						v135 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
						return v6
					} else {
						v40 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
						v43 = v6 + int32(11)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
						if v44 != int32(255) {
							v130 = v44
							v131 = v43
							v133 = v130 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
							v135 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
							return v6
						} else {
							v47 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)) = uint8(v47)
							v50 = v6 + int32(10)
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
							if v51 != int32(255) {
								v130 = v51
								v131 = v50
								v133 = v130 + int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
								v135 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
								return v6
							} else {
								v54 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v54)
								v57 = v6 + int32(9)
								v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
								if v58 != int32(255) {
									v130 = v58
									v131 = v57
									v133 = v130 + int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
									v135 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
									return v6
								} else {
									v61 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v61)
									v64 = v6 + int32(8)
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
									if v65 != int32(255) {
										v130 = v65
										v131 = v64
										v133 = v130 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
										v135 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
										return v6
									} else {
										v68 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v6)+8)) = uint8(v68)
										v71 = v6 + int32(7)
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
										if v72 != int32(255) {
											v130 = v72
											v131 = v71
											v133 = v130 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
											v135 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
											return v6
										} else {
											v75 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)) = uint8(v75)
											v78 = v6 + int32(6)
											v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
											if v79 != int32(255) {
												v130 = v79
												v131 = v78
												v133 = v130 + int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
												v135 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
												return v6
											} else {
												v82 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v6)+6)) = uint8(v82)
												v85 = v6 + int32(5)
												v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
												if v86 != int32(255) {
													v130 = v86
													v131 = v85
													v133 = v130 + int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
													v135 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
													return v6
												} else {
													v89 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)) = uint8(v89)
													v92 = v6 + int32(4)
													v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
													if v93 != int32(255) {
														v130 = v93
														v131 = v92
														v133 = v130 + int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
														v135 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
														return v6
													} else {
														v96 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v96)
														v99 = v6 + int32(3)
														v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
														if v100 != int32(255) {
															v130 = v100
															v131 = v99
															v133 = v130 + int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
															v135 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
															return v6
														} else {
															v103 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)) = uint8(v103)
															v106 = v6 + int32(2)
															v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
															if v107 != int32(255) {
																v130 = v107
																v131 = v106
																v133 = v130 + int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																v135 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																return v6
															} else {
																v110 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)) = uint8(v110)
																v113 = v6 + int32(1)
																v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
																if v114 != int32(255) {
																	v130 = v114
																	v131 = v113
																	v133 = v130 + int32(1)
																	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																	v135 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																	return v6
																} else {
																	v117 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)) = uint8(v117)
																	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
																	if v119 != int32(255) {
																		v130 = v119
																		v131 = v6
																		v133 = v130 + int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v133)
																		v135 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v135)
																		return v6
																	} else {
																		v122 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v122)
																		F_pfree(m, v6)
																		mBase = m.M
																		v125 = m.ExcPending
																		if v125 != 0 {
																			return int32(0)
																		} else {
																			v126 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v126)
																			return int32(0)
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
		}
	}
}
func F_uuid_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(1543)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
	if v9 == int32(1) {
		v12 = int32(_a_F_uuid_sortsupport_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_uuid_sortsupport[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, _c_F_uuid_sortsupport[0])) = v15
		v18 = F_palloc(m, int32(40))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v22)
			*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
			F_initHyperLogLog(m, v18+int32(16), int32(10))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(1543)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(1544)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(1545)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(116)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = v18
				*(*int32)(unsafe.Add(mBase, _c_F_uuid_sortsupport[0])) = v13
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
