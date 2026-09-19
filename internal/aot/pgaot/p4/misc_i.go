package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeTimeouts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	v1 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[0])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[1])) = v1
	v11 = v1
	for {
		v13 = int32(40)
		v14 = v11 * v13
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[3]))) = v11
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[4]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[5]))) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[6]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[7]))) = v15
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v15)
		v29 = v11 | int32(1)
		v31 = v29 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[3]))) = v29
		*(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[4]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[5]))) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[6]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[7]))) = v15
		*(*uint8)(unsafe.Add(mBase, uint32(v31)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v15)
		v46 = v11 | int32(2)
		v48 = v46 * v13
		*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[3]))) = v46
		*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[4]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[5]))) = v15
		*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[6]))) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[7]))) = v15
		*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v15)
		if v11 != int32(20) {
			v65 = v11 | int32(3)
			v67 = v65 * int32(40)
			v68 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[2]))) = uint8(v68)
			*(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[3]))) = v65
			v71 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[4]))) = v71
			*(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[5]))) = v68
			*(*int64)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[6]))) = v71
			*(*int32)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[7]))) = v68
			*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_c_F_InitializeTimeouts[8]))) = uint8(v68)
			v11 = v11 + int32(4)
			continue
		} else {
			break
		}
		break
	}
	v84 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeTimeouts[9])) = uint8(v84)
	v87 = int32(1769)
	v89 = m.G0
	v91 = v89 - int32(32)
	m.G0 = v91
	switch int32(1771) {
	case 0, 2:
		v101 = v87
	default:
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeTimeouts[10])) = v87
		v101 = int32(_a_F_InitializeTimeouts_0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v101
	F_sigemptyset(m, v91+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v91)+24)) = int32(268435456)
	v115 = F___sigaction(m, int32(14), v91+int32(12), int32(0))
	mBase = m.M
	m.G0 = v91 + int32(32)
	return
}
func F_InputFunctionCall(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn13834(m, l0, l1, l2, l3, int32(_a_F_InputFunctionCall_0), int32(1560), int32(_a_F_InputFunctionCall_1), int32(1554), int32(_a_F_InputFunctionCall_2))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F___isspace(m *base.Module, l0 int32) int32 {
	return base.B2i32(l0 == int32(32)) | base.B2i32(base.Ui32(l0-int32(9)) < base.Ui32(int32(5)))
}
func F__intbig_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F__intbig_in_0), int32(36), int32(_a_F__intbig_in_1), int32(_a_F__intbig_in_2), int32(_a_F__intbig_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_i2tod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = F_Float8GetDatum(m, base.F64_convert_i32_s(v2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_i4tod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_Float8GetDatum(m, base.F64_convert_i32_s(v2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_i8tof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return base.I32_reinterpret_f32(base.F32_convert_i64_s(v3))
}
func F_inclusion_get_procinfo(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+16))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+112)))
	if v15 != 0 {
		v54 = int32(0)
		m.G0 = v8 + int32(16)
		return v54
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
		if v16 != 0 {
			v54 = v14
			m.G0 = v8 + int32(16)
			return v54
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v18 = base.I32_extend16_s(l1)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+216))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v20+v22*(v18-int32(1))<<(uint(int32(2))%32)+int32(44)-int32(4))))
			if v34 != 0 {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v37 = F_index_getprocinfo(m, v35, v18, int32(11))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v37)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v42
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v44
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v46
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
					*(*int64)(unsafe.Add(mBase, uint32(v14))) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v41
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
					v54 = v14
					m.G0 = v8 + int32(16)
					return v54
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(117833860))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_inclusion_get_procinfo_0), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(11)
							F_errdetail_internal(m, int32(_a_F_inclusion_get_procinfo_1), v8)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_inclusion_get_procinfo_2), int32(577), int32(_a_F_inclusion_get_procinfo_3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
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
		}
	}
}
func F_infobits_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	F_appendStringInfo(m, l0, int32(_a_F_infobits_desc_0), v7)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l1&int32(1) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	if l1&int32(2) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_2))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if l1&int32(4) != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l1&int32(8) != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l1&int32(16) != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_infobits_desc_5))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v39-int32(1)))))
	if v43 == int32(32) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v47 = v39 - int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
	v50 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v38+v47))) = uint8(v50)
	goto L25
L24:
	;
	goto L25
L25:
	;
	F_appendStringInfoChar(m, l0, int32(93))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_initTrie(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v504 int32
	_ = v504
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v525 int32
	_ = v525
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v591 int32
	_ = v591
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v704 int32
	_ = v704
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v753 int32
	_ = v753
	var v754 int64
	_ = v754
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	v2 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(256)
	m.G0 = v25
	v28 = l0
	v29 = v25
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v37 = int32(-1)
	v45 = v2
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v37 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v753 = int32(m.ExcTag)
	v754 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v753 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L7:
	;
	v128 = v30
	v129 = v31
	v132 = v34
	v135 = v125
	v143 = v45
	goto L20
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+228)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_initTrie[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	v63 = F_get_tsearch_config_filename(m, v28, int32(_a_F_initTrie_0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v121 = v32
	v122 = v33
	v123 = v35
	v125 = int32(1)
	goto L7
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	v73 = F_tsearch_readline_begin(m, v29+int32(184), v63)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L12
	}
L12:
	;
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_initTrie[1]))
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_initTrie[2]))
	v121 = v76
	v122 = v78
	v123 = v60
	v125 = int32(0)
	goto L7
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	F_errcode(m, int32(22))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v63
	F_errmsg(m, int32(_a_F_initTrie_1), v29)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v60
	F_errfinish(m, int32(_a_F_initTrie_2), int32(109), int32(_a_F_initTrie_3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	if v135 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v150 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+183)) = uint8(v150)
	goto L25
L23:
	;
	goto L24
L24:
	;
	if v143 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	v153 = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v29 + int32(12)
	goto L28
L26:
	;
	v135 = int32(1)
	v143 = int32(0)
	goto L20
L28:
	;
	goto L26
L29:
	;
	v128 = v680
	v129 = v681
	v132 = v684
	v135 = int32(0)
	goto L20
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[0])) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_pg_re_throw(m)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L134
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[1])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[2])) = v122
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+183)))
	if v704 != 0 {
		goto L29
	} else {
		goto L132
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v132
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[2])) = v29 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v175 = F_tsearch_readline(m, v29+int32(184))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[1])) = v121
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[2])) = v122
	v655 = int32(_a_F_initTrie_4)
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_initTrie[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_initTrie[0])) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v665 = F_CopyErrorData(m)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L129
	}
L35:
	;
	if v175 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v179 = v128
	v180 = v129
	v183 = v132
	v195 = v175
	goto L39
L37:
	;
	v629 = v128
	v630 = v129
	v633 = v132
	goto L38
L38:
	;
	v649 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+183)) = uint8(v649)
	v680 = v629
	v681 = v630
	v684 = v633
	goto L31
L39:
	;
	v199 = int32(0)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v205 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v629 = v581
	v630 = v582
	v633 = v625
	goto L38
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_pfree(m, v591)
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L125
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v549
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v29)+228))
	v576 = F_placeChar(m, v575, v289, v291, v559, v555)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L124
	}
L43:
	;
	switch v450 + int32(2) {
	case 0:
		goto L112
	case 1:
		goto L113
	default:
		v581 = v441
		v582 = v442
		v591 = v451
		goto L41
	}
L44:
	;
	v338 = int32(0)
	v339 = base.B2i32(v290 <= v338)
	if (v339|(v293^int32(-1)))&int32(1) == v338 {
		goto L89
	} else {
		goto L90
	}
L45:
	;
	v214 = v199
	v215 = v195
	v216 = v199
	v217 = v199
	v221 = v199
	v222 = v199
	v225 = v199
	goto L48
L46:
	;
	v315 = v199
	v318 = v199
	v323 = v199
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v335 = F_palloc(m, v315)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L85
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v234 = F_pg_mblen_cstr(m, v215)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L50
	}
L49:
	;
	v300 = base.B2i32(base.Ui32(v290-int32(1)) < base.Ui32(int32(2)))
	if base.Ui32(v290-int32(1)) < base.Ui32(int32(2)) {
		goto L78
	} else {
		goto L79
	}
L50:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v236-int32(9)))&base.B2i32(v236 != int32(32)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v294 = v288 + v234
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v295 != 0 {
		v214 = v287
		v215 = v294
		v216 = v289
		v217 = v290
		v221 = v291
		v222 = v292
		v225 = v293
		goto L48
	} else {
		goto L77
	}
L52:
	;
	v287 = v282
	v288 = v283
	v289 = v216
	v290 = v285
	v291 = v221
	v292 = v222
	v293 = v225
	goto L51
L53:
	;
	if v217 == int32(3) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v259 = int32(1)
	switch v217 {
	case 0:
		v287 = v214
		v288 = v215
		v289 = v215
		v290 = v259
		v291 = v234
		v292 = v222
		v293 = v225
		goto L51
	case 1:
		goto L69
	case 2:
		goto L68
	case 3:
		goto L67
	case 4:
		goto L66
	default:
		goto L65
	}
L56:
	;
	v251 = int32(5)
	goto L58
L57:
	;
	v251 = v217
	goto L58
L58:
	;
	if v217 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v254 = int32(2)
	goto L61
L60:
	;
	v254 = v251
	goto L61
L61:
	;
	if v254 == int32(4) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v257 = v234
	goto L64
L63:
	;
	v257 = int32(0)
	goto L64
L64:
	;
	v282 = v214 + v257
	v283 = v215
	v285 = v254
	goto L52
L65:
	;
	v282 = v214
	v283 = v215
	v285 = int32(-1)
	goto L52
L66:
	;
	v269 = v214 + v234
	v270 = int32(4)
	if v236 != int32(34) {
		v282 = v269
		v283 = v215
		v285 = v270
		goto L52
	} else {
		goto L73
	}
L67:
	;
	v282 = v214 + v234
	v283 = v215
	v285 = int32(3)
	goto L52
L68:
	;
	v264 = base.B2i32(v236 == int32(34))
	if v236 == int32(34) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v287 = v214
	v288 = v215
	v289 = v216
	v290 = v259
	v291 = v234 + v221
	v292 = v222
	v293 = v225
	goto L51
L70:
	;
	v265 = int32(4)
	goto L72
L71:
	;
	v265 = int32(3)
	goto L72
L72:
	;
	v287 = v234
	v288 = v215
	v289 = v216
	v290 = v265
	v291 = v221
	v292 = v215
	v293 = v264 | v225
	goto L51
L73:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v273 != int32(34) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v287 = v269
	v288 = v215
	v289 = v216
	v290 = int32(5)
	v291 = v221
	v292 = v222
	v293 = v225
	goto L51
L75:
	;
	goto L76
L76:
	;
	v277 = int32(1)
	v282 = v269 + v277
	v283 = v215 + v277
	v285 = v270
	goto L52
L77:
	;
	goto L49
L78:
	;
	v301 = int32(_a_F_initTrie_5)
	goto L80
L79:
	;
	v301 = v292
	goto L80
L80:
	;
	if base.Ui32(v290-int32(1)) < base.Ui32(int32(2)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v303 = int32(0)
	goto L83
L82:
	;
	v303 = v287
	goto L83
L83:
	;
	if v290 != int32(4) {
		goto L44
	} else {
		goto L84
	}
L84:
	;
	v315 = v303
	v318 = int32(-2)
	v323 = v301
	goto L47
L85:
	;
	if v315 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	base.MemoryCopy(m, v335, v323, v315)
	goto L88
L87:
	;
	goto L88
L88:
	;
	v441 = v179
	v442 = v335
	v450 = v318
	v451 = v335
	goto L43
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v355 = F_palloc(m, v303-int32(2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v434 = F_palloc(m, v303)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L107
	}
L92:
	;
	if v303 < int32(3) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v290 <= v338 {
		v441 = v179
		v442 = v180
		v450 = v290
		v451 = v355
		goto L43
	} else {
		goto L106
	}
L94:
	;
	v414 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v360 = int32(1)
	v372 = int32(0)
	v373 = v360
	goto L97
L97:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v301))))
	*(*uint8)(unsafe.Add(mBase, uint32(v372+v355))) = uint8(v388)
	v391 = v372 + int32(1)
	if v388 == int32(34) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v414 = v391
	goto L93
L99:
	;
	v395 = v373 + int32(1)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v301))))
	if v397 == int32(34) {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v402 = v373
	goto L101
L101:
	;
	v404 = v402 + int32(1)
	if v404 < v303-v360 {
		v372 = v391
		v373 = v404
		goto L97
	} else {
		goto L105
	}
L102:
	;
	v400 = v395
	goto L104
L103:
	;
	v400 = v373
	goto L104
L104:
	;
	v402 = v400
	goto L101
L105:
	;
	goto L98
L106:
	;
	v549 = v179
	v555 = v414
	v559 = v355
	goto L42
L107:
	;
	if v303 != 0 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	base.MemoryCopy(m, v434, v301, v303)
	goto L110
L109:
	;
	goto L110
L110:
	;
	if int32(0) < v290 {
		v549 = v434
		v555 = v303
		v559 = v434
		goto L42
	} else {
		goto L111
	}
L111:
	;
	v441 = v434
	v442 = v180
	v450 = v290
	v451 = v434
	goto L43
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v513 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L119
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v471 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L114
	}
L114:
	;
	if v471 == int32(0) {
		v581 = v441
		v582 = v442
		v591 = v451
		goto L41
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errcode(m, int32(22))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errmsg(m, int32(_a_F_initTrie_6), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errfinish(m, int32(_a_F_initTrie_2), int32(267), int32(_a_F_initTrie_3))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L118
	}
L118:
	;
	v581 = v441
	v582 = v442
	v591 = v451
	goto L41
L119:
	;
	if v513 == int32(0) {
		v581 = v441
		v582 = v442
		v591 = v451
		goto L41
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errcode(m, int32(22))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errmsg(m, int32(_a_F_initTrie_7), int32(0))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v441
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_errfinish(m, int32(_a_F_initTrie_2), int32(271), int32(_a_F_initTrie_3))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L123
	}
L123:
	;
	v581 = v441
	v582 = v442
	v591 = v451
	goto L41
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+228)) = v576
	v581 = v549
	v582 = v180
	v591 = v559
	goto L41
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_pfree(m, v195)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	v625 = F_tsearch_readline(m, v29+int32(184))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L127
	}
L127:
	;
	if v625 != 0 {
		v179 = v581
		v180 = v582
		v183 = v625
		v195 = v625
		goto L39
	} else {
		goto L128
	}
L128:
	;
	goto L40
L129:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v665)+28))
	if v667 != int32(84017282) {
		goto L30
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_FlushErrorState(m)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L131
	}
L131:
	;
	v680 = v128
	v681 = v129
	v684 = v132
	goto L31
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v680
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v684
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v29)+244)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v29)+252)) = v123
	F_tsearch_readline_end(m, v29+int32(184))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		v731 = v28
		v732 = v29
		goto L6
	} else {
		goto L133
	}
L133:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v29)+228))
	m.G0 = v29 + int32(256)
	return v715
L134:
	;
	goto L3
L135:
	;
	v758 = int32(v754)
	m.G0 = v732
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v758)+4))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v758)))
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v761)))
	if v732+int32(12) == v764 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	m.ExcPending = 1
	goto L144
L137:
	;
	if v768 != 0 {
		goto L141
	} else {
		goto L142
	}
L138:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v761)+4))
	v768 = v766
	goto L140
L139:
	;
	v768 = int32(0)
	goto L140
L140:
	;
	goto L137
L141:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v732)+252))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v732)+248))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v732)+244))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v732)+240))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v732)+236))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v732)+232))
	v28 = v731
	v29 = v732
	v30 = v773
	v31 = v772
	v32 = v771
	v33 = v770
	v34 = v774
	v35 = v769
	v37 = v768
	v45 = v760
	goto L1
L142:
	;
	goto L143
L143:
	;
	F___wasm_longjmp(m, v761, v760)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	return int32(0)
L145:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_init_MultiFuncCall(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_init_MultiFuncCall_0), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(143), int32(_a_F_init_MultiFuncCall_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
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
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v8 != int32(383) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_init_MultiFuncCall_0), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(143), int32(_a_F_init_MultiFuncCall_2))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v12 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
				v51 = F_AllocSetContextCreateInternal(m, v46, int32(_a_F_init_MultiFuncCall_3), int32(0), int32(1024), int32(_a_F_init_MultiFuncCall_4))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v54 = F_MemoryContextAllocZero(m, v51, int32(32))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v54))) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v56
						*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v56
						*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = v51
						v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v54
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_RegisterExprContextCallback(m, v67, int32(1613), v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v54
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_init_MultiFuncCall_5), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_init_MultiFuncCall_1), int32(193), int32(_a_F_init_MultiFuncCall_2))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
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
}
func F_init_work(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
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
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v902 int32
	_ = v902
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1062 int32
	_ = v1062
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1141 int32
	_ = v1141
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1187 int32
	_ = v1187
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1225 int32
	_ = v1225
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1240 int32
	_ = v1240
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1350 int32
	_ = v1350
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1637 int32
	_ = v1637
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1657 int32
	_ = v1657
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1674 int32
	_ = v1674
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1689 int32
	_ = v1689
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1704 int32
	_ = v1704
	var v1707 int32
	_ = v1707
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1746 int32
	_ = v1746
	var v1751 int32
	_ = v1751
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1776 int32
	_ = v1776
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1811 int32
	_ = v1811
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1836 int32
	_ = v1836
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1872 int32
	_ = v1872
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1907 int32
	_ = v1907
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1928 int32
	_ = v1928
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1935 int32
	_ = v1935
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1945 int32
	_ = v1945
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1963 int32
	_ = v1963
	var v1972 int32
	_ = v1972
	var v1980 int32
	_ = v1980
	var v1990 int32
	_ = v1990
	var v1994 int32
	_ = v1994
	var v1998 int32
	_ = v1998
	v16 = F_pgp_init(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(l3)+32)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = int32(-1)
	v30 = int32(0)
	if base.B2i32(l2 == v30)|v16 == v30 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = int32(1)
	v36 = l2 + v35
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v39 = v37 & v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v37 == v35 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v1980 = v16
	goto L5
L5:
	;
	if v1980 == int32(0) {
		goto L597
	} else {
		goto L598
	}
L6:
	;
	v184 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v174+v179))) = uint8(v184)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v188 == v184 {
		v1963 = v184
		goto L41
	} else {
		goto L42
	}
L7:
	;
	if v39 != 0 {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v74 = F_palloc(m, v71+int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L20
	}
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v71 = int32(base.Ui32(v65)>>(uint(int32(2))%32)) - int32(4)
	goto L8
L10:
	;
	v63 = F_palloc(m, int32(5))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L19
	}
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if base.Ui32((v43-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v39 == int32(0) {
		goto L9
	} else {
		goto L18
	}
L14:
	;
	if v43 == int32(18) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = int32(16)
	goto L17
L16:
	;
	v54 = int32(0)
	goto L17
L17:
	;
	v71 = v54
	goto L8
L18:
	;
	v57 = int32(1)
	v71 = int32(base.Ui32(v37)>>(uint(v57)%32)) - v57
	goto L8
L19:
	;
	v78 = int32(4)
	v79 = v63
	goto L7
L20:
	;
	if v71 <= int32(0) {
		v174 = v71
		v179 = v74
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v78 = v71
	v79 = v74
	goto L7
L22:
	;
	v82 = v36
	goto L24
L23:
	;
	v82 = l2 + int32(4)
	goto L24
L24:
	;
	v83 = int32(0)
	if v78 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v93 = v83
	v97 = int32(0)
	goto L28
L26:
	;
	v143 = v83
	goto L27
L27:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143+v82))))
	if base.Ui32((v158-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93+v82))))
	if base.Ui32((v108-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v78&int32(1) == int32(0) {
		v174 = v78
		v179 = v79
		goto L6
	} else {
		goto L37
	}
L30:
	;
	v117 = v108 | int32(32)
	goto L32
L31:
	;
	v117 = v108
	goto L32
L32:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v93+v79))) = uint8(v117)
	v120 = v93 | int32(1)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82+v120))))
	if base.Ui32((v123-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v132 = v123 | int32(32)
	goto L35
L34:
	;
	v132 = v123
	goto L35
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v79+v120))) = uint8(v132)
	v134 = int32(2)
	v135 = v93 + v134
	v137 = v97 + v134
	if v137 != v78&int32(2147483646) {
		v93 = v135
		v97 = v137
		goto L28
	} else {
		goto L36
	}
L36:
	;
	goto L29
L37:
	;
	v143 = v135
	goto L27
L38:
	;
	v167 = v158 | int32(32)
	goto L40
L39:
	;
	v167 = v158
	goto L40
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v143+v79))) = uint8(v167)
	v174 = v78
	v179 = v79
	goto L6
L41:
	;
	F_pfree(m, v179)
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L596
	}
L42:
	;
	v196 = v179
	v197 = v188
	goto L43
L43:
	;
	switch v197 - int32(32) {
	case 0:
		goto L48
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28:
		goto L47
	case 12, 29:
		goto L46
	default:
		goto L49
	}
L44:
	;
	v1963 = v1954
	goto L41
L45:
	;
	v253 = int32(-13)
	v256 = v239
	goto L56
L46:
	;
	v239 = v196 + int32(1)
	goto L45
L47:
	;
	v217 = v196
	v221 = v197
	goto L51
L48:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	v196 = v196 + int32(1)
	v197 = v214
	goto L43
L49:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v197-int32(9)) {
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	switch v221 {
	case 0, 9, 10, 32, 44:
		v239 = v217
		goto L45
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43:
		goto L53
	default:
		goto L54
	}
L53:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	v217 = v217 + int32(1)
	v221 = v232
	goto L51
L54:
	;
	if v221 == int32(61) {
		v239 = v217
		goto L45
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	if base.B2i32(base.Ui32(v269-int32(9)) < base.Ui32(int32(2)))|base.B2i32(v269 == int32(32)) != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v345 = v283 + v341
	v351 = v345
	goto L73
L58:
	;
	v256 = v256 + int32(1)
	goto L56
L59:
	;
	if v269 != int32(61) {
		v1963 = v253
		goto L41
	} else {
		goto L62
	}
L60:
	;
	goto L57
L61:
	;
	goto L60
L62:
	;
	v283 = v256
	goto L63
L63:
	;
	v296 = int32(1)
	v298 = v283 + v296
	v299 = int32(2)
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if base.Ui32(v300-int32(9)) < base.Ui32(v299) {
		v283 = v298
		goto L63
	} else {
		goto L65
	}
L64:
	;
	v312 = v300
	v313 = v296
	goto L67
L65:
	;
	switch v300 - int32(32) {
	case 0:
		v283 = v298
		goto L63
	default:
		goto L66
	case 12, 29:
		v341 = v299
		goto L61
	}
L66:
	;
	goto L64
L67:
	;
	v323 = v312 & int32(255)
	switch v323 {
	case 0, 9, 10, 32, 44:
		v341 = v313
		goto L61
	case 1, 2, 3, 4, 5, 6, 7, 8, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43:
		goto L69
	default:
		goto L70
	}
L69:
	;
	v327 = v313 + int32(1)
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283+v327))))
	v312 = v329
	v313 = v327
	goto L67
L70:
	;
	if v323 != int32(61) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v341 = v313
	goto L61
L72:
	;
	v369 = int32(0)
	if base.B2i32(v197 == v369)|base.B2i32(v300 == v369)|base.B2i32(v341 == int32(1)) != 0 {
		v1963 = v253
		goto L41
	} else {
		goto L78
	}
L73:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	switch v361 {
	case 0:
		v368 = v351
		goto L72
	default:
		goto L75
	case 9, 10, 32:
		goto L76
	}
L74:
	;
	if v361 != int32(44) {
		v1963 = v253
		goto L41
	} else {
		goto L77
	}
L75:
	;
	goto L74
L76:
	;
	v351 = v351 + int32(1)
	goto L73
L77:
	;
	v368 = v351 + int32(1)
	goto L72
L78:
	;
	v378 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v196+(v239-v196)))) = uint8(v378)
	*(*uint8)(unsafe.Add(mBase, uint32(v345))) = uint8(v378)
	v382 = int32(_a_F_init_work_0)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[0])))
	if base.B2i32(v385 == v378)|base.B2i32(v385 != v388) != 0 {
		v406 = v385
		v407 = v388
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v1955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368))))
	if v1955 != 0 {
		v196 = v368
		v197 = v1955
		goto L43
	} else {
		goto L595
	}
L80:
	;
	v1155 = int32(_a_F_init_work_1)
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[1])))
	if base.B2i32(v1158 == int32(0))|base.B2i32(v1158 != v1161) != 0 {
		v1179 = v1158
		v1180 = v1161
		goto L352
	} else {
		goto L353
	}
L81:
	;
	if int32(0) <= v1152 {
		v1954 = v1152
		goto L79
	} else {
		goto L350
	}
L82:
	;
	if v406-v407 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	goto L82
L84:
	;
	v391 = v196
	v392 = v382
	goto L85
L85:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v392)+1)))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391)+1)))
	if v396 == int32(0) {
		v406 = v396
		v407 = v395
		goto L83
	} else {
		goto L87
	}
L86:
	;
	v406 = v396
	v407 = v395
	goto L83
L87:
	;
	v399 = int32(1)
	if v396 == v395 {
		v391 = v391 + v399
		v392 = v392 + v399
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v411 = F_pgp_get_cipher_code(m, v298)
	mBase = m.M
	if v411 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	goto L91
L91:
	;
	v417 = int32(_a_F_init_work_2)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[2])))
	if base.B2i32(v420 == int32(0))|base.B2i32(v420 != v423) != 0 {
		v441 = v420
		v442 = v423
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v1152 = v416
	goto L81
L93:
	;
	v416 = v411
	goto L92
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v411
	v416 = int32(0)
	goto L92
L96:
	;
	if v441-v442 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	goto L96
L98:
	;
	v426 = v196
	v427 = v417
	goto L99
L99:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v427)+1)))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426)+1)))
	if v431 == int32(0) {
		v441 = v431
		v442 = v430
		goto L97
	} else {
		goto L101
	}
L100:
	;
	v441 = v431
	v442 = v430
	goto L97
L101:
	;
	v434 = int32(1)
	if v431 == v430 {
		v426 = v426 + v434
		v427 = v427 + v434
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v449 = v298
	goto L107
L104:
	;
	goto L105
L105:
	;
	v498 = int32(_a_F_init_work_3)
	v501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v504 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[3])))
	if base.B2i32(v501 == int32(0))|base.B2i32(v501 != v504) != 0 {
		v522 = v501
		v523 = v504
		goto L124
	} else {
		goto L125
	}
L106:
	;
	v494 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+72)) = base.B2i32(v493 != v494)
	goto L122
L107:
	;
	v454 = v449 + int32(1)
	v455 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449))))
	v456 = F___isspace(m, v455)
	mBase = m.M
	if v456 != 0 {
		v449 = v454
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v457 = int32(1)
	switch v455&int32(255) - int32(43) {
	case 0:
		v463 = v457
		goto L111
	default:
		v465 = v455
		v466 = v449
		v467 = v457
		goto L110
	case 2:
		goto L112
	}
L109:
	;
	goto L108
L110:
	;
	v468 = int32(0)
	v470 = v465 - int32(48)
	if base.Ui32(v470) <= base.Ui32(int32(9)) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	v464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v454))))
	v465 = v464
	v466 = v454
	v467 = v463
	goto L110
L112:
	;
	v463 = int32(0)
	goto L111
L113:
	;
	v473 = v468
	v474 = v470
	v475 = v466
	goto L116
L114:
	;
	v487 = v468
	goto L115
L115:
	;
	if v467 != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v477 = int32(10)
	v479 = v473*v477 - v474
	v480 = int32(*(*int8)(unsafe.Add(mBase, uint32(v475)+1)))
	v484 = v480 - int32(48)
	if base.Ui32(v484) < base.Ui32(v477) {
		v473 = v479
		v474 = v484
		v475 = v475 + int32(1)
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v487 = v479
	goto L115
L118:
	;
	goto L117
L119:
	;
	v493 = int32(0) - v487
	goto L121
L120:
	;
	v493 = v487
	goto L121
L121:
	;
	goto L106
L122:
	;
	v1152 = v494
	goto L81
L123:
	;
	if v522-v523 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	goto L123
L125:
	;
	v507 = v196
	v508 = v498
	goto L126
L126:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507)+1)))
	if v512 == int32(0) {
		v522 = v512
		v523 = v511
		goto L124
	} else {
		goto L128
	}
L127:
	;
	v522 = v512
	v523 = v511
	goto L124
L128:
	;
	v515 = int32(1)
	if v512 == v511 {
		v507 = v507 + v515
		v508 = v508 + v515
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v530 = v298
	goto L134
L131:
	;
	goto L132
L132:
	;
	v579 = int32(_a_F_init_work_4)
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[4])))
	if base.B2i32(v582 == int32(0))|base.B2i32(v582 != v585) != 0 {
		v603 = v582
		v604 = v585
		goto L151
	} else {
		goto L152
	}
L133:
	;
	v575 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+76)) = base.B2i32(v574 != v575)
	goto L149
L134:
	;
	v535 = v530 + int32(1)
	v536 = int32(*(*int8)(unsafe.Add(mBase, uint32(v530))))
	v537 = F___isspace(m, v536)
	mBase = m.M
	if v537 != 0 {
		v530 = v535
		goto L134
	} else {
		goto L136
	}
L135:
	;
	v538 = int32(1)
	switch v536&int32(255) - int32(43) {
	case 0:
		v544 = v538
		goto L138
	default:
		v546 = v536
		v547 = v530
		v548 = v538
		goto L137
	case 2:
		goto L139
	}
L136:
	;
	goto L135
L137:
	;
	v549 = int32(0)
	v551 = v546 - int32(48)
	if base.Ui32(v551) <= base.Ui32(int32(9)) {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v545 = int32(*(*int8)(unsafe.Add(mBase, uint32(v535))))
	v546 = v545
	v547 = v535
	v548 = v544
	goto L137
L139:
	;
	v544 = int32(0)
	goto L138
L140:
	;
	v554 = v549
	v555 = v551
	v556 = v547
	goto L143
L141:
	;
	v568 = v549
	goto L142
L142:
	;
	if v548 != 0 {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v558 = int32(10)
	v560 = v554*v558 - v555
	v561 = int32(*(*int8)(unsafe.Add(mBase, uint32(v556)+1)))
	v565 = v561 - int32(48)
	if base.Ui32(v565) < base.Ui32(v558) {
		v554 = v560
		v555 = v565
		v556 = v556 + int32(1)
		goto L143
	} else {
		goto L145
	}
L144:
	;
	v568 = v560
	goto L142
L145:
	;
	goto L144
L146:
	;
	v574 = int32(0) - v568
	goto L148
L147:
	;
	v574 = v568
	goto L148
L148:
	;
	goto L133
L149:
	;
	v1152 = v575
	goto L81
L150:
	;
	if v603-v604 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L151:
	;
	goto L150
L152:
	;
	v588 = v196
	v589 = v579
	goto L153
L153:
	;
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588)+1)))
	if v593 == int32(0) {
		v603 = v593
		v604 = v592
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v603 = v593
	v604 = v592
	goto L151
L155:
	;
	v596 = int32(1)
	if v593 == v592 {
		v588 = v588 + v596
		v589 = v589 + v596
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v611 = v298
	goto L161
L158:
	;
	goto L159
L159:
	;
	v667 = int32(_a_F_init_work_5)
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v673 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[5])))
	if base.B2i32(v670 == int32(0))|base.B2i32(v670 != v673) != 0 {
		v691 = v670
		v692 = v673
		goto L181
	} else {
		goto L182
	}
L160:
	;
	if base.B2i32(v655 == int32(2))|base.B2i32(base.Ui32(int32(3)) < base.Ui32(v655)) != 0 {
		goto L177
	} else {
		goto L178
	}
L161:
	;
	v616 = v611 + int32(1)
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v611))))
	v618 = F___isspace(m, v617)
	mBase = m.M
	if v618 != 0 {
		v611 = v616
		goto L161
	} else {
		goto L163
	}
L162:
	;
	v619 = int32(1)
	switch v617&int32(255) - int32(43) {
	case 0:
		v625 = v619
		goto L165
	default:
		v627 = v617
		v628 = v611
		v629 = v619
		goto L164
	case 2:
		goto L166
	}
L163:
	;
	goto L162
L164:
	;
	v630 = int32(0)
	v632 = v627 - int32(48)
	if base.Ui32(v632) <= base.Ui32(int32(9)) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	v626 = int32(*(*int8)(unsafe.Add(mBase, uint32(v616))))
	v627 = v626
	v628 = v616
	v629 = v625
	goto L164
L166:
	;
	v625 = int32(0)
	goto L165
L167:
	;
	v635 = v630
	v636 = v632
	v637 = v628
	goto L170
L168:
	;
	v649 = v630
	goto L169
L169:
	;
	if v629 != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	v639 = int32(10)
	v641 = v635*v639 - v636
	v642 = int32(*(*int8)(unsafe.Add(mBase, uint32(v637)+1)))
	v646 = v642 - int32(48)
	if base.Ui32(v646) < base.Ui32(v639) {
		v635 = v641
		v636 = v646
		v637 = v637 + int32(1)
		goto L170
	} else {
		goto L172
	}
L171:
	;
	v649 = v641
	goto L169
L172:
	;
	goto L171
L173:
	;
	v655 = int32(0) - v649
	goto L175
L174:
	;
	v655 = v649
	goto L175
L175:
	;
	goto L160
L176:
	;
	v1152 = v666
	goto L81
L177:
	;
	v666 = int32(-13)
	goto L179
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v655
	v666 = int32(0)
	goto L179
L179:
	;
	goto L176
L180:
	;
	if v691-v692 == int32(0) {
		goto L187
	} else {
		goto L188
	}
L181:
	;
	goto L180
L182:
	;
	v676 = v196
	v677 = v667
	goto L183
L183:
	;
	v680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v677)+1)))
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v676)+1)))
	if v681 == int32(0) {
		v691 = v681
		v692 = v680
		goto L181
	} else {
		goto L185
	}
L184:
	;
	v691 = v681
	v692 = v680
	goto L181
L185:
	;
	v684 = int32(1)
	if v681 == v680 {
		v676 = v676 + v684
		v677 = v677 + v684
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v699 = v298
	goto L191
L188:
	;
	goto L189
L189:
	;
	v756 = int32(_a_F_init_work_6)
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v762 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[6])))
	if base.B2i32(v759 == int32(0))|base.B2i32(v759 != v762) != 0 {
		v780 = v759
		v781 = v762
		goto L211
	} else {
		goto L212
	}
L190:
	;
	v745 = int32(-13)
	if base.Ui32(int32(65010688)) < base.Ui32(v743-int32(1024)) {
		v755 = v745
		goto L207
	} else {
		goto L208
	}
L191:
	;
	v704 = v699 + int32(1)
	v705 = int32(*(*int8)(unsafe.Add(mBase, uint32(v699))))
	v706 = F___isspace(m, v705)
	mBase = m.M
	if v706 != 0 {
		v699 = v704
		goto L191
	} else {
		goto L193
	}
L192:
	;
	v707 = int32(1)
	switch v705&int32(255) - int32(43) {
	case 0:
		v713 = v707
		goto L195
	default:
		v715 = v705
		v716 = v699
		v717 = v707
		goto L194
	case 2:
		goto L196
	}
L193:
	;
	goto L192
L194:
	;
	v718 = int32(0)
	v720 = v715 - int32(48)
	if base.Ui32(v720) <= base.Ui32(int32(9)) {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	v714 = int32(*(*int8)(unsafe.Add(mBase, uint32(v704))))
	v715 = v714
	v716 = v704
	v717 = v713
	goto L194
L196:
	;
	v713 = int32(0)
	goto L195
L197:
	;
	v723 = v718
	v724 = v720
	v725 = v716
	goto L200
L198:
	;
	v737 = v718
	goto L199
L199:
	;
	if v717 != 0 {
		goto L203
	} else {
		goto L204
	}
L200:
	;
	v727 = int32(10)
	v729 = v723*v727 - v724
	v730 = int32(*(*int8)(unsafe.Add(mBase, uint32(v725)+1)))
	v734 = v730 - int32(48)
	if base.Ui32(v734) < base.Ui32(v727) {
		v723 = v729
		v724 = v734
		v725 = v725 + int32(1)
		goto L200
	} else {
		goto L202
	}
L201:
	;
	v737 = v729
	goto L199
L202:
	;
	goto L201
L203:
	;
	v743 = int32(0) - v737
	goto L205
L204:
	;
	v743 = v737
	goto L205
L205:
	;
	goto L190
L206:
	;
	v1152 = v755
	goto L81
L207:
	;
	goto L206
L208:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v40)+44))
	if v750 != int32(3) {
		v755 = v745
		goto L207
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+48)) = v743
	v755 = int32(0)
	goto L207
L210:
	;
	if v780-v781 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L211:
	;
	goto L210
L212:
	;
	v765 = v196
	v766 = v756
	goto L213
L213:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v766)+1)))
	v770 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v765)+1)))
	if v770 == int32(0) {
		v780 = v770
		v781 = v769
		goto L211
	} else {
		goto L215
	}
L214:
	;
	v780 = v770
	v781 = v769
	goto L211
L215:
	;
	v773 = int32(1)
	if v770 == v769 {
		v765 = v765 + v773
		v766 = v766 + v773
		goto L213
	} else {
		goto L216
	}
L216:
	;
	goto L214
L217:
	;
	v785 = F_pgp_get_digest_code(m, v298)
	mBase = m.M
	if v785 < int32(0) {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	goto L219
L219:
	;
	v791 = int32(_a_F_init_work_7)
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v797 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[7])))
	if base.B2i32(v794 == int32(0))|base.B2i32(v794 != v797) != 0 {
		v815 = v794
		v816 = v797
		goto L225
	} else {
		goto L226
	}
L220:
	;
	v1152 = v790
	goto L81
L221:
	;
	v790 = v785
	goto L220
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+52)) = v785
	v790 = int32(0)
	goto L220
L224:
	;
	if v815-v816 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L225:
	;
	goto L224
L226:
	;
	v800 = v196
	v801 = v791
	goto L227
L227:
	;
	v804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801)+1)))
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v800)+1)))
	if v805 == int32(0) {
		v815 = v805
		v816 = v804
		goto L225
	} else {
		goto L229
	}
L228:
	;
	v815 = v805
	v816 = v804
	goto L225
L229:
	;
	v808 = int32(1)
	if v805 == v804 {
		v800 = v800 + v808
		v801 = v801 + v808
		goto L227
	} else {
		goto L230
	}
L230:
	;
	goto L228
L231:
	;
	v820 = F_pgp_get_cipher_code(m, v298)
	mBase = m.M
	if v820 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L232:
	;
	goto L233
L233:
	;
	v826 = int32(_a_F_init_work_8)
	v829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[8])))
	if base.B2i32(v829 == int32(0))|base.B2i32(v829 != v832) != 0 {
		v850 = v829
		v851 = v832
		goto L239
	} else {
		goto L240
	}
L234:
	;
	v1152 = v825
	goto L81
L235:
	;
	v825 = v820
	goto L234
L236:
	;
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+56)) = v820
	v825 = int32(0)
	goto L234
L238:
	;
	if v850-v851 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L239:
	;
	goto L238
L240:
	;
	v835 = v196
	v836 = v826
	goto L241
L241:
	;
	v839 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v836)+1)))
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+1)))
	if v840 == int32(0) {
		v850 = v840
		v851 = v839
		goto L239
	} else {
		goto L243
	}
L242:
	;
	v850 = v840
	v851 = v839
	goto L239
L243:
	;
	v843 = int32(1)
	if v840 == v839 {
		v835 = v835 + v843
		v836 = v836 + v843
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v858 = v298
	goto L249
L246:
	;
	goto L247
L247:
	;
	v909 = int32(_a_F_init_work_9)
	v912 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v915 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[9])))
	if base.B2i32(v912 == int32(0))|base.B2i32(v912 != v915) != 0 {
		v933 = v912
		v934 = v915
		goto L269
	} else {
		goto L270
	}
L248:
	;
	if base.Ui32(v902) <= base.Ui32(int32(3)) {
		goto L265
	} else {
		goto L266
	}
L249:
	;
	v863 = v858 + int32(1)
	v864 = int32(*(*int8)(unsafe.Add(mBase, uint32(v858))))
	v865 = F___isspace(m, v864)
	mBase = m.M
	if v865 != 0 {
		v858 = v863
		goto L249
	} else {
		goto L251
	}
L250:
	;
	v866 = int32(1)
	switch v864&int32(255) - int32(43) {
	case 0:
		v872 = v866
		goto L253
	default:
		v874 = v864
		v875 = v858
		v876 = v866
		goto L252
	case 2:
		goto L254
	}
L251:
	;
	goto L250
L252:
	;
	v877 = int32(0)
	v879 = v874 - int32(48)
	if base.Ui32(v879) <= base.Ui32(int32(9)) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	v873 = int32(*(*int8)(unsafe.Add(mBase, uint32(v863))))
	v874 = v873
	v875 = v863
	v876 = v872
	goto L252
L254:
	;
	v872 = int32(0)
	goto L253
L255:
	;
	v882 = v877
	v883 = v879
	v884 = v875
	goto L258
L256:
	;
	v896 = v877
	goto L257
L257:
	;
	if v876 != 0 {
		goto L261
	} else {
		goto L262
	}
L258:
	;
	v886 = int32(10)
	v888 = v882*v886 - v883
	v889 = int32(*(*int8)(unsafe.Add(mBase, uint32(v884)+1)))
	v893 = v889 - int32(48)
	if base.Ui32(v893) < base.Ui32(v886) {
		v882 = v888
		v883 = v893
		v884 = v884 + int32(1)
		goto L258
	} else {
		goto L260
	}
L259:
	;
	v896 = v888
	goto L257
L260:
	;
	goto L259
L261:
	;
	v902 = int32(0) - v896
	goto L263
L262:
	;
	v902 = v896
	goto L263
L263:
	;
	goto L248
L264:
	;
	v1152 = v908
	goto L81
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+64)) = v902
	v908 = int32(0)
	goto L267
L266:
	;
	v908 = int32(-13)
	goto L267
L267:
	;
	goto L264
L268:
	;
	if v933-v934 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L269:
	;
	goto L268
L270:
	;
	v918 = v196
	v919 = v909
	goto L271
L271:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v919)+1)))
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v918)+1)))
	if v923 == int32(0) {
		v933 = v923
		v934 = v922
		goto L269
	} else {
		goto L273
	}
L272:
	;
	v933 = v923
	v934 = v922
	goto L269
L273:
	;
	v926 = int32(1)
	if v923 == v922 {
		v918 = v918 + v926
		v919 = v919 + v926
		goto L271
	} else {
		goto L274
	}
L274:
	;
	goto L272
L275:
	;
	v941 = v298
	goto L279
L276:
	;
	goto L277
L277:
	;
	v992 = int32(_a_F_init_work_10)
	v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v998 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[10])))
	if base.B2i32(v995 == int32(0))|base.B2i32(v995 != v998) != 0 {
		v1016 = v995
		v1017 = v998
		goto L299
	} else {
		goto L300
	}
L278:
	;
	if base.Ui32(v985) <= base.Ui32(int32(9)) {
		goto L295
	} else {
		goto L296
	}
L279:
	;
	v946 = v941 + int32(1)
	v947 = int32(*(*int8)(unsafe.Add(mBase, uint32(v941))))
	v948 = F___isspace(m, v947)
	mBase = m.M
	if v948 != 0 {
		v941 = v946
		goto L279
	} else {
		goto L281
	}
L280:
	;
	v949 = int32(1)
	switch v947&int32(255) - int32(43) {
	case 0:
		v955 = v949
		goto L283
	default:
		v957 = v947
		v958 = v941
		v959 = v949
		goto L282
	case 2:
		goto L284
	}
L281:
	;
	goto L280
L282:
	;
	v960 = int32(0)
	v962 = v957 - int32(48)
	if base.Ui32(v962) <= base.Ui32(int32(9)) {
		goto L285
	} else {
		goto L286
	}
L283:
	;
	v956 = int32(*(*int8)(unsafe.Add(mBase, uint32(v946))))
	v957 = v956
	v958 = v946
	v959 = v955
	goto L282
L284:
	;
	v955 = int32(0)
	goto L283
L285:
	;
	v965 = v960
	v966 = v962
	v967 = v958
	goto L288
L286:
	;
	v979 = v960
	goto L287
L287:
	;
	if v959 != 0 {
		goto L291
	} else {
		goto L292
	}
L288:
	;
	v969 = int32(10)
	v971 = v965*v969 - v966
	v972 = int32(*(*int8)(unsafe.Add(mBase, uint32(v967)+1)))
	v976 = v972 - int32(48)
	if base.Ui32(v976) < base.Ui32(v969) {
		v965 = v971
		v966 = v976
		v967 = v967 + int32(1)
		goto L288
	} else {
		goto L290
	}
L289:
	;
	v979 = v971
	goto L287
L290:
	;
	goto L289
L291:
	;
	v985 = int32(0) - v979
	goto L293
L292:
	;
	v985 = v979
	goto L293
L293:
	;
	goto L278
L294:
	;
	v1152 = v991
	goto L81
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40)+68)) = v985
	v991 = int32(0)
	goto L297
L296:
	;
	v991 = int32(-13)
	goto L297
L297:
	;
	goto L294
L298:
	;
	if v1016-v1017 == int32(0) {
		goto L305
	} else {
		goto L306
	}
L299:
	;
	goto L298
L300:
	;
	v1001 = v196
	v1002 = v992
	goto L301
L301:
	;
	v1005 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1002)+1)))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+1)))
	if v1006 == int32(0) {
		v1016 = v1006
		v1017 = v1005
		goto L299
	} else {
		goto L303
	}
L302:
	;
	v1016 = v1006
	v1017 = v1005
	goto L299
L303:
	;
	v1009 = int32(1)
	if v1006 == v1005 {
		v1001 = v1001 + v1009
		v1002 = v1002 + v1009
		goto L301
	} else {
		goto L304
	}
L304:
	;
	goto L302
L305:
	;
	v1024 = v298
	goto L309
L306:
	;
	goto L307
L307:
	;
	v1073 = int32(_a_F_init_work_11)
	v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[11])))
	if base.B2i32(v1076 == int32(0))|base.B2i32(v1076 != v1079) != 0 {
		v1097 = v1076
		v1098 = v1079
		goto L326
	} else {
		goto L327
	}
L308:
	;
	v1069 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+84)) = base.B2i32(v1068 != v1069)
	goto L324
L309:
	;
	v1029 = v1024 + int32(1)
	v1030 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1024))))
	v1031 = F___isspace(m, v1030)
	mBase = m.M
	if v1031 != 0 {
		v1024 = v1029
		goto L309
	} else {
		goto L311
	}
L310:
	;
	v1032 = int32(1)
	switch v1030&int32(255) - int32(43) {
	case 0:
		v1038 = v1032
		goto L313
	default:
		v1040 = v1030
		v1041 = v1024
		v1042 = v1032
		goto L312
	case 2:
		goto L314
	}
L311:
	;
	goto L310
L312:
	;
	v1043 = int32(0)
	v1045 = v1040 - int32(48)
	if base.Ui32(v1045) <= base.Ui32(int32(9)) {
		goto L315
	} else {
		goto L316
	}
L313:
	;
	v1039 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1029))))
	v1040 = v1039
	v1041 = v1029
	v1042 = v1038
	goto L312
L314:
	;
	v1038 = int32(0)
	goto L313
L315:
	;
	v1048 = v1043
	v1049 = v1045
	v1050 = v1041
	goto L318
L316:
	;
	v1062 = v1043
	goto L317
L317:
	;
	if v1042 != 0 {
		goto L321
	} else {
		goto L322
	}
L318:
	;
	v1052 = int32(10)
	v1054 = v1048*v1052 - v1049
	v1055 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1050)+1)))
	v1059 = v1055 - int32(48)
	if base.Ui32(v1059) < base.Ui32(v1052) {
		v1048 = v1054
		v1049 = v1059
		v1050 = v1050 + int32(1)
		goto L318
	} else {
		goto L320
	}
L319:
	;
	v1062 = v1054
	goto L317
L320:
	;
	goto L319
L321:
	;
	v1068 = int32(0) - v1062
	goto L323
L322:
	;
	v1068 = v1062
	goto L323
L323:
	;
	goto L308
L324:
	;
	v1152 = v1069
	goto L81
L325:
	;
	if v1097-v1098 != 0 {
		goto L80
	} else {
		goto L332
	}
L326:
	;
	goto L325
L327:
	;
	v1082 = v196
	v1083 = v1073
	goto L328
L328:
	;
	v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1083)+1)))
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1082)+1)))
	if v1087 == int32(0) {
		v1097 = v1087
		v1098 = v1086
		goto L326
	} else {
		goto L330
	}
L329:
	;
	v1097 = v1087
	v1098 = v1086
	goto L326
L330:
	;
	v1090 = int32(1)
	if v1087 == v1086 {
		v1082 = v1082 + v1090
		v1083 = v1083 + v1090
		goto L328
	} else {
		goto L331
	}
L331:
	;
	goto L329
L332:
	;
	v1103 = v298
	goto L334
L333:
	;
	v1148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v40)+88)) = base.B2i32(v1147 != v1148)
	goto L349
L334:
	;
	v1108 = v1103 + int32(1)
	v1109 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1103))))
	v1110 = F___isspace(m, v1109)
	mBase = m.M
	if v1110 != 0 {
		v1103 = v1108
		goto L334
	} else {
		goto L336
	}
L335:
	;
	v1111 = int32(1)
	switch v1109&int32(255) - int32(43) {
	case 0:
		v1117 = v1111
		goto L338
	default:
		v1119 = v1109
		v1120 = v1103
		v1121 = v1111
		goto L337
	case 2:
		goto L339
	}
L336:
	;
	goto L335
L337:
	;
	v1122 = int32(0)
	v1124 = v1119 - int32(48)
	if base.Ui32(v1124) <= base.Ui32(int32(9)) {
		goto L340
	} else {
		goto L341
	}
L338:
	;
	v1118 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1108))))
	v1119 = v1118
	v1120 = v1108
	v1121 = v1117
	goto L337
L339:
	;
	v1117 = int32(0)
	goto L338
L340:
	;
	v1127 = v1122
	v1128 = v1124
	v1129 = v1120
	goto L343
L341:
	;
	v1141 = v1122
	goto L342
L342:
	;
	if v1121 != 0 {
		goto L346
	} else {
		goto L347
	}
L343:
	;
	v1131 = int32(10)
	v1133 = v1127*v1131 - v1128
	v1134 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1129)+1)))
	v1138 = v1134 - int32(48)
	if base.Ui32(v1138) < base.Ui32(v1131) {
		v1127 = v1133
		v1128 = v1138
		v1129 = v1129 + int32(1)
		goto L343
	} else {
		goto L345
	}
L344:
	;
	v1141 = v1133
	goto L342
L345:
	;
	goto L344
L346:
	;
	v1147 = int32(0) - v1141
	goto L348
L347:
	;
	v1147 = v1141
	goto L348
L348:
	;
	goto L333
L349:
	;
	v1152 = v1148
	goto L81
L350:
	;
	v1963 = v1152
	goto L41
L351:
	;
	if v1179-v1180 == int32(0) {
		goto L358
	} else {
		goto L359
	}
L352:
	;
	goto L351
L353:
	;
	v1164 = v196
	v1165 = v1155
	goto L354
L354:
	;
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1165)+1)))
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1164)+1)))
	if v1169 == int32(0) {
		v1179 = v1169
		v1180 = v1168
		goto L352
	} else {
		goto L356
	}
L355:
	;
	v1179 = v1169
	v1180 = v1168
	goto L352
L356:
	;
	v1172 = int32(1)
	if v1169 == v1168 {
		v1164 = v1164 + v1172
		v1165 = v1165 + v1172
		goto L354
	} else {
		goto L357
	}
L357:
	;
	goto L355
L358:
	;
	v1187 = v298
	goto L362
L359:
	;
	goto L360
L360:
	;
	v1234 = int32(_a_F_init_work_12)
	v1237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1240 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[12])))
	if base.B2i32(v1237 == int32(0))|base.B2i32(v1237 != v1240) != 0 {
		v1258 = v1237
		v1259 = v1240
		goto L378
	} else {
		goto L379
	}
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v1231
	v1954 = int32(0)
	goto L79
L362:
	;
	v1192 = v1187 + int32(1)
	v1193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1187))))
	v1194 = F___isspace(m, v1193)
	mBase = m.M
	if v1194 != 0 {
		v1187 = v1192
		goto L362
	} else {
		goto L364
	}
L363:
	;
	v1195 = int32(1)
	switch v1193&int32(255) - int32(43) {
	case 0:
		v1201 = v1195
		goto L366
	default:
		v1203 = v1193
		v1204 = v1187
		v1205 = v1195
		goto L365
	case 2:
		goto L367
	}
L364:
	;
	goto L363
L365:
	;
	v1206 = int32(0)
	v1208 = v1203 - int32(48)
	if base.Ui32(v1208) <= base.Ui32(int32(9)) {
		goto L368
	} else {
		goto L369
	}
L366:
	;
	v1202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1192))))
	v1203 = v1202
	v1204 = v1192
	v1205 = v1201
	goto L365
L367:
	;
	v1201 = int32(0)
	goto L366
L368:
	;
	v1211 = v1206
	v1212 = v1208
	v1213 = v1204
	goto L371
L369:
	;
	v1225 = v1206
	goto L370
L370:
	;
	if v1205 != 0 {
		goto L374
	} else {
		goto L375
	}
L371:
	;
	v1215 = int32(10)
	v1217 = v1211*v1215 - v1212
	v1218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1213)+1)))
	v1222 = v1218 - int32(48)
	if base.Ui32(v1222) < base.Ui32(v1215) {
		v1211 = v1217
		v1212 = v1222
		v1213 = v1213 + int32(1)
		goto L371
	} else {
		goto L373
	}
L372:
	;
	v1225 = v1217
	goto L370
L373:
	;
	goto L372
L374:
	;
	v1231 = int32(0) - v1225
	goto L376
L375:
	;
	v1231 = v1225
	goto L376
L376:
	;
	goto L361
L377:
	;
	if v1258-v1259 == int32(0) {
		goto L384
	} else {
		goto L385
	}
L378:
	;
	goto L377
L379:
	;
	v1243 = v196
	v1244 = v1234
	goto L380
L380:
	;
	v1247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1244)+1)))
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1243)+1)))
	if v1248 == int32(0) {
		v1258 = v1248
		v1259 = v1247
		goto L378
	} else {
		goto L382
	}
L381:
	;
	v1258 = v1248
	v1259 = v1247
	goto L378
L382:
	;
	v1251 = int32(1)
	if v1248 == v1247 {
		v1243 = v1243 + v1251
		v1244 = v1244 + v1251
		goto L380
	} else {
		goto L383
	}
L383:
	;
	goto L381
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1268 = F_pg_strcasecmp(m, int32(_a_F_init_work_13), v298)
	mBase = m.M
	if v1268 == int32(0) {
		v1311 = int32(_a_F_init_work_14)
		goto L389
	} else {
		goto L390
	}
L385:
	;
	goto L386
L386:
	;
	v1316 = int32(_a_F_init_work_15)
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[13])))
	if base.B2i32(v1319 == int32(0))|base.B2i32(v1319 != v1322) != 0 {
		v1340 = v1319
		v1341 = v1322
		goto L400
	} else {
		goto L401
	}
L387:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v1313
	v1954 = int32(0)
	goto L79
L388:
	;
	goto L387
L389:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1311)+4))
	v1313 = v1312
	goto L388
L390:
	;
	v1273 = F_pg_strcasecmp(m, int32(_a_F_init_work_16), v298)
	mBase = m.M
	if v1273 == int32(0) {
		v1311 = int32(_a_F_init_work_17)
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v1278 = F_pg_strcasecmp(m, int32(_a_F_init_work_18), v298)
	mBase = m.M
	if v1278 == int32(0) {
		v1311 = int32(_a_F_init_work_19)
		goto L389
	} else {
		goto L392
	}
L392:
	;
	v1283 = F_pg_strcasecmp(m, int32(_a_F_init_work_20), v298)
	mBase = m.M
	if v1283 == int32(0) {
		v1311 = int32(_a_F_init_work_21)
		goto L389
	} else {
		goto L393
	}
L393:
	;
	v1288 = F_pg_strcasecmp(m, int32(_a_F_init_work_22), v298)
	mBase = m.M
	if v1288 == int32(0) {
		v1311 = int32(_a_F_init_work_23)
		goto L389
	} else {
		goto L394
	}
L394:
	;
	v1293 = F_pg_strcasecmp(m, int32(_a_F_init_work_24), v298)
	mBase = m.M
	if v1293 == int32(0) {
		v1311 = int32(_a_F_init_work_25)
		goto L389
	} else {
		goto L395
	}
L395:
	;
	v1298 = F_pg_strcasecmp(m, int32(_a_F_init_work_26), v298)
	mBase = m.M
	if v1298 == int32(0) {
		v1311 = int32(_a_F_init_work_27)
		goto L389
	} else {
		goto L396
	}
L396:
	;
	v1303 = F_pg_strcasecmp(m, int32(_a_F_init_work_28), v298)
	mBase = m.M
	if v1303 == int32(0) {
		v1311 = int32(_a_F_init_work_29)
		goto L389
	} else {
		goto L397
	}
L397:
	;
	v1308 = F_pg_strcasecmp(m, int32(_a_F_init_work_30), v298)
	mBase = m.M
	if v1308 != 0 {
		v1313 = int32(-103)
		goto L388
	} else {
		goto L398
	}
L398:
	;
	v1311 = int32(_a_F_init_work_31)
	goto L389
L399:
	;
	if v1340-v1341 == int32(0) {
		goto L406
	} else {
		goto L407
	}
L400:
	;
	goto L399
L401:
	;
	v1325 = v196
	v1326 = v1316
	goto L402
L402:
	;
	v1329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1326)+1)))
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1325)+1)))
	if v1330 == int32(0) {
		v1340 = v1330
		v1341 = v1329
		goto L400
	} else {
		goto L404
	}
L403:
	;
	v1340 = v1330
	v1341 = v1329
	goto L400
L404:
	;
	v1333 = int32(1)
	if v1330 == v1329 {
		v1325 = v1325 + v1333
		v1326 = v1326 + v1333
		goto L402
	} else {
		goto L405
	}
L405:
	;
	goto L403
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1350 = v298
	goto L410
L407:
	;
	goto L408
L408:
	;
	v1397 = int32(_a_F_init_work_32)
	v1400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[14])))
	if base.B2i32(v1400 == int32(0))|base.B2i32(v1400 != v1403) != 0 {
		v1421 = v1400
		v1422 = v1403
		goto L426
	} else {
		goto L427
	}
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v1394
	v1954 = int32(0)
	goto L79
L410:
	;
	v1355 = v1350 + int32(1)
	v1356 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1350))))
	v1357 = F___isspace(m, v1356)
	mBase = m.M
	if v1357 != 0 {
		v1350 = v1355
		goto L410
	} else {
		goto L412
	}
L411:
	;
	v1358 = int32(1)
	switch v1356&int32(255) - int32(43) {
	case 0:
		v1364 = v1358
		goto L414
	default:
		v1366 = v1356
		v1367 = v1350
		v1368 = v1358
		goto L413
	case 2:
		goto L415
	}
L412:
	;
	goto L411
L413:
	;
	v1369 = int32(0)
	v1371 = v1366 - int32(48)
	if base.Ui32(v1371) <= base.Ui32(int32(9)) {
		goto L416
	} else {
		goto L417
	}
L414:
	;
	v1365 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1355))))
	v1366 = v1365
	v1367 = v1355
	v1368 = v1364
	goto L413
L415:
	;
	v1364 = int32(0)
	goto L414
L416:
	;
	v1374 = v1369
	v1375 = v1371
	v1376 = v1367
	goto L419
L417:
	;
	v1388 = v1369
	goto L418
L418:
	;
	if v1368 != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	v1378 = int32(10)
	v1380 = v1374*v1378 - v1375
	v1381 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1376)+1)))
	v1385 = v1381 - int32(48)
	if base.Ui32(v1385) < base.Ui32(v1378) {
		v1374 = v1380
		v1375 = v1385
		v1376 = v1376 + int32(1)
		goto L419
	} else {
		goto L421
	}
L420:
	;
	v1388 = v1380
	goto L418
L421:
	;
	goto L420
L422:
	;
	v1394 = int32(0) - v1388
	goto L424
L423:
	;
	v1394 = v1388
	goto L424
L424:
	;
	goto L409
L425:
	;
	if v1421-v1422 == int32(0) {
		goto L432
	} else {
		goto L433
	}
L426:
	;
	goto L425
L427:
	;
	v1406 = v196
	v1407 = v1397
	goto L428
L428:
	;
	v1410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1407)+1)))
	v1411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406)+1)))
	if v1411 == int32(0) {
		v1421 = v1411
		v1422 = v1410
		goto L426
	} else {
		goto L430
	}
L429:
	;
	v1421 = v1411
	v1422 = v1410
	goto L426
L430:
	;
	v1414 = int32(1)
	if v1411 == v1410 {
		v1406 = v1406 + v1414
		v1407 = v1407 + v1414
		goto L428
	} else {
		goto L431
	}
L431:
	;
	goto L429
L432:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1431 = v298
	goto L436
L433:
	;
	goto L434
L434:
	;
	v1478 = int32(_a_F_init_work_33)
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1484 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[15])))
	if base.B2i32(v1481 == int32(0))|base.B2i32(v1481 != v1484) != 0 {
		v1502 = v1481
		v1503 = v1484
		goto L452
	} else {
		goto L453
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v1475
	v1954 = int32(0)
	goto L79
L436:
	;
	v1436 = v1431 + int32(1)
	v1437 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1431))))
	v1438 = F___isspace(m, v1437)
	mBase = m.M
	if v1438 != 0 {
		v1431 = v1436
		goto L436
	} else {
		goto L438
	}
L437:
	;
	v1439 = int32(1)
	switch v1437&int32(255) - int32(43) {
	case 0:
		v1445 = v1439
		goto L440
	default:
		v1447 = v1437
		v1448 = v1431
		v1449 = v1439
		goto L439
	case 2:
		goto L441
	}
L438:
	;
	goto L437
L439:
	;
	v1450 = int32(0)
	v1452 = v1447 - int32(48)
	if base.Ui32(v1452) <= base.Ui32(int32(9)) {
		goto L442
	} else {
		goto L443
	}
L440:
	;
	v1446 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1436))))
	v1447 = v1446
	v1448 = v1436
	v1449 = v1445
	goto L439
L441:
	;
	v1445 = int32(0)
	goto L440
L442:
	;
	v1455 = v1450
	v1456 = v1452
	v1457 = v1448
	goto L445
L443:
	;
	v1469 = v1450
	goto L444
L444:
	;
	if v1449 != 0 {
		goto L448
	} else {
		goto L449
	}
L445:
	;
	v1459 = int32(10)
	v1461 = v1455*v1459 - v1456
	v1462 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1457)+1)))
	v1466 = v1462 - int32(48)
	if base.Ui32(v1466) < base.Ui32(v1459) {
		v1455 = v1461
		v1456 = v1466
		v1457 = v1457 + int32(1)
		goto L445
	} else {
		goto L447
	}
L446:
	;
	v1469 = v1461
	goto L444
L447:
	;
	goto L446
L448:
	;
	v1475 = int32(0) - v1469
	goto L450
L449:
	;
	v1475 = v1469
	goto L450
L450:
	;
	goto L435
L451:
	;
	if v1502-v1503 == int32(0) {
		goto L458
	} else {
		goto L459
	}
L452:
	;
	goto L451
L453:
	;
	v1487 = v196
	v1488 = v1478
	goto L454
L454:
	;
	v1491 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1488)+1)))
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1487)+1)))
	if v1492 == int32(0) {
		v1502 = v1492
		v1503 = v1491
		goto L452
	} else {
		goto L456
	}
L455:
	;
	v1502 = v1492
	v1503 = v1491
	goto L452
L456:
	;
	v1495 = int32(1)
	if v1492 == v1491 {
		v1487 = v1487 + v1495
		v1488 = v1488 + v1495
		goto L454
	} else {
		goto L457
	}
L457:
	;
	goto L455
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1512 = v298
	goto L462
L459:
	;
	goto L460
L460:
	;
	v1559 = int32(_a_F_init_work_34)
	v1562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[16])))
	if base.B2i32(v1562 == int32(0))|base.B2i32(v1562 != v1565) != 0 {
		v1583 = v1562
		v1584 = v1565
		goto L478
	} else {
		goto L479
	}
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v1556
	v1954 = int32(0)
	goto L79
L462:
	;
	v1517 = v1512 + int32(1)
	v1518 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1512))))
	v1519 = F___isspace(m, v1518)
	mBase = m.M
	if v1519 != 0 {
		v1512 = v1517
		goto L462
	} else {
		goto L464
	}
L463:
	;
	v1520 = int32(1)
	switch v1518&int32(255) - int32(43) {
	case 0:
		v1526 = v1520
		goto L466
	default:
		v1528 = v1518
		v1529 = v1512
		v1530 = v1520
		goto L465
	case 2:
		goto L467
	}
L464:
	;
	goto L463
L465:
	;
	v1531 = int32(0)
	v1533 = v1528 - int32(48)
	if base.Ui32(v1533) <= base.Ui32(int32(9)) {
		goto L468
	} else {
		goto L469
	}
L466:
	;
	v1527 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1517))))
	v1528 = v1527
	v1529 = v1517
	v1530 = v1526
	goto L465
L467:
	;
	v1526 = int32(0)
	goto L466
L468:
	;
	v1536 = v1531
	v1537 = v1533
	v1538 = v1529
	goto L471
L469:
	;
	v1550 = v1531
	goto L470
L470:
	;
	if v1530 != 0 {
		goto L474
	} else {
		goto L475
	}
L471:
	;
	v1540 = int32(10)
	v1542 = v1536*v1540 - v1537
	v1543 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1538)+1)))
	v1547 = v1543 - int32(48)
	if base.Ui32(v1547) < base.Ui32(v1540) {
		v1536 = v1542
		v1537 = v1547
		v1538 = v1538 + int32(1)
		goto L471
	} else {
		goto L473
	}
L472:
	;
	v1550 = v1542
	goto L470
L473:
	;
	goto L472
L474:
	;
	v1556 = int32(0) - v1550
	goto L476
L475:
	;
	v1556 = v1550
	goto L476
L476:
	;
	goto L461
L477:
	;
	if v1583-v1584 == int32(0) {
		goto L484
	} else {
		goto L485
	}
L478:
	;
	goto L477
L479:
	;
	v1568 = v196
	v1569 = v1559
	goto L480
L480:
	;
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1569)+1)))
	v1573 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1568)+1)))
	if v1573 == int32(0) {
		v1583 = v1573
		v1584 = v1572
		goto L478
	} else {
		goto L482
	}
L481:
	;
	v1583 = v1573
	v1584 = v1572
	goto L478
L482:
	;
	v1576 = int32(1)
	if v1573 == v1572 {
		v1568 = v1568 + v1576
		v1569 = v1569 + v1576
		goto L480
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1593 = v298
	goto L488
L485:
	;
	goto L486
L486:
	;
	v1640 = int32(_a_F_init_work_35)
	v1643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1646 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[17])))
	if base.B2i32(v1643 == int32(0))|base.B2i32(v1643 != v1646) != 0 {
		v1664 = v1643
		v1665 = v1646
		goto L504
	} else {
		goto L505
	}
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v1637
	v1954 = int32(0)
	goto L79
L488:
	;
	v1598 = v1593 + int32(1)
	v1599 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1593))))
	v1600 = F___isspace(m, v1599)
	mBase = m.M
	if v1600 != 0 {
		v1593 = v1598
		goto L488
	} else {
		goto L490
	}
L489:
	;
	v1601 = int32(1)
	switch v1599&int32(255) - int32(43) {
	case 0:
		v1607 = v1601
		goto L492
	default:
		v1609 = v1599
		v1610 = v1593
		v1611 = v1601
		goto L491
	case 2:
		goto L493
	}
L490:
	;
	goto L489
L491:
	;
	v1612 = int32(0)
	v1614 = v1609 - int32(48)
	if base.Ui32(v1614) <= base.Ui32(int32(9)) {
		goto L494
	} else {
		goto L495
	}
L492:
	;
	v1608 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1598))))
	v1609 = v1608
	v1610 = v1598
	v1611 = v1607
	goto L491
L493:
	;
	v1607 = int32(0)
	goto L492
L494:
	;
	v1617 = v1612
	v1618 = v1614
	v1619 = v1610
	goto L497
L495:
	;
	v1631 = v1612
	goto L496
L496:
	;
	if v1611 != 0 {
		goto L500
	} else {
		goto L501
	}
L497:
	;
	v1621 = int32(10)
	v1623 = v1617*v1621 - v1618
	v1624 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1619)+1)))
	v1628 = v1624 - int32(48)
	if base.Ui32(v1628) < base.Ui32(v1621) {
		v1617 = v1623
		v1618 = v1628
		v1619 = v1619 + int32(1)
		goto L497
	} else {
		goto L499
	}
L498:
	;
	v1631 = v1623
	goto L496
L499:
	;
	goto L498
L500:
	;
	v1637 = int32(0) - v1631
	goto L502
L501:
	;
	v1637 = v1631
	goto L502
L502:
	;
	goto L487
L503:
	;
	if v1664-v1665 == int32(0) {
		goto L510
	} else {
		goto L511
	}
L504:
	;
	goto L503
L505:
	;
	v1649 = v196
	v1650 = v1640
	goto L506
L506:
	;
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650)+1)))
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1649)+1)))
	if v1654 == int32(0) {
		v1664 = v1654
		v1665 = v1653
		goto L504
	} else {
		goto L508
	}
L507:
	;
	v1664 = v1654
	v1665 = v1653
	goto L504
L508:
	;
	v1657 = int32(1)
	if v1654 == v1653 {
		v1649 = v1649 + v1657
		v1650 = v1650 + v1657
		goto L506
	} else {
		goto L509
	}
L509:
	;
	goto L507
L510:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1674 = F_pg_strcasecmp(m, int32(_a_F_init_work_36), v298)
	mBase = m.M
	if v1674 == int32(0) {
		v1707 = int32(_a_F_init_work_37)
		goto L515
	} else {
		goto L516
	}
L511:
	;
	goto L512
L512:
	;
	v1712 = int32(_a_F_init_work_38)
	v1715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1718 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[18])))
	if base.B2i32(v1715 == int32(0))|base.B2i32(v1715 != v1718) != 0 {
		v1736 = v1715
		v1737 = v1718
		goto L524
	} else {
		goto L525
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v1709
	v1954 = int32(0)
	goto L79
L514:
	;
	goto L513
L515:
	;
	v1708 = *(*int32)(unsafe.Add(mBase, uint32(v1707)+4))
	v1709 = v1708
	goto L514
L516:
	;
	v1679 = F_pg_strcasecmp(m, int32(_a_F_init_work_39), v298)
	mBase = m.M
	if v1679 == int32(0) {
		v1707 = int32(_a_F_init_work_40)
		goto L515
	} else {
		goto L517
	}
L517:
	;
	v1684 = F_pg_strcasecmp(m, int32(_a_F_init_work_41), v298)
	mBase = m.M
	if v1684 == int32(0) {
		v1707 = int32(_a_F_init_work_42)
		goto L515
	} else {
		goto L518
	}
L518:
	;
	v1689 = F_pg_strcasecmp(m, int32(_a_F_init_work_43), v298)
	mBase = m.M
	if v1689 == int32(0) {
		v1707 = int32(_a_F_init_work_44)
		goto L515
	} else {
		goto L519
	}
L519:
	;
	v1694 = F_pg_strcasecmp(m, int32(_a_F_init_work_45), v298)
	mBase = m.M
	if v1694 == int32(0) {
		v1707 = int32(_a_F_init_work_46)
		goto L515
	} else {
		goto L520
	}
L520:
	;
	v1699 = F_pg_strcasecmp(m, int32(_a_F_init_work_47), v298)
	mBase = m.M
	if v1699 == int32(0) {
		v1707 = int32(_a_F_init_work_48)
		goto L515
	} else {
		goto L521
	}
L521:
	;
	v1704 = F_pg_strcasecmp(m, int32(_a_F_init_work_49), v298)
	mBase = m.M
	if v1704 != 0 {
		v1709 = int32(-104)
		goto L514
	} else {
		goto L522
	}
L522:
	;
	v1707 = int32(_a_F_init_work_50)
	goto L515
L523:
	;
	if v1736-v1737 == int32(0) {
		goto L530
	} else {
		goto L531
	}
L524:
	;
	goto L523
L525:
	;
	v1721 = v196
	v1722 = v1712
	goto L526
L526:
	;
	v1725 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1722)+1)))
	v1726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1721)+1)))
	if v1726 == int32(0) {
		v1736 = v1726
		v1737 = v1725
		goto L524
	} else {
		goto L528
	}
L527:
	;
	v1736 = v1726
	v1737 = v1725
	goto L524
L528:
	;
	v1729 = int32(1)
	if v1726 == v1725 {
		v1721 = v1721 + v1729
		v1722 = v1722 + v1729
		goto L526
	} else {
		goto L529
	}
L529:
	;
	goto L527
L530:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1746 = F_pg_strcasecmp(m, int32(_a_F_init_work_13), v298)
	mBase = m.M
	if v1746 == int32(0) {
		v1789 = int32(_a_F_init_work_14)
		goto L535
	} else {
		goto L536
	}
L531:
	;
	goto L532
L532:
	;
	v1794 = int32(_a_F_init_work_51)
	v1797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1800 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[19])))
	if base.B2i32(v1797 == int32(0))|base.B2i32(v1797 != v1800) != 0 {
		v1818 = v1797
		v1819 = v1800
		goto L546
	} else {
		goto L547
	}
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v1791
	v1954 = int32(0)
	goto L79
L534:
	;
	goto L533
L535:
	;
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(v1789)+4))
	v1791 = v1790
	goto L534
L536:
	;
	v1751 = F_pg_strcasecmp(m, int32(_a_F_init_work_16), v298)
	mBase = m.M
	if v1751 == int32(0) {
		v1789 = int32(_a_F_init_work_17)
		goto L535
	} else {
		goto L537
	}
L537:
	;
	v1756 = F_pg_strcasecmp(m, int32(_a_F_init_work_18), v298)
	mBase = m.M
	if v1756 == int32(0) {
		v1789 = int32(_a_F_init_work_19)
		goto L535
	} else {
		goto L538
	}
L538:
	;
	v1761 = F_pg_strcasecmp(m, int32(_a_F_init_work_20), v298)
	mBase = m.M
	if v1761 == int32(0) {
		v1789 = int32(_a_F_init_work_21)
		goto L535
	} else {
		goto L539
	}
L539:
	;
	v1766 = F_pg_strcasecmp(m, int32(_a_F_init_work_22), v298)
	mBase = m.M
	if v1766 == int32(0) {
		v1789 = int32(_a_F_init_work_23)
		goto L535
	} else {
		goto L540
	}
L540:
	;
	v1771 = F_pg_strcasecmp(m, int32(_a_F_init_work_24), v298)
	mBase = m.M
	if v1771 == int32(0) {
		v1789 = int32(_a_F_init_work_25)
		goto L535
	} else {
		goto L541
	}
L541:
	;
	v1776 = F_pg_strcasecmp(m, int32(_a_F_init_work_26), v298)
	mBase = m.M
	if v1776 == int32(0) {
		v1789 = int32(_a_F_init_work_27)
		goto L535
	} else {
		goto L542
	}
L542:
	;
	v1781 = F_pg_strcasecmp(m, int32(_a_F_init_work_28), v298)
	mBase = m.M
	if v1781 == int32(0) {
		v1789 = int32(_a_F_init_work_29)
		goto L535
	} else {
		goto L543
	}
L543:
	;
	v1786 = F_pg_strcasecmp(m, int32(_a_F_init_work_30), v298)
	mBase = m.M
	if v1786 != 0 {
		v1791 = int32(-103)
		goto L534
	} else {
		goto L544
	}
L544:
	;
	v1789 = int32(_a_F_init_work_31)
	goto L535
L545:
	;
	if v1818-v1819 == int32(0) {
		goto L552
	} else {
		goto L553
	}
L546:
	;
	goto L545
L547:
	;
	v1803 = v196
	v1804 = v1794
	goto L548
L548:
	;
	v1807 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1804)+1)))
	v1808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1803)+1)))
	if v1808 == int32(0) {
		v1818 = v1808
		v1819 = v1807
		goto L546
	} else {
		goto L550
	}
L549:
	;
	v1818 = v1808
	v1819 = v1807
	goto L546
L550:
	;
	v1811 = int32(1)
	if v1808 == v1807 {
		v1803 = v1803 + v1811
		v1804 = v1804 + v1811
		goto L548
	} else {
		goto L551
	}
L551:
	;
	goto L549
L552:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1828 = v298
	goto L556
L553:
	;
	goto L554
L554:
	;
	v1875 = int32(_a_F_init_work_52)
	v1878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_init_work[20])))
	if base.B2i32(v1878 == int32(0))|base.B2i32(v1878 != v1881) != 0 {
		v1899 = v1878
		v1900 = v1881
		goto L572
	} else {
		goto L573
	}
L555:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+28)) = v1872
	v1954 = int32(0)
	goto L79
L556:
	;
	v1833 = v1828 + int32(1)
	v1834 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1828))))
	v1835 = F___isspace(m, v1834)
	mBase = m.M
	if v1835 != 0 {
		v1828 = v1833
		goto L556
	} else {
		goto L558
	}
L557:
	;
	v1836 = int32(1)
	switch v1834&int32(255) - int32(43) {
	case 0:
		v1842 = v1836
		goto L560
	default:
		v1844 = v1834
		v1845 = v1828
		v1846 = v1836
		goto L559
	case 2:
		goto L561
	}
L558:
	;
	goto L557
L559:
	;
	v1847 = int32(0)
	v1849 = v1844 - int32(48)
	if base.Ui32(v1849) <= base.Ui32(int32(9)) {
		goto L562
	} else {
		goto L563
	}
L560:
	;
	v1843 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1833))))
	v1844 = v1843
	v1845 = v1833
	v1846 = v1842
	goto L559
L561:
	;
	v1842 = int32(0)
	goto L560
L562:
	;
	v1852 = v1847
	v1853 = v1849
	v1854 = v1845
	goto L565
L563:
	;
	v1866 = v1847
	goto L564
L564:
	;
	if v1846 != 0 {
		goto L568
	} else {
		goto L569
	}
L565:
	;
	v1856 = int32(10)
	v1858 = v1852*v1856 - v1853
	v1859 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1854)+1)))
	v1863 = v1859 - int32(48)
	if base.Ui32(v1863) < base.Ui32(v1856) {
		v1852 = v1858
		v1853 = v1863
		v1854 = v1854 + int32(1)
		goto L565
	} else {
		goto L567
	}
L566:
	;
	v1866 = v1858
	goto L564
L567:
	;
	goto L566
L568:
	;
	v1872 = int32(0) - v1866
	goto L570
L569:
	;
	v1872 = v1866
	goto L570
L570:
	;
	goto L555
L571:
	;
	if v1899-v1900 != 0 {
		v1963 = v253
		goto L41
	} else {
		goto L578
	}
L572:
	;
	goto L571
L573:
	;
	v1884 = v196
	v1885 = v1875
	goto L574
L574:
	;
	v1888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1885)+1)))
	v1889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884)+1)))
	if v1889 == int32(0) {
		v1899 = v1889
		v1900 = v1888
		goto L572
	} else {
		goto L576
	}
L575:
	;
	v1899 = v1889
	v1900 = v1888
	goto L572
L576:
	;
	v1892 = int32(1)
	if v1889 == v1888 {
		v1884 = v1884 + v1892
		v1885 = v1885 + v1892
		goto L574
	} else {
		goto L577
	}
L577:
	;
	goto L575
L578:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = int32(1)
	v1907 = v298
	goto L580
L579:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v1951
	v1954 = int32(0)
	goto L79
L580:
	;
	v1912 = v1907 + int32(1)
	v1913 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1907))))
	v1914 = F___isspace(m, v1913)
	mBase = m.M
	if v1914 != 0 {
		v1907 = v1912
		goto L580
	} else {
		goto L582
	}
L581:
	;
	v1915 = int32(1)
	switch v1913&int32(255) - int32(43) {
	case 0:
		v1921 = v1915
		goto L584
	default:
		v1923 = v1913
		v1924 = v1907
		v1925 = v1915
		goto L583
	case 2:
		goto L585
	}
L582:
	;
	goto L581
L583:
	;
	v1926 = int32(0)
	v1928 = v1923 - int32(48)
	if base.Ui32(v1928) <= base.Ui32(int32(9)) {
		goto L586
	} else {
		goto L587
	}
L584:
	;
	v1922 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1912))))
	v1923 = v1922
	v1924 = v1912
	v1925 = v1921
	goto L583
L585:
	;
	v1921 = int32(0)
	goto L584
L586:
	;
	v1931 = v1926
	v1932 = v1928
	v1933 = v1924
	goto L589
L587:
	;
	v1945 = v1926
	goto L588
L588:
	;
	if v1925 != 0 {
		goto L592
	} else {
		goto L593
	}
L589:
	;
	v1935 = int32(10)
	v1937 = v1931*v1935 - v1932
	v1938 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1933)+1)))
	v1942 = v1938 - int32(48)
	if base.Ui32(v1942) < base.Ui32(v1935) {
		v1931 = v1937
		v1932 = v1942
		v1933 = v1933 + int32(1)
		goto L589
	} else {
		goto L591
	}
L590:
	;
	v1945 = v1937
	goto L588
L591:
	;
	goto L590
L592:
	;
	v1951 = int32(0) - v1945
	goto L594
L593:
	;
	v1951 = v1945
	goto L594
L594:
	;
	goto L579
L595:
	;
	goto L44
L596:
	;
	v1980 = v1963
	goto L5
L597:
	;
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v1990 != 0 {
		goto L600
	} else {
		goto L601
	}
L598:
	;
	goto L599
L599:
	;
	F_px_THROW_ERROR(m, v1980)
	mBase = m.M
	v1998 = m.ExcPending
	if v1998 != 0 {
		goto L1
	} else {
		goto L605
	}
L600:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_init_work[21])) = int32(_a_F_init_work_53)
	goto L603
L601:
	;
	goto L602
L602:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v1994)+80)) = l1
	goto L604
L603:
	;
	goto L602
L604:
	;
	return
L605:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_initcap(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = F_str_initcap(m, v18, v45, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = F_cstring_to_text(m, v47)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v47)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					return v49
				}
			}
		}
	}
}
func F_initialize_reloptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v486 int32
	_ = v486
	v1 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[0]))
	if v11 != 0 {
		v12 = v1
		for {
			v22 = v12 + int32(1)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v22*int32(28))+uint32(_c_F_initialize_reloptions[0])))
			if v27 != 0 {
				v12 = v22
				continue
			} else {
				break
			}
			break
		}
		v28 = v22
	} else {
		v28 = v1
	}
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[1]))
	if v38 != 0 {
		v39 = v28
		v40 = v1
		for {
			v48 = int32(1)
			v49 = v39 + v48
			v51 = v40 + v48
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v51*int32(36))+uint32(_c_F_initialize_reloptions[1])))
			if v56 != 0 {
				v39 = v49
				v40 = v51
				continue
			} else {
				break
			}
			break
		}
		v57 = v49
	} else {
		v57 = v28
	}
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[2]))
	if v68 != 0 {
		v69 = v57
		v70 = int32(0)
		for {
			v78 = int32(1)
			v79 = v69 + v78
			v81 = v70 + v78
			v86 = *(*int32)(unsafe.Add(mBase, uint32(v81*int32(48))+uint32(_c_F_initialize_reloptions[2])))
			if v86 != 0 {
				v69 = v79
				v70 = v81
				continue
			} else {
				break
			}
			break
		}
		v87 = v79
	} else {
		v87 = v57
	}
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[3]))
	if v98 != 0 {
		v99 = v87
		v100 = int32(0)
		for {
			v108 = int32(1)
			v109 = v99 + v108
			v111 = v100 + v108
			v116 = *(*int32)(unsafe.Add(mBase, uint32(v111*int32(36))+uint32(_c_F_initialize_reloptions[3])))
			if v116 != 0 {
				v99 = v109
				v100 = v111
				continue
			} else {
				break
			}
			break
		}
		v117 = v109
	} else {
		v117 = v87
	}
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[4]))
	if v128 != 0 {
		v129 = v117
		v130 = int32(0)
		for {
			v138 = int32(1)
			v139 = v129 + v138
			v141 = v130 + v138
			v146 = *(*int32)(unsafe.Add(mBase, uint32(v141*int32(44))+uint32(_c_F_initialize_reloptions[4])))
			if v146 != 0 {
				v129 = v139
				v130 = v141
				continue
			} else {
				break
			}
			break
		}
		v147 = v139
	} else {
		v147 = v117
	}
	v156 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[5]))
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[6]))
	if v161 != 0 {
		F_pfree(m, v161)
		mBase = m.M
		v163 = m.ExcPending
		if v163 != 0 {
			return
		} else {
			v166 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[7]))
			v171 = F_MemoryContextAlloc(m, v166, (v158+v147)<<(uint(int32(2))%32)+int32(4))
			mBase = m.M
			v172 = m.ExcPending
			if v172 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[6])) = v171
				v175 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[0]))
				if v175 != 0 {
					v177 = v156
					for {
						v187 = v171 + v177<<(uint(int32(2))%32)
						v189 = v177 * int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189 + int32(_a_F_initialize_reloptions_0)
						*(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_initialize_reloptions[8]))) = int32(0)
						v197 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
						v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
						v199 = F_strlen(m, v198)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v199
						v202 = v177 + int32(1)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_initialize_reloptions[9])))
						if v205 != 0 {
							v177 = v202
							continue
						} else {
							break
						}
						break
					}
					v207 = v202
				} else {
					v207 = v156
				}
				v217 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[1]))
				if v217 != 0 {
					v218 = int32(0)
					v219 = v207
					for {
						v229 = v171 + v219<<(uint(int32(2))%32)
						v231 = v218 * int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v229))) = v231 + int32(_a_F_initialize_reloptions_1)
						v237 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_c_F_initialize_reloptions[10]))) = v237
						v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
						v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
						v241 = F_strlen(m, v240)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v241
						v246 = v219 + v237
						v249 = *(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_c_F_initialize_reloptions[11])))
						if v249 != 0 {
							v218 = v218 + v237
							v219 = v246
							continue
						} else {
							break
						}
						break
					}
					v251 = v246
				} else {
					v251 = v207
				}
				v261 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[2]))
				if v261 != 0 {
					v262 = int32(0)
					v263 = v251
					for {
						v271 = int32(2)
						v273 = v171 + v263<<(uint(v271)%32)
						v275 = v262 * int32(48)
						*(*int32)(unsafe.Add(mBase, uint32(v273))) = v275 + int32(_a_F_initialize_reloptions_2)
						*(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_initialize_reloptions[12]))) = v271
						v283 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
						v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
						v285 = F_strlen(m, v284)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v283)+16)) = v285
						v287 = int32(1)
						v290 = v263 + v287
						v293 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_initialize_reloptions[13])))
						if v293 != 0 {
							v262 = v262 + v287
							v263 = v290
							continue
						} else {
							break
						}
						break
					}
					v295 = v290
				} else {
					v295 = v251
				}
				v305 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[3]))
				if v305 != 0 {
					v306 = int32(0)
					v307 = v295
					for {
						v317 = v171 + v307<<(uint(int32(2))%32)
						v319 = v306 * int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v317))) = v319 + int32(_a_F_initialize_reloptions_3)
						*(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_initialize_reloptions[14]))) = int32(3)
						v327 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
						v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
						v329 = F_strlen(m, v328)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = v329
						v331 = int32(1)
						v334 = v307 + v331
						v337 = *(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_initialize_reloptions[15])))
						if v337 != 0 {
							v306 = v306 + v331
							v307 = v334
							continue
						} else {
							break
						}
						break
					}
					v339 = v334
				} else {
					v339 = v295
				}
				v349 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[4]))
				if v349 != 0 {
					v350 = int32(0)
					v351 = v339
					for {
						v361 = v171 + v351<<(uint(int32(2))%32)
						v363 = v350 * int32(44)
						*(*int32)(unsafe.Add(mBase, uint32(v361))) = v363 + int32(_a_F_initialize_reloptions_4)
						*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_initialize_reloptions[16]))) = int32(4)
						v371 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
						v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
						v373 = F_strlen(m, v372)
						mBase = m.M
						*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = v373
						v375 = int32(1)
						v378 = v351 + v375
						v381 = *(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_initialize_reloptions[17])))
						if v381 != 0 {
							v350 = v350 + v375
							v351 = v378
							continue
						} else {
							break
						}
						break
					}
					v383 = v378
				} else {
					v383 = v339
				}
				v392 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[5]))
				if v392 <= int32(0) {
					v472 = v383
				} else {
					v396 = v392 & int32(3)
					v398 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[18]))
					if base.Ui32(v392) < base.Ui32(int32(4)) {
						v437 = int32(0)
						v438 = v383
						v447 = v437
						v448 = v438
						v449 = int32(0)
						for {
							v456 = int32(2)
							v462 = *(*int32)(unsafe.Add(mBase, uint32(v398+v447<<(uint(v456)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v171+v448<<(uint(v456)%32)))) = v462
							v464 = int32(1)
							v467 = v448 + v464
							v469 = v449 + v464
							if v469 != v396 {
								v447 = v447 + v464
								v448 = v467
								v449 = v469
								continue
							} else {
								break
							}
							break
						}
						v472 = v467
					} else {
						v405 = int32(0)
						v406 = v383
						v413 = v1
						for {
							v414 = int32(2)
							v416 = v171 + v406<<(uint(v414)%32)
							v419 = v398 + v405<<(uint(v414)%32)
							v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
							*(*int32)(unsafe.Add(mBase, uint32(v416))) = v420
							v422 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = v422
							v424 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v416)+8)) = v424
							v426 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v416)+12)) = v426
							v428 = int32(4)
							v429 = v405 + v428
							v431 = v406 + v428
							v433 = v413 + v428
							if v433 != v392&int32(2147483644) {
								v405 = v429
								v406 = v431
								v413 = v433
								continue
							} else {
								break
							}
							break
						}
						if v396 == int32(0) {
							v472 = v431
						} else {
							v437 = v429
							v438 = v431
							v447 = v437
							v448 = v438
							v449 = int32(0)
							for {
								v456 = int32(2)
								v462 = *(*int32)(unsafe.Add(mBase, uint32(v398+v447<<(uint(v456)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v171+v448<<(uint(v456)%32)))) = v462
								v464 = int32(1)
								v467 = v448 + v464
								v469 = v449 + v464
								if v469 != v396 {
									v447 = v447 + v464
									v448 = v467
									v449 = v469
									continue
								} else {
									break
								}
								break
							}
							v472 = v467
						}
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v171+v472<<(uint(int32(2))%32)))) = int32(0)
				v486 = int32(1)
				*(*uint8)(unsafe.Add(mBase, _c_F_initialize_reloptions[19])) = uint8(v486)
				return
			}
		}
	} else {
		v166 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[7]))
		v171 = F_MemoryContextAlloc(m, v166, (v158+v147)<<(uint(int32(2))%32)+int32(4))
		mBase = m.M
		v172 = m.ExcPending
		if v172 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[6])) = v171
			v175 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[0]))
			if v175 != 0 {
				v177 = v156
				for {
					v187 = v171 + v177<<(uint(int32(2))%32)
					v189 = v177 * int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189 + int32(_a_F_initialize_reloptions_0)
					*(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_initialize_reloptions[8]))) = int32(0)
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
					v199 = F_strlen(m, v198)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v197)+16)) = v199
					v202 = v177 + int32(1)
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v189)+uint32(_c_F_initialize_reloptions[9])))
					if v205 != 0 {
						v177 = v202
						continue
					} else {
						break
					}
					break
				}
				v207 = v202
			} else {
				v207 = v156
			}
			v217 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[1]))
			if v217 != 0 {
				v218 = int32(0)
				v219 = v207
				for {
					v229 = v171 + v219<<(uint(int32(2))%32)
					v231 = v218 * int32(36)
					*(*int32)(unsafe.Add(mBase, uint32(v229))) = v231 + int32(_a_F_initialize_reloptions_1)
					v237 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_c_F_initialize_reloptions[10]))) = v237
					v239 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
					v240 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
					v241 = F_strlen(m, v240)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v239)+16)) = v241
					v246 = v219 + v237
					v249 = *(*int32)(unsafe.Add(mBase, uint32(v231)+uint32(_c_F_initialize_reloptions[11])))
					if v249 != 0 {
						v218 = v218 + v237
						v219 = v246
						continue
					} else {
						break
					}
					break
				}
				v251 = v246
			} else {
				v251 = v207
			}
			v261 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[2]))
			if v261 != 0 {
				v262 = int32(0)
				v263 = v251
				for {
					v271 = int32(2)
					v273 = v171 + v263<<(uint(v271)%32)
					v275 = v262 * int32(48)
					*(*int32)(unsafe.Add(mBase, uint32(v273))) = v275 + int32(_a_F_initialize_reloptions_2)
					*(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_initialize_reloptions[12]))) = v271
					v283 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
					v285 = F_strlen(m, v284)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v283)+16)) = v285
					v287 = int32(1)
					v290 = v263 + v287
					v293 = *(*int32)(unsafe.Add(mBase, uint32(v275)+uint32(_c_F_initialize_reloptions[13])))
					if v293 != 0 {
						v262 = v262 + v287
						v263 = v290
						continue
					} else {
						break
					}
					break
				}
				v295 = v290
			} else {
				v295 = v251
			}
			v305 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[3]))
			if v305 != 0 {
				v306 = int32(0)
				v307 = v295
				for {
					v317 = v171 + v307<<(uint(int32(2))%32)
					v319 = v306 * int32(36)
					*(*int32)(unsafe.Add(mBase, uint32(v317))) = v319 + int32(_a_F_initialize_reloptions_3)
					*(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_initialize_reloptions[14]))) = int32(3)
					v327 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
					v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
					v329 = F_strlen(m, v328)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v327)+16)) = v329
					v331 = int32(1)
					v334 = v307 + v331
					v337 = *(*int32)(unsafe.Add(mBase, uint32(v319)+uint32(_c_F_initialize_reloptions[15])))
					if v337 != 0 {
						v306 = v306 + v331
						v307 = v334
						continue
					} else {
						break
					}
					break
				}
				v339 = v334
			} else {
				v339 = v295
			}
			v349 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[4]))
			if v349 != 0 {
				v350 = int32(0)
				v351 = v339
				for {
					v361 = v171 + v351<<(uint(int32(2))%32)
					v363 = v350 * int32(44)
					*(*int32)(unsafe.Add(mBase, uint32(v361))) = v363 + int32(_a_F_initialize_reloptions_4)
					*(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_initialize_reloptions[16]))) = int32(4)
					v371 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
					v372 = *(*int32)(unsafe.Add(mBase, uint32(v371)))
					v373 = F_strlen(m, v372)
					mBase = m.M
					*(*int32)(unsafe.Add(mBase, uint32(v371)+16)) = v373
					v375 = int32(1)
					v378 = v351 + v375
					v381 = *(*int32)(unsafe.Add(mBase, uint32(v363)+uint32(_c_F_initialize_reloptions[17])))
					if v381 != 0 {
						v350 = v350 + v375
						v351 = v378
						continue
					} else {
						break
					}
					break
				}
				v383 = v378
			} else {
				v383 = v339
			}
			v392 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[5]))
			if v392 <= int32(0) {
				v472 = v383
			} else {
				v396 = v392 & int32(3)
				v398 = *(*int32)(unsafe.Add(mBase, _c_F_initialize_reloptions[18]))
				if base.Ui32(v392) < base.Ui32(int32(4)) {
					v437 = int32(0)
					v438 = v383
					v447 = v437
					v448 = v438
					v449 = int32(0)
					for {
						v456 = int32(2)
						v462 = *(*int32)(unsafe.Add(mBase, uint32(v398+v447<<(uint(v456)%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v171+v448<<(uint(v456)%32)))) = v462
						v464 = int32(1)
						v467 = v448 + v464
						v469 = v449 + v464
						if v469 != v396 {
							v447 = v447 + v464
							v448 = v467
							v449 = v469
							continue
						} else {
							break
						}
						break
					}
					v472 = v467
				} else {
					v405 = int32(0)
					v406 = v383
					v413 = v1
					for {
						v414 = int32(2)
						v416 = v171 + v406<<(uint(v414)%32)
						v419 = v398 + v405<<(uint(v414)%32)
						v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
						*(*int32)(unsafe.Add(mBase, uint32(v416))) = v420
						v422 = *(*int32)(unsafe.Add(mBase, uint32(v419)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v416)+4)) = v422
						v424 = *(*int32)(unsafe.Add(mBase, uint32(v419)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v416)+8)) = v424
						v426 = *(*int32)(unsafe.Add(mBase, uint32(v419)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v416)+12)) = v426
						v428 = int32(4)
						v429 = v405 + v428
						v431 = v406 + v428
						v433 = v413 + v428
						if v433 != v392&int32(2147483644) {
							v405 = v429
							v406 = v431
							v413 = v433
							continue
						} else {
							break
						}
						break
					}
					if v396 == int32(0) {
						v472 = v431
					} else {
						v437 = v429
						v438 = v431
						v447 = v437
						v448 = v438
						v449 = int32(0)
						for {
							v456 = int32(2)
							v462 = *(*int32)(unsafe.Add(mBase, uint32(v398+v447<<(uint(v456)%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v171+v448<<(uint(v456)%32)))) = v462
							v464 = int32(1)
							v467 = v448 + v464
							v469 = v449 + v464
							if v469 != v396 {
								v447 = v447 + v464
								v448 = v467
								v449 = v469
								continue
							} else {
								break
							}
							break
						}
						v472 = v467
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v171+v472<<(uint(int32(2))%32)))) = int32(0)
			v486 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_initialize_reloptions[19])) = uint8(v486)
			return
		}
	}
}
func F_innerrel_is_unique_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v555 int32
	_ = v555
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v766 int32
	_ = v766
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v9
	if l5 == v9 {
		v766 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v766
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v25 != 0 {
		v87 = int32(0)
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v93 == int32(0) {
		v766 = v9
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v93 = v87
	goto L3
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	switch v26 {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L6:
	;
	v87 = int32(0)
	goto L4
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59<<(uint(int32(2))%32))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+36))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+120))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+38)))
	if v66 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+108))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v30 <= int32(0) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = int32(0)
	if v33 < v30 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v30
	goto L13
L12:
	;
	v37 = v33
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v40 = v33
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38+v40<<(uint(int32(2))%32))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+101)))
	if v47 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L6
L16:
	;
	v56 = v40 + int32(1)
	if v56 != v37 {
		v40 = v56
		goto L14
	} else {
		goto L20
	}
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+103)))
	if v50 != int32(1) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	if v53 != 0 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v93 = int32(1)
	goto L3
L20:
	;
	goto L15
L21:
	;
	v69 = int32(1)
	if v65 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v65 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64)+100))
	if v70 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+108))
	if v71 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L26
	}
L26:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+36)))
	if v72 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+112))
	if v73 != 0 {
		v87 = v69
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)+144))
	if v74 == int32(0) {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v87 = v69
	goto L4
L30:
	;
	v93 = int32(1)
	goto L3
L31:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+176))
	if v96 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	if v258 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L33:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 <= int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v110 = v9
	goto L35
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v116+v110<<(uint(int32(2))%32))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
	if l7 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L32
L37:
	;
	v241 = v110 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v241 < v242 {
		v110 = v241
		goto L35
	} else {
		goto L69
	}
L38:
	;
	v124 = int32(0)
	if v121 == v124 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v181 = int32(0)
	if base.B2i32(v121 == v181)|base.B2i32(l2 == v181) != 0 {
		v227 = base.B2i32(v121|l2 == v181)
		goto L57
	} else {
		goto L58
	}
L41:
	;
	if v177 == int32(0) {
		goto L37
	} else {
		goto L55
	}
L42:
	;
	v177 = int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	if l2 == int32(0) {
		v170 = v124
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v177 = v170
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v134 < v133 {
		v170 = v124
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v136 = int32(1)
	if v133 <= v136 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v139 = v136
	goto L50
L49:
	;
	v139 = v133
	goto L50
L50:
	;
	v140 = int32(8)
	v145 = int32(0)
	goto L51
L51:
	;
	v152 = v145 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v121+v140+v152)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l2+v140+v152)))
	v159 = v154 & (v156 ^ int32(-1))
	v161 = base.B2i32(v159 == int32(0))
	if v159 != 0 {
		v170 = v161
		goto L45
	} else {
		goto L53
	}
L52:
	;
	v170 = v161
	goto L45
L53:
	;
	v163 = v145 + int32(1)
	if v163 != v139 {
		v145 = v163
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v766 = int32(1)
	goto L1
L56:
	;
	if v227 == int32(0) {
		goto L37
	} else {
		goto L67
	}
L57:
	;
	goto L56
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v195 != v196 {
		v227 = int32(0)
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v198 = int32(1)
	if v195 <= v198 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v201 = v198
	goto L62
L61:
	;
	v201 = v195
	goto L62
L62:
	;
	v202 = int32(8)
	v207 = int32(0)
	goto L63
L63:
	;
	v215 = v207 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v121+v202+v215)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l2+v202+v215)))
	v220 = base.B2i32(v217 == v219)
	if v217 != v219 {
		v227 = v220
		goto L57
	} else {
		goto L65
	}
L64:
	;
	v227 = v220
	goto L57
L65:
	;
	v223 = v207 + int32(1)
	if v223 != v201 {
		v207 = v223
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)))
	if v234 != int32(1) {
		goto L37
	} else {
		goto L68
	}
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v237
	v766 = int32(1)
	goto L1
L69:
	;
	goto L36
L70:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if int32(0) < v359 {
		goto L93
	} else {
		goto L94
	}
L71:
	;
	v261 = int32(0)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v262 <= v261 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v273 = v261
	goto L73
L73:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)+12))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v279+v273<<(uint(int32(2))%32))))
	v284 = int32(0)
	if l2 == v284 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v766 = int32(0)
	goto L1
L75:
	;
	if v337 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v337 = int32(1)
	goto L75
L77:
	;
	goto L78
L78:
	;
	if v283 == int32(0) {
		v330 = v284
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v337 = v330
	goto L75
L80:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	if v294 < v293 {
		v330 = v284
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v296 = int32(1)
	if v293 <= v296 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v299 = v296
	goto L84
L83:
	;
	v299 = v293
	goto L84
L84:
	;
	v300 = int32(8)
	v305 = int32(0)
	goto L85
L85:
	;
	v312 = v305 << (uint(int32(2)) % 32)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l2+v300+v312)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v283+v300+v312)))
	v319 = v314 & (v316 ^ int32(-1))
	v321 = base.B2i32(v319 == int32(0))
	if v319 != 0 {
		v330 = v321
		goto L79
	} else {
		goto L87
	}
L86:
	;
	v330 = v321
	goto L79
L87:
	;
	v323 = v305 + int32(1)
	if v323 != v299 {
		v305 = v323
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v341 = v273 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	if v341 < v342 {
		v273 = v341
		goto L73
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L74
L92:
	;
	goto L70
L93:
	;
	v375 = int32(0)
	v380 = v9
	goto L96
L94:
	;
	v706 = v9
	goto L95
L95:
	;
	if l7 != 0 {
		goto L185
	} else {
		goto L186
	}
L96:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v381+v375<<(uint(int32(2))%32))))
	if int32(1)<<(uint(l4)%32)&int32(174) != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v706 = v688
	goto L95
L98:
	;
	v690 = v375 + int32(1)
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v690 < v691 {
		v375 = v690
		v380 = v688
		goto L96
	} else {
		goto L184
	}
L99:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+8)))
	if v386 != 0 {
		v688 = v380
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385)+9)))
	if v444 != int32(1) {
		v688 = v380
		goto L98
	} else {
		goto L118
	}
L102:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v385)+32))
	v388 = int32(0)
	if v387 == v388 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v441 == int32(0) {
		v688 = v380
		goto L98
	} else {
		goto L117
	}
L104:
	;
	v441 = int32(1)
	goto L103
L105:
	;
	goto L106
L106:
	;
	if l1 == int32(0) {
		v434 = v388
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v441 = v434
	goto L103
L108:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v398 < v397 {
		v434 = v388
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v400 = int32(1)
	if v397 <= v400 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v403 = v400
	goto L112
L111:
	;
	v403 = v397
	goto L112
L112:
	;
	v404 = int32(8)
	v409 = int32(0)
	goto L113
L113:
	;
	v416 = v409 << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v387+v404+v416)))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l1+v404+v416)))
	v423 = v418 & (v420 ^ int32(-1))
	v425 = base.B2i32(v423 == int32(0))
	if v423 != 0 {
		v434 = v425
		goto L107
	} else {
		goto L115
	}
L114:
	;
	v434 = v425
	goto L107
L115:
	;
	v427 = v409 + int32(1)
	if v427 != v403 {
		v409 = v427
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	goto L101
L118:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v385)+96))
	if v447 == int32(0) {
		v688 = v380
		goto L98
	} else {
		goto L119
	}
L119:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v385)+44))
	v452 = int32(0)
	if v451 == v452 {
		goto L123
	} else {
		goto L124
	}
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v385)+120)) = uint8(v681)
	v683 = F_lappend(m, v380, v385)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L182
	} else {
		goto L183
	}
L121:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v385)+44))
	v567 = int32(0)
	if v566 == v567 {
		goto L153
	} else {
		goto L154
	}
L122:
	;
	if v505 == int32(0) {
		goto L121
	} else {
		goto L136
	}
L123:
	;
	v505 = int32(1)
	goto L122
L124:
	;
	goto L125
L125:
	;
	if l2 == int32(0) {
		v498 = v452
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v505 = v498
	goto L122
L127:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v462 < v461 {
		v498 = v452
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v464 = int32(1)
	if v461 <= v464 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v467 = v464
	goto L131
L130:
	;
	v467 = v461
	goto L131
L131:
	;
	v468 = int32(8)
	v473 = int32(0)
	goto L132
L132:
	;
	v480 = v473 << (uint(int32(2)) % 32)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v451+v468+v480)))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l2+v468+v480)))
	v487 = v482 & (v484 ^ int32(-1))
	v489 = base.B2i32(v487 == int32(0))
	if v487 != 0 {
		v498 = v489
		goto L126
	} else {
		goto L134
	}
L133:
	;
	v498 = v489
	goto L126
L134:
	;
	v491 = v473 + int32(1)
	if v491 != v467 {
		v473 = v491
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v385)+48))
	v509 = int32(0)
	if v508 == v509 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	if v562 == int32(0) {
		goto L121
	} else {
		goto L151
	}
L138:
	;
	v562 = int32(1)
	goto L137
L139:
	;
	goto L140
L140:
	;
	if v450 == int32(0) {
		v555 = v509
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v562 = v555
	goto L137
L142:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v519 < v518 {
		v555 = v509
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v521 = int32(1)
	if v518 <= v521 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v524 = v521
	goto L146
L145:
	;
	v524 = v518
	goto L146
L146:
	;
	v525 = int32(8)
	v530 = int32(0)
	goto L147
L147:
	;
	v537 = v530 << (uint(int32(2)) % 32)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v508+v525+v537)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v450+v525+v537)))
	v544 = v539 & (v541 ^ int32(-1))
	v546 = base.B2i32(v544 == int32(0))
	if v544 != 0 {
		v555 = v546
		goto L141
	} else {
		goto L149
	}
L148:
	;
	v555 = v546
	goto L141
L149:
	;
	v548 = v530 + int32(1)
	if v548 != v524 {
		v530 = v548
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v681 = int32(1)
	goto L120
L152:
	;
	if v620 == int32(0) {
		v688 = v380
		goto L98
	} else {
		goto L166
	}
L153:
	;
	v620 = int32(1)
	goto L152
L154:
	;
	goto L155
L155:
	;
	if v450 == int32(0) {
		v613 = v567
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v620 = v613
	goto L152
L157:
	;
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v566)+4))
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v577 < v576 {
		v613 = v567
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v579 = int32(1)
	if v576 <= v579 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v582 = v579
	goto L161
L160:
	;
	v582 = v576
	goto L161
L161:
	;
	v583 = int32(8)
	v588 = int32(0)
	goto L162
L162:
	;
	v595 = v588 << (uint(int32(2)) % 32)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v566+v583+v595)))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v450+v583+v595)))
	v602 = v597 & (v599 ^ int32(-1))
	v604 = base.B2i32(v602 == int32(0))
	if v602 != 0 {
		v613 = v604
		goto L156
	} else {
		goto L164
	}
L163:
	;
	v613 = v604
	goto L156
L164:
	;
	v606 = v588 + int32(1)
	if v606 != v582 {
		v588 = v606
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v385)+48))
	v624 = int32(0)
	if v623 == v624 {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v677 == int32(0) {
		v688 = v380
		goto L98
	} else {
		goto L181
	}
L168:
	;
	v677 = int32(1)
	goto L167
L169:
	;
	goto L170
L170:
	;
	if l2 == int32(0) {
		v670 = v624
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v677 = v670
	goto L167
L172:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v623)+4))
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v634 < v633 {
		v670 = v624
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v636 = int32(1)
	if v633 <= v636 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v639 = v636
	goto L176
L175:
	;
	v639 = v633
	goto L176
L176:
	;
	v640 = int32(8)
	v645 = int32(0)
	goto L177
L177:
	;
	v652 = v645 << (uint(int32(2)) % 32)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v623+v640+v652)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l2+v640+v652)))
	v659 = v654 & (v656 ^ int32(-1))
	v661 = base.B2i32(v659 == int32(0))
	if v659 != 0 {
		v670 = v661
		goto L171
	} else {
		goto L179
	}
L178:
	;
	v670 = v661
	goto L171
L179:
	;
	v663 = v645 + int32(1)
	if v663 != v639 {
		v645 = v663
		goto L177
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	v681 = int32(0)
	goto L120
L182:
	;
	return int32(0)
L183:
	;
	v688 = v683
	goto L98
L184:
	;
	goto L97
L185:
	;
	v710 = v17 + int32(12)
	goto L187
L186:
	;
	v710 = int32(0)
	goto L187
L187:
	;
	v711 = F_rel_is_distinct_for(m, l0, l3, v706, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L182
	} else {
		goto L188
	}
L188:
	;
	if v711 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v713 = int32(_a_F_innerrel_is_unique_ext_0)
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0]))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v716
	v719 = F_palloc0(m, int32(16))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L182
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	if l6 != 0 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v719))) = int32(329)
	v723 = F_bms_copy(m, l2)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L182
	} else {
		goto L193
	}
L193:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v719)+8)) = uint8(base.B2i32(l7 != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v719)+4)) = v723
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v719)+12)) = v729
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l3)+176))
	v732 = F_lappend(m, v731, v719)
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L182
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+176)) = v732
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v714
	v737 = int32(1)
	if l7 == int32(0) {
		v766 = v737
		goto L1
	} else {
		goto L195
	}
L195:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v740
	v766 = v737
	goto L1
L196:
	;
	v745 = int32(_a_F_innerrel_is_unique_ext_0)
	v746 = *(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0]))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l3)+180))
	v751 = F_bms_copy(m, l2)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L182
	} else {
		goto L199
	}
L197:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+368))
	if v742 != 0 {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v766 = int32(0)
	goto L1
L199:
	;
	v753 = F_lappend(m, v750, v751)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L182
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+180)) = v753
	*(*int32)(unsafe.Add(mBase, _c_F_innerrel_is_unique_ext[0])) = v746
	v766 = int32(0)
	goto L1
}
func F_int24div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int24div_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int24div_1), int32(1068), int32(_a_F_int24div_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
		v25 = base.I32_div_s(v24, v3)
		return v25
	}
}
func F_int24ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 != v3)
}
func F_int28div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int28div_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int28div_1), int32(1164), int32(_a_F_int28div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
		v26 = base.I64_div_s(v25, v4)
		v27 = F_Int64GetDatum(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_int2div(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = int32(_a_F_int2div_0)
	v7 = v5 & v6
	if v7 != v6 {
		if v7 != 0 {
			v38 = base.I32_div_s(base.I32_extend16_s(v4), base.I32_extend16_s(v5))
			v41 = v38 << (uint(int32(16)) % 32)
			return v41 >> (uint(int32(16)) % 32)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33816706))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2div_1), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2div_2), int32(988), int32(_a_F_int2div_3))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
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
		if v4&int32(_a_F_int2div_0) == int32(_a_F_int2div_4) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int2div_5), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int2div_2), int32(1004), int32(_a_F_int2div_3))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
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
			v41 = int32(0) - v4<<(uint(int32(16))%32)
			return v41 >> (uint(int32(16)) % 32)
		}
	}
}
func F_int2gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v2)
}
func F_int2or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return base.I32_extend16_s(v2 | v3)
}
func F_int2shr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2 >> (uint(v3) % 32)
}
func F_int2vectorsend(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int42ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 != v3)
}
func F_int48div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int48div_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int48div_1), int32(1022), int32(_a_F_int48div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
		v26 = base.I64_div_s(v25, v4)
		v27 = F_Int64GetDatum(m, v26)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_int48gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 < v4)
}
func F_int4and(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 & v3
}
func F_int4div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v4 + int32(1) {
	case 0:
		if v3 == int32(-2147483648) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_int4div_0), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_int4div_1), int32(888), int32(_a_F_int4div_2))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
			return int32(0) - v3
		}
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4div_3), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4div_1), int32(872), int32(_a_F_int4div_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v30 = base.I32_div_s(v3, v4)
		return v30
	}
}
func F_int4or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 | v3
}
func F_int4range_subdiff(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_Float8GetDatum(m, base.F64_sub(base.F64_convert_i32_s(v2), base.F64_convert_i32_s(v4)))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_int4shr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return v2 >> (uint(v3) % 32)
}
func F_int64_div_fast_to_numeric(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v67 int64
	_ = v67
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int64
	_ = v147
	var v155 int64
	_ = v155
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v205 int64
	_ = v205
	var v208 int64
	_ = v208
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v215 int64
	_ = v215
	var v219 int64
	_ = v219
	var v226 int64
	_ = v226
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 int64
	_ = v310
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v382 int64
	_ = v382
	var v385 int64
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v404 int64
	_ = v404
	var v407 int64
	_ = v407
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	v3 = int32(0)
	v13 = int64(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v13
	v28 = l1 >> (uint(int32(2)) % 32)
	v30 = l1 & int32(3)
	if v30 != 0 {
		v32 = v17 + int32(-48)
		v33 = int64(63)
		v34 = l0 >> (uint(v33) % 64)
		v39 = int64(*(*int32)(unsafe.Add(mBase, uint32((int32(4)-v30)<<(uint(int32(2))%32))+uint32(_c_F_int64_div_fast_to_numeric[0]))))
		v41 = v39 >> (uint(v33) % 64)
		v46 = int64(32)
		v47 = int64(base.Ui64(v39) >> (uint(v46) % 64))
		v49 = int64(base.Ui64(l0) >> (uint(v46) % 64))
		v52 = int64(4294967295)
		v53 = v39 & v52
		v55 = l0 & v52
		v56 = v53 * v55
		v60 = int64(base.Ui64(v56)>>(uint(v46)%64)) + v53*v49
		v67 = v55*v47 + v60&v52
		*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = l0*v41 + v34*v39 + v47*v49 + int64(base.Ui64(v60)>>(uint(v46)%64)) + int64(base.Ui64(v67)>>(uint(v46)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v32))) = v56&v52 | v67<<(uint(v46)%64)
		v78 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
		if v78 != v79>>(uint(int64(63))%64) {
			v87 = int64(32)
			v88 = int64(base.Ui64(l0) >> (uint(v87) % 64))
			v90 = int64(base.Ui64(v39) >> (uint(v87) % 64))
			v93 = int64(4294967295)
			v94 = l0 & v93
			v96 = v39 & v93
			v97 = v94 * v96
			v101 = int64(base.Ui64(v97)>>(uint(v87)%64)) + v94*v90
			v108 = v96*v88 + v101&v93
			*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v39*v34 + v41*l0 + v88*v90 + int64(base.Ui64(v101)>>(uint(v87)%64)) + int64(base.Ui64(v108)>>(uint(v87)%64))
			*(*int64)(unsafe.Add(mBase, uint32(v19))) = v97&v93 | v108<<(uint(v87)%64)
			v119 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
			v120 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
			v121 = m.G0
			v123 = v121 - int32(32)
			m.G0 = v123
			v126 = v17 + int32(-24)
			v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
			if v127 != 0 {
				F_pfree(m, v127)
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					v133 = F_palloc(m, int32(22))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v133
						v136 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v133))) = uint16(v136)
						v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v139 + int32(2)
						if v120 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = int64(16384)
							v147 = int64(0)
							v160 = v147 - v119
							v161 = v147 - (v120 + base.I64_extend_i32_u(base.B2i32(v119 != v147)))
							v164 = v160
							v167 = v139 + int32(22)
							v169 = v136
							v176 = v161
							for {
								v180 = int32(16)
								v181 = v123 + v180
								v184 = m.G0
								v186 = v184 - v180
								m.G0 = v186
								F___udivmodti4(m, v186, v164, v176, int64(10000), int64(0))
								mBase = m.M
								v190 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
								v191 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
								*(*int64)(unsafe.Add(mBase, uint32(v181))) = v191
								*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v190
								m.G0 = v186 + v180
								v197 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
								v198 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
								v199 = int64(55536)
								v200 = int64(0)
								v205 = int64(32)
								v208 = int64(base.Ui64(v197) >> (uint(v205) % 64))
								v211 = int64(4294967295)
								v214 = v197 & v211
								v215 = v199 * v214
								v219 = int64(base.Ui64(v215)>>(uint(v205)%64)) + v199*v208
								v226 = v214*v200 + v219&v211
								*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v197*v200 + v198*v199 + v200*v208 + int64(base.Ui64(v219)>>(uint(v205)%64)) + int64(base.Ui64(v226)>>(uint(v205)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v123))) = v215&v211 | v226<<(uint(v205)%64)
								v238 = v167 - int32(2)
								v239 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
								v240 = v239 + v164
								*(*uint16)(unsafe.Add(mBase, uint32(v238))) = uint16(v240)
								v244 = int64(0)
								v249 = v169 + int32(1)
								if v176 == v244 {
									v250 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v164))
								} else {
									v250 = base.B2i32(v176 != v244)
								}
								if v250 != 0 {
									v164 = v197
									v167 = v238
									v169 = v249
									v176 = v198
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v238
							v257 = v249
							v259 = v169
						} else {
							v155 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v155
							if v119|v120 == v155 {
								v257 = v136
								v259 = v3
							} else {
								v160 = v119
								v161 = v120
								v164 = v160
								v167 = v139 + int32(22)
								v169 = v136
								v176 = v161
								for {
									v180 = int32(16)
									v181 = v123 + v180
									v184 = m.G0
									v186 = v184 - v180
									m.G0 = v186
									F___udivmodti4(m, v186, v164, v176, int64(10000), int64(0))
									mBase = m.M
									v190 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
									v191 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
									*(*int64)(unsafe.Add(mBase, uint32(v181))) = v191
									*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v190
									m.G0 = v186 + v180
									v197 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
									v198 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
									v199 = int64(55536)
									v200 = int64(0)
									v205 = int64(32)
									v208 = int64(base.Ui64(v197) >> (uint(v205) % 64))
									v211 = int64(4294967295)
									v214 = v197 & v211
									v215 = v199 * v214
									v219 = int64(base.Ui64(v215)>>(uint(v205)%64)) + v199*v208
									v226 = v214*v200 + v219&v211
									*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v197*v200 + v198*v199 + v200*v208 + int64(base.Ui64(v219)>>(uint(v205)%64)) + int64(base.Ui64(v226)>>(uint(v205)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v123))) = v215&v211 | v226<<(uint(v205)%64)
									v238 = v167 - int32(2)
									v239 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
									v240 = v239 + v164
									*(*uint16)(unsafe.Add(mBase, uint32(v238))) = uint16(v240)
									v244 = int64(0)
									v249 = v169 + int32(1)
									if v176 == v244 {
										v250 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v164))
									} else {
										v250 = base.B2i32(v176 != v244)
									}
									if v250 != 0 {
										v164 = v197
										v167 = v238
										v169 = v249
										v176 = v198
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v238
								v257 = v249
								v259 = v169
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v259
						*(*int32)(unsafe.Add(mBase, uint32(v126))) = v257
						m.G0 = v123 + int32(32)
						v273 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
						v274 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
						v347 = v273
						v351 = v274
						v434 = v347
						v438 = v351
						v439 = v28 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v438 - v439
						v449 = int32(0)
						if v449 < l1 {
							v452 = l1
						} else {
							v452 = v449
						}
						*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v452
						v457 = F_make_result_opt_error(m, v17+int32(-24), int32(0))
						mBase = m.M
						v458 = m.ExcPending
						if v458 != 0 {
							return int32(0)
						} else {
							if v434 != 0 {
								F_pfree(m, v434)
								mBase = m.M
								v460 = m.ExcPending
								if v460 != 0 {
									return int32(0)
								} else {
									m.G0 = v19 - int32(-64)
									return v457
								}
							} else {
								m.G0 = v19 - int32(-64)
								return v457
							}
						}
					}
				}
			} else {
				v133 = F_palloc(m, int32(22))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v126)+16)) = v133
					v136 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v133))) = uint16(v136)
					v139 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v139 + int32(2)
					if v120 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = int64(16384)
						v147 = int64(0)
						v160 = v147 - v119
						v161 = v147 - (v120 + base.I64_extend_i32_u(base.B2i32(v119 != v147)))
						v164 = v160
						v167 = v139 + int32(22)
						v169 = v136
						v176 = v161
						for {
							v180 = int32(16)
							v181 = v123 + v180
							v184 = m.G0
							v186 = v184 - v180
							m.G0 = v186
							F___udivmodti4(m, v186, v164, v176, int64(10000), int64(0))
							mBase = m.M
							v190 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
							v191 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
							*(*int64)(unsafe.Add(mBase, uint32(v181))) = v191
							*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v190
							m.G0 = v186 + v180
							v197 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
							v198 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
							v199 = int64(55536)
							v200 = int64(0)
							v205 = int64(32)
							v208 = int64(base.Ui64(v197) >> (uint(v205) % 64))
							v211 = int64(4294967295)
							v214 = v197 & v211
							v215 = v199 * v214
							v219 = int64(base.Ui64(v215)>>(uint(v205)%64)) + v199*v208
							v226 = v214*v200 + v219&v211
							*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v197*v200 + v198*v199 + v200*v208 + int64(base.Ui64(v219)>>(uint(v205)%64)) + int64(base.Ui64(v226)>>(uint(v205)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v123))) = v215&v211 | v226<<(uint(v205)%64)
							v238 = v167 - int32(2)
							v239 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
							v240 = v239 + v164
							*(*uint16)(unsafe.Add(mBase, uint32(v238))) = uint16(v240)
							v244 = int64(0)
							v249 = v169 + int32(1)
							if v176 == v244 {
								v250 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v164))
							} else {
								v250 = base.B2i32(v176 != v244)
							}
							if v250 != 0 {
								v164 = v197
								v167 = v238
								v169 = v249
								v176 = v198
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v238
						v257 = v249
						v259 = v169
					} else {
						v155 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v126)+8)) = v155
						if v119|v120 == v155 {
							v257 = v136
							v259 = v3
						} else {
							v160 = v119
							v161 = v120
							v164 = v160
							v167 = v139 + int32(22)
							v169 = v136
							v176 = v161
							for {
								v180 = int32(16)
								v181 = v123 + v180
								v184 = m.G0
								v186 = v184 - v180
								m.G0 = v186
								F___udivmodti4(m, v186, v164, v176, int64(10000), int64(0))
								mBase = m.M
								v190 = *(*int64)(unsafe.Add(mBase, uint32(v186)+8))
								v191 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
								*(*int64)(unsafe.Add(mBase, uint32(v181))) = v191
								*(*int64)(unsafe.Add(mBase, uint32(v181)+8)) = v190
								m.G0 = v186 + v180
								v197 = *(*int64)(unsafe.Add(mBase, uint32(v123)+16))
								v198 = *(*int64)(unsafe.Add(mBase, uint32(v123)+24))
								v199 = int64(55536)
								v200 = int64(0)
								v205 = int64(32)
								v208 = int64(base.Ui64(v197) >> (uint(v205) % 64))
								v211 = int64(4294967295)
								v214 = v197 & v211
								v215 = v199 * v214
								v219 = int64(base.Ui64(v215)>>(uint(v205)%64)) + v199*v208
								v226 = v214*v200 + v219&v211
								*(*int64)(unsafe.Add(mBase, uint32(v123)+8)) = v197*v200 + v198*v199 + v200*v208 + int64(base.Ui64(v219)>>(uint(v205)%64)) + int64(base.Ui64(v226)>>(uint(v205)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v123))) = v215&v211 | v226<<(uint(v205)%64)
								v238 = v167 - int32(2)
								v239 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
								v240 = v239 + v164
								*(*uint16)(unsafe.Add(mBase, uint32(v238))) = uint16(v240)
								v244 = int64(0)
								v249 = v169 + int32(1)
								if v176 == v244 {
									v250 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v164))
								} else {
									v250 = base.B2i32(v176 != v244)
								}
								if v250 != 0 {
									v164 = v197
									v167 = v238
									v169 = v249
									v176 = v198
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v126)+20)) = v238
							v257 = v249
							v259 = v169
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v259
					*(*int32)(unsafe.Add(mBase, uint32(v126))) = v257
					m.G0 = v123 + int32(32)
					v273 = *(*int32)(unsafe.Add(mBase, uint32(v19)+56))
					v274 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
					v347 = v273
					v351 = v274
					v434 = v347
					v438 = v351
					v439 = v28 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v438 - v439
					v449 = int32(0)
					if v449 < l1 {
						v452 = l1
					} else {
						v452 = v449
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v452
					v457 = F_make_result_opt_error(m, v17+int32(-24), int32(0))
					mBase = m.M
					v458 = m.ExcPending
					if v458 != 0 {
						return int32(0)
					} else {
						if v434 != 0 {
							F_pfree(m, v434)
							mBase = m.M
							v460 = m.ExcPending
							if v460 != 0 {
								return int32(0)
							} else {
								m.G0 = v19 - int32(-64)
								return v457
							}
						} else {
							m.G0 = v19 - int32(-64)
							return v457
						}
					}
				}
			}
		} else {
			v276 = F_palloc(m, int32(12))
			mBase = m.M
			v277 = m.ExcPending
			if v277 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v276
				v279 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v276))) = uint16(v279)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v276 + int32(2)
				if v79 < int64(0) {
					*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(16384)
					v295 = int64(0) - v79
					v302 = v276 + int32(12)
					v303 = v279
					v310 = v295
					for {
						v315 = v302 - int32(2)
						v317 = base.I64_div_u_s(v310, int64(10000))
						v320 = v317*int64(55536) + v310
						*(*uint16)(unsafe.Add(mBase, uint32(v315))) = uint16(v320)
						v323 = v303 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v310) {
							v302 = v315
							v303 = v323
							v310 = v317
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v315
					v332 = v323
					v334 = v303
				} else {
					v291 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v291
					if v79 == v291 {
						v332 = v279
						v334 = v3
					} else {
						v295 = v79
						v302 = v276 + int32(12)
						v303 = v279
						v310 = v295
						for {
							v315 = v302 - int32(2)
							v317 = base.I64_div_u_s(v310, int64(10000))
							v320 = v317*int64(55536) + v310
							*(*uint16)(unsafe.Add(mBase, uint32(v315))) = uint16(v320)
							v323 = v303 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v310) {
								v302 = v315
								v303 = v323
								v310 = v317
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v315
						v332 = v323
						v334 = v303
					}
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v332
				v347 = v276
				v351 = v334
				v434 = v347
				v438 = v351
				v439 = v28 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v438 - v439
				v449 = int32(0)
				if v449 < l1 {
					v452 = l1
				} else {
					v452 = v449
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v452
				v457 = F_make_result_opt_error(m, v17+int32(-24), int32(0))
				mBase = m.M
				v458 = m.ExcPending
				if v458 != 0 {
					return int32(0)
				} else {
					if v434 != 0 {
						F_pfree(m, v434)
						mBase = m.M
						v460 = m.ExcPending
						if v460 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 - int32(-64)
							return v457
						}
					} else {
						m.G0 = v19 - int32(-64)
						return v457
					}
				}
			}
		}
	} else {
		v363 = F_palloc(m, int32(12))
		mBase = m.M
		v364 = m.ExcPending
		if v364 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v363
			v366 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v363))) = uint16(v366)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v363 + int32(2)
			if l0 < int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(16384)
				v382 = int64(0) - l0
				v385 = v382
				v389 = v363 + int32(12)
				v390 = v366
				for {
					v402 = v389 - int32(2)
					v404 = base.I64_div_u_s(v385, int64(10000))
					v407 = v404*int64(55536) + v385
					*(*uint16)(unsafe.Add(mBase, uint32(v402))) = uint16(v407)
					v410 = v390 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v385) {
						v385 = v404
						v389 = v402
						v390 = v410
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v402
				v419 = v410
				v421 = v390
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(0)
				if l0 == int64(0) {
					v419 = v366
					v421 = v3
				} else {
					v382 = l0
					v385 = v382
					v389 = v363 + int32(12)
					v390 = v366
					for {
						v402 = v389 - int32(2)
						v404 = base.I64_div_u_s(v385, int64(10000))
						v407 = v404*int64(55536) + v385
						*(*uint16)(unsafe.Add(mBase, uint32(v402))) = uint16(v407)
						v410 = v390 + int32(1)
						if base.Ui64(int64(9999)) < base.Ui64(v385) {
							v385 = v404
							v389 = v402
							v390 = v410
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v402
					v419 = v410
					v421 = v390
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v419
			v434 = v363
			v438 = v421
			v439 = v28
			*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v438 - v439
			v449 = int32(0)
			if v449 < l1 {
				v452 = l1
			} else {
				v452 = v449
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v452
			v457 = F_make_result_opt_error(m, v17+int32(-24), int32(0))
			mBase = m.M
			v458 = m.ExcPending
			if v458 != 0 {
				return int32(0)
			} else {
				if v434 != 0 {
					F_pfree(m, v434)
					mBase = m.M
					v460 = m.ExcPending
					if v460 != 0 {
						return int32(0)
					} else {
						m.G0 = v19 - int32(-64)
						return v457
					}
				} else {
					m.G0 = v19 - int32(-64)
					return v457
				}
			}
		}
	}
}
func F_int82le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v4)
}
func F_int8in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = F_pg_strtoint64_safe(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Int64GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_int8inc_any(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_int8inc(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_int8or(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_Int64GetDatum(m, v3|v5)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_intarray_add_elem(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = F_array_contains_nulls(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v9 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67108994))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_intarray_add_elem_0), int32(0))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intarray_add_elem_1), int32(360), int32(_a_F_intarray_add_elem_2))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
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
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v16 = F_ArrayGetNItemsSafe(m, v13, l0+int32(16))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if v16 < int32(0) {
						v21 = F_construct_empty_array(m, int32(23))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							if v23 == int32(0) {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
								v75 = v21
								v76 = v21 + (v26<<(uint(int32(3))%32)+int32(23))&int32(-8)
							} else {
								v75 = v21
								v76 = v23 + v21
							}
							*(*int32)(unsafe.Add(mBase, uint32(v76+v16<<(uint(int32(2))%32)))) = l1
							return v75
						}
					} else {
						v36 = v16 + int32(1)
						v40 = v36<<(uint(int32(2))%32) + int32(24)
						v41 = F_palloc0(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v36
							*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = int32(23)
							*(*int64)(unsafe.Add(mBase, uint32(v41)+4)) = int64(1)
							*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
							v54 = v41 + int32(24)
							if v16 == int32(0) {
								v75 = v41
								v76 = v54
							} else {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v57 == int32(0) {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v67 = (v60<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								} else {
									v67 = v57
								}
								v69 = v16 << (uint(int32(2)) % 32)
								if v69 == int32(0) {
									v75 = v41
									v76 = v54
								} else {
									base.MemoryCopy(m, v54, l0+v67, v69)
									v75 = v41
									v76 = v54
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v76+v16<<(uint(int32(2))%32)))) = l1
							return v75
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = F_ArrayGetNItemsSafe(m, v13, l0+int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if v16 < int32(0) {
				v21 = F_construct_empty_array(m, int32(23))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
					if v23 == int32(0) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
						v75 = v21
						v76 = v21 + (v26<<(uint(int32(3))%32)+int32(23))&int32(-8)
					} else {
						v75 = v21
						v76 = v23 + v21
					}
					*(*int32)(unsafe.Add(mBase, uint32(v76+v16<<(uint(int32(2))%32)))) = l1
					return v75
				}
			} else {
				v36 = v16 + int32(1)
				v40 = v36<<(uint(int32(2))%32) + int32(24)
				v41 = F_palloc0(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = int32(23)
					*(*int64)(unsafe.Add(mBase, uint32(v41)+4)) = int64(1)
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
					v54 = v41 + int32(24)
					if v16 == int32(0) {
						v75 = v41
						v76 = v54
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v57 == int32(0) {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v67 = (v60<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						} else {
							v67 = v57
						}
						v69 = v16 << (uint(int32(2)) % 32)
						if v69 == int32(0) {
							v75 = v41
							v76 = v54
						} else {
							base.MemoryCopy(m, v54, l0+v67, v69)
							v75 = v41
							v76 = v54
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v76+v16<<(uint(int32(2))%32)))) = l1
					return v75
				}
			}
		}
	}
}
func F_inter_sb(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_box_interpt_lseg(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_interpret_func_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = int32(2281)
		v14 = int32(1)
		v18 = F_LookupFuncName(m, v8, v14, v6+int32(28), v14)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v18 != 0 {
				v20 = F_get_func_rettype(m, v18)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != int32(2281) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								v61 = F_NameListToString(m, v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(_a_F_interpret_func_support_0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v61
									F_errmsg(m, int32(_a_F_interpret_func_support_1), v6+int32(16))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(708), int32(_a_F_interpret_func_support_3))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
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
						v24 = F_superuser(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							if v24 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_interpret_func_support_4), int32(0))
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(718), int32(_a_F_interpret_func_support_3))
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
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
								m.G0 = v6 + int32(32)
								return v18
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(52461700))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v43 = F_func_signature_string(m, v8, int32(1), int32(0), v6+int32(28))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v43
							F_errmsg(m, int32(_a_F_interpret_func_support_5), v6)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_interpret_func_support_2), int32(702), int32(_a_F_interpret_func_support_3))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
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
		}
	}
}
func F_intervaltypmodout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_palloc(m, int32(64))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v10 < int32(0) {
			v18 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v18)
			m.G0 = v8 + int32(48)
			return v12
		} else {
			v21 = int32(base.Ui32(v10) >> (uint(int32(16)) % 32))
			if base.Ui32(v21) <= base.Ui32(int32(3071)) {
				switch v21 - int32(2) {
				case 0:
					v70 = int32(_a_F_intervaltypmodout_0)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 1, 3, 5:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
						F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 2:
					v70 = int32(_a_F_intervaltypmodout_7)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 4:
					v70 = int32(_a_F_intervaltypmodout_8)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				case 6:
					v70 = int32(_a_F_intervaltypmodout_9)
					v71 = int32(_a_F_intervaltypmodout_1)
					v72 = v10 & v71
					if v72 != v71 {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
						v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
						v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(48)
							return v12
						}
					}
				default:
					switch v21 - int32(1024) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_10)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_11)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 == int32(2048) {
							v70 = int32(_a_F_intervaltypmodout_12)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
								F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
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
			} else {
				if base.Ui32(v21) <= base.Ui32(int32(_a_F_intervaltypmodout_13)) {
					switch v21 - int32(3072) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_14)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_15)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 != int32(_a_F_intervaltypmodout_16) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
								F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v70 = int32(_a_F_intervaltypmodout_17)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						}
					}
				} else {
					switch v21 - int32(_a_F_intervaltypmodout_18) {
					case 0:
						v70 = int32(_a_F_intervaltypmodout_19)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					case 1, 2, 3, 4, 5, 6, 7:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
							F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 8:
						v70 = int32(_a_F_intervaltypmodout_20)
						v71 = int32(_a_F_intervaltypmodout_1)
						v72 = v10 & v71
						if v72 != v71 {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
							v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
							v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								m.G0 = v8 + int32(48)
								return v12
							}
						}
					default:
						if v21 == int32(_a_F_intervaltypmodout_21) {
							v70 = int32(_a_F_intervaltypmodout_22)
							v71 = int32(_a_F_intervaltypmodout_1)
							v72 = v10 & v71
							if v72 != v71 {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
								v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
								v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									m.G0 = v8 + int32(48)
									return v12
								}
							}
						} else {
							if v21 != int32(_a_F_intervaltypmodout_23) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
									F_errmsg_internal(m, int32(_a_F_intervaltypmodout_4), v8)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_intervaltypmodout_5), int32(1188), int32(_a_F_intervaltypmodout_6))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v70 = int32(_a_F_intervaltypmodout_24)
								v71 = int32(_a_F_intervaltypmodout_1)
								v72 = v10 & v71
								if v72 != v71 {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v70
									v81 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_2), v8+int32(32))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v12
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v70
									v88 = F_pg_snprintf(m, v12, int32(64), int32(_a_F_intervaltypmodout_3), v8+int32(16))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										m.G0 = v8 + int32(48)
										return v12
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_inv_seek(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v103 int32
	_ = v103
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	switch l2 {
	case 0:
		v131 = l1
		goto L4
	case 1:
		goto L5
	case 2:
		goto L7
	default:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L49
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L17
	} else {
		goto L45
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L17
	} else {
		goto L42
	}
L4:
	;
	if base.Ui64(int64(4398046509057)) <= base.Ui64(v131) {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v129 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v131 = v129 + l1
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L37
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0]))
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v20 = v16
	goto L10
L9:
	;
	v20 = int32(0)
	goto L10
L10:
	;
	if v20 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v23 = int32(_a_F_inv_seek_3)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2]))
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2])) = v27
	if v16 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v52 = v13 + int32(48)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v52, int32(1), int32(3), int32(184), v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L17
	} else {
		goto L23
	}
L14:
	;
	v39 = v19
	goto L16
L15:
	;
	v32 = F_table_open(m, int32(2613), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v39 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	return int64(0)
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0])) = v32
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	v39 = v38
	goto L16
L19:
	;
	v45 = F_index_open(m, int32(2683), int32(3))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[2])) = v24
	goto L13
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1])) = v45
	goto L21
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[0]))
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_inv_seek[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = F_systable_beginscan_ordered(m, v60, v62, v63, int32(1), v52)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	F_systable_endscan_ordered(m, v65)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L17
	} else {
		goto L36
	}
L25:
	;
	v68 = F_systable_getnext_ordered(m, v65, int32(-1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	if v68 == int32(0) {
		v109 = int64(0)
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+20)))
	if v73&int32(1) != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v77 = v72 + v76
	v79 = v77 + int32(8)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
	v82 = v80 & int32(3)
	if v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v83 = F_detoast_attr(m, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L17
	} else {
		goto L32
	}
L30:
	;
	v85 = v79
	goto L31
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v90 = int32(base.Ui32(v86)>>(uint(int32(2))%32)) - int32(4)
	if base.Ui32(v86-int32(_a_F_inv_seek_6)) <= base.Ui32(int32(-8197)) {
		goto L2
	} else {
		goto L33
	}
L32:
	;
	v85 = v83
	goto L31
L33:
	;
	v96 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+4)))
	v99 = base.I64_extend_i32_s(v90) + v96<<(uint(int64(11))%64)
	if v82 == int32(0) {
		v109 = v99
		goto L24
	} else {
		goto L34
	}
L34:
	;
	F_pfree(m, v85)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L35
	}
L35:
	;
	v109 = v99
	goto L24
L36:
	;
	v131 = l1 + v109
	goto L4
L37:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(_a_F_inv_seek_9), v13)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(417), int32(_a_F_inv_seek_2))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L17
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	m.G0 = v13 + int32(96)
	return v131
L42:
	;
	F_errmsg_internal(m, int32(_a_F_inv_seek_4), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(374), int32(_a_F_inv_seek_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L17
	} else {
		goto L46
	}
L46:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v90
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v166
	F_errmsg(m, int32(_a_F_inv_seek_7), v13+int32(32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(153), int32(_a_F_inv_seek_8))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v131
	F_errmsg_internal(m, int32(_a_F_inv_seek_0), v13+int32(16))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_inv_seek_1), int32(430), int32(_a_F_inv_seek_2))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_is_pseudo_constant_clause_relids(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	if l1 == int32(0) {
		v7 = F_contain_volatile_functions_walker(m, l0, int32(0))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				v15 = int32(1)
			} else {
				v15 = int32(0)
			}
			return v15
		}
	} else {
		v15 = int32(0)
		return v15
	}
}
func F_isalpha(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0|int32(32)-int32(97)) < base.Ui32(int32(26)))
}
func F_isatty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v9 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v5+int32(8))
	mBase = m.M
	if v9 == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+8)))
		if v14 == int32(2) {
			v22 = int32(1)
		} else {
			v17 = int32(59)
			*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v17
			v22 = int32(0)
		}
	} else {
		v17 = v9
		*(*int32)(unsafe.Add(mBase, _c_F_isatty[0])) = v17
		v22 = int32(0)
	}
	m.G0 = v5 + int32(32)
	return v22
}
func F_iso_to_win866(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13938(m, l0, int32(_a_F_iso_to_win866_0), int32(20), int32(25))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_iswalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	if base.Ui32(int32(10)) <= base.Ui32(l0-int32(48)) {
		if base.Ui32(l0) <= base.Ui32(int32(_a_F_iswalnum_0)) {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswalnum[0]))))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v14<<(uint(int32(5))%32))+uint32(_c_F_iswalnum[0]))))
			v26 = int32(base.Ui32(v18)>>(uint(l0&int32(7))%32)) & int32(1)
		} else {
			v26 = base.B2i32(base.Ui32(l0) < base.Ui32(int32(_a_F_iswalnum_1)))
		}
		v30 = base.B2i32(v26 != int32(0))
	} else {
		v30 = int32(1)
	}
	return v30
}
func F_iterate_json_values(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = F_palloc0(m, int32(40))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = F_palloc0(m, int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			v22 = F_pg_detoast_datum_packed(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = int32(1)
				v25 = v22 + v24
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				v30 = v28 & v24
				if v30 != 0 {
					v31 = v25
				} else {
					v31 = v22 + int32(4)
				}
				if v28 == int32(1) {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
					if v37 == int32(18) {
						v40 = int32(16)
					} else {
						v40 = int32(0)
					}
					if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v47 = int32(4)
					} else {
						v47 = v40
					}
					v58 = v47
				} else {
					v48 = int32(1)
					if v30 != 0 {
						v58 = int32(base.Ui32(v28)>>(uint(v48)%32)) - v48
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
						v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_iterate_json_values[0]))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				v63 = F_makeJsonLexContextCstringLen(m, v12+int32(12), v31, v58, v61, int32(1))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(1157)
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = int32(1378)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(1379)
					v76 = v12 + int32(12)
					v77 = F_pg_parse_json(m, v76, v15)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return
					} else {
						if v77 != 0 {
							F_json_errsave_error(m, v77, v76, int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_freeJsonLexContext(m, v12+int32(12))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									m.G0 = v12 + int32(80)
									return
								}
							}
						} else {
							F_freeJsonLexContext(m, v12+int32(12))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								m.G0 = v12 + int32(80)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_itmin2interval(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+12)))
	v5 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+16)))
	v8 = v4 + v5*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v8-int64(2147483648)) {
		*(*uint32)(unsafe.Add(mBase, uint32(l1)+12)) = uint32(v8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v14
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(l1))) = v16
	} else {
	}
	return
}
