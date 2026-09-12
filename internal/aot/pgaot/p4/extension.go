package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__equalCreateExtensionStmt(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v7 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v49
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
	if v38 != v39 {
		goto L17
	} else {
		goto L18
	}
L3:
	;
	if v6 == int32(0) {
		v49 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if v6 == v7 {
		goto L2
	} else {
		goto L16
	}
L6:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v33-v32 != 0 {
		v49 = v3
		goto L1
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v17 = v7
	v18 = v6
	goto L11
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v32 = v21
	v33 = v22
	goto L8
L13:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L2
L16:
	;
	return int32(0)
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v45 = F_equal(m, v43, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v49 = v45
	goto L1
}
func F_get_extension_control_directories(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	v9 = m.G0
	v11 = v9 - int32(1056)
	m.G0 = v11
	F_get_share_path(m, v11+int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v11 + int32(32)
	v25 = F_psprintf(m, int32(267173), v11+int32(16))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[442]))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v11 + int32(1056)
	return v283
L5:
	;
	v33 = F_lappend(m, int32(0), v25)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v35 = F_pstrdup(m, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v283 = v33
	goto L4
L9:
	;
	v37 = v35
	v40 = int32(0)
	goto L10
L10:
	;
	v46 = v37
	goto L14
L11:
	;
	v283 = v276
	goto L4
L12:
	;
	v119 = v117 + int32(1)
	v120 = F_palloc(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L13:
	;
	if v56 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v48 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L13
L16:
	;
	goto L15
L17:
	;
	v56 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v48 == int32(58) {
		v56 = v46
		goto L16
	} else {
		goto L20
	}
L20:
	;
	v46 = v46 + int32(1)
	goto L14
L21:
	;
	if v37&int32(3) == int32(0) {
		v82 = v37
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L23
L23:
	;
	v117 = v56 - v37
	goto L12
L24:
	;
	v117 = v115
	goto L12
L25:
	;
	v115 = v107 - v37
	goto L24
L26:
	;
	v86 = v82
	goto L35
L27:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v66 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v115 = int32(0)
	goto L24
L29:
	;
	goto L30
L30:
	;
	v71 = v37
	goto L31
L31:
	;
	v75 = v71 + int32(1)
	if v75&int32(3) == int32(0) {
		v82 = v75
		goto L26
	} else {
		goto L33
	}
L32:
	;
	v107 = v75
	goto L25
L33:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v80 != 0 {
		v71 = v75
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v95 = int32(-2139062144)
	if (int32(16843008)-v92|v92)&v95 == v95 {
		v86 = v86 + int32(4)
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v101 = v86
	goto L38
L37:
	;
	goto L36
L38:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v105 != 0 {
		v101 = v101 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v107 = v101
	goto L25
L40:
	;
	goto L39
L41:
	;
	if v119 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v236 = int32(284401)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, _consts[443])))
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v240 == int32(0) {
		v259 = v239
		v260 = v240
		goto L76
	} else {
		goto L77
	}
L43:
	;
	v233 = F_strlen(m, v229)
	mBase = m.M
	goto L42
L44:
	;
	v229 = v37
	goto L43
L45:
	;
	goto L46
L46:
	;
	v127 = v119 - int32(1)
	if (v120^v37)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v226)
	v229 = v222
	goto L43
L48:
	;
	v207 = v202
	v208 = v203
	v209 = v204
	goto L70
L49:
	;
	if v197 == int32(0) {
		v222 = v195
		v223 = v196
		goto L47
	} else {
		goto L69
	}
L50:
	;
	v195 = v37
	v196 = v120
	v197 = v127
	goto L49
L51:
	;
	goto L52
L52:
	;
	v131 = int32(0)
	if v37&int32(3) == v131 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v164 == int32(0) {
		v222 = v161
		v223 = v162
		goto L47
	} else {
		goto L62
	}
L54:
	;
	v161 = v37
	v162 = v120
	v163 = v127
	v164 = base.B2i32(v127 != v131)
	goto L53
L55:
	;
	if v127 == int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v140 = v37
	v141 = v120
	v142 = v127
	goto L57
L57:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v144)
	if v144 == int32(0) {
		v202 = v140
		v203 = v141
		v204 = v142
		goto L48
	} else {
		goto L59
	}
L58:
	;
	v161 = v155
	v162 = v149
	v163 = v151
	v164 = v153
	goto L53
L59:
	;
	v148 = int32(1)
	v149 = v141 + v148
	v151 = v142 - v148
	v152 = int32(0)
	v153 = base.B2i32(v151 != v152)
	v155 = v140 + v148
	if v155&int32(3) == v152 {
		v161 = v155
		v162 = v149
		v163 = v151
		v164 = v153
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v151 != 0 {
		v140 = v155
		v141 = v149
		v142 = v151
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v167 == int32(0) {
		v195 = v161
		v196 = v162
		v197 = v163
		goto L49
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(v163) < base.Ui32(int32(4)) {
		v195 = v161
		v196 = v162
		v197 = v163
		goto L49
	} else {
		goto L64
	}
L64:
	;
	v173 = v161
	v174 = v162
	v175 = v163
	goto L65
L65:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v181 = int32(-2139062144)
	if (int32(16843008)-v178|v178)&v181 != v181 {
		v202 = v173
		v203 = v174
		v204 = v175
		goto L48
	} else {
		goto L67
	}
L66:
	;
	v195 = v189
	v196 = v187
	v197 = v191
	goto L49
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174))) = v178
	v186 = int32(4)
	v187 = v174 + v186
	v189 = v173 + v186
	v191 = v175 - v186
	if base.Ui32(int32(3)) < base.Ui32(v191) {
		v173 = v189
		v174 = v187
		v175 = v191
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v202 = v195
	v203 = v196
	v204 = v197
	goto L48
L70:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v211)
	if v211 == int32(0) {
		v222 = v207
		v223 = v208
		goto L47
	} else {
		goto L72
	}
L71:
	;
	v222 = v218
	v223 = v216
	goto L47
L72:
	;
	v215 = int32(1)
	v216 = v208 + v215
	v218 = v207 + v215
	v220 = v209 - v215
	if v220 != 0 {
		v207 = v218
		v208 = v216
		v209 = v220
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	F_pfree(m, v120)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L88
	}
L75:
	;
	if v260-v259 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L76:
	;
	goto L75
L77:
	;
	if v239 != v240 {
		v259 = v239
		v260 = v240
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v244 = v120
	v245 = v236
	goto L79
L79:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+1)))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+1)))
	if v249 == int32(0) {
		v259 = v248
		v260 = v249
		goto L76
	} else {
		goto L81
	}
L80:
	;
	v259 = v248
	v260 = v249
	goto L76
L81:
	;
	v252 = int32(1)
	if v248 == v249 {
		v244 = v244 + v252
		v245 = v245 + v252
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v265 = F_substitute_path_macro(m, v120, int32(284401), v25)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v120
	v269 = F_psprintf(m, int32(267173), v11)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	v271 = v265
	goto L74
L87:
	;
	v271 = v269
	goto L74
L88:
	;
	F_canonicalize_path_enc(m, v271)
	mBase = m.M
	v276 = F_lappend(m, v40, v271)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v117))))
	if v279 != 0 {
		v37 = v37 + v119
		v40 = v276
		goto L10
	} else {
		goto L90
	}
L90:
	;
	goto L11
}
func F_parse_extension_control_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
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
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
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
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v904 int32
	_ = v904
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(240)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+236)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v14)+232)) = v3
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L12
	} else {
		goto L287
	}
L2:
	;
	if v192&int32(3) == int32(0) {
		v224 = v192
		goto L58
	} else {
		goto L59
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v56
	v77 = F_psprintf(m, int32(296353), v14+int32(208))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L26
	}
L4:
	;
	if v69 == int32(0) {
		goto L1
	} else {
		goto L24
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v57 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L8:
	;
	v42 = F_palloc(m, int32(1024))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L19
	}
L9:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_pstrdup(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(47) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return
L13:
	;
	v40 = v24
	goto L8
L14:
	;
	v29 = F_pstrdup(m, v20)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = v31
	v37 = F_psprintf(m, int32(174083), v14+int32(192))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v40 = v29
	goto L8
L18:
	;
	v40 = v37
	goto L8
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+184)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v14)+180)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = v40
	v52 = F_pg_snprintf(m, v42, int32(1024), int32(296346), v14+int32(176))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_pfree(m, v40)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v69 = v42
	goto L4
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v57
	v65 = F_psprintf(m, int32(296332), v14+int32(224))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	v69 = v65
	goto L4
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v192 = v72
	v195 = v69
	goto L2
L25:
	;
	if v132 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L26:
	;
	v79 = F_get_extension_control_directories(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v81 = int32(0)
	v82 = m.G0
	v84 = v82 - int32(32)
	m.G0 = v84
	if v79 == v81 {
		v132 = v81
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L44
	}
L29:
	;
	m.G0 = v84 + int32(32)
	goto L25
L30:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v88 <= int32(0) {
		v132 = v81
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v99 = v3
	goto L32
L32:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v99<<(uint(int32(2))%32))))
	v107 = F_pstrdup(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L34
	}
L33:
	;
	v132 = v81
	goto L29
L34:
	;
	F_canonicalize_path_enc(m, v107)
	mBase = m.M
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v110 != int32(47) {
		goto L28
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = v107
	v116 = F_psprintf(m, int32(174083), v84)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	v118 = F_pg_file_exists(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	if v118 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = v116
	goto L29
L39:
	;
	goto L40
L40:
	;
	F_pfree(m, v107)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v116)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	v125 = v99 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v125 < v126 {
		v99 = v125
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(316023)
	F_errmsg(m, int32(316268), v84+int32(16))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(487544), int32(4047), int32(152319))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	v166 = F_strlen(m, v132)
	mBase = m.M
	v173 = v166 + int32(1)
	goto L51
L49:
	;
	v187 = F_pnstrdup(m, v132, v185-v132)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L55
	}
L50:
	;
	goto L49
L51:
	;
	v175 = int32(0)
	if v173 == v175 {
		v185 = v175
		goto L50
	} else {
		goto L53
	}
L52:
	;
	v185 = v180
	goto L50
L53:
	;
	v179 = v173 - int32(1)
	v180 = v132 + v179
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v181 != int32(47) {
		v173 = v179
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v187
	v192 = v187
	v195 = v132
	goto L2
L56:
	;
	v260 = F_pnstrdup(m, v192, v257-int32(10))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L12
	} else {
		goto L73
	}
L57:
	;
	v257 = v249 - v192
	goto L56
L58:
	;
	v228 = v224
	goto L67
L59:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v208 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v257 = int32(0)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v213 = v192
	goto L63
L63:
	;
	v217 = v213 + int32(1)
	if v217&int32(3) == int32(0) {
		v224 = v217
		goto L58
	} else {
		goto L65
	}
L64:
	;
	v249 = v217
	goto L57
L65:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	if v222 != 0 {
		v213 = v217
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v237 = int32(-2139062144)
	if (int32(16843008)-v234|v234)&v237 == v237 {
		v228 = v228 + int32(4)
		goto L67
	} else {
		goto L69
	}
L68:
	;
	v243 = v228
	goto L70
L69:
	;
	goto L68
L70:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v247 != 0 {
		v243 = v243 + int32(1)
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v249 = v243
	goto L57
L72:
	;
	goto L71
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v260
	v264 = F_AllocateFile(m, v195, int32(227078))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L12
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v195)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L12
	} else {
		goto L286
	}
L75:
	;
	if v264 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	if l1 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L78
L78:
	;
	v295 = F_ParseConfigFp(m, v264, v195, int32(0), int32(21), v14+int32(236), v14+int32(232))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L12
	} else {
		goto L87
	}
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v269 == int32(44) {
		goto L74
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v195
	F_errmsg(m, int32(293691), v14+int32(16))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L12
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(487544), int32(696), int32(379680))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L12
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v297 = F_FreeFile(m, v264)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L12
	} else {
		goto L88
	}
L88:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	if v299 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v312 = v299
	goto L92
L90:
	;
	v867 = int32(0)
	goto L91
L91:
	;
	F_FreeConfigVariables(m, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L12
	} else {
		goto L279
	}
L92:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	v322 = int32(13485)
	v325 = int32(*(*uint8)(unsafe.Add(mBase, _consts[431])))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v326 == int32(0) {
		v345 = v325
		v346 = v326
		goto L98
	} else {
		goto L99
	}
L93:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v14)+236))
	v867 = v854
	goto L91
L94:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v312)+24))
	if v853 != 0 {
		v312 = v853
		goto L92
	} else {
		goto L278
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L12
	} else {
		goto L274
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L12
	} else {
		goto L270
	}
L97:
	;
	if v346-v345 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L98:
	;
	goto L97
L99:
	;
	if v325 != v326 {
		v345 = v325
		v346 = v326
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v330 = v321
	v331 = v322
	goto L101
L101:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+1)))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+1)))
	if v335 == int32(0) {
		v345 = v334
		v346 = v335
		goto L98
	} else {
		goto L103
	}
L102:
	;
	v345 = v334
	v346 = v335
	goto L98
L103:
	;
	v338 = int32(1)
	if v334 == v335 {
		v330 = v330 + v338
		v331 = v331 + v338
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	if l1 != 0 {
		goto L96
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v354 = int32(266768)
	v357 = int32(*(*uint8)(unsafe.Add(mBase, _consts[432])))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v358 == int32(0) {
		v377 = v357
		v378 = v358
		goto L111
	} else {
		goto L112
	}
L108:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v351 = F_pstrdup(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L12
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v351
	goto L94
L110:
	;
	if v378-v377 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L111:
	;
	goto L110
L112:
	;
	if v357 != v358 {
		v377 = v357
		v378 = v358
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v362 = v321
	v363 = v354
	goto L114
L114:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+1)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+1)))
	if v367 == int32(0) {
		v377 = v366
		v378 = v367
		goto L111
	} else {
		goto L116
	}
L115:
	;
	v377 = v366
	v378 = v367
	goto L111
L116:
	;
	v370 = int32(1)
	if v366 == v367 {
		v362 = v362 + v370
		v363 = v363 + v370
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if l1 != 0 {
		goto L95
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v386 = int32(371071)
	v389 = int32(*(*uint8)(unsafe.Add(mBase, _consts[433])))
	v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v390 == int32(0) {
		v409 = v389
		v410 = v390
		goto L124
	} else {
		goto L125
	}
L121:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v383 = F_pstrdup(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L12
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v383
	goto L94
L123:
	;
	if v410-v409 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L124:
	;
	goto L123
L125:
	;
	if v389 != v390 {
		v409 = v389
		v410 = v390
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v394 = v321
	v395 = v386
	goto L127
L127:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+1)))
	if v399 == int32(0) {
		v409 = v398
		v410 = v399
		goto L124
	} else {
		goto L129
	}
L128:
	;
	v409 = v398
	v410 = v399
	goto L124
L129:
	;
	v402 = int32(1)
	if v398 == v399 {
		v394 = v394 + v402
		v395 = v395 + v402
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v415 = F_pstrdup(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L12
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v418 = int32(93526)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, _consts[434])))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v422 == int32(0) {
		v441 = v421
		v442 = v422
		goto L136
	} else {
		goto L137
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v415
	goto L94
L135:
	;
	if v442-v441 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L136:
	;
	goto L135
L137:
	;
	if v421 != v422 {
		v441 = v421
		v442 = v422
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v426 = v321
	v427 = v418
	goto L139
L139:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	if v431 == int32(0) {
		v441 = v430
		v442 = v431
		goto L136
	} else {
		goto L141
	}
L140:
	;
	v441 = v430
	v442 = v431
	goto L136
L141:
	;
	v434 = int32(1)
	if v430 == v431 {
		v426 = v426 + v434
		v427 = v427 + v434
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v447 = F_pstrdup(m, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L12
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v450 = int32(497384)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, _consts[435])))
	v454 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v454 == int32(0) {
		v473 = v453
		v474 = v454
		goto L148
	} else {
		goto L149
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v447
	goto L94
L147:
	;
	if v474-v473 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L148:
	;
	goto L147
L149:
	;
	if v453 != v454 {
		v473 = v453
		v474 = v454
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v458 = v321
	v459 = v450
	goto L151
L151:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+1)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458)+1)))
	if v463 == int32(0) {
		v473 = v462
		v474 = v463
		goto L148
	} else {
		goto L153
	}
L152:
	;
	v473 = v462
	v474 = v463
	goto L148
L153:
	;
	v466 = int32(1)
	if v462 == v463 {
		v458 = v458 + v466
		v459 = v459 + v466
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v479 = F_pstrdup(m, v478)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L12
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v482 = int32(385186)
	v485 = int32(*(*uint8)(unsafe.Add(mBase, _consts[436])))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v486 == int32(0) {
		v505 = v485
		v506 = v486
		goto L160
	} else {
		goto L161
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v479
	goto L94
L159:
	;
	if v506-v505 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L160:
	;
	goto L159
L161:
	;
	if v485 != v486 {
		v505 = v485
		v506 = v486
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v490 = v321
	v491 = v482
	goto L163
L163:
	;
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v491)+1)))
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490)+1)))
	if v495 == int32(0) {
		v505 = v494
		v506 = v495
		goto L160
	} else {
		goto L165
	}
L164:
	;
	v505 = v494
	v506 = v495
	goto L160
L165:
	;
	v498 = int32(1)
	if v494 == v495 {
		v490 = v490 + v498
		v491 = v491 + v498
		goto L163
	} else {
		goto L166
	}
L166:
	;
	goto L164
L167:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v511 = F_strlen(m, v510)
	mBase = m.M
	v512 = F_parse_bool_with_len(m, v510, v511, l0+int32(32))
	mBase = m.M
	goto L170
L168:
	;
	goto L169
L169:
	;
	v532 = int32(214174)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, _consts[437])))
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v536 == int32(0) {
		v555 = v535
		v556 = v536
		goto L177
	} else {
		goto L178
	}
L170:
	;
	if v512 != 0 {
		goto L94
	} else {
		goto L171
	}
L171:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L12
	} else {
		goto L172
	}
L172:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L12
	} else {
		goto L173
	}
L173:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v520
	F_errmsg(m, int32(339685), v14-int32(-64))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L12
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(487544), int32(751), int32(379680))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L12
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L176:
	;
	if v556-v555 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L177:
	;
	goto L176
L178:
	;
	if v535 != v536 {
		v555 = v535
		v556 = v536
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v540 = v321
	v541 = v532
	goto L180
L180:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541)+1)))
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v540)+1)))
	if v545 == int32(0) {
		v555 = v544
		v556 = v545
		goto L177
	} else {
		goto L182
	}
L181:
	;
	v555 = v544
	v556 = v545
	goto L177
L182:
	;
	v548 = int32(1)
	if v544 == v545 {
		v540 = v540 + v548
		v541 = v541 + v548
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v561 = F_strlen(m, v560)
	mBase = m.M
	v562 = F_parse_bool_with_len(m, v560, v561, l0+int32(33))
	mBase = m.M
	goto L187
L185:
	;
	goto L186
L186:
	;
	v582 = int32(433392)
	v585 = int32(*(*uint8)(unsafe.Add(mBase, _consts[438])))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v586 == int32(0) {
		v605 = v585
		v606 = v586
		goto L194
	} else {
		goto L195
	}
L187:
	;
	if v562 != 0 {
		goto L94
	} else {
		goto L188
	}
L188:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L12
	} else {
		goto L189
	}
L189:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L12
	} else {
		goto L190
	}
L190:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v570
	F_errmsg(m, int32(339685), v14+int32(80))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L12
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(487544), int32(759), int32(379680))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L12
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	if v606-v605 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L194:
	;
	goto L193
L195:
	;
	if v585 != v586 {
		v605 = v585
		v606 = v586
		goto L194
	} else {
		goto L196
	}
L196:
	;
	v590 = v321
	v591 = v582
	goto L197
L197:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v591)+1)))
	v595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	if v595 == int32(0) {
		v605 = v594
		v606 = v595
		goto L194
	} else {
		goto L199
	}
L198:
	;
	v605 = v594
	v606 = v595
	goto L194
L199:
	;
	v598 = int32(1)
	if v594 == v595 {
		v590 = v590 + v598
		v591 = v591 + v598
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v611 = F_strlen(m, v610)
	mBase = m.M
	v612 = F_parse_bool_with_len(m, v610, v611, l0+int32(34))
	mBase = m.M
	goto L204
L202:
	;
	goto L203
L203:
	;
	v632 = int32(330076)
	v635 = int32(*(*uint8)(unsafe.Add(mBase, _consts[408])))
	v636 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v636 == int32(0) {
		v655 = v635
		v656 = v636
		goto L211
	} else {
		goto L212
	}
L204:
	;
	if v612 != 0 {
		goto L94
	} else {
		goto L205
	}
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L12
	} else {
		goto L206
	}
L206:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L12
	} else {
		goto L207
	}
L207:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v620
	F_errmsg(m, int32(339685), v14+int32(96))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L12
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(487544), int32(767), int32(379680))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L12
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	if v656-v655 == int32(0) {
		goto L218
	} else {
		goto L219
	}
L211:
	;
	goto L210
L212:
	;
	if v635 != v636 {
		v655 = v635
		v656 = v636
		goto L211
	} else {
		goto L213
	}
L213:
	;
	v640 = v321
	v641 = v632
	goto L214
L214:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+1)))
	if v645 == int32(0) {
		v655 = v644
		v656 = v645
		goto L211
	} else {
		goto L216
	}
L215:
	;
	v655 = v644
	v656 = v645
	goto L211
L216:
	;
	v648 = int32(1)
	if v644 == v645 {
		v640 = v640 + v648
		v641 = v641 + v648
		goto L214
	} else {
		goto L217
	}
L217:
	;
	goto L215
L218:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v662 = F_pg_char_to_encoding_private(m, v660)
	mBase = m.M
	if base.Ui32(int32(35)) <= base.Ui32(v662) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	goto L220
L220:
	;
	v688 = int32(158894)
	v691 = int32(*(*uint8)(unsafe.Add(mBase, _consts[439])))
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v692 == int32(0) {
		v711 = v691
		v712 = v692
		goto L231
	} else {
		goto L232
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v665
	if int32(0) <= v665 {
		goto L94
	} else {
		goto L225
	}
L222:
	;
	v665 = int32(-1)
	goto L224
L223:
	;
	v665 = v662
	goto L224
L224:
	;
	goto L221
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L12
	} else {
		goto L226
	}
L226:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L12
	} else {
		goto L227
	}
L227:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v676
	F_errmsg(m, int32(374910), v14+int32(112))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L12
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(487544), int32(776), int32(379680))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L12
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	if v712-v711 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L231:
	;
	goto L230
L232:
	;
	if v691 != v692 {
		v711 = v691
		v712 = v692
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v696 = v321
	v697 = v688
	goto L234
L234:
	;
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v697)+1)))
	v701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v696)+1)))
	if v701 == int32(0) {
		v711 = v700
		v712 = v701
		goto L231
	} else {
		goto L236
	}
L235:
	;
	v711 = v700
	v712 = v701
	goto L231
L236:
	;
	v704 = int32(1)
	if v700 == v701 {
		v696 = v696 + v704
		v697 = v697 + v704
		goto L234
	} else {
		goto L237
	}
L237:
	;
	goto L235
L238:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v717 = F_pstrdup(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L12
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v741 = int32(351075)
	v744 = int32(*(*uint8)(unsafe.Add(mBase, _consts[440])))
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v321))))
	if v745 == int32(0) {
		v764 = v744
		v765 = v745
		goto L249
	} else {
		goto L250
	}
L241:
	;
	v720 = F_SplitIdentifierString(m, v717, int32(44), l0+int32(40))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L12
	} else {
		goto L242
	}
L242:
	;
	if v720 != 0 {
		goto L94
	} else {
		goto L243
	}
L243:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L12
	} else {
		goto L244
	}
L244:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L12
	} else {
		goto L245
	}
L245:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v729
	F_errmsg(m, int32(160615), v14+int32(128))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L12
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(487544), int32(790), int32(379680))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L12
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	if v765-v764 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L249:
	;
	goto L248
L250:
	;
	if v744 != v745 {
		v764 = v744
		v765 = v745
		goto L249
	} else {
		goto L251
	}
L251:
	;
	v749 = v321
	v750 = v741
	goto L252
L252:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v750)+1)))
	v754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749)+1)))
	if v754 == int32(0) {
		v764 = v753
		v765 = v754
		goto L249
	} else {
		goto L254
	}
L253:
	;
	v764 = v753
	v765 = v754
	goto L249
L254:
	;
	v757 = int32(1)
	if v753 == v754 {
		v749 = v749 + v757
		v750 = v750 + v757
		goto L252
	} else {
		goto L255
	}
L255:
	;
	goto L253
L256:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v312)+4))
	v770 = F_pstrdup(m, v769)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L12
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L12
	} else {
		goto L266
	}
L259:
	;
	v773 = F_SplitIdentifierString(m, v770, int32(44), l0+int32(44))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L12
	} else {
		goto L260
	}
L260:
	;
	if v773 != 0 {
		goto L94
	} else {
		goto L261
	}
L261:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L12
	} else {
		goto L262
	}
L262:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L12
	} else {
		goto L263
	}
L263:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = v782
	F_errmsg(m, int32(160615), v14+int32(144))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L12
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(487544), int32(805), int32(379680))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L12
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L12
	} else {
		goto L267
	}
L267:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v195
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v801
	F_errmsg(m, int32(690607), v14+int32(160))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L12
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(487544), int32(812), int32(379680))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L12
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L12
	} else {
		goto L271
	}
L271:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v821
	F_errmsg(m, int32(381063), v14+int32(32))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L12
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(487544), int32(719), int32(379680))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L12
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L12
	} else {
		goto L275
	}
L275:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v840
	F_errmsg(m, int32(381063), v14+int32(48))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L12
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(487544), int32(729), int32(379680))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L12
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	goto L93
L279:
	;
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v870 != int32(1) {
		goto L74
	} else {
		goto L280
	}
L280:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v873 == int32(0) {
		goto L74
	} else {
		goto L281
	}
L281:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L12
	} else {
		goto L282
	}
L282:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L12
	} else {
		goto L283
	}
L283:
	;
	F_errmsg(m, int32(338105), int32(0))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L12
	} else {
		goto L284
	}
L284:
	;
	F_errfinish(m, int32(487544), int32(820), int32(379680))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L12
	} else {
		goto L285
	}
L285:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L286:
	;
	m.G0 = v14 + int32(240)
	return
L287:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L12
	} else {
		goto L288
	}
L288:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v926
	F_errmsg(m, int32(388752), v14)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L12
	} else {
		goto L289
	}
L289:
	;
	F_errhint(m, int32(601461), int32(0))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L12
	} else {
		goto L290
	}
L290:
	;
	F_errfinish(m, int32(487544), int32(673), int32(379680))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L12
	} else {
		goto L291
	}
L291:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
