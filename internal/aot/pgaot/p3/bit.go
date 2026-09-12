package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if base.Ui32(v16-int32(2147483641)) < base.Ui32(int32(-2147483640)) {
			v66 = v12
			m.G0 = v9 + int32(16)
			return v66
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			if v21 == v16 {
				v66 = v12
				m.G0 = v9 + int32(16)
				return v66
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(101187714))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v16
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v82
							F_errmsg(m, int32(637289), v9)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(470441), int32(407), int32(96308))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v29 = int32(base.Ui32(v16+int32(7)) >> (uint(int32(3)) % 32))
					v31 = v29 + int32(8)
					v32 = F_palloc0(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v16
						v35 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31 << (uint(v35) % 32)
						v38 = int32(8)
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v46 = int32(base.Ui32(v42)>>(uint(v35)%32)) - v38
						if base.Ui32(v29) < base.Ui32(v46) {
							v48 = v29
						} else {
							v48 = v46
						}
						if v48 != 0 {
							v49 = F__emscripten_memcpy_bulkmem(m, v32+v38, v12+v38, v48)
							mBase = m.M
						} else {
						}
						v55 = v31<<(uint(int32(3))%32) - v16 + int32(-64)
						if v55 <= int32(0) {
							v66 = v32
						} else {
							v60 = v32 + v31 - int32(1)
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v64 = v61 & (int32(255) << (uint(v55) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v64)
							v66 = v32
						}
						m.G0 = v9 + int32(16)
						return v66
					}
				}
			}
		}
	}
}
func F_bit_bit_count(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int64
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int64
	_ = v98
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int64
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int64
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v245 int64
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(8)
		v12 = v7 + v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v15 = int32(base.Ui32(v13) >> (uint(int32(2)) % 32))
		v17 = v15 - v11
		if base.Ui32(v13) <= base.Ui32(int32(47)) {
			if v17 == int32(0) {
				v23 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			} else {
				v26 = int32(3)
				v27 = v15 & v26
				if base.Ui32(v15-int32(9)) < base.Ui32(v26) {
					v65 = v12
					v69 = v5
				} else {
					v35 = v12
					v36 = int32(0)
					v39 = v5
					for {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+3)))
						v43 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v40)+uint32(_consts[1045]))))
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+2)))
						v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v44)+uint32(_consts[1045]))))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
						v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_consts[1045]))))
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
						v55 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1045]))))
						v59 = v43 + (v47 + (v51 + (v39 + v55)))
						v60 = int32(4)
						v61 = v35 + v60
						v63 = v36 + v60
						if v63 != v17&int32(-4) {
							v35 = v61
							v36 = v63
							v39 = v59
							continue
						} else {
							break
						}
						break
					}
					v65 = v61
					v69 = v59
				}
				if v27 != 0 {
					v71 = v65
					v72 = int32(0)
					v75 = v69
					for {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
						v79 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v76)+uint32(_consts[1045]))))
						v80 = v75 + v79
						v81 = int32(1)
						v84 = v72 + v81
						if v84 != v27 {
							v71 = v71 + v81
							v72 = v84
							v75 = v80
							continue
						} else {
							break
						}
						break
					}
					v90 = v80
				} else {
					v90 = v69
				}
				v91 = F_Int64GetDatum(m, v90)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					return v91
				}
			}
		} else {
			v98 = int64(0)
			if v17 < int32(4) {
				v177 = v12
				v178 = v17
				v183 = v98
			} else {
				if v12 != (v7+int32(11))&int32(-4) {
					v177 = v12
					v178 = v17
					v183 = v98
				} else {
					v107 = v17 - int32(4)
					v111 = int32(base.Ui32(v107)>>(uint(int32(2))%32)) + int32(1)
					v113 = v111 & int32(3)
					if base.Ui32(v107) < base.Ui32(int32(12)) {
						v149 = v12
						v150 = v17
						v155 = v98
					} else {
						v119 = v12
						v120 = v17
						v121 = int32(0)
						v125 = v98
						for {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v119)+8))
							v132 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
							v141 = base.I64_extend_i32_u(base.I32_popcnt(v126)) + (base.I64_extend_i32_u(base.I32_popcnt(v129)) + (base.I64_extend_i32_u(base.I32_popcnt(v132)) + (v125 + base.I64_extend_i32_u(base.I32_popcnt(v135)))))
							v142 = int32(16)
							v143 = v120 - v142
							v145 = v119 + v142
							v147 = v121 + int32(4)
							if v147 != v111&int32(2147483644) {
								v119 = v145
								v120 = v143
								v121 = v147
								v125 = v141
								continue
							} else {
								break
							}
							break
						}
						v149 = v145
						v150 = v143
						v155 = v141
					}
					if v113 == int32(0) {
						v177 = v149
						v178 = v150
						v183 = v155
					} else {
						v160 = v150
						v161 = v149
						v162 = int32(0)
						v165 = v155
						for {
							v166 = int32(4)
							v167 = v160 - v166
							v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
							v171 = v165 + base.I64_extend_i32_u(base.I32_popcnt(v168))
							v173 = v161 + v166
							v175 = v162 + int32(1)
							if v175 != v113 {
								v160 = v167
								v161 = v173
								v162 = v175
								v165 = v171
								continue
							} else {
								break
							}
							break
						}
						v177 = v173
						v178 = v167
						v183 = v171
					}
				}
			}
			if v178 == int32(0) {
				v256 = v183
			} else {
				v187 = v178 & int32(3)
				if v187 == int32(0) {
					v210 = v177
					v212 = v178
					v216 = v183
				} else {
					v193 = v178
					v194 = v177
					v195 = int32(0)
					v197 = v183
					for {
						v198 = int32(1)
						v199 = v193 - v198
						v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
						v203 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v200)+uint32(_consts[1045]))))
						v204 = v197 + v203
						v206 = v194 + v198
						v208 = v195 + v198
						if v208 != v187 {
							v193 = v199
							v194 = v206
							v195 = v208
							v197 = v204
							continue
						} else {
							break
						}
						break
					}
					v210 = v206
					v212 = v199
					v216 = v204
				}
				if base.Ui32(v178) < base.Ui32(int32(4)) {
					v256 = v216
				} else {
					v219 = v210
					v221 = v212
					v225 = v216
					for {
						v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+3)))
						v229 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v226)+uint32(_consts[1045]))))
						v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+2)))
						v233 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_consts[1045]))))
						v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
						v237 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v234)+uint32(_consts[1045]))))
						v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
						v241 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v238)+uint32(_consts[1045]))))
						v245 = v229 + (v233 + (v237 + (v225 + v241)))
						v246 = int32(4)
						v249 = v221 - v246
						if v249 != 0 {
							v219 = v219 + v246
							v221 = v249
							v225 = v245
							continue
						} else {
							break
						}
						break
					}
					v256 = v245
				}
			}
			v257 = F_Int64GetDatum(m, v256)
			mBase = m.M
			v258 = m.ExcPending
			if v258 != 0 {
				return int32(0)
			} else {
				return v257
			}
		}
	}
}
func F_bit_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	switch v16 - int32(88) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v25 = v15
		goto L4
	case 10:
		goto L5
	default:
		goto L6
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v338
L2:
	;
	if int32(0) < v13 {
		goto L55
	} else {
		goto L56
	}
L3:
	;
	v85 = v15 + int32(1)
	if v85&int32(3) == int32(0) {
		v109 = v85
		goto L28
	} else {
		goto L29
	}
L4:
	;
	if v25&int32(3) == int32(0) {
		v50 = v25
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v25 = v15 + int32(1)
	goto L4
L6:
	;
	if v16 == int32(120) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v16 != int32(66) {
		v25 = v15
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v169 = v25
	v170 = int32(1)
	v171 = v83
	goto L2
L10:
	;
	v83 = v75 - v25
	goto L9
L11:
	;
	v54 = v50
	goto L20
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v83 = int32(0)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = v25
	goto L16
L16:
	;
	v43 = v39 + int32(1)
	if v43&int32(3) == int32(0) {
		v50 = v43
		goto L11
	} else {
		goto L18
	}
L17:
	;
	v75 = v43
	goto L10
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v48 != 0 {
		v39 = v43
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v63 = int32(-2139062144)
	if (int32(16843008)-v60|v60)&v63 == v63 {
		v54 = v54 + int32(4)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v69 = v54
	goto L23
L22:
	;
	goto L21
L23:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v73 != 0 {
		v69 = v69 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v75 = v69
	goto L10
L25:
	;
	goto L24
L26:
	;
	if int32(536870911) <= v142 {
		goto L43
	} else {
		goto L44
	}
L27:
	;
	v142 = v134 - v85
	goto L26
L28:
	;
	v113 = v109
	goto L37
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v93 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v142 = int32(0)
	goto L26
L31:
	;
	goto L32
L32:
	;
	v98 = v85
	goto L33
L33:
	;
	v102 = v98 + int32(1)
	if v102&int32(3) == int32(0) {
		v109 = v102
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v134 = v102
	goto L27
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v107 != 0 {
		v98 = v102
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v122 = int32(-2139062144)
	if (int32(16843008)-v119|v119)&v122 == v122 {
		v113 = v113 + int32(4)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v128 = v113
	goto L40
L39:
	;
	goto L38
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v132 != 0 {
		v128 = v128 + int32(1)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v134 = v128
	goto L27
L42:
	;
	goto L41
L43:
	;
	v145 = F_errsave_start(m, v14)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v169 = v85
	v170 = v2
	v171 = v142 << (uint(int32(2)) % 32)
	goto L2
L46:
	;
	return int32(0)
L47:
	;
	if v145 == int32(0) {
		v338 = v2
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483640)
	F_errmsg(m, int32(638247), v9+int32(-16))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, v14, int32(470441), int32(199), int32(265574))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v338 = v2
	goto L1
L52:
	;
	v285 = v189
	v286 = int32(128)
	v287 = v169
	goto L86
L53:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v212 == int32(0) {
		v338 = v182
		goto L1
	} else {
		goto L66
	}
L54:
	;
	v193 = F_errsave_start(m, v14)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L46
	} else {
		goto L61
	}
L55:
	;
	if v171 != v13 {
		goto L54
	} else {
		goto L58
	}
L56:
	;
	v175 = v171
	goto L57
L57:
	;
	v178 = int32(8)
	v179 = base.I32_div_s(v175+int32(7), v178)
	v181 = v179 + v178
	v182 = F_palloc0(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L46
	} else {
		goto L59
	}
L58:
	;
	v175 = v13
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v181 << (uint(int32(2)) % 32)
	v189 = v182 + int32(8)
	if v170 == int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	goto L52
L61:
	;
	if v193 == int32(0) {
		v338 = v2
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(101187714))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L46
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v171
	F_errmsg(m, int32(637289), v9+int32(-32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L46
	} else {
		goto L64
	}
L64:
	;
	F_errsave_finish(m, v14, int32(470441), int32(213), int32(265574))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L46
	} else {
		goto L65
	}
L65:
	;
	v338 = v2
	goto L1
L66:
	;
	v215 = v189
	v216 = v212
	v217 = v169
	v222 = v2
	goto L67
L67:
	;
	v224 = v216 - int32(48)
	if base.Ui32(v224&int32(255)) < base.Ui32(int32(10)) {
		v245 = v224
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v338 = v182
	goto L1
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v215))) = uint8(v278)
	v283 = v217 + int32(1)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v284 != 0 {
		v215 = v280
		v216 = v284
		v217 = v283
		v222 = v279
		goto L67
	} else {
		goto L85
	}
L70:
	;
	v278 = v245 << (uint(int32(4)) % 32)
	v279 = int32(1)
	v280 = v215
	goto L69
L71:
	;
	v253 = int32(0)
	v254 = F_errsave_start(m, v14)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L46
	} else {
		goto L79
	}
L72:
	;
	if v222 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L73:
	;
	if base.Ui32((v216-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v245 = v216 - int32(55)
	goto L72
L75:
	;
	goto L76
L76:
	;
	if base.Ui32(int32(5)) < base.Ui32((v216-int32(97))&int32(255)) {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v245 = v216 - int32(87)
	goto L72
L78:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v278 = v248 | v245
	v279 = int32(0)
	v280 = v215 + int32(1)
	goto L69
L79:
	;
	if v254 == int32(0) {
		v338 = v253
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L46
	} else {
		goto L81
	}
L81:
	;
	v261 = F_pg_mblen_cstr(m, v217)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L46
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v261
	F_errmsg(m, int32(95827), v9+int32(-48))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L46
	} else {
		goto L83
	}
L83:
	;
	F_errsave_finish(m, v14, int32(470441), int32(260), int32(265574))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L46
	} else {
		goto L84
	}
L84:
	;
	v338 = v253
	goto L1
L85:
	;
	goto L68
L86:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	switch v293 - int32(48) {
	case 0:
		goto L89
	case 1:
		goto L90
	default:
		goto L88
	}
L87:
	;
	if v293 == int32(0) {
		v338 = v182
		goto L1
	} else {
		goto L94
	}
L88:
	;
	goto L87
L89:
	;
	v301 = v286 & int32(255)
	v305 = base.B2i32(base.Ui32(v301) < base.Ui32(int32(2)))
	if base.Ui32(v301) < base.Ui32(int32(2)) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v297 = v296 | v286
	*(*uint8)(unsafe.Add(mBase, uint32(v285))) = uint8(v297)
	goto L89
L91:
	;
	v306 = int32(-128)
	goto L93
L92:
	;
	v306 = int32(base.Ui32(v301) >> (uint(int32(1)) % 32))
	goto L93
L93:
	;
	v285 = v285 + v305
	v286 = v306
	v287 = v287 + int32(1)
	goto L86
L94:
	;
	v312 = int32(0)
	v313 = F_errsave_start(m, v14)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L46
	} else {
		goto L95
	}
L95:
	;
	if v313 == int32(0) {
		v338 = v312
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L46
	} else {
		goto L97
	}
L97:
	;
	v320 = F_pg_mblen_cstr(m, v287)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L46
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v320
	F_errmsg(m, int32(95766), v11)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L46
	} else {
		goto L99
	}
L99:
	;
	F_errsave_finish(m, v14, int32(470441), int32(235), int32(265574))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L46
	} else {
		goto L100
	}
L100:
	;
	v338 = v312
	goto L1
}
