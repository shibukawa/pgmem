package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bit(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if base.Ui32(v17-int32(2147483641)) < base.Ui32(int32(-2147483640)) {
			v66 = v13
			m.G0 = v10 + int32(16)
			return v66
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v22 == v17 {
				v66 = v13
				m.G0 = v10 + int32(16)
				return v66
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				if v24 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(101187714))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v17
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v83
							F_errmsg(m, int32(_a_F_bit_0), v10)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_bit_1), int32(407), int32(_a_F_bit_2))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
					v30 = int32(base.Ui32(v17+int32(7)) >> (uint(int32(3)) % 32))
					v32 = v30 + int32(8)
					v33 = F_palloc0(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v17
						v36 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v33))) = v32 << (uint(v36) % 32)
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v43 = int32(base.Ui32(v39)>>(uint(v36)%32)) - int32(8)
						if base.Ui32(v30) < base.Ui32(v43) {
							v45 = v30
						} else {
							v45 = v43
						}
						if v45 != 0 {
							v46 = int32(8)
							base.MemoryCopy(m, v33+v46, v13+v46, v45)
						} else {
						}
						v55 = v32<<(uint(int32(3))%32) - v17 + int32(-64)
						if v55 <= int32(0) {
							v66 = v33
						} else {
							v60 = v33 + v32 - int32(1)
							v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v64 = v61 & (int32(255) << (uint(v55) % 32))
							*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v64)
							v66 = v33
						}
						m.G0 = v10 + int32(16)
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
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v55 int64
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int64
	_ = v95
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int64
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	v6 = int64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(8)
		v13 = v8 + v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
		v18 = v16 - v12
		if base.Ui32(v14) <= base.Ui32(int32(47)) {
			if v18 == int32(0) {
				v24 = F_Int64GetDatum(m, int64(0))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					return v24
				}
			} else {
				v27 = int32(3)
				v28 = v16 & v27
				if base.Ui32(v27) <= base.Ui32(v16-int32(9)) {
					v36 = v13
					v38 = int32(0)
					v41 = v6
					for {
						v42 = int32(4)
						v43 = v36 + v42
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+3)))
						v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v44)+uint32(_c_F_bit_bit_count[0]))))
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+2)))
						v47 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_c_F_bit_bit_count[0]))))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
						v49 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_bit_bit_count[0]))))
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
						v51 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_bit_bit_count[0]))))
						v55 = v45 + (v47 + (v49 + (v41 + v51)))
						v57 = v38 + v42
						if v57 != v18&int32(-4) {
							v36 = v43
							v38 = v57
							v41 = v55
							continue
						} else {
							break
						}
						break
					}
					if v28 == int32(0) {
						v87 = v55
					} else {
						v61 = v43
						v66 = v55
						v68 = v61
						v69 = int32(0)
						v73 = v66
						for {
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
							v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_bit_bit_count[0]))))
							v76 = v73 + v75
							v77 = int32(1)
							v80 = v69 + v77
							if v80 != v28 {
								v68 = v68 + v77
								v69 = v80
								v73 = v76
								continue
							} else {
								break
							}
							break
						}
						v87 = v76
					}
				} else {
					v61 = v13
					v66 = v6
					v68 = v61
					v69 = int32(0)
					v73 = v66
					for {
						v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
						v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_bit_bit_count[0]))))
						v76 = v73 + v75
						v77 = int32(1)
						v80 = v69 + v77
						if v80 != v28 {
							v68 = v68 + v77
							v69 = v80
							v73 = v76
							continue
						} else {
							break
						}
						break
					}
					v87 = v76
				}
				v88 = F_Int64GetDatum(m, v87)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					return v88
				}
			}
		} else {
			v95 = int64(0)
			if base.B2i32(v13 != (v8+int32(11))&int32(-4))|base.B2i32(v18 < int32(4)) != 0 {
				v174 = v13
				v175 = v18
				v180 = v95
			} else {
				v105 = v18 - int32(4)
				v109 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) + int32(1)
				v111 = v109 & int32(3)
				if base.Ui32(int32(12)) <= base.Ui32(v105) {
					v116 = v13
					v117 = v18
					v120 = int32(0)
					v122 = v95
					for {
						v123 = int32(16)
						v124 = v117 - v123
						v126 = v116 + v123
						v127 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v116)+8))
						v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
						v142 = base.I64_extend_i32_u(base.I32_popcnt(v127)) + (base.I64_extend_i32_u(base.I32_popcnt(v130)) + (base.I64_extend_i32_u(base.I32_popcnt(v133)) + (v122 + base.I64_extend_i32_u(base.I32_popcnt(v136)))))
						v144 = v120 + int32(4)
						if v144 != v109&int32(2147483644) {
							v116 = v126
							v117 = v124
							v120 = v144
							v122 = v142
							continue
						} else {
							break
						}
						break
					}
					if v111 == int32(0) {
						v174 = v126
						v175 = v124
						v180 = v142
					} else {
						v148 = v126
						v149 = v124
						v154 = v142
						v156 = v148
						v157 = v149
						v158 = int32(0)
						v162 = v154
						for {
							v163 = int32(4)
							v164 = v157 - v163
							v166 = v156 + v163
							v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
							v170 = v162 + base.I64_extend_i32_u(base.I32_popcnt(v167))
							v172 = v158 + int32(1)
							if v172 != v111 {
								v156 = v166
								v157 = v164
								v158 = v172
								v162 = v170
								continue
							} else {
								break
							}
							break
						}
						v174 = v166
						v175 = v164
						v180 = v170
					}
				} else {
					v148 = v13
					v149 = v18
					v154 = v95
					v156 = v148
					v157 = v149
					v158 = int32(0)
					v162 = v154
					for {
						v163 = int32(4)
						v164 = v157 - v163
						v166 = v156 + v163
						v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
						v170 = v162 + base.I64_extend_i32_u(base.I32_popcnt(v167))
						v172 = v158 + int32(1)
						if v172 != v111 {
							v156 = v166
							v157 = v164
							v158 = v172
							v162 = v170
							continue
						} else {
							break
						}
						break
					}
					v174 = v166
					v175 = v164
					v180 = v170
				}
			}
			if v175 == int32(0) {
				v243 = v180
			} else {
				v184 = v175 & int32(3)
				if v184 == int32(0) {
					v205 = v174
					v207 = v175
					v211 = v180
				} else {
					v188 = v174
					v190 = v175
					v192 = int32(0)
					v194 = v180
					for {
						v195 = int32(1)
						v196 = v188 + v195
						v198 = v190 - v195
						v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
						v200 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v199)+uint32(_c_F_bit_bit_count[0]))))
						v201 = v194 + v200
						v203 = v192 + v195
						if v203 != v184 {
							v188 = v196
							v190 = v198
							v192 = v203
							v194 = v201
							continue
						} else {
							break
						}
						break
					}
					v205 = v196
					v207 = v198
					v211 = v201
				}
				if base.Ui32(v175) < base.Ui32(int32(4)) {
					v243 = v211
				} else {
					v214 = v205
					v216 = v207
					v220 = v211
					for {
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+3)))
						v222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_bit_bit_count[0]))))
						v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+2)))
						v224 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v223)+uint32(_c_F_bit_bit_count[0]))))
						v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+1)))
						v226 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v225)+uint32(_c_F_bit_bit_count[0]))))
						v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
						v228 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v227)+uint32(_c_F_bit_bit_count[0]))))
						v232 = v222 + (v224 + (v226 + (v220 + v228)))
						v233 = int32(4)
						v236 = v216 - v233
						if v236 != 0 {
							v214 = v214 + v233
							v216 = v236
							v220 = v232
							continue
						} else {
							break
						}
						break
					}
					v243 = v232
				}
			}
			v244 = F_Int64GetDatum(m, v243)
			mBase = m.M
			v245 = m.ExcPending
			if v245 != 0 {
				return int32(0)
			} else {
				return v244
			}
		}
	}
}
func F_bit_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	switch v17 - int32(88) {
	case 0:
		goto L3
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		v26 = v16
		goto L4
	case 10:
		goto L5
	default:
		goto L6
	}
L1:
	;
	m.G0 = v12 - int32(-64)
	return v230
L2:
	;
	if int32(0) < v15 {
		goto L21
	} else {
		goto L22
	}
L3:
	;
	v30 = v16 + int32(1)
	v31 = F_strlen(m, v30)
	mBase = m.M
	if int32(536870911) <= v31 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v28 = F_strlen(m, v26)
	mBase = m.M
	v57 = v26
	v59 = int32(1)
	v60 = v28
	goto L2
L5:
	;
	v26 = v16 + int32(1)
	goto L4
L6:
	;
	if v17 == int32(120) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	if v17 != int32(66) {
		v26 = v16
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v34 = F_errsave_start(m, v14)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v57 = v30
	v59 = v2
	v60 = v31 << (uint(int32(2)) % 32)
	goto L2
L12:
	;
	return int32(0)
L13:
	;
	if v34 == int32(0) {
		v230 = v2
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(2147483640)
	F_errmsg(m, int32(_a_F_bit_in_0), v10+int32(-16))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, v14, int32(_a_F_bit_in_1), int32(199), int32(_a_F_bit_in_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v230 = v2
	goto L1
L18:
	;
	v175 = v57
	v176 = v78
	v179 = int32(128)
	goto L52
L19:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	if v101 == int32(0) {
		v230 = v71
		goto L1
	} else {
		goto L32
	}
L20:
	;
	v82 = F_errsave_start(m, v14)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L12
	} else {
		goto L27
	}
L21:
	;
	if v60 != v15 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v64 = v60
	goto L23
L23:
	;
	v67 = int32(8)
	v68 = base.I32_div_s(v64+int32(7), v67)
	v70 = v68 + v67
	v71 = F_palloc0(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L25
	}
L24:
	;
	v64 = v15
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v70 << (uint(int32(2)) % 32)
	v78 = v71 + int32(8)
	if v59 == int32(0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	if v82 == int32(0) {
		v230 = v2
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(101187714))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v60
	F_errmsg(m, int32(_a_F_bit_in_3), v10+int32(-32))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L30
	}
L30:
	;
	F_errsave_finish(m, v14, int32(_a_F_bit_in_1), int32(213), int32(_a_F_bit_in_2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	v230 = v2
	goto L1
L32:
	;
	v104 = v57
	v105 = v78
	v109 = v101
	v112 = v2
	goto L33
L33:
	;
	v114 = v109 - int32(48)
	if base.Ui32(v114&int32(255)) < base.Ui32(int32(10)) {
		v135 = v114
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v230 = v71
	goto L1
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v105))) = uint8(v168)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v172 != 0 {
		v104 = v104 + int32(1)
		v105 = v170
		v109 = v172
		v112 = v169
		goto L33
	} else {
		goto L51
	}
L36:
	;
	v168 = v135 << (uint(int32(4)) % 32)
	v169 = int32(1)
	v170 = v105
	goto L35
L37:
	;
	v143 = int32(0)
	v144 = F_errsave_start(m, v14)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L45
	}
L38:
	;
	if v112 == int32(0) {
		goto L36
	} else {
		goto L44
	}
L39:
	;
	if base.Ui32((v109-int32(65))&int32(255)) <= base.Ui32(int32(5)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v135 = v109 - int32(55)
	goto L38
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(5)) < base.Ui32((v109-int32(97))&int32(255)) {
		goto L37
	} else {
		goto L43
	}
L43:
	;
	v135 = v109 - int32(87)
	goto L38
L44:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v168 = v138 | v135
	v169 = int32(0)
	v170 = v105 + int32(1)
	goto L35
L45:
	;
	if v144 == int32(0) {
		v230 = v143
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v151 = F_pg_mblen_cstr(m, v104)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v151
	F_errmsg(m, int32(_a_F_bit_in_4), v10+int32(-48))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, v14, int32(_a_F_bit_in_1), int32(260), int32(_a_F_bit_in_2))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L50
	}
L50:
	;
	v230 = v143
	goto L1
L51:
	;
	goto L34
L52:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	switch v184 - int32(48) {
	case 0:
		goto L55
	case 1:
		goto L56
	default:
		goto L54
	}
L53:
	;
	if v184 == int32(0) {
		v230 = v71
		goto L1
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v190 = int32(1)
	v195 = int32(base.Ui32(v179)>>(uint(v190)%32)) & int32(127)
	if v195 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v188 = v187 | v179
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v188)
	goto L55
L57:
	;
	v197 = v195
	goto L59
L58:
	;
	v197 = int32(-128)
	goto L59
L59:
	;
	v175 = v175 + v190
	v176 = v176 + base.B2i32(v195 == int32(0))
	v179 = v197
	goto L52
L60:
	;
	v203 = int32(0)
	v204 = F_errsave_start(m, v14)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	if v204 == int32(0) {
		v230 = v203
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	v211 = F_pg_mblen_cstr(m, v175)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v211
	F_errmsg(m, int32(_a_F_bit_in_5), v12)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	F_errsave_finish(m, v14, int32(_a_F_bit_in_1), int32(235), int32(_a_F_bit_in_2))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v230 = v203
	goto L1
}
