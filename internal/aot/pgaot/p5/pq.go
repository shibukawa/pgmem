package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	if int32(0) <= l2 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if l2 <= v7-v8 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if l2 != 0 {
				v30 = F__emscripten_memcpy_bulkmem(m, l1, v28+v8, l2)
				mBase = m.M
			} else {
			}
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32 + l2
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_errmsg(m, int32(404899), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						F_errfinish(m, int32(493920), int32(533), int32(158758))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errmsg(m, int32(404899), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(493920), int32(533), int32(158758))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_pq_endmessage_reuse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+12)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v2, v3, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_pq_getkeepalivesidle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v22 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v22 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
			if v13 != 0 {
				v22 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
				if v14 != 0 {
					v22 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(384))))
					v22 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v22
}
func F_pq_getmessage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v449 int32
	_ = v449
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v22 = v3
	v23 = v3
	v24 = v3
	v25 = int32(-1)
	v26 = v3
	v27 = v3
	v28 = v3
	v31 = v16
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v25 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	goto L2
L5:
	;
	v483 = int32(m.ExcTag)
	v484 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v483 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v215
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	v372 = int32(0)
	if v209 == v372 {
		v419 = v372
		goto L72
	} else {
		goto L73
	}
L7:
	;
	m.G0 = v16 + int32(32)
	return v352
L8:
	;
	v342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v342)
	v352 = v342
	goto L7
L9:
	;
	v35 = v31 - int32(16)
	m.G0 = v35
	v38 = v35 - int32(160)
	m.G0 = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v46)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v46
	goto L12
L10:
	;
	v209 = v22
	v210 = v23
	v211 = v24
	v213 = v26
	v214 = v27
	v215 = v28
	v218 = v31
	goto L11
L11:
	;
	if v210 != 0 {
		goto L6
	} else {
		goto L46
	}
L12:
	;
	v63 = v35
	v64 = int32(4)
	goto L13
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v67 <= v69 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v132 = int32(24)
	v134 = int32(65280)
	v136 = int32(8)
	v146 = v131<<(uint(v132)%32) | v131&v134<<(uint(v136)%32) | (int32(base.Ui32(v131)>>(uint(v136)%32))&v134 | int32(base.Ui32(v131)>>(uint(v132)%32)))
	if base.B2i32(int32(4) <= v146)&base.B2i32(v146 <= l1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	v76 = F_pq_recvbuf(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v121 = v67 - v69
	if base.Ui32(v121) < base.Ui32(v64) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if v76 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	v87 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v89 = int32(-1)
	if v87 == int32(0) {
		v352 = v89
		goto L7
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errcode(m, int32(16908800))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errmsg(m, int32(420732), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errfinish(m, int32(497558), int32(1217), int32(404330))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v352 = v89
	goto L7
L25:
	;
	v123 = v121
	goto L27
L26:
	;
	v123 = v64
	goto L27
L27:
	;
	if v123 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[380])) = v123 + v69
	v130 = v64 - v123
	if v130 != 0 {
		v63 = v123 + v125
		v64 = v130
		goto L13
	} else {
		goto L32
	}
L29:
	;
	v124 = F__emscripten_memcpy_bulkmem(m, v63, v69+int32(4414656), v123)
	mBase = m.M
	v125 = v124
	goto L31
L30:
	;
	v125 = v63
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L14
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	v160 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v131 == int32(67108864) {
		goto L8
	} else {
		goto L41
	}
L36:
	;
	v162 = int32(-1)
	if v160 == int32(0) {
		v352 = v162
		goto L7
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errcode(m, int32(16908800))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errmsg(m, int32(321317), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v35
	F_errfinish(m, int32(497558), int32(1227), int32(404330))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		v482 = v38
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v352 = v162
	goto L7
L41:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v200 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v16 + int32(8)
	goto L45
L43:
	;
	v209 = v146 - int32(4)
	v210 = int32(0)
	v211 = v38
	v213 = v35
	v214 = v200
	v215 = v198
	v218 = v38
	goto L11
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_enlargeStringInfo(m, l0, v209)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v215
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v214
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v209 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v243 = v232
	v244 = v209
	goto L51
L49:
	;
	v322 = v232
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v209
	v327 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v209+v322))) = uint8(v327)
	goto L8
L51:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	v249 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v247 <= v249 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v322 = v311
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	v256 = F_pq_recvbuf(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v301 = v247 - v249
	if base.Ui32(v301) < base.Ui32(v244) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	if v256 == int32(0) {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	v267 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v269 = int32(-1)
	if v267 == int32(0) {
		v352 = v269
		goto L7
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errcode(m, int32(16908800))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errmsg(m, int32(96552), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errfinish(m, int32(497558), int32(1262), int32(404330))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v352 = v269
	goto L7
L63:
	;
	v303 = v301
	goto L65
L64:
	;
	v303 = v244
	goto L65
L65:
	;
	if v303 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[380])) = v303 + v249
	v310 = v244 - v303
	if v310 != 0 {
		v243 = v303 + v305
		v244 = v310
		goto L51
	} else {
		goto L70
	}
L67:
	;
	v304 = F__emscripten_memcpy_bulkmem(m, v243, v249+int32(4414656), v303)
	mBase = m.M
	v305 = v304
	goto L69
L68:
	;
	v305 = v243
	goto L69
L69:
	;
	goto L66
L70:
	;
	goto L52
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	v462 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v462)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_pg_re_throw(m)
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L91
	}
L72:
	;
	if v419 != int32(-1) {
		goto L71
	} else {
		goto L85
	}
L73:
	;
	v386 = v209
	goto L74
L74:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	v391 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v389 <= v391 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v419 = int32(0)
	goto L72
L76:
	;
	v393 = F_pq_recvbuf(m)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v399 = v389 - v391
	if base.Ui32(v399) < base.Ui32(v386) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	if v393 == int32(0) {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	v419 = int32(-1)
	goto L72
L81:
	;
	v401 = v399
	goto L83
L82:
	;
	v401 = v386
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, _consts[380])) = v401 + v391
	v404 = v386 - v401
	if v404 != 0 {
		v386 = v404
		goto L74
	} else {
		goto L84
	}
L84:
	;
	goto L75
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	v429 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L86
	}
L86:
	;
	if v429 == int32(0) {
		goto L71
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errcode(m, int32(16908800))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errmsg(m, int32(96552), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v214
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v213
	F_errfinish(m, int32(497558), int32(1249), int32(404330))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		v482 = v218
		goto L5
	} else {
		goto L90
	}
L90:
	;
	goto L71
L91:
	;
	goto L4
L92:
	;
	v488 = int32(v484)
	m.G0 = v482
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v488)+4))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	if v16+int32(8) == v495 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	m.ExcPending = 1
	goto L101
L94:
	;
	if v498 != 0 {
		goto L98
	} else {
		goto L99
	}
L95:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	v498 = v497
	goto L97
L96:
	;
	v498 = int32(0)
	goto L97
L97:
	;
	goto L94
L98:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v22 = v501
	v23 = v490
	v24 = v500
	v25 = v498
	v26 = v499
	v27 = v502
	v28 = v503
	v31 = v482
	goto L1
L99:
	;
	goto L100
L100:
	;
	F___wasm_longjmp(m, v491, v490)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	return int32(0)
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_putemptymessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = m.T0[v6].(func(*base.Module, int32, int32, int32) int32)(m, l0, v2, v2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_pq_puttextmessage(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = int32(522613)
	v6 = F_strlen(m, v4)
	mBase = m.M
	v7 = F_pg_server_to_client(m, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[219]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		if v7 != int32(522613) {
			v15 = F_strlen(m, v7)
			mBase = m.M
			v18 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), v7, v15+int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v26 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), int32(522613), v6+int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_sendstring(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v5 = F_strlen(m, l1)
	v6 = F_pg_server_to_client(m, l1, v5)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 != v6 {
			v9 = F_strlen(m, v6)
			F_appendBinaryStringInfoNT(m, l0, v6, v9+int32(1))
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_pfree(m, v6)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_appendBinaryStringInfoNT(m, l0, l1, v5+int32(1))
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_settcpusertimeout(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+412))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
					} else {
					}
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+412)) = v50
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v39 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v39 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+412))
							if v30 != 0 {
								v39 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
								if v31 != 0 {
									v39 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(396))))
									v39 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v39 {
						if l0 == int32(0) {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v47
						} else {
						}
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+412)) = v50
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
func F_pq_startmsgread(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _consts[378])))
	if v2 != 0 {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_errmsg(m, int32(68075), int32(0))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					F_errfinish(m, int32(497558), int32(1151), int32(464512))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v20 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[378])) = uint8(v20)
		return
	}
}
