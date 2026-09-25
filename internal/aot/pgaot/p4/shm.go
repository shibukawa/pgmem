package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_shm_mq_get_sender(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_shm_mq_get_sender_0), int32(261), int32(_a_F_shm_mq_get_sender_1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v14 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v14))
			return v13
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v14))
		return v13
	}
}
func F_shm_mq_send(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
	v16 = F_shm_mq_sendv(m, l0, v9+int32(8), int32(1), l3, l4)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v9 + int32(16)
		return v16
	}
}
func F_shm_mq_sendv(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v378 int64
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v468 int32
	_ = v468
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 < l2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	goto L24
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L18
	} else {
		goto L19
	}
L3:
	;
	v23 = l2 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(0)
	v147 = v6
	goto L1
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v110
	if base.Ui32(int32(1073741823)) < base.Ui32(v110) {
		goto L2
	} else {
		goto L17
	}
L7:
	;
	v33 = v6
	v34 = v6
	v40 = v6
	goto L10
L8:
	;
	v65 = v6
	v72 = v6
	goto L9
L9:
	;
	v79 = v65
	v85 = v6
	v86 = v72
	goto L14
L10:
	;
	v44 = l1 + v33<<(uint(int32(3))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v52 = v45 + (v46 + (v47 + (v48 + v40)))
	v53 = int32(4)
	v54 = v33 + v53
	v56 = v34 + v53
	if v56 != l2&int32(2147483644) {
		v33 = v54
		v34 = v56
		v40 = v52
		goto L10
	} else {
		goto L12
	}
L11:
	;
	if v23 == int32(0) {
		v110 = v52
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v65 = v54
	v72 = v52
	goto L9
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1+v79<<(uint(int32(3))%32))+4))
	v92 = v91 + v86
	v93 = int32(1)
	v96 = v85 + v93
	if v96 != v23 {
		v79 = v79 + v93
		v85 = v96
		v86 = v92
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v110 = v92
	goto L6
L16:
	;
	goto L15
L17:
	;
	v147 = v110
	goto L1
L18:
	;
	return int32(0)
L19:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v110
	F_errmsg(m, int32(_a_F_shm_mq_sendv_0), v17)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_shm_mq_sendv_1), int32(384), int32(_a_F_shm_mq_sendv_2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	m.G0 = v17 + int32(32)
	return v468
L24:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v165 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v468 = int32(1)
	goto L23
L26:
	;
	v171 = v164
	v172 = v164
	v174 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	v441 = F_shm_mq_send_bytes(m, l0, int32(4)-v164, v17+int32(28)+v164, l3, v17+int32(24))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L18
	} else {
		goto L97
	}
L29:
	;
	v182 = l1 + v174<<(uint(int32(3))%32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if base.Ui32(v183) <= base.Ui32(v171) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v331)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v331
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)))
	if v335 != 0 {
		goto L65
	} else {
		goto L66
	}
L31:
	;
	goto L30
L32:
	;
	if base.Ui32(v308) < base.Ui32(v147) {
		v171 = v307
		v172 = v308
		v174 = v310
		goto L29
	} else {
		goto L64
	}
L33:
	;
	v186 = v174 + int32(1)
	if l2 <= v186 {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v193 = v174 + int32(1)
	if base.B2i32(base.Ui32(v171+int32(8)) <= base.Ui32(v183))|base.B2i32(l2 <= v193) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v307 = v171 - v183
	v308 = v172
	v310 = v186
	goto L32
L37:
	;
	v286 = F_shm_mq_send_bytes(m, l0, v280, v17+int32(16), l3, v17+int32(24))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L18
	} else {
		goto L59
	}
L38:
	;
	v204 = v171
	v205 = int32(0)
	v207 = v174
	v210 = v183
	goto L41
L39:
	;
	goto L40
L40:
	;
	v253 = v183 - v171
	if v193 < l2 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	v221 = v204
	v222 = v205
	goto L44
L43:
	;
	v245 = v221 - v210
	v247 = v207 + int32(1)
	if l2 <= v247 {
		v279 = v245
		v280 = v222
		v281 = v247
		goto L37
	} else {
		goto L48
	}
L44:
	;
	if base.Ui32(v210) <= base.Ui32(v221) {
		goto L43
	} else {
		goto L46
	}
L45:
	;
	v279 = v239
	v280 = int32(8)
	v281 = v207
	goto L37
L46:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1+v207<<(uint(int32(3))%32))))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234+v221))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(16)+v222))) = uint8(v236)
	v238 = int32(1)
	v239 = v221 + v238
	v241 = v222 + v238
	if v241 != int32(8) {
		v221 = v239
		v222 = v241
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1+v247<<(uint(int32(3))%32))+4))
	v204 = v245
	v205 = v222
	v207 = v247
	v210 = v252
	goto L41
L49:
	;
	v257 = v253 & int32(-8)
	goto L51
L50:
	;
	v257 = v253
	goto L51
L51:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v262 = F_shm_mq_send_bytes(m, l0, v257, v258+v171, l3, v17+int32(24))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	if v262 == int32(2) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v266)
	v468 = int32(2)
	goto L23
L54:
	;
	goto L55
L55:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v273 = v271 + v272
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v273
	if v262 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v307 = v171 + v271
	v308 = v273
	v310 = v174
	goto L32
L57:
	;
	goto L58
L58:
	;
	v468 = int32(1)
	goto L23
L59:
	;
	if v286 == int32(2) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v290 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v290
	v468 = int32(2)
	goto L23
L61:
	;
	goto L62
L62:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v297 = v295 + v296
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v297
	if v286 == int32(0) {
		v307 = v279
		v308 = v297
		v310 = v281
		goto L32
	} else {
		goto L63
	}
L63:
	;
	v468 = int32(1)
	goto L23
L64:
	;
	goto L31
L65:
	;
	v468 = int32(2)
	goto L23
L66:
	;
	goto L67
L67:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v337 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l4 != 0 {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v359 = v340
	goto L68
L70:
	;
	goto L71
L71:
	;
	v343 = base.AtomicRmwXchg32(m, v19, int32(0), int32(1))
	if v343 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_s_lock(m, v19, int32(_a_F_shm_mq_sendv_1), int32(526), int32(_a_F_shm_mq_sendv_2))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L18
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v350 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v19))), uint32(v350))
	if v349 == v350 {
		v359 = v350
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v356 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v356)
	v359 = v349
	goto L68
L77:
	;
	v366 = int32(0)
	v370 = base.AtomicRmwOr32(m, v366, int32(_a_F_shm_mq_sendv_3), v366)
	v372 = int64(0)
	v374 = int32(24)
	v375 = base.AtomicRmwCmpxchg64(m, v19, v374, v372, v372)
	v378 = base.AtomicRmwXchg64(m, v19, v374, base.I64_extend_i32_u(v360)+v375)
	if v359 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if base.Ui32(int32(base.Ui32(v361)>>(uint(int32(2))%32))) < base.Ui32(v360) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v468 = int32(0)
	goto L23
L80:
	;
	v380 = v359 + int32(20)
	v381 = int32(0)
	v384 = base.AtomicRmwOr32(m, v381, int32(_a_F_shm_mq_sendv_4), v381)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if v385 != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	v468 = v366
	goto L23
L83:
	;
	goto L82
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = int32(1)
	v388 = int32(0)
	v391 = base.AtomicRmwOr32(m, v388, int32(_a_F_shm_mq_sendv_4), v388)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	if v392 == v388 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v380)+12))
	if v395 == int32(0) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[0]))
	if v399 == v395 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v401 = m.G0
	v403 = v401 - int32(16)
	m.G0 = v403
	v406 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[1]))
	if v406 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v429 = F_pgmem_kill(m, v395, int32(23))
	mBase = m.M
	goto L84
L91:
	;
	m.G0 = v403 + int32(16)
	goto L83
L92:
	;
	v409 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v403)+15)) = uint8(v409)
	goto L93
L93:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[2]))
	v417 = F_write(m, v413, v403+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v417 {
		goto L91
	} else {
		goto L95
	}
L94:
	;
	goto L91
L95:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[3]))
	if v421 == int32(27) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	if v441 == int32(2) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v445 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v445)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v445
	v468 = int32(2)
	goto L23
L99:
	;
	goto L100
L100:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v452 = v450 + v451
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v452
	if base.Ui32(int32(4)) <= base.Ui32(v452) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v456)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L103
L102:
	;
	goto L103
L103:
	;
	if v441 == int32(0) {
		goto L24
	} else {
		goto L104
	}
L104:
	;
	goto L25
}
func F_shm_mq_set_receiver(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_shm_mq_set_receiver_0), int32(210), int32(_a_F_shm_mq_set_receiver_1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v13 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v13))
	if v12 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v17 = v12 + int32(20)
	v18 = int32(0)
	v21 = base.AtomicRmwOr32(m, v18, int32(_a_F_shm_mq_set_receiver_2), v18)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	goto L8
L10:
	;
	goto L9
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(1)
	v25 = int32(0)
	v28 = base.AtomicRmwOr32(m, v25, int32(_a_F_shm_mq_set_receiver_2), v25)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v29 == v25 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v32 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[0]))
	if v36 == v32 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v38 = m.G0
	v40 = v38 - int32(16)
	m.G0 = v40
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[1]))
	if v43 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v66 = F_pgmem_kill(m, v32, int32(23))
	mBase = m.M
	goto L10
L17:
	;
	m.G0 = v40 + int32(16)
	goto L9
L18:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v40)+15)) = uint8(v46)
	goto L19
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[2]))
	v54 = F_write(m, v50, v40+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v54 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[3]))
	if v58 == int32(27) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
