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
			F_GUC_flex_fatal(m, int32(_a_F_GUC_yyensure_buffer_stack_0))
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
				F_GUC_flex_fatal(m, int32(_a_F_GUC_yyensure_buffer_stack_0))
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
				*(*int64)(unsafe.Add(mBase, uint32(v32)+24)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v33
				*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
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
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
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
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 < v20 {
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
	m.G0 = v12 + int32(16)
	return
L4:
	;
	v36 = F_array_ref(m, l0, v12+int32(12), v12+int32(11))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = F_text_to_cstring(m, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v301 = v299 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v301
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v301 <= v303 {
		goto L4
	} else {
		goto L86
	}
L11:
	;
	v250 = v242
	goto L69
L12:
	;
	v43 = int32(_a_F_TransformGUCArray_0)
	v47 = m.G0
	v49 = v47 - int32(32)
	m.G0 = v49
	v51 = int32(*(*int8)(unsafe.Add(mBase, _c_F_TransformGUCArray[0])))
	if v51 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109+v41))))
	if v111 == int32(61) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	m.G0 = v49 + int32(32)
	v109 = v102 - v41
	goto L13
L15:
	;
	F___memset(m, v49, int32(0), int32(32))
	mBase = m.M
	v57 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransformGUCArray[0])))
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransformGUCArray[1])))
	if v52 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v53 = F___strchrnul(m, v41, v51)
	mBase = m.M
	v102 = v53
	goto L14
L19:
	;
	goto L18
L20:
	;
	v59 = v43
	v60 = v57
	goto L23
L21:
	;
	goto L22
L22:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v81 == int32(0) {
		v102 = v41
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v67 = v49 + int32(base.Ui32(v60)>>(uint(int32(3))%32))&int32(28)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v69 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v68 | v69<<(uint(v60)%32)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v73 != 0 {
		v59 = v59 + v69
		v60 = v73
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
	v85 = v41
	v86 = v81
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(base.Ui32(v86)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v94)>>(uint(v86)%32))&int32(1) != 0 {
		v102 = v85
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v102 = v100
	goto L14
L29:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v100 = v85 + int32(1)
	if v98 != 0 {
		v85 = v100
		v86 = v98
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v115 = v109 + int32(1)
	v116 = F_palloc(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v240 = F_pstrdup(m, v41)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L67
	}
L34:
	;
	if v115 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v237 = F_pstrdup(m, v115+v41)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L66
	}
L36:
	;
	v233 = F_strlen(m, v229)
	mBase = m.M
	goto L35
L37:
	;
	v229 = v41
	goto L36
L38:
	;
	goto L39
L39:
	;
	v123 = v115 - int32(1)
	if (v116^v41)&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v226 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v226)
	v229 = v222
	goto L36
L41:
	;
	v207 = v202
	v208 = v203
	v209 = v204
	goto L62
L42:
	;
	if v197 == int32(0) {
		v222 = v195
		v223 = v196
		goto L40
	} else {
		goto L61
	}
L43:
	;
	v195 = v41
	v196 = v116
	v197 = v123
	goto L42
L44:
	;
	goto L45
L45:
	;
	v127 = int32(0)
	if base.B2i32(v41&int32(3) == v127)|base.B2i32(v123 == v127) == v127 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v163 == int32(0) {
		v222 = v160
		v223 = v161
		goto L40
	} else {
		goto L55
	}
L47:
	;
	v139 = v41
	v140 = v116
	v141 = v123
	goto L50
L48:
	;
	goto L49
L49:
	;
	v160 = v41
	v161 = v116
	v162 = v123
	v163 = base.B2i32(v123 != v127)
	goto L46
L50:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v143)
	if v143 == int32(0) {
		v202 = v139
		v203 = v140
		v204 = v141
		goto L41
	} else {
		goto L52
	}
L51:
	;
	v160 = v154
	v161 = v148
	v162 = v150
	v163 = v152
	goto L46
L52:
	;
	v147 = int32(1)
	v148 = v140 + v147
	v150 = v141 - v147
	v151 = int32(0)
	v152 = base.B2i32(v150 != v151)
	v154 = v139 + v147
	if v154&int32(3) == v151 {
		v160 = v154
		v161 = v148
		v162 = v150
		v163 = v152
		goto L46
	} else {
		goto L53
	}
L53:
	;
	if v150 != 0 {
		v139 = v154
		v140 = v148
		v141 = v150
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if base.B2i32(v166 == int32(0))|base.B2i32(base.Ui32(v162) < base.Ui32(int32(4))) != 0 {
		v195 = v160
		v196 = v161
		v197 = v162
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v173 = v160
	v174 = v161
	v175 = v162
	goto L57
L57:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v181 = int32(-2139062144)
	if (int32(16843008)-v178|v178)&v181 != v181 {
		v202 = v173
		v203 = v174
		v204 = v175
		goto L41
	} else {
		goto L59
	}
L58:
	;
	v195 = v189
	v196 = v187
	v197 = v191
	goto L42
L59:
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
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v202 = v195
	v203 = v196
	v204 = v197
	goto L41
L62:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v211)
	if v211 == int32(0) {
		v222 = v207
		v223 = v208
		goto L40
	} else {
		goto L64
	}
L63:
	;
	v222 = v218
	v223 = v216
	goto L40
L64:
	;
	v215 = int32(1)
	v216 = v208 + v215
	v218 = v207 + v215
	v220 = v209 - v215
	if v220 != 0 {
		v207 = v218
		v208 = v216
		v209 = v220
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v242 = v116
	v244 = v237
	goto L11
L67:
	;
	v242 = v240
	v244 = int32(0)
	goto L11
L68:
	;
	F_pfree(m, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L85
	}
L69:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	if v254 != int32(45) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v280 = F_lappend(m, v279, v242)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L6
	} else {
		goto L83
	}
L71:
	;
	goto L70
L72:
	;
	v250 = v250 + int32(1)
	goto L69
L73:
	;
	if v254 != 0 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v275 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v250))) = uint8(v275)
	goto L72
L76:
	;
	if v244 != 0 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v259 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v259 == int32(0) {
		v287 = v242
		goto L68
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v242
	F_errmsg(m, int32(_a_F_TransformGUCArray_1), v12)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_TransformGUCArray_2), int32(_a_F_TransformGUCArray_3), int32(_a_F_TransformGUCArray_4))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v287 = v242
	goto L68
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v280
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v284 = F_lappend(m, v283, v244)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v284
	v287 = v41
	goto L68
L85:
	;
	goto L10
L86:
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	v1 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[0]))
	v17 = F_AllocSetContextCreateInternal(m, v12, int32(_a_F_build_guc_variables_0), v1, int32(_a_F_build_guc_variables_1), int32(_a_F_build_guc_variables_2))
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
	*(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[1])) = v17
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[2]))
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = v1
	goto L6
L4:
	;
	v36 = v1
	goto L5
L5:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[3]))
	if v42 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v28 = v22 * int32(120)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_build_guc_variables[4]))) = int32(0)
	v34 = v22 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_build_guc_variables[5])))
	if v35 != 0 {
		v22 = v34
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v36 = v34
	goto L5
L8:
	;
	goto L7
L9:
	;
	v43 = v36
	v44 = v1
	goto L12
L10:
	;
	v63 = v36
	goto L11
L11:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[6]))
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v48 = int32(7)
	v52 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44<<(uint(v48)%32))+uint32(_c_F_build_guc_variables[7]))) = v52
	v55 = v43 + v52
	v57 = v44 + v52
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(v48)%32))+uint32(_c_F_build_guc_variables[3])))
	if v62 != 0 {
		v43 = v55
		v44 = v57
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v63 = v55
	goto L11
L14:
	;
	goto L13
L15:
	;
	v71 = v63
	v72 = int32(0)
	goto L18
L16:
	;
	v87 = v63
	goto L17
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[8]))
	if v94 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v77 = v72 * int32(152)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_build_guc_variables[9]))) = int32(2)
	v82 = int32(1)
	v85 = v71 + v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_build_guc_variables[10])))
	if v86 != 0 {
		v71 = v85
		v72 = v72 + v82
		goto L18
	} else {
		goto L20
	}
L19:
	;
	v87 = v85
	goto L17
L20:
	;
	goto L19
L21:
	;
	v95 = v87
	v96 = int32(0)
	goto L24
L22:
	;
	v111 = v87
	goto L23
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[11]))
	if v118 != 0 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v101 = v96 * int32(120)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_build_guc_variables[12]))) = int32(3)
	v106 = int32(1)
	v109 = v95 + v106
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_build_guc_variables[13])))
	if v110 != 0 {
		v95 = v109
		v96 = v96 + v106
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v111 = v109
	goto L23
L26:
	;
	goto L25
L27:
	;
	v119 = v111
	v120 = int32(0)
	goto L30
L28:
	;
	v135 = v111
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(1641)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(1642)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = int64(34359738372)
	v151 = base.I32_div_s(v135, int32(4))
	v156 = F_hash_create(m, int32(_a_F_build_guc_variables_3), v151+v135, v6+int32(-48), int32(1224))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	v125 = v120 * int32(124)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_build_guc_variables[14]))) = int32(4)
	v130 = int32(1)
	v133 = v119 + v130
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_build_guc_variables[15])))
	if v134 != 0 {
		v119 = v133
		v120 = v120 + v130
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v135 = v133
	goto L29
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16])) = v156
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[2]))
	if v160 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v162 = int32(0)
	goto L37
L35:
	;
	goto L36
L36:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[3]))
	if v188 != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16]))
	v169 = v162 * int32(120)
	v171 = v169 + int32(_a_F_build_guc_variables_4)
	v175 = F_hash_search(m, v167, v171, int32(1), v6+int32(-49))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v171
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_build_guc_variables[5])))
	if v180 != 0 {
		v162 = v162 + int32(1)
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v189 = int32(0)
	goto L44
L42:
	;
	goto L43
L43:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[6]))
	if v220 != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16]))
	v199 = v189<<(uint(int32(7))%32) + int32(_a_F_build_guc_variables_5)
	v203 = F_hash_search(m, v195, v199, int32(1), v6+int32(-49))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v203)+4)) = v199
	v207 = v189 + int32(1)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v207<<(uint(int32(7))%32))+uint32(_c_F_build_guc_variables[3])))
	if v212 != 0 {
		v189 = v207
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v221 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[8]))
	if v248 != 0 {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16]))
	v229 = v221 * int32(152)
	v231 = v229 + int32(_a_F_build_guc_variables_6)
	v235 = F_hash_search(m, v227, v231, int32(1), v6+int32(-49))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v231
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v229)+uint32(_c_F_build_guc_variables[10])))
	if v240 != 0 {
		v221 = v221 + int32(1)
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v249 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v276 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[11]))
	if v276 != 0 {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16]))
	v257 = v249 * int32(120)
	v259 = v257 + int32(_a_F_build_guc_variables_7)
	v263 = F_hash_search(m, v255, v259, int32(1), v6+int32(-49))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v263)+4)) = v259
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v257)+uint32(_c_F_build_guc_variables[13])))
	if v268 != 0 {
		v249 = v249 + int32(1)
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v277 = int32(0)
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
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[16]))
	v285 = v277 * int32(124)
	v287 = v285 + int32(_a_F_build_guc_variables_8)
	v291 = F_hash_search(m, v283, v287, int32(1), v6+int32(-49))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L64
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v291)+4)) = v287
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v285)+uint32(_c_F_build_guc_variables[15])))
	if v296 != 0 {
		v277 = v277 + int32(1)
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
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	v2 = int32(0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v7 == v2 {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_0)
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v218 = F_pstrdup(m, v213)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L73
	} else {
		goto L74
	}
L2:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v213 = v212
	goto L1
L3:
	;
	v125 = l0
	v126 = v2
	goto L45
L4:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	if v121 != 0 {
		goto L3
	} else {
		goto L43
	}
L5:
	;
	if base.Ui32((v7-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v18 = v7 | int32(32)
	goto L8
L7:
	;
	v18 = v7
	goto L8
L8:
	;
	if v18 != int32(115) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	if v22 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_1)
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if base.Ui32((v22-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = v22 | int32(32)
	goto L13
L12:
	;
	v33 = v22
	goto L13
L13:
	;
	if v33 != int32(111) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if v37 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_2)
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if base.Ui32((v37-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = v37 | int32(32)
	goto L18
L17:
	;
	v48 = v37
	goto L18
L18:
	;
	if v48 != int32(114) {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	if v52 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_3)
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if base.Ui32((v52-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v63 = v52 | int32(32)
	goto L23
L22:
	;
	v63 = v52
	goto L23
L23:
	;
	if v63 != int32(116) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v67 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_4)
		goto L4
	} else {
		goto L25
	}
L25:
	;
	if v67 != int32(95) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v73 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_5)
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v84 = v73 | int32(32)
	goto L30
L29:
	;
	v84 = v73
	goto L30
L30:
	;
	if v84 != int32(109) {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v88 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_6)
		goto L4
	} else {
		goto L32
	}
L32:
	;
	if base.Ui32((v88-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v99 = v88 | int32(32)
	goto L35
L34:
	;
	v99 = v88
	goto L35
L35:
	;
	if v99 != int32(101) {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	if v103 == int32(0) {
		v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_7)
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if base.Ui32((v103-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v114 = v103 | int32(32)
	goto L40
L39:
	;
	v114 = v103
	goto L40
L40:
	;
	if v114 != int32(109) {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v117 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	v120 = int32(_a_F_convert_GUC_name_for_parameter_acl_8)
	goto L4
L43:
	;
	v211 = int32(_a_F_convert_GUC_name_for_parameter_acl_9)
	goto L2
L44:
	;
	v168 = l0
	v169 = int32(0)
	goto L59
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v129 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v126 != int32(10) {
		goto L44
	} else {
		goto L58
	}
L47:
	;
	if v126 == int32(10) {
		goto L44
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	v134 = int32(1)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+uint32(_c_F_convert_GUC_name_for_parameter_acl[0]))))
	if base.Ui32((v138-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v147 = v138 | int32(32)
	goto L53
L52:
	;
	v147 = v138
	goto L53
L53:
	;
	v148 = int32(255)
	if base.Ui32((v129-int32(65))&v148) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v158 = v129 | int32(32)
	goto L56
L55:
	;
	v158 = v129
	goto L56
L56:
	;
	if v147&v148 == v158 {
		v125 = v125 + v134
		v126 = v126 + v134
		goto L45
	} else {
		goto L57
	}
L57:
	;
	goto L44
L58:
	;
	v211 = int32(_a_F_convert_GUC_name_for_parameter_acl_10)
	goto L2
L59:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if v172 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	if v169 != int32(14) {
		v213 = l0
		goto L1
	} else {
		goto L72
	}
L61:
	;
	if v169 == int32(14) {
		v213 = l0
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v177 = int32(1)
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+uint32(_c_F_convert_GUC_name_for_parameter_acl[1]))))
	if base.Ui32((v181-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v190 = v181 | int32(32)
	goto L67
L66:
	;
	v190 = v181
	goto L67
L67:
	;
	v191 = int32(255)
	if base.Ui32((v172-int32(65))&v191) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v201 = v172 | int32(32)
	goto L70
L69:
	;
	v201 = v172
	goto L70
L70:
	;
	if v190&v191 == v201 {
		v168 = v168 + v177
		v169 = v169 + v177
		goto L59
	} else {
		goto L71
	}
L71:
	;
	v213 = l0
	goto L1
L72:
	;
	v211 = int32(_a_F_convert_GUC_name_for_parameter_acl_11)
	goto L2
L73:
	;
	return int32(0)
L74:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if v222 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v223 = v218
	v225 = v222
	goto L78
L76:
	;
	goto L77
L77:
	;
	return v218
L78:
	;
	if base.Ui32((v225-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L77
L80:
	;
	v235 = v225 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v235)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v237 != 0 {
		v223 = v223 + int32(1)
		v225 = v237
		goto L78
	} else {
		goto L83
	}
L83:
	;
	goto L79
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
	var v23 int32
	_ = v23
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
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
			if v23 != 0 {
				v7 = v23
				v8 = v22
				v9 = v9 + int32(1)
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
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_guc_strdup[0]))
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
					F_errcode(m, int32(_a_F_guc_strdup_0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_guc_strdup_1), int32(0))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_guc_strdup_2), int32(647), int32(_a_F_guc_strdup_3))
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
			if v9 == int32(0) {
			} else {
				base.MemoryCopy(m, v11, l1, v9)
			}
			return v11
		}
	}
}
