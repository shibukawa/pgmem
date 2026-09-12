package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_inet_gist_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = v15 + int32(4)
	v18 = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+3)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v25 = v23 - v18
	if v25 <= int32(0) {
		v147 = v20
		v150 = v22
		v151 = v21
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v163 = F_palloc0(m, int32(20))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	v29 = v19 + int32(4)
	v30 = v20
	v33 = v22
	v34 = v21
	v35 = v22
	v36 = v18
	goto L3
L3:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v17+v36<<(uint(int32(4))%32))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+2)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+3)))
	if v30 < v52 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v139 == v137 {
		v147 = v135
		v150 = v139
		v151 = v136
		goto L1
	} else {
		goto L37
	}
L5:
	;
	v54 = v30
	goto L7
L6:
	;
	v54 = v52
	goto L7
L7:
	;
	if int32(0) < v54 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = v47 + int32(4)
	v59 = int32(0)
	v64 = int32(8)
	v65 = base.I32_div_s(v54, v64)
	if v64 <= v54 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v135 = v54
	goto L10
L10:
	;
	if base.Ui32(v34) < base.Ui32(v48) {
		goto L27
	} else {
		goto L28
	}
L11:
	;
	v135 = v129 + v127<<(uint(int32(3))%32)
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
	v71 = v59
	goto L17
L15:
	;
	v88 = v59
	goto L16
L16:
	;
	v95 = v54 - v65<<(uint(int32(3))%32)
	if v95 == int32(0) {
		v127 = v88
		v129 = v59
		goto L12
	} else {
		goto L23
	}
L17:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v71))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v71))))
	if v77 != v79 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = v65
	goto L16
L19:
	;
	v105 = v71
	v106 = int32(7)
	v108 = v77
	v109 = v79
	goto L13
L20:
	;
	goto L21
L21:
	;
	v83 = v71 + int32(1)
	if v83 != v65 {
		v71 = v83
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v88))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v88))))
	v105 = v88
	v106 = v95
	v108 = v101
	v109 = v99
	goto L13
L24:
	;
	if int32(base.Ui32(v108^v109)>>(uint(int32(8)-v115)%32)) != 0 {
		v115 = v115 - int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v127 = v105
	v129 = v115
	goto L12
L26:
	;
	goto L25
L27:
	;
	v136 = v34
	goto L29
L28:
	;
	v136 = v48
	goto L29
L29:
	;
	if base.Ui32(v50) < base.Ui32(v35) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v137 = v35
	goto L32
L31:
	;
	v137 = v50
	goto L32
L32:
	;
	if base.Ui32(v33) < base.Ui32(v50) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v139 = v33
	goto L35
L34:
	;
	v139 = v50
	goto L35
L35:
	;
	v141 = v36 + int32(1)
	if v141 <= v25 {
		v30 = v135
		v33 = v139
		v34 = v136
		v35 = v137
		v36 = v141
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L4
L37:
	;
	v144 = int32(0)
	v147 = v144
	v150 = v144
	v151 = v144
	goto L1
L38:
	;
	return int32(0)
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+3)) = uint8(v147)
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)) = uint8(v151)
	*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)) = uint8(v150)
	if int32(0) < v147 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v172 = int32(4)
	v179 = base.I32_div_s(v147+int32(7), int32(8))
	if v179 != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	goto L42
L42:
	;
	v185 = base.I32_div_s(v147, int32(8))
	v188 = v147 - v185<<(uint(int32(3))%32)
	if v188 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	goto L42
L44:
	;
	v180 = F__emscripten_memcpy_bulkmem(m, v163+v172, v161+v172, v179)
	mBase = m.M
	goto L46
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v191 = v163 + v185 + int32(4)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v195 = v192 & (int32(-256) >> (uint(v188) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v195)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v199 = v197
	goto L49
L48:
	;
	v199 = v150
	goto L49
L49:
	;
	if v199&int32(255) == int32(3) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v204 = int32(41)
	goto L52
L51:
	;
	v204 = int32(17)
	goto L52
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v204)
	return v163
}
func F_inet_in(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_network_in(m, v2, int32(0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_inet_server_port(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_inet_server_port[0]))
	if v10 == v2 {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v39 = v2
		m.G0 = v7 + int32(32)
		return v39
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)))
		switch v15 - int32(2) {
		case 0, 8:
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v7))) = uint8(v20)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+140))
			v29 = F_pg_getnameinfo_all(m, v10+int32(12), v24, v20, v20, v7, int32(32), int32(3))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				if v29 != 0 {
					v33 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
					v39 = v2
					m.G0 = v7 + int32(32)
					return v39
				} else {
					v37 = F_DirectFunctionCall1Coll(m, int32(1424), int32(0), v7)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = v37
						m.G0 = v7 + int32(32)
						return v39
					}
				}
			}
		default:
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v39 = v2
			m.G0 = v7 + int32(32)
			return v39
		}
	}
}
func F_inet_spg_consistent_bitmap(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
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
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(1)
	goto L3
L2:
	;
	v22 = int32(15)
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
	return v22
L5:
	;
	goto L6
L6:
	;
	v26 = int32(1)
	v27 = l0 + v26
	v29 = l0 + int32(4)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v30&v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = v27
	goto L9
L8:
	;
	v33 = v29
	goto L9
L9:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	v36 = int32(base.Ui32(v34) >> (uint(int32(3)) % 32))
	v42 = int32(1) << (uint((v34^int32(-1))&int32(7)) % 32)
	v47 = v22
	v56 = int32(0)
	goto L10
L10:
	;
	v64 = l2 + v56*int32(48)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+44))
	v66 = F_pg_detoast_datum_packed(m, v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return v503
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+6)))
	v71 = int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v75 = v73 & v71
	if v75 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v498 = v56 + int32(1)
	if v498 != l1 {
		v47 = v488
		v56 = v498
		goto L10
	} else {
		goto L195
	}
L16:
	;
	v76 = v71
	goto L18
L17:
	;
	v76 = int32(4)
	goto L18
L18:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v76))))
	v79 = int32(1)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v83 = v81 & v79
	if v83 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v84 = v79
	goto L21
L20:
	;
	v84 = int32(4)
	goto L21
L21:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v84))))
	if v78 != v86 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v88 = int32(0)
	switch v70 - int32(19) {
	case 0:
		goto L25
	case 1, 2:
		goto L27
	case 3, 4:
		goto L26
	default:
		v503 = v88
		goto L12
	}
L23:
	;
	goto L24
L24:
	;
	v94 = v66 + int32(1)
	v96 = v66 + int32(4)
	if v75 != 0 {
		goto L32
	} else {
		goto L33
	}
L25:
	;
	if v47 != 0 {
		v488 = v47
		goto L15
	} else {
		goto L31
	}
L26:
	;
	if base.Ui32(v86) < base.Ui32(v78) {
		v503 = v88
		goto L12
	} else {
		goto L30
	}
L27:
	;
	if base.Ui32(v78) < base.Ui32(v86) {
		v503 = v88
		goto L12
	} else {
		goto L28
	}
L28:
	;
	if v47 != 0 {
		v488 = v47
		goto L15
	} else {
		goto L29
	}
L29:
	;
	v503 = v88
	goto L12
L30:
	;
	goto L25
L31:
	;
	v503 = v88
	goto L12
L32:
	;
	v97 = v94
	goto L34
L33:
	;
	v97 = v96
	goto L34
L34:
	;
	v99 = v70 - int32(18)
	switch v99 {
	case 0:
		goto L37
	default:
		v135 = v47
		goto L35
	case 6:
		goto L36
	case 7:
		goto L40
	case 8:
		goto L39
	case 9:
		goto L38
	}
L35:
	;
	v137 = int32(0)
	if v135 == v137 {
		v503 = v137
		goto L12
	} else {
		goto L61
	}
L36:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if base.Ui32(v132) < base.Ui32(v34) {
		goto L58
	} else {
		goto L59
	}
L37:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if base.Ui32(v34) < base.Ui32(v121) {
		goto L52
	} else {
		goto L53
	}
L38:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v114 == v34 {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if v34 == v105-int32(1) {
		goto L44
	} else {
		goto L45
	}
L40:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if base.Ui32(v34) < base.Ui32(v102) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v104 = v47 & int32(12)
	goto L43
L42:
	;
	v104 = v47
	goto L43
L43:
	;
	v135 = v104
	goto L35
L44:
	;
	v135 = v47 & int32(3)
	goto L35
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(v34) < base.Ui32(v105) {
		v135 = v47
		goto L35
	} else {
		goto L47
	}
L47:
	;
	return int32(0)
L48:
	;
	v135 = v47 & int32(3)
	goto L35
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v34) <= base.Ui32(v114) {
		v135 = v47
		goto L35
	} else {
		goto L51
	}
L51:
	;
	return int32(0)
L52:
	;
	v135 = v47 & int32(12)
	goto L35
L53:
	;
	goto L54
L54:
	;
	if v34 == v121 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v135 = v47 & int32(3)
	goto L35
L56:
	;
	goto L57
L57:
	;
	return int32(0)
L58:
	;
	v134 = v47
	goto L60
L59:
	;
	v134 = v47 & int32(12)
	goto L60
L60:
	;
	v135 = v134
	goto L35
L61:
	;
	if v83 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v140 = v27
	goto L64
L63:
	;
	v140 = v29
	goto L64
L64:
	;
	v141 = int32(2)
	v142 = v140 + v141
	v144 = v97 + v141
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	if base.Ui32(v34) < base.Ui32(v145) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v147 = v34
	goto L67
L66:
	;
	v147 = v145
	goto L67
L67:
	;
	v152 = base.I32_div_s(v147, int32(8))
	v153 = F_memcmp(m, v142, v144, v152)
	mBase = m.M
	if v153 != 0 {
		v239 = v153
		goto L70
	} else {
		goto L71
	}
L68:
	;
	if v249 != 0 {
		goto L89
	} else {
		goto L90
	}
L69:
	;
	if v240 != 0 {
		goto L86
	} else {
		goto L87
	}
L70:
	;
	v249 = v239
	goto L68
L71:
	;
	v154 = int32(0)
	v157 = v147 - v152<<(uint(int32(3))%32)
	if v157 <= v154 {
		v239 = v154
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142+v152))))
	v162 = int32(128)
	v163 = v161 & v162
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144+v152))))
	if v163 != v165&v162 {
		v240 = v163
		goto L69
	} else {
		goto L73
	}
L73:
	;
	if v157 == int32(1) {
		v239 = v154
		goto L70
	} else {
		goto L74
	}
L74:
	;
	v171 = int32(1)
	v173 = int32(128)
	v174 = v161 << (uint(v171) % 32) & v173
	if v174 != v165<<(uint(v171)%32)&v173 {
		v240 = v174
		goto L69
	} else {
		goto L75
	}
L75:
	;
	if v157 < int32(3) {
		v239 = v154
		goto L70
	} else {
		goto L76
	}
L76:
	;
	v182 = int32(2)
	v184 = int32(128)
	v185 = v161 << (uint(v182) % 32) & v184
	if v185 != v165<<(uint(v182)%32)&v184 {
		v240 = v185
		goto L69
	} else {
		goto L77
	}
L77:
	;
	if v157 == int32(3) {
		v239 = v154
		goto L70
	} else {
		goto L78
	}
L78:
	;
	v193 = int32(3)
	v195 = int32(128)
	v196 = v161 << (uint(v193) % 32) & v195
	if v196 != v165<<(uint(v193)%32)&v195 {
		v240 = v196
		goto L69
	} else {
		goto L79
	}
L79:
	;
	if v157 < int32(5) {
		v239 = v154
		goto L70
	} else {
		goto L80
	}
L80:
	;
	v204 = int32(4)
	v206 = int32(128)
	v207 = v161 << (uint(v204) % 32) & v206
	if v207 != v165<<(uint(v204)%32)&v206 {
		v240 = v207
		goto L69
	} else {
		goto L81
	}
L81:
	;
	if v157 == int32(5) {
		v239 = v154
		goto L70
	} else {
		goto L82
	}
L82:
	;
	v215 = int32(5)
	v217 = int32(128)
	v218 = v161 << (uint(v215) % 32) & v217
	if v218 != v165<<(uint(v215)%32)&v217 {
		v240 = v218
		goto L69
	} else {
		goto L83
	}
L83:
	;
	if v157 < int32(7) {
		v239 = v154
		goto L70
	} else {
		goto L84
	}
L84:
	;
	v226 = int32(6)
	v228 = int32(128)
	v229 = v161 << (uint(v226) % 32) & v228
	if v229 != v165<<(uint(v226)%32)&v228 {
		v240 = v229
		goto L69
	} else {
		goto L85
	}
L85:
	;
	v239 = v154
	goto L70
L86:
	;
	v243 = int32(1)
	goto L88
L87:
	;
	v243 = int32(-1)
	goto L88
L88:
	;
	v249 = v243
	goto L68
L89:
	;
	switch v70 - int32(19) {
	case 0:
		v488 = v135
		goto L15
	case 1, 2:
		goto L93
	case 3, 4:
		goto L92
	default:
		v503 = v137
		goto L12
	}
L90:
	;
	goto L91
L91:
	;
	if v135&int32(12) == int32(0) {
		v282 = v135
		goto L98
	} else {
		goto L99
	}
L92:
	;
	if int32(0) <= v249 {
		v488 = v135
		goto L15
	} else {
		goto L95
	}
L93:
	;
	if v249 <= int32(0) {
		v488 = v135
		goto L15
	} else {
		goto L94
	}
L94:
	;
	v503 = v137
	goto L12
L95:
	;
	v503 = v137
	goto L12
L96:
	;
	if v315&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L97:
	;
	switch v70 - int32(20) {
	case 0, 1:
		goto L117
	case 2, 3:
		goto L116
	default:
		goto L115
	}
L98:
	;
	if base.Ui32((v70-int32(24))&int32(_a_F_inet_spg_consistent_bitmap_0)) < base.Ui32(int32(_a_F_inet_spg_consistent_bitmap_1)) {
		v488 = v282
		goto L15
	} else {
		goto L114
	}
L99:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v260&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v263 = v94
	goto L102
L101:
	;
	v263 = v96
	goto L102
L102:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+1)))
	if base.Ui32(v264) <= base.Ui32(v34) {
		v282 = v135
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263+v36)+2)))
	v268 = v42 & v267
	switch v70 - int32(19) {
	case 0:
		v314 = v135
		v315 = v260
		goto L96
	case 1, 2:
		goto L109
	case 3, 4:
		goto L108
	default:
		goto L107
	}
L104:
	;
	if v279 == int32(0) {
		v503 = v137
		goto L12
	} else {
		goto L113
	}
L105:
	;
	v279 = v135 & int32(11)
	goto L104
L106:
	;
	v279 = v135 & int32(7)
	goto L104
L107:
	;
	if v268 != 0 {
		goto L105
	} else {
		goto L112
	}
L108:
	;
	if v268 == int32(0) {
		v291 = v135
		goto L97
	} else {
		goto L111
	}
L109:
	;
	if v268 == int32(0) {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v291 = v135
	goto L97
L111:
	;
	goto L105
L112:
	;
	goto L106
L113:
	;
	v282 = v279
	goto L98
L114:
	;
	v291 = v282
	goto L97
L115:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v314 = v291
	v315 = v313
	goto L96
L116:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v305&int32(1) != 0 {
		goto L126
	} else {
		goto L127
	}
L117:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v296&int32(1) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v299 = v94
	goto L120
L119:
	;
	v299 = v96
	goto L120
L120:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299)+1)))
	if v300 == v34 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v303 = v291 & int32(3)
	if v303 != 0 {
		v314 = v303
		v315 = v296
		goto L96
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	if base.Ui32(v34) <= base.Ui32(v300) {
		v314 = v291
		v315 = v296
		goto L96
	} else {
		goto L125
	}
L124:
	;
	v503 = v137
	goto L12
L125:
	;
	v503 = v137
	goto L12
L126:
	;
	v308 = v94
	goto L128
L127:
	;
	v308 = v96
	goto L128
L128:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+1)))
	if base.Ui32(v309) <= base.Ui32(v34) {
		v314 = v291
		v315 = v305
		goto L96
	} else {
		goto L129
	}
L129:
	;
	v312 = v291 & int32(12)
	if v312 != 0 {
		v314 = v312
		v315 = v305
		goto L96
	} else {
		goto L130
	}
L130:
	;
	v503 = v137
	goto L12
L131:
	;
	v320 = v94
	goto L133
L132:
	;
	v320 = v96
	goto L133
L133:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320)+1)))
	if v34 != v321 {
		v488 = v314
		goto L15
	} else {
		goto L134
	}
L134:
	;
	if l3 != 0 {
		v350 = v314
		goto L135
	} else {
		goto L136
	}
L135:
	;
	if l3 == int32(0) {
		v488 = v350
		goto L15
	} else {
		goto L152
	}
L136:
	;
	if v314&int32(3) == int32(0) {
		v350 = v314
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320))))
	if v329 == int32(2) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v332 = int32(32)
	goto L140
L139:
	;
	v332 = int32(128)
	goto L140
L140:
	;
	if base.Ui32(v332) <= base.Ui32(v34) {
		v350 = v314
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v320+v36)+2)))
	v336 = v42 & v335
	switch v70 - int32(19) {
	case 0:
		v350 = v314
		goto L135
	case 1, 2:
		goto L147
	case 3, 4:
		goto L146
	default:
		goto L145
	}
L142:
	;
	if v347 == int32(0) {
		v503 = v137
		goto L12
	} else {
		goto L151
	}
L143:
	;
	v347 = v314 & int32(14)
	goto L142
L144:
	;
	v347 = v314 & int32(13)
	goto L142
L145:
	;
	if v336 != 0 {
		goto L143
	} else {
		goto L150
	}
L146:
	;
	if v336 == int32(0) {
		v350 = v314
		goto L135
	} else {
		goto L149
	}
L147:
	;
	if v336 == int32(0) {
		goto L144
	} else {
		goto L148
	}
L148:
	;
	v350 = v314
	goto L135
L149:
	;
	goto L143
L150:
	;
	goto L144
L151:
	;
	v350 = v347
	goto L135
L152:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v356 = v354 & int32(1)
	if v356 != 0 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v357 = v27
	goto L155
L154:
	;
	v357 = v29
	goto L155
L155:
	;
	v358 = int32(2)
	v359 = v357 + v358
	v361 = v320 + v358
	if v356 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v366 = int32(1)
	goto L158
L157:
	;
	v366 = int32(4)
	goto L158
L158:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v366))))
	if v368 == int32(2) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v371 = int32(32)
	goto L161
L160:
	;
	v371 = int32(128)
	goto L161
L161:
	;
	v376 = base.I32_div_s(v371, int32(8))
	v377 = F_memcmp(m, v359, v361, v376)
	mBase = m.M
	if v377 != 0 {
		v463 = v377
		goto L164
	} else {
		goto L165
	}
L162:
	;
	switch v99 - int32(1) {
	case 0:
		goto L183
	case 1:
		goto L188
	case 2:
		goto L187
	case 3:
		goto L184
	case 4:
		goto L185
	default:
		goto L186
	}
L163:
	;
	if v464 != 0 {
		goto L180
	} else {
		goto L181
	}
L164:
	;
	v473 = v463
	goto L162
L165:
	;
	v378 = int32(0)
	v381 = v371 - v376<<(uint(int32(3))%32)
	if v381 <= v378 {
		v463 = v378
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+v376))))
	v386 = int32(128)
	v387 = v385 & v386
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361+v376))))
	if v387 != v389&v386 {
		v464 = v387
		goto L163
	} else {
		goto L167
	}
L167:
	;
	if v381 == int32(1) {
		v463 = v378
		goto L164
	} else {
		goto L168
	}
L168:
	;
	v395 = int32(1)
	v397 = int32(128)
	v398 = v385 << (uint(v395) % 32) & v397
	if v398 != v389<<(uint(v395)%32)&v397 {
		v464 = v398
		goto L163
	} else {
		goto L169
	}
L169:
	;
	if v381 < int32(3) {
		v463 = v378
		goto L164
	} else {
		goto L170
	}
L170:
	;
	v406 = int32(2)
	v408 = int32(128)
	v409 = v385 << (uint(v406) % 32) & v408
	if v409 != v389<<(uint(v406)%32)&v408 {
		v464 = v409
		goto L163
	} else {
		goto L171
	}
L171:
	;
	if v381 == int32(3) {
		v463 = v378
		goto L164
	} else {
		goto L172
	}
L172:
	;
	v417 = int32(3)
	v419 = int32(128)
	v420 = v385 << (uint(v417) % 32) & v419
	if v420 != v389<<(uint(v417)%32)&v419 {
		v464 = v420
		goto L163
	} else {
		goto L173
	}
L173:
	;
	if v381 < int32(5) {
		v463 = v378
		goto L164
	} else {
		goto L174
	}
L174:
	;
	v428 = int32(4)
	v430 = int32(128)
	v431 = v385 << (uint(v428) % 32) & v430
	if v431 != v389<<(uint(v428)%32)&v430 {
		v464 = v431
		goto L163
	} else {
		goto L175
	}
L175:
	;
	if v381 == int32(5) {
		v463 = v378
		goto L164
	} else {
		goto L176
	}
L176:
	;
	v439 = int32(5)
	v441 = int32(128)
	v442 = v385 << (uint(v439) % 32) & v441
	if v442 != v389<<(uint(v439)%32)&v441 {
		v464 = v442
		goto L163
	} else {
		goto L177
	}
L177:
	;
	if v381 < int32(7) {
		v463 = v378
		goto L164
	} else {
		goto L178
	}
L178:
	;
	v450 = int32(6)
	v452 = int32(128)
	v453 = v385 << (uint(v450) % 32) & v452
	if v453 != v389<<(uint(v450)%32)&v452 {
		v464 = v453
		goto L163
	} else {
		goto L179
	}
L179:
	;
	v463 = v378
	goto L164
L180:
	;
	v467 = int32(1)
	goto L182
L181:
	;
	v467 = int32(-1)
	goto L182
L182:
	;
	v473 = v467
	goto L162
L183:
	;
	if v473 == int32(0) {
		v503 = v137
		goto L12
	} else {
		goto L194
	}
L184:
	;
	if int32(0) < v473 {
		v488 = v350
		goto L15
	} else {
		goto L193
	}
L185:
	;
	if int32(0) <= v473 {
		v488 = v350
		goto L15
	} else {
		goto L192
	}
L186:
	;
	if v473 == int32(0) {
		v488 = v350
		goto L15
	} else {
		goto L191
	}
L187:
	;
	if v473 <= int32(0) {
		v488 = v350
		goto L15
	} else {
		goto L190
	}
L188:
	;
	if v473 < int32(0) {
		v488 = v350
		goto L15
	} else {
		goto L189
	}
L189:
	;
	v503 = v137
	goto L12
L190:
	;
	v503 = v137
	goto L12
L191:
	;
	v503 = v137
	goto L12
L192:
	;
	v503 = v137
	goto L12
L193:
	;
	v503 = v137
	goto L12
L194:
	;
	v488 = v350
	goto L15
L195:
	;
	v503 = v488
	goto L12
}
