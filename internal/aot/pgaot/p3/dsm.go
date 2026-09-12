package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_create(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int64
	_ = v349
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v3
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[824])))
	if v22 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[824])) = uint8(v26)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_ResourceOwnerEnlarge(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v39 = F_MemoryContextAlloc(m, v37, int32(36))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[826]))
	if v42 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v45 = int32(4126868)
	*(*int32)(unsafe.Add(mBase, _consts[827])) = v45
	v49 = v45
	goto L12
L11:
	;
	v49 = v42
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(4126868)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v39
	*(*int32)(unsafe.Add(mBase, _consts[826])) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v39)+24)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v39)+16)) = int64(4294967295)
	v61 = *(*int32)(unsafe.Add(mBase, _consts[179]))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v61
	if v61 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_ResourceOwnerRemember(m, v61, v39, int32(1630476))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v67 = v39 + int32(28)
	v69 = v39 + int32(24)
	v71 = v39 + int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = int32(0)
	if v29 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L15
L17:
	;
	v184 = int32(0)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[828]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v187 != 0 {
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v79 = F_LWLockAcquire(m, v75+int32(4352), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v109 = v3
	goto L20
L20:
	;
	goto L27
L21:
	;
	v81 = int32(12)
	v87 = int32(base.Ui32(l0)>>(uint(v81)%32)) + base.B2i32(l0&int32(4095) != int32(0))
	v90 = F_FreePageManagerGet(m, v29, v87, v17+v81)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	if v90 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[825]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v95 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v93 + v94<<(uint(v95)%32)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v87 << (uint(v95) % 32)
	v177 = v87
	v178 = int32(1)
	goto L17
L24:
	;
	goto L25
L25:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v104+int32(4352))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v109 = v87
	goto L20
L27:
	;
	v126 = int32(4604072)
	v127 = int32(4604064)
	v128 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v130 = *(*int64)(unsafe.Add(mBase, _consts[818]))
	v131 = v128 ^ v130
	*(*int64)(unsafe.Add(mBase, _consts[818])) = base.I64_rotl(v131, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v131<<(uint(int64(16))%64) ^ base.I64_rotl(v128, int64(24)) ^ v131
	goto L29
L28:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v168 = F_LWLockAcquire(m, v164+int32(4352), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L33
	}
L29:
	;
	v153 = base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v128*int64(5), int64(7))*int64(9))>>(uint(int64(32))%64))) << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v153
	if v153 == int32(0) {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v159 = F_dsm_impl_op(m, int32(0), v153, l0, v71, v69, v67, int32(21))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v159 == int32(0) {
		goto L27
	} else {
		goto L32
	}
L32:
	;
	goto L28
L33:
	;
	v177 = v109
	v178 = v3
	goto L17
L34:
	;
	m.G0 = v17 + int32(16)
	return v430
L35:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v422+int32(4352))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L74
	}
L36:
	;
	v192 = v184
	goto L39
L37:
	;
	v278 = v184
	goto L38
L38:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v186)+8))
	if base.Ui32(v290) <= base.Ui32(v187) {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v205 = v192 * int32(24)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v186+int32(16)+v205)))
	if v207 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v278 = v187 * int32(24)
	goto L38
L41:
	;
	if v178 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v272 = v192 + int32(1)
	if v272 != v187 {
		v192 = v272
		goto L39
	} else {
		goto L48
	}
L44:
	;
	v214 = int32(4604072)
	v215 = int32(4604064)
	v216 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v218 = *(*int64)(unsafe.Add(mBase, _consts[818]))
	v219 = v216 ^ v218
	*(*int64)(unsafe.Add(mBase, _consts[818])) = base.I64_rotl(v219, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v219<<(uint(int64(16))%64) ^ base.I64_rotl(v216, int64(24)) ^ v219
	goto L47
L45:
	;
	v256 = v186
	goto L46
L46:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v256+v192*int32(24))+12)) = v261
	v263 = v256 + v205
	v264 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v263)+32)) = uint8(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v263)+28)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v263)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v192
	goto L35
L47:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[828]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v192<<(uint(int32(1))%32) | base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v216*int64(5), int64(7))*int64(9))>>(uint(int64(32))%64)))<<(uint(int32(32)-base.I32_clz(v243))%32) | int32(1)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v252 = v242 + v205
	*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v251
	v256 = v242
	goto L46
L48:
	;
	goto L40
L49:
	;
	if v178 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	goto L51
L51:
	;
	if v178 != 0 {
		goto L70
	} else {
		goto L71
	}
L52:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v313 != 0 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_FreePageManagerPut(m, v29, v292, v177)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v302+int32(4352))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L7
	} else {
		goto L58
	}
L56:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v296+int32(4352))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v311 = F_dsm_impl_op(m, int32(3), v308, int32(0), v71, v69, v67, int32(19))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	goto L52
L60:
	;
	F_ResourceOwnerForget(m, v313, v39, int32(1630476))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+4)) = v318
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = v320
	F_pfree(m, v39)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	if l1&int32(1) != 0 {
		v430 = int32(0)
		goto L34
	} else {
		goto L65
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(123175), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(498565), int32(626), int32(356457))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v347 = int32(4604072)
	v348 = int32(4604064)
	v349 = *(*int64)(unsafe.Add(mBase, _consts[819]))
	v351 = *(*int64)(unsafe.Add(mBase, _consts[818]))
	v352 = v349 ^ v351
	*(*int64)(unsafe.Add(mBase, _consts[818])) = base.I64_rotl(v352, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[819])) = v352<<(uint(int64(16))%64) ^ base.I64_rotl(v349, int64(24)) ^ v352
	goto L73
L71:
	;
	v389 = v186
	goto L72
L72:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v394 = v389 + v187*int32(24)
	v395 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v394)+32)) = uint8(v395)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+28)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v394)+16)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v394)+12)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v187
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v403 + int32(1)
	goto L35
L73:
	;
	v375 = *(*int32)(unsafe.Add(mBase, _consts[828]))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v375)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v187<<(uint(int32(1))%32) | base.I32_wrap_i64(int64(base.Ui64(base.I64_rotl(v349*int64(5), int64(7))*int64(9))>>(uint(int64(32))%64)))<<(uint(int32(32)-base.I32_clz(v376))%32) | int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v385 = v278 + v375
	*(*int32)(unsafe.Add(mBase, uint32(v385)+24)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v385)+20)) = v384
	v389 = v375
	goto L72
L74:
	;
	v430 = v39
	goto L34
}
