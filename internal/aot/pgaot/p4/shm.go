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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_shm_mq_get_sender_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v12 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v12))
			return v11
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v12))
		return v11
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
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int64
	_ = v370
	var v372 int32
	_ = v372
	var v373 int64
	_ = v373
	var v376 int64
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v466 int32
	_ = v466
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
	F_errfinish(m, int32(_a_F_shm_mq_sendv_1), int32(386), int32(_a_F_shm_mq_sendv_2))
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
	return v466
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
	v466 = int32(1)
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
	v439 = F_shm_mq_send_bytes(m, l0, int32(4)-v164, v17+int32(28)+v164, l3, v17+int32(24))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
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
	v466 = int32(2)
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
	v466 = int32(1)
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
	v466 = int32(2)
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
	v466 = int32(1)
	goto L23
L64:
	;
	goto L31
L65:
	;
	v466 = int32(2)
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
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l4 != 0 {
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v357 = v340
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
	F_s_lock(m, v19, int32(_a_F_shm_mq_sendv_3))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L18
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v348 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v19))), uint32(v348))
	if v347 == v348 {
		v357 = v348
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v354 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v354)
	v357 = v347
	goto L68
L77:
	;
	v364 = int32(0)
	v368 = base.AtomicRmwOr32(m, v364, int32(_a_F_shm_mq_sendv_4), v364)
	v370 = int64(0)
	v372 = int32(24)
	v373 = base.AtomicRmwCmpxchg64(m, v19, v372, v370, v370)
	v376 = base.AtomicRmwXchg64(m, v19, v372, base.I64_extend_i32_u(v358)+v373)
	if v357 != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if base.Ui32(int32(base.Ui32(v359)>>(uint(int32(2))%32))) < base.Ui32(v358) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v466 = int32(0)
	goto L23
L80:
	;
	v378 = v357 + int32(316)
	v379 = int32(0)
	v382 = base.AtomicRmwOr32(m, v379, int32(_a_F_shm_mq_sendv_5), v379)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if v383 != 0 {
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
	v466 = v364
	goto L23
L83:
	;
	goto L82
L84:
	;
	goto L83
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v378))) = int32(1)
	v386 = int32(0)
	v389 = base.AtomicRmwOr32(m, v386, int32(_a_F_shm_mq_sendv_5), v386)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v390 == v386 {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	if v393 == int32(0) {
		goto L84
	} else {
		goto L87
	}
L87:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[0]))
	if v397 == v393 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v399 = m.G0
	v401 = v399 - int32(16)
	m.G0 = v401
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[1]))
	if v404 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v427 = F_pgmem_kill(m, v393, int32(23))
	mBase = m.M
	goto L84
L91:
	;
	m.G0 = v401 + int32(16)
	goto L83
L92:
	;
	v407 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+15)) = uint8(v407)
	goto L93
L93:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[2]))
	v415 = F_write(m, v411, v401+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v415 {
		goto L91
	} else {
		goto L95
	}
L94:
	;
	goto L91
L95:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_sendv[3]))
	if v419 == int32(27) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	if v439 == int32(2) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v443)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v443
	v466 = int32(2)
	goto L23
L99:
	;
	goto L100
L100:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v450 = v448 + v449
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v450
	if base.Ui32(int32(4)) <= base.Ui32(v450) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v454)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L103
L102:
	;
	goto L103
L103:
	;
	if v439 == int32(0) {
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_s_lock(m, l0, int32(_a_F_shm_mq_set_receiver_0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
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
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v11))
	if v10 != 0 {
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
	v15 = v10 + int32(316)
	v16 = int32(0)
	v19 = base.AtomicRmwOr32(m, v16, int32(_a_F_shm_mq_set_receiver_1), v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v20 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
	v23 = int32(0)
	v26 = base.AtomicRmwOr32(m, v23, int32(_a_F_shm_mq_set_receiver_1), v23)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v27 == v23 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v30 == int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[0]))
	if v34 == v30 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v36 = m.G0
	v38 = v36 - int32(16)
	m.G0 = v38
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[1]))
	if v41 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v64 = F_pgmem_kill(m, v30, int32(23))
	mBase = m.M
	goto L10
L17:
	;
	m.G0 = v38 + int32(16)
	goto L9
L18:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+15)) = uint8(v44)
	goto L19
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[2]))
	v52 = F_write(m, v48, v38+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v52 {
		goto L17
	} else {
		goto L21
	}
L20:
	;
	goto L17
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_shm_mq_set_receiver[3]))
	if v56 == int32(27) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
