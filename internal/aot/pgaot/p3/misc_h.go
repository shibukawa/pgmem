package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_HaveRegisteredOrActiveSnapshot(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	if v3 != 0 {
		return int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _consts[1198]))
		v9 = *(*int32)(unsafe.Add(mBase, _consts[1195]))
		if v9 == int32(0) {
			return base.B2i32(v7 != int32(0))
		} else {
			if v7 == int32(0) {
				return base.B2i32(v7 != int32(0))
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if v14 != 0 {
					return base.B2i32(v7 != int32(0))
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_hamming_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			if v16 != v17 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v27
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v26
						F_errmsg(m, int32(54318), v6)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(512331), int32(39), int32(154322))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
				v41 = int32(8)
				v49 = *(*int32)(unsafe.Add(mBase, _consts[1365]))
				v50 = m.T0[v49].(func(*base.Module, int32, int32, int32, int64) int64)(m, int32(base.Ui32(v38)>>(uint(int32(2))%32))-v41, v9+v41, v14+v41, int64(0))
				mBase = m.M
				v52 = F_Float8GetDatum(m, base.F64_convert_i64_u(v50))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v52
				}
			}
		}
	}
}
func F_handle_pm_child_exit_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[657])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_handle_pm_reload_request_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, _consts[656])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	F_SetLatch(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_handle_streamed_transaction(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v76 int64
	_ = v76
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	v1 = l0
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	v18 = *(*int32)(unsafe.Add(mBase, _consts[588]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v19 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v384
L2:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v46 = F_pq_getmsgint(m, l1, int32(4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L18
	}
L3:
	;
	v26 = F_pa_find_worker(m, v16)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v22 != int32(3) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v40 = int32(0)
	v41 = int32(4)
	goto L2
L6:
	;
	return int32(0)
L7:
	;
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+12)))
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v34 = int32(0)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[693])))
	if v36 == v34 {
		v384 = v34
		goto L1
	} else {
		goto L14
	}
L11:
	;
	v33 = int32(3)
	goto L13
L12:
	;
	v33 = int32(2)
	goto L13
L13:
	;
	v40 = v26
	v41 = v33
	goto L2
L14:
	;
	v40 = v34
	v41 = int32(1)
	goto L2
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v1)
	v352 = int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v353 - v354 + v352
	v360 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	F_BufFileWrite(m, v360, v13, int32(4))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L6
	} else {
		goto L86
	}
L16:
	;
	v310 = int32(4446460)
	v311 = *(*int32)(unsafe.Add(mBase, _consts[695]))
	v312 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v305+v311<<(uint(v312)%32)))) = v46
	v317 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	v319 = *(*int32)(unsafe.Add(mBase, _consts[695]))
	v322 = v305 + v319<<(uint(v312)%32)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v317)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v322+v312))) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v317)+32))
	v330 = int64(*(*int32)(unsafe.Add(mBase, uint32(v317)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v322+int32(8)))) = v329 + v330
	goto L85
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[696])) = int32(128)
	v289 = int32(4536272)
	v290 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v293 = *(*int32)(unsafe.Add(mBase, _consts[697]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v293
	v296 = F_palloc(m, int32(2048))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L6
	} else {
		goto L84
	}
L18:
	;
	if v46 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	switch v41 - int32(2) {
	case 0:
		goto L24
	case 1:
		goto L23
	case 2:
		goto L22
	default:
		goto L25
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L6
	} else {
		goto L80
	}
L22:
	;
	v143 = int32(4446452)
	v145 = *(*int32)(unsafe.Add(mBase, _consts[698]))
	*(*int32)(unsafe.Add(mBase, _consts[698])) = v145 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	v151 = m.G0
	v153 = v151 - int32(96)
	m.G0 = v153
	if v150 == v46 {
		goto L43
	} else {
		goto L44
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1)
	v113 = v43 - v42
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v113 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	F_BufFileWrite(m, v118, v13+int32(8), int32(4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L40
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v102 = F_pa_send_data(m, v40, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L35
	}
L25:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[692]))
	if v51 == v46 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[699]))
	if v54 == v46 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[700]))
	*(*int32)(unsafe.Add(mBase, _consts[699])) = v46
	v61 = *(*int32)(unsafe.Add(mBase, _consts[695]))
	if v61 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v76 = base.I64_extend_i32_u(v61)
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v57-int32(16)+base.I32_wrap_i64(v76)<<(uint(int32(4))%32))))
	if v81 == v46 {
		goto L15
	} else {
		goto L31
	}
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[696]))
	if v61 != v90 {
		v305 = v57
		goto L16
	} else {
		goto L33
	}
L31:
	;
	if base.B2i32(v76 < int64(2)) == int32(0) {
		v76 = v76 - int64(1)
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[696])) = v61 << (uint(int32(1)) % 32)
	v98 = F_repalloc(m, v57, v61<<(uint(int32(5))%32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v305 = v98
	goto L16
L35:
	;
	if v102 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v384 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
	goto L1
L37:
	;
	goto L38
L38:
	;
	F_pa_switch_to_partial_serialize(m, v40, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	F_BufFileWrite(m, v125, v13+int32(15), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v113
	v133 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	F_BufFileWrite(m, v133, v42+v44, v113)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v384 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
	goto L1
L43:
	;
	m.G0 = v153 + int32(96)
	v384 = int32(0)
	goto L1
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v158 = int32(0)
	if v157 == v158 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v196 != 0 {
		goto L43
	} else {
		goto L58
	}
L46:
	;
	v196 = int32(0)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v164 <= int32(0) {
		v189 = v158
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v196 = v189
	goto L45
L50:
	;
	v167 = int32(0)
	if v167 < v164 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v170 = v164
	goto L53
L52:
	;
	v170 = v167
	goto L53
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v173 = int32(0)
	goto L54
L54:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173<<(uint(int32(2))%32))))
	v182 = base.B2i32(v181 == v46)
	if v181 == v46 {
		v189 = v182
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v189 = v182
	goto L49
L56:
	;
	v184 = v173 + int32(1)
	if v184 != v170 {
		v173 = v184
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _consts[675]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v46
	v208 = F_pg_snprintf(m, v153+int32(32), int32(64), int32(39423), v153+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v212 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	if v212 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v153 + int32(32)
	F_errmsg_internal(m, int32(225090), v153)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+24))
	goto L66
L64:
	;
	F_errfinish(m, int32(506954), int32(1380), int32(153137))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v227)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	goto L70
L68:
	;
	goto L69
L69:
	;
	F_DefineSavepoint(m, v153+int32(32))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L77
	}
L70:
	;
	if base.B2i32(v234 == int32(2)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L6
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	F_BeginTransactionBlock(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L6
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	goto L69
L77:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v251 = int32(4536272)
	v252 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v255 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v255
	v258 = *(*int32)(unsafe.Add(mBase, _consts[701]))
	v259 = F_lappend_xid(m, v258, v46)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v252
	*(*int32)(unsafe.Add(mBase, _consts[701])) = v259
	goto L43
L80:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errmsg_internal(m, int32(262234), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(506994), int32(584), int32(261615))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v290
	v305 = v296
	goto L16
L85:
	;
	*(*int32)(unsafe.Add(mBase, _consts[700])) = v305
	v335 = int32(4446460)
	v337 = *(*int32)(unsafe.Add(mBase, _consts[695]))
	*(*int32)(unsafe.Add(mBase, _consts[695])) = v337 + int32(1)
	goto L15
L86:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	F_BufFileWrite(m, v365, v13+int32(7), int32(1))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v373 = v371 - v372
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v373
	v376 = *(*int32)(unsafe.Add(mBase, _consts[694]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_BufFileWrite(m, v376, v372+v377, v373)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v384 = v352
	goto L1
}
func F_has_largeobject_privilege_name_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_get_role_oid_or_public(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = F_convert_any_priv_string(m, v13, int32(1673888))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v16&int64(4) == int64(0) {
					v23 = *(*int32)(unsafe.Add(mBase, _consts[113]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					v25 = v24
				} else {
					v25 = int32(0)
				}
				v26 = F_LargeObjectExistsWithSnapshot(m, v11, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, _consts[565])))
						if v30 != 0 {
							v39 = int32(1)
							return v39
						} else {
							v31 = F_pg_largeobject_aclcheck_snapshot(m, v11, v7, v16, v25)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v31 == int32(0))
							}
						}
					} else {
						v36 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
						v39 = int32(0)
						return v39
					}
				}
			}
		}
	}
}
func F_has_useful_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v5 != 0 {
		v11 = v4
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
		if v6 != 0 {
			v11 = v4
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
			if v7 != 0 {
				v11 = v4
			} else {
				v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				v11 = base.B2i32(v8 != int32(0))
			}
		}
	}
	return v11
}
func F_hashagg_spill_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l3
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v91 = F_ExecFetchSlotMinimalTuple(m, v85, v12+int32(11))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L22
	}
L2:
	;
	v85 = l2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+268))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
	if v18 < v17 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_slot_getsomeattrs_int(m, l2, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	m.T0[v23].(func(*base.Module, int32))(m, v16)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) < v27 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v35 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v75 = v73 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)) = uint16(v78)
	goto L21
L14:
	;
	v40 = int32(1)
	v42 = v35 + v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v44 = F_bms_is_member(m, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	if v44 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = v35 << (uint(int32(2)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50+v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v47+v48))) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v35))))
	v57 = v56
	goto L19
L18:
	;
	v57 = v40
	goto L19
L19:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v58+v35))) = uint8(v57)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v42 < v62 {
		v35 = v42
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v85 = v16
	goto L1
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v94 <= int32(31) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v100 = int32(base.Ui32(v97&l3) >> (uint(v94) % 32))
	goto L25
L24:
	;
	v100 = int32(0)
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v104 = v101 + v100<<(uint(int32(3))%32)
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v104)))
	*(*int64)(unsafe.Add(mBase, uint32(v104))) = v105 + int64(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v110 = int32(24)
	v112 = v109 + v100*v110
	v117 = int32(711645284)
	v120 = l3 - int32(1636608428) ^ v117 - int32(1455628627)
	v125 = v120 ^ int32(-1636608428) - base.I32_rotl(v120, int32(25))
	v130 = v125 ^ v117 - base.I32_rotl(v125, int32(16))
	v134 = v130 ^ v120 - base.I32_rotl(v130, int32(4))
	v138 = v134 ^ v125 - base.I32_rotl(v134, int32(14))
	v142 = v138 ^ v130 - base.I32_rotl(v138, v110)
	goto L26
L26:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v147 = int32(32) - v146
	v148 = v142 << (uint(v146) % 32)
	if v148 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v100<<(uint(int32(2))%32))))
	F_LogicalTapeWrite(m, v179, v12+int32(12), int32(4))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L37
	}
L28:
	;
	v155 = int32(32) - (base.I32_clz(v148) ^ int32(31))
	v156 = int32(255)
	if base.Ui32(v147&v156) < base.Ui32(v155&v156) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v165 = v147 + int32(1)
	goto L30
L30:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	v168 = v166 + int32(base.Ui32(v142)>>(uint(v147)%32))
	v170 = v165 & int32(255)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	if base.Ui32(v171) < base.Ui32(v170) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v161 = v147 + int32(1)
	goto L33
L32:
	;
	v161 = v155
	goto L33
L33:
	;
	v165 = v161
	goto L30
L34:
	;
	v173 = v170
	goto L36
L35:
	;
	v173 = v171
	goto L36
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v168))) = uint8(v173)
	goto L27
L37:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	F_LogicalTapeWrite(m, v179, v91, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v188 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pfree(m, v91)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	m.G0 = v12 + int32(16)
	return
L42:
	;
	goto L41
}
func F_hashboolextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int64(0) {
		v14 = int32(-1636608428)
		v52 = v14
		v54 = v14
		v57 = v14
	} else {
		v17 = base.I32_wrap_i64(v6)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v28 = v17 - v22 - int32(1636608428) ^ base.I32_rotl(v22, int32(6))
		v30 = v17 + int32(1021750440)
		v31 = v22 + v30
		v32 = v28 + v31
		v36 = v30 - v28 ^ base.I32_rotl(v28, int32(8))
		v40 = v31 - v36 ^ base.I32_rotl(v36, int32(16))
		v44 = v32 - v40 ^ base.I32_rotl(v40, int32(19))
		v45 = v36 + v32
		v46 = v40 + v45
		v52 = v44 + v46
		v54 = v46
		v57 = v45 - v44 ^ base.I32_rotl(v44, int32(4))
	}
	v59 = int32(14)
	v61 = v52 ^ v57 - base.I32_rotl(v52, v59)
	v66 = v61 ^ (base.B2i32(v2 != int32(0)) + v54) - base.I32_rotl(v61, int32(11))
	v70 = v66 ^ v52 - base.I32_rotl(v66, int32(25))
	v74 = v70 ^ v61 - base.I32_rotl(v70, int32(16))
	v78 = v74 ^ v66 - base.I32_rotl(v74, int32(4))
	v82 = v78 ^ v70 - base.I32_rotl(v78, v59)
	v92 = F_Int64GetDatum(m, base.I64_extend_i32_u(v82)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v82^v74-base.I32_rotl(v82, int32(24))))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		return int32(0)
	} else {
		return v92
	}
}
func F_hashfloat4extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 float32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 float64
	_ = v20
	var v26 float64
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.F32_eq(v12, float32(0)) != 0 {
		v15 = F_Int64GetDatum(m, v11)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v349 = v15
			m.G0 = v8 + int32(16)
			return v349
		}
	} else {
		v20 = base.F64_promote_f32(v12)
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807)) {
			v26 = math.Float64frombits(uint64(0x7ff8000000000000))
		} else {
			v26 = v20
		}
		*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v26
		v29 = v8 + int32(8)
		v36 = int32(-1636608424)
		if v11 == int64(0) {
			v73 = v36
			v75 = v36
			v77 = v36
		} else {
			v39 = base.I32_wrap_i64(v11)
			v41 = v39 + int32(1021750448)
			v45 = int32(4)
			v47 = base.I32_wrap_i64(int64(base.Ui64(v11)>>(uint(int64(32))%64))) ^ base.I32_rotl(v36, v45)
			v51 = v36 + v39 - v47 ^ base.I32_rotl(v47, int32(6))
			v55 = v41 - v51 ^ base.I32_rotl(v51, int32(8))
			v56 = v47 + v41
			v57 = v51 + v56
			v58 = v55 + v57
			v62 = v56 - v55 ^ base.I32_rotl(v55, int32(16))
			v66 = v57 - v62 ^ base.I32_rotl(v62, int32(19))
			v71 = v62 + v58
			v73 = v71
			v75 = v58 - v66 ^ base.I32_rotl(v66, v45)
			v77 = v66 + v71
		}
		if v29&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v301 = v73
				v302 = v77
				v303 = v75
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 1:
				v294 = v73
				v295 = v77
				v296 = v75
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 2:
				v287 = v73
				v288 = v77
				v289 = v75
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 3:
				v281 = v77
				v282 = v75
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 4:
				v277 = v77
				v278 = v75
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 5:
				v271 = v77
				v272 = v75
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 6:
				v265 = v77
				v266 = v75
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v271 = v267<<(uint(int32(16))%32) + v265
				v272 = v266
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 7:
				v260 = v75
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
				v265 = v261<<(uint(int32(24))%32) + v77
				v266 = v260
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v271 = v267<<(uint(int32(16))%32) + v265
				v272 = v266
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 8:
				v255 = v75
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v260 = v256<<(uint(int32(8))%32) + v255
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
				v265 = v261<<(uint(int32(24))%32) + v77
				v266 = v260
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v271 = v267<<(uint(int32(16))%32) + v265
				v272 = v266
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 9:
				v250 = v75
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+9)))
				v255 = v251<<(uint(int32(16))%32) + v250
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v260 = v256<<(uint(int32(8))%32) + v255
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
				v265 = v261<<(uint(int32(24))%32) + v77
				v266 = v260
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v271 = v267<<(uint(int32(16))%32) + v265
				v272 = v266
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			case 10:
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)))
				v250 = v246<<(uint(int32(24))%32) + v75
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+9)))
				v255 = v251<<(uint(int32(16))%32) + v250
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v260 = v256<<(uint(int32(8))%32) + v255
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+7)))
				v265 = v261<<(uint(int32(24))%32) + v77
				v266 = v260
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v271 = v267<<(uint(int32(16))%32) + v265
				v272 = v266
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v277 = v273<<(uint(int32(8))%32) + v271
				v278 = v272
				v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v281 = v277 + v279
				v282 = v278
				v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+3)))
				v287 = v283<<(uint(int32(24))%32) + v73
				v288 = v281
				v289 = v282
				v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v294 = v290<<(uint(int32(16))%32) + v287
				v295 = v288
				v296 = v289
				v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v301 = v297<<(uint(int32(8))%32) + v294
				v302 = v295
				v303 = v296
				v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v301 + v304
				v310 = v302
				v311 = v303
			default:
				v309 = v73
				v310 = v77
				v311 = v75
			}
		} else {
			switch int32(7) {
			case 0:
				v243 = v73
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v243 + v244
				v310 = v77
				v311 = v75
			case 1:
				v238 = v73
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v243 + v244
				v310 = v77
				v311 = v75
			case 2:
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+2)))
				v238 = v234<<(uint(int32(16))%32) + v73
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
				v309 = v243 + v244
				v310 = v77
				v311 = v75
			case 3:
				v231 = v77
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v309 = v232 + v73
				v310 = v231
				v311 = v75
			case 4:
				v228 = v77
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v231 = v228 + v229
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v309 = v232 + v73
				v310 = v231
				v311 = v75
			case 5:
				v223 = v77
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v228 = v224<<(uint(int32(8))%32) + v223
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v231 = v228 + v229
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v309 = v232 + v73
				v310 = v231
				v311 = v75
			case 6:
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+6)))
				v223 = v219<<(uint(int32(16))%32) + v77
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)))
				v228 = v224<<(uint(int32(8))%32) + v223
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)))
				v231 = v228 + v229
				v232 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v309 = v232 + v73
				v310 = v231
				v311 = v75
			case 7:
				v214 = v75
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v309 = v215 + v73
				v310 = v217 + v77
				v311 = v214
			case 8:
				v209 = v75
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v309 = v215 + v73
				v310 = v217 + v77
				v311 = v214
			case 9:
				v204 = v75
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+9)))
				v209 = v205<<(uint(int32(16))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v309 = v215 + v73
				v310 = v217 + v77
				v311 = v214
			case 10:
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+10)))
				v204 = v200<<(uint(int32(24))%32) + v75
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+9)))
				v209 = v205<<(uint(int32(16))%32) + v204
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+8)))
				v214 = v210<<(uint(int32(8))%32) + v209
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
				v309 = v215 + v73
				v310 = v217 + v77
				v311 = v214
			default:
				v309 = v73
				v310 = v77
				v311 = v75
			}
		}
		v314 = int32(14)
		v316 = v310 ^ v311 - base.I32_rotl(v310, v314)
		v320 = v316 ^ v309 - base.I32_rotl(v316, int32(11))
		v324 = v320 ^ v310 - base.I32_rotl(v320, int32(25))
		v328 = v324 ^ v316 - base.I32_rotl(v324, int32(16))
		v332 = v328 ^ v320 - base.I32_rotl(v328, int32(4))
		v336 = v332 ^ v324 - base.I32_rotl(v332, v314)
		v346 = F_Int64GetDatum(m, base.I64_extend_i32_u(v336)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v328^v336-base.I32_rotl(v336, int32(24))))
		mBase = m.M
		v347 = m.ExcPending
		if v347 != 0 {
			return int32(0)
		} else {
			v349 = v346
			m.G0 = v8 + int32(16)
			return v349
		}
	}
}
func F_hashint4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return v28 ^ v20 - base.I32_rotl(v28, int32(24))
}
func F_hashint8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v16 = int32(711645284)
	v19 = base.I32_wrap_i64(v4>>(uint(int64(63))%64)^int64(base.Ui64(v4)>>(uint(int64(32))%64))^v4) - int32(1636608428) ^ v16 - int32(1455628627)
	v24 = v19 ^ int32(-1636608428) - base.I32_rotl(v19, int32(25))
	v29 = v24 ^ v16 - base.I32_rotl(v24, int32(16))
	v33 = v29 ^ v19 - base.I32_rotl(v29, int32(4))
	v37 = v33 ^ v24 - base.I32_rotl(v33, int32(14))
	return v37 ^ v29 - base.I32_rotl(v37, int32(24))
}
func F_hashtranslatestrategy(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	if l0 == int32(1) {
		v7 = int32(3)
	} else {
		v7 = int32(0)
	}
	return v7
}
func F_hba_authname(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[572])))
	return v6
}
func F_heap2_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int64
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int64
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v337 int32
	_ = v337
	var v357 int32
	_ = v357
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+96))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+48)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v22, v23, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v27 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	switch int32(base.Ui32(v20)>>(uint(int32(4))%32))&int32(7) - int32(5) {
	case 0:
		goto L6
	default:
		goto L3
	case 2:
		goto L5
	}
L5:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v209 != 0 {
		goto L3
	} else {
		goto L44
	}
L6:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v37 = F_SnapBuildProcessChange(m, v21, v23, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v37 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v41 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v42 = m.G0
	v44 = v42 - int32(16)
	m.G0 = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+64))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v49&int32(8) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v44 + int32(16)
	return
L11:
	;
	v54 = int32(0)
	F_XLogRecGetBlockTag(m, v46, v54, v44, v54, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+88))
	if v59 != v61 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v63 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+56)))
	v66 = F_filter_by_origin_cb_wrapper(m, l0, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v68 = int32(0)
	v70 = v44 + int32(12)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+72))
	if v73 < v68 {
		v95 = v68
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v66 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
	if v99 == int32(0) {
		goto L10
	} else {
		goto L30
	}
L20:
	;
	v98 = v95
	goto L19
L21:
	;
	v79 = v72 + int32(76)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 != int32(1) {
		v95 = v68
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+43)))
	if v83 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v70 == int32(0) {
		v95 = v68
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v70 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v88
	v98 = v88
	goto L19
L27:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v70))) = v91
	goto L29
L28:
	;
	goto L29
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+44))
	v95 = v93
	goto L20
L30:
	;
	v105 = int32(0)
	v106 = v98
	goto L31
L31:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v120 = F_ReorderBufferAllocChange(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L10
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+16)) = uint16(v125)
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v120)+20)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+28)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v135 = (v106 + int32(1)) & int32(-2)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135))))
	v137 = F_ReorderBufferAllocTupleBuf(m, v131, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+40)) = v137
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+12)) = v140
	*(*uint16)(unsafe.Add(mBase, uint32(v137)+8)) = uint16(v140)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = int32(-1)
	v146 = int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v136 + v146
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v149))) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v149)+15)) = v150
	*(*int64)(unsafe.Add(mBase, uint32(v149)+8)) = v150
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v160 = v135 + int32(7)
	if v136 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+20)) = uint16(v163)
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v149)+18)) = uint16(v165)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+22)) = uint8(v167)
	v170 = v105 + int32(1)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if v171&int32(2) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v161 = F__emscripten_memcpy_bulkmem(m, v156+v146, v160, v136)
	mBase = m.M
	goto L38
L37:
	;
	goto L38
L38:
	;
	goto L35
L39:
	;
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
	v177 = base.B2i32(v170 == v174)
	goto L41
L40:
	;
	v177 = int32(0)
	goto L41
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+32)) = uint8(v177)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+36))
	v182 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferQueueChange(m, v179, v181, v182, v120, int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
	if base.Ui32(v170) < base.Ui32(v187) {
		v105 = v170
		v106 = v160 + v136
		goto L31
	} else {
		goto L43
	}
L43:
	;
	goto L32
L44:
	;
	v210 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+96))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+64))
	v214 = m.G0
	v216 = v214 - int32(32)
	m.G0 = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F_ReorderBufferXidSetCatalogChanges(m, v218, v23, v210)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v213)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+24)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v213)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+16)) = v228
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v216)+12)) = uint16(v230)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v213)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v232
	v235 = v216 + int32(16)
	v237 = v216 + int32(8)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v225)+124))
	v240 = F_MemoryContextAlloc(m, v238, int32(64))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v242 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v240))) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v240)+56)) = v242
	v247 = v240 + int32(48)
	*(*int64)(unsafe.Add(mBase, uint32(v247))) = v242
	v251 = v240 + int32(40)
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v242
	v255 = v240 + int32(32)
	*(*int64)(unsafe.Add(mBase, uint32(v255))) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v240)+24)) = v242
	*(*int64)(unsafe.Add(mBase, uint32(v240)+16)) = v242
	v263 = v240 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v263))) = v242
	v267 = F_ReorderBufferTXNByXid(m, v225, v224, int32(0), v210)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v235)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+28)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v235)))
	*(*int64)(unsafe.Add(mBase, uint32(v240)+20)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v273
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v240)+36)) = uint16(v275)
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v240)+44)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v223
	*(*int32)(unsafe.Add(mBase, uint32(v240)+12)) = v267
	*(*int64)(unsafe.Add(mBase, uint32(v240))) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v263))) = int32(7)
	v285 = v267 + int32(136)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v267)+140))
	if v286 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267)+140)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v267)+136)) = v285
	goto L50
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+56)) = v285
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v267)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+52)) = v292
	v295 = v240 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v292)+4)) = v295
	*(*int32)(unsafe.Add(mBase, uint32(v267)+136)) = v295
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v267)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v267)+144)) = v298 + int64(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	if v303 != int32(-1) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+124))
	v332 = F_MemoryContextAlloc(m, v330, int32(64))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L65
	}
L52:
	;
	if base.Ui32(v302) < base.Ui32(v303) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v302 != int32(-1) {
		v326 = v302
		goto L51
	} else {
		goto L61
	}
L55:
	;
	v307 = v303
	goto L57
L56:
	;
	v307 = v302
	goto L57
L57:
	;
	if v302 == int32(-1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v310 = v303
	goto L60
L59:
	;
	v310 = v307
	goto L60
L60:
	;
	v326 = v310
	goto L51
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(475567), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(511831), int32(716), int32(447346))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v334 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v332)+16)) = v334
	v337 = v332 + int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v337))) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332))) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332)+56)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332)+48)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332)+40)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332)+32)) = v334
	*(*int64)(unsafe.Add(mBase, uint32(v332)+24)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v332)+20)) = v326 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v337))) = int32(6)
	F_ReorderBufferQueueChange(m, v329, v23, v210, v332, int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	m.G0 = v216 + int32(32)
	goto L3
}
func F_heap2_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	if base.Ui32(l0) <= base.Ui32(int32(223)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[77])))
		v12 = v11
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_heap2_redo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v579 int32
	_ = v579
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v732 int32
	_ = v732
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v773 int32
	_ = v773
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v860 int32
	_ = v860
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v908 int32
	_ = v908
	var v916 int32
	_ = v916
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v967 int32
	_ = v967
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int64
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int64
	_ = v1034
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1100 int32
	_ = v1100
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1134 int32
	_ = v1134
	var v1136 int64
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1157 int32
	_ = v1157
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1201 int32
	_ = v1201
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1222 int32
	_ = v1222
	var v1224 int64
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1251 int32
	_ = v1251
	var v1252 int64
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1267 int32
	_ = v1267
	var v1269 int64
	_ = v1269
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1343 int32
	_ = v1343
	var v1349 int32
	_ = v1349
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1400 int32
	_ = v1400
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1466 int32
	_ = v1466
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int64
	_ = v1475
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1496 int32
	_ = v1496
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1563 int32
	_ = v1563
	var v1568 int32
	_ = v1568
	var v1577 int32
	_ = v1577
	var v1588 int32
	_ = v1588
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1609 int32
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1636 int32
	_ = v1636
	var v1652 int32
	_ = v1652
	var v1654 int32
	_ = v1654
	var v1662 int32
	_ = v1662
	var v1664 int64
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1671 int64
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int64
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1698 int32
	_ = v1698
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1717 int32
	_ = v1717
	var v1721 int32
	_ = v1721
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1744 int32
	_ = v1744
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1773 int32
	_ = v1773
	var v1792 int32
	_ = v1792
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1803 int64
	_ = v1803
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1814 int32
	_ = v1814
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1825 int64
	_ = v1825
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1828 int64
	_ = v1828
	var v1836 int64
	_ = v1836
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1854 int32
	_ = v1854
	var v1857 int64
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1877 int64
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1900 int32
	_ = v1900
	var v1905 int32
	_ = v1905
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1930 int32
	_ = v1930
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1952 int32
	_ = v1952
	var v1957 int32
	_ = v1957
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1978 int32
	_ = v1978
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1989 int32
	_ = v1989
	var v1990 int64
	_ = v1990
	var v1999 int32
	_ = v1999
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2018 int32
	_ = v2018
	var v2023 int32
	_ = v2023
	var v2050 int32
	_ = v2050
	var v2054 int32
	_ = v2054
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2067 int32
	_ = v2067
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2080 int32
	_ = v2080
	var v2085 int32
	_ = v2085
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2099 int32
	_ = v2099
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(8352)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v26 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25)+48)))
	switch int32(base.Ui32(v26)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0, 1, 2:
		goto L10
	case 3:
		goto L9
	case 4:
		goto L8
	case 5:
		goto L7
	case 6:
		goto L5
	default:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L11
	} else {
		goto L416
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L11
	} else {
		goto L413
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L11
	} else {
		goto L410
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2050 = m.ExcPending
	if v2050 != 0 {
		goto L11
	} else {
		goto L407
	}
L5:
	;
	m.G0 = v23 + int32(8352)
	return
L6:
	;
	v1819 = m.G0
	v1821 = v1819 - int32(1136)
	m.G0 = v1821
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1824 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+64))
	v1825 = *(*int64)(unsafe.Add(mBase, uint32(v1824)+4))
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1824)))
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(v1823)+36))
	v1828 = *(*int64)(unsafe.Add(mBase, uint32(v1824)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v1821)+96)) = uint32(v1828)
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+104)) = v1827
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+100)) = v1826
	*(*int64)(unsafe.Add(mBase, uint32(v1821)+84)) = v1825
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+80)) = int32(159799)
	v1836 = int64(base.Ui64(v1828) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1821)+92)) = uint32(v1836)
	v1844 = F_pg_snprintf(m, v1821+int32(112), int32(1024), int32(30689), v1821+int32(80))
	mBase = m.M
	v1845 = m.ExcPending
	if v1845 != 0 {
		goto L11
	} else {
		goto L356
	}
L7:
	;
	v1671 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v1673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672)+7)))
	if v1673&int32(1) != 0 {
		goto L325
	} else {
		goto L326
	}
L8:
	;
	v1252 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v1254 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v1254, v23+int32(8336), v1254, v23+int32(8348))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L11
	} else {
		goto L227
	}
L9:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v1015 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1016 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32]))) = v1016
	F_XLogRecGetBlockTag(m, l0, int32(1), v23+int32(168), v1016, v23+int32(8332))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L11
	} else {
		goto L165
	}
L10:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v35 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v35, v23+int32(168), v35, v23+int32(8348))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v43&int32(8) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = int32(0)
	v69 = v43 & int32(4)
	v74 = F_XLogReadBufferForRedoExtended(m, l0, v66, v66, int32(base.Ui32(v69)>>(uint(int32(2))%32)), v23+int32(8336))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if base.Ui32(v49) < base.Ui32(int32(2)) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v34)+2))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v53
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v55
	F_ResolveRecoveryConflictWithSnapshot(m, v52, int32(base.Ui32(v43&int32(2))>>(uint(int32(1))%32)), v23+int32(24))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v74 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	if v78 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	if v908 == int32(0) {
		goto L5
	} else {
		goto L135
	}
L21:
	;
	v97 = int32(0)
	v99 = v23 + int32(140)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+72))
	if v102 < v97 {
		v124 = v97
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v82+(v78^int32(-1))<<(uint(int32(2))%32))))
	v96 = v88
	goto L21
L23:
	;
	goto L24
L24:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v96 = v90 + v78<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	v129 = v23 + int32(144)
	v131 = v23 + int32(136)
	v133 = v23 + int32(132)
	v135 = v23 + int32(156)
	v139 = v23 + int32(152)
	v141 = v23 + int32(164)
	v143 = v23 + int32(148)
	v145 = v23 + int32(160)
	if v43&int32(16) != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v127 = v124
	goto L25
L27:
	;
	v108 = v101 + int32(76)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
	if v109 != int32(1) {
		v124 = v97
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+43)))
	if v112 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v99 == int32(0) {
		v124 = v97
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v117
	v127 = v117
	goto L25
L33:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v120
	goto L35
L34:
	;
	goto L35
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v108)+44))
	v124 = v122
	goto L26
L36:
	;
	if v43&int32(32) != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v148
	v151 = v127 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v161 = v151 + v153*int32(12)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v157 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v157
	v161 = v127
	goto L36
L40:
	;
	if v43&int32(64) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161))))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v164
	v166 = int32(2)
	v167 = v161 + v166
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34]))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v177 = v167 + v169<<(uint(v166)%32)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34]))) = v173
	v177 = v161
	goto L40
L44:
	;
	if base.I32_extend8_s(v43) < int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177))))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v180
	v183 = v177 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v193 = v183 + v185<<(uint(int32(1))%32)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v189
	v193 = v177
	goto L44
L48:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v23)+148))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v23)+152))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	if int32(0) < v215 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v197
	v200 = v193 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v200 + v202<<(uint(int32(1))%32)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v207 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v193
	goto L48
L52:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v23)+144))
	if int32(0) < v753 {
		goto L116
	} else {
		goto L117
	}
L53:
	;
	v222 = int32(0)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v23)+160))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	if v227 < v222 {
		goto L58
	} else {
		goto L59
	}
L54:
	;
	if int32(0) < v214 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	if v213 <= int32(0) {
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	if v215 <= int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231+(v227^int32(-1))<<(uint(int32(2))%32))))
	v245 = v237
	goto L57
L59:
	;
	goto L60
L60:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v245 = v239 + v227<<(uint(int32(13))%32) + int32(-8192)
	goto L57
L61:
	;
	if v214 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L62:
	;
	v248 = int32(1)
	v251 = v245 + int32(24)
	if v215 != v248 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v258 = v224
	v261 = int32(0)
	goto L66
L64:
	;
	v307 = v224
	goto L65
L65:
	;
	if v215&v248 == int32(0) {
		goto L61
	} else {
		goto L69
	}
L66:
	;
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258))))
	v278 = int32(2)
	v281 = int32(4)
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+2)))
	v284 = int32(32767)
	v286 = int32(65536)
	*(*int32)(unsafe.Add(mBase, uint32(v277<<(uint(v278)%32)+v251-v281))) = v283&v284 | v286
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+4)))
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v289<<(uint(v278)%32)+v251-v281))) = v295&v284 | v286
	v302 = v258 + int32(8)
	v304 = v261 + v278
	if v304 != v215&int32(2147483646) {
		v258 = v302
		v261 = v304
		goto L66
	} else {
		goto L68
	}
L67:
	;
	v307 = v302
	goto L65
L68:
	;
	goto L67
L69:
	;
	v328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307))))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v328<<(uint(int32(2))%32)+v251-int32(4)))) = v334&int32(32767) | int32(65536)
	goto L61
L70:
	;
	if v213 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v363 = v214 & int32(3)
	v365 = v245 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v214) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v376 = int32(0)
	v377 = v225
	goto L75
L73:
	;
	v434 = v225
	goto L74
L74:
	;
	if v363 == int32(0) {
		goto L70
	} else {
		goto L78
	}
L75:
	;
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377))))
	v392 = int32(2)
	v395 = int32(4)
	v397 = int32(98304)
	*(*int32)(unsafe.Add(mBase, uint32(v391<<(uint(v392)%32)+v365-v395))) = v397
	v399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v399<<(uint(v392)%32)+v365-v395))) = v397
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v407<<(uint(v392)%32)+v365-v395))) = v397
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v377)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v415<<(uint(v392)%32)+v365-v395))) = v397
	v424 = v377 + int32(8)
	v426 = v376 + v395
	if v426 != v214&int32(2147483644) {
		v376 = v426
		v377 = v424
		goto L75
	} else {
		goto L77
	}
L76:
	;
	v434 = v424
	goto L74
L77:
	;
	goto L76
L78:
	;
	v456 = int32(0)
	v457 = v434
	goto L79
L79:
	;
	v471 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457))))
	v472 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v471<<(uint(v472)%32)+v365-int32(4)))) = int32(98304)
	v482 = v456 + int32(1)
	if v482 != v363 {
		v456 = v482
		v457 = v457 + v472
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L70
L81:
	;
	goto L80
L82:
	;
	if v69 == v222 {
		goto L95
	} else {
		goto L96
	}
L83:
	;
	v507 = v213 & int32(3)
	v509 = v245 + int32(24)
	if base.Ui32(int32(4)) <= base.Ui32(v213) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v516 = int32(0)
	v522 = v226
	goto L87
L85:
	;
	v579 = v226
	goto L86
L86:
	;
	if v507 == int32(0) {
		goto L82
	} else {
		goto L90
	}
L87:
	;
	v535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522))))
	v536 = int32(2)
	v539 = int32(4)
	v541 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v535<<(uint(v536)%32)+v509-v539))) = v541
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v543<<(uint(v536)%32)+v509-v539))) = v541
	v551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v551<<(uint(v536)%32)+v509-v539))) = v541
	v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v522)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v559<<(uint(v536)%32)+v509-v539))) = v541
	v568 = v522 + int32(8)
	v570 = v516 + v539
	if v570 != v213&int32(2147483644) {
		v516 = v570
		v522 = v568
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v579 = v568
	goto L86
L89:
	;
	goto L88
L90:
	;
	v596 = int32(0)
	v602 = v579
	goto L91
L91:
	;
	v615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v602))))
	v616 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v615<<(uint(v616)%32)+v509-int32(4)))) = int32(0)
	v626 = v596 + int32(1)
	if v626 != v507 {
		v596 = v626
		v602 = v602 + v616
		goto L91
	} else {
		goto L93
	}
L92:
	;
	goto L82
L93:
	;
	goto L92
L94:
	;
	goto L52
L95:
	;
	v648 = int32(0)
	v654 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+12)))
	if base.Ui32(v654) < base.Ui32(int32(25)) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	goto L97
L97:
	;
	F_PageRepairFragmentation(m, v245)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L11
	} else {
		goto L115
	}
L98:
	;
	goto L94
L99:
	;
	v720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+10)))
	v722 = v720 & int32(65534)
	*(*uint16)(unsafe.Add(mBase, uint32(v245)+10)) = uint16(v722)
	goto L98
L100:
	;
	v662 = int32(base.Ui32(v654+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v662 == int32(0) {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v668 = v662
	v669 = v648
	v672 = v648
	goto L103
L102:
	;
	if int32(0) < v698 {
		goto L111
	} else {
		goto L112
	}
L103:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v668&int32(65535)<<(uint(int32(2))%32)+(v245+int32(24))-int32(4))))
	v683 = v681 & int32(98304)
	if v668 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v698 = v692
	v700 = int32(0)
	goto L102
L105:
	;
	v695 = v668 - int32(1)
	if v695 != 0 {
		v668 = v695
		v669 = v692
		v672 = v693
		goto L103
	} else {
		goto L110
	}
L106:
	;
	if v683 != 0 {
		v692 = v669
		v693 = v672
		goto L105
	} else {
		goto L109
	}
L107:
	;
	if v672 != 0 {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v686 = int32(0)
	v692 = v669 + base.B2i32(v683 == v686)
	v693 = base.B2i32(v683 != v686)
	goto L105
L109:
	;
	v698 = v669
	v700 = int32(1)
	goto L102
L110:
	;
	goto L104
L111:
	;
	v705 = v654 - v698<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v245)+12)) = uint16(v705)
	goto L113
L112:
	;
	goto L113
L113:
	;
	if v700 == int32(0) {
		goto L99
	} else {
		goto L114
	}
L114:
	;
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v245)+10)))
	v711 = v709 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v245)+10)) = uint16(v711)
	goto L98
L115:
	;
	goto L94
L116:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	v773 = v2
	goto L119
L117:
	;
	goto L118
L118:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_rotr(v33, int64(32))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	F_MarkBufferDirty(m, v885)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L11
	} else {
		goto L134
	}
L119:
	;
	v781 = v758 + v773*int32(12)
	v783 = v781 + int32(10)
	v784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v783))))
	if v784 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	goto L118
L121:
	;
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v781)))
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v781)+6)))
	v787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v781)+4)))
	v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v781)+8)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v795 = int32(0)
	v796 = v794
	goto L124
L122:
	;
	goto L123
L123:
	;
	v860 = v773 + int32(1)
	if v860 != v753 {
		v773 = v860
		goto L119
	} else {
		goto L133
	}
L124:
	;
	v815 = int32(2)
	v816 = v796 + v815
	*(*int32)(unsafe.Add(mBase, uint32(v23)+132)) = v816
	v818 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796))))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818<<(uint(v815)%32)+(v96+int32(24))-int32(4))))
	v827 = v96 + v824&int32(32767)
	*(*int32)(unsafe.Add(mBase, uint32(v827)+4)) = v785
	if v788&int32(2) != 0 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	goto L123
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v827)+8)) = int32(2)
	goto L128
L127:
	;
	goto L128
L128:
	;
	if v788&int32(4) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v827)+8)) = int32(0)
	goto L131
L130:
	;
	goto L131
L131:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v827)+18)) = uint16(v787)
	*(*uint16)(unsafe.Add(mBase, uint32(v827)+20)) = uint16(v786)
	v836 = v795 + int32(1)
	v837 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v783))))
	if base.Ui32(v836) < base.Ui32(v837) {
		v795 = v836
		v796 = v816
		goto L124
	} else {
		goto L132
	}
L132:
	;
	goto L125
L133:
	;
	goto L120
L134:
	;
	goto L20
L135:
	;
	if base.Ui32(int32(32)) <= base.Ui32(v43) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	if v908 < int32(0) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	F_UnlockReleaseBuffer(m, v908)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L11
	} else {
		goto L164
	}
L139:
	;
	v934 = int32(4)
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v930)+14)))
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v930)+12)))
	v937 = v935 - v936
	if v937 <= v934 {
		goto L144
	} else {
		goto L145
	}
L140:
	;
	v916 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v916+(v908^int32(-1))<<(uint(int32(2))%32))))
	v930 = v922
	goto L139
L141:
	;
	goto L142
L142:
	;
	v924 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v930 = v924 + v908<<(uint(int32(13))%32) + int32(-8192)
	goto L139
L143:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	F_UnlockReleaseBuffer(m, v1000)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L11
	} else {
		goto L162
	}
L144:
	;
	v940 = v934
	goto L146
L145:
	;
	v940 = v937
	goto L146
L146:
	;
	v942 = v940 - int32(4)
	if v942 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v999 = int32(0)
	goto L143
L148:
	;
	goto L149
L149:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v936) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v999 = v942
	goto L143
L151:
	;
	v953 = int32(base.Ui32(v936+int32(262120)) >> (uint(int32(2)) % 32))
	goto L153
L152:
	;
	v953 = int32(0)
	goto L153
L153:
	;
	if base.Ui32(v953&int32(65535)) < base.Ui32(int32(291)) {
		goto L150
	} else {
		goto L154
	}
L154:
	;
	v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+10)))
	if v958&int32(1) == int32(0) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v999 = int32(0)
	goto L143
L156:
	;
	goto L157
L157:
	;
	v967 = int32(1)
	goto L158
L158:
	;
	v978 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v967&int32(65535)<<(uint(int32(2))%32)+(v930+int32(24))-int32(3)))))
	if v978&int32(384) == int32(0) {
		goto L150
	} else {
		goto L160
	}
L159:
	;
	v999 = int32(0)
	goto L143
L160:
	;
	v984 = v967 + int32(1)
	v985 = int32(65535)
	if base.Ui32(v984&v985) <= base.Ui32(v953&v985) {
		v967 = v984
		goto L158
	} else {
		goto L161
	}
L161:
	;
	goto L159
L162:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v1003
	v1005 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v1005
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_XLogRecordPageWithFreeSpace(m, v23+int32(8), v1009, v999)
	mBase = m.M
	v1011 = m.ExcPending
	if v1011 != 0 {
		goto L11
	} else {
		goto L163
	}
L163:
	;
	goto L5
L164:
	;
	goto L5
L165:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, _consts[33]))
	if base.Ui32(int32(2)) <= base.Ui32(v1027) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1031 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+4)))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1032
	v1034 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v1034
	F_ResolveRecoveryConflictWithSnapshot(m, v1030, int32(base.Ui32(v1031&int32(4))>>(uint(int32(2))%32)), v23+int32(72))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L11
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v1049 = F_XLogReadBufferForRedo(m, l0, int32(1), v23+int32(8348))
	mBase = m.M
	v1050 = m.ExcPending
	if v1050 != 0 {
		goto L11
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	if v1049 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	if v1053 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	goto L173
L173:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	if v1094 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L174:
	;
	v1072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1071)+10)))
	v1074 = v1072 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1071)+10)) = uint16(v1074)
	v1077 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+252))
	goto L179
L175:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1057+(v1053^int32(-1))<<(uint(int32(2))%32))))
	v1071 = v1063
	goto L174
L176:
	;
	goto L177
L177:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1071 = v1065 + v1053<<(uint(int32(13))%32) + int32(-8192)
	goto L174
L178:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_MarkBufferDirty(m, v1090)
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L11
	} else {
		goto L184
	}
L179:
	;
	if base.B2i32(v1078 != int32(0)) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1084 = int32(*(*uint8)(unsafe.Add(mBase, _consts[37])))
	if v1084 != int32(1) {
		goto L178
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1071))) = base.I64_rotr(v1015, int64(32))
	goto L178
L183:
	;
	goto L182
L184:
	;
	goto L173
L185:
	;
	v1144 = int32(0)
	v1149 = F_XLogReadBufferForRedoExtended(m, l0, v1144, int32(3), v1144, v23+int32(8336))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L11
	} else {
		goto L198
	}
L186:
	;
	if v1094 < int32(0) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v1115 = int32(4)
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114)+14)))
	v1117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1114)+12)))
	v1118 = v1116 - v1117
	if v1118 <= v1115 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v1100 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1100+(v1094^int32(-1))<<(uint(int32(2))%32))))
	v1114 = v1106
	goto L187
L189:
	;
	goto L190
L190:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1114 = v1108 + v1094<<(uint(int32(13))%32) + int32(-8192)
	goto L187
L191:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_UnlockReleaseBuffer(m, v1124)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L11
	} else {
		goto L195
	}
L192:
	;
	v1121 = v1115
	goto L194
L193:
	;
	v1121 = v1118
	goto L194
L194:
	;
	goto L191
L195:
	;
	v1127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+4)))
	if v1127&int32(3) == int32(0) {
		goto L185
	} else {
		goto L196
	}
L196:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23-int32(-64)))) = v1134
	v1136 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v1136
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	F_XLogRecordPageWithFreeSpace(m, v23+int32(56), v1140, v1121-int32(4))
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		goto L11
	} else {
		goto L197
	}
L197:
	;
	goto L185
L198:
	;
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	if v1149 == int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if v1151 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	goto L201
L201:
	;
	if v1151 == int32(0) {
		goto L5
	} else {
		goto L225
	}
L202:
	;
	v1172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1171)+14)))
	if v1172 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v1157 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1157+(v1151^int32(-1))<<(uint(int32(2))%32))))
	v1171 = v1163
	goto L202
L204:
	;
	goto L205
L205:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1171 = v1165 + v1151<<(uint(int32(13))%32) + int32(-8192)
	goto L202
L206:
	;
	if v1171&int32(3) != 0 {
		goto L211
	} else {
		goto L212
	}
L207:
	;
	v1217 = v1151
	goto L208
L208:
	;
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1014)+4)))
	F_LockBuffer(m, v1217, int32(0))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L11
	} else {
		goto L219
	}
L209:
	;
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	v1217 = v1216
	goto L208
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1171)+10)) = int32(1572864)
	v1207 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v1171)+18)) = uint16(v1207)
	v1213 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v1171)+16)) = uint16(v1213)
	*(*uint16)(unsafe.Add(mBase, uint32(v1171)+14)) = uint16(v1213)
	goto L209
L211:
	;
	v1201 = F___memset(m, v1171, int32(0), int32(8192))
	mBase = m.M
	goto L210
L212:
	;
	goto L211
L219:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v1222
	v1224 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v1224
	v1228 = F_CreateFakeRelcacheEntry(m, v23+int32(40))
	mBase = m.M
	v1229 = m.ExcPending
	if v1229 != 0 {
		goto L11
	} else {
		goto L220
	}
L220:
	;
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	F_visibilitymap_pin(m, v1228, v1230, v23+int32(8336))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L11
	} else {
		goto L221
	}
L221:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1241 = F_visibilitymap_set(m, v1228, v1235, int32(0), v1015, v1237, v1238, v1218&int32(3))
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L11
	} else {
		goto L222
	}
L222:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	F_ReleaseBuffer(m, v1243)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L11
	} else {
		goto L223
	}
L223:
	;
	F_pfree(m, v1228)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L11
	} else {
		goto L224
	}
L224:
	;
	goto L5
L225:
	;
	F_UnlockReleaseBuffer(m, v1151)
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L11
	} else {
		goto L226
	}
L226:
	;
	goto L5
L227:
	;
	v1262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	if v1262&int32(1) != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1267 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[38])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v1267
	v1269 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+104)) = v1269
	v1273 = F_CreateFakeRelcacheEntry(m, v23+int32(104))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L11
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	if v26 < int32(0) {
		goto L238
	} else {
		goto L239
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = int32(0)
	v1277 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_visibilitymap_pin(m, v1273, v1277, v23+int32(168))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	v1285 = F_visibilitymap_clear(m, v1282, v1283, int32(3))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L11
	} else {
		goto L233
	}
L233:
	;
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	F_ReleaseBuffer(m, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L11
	} else {
		goto L234
	}
L234:
	;
	F_pfree(m, v1273)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L11
	} else {
		goto L235
	}
L235:
	;
	goto L230
L236:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	if v1652 != 0 {
		goto L319
	} else {
		goto L320
	}
L237:
	;
	v1365 = int32(0)
	v1367 = v23 + int32(164)
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1369)+72))
	if v1370 < v1365 {
		v1392 = v1365
		goto L259
	} else {
		goto L260
	}
L238:
	;
	v1296 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L11
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	v1358 = int32(0)
	v1362 = F_XLogReadBufferForRedo(m, l0, v1358, v23+int32(8332))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L11
	} else {
		goto L256
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34]))) = v1296
	if v1296 < int32(0) {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	if v1316&int32(3) != 0 {
		goto L248
	} else {
		goto L249
	}
L243:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1302+(v1296^int32(-1))<<(uint(int32(2))%32))))
	v1316 = v1308
	goto L242
L244:
	;
	goto L245
L245:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1316 = v1310 + v1296<<(uint(int32(13))%32) + int32(-8192)
	goto L242
L246:
	;
	goto L237
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1316)+10)) = int32(1572864)
	v1349 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v1316)+18)) = uint16(v1349)
	v1355 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v1316)+16)) = uint16(v1355)
	*(*uint16)(unsafe.Add(mBase, uint32(v1316)+14)) = uint16(v1355)
	goto L246
L248:
	;
	v1343 = F___memset(m, v1316, int32(0), int32(8192))
	mBase = m.M
	goto L247
L249:
	;
	goto L248
L256:
	;
	if v1362 != 0 {
		v1633 = v1358
		v1636 = v2
		goto L236
	} else {
		goto L257
	}
L257:
	;
	goto L237
L258:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	if v1396 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L259:
	;
	v1395 = v1392
	goto L258
L260:
	;
	v1376 = v1369 + int32(76)
	v1377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376))))
	if v1377 != int32(1) {
		v1392 = v1365
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1376)+43)))
	if v1380 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	if v1367 == int32(0) {
		v1392 = v1365
		goto L259
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	if v1367 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1385 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1367))) = v1385
	v1395 = v1385
	goto L258
L266:
	;
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1376)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1367))) = v1388
	goto L268
L267:
	;
	goto L268
L268:
	;
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+44))
	v1392 = v1390
	goto L259
L269:
	;
	v1415 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	v1416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1253)+2)))
	if v1416 != 0 {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1400+(v1396^int32(-1))<<(uint(int32(2))%32))))
	v1414 = v1406
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1414 = v1408 + v1396<<(uint(int32(13))%32) + int32(-8192)
	goto L269
L273:
	;
	v1426 = v23 + int32(176)
	v1428 = v1395
	v1431 = v2
	goto L276
L274:
	;
	v1520 = v1395
	goto L275
L275:
	;
	if v1520 != v1395+v1415 {
		goto L2
	} else {
		goto L292
	}
L276:
	;
	if int32(0) <= v26 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1520 = v1514
	goto L275
L278:
	;
	v1450 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1253+int32(4)+v1431<<(uint(int32(1))%32)))))
	v1453 = v1450
	goto L280
L279:
	;
	v1453 = v1431 + int32(1)
	goto L280
L280:
	;
	v1455 = v1453 & int32(65535)
	v1456 = int32(1)
	v1457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+12)))
	if base.Ui32(v1457) < base.Ui32(int32(25)) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1466 = v1456
	goto L283
L282:
	;
	v1466 = int32(base.Ui32(v1457+int32(262120))>>(uint(int32(2))%32)) + v1456
	goto L283
L283:
	;
	if base.Ui32(v1466&int32(65535)) < base.Ui32(v1455) {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	v1473 = (v1428 + int32(1)) & int32(-2)
	v1474 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473))))
	v1475 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23+int32(183)))) = v1475
	*(*int64)(unsafe.Add(mBase, uint32(v1426))) = v1475
	*(*int64)(unsafe.Add(mBase, uint32(v23)+168)) = v1475
	v1482 = v1473 + int32(7)
	if v1474 != 0 {
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+186)) = uint16(v1485)
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1473)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1487)
	v1489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1473)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+190)) = uint8(v1489)
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1491)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v1426))) = int32(0)
	v1496 = v1487 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1496)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v1492
	v1499 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+182)) = uint16(v1499)
	v1502 = int32(base.Ui32(v1499) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+180)) = uint16(v1502)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+184)) = uint16(v1453)
	v1510 = F_PageAddItemExtended(m, v1414, v23+int32(168), v1474+int32(23), v1455, int32(3))
	mBase = m.M
	v1511 = m.ExcPending
	if v1511 != 0 {
		goto L11
	} else {
		goto L289
	}
L286:
	;
	v1483 = F__emscripten_memcpy_bulkmem(m, v23+int32(191), v1482, v1474)
	mBase = m.M
	goto L288
L287:
	;
	goto L288
L288:
	;
	goto L285
L289:
	;
	if v1510 == int32(0) {
		goto L3
	} else {
		goto L290
	}
L290:
	;
	v1514 = v1474 + v1482
	v1516 = v1431 + int32(1)
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1253)+2)))
	if base.Ui32(v1516) < base.Ui32(v1517) {
		v1428 = v1514
		v1431 = v1516
		goto L276
	} else {
		goto L291
	}
L291:
	;
	goto L277
L292:
	;
	v1544 = int32(4)
	v1545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+14)))
	v1546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+12)))
	v1547 = v1545 - v1546
	if v1547 <= v1544 {
		goto L294
	} else {
		goto L295
	}
L293:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1414))) = base.I64_rotr(v1252, int64(32))
	v1613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	if v1613&int32(1) != 0 {
		goto L312
	} else {
		goto L313
	}
L294:
	;
	v1550 = v1544
	goto L296
L295:
	;
	v1550 = v1547
	goto L296
L296:
	;
	v1552 = v1550 - int32(4)
	if v1552 == int32(0) {
		goto L297
	} else {
		goto L298
	}
L297:
	;
	v1609 = int32(0)
	goto L293
L298:
	;
	goto L299
L299:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1546) {
		goto L301
	} else {
		goto L302
	}
L300:
	;
	v1609 = v1552
	goto L293
L301:
	;
	v1563 = int32(base.Ui32(v1546+int32(262120)) >> (uint(int32(2)) % 32))
	goto L303
L302:
	;
	v1563 = int32(0)
	goto L303
L303:
	;
	if base.Ui32(v1563&int32(65535)) < base.Ui32(int32(291)) {
		goto L300
	} else {
		goto L304
	}
L304:
	;
	v1568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414)+10)))
	if v1568&int32(1) == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1609 = int32(0)
	goto L293
L306:
	;
	goto L307
L307:
	;
	v1577 = int32(1)
	goto L308
L308:
	;
	v1588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1577&int32(65535)<<(uint(int32(2))%32)+(v1414+int32(24))-int32(3)))))
	if v1588&int32(384) == int32(0) {
		goto L300
	} else {
		goto L310
	}
L309:
	;
	v1609 = int32(0)
	goto L293
L310:
	;
	v1594 = v1577 + int32(1)
	v1595 = int32(65535)
	if base.Ui32(v1594&v1595) <= base.Ui32(v1563&v1595) {
		v1577 = v1594
		goto L308
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	v1616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+10)))
	v1618 = v1616 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v1414)+10)) = uint16(v1618)
	v1620 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1253))))
	v1621 = v1620
	goto L314
L313:
	;
	v1621 = v1613
	goto L314
L314:
	;
	if v1621&int32(32) != 0 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+10)))
	v1626 = v1624 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1414)+10)) = uint16(v1626)
	goto L317
L316:
	;
	goto L317
L317:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[34])))
	F_MarkBufferDirty(m, v1628)
	mBase = m.M
	v1630 = m.ExcPending
	if v1630 != 0 {
		goto L11
	} else {
		goto L318
	}
L318:
	;
	v1633 = v1609
	v1636 = int32(1)
	goto L236
L319:
	;
	F_UnlockReleaseBuffer(m, v1652)
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L11
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	if v1636&base.B2i32(base.Ui32(v1633) < base.Ui32(int32(1638))) == int32(0) {
		goto L5
	} else {
		goto L323
	}
L322:
	;
	goto L321
L323:
	;
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[38])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = v1662
	v1664 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = v1664
	v1668 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_XLogRecordPageWithFreeSpace(m, v23+int32(88), v1668, v1633)
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L11
	} else {
		goto L324
	}
L324:
	;
	goto L5
L325:
	;
	v1676 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32]))) = v1676
	F_XLogRecGetBlockTag(m, l0, v1676, v23+int32(168), v1676, v23+int32(8348))
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L11
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1713 = F_XLogReadBufferForRedo(m, l0, int32(0), v23+int32(168))
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L11
	} else {
		goto L334
	}
L328:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v1686
	v1688 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v1688
	v1692 = F_CreateFakeRelcacheEntry(m, v23+int32(120))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L11
	} else {
		goto L329
	}
L329:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	F_visibilitymap_pin(m, v1692, v1694, v23+int32(8336))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L11
	} else {
		goto L330
	}
L330:
	;
	v1699 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[35])))
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	v1702 = F_visibilitymap_clear(m, v1699, v1700, int32(2))
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L11
	} else {
		goto L331
	}
L331:
	;
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[32])))
	F_ReleaseBuffer(m, v1704)
	mBase = m.M
	v1706 = m.ExcPending
	if v1706 != 0 {
		goto L11
	} else {
		goto L332
	}
L332:
	;
	F_pfree(m, v1692)
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L11
	} else {
		goto L333
	}
L333:
	;
	goto L327
L334:
	;
	if v1713 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v1717 < int32(0) {
		goto L339
	} else {
		goto L340
	}
L336:
	;
	goto L337
L337:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v1814 == int32(0) {
		goto L5
	} else {
		goto L354
	}
L338:
	;
	v1736 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1735)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1736) {
		goto L342
	} else {
		goto L343
	}
L339:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1721+(v1717^int32(-1))<<(uint(int32(2))%32))))
	v1735 = v1727
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1735 = v1729 + v1717<<(uint(int32(13))%32) + int32(-8192)
	goto L338
L342:
	;
	v1744 = int32(base.Ui32(v1736+int32(262120)) >> (uint(int32(2)) % 32))
	goto L344
L343:
	;
	v1744 = int32(0)
	goto L344
L344:
	;
	v1747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1672)+4)))
	if base.Ui32(v1744&int32(65535)) < base.Ui32(v1747) {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v1747<<(uint(int32(2))%32)+v1735)+20))
	if v1752&int32(98304) != int32(32768) {
		goto L1
	} else {
		goto L346
	}
L346:
	;
	v1759 = v1735 + v1752&int32(32767)
	v1760 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759)+20)))
	v1762 = v1760 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+20)) = uint16(v1762)
	v1764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1759)+18)))
	v1766 = v1764 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+18)) = uint16(v1766)
	v1768 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1672)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+18)) = uint16(v1766)
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+20)) = uint16(v1762)
	if v1768&int32(15) != 0 {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1773 = int32(1)
	v1792 = v1768<<(uint(v1773)%32)&int32(16) | (v1768<<(uint(int32(4))%32)&int32(64) | (v1768<<(uint(int32(6))%32)&int32(128) | v1768&v1773<<(uint(int32(12))%32))) | v1762
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+20)) = uint16(v1792)
	goto L349
L348:
	;
	goto L349
L349:
	;
	if v1768&int32(16) != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1797 = v1764 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+18)) = uint16(v1797)
	goto L352
L351:
	;
	goto L352
L352:
	;
	v1799 = *(*int32)(unsafe.Add(mBase, uint32(v1672)))
	*(*int32)(unsafe.Add(mBase, uint32(v1759)+4)) = v1799
	*(*uint32)(unsafe.Add(mBase, uint32(v1735)+4)) = uint32(v1671)
	v1803 = int64(base.Ui64(v1671) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1735))) = uint32(v1803)
	v1805 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	F_MarkBufferDirty(m, v1805)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L11
	} else {
		goto L353
	}
L353:
	;
	goto L337
L354:
	;
	F_UnlockReleaseBuffer(m, v1814)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L11
	} else {
		goto L355
	}
L355:
	;
	goto L5
L356:
	;
	v1849 = F_OpenTransientFile(m, v1821+int32(112), int32(65))
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L11
	} else {
		goto L360
	}
L357:
	;
	goto L5
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L11
	} else {
		goto L403
	}
L359:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L11
	} else {
		goto L399
	}
L360:
	;
	if int32(0) <= v1849 {
		goto L361
	} else {
		goto L362
	}
L361:
	;
	v1854 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v1854))) = int32(167772198)
	v1857 = *(*int64)(unsafe.Add(mBase, uint32(v1824)+16))
	v1858 = F_ftruncate(m, v1849, v1857)
	mBase = m.M
	if v1858 != 0 {
		goto L359
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L11
	} else {
		goto L395
	}
L364:
	;
	v1859 = int32(4142700)
	v1860 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v1861 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1860))) = v1861
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1863)+64))
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(v1824)+24))
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v1861
	v1870 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v1870))) = int32(167772196)
	v1876 = v1865 * int32(36)
	v1877 = *(*int64)(unsafe.Add(mBase, uint32(v1824)+16))
	v1878 = F_pwrite(m, v1849, v1864+int32(40), v1876, v1877)
	mBase = m.M
	if v1878 != v1876 {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1881 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v1881 == int32(0) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	goto L367
L367:
	;
	v1906 = int32(4142700)
	v1907 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	v1908 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1907))) = v1908
	v1911 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v1911))) = int32(167772195)
	v1916 = int32(*(*uint8)(unsafe.Add(mBase, _consts[41])))
	if v1916 != int32(1) {
		v1930 = v1908
		goto L377
	} else {
		goto L378
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(51)
	goto L370
L369:
	;
	goto L370
L370:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L11
	} else {
		goto L371
	}
L371:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L11
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+48)) = v1821 + int32(112)
	F_errmsg(m, int32(304496), v1821+int32(48))
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L11
	} else {
		goto L373
	}
L373:
	;
	F_errfinish(m, int32(507593), int32(1122), int32(357348))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L11
	} else {
		goto L374
	}
L374:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L375:
	;
	v1959 = *(*int32)(unsafe.Add(mBase, _consts[39]))
	*(*int32)(unsafe.Add(mBase, uint32(v1959))) = int32(0)
	v1962 = F_CloseTransientFile(m, v1849)
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L11
	} else {
		goto L393
	}
L376:
	;
	if v1930 == int32(0) {
		goto L375
	} else {
		goto L383
	}
L377:
	;
	goto L376
L378:
	;
	goto L379
L379:
	;
	v1921 = F_fsync(m, v1849)
	mBase = m.M
	if v1921 != int32(-1) {
		v1930 = v1921
		goto L377
	} else {
		goto L381
	}
L380:
	;
	v1930 = int32(-1)
	goto L377
L381:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v1925 == int32(27) {
		goto L379
	} else {
		goto L382
	}
L382:
	;
	goto L380
L383:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, _consts[42])))
	if v1936 != 0 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	v1939 = F_errstart(m, v1937, int32(0))
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L11
	} else {
		goto L388
	}
L385:
	;
	v1937 = int32(21)
	goto L387
L386:
	;
	v1937 = int32(23)
	goto L387
L387:
	;
	goto L384
L388:
	;
	if v1939 == int32(0) {
		goto L375
	} else {
		goto L389
	}
L389:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L11
	} else {
		goto L390
	}
L390:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+32)) = v1821 + int32(112)
	F_errmsg(m, int32(305761), v1821+int32(32))
	mBase = m.M
	v1952 = m.ExcPending
	if v1952 != 0 {
		goto L11
	} else {
		goto L391
	}
L391:
	;
	F_errfinish(m, int32(507593), int32(1135), int32(357348))
	mBase = m.M
	v1957 = m.ExcPending
	if v1957 != 0 {
		goto L11
	} else {
		goto L392
	}
L392:
	;
	goto L375
L393:
	;
	if v1962 != 0 {
		goto L358
	} else {
		goto L394
	}
L394:
	;
	m.G0 = v1821 + int32(1136)
	goto L357
L395:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L11
	} else {
		goto L396
	}
L396:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821))) = v1821 + int32(112)
	F_errmsg(m, int32(305492), v1821)
	mBase = m.M
	v1978 = m.ExcPending
	if v1978 != 0 {
		goto L11
	} else {
		goto L397
	}
L397:
	;
	F_errfinish(m, int32(507593), int32(1094), int32(357348))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L11
	} else {
		goto L398
	}
L398:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L399:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1989 = m.ExcPending
	if v1989 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	v1990 = *(*int64)(unsafe.Add(mBase, uint32(v1824)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v1821)+68)) = uint32(v1990)
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+64)) = v1821 + int32(112)
	F_errmsg(m, int32(298569), v1821-int32(-64))
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L11
	} else {
		goto L401
	}
L401:
	;
	F_errfinish(m, int32(507593), int32(1105), int32(357348))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L11
	} else {
		goto L402
	}
L402:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L403:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L11
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1821)+16)) = v1821 + int32(112)
	F_errmsg(m, int32(305556), v1821+int32(16))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L11
	} else {
		goto L405
	}
L405:
	;
	F_errfinish(m, int32(507593), int32(1141), int32(357348))
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L11
	} else {
		goto L406
	}
L406:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L407:
	;
	F_errmsg_internal(m, int32(233040), int32(0))
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L11
	} else {
		goto L408
	}
L408:
	;
	F_errfinish(m, int32(510063), int32(619), int32(83328))
	mBase = m.M
	v2059 = m.ExcPending
	if v2059 != 0 {
		goto L11
	} else {
		goto L409
	}
L409:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L410:
	;
	F_errmsg_internal(m, int32(393243), int32(0))
	mBase = m.M
	v2067 = m.ExcPending
	if v2067 != 0 {
		goto L11
	} else {
		goto L411
	}
L411:
	;
	F_errfinish(m, int32(510063), int32(645), int32(83328))
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L11
	} else {
		goto L412
	}
L412:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L413:
	;
	F_errmsg_internal(m, int32(331680), int32(0))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L11
	} else {
		goto L414
	}
L414:
	;
	F_errfinish(m, int32(510063), int32(648), int32(83328))
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L11
	} else {
		goto L415
	}
L415:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L416:
	;
	F_errmsg_internal(m, int32(242536), int32(0))
	mBase = m.M
	v2094 = m.ExcPending
	if v2094 != 0 {
		goto L11
	} else {
		goto L417
	}
L417:
	;
	F_errfinish(m, int32(510063), int32(1113), int32(458235))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L11
	} else {
		goto L418
	}
L418:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heapgettup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = l0 - int32(-64)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v21 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v101 = v91
	v102 = v92
	v103 = v93
	v104 = v96
	v105 = v95
	goto L24
L2:
	;
	v91 = v87
	v92 = v28
	v93 = v88
	v95 = v46
	v96 = int32(1)
	goto L1
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v24, int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v91 = v5
	v92 = v5
	v93 = v5
	v95 = v5
	v96 = int32(0)
	goto L1
L6:
	;
	return
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v28 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if l1 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(v28^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L8
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v46 = v40 + v28<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v49) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
	v73 = int32(65535)
	v74 = int32(base.Ui32(v68+int32(262120))>>(uint(int32(2))%32)) & v73
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)))
	v79 = (v75 - int32(1)) & v73
	if base.Ui32(v74) < base.Ui32(v79) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v59 = int32(base.Ui32(v49+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	goto L17
L16:
	;
	v59 = int32(0)
	goto L17
L17:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)))
	v61 = int32(1)
	v62 = v60 + v61
	v87 = v59 - v62&int32(65535) + v61
	v88 = v62
	goto L2
L18:
	;
	v81 = v74
	goto L20
L19:
	;
	v81 = v79
	goto L20
L20:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v68) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v85 = v81
	goto L23
L22:
	;
	v85 = int32(0)
	goto L23
L23:
	;
	v87 = v85
	v88 = v85
	goto L2
L24:
	;
	if v104 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	m.G0 = v17 + int32(16)
	return
L26:
	;
	goto L25
L27:
	;
	v101 = v276
	v102 = v277
	v103 = v278
	v104 = int32(0)
	goto L24
L28:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v311, int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L6
	} else {
		goto L68
	}
L29:
	;
	v289 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v289
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v289
	v293 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v293
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v293)
	goto L26
L30:
	;
	F_heap_fetch_next_buffer(m, l0, l1)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if int32(0) < v101 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v115 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_LockBuffer(m, v115, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v122 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v140)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v141) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126+(v122^int32(-1))<<(uint(int32(2))%32))))
	v140 = v132
	goto L36
L38:
	;
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v140 = v134 + v122<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	v149 = int32(base.Ui32(v141+int32(262120)) >> (uint(int32(2)) % 32))
	goto L42
L41:
	;
	v149 = int32(0)
	goto L42
L42:
	;
	if l1 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v152 = int32(1)
	goto L45
L44:
	;
	v152 = v149
	goto L45
L45:
	;
	v101 = v149 & int32(65535)
	v102 = v122
	v103 = v152
	v104 = int32(1)
	v105 = v140
	goto L24
L46:
	;
	v164 = v101
	v166 = v103
	goto L49
L47:
	;
	v276 = v101
	v277 = v102
	v278 = v103
	goto L48
L48:
	;
	F_LockBuffer(m, v277, int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L6
	} else {
		goto L67
	}
L49:
	;
	v180 = v166&int32(65535)<<(uint(int32(2))%32) + (v105 + int32(24)) - int32(4)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	if v181&int32(98304) != int32(32768) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v276 = v268
	v277 = v271
	v278 = v266
	goto L48
L51:
	;
	v266 = l1 + v166
	v267 = int32(1)
	v268 = v164 - v267
	if v267 < v164 {
		v164 = v268
		v166 = v266
		goto L49
	} else {
		goto L66
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v105 + v181&int32(32767)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v166)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v192)
	v195 = int32(base.Ui32(v192) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v190) >> (uint(int32(17)) % 32))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v202 = F_HeapTupleSatisfiesVisibility(m, v20, v200, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, v202, v204, v20, v205, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v202 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	if l3 == int32(0) {
		goto L28
	} else {
		goto L56
	}
L56:
	;
	if l2 == int32(0) {
		goto L28
	} else {
		goto L57
	}
L57:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	v221 = l3
	v222 = l2
	goto L58
L58:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221))))
	if v231&int32(1) != 0 {
		goto L51
	} else {
		goto L60
	}
L59:
	;
	goto L28
L60:
	;
	v234 = int32(*(*int16)(unsafe.Add(mBase, uint32(v221)+4)))
	v237 = F_heap_getattr_1(m, v20, v234, v216, v17+int32(15))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v239 != 0 {
		goto L51
	} else {
		goto L62
	}
L62:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v221)+44))
	v244 = F_FunctionCall2Coll(m, v221+int32(16), v242, v237, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	if v244 == int32(0) {
		goto L51
	} else {
		goto L64
	}
L64:
	;
	v251 = v222 - int32(1)
	if v251 != 0 {
		v221 = v221 + int32(48)
		v222 = v251
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	goto L50
L67:
	;
	goto L27
L68:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)) = uint16(v166)
	goto L26
}
func F_hemdist_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v247 int64
	_ = v247
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int64
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v290 int64
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int64
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int64
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int64
	_ = v332
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v365 int64
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v374 int64
	_ = v374
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v379 int32
	_ = v379
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v386 int64
	_ = v386
	var v387 int32
	_ = v387
	var v390 int64
	_ = v390
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v405 int64
	_ = v405
	var v410 int64
	_ = v410
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int64
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v453 int64
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int64
	_ = v467
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int64
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v483 int64
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int64
	_ = v495
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int64
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int64
	_ = v515
	var v516 int64
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v528 int64
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v537 int64
	_ = v537
	var v538 int32
	_ = v538
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v549 int64
	_ = v549
	var v550 int32
	_ = v550
	var v553 int64
	_ = v553
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v568 int64
	_ = v568
	var v572 int64
	_ = v572
	var v584 int64
	_ = v584
	v4 = int64(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = int32(2)
	v12 = v10 & v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v13&v11 != 0 {
		if v12 != 0 {
			return int32(0)
		} else {
			v19 = l1 + int32(8)
			v20 = int32(3)
			if v20 < l2 {
				v247 = int64(0)
				if l2 < int32(4) {
					v326 = v19
					v327 = l2
					v332 = v247
				} else {
					if v19 != (l1+int32(11))&int32(-4) {
						v326 = v19
						v327 = l2
						v332 = v247
					} else {
						v256 = l2 - int32(4)
						v260 = int32(base.Ui32(v256)>>(uint(int32(2))%32)) + int32(1)
						v262 = v260 & int32(3)
						if base.Ui32(v256) < base.Ui32(int32(12)) {
							v298 = v19
							v299 = l2
							v304 = v247
						} else {
							v268 = v19
							v269 = l2
							v270 = int32(0)
							v274 = v247
							for {
								v275 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
								v278 = *(*int32)(unsafe.Add(mBase, uint32(v268)+8))
								v281 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
								v290 = base.I64_extend_i32_u(base.I32_popcnt(v275)) + (base.I64_extend_i32_u(base.I32_popcnt(v278)) + (base.I64_extend_i32_u(base.I32_popcnt(v281)) + (v274 + base.I64_extend_i32_u(base.I32_popcnt(v284)))))
								v291 = int32(16)
								v292 = v269 - v291
								v294 = v268 + v291
								v296 = v270 + int32(4)
								if v296 != v260&int32(2147483644) {
									v268 = v294
									v269 = v292
									v270 = v296
									v274 = v290
									continue
								} else {
									break
								}
								break
							}
							v298 = v294
							v299 = v292
							v304 = v290
						}
						if v262 == int32(0) {
							v326 = v298
							v327 = v299
							v332 = v304
						} else {
							v309 = v299
							v310 = v298
							v311 = int32(0)
							v314 = v304
							for {
								v315 = int32(4)
								v316 = v309 - v315
								v317 = *(*int32)(unsafe.Add(mBase, uint32(v310)))
								v320 = v314 + base.I64_extend_i32_u(base.I32_popcnt(v317))
								v322 = v310 + v315
								v324 = v311 + int32(1)
								if v324 != v262 {
									v309 = v316
									v310 = v322
									v311 = v324
									v314 = v320
									continue
								} else {
									break
								}
								break
							}
							v326 = v322
							v327 = v316
							v332 = v320
						}
					}
				}
				if v327 == int32(0) {
					v405 = v332
				} else {
					v336 = v327 & int32(3)
					if v336 == int32(0) {
						v359 = v326
						v361 = v327
						v365 = v332
					} else {
						v342 = v327
						v343 = v326
						v344 = int32(0)
						v346 = v332
						for {
							v347 = int32(1)
							v348 = v342 - v347
							v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343))))
							v352 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v349)+uint32(_consts[1053]))))
							v353 = v346 + v352
							v355 = v343 + v347
							v357 = v344 + v347
							if v357 != v336 {
								v342 = v348
								v343 = v355
								v344 = v357
								v346 = v353
								continue
							} else {
								break
							}
							break
						}
						v359 = v355
						v361 = v348
						v365 = v353
					}
					if base.Ui32(v327) < base.Ui32(int32(4)) {
						v405 = v365
					} else {
						v368 = v359
						v370 = v361
						v374 = v365
						for {
							v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+3)))
							v378 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v375)+uint32(_consts[1053]))))
							v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+2)))
							v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v379)+uint32(_consts[1053]))))
							v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+1)))
							v386 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v383)+uint32(_consts[1053]))))
							v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
							v390 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v387)+uint32(_consts[1053]))))
							v394 = v378 + (v382 + (v386 + (v374 + v390)))
							v395 = int32(4)
							v398 = v370 - v395
							if v398 != 0 {
								v368 = v368 + v395
								v370 = v398
								v374 = v394
								continue
							} else {
								break
							}
							break
						}
						v405 = v394
					}
				}
				v584 = v405
			} else {
				if l2 == int32(0) {
					v584 = v4
				} else {
					v27 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v65 = v19
						v67 = v4
					} else {
						v34 = v19
						v36 = v4
						v37 = int32(0)
						for {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
							v45 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1053]))))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
							v49 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_consts[1053]))))
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+2)))
							v53 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v51)+uint32(_consts[1053]))))
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+3)))
							v57 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55)+uint32(_consts[1053]))))
							v58 = v36 + v45 + v49 + v53 + v57
							v59 = int32(4)
							v60 = v34 + v59
							v62 = v37 + v59
							if v62 != l2&int32(-4) {
								v34 = v60
								v36 = v58
								v37 = v62
								continue
							} else {
								break
							}
							break
						}
						v65 = v60
						v67 = v58
					}
					if v27 == int32(0) {
						v584 = v67
					} else {
						v77 = v65
						v78 = int32(0)
						v79 = v67
						for {
							v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
							v88 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v85)+uint32(_consts[1053]))))
							v89 = v79 + v88
							v90 = int32(1)
							v93 = v78 + v90
							if v93 != v27 {
								v77 = v77 + v90
								v78 = v93
								v79 = v89
								continue
							} else {
								break
							}
							break
						}
						v584 = v89
					}
				}
			}
			return l2<<(uint(v20)%32) - base.I32_wrap_i64(v584)
		}
	} else {
		if v12 != 0 {
			v96 = l0 + int32(8)
			v97 = int32(3)
			if v97 < l2 {
				v410 = int64(0)
				if l2 < int32(4) {
					v489 = v96
					v490 = l2
					v495 = v410
				} else {
					if v96 != (l0+int32(11))&int32(-4) {
						v489 = v96
						v490 = l2
						v495 = v410
					} else {
						v419 = l2 - int32(4)
						v423 = int32(base.Ui32(v419)>>(uint(int32(2))%32)) + int32(1)
						v425 = v423 & int32(3)
						if base.Ui32(v419) < base.Ui32(int32(12)) {
							v461 = v96
							v462 = l2
							v467 = v410
						} else {
							v431 = v96
							v432 = l2
							v433 = int32(0)
							v437 = v410
							for {
								v438 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
								v441 = *(*int32)(unsafe.Add(mBase, uint32(v431)+8))
								v444 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
								v447 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
								v453 = base.I64_extend_i32_u(base.I32_popcnt(v438)) + (base.I64_extend_i32_u(base.I32_popcnt(v441)) + (base.I64_extend_i32_u(base.I32_popcnt(v444)) + (v437 + base.I64_extend_i32_u(base.I32_popcnt(v447)))))
								v454 = int32(16)
								v455 = v432 - v454
								v457 = v431 + v454
								v459 = v433 + int32(4)
								if v459 != v423&int32(2147483644) {
									v431 = v457
									v432 = v455
									v433 = v459
									v437 = v453
									continue
								} else {
									break
								}
								break
							}
							v461 = v457
							v462 = v455
							v467 = v453
						}
						if v425 == int32(0) {
							v489 = v461
							v490 = v462
							v495 = v467
						} else {
							v472 = v462
							v473 = v461
							v474 = int32(0)
							v477 = v467
							for {
								v478 = int32(4)
								v479 = v472 - v478
								v480 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
								v483 = v477 + base.I64_extend_i32_u(base.I32_popcnt(v480))
								v485 = v473 + v478
								v487 = v474 + int32(1)
								if v487 != v425 {
									v472 = v479
									v473 = v485
									v474 = v487
									v477 = v483
									continue
								} else {
									break
								}
								break
							}
							v489 = v485
							v490 = v479
							v495 = v483
						}
					}
				}
				if v490 == int32(0) {
					v568 = v495
				} else {
					v499 = v490 & int32(3)
					if v499 == int32(0) {
						v522 = v489
						v524 = v490
						v528 = v495
					} else {
						v505 = v490
						v506 = v489
						v507 = int32(0)
						v509 = v495
						for {
							v510 = int32(1)
							v511 = v505 - v510
							v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
							v515 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v512)+uint32(_consts[1053]))))
							v516 = v509 + v515
							v518 = v506 + v510
							v520 = v507 + v510
							if v520 != v499 {
								v505 = v511
								v506 = v518
								v507 = v520
								v509 = v516
								continue
							} else {
								break
							}
							break
						}
						v522 = v518
						v524 = v511
						v528 = v516
					}
					if base.Ui32(v490) < base.Ui32(int32(4)) {
						v568 = v528
					} else {
						v531 = v522
						v533 = v524
						v537 = v528
						for {
							v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+3)))
							v541 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v538)+uint32(_consts[1053]))))
							v542 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+2)))
							v545 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v542)+uint32(_consts[1053]))))
							v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531)+1)))
							v549 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v546)+uint32(_consts[1053]))))
							v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
							v553 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v550)+uint32(_consts[1053]))))
							v557 = v541 + (v545 + (v549 + (v537 + v553)))
							v558 = int32(4)
							v561 = v533 - v558
							if v561 != 0 {
								v531 = v531 + v558
								v533 = v561
								v537 = v557
								continue
							} else {
								break
							}
							break
						}
						v568 = v557
					}
				}
				v572 = v568
			} else {
				if l2 == int32(0) {
					v572 = v4
				} else {
					v104 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v142 = v96
						v144 = v4
					} else {
						v111 = v96
						v113 = v4
						v114 = int32(0)
						for {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
							v122 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[1053]))))
							v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
							v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[1053]))))
							v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+2)))
							v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1053]))))
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+3)))
							v134 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[1053]))))
							v135 = v113 + v122 + v126 + v130 + v134
							v136 = int32(4)
							v137 = v111 + v136
							v139 = v114 + v136
							if v139 != l2&int32(-4) {
								v111 = v137
								v113 = v135
								v114 = v139
								continue
							} else {
								break
							}
							break
						}
						v142 = v137
						v144 = v135
					}
					if v104 == int32(0) {
						v572 = v144
					} else {
						v154 = v142
						v155 = int32(0)
						v156 = v144
						for {
							v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154))))
							v165 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_consts[1053]))))
							v166 = v156 + v165
							v167 = int32(1)
							v170 = v155 + v167
							if v170 != v104 {
								v154 = v154 + v167
								v155 = v170
								v156 = v166
								continue
							} else {
								break
							}
							break
						}
						v572 = v166
					}
				}
			}
			return l2<<(uint(v97)%32) - base.I32_wrap_i64(v572)
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v176 = int32(8)
				v177 = l1 + v176
				v179 = l0 + v176
				v180 = int32(1)
				if l2 == v180 {
					v184 = int32(0)
					v224 = v184
					v225 = v184
				} else {
					v188 = int32(0)
					v191 = v188
					v192 = v188
					v195 = int32(0)
					for {
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v191))))
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191+v177))))
						v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201^v203)+uint32(_consts[1053]))))
						v209 = v191 | int32(1)
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v209))))
						v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v209))))
						v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211^v213)+uint32(_consts[1053]))))
						v217 = v192 + v206 + v216
						v218 = int32(2)
						v219 = v191 + v218
						v221 = v195 + v218
						if v221 != l2&int32(2147483646) {
							v191 = v219
							v192 = v217
							v195 = v221
							continue
						} else {
							break
						}
						break
					}
					v224 = v219
					v225 = v217
				}
				if l2&v180 != 0 {
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v224))))
					v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224+v177))))
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233^v235)+uint32(_consts[1053]))))
					v241 = v225 + v239
				} else {
					v241 = v225
				}
				return v241
			}
		}
	}
}
func F_hemdistcache_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int64
	_ = v35
	var v42 int32
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v162 int32
	_ = v162
	var v165 int64
	_ = v165
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v245 int64
	_ = v245
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int64
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int64
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int64
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int64
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int64
	_ = v330
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int64
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v363 int64
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v376 int64
	_ = v376
	var v377 int32
	_ = v377
	var v380 int64
	_ = v380
	var v381 int32
	_ = v381
	var v384 int64
	_ = v384
	var v385 int32
	_ = v385
	var v388 int64
	_ = v388
	var v392 int64
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v403 int64
	_ = v403
	var v408 int64
	_ = v408
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v451 int64
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int64
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int64
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int64
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int64
	_ = v493
	var v497 int32
	_ = v497
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int64
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int64
	_ = v513
	var v514 int64
	_ = v514
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v526 int64
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v539 int64
	_ = v539
	var v540 int32
	_ = v540
	var v543 int64
	_ = v543
	var v544 int32
	_ = v544
	var v547 int64
	_ = v547
	var v548 int32
	_ = v548
	var v551 int64
	_ = v551
	var v555 int64
	_ = v555
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v566 int64
	_ = v566
	var v570 int64
	_ = v570
	var v584 int64
	_ = v584
	v4 = int64(0)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(1) {
		if v10&int32(1) != 0 {
			return int32(0)
		} else {
			v18 = int32(3)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v18 < l2 {
				v245 = int64(0)
				if l2 < int32(4) {
					v324 = v20
					v325 = l2
					v330 = v245
				} else {
					if v20 != (v20+int32(3))&int32(-4) {
						v324 = v20
						v325 = l2
						v330 = v245
					} else {
						v254 = l2 - int32(4)
						v258 = int32(base.Ui32(v254)>>(uint(int32(2))%32)) + int32(1)
						v260 = v258 & int32(3)
						if base.Ui32(v254) < base.Ui32(int32(12)) {
							v296 = v20
							v297 = l2
							v302 = v245
						} else {
							v266 = v20
							v267 = l2
							v268 = int32(0)
							v272 = v245
							for {
								v273 = *(*int32)(unsafe.Add(mBase, uint32(v266)+12))
								v276 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
								v279 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
								v282 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
								v288 = base.I64_extend_i32_u(base.I32_popcnt(v273)) + (base.I64_extend_i32_u(base.I32_popcnt(v276)) + (base.I64_extend_i32_u(base.I32_popcnt(v279)) + (v272 + base.I64_extend_i32_u(base.I32_popcnt(v282)))))
								v289 = int32(16)
								v290 = v267 - v289
								v292 = v266 + v289
								v294 = v268 + int32(4)
								if v294 != v258&int32(2147483644) {
									v266 = v292
									v267 = v290
									v268 = v294
									v272 = v288
									continue
								} else {
									break
								}
								break
							}
							v296 = v292
							v297 = v290
							v302 = v288
						}
						if v260 == int32(0) {
							v324 = v296
							v325 = v297
							v330 = v302
						} else {
							v307 = v297
							v308 = v296
							v309 = int32(0)
							v312 = v302
							for {
								v313 = int32(4)
								v314 = v307 - v313
								v315 = *(*int32)(unsafe.Add(mBase, uint32(v308)))
								v318 = v312 + base.I64_extend_i32_u(base.I32_popcnt(v315))
								v320 = v308 + v313
								v322 = v309 + int32(1)
								if v322 != v260 {
									v307 = v314
									v308 = v320
									v309 = v322
									v312 = v318
									continue
								} else {
									break
								}
								break
							}
							v324 = v320
							v325 = v314
							v330 = v318
						}
					}
				}
				if v325 == int32(0) {
					v403 = v330
				} else {
					v334 = v325 & int32(3)
					if v334 == int32(0) {
						v357 = v324
						v359 = v325
						v363 = v330
					} else {
						v340 = v325
						v341 = v324
						v342 = int32(0)
						v344 = v330
						for {
							v345 = int32(1)
							v346 = v340 - v345
							v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341))))
							v350 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v347)+uint32(_consts[1053]))))
							v351 = v344 + v350
							v353 = v341 + v345
							v355 = v342 + v345
							if v355 != v334 {
								v340 = v346
								v341 = v353
								v342 = v355
								v344 = v351
								continue
							} else {
								break
							}
							break
						}
						v357 = v353
						v359 = v346
						v363 = v351
					}
					if base.Ui32(v325) < base.Ui32(int32(4)) {
						v403 = v363
					} else {
						v366 = v357
						v368 = v359
						v372 = v363
						for {
							v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+3)))
							v376 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v373)+uint32(_consts[1053]))))
							v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+2)))
							v380 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v377)+uint32(_consts[1053]))))
							v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+1)))
							v384 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v381)+uint32(_consts[1053]))))
							v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
							v388 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v385)+uint32(_consts[1053]))))
							v392 = v376 + (v380 + (v384 + (v372 + v388)))
							v393 = int32(4)
							v396 = v368 - v393
							if v396 != 0 {
								v366 = v366 + v393
								v368 = v396
								v372 = v392
								continue
							} else {
								break
							}
							break
						}
						v403 = v392
					}
				}
				v584 = v403
			} else {
				if l2 == int32(0) {
					v584 = v4
				} else {
					v26 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v63 = v20
						v66 = v4
					} else {
						v32 = v20
						v33 = int32(0)
						v35 = v4
						for {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
							v44 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1053]))))
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
							v48 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[1053]))))
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+2)))
							v52 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v50)+uint32(_consts[1053]))))
							v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+3)))
							v56 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v54)+uint32(_consts[1053]))))
							v57 = v35 + v44 + v48 + v52 + v56
							v58 = int32(4)
							v59 = v32 + v58
							v61 = v33 + v58
							if v61 != l2&int32(-4) {
								v32 = v59
								v33 = v61
								v35 = v57
								continue
							} else {
								break
							}
							break
						}
						v63 = v59
						v66 = v57
					}
					if v26 == int32(0) {
						v584 = v66
					} else {
						v75 = v63
						v77 = int32(0)
						v78 = v66
						for {
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v87 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1053]))))
							v88 = v78 + v87
							v89 = int32(1)
							v92 = v77 + v89
							if v92 != v26 {
								v75 = v75 + v89
								v77 = v92
								v78 = v88
								continue
							} else {
								break
							}
							break
						}
						v584 = v88
					}
				}
			}
			return l2<<(uint(v18)%32) + (base.I32_wrap_i64(v584) ^ int32(-1))
		}
	} else {
		if v10&int32(1) != 0 {
			v96 = int32(3)
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v96 < l2 {
				v408 = int64(0)
				if l2 < int32(4) {
					v487 = v98
					v488 = l2
					v493 = v408
				} else {
					if v98 != (v98+int32(3))&int32(-4) {
						v487 = v98
						v488 = l2
						v493 = v408
					} else {
						v417 = l2 - int32(4)
						v421 = int32(base.Ui32(v417)>>(uint(int32(2))%32)) + int32(1)
						v423 = v421 & int32(3)
						if base.Ui32(v417) < base.Ui32(int32(12)) {
							v459 = v98
							v460 = l2
							v465 = v408
						} else {
							v429 = v98
							v430 = l2
							v431 = int32(0)
							v435 = v408
							for {
								v436 = *(*int32)(unsafe.Add(mBase, uint32(v429)+12))
								v439 = *(*int32)(unsafe.Add(mBase, uint32(v429)+8))
								v442 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
								v445 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
								v451 = base.I64_extend_i32_u(base.I32_popcnt(v436)) + (base.I64_extend_i32_u(base.I32_popcnt(v439)) + (base.I64_extend_i32_u(base.I32_popcnt(v442)) + (v435 + base.I64_extend_i32_u(base.I32_popcnt(v445)))))
								v452 = int32(16)
								v453 = v430 - v452
								v455 = v429 + v452
								v457 = v431 + int32(4)
								if v457 != v421&int32(2147483644) {
									v429 = v455
									v430 = v453
									v431 = v457
									v435 = v451
									continue
								} else {
									break
								}
								break
							}
							v459 = v455
							v460 = v453
							v465 = v451
						}
						if v423 == int32(0) {
							v487 = v459
							v488 = v460
							v493 = v465
						} else {
							v470 = v460
							v471 = v459
							v472 = int32(0)
							v475 = v465
							for {
								v476 = int32(4)
								v477 = v470 - v476
								v478 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
								v481 = v475 + base.I64_extend_i32_u(base.I32_popcnt(v478))
								v483 = v471 + v476
								v485 = v472 + int32(1)
								if v485 != v423 {
									v470 = v477
									v471 = v483
									v472 = v485
									v475 = v481
									continue
								} else {
									break
								}
								break
							}
							v487 = v483
							v488 = v477
							v493 = v481
						}
					}
				}
				if v488 == int32(0) {
					v566 = v493
				} else {
					v497 = v488 & int32(3)
					if v497 == int32(0) {
						v520 = v487
						v522 = v488
						v526 = v493
					} else {
						v503 = v488
						v504 = v487
						v505 = int32(0)
						v507 = v493
						for {
							v508 = int32(1)
							v509 = v503 - v508
							v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
							v513 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v510)+uint32(_consts[1053]))))
							v514 = v507 + v513
							v516 = v504 + v508
							v518 = v505 + v508
							if v518 != v497 {
								v503 = v509
								v504 = v516
								v505 = v518
								v507 = v514
								continue
							} else {
								break
							}
							break
						}
						v520 = v516
						v522 = v509
						v526 = v514
					}
					if base.Ui32(v488) < base.Ui32(int32(4)) {
						v566 = v526
					} else {
						v529 = v520
						v531 = v522
						v535 = v526
						for {
							v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+3)))
							v539 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v536)+uint32(_consts[1053]))))
							v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+2)))
							v543 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v540)+uint32(_consts[1053]))))
							v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529)+1)))
							v547 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v544)+uint32(_consts[1053]))))
							v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
							v551 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v548)+uint32(_consts[1053]))))
							v555 = v539 + (v543 + (v547 + (v535 + v551)))
							v556 = int32(4)
							v559 = v531 - v556
							if v559 != 0 {
								v529 = v529 + v556
								v531 = v559
								v535 = v555
								continue
							} else {
								break
							}
							break
						}
						v566 = v555
					}
				}
				v570 = v566
			} else {
				if l2 == int32(0) {
					v570 = v4
				} else {
					v104 = l2 & int32(3)
					if base.Ui32(l2) < base.Ui32(int32(4)) {
						v141 = v98
						v144 = v4
					} else {
						v110 = v98
						v111 = int32(0)
						v113 = v4
						for {
							v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
							v122 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_consts[1053]))))
							v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
							v126 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v124)+uint32(_consts[1053]))))
							v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+2)))
							v130 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1053]))))
							v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+3)))
							v134 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v132)+uint32(_consts[1053]))))
							v135 = v113 + v122 + v126 + v130 + v134
							v136 = int32(4)
							v137 = v110 + v136
							v139 = v111 + v136
							if v139 != l2&int32(-4) {
								v110 = v137
								v111 = v139
								v113 = v135
								continue
							} else {
								break
							}
							break
						}
						v141 = v137
						v144 = v135
					}
					if v104 == int32(0) {
						v570 = v144
					} else {
						v153 = v141
						v155 = int32(0)
						v156 = v144
						for {
							v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
							v165 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v162)+uint32(_consts[1053]))))
							v166 = v156 + v165
							v167 = int32(1)
							v170 = v155 + v167
							if v170 != v104 {
								v153 = v153 + v167
								v155 = v170
								v156 = v166
								continue
							} else {
								break
							}
							break
						}
						v570 = v166
					}
				}
			}
			return l2<<(uint(v96)%32) + (base.I32_wrap_i64(v570) ^ int32(-1))
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v178 = int32(1)
				if l2 == v178 {
					v182 = int32(0)
					v221 = v182
					v223 = v182
				} else {
					v186 = int32(0)
					v188 = v186
					v190 = v186
					v193 = int32(0)
					for {
						v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v177))))
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188+v176))))
						v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199^v201)+uint32(_consts[1053]))))
						v207 = v188 | int32(1)
						v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v207))))
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v177))))
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209^v211)+uint32(_consts[1053]))))
						v215 = v190 + v204 + v214
						v216 = int32(2)
						v217 = v188 + v216
						v219 = v193 + v216
						if v219 != l2&int32(2147483646) {
							v188 = v217
							v190 = v215
							v193 = v219
							continue
						} else {
							break
						}
						break
					}
					v221 = v217
					v223 = v215
				}
				if l2&v178 != 0 {
					v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v176))))
					v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v177))))
					v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231^v233)+uint32(_consts[1053]))))
					v239 = v223 + v237
				} else {
					v239 = v223
				}
				return v239
			}
		}
	}
}
func F_hmac_block_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v4 = m.T0[v3].(func(*base.Module, int32) int32)(m, v2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_hs_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(5470), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_htonl(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
