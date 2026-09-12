package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_shm_mq_get_sender(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l0, int32(515238), int32(261), int32(236957))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			return v15
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		return v15
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int64
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v415 int32
	_ = v415
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
	goto L25
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v23 = l2 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(l2) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(0)
	v144 = v6
	goto L1
L6:
	;
	v33 = v6
	v39 = v6
	v40 = v6
	goto L9
L7:
	;
	v63 = v6
	v69 = v6
	goto L8
L8:
	;
	if v23 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v44 = l1 + v33<<(uint(int32(3))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+28))
	v52 = v39 + v45 + v47 + v49 + v51
	v53 = int32(4)
	v54 = v33 + v53
	v56 = v40 + v53
	if v56 != l2&int32(2147483644) {
		v33 = v54
		v39 = v52
		v40 = v56
		goto L9
	} else {
		goto L11
	}
L10:
	;
	v63 = v54
	v69 = v52
	goto L8
L11:
	;
	goto L10
L12:
	;
	v77 = v63
	v79 = v6
	v83 = v69
	goto L15
L13:
	;
	v107 = v69
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v107
	if base.Ui32(int32(1073741823)) < base.Ui32(v107) {
		goto L2
	} else {
		goto L18
	}
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1+v77<<(uint(int32(3))%32))+4))
	v90 = v83 + v89
	v91 = int32(1)
	v94 = v79 + v91
	if v94 != v23 {
		v77 = v77 + v91
		v79 = v94
		v83 = v90
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v107 = v90
	goto L14
L17:
	;
	goto L16
L18:
	;
	v144 = v107
	goto L1
L19:
	;
	return int32(0)
L20:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v107
	F_errmsg(m, int32(362213), v17)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(515238), int32(384), int32(36813))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	m.G0 = v17 + int32(32)
	return v415
L25:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v163 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v415 = int32(1)
	goto L24
L27:
	;
	v169 = v162
	v170 = v162
	v173 = int32(0)
	goto L30
L28:
	;
	goto L29
L29:
	;
	v379 = F_shm_mq_send_bytes(m, l0, int32(4)-v162, v17+int32(28)+v162, l3, v17+int32(24))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L19
	} else {
		goto L83
	}
L30:
	;
	v180 = l1 + v173<<(uint(int32(3))%32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if base.Ui32(v181) <= base.Ui32(v169) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v326)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v326
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+36)))
	if v331 != 0 {
		v415 = int32(2)
		goto L24
	} else {
		goto L66
	}
L32:
	;
	goto L31
L33:
	;
	if base.Ui32(v303) < base.Ui32(v144) {
		v169 = v302
		v170 = v303
		v173 = v306
		goto L30
	} else {
		goto L65
	}
L34:
	;
	v184 = v173 + int32(1)
	if l2 <= v184 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v188 = v173 + int32(1)
	if l2 <= v188 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v302 = v169 - v181
	v303 = v170
	v306 = v184
	goto L33
L38:
	;
	v281 = F_shm_mq_send_bytes(m, l0, v275, v17+int32(16), l3, v17+int32(24))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L19
	} else {
		goto L60
	}
L39:
	;
	v248 = v181 - v169
	if v188 < l2 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	if base.Ui32(v169+int32(8)) <= base.Ui32(v181) {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v199 = v169
	v200 = int32(0)
	v201 = v181
	v203 = v173
	goto L42
L42:
	;
	v216 = v199
	v217 = v200
	goto L45
L44:
	;
	v240 = v216 - v201
	v242 = v203 + int32(1)
	if l2 <= v242 {
		v274 = v240
		v275 = v217
		v276 = v242
		goto L38
	} else {
		goto L49
	}
L45:
	;
	if base.Ui32(v201) <= base.Ui32(v216) {
		goto L44
	} else {
		goto L47
	}
L46:
	;
	v274 = v234
	v275 = int32(8)
	v276 = v203
	goto L38
L47:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1+v203<<(uint(int32(3))%32))))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229+v216))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17+int32(16)+v217))) = uint8(v231)
	v233 = int32(1)
	v234 = v216 + v233
	v236 = v217 + v233
	if v236 != int32(8) {
		v216 = v234
		v217 = v236
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1+v242<<(uint(int32(3))%32))+4))
	v199 = v240
	v200 = v217
	v201 = v247
	v203 = v242
	goto L42
L50:
	;
	v252 = v248 & int32(-8)
	goto L52
L51:
	;
	v252 = v248
	goto L52
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v257 = F_shm_mq_send_bytes(m, l0, v252, v253+v169, l3, v17+int32(24))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L19
	} else {
		goto L53
	}
L53:
	;
	if v257 == int32(2) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v261 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v261
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v261)
	v415 = int32(2)
	goto L24
L55:
	;
	goto L56
L56:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v268 = v266 + v267
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v268
	if v257 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v302 = v169 + v266
	v303 = v268
	v306 = v173
	goto L33
L58:
	;
	goto L59
L59:
	;
	v415 = int32(1)
	goto L24
L60:
	;
	if v281 == int32(2) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v285)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v285
	v415 = int32(2)
	goto L24
L62:
	;
	goto L63
L63:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v292 = v290 + v291
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v292
	if v281 == int32(0) {
		v302 = v274
		v303 = v292
		v306 = v276
		goto L33
	} else {
		goto L64
	}
L64:
	;
	v415 = int32(1)
	goto L24
L65:
	;
	goto L32
L66:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v332 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if l4 != 0 {
		goto L76
	} else {
		goto L77
	}
L68:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v354 = v335
	goto L67
L69:
	;
	goto L70
L70:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(1)
	if v336 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_s_lock(m, v19, int32(515238), int32(526), int32(36813))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L19
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v344 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v344
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v347 == v344 {
		v354 = v344
		goto L67
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v350 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)) = uint8(v350)
	v354 = v347
	goto L67
L76:
	;
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v361 + base.I64_extend_i32_u(v355)
	if v354 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	if base.Ui32(int32(base.Ui32(v356)>>(uint(int32(2))%32))) < base.Ui32(v355) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v415 = int32(0)
	goto L24
L79:
	;
	F_SetLatch(m, v354+int32(20))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L19
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v369 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v369
	v415 = v369
	goto L24
L82:
	;
	goto L81
L83:
	;
	if v379 == int32(2) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v383 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v383)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v383
	v415 = int32(2)
	goto L24
L85:
	;
	goto L86
L86:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v390 = v388 + v389
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v390
	if base.Ui32(int32(4)) <= base.Ui32(v390) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v394 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v394)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	goto L89
L88:
	;
	goto L89
L89:
	;
	if v379 == int32(0) {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	goto L26
}
func F_shm_mq_set_receiver(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v4 != 0 {
		F_s_lock(m, l0, int32(515238), int32(210), int32(224249))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v15 != 0 {
				F_SetLatch(m, v15+int32(20))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v15 != 0 {
			F_SetLatch(m, v15+int32(20))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
