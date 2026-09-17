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
									F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2255), int32(_a_F_numeric_abbrev_abort_3))
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
										F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2276), int32(_a_F_numeric_abbrev_abort_3))
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
										F_errfinish(m, int32(_a_F_numeric_abbrev_abort_2), int32(2284), int32(_a_F_numeric_abbrev_abort_3))
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
func F_numeric_accum(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13949(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_numeric_add(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_add_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_ceil(m *base.Module, l0 int32) int32 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v357
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
	v353 = F_make_result_opt_error(m, v12, int32(0))
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
func F_numeric_cmp_abbrev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return base.B2i32(l0 < l1) - base.B2i32(l1 < l0)
}
func F_numeric_combine(m *base.Module, l0 int32) int32 {
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int64
	_ = v198
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
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
		switch v16 - int32(429) {
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
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+368))
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
		v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v42 == int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v46 = v45
		} else {
			v46 = v2
		}
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v47 != 0 {
			v226 = v46
			m.G0 = v8 + int32(32)
			return v226
		} else {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v48 == int32(0) {
				v226 = v46
				m.G0 = v8 + int32(32)
				return v226
			} else {
				if v46 == int32(0) {
					v53 = int32(_a_F_numeric_combine_0)
					v54 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v56
					v59 = F_palloc0(m, int32(112))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
						v66 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v66
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+8)) = v68
						v70 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+88)) = v70
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+96)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+104)) = v74
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v59)+72)) = v76
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v59)+80)) = v78
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
						v83 = F_palloc(m, v80<<(uint(int32(2))%32))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v59)+36)) = v83
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							v89 = F_palloc(m, v86<<(uint(int32(2))%32))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v59)+40)) = v89
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v94 = v92 << (uint(int32(2)) % 32)
								if v94 != 0 {
									v95 = *(*int32)(unsafe.Add(mBase, uint32(v59)+36))
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v48)+36))
									base.MemoryCopy(m, v95, v96, v94)
								} else {
								}
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								v100 = v98 << (uint(int32(2)) % 32)
								if v100 != 0 {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v59)+40))
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v48)+40))
									base.MemoryCopy(m, v101, v102, v100)
								} else {
								}
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v48)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+28)) = v104
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v106
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v110
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
								v115 = F_palloc(m, v112<<(uint(int32(2))%32))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v59)+64)) = v115
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
									v121 = F_palloc(m, v118<<(uint(int32(2))%32))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v59)+68)) = v121
										v124 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
										v126 = v124 << (uint(int32(2)) % 32)
										if v126 != 0 {
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v59)+64))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v48)+64))
											base.MemoryCopy(m, v127, v128, v126)
										} else {
										}
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
										v132 = v130 << (uint(int32(2)) % 32)
										if v132 != 0 {
											v133 = *(*int32)(unsafe.Add(mBase, uint32(v59)+68))
											v134 = *(*int32)(unsafe.Add(mBase, uint32(v48)+68))
											base.MemoryCopy(m, v133, v134, v132)
										} else {
										}
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v48)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+56)) = v136
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v48)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+44)) = v138
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+48)) = v140
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+52)) = v142
										v220 = v59
										v222 = v54
										*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v222
										v226 = v220
										m.G0 = v8 + int32(32)
										return v226
									}
								}
							}
						}
					}
				} else {
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v46)+8))
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v144 + v145
					v148 = *(*int64)(unsafe.Add(mBase, uint32(v46)+88))
					v149 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+88)) = v148 + v149
					v152 = *(*int64)(unsafe.Add(mBase, uint32(v46)+96))
					v153 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+96)) = v152 + v153
					v156 = *(*int64)(unsafe.Add(mBase, uint32(v46)+104))
					v157 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v46)+104)) = v156 + v157
					v160 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					if v160 <= int64(0) {
						v226 = v46
						m.G0 = v8 + int32(32)
						return v226
					} else {
						v163 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						v164 = *(*int32)(unsafe.Add(mBase, uint32(v46)+72))
						if v164 < v163 {
							*(*int32)(unsafe.Add(mBase, uint32(v46)+72)) = v163
							v167 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v167
						} else {
							if v164 != v163 {
							} else {
								v170 = *(*int64)(unsafe.Add(mBase, uint32(v46)+80))
								v171 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v46)+80)) = v170 + v171
							}
						}
						v174 = int32(_a_F_numeric_combine_0)
						v175 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0]))
						v177 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v177
						v179 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v179
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v179
						*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v179
						v188 = v8 + int32(8)
						F_accum_sum_final(m, v48+int32(16), v188)
						mBase = m.M
						v190 = m.ExcPending
						if v190 != 0 {
							return int32(0)
						} else {
							F_accum_sum_add(m, v46+int32(16), v188)
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								v195 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
								if v195 != 0 {
									F_pfree(m, v195)
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int32(0)
									} else {
										v198 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v198
										*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v198
										*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v198
										v207 = v8 + int32(8)
										F_accum_sum_final(m, v48+int32(44), v207)
										mBase = m.M
										v209 = m.ExcPending
										if v209 != 0 {
											return int32(0)
										} else {
											F_accum_sum_add(m, v46+int32(44), v207)
											mBase = m.M
											v213 = m.ExcPending
											if v213 != 0 {
												return int32(0)
											} else {
												v214 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
												if v214 == int32(0) {
													v220 = v46
													v222 = v175
													*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v222
													v226 = v220
													m.G0 = v8 + int32(32)
													return v226
												} else {
													F_pfree(m, v214)
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return int32(0)
													} else {
														v220 = v46
														v222 = v175
														*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v222
														v226 = v220
														m.G0 = v8 + int32(32)
														return v226
													}
												}
											}
										}
									}
								} else {
									v198 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v198
									*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v198
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v198
									v207 = v8 + int32(8)
									F_accum_sum_final(m, v48+int32(44), v207)
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return int32(0)
									} else {
										F_accum_sum_add(m, v46+int32(44), v207)
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											v214 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
											if v214 == int32(0) {
												v220 = v46
												v222 = v175
												*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v222
												v226 = v220
												m.G0 = v8 + int32(32)
												return v226
											} else {
												F_pfree(m, v214)
												mBase = m.M
												v218 = m.ExcPending
												if v218 != 0 {
													return int32(0)
												} else {
													v220 = v46
													v222 = v175
													*(*int32)(unsafe.Add(mBase, _c_F_numeric_combine[0])) = v222
													v226 = v220
													m.G0 = v8 + int32(32)
													return v226
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v236 = m.ExcPending
		if v236 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_combine_1), int32(0))
			mBase = m.M
			v240 = m.ExcPending
			if v240 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_combine_2), int32(_a_F_numeric_combine_3), int32(_a_F_numeric_combine_4))
				mBase = m.M
				v245 = m.ExcPending
				if v245 != 0 {
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
func F_numeric_div(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_div_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_exp(m *base.Module, l0 int32) int32 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(_a_F_numeric_exp_0)) <= base.Ui32(v16) {
			if v17 == int32(-4096) {
				v24 = F_make_result_opt_error(m, int32(_a_F_numeric_exp_1), int32(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v127 = v24
					m.G0 = v9 + int32(48)
					return v127
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v29 = F_palloc(m, int32(base.Ui32(v26)>>(uint(int32(2))%32)))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
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
					return v127
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
				return int32(0)
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
					return int32(0)
				} else {
					v120 = F_make_result_opt_error(m, v9, int32(0))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
						if v122 == int32(0) {
							v127 = v120
							m.G0 = v9 + int32(48)
							return v127
						} else {
							F_pfree(m, v122)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								v127 = v120
								m.G0 = v9 + int32(48)
								return v127
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	v6 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v25 = int32(_a_F_numeric_fast_cmp_0)
			v26 = v24 & v25
			if v26 == v25 {
				if v24 != int32(_a_F_numeric_fast_cmp_1) {
					if v24 != int32(_a_F_numeric_fast_cmp_0) {
						if v23 != int32(_a_F_numeric_fast_cmp_2) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v162 = v45
					} else {
						v162 = base.B2i32(v23 != int32(_a_F_numeric_fast_cmp_0))
					}
				} else {
					if v23 == int32(_a_F_numeric_fast_cmp_0) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v23 != int32(_a_F_numeric_fast_cmp_1))
					}
					v162 = v40
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_fast_cmp_0)) <= base.Ui32(v23) {
					if v23 == int32(_a_F_numeric_fast_cmp_2) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v162 = v52
				} else {
					v54 = v6 + int32(6)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v62 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v63 = int32(-8)
					} else {
						v63 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v75 = v65
					} else {
						v75 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v77 = int32(base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))+v63) >> (uint(int32(1)) % 32))
					v79 = v10 + int32(6)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v23))
					if int32(0) <= base.I32_extend16_s(v23) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v23) {
						v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79))))
						v100 = v90
					} else {
						v100 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					v101 = int32(1)
					v102 = int32(base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))+v88) >> (uint(v101) % 32))
					v108 = v23 & int32(_a_F_numeric_fast_cmp_0)
					if v108 == int32(_a_F_numeric_fast_cmp_3) {
						v111 = v23 << (uint(v101) % 32) & int32(_a_F_numeric_fast_cmp_4)
					} else {
						v111 = v108
					}
					if v77 == int32(0) {
						if v102 == int32(0) {
							v162 = int32(0)
						} else {
							if v111 == int32(_a_F_numeric_fast_cmp_4) {
								v121 = int32(1)
							} else {
								v121 = int32(-1)
							}
							v162 = v121
						}
					} else {
						if v26 == int32(_a_F_numeric_fast_cmp_3) {
							v128 = v24 << (uint(int32(1)) % 32) & int32(_a_F_numeric_fast_cmp_4)
						} else {
							v128 = v26
						}
						if v102 == int32(0) {
							if v128 != 0 {
								v133 = int32(-1)
							} else {
								v133 = int32(1)
							}
							v162 = v133
						} else {
							if int32(0) <= base.I32_extend16_s(v24) {
								v136 = v6 + int32(8)
							} else {
								v136 = v54
							}
							if int32(0) <= base.I32_extend16_s(v23) {
								v139 = v10 + int32(8)
							} else {
								v139 = v79
							}
							if v128 == int32(0) {
								if v111 == int32(_a_F_numeric_fast_cmp_4) {
									v162 = int32(1)
								} else {
									v145 = F_cmp_abs_common(m, v136, v77, v75, v139, v102, v100)
									mBase = m.M
									v162 = v145
								}
							} else {
								if v111 == int32(0) {
									v162 = int32(-1)
								} else {
									v149 = F_cmp_abs_common(m, v139, v102, v100, v136, v77, v75)
									mBase = m.M
									v162 = v149
								}
							}
						}
					}
				}
			}
			if l0 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v165 = m.ExcPending
				if v165 != 0 {
					return int32(0)
				} else {
					if l1 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							return v162
						}
					} else {
						return v162
					}
				}
			} else {
				if l1 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int32(0)
					} else {
						return v162
					}
				} else {
					return v162
				}
			}
		}
	}
}
func F_numeric_gt(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_gt_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_gt_1) {
					if v25 != int32(_a_F_numeric_gt_0) {
						if v24 != int32(_a_F_numeric_gt_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_gt_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_gt_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_gt_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_gt_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_gt_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_gt_0)
					if v109 == int32(_a_F_numeric_gt_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_gt_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_gt_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_gt_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_gt_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_gt_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v163)
						}
					} else {
						return base.B2i32(int32(0) < v163)
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return base.B2i32(int32(0) < v163)
					}
				} else {
					return base.B2i32(int32(0) < v163)
				}
			}
		}
	}
}
func F_numeric_le(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_le_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_le_1) {
					if v25 != int32(_a_F_numeric_le_0) {
						if v24 != int32(_a_F_numeric_le_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_le_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_le_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_le_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_le_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_le_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_le_0)
					if v109 == int32(_a_F_numeric_le_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_le_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_le_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_le_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_le_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_le_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v163 <= int32(0))
						}
					} else {
						return base.B2i32(v163 <= int32(0))
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v163 <= int32(0))
					}
				} else {
					return base.B2i32(v163 <= int32(0))
				}
			}
		}
	}
}
func F_numeric_mod_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int64
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	if l2 != 0 {
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v14)
	} else {
	}
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = base.I32_extend16_s(v16)
	if base.Ui32(v16) <= base.Ui32(int32(_a_F_numeric_mod_opt_error_0)) {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v21 = base.I32_extend16_s(v20)
		if base.Ui32(int32(_a_F_numeric_mod_opt_error_0)) < base.Ui32(v20) {
			v49 = v21
			if v49&int32(_a_F_numeric_mod_opt_error_1) != int32(_a_F_numeric_mod_opt_error_2) {
				if v17&int32(-8193) == int32(-12288) {
					if base.Ui32(int32(_a_F_numeric_mod_opt_error_0)) < base.Ui32(v49&int32(_a_F_numeric_mod_opt_error_1)) {
						v103 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v242 = v103
							m.G0 = v12 + int32(96)
							return v242
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v49) {
							v78 = int32(-8)
						} else {
							v78 = int32(-6)
						}
						if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v70)>>(uint(int32(2))%32))+v78) {
							v103 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v242 = v103
								m.G0 = v12 + int32(96)
								return v242
							}
						} else {
							if l2 != 0 {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v82)
								v242 = int32(0)
								m.G0 = v12 + int32(96)
								return v242
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_mod_opt_error_4), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_numeric_mod_opt_error_5), int32(3517), int32(_a_F_numeric_mod_opt_error_6))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
						}
					}
				} else {
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = F_palloc(m, int32(base.Ui32(v105)>>(uint(int32(2))%32)))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
						if v112 == int32(0) {
							v242 = v108
						} else {
							base.MemoryCopy(m, v108, l0, v112)
							v242 = v108
						}
						m.G0 = v12 + int32(96)
						return v242
					}
				}
			} else {
				v58 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v242 = v58
					m.G0 = v12 + int32(96)
					return v242
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v31 = int32(-8)
			} else {
				v31 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(base.Ui32(int32(base.Ui32(v24)>>(uint(int32(2))%32))+v31) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v117 = v116
			} else {
				v117 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v117
			v119 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v119
			v128 = base.B2i32(v17 < v119)
			if v17 < v119 {
				v129 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v129 = v16 & int32(_a_F_numeric_mod_opt_error_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v129
			v136 = v16 & int32(_a_F_numeric_mod_opt_error_2)
			if v136 == int32(_a_F_numeric_mod_opt_error_8) {
				v139 = v16 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mod_opt_error_9)
			} else {
				v139 = v136
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v139
			if v17 < v119 {
				v143 = int32(6)
			} else {
				v143 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = l0 + v143
			v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v152 = base.B2i32(int32(0) <= v21)
			if int32(0) <= v21 {
				v153 = int32(-8)
			} else {
				v153 = int32(-6)
			}
			v156 = int32(base.Ui32(int32(base.Ui32(v146)>>(uint(int32(2))%32))+v153) >> (uint(int32(1)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v156
			if int32(0) <= v21 {
				v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v168 = v158
			} else {
				v168 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v168
			v170 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v170
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v170
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v170
			v176 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v176
			v185 = base.B2i32(v21 < v176)
			if v21 < v176 {
				v186 = int32(base.Ui32(v20)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v186 = v20 & int32(_a_F_numeric_mod_opt_error_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v186
			v193 = v20 & int32(_a_F_numeric_mod_opt_error_2)
			if v193 == int32(_a_F_numeric_mod_opt_error_8) {
				v196 = v20 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mod_opt_error_9)
			} else {
				v196 = v193
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v196
			if v21 < v176 {
				v200 = int32(6)
			} else {
				v200 = int32(8)
			}
			v201 = l1 + v200
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v201
			if l2 == int32(0) {
				v209 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v209
				*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v209
				*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v209
				v216 = v12 + int32(48)
				v218 = v12 + int32(24)
				v220 = v12 + int32(72)
				v221 = int32(0)
				F_div_var(m, v216, v218, v220, v221, v221, int32(1))
				mBase = m.M
				v225 = m.ExcPending
				if v225 != 0 {
					return int32(0)
				} else {
					F_mul_var(m, v218, v220, v220, v186)
					mBase = m.M
					v227 = m.ExcPending
					if v227 != 0 {
						return int32(0)
					} else {
						F_sub_var(m, v216, v220, v12)
						mBase = m.M
						v229 = m.ExcPending
						if v229 != 0 {
							return int32(0)
						} else {
							v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
							if v230 != 0 {
								F_pfree(m, v230)
								mBase = m.M
								v232 = m.ExcPending
								if v232 != 0 {
									return int32(0)
								} else {
									v234 = F_make_result_opt_error(m, v12, int32(0))
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
										return int32(0)
									} else {
										v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
										if v236 == int32(0) {
											v242 = v234
											m.G0 = v12 + int32(96)
											return v242
										} else {
											F_pfree(m, v236)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v242 = v234
												m.G0 = v12 + int32(96)
												return v242
											}
										}
									}
								}
							} else {
								v234 = F_make_result_opt_error(m, v12, int32(0))
								mBase = m.M
								v235 = m.ExcPending
								if v235 != 0 {
									return int32(0)
								} else {
									v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
									if v236 == int32(0) {
										v242 = v234
										m.G0 = v12 + int32(96)
										return v242
									} else {
										F_pfree(m, v236)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int32(0)
										} else {
											v242 = v234
											m.G0 = v12 + int32(96)
											return v242
										}
									}
								}
							}
						}
					}
				}
			} else {
				if v156 != 0 {
					v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
					if v205 != 0 {
						v209 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v209
						*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v209
						*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v209
						v216 = v12 + int32(48)
						v218 = v12 + int32(24)
						v220 = v12 + int32(72)
						v221 = int32(0)
						F_div_var(m, v216, v218, v220, v221, v221, int32(1))
						mBase = m.M
						v225 = m.ExcPending
						if v225 != 0 {
							return int32(0)
						} else {
							F_mul_var(m, v218, v220, v220, v186)
							mBase = m.M
							v227 = m.ExcPending
							if v227 != 0 {
								return int32(0)
							} else {
								F_sub_var(m, v216, v220, v12)
								mBase = m.M
								v229 = m.ExcPending
								if v229 != 0 {
									return int32(0)
								} else {
									v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
									if v230 != 0 {
										F_pfree(m, v230)
										mBase = m.M
										v232 = m.ExcPending
										if v232 != 0 {
											return int32(0)
										} else {
											v234 = F_make_result_opt_error(m, v12, int32(0))
											mBase = m.M
											v235 = m.ExcPending
											if v235 != 0 {
												return int32(0)
											} else {
												v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
												if v236 == int32(0) {
													v242 = v234
													m.G0 = v12 + int32(96)
													return v242
												} else {
													F_pfree(m, v236)
													mBase = m.M
													v240 = m.ExcPending
													if v240 != 0 {
														return int32(0)
													} else {
														v242 = v234
														m.G0 = v12 + int32(96)
														return v242
													}
												}
											}
										}
									} else {
										v234 = F_make_result_opt_error(m, v12, int32(0))
										mBase = m.M
										v235 = m.ExcPending
										if v235 != 0 {
											return int32(0)
										} else {
											v236 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
											if v236 == int32(0) {
												v242 = v234
												m.G0 = v12 + int32(96)
												return v242
											} else {
												F_pfree(m, v236)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													v242 = v234
													m.G0 = v12 + int32(96)
													return v242
												}
											}
										}
									}
								}
							}
						}
					} else {
						v206 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v206)
						v242 = int32(0)
						m.G0 = v12 + int32(96)
						return v242
					}
				} else {
					v206 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v206)
					v242 = int32(0)
					m.G0 = v12 + int32(96)
					return v242
				}
			}
		}
	} else {
		if v17 == int32(-16384) {
			v58 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v242 = v58
				m.G0 = v12 + int32(96)
				return v242
			}
		} else {
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v49 = v47
			if v49&int32(_a_F_numeric_mod_opt_error_1) != int32(_a_F_numeric_mod_opt_error_2) {
				if v17&int32(-8193) == int32(-12288) {
					if base.Ui32(int32(_a_F_numeric_mod_opt_error_0)) < base.Ui32(v49&int32(_a_F_numeric_mod_opt_error_1)) {
						v103 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v242 = v103
							m.G0 = v12 + int32(96)
							return v242
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v49) {
							v78 = int32(-8)
						} else {
							v78 = int32(-6)
						}
						if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v70)>>(uint(int32(2))%32))+v78) {
							v103 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v242 = v103
								m.G0 = v12 + int32(96)
								return v242
							}
						} else {
							if l2 != 0 {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v82)
								v242 = int32(0)
								m.G0 = v12 + int32(96)
								return v242
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_mod_opt_error_4), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_numeric_mod_opt_error_5), int32(3517), int32(_a_F_numeric_mod_opt_error_6))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
						}
					}
				} else {
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = F_palloc(m, int32(base.Ui32(v105)>>(uint(int32(2))%32)))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
						if v112 == int32(0) {
							v242 = v108
						} else {
							base.MemoryCopy(m, v108, l0, v112)
							v242 = v108
						}
						m.G0 = v12 + int32(96)
						return v242
					}
				}
			} else {
				v58 = F_make_result_opt_error(m, int32(_a_F_numeric_mod_opt_error_3), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v242 = v58
					m.G0 = v12 + int32(96)
					return v242
				}
			}
		}
	}
}
func F_numeric_ne(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_ne_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_ne_1) {
					if v25 != int32(_a_F_numeric_ne_0) {
						if v24 != int32(_a_F_numeric_ne_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_ne_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_ne_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_ne_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_ne_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_ne_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_ne_0)
					if v109 == int32(_a_F_numeric_ne_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_ne_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_ne_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_ne_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_ne_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_ne_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v163 != int32(0))
						}
					} else {
						return base.B2i32(v163 != int32(0))
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v163 != int32(0))
					}
				} else {
					return base.B2i32(v163 != int32(0))
				}
			}
		}
	}
}
func F_numeric_sign(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
		if v11 == int32(_a_F_numeric_sign_0) {
			v47 = int32(_a_F_numeric_sign_1)
		} else {
			v14 = base.I32_extend16_s(v11)
			v15 = int32(_a_F_numeric_sign_0)
			v16 = v11 & v15
			if v16 == v15 {
				if v14 != int32(-12288) {
					v47 = int32(_a_F_numeric_sign_2)
				} else {
					v47 = int32(_a_F_numeric_sign_3)
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if int32(0) <= v14 {
					v30 = int32(-8)
				} else {
					v30 = int32(-6)
				}
				if base.Ui32(int32(base.Ui32(v23)>>(uint(int32(2))%32))+v30) < base.Ui32(int32(2)) {
					v47 = int32(_a_F_numeric_sign_4)
				} else {
					if v16 == int32(_a_F_numeric_sign_5) {
						v40 = v11 << (uint(int32(1)) % 32) & int32(_a_F_numeric_sign_6)
					} else {
						v40 = v16
					}
					if v40 == int32(_a_F_numeric_sign_6) {
						v47 = int32(_a_F_numeric_sign_2)
					} else {
						v47 = int32(_a_F_numeric_sign_3)
					}
				}
			}
		}
		v49 = F_make_result_opt_error(m, v47, int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func F_numeric_to_number(m *base.Module, l0 int32) int32 {
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
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
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
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
				v153 = int32(0)
				m.G0 = v10 + int32(48)
				return v153
			} else {
				v58 = v10 + int32(12)
				v61 = F_NUM_cache(m, v49, v58, v18, v10+int32(11))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v67 = F_palloc(m, v49<<(uint(int32(3))%32)|int32(1))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = int32(1)
						v70 = v13 + v69
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v75 = v73 & v69
						if v75 != 0 {
							v76 = v70
						} else {
							v76 = v13 + int32(4)
						}
						if v73 == int32(1) {
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
							if v82 == int32(18) {
								v85 = int32(16)
							} else {
								v85 = int32(0)
							}
							if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v92 = int32(4)
							} else {
								v92 = v85
							}
							v103 = v92
						} else {
							v93 = int32(1)
							if v75 != 0 {
								v103 = int32(base.Ui32(v73)>>(uint(v93)%32)) - v93
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v103 = int32(base.Ui32(v97)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v104 = int32(0)
						F_NUM_processor(m, v61, v58, v76, v67, v103, v104, v104, v104)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
							if v114 == int32(1) {
								F_pfree(m, v61)
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
									return int32(0)
								} else {
									v120 = int32(0)
									v127 = F_DirectFunctionCall3Coll(m, int32(408), v120, v67, v120, (v109+(v110+v111))<<(uint(int32(16))%32)|v110+int32(4))
									mBase = m.M
									v128 = m.ExcPending
									if v128 != 0 {
										return int32(0)
									} else {
										v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v129&int32(8) != 0 {
											v133 = int32(0)
											v137 = F_int64_to_numeric(m, int64(10))
											mBase = m.M
											v138 = m.ExcPending
											if v138 != 0 {
												return int32(0)
											} else {
												v142 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v109))
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return int32(0)
												} else {
													v144 = F_DirectFunctionCall2Coll(m, int32(1297), v133, v137, v142)
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return int32(0)
													} else {
														v146 = F_pg_detoast_datum(m, v144)
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return int32(0)
														} else {
															v148 = F_DirectFunctionCall2Coll(m, int32(1262), v133, v127, v146)
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																v150 = v148
																F_pfree(m, v67)
																mBase = m.M
																v152 = m.ExcPending
																if v152 != 0 {
																	return int32(0)
																} else {
																	v153 = v150
																	m.G0 = v10 + int32(48)
																	return v153
																}
															}
														}
													}
												}
											}
										} else {
											v150 = v127
											F_pfree(m, v67)
											mBase = m.M
											v152 = m.ExcPending
											if v152 != 0 {
												return int32(0)
											} else {
												v153 = v150
												m.G0 = v10 + int32(48)
												return v153
											}
										}
									}
								}
							} else {
								v120 = int32(0)
								v127 = F_DirectFunctionCall3Coll(m, int32(408), v120, v67, v120, (v109+(v110+v111))<<(uint(int32(16))%32)|v110+int32(4))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v129&int32(8) != 0 {
										v133 = int32(0)
										v137 = F_int64_to_numeric(m, int64(10))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											v142 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v109))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return int32(0)
											} else {
												v144 = F_DirectFunctionCall2Coll(m, int32(1297), v133, v137, v142)
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return int32(0)
												} else {
													v146 = F_pg_detoast_datum(m, v144)
													mBase = m.M
													v147 = m.ExcPending
													if v147 != 0 {
														return int32(0)
													} else {
														v148 = F_DirectFunctionCall2Coll(m, int32(1262), v133, v127, v146)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															v150 = v148
															F_pfree(m, v67)
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																v153 = v150
																m.G0 = v10 + int32(48)
																return v153
															}
														}
													}
												}
											}
										}
									} else {
										v150 = v127
										F_pfree(m, v67)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											v153 = v150
											m.G0 = v10 + int32(48)
											return v153
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
func F_numeric_trim_scale(m *base.Module, l0 int32) int32 {
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v124
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
	v122 = F_make_result_opt_error(m, v9+int32(8), int32(0))
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
