package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogCheckpointEnd(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v71 int64
	_ = v71
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v105 int32
	_ = v105
	var v108 int64
	_ = v108
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v200 float64
	_ = v200
	var v202 float64
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v225 int64
	_ = v225
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v25 = m.G0
	v26 = int32(16)
	v27 = v25 - v26
	m.G0 = v27
	F___gettimeofday(m, v27)
	mBase = m.M
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
	v31 = int64(*(*int32)(unsafe.Add(mBase, uint32(v27)+8)))
	m.G0 = v27 + v26
	*(*int64)(unsafe.Add(mBase, _consts[231])) = v31 + v30*int64(1000000) - int64(946684800000000)
	v42 = *(*int64)(unsafe.Add(mBase, _consts[232]))
	v44 = *(*int64)(unsafe.Add(mBase, _consts[233]))
	if v44 <= v42 {
		v61 = int32(0)
	} else {
		v47 = int32(2147483647)
		v50 = v44 - v42
		if base.B2i32(int64(0) < v42)^base.B2i32(v50 < v44) != 0 {
			v61 = v47
		} else {
			if int64(2147483646000) < v50 {
				v61 = v47
			} else {
				v58 = base.I64_div_s(v50+int64(999), int64(1000))
				v61 = base.I32_wrap_i64(v58)
			}
		}
	}
	v63 = *(*int64)(unsafe.Add(mBase, _consts[233]))
	v65 = *(*int64)(unsafe.Add(mBase, _consts[234]))
	if v65 <= v63 {
		v82 = int32(0)
	} else {
		v68 = int32(2147483647)
		v71 = v65 - v63
		if base.B2i32(int64(0) < v63)^base.B2i32(v71 < v65) != 0 {
			v82 = v68
		} else {
			if int64(2147483646000) < v71 {
				v82 = v68
			} else {
				v79 = base.I64_div_s(v71+int64(999), int64(1000))
				v82 = base.I32_wrap_i64(v79)
			}
		}
	}
	v83 = int32(4495168)
	v85 = *(*int64)(unsafe.Add(mBase, _consts[235]))
	*(*int64)(unsafe.Add(mBase, _consts[235])) = v85 + base.I64_extend_i32_s(v61)
	v89 = int32(4495176)
	v91 = *(*int64)(unsafe.Add(mBase, _consts[236]))
	*(*int64)(unsafe.Add(mBase, _consts[236])) = v91 + base.I64_extend_i32_s(v82)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _consts[237])))
	if v96 != int32(1) {
		m.G0 = v19 + int32(112)
		return
	} else {
		v100 = *(*int64)(unsafe.Add(mBase, _consts[238]))
		v102 = *(*int64)(unsafe.Add(mBase, _consts[231]))
		if v102 <= v100 {
			v119 = int32(0)
		} else {
			v105 = int32(2147483647)
			v108 = v102 - v100
			if base.B2i32(int64(0) < v100)^base.B2i32(v108 < v102) != 0 {
				v119 = v105
			} else {
				if int64(2147483646000) < v108 {
					v119 = v105
				} else {
					v116 = base.I64_div_s(v108+int64(999), int64(1000))
					v119 = base.I32_wrap_i64(v116)
				}
			}
		}
		v121 = *(*int64)(unsafe.Add(mBase, _consts[239]))
		v125 = base.I64_div_u_s(v121+int64(999), int64(1000))
		v127 = *(*int32)(unsafe.Add(mBase, _consts[240]))
		if int32(0) < v127 {
			v131 = *(*int64)(unsafe.Add(mBase, _consts[241]))
			v133 = base.I64_div_u_s(v131, base.I64_extend_i32_u(v127))
			v137 = base.I64_div_u_s(v133+int64(999), int64(1000))
			v139 = base.I32_wrap_i64(v137)
		} else {
			v139 = int32(0)
		}
		v142 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return
		} else {
			if v142 == int32(0) {
				m.G0 = v19 + int32(112)
				return
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, _consts[109]))
				v148 = *(*int64)(unsafe.Add(mBase, uint32(v147)+40))
				v149 = *(*int64)(unsafe.Add(mBase, uint32(v147)+32))
				v151 = *(*int64)(unsafe.Add(mBase, _consts[242]))
				*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v151
				v154 = *(*int64)(unsafe.Add(mBase, _consts[243]))
				*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v154
				v156 = int32(1000)
				v157 = base.I32_div_s(v61, v156)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v157
				v160 = base.I32_div_s(v82, v156)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v160
				v163 = base.I32_div_s(v119, v156)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v163
				v166 = *(*int32)(unsafe.Add(mBase, _consts[240]))
				*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v166
				v168 = base.I32_wrap_i64(v125)
				v170 = base.I32_div_s(v168, v156)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v170
				v173 = base.I32_div_s(v139, v156)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v173
				*(*uint32)(unsafe.Add(mBase, uint32(v19)+88)) = uint32(v149)
				*(*uint32)(unsafe.Add(mBase, uint32(v19)+96)) = uint32(v148)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v61 - v157*v156
				*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v82 - v160*v156
				*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v119 - v163*v156
				*(*int32)(unsafe.Add(mBase, uint32(v19-int32(-64)))) = v168 - v170*v156
				*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v139 - v173*v156
				v200 = *(*float64)(unsafe.Add(mBase, _consts[244]))
				v202 = base.F64_mul(v200, float64(0.0009765625))
				if base.F64_lt(base.F64_abs(v202), float64(2.147483648e+09)) != 0 {
					v206 = base.I32_trunc_f64_s(v202)
					v208 = v206
				} else {
					v208 = int32(-2147483648)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v208
				v211 = *(*float64)(unsafe.Add(mBase, _consts[245]))
				v213 = base.F64_mul(v211, float64(0.0009765625))
				if base.F64_lt(base.F64_abs(v213), float64(2.147483648e+09)) != 0 {
					v217 = base.I32_trunc_f64_s(v213)
					v219 = v217
				} else {
					v219 = int32(-2147483648)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v219
				v221 = int64(32)
				v222 = int64(base.Ui64(v149) >> (uint(v221) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v19)+84)) = uint32(v222)
				v225 = int64(base.Ui64(v148) >> (uint(v221) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v19)+92)) = uint32(v225)
				v228 = *(*int32)(unsafe.Add(mBase, _consts[246]))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v228
				v234 = *(*int32)(unsafe.Add(mBase, _consts[136]))
				*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = base.F64_div(base.F64_mul(base.F64_convert_i32_s(v228), float64(100)), base.F64_convert_i32_s(v234))
				if l0 != 0 {
					v240 = int32(510615)
				} else {
					v240 = int32(510901)
				}
				F_errmsg(m, v240, v19)
				mBase = m.M
				v242 = m.ExcPending
				if v242 != 0 {
					return
				} else {
					if l0 != 0 {
						v246 = int32(6792)
					} else {
						v246 = int32(6816)
					}
					F_errfinish(m, int32(498262), v246, int32(429497))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return
					} else {
						m.G0 = v19 + int32(112)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v345 int32
	_ = v345
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v455 int32
	_ = v455
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v545 int32
	_ = v545
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v634 int32
	_ = v634
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_pstrdup(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(48)
	return v720
L2:
	;
	return int32(0)
L3:
	;
	v19 = int32(44)
	v22 = F_SplitIdentifierString(m, v15, v19, v12+v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v22 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v28
	goto L8
L6:
	;
	goto L7
L7:
	;
	v45 = int32(1)
	v46 = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v47 == v46 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(141552)
	v37 = F_format_elog_string(m, int32(665384), v12+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[826])) = v37
	F_pfree(m, v15)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_list_free(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
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
	F_pfree(m, v15)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L2
	} else {
		goto L231
	}
L13:
	;
	v698 = int32(0)
	v704 = v45
	goto L12
L14:
	;
	goto L15
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v698 = int32(0)
	v704 = v45
	goto L12
L17:
	;
	goto L18
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v60 = v56
	v61 = int32(339094)
	goto L21
L19:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if int32(2) <= v680 {
		goto L226
	} else {
		goto L227
	}
L20:
	;
	if v98 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v64 == v65 {
		v87 = v64
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v98 = int32(0)
	goto L20
L23:
	;
	v89 = int32(1)
	if v87 != 0 {
		v60 = v60 + v89
		v61 = v61 + v89
		goto L21
	} else {
		goto L32
	}
L24:
	;
	if base.Ui32((v64-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = v64 | int32(32)
	goto L27
L26:
	;
	v75 = v64
	goto L27
L27:
	;
	if base.Ui32((v65-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v84 = v65 | int32(32)
	goto L30
L29:
	;
	v84 = v65
	goto L30
L30:
	;
	if v75 == v84 {
		v87 = v75
		goto L23
	} else {
		goto L31
	}
L31:
	;
	v98 = v75 - v84
	goto L20
L32:
	;
	goto L22
L33:
	;
	v679 = int32(1638576)
	goto L19
L34:
	;
	goto L35
L35:
	;
	v105 = v56
	v106 = int32(361486)
	goto L37
L36:
	;
	if v143 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L37:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	if v109 == v110 {
		v132 = v109
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v143 = int32(0)
	goto L36
L39:
	;
	v134 = int32(1)
	if v132 != 0 {
		v105 = v105 + v134
		v106 = v106 + v134
		goto L37
	} else {
		goto L48
	}
L40:
	;
	if base.Ui32((v109-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v120 = v109 | int32(32)
	goto L43
L42:
	;
	v120 = v109
	goto L43
L43:
	;
	if base.Ui32((v110-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v129 = v110 | int32(32)
	goto L46
L45:
	;
	v129 = v110
	goto L46
L46:
	;
	if v120 == v129 {
		v132 = v120
		goto L39
	} else {
		goto L47
	}
L47:
	;
	v143 = v120 - v129
	goto L36
L48:
	;
	goto L38
L49:
	;
	v679 = int32(1638588)
	goto L19
L50:
	;
	goto L51
L51:
	;
	v150 = v56
	v151 = int32(241050)
	goto L53
L52:
	;
	if v188 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L53:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150))))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v154 == v155 {
		v177 = v154
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v188 = int32(0)
	goto L52
L55:
	;
	v179 = int32(1)
	if v177 != 0 {
		v150 = v150 + v179
		v151 = v151 + v179
		goto L53
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v154-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v165 = v154 | int32(32)
	goto L59
L58:
	;
	v165 = v154
	goto L59
L59:
	;
	if base.Ui32((v155-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v174 = v155 | int32(32)
	goto L62
L61:
	;
	v174 = v155
	goto L62
L62:
	;
	if v165 == v174 {
		v177 = v165
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v188 = v165 - v174
	goto L52
L64:
	;
	goto L54
L65:
	;
	v679 = int32(1638600)
	goto L19
L66:
	;
	goto L67
L67:
	;
	v195 = v56
	v196 = int32(570745)
	goto L69
L68:
	;
	if v233 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L69:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	if v199 == v200 {
		v222 = v199
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v233 = int32(0)
	goto L68
L71:
	;
	v224 = int32(1)
	if v222 != 0 {
		v195 = v195 + v224
		v196 = v196 + v224
		goto L69
	} else {
		goto L80
	}
L72:
	;
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v210 = v199 | int32(32)
	goto L75
L74:
	;
	v210 = v199
	goto L75
L75:
	;
	if base.Ui32((v200-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v219 = v200 | int32(32)
	goto L78
L77:
	;
	v219 = v200
	goto L78
L78:
	;
	if v210 == v219 {
		v222 = v210
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v233 = v210 - v219
	goto L68
L80:
	;
	goto L70
L81:
	;
	v679 = int32(1638612)
	goto L19
L82:
	;
	goto L83
L83:
	;
	v240 = v56
	v241 = int32(273222)
	goto L85
L84:
	;
	if v278 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L85:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v244 == v245 {
		v267 = v244
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v278 = int32(0)
	goto L84
L87:
	;
	v269 = int32(1)
	if v267 != 0 {
		v240 = v240 + v269
		v241 = v241 + v269
		goto L85
	} else {
		goto L96
	}
L88:
	;
	if base.Ui32((v244-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v255 = v244 | int32(32)
	goto L91
L90:
	;
	v255 = v244
	goto L91
L91:
	;
	if base.Ui32((v245-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v264 = v245 | int32(32)
	goto L94
L93:
	;
	v264 = v245
	goto L94
L94:
	;
	if v255 == v264 {
		v267 = v255
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v278 = v255 - v264
	goto L84
L96:
	;
	goto L86
L97:
	;
	v679 = int32(1638624)
	goto L19
L98:
	;
	goto L99
L99:
	;
	v285 = v56
	v286 = int32(344470)
	goto L101
L100:
	;
	if v323 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L101:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	if v289 == v290 {
		v312 = v289
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v323 = int32(0)
	goto L100
L103:
	;
	v314 = int32(1)
	if v312 != 0 {
		v285 = v285 + v314
		v286 = v286 + v314
		goto L101
	} else {
		goto L112
	}
L104:
	;
	if base.Ui32((v289-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v300 = v289 | int32(32)
	goto L107
L106:
	;
	v300 = v289
	goto L107
L107:
	;
	if base.Ui32((v290-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v309 = v290 | int32(32)
	goto L110
L109:
	;
	v309 = v290
	goto L110
L110:
	;
	if v300 == v309 {
		v312 = v300
		goto L103
	} else {
		goto L111
	}
L111:
	;
	v323 = v300 - v309
	goto L100
L112:
	;
	goto L102
L113:
	;
	v679 = int32(1638636)
	goto L19
L114:
	;
	goto L115
L115:
	;
	v330 = v56
	v331 = int32(157316)
	goto L117
L116:
	;
	if v368 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L117:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330))))
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331))))
	if v334 == v335 {
		v357 = v334
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v368 = int32(0)
	goto L116
L119:
	;
	v359 = int32(1)
	if v357 != 0 {
		v330 = v330 + v359
		v331 = v331 + v359
		goto L117
	} else {
		goto L128
	}
L120:
	;
	if base.Ui32((v334-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v345 = v334 | int32(32)
	goto L123
L122:
	;
	v345 = v334
	goto L123
L123:
	;
	if base.Ui32((v335-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v354 = v335 | int32(32)
	goto L126
L125:
	;
	v354 = v335
	goto L126
L126:
	;
	if v345 == v354 {
		v357 = v345
		goto L119
	} else {
		goto L127
	}
L127:
	;
	v368 = v345 - v354
	goto L116
L128:
	;
	goto L118
L129:
	;
	v679 = int32(1638648)
	goto L19
L130:
	;
	goto L131
L131:
	;
	v375 = v56
	v376 = int32(562633)
	goto L133
L132:
	;
	if v413 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L133:
	;
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375))))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376))))
	if v379 == v380 {
		v402 = v379
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v413 = int32(0)
	goto L132
L135:
	;
	v404 = int32(1)
	if v402 != 0 {
		v375 = v375 + v404
		v376 = v376 + v404
		goto L133
	} else {
		goto L144
	}
L136:
	;
	if base.Ui32((v379-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v390 = v379 | int32(32)
	goto L139
L138:
	;
	v390 = v379
	goto L139
L139:
	;
	if base.Ui32((v380-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v399 = v380 | int32(32)
	goto L142
L141:
	;
	v399 = v380
	goto L142
L142:
	;
	if v390 == v399 {
		v402 = v390
		goto L135
	} else {
		goto L143
	}
L143:
	;
	v413 = v390 - v399
	goto L132
L144:
	;
	goto L134
L145:
	;
	v679 = int32(1638660)
	goto L19
L146:
	;
	goto L147
L147:
	;
	v417 = int32(0)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v418 <= v417 {
		v698 = v417
		v704 = v45
		goto L12
	} else {
		goto L148
	}
L148:
	;
	v424 = v417
	v425 = int32(0)
	goto L149
L149:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v432+v425<<(uint(int32(2))%32))))
	v440 = v436
	v441 = int32(84182)
	goto L154
L150:
	;
	v668 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v668
	goto L224
L151:
	;
	goto L150
L152:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v659)+4))
	v661 = v660 | v424
	v663 = v425 + int32(1)
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	if v663 < v664 {
		v424 = v661
		v425 = v663
		goto L149
	} else {
		goto L223
	}
L153:
	;
	if v478 == int32(0) {
		v659 = int32(1638672)
		goto L152
	} else {
		goto L166
	}
L154:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v440))))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v441))))
	if v444 == v445 {
		v467 = v444
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v478 = int32(0)
	goto L153
L156:
	;
	v469 = int32(1)
	if v467 != 0 {
		v440 = v440 + v469
		v441 = v441 + v469
		goto L154
	} else {
		goto L165
	}
L157:
	;
	if base.Ui32((v444-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v455 = v444 | int32(32)
	goto L160
L159:
	;
	v455 = v444
	goto L160
L160:
	;
	if base.Ui32((v445-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v464 = v445 | int32(32)
	goto L163
L162:
	;
	v464 = v445
	goto L163
L163:
	;
	if v455 == v464 {
		v467 = v455
		goto L156
	} else {
		goto L164
	}
L164:
	;
	v478 = v455 - v464
	goto L153
L165:
	;
	goto L155
L166:
	;
	v485 = v436
	v486 = int32(265897)
	goto L168
L167:
	;
	if v523 == int32(0) {
		v659 = int32(1638684)
		goto L152
	} else {
		goto L180
	}
L168:
	;
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485))))
	v490 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v486))))
	if v489 == v490 {
		v512 = v489
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v523 = int32(0)
	goto L167
L170:
	;
	v514 = int32(1)
	if v512 != 0 {
		v485 = v485 + v514
		v486 = v486 + v514
		goto L168
	} else {
		goto L179
	}
L171:
	;
	if base.Ui32((v489-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v500 = v489 | int32(32)
	goto L174
L173:
	;
	v500 = v489
	goto L174
L174:
	;
	if base.Ui32((v490-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v509 = v490 | int32(32)
	goto L177
L176:
	;
	v509 = v490
	goto L177
L177:
	;
	if v500 == v509 {
		v512 = v500
		goto L170
	} else {
		goto L178
	}
L178:
	;
	v523 = v500 - v509
	goto L167
L179:
	;
	goto L169
L180:
	;
	v530 = v436
	v531 = int32(258650)
	goto L182
L181:
	;
	if v568 == int32(0) {
		v659 = int32(1638696)
		goto L152
	} else {
		goto L194
	}
L182:
	;
	v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v530))))
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v531))))
	if v534 == v535 {
		v557 = v534
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v568 = int32(0)
	goto L181
L184:
	;
	v559 = int32(1)
	if v557 != 0 {
		v530 = v530 + v559
		v531 = v531 + v559
		goto L182
	} else {
		goto L193
	}
L185:
	;
	if base.Ui32((v534-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v545 = v534 | int32(32)
	goto L188
L187:
	;
	v545 = v534
	goto L188
L188:
	;
	if base.Ui32((v535-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v554 = v535 | int32(32)
	goto L191
L190:
	;
	v554 = v535
	goto L191
L191:
	;
	if v545 == v554 {
		v557 = v545
		goto L184
	} else {
		goto L192
	}
L192:
	;
	v568 = v545 - v554
	goto L181
L193:
	;
	goto L183
L194:
	;
	v575 = v436
	v576 = int32(142971)
	goto L196
L195:
	;
	if v613 == int32(0) {
		v659 = int32(1638708)
		goto L152
	} else {
		goto L208
	}
L196:
	;
	v579 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575))))
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v576))))
	if v579 == v580 {
		v602 = v579
		goto L198
	} else {
		goto L199
	}
L197:
	;
	v613 = int32(0)
	goto L195
L198:
	;
	v604 = int32(1)
	if v602 != 0 {
		v575 = v575 + v604
		v576 = v576 + v604
		goto L196
	} else {
		goto L207
	}
L199:
	;
	if base.Ui32((v579-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v590 = v579 | int32(32)
	goto L202
L201:
	;
	v590 = v579
	goto L202
L202:
	;
	if base.Ui32((v580-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v599 = v580 | int32(32)
	goto L205
L204:
	;
	v599 = v580
	goto L205
L205:
	;
	if v590 == v599 {
		v602 = v590
		goto L198
	} else {
		goto L206
	}
L206:
	;
	v613 = v590 - v599
	goto L195
L207:
	;
	goto L197
L208:
	;
	v619 = v436
	v620 = int32(305463)
	goto L210
L209:
	;
	if v657 != 0 {
		goto L151
	} else {
		goto L222
	}
L210:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619))))
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v620))))
	if v623 == v624 {
		v646 = v623
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v657 = int32(0)
	goto L209
L212:
	;
	v648 = int32(1)
	if v646 != 0 {
		v619 = v619 + v648
		v620 = v620 + v648
		goto L210
	} else {
		goto L221
	}
L213:
	;
	if base.Ui32((v623-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v634 = v623 | int32(32)
	goto L216
L215:
	;
	v634 = v623
	goto L216
L216:
	;
	if base.Ui32((v624-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v643 = v624 | int32(32)
	goto L219
L218:
	;
	v643 = v624
	goto L219
L219:
	;
	if v634 == v643 {
		v646 = v634
		goto L212
	} else {
		goto L220
	}
L220:
	;
	v657 = v634 - v643
	goto L209
L221:
	;
	goto L211
L222:
	;
	v659 = int32(1638720)
	goto L152
L223:
	;
	v698 = v661
	v704 = v45
	goto L12
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v436
	v676 = F_format_elog_string(m, int32(666057), v12+int32(16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L2
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, _consts[826])) = v676
	v698 = v424
	v704 = int32(0)
	goto L12
L226:
	;
	v685 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v685
	goto L229
L227:
	;
	goto L228
L228:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v679)+4))
	v698 = v695
	v704 = v45
	goto L12
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v56
	v691 = F_format_elog_string(m, int32(589829), v12)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L2
	} else {
		goto L230
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, _consts[826])) = v691
	v698 = int32(0)
	v704 = int32(0)
	goto L12
L231:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	F_list_free(m, v707)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L2
	} else {
		goto L232
	}
L232:
	;
	if v704 == int32(0) {
		v720 = v46
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v713 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L2
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v713
	if v713 == int32(0) {
		v720 = v46
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v713))) = v698
	v720 = int32(1)
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
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
				v30 = *(*int32)(unsafe.Add(mBase, _consts[68]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v25<<(uint(int32(3))%32))))
				v35 = v34
			} else {
				v35 = v25
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
			if v39&int32(32) != 0 {
				v43 = *(*int32)(unsafe.Add(mBase, _consts[68]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v38<<(uint(int32(3))%32))+4))
				v48 = v47
			} else {
				v48 = v38
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v48
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v68 = v50
		} else {
			if v21&int32(2048) != 0 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
				v65 = int32(-1)
			} else {
				if v21&int32(128) != 0 {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
					v65 = int32(-1)
				} else {
					if v21&int32(4176) != int32(64) {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(-1)
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						v65 = v64
					} else {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v59
						v65 = int32(-1)
					}
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v65
			v68 = int32(-1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v68
		F_XLogBeginInsert(m)
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return
		} else {
			F_XLogRegisterData(m, v7+int32(12), int32(34))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				v79 = F_XLogInsert(m, int32(9), int32(112))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 < int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v12+(l0^int32(-1))<<(uint(int32(2))%32))))
		v26 = v18
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v26 = v20 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v28 = v7 + int32(20)
	if l0 < int32(0) {
		v37 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		v50 = v37 + (l0^int32(-1))<<(uint(int32(6))%32)
	} else {
		v44 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	F_log_newpage(m, v7+int32(20), v61, v62, v26, l1)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		return
	} else {
		m.G0 = v7 + int32(32)
		return
	}
}
func F_show_log_file_mode(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[429]))
	*(*int32)(unsafe.Add(mBase, uint32(v4))) = v7
	v12 = F_pg_snprintf(m, int32(4412733), int32(12), int32(242827), v4)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		m.G0 = v4 + int32(16)
		return int32(4412733)
	}
}
