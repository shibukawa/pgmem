package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_numeric_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v23 int32
	_ = v23
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 float64
	_ = v48
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
	var v72 int32
	_ = v72
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v83 int32
	_ = v83
	var v84 float64
	_ = v84
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v194 float64
	_ = v194
	var v196 float64
	_ = v196
	var v198 float64
	_ = v198
	var v211 float64
	_ = v211
	var v221 float64
	_ = v221
	var v232 float64
	_ = v232
	var v244 float64
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int64
	_ = v294
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int64
	_ = v324
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l0 < int32(_a_F_numeric_abbrev_abort_0) {
		v340 = v3
		m.G0 = v10 + int32(96)
		return v340
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
		if v15 < int64(10000) {
			v340 = v3
			m.G0 = v10 + int32(96)
			return v340
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
			if v18 != int32(1) {
				v340 = v3
				m.G0 = v10 + int32(96)
				return v340
			} else {
				v22 = v14 + int32(24)
				v23 = int32(0)
				v30 = float64(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v32 != 0 {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					if v32 != int32(1) {
						v42 = v23
						v43 = v23
						v48 = v30
						for {
							v50 = float64(1)
							v51 = v42 + v33
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
							v53 = F_scalbn(m, v50, v52)
							mBase = m.M
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v57 = F_scalbn(m, v50, v56)
							mBase = m.M
							v62 = base.F64_add(base.F64_add(v48, base.F64_div(v50, v57)), base.F64_div(v50, v53))
							v63 = int32(2)
							v64 = v42 + v63
							v66 = v43 + v63
							if v66 != v32&int32(-2) {
								v42 = v64
								v43 = v66
								v48 = v62
								continue
							} else {
								break
							}
							break
						}
						if v32&int32(1) == int32(0) {
							v95 = v62
						} else {
							v72 = v64
							v78 = v62
							v80 = float64(1)
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v33))))
							v84 = F_scalbn(m, v80, v83)
							mBase = m.M
							v95 = base.F64_add(v78, base.F64_div(v80, v84))
						}
					} else {
						v72 = v23
						v78 = v30
						v80 = float64(1)
						v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72+v33))))
						v84 = F_scalbn(m, v80, v83)
						mBase = m.M
						v95 = base.F64_add(v78, base.F64_div(v80, v84))
					}
					v97 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v98 = base.F64_div(v97, v95)
					v99 = base.F64_convert_i32_u(v32)
					if base.F64_le(v98, base.F64_mul(v99, float64(2.5))) == int32(0) {
						v211 = v98
						if base.F64_gt(v211, float64(1.4316557653333333e+08)) == int32(0) {
							v232 = v211
						} else {
							v221 = F_log(m, base.F64_add(base.F64_mul(v211, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v232 = base.F64_mul(v221, float64(-4.294967296e+09))
						}
						v244 = v232
					} else {
						v106 = v32 & int32(3)
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						v108 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v117 = int32(0)
							v118 = v108
							v119 = v108
							for {
								v126 = v118 + v107
								v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
								v128 = int32(0)
								v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+1)))
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+2)))
								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+3)))
								v142 = v119 + base.B2i32(v127 == v128) + base.B2i32(v131 == v128) + base.B2i32(v135 == v128) + base.B2i32(v139 == v128)
								v143 = int32(4)
								v144 = v118 + v143
								v146 = v117 + v143
								if v146 != v32&int32(-4) {
									v117 = v146
									v118 = v144
									v119 = v142
									continue
								} else {
									break
								}
								break
							}
							if v106 == int32(0) {
								v183 = v142
							} else {
								v152 = v144
								v153 = v142
								v162 = v152
								v163 = v153
								v166 = v108
								for {
									v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v107))))
									v174 = v163 + base.B2i32(v171 == int32(0))
									v175 = int32(1)
									v178 = v166 + v175
									if v178 != v106 {
										v162 = v162 + v175
										v163 = v174
										v166 = v178
										continue
									} else {
										break
									}
									break
								}
								v183 = v174
							}
						} else {
							v152 = v108
							v153 = v108
							v162 = v152
							v163 = v153
							v166 = v108
							for {
								v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v107))))
								v174 = v163 + base.B2i32(v171 == int32(0))
								v175 = int32(1)
								v178 = v166 + v175
								if v178 != v106 {
									v162 = v162 + v175
									v163 = v174
									v166 = v178
									continue
								} else {
									break
								}
								break
							}
							v183 = v174
						}
						if v183 == int32(0) {
							v232 = v98
							v244 = v232
						} else {
							v194 = F_log(m, base.F64_div(v99, base.F64_convert_i32_s(v183)))
							mBase = m.M
							v244 = base.F64_mul(v194, v99)
						}
					}
				} else {
					v196 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v198 = base.F64_div(v196, float64(0))
					if base.F64_le(v198, base.F64_mul(base.F64_convert_i32_u(v32), float64(2.5))) != 0 {
						v232 = v198
					} else {
						v211 = v198
						if base.F64_gt(v211, float64(1.4316557653333333e+08)) == int32(0) {
							v232 = v211
						} else {
							v221 = F_log(m, base.F64_add(base.F64_mul(v211, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v232 = base.F64_mul(v221, float64(-4.294967296e+09))
						}
					}
					v244 = v232
				}
				if base.F64_gt(v244, float64(100000)) != 0 {
					v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_abbrev_abort[0])))
					if v248 != int32(1) {
						v272 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v272)
						v340 = v3
						m.G0 = v10 + int32(96)
						return v340
					} else {
						v253 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v256 = m.ExcPending
						if v256 != 0 {
							return int32(0)
						} else {
							if v253 == int32(0) {
								v272 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v272)
								v340 = v3
								m.G0 = v10 + int32(96)
								return v340
							} else {
								v259 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v259
								*(*float64)(unsafe.Add(mBase, uint32(v10))) = v244
								F_errmsg_internal(m, int32(_a_F_numeric_abbrev_abort_1), v10)
								mBase = m.M
								v265 = m.ExcPending
								if v265 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2241), int32(_a_F_numeric_abbrev_abort_3))
									mBase = m.M
									v270 = m.ExcPending
									if v270 != 0 {
										return int32(0)
									} else {
										v272 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v272)
										v340 = v3
										m.G0 = v10 + int32(96)
										return v340
									}
								}
							}
						}
					}
				} else {
					v275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_abbrev_abort[0])))
					v276 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v276), float64(10000)), float64(0.5)), v244) != 0 {
						v283 = int32(1)
						if v275&v283 == int32(0) {
							v340 = v283
							m.G0 = v10 + int32(96)
							return v340
						} else {
							v290 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v291 = m.ExcPending
							if v291 != 0 {
								return int32(0)
							} else {
								if v290 == int32(0) {
									v340 = v283
									m.G0 = v10 + int32(96)
									return v340
								} else {
									v294 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v294
									*(*float64)(unsafe.Add(mBase, uint32(v10)+32)) = v244
									*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v294), float64(10000)), float64(0.5))
									F_errmsg_internal(m, int32(_a_F_numeric_abbrev_abort_4), v10+int32(32))
									mBase = m.M
									v308 = m.ExcPending
									if v308 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2262), int32(_a_F_numeric_abbrev_abort_3))
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v340 = v283
											m.G0 = v10 + int32(96)
											return v340
										}
									}
								}
							}
						}
					} else {
						if v275&int32(1) == int32(0) {
							v340 = v3
							m.G0 = v10 + int32(96)
							return v340
						} else {
							v320 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v321 = m.ExcPending
							if v321 != 0 {
								return int32(0)
							} else {
								if v320 == int32(0) {
									v340 = v3
									m.G0 = v10 + int32(96)
									return v340
								} else {
									v324 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v324
									*(*float64)(unsafe.Add(mBase, uint32(v10)+64)) = v244
									F_errmsg_internal(m, int32(_a_F_numeric_abbrev_abort_5), v10-int32(-64))
									mBase = m.M
									v332 = m.ExcPending
									if v332 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2270), int32(_a_F_numeric_abbrev_abort_3))
										mBase = m.M
										v337 = m.ExcPending
										if v337 != 0 {
											return int32(0)
										} else {
											v340 = v3
											m.G0 = v10 + int32(96)
											return v340
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
func F_numeric_accum(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14334(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_numeric_add(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v11 = F_numeric_add_safe(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
func F_numeric_ceil(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v100 int64
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v343 int64
	_ = v343
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = int32(base.Ui32(v19) >> (uint(int32(2)) % 32))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	if base.Ui32(int32(_a_F_numeric_ceil_0)) <= base.Ui32(v22) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v12 + int32(48)
	return base.I64_extend_i32_u(v357)
L4:
	;
	v25 = F_palloc(m, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v35 = base.I32_extend16_s(v22)
	v37 = base.B2i32(int32(0) <= v35)
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
	if v29 == int32(0) {
		v357 = v25
		goto L3
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, v25, v15, v29)
	v357 = v25
	goto L3
L9:
	;
	v38 = int32(-8)
	goto L11
L10:
	;
	v38 = int32(-6)
	goto L11
L11:
	;
	v39 = v21 + v38
	v41 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v41
	if int32(0) <= v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v15)+6)))
	v53 = v43
	goto L14
L13:
	;
	v53 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v53
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v55
	v60 = base.B2i32(v35 < v55)
	if v35 < v55 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = int32(6)
	goto L17
L16:
	;
	v61 = int32(8)
	goto L17
L17:
	;
	v62 = v15 + v61
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v62
	if v35 < v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(base.Ui32(v22)>>(uint(int32(7))%32)) & int32(63)
	goto L20
L19:
	;
	v70 = v22 & int32(_a_F_numeric_ceil_1)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v70
	v77 = v22 & int32(_a_F_numeric_ceil_0)
	if v77 == int32(_a_F_numeric_ceil_2) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v80 = v22 << (uint(int32(1)) % 32) & int32(_a_F_numeric_ceil_3)
	goto L23
L22:
	;
	v80 = v77
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v80
	v83 = v39 & int32(-2)
	v86 = F_palloc(m, v83+int32(2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v88)
	if base.B2i32(v41 == v88)|base.B2i32(v83 == v88) == v88 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	base.MemoryCopy(m, v86+int32(2), v62, v83)
	goto L27
L26:
	;
	goto L27
L27:
	;
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v100
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v86
	v107 = int32(2)
	v108 = v86 + v107
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v112 = v110 << (uint(v107) % 32)
	if base.Ui32(int32(2147483644)) <= base.Ui32(v112) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v129
	if v80 != 0 {
		v321 = v129
		goto L35
	} else {
		goto L36
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+28)) = int64(0)
	v117 = int32(0)
	v127 = v117
	v129 = v117
	goto L28
L30:
	;
	goto L31
L31:
	;
	v122 = base.I32_div_s(v112+int32(7), int32(4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	if v122 < v123 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v125 = v122
	goto L34
L33:
	;
	v125 = v123
	goto L34
L34:
	;
	v127 = v110
	v129 = v125
	goto L28
L35:
	;
	v323 = v321 << (uint(int32(1)) % 32)
	v326 = F_palloc(m, v323+int32(2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L89
	}
L36:
	;
	if v41 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v316 = v12 + int32(24)
	F_add_var(m, v316, int32(_a_F_numeric_ceil_4), v316)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L88
	}
L38:
	;
	if v129 != 0 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v129 == int32(0) {
		goto L37
	} else {
		goto L42
	}
L41:
	;
	v321 = int32(0)
	goto L35
L42:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v12)+32))
	if v136 == int32(_a_F_numeric_ceil_3) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v139 = int32(0)
	if base.B2i32(v127 < v53)&base.B2i32(v139 < v41) == v139 {
		v170 = v53
		v174 = v139
		goto L46
	} else {
		goto L47
	}
L44:
	;
	if v312 == int32(0) {
		v321 = v129
		goto L35
	} else {
		goto L87
	}
L45:
	;
	v312 = v302
	goto L44
L46:
	;
	if base.B2i32(v129 <= int32(0))|base.B2i32(v127 <= v170) != 0 {
		v206 = v127
		v208 = v139
		goto L53
	} else {
		goto L54
	}
L47:
	;
	v151 = v53
	v155 = v139
	goto L48
L48:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v155<<(uint(int32(1))%32)))))
	if v161 != 0 {
		v302 = int32(1)
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v170 = v165
	v174 = v163
	goto L46
L50:
	;
	v162 = int32(1)
	v163 = v155 + v162
	v165 = v151 - v162
	if v165 <= v127 {
		v170 = v165
		v174 = v163
		goto L46
	} else {
		goto L51
	}
L51:
	;
	if v163 < v41 {
		v151 = v165
		v155 = v163
		goto L48
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	if v170 != v206 {
		v248 = v174
		v249 = v208
		goto L60
	} else {
		goto L61
	}
L54:
	;
	v187 = v127
	v189 = v139
	goto L55
L55:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+v189<<(uint(int32(1))%32)))))
	if v194 != 0 {
		v302 = int32(-1)
		goto L45
	} else {
		goto L57
	}
L56:
	;
	v206 = v198
	v208 = v196
	goto L53
L57:
	;
	v195 = int32(1)
	v196 = v189 + v195
	v198 = v187 - v195
	if v198 <= v170 {
		v206 = v198
		v208 = v196
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if v196 < v129 {
		v187 = v198
		v189 = v196
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	if v41 < v248 {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v217 = v174
	v218 = v208
	goto L62
L62:
	;
	if base.B2i32(v41 <= v217)|base.B2i32(v129 <= v218) != 0 {
		v248 = v217
		v249 = v218
		goto L60
	} else {
		goto L64
	}
L63:
	;
	if base.I32_extend16_s(v234) < base.I32_extend16_s(v232) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	v223 = int32(1)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v217<<(uint(v223)%32)))))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218<<(uint(v223)%32)+v108))))
	if v232 == v234 {
		v217 = v217 + v223
		v218 = v218 + v223
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v241 = int32(1)
	goto L68
L67:
	;
	v241 = int32(-1)
	goto L68
L68:
	;
	v312 = v241
	goto L44
L69:
	;
	v252 = v248
	goto L71
L70:
	;
	v252 = v41
	goto L71
L71:
	;
	v259 = v248
	goto L72
L72:
	;
	if v252 == v259 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v302 = v285
	goto L45
L74:
	;
	if v129 < v249 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	v285 = int32(1)
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v259<<(uint(v285)%32)))))
	if v291 == int32(0) {
		v259 = v259 + v285
		goto L72
	} else {
		goto L86
	}
L77:
	;
	v264 = v249
	goto L79
L78:
	;
	v264 = v129
	goto L79
L79:
	;
	v272 = v249
	goto L80
L80:
	;
	if v264 == v272 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v302 = int32(-1)
	goto L45
L82:
	;
	v312 = int32(0)
	goto L44
L83:
	;
	goto L84
L84:
	;
	v276 = int32(1)
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272<<(uint(v276)%32)+v108))))
	if v281 == int32(0) {
		v272 = v272 + v276
		goto L80
	} else {
		goto L85
	}
L85:
	;
	goto L81
L86:
	;
	goto L73
L87:
	;
	goto L37
L88:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v321 = v320
	goto L35
L89:
	;
	v328 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v326))) = uint16(v328)
	if base.B2i32(v323 == v328)|base.B2i32(v321 <= v328) == v328 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	base.MemoryCopy(m, v326+int32(2), v339, v323)
	goto L92
L91:
	;
	goto L92
L92:
	;
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v341
	v343 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v343
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v326 + int32(2)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v349 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_pfree(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v353 = F_make_result_safe(m, v12, int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L97
	}
L96:
	;
	goto L95
L97:
	;
	F_pfree(m, v326)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v357 = v353
	goto L3
}
func F_numeric_cmp_abbrev(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	return base.B2i32(l0 < l1) - base.B2i32(l1 < l0)
}
func F_numeric_combine(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
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
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = v8 + int32(4)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == v2 {
		v30 = int32(0)
		if v11 == v30 {
			v38 = v30
		} else {
			v33 = v30
			v34 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
			v38 = v34
		}
		v41 = v38
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(435) {
		case 0:
			if v11 == int32(0) {
				v41 = int32(1)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
				v33 = v23
				v34 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		case 1:
			if v11 == int32(0) {
				v41 = int32(2)
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+376))
				v33 = v28
				v34 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
				v41 = v38
			}
		default:
			v30 = int32(0)
			if v11 == v30 {
				v38 = v30
			} else {
				v33 = v30
				v34 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v33
				v38 = v34
			}
			v41 = v38
		}
	}
	if v41 != 0 {
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v42 == int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v46 = v45
		} else {
			v46 = v2
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v47 == int32(0) {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v50 != 0 {
				if v46 == int32(0) {
					v57 = int32(_a_F_numeric_combine_0)
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v60
					v63 = F_palloc0(m, int32(112))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						v67 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v67)
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+4)) = v70
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v63)+8)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v50)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v63)+88)) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v50)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v63)+96)) = v76
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v50)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v63)+104)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+72)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v63)+80)) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						v87 = F_palloc(m, v84<<(uint(int32(2))%32))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v63)+36)) = v87
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
							v93 = F_palloc(m, v90<<(uint(int32(2))%32))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v63)+40)) = v93
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								v98 = v96 << (uint(int32(2)) % 32)
								if v98 != 0 {
									v99 = *(*int32)(unsafe.Add(mBase, uint32(v63)+36))
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
									base.MemoryCopy(m, v99, v100, v98)
								} else {
								}
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								v104 = v102 << (uint(int32(2)) % 32)
								if v104 != 0 {
									v105 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
									v106 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
									base.MemoryCopy(m, v105, v106, v104)
								} else {
								}
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v63)+28)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = v110
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = v112
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = v114
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
								v119 = F_palloc(m, v116<<(uint(int32(2))%32))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v63)+64)) = v119
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
									v125 = F_palloc(m, v122<<(uint(int32(2))%32))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v63)+68)) = v125
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										v130 = v128 << (uint(int32(2)) % 32)
										if v130 != 0 {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v63)+64))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v50)+64))
											base.MemoryCopy(m, v131, v132, v130)
										} else {
										}
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										v136 = v134 << (uint(int32(2)) % 32)
										if v136 != 0 {
											v137 = *(*int32)(unsafe.Add(mBase, uint32(v63)+68))
											v138 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
											base.MemoryCopy(m, v137, v138, v136)
										} else {
										}
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v63)+56)) = v140
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v63)+44)) = v142
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v63)+48)) = v144
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
										*(*int32)(unsafe.Add(mBase, uint32(v63)+52)) = v146
										*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v58
										v226 = v63
										m.G0 = v8 + int32(32)
										return base.I64_extend_i32_u(v226)
									}
								}
							}
						}
					}
				} else {
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v150 + v151
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v46)+88))
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v50)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+88)) = v154 + v155
					v158 = *(*int64)(unsafe.Add(mBase, uint32(v46)+96))
					v159 = *(*int64)(unsafe.Add(mBase, uint32(v50)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+96)) = v158 + v159
					v162 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
					v163 = *(*int64)(unsafe.Add(mBase, uint32(v50)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+104)) = v162 + v163
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
					if v166 <= int64(0) {
						v226 = v46
						m.G0 = v8 + int32(32)
						return base.I64_extend_i32_u(v226)
					} else {
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
						if v170 < v169 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+72)) = v169
							v173 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v173
						} else {
							if v169 != v170 {
							} else {
								v176 = *(*int64)(unsafe.Add(mBase, uint32(v46)+80))
								v177 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v176 + v177
							}
						}
						v180 = int32(_a_F_numeric_combine_0)
						v181 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v183
						v185 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v185
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v185
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v185
						v194 = v8 + int32(8)
						F_accum_sum_final(m, v50+int32(16), v194)
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return int64(0)
						} else {
							F_accum_sum_add(m, v46+int32(16), v194)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return int64(0)
							} else {
								v201 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								if v201 != 0 {
									F_pfree(m, v201)
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int64(0)
									} else {
										v204 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v204
										*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v204
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v204
										v213 = v8 + int32(8)
										F_accum_sum_final(m, v50+int32(44), v213)
										mBase = m.M
										v215 = m.ExcPending
										if v215 != 0 {
											return int64(0)
										} else {
											F_accum_sum_add(m, v46+int32(44), v213)
											mBase = m.M
											v219 = m.ExcPending
											if v219 != 0 {
												return int64(0)
											} else {
												v220 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
												if v220 != 0 {
													F_pfree(m, v220)
													mBase = m.M
													v222 = m.ExcPending
													if v222 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v181
														v226 = v46
														m.G0 = v8 + int32(32)
														return base.I64_extend_i32_u(v226)
													}
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v181
													v226 = v46
													m.G0 = v8 + int32(32)
													return base.I64_extend_i32_u(v226)
												}
											}
										}
									}
								} else {
									v204 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v204
									*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v204
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v204
									v213 = v8 + int32(8)
									F_accum_sum_final(m, v50+int32(44), v213)
									mBase = m.M
									v215 = m.ExcPending
									if v215 != 0 {
										return int64(0)
									} else {
										F_accum_sum_add(m, v46+int32(44), v213)
										mBase = m.M
										v219 = m.ExcPending
										if v219 != 0 {
											return int64(0)
										} else {
											v220 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
											if v220 != 0 {
												F_pfree(m, v220)
												mBase = m.M
												v222 = m.ExcPending
												if v222 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v181
													v226 = v46
													m.G0 = v8 + int32(32)
													return base.I64_extend_i32_u(v226)
												}
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v181
												v226 = v46
												m.G0 = v8 + int32(32)
												return base.I64_extend_i32_u(v226)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v46 != 0 {
					v226 = v46
				} else {
					v52 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
					v226 = int32(0)
				}
				m.G0 = v8 + int32(32)
				return base.I64_extend_i32_u(v226)
			}
		} else {
			if v46 != 0 {
				v226 = v46
			} else {
				v52 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v52)
				v226 = int32(0)
			}
			m.G0 = v8 + int32(32)
			return base.I64_extend_i32_u(v226)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v237 = m.ExcPending
		if v237 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_combine_1), int32(0))
			mBase = m.M
			v241 = m.ExcPending
			if v241 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_combine_2), int32(_a_F_numeric_combine_3), int32(_a_F_numeric_combine_4))
				mBase = m.M
				v246 = m.ExcPending
				if v246 != 0 {
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
func F_numeric_div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v11 = F_numeric_div_safe(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
func F_numeric_exp(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v101 float64
	_ = v101
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(_a_F_numeric_exp_0)) <= base.Ui32(v16) {
			if v17 == int32(-4096) {
				v24 = F_make_result_safe(m, int32(_a_F_numeric_exp_1), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v127 = v24
					m.G0 = v9 + int32(48)
					return base.I64_extend_i32_u(v127)
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v29 = F_palloc(m, int32(base.Ui32(v26)>>(uint(int32(2))%32)))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v33 = int32(base.Ui32(v31) >> (uint(int32(2)) % 32))
					if v33 == int32(0) {
						v127 = v29
					} else {
						base.MemoryCopy(m, v29, v12, v33)
						v127 = v29
					}
					m.G0 = v9 + int32(48)
					return base.I64_extend_i32_u(v127)
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v43 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v44 = int32(-8)
			} else {
				v44 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(int32(base.Ui32(v37)>>(uint(int32(2))%32))+v44) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v59 = v49
			} else {
				v59 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v59
			v61 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v61
			v63 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v63
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v63
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v63
			v72 = base.B2i32(v17 < v61)
			if v17 < v61 {
				v73 = int32(6)
			} else {
				v73 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v12 + v73
			v81 = v16 & int32(_a_F_numeric_exp_0)
			if v81 == int32(_a_F_numeric_exp_2) {
				v84 = v16 << (uint(int32(1)) % 32) & int32(_a_F_numeric_exp_3)
			} else {
				v84 = v81
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v84
			if v17 < v61 {
				v92 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v92 = v16 & int32(_a_F_numeric_exp_4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v92
			v95 = v9 + int32(24)
			v98 = F_numericvar_to_double_no_overflow(m, v95)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return int64(0)
			} else {
				v101 = base.F64_mul(v98, float64(0.434294481903252))
				v102 = float64(-2000)
				if base.F64_gt(v101, v102) != 0 {
					v105 = v101
				} else {
					v105 = v102
				}
				v106 = float64(2000)
				if base.F64_lt(v105, v106) != 0 {
					v109 = v105
				} else {
					v109 = v106
				}
				v111 = int32(16) - base.I32_trunc_sat_f64_s(v109)
				if v92 < v111 {
					v113 = v111
				} else {
					v113 = v92
				}
				if int32(1000) <= v113 {
					v116 = int32(1000)
				} else {
					v116 = v113
				}
				F_exp_var(m, v95, v9, v116)
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int64(0)
				} else {
					v120 = F_make_result_safe(m, v9, int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int64(0)
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
						if v122 == int32(0) {
							v127 = v120
							m.G0 = v9 + int32(48)
							return base.I64_extend_i32_u(v127)
						} else {
							F_pfree(m, v122)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								v127 = v120
								m.G0 = v9 + int32(48)
								return base.I64_extend_i32_u(v127)
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_fast_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	v8 = base.I32_wrap_i64(l0)
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = base.I32_wrap_i64(l1)
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v28 = base.I32_extend16_s(v27)
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			if base.Ui32(int32(_a_F_numeric_fast_cmp_0)) <= base.Ui32(v29) {
				if v29 != int32(_a_F_numeric_fast_cmp_1) {
					if v29 != int32(_a_F_numeric_fast_cmp_0) {
						if v28 != int32(-4096) {
							v48 = int32(-1)
						} else {
							v48 = int32(0)
						}
						v169 = v48
					} else {
						v169 = base.B2i32(v28 != int32(-16384))
					}
				} else {
					if v28 == int32(-16384) {
						v43 = int32(-1)
					} else {
						v43 = base.B2i32(v28 != int32(-12288))
					}
					v169 = v43
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v28) {
					if v28 == int32(-4096) {
						v55 = int32(1)
					} else {
						v55 = int32(-1)
					}
					v169 = v55
				} else {
					v57 = v9 + int32(6)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v63 = base.I32_extend16_s(v29)
					v65 = base.B2i32(int32(0) <= v63)
					if int32(0) <= v63 {
						v66 = int32(-8)
					} else {
						v66 = int32(-6)
					}
					if int32(0) <= v63 {
						v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57))))
						v78 = v68
					} else {
						v78 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = int32(base.Ui32(int32(base.Ui32(v58)>>(uint(int32(2))%32))+v66) >> (uint(int32(1)) % 32))
					v82 = v14 + int32(6)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v89 = base.B2i32(int32(0) <= v28)
					if int32(0) <= v28 {
						v90 = int32(-8)
					} else {
						v90 = int32(-6)
					}
					if int32(0) <= v28 {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v27<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v27&int32(63)
					}
					v103 = int32(1)
					v104 = int32(base.Ui32(int32(base.Ui32(v83)>>(uint(int32(2))%32))+v90) >> (uint(v103) % 32))
					v110 = v27 & int32(_a_F_numeric_fast_cmp_0)
					if v110 == int32(_a_F_numeric_fast_cmp_2) {
						v113 = v27 << (uint(v103) % 32) & int32(_a_F_numeric_fast_cmp_3)
					} else {
						v113 = v110
					}
					if v80 == int32(0) {
						if v104 == int32(0) {
							v169 = int32(0)
						} else {
							if v113 == int32(_a_F_numeric_fast_cmp_3) {
								v123 = int32(1)
							} else {
								v123 = int32(-1)
							}
							v169 = v123
						}
					} else {
						v129 = v29 & int32(_a_F_numeric_fast_cmp_0)
						if v129 == int32(_a_F_numeric_fast_cmp_2) {
							v132 = v29 << (uint(int32(1)) % 32) & int32(_a_F_numeric_fast_cmp_3)
						} else {
							v132 = v129
						}
						if v104 == int32(0) {
							if v132 != 0 {
								v137 = int32(-1)
							} else {
								v137 = int32(1)
							}
							v169 = v137
						} else {
							if v63 < int32(0) {
								v142 = v57
							} else {
								v142 = v9 + int32(8)
							}
							if v28 < int32(0) {
								v147 = v82
							} else {
								v147 = v14 + int32(8)
							}
							if v132 == int32(0) {
								if v113 == int32(_a_F_numeric_fast_cmp_3) {
									v169 = int32(1)
								} else {
									v153 = F_cmp_abs_common(m, v142, v80, v78, v147, v104, v102)
									mBase = m.M
									v169 = v153
								}
							} else {
								if v113 == int32(0) {
									v169 = int32(-1)
								} else {
									v157 = F_cmp_abs_common(m, v147, v104, v102, v142, v80, v78)
									mBase = m.M
									v169 = v157
								}
							}
						}
					}
				}
			}
			if v9 != v8 {
				F_pfree(m, v9)
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
					return int32(0)
				} else {
					if v14 != v13 {
						F_pfree(m, v14)
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							return v169
						}
					} else {
						return v169
					}
				}
			} else {
				if v14 != v13 {
					F_pfree(m, v14)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return int32(0)
					} else {
						return v169
					}
				} else {
					return v169
				}
			}
		}
	}
}
func F_numeric_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = base.I32_extend16_s(v24)
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			if base.Ui32(int32(_a_F_numeric_gt_0)) <= base.Ui32(v26) {
				if v26 != int32(_a_F_numeric_gt_1) {
					if v26 != int32(_a_F_numeric_gt_0) {
						if v25 != int32(-4096) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v166 = v45
					} else {
						v166 = base.B2i32(v25 != int32(-16384))
					}
				} else {
					if v25 == int32(-16384) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v25 != int32(-12288))
					}
					v166 = v40
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v25) {
					if v25 == int32(-4096) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v166 = v52
				} else {
					v54 = v6 + int32(6)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v60 = base.I32_extend16_s(v26)
					v62 = base.B2i32(int32(0) <= v60)
					if int32(0) <= v60 {
						v63 = int32(-8)
					} else {
						v63 = int32(-6)
					}
					if int32(0) <= v60 {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v75 = v65
					} else {
						v75 = v26<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v26&int32(63)
					}
					v77 = int32(base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))+v63) >> (uint(int32(1)) % 32))
					v79 = v11 + int32(6)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v86 = base.B2i32(int32(0) <= v25)
					if int32(0) <= v25 {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= v25 {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79))))
						v99 = v89
					} else {
						v99 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v24 & int32(_a_F_numeric_gt_0)
					if v107 == int32(_a_F_numeric_gt_2) {
						v110 = v24 << (uint(v100) % 32) & int32(_a_F_numeric_gt_3)
					} else {
						v110 = v107
					}
					if v77 == int32(0) {
						if v101 == int32(0) {
							v166 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_gt_3) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v166 = v120
						}
					} else {
						v126 = v26 & int32(_a_F_numeric_gt_0)
						if v126 == int32(_a_F_numeric_gt_2) {
							v129 = v26 << (uint(int32(1)) % 32) & int32(_a_F_numeric_gt_3)
						} else {
							v129 = v126
						}
						if v101 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v166 = v134
						} else {
							if v60 < int32(0) {
								v139 = v54
							} else {
								v139 = v6 + int32(8)
							}
							if v25 < int32(0) {
								v144 = v79
							} else {
								v144 = v11 + int32(8)
							}
							if v129 == int32(0) {
								if v110 == int32(_a_F_numeric_gt_3) {
									v166 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v139, v77, v75, v144, v101, v99)
									mBase = m.M
									v166 = v150
								}
							} else {
								if v110 == int32(0) {
									v166 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v144, v101, v99, v139, v77, v75)
									mBase = m.M
									v166 = v154
								}
							}
						}
					}
				}
			}
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v167 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int64(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v171 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) < v166))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v166))
					}
				}
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v171 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) < v166))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(int32(0) < v166))
				}
			}
		}
	}
}
func F_numeric_le(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = base.I32_extend16_s(v24)
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			if base.Ui32(int32(_a_F_numeric_le_0)) <= base.Ui32(v26) {
				if v26 != int32(_a_F_numeric_le_1) {
					if v26 != int32(_a_F_numeric_le_0) {
						if v25 != int32(-4096) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v166 = v45
					} else {
						v166 = base.B2i32(v25 != int32(-16384))
					}
				} else {
					if v25 == int32(-16384) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v25 != int32(-12288))
					}
					v166 = v40
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v25) {
					if v25 == int32(-4096) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v166 = v52
				} else {
					v54 = v6 + int32(6)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v60 = base.I32_extend16_s(v26)
					v62 = base.B2i32(int32(0) <= v60)
					if int32(0) <= v60 {
						v63 = int32(-8)
					} else {
						v63 = int32(-6)
					}
					if int32(0) <= v60 {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v75 = v65
					} else {
						v75 = v26<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v26&int32(63)
					}
					v77 = int32(base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))+v63) >> (uint(int32(1)) % 32))
					v79 = v11 + int32(6)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v86 = base.B2i32(int32(0) <= v25)
					if int32(0) <= v25 {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= v25 {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79))))
						v99 = v89
					} else {
						v99 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v24 & int32(_a_F_numeric_le_0)
					if v107 == int32(_a_F_numeric_le_2) {
						v110 = v24 << (uint(v100) % 32) & int32(_a_F_numeric_le_3)
					} else {
						v110 = v107
					}
					if v77 == int32(0) {
						if v101 == int32(0) {
							v166 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_le_3) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v166 = v120
						}
					} else {
						v126 = v26 & int32(_a_F_numeric_le_0)
						if v126 == int32(_a_F_numeric_le_2) {
							v129 = v26 << (uint(int32(1)) % 32) & int32(_a_F_numeric_le_3)
						} else {
							v129 = v126
						}
						if v101 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v166 = v134
						} else {
							if v60 < int32(0) {
								v139 = v54
							} else {
								v139 = v6 + int32(8)
							}
							if v25 < int32(0) {
								v144 = v79
							} else {
								v144 = v11 + int32(8)
							}
							if v129 == int32(0) {
								if v110 == int32(_a_F_numeric_le_3) {
									v166 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v139, v77, v75, v144, v101, v99)
									mBase = m.M
									v166 = v150
								}
							} else {
								if v110 == int32(0) {
									v166 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v144, v101, v99, v139, v77, v75)
									mBase = m.M
									v166 = v154
								}
							}
						}
					}
				}
			}
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v167 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int64(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v171 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v166 <= int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v166 <= int32(0)))
					}
				}
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v171 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v166 <= int32(0)))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(v166 <= int32(0)))
				}
			}
		}
	}
}
func F_numeric_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = base.I32_extend16_s(v24)
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			if base.Ui32(int32(_a_F_numeric_ne_0)) <= base.Ui32(v26) {
				if v26 != int32(_a_F_numeric_ne_1) {
					if v26 != int32(_a_F_numeric_ne_0) {
						if v25 != int32(-4096) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v166 = v45
					} else {
						v166 = base.B2i32(v25 != int32(-16384))
					}
				} else {
					if v25 == int32(-16384) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v25 != int32(-12288))
					}
					v166 = v40
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v25) {
					if v25 == int32(-4096) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v166 = v52
				} else {
					v54 = v6 + int32(6)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v60 = base.I32_extend16_s(v26)
					v62 = base.B2i32(int32(0) <= v60)
					if int32(0) <= v60 {
						v63 = int32(-8)
					} else {
						v63 = int32(-6)
					}
					if int32(0) <= v60 {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v75 = v65
					} else {
						v75 = v26<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v26&int32(63)
					}
					v77 = int32(base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))+v63) >> (uint(int32(1)) % 32))
					v79 = v11 + int32(6)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v86 = base.B2i32(int32(0) <= v25)
					if int32(0) <= v25 {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= v25 {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79))))
						v99 = v89
					} else {
						v99 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v24 & int32(_a_F_numeric_ne_0)
					if v107 == int32(_a_F_numeric_ne_2) {
						v110 = v24 << (uint(v100) % 32) & int32(_a_F_numeric_ne_3)
					} else {
						v110 = v107
					}
					if v77 == int32(0) {
						if v101 == int32(0) {
							v166 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_ne_3) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v166 = v120
						}
					} else {
						v126 = v26 & int32(_a_F_numeric_ne_0)
						if v126 == int32(_a_F_numeric_ne_2) {
							v129 = v26 << (uint(int32(1)) % 32) & int32(_a_F_numeric_ne_3)
						} else {
							v129 = v126
						}
						if v101 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v166 = v134
						} else {
							if v60 < int32(0) {
								v139 = v54
							} else {
								v139 = v6 + int32(8)
							}
							if v25 < int32(0) {
								v144 = v79
							} else {
								v144 = v11 + int32(8)
							}
							if v129 == int32(0) {
								if v110 == int32(_a_F_numeric_ne_3) {
									v166 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v139, v77, v75, v144, v101, v99)
									mBase = m.M
									v166 = v150
								}
							} else {
								if v110 == int32(0) {
									v166 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v144, v101, v99, v139, v77, v75)
									mBase = m.M
									v166 = v154
								}
							}
						}
					}
				}
			}
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v167 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int64(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v171 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v166 != int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v166 != int32(0)))
					}
				}
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v171 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v166 != int32(0)))
					}
				} else {
					return base.I64_extend_i32_u(base.B2i32(v166 != int32(0)))
				}
			}
		}
	}
}
func F_numeric_sign(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
		if v9 == int32(_a_F_numeric_sign_0) {
			v45 = int32(_a_F_numeric_sign_1)
		} else {
			if base.Ui32(int32(_a_F_numeric_sign_0)) <= base.Ui32(v9) {
				if v9 != int32(_a_F_numeric_sign_2) {
					v45 = int32(_a_F_numeric_sign_3)
				} else {
					v45 = int32(_a_F_numeric_sign_4)
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				if int32(0) <= base.I32_extend16_s(v9) {
					v26 = int32(-8)
				} else {
					v26 = int32(-6)
				}
				if base.Ui32(int32(base.Ui32(v18)>>(uint(int32(2))%32))+v26) < base.Ui32(int32(2)) {
					v45 = int32(_a_F_numeric_sign_5)
				} else {
					v35 = v9 & int32(_a_F_numeric_sign_0)
					if v35 == int32(_a_F_numeric_sign_6) {
						v38 = v9 << (uint(int32(1)) % 32) & int32(_a_F_numeric_sign_7)
					} else {
						v38 = v35
					}
					if v38 == int32(_a_F_numeric_sign_7) {
						v45 = int32(_a_F_numeric_sign_3)
					} else {
						v45 = int32(_a_F_numeric_sign_4)
					}
				}
			}
		}
		v47 = F_make_result_safe(m, v45, int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v47)
		}
	}
}
func F_numeric_to_number(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v163 int64
	_ = v163
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v20 == int32(1) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v49 = v36
			} else {
				v37 = int32(1)
				if v20&v37 != 0 {
					v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if base.Ui32(v49-int32(268435455)) <= base.Ui32(int32(-268435455)) {
				v54 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
				v163 = int64(0)
				m.G0 = v10 + int32(48)
				return v163
			} else {
				v57 = v10 + int32(12)
				v60 = F_NUM_cache(m, v49, v57, v18, v10+int32(11))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					v66 = F_palloc(m, v49<<(uint(int32(3))%32)|int32(1))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						v68 = int32(1)
						v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v72 = v70 & v68
						if v72 != 0 {
							v73 = v68
						} else {
							v73 = int32(4)
						}
						if v70 == int32(1) {
							v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
							if v80 == int32(18) {
								v83 = int32(16)
							} else {
								v83 = int32(0)
							}
							if base.Ui32((v80-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v90 = int32(4)
							} else {
								v90 = v83
							}
							v101 = v90
						} else {
							v91 = int32(1)
							if v72 != 0 {
								v101 = int32(base.Ui32(v70)>>(uint(v91)%32)) - v91
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v101 = int32(base.Ui32(v95)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v102 = int32(0)
						F_NUM_processor(m, v60, v57, v13+v73, v66, v101, v102, v102, v102)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
							if v112 == int32(1) {
								F_pfree(m, v60)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									v127 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v66), int64(0), base.I64_extend_i32_s((v107+(v108+v109))<<(uint(int32(16))%32)|v108+int32(4)))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int64(0)
									} else {
										v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v129&int32(8) != 0 {
											v133 = int32(0)
											v137 = F_int64_to_numeric(m, int64(10))
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int64(0)
											} else {
												v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
												v144 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v141))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int64(0)
												} else {
													v147 = F_DirectFunctionCall2Coll(m, int32(1435), v133, base.I64_extend_i32_u(v137), base.I64_extend_i32_u(v144))
													mBase = m.M
													v148 = m.ExcPending
													if v148 != 0 {
														return int64(0)
													} else {
														v150 = F_pg_detoast_datum(m, base.I32_wrap_i64(v147))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int64(0)
														} else {
															v153 = F_DirectFunctionCall2Coll(m, int32(1408), v133, v127, base.I64_extend_i32_u(v150))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return int64(0)
															} else {
																v155 = v153
																F_pfree(m, v66)
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int64(0)
																} else {
																	v163 = v155
																	m.G0 = v10 + int32(48)
																	return v163
																}
															}
														}
													}
												}
											}
										} else {
											v155 = v127
											F_pfree(m, v66)
											mBase = m.M
											v157 = m.ExcPending
											if v157 != 0 {
												return int64(0)
											} else {
												v163 = v155
												m.G0 = v10 + int32(48)
												return v163
											}
										}
									}
								}
							} else {
								v127 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v66), int64(0), base.I64_extend_i32_s((v107+(v108+v109))<<(uint(int32(16))%32)|v108+int32(4)))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int64(0)
								} else {
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v129&int32(8) != 0 {
										v133 = int32(0)
										v137 = F_int64_to_numeric(m, int64(10))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int64(0)
										} else {
											v141 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
											v144 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v141))
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return int64(0)
											} else {
												v147 = F_DirectFunctionCall2Coll(m, int32(1435), v133, base.I64_extend_i32_u(v137), base.I64_extend_i32_u(v144))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int64(0)
												} else {
													v150 = F_pg_detoast_datum(m, base.I32_wrap_i64(v147))
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int64(0)
													} else {
														v153 = F_DirectFunctionCall2Coll(m, int32(1408), v133, v127, base.I64_extend_i32_u(v150))
														mBase = m.M
														v154 = m.ExcPending
														if v154 != 0 {
															return int64(0)
														} else {
															v155 = v153
															F_pfree(m, v66)
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int64(0)
															} else {
																v163 = v155
																m.G0 = v10 + int32(48)
																return v163
															}
														}
													}
												}
											}
										}
									} else {
										v155 = v127
										F_pfree(m, v66)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int64(0)
										} else {
											v163 = v155
											m.G0 = v10 + int32(48)
											return v163
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
func F_numeric_trim_scale(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = int32(base.Ui32(v16) >> (uint(int32(2)) % 32))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
	if base.Ui32(int32(_a_F_numeric_trim_scale_0)) <= base.Ui32(v19) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return base.I64_extend_i32_u(v124)
L4:
	;
	v22 = F_palloc(m, v18)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v32 = base.I32_extend16_s(v19)
	v34 = base.B2i32(int32(0) <= v32)
	if int32(0) <= v32 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v26 = int32(base.Ui32(v24) >> (uint(int32(2)) % 32))
	if v26 == int32(0) {
		v124 = v22
		goto L3
	} else {
		goto L8
	}
L8:
	;
	base.MemoryCopy(m, v22, v12, v26)
	v124 = v22
	goto L3
L9:
	;
	v35 = int32(-8)
	goto L11
L10:
	;
	v35 = int32(-6)
	goto L11
L11:
	;
	v38 = int32(base.Ui32(v18+v35) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v38
	if int32(0) <= v32 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
	v50 = v40
	goto L14
L13:
	;
	v50 = v19<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v19&int32(63)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v50
	v57 = v19 & int32(_a_F_numeric_trim_scale_0)
	if v57 == int32(_a_F_numeric_trim_scale_1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v60 = v19 << (uint(int32(1)) % 32) & int32(_a_F_numeric_trim_scale_2)
	goto L17
L16:
	;
	v60 = v57
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v62
	if v32 < v62 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = int32(6)
	goto L20
L19:
	;
	v69 = int32(8)
	goto L20
L20:
	;
	v70 = v12 + v69
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v70
	v72 = v38
	goto L22
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v113
	v122 = F_make_result_safe(m, v9+int32(8), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L22:
	;
	if v72 <= int32(0) {
		v113 = v62
		goto L21
	} else {
		goto L24
	}
L23:
	;
	v90 = (v81 - v50) << (uint(int32(2)) % 32)
	if v90 <= int32(0) {
		v113 = v62
		goto L21
	} else {
		goto L26
	}
L24:
	;
	v80 = int32(1)
	v81 = v72 - v80
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+v81<<(uint(v80)%32)))))
	if v85 == int32(0) {
		v72 = v81
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v95 = base.I32_rem_s(base.I32_extend16_s(v85), int32(10))
	if v95 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v113 = v90
	goto L21
L28:
	;
	goto L29
L29:
	;
	v97 = v90
	v98 = v85
	goto L30
L30:
	;
	v103 = v97 - int32(1)
	v105 = int32(10)
	v106 = base.I32_div_s(base.I32_extend16_s(v98), v105)
	v109 = base.I32_rem_s(base.I32_extend16_s(v106), v105)
	if v109 == int32(0) {
		v97 = v103
		v98 = v106
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v113 = v103
	goto L21
L32:
	;
	goto L31
L33:
	;
	v124 = v122
	goto L3
}
