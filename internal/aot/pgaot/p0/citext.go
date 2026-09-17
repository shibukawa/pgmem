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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	v52 = F_str_tolower(m, v23, v50, int32(100))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v29 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v40 = int32(1)
	if v22 != 0 {
		v50 = int32(base.Ui32(v20)>>(uint(v40)%32)) - v40
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v32 = int32(16)
	goto L13
L12:
	;
	v32 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v29-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v39 = int32(4)
	goto L16
L15:
	;
	v39 = v32
	goto L16
L16:
	;
	v50 = v39
	goto L7
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v54 = int32(1)
	v55 = v14 + v54
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v60 = v58 & v54
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v61 = v55
	goto L21
L20:
	;
	v61 = v14 + int32(4)
	goto L21
L21:
	;
	if v58 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v90 = F_str_tolower(m, v61, v88, int32(100))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v67 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v78 = int32(1)
	if v60 != 0 {
		v88 = int32(base.Ui32(v58)>>(uint(v78)%32)) - v78
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v70 = int32(16)
	goto L28
L27:
	;
	v70 = int32(0)
	goto L28
L28:
	;
	if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v77 = int32(4)
	goto L31
L30:
	;
	v77 = v70
	goto L31
L31:
	;
	v88 = v77
	goto L22
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v88 = int32(base.Ui32(v82)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if base.B2i32(v94 == int32(0))|base.B2i32(v94 != v97) != 0 {
		v115 = v94
		v116 = v97
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v52)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v100 = v52
	v101 = v90
	goto L37
L37:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v105
		v116 = v104
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v115 = v105
	v116 = v104
	goto L35
L39:
	;
	v108 = int32(1)
	if v105 == v104 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_pfree(m, v90)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v122 != v9 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v9)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v126 != v14 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v14)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	return base.B2i32(v115-v116 == int32(0))
L50:
	;
	goto L49
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
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
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
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
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_str_tolower(m, v18, v45, int32(100))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = F_strlen(m, v47)
			mBase = m.M
			v55 = v49 - int32(1636608432)
			if v47&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v49) {
					v164 = v47
					v165 = v49
					v166 = v55
					v167 = v55
					v168 = v55
					for {
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
						v171 = v170 + v167
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v164)+8))
						v175 = v174 + v168
						v177 = int32(4)
						v179 = v172 + v166 - v175 ^ base.I32_rotl(v175, v177)
						v183 = v171 - v179 ^ base.I32_rotl(v179, int32(6))
						v184 = v175 + v171
						v185 = v179 + v184
						v186 = v183 + v185
						v190 = v184 - v183 ^ base.I32_rotl(v183, int32(8))
						v194 = v185 - v190 ^ base.I32_rotl(v190, int32(16))
						v198 = v186 - v194 ^ base.I32_rotl(v194, int32(19))
						v199 = v190 + v186
						v200 = v194 + v199
						v201 = v198 + v200
						v205 = v199 - v198 ^ base.I32_rotl(v198, v177)
						v206 = int32(12)
						v207 = v164 + v206
						v209 = v165 - v206
						if base.Ui32(int32(11)) < base.Ui32(v209) {
							v164 = v207
							v165 = v209
							v166 = v200
							v167 = v201
							v168 = v205
							continue
						} else {
							break
						}
						break
					}
					v212 = v207
					v213 = v209
					v214 = v200
					v215 = v201
					v216 = v205
				} else {
					v212 = v47
					v213 = v49
					v214 = v55
					v215 = v55
					v216 = v55
				}
				switch v213 - int32(1) {
				case 0:
					v275 = v214
					v276 = v215
					v277 = v216
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 1:
					v268 = v214
					v269 = v215
					v270 = v216
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 2:
					v261 = v214
					v262 = v215
					v263 = v216
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 3:
					v255 = v215
					v256 = v216
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 4:
					v251 = v215
					v252 = v216
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 5:
					v245 = v215
					v246 = v216
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 6:
					v239 = v215
					v240 = v216
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
					v245 = v241<<(uint(int32(16))%32) + v239
					v246 = v240
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 7:
					v234 = v216
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
					v239 = v235<<(uint(int32(24))%32) + v215
					v240 = v234
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
					v245 = v241<<(uint(int32(16))%32) + v239
					v246 = v240
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 8:
					v229 = v216
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
					v239 = v235<<(uint(int32(24))%32) + v215
					v240 = v234
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
					v245 = v241<<(uint(int32(16))%32) + v239
					v246 = v240
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 9:
					v224 = v216
					v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+9)))
					v229 = v225<<(uint(int32(16))%32) + v224
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
					v239 = v235<<(uint(int32(24))%32) + v215
					v240 = v234
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
					v245 = v241<<(uint(int32(16))%32) + v239
					v246 = v240
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				case 10:
					v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+10)))
					v224 = v220<<(uint(int32(24))%32) + v216
					v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+9)))
					v229 = v225<<(uint(int32(16))%32) + v224
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+8)))
					v234 = v230<<(uint(int32(8))%32) + v229
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+7)))
					v239 = v235<<(uint(int32(24))%32) + v215
					v240 = v234
					v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+6)))
					v245 = v241<<(uint(int32(16))%32) + v239
					v246 = v240
					v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+5)))
					v251 = v247<<(uint(int32(8))%32) + v245
					v252 = v246
					v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
					v255 = v251 + v253
					v256 = v252
					v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+3)))
					v261 = v257<<(uint(int32(24))%32) + v214
					v262 = v255
					v263 = v256
					v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+2)))
					v268 = v264<<(uint(int32(16))%32) + v261
					v269 = v262
					v270 = v263
					v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+1)))
					v275 = v271<<(uint(int32(8))%32) + v268
					v276 = v269
					v277 = v270
					v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
					v282 = v275 + v278
					v283 = v276
					v284 = v277
				default:
					v282 = v214
					v283 = v215
					v284 = v216
				}
			} else {
				if base.Ui32(v49) < base.Ui32(int32(12)) {
					v110 = v47
					v111 = v49
					v112 = v55
					v113 = v55
					v114 = v55
				} else {
					v62 = v47
					v63 = v49
					v64 = v55
					v65 = v55
					v66 = v55
					for {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
						v69 = v68 + v65
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
						v73 = v72 + v66
						v75 = int32(4)
						v77 = v70 + v64 - v73 ^ base.I32_rotl(v73, v75)
						v81 = v69 - v77 ^ base.I32_rotl(v77, int32(6))
						v82 = v73 + v69
						v83 = v77 + v82
						v84 = v81 + v83
						v88 = v82 - v81 ^ base.I32_rotl(v81, int32(8))
						v92 = v83 - v88 ^ base.I32_rotl(v88, int32(16))
						v96 = v84 - v92 ^ base.I32_rotl(v92, int32(19))
						v97 = v88 + v84
						v98 = v92 + v97
						v99 = v96 + v98
						v103 = v97 - v96 ^ base.I32_rotl(v96, v75)
						v104 = int32(12)
						v105 = v62 + v104
						v107 = v63 - v104
						if base.Ui32(int32(11)) < base.Ui32(v107) {
							v62 = v105
							v63 = v107
							v64 = v98
							v65 = v99
							v66 = v103
							continue
						} else {
							break
						}
						break
					}
					v110 = v105
					v111 = v107
					v112 = v98
					v113 = v99
					v114 = v103
				}
				switch v111 - int32(1) {
				case 0:
					v161 = v112
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
					v282 = v161 + v162
					v283 = v113
					v284 = v114
				case 1:
					v156 = v112
					v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
					v161 = v157<<(uint(int32(8))%32) + v156
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
					v282 = v161 + v162
					v283 = v113
					v284 = v114
				case 2:
					v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+2)))
					v156 = v152<<(uint(int32(16))%32) + v112
					v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
					v161 = v157<<(uint(int32(8))%32) + v156
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
					v282 = v161 + v162
					v283 = v113
					v284 = v114
				case 3:
					v149 = v113
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v282 = v150 + v112
					v283 = v149
					v284 = v114
				case 4:
					v146 = v113
					v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
					v149 = v146 + v147
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v282 = v150 + v112
					v283 = v149
					v284 = v114
				case 5:
					v141 = v113
					v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+5)))
					v146 = v142<<(uint(int32(8))%32) + v141
					v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
					v149 = v146 + v147
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v282 = v150 + v112
					v283 = v149
					v284 = v114
				case 6:
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+6)))
					v141 = v137<<(uint(int32(16))%32) + v113
					v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+5)))
					v146 = v142<<(uint(int32(8))%32) + v141
					v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+4)))
					v149 = v146 + v147
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v282 = v150 + v112
					v283 = v149
					v284 = v114
				case 7:
					v132 = v114
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
					v282 = v133 + v112
					v283 = v135 + v113
					v284 = v132
				case 8:
					v127 = v114
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+8)))
					v132 = v128<<(uint(int32(8))%32) + v127
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
					v282 = v133 + v112
					v283 = v135 + v113
					v284 = v132
				case 9:
					v122 = v114
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+9)))
					v127 = v123<<(uint(int32(16))%32) + v122
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+8)))
					v132 = v128<<(uint(int32(8))%32) + v127
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
					v282 = v133 + v112
					v283 = v135 + v113
					v284 = v132
				case 10:
					v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+10)))
					v122 = v118<<(uint(int32(24))%32) + v114
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+9)))
					v127 = v123<<(uint(int32(16))%32) + v122
					v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+8)))
					v132 = v128<<(uint(int32(8))%32) + v127
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
					v282 = v133 + v112
					v283 = v135 + v113
					v284 = v132
				default:
					v282 = v112
					v283 = v113
					v284 = v114
				}
			}
			v287 = int32(14)
			v289 = v283 ^ v284 - base.I32_rotl(v283, v287)
			v293 = v289 ^ v282 - base.I32_rotl(v289, int32(11))
			v297 = v293 ^ v283 - base.I32_rotl(v293, int32(25))
			v301 = v297 ^ v289 - base.I32_rotl(v297, int32(16))
			v305 = v301 ^ v293 - base.I32_rotl(v301, int32(4))
			v309 = v305 ^ v297 - base.I32_rotl(v305, v287)
			F_pfree(m, v47)
			mBase = m.M
			v315 = m.ExcPending
			if v315 != 0 {
				return int32(0)
			} else {
				v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v316 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v319 = m.ExcPending
					if v319 != 0 {
						return int32(0)
					} else {
						return v309 ^ v301 - base.I32_rotl(v309, int32(24))
					}
				} else {
					return v309 ^ v301 - base.I32_rotl(v309, int32(24))
				}
			}
		}
	}
}
