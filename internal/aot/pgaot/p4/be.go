package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_be_lo_export(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	var v219 int32
	_ = v219
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v298 int32
	_ = v298
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int64
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v23 = v2
	v24 = v2
	v25 = v2
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = int32(-1)
	v32 = v16
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L4
L3:
	;
	m.G0 = v16 + int32(80)
	return int32(1)
L4:
	;
	if v31 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v375 = int32(m.ExcTag)
	v376 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v375 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L7:
	;
	v37 = v32 + int32(-8192)
	m.G0 = v37
	v40 = v37 - int32(1024)
	m.G0 = v40
	v43 = v40 - int32(160)
	m.G0 = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(28))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v37
	v54 = F_pg_detoast_datum_packed(m, v45)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		v373 = v43
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v105 = v23
	v106 = v24
	v107 = v25
	v108 = v26
	v109 = v27
	v110 = v28
	v111 = v29
	v112 = v30
	v114 = v32
	goto L9
L9:
	;
	if v112 != 0 {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v57 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v37
	v68 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v69 = F_inv_open(m, v46, int32(262144), v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		v373 = v43
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v37
	F_text_to_cstring_buffer(m, v54, v40, int32(1024))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		v373 = v43
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v37
	v90 = int32(4452996)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = int32(18)
	v94 = F___syscall_ret(m, v91)
	mBase = m.M
	goto L13
L13:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	v99 = *(*int32)(unsafe.Add(mBase, _consts[425]))
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v16 + int32(48)
	goto L17
L15:
	;
	v105 = v37
	v106 = v40
	v107 = v99
	v108 = v97
	v109 = v43
	v110 = v94
	v111 = v69
	v112 = int32(0)
	v114 = v43
	goto L9
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v108
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v127 = int32(4452996)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v110
	v131 = F___syscall_ret(m, v128)
	mBase = m.M
	goto L21
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v152 = F_OpenTransientFilePerm(m, v106, int32(577), int32(420))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_pg_re_throw(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L1
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v108
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v166 = int32(4452996)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	*(*int32)(unsafe.Add(mBase, _consts[517])) = v110
	v170 = F___syscall_ret(m, v167)
	mBase = m.M
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v108
	*(*int32)(unsafe.Add(mBase, _consts[425])) = v107
	if v152 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L33
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errcode_for_file_access(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v106
	F_errmsg(m, int32(313331), v16)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errfinish(m, int32(520635), int32(527), int32(84976))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L1
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v306 = F_CloseTransientFile(m, v152)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L42
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v241 = F_inv_read(m, v111, v105, int32(8192))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L35
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L38
	}
L35:
	;
	if v241 <= int32(0) {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	v252 = F_write(m, v152, v105, v241)
	mBase = m.M
	if v252 == v241 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errcode_for_file_access(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v106
	F_errmsg(m, int32(313294), v16+int32(16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errfinish(m, int32(520635), int32(539), int32(84976))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L41
	}
L41:
	;
	goto L1
L42:
	;
	if v306 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_pfree(m, v111)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L50
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errcode_for_file_access(m)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v106
	F_errmsg(m, int32(314465), v16+int32(32))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v107
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v105
	F_errfinish(m, int32(520635), int32(546), int32(84976))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		v373 = v114
		goto L6
	} else {
		goto L49
	}
L49:
	;
	goto L1
L50:
	;
	goto L5
L51:
	;
	v380 = int32(v376)
	m.G0 = v373
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v380)+4))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v16+int32(48) == v387 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	m.ExcPending = 1
	goto L60
L53:
	;
	if v390 != 0 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v383)+4))
	v390 = v389
	goto L56
L55:
	;
	v390 = int32(0)
	goto L56
L56:
	;
	goto L53
L57:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v23 = v391
	v24 = v392
	v25 = v396
	v26 = v397
	v27 = v393
	v28 = v395
	v29 = v394
	v30 = v382
	v31 = v390
	v32 = v373
	goto L2
L58:
	;
	goto L59
L59:
	;
	F___wasm_longjmp(m, v383, v382)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	return int32(0)
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_be_lo_open(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v8&int32(131072) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_PreventCommandIfReadOnly(m, int32(716333))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[513])) = uint8(v17)
	v20 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[36]))
	v30 = F_AllocSetContextCreateInternal(m, v25, int32(304086), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	v33 = v20
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[515]))
	if int32(0) < v35 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _consts[514])) = v30
	v33 = v30
	goto L8
L10:
	;
	v88 = F_inv_open(m, v7, v8, v85)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[515])) = v69
	*(*int32)(unsafe.Add(mBase, _consts[516])) = v75
	v81 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v83 = v70
	v85 = v81
	goto L10
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	v41 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v67 = F_MemoryContextAllocZero(m, v33, int32(256))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L22
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40+v41<<(uint(int32(2))%32))))
	if v50 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v62 = F_repalloc0(m, v40, v35<<(uint(int32(2))%32), v35<<(uint(int32(3))%32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L21
	}
L17:
	;
	v83 = v41
	v85 = v33
	goto L10
L18:
	;
	goto L19
L19:
	;
	v54 = v41 + int32(1)
	if v54 != v35 {
		v41 = v54
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v69 = v35 << (uint(int32(1)) % 32)
	v70 = v35
	v75 = v62
	goto L11
L22:
	;
	v69 = int32(64)
	v70 = int32(0)
	v75 = v67
	goto L11
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v94 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[214]))
	v97 = F_RegisterSnapshotOnOwner(m, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[516]))
	*(*int32)(unsafe.Add(mBase, uint32(v101+v83<<(uint(int32(2))%32)))) = v88
	return v83
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v97
	goto L27
}
func F_be_lowrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_PreventCommandIfReadOnly(m, int32(718446))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v18 = int32(4)
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v20&int32(254) == int32(2) {
					v29 = v18
				} else {
					v29 = base.B2i32(v20 == int32(18)) << (uint(v18) % 32)
				}
				if v20 == int32(1) {
					v32 = v18
				} else {
					v32 = v29
				}
				v45 = v32
			} else {
				v33 = int32(1)
				if v15&v33 != 0 {
					v45 = int32(base.Ui32(v15)>>(uint(v33)%32)) - v33
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v46 = int32(1)
			if v15&v46 != 0 {
				v50 = v46
			} else {
				v50 = int32(4)
			}
			v52 = m.G0
			v54 = v52 - int32(32)
			m.G0 = v54
			if v6 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v6
						F_errmsg(m, int32(506257), v54)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(520635), int32(190), int32(368416))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, _consts[515]))
				if v59 <= v6 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v54))) = v6
							F_errmsg(m, int32(506257), v54)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(520635), int32(190), int32(368416))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, _consts[516]))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v6<<(uint(int32(2))%32))))
					if v66 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v86 = m.ExcPending
							if v86 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v6
								F_errmsg(m, int32(506257), v54)
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520635), int32(190), int32(368416))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+24)))
						if v69&int32(2) == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v54)+16)) = v6
									F_errmsg(m, int32(346902), v54+int32(16))
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520635), int32(198), int32(368416))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v74 = F_inv_write(m, v66, v8+v50, v45)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								m.G0 = v54 + int32(32)
								return v74
							}
						}
					}
				}
			}
		}
	}
}
