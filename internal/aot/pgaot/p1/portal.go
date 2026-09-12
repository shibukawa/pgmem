package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PortalRun(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v298 int32
	_ = v298
	var v313 int32
	_ = v313
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	v7 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(48)
	m.G0 = v24
	v34 = v7
	v35 = v7
	v36 = v7
	v37 = v7
	v38 = v7
	v39 = v7
	v40 = v7
	v41 = v7
	v42 = v7
	v43 = v7
	v44 = v7
	v45 = v24
	v46 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	m.G0 = v24 + int32(48)
	return v362 & int32(1)
L3:
	;
	goto L2
L4:
	;
	if v46 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v421 = int32(m.ExcTag)
	v422 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v421 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L7:
	;
	v51 = v45 - int32(160)
	m.G0 = v51
	v54 = base.B2i32(l5 == int32(0))
	if l5 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v167 = v34
	v168 = v35
	v169 = v36
	v170 = v37
	v171 = v38
	v172 = v39
	v173 = v40
	v174 = v41
	v175 = v42
	v176 = v43
	v177 = v44
	v178 = v45
	goto L9
L9:
	;
	if v177 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, _consts[903])))
	if v70 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	F_MarkPortalActive(m, l0)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		v418 = v51
		goto L6
	} else {
		goto L24
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v73 == int32(4) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	v88 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		v418 = v51
		goto L6
	} else {
		goto L17
	}
L17:
	;
	if v88 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	F_errmsg_internal(m, int32(235423), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		v418 = v51
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	F_getrusage(m, int32(4378856))
	mBase = m.M
	F___gettimeofday(m, int32(4379008))
	mBase = m.M
	goto L23
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v54)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v51
	F_errfinish(m, int32(475686), int32(708), int32(235423))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		v418 = v51
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L14
L24:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v149 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	v151 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v153 = *(*int32)(unsafe.Add(mBase, _consts[904]))
	v155 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v157 = *(*int32)(unsafe.Add(mBase, _consts[905]))
	v159 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v161 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v24 + int32(4)
	goto L28
L26:
	;
	v167 = v151
	v168 = v155
	v169 = v51
	v170 = v54
	v171 = v149
	v172 = v147
	v173 = v153
	v174 = v157
	v175 = v159
	v176 = v161
	v177 = int32(0)
	v178 = v51
	goto L9
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v172
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v171
	*(*int32)(unsafe.Add(mBase, _consts[905])) = v174
	*(*int32)(unsafe.Add(mBase, _consts[904])) = v173
	v375 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v167 == v175 {
		goto L63
	} else {
		goto L64
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[905])) = l0
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v169
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v185 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v172
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v328 = v170 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v328)
	F_MarkPortalFailed(m, l0)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L55
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v185
	goto L36
L35:
	;
	goto L36
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v189
	*(*int32)(unsafe.Add(mBase, _consts[904])) = v189
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	switch v193 {
	case 0:
		goto L39
	case 1, 2, 3:
		goto L40
	case 4:
		goto L38
	default:
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v278 = v170 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v278)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L52
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v243 = v170 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v243)
	F_PortalRunMulti(m, l0, l2, int32(0), l3, l4, l5)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L47
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v218 = int32(1)
	v219 = v170 & v218
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v219)
	v222 = F_PortalRunSelect(m, l0, v218, l1, l3)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L43
	}
L40:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v194 != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v205 = v170 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v205)
	F_FillPortalStore(m, l0, l2)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	if v219 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(2)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+117)))
	v362 = v232
	goto L29
L45:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v224 == int32(0) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l5)+8)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v224
	goto L44
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v257 = int32(1)
	v259 = v170 & v257
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v259)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = int32(4)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v263 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	m.T0[v263].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v362 = v257
	goto L29
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L50
L52:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v278)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v284
	F_errmsg_internal(m, int32(464286), v24)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v278)
	F_errfinish(m, int32(475686), int32(800), int32(235423))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L30
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[905])) = v174
	*(*int32)(unsafe.Add(mBase, _consts[904])) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	v340 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	if v167 == v175 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v342 = v340
	goto L58
L57:
	;
	v342 = v167
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v342
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	v348 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	if v168 == v176 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v350 = v348
	goto L61
L60:
	;
	v350 = v168
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v328)
	F_pg_re_throw(m)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L62
	}
L62:
	;
	goto L30
L63:
	;
	v377 = v375
	goto L65
L64:
	;
	v377 = v167
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v377
	v381 = *(*int32)(unsafe.Add(mBase, _consts[176]))
	if v168 == v176 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v383 = v381
	goto L68
L67:
	;
	v383 = v168
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, _consts[175])) = v383
	v386 = int32(*(*uint8)(unsafe.Add(mBase, _consts[903])))
	if v386 != int32(1) {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v389 == int32(4) {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v24)+24)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v24)+32)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v169
	v402 = v170 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)) = uint8(v402)
	F_ShowUsage(m, int32(506543))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		v418 = v178
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L5
L72:
	;
	v426 = int32(v422)
	m.G0 = v418
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v24+int32(4) == v433 {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	m.ExcPending = 1
	goto L81
L74:
	;
	if v436 != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v429)+4))
	v436 = v435
	goto L77
L76:
	;
	v436 = int32(0)
	goto L77
L77:
	;
	goto L74
L78:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v24)+28))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v34 = v444
	v35 = v442
	v36 = v437
	v37 = v438
	v38 = v445
	v39 = v446
	v40 = v443
	v41 = v441
	v42 = v440
	v43 = v439
	v44 = v428
	v45 = v418
	v46 = v436
	goto L1
L79:
	;
	goto L80
L80:
	;
	F___wasm_longjmp(m, v429, v428)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	return int32(0)
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
