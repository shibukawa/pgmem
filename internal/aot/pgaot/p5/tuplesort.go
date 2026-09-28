package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_tuplesort_estimate_shared(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = F_mul_size(m, int32(8), l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_add_size(m, v3, int32(72))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return (v8 + int32(7)) & int32(-8)
		}
	}
}
func F_tuplesort_getgintuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(_a_F_tuplesort_getgintuple_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v13
	v18 = F_tuplesort_gettuple_common(m, l0, int32(1), v8+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if v18 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v11
			v34 = v3
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getgintuple[0])) = v11
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			if v28 == int32(0) {
				v34 = v3
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
				v34 = v28
			}
		}
		m.G0 = v8 + int32(32)
		return v34
	}
}
func F_tuplesort_getheaptuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int32(_a_F_tuplesort_getheaptuple_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0])) = v11
	v16 = F_tuplesort_gettuple_common(m, l0, int32(1), v6+int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_getheaptuple[0])) = v9
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		m.G0 = v6 + int32(32)
		if v16 != 0 {
			v27 = v22
		} else {
			v27 = int32(0)
		}
		return v27
	}
}
func F_tuplesort_gettuple_common(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v228 int64
	_ = v228
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v303 int64
	_ = v303
	var v305 int64
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v326 int32
	_ = v326
	var v327 int64
	_ = v327
	var v329 int64
	_ = v329
	var v331 int64
	_ = v331
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int64
	_ = v404
	var v406 int64
	_ = v406
	var v408 int64
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v429 int32
	_ = v429
	var v430 int64
	_ = v430
	var v432 int64
	_ = v432
	var v434 int64
	_ = v434
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	switch v17 - int32(3) {
	case 0:
		goto L8
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L17
	} else {
		goto L141
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L17
	} else {
		goto L138
	}
L3:
	;
	m.G0 = v15 + int32(32)
	return v453
L4:
	;
	v450 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v450)
	v453 = v100
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L17
	} else {
		goto L135
	}
L6:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v201 != 0 {
		goto L70
	} else {
		goto L71
	}
L7:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v20 < v21 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v60 = int32(0)
	if v20 <= v60 {
		v453 = v60
		goto L3
	} else {
		goto L21
	}
L12:
	;
	v23 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v20 + v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v30 = v27 + v20*int32(24)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v33
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v35
	v453 = v23
	goto L3
L13:
	;
	goto L14
L14:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v37)
	v39 = int32(0)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)))
	if v40 != v37 {
		v453 = v39
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v20 < v43 {
		v453 = v39
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_0), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1395), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L21:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if v63 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v75 = int32(24)
	v79 = v74 + v73*v75 - v75
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v79)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v79)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v84
	v453 = int32(1)
	goto L3
L23:
	;
	v66 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v66)
	v73 = v20
	goto L22
L24:
	;
	goto L25
L25:
	;
	v68 = int32(1)
	v69 = v20 - v68
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v69
	if v20 == v68 {
		v453 = v60
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v73 = v69
	goto L22
L27:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v87) < base.Ui32(v88) {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	goto L29
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)))
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L29
L31:
	;
	F_pfree(m, v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L17
	} else {
		goto L34
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v90) <= base.Ui32(v87) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v92
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v87
	goto L30
L34:
	;
	goto L30
L35:
	;
	v100 = int32(0)
	if v99&int32(1) != 0 {
		v453 = v100
		goto L3
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	if v99&int32(1) != 0 {
		goto L45
	} else {
		goto L46
	}
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v105 = F_LogicalTapeRead(m, v103, v15, int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	if v105 != int32(4) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v109 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v113].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v112, v109)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L17
	} else {
		goto L42
	}
L42:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v116
	v453 = int32(1)
	goto L3
L43:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v188 = F_getlen(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L17
	} else {
		goto L66
	}
L44:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+208)) = uint8(v181)
	goto L43
L45:
	;
	v124 = F_LogicalTapeBackspace(m, v119, int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L17
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v139 = int32(0)
	v141 = F_LogicalTapeBackspace(m, v119, int32(4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L55
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L17
	} else {
		goto L50
	}
L49:
	;
	switch v124 {
	case 0:
		v453 = int32(0)
		goto L3
	default:
		goto L48
	case 8:
		goto L44
	}
L50:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_3), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1478), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v157 = F_getlen(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L17
	} else {
		goto L59
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	switch v141 {
	case 0:
		v453 = v139
		goto L3
	default:
		goto L54
	case 4:
		goto L53
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_3), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1492), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L17
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v161 = v157 + int32(8)
	v162 = F_LogicalTapeBackspace(m, v159, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L17
	} else {
		goto L60
	}
L60:
	;
	if v162 == v157+int32(4) {
		v453 = v139
		goto L3
	} else {
		goto L61
	}
L61:
	;
	if v161 == v162 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L17
	} else {
		goto L63
	}
L63:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_4), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1512), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L17
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v191 = F_LogicalTapeBackspace(m, v190, v188)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L17
	} else {
		goto L67
	}
L67:
	;
	if v191 != v188 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v195].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v194, v188)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L17
	} else {
		goto L69
	}
L69:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v198
	v453 = int32(1)
	goto L3
L70:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if base.Ui32(v201) < base.Ui32(v202) {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v213 <= int32(0) {
		goto L78
	} else {
		goto L79
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	goto L72
L74:
	;
	F_pfree(m, v201)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L17
	} else {
		goto L77
	}
L75:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if base.Ui32(v204) <= base.Ui32(v201) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v201))) = v206
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v201
	goto L73
L77:
	;
	goto L73
L78:
	;
	v453 = int32(0)
	goto L3
L79:
	;
	goto L80
L80:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217+v219<<(uint(int32(2))%32))))
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v218)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v218)))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v228
	*(*uint32)(unsafe.Add(mBase, uint32(l0)+164)) = uint32(v228)
	v234 = F_LogicalTapeRead(m, v223, v15+int32(28), int32(4))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	if v234 != int32(4) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	if v238 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v242 = int32(1)
	v243 = v241 - v242
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v243
	if int32(0) < v243 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	m.T0[v351].(func(*base.Module, int32, int32, int32, int32))(m, l0, v15, v223, v238)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L17
	} else {
		goto L112
	}
L86:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettuple_common[0]))
	if v252 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v345 - int32(1)
	F_LogicalTapeClose(m, v223)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L17
	} else {
		goto L111
	}
L89:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	v256 = v243
	goto L91
L91:
	;
	v257 = v243*int32(24) + v248
	if base.Ui32(v256) < base.Ui32(int32(2)) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v256 = v255
	goto L91
L93:
	;
	v326 = v248 + v314*int32(24)
	v327 = *(*int64)(unsafe.Add(mBase, uint32(v257)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v326)+16)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v257)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v326)+8)) = v329
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v257)))
	*(*int64)(unsafe.Add(mBase, uint32(v326))) = v331
	goto L88
L94:
	;
	v314 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v265 = int32(1)
	v266 = v4
	v272 = v4
	goto L97
L97:
	;
	v275 = v272 + int32(2)
	if base.Ui32(v256) <= base.Ui32(v275) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v314 = v289
	goto L93
L99:
	;
	v289 = v265
	goto L101
L100:
	;
	v277 = int32(24)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v284 = m.T0[v283].(func(*base.Module, int32, int32, int32) int32)(m, v248+v265*v277, v248+v275*v277, l0)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L17
	} else {
		goto L102
	}
L101:
	;
	v292 = v248 + v289*int32(24)
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v294 = m.T0[v293].(func(*base.Module, int32, int32, int32) int32)(m, v257, v292, l0)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L17
	} else {
		goto L106
	}
L102:
	;
	if int32(0) < v284 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v288 = v275
	goto L105
L104:
	;
	v288 = v265
	goto L105
L105:
	;
	v289 = v288
	goto L101
L106:
	;
	if v294 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v314 = v266
	goto L93
L108:
	;
	goto L109
L109:
	;
	v300 = v248 + v266*int32(24)
	v301 = *(*int64)(unsafe.Add(mBase, uint32(v292)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v300)+16)) = v301
	v303 = *(*int64)(unsafe.Add(mBase, uint32(v292)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v300)+8)) = v303
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v292)))
	*(*int64)(unsafe.Add(mBase, uint32(v300))) = v305
	v307 = int32(1)
	v308 = v289 << (uint(v307) % 32)
	v310 = v308 | v307
	if base.Ui32(v310) < base.Ui32(v256) {
		v265 = v310
		v266 = v289
		v272 = v308
		goto L97
	} else {
		goto L110
	}
L110:
	;
	goto L98
L111:
	;
	v453 = v242
	goto L3
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v219
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettuple_common[0]))
	if v357 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L17
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if base.Ui32(v360) < base.Ui32(int32(2)) {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	v429 = v355 + v417*int32(24)
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v429)+16)) = v430
	v432 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v429)+8)) = v432
	v434 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v429))) = v434
	v453 = int32(1)
	goto L3
L118:
	;
	v417 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v366 = int32(1)
	v369 = v4
	v371 = v4
	goto L121
L121:
	;
	v378 = v371 + int32(2)
	if base.Ui32(v360) <= base.Ui32(v378) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v417 = v392
	goto L117
L123:
	;
	v392 = v366
	goto L125
L124:
	;
	v380 = int32(24)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v387 = m.T0[v386].(func(*base.Module, int32, int32, int32) int32)(m, v355+v366*v380, v355+v378*v380, l0)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L17
	} else {
		goto L126
	}
L125:
	;
	v395 = v355 + v392*int32(24)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v397 = m.T0[v396].(func(*base.Module, int32, int32, int32) int32)(m, v15, v395, l0)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L17
	} else {
		goto L130
	}
L126:
	;
	if int32(0) < v387 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v391 = v378
	goto L129
L128:
	;
	v391 = v366
	goto L129
L129:
	;
	v392 = v391
	goto L125
L130:
	;
	if v397 <= int32(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v417 = v369
	goto L117
L132:
	;
	goto L133
L133:
	;
	v403 = v355 + v369*int32(24)
	v404 = *(*int64)(unsafe.Add(mBase, uint32(v395)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+16)) = v404
	v406 = *(*int64)(unsafe.Add(mBase, uint32(v395)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v403)+8)) = v406
	v408 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
	*(*int64)(unsafe.Add(mBase, uint32(v403))) = v408
	v410 = int32(1)
	v411 = v392 << (uint(v410) % 32)
	v413 = v411 | v410
	if base.Ui32(v413) < base.Ui32(v360) {
		v366 = v413
		v369 = v392
		v371 = v411
		goto L121
	} else {
		goto L134
	}
L134:
	;
	goto L122
L135:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_5), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L17
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1595), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L17
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
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_4), int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L17
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(1525), int32(_a_F_tuplesort_gettuple_common_2))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L17
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_errmsg_internal(m, int32(_a_F_tuplesort_gettuple_common_6), int32(0))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L17
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_tuplesort_gettuple_common_1), int32(3173), int32(_a_F_tuplesort_gettuple_common_7))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tuplesort_gettupleslot(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(_a_F_tuplesort_gettupleslot_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v15
	v19 = F_tuplesort_gettuple_common(m, l0, l1, v10+int32(8))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v13
			*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
			m.T0[v52].(func(*base.Module, int32))(m, l3)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v57 = int32(0)
				m.G0 = v10 + int32(32)
				return v57
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_tuplesort_gettupleslot[0])) = v13
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			if v31 == int32(0) {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
				m.T0[v52].(func(*base.Module, int32))(m, l3)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v57 = int32(0)
					m.G0 = v10 + int32(32)
					return v57
				}
			} else {
				if l4 == int32(0) {
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
					if v37 == int32(0) {
					} else {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
						*(*int64)(unsafe.Add(mBase, uint32(l4))) = v40
					}
				}
				if l2 != 0 {
					v43 = F_heap_copy_minimal_tuple(m, v31, int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v43
						v46 = v43
						v47 = F_ExecStoreMinimalTuple(m, v46, l3, l2)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							v57 = int32(1)
							m.G0 = v10 + int32(32)
							return v57
						}
					}
				} else {
					v46 = v31
					v47 = F_ExecStoreMinimalTuple(m, v46, l3, l2)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v57 = int32(1)
						m.G0 = v10 + int32(32)
						return v57
					}
				}
			}
		}
	}
}
func F_tuplesort_initialize_shared(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	v4 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v4))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	F_SharedFileSetInit(m, l0+int32(12), l2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = l1
		if l1 <= int32(0) {
		} else {
			v17 = l1 << (uint(int32(3)) % 32)
			if v17 == int32(0) {
			} else {
				base.MemoryFill(m, l0+int32(72), int32(0), v17)
			}
		}
		return
	}
}
