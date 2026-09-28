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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v155 float64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v302 int32
	_ = v302
	var v303 float64
	_ = v303
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v342 int32
	_ = v342
	var v344 float64
	_ = v344
	var v351 float64
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v429 float64
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v495 int32
	_ = v495
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = int32(_a_F_gistBuildCallback_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[0])) = v29
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
	*(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[0])) = v26
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
	v97 = v70
	goto L14
L14:
	;
	if v97 == int32(2) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+32))
	v90 = F_pow(m, base.F64_div(base.F64_convert_i32_u(int32(_a_F_gistBuildCallback_1)-v80), base.F64_div(base.F64_convert_i64_s(v83), base.F64_convert_i64_s(v73))), base.F64_convert_i32_s(v88))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v78)+36)) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_add(v90, v90)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v97 = v95
	goto L14
L16:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v151 = int32(_a_F_gistBuildCallback_1) - v150
	v152 = *(*int64)(unsafe.Add(mBase, uint32(l5)+32))
	v155 = base.F64_div(base.F64_convert_i64_s(v152), base.F64_convert_i64_s(v148))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+52))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 <= int32(0) {
		goto L36
	} else {
		goto L37
	}
L17:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+24)))
	if v102 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	v140 = v97
	goto L19
L19:
	;
	if v140 != int32(3) {
		goto L11
	} else {
		goto L33
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[1]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v105 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v131 = v105
	goto L23
L22:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+16)) = v109
	v113 = F_smgropen(m, v23+int32(16), v106)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v133 = F_smgrnblocks(m, v131, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L29
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v113
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+72))
	if v117 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = v129
	goto L23
L26:
	;
	v125 = v117
	goto L28
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v113)+72))
	v125 = v123
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+72)) = v125 + int32(1)
	goto L25
L29:
	;
	if base.Ui32(v104) < base.Ui32(v133) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v136 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	v148 = v136
	goto L16
L31:
	;
	goto L32
L32:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
	v140 = v137
	goto L19
L33:
	;
	v143 = *(*int64)(unsafe.Add(mBase, uint32(l5)+24))
	if v143 < int64(4096) {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v148 = v143
	goto L16
L35:
	;
	v302 = base.I32_div_u_s(v151, base.I32_trunc_sat_f64_u(v155))
	v303 = base.F64_convert_i32_u(v302)
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[2]))
	v312 = base.I32_div_u_s(v151, v282)
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[1]))
	v317 = base.I32_div_s(v315, int32(4))
	v323 = int32(1)
	goto L64
L36:
	;
	v282 = int32(8)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v164 = v159 & int32(3)
	v165 = int32(8)
	v166 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v159) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v173 = v165
	v174 = int32(0)
	v175 = v166
	goto L42
L40:
	;
	v227 = v165
	v229 = v166
	goto L41
L41:
	;
	v247 = int32(0)
	v248 = v227
	v250 = v229
	goto L58
L42:
	;
	v195 = v158 + v175<<(uint(int32(3))%32)
	v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+30)))
	if v196 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v164 == int32(0) {
		v282 = v218
		goto L35
	} else {
		goto L57
	}
L44:
	;
	v199 = int32(4)
	goto L46
L45:
	;
	v199 = v196
	goto L46
L46:
	;
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+38)))
	if v202 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v205 = int32(4)
	goto L49
L48:
	;
	v205 = v202
	goto L49
L49:
	;
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+46)))
	if v208 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v211 = int32(4)
	goto L52
L51:
	;
	v211 = v208
	goto L52
L52:
	;
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+54)))
	if v214 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v217 = int32(4)
	goto L55
L54:
	;
	v217 = v214
	goto L55
L55:
	;
	v218 = v173 + v199 + v205 + v211 + v217
	v219 = int32(4)
	v220 = v175 + v219
	v222 = v174 + v219
	if v222 != v159&int32(2147483644) {
		v173 = v218
		v174 = v222
		v175 = v220
		goto L42
	} else {
		goto L56
	}
L56:
	;
	goto L43
L57:
	;
	v227 = v218
	v229 = v220
	goto L41
L58:
	;
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v158+v250<<(uint(int32(3))%32))+30)))
	if v271 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v282 = v275
	goto L35
L60:
	;
	v274 = int32(4)
	goto L62
L61:
	;
	v274 = v271
	goto L62
L62:
	;
	v275 = v248 + v274
	v276 = int32(1)
	v279 = v247 + v276
	if v279 != v164 {
		v247 = v279
		v248 = v275
		v250 = v250 + v276
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
L64:
	;
	v342 = v323 + int32(1)
	v344 = F_pow(m, v303, base.F64_convert_i32_s(v342))
	mBase = m.M
	if base.F64_gt(base.F64_div(base.F64_sub(float64(1), v344), base.F64_sub(float64(1), v303)), base.F64_convert_i32_s(v317)) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v355 = int32(0)
	v357 = v323 - int32(1)
	if v355 < v357 {
		goto L70
	} else {
		goto L71
	}
L66:
	;
	v351 = F_pow(m, base.F64_convert_i32_u(v312), base.F64_convert_i32_s(v323))
	mBase = m.M
	if base.F64_gt(v351, base.F64_mul(base.F64_mul(base.F64_convert_i32_s(v306), float64(1024)), float64(0.0001220703125))) == int32(0) {
		v323 = v342
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
	v361 = int32(0)
	v364 = v355
	goto L73
L71:
	;
	goto L72
L72:
	;
	v525 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L97
	}
L73:
	;
	v381 = F_ReadBuffer(m, v157, v364)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	F_UnlockReleaseBuffer(m, v381)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L85
	}
L75:
	;
	F_LockBufferInternal(m, v381, int32(1))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v381 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v403)+16)))
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404+v403)+12)))
	if v406&int32(1) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[3]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389+(v381^int32(-1))<<(uint(int32(2))%32))))
	v403 = v395
	goto L77
L79:
	;
	goto L80
L80:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[4]))
	v403 = v397 + v381<<(uint(int32(13))%32) + int32(-8192)
	goto L77
L81:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v403)+24))
	v414 = v403 + v411&int32(_a_F_gistBuildCallback_2)
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414)+2)))
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v414))))
	F_UnlockReleaseBuffer(m, v381)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L74
L84:
	;
	v361 = v361 + int32(1)
	v364 = v415 | v416<<(uint(int32(16))%32)
	goto L73
L85:
	;
	v429 = F_pow(m, base.F64_div(base.F64_convert_i32_u(v151), v155), base.F64_convert_i32_u(v357))
	mBase = m.M
	v432 = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_add(v429, v429)))
	v433 = m.G0
	v435 = v433 - int32(48)
	m.G0 = v435
	v438 = F_palloc(m, int32(64))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+32)) = v357
	*(*int32)(unsafe.Add(mBase, uint32(v438)+36)) = v432
	v443 = F_BufFileCreateTemp(m, int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v438)+16)) = int64(137438953472)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v443
	v452 = F_palloc_mul(m, int32(4), int32(32))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+12)) = v452
	v456 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v438))) = v456
	*(*int32)(unsafe.Add(mBase, uint32(v435)+36)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v435)+8)) = int64(103079215108)
	v464 = F_hash_create(m, int32(_a_F_gistBuildCallback_3), int64(1024), v435, int32(1064))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v466 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+44)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v438)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+24)) = v464
	v473 = F_palloc_mul(m, int32(4), v466)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+40)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = int32(0)
	v478 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+56)) = v478
	v482 = F_palloc_mul(m, int32(4), v478)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v438)+60)) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v438)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v438)+48)) = v482
	m.G0 = v435 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+40)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = int64(34359738372)
	v495 = *(*int32)(unsafe.Add(mBase, _c_F_gistBuildCallback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v495
	v502 = F_hash_create(m, int32(_a_F_gistBuildCallback_4), int64(1024), v23+int32(32), int32(1064))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+44)) = v502
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = int32(4)
	v509 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v509 == int32(0) {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v357
	F_errmsg_internal(m, int32(_a_F_gistBuildCallback_5), v23)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_gistBuildCallback_6), int32(778), int32(_a_F_gistBuildCallback_7))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L11
L97:
	;
	if v525 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	F_errmsg_internal(m, int32(_a_F_gistBuildCallback_8), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = int32(1)
	goto L11
L101:
	;
	F_errfinish(m, int32(_a_F_gistBuildCallback_6), int32(757), int32(_a_F_gistBuildCallback_7))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	goto L100
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
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
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
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int64
	_ = v314
	var v315 int64
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v346 int32
	_ = v346
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
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
	v486 = m.ExcPending
	if v486 != 0 {
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
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L6
L8:
	;
	goto L9
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[1]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L6
L10:
	;
	v68 = int32(base.Ui32(v60+int32(_a_F_gistFindCorrectParent_0)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v68 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v55-int32(1))&int32(_a_F_gistFindCorrectParent_1)) < base.Ui32(v68&int32(_a_F_gistFindCorrectParent_1)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v53+v55<<(uint(int32(2))%32))+20))
	v79 = v53 + v76&int32(_a_F_gistFindCorrectParent_2)
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
	v94 = v35
	goto L17
L16:
	;
	goto L15
L17:
	;
	v101 = v89 & int32(_a_F_gistFindCorrectParent_1)
	if base.Ui32(v101) < base.Ui32(int32(25)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+28))
	if v205 != 0 {
		goto L39
	} else {
		goto L40
	}
L19:
	;
	v163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v91+v163)+8))
	v166 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+26)) = uint16(v166)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v165
	F_UnlockReleaseBuffer(m, v94)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L28
	}
L20:
	;
	v109 = int32(base.Ui32(v101+int32(_a_F_gistFindCorrectParent_0))>>(uint(int32(2))%32)) & int32(_a_F_gistFindCorrectParent_1)
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
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v91+int32(20)+v118&int32(_a_F_gistFindCorrectParent_1)<<(uint(int32(2))%32))))
	v137 = v91 + v134&int32(_a_F_gistFindCorrectParent_2)
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137))))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+2)))
	if v114 == v138<<(uint(int32(16))%32)|v141 {
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
	v146 = v118 + int32(1)
	if base.Ui32(v146&int32(_a_F_gistFindCorrectParent_1)) <= base.Ui32(v109) {
		v118 = v146
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v171 != int32(-1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v174 = F_ReadBuffer(m, l0, v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v174
	F_LockBufferInternal(m, v174, int32(3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	F_gistcheckpage(m, l0, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v183 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v201
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+12)))
	v89 = v203
	v91 = v201
	v94 = v183
	goto L17
L36:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[0]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187+(v183^int32(-1))<<(uint(int32(2))%32))))
	v201 = v193
	goto L35
L37:
	;
	goto L38
L38:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[1]))
	v201 = v195 + v183<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v208 = v205
	goto L42
L40:
	;
	goto L41
L41:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v238 = F_palloc0(m, int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L46
	}
L42:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	F_ReleaseBuffer(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v208)+28))
	if v222 != 0 {
		v208 = v222
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v240 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v238)+26)) = uint16(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v238
	v249 = F_list_make1_impl(m, int32(1), v16+int32(8))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L50
	}
L47:
	;
	goto L3
L48:
	;
	F_UnlockReleaseBuffer(m, v271)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L88
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L4
	} else {
		goto L85
	}
L50:
	;
	if v249 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v257 = v249
	goto L52
L52:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	v268 = F_list_delete_first(m, v257)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v271 = F_ReadBuffer(m, l0, v270)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_LockBufferInternal(m, v271, int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_gistcheckpage(m, l0, v271)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	if v271 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+16)))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v295)+12)))
	if v298&int32(1) != 0 {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[0]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v281+(v271^int32(-1))<<(uint(int32(2))%32))))
	v295 = v287
	goto L58
L60:
	;
	goto L61
L61:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[1]))
	v295 = v289 + v271<<(uint(int32(13))%32) + int32(-8192)
	goto L58
L62:
	;
	F_UnlockReleaseBuffer(m, v271)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L4
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v303 = F_BufferGetLSNAtomic(m, v271)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L66
	}
L65:
	;
	goto L49
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v267)+16)) = v303
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+16)))
	v307 = v295 + v306
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+12)))
	if v308&int32(8) != 0 {
		goto L47
	} else {
		goto L67
	}
L67:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v267)+28))
	if v311 == int32(0) {
		v336 = v268
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+12)))
	if base.Ui32(v338) < base.Ui32(int32(25)) {
		v397 = v336
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v314 = *(*int64)(unsafe.Add(mBase, uint32(v311)+16))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v307)))
	if base.Ui64(base.I64_rotl(v315, int64(32))) <= base.Ui64(v314) {
		v336 = v268
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v307)+8))
	if v319 == int32(-1) {
		v336 = v268
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v323 = F_palloc0(m, int32(32))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+16)))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v295+v325)+8))
	v328 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v323)+26)) = uint16(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v327
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v267)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v323)+28)) = v331
	v333 = F_lcons(m, v323, v268)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	v336 = v333
	goto L68
L74:
	;
	F_UnlockReleaseBuffer(m, v271)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L83
	}
L75:
	;
	v346 = int32(base.Ui32(v338+int32(_a_F_gistFindCorrectParent_0))>>(uint(int32(2))%32)) & int32(_a_F_gistFindCorrectParent_1)
	if v346 == int32(0) {
		v397 = v336
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v354 = int32(1)
	v356 = v336
	goto L77
L77:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v295+int32(20)+v354&int32(_a_F_gistFindCorrectParent_1)<<(uint(int32(2))%32))))
	v373 = v295 + v370&int32(_a_F_gistFindCorrectParent_2)
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v373))))
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v373)+2)))
	v378 = v374<<(uint(int32(16))%32) | v377
	if v378 == v236 {
		goto L48
	} else {
		goto L79
	}
L78:
	;
	v397 = v386
	goto L74
L79:
	;
	v381 = F_palloc0(m, int32(32))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v381)+28)) = v267
	*(*uint16)(unsafe.Add(mBase, uint32(v381)+26)) = uint16(v354)
	*(*int32)(unsafe.Add(mBase, uint32(v381))) = v378
	v386 = F_lappend(m, v356, v381)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v389 = v354 + int32(1)
	if base.Ui32(v389&int32(_a_F_gistFindCorrectParent_1)) <= base.Ui32(v346) {
		v354 = v389
		v356 = v386
		goto L77
	} else {
		goto L82
	}
L82:
	;
	goto L78
L83:
	;
	if v397 != 0 {
		v257 = v397
		goto L52
	} else {
		goto L84
	}
L84:
	;
	goto L53
L85:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v425 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gistFindCorrectParent_3), v16)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_gistFindCorrectParent_4), int32(1017), int32(_a_F_gistFindCorrectParent_5))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
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
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v354)
	v443 = v267
	goto L89
L89:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v443)))
	v455 = F_ReadBuffer(m, l0, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L4
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v267
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	F_LockBufferInternal(m, v479, int32(3))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L97
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443)+4)) = v455
	if v455 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v443)+8)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v443)+28))
	if v477 != 0 {
		v443 = v477
		goto L89
	} else {
		goto L96
	}
L93:
	;
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[0]))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v461+(v455^int32(-1))<<(uint(int32(2))%32))))
	v475 = v467
	goto L92
L94:
	;
	goto L95
L95:
	;
	v469 = *(*int32)(unsafe.Add(mBase, _c_F_gistFindCorrectParent[1]))
	v475 = v469 + v455<<(uint(int32(13))%32) + int32(-8192)
	goto L92
L96:
	;
	goto L90
L97:
	;
	goto L2
L98:
	;
	F_errmsg_internal(m, int32(_a_F_gistFindCorrectParent_6), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(_a_F_gistFindCorrectParent_4), int32(963), int32(_a_F_gistFindCorrectParent_5))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l1
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_gistProcessItup[0]))
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
	return v150
L7:
	;
	v141 = F_gistGetNodeBuffer(m, v20, v38, v32)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L40
	}
L8:
	;
	v128 = F_ReadBuffer(m, v19, v117)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L37
	}
L9:
	;
	v117 = l2
	v120 = int32(-1)
	v126 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v32 = l3
	v38 = l2
	goto L12
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v45 = base.I32_rem_s(v32, v44)
	if v45|base.B2i32(v32 == l3) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v117 = v89
	v120 = v112
	v126 = v76
	goto L8
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	if v32 != v50 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v52 = F_ReadBuffer(m, v19, v38)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	F_LockBufferInternal(m, v52, int32(3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	if v52 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v76 = F_gistchoose(m, v19, v74, v75, v21)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L24
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_gistProcessItup[1]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+(v52^int32(-1))<<(uint(int32(2))%32))))
	v74 = v66
	goto L20
L22:
	;
	goto L23
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_gistProcessItup[2]))
	v74 = v68 + v52<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L24:
	;
	v78 = int32(2)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76<<(uint(v78)%32)+v74)+20))
	v84 = v81&int32(_a_F_gistProcessItup_0) + v74
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84))))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+2)))
	v89 = v85<<(uint(int32(16))%32) | v88
	if v78 <= v32 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v89
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v99 = F_hash_search(m, v93, v16+int32(12), int32(1), v16+int32(11))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v102 = F_gistgetadjusted(m, v19, v84, v75, v21)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v38
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v102
	if v102 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v114 = v32 - int32(1)
	if v114 != 0 {
		v32 = v114
		v38 = v89
		goto L12
	} else {
		goto L36
	}
L31:
	;
	v108 = F_gistbufferinginserttuples(m, l0, v52, v32, v16, int32(1), v76, int32(-1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_UnlockReleaseBuffer(m, v52)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L35
	}
L34:
	;
	v112 = v108
	goto L30
L35:
	;
	v112 = v38
	goto L30
L36:
	;
	goto L13
L37:
	;
	F_LockBufferInternal(m, v128, int32(3))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v133 = int32(0)
	v139 = F_gistbufferinginserttuples(m, l0, v128, v133, v16+int32(4), int32(1), v133, v120, v126)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v150 = v133
	goto L6
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_gistPushItupToNodeBuffer(m, v20, v141, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v150 = base.B2i32(v147 < v146)
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
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
	if base.Ui32(v26-int32(4)) < base.Ui32(v12&int32(_a_F_gist_indexsortbuild_levelstate_add_0)+int32(4)) {
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
					F_PageInit(m, v69, int32(_a_F_gist_indexsortbuild_levelstate_add_1), int32(16))
					mBase = m.M
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
					v74 = v69 + v73
					v75 = int32(_a_F_gist_indexsortbuild_levelstate_add_2)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
					*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v14+v83<<(uint(int32(2))%32))))
					F_gistfillbuffer(m, v87, v9+int32(12), int32(1), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					v56 = F_palloc0(m, int32(_a_F_gist_indexsortbuild_levelstate_add_1))
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
						F_PageInit(m, v69, int32(_a_F_gist_indexsortbuild_levelstate_add_1), int32(16))
						mBase = m.M
						v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
						v74 = v69 + v73
						v75 = int32(_a_F_gist_indexsortbuild_levelstate_add_2)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
						*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v14+v83<<(uint(int32(2))%32))))
						F_gistfillbuffer(m, v87, v9+int32(12), int32(1), int32(0))
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
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
				F_PageInit(m, v69, int32(_a_F_gist_indexsortbuild_levelstate_add_1), int32(16))
				mBase = m.M
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
				v74 = v69 + v73
				v75 = int32(_a_F_gist_indexsortbuild_levelstate_add_2)
				*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
				*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
				*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v87 = *(*int32)(unsafe.Add(mBase, uint32(v14+v83<<(uint(int32(2))%32))))
				F_gistfillbuffer(m, v87, v9+int32(12), int32(1), int32(0))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				v56 = F_palloc0(m, int32(_a_F_gist_indexsortbuild_levelstate_add_1))
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
					F_PageInit(m, v69, int32(_a_F_gist_indexsortbuild_levelstate_add_1), int32(16))
					mBase = m.M
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+16)))
					v74 = v69 + v73
					v75 = int32(_a_F_gist_indexsortbuild_levelstate_add_2)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+14)) = uint16(v75)
					*(*uint16)(unsafe.Add(mBase, uint32(v74)+12)) = uint16(v41)
					*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = int32(-1)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v14+v83<<(uint(int32(2))%32))))
					F_gistfillbuffer(m, v87, v9+int32(12), int32(1), int32(0))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		}
	} else {
		v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v14+v83<<(uint(int32(2))%32))))
		F_gistfillbuffer(m, v87, v9+int32(12), int32(1), int32(0))
		mBase = m.M
		v93 = m.ExcPending
		if v93 != 0 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
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
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	v3 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v3)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	v12 = v10 & int32(_a_F_gist_mask_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v12)
	F_mask_unused_space(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		*(*int64)(unsafe.Add(mBase, uint32(l0+v16))) = int64(0)
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v21 = l0 + v20
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)))
		v24 = v22 | int32(8)
		*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)) = uint16(v24)
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
		v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v26)+12)))
		if v28&int32(1) != 0 {
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
			if base.Ui32(v37) < base.Ui32(int32(25)) {
			} else {
				v43 = int32(base.Ui32(v37+int32(_a_F_gist_mask_1)) >> (uint(int32(2)) % 32))
				if v43&int32(_a_F_gist_mask_2) == int32(0) {
				} else {
					v48 = int32(1)
					v50 = l0 + int32(20)
					v54 = (v43 + v48) & int32(_a_F_gist_mask_2)
					if base.Ui32(int32(3)) <= base.Ui32(v54) {
						v57 = int32(2)
						if base.Ui32(v54) <= base.Ui32(v57) {
							v60 = v57
						} else {
							v60 = v54
						}
						v61 = int32(1)
						v62 = v60 - v61
						v70 = v61
						v71 = int32(0)
						for {
							v78 = v50 + v70<<(uint(int32(2))%32)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
							if v79&int32(_a_F_gist_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v78))) = v79 & int32(-98305)
							} else {
							}
							v86 = v78 + int32(4)
							v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
							if v87&int32(_a_F_gist_mask_3) != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v86))) = v87 & int32(-98305)
							} else {
							}
							v93 = int32(2)
							v94 = v70 + v93
							v96 = v71 + v93
							if v96 != v62&int32(-2) {
								v70 = v94
								v71 = v96
								continue
							} else {
								break
							}
							break
						}
						if v62&v61 == int32(0) {
						} else {
							v101 = v94
							v109 = v50 + v101<<(uint(int32(2))%32)
							v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
							if v110&int32(_a_F_gist_mask_3) == int32(0) {
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110 & int32(-98305)
							}
						}
					} else {
						v101 = v48
						v109 = v50 + v101<<(uint(int32(2))%32)
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						if v110&int32(_a_F_gist_mask_3) == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v109))) = v110 & int32(-98305)
						}
					}
				}
			}
			v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
			v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v125)+12)))
			v128 = v125
			v129 = v127
		} else {
			v128 = v26
			v129 = v28
		}
		v132 = v129 & int32(_a_F_gist_mask_4)
		*(*uint16)(unsafe.Add(mBase, uint32(l0+v128)+12)) = uint16(v132)
		return
	}
}
func F_gist_translate_cmptype_btree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = v2 - int32(1)
	if base.Ui32(v4) < base.Ui32(int32(5)) {
		v11 = base.I64_extend_i32_u(v4) + int64(1)
	} else {
		v11 = int64(0)
	}
	return v11
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
		*(*int32)(unsafe.Add(mBase, _c_F_gist_xlog_startup[0])) = v2
		return
	}
}
