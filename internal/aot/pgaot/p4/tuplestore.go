package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetTuplestoreDestReceiverParams(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	return
}
func F_build_tuplestore_recursively(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	v26 = m.G0
	v28 = v26 - int32(320)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l12)))
	if base.B2i32(int32(0) < l9)&base.B2i32(l9 < l7) != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L7
	} else {
		goto L112
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L105
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L7
	} else {
		goto L98
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L93
	}
L5:
	;
	m.G0 = v28 + int32(320)
	return
L6:
	;
	v36 = v28 + int32(304)
	F_initStringInfo(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return
L8:
	;
	v39 = F_quote_literal_cstr(m, l5)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if l11 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l10 != 0 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+220)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+216)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+212)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+208)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+204)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+200)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28)+196)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+192)) = l0
	F_appendStringInfo(m, v36, int32(_a_F_build_tuplestore_recursively_0), v28+int32(192))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+176)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v28)+172)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+168)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+164)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v28)+160)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v28)+156)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+152)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28)+148)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v28)+144)) = l0
	F_appendStringInfo(m, v28+int32(304), int32(_a_F_build_tuplestore_recursively_1), v28+int32(144))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	v74 = int32(0)
	goto L10
L15:
	;
	v74 = int32(4)
	goto L10
L16:
	;
	v79 = v74 | int32(16)
	goto L18
L17:
	;
	v79 = v74 + int32(12)
	goto L18
L18:
	;
	v80 = F_palloc(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	if l7 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v28)+128)) = v84
	v90 = v28 + int32(292)
	v94 = F_pg_sprintf(m, v90, int32(_a_F_build_tuplestore_recursively_2), v28+int32(128))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	v122 = l7
	goto L22
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v28)+304))
	v126 = F_SPI_execute(m, v123, int32(1), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L36
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v90
	if l10 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = l5
	goto L26
L25:
	;
	goto L26
L26:
	;
	if l11 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v98 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = v98
	v104 = v28 + int32(280)
	v108 = F_pg_sprintf(m, v104, int32(_a_F_build_tuplestore_recursively_2), v28+int32(112))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v116 = F_BuildTupleFromCStrings(m, l12, v80)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L34
	}
L30:
	;
	if l10 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v112 = int32(16)
	goto L33
L32:
	;
	v112 = int32(12)
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80+v112))) = v104
	goto L29
L34:
	;
	F_tuplestore_puttuple(m, l13, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v122 = int32(1)
	goto L22
L36:
	;
	if v126 != int32(5) {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	v131 = *(*int64)(unsafe.Add(mBase, _c_F_build_tuplestore_recursively[0]))
	if v131 == int64(0) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_build_tuplestore_recursively[1]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	if v137 <= int32(1) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v140 = int32(4)
	v142 = v136 + v137<<(uint(v140)%32)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+96))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v147 = v30 + v144<<(uint(v140)%32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+96))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v147)+88))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v142)+88))
	if base.B2i32(v143 != v148)&base.B2i32(int32(0) <= v148)|base.B2i32(v153 != v154) != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v142)+196))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+196))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v147)+188))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v142)+188))
	if base.B2i32(v157 != v158)&base.B2i32(int32(0) <= v158)|base.B2i32(v163 != v164) != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_initStringInfo(m, v28+int32(264))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	F_initStringInfo(m, v28+int32(248))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_initStringInfo(m, v28+int32(232))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v204 = int64(0)
	goto L45
L45:
	;
	v207 = v28 + int32(264)
	F_appendStringInfoString(m, v207, l6)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L7
	} else {
		goto L47
	}
L46:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v28)+264))
	if v326 != 0 {
		goto L83
	} else {
		goto L84
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = l4
	F_appendStringInfo(m, v28+int32(248), int32(_a_F_build_tuplestore_recursively_3), v28-int32(-64))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220+base.I32_wrap_i64(v204)<<(uint(int32(2))%32))))
	v227 = F_SPI_getvalue(m, v225, v136, int32(1))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v230 = F_SPI_getvalue(m, v225, v136, int32(2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v122
	v238 = F_pg_sprintf(m, v28+int32(292), int32(_a_F_build_tuplestore_recursively_2), v28+int32(48))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if v227 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+40)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = l4
	F_appendStringInfo(m, v28+int32(232), int32(_a_F_build_tuplestore_recursively_3), v28+int32(32))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v28)+264))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v28 + int32(292)
	if l10 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	v252 = F_strstr(m, v250, v251)
	mBase = m.M
	if v252 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = l4
	F_appendStringInfo(m, v207, int32(_a_F_build_tuplestore_recursively_4), v28+int32(16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v260
	goto L60
L59:
	;
	goto L60
L60:
	;
	if l11 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v284 = F_BuildTupleFromCStrings(m, l12, v80)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L7
	} else {
		goto L67
	}
L62:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v269 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v269
	v275 = v28 + int32(280)
	v277 = F_pg_sprintf(m, v275, int32(_a_F_build_tuplestore_recursively_2), v28)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	if l10 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v275
	goto L61
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v28 + int32(280)
	goto L61
L67:
	;
	F_tuplestore_puttuple(m, l13, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v284)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	if v227 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_build_tuplestore_recursively(m, l0, l1, l2, l3, l4, v227, v260, v122+int32(1), l8, l9, l10, l11, l12, l13)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L7
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v230 != 0 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	F_pfree(m, v227)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	F_pfree(m, v230)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v297 = v28 + int32(264)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v299)
	*(*int32)(unsafe.Add(mBase, uint32(v297)+12)) = v299
	*(*int32)(unsafe.Add(mBase, uint32(v297)+4)) = v299
	goto L79
L78:
	;
	goto L77
L79:
	;
	v306 = v28 + int32(248)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v307))) = uint8(v308)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+12)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v308
	goto L80
L80:
	;
	v315 = v28 + int32(232)
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v316))) = uint8(v317)
	*(*int32)(unsafe.Add(mBase, uint32(v315)+12)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = v317
	goto L81
L81:
	;
	v324 = v204 + int64(1)
	if v324 != v131 {
		v204 = v324
		goto L45
	} else {
		goto L82
	}
L82:
	;
	goto L46
L83:
	;
	F_pfree(m, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L7
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v28)+248))
	if v331 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+264)) = int32(0)
	goto L85
L87:
	;
	F_pfree(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v28)+232))
	if v336 == int32(0) {
		goto L5
	} else {
		goto L91
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+248)) = int32(0)
	goto L89
L91:
	;
	F_pfree(m, v336)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L7
	} else {
		goto L92
	}
L92:
	;
	goto L5
L93:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_build_tuplestore_recursively_5), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	F_errdetail(m, int32(_a_F_build_tuplestore_recursively_6), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_build_tuplestore_recursively_7), int32(1481), int32(_a_F_build_tuplestore_recursively_8))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_build_tuplestore_recursively_9), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L7
	} else {
		goto L100
	}
L100:
	;
	v400 = F_format_type_with_typemod(m, v154, v143)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L7
	} else {
		goto L101
	}
L101:
	;
	v402 = F_format_type_with_typemod(m, v153, v148)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v400
	F_errdetail(m, int32(_a_F_build_tuplestore_recursively_10), v28+int32(96))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_build_tuplestore_recursively_7), int32(1498), int32(_a_F_build_tuplestore_recursively_8))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L7
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_build_tuplestore_recursively_9), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	v427 = F_format_type_with_typemod(m, v164, v157)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L7
	} else {
		goto L108
	}
L108:
	;
	v429 = F_format_type_with_typemod(m, v163, v158)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v427
	F_errdetail(m, int32(_a_F_build_tuplestore_recursively_11), v28+int32(80))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(_a_F_build_tuplestore_recursively_7), int32(1511), int32(_a_F_build_tuplestore_recursively_8))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L7
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_errcode(m, int32(151388292))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	F_errmsg(m, int32(_a_F_build_tuplestore_recursively_12), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L7
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_build_tuplestore_recursively_7), int32(1340), int32(_a_F_build_tuplestore_recursively_13))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplestore_begin_heap(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v2 = l1
	v7 = F_palloc0(m, int32(120))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v13 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v7)+10)) = uint16(v13)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+9)) = uint8(v2)
		if l0 != 0 {
			v18 = int32(12)
		} else {
			v18 = int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v20
		v26 = base.I64_extend_i32_s(l2) << (uint(int64(10)) % 64)
		*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = v26
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v26
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_begin_heap[0]))
		v35 = F_GenerationContextCreate(m, v30, int32(_a_F_tuplestore_begin_heap_0), v20, int32(_a_F_tuplestore_begin_heap_1), int32(_a_F_tuplestore_begin_heap_2))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v35
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_begin_heap[1]))
			v40 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+76)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v7)+56)) = v39
			v43 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v7)+88)) = uint8(v43)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+84)) = int32(_a_F_tuplestore_begin_heap_3)
			*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = v40
			v50 = F_palloc(m, int32(_a_F_tuplestore_begin_heap_4))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+72)) = v50
				v53 = F_GetMemoryChunkSpace(m, v50)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+104)) = int32(8)
					*(*int64)(unsafe.Add(mBase, uint32(v7)+96)) = int64(4294967296)
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v7)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v59 - base.I64_extend_i32_u(v53)
					v64 = F_palloc(m, int32(192))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+92)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v64))) = v18
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						v69 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v68)+4)) = uint8(v69)
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)+92))
						*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v69
						*(*int32)(unsafe.Add(mBase, uint32(v7)+68)) = int32(1851)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+64)) = int32(1852)
						*(*int32)(unsafe.Add(mBase, uint32(v7)+60)) = int32(1853)
						return v7
					}
				}
			}
		}
	}
}
func F_tuplestore_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_tuplestore_gettuple(m, l0, l1, v8+int32(15))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if l2 == int32(0) {
				v25 = v16
				v26 = v12
				v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(16)
					return base.B2i32(v12 != int32(0))
				}
			} else {
				if v16&int32(1) != 0 {
					v25 = v16
					v26 = v12
					v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return base.B2i32(v12 != int32(0))
					}
				} else {
					v22 = F_heap_copy_minimal_tuple(m, v12, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = int32(1)
						v26 = v22
						v29 = F_ExecStoreMinimalTuple(m, v26, l3, v25&int32(1))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return base.B2i32(v12 != int32(0))
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
			m.T0[v32].(func(*base.Module, int32))(m, l3)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(16)
				return base.B2i32(v12 != int32(0))
			}
		}
	}
}
func F_tuplestore_rescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v7 = v3 + v4*int32(24)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v8 {
	case 0:
		v53 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v53
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v53)
		return
	case 1:
		*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(0)
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v11)
		return
	case 2:
		v15 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)) = uint8(v15)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = F_BufFileSeek(m, v17, v15, int64(0), v15)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 == int32(0) {
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_tuplestore_rescan_0), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_tuplestore_rescan_1), int32(1308), int32(_a_F_tuplestore_rescan_2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_tuplestore_rescan_3), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_tuplestore_rescan_1), int32(1311), int32(_a_F_tuplestore_rescan_2))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
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
func F_tuplestore_skiptuples(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v98 int32
	_ = v98
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 <= int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v98
L2:
	;
	v98 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v98 = int32(0)
	goto L1
L6:
	;
	v17 = l1
	goto L9
L7:
	;
	goto L8
L8:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v49 = v45 + v46*int32(24)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
	if l2 != 0 {
		goto L23
	} else {
		goto L24
	}
L9:
	;
	v25 = F_tuplestore_gettuple(m, l0, l2, v10+int32(15))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v98 = int32(1)
	goto L1
L11:
	;
	return int32(0)
L12:
	;
	if v25 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
	if v31 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_pfree(m, v25)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_tuplestore_skiptuples[0]))
	if v37 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v40 = int64(1)
	if base.Ui64(v40) < base.Ui64(v17) {
		v17 = v17 - v40
		goto L9
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L10
L23:
	;
	if v50&int32(1) != 0 {
		v98 = v4
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v50&int32(1) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	if l1 <= base.I64_extend_i32_s(v53-v54) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v54 + base.I32_wrap_i64(l1)
	v98 = int32(1)
	goto L1
L28:
	;
	goto L29
L29:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v62)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v53
	v98 = v4
	goto L1
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v76 < base.I64_extend_i32_s(v77-v78) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v76 = l1
	v77 = v69
	goto L30
L32:
	;
	goto L33
L33:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)) = uint8(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v70
	v76 = l1 - int64(1)
	v77 = v70
	goto L30
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v77 - base.I32_wrap_i64(v76)
	v98 = int32(1)
	goto L1
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = v78
	goto L5
}
