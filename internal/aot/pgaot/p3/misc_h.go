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
						F_errmsg(m, int32(_a_F_hamming_distance_0), v6)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hamming_distance_1), int32(39), int32(_a_F_hamming_distance_2))
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
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_hamming_distance[0]))
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
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_child_exit_signal[5]))
	if v39 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
}
func F_handle_pm_reload_request_signal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	*(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[0])) = int32(1)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[1]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	if v7 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v10 == int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v13 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[2]))
	if v17 == v13 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[3]))
	if v24 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v47 = F_pgmem_kill(m, v13, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v21 + int32(16)
	goto L1
L10:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+15)) = uint8(v27)
	goto L11
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[4]))
	v35 = F_write(m, v31, v21+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v35 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_handle_pm_reload_request_signal[5]))
	if v39 == int32(27) {
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
	var v74 int64
	_ = v74
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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	v1 = l0
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[1]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+16)))
	if v19 != int32(1) {
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
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[2])))
	if v36 == v34 {
		v382 = v34
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
	v350 = int32(1)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v351 - v352 + v350
	v358 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v358, v13, int32(4))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L86
	}
L16:
	;
	v308 = int32(_a_F_handle_streamed_transaction_0)
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	v310 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v301+v309<<(uint(v310)%32)))) = v46
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	v320 = v301 + v317<<(uint(v310)%32)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v315)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v320+v310))) = v325
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v315)+32))
	v328 = int64(*(*int32)(unsafe.Add(mBase, uint32(v315)+40)))
	*(*int64)(unsafe.Add(mBase, uint32(v320+int32(8)))) = v327 + v328
	goto L85
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5])) = int32(128)
	v287 = int32(_a_F_handle_streamed_transaction_1)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6]))
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v291
	v294 = F_palloc(m, int32(2048))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L6
	} else {
		goto L80
	}
L22:
	;
	v143 = int32(_a_F_handle_streamed_transaction_2)
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[8])) = v145 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
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
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
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
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[0]))
	if v51 == v46 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[9]))
	if v54 == v46 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[10]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[9])) = v46
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	if v61 == int32(0) {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v74 = base.I64_extend_i32_u(v61)
	goto L29
L29:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v57+base.I32_wrap_i64(v74)<<(uint(int32(4))%32)-int32(16))))
	if v81 == v46 {
		goto L15
	} else {
		goto L31
	}
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5]))
	if v61 != v90 {
		v301 = v57
		goto L16
	} else {
		goto L33
	}
L31:
	;
	if base.B2i32(v74 < int64(2)) == int32(0) {
		v74 = v74 - int64(1)
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[5])) = v61 << (uint(int32(1)) % 32)
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
	v301 = v98
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
	v382 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
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
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
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
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
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
	v382 = base.B2i32(v1 != int32(82)) & base.B2i32(v1 != int32(89))
	goto L1
L43:
	;
	m.G0 = v153 + int32(96)
	v382 = int32(0)
	goto L1
L44:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11]))
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
		v190 = v158
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v196 = v190
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
		v190 = v182
		goto L49
	} else {
		goto L56
	}
L55:
	;
	v190 = v182
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
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[12]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v46
	v203 = v153 + int32(32)
	v208 = F_pg_snprintf(m, v203, int32(64), int32(_a_F_handle_streamed_transaction_3), v153+int32(16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v203
	F_errmsg_internal(m, int32(_a_F_handle_streamed_transaction_4), v153)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[13]))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+24))
	goto L66
L64:
	;
	F_errfinish(m, int32(_a_F_handle_streamed_transaction_5), int32(1380), int32(_a_F_handle_streamed_transaction_6))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	if base.B2i32(base.Ui32(int32(1)) < base.Ui32(v225)) == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[13]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+20))
	goto L70
L68:
	;
	goto L69
L69:
	;
	F_DefineSavepoint(m, v153+int32(32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L6
	} else {
		goto L77
	}
L70:
	;
	if base.B2i32(v232 == int32(2)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
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
	v240 = m.ExcPending
	if v240 != 0 {
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
	v242 = m.ExcPending
	if v242 != 0 {
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v249 = int32(_a_F_handle_streamed_transaction_1)
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6]))
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v253
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11]))
	v257 = F_lappend_xid(m, v256, v46)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v250
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[11])) = v257
	goto L43
L80:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	F_errmsg_internal(m, int32(_a_F_handle_streamed_transaction_7), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_handle_streamed_transaction_8), int32(584), int32(_a_F_handle_streamed_transaction_9))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
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
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[6])) = v288
	v301 = v294
	goto L16
L85:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[10])) = v301
	v333 = int32(_a_F_handle_streamed_transaction_0)
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[4])) = v335 + int32(1)
	goto L15
L86:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	F_BufFileWrite(m, v363, v13+int32(7), int32(1))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v371 = v369 - v370
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v371
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_handle_streamed_transaction[3]))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	F_BufFileWrite(m, v374, v370+v375, v371)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v382 = v350
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
			v16 = F_convert_any_priv_string(m, v13, int32(_a_F_has_largeobject_privilege_name_id_0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
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
					return int32(0)
				} else {
					if v26 != 0 {
						v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_name_id[1])))
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
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
	v90 = F_ExecFetchSlotMinimalTuple(m, v84, v12+int32(11))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L8
	} else {
		goto L22
	}
L2:
	;
	v84 = l2
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
	v34 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v74 = v72 & int32(_a_F_hashagg_spill_tuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)) = uint16(v74)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)) = uint16(v77)
	goto L21
L14:
	;
	v39 = int32(1)
	v41 = v34 + v39
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v43 = F_bms_is_member(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = v34 << (uint(int32(2)) % 32)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49+v46)))
	*(*int32)(unsafe.Add(mBase, uint32(v46+v47))) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v34))))
	v56 = v55
	goto L19
L18:
	;
	v56 = v39
	goto L19
L19:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v57+v34))) = uint8(v56)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v41 < v61 {
		v34 = v41
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v84 = v16
	goto L1
L22:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v93 <= int32(31) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v99 = int32(base.Ui32(v96&l3) >> (uint(v93) % 32))
	goto L25
L24:
	;
	v99 = int32(0)
	goto L25
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = v100 + v99<<(uint(int32(3))%32)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v103))) = v104 + int64(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v109 = int32(24)
	v111 = v108 + v99*v109
	v116 = int32(711645284)
	v119 = l3 - int32(1636608428) ^ v116 - int32(1455628627)
	v124 = v119 ^ int32(-1636608428) - base.I32_rotl(v119, int32(25))
	v129 = v124 ^ v116 - base.I32_rotl(v124, int32(16))
	v133 = v129 ^ v119 - base.I32_rotl(v129, int32(4))
	v137 = v133 ^ v124 - base.I32_rotl(v133, int32(14))
	v141 = v137 ^ v129 - base.I32_rotl(v137, v109)
	goto L26
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111))))
	v147 = int32(32) - v146
	v149 = v144 + int32(base.Ui32(v141)>>(uint(v147)%32))
	v150 = v141 << (uint(v146) % 32)
	if v150 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v99<<(uint(int32(2))%32))))
	F_LogicalTapeWrite(m, v178, v12+int32(12), int32(4))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L37
	}
L28:
	;
	v157 = int32(32) - (base.I32_clz(v150) ^ int32(31))
	v158 = int32(255)
	if base.Ui32(v147&v158) < base.Ui32(v157&v158) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v167 = v147 + int32(1)
	goto L30
L30:
	;
	v169 = v167 & int32(255)
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if base.Ui32(v170) < base.Ui32(v169) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v163 = v147 + int32(1)
	goto L33
L32:
	;
	v163 = v157
	goto L33
L33:
	;
	v167 = v163
	goto L30
L34:
	;
	v172 = v169
	goto L36
L35:
	;
	v172 = v170
	goto L36
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v172)
	goto L27
L37:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	F_LogicalTapeWrite(m, v178, v90, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	if v187 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pfree(m, v90)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	if v6 == int64(0) {
		v13 = int32(-1636608428)
		v52 = v13
		v53 = v13
		v56 = int32(0)
	} else {
		v16 = base.I32_wrap_i64(v6)
		v18 = v16 + int32(1021750440)
		v23 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v29 = v16 - v23 - int32(1636608428) ^ base.I32_rotl(v23, int32(6))
		v33 = v18 - v29 ^ base.I32_rotl(v29, int32(8))
		v34 = v23 + v18
		v35 = v29 + v34
		v36 = v33 + v35
		v40 = v34 - v33 ^ base.I32_rotl(v33, int32(16))
		v44 = v35 - v40 ^ base.I32_rotl(v40, int32(19))
		v49 = v40 + v36
		v50 = v44 + v49
		v52 = v50
		v53 = v49
		v56 = v36 - v44 ^ base.I32_rotl(v44, int32(4)) ^ v50
	}
	v57 = int32(14)
	v59 = v56 - base.I32_rotl(v52, v57)
	v64 = v59 ^ (base.B2i32(v2 != int32(0)) + v53) - base.I32_rotl(v59, int32(11))
	v68 = v52 ^ v64 - base.I32_rotl(v64, int32(25))
	v72 = v68 ^ v59 - base.I32_rotl(v68, int32(16))
	v76 = v72 ^ v64 - base.I32_rotl(v72, int32(4))
	v80 = v76 ^ v68 - base.I32_rotl(v76, v57)
	v90 = F_Int64GetDatum(m, base.I64_extend_i32_u(v80)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v80^v72-base.I32_rotl(v80, int32(24))))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return int32(0)
	} else {
		return v90
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
			v56 = v41 + v47
			v57 = v51 + v56
			v58 = v55 + v57
			v62 = v56 - v55 ^ base.I32_rotl(v55, int32(16))
			v66 = v57 - v62 ^ base.I32_rotl(v62, int32(19))
			v71 = v58 + v62
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
		v346 = F_Int64GetDatum(m, base.I64_extend_i32_u(v336)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v336^v328-base.I32_rotl(v336, int32(24))))
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
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_hba_authname[0])))
	return v4
}
func F_heap2_decode(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v196 int32
	_ = v196
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int64
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int64
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v334 int32
	_ = v334
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+96))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+48)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+36))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferProcessXid(m, v18, v19, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v23 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	switch int32(base.Ui32(v16)>>(uint(int32(4))%32))&int32(7) - int32(5) {
	case 0:
		goto L6
	default:
		goto L3
	case 2:
		goto L5
	}
L5:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v196 != 0 {
		goto L3
	} else {
		goto L43
	}
L6:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v33 = F_SnapBuildProcessChange(m, v17, v19, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v37 != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v38 = m.G0
	v40 = v38 - int32(16)
	m.G0 = v40
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+64))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45&int32(8) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	m.G0 = v40 + int32(16)
	return
L11:
	;
	v50 = int32(0)
	F_XLogRecGetBlockTag(m, v42, v50, v40, v50, v50)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+88))
	if v55 != v57 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+56)))
	v62 = F_filter_by_origin_cb_wrapper(m, l0, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v64 = int32(0)
	v66 = v40 + int32(12)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+72))
	if v69 < v64 {
		v91 = v64
		goto L20
	} else {
		goto L21
	}
L17:
	;
	if v62 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	if v95 == int32(0) {
		goto L10
	} else {
		goto L30
	}
L20:
	;
	v94 = v91
	goto L19
L21:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68+int32(0))+76)))
	if v74 != int32(1) {
		v91 = v64
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v78 = v68 + int32(76)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+43)))
	if v79 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v66 == int32(0) {
		v91 = v64
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v84
	v94 = v84
	goto L19
L27:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v66))) = v87
	goto L29
L28:
	;
	goto L29
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+44))
	v91 = v89
	goto L20
L30:
	;
	v102 = v94
	v108 = int32(0)
	goto L31
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v112 = F_ReorderBufferAllocChange(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L10
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)) = uint16(v117)
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v112)+20)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v127 = (v102 + int32(1)) & int32(-2)
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	v129 = F_ReorderBufferAllocTupleBuf(m, v123, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+40)) = v129
	v132 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+12)) = v132
	*(*uint16)(unsafe.Add(mBase, uint32(v129)+8)) = uint16(v132)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v128 + int32(23)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v142 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v141)+15)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v141)+8)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v141))) = v142
	v149 = v127 + int32(7)
	if v128 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	base.MemoryCopy(m, v150+int32(23), v149, v128)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+20)) = uint16(v154)
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v141)+18)) = uint16(v156)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+22)) = uint8(v158)
	v161 = v108 + int32(1)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v162&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v168 = base.B2i32(v161 == v165)
	goto L40
L39:
	;
	v168 = int32(0)
	goto L40
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+32)) = uint8(v168)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v42)+96))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+36))
	v173 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	F_ReorderBufferQueueChange(m, v170, v172, v173, v112, int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	if base.Ui32(v161) < base.Ui32(v178) {
		v102 = v128 + v149
		v108 = v161
		goto L31
	} else {
		goto L42
	}
L42:
	;
	goto L32
L43:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)+96))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+64))
	v201 = m.G0
	v203 = v201 - int32(32)
	m.G0 = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	F_ReorderBufferXidSetCatalogChanges(m, v205, v19, v197)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v200)+12))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v200)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+24)) = v213
	v215 = *(*int64)(unsafe.Add(mBase, uint32(v200)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v203)+16)) = v215
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v200)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v203)+8)) = v217
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v203)+12)) = uint16(v219)
	v222 = v203 + int32(16)
	v224 = v203 + int32(8)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v212)+124))
	v227 = F_MemoryContextAlloc(m, v225, int32(64))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v229 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v227)+56)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+48)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+40)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+32)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+24)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+16)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227)+8)) = v229
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = v229
	v246 = F_ReorderBufferTXNByXid(m, v212, v211, int32(0), v197)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+28)) = v248
	v250 = *(*int64)(unsafe.Add(mBase, uint32(v222)))
	*(*int64)(unsafe.Add(mBase, uint32(v227)+20)) = v250
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+32)) = v252
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v227)+36)) = uint16(v254)
	*(*int32)(unsafe.Add(mBase, uint32(v227)+48)) = v208
	*(*int32)(unsafe.Add(mBase, uint32(v227)+44)) = v209
	*(*int32)(unsafe.Add(mBase, uint32(v227)+40)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v227)+12)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v227))) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v227)+8)) = int32(7)
	v264 = v246 + int32(136)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v246)+140))
	if v265 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+140)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v246)+136)) = v264
	goto L49
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+56)) = v264
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v246)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+52)) = v271
	v274 = v227 + int32(52)
	*(*int32)(unsafe.Add(mBase, uint32(v271)+4)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v246)+136)) = v274
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v246)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v246)+144)) = v277 + int64(1)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	if v282 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v17)+56))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+124))
	v311 = F_MemoryContextAlloc(m, v309, int32(64))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	if base.Ui32(v281) < base.Ui32(v282) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	if v281 != int32(-1) {
		v305 = v281
		goto L50
	} else {
		goto L60
	}
L54:
	;
	v286 = v282
	goto L56
L55:
	;
	v286 = v281
	goto L56
L56:
	;
	if v281 == int32(-1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v289 = v282
	goto L59
L58:
	;
	v289 = v286
	goto L59
L59:
	;
	v305 = v289
	goto L50
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_decode_0), int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_heap2_decode_1), int32(716), int32(_a_F_heap2_decode_2))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v313 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v311)+16)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+8)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+56)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+48)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+40)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+32)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311)+24)) = v313
	*(*int64)(unsafe.Add(mBase, uint32(v311))) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v311)+20)) = v305 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v311)+8)) = int32(6)
	F_ReorderBufferQueueChange(m, v308, v19, v197, v311, int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	m.G0 = v203 + int32(32)
	goto L3
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
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
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v686 int32
	_ = v686
	var v705 int32
	_ = v705
	var v710 int32
	_ = v710
	var v725 int32
	_ = v725
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v802 int32
	_ = v802
	var v825 int32
	_ = v825
	var v827 int32
	_ = v827
	var v846 int32
	_ = v846
	var v854 int32
	_ = v854
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v905 int32
	_ = v905
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int64
	_ = v939
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int64
	_ = v951
	var v952 int32
	_ = v952
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int64
	_ = v970
	var v979 int32
	_ = v979
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1018 int32
	_ = v1018
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1068 int64
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1136 int32
	_ = v1136
	var v1150 int32
	_ = v1150
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1167 int64
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1195 int64
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int64
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1243 int32
	_ = v1243
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1299 int32
	_ = v1299
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1325 int32
	_ = v1325
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1356 int32
	_ = v1356
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1410 int32
	_ = v1410
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int64
	_ = v1419
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1504 int32
	_ = v1504
	var v1509 int32
	_ = v1509
	var v1518 int32
	_ = v1518
	var v1527 int32
	_ = v1527
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1548 int32
	_ = v1548
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1575 int32
	_ = v1575
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1597 int64
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1606 int64
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int64
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1657 int32
	_ = v1657
	var v1663 int32
	_ = v1663
	var v1665 int32
	_ = v1665
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1680 int32
	_ = v1680
	var v1687 int32
	_ = v1687
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1727 int32
	_ = v1727
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1738 int64
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1742 int32
	_ = v1742
	var v1749 int32
	_ = v1749
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int64
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1763 int64
	_ = v1763
	var v1771 int64
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1790 int64
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int32
	_ = v1798
	var v1803 int32
	_ = v1803
	var v1809 int32
	_ = v1809
	var v1810 int64
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1814 int32
	_ = v1814
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1844 int32
	_ = v1844
	var v1849 int32
	_ = v1849
	var v1854 int32
	_ = v1854
	var v1858 int32
	_ = v1858
	var v1863 int32
	_ = v1863
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1872 int32
	_ = v1872
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1885 int32
	_ = v1885
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1911 int32
	_ = v1911
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1923 int64
	_ = v1923
	var v1932 int32
	_ = v1932
	var v1937 int32
	_ = v1937
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1951 int32
	_ = v1951
	var v1956 int32
	_ = v1956
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
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
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(_a_F_heap2_redo_0)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v24 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+48)))
	switch int32(base.Ui32(v24)>>(uint(int32(4))%32))&int32(7) - int32(1) {
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
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L11
	} else {
		goto L415
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v2007 = m.ExcPending
	if v2007 != 0 {
		goto L11
	} else {
		goto L412
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L11
	} else {
		goto L409
	}
L4:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L11
	} else {
		goto L406
	}
L5:
	;
	m.G0 = v21 + int32(_a_F_heap2_redo_0)
	return
L6:
	;
	v1754 = m.G0
	v1756 = v1754 - int32(1136)
	m.G0 = v1756
	v1758 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+64))
	v1760 = *(*int64)(unsafe.Add(mBase, uint32(v1759)+4))
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1759)))
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1758)+36))
	v1763 = *(*int64)(unsafe.Add(mBase, uint32(v1759)+32))
	*(*uint32)(unsafe.Add(mBase, uint32(v1756)+96)) = uint32(v1763)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+104)) = v1762
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+100)) = v1761
	*(*int64)(unsafe.Add(mBase, uint32(v1756)+84)) = v1760
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+80)) = int32(_a_F_heap2_redo_1)
	v1771 = int64(base.Ui64(v1763) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1756)+92)) = uint32(v1771)
	v1774 = v1756 + int32(112)
	v1779 = F_pg_snprintf(m, v1774, int32(1024), int32(_a_F_heap2_redo_2), v1756+int32(80))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L11
	} else {
		goto L355
	}
L7:
	;
	v1606 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607)+7)))
	if v1608&int32(1) != 0 {
		goto L324
	} else {
		goto L325
	}
L8:
	;
	v1195 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v1197 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v1197, v21+int32(_a_F_heap2_redo_3), v1197, v21+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v1204 = m.ExcPending
	if v1204 != 0 {
		goto L11
	} else {
		goto L226
	}
L9:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v951 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v952 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0]))) = v952
	F_XLogRecGetBlockTag(m, l0, int32(1), v21+int32(168), v952, v21+int32(_a_F_heap2_redo_5))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L11
	} else {
		goto L163
	}
L10:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v33 = int32(0)
	F_XLogRecGetBlockTag(m, l0, v33, v21+int32(168), v33, v21+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v41&int32(8) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v63 = int32(0)
	v66 = v41 & int32(4)
	v71 = F_XLogReadBufferForRedoExtended(m, l0, v63, v63, int32(base.Ui32(v66)>>(uint(int32(2))%32)), v21+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[1]))
	if base.Ui32(v47) < base.Ui32(int32(2)) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+2))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v51
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v53
	F_ResolveRecoveryConflictWithSnapshot(m, v50, int32(base.Ui32(v41&int32(2))>>(uint(int32(1))%32)), v21+int32(24))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	if v71 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	if v75 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	if v846 == int32(0) {
		goto L5
	} else {
		goto L133
	}
L21:
	;
	v94 = int32(0)
	v96 = v21 + int32(140)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+72))
	if v99 < v94 {
		v121 = v94
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v75^int32(-1))<<(uint(int32(2))%32))))
	v93 = v85
	goto L21
L23:
	;
	goto L24
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v93 = v87 + v75<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	v126 = v21 + int32(144)
	v128 = v21 + int32(136)
	v130 = v21 + int32(132)
	v132 = v21 + int32(156)
	v136 = v21 + int32(152)
	v138 = v21 + int32(164)
	v140 = v21 + int32(148)
	v142 = v21 + int32(160)
	if v41&int32(16) != 0 {
		goto L37
	} else {
		goto L38
	}
L26:
	;
	v124 = v121
	goto L25
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98+int32(0))+76)))
	if v104 != int32(1) {
		v121 = v94
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v108 = v98 + int32(76)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+43)))
	if v109 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v96 == int32(0) {
		v121 = v94
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v114 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v114
	v124 = v114
	goto L25
L33:
	;
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v117
	goto L35
L34:
	;
	goto L35
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+44))
	v121 = v119
	goto L26
L36:
	;
	if v41&int32(32) != 0 {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v145 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v145
	v148 = v124 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v158 = v148 + v150*int32(12)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v154
	v158 = v124
	goto L36
L40:
	;
	if v41&int32(64) != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158))))
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v161
	v163 = int32(2)
	v164 = v158 + v163
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4]))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v174 = v164 + v166<<(uint(v163)%32)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4]))) = v170
	v174 = v158
	goto L40
L44:
	;
	if base.I32_extend8_s(v41) < int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v174))))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v177
	v180 = v174 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v190 = v180 + v182<<(uint(int32(1))%32)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v186
	v190 = v174
	goto L44
L48:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v21)+156))
	v211 = int32(0)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v21)+152))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v21)+148))
	if base.B2i32(base.B2i32(v211 < v210)|base.B2i32(v211 < v213) == v211)&base.B2i32(v219 <= v211) == v211 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v194
	v197 = v190 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v197 + v199<<(uint(int32(1))%32)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v204 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v142))) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v190
	goto L48
L52:
	;
	v225 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v21)+164))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v21)+160))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	if v230 < v225 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v21)+144))
	if int32(0) < v705 {
		goto L114
	} else {
		goto L115
	}
L55:
	;
	if v210 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234+(v230^int32(-1))<<(uint(int32(2))%32))))
	v248 = v240
	goto L55
L57:
	;
	goto L58
L58:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v248 = v242 + v230<<(uint(int32(13))%32) + int32(-8192)
	goto L55
L59:
	;
	if v213 <= int32(0) {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v252 = v248 + int32(20)
	if v210 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v262 = v227
	v264 = int32(0)
	goto L64
L62:
	;
	v307 = v227
	goto L63
L63:
	;
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307))))
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v252+v323<<(uint(int32(2))%32)))) = v327&int32(_a_F_heap2_redo_6) | int32(_a_F_heap2_redo_7)
	goto L59
L64:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262))))
	v279 = int32(2)
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+2)))
	v283 = int32(_a_F_heap2_redo_6)
	v285 = int32(_a_F_heap2_redo_7)
	*(*int32)(unsafe.Add(mBase, uint32(v252+v278<<(uint(v279)%32)))) = v282&v283 | v285
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+4)))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v252+v288<<(uint(v279)%32)))) = v292&v283 | v285
	v299 = v262 + int32(8)
	v301 = v264 + v279
	if v301 != v210&int32(2147483646) {
		v262 = v299
		v264 = v301
		goto L64
	} else {
		goto L66
	}
L65:
	;
	if v210&int32(1) == int32(0) {
		goto L59
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	v307 = v299
	goto L63
L68:
	;
	if v219 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L69:
	;
	v354 = v213 & int32(3)
	v356 = v248 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v213) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v362 = v228
	v366 = int32(0)
	goto L73
L71:
	;
	v411 = v228
	goto L72
L72:
	;
	v430 = v411
	v434 = int32(0)
	goto L77
L73:
	;
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362))))
	v381 = int32(2)
	v384 = int32(_a_F_heap2_redo_8)
	*(*int32)(unsafe.Add(mBase, uint32(v356+v380<<(uint(v381)%32)))) = v384
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v386<<(uint(v381)%32)))) = v384
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v392<<(uint(v381)%32)))) = v384
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v356+v398<<(uint(v381)%32)))) = v384
	v405 = v362 + int32(8)
	v407 = v366 + int32(4)
	if v407 != v213&int32(2147483644) {
		v362 = v405
		v366 = v407
		goto L73
	} else {
		goto L75
	}
L74:
	;
	if v354 == int32(0) {
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v411 = v405
	goto L72
L77:
	;
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v430))))
	v449 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v356+v448<<(uint(v449)%32)))) = int32(_a_F_heap2_redo_8)
	v457 = v434 + int32(1)
	if v457 != v354 {
		v430 = v430 + v449
		v434 = v457
		goto L77
	} else {
		goto L79
	}
L78:
	;
	goto L68
L79:
	;
	goto L78
L80:
	;
	if v66 == v225 {
		goto L93
	} else {
		goto L94
	}
L81:
	;
	v480 = v219 & int32(3)
	v482 = v248 + int32(20)
	if base.Ui32(int32(4)) <= base.Ui32(v219) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v490 = int32(0)
	v491 = v229
	goto L85
L83:
	;
	v540 = v229
	goto L84
L84:
	;
	v558 = int32(0)
	v559 = v540
	goto L89
L85:
	;
	v506 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491))))
	v507 = int32(2)
	v510 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v482+v506<<(uint(v507)%32)))) = v510
	v512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v482+v512<<(uint(v507)%32)))) = v510
	v518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v482+v518<<(uint(v507)%32)))) = v510
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v491)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v482+v524<<(uint(v507)%32)))) = v510
	v531 = v491 + int32(8)
	v533 = v490 + int32(4)
	if v533 != v219&int32(2147483644) {
		v490 = v533
		v491 = v531
		goto L85
	} else {
		goto L87
	}
L86:
	;
	if v480 == int32(0) {
		goto L80
	} else {
		goto L88
	}
L87:
	;
	goto L86
L88:
	;
	v540 = v531
	goto L84
L89:
	;
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v559))))
	v575 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v482+v574<<(uint(v575)%32)))) = int32(0)
	v583 = v558 + int32(1)
	if v583 != v480 {
		v558 = v583
		v559 = v559 + v575
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L80
L91:
	;
	goto L90
L92:
	;
	goto L54
L93:
	;
	v603 = int32(0)
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+12)))
	if base.Ui32(v609) < base.Ui32(int32(25)) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	F_PageRepairFragmentation(m, v248)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L11
	} else {
		goto L113
	}
L96:
	;
	goto L92
L97:
	;
	v674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+10)))
	v676 = v674 & int32(_a_F_heap2_redo_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+10)) = uint16(v676)
	goto L96
L98:
	;
	v617 = int32(base.Ui32(v609+int32(_a_F_heap2_redo_10))>>(uint(int32(2))%32)) & int32(_a_F_heap2_redo_11)
	if v617 == int32(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v623 = v617
	v625 = v603
	v628 = v603
	goto L101
L100:
	;
	if int32(0) < v652 {
		goto L109
	} else {
		goto L110
	}
L101:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v248+int32(20)+v623<<(uint(int32(2))%32))))
	v634 = v632 & int32(_a_F_heap2_redo_8)
	if base.B2i32(v623 == int32(1))|v628 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v652 = v646
	v654 = int32(0)
	goto L100
L103:
	;
	v649 = v623 - int32(1)
	if v649 != 0 {
		v623 = v649
		v625 = v646
		v628 = v647
		goto L101
	} else {
		goto L108
	}
L104:
	;
	v640 = int32(0)
	v646 = v625 + base.B2i32(v634 == v640)
	v647 = base.B2i32(v634 != v640)
	goto L103
L105:
	;
	goto L106
L106:
	;
	if v634 != 0 {
		v646 = v625
		v647 = v628
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v652 = v625
	v654 = int32(1)
	goto L100
L108:
	;
	goto L102
L109:
	;
	v659 = v609 - v652<<(uint(int32(2))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+12)) = uint16(v659)
	goto L111
L110:
	;
	goto L111
L111:
	;
	if v654 == int32(0) {
		goto L97
	} else {
		goto L112
	}
L112:
	;
	v663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+10)))
	v665 = v663 | int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v248)+10)) = uint16(v665)
	goto L96
L113:
	;
	goto L92
L114:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v21)+136))
	v725 = v2
	goto L117
L115:
	;
	goto L116
L116:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = base.I64_rotr(v31, int64(32))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	F_MarkBufferDirty(m, v825)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L11
	} else {
		goto L132
	}
L117:
	;
	v731 = v710 + v725*int32(12)
	v732 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731)+10)))
	if v732 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L116
L119:
	;
	v733 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731)+6)))
	v734 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731)+4)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v731)))
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+8)))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v21)+132))
	v745 = v742
	v746 = int32(0)
	goto L122
L120:
	;
	goto L121
L121:
	;
	v802 = v725 + int32(1)
	if v802 != v705 {
		v725 = v802
		goto L117
	} else {
		goto L131
	}
L122:
	;
	v761 = int32(2)
	v762 = v745 + v761
	*(*int32)(unsafe.Add(mBase, uint32(v21)+132)) = v762
	v764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v745))))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v93+int32(20)+v764<<(uint(v761)%32))))
	v771 = v93 + v768&int32(_a_F_heap2_redo_6)
	*(*int32)(unsafe.Add(mBase, uint32(v771)+4)) = v735
	if v736&int32(2) != 0 {
		goto L124
	} else {
		goto L125
	}
L123:
	;
	goto L121
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+8)) = int32(2)
	goto L126
L125:
	;
	goto L126
L126:
	;
	if v736&int32(4) != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+8)) = int32(0)
	goto L129
L128:
	;
	goto L129
L129:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v771)+18)) = uint16(v734)
	*(*uint16)(unsafe.Add(mBase, uint32(v771)+20)) = uint16(v733)
	v780 = v746 + int32(1)
	v781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v731)+10)))
	if base.Ui32(v780) < base.Ui32(v781) {
		v745 = v762
		v746 = v780
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
	goto L20
L133:
	;
	if base.Ui32(int32(32)) <= base.Ui32(v41) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	if v846 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L136
L136:
	;
	F_UnlockReleaseBuffer(m, v846)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L11
	} else {
		goto L162
	}
L137:
	;
	v872 = int32(4)
	v873 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v868)+14)))
	v874 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v868)+12)))
	v875 = v873 - v874
	if v875 <= v872 {
		goto L142
	} else {
		goto L143
	}
L138:
	;
	v854 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v854+(v846^int32(-1))<<(uint(int32(2))%32))))
	v868 = v860
	goto L137
L139:
	;
	goto L140
L140:
	;
	v862 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v868 = v862 + v846<<(uint(int32(13))%32) + int32(-8192)
	goto L137
L141:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	F_UnlockReleaseBuffer(m, v936)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L11
	} else {
		goto L160
	}
L142:
	;
	v878 = v872
	goto L144
L143:
	;
	v878 = v875
	goto L144
L144:
	;
	v880 = v878 - int32(4)
	if v880 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v935 = int32(0)
	goto L141
L146:
	;
	goto L147
L147:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v874) {
		goto L149
	} else {
		goto L150
	}
L148:
	;
	v935 = v880
	goto L141
L149:
	;
	v891 = int32(base.Ui32(v874+int32(_a_F_heap2_redo_10)) >> (uint(int32(2)) % 32))
	goto L151
L150:
	;
	v891 = int32(0)
	goto L151
L151:
	;
	if base.Ui32(v891&int32(_a_F_heap2_redo_11)) < base.Ui32(int32(291)) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868)+10)))
	if v896&int32(1) == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v935 = int32(0)
	goto L141
L154:
	;
	goto L155
L155:
	;
	v905 = int32(1)
	goto L156
L156:
	;
	v914 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v868+int32(20)+v905&int32(_a_F_heap2_redo_11)<<(uint(int32(2))%32))+1)))
	if v914&int32(384) == int32(0) {
		goto L148
	} else {
		goto L158
	}
L157:
	;
	v935 = int32(0)
	goto L141
L158:
	;
	v920 = v905 + int32(1)
	v921 = int32(_a_F_heap2_redo_11)
	if base.Ui32(v920&v921) <= base.Ui32(v891&v921) {
		v905 = v920
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v939 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v939
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v941
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_XLogRecordPageWithFreeSpace(m, v21+int32(8), v945, v935)
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L11
	} else {
		goto L161
	}
L161:
	;
	goto L5
L162:
	;
	goto L5
L163:
	;
	v963 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[1]))
	if base.Ui32(int32(2)) <= base.Ui32(v963) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v967 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+4)))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v968
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v970
	F_ResolveRecoveryConflictWithSnapshot(m, v966, int32(base.Ui32(v967&int32(4))>>(uint(int32(2))%32)), v21+int32(72))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L11
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v983 = F_XLogReadBufferForRedo(m, l0, int32(1), v21+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L11
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	if v983 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	if v987 < int32(0) {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	goto L171
L171:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	if v1030 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L172:
	;
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1005)+10)))
	v1008 = v1006 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1005)+10)) = uint16(v1008)
	v1011 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[6]))
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1011)+252))
	goto L177
L173:
	;
	v991 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v991+(v987^int32(-1))<<(uint(int32(2))%32))))
	v1005 = v997
	goto L172
L174:
	;
	goto L175
L175:
	;
	v999 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1005 = v999 + v987<<(uint(int32(13))%32) + int32(-8192)
	goto L172
L176:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_MarkBufferDirty(m, v1026)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L11
	} else {
		goto L182
	}
L177:
	;
	if base.B2i32(v1012 != int32(0)) == int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v1018 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[7])))
	if v1018&int32(1) == int32(0) {
		goto L176
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1005))) = base.I64_rotr(v951, int64(32))
	goto L176
L181:
	;
	goto L180
L182:
	;
	goto L171
L183:
	;
	v1078 = int32(0)
	v1083 = F_XLogReadBufferForRedoExtended(m, l0, v1078, int32(3), v1078, v21+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L11
	} else {
		goto L196
	}
L184:
	;
	if v1030 < int32(0) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v1051 = int32(4)
	v1052 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1050)+14)))
	v1053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1050)+12)))
	v1054 = v1052 - v1053
	if v1054 <= v1051 {
		goto L190
	} else {
		goto L191
	}
L186:
	;
	v1036 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1036+(v1030^int32(-1))<<(uint(int32(2))%32))))
	v1050 = v1042
	goto L185
L187:
	;
	goto L188
L188:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1050 = v1044 + v1030<<(uint(int32(13))%32) + int32(-8192)
	goto L185
L189:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_UnlockReleaseBuffer(m, v1060)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L11
	} else {
		goto L193
	}
L190:
	;
	v1057 = v1051
	goto L192
L191:
	;
	v1057 = v1054
	goto L192
L192:
	;
	goto L189
L193:
	;
	v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+4)))
	if v1063&int32(3) == int32(0) {
		goto L183
	} else {
		goto L194
	}
L194:
	;
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v1068
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v1070
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	F_XLogRecordPageWithFreeSpace(m, v21+int32(56), v1074, v1057-int32(4))
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L11
	} else {
		goto L195
	}
L195:
	;
	goto L183
L196:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	if v1083 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if v1085 < int32(0) {
		goto L201
	} else {
		goto L202
	}
L198:
	;
	goto L199
L199:
	;
	if v1085 == int32(0) {
		goto L5
	} else {
		goto L224
	}
L200:
	;
	v1106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+14)))
	if v1106 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L201:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1091+(v1085^int32(-1))<<(uint(int32(2))%32))))
	v1105 = v1097
	goto L200
L202:
	;
	goto L203
L203:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1105 = v1099 + v1085<<(uint(int32(13))%32) + int32(-8192)
	goto L200
L204:
	;
	v1109 = int32(_a_F_heap2_redo_12)
	v1110 = int32(0)
	if v1110|(v1105&int32(3)|int32(1)) == v1110 {
		goto L209
	} else {
		goto L210
	}
L205:
	;
	v1160 = v1085
	goto L206
L206:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v950)+4)))
	F_LockBuffer(m, v1160, int32(0))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L11
	} else {
		goto L218
	}
L207:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	v1160 = v1159
	goto L206
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1105)+10)) = int32(_a_F_heap2_redo_13)
	v1150 = int32(_a_F_heap2_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1105)+18)) = uint16(v1150)
	v1156 = int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1105)+16)) = uint16(v1156)
	*(*uint16)(unsafe.Add(mBase, uint32(v1105)+14)) = uint16(v1156)
	goto L207
L209:
	;
	goto L212
L210:
	;
	goto L211
L211:
	;
	goto L217
L212:
	;
	v1127 = v1105 + v1109
	v1129 = v1105 + int32(4)
	if base.Ui32(v1129) < base.Ui32(v1127) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1131 = v1127
	goto L215
L214:
	;
	v1131 = v1129
	goto L215
L215:
	;
	v1136 = (v1105^int32(-1)+v1131)&int32(-4) + int32(4)
	if v1136 == int32(0) {
		goto L208
	} else {
		goto L216
	}
L216:
	;
	base.MemoryFill(m, v1105, int32(0), v1136)
	goto L208
L217:
	;
	base.MemoryFill(m, v1105, int32(0), v1109)
	goto L208
L218:
	;
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v1165
	v1167 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = v1167
	v1171 = F_CreateFakeRelcacheEntry(m, v21+int32(40))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L11
	} else {
		goto L219
	}
L219:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	F_visibilitymap_pin(m, v1171, v1173, v21+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L11
	} else {
		goto L220
	}
L220:
	;
	v1178 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v950)))
	v1184 = F_visibilitymap_set(m, v1171, v1178, int32(0), v951, v1180, v1181, v1161&int32(3))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L11
	} else {
		goto L221
	}
L221:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	F_ReleaseBuffer(m, v1186)
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L11
	} else {
		goto L222
	}
L222:
	;
	F_pfree(m, v1171)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L11
	} else {
		goto L223
	}
L223:
	;
	goto L5
L224:
	;
	F_UnlockReleaseBuffer(m, v1085)
	mBase = m.M
	v1194 = m.ExcPending
	if v1194 != 0 {
		goto L11
	} else {
		goto L225
	}
L225:
	;
	goto L5
L226:
	;
	v1205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196))))
	if v1205&int32(1) != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+112)) = v1208
	v1210 = *(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+104)) = v1210
	v1214 = F_CreateFakeRelcacheEntry(m, v21+int32(104))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
		goto L11
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	if v24 < int32(0) {
		goto L237
	} else {
		goto L238
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = int32(0)
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_visibilitymap_pin(m, v1214, v1218, v21+int32(168))
	mBase = m.M
	v1222 = m.ExcPending
	if v1222 != 0 {
		goto L11
	} else {
		goto L231
	}
L231:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	v1226 = F_visibilitymap_clear(m, v1223, v1224, int32(3))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L11
	} else {
		goto L232
	}
L232:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	F_ReleaseBuffer(m, v1228)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L11
	} else {
		goto L233
	}
L233:
	;
	F_pfree(m, v1214)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L11
	} else {
		goto L234
	}
L234:
	;
	goto L229
L235:
	;
	v1589 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	if v1589 != 0 {
		goto L318
	} else {
		goto L319
	}
L236:
	;
	v1315 = int32(0)
	v1317 = v21 + int32(164)
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1319)+72))
	if v1320 < v1315 {
		v1342 = v1315
		goto L259
	} else {
		goto L260
	}
L237:
	;
	v1237 = F_XLogInitBufferForRedo(m, l0, int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L11
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v1308 = int32(0)
	v1312 = F_XLogReadBufferForRedo(m, l0, v1308, v21+int32(_a_F_heap2_redo_5))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L11
	} else {
		goto L256
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4]))) = v1237
	if v1237 < int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1258 = int32(_a_F_heap2_redo_12)
	v1259 = int32(0)
	if v1259|(v1257&int32(3)|int32(1)) == v1259 {
		goto L247
	} else {
		goto L248
	}
L242:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v1243+(v1237^int32(-1))<<(uint(int32(2))%32))))
	v1257 = v1249
	goto L241
L243:
	;
	goto L244
L244:
	;
	v1251 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1257 = v1251 + v1237<<(uint(int32(13))%32) + int32(-8192)
	goto L241
L245:
	;
	goto L236
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1257)+10)) = int32(_a_F_heap2_redo_13)
	v1299 = int32(_a_F_heap2_redo_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1257)+18)) = uint16(v1299)
	v1305 = int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1257)+16)) = uint16(v1305)
	*(*uint16)(unsafe.Add(mBase, uint32(v1257)+14)) = uint16(v1305)
	goto L245
L247:
	;
	goto L250
L248:
	;
	goto L249
L249:
	;
	goto L255
L250:
	;
	v1276 = v1257 + v1258
	v1278 = v1257 + int32(4)
	if base.Ui32(v1278) < base.Ui32(v1276) {
		goto L251
	} else {
		goto L252
	}
L251:
	;
	v1280 = v1276
	goto L253
L252:
	;
	v1280 = v1278
	goto L253
L253:
	;
	v1285 = (v1257^int32(-1)+v1280)&int32(-4) + int32(4)
	if v1285 == int32(0) {
		goto L246
	} else {
		goto L254
	}
L254:
	;
	base.MemoryFill(m, v1257, int32(0), v1285)
	goto L246
L255:
	;
	base.MemoryFill(m, v1257, int32(0), v1258)
	goto L246
L256:
	;
	if v1312 != 0 {
		v1573 = v1308
		v1575 = v2
		goto L235
	} else {
		goto L257
	}
L257:
	;
	goto L236
L258:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	if v1346 < int32(0) {
		goto L270
	} else {
		goto L271
	}
L259:
	;
	v1345 = v1342
	goto L258
L260:
	;
	v1325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1319+int32(0))+76)))
	if v1325 != int32(1) {
		v1342 = v1315
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1329 = v1319 + int32(76)
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+43)))
	if v1330 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	if v1317 == int32(0) {
		v1342 = v1315
		goto L259
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	if v1317 != 0 {
		goto L266
	} else {
		goto L267
	}
L265:
	;
	v1335 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1317))) = v1335
	v1345 = v1335
	goto L258
L266:
	;
	v1338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1329)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v1317))) = v1338
	goto L268
L267:
	;
	goto L268
L268:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+44))
	v1342 = v1340
	goto L259
L269:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v21)+164))
	v1366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1196)+2)))
	if v1366 != 0 {
		goto L273
	} else {
		goto L274
	}
L270:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v1350+(v1346^int32(-1))<<(uint(int32(2))%32))))
	v1364 = v1356
	goto L269
L271:
	;
	goto L272
L272:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1364 = v1358 + v1346<<(uint(int32(13))%32) + int32(-8192)
	goto L269
L273:
	;
	v1375 = v1345
	v1377 = v2
	goto L276
L274:
	;
	v1464 = v1345
	goto L275
L275:
	;
	if v1464 != v1345+v1365 {
		goto L2
	} else {
		goto L291
	}
L276:
	;
	if int32(0) <= v24 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1464 = v1457
	goto L275
L278:
	;
	v1394 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1196+int32(4)+v1377<<(uint(int32(1))%32)))))
	v1397 = v1394
	goto L280
L279:
	;
	v1397 = v1377 + int32(1)
	goto L280
L280:
	;
	v1399 = v1397 & int32(_a_F_heap2_redo_11)
	v1400 = int32(1)
	v1401 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+12)))
	if base.Ui32(v1401) < base.Ui32(int32(25)) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1410 = v1400
	goto L283
L282:
	;
	v1410 = int32(base.Ui32(v1401+int32(_a_F_heap2_redo_10))>>(uint(int32(2))%32)) + v1400
	goto L283
L283:
	;
	if base.Ui32(v1410&int32(_a_F_heap2_redo_11)) < base.Ui32(v1399) {
		goto L4
	} else {
		goto L284
	}
L284:
	;
	v1417 = (v1375 + int32(1)) & int32(-2)
	v1418 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1417))))
	v1419 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+183)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v21)+176)) = v1419
	*(*int64)(unsafe.Add(mBase, uint32(v21)+168)) = v1419
	v1426 = v1417 + int32(7)
	if v1418 != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	base.MemoryCopy(m, v21+int32(191), v1426, v1418)
	goto L287
L286:
	;
	goto L287
L287:
	;
	v1428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1417)+2)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+186)) = uint16(v1428)
	v1430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1417)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+188)) = uint16(v1430)
	v1432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+190)) = uint8(v1432)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v1434)+36))
	v1437 = v1430 & int32(_a_F_heap2_redo_15)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+188)) = uint16(v1437)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+176)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+168)) = v1435
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+182)) = uint16(v1442)
	v1445 = int32(base.Ui32(v1442) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+180)) = uint16(v1445)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+184)) = uint16(v1397)
	v1453 = F_PageAddItemExtended(m, v1364, v21+int32(168), v1418+int32(23), v1399, int32(3))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L11
	} else {
		goto L288
	}
L288:
	;
	if v1453 == int32(0) {
		goto L3
	} else {
		goto L289
	}
L289:
	;
	v1457 = v1418 + v1426
	v1459 = v1377 + int32(1)
	v1460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1196)+2)))
	if base.Ui32(v1459) < base.Ui32(v1460) {
		v1375 = v1457
		v1377 = v1459
		goto L276
	} else {
		goto L290
	}
L290:
	;
	goto L277
L291:
	;
	v1485 = int32(4)
	v1486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+14)))
	v1487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+12)))
	v1488 = v1486 - v1487
	if v1488 <= v1485 {
		goto L293
	} else {
		goto L294
	}
L292:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1364))) = base.I64_rotr(v1195, int64(32))
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196))))
	if v1552&int32(1) != 0 {
		goto L311
	} else {
		goto L312
	}
L293:
	;
	v1491 = v1485
	goto L295
L294:
	;
	v1491 = v1488
	goto L295
L295:
	;
	v1493 = v1491 - int32(4)
	if v1493 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1548 = int32(0)
	goto L292
L297:
	;
	goto L298
L298:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1487) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1548 = v1493
	goto L292
L300:
	;
	v1504 = int32(base.Ui32(v1487+int32(_a_F_heap2_redo_10)) >> (uint(int32(2)) % 32))
	goto L302
L301:
	;
	v1504 = int32(0)
	goto L302
L302:
	;
	if base.Ui32(v1504&int32(_a_F_heap2_redo_11)) < base.Ui32(int32(291)) {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	v1509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1364)+10)))
	if v1509&int32(1) == int32(0) {
		goto L304
	} else {
		goto L305
	}
L304:
	;
	v1548 = int32(0)
	goto L292
L305:
	;
	goto L306
L306:
	;
	v1518 = int32(1)
	goto L307
L307:
	;
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364+int32(20)+v1518&int32(_a_F_heap2_redo_11)<<(uint(int32(2))%32))+1)))
	if v1527&int32(384) == int32(0) {
		goto L299
	} else {
		goto L309
	}
L308:
	;
	v1548 = int32(0)
	goto L292
L309:
	;
	v1533 = v1518 + int32(1)
	v1534 = int32(_a_F_heap2_redo_11)
	if base.Ui32(v1533&v1534) <= base.Ui32(v1504&v1534) {
		v1518 = v1533
		goto L307
	} else {
		goto L310
	}
L310:
	;
	goto L308
L311:
	;
	v1555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+10)))
	v1557 = v1555 & int32(_a_F_heap2_redo_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v1364)+10)) = uint16(v1557)
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196))))
	v1560 = v1559
	goto L313
L312:
	;
	v1560 = v1552
	goto L313
L313:
	;
	if v1560&int32(32) != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1364)+10)))
	v1565 = v1563 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1364)+10)) = uint16(v1565)
	goto L316
L315:
	;
	goto L316
L316:
	;
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[4])))
	F_MarkBufferDirty(m, v1567)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L11
	} else {
		goto L317
	}
L317:
	;
	v1573 = v1548
	v1575 = int32(1)
	goto L235
L318:
	;
	F_UnlockReleaseBuffer(m, v1589)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L11
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	if v1575&base.B2i32(base.Ui32(v1573) < base.Ui32(int32(1638))) == int32(0) {
		goto L5
	} else {
		goto L322
	}
L321:
	;
	goto L320
L322:
	;
	v1597 = *(*int64)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+88)) = v1597
	v1599 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = v1599
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_XLogRecordPageWithFreeSpace(m, v21+int32(88), v1603, v1573)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L11
	} else {
		goto L323
	}
L323:
	;
	goto L5
L324:
	;
	v1611 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0]))) = v1611
	F_XLogRecGetBlockTag(m, l0, v1611, v21+int32(168), v1611, v21+int32(_a_F_heap2_redo_4))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L11
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	v1648 = F_XLogReadBufferForRedo(m, l0, int32(0), v21+int32(168))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L11
	} else {
		goto L333
	}
L327:
	;
	v1621 = *(*int32)(unsafe.Add(mBase, uint32(v21)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v1621
	v1623 = *(*int64)(unsafe.Add(mBase, uint32(v21)+168))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+120)) = v1623
	v1627 = F_CreateFakeRelcacheEntry(m, v21+int32(120))
	mBase = m.M
	v1628 = m.ExcPending
	if v1628 != 0 {
		goto L11
	} else {
		goto L328
	}
L328:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	F_visibilitymap_pin(m, v1627, v1629, v21+int32(_a_F_heap2_redo_3))
	mBase = m.M
	v1633 = m.ExcPending
	if v1633 != 0 {
		goto L11
	} else {
		goto L329
	}
L329:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[5])))
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	v1637 = F_visibilitymap_clear(m, v1634, v1635, int32(2))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L11
	} else {
		goto L330
	}
L330:
	;
	v1639 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_heap2_redo[0])))
	F_ReleaseBuffer(m, v1639)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L11
	} else {
		goto L331
	}
L331:
	;
	F_pfree(m, v1627)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L11
	} else {
		goto L332
	}
L332:
	;
	goto L326
L333:
	;
	if v1648 == int32(0) {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1607)+4)))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	if v1653 < int32(0) {
		goto L338
	} else {
		goto L339
	}
L335:
	;
	goto L336
L336:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	if v1749 == int32(0) {
		goto L5
	} else {
		goto L353
	}
L337:
	;
	v1672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1671)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1672) {
		goto L341
	} else {
		goto L342
	}
L338:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[2]))
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1657+(v1653^int32(-1))<<(uint(int32(2))%32))))
	v1671 = v1663
	goto L337
L339:
	;
	goto L340
L340:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[3]))
	v1671 = v1665 + v1653<<(uint(int32(13))%32) + int32(-8192)
	goto L337
L341:
	;
	v1680 = int32(base.Ui32(v1672+int32(_a_F_heap2_redo_10)) >> (uint(int32(2)) % 32))
	goto L343
L342:
	;
	v1680 = int32(0)
	goto L343
L343:
	;
	if base.Ui32(v1680&int32(_a_F_heap2_redo_11)) < base.Ui32(v1652) {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1671+v1652<<(uint(int32(2))%32))+20))
	if v1687&int32(_a_F_heap2_redo_8) != int32(_a_F_heap2_redo_17) {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	v1694 = v1671 + v1687&int32(_a_F_heap2_redo_6)
	v1695 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)))
	v1697 = v1695 & int32(_a_F_heap2_redo_18)
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)) = uint16(v1697)
	v1699 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1694)+18)))
	v1701 = v1699 & int32(-8193)
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+18)) = uint16(v1701)
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+18)) = uint16(v1701)
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)) = uint16(v1697)
	if v1703&int32(15) != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1708 = int32(1)
	v1727 = v1703<<(uint(v1708)%32)&int32(16) | (v1703<<(uint(int32(4))%32)&int32(64) | (v1703<<(uint(int32(6))%32)&int32(128) | v1703&v1708<<(uint(int32(12))%32))) | v1697
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)) = uint16(v1727)
	goto L348
L347:
	;
	goto L348
L348:
	;
	if v1703&int32(16) != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1732 = v1699 | int32(_a_F_heap2_redo_12)
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+18)) = uint16(v1732)
	goto L351
L350:
	;
	goto L351
L351:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1607)))
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+4)) = v1734
	*(*uint32)(unsafe.Add(mBase, uint32(v1671)+4)) = uint32(v1606)
	v1738 = int64(base.Ui64(v1606) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v1671))) = uint32(v1738)
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v21)+168))
	F_MarkBufferDirty(m, v1740)
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L11
	} else {
		goto L352
	}
L352:
	;
	goto L336
L353:
	;
	F_UnlockReleaseBuffer(m, v1749)
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L11
	} else {
		goto L354
	}
L354:
	;
	goto L5
L355:
	;
	v1782 = F_OpenTransientFile(m, v1774, int32(65))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L11
	} else {
		goto L359
	}
L356:
	;
	goto L5
L357:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L11
	} else {
		goto L402
	}
L358:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1920 = m.ExcPending
	if v1920 != 0 {
		goto L11
	} else {
		goto L398
	}
L359:
	;
	if int32(0) <= v1782 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1787 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1787))) = int32(167772198)
	v1790 = *(*int64)(unsafe.Add(mBase, uint32(v1759)+16))
	v1791 = F_ftruncate(m, v1782, v1790)
	mBase = m.M
	if v1791 != 0 {
		goto L358
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1903 = m.ExcPending
	if v1903 != 0 {
		goto L11
	} else {
		goto L394
	}
L363:
	;
	v1792 = int32(_a_F_heap2_redo_19)
	v1793 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	v1794 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1793))) = v1794
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1796)+64))
	v1798 = *(*int32)(unsafe.Add(mBase, uint32(v1759)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10])) = v1794
	v1803 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1803))) = int32(167772196)
	v1809 = v1798 * int32(36)
	v1810 = *(*int64)(unsafe.Add(mBase, uint32(v1759)+16))
	v1811 = F_pwrite(m, v1782, v1797+int32(40), v1809, v1810)
	mBase = m.M
	if v1811 != v1809 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10]))
	if v1814 == int32(0) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	goto L366
L366:
	;
	v1839 = int32(_a_F_heap2_redo_19)
	v1840 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	v1841 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1840))) = v1841
	v1844 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1844))) = int32(167772195)
	v1849 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[11])))
	if v1849 != int32(1) {
		v1863 = v1841
		goto L376
	} else {
		goto L377
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10])) = int32(51)
	goto L369
L368:
	;
	goto L369
L369:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L11
	} else {
		goto L370
	}
L370:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L11
	} else {
		goto L371
	}
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+48)) = v1756 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_20), v1756+int32(48))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L11
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_21), int32(1122), int32(_a_F_heap2_redo_22))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L11
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v1892))) = int32(0)
	v1895 = F_CloseTransientFile(m, v1782)
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L11
	} else {
		goto L392
	}
L375:
	;
	if v1863 == int32(0) {
		goto L374
	} else {
		goto L382
	}
L376:
	;
	goto L375
L377:
	;
	goto L378
L378:
	;
	v1854 = F_fsync(m, v1782)
	mBase = m.M
	if v1854 != int32(-1) {
		v1863 = v1854
		goto L376
	} else {
		goto L380
	}
L379:
	;
	v1863 = int32(-1)
	goto L376
L380:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, _c_F_heap2_redo[10]))
	if v1858 == int32(27) {
		goto L378
	} else {
		goto L381
	}
L381:
	;
	goto L379
L382:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap2_redo[12])))
	if v1869 != 0 {
		goto L384
	} else {
		goto L385
	}
L383:
	;
	v1872 = F_errstart(m, v1870, int32(0))
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L11
	} else {
		goto L387
	}
L384:
	;
	v1870 = int32(21)
	goto L386
L385:
	;
	v1870 = int32(23)
	goto L386
L386:
	;
	goto L383
L387:
	;
	if v1872 == int32(0) {
		goto L374
	} else {
		goto L388
	}
L388:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L11
	} else {
		goto L389
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+32)) = v1756 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_23), v1756+int32(32))
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L11
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_21), int32(1135), int32(_a_F_heap2_redo_22))
	mBase = m.M
	v1890 = m.ExcPending
	if v1890 != 0 {
		goto L11
	} else {
		goto L391
	}
L391:
	;
	goto L374
L392:
	;
	if v1895 != 0 {
		goto L357
	} else {
		goto L393
	}
L393:
	;
	m.G0 = v1756 + int32(1136)
	goto L356
L394:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L11
	} else {
		goto L395
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1756))) = v1756 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_24), v1756)
	mBase = m.M
	v1911 = m.ExcPending
	if v1911 != 0 {
		goto L11
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_21), int32(1094), int32(_a_F_heap2_redo_22))
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L11
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L398:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1922 = m.ExcPending
	if v1922 != 0 {
		goto L11
	} else {
		goto L399
	}
L399:
	;
	v1923 = *(*int64)(unsafe.Add(mBase, uint32(v1759)+16))
	*(*uint32)(unsafe.Add(mBase, uint32(v1756)+68)) = uint32(v1923)
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+64)) = v1756 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_25), v1756-int32(-64))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L11
	} else {
		goto L400
	}
L400:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_21), int32(1105), int32(_a_F_heap2_redo_22))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L11
	} else {
		goto L401
	}
L401:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L402:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v1943 = m.ExcPending
	if v1943 != 0 {
		goto L11
	} else {
		goto L403
	}
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1756)+16)) = v1756 + int32(112)
	F_errmsg(m, int32(_a_F_heap2_redo_26), v1756+int32(16))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L11
	} else {
		goto L404
	}
L404:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_21), int32(1141), int32(_a_F_heap2_redo_22))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L11
	} else {
		goto L405
	}
L405:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L406:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_27), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L11
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_28), int32(619), int32(_a_F_heap2_redo_29))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L11
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_30), int32(0))
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L11
	} else {
		goto L410
	}
L410:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_28), int32(645), int32(_a_F_heap2_redo_29))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L11
	} else {
		goto L411
	}
L411:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L412:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_31), int32(0))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L11
	} else {
		goto L413
	}
L413:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_28), int32(648), int32(_a_F_heap2_redo_29))
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L11
	} else {
		goto L414
	}
L414:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L415:
	;
	F_errmsg_internal(m, int32(_a_F_heap2_redo_32), int32(0))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L11
	} else {
		goto L416
	}
L416:
	;
	F_errfinish(m, int32(_a_F_heap2_redo_28), int32(1113), int32(_a_F_heap2_redo_33))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L11
	} else {
		goto L417
	}
L417:
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
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
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
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
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)))
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
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)))
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
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_LockBuffer(m, v310, int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L67
	}
L29:
	;
	v288 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v288
	v292 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v292
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)) = uint8(v292)
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
	F_LockBuffer(m, v276, int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
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
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v105 + v179&int32(_a_F_heapgettup_4)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v166)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v190)
	v193 = int32(base.Ui32(v190) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v193)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v188) >> (uint(int32(17)) % 32))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v220)+44))
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
	if v243 == int32(0) {
		goto L51
	} else {
		goto L63
	}
L63:
	;
	v250 = v221 - int32(1)
	if v250 != 0 {
		v220 = v220 + int32(48)
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
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+50)) = uint16(v166)
	goto L26
}
func F_hemdist_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	v5 = Fn13925(m, l0, l1, l2, int32(2))
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
	var v17 int32
	_ = v17
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
	var v95 int32
	_ = v95
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
	var v251 int64
	_ = v251
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v298 int64
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int64
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int64
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v367 int64
	_ = v367
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v376 int64
	_ = v376
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v381 int32
	_ = v381
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int64
	_ = v384
	var v388 int64
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int64
	_ = v399
	var v404 int64
	_ = v404
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int64
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
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
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int64
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v471 int64
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int64
	_ = v479
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int64
	_ = v489
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int64
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v510 int64
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v520 int64
	_ = v520
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v529 int64
	_ = v529
	var v530 int32
	_ = v530
	var v531 int64
	_ = v531
	var v532 int32
	_ = v532
	var v533 int64
	_ = v533
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v552 int64
	_ = v552
	var v560 int64
	_ = v560
	var v573 int64
	_ = v573
	v8 = int64(0)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 == int32(1) {
		if v9&int32(1) != 0 {
			return int32(0)
		} else {
			v17 = int32(3)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v17 < l2 {
				v251 = int64(0)
				if base.B2i32(v19 != (v19+int32(3))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v330 = v19
					v331 = l2
					v336 = v251
				} else {
					v261 = l2 - int32(4)
					v265 = int32(base.Ui32(v261)>>(uint(int32(2))%32)) + int32(1)
					v267 = v265 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v261) {
						v272 = v19
						v273 = l2
						v276 = int32(0)
						v278 = v251
						for {
							v279 = int32(16)
							v280 = v273 - v279
							v282 = v272 + v279
							v283 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
							v286 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
							v289 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
							v292 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
							v298 = base.I64_extend_i32_u(base.I32_popcnt(v283)) + (base.I64_extend_i32_u(base.I32_popcnt(v286)) + (base.I64_extend_i32_u(base.I32_popcnt(v289)) + (v278 + base.I64_extend_i32_u(base.I32_popcnt(v292)))))
							v300 = v276 + int32(4)
							if v300 != v265&int32(2147483644) {
								v272 = v282
								v273 = v280
								v276 = v300
								v278 = v298
								continue
							} else {
								break
							}
							break
						}
						if v267 == int32(0) {
							v330 = v282
							v331 = v280
							v336 = v298
						} else {
							v304 = v282
							v305 = v280
							v310 = v298
							v312 = v304
							v313 = v305
							v314 = int32(0)
							v318 = v310
							for {
								v319 = int32(4)
								v320 = v313 - v319
								v322 = v312 + v319
								v323 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
								v326 = v318 + base.I64_extend_i32_u(base.I32_popcnt(v323))
								v328 = v314 + int32(1)
								if v328 != v267 {
									v312 = v322
									v313 = v320
									v314 = v328
									v318 = v326
									continue
								} else {
									break
								}
								break
							}
							v330 = v322
							v331 = v320
							v336 = v326
						}
					} else {
						v304 = v19
						v305 = l2
						v310 = v251
						v312 = v304
						v313 = v305
						v314 = int32(0)
						v318 = v310
						for {
							v319 = int32(4)
							v320 = v313 - v319
							v322 = v312 + v319
							v323 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
							v326 = v318 + base.I64_extend_i32_u(base.I32_popcnt(v323))
							v328 = v314 + int32(1)
							if v328 != v267 {
								v312 = v322
								v313 = v320
								v314 = v328
								v318 = v326
								continue
							} else {
								break
							}
							break
						}
						v330 = v322
						v331 = v320
						v336 = v326
					}
				}
				if v331 == int32(0) {
					v399 = v336
				} else {
					v340 = v331 & int32(3)
					if v340 == int32(0) {
						v361 = v330
						v363 = v331
						v367 = v336
					} else {
						v344 = v330
						v346 = v331
						v348 = int32(0)
						v350 = v336
						for {
							v351 = int32(1)
							v352 = v344 + v351
							v354 = v346 - v351
							v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
							v356 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v355)+uint32(_c_F_hemdistcache_2[0]))))
							v357 = v350 + v356
							v359 = v348 + v351
							if v359 != v340 {
								v344 = v352
								v346 = v354
								v348 = v359
								v350 = v357
								continue
							} else {
								break
							}
							break
						}
						v361 = v352
						v363 = v354
						v367 = v357
					}
					if base.Ui32(v331) < base.Ui32(int32(4)) {
						v399 = v367
					} else {
						v370 = v361
						v372 = v363
						v376 = v367
						for {
							v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+3)))
							v378 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v377)+uint32(_c_F_hemdistcache_2[0]))))
							v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+2)))
							v380 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v379)+uint32(_c_F_hemdistcache_2[0]))))
							v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+1)))
							v382 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v381)+uint32(_c_F_hemdistcache_2[0]))))
							v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370))))
							v384 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v383)+uint32(_c_F_hemdistcache_2[0]))))
							v388 = v378 + (v380 + (v382 + (v376 + v384)))
							v389 = int32(4)
							v392 = v372 - v389
							if v392 != 0 {
								v370 = v370 + v389
								v372 = v392
								v376 = v388
								continue
							} else {
								break
							}
							break
						}
						v399 = v388
					}
				}
				v573 = v399
			} else {
				if l2 == int32(0) {
					v573 = v8
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
							v573 = v58
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
							v573 = v87
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
						v573 = v87
					}
				}
			}
			return l2<<(uint(v17)%32) + (base.I32_wrap_i64(v573) ^ int32(-1))
		}
	} else {
		if v9&int32(1) != 0 {
			v95 = int32(3)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v95 < l2 {
				v404 = int64(0)
				if base.B2i32(v97 != (v97+int32(3))&int32(-4))|base.B2i32(l2 < int32(4)) != 0 {
					v483 = v97
					v484 = l2
					v489 = v404
				} else {
					v414 = l2 - int32(4)
					v418 = int32(base.Ui32(v414)>>(uint(int32(2))%32)) + int32(1)
					v420 = v418 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v414) {
						v425 = v97
						v426 = l2
						v429 = int32(0)
						v431 = v404
						for {
							v432 = int32(16)
							v433 = v426 - v432
							v435 = v425 + v432
							v436 = *(*int32)(unsafe.Add(mBase, uint32(v425)+12))
							v439 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
							v442 = *(*int32)(unsafe.Add(mBase, uint32(v425)+4))
							v445 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
							v451 = base.I64_extend_i32_u(base.I32_popcnt(v436)) + (base.I64_extend_i32_u(base.I32_popcnt(v439)) + (base.I64_extend_i32_u(base.I32_popcnt(v442)) + (v431 + base.I64_extend_i32_u(base.I32_popcnt(v445)))))
							v453 = v429 + int32(4)
							if v453 != v418&int32(2147483644) {
								v425 = v435
								v426 = v433
								v429 = v453
								v431 = v451
								continue
							} else {
								break
							}
							break
						}
						if v420 == int32(0) {
							v483 = v435
							v484 = v433
							v489 = v451
						} else {
							v457 = v435
							v458 = v433
							v463 = v451
							v465 = v457
							v466 = v458
							v467 = int32(0)
							v471 = v463
							for {
								v472 = int32(4)
								v473 = v466 - v472
								v475 = v465 + v472
								v476 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
								v479 = v471 + base.I64_extend_i32_u(base.I32_popcnt(v476))
								v481 = v467 + int32(1)
								if v481 != v420 {
									v465 = v475
									v466 = v473
									v467 = v481
									v471 = v479
									continue
								} else {
									break
								}
								break
							}
							v483 = v475
							v484 = v473
							v489 = v479
						}
					} else {
						v457 = v97
						v458 = l2
						v463 = v404
						v465 = v457
						v466 = v458
						v467 = int32(0)
						v471 = v463
						for {
							v472 = int32(4)
							v473 = v466 - v472
							v475 = v465 + v472
							v476 = *(*int32)(unsafe.Add(mBase, uint32(v465)))
							v479 = v471 + base.I64_extend_i32_u(base.I32_popcnt(v476))
							v481 = v467 + int32(1)
							if v481 != v420 {
								v465 = v475
								v466 = v473
								v467 = v481
								v471 = v479
								continue
							} else {
								break
							}
							break
						}
						v483 = v475
						v484 = v473
						v489 = v479
					}
				}
				if v484 == int32(0) {
					v552 = v489
				} else {
					v493 = v484 & int32(3)
					if v493 == int32(0) {
						v514 = v483
						v516 = v484
						v520 = v489
					} else {
						v497 = v483
						v499 = v484
						v501 = int32(0)
						v503 = v489
						for {
							v504 = int32(1)
							v505 = v497 + v504
							v507 = v499 - v504
							v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
							v509 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v508)+uint32(_c_F_hemdistcache_2[0]))))
							v510 = v503 + v509
							v512 = v501 + v504
							if v512 != v493 {
								v497 = v505
								v499 = v507
								v501 = v512
								v503 = v510
								continue
							} else {
								break
							}
							break
						}
						v514 = v505
						v516 = v507
						v520 = v510
					}
					if base.Ui32(v484) < base.Ui32(int32(4)) {
						v552 = v520
					} else {
						v523 = v514
						v525 = v516
						v529 = v520
						for {
							v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+3)))
							v531 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v530)+uint32(_c_F_hemdistcache_2[0]))))
							v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+2)))
							v533 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v532)+uint32(_c_F_hemdistcache_2[0]))))
							v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523)+1)))
							v535 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v534)+uint32(_c_F_hemdistcache_2[0]))))
							v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v523))))
							v537 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v536)+uint32(_c_F_hemdistcache_2[0]))))
							v541 = v531 + (v533 + (v535 + (v529 + v537)))
							v542 = int32(4)
							v545 = v525 - v542
							if v545 != 0 {
								v523 = v523 + v542
								v525 = v545
								v529 = v541
								continue
							} else {
								break
							}
							break
						}
						v552 = v541
					}
				}
				v560 = v552
			} else {
				if l2 == int32(0) {
					v560 = v8
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
							v560 = v136
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
							v560 = v165
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
						v560 = v165
					}
				}
			}
			return l2<<(uint(v95)%32) + (base.I32_wrap_i64(v560) ^ int32(-1))
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
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hs_contained_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
