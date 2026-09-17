package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogCheckpointEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v49 int64
	_ = v49
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v71 int64
	_ = v71
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int64
	_ = v86
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v109 int64
	_ = v109
	var v118 int64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v156 int64
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v202 float64
	_ = v202
	var v203 float64
	_ = v203
	var v208 float64
	_ = v208
	var v213 int64
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v24 = m.G0
	v25 = int32(16)
	v26 = v24 - v25
	m.G0 = v26
	F_gettimeofday(m, v26)
	mBase = m.M
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	v30 = int64(*(*int32)(unsafe.Add(mBase, uint32(v26)+8)))
	m.G0 = v26 + v25
	*(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[0])) = v30 + v29*int64(1000000) - int64(946684800000000)
	v41 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[1]))
	v43 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[2]))
	if v43 <= v41 {
		v61 = int32(0)
	} else {
		v49 = v43 - v41
		if base.B2i32(int64(0) < v41)^base.B2i32(v49 < v43)|base.B2i32(int64(2147483646000) < v49) != 0 {
			v61 = int32(2147483647)
		} else {
			v58 = base.I64_div_s(v49+int64(999), int64(1000))
			v61 = base.I32_wrap_i64(v58)
		}
	}
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[2]))
	v65 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[3]))
	if v65 <= v63 {
		v83 = int32(0)
	} else {
		v71 = v65 - v63
		if base.B2i32(int64(0) < v63)^base.B2i32(v71 < v65)|base.B2i32(int64(2147483646000) < v71) != 0 {
			v83 = int32(2147483647)
		} else {
			v80 = base.I64_div_s(v71+int64(999), int64(1000))
			v83 = base.I32_wrap_i64(v80)
		}
	}
	v84 = int32(_a_F_LogCheckpointEnd_0)
	v86 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[4])) = v86 + base.I64_extend_i32_s(v61)
	v90 = int32(_a_F_LogCheckpointEnd_1)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[5]))
	*(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[5])) = v92 + base.I64_extend_i32_s(v83)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[6])))
	if v97 != int32(1) {
		m.G0 = v18 + int32(112)
		return
	} else {
		v101 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[7]))
		v103 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[0]))
		if v103 <= v101 {
			v121 = int32(0)
		} else {
			v109 = v103 - v101
			if base.B2i32(int64(0) < v101)^base.B2i32(v109 < v103)|base.B2i32(int64(2147483646000) < v109) != 0 {
				v121 = int32(2147483647)
			} else {
				v118 = base.I64_div_s(v109+int64(999), int64(1000))
				v121 = base.I32_wrap_i64(v118)
			}
		}
		v123 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[8]))
		v127 = base.I64_div_u_s(v123+int64(999), int64(1000))
		v129 = *(*int32)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[9]))
		if int32(0) < v129 {
			v133 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[10]))
			v135 = base.I64_div_u_s(v133, base.I64_extend_i32_u(v129))
			v139 = base.I64_div_u_s(v135+int64(999), int64(1000))
			v141 = base.I32_wrap_i64(v139)
		} else {
			v141 = int32(0)
		}
		v144 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v145 = m.ExcPending
		if v145 != 0 {
			return
		} else {
			if v144 == int32(0) {
				m.G0 = v18 + int32(112)
				return
			} else {
				v149 = *(*int32)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[11]))
				v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)+40))
				v151 = *(*int64)(unsafe.Add(mBase, uint32(v149)+32))
				v153 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[12]))
				*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v153
				v156 = *(*int64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[13]))
				*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v156
				v158 = int32(1000)
				v159 = base.I32_div_s(v61, v158)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v159
				v162 = base.I32_div_s(v83, v158)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v162
				v165 = base.I32_div_s(v121, v158)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v165
				v168 = *(*int32)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[9]))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v168
				v170 = base.I32_wrap_i64(v127)
				v172 = base.I32_div_s(v170, v158)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v172
				v175 = base.I32_div_s(v141, v158)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v175
				*(*uint32)(unsafe.Add(mBase, uint32(v18)+88)) = uint32(v151)
				*(*uint32)(unsafe.Add(mBase, uint32(v18)+96)) = uint32(v150)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v61 - v159*v158
				*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v83 - v162*v158
				*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v121 - v165*v158
				*(*int32)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = v170 - v172*v158
				*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v141 - v175*v158
				v202 = *(*float64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[14]))
				v203 = float64(0.0009765625)
				*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = base.I32_trunc_sat_f64_s(base.F64_mul(v202, v203))
				v208 = *(*float64)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[15]))
				*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = base.I32_trunc_sat_f64_s(base.F64_mul(v208, v203))
				v213 = int64(32)
				v214 = int64(base.Ui64(v151) >> (uint(v213) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v18)+84)) = uint32(v214)
				v217 = int64(base.Ui64(v150) >> (uint(v213) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v18)+92)) = uint32(v217)
				v220 = *(*int32)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[16]))
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v220
				v226 = *(*int32)(unsafe.Add(mBase, _c_F_LogCheckpointEnd[17]))
				*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = base.F64_div(base.F64_mul(base.F64_convert_i32_s(v220), float64(100)), base.F64_convert_i32_s(v226))
				if l0 != 0 {
					v232 = int32(_a_F_LogCheckpointEnd_2)
				} else {
					v232 = int32(_a_F_LogCheckpointEnd_3)
				}
				F_errmsg(m, v232, v18)
				mBase = m.M
				v234 = m.ExcPending
				if v234 != 0 {
					return
				} else {
					if l0 != 0 {
						v238 = int32(_a_F_LogCheckpointEnd_4)
					} else {
						v238 = int32(_a_F_LogCheckpointEnd_5)
					}
					F_errfinish(m, int32(_a_F_LogCheckpointEnd_6), v238, int32(_a_F_LogCheckpointEnd_7))
					mBase = m.M
					v241 = m.ExcPending
					if v241 != 0 {
						return
					} else {
						m.G0 = v18 + int32(112)
						return
					}
				}
			}
		}
	}
}
func F_check_log_connections(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v629 int32
	_ = v629
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v706 int32
	_ = v706
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pstrdup(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(48)
	return v720
L2:
	;
	return int32(0)
L3:
	;
	v18 = int32(44)
	v21 = F_SplitIdentifierString(m, v14, v18, v11+v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v21 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[1])) = v27
	goto L8
L6:
	;
	goto L7
L7:
	;
	v44 = int32(0)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v46 == v44 {
		v686 = v44
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_check_log_connections_0)
	v36 = F_format_elog_string(m, int32(_a_F_check_log_connections_1), v11+int32(32))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[2])) = v36
	F_pfree(m, v14)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_list_free(m, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v720 = int32(0)
	goto L1
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[2])) = v706
	F_pfree(m, v14)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L2
	} else {
		goto L232
	}
L13:
	;
	F_pfree(m, v14)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L2
	} else {
		goto L228
	}
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v49 == int32(0) {
		v686 = v44
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = v53
	v58 = int32(_a_F_check_log_connections_2)
	goto L18
L16:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if int32(2) <= v672 {
		goto L223
	} else {
		goto L224
	}
L17:
	;
	if v95 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v61 == v62 {
		v84 = v61
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v95 = int32(0)
	goto L17
L20:
	;
	v86 = int32(1)
	if v84 != 0 {
		v57 = v57 + v86
		v58 = v58 + v86
		goto L18
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v61-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = v61 | int32(32)
	goto L24
L23:
	;
	v72 = v61
	goto L24
L24:
	;
	if base.Ui32((v62-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = v62 | int32(32)
	goto L27
L26:
	;
	v81 = v62
	goto L27
L27:
	;
	if v72 == v81 {
		v84 = v72
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v95 = v72 - v81
	goto L17
L29:
	;
	goto L19
L30:
	;
	v671 = int32(_a_F_check_log_connections_3)
	goto L16
L31:
	;
	goto L32
L32:
	;
	v102 = v53
	v103 = int32(_a_F_check_log_connections_4)
	goto L34
L33:
	;
	if v140 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v106 == v107 {
		v129 = v106
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v140 = int32(0)
	goto L33
L36:
	;
	v131 = int32(1)
	if v129 != 0 {
		v102 = v102 + v131
		v103 = v103 + v131
		goto L34
	} else {
		goto L45
	}
L37:
	;
	if base.Ui32((v106-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v117 = v106 | int32(32)
	goto L40
L39:
	;
	v117 = v106
	goto L40
L40:
	;
	if base.Ui32((v107-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v126 = v107 | int32(32)
	goto L43
L42:
	;
	v126 = v107
	goto L43
L43:
	;
	if v117 == v126 {
		v129 = v117
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v140 = v117 - v126
	goto L33
L45:
	;
	goto L35
L46:
	;
	v671 = int32(_a_F_check_log_connections_5)
	goto L16
L47:
	;
	goto L48
L48:
	;
	v147 = v53
	v148 = int32(_a_F_check_log_connections_6)
	goto L50
L49:
	;
	if v185 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v151 == v152 {
		v174 = v151
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v185 = int32(0)
	goto L49
L52:
	;
	v176 = int32(1)
	if v174 != 0 {
		v147 = v147 + v176
		v148 = v148 + v176
		goto L50
	} else {
		goto L61
	}
L53:
	;
	if base.Ui32((v151-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v162 = v151 | int32(32)
	goto L56
L55:
	;
	v162 = v151
	goto L56
L56:
	;
	if base.Ui32((v152-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v171 = v152 | int32(32)
	goto L59
L58:
	;
	v171 = v152
	goto L59
L59:
	;
	if v162 == v171 {
		v174 = v162
		goto L52
	} else {
		goto L60
	}
L60:
	;
	v185 = v162 - v171
	goto L49
L61:
	;
	goto L51
L62:
	;
	v671 = int32(_a_F_check_log_connections_7)
	goto L16
L63:
	;
	goto L64
L64:
	;
	v192 = v53
	v193 = int32(_a_F_check_log_connections_8)
	goto L66
L65:
	;
	if v230 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L66:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v196 == v197 {
		v219 = v196
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v230 = int32(0)
	goto L65
L68:
	;
	v221 = int32(1)
	if v219 != 0 {
		v192 = v192 + v221
		v193 = v193 + v221
		goto L66
	} else {
		goto L77
	}
L69:
	;
	if base.Ui32((v196-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v207 = v196 | int32(32)
	goto L72
L71:
	;
	v207 = v196
	goto L72
L72:
	;
	if base.Ui32((v197-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v216 = v197 | int32(32)
	goto L75
L74:
	;
	v216 = v197
	goto L75
L75:
	;
	if v207 == v216 {
		v219 = v207
		goto L68
	} else {
		goto L76
	}
L76:
	;
	v230 = v207 - v216
	goto L65
L77:
	;
	goto L67
L78:
	;
	v671 = int32(_a_F_check_log_connections_9)
	goto L16
L79:
	;
	goto L80
L80:
	;
	v237 = v53
	v238 = int32(_a_F_check_log_connections_10)
	goto L82
L81:
	;
	if v275 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L82:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	if v241 == v242 {
		v264 = v241
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v275 = int32(0)
	goto L81
L84:
	;
	v266 = int32(1)
	if v264 != 0 {
		v237 = v237 + v266
		v238 = v238 + v266
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if base.Ui32((v241-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v252 = v241 | int32(32)
	goto L88
L87:
	;
	v252 = v241
	goto L88
L88:
	;
	if base.Ui32((v242-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v261 = v242 | int32(32)
	goto L91
L90:
	;
	v261 = v242
	goto L91
L91:
	;
	if v252 == v261 {
		v264 = v252
		goto L84
	} else {
		goto L92
	}
L92:
	;
	v275 = v252 - v261
	goto L81
L93:
	;
	goto L83
L94:
	;
	v671 = int32(_a_F_check_log_connections_11)
	goto L16
L95:
	;
	goto L96
L96:
	;
	v282 = v53
	v283 = int32(_a_F_check_log_connections_12)
	goto L98
L97:
	;
	if v320 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L98:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283))))
	if v286 == v287 {
		v309 = v286
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v320 = int32(0)
	goto L97
L100:
	;
	v311 = int32(1)
	if v309 != 0 {
		v282 = v282 + v311
		v283 = v283 + v311
		goto L98
	} else {
		goto L109
	}
L101:
	;
	if base.Ui32((v286-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v297 = v286 | int32(32)
	goto L104
L103:
	;
	v297 = v286
	goto L104
L104:
	;
	if base.Ui32((v287-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v306 = v287 | int32(32)
	goto L107
L106:
	;
	v306 = v287
	goto L107
L107:
	;
	if v297 == v306 {
		v309 = v297
		goto L100
	} else {
		goto L108
	}
L108:
	;
	v320 = v297 - v306
	goto L97
L109:
	;
	goto L99
L110:
	;
	v671 = int32(_a_F_check_log_connections_13)
	goto L16
L111:
	;
	goto L112
L112:
	;
	v327 = v53
	v328 = int32(_a_F_check_log_connections_14)
	goto L114
L113:
	;
	if v365 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L114:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v328))))
	if v331 == v332 {
		v354 = v331
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v365 = int32(0)
	goto L113
L116:
	;
	v356 = int32(1)
	if v354 != 0 {
		v327 = v327 + v356
		v328 = v328 + v356
		goto L114
	} else {
		goto L125
	}
L117:
	;
	if base.Ui32((v331-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v342 = v331 | int32(32)
	goto L120
L119:
	;
	v342 = v331
	goto L120
L120:
	;
	if base.Ui32((v332-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v351 = v332 | int32(32)
	goto L123
L122:
	;
	v351 = v332
	goto L123
L123:
	;
	if v342 == v351 {
		v354 = v342
		goto L116
	} else {
		goto L124
	}
L124:
	;
	v365 = v342 - v351
	goto L113
L125:
	;
	goto L115
L126:
	;
	v671 = int32(_a_F_check_log_connections_15)
	goto L16
L127:
	;
	goto L128
L128:
	;
	v372 = v53
	v373 = int32(_a_F_check_log_connections_16)
	goto L130
L129:
	;
	if v410 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L130:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
	if v376 == v377 {
		v399 = v376
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v410 = int32(0)
	goto L129
L132:
	;
	v401 = int32(1)
	if v399 != 0 {
		v372 = v372 + v401
		v373 = v373 + v401
		goto L130
	} else {
		goto L141
	}
L133:
	;
	if base.Ui32((v376-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v387 = v376 | int32(32)
	goto L136
L135:
	;
	v387 = v376
	goto L136
L136:
	;
	if base.Ui32((v377-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v396 = v377 | int32(32)
	goto L139
L138:
	;
	v396 = v377
	goto L139
L139:
	;
	if v387 == v396 {
		v399 = v387
		goto L132
	} else {
		goto L140
	}
L140:
	;
	v410 = v387 - v396
	goto L129
L141:
	;
	goto L131
L142:
	;
	v671 = int32(_a_F_check_log_connections_17)
	goto L16
L143:
	;
	goto L144
L144:
	;
	v414 = int32(0)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v415 <= v414 {
		v686 = v414
		goto L13
	} else {
		goto L145
	}
L145:
	;
	v420 = v414
	v422 = int32(0)
	goto L146
L146:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v427+v422<<(uint(int32(2))%32))))
	v435 = v431
	v436 = int32(_a_F_check_log_connections_18)
	goto L151
L147:
	;
	v662 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[1])) = v662
	goto L221
L148:
	;
	goto L147
L149:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v654)+4))
	v656 = v420 | v655
	v658 = v422 + int32(1)
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v658 < v659 {
		v420 = v656
		v422 = v658
		goto L146
	} else {
		goto L220
	}
L150:
	;
	if v473 == int32(0) {
		v654 = int32(_a_F_check_log_connections_19)
		goto L149
	} else {
		goto L163
	}
L151:
	;
	v439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v436))))
	if v439 == v440 {
		v462 = v439
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v473 = int32(0)
	goto L150
L153:
	;
	v464 = int32(1)
	if v462 != 0 {
		v435 = v435 + v464
		v436 = v436 + v464
		goto L151
	} else {
		goto L162
	}
L154:
	;
	if base.Ui32((v439-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v450 = v439 | int32(32)
	goto L157
L156:
	;
	v450 = v439
	goto L157
L157:
	;
	if base.Ui32((v440-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v459 = v440 | int32(32)
	goto L160
L159:
	;
	v459 = v440
	goto L160
L160:
	;
	if v450 == v459 {
		v462 = v450
		goto L153
	} else {
		goto L161
	}
L161:
	;
	v473 = v450 - v459
	goto L150
L162:
	;
	goto L152
L163:
	;
	v480 = v431
	v481 = int32(_a_F_check_log_connections_20)
	goto L165
L164:
	;
	if v518 == int32(0) {
		v654 = int32(_a_F_check_log_connections_21)
		goto L149
	} else {
		goto L177
	}
L165:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v484 == v485 {
		v507 = v484
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v518 = int32(0)
	goto L164
L167:
	;
	v509 = int32(1)
	if v507 != 0 {
		v480 = v480 + v509
		v481 = v481 + v509
		goto L165
	} else {
		goto L176
	}
L168:
	;
	if base.Ui32((v484-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v495 = v484 | int32(32)
	goto L171
L170:
	;
	v495 = v484
	goto L171
L171:
	;
	if base.Ui32((v485-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v504 = v485 | int32(32)
	goto L174
L173:
	;
	v504 = v485
	goto L174
L174:
	;
	if v495 == v504 {
		v507 = v495
		goto L167
	} else {
		goto L175
	}
L175:
	;
	v518 = v495 - v504
	goto L164
L176:
	;
	goto L166
L177:
	;
	v525 = v431
	v526 = int32(_a_F_check_log_connections_22)
	goto L179
L178:
	;
	if v563 == int32(0) {
		v654 = int32(_a_F_check_log_connections_23)
		goto L149
	} else {
		goto L191
	}
L179:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525))))
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526))))
	if v529 == v530 {
		v552 = v529
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v563 = int32(0)
	goto L178
L181:
	;
	v554 = int32(1)
	if v552 != 0 {
		v525 = v525 + v554
		v526 = v526 + v554
		goto L179
	} else {
		goto L190
	}
L182:
	;
	if base.Ui32((v529-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v540 = v529 | int32(32)
	goto L185
L184:
	;
	v540 = v529
	goto L185
L185:
	;
	if base.Ui32((v530-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v549 = v530 | int32(32)
	goto L188
L187:
	;
	v549 = v530
	goto L188
L188:
	;
	if v540 == v549 {
		v552 = v540
		goto L181
	} else {
		goto L189
	}
L189:
	;
	v563 = v540 - v549
	goto L178
L190:
	;
	goto L180
L191:
	;
	v570 = v431
	v571 = int32(_a_F_check_log_connections_24)
	goto L193
L192:
	;
	if v608 == int32(0) {
		v654 = int32(_a_F_check_log_connections_25)
		goto L149
	} else {
		goto L205
	}
L193:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v571))))
	if v574 == v575 {
		v597 = v574
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v608 = int32(0)
	goto L192
L195:
	;
	v599 = int32(1)
	if v597 != 0 {
		v570 = v570 + v599
		v571 = v571 + v599
		goto L193
	} else {
		goto L204
	}
L196:
	;
	if base.Ui32((v574-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v585 = v574 | int32(32)
	goto L199
L198:
	;
	v585 = v574
	goto L199
L199:
	;
	if base.Ui32((v575-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v594 = v575 | int32(32)
	goto L202
L201:
	;
	v594 = v575
	goto L202
L202:
	;
	if v585 == v594 {
		v597 = v585
		goto L195
	} else {
		goto L203
	}
L203:
	;
	v608 = v585 - v594
	goto L192
L204:
	;
	goto L194
L205:
	;
	v614 = v431
	v615 = int32(_a_F_check_log_connections_26)
	goto L207
L206:
	;
	if v652 != 0 {
		goto L148
	} else {
		goto L219
	}
L207:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614))))
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v615))))
	if v618 == v619 {
		v641 = v618
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v652 = int32(0)
	goto L206
L209:
	;
	v643 = int32(1)
	if v641 != 0 {
		v614 = v614 + v643
		v615 = v615 + v643
		goto L207
	} else {
		goto L218
	}
L210:
	;
	if base.Ui32((v618-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v629 = v618 | int32(32)
	goto L213
L212:
	;
	v629 = v618
	goto L213
L213:
	;
	if base.Ui32((v619-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v638 = v619 | int32(32)
	goto L216
L215:
	;
	v638 = v619
	goto L216
L216:
	;
	if v629 == v638 {
		v641 = v629
		goto L209
	} else {
		goto L217
	}
L217:
	;
	v652 = v629 - v638
	goto L206
L218:
	;
	goto L208
L219:
	;
	v654 = int32(_a_F_check_log_connections_27)
	goto L149
L220:
	;
	v686 = v656
	goto L13
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v431
	v669 = F_format_elog_string(m, int32(_a_F_check_log_connections_28), v11+int32(16))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L2
	} else {
		goto L222
	}
L222:
	;
	v706 = v669
	goto L12
L223:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_log_connections[1])) = v676
	goto L226
L224:
	;
	goto L225
L225:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	v686 = v683
	goto L13
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v53
	v681 = F_format_elog_string(m, int32(_a_F_check_log_connections_29), v11)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L2
	} else {
		goto L227
	}
L227:
	;
	v706 = v681
	goto L12
L228:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_list_free(m, v694)
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L2
	} else {
		goto L229
	}
L229:
	;
	v698 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L2
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v698
	if v698 == int32(0) {
		v720 = v44
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v698))) = v686
	v720 = int32(1)
	goto L1
L232:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	F_list_free(m, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L2
	} else {
		goto L233
	}
L233:
	;
	v720 = v44
	goto L1
}
func F_log_heap_new_cid(m *base.Module, l0 int32, l1 int32) {
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v10 = F_GetTopTransactionId(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v13
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7)+28)) = v15
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+44)) = uint16(v17)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+40)) = v19
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+20)))
		if v21&int32(32) != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
			if v26&int32(32) != 0 {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_new_cid[0]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v25<<(uint(int32(3))%32))))
				v35 = v34
			} else {
				v35 = v25
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
			if v39&int32(32) != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_new_cid[0]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(3))%32))+4))
				v48 = v47
			} else {
				v48 = v38
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v48
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v74 = v50
		} else {
			v56 = int32(0)
			if base.B2i32(v21&int32(2048)|v21&int32(128) == v56)&base.B2i32(v21&int32(_a_F_log_heap_new_cid_0) != int32(64)) == v56 {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v65
				v71 = int32(-1)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(-1)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v71 = v70
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v71
			v74 = int32(-1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v74
		F_XLogBeginInsert(m)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			F_XLogRegisterData(m, v7+int32(12), int32(34))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return
			} else {
				v85 = F_XLogInsert(m, int32(9), int32(112))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			}
		}
	}
}
func F_log_newpage_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_buffer[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_buffer[1]))
		v26 = v20 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v28 = v7 + int32(20)
	if l0 < int32(0) {
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_buffer[2]))
		v50 = v37 + (l0^int32(-1))<<(uint(int32(6))%32)
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_buffer[3]))
		v50 = v44 + l0<<(uint(int32(6))%32) + int32(-64)
	}
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+8)) = v52
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v51
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(16)))) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v7+int32(12)))) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_log_newpage(m, v28, v59, v60, v26, l1)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		return
	} else {
		m.G0 = v7 + int32(32)
		return
	}
}
func F_show_log_file_mode(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13992(m, int32(_a_F_show_log_file_mode_0), int32(_a_F_show_log_file_mode_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
