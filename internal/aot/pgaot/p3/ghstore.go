package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v510 int64
	_ = v510
	v2 = int32(0)
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = base.I32_wrap_i64(v16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == v2 {
		v36 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v36&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v23 == int32(0) {
		v36 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 != int32(7) {
		v36 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v29 != int32(17) {
		v36 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+32)))
	v36 = v32 ^ int32(1)
	goto L2
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = F_get_fn_opclass_options(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v45 = int32(16)
	goto L9
L9:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+18)))
	if v46 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	return int64(0)
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v45 = v44
	goto L9
L12:
	;
	return v510
L13:
	;
	v510 = base.I64_extend_i32_u(v480)
	goto L12
L14:
	;
	v468 = F_palloc(m, int32(24))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L10
	} else {
		goto L71
	}
L15:
	;
	v50 = v45 + int32(8)
	v51 = F_palloc(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399)+4)))
	if v400&int32(4) != 0 {
		goto L60
	} else {
		goto L61
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
	v59 = v51 + int32(8)
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	base.MemoryFill(m, v59, int32(0), v45)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v63 = F_hstoreUpgrade(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v67 = v65 & int32(268435455)
	if v67 == int32(0) {
		v457 = v51
		goto L14
	} else {
		goto L23
	}
L23:
	;
	v71 = v63 + int32(8)
	v72 = int32(3)
	v74 = v71 + v67<<(uint(v72)%32)
	v76 = v45 << (uint(v72) % 32)
	v88 = v2
	goto L24
L24:
	;
	v94 = v71 + v88<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v457 = v51
	goto L14
L26:
	;
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v110 = v95 & int32(1073741823)
	v111 = v74
	goto L26
L28:
	;
	goto L29
L29:
	;
	v100 = int32(1073741823)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v94-int32(4))))
	v106 = v104 & v100
	v110 = v95&v100 - v106
	v111 = v106 + v74
	goto L26
L30:
	;
	v112 = int32(-1)
	if v110 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v226 = int32(0)
	goto L32
L32:
	;
	v227 = base.I32_rem_u_s(v226, v76)
	v230 = v59 + int32(base.Ui32(v227)>>(uint(int32(3))%32))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v236 = v231 | int32(1)<<(uint(v227&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v236)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v238&int32(1073741824) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L33:
	;
	v226 = v194 ^ int32(-1)
	goto L32
L34:
	;
	v120 = v111
	v121 = v112
	v122 = int32(0)
	goto L37
L35:
	;
	v166 = v111
	v167 = v112
	goto L36
L36:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v189 = *(*int32)(unsafe.Add(mBase, uint32((v167^v181)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_compress[0])))
	v194 = v189 ^ int32(base.Ui32(v167)>>(uint(int32(8))%32))
	goto L33
L37:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v137 = int32(255)
	v139 = int32(2)
	v143 = *(*int32)(unsafe.Add(mBase, uint32((v135^v121)&v137<<(uint(v139)%32))+uint32(_c_F_ghstore_compress[0])))
	v144 = int32(8)
	v146 = v143 ^ int32(base.Ui32(v121)>>(uint(v144)%32))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32((v146^v147)&v137<<(uint(v139)%32))+uint32(_c_F_ghstore_compress[0])))
	v158 = v155 ^ int32(base.Ui32(v146)>>(uint(v144)%32))
	v160 = v120 + v139
	v162 = v122 + v139
	if v162 != v110&int32(-2) {
		v120 = v160
		v121 = v158
		v122 = v162
		goto L37
	} else {
		goto L39
	}
L38:
	;
	if v110&int32(1) == int32(0) {
		v194 = v158
		goto L33
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	v166 = v160
	v167 = v158
	goto L36
L41:
	;
	if v238 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	v397 = v88 + int32(1)
	if v397 != v67 {
		v88 = v397
		goto L24
	} else {
		goto L59
	}
L44:
	;
	if v253 != 0 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v253 = v238 & int32(1073741823)
	v254 = v74
	goto L44
L46:
	;
	goto L47
L47:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v249 = v247 & int32(1073741823)
	v253 = v238 - v249
	v254 = v249 + v74
	goto L44
L48:
	;
	v255 = int32(-1)
	if v253 != int32(1) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v369 = int32(0)
	goto L50
L50:
	;
	v370 = base.I32_rem_u_s(v369, v76)
	v373 = v59 + int32(base.Ui32(v370)>>(uint(int32(3))%32))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	v379 = v374 | int32(1)<<(uint(v370&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v373))) = uint8(v379)
	goto L43
L51:
	;
	v369 = v337 ^ int32(-1)
	goto L50
L52:
	;
	v263 = v254
	v264 = v255
	v265 = int32(0)
	goto L55
L53:
	;
	v309 = v254
	v310 = v255
	goto L54
L54:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32((v310^v324)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_compress[0])))
	v337 = v332 ^ int32(base.Ui32(v310)>>(uint(int32(8))%32))
	goto L51
L55:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263))))
	v280 = int32(255)
	v282 = int32(2)
	v286 = *(*int32)(unsafe.Add(mBase, uint32((v278^v264)&v280<<(uint(v282)%32))+uint32(_c_F_ghstore_compress[0])))
	v287 = int32(8)
	v289 = v286 ^ int32(base.Ui32(v264)>>(uint(v287)%32))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	v298 = *(*int32)(unsafe.Add(mBase, uint32((v289^v290)&v280<<(uint(v282)%32))+uint32(_c_F_ghstore_compress[0])))
	v301 = v298 ^ int32(base.Ui32(v289)>>(uint(v287)%32))
	v303 = v263 + v282
	v305 = v265 + v282
	if v305 != v253&int32(-2) {
		v263 = v303
		v264 = v301
		v265 = v305
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v253&int32(1) == int32(0) {
		v337 = v301
		goto L51
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v309 = v303
	v310 = v301
	goto L54
L59:
	;
	goto L25
L60:
	;
	v480 = v17
	goto L13
L61:
	;
	goto L62
L62:
	;
	v403 = int32(0)
	if v403 < v45 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v410 = v403
	goto L66
L64:
	;
	goto L65
L65:
	;
	v448 = F_palloc(m, int32(8))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L10
	} else {
		goto L70
	}
L66:
	;
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+(v399+int32(8))))))
	if v426 != int32(255) {
		v510 = v16 & int64(4294967295)
		goto L12
	} else {
		goto L68
	}
L67:
	;
	goto L65
L68:
	;
	v430 = v410 + int32(1)
	if v430 != v45 {
		v410 = v430
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v448))) = int64(17179869216)
	v457 = v448
	goto L14
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v468))) = base.I64_extend_i32_u(v457)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+8)) = v472
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v468)+12)) = v474
	v476 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+16)))
	v477 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v468)+18)) = uint8(v477)
	*(*uint16)(unsafe.Add(mBase, uint32(v468)+16)) = uint16(v476)
	v480 = v468
	goto L13
}
func F_ghstore_options(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_ghstore_options_0), int32(_a_F_ghstore_options_1), int32(16), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
