package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v419 int32
	_ = v419
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	v2 = int32(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 == v2 {
		v34 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v34&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	if v21 == int32(0) {
		v34 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v24 != int32(7) {
		v34 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v27 != int32(17) {
		v34 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	v34 = v30 ^ int32(1)
	goto L2
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_get_fn_opclass_options(m, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v43 = int32(16)
	goto L9
L9:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+14)))
	if v44 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v43 = v42
	goto L9
L12:
	;
	v429 = F_palloc(m, int32(16))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L70
	}
L13:
	;
	v48 = v43 + int32(8)
	v49 = F_palloc(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363)+4)))
	if v364&int32(4) != 0 {
		goto L58
	} else {
		goto L59
	}
L16:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v48 << (uint(int32(2)) % 32)
	v60 = F__emscripten_memset_bulkmem(m, v49+int32(8), base.I32_extend8_s(v51), v43)
	mBase = m.M
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v62 = F_hstoreUpgrade(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v66 = v64 & int32(268435455)
	if v66 == int32(0) {
		v419 = v49
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v70 = v62 + int32(8)
	v71 = int32(3)
	v73 = v70 + v66<<(uint(v71)%32)
	v75 = v43 << (uint(v71) % 32)
	v80 = v2
	goto L20
L20:
	;
	v92 = v70 + v80<<(uint(int32(3))%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v419 = v49
	goto L12
L22:
	;
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v108 = v93 & int32(1073741823)
	v109 = v73
	goto L22
L24:
	;
	goto L25
L25:
	;
	v98 = int32(1073741823)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92-int32(4))))
	v104 = v102 & v98
	v108 = v93&v98 - v104
	v109 = v104 + v73
	goto L22
L26:
	;
	v110 = int32(1)
	v112 = int32(-1)
	if v108 != v110 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v204 = int32(0)
	goto L28
L28:
	;
	v205 = base.I32_rem_u_s(v204, v75)
	v206 = int32(3)
	v208 = v60 + int32(base.Ui32(v205)>>(uint(v206)%32))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	v214 = v209 | int32(1)<<(uint(v205&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v208))) = uint8(v214)
	v220 = v70 + v80<<(uint(v206)%32) + int32(4)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v221&int32(1073741824) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	v118 = v109
	v119 = v112
	v120 = int32(0)
	goto L32
L30:
	;
	v160 = v109
	v161 = v112
	goto L31
L31:
	;
	if v108&v110 != 0 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	v135 = int32(255)
	v137 = int32(2)
	v140 = *(*int32)(unsafe.Add(mBase, uint32((v133^v119)&v135<<(uint(v137)%32))+uint32(_consts[1056])))
	v141 = int32(8)
	v143 = v140 ^ int32(base.Ui32(v119)>>(uint(v141)%32))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32((v143^v144)&v135<<(uint(v137)%32))+uint32(_consts[1056])))
	v154 = v151 ^ int32(base.Ui32(v143)>>(uint(v141)%32))
	v156 = v118 + v137
	v158 = v120 + v137
	if v158 != v108&int32(-2) {
		v118 = v156
		v119 = v154
		v120 = v158
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v160 = v156
	v161 = v154
	goto L31
L34:
	;
	goto L33
L35:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v182 = *(*int32)(unsafe.Add(mBase, uint32((v161^v174)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1056])))
	v186 = v182 ^ int32(base.Ui32(v161)>>(uint(int32(8))%32))
	goto L37
L36:
	;
	v186 = v161
	goto L37
L37:
	;
	v204 = v186 ^ int32(-1)
	goto L28
L38:
	;
	if v221 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L40
L40:
	;
	v361 = v80 + int32(1)
	if v361 != v66 {
		v80 = v361
		goto L20
	} else {
		goto L57
	}
L41:
	;
	if v238 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v238 = v221 & int32(1073741823)
	v239 = v73
	goto L41
L43:
	;
	goto L44
L44:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220-int32(4))))
	v234 = v232 & int32(1073741823)
	v238 = v221 - v234
	v239 = v234 + v73
	goto L41
L45:
	;
	v240 = int32(1)
	v242 = int32(-1)
	if v238 != v240 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v334 = int32(0)
	goto L47
L47:
	;
	v335 = base.I32_rem_u_s(v334, v75)
	v338 = v60 + int32(base.Ui32(v335)>>(uint(int32(3))%32))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	v344 = v339 | int32(1)<<(uint(v335&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v338))) = uint8(v344)
	goto L40
L48:
	;
	v248 = v239
	v249 = v242
	v250 = int32(0)
	goto L51
L49:
	;
	v290 = v239
	v291 = v242
	goto L50
L50:
	;
	if v238&v240 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v265 = int32(255)
	v267 = int32(2)
	v270 = *(*int32)(unsafe.Add(mBase, uint32((v263^v249)&v265<<(uint(v267)%32))+uint32(_consts[1056])))
	v271 = int32(8)
	v273 = v270 ^ int32(base.Ui32(v249)>>(uint(v271)%32))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+1)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32((v273^v274)&v265<<(uint(v267)%32))+uint32(_consts[1056])))
	v284 = v281 ^ int32(base.Ui32(v273)>>(uint(v271)%32))
	v286 = v248 + v267
	v288 = v250 + v267
	if v288 != v238&int32(-2) {
		v248 = v286
		v249 = v284
		v250 = v288
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v290 = v286
	v291 = v284
	goto L50
L53:
	;
	goto L52
L54:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290))))
	v312 = *(*int32)(unsafe.Add(mBase, uint32((v291^v304)&int32(255)<<(uint(int32(2))%32))+uint32(_consts[1056])))
	v316 = v312 ^ int32(base.Ui32(v291)>>(uint(int32(8))%32))
	goto L56
L55:
	;
	v316 = v291
	goto L56
L56:
	;
	v334 = v316 ^ int32(-1)
	goto L47
L57:
	;
	goto L21
L58:
	;
	return v15
L59:
	;
	goto L60
L60:
	;
	v368 = int32(0)
	if v43 <= v368 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v410 = F_palloc(m, int32(8))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L69
	}
L62:
	;
	v373 = v368
	goto L63
L63:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+(v363+int32(8))))))
	if v388 == int32(255) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	return v15
L65:
	;
	v392 = v373 + int32(1)
	if v43 != v392 {
		v373 = v392
		goto L63
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	goto L61
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v410))) = int64(17179869216)
	v419 = v410
	goto L12
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429))) = v419
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v429)+4)) = v432
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v429)+8)) = v434
	v436 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	v437 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v429)+14)) = uint8(v437)
	*(*uint16)(unsafe.Add(mBase, uint32(v429)+12)) = uint16(v436)
	return v429
}
func F_ghstore_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(282166), int32(159271), int32(16), int32(1), int32(2024))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
