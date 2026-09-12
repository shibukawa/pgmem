package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_btree_xlog_cleanup(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[136])) = int32(0)
		return
	}
}
func F_btree_xlog_split(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v365 int32
	_ = v365
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	v3 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+64))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	F_XLogRecGetBlockTag(m, l1, v3, v3, v3, v24+int32(20))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v38 = int32(0)
	F_XLogRecGetBlockTag(m, l1, int32(1), v38, v38, v24+int32(16))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v45 = int32(0)
	v48 = v24 + int32(12)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+72))
	if v52 < int32(2) {
		v76 = v45
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v76 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	v58 = v51 + int32(180)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v59 != int32(1) {
		v76 = v45
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L9
L9:
	;
	goto L10
L10:
	;
	goto L12
L12:
	;
	goto L13
L13:
	;
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v68
	goto L16
L15:
	;
	goto L16
L16:
	;
	v76 = int32(1)
	goto L5
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(0)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v29 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v134 = F_XLogInitBufferForRedo(m, l1, int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L34
	}
L22:
	;
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	v87 = F_XLogReadBufferForRedo(m, l1, int32(3), v24+int32(28))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v87 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v91 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v125 == int32(0) {
		goto L21
	} else {
		goto L32
	}
L27:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+16)))
	v111 = v110 + v109
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+12)))
	v114 = v112 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v111)+12)) = uint16(v114)
	*(*uint32)(unsafe.Add(mBase, uint32(v109)+4)) = uint32(v83)
	v118 = int64(base.Ui64(v83) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v109))) = uint32(v118)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_MarkBufferDirty(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v91^int32(-1))<<(uint(int32(2))%32))))
	v109 = v101
	goto L27
L29:
	;
	goto L30
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v109 = v103 + v91<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	goto L26
L32:
	;
	F_UnlockReleaseBuffer(m, v125)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L21
L34:
	;
	v138 = v24 + int32(24)
	v139 = int32(0)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+72))
	if v141 < int32(1) {
		v163 = v139
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v134 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v166 = v163
	goto L35
L37:
	;
	v147 = v140 + int32(128)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v148 != int32(1) {
		v163 = v139
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+43)))
	if v151 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v138 == int32(0) {
		v163 = v139
		goto L36
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v138 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v156 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v156
	v166 = v156
	goto L35
L43:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v159
	goto L45
L44:
	;
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v163 = v161
	goto L36
L46:
	;
	F_PageInit(m, v184, int32(8192), int32(16))
	mBase = m.M
	goto L50
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+(v134^int32(-1))<<(uint(int32(2))%32))))
	v184 = v176
	goto L46
L48:
	;
	goto L49
L49:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v184 = v178 + v134<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+16)))
	v189 = v184 + v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+4)) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v195 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+14)) = uint16(v195)
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+12)) = uint16(base.B2i32(v29 == v195))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+8)) = v194
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	F__bt_restore_page(m, v184, v166, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = base.I64_rotr(v26, int64(32))
	F_MarkBufferDirty(m, v134)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v211 = base.I32_wrap_i64(int64(base.Ui64(v26) >> (uint(int64(32)) % 64)))
	v212 = base.I32_wrap_i64(v26)
	v216 = F_XLogReadBufferForRedo(m, l1, int32(0), v24+int32(28))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L57
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L150
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L147
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L144
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L141
	}
L57:
	;
	if v216 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v220 < int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v500 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L61:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+16)))
	v240 = int32(0)
	v243 = v24 + int32(24)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+72))
	if v246 < v240 {
		v268 = v240
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v220^int32(-1))<<(uint(int32(2))%32))))
	v238 = v230
	goto L61
L63:
	;
	goto L64
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v238 = v232 + v220<<(uint(int32(13))%32) + int32(-8192)
	goto L61
L65:
	;
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	if l0 != 0 {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v271 = v268
	goto L65
L67:
	;
	v252 = v245 + int32(76)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	if v253 != int32(1) {
		v268 = v240
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252)+43)))
	if v256 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v243 == int32(0) {
		v268 = v240
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v243 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v261 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v261
	v271 = v261
	goto L65
L73:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v243))) = v264
	goto L75
L74:
	;
	goto L75
L75:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v252)+44))
	v268 = v266
	goto L66
L76:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309)+6)))
	v322 = (v316&int32(8191) + int32(7)) & int32(16376)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v315 - v322
	v325 = F_PageGetTempPageCopySpecial(m, v238)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L85
	}
L77:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+6)))
	v283 = (v277&int32(8191) + int32(7)) & int32(16376)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v276 - v283
	v286 = v271 + v283
	if l0 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	if v272&int32(65535) != 0 {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v309 = v271
	v311 = v240
	v312 = v3
	v313 = v3
	v314 = int32(0)
	goto L76
L80:
	;
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	v293 = F_CopyIndexTuple(m, v271)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	if v272&int32(65535) != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v309 = v286
	v311 = v240
	v312 = v283
	v313 = v271
	v314 = int32(0)
	goto L76
L83:
	;
	v298 = (v292 - int32(1)) & int32(65535)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298<<(uint(int32(2))%32)+v238)+20))
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	v307 = F__bt_swap_posting(m, v293, v238+v302&int32(32767), v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v309 = v286
	v311 = v298
	v312 = v283
	v313 = v293
	v314 = v307
	goto L76
L85:
	;
	v329 = F_PageAddItemExtended(m, v325, v309, v322, int32(1), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v329 == int32(0) {
		goto L56
	} else {
		goto L87
	}
L87:
	;
	v333 = int32(2)
	v336 = v238 + v239
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+4))
	if v337 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v338 = v333
	goto L90
L89:
	;
	v338 = int32(1)
	goto L90
L90:
	;
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if base.Ui32(v338) < base.Ui32(v339) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v345 = v338
	v347 = v333
	goto L94
L92:
	;
	v432 = v338
	v434 = v333
	goto L93
L93:
	;
	if l0 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L94:
	;
	if v345 == v311 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v432 = v427
	v434 = v425
	goto L93
L96:
	;
	v424 = int32(1)
	v425 = v422 + v424
	v427 = v345 + v424
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if base.Ui32(v427) < base.Ui32(v428) {
		v345 = v427
		v347 = v425
		goto L94
	} else {
		goto L112
	}
L97:
	;
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v314)+6)))
	v375 = F_PageAddItemExtended(m, v325, v314, (v365&int32(8191)+int32(7))&int32(16376), v347&int32(65535), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	if l0 == int32(0) {
		v403 = v347
		goto L105
	} else {
		goto L106
	}
L100:
	;
	if v375 != 0 {
		v422 = v347
		goto L96
	} else {
		goto L101
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg_internal(m, int32(102267), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(496253), int32(390), int32(102154))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
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
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v345<<(uint(int32(2))%32)+(v238+int32(24))-int32(4))))
	v418 = F_PageAddItemExtended(m, v325, v238+v409&int32(32767), int32(base.Ui32(v409)>>(uint(int32(17))%32)), v403&int32(65535), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L110
	}
L106:
	;
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	if v345 != v392 {
		v403 = v347
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v397 = F_PageAddItemExtended(m, v325, v313, v312, v347&int32(65535), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v397 == int32(0) {
		goto L55
	} else {
		goto L109
	}
L109:
	;
	v403 = v347 + int32(1)
	goto L105
L110:
	;
	if v418 == int32(0) {
		goto L54
	} else {
		goto L111
	}
L111:
	;
	v422 = v403
	goto L96
L112:
	;
	goto L95
L113:
	;
	F_PageRestoreTempPage(m, v325, v238)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L118
	}
L114:
	;
	v453 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	if v453 != v432&int32(65535) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v460 = F_PageAddItemExtended(m, v325, v313, v312, v434&int32(65535), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v460 == int32(0) {
		goto L53
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	if v29 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v468 = int32(128)
	goto L121
L120:
	;
	v468 = int32(129)
	goto L121
L121:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v336)+12)) = uint16(v468)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v471 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v336)+14)) = uint16(v471)
	*(*int32)(unsafe.Add(mBase, uint32(v336)+4)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v211
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_MarkBufferDirty(m, v476)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L60
L123:
	;
	F_UnlockReleaseBuffer(m, v134)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L136
	}
L124:
	;
	v506 = F_XLogReadBufferForRedo(m, l1, int32(2), v24+int32(8))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v506 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v510 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v539 == int32(0) {
		goto L123
	} else {
		goto L134
	}
L129:
	;
	v529 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v528)+16)))
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v529+v528))) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v528)+4)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = v211
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_MarkBufferDirty(m, v535)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v514+(v510^int32(-1))<<(uint(int32(2))%32))))
	v528 = v520
	goto L129
L131:
	;
	goto L132
L132:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v528 = v522 + v510<<(uint(int32(13))%32) + int32(-8192)
	goto L129
L133:
	;
	goto L128
L134:
	;
	F_UnlockReleaseBuffer(m, v539)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L123
L136:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v547 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_UnlockReleaseBuffer(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	m.G0 = v24 + int32(32)
	return
L140:
	;
	goto L139
L141:
	;
	F_errmsg_internal(m, int32(102171), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(496253), int32(373), int32(102154))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errmsg_internal(m, int32(102219), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(496253), int32(400), int32(102154))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	F_errmsg_internal(m, int32(102328), int32(0))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(496253), int32(409), int32(102154))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	F_errmsg_internal(m, int32(102219), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(496253), int32(418), int32(102154))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btree_xlog_updates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v15 = l2
	v19 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L16
	}
L2:
	;
	v23 = int32(1)
	v25 = l1 + v19<<(uint(v23)%32)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26<<(uint(int32(2))%32)+(l0+int32(24))-int32(4))))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	v38 = F_palloc(m, v33<<(uint(v23)%32)+int32(8))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	return
L5:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = l0 + v32&int32(32767)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)) = uint16(v40)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)) = uint16(v46)
	v51 = v15 + int32(2)
	v53 = v46 << (uint(int32(1)) % 32)
	if v53 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F__bt_update_posting(m, v38)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L10
	}
L7:
	;
	v54 = F__emscripten_memcpy_bulkmem(m, v38+int32(8), v51, v53)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+6)))
	v67 = F_PageIndexTupleOverwrite(m, l0, v58, v59, (v60&int32(8191)+int32(7))&int32(16376))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	if v67 == int32(0) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_pfree(m, v38)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	v77 = int32(1)
	v81 = v19 + v77
	if v81 != l3 {
		v15 = v51 + v76<<(uint(v77)%32)
		v19 = v81
		goto L2
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	F_errmsg_internal(m, int32(290297), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(496253), int32(585), int32(160077))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
