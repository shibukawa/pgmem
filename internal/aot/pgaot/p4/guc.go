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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v4 < v21 {
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
	v38 = F_array_ref(m, l0, v13+int32(12), v13+int32(11))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if v40 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = F_text_to_cstring(m, base.I32_wrap_i64(v38))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v306 = v304 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v306
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v306 <= v308 {
		goto L4
	} else {
		goto L86
	}
L11:
	;
	v253 = v245
	goto L69
L12:
	;
	v46 = int32(_a_F_TransformGUCArray_0)
	v50 = m.G0
	v52 = v50 - int32(32)
	m.G0 = v52
	v54 = int32(*(*int8)(unsafe.Add(mBase, _c_F_TransformGUCArray[0])))
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v44))))
	if v114 == int32(61) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	m.G0 = v52 + int32(32)
	v112 = v105 - v44
	goto L13
L15:
	;
	F___memset(m, v52, int32(0), int32(32))
	mBase = m.M
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransformGUCArray[0])))
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TransformGUCArray[1])))
	if v55 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v56 = F___strchrnul(m, v44, v54)
	mBase = m.M
	v105 = v56
	goto L14
L19:
	;
	goto L18
L20:
	;
	v62 = v46
	v63 = v60
	goto L23
L21:
	;
	goto L22
L22:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v84 == int32(0) {
		v105 = v44
		goto L14
	} else {
		goto L26
	}
L23:
	;
	v70 = v52 + int32(base.Ui32(v63)>>(uint(int32(3))%32))&int32(28)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v71 | v72<<(uint(v63)%32)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	if v76 != 0 {
		v62 = v62 + v72
		v63 = v76
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
	v88 = v44
	v89 = v84
	goto L27
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v52+int32(base.Ui32(v89)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v97)>>(uint(v89)%32))&int32(1) != 0 {
		v105 = v88
		goto L14
	} else {
		goto L29
	}
L28:
	;
	v105 = v103
	goto L14
L29:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v103 = v88 + int32(1)
	if v101 != 0 {
		v88 = v103
		v89 = v101
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v118 = v112 + int32(1)
	v119 = F_palloc(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v243 = F_pstrdup(m, v44)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L67
	}
L34:
	;
	if v118 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v240 = F_pstrdup(m, v118+v44)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L66
	}
L36:
	;
	v236 = F_strlen(m, v232)
	mBase = m.M
	goto L35
L37:
	;
	v232 = v44
	goto L36
L38:
	;
	goto L39
L39:
	;
	v126 = v118 - int32(1)
	if (v119^v44)&int32(3) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v229 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v229)
	v232 = v225
	goto L36
L41:
	;
	v210 = v205
	v211 = v206
	v212 = v207
	goto L62
L42:
	;
	if v200 == int32(0) {
		v225 = v198
		v226 = v199
		goto L40
	} else {
		goto L61
	}
L43:
	;
	v198 = v44
	v199 = v119
	v200 = v126
	goto L42
L44:
	;
	goto L45
L45:
	;
	v130 = int32(0)
	if base.B2i32(v44&int32(3) == v130)|base.B2i32(v126 == v130) == v130 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v166 == int32(0) {
		v225 = v163
		v226 = v164
		goto L40
	} else {
		goto L55
	}
L47:
	;
	v142 = v44
	v143 = v119
	v144 = v126
	goto L50
L48:
	;
	goto L49
L49:
	;
	v163 = v44
	v164 = v119
	v165 = v126
	v166 = base.B2i32(v126 != v130)
	goto L46
L50:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v146)
	if v146 == int32(0) {
		v205 = v142
		v206 = v143
		v207 = v144
		goto L41
	} else {
		goto L52
	}
L51:
	;
	v163 = v157
	v164 = v151
	v165 = v153
	v166 = v155
	goto L46
L52:
	;
	v150 = int32(1)
	v151 = v143 + v150
	v153 = v144 - v150
	v154 = int32(0)
	v155 = base.B2i32(v153 != v154)
	v157 = v142 + v150
	if v157&int32(3) == v154 {
		v163 = v157
		v164 = v151
		v165 = v153
		v166 = v155
		goto L46
	} else {
		goto L53
	}
L53:
	;
	if v153 != 0 {
		v142 = v157
		v143 = v151
		v144 = v153
		goto L50
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if base.B2i32(v169 == int32(0))|base.B2i32(base.Ui32(v165) < base.Ui32(int32(4))) != 0 {
		v198 = v163
		v199 = v164
		v200 = v165
		goto L42
	} else {
		goto L56
	}
L56:
	;
	v176 = v163
	v177 = v164
	v178 = v165
	goto L57
L57:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v184 = int32(-2139062144)
	if (int32(16843008)-v181|v181)&v184 != v184 {
		v205 = v176
		v206 = v177
		v207 = v178
		goto L41
	} else {
		goto L59
	}
L58:
	;
	v198 = v192
	v199 = v190
	v200 = v194
	goto L42
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177))) = v181
	v189 = int32(4)
	v190 = v177 + v189
	v192 = v176 + v189
	v194 = v178 - v189
	if base.Ui32(int32(3)) < base.Ui32(v194) {
		v176 = v192
		v177 = v190
		v178 = v194
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v205 = v198
	v206 = v199
	v207 = v200
	goto L41
L62:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v214)
	if v214 == int32(0) {
		v225 = v210
		v226 = v211
		goto L40
	} else {
		goto L64
	}
L63:
	;
	v225 = v221
	v226 = v219
	goto L40
L64:
	;
	v218 = int32(1)
	v219 = v211 + v218
	v221 = v210 + v218
	v223 = v212 - v218
	if v223 != 0 {
		v210 = v221
		v211 = v219
		v212 = v223
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v245 = v119
	v247 = v240
	goto L11
L67:
	;
	v245 = v243
	v247 = int32(0)
	goto L11
L68:
	;
	F_pfree(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L6
	} else {
		goto L85
	}
L69:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
	if v258 != int32(45) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v284 = F_lappend(m, v283, v245)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L83
	}
L71:
	;
	goto L70
L72:
	;
	v253 = v253 + int32(1)
	goto L69
L73:
	;
	if v258 != 0 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v279 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v279)
	goto L72
L76:
	;
	if v247 != 0 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v263 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v263 == int32(0) {
		v291 = v245
		goto L68
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v245
	F_errmsg(m, int32(_a_F_TransformGUCArray_1), v13)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_TransformGUCArray_2), int32(_a_F_TransformGUCArray_3), int32(_a_F_TransformGUCArray_4))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v291 = v245
	goto L68
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v284
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v288 = F_lappend(m, v287, v247)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v288
	v291 = v44
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[0]))
	v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_build_guc_variables_0), v1, int32(_a_F_build_guc_variables_1), int32(_a_F_build_guc_variables_2))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[1])) = v16
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[2]))
	if v20 != 0 {
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
	v33 = v1
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = int32(1856)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(1857)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(34359738372)
	v46 = base.I32_div_s(v33, int32(4))
	v52 = F_hash_create(m, int32(_a_F_build_guc_variables_3), base.I64_extend_i32_s(v46+v33), v5+int32(-48), int32(1224))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v26 = v22 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26*int32(152))+uint32(_c_F_build_guc_variables[2])))
	if v31 != 0 {
		v22 = v26
		goto L6
	} else {
		goto L8
	}
L7:
	;
	v33 = v26
	goto L5
L8:
	;
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[3])) = v52
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[2]))
	if v56 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = v1
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v7 - int32(-64)
	return
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_build_guc_variables[3]))
	v64 = v59 * int32(152)
	v66 = v64 + int32(_a_F_build_guc_variables_4)
	v70 = F_hash_search(m, v62, v66, int32(1), v5+int32(-49))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = v66
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+uint32(_c_F_build_guc_variables[4])))
	if v75 != 0 {
		v59 = v59 + int32(1)
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
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
							F_errfinish(m, int32(_a_F_guc_strdup_2), int32(646), int32(_a_F_guc_strdup_3))
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
