package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FetchPortalTargetList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v3 != int32(4) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v9 != 0 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v46 = int32(0)
	goto L3
L3:
	;
	return v46
L4:
	;
	v42 = F_FetchStatementTargetList(m, v40)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	goto L4
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v10 <= int32(0) {
		v40 = int32(0)
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v40 = int32(0)
	goto L5
L9:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v16 = v10
	goto L12
L11:
	;
	v16 = v13
	goto L12
L12:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v19 = int32(0)
	goto L13
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v17+v19<<(uint(int32(2))%32))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+26)))
	if v27 == int32(1) {
		v40 = v26
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L8
L15:
	;
	v31 = v19 + int32(1)
	if v31 != v16 {
		v19 = v31
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return int32(0)
L18:
	;
	v46 = v42
	goto L3
}
func F_PortalRunFetch(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v151 int64
	_ = v151
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int64
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	var v219 int64
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v267 int32
	_ = v267
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v305 int64
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v319 int64
	_ = v319
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int64
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int64
	_ = v353
	var v383 int32
	_ = v383
	var v397 int32
	_ = v397
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v411 int64
	_ = v411
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(32)
	m.G0 = v20
	v28 = v5
	v29 = v5
	v30 = v5
	v31 = v5
	v32 = v5
	v33 = v5
	v34 = v5
	v35 = v20
	v36 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v36 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v410 = int32(m.ExcTag)
	v411 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v410 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L7:
	;
	v43 = v35 - int32(160)
	m.G0 = v43
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v43
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		v405 = v43
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v69 = v28
	v70 = v29
	v71 = v30
	v72 = v31
	v73 = v32
	v74 = v33
	v75 = v34
	v76 = v35
	goto L9
L9:
	;
	if v70 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v57 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	v59 = *(*int32)(unsafe.Add(mBase, _consts[413]))
	v61 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v63 = *(*int32)(unsafe.Add(mBase, _consts[259]))
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v20 + int32(4)
	goto L14
L12:
	;
	v69 = v43
	v70 = int32(0)
	v71 = v57
	v72 = v55
	v73 = v59
	v74 = v61
	v75 = v63
	v76 = v43
	goto L9
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[259])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v69
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v83 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v72
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L118
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v83
	goto L20
L19:
	;
	goto L20
L20:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v87
	v89 = int32(4520272)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v87
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if base.Ui32(int32(3)) <= base.Ui32(v93-int32(1)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	switch l1 {
	case 0:
		goto L38
	case 1:
		goto L42
	case 2:
		goto L41
	case 3:
		goto L40
	default:
		goto L39
	}
L22:
	;
	if v93 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v131 != 0 {
		goto L21
	} else {
		goto L29
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_errmsg_internal(m, int32(20583), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_errfinish(m, int32(493485), int32(1434), int32(325260))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_FillPortalStore(m, l0, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L21
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v72
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v71
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v90
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v74
	*(*int32)(unsafe.Add(mBase, _consts[259])) = v75
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v73
	m.G0 = v20 + int32(32)
	return v353
L32:
	;
	v353 = base.I64_extend_i32_u(v344) & int64(1)
	goto L31
L33:
	;
	v342 = F_PortalRunSelect(m, l0, v340, v339, l3)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L117
	}
L34:
	;
	v331 = int32(1)
	v335 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v336 = F_PortalRunSelect(m, l0, int32(0), v331, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L116
	}
L35:
	;
	v308 = l2 >> (uint(int32(31)) % 32)
	v310 = l2 ^ v308 - v308
	if base.B2i32(v310 != int32(2147483647))|v287 != 0 {
		v339 = v310
		v340 = v287
		goto L33
	} else {
		goto L106
	}
L36:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v289 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L37:
	;
	if l2 != 0 {
		goto L35
	} else {
		goto L97
	}
L38:
	;
	v287 = base.B2i32(int32(0) <= l2)
	goto L37
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L94
	}
L40:
	;
	if int32(0) < l2 {
		goto L80
	} else {
		goto L81
	}
L41:
	;
	if int32(0) < l2 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v287 = int32(base.Ui32(l2) >> (uint(int32(31)) % 32))
	goto L37
L43:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v155 = l2 - int32(1)
	if base.B2i32(base.Ui64(v151) <= base.Ui64(int64(2147483646)))&base.B2i32(base.Ui64(int64(base.Ui64(v151)>>(uint(int64(1))%64))) < base.Ui64(base.I64_extend_i32_u(v155))) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	goto L45
L45:
	;
	if l2 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L46:
	;
	v196 = int32(1)
	v198 = F_PortalRunSelect(m, l0, v196, v196, l3)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L59
	}
L47:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v174 = v172 + base.I32_wrap_i64(v151)
	if base.Ui32(l2) <= base.Ui32(v174) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if l2 == int32(1) {
		goto L46
	} else {
		goto L51
	}
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v170 = F_PortalRunSelect(m, l0, int32(1), v155, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L52
	}
L52:
	;
	goto L46
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v182 = F_PortalRunSelect(m, l0, int32(0), v174-l2+int32(1), v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(l2) <= base.Ui32(v174+int32(1)) {
		goto L46
	} else {
		goto L57
	}
L56:
	;
	goto L46
L57:
	;
	v192 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v193 = F_PortalRunSelect(m, l0, int32(1), l2+(v174^int32(-1)), v192)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L58
	}
L58:
	;
	goto L46
L59:
	;
	v353 = v198
	goto L31
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v206 = F_PortalRunSelect(m, l0, int32(1), int32(2147483647), v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L69
	}
L63:
	;
	if l2 != int32(-1) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v215 = F_PortalRunSelect(m, l0, int32(0), l2^int32(-1), v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v219 = F_PortalRunSelect(m, l0, int32(0), int32(1), l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	v353 = v219
	goto L31
L69:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v223 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v223)+20)) = l3
	goto L72
L71:
	;
	goto L72
L72:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v225 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v228 = F_RunFromStore(m, l0, int32(0), int64(0), l3)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	F_PushActiveSnapshot(m, v230)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L77
	}
L76:
	;
	v353 = v228
	goto L31
L77:
	;
	F_ExecutorRun(m, v223, int32(0), int64(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L78
	}
L78:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v223)+40))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v237)+112))
	F_PopActiveSnapshot(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L79
	}
L79:
	;
	v353 = v238
	goto L31
L80:
	;
	if l2 != int32(1) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if int32(0) <= l2 {
		v288 = int32(1)
		goto L36
	} else {
		goto L88
	}
L83:
	;
	v246 = int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v251 = F_PortalRunSelect(m, l0, v246, l2-v246, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v253 = int32(1)
	v255 = F_PortalRunSelect(m, l0, v253, v253, l3)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v353 = v255
	goto L31
L88:
	;
	if l2 != int32(-1) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[922]))
	v266 = F_PortalRunSelect(m, l0, int32(0), l2^int32(-1), v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v270 = F_PortalRunSelect(m, l0, int32(0), int32(1), l3)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L93
	}
L92:
	;
	goto L91
L93:
	;
	v353 = v270
	goto L31
L94:
	;
	F_errmsg_internal(m, int32(255427), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(493485), int32(1605), int32(325258))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	v288 = v287
	goto L36
L98:
	;
	if v292&int32(1) == int32(0) {
		goto L34
	} else {
		goto L104
	}
L99:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v293 != 0 {
		goto L98
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v296 = int32(0)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v297 == v296 {
		v344 = v296
		goto L32
	} else {
		goto L103
	}
L102:
	;
	v344 = v292 ^ int32(1)
	goto L32
L103:
	;
	v339 = v296
	v340 = v288
	goto L33
L104:
	;
	v305 = F_PortalRunSelect(m, l0, v288, int32(0), l3)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L105
	}
L105:
	;
	v353 = v305
	goto L31
L106:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v314 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v317 = F_PortalRunSelect(m, l0, int32(0), int32(2147483647), l3)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v319 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if v319 == int64(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	v353 = v317
	goto L31
L111:
	;
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L114
	}
L112:
	;
	goto L113
L113:
	;
	v325 = int64(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	F_DoPortalRewind(m, l0)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L115
	}
L114:
	;
	v353 = int64(0)
	goto L31
L115:
	;
	v353 = v319 - (v325 ^ int64(1))
	goto L31
L116:
	;
	v339 = int32(1)
	v340 = v331
	goto L33
L117:
	;
	v353 = v342
	goto L31
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v74
	*(*int32)(unsafe.Add(mBase, _consts[259])) = v75
	*(*int32)(unsafe.Add(mBase, _consts[413])) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v69
	F_pg_re_throw(m)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		v405 = v76
		goto L6
	} else {
		goto L119
	}
L119:
	;
	goto L5
L120:
	;
	v415 = int32(v411)
	m.G0 = v405
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v415)+4))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v415)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if v20+int32(4) == v422 {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	m.ExcPending = 1
	goto L129
L122:
	;
	if v425 != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v425 = v424
	goto L125
L124:
	;
	v425 = int32(0)
	goto L125
L125:
	;
	goto L122
L126:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v28 = v426
	v29 = v417
	v30 = v430
	v31 = v431
	v32 = v429
	v33 = v428
	v34 = v427
	v35 = v405
	v36 = v425
	goto L1
L127:
	;
	goto L128
L128:
	;
	F___wasm_longjmp(m, v418, v417)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	return int64(0)
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
