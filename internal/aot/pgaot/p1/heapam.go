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
	var v26 int32
	_ = v26
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
		*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v5
		if v5 != 0 {
			F_LockBufferInternal(m, v5, int32(1))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				if v13 < int32(0) {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_block[0]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v13^int32(-1))*int32(56))+16))
					v32 = v23
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_block[1]))
					v26 = int32(56)
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+v13*v26-v26)+16))
					v32 = v31
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v32
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
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
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int64
	_ = v438
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v473 int32
	_ = v473
	v22 = m.G0
	v24 = v22 - int32(624)
	m.G0 = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v27) <= base.Ui32(v26) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v24 + int32(624)
	return v473
L2:
	;
	v30 = l0 + int32(116)
	goto L5
L3:
	;
	v363 = v26
	goto L4
L4:
	;
	v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v363<<(uint(int32(1))%32))+116)))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v385 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+108)) = int64(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v58 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v363 = v355
	goto L4
L7:
	;
	F_ReleaseBuffer(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v68 = F_read_stream_next_buffer(m, v65, v24+int32(620))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v68
	v71 = int32(0)
	if v68 == v71 {
		v473 = v71
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v24)+620))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	if v76 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v81 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v90 = v81
	v91 = v81
	goto L18
L15:
	;
	v133 = int32(-1)
	goto L16
L16:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v134)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_heap_page_prune_opt(m, v139, v140, l0+int32(104), int32(base.Ui32(v141&int32(1024))>>(uint(int32(10))%32)))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L10
	} else {
		goto L33
	}
L17:
	;
	v133 = v125
	goto L16
L18:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v85+int32(8)+v91<<(uint(int32(2))%32))))
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
	v104 = v90
	v106 = v91<<(uint(int32(5))%32) | int32(1)
	goto L23
L21:
	;
	v125 = v90
	goto L22
L22:
	;
	v130 = v91 + int32(1)
	if v130 != int32(10) {
		v90 = v125
		v91 = v130
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
	v125 = v118
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
	v118 = v104
	goto L27
L27:
	;
	v119 = int32(1)
	v122 = int32(base.Ui32(v102) >> (uint(v119) % 32))
	if v122 != 0 {
		v102 = v122
		v104 = v118
		v106 = v106 + v119
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
	v118 = v104 + int32(1)
	goto L27
L31:
	;
	goto L24
L32:
	;
	goto L19
L33:
	;
	F_LockBufferInternal(m, v140, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	if v151 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_UnlockBuffer(m, v140)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L10
	} else {
		goto L70
	}
L36:
	;
	if v133 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v211 = int32(0)
	if v140 < v211 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v332 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v158 = int32(base.Ui32(v136) >> (uint(int32(16)) % 32))
	v159 = int32(0)
	v166 = v159
	v168 = v159
	goto L42
L42:
	;
	v184 = int32(1)
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(32)+v166<<(uint(v184)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+30)) = uint16(v187)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+28)) = uint16(v136)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+26)) = uint16(v158)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v198 = F_heap_hot_search_buffer(m, v24+int32(26), v193, v140, v138, v24+int32(4), int32(0), v184)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v332 = v207
	goto L35
L44:
	;
	if v198 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v200 = int32(1)
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v168<<(uint(v200)%32)))) = uint16(v203)
	v207 = v168 + v200
	goto L47
L46:
	;
	v207 = v168
	goto L47
L47:
	;
	v209 = v166 + int32(1)
	if v209 != v133 {
		v166 = v209
		v168 = v207
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+12)))
	if base.Ui32(v230) < base.Ui32(int32(25)) {
		v332 = v211
		goto L35
	} else {
		goto L53
	}
L50:
	;
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[0]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215+(v140^int32(-1))<<(uint(int32(2))%32))))
	v229 = v221
	goto L49
L51:
	;
	goto L52
L52:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[1]))
	v229 = v223 + v140<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v234 = v230 + int32(_a_F_heapam_scan_bitmap_next_tuple_0)
	if v234&int32(_a_F_heapam_scan_bitmap_next_tuple_1) == int32(0) {
		v332 = v211
		goto L35
	} else {
		goto L54
	}
L54:
	;
	v240 = int32(base.Ui32(v136) >> (uint(int32(16)) % 32))
	v253 = int32(1)
	v255 = v211
	goto L55
L55:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v229+int32(20)+v253<<(uint(int32(2))%32))))
	if v272&int32(_a_F_heapam_scan_bitmap_next_tuple_2) == int32(_a_F_heapam_scan_bitmap_next_tuple_3) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v332 = v319
	goto L35
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = int32(base.Ui32(v272) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v229 + v272&int32(_a_F_heapam_scan_bitmap_next_tuple_4)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+12)) = uint16(v253)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+10)) = uint16(v136)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+8)) = uint16(v240)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v285
	v292 = F_HeapTupleSatisfiesVisibility(m, v24+int32(4), v138, v140)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L60
	}
L58:
	;
	v319 = v255
	goto L59
L59:
	;
	if v253 != int32(base.Ui32(v234)>>(uint(int32(2))%32))&int32(_a_F_heapam_scan_bitmap_next_tuple_5) {
		v253 = v253 + int32(1)
		v255 = v319
		goto L55
	} else {
		goto L69
	}
L60:
	;
	if v292 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v294 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v255<<(uint(v294)%32)))) = uint16(v253)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+20)))
	v303 = int32(768)
	if v302&v303 != v303 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v312 = v255
	goto L63
L63:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v292, v314, v24+int32(4), v140, v138)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L68
	}
L64:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v309 = v307
	goto L66
L65:
	;
	v309 = int32(2)
	goto L66
L66:
	;
	F_PredicateLockTID(m, v300, v24+int32(8), v138, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v312 = v255 + v294
	goto L63
L68:
	;
	v319 = v312
	goto L59
L69:
	;
	goto L56
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v332
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)))
	if v349 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v350 = l3
	goto L73
L72:
	;
	v350 = l4
	goto L73
L73:
	;
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v350)))
	*(*int64)(unsafe.Add(mBase, uint32(v350))) = v351 + int64(1)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui32(v356) <= base.Ui32(v355) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L6
L75:
	;
	v406 = v382<<(uint(int32(2))%32) + v403 + int32(20)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v407&int32(_a_F_heapam_scan_bitmap_next_tuple_4) + v403
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(base.Ui32(v412) >> (uint(int32(17)) % 32))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v416)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)) = uint16(v382)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+74)) = uint16(v419)
	v422 = int32(base.Ui32(v419) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v422)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v417
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v416)+272))
	if v427 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[0]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389+(v385^int32(-1))<<(uint(int32(2))%32))))
	v403 = v395
	goto L75
L77:
	;
	goto L78
L78:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[1]))
	v403 = v397 + v385<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	F_ExecStoreBufferHeapTuple(m, l0+int32(68), l1, v443)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L10
	} else {
		goto L85
	}
L80:
	;
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+268)))
	if v430 != int32(1) {
		v443 = v385
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v437 = v427
	goto L82
L82:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v437)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v437)+32)) = v438 + int64(1)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v443 = v442
	goto L79
L83:
	;
	F_pgstat_assoc_relation(m, v416)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+272))
	v437 = v436
	goto L82
L85:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v448 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v447 + v448
	v473 = v448
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
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v20
		F_heap_insert(m, l0, v18, l2, l3, l4)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v25)
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v27
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	v10 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	*(*uint8)(unsafe.Add(mBase, uint32(l8)+16)) = uint8(v10)
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+60)) = uint16(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v24
	v27 = l3 + int32(52)
	v29 = l7 & int32(1)
	v32 = F_heap_lock_tuple(m, l0, v27, l4, l5, l6, v29, v18+int32(108), l8)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = l7 & int32(2)
	if base.B2i32(v37 == int32(0))|base.B2i32(v32 != int32(3)) != 0 {
		v411 = v32
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v18 + int32(112)
	return v483
L4:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v471)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L138
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L134
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L130
	}
L7:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v417
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v417
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ExecStorePinnedBufferHeapTuple(m, v27, l3, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L129
	}
L8:
	;
	v44 = l3 + int32(56)
	goto L9
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v411 = v396
	goto L7
L11:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+2)))
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8))))
	v65 = int32(16)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	if v63|v64<<(uint(v65)%32) == v68|v69<<(uint(v65)%32) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v79 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	goto L12
L14:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+4)))
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	if v75 == v76 {
		v79 = int32(1)
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v79 = int32(0)
	goto L13
L17:
	;
	goto L16
L18:
	;
	v483 = int32(4)
	goto L3
L19:
	;
	goto L20
L20:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l8)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v81)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l8)+16)) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(4)
	v101 = v87
	goto L21
L21:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v105 == int32(_a_F_heapam_tuple_lock_0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v387)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L125
	}
L23:
	;
	v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	if v108&v109 == int32(_a_F_heapam_tuple_lock_1) {
		goto L6
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v115
	v122 = F_heap_fetch(m, l0, v18+int32(32), v27, v18+int32(108), int32(1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	if v122 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L22
L29:
	;
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
	v127 = int32(768)
	if v126&v127 != v127 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	if v124 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L32:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v133 = v131
	goto L34
L33:
	;
	v133 = int32(2)
	goto L34
L34:
	;
	if v133 != v101 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v470 = int32(4)
	goto L4
L36:
	;
	goto L37
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	if v136 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	if v137 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if base.Ui32(v101) < base.Ui32(int32(3)) {
		goto L56
	} else {
		goto L57
	}
L42:
	;
	switch l6 {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	default:
		goto L21
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[0])))
	v153 = F_ConditionalXactLockTableWait(m, v150, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L49
	}
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	v147 = F_ConditionalXactLockTableWait(m, v145, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
	F_XactLockTableWait(m, v141, l0, v44, int32(7))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L21
L47:
	;
	if v147 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v483 = int32(6)
	goto L3
L49:
	;
	if v153 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(50463045))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v162 + int32(4)
	F_errmsg(m, int32(_a_F_heapam_tuple_lock_2), v18)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(380), int32(_a_F_heapam_tuple_lock_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	if v305 == int32(0) {
		goto L28
	} else {
		goto L95
	}
L56:
	;
	v305 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[1]))
	if v185 == v101 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v305 = int32(1)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[2]))
	if v189 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v305 = v295
	goto L55
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[3]))
	if v193 == int32(0) {
		v295 = int32(0)
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[4]))
	v265 = int32(0)
	v268 = v189 - int32(1)
	goto L85
L66:
	;
	v198 = v193
	goto L67
L67:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v198)+20))
	if v204 == int32(4) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v295 = int32(0)
	goto L62
L69:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v198)+80))
	if v258 != 0 {
		v198 = v258
		goto L67
	} else {
		goto L84
	}
L70:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if v207 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v210 = int32(1)
	if v101 == v207 {
		v295 = v210
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v198)+52))
	v214 = v212 - int32(1)
	if v214 < int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v198)+48))
	v220 = int32(0)
	v223 = v214
	goto L74
L74:
	;
	v228 = int32(2)
	v229 = base.I32_div_s(v223-v220, v228)
	v230 = v229 + v220
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v217+v230<<(uint(v228)%32))))
	if v234 == v101 {
		v295 = v210
		goto L62
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v243 = base.B2i32(v234-v101 < int32(0)) | base.B2i32(base.Ui32(v234) < base.Ui32(int32(3)))
	if v243 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v244 = v230 + int32(1)
	goto L79
L78:
	;
	v244 = v220
	goto L79
L79:
	;
	if v243 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v247 = v223
	goto L82
L81:
	;
	v247 = v230 - int32(1)
	goto L82
L82:
	;
	if v244 <= v247 {
		v220 = v244
		v223 = v247
		goto L74
	} else {
		goto L83
	}
L83:
	;
	goto L75
L84:
	;
	goto L68
L85:
	;
	v273 = int32(2)
	v274 = base.I32_div_s(v268-v265, v273)
	v275 = v274 + v265
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v263+v275<<(uint(v273)%32))))
	v280 = base.B2i32(v279 == v101)
	if v279 == v101 {
		v295 = v280
		goto L62
	} else {
		goto L87
	}
L86:
	;
	v295 = v280
	goto L62
L87:
	;
	v283 = base.B2i32(base.Ui32(v279) < base.Ui32(v101))
	if base.Ui32(v279) < base.Ui32(v101) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v284 = v275 + int32(1)
	goto L90
L89:
	;
	v284 = v265
	goto L90
L90:
	;
	if base.Ui32(v279) < base.Ui32(v101) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v287 = v268
	goto L93
L92:
	;
	v287 = v275 - int32(1)
	goto L93
L93:
	;
	if v284 <= v287 {
		v265 = v284
		v268 = v287
		goto L85
	} else {
		goto L94
	}
L94:
	;
	goto L86
L95:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v308)+8))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+20)))
	if v311&int32(32) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if base.Ui32(v320) < base.Ui32(l4) {
		goto L28
	} else {
		goto L100
	}
L97:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[5]))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v310<<(uint(int32(3))%32))))
	v320 = v319
	goto L99
L98:
	;
	v320 = v310
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+8)) = v101
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+20)))
	if v326&int32(32) != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+12)) = v335
	v470 = int32(2)
	goto L4
L102:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[5]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v330+v325<<(uint(int32(3))%32))))
	v335 = v334
	goto L104
L103:
	;
	v335 = v325
	goto L104
L104:
	;
	goto L101
L105:
	;
	v483 = int32(4)
	goto L3
L106:
	;
	goto L107
L107:
	;
	v340 = int32(4)
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
	v342 = int32(768)
	if v341&v342 != v342 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v348 = v346
	goto L110
L109:
	;
	v348 = int32(2)
	goto L110
L110:
	;
	if v348 != v101 {
		v470 = v340
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v351 = v124 + int32(12)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	v354 = int32(16)
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351)+2)))
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	if v352|v353<<(uint(v354)%32) == v357|v358<<(uint(v354)%32) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v368 != 0 {
		v470 = v340
		goto L4
	} else {
		goto L118
	}
L113:
	;
	goto L112
L114:
	;
	v364 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351)+4)))
	if v364 == v365 {
		v368 = int32(1)
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v368 = int32(0)
	goto L113
L117:
	;
	goto L116
L118:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v370)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v372
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v375 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v374)+20)))
	if v375&int32(_a_F_heapam_tuple_lock_5) == int32(_a_F_heapam_tuple_lock_6) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L124
	}
L120:
	;
	v380 = F_HeapTupleGetUpdateXid(m, v374)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v374)+4))
	v383 = v382
	goto L119
L123:
	;
	v383 = v380
	goto L119
L124:
	;
	v101 = v383
	goto L21
L125:
	;
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v390)
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v392
	v396 = F_heap_lock_tuple(m, l0, v27, l4, l5, l6, v29, v18+int32(108), l8)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v37 == int32(0) {
		v411 = v396
		goto L7
	} else {
		goto L127
	}
L127:
	;
	if v396 == int32(3) {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	goto L10
L129:
	;
	v483 = v411
	goto L3
L130:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_heapam_tuple_lock_7), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(324), int32(_a_F_heapam_tuple_lock_4))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v446 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+58)))
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v449 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+60)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v448 + int32(4)
	v456 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v446 | v447<<(uint(v456)%32)
	F_errmsg_internal(m, int32(_a_F_heapam_tuple_lock_8), v18+v456)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(354), int32(_a_F_heapam_tuple_lock_4))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v483 = v470
	goto L3
}
