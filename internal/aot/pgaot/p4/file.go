package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_copy_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	v8 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v15 = F_palloc(m, int32(65536))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v18 = F_OpenTransientFile(m, l0, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L62
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L58
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L54
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L50
	}
L7:
	;
	if int32(0) <= v18 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v23 = F_OpenTransientFile(m, l1, int32(194))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L46
	}
L11:
	;
	if v23 < int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v34 = v8
	v35 = v8
	goto L13
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v44 < v34 {
		goto L37
	} else {
		goto L38
	}
L15:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if int64(1048576) <= v34-v35 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	v43 = F_fsync(m, v23)
	mBase = m.M
	goto L22
L20:
	;
	v44 = v35
	goto L21
L21:
	;
	v45 = int32(4102524)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(167772175)
	v50 = F_read(m, v18, v15, int32(65536))
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v53
	if v50 < v53 {
		goto L5
	} else {
		goto L23
	}
L22:
	;
	v44 = v34
	goto L21
L23:
	;
	if v50 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(167772176)
	v64 = F_write(m, v23, v15, v50)
	mBase = m.M
	if v64 != v50 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	goto L14
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v67 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
	v34 = v34 + base.I64_extend_i32_u(v50)
	v35 = v44
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(51)
	goto L32
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = l1
	F_errmsg(m, int32(296551), v12+int32(80))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(491941), int32(213), int32(384002))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v97 = F_fsync(m, v23)
	mBase = m.M
	goto L40
L38:
	;
	goto L39
L39:
	;
	v98 = F_CloseTransientFile(m, v23)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	if v98 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v100 = F_CloseTransientFile(m, v18)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v100 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v15)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	m.G0 = v12 + int32(96)
	return
L46:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg(m, int32(296868), v12)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(491941), int32(167), int32(384002))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	F_errmsg(m, int32(297547), v12+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(491941), int32(173), int32(384002))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l0
	F_errmsg(m, int32(297787), v12+int32(32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(491941), int32(201), int32(384002))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l1
	F_errmsg(m, int32(297611), v12-int32(-64))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(491941), int32(224), int32(384002))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l0
	F_errmsg(m, int32(297611), v12+int32(48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(491941), int32(229), int32(384002))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_sendFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v168 int64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v218 int64
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v294 int64
	_ = v294
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v377 int64
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int64
	_ = v450
	var v454 int32
	_ = v454
	var v455 int64
	_ = v455
	var v459 int32
	_ = v459
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	v19 = m.G0
	v21 = v19 - int32(8288)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[307]))) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[308]))) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[309]))) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	v30 = F_pg_checksum_init(m, v21+int32(8268), v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L3
	} else {
		goto L130
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L3
	} else {
		goto L127
	}
L3:
	;
	return int32(0)
L4:
	;
	if int32(0) <= v30 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v37 = F_OpenTransientFile(m, l1, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L3
	} else {
		goto L124
	}
L8:
	;
	m.G0 = v21 + int32(8288)
	return int32(base.Ui32(v37^int32(-1)) >> (uint(int32(31)) % 32))
L9:
	;
	if v37 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if l4 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v62 = int32(0)
	F__tarWriteHeader(m, l0, l2, v62, l3, v62)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L3
	} else {
		goto L21
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v42 == int32(44) {
		goto L8
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L3
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = l1
	F_errmsg(m, int32(296868), v21+int32(48))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(492539), int32(1600), int32(387745))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L3
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
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[310])))
	if v68 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[275]))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+252))
	goto L25
L23:
	;
	v79 = v62
	goto L24
L24:
	;
	if l11 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v79 = base.B2i32(v73 != int32(0)) & base.B2i32(l7 != int32(0))
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[311]))) = int32(-743563507)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[312]))) = int32(0)
	F_push_to_sink(m, l0, v21+int32(8268), v21+int32(8260), v21+int32(8264), int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	v168 = int64(0)
	goto L28
L28:
	;
	v170 = l8 << (uint(int32(17)) % 32)
	v171 = int32(0)
	v177 = v171
	v180 = v171
	v185 = v79
	v189 = v168
	goto L48
L29:
	;
	F_push_to_sink(m, l0, v21+int32(8268), v21+int32(8260), v21+int32(8284), int32(4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_push_to_sink(m, l0, v21+int32(8268), v21+int32(8260), v21+int32(8280), int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v116 = l10 << (uint(int32(2)) % 32)
	F_push_to_sink(m, l0, v21+int32(8268), v21+int32(8260), l11, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	if l10 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v148 != 0 {
		goto L42
	} else {
		goto L43
	}
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[312])))
	v121 = v119 & int32(8191)
	if v121 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v146 = int64(12)
	goto L36
L36:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[312])))
	v148 = v147
	v150 = v146
	goto L33
L37:
	;
	v148 = v119
	v150 = int64(12)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v129 = int32(8192) - v121
	v131 = F__emscripten_memset_bulkmem(m, v21-int32(-64), base.I32_extend8_s(int32(0)), v129)
	mBase = m.M
	goto L40
L40:
	;
	F_push_to_sink(m, l0, v21+int32(8268), v21+int32(8260), v21-int32(-64), v129)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v146 = base.I64_extend_i32_u(int32(8204) - v121)
	goto L36
L42:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+8))
	m.T0[v152].(func(*base.Module, int32, int32))(m, l0, v148)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v168 = v150 + base.I64_extend_i32_u(v116)
	goto L28
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = F_pg_checksum_update(m, v21+int32(8268), v157, v148)
	mBase = m.M
	if v158 < int32(0) {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v189 < v294 {
		goto L76
	} else {
		goto L77
	}
L48:
	;
	if l11 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L73
	}
L50:
	;
	if v185&int32(1) == int32(0) {
		goto L60
	} else {
		goto L61
	}
L51:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v193 <= v189 {
		goto L47
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(l10) <= base.Ui32(v177) {
		goto L47
	} else {
		goto L56
	}
L54:
	;
	v202 = F_read_file_data_into_buffer(m, l0, l1, v37, v189, base.I32_wrap_i64(v193-v189), v180+v170, v185&int32(1), v21+int32(8276))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	v224 = v177
	v226 = v202
	goto L50
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l11+v177<<(uint(int32(2))%32))))
	v218 = F_read_file_data_into_buffer(m, l0, l1, v37, base.I64_extend_i32_u(v208<<(uint(int32(13))%32)), int32(8192), v208+v170, v185&int32(1), v21+int32(8276))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	if v218 < int64(8192) {
		goto L47
	} else {
		goto L58
	}
L58:
	;
	v224 = v177 + int32(1)
	v226 = v218
	goto L50
L59:
	;
	if v226 == int64(0) {
		goto L47
	} else {
		goto L70
	}
L60:
	;
	v259 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v226&int64(8191) == int64(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v259 = int32(1)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v237 = int32(0)
	v240 = F_errstart(m, int32(19), v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	if v240 == int32(0) {
		v259 = v237
		goto L59
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = int32(8192)
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+40)) = uint32(v226)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	F_errmsg(m, int32(225712), v21+int32(32))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(492539), int32(1756), int32(387745))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v259 = v237
	goto L59
L70:
	;
	v262 = base.I32_wrap_i64(v226)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+8))
	m.T0[v264].(func(*base.Module, int32, int32))(m, l0, v262)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v269 = base.I32_div_s(v262, int32(8192))
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v274 = F_pg_checksum_update(m, v21+int32(8268), v273, v262)
	mBase = m.M
	if int32(0) <= v274 {
		v177 = v224
		v180 = v269 + v180
		v185 = v259
		v189 = v189 + v226
		goto L48
	} else {
		goto L72
	}
L72:
	;
	goto L49
L73:
	;
	F_errmsg_internal(m, int32(232245), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(492539), int32(1785), int32(387745))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	v312 = v189
	v313 = v294
	goto L79
L77:
	;
	v377 = v189
	goto L78
L78:
	;
	v379 = base.I32_wrap_i64(v377)
	v384 = (v379+int32(511))&int32(-512) - v379
	if int32(0) < v384 {
		goto L97
	} else {
		goto L98
	}
L79:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v316 = base.I32_wrap_i64(v313 - v312)
	if base.Ui32(v314) < base.Ui32(v316) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v377 = v358
	goto L78
L81:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v350 = F_pg_checksum_update(m, v21+int32(8268), v349, v318)
	mBase = m.M
	if v350 < int32(0) {
		goto L1
	} else {
		goto L94
	}
L82:
	;
	v344 = F__emscripten_memset_bulkmem(m, v319, base.I32_extend8_s(int32(0)), v341)
	mBase = m.M
	goto L93
L83:
	;
	v318 = v314
	goto L85
L84:
	;
	v318 = v316
	goto L85
L85:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v319&int32(3) != 0 {
		v341 = v318
		goto L82
	} else {
		goto L86
	}
L86:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v318) {
		v341 = v318
		goto L82
	} else {
		goto L87
	}
L87:
	;
	if v318&int32(3) != 0 {
		v341 = v318
		goto L82
	} else {
		goto L88
	}
L88:
	;
	v326 = v319 + v318
	if base.Ui32(v326) <= base.Ui32(v319) {
		goto L81
	} else {
		goto L89
	}
L89:
	;
	v331 = v319 + int32(4)
	if base.Ui32(v331) < base.Ui32(v326) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v333 = v326
	goto L92
L91:
	;
	v333 = v331
	goto L92
L92:
	;
	v341 = (v319^int32(-1)+v333)&int32(-4) + int32(4)
	goto L82
L93:
	;
	goto L81
L94:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v353)+8))
	m.T0[v354].(func(*base.Module, int32, int32))(m, l0, v318)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L3
	} else {
		goto L95
	}
L95:
	;
	v358 = v312 + base.I64_extend_i32_u(v318)
	v359 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v358 < v359 {
		v312 = v358
		v313 = v359
		goto L79
	} else {
		goto L96
	}
L96:
	;
	goto L80
L97:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v387&int32(3) != 0 {
		v409 = v384
		goto L101
	} else {
		goto L102
	}
L98:
	;
	goto L99
L99:
	;
	v422 = F_CloseTransientFile(m, v37)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L3
	} else {
		goto L111
	}
L100:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v415)+8))
	m.T0[v416].(func(*base.Module, int32, int32))(m, l0, v384)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L3
	} else {
		goto L110
	}
L101:
	;
	v412 = F__emscripten_memset_bulkmem(m, v387, base.I32_extend8_s(int32(0)), v409)
	mBase = m.M
	goto L109
L102:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v384) {
		v409 = v384
		goto L101
	} else {
		goto L103
	}
L103:
	;
	if v384&int32(3) != 0 {
		v409 = v384
		goto L101
	} else {
		goto L104
	}
L104:
	;
	v394 = v387 + v384
	if base.Ui32(v394) <= base.Ui32(v387) {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v399 = v387 + int32(4)
	if base.Ui32(v399) < base.Ui32(v394) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v401 = v394
	goto L108
L107:
	;
	v401 = v399
	goto L108
L108:
	;
	v409 = (v387^int32(-1)+v401)&int32(-4) + int32(4)
	goto L101
L109:
	;
	goto L100
L110:
	;
	goto L99
L111:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_consts[309])))
	if int32(2) <= v424 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v429 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v448 = int32(4383208)
	v450 = *(*int64)(unsafe.Add(mBase, _consts[313]))
	*(*int64)(unsafe.Add(mBase, _consts[313])) = v450 + base.I64_extend_i32_s(v424)
	v454 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v455 = *(*int64)(unsafe.Add(mBase, uint32(l3)+56))
	F_AddFileToBackupManifest(m, l9, l6, l2, v454, v455, v21+int32(8268))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L3
	} else {
		goto L123
	}
L115:
	;
	if v429 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l1
	F_errmsg_plural(m, int32(361174), int32(160707), v424, v21+int32(16))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	F_pgstat_prepare_report_checksum_failure(m, l5)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	F_errfinish(m, int32(492539), int32(1818), int32(387745))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L3
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	F_pgstat_report_checksum_failures_in_db(m, l5, v424)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	goto L114
L123:
	;
	goto L8
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
	F_errmsg_internal(m, int32(698577), v21)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L3
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(492539), int32(1591), int32(387745))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L3
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errmsg_internal(m, int32(232245), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(492539), int32(1666), int32(387745))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_errmsg_internal(m, int32(232245), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(492539), int32(1798), int32(387745))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L3
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
