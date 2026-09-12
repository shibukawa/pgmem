package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_scan_analyze_next_block(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v5 = F_read_stream_next_buffer(m, l1, int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v5
		if v5 != 0 {
			F_LockBuffer(m, v5, int32(1))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v13 < int32(0) {
					v17 = *(*int32)(unsafe.Add(mBase, _consts[8]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v13^int32(-1))<<(uint(int32(6))%32))+16))
					v32 = v23
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+v13<<(uint(int32(6))%32)+int32(-64))+16))
					v32 = v31
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v32
				return base.B2i32(v5 != int32(0))
			}
		} else {
			return base.B2i32(v5 != int32(0))
		}
	}
}
func F_heapam_scan_bitmap_next_tuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v56 int32
	_ = v56
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int64
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int64
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v477 int32
	_ = v477
	v22 = m.G0
	v24 = v22 - int32(624)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v27) <= base.Ui32(v26) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v24 + int32(624)
	return v477
L2:
	;
	v30 = l0 + int32(108)
	goto L5
L3:
	;
	v367 = v26
	goto L4
L4:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v367<<(uint(int32(1))%32))+108)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v387 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+100)) = int64(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v56 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v367 = v359
	goto L4
L7:
	;
	F_ReleaseBuffer(m, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v66 = F_read_stream_next_buffer(m, v63, v24+int32(620))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v66
	v69 = int32(0)
	if v66 == v69 {
		v477 = v69
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)+620))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
	if v74 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v79 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v89 = v79
	v90 = v79
	goto L18
L15:
	;
	v137 = int32(-1)
	goto L16
L16:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v138)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_heap_page_prune_opt(m, v143, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L33
	}
L17:
	;
	v137 = v128
	goto L16
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84+int32(8)+v90<<(uint(int32(2))%32))))
	if v97 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v102 = v97
	v104 = v89
	v106 = v90<<(uint(int32(5))%32) | int32(1)
	goto L23
L21:
	;
	v128 = v89
	goto L22
L22:
	;
	v134 = v90 + int32(1)
	if v134 != int32(10) {
		v89 = v128
		v90 = v134
		goto L18
	} else {
		goto L32
	}
L23:
	;
	if v102&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v128 = v119
	goto L22
L25:
	;
	if base.Ui32(v104) < base.Ui32(int32(291)) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v119 = v104
	goto L27
L27:
	;
	v120 = int32(1)
	if base.Ui32(v120) < base.Ui32(v102) {
		v102 = int32(base.Ui32(v102) >> (uint(v120) % 32))
		v104 = v119
		v106 = v106 + v120
		goto L23
	} else {
		goto L31
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(32)+v104<<(uint(int32(1))%32)))) = uint16(v106)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v119 = v104 + int32(1)
	goto L27
L31:
	;
	goto L24
L32:
	;
	goto L19
L33:
	;
	F_LockBuffer(m, v144, int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
	if v150 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_LockBuffer(m, v144, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L70
	}
L36:
	;
	if v137 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v210 = int32(0)
	if v144 < v210 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v335 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v157 = int32(base.Ui32(v140) >> (uint(int32(16)) % 32))
	v158 = int32(0)
	v165 = v158
	v167 = v158
	goto L42
L42:
	;
	v183 = int32(1)
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(32)+v165<<(uint(v183)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+30)) = uint16(v186)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v140)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+26)) = uint16(v157)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v197 = F_heap_hot_search_buffer(m, v24+int32(26), v192, v144, v142, v24+int32(4), int32(0), v183)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v335 = v206
	goto L35
L44:
	;
	if v197 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v199 = int32(1)
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v167<<(uint(v199)%32)))) = uint16(v202)
	v206 = v167 + v199
	goto L47
L46:
	;
	v206 = v167
	goto L47
L47:
	;
	v208 = v165 + int32(1)
	if v208 != v137 {
		v165 = v208
		v167 = v206
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+12)))
	if base.Ui32(v229) < base.Ui32(int32(25)) {
		v335 = v210
		goto L35
	} else {
		goto L53
	}
L50:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+(v144^int32(-1))<<(uint(int32(2))%32))))
	v228 = v220
	goto L49
L51:
	;
	goto L52
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v228 = v222 + v144<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v233 = v229 + int32(262120)
	if v233&int32(262140) == int32(0) {
		v335 = v210
		goto L35
	} else {
		goto L54
	}
L54:
	;
	v239 = int32(base.Ui32(v140) >> (uint(int32(16)) % 32))
	v252 = int32(1)
	v254 = v210
	goto L55
L55:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v252<<(uint(int32(2))%32)+(v228+int32(24))-int32(4))))
	if v273&int32(98304) == int32(32768) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v335 = v321
	goto L35
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(base.Ui32(v273) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v228 + v273&int32(32767)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)) = uint16(v252)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)) = uint16(v140)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)) = uint16(v239)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v286
	v293 = F_HeapTupleSatisfiesVisibility(m, v24+int32(4), v142, v144)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L60
	}
L58:
	;
	v321 = v254
	goto L59
L59:
	;
	if v252 != int32(base.Ui32(v233)>>(uint(int32(2))%32))&int32(65535) {
		v252 = v252 + int32(1)
		v254 = v321
		goto L55
	} else {
		goto L69
	}
L60:
	;
	if v293 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v295 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v254<<(uint(v295)%32)))) = uint16(v252)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303)+20)))
	v305 = int32(768)
	if v304&v305 != v305 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v313 = v254
	goto L63
L63:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v293, v316, v24+int32(4), v144, v142)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L10
	} else {
		goto L68
	}
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303)))
	v310 = v309
	goto L66
L65:
	;
	v310 = int32(2)
	goto L66
L66:
	;
	F_PredicateLockTID(m, v302, v24+int32(8), v142, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v313 = v254 + v295
	goto L63
L68:
	;
	v321 = v313
	goto L59
L69:
	;
	goto L56
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v335
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+4)))
	if v353 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v354 = l3
	goto L73
L72:
	;
	v354 = l4
	goto L73
L73:
	;
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v354)))
	*(*int64)(unsafe.Add(mBase, uint32(v354))) = v355 + int64(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v360) <= base.Ui32(v359) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L6
L75:
	;
	v410 = v405 + v386<<(uint(int32(2))%32) + int32(20)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v411&int32(32767) + v405
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v416) >> (uint(int32(17)) % 32))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v386)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v423)
	v426 = int32(base.Ui32(v423) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v426)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v421
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v420)+272))
	if v431 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v391+(v387^int32(-1))<<(uint(int32(2))%32))))
	v405 = v397
	goto L75
L77:
	;
	goto L78
L78:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v405 = v399 + v387<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l1, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L10
	} else {
		goto L85
	}
L80:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v420)+268)))
	if v434 != int32(1) {
		v447 = v387
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v441 = v431
	goto L82
L82:
	;
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v441)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v441)+32)) = v442 + int64(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v447 = v446
	goto L79
L83:
	;
	F_pgstat_assoc_relation(m, v420)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v439)+272))
	v441 = v440
	goto L82
L85:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v452 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v451 + v452
	v477 = v452
	goto L1
}
func F_heapam_tuple_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v13)
	v18 = F_ExecFetchSlotHeapTuple(m, l1, v13, v11+int32(15))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v20
		F_heap_insert(m, l0, v18, l2, l3, l4)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v27
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
			if v29 == int32(1) {
				F_pfree(m, v18)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v11 + int32(16)
					return
				}
			} else {
				m.G0 = v11 + int32(16)
				return
			}
		}
	}
}
func F_heapam_tuple_lock(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v407 int32
	_ = v407
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v482 int32
	_ = v482
	v10 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(112)
	m.G0 = v20
	*(*uint8)(unsafe.Add(mBase, uint32(l8)+16)) = uint8(v10)
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v26
	v29 = l3 + int32(48)
	v31 = l7 & int32(1)
	v34 = F_heap_lock_tuple(m, l0, v29, l4, l5, l6, v31, v20+int32(108), l8)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v39 = l7 & int32(2)
	if v39 == int32(0) {
		v407 = v34
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v20 + int32(112)
	return v482
L4:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ReleaseBuffer(m, v470)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L139
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L135
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L131
	}
L7:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v415
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v415
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ExecStorePinnedBufferHeapTuple(m, v29, l3, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L130
	}
L8:
	;
	if v34 != int32(3) {
		v407 = v34
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v45 = l3 + int32(52)
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ReleaseBuffer(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v407 = v392
	goto L7
L12:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+2)))
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8))))
	v68 = int32(16)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+2)))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
	if v66|v67<<(uint(v68)%32) == v71|v72<<(uint(v68)%32) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if v82 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+4)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	if v78 == v79 {
		v82 = int32(1)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v82 = int32(0)
	goto L14
L18:
	;
	goto L17
L19:
	;
	v482 = int32(4)
	goto L3
L20:
	;
	goto L21
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
	v86 = int32(4)
	v87 = l1 + v86
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v88)
	v90 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8)+16)) = uint8(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v86
	v106 = v92
	goto L22
L22:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v112 == int32(65533) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ReleaseBuffer(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L126
	}
L24:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v115&v116 == int32(65535) {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v120
	v123 = l3 + int32(56)
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v124)
	v131 = F_heap_fetch(m, l0, v20+int32(32), v29, v20+int32(108), int32(1))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	if v131 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L23
L30:
	;
	v134 = int32(2)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+20)))
	v136 = int32(768)
	if v135&v136 != v136 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	if v133 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v141 = v140
	goto L35
L34:
	;
	v141 = v134
	goto L35
L35:
	;
	if v141 != v106 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v468 = int32(4)
	goto L4
L37:
	;
	goto L38
L38:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	if v144 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	if v145 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ReleaseBuffer(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(v106) < base.Ui32(int32(3)) {
		goto L57
	} else {
		goto L58
	}
L43:
	;
	switch l6 {
	case 0:
		goto L46
	case 1:
		goto L45
	case 2:
		goto L44
	default:
		goto L22
	}
L44:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[75])))
	v161 = F_ConditionalXactLockTableWait(m, v158, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L50
	}
L45:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v155 = F_ConditionalXactLockTableWait(m, v153, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	F_XactLockTableWait(m, v149, l0, v45, int32(7))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	goto L22
L48:
	;
	if v155 != 0 {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v482 = int32(6)
	goto L3
L50:
	;
	if v161 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v170 + int32(4)
	F_errmsg(m, int32(685543), v20)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(490370), int32(471), int32(313938))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	if v301 == int32(0) {
		goto L29
	} else {
		goto L96
	}
L57:
	;
	v301 = int32(0)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v192 == v106 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v301 = int32(1)
	goto L56
L61:
	;
	goto L62
L62:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v196 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v301 = v293
	goto L56
L64:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v200 == int32(0) {
		v293 = int32(0)
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v264 = int32(0)
	v266 = v196 - int32(1)
	goto L86
L67:
	;
	v205 = v200
	goto L68
L68:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	if v210 == int32(4) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v293 = int32(0)
	goto L63
L70:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v205)+80))
	if v257 != 0 {
		v205 = v257
		goto L68
	} else {
		goto L85
	}
L71:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	if v213 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v216 = int32(1)
	if v106 == v213 {
		v293 = v216
		goto L63
	} else {
		goto L73
	}
L73:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v205)+52))
	v220 = v218 - int32(1)
	if v220 < int32(0) {
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v225 = int32(0)
	v227 = v220
	goto L75
L75:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v205)+48))
	v233 = int32(2)
	v234 = base.I32_div_s(v227-v225, v233)
	v235 = v234 + v225
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v231+v235<<(uint(v233)%32))))
	if v239 == v106 {
		v293 = v216
		goto L63
	} else {
		goto L77
	}
L76:
	;
	goto L70
L77:
	;
	v243 = F_TransactionIdPrecedes(m, v239, v106)
	mBase = m.M
	if v243 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v244 = v235 + int32(1)
	goto L80
L79:
	;
	v244 = v225
	goto L80
L80:
	;
	if v243 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v247 = v227
	goto L83
L82:
	;
	v247 = v235 - int32(1)
	goto L83
L83:
	;
	if v244 <= v247 {
		v225 = v244
		v227 = v247
		goto L75
	} else {
		goto L84
	}
L84:
	;
	goto L76
L85:
	;
	goto L69
L86:
	;
	v271 = int32(2)
	v272 = base.I32_div_s(v266-v264, v271)
	v273 = v272 + v264
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v262+v273<<(uint(v271)%32))))
	v278 = base.B2i32(v277 == v106)
	if v277 == v106 {
		v293 = v278
		goto L63
	} else {
		goto L88
	}
L87:
	;
	v293 = v278
	goto L63
L88:
	;
	v281 = base.B2i32(base.Ui32(v277) < base.Ui32(v106))
	if base.Ui32(v277) < base.Ui32(v106) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v282 = v273 + int32(1)
	goto L91
L90:
	;
	v282 = v264
	goto L91
L91:
	;
	if base.Ui32(v277) < base.Ui32(v106) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v285 = v266
	goto L94
L93:
	;
	v285 = v273 - int32(1)
	goto L94
L94:
	;
	if v282 <= v285 {
		v264 = v282
		v266 = v285
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L87
L96:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v304)+8))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304)+20)))
	if v307&int32(32) != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if base.Ui32(v316) < base.Ui32(l4) {
		goto L29
	} else {
		goto L101
	}
L98:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v311+v306<<(uint(int32(3))%32))))
	v316 = v315
	goto L100
L99:
	;
	v316 = v306
	goto L100
L100:
	;
	goto L97
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+8)) = v106
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v319)+8))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+20)))
	if v322&int32(32) != 0 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+12)) = v331
	v468 = v134
	goto L4
L103:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v326+v321<<(uint(int32(3))%32))))
	v331 = v330
	goto L105
L104:
	;
	v331 = v321
	goto L105
L105:
	;
	goto L102
L106:
	;
	v482 = int32(4)
	goto L3
L107:
	;
	goto L108
L108:
	;
	v337 = int32(4)
	v338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+20)))
	v339 = int32(768)
	if v338&v339 != v339 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v344 = v343
	goto L111
L110:
	;
	v344 = int32(2)
	goto L111
L111:
	;
	if v344 != v106 {
		v468 = v337
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v347 = v133 + int32(12)
	v348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+2)))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
	v350 = int32(16)
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+2)))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347))))
	if v348|v349<<(uint(v350)%32) == v353|v354<<(uint(v350)%32) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if v364 != 0 {
		v468 = v337
		goto L4
	} else {
		goto L119
	}
L114:
	;
	goto L113
L115:
	;
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347)+4)))
	if v360 == v361 {
		v364 = int32(1)
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v364 = int32(0)
	goto L114
L118:
	;
	goto L117
L119:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v366
	v368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v365)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v87))) = uint16(v368)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v370)+20)))
	if v371&int32(6272) == int32(4096) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	F_ReleaseBuffer(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L125
	}
L121:
	;
	v376 = F_HeapTupleGetUpdateXid(m, v370)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	v379 = v378
	goto L120
L124:
	;
	v379 = v376
	goto L120
L125:
	;
	v106 = v379
	goto L22
L126:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v386)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v388
	v392 = F_heap_lock_tuple(m, l0, v29, l4, l5, l6, v31, v20+int32(108), l8)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v39 == int32(0) {
		v407 = v392
		goto L7
	} else {
		goto L128
	}
L128:
	;
	if v392 == int32(3) {
		goto L10
	} else {
		goto L129
	}
L129:
	;
	goto L11
L130:
	;
	v482 = v407
	goto L3
L131:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errmsg(m, int32(352573), int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(490370), int32(415), int32(313938))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v444 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+54)))
	v445 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+52)))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v446 + int32(4)
	v454 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v444 | v445<<(uint(v454)%32)
	F_errmsg_internal(m, int32(698423), v20+v454)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(490370), int32(445), int32(313938))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v482 = v468
	goto L3
}
