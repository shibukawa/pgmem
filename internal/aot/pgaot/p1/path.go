package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fetch_search_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	F_recomputeNamespacePath(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fetch_search_path[0])))
	if v8 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_AccessTempTableNamespace(m, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_search_path[1]))
	v18 = F_list_copy(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return v34
L9:
	;
	if l0|base.B2i32(v18 == int32(0)) != 0 {
		v34 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = v18
	goto L11
L11:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_search_path[2]))
	if v26 == v28 {
		v34 = v24
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v34 = int32(0)
	goto L8
L13:
	;
	v30 = F_list_delete_first(m, v24)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v30 != 0 {
		v24 = v30
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
}
func F_find_update_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v275 = int32(0)
	if l1 == l2 {
		v305 = v275
		goto L72
	} else {
		goto L73
	}
L2:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v76 <= int32(0) {
		goto L1
	} else {
		goto L16
	}
L3:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v60 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v60
	if l0 == v60 {
		goto L1
	} else {
		goto L15
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 < v18 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v25 = v17
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	goto L2
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v25<<(uint(int32(2))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+12)) = int64(2147483647)
	v40 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+9)) = uint8(v40)
	v43 = v25 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v25 = v43
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	goto L2
L16:
	;
	v85 = v76
	goto L17
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = int32(0)
	if v85 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L1
L19:
	;
	if v159 == int32(0) {
		goto L1
	} else {
		goto L44
	}
L20:
	;
	v105 = v92
	v107 = int32(0)
	v108 = v92
	goto L23
L21:
	;
	v139 = v92
	v142 = v92
	goto L22
L22:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v91+v142<<(uint(int32(2))%32))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+9)))
	if v151 != 0 {
		v159 = v139
		goto L19
	} else {
		goto L39
	}
L23:
	;
	v115 = v91 + v108<<(uint(int32(2))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+9)))
	if v117 != 0 {
		v121 = v105
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v85&int32(1) == int32(0) {
		v159 = v127
		goto L19
	} else {
		goto L38
	}
L25:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+9)))
	if v123 != 0 {
		v127 = v121
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if v118 <= v119 {
		v121 = v105
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v121 = v116
	goto L25
L30:
	;
	goto L29
L31:
	;
	v128 = int32(2)
	v129 = v108 + v128
	v131 = v107 + v128
	if v131 != v85&int32(2147483646) {
		v105 = v127
		v107 = v131
		v108 = v129
		goto L23
	} else {
		goto L37
	}
L32:
	;
	if v121 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	if v124 <= v125 {
		v127 = v121
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v127 = v122
	goto L31
L36:
	;
	goto L35
L37:
	;
	goto L24
L38:
	;
	v139 = v127
	v142 = v129
	goto L22
L39:
	;
	if v139 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	if v152 <= v153 {
		v159 = v139
		goto L19
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v159 = v150
	goto L19
L43:
	;
	goto L42
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	if v169 == int32(2147483647) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+9)) = uint8(v172)
	if l2 == v159 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v175 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v260 {
		v85 = v260
		goto L17
	} else {
		goto L71
	}
L48:
	;
	v178 = int32(0)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v179 <= v178 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v188 = v178
	goto L50
L50:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194+v188<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L47
L52:
	;
	v245 = v188 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v245 < v246 {
		v188 = v245
		goto L50
	} else {
		goto L70
	}
L53:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+8)))
	if v199 != 0 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v159)+12))
	v202 = v200 + int32(1)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)+12))
	if v202 < v203 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v198)+12)) = v202
	goto L52
L58:
	;
	goto L59
L59:
	;
	if v202 != v203 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v198)+16))
	if v208 == int32(0) {
		goto L52
	} else {
		goto L61
	}
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	if base.B2i32(v215 == int32(0))|base.B2i32(v215 != v218) != 0 {
		v236 = v215
		v237 = v218
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if int32(0) <= v236-v237 {
		goto L52
	} else {
		goto L69
	}
L63:
	;
	goto L62
L64:
	;
	v221 = v211
	v222 = v212
	goto L65
L65:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v226 == int32(0) {
		v236 = v226
		v237 = v225
		goto L63
	} else {
		goto L67
	}
L66:
	;
	v236 = v226
	v237 = v225
	goto L63
L67:
	;
	v229 = int32(1)
	if v226 == v225 {
		v221 = v221 + v229
		v222 = v222 + v229
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v198)+16)) = v159
	goto L52
L70:
	;
	goto L51
L71:
	;
	goto L18
L72:
	;
	return v305
L73:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v277&int32(1) == int32(0) {
		v305 = v275
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v284 = l2
	v286 = v275
	goto L75
L75:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v295 = F_lcons(m, v294, v286)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v305 = v295
	goto L72
L77:
	;
	return int32(0)
L78:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v284)+16))
	if v299 != l1 {
		v284 = v299
		v286 = v295
		goto L75
	} else {
		goto L79
	}
L79:
	;
	goto L76
}
func F_path_is_reparameterizable_by_child(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1089 int32
	_ = v1089
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1255 int32
	_ = v1255
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1332 int32
	_ = v1332
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1386 int32
	_ = v1386
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1506 int32
	_ = v1506
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1543 int32
	_ = v1543
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1593 int32
	_ = v1593
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1662 int32
	_ = v1662
	var v1668 int32
	_ = v1668
	var v1678 int32
	_ = v1678
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == int32(0) {
		v1678 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1678
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v11 = int32(0)
	if base.B2i32(v9 == v11)|base.B2i32(v10 == v11) != 0 {
		v56 = v11
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v56 == int32(0) {
		v1678 = v5
		goto L1
	} else {
		goto L16
	}
L4:
	;
	goto L3
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v21 < v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = v21
	goto L8
L7:
	;
	v24 = v22
	goto L8
L8:
	;
	if v24 <= int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v27 = int32(1)
	goto L11
L10:
	;
	v27 = v24
	goto L11
L11:
	;
	v28 = int32(8)
	v33 = int32(0)
	goto L12
L12:
	;
	v40 = v33 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10+v28+v40)))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v9+v28+v40)))
	v45 = v42 & v44
	v47 = base.B2i32(v45 != int32(0))
	if v45 != 0 {
		v56 = v47
		goto L4
	} else {
		goto L14
	}
L13:
	;
	v56 = v47
	goto L4
L14:
	;
	v49 = v33 + int32(1)
	if v49 != v27 {
		v33 = v49
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v60 - int32(279) {
	case 0, 1:
		goto L17
	default:
		v1678 = v59
		goto L1
	case 3:
		goto L27
	case 4:
		goto L26
	case 5:
		goto L25
	case 9:
		goto L24
	case 10:
		goto L23
	case 11:
		goto L21
	case 14:
		goto L20
	case 15:
		goto L19
	case 17:
		goto L18
	case 19, 20, 21:
		goto L22
	}
L17:
	;
	v1678 = int32(1)
	goto L1
L18:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1534 = int32(1)
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+16))
	if v1535 == int32(0) {
		v1662 = v1534
		goto L618
	} else {
		goto L619
	}
L19:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1396 = int32(1)
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+16))
	if v1397 == int32(0) {
		v1524 = v1396
		goto L562
	} else {
		goto L563
	}
L20:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1258 = int32(1)
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+16))
	if v1259 == int32(0) {
		v1386 = v1258
		goto L506
	} else {
		goto L507
	}
L21:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v1096 == int32(0) {
		goto L17
	} else {
		goto L442
	}
L22:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v821 = int32(1)
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v818)+16))
	if v822 == int32(0) {
		v949 = v821
		goto L331
	} else {
		goto L332
	}
L23:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v659 == int32(0) {
		goto L17
	} else {
		goto L267
	}
L24:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v519 == int32(0) {
		goto L17
	} else {
		goto L210
	}
L25:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v360 == int32(0) {
		goto L17
	} else {
		goto L147
	}
L26:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v201 == int32(0) {
		goto L17
	} else {
		goto L84
	}
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v66 = int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	if v67 == int32(0) {
		v194 = v66
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v200 != 0 {
		goto L17
	} else {
		goto L83
	}
L29:
	;
	v200 = v194
	goto L28
L30:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v72 = F_bms_overlap(m, v70, v71)
	mBase = m.M
	if v72 == int32(0) {
		v194 = v66
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v75 = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	switch v76 - int32(279) {
	case 0, 1:
		goto L32
	default:
		v194 = v75
		goto L29
	case 3:
		goto L42
	case 4:
		goto L41
	case 5:
		goto L40
	case 9:
		goto L39
	case 10:
		goto L38
	case 11:
		goto L36
	case 14:
		goto L35
	case 15:
		goto L34
	case 17:
		goto L33
	case 19, 20, 21:
		goto L37
	}
L32:
	;
	v194 = int32(1)
	goto L29
L33:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	v184 = F_path_is_reparameterizable_by_child(m, v183, l1)
	mBase = m.M
	if v184 == int32(0) {
		v194 = v75
		goto L29
	} else {
		goto L82
	}
L34:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	v182 = F_path_is_reparameterizable_by_child(m, v181, l1)
	mBase = m.M
	if v182 != 0 {
		goto L32
	} else {
		goto L81
	}
L35:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	v180 = F_path_is_reparameterizable_by_child(m, v179, l1)
	mBase = m.M
	if v180 != 0 {
		goto L32
	} else {
		goto L80
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v157 == int32(0) {
		goto L32
	} else {
		goto L72
	}
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v63)+80))
	v152 = F_path_is_reparameterizable_by_child(m, v151, l1)
	mBase = m.M
	if v152 == int32(0) {
		v194 = v75
		goto L29
	} else {
		goto L70
	}
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	if v129 == int32(0) {
		goto L32
	} else {
		goto L62
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v125 == int32(0) {
		goto L32
	} else {
		goto L60
	}
L40:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v103 == int32(0) {
		goto L32
	} else {
		goto L52
	}
L41:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v81 == int32(0) {
		goto L32
	} else {
		goto L44
	}
L42:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	v80 = F_path_is_reparameterizable_by_child(m, v79, l1)
	mBase = m.M
	if v80 != 0 {
		goto L32
	} else {
		goto L43
	}
L43:
	;
	v194 = v75
	goto L29
L44:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 <= v84 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v88 = v84
	goto L46
L46:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v88<<(uint(int32(2))%32))))
	v97 = F_path_is_reparameterizable_by_child(m, v96, l1)
	mBase = m.M
	if v97 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v200 = int32(0)
	goto L28
L48:
	;
	v99 = v88 + int32(1)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v99 < v100 {
		v88 = v99
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	goto L47
L51:
	;
	goto L32
L52:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 <= v106 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v110 = v106
	goto L54
L54:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114+v110<<(uint(int32(2))%32))))
	v119 = F_path_is_reparameterizable_by_child(m, v118, l1)
	mBase = m.M
	if v119 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v200 = int32(0)
	goto L28
L56:
	;
	v121 = v110 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v121 < v122 {
		v110 = v121
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L32
L60:
	;
	v128 = F_path_is_reparameterizable_by_child(m, v125, l1)
	mBase = m.M
	if v128 != 0 {
		goto L32
	} else {
		goto L61
	}
L61:
	;
	v194 = v75
	goto L29
L62:
	;
	v132 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v133 <= v132 {
		goto L32
	} else {
		goto L63
	}
L63:
	;
	v136 = v132
	goto L64
L64:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(int32(2))%32))))
	v145 = F_path_is_reparameterizable_by_child(m, v144, l1)
	mBase = m.M
	if v145 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v200 = int32(0)
	goto L28
L66:
	;
	v147 = v136 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	if v147 < v148 {
		v136 = v147
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	goto L32
L70:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v63)+84))
	v156 = F_path_is_reparameterizable_by_child(m, v155, l1)
	mBase = m.M
	if v156 != 0 {
		goto L32
	} else {
		goto L71
	}
L71:
	;
	v194 = v75
	goto L29
L72:
	;
	v160 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v161 <= v160 {
		goto L32
	} else {
		goto L73
	}
L73:
	;
	v164 = v160
	goto L74
L74:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v168+v164<<(uint(int32(2))%32))))
	v173 = F_path_is_reparameterizable_by_child(m, v172, l1)
	mBase = m.M
	if v173 != 0 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v200 = int32(0)
	goto L28
L76:
	;
	v175 = v164 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v175 < v176 {
		v164 = v175
		goto L74
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	goto L75
L79:
	;
	goto L32
L80:
	;
	v194 = v75
	goto L29
L81:
	;
	v194 = v75
	goto L29
L82:
	;
	goto L32
L83:
	;
	v1678 = v59
	goto L1
L84:
	;
	v204 = int32(0)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v205 <= v204 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	v208 = v204
	goto L86
L86:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v212+v208<<(uint(int32(2))%32))))
	v219 = int32(1)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	if v220 == int32(0) {
		v347 = v219
		goto L89
	} else {
		goto L90
	}
L87:
	;
	return int32(0)
L88:
	;
	if v353 != 0 {
		goto L143
	} else {
		goto L144
	}
L89:
	;
	v353 = v347
	goto L88
L90:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v225 = F_bms_overlap(m, v223, v224)
	mBase = m.M
	if v225 == int32(0) {
		v347 = v219
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v228 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	switch v229 - int32(279) {
	case 0, 1:
		goto L92
	default:
		v347 = v228
		goto L89
	case 3:
		goto L102
	case 4:
		goto L101
	case 5:
		goto L100
	case 9:
		goto L99
	case 10:
		goto L98
	case 11:
		goto L96
	case 14:
		goto L95
	case 15:
		goto L94
	case 17:
		goto L93
	case 19, 20, 21:
		goto L97
	}
L92:
	;
	v347 = int32(1)
	goto L89
L93:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v337 = F_path_is_reparameterizable_by_child(m, v336, l1)
	mBase = m.M
	if v337 == int32(0) {
		v347 = v228
		goto L89
	} else {
		goto L142
	}
L94:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v335 = F_path_is_reparameterizable_by_child(m, v334, l1)
	mBase = m.M
	if v335 != 0 {
		goto L92
	} else {
		goto L141
	}
L95:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v333 = F_path_is_reparameterizable_by_child(m, v332, l1)
	mBase = m.M
	if v333 != 0 {
		goto L92
	} else {
		goto L140
	}
L96:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	if v310 == int32(0) {
		goto L92
	} else {
		goto L132
	}
L97:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v216)+80))
	v305 = F_path_is_reparameterizable_by_child(m, v304, l1)
	mBase = m.M
	if v305 == int32(0) {
		v347 = v228
		goto L89
	} else {
		goto L130
	}
L98:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v216)+76))
	if v282 == int32(0) {
		goto L92
	} else {
		goto L122
	}
L99:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	if v278 == int32(0) {
		goto L92
	} else {
		goto L120
	}
L100:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	if v256 == int32(0) {
		goto L92
	} else {
		goto L112
	}
L101:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	if v234 == int32(0) {
		goto L92
	} else {
		goto L104
	}
L102:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v216)+72))
	v233 = F_path_is_reparameterizable_by_child(m, v232, l1)
	mBase = m.M
	if v233 != 0 {
		goto L92
	} else {
		goto L103
	}
L103:
	;
	v347 = v228
	goto L89
L104:
	;
	v237 = int32(0)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v238 <= v237 {
		goto L92
	} else {
		goto L105
	}
L105:
	;
	v241 = v237
	goto L106
L106:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v245+v241<<(uint(int32(2))%32))))
	v250 = F_path_is_reparameterizable_by_child(m, v249, l1)
	mBase = m.M
	if v250 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v353 = int32(0)
	goto L88
L108:
	;
	v252 = v241 + int32(1)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	if v252 < v253 {
		v241 = v252
		goto L106
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	goto L107
L111:
	;
	goto L92
L112:
	;
	v259 = int32(0)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v260 <= v259 {
		goto L92
	} else {
		goto L113
	}
L113:
	;
	v263 = v259
	goto L114
L114:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v256)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v267+v263<<(uint(int32(2))%32))))
	v272 = F_path_is_reparameterizable_by_child(m, v271, l1)
	mBase = m.M
	if v272 != 0 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v353 = int32(0)
	goto L88
L116:
	;
	v274 = v263 + int32(1)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v256)+4))
	if v274 < v275 {
		v263 = v274
		goto L114
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	goto L92
L120:
	;
	v281 = F_path_is_reparameterizable_by_child(m, v278, l1)
	mBase = m.M
	if v281 != 0 {
		goto L92
	} else {
		goto L121
	}
L121:
	;
	v347 = v228
	goto L89
L122:
	;
	v285 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v286 <= v285 {
		goto L92
	} else {
		goto L123
	}
L123:
	;
	v289 = v285
	goto L124
L124:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v289<<(uint(int32(2))%32))))
	v298 = F_path_is_reparameterizable_by_child(m, v297, l1)
	mBase = m.M
	if v298 != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v353 = int32(0)
	goto L88
L126:
	;
	v300 = v289 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v300 < v301 {
		v289 = v300
		goto L124
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	goto L125
L129:
	;
	goto L92
L130:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v216)+84))
	v309 = F_path_is_reparameterizable_by_child(m, v308, l1)
	mBase = m.M
	if v309 != 0 {
		goto L92
	} else {
		goto L131
	}
L131:
	;
	v347 = v228
	goto L89
L132:
	;
	v313 = int32(0)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	if v314 <= v313 {
		goto L92
	} else {
		goto L133
	}
L133:
	;
	v317 = v313
	goto L134
L134:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321+v317<<(uint(int32(2))%32))))
	v326 = F_path_is_reparameterizable_by_child(m, v325, l1)
	mBase = m.M
	if v326 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	v353 = int32(0)
	goto L88
L136:
	;
	v328 = v317 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	if v328 < v329 {
		v317 = v328
		goto L134
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	goto L135
L139:
	;
	goto L92
L140:
	;
	v347 = v228
	goto L89
L141:
	;
	v347 = v228
	goto L89
L142:
	;
	goto L92
L143:
	;
	v355 = v208 + int32(1)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v355 < v356 {
		v208 = v355
		goto L86
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	goto L87
L146:
	;
	goto L17
L147:
	;
	v363 = int32(0)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	if v364 <= v363 {
		goto L17
	} else {
		goto L148
	}
L148:
	;
	v367 = v363
	goto L149
L149:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v360)+12))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v371+v367<<(uint(int32(2))%32))))
	v378 = int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)+16))
	if v379 == int32(0) {
		v506 = v378
		goto L152
	} else {
		goto L153
	}
L150:
	;
	return int32(0)
L151:
	;
	if v512 != 0 {
		goto L206
	} else {
		goto L207
	}
L152:
	;
	v512 = v506
	goto L151
L153:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v384 = F_bms_overlap(m, v382, v383)
	mBase = m.M
	if v384 == int32(0) {
		v506 = v378
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v387 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	switch v388 - int32(279) {
	case 0, 1:
		goto L155
	default:
		v506 = v387
		goto L152
	case 3:
		goto L165
	case 4:
		goto L164
	case 5:
		goto L163
	case 9:
		goto L162
	case 10:
		goto L161
	case 11:
		goto L159
	case 14:
		goto L158
	case 15:
		goto L157
	case 17:
		goto L156
	case 19, 20, 21:
		goto L160
	}
L155:
	;
	v506 = int32(1)
	goto L152
L156:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	v496 = F_path_is_reparameterizable_by_child(m, v495, l1)
	mBase = m.M
	if v496 == int32(0) {
		v506 = v387
		goto L152
	} else {
		goto L205
	}
L157:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	v494 = F_path_is_reparameterizable_by_child(m, v493, l1)
	mBase = m.M
	if v494 != 0 {
		goto L155
	} else {
		goto L204
	}
L158:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	v492 = F_path_is_reparameterizable_by_child(m, v491, l1)
	mBase = m.M
	if v492 != 0 {
		goto L155
	} else {
		goto L203
	}
L159:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	if v469 == int32(0) {
		goto L155
	} else {
		goto L195
	}
L160:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v375)+80))
	v464 = F_path_is_reparameterizable_by_child(m, v463, l1)
	mBase = m.M
	if v464 == int32(0) {
		v506 = v387
		goto L152
	} else {
		goto L193
	}
L161:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v375)+76))
	if v441 == int32(0) {
		goto L155
	} else {
		goto L185
	}
L162:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	if v437 == int32(0) {
		goto L155
	} else {
		goto L183
	}
L163:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	if v415 == int32(0) {
		goto L155
	} else {
		goto L175
	}
L164:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	if v393 == int32(0) {
		goto L155
	} else {
		goto L167
	}
L165:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v375)+72))
	v392 = F_path_is_reparameterizable_by_child(m, v391, l1)
	mBase = m.M
	if v392 != 0 {
		goto L155
	} else {
		goto L166
	}
L166:
	;
	v506 = v387
	goto L152
L167:
	;
	v396 = int32(0)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v397 <= v396 {
		goto L155
	} else {
		goto L168
	}
L168:
	;
	v400 = v396
	goto L169
L169:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v393)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v404+v400<<(uint(int32(2))%32))))
	v409 = F_path_is_reparameterizable_by_child(m, v408, l1)
	mBase = m.M
	if v409 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v512 = int32(0)
	goto L151
L171:
	;
	v411 = v400 + int32(1)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v393)+4))
	if v411 < v412 {
		v400 = v411
		goto L169
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	goto L170
L174:
	;
	goto L155
L175:
	;
	v418 = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if v419 <= v418 {
		goto L155
	} else {
		goto L176
	}
L176:
	;
	v422 = v418
	goto L177
L177:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v415)+12))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426+v422<<(uint(int32(2))%32))))
	v431 = F_path_is_reparameterizable_by_child(m, v430, l1)
	mBase = m.M
	if v431 != 0 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v512 = int32(0)
	goto L151
L179:
	;
	v433 = v422 + int32(1)
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	if v433 < v434 {
		v422 = v433
		goto L177
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	goto L178
L182:
	;
	goto L155
L183:
	;
	v440 = F_path_is_reparameterizable_by_child(m, v437, l1)
	mBase = m.M
	if v440 != 0 {
		goto L155
	} else {
		goto L184
	}
L184:
	;
	v506 = v387
	goto L152
L185:
	;
	v444 = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v445 <= v444 {
		goto L155
	} else {
		goto L186
	}
L186:
	;
	v448 = v444
	goto L187
L187:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v441)+12))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452+v448<<(uint(int32(2))%32))))
	v457 = F_path_is_reparameterizable_by_child(m, v456, l1)
	mBase = m.M
	if v457 != 0 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	v512 = int32(0)
	goto L151
L189:
	;
	v459 = v448 + int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	if v459 < v460 {
		v448 = v459
		goto L187
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	goto L188
L192:
	;
	goto L155
L193:
	;
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v375)+84))
	v468 = F_path_is_reparameterizable_by_child(m, v467, l1)
	mBase = m.M
	if v468 != 0 {
		goto L155
	} else {
		goto L194
	}
L194:
	;
	v506 = v387
	goto L152
L195:
	;
	v472 = int32(0)
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v473 <= v472 {
		goto L155
	} else {
		goto L196
	}
L196:
	;
	v476 = v472
	goto L197
L197:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v469)+12))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480+v476<<(uint(int32(2))%32))))
	v485 = F_path_is_reparameterizable_by_child(m, v484, l1)
	mBase = m.M
	if v485 != 0 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v512 = int32(0)
	goto L151
L199:
	;
	v487 = v476 + int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v469)+4))
	if v487 < v488 {
		v476 = v487
		goto L197
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	goto L198
L202:
	;
	goto L155
L203:
	;
	v506 = v387
	goto L152
L204:
	;
	v506 = v387
	goto L152
L205:
	;
	goto L155
L206:
	;
	v514 = v367 + int32(1)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v360)+4))
	if v514 < v515 {
		v367 = v514
		goto L149
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	goto L150
L209:
	;
	goto L17
L210:
	;
	v524 = int32(1)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v519)+16))
	if v525 == int32(0) {
		v652 = v524
		goto L212
	} else {
		goto L213
	}
L211:
	;
	if v658 != 0 {
		goto L17
	} else {
		goto L266
	}
L212:
	;
	v658 = v652
	goto L211
L213:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v525)+4))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v530 = F_bms_overlap(m, v528, v529)
	mBase = m.M
	if v530 == int32(0) {
		v652 = v524
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v533 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	switch v534 - int32(279) {
	case 0, 1:
		goto L215
	default:
		v652 = v533
		goto L212
	case 3:
		goto L225
	case 4:
		goto L224
	case 5:
		goto L223
	case 9:
		goto L222
	case 10:
		goto L221
	case 11:
		goto L219
	case 14:
		goto L218
	case 15:
		goto L217
	case 17:
		goto L216
	case 19, 20, 21:
		goto L220
	}
L215:
	;
	v652 = int32(1)
	goto L212
L216:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	v642 = F_path_is_reparameterizable_by_child(m, v641, l1)
	mBase = m.M
	if v642 == int32(0) {
		v652 = v533
		goto L212
	} else {
		goto L265
	}
L217:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	v640 = F_path_is_reparameterizable_by_child(m, v639, l1)
	mBase = m.M
	if v640 != 0 {
		goto L215
	} else {
		goto L264
	}
L218:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	v638 = F_path_is_reparameterizable_by_child(m, v637, l1)
	mBase = m.M
	if v638 != 0 {
		goto L215
	} else {
		goto L263
	}
L219:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	if v615 == int32(0) {
		goto L215
	} else {
		goto L255
	}
L220:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v519)+80))
	v610 = F_path_is_reparameterizable_by_child(m, v609, l1)
	mBase = m.M
	if v610 == int32(0) {
		v652 = v533
		goto L212
	} else {
		goto L253
	}
L221:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v519)+76))
	if v587 == int32(0) {
		goto L215
	} else {
		goto L245
	}
L222:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	if v583 == int32(0) {
		goto L215
	} else {
		goto L243
	}
L223:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	if v561 == int32(0) {
		goto L215
	} else {
		goto L235
	}
L224:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	if v539 == int32(0) {
		goto L215
	} else {
		goto L227
	}
L225:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v519)+72))
	v538 = F_path_is_reparameterizable_by_child(m, v537, l1)
	mBase = m.M
	if v538 != 0 {
		goto L215
	} else {
		goto L226
	}
L226:
	;
	v652 = v533
	goto L212
L227:
	;
	v542 = int32(0)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v543 <= v542 {
		goto L215
	} else {
		goto L228
	}
L228:
	;
	v546 = v542
	goto L229
L229:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v539)+12))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550+v546<<(uint(int32(2))%32))))
	v555 = F_path_is_reparameterizable_by_child(m, v554, l1)
	mBase = m.M
	if v555 != 0 {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	v658 = int32(0)
	goto L211
L231:
	;
	v557 = v546 + int32(1)
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v539)+4))
	if v557 < v558 {
		v546 = v557
		goto L229
	} else {
		goto L234
	}
L232:
	;
	goto L233
L233:
	;
	goto L230
L234:
	;
	goto L215
L235:
	;
	v564 = int32(0)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v565 <= v564 {
		goto L215
	} else {
		goto L236
	}
L236:
	;
	v568 = v564
	goto L237
L237:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v572+v568<<(uint(int32(2))%32))))
	v577 = F_path_is_reparameterizable_by_child(m, v576, l1)
	mBase = m.M
	if v577 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v658 = int32(0)
	goto L211
L239:
	;
	v579 = v568 + int32(1)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v561)+4))
	if v579 < v580 {
		v568 = v579
		goto L237
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	goto L238
L242:
	;
	goto L215
L243:
	;
	v586 = F_path_is_reparameterizable_by_child(m, v583, l1)
	mBase = m.M
	if v586 != 0 {
		goto L215
	} else {
		goto L244
	}
L244:
	;
	v652 = v533
	goto L212
L245:
	;
	v590 = int32(0)
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v591 <= v590 {
		goto L215
	} else {
		goto L246
	}
L246:
	;
	v594 = v590
	goto L247
L247:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v587)+12))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v598+v594<<(uint(int32(2))%32))))
	v603 = F_path_is_reparameterizable_by_child(m, v602, l1)
	mBase = m.M
	if v603 != 0 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v658 = int32(0)
	goto L211
L249:
	;
	v605 = v594 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v587)+4))
	if v605 < v606 {
		v594 = v605
		goto L247
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	goto L248
L252:
	;
	goto L215
L253:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v519)+84))
	v614 = F_path_is_reparameterizable_by_child(m, v613, l1)
	mBase = m.M
	if v614 != 0 {
		goto L215
	} else {
		goto L254
	}
L254:
	;
	v652 = v533
	goto L212
L255:
	;
	v618 = int32(0)
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if v619 <= v618 {
		goto L215
	} else {
		goto L256
	}
L256:
	;
	v622 = v618
	goto L257
L257:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v615)+12))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v626+v622<<(uint(int32(2))%32))))
	v631 = F_path_is_reparameterizable_by_child(m, v630, l1)
	mBase = m.M
	if v631 != 0 {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v658 = int32(0)
	goto L211
L259:
	;
	v633 = v622 + int32(1)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v615)+4))
	if v633 < v634 {
		v622 = v633
		goto L257
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	goto L258
L262:
	;
	goto L215
L263:
	;
	v652 = v533
	goto L212
L264:
	;
	v652 = v533
	goto L212
L265:
	;
	goto L215
L266:
	;
	v1678 = v59
	goto L1
L267:
	;
	v662 = int32(0)
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v663 <= v662 {
		goto L17
	} else {
		goto L268
	}
L268:
	;
	v666 = v662
	goto L269
L269:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v659)+12))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+v666<<(uint(int32(2))%32))))
	v677 = int32(1)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v674)+16))
	if v678 == int32(0) {
		v805 = v677
		goto L272
	} else {
		goto L273
	}
L270:
	;
	return int32(0)
L271:
	;
	if v811 != 0 {
		goto L326
	} else {
		goto L327
	}
L272:
	;
	v811 = v805
	goto L271
L273:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v678)+4))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v683 = F_bms_overlap(m, v681, v682)
	mBase = m.M
	if v683 == int32(0) {
		v805 = v677
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v686 = int32(0)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	switch v687 - int32(279) {
	case 0, 1:
		goto L275
	default:
		v805 = v686
		goto L272
	case 3:
		goto L285
	case 4:
		goto L284
	case 5:
		goto L283
	case 9:
		goto L282
	case 10:
		goto L281
	case 11:
		goto L279
	case 14:
		goto L278
	case 15:
		goto L277
	case 17:
		goto L276
	case 19, 20, 21:
		goto L280
	}
L275:
	;
	v805 = int32(1)
	goto L272
L276:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	v795 = F_path_is_reparameterizable_by_child(m, v794, l1)
	mBase = m.M
	if v795 == int32(0) {
		v805 = v686
		goto L272
	} else {
		goto L325
	}
L277:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	v793 = F_path_is_reparameterizable_by_child(m, v792, l1)
	mBase = m.M
	if v793 != 0 {
		goto L275
	} else {
		goto L324
	}
L278:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	v791 = F_path_is_reparameterizable_by_child(m, v790, l1)
	mBase = m.M
	if v791 != 0 {
		goto L275
	} else {
		goto L323
	}
L279:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	if v768 == int32(0) {
		goto L275
	} else {
		goto L315
	}
L280:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v674)+80))
	v763 = F_path_is_reparameterizable_by_child(m, v762, l1)
	mBase = m.M
	if v763 == int32(0) {
		v805 = v686
		goto L272
	} else {
		goto L313
	}
L281:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v674)+76))
	if v740 == int32(0) {
		goto L275
	} else {
		goto L305
	}
L282:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	if v736 == int32(0) {
		goto L275
	} else {
		goto L303
	}
L283:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	if v714 == int32(0) {
		goto L275
	} else {
		goto L295
	}
L284:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	if v692 == int32(0) {
		goto L275
	} else {
		goto L287
	}
L285:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v674)+72))
	v691 = F_path_is_reparameterizable_by_child(m, v690, l1)
	mBase = m.M
	if v691 != 0 {
		goto L275
	} else {
		goto L286
	}
L286:
	;
	v805 = v686
	goto L272
L287:
	;
	v695 = int32(0)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	if v696 <= v695 {
		goto L275
	} else {
		goto L288
	}
L288:
	;
	v699 = v695
	goto L289
L289:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v692)+12))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v703+v699<<(uint(int32(2))%32))))
	v708 = F_path_is_reparameterizable_by_child(m, v707, l1)
	mBase = m.M
	if v708 != 0 {
		goto L291
	} else {
		goto L292
	}
L290:
	;
	v811 = int32(0)
	goto L271
L291:
	;
	v710 = v699 + int32(1)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v692)+4))
	if v710 < v711 {
		v699 = v710
		goto L289
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	goto L290
L294:
	;
	goto L275
L295:
	;
	v717 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v718 <= v717 {
		goto L275
	} else {
		goto L296
	}
L296:
	;
	v721 = v717
	goto L297
L297:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v725+v721<<(uint(int32(2))%32))))
	v730 = F_path_is_reparameterizable_by_child(m, v729, l1)
	mBase = m.M
	if v730 != 0 {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	v811 = int32(0)
	goto L271
L299:
	;
	v732 = v721 + int32(1)
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v732 < v733 {
		v721 = v732
		goto L297
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	goto L298
L302:
	;
	goto L275
L303:
	;
	v739 = F_path_is_reparameterizable_by_child(m, v736, l1)
	mBase = m.M
	if v739 != 0 {
		goto L275
	} else {
		goto L304
	}
L304:
	;
	v805 = v686
	goto L272
L305:
	;
	v743 = int32(0)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	if v744 <= v743 {
		goto L275
	} else {
		goto L306
	}
L306:
	;
	v747 = v743
	goto L307
L307:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v740)+12))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v751+v747<<(uint(int32(2))%32))))
	v756 = F_path_is_reparameterizable_by_child(m, v755, l1)
	mBase = m.M
	if v756 != 0 {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v811 = int32(0)
	goto L271
L309:
	;
	v758 = v747 + int32(1)
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	if v758 < v759 {
		v747 = v758
		goto L307
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	goto L308
L312:
	;
	goto L275
L313:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v674)+84))
	v767 = F_path_is_reparameterizable_by_child(m, v766, l1)
	mBase = m.M
	if v767 != 0 {
		goto L275
	} else {
		goto L314
	}
L314:
	;
	v805 = v686
	goto L272
L315:
	;
	v771 = int32(0)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	if v772 <= v771 {
		goto L275
	} else {
		goto L316
	}
L316:
	;
	v775 = v771
	goto L317
L317:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v768)+12))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v779+v775<<(uint(int32(2))%32))))
	v784 = F_path_is_reparameterizable_by_child(m, v783, l1)
	mBase = m.M
	if v784 != 0 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v811 = int32(0)
	goto L271
L319:
	;
	v786 = v775 + int32(1)
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v768)+4))
	if v786 < v787 {
		v775 = v786
		goto L317
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	goto L318
L322:
	;
	goto L275
L323:
	;
	v805 = v686
	goto L272
L324:
	;
	v805 = v686
	goto L272
L325:
	;
	goto L275
L326:
	;
	v813 = v666 + int32(1)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	if v813 < v814 {
		v666 = v813
		goto L269
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	goto L270
L329:
	;
	goto L17
L330:
	;
	if v955 == int32(0) {
		v1678 = v59
		goto L1
	} else {
		goto L385
	}
L331:
	;
	v955 = v949
	goto L330
L332:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v822)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v827 = F_bms_overlap(m, v825, v826)
	mBase = m.M
	if v827 == int32(0) {
		v949 = v821
		goto L331
	} else {
		goto L333
	}
L333:
	;
	v830 = int32(0)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	switch v831 - int32(279) {
	case 0, 1:
		goto L334
	default:
		v949 = v830
		goto L331
	case 3:
		goto L344
	case 4:
		goto L343
	case 5:
		goto L342
	case 9:
		goto L341
	case 10:
		goto L340
	case 11:
		goto L338
	case 14:
		goto L337
	case 15:
		goto L336
	case 17:
		goto L335
	case 19, 20, 21:
		goto L339
	}
L334:
	;
	v949 = int32(1)
	goto L331
L335:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	v939 = F_path_is_reparameterizable_by_child(m, v938, l1)
	mBase = m.M
	if v939 == int32(0) {
		v949 = v830
		goto L331
	} else {
		goto L384
	}
L336:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	v937 = F_path_is_reparameterizable_by_child(m, v936, l1)
	mBase = m.M
	if v937 != 0 {
		goto L334
	} else {
		goto L383
	}
L337:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	v935 = F_path_is_reparameterizable_by_child(m, v934, l1)
	mBase = m.M
	if v935 != 0 {
		goto L334
	} else {
		goto L382
	}
L338:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	if v912 == int32(0) {
		goto L334
	} else {
		goto L374
	}
L339:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v818)+80))
	v907 = F_path_is_reparameterizable_by_child(m, v906, l1)
	mBase = m.M
	if v907 == int32(0) {
		v949 = v830
		goto L331
	} else {
		goto L372
	}
L340:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v818)+76))
	if v884 == int32(0) {
		goto L334
	} else {
		goto L364
	}
L341:
	;
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	if v880 == int32(0) {
		goto L334
	} else {
		goto L362
	}
L342:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	if v858 == int32(0) {
		goto L334
	} else {
		goto L354
	}
L343:
	;
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	if v836 == int32(0) {
		goto L334
	} else {
		goto L346
	}
L344:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v818)+72))
	v835 = F_path_is_reparameterizable_by_child(m, v834, l1)
	mBase = m.M
	if v835 != 0 {
		goto L334
	} else {
		goto L345
	}
L345:
	;
	v949 = v830
	goto L331
L346:
	;
	v839 = int32(0)
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	if v840 <= v839 {
		goto L334
	} else {
		goto L347
	}
L347:
	;
	v843 = v839
	goto L348
L348:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v836)+12))
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v847+v843<<(uint(int32(2))%32))))
	v852 = F_path_is_reparameterizable_by_child(m, v851, l1)
	mBase = m.M
	if v852 != 0 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	v955 = int32(0)
	goto L330
L350:
	;
	v854 = v843 + int32(1)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	if v854 < v855 {
		v843 = v854
		goto L348
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	goto L349
L353:
	;
	goto L334
L354:
	;
	v861 = int32(0)
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v862 <= v861 {
		goto L334
	} else {
		goto L355
	}
L355:
	;
	v865 = v861
	goto L356
L356:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v858)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v869+v865<<(uint(int32(2))%32))))
	v874 = F_path_is_reparameterizable_by_child(m, v873, l1)
	mBase = m.M
	if v874 != 0 {
		goto L358
	} else {
		goto L359
	}
L357:
	;
	v955 = int32(0)
	goto L330
L358:
	;
	v876 = v865 + int32(1)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v858)+4))
	if v876 < v877 {
		v865 = v876
		goto L356
	} else {
		goto L361
	}
L359:
	;
	goto L360
L360:
	;
	goto L357
L361:
	;
	goto L334
L362:
	;
	v883 = F_path_is_reparameterizable_by_child(m, v880, l1)
	mBase = m.M
	if v883 != 0 {
		goto L334
	} else {
		goto L363
	}
L363:
	;
	v949 = v830
	goto L331
L364:
	;
	v887 = int32(0)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v884)+4))
	if v888 <= v887 {
		goto L334
	} else {
		goto L365
	}
L365:
	;
	v891 = v887
	goto L366
L366:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v884)+12))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895+v891<<(uint(int32(2))%32))))
	v900 = F_path_is_reparameterizable_by_child(m, v899, l1)
	mBase = m.M
	if v900 != 0 {
		goto L368
	} else {
		goto L369
	}
L367:
	;
	v955 = int32(0)
	goto L330
L368:
	;
	v902 = v891 + int32(1)
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v884)+4))
	if v902 < v903 {
		v891 = v902
		goto L366
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	goto L367
L371:
	;
	goto L334
L372:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v818)+84))
	v911 = F_path_is_reparameterizable_by_child(m, v910, l1)
	mBase = m.M
	if v911 != 0 {
		goto L334
	} else {
		goto L373
	}
L373:
	;
	v949 = v830
	goto L331
L374:
	;
	v915 = int32(0)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	if v916 <= v915 {
		goto L334
	} else {
		goto L375
	}
L375:
	;
	v919 = v915
	goto L376
L376:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v912)+12))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v923+v919<<(uint(int32(2))%32))))
	v928 = F_path_is_reparameterizable_by_child(m, v927, l1)
	mBase = m.M
	if v928 != 0 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	v955 = int32(0)
	goto L330
L378:
	;
	v930 = v919 + int32(1)
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	if v930 < v931 {
		v919 = v930
		goto L376
	} else {
		goto L381
	}
L379:
	;
	goto L380
L380:
	;
	goto L377
L381:
	;
	goto L334
L382:
	;
	v949 = v830
	goto L331
L383:
	;
	v949 = v830
	goto L331
L384:
	;
	goto L334
L385:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v961 = int32(1)
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v958)+16))
	if v962 == int32(0) {
		v1089 = v961
		goto L387
	} else {
		goto L388
	}
L386:
	;
	if v1095 != 0 {
		goto L17
	} else {
		goto L441
	}
L387:
	;
	v1095 = v1089
	goto L386
L388:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v962)+4))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v967 = F_bms_overlap(m, v965, v966)
	mBase = m.M
	if v967 == int32(0) {
		v1089 = v961
		goto L387
	} else {
		goto L389
	}
L389:
	;
	v970 = int32(0)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v958)))
	switch v971 - int32(279) {
	case 0, 1:
		goto L390
	default:
		v1089 = v970
		goto L387
	case 3:
		goto L400
	case 4:
		goto L399
	case 5:
		goto L398
	case 9:
		goto L397
	case 10:
		goto L396
	case 11:
		goto L394
	case 14:
		goto L393
	case 15:
		goto L392
	case 17:
		goto L391
	case 19, 20, 21:
		goto L395
	}
L390:
	;
	v1089 = int32(1)
	goto L387
L391:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	v1079 = F_path_is_reparameterizable_by_child(m, v1078, l1)
	mBase = m.M
	if v1079 == int32(0) {
		v1089 = v970
		goto L387
	} else {
		goto L440
	}
L392:
	;
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	v1077 = F_path_is_reparameterizable_by_child(m, v1076, l1)
	mBase = m.M
	if v1077 != 0 {
		goto L390
	} else {
		goto L439
	}
L393:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	v1075 = F_path_is_reparameterizable_by_child(m, v1074, l1)
	mBase = m.M
	if v1075 != 0 {
		goto L390
	} else {
		goto L438
	}
L394:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	if v1052 == int32(0) {
		goto L390
	} else {
		goto L430
	}
L395:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v958)+80))
	v1047 = F_path_is_reparameterizable_by_child(m, v1046, l1)
	mBase = m.M
	if v1047 == int32(0) {
		v1089 = v970
		goto L387
	} else {
		goto L428
	}
L396:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v958)+76))
	if v1024 == int32(0) {
		goto L390
	} else {
		goto L420
	}
L397:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	if v1020 == int32(0) {
		goto L390
	} else {
		goto L418
	}
L398:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	if v998 == int32(0) {
		goto L390
	} else {
		goto L410
	}
L399:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	if v976 == int32(0) {
		goto L390
	} else {
		goto L402
	}
L400:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v958)+72))
	v975 = F_path_is_reparameterizable_by_child(m, v974, l1)
	mBase = m.M
	if v975 != 0 {
		goto L390
	} else {
		goto L401
	}
L401:
	;
	v1089 = v970
	goto L387
L402:
	;
	v979 = int32(0)
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	if v980 <= v979 {
		goto L390
	} else {
		goto L403
	}
L403:
	;
	v983 = v979
	goto L404
L404:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v976)+12))
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v987+v983<<(uint(int32(2))%32))))
	v992 = F_path_is_reparameterizable_by_child(m, v991, l1)
	mBase = m.M
	if v992 != 0 {
		goto L406
	} else {
		goto L407
	}
L405:
	;
	v1095 = int32(0)
	goto L386
L406:
	;
	v994 = v983 + int32(1)
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v976)+4))
	if v994 < v995 {
		v983 = v994
		goto L404
	} else {
		goto L409
	}
L407:
	;
	goto L408
L408:
	;
	goto L405
L409:
	;
	goto L390
L410:
	;
	v1001 = int32(0)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v998)+4))
	if v1002 <= v1001 {
		goto L390
	} else {
		goto L411
	}
L411:
	;
	v1005 = v1001
	goto L412
L412:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v998)+12))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1009+v1005<<(uint(int32(2))%32))))
	v1014 = F_path_is_reparameterizable_by_child(m, v1013, l1)
	mBase = m.M
	if v1014 != 0 {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	v1095 = int32(0)
	goto L386
L414:
	;
	v1016 = v1005 + int32(1)
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v998)+4))
	if v1016 < v1017 {
		v1005 = v1016
		goto L412
	} else {
		goto L417
	}
L415:
	;
	goto L416
L416:
	;
	goto L413
L417:
	;
	goto L390
L418:
	;
	v1023 = F_path_is_reparameterizable_by_child(m, v1020, l1)
	mBase = m.M
	if v1023 != 0 {
		goto L390
	} else {
		goto L419
	}
L419:
	;
	v1089 = v970
	goto L387
L420:
	;
	v1027 = int32(0)
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	if v1028 <= v1027 {
		goto L390
	} else {
		goto L421
	}
L421:
	;
	v1031 = v1027
	goto L422
L422:
	;
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+12))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1035+v1031<<(uint(int32(2))%32))))
	v1040 = F_path_is_reparameterizable_by_child(m, v1039, l1)
	mBase = m.M
	if v1040 != 0 {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	v1095 = int32(0)
	goto L386
L424:
	;
	v1042 = v1031 + int32(1)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1024)+4))
	if v1042 < v1043 {
		v1031 = v1042
		goto L422
	} else {
		goto L427
	}
L425:
	;
	goto L426
L426:
	;
	goto L423
L427:
	;
	goto L390
L428:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v958)+84))
	v1051 = F_path_is_reparameterizable_by_child(m, v1050, l1)
	mBase = m.M
	if v1051 != 0 {
		goto L390
	} else {
		goto L429
	}
L429:
	;
	v1089 = v970
	goto L387
L430:
	;
	v1055 = int32(0)
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+4))
	if v1056 <= v1055 {
		goto L390
	} else {
		goto L431
	}
L431:
	;
	v1059 = v1055
	goto L432
L432:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+12))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1063+v1059<<(uint(int32(2))%32))))
	v1068 = F_path_is_reparameterizable_by_child(m, v1067, l1)
	mBase = m.M
	if v1068 != 0 {
		goto L434
	} else {
		goto L435
	}
L433:
	;
	v1095 = int32(0)
	goto L386
L434:
	;
	v1070 = v1059 + int32(1)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1052)+4))
	if v1070 < v1071 {
		v1059 = v1070
		goto L432
	} else {
		goto L437
	}
L435:
	;
	goto L436
L436:
	;
	goto L433
L437:
	;
	goto L390
L438:
	;
	v1089 = v970
	goto L387
L439:
	;
	v1089 = v970
	goto L387
L440:
	;
	goto L390
L441:
	;
	v1678 = v59
	goto L1
L442:
	;
	v1099 = int32(0)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	if v1100 <= v1099 {
		goto L17
	} else {
		goto L443
	}
L443:
	;
	v1103 = v1099
	goto L444
L444:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+12))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1107+v1103<<(uint(int32(2))%32))))
	v1114 = int32(1)
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+16))
	if v1115 == int32(0) {
		v1242 = v1114
		goto L447
	} else {
		goto L448
	}
L445:
	;
	return int32(0)
L446:
	;
	if v1248 != 0 {
		goto L501
	} else {
		goto L502
	}
L447:
	;
	v1248 = v1242
	goto L446
L448:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v1115)+4))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1120 = F_bms_overlap(m, v1118, v1119)
	mBase = m.M
	if v1120 == int32(0) {
		v1242 = v1114
		goto L447
	} else {
		goto L449
	}
L449:
	;
	v1123 = int32(0)
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	switch v1124 - int32(279) {
	case 0, 1:
		goto L450
	default:
		v1242 = v1123
		goto L447
	case 3:
		goto L460
	case 4:
		goto L459
	case 5:
		goto L458
	case 9:
		goto L457
	case 10:
		goto L456
	case 11:
		goto L454
	case 14:
		goto L453
	case 15:
		goto L452
	case 17:
		goto L451
	case 19, 20, 21:
		goto L455
	}
L450:
	;
	v1242 = int32(1)
	goto L447
L451:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	v1232 = F_path_is_reparameterizable_by_child(m, v1231, l1)
	mBase = m.M
	if v1232 == int32(0) {
		v1242 = v1123
		goto L447
	} else {
		goto L500
	}
L452:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	v1230 = F_path_is_reparameterizable_by_child(m, v1229, l1)
	mBase = m.M
	if v1230 != 0 {
		goto L450
	} else {
		goto L499
	}
L453:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	v1228 = F_path_is_reparameterizable_by_child(m, v1227, l1)
	mBase = m.M
	if v1228 != 0 {
		goto L450
	} else {
		goto L498
	}
L454:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	if v1205 == int32(0) {
		goto L450
	} else {
		goto L490
	}
L455:
	;
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+80))
	v1200 = F_path_is_reparameterizable_by_child(m, v1199, l1)
	mBase = m.M
	if v1200 == int32(0) {
		v1242 = v1123
		goto L447
	} else {
		goto L488
	}
L456:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+76))
	if v1177 == int32(0) {
		goto L450
	} else {
		goto L480
	}
L457:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	if v1173 == int32(0) {
		goto L450
	} else {
		goto L478
	}
L458:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	if v1151 == int32(0) {
		goto L450
	} else {
		goto L470
	}
L459:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	if v1129 == int32(0) {
		goto L450
	} else {
		goto L462
	}
L460:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+72))
	v1128 = F_path_is_reparameterizable_by_child(m, v1127, l1)
	mBase = m.M
	if v1128 != 0 {
		goto L450
	} else {
		goto L461
	}
L461:
	;
	v1242 = v1123
	goto L447
L462:
	;
	v1132 = int32(0)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1133 <= v1132 {
		goto L450
	} else {
		goto L463
	}
L463:
	;
	v1136 = v1132
	goto L464
L464:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+12))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1140+v1136<<(uint(int32(2))%32))))
	v1145 = F_path_is_reparameterizable_by_child(m, v1144, l1)
	mBase = m.M
	if v1145 != 0 {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v1248 = int32(0)
	goto L446
L466:
	;
	v1147 = v1136 + int32(1)
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1129)+4))
	if v1147 < v1148 {
		v1136 = v1147
		goto L464
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	goto L465
L469:
	;
	goto L450
L470:
	;
	v1154 = int32(0)
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if v1155 <= v1154 {
		goto L450
	} else {
		goto L471
	}
L471:
	;
	v1158 = v1154
	goto L472
L472:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+12))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1162+v1158<<(uint(int32(2))%32))))
	v1167 = F_path_is_reparameterizable_by_child(m, v1166, l1)
	mBase = m.M
	if v1167 != 0 {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	v1248 = int32(0)
	goto L446
L474:
	;
	v1169 = v1158 + int32(1)
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1151)+4))
	if v1169 < v1170 {
		v1158 = v1169
		goto L472
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	goto L473
L477:
	;
	goto L450
L478:
	;
	v1176 = F_path_is_reparameterizable_by_child(m, v1173, l1)
	mBase = m.M
	if v1176 != 0 {
		goto L450
	} else {
		goto L479
	}
L479:
	;
	v1242 = v1123
	goto L447
L480:
	;
	v1180 = int32(0)
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1177)+4))
	if v1181 <= v1180 {
		goto L450
	} else {
		goto L481
	}
L481:
	;
	v1184 = v1180
	goto L482
L482:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1177)+12))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1188+v1184<<(uint(int32(2))%32))))
	v1193 = F_path_is_reparameterizable_by_child(m, v1192, l1)
	mBase = m.M
	if v1193 != 0 {
		goto L484
	} else {
		goto L485
	}
L483:
	;
	v1248 = int32(0)
	goto L446
L484:
	;
	v1195 = v1184 + int32(1)
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1177)+4))
	if v1195 < v1196 {
		v1184 = v1195
		goto L482
	} else {
		goto L487
	}
L485:
	;
	goto L486
L486:
	;
	goto L483
L487:
	;
	goto L450
L488:
	;
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+84))
	v1204 = F_path_is_reparameterizable_by_child(m, v1203, l1)
	mBase = m.M
	if v1204 != 0 {
		goto L450
	} else {
		goto L489
	}
L489:
	;
	v1242 = v1123
	goto L447
L490:
	;
	v1208 = int32(0)
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1209 <= v1208 {
		goto L450
	} else {
		goto L491
	}
L491:
	;
	v1212 = v1208
	goto L492
L492:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1216+v1212<<(uint(int32(2))%32))))
	v1221 = F_path_is_reparameterizable_by_child(m, v1220, l1)
	mBase = m.M
	if v1221 != 0 {
		goto L494
	} else {
		goto L495
	}
L493:
	;
	v1248 = int32(0)
	goto L446
L494:
	;
	v1223 = v1212 + int32(1)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+4))
	if v1223 < v1224 {
		v1212 = v1223
		goto L492
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	goto L493
L497:
	;
	goto L450
L498:
	;
	v1242 = v1123
	goto L447
L499:
	;
	v1242 = v1123
	goto L447
L500:
	;
	goto L450
L501:
	;
	v1250 = v1103 + int32(1)
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(v1096)+4))
	if v1250 < v1251 {
		v1103 = v1250
		goto L444
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	goto L445
L504:
	;
	goto L17
L505:
	;
	if v1392 != 0 {
		goto L17
	} else {
		goto L560
	}
L506:
	;
	v1392 = v1386
	goto L505
L507:
	;
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(v1259)+4))
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1264 = F_bms_overlap(m, v1262, v1263)
	mBase = m.M
	if v1264 == int32(0) {
		v1386 = v1258
		goto L506
	} else {
		goto L508
	}
L508:
	;
	v1267 = int32(0)
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1255)))
	switch v1268 - int32(279) {
	case 0, 1:
		goto L509
	default:
		v1386 = v1267
		goto L506
	case 3:
		goto L519
	case 4:
		goto L518
	case 5:
		goto L517
	case 9:
		goto L516
	case 10:
		goto L515
	case 11:
		goto L513
	case 14:
		goto L512
	case 15:
		goto L511
	case 17:
		goto L510
	case 19, 20, 21:
		goto L514
	}
L509:
	;
	v1386 = int32(1)
	goto L506
L510:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	v1376 = F_path_is_reparameterizable_by_child(m, v1375, l1)
	mBase = m.M
	if v1376 == int32(0) {
		v1386 = v1267
		goto L506
	} else {
		goto L559
	}
L511:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	v1374 = F_path_is_reparameterizable_by_child(m, v1373, l1)
	mBase = m.M
	if v1374 != 0 {
		goto L509
	} else {
		goto L558
	}
L512:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	v1372 = F_path_is_reparameterizable_by_child(m, v1371, l1)
	mBase = m.M
	if v1372 != 0 {
		goto L509
	} else {
		goto L557
	}
L513:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	if v1349 == int32(0) {
		goto L509
	} else {
		goto L549
	}
L514:
	;
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+80))
	v1344 = F_path_is_reparameterizable_by_child(m, v1343, l1)
	mBase = m.M
	if v1344 == int32(0) {
		v1386 = v1267
		goto L506
	} else {
		goto L547
	}
L515:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+76))
	if v1321 == int32(0) {
		goto L509
	} else {
		goto L539
	}
L516:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	if v1317 == int32(0) {
		goto L509
	} else {
		goto L537
	}
L517:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	if v1295 == int32(0) {
		goto L509
	} else {
		goto L529
	}
L518:
	;
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	if v1273 == int32(0) {
		goto L509
	} else {
		goto L521
	}
L519:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+72))
	v1272 = F_path_is_reparameterizable_by_child(m, v1271, l1)
	mBase = m.M
	if v1272 != 0 {
		goto L509
	} else {
		goto L520
	}
L520:
	;
	v1386 = v1267
	goto L506
L521:
	;
	v1276 = int32(0)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+4))
	if v1277 <= v1276 {
		goto L509
	} else {
		goto L522
	}
L522:
	;
	v1280 = v1276
	goto L523
L523:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+12))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1284+v1280<<(uint(int32(2))%32))))
	v1289 = F_path_is_reparameterizable_by_child(m, v1288, l1)
	mBase = m.M
	if v1289 != 0 {
		goto L525
	} else {
		goto L526
	}
L524:
	;
	v1392 = int32(0)
	goto L505
L525:
	;
	v1291 = v1280 + int32(1)
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1273)+4))
	if v1291 < v1292 {
		v1280 = v1291
		goto L523
	} else {
		goto L528
	}
L526:
	;
	goto L527
L527:
	;
	goto L524
L528:
	;
	goto L509
L529:
	;
	v1298 = int32(0)
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	if v1299 <= v1298 {
		goto L509
	} else {
		goto L530
	}
L530:
	;
	v1302 = v1298
	goto L531
L531:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+12))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1306+v1302<<(uint(int32(2))%32))))
	v1311 = F_path_is_reparameterizable_by_child(m, v1310, l1)
	mBase = m.M
	if v1311 != 0 {
		goto L533
	} else {
		goto L534
	}
L532:
	;
	v1392 = int32(0)
	goto L505
L533:
	;
	v1313 = v1302 + int32(1)
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+4))
	if v1313 < v1314 {
		v1302 = v1313
		goto L531
	} else {
		goto L536
	}
L534:
	;
	goto L535
L535:
	;
	goto L532
L536:
	;
	goto L509
L537:
	;
	v1320 = F_path_is_reparameterizable_by_child(m, v1317, l1)
	mBase = m.M
	if v1320 != 0 {
		goto L509
	} else {
		goto L538
	}
L538:
	;
	v1386 = v1267
	goto L506
L539:
	;
	v1324 = int32(0)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+4))
	if v1325 <= v1324 {
		goto L509
	} else {
		goto L540
	}
L540:
	;
	v1328 = v1324
	goto L541
L541:
	;
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+12))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1332+v1328<<(uint(int32(2))%32))))
	v1337 = F_path_is_reparameterizable_by_child(m, v1336, l1)
	mBase = m.M
	if v1337 != 0 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	v1392 = int32(0)
	goto L505
L543:
	;
	v1339 = v1328 + int32(1)
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1321)+4))
	if v1339 < v1340 {
		v1328 = v1339
		goto L541
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	goto L542
L546:
	;
	goto L509
L547:
	;
	v1347 = *(*int32)(unsafe.Add(mBase, uint32(v1255)+84))
	v1348 = F_path_is_reparameterizable_by_child(m, v1347, l1)
	mBase = m.M
	if v1348 != 0 {
		goto L509
	} else {
		goto L548
	}
L548:
	;
	v1386 = v1267
	goto L506
L549:
	;
	v1352 = int32(0)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+4))
	if v1353 <= v1352 {
		goto L509
	} else {
		goto L550
	}
L550:
	;
	v1356 = v1352
	goto L551
L551:
	;
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+12))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v1360+v1356<<(uint(int32(2))%32))))
	v1365 = F_path_is_reparameterizable_by_child(m, v1364, l1)
	mBase = m.M
	if v1365 != 0 {
		goto L553
	} else {
		goto L554
	}
L552:
	;
	v1392 = int32(0)
	goto L505
L553:
	;
	v1367 = v1356 + int32(1)
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+4))
	if v1367 < v1368 {
		v1356 = v1367
		goto L551
	} else {
		goto L556
	}
L554:
	;
	goto L555
L555:
	;
	goto L552
L556:
	;
	goto L509
L557:
	;
	v1386 = v1267
	goto L506
L558:
	;
	v1386 = v1267
	goto L506
L559:
	;
	goto L509
L560:
	;
	v1678 = v59
	goto L1
L561:
	;
	if v1530 != 0 {
		goto L17
	} else {
		goto L616
	}
L562:
	;
	v1530 = v1524
	goto L561
L563:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+4))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1402 = F_bms_overlap(m, v1400, v1401)
	mBase = m.M
	if v1402 == int32(0) {
		v1524 = v1396
		goto L562
	} else {
		goto L564
	}
L564:
	;
	v1405 = int32(0)
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1393)))
	switch v1406 - int32(279) {
	case 0, 1:
		goto L565
	default:
		v1524 = v1405
		goto L562
	case 3:
		goto L575
	case 4:
		goto L574
	case 5:
		goto L573
	case 9:
		goto L572
	case 10:
		goto L571
	case 11:
		goto L569
	case 14:
		goto L568
	case 15:
		goto L567
	case 17:
		goto L566
	case 19, 20, 21:
		goto L570
	}
L565:
	;
	v1524 = int32(1)
	goto L562
L566:
	;
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	v1514 = F_path_is_reparameterizable_by_child(m, v1513, l1)
	mBase = m.M
	if v1514 == int32(0) {
		v1524 = v1405
		goto L562
	} else {
		goto L615
	}
L567:
	;
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	v1512 = F_path_is_reparameterizable_by_child(m, v1511, l1)
	mBase = m.M
	if v1512 != 0 {
		goto L565
	} else {
		goto L614
	}
L568:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	v1510 = F_path_is_reparameterizable_by_child(m, v1509, l1)
	mBase = m.M
	if v1510 != 0 {
		goto L565
	} else {
		goto L613
	}
L569:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	if v1487 == int32(0) {
		goto L565
	} else {
		goto L605
	}
L570:
	;
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+80))
	v1482 = F_path_is_reparameterizable_by_child(m, v1481, l1)
	mBase = m.M
	if v1482 == int32(0) {
		v1524 = v1405
		goto L562
	} else {
		goto L603
	}
L571:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+76))
	if v1459 == int32(0) {
		goto L565
	} else {
		goto L595
	}
L572:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	if v1455 == int32(0) {
		goto L565
	} else {
		goto L593
	}
L573:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	if v1433 == int32(0) {
		goto L565
	} else {
		goto L585
	}
L574:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	if v1411 == int32(0) {
		goto L565
	} else {
		goto L577
	}
L575:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+72))
	v1410 = F_path_is_reparameterizable_by_child(m, v1409, l1)
	mBase = m.M
	if v1410 != 0 {
		goto L565
	} else {
		goto L576
	}
L576:
	;
	v1524 = v1405
	goto L562
L577:
	;
	v1414 = int32(0)
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+4))
	if v1415 <= v1414 {
		goto L565
	} else {
		goto L578
	}
L578:
	;
	v1418 = v1414
	goto L579
L579:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+12))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1422+v1418<<(uint(int32(2))%32))))
	v1427 = F_path_is_reparameterizable_by_child(m, v1426, l1)
	mBase = m.M
	if v1427 != 0 {
		goto L581
	} else {
		goto L582
	}
L580:
	;
	v1530 = int32(0)
	goto L561
L581:
	;
	v1429 = v1418 + int32(1)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1411)+4))
	if v1429 < v1430 {
		v1418 = v1429
		goto L579
	} else {
		goto L584
	}
L582:
	;
	goto L583
L583:
	;
	goto L580
L584:
	;
	goto L565
L585:
	;
	v1436 = int32(0)
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+4))
	if v1437 <= v1436 {
		goto L565
	} else {
		goto L586
	}
L586:
	;
	v1440 = v1436
	goto L587
L587:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+12))
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1444+v1440<<(uint(int32(2))%32))))
	v1449 = F_path_is_reparameterizable_by_child(m, v1448, l1)
	mBase = m.M
	if v1449 != 0 {
		goto L589
	} else {
		goto L590
	}
L588:
	;
	v1530 = int32(0)
	goto L561
L589:
	;
	v1451 = v1440 + int32(1)
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1433)+4))
	if v1451 < v1452 {
		v1440 = v1451
		goto L587
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	goto L588
L592:
	;
	goto L565
L593:
	;
	v1458 = F_path_is_reparameterizable_by_child(m, v1455, l1)
	mBase = m.M
	if v1458 != 0 {
		goto L565
	} else {
		goto L594
	}
L594:
	;
	v1524 = v1405
	goto L562
L595:
	;
	v1462 = int32(0)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+4))
	if v1463 <= v1462 {
		goto L565
	} else {
		goto L596
	}
L596:
	;
	v1466 = v1462
	goto L597
L597:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+12))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1470+v1466<<(uint(int32(2))%32))))
	v1475 = F_path_is_reparameterizable_by_child(m, v1474, l1)
	mBase = m.M
	if v1475 != 0 {
		goto L599
	} else {
		goto L600
	}
L598:
	;
	v1530 = int32(0)
	goto L561
L599:
	;
	v1477 = v1466 + int32(1)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1459)+4))
	if v1477 < v1478 {
		v1466 = v1477
		goto L597
	} else {
		goto L602
	}
L600:
	;
	goto L601
L601:
	;
	goto L598
L602:
	;
	goto L565
L603:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1393)+84))
	v1486 = F_path_is_reparameterizable_by_child(m, v1485, l1)
	mBase = m.M
	if v1486 != 0 {
		goto L565
	} else {
		goto L604
	}
L604:
	;
	v1524 = v1405
	goto L562
L605:
	;
	v1490 = int32(0)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+4))
	if v1491 <= v1490 {
		goto L565
	} else {
		goto L606
	}
L606:
	;
	v1494 = v1490
	goto L607
L607:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+12))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1498+v1494<<(uint(int32(2))%32))))
	v1503 = F_path_is_reparameterizable_by_child(m, v1502, l1)
	mBase = m.M
	if v1503 != 0 {
		goto L609
	} else {
		goto L610
	}
L608:
	;
	v1530 = int32(0)
	goto L561
L609:
	;
	v1505 = v1494 + int32(1)
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v1487)+4))
	if v1505 < v1506 {
		v1494 = v1505
		goto L607
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	goto L608
L612:
	;
	goto L565
L613:
	;
	v1524 = v1405
	goto L562
L614:
	;
	v1524 = v1405
	goto L562
L615:
	;
	goto L565
L616:
	;
	v1678 = v59
	goto L1
L617:
	;
	if v1668 == int32(0) {
		v1678 = v59
		goto L1
	} else {
		goto L672
	}
L618:
	;
	v1668 = v1662
	goto L617
L619:
	;
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(v1535)+4))
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1540 = F_bms_overlap(m, v1538, v1539)
	mBase = m.M
	if v1540 == int32(0) {
		v1662 = v1534
		goto L618
	} else {
		goto L620
	}
L620:
	;
	v1543 = int32(0)
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v1531)))
	switch v1544 - int32(279) {
	case 0, 1:
		goto L621
	default:
		v1662 = v1543
		goto L618
	case 3:
		goto L631
	case 4:
		goto L630
	case 5:
		goto L629
	case 9:
		goto L628
	case 10:
		goto L627
	case 11:
		goto L625
	case 14:
		goto L624
	case 15:
		goto L623
	case 17:
		goto L622
	case 19, 20, 21:
		goto L626
	}
L621:
	;
	v1662 = int32(1)
	goto L618
L622:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	v1652 = F_path_is_reparameterizable_by_child(m, v1651, l1)
	mBase = m.M
	if v1652 == int32(0) {
		v1662 = v1543
		goto L618
	} else {
		goto L671
	}
L623:
	;
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	v1650 = F_path_is_reparameterizable_by_child(m, v1649, l1)
	mBase = m.M
	if v1650 != 0 {
		goto L621
	} else {
		goto L670
	}
L624:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	v1648 = F_path_is_reparameterizable_by_child(m, v1647, l1)
	mBase = m.M
	if v1648 != 0 {
		goto L621
	} else {
		goto L669
	}
L625:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	if v1625 == int32(0) {
		goto L621
	} else {
		goto L661
	}
L626:
	;
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+80))
	v1620 = F_path_is_reparameterizable_by_child(m, v1619, l1)
	mBase = m.M
	if v1620 == int32(0) {
		v1662 = v1543
		goto L618
	} else {
		goto L659
	}
L627:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+76))
	if v1597 == int32(0) {
		goto L621
	} else {
		goto L651
	}
L628:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	if v1593 == int32(0) {
		goto L621
	} else {
		goto L649
	}
L629:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	if v1571 == int32(0) {
		goto L621
	} else {
		goto L641
	}
L630:
	;
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	if v1549 == int32(0) {
		goto L621
	} else {
		goto L633
	}
L631:
	;
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+72))
	v1548 = F_path_is_reparameterizable_by_child(m, v1547, l1)
	mBase = m.M
	if v1548 != 0 {
		goto L621
	} else {
		goto L632
	}
L632:
	;
	v1662 = v1543
	goto L618
L633:
	;
	v1552 = int32(0)
	v1553 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	if v1553 <= v1552 {
		goto L621
	} else {
		goto L634
	}
L634:
	;
	v1556 = v1552
	goto L635
L635:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+12))
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1560+v1556<<(uint(int32(2))%32))))
	v1565 = F_path_is_reparameterizable_by_child(m, v1564, l1)
	mBase = m.M
	if v1565 != 0 {
		goto L637
	} else {
		goto L638
	}
L636:
	;
	v1668 = int32(0)
	goto L617
L637:
	;
	v1567 = v1556 + int32(1)
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1549)+4))
	if v1567 < v1568 {
		v1556 = v1567
		goto L635
	} else {
		goto L640
	}
L638:
	;
	goto L639
L639:
	;
	goto L636
L640:
	;
	goto L621
L641:
	;
	v1574 = int32(0)
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+4))
	if v1575 <= v1574 {
		goto L621
	} else {
		goto L642
	}
L642:
	;
	v1578 = v1574
	goto L643
L643:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+12))
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1582+v1578<<(uint(int32(2))%32))))
	v1587 = F_path_is_reparameterizable_by_child(m, v1586, l1)
	mBase = m.M
	if v1587 != 0 {
		goto L645
	} else {
		goto L646
	}
L644:
	;
	v1668 = int32(0)
	goto L617
L645:
	;
	v1589 = v1578 + int32(1)
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1571)+4))
	if v1589 < v1590 {
		v1578 = v1589
		goto L643
	} else {
		goto L648
	}
L646:
	;
	goto L647
L647:
	;
	goto L644
L648:
	;
	goto L621
L649:
	;
	v1596 = F_path_is_reparameterizable_by_child(m, v1593, l1)
	mBase = m.M
	if v1596 != 0 {
		goto L621
	} else {
		goto L650
	}
L650:
	;
	v1662 = v1543
	goto L618
L651:
	;
	v1600 = int32(0)
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(v1597)+4))
	if v1601 <= v1600 {
		goto L621
	} else {
		goto L652
	}
L652:
	;
	v1604 = v1600
	goto L653
L653:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v1597)+12))
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1608+v1604<<(uint(int32(2))%32))))
	v1613 = F_path_is_reparameterizable_by_child(m, v1612, l1)
	mBase = m.M
	if v1613 != 0 {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	v1668 = int32(0)
	goto L617
L655:
	;
	v1615 = v1604 + int32(1)
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1597)+4))
	if v1615 < v1616 {
		v1604 = v1615
		goto L653
	} else {
		goto L658
	}
L656:
	;
	goto L657
L657:
	;
	goto L654
L658:
	;
	goto L621
L659:
	;
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1531)+84))
	v1624 = F_path_is_reparameterizable_by_child(m, v1623, l1)
	mBase = m.M
	if v1624 != 0 {
		goto L621
	} else {
		goto L660
	}
L660:
	;
	v1662 = v1543
	goto L618
L661:
	;
	v1628 = int32(0)
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+4))
	if v1629 <= v1628 {
		goto L621
	} else {
		goto L662
	}
L662:
	;
	v1632 = v1628
	goto L663
L663:
	;
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+12))
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1636+v1632<<(uint(int32(2))%32))))
	v1641 = F_path_is_reparameterizable_by_child(m, v1640, l1)
	mBase = m.M
	if v1641 != 0 {
		goto L665
	} else {
		goto L666
	}
L664:
	;
	v1668 = int32(0)
	goto L617
L665:
	;
	v1643 = v1632 + int32(1)
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1625)+4))
	if v1643 < v1644 {
		v1632 = v1643
		goto L663
	} else {
		goto L668
	}
L666:
	;
	goto L667
L667:
	;
	goto L664
L668:
	;
	goto L621
L669:
	;
	v1662 = v1543
	goto L618
L670:
	;
	v1662 = v1543
	goto L618
L671:
	;
	goto L621
L672:
	;
	goto L17
}
func F_path_n_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v12 <= v11)
		}
	}
}
func F_path_n_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			return base.B2i32(v11 < v12)
		}
	}
}
func F_path_npoints(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		return v7
	}
}
