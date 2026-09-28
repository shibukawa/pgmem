package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_citext_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(1)
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	v19 = v17 & v15
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = v15
	goto L6
L5:
	;
	v20 = int32(4)
	goto L6
L6:
	;
	if v17 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v50 = F_str_tolower(m, v8+v20, v48, int32(100))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
	if v27 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v38 = int32(1)
	if v19 != 0 {
		v48 = int32(base.Ui32(v17)>>(uint(v38)%32)) - v38
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v30 = int32(16)
	goto L13
L12:
	;
	v30 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v27-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v37 = int32(4)
	goto L16
L15:
	;
	v37 = v30
	goto L16
L16:
	;
	v48 = v37
	goto L7
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v52 = int32(1)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v56 = v54 & v52
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v57 = v52
	goto L21
L20:
	;
	v57 = int32(4)
	goto L21
L21:
	;
	if v54 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v87 = F_str_tolower(m, v13+v57, v85, int32(100))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L33
	}
L23:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v64 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v75 = int32(1)
	if v56 != 0 {
		v85 = int32(base.Ui32(v54)>>(uint(v75)%32)) - v75
		goto L22
	} else {
		goto L32
	}
L26:
	;
	v67 = int32(16)
	goto L28
L27:
	;
	v67 = int32(0)
	goto L28
L28:
	;
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v74 = int32(4)
	goto L31
L30:
	;
	v74 = v67
	goto L31
L31:
	;
	v85 = v74
	goto L22
L32:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v85 = int32(base.Ui32(v79)>>(uint(int32(2))%32)) - int32(4)
	goto L22
L33:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if base.B2i32(v91 == int32(0))|base.B2i32(v91 != v94) != 0 {
		v112 = v91
		v113 = v94
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_pfree(m, v50)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	goto L34
L36:
	;
	v97 = v50
	v98 = v87
	goto L37
L37:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v102 == int32(0) {
		v112 = v102
		v113 = v101
		goto L35
	} else {
		goto L39
	}
L38:
	;
	v112 = v102
	v113 = v101
	goto L35
L39:
	;
	v105 = int32(1)
	if v102 == v101 {
		v97 = v97 + v105
		v98 = v98 + v105
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	F_pfree(m, v87)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v119 != v8 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_pfree(m, v8)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v123 != v13 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	F_pfree(m, v13)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	return base.I64_extend_i32_u(base.B2i32(v112-v113 == int32(0)))
L50:
	;
	goto L49
}
func F_citext_ge(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v14 = F_citextcmp(m, v6, v11, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v16 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int64(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v20 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) <= v14))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v14))
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v14))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) <= v14))
					}
				}
			}
		}
	}
}
func F_citext_hash(m *base.Module, l0 int32) int64 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
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
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = int32(1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v14 = v12 & v10
		if v14 != 0 {
			v15 = v10
		} else {
			v15 = int32(4)
		}
		if v12 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if v22 == int32(18) {
				v25 = int32(16)
			} else {
				v25 = int32(0)
			}
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v32 = int32(4)
			} else {
				v32 = v25
			}
			v43 = v32
		} else {
			v33 = int32(1)
			if v14 != 0 {
				v43 = int32(base.Ui32(v12)>>(uint(v33)%32)) - v33
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v45 = F_str_tolower(m, v6+v15, v43, int32(100))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			v47 = F_strlen(m, v45)
			mBase = m.M
			v53 = v47 - int32(1636608432)
			if v45&int32(3) != 0 {
				if base.Ui32(int32(11)) < base.Ui32(v47) {
					v162 = v45
					v163 = v47
					v164 = v53
					v165 = v53
					v166 = v53
					for {
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
						v169 = v168 + v165
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
						v173 = v172 + v166
						v175 = int32(4)
						v177 = v170 + v164 - v173 ^ base.I32_rotl(v173, v175)
						v181 = v169 - v177 ^ base.I32_rotl(v177, int32(6))
						v182 = v173 + v169
						v183 = v177 + v182
						v184 = v181 + v183
						v188 = v182 - v181 ^ base.I32_rotl(v181, int32(8))
						v192 = v183 - v188 ^ base.I32_rotl(v188, int32(16))
						v196 = v184 - v192 ^ base.I32_rotl(v192, int32(19))
						v197 = v188 + v184
						v198 = v192 + v197
						v199 = v196 + v198
						v203 = v197 - v196 ^ base.I32_rotl(v196, v175)
						v204 = int32(12)
						v205 = v162 + v204
						v207 = v163 - v204
						if base.Ui32(int32(11)) < base.Ui32(v207) {
							v162 = v205
							v163 = v207
							v164 = v198
							v165 = v199
							v166 = v203
							continue
						} else {
							break
						}
						break
					}
					v210 = v205
					v211 = v207
					v212 = v198
					v213 = v199
					v214 = v203
				} else {
					v210 = v45
					v211 = v47
					v212 = v53
					v213 = v53
					v214 = v53
				}
				switch v211 - int32(1) {
				case 0:
					v273 = v212
					v274 = v213
					v275 = v214
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 1:
					v266 = v212
					v267 = v213
					v268 = v214
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 2:
					v259 = v212
					v260 = v213
					v261 = v214
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 3:
					v253 = v213
					v254 = v214
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 4:
					v249 = v213
					v250 = v214
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 5:
					v243 = v213
					v244 = v214
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 6:
					v237 = v213
					v238 = v214
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
					v243 = v239<<(uint(int32(16))%32) + v237
					v244 = v238
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 7:
					v232 = v214
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
					v237 = v233<<(uint(int32(24))%32) + v213
					v238 = v232
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
					v243 = v239<<(uint(int32(16))%32) + v237
					v244 = v238
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 8:
					v227 = v214
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+8)))
					v232 = v228<<(uint(int32(8))%32) + v227
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
					v237 = v233<<(uint(int32(24))%32) + v213
					v238 = v232
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
					v243 = v239<<(uint(int32(16))%32) + v237
					v244 = v238
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 9:
					v222 = v214
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+9)))
					v227 = v223<<(uint(int32(16))%32) + v222
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+8)))
					v232 = v228<<(uint(int32(8))%32) + v227
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
					v237 = v233<<(uint(int32(24))%32) + v213
					v238 = v232
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
					v243 = v239<<(uint(int32(16))%32) + v237
					v244 = v238
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				case 10:
					v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+10)))
					v222 = v218<<(uint(int32(24))%32) + v214
					v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+9)))
					v227 = v223<<(uint(int32(16))%32) + v222
					v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+8)))
					v232 = v228<<(uint(int32(8))%32) + v227
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+7)))
					v237 = v233<<(uint(int32(24))%32) + v213
					v238 = v232
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+6)))
					v243 = v239<<(uint(int32(16))%32) + v237
					v244 = v238
					v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+5)))
					v249 = v245<<(uint(int32(8))%32) + v243
					v250 = v244
					v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+4)))
					v253 = v249 + v251
					v254 = v250
					v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+3)))
					v259 = v255<<(uint(int32(24))%32) + v212
					v260 = v253
					v261 = v254
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+2)))
					v266 = v262<<(uint(int32(16))%32) + v259
					v267 = v260
					v268 = v261
					v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
					v273 = v269<<(uint(int32(8))%32) + v266
					v274 = v267
					v275 = v268
					v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
					v280 = v273 + v276
					v281 = v274
					v282 = v275
				default:
					v280 = v212
					v281 = v213
					v282 = v214
				}
			} else {
				if base.Ui32(v47) < base.Ui32(int32(12)) {
					v108 = v45
					v109 = v47
					v110 = v53
					v111 = v53
					v112 = v53
				} else {
					v60 = v45
					v61 = v47
					v62 = v53
					v63 = v53
					v64 = v53
					for {
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
						v67 = v66 + v63
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
						v71 = v70 + v64
						v73 = int32(4)
						v75 = v68 + v62 - v71 ^ base.I32_rotl(v71, v73)
						v79 = v67 - v75 ^ base.I32_rotl(v75, int32(6))
						v80 = v71 + v67
						v81 = v75 + v80
						v82 = v79 + v81
						v86 = v80 - v79 ^ base.I32_rotl(v79, int32(8))
						v90 = v81 - v86 ^ base.I32_rotl(v86, int32(16))
						v94 = v82 - v90 ^ base.I32_rotl(v90, int32(19))
						v95 = v86 + v82
						v96 = v90 + v95
						v97 = v94 + v96
						v101 = v95 - v94 ^ base.I32_rotl(v94, v73)
						v102 = int32(12)
						v103 = v60 + v102
						v105 = v61 - v102
						if base.Ui32(int32(11)) < base.Ui32(v105) {
							v60 = v103
							v61 = v105
							v62 = v96
							v63 = v97
							v64 = v101
							continue
						} else {
							break
						}
						break
					}
					v108 = v103
					v109 = v105
					v110 = v96
					v111 = v97
					v112 = v101
				}
				switch v109 - int32(1) {
				case 0:
					v159 = v110
					v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
					v280 = v159 + v160
					v281 = v111
					v282 = v112
				case 1:
					v154 = v110
					v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
					v159 = v155<<(uint(int32(8))%32) + v154
					v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
					v280 = v159 + v160
					v281 = v111
					v282 = v112
				case 2:
					v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+2)))
					v154 = v150<<(uint(int32(16))%32) + v110
					v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
					v159 = v155<<(uint(int32(8))%32) + v154
					v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
					v280 = v159 + v160
					v281 = v111
					v282 = v112
				case 3:
					v147 = v111
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v280 = v148 + v110
					v281 = v147
					v282 = v112
				case 4:
					v144 = v111
					v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
					v147 = v144 + v145
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v280 = v148 + v110
					v281 = v147
					v282 = v112
				case 5:
					v139 = v111
					v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+5)))
					v144 = v140<<(uint(int32(8))%32) + v139
					v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
					v147 = v144 + v145
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v280 = v148 + v110
					v281 = v147
					v282 = v112
				case 6:
					v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+6)))
					v139 = v135<<(uint(int32(16))%32) + v111
					v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+5)))
					v144 = v140<<(uint(int32(8))%32) + v139
					v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
					v147 = v144 + v145
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v280 = v148 + v110
					v281 = v147
					v282 = v112
				case 7:
					v130 = v112
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					v280 = v131 + v110
					v281 = v133 + v111
					v282 = v130
				case 8:
					v125 = v112
					v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+8)))
					v130 = v126<<(uint(int32(8))%32) + v125
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					v280 = v131 + v110
					v281 = v133 + v111
					v282 = v130
				case 9:
					v120 = v112
					v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+9)))
					v125 = v121<<(uint(int32(16))%32) + v120
					v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+8)))
					v130 = v126<<(uint(int32(8))%32) + v125
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					v280 = v131 + v110
					v281 = v133 + v111
					v282 = v130
				case 10:
					v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
					v120 = v116<<(uint(int32(24))%32) + v112
					v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+9)))
					v125 = v121<<(uint(int32(16))%32) + v120
					v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+8)))
					v130 = v126<<(uint(int32(8))%32) + v125
					v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
					v280 = v131 + v110
					v281 = v133 + v111
					v282 = v130
				default:
					v280 = v110
					v281 = v111
					v282 = v112
				}
			}
			v285 = int32(14)
			v287 = v281 ^ v282 - base.I32_rotl(v281, v285)
			v291 = v287 ^ v280 - base.I32_rotl(v287, int32(11))
			v295 = v291 ^ v281 - base.I32_rotl(v291, int32(25))
			v299 = v295 ^ v287 - base.I32_rotl(v295, int32(16))
			v303 = v299 ^ v291 - base.I32_rotl(v299, int32(4))
			v307 = v303 ^ v295 - base.I32_rotl(v303, v285)
			F_pfree(m, v45)
			mBase = m.M
			v313 = m.ExcPending
			if v313 != 0 {
				return int64(0)
			} else {
				v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v314 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v317 = m.ExcPending
					if v317 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v307 ^ v299 - base.I32_rotl(v307, int32(24)))
					}
				} else {
					return base.I64_extend_i32_u(v307 ^ v299 - base.I32_rotl(v307, int32(24)))
				}
			}
		}
	}
}
