package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitReadBuffers(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v176 int64
	_ = v176
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v223 int32
	_ = v223
	var v227 int64
	_ = v227
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v243 int64
	_ = v243
	var v247 int32
	_ = v247
	var v260 int32
	_ = v260
	var v264 int64
	_ = v264
	var v268 int64
	_ = v268
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v320 int64
	_ = v320
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v335 int64
	_ = v335
	var v339 int32
	_ = v339
	var v341 int64
	_ = v341
	var v351 int32
	_ = v351
	var v355 int64
	_ = v355
	var v359 int64
	_ = v359
	var v361 int32
	_ = v361
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int64
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v419 int32
	_ = v419
	var v424 int64
	_ = v424
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v21 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = l0 + int32(36)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	goto L8
L2:
	;
	v31 = int32(1)
	v32 = int32(3)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v27 = F_IOContextForStrategy(m, v26)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v31 = v2
	v32 = v27
	goto L1
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L5
	} else {
		goto L99
	}
L8:
	;
	if base.B2i32(v35 != int32(-1)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[0]))
	if v41 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v43 = l0 + int32(48)
	v45 = l0 + int32(56)
	v56 = v2
	goto L13
L12:
	;
	goto L11
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	goto L16
L14:
	;
	m.G0 = v19 + int32(32)
	return v454
L15:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+30)))
	if v460 != v461 {
		goto L91
	} else {
		goto L92
	}
L16:
	;
	if base.B2i32(v62 != int32(-1)) == int32(0) {
		v454 = v56
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
	if v291 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L19:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	if v70&int32(448) != 0 {
		v285 = v56
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[1]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v79 = v75 + v76<<(uint(int32(7))%32)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v81 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+8)))
	v82 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v34)+4)))
	v83 = int32(0)
	v86 = base.AtomicRmwOr32(m, v83, int32(_a_F_WaitReadBuffers_0), v83)
	v87 = int32(1)
	v92 = v81 | v82<<(uint(int64(32))%64)
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	if base.B2i32(v80 == v83)|base.B2i32(v92 != v93) != 0 {
		v131 = v87
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	if v131 != 0 {
		v285 = v56
		goto L18
	} else {
		goto L32
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[2]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[3]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+40))
	if v101 == int32(0) {
		v119 = v80
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v121 = v119 & int32(254)
	v122 = int32(6)
	if base.B2i32(v121 != v122)|base.B2i32(v97 != v98) != 0 {
		v131 = base.B2i32(v121 == v122)
		goto L23
	} else {
		goto L30
	}
L26:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+3)))
	if v104&int32(1) != 0 {
		v119 = v80
		goto L25
	} else {
		goto L27
	}
L27:
	;
	m.T0[v101].(func(*base.Module, int32, int64))(m, v79, v92)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	v110 = int32(0)
	v113 = base.AtomicRmwOr32(m, v110, int32(_a_F_WaitReadBuffers_0), v110)
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v79)+48))
	if base.B2i32(v109 == v110)|base.B2i32(v116 != v92) != 0 {
		v131 = v87
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v119 = v109
	goto L25
L30:
	;
	F_pgaio_io_reclaim(m, v79)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v131 = int32(1)
	goto L23
L32:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[4])))
	v140 = m.G0
	v142 = v140 - int32(16)
	m.G0 = v142
	if v137 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_pgaio_wref_wait(m, v34)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L37
	}
L34:
	;
	F___clock_gettime(m, int32(1), v142)
	mBase = m.M
	v146 = int64(*(*int32)(unsafe.Add(mBase, uint32(v142)+8)))
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
	v151 = v146 + v147*int64(1000000000)
	goto L36
L35:
	;
	v151 = int64(0)
	goto L36
L36:
	;
	m.G0 = v142 + int32(16)
	goto L33
L37:
	;
	v158 = int32(0)
	v159 = int64(0)
	v163 = m.G0
	v165 = v163 - int32(16)
	m.G0 = v165
	if v151 != v159 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v285 = int32(1)
	goto L18
L39:
	;
	F___clock_gettime(m, int32(1), v165)
	mBase = m.M
	v171 = int64(*(*int32)(unsafe.Add(mBase, uint32(v165)+8)))
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v165)))
	v176 = v171 + (v172*int64(1000000000) - v151)
	if v31 == int32(2) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v260 = v31*int32(320) + v32<<(uint(int32(6))%32)
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_WaitReadBuffers[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_WaitReadBuffers[5]))) = v264 + base.I64_extend_i32_u(v158)
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_WaitReadBuffers[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v260)+uint32(_c_F_WaitReadBuffers[6]))) = v268 + v159
	F_pgstat_count_backend_io_op(m, v31, v32, int32(6), v158, v159)
	mBase = m.M
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[7])) = uint8(v273)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[8])) = uint8(v273)
	m.G0 = v165 + int32(16)
	goto L38
L42:
	;
	v223 = v31*int32(320) + v32<<(uint(int32(6))%32)
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_c_F_WaitReadBuffers[9])))
	*(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_c_F_WaitReadBuffers[9]))) = v227 + v176
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[10]))
	v238 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v231))|base.B2i32(int32(1)<<(uint(v231)%32)&int32(_a_F_WaitReadBuffers_1) == v238) == v238 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	goto L45
L45:
	;
	goto L46
L46:
	;
	goto L49
L49:
	;
	v202 = int32(_a_F_WaitReadBuffers_2)
	v204 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[11]))
	v206 = base.I64_div_s(v176, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[11])) = v204 + v206
	switch v31 {
	case 0:
		goto L51
	case 1:
		goto L50
	default:
		goto L42
	}
L50:
	;
	v214 = int32(_a_F_WaitReadBuffers_3)
	v216 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[12])) = v216 + v176
	goto L42
L51:
	;
	v209 = int32(_a_F_WaitReadBuffers_4)
	v211 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[13]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[13])) = v211 + v176
	goto L42
L52:
	;
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_c_F_WaitReadBuffers[14])))
	*(*int64)(unsafe.Add(mBase, uint32(v223)+uint32(_c_F_WaitReadBuffers[14]))) = v243 + v176
	v247 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[7])) = uint8(v247)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[15])) = uint8(v247)
	goto L54
L53:
	;
	goto L54
L54:
	;
	goto L41
L55:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v295 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+32)))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(int32(2))%32))))
	if v299 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v401 = int32(base.Ui32(v397)>>(uint(int32(6))%32)) & int32(7)
	if v401 == int32(4) {
		goto L80
	} else {
		goto L81
	}
L58:
	;
	v317 = int64(0)
	v320 = base.AtomicRmwCmpxchg64(m, v316, int32(24), v317, v317)
	if v320&int64(16777216) == v317 {
		v454 = v285
		goto L15
	} else {
		goto L62
	}
L59:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[16]))
	v316 = v303 + (v299^int32(-1))*int32(56)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[17]))
	v311 = int32(56)
	v316 = v310 + v299*v311 - v311
	goto L58
L62:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	v327 = v325 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v327)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v330 == int32(116) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v351 = v31*int32(320) + v32<<(uint(int32(6))%32)
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_WaitReadBuffers[18])))
	*(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_WaitReadBuffers[18]))) = v355 + int64(1)
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_WaitReadBuffers[19])))
	*(*int64)(unsafe.Add(mBase, uint32(v351)+uint32(_c_F_WaitReadBuffers[19]))) = v359
	v361 = int32(1)
	F_pgstat_count_backend_io_op(m, v31, v32, int32(2), v361, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[7])) = uint8(v361)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[8])) = uint8(v361)
	goto L67
L64:
	;
	v333 = int32(_a_F_WaitReadBuffers_5)
	v335 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[20]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[20])) = v335 + int64(1)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v339 = int32(_a_F_WaitReadBuffers_6)
	v341 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[21]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[21])) = v341 + int64(1)
	goto L63
L67:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[22])))
	if v371 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v374 = int32(_a_F_WaitReadBuffers_7)
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[23]))
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[24]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[23])) = v376 + v378
	goto L70
L69:
	;
	goto L70
L70:
	;
	if v329 == int32(0) {
		v454 = v285
		goto L15
	} else {
		goto L71
	}
L71:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v329)+272))
	if v383 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329)+268)))
	if v386 != int32(1) {
		v454 = v285
		goto L15
	} else {
		goto L75
	}
L73:
	;
	v392 = v383
	goto L74
L74:
	;
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v392)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v392)+120)) = v393 + int64(1)
	v454 = v285
	goto L15
L75:
	;
	F_pgstat_assoc_relation(m, v329)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v329)+272))
	v392 = v391
	goto L74
L77:
	;
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	v449 = v448 + v447
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v449)
	v454 = v285
	goto L15
L78:
	;
	if v397&int32(448) != int32(128) {
		v447 = v406
		goto L77
	} else {
		goto L85
	}
L79:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v414
	F_pgaio_result_report(m, v19+int32(8), v45, v413)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L84
	}
L80:
	;
	v412 = int32(0)
	v413 = int32(21)
	goto L79
L81:
	;
	goto L82
L82:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32(int32(1)) < base.Ui32(v401-int32(3)) {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v412 = v406
	v413 = int32(19)
	goto L79
L84:
	;
	v447 = v412
	goto L77
L85:
	;
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v424
	F_pgaio_result_report(m, v19+int32(16), v45, int32(14))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	v433 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	if v433 == int32(0) {
		v447 = v406
		goto L77
	} else {
		goto L88
	}
L88:
	;
	F_errmsg_internal(m, int32(_a_F_WaitReadBuffers_8), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_WaitReadBuffers_9), int32(1741), int32(_a_F_WaitReadBuffers_10))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	v447 = v406
	goto L77
L91:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[25]))
	if v464 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	goto L14
L94:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L5
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v469 = F_AsyncReadBuffers(m, l0, v19+int32(28))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L5
	} else {
		goto L98
	}
L97:
	;
	goto L96
L98:
	;
	v56 = int32(1)
	goto L13
L99:
	;
	F_errmsg_internal(m, int32(_a_F_WaitReadBuffers_11), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_WaitReadBuffers_9), int32(1789), int32(_a_F_WaitReadBuffers_12))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___wasi_syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	if l0 == int32(0) {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F___wasi_syscall_ret[0])) = l0
		return int32(-1)
	}
}
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v3) {
		v6 = v3
	} else {
		v6 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l0
	{
		m.ExcTag = uint32(int32(0))
		m.ExcVals[0] = uint64(uint32(l0 + int32(8)))
		m.ExcPending = 1
	}
	return
}
func F___wasm_setjmp(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
	return
}
func F_waitonlock_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = v7 + int32(16)
	F_initStringInfo(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_DescribeLockTag(m, v11, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_set_errcontext_domain(m, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
				v20 = int32(2)
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19<<(uint(v20)%32))+uint32(_c_F_waitonlock_error_callback[0])))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+v9<<(uint(v20)%32))))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v29
				F_errcontext_msg(m, int32(_a_F_waitonlock_error_callback_0), v7)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			}
		}
	}
}
func F_walkdir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(2064)
	m.G0 = v11
	v13 = F_AllocateDir(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_ReadDirExtended(m, v13, l0, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v13)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_walkdir[0]))
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+19)))
	if v29 != int32(46) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v66 = F_ReadDirExtended(m, v13, l0, l3)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v22 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v46 = v11 + int32(16)
	v49 = F_pg_snprintf(m, v46, int32(2048), int32(_a_F_walkdir_0), v11)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v32 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v35 != int32(46) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)))
	if v38 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v51 = F_get_dirent_type(m, v46, v22, l2, l3)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_walkdir(m, v11+int32(16), l1, int32(0), l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, v11+int32(16), int32(0), l3)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	switch v51 - int32(2) {
	case 0:
		goto L21
	case 1:
		goto L20
	default:
		goto L13
	}
L23:
	;
	goto L13
L24:
	;
	goto L13
L25:
	;
	if v66 != 0 {
		v22 = v66
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	if v13 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l3)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	m.G0 = v11 + int32(2064)
	return
L31:
	;
	goto L30
}
func F_wcrtomb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	v2 = l1
	if l0 != 0 {
		if base.Ui32(v2) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
			return int32(1)
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[0]))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v7 == int32(0) {
				if v2&int32(-128) == int32(_a_F_wcrtomb_0) {
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
					return int32(1)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[1])) = int32(25)
					v91 = int32(-1)
					return v91
				}
			} else {
				if base.Ui32(v2) <= base.Ui32(int32(2047)) {
					v19 = v2&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v19)
					v24 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v24)
					return int32(2)
				} else {
					if base.B2i32(v2&int32(-8192) != int32(_a_F_wcrtomb_1))&base.B2i32(base.Ui32(int32(_a_F_wcrtomb_2)) <= base.Ui32(v2)) == int32(0) {
						v37 = int32(63)
						v39 = int32(128)
						v40 = v2&v37 | v39
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v40)
						v45 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v45)
						v52 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v37 | v39
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v52)
						return int32(3)
					} else {
						if base.Ui32(v2-int32(_a_F_wcrtomb_3)) <= base.Ui32(int32(_a_F_wcrtomb_4)) {
							v60 = int32(63)
							v62 = int32(128)
							v63 = v2&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v63)
							v68 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v68)
							v75 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v75)
							v82 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v82)
							return int32(4)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[1])) = int32(25)
							v91 = int32(-1)
							return v91
						}
					}
				}
			}
		}
	} else {
		v91 = int32(1)
		return v91
	}
}
func F_wctomb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v2 = l1
	if l0 == int32(0) {
		return int32(0)
	} else {
		if l0 != 0 {
			if base.Ui32(v2) <= base.Ui32(int32(127)) {
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
				v95 = int32(1)
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, _c_F_wctomb[0]))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v11 == int32(0) {
					if v2&int32(-128) == int32(_a_F_wctomb_0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
						v95 = int32(1)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_wctomb[1])) = int32(25)
						v92 = int32(-1)
						v95 = v92
					}
				} else {
					if base.Ui32(v2) <= base.Ui32(int32(2047)) {
						v23 = v2&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v23)
						v28 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v28)
						v95 = int32(2)
					} else {
						if base.B2i32(v2&int32(-8192) != int32(_a_F_wctomb_1))&base.B2i32(base.Ui32(int32(_a_F_wctomb_2)) <= base.Ui32(v2)) == int32(0) {
							v40 = int32(63)
							v42 = int32(128)
							v43 = v2&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
							v48 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
							v55 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v55)
							v95 = int32(3)
						} else {
							if base.Ui32(v2-int32(_a_F_wctomb_3)) <= base.Ui32(int32(_a_F_wctomb_4)) {
								v62 = int32(63)
								v64 = int32(128)
								v65 = v2&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v65)
								v70 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v70)
								v77 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v77)
								v84 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v84)
								v95 = int32(4)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_wctomb[1])) = int32(25)
								v92 = int32(-1)
								v95 = v92
							}
						}
					}
				}
			}
		} else {
			v92 = int32(1)
			v95 = v92
		}
		return v95
	}
}
func F_win1251_to_win866(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_win1251_to_win866_0), int32(20), int32(23))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_win866_to_iso(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_win866_to_iso_0), int32(25), int32(20))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_write_syslogger_file(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	if l2&int32(8) != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[0]))
		if v9 != 0 {
			v21 = v9
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[1]))
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[2]))
			if v12 != 0 {
				v15 = v12
			} else {
				v15 = v14
			}
			if int32(base.Ui32(l2&int32(16))>>(uint(int32(4))%32)) != 0 {
				v20 = v15
			} else {
				v20 = v14
			}
			v21 = v20
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[1]))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[2]))
		if v12 != 0 {
			v15 = v12
		} else {
			v15 = v14
		}
		if int32(base.Ui32(l2&int32(16))>>(uint(int32(4))%32)) != 0 {
			v20 = v15
		} else {
			v20 = v14
		}
		v21 = v20
	}
	v24 = F_fwrite(m, l0, int32(1), l1, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		if v24 != l1 {
			F_write_stderr(m, int32(_a_F_write_syslogger_file_0), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_writetup_datum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)))
	if v12 != 0 {
		v26 = v4
		v27 = v4
		v28 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27 + v28
		v32 = v9 + int32(12)
		F_LogicalTapeWrite(m, l1, v32, v28)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			F_LogicalTapeWrite(m, l1, v26, v27)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v38&int32(1) != 0 {
					F_LogicalTapeWrite(m, l1, v32, int32(4))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v13 == int32(0) {
			v16 = int32(8)
			v26 = l2 + v16
			v27 = v16
			v28 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27 + v28
			v32 = v9 + int32(12)
			F_LogicalTapeWrite(m, l1, v32, v28)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_LogicalTapeWrite(m, l1, v26, v27)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v38&int32(1) != 0 {
						F_LogicalTapeWrite(m, l1, v32, int32(4))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			v24 = F_datumGetSize(m, base.I64_extend_i32_u(v19), int32(0), v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = v19
				v27 = v24
				v28 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v27 + v28
				v32 = v9 + int32(12)
				F_LogicalTapeWrite(m, l1, v32, v28)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_LogicalTapeWrite(m, l1, v26, v27)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v38&int32(1) != 0 {
							F_LogicalTapeWrite(m, l1, v32, int32(4))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_writetup_heap_2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v26 int32
	_ = v26
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
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v10 - int32(6)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = v8 + int32(12)
	F_BufFileWrite(m, v14, v16, int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = int32(10)
		F_BufFileWrite(m, v20, l1+v21, v10-v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v27 == int32(1) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_BufFileWrite(m, v30, v16, int32(4))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = F_GetMemoryChunkSpace(m, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + base.I64_extend_i32_u(v34)
						F_pfree(m, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v34 = F_GetMemoryChunkSpace(m, l1)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + base.I64_extend_i32_u(v34)
					F_pfree(m, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_writetup_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11 + v12
	v16 = v8 + int32(12)
	F_LogicalTapeWrite(m, l1, v16, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		F_LogicalTapeWrite(m, l1, v10, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v23&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v16, int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
