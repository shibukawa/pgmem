package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = v16 + int32(8)
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+3)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v26 = v24 - v19
	if v26 <= int32(0) {
		v149 = v21
		v152 = v23
		v153 = v22
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v166 = F_palloc0(m, int32(20))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v30 = v20 + int32(4)
	v31 = v21
	v33 = v23
	v34 = v23
	v35 = v22
	v36 = v19
	goto L3
L3:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v18+v36*int32(24))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+2)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+3)))
	if v31 < v54 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v139 == v141 {
		v149 = v137
		v152 = v141
		v153 = v138
		goto L1
	} else {
		goto L37
	}
L5:
	;
	v56 = v31
	goto L7
L6:
	;
	v56 = v54
	goto L7
L7:
	;
	if int32(0) < v56 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v60 = v49 + int32(4)
	v61 = int32(0)
	v66 = int32(8)
	v67 = base.I32_div_s(v56, v66)
	if v66 <= v56 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v137 = v56
	goto L10
L10:
	;
	if base.Ui32(v35) < base.Ui32(v50) {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	v137 = v133 + v129<<(uint(int32(3))%32)
	goto L10
L12:
	;
	goto L11
L13:
	;
	v115 = v106
	goto L24
L14:
	;
	v73 = v61
	goto L17
L15:
	;
	v90 = v61
	goto L16
L16:
	;
	v97 = v56 - v67<<(uint(int32(3))%32)
	if v97 == int32(0) {
		v129 = v90
		v133 = v61
		goto L12
	} else {
		goto L23
	}
L17:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v73))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v73))))
	if v79 != v81 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v90 = v67
	goto L16
L19:
	;
	v106 = int32(7)
	v107 = v73
	v109 = v79
	v110 = v81
	goto L13
L20:
	;
	goto L21
L21:
	;
	v85 = v73 + int32(1)
	if v85 != v67 {
		v73 = v85
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v90))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v90))))
	v106 = v97
	v107 = v90
	v109 = v103
	v110 = v101
	goto L13
L24:
	;
	if int32(base.Ui32(v109^v110)>>(uint(int32(8)-v115)%32)) != 0 {
		v115 = v115 - int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v129 = v107
	v133 = v115
	goto L12
L26:
	;
	goto L25
L27:
	;
	v138 = v35
	goto L29
L28:
	;
	v138 = v50
	goto L29
L29:
	;
	if base.Ui32(v52) < base.Ui32(v33) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v139 = v33
	goto L32
L31:
	;
	v139 = v52
	goto L32
L32:
	;
	if base.Ui32(v34) < base.Ui32(v52) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = v34
	goto L35
L34:
	;
	v141 = v52
	goto L35
L35:
	;
	v143 = v36 + int32(1)
	if v143 <= v26 {
		v31 = v137
		v33 = v139
		v34 = v141
		v35 = v138
		v36 = v143
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	v146 = int32(0)
	v149 = v146
	v152 = v146
	v153 = v146
	goto L1
L38:
	;
	return int64(0)
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+3)) = uint8(v149)
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)) = uint8(v153)
	*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)) = uint8(v152)
	if v149 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v191 = base.I32_div_s(v149, int32(8))
	v194 = v149 - v191<<(uint(int32(3))%32)
	if v194 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v178 = base.I32_div_s(v149+int32(7), int32(8))
	if v178 == int32(0) {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v181 = int32(4)
	base.MemoryCopy(m, v166+v181, base.I32_wrap_i64(v164)+v181, v178)
	goto L40
L43:
	;
	v197 = v191 + v166 + int32(4)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v201 = v198 & (int32(-256) >> (uint(v194) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v201)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
	v205 = v203
	goto L45
L44:
	;
	v205 = v152
	goto L45
L45:
	;
	if v205&int32(255) == int32(3) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v210 = int32(41)
	goto L48
L47:
	;
	v210 = int32(17)
	goto L48
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v210)
	return base.I64_extend_i32_u(v166)
}
func F_inet_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_network_in(m, v2, int32(0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v5)
	}
}
func F_inet_server_port(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_inet_server_port[0]))
	if v10 == int32(0) {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v43 = int64(0)
		m.G0 = v7 + int32(32)
		return v43
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
		switch v16 - int32(2) {
		case 0, 8:
			v22 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v22)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v10)+140))
			v31 = F_pg_getnameinfo_all(m, v10+int32(12), v26, v22, v22, v7, int32(32), int32(3))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				if v31 != 0 {
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
					v43 = int64(0)
					m.G0 = v7 + int32(32)
					return v43
				} else {
					v41 = F_DirectFunctionCall1Coll(m, int32(1142), int32(0), base.I64_extend_i32_u(v7))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						v43 = v41
						m.G0 = v7 + int32(32)
						return v43
					}
				}
			}
		default:
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v19)
			v43 = int64(0)
			m.G0 = v7 + int32(32)
			return v43
		}
	}
}
func F_inet_spg_consistent_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = int32(1)
	goto L3
L2:
	;
	v19 = int32(15)
	goto L3
L3:
	;
	if l1 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v19
L5:
	;
	goto L6
L6:
	;
	v23 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v25&v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v23
	goto L9
L8:
	;
	v28 = int32(4)
	goto L9
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v28)+1)))
	v32 = int32(base.Ui32(v30) >> (uint(int32(3)) % 32))
	v38 = int32(1) << (uint((v30^int32(-1))&int32(7)) % 32)
	v43 = v19
	v54 = int32(0)
	goto L10
L10:
	;
	v57 = l2 + v54*int32(56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+48))
	v59 = F_pg_detoast_datum_packed(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return v498
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+6)))
	v64 = int32(1)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v66&v64 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v494 = v54 + int32(1)
	if v494 != l1 {
		v43 = v488
		v54 = v494
		goto L10
	} else {
		goto L182
	}
L16:
	;
	v69 = v64
	goto L18
L17:
	;
	v69 = int32(4)
	goto L18
L18:
	;
	v70 = v59 + v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v72 = int32(1)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v74&v72 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v77 = v72
	goto L21
L20:
	;
	v77 = int32(4)
	goto L21
L21:
	;
	v78 = l0 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v71 != v79 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v81 = int32(0)
	switch v63 - int32(19) {
	case 0:
		goto L25
	case 1, 2:
		goto L27
	case 3, 4:
		goto L26
	default:
		v498 = v81
		goto L12
	}
L23:
	;
	goto L24
L24:
	;
	v87 = v63 - int32(18)
	switch v87 {
	case 0:
		goto L34
	default:
		v123 = v43
		goto L32
	case 6:
		goto L33
	case 7:
		goto L37
	case 8:
		goto L36
	case 9:
		goto L35
	}
L25:
	;
	if v43 != 0 {
		v488 = v43
		goto L15
	} else {
		goto L31
	}
L26:
	;
	if base.Ui32(v79) < base.Ui32(v71) {
		v498 = v81
		goto L12
	} else {
		goto L30
	}
L27:
	;
	if base.Ui32(v71) < base.Ui32(v79) {
		v498 = v81
		goto L12
	} else {
		goto L28
	}
L28:
	;
	if v43 != 0 {
		v488 = v43
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v498 = v81
	goto L12
L30:
	;
	goto L25
L31:
	;
	v498 = v81
	goto L12
L32:
	;
	v125 = int32(0)
	if v123 == v125 {
		v498 = v125
		goto L12
	} else {
		goto L58
	}
L33:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if base.Ui32(v120) < base.Ui32(v30) {
		goto L55
	} else {
		goto L56
	}
L34:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if base.Ui32(v30) < base.Ui32(v109) {
		goto L49
	} else {
		goto L50
	}
L35:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v102 == v30 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v30 == v93-int32(1) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if base.Ui32(v30) < base.Ui32(v90) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v92 = v43 & int32(12)
	goto L40
L39:
	;
	v92 = v43
	goto L40
L40:
	;
	v123 = v92
	goto L32
L41:
	;
	v123 = v43 & int32(3)
	goto L32
L42:
	;
	goto L43
L43:
	;
	if base.Ui32(v30) < base.Ui32(v93) {
		v123 = v43
		goto L32
	} else {
		goto L44
	}
L44:
	;
	return int32(0)
L45:
	;
	v123 = v43 & int32(3)
	goto L32
L46:
	;
	goto L47
L47:
	;
	if base.Ui32(v30) <= base.Ui32(v102) {
		v123 = v43
		goto L32
	} else {
		goto L48
	}
L48:
	;
	return int32(0)
L49:
	;
	v123 = v43 & int32(12)
	goto L32
L50:
	;
	goto L51
L51:
	;
	if v30 == v109 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v123 = v43 & int32(3)
	goto L32
L53:
	;
	goto L54
L54:
	;
	return int32(0)
L55:
	;
	v122 = v43
	goto L57
L56:
	;
	v122 = v43 & int32(12)
	goto L57
L57:
	;
	v123 = v122
	goto L32
L58:
	;
	v128 = int32(2)
	v129 = v78 + v128
	v131 = v70 + v128
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if base.Ui32(v30) < base.Ui32(v132) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v134 = v30
	goto L61
L60:
	;
	v134 = v132
	goto L61
L61:
	;
	v139 = base.I32_div_s(v134, int32(8))
	v140 = F_memcmp(m, v129, v131, v139)
	mBase = m.M
	if v140 != 0 {
		v226 = v140
		goto L64
	} else {
		goto L65
	}
L62:
	;
	if v236 != 0 {
		goto L83
	} else {
		goto L84
	}
L63:
	;
	if v227 != 0 {
		goto L80
	} else {
		goto L81
	}
L64:
	;
	v236 = v226
	goto L62
L65:
	;
	v141 = int32(0)
	v144 = v134 - v139<<(uint(int32(3))%32)
	if v144 <= v141 {
		v226 = v141
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129+v139))))
	v149 = int32(128)
	v150 = v148 & v149
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131+v139))))
	if v150 != v152&v149 {
		v227 = v150
		goto L63
	} else {
		goto L67
	}
L67:
	;
	if v144 == int32(1) {
		v226 = v141
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v158 = int32(1)
	v160 = int32(128)
	v161 = v148 << (uint(v158) % 32) & v160
	if v161 != v152<<(uint(v158)%32)&v160 {
		v227 = v161
		goto L63
	} else {
		goto L69
	}
L69:
	;
	if v144 < int32(3) {
		v226 = v141
		goto L64
	} else {
		goto L70
	}
L70:
	;
	v169 = int32(2)
	v171 = int32(128)
	v172 = v148 << (uint(v169) % 32) & v171
	if v172 != v152<<(uint(v169)%32)&v171 {
		v227 = v172
		goto L63
	} else {
		goto L71
	}
L71:
	;
	if v144 == int32(3) {
		v226 = v141
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v180 = int32(3)
	v182 = int32(128)
	v183 = v148 << (uint(v180) % 32) & v182
	if v183 != v152<<(uint(v180)%32)&v182 {
		v227 = v183
		goto L63
	} else {
		goto L73
	}
L73:
	;
	if v144 < int32(5) {
		v226 = v141
		goto L64
	} else {
		goto L74
	}
L74:
	;
	v191 = int32(4)
	v193 = int32(128)
	v194 = v148 << (uint(v191) % 32) & v193
	if v194 != v152<<(uint(v191)%32)&v193 {
		v227 = v194
		goto L63
	} else {
		goto L75
	}
L75:
	;
	if v144 == int32(5) {
		v226 = v141
		goto L64
	} else {
		goto L76
	}
L76:
	;
	v202 = int32(5)
	v204 = int32(128)
	v205 = v148 << (uint(v202) % 32) & v204
	if v205 != v152<<(uint(v202)%32)&v204 {
		v227 = v205
		goto L63
	} else {
		goto L77
	}
L77:
	;
	if v144 < int32(7) {
		v226 = v141
		goto L64
	} else {
		goto L78
	}
L78:
	;
	v213 = int32(6)
	v215 = int32(128)
	v216 = v148 << (uint(v213) % 32) & v215
	if v216 != v152<<(uint(v213)%32)&v215 {
		v227 = v216
		goto L63
	} else {
		goto L79
	}
L79:
	;
	v226 = v141
	goto L64
L80:
	;
	v230 = int32(1)
	goto L82
L81:
	;
	v230 = int32(-1)
	goto L82
L82:
	;
	v236 = v230
	goto L62
L83:
	;
	switch v63 - int32(19) {
	case 0:
		v488 = v123
		goto L15
	case 1, 2:
		goto L87
	case 3, 4:
		goto L86
	default:
		v498 = v125
		goto L12
	}
L84:
	;
	goto L85
L85:
	;
	if v123&int32(12) == int32(0) {
		v272 = v123
		goto L92
	} else {
		goto L93
	}
L86:
	;
	if int32(0) <= v236 {
		v488 = v123
		goto L15
	} else {
		goto L89
	}
L87:
	;
	if v236 <= int32(0) {
		v488 = v123
		goto L15
	} else {
		goto L88
	}
L88:
	;
	v498 = v125
	goto L12
L89:
	;
	v498 = v125
	goto L12
L90:
	;
	v315 = int32(1)
	if v314&v315 != 0 {
		goto L122
	} else {
		goto L123
	}
L91:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	switch v63 - int32(20) {
	case 0, 1:
		goto L111
	case 2, 3:
		goto L110
	default:
		v312 = v281
		v314 = v284
		goto L90
	}
L92:
	;
	if base.Ui32((v63-int32(24))&int32(_a_F_inet_spg_consistent_bitmap_0)) < base.Ui32(int32(_a_F_inet_spg_consistent_bitmap_1)) {
		v488 = v272
		goto L15
	} else {
		goto L108
	}
L93:
	;
	v247 = int32(1)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v249&v247 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v252 = v247
	goto L96
L95:
	;
	v252 = int32(4)
	goto L96
L96:
	;
	v253 = v59 + v252
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
	if base.Ui32(v254) <= base.Ui32(v30) {
		v272 = v123
		goto L92
	} else {
		goto L97
	}
L97:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253+v32)+2)))
	v258 = v38 & v257
	switch v63 - int32(19) {
	case 0:
		v312 = v123
		v314 = v249
		goto L90
	case 1, 2:
		goto L103
	case 3, 4:
		goto L102
	default:
		goto L101
	}
L98:
	;
	if v269 == int32(0) {
		v498 = v125
		goto L12
	} else {
		goto L107
	}
L99:
	;
	v269 = v123 & int32(11)
	goto L98
L100:
	;
	v269 = v123 & int32(7)
	goto L98
L101:
	;
	if v258 != 0 {
		goto L99
	} else {
		goto L106
	}
L102:
	;
	if v258 == int32(0) {
		v281 = v123
		goto L91
	} else {
		goto L105
	}
L103:
	;
	if v258 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v281 = v123
	goto L91
L105:
	;
	goto L99
L106:
	;
	goto L100
L107:
	;
	v272 = v269
	goto L92
L108:
	;
	v281 = v272
	goto L91
L109:
	;
	if v309 == int32(0) {
		v498 = v125
		goto L12
	} else {
		goto L121
	}
L110:
	;
	v298 = int32(1)
	if v284&v298 != 0 {
		goto L117
	} else {
		goto L118
	}
L111:
	;
	v289 = int32(1)
	if v284&v289 != 0 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v293 = v289
	goto L114
L113:
	;
	v293 = int32(4)
	goto L114
L114:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v293)+1)))
	if v295 == v30 {
		v309 = v281 & int32(3)
		goto L109
	} else {
		goto L115
	}
L115:
	;
	if base.Ui32(v30) <= base.Ui32(v295) {
		v312 = v281
		v314 = v284
		goto L90
	} else {
		goto L116
	}
L116:
	;
	v498 = v125
	goto L12
L117:
	;
	v302 = v298
	goto L119
L118:
	;
	v302 = int32(4)
	goto L119
L119:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v302)+1)))
	if base.Ui32(v304) <= base.Ui32(v30) {
		v312 = v281
		v314 = v284
		goto L90
	} else {
		goto L120
	}
L120:
	;
	v309 = v281 & int32(12)
	goto L109
L121:
	;
	v312 = v309
	v314 = v284
	goto L90
L122:
	;
	v319 = v315
	goto L124
L123:
	;
	v319 = int32(4)
	goto L124
L124:
	;
	v320 = v59 + v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)))
	if v30 != v321 {
		v488 = v312
		goto L15
	} else {
		goto L125
	}
L125:
	;
	if l3|base.B2i32(v312&int32(3) == int32(0)) != 0 {
		v351 = v312
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if l3 == int32(0) {
		v488 = v351
		goto L15
	} else {
		goto L142
	}
L127:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	if v330 == int32(2) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v333 = int32(32)
	goto L130
L129:
	;
	v333 = int32(128)
	goto L130
L130:
	;
	if base.Ui32(v333) <= base.Ui32(v30) {
		v351 = v312
		goto L126
	} else {
		goto L131
	}
L131:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v32)+2)))
	v337 = v38 & v336
	switch v63 - int32(19) {
	case 0:
		v351 = v312
		goto L126
	case 1, 2:
		goto L137
	case 3, 4:
		goto L136
	default:
		goto L135
	}
L132:
	;
	if v348 == int32(0) {
		v498 = v125
		goto L12
	} else {
		goto L141
	}
L133:
	;
	v348 = v312 & int32(14)
	goto L132
L134:
	;
	v348 = v312 & int32(13)
	goto L132
L135:
	;
	if v337 != 0 {
		goto L133
	} else {
		goto L140
	}
L136:
	;
	if v337 == int32(0) {
		v351 = v312
		goto L126
	} else {
		goto L139
	}
L137:
	;
	if v337 == int32(0) {
		goto L134
	} else {
		goto L138
	}
L138:
	;
	v351 = v312
	goto L126
L139:
	;
	goto L133
L140:
	;
	goto L134
L141:
	;
	v351 = v348
	goto L126
L142:
	;
	v355 = int32(1)
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v357&v355 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v360 = v355
	goto L145
L144:
	;
	v360 = int32(4)
	goto L145
L145:
	;
	v361 = l0 + v360
	v362 = int32(2)
	v363 = v361 + v362
	v365 = v320 + v362
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361))))
	if v368 == v362 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v371 = int32(32)
	goto L148
L147:
	;
	v371 = int32(128)
	goto L148
L148:
	;
	v376 = base.I32_div_s(v371, int32(8))
	v377 = F_memcmp(m, v363, v365, v376)
	mBase = m.M
	if v377 != 0 {
		v463 = v377
		goto L151
	} else {
		goto L152
	}
L149:
	;
	switch v87 - int32(1) {
	case 0:
		goto L170
	case 1:
		goto L175
	case 2:
		goto L174
	case 3:
		goto L171
	case 4:
		goto L172
	default:
		goto L173
	}
L150:
	;
	if v464 != 0 {
		goto L167
	} else {
		goto L168
	}
L151:
	;
	v473 = v463
	goto L149
L152:
	;
	v378 = int32(0)
	v381 = v371 - v376<<(uint(int32(3))%32)
	if v381 <= v378 {
		v463 = v378
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363+v376))))
	v386 = int32(128)
	v387 = v385 & v386
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365+v376))))
	if v387 != v389&v386 {
		v464 = v387
		goto L150
	} else {
		goto L154
	}
L154:
	;
	if v381 == int32(1) {
		v463 = v378
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v395 = int32(1)
	v397 = int32(128)
	v398 = v385 << (uint(v395) % 32) & v397
	if v398 != v389<<(uint(v395)%32)&v397 {
		v464 = v398
		goto L150
	} else {
		goto L156
	}
L156:
	;
	if v381 < int32(3) {
		v463 = v378
		goto L151
	} else {
		goto L157
	}
L157:
	;
	v406 = int32(2)
	v408 = int32(128)
	v409 = v385 << (uint(v406) % 32) & v408
	if v409 != v389<<(uint(v406)%32)&v408 {
		v464 = v409
		goto L150
	} else {
		goto L158
	}
L158:
	;
	if v381 == int32(3) {
		v463 = v378
		goto L151
	} else {
		goto L159
	}
L159:
	;
	v417 = int32(3)
	v419 = int32(128)
	v420 = v385 << (uint(v417) % 32) & v419
	if v420 != v389<<(uint(v417)%32)&v419 {
		v464 = v420
		goto L150
	} else {
		goto L160
	}
L160:
	;
	if v381 < int32(5) {
		v463 = v378
		goto L151
	} else {
		goto L161
	}
L161:
	;
	v428 = int32(4)
	v430 = int32(128)
	v431 = v385 << (uint(v428) % 32) & v430
	if v431 != v389<<(uint(v428)%32)&v430 {
		v464 = v431
		goto L150
	} else {
		goto L162
	}
L162:
	;
	if v381 == int32(5) {
		v463 = v378
		goto L151
	} else {
		goto L163
	}
L163:
	;
	v439 = int32(5)
	v441 = int32(128)
	v442 = v385 << (uint(v439) % 32) & v441
	if v442 != v389<<(uint(v439)%32)&v441 {
		v464 = v442
		goto L150
	} else {
		goto L164
	}
L164:
	;
	if v381 < int32(7) {
		v463 = v378
		goto L151
	} else {
		goto L165
	}
L165:
	;
	v450 = int32(6)
	v452 = int32(128)
	v453 = v385 << (uint(v450) % 32) & v452
	if v453 != v389<<(uint(v450)%32)&v452 {
		v464 = v453
		goto L150
	} else {
		goto L166
	}
L166:
	;
	v463 = v378
	goto L151
L167:
	;
	v467 = int32(1)
	goto L169
L168:
	;
	v467 = int32(-1)
	goto L169
L169:
	;
	v473 = v467
	goto L149
L170:
	;
	if v473 == int32(0) {
		v498 = v125
		goto L12
	} else {
		goto L181
	}
L171:
	;
	if int32(0) < v473 {
		v488 = v351
		goto L15
	} else {
		goto L180
	}
L172:
	;
	if int32(0) <= v473 {
		v488 = v351
		goto L15
	} else {
		goto L179
	}
L173:
	;
	if v473 == int32(0) {
		v488 = v351
		goto L15
	} else {
		goto L178
	}
L174:
	;
	if v473 <= int32(0) {
		v488 = v351
		goto L15
	} else {
		goto L177
	}
L175:
	;
	if v473 < int32(0) {
		v488 = v351
		goto L15
	} else {
		goto L176
	}
L176:
	;
	v498 = v125
	goto L12
L177:
	;
	v498 = v125
	goto L12
L178:
	;
	v498 = v125
	goto L12
L179:
	;
	v498 = v125
	goto L12
L180:
	;
	v498 = v125
	goto L12
L181:
	;
	v488 = v351
	goto L15
L182:
	;
	v498 = v488
	goto L12
}
