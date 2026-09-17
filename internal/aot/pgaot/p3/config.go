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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v13 = F_find_option(m, l0, v3, l1, int32(21))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(80)
	return v98
L2:
	;
	return int32(0)
L3:
	;
	if v13 == int32(0) {
		v98 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	switch v19 {
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
		v98 = v3
		goto L1
	}
L5:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v30
	v32 = int32(_a_F_GetConfigOption_0)
	v38 = F_pg_snprintf(m, v32, int32(256), int32(_a_F_GetConfigOption_1), v9+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
	v23 = int32(_a_F_GetConfigOption_0)
	v27 = F_pg_snprintf(m, v23, int32(256), int32(_a_F_GetConfigOption_2), v9)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v98 = v23
	goto L1
L11:
	;
	v98 = v32
	goto L1
L12:
	;
	v43 = v41
	goto L14
L13:
	;
	v43 = int32(_a_F_GetConfigOption_3)
	goto L14
L14:
	;
	v98 = v43
	goto L1
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L23
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v49 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v52 == v45 {
		v98 = v49
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v54 = v46
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v60 == int32(0) {
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v98 = v60
	goto L1
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v45 != v63 {
		v54 = v54 + int32(12)
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v45
	F_errmsg_internal(m, int32(_a_F_GetConfigOption_4), v9+int32(32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_GetConfigOption_5), int32(3036), int32(_a_F_GetConfigOption_6))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v94 = int32(_a_F_GetConfigOption_7)
	goto L28
L27:
	;
	v94 = int32(_a_F_GetConfigOption_8)
	goto L28
L28:
	;
	v98 = v94
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
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessConfigFile[0])))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFile[1]))
	v13 = F_AllocSetContextCreateInternal(m, v8, int32(_a_F_ProcessConfigFile_0), int32(0), int32(_a_F_ProcessConfigFile_1), int32(_a_F_ProcessConfigFile_2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(_a_F_ProcessConfigFile_3)
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFile[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFile[1])) = v13
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
			*(*int32)(unsafe.Add(mBase, _c_F_ProcessConfigFile[1])) = v16
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
	var v52 int32
	_ = v52
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(240)
	m.G0 = v8
	if l0 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L26
	} else {
		goto L98
	}
L2:
	;
	m.G0 = v8 + int32(240)
	return v346
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = l1
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_0), v8+int32(128))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L26
	} else {
		goto L97
	}
L4:
	;
	v136 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	v143 = F___fstatat(m, int32(-100), v138, v8+int32(144), v136)
	mBase = m.M
	goto L47
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = int32(_a_F_SelectConfigFiles_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v60
	v119 = F_pg_sprintf(m, v92, int32(_a_F_SelectConfigFiles_2), v8+int32(96))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L26
	} else {
		goto L44
	}
L6:
	;
	v101 = F_make_absolute_path(m, v98)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L26
	} else {
		goto L42
	}
L7:
	;
	v59 = l0
	goto L9
L8:
	;
	v10 = int32(_a_F_SelectConfigFiles_3)
	v11 = int32(0)
	v16 = F___strchrnul(m, v10, int32(61))
	mBase = m.M
	if v10 == v16 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v60 = F_make_absolute_path(m, v59)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L26
	} else {
		goto L27
	}
L10:
	;
	v59 = v58
	goto L9
L11:
	;
	v58 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v19 = v16 - v10
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_SelectConfigFiles[1]))))
	if v21 != 0 {
		v52 = v11
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v52
	goto L10
L15:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[2]))
	if v23 == int32(0) {
		v52 = v11
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 == int32(0) {
		v52 = v11
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v30 = v23
	v31 = v26
	goto L18
L18:
	;
	v34 = F_strncmp(m, v10, v31, v19)
	mBase = m.M
	if v34 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v52 = v38 + int32(1)
	goto L14
L20:
	;
	goto L19
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v38 = v37 + v19
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 == int32(61) {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v43 != 0 {
		v30 = v30 + int32(4)
		v31 = v43
		goto L18
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v52 = v11
	goto L14
L26:
	;
	return int32(0)
L27:
	;
	if v60 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v68 = F___fstatat(m, int32(-100), v60, v8+int32(144), int32(0))
	mBase = m.M
	goto L31
L29:
	;
	goto L30
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	if v95 == int32(0) {
		goto L3
	} else {
		goto L41
	}
L31:
	;
	if v68 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = l1
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_4), v8+int32(112))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L26
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	if v85 != 0 {
		v98 = v85
		goto L6
	} else {
		goto L38
	}
L35:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[3]))
	if v77 != int32(44) {
		v346 = v3
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_5), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v346 = v3
	goto L2
L38:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v88 = F_strlen(m, v60)
	mBase = m.M
	v92 = F_MemoryContextAllocExtended(m, v87, v88+int32(17), int32(2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L26
	} else {
		goto L39
	}
L39:
	;
	if v92 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	goto L1
L41:
	;
	v98 = v95
	goto L6
L42:
	;
	v103 = int32(1)
	v104 = int32(10)
	v106 = int32(0)
	v110 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_6), int32(0), v101, v103, v104, v104, v106, v103, v106, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L26
	} else {
		goto L43
	}
L43:
	;
	F_emscripten_builtin_free(m, v101)
	mBase = m.M
	goto L4
L44:
	;
	v122 = int32(0)
	v123 = int32(1)
	v124 = int32(10)
	v130 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_6), v122, v92, v123, v124, v124, v122, v123, v122, v122)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	F_pfree(m, v92)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L26
	} else {
		goto L46
	}
L46:
	;
	goto L4
L47:
	;
	if v143 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v146
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_7), v8+int32(80))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L26
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L26
	} else {
		goto L52
	}
L51:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v346 = v136
	goto L2
L52:
	;
	v158 = int32(0)
	v161 = F_find_option(m, int32(_a_F_SelectConfigFiles_8), v158, v158, int32(23))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v161)+92))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v164|v60 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l1
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v170
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_9), v8-int32(-64))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L26
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v164 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v346 = v136
	goto L2
L58:
	;
	v177 = v164
	goto L60
L59:
	;
	v177 = v60
	goto L60
L60:
	;
	v178 = F_make_absolute_path(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v180 = int32(_a_F_SelectConfigFiles_10)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[5]))
	F_emscripten_builtin_free(m, v181)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[5])) = v178
	v186 = int32(0)
	v189 = int32(1)
	v190 = int32(10)
	v196 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_8), v186, v178, v189, v190, v190, v186, v189, v186, v186)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	v202 = int32(0)
	v204 = int32(1)
	v211 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_11), v202, int32(_a_F_SelectConfigFiles_12), v204, v204, int32(10), v202, v204, v202, v202)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[6]))
	if v214 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	v346 = int32(1)
	goto L2
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_SelectConfigFiles_13)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v60
	v320 = F_pg_sprintf(m, v304, int32(_a_F_SelectConfigFiles_2), v8)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L26
	} else {
		goto L94
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v309
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_14), v8+int32(48))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L26
	} else {
		goto L93
	}
L68:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v300 = F_strlen(m, v60)
	mBase = m.M
	v304 = F_MemoryContextAllocExtended(m, v299, v300+int32(15), int32(2))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L26
	} else {
		goto L91
	}
L69:
	;
	if v60 == int32(0) {
		goto L67
	} else {
		goto L90
	}
L70:
	;
	v282 = F_make_absolute_path(m, v279)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L26
	} else {
		goto L88
	}
L71:
	;
	if v60 != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	goto L73
L73:
	;
	v262 = F_make_absolute_path(m, v214)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L26
	} else {
		goto L85
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_SelectConfigFiles_15)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v60
	v241 = F_pg_sprintf(m, v223, int32(_a_F_SelectConfigFiles_2), v8+int32(16))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L26
	} else {
		goto L81
	}
L75:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v219 = F_strlen(m, v60)
	mBase = m.M
	v223 = F_MemoryContextAllocExtended(m, v218, v219+int32(13), int32(2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L26
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v228
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_16), v8+int32(32))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L26
	} else {
		goto L80
	}
L78:
	;
	if v223 != 0 {
		goto L74
	} else {
		goto L79
	}
L79:
	;
	goto L1
L80:
	;
	v346 = int32(0)
	goto L2
L81:
	;
	v244 = int32(0)
	v245 = int32(1)
	v246 = int32(10)
	v252 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_17), v244, v223, v245, v246, v246, v244, v245, v244, v244)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L26
	} else {
		goto L82
	}
L82:
	;
	F_pfree(m, v223)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L26
	} else {
		goto L83
	}
L83:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[7]))
	if v257 == int32(0) {
		goto L68
	} else {
		goto L84
	}
L84:
	;
	v279 = v257
	goto L70
L85:
	;
	v264 = int32(1)
	v265 = int32(10)
	v267 = int32(0)
	v271 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_17), int32(0), v262, v264, v265, v265, v267, v264, v267, v267)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L26
	} else {
		goto L86
	}
L86:
	;
	F_emscripten_builtin_free(m, v262)
	mBase = m.M
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[7]))
	if v275 == int32(0) {
		goto L69
	} else {
		goto L87
	}
L87:
	;
	v279 = v275
	goto L70
L88:
	;
	v284 = int32(1)
	v285 = int32(10)
	v287 = int32(0)
	v291 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_18), int32(0), v282, v284, v285, v285, v287, v284, v287, v287)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L89
	}
L89:
	;
	F_emscripten_builtin_free(m, v282)
	mBase = m.M
	goto L65
L90:
	;
	goto L68
L91:
	;
	if v304 != 0 {
		goto L66
	} else {
		goto L92
	}
L92:
	;
	goto L1
L93:
	;
	v346 = int32(0)
	goto L2
L94:
	;
	v323 = int32(0)
	v324 = int32(1)
	v325 = int32(10)
	v331 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_18), v323, v304, v324, v325, v325, v323, v324, v323, v323)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L26
	} else {
		goto L95
	}
L95:
	;
	F_pfree(m, v304)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L26
	} else {
		goto L96
	}
L96:
	;
	goto L65
L97:
	;
	v346 = v3
	goto L2
L98:
	;
	F_errcode(m, int32(_a_F_SelectConfigFiles_19))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L26
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_SelectConfigFiles_20), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L26
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_SelectConfigFiles_21), int32(647), int32(_a_F_SelectConfigFiles_22))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L26
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_set_config_option[0]))
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
