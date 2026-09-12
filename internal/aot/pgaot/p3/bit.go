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
							F_errmsg(m, int32(661529), v9)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(491490), int32(407), int32(103301))
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
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
	return v226
L2:
	;
	if int32(0) < v13 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v29 = v15 + int32(1)
	v30 = F_strlen(m, v29)
	mBase = m.M
	if int32(536870911) <= v30 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v27 = F_strlen(m, v25)
	mBase = m.M
	v57 = v25
	v58 = int32(1)
	v59 = v27
	goto L2
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
	v33 = F_errsave_start(m, v14)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v57 = v29
	v58 = v2
	v59 = v30 << (uint(int32(2)) % 32)
	goto L2
L12:
	;
	return int32(0)
L13:
	;
	if v33 == int32(0) {
		v226 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483640)
	F_errmsg(m, int32(662487), v9+int32(-16))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, v14, int32(491490), int32(199), int32(278463))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v226 = v2
	goto L1
L18:
	;
	v173 = v77
	v174 = int32(128)
	v175 = v57
	goto L52
L19:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v100 == int32(0) {
		v226 = v70
		goto L1
	} else {
		goto L32
	}
L20:
	;
	v81 = F_errsave_start(m, v14)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L27
	}
L21:
	;
	if v59 != v13 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v63 = v59
	goto L23
L23:
	;
	v66 = int32(8)
	v67 = base.I32_div_s(v63+int32(7), v66)
	v69 = v67 + v66
	v70 = F_palloc0(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	v63 = v13
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v69 << (uint(int32(2)) % 32)
	v77 = v70 + int32(8)
	if v58 == int32(0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if v81 == int32(0) {
		v226 = v2
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(101187714))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v59
	F_errmsg(m, int32(661529), v9+int32(-32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, v14, int32(491490), int32(213), int32(278463))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v226 = v2
	goto L1
L32:
	;
	v103 = v77
	v104 = v100
	v105 = v57
	v110 = v2
	goto L33
L33:
	;
	v112 = v104 - int32(48)
	if base.Ui32(v112&int32(255)) < base.Ui32(int32(10)) {
		v133 = v112
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v226 = v70
	goto L1
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v166)
	v171 = v105 + int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	if v172 != 0 {
		v103 = v168
		v104 = v172
		v105 = v171
		v110 = v167
		goto L33
	} else {
		goto L51
	}
L36:
	;
	v166 = v133 << (uint(int32(4)) % 32)
	v167 = int32(1)
	v168 = v103
	goto L35
L37:
	;
	v141 = int32(0)
	v142 = F_errsave_start(m, v14)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L12
	} else {
		goto L45
	}
L38:
	;
	if v110 == int32(0) {
		goto L36
	} else {
		goto L44
	}
L39:
	;
	if base.Ui32((v104-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v133 = v104 - int32(55)
	goto L38
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(5)) < base.Ui32((v104-int32(97))&int32(255)) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v133 = v104 - int32(87)
	goto L38
L44:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v166 = v136 | v133
	v167 = int32(0)
	v168 = v103 + int32(1)
	goto L35
L45:
	;
	if v142 == int32(0) {
		v226 = v141
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v149 = F_pg_mblen_cstr(m, v105)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v149
	F_errmsg(m, int32(102630), v9+int32(-48))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, v14, int32(491490), int32(260), int32(278463))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v226 = v141
	goto L1
L51:
	;
	goto L34
L52:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	switch v181 - int32(48) {
	case 0:
		goto L55
	case 1:
		goto L56
	default:
		goto L54
	}
L53:
	;
	if v181 == int32(0) {
		v226 = v70
		goto L1
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v189 = v174 & int32(255)
	v193 = base.B2i32(base.Ui32(v189) < base.Ui32(int32(2)))
	if base.Ui32(v189) < base.Ui32(int32(2)) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	v185 = v184 | v174
	*(*uint8)(unsafe.Add(mBase, uint32(v173))) = uint8(v185)
	goto L55
L57:
	;
	v194 = int32(-128)
	goto L59
L58:
	;
	v194 = int32(base.Ui32(v189) >> (uint(int32(1)) % 32))
	goto L59
L59:
	;
	v173 = v173 + v193
	v174 = v194
	v175 = v175 + int32(1)
	goto L52
L60:
	;
	v200 = int32(0)
	v201 = F_errsave_start(m, v14)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	if v201 == int32(0) {
		v226 = v200
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v208 = F_pg_mblen_cstr(m, v175)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v208
	F_errmsg(m, int32(102569), v11)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v14, int32(491490), int32(235), int32(278463))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v226 = v200
	goto L1
}
