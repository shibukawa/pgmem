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
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_block[0]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v17+(v13^int32(-1))<<(uint(int32(6))%32))+16))
					v32 = v23
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_analyze_next_block[1]))
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v460 int32
	_ = v460
	v21 = m.G0
	v23 = v21 - int32(624)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v26) <= base.Ui32(v25) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v23 + int32(624)
	return v460
L2:
	;
	v29 = l0 + int32(108)
	goto L5
L3:
	;
	v352 = v25
	goto L4
L4:
	;
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v352<<(uint(int32(1))%32))+108)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v373 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+100)) = int64(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v54 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v352 = v344
	goto L4
L7:
	;
	F_ReleaseBuffer(m, v54)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v64 = F_read_stream_next_buffer(m, v61, v23+int32(620))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v64
	v67 = int32(0)
	if v64 == v67 {
		v460 = v67
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v23)+620))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v72 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v77 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v86 = v77
	v87 = v77
	goto L18
L15:
	;
	v129 = int32(-1)
	goto L16
L16:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v130)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_heap_page_prune_opt(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L33
	}
L17:
	;
	v129 = v121
	goto L16
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v81+int32(8)+v87<<(uint(int32(2))%32))))
	if v93 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v98 = v93
	v100 = v86
	v102 = v87<<(uint(int32(5))%32) | int32(1)
	goto L23
L21:
	;
	v121 = v86
	goto L22
L22:
	;
	v126 = v87 + int32(1)
	if v126 != int32(10) {
		v86 = v121
		v87 = v126
		goto L18
	} else {
		goto L32
	}
L23:
	;
	if v98&int32(1) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v121 = v114
	goto L22
L25:
	;
	if base.Ui32(v100) < base.Ui32(int32(291)) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v114 = v100
	goto L27
L27:
	;
	v115 = int32(1)
	v118 = int32(base.Ui32(v98) >> (uint(v115) % 32))
	if v118 != 0 {
		v98 = v118
		v100 = v114
		v102 = v102 + v115
		goto L23
	} else {
		goto L31
	}
L28:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(32)+v100<<(uint(int32(1))%32)))) = uint16(v102)
	goto L30
L29:
	;
	goto L30
L30:
	;
	v114 = v100 + int32(1)
	goto L27
L31:
	;
	goto L24
L32:
	;
	goto L19
L33:
	;
	F_LockBuffer(m, v136, int32(1))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v142 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	F_LockBuffer(m, v136, int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L70
	}
L36:
	;
	if v129 <= int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v201 = int32(0)
	if v136 < v201 {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v321 = int32(0)
	goto L35
L40:
	;
	goto L41
L41:
	;
	v149 = int32(base.Ui32(v132) >> (uint(int32(16)) % 32))
	v150 = int32(0)
	v157 = v150
	v159 = v150
	goto L42
L42:
	;
	v174 = int32(1)
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(32)+v157<<(uint(v174)%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)) = uint16(v177)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+28)) = uint16(v132)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+26)) = uint16(v149)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v188 = F_heap_hot_search_buffer(m, v23+int32(26), v183, v136, v134, v23+int32(4), int32(0), v174)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	v321 = v197
	goto L35
L44:
	;
	if v188 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v190 = int32(1)
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+30)))
	*(*uint16)(unsafe.Add(mBase, uint32(v29+v159<<(uint(v190)%32)))) = uint16(v193)
	v197 = v159 + v190
	goto L47
L46:
	;
	v197 = v159
	goto L47
L47:
	;
	v199 = v157 + int32(1)
	if v199 != v129 {
		v157 = v199
		v159 = v197
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+12)))
	if base.Ui32(v220) < base.Ui32(int32(25)) {
		v321 = v201
		goto L35
	} else {
		goto L53
	}
L50:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[0]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v205+(v136^int32(-1))<<(uint(int32(2))%32))))
	v219 = v211
	goto L49
L51:
	;
	goto L52
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[1]))
	v219 = v213 + v136<<(uint(int32(13))%32) + int32(-8192)
	goto L49
L53:
	;
	v224 = v220 + int32(_a_F_heapam_scan_bitmap_next_tuple_0)
	if v224&int32(_a_F_heapam_scan_bitmap_next_tuple_1) == int32(0) {
		v321 = v201
		goto L35
	} else {
		goto L54
	}
L54:
	;
	v230 = int32(base.Ui32(v132) >> (uint(int32(16)) % 32))
	v243 = int32(1)
	v245 = v201
	goto L55
L55:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v219+int32(20)+v243<<(uint(int32(2))%32))))
	if v261&int32(_a_F_heapam_scan_bitmap_next_tuple_2) == int32(_a_F_heapam_scan_bitmap_next_tuple_3) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v321 = v308
	goto L35
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = int32(base.Ui32(v261) >> (uint(int32(17)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v219 + v261&int32(_a_F_heapam_scan_bitmap_next_tuple_4)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+12)) = uint16(v243)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+10)) = uint16(v132)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v274
	v281 = F_HeapTupleSatisfiesVisibility(m, v23+int32(4), v134, v136)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L60
	}
L58:
	;
	v308 = v245
	goto L59
L59:
	;
	if v243 != int32(base.Ui32(v224)>>(uint(int32(2))%32))&int32(_a_F_heapam_scan_bitmap_next_tuple_5) {
		v243 = v243 + int32(1)
		v245 = v308
		goto L55
	} else {
		goto L69
	}
L60:
	;
	if v281 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v283 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v29+v245<<(uint(v283)%32)))) = uint16(v243)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v23)+20))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v290)+20)))
	v292 = int32(768)
	if v291&v292 != v292 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v301 = v245
	goto L63
L63:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_HeapCheckForSerializableConflictOut(m, v281, v303, v23+int32(4), v136, v134)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L10
	} else {
		goto L68
	}
L64:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v298 = v296
	goto L66
L65:
	;
	v298 = int32(2)
	goto L66
L66:
	;
	F_PredicateLockTID(m, v289, v23+int32(8), v134, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v301 = v245 + v283
	goto L63
L68:
	;
	v308 = v301
	goto L59
L69:
	;
	goto L56
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v321
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v338 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v339 = l3
	goto L73
L72:
	;
	v339 = l4
	goto L73
L73:
	;
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v339)))
	*(*int64)(unsafe.Add(mBase, uint32(v339))) = v340 + int64(1)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if base.Ui32(v345) <= base.Ui32(v344) {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	goto L6
L75:
	;
	v394 = v370<<(uint(int32(2))%32) + v391 + int32(20)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v395&int32(_a_F_heapam_scan_bitmap_next_tuple_4) + v391
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(base.Ui32(v400) >> (uint(int32(17)) % 32))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+72)) = uint16(v370)
	v407 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+70)) = uint16(v407)
	v410 = int32(base.Ui32(v407) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+68)) = uint16(v410)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v405
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v404)+272))
	if v415 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[0]))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v377+(v373^int32(-1))<<(uint(int32(2))%32))))
	v391 = v383
	goto L75
L77:
	;
	goto L78
L78:
	;
	v385 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_scan_bitmap_next_tuple[1]))
	v391 = v385 + v373<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l1, v431)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L10
	} else {
		goto L85
	}
L80:
	;
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+268)))
	if v418 != int32(1) {
		v431 = v373
		goto L79
	} else {
		goto L83
	}
L81:
	;
	v425 = v415
	goto L82
L82:
	;
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v425)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v425)+32)) = v426 + int64(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v431 = v430
	goto L79
L83:
	;
	F_pgstat_assoc_relation(m, v404)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)+272))
	v425 = v424
	goto L82
L85:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v436 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v435 + v436
	v460 = v436
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
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	v10 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	*(*uint8)(unsafe.Add(mBase, uint32(l8)+16)) = uint8(v10)
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)) = uint16(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v24
	v27 = l3 + int32(48)
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
		v399 = v32
		goto L7
	} else {
		goto L8
	}
L3:
	;
	m.G0 = v18 + int32(112)
	return v471
L4:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L138
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L134
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L130
	}
L7:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v405
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ExecStorePinnedBufferHeapTuple(m, v27, l3, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L129
	}
L8:
	;
	v44 = l3 + int32(52)
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
	v399 = v384
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
	v471 = int32(4)
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
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v375)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
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
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
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
	v458 = int32(4)
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
	v471 = int32(6)
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
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(471), int32(_a_F_heapam_tuple_lock_4))
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
	if v293 == int32(0) {
		goto L28
	} else {
		goto L95
	}
L56:
	;
	v293 = int32(0)
	goto L55
L57:
	;
	goto L58
L58:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[1]))
	if v184 == v101 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v293 = int32(1)
	goto L55
L60:
	;
	goto L61
L61:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[2]))
	if v188 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v293 = v285
	goto L55
L63:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[3]))
	if v192 == int32(0) {
		v285 = int32(0)
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[4]))
	v256 = int32(0)
	v258 = v188 - int32(1)
	goto L85
L66:
	;
	v197 = v192
	goto L67
L67:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)+20))
	if v202 == int32(4) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v285 = int32(0)
	goto L62
L69:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v197)+80))
	if v249 != 0 {
		v197 = v249
		goto L67
	} else {
		goto L84
	}
L70:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	if v205 == int32(0) {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v208 = int32(1)
	if v101 == v205 {
		v285 = v208
		goto L62
	} else {
		goto L72
	}
L72:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v197)+52))
	v212 = v210 - int32(1)
	if v212 < int32(0) {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v217 = int32(0)
	v219 = v212
	goto L74
L74:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v197)+48))
	v225 = int32(2)
	v226 = base.I32_div_s(v219-v217, v225)
	v227 = v226 + v217
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223+v227<<(uint(v225)%32))))
	if v231 == v101 {
		v285 = v208
		goto L62
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v235 = F_TransactionIdPrecedes(m, v231, v101)
	mBase = m.M
	if v235 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v236 = v227 + int32(1)
	goto L79
L78:
	;
	v236 = v217
	goto L79
L79:
	;
	if v235 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v239 = v219
	goto L82
L81:
	;
	v239 = v227 - int32(1)
	goto L82
L82:
	;
	if v236 <= v239 {
		v217 = v236
		v219 = v239
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
	v263 = int32(2)
	v264 = base.I32_div_s(v258-v256, v263)
	v265 = v264 + v256
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v254+v265<<(uint(v263)%32))))
	v270 = base.B2i32(v269 == v101)
	if v269 == v101 {
		v285 = v270
		goto L62
	} else {
		goto L87
	}
L86:
	;
	v285 = v270
	goto L62
L87:
	;
	v273 = base.B2i32(base.Ui32(v269) < base.Ui32(v101))
	if base.Ui32(v269) < base.Ui32(v101) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v274 = v265 + int32(1)
	goto L90
L89:
	;
	v274 = v256
	goto L90
L90:
	;
	if base.Ui32(v269) < base.Ui32(v101) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v277 = v258
	goto L93
L92:
	;
	v277 = v265 - int32(1)
	goto L93
L93:
	;
	if v274 <= v277 {
		v256 = v274
		v258 = v277
		goto L85
	} else {
		goto L94
	}
L94:
	;
	goto L86
L95:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v296)+8))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+20)))
	if v299&int32(32) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if base.Ui32(v308) < base.Ui32(l4) {
		goto L28
	} else {
		goto L100
	}
L97:
	;
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[5]))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v303+v298<<(uint(int32(3))%32))))
	v308 = v307
	goto L99
L98:
	;
	v308 = v298
	goto L99
L99:
	;
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+8)) = v101
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v311)+8))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+20)))
	if v314&int32(32) != 0 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8)+12)) = v323
	v458 = int32(2)
	goto L4
L102:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_tuple_lock[5]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318+v313<<(uint(int32(3))%32))))
	v323 = v322
	goto L104
L103:
	;
	v323 = v313
	goto L104
L104:
	;
	goto L101
L105:
	;
	v471 = int32(4)
	goto L3
L106:
	;
	goto L107
L107:
	;
	v328 = int32(4)
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+20)))
	v330 = int32(768)
	if v329&v330 != v330 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v336 = v334
	goto L110
L109:
	;
	v336 = int32(2)
	goto L110
L110:
	;
	if v336 != v101 {
		v458 = v328
		goto L4
	} else {
		goto L111
	}
L111:
	;
	v339 = v124 + int32(12)
	v340 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	v342 = int32(16)
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339)+2)))
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339))))
	if v340|v341<<(uint(v342)%32) == v345|v346<<(uint(v342)%32) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	if v356 != 0 {
		v458 = v328
		goto L4
	} else {
		goto L118
	}
L113:
	;
	goto L112
L114:
	;
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v339)+4)))
	if v352 == v353 {
		v356 = int32(1)
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v356 = int32(0)
	goto L113
L117:
	;
	goto L116
L118:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v358)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v357)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362)+20)))
	if v363&int32(_a_F_heapam_tuple_lock_5) == int32(_a_F_heapam_tuple_lock_6) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v18)+108))
	F_ReleaseBuffer(m, v372)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L124
	}
L120:
	;
	v368 = F_HeapTupleGetUpdateXid(m, v362)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L123
	}
L121:
	;
	goto L122
L122:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v362)+4))
	v371 = v370
	goto L119
L123:
	;
	v371 = v368
	goto L119
L124:
	;
	v101 = v371
	goto L21
L125:
	;
	v378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v378)
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v380
	v384 = F_heap_lock_tuple(m, l0, v27, l4, l5, l6, v29, v18+int32(108), l8)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v37 == int32(0) {
		v399 = v384
		goto L7
	} else {
		goto L127
	}
L127:
	;
	if v384 == int32(3) {
		goto L9
	} else {
		goto L128
	}
L128:
	;
	goto L10
L129:
	;
	v471 = v399
	goto L3
L130:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errmsg(m, int32(_a_F_heapam_tuple_lock_7), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(415), int32(_a_F_heapam_tuple_lock_4))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v434 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+54)))
	v435 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+52)))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v437 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+56)))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v18)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v436 + int32(4)
	v444 = int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v434 | v435<<(uint(v444)%32)
	F_errmsg_internal(m, int32(_a_F_heapam_tuple_lock_8), v18+v444)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_heapam_tuple_lock_3), int32(445), int32(_a_F_heapam_tuple_lock_4))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
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
	v471 = v458
	goto L3
}
