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
	v15 = F_palloc(m, int32(_a_F_copy_file_0))
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
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[0]))
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
	v45 = int32(_a_F_copy_file_1)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(167772175)
	v50 = F_read(m, v18, v15, int32(_a_F_copy_file_0))
	mBase = m.M
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[1]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_copy_file[2])) = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[1]))
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
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[2]))
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
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_copy_file[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
	v34 = v34 + base.I64_extend_i32_u(v50)
	v35 = v44
	goto L13
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_copy_file[2])) = int32(51)
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
	F_errmsg(m, int32(_a_F_copy_file_2), v12+int32(80))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(213), int32(_a_F_copy_file_4))
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
	F_errmsg(m, int32(_a_F_copy_file_5), v12)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(167), int32(_a_F_copy_file_4))
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
	F_errmsg(m, int32(_a_F_copy_file_6), v12+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(173), int32(_a_F_copy_file_4))
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
	F_errmsg(m, int32(_a_F_copy_file_7), v12+int32(32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(201), int32(_a_F_copy_file_4))
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
	F_errmsg(m, int32(_a_F_copy_file_8), v12-int32(-64))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(224), int32(_a_F_copy_file_4))
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
	F_errmsg(m, int32(_a_F_copy_file_8), v12+int32(48))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_copy_file_3), int32(229), int32(_a_F_copy_file_4))
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
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v155 int64
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v281 int64
	_ = v281
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v348 int64
	_ = v348
	var v366 int64
	_ = v366
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int64
	_ = v435
	var v439 int32
	_ = v439
	var v440 int64
	_ = v440
	var v444 int32
	_ = v444
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	v19 = m.G0
	v21 = v19 - int32(_a_F_sendFile_0)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[0]))) = l12
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[1]))) = l10
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[2]))) = int32(0)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	v30 = F_pg_checksum_init(m, v21+int32(_a_F_sendFile_1), v29)
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
	v500 = m.ExcPending
	if v500 != 0 {
		goto L3
	} else {
		goto L132
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L3
	} else {
		goto L129
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
	v474 = m.ExcPending
	if v474 != 0 {
		goto L3
	} else {
		goto L126
	}
L8:
	;
	m.G0 = v21 + int32(_a_F_sendFile_0)
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
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_sendFile[3]))
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
	F_errmsg(m, int32(_a_F_sendFile_2), v21+int32(48))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1600), int32(_a_F_sendFile_4))
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
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_sendFile[4])))
	if v68 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_sendFile[5]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[6]))) = int32(-743563507)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[7]))) = int32(0)
	v85 = v21 + int32(_a_F_sendFile_1)
	v87 = v21 + int32(_a_F_sendFile_5)
	F_push_to_sink(m, l0, v85, v87, v21+int32(_a_F_sendFile_6), int32(4))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	v155 = int64(0)
	goto L28
L28:
	;
	v157 = l8 << (uint(int32(17)) % 32)
	v158 = int32(0)
	v164 = v158
	v167 = v158
	v172 = v79
	v176 = v155
	goto L50
L29:
	;
	F_push_to_sink(m, l0, v85, v87, v21+int32(_a_F_sendFile_7), int32(4))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	F_push_to_sink(m, l0, v85, v87, v21+int32(_a_F_sendFile_8), int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v104 = l10 << (uint(int32(2)) % 32)
	F_push_to_sink(m, l0, v85, v87, l11, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
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
	if v135 != 0 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[7])))
	v109 = v107 & int32(_a_F_sendFile_9)
	if v109 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v133 = int64(12)
	goto L36
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[7])))
	v135 = v134
	v137 = v133
	goto L33
L37:
	;
	v135 = v107
	v137 = int64(12)
	goto L33
L38:
	;
	goto L39
L39:
	;
	v114 = int32(_a_F_sendFile_10) - v109
	if v114 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryFill(m, v21-int32(-64), int32(0), v114)
	goto L42
L41:
	;
	goto L42
L42:
	;
	F_push_to_sink(m, l0, v21+int32(_a_F_sendFile_1), v21+int32(_a_F_sendFile_5), v21-int32(-64), v114)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v133 = base.I64_extend_i32_u(int32(_a_F_sendFile_11) - v109)
	goto L36
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	m.T0[v139].(func(*base.Module, int32, int32))(m, l0, v135)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v155 = v137 + base.I64_extend_i32_u(v104)
	goto L28
L47:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v145 = F_pg_checksum_update(m, v21+int32(_a_F_sendFile_1), v144, v135)
	mBase = m.M
	if v145 < int32(0) {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v281 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v176 < v281 {
		goto L78
	} else {
		goto L79
	}
L50:
	;
	if l11 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L3
	} else {
		goto L75
	}
L52:
	;
	if v172&int32(1) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L53:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v180 <= v176 {
		goto L49
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.Ui32(l10) <= base.Ui32(v164) {
		goto L49
	} else {
		goto L58
	}
L56:
	;
	v189 = F_read_file_data_into_buffer(m, l0, l1, v37, v176, base.I32_wrap_i64(v180-v176), v167+v157, v172&int32(1), v21+int32(_a_F_sendFile_12))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	v211 = v164
	v213 = v189
	goto L52
L58:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l11+v164<<(uint(int32(2))%32))))
	v205 = F_read_file_data_into_buffer(m, l0, l1, v37, base.I64_extend_i32_u(v195<<(uint(int32(13))%32)), int32(_a_F_sendFile_10), v157+v195, v172&int32(1), v21+int32(_a_F_sendFile_12))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if v205 < int64(8192) {
		goto L49
	} else {
		goto L60
	}
L60:
	;
	v211 = v164 + int32(1)
	v213 = v205
	goto L52
L61:
	;
	if v213 == int64(0) {
		goto L49
	} else {
		goto L72
	}
L62:
	;
	v246 = int32(0)
	goto L61
L63:
	;
	goto L64
L64:
	;
	if v213&int64(8191) == int64(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v246 = int32(1)
	goto L61
L66:
	;
	goto L67
L67:
	;
	v224 = int32(0)
	v227 = F_errstart(m, int32(19), v224)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	if v227 == int32(0) {
		v246 = v224
		goto L61
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = int32(_a_F_sendFile_10)
	*(*uint32)(unsafe.Add(mBase, uint32(v21)+40)) = uint32(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l1
	F_errmsg(m, int32(_a_F_sendFile_13), v21+int32(32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L3
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1756), int32(_a_F_sendFile_4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	v246 = v224
	goto L61
L72:
	;
	v249 = base.I32_wrap_i64(v213)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+8))
	m.T0[v251].(func(*base.Module, int32, int32))(m, l0, v249)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v256 = base.I32_div_s(v249, int32(_a_F_sendFile_10))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v261 = F_pg_checksum_update(m, v21+int32(_a_F_sendFile_1), v260, v249)
	mBase = m.M
	if int32(0) <= v261 {
		v164 = v211
		v167 = v256 + v167
		v172 = v246
		v176 = v176 + v213
		goto L50
	} else {
		goto L74
	}
L74:
	;
	goto L51
L75:
	;
	F_errmsg_internal(m, int32(_a_F_sendFile_14), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1785), int32(_a_F_sendFile_4))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v299 = v176
	v300 = v281
	goto L81
L79:
	;
	v366 = v176
	goto L80
L80:
	;
	v368 = base.I32_wrap_i64(v366)
	v373 = (v368+int32(511))&int32(-512) - v368
	if int32(0) < v373 {
		goto L99
	} else {
		goto L100
	}
L81:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v303 = base.I32_wrap_i64(v300 - v299)
	if base.Ui32(v301) < base.Ui32(v303) {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	v366 = v347
	goto L80
L83:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v339 = F_pg_checksum_update(m, v21+int32(_a_F_sendFile_1), v338, v305)
	mBase = m.M
	if v339 < int32(0) {
		goto L1
	} else {
		goto L96
	}
L84:
	;
	if v329 == int32(0) {
		goto L83
	} else {
		goto L95
	}
L85:
	;
	v305 = v301
	goto L87
L86:
	;
	v305 = v303
	goto L87
L87:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v306&int32(3) != 0 {
		v329 = v305
		goto L84
	} else {
		goto L88
	}
L88:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v305) {
		v329 = v305
		goto L84
	} else {
		goto L89
	}
L89:
	;
	if v305&int32(3) != 0 {
		v329 = v305
		goto L84
	} else {
		goto L90
	}
L90:
	;
	if v305 == int32(0) {
		goto L83
	} else {
		goto L91
	}
L91:
	;
	v317 = v305 + v306
	v319 = v306 + int32(4)
	if base.Ui32(v319) < base.Ui32(v317) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v321 = v317
	goto L94
L93:
	;
	v321 = v319
	goto L94
L94:
	;
	v329 = (v306^int32(-1)+v321)&int32(-4) + int32(4)
	goto L84
L95:
	;
	base.MemoryFill(m, v306, int32(0), v329)
	goto L83
L96:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	m.T0[v343].(func(*base.Module, int32, int32))(m, l0, v305)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	v347 = v299 + base.I64_extend_i32_u(v305)
	v348 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
	if v347 < v348 {
		v299 = v347
		v300 = v348
		goto L81
	} else {
		goto L98
	}
L98:
	;
	goto L82
L99:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v376&int32(3) != 0 {
		v397 = v373
		goto L102
	} else {
		goto L103
	}
L100:
	;
	goto L101
L101:
	;
	v407 = F_CloseTransientFile(m, v37)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L3
	} else {
		goto L113
	}
L102:
	;
	if v397 != 0 {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v373) {
		v397 = v373
		goto L102
	} else {
		goto L104
	}
L104:
	;
	if v373&int32(3) != 0 {
		v397 = v373
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v385 = v373 + v376
	v387 = v376 + int32(4)
	if base.Ui32(v387) < base.Ui32(v385) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v389 = v385
	goto L108
L107:
	;
	v389 = v387
	goto L108
L108:
	;
	v397 = (v376^int32(-1)+v389)&int32(-4) + int32(4)
	goto L102
L109:
	;
	base.MemoryFill(m, v376, int32(0), v397)
	goto L111
L110:
	;
	goto L111
L111:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	m.T0[v401].(func(*base.Module, int32, int32))(m, l0, v373)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	goto L101
L113:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v21)+uint32(_c_F_sendFile[2])))
	if int32(2) <= v409 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v414 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L3
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v433 = int32(_a_F_sendFile_15)
	v435 = *(*int64)(unsafe.Add(mBase, _c_F_sendFile[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_sendFile[8])) = v435 + base.I64_extend_i32_s(v409)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v440 = *(*int64)(unsafe.Add(mBase, uint32(l3)+56))
	F_AddFileToBackupManifest(m, l9, l6, l2, v439, v440, v21+int32(_a_F_sendFile_1))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L3
	} else {
		goto L125
	}
L117:
	;
	if v414 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v409
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l1
	F_errmsg_plural(m, int32(_a_F_sendFile_16), int32(_a_F_sendFile_17), v409, v21+int32(16))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	F_pgstat_prepare_report_checksum_failure(m, l5)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L123
	}
L121:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1818), int32(_a_F_sendFile_4))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L3
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	F_pgstat_report_checksum_failures_in_db(m, l5, v409)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L3
	} else {
		goto L124
	}
L124:
	;
	goto L116
L125:
	;
	goto L8
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
	F_errmsg_internal(m, int32(_a_F_sendFile_18), v21)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1591), int32(_a_F_sendFile_4))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	F_errmsg_internal(m, int32(_a_F_sendFile_14), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1666), int32(_a_F_sendFile_4))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errmsg_internal(m, int32(_a_F_sendFile_14), int32(0))
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L3
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_sendFile_3), int32(1798), int32(_a_F_sendFile_4))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L3
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
