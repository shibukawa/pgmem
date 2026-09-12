package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetConfigOption(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v12 = F_find_option(m, l0, v3, l1, int32(21))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return v91
L2:
	;
	return int32(0)
L3:
	;
	if v12 == int32(0) {
		v91 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	switch v18 {
	case 0:
		goto L5
	case 1:
		goto L9
	case 2:
		goto L8
	case 3:
		goto L7
	case 4:
		goto L6
	default:
		v91 = v3
		goto L1
	}
L5:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)+100))
	if v45 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	*(*float64)(unsafe.Add(mBase, uint32(v8)+16)) = v29
	v31 = int32(4513776)
	v37 = F_pg_snprintf(m, v31, int32(256), int32(338438), v8+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v20
	v22 = int32(4513776)
	v26 = F_pg_snprintf(m, v22, int32(256), int32(488641), v8)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v91 = v22
	goto L1
L11:
	;
	v91 = v31
	goto L1
L12:
	;
	v42 = v40
	goto L14
L13:
	;
	v42 = int32(757756)
	goto L14
L14:
	;
	v91 = v42
	goto L1
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L2
	} else {
		goto L22
	}
L16:
	;
	v48 = v45
	goto L17
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v53 == int32(0) {
		goto L15
	} else {
		goto L19
	}
L18:
	;
	goto L15
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v56 == v44 {
		v91 = v53
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v59 = v48 + int32(12)
	if v59 != 0 {
		v48 = v59
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v44
	F_errmsg_internal(m, int32(181726), v8+int32(32))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(500207), int32(3036), int32(345230))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	v86 = int32(273318)
	goto L27
L26:
	;
	v86 = int32(339246)
	goto L27
L27:
	;
	v91 = v86
	goto L1
}
func F_ProcessConfigFile(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	v8 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v13 = F_AllocSetContextCreateInternal(m, v8, int32(329726), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(4515712)
		v16 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v13
		if v6 != 0 {
			v22 = int32(13)
		} else {
			v22 = int32(15)
		}
		v23 = F_ProcessConfigFileInternal(m, l0, int32(1), v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
			F_MemoryContextDelete(m, v13)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_SelectConfigFiles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
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
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
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
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(240)
	m.G0 = v8
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	m.G0 = v8 + int32(240)
	return v390
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = l1
	F_write_stderr(m, int32(755630), v8+int32(128))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L25
	} else {
		goto L108
	}
L3:
	;
	v152 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	v159 = F___fstatat(m, int32(-100), v154, v8+int32(144), v152)
	mBase = m.M
	goto L50
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = int32(338955)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v60
	v135 = F_pg_sprintf(m, v92, int32(177111), v8+int32(96))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L25
	} else {
		goto L47
	}
L5:
	;
	v117 = F_make_absolute_path(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L25
	} else {
		goto L45
	}
L6:
	;
	v59 = l0
	goto L8
L7:
	;
	v10 = int32(545386)
	v11 = int32(0)
	v16 = F___strchrnul(m, v10, int32(61))
	mBase = m.M
	if v10 == v16 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v60 = F_make_absolute_path(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L9:
	;
	v59 = v58
	goto L8
L10:
	;
	v58 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v19 = v16 - v10
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[1130]))))
	if v21 != 0 {
		v51 = v11
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v58 = v51
	goto L9
L14:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[1115]))
	if v23 == int32(0) {
		v51 = v11
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 == int32(0) {
		v51 = v11
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v30 = v23
	v31 = v26
	goto L17
L17:
	;
	v34 = F_strncmp(m, v10, v31, v19)
	mBase = m.M
	if v34 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v51 = v38 + int32(1)
	goto L13
L19:
	;
	goto L18
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v38 = v37 + v19
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 == int32(61) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v43 != 0 {
		v30 = v30 + int32(4)
		v31 = v43
		goto L17
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v51 = v11
	goto L13
L25:
	;
	return int32(0)
L26:
	;
	if v60 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v68 = F___fstatat(m, int32(-100), v60, v8+int32(144), int32(0))
	mBase = m.M
	goto L30
L28:
	;
	goto L29
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	if v111 == int32(0) {
		goto L2
	} else {
		goto L44
	}
L30:
	;
	if v68 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = l1
	F_write_stderr(m, int32(749307), v8+int32(112))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L25
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	if v85 != 0 {
		v114 = v85
		goto L5
	} else {
		goto L37
	}
L34:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v77 != int32(44) {
		v390 = v3
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_write_stderr(m, int32(754732), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v390 = v3
	goto L1
L37:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1131]))
	v88 = F_strlen(m, v60)
	mBase = m.M
	v92 = F_MemoryContextAllocExtended(m, v87, v88+int32(17), int32(2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	if v92 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L25
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(500207), int32(647), int32(489277))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v114 = v111
	goto L5
L45:
	;
	v119 = int32(1)
	v120 = int32(10)
	v122 = int32(0)
	v126 = F_set_config_with_handle(m, int32(387017), int32(0), v117, v119, v120, v120, v122, v119, v122, v122)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	F_emscripten_builtin_free(m, v117)
	mBase = m.M
	goto L3
L47:
	;
	v138 = int32(0)
	v139 = int32(1)
	v140 = int32(10)
	v146 = F_set_config_with_handle(m, int32(387017), v138, v92, v139, v140, v140, v138, v139, v138, v138)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L48
	}
L48:
	;
	F_pfree(m, v92)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	goto L3
L50:
	;
	if v159 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v162
	F_write_stderr(m, int32(749348), v8+int32(80))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L25
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L25
	} else {
		goto L55
	}
L54:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v390 = v152
	goto L1
L55:
	;
	v174 = int32(0)
	v177 = F_find_option(m, int32(13273), v174, v174, int32(23))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+92))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v180|v60 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l1
	v186 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v186
	F_write_stderr(m, int32(755090), v8-int32(-64))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L25
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v180 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v390 = v152
	goto L1
L61:
	;
	v193 = v180
	goto L63
L62:
	;
	v193 = v60
	goto L63
L63:
	;
	v194 = F_make_absolute_path(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L25
	} else {
		goto L64
	}
L64:
	;
	v196 = int32(4510380)
	v197 = *(*int32)(unsafe.Add(mBase, _consts[253]))
	F_emscripten_builtin_free(m, v197)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _consts[253])) = v194
	v202 = int32(0)
	v205 = int32(1)
	v206 = int32(10)
	v212 = F_set_config_with_handle(m, int32(13273), v202, v194, v205, v206, v206, v202, v205, v202, v202)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L25
	} else {
		goto L65
	}
L65:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	v218 = int32(0)
	v220 = int32(1)
	v227 = F_set_config_with_handle(m, int32(143510), v218, int32(98641), v220, v220, int32(10), v218, v220, v218, v218)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[1132]))
	if v230 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L68:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v390 = int32(1)
	goto L1
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(338920)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v60
	v366 = F_pg_sprintf(m, v334, int32(177111), v8)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L25
	} else {
		goto L105
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
	v355 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v355
	F_write_stderr(m, int32(755270), v8+int32(48))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L25
	} else {
		goto L104
	}
L71:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _consts[1131]))
	v330 = F_strlen(m, v60)
	mBase = m.M
	v334 = F_MemoryContextAllocExtended(m, v329, v330+int32(15), int32(2))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L25
	} else {
		goto L98
	}
L72:
	;
	if v60 == int32(0) {
		goto L70
	} else {
		goto L97
	}
L73:
	;
	v313 = F_make_absolute_path(m, v310)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L25
	} else {
		goto L95
	}
L74:
	;
	if v60 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	goto L76
L76:
	;
	v294 = F_make_absolute_path(m, v230)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L25
	} else {
		goto L92
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(338971)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v60
	v273 = F_pg_sprintf(m, v239, int32(177111), v8+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L25
	} else {
		goto L88
	}
L78:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _consts[1131]))
	v235 = F_strlen(m, v60)
	mBase = m.M
	v239 = F_MemoryContextAllocExtended(m, v234, v235+int32(13), int32(2))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L25
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	v260 = *(*int32)(unsafe.Add(mBase, _consts[1129]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v260
	F_write_stderr(m, int32(755452), v8+int32(32))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L25
	} else {
		goto L87
	}
L81:
	;
	if v239 != 0 {
		goto L77
	} else {
		goto L82
	}
L82:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(500207), int32(647), int32(489277))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L25
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
	v390 = int32(0)
	goto L1
L88:
	;
	v276 = int32(0)
	v277 = int32(1)
	v278 = int32(10)
	v284 = F_set_config_with_handle(m, int32(387188), v276, v239, v277, v278, v278, v276, v277, v276, v276)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L25
	} else {
		goto L89
	}
L89:
	;
	F_pfree(m, v239)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[1133]))
	if v289 == int32(0) {
		goto L71
	} else {
		goto L91
	}
L91:
	;
	v310 = v289
	goto L73
L92:
	;
	v296 = int32(1)
	v297 = int32(10)
	v299 = int32(0)
	v303 = F_set_config_with_handle(m, int32(387188), int32(0), v294, v296, v297, v297, v299, v296, v299, v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	F_emscripten_builtin_free(m, v294)
	mBase = m.M
	v307 = *(*int32)(unsafe.Add(mBase, _consts[1133]))
	if v307 == int32(0) {
		goto L72
	} else {
		goto L94
	}
L94:
	;
	v310 = v307
	goto L73
L95:
	;
	v315 = int32(1)
	v316 = int32(10)
	v318 = int32(0)
	v322 = F_set_config_with_handle(m, int32(386797), int32(0), v313, v315, v316, v316, v318, v315, v318, v318)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L25
	} else {
		goto L96
	}
L96:
	;
	F_emscripten_builtin_free(m, v313)
	mBase = m.M
	goto L68
L97:
	;
	goto L71
L98:
	;
	if v334 != 0 {
		goto L69
	} else {
		goto L99
	}
L99:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L25
	} else {
		goto L100
	}
L100:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L25
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(13904), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L25
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(500207), int32(647), int32(489277))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L25
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v390 = int32(0)
	goto L1
L105:
	;
	v369 = int32(0)
	v370 = int32(1)
	v371 = int32(10)
	v377 = F_set_config_with_handle(m, int32(386797), v369, v334, v370, v371, v371, v369, v370, v369, v369)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L25
	} else {
		goto L106
	}
L106:
	;
	F_pfree(m, v334)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L25
	} else {
		goto L107
	}
L107:
	;
	goto L68
L108:
	;
	v390 = v3
	goto L1
}
func F_set_config_option(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	if base.B2i32(l3 != int32(9))&base.B2i32(base.Ui32(l3) <= base.Ui32(int32(10))) != 0 {
		v16 = int32(10)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v16 = v15
	}
	v17 = int32(0)
	v19 = F_set_config_with_handle(m, l0, int32(0), l1, l2, l3, v16, l4, l5, v17, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		return
	}
}
