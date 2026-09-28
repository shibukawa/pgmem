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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int64
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int64
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+64))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	F_XLogRecGetBlockTag(m, l1, v3, v3, v3, v22+int32(20))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = int32(0)
	F_XLogRecGetBlockTag(m, l1, int32(1), v36, v36, v22+int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v43 = int32(0)
	v46 = v22 + int32(12)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+72))
	if v50 < int32(2) {
		v74 = v43
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v74 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	goto L4
L6:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49+int32(104))+76)))
	if v55 != int32(1) {
		v74 = v43
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
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v49+int32(180))+20))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v66
	goto L16
L15:
	;
	goto L16
L16:
	;
	v74 = int32(1)
	goto L5
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
	goto L20
L19:
	;
	goto L20
L20:
	;
	if v27 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v131 = F_XLogInitBufferForRedo(m, l1, int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L34
	}
L22:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	v85 = F_XLogReadBufferForRedo(m, l1, int32(3), v22+int32(28))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v85 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v89 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v122 == int32(0) {
		goto L21
	} else {
		goto L32
	}
L27:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v107)+16)))
	v109 = v108 + v107
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+12)))
	v112 = v110 & int32(_a_F_btree_xlog_split_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+12)) = uint16(v112)
	*(*int64)(unsafe.Add(mBase, uint32(v107))) = base.I64_rotl(v81, int64(32))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_MarkBufferDirty(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v89^int32(-1))<<(uint(int32(2))%32))))
	v107 = v99
	goto L27
L29:
	;
	goto L30
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v107 = v101 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L27
L31:
	;
	goto L26
L32:
	;
	F_UnlockReleaseBuffer(m, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L21
L34:
	;
	v135 = v22 + int32(24)
	v136 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+72))
	if v138 < int32(1) {
		v160 = v136
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v131 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L36:
	;
	v163 = v160
	goto L35
L37:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137+int32(52))+76)))
	if v143 != int32(1) {
		v160 = v136
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v147 = v137 + int32(128)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+43)))
	if v148 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v135 == int32(0) {
		v160 = v136
		goto L36
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v135 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v153 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v153
	v163 = v153
	goto L35
L43:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v156
	goto L45
L44:
	;
	goto L45
L45:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v160 = v158
	goto L36
L46:
	;
	F_PageInit(m, v181, int32(_a_F_btree_xlog_split_1), int32(16))
	mBase = m.M
	goto L50
L47:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+(v131^int32(-1))<<(uint(int32(2))%32))))
	v181 = v173
	goto L46
L48:
	;
	goto L49
L49:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v181 = v175 + v131<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v186 = v181 + v185
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v192 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+14)) = uint16(v192)
	*(*uint16)(unsafe.Add(mBase, uint32(v186)+12)) = uint16(base.B2i32(v27 == v192))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+8)) = v191
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	F__bt_restore_page(m, v181, v163, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v202 = base.I64_rotl(v24, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v181))) = v202
	F_MarkBufferDirty(m, v131)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v209 = F_XLogReadBufferForRedo(m, l1, int32(0), v22+int32(28))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L57
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L150
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L147
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L144
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L141
	}
L57:
	;
	if v209 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v213 < int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L60
L60:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v482 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L61:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+16)))
	v233 = int32(0)
	v235 = v22 + int32(24)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v237)+72))
	if v238 < v233 {
		v260 = v233
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217+(v213^int32(-1))<<(uint(int32(2))%32))))
	v231 = v223
	goto L61
L63:
	;
	goto L64
L64:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v231 = v225 + v213<<(uint(int32(13))%32) + int32(-8192)
	goto L61
L65:
	;
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+8)))
	if v264|l0 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L66:
	;
	v263 = v260
	goto L65
L67:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237+int32(0))+76)))
	if v243 != int32(1) {
		v260 = v233
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v247 = v237 + int32(76)
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+43)))
	if v248 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v235 == int32(0) {
		v260 = v233
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v235 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v253 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v253
	v263 = v253
	goto L65
L73:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v235))) = v256
	goto L75
L74:
	;
	goto L75
L75:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v247)+44))
	v260 = v258
	goto L66
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
	v313 = (v307&int32(_a_F_btree_xlog_split_2) + int32(7)) & int32(_a_F_btree_xlog_split_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v306 - v313
	v316 = F_PageGetTempPageCopySpecial(m, v231)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L85
	}
L77:
	;
	v300 = v263
	v302 = v3
	v303 = v3
	v304 = v3
	v305 = int32(0)
	goto L76
L78:
	;
	goto L79
L79:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+6)))
	v276 = (v270&int32(_a_F_btree_xlog_split_2) + int32(7)) & int32(_a_F_btree_xlog_split_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v269 - v276
	v279 = v263 + v276
	if l0 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	v284 = F_CopyIndexTuple(m, v263)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	if v264 != 0 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v300 = v279
	v302 = v3
	v303 = v276
	v304 = v263
	v305 = int32(0)
	goto L76
L83:
	;
	v289 = (v283 - int32(1)) & int32(_a_F_btree_xlog_split_4)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v231+v289<<(uint(int32(2))%32))+20))
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+8)))
	v298 = F__bt_swap_posting(m, v284, v231+v293&int32(_a_F_btree_xlog_split_5), v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v300 = v279
	v302 = v289
	v303 = v276
	v304 = v284
	v305 = v298
	goto L76
L85:
	;
	v320 = F_PageAddItemExtended(m, v316, v300, v313, int32(1), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v320 == int32(0) {
		goto L56
	} else {
		goto L87
	}
L87:
	;
	v324 = int32(2)
	v327 = v231 + v232
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	if v328 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v329 = v324
	goto L90
L89:
	;
	v329 = int32(1)
	goto L90
L90:
	;
	v330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if base.Ui32(v329) < base.Ui32(v330) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v336 = v329
	v338 = v324
	goto L94
L92:
	;
	v419 = v329
	v421 = v324
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
	if v336 == v302 {
		goto L97
	} else {
		goto L98
	}
L95:
	;
	v419 = v414
	v421 = v412
	goto L93
L96:
	;
	v411 = int32(1)
	v412 = v409 + v411
	v414 = v336 + v411
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if base.Ui32(v414) < base.Ui32(v415) {
		v336 = v414
		v338 = v412
		goto L94
	} else {
		goto L112
	}
L97:
	;
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v305)+6)))
	v364 = F_PageAddItemExtended(m, v316, v305, (v354&int32(_a_F_btree_xlog_split_2)+int32(7))&int32(_a_F_btree_xlog_split_3), v338&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
		v392 = v338
		goto L105
	} else {
		goto L106
	}
L100:
	;
	if v364 != 0 {
		v409 = v338
		goto L96
	} else {
		goto L101
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_6), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(383), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
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
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v231+int32(20)+v336<<(uint(int32(2))%32))))
	v405 = F_PageAddItemExtended(m, v316, v231+v396&int32(_a_F_btree_xlog_split_5), int32(base.Ui32(v396)>>(uint(int32(17))%32)), v392&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L110
	}
L106:
	;
	v381 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	if v336 != v381 {
		v392 = v338
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v386 = F_PageAddItemExtended(m, v316, v304, v303, v338&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v386 == int32(0) {
		goto L55
	} else {
		goto L109
	}
L109:
	;
	v392 = v338 + int32(1)
	goto L105
L110:
	;
	if v405 == int32(0) {
		goto L54
	} else {
		goto L111
	}
L111:
	;
	v409 = v392
	goto L96
L112:
	;
	goto L95
L113:
	;
	F_PageRestoreTempPage(m, v316, v231)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L1
	} else {
		goto L118
	}
L114:
	;
	v438 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+6)))
	if v438 != v419&int32(_a_F_btree_xlog_split_4) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v445 = F_PageAddItemExtended(m, v316, v304, v303, v421&int32(_a_F_btree_xlog_split_4), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v445 == int32(0) {
		goto L53
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	if v27 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v453 = int32(128)
	goto L121
L120:
	;
	v453 = int32(129)
	goto L121
L121:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v327)+12)) = uint16(v453)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v456 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v327)+14)) = uint16(v456)
	*(*int32)(unsafe.Add(mBase, uint32(v327)+4)) = v455
	*(*int64)(unsafe.Add(mBase, uint32(v231))) = v202
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_MarkBufferDirty(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L60
L123:
	;
	F_UnlockReleaseBuffer(m, v131)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L136
	}
L124:
	;
	v488 = F_XLogReadBufferForRedo(m, l1, int32(2), v22+int32(8))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v488 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v492 < int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	goto L128
L128:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v520 == int32(0) {
		goto L123
	} else {
		goto L134
	}
L129:
	;
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v510)+16)))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v511+v510))) = v513
	*(*int64)(unsafe.Add(mBase, uint32(v510))) = v202
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_MarkBufferDirty(m, v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L133
	}
L130:
	;
	v496 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[0]))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v496+(v492^int32(-1))<<(uint(int32(2))%32))))
	v510 = v502
	goto L129
L131:
	;
	goto L132
L132:
	;
	v504 = *(*int32)(unsafe.Add(mBase, _c_F_btree_xlog_split[1]))
	v510 = v504 + v492<<(uint(int32(13))%32) + int32(-8192)
	goto L129
L133:
	;
	goto L128
L134:
	;
	F_UnlockReleaseBuffer(m, v520)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L123
L136:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v528 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	F_UnlockReleaseBuffer(m, v528)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	m.G0 = v22 + int32(32)
	return
L140:
	;
	goto L139
L141:
	;
	F_errmsg_internal(m, int32(_a_F_btree_xlog_split_9), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(368), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
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
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(392), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
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
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(400), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
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
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_btree_xlog_split_7), int32(408), int32(_a_F_btree_xlog_split_8))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
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
	F_errstart_cold(m, int32(24), int32(0))
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
	F_errfinish(m, int32(_a_F_btree_xlog_updates_4), int32(573), int32(_a_F_btree_xlog_updates_5))
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
