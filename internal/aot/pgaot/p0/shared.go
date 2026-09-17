package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_DeleteSharedSecurityLabel(m *base.Module, l0 int32, l1 int32) {
	var v6 int32
	_ = v6
	Fn13826(m, l0, l1, int32(3593), int32(3592))
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
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(224)
	m.G0 = v18
	goto L2
L1:
	;
	if v223 == int32(0) {
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
	v294 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = l0
	F_errstart_cold(m, int32(21), v294)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
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
	F_ScanKeyInit(m, v57, int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_ScanKeyInit(m, v18+int32(176), int32(6), int32(3), int32(184), l1)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v74 = F_systable_beginscan(m, v54, int32(1233), int32(1), int32(0), int32(2), v57)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v76 = F_systable_getnext(m, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v76 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v80 = v76
	v85 = int32(128)
	v86 = v5
	v88 = v5
	v89 = v40
	goto L18
L16:
	;
	v221 = v5
	v223 = v5
	v224 = v40
	goto L17
L17:
	;
	F_systable_endscan(m, v74)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
	} else {
		goto L45
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+22)))
	v96 = v94 + v95
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v103 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	v221 = v204
	v223 = v206
	v224 = v207
	goto L17
L20:
	;
	v212 = F_systable_getnext(m, v74)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L43
	}
L21:
	;
	v189 = F_palloc(m, int32(8))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L41
	}
L22:
	;
	v146 = int32(0)
	goto L35
L23:
	;
	if v85 <= v86 {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_checkSharedDependencies[0]))
	if v103 == v107 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v88 == int32(0) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v111 <= int32(0) {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v114 = int32(0)
	if v114 < v111 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v117 = v111
	goto L30
L29:
	;
	v117 = v114
	goto L30
L30:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	goto L22
L31:
	;
	v123 = F_repalloc(m, v89, v85*int32(40))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L6
	} else {
		goto L34
	}
L32:
	;
	v127 = v85
	v128 = v89
	goto L33
L33:
	;
	v131 = v128 + v86*int32(20)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v18)+116))
	*(*int64)(unsafe.Add(mBase, uint32(v131))) = v134
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+12)) = uint8(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_checkSharedDependencies[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = base.B2i32(v138 != v140)
	v203 = v127
	v204 = v86 + int32(1)
	v206 = v88
	v207 = v128
	goto L20
L34:
	;
	v127 = v85 << (uint(int32(1)) % 32)
	v128 = v123
	goto L33
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v118+v146<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v103 != v164 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v169 + int32(1)
	v203 = v85
	v204 = v86
	v206 = v88
	v207 = v89
	goto L20
L37:
	;
	v167 = v146 + int32(1)
	if v117 != v167 {
		v146 = v167
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
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v191
	v195 = F_lappend(m, v88, v189)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v203 = v85
	v204 = v86
	v206 = v195
	v207 = v89
	goto L20
L43:
	;
	if v212 != 0 {
		v80 = v212
		v85 = v203
		v86 = v204
		v88 = v206
		v89 = v207
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if int32(2) <= v221 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v246 = int32(0)
	v249 = v246
	v250 = v246
	v255 = v246
	goto L53
L48:
	;
	F_pg_qsort(m, v224, v221, int32(20), int32(475))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v240 = int32(0)
	if v221 != int32(1) {
		v319 = v240
		v325 = v240
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
	if v249 <= int32(99) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v319 = v279
	v325 = v281
	goto L1
L55:
	;
	v286 = v224 + v250*int32(20)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+16))
	v288 = int32(*(*int8)(unsafe.Add(mBase, uint32(v286)+12)))
	F_storeObjectDescription(m, v18+int32(84), v287, v286, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L6
	} else {
		goto L60
	}
L56:
	;
	v270 = v224 + v250*int32(20)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+16))
	v272 = int32(*(*int8)(unsafe.Add(mBase, uint32(v270)+12)))
	F_storeObjectDescription(m, v18+int32(100), v271, v270, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v279 = v249
	v281 = v255 + int32(1)
	goto L55
L59:
	;
	v279 = v249 + int32(1)
	v281 = v255
	goto L55
L60:
	;
	v292 = v250 + int32(1)
	if v292 != v221 {
		v249 = v279
		v250 = v292
		v255 = v281
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
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	v308 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v308
	F_errmsg(m, int32(_a_F_checkSharedDependencies_1), v18)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_checkSharedDependencies_2), int32(704), int32(_a_F_checkSharedDependencies_3))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
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
	F_pfree(m, v224)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L6
	} else {
		goto L103
	}
L68:
	;
	v445 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v337 = int32(0)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v338 <= v337 {
		v445 = v337
		goto L67
	} else {
		goto L71
	}
L71:
	;
	v342 = v319
	v343 = int32(0)
	v349 = v337
	goto L72
L72:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v343<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = int32(1262)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v364
	if int32(100) <= v342 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	v445 = v404
	goto L67
L74:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v411 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L6
	} else {
		goto L89
	}
L75:
	;
	v403 = v342
	v404 = v349 + int32(1)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v373 = v342 + int32(1)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v378 = F_getObjectDescription(m, v18+int32(116), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L6
	} else {
		goto L78
	}
L78:
	;
	if v378 == int32(0) {
		v403 = v373
		v404 = v349
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v382 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_appendStringInfoChar(m, v18+int32(100), int32(10))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L6
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v374
	if v374 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	goto L82
L84:
	;
	v396 = int32(_a_F_checkSharedDependencies_4)
	goto L86
L85:
	;
	v396 = int32(_a_F_checkSharedDependencies_5)
	goto L86
L86:
	;
	F_appendStringInfo(m, v18+int32(100), v396, v18-int32(-64))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	F_pfree(m, v378)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v403 = v373
	v404 = v349
	goto L74
L89:
	;
	if v411 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v413 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	v435 = v343 + int32(1)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	if v435 < v436 {
		v342 = v403
		v343 = v435
		v349 = v404
		goto L72
	} else {
		goto L102
	}
L93:
	;
	F_appendStringInfoChar(m, v18+int32(84), int32(10))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v407
	if v407 == int32(1) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v427 = int32(_a_F_checkSharedDependencies_4)
	goto L99
L98:
	;
	v427 = int32(_a_F_checkSharedDependencies_5)
	goto L99
L99:
	;
	F_appendStringInfo(m, v18+int32(84), v427, v18+int32(48))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_pfree(m, v411)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
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
	F_list_free_deep(m, v223)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v457 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	m.G0 = v18 + int32(224)
	return base.B2i32(v457 != int32(0))
L106:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	F_pfree(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	if int32(0) < v325 {
		goto L111
	} else {
		goto L112
	}
L109:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	F_pfree(m, v463)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	v466 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v466
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v466
	goto L105
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v325
	if v325 == int32(1) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	if int32(0) < v445 {
		goto L118
	} else {
		goto L119
	}
L114:
	;
	v479 = int32(_a_F_checkSharedDependencies_6)
	goto L116
L115:
	;
	v479 = int32(_a_F_checkSharedDependencies_7)
	goto L116
L116:
	;
	F_appendStringInfo(m, v18+int32(100), v479, v18+int32(32))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L6
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v445
	if v445 == int32(1) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v18)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v498
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v500
	goto L105
L121:
	;
	v493 = int32(_a_F_checkSharedDependencies_8)
	goto L123
L122:
	;
	v493 = int32(_a_F_checkSharedDependencies_9)
	goto L123
L123:
	;
	F_appendStringInfo(m, v18+int32(100), v493, v18+int32(16))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
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
	v37 = v27 << (uint(int32(4)) % 32)
	v39 = int32(20)
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
