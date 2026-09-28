package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ExecOnConflictLockRow(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	v7 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+64))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l3)+188))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+104))
	v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l3, l2, v15, l1, v16, l4, v7, v7, v11+int32(28))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L73
	}
L2:
	;
	m.G0 = v11 + int32(48)
	return v251
L3:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+12))
	m.T0[v247].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L72
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L69
	}
L5:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[0]))
	if v214 < int32(2) {
		goto L3
	} else {
		goto L64
	}
L6:
	;
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[0]))
	if v194 < int32(2) {
		goto L3
	} else {
		goto L59
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L56
	}
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v32 = m.T0[v31].(func(*base.Module, int32, int32, int32) int64)(m, l1, int32(-2), v11+int32(27))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	switch v23 {
	case 0:
		v251 = int32(1)
		goto L2
	case 1:
		goto L8
	case 2:
		goto L7
	case 3:
		goto L6
	case 4:
		goto L5
	default:
		goto L4
	}
L11:
	;
	v34 = base.I32_wrap_i64(v32)
	if base.Ui32(v34) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L52
	}
L13:
	;
	v166 = int32(0)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[1]))
	if v46 == v34 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v166 = int32(1)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[2]))
	if v50 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v166 = v156
	goto L12
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[3]))
	if v54 == int32(0) {
		v156 = int32(0)
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ExecOnConflictLockRow[4]))
	v126 = int32(0)
	v129 = v50 - int32(1)
	goto L42
L23:
	;
	v59 = v54
	goto L24
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	if v65 == int32(4) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v156 = int32(0)
	goto L19
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v59)+80))
	if v119 != 0 {
		v59 = v119
		goto L24
	} else {
		goto L41
	}
L27:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v68 == int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v71 = int32(1)
	if v34 == v68 {
		v156 = v71
		goto L19
	} else {
		goto L29
	}
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v59)+52))
	v75 = v73 - int32(1)
	if v75 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v59)+48))
	v81 = int32(0)
	v84 = v75
	goto L31
L31:
	;
	v89 = int32(2)
	v90 = base.I32_div_s(v84-v81, v89)
	v91 = v90 + v81
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v78+v91<<(uint(v89)%32))))
	if v95 == v34 {
		v156 = v71
		goto L19
	} else {
		goto L33
	}
L32:
	;
	goto L26
L33:
	;
	v104 = base.B2i32(v95-v34 < int32(0)) | base.B2i32(base.Ui32(v95) < base.Ui32(int32(3)))
	if v104 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v105 = v91 + int32(1)
	goto L36
L35:
	;
	v105 = v81
	goto L36
L36:
	;
	if v104 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v108 = v84
	goto L39
L38:
	;
	v108 = v91 - int32(1)
	goto L39
L39:
	;
	if v105 <= v108 {
		v81 = v105
		v84 = v108
		goto L31
	} else {
		goto L40
	}
L40:
	;
	goto L32
L41:
	;
	goto L25
L42:
	;
	v134 = int32(2)
	v135 = base.I32_div_s(v129-v126, v134)
	v136 = v135 + v126
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v124+v136<<(uint(v134)%32))))
	v141 = base.B2i32(v140 == v34)
	if v140 == v34 {
		v156 = v141
		goto L19
	} else {
		goto L44
	}
L43:
	;
	v156 = v141
	goto L19
L44:
	;
	v144 = base.B2i32(base.Ui32(v140) < base.Ui32(v34))
	if base.Ui32(v140) < base.Ui32(v34) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v145 = v136 + int32(1)
	goto L47
L46:
	;
	v145 = v126
	goto L47
L47:
	;
	if base.Ui32(v140) < base.Ui32(v34) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v148 = v129
	goto L50
L49:
	;
	v148 = v136 - int32(1)
	goto L50
L50:
	;
	if v145 <= v148 {
		v126 = v145
		v129 = v148
		goto L42
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	if v166 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_ExecOnConflictLockRow_0), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2833), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
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
	F_errmsg_internal(m, int32(_a_F_ExecOnConflictLockRow_3), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2843), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_ExecOnConflictLockRow_4), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2850), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_ExecOnConflictLockRow_5), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2866), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v23
	F_errmsg_internal(m, int32(_a_F_ExecOnConflictLockRow_6), v11)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2873), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v251 = int32(0)
	goto L2
L73:
	;
	if l5 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v261 = int32(_a_F_ExecOnConflictLockRow_7)
	goto L76
L75:
	;
	v261 = int32(_a_F_ExecOnConflictLockRow_8)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v261
	F_errmsg(m, int32(_a_F_ExecOnConflictLockRow_9), v11+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L77
	}
L77:
	;
	F_errhint(m, int32(_a_F_ExecOnConflictLockRow_10), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_ExecOnConflictLockRow_1), int32(2830), int32(_a_F_ExecOnConflictLockRow_2))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_on_ppath(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v28 float64
	_ = v28
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v63 float64
	_ = v63
	var v64 int32
	_ = v64
	var v65 float64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 int64
	_ = v90
	var v92 int64
	_ = v92
	var v95 float64
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v111 float64
	_ = v111
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v147 float64
	_ = v147
	var v154 int32
	_ = v154
	var v155 float64
	_ = v155
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v193 float64
	_ = v193
	var v194 int32
	_ = v194
	var v195 float64
	_ = v195
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v200 float64
	_ = v200
	var v211 float64
	_ = v211
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v232 int64
	_ = v232
	var v234 int64
	_ = v234
	var v237 float64
	_ = v237
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v253 float64
	_ = v253
	var v261 float64
	_ = v261
	var v266 float64
	_ = v266
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v277 float64
	_ = v277
	var v278 float64
	_ = v278
	var v280 float64
	_ = v280
	var v282 float64
	_ = v282
	var v289 float64
	_ = v289
	var v295 float64
	_ = v295
	var v297 float64
	_ = v297
	var v307 float64
	_ = v307
	var v308 int32
	_ = v308
	var v309 float64
	_ = v309
	var v312 int32
	_ = v312
	var v313 float64
	_ = v313
	var v314 float64
	_ = v314
	var v315 float64
	_ = v315
	var v317 float64
	_ = v317
	var v330 float64
	_ = v330
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v335 float64
	_ = v335
	var v337 float64
	_ = v337
	var v350 float64
	_ = v350
	var v351 int32
	_ = v351
	var v352 float64
	_ = v352
	var v353 int64
	_ = v353
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 float64
	_ = v366
	var v367 float64
	_ = v367
	var v370 int32
	_ = v370
	var v371 float64
	_ = v371
	var v372 int64
	_ = v372
	var v374 int64
	_ = v374
	var v377 float64
	_ = v377
	var v380 int64
	_ = v380
	var v382 int64
	_ = v382
	var v393 float64
	_ = v393
	var v401 float64
	_ = v401
	var v406 float64
	_ = v406
	var v407 float64
	_ = v407
	var v408 float64
	_ = v408
	var v417 float64
	_ = v417
	var v418 float64
	_ = v418
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v429 float64
	_ = v429
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v461 int64
	_ = v461
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v461
L4:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v26 = base.F64_sub(v24, v25)
	v28 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v26), v28)|base.F64_eq(base.F64_abs(v24), v28)|base.F64_eq(base.F64_abs(v25), v28) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v444 = F_point_inside(m, v14, v20, v16+int32(16))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L105
	}
L7:
	;
	v41 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v43 = v26
	goto L9
L9:
	;
	v45 = v20 - int32(1)
	v46 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	v48 = base.F64_sub(v46, v47)
	v50 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v48), v50)|base.F64_eq(base.F64_abs(v46), v50)|base.F64_eq(base.F64_abs(v47), v50) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v43 = v41
	goto L9
L11:
	;
	v63 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v65 = v48
	goto L13
L13:
	;
	v67 = v16 + int32(16)
	v68 = int32(0)
	if v68 < v45 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v65 = v63
	goto L13
L15:
	;
	v71 = v45
	goto L17
L16:
	;
	v71 = v68
	goto L17
L17:
	;
	v80 = m.G0
	v82 = v80 - int32(32)
	m.G0 = v82
	v84 = base.F64_abs(v43)
	v85 = base.F64_abs(v65)
	v88 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)))
	if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v154 = int32(0)
	v155 = v147
	goto L38
L19:
	;
	m.G0 = v82 + int32(32)
	goto L18
L20:
	;
	v89 = v84
	goto L22
L21:
	;
	v89 = v85
	goto L22
L22:
	;
	v90 = base.I64_reinterpret_f64(v89)
	v92 = int64(base.Ui64(v90) >> (uint(int64(52)) % 64))
	if v92 == int64(2047) {
		v147 = v89
		goto L19
	} else {
		goto L23
	}
L23:
	;
	if base.Ui64(base.I64_reinterpret_f64(v84)) < base.Ui64(base.I64_reinterpret_f64(v85)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v95 = v85
	goto L26
L25:
	;
	v95 = v84
	goto L26
L26:
	;
	if v90 == int64(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v147 = v95
	goto L19
L28:
	;
	v98 = base.I64_reinterpret_f64(v95)
	v100 = int64(base.Ui64(v98) >> (uint(int64(52)) % 64))
	if v100 == int64(2047) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if int32(65) <= base.I32_wrap_i64(v100)-base.I32_wrap_i64(v92) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v147 = base.F64_add(v84, v85)
	goto L19
L31:
	;
	goto L32
L32:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v98) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_sq(m, v82+int32(24), v82+int32(16), v124)
	mBase = m.M
	F_sq(m, v82+int32(8), v82, v125)
	mBase = m.M
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
	v136 = *(*float64)(unsafe.Add(mBase, uint32(v82)+16))
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v82)+8))
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v82)+24))
	v147 = base.F64_mul(v126, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v135, v136), v138), v140)))
	goto L19
L34:
	;
	v111 = float64(1.90109156629516e-211)
	v124 = base.F64_mul(v95, v111)
	v125 = base.F64_mul(v89, v111)
	v126 = float64(5.260135901548374e+210)
	goto L33
L35:
	;
	goto L36
L36:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v90) {
		v124 = v95
		v125 = v89
		v126 = float64(1)
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v119 = float64(5.260135901548374e+210)
	v124 = base.F64_mul(v95, v119)
	v125 = base.F64_mul(v89, v119)
	v126 = float64(1.90109156629516e-211)
	goto L33
L38:
	;
	if v154 == v71 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v461 = v353
	goto L3
L40:
	;
	return int64(0)
L41:
	;
	goto L42
L42:
	;
	v171 = math.Float64frombits(uint64(0x7ff0000000000000))
	v173 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v175 = v154 + int32(1)
	v178 = v67 + v175<<(uint(int32(4))%32)
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v178)))
	v180 = base.F64_sub(v173, v179)
	if base.F64_ne(base.F64_abs(v180), v171)|base.F64_eq(base.F64_abs(v173), v171)|base.F64_eq(base.F64_abs(v179), v171) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v195 = v180
	goto L45
L44:
	;
	v193 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v197 = *(*float64)(unsafe.Add(mBase, uint32(v178)+8))
	v198 = base.F64_sub(v196, v197)
	v200 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v198), v200)|base.F64_eq(base.F64_abs(v196), v200)|base.F64_eq(base.F64_abs(v197), v200) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v195 = v193
	goto L45
L47:
	;
	v213 = v198
	goto L49
L48:
	;
	v211 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v222 = m.G0
	v224 = v222 - int32(32)
	m.G0 = v224
	v226 = base.F64_abs(v195)
	v227 = base.F64_abs(v213)
	v230 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)))
	if base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v213 = v211
	goto L49
L51:
	;
	v295 = base.F64_add(v155, v289)
	v297 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v155), v171)|base.F64_ne(base.F64_abs(v295), v297)|base.F64_eq(base.F64_abs(v289), v297) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L52:
	;
	m.G0 = v224 + int32(32)
	goto L51
L53:
	;
	v231 = v226
	goto L55
L54:
	;
	v231 = v227
	goto L55
L55:
	;
	v232 = base.I64_reinterpret_f64(v231)
	v234 = int64(base.Ui64(v232) >> (uint(int64(52)) % 64))
	if v234 == int64(2047) {
		v289 = v231
		goto L52
	} else {
		goto L56
	}
L56:
	;
	if base.Ui64(base.I64_reinterpret_f64(v226)) < base.Ui64(base.I64_reinterpret_f64(v227)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v237 = v227
	goto L59
L58:
	;
	v237 = v226
	goto L59
L59:
	;
	if v232 == int64(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v289 = v237
	goto L52
L61:
	;
	v240 = base.I64_reinterpret_f64(v237)
	v242 = int64(base.Ui64(v240) >> (uint(int64(52)) % 64))
	if v242 == int64(2047) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if int32(65) <= base.I32_wrap_i64(v242)-base.I32_wrap_i64(v234) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v289 = base.F64_add(v226, v227)
	goto L52
L64:
	;
	goto L65
L65:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v240) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	F_sq(m, v224+int32(24), v224+int32(16), v266)
	mBase = m.M
	F_sq(m, v224+int32(8), v224, v267)
	mBase = m.M
	v277 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
	v278 = *(*float64)(unsafe.Add(mBase, uint32(v224)+16))
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v224)+8))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v224)+24))
	v289 = base.F64_mul(v268, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v277, v278), v280), v282)))
	goto L52
L67:
	;
	v253 = float64(1.90109156629516e-211)
	v266 = base.F64_mul(v237, v253)
	v267 = base.F64_mul(v231, v253)
	v268 = float64(5.260135901548374e+210)
	goto L66
L68:
	;
	goto L69
L69:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v232) {
		v266 = v237
		v267 = v231
		v268 = float64(1)
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v261 = float64(5.260135901548374e+210)
	v266 = base.F64_mul(v237, v261)
	v267 = base.F64_mul(v231, v261)
	v268 = float64(1.90109156629516e-211)
	goto L66
L71:
	;
	v307 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v309 = v295
	goto L73
L73:
	;
	v312 = v67 + v154<<(uint(int32(4))%32)
	v313 = *(*float64)(unsafe.Add(mBase, uint32(v312)))
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v178)))
	v315 = base.F64_sub(v313, v314)
	v317 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v315), v317)|base.F64_eq(base.F64_abs(v313), v317)|base.F64_eq(base.F64_abs(v314), v317) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v309 = v307
	goto L73
L75:
	;
	v330 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	v332 = v315
	goto L77
L77:
	;
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v312)+8))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v178)+8))
	v335 = base.F64_sub(v333, v334)
	v337 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v335), v337)|base.F64_eq(base.F64_abs(v333), v337)|base.F64_eq(base.F64_abs(v334), v337) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v332 = v330
	goto L77
L79:
	;
	v350 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	v352 = v335
	goto L81
L81:
	;
	v353 = int64(1)
	v362 = m.G0
	v364 = v362 - int32(32)
	m.G0 = v364
	v366 = base.F64_abs(v332)
	v367 = base.F64_abs(v352)
	v370 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v366)) < base.Ui64(base.I64_reinterpret_f64(v367)))
	if base.Ui64(base.I64_reinterpret_f64(v366)) < base.Ui64(base.I64_reinterpret_f64(v367)) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v352 = v350
	goto L81
L83:
	;
	if base.F64_eq(v309, v429) != 0 {
		v461 = v353
		goto L3
	} else {
		goto L103
	}
L84:
	;
	m.G0 = v364 + int32(32)
	goto L83
L85:
	;
	v371 = v366
	goto L87
L86:
	;
	v371 = v367
	goto L87
L87:
	;
	v372 = base.I64_reinterpret_f64(v371)
	v374 = int64(base.Ui64(v372) >> (uint(int64(52)) % 64))
	if v374 == int64(2047) {
		v429 = v371
		goto L84
	} else {
		goto L88
	}
L88:
	;
	if base.Ui64(base.I64_reinterpret_f64(v366)) < base.Ui64(base.I64_reinterpret_f64(v367)) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v377 = v367
	goto L91
L90:
	;
	v377 = v366
	goto L91
L91:
	;
	if v372 == int64(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v429 = v377
	goto L84
L93:
	;
	v380 = base.I64_reinterpret_f64(v377)
	v382 = int64(base.Ui64(v380) >> (uint(int64(52)) % 64))
	if v382 == int64(2047) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	if int32(65) <= base.I32_wrap_i64(v382)-base.I32_wrap_i64(v374) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v429 = base.F64_add(v366, v367)
	goto L84
L96:
	;
	goto L97
L97:
	;
	if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v380) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	F_sq(m, v364+int32(24), v364+int32(16), v406)
	mBase = m.M
	F_sq(m, v364+int32(8), v364, v407)
	mBase = m.M
	v417 = *(*float64)(unsafe.Add(mBase, uint32(v364)))
	v418 = *(*float64)(unsafe.Add(mBase, uint32(v364)+16))
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v364)+8))
	v422 = *(*float64)(unsafe.Add(mBase, uint32(v364)+24))
	v429 = base.F64_mul(v408, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v417, v418), v420), v422)))
	goto L84
L99:
	;
	v393 = float64(1.90109156629516e-211)
	v406 = base.F64_mul(v377, v393)
	v407 = base.F64_mul(v371, v393)
	v408 = float64(5.260135901548374e+210)
	goto L98
L100:
	;
	goto L101
L101:
	;
	if base.Ui64(int64(2580562586483294207)) < base.Ui64(v372) {
		v406 = v377
		v407 = v371
		v408 = float64(1)
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v401 = float64(5.260135901548374e+210)
	v406 = base.F64_mul(v377, v401)
	v407 = base.F64_mul(v371, v401)
	v408 = float64(1.90109156629516e-211)
	goto L98
L103:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v309, v429)), float64(1e-06)) == int32(0) {
		v154 = v175
		v155 = v289
		goto L38
	} else {
		goto L104
	}
L104:
	;
	goto L39
L105:
	;
	v461 = base.I64_extend_i32_u(base.B2i32(v444 != int32(0)))
	goto L3
}
func F_on_proc_exit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[0]))
	if v5 < int32(20) {
		v9 = v5 << (uint(int32(4)) % 32)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_on_proc_exit[1]))) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+uint32(_c_F_on_proc_exit[2]))) = l0
		*(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[0])) = v5 + int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_on_proc_exit[3])))
		if v22 == int32(0) {
			v27 = *(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[4]))
			if v27 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _c_F_on_proc_exit[4])) = v27 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27<<(uint(int32(2))%32))+uint32(_c_F_on_proc_exit[5]))) = int32(1196)
			} else {
			}
			v44 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_on_proc_exit[3])) = uint8(v44)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_on_proc_exit_0), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_on_proc_exit_1), int32(321), int32(_a_F_on_proc_exit_2))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_on_ps(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_lseg_contain_point(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_on_sb(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v48 int64
	_ = v48
	v8 = int64(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.F64_ge(v10, v12) == int32(0) {
		v48 = v8
	} else {
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		if base.F64_ge(v12, v16) == int32(0) {
			v48 = v8
		} else {
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
			if base.F64_ge(v20, v21) == int32(0) {
				v48 = v8
			} else {
				v25 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				if base.F64_ge(v21, v25) == int32(0) {
					v48 = v8
				} else {
					v29 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
					v31 = int32(0)
					if base.B2i32(base.F64_ge(v10, v29) == v31)|base.B2i32(base.F64_le(v16, v29) == v31) != 0 {
						v48 = v8
					} else {
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_ge(v20, v37) == int32(0) {
							v48 = v8
						} else {
							v48 = base.I64_extend_i32_u(base.F64_ge(v37, v25))
						}
					}
				}
			}
		}
	}
	return v48
}
