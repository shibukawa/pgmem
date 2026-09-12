package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildSerialize(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	v12 = m.G0
	v14 = v12 - int32(2304)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 < int32(2) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L9
	} else {
		goto L111
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L9
	} else {
		goto L107
	}
L3:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v403 = F_CloseTransientFile(m, v251)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L9
	} else {
		goto L102
	}
L4:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v376 = F_CloseTransientFile(m, v251)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L9
	} else {
		goto L94
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L9
	} else {
		goto L90
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L9
	} else {
		goto L86
	}
L7:
	;
	m.G0 = v14 + int32(2304)
	return
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = int32(111036)
	v21 = base.I32_wrap_i64(l1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+152)) = v21
	v25 = base.I32_wrap_i64(int64(base.Ui64(l1) >> (uint(int64(32)) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+148)) = v25
	v32 = F_pg_sprintf(m, v14+int32(256), int32(225837), v14+int32(144))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v40 = F___fstatat(m, int32(-100), v14+int32(256), v14+int32(160), int32(0))
	mBase = m.M
	goto L12
L11:
	;
	v78 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L24
	}
L12:
	;
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v42 == int32(44) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_fsync_fname(m, v14+int32(256), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L9
	} else {
		goto L21
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = v14 + int32(256)
	F_errmsg(m, int32(282964), v14+int32(128))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(476291), int32(1546), int32(325376))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	F_fsync_fname(m, int32(111036), int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v74)+136)) = l1
	goto L23
L23:
	;
	goto L7
L24:
	;
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = v14 + int32(256)
	F_errmsg_internal(m, int32(172396), v14+int32(112))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = int32(111036)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+100)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v21
	v98 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v98
	v105 = F_pg_sprintf(m, v14+int32(1280), int32(223541), v14+int32(96))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	F_errfinish(m, int32(476291), int32(1573), int32(325376))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v109 = F_unlink(m, v14+int32(1280))
	mBase = m.M
	if v109 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v111 != int32(44) {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v114 = int32(4442576)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+28))
	if v120 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v123 = F_palloc(m, v120<<(uint(int32(2))%32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L38
	}
L36:
	;
	v181 = int32(0)
	goto L37
L37:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+28))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v189 = (v183+v184)<<(uint(int32(2))%32) + int32(104)
	v190 = F_palloc0(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L9
	} else {
		goto L46
	}
L38:
	;
	v125 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v119)+24))
	if v126 == v125 {
		v157 = v125
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pg_qsort(m, v123, v157, int32(4), int32(185))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L9
	} else {
		goto L45
	}
L40:
	;
	v130 = v119 + int32(20)
	if v126 == v130 {
		v157 = v125
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v135 = v125
	v136 = v126
	goto L42
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v136-int32(192))))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v135<<(uint(int32(2))%32)))) = v148
	v151 = v135 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v152 != v130 {
		v135 = v151
		v136 = v152
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v157 = v151
	goto L39
L44:
	;
	goto L43
L45:
	;
	v181 = v123
	goto L37
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+12)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(-2925404159)
	v198 = int32(8)
	v201 = m.Env.Pgmem_crc32c(m, int32(-1), v190+v198, v198)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v201
	goto L48
L47:
	;
	v208 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+100)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v190)+92)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v190)+72)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v190)+56)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v190)+20)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v190)+96)) = v183
	v220 = m.Env.Pgmem_crc32c(m, v201, v206, int32(88))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v220
	v223 = v190 + int32(104)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v224 != 0 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v206 = F__emscripten_memcpy_bulkmem(m, v190+int32(16), l0, int32(80))
	mBase = m.M
	goto L50
L50:
	;
	goto L47
L51:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v227 = v224 << (uint(int32(2)) % 32)
	if v227 != 0 {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	v233 = v220
	v234 = v223
	goto L53
L53:
	;
	if v183 != 0 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	v230 = m.Env.Pgmem_crc32c(m, v220, v229, v227)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v230
	v233 = v230
	v234 = v229 + v227
	goto L53
L55:
	;
	v228 = F__emscripten_memcpy_bulkmem(m, v223, v225, v227)
	mBase = m.M
	v229 = v228
	goto L57
L56:
	;
	v229 = v223
	goto L57
L57:
	;
	goto L54
L58:
	;
	v237 = v183 << (uint(int32(2)) % 32)
	if v237 != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v244 = v233
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v244 ^ int32(-1)
	v251 = F_OpenTransientFile(m, v14+int32(1280), int32(193))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L9
	} else {
		goto L65
	}
L61:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	v241 = m.Env.Pgmem_crc32c(m, v240, v239, v237)
	mBase = m.M
	v244 = v241
	goto L60
L62:
	;
	v238 = F__emscripten_memcpy_bulkmem(m, v234, v181, v237)
	mBase = m.M
	v239 = v238
	goto L64
L63:
	;
	v239 = v234
	goto L64
L64:
	;
	goto L61
L65:
	;
	if v251 < int32(0) {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = int32(167772216)
	v262 = F_write(m, v251, v190, v189)
	mBase = m.M
	if v262 != v189 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v264 = int32(4062620)
	v265 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266
	v269 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v269))) = int32(167772215)
	v274 = int32(*(*uint8)(unsafe.Add(mBase, _consts[162])))
	if v274 != int32(1) {
		v288 = v266
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v288 != 0 {
		goto L3
	} else {
		goto L75
	}
L69:
	;
	goto L68
L70:
	;
	goto L71
L71:
	;
	v279 = F_fsync(m, v251)
	mBase = m.M
	if v279 != int32(-1) {
		v288 = v279
		goto L69
	} else {
		goto L73
	}
L72:
	;
	v288 = int32(-1)
	goto L69
L73:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v283 == int32(27) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _consts[165]))
	*(*int32)(unsafe.Add(mBase, uint32(v290))) = int32(0)
	v293 = F_CloseTransientFile(m, v251)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
	} else {
		goto L76
	}
L76:
	;
	if v293 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	F_fsync_fname(m, int32(111036), int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v303 = F_rename(m, v14+int32(1280), v14+int32(256))
	mBase = m.M
	if v303 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_fsync_fname(m, v14+int32(256), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_fsync_fname(m, int32(111036), int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = l1
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v115
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v316)+136)) = l1
	goto L82
L82:
	;
	F_pfree(m, v190)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L9
	} else {
		goto L83
	}
L83:
	;
	if v181 == int32(0) {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	F_pfree(m, v181)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	goto L7
L86:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v14 + int32(1280)
	F_errmsg(m, int32(284416), v14+int32(80))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(476291), int32(1589), int32(325376))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L9
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v14 + int32(1280)
	F_errmsg(m, int32(283862), v14)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(476291), int32(1651), int32(325376))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L9
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
	if v375 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v380 = v375
	goto L97
L96:
	;
	v380 = int32(51)
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = v380
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L9
	} else {
		goto L98
	}
L98:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L9
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v14 + int32(1280)
	F_errmsg(m, int32(283545), v14-int32(-64))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L9
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(476291), int32(1665), int32(325376))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L9
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, _consts[166])) = v402
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L9
	} else {
		goto L103
	}
L103:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L9
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v14 + int32(1280)
	F_errmsg(m, int32(284810), v14+int32(48))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L9
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(476291), int32(1689), int32(325376))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L9
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L9
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v14 + int32(1280)
	F_errmsg(m, int32(284605), v14+int32(32))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L9
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(476291), int32(1696), int32(325376))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L9
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L9
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v14 + int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v14 + int32(1280)
	F_errmsg(m, int32(282611), v14+int32(16))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L9
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(476291), int32(1709), int32(325376))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SnapBuildXactNeedsSkip(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	return base.B2i32(base.Ui64(l1) < base.Ui64(v3))
}
