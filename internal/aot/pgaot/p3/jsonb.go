package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonbDeepContains(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v180 int64
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	F_check_stack_depth(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = F_JsonbIteratorNext(m, l0, v12+int32(60), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L82
	}
L4:
	;
	m.G0 = v12 + int32(80)
	return v274
L5:
	;
	v26 = F_JsonbIteratorNext(m, l1, v12+int32(40), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v21 != v26 {
		v274 = v3
		goto L4
	} else {
		goto L7
	}
L7:
	;
	switch v21 - int32(4) {
	case 0:
		goto L8
	default:
		goto L3
	case 2:
		goto L9
	}
L8:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+72)))
	if v107 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+64))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v31 < v32 {
		v274 = v3
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v37 = F_JsonbIteratorNext(m, l1, v12+int32(40), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v37 == int32(7) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v274 = int32(1)
	goto L4
L13:
	;
	goto L14
L14:
	;
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v57 = F_getKeyJsonValueFromContainer(m, v52, v53, v54, v12+int32(20))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v274 = int32(1)
	goto L4
L17:
	;
	if v57 == int32(0) {
		v274 = v3
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v64 = F_JsonbIteratorNext(m, l1, v12+int32(40), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v66 != v67 {
		v274 = v3
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if base.B2i32(v66 != int32(32))&base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v66)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v101 = F_JsonbIteratorNext(m, l1, v12+int32(40), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v78 = F_equalsJsonbScalarValue(m, v57, v12+int32(40))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v82 = F_iteratorFromContainer(m, v80, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	if v78 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v274 = v3
	goto L4
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v87 = F_iteratorFromContainer(m, v85, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v87
	v94 = F_JsonbDeepContains(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v94 == int32(0) {
		v274 = v3
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	if v101 != int32(7) {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	goto L16
L33:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)))
	if v110 != int32(1) {
		v274 = v3
		goto L4
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v113 = int32(1)
	v117 = F_JsonbIteratorNext(m, l1, v12+int32(40), v113)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	if v117 == int32(5) {
		v274 = v113
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v126 = v106
	v128 = v3
	goto L39
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v130))&base.B2i32(v130 != int32(32)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v274 = v263
	goto L4
L41:
	;
	v263 = int32(1)
	v267 = F_JsonbIteratorNext(m, l1, v12+int32(40), v263)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L80
	}
L42:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v143 = F_findJsonbValueFromContainer(m, v139, int32(1073741824), v12+int32(40))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v128 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	if v143 != 0 {
		v259 = v126
		v261 = v128
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v274 = int32(0)
	goto L4
L47:
	;
	if v246 != v251 {
		v259 = v251
		v261 = v249
		goto L41
	} else {
		goto L79
	}
L48:
	;
	v240 = int32(0)
	v246 = v240
	v249 = v128
	v251 = v240
	goto L47
L49:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	v207 = int32(0)
	goto L64
L50:
	;
	v148 = int32(0)
	v151 = F_palloc(m, v126*int32(20))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if v126 == int32(0) {
		goto L48
	} else {
		goto L63
	}
L53:
	;
	if v126 == int32(0) {
		v274 = v148
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v159 = v148
	v160 = int32(0)
	goto L55
L55:
	;
	v168 = F_JsonbIteratorNext(m, l0, v12+int32(60), int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	if v184 != 0 {
		v195 = v184
		v199 = v151
		goto L49
	} else {
		goto L62
	}
L57:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	if v170 == int32(18) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v175 = v151 + v159*int32(20)
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v12)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v175))) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v12)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+16)) = v178
	v180 = *(*int64)(unsafe.Add(mBase, uint32(v12)+68))
	*(*int64)(unsafe.Add(mBase, uint32(v175)+8)) = v180
	v184 = v159 + int32(1)
	goto L60
L59:
	;
	v184 = v159
	goto L60
L60:
	;
	v187 = v160 + int32(1)
	if v187 != v126 {
		v159 = v184
		v160 = v187
		goto L55
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	v274 = int32(0)
	goto L4
L63:
	;
	v195 = v126
	v199 = v128
	goto L49
L64:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v199+v207*int32(20))+8))
	v217 = F_iteratorFromContainer(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	v274 = int32(0)
	goto L4
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v217
	v221 = F_iteratorFromContainer(m, v202, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v221
	v228 = F_JsonbDeepContains(m, v12+int32(20), v12+int32(16))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v230 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_pfree(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v233 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L71
L73:
	;
	F_pfree(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v228 != 0 {
		v246 = v207
		v249 = v199
		v251 = v195
		goto L47
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	v237 = v207 + int32(1)
	if v237 != v195 {
		v207 = v237
		goto L64
	} else {
		goto L78
	}
L78:
	;
	goto L65
L79:
	;
	v274 = int32(0)
	goto L4
L80:
	;
	if v267 != int32(5) {
		v126 = v259
		v128 = v261
		goto L39
	} else {
		goto L81
	}
L81:
	;
	goto L40
L82:
	;
	F_errmsg_internal(m, int32(369421), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(498906), int32(1314), int32(149956))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_JsonbHashScalarValue(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
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
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v5 {
	case 0:
		v294 = int32(1)
		v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v295, int32(1)) ^ v294
		return
	case 1:
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v35 = v29 - int32(1636608432)
		if v28&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v29) {
				v144 = v28
				v145 = v29
				v146 = v35
				v147 = v35
				v148 = v35
				for {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
					v151 = v150 + v147
					v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
					v155 = v154 + v148
					v157 = int32(4)
					v159 = v152 + v146 - v155 ^ base.I32_rotl(v155, v157)
					v163 = v151 - v159 ^ base.I32_rotl(v159, int32(6))
					v164 = v155 + v151
					v165 = v159 + v164
					v166 = v163 + v165
					v170 = v164 - v163 ^ base.I32_rotl(v163, int32(8))
					v174 = v165 - v170 ^ base.I32_rotl(v170, int32(16))
					v178 = v166 - v174 ^ base.I32_rotl(v174, int32(19))
					v179 = v170 + v166
					v180 = v174 + v179
					v181 = v178 + v180
					v185 = v179 - v178 ^ base.I32_rotl(v178, v157)
					v186 = int32(12)
					v187 = v144 + v186
					v189 = v145 - v186
					if base.Ui32(int32(11)) < base.Ui32(v189) {
						v144 = v187
						v145 = v189
						v146 = v180
						v147 = v181
						v148 = v185
						continue
					} else {
						break
					}
					break
				}
				v192 = v187
				v193 = v189
				v194 = v180
				v195 = v181
				v196 = v185
			} else {
				v192 = v28
				v193 = v29
				v194 = v35
				v195 = v35
				v196 = v35
			}
			switch v193 - int32(1) {
			case 0:
				v255 = v194
				v256 = v195
				v257 = v196
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 1:
				v248 = v194
				v249 = v195
				v250 = v196
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 2:
				v241 = v194
				v242 = v195
				v243 = v196
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 3:
				v235 = v195
				v236 = v196
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 4:
				v231 = v195
				v232 = v196
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 5:
				v225 = v195
				v226 = v196
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 6:
				v219 = v195
				v220 = v196
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+6)))
				v225 = v221<<(uint(int32(16))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 7:
				v214 = v196
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+7)))
				v219 = v215<<(uint(int32(24))%32) + v195
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+6)))
				v225 = v221<<(uint(int32(16))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 8:
				v209 = v196
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+7)))
				v219 = v215<<(uint(int32(24))%32) + v195
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+6)))
				v225 = v221<<(uint(int32(16))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 9:
				v204 = v196
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+9)))
				v209 = v205<<(uint(int32(16))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+7)))
				v219 = v215<<(uint(int32(24))%32) + v195
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+6)))
				v225 = v221<<(uint(int32(16))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			case 10:
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+10)))
				v204 = v200<<(uint(int32(24))%32) + v196
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+9)))
				v209 = v205<<(uint(int32(16))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+7)))
				v219 = v215<<(uint(int32(24))%32) + v195
				v220 = v214
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+6)))
				v225 = v221<<(uint(int32(16))%32) + v219
				v226 = v220
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+5)))
				v231 = v227<<(uint(int32(8))%32) + v225
				v232 = v226
				v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
				v235 = v231 + v233
				v236 = v232
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+3)))
				v241 = v237<<(uint(int32(24))%32) + v194
				v242 = v235
				v243 = v236
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+2)))
				v248 = v244<<(uint(int32(16))%32) + v241
				v249 = v242
				v250 = v243
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
				v255 = v251<<(uint(int32(8))%32) + v248
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
				v262 = v255 + v258
				v263 = v256
				v264 = v257
			default:
				v262 = v194
				v263 = v195
				v264 = v196
			}
		} else {
			if base.Ui32(v29) < base.Ui32(int32(12)) {
				v90 = v28
				v91 = v29
				v92 = v35
				v93 = v35
				v94 = v35
			} else {
				v42 = v28
				v43 = v29
				v44 = v35
				v45 = v35
				v46 = v35
				for {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					v49 = v48 + v45
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
					v53 = v52 + v46
					v55 = int32(4)
					v57 = v50 + v44 - v53 ^ base.I32_rotl(v53, v55)
					v61 = v49 - v57 ^ base.I32_rotl(v57, int32(6))
					v62 = v53 + v49
					v63 = v57 + v62
					v64 = v61 + v63
					v68 = v62 - v61 ^ base.I32_rotl(v61, int32(8))
					v72 = v63 - v68 ^ base.I32_rotl(v68, int32(16))
					v76 = v64 - v72 ^ base.I32_rotl(v72, int32(19))
					v77 = v68 + v64
					v78 = v72 + v77
					v79 = v76 + v78
					v83 = v77 - v76 ^ base.I32_rotl(v76, v55)
					v84 = int32(12)
					v85 = v42 + v84
					v87 = v43 - v84
					if base.Ui32(int32(11)) < base.Ui32(v87) {
						v42 = v85
						v43 = v87
						v44 = v78
						v45 = v79
						v46 = v83
						continue
					} else {
						break
					}
					break
				}
				v90 = v85
				v91 = v87
				v92 = v78
				v93 = v79
				v94 = v83
			}
			switch v91 - int32(1) {
			case 0:
				v141 = v92
				v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
				v262 = v141 + v142
				v263 = v93
				v264 = v94
			case 1:
				v136 = v92
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
				v141 = v137<<(uint(int32(8))%32) + v136
				v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
				v262 = v141 + v142
				v263 = v93
				v264 = v94
			case 2:
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+2)))
				v136 = v132<<(uint(int32(16))%32) + v92
				v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
				v141 = v137<<(uint(int32(8))%32) + v136
				v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
				v262 = v141 + v142
				v263 = v93
				v264 = v94
			case 3:
				v129 = v93
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v262 = v130 + v92
				v263 = v129
				v264 = v94
			case 4:
				v126 = v93
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
				v129 = v126 + v127
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v262 = v130 + v92
				v263 = v129
				v264 = v94
			case 5:
				v121 = v93
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
				v126 = v122<<(uint(int32(8))%32) + v121
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
				v129 = v126 + v127
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v262 = v130 + v92
				v263 = v129
				v264 = v94
			case 6:
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+6)))
				v121 = v117<<(uint(int32(16))%32) + v93
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+5)))
				v126 = v122<<(uint(int32(8))%32) + v121
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)))
				v129 = v126 + v127
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v262 = v130 + v92
				v263 = v129
				v264 = v94
			case 7:
				v112 = v94
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
				v262 = v113 + v92
				v263 = v115 + v93
				v264 = v112
			case 8:
				v107 = v94
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
				v262 = v113 + v92
				v263 = v115 + v93
				v264 = v112
			case 9:
				v102 = v94
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)))
				v107 = v103<<(uint(int32(16))%32) + v102
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
				v262 = v113 + v92
				v263 = v115 + v93
				v264 = v112
			case 10:
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+10)))
				v102 = v98<<(uint(int32(24))%32) + v94
				v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)))
				v107 = v103<<(uint(int32(16))%32) + v102
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
				v112 = v108<<(uint(int32(8))%32) + v107
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
				v262 = v113 + v92
				v263 = v115 + v93
				v264 = v112
			default:
				v262 = v92
				v263 = v93
				v264 = v94
			}
		}
		v267 = int32(14)
		v269 = v263 ^ v264 - base.I32_rotl(v263, v267)
		v273 = v269 ^ v262 - base.I32_rotl(v269, int32(11))
		v277 = v273 ^ v263 - base.I32_rotl(v273, int32(25))
		v281 = v277 ^ v269 - base.I32_rotl(v277, int32(16))
		v285 = v281 ^ v273 - base.I32_rotl(v281, int32(4))
		v289 = v285 ^ v277 - base.I32_rotl(v285, v267)
		v294 = v289 ^ v281 - base.I32_rotl(v289, int32(24))
		v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v295, int32(1)) ^ v294
		return
	case 2:
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = F_DirectFunctionCall1Coll(m, int32(1345), int32(0), v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v294 = v9
			v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v295, int32(1)) ^ v294
			return
		}
	case 3:
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v13 != 0 {
			v14 = int32(2)
		} else {
			v14 = int32(4)
		}
		v294 = v14
		v295 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v295, int32(1)) ^ v294
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(369516), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errfinish(m, int32(498906), int32(1353), int32(348415))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_JsonbToCStringWorker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	if l0 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = F_makeStringInfo(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v23 = l0
	goto L3
L3:
	;
	if l3 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v23 = v19
	goto L3
L6:
	;
	v26 = int32(1)
	goto L8
L7:
	;
	v26 = int32(2)
	goto L8
L8:
	;
	if l2 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = int32(64)
	goto L11
L10:
	;
	v30 = l2
	goto L11
L11:
	;
	F_enlargeStringInfo(m, v23, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v33 = F_JsonbIteratorInit(m, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v33
	v37 = int32(0)
	v40 = v37
	v41 = v37
	v43 = v5
	v44 = int32(1)
	v46 = v5
	goto L14
L14:
	;
	v53 = v40
	v54 = v41
	v56 = v43
	v57 = v44
	v59 = v46
	v61 = int32(0)
	goto L16
L16:
	;
	if v53&int32(1) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v40 = int32(0)
	v41 = v313
	v43 = v310
	v44 = v314
	v46 = l3
	goto L14
L18:
	;
	goto L17
L19:
	;
	v53 = int32(0)
	v54 = v444
	v56 = v445
	v57 = v446
	v59 = l3
	goto L16
L20:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	m.G0 = v15 + int32(32)
	return v438
L21:
	;
	v73 = F_JsonbIteratorNext(m, v15+int32(28), v15+int32(8), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	v77 = v54
	goto L23
L23:
	;
	switch v77 - int32(1) {
	case 0:
		goto L32
	default:
		goto L26
	case 2:
		goto L31
	case 3:
		goto L34
	case 4:
		goto L30
	case 5:
		goto L33
	case 6:
		goto L29
	}
L24:
	;
	if v73 == int32(0) {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v77 = v73
	goto L23
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L4
	} else {
		goto L131
	}
L27:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v413 = int32(125)
	*(*uint8)(unsafe.Add(mBase, uint32(v411+v396))) = uint8(v413)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v417 = v415 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v417
	v419 = int32(0)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*uint8)(unsafe.Add(mBase, uint32(v420+v417))) = uint8(v419)
	v444 = int32(7)
	v445 = v366
	v446 = v419
	goto L19
L28:
	;
	v444 = v409
	v445 = v406
	v446 = int32(0)
	goto L19
L29:
	;
	v365 = int32(1)
	v366 = v56 - v365
	if v59&v365 != 0 {
		goto L120
	} else {
		goto L121
	}
L30:
	;
	v309 = int32(1)
	v310 = v56 - v309
	v313 = int32(5)
	v314 = int32(0)
	if v61&v309 != 0 {
		v53 = v314
		v54 = v313
		v56 = v310
		v57 = v314
		v59 = l3
		v61 = v309
		goto L16
	} else {
		goto L106
	}
L31:
	;
	if v57 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L32:
	;
	if v57 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L33:
	;
	if v57 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L34:
	;
	if v57 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_appendBinaryStringInfo(m, v23, int32(748219), v26)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v85 = int32(1)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v87 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L37
L39:
	;
	if v59&(v53^int32(-1))&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v143 = v85
	goto L41
L41:
	;
	v53 = int32(0)
	v54 = int32(4)
	v56 = v56 + int32(1)
	v57 = v85
	v59 = l3
	v61 = v143
	goto L16
L42:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v95 <= v96+int32(1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v121 <= v122+int32(1) {
		goto L52
	} else {
		goto L53
	}
L45:
	;
	F_appendStringInfoSpaces(m, v23, v56<<(uint(int32(2))%32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L50
	}
L46:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v105 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v103+v96))) = uint8(v105)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v109 = v107 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111+v109))) = uint8(v113)
	goto L45
L49:
	;
	goto L45
L50:
	;
	goto L44
L51:
	;
	v143 = v61
	goto L41
L52:
	;
	F_appendStringInfoChar(m, v23, int32(91))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v131 = int32(91)
	*(*uint8)(unsafe.Add(mBase, uint32(v129+v122))) = uint8(v131)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v135 = v133 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v137+v135))) = uint8(v139)
	goto L51
L55:
	;
	goto L51
L56:
	;
	F_appendBinaryStringInfo(m, v23, int32(748219), v26)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v59&(v53^int32(-1))&int32(1) != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v158 <= v159+int32(1) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v184 <= v185+int32(1) {
		goto L70
	} else {
		goto L71
	}
L63:
	;
	F_appendStringInfoSpaces(m, v23, v56<<(uint(int32(2))%32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L68
	}
L64:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v168 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v166+v159))) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v172 = v170 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174+v172))) = uint8(v176)
	goto L63
L67:
	;
	goto L63
L68:
	;
	goto L62
L69:
	;
	v205 = int32(1)
	v444 = int32(6)
	v445 = v56 + v205
	v446 = v205
	goto L19
L70:
	;
	F_appendStringInfoChar(m, v23, int32(123))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v194 = int32(123)
	*(*uint8)(unsafe.Add(mBase, uint32(v192+v185))) = uint8(v194)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v198 = v196 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v202 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200+v198))) = uint8(v202)
	goto L69
L73:
	;
	goto L69
L74:
	;
	F_appendBinaryStringInfo(m, v23, int32(748219), v26)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if v59&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v216 <= v217+int32(1) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	goto L80
L80:
	;
	F_jsonb_put_escaped_value(m, v23, v15+int32(8))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L87
	}
L81:
	;
	F_appendStringInfoSpaces(m, v23, v56<<(uint(int32(2))%32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L4
	} else {
		goto L86
	}
L82:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v226 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v224+v217))) = uint8(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v230 = v228 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v230
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v232+v230))) = uint8(v234)
	goto L81
L85:
	;
	goto L81
L86:
	;
	goto L80
L87:
	;
	F_appendBinaryStringInfo(m, v23, int32(747976), int32(2))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	v250 = int32(1)
	v257 = F_JsonbIteratorNext(m, v15+int32(28), v15+int32(8), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v257 != int32(2) {
		v53 = v250
		v54 = v257
		v57 = v250
		v59 = l3
		goto L16
	} else {
		goto L90
	}
L90:
	;
	F_jsonb_put_escaped_value(m, v23, v15+int32(8))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v406 = v56
	v409 = int32(2)
	goto L28
L92:
	;
	F_appendBinaryStringInfo(m, v23, int32(748219), v26)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if (v61|(v59^int32(-1)))&int32(1) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v278 <= v279+int32(1) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	goto L98
L98:
	;
	F_jsonb_put_escaped_value(m, v23, v15+int32(8))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L105
	}
L99:
	;
	F_appendStringInfoSpaces(m, v23, v56<<(uint(int32(2))%32))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L104
	}
L100:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v288 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v286+v279))) = uint8(v288)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v292 = v290 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v296 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v294+v292))) = uint8(v296)
	goto L99
L103:
	;
	goto L99
L104:
	;
	goto L98
L105:
	;
	v406 = v56
	v409 = int32(3)
	goto L28
L106:
	;
	if v59&int32(1) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v319 <= v320+int32(1) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v345 <= v346+int32(1) {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	F_appendStringInfoSpaces(m, v23, v310<<(uint(int32(2))%32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L115
	}
L111:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v329 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v327+v320))) = uint8(v329)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v333 = v331 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v337 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v335+v333))) = uint8(v337)
	goto L110
L114:
	;
	goto L110
L115:
	;
	goto L109
L116:
	;
	F_appendStringInfoChar(m, v23, int32(93))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v355 = int32(93)
	*(*uint8)(unsafe.Add(mBase, uint32(v353+v346))) = uint8(v355)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v359 = v357 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v359
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v363 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v361+v359))) = uint8(v363)
	goto L18
L119:
	;
	goto L18
L120:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v369 <= v370+int32(1) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v396+int32(1) < v395 {
		goto L27
	} else {
		goto L129
	}
L123:
	;
	F_appendStringInfoSpaces(m, v23, v366<<(uint(int32(2))%32))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L128
	}
L124:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v379 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v377+v370))) = uint8(v379)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v383 = v381 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v385+v383))) = uint8(v387)
	goto L123
L127:
	;
	goto L123
L128:
	;
	goto L122
L129:
	;
	F_appendStringInfoChar(m, v23, int32(125))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v406 = v366
	v409 = int32(7)
	goto L28
L131:
	;
	F_errmsg_internal(m, int32(370023), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(501581), int32(603), int32(221726))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_JsonbTypeName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 {
	case 0:
		v101 = int32(304327)
		m.G0 = v7 - int32(-64)
		return v101
	case 1:
		v101 = int32(331637)
		m.G0 = v7 - int32(-64)
		return v101
	case 2:
		v101 = int32(229051)
		m.G0 = v7 - int32(-64)
		return v101
	case 3:
		v101 = int32(284909)
		m.G0 = v7 - int32(-64)
		return v101
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int32(0)
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v69
			F_errmsg_internal(m, int32(486927), v7)
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(501581), int32(217), int32(383366))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v101 = int32(26230)
		m.G0 = v7 - int32(-64)
		return v101
	case 17:
		v101 = int32(111875)
		m.G0 = v7 - int32(-64)
		return v101
	case 18:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = F_JsonbExtractScalar(m, v11, v5+int32(-20))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				v20 = F_JsonbTypeName(m, v5+int32(-20))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v101 = v20
					m.G0 = v7 - int32(-64)
					return v101
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				if v22&int32(1073741824) != 0 {
					v101 = int32(26230)
					m.G0 = v7 - int32(-64)
					return v101
				} else {
					if v22&int32(536870912) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v86
							F_errmsg_internal(m, int32(29817), v5+int32(-48))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(501581), int32(171), int32(383299))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v101 = int32(111875)
						m.G0 = v7 - int32(-64)
						return v101
					}
				}
			}
		}
	case 32:
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v34 <= int32(1183) {
			switch v34 - int32(1082) {
			case 0:
				v101 = int32(358309)
				m.G0 = v7 - int32(-64)
				return v101
			case 1:
				v101 = int32(373652)
				m.G0 = v7 - int32(-64)
				return v101
			default:
				if v34 == int32(1114) {
					v101 = int32(373624)
					m.G0 = v7 - int32(-64)
					return v101
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v53
						F_errmsg_internal(m, int32(486991), v5+int32(-32))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(501581), int32(213), int32(383366))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
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
		} else {
			if v34 != int32(1184) {
				if v34 != int32(1266) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v53
						F_errmsg_internal(m, int32(486991), v5+int32(-32))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(501581), int32(213), int32(383366))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v101 = int32(373700)
					m.G0 = v7 - int32(-64)
					return v101
				}
			} else {
				v101 = int32(373675)
				m.G0 = v7 - int32(-64)
				return v101
			}
		}
	}
}
func F_compareJsonbContainers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
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
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = F_iteratorFromContainer(m, l0, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+44)) = v10
	v16 = F_iteratorFromContainer(m, l1, int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v16
	goto L4
L4:
	;
	v28 = F_JsonbIteratorNext(m, v7+int32(44), v7+int32(20), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v135 != 0 {
		goto L62
	} else {
		goto L63
	}
L6:
	;
	goto L5
L7:
	;
	if v129 == int32(0) {
		goto L4
	} else {
		goto L61
	}
L8:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v125 = F_varstr_cmp(m, v120, v121, v122, v123, int32(100))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L60
	}
L9:
	;
	v33 = F_JsonbIteratorNext(m, v7+int32(40), v7, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v28 == v33 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v28 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if base.Ui32(v117) < base.Ui32(v116) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v134 = int32(0)
	goto L6
L15:
	;
	goto L16
L16:
	;
	v39 = int32(0)
	v40 = int32(5)
	if v28&v40 == v40 {
		v129 = v39
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v44 == v45 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	switch v44 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L26
	case 2:
		goto L25
	default:
		v129 = v39
		goto L7
	case 15:
		goto L24
	case 16:
		goto L23
	case 17:
		goto L22
	case 31:
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(v45) < base.Ui32(v44) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L51
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L48
	}
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v77 == v78 {
		v129 = v39
		goto L7
	} else {
		goto L44
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v62 == v63 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
	if v55 == v56 {
		v129 = v39
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v53 = F_DirectFunctionCall2Coll(m, int32(1343), int32(0), v51, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v129 = v53
	goto L7
L28:
	;
	if base.Ui32(v56) < base.Ui32(v55) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v61 = int32(1)
	goto L31
L30:
	;
	v61 = int32(-1)
	goto L31
L31:
	;
	v134 = v61
	goto L6
L32:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+32)))
	if v67 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v63 < v62 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v68 = int32(-1)
	goto L37
L36:
	;
	v68 = int32(1)
	goto L37
L37:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v67 != v70 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v72 = v68
	goto L40
L39:
	;
	v72 = int32(0)
	goto L40
L40:
	;
	v129 = v72
	goto L7
L41:
	;
	v76 = int32(1)
	goto L43
L42:
	;
	v76 = int32(-1)
	goto L43
L43:
	;
	v134 = v76
	goto L6
L44:
	;
	if v78 < v77 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v83 = int32(1)
	goto L47
L46:
	;
	v83 = int32(-1)
	goto L47
L47:
	;
	v134 = v83
	goto L6
L48:
	;
	F_errmsg_internal(m, int32(346499), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(498906), int32(264), int32(134186))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errmsg_internal(m, int32(347291), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(498906), int32(267), int32(134186))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v113 = int32(1)
	goto L56
L55:
	;
	v113 = int32(-1)
	goto L56
L56:
	;
	v134 = v113
	goto L6
L57:
	;
	v119 = int32(1)
	goto L59
L58:
	;
	v119 = int32(-1)
	goto L59
L59:
	;
	v134 = v119
	goto L6
L60:
	;
	v129 = v125
	goto L7
L61:
	;
	v134 = v129
	goto L6
L62:
	;
	v136 = v135
	goto L65
L63:
	;
	goto L64
L64:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	if v147 != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v136)+36))
	F_pfree(m, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	if v140 != 0 {
		v136 = v140
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v148 = v147
	goto L72
L70:
	;
	goto L71
L71:
	;
	m.G0 = v7 + int32(48)
	return v134
L72:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	F_pfree(m, v148)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L71
L74:
	;
	if v152 != 0 {
		v148 = v152
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
}
func F_findJsonbValueFromContainer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
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
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = v9 & int32(268435455)
	if v11 == v4 {
		v196 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v196
L2:
	;
	v14 = l1 & v9
	if v14&int32(1073741824) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = l0 + int32(4)
	v21 = v18 + v11<<(uint(int32(2))%32)
	v23 = F_palloc(m, int32(20))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	if v14&int32(536870912) == int32(0) {
		v196 = v4
		goto L1
	} else {
		goto L46
	}
L6:
	;
	return int32(0)
L7:
	;
	v27 = int32(0)
	v30 = v27
	v33 = v27
	goto L8
L8:
	;
	v40 = l0 + int32(4)
	v43 = v40 + v30<<(uint(int32(2))%32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	switch int32(base.Ui32(v44)>>(uint(int32(28))%32)) & int32(7) {
	case 0:
		goto L15
	case 1:
		goto L14
	case 2:
		goto L12
	case 3:
		goto L13
	case 4:
		goto L16
	default:
		goto L11
	}
L9:
	;
	F_pfree(m, v23)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L45
	}
L10:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v161 == v162 {
		goto L36
	} else {
		goto L37
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(18)
	v109 = (v33 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v21 + v109
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v112 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v100 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)) = uint8(v100)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(3)
	goto L10
L13:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+4)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(3)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v21 + (v33+int32(3))&int32(-4)
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v21 + v33
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v55 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(0)
	goto L10
L17:
	;
	v60 = v30
	v61 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v55 & int32(268435455)
	goto L10
L20:
	;
	v67 = v60 - int32(1)
	if int32(0) <= v67 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v55&int32(268435455) - v79
	goto L10
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v40+v67<<(uint(int32(2))%32))))
	v76 = v73&int32(268435455) + v61
	if int32(0) <= v73 {
		v60 = v67
		v61 = v76
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v79 = v61
	goto L24
L24:
	;
	goto L21
L25:
	;
	v79 = v76
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v150 + (v33 - v109)
	goto L10
L27:
	;
	v117 = v30
	v118 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v150 = v112 & int32(268435455)
	goto L26
L30:
	;
	v124 = v117 - int32(1)
	if int32(0) <= v124 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v150 = v112&int32(268435455) - v136
	goto L26
L32:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v40+v124<<(uint(int32(2))%32))))
	v133 = v130&int32(268435455) + v118
	if int32(0) <= v130 {
		v117 = v124
		v118 = v133
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v136 = v118
	goto L34
L34:
	;
	goto L31
L35:
	;
	v136 = v133
	goto L34
L36:
	;
	v164 = F_equalsJsonbScalarValue(m, l2, v23)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v166 = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v18+v30<<(uint(int32(2))%32))))
	if v166 <= v170 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v164 != 0 {
		v196 = v23
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v173 = v33
	goto L43
L42:
	;
	v173 = v166
	goto L43
L43:
	;
	v178 = v30 + int32(1)
	if v178 != v11 {
		v30 = v178
		v33 = v173 + v170&int32(268435455)
		goto L8
	} else {
		goto L44
	}
L44:
	;
	goto L9
L45:
	;
	return int32(0)
L46:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v191 = F_getKeyJsonValueFromContainer(m, l0, v188, v189, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v196 = v191
	goto L1
}
func F_jsonb_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v83
L2:
	;
	v12 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
	v83 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = F_palloc(m, int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v69 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v69
	v74 = F_pushJsonbValue(m, v7, int32(5), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)) = uint8(v35)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+29)) = uint8(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v23
	v41 = v39
	goto L13
L11:
	;
	v61 = v23
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(0)
	v69 = v23
	goto L7
L13:
	;
	v45 = F_palloc(m, int32(32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v61 = v45
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)) = uint8(v56)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+29)) = uint8(v58)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v60 != 0 {
		v40 = v45
		v41 = v60
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v74
	v77 = F_JsonbValueToJsonb(m, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v83 = v77
	goto L1
}
func F_jsonb_agg_strict_transfn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_agg_transfn_worker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_bool(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v19 = F_JsonbExtractScalar(m, v11+int32(4), v8+int32(12))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v19 != 0 {
				switch v21 {
				case 0:
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v22 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
							v37 = int32(0)
							m.G0 = v8 + int32(32)
							return v37
						}
					} else {
						v26 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
						v37 = int32(0)
						m.G0 = v8 + int32(32)
						return v37
					}
				default:
					F_cannotCastJsonbValue(m, v21, int32(284909))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				case 3:
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v32 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
							v37 = v36
							m.G0 = v8 + int32(32)
							return v37
						}
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
						v37 = v36
						m.G0 = v8 + int32(32)
						return v37
					}
				}
			} else {
				F_cannotCastJsonbValue(m, v21, int32(284909))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
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
func F_jsonb_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return v20
							}
						} else {
							return v20
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return v20
						}
					} else {
						return v20
					}
				}
			}
		}
	}
}
func F_jsonb_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v20 == int32(0))
							}
						} else {
							return base.B2i32(v20 == int32(0))
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v20 == int32(0))
						}
					} else {
						return base.B2i32(v20 == int32(0))
					}
				}
			}
		}
	}
}
func F_jsonb_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v14 = l0 + int32(28)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v20 = F_compareJsonbContainers(m, v7+int32(4), v16+int32(4))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v22 != v7 {
					F_pfree(m, v7)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						if v26 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v29 = m.ExcPending
							if v29 != 0 {
								return int32(0)
							} else {
								return int32(base.Ui32(v20^int32(-1)) >> (uint(int32(31)) % 32))
							}
						} else {
							return int32(base.Ui32(v20^int32(-1)) >> (uint(int32(31)) % 32))
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if v26 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v20^int32(-1)) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v20^int32(-1)) >> (uint(int32(31)) % 32))
					}
				}
			}
		}
	}
}
func F_jsonb_hash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v16&int32(268435455) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = F_JsonbIteratorInit(m, v10+int32(4))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v69 = int32(0)
	goto L5
L5:
	;
	m.G0 = v7 + int32(32)
	return v69
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v21
	goto L8
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v62 != v10 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v33 = F_JsonbIteratorNext(m, v7+int32(28), v7+int32(8), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	F_JsonbHashScalarValue(m, v7+int32(8), v7+int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v39 ^ int32(536870912)
	goto L8
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v35 ^ int32(1073741824)
	goto L8
L14:
	;
	switch v33 {
	case 0:
		goto L7
	case 1, 2, 3:
		goto L11
	case 4:
		goto L13
	case 5, 7:
		goto L8
	case 6:
		goto L12
	default:
		goto L10
	}
L15:
	;
	goto L8
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v33
	F_errmsg_internal(m, int32(489070), v7)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(497304), int32(286), int32(324095))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	F_pfree(m, v10)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v69 = v66
	goto L5
L22:
	;
	goto L21
}
func F_jsonb_in_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = F_pushJsonbValue(m, l0, int32(6), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v10)
		return int32(0)
	}
}
func F_jsonb_object_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v83
L2:
	;
	v12 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v12)
	v83 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+4)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = F_palloc(m, int32(32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v69 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v69
	v74 = F_pushJsonbValue(m, v7, int32(7), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v29
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+28)) = uint8(v35)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+29)) = uint8(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v23
	v41 = v39
	goto L13
L11:
	;
	v61 = v23
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(0)
	v69 = v23
	goto L7
L13:
	;
	v45 = F_palloc(m, int32(32))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v61 = v45
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+16)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v41)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	*(*int64)(unsafe.Add(mBase, uint32(v45))) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v45)+20)) = v54
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+28)) = uint8(v56)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v45)+29)) = uint8(v58)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	if v60 != 0 {
		v40 = v45
		v41 = v60
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v74
	v77 = F_JsonbValueToJsonb(m, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v83 = v77
	goto L1
}
func F_jsonb_object_agg_transfn(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_jsonb_object_field(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
			if v19&int32(32) == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v81 = int32(0)
				m.G0 = v9 + int32(32)
				return v81
			} else {
				v26 = int32(4)
				v28 = int32(1)
				v29 = v17 + v28
				v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v34 = v32 & v28
				if v34 != 0 {
					v35 = v29
				} else {
					v35 = v17 + v26
				}
				if v32 == int32(1) {
					v38 = int32(4)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
					if v40&int32(254) == int32(2) {
						v49 = v38
					} else {
						v49 = base.B2i32(v40 == int32(18)) << (uint(v38) % 32)
					}
					if v40 == int32(1) {
						v52 = v38
					} else {
						v52 = v49
					}
					v63 = v52
				} else {
					v53 = int32(1)
					if v34 != 0 {
						v63 = int32(base.Ui32(v32)>>(uint(v53)%32)) - v53
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v63 = int32(base.Ui32(v57)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v66 = F_getKeyJsonValueFromContainer(m, v12+v26, v35, v63, v9+int32(12))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					if v66 != 0 {
						v68 = F_JsonbValueToJsonb(m, v66)
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							v81 = v68
							m.G0 = v9 + int32(32)
							return v81
						}
					} else {
						v70 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
						v81 = int32(0)
						m.G0 = v9 + int32(32)
						return v81
					}
				}
			}
		}
	}
}
func F_jsonb_path_query(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_pretty(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_makeStringInfo(m)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v16 = F_JsonbToCStringWorker(m, v8, v4+int32(4), int32(base.Ui32(v12)>>(uint(int32(2))%32)), int32(1))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v20 = F_cstring_to_text_with_len(m, v18, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		}
	}
}
func F_jsonb_set(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(0)
				v25 = v8 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(18)
				v28 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v19 + v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(base.Ui32(v31)>>(uint(int32(2))%32)) - v28
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
				if v37 < int32(2) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
					if v40&int32(268435456) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(230218), int32(0))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(496116), int32(4892), int32(106678))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
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
						if v40&int32(268435455)|v21 == int32(0) {
							v82 = v11
							m.G0 = v8 + int32(48)
							return v82
						} else {
							F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(24), v8+int32(20), v8+int32(16))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
								if v57 == int32(0) {
									v82 = v11
									m.G0 = v8 + int32(48)
									return v82
								} else {
									v62 = F_JsonbIteratorInit(m, v11+int32(4))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v62
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
										if v21 != 0 {
											v77 = int32(1)
										} else {
											v77 = int32(4)
										}
										v78 = F_setPath(m, v8+int32(12), v67, v68, v69, v8+int32(8), int32(0), v8+int32(28), v77)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											v80 = F_JsonbValueToJsonb(m, v78)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = v80
												m.G0 = v8 + int32(48)
												return v82
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
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(352845954))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(118749), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(496116), int32(4887), int32(106678))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
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
	}
}
func F_jsonb_subscript_fetch_old(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = int32(0)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
			v22 = F_jsonb_get_element(m, v15, v17, v18, v4+int32(52), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v22
				return
			}
		}
	}
}
func F_jsonb_subscript_handler(m *base.Module, l0 int32) int32 {
	return int32(1670468)
}
func F_jsonb_to_tsvector(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = l0 + int32(28)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = F_parse_jsonb_index_flags(m, v19)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_getTSCurrentConfig(m)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v23
					v26 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v26
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v9 + int32(8)
					F_iterate_jsonb_values(m, v12, v21, v9+int32(24))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v39 = F_make_tsvector(m, v9+int32(8))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v41 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									if v45 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return int32(0)
										} else {
											m.G0 = v9 + int32(32)
											return v39
										}
									} else {
										m.G0 = v9 + int32(32)
										return v39
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								if v45 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										m.G0 = v9 + int32(32)
										return v39
									}
								} else {
									m.G0 = v9 + int32(32)
									return v39
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	if l1&int32(-2) != int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(32)
	return v260
L2:
	;
	if l2 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L3:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v17 - int32(16) {
	case 0:
		goto L5
	case 1:
		goto L6
	default:
		goto L2
	}
L5:
	;
	v85 = F_palloc(m, int32(32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L7
	} else {
		goto L19
	}
L6:
	;
	v21 = F_palloc(m, int32(32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+28)) = uint16(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(17)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v26
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(4)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v42 = F_palloc(m, v39*int32(44))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+8)) = v42
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v46 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v82 = F_pushJsonbValueScalar(m, l0, int32(7), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L18
	}
L13:
	;
	v58 = v53 * int32(44)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v61 = F_pushJsonbValueScalar(m, l0, int32(1), v58+v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v68 = F_pushJsonbValue(m, l0, int32(2), v64+v58+int32(20))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v71 = v53 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v71 < v72 {
		v53 = v71
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v260 = v82
	goto L1
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+28)) = uint16(v88)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(16)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v88
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+12)) = uint8(v88)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v100)+20)) = int32(4)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+20))
	v107 = F_palloc(m, v104*int32(20))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v107
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v111 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v117 = v4
	goto L24
L22:
	;
	goto L23
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v138
	if v138 == int32(0) {
		v260 = v137
		goto L1
	} else {
		goto L28
	}
L24:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v125 = F_pushJsonbValue(m, l0, int32(3), v121+v117*int32(20))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L7
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v128 = v117 + int32(1)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v128 < v129 {
		v117 = v128
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	switch v142 - int32(16) {
	case 0:
		goto L31
	case 1:
		goto L30
	default:
		goto L29
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L33
	}
L30:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v147 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v154 = v151 + v147*int32(44)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+36)) = v155
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v137)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+28)) = v157
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+20)) = v159
	v260 = v137
	goto L1
L31:
	;
	F_appendElement(m, v138, v137)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v260 = v137
	goto L1
L33:
	;
	F_errmsg_internal(m, int32(369421), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(498906), int32(720), int32(230446))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v187 = F_iteratorFromContainer(m, v185, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L7
	} else {
		goto L42
	}
L37:
	;
	v183 = F_pushJsonbValueScalar(m, l0, l1, l2)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L41
	}
L38:
	;
	if base.Ui32(l1-int32(4)) < base.Ui32(int32(-2)) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v180 == int32(18) {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v260 = v183
	goto L1
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v187
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+3)))
	if v191&int32(16) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v229 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L50
	}
L44:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v196 == int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v204 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v211 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v215 = F_pushJsonbValueScalar(m, l0, l1, v9+int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v222 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(1))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v260 = v215
	goto L1
L50:
	;
	if v229 == int32(0) {
		v260 = v4
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v235 = v229
	goto L52
L52:
	;
	v240 = v9 + int32(8)
	if base.Ui32(v235) < base.Ui32(int32(4)) {
		v247 = v240
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v260 = v248
	goto L1
L54:
	;
	v248 = F_pushJsonbValueScalar(m, l0, v235, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L60
	}
L55:
	;
	if v235 == int32(4) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
	if v245 != 0 {
		v247 = v240
		goto L54
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v247 = int32(0)
	goto L54
L59:
	;
	goto L58
L60:
	;
	v255 = F_JsonbIteratorNext(m, v9+int32(28), v9+int32(8), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	if v255 != 0 {
		v235 = v255
		goto L52
	} else {
		goto L62
	}
L62:
	;
	goto L53
}
func F_pushJsonbValueScalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v182 int32
	_ = v182
	var __phi182 int32
	_ = __phi182
	var v183 int32
	_ = v183
	var __phi183 int32
	_ = __phi183
	var v184 int32
	_ = v184
	var __phi184 int32
	_ = __phi184
	var v186 int32
	_ = v186
	var __phi186 int32
	_ = __phi186
	var v187 int32
	_ = v187
	var __phi187 int32
	_ = __phi187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
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
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
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
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	switch l1 - int32(1) {
	case 0:
		goto L10
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L11
	case 4:
		goto L6
	case 5:
		goto L4
	case 6:
		goto L7
	default:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L12
	} else {
		goto L98
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L12
	} else {
		goto L94
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v395
L4:
	;
	v371 = F_palloc(m, int32(32))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L12
	} else {
		goto L92
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L12
	} else {
		goto L89
	}
L6:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v320)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v321
	if v321 == int32(0) {
		v395 = v320
		goto L3
	} else {
		goto L81
	}
L7:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+29)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+28)))
	v112 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v112)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if int32(2) <= v115 {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_appendElement(m, v105, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L27
	}
L9:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v90 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v97 = v94 + v90*int32(44)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+36)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+28)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+20)) = v102
	v395 = int32(0)
	goto L3
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if base.Ui32(int32(24403223)) <= base.Ui32(v55) {
		goto L2
	} else {
		goto L21
	}
L11:
	;
	v18 = F_palloc(m, int32(32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return int32(0)
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)) = uint16(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(16)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v23
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+20)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	v50 = F_palloc(m, v47*int32(20))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L20
	}
L15:
	;
	v43 = int32(4)
	goto L14
L16:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+12)) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v35 <= int32(0) {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+12)) = uint8(v39)
	goto L15
L19:
	;
	v43 = v35
	goto L14
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v50
	v395 = v18
	goto L3
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
	if base.Ui32(v55) < base.Ui32(v58) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v73 = int32(44)
	v75 = v72 + v71*v73
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v75)+8)) = v80
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82+v83*v73)+40)) = v83
	v395 = int32(0)
	goto L3
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v71 = v55
	v72 = v60
	goto L22
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v58 << (uint(int32(1)) % 32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v67 = F_repalloc(m, v64, v58*int32(88))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v71 = v70
	v72 = v67
	goto L22
L27:
	;
	v395 = int32(0)
	goto L3
L28:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	F_qsort_arg(m, v118, v115, int32(44), int32(1342), v13+int32(15))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L12
	} else {
		goto L31
	}
L29:
	;
	v129 = v112
	goto L30
L30:
	;
	v130 = int32(1)
	if v129&v130|v110&v130 == int32(0) {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v125&v111&int32(1) != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v129 = v125
	goto L30
L33:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v110&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v168 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L35:
	;
	if v137 <= int32(0) {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v137 <= int32(0) {
		goto L6
	} else {
		goto L43
	}
L38:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v146 = v137
	v149 = v142
	goto L39
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	if v153 != 0 {
		v168 = v146
		v171 = v149
		goto L34
	} else {
		goto L41
	}
L40:
	;
	goto L6
L41:
	;
	v154 = int32(1)
	v155 = v146 - v154
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v155
	v158 = v149 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v158
	if v154 < v146 {
		v146 = v155
		v149 = v158
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v168 = v137
	v171 = v164
	goto L34
L44:
	;
	v305 = int32(44)
	v308 = base.I32_div_s(v299-v300+v305, v305)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v308
	goto L6
L45:
	;
	v299 = v171
	v300 = v171
	goto L44
L46:
	;
	goto L47
L47:
	;
	__phi182 = v171
	__phi183 = v171 + int32(44)
	__phi184 = v168
	__phi186 = v171
	__phi187 = v171
	v182 = __phi182
	v183 = __phi183
	v184 = __phi184
	v186 = __phi186
	v187 = __phi187
	goto L48
L48:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v182)+48))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v191 == v192 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v299 = v286
	v300 = v287
	goto L44
L50:
	;
	v288 = int32(44)
	v289 = v183 + v288
	v292 = base.I32_div_s(v289-v287, v288)
	if v292 < v285 {
		__phi182 = v183
		__phi183 = v289
		__phi184 = v285
		__phi186 = v286
		__phi187 = v287
		v182 = __phi182
		v183 = __phi183
		v184 = __phi184
		v186 = __phi186
		v187 = __phi187
		goto L48
	} else {
		goto L80
	}
L51:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v182)+52))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if base.Ui32(int32(4)) <= base.Ui32(v191) {
		goto L57
	} else {
		goto L58
	}
L52:
	;
	goto L53
L53:
	;
	if v110&int32(1) != 0 {
		goto L73
	} else {
		goto L74
	}
L54:
	;
	if v257 == int32(0) {
		v285 = v184
		v286 = v186
		v287 = v187
		goto L50
	} else {
		goto L72
	}
L55:
	;
	v257 = int32(0)
	goto L54
L56:
	;
	v231 = v226
	v232 = v227
	v233 = v228
	goto L66
L57:
	;
	if (v194|v195)&int32(3) != 0 {
		v226 = v194
		v227 = v195
		v228 = v191
		goto L56
	} else {
		goto L60
	}
L58:
	;
	v219 = v194
	v220 = v195
	v221 = v191
	goto L59
L59:
	;
	if v221 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v203 = v194
	v204 = v195
	v205 = v191
	goto L61
L61:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	if v208 != v209 {
		v226 = v203
		v227 = v204
		v228 = v205
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v219 = v214
	v220 = v212
	v221 = v216
	goto L59
L63:
	;
	v211 = int32(4)
	v212 = v204 + v211
	v214 = v203 + v211
	v216 = v205 - v211
	if base.Ui32(int32(3)) < base.Ui32(v216) {
		v203 = v214
		v204 = v212
		v205 = v216
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v226 = v219
	v227 = v220
	v228 = v221
	goto L56
L66:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232))))
	if v236 == v237 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v257 = v236 - v237
	goto L54
L68:
	;
	v239 = int32(1)
	v244 = v233 - v239
	if v244 != 0 {
		v231 = v231 + v239
		v232 = v232 + v239
		v233 = v244
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L55
L72:
	;
	goto L53
L73:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v182-int32(-64))))
	if v262 == int32(0) {
		v285 = v184
		v286 = v186
		v287 = v187
		goto L50
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v266 = v186 + int32(44)
	if v186 != v182 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v183)))
	*(*int64)(unsafe.Add(mBase, uint32(v266))) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v183)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v266)+40)) = v270
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v183)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v266)+32)) = v272
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v183)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v266)+24)) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v183)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v266)+16)) = v276
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v183)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v266)+8)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v282 = v280
	v283 = v281
	goto L79
L78:
	;
	v282 = v184
	v283 = v187
	goto L79
L79:
	;
	v285 = v282
	v286 = v266
	v287 = v283
	goto L50
L80:
	;
	goto L49
L81:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
	switch v325 - int32(16) {
	case 0:
		goto L84
	case 1:
		goto L83
	default:
		goto L82
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L12
	} else {
		goto L86
	}
L83:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v321)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v321)+4)) = v330 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	v337 = v334 + v330*int32(44)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v320)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v337)+36)) = v338
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v320)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v337)+28)) = v340
	v342 = *(*int64)(unsafe.Add(mBase, uint32(v320)))
	*(*int64)(unsafe.Add(mBase, uint32(v337)+20)) = v342
	v395 = v320
	goto L3
L84:
	;
	F_appendElement(m, v321, v320)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	v395 = v320
	goto L3
L86:
	;
	F_errmsg_internal(m, int32(369421), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L12
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(498906), int32(720), int32(230446))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(283100), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L12
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(498906), int32(725), int32(230446))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L12
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v374 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v371)+28)) = uint16(v374)
	*(*int32)(unsafe.Add(mBase, uint32(v371)+24)) = v373
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v371
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = int32(17)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v374
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v383)+20)) = int32(4)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+20))
	v390 = F_palloc(m, v387*int32(44))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L12
	} else {
		goto L93
	}
L93:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v392)+8)) = v390
	v395 = v371
	goto L3
L94:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L12
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(24403223)
	F_errmsg(m, int32(672813), v13)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L12
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(498906), int32(761), int32(22727))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L12
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(786562))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L12
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(346526), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L12
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(498906), int32(1969), int32(111952))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L12
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
