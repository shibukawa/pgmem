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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v315 int32
	_ = v315
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
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
	v453 = F_palloc(m, int32(16))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
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
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387)+4)))
	if v388&int32(4) != 0 {
		goto L58
	} else {
		goto L59
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v48 << (uint(int32(2)) % 32)
	v57 = v49 + int32(8)
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	base.MemoryFill(m, v57, int32(0), v43)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v61 = F_hstoreUpgrade(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v65 = v63 & int32(268435455)
	if v65 == int32(0) {
		v443 = v49
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v69 = v61 + int32(8)
	v70 = int32(3)
	v72 = v69 + v65<<(uint(v70)%32)
	v74 = v43 << (uint(v70) % 32)
	v86 = v2
	goto L22
L22:
	;
	v91 = v69 + v86<<(uint(int32(3))%32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	if v92 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v443 = v49
	goto L12
L24:
	;
	if v107 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v107 = v92 & int32(1073741823)
	v108 = v72
	goto L24
L26:
	;
	goto L27
L27:
	;
	v97 = int32(1073741823)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v91-int32(4))))
	v103 = v101 & v97
	v107 = v92&v97 - v103
	v108 = v103 + v72
	goto L24
L28:
	;
	v109 = int32(-1)
	if v107 != int32(1) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v219 = int32(0)
	goto L30
L30:
	;
	v220 = base.I32_rem_u_s(v219, v74)
	v223 = v57 + int32(base.Ui32(v220)>>(uint(int32(3))%32))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v229 = v224 | int32(1)<<(uint(v220&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v223))) = uint8(v229)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	if v231&int32(1073741824) == int32(0) {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v219 = v189 ^ int32(-1)
	goto L30
L32:
	;
	v117 = v108
	v118 = v109
	v119 = int32(0)
	goto L35
L33:
	;
	v162 = v108
	v163 = v109
	goto L34
L34:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32((v163^v176)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_compress[0])))
	v189 = v184 ^ int32(base.Ui32(v163)>>(uint(int32(8))%32))
	goto L31
L35:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v133 = int32(255)
	v135 = int32(2)
	v139 = *(*int32)(unsafe.Add(mBase, uint32((v131^v118)&v133<<(uint(v135)%32))+uint32(_c_F_ghstore_compress[0])))
	v140 = int32(8)
	v142 = v139 ^ int32(base.Ui32(v118)>>(uint(v140)%32))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32((v142^v143)&v133<<(uint(v135)%32))+uint32(_c_F_ghstore_compress[0])))
	v154 = v151 ^ int32(base.Ui32(v142)>>(uint(v140)%32))
	v156 = v117 + v135
	v158 = v119 + v135
	if v158 != v107&int32(-2) {
		v117 = v156
		v118 = v154
		v119 = v158
		goto L35
	} else {
		goto L37
	}
L36:
	;
	if v107&int32(1) == int32(0) {
		v189 = v154
		goto L31
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v162 = v156
	v163 = v154
	goto L34
L39:
	;
	if v231 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	goto L41
L41:
	;
	v385 = v86 + int32(1)
	if v385 != v65 {
		v86 = v385
		goto L22
	} else {
		goto L57
	}
L42:
	;
	if v246 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v246 = v231 & int32(1073741823)
	v247 = v72
	goto L42
L44:
	;
	goto L45
L45:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v242 = v240 & int32(1073741823)
	v246 = v231 - v242
	v247 = v242 + v72
	goto L42
L46:
	;
	v248 = int32(-1)
	if v246 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v358 = int32(0)
	goto L48
L48:
	;
	v359 = base.I32_rem_u_s(v358, v74)
	v362 = v57 + int32(base.Ui32(v359)>>(uint(int32(3))%32))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	v368 = v363 | int32(1)<<(uint(v359&int32(7))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v362))) = uint8(v368)
	goto L41
L49:
	;
	v358 = v328 ^ int32(-1)
	goto L48
L50:
	;
	v256 = v247
	v257 = v248
	v258 = int32(0)
	goto L53
L51:
	;
	v301 = v247
	v302 = v248
	goto L52
L52:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301))))
	v323 = *(*int32)(unsafe.Add(mBase, uint32((v302^v315)&int32(255)<<(uint(int32(2))%32))+uint32(_c_F_ghstore_compress[0])))
	v328 = v323 ^ int32(base.Ui32(v302)>>(uint(int32(8))%32))
	goto L49
L53:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256))))
	v272 = int32(255)
	v274 = int32(2)
	v278 = *(*int32)(unsafe.Add(mBase, uint32((v270^v257)&v272<<(uint(v274)%32))+uint32(_c_F_ghstore_compress[0])))
	v279 = int32(8)
	v281 = v278 ^ int32(base.Ui32(v257)>>(uint(v279)%32))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+1)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32((v281^v282)&v272<<(uint(v274)%32))+uint32(_c_F_ghstore_compress[0])))
	v293 = v290 ^ int32(base.Ui32(v281)>>(uint(v279)%32))
	v295 = v256 + v274
	v297 = v258 + v274
	if v297 != v246&int32(-2) {
		v256 = v295
		v257 = v293
		v258 = v297
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v246&int32(1) == int32(0) {
		v328 = v293
		goto L49
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v301 = v295
	v302 = v293
	goto L52
L57:
	;
	goto L23
L58:
	;
	return v15
L59:
	;
	goto L60
L60:
	;
	v392 = int32(0)
	if v43 <= v392 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v434 = F_palloc(m, int32(8))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L10
	} else {
		goto L69
	}
L62:
	;
	v397 = v392
	goto L63
L63:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+(v387+int32(8))))))
	if v412 == int32(255) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	return v15
L65:
	;
	v416 = v397 + int32(1)
	if v43 != v416 {
		v397 = v416
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
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = int64(17179869216)
	v443 = v434
	goto L12
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = v443
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+4)) = v456
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v453)+8)) = v458
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+12)))
	v461 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v453)+14)) = uint8(v461)
	*(*uint16)(unsafe.Add(mBase, uint32(v453)+12)) = uint16(v460)
	return v453
}
func F_ghstore_options(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(8)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(0)
	F_add_local_int_reloption(m, v2, int32(_a_F_ghstore_options_0), int32(_a_F_ghstore_options_1), int32(16), int32(1), int32(2024))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
