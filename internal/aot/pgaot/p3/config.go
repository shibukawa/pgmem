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
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v93 != 0 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
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
	F_errfinish(m, int32(_a_F_GetConfigOption_5), int32(2938), int32(_a_F_GetConfigOption_6))
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
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(272)
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
	v416 = m.ExcPending
	if v416 != 0 {
		goto L26
	} else {
		goto L112
	}
L2:
	;
	F_emscripten_builtin_free(m, v60)
	mBase = m.M
	m.G0 = v8 + int32(272)
	return v403
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = l1
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_0), v8+int32(160))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L26
	} else {
		goto L111
	}
L4:
	;
	v136 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	v143 = F___fstatat(m, int32(-100), v138, v8+int32(176), v136)
	mBase = m.M
	goto L47
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = int32(_a_F_SelectConfigFiles_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v60
	v119 = F_pg_sprintf(m, v92, int32(_a_F_SelectConfigFiles_2), v8+int32(128))
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
	v68 = F___fstatat(m, int32(-100), v60, v8+int32(176), int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = l1
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_4), v8+int32(144))
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
		v403 = v3
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
	v403 = v3
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = l1
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v146
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_7), v8+int32(112))
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
	v155 = m.ExcPending
	if v155 != 0 {
		goto L26
	} else {
		goto L52
	}
L51:
	;
	v403 = v136
	goto L2
L52:
	;
	v157 = int32(0)
	v160 = F_find_option(m, int32(_a_F_SelectConfigFiles_8), v157, v157, int32(24))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v160)+96))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v163|v60 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = l1
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+100)) = v169
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_9), v8+int32(96))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L26
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v163 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v403 = v136
	goto L2
L58:
	;
	v176 = v163
	goto L60
L59:
	;
	v176 = v60
	goto L60
L60:
	;
	F_SetDataDir(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L26
	} else {
		goto L61
	}
L61:
	;
	v180 = int32(0)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[5]))
	v183 = int32(1)
	v184 = int32(10)
	v190 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_8), v180, v182, v183, v184, v184, v180, v183, v180, v180)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L26
	} else {
		goto L62
	}
L62:
	;
	F_ProcessConfigFile(m, int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L26
	} else {
		goto L63
	}
L63:
	;
	v196 = int32(0)
	v198 = int32(1)
	v205 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_10), v196, int32(_a_F_SelectConfigFiles_11), v198, v198, int32(10), v196, v198, v196, v196)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L26
	} else {
		goto L64
	}
L64:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[6]))
	if v208 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L65:
	;
	v367 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v368 = F_strlen(m, v60)
	mBase = m.M
	v372 = F_MemoryContextAllocExtended(m, v367, v368+int32(15), int32(2))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L26
	} else {
		goto L104
	}
L66:
	;
	v351 = F_make_absolute_path(m, v347)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L26
	} else {
		goto L102
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_SelectConfigFiles_12)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v60
	v326 = F_pg_sprintf(m, v308, int32(_a_F_SelectConfigFiles_2), v8+int32(16))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L26
	} else {
		goto L98
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l1
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v313
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_13), v8+int32(80))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L26
	} else {
		goto L97
	}
L69:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v304 = F_strlen(m, v60)
	mBase = m.M
	v308 = F_MemoryContextAllocExtended(m, v303, v304+int32(15), int32(2))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L26
	} else {
		goto L95
	}
L70:
	;
	if v60 == int32(0) {
		goto L68
	} else {
		goto L94
	}
L71:
	;
	v275 = F_make_absolute_path(m, v272)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L26
	} else {
		goto L89
	}
L72:
	;
	if v60 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	goto L74
L74:
	;
	v256 = F_make_absolute_path(m, v208)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L26
	} else {
		goto L86
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(_a_F_SelectConfigFiles_14)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v60
	v235 = F_pg_sprintf(m, v217, int32(_a_F_SelectConfigFiles_2), v8+int32(48))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L26
	} else {
		goto L82
	}
L76:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[4]))
	v213 = F_strlen(m, v60)
	mBase = m.M
	v217 = F_MemoryContextAllocExtended(m, v212, v213+int32(13), int32(2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L26
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = l1
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v222
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_15), v8-int32(-64))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L26
	} else {
		goto L81
	}
L79:
	;
	if v217 != 0 {
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L1
L81:
	;
	v403 = int32(0)
	goto L2
L82:
	;
	v238 = int32(0)
	v239 = int32(1)
	v240 = int32(10)
	v246 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_16), v238, v217, v239, v240, v240, v238, v239, v238, v238)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L26
	} else {
		goto L83
	}
L83:
	;
	F_pfree(m, v217)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L26
	} else {
		goto L84
	}
L84:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[7]))
	if v251 == int32(0) {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v272 = v251
	goto L71
L86:
	;
	v258 = int32(1)
	v259 = int32(10)
	v261 = int32(0)
	v265 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_16), int32(0), v256, v258, v259, v259, v261, v258, v261, v261)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L26
	} else {
		goto L87
	}
L87:
	;
	F_emscripten_builtin_free(m, v256)
	mBase = m.M
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[7]))
	if v269 == int32(0) {
		goto L70
	} else {
		goto L88
	}
L88:
	;
	v272 = v269
	goto L71
L89:
	;
	v277 = int32(1)
	v278 = int32(10)
	v280 = int32(0)
	v284 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_17), int32(0), v275, v277, v278, v278, v280, v277, v280, v280)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L26
	} else {
		goto L90
	}
L90:
	;
	F_emscripten_builtin_free(m, v275)
	mBase = m.M
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[8]))
	if v288 != 0 {
		v347 = v288
		goto L66
	} else {
		goto L91
	}
L91:
	;
	if v60 != 0 {
		goto L65
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v292
	F_write_stderr(m, int32(_a_F_SelectConfigFiles_18), v8+int32(32))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L26
	} else {
		goto L93
	}
L93:
	;
	v403 = int32(0)
	goto L2
L94:
	;
	goto L69
L95:
	;
	if v308 != 0 {
		goto L67
	} else {
		goto L96
	}
L96:
	;
	goto L1
L97:
	;
	v403 = int32(0)
	goto L2
L98:
	;
	v329 = int32(0)
	v330 = int32(1)
	v331 = int32(10)
	v337 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_17), v329, v308, v330, v331, v331, v329, v330, v329, v329)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L26
	} else {
		goto L99
	}
L99:
	;
	F_pfree(m, v308)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L26
	} else {
		goto L100
	}
L100:
	;
	v342 = *(*int32)(unsafe.Add(mBase, _c_F_SelectConfigFiles[8]))
	if v342 == int32(0) {
		goto L65
	} else {
		goto L101
	}
L101:
	;
	v347 = v342
	goto L66
L102:
	;
	v353 = int32(1)
	v354 = int32(10)
	v356 = int32(0)
	v360 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_19), int32(0), v351, v353, v354, v354, v356, v353, v356, v356)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L26
	} else {
		goto L103
	}
L103:
	;
	F_emscripten_builtin_free(m, v351)
	mBase = m.M
	v403 = int32(1)
	goto L2
L104:
	;
	if v372 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	goto L1
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_SelectConfigFiles_20)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v60
	v380 = F_pg_sprintf(m, v372, int32(_a_F_SelectConfigFiles_2), v8)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L26
	} else {
		goto L108
	}
L108:
	;
	v382 = int32(1)
	v384 = int32(0)
	v386 = int32(10)
	v392 = F_set_config_with_handle(m, int32(_a_F_SelectConfigFiles_19), v384, v372, v382, v386, v386, v384, v382, v384, v384)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L26
	} else {
		goto L109
	}
L109:
	;
	F_pfree(m, v372)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L26
	} else {
		goto L110
	}
L110:
	;
	v403 = v382
	goto L2
L111:
	;
	v403 = v3
	goto L2
L112:
	;
	F_errcode(m, int32(_a_F_SelectConfigFiles_21))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L26
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_SelectConfigFiles_22), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L26
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_SelectConfigFiles_23), int32(646), int32(_a_F_SelectConfigFiles_24))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L26
	} else {
		goto L115
	}
L115:
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
