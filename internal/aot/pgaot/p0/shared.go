package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteSharedSecurityLabel(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn14211(m, l0, l1, int32(3593), int32(3592))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_checkSharedDependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
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
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(240)
	m.G0 = v18
	goto L2
L1:
	;
	if v225 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L2:
	;
	if base.B2i32(base.B2i32(l0 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_checkSharedDependencies_0)) < base.Ui32(l1)) == v5)&((base.B2i32(l0 != int32(2615))|base.B2i32(l1 != int32(2200)))&base.B2i32(l0 != int32(1262))) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v40 = F_palloc(m, int32(2560))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v296 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = l0
	F_errstart_cold(m, int32(21), v296)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L62
	}
L6:
	;
	return int32(0)
L7:
	;
	F_initStringInfo(m, v18+int32(100))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_initStringInfo(m, v18+int32(84))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v54 = F_table_open(m, int32(1214), int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v57 = v18 + int32(128)
	F_ScanKeyInit(m, v57, int32(5), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v64 = int32(184)
	F_ScanKeyInit(m, v18+v64, int32(6), int32(3), v64, base.I64_extend_i32_u(l1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v76 = F_systable_beginscan(m, v54, int32(1233), int32(1), int32(0), int32(2), v57)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v78 = F_systable_getnext(m, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v78 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v82 = v78
	v87 = int32(128)
	v88 = v5
	v90 = v5
	v91 = v40
	goto L18
L16:
	;
	v223 = v5
	v225 = v5
	v226 = v40
	goto L17
L17:
	;
	F_systable_endscan(m, v76)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L45
	}
L18:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v98 = v96 + v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v103
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v105 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v223 = v206
	v225 = v208
	v226 = v209
	goto L17
L20:
	;
	v214 = F_systable_getnext(m, v76)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L6
	} else {
		goto L43
	}
L21:
	;
	v191 = F_palloc(m, int32(8))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L41
	}
L22:
	;
	v148 = int32(0)
	goto L35
L23:
	;
	if v87 <= v88 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_checkSharedDependencies[0]))
	if v105 == v109 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v90 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v113 <= int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v116 = int32(0)
	if v116 < v113 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v119 = v113
	goto L30
L29:
	;
	v119 = v116
	goto L30
L30:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	goto L22
L31:
	;
	v125 = F_repalloc(m, v91, v87*int32(40))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	v129 = v87
	v130 = v91
	goto L33
L33:
	;
	v133 = v130 + v88*int32(20)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+8)) = v134
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int64)(unsafe.Add(mBase, uint32(v133))) = v136
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v133)+12)) = uint8(v138)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_checkSharedDependencies[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+16)) = base.B2i32(v140 != v142)
	v205 = v129
	v206 = v88 + int32(1)
	v208 = v90
	v209 = v130
	goto L20
L34:
	;
	v129 = v87 << (uint(int32(1)) % 32)
	v130 = v125
	goto L33
L35:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v120+v148<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v105 != v166 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v171 + int32(1)
	v205 = v87
	v206 = v88
	v208 = v90
	v209 = v91
	goto L20
L37:
	;
	v169 = v148 + int32(1)
	if v119 != v169 {
		v148 = v169
		goto L35
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L36
L40:
	;
	goto L21
L41:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v191))) = v193
	v197 = F_lappend(m, v90, v191)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v205 = v87
	v206 = v88
	v208 = v197
	v209 = v91
	goto L20
L43:
	;
	if v214 != 0 {
		v82 = v214
		v87 = v205
		v88 = v206
		v90 = v208
		v91 = v209
		goto L18
	} else {
		goto L44
	}
L44:
	;
	goto L19
L45:
	;
	F_relation_close(m, v54, int32(1))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if int32(2) <= v223 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v248 = int32(0)
	v251 = v248
	v252 = v248
	v257 = v248
	goto L53
L48:
	;
	F_pg_qsort(m, v226, v223, int32(20), int32(510))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v242 = int32(0)
	if v223 != int32(1) {
		v321 = v242
		v327 = v242
		goto L1
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	goto L47
L53:
	;
	if v251 <= int32(99) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v321 = v281
	v327 = v283
	goto L1
L55:
	;
	v288 = v226 + v252*int32(20)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	v290 = int32(*(*int8)(unsafe.Add(mBase, uint32(v288)+12)))
	F_storeObjectDescription(m, v18+int32(84), v289, v288, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L60
	}
L56:
	;
	v272 = v226 + v252*int32(20)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)+16))
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v272)+12)))
	F_storeObjectDescription(m, v18+int32(100), v273, v272, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v281 = v251
	v283 = v257 + int32(1)
	goto L55
L59:
	;
	v281 = v251 + int32(1)
	v283 = v257
	goto L55
L60:
	;
	v294 = v252 + int32(1)
	if v294 != v223 {
		v251 = v281
		v252 = v294
		v257 = v283
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v310 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v310
	F_errmsg(m, int32(_a_F_checkSharedDependencies_1), v18)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_checkSharedDependencies_2), int32(704), int32(_a_F_checkSharedDependencies_3))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_pfree(m, v226)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L103
	}
L68:
	;
	v447 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v339 = int32(0)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v340 <= v339 {
		v447 = v339
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v344 = v321
	v345 = int32(0)
	v351 = v339
	goto L72
L72:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v345<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(1262)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v366
	if int32(100) <= v344 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v447 = v406
	goto L67
L74:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	v413 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L6
	} else {
		goto L89
	}
L75:
	;
	v405 = v344
	v406 = v351 + int32(1)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v375 = v344 + int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v363)+4))
	v380 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v380 == int32(0) {
		v405 = v375
		v406 = v351
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v384 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_appendStringInfoChar(m, v18+int32(100), int32(10))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v376
	if v376 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	v398 = int32(_a_F_checkSharedDependencies_4)
	goto L86
L85:
	;
	v398 = int32(_a_F_checkSharedDependencies_5)
	goto L86
L86:
	;
	F_appendStringInfo(m, v18+int32(100), v398, v18-int32(-64))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_pfree(m, v380)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v405 = v375
	v406 = v351
	goto L74
L89:
	;
	if v413 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v415 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v437 = v345 + int32(1)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v437 < v438 {
		v344 = v405
		v345 = v437
		v351 = v406
		goto L72
	} else {
		goto L102
	}
L93:
	;
	F_appendStringInfoChar(m, v18+int32(84), int32(10))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v409
	if v409 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v429 = int32(_a_F_checkSharedDependencies_4)
	goto L99
L98:
	;
	v429 = int32(_a_F_checkSharedDependencies_5)
	goto L99
L99:
	;
	F_appendStringInfo(m, v18+int32(84), v429, v18+int32(48))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v413)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	goto L92
L102:
	;
	goto L73
L103:
	;
	F_list_free_deep(m, v225)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v459 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	m.G0 = v18 + int32(240)
	return base.B2i32(v459 != int32(0))
L106:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	F_pfree(m, v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if int32(0) < v327 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	F_pfree(m, v465)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v468 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v468
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v468
	goto L105
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v327
	if v327 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if int32(0) < v447 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	v481 = int32(_a_F_checkSharedDependencies_6)
	goto L116
L115:
	;
	v481 = int32(_a_F_checkSharedDependencies_7)
	goto L116
L116:
	;
	F_appendStringInfo(m, v18+int32(100), v481, v18+int32(32))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v447
	if v447 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v500
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v502
	goto L105
L121:
	;
	v495 = int32(_a_F_checkSharedDependencies_8)
	goto L123
L122:
	;
	v495 = int32(_a_F_checkSharedDependencies_9)
	goto L123
L123:
	;
	F_appendStringInfo(m, v18+int32(100), v495, v18+int32(16))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L120
}
func F_shared_record_table_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v5 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v15 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_dsa_get_address(m, l3, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v13
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v14 = v9
	goto L1
L7:
	;
	v23 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v27 != v28 {
		v79 = v23
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = F_dsa_get_address(m, l3, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = v21
	goto L7
L11:
	;
	v22 = v19
	goto L7
L12:
	;
	return v79 ^ int32(1)
L13:
	;
	goto L12
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v30 != v31 {
		v79 = v23
		goto L13
	} else {
		goto L15
	}
L15:
	;
	if v27 <= int32(0) {
		v79 = int32(1)
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v37 = v27 << (uint(int32(3)) % 32)
	v39 = int32(28)
	v45 = int32(0)
	goto L17
L17:
	;
	v52 = v45 * int32(100)
	v53 = v14 + v37 + v39 + v52
	v54 = int32(4)
	v56 = v52 + (v22 + v37 + v39)
	v59 = F_strcmp(m, v53+v54, v56+v54)
	mBase = m.M
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v79 = int32(0)
	goto L13
L19:
	;
	goto L18
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v56)+68))
	if v60 != v61 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	if v63 != v64 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v53)+96))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56)+96))
	if v66 != v67 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+91)))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+91)))
	if v69 != v70 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v72 = int32(1)
	v74 = v45 + v72
	if v27 != v74 {
		v45 = v74
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v79 = v72
	goto L13
}
func F_shared_ts_free_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	v4 = int32(0)
	F_check_stack_depth(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_dsa_get_address(m, v11, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L8
	}
L3:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v170, l1)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L53
	}
L4:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v128 == int32(0) {
		goto L3
	} else {
		goto L42
	}
L5:
	;
	v95 = v4
	goto L31
L6:
	;
	v59 = v4
	goto L20
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	switch v14 {
	case 0:
		goto L4
	case 1:
		goto L7
	case 2:
		goto L6
	case 3:
		goto L5
	default:
		goto L3
	}
L9:
	;
	v28 = v4
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v28<<(uint(int32(2))%32))))
	if base.B2i32(l2 <= int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L3
L12:
	;
	v46 = v28 + int32(1)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if base.Ui32(v46) < base.Ui32(v47) {
		v28 = v46
		goto L10
	} else {
		goto L19
	}
L13:
	;
	F_shared_ts_free_recurse(m, l0, v35, l2-int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v35&int32(1) != 0 {
		goto L12
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v42, v35)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	goto L11
L20:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+(v12+int32(12))))))
	if v64 == int32(255) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L3
L22:
	;
	v82 = v59 + int32(1)
	if v82 != int32(256) {
		v59 = v82
		goto L20
	} else {
		goto L30
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(268)+v64<<(uint(int32(2))%32))))
	if int32(0) < l2 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_shared_ts_free_recurse(m, l0, v70, l2-int32(8))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v70&int32(1) != 0 {
		goto L22
	} else {
		goto L28
	}
L27:
	;
	goto L22
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v77, v70)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	goto L22
L30:
	;
	goto L21
L31:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)+int32(base.Ui32(v95)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v104)>>(uint(v95)%32))&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L3
L33:
	;
	v125 = v95 + int32(1)
	if v125 != int32(256) {
		v95 = v125
		goto L31
	} else {
		goto L41
	}
L34:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(36)+v95<<(uint(int32(2))%32))))
	if int32(0) < l2 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_shared_ts_free_recurse(m, l0, v113, l2-int32(8))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v113&int32(1) != 0 {
		goto L33
	} else {
		goto L39
	}
L38:
	;
	goto L33
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v120, v113)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	goto L32
L42:
	;
	v131 = int32(8)
	v141 = v4
	goto L43
L43:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v12+v131+v141<<(uint(int32(2))%32))))
	if base.B2i32(l2 <= int32(0)) == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L3
L45:
	;
	v159 = v141 + int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	if base.Ui32(v159) < base.Ui32(v160) {
		v141 = v159
		goto L43
	} else {
		goto L52
	}
L46:
	;
	F_shared_ts_free_recurse(m, l0, v148, l2-v131)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v148&int32(1) != 0 {
		goto L45
	} else {
		goto L50
	}
L49:
	;
	goto L45
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_dsa_free(m, v155, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	goto L44
L53:
	;
	return
}
