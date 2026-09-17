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
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_cleanup[0]))
	F_MemoryContextDelete(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_cleanup[0])) = int32(0)
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
	var v57 int32
	_ = v57
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
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
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
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
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
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51+int32(104))+76)))
	if v57 != int32(1) {
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v51+int32(180))+20))
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
	v114 = v112 & int32(_a_F_btree_xlog_split_0)
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
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v91^int32(-1))<<(uint(int32(2))%32))))
	v109 = v101
	goto L27
L29:
	;
	goto L30
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
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
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+int32(52))+76)))
	if v146 != int32(1) {
		v163 = v139
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v150 = v140 + int32(128)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+43)))
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
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v159
	goto L45
L44:
	;
	goto L45
L45:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v150)+44))
	v163 = v161
	goto L36
L46:
	;
	F_PageInit(m, v184, int32(_a_F_btree_xlog_split_1), int32(16))
	mBase = m.M
	goto L50
L47:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+(v134^int32(-1))<<(uint(int32(2))%32))))
	v184 = v176
	goto L46
L48:
	;
	goto L49
L49:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
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
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L150
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L147
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L144
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
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
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	if v496 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L61:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+16)))
	v240 = int32(0)
	v242 = v24 + int32(24)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+72))
	if v245 < v240 {
		v267 = v240
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v224+(v220^int32(-1))<<(uint(int32(2))%32))))
	v238 = v230
	goto L61
L63:
	;
	goto L64
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v238 = v232 + v220<<(uint(int32(13))%32) + int32(-8192)
	goto L61
L65:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	if v271|l0 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v270 = v267
	goto L65
L67:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244+int32(0))+76)))
	if v250 != int32(1) {
		v267 = v240
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v254 = v244 + int32(76)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+43)))
	if v255 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v242 == int32(0) {
		v267 = v240
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v242 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v260 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v260
	v270 = v260
	goto L65
L73:
	;
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v263
	goto L75
L74:
	;
	goto L75
L75:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v254)+44))
	v267 = v265
	goto L66
L76:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+6)))
	v320 = (v314&int32(_a_F_btree_xlog_split_2) + int32(7)) & int32(_a_F_btree_xlog_split_3)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v313 - v320
	v323 = F_PageGetTempPageCopySpecial(m, v238)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L85
	}
L77:
	;
	v307 = v270
	v309 = v3
	v310 = v3
	v311 = v3
	v312 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270)+6)))
	v283 = (v277&int32(_a_F_btree_xlog_split_2) + int32(7)) & int32(_a_F_btree_xlog_split_3)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v276 - v283
	v286 = v270 + v283
	if l0 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	v291 = F_CopyIndexTuple(m, v270)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	if v271 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v307 = v286
	v309 = v3
	v310 = v283
	v311 = v270
	v312 = int32(0)
	goto L76
L83:
	;
	v296 = (v290 - int32(1)) & int32(_a_F_btree_xlog_split_4)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v238+v296<<(uint(int32(2))%32))+20))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+8)))
	v305 = F__bt_swap_posting(m, v291, v238+v300&int32(_a_F_btree_xlog_split_5), v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v307 = v286
	v309 = v296
	v310 = v283
	v311 = v291
	v312 = v305
	goto L76
L85:
	;
	v327 = F_PageAddItemExtended(m, v323, v307, v320, int32(1), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v327 == int32(0) {
		goto L56
	} else {
		goto L87
	}
L87:
	;
	v331 = int32(2)
	v334 = v238 + v239
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v335 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v336 = v331
	goto L90
L89:
	;
	v336 = int32(1)
	goto L90
L90:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if base.Ui32(v336) < base.Ui32(v337) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v343 = v336
	v345 = v331
	goto L94
L92:
	;
	v428 = v336
	v430 = v331
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
	if v343 == v309 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v428 = v423
	v430 = v421
	goto L93
L96:
	;
	v420 = int32(1)
	v421 = v418 + v420
	v423 = v343 + v420
	v424 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if base.Ui32(v423) < base.Ui32(v424) {
		v343 = v423
		v345 = v421
		goto L94
	} else {
		goto L112
	}
L97:
	;
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v312)+6)))
	v373 = F_PageAddItemExtended(m, v323, v312, (v363&int32(_a_F_btree_xlog_split_2)+int32(7))&int32(_a_F_btree_xlog_split_3), v345&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
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
		v401 = v345
		goto L105
	} else {
		goto L106
	}
L100:
	;
	if v373 != 0 {
		v418 = v345
		goto L96
	} else {
		goto L101
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_6), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(390), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
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
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v238+int32(20)+v343<<(uint(int32(2))%32))))
	v414 = F_PageAddItemExtended(m, v323, v238+v405&int32(_a_F_btree_xlog_split_5), int32(base.Ui32(v405)>>(uint(int32(17))%32)), v401&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L110
	}
L106:
	;
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	if v343 != v390 {
		v401 = v345
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v395 = F_PageAddItemExtended(m, v323, v311, v310, v345&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v395 == int32(0) {
		goto L55
	} else {
		goto L109
	}
L109:
	;
	v401 = v345 + int32(1)
	goto L105
L110:
	;
	if v414 == int32(0) {
		goto L54
	} else {
		goto L111
	}
L111:
	;
	v418 = v401
	goto L96
L112:
	;
	goto L95
L113:
	;
	F_PageRestoreTempPage(m, v323, v238)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L118
	}
L114:
	;
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+6)))
	if v449 != v428&int32(_a_F_btree_xlog_split_4) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v456 = F_PageAddItemExtended(m, v323, v311, v310, v430&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v456 == int32(0) {
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
	v464 = int32(128)
	goto L121
L120:
	;
	v464 = int32(129)
	goto L121
L121:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v334)+12)) = uint16(v464)
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v467 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v334)+14)) = uint16(v467)
	*(*int32)(unsafe.Add(mBase, uint32(v334)+4)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v238)+4)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v211
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	F_MarkBufferDirty(m, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
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
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L136
	}
L124:
	;
	v502 = F_XLogReadBufferForRedo(m, l1, int32(2), v24+int32(8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v502 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v506 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v535 == int32(0) {
		goto L123
	} else {
		goto L134
	}
L129:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v524)+16)))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v525+v524))) = v527
	*(*int32)(unsafe.Add(mBase, uint32(v524)+4)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v524))) = v211
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	F_MarkBufferDirty(m, v531)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510+(v506^int32(-1))<<(uint(int32(2))%32))))
	v524 = v516
	goto L129
L131:
	;
	goto L132
L132:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v524 = v518 + v506<<(uint(int32(13))%32) + int32(-8192)
	goto L129
L133:
	;
	goto L128
L134:
	;
	F_UnlockReleaseBuffer(m, v535)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L123
L136:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	if v543 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_UnlockReleaseBuffer(m, v543)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_9), int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(373), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_10), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(400), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_11), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(409), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_10), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(418), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v15 = l2
	v21 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L15
	}
L2:
	;
	v23 = int32(1)
	v25 = l1 + v21<<(uint(v23)%32)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v26<<(uint(int32(2))%32))))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	v36 = F_palloc(m, v31<<(uint(v23)%32)+int32(8))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = l0 + v30&int32(_a_F_btree_xlog_updates_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)) = uint16(v38)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)) = uint16(v44)
	v47 = v15 + int32(2)
	v49 = v44 << (uint(int32(1)) % 32)
	if v49 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	base.MemoryCopy(m, v36+int32(8), v47, v49)
	goto L8
L7:
	;
	goto L8
L8:
	;
	F__bt_update_posting(m, v36)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+6)))
	v64 = F_PageIndexTupleOverwrite(m, l0, v55, v56, (v57&int32(_a_F_btree_xlog_updates_1)+int32(7))&int32(_a_F_btree_xlog_updates_2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v64 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_pfree(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_pfree(m, v36)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15))))
	v74 = int32(1)
	v78 = v21 + v74
	if v78 != l3 {
		v15 = v47 + v73<<(uint(v74)%32)
		v21 = v78
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L3
L15:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_updates_3), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_updates_4), int32(585), int32(_a_F_btree_xlog_updates_5))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
