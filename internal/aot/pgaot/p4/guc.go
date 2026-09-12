package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GUC_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_emscripten_builtin_malloc(m, int32(4))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
		if v8 == int32(0) {
			F_GUC_flex_fatal(m, int32(684192))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
			return
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v17-int32(1)) <= base.Ui32(v16) {
			v22 = v17 + int32(8)
			v25 = F_emscripten_builtin_realloc(m, v4, v22<<(uint(int32(2))%32))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v25
			if v25 == int32(0) {
				F_GUC_flex_fatal(m, int32(684192))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v32 = v25 + v29<<(uint(int32(2))%32)
				v33 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v22
				return
			}
		} else {
			return
		}
	}
}
func F_TransformGUCArray(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
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
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	v4 = int32(0)
	v11 = m.G0
	v12 = int32(16)
	v13 = v11 - v12
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	v22 = l0 + v12
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v4 < v23 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v40 = F_array_ref(m, l0, v13+int32(12), v13+int32(11))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if v42 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v45 = F_text_to_cstring(m, v40)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v303 = v301 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v303
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v303 <= v305 {
		goto L4
	} else {
		goto L89
	}
L11:
	;
	v248 = v243
	goto L72
L12:
	;
	v47 = int32(546484)
	v51 = m.G0
	v53 = v51 - int32(32)
	m.G0 = v53
	v55 = int32(*(*int8)(unsafe.Add(mBase, _consts[1446])))
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v45))))
	if v115 == int32(61) {
		goto L33
	} else {
		goto L34
	}
L14:
	;
	m.G0 = v53 + int32(32)
	v113 = v108 - v45
	goto L13
L15:
	;
	v60 = F___memset(m, v53, int32(0), int32(32))
	mBase = m.M
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1446])))
	if v61 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1447])))
	if v56 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = F___strchrnul(m, v45, v55)
	mBase = m.M
	v108 = v57
	goto L14
L19:
	;
	goto L18
L20:
	;
	v63 = v47
	v64 = v61
	goto L23
L21:
	;
	goto L22
L22:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v85 == int32(0) {
		v108 = v45
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v71 = v53 + int32(base.Ui32(v64)>>(uint(int32(3))%32))&int32(28)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 | v73<<(uint(v64)%32)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)))
	if v77 != 0 {
		v63 = v63 + v73
		v64 = v77
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L22
L25:
	;
	goto L24
L26:
	;
	v89 = v45
	v90 = v85
	goto L27
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(base.Ui32(v90)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v98)>>(uint(v90)%32))&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v108 = v104
	goto L14
L29:
	;
	v108 = v89
	goto L14
L30:
	;
	goto L31
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+1)))
	v104 = v89 + int32(1)
	if v102 != 0 {
		v89 = v104
		v90 = v102
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v119 = v113 + int32(1)
	v120 = F_palloc(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v240 = F_pstrdup(m, v45)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L70
	}
L36:
	;
	if v119 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v237 = F_pstrdup(m, v119+v45)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L69
	}
L38:
	;
	v233 = F_strlen(m, v229)
	mBase = m.M
	goto L37
L39:
	;
	v229 = v45
	goto L38
L40:
	;
	goto L41
L41:
	;
	v127 = v119 - int32(1)
	if (v120^v45)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v226)
	v229 = v222
	goto L38
L43:
	;
	v207 = v202
	v208 = v203
	v209 = v204
	goto L65
L44:
	;
	if v197 == int32(0) {
		v222 = v195
		v223 = v196
		goto L42
	} else {
		goto L64
	}
L45:
	;
	v195 = v45
	v196 = v120
	v197 = v127
	goto L44
L46:
	;
	goto L47
L47:
	;
	v131 = int32(0)
	if v45&int32(3) == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v164 == int32(0) {
		v222 = v161
		v223 = v162
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v161 = v45
	v162 = v120
	v163 = v127
	v164 = base.B2i32(v127 != v131)
	goto L48
L50:
	;
	if v127 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v140 = v45
	v141 = v120
	v142 = v127
	goto L52
L52:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	*(*uint8)(unsafe.Add(mBase, uint32(v141))) = uint8(v144)
	if v144 == int32(0) {
		v202 = v140
		v203 = v141
		v204 = v142
		goto L43
	} else {
		goto L54
	}
L53:
	;
	v161 = v155
	v162 = v149
	v163 = v151
	v164 = v153
	goto L48
L54:
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
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v151 != 0 {
		v140 = v155
		v141 = v149
		v142 = v151
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	if v167 == int32(0) {
		v195 = v161
		v196 = v162
		v197 = v163
		goto L44
	} else {
		goto L58
	}
L58:
	;
	if base.Ui32(v163) < base.Ui32(int32(4)) {
		v195 = v161
		v196 = v162
		v197 = v163
		goto L44
	} else {
		goto L59
	}
L59:
	;
	v173 = v161
	v174 = v162
	v175 = v163
	goto L60
L60:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v181 = int32(-2139062144)
	if (int32(16843008)-v178|v178)&v181 != v181 {
		v202 = v173
		v203 = v174
		v204 = v175
		goto L43
	} else {
		goto L62
	}
L61:
	;
	v195 = v189
	v196 = v187
	v197 = v191
	goto L44
L62:
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
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v202 = v195
	v203 = v196
	v204 = v197
	goto L43
L65:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v211)
	if v211 == int32(0) {
		v222 = v207
		v223 = v208
		goto L42
	} else {
		goto L67
	}
L66:
	;
	v222 = v218
	v223 = v216
	goto L42
L67:
	;
	v215 = int32(1)
	v216 = v208 + v215
	v218 = v207 + v215
	v220 = v209 - v215
	if v220 != 0 {
		v207 = v218
		v208 = v216
		v209 = v220
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v243 = v120
	v244 = v237
	goto L11
L70:
	;
	v243 = v240
	v244 = int32(0)
	goto L11
L71:
	;
	F_pfree(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L88
	}
L72:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v255 != int32(45) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v281 = F_lappend(m, v280, v243)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L6
	} else {
		goto L86
	}
L74:
	;
	goto L73
L75:
	;
	v248 = v248 + int32(1)
	goto L72
L76:
	;
	if v255 != 0 {
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v276 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v248))) = uint8(v276)
	goto L75
L79:
	;
	if v244 != 0 {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v260 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	if v260 == int32(0) {
		v288 = v243
		goto L71
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v243
	F_errmsg(m, int32(700806), v13)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(499912), int32(6444), int32(26324))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	v288 = v243
	goto L71
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v281
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v285 = F_lappend(m, v284, v244)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v285
	v288 = v45
	goto L71
L88:
	;
	goto L10
L89:
	;
	goto L5
}
func F_build_guc_variables(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	v1 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v17 = F_AllocSetContextCreateInternal(m, v12, int32(61985), v1, int32(8192), int32(8388608))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1436])) = v17
	v21 = *(*int32)(unsafe.Add(mBase, _consts[1437]))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v1
	v25 = int32(4122480)
	goto L6
L4:
	;
	v37 = v1
	goto L5
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = int32(0)
	v31 = v23 + int32(1)
	v33 = v31 * int32(120)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[1437])))
	if v36 != 0 {
		v23 = v31
		v25 = v33 + int32(4122480)
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v37 = v31
	goto L5
L8:
	;
	goto L7
L9:
	;
	v45 = v37
	v46 = v1
	v47 = int32(4136400)
	goto L12
L10:
	;
	v61 = v37
	goto L11
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	if v68 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v50 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v50
	v53 = v45 + v50
	v55 = v46 + v50
	v57 = v55 << (uint(int32(7)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_consts[1438])))
	if v60 != 0 {
		v45 = v53
		v46 = v55
		v47 = v57 + int32(4136400)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v61 = v53
	goto L11
L14:
	;
	goto L13
L15:
	;
	v70 = v61
	v71 = int32(0)
	v72 = int32(4155344)
	goto L18
L16:
	;
	v86 = v61
	goto L17
L17:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[1440]))
	if v93 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = int32(2)
	v77 = int32(1)
	v78 = v70 + v77
	v80 = v71 + v77
	v82 = v80 * int32(152)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)+uint32(_consts[1439])))
	if v85 != 0 {
		v70 = v78
		v71 = v80
		v72 = v82 + int32(4155344)
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v86 = v78
	goto L17
L20:
	;
	goto L19
L21:
	;
	v95 = v86
	v96 = int32(0)
	v97 = int32(4159456)
	goto L24
L22:
	;
	v111 = v86
	goto L23
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _consts[1441]))
	if v118 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = int32(3)
	v102 = int32(1)
	v103 = v95 + v102
	v105 = v96 + v102
	v107 = v105 * int32(120)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1440])))
	if v110 != 0 {
		v95 = v103
		v96 = v105
		v97 = v107 + int32(4159456)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v111 = v103
	goto L23
L26:
	;
	goto L25
L27:
	;
	v120 = v111
	v121 = int32(0)
	v122 = int32(4168592)
	goto L30
L28:
	;
	v136 = v111
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(1657)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1658)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(34359738372)
	v152 = base.I32_div_s(v136, int32(4))
	v157 = F_hash_create(m, int32(393987), v152+v136, v6+int32(-48), int32(1224))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v122)+24)) = int32(4)
	v127 = int32(1)
	v128 = v120 + v127
	v130 = v121 + v127
	v132 = v130 * int32(124)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[1441])))
	if v135 != 0 {
		v120 = v128
		v121 = v130
		v122 = v132 + int32(4168592)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v136 = v128
	goto L29
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[431])) = v157
	v161 = *(*int32)(unsafe.Add(mBase, _consts[1437]))
	if v161 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v163 = int32(4122480)
	v164 = int32(0)
	goto L37
L35:
	;
	goto L36
L36:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _consts[1438]))
	if v190 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	v173 = F_hash_search(m, v169, v163, int32(1), v6+int32(-49))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v163
	v177 = v164 + int32(1)
	v179 = v177 * int32(120)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v179)+uint32(_consts[1437])))
	if v182 != 0 {
		v163 = v179 + int32(4122480)
		v164 = v177
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v192 = int32(4136400)
	v193 = int32(0)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	if v219 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	v202 = F_hash_search(m, v198, v192, int32(1), v6+int32(-49))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v192
	v206 = v193 + int32(1)
	v208 = v206 << (uint(int32(7)) % 32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v208)+uint32(_consts[1438])))
	if v211 != 0 {
		v192 = v208 + int32(4136400)
		v193 = v206
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v221 = int32(4155344)
	v222 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[1440]))
	if v248 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	v231 = F_hash_search(m, v227, v221, int32(1), v6+int32(-49))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v221
	v235 = v222 + int32(1)
	v237 = v235 * int32(152)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+uint32(_consts[1439])))
	if v240 != 0 {
		v221 = v237 + int32(4155344)
		v222 = v235
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v250 = int32(4159456)
	v251 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _consts[1441]))
	if v277 != 0 {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	v260 = F_hash_search(m, v256, v250, int32(1), v6+int32(-49))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+4)) = v250
	v264 = v251 + int32(1)
	v266 = v264 * int32(120)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)+uint32(_consts[1440])))
	if v269 != 0 {
		v250 = v266 + int32(4159456)
		v251 = v264
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v279 = int32(4168592)
	v280 = int32(0)
	goto L65
L63:
	;
	goto L64
L64:
	;
	m.G0 = v8 - int32(-64)
	return
L65:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	v289 = F_hash_search(m, v285, v279, int32(1), v6+int32(-49))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v289)+4)) = v279
	v293 = v280 + int32(1)
	v295 = v293 * int32(124)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v295)+uint32(_consts[1441])))
	if v298 != 0 {
		v279 = v295 + int32(4168592)
		v280 = v293
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
}
func F_convert_GUC_name_for_parameter_acl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	v2 = int32(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == v2 {
		v140 = int32(291521)
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v346 = F_pstrdup(m, v340)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L121
	} else {
		goto L122
	}
L2:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v334<<(uint(int32(2))%32))+uint32(_consts[1442])))
	v340 = v339
	goto L1
L3:
	;
	v155 = int32(291530)
	v157 = l0
	goto L66
L4:
	;
	if v149 != 0 {
		goto L3
	} else {
		goto L63
	}
L5:
	;
	v149 = v144 + base.I32_extend8_s(v143)
	goto L4
L6:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v141 != 0 {
		goto L60
	} else {
		goto L61
	}
L7:
	;
	if base.Ui32((v10-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v21 = v10 | int32(32)
	goto L10
L9:
	;
	v21 = v10
	goto L10
L10:
	;
	if v21 != int32(115) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v143 = v21
	v144 = int32(-115)
	goto L5
L12:
	;
	goto L13
L13:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v26 == int32(0) {
		v140 = int32(291522)
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32((v26-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v37 = v26 | int32(32)
	goto L17
L16:
	;
	v37 = v26
	goto L17
L17:
	;
	if v37 != int32(111) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v143 = v37
	v144 = int32(-111)
	goto L5
L19:
	;
	goto L20
L20:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v42 == int32(0) {
		v140 = int32(291523)
		goto L6
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32((v42-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v53 = v42 | int32(32)
	goto L24
L23:
	;
	v53 = v42
	goto L24
L24:
	;
	if v53 != int32(114) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v143 = v53
	v144 = int32(-114)
	goto L5
L26:
	;
	goto L27
L27:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v58 == int32(0) {
		v140 = int32(291524)
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v69 = v58 | int32(32)
	goto L31
L30:
	;
	v69 = v58
	goto L31
L31:
	;
	if v69 != int32(116) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v143 = v69
	v144 = int32(-116)
	goto L5
L33:
	;
	goto L34
L34:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v74 == int32(0) {
		v140 = int32(291525)
		goto L6
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v85 = v74 | int32(32)
	goto L38
L37:
	;
	v85 = v74
	goto L38
L38:
	;
	if v85 != int32(95) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v143 = v85
	v144 = int32(-95)
	goto L5
L40:
	;
	goto L41
L41:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v90 == int32(0) {
		v140 = int32(291526)
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v93 = int32(-109)
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v102 = v90 | int32(32)
	goto L45
L44:
	;
	v102 = v90
	goto L45
L45:
	;
	if v102 != int32(109) {
		v143 = v102
		v144 = v93
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v106 == int32(0) {
		v140 = int32(291527)
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if base.Ui32((v106-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v117 = v106 | int32(32)
	goto L50
L49:
	;
	v117 = v106
	goto L50
L50:
	;
	if v117 != int32(101) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v143 = v117
	v144 = int32(-101)
	goto L5
L52:
	;
	goto L53
L53:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	if v122 == int32(0) {
		v140 = int32(291528)
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if base.Ui32((v122-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v133 = v122 | int32(32)
	goto L57
L56:
	;
	v133 = v122
	goto L57
L57:
	;
	if v133 != int32(109) {
		v143 = v133
		v144 = v93
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v136 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v140 = int32(291529)
	goto L6
L60:
	;
	v142 = int32(-1)
	goto L62
L61:
	;
	v142 = v2
	goto L62
L62:
	;
	v149 = v142
	goto L4
L63:
	;
	v334 = int32(1)
	goto L2
L64:
	;
	v245 = int32(342573)
	v247 = l0
	goto L94
L65:
	;
	if v236 != 0 {
		goto L64
	} else {
		goto L92
	}
L66:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v162 != 0 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v236 = base.I32_extend8_s(v226) - base.I32_extend8_s(v224)
	goto L65
L68:
	;
	goto L67
L69:
	;
	v197 = int32(2)
	if base.Ui32((v187-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L85
	} else {
		goto L86
	}
L70:
	;
	if v155 == int32(291540) {
		goto L64
	} else {
		goto L73
	}
L71:
	;
	v192 = v155
	goto L72
L72:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v193 != 0 {
		goto L82
	} else {
		goto L83
	}
L73:
	;
	if base.Ui32((v162-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v173 = v162 | int32(32)
	goto L76
L75:
	;
	v173 = v162
	goto L76
L76:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if base.Ui32((v174-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v183 = v174 | int32(32)
	goto L79
L78:
	;
	v183 = v174
	goto L79
L79:
	;
	if v173 != v183&int32(255) {
		v224 = v183
		v226 = v173
		goto L68
	} else {
		goto L80
	}
L80:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	if v187 != 0 {
		goto L69
	} else {
		goto L81
	}
L81:
	;
	v192 = v155 + int32(1)
	goto L72
L82:
	;
	v194 = int32(-1)
	goto L84
L83:
	;
	v194 = int32(0)
	goto L84
L84:
	;
	v236 = v194
	goto L65
L85:
	;
	v209 = v187 | int32(32)
	goto L87
L86:
	;
	v209 = v187
	goto L87
L87:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+int32(1)))))
	if base.Ui32((v210-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v219 = v210 | int32(32)
	goto L90
L89:
	;
	v219 = v210
	goto L90
L90:
	;
	if v209 == v219&int32(255) {
		v155 = v155 + v197
		v157 = v157 + v197
		goto L66
	} else {
		goto L91
	}
L91:
	;
	v224 = v219
	v226 = v209
	goto L68
L92:
	;
	v334 = int32(3)
	goto L2
L93:
	;
	if v326 != 0 {
		v340 = l0
		goto L1
	} else {
		goto L120
	}
L94:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v252 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v326 = base.I32_extend8_s(v316) - base.I32_extend8_s(v314)
	goto L93
L96:
	;
	goto L95
L97:
	;
	v287 = int32(2)
	if base.Ui32((v277-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L113
	} else {
		goto L114
	}
L98:
	;
	if v245 == int32(342587) {
		v340 = l0
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v282 = v245
	goto L100
L100:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	if v283 != 0 {
		goto L110
	} else {
		goto L111
	}
L101:
	;
	if base.Ui32((v252-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v263 = v252 | int32(32)
	goto L104
L103:
	;
	v263 = v252
	goto L104
L104:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if base.Ui32((v264-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v273 = v264 | int32(32)
	goto L107
L106:
	;
	v273 = v264
	goto L107
L107:
	;
	if v263 != v273&int32(255) {
		v314 = v273
		v316 = v263
		goto L96
	} else {
		goto L108
	}
L108:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	if v277 != 0 {
		goto L97
	} else {
		goto L109
	}
L109:
	;
	v282 = v245 + int32(1)
	goto L100
L110:
	;
	v284 = int32(-1)
	goto L112
L111:
	;
	v284 = int32(0)
	goto L112
L112:
	;
	v326 = v284
	goto L93
L113:
	;
	v299 = v277 | int32(32)
	goto L115
L114:
	;
	v299 = v277
	goto L115
L115:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+int32(1)))))
	if base.Ui32((v300-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v309 = v300 | int32(32)
	goto L118
L117:
	;
	v309 = v300
	goto L118
L118:
	;
	if v299 == v309&int32(255) {
		v245 = v245 + v287
		v247 = v247 + v287
		goto L94
	} else {
		goto L119
	}
L119:
	;
	v314 = v309
	v316 = v299
	goto L96
L120:
	;
	v334 = int32(5)
	goto L2
L121:
	;
	return int32(0)
L122:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346))))
	if v350 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v352 = v350
	v354 = v346
	goto L126
L124:
	;
	goto L125
L125:
	;
	return v346
L126:
	;
	if base.Ui32((v352-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	goto L125
L128:
	;
	v364 = v352 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v354))) = uint8(v364)
	goto L130
L129:
	;
	goto L130
L130:
	;
	v367 = v354 + int32(1)
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367))))
	if v368 != 0 {
		v352 = v368
		v354 = v367
		goto L126
	} else {
		goto L131
	}
L131:
	;
	goto L127
}
func F_guc_name_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 != 0 {
		v7 = v6
		v8 = v3
		v9 = v5
		for {
			if base.Ui32((v7-int32(65))&int32(255)) < base.Ui32(int32(26)) {
				v18 = v7 | int32(32)
			} else {
				v18 = v7
			}
			v22 = base.I32_extend8_s(v18) ^ base.I32_rotl(v8, int32(5))
			v24 = v9 + int32(1)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v25 != 0 {
				v7 = v25
				v8 = v22
				v9 = v24
				continue
			} else {
				break
			}
			break
		}
		v27 = v22
	} else {
		v27 = v3
	}
	return v27
}
func F_guc_strdup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1436]))
	v7 = F_strlen(m, l1)
	mBase = m.M
	v9 = v7 + int32(1)
	v11 = F_MemoryContextAllocExtended(m, v6, v9, int32(2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v18 = F_errstart(m, l0, int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					return v11
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(13904), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499912), int32(647), int32(489018))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								return v11
							}
						}
					}
				}
			}
		} else {
			if v9 != 0 {
				v35 = F__emscripten_memcpy_bulkmem(m, v11, l1, v9)
				mBase = m.M
			} else {
			}
			return v11
		}
	}
}
