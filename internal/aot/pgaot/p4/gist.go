package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistBuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
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
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int64
	_ = v148
	var v153 int64
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int64
	_ = v160
	var v163 float64
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v321 int32
	_ = v321
	var v323 float64
	_ = v323
	var v330 float64
	_ = v330
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v356 float64
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 float64
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v491 int32
	_ = v491
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = int32(4476144)
	v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v29
	v32 = F_gistFormTuple(m, v28, l0, l2, l3, int32(1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v36
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+24)) = v38 + int64(1)
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v43 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+32)) = v42 + v43&int64(8191)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v48 == int32(4) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v26
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	F_MemoryContextReset(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+60))
	v54 = F_gistProcessItup(m, l5, v32, int32(0), v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_gistdoinsert(m, l0, v32, v58, v59, v60, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	F_gistProcessEmptyingQueue(m, l5)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L3
L9:
	;
	goto L3
L10:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	if v70 == int32(4) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	m.G0 = v23 + int32(80)
	return
L12:
	;
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	if v73&int64(4095) != int64(0) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	v102 = v70
	goto L14
L14:
	;
	if v102 == int32(2) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v82 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+32))
	v90 = F_pow(m, base.F64_div(base.F64_convert_i32_u(int32(8148)-v79), base.F64_div(base.F64_convert_i64_s(v82), base.F64_convert_i64_s(v73))), base.F64_convert_i32_s(v88))
	mBase = m.M
	v92 = base.F64_nearest(base.F64_add(v90, v90))
	if base.F64_lt(base.F64_abs(v92), float64(2.147483648e+09)) != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+36)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v102 = v100
	goto L14
L17:
	;
	v96 = base.I32_trunc_f64_s(v92)
	v98 = v96
	goto L16
L18:
	;
	goto L19
L19:
	;
	v98 = int32(-2147483648)
	goto L16
L20:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+52))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v160 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v163 = base.F64_div(base.F64_convert_i64_s(v160), base.F64_convert_i64_s(v153))
	if base.F64_lt(v163, float64(4.294967296e+09))&base.F64_ge(v163, float64(0)) != 0 {
		goto L40
	} else {
		goto L41
	}
L21:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+24)))
	if v107 != 0 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	v145 = v102
	goto L23
L23:
	;
	if v145 != int32(3) {
		goto L11
	} else {
		goto L37
	}
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v110 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v136 = v110
	goto L27
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v112
	v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v114
	v118 = F_smgropen(m, v23+int32(16), v111)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	v138 = F_smgrnblocks(m, v136, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v118
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118)+72))
	if v122 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v136 = v134
	goto L27
L30:
	;
	v130 = v122
	goto L32
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)+76))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v118)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+72))
	v130 = v128
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+72)) = v130 + int32(1)
	goto L29
L33:
	;
	if base.Ui32(v109) < base.Ui32(v138) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v141 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	v153 = v141
	goto L20
L35:
	;
	goto L36
L36:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v145 = v142
	goto L23
L37:
	;
	v148 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	if v148 < int64(4096) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v153 = v148
	goto L20
L39:
	;
	v173 = int32(8148) - v154
	if v157 <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v169 = base.I32_trunc_f64_u(v163)
	v171 = v169
	goto L39
L41:
	;
	goto L42
L42:
	;
	v171 = int32(0)
	goto L39
L43:
	;
	v281 = base.I32_div_u_s(v173, v171)
	v282 = base.F64_convert_i32_u(v281)
	v285 = *(*int32)(unsafe.Add(mBase, _consts[58]))
	v291 = base.I32_div_u_s(v173, v261)
	v294 = *(*int32)(unsafe.Add(mBase, _consts[79]))
	v296 = base.I32_div_s(v294, int32(4))
	v302 = int32(1)
	goto L64
L44:
	;
	v261 = int32(8)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v175 = int32(1)
	v178 = v156 + int32(24)
	if v157 == v175 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v157&v175 == int32(0) {
		v261 = v231
		goto L43
	} else {
		goto L60
	}
L48:
	;
	v231 = int32(8)
	v250 = int32(0)
	goto L47
L49:
	;
	goto L50
L50:
	;
	v186 = int32(0)
	v188 = v186
	v189 = int32(8)
	v191 = v186
	goto L51
L51:
	;
	v208 = int32(4)
	v211 = v178 + v191<<(uint(v208)%32)
	v212 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211))))
	if v212 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v231 = v222
	v250 = v224 << (uint(int32(4)) % 32)
	goto L47
L53:
	;
	v215 = v208
	goto L55
L54:
	;
	v215 = v212
	goto L55
L55:
	;
	v218 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211)+16)))
	if v218 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v221 = int32(4)
	goto L58
L57:
	;
	v221 = v218
	goto L58
L58:
	;
	v222 = v189 + v215 + v221
	v223 = int32(2)
	v224 = v191 + v223
	v226 = v188 + v223
	if v226 != v157&int32(2147483646) {
		v188 = v226
		v189 = v222
		v191 = v224
		goto L51
	} else {
		goto L59
	}
L59:
	;
	goto L52
L60:
	;
	v255 = int32(*(*int16)(unsafe.Add(mBase, uint32(v250+v178))))
	if v255 < int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v258 = int32(4)
	goto L63
L62:
	;
	v258 = v255
	goto L63
L63:
	;
	v261 = v231 + v258
	goto L43
L64:
	;
	v321 = v302 + int32(1)
	v323 = F_pow(m, v282, base.F64_convert_i32_s(v321))
	mBase = m.M
	if base.F64_gt(base.F64_div(base.F64_sub(float64(1), v323), base.F64_sub(float64(1), v282)), base.F64_convert_i32_s(v296)) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v335 = v302 - int32(1)
	if v335 <= int32(0) {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v330 = F_pow(m, base.F64_convert_i32_u(v291), base.F64_convert_i32_s(v302))
	mBase = m.M
	if base.F64_gt(v330, base.F64_mul(base.F64_mul(base.F64_convert_i32_s(v285), float64(1024)), float64(0.0001220703125))) == int32(0) {
		v302 = v321
		goto L64
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	goto L68
L70:
	;
	v340 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v356 = F_pow(m, base.F64_div(base.F64_convert_i32_u(v173), v163), base.F64_convert_i32_u(v335))
	mBase = m.M
	v358 = int32(0)
	v360 = v358
	v363 = v358
	goto L79
L73:
	;
	if v340 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_errmsg_internal(m, int32(426463), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = int32(1)
	goto L11
L77:
	;
	F_errfinish(m, int32(494176), int32(757), int32(327523))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v380 = F_ReadBuffer(m, v155, v363)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	F_UnlockReleaseBuffer(m, v380)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L91
	}
L81:
	;
	F_LockBuffer(m, v380, int32(1))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v380 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v402)+16)))
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403+v402)+12)))
	if v405&int32(1) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v388+(v380^int32(-1))<<(uint(int32(2))%32))))
	v402 = v394
	goto L83
L85:
	;
	goto L86
L86:
	;
	v396 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v402 = v396 + v380<<(uint(int32(13))%32) + int32(-8192)
	goto L83
L87:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v402)+24))
	v413 = v402 + v410&int32(32767)
	v414 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413)+2)))
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v413))))
	F_UnlockReleaseBuffer(m, v380)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	goto L80
L90:
	;
	v360 = v360 + int32(1)
	v363 = v414 | v415<<(uint(int32(16))%32)
	goto L79
L91:
	;
	v425 = base.F64_nearest(base.F64_add(v356, v356))
	if base.F64_lt(base.F64_abs(v425), float64(2.147483648e+09)) != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v432 = m.G0
	v434 = v432 - int32(48)
	m.G0 = v434
	v437 = F_palloc(m, int32(64))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L96
	}
L93:
	;
	v429 = base.I32_trunc_f64_s(v425)
	v431 = v429
	goto L92
L94:
	;
	goto L95
L95:
	;
	v431 = int32(-2147483648)
	goto L92
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+32)) = v335
	*(*int32)(unsafe.Add(mBase, uint32(v437)+36)) = v431
	v442 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v437)+16)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+4)) = v442
	v450 = F_palloc(m, int32(128))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+12)) = v450
	v454 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v437))) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v434)+40)) = v454
	*(*int64)(unsafe.Add(mBase, uint32(v434)+16)) = int64(103079215108)
	v462 = F_hash_create(m, int32(133625), int32(1024), v434, int32(1064))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+44)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+24)) = v462
	v470 = F_palloc(m, int32(4))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+40)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+56)) = int32(32)
	v478 = F_palloc(m, int32(128))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v437)+60)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v437)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v437)+48)) = v478
	m.G0 = v434 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+40)) = v437
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = int64(34359738372)
	v491 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+72)) = v491
	v498 = F_hash_create(m, int32(236113), int32(1024), v23+int32(32), int32(1064))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+44)) = v498
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = int32(4)
	v505 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v505 == int32(0) {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v335
	F_errmsg_internal(m, int32(474952), v23)
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(494176), int32(778), int32(327523))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	goto L11
}
func F_gistFindCorrectParent(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
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
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int64
	_ = v316
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	goto L2
L1:
	;
	m.G0 = v16 + int32(16)
	return
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	F_gistcheckpage(m, l0, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L4
	} else {
		goto L98
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 < int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v53
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v60) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L6
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v68 = int32(base.Ui32(v60+int32(262120)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v68 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v55-int32(1))&int32(65535)) < base.Ui32(v68&int32(65535)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32)+v53)+20))
	v79 = v53 + v76&int32(32767)
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79))))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+2)))
	if v72 == v80<<(uint(int32(16))%32)|v83 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v89 = v60
	v91 = v53
	v95 = v35
	goto L17
L16:
	;
	goto L15
L17:
	;
	v101 = v89 & int32(65535)
	if base.Ui32(v101) < base.Ui32(int32(25)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+28))
	if v207 != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v91+v165)+8))
	v168 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+26)) = uint16(v168)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v167
	F_UnlockReleaseBuffer(m, v95)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L28
	}
L20:
	;
	v109 = int32(base.Ui32(v101+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v109 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v118 = int32(1)
	goto L22
L22:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v118&int32(65535)<<(uint(int32(2))%32)+(v91+int32(24))-int32(4))))
	v139 = v91 + v136&int32(32767)
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139))))
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+2)))
	if v114 == v140<<(uint(int32(16))%32)|v143 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L19
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v118)
	goto L1
L25:
	;
	goto L26
L26:
	;
	v148 = v118 + int32(1)
	if base.Ui32(v148&int32(65535)) <= base.Ui32(v109) {
		v118 = v148
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v173 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v176 = F_ReadBuffer(m, l0, v173)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L18
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v176
	F_LockBuffer(m, v176, int32(2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	F_gistcheckpage(m, l0, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v185 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v203
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+12)))
	v89 = v205
	v91 = v203
	v95 = v185
	goto L17
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189+(v185^int32(-1))<<(uint(int32(2))%32))))
	v203 = v195
	goto L35
L37:
	;
	goto L38
L38:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v203 = v197 + v185<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v210 = v207
	goto L42
L40:
	;
	goto L41
L41:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v240 = F_palloc0(m, int32(32))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L46
	}
L42:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	F_ReleaseBuffer(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v210)+28))
	if v224 != 0 {
		v210 = v224
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v242 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v240)+26)) = uint16(v242)
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v240
	v251 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	goto L3
L48:
	;
	F_UnlockReleaseBuffer(m, v273)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L88
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L85
	}
L50:
	;
	if v251 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v259 = v251
	goto L52
L52:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v259)+12))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v270 = F_list_delete_first(m, v259)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L4
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	v273 = F_ReadBuffer(m, l0, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_LockBuffer(m, v273, int32(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_gistcheckpage(m, l0, v273)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v273 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+16)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298+v297)+12)))
	if v300&int32(1) != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v283+(v273^int32(-1))<<(uint(int32(2))%32))))
	v297 = v289
	goto L58
L60:
	;
	goto L61
L61:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v297 = v291 + v273<<(uint(int32(13))%32) + int32(-8192)
	goto L58
L62:
	;
	F_UnlockReleaseBuffer(m, v273)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v305 = F_BufferGetLSNAtomic(m, v273)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L49
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v269)+16)) = v305
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+16)))
	v309 = v297 + v308
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+12)))
	if v310&int32(8) != 0 {
		goto L47
	} else {
		goto L67
	}
L67:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v269)+28))
	if v313 == int32(0) {
		v341 = v270
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+12)))
	if base.Ui32(v342) < base.Ui32(int32(25)) {
		v403 = v341
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v313)+16))
	v317 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v309)+4)))
	v318 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v309))))
	if base.Ui64(v317|v318<<(uint(int64(32))%64)) <= base.Ui64(v316) {
		v341 = v270
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v323 == int32(-1) {
		v341 = v270
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v327 = F_palloc0(m, int32(32))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+16)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v297+v329)+8))
	v332 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v327)+26)) = uint16(v332)
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v331
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v269)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v327)+28)) = v335
	v337 = F_lcons(m, v327, v270)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v341 = v337
	goto L68
L74:
	;
	F_UnlockReleaseBuffer(m, v273)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L83
	}
L75:
	;
	v350 = int32(base.Ui32(v342+int32(262120))>>(uint(int32(2))%32)) & int32(65535)
	if v350 == int32(0) {
		v403 = v341
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v358 = int32(1)
	v360 = v341
	goto L77
L77:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v358&int32(65535)<<(uint(int32(2))%32)+(v297+int32(24))-int32(4))))
	v379 = v297 + v376&int32(32767)
	v380 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379))))
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v379)+2)))
	v384 = v380<<(uint(int32(16))%32) | v383
	if v384 == v238 {
		goto L48
	} else {
		goto L79
	}
L78:
	;
	v403 = v392
	goto L74
L79:
	;
	v387 = F_palloc0(m, int32(32))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v387)+28)) = v269
	*(*uint16)(unsafe.Add(mBase, uint32(v387)+26)) = uint16(v358)
	*(*int32)(unsafe.Add(mBase, uint32(v387))) = v384
	v392 = F_lappend(m, v360, v387)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v395 = v358 + int32(1)
	if base.Ui32(v395&int32(65535)) <= base.Ui32(v350) {
		v358 = v395
		v360 = v392
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if v403 != 0 {
		v259 = v403
		goto L52
	} else {
		goto L84
	}
L84:
	;
	goto L53
L85:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v431 + int32(4)
	F_errmsg_internal(m, int32(47640), v16)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(487350), int32(1017), int32(318724))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v358)
	v449 = v269
	goto L89
L89:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	v461 = F_ReadBuffer(m, l0, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v269
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	F_LockBuffer(m, v485, int32(2))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L4
	} else {
		goto L97
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449)+4)) = v461
	if v461 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v449)+8)) = v481
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v449)+28))
	if v483 != 0 {
		v449 = v483
		goto L89
	} else {
		goto L96
	}
L93:
	;
	v467 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v467+(v461^int32(-1))<<(uint(int32(2))%32))))
	v481 = v473
	goto L92
L94:
	;
	goto L95
L95:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v481 = v475 + v461<<(uint(int32(13))%32) + int32(-8192)
	goto L92
L96:
	;
	goto L90
L97:
	;
	goto L2
L98:
	;
	F_errmsg_internal(m, int32(346691), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(487350), int32(963), int32(318724))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_gistProcessItup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l1
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l3 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	m.G0 = v16 + int32(16)
	return v147
L7:
	;
	v138 = F_gistGetNodeBuffer(m, v20, v39, v32)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L40
	}
L8:
	;
	v125 = F_ReadBuffer(m, v19, v114)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L37
	}
L9:
	;
	v114 = l2
	v117 = int32(-1)
	v123 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v32 = l3
	v39 = l2
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v45 = base.I32_rem_s(v32, v44)
	if v45 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v114 = v86
	v117 = v109
	v123 = v73
	goto L8
L14:
	;
	v49 = F_ReadBuffer(m, v19, v39)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L18
	}
L15:
	;
	if v32 == l3 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v32 != v47 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_LockBuffer(m, v49, int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v49 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v73 = F_gistchoose(m, v19, v71, v72, v21)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L24
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v49^int32(-1))<<(uint(int32(2))%32))))
	v71 = v63
	goto L20
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v71 = v65 + v49<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L24:
	;
	v75 = int32(2)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v71+v73<<(uint(v75)%32))+20))
	v81 = v78&int32(32767) + v71
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81))))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+2)))
	v86 = v82<<(uint(int32(16))%32) | v85
	if v75 <= v32 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v96 = F_hash_search(m, v90, v16+int32(12), int32(1), v16+int32(11))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v99 = F_gistgetadjusted(m, v19, v81, v72, v21)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v39
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v99
	if v99 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v111 = v32 - int32(1)
	if v111 != 0 {
		v32 = v111
		v39 = v86
		goto L12
	} else {
		goto L36
	}
L31:
	;
	v105 = F_gistbufferinginserttuples(m, l0, v49, v32, v16, int32(1), v73, int32(-1), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_UnlockReleaseBuffer(m, v49)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	v109 = v105
	goto L30
L35:
	;
	v109 = v39
	goto L30
L36:
	;
	goto L13
L37:
	;
	F_LockBuffer(m, v125, int32(2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v130 = int32(0)
	v136 = F_gistbufferinginserttuples(m, l0, v125, v130, v16+int32(4), int32(1), v130, v117, v123)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v147 = v130
	goto L6
L40:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_gistPushItupToNodeBuffer(m, v20, v138, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v147 = base.B2i32(v144 < v143)
	goto L6
}
func F_gist_indexsortbuild_levelstate_add(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v14 = l1 + int32(12)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14+v15<<(uint(int32(2))%32))))
	v20 = int32(4)
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+14)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+12)))
	v23 = v21 - v22
	if v23 <= v20 {
		v26 = v20
	} else {
		v26 = v23
	}
	if base.Ui32(v26-int32(4)) < base.Ui32(v12&int32(8191)+int32(4)) {
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v14+v34<<(uint(int32(2))%32))))
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38+v39)+12)))
		v43 = v34 + int32(1)
		if v43 == int32(4) {
			F_gist_indexsortbuild_levelstate_flush(m, l0, l1)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v50 = v48
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v14+v50<<(uint(int32(2))%32))))
				if v54 != 0 {
					v69 = v54
					F_PageInit(m, v69, int32(8192), int32(16))
					mBase = m.M
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
					v74 = v69 + v73
					v75 = int32(65409)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
					*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
					F_gistfillbuffer(m, v86, v9+int32(12), int32(1), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					v56 = F_palloc0(m, int32(8192))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v59 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v14+v58<<(uint(v59)%32)))) = v56
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v14+v63<<(uint(v59)%32))))
						v69 = v67
						F_PageInit(m, v69, int32(8192), int32(16))
						mBase = m.M
						v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
						v74 = v69 + v73
						v75 = int32(65409)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
						F_gistfillbuffer(m, v86, v9+int32(12), int32(1), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v43
			v50 = v43
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v14+v50<<(uint(int32(2))%32))))
			if v54 != 0 {
				v69 = v54
				F_PageInit(m, v69, int32(8192), int32(16))
				mBase = m.M
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
				v74 = v69 + v73
				v75 = int32(65409)
				*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
				*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v86 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
				F_gistfillbuffer(m, v86, v9+int32(12), int32(1), int32(0))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				v56 = F_palloc0(m, int32(8192))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v59 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v14+v58<<(uint(v59)%32)))) = v56
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v14+v63<<(uint(v59)%32))))
					v69 = v67
					F_PageInit(m, v69, int32(8192), int32(16))
					mBase = m.M
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
					v74 = v69 + v73
					v75 = int32(65409)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
					*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
					F_gistfillbuffer(m, v86, v9+int32(12), int32(1), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v86 = *(*int32)(unsafe.Add(mBase, uint32(v14+v82<<(uint(int32(2))%32))))
		F_gistfillbuffer(m, v86, v9+int32(12), int32(1), int32(0))
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_gist_mask(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(65528)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0+v16))) = v18
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		*(*int32)(unsafe.Add(mBase, uint32(l0+v20)+4)) = v18
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v25 = l0 + v24
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+12)))
		v28 = v26 | int32(8)
		*(*uint16)(unsafe.Add(mBase, uint32(v25)+12)) = uint16(v28)
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v30)+12)))
		if v32&int32(1) != 0 {
			v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			if base.Ui32(v42) < base.Ui32(int32(25)) {
			} else {
				v48 = int32(base.Ui32(v42+int32(262120)) >> (uint(int32(2)) % 32))
				if v48&int32(65535) == int32(0) {
				} else {
					v53 = int32(1)
					v54 = int32(2)
					v58 = (v48 + v53) & int32(65535)
					if base.Ui32(v58) <= base.Ui32(v54) {
						v61 = v54
					} else {
						v61 = v58
					}
					v62 = int32(1)
					v63 = v61 - v62
					v67 = l0 + int32(24)
					v68 = int32(0)
					if base.Ui32(int32(3)) <= base.Ui32(v58) {
						v73 = v68
						v74 = v53
						for {
							v83 = v74<<(uint(int32(2))%32) + v67
							v85 = v83 - int32(4)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
							if v86&int32(98304) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v85))) = v86 & int32(-98305)
							} else {
							}
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
							if v92&int32(98304) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v83))) = v92 & int32(-98305)
							} else {
							}
							v98 = int32(2)
							v101 = v73 + v98
							if v101 != v63&int32(-2) {
								v73 = v101
								v74 = v74 + v98
								continue
							} else {
								break
							}
							break
						}
						v105 = v74 + int32(1)
					} else {
						v105 = v68
					}
					if v63&v62 == int32(0) {
					} else {
						v117 = v67 + v105<<(uint(int32(2))%32)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
						if v118&int32(98304) == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v117))) = v118 & int32(-98305)
						}
					}
				}
			}
			v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v134)+12)))
			v137 = v134
			v138 = v136
		} else {
			v137 = v30
			v138 = v32
		}
		v141 = v138 & int32(65519)
		*(*uint16)(unsafe.Add(mBase, uint32(l0+v137)+12)) = uint16(v141)
		return
	}
}
func F_gist_translate_cmptype_btree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v2-int32(1)) < base.Ui32(int32(5)) {
		v8 = v2
	} else {
		v8 = int32(0)
	}
	return v8
}
func F_gist_xlog_startup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_createTempGistContext(m)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[81])) = v2
		return
	}
}
