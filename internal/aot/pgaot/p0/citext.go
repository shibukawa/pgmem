package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_eq(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(1)
	v17 = v9 + v16
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v22 = v20 & v16
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v17
	goto L6
L5:
	;
	v23 = v9 + int32(4)
	goto L6
L6:
	;
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = F_str_tolower(m, v23, v51, int32(100))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = v26
	goto L13
L12:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L13
L13:
	;
	if v28 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v26
	goto L16
L15:
	;
	v40 = v37
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v55 = int32(1)
	v56 = v14 + v55
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v61 = v59 & v55
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v62 = v56
	goto L21
L20:
	;
	v62 = v14 + int32(4)
	goto L21
L21:
	;
	if v59 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v92 = F_str_tolower(m, v62, v90, int32(100))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v65 = int32(4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v67&int32(254) == int32(2) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v80 = int32(1)
	if v61 != 0 {
		v90 = int32(base.Ui32(v59)>>(uint(v80)%32)) - v80
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v76 = v65
	goto L28
L27:
	;
	v76 = base.B2i32(v67 == int32(18)) << (uint(v65) % 32)
	goto L28
L28:
	;
	if v67 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v79 = v65
	goto L31
L30:
	;
	v79 = v76
	goto L31
L31:
	;
	v90 = v79
	goto L22
L32:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v97 == int32(0) {
		v116 = v96
		v117 = v97
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v53)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L42
	}
L35:
	;
	goto L34
L36:
	;
	if v96 != v97 {
		v116 = v96
		v117 = v97
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v101 = v53
	v102 = v92
	goto L38
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v105
		v117 = v106
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v116 = v105
	v117 = v106
	goto L35
L40:
	;
	v109 = int32(1)
	if v105 == v106 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_pfree(m, v92)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v123 != v9 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_pfree(m, v9)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v127 != v14 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L46
L48:
	;
	F_pfree(m, v14)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	return base.B2i32(v117-v116 == int32(0))
L51:
	;
	goto L50
}
func F_citext_ge(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v14^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_citext_hash(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v48 = F_str_tolower(m, v18, v46, int32(100))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return int32(0)
		} else {
			v50 = F_strlen(m, v48)
			mBase = m.M
			v56 = v50 - int32(1636608432)
			if v48&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v50) {
					v165 = v48
					v166 = v50
					v167 = v56
					v168 = v56
					v169 = v56
					for {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
						v172 = v171 + v168
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
						v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
						v176 = v175 + v169
						v178 = int32(4)
						v180 = v173 + v167 - v176 ^ base.I32_rotl(v176, v178)
						v184 = v172 - v180 ^ base.I32_rotl(v180, int32(6))
						v185 = v176 + v172
						v186 = v180 + v185
						v187 = v184 + v186
						v191 = v185 - v184 ^ base.I32_rotl(v184, int32(8))
						v195 = v186 - v191 ^ base.I32_rotl(v191, int32(16))
						v199 = v187 - v195 ^ base.I32_rotl(v195, int32(19))
						v200 = v191 + v187
						v201 = v195 + v200
						v202 = v199 + v201
						v206 = v200 - v199 ^ base.I32_rotl(v199, v178)
						v207 = int32(12)
						v208 = v165 + v207
						v210 = v166 - v207
						if base.Ui32(int32(11)) < base.Ui32(v210) {
							v165 = v208
							v166 = v210
							v167 = v201
							v168 = v202
							v169 = v206
							continue
						} else {
							break
						}
						break
					}
					v213 = v208
					v214 = v210
					v215 = v201
					v216 = v202
					v217 = v206
				} else {
					v213 = v48
					v214 = v50
					v215 = v56
					v216 = v56
					v217 = v56
				}
				switch v214 - int32(1) {
				case 0:
					v276 = v215
					v277 = v216
					v278 = v217
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 1:
					v269 = v215
					v270 = v216
					v271 = v217
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 2:
					v262 = v215
					v263 = v216
					v264 = v217
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 3:
					v256 = v216
					v257 = v217
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 4:
					v252 = v216
					v253 = v217
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 5:
					v246 = v216
					v247 = v217
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 6:
					v240 = v216
					v241 = v217
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
					v246 = v242<<(uint(int32(16))%32) + v240
					v247 = v241
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 7:
					v235 = v217
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+7)))
					v240 = v236<<(uint(int32(24))%32) + v216
					v241 = v235
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
					v246 = v242<<(uint(int32(16))%32) + v240
					v247 = v241
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 8:
					v230 = v217
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+8)))
					v235 = v231<<(uint(int32(8))%32) + v230
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+7)))
					v240 = v236<<(uint(int32(24))%32) + v216
					v241 = v235
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
					v246 = v242<<(uint(int32(16))%32) + v240
					v247 = v241
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 9:
					v225 = v217
					v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+9)))
					v230 = v226<<(uint(int32(16))%32) + v225
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+8)))
					v235 = v231<<(uint(int32(8))%32) + v230
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+7)))
					v240 = v236<<(uint(int32(24))%32) + v216
					v241 = v235
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
					v246 = v242<<(uint(int32(16))%32) + v240
					v247 = v241
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				case 10:
					v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+10)))
					v225 = v221<<(uint(int32(24))%32) + v217
					v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+9)))
					v230 = v226<<(uint(int32(16))%32) + v225
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+8)))
					v235 = v231<<(uint(int32(8))%32) + v230
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+7)))
					v240 = v236<<(uint(int32(24))%32) + v216
					v241 = v235
					v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+6)))
					v246 = v242<<(uint(int32(16))%32) + v240
					v247 = v241
					v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+5)))
					v252 = v248<<(uint(int32(8))%32) + v246
					v253 = v247
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+4)))
					v256 = v252 + v254
					v257 = v253
					v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+3)))
					v262 = v258<<(uint(int32(24))%32) + v215
					v263 = v256
					v264 = v257
					v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+2)))
					v269 = v265<<(uint(int32(16))%32) + v262
					v270 = v263
					v271 = v264
					v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
					v276 = v272<<(uint(int32(8))%32) + v269
					v277 = v270
					v278 = v271
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
					v283 = v276 + v279
					v284 = v277
					v285 = v278
				default:
					v283 = v215
					v284 = v216
					v285 = v217
				}
			} else {
				if base.Ui32(v50) < base.Ui32(int32(12)) {
					v111 = v48
					v112 = v50
					v113 = v56
					v114 = v56
					v115 = v56
				} else {
					v63 = v48
					v64 = v50
					v65 = v56
					v66 = v56
					v67 = v56
					for {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
						v70 = v69 + v66
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						v74 = v73 + v67
						v76 = int32(4)
						v78 = v71 + v65 - v74 ^ base.I32_rotl(v74, v76)
						v82 = v70 - v78 ^ base.I32_rotl(v78, int32(6))
						v83 = v74 + v70
						v84 = v78 + v83
						v85 = v82 + v84
						v89 = v83 - v82 ^ base.I32_rotl(v82, int32(8))
						v93 = v84 - v89 ^ base.I32_rotl(v89, int32(16))
						v97 = v85 - v93 ^ base.I32_rotl(v93, int32(19))
						v98 = v89 + v85
						v99 = v93 + v98
						v100 = v97 + v99
						v104 = v98 - v97 ^ base.I32_rotl(v97, v76)
						v105 = int32(12)
						v106 = v63 + v105
						v108 = v64 - v105
						if base.Ui32(int32(11)) < base.Ui32(v108) {
							v63 = v106
							v64 = v108
							v65 = v99
							v66 = v100
							v67 = v104
							continue
						} else {
							break
						}
						break
					}
					v111 = v106
					v112 = v108
					v113 = v99
					v114 = v100
					v115 = v104
				}
				switch v112 - int32(1) {
				case 0:
					v162 = v113
					v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
					v283 = v162 + v163
					v284 = v114
					v285 = v115
				case 1:
					v157 = v113
					v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
					v162 = v158<<(uint(int32(8))%32) + v157
					v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
					v283 = v162 + v163
					v284 = v114
					v285 = v115
				case 2:
					v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+2)))
					v157 = v153<<(uint(int32(16))%32) + v113
					v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
					v162 = v158<<(uint(int32(8))%32) + v157
					v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
					v283 = v162 + v163
					v284 = v114
					v285 = v115
				case 3:
					v150 = v114
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v283 = v151 + v113
					v284 = v150
					v285 = v115
				case 4:
					v147 = v114
					v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
					v150 = v147 + v148
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v283 = v151 + v113
					v284 = v150
					v285 = v115
				case 5:
					v142 = v114
					v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
					v147 = v143<<(uint(int32(8))%32) + v142
					v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
					v150 = v147 + v148
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v283 = v151 + v113
					v284 = v150
					v285 = v115
				case 6:
					v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+6)))
					v142 = v138<<(uint(int32(16))%32) + v114
					v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+5)))
					v147 = v143<<(uint(int32(8))%32) + v142
					v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
					v150 = v147 + v148
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v283 = v151 + v113
					v284 = v150
					v285 = v115
				case 7:
					v133 = v115
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
					v283 = v134 + v113
					v284 = v136 + v114
					v285 = v133
				case 8:
					v128 = v115
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+8)))
					v133 = v129<<(uint(int32(8))%32) + v128
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
					v283 = v134 + v113
					v284 = v136 + v114
					v285 = v133
				case 9:
					v123 = v115
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
					v128 = v124<<(uint(int32(16))%32) + v123
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+8)))
					v133 = v129<<(uint(int32(8))%32) + v128
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
					v283 = v134 + v113
					v284 = v136 + v114
					v285 = v133
				case 10:
					v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+10)))
					v123 = v119<<(uint(int32(24))%32) + v115
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+9)))
					v128 = v124<<(uint(int32(16))%32) + v123
					v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+8)))
					v133 = v129<<(uint(int32(8))%32) + v128
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
					v283 = v134 + v113
					v284 = v136 + v114
					v285 = v133
				default:
					v283 = v113
					v284 = v114
					v285 = v115
				}
			}
			v288 = int32(14)
			v290 = v284 ^ v285 - base.I32_rotl(v284, v288)
			v294 = v290 ^ v283 - base.I32_rotl(v290, int32(11))
			v298 = v294 ^ v284 - base.I32_rotl(v294, int32(25))
			v302 = v298 ^ v290 - base.I32_rotl(v298, int32(16))
			v306 = v302 ^ v294 - base.I32_rotl(v302, int32(4))
			v310 = v306 ^ v298 - base.I32_rotl(v306, v288)
			F_pfree(m, v48)
			mBase = m.M
			v316 = m.ExcPending
			if v316 != 0 {
				return int32(0)
			} else {
				v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v317 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v320 = m.ExcPending
					if v320 != 0 {
						return int32(0)
					} else {
						return v310 ^ v302 - base.I32_rotl(v310, int32(24))
					}
				} else {
					return v310 ^ v302 - base.I32_rotl(v310, int32(24))
				}
			}
		}
	}
}
