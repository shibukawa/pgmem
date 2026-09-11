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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
		goto L8
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
	if l0 != 0 {
		v33 = v18
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return v33
L10:
	;
	if v18 == int32(0) {
		v33 = v18
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v23 = v18
	goto L12
L12:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_search_path[2]))
	if v25 == v27 {
		v33 = v23
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v33 = int32(0)
	goto L9
L14:
	;
	v29 = F_list_delete_first(m, v23)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v29 != 0 {
		v23 = v29
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	if l4 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v264 = int32(0)
	if l1 == l2 {
		v294 = v264
		goto L73
	} else {
		goto L74
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
	v88 = v76
	goto L17
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v92 = int32(0)
	if v88 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L1
L19:
	;
	v103 = v92
	v104 = v92
	v107 = int32(0)
	goto L22
L20:
	;
	v135 = v92
	v136 = v92
	goto L21
L21:
	;
	if v88&int32(1) == int32(0) {
		v155 = v135
		goto L37
	} else {
		goto L38
	}
L22:
	;
	v113 = v91 + v104<<(uint(int32(2))%32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+9)))
	if v115 != 0 {
		v119 = v103
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v135 = v125
	v136 = v127
	goto L21
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+9)))
	if v121 != 0 {
		v125 = v119
		goto L30
	} else {
		goto L31
	}
L25:
	;
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	if v116 <= v117 {
		v119 = v103
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v119 = v114
	goto L24
L29:
	;
	goto L28
L30:
	;
	v126 = int32(2)
	v127 = v104 + v126
	v129 = v107 + v126
	if v129 != v88&int32(2147483646) {
		v103 = v125
		v104 = v127
		v107 = v129
		goto L22
	} else {
		goto L36
	}
L31:
	;
	if v119 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	if v122 <= v123 {
		v125 = v119
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v125 = v120
	goto L30
L35:
	;
	goto L34
L36:
	;
	goto L23
L37:
	;
	if v155 == int32(0) {
		goto L1
	} else {
		goto L44
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v91+v136<<(uint(int32(2))%32))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+9)))
	if v151 != 0 {
		v155 = v135
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if v135 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	if v152 <= v153 {
		v155 = v135
		goto L37
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v155 = v150
	goto L37
L43:
	;
	goto L42
L44:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	if v159 == int32(2147483647) {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v155)+9)) = uint8(v162)
	if l2 == v155 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v165 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v249 {
		v88 = v249
		goto L17
	} else {
		goto L72
	}
L48:
	;
	v168 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v169 <= v168 {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v180 = v168
	goto L50
L50:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v180<<(uint(int32(2))%32))))
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
	v234 = v180 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v234 < v235 {
		v180 = v234
		goto L50
	} else {
		goto L71
	}
L53:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+8)))
	if v189 != 0 {
		goto L52
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v192 = v190 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if v192 < v193 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v188)+12)) = v192
	goto L52
L58:
	;
	goto L59
L59:
	;
	if v192 != v193 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	if v198 == int32(0) {
		goto L52
	} else {
		goto L61
	}
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201))))
	if v206 == int32(0) {
		v225 = v205
		v226 = v206
		goto L63
	} else {
		goto L64
	}
L62:
	;
	if int32(0) <= v226-v225 {
		goto L52
	} else {
		goto L70
	}
L63:
	;
	goto L62
L64:
	;
	if v205 != v206 {
		v225 = v205
		v226 = v206
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v210 = v201
	v211 = v202
	goto L66
L66:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+1)))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)))
	if v215 == int32(0) {
		v225 = v214
		v226 = v215
		goto L63
	} else {
		goto L68
	}
L67:
	;
	v225 = v214
	v226 = v215
	goto L63
L68:
	;
	v218 = int32(1)
	if v214 == v215 {
		v210 = v210 + v218
		v211 = v211 + v218
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+16)) = v155
	goto L52
L71:
	;
	goto L51
L72:
	;
	goto L18
L73:
	;
	return v294
L74:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v266&int32(1) == int32(0) {
		v294 = v264
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v273 = l2
	v275 = v264
	goto L76
L76:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v284 = F_lcons(m, v283, v275)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v294 = v284
	goto L73
L78:
	;
	return int32(0)
L79:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	if v288 != l1 {
		v273 = v288
		v275 = v284
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
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
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
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
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
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
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
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
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
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
	var v804 int32
	_ = v804
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
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
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
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
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
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
	var v1241 int32
	_ = v1241
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1294 int32
	_ = v1294
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1355 int32
	_ = v1355
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
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
	var v1385 int32
	_ = v1385
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1443 int32
	_ = v1443
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1508 int32
	_ = v1508
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
	var v1523 int32
	_ = v1523
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
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
	var v1661 int32
	_ = v1661
	var v1667 int32
	_ = v1667
	var v1677 int32
	_ = v1677
	v5 = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == int32(0) {
		v1677 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1677
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v11 = int32(0)
	if v9 == v11 {
		v52 = v11
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v52 == int32(0) {
		v1677 = v5
		goto L1
	} else {
		goto L17
	}
L4:
	;
	goto L3
L5:
	;
	if v10 == int32(0) {
		v52 = v11
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v20 < v21 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v20
	goto L9
L8:
	;
	v23 = v21
	goto L9
L9:
	;
	if v23 <= int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = int32(1)
	goto L12
L11:
	;
	v26 = v23
	goto L12
L12:
	;
	v27 = int32(8)
	v32 = int32(0)
	goto L13
L13:
	;
	v39 = v32 << (uint(int32(2)) % 32)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10+v27+v39)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v9+v27))))
	v44 = v41 & v43
	v46 = base.B2i32(v44 != int32(0))
	if v44 != 0 {
		v52 = v46
		goto L4
	} else {
		goto L15
	}
L14:
	;
	v52 = v46
	goto L4
L15:
	;
	v48 = v32 + int32(1)
	if v48 != v26 {
		v32 = v48
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v59 - int32(279) {
	case 0, 1:
		goto L18
	default:
		v1677 = v58
		goto L1
	case 3:
		goto L28
	case 4:
		goto L27
	case 5:
		goto L26
	case 9:
		goto L25
	case 10:
		goto L24
	case 11:
		goto L22
	case 14:
		goto L21
	case 15:
		goto L20
	case 17:
		goto L19
	case 19, 20, 21:
		goto L23
	}
L18:
	;
	v1677 = int32(1)
	goto L1
L19:
	;
	v1530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1533 = int32(1)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+16))
	if v1534 == int32(0) {
		v1661 = v1533
		goto L619
	} else {
		goto L620
	}
L20:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1395 = int32(1)
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+16))
	if v1396 == int32(0) {
		v1523 = v1395
		goto L563
	} else {
		goto L564
	}
L21:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v1257 = int32(1)
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+16))
	if v1258 == int32(0) {
		v1385 = v1257
		goto L507
	} else {
		goto L508
	}
L22:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v1095 == int32(0) {
		goto L18
	} else {
		goto L443
	}
L23:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v820 = int32(1)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v817)+16))
	if v821 == int32(0) {
		v948 = v820
		goto L332
	} else {
		goto L333
	}
L24:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v658 == int32(0) {
		goto L18
	} else {
		goto L268
	}
L25:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v518 == int32(0) {
		goto L18
	} else {
		goto L211
	}
L26:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v359 == int32(0) {
		goto L18
	} else {
		goto L148
	}
L27:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v200 == int32(0) {
		goto L18
	} else {
		goto L85
	}
L28:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v65 = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v66 == int32(0) {
		v193 = v65
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v199 != 0 {
		goto L18
	} else {
		goto L84
	}
L30:
	;
	v199 = v193
	goto L29
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v71 = F_bms_overlap(m, v69, v70)
	mBase = m.M
	if v71 == int32(0) {
		v193 = v65
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	switch v75 - int32(279) {
	case 0, 1:
		goto L33
	default:
		v193 = v74
		goto L30
	case 3:
		goto L43
	case 4:
		goto L42
	case 5:
		goto L41
	case 9:
		goto L40
	case 10:
		goto L39
	case 11:
		goto L37
	case 14:
		goto L36
	case 15:
		goto L35
	case 17:
		goto L34
	case 19, 20, 21:
		goto L38
	}
L33:
	;
	v193 = int32(1)
	goto L30
L34:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v183 = F_path_is_reparameterizable_by_child(m, v182, l1)
	mBase = m.M
	if v183 == int32(0) {
		v193 = v74
		goto L30
	} else {
		goto L83
	}
L35:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v181 = F_path_is_reparameterizable_by_child(m, v180, l1)
	mBase = m.M
	if v181 != 0 {
		goto L33
	} else {
		goto L82
	}
L36:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v179 = F_path_is_reparameterizable_by_child(m, v178, l1)
	mBase = m.M
	if v179 != 0 {
		goto L33
	} else {
		goto L81
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	if v156 == int32(0) {
		goto L33
	} else {
		goto L73
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v62)+80))
	v151 = F_path_is_reparameterizable_by_child(m, v150, l1)
	mBase = m.M
	if v151 == int32(0) {
		v193 = v74
		goto L30
	} else {
		goto L71
	}
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v62)+76))
	if v128 == int32(0) {
		goto L33
	} else {
		goto L63
	}
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	if v124 == int32(0) {
		goto L33
	} else {
		goto L61
	}
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	if v102 == int32(0) {
		goto L33
	} else {
		goto L53
	}
L42:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	if v80 == int32(0) {
		goto L33
	} else {
		goto L45
	}
L43:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)+72))
	v79 = F_path_is_reparameterizable_by_child(m, v78, l1)
	mBase = m.M
	if v79 != 0 {
		goto L33
	} else {
		goto L44
	}
L44:
	;
	v193 = v74
	goto L30
L45:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v84 <= v83 {
		goto L33
	} else {
		goto L46
	}
L46:
	;
	v87 = v83
	goto L47
L47:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91+v87<<(uint(int32(2))%32))))
	v96 = F_path_is_reparameterizable_by_child(m, v95, l1)
	mBase = m.M
	if v96 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v199 = int32(0)
	goto L29
L49:
	;
	v98 = v87 + int32(1)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v98 < v99 {
		v87 = v98
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L33
L53:
	;
	v105 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v106 <= v105 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	v109 = v105
	goto L55
L55:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v109<<(uint(int32(2))%32))))
	v118 = F_path_is_reparameterizable_by_child(m, v117, l1)
	mBase = m.M
	if v118 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v199 = int32(0)
	goto L29
L57:
	;
	v120 = v109 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v120 < v121 {
		v109 = v120
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L33
L61:
	;
	v127 = F_path_is_reparameterizable_by_child(m, v124, l1)
	mBase = m.M
	if v127 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	v193 = v74
	goto L30
L63:
	;
	v131 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v132 <= v131 {
		goto L33
	} else {
		goto L64
	}
L64:
	;
	v135 = v131
	goto L65
L65:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v135<<(uint(int32(2))%32))))
	v144 = F_path_is_reparameterizable_by_child(m, v143, l1)
	mBase = m.M
	if v144 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v199 = int32(0)
	goto L29
L67:
	;
	v146 = v135 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	if v146 < v147 {
		v135 = v146
		goto L65
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L33
L71:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v62)+84))
	v155 = F_path_is_reparameterizable_by_child(m, v154, l1)
	mBase = m.M
	if v155 != 0 {
		goto L33
	} else {
		goto L72
	}
L72:
	;
	v193 = v74
	goto L30
L73:
	;
	v159 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v160 <= v159 {
		goto L33
	} else {
		goto L74
	}
L74:
	;
	v163 = v159
	goto L75
L75:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v163<<(uint(int32(2))%32))))
	v172 = F_path_is_reparameterizable_by_child(m, v171, l1)
	mBase = m.M
	if v172 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v199 = int32(0)
	goto L29
L77:
	;
	v174 = v163 + int32(1)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	if v174 < v175 {
		v163 = v174
		goto L75
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L76
L80:
	;
	goto L33
L81:
	;
	v193 = v74
	goto L30
L82:
	;
	v193 = v74
	goto L30
L83:
	;
	goto L33
L84:
	;
	v1677 = v58
	goto L1
L85:
	;
	v203 = int32(0)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v204 <= v203 {
		goto L18
	} else {
		goto L86
	}
L86:
	;
	v207 = v203
	goto L87
L87:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+v207<<(uint(int32(2))%32))))
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v219 == int32(0) {
		v346 = v218
		goto L90
	} else {
		goto L91
	}
L88:
	;
	return int32(0)
L89:
	;
	if v352 != 0 {
		goto L144
	} else {
		goto L145
	}
L90:
	;
	v352 = v346
	goto L89
L91:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v224 = F_bms_overlap(m, v222, v223)
	mBase = m.M
	if v224 == int32(0) {
		v346 = v218
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v227 = int32(0)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	switch v228 - int32(279) {
	case 0, 1:
		goto L93
	default:
		v346 = v227
		goto L90
	case 3:
		goto L103
	case 4:
		goto L102
	case 5:
		goto L101
	case 9:
		goto L100
	case 10:
		goto L99
	case 11:
		goto L97
	case 14:
		goto L96
	case 15:
		goto L95
	case 17:
		goto L94
	case 19, 20, 21:
		goto L98
	}
L93:
	;
	v346 = int32(1)
	goto L90
L94:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	v336 = F_path_is_reparameterizable_by_child(m, v335, l1)
	mBase = m.M
	if v336 == int32(0) {
		v346 = v227
		goto L90
	} else {
		goto L143
	}
L95:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	v334 = F_path_is_reparameterizable_by_child(m, v333, l1)
	mBase = m.M
	if v334 != 0 {
		goto L93
	} else {
		goto L142
	}
L96:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	v332 = F_path_is_reparameterizable_by_child(m, v331, l1)
	mBase = m.M
	if v332 != 0 {
		goto L93
	} else {
		goto L141
	}
L97:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	if v309 == int32(0) {
		goto L93
	} else {
		goto L133
	}
L98:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v215)+80))
	v304 = F_path_is_reparameterizable_by_child(m, v303, l1)
	mBase = m.M
	if v304 == int32(0) {
		v346 = v227
		goto L90
	} else {
		goto L131
	}
L99:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v215)+76))
	if v281 == int32(0) {
		goto L93
	} else {
		goto L123
	}
L100:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	if v277 == int32(0) {
		goto L93
	} else {
		goto L121
	}
L101:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	if v255 == int32(0) {
		goto L93
	} else {
		goto L113
	}
L102:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	if v233 == int32(0) {
		goto L93
	} else {
		goto L105
	}
L103:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v215)+72))
	v232 = F_path_is_reparameterizable_by_child(m, v231, l1)
	mBase = m.M
	if v232 != 0 {
		goto L93
	} else {
		goto L104
	}
L104:
	;
	v346 = v227
	goto L90
L105:
	;
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v237 <= v236 {
		goto L93
	} else {
		goto L106
	}
L106:
	;
	v240 = v236
	goto L107
L107:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244+v240<<(uint(int32(2))%32))))
	v249 = F_path_is_reparameterizable_by_child(m, v248, l1)
	mBase = m.M
	if v249 != 0 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v352 = int32(0)
	goto L89
L109:
	;
	v251 = v240 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v251 < v252 {
		v240 = v251
		goto L107
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	goto L108
L112:
	;
	goto L93
L113:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v259 <= v258 {
		goto L93
	} else {
		goto L114
	}
L114:
	;
	v262 = v258
	goto L115
L115:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v266+v262<<(uint(int32(2))%32))))
	v271 = F_path_is_reparameterizable_by_child(m, v270, l1)
	mBase = m.M
	if v271 != 0 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v352 = int32(0)
	goto L89
L117:
	;
	v273 = v262 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	if v273 < v274 {
		v262 = v273
		goto L115
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	goto L116
L120:
	;
	goto L93
L121:
	;
	v280 = F_path_is_reparameterizable_by_child(m, v277, l1)
	mBase = m.M
	if v280 != 0 {
		goto L93
	} else {
		goto L122
	}
L122:
	;
	v346 = v227
	goto L90
L123:
	;
	v284 = int32(0)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v285 <= v284 {
		goto L93
	} else {
		goto L124
	}
L124:
	;
	v288 = v284
	goto L125
L125:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v281)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v288<<(uint(int32(2))%32))))
	v297 = F_path_is_reparameterizable_by_child(m, v296, l1)
	mBase = m.M
	if v297 != 0 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v352 = int32(0)
	goto L89
L127:
	;
	v299 = v288 + int32(1)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	if v299 < v300 {
		v288 = v299
		goto L125
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	goto L126
L130:
	;
	goto L93
L131:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v215)+84))
	v308 = F_path_is_reparameterizable_by_child(m, v307, l1)
	mBase = m.M
	if v308 != 0 {
		goto L93
	} else {
		goto L132
	}
L132:
	;
	v346 = v227
	goto L90
L133:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v313 <= v312 {
		goto L93
	} else {
		goto L134
	}
L134:
	;
	v316 = v312
	goto L135
L135:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320+v316<<(uint(int32(2))%32))))
	v325 = F_path_is_reparameterizable_by_child(m, v324, l1)
	mBase = m.M
	if v325 != 0 {
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v352 = int32(0)
	goto L89
L137:
	;
	v327 = v316 + int32(1)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	if v327 < v328 {
		v316 = v327
		goto L135
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	goto L136
L140:
	;
	goto L93
L141:
	;
	v346 = v227
	goto L90
L142:
	;
	v346 = v227
	goto L90
L143:
	;
	goto L93
L144:
	;
	v354 = v207 + int32(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v354 < v355 {
		v207 = v354
		goto L87
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	goto L88
L147:
	;
	goto L18
L148:
	;
	v362 = int32(0)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v363 <= v362 {
		goto L18
	} else {
		goto L149
	}
L149:
	;
	v366 = v362
	goto L150
L150:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v359)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v370+v366<<(uint(int32(2))%32))))
	v377 = int32(1)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v374)+16))
	if v378 == int32(0) {
		v505 = v377
		goto L153
	} else {
		goto L154
	}
L151:
	;
	return int32(0)
L152:
	;
	if v511 != 0 {
		goto L207
	} else {
		goto L208
	}
L153:
	;
	v511 = v505
	goto L152
L154:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v383 = F_bms_overlap(m, v381, v382)
	mBase = m.M
	if v383 == int32(0) {
		v505 = v377
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v386 = int32(0)
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v374)))
	switch v387 - int32(279) {
	case 0, 1:
		goto L156
	default:
		v505 = v386
		goto L153
	case 3:
		goto L166
	case 4:
		goto L165
	case 5:
		goto L164
	case 9:
		goto L163
	case 10:
		goto L162
	case 11:
		goto L160
	case 14:
		goto L159
	case 15:
		goto L158
	case 17:
		goto L157
	case 19, 20, 21:
		goto L161
	}
L156:
	;
	v505 = int32(1)
	goto L153
L157:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	v495 = F_path_is_reparameterizable_by_child(m, v494, l1)
	mBase = m.M
	if v495 == int32(0) {
		v505 = v386
		goto L153
	} else {
		goto L206
	}
L158:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	v493 = F_path_is_reparameterizable_by_child(m, v492, l1)
	mBase = m.M
	if v493 != 0 {
		goto L156
	} else {
		goto L205
	}
L159:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	v491 = F_path_is_reparameterizable_by_child(m, v490, l1)
	mBase = m.M
	if v491 != 0 {
		goto L156
	} else {
		goto L204
	}
L160:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	if v468 == int32(0) {
		goto L156
	} else {
		goto L196
	}
L161:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v374)+80))
	v463 = F_path_is_reparameterizable_by_child(m, v462, l1)
	mBase = m.M
	if v463 == int32(0) {
		v505 = v386
		goto L153
	} else {
		goto L194
	}
L162:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v374)+76))
	if v440 == int32(0) {
		goto L156
	} else {
		goto L186
	}
L163:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	if v436 == int32(0) {
		goto L156
	} else {
		goto L184
	}
L164:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	if v414 == int32(0) {
		goto L156
	} else {
		goto L176
	}
L165:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	if v392 == int32(0) {
		goto L156
	} else {
		goto L168
	}
L166:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v374)+72))
	v391 = F_path_is_reparameterizable_by_child(m, v390, l1)
	mBase = m.M
	if v391 != 0 {
		goto L156
	} else {
		goto L167
	}
L167:
	;
	v505 = v386
	goto L153
L168:
	;
	v395 = int32(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v396 <= v395 {
		goto L156
	} else {
		goto L169
	}
L169:
	;
	v399 = v395
	goto L170
L170:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v392)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403+v399<<(uint(int32(2))%32))))
	v408 = F_path_is_reparameterizable_by_child(m, v407, l1)
	mBase = m.M
	if v408 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	v511 = int32(0)
	goto L152
L172:
	;
	v410 = v399 + int32(1)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v410 < v411 {
		v399 = v410
		goto L170
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	goto L171
L175:
	;
	goto L156
L176:
	;
	v417 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v418 <= v417 {
		goto L156
	} else {
		goto L177
	}
L177:
	;
	v421 = v417
	goto L178
L178:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v414)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v425+v421<<(uint(int32(2))%32))))
	v430 = F_path_is_reparameterizable_by_child(m, v429, l1)
	mBase = m.M
	if v430 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v511 = int32(0)
	goto L152
L180:
	;
	v432 = v421 + int32(1)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v414)+4))
	if v432 < v433 {
		v421 = v432
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L179
L183:
	;
	goto L156
L184:
	;
	v439 = F_path_is_reparameterizable_by_child(m, v436, l1)
	mBase = m.M
	if v439 != 0 {
		goto L156
	} else {
		goto L185
	}
L185:
	;
	v505 = v386
	goto L153
L186:
	;
	v443 = int32(0)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v444 <= v443 {
		goto L156
	} else {
		goto L187
	}
L187:
	;
	v447 = v443
	goto L188
L188:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451+v447<<(uint(int32(2))%32))))
	v456 = F_path_is_reparameterizable_by_child(m, v455, l1)
	mBase = m.M
	if v456 != 0 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v511 = int32(0)
	goto L152
L190:
	;
	v458 = v447 + int32(1)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v458 < v459 {
		v447 = v458
		goto L188
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L189
L193:
	;
	goto L156
L194:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v374)+84))
	v467 = F_path_is_reparameterizable_by_child(m, v466, l1)
	mBase = m.M
	if v467 != 0 {
		goto L156
	} else {
		goto L195
	}
L195:
	;
	v505 = v386
	goto L153
L196:
	;
	v471 = int32(0)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v472 <= v471 {
		goto L156
	} else {
		goto L197
	}
L197:
	;
	v475 = v471
	goto L198
L198:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v468)+12))
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v479+v475<<(uint(int32(2))%32))))
	v484 = F_path_is_reparameterizable_by_child(m, v483, l1)
	mBase = m.M
	if v484 != 0 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v511 = int32(0)
	goto L152
L200:
	;
	v486 = v475 + int32(1)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v468)+4))
	if v486 < v487 {
		v475 = v486
		goto L198
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	goto L199
L203:
	;
	goto L156
L204:
	;
	v505 = v386
	goto L153
L205:
	;
	v505 = v386
	goto L153
L206:
	;
	goto L156
L207:
	;
	v513 = v366 + int32(1)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v513 < v514 {
		v366 = v513
		goto L150
	} else {
		goto L210
	}
L208:
	;
	goto L209
L209:
	;
	goto L151
L210:
	;
	goto L18
L211:
	;
	v523 = int32(1)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v518)+16))
	if v524 == int32(0) {
		v651 = v523
		goto L213
	} else {
		goto L214
	}
L212:
	;
	if v657 != 0 {
		goto L18
	} else {
		goto L267
	}
L213:
	;
	v657 = v651
	goto L212
L214:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v524)+4))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v529 = F_bms_overlap(m, v527, v528)
	mBase = m.M
	if v529 == int32(0) {
		v651 = v523
		goto L213
	} else {
		goto L215
	}
L215:
	;
	v532 = int32(0)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	switch v533 - int32(279) {
	case 0, 1:
		goto L216
	default:
		v651 = v532
		goto L213
	case 3:
		goto L226
	case 4:
		goto L225
	case 5:
		goto L224
	case 9:
		goto L223
	case 10:
		goto L222
	case 11:
		goto L220
	case 14:
		goto L219
	case 15:
		goto L218
	case 17:
		goto L217
	case 19, 20, 21:
		goto L221
	}
L216:
	;
	v651 = int32(1)
	goto L213
L217:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	v641 = F_path_is_reparameterizable_by_child(m, v640, l1)
	mBase = m.M
	if v641 == int32(0) {
		v651 = v532
		goto L213
	} else {
		goto L266
	}
L218:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	v639 = F_path_is_reparameterizable_by_child(m, v638, l1)
	mBase = m.M
	if v639 != 0 {
		goto L216
	} else {
		goto L265
	}
L219:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	v637 = F_path_is_reparameterizable_by_child(m, v636, l1)
	mBase = m.M
	if v637 != 0 {
		goto L216
	} else {
		goto L264
	}
L220:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	if v614 == int32(0) {
		goto L216
	} else {
		goto L256
	}
L221:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v518)+80))
	v609 = F_path_is_reparameterizable_by_child(m, v608, l1)
	mBase = m.M
	if v609 == int32(0) {
		v651 = v532
		goto L213
	} else {
		goto L254
	}
L222:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v518)+76))
	if v586 == int32(0) {
		goto L216
	} else {
		goto L246
	}
L223:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	if v582 == int32(0) {
		goto L216
	} else {
		goto L244
	}
L224:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	if v560 == int32(0) {
		goto L216
	} else {
		goto L236
	}
L225:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	if v538 == int32(0) {
		goto L216
	} else {
		goto L228
	}
L226:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v518)+72))
	v537 = F_path_is_reparameterizable_by_child(m, v536, l1)
	mBase = m.M
	if v537 != 0 {
		goto L216
	} else {
		goto L227
	}
L227:
	;
	v651 = v532
	goto L213
L228:
	;
	v541 = int32(0)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v542 <= v541 {
		goto L216
	} else {
		goto L229
	}
L229:
	;
	v545 = v541
	goto L230
L230:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v538)+12))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549+v545<<(uint(int32(2))%32))))
	v554 = F_path_is_reparameterizable_by_child(m, v553, l1)
	mBase = m.M
	if v554 != 0 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v657 = int32(0)
	goto L212
L232:
	;
	v556 = v545 + int32(1)
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v538)+4))
	if v556 < v557 {
		v545 = v556
		goto L230
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	goto L231
L235:
	;
	goto L216
L236:
	;
	v563 = int32(0)
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	if v564 <= v563 {
		goto L216
	} else {
		goto L237
	}
L237:
	;
	v567 = v563
	goto L238
L238:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v560)+12))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v571+v567<<(uint(int32(2))%32))))
	v576 = F_path_is_reparameterizable_by_child(m, v575, l1)
	mBase = m.M
	if v576 != 0 {
		goto L240
	} else {
		goto L241
	}
L239:
	;
	v657 = int32(0)
	goto L212
L240:
	;
	v578 = v567 + int32(1)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v560)+4))
	if v578 < v579 {
		v567 = v578
		goto L238
	} else {
		goto L243
	}
L241:
	;
	goto L242
L242:
	;
	goto L239
L243:
	;
	goto L216
L244:
	;
	v585 = F_path_is_reparameterizable_by_child(m, v582, l1)
	mBase = m.M
	if v585 != 0 {
		goto L216
	} else {
		goto L245
	}
L245:
	;
	v651 = v532
	goto L213
L246:
	;
	v589 = int32(0)
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	if v590 <= v589 {
		goto L216
	} else {
		goto L247
	}
L247:
	;
	v593 = v589
	goto L248
L248:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v586)+12))
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v597+v593<<(uint(int32(2))%32))))
	v602 = F_path_is_reparameterizable_by_child(m, v601, l1)
	mBase = m.M
	if v602 != 0 {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	v657 = int32(0)
	goto L212
L250:
	;
	v604 = v593 + int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	if v604 < v605 {
		v593 = v604
		goto L248
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	goto L249
L253:
	;
	goto L216
L254:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v518)+84))
	v613 = F_path_is_reparameterizable_by_child(m, v612, l1)
	mBase = m.M
	if v613 != 0 {
		goto L216
	} else {
		goto L255
	}
L255:
	;
	v651 = v532
	goto L213
L256:
	;
	v617 = int32(0)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	if v618 <= v617 {
		goto L216
	} else {
		goto L257
	}
L257:
	;
	v621 = v617
	goto L258
L258:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v614)+12))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v625+v621<<(uint(int32(2))%32))))
	v630 = F_path_is_reparameterizable_by_child(m, v629, l1)
	mBase = m.M
	if v630 != 0 {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	v657 = int32(0)
	goto L212
L260:
	;
	v632 = v621 + int32(1)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	if v632 < v633 {
		v621 = v632
		goto L258
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	goto L259
L263:
	;
	goto L216
L264:
	;
	v651 = v532
	goto L213
L265:
	;
	v651 = v532
	goto L213
L266:
	;
	goto L216
L267:
	;
	v1677 = v58
	goto L1
L268:
	;
	v661 = int32(0)
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	if v662 <= v661 {
		goto L18
	} else {
		goto L269
	}
L269:
	;
	v665 = v661
	goto L270
L270:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v658)+12))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v669+v665<<(uint(int32(2))%32))))
	v676 = int32(1)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v673)+16))
	if v677 == int32(0) {
		v804 = v676
		goto L273
	} else {
		goto L274
	}
L271:
	;
	return int32(0)
L272:
	;
	if v810 != 0 {
		goto L327
	} else {
		goto L328
	}
L273:
	;
	v810 = v804
	goto L272
L274:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v677)+4))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v682 = F_bms_overlap(m, v680, v681)
	mBase = m.M
	if v682 == int32(0) {
		v804 = v676
		goto L273
	} else {
		goto L275
	}
L275:
	;
	v685 = int32(0)
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v673)))
	switch v686 - int32(279) {
	case 0, 1:
		goto L276
	default:
		v804 = v685
		goto L273
	case 3:
		goto L286
	case 4:
		goto L285
	case 5:
		goto L284
	case 9:
		goto L283
	case 10:
		goto L282
	case 11:
		goto L280
	case 14:
		goto L279
	case 15:
		goto L278
	case 17:
		goto L277
	case 19, 20, 21:
		goto L281
	}
L276:
	;
	v804 = int32(1)
	goto L273
L277:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	v794 = F_path_is_reparameterizable_by_child(m, v793, l1)
	mBase = m.M
	if v794 == int32(0) {
		v804 = v685
		goto L273
	} else {
		goto L326
	}
L278:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	v792 = F_path_is_reparameterizable_by_child(m, v791, l1)
	mBase = m.M
	if v792 != 0 {
		goto L276
	} else {
		goto L325
	}
L279:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	v790 = F_path_is_reparameterizable_by_child(m, v789, l1)
	mBase = m.M
	if v790 != 0 {
		goto L276
	} else {
		goto L324
	}
L280:
	;
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	if v767 == int32(0) {
		goto L276
	} else {
		goto L316
	}
L281:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v673)+80))
	v762 = F_path_is_reparameterizable_by_child(m, v761, l1)
	mBase = m.M
	if v762 == int32(0) {
		v804 = v685
		goto L273
	} else {
		goto L314
	}
L282:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v673)+76))
	if v739 == int32(0) {
		goto L276
	} else {
		goto L306
	}
L283:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	if v735 == int32(0) {
		goto L276
	} else {
		goto L304
	}
L284:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	if v713 == int32(0) {
		goto L276
	} else {
		goto L296
	}
L285:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	if v691 == int32(0) {
		goto L276
	} else {
		goto L288
	}
L286:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v673)+72))
	v690 = F_path_is_reparameterizable_by_child(m, v689, l1)
	mBase = m.M
	if v690 != 0 {
		goto L276
	} else {
		goto L287
	}
L287:
	;
	v804 = v685
	goto L273
L288:
	;
	v694 = int32(0)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v695 <= v694 {
		goto L276
	} else {
		goto L289
	}
L289:
	;
	v698 = v694
	goto L290
L290:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v691)+12))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v702+v698<<(uint(int32(2))%32))))
	v707 = F_path_is_reparameterizable_by_child(m, v706, l1)
	mBase = m.M
	if v707 != 0 {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	v810 = int32(0)
	goto L272
L292:
	;
	v709 = v698 + int32(1)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v691)+4))
	if v709 < v710 {
		v698 = v709
		goto L290
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	goto L291
L295:
	;
	goto L276
L296:
	;
	v716 = int32(0)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v717 <= v716 {
		goto L276
	} else {
		goto L297
	}
L297:
	;
	v720 = v716
	goto L298
L298:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v724+v720<<(uint(int32(2))%32))))
	v729 = F_path_is_reparameterizable_by_child(m, v728, l1)
	mBase = m.M
	if v729 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v810 = int32(0)
	goto L272
L300:
	;
	v731 = v720 + int32(1)
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v713)+4))
	if v731 < v732 {
		v720 = v731
		goto L298
	} else {
		goto L303
	}
L301:
	;
	goto L302
L302:
	;
	goto L299
L303:
	;
	goto L276
L304:
	;
	v738 = F_path_is_reparameterizable_by_child(m, v735, l1)
	mBase = m.M
	if v738 != 0 {
		goto L276
	} else {
		goto L305
	}
L305:
	;
	v804 = v685
	goto L273
L306:
	;
	v742 = int32(0)
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v739)+4))
	if v743 <= v742 {
		goto L276
	} else {
		goto L307
	}
L307:
	;
	v746 = v742
	goto L308
L308:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v739)+12))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v750+v746<<(uint(int32(2))%32))))
	v755 = F_path_is_reparameterizable_by_child(m, v754, l1)
	mBase = m.M
	if v755 != 0 {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v810 = int32(0)
	goto L272
L310:
	;
	v757 = v746 + int32(1)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v739)+4))
	if v757 < v758 {
		v746 = v757
		goto L308
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	goto L309
L313:
	;
	goto L276
L314:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v673)+84))
	v766 = F_path_is_reparameterizable_by_child(m, v765, l1)
	mBase = m.M
	if v766 != 0 {
		goto L276
	} else {
		goto L315
	}
L315:
	;
	v804 = v685
	goto L273
L316:
	;
	v770 = int32(0)
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	if v771 <= v770 {
		goto L276
	} else {
		goto L317
	}
L317:
	;
	v774 = v770
	goto L318
L318:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v767)+12))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v778+v774<<(uint(int32(2))%32))))
	v783 = F_path_is_reparameterizable_by_child(m, v782, l1)
	mBase = m.M
	if v783 != 0 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v810 = int32(0)
	goto L272
L320:
	;
	v785 = v774 + int32(1)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	if v785 < v786 {
		v774 = v785
		goto L318
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	goto L319
L323:
	;
	goto L276
L324:
	;
	v804 = v685
	goto L273
L325:
	;
	v804 = v685
	goto L273
L326:
	;
	goto L276
L327:
	;
	v812 = v665 + int32(1)
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v658)+4))
	if v812 < v813 {
		v665 = v812
		goto L270
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	goto L271
L330:
	;
	goto L18
L331:
	;
	if v954 == int32(0) {
		v1677 = v58
		goto L1
	} else {
		goto L386
	}
L332:
	;
	v954 = v948
	goto L331
L333:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v821)+4))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v826 = F_bms_overlap(m, v824, v825)
	mBase = m.M
	if v826 == int32(0) {
		v948 = v820
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v829 = int32(0)
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v817)))
	switch v830 - int32(279) {
	case 0, 1:
		goto L335
	default:
		v948 = v829
		goto L332
	case 3:
		goto L345
	case 4:
		goto L344
	case 5:
		goto L343
	case 9:
		goto L342
	case 10:
		goto L341
	case 11:
		goto L339
	case 14:
		goto L338
	case 15:
		goto L337
	case 17:
		goto L336
	case 19, 20, 21:
		goto L340
	}
L335:
	;
	v948 = int32(1)
	goto L332
L336:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	v938 = F_path_is_reparameterizable_by_child(m, v937, l1)
	mBase = m.M
	if v938 == int32(0) {
		v948 = v829
		goto L332
	} else {
		goto L385
	}
L337:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	v936 = F_path_is_reparameterizable_by_child(m, v935, l1)
	mBase = m.M
	if v936 != 0 {
		goto L335
	} else {
		goto L384
	}
L338:
	;
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	v934 = F_path_is_reparameterizable_by_child(m, v933, l1)
	mBase = m.M
	if v934 != 0 {
		goto L335
	} else {
		goto L383
	}
L339:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	if v911 == int32(0) {
		goto L335
	} else {
		goto L375
	}
L340:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v817)+80))
	v906 = F_path_is_reparameterizable_by_child(m, v905, l1)
	mBase = m.M
	if v906 == int32(0) {
		v948 = v829
		goto L332
	} else {
		goto L373
	}
L341:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v817)+76))
	if v883 == int32(0) {
		goto L335
	} else {
		goto L365
	}
L342:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	if v879 == int32(0) {
		goto L335
	} else {
		goto L363
	}
L343:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	if v857 == int32(0) {
		goto L335
	} else {
		goto L355
	}
L344:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	if v835 == int32(0) {
		goto L335
	} else {
		goto L347
	}
L345:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v817)+72))
	v834 = F_path_is_reparameterizable_by_child(m, v833, l1)
	mBase = m.M
	if v834 != 0 {
		goto L335
	} else {
		goto L346
	}
L346:
	;
	v948 = v829
	goto L332
L347:
	;
	v838 = int32(0)
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	if v839 <= v838 {
		goto L335
	} else {
		goto L348
	}
L348:
	;
	v842 = v838
	goto L349
L349:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v835)+12))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v846+v842<<(uint(int32(2))%32))))
	v851 = F_path_is_reparameterizable_by_child(m, v850, l1)
	mBase = m.M
	if v851 != 0 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v954 = int32(0)
	goto L331
L351:
	;
	v853 = v842 + int32(1)
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v835)+4))
	if v853 < v854 {
		v842 = v853
		goto L349
	} else {
		goto L354
	}
L352:
	;
	goto L353
L353:
	;
	goto L350
L354:
	;
	goto L335
L355:
	;
	v860 = int32(0)
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v861 <= v860 {
		goto L335
	} else {
		goto L356
	}
L356:
	;
	v864 = v860
	goto L357
L357:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v857)+12))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v868+v864<<(uint(int32(2))%32))))
	v873 = F_path_is_reparameterizable_by_child(m, v872, l1)
	mBase = m.M
	if v873 != 0 {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	v954 = int32(0)
	goto L331
L359:
	;
	v875 = v864 + int32(1)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v857)+4))
	if v875 < v876 {
		v864 = v875
		goto L357
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	goto L358
L362:
	;
	goto L335
L363:
	;
	v882 = F_path_is_reparameterizable_by_child(m, v879, l1)
	mBase = m.M
	if v882 != 0 {
		goto L335
	} else {
		goto L364
	}
L364:
	;
	v948 = v829
	goto L332
L365:
	;
	v886 = int32(0)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	if v887 <= v886 {
		goto L335
	} else {
		goto L366
	}
L366:
	;
	v890 = v886
	goto L367
L367:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v883)+12))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v894+v890<<(uint(int32(2))%32))))
	v899 = F_path_is_reparameterizable_by_child(m, v898, l1)
	mBase = m.M
	if v899 != 0 {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	v954 = int32(0)
	goto L331
L369:
	;
	v901 = v890 + int32(1)
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v883)+4))
	if v901 < v902 {
		v890 = v901
		goto L367
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	goto L368
L372:
	;
	goto L335
L373:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v817)+84))
	v910 = F_path_is_reparameterizable_by_child(m, v909, l1)
	mBase = m.M
	if v910 != 0 {
		goto L335
	} else {
		goto L374
	}
L374:
	;
	v948 = v829
	goto L332
L375:
	;
	v914 = int32(0)
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v911)+4))
	if v915 <= v914 {
		goto L335
	} else {
		goto L376
	}
L376:
	;
	v918 = v914
	goto L377
L377:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v911)+12))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v922+v918<<(uint(int32(2))%32))))
	v927 = F_path_is_reparameterizable_by_child(m, v926, l1)
	mBase = m.M
	if v927 != 0 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v954 = int32(0)
	goto L331
L379:
	;
	v929 = v918 + int32(1)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v911)+4))
	if v929 < v930 {
		v918 = v929
		goto L377
	} else {
		goto L382
	}
L380:
	;
	goto L381
L381:
	;
	goto L378
L382:
	;
	goto L335
L383:
	;
	v948 = v829
	goto L332
L384:
	;
	v948 = v829
	goto L332
L385:
	;
	goto L335
L386:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v960 = int32(1)
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v957)+16))
	if v961 == int32(0) {
		v1088 = v960
		goto L388
	} else {
		goto L389
	}
L387:
	;
	if v1094 != 0 {
		goto L18
	} else {
		goto L442
	}
L388:
	;
	v1094 = v1088
	goto L387
L389:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v961)+4))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v966 = F_bms_overlap(m, v964, v965)
	mBase = m.M
	if v966 == int32(0) {
		v1088 = v960
		goto L388
	} else {
		goto L390
	}
L390:
	;
	v969 = int32(0)
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v957)))
	switch v970 - int32(279) {
	case 0, 1:
		goto L391
	default:
		v1088 = v969
		goto L388
	case 3:
		goto L401
	case 4:
		goto L400
	case 5:
		goto L399
	case 9:
		goto L398
	case 10:
		goto L397
	case 11:
		goto L395
	case 14:
		goto L394
	case 15:
		goto L393
	case 17:
		goto L392
	case 19, 20, 21:
		goto L396
	}
L391:
	;
	v1088 = int32(1)
	goto L388
L392:
	;
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	v1078 = F_path_is_reparameterizable_by_child(m, v1077, l1)
	mBase = m.M
	if v1078 == int32(0) {
		v1088 = v969
		goto L388
	} else {
		goto L441
	}
L393:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	v1076 = F_path_is_reparameterizable_by_child(m, v1075, l1)
	mBase = m.M
	if v1076 != 0 {
		goto L391
	} else {
		goto L440
	}
L394:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	v1074 = F_path_is_reparameterizable_by_child(m, v1073, l1)
	mBase = m.M
	if v1074 != 0 {
		goto L391
	} else {
		goto L439
	}
L395:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	if v1051 == int32(0) {
		goto L391
	} else {
		goto L431
	}
L396:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v957)+80))
	v1046 = F_path_is_reparameterizable_by_child(m, v1045, l1)
	mBase = m.M
	if v1046 == int32(0) {
		v1088 = v969
		goto L388
	} else {
		goto L429
	}
L397:
	;
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v957)+76))
	if v1023 == int32(0) {
		goto L391
	} else {
		goto L421
	}
L398:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	if v1019 == int32(0) {
		goto L391
	} else {
		goto L419
	}
L399:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	if v997 == int32(0) {
		goto L391
	} else {
		goto L411
	}
L400:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	if v975 == int32(0) {
		goto L391
	} else {
		goto L403
	}
L401:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v957)+72))
	v974 = F_path_is_reparameterizable_by_child(m, v973, l1)
	mBase = m.M
	if v974 != 0 {
		goto L391
	} else {
		goto L402
	}
L402:
	;
	v1088 = v969
	goto L388
L403:
	;
	v978 = int32(0)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v975)+4))
	if v979 <= v978 {
		goto L391
	} else {
		goto L404
	}
L404:
	;
	v982 = v978
	goto L405
L405:
	;
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v975)+12))
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v986+v982<<(uint(int32(2))%32))))
	v991 = F_path_is_reparameterizable_by_child(m, v990, l1)
	mBase = m.M
	if v991 != 0 {
		goto L407
	} else {
		goto L408
	}
L406:
	;
	v1094 = int32(0)
	goto L387
L407:
	;
	v993 = v982 + int32(1)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v975)+4))
	if v993 < v994 {
		v982 = v993
		goto L405
	} else {
		goto L410
	}
L408:
	;
	goto L409
L409:
	;
	goto L406
L410:
	;
	goto L391
L411:
	;
	v1000 = int32(0)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v997)+4))
	if v1001 <= v1000 {
		goto L391
	} else {
		goto L412
	}
L412:
	;
	v1004 = v1000
	goto L413
L413:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v997)+12))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1008+v1004<<(uint(int32(2))%32))))
	v1013 = F_path_is_reparameterizable_by_child(m, v1012, l1)
	mBase = m.M
	if v1013 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	v1094 = int32(0)
	goto L387
L415:
	;
	v1015 = v1004 + int32(1)
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v997)+4))
	if v1015 < v1016 {
		v1004 = v1015
		goto L413
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	goto L414
L418:
	;
	goto L391
L419:
	;
	v1022 = F_path_is_reparameterizable_by_child(m, v1019, l1)
	mBase = m.M
	if v1022 != 0 {
		goto L391
	} else {
		goto L420
	}
L420:
	;
	v1088 = v969
	goto L388
L421:
	;
	v1026 = int32(0)
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+4))
	if v1027 <= v1026 {
		goto L391
	} else {
		goto L422
	}
L422:
	;
	v1030 = v1026
	goto L423
L423:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+12))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1034+v1030<<(uint(int32(2))%32))))
	v1039 = F_path_is_reparameterizable_by_child(m, v1038, l1)
	mBase = m.M
	if v1039 != 0 {
		goto L425
	} else {
		goto L426
	}
L424:
	;
	v1094 = int32(0)
	goto L387
L425:
	;
	v1041 = v1030 + int32(1)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+4))
	if v1041 < v1042 {
		v1030 = v1041
		goto L423
	} else {
		goto L428
	}
L426:
	;
	goto L427
L427:
	;
	goto L424
L428:
	;
	goto L391
L429:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v957)+84))
	v1050 = F_path_is_reparameterizable_by_child(m, v1049, l1)
	mBase = m.M
	if v1050 != 0 {
		goto L391
	} else {
		goto L430
	}
L430:
	;
	v1088 = v969
	goto L388
L431:
	;
	v1054 = int32(0)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1055 <= v1054 {
		goto L391
	} else {
		goto L432
	}
L432:
	;
	v1058 = v1054
	goto L433
L433:
	;
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+12))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1062+v1058<<(uint(int32(2))%32))))
	v1067 = F_path_is_reparameterizable_by_child(m, v1066, l1)
	mBase = m.M
	if v1067 != 0 {
		goto L435
	} else {
		goto L436
	}
L434:
	;
	v1094 = int32(0)
	goto L387
L435:
	;
	v1069 = v1058 + int32(1)
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1069 < v1070 {
		v1058 = v1069
		goto L433
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	goto L434
L438:
	;
	goto L391
L439:
	;
	v1088 = v969
	goto L388
L440:
	;
	v1088 = v969
	goto L388
L441:
	;
	goto L391
L442:
	;
	v1677 = v58
	goto L1
L443:
	;
	v1098 = int32(0)
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if v1099 <= v1098 {
		goto L18
	} else {
		goto L444
	}
L444:
	;
	v1102 = v1098
	goto L445
L445:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+12))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v1102<<(uint(int32(2))%32))))
	v1113 = int32(1)
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+16))
	if v1114 == int32(0) {
		v1241 = v1113
		goto L448
	} else {
		goto L449
	}
L446:
	;
	return int32(0)
L447:
	;
	if v1247 != 0 {
		goto L502
	} else {
		goto L503
	}
L448:
	;
	v1247 = v1241
	goto L447
L449:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1114)+4))
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1119 = F_bms_overlap(m, v1117, v1118)
	mBase = m.M
	if v1119 == int32(0) {
		v1241 = v1113
		goto L448
	} else {
		goto L450
	}
L450:
	;
	v1122 = int32(0)
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1110)))
	switch v1123 - int32(279) {
	case 0, 1:
		goto L451
	default:
		v1241 = v1122
		goto L448
	case 3:
		goto L461
	case 4:
		goto L460
	case 5:
		goto L459
	case 9:
		goto L458
	case 10:
		goto L457
	case 11:
		goto L455
	case 14:
		goto L454
	case 15:
		goto L453
	case 17:
		goto L452
	case 19, 20, 21:
		goto L456
	}
L451:
	;
	v1241 = int32(1)
	goto L448
L452:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	v1231 = F_path_is_reparameterizable_by_child(m, v1230, l1)
	mBase = m.M
	if v1231 == int32(0) {
		v1241 = v1122
		goto L448
	} else {
		goto L501
	}
L453:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	v1229 = F_path_is_reparameterizable_by_child(m, v1228, l1)
	mBase = m.M
	if v1229 != 0 {
		goto L451
	} else {
		goto L500
	}
L454:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	v1227 = F_path_is_reparameterizable_by_child(m, v1226, l1)
	mBase = m.M
	if v1227 != 0 {
		goto L451
	} else {
		goto L499
	}
L455:
	;
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	if v1204 == int32(0) {
		goto L451
	} else {
		goto L491
	}
L456:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+80))
	v1199 = F_path_is_reparameterizable_by_child(m, v1198, l1)
	mBase = m.M
	if v1199 == int32(0) {
		v1241 = v1122
		goto L448
	} else {
		goto L489
	}
L457:
	;
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+76))
	if v1176 == int32(0) {
		goto L451
	} else {
		goto L481
	}
L458:
	;
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	if v1172 == int32(0) {
		goto L451
	} else {
		goto L479
	}
L459:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	if v1150 == int32(0) {
		goto L451
	} else {
		goto L471
	}
L460:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	if v1128 == int32(0) {
		goto L451
	} else {
		goto L463
	}
L461:
	;
	v1126 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+72))
	v1127 = F_path_is_reparameterizable_by_child(m, v1126, l1)
	mBase = m.M
	if v1127 != 0 {
		goto L451
	} else {
		goto L462
	}
L462:
	;
	v1241 = v1122
	goto L448
L463:
	;
	v1131 = int32(0)
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+4))
	if v1132 <= v1131 {
		goto L451
	} else {
		goto L464
	}
L464:
	;
	v1135 = v1131
	goto L465
L465:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+12))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1139+v1135<<(uint(int32(2))%32))))
	v1144 = F_path_is_reparameterizable_by_child(m, v1143, l1)
	mBase = m.M
	if v1144 != 0 {
		goto L467
	} else {
		goto L468
	}
L466:
	;
	v1247 = int32(0)
	goto L447
L467:
	;
	v1146 = v1135 + int32(1)
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1128)+4))
	if v1146 < v1147 {
		v1135 = v1146
		goto L465
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	goto L466
L470:
	;
	goto L451
L471:
	;
	v1153 = int32(0)
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+4))
	if v1154 <= v1153 {
		goto L451
	} else {
		goto L472
	}
L472:
	;
	v1157 = v1153
	goto L473
L473:
	;
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+12))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1161+v1157<<(uint(int32(2))%32))))
	v1166 = F_path_is_reparameterizable_by_child(m, v1165, l1)
	mBase = m.M
	if v1166 != 0 {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	v1247 = int32(0)
	goto L447
L475:
	;
	v1168 = v1157 + int32(1)
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+4))
	if v1168 < v1169 {
		v1157 = v1168
		goto L473
	} else {
		goto L478
	}
L476:
	;
	goto L477
L477:
	;
	goto L474
L478:
	;
	goto L451
L479:
	;
	v1175 = F_path_is_reparameterizable_by_child(m, v1172, l1)
	mBase = m.M
	if v1175 != 0 {
		goto L451
	} else {
		goto L480
	}
L480:
	;
	v1241 = v1122
	goto L448
L481:
	;
	v1179 = int32(0)
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+4))
	if v1180 <= v1179 {
		goto L451
	} else {
		goto L482
	}
L482:
	;
	v1183 = v1179
	goto L483
L483:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+12))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1187+v1183<<(uint(int32(2))%32))))
	v1192 = F_path_is_reparameterizable_by_child(m, v1191, l1)
	mBase = m.M
	if v1192 != 0 {
		goto L485
	} else {
		goto L486
	}
L484:
	;
	v1247 = int32(0)
	goto L447
L485:
	;
	v1194 = v1183 + int32(1)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1176)+4))
	if v1194 < v1195 {
		v1183 = v1194
		goto L483
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	goto L484
L488:
	;
	goto L451
L489:
	;
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+84))
	v1203 = F_path_is_reparameterizable_by_child(m, v1202, l1)
	mBase = m.M
	if v1203 != 0 {
		goto L451
	} else {
		goto L490
	}
L490:
	;
	v1241 = v1122
	goto L448
L491:
	;
	v1207 = int32(0)
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+4))
	if v1208 <= v1207 {
		goto L451
	} else {
		goto L492
	}
L492:
	;
	v1211 = v1207
	goto L493
L493:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+12))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1215+v1211<<(uint(int32(2))%32))))
	v1220 = F_path_is_reparameterizable_by_child(m, v1219, l1)
	mBase = m.M
	if v1220 != 0 {
		goto L495
	} else {
		goto L496
	}
L494:
	;
	v1247 = int32(0)
	goto L447
L495:
	;
	v1222 = v1211 + int32(1)
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v1204)+4))
	if v1222 < v1223 {
		v1211 = v1222
		goto L493
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	goto L494
L498:
	;
	goto L451
L499:
	;
	v1241 = v1122
	goto L448
L500:
	;
	v1241 = v1122
	goto L448
L501:
	;
	goto L451
L502:
	;
	v1249 = v1102 + int32(1)
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if v1249 < v1250 {
		v1102 = v1249
		goto L445
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	goto L446
L505:
	;
	goto L18
L506:
	;
	if v1391 != 0 {
		goto L18
	} else {
		goto L561
	}
L507:
	;
	v1391 = v1385
	goto L506
L508:
	;
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(v1258)+4))
	v1262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1263 = F_bms_overlap(m, v1261, v1262)
	mBase = m.M
	if v1263 == int32(0) {
		v1385 = v1257
		goto L507
	} else {
		goto L509
	}
L509:
	;
	v1266 = int32(0)
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v1254)))
	switch v1267 - int32(279) {
	case 0, 1:
		goto L510
	default:
		v1385 = v1266
		goto L507
	case 3:
		goto L520
	case 4:
		goto L519
	case 5:
		goto L518
	case 9:
		goto L517
	case 10:
		goto L516
	case 11:
		goto L514
	case 14:
		goto L513
	case 15:
		goto L512
	case 17:
		goto L511
	case 19, 20, 21:
		goto L515
	}
L510:
	;
	v1385 = int32(1)
	goto L507
L511:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	v1375 = F_path_is_reparameterizable_by_child(m, v1374, l1)
	mBase = m.M
	if v1375 == int32(0) {
		v1385 = v1266
		goto L507
	} else {
		goto L560
	}
L512:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	v1373 = F_path_is_reparameterizable_by_child(m, v1372, l1)
	mBase = m.M
	if v1373 != 0 {
		goto L510
	} else {
		goto L559
	}
L513:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	v1371 = F_path_is_reparameterizable_by_child(m, v1370, l1)
	mBase = m.M
	if v1371 != 0 {
		goto L510
	} else {
		goto L558
	}
L514:
	;
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	if v1348 == int32(0) {
		goto L510
	} else {
		goto L550
	}
L515:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+80))
	v1343 = F_path_is_reparameterizable_by_child(m, v1342, l1)
	mBase = m.M
	if v1343 == int32(0) {
		v1385 = v1266
		goto L507
	} else {
		goto L548
	}
L516:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+76))
	if v1320 == int32(0) {
		goto L510
	} else {
		goto L540
	}
L517:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	if v1316 == int32(0) {
		goto L510
	} else {
		goto L538
	}
L518:
	;
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	if v1294 == int32(0) {
		goto L510
	} else {
		goto L530
	}
L519:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	if v1272 == int32(0) {
		goto L510
	} else {
		goto L522
	}
L520:
	;
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+72))
	v1271 = F_path_is_reparameterizable_by_child(m, v1270, l1)
	mBase = m.M
	if v1271 != 0 {
		goto L510
	} else {
		goto L521
	}
L521:
	;
	v1385 = v1266
	goto L507
L522:
	;
	v1275 = int32(0)
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	if v1276 <= v1275 {
		goto L510
	} else {
		goto L523
	}
L523:
	;
	v1279 = v1275
	goto L524
L524:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+12))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1283+v1279<<(uint(int32(2))%32))))
	v1288 = F_path_is_reparameterizable_by_child(m, v1287, l1)
	mBase = m.M
	if v1288 != 0 {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	v1391 = int32(0)
	goto L506
L526:
	;
	v1290 = v1279 + int32(1)
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1272)+4))
	if v1290 < v1291 {
		v1279 = v1290
		goto L524
	} else {
		goto L529
	}
L527:
	;
	goto L528
L528:
	;
	goto L525
L529:
	;
	goto L510
L530:
	;
	v1297 = int32(0)
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+4))
	if v1298 <= v1297 {
		goto L510
	} else {
		goto L531
	}
L531:
	;
	v1301 = v1297
	goto L532
L532:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+12))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1305+v1301<<(uint(int32(2))%32))))
	v1310 = F_path_is_reparameterizable_by_child(m, v1309, l1)
	mBase = m.M
	if v1310 != 0 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	v1391 = int32(0)
	goto L506
L534:
	;
	v1312 = v1301 + int32(1)
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1294)+4))
	if v1312 < v1313 {
		v1301 = v1312
		goto L532
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	goto L533
L537:
	;
	goto L510
L538:
	;
	v1319 = F_path_is_reparameterizable_by_child(m, v1316, l1)
	mBase = m.M
	if v1319 != 0 {
		goto L510
	} else {
		goto L539
	}
L539:
	;
	v1385 = v1266
	goto L507
L540:
	;
	v1323 = int32(0)
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+4))
	if v1324 <= v1323 {
		goto L510
	} else {
		goto L541
	}
L541:
	;
	v1327 = v1323
	goto L542
L542:
	;
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+12))
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1331+v1327<<(uint(int32(2))%32))))
	v1336 = F_path_is_reparameterizable_by_child(m, v1335, l1)
	mBase = m.M
	if v1336 != 0 {
		goto L544
	} else {
		goto L545
	}
L543:
	;
	v1391 = int32(0)
	goto L506
L544:
	;
	v1338 = v1327 + int32(1)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1320)+4))
	if v1338 < v1339 {
		v1327 = v1338
		goto L542
	} else {
		goto L547
	}
L545:
	;
	goto L546
L546:
	;
	goto L543
L547:
	;
	goto L510
L548:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+84))
	v1347 = F_path_is_reparameterizable_by_child(m, v1346, l1)
	mBase = m.M
	if v1347 != 0 {
		goto L510
	} else {
		goto L549
	}
L549:
	;
	v1385 = v1266
	goto L507
L550:
	;
	v1351 = int32(0)
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+4))
	if v1352 <= v1351 {
		goto L510
	} else {
		goto L551
	}
L551:
	;
	v1355 = v1351
	goto L552
L552:
	;
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+12))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v1359+v1355<<(uint(int32(2))%32))))
	v1364 = F_path_is_reparameterizable_by_child(m, v1363, l1)
	mBase = m.M
	if v1364 != 0 {
		goto L554
	} else {
		goto L555
	}
L553:
	;
	v1391 = int32(0)
	goto L506
L554:
	;
	v1366 = v1355 + int32(1)
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1348)+4))
	if v1366 < v1367 {
		v1355 = v1366
		goto L552
	} else {
		goto L557
	}
L555:
	;
	goto L556
L556:
	;
	goto L553
L557:
	;
	goto L510
L558:
	;
	v1385 = v1266
	goto L507
L559:
	;
	v1385 = v1266
	goto L507
L560:
	;
	goto L510
L561:
	;
	v1677 = v58
	goto L1
L562:
	;
	if v1529 != 0 {
		goto L18
	} else {
		goto L617
	}
L563:
	;
	v1529 = v1523
	goto L562
L564:
	;
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1396)+4))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1401 = F_bms_overlap(m, v1399, v1400)
	mBase = m.M
	if v1401 == int32(0) {
		v1523 = v1395
		goto L563
	} else {
		goto L565
	}
L565:
	;
	v1404 = int32(0)
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	switch v1405 - int32(279) {
	case 0, 1:
		goto L566
	default:
		v1523 = v1404
		goto L563
	case 3:
		goto L576
	case 4:
		goto L575
	case 5:
		goto L574
	case 9:
		goto L573
	case 10:
		goto L572
	case 11:
		goto L570
	case 14:
		goto L569
	case 15:
		goto L568
	case 17:
		goto L567
	case 19, 20, 21:
		goto L571
	}
L566:
	;
	v1523 = int32(1)
	goto L563
L567:
	;
	v1512 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	v1513 = F_path_is_reparameterizable_by_child(m, v1512, l1)
	mBase = m.M
	if v1513 == int32(0) {
		v1523 = v1404
		goto L563
	} else {
		goto L616
	}
L568:
	;
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	v1511 = F_path_is_reparameterizable_by_child(m, v1510, l1)
	mBase = m.M
	if v1511 != 0 {
		goto L566
	} else {
		goto L615
	}
L569:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	v1509 = F_path_is_reparameterizable_by_child(m, v1508, l1)
	mBase = m.M
	if v1509 != 0 {
		goto L566
	} else {
		goto L614
	}
L570:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	if v1486 == int32(0) {
		goto L566
	} else {
		goto L606
	}
L571:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+80))
	v1481 = F_path_is_reparameterizable_by_child(m, v1480, l1)
	mBase = m.M
	if v1481 == int32(0) {
		v1523 = v1404
		goto L563
	} else {
		goto L604
	}
L572:
	;
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+76))
	if v1458 == int32(0) {
		goto L566
	} else {
		goto L596
	}
L573:
	;
	v1454 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	if v1454 == int32(0) {
		goto L566
	} else {
		goto L594
	}
L574:
	;
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	if v1432 == int32(0) {
		goto L566
	} else {
		goto L586
	}
L575:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	if v1410 == int32(0) {
		goto L566
	} else {
		goto L578
	}
L576:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	v1409 = F_path_is_reparameterizable_by_child(m, v1408, l1)
	mBase = m.M
	if v1409 != 0 {
		goto L566
	} else {
		goto L577
	}
L577:
	;
	v1523 = v1404
	goto L563
L578:
	;
	v1413 = int32(0)
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+4))
	if v1414 <= v1413 {
		goto L566
	} else {
		goto L579
	}
L579:
	;
	v1417 = v1413
	goto L580
L580:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+12))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v1421+v1417<<(uint(int32(2))%32))))
	v1426 = F_path_is_reparameterizable_by_child(m, v1425, l1)
	mBase = m.M
	if v1426 != 0 {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	v1529 = int32(0)
	goto L562
L582:
	;
	v1428 = v1417 + int32(1)
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1410)+4))
	if v1428 < v1429 {
		v1417 = v1428
		goto L580
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	goto L581
L585:
	;
	goto L566
L586:
	;
	v1435 = int32(0)
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+4))
	if v1436 <= v1435 {
		goto L566
	} else {
		goto L587
	}
L587:
	;
	v1439 = v1435
	goto L588
L588:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+12))
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1443+v1439<<(uint(int32(2))%32))))
	v1448 = F_path_is_reparameterizable_by_child(m, v1447, l1)
	mBase = m.M
	if v1448 != 0 {
		goto L590
	} else {
		goto L591
	}
L589:
	;
	v1529 = int32(0)
	goto L562
L590:
	;
	v1450 = v1439 + int32(1)
	v1451 = *(*int32)(unsafe.Add(mBase, uint32(v1432)+4))
	if v1450 < v1451 {
		v1439 = v1450
		goto L588
	} else {
		goto L593
	}
L591:
	;
	goto L592
L592:
	;
	goto L589
L593:
	;
	goto L566
L594:
	;
	v1457 = F_path_is_reparameterizable_by_child(m, v1454, l1)
	mBase = m.M
	if v1457 != 0 {
		goto L566
	} else {
		goto L595
	}
L595:
	;
	v1523 = v1404
	goto L563
L596:
	;
	v1461 = int32(0)
	v1462 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+4))
	if v1462 <= v1461 {
		goto L566
	} else {
		goto L597
	}
L597:
	;
	v1465 = v1461
	goto L598
L598:
	;
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+12))
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1469+v1465<<(uint(int32(2))%32))))
	v1474 = F_path_is_reparameterizable_by_child(m, v1473, l1)
	mBase = m.M
	if v1474 != 0 {
		goto L600
	} else {
		goto L601
	}
L599:
	;
	v1529 = int32(0)
	goto L562
L600:
	;
	v1476 = v1465 + int32(1)
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1458)+4))
	if v1476 < v1477 {
		v1465 = v1476
		goto L598
	} else {
		goto L603
	}
L601:
	;
	goto L602
L602:
	;
	goto L599
L603:
	;
	goto L566
L604:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+84))
	v1485 = F_path_is_reparameterizable_by_child(m, v1484, l1)
	mBase = m.M
	if v1485 != 0 {
		goto L566
	} else {
		goto L605
	}
L605:
	;
	v1523 = v1404
	goto L563
L606:
	;
	v1489 = int32(0)
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+4))
	if v1490 <= v1489 {
		goto L566
	} else {
		goto L607
	}
L607:
	;
	v1493 = v1489
	goto L608
L608:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+12))
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1497+v1493<<(uint(int32(2))%32))))
	v1502 = F_path_is_reparameterizable_by_child(m, v1501, l1)
	mBase = m.M
	if v1502 != 0 {
		goto L610
	} else {
		goto L611
	}
L609:
	;
	v1529 = int32(0)
	goto L562
L610:
	;
	v1504 = v1493 + int32(1)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1486)+4))
	if v1504 < v1505 {
		v1493 = v1504
		goto L608
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	goto L609
L613:
	;
	goto L566
L614:
	;
	v1523 = v1404
	goto L563
L615:
	;
	v1523 = v1404
	goto L563
L616:
	;
	goto L566
L617:
	;
	v1677 = v58
	goto L1
L618:
	;
	if v1667 == int32(0) {
		v1677 = v58
		goto L1
	} else {
		goto L673
	}
L619:
	;
	v1667 = v1661
	goto L618
L620:
	;
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1534)+4))
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	v1539 = F_bms_overlap(m, v1537, v1538)
	mBase = m.M
	if v1539 == int32(0) {
		v1661 = v1533
		goto L619
	} else {
		goto L621
	}
L621:
	;
	v1542 = int32(0)
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1530)))
	switch v1543 - int32(279) {
	case 0, 1:
		goto L622
	default:
		v1661 = v1542
		goto L619
	case 3:
		goto L632
	case 4:
		goto L631
	case 5:
		goto L630
	case 9:
		goto L629
	case 10:
		goto L628
	case 11:
		goto L626
	case 14:
		goto L625
	case 15:
		goto L624
	case 17:
		goto L623
	case 19, 20, 21:
		goto L627
	}
L622:
	;
	v1661 = int32(1)
	goto L619
L623:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	v1651 = F_path_is_reparameterizable_by_child(m, v1650, l1)
	mBase = m.M
	if v1651 == int32(0) {
		v1661 = v1542
		goto L619
	} else {
		goto L672
	}
L624:
	;
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	v1649 = F_path_is_reparameterizable_by_child(m, v1648, l1)
	mBase = m.M
	if v1649 != 0 {
		goto L622
	} else {
		goto L671
	}
L625:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	v1647 = F_path_is_reparameterizable_by_child(m, v1646, l1)
	mBase = m.M
	if v1647 != 0 {
		goto L622
	} else {
		goto L670
	}
L626:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	if v1624 == int32(0) {
		goto L622
	} else {
		goto L662
	}
L627:
	;
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+80))
	v1619 = F_path_is_reparameterizable_by_child(m, v1618, l1)
	mBase = m.M
	if v1619 == int32(0) {
		v1661 = v1542
		goto L619
	} else {
		goto L660
	}
L628:
	;
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+76))
	if v1596 == int32(0) {
		goto L622
	} else {
		goto L652
	}
L629:
	;
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	if v1592 == int32(0) {
		goto L622
	} else {
		goto L650
	}
L630:
	;
	v1570 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	if v1570 == int32(0) {
		goto L622
	} else {
		goto L642
	}
L631:
	;
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	if v1548 == int32(0) {
		goto L622
	} else {
		goto L634
	}
L632:
	;
	v1546 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+72))
	v1547 = F_path_is_reparameterizable_by_child(m, v1546, l1)
	mBase = m.M
	if v1547 != 0 {
		goto L622
	} else {
		goto L633
	}
L633:
	;
	v1661 = v1542
	goto L619
L634:
	;
	v1551 = int32(0)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+4))
	if v1552 <= v1551 {
		goto L622
	} else {
		goto L635
	}
L635:
	;
	v1555 = v1551
	goto L636
L636:
	;
	v1559 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+12))
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1559+v1555<<(uint(int32(2))%32))))
	v1564 = F_path_is_reparameterizable_by_child(m, v1563, l1)
	mBase = m.M
	if v1564 != 0 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	v1667 = int32(0)
	goto L618
L638:
	;
	v1566 = v1555 + int32(1)
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1548)+4))
	if v1566 < v1567 {
		v1555 = v1566
		goto L636
	} else {
		goto L641
	}
L639:
	;
	goto L640
L640:
	;
	goto L637
L641:
	;
	goto L622
L642:
	;
	v1573 = int32(0)
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+4))
	if v1574 <= v1573 {
		goto L622
	} else {
		goto L643
	}
L643:
	;
	v1577 = v1573
	goto L644
L644:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+12))
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1581+v1577<<(uint(int32(2))%32))))
	v1586 = F_path_is_reparameterizable_by_child(m, v1585, l1)
	mBase = m.M
	if v1586 != 0 {
		goto L646
	} else {
		goto L647
	}
L645:
	;
	v1667 = int32(0)
	goto L618
L646:
	;
	v1588 = v1577 + int32(1)
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v1570)+4))
	if v1588 < v1589 {
		v1577 = v1588
		goto L644
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	goto L645
L649:
	;
	goto L622
L650:
	;
	v1595 = F_path_is_reparameterizable_by_child(m, v1592, l1)
	mBase = m.M
	if v1595 != 0 {
		goto L622
	} else {
		goto L651
	}
L651:
	;
	v1661 = v1542
	goto L619
L652:
	;
	v1599 = int32(0)
	v1600 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+4))
	if v1600 <= v1599 {
		goto L622
	} else {
		goto L653
	}
L653:
	;
	v1603 = v1599
	goto L654
L654:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+12))
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1607+v1603<<(uint(int32(2))%32))))
	v1612 = F_path_is_reparameterizable_by_child(m, v1611, l1)
	mBase = m.M
	if v1612 != 0 {
		goto L656
	} else {
		goto L657
	}
L655:
	;
	v1667 = int32(0)
	goto L618
L656:
	;
	v1614 = v1603 + int32(1)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1596)+4))
	if v1614 < v1615 {
		v1603 = v1614
		goto L654
	} else {
		goto L659
	}
L657:
	;
	goto L658
L658:
	;
	goto L655
L659:
	;
	goto L622
L660:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1530)+84))
	v1623 = F_path_is_reparameterizable_by_child(m, v1622, l1)
	mBase = m.M
	if v1623 != 0 {
		goto L622
	} else {
		goto L661
	}
L661:
	;
	v1661 = v1542
	goto L619
L662:
	;
	v1627 = int32(0)
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+4))
	if v1628 <= v1627 {
		goto L622
	} else {
		goto L663
	}
L663:
	;
	v1631 = v1627
	goto L664
L664:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+12))
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v1635+v1631<<(uint(int32(2))%32))))
	v1640 = F_path_is_reparameterizable_by_child(m, v1639, l1)
	mBase = m.M
	if v1640 != 0 {
		goto L666
	} else {
		goto L667
	}
L665:
	;
	v1667 = int32(0)
	goto L618
L666:
	;
	v1642 = v1631 + int32(1)
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+4))
	if v1642 < v1643 {
		v1631 = v1642
		goto L664
	} else {
		goto L669
	}
L667:
	;
	goto L668
L668:
	;
	goto L665
L669:
	;
	goto L622
L670:
	;
	v1661 = v1542
	goto L619
L671:
	;
	v1661 = v1542
	goto L619
L672:
	;
	goto L622
L673:
	;
	goto L18
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
