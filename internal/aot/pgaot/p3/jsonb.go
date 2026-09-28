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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
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
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(112)
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
	v21 = F_JsonbIteratorNext(m, l0, v12+int32(80), int32(0))
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
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L81
	}
L4:
	;
	m.G0 = v12 + int32(112)
	return v262
L5:
	;
	v26 = F_JsonbIteratorNext(m, l1, v12+int32(48), int32(0))
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
		v262 = v3
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
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+96)))
	if v105 == int32(1) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v31 < v32 {
		v262 = v3
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v37 = F_JsonbIteratorNext(m, l1, v12+int32(48), int32(0))
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
	v262 = int32(1)
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
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v57 = F_getKeyJsonValueFromContainer(m, v52, v53, v54, v12+int32(16))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v262 = int32(1)
	goto L4
L17:
	;
	if v57 == int32(0) {
		v262 = v3
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v62 = v12 + int32(48)
	v64 = F_JsonbIteratorNext(m, l1, v62, int32(1))
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if v66 != v67 {
		v262 = v3
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
	v99 = F_JsonbIteratorNext(m, l1, v12+int32(48), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v76 = F_equalsJsonbScalarValue(m, v57, v62)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v80 = F_iteratorFromContainer(m, v78, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	if v76 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v262 = v3
	goto L4
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v80
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v85 = F_iteratorFromContainer(m, v83, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v85
	v92 = F_JsonbDeepContains(m, v12+int32(12), v12+int32(8))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v92 == int32(0) {
		v262 = v3
		goto L4
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	if v99 != int32(7) {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	goto L16
L33:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+64)))
	if v108&int32(1) == int32(0) {
		v262 = v3
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
	v117 = F_JsonbIteratorNext(m, l1, v12+int32(48), v113)
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
		v262 = v113
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v125 = v104
	v128 = v3
	goto L39
L39:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v130))&base.B2i32(v130 != int32(32)) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v262 = v251
	goto L4
L41:
	;
	v251 = int32(1)
	v255 = F_JsonbIteratorNext(m, l1, v12+int32(48), v251)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L79
	}
L42:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v143 = F_findJsonbValueFromContainer(m, v139, int32(1073741824), v12+int32(48))
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
		goto L48
	} else {
		goto L49
	}
L45:
	;
	if v143 != 0 {
		v246 = v125
		v249 = v128
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v262 = int32(0)
	goto L4
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v208 = int32(0)
	goto L63
L48:
	;
	v148 = int32(0)
	v150 = F_palloc_mul(m, int32(32), v125)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v125 != 0 {
		v195 = v125
		v199 = v128
		goto L47
	} else {
		goto L61
	}
L51:
	;
	if v125 == int32(0) {
		v262 = v148
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v158 = v148
	v160 = int32(0)
	goto L53
L53:
	;
	v167 = F_JsonbIteratorNext(m, l0, v12+int32(80), int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	if v185 != 0 {
		v195 = v185
		v199 = v150
		goto L47
	} else {
		goto L60
	}
L55:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v12)+80))
	if v169 == int32(18) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v174 = v150 + v158<<(uint(int32(5))%32)
	v175 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v174)+24)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v174)+16)) = v177
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
	*(*int64)(unsafe.Add(mBase, uint32(v174)+8)) = v179
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v12)+80))
	*(*int64)(unsafe.Add(mBase, uint32(v174))) = v181
	v185 = v158 + int32(1)
	goto L58
L57:
	;
	v185 = v158
	goto L58
L58:
	;
	v188 = v160 + int32(1)
	if v188 != v125 {
		v158 = v185
		v160 = v188
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v262 = int32(0)
	goto L4
L61:
	;
	v262 = int32(0)
	goto L4
L62:
	;
	if v208 != v195 {
		v246 = v195
		v249 = v199
		goto L41
	} else {
		goto L78
	}
L63:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v199+v208<<(uint(int32(5))%32))+12))
	v217 = F_iteratorFromContainer(m, v215, int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v262 = int32(0)
	goto L4
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v217
	v221 = F_iteratorFromContainer(m, v202, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v221
	v228 = F_JsonbDeepContains(m, v12+int32(16), v12+int32(12))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v230 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	F_pfree(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v233 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	F_pfree(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	if v228 != 0 {
		goto L62
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v237 = v208 + int32(1)
	if v237 != v195 {
		v208 = v237
		goto L63
	} else {
		goto L77
	}
L77:
	;
	goto L64
L78:
	;
	v262 = int32(0)
	goto L4
L79:
	;
	if v255 != int32(5) {
		v125 = v246
		v128 = v249
		goto L39
	} else {
		goto L80
	}
L80:
	;
	goto L40
L81:
	;
	F_errmsg_internal(m, int32(_a_F_JsonbDeepContains_0), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_JsonbDeepContains_1), int32(1428), int32(_a_F_JsonbDeepContains_2))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
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
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
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
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v5 {
	case 0:
		v295 = int32(1)
		v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v296, int32(1)) ^ v295
		return
	case 1:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v36 = v30 - int32(1636608432)
		if v29&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v30) {
				v145 = v29
				v146 = v30
				v147 = v36
				v148 = v36
				v149 = v36
				for {
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
					v152 = v151 + v148
					v153 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
					v155 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
					v156 = v155 + v149
					v158 = int32(4)
					v160 = v153 + v147 - v156 ^ base.I32_rotl(v156, v158)
					v164 = v152 - v160 ^ base.I32_rotl(v160, int32(6))
					v165 = v156 + v152
					v166 = v160 + v165
					v167 = v164 + v166
					v171 = v165 - v164 ^ base.I32_rotl(v164, int32(8))
					v175 = v166 - v171 ^ base.I32_rotl(v171, int32(16))
					v179 = v167 - v175 ^ base.I32_rotl(v175, int32(19))
					v180 = v171 + v167
					v181 = v175 + v180
					v182 = v179 + v181
					v186 = v180 - v179 ^ base.I32_rotl(v179, v158)
					v187 = int32(12)
					v188 = v145 + v187
					v190 = v146 - v187
					if base.Ui32(int32(11)) < base.Ui32(v190) {
						v145 = v188
						v146 = v190
						v147 = v181
						v148 = v182
						v149 = v186
						continue
					} else {
						break
					}
					break
				}
				v193 = v188
				v194 = v190
				v195 = v181
				v196 = v182
				v197 = v186
			} else {
				v193 = v29
				v194 = v30
				v195 = v36
				v196 = v36
				v197 = v36
			}
			switch v194 - int32(1) {
			case 0:
				v256 = v195
				v257 = v196
				v258 = v197
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 1:
				v249 = v195
				v250 = v196
				v251 = v197
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 2:
				v242 = v195
				v243 = v196
				v244 = v197
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 3:
				v236 = v196
				v237 = v197
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 4:
				v232 = v196
				v233 = v197
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 5:
				v226 = v196
				v227 = v197
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 6:
				v220 = v196
				v221 = v197
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
				v226 = v222<<(uint(int32(16))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 7:
				v215 = v197
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+7)))
				v220 = v216<<(uint(int32(24))%32) + v196
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
				v226 = v222<<(uint(int32(16))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 8:
				v210 = v197
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+8)))
				v215 = v211<<(uint(int32(8))%32) + v210
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+7)))
				v220 = v216<<(uint(int32(24))%32) + v196
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
				v226 = v222<<(uint(int32(16))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 9:
				v205 = v197
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+9)))
				v210 = v206<<(uint(int32(16))%32) + v205
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+8)))
				v215 = v211<<(uint(int32(8))%32) + v210
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+7)))
				v220 = v216<<(uint(int32(24))%32) + v196
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
				v226 = v222<<(uint(int32(16))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			case 10:
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+10)))
				v205 = v201<<(uint(int32(24))%32) + v197
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+9)))
				v210 = v206<<(uint(int32(16))%32) + v205
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+8)))
				v215 = v211<<(uint(int32(8))%32) + v210
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+7)))
				v220 = v216<<(uint(int32(24))%32) + v196
				v221 = v215
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+6)))
				v226 = v222<<(uint(int32(16))%32) + v220
				v227 = v221
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+5)))
				v232 = v228<<(uint(int32(8))%32) + v226
				v233 = v227
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)))
				v236 = v232 + v234
				v237 = v233
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+3)))
				v242 = v238<<(uint(int32(24))%32) + v195
				v243 = v236
				v244 = v237
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+2)))
				v249 = v245<<(uint(int32(16))%32) + v242
				v250 = v243
				v251 = v244
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
				v256 = v252<<(uint(int32(8))%32) + v249
				v257 = v250
				v258 = v251
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
				v263 = v256 + v259
				v264 = v257
				v265 = v258
			default:
				v263 = v195
				v264 = v196
				v265 = v197
			}
		} else {
			if base.Ui32(v30) < base.Ui32(int32(12)) {
				v91 = v29
				v92 = v30
				v93 = v36
				v94 = v36
				v95 = v36
			} else {
				v43 = v29
				v44 = v30
				v45 = v36
				v46 = v36
				v47 = v36
				for {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v50 = v49 + v46
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
					v54 = v53 + v47
					v56 = int32(4)
					v58 = v51 + v45 - v54 ^ base.I32_rotl(v54, v56)
					v62 = v50 - v58 ^ base.I32_rotl(v58, int32(6))
					v63 = v54 + v50
					v64 = v58 + v63
					v65 = v62 + v64
					v69 = v63 - v62 ^ base.I32_rotl(v62, int32(8))
					v73 = v64 - v69 ^ base.I32_rotl(v69, int32(16))
					v77 = v65 - v73 ^ base.I32_rotl(v73, int32(19))
					v78 = v69 + v65
					v79 = v73 + v78
					v80 = v77 + v79
					v84 = v78 - v77 ^ base.I32_rotl(v77, v56)
					v85 = int32(12)
					v86 = v43 + v85
					v88 = v44 - v85
					if base.Ui32(int32(11)) < base.Ui32(v88) {
						v43 = v86
						v44 = v88
						v45 = v79
						v46 = v80
						v47 = v84
						continue
					} else {
						break
					}
					break
				}
				v91 = v86
				v92 = v88
				v93 = v79
				v94 = v80
				v95 = v84
			}
			switch v92 - int32(1) {
			case 0:
				v142 = v93
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v263 = v142 + v143
				v264 = v94
				v265 = v95
			case 1:
				v137 = v93
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v142 = v138<<(uint(int32(8))%32) + v137
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v263 = v142 + v143
				v264 = v94
				v265 = v95
			case 2:
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+2)))
				v137 = v133<<(uint(int32(16))%32) + v93
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
				v142 = v138<<(uint(int32(8))%32) + v137
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
				v263 = v142 + v143
				v264 = v94
				v265 = v95
			case 3:
				v130 = v94
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v263 = v131 + v93
				v264 = v130
				v265 = v95
			case 4:
				v127 = v94
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v130 = v127 + v128
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v263 = v131 + v93
				v264 = v130
				v265 = v95
			case 5:
				v122 = v94
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v130 = v127 + v128
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v263 = v131 + v93
				v264 = v130
				v265 = v95
			case 6:
				v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+6)))
				v122 = v118<<(uint(int32(16))%32) + v94
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+5)))
				v127 = v123<<(uint(int32(8))%32) + v122
				v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
				v130 = v127 + v128
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v263 = v131 + v93
				v264 = v130
				v265 = v95
			case 7:
				v113 = v95
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				v263 = v114 + v93
				v264 = v116 + v94
				v265 = v113
			case 8:
				v108 = v95
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
				v113 = v109<<(uint(int32(8))%32) + v108
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				v263 = v114 + v93
				v264 = v116 + v94
				v265 = v113
			case 9:
				v103 = v95
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
				v108 = v104<<(uint(int32(16))%32) + v103
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
				v113 = v109<<(uint(int32(8))%32) + v108
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				v263 = v114 + v93
				v264 = v116 + v94
				v265 = v113
			case 10:
				v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+10)))
				v103 = v99<<(uint(int32(24))%32) + v95
				v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+9)))
				v108 = v104<<(uint(int32(16))%32) + v103
				v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+8)))
				v113 = v109<<(uint(int32(8))%32) + v108
				v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
				v263 = v114 + v93
				v264 = v116 + v94
				v265 = v113
			default:
				v263 = v93
				v264 = v94
				v265 = v95
			}
		}
		v268 = int32(14)
		v270 = v264 ^ v265 - base.I32_rotl(v264, v268)
		v274 = v270 ^ v263 - base.I32_rotl(v270, int32(11))
		v278 = v274 ^ v264 - base.I32_rotl(v274, int32(25))
		v282 = v278 ^ v270 - base.I32_rotl(v278, int32(16))
		v286 = v282 ^ v274 - base.I32_rotl(v282, int32(4))
		v290 = v286 ^ v278 - base.I32_rotl(v286, v268)
		v295 = v290 ^ v282 - base.I32_rotl(v290, int32(24))
		v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v296, int32(1)) ^ v295
		return
	case 2:
		v8 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+8)))
		v9 = F_DirectFunctionCall1Coll(m, int32(1469), int32(0), v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v295 = base.I32_wrap_i64(v9)
			v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v296, int32(1)) ^ v295
			return
		}
	case 3:
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v14 != 0 {
			v15 = int32(2)
		} else {
			v15 = int32(4)
		}
		v295 = v15
		v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_rotl(v296, int32(1)) ^ v295
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_JsonbHashScalarValue_0), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_JsonbHashScalarValue_1), int32(1467), int32(_a_F_JsonbHashScalarValue_2))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
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
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
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
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(48)
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v33
	v37 = int32(0)
	v40 = v37
	v41 = v37
	v43 = v5
	v44 = int32(1)
	v50 = v5
	goto L14
L14:
	;
	v53 = v40
	v54 = v41
	v56 = v43
	v57 = v44
	v61 = int32(0)
	v63 = v50
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
	v41 = v309
	v43 = v306
	v44 = v310
	v50 = l3
	goto L14
L18:
	;
	goto L17
L19:
	;
	v53 = int32(0)
	v54 = v441
	v56 = v442
	v57 = v443
	v63 = l3
	goto L16
L20:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	m.G0 = v15 + int32(48)
	return v435
L21:
	;
	v73 = F_JsonbIteratorNext(m, v15+int32(44), v15+int32(8), int32(0))
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
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L131
	}
L27:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v410 = int32(125)
	*(*uint8)(unsafe.Add(mBase, uint32(v408+v392))) = uint8(v410)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v414 = v412 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v414
	v416 = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	*(*uint8)(unsafe.Add(mBase, uint32(v417+v414))) = uint8(v416)
	v441 = int32(7)
	v442 = v362
	v443 = v416
	goto L19
L28:
	;
	v441 = v406
	v442 = v402
	v443 = int32(0)
	goto L19
L29:
	;
	v361 = int32(1)
	v362 = v56 - v361
	if v63&v361 != 0 {
		goto L120
	} else {
		goto L121
	}
L30:
	;
	v305 = int32(1)
	v306 = v56 - v305
	v309 = int32(5)
	v310 = int32(0)
	if v61&v305 != 0 {
		v53 = v310
		v54 = v309
		v56 = v306
		v57 = v310
		v61 = v305
		v63 = l3
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
	F_appendBinaryStringInfo(m, v23, int32(_a_F_JsonbToCStringWorker_0), v26)
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
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+24)))
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
	if v63&(v53^int32(-1))&int32(1) != 0 {
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
	v61 = v143
	v63 = l3
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
	F_appendBinaryStringInfo(m, v23, int32(_a_F_JsonbToCStringWorker_0), v26)
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
	if v63&(v53^int32(-1))&int32(1) != 0 {
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
	v441 = int32(6)
	v442 = v56 + v205
	v443 = v205
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
	F_appendBinaryStringInfo(m, v23, int32(_a_F_JsonbToCStringWorker_0), v26)
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
	if v63&int32(1) != 0 {
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
	v243 = v15 + int32(8)
	F_jsonb_put_escaped_value(m, v23, v243)
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
	F_appendBinaryStringInfo(m, v23, int32(_a_F_JsonbToCStringWorker_1), int32(2))
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
	v255 = F_JsonbIteratorNext(m, v15+int32(44), v243, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if v255 != int32(2) {
		v53 = v250
		v54 = v255
		v57 = v250
		v63 = l3
		goto L16
	} else {
		goto L90
	}
L90:
	;
	F_jsonb_put_escaped_value(m, v23, v243)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v402 = v56
	v406 = int32(2)
	goto L28
L92:
	;
	F_appendBinaryStringInfo(m, v23, int32(_a_F_JsonbToCStringWorker_0), v26)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if (v61|(v63^int32(-1)))&int32(1) == int32(0) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v274 <= v275+int32(1) {
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
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L105
	}
L99:
	;
	F_appendStringInfoSpaces(m, v23, v56<<(uint(int32(2))%32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L104
	}
L100:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v284 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v282+v275))) = uint8(v284)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v288 = v286 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v292 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290+v288))) = uint8(v292)
	goto L99
L103:
	;
	goto L99
L104:
	;
	goto L98
L105:
	;
	v402 = v56
	v406 = int32(3)
	goto L28
L106:
	;
	if v63&int32(1) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v315 <= v316+int32(1) {
		goto L111
	} else {
		goto L112
	}
L108:
	;
	goto L109
L109:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v341 <= v342+int32(1) {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	F_appendStringInfoSpaces(m, v23, v306<<(uint(int32(2))%32))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L115
	}
L111:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L4
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v325 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v323+v316))) = uint8(v325)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v329 = v327 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v329
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v333 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v331+v329))) = uint8(v333)
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
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v351 = int32(93)
	*(*uint8)(unsafe.Add(mBase, uint32(v349+v342))) = uint8(v351)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v355 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v359 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v357+v355))) = uint8(v359)
	goto L18
L119:
	;
	goto L18
L120:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v365 <= v366+int32(1) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v392+int32(1) < v391 {
		goto L27
	} else {
		goto L129
	}
L123:
	;
	F_appendStringInfoSpaces(m, v23, v362<<(uint(int32(2))%32))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L128
	}
L124:
	;
	F_appendStringInfoChar(m, v23, int32(10))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v375 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v373+v366))) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v379 = v377 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v379
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v383 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v381+v379))) = uint8(v383)
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
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L130
	}
L130:
	;
	v402 = v362
	v406 = int32(7)
	goto L28
L131:
	;
	F_errmsg_internal(m, int32(_a_F_JsonbToCStringWorker_2), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_JsonbToCStringWorker_3), int32(595), int32(_a_F_JsonbToCStringWorker_4))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	v5 = m.G0
	v7 = v5 - int32(80)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v10 {
	case 0:
		v99 = int32(_a_F_JsonbTypeName_0)
		m.G0 = v7 + int32(80)
		return v99
	case 1:
		v99 = int32(_a_F_JsonbTypeName_1)
		m.G0 = v7 + int32(80)
		return v99
	case 2:
		v99 = int32(_a_F_JsonbTypeName_2)
		m.G0 = v7 + int32(80)
		return v99
	case 3:
		v99 = int32(_a_F_JsonbTypeName_3)
		m.G0 = v7 + int32(80)
		return v99
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v66 = m.ExcPending
		if v66 != 0 {
			return int32(0)
		} else {
			v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v67
			F_errmsg_internal(m, int32(_a_F_JsonbTypeName_4), v7)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_JsonbTypeName_5), int32(209), int32(_a_F_JsonbTypeName_6))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v99 = int32(_a_F_JsonbTypeName_7)
		m.G0 = v7 + int32(80)
		return v99
	case 17:
		v99 = int32(_a_F_JsonbTypeName_8)
		m.G0 = v7 + int32(80)
		return v99
	case 18:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = v7 + int32(48)
		v14 = F_JsonbExtractScalar(m, v11, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v14 != 0 {
				v18 = F_JsonbTypeName(m, v13)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v99 = v18
					m.G0 = v7 + int32(80)
					return v99
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				if v20&int32(1073741824) != 0 {
					v99 = int32(_a_F_JsonbTypeName_7)
					m.G0 = v7 + int32(80)
					return v99
				} else {
					if v20&int32(536870912) == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v84
							F_errmsg_internal(m, int32(_a_F_JsonbTypeName_9), v7+int32(16))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_JsonbTypeName_5), int32(163), int32(_a_F_JsonbTypeName_10))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v99 = int32(_a_F_JsonbTypeName_8)
						m.G0 = v7 + int32(80)
						return v99
					}
				}
			}
		}
	case 32:
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v32 <= int32(1183) {
			switch v32 - int32(1082) {
			case 0:
				v99 = int32(_a_F_JsonbTypeName_11)
				m.G0 = v7 + int32(80)
				return v99
			case 1:
				v99 = int32(_a_F_JsonbTypeName_12)
				m.G0 = v7 + int32(80)
				return v99
			default:
				if v32 == int32(1114) {
					v99 = int32(_a_F_JsonbTypeName_13)
					m.G0 = v7 + int32(80)
					return v99
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v51
						F_errmsg_internal(m, int32(_a_F_JsonbTypeName_14), v7+int32(32))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonbTypeName_5), int32(205), int32(_a_F_JsonbTypeName_6))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
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
			if v32 != int32(1184) {
				if v32 != int32(1266) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v51
						F_errmsg_internal(m, int32(_a_F_JsonbTypeName_14), v7+int32(32))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_JsonbTypeName_5), int32(205), int32(_a_F_JsonbTypeName_6))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v99 = int32(_a_F_JsonbTypeName_15)
					m.G0 = v7 + int32(80)
					return v99
				}
			} else {
				v99 = int32(_a_F_JsonbTypeName_16)
				m.G0 = v7 + int32(80)
				return v99
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	v5 = m.G0
	v7 = v5 - int32(80)
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+76)) = v10
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v16
	goto L4
L4:
	;
	v28 = F_JsonbIteratorNext(m, v7+int32(76), v7+int32(40), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L5:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v7)+76))
	if v138 != 0 {
		goto L62
	} else {
		goto L63
	}
L6:
	;
	goto L5
L7:
	;
	if v132 == int32(0) {
		goto L4
	} else {
		goto L61
	}
L8:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v7)+52))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v128 = F_varstr_cmp(m, v123, v124, v125, v126, int32(100))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L60
	}
L9:
	;
	v35 = F_JsonbIteratorNext(m, v7+int32(72), v7+int32(8), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v28 == v35 {
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if base.Ui32(v120) < base.Ui32(v119) {
		goto L57
	} else {
		goto L58
	}
L14:
	;
	v137 = int32(0)
	goto L6
L15:
	;
	goto L16
L16:
	;
	v41 = int32(0)
	v42 = int32(5)
	if v28&v42 == v42 {
		v132 = v41
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	if v46 == v47 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	switch v46 - int32(1) {
	case 0:
		goto L8
	case 1:
		goto L26
	case 2:
		goto L25
	default:
		v132 = v41
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
	if base.Ui32(v47) < base.Ui32(v46) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L51
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L48
	}
L23:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v80 == v81 {
		v132 = v41
		goto L7
	} else {
		goto L44
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v65 == v66 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+48)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)))
	if v58 == v59 {
		v132 = v41
		goto L7
	} else {
		goto L28
	}
L26:
	;
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+48)))
	v54 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v7)+16)))
	v55 = F_DirectFunctionCall2Coll(m, int32(1467), int32(0), v53, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v132 = base.I32_wrap_i64(v55)
	goto L7
L28:
	;
	if base.Ui32(v59) < base.Ui32(v58) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v64 = int32(1)
	goto L31
L30:
	;
	v64 = int32(-1)
	goto L31
L31:
	;
	v137 = v64
	goto L6
L32:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+56)))
	if v70 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v66 < v65 {
		goto L41
	} else {
		goto L42
	}
L35:
	;
	v71 = int32(-1)
	goto L37
L36:
	;
	v71 = int32(1)
	goto L37
L37:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+24)))
	if v70 != v73 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v75 = v71
	goto L40
L39:
	;
	v75 = int32(0)
	goto L40
L40:
	;
	v132 = v75
	goto L7
L41:
	;
	v79 = int32(1)
	goto L43
L42:
	;
	v79 = int32(-1)
	goto L43
L43:
	;
	v137 = v79
	goto L6
L44:
	;
	if v81 < v80 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v86 = int32(1)
	goto L47
L46:
	;
	v86 = int32(-1)
	goto L47
L47:
	;
	v137 = v86
	goto L6
L48:
	;
	F_errmsg_internal(m, int32(_a_F_compareJsonbContainers_0), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_compareJsonbContainers_1), int32(267), int32(_a_F_compareJsonbContainers_2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_compareJsonbContainers_3), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_compareJsonbContainers_1), int32(270), int32(_a_F_compareJsonbContainers_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
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
	v116 = int32(1)
	goto L56
L55:
	;
	v116 = int32(-1)
	goto L56
L56:
	;
	v137 = v116
	goto L6
L57:
	;
	v122 = int32(1)
	goto L59
L58:
	;
	v122 = int32(-1)
	goto L59
L59:
	;
	v137 = v122
	goto L6
L60:
	;
	v132 = v128
	goto L7
L61:
	;
	v137 = v132
	goto L6
L62:
	;
	v139 = v138
	goto L65
L63:
	;
	goto L64
L64:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v7)+72))
	if v150 != 0 {
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+36))
	F_pfree(m, v139)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	if v143 != 0 {
		v139 = v143
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v151 = v150
	goto L72
L70:
	;
	goto L71
L71:
	;
	m.G0 = v7 + int32(80)
	return v137
L72:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+36))
	F_pfree(m, v151)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L71
L74:
	;
	if v155 != 0 {
		v151 = v155
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
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
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
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v4 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = v9 & int32(268435455)
	if v11 == v4 {
		v204 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v204
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
	v23 = F_palloc(m, int32(32))
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
		v204 = v4
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
	v32 = v27
	goto L8
L8:
	;
	v44 = l0 + v30<<(uint(int32(2))%32) + int32(4)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	switch int32(base.Ui32(v45)>>(uint(int32(28))%32)) & int32(7) {
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
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L45
	}
L10:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v168 == v169 {
		goto L36
	} else {
		goto L37
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(18)
	v112 = (v32 + int32(3)) & int32(-4)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v21 + v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v115 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v103)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(3)
	goto L10
L13:
	;
	v99 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v99)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(3)
	goto L10
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v21 + (v32+int32(3))&int32(-4)
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v21 + v32
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v56 < int32(0) {
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
	v61 = v30
	v62 = int32(0)
	goto L20
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v56 & int32(268435455)
	goto L10
L20:
	;
	v69 = v61 - int32(1)
	if int32(0) <= v69 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v56&int32(268435455) - v82
	goto L10
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0+v61<<(uint(int32(2))%32))))
	v78 = v75&int32(268435455) + v62
	if int32(0) <= v75 {
		v61 = v69
		v62 = v78
		goto L20
	} else {
		goto L25
	}
L23:
	;
	v82 = v62
	goto L24
L24:
	;
	goto L21
L25:
	;
	v82 = v78
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v156 + (v32 - v112)
	goto L10
L27:
	;
	v120 = v30
	v121 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v156 = v115 & int32(268435455)
	goto L26
L30:
	;
	v128 = v120 - int32(1)
	if int32(0) <= v128 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v156 = v115&int32(268435455) - v141
	goto L26
L32:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0+v120<<(uint(int32(2))%32))))
	v137 = v134&int32(268435455) + v121
	if int32(0) <= v134 {
		v120 = v128
		v121 = v137
		goto L30
	} else {
		goto L35
	}
L33:
	;
	v141 = v121
	goto L34
L34:
	;
	goto L31
L35:
	;
	v141 = v137
	goto L34
L36:
	;
	v171 = F_equalsJsonbScalarValue(m, l2, v23)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v173 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v18+v30<<(uint(int32(2))%32))))
	if v173 <= v177 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v171 != 0 {
		v204 = v23
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v180 = v32
	goto L43
L42:
	;
	v180 = v173
	goto L43
L43:
	;
	v185 = v30 + int32(1)
	if v185 != v11 {
		v30 = v185
		v32 = v180 + v177&int32(268435455)
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
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v198 = F_getKeyJsonValueFromContainer(m, l0, v195, v196, int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v204 = v198
	goto L1
}
func F_jsonb_agg_finalfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14320(m, l0, int32(5))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_agg_strict_transfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_agg_transfn_worker(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_bool(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v18 = F_JsonbExtractScalar(m, v12+int32(4), v9)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			if v18 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_cannotCastJsonbValue(m, v20, int32(_a_F_jsonb_bool_0), v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v45 = int64(0)
					m.G0 = v9 + int32(32)
					return v45
				}
			} else {
				switch v20 {
				case 0:
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v28 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							v32 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
							v45 = int64(0)
							m.G0 = v9 + int32(32)
							return v45
						}
					} else {
						v32 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
						v45 = int64(0)
						m.G0 = v9 + int32(32)
						return v45
					}
				default:
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_cannotCastJsonbValue(m, v20, int32(_a_F_jsonb_bool_0), v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v45 = int64(0)
						m.G0 = v9 + int32(32)
						return v45
					}
				case 3:
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v40 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)))
							v45 = v44
							m.G0 = v9 + int32(32)
							return v45
						}
					} else {
						v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)))
						v45 = v44
						m.G0 = v9 + int32(32)
						return v45
					}
				}
			}
		}
	}
}
func F_jsonb_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_s(v17)
							}
						} else {
							return base.I64_extend_i32_s(v17)
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v17)
						}
					} else {
						return base.I64_extend_i32_s(v17)
					}
				}
			}
		}
	}
}
func F_jsonb_eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v17 == int32(0)))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 == int32(0)))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v17 == int32(0)))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(v17 == int32(0)))
					}
				}
			}
		}
	}
}
func F_jsonb_ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v17 = F_compareJsonbContainers(m, v6+int32(4), v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v19 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v23 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(int32(0) <= v17))
							}
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v17))
						}
					}
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v23 != v13 {
						F_pfree(m, v13)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(int32(0) <= v17))
						}
					} else {
						return base.I64_extend_i32_u(base.B2i32(int32(0) <= v17))
					}
				}
			}
		}
	}
}
func F_jsonb_hash(m *base.Module, l0 int32) int64 {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v75 int64
	_ = v75
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v17&int32(268435455) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = F_JsonbIteratorInit(m, v11+int32(4))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v75 = int64(0)
	goto L5
L5:
	;
	m.G0 = v8 + int32(48)
	return v75
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v22
	goto L8
L7:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v64 != v11 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v35 = F_JsonbIteratorNext(m, v8+int32(44), v8+int32(8), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L16
	}
L10:
	;
	goto L9
L11:
	;
	F_JsonbHashScalarValue(m, v8+int32(8), v8+int32(4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v41 ^ int32(536870912)
	goto L8
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v37 ^ int32(1073741824)
	goto L8
L14:
	;
	switch v35 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v35
	F_errmsg_internal(m, int32(_a_F_jsonb_hash_0), v8)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_jsonb_hash_1), int32(286), int32(_a_F_jsonb_hash_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	F_pfree(m, v11)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+4)))
	v75 = v68
	goto L5
L22:
	;
	goto L21
}
func F_jsonb_in_object_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	F_pushJsonbValue(m, l0, int32(6), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)) = uint8(v9)
		return int32(0)
	}
}
func F_jsonb_object_agg_finalfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14320(m, l0, int32(7))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_object_agg_transfn(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_jsonb_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_jsonb_object_field(m *base.Module, l0 int32) int64 {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
			if v19&int32(32) == int32(0) {
				v68 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
				v74 = int64(0)
				m.G0 = v9 + int32(32)
				return v74
			} else {
				v24 = int32(4)
				v26 = int32(1)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				v30 = v28 & v26
				if v30 != 0 {
					v31 = v26
				} else {
					v31 = v24
				}
				if v28 == int32(1) {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
					if v38 == int32(18) {
						v41 = int32(16)
					} else {
						v41 = int32(0)
					}
					if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v48 = int32(4)
					} else {
						v48 = v41
					}
					v59 = v48
				} else {
					v49 = int32(1)
					if v30 != 0 {
						v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v60 = F_getKeyJsonValueFromContainer(m, v12+v24, v17+v31, v59, v9)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					if v60 != 0 {
						v62 = F_JsonbValueToJsonb(m, v60)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int64(0)
						} else {
							v74 = base.I64_extend_i32_u(v62)
							m.G0 = v9 + int32(32)
							return v74
						}
					} else {
						v68 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v68)
						v74 = int64(0)
						m.G0 = v9 + int32(32)
						return v74
					}
				}
			}
		}
	}
}
func F_jsonb_path_query(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_pretty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v25 int32
	_ = v25
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		F_initStringInfo(m, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v20 = F_JsonbToCStringWorker(m, v5, v8+int32(4), int32(base.Ui32(v16)>>(uint(int32(2))%32)), int32(1))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				v24 = F_cstring_to_text_with_len(m, v22, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					m.G0 = v5 + int32(16)
					return base.I64_extend_i32_u(v24)
				}
			}
		}
	}
}
func F_jsonb_set(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v20 = F_pg_detoast_datum(m, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(0)
				v25 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v25
				*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v25
				v30 = v9 + int32(48)
				*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(18)
				v33 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v20 + v33
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - v33
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v42 < int32(2) {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
					if v45&int32(268435456) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_jsonb_set_0), int32(0))
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_jsonb_set_1), int32(_a_F_jsonb_set_2), int32(_a_F_jsonb_set_3))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if base.B2i32(v45&int32(268435455) == int32(0))&base.B2i32(v22 == int64(0)) != 0 {
							v90 = v12
							m.G0 = v9 + int32(80)
							return base.I64_extend_i32_u(v90)
						} else {
							F_deconstruct_array_builtin(m, v17, int32(25), v9+int32(44), v9+int32(40), v9+int32(36))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int64(0)
							} else {
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
								if v64 == int32(0) {
									v90 = v12
									m.G0 = v9 + int32(80)
									return base.I64_extend_i32_u(v90)
								} else {
									v69 = F_JsonbIteratorInit(m, v12+int32(4))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v69
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
										v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										if v22 != int64(0) {
											v84 = int32(1)
										} else {
											v84 = int32(4)
										}
										F_setPath(m, v9+int32(32), v74, v75, v76, v9+int32(8), int32(0), v30, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int64(0)
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
											v88 = F_JsonbValueToJsonb(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int64(0)
											} else {
												v90 = v88
												m.G0 = v9 + int32(80)
												return base.I64_extend_i32_u(v90)
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
					v99 = m.ExcPending
					if v99 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(352845954))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_jsonb_set_4), int32(0))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_jsonb_set_1), int32(_a_F_jsonb_set_5), int32(_a_F_jsonb_set_3))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int64(0)
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
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+64)) = uint8(v9)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = int64(0)
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
			v22 = F_jsonb_get_element(m, v15, v17, v18, v4-int32(-64), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v4)+56)) = v22
				return
			}
		}
	}
}
func F_jsonb_subscript_handler(m *base.Module, l0 int32) int64 {
	return int64(1745092)
}
func F_jsonb_to_tsvector(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = F_parse_jsonb_index_flags(m, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_getTSCurrentConfig(m)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v21
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v24
					v29 = v9 + int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v29
					F_iterate_jsonb_values(m, v12, v19, v9+int32(24))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int64(0)
					} else {
						v35 = F_make_tsvector(m, v29)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v37 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v41 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(32)
											return base.I64_extend_i32_u(v35)
										}
									} else {
										m.G0 = v9 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								}
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v41 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(32)
										return base.I64_extend_i32_u(v35)
									}
								} else {
									m.G0 = v9 + int32(32)
									return base.I64_extend_i32_u(v35)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pushJsonbValue(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(48)
	return
L2:
	;
	switch v12 - int32(16) {
	case 0:
		goto L10
	case 1:
		goto L11
	default:
		goto L9
	}
L3:
	;
	F_pushJsonbValueScalar(m, l0, l1, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = int32(4)
	if base.B2i32(base.Ui32(v12) < base.Ui32(v13))|base.B2i32(base.Ui32(l1-v13) < base.Ui32(int32(-2))) != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v12 != int32(32) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	return
L8:
	;
	goto L1
L9:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v148 = F_iteratorFromContainer(m, v146, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L7
	} else {
		goto L45
	}
L10:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValue[0]))
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValue[0]))
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = v27
	goto L14
L13:
	;
	v30 = v29
	goto L14
L14:
	;
	v32 = F_MemoryContextAlloc(m, v30, int32(48))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+40)) = uint16(v35)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+36)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(17)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValue[0]))
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = v45
	goto L18
L17:
	;
	v48 = v47
	goto L18
L18:
	;
	v50 = F_MemoryContextAlloc(m, v48, int32(288))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+12)) = v50
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v53 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v58 = int32(0)
	goto L23
L21:
	;
	goto L22
L22:
	;
	F_pushJsonbValueScalar(m, l0, int32(7), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L28
	}
L23:
	;
	v64 = v58 * int32(72)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_pushJsonbValueScalar(m, l0, int32(1), v64+v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_pushJsonbValue(m, l0, int32(2), v70+v64+int32(32))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v77 = v58 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v77 < v78 {
		v58 = v77
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	goto L1
L29:
	;
	v92 = v89
	goto L31
L30:
	;
	v92 = v91
	goto L31
L31:
	;
	v94 = F_MemoryContextAlloc(m, v92, int32(48))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v97 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v94)+40)) = uint16(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+36)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v97)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(16)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValue[0]))
	if v109 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v112 = v109
	goto L35
L34:
	;
	v112 = v111
	goto L35
L35:
	;
	v114 = F_MemoryContextAlloc(m, v112, int32(128))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v114
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if int32(0) < v117 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v122 = int32(0)
	goto L40
L38:
	;
	goto L39
L39:
	;
	F_pushJsonbValueScalar(m, l0, int32(5), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L7
	} else {
		goto L44
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_pushJsonbValue(m, l0, int32(3), v127+v122<<(uint(int32(5))%32))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L42
	}
L41:
	;
	goto L39
L42:
	;
	v134 = v122 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v134 < v135 {
		v122 = v134
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	goto L1
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+3)))
	if v152&int32(16) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v180 = F_JsonbIteratorNext(m, v8+int32(44), v8+int32(8), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L7
	} else {
		goto L53
	}
L47:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v157 == int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v161 = v8 + int32(44)
	v163 = v8 + int32(8)
	v165 = F_JsonbIteratorNext(m, v161, v163, int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v168 = F_JsonbIteratorNext(m, v161, v163, int32(1))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	F_pushJsonbValueScalar(m, l0, l1, v163)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	v173 = F_JsonbIteratorNext(m, v161, v163, int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	goto L1
L53:
	;
	if v180 == int32(0) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v186 = v180
	goto L55
L55:
	;
	v190 = v8 + int32(8)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
	if v192&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L1
L57:
	;
	v195 = v190
	goto L59
L58:
	;
	v195 = int32(0)
	goto L59
L59:
	;
	if v186 == int32(4) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v199 = v195
	goto L62
L61:
	;
	v199 = int32(0)
	goto L62
L62:
	;
	if base.Ui32(v186) < base.Ui32(int32(4)) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v202 = v190
	goto L65
L64:
	;
	v202 = v199
	goto L65
L65:
	;
	F_pushJsonbValueScalar(m, l0, v186, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v208 = F_JsonbIteratorNext(m, v8+int32(44), v190, int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	if v208 != 0 {
		v186 = v208
		goto L55
	} else {
		goto L68
	}
L68:
	;
	goto L56
}
func F_pushJsonbValueScalar(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
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
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v306 int64
	_ = v306
	var v308 int64
	_ = v308
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
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
	v403 = m.ExcPending
	if v403 != 0 {
		goto L15
	} else {
		goto L106
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L15
	} else {
		goto L102
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return
L4:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValueScalar[0]))
	if v343 != 0 {
		goto L94
	} else {
		goto L95
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L15
	} else {
		goto L91
	}
L6:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v291
	if v291 != 0 {
		goto L81
	} else {
		goto L82
	}
L7:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+41)))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+40)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)) = uint8(v127)
	if int32(2) <= v126 {
		goto L39
	} else {
		goto L40
	}
L8:
	;
	F_appendElement(m, l0, l2, int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L15
	} else {
		goto L38
	}
L9:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	v101 = v97 + v98*int32(72)
	v102 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v101)+56)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v101)+48)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v101)+40)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v101)+32)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v110 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_copyScalarSubstructure(m, v101+int32(32), v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L15
	} else {
		goto L37
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	if base.Ui32(v58) < base.Ui32(v59) {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValueScalar[0]))
	if v18 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v21 = v18
	goto L14
L13:
	;
	v21 = v20
	goto L14
L14:
	;
	v23 = F_MemoryContextAlloc(m, v21, int32(48))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+40)) = uint16(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+36)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(16)
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValueScalar[0]))
	if v50 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = int32(4)
	v49 = int32(128)
	goto L17
L19:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)) = uint8(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v36 <= int32(0) {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+16)) = uint8(v42)
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v36
	v49 = v36 << (uint(int32(5)) % 32)
	goto L17
L23:
	;
	v53 = v50
	goto L25
L24:
	;
	v53 = v52
	goto L25
L25:
	;
	v54 = F_MemoryContextAlloc(m, v53, v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v54
	goto L3
L27:
	;
	v82 = v79 + v78*int32(72)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v82)+16)) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+64)) = v91
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_copyScalarSubstructure(m, v82, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L15
	} else {
		goto L36
	}
L28:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v78 = v58
	v79 = v61
	goto L27
L29:
	;
	goto L30
L30:
	;
	if base.Ui32(int32(14913080)) <= base.Ui32(v58) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v64 = int32(14913080)
	v66 = v59 << (uint(int32(1)) % 32)
	if base.Ui32(v64) <= base.Ui32(v66) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v69 = v64
	goto L34
L33:
	;
	v69 = v66
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+32)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v74 = F_repalloc(m, v71, v69*int32(72))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
	v78 = v77
	v79 = v74
	goto L27
L36:
	;
	goto L3
L37:
	;
	goto L3
L38:
	;
	goto L3
L39:
	;
	F_qsort_arg(m, v123, v126, int32(72), int32(1466), v14+int32(15))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L15
	} else {
		goto L42
	}
L40:
	;
	v142 = v127
	goto L41
L41:
	;
	v143 = int32(1)
	if v142&v143|v124&v143 == int32(0) {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v138&v125&int32(1) != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v142 = v138
	goto L41
L44:
	;
	v150 = int32(0)
	if v150 < v126 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v155 = int32(0)
	v156 = v150
	goto L48
L46:
	;
	v269 = v150
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+8)) = v269
	goto L6
L48:
	;
	v167 = v123 + v155*int32(72)
	if v156 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v269 = v261
	goto L47
L50:
	;
	v265 = v155 + int32(1)
	if v265 != v126 {
		v155 = v265
		v156 = v261
		goto L48
	} else {
		goto L80
	}
L51:
	;
	if v124&int32(1) != 0 {
		goto L73
	} else {
		goto L74
	}
L52:
	;
	v172 = v123 + v156*int32(72)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172+int32(-64))))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	if v175 != v176 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v172-int32(60))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v175) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v243 == int32(0) {
		v261 = v156
		goto L50
	} else {
		goto L72
	}
L55:
	;
	v243 = int32(0)
	goto L54
L56:
	;
	v217 = v212
	v218 = v213
	v219 = v214
	goto L66
L57:
	;
	if (v180|v181)&int32(3) != 0 {
		v212 = v180
		v213 = v181
		v214 = v175
		goto L56
	} else {
		goto L60
	}
L58:
	;
	v205 = v180
	v206 = v181
	v207 = v175
	goto L59
L59:
	;
	if v207 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v189 = v180
	v190 = v181
	v191 = v175
	goto L61
L61:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	if v194 != v195 {
		v212 = v189
		v213 = v190
		v214 = v191
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v205 = v200
	v206 = v198
	v207 = v202
	goto L59
L63:
	;
	v197 = int32(4)
	v198 = v190 + v197
	v200 = v189 + v197
	v202 = v191 - v197
	if base.Ui32(int32(3)) < base.Ui32(v202) {
		v189 = v200
		v190 = v198
		v191 = v202
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v212 = v205
	v213 = v206
	v214 = v207
	goto L56
L66:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v222 == v223 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v243 = v222 - v223
	goto L54
L68:
	;
	v225 = int32(1)
	v230 = v219 - v225
	if v230 != 0 {
		v217 = v217 + v225
		v218 = v218 + v225
		v219 = v230
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
	goto L51
L73:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	if v250 == int32(0) {
		v261 = v156
		goto L50
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if v156 < v155 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	v254 = int32(72)
	base.MemoryCopy(m, v123+v156*v254, v167, v254)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v261 = v156 + int32(1)
	goto L50
L80:
	;
	goto L49
L81:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	switch v293 - int32(16) {
	case 0:
		goto L86
	case 1:
		goto L85
	default:
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v290
	goto L3
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L15
	} else {
		goto L88
	}
L85:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	v303 = v299 + v300*int32(72)
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v290)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+56)) = v304
	v306 = *(*int64)(unsafe.Add(mBase, uint32(v290)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+48)) = v306
	v308 = *(*int64)(unsafe.Add(mBase, uint32(v290)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+40)) = v308
	v310 = *(*int64)(unsafe.Add(mBase, uint32(v290)))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+32)) = v310
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v291)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v312 + int32(1)
	goto L3
L86:
	;
	F_appendElement(m, l0, v290, int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L15
	} else {
		goto L87
	}
L87:
	;
	goto L3
L88:
	;
	F_errmsg_internal(m, int32(_a_F_pushJsonbValueScalar_0), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L15
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_pushJsonbValueScalar_1), int32(748), int32(_a_F_pushJsonbValueScalar_2))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L15
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errmsg_internal(m, int32(_a_F_pushJsonbValueScalar_3), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L15
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_pushJsonbValueScalar_1), int32(755), int32(_a_F_pushJsonbValueScalar_2))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L15
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	v346 = v343
	goto L96
L95:
	;
	v346 = v345
	goto L96
L96:
	;
	v348 = F_MemoryContextAlloc(m, v346, int32(48))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L15
	} else {
		goto L97
	}
L97:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v351 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v348)+40)) = uint16(v351)
	*(*int32)(unsafe.Add(mBase, uint32(v348)+36)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v348
	*(*int32)(unsafe.Add(mBase, uint32(v348)+32)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v348)+8)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = int32(17)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_pushJsonbValueScalar[0]))
	if v361 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v364 = v361
	goto L100
L99:
	;
	v364 = v363
	goto L100
L100:
	;
	v366 = F_MemoryContextAlloc(m, v364, int32(288))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L15
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348)+12)) = v366
	goto L3
L102:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L15
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(14913080)
	F_errmsg(m, int32(_a_F_pushJsonbValueScalar_4), v14)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L15
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_pushJsonbValueScalar_1), int32(800), int32(_a_F_pushJsonbValueScalar_5))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L15
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(_a_F_pushJsonbValueScalar_6))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L15
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_pushJsonbValueScalar_7), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_pushJsonbValueScalar_1), int32(2085), int32(_a_F_pushJsonbValueScalar_8))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
