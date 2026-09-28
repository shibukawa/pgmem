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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_HaveRegisteredOrActiveSnapshot[0]))
	if v3 != 0 {
		return int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_HaveRegisteredOrActiveSnapshot[1]))
		v8 = int32(0)
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_HaveRegisteredOrActiveSnapshot[2]))
		if base.B2i32(v7 == v8)|base.B2i32(v11 == v8) != 0 {
			return base.B2i32(v11 != int32(0))
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			if v15 != 0 {
				return base.B2i32(v11 != int32(0))
			} else {
				return int32(0)
			}
		}
	}
}
func F_hamming_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
			if v17 != v18 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
						F_errmsg(m, int32(_a_F_hamming_distance_0), v7)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_hamming_distance_1), int32(39), int32(_a_F_hamming_distance_2))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v42 = int32(8)
				v50 = *(*int32)(unsafe.Add(mBase, _c_F_hamming_distance[0]))
				v51 = m.T0[v50].(func(*base.Module, int32, int32, int32, int64) int64)(m, int32(base.Ui32(v39)>>(uint(int32(2))%32))-v42, v10+v42, v15+v42, int64(0))
				mBase = m.M
				m.G0 = v7 + int32(16)
				return base.I64_reinterpret_f64(base.F64_convert_i64_u(v51))
			}
		}
	}
}
func F_handle_pm_child_exit_signal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_handle_pm_child_exit_signal_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_handle_pm_child_exit_signal_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_handle_pm_reload_request_signal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[0])) = int32(1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[1]))
	v8 = int32(0)
	v11 = base.AtomicRmwOr32(m, v8, int32(_a_F_handle_pm_reload_request_signal_0), v8)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(1)
	v15 = int32(0)
	v18 = base.AtomicRmwOr32(m, v15, int32(_a_F_handle_pm_reload_request_signal_0), v15)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v19 == v15 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	if v22 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[2]))
	if v26 == v22 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[3]))
	if v33 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v56 = F_pgmem_kill(m, v22, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v30 + int32(16)
	goto L1
L10:
	;
	v36 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+15)) = uint8(v36)
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[4]))
	v44 = F_write(m, v40, v30+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v44 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[5]))
	if v48 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v75 int64
	_ = v75
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
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
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	v1 = l0
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[1]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v19 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v382
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v47 = F_pq_getmsgint(m, l1, int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L19
	}
L3:
	;
	v22 = int32(4)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v23 == v22 {
		v41 = int32(0)
		v42 = v22
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = F_pa_find_worker(m, v16)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return int32(0)
L8:
	;
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+12)))
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v35 = int32(0)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[2])))
	if v38 == v35 {
		v382 = v35
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v34 = int32(3)
	goto L14
L13:
	;
	v34 = int32(2)
	goto L14
L14:
	;
	v41 = v27
	v42 = v34
	goto L2
L15:
	;
	v41 = v35
	v42 = int32(1)
	goto L2
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+7)) = uint8(v1)
	v351 = int32(1)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v352 - v353 + v351
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v359, v13, int32(4))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L87
	}
L17:
	;
	v309 = int32(_a_F_handle_streamed_transaction_0)
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	v311 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v302+v310<<(uint(v311)%32)))) = v47
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	v321 = v302 + v318<<(uint(v311)%32)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v316)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v321+v311))) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v316)+40))
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v316)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v321+int32(8)))) = v328 + v329
	goto L86
L18:
	;
	v285 = int32(128)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5])) = v285
	v287 = int32(_a_F_handle_streamed_transaction_1)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6]))
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v291
	v295 = F_palloc_mul(m, int32(16), v285)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L7
	} else {
		goto L85
	}
L19:
	;
	if v47 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	switch v42 - int32(2) {
	case 0:
		goto L25
	case 1:
		goto L24
	case 2:
		goto L23
	default:
		goto L26
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L7
	} else {
		goto L81
	}
L23:
	;
	v143 = int32(_a_F_handle_streamed_transaction_2)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[8])) = v145 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
	v151 = m.G0
	v153 = v151 - int32(96)
	m.G0 = v153
	if v150 == v47 {
		goto L44
	} else {
		goto L45
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)) = uint8(v1)
	v113 = v44 - v43
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v113 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v118, v13+int32(8), int32(4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L41
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v102 = F_pa_send_data(m, v41, v100, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L36
	}
L26:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
	if v52 == v47 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[9]))
	if v55 == v47 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[9])) = v47
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	if v62 == int32(0) {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v75 = base.I64_extend_i32_u(v62)
	goto L30
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58+base.I32_wrap_i64(v75)<<(uint(int32(4))%32)-int32(16))))
	if v82 == v47 {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5]))
	if v62 != v91 {
		v302 = v58
		goto L17
	} else {
		goto L34
	}
L32:
	;
	if base.B2i32(v75 < int64(2)) == int32(0) {
		v75 = v75 - int64(1)
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v95 = v62 << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5])) = v95
	v98 = F_repalloc_mul(m, v58, int32(16), v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v302 = v98
	goto L17
L36:
	;
	if v102 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v382 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
	goto L1
L38:
	;
	goto L39
L39:
	;
	F_pa_switch_to_partial_serialize(m, v41, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	goto L24
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v125, v13+int32(15), int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v113
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v133, v43+v45, v113)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v382 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
	goto L1
L44:
	;
	m.G0 = v153 + int32(96)
	v382 = int32(0)
	goto L1
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11]))
	v158 = int32(0)
	if v157 == v158 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v196 != 0 {
		goto L44
	} else {
		goto L59
	}
L47:
	;
	v196 = int32(0)
	goto L46
L48:
	;
	goto L49
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v164 <= int32(0) {
		v190 = v158
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v196 = v190
	goto L46
L51:
	;
	v167 = int32(0)
	if v167 < v164 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v170 = v164
	goto L54
L53:
	;
	v170 = v167
	goto L54
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v173 = int32(0)
	goto L55
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v171+v173<<(uint(int32(2))%32))))
	v182 = base.B2i32(v181 == v47)
	if v181 == v47 {
		v190 = v182
		goto L50
	} else {
		goto L57
	}
L56:
	;
	v190 = v182
	goto L50
L57:
	;
	v184 = v173 + int32(1)
	if v184 != v170 {
		v173 = v184
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[12]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v47
	v203 = v153 + int32(32)
	v208 = F_pg_snprintf(m, v203, int32(64), int32(_a_F_handle_streamed_transaction_3), v153+int32(16))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v212 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	if v212 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v203
	F_errmsg_internal(m, int32(_a_F_handle_streamed_transaction_4), v153)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L7
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[13]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	goto L67
L65:
	;
	F_errfinish(m, int32(_a_F_handle_streamed_transaction_5), int32(1392), int32(_a_F_handle_streamed_transaction_6))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v225)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[13]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	goto L71
L69:
	;
	goto L70
L70:
	;
	F_DefineSavepoint(m, v153+int32(32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L7
	} else {
		goto L78
	}
L71:
	;
	if base.B2i32(v232 == int32(2)) == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_BeginTransactionBlock(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L7
	} else {
		goto L77
	}
L77:
	;
	goto L70
L78:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L7
	} else {
		goto L79
	}
L79:
	;
	v249 = int32(_a_F_handle_streamed_transaction_1)
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6]))
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v253
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11]))
	v257 = F_lappend_xid(m, v256, v47)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L7
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v250
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11])) = v257
	goto L44
L81:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L7
	} else {
		goto L82
	}
L82:
	;
	F_errmsg_internal(m, int32(_a_F_handle_streamed_transaction_7), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_handle_streamed_transaction_8), int32(817), int32(_a_F_handle_streamed_transaction_9))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L7
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v288
	v302 = v295
	goto L17
L86:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[10])) = v302
	v334 = int32(_a_F_handle_streamed_transaction_0)
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4])) = v336 + int32(1)
	goto L16
L87:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v364, v13+int32(7), int32(1))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v372 = v370 - v371
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v372
	v375 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_BufFileWrite(m, v375, v371+v376, v372)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v382 = v351
	goto L1
}
func F_has_largeobject_privilege_name_id(m *base.Module, l0 int32) int64 {
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
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_get_role_oid_or_public(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v16 = F_convert_any_priv_string(m, v13, int32(_a_F_has_largeobject_privilege_name_id_0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				if v16&int64(4) == int64(0) {
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_name_id[0]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
					v25 = v24
				} else {
					v25 = int32(0)
				}
				v26 = F_LargeObjectExistsWithSnapshot(m, v11, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int64(0)
				} else {
					if v26 != 0 {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_name_id[1])))
						if v30 != 0 {
							v40 = int64(1)
							return v40
						} else {
							v31 = F_pg_largeobject_aclcheck_snapshot(m, v11, v7, v16, v25)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v31 == int32(0)))
							}
						}
					} else {
						v37 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v37)
						v40 = int64(0)
						return v40
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
	var v10 int32
	_ = v10
	v4 = int32(1)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
	if v5 != 0 {
		v10 = v4
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
		if v6 != 0 {
			v10 = v4
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
			v10 = base.B2i32(v7 != int32(0))
		}
	}
	return v10
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
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
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
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
	v92 = F_ExecFetchSlotMinimalTuple(m, v86, v12+int32(11))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L22
	}
L2:
	;
	v86 = l2
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
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	m.T0[v21].(func(*base.Module, int32, int32))(m, l2, v17)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	m.T0[v25].(func(*base.Module, int32))(m, v16)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if int32(0) < v29 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v76 = v74 & int32(_a_F_hashagg_spill_tuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v76)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)) = uint16(v79)
	goto L21
L14:
	;
	v41 = int32(1)
	v43 = v36 + v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v45 = F_bms_is_member(m, v43, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v48 = v36 << (uint(int32(3)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v51+v48)))
	*(*int64)(unsafe.Add(mBase, uint32(v48+v49))) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v36))))
	v58 = v57
	goto L19
L18:
	;
	v58 = v41
	goto L19
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v59+v36))) = uint8(v58)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v43 < v63 {
		v36 = v43
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v86 = v16
	goto L1
L22:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v95 <= int32(31) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v101 = int32(base.Ui32(v98&l3) >> (uint(v95) % 32))
	goto L25
L24:
	;
	v101 = int32(0)
	goto L25
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v105 = v102 + v101<<(uint(int32(3))%32)
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)))
	*(*int64)(unsafe.Add(mBase, uint32(v105))) = v106 + int64(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v111 = int32(24)
	v113 = v110 + v101*v111
	v118 = int32(711645284)
	v121 = l3 - int32(1636608428) ^ v118 - int32(1455628627)
	v126 = v121 ^ int32(-1636608428) - base.I32_rotl(v121, int32(25))
	v131 = v126 ^ v118 - base.I32_rotl(v126, int32(16))
	v135 = v131 ^ v121 - base.I32_rotl(v131, int32(4))
	v139 = v135 ^ v126 - base.I32_rotl(v135, int32(14))
	v143 = v139 ^ v131 - base.I32_rotl(v139, v111)
	goto L26
L26:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v113)+16))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v149 = int32(32) - v148
	v151 = v146 + int32(base.Ui32(v143)>>(uint(v149)%32))
	v152 = v143 << (uint(v148) % 32)
	if v152 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v101<<(uint(int32(2))%32))))
	F_LogicalTapeWrite(m, v180, v12+int32(12), int32(4))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L8
	} else {
		goto L37
	}
L28:
	;
	v159 = int32(32) - (base.I32_clz(v152) ^ int32(31))
	v160 = int32(255)
	if base.Ui32(v149&v160) < base.Ui32(v159&v160) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v169 = v149 + int32(1)
	goto L30
L30:
	;
	v171 = v169 & int32(255)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if base.Ui32(v172) < base.Ui32(v171) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v165 = v149 + int32(1)
	goto L33
L32:
	;
	v165 = v159
	goto L33
L33:
	;
	v169 = v165
	goto L30
L34:
	;
	v174 = v171
	goto L36
L35:
	;
	v174 = v172
	goto L36
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v174)
	goto L27
L37:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	F_LogicalTapeWrite(m, v180, v92, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v189 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pfree(m, v92)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
func F_hashboolextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v5 == v3 {
		v12 = int32(-1636608428)
		v51 = v12
		v52 = v12
		v55 = int32(0)
	} else {
		v15 = base.I32_wrap_i64(v5)
		v17 = v15 + int32(1021750440)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v28 = v15 - v22 - int32(1636608428) ^ base.I32_rotl(v22, int32(6))
		v32 = v17 - v28 ^ base.I32_rotl(v28, int32(8))
		v33 = v22 + v17
		v34 = v28 + v33
		v35 = v32 + v34
		v39 = v33 - v32 ^ base.I32_rotl(v32, int32(16))
		v43 = v34 - v39 ^ base.I32_rotl(v39, int32(19))
		v48 = v39 + v35
		v49 = v43 + v48
		v51 = v49
		v52 = v48
		v55 = v35 - v43 ^ base.I32_rotl(v43, int32(4)) ^ v49
	}
	v56 = int32(14)
	v58 = v55 - base.I32_rotl(v51, v56)
	v63 = v58 ^ (base.B2i32(v2 != v3) + v52) - base.I32_rotl(v58, int32(11))
	v67 = v51 ^ v63 - base.I32_rotl(v63, int32(25))
	v71 = v67 ^ v58 - base.I32_rotl(v67, int32(16))
	v75 = v71 ^ v63 - base.I32_rotl(v71, int32(4))
	v79 = v75 ^ v67 - base.I32_rotl(v75, v56)
	return base.I64_extend_i32_u(v79)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v79^v71-base.I32_rotl(v79, int32(24)))
}
func F_hashfloat4extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 float32
	_ = v11
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
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
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
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
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v341 int64
	_ = v341
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.F32_ne(v11, float32(0)) != 0 {
		v15 = base.F64_promote_f32(v11)
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v15)&int64(9223372036854775807)) {
			v21 = math.Float64frombits(uint64(0x7ff8000000000000))
		} else {
			v21 = v15
		}
		*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v21
		v24 = v8 + int32(8)
		v31 = int32(-1636608424)
		if v10 == int64(0) {
			v68 = v31
			v70 = v31
			v72 = v31
		} else {
			v34 = base.I32_wrap_i64(v10)
			v36 = v34 + int32(1021750448)
			v40 = int32(4)
			v42 = base.I32_wrap_i64(int64(base.Ui64(v10)>>(uint(int64(32))%64))) ^ base.I32_rotl(v31, v40)
			v46 = v31 + v34 - v42 ^ base.I32_rotl(v42, int32(6))
			v50 = v36 - v46 ^ base.I32_rotl(v46, int32(8))
			v51 = v36 + v42
			v52 = v46 + v51
			v53 = v50 + v52
			v57 = v51 - v50 ^ base.I32_rotl(v50, int32(16))
			v61 = v52 - v57 ^ base.I32_rotl(v57, int32(19))
			v66 = v53 + v57
			v68 = v66
			v70 = v53 - v61 ^ base.I32_rotl(v61, v40)
			v72 = v61 + v66
		}
		if v24&int32(3) != 0 {
			switch int32(7) {
			case 0:
				v296 = v68
				v297 = v72
				v298 = v70
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 1:
				v289 = v68
				v290 = v72
				v291 = v70
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 2:
				v282 = v68
				v283 = v72
				v284 = v70
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 3:
				v276 = v72
				v277 = v70
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 4:
				v272 = v72
				v273 = v70
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 5:
				v266 = v72
				v267 = v70
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 6:
				v260 = v72
				v261 = v70
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v266 = v262<<(uint(int32(16))%32) + v260
				v267 = v261
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 7:
				v255 = v70
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)))
				v260 = v256<<(uint(int32(24))%32) + v72
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v266 = v262<<(uint(int32(16))%32) + v260
				v267 = v261
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 8:
				v250 = v70
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v255 = v251<<(uint(int32(8))%32) + v250
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)))
				v260 = v256<<(uint(int32(24))%32) + v72
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v266 = v262<<(uint(int32(16))%32) + v260
				v267 = v261
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 9:
				v245 = v70
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
				v250 = v246<<(uint(int32(16))%32) + v245
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v255 = v251<<(uint(int32(8))%32) + v250
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)))
				v260 = v256<<(uint(int32(24))%32) + v72
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v266 = v262<<(uint(int32(16))%32) + v260
				v267 = v261
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			case 10:
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+10)))
				v245 = v241<<(uint(int32(24))%32) + v70
				v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
				v250 = v246<<(uint(int32(16))%32) + v245
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v255 = v251<<(uint(int32(8))%32) + v250
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+7)))
				v260 = v256<<(uint(int32(24))%32) + v72
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v266 = v262<<(uint(int32(16))%32) + v260
				v267 = v261
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v272 = v268<<(uint(int32(8))%32) + v266
				v273 = v267
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v276 = v272 + v274
				v277 = v273
				v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+3)))
				v282 = v278<<(uint(int32(24))%32) + v68
				v283 = v276
				v284 = v277
				v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v289 = v285<<(uint(int32(16))%32) + v282
				v290 = v283
				v291 = v284
				v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v296 = v292<<(uint(int32(8))%32) + v289
				v297 = v290
				v298 = v291
				v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v296 + v299
				v305 = v297
				v306 = v298
			default:
				v304 = v68
				v305 = v72
				v306 = v70
			}
		} else {
			switch int32(7) {
			case 0:
				v238 = v68
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v238 + v239
				v305 = v72
				v306 = v70
			case 1:
				v233 = v68
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v238 = v234<<(uint(int32(8))%32) + v233
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v238 + v239
				v305 = v72
				v306 = v70
			case 2:
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
				v233 = v229<<(uint(int32(16))%32) + v68
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
				v238 = v234<<(uint(int32(8))%32) + v233
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
				v304 = v238 + v239
				v305 = v72
				v306 = v70
			case 3:
				v226 = v72
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v304 = v227 + v68
				v305 = v226
				v306 = v70
			case 4:
				v223 = v72
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v226 = v223 + v224
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v304 = v227 + v68
				v305 = v226
				v306 = v70
			case 5:
				v218 = v72
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v223 = v219<<(uint(int32(8))%32) + v218
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v226 = v223 + v224
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v304 = v227 + v68
				v305 = v226
				v306 = v70
			case 6:
				v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
				v218 = v214<<(uint(int32(16))%32) + v72
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+5)))
				v223 = v219<<(uint(int32(8))%32) + v218
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+4)))
				v226 = v223 + v224
				v227 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v304 = v227 + v68
				v305 = v226
				v306 = v70
			case 7:
				v209 = v70
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v212 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v304 = v210 + v68
				v305 = v212 + v72
				v306 = v209
			case 8:
				v204 = v70
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v212 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v304 = v210 + v68
				v305 = v212 + v72
				v306 = v209
			case 9:
				v199 = v70
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
				v204 = v200<<(uint(int32(16))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v212 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v304 = v210 + v68
				v305 = v212 + v72
				v306 = v209
			case 10:
				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+10)))
				v199 = v195<<(uint(int32(24))%32) + v70
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+9)))
				v204 = v200<<(uint(int32(16))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+8)))
				v209 = v205<<(uint(int32(8))%32) + v204
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v212 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
				v304 = v210 + v68
				v305 = v212 + v72
				v306 = v209
			default:
				v304 = v68
				v305 = v72
				v306 = v70
			}
		}
		v309 = int32(14)
		v311 = v305 ^ v306 - base.I32_rotl(v305, v309)
		v315 = v311 ^ v304 - base.I32_rotl(v311, int32(11))
		v319 = v315 ^ v305 - base.I32_rotl(v315, int32(25))
		v323 = v319 ^ v311 - base.I32_rotl(v319, int32(16))
		v327 = v323 ^ v315 - base.I32_rotl(v323, int32(4))
		v331 = v327 ^ v319 - base.I32_rotl(v327, v309)
		v341 = base.I64_extend_i32_u(v331)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v331^v323-base.I32_rotl(v331, int32(24)))
	} else {
		v341 = v10
	}
	m.G0 = v8 + int32(16)
	return v341
}
func F_hashint4(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = int32(711645284)
	v10 = v2 - int32(1636608428) ^ v7 - int32(1455628627)
	v15 = v10 ^ int32(-1636608428) - base.I32_rotl(v10, int32(25))
	v20 = v15 ^ v7 - base.I32_rotl(v15, int32(16))
	v24 = v20 ^ v10 - base.I32_rotl(v20, int32(4))
	v28 = v24 ^ v15 - base.I32_rotl(v24, int32(14))
	return base.I64_extend_i32_u(v28 ^ v20 - base.I32_rotl(v28, int32(24)))
}
func F_hashint8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = int32(711645284)
	v18 = base.I32_wrap_i64(v3>>(uint(int64(63))%64)^int64(base.Ui64(v3)>>(uint(int64(32))%64))^v3) - int32(1636608428) ^ v15 - int32(1455628627)
	v23 = v18 ^ int32(-1636608428) - base.I32_rotl(v18, int32(25))
	v28 = v23 ^ v15 - base.I32_rotl(v23, int32(16))
	v32 = v28 ^ v18 - base.I32_rotl(v28, int32(4))
	v36 = v32 ^ v23 - base.I32_rotl(v32, int32(14))
	return base.I64_extend_i32_u(v36 ^ v28 - base.I32_rotl(v36, int32(24)))
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
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_hba_authname[0])))
	return v4
}
func F_heap2_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int64
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int64
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int64
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int64
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int64
	_ = v374
	var v395 int32
	_ = v395
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+96))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+48)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v23, v24, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v28 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L80
	}
L4:
	;
	m.G0 = v17 + int32(16)
	return
L5:
	;
	v32 = v21 & int32(112)
	v33 = int32(4)
	switch int32(base.Ui32(v32)>>(uint(v33)%32)) - v33 {
	case 0:
		goto L3
	case 1:
		goto L7
	default:
		goto L4
	case 3:
		goto L6
	}
L6:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v257 != 0 {
		goto L4
	} else {
		goto L57
	}
L7:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v38 = F_SnapBuildProcessChange(m, v22, v24, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v38 == int32(0) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v42 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v43 = int32(0)
	v45 = m.G0
	v47 = v45 - int32(16)
	m.G0 = v47
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[0]))
	if v50 == v43 {
		v93 = v43
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v93 != 0 {
		goto L4
	} else {
		goto L23
	}
L12:
	;
	m.G0 = v47 + int32(16)
	goto L11
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v54 = int32(0)
	v60 = F_XLogRecGetBlockTagExtended(m, v53, v54, v47+int32(4), v54, v54, v54)
	mBase = m.M
	if v60 == v54 {
		v93 = v43
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[0]))
	if v63 != v65 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[1]))
	if base.B2i32(v76 == int32(0))|base.B2i32(v63 != v76) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[2]))
	if v67 != v69 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[3]))
	if v71 == v73 {
		v93 = v43
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v93 = int32(1)
	goto L12
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[4]))
	if v82 != v84 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_decode[5]))
	if v86 == v88 {
		v93 = int32(0)
		goto L12
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v97 = m.G0
	v99 = v97 - int32(16)
	m.G0 = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+64))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v104&int32(8) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	m.G0 = v99 + int32(16)
	goto L4
L25:
	;
	v109 = int32(0)
	F_XLogRecGetBlockTag(m, v101, v109, v99, v109, v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+88))
	if v114 != v116 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v118 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+56)))
	v121 = F_filter_by_origin_cb_wrapper(m, l0, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v123 = int32(0)
	v125 = v99 + int32(12)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	if v128 < v123 {
		v150 = v123
		goto L34
	} else {
		goto L35
	}
L31:
	;
	if v121 != 0 {
		goto L24
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+2)))
	if v154 == int32(0) {
		goto L24
	} else {
		goto L44
	}
L34:
	;
	v153 = v150
	goto L33
L35:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+int32(0))+76)))
	if v133 != int32(1) {
		v150 = v123
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v137 = v127 + int32(76)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+43)))
	if v138 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if v125 == int32(0) {
		v150 = v123
		goto L34
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if v125 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v143
	v153 = v143
	goto L33
L41:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v146
	goto L43
L42:
	;
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v137)+44))
	v150 = v148
	goto L34
L44:
	;
	v161 = v153
	v167 = int32(0)
	goto L45
L45:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v172 = F_ReorderBufferAllocChange(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	goto L24
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+8)) = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v172)+16)) = uint16(v177)
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v172)+20)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v99)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v172)+28)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v187 = (v161 + int32(1)) & int32(-2)
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187))))
	v189 = F_ReorderBufferAllocTupleBuf(m, v183, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+40)) = v189
	v192 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = v192
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+8)) = uint16(v192)
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v188 + int32(23)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v202 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v201)+15)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v201)+8)) = v202
	*(*int64)(unsafe.Add(mBase, uint32(v201))) = v202
	v209 = v187 + int32(7)
	if v188 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	base.MemoryCopy(m, v210+int32(23), v209, v188)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v201)+20)) = uint16(v214)
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v201)+18)) = uint16(v216)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+22)) = uint8(v218)
	v221 = v167 + int32(1)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v222&int32(2) != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+2)))
	v228 = base.B2i32(v221 == v225)
	goto L54
L53:
	;
	v228 = int32(0)
	goto L54
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v172)+32)) = uint8(v228)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v101)+96))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+36))
	v233 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferQueueChange(m, v230, v232, v233, v172, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+2)))
	if base.Ui32(v221) < base.Ui32(v238) {
		v161 = v188 + v209
		v167 = v221
		goto L45
	} else {
		goto L56
	}
L56:
	;
	goto L46
L57:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+96))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+64))
	v262 = m.G0
	v264 = v262 - int32(32)
	m.G0 = v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	F_ReorderBufferXidSetCatalogChanges(m, v266, v24, v258)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v261)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+24)) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v261)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+16)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v261)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+8)) = v278
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v264)+12)) = uint16(v280)
	v283 = v264 + int32(16)
	v285 = v264 + int32(8)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v273)+124))
	v288 = F_MemoryContextAlloc(m, v286, int32(64))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v290 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v288)+56)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+48)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+40)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+32)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+24)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+16)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288)+8)) = v290
	*(*int64)(unsafe.Add(mBase, uint32(v288))) = v290
	v307 = F_ReorderBufferTXNByXid(m, v273, v272, int32(0), v258)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+28)) = v309
	v311 = *(*int64)(unsafe.Add(mBase, uint32(v283)))
	*(*int64)(unsafe.Add(mBase, uint32(v288)+20)) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+32)) = v313
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v285)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v288)+36)) = uint16(v315)
	*(*int32)(unsafe.Add(mBase, uint32(v288)+48)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v288)+44)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v288)+40)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v288)+12)) = v307
	*(*int64)(unsafe.Add(mBase, uint32(v288))) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v288)+8)) = int32(7)
	v325 = v307 + int32(136)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v307)+140))
	if v326 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307)+140)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v307)+136)) = v325
	goto L63
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v288)+56)) = v325
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v307)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v288)+52)) = v332
	v335 = v288 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v332)+4)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v307)+136)) = v335
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v307)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v307)+144)) = v338 + int64(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v261)+4))
	if v343 != int32(-1) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)+124))
	v372 = F_MemoryContextAlloc(m, v370, int32(64))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L78
	}
L65:
	;
	if base.Ui32(v342) < base.Ui32(v343) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v342 != int32(-1) {
		v366 = v342
		goto L64
	} else {
		goto L74
	}
L68:
	;
	v347 = v343
	goto L70
L69:
	;
	v347 = v342
	goto L70
L70:
	;
	if v342 == int32(-1) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v350 = v343
	goto L73
L72:
	;
	v350 = v347
	goto L73
L73:
	;
	v366 = v350
	goto L64
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_decode_0), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_heap2_decode_1), int32(719), int32(_a_F_heap2_decode_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v374 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v372)+16)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+8)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+56)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+48)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+40)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+32)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372)+24)) = v374
	*(*int64)(unsafe.Add(mBase, uint32(v372))) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v372)+20)) = v366 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = int32(6)
	F_ReorderBufferQueueChange(m, v369, v24, v258, v372, int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	m.G0 = v264 + int32(32)
	goto L4
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v32
	F_errmsg_internal(m, int32(_a_F_heap2_decode_3), v17)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_heap2_decode_4), int32(461), int32(_a_F_heap2_decode_5))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_heap2_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if base.Ui32(l0) <= base.Ui32(int32(223)) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_heap2_identify[0])))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10
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
	var v28 int32
	_ = v28
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v562 int32
	_ = v562
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v606 int32
	_ = v606
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v711 int32
	_ = v711
	var v732 int32
	_ = v732
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v774 int32
	_ = v774
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v838 int32
	_ = v838
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v957 int32
	_ = v957
	var v962 int32
	_ = v962
	var v971 int32
	_ = v971
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1070 int32
	_ = v1070
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int64
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1118 int64
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int64
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1141 int64
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int64
	_ = v1317
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1340 int32
	_ = v1340
	var v1343 int32
	_ = v1343
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1407 int32
	_ = v1407
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int64
	_ = v1416
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1471 int32
	_ = v1471
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1528 int32
	_ = v1528
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1560 int32
	_ = v1560
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1574 int32
	_ = v1574
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1597 int32
	_ = v1597
	var v1602 int32
	_ = v1602
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1653 int64
	_ = v1653
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1689 int32
	_ = v1689
	var v1691 int32
	_ = v1691
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1723 int32
	_ = v1723
	var v1728 int32
	_ = v1728
	var v1742 int32
	_ = v1742
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1753 int64
	_ = v1753
	var v1755 int32
	_ = v1755
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1768 int32
	_ = v1768
	var v1772 int32
	_ = v1772
	var v1773 int64
	_ = v1773
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1785 int32
	_ = v1785
	var v1786 int64
	_ = v1786
	var v1788 int32
	_ = v1788
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1833 int32
	_ = v1833
	var v1840 int32
	_ = v1840
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1880 int32
	_ = v1880
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1892 int32
	_ = v1892
	var v1894 int32
	_ = v1894
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1915 int32
	_ = v1915
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1927 int64
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1930 int64
	_ = v1930
	var v1938 int64
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1946 int32
	_ = v1946
	var v1947 int32
	_ = v1947
	var v1949 int32
	_ = v1949
	var v1950 int32
	_ = v1950
	var v1954 int32
	_ = v1954
	var v1957 int64
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1959 int32
	_ = v1959
	var v1960 int32
	_ = v1960
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1964 int32
	_ = v1964
	var v1965 int32
	_ = v1965
	var v1970 int32
	_ = v1970
	var v1976 int32
	_ = v1976
	var v1977 int64
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2011 int32
	_ = v2011
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2052 int32
	_ = v2052
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2078 int32
	_ = v2078
	var v2083 int32
	_ = v2083
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2090 int64
	_ = v2090
	var v2099 int32
	_ = v2099
	var v2104 int32
	_ = v2104
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2118 int32
	_ = v2118
	var v2123 int32
	_ = v2123
	var v2170 int32
	_ = v2170
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2203 int32
	_ = v2203
	var v2207 int32
	_ = v2207
	var v2212 int32
	_ = v2212
	var v2216 int32
	_ = v2216
	var v2220 int32
	_ = v2220
	var v2225 int32
	_ = v2225
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2252 int32
	_ = v2252
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(_a_F_heap2_redo_0)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+48)))
	v28 = v26 & int32(240)
	switch int32(base.Ui32(v28)>>(uint(int32(4))%32))&int32(7) - int32(1) {
	case 0, 1, 2:
		goto L11
	case 3:
		goto L8
	case 4:
		goto L10
	case 5:
		goto L9
	case 6:
		goto L6
	default:
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L12
	} else {
		goto L451
	}
L2:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L12
	} else {
		goto L448
	}
L3:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L12
	} else {
		goto L445
	}
L4:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L12
	} else {
		goto L442
	}
L5:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L12
	} else {
		goto L439
	}
L6:
	;
	m.G0 = v23 + int32(_a_F_heap2_redo_0)
	return
L7:
	;
	v1921 = m.G0
	v1923 = v1921 - int32(1136)
	m.G0 = v1923
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+64))
	v1927 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+4))
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1926)))
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1925)+36))
	v1930 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v1923)+96)) = uint32(v1930)
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+104)) = v1929
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+100)) = v1928
	*(*int64)(unsafe.Add(mBase, uint32(v1923)+84)) = v1927
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+80)) = int32(_a_F_heap2_redo_1)
	v1938 = int64(base.Ui64(v1930) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1923)+92)) = uint32(v1938)
	v1941 = v1923 + int32(112)
	v1946 = F_pg_snprintf(m, v1941, int32(1024), int32(_a_F_heap2_redo_2), v1923+int32(80))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L12
	} else {
		goto L388
	}
L8:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1909 = m.ExcPending
	if v1909 != 0 {
		goto L12
	} else {
		goto L385
	}
L9:
	;
	v1773 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774)+7)))
	if v1775&int32(1) != 0 {
		goto L357
	} else {
		goto L358
	}
L10:
	;
	v1125 = base.I32_extend8_s(v26)
	v1126 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1127 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+164)) = v1127
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	F_XLogRecGetBlockTag(m, l0, v1127, v23+int32(_a_F_heap2_redo_3), v1127, v23+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v1137 = m.ExcPending
	if v1137 != 0 {
		goto L12
	} else {
		goto L207
	}
L11:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0]))) = v37
	F_XLogRecGetBlockTag(m, l0, v37, v23+int32(168), v37, v23+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if v47&int32(8) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v70 = v47 & int32(256)
	v71 = int32(0)
	v74 = v47 & int32(4)
	v79 = F_XLogReadBufferForRedoExtended(m, l0, v71, v71, int32(base.Ui32(v74)>>(uint(int32(2))%32)), v23+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L12
	} else {
		goto L19
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[1]))
	if base.Ui32(v53) < base.Ui32(int32(2)) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)+2))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v59
	F_ResolveRecoveryConflictWithSnapshot(m, v56, int32(base.Ui32(v47&int32(2))>>(uint(int32(1))%32)), v23+int32(32))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	if v907 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L19:
	;
	if v79 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	if v81 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v100 = int32(0)
	v102 = v23 + int32(136)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+72))
	if v105 < v100 {
		v127 = v100
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85+(v81^int32(-1))<<(uint(int32(2))%32))))
	v99 = v91
	goto L21
L23:
	;
	goto L24
L24:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v99 = v93 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	v132 = v23 + int32(140)
	v134 = v23 + int32(132)
	v136 = v23 + int32(128)
	v138 = v23 + int32(152)
	v140 = v23 + int32(164)
	v142 = v23 + int32(148)
	v144 = v23 + int32(160)
	v146 = v23 + int32(144)
	v148 = v23 + int32(156)
	if v47&int32(16) != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v130 = v127
	goto L25
L27:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+int32(0))+76)))
	if v110 != int32(1) {
		v127 = v100
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v114 = v104 + int32(76)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+43)))
	if v115 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v102 == int32(0) {
		v127 = v100
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v102 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v120 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v120
	v130 = v120
	goto L25
L33:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v123
	goto L35
L34:
	;
	goto L35
L35:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v114)+44))
	v127 = v125
	goto L26
L36:
	;
	if v47&int32(32) != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130))))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v151
	v154 = v130 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v164 = v154 + v156*int32(12)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v160 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v134))) = v160
	v164 = v130
	goto L36
L40:
	;
	if v47&int32(64) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164))))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v167
	v169 = int32(2)
	v170 = v164 + v169
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v181 = v170 + v172<<(uint(v169)%32)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v176 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v176
	v181 = v164
	goto L40
L44:
	;
	if v47&int32(128) != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181))))
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v184
	v187 = v181 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v198 = v187 + v189<<(uint(int32(1))%32)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v193 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v193
	v198 = v181
	goto L44
L48:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v23)+152))
	v218 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v23)+148))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v23)+144))
	v227 = base.B2i32(v218 < v217) | base.B2i32(v218 < v220) | base.B2i32(v218 < v224)
	if v227 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v201
	v204 = v198 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v204 + v206<<(uint(int32(1))%32)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v211
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v198
	goto L48
L52:
	;
	v228 = int32(0)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v23)+160))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	if v233 < v228 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	goto L54
L54:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v23)+140))
	if int32(0) < v732 {
		goto L114
	} else {
		goto L115
	}
L55:
	;
	goto L54
L56:
	;
	if v217 <= int32(0) {
		goto L60
	} else {
		goto L61
	}
L57:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v237+(v233^int32(-1))<<(uint(int32(2))%32))))
	v251 = v243
	goto L56
L58:
	;
	goto L59
L59:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v251 = v245 + v233<<(uint(int32(13))%32) + int32(-8192)
	goto L56
L60:
	;
	if v220 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v255 = v251 + int32(20)
	if v217 != int32(1) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v264 = v230
	v266 = int32(0)
	goto L65
L63:
	;
	v311 = v230
	goto L64
L64:
	;
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311))))
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v311)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v330<<(uint(int32(2))%32)))) = v334&int32(_a_F_heap2_redo_5) | int32(_a_F_heap2_redo_6)
	goto L60
L65:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264))))
	v284 = int32(2)
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+2)))
	v288 = int32(_a_F_heap2_redo_5)
	v290 = int32(_a_F_heap2_redo_6)
	*(*int32)(unsafe.Add(mBase, uint32(v255+v283<<(uint(v284)%32)))) = v287&v288 | v290
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+4)))
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v255+v293<<(uint(v284)%32)))) = v297&v288 | v290
	v304 = v264 + int32(8)
	v306 = v266 + v284
	if v306 != v217&int32(2147483646) {
		v264 = v304
		v266 = v306
		goto L65
	} else {
		goto L67
	}
L66:
	;
	if v217&int32(1) == int32(0) {
		goto L60
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v311 = v304
	goto L64
L69:
	;
	if v224 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L70:
	;
	v363 = v220 & int32(3)
	v365 = v251 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v220) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v374 = int32(0)
	v375 = v231
	goto L74
L72:
	;
	v426 = v231
	goto L73
L73:
	;
	v446 = int32(0)
	v447 = v426
	goto L78
L74:
	;
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375))))
	v392 = int32(2)
	v395 = int32(_a_F_heap2_redo_7)
	*(*int32)(unsafe.Add(mBase, uint32(v365+v391<<(uint(v392)%32)))) = v395
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v365+v397<<(uint(v392)%32)))) = v395
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v365+v403<<(uint(v392)%32)))) = v395
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v365+v409<<(uint(v392)%32)))) = v395
	v416 = v375 + int32(8)
	v418 = v374 + int32(4)
	if v418 != v220&int32(2147483644) {
		v374 = v418
		v375 = v416
		goto L74
	} else {
		goto L76
	}
L75:
	;
	if v363 == int32(0) {
		goto L69
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	v426 = v416
	goto L73
L78:
	;
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v447))))
	v464 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v365+v463<<(uint(v464)%32)))) = int32(_a_F_heap2_redo_7)
	v472 = v446 + int32(1)
	if v472 != v363 {
		v446 = v472
		v447 = v447 + v464
		goto L78
	} else {
		goto L80
	}
L79:
	;
	goto L69
L80:
	;
	goto L79
L81:
	;
	if v74 == v228 {
		goto L93
	} else {
		goto L94
	}
L82:
	;
	v497 = v224 & int32(3)
	v499 = v251 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v224) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v506 = int32(0)
	v511 = v232
	goto L86
L84:
	;
	v562 = v232
	goto L85
L85:
	;
	v578 = int32(0)
	v583 = v562
	goto L90
L86:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511))))
	v526 = int32(2)
	v529 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v499+v525<<(uint(v526)%32)))) = v529
	v531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v499+v531<<(uint(v526)%32)))) = v529
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v499+v537<<(uint(v526)%32)))) = v529
	v543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v511)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v499+v543<<(uint(v526)%32)))) = v529
	v550 = v511 + int32(8)
	v552 = v506 + int32(4)
	if v552 != v224&int32(2147483644) {
		v506 = v552
		v511 = v550
		goto L86
	} else {
		goto L88
	}
L87:
	;
	if v497 == int32(0) {
		goto L81
	} else {
		goto L89
	}
L88:
	;
	goto L87
L89:
	;
	v562 = v550
	goto L85
L90:
	;
	v597 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v583))))
	v598 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v499+v597<<(uint(v598)%32)))) = int32(0)
	v606 = v578 + int32(1)
	if v606 != v497 {
		v578 = v606
		v583 = v583 + v598
		goto L90
	} else {
		goto L92
	}
L91:
	;
	goto L81
L92:
	;
	goto L91
L93:
	;
	v628 = int32(0)
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+12)))
	if base.Ui32(v634) < base.Ui32(int32(25)) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	F_PageRepairFragmentation(m, v251)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L12
	} else {
		goto L113
	}
L96:
	;
	goto L55
L97:
	;
	v699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+10)))
	v701 = v699 & int32(_a_F_heap2_redo_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+10)) = uint16(v701)
	goto L96
L98:
	;
	v642 = int32(base.Ui32(v634+int32(_a_F_heap2_redo_9))>>(uint(int32(2))%32)) & int32(_a_F_heap2_redo_10)
	if v642 == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v648 = v642
	v650 = v628
	v653 = v628
	goto L101
L100:
	;
	if int32(0) < v677 {
		goto L109
	} else {
		goto L110
	}
L101:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v251+int32(20)+v648<<(uint(int32(2))%32))))
	v659 = v657 & int32(_a_F_heap2_redo_7)
	if base.B2i32(v648 == int32(1))|v653 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v677 = v671
	v679 = int32(0)
	goto L100
L103:
	;
	v674 = v648 - int32(1)
	if v674 != 0 {
		v648 = v674
		v650 = v671
		v653 = v672
		goto L101
	} else {
		goto L108
	}
L104:
	;
	v665 = int32(0)
	v671 = v650 + base.B2i32(v659 == v665)
	v672 = base.B2i32(v659 != v665)
	goto L103
L105:
	;
	goto L106
L106:
	;
	if v659 != 0 {
		v671 = v650
		v672 = v653
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v677 = v650
	v679 = int32(1)
	goto L100
L108:
	;
	goto L102
L109:
	;
	v684 = v634 - v677<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+12)) = uint16(v684)
	goto L111
L110:
	;
	goto L111
L111:
	;
	if v679 == int32(0) {
		goto L97
	} else {
		goto L112
	}
L112:
	;
	v688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+10)))
	v690 = v688 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v251)+10)) = uint16(v690)
	goto L96
L113:
	;
	goto L55
L114:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v739 = v737
	v742 = v732
	v754 = v2
	goto L117
L115:
	;
	goto L116
L116:
	;
	if v70 != 0 {
		goto L132
	} else {
		goto L133
	}
L117:
	;
	v759 = v754 * int32(12)
	v760 = v739 + v759
	v761 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v760)+10)))
	if v761 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v762 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v760)+6)))
	v763 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v760)+4)))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v760)))
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760)+8)))
	v774 = int32(0)
	goto L122
L120:
	;
	v818 = v739
	v821 = v742
	goto L121
L121:
	;
	v838 = v754 + int32(1)
	if v838 < v821 {
		v739 = v818
		v742 = v821
		v754 = v838
		goto L117
	} else {
		goto L131
	}
L122:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	v792 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v791 + v792
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v791))))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v99+int32(20)+v795<<(uint(v792)%32))))
	v802 = v99 + v799&int32(_a_F_heap2_redo_5)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+4)) = v764
	if v765&int32(2) != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v23)+140))
	v818 = v812
	v821 = v816
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+8)) = int32(2)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v765&int32(4) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+8)) = int32(0)
	goto L129
L128:
	;
	goto L129
L129:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+18)) = uint16(v763)
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+20)) = uint16(v762)
	v811 = v774 + int32(1)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v23)+132))
	v814 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v812+v759)+10)))
	if base.Ui32(v811) < base.Ui32(v814) {
		v774 = v811
		goto L122
	} else {
		goto L130
	}
L130:
	;
	goto L123
L131:
	;
	goto L118
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = int32(0)
	v862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+10)))
	v864 = v862 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+10)) = uint16(v864)
	goto L134
L133:
	;
	goto L134
L134:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	F_MarkBufferDirty(m, v866)
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L12
	} else {
		goto L135
	}
L135:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v23)+140))
	if v227|base.B2i32(int32(0) < v869) != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v99))) = base.I64_rotl(v35, int64(32))
	goto L18
L137:
	;
	if v70 == int32(0) {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	v876 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[5])))
	if v876 != 0 {
		goto L136
	} else {
		goto L139
	}
L139:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[6]))
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v878)+268))
	goto L140
L140:
	;
	if base.B2i32(v879 != int32(0)) == int32(0) {
		goto L18
	} else {
		goto L141
	}
L141:
	;
	goto L136
L142:
	;
	if v70 == int32(0) {
		goto L174
	} else {
		goto L175
	}
L143:
	;
	v910 = int32(0)
	v1010 = v910
	v1011 = v910
	goto L142
L144:
	;
	goto L145
L145:
	;
	v913 = v47 & int32(480)
	if v913 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	F_UnlockReleaseBuffer(m, v1003)
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L12
	} else {
		goto L173
	}
L147:
	;
	v1003 = v907
	v1004 = int32(0)
	goto L146
L148:
	;
	goto L149
L149:
	;
	if v907 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v938 = int32(4)
	v939 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+14)))
	v940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+12)))
	v941 = v939 - v940
	if v941 <= v938 {
		goto L155
	} else {
		goto L156
	}
L151:
	;
	v920 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v920+(v907^int32(-1))<<(uint(int32(2))%32))))
	v934 = v926
	goto L150
L152:
	;
	goto L153
L153:
	;
	v928 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v934 = v928 + v907<<(uint(int32(13))%32) + int32(-8192)
	goto L150
L154:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	v1003 = v1002
	v1004 = v1001
	goto L146
L155:
	;
	v944 = v938
	goto L157
L156:
	;
	v944 = v941
	goto L157
L157:
	;
	v946 = v944 - int32(4)
	if v946 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1001 = int32(0)
	goto L154
L159:
	;
	goto L160
L160:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v940) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v1001 = v946
	goto L154
L162:
	;
	v957 = int32(base.Ui32(v940+int32(_a_F_heap2_redo_9)) >> (uint(int32(2)) % 32))
	goto L164
L163:
	;
	v957 = int32(0)
	goto L164
L164:
	;
	if base.Ui32(v957&int32(_a_F_heap2_redo_10)) < base.Ui32(int32(291)) {
		goto L161
	} else {
		goto L165
	}
L165:
	;
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+10)))
	if v962&int32(1) == int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1001 = int32(0)
	goto L154
L167:
	;
	goto L168
L168:
	;
	v971 = int32(1)
	goto L169
L169:
	;
	v980 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934+int32(20)+v971&int32(_a_F_heap2_redo_10)<<(uint(int32(2))%32))+1)))
	if v980&int32(384) == int32(0) {
		goto L161
	} else {
		goto L171
	}
L170:
	;
	v1001 = int32(0)
	goto L154
L171:
	;
	v986 = v971 + int32(1)
	v987 = int32(_a_F_heap2_redo_10)
	if base.Ui32(v986&v987) <= base.Ui32(v957&v987) {
		v971 = v986
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v1010 = base.B2i32(v913 != int32(0))
	v1011 = v1004
	goto L142
L174:
	;
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	if v1113 != 0 {
		goto L201
	} else {
		goto L202
	}
L175:
	;
	v1019 = F_XLogReadBufferForRedoExtended(m, l0, int32(1), int32(3), int32(0), v23+int32(_a_F_heap2_redo_11))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L12
	} else {
		goto L176
	}
L176:
	;
	if v1019 != 0 {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	if v1021 < int32(0) {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v1040 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1039)+14)))
	if v1040 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1025+(v1021^int32(-1))<<(uint(int32(2))%32))))
	v1039 = v1031
	goto L178
L180:
	;
	goto L181
L181:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1039 = v1033 + v1021<<(uint(int32(13))%32) + int32(-8192)
	goto L178
L182:
	;
	v1043 = int32(_a_F_heap2_redo_12)
	v1044 = int32(0)
	if v1044|(v1039&int32(3)|int32(1)) == v1044 {
		goto L187
	} else {
		goto L188
	}
L183:
	;
	v1094 = v1021
	goto L184
L184:
	;
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v1097
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	if v47&int32(512) != 0 {
		goto L196
	} else {
		goto L197
	}
L185:
	;
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	v1094 = v1093
	goto L184
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1039)+10)) = int32(_a_F_heap2_redo_13)
	v1084 = int32(_a_F_heap2_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1039)+18)) = uint16(v1084)
	v1090 = int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1039)+16)) = uint16(v1090)
	*(*uint16)(unsafe.Add(mBase, uint32(v1039)+14)) = uint16(v1090)
	goto L185
L187:
	;
	goto L190
L188:
	;
	goto L189
L189:
	;
	goto L195
L190:
	;
	v1061 = v1039 + v1043
	v1063 = v1039 + int32(4)
	if base.Ui32(v1063) < base.Ui32(v1061) {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v1065 = v1061
	goto L193
L192:
	;
	v1065 = v1063
	goto L193
L193:
	;
	v1070 = (v1039^int32(-1)+v1065)&int32(-4) + int32(4)
	if v1070 == int32(0) {
		goto L186
	} else {
		goto L194
	}
L194:
	;
	base.MemoryFill(m, v1039, int32(0), v1070)
	goto L186
L195:
	;
	base.MemoryFill(m, v1039, int32(0), v1043)
	goto L186
L196:
	;
	v1104 = int32(3)
	goto L198
L197:
	;
	v1104 = int32(1)
	goto L198
L198:
	;
	v1105 = F_visibilitymap_set(m, v1099, v1094, v1104)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L12
	} else {
		goto L199
	}
L199:
	;
	if v1105 == v1104 {
		goto L174
	} else {
		goto L200
	}
L200:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1039))) = base.I64_rotl(v35, int64(32))
	goto L174
L201:
	;
	F_UnlockReleaseBuffer(m, v1113)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L12
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if v1010 == int32(0) {
		goto L6
	} else {
		goto L205
	}
L204:
	;
	goto L203
L205:
	;
	v1118 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v1118
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v1120
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	F_XLogRecordPageWithFreeSpace(m, v23, v1122, v1011)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L12
	} else {
		goto L206
	}
L206:
	;
	goto L6
L207:
	;
	v1138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	if v1138&int32(1) != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v1141 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v1141
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+88)) = v1143
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	F_heap_xlog_vm_clear(m, l0, v23+int32(80), v1147, int32(3))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L12
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	if v1125 < int32(0) {
		goto L214
	} else {
		goto L215
	}
L211:
	;
	goto L210
L212:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	if v1544 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L213:
	;
	v1233 = int32(0)
	v1235 = v23 + int32(160)
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+72))
	if v1238 < v1233 {
		v1260 = v1233
		goto L236
	} else {
		goto L237
	}
L214:
	;
	v1154 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L12
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v1225 = int32(0)
	v1229 = F_XLogReadBufferForRedo(m, l0, v1225, v23+int32(_a_F_heap2_redo_11))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L12
	} else {
		goto L233
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0]))) = v1154
	if v1154 < int32(0) {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1175 = int32(_a_F_heap2_redo_12)
	v1176 = int32(0)
	if v1176|(v1174&int32(3)|int32(1)) == v1176 {
		goto L224
	} else {
		goto L225
	}
L219:
	;
	v1160 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1160+(v1154^int32(-1))<<(uint(int32(2))%32))))
	v1174 = v1166
	goto L218
L220:
	;
	goto L221
L221:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1174 = v1168 + v1154<<(uint(int32(13))%32) + int32(-8192)
	goto L218
L222:
	;
	goto L213
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1174)+10)) = int32(_a_F_heap2_redo_13)
	v1216 = int32(_a_F_heap2_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1174)+18)) = uint16(v1216)
	v1222 = int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1174)+16)) = uint16(v1222)
	*(*uint16)(unsafe.Add(mBase, uint32(v1174)+14)) = uint16(v1222)
	goto L222
L224:
	;
	goto L227
L225:
	;
	goto L226
L226:
	;
	goto L232
L227:
	;
	v1193 = v1174 + v1175
	v1195 = v1174 + int32(4)
	if base.Ui32(v1195) < base.Ui32(v1193) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1197 = v1193
	goto L230
L229:
	;
	v1197 = v1195
	goto L230
L230:
	;
	v1202 = (v1174^int32(-1)+v1197)&int32(-4) + int32(4)
	if v1202 == int32(0) {
		goto L223
	} else {
		goto L231
	}
L231:
	;
	base.MemoryFill(m, v1174, int32(0), v1202)
	goto L223
L232:
	;
	base.MemoryFill(m, v1174, int32(0), v1175)
	goto L223
L233:
	;
	if v1229 != 0 {
		v1528 = v1225
		goto L212
	} else {
		goto L234
	}
L234:
	;
	goto L213
L235:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	if v1264 < int32(0) {
		goto L247
	} else {
		goto L248
	}
L236:
	;
	v1263 = v1260
	goto L235
L237:
	;
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1237+int32(0))+76)))
	if v1243 != int32(1) {
		v1260 = v1233
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1247 = v1237 + int32(76)
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247)+43)))
	if v1248 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	if v1235 == int32(0) {
		v1260 = v1233
		goto L236
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	if v1235 != 0 {
		goto L243
	} else {
		goto L244
	}
L242:
	;
	v1253 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1235))) = v1253
	v1263 = v1253
	goto L235
L243:
	;
	v1256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1247)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1235))) = v1256
	goto L245
L244:
	;
	goto L245
L245:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1247)+44))
	v1260 = v1258
	goto L236
L246:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v23)+160))
	v1284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1129)+2)))
	if v1284 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L247:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1268+(v1264^int32(-1))<<(uint(int32(2))%32))))
	v1282 = v1274
	goto L246
L248:
	;
	goto L249
L249:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1282 = v1276 + v1264<<(uint(int32(13))%32) + int32(-8192)
	goto L246
L250:
	;
	if v1460 != v1263+v1283 {
		goto L3
	} else {
		goto L282
	}
L251:
	;
	v1460 = v1263
	v1471 = v2
	goto L250
L252:
	;
	goto L253
L253:
	;
	v1288 = v1129 + int32(4)
	if int32(0) <= v1125 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1288))))
	v1293 = v1292
	goto L256
L255:
	;
	v1293 = int32(1)
	goto L256
L256:
	;
	v1294 = int32(1)
	v1295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+12)))
	if base.Ui32(v1295) < base.Ui32(int32(25)) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1304 = v1294
	goto L259
L258:
	;
	v1304 = int32(base.Ui32(v1295+int32(_a_F_heap2_redo_9))>>(uint(int32(2))%32)) + v1294
	goto L259
L259:
	;
	v1305 = int32(_a_F_heap2_redo_10)
	if base.Ui32(v1304&v1305) < base.Ui32(v1293&v1305) {
		goto L5
	} else {
		goto L260
	}
L260:
	;
	v1311 = v23 + int32(191)
	v1315 = (v1263 + int32(1)) & int32(-2)
	v1316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1315))))
	v1317 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+183)) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v23)+176)) = v1317
	*(*int64)(unsafe.Add(mBase, uint32(v23)+168)) = v1317
	v1324 = v1315 + int32(7)
	if v1316 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	base.MemoryCopy(m, v1311, v1324, v1316)
	goto L263
L262:
	;
	goto L263
L263:
	;
	v1326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1315)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+186)) = uint16(v1326)
	v1328 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1315)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1328)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+190)) = uint8(v1330)
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1333 = *(*int32)(unsafe.Add(mBase, uint32(v1332)+36))
	v1335 = v1328 & int32(_a_F_heap2_redo_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1335)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v1333
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+182)) = uint16(v1340)
	v1343 = int32(base.Ui32(v1340) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+180)) = uint16(v1343)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+184)) = uint16(v1293)
	v1353 = F_PageAddItemExtended(m, v1282, v23+int32(168), v1316+int32(23), v1293&int32(_a_F_heap2_redo_10), int32(3))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L12
	} else {
		goto L264
	}
L264:
	;
	if v1353 == int32(0) {
		goto L4
	} else {
		goto L265
	}
L265:
	;
	v1357 = v1316 + v1324
	v1358 = int32(768)
	v1361 = base.B2i32(v1328&v1358 == v1358)
	v1362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1129)+2)))
	if base.Ui32(v1362) < base.Ui32(int32(2)) {
		v1460 = v1357
		v1471 = v1361
		goto L250
	} else {
		goto L266
	}
L266:
	;
	v1369 = v1357
	v1372 = int32(1)
	goto L267
L267:
	;
	if int32(0) <= v1125 {
		goto L269
	} else {
		goto L270
	}
L268:
	;
	v1460 = v1454
	v1471 = v1361
	goto L250
L269:
	;
	v1391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1288+v1372<<(uint(int32(1))%32)))))
	v1394 = v1391
	goto L271
L270:
	;
	v1394 = v1372 + int32(1)
	goto L271
L271:
	;
	v1396 = v1394 & int32(_a_F_heap2_redo_10)
	v1397 = int32(1)
	v1398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+12)))
	if base.Ui32(v1398) < base.Ui32(int32(25)) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1407 = v1397
	goto L274
L273:
	;
	v1407 = int32(base.Ui32(v1398+int32(_a_F_heap2_redo_9))>>(uint(int32(2))%32)) + v1397
	goto L274
L274:
	;
	if base.Ui32(v1407&int32(_a_F_heap2_redo_10)) < base.Ui32(v1396) {
		goto L5
	} else {
		goto L275
	}
L275:
	;
	v1414 = (v1369 + int32(1)) & int32(-2)
	v1415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414))))
	v1416 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+183)) = v1416
	*(*int64)(unsafe.Add(mBase, uint32(v23)+176)) = v1416
	*(*int64)(unsafe.Add(mBase, uint32(v23)+168)) = v1416
	v1423 = v1414 + int32(7)
	if v1415 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	base.MemoryCopy(m, v1311, v1423, v1415)
	goto L278
L277:
	;
	goto L278
L278:
	;
	v1425 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+186)) = uint16(v1425)
	v1427 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1414)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1427)
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1414)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+190)) = uint8(v1429)
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1432 = *(*int32)(unsafe.Add(mBase, uint32(v1431)+36))
	v1434 = v1427 & int32(_a_F_heap2_redo_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+188)) = uint16(v1434)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+168)) = v1432
	v1439 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+182)) = uint16(v1439)
	v1442 = int32(base.Ui32(v1439) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+180)) = uint16(v1442)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+184)) = uint16(v1394)
	v1450 = F_PageAddItemExtended(m, v1282, v23+int32(168), v1415+int32(23), v1396, int32(3))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L12
	} else {
		goto L279
	}
L279:
	;
	if v1450 == int32(0) {
		goto L4
	} else {
		goto L280
	}
L280:
	;
	v1454 = v1423 + v1415
	v1456 = v1372 + int32(1)
	v1457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1129)+2)))
	if base.Ui32(v1456) < base.Ui32(v1457) {
		v1369 = v1454
		v1372 = v1456
		goto L267
	} else {
		goto L281
	}
L281:
	;
	goto L268
L282:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1282))) = base.I64_rotl(v1126, int64(32))
	v1484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	if v1484&int32(1) != 0 {
		goto L284
	} else {
		goto L285
	}
L283:
	;
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	F_MarkBufferDirty(m, v1520)
	mBase = m.M
	v1522 = m.ExcPending
	if v1522 != 0 {
		goto L12
	} else {
		goto L298
	}
L284:
	;
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+10)))
	v1489 = v1487 & int32(_a_F_heap2_redo_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v1282)+10)) = uint16(v1489)
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	v1492 = v1491
	goto L286
L285:
	;
	v1492 = v1484
	goto L286
L286:
	;
	if v1492&int32(32) != 0 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1282)+20)) = int32(0)
	v1497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1282)+10)))
	v1499 = v1497 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1282)+10)) = uint16(v1499)
	goto L283
L288:
	;
	goto L289
L289:
	;
	if v1471 != 0 {
		goto L283
	} else {
		goto L290
	}
L290:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1501)+36))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1282)+20))
	if v1503 == int32(0) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1282)+20)) = v1502
	goto L283
L292:
	;
	v1506 = int32(3)
	if base.B2i32(base.Ui32(v1503) < base.Ui32(v1506))|base.B2i32(base.Ui32(v1502) < base.Ui32(v1506)) == int32(0) {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	if v1502-v1503 < int32(0) {
		goto L291
	} else {
		goto L296
	}
L294:
	;
	goto L295
L295:
	;
	if base.Ui32(v1503) <= base.Ui32(v1502) {
		goto L283
	} else {
		goto L297
	}
L296:
	;
	goto L283
L297:
	;
	goto L291
L298:
	;
	v1528 = int32(1)
	goto L212
L299:
	;
	v1665 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0]))) = v1665
	v1667 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	if v1667&int32(32) == v1665 {
		goto L331
	} else {
		goto L332
	}
L300:
	;
	v1547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	v1550 = int32(0)
	if base.B2i32(v1547&int32(32) == v1550)&(v1528^int32(-1)) == v1550 {
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if v1544 < int32(0) {
		goto L305
	} else {
		goto L306
	}
L302:
	;
	goto L303
L303:
	;
	F_UnlockReleaseBuffer(m, v1544)
	mBase = m.M
	v1663 = m.ExcPending
	if v1663 != 0 {
		goto L12
	} else {
		goto L330
	}
L304:
	;
	v1578 = int32(4)
	v1579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1574)+14)))
	v1580 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1574)+12)))
	v1581 = v1579 - v1580
	if v1581 <= v1578 {
		goto L309
	} else {
		goto L310
	}
L305:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1560+(v1544^int32(-1))<<(uint(int32(2))%32))))
	v1574 = v1566
	goto L304
L306:
	;
	goto L307
L307:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1574 = v1568 + v1544<<(uint(int32(13))%32) + int32(-8192)
	goto L304
L308:
	;
	v1642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129))))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[0])))
	F_UnlockReleaseBuffer(m, v1643)
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L12
	} else {
		goto L327
	}
L309:
	;
	v1584 = v1578
	goto L311
L310:
	;
	v1584 = v1581
	goto L311
L311:
	;
	v1586 = v1584 - int32(4)
	if v1586 == int32(0) {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v1641 = int32(0)
	goto L308
L313:
	;
	goto L314
L314:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1580) {
		goto L316
	} else {
		goto L317
	}
L315:
	;
	v1641 = v1586
	goto L308
L316:
	;
	v1597 = int32(base.Ui32(v1580+int32(_a_F_heap2_redo_9)) >> (uint(int32(2)) % 32))
	goto L318
L317:
	;
	v1597 = int32(0)
	goto L318
L318:
	;
	if base.Ui32(v1597&int32(_a_F_heap2_redo_10)) < base.Ui32(int32(291)) {
		goto L315
	} else {
		goto L319
	}
L319:
	;
	v1602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1574)+10)))
	if v1602&int32(1) == int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1641 = int32(0)
	goto L308
L321:
	;
	goto L322
L322:
	;
	v1611 = int32(1)
	goto L323
L323:
	;
	v1620 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1574+int32(20)+v1611&int32(_a_F_heap2_redo_10)<<(uint(int32(2))%32))+1)))
	if v1620&int32(384) == int32(0) {
		goto L315
	} else {
		goto L325
	}
L324:
	;
	v1641 = int32(0)
	goto L308
L325:
	;
	v1626 = v1611 + int32(1)
	v1627 = int32(_a_F_heap2_redo_10)
	if base.Ui32(v1626&v1627) <= base.Ui32(v1597&v1627) {
		v1611 = v1626
		goto L323
	} else {
		goto L326
	}
L326:
	;
	goto L324
L327:
	;
	if base.B2i32(v1642&int32(32) == int32(0))&base.B2i32(base.Ui32(int32(1637)) < base.Ui32(v1641)) != 0 {
		goto L299
	} else {
		goto L328
	}
L328:
	;
	v1653 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v1653
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v1655
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	F_XLogRecordPageWithFreeSpace(m, v23-int32(-64), v1659, v1641)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L12
	} else {
		goto L329
	}
L329:
	;
	goto L299
L330:
	;
	goto L299
L331:
	;
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	if v1768 == int32(0) {
		goto L6
	} else {
		goto L355
	}
L332:
	;
	v1677 = F_XLogReadBufferForRedoExtended(m, l0, int32(1), int32(3), int32(0), v23+int32(164))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L12
	} else {
		goto L333
	}
L333:
	;
	if v1677 != 0 {
		goto L331
	} else {
		goto L334
	}
L334:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	if v1679 < int32(0) {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	v1698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1697)+14)))
	if v1698 == int32(0) {
		goto L339
	} else {
		goto L340
	}
L336:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1683+(v1679^int32(-1))<<(uint(int32(2))%32))))
	v1697 = v1689
	goto L335
L337:
	;
	goto L338
L338:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1697 = v1691 + v1679<<(uint(int32(13))%32) + int32(-8192)
	goto L335
L339:
	;
	v1701 = int32(_a_F_heap2_redo_12)
	v1702 = int32(0)
	if v1702|(v1697&int32(3)|int32(1)) == v1702 {
		goto L344
	} else {
		goto L345
	}
L340:
	;
	v1752 = v1679
	goto L341
L341:
	;
	v1753 = *(*int64)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v1753
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+56)) = v1755
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[7])))
	v1759 = F_visibilitymap_set(m, v1757, v1752, int32(3))
	mBase = m.M
	v1760 = m.ExcPending
	if v1760 != 0 {
		goto L12
	} else {
		goto L353
	}
L342:
	;
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v23)+164))
	v1752 = v1751
	goto L341
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1697)+10)) = int32(_a_F_heap2_redo_13)
	v1742 = int32(_a_F_heap2_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1697)+18)) = uint16(v1742)
	v1748 = int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1697)+16)) = uint16(v1748)
	*(*uint16)(unsafe.Add(mBase, uint32(v1697)+14)) = uint16(v1748)
	goto L342
L344:
	;
	goto L347
L345:
	;
	goto L346
L346:
	;
	goto L352
L347:
	;
	v1719 = v1697 + v1701
	v1721 = v1697 + int32(4)
	if base.Ui32(v1721) < base.Ui32(v1719) {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v1723 = v1719
	goto L350
L349:
	;
	v1723 = v1721
	goto L350
L350:
	;
	v1728 = (v1697^int32(-1)+v1723)&int32(-4) + int32(4)
	if v1728 == int32(0) {
		goto L343
	} else {
		goto L351
	}
L351:
	;
	base.MemoryFill(m, v1697, int32(0), v1728)
	goto L343
L352:
	;
	base.MemoryFill(m, v1697, int32(0), v1701)
	goto L343
L353:
	;
	if v1759 == int32(3) {
		goto L331
	} else {
		goto L354
	}
L354:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1697))) = base.I64_rotl(v1126, int64(32))
	goto L331
L355:
	;
	F_UnlockReleaseBuffer(m, v1768)
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L12
	} else {
		goto L356
	}
L356:
	;
	goto L6
L357:
	;
	v1778 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v1778, v23+int32(168), v1778, v23+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v1785 = m.ExcPending
	if v1785 != 0 {
		goto L12
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1799 = F_XLogReadBufferForRedo(m, l0, int32(0), v23+int32(168))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L12
	} else {
		goto L362
	}
L360:
	;
	v1786 = *(*int64)(unsafe.Add(mBase, uint32(v23)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v1786
	v1788 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v1788
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_c_F_heap2_redo[2])))
	F_heap_xlog_vm_clear(m, l0, v23+int32(96), v1792, int32(2))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L12
	} else {
		goto L361
	}
L361:
	;
	goto L359
L362:
	;
	if v1799 == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v1803 < int32(0) {
		goto L367
	} else {
		goto L368
	}
L364:
	;
	goto L365
L365:
	;
	v1901 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	if v1901 == int32(0) {
		goto L6
	} else {
		goto L383
	}
L366:
	;
	v1822 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1774)+4)))
	if v1822 == int32(0) {
		goto L2
	} else {
		goto L370
	}
L367:
	;
	v1807 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1813 = *(*int32)(unsafe.Add(mBase, uint32(v1807+(v1803^int32(-1))<<(uint(int32(2))%32))))
	v1821 = v1813
	goto L366
L368:
	;
	goto L369
L369:
	;
	v1815 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[4]))
	v1821 = v1815 + v1803<<(uint(int32(13))%32) + int32(-8192)
	goto L366
L370:
	;
	v1825 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1821)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1825) {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v1833 = int32(base.Ui32(v1825+int32(_a_F_heap2_redo_9)) >> (uint(int32(2)) % 32))
	goto L373
L372:
	;
	v1833 = int32(0)
	goto L373
L373:
	;
	if base.Ui32(v1833&int32(_a_F_heap2_redo_10)) < base.Ui32(v1822) {
		goto L2
	} else {
		goto L374
	}
L374:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1821+v1822<<(uint(int32(2))%32))+20))
	if v1840&int32(_a_F_heap2_redo_7) != int32(_a_F_heap2_redo_17) {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v1847 = v1821 + v1840&int32(_a_F_heap2_redo_5)
	v1848 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1847)+20)))
	v1850 = v1848 & int32(_a_F_heap2_redo_18)
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+20)) = uint16(v1850)
	v1852 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1847)+18)))
	v1854 = v1852 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+18)) = uint16(v1854)
	v1856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+18)) = uint16(v1854)
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+20)) = uint16(v1850)
	if v1856&int32(15) != 0 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1861 = int32(1)
	v1880 = v1856<<(uint(v1861)%32)&int32(16) | (v1856<<(uint(int32(4))%32)&int32(64) | (v1856<<(uint(int32(6))%32)&int32(128) | v1856&v1861<<(uint(int32(12))%32))) | v1850
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+20)) = uint16(v1880)
	goto L378
L377:
	;
	goto L378
L378:
	;
	if v1856&int32(16) != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1885 = v1852 | int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1847)+18)) = uint16(v1885)
	goto L381
L380:
	;
	goto L381
L381:
	;
	v1887 = *(*int32)(unsafe.Add(mBase, uint32(v1774)))
	*(*int32)(unsafe.Add(mBase, uint32(v1847)+4)) = v1887
	*(*int64)(unsafe.Add(mBase, uint32(v1821))) = base.I64_rotl(v1773, int64(32))
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v23)+168))
	F_MarkBufferDirty(m, v1892)
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L12
	} else {
		goto L382
	}
L382:
	;
	goto L365
L383:
	;
	F_UnlockReleaseBuffer(m, v1901)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L12
	} else {
		goto L384
	}
L384:
	;
	goto L6
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v28
	F_errmsg_internal(m, int32(_a_F_heap2_redo_19), v23+int32(112))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L12
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(1374), int32(_a_F_heap2_redo_21))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L12
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	v1949 = F_OpenTransientFile(m, v1941, int32(65))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L12
	} else {
		goto L392
	}
L389:
	;
	goto L6
L390:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L12
	} else {
		goto L435
	}
L391:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
		goto L12
	} else {
		goto L431
	}
L392:
	;
	if int32(0) <= v1949 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1954))) = int32(167772200)
	v1957 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	v1958 = F_ftruncate(m, v1949, v1957)
	mBase = m.M
	if v1958 != 0 {
		goto L391
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L12
	} else {
		goto L427
	}
L396:
	;
	v1959 = int32(_a_F_heap2_redo_22)
	v1960 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	v1961 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1960))) = v1961
	v1963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1964 = *(*int32)(unsafe.Add(mBase, uint32(v1963)+64))
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1926)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10])) = v1961
	v1970 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1970))) = int32(167772198)
	v1976 = v1965 * int32(36)
	v1977 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	v1978 = F_pwrite(m, v1949, v1964+int32(40), v1976, v1977)
	mBase = m.M
	if v1978 != v1976 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10]))
	if v1981 == int32(0) {
		goto L400
	} else {
		goto L401
	}
L398:
	;
	goto L399
L399:
	;
	v2006 = int32(_a_F_heap2_redo_22)
	v2007 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	v2008 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2007))) = v2008
	v2011 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v2011))) = int32(167772197)
	v2016 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[11])))
	if v2016 != int32(1) {
		v2030 = v2008
		goto L409
	} else {
		goto L410
	}
L400:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10])) = int32(51)
	goto L402
L401:
	;
	goto L402
L402:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L12
	} else {
		goto L403
	}
L403:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L12
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+48)) = v1923 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_23), v1923+int32(48))
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L12
	} else {
		goto L405
	}
L405:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_24), int32(1125), int32(_a_F_heap2_redo_25))
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L12
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
	v2059 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v2059))) = int32(0)
	v2062 = F_CloseTransientFile(m, v1949)
	mBase = m.M
	v2063 = m.ExcPending
	if v2063 != 0 {
		goto L12
	} else {
		goto L425
	}
L408:
	;
	if v2030 == int32(0) {
		goto L407
	} else {
		goto L415
	}
L409:
	;
	goto L408
L410:
	;
	goto L411
L411:
	;
	v2021 = F_fsync(m, v1949)
	mBase = m.M
	if v2021 != int32(-1) {
		v2030 = v2021
		goto L409
	} else {
		goto L413
	}
L412:
	;
	v2030 = int32(-1)
	goto L409
L413:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10]))
	if v2025 == int32(27) {
		goto L411
	} else {
		goto L414
	}
L414:
	;
	goto L412
L415:
	;
	v2036 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[12])))
	if v2036 != 0 {
		goto L417
	} else {
		goto L418
	}
L416:
	;
	v2039 = F_errstart(m, v2037, int32(0))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L12
	} else {
		goto L420
	}
L417:
	;
	v2037 = int32(21)
	goto L419
L418:
	;
	v2037 = int32(24)
	goto L419
L419:
	;
	goto L416
L420:
	;
	if v2039 == int32(0) {
		goto L407
	} else {
		goto L421
	}
L421:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L12
	} else {
		goto L422
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+32)) = v1923 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_26), v1923+int32(32))
	mBase = m.M
	v2052 = m.ExcPending
	if v2052 != 0 {
		goto L12
	} else {
		goto L423
	}
L423:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_24), int32(1138), int32(_a_F_heap2_redo_25))
	mBase = m.M
	v2057 = m.ExcPending
	if v2057 != 0 {
		goto L12
	} else {
		goto L424
	}
L424:
	;
	goto L407
L425:
	;
	if v2062 != 0 {
		goto L390
	} else {
		goto L426
	}
L426:
	;
	m.G0 = v1923 + int32(1136)
	goto L389
L427:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L12
	} else {
		goto L428
	}
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1923))) = v1923 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_27), v1923)
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L12
	} else {
		goto L429
	}
L429:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_24), int32(1097), int32(_a_F_heap2_redo_25))
	mBase = m.M
	v2083 = m.ExcPending
	if v2083 != 0 {
		goto L12
	} else {
		goto L430
	}
L430:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L431:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L12
	} else {
		goto L432
	}
L432:
	;
	v2090 = *(*int64)(unsafe.Add(mBase, uint32(v1926)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v1923)+68)) = uint32(v2090)
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+64)) = v1923 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_28), v1923-int32(-64))
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L12
	} else {
		goto L433
	}
L433:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_24), int32(1108), int32(_a_F_heap2_redo_25))
	mBase = m.M
	v2104 = m.ExcPending
	if v2104 != 0 {
		goto L12
	} else {
		goto L434
	}
L434:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L435:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v2110 = m.ExcPending
	if v2110 != 0 {
		goto L12
	} else {
		goto L436
	}
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1923)+16)) = v1923 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_29), v1923+int32(16))
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L12
	} else {
		goto L437
	}
L437:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_24), int32(1144), int32(_a_F_heap2_redo_25))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L12
	} else {
		goto L438
	}
L438:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L439:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_30), int32(0))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L12
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(611), int32(_a_F_heap2_redo_31))
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L12
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_32), int32(0))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L12
	} else {
		goto L443
	}
L443:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(641), int32(_a_F_heap2_redo_31))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L12
	} else {
		goto L444
	}
L444:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L445:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_33), int32(0))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L12
	} else {
		goto L446
	}
L446:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(644), int32(_a_F_heap2_redo_31))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L12
	} else {
		goto L447
	}
L447:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L448:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_34), int32(0))
	mBase = m.M
	v2234 = m.ExcPending
	if v2234 != 0 {
		goto L12
	} else {
		goto L449
	}
L449:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(1229), int32(_a_F_heap2_redo_35))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L12
	} else {
		goto L450
	}
L450:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L451:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_36), int32(0))
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L12
	} else {
		goto L452
	}
L452:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_20), int32(1232), int32(_a_F_heap2_redo_35))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L12
	} else {
		goto L453
	}
L453:
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = l0 + int32(68)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
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
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_LockBufferInternal(m, v24, int32(1))
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
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(v28^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L8
L10:
	;
	goto L11
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup[1]))
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
	v73 = int32(_a_F_heapgettup_0)
	v74 = int32(base.Ui32(v68+int32(_a_F_heapgettup_1))>>(uint(int32(2))%32)) & v73
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	v79 = (v75 - int32(1)) & v73
	if base.Ui32(v74) < base.Ui32(v79) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v59 = int32(base.Ui32(v49+int32(_a_F_heapgettup_1))>>(uint(int32(2))%32)) & int32(_a_F_heapgettup_0)
	goto L17
L16:
	;
	v59 = int32(0)
	goto L17
L17:
	;
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
	v61 = int32(1)
	v62 = v60 + v61
	v87 = v59 - v62&int32(_a_F_heapgettup_0) + v61
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
	v101 = v275
	v102 = v276
	v103 = v277
	v104 = int32(0)
	goto L24
L28:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	F_UnlockBuffer(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L67
	}
L29:
	;
	v287 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v287
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v287
	v291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v291
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)) = uint8(v291)
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
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v115 == int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	F_LockBufferInternal(m, v115, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
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
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup[0]))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v126+(v122^int32(-1))<<(uint(int32(2))%32))))
	v140 = v132
	goto L36
L38:
	;
	goto L39
L39:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_heapgettup[1]))
	v140 = v134 + v122<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	v149 = int32(base.Ui32(v141+int32(_a_F_heapgettup_1)) >> (uint(int32(2)) % 32))
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
	v101 = v149 & int32(_a_F_heapgettup_0)
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
	v275 = v101
	v276 = v102
	v277 = v103
	goto L48
L48:
	;
	F_UnlockBuffer(m, v276)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L6
	} else {
		goto L66
	}
L49:
	;
	v178 = v105 + int32(20) + v166&int32(_a_F_heapgettup_0)<<(uint(int32(2))%32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179&int32(_a_F_heapgettup_2) != int32(_a_F_heapgettup_3) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v275 = v267
	v276 = v270
	v277 = v265
	goto L48
L51:
	;
	v265 = l1 + v166
	v266 = int32(1)
	v267 = v164 - v266
	if v266 < v164 {
		v164 = v267
		v166 = v265
		goto L49
	} else {
		goto L65
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v105 + v179&int32(_a_F_heapgettup_4)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v166)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v190)
	v193 = int32(base.Ui32(v190) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v193)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(base.Ui32(v188) >> (uint(int32(17)) % 32))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v200 = F_HeapTupleSatisfiesVisibility(m, v20, v198, v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_HeapCheckForSerializableConflictOut(m, v200, v202, v20, v203, v204)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	if v200 == int32(0) {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v209 = int32(0)
	if base.B2i32(l3 == v209)|base.B2i32(l2 == v209) != 0 {
		goto L28
	} else {
		goto L56
	}
L56:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+52))
	v220 = l3
	v221 = l2
	goto L57
L57:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220))))
	if v230&int32(1) != 0 {
		goto L51
	} else {
		goto L59
	}
L58:
	;
	goto L28
L59:
	;
	v233 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220)+4)))
	v236 = F_heap_getattr_1(m, v20, v233, v215, v17+int32(15))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+15)))
	if v238 != 0 {
		goto L51
	} else {
		goto L61
	}
L61:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v220)+48))
	v243 = F_FunctionCall2Coll(m, v220+int32(16), v241, v236, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L62
	}
L62:
	;
	if v243 == int64(0) {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v250 = v221 - int32(1)
	if v250 != 0 {
		v220 = v220 + int32(56)
		v221 = v250
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	goto L50
L66:
	;
	goto L27
L67:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)) = uint16(v166)
	goto L26
}
func F_hemdist_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	v5 = Fn14312(m, l0, l1, l2, int32(2))
	return v5
}
func F_hemdistcache_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v128 int64
	_ = v128
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v279 int64
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int64
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int64
	_ = v308
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int64
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int64
	_ = v341
	var v342 int32
	_ = v342
	var v343 int64
	_ = v343
	var v344 int32
	_ = v344
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v347 int64
	_ = v347
	var v351 int64
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int64
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v370 int32
	_ = v370
	var v371 int64
	_ = v371
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int64
	_ = v380
	var v390 int64
	_ = v390
	var v398 int64
	_ = v398
	var v411 int64
	_ = v411
	v8 = int64(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(1) {
		if v9&int32(1) != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if int32(7) < l2 {
				v247 = int64(0)
				v248 = int32(0)
				if l2 == v248 {
					v318 = int64(0)
				} else {
					v255 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v260 = v19
						v262 = v247
						v265 = v248
						for {
							v266 = int32(4)
							v267 = v260 + v266
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+3)))
							v269 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_c_F_hemdistcache_2[0]))))
							v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+2)))
							v271 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v270)+uint32(_c_F_hemdistcache_2[0]))))
							v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
							v273 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v272)+uint32(_c_F_hemdistcache_2[0]))))
							v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
							v275 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v274)+uint32(_c_F_hemdistcache_2[0]))))
							v279 = v269 + (v271 + (v273 + (v262 + v275)))
							v281 = v265 + v266
							if v281 != l2&int32(-4) {
								v260 = v267
								v262 = v279
								v265 = v281
								continue
							} else {
								break
							}
							break
						}
						if v255 == int32(0) {
							v308 = v279
						} else {
							v285 = v267
							v287 = v279
							v292 = v285
							v293 = int32(0)
							v294 = v287
							for {
								v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
								v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v298)+uint32(_c_F_hemdistcache_2[0]))))
								v300 = v294 + v299
								v301 = int32(1)
								v304 = v293 + v301
								if v304 != v255 {
									v292 = v292 + v301
									v293 = v304
									v294 = v300
									continue
								} else {
									break
								}
								break
							}
							v308 = v300
						}
					} else {
						v285 = v19
						v287 = v247
						v292 = v285
						v293 = int32(0)
						v294 = v287
						for {
							v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
							v299 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v298)+uint32(_c_F_hemdistcache_2[0]))))
							v300 = v294 + v299
							v301 = int32(1)
							v304 = v293 + v301
							if v304 != v255 {
								v292 = v292 + v301
								v293 = v304
								v294 = v300
								continue
							} else {
								break
							}
							break
						}
						v308 = v300
					}
					v318 = v308
				}
				v411 = v318
			} else {
				if l2 == int32(0) {
					v411 = v8
				} else {
					v25 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v31 = v19
						v32 = int32(0)
						v38 = v8
						for {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)))
							v42 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v39)+uint32(_c_F_hemdistcache_2[0]))))
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)))
							v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_hemdistcache_2[0]))))
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
							v50 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v47)+uint32(_c_F_hemdistcache_2[0]))))
							v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
							v54 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v51)+uint32(_c_F_hemdistcache_2[0]))))
							v58 = v42 + (v46 + (v50 + (v38 + v54)))
							v59 = int32(4)
							v60 = v31 + v59
							v62 = v32 + v59
							if v62 != l2&int32(-4) {
								v31 = v60
								v32 = v62
								v38 = v58
								continue
							} else {
								break
							}
							break
						}
						if v25 == int32(0) {
							v411 = v58
						} else {
							v66 = v60
							v73 = v58
							v75 = v66
							v76 = int32(0)
							v82 = v73
							for {
								v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
								v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_hemdistcache_2[0]))))
								v87 = v82 + v86
								v88 = int32(1)
								v91 = v76 + v88
								if v91 != v25 {
									v75 = v75 + v88
									v76 = v91
									v82 = v87
									continue
								} else {
									break
								}
								break
							}
							v411 = v87
						}
					} else {
						v66 = v19
						v73 = v8
						v75 = v66
						v76 = int32(0)
						v82 = v73
						for {
							v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
							v86 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v83)+uint32(_c_F_hemdistcache_2[0]))))
							v87 = v82 + v86
							v88 = int32(1)
							v91 = v76 + v88
							if v91 != v25 {
								v75 = v75 + v88
								v76 = v91
								v82 = v87
								continue
							} else {
								break
							}
							break
						}
						v411 = v87
					}
				}
			}
			return l2<<(uint(int32(3))%32) + (base.I32_wrap_i64(v411) ^ int32(-1))
		}
	} else {
		if v9&int32(1) != 0 {
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if int32(7) < l2 {
				v319 = int64(0)
				v320 = int32(0)
				if l2 == v320 {
					v390 = int64(0)
				} else {
					v327 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v332 = v97
						v334 = v319
						v337 = v320
						for {
							v338 = int32(4)
							v339 = v332 + v338
							v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+3)))
							v341 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v340)+uint32(_c_F_hemdistcache_2[0]))))
							v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+2)))
							v343 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v342)+uint32(_c_F_hemdistcache_2[0]))))
							v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332)+1)))
							v345 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v344)+uint32(_c_F_hemdistcache_2[0]))))
							v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332))))
							v347 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v346)+uint32(_c_F_hemdistcache_2[0]))))
							v351 = v341 + (v343 + (v345 + (v334 + v347)))
							v353 = v337 + v338
							if v353 != l2&int32(-4) {
								v332 = v339
								v334 = v351
								v337 = v353
								continue
							} else {
								break
							}
							break
						}
						if v327 == int32(0) {
							v380 = v351
						} else {
							v357 = v339
							v359 = v351
							v364 = v357
							v365 = int32(0)
							v366 = v359
							for {
								v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
								v371 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v370)+uint32(_c_F_hemdistcache_2[0]))))
								v372 = v366 + v371
								v373 = int32(1)
								v376 = v365 + v373
								if v376 != v327 {
									v364 = v364 + v373
									v365 = v376
									v366 = v372
									continue
								} else {
									break
								}
								break
							}
							v380 = v372
						}
					} else {
						v357 = v97
						v359 = v319
						v364 = v357
						v365 = int32(0)
						v366 = v359
						for {
							v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
							v371 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v370)+uint32(_c_F_hemdistcache_2[0]))))
							v372 = v366 + v371
							v373 = int32(1)
							v376 = v365 + v373
							if v376 != v327 {
								v364 = v364 + v373
								v365 = v376
								v366 = v372
								continue
							} else {
								break
							}
							break
						}
						v380 = v372
					}
					v390 = v380
				}
				v398 = v390
			} else {
				if l2 == int32(0) {
					v398 = v8
				} else {
					v103 = l2 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(l2) {
						v109 = v97
						v110 = int32(0)
						v116 = v8
						for {
							v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+3)))
							v120 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_hemdistcache_2[0]))))
							v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+2)))
							v124 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_hemdistcache_2[0]))))
							v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
							v128 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_hemdistcache_2[0]))))
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
							v132 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v129)+uint32(_c_F_hemdistcache_2[0]))))
							v136 = v120 + (v124 + (v128 + (v116 + v132)))
							v137 = int32(4)
							v138 = v109 + v137
							v140 = v110 + v137
							if v140 != l2&int32(-4) {
								v109 = v138
								v110 = v140
								v116 = v136
								continue
							} else {
								break
							}
							break
						}
						if v103 == int32(0) {
							v398 = v136
						} else {
							v144 = v138
							v151 = v136
							v153 = v144
							v154 = int32(0)
							v160 = v151
							for {
								v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
								v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v161)+uint32(_c_F_hemdistcache_2[0]))))
								v165 = v160 + v164
								v166 = int32(1)
								v169 = v154 + v166
								if v169 != v103 {
									v153 = v153 + v166
									v154 = v169
									v160 = v165
									continue
								} else {
									break
								}
								break
							}
							v398 = v165
						}
					} else {
						v144 = v97
						v151 = v8
						v153 = v144
						v154 = int32(0)
						v160 = v151
						for {
							v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153))))
							v164 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v161)+uint32(_c_F_hemdistcache_2[0]))))
							v165 = v160 + v164
							v166 = int32(1)
							v169 = v154 + v166
							if v169 != v103 {
								v153 = v153 + v166
								v154 = v169
								v160 = v165
								continue
							} else {
								break
							}
							break
						}
						v398 = v165
					}
				}
			}
			return l2<<(uint(int32(3))%32) + (base.I32_wrap_i64(v398) ^ int32(-1))
		} else {
			if l2 <= int32(0) {
				return int32(0)
			} else {
				v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v177 = int32(0)
				if l2 != int32(1) {
					v186 = v177
					v187 = v177
					v188 = int32(0)
					for {
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v176))))
						v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186+v175))))
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195^v197)+uint32(_c_F_hemdistcache_2[0]))))
						v204 = v186 | int32(1)
						v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v204))))
						v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204+v176))))
						v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206^v208)+uint32(_c_F_hemdistcache_2[0]))))
						v213 = v187 + v201 + v212
						v214 = int32(2)
						v215 = v186 + v214
						v217 = v188 + v214
						if v217 != l2&int32(2147483646) {
							v186 = v215
							v187 = v213
							v188 = v217
							continue
						} else {
							break
						}
						break
					}
					if l2&int32(1) == int32(0) {
						v239 = v213
					} else {
						v221 = v215
						v222 = v213
						v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v175))))
						v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v176))))
						v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230^v232)+uint32(_c_F_hemdistcache_2[0]))))
						v239 = v222 + v236
					}
				} else {
					v221 = v177
					v222 = v177
					v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v175))))
					v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221+v176))))
					v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230^v232)+uint32(_c_F_hemdistcache_2[0]))))
					v239 = v222 + v236
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
func F_hs_contained(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hs_contained_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
