package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enable_timeout(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var __phi102 int32
	_ = __phi102
	var v103 int32
	_ = v103
	var __phi103 int32
	_ = __phi103
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var __phi224 int32
	_ = __phi224
	var v227 int32
	_ = v227
	var __phi227 int32
	_ = __phi227
	var v229 int32
	_ = v229
	var __phi229 int32
	_ = __phi229
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = l0 * int32(40)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[844]))))
	if v21 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v26 <= int32(0) {
		v69 = int32(-1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v157 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v159 <= v157 {
		v195 = v157
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1223])))
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v90)+4)) = uint8(v91)
	v94 = v33 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v94 < v96 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v33 = int32(0)
	goto L8
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v33 < v55 {
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v42 = v33 << (uint(int32(2)) % 32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+uint32(_consts[1223])))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v69 = int32(-1)
	goto L5
L10:
	;
	v49 = v33 + int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v49 < v51 {
		v33 = v49
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v69 = v33
	goto L5
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v69
	v76 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v76 - int32(1)
	F_errmsg_internal(m, int32(466854), v15+int32(16))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(492709), int32(143), int32(28113))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	__phi102 = v33
	__phi103 = v94
	v102 = __phi102
	v103 = __phi103
	goto L20
L18:
	;
	goto L19
L19:
	;
	v137 = int32(4514424)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	*(*int32)(unsafe.Add(mBase, _consts[843])) = v139 - int32(1)
	goto L3
L20:
	;
	v110 = int32(2)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v103<<(uint(v110)%32))+uint32(_consts[1223])))
	*(*int32)(unsafe.Add(mBase, uint32(v102<<(uint(v110)%32))+uint32(_consts[1223]))) = v118
	v121 = v103 + int32(1)
	v123 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v121 < v123 {
		__phi102 = v103
		__phi103 = v121
		v102 = __phi102
		v103 = __phi103
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1224]))) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1225]))) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[1226]))) = l1
	v204 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[845]))) = uint8(v204)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v195 <= v207 {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v168 = v157
	goto L25
L25:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v168<<(uint(int32(2))%32))+uint32(_consts[1223])))
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v178)+24))
	if l2 < v179 {
		v195 = v168
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v195 = v185
	goto L23
L27:
	;
	if l2 == v179 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if base.Ui32(l0) < base.Ui32(v182) {
		v195 = v168
		goto L23
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v185 = v168 + int32(1)
	v187 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	if v185 < v187 {
		v168 = v185
		goto L25
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L26
L33:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[844]))) = uint8(v209)
	v213 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	v215 = v213 - v209
	if v215 < v195 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L49
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v195<<(uint(int32(2))%32))+uint32(_consts[1223]))) = v18 + int32(4514432)
	v323 = int32(4514424)
	v325 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	*(*int32)(unsafe.Add(mBase, _consts[843])) = v325 + int32(1)
	m.G0 = v15 + int32(32)
	return
L37:
	;
	v217 = v213 - v195
	v221 = v217 & int32(3)
	if v221 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if base.Ui32(v217-int32(1)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L45
	}
L39:
	;
	v255 = v213
	v256 = v215
	goto L38
L40:
	;
	goto L41
L41:
	;
	__phi224 = v213
	__phi227 = int32(0)
	__phi229 = v215
	v224 = __phi224
	v227 = __phi227
	v229 = __phi229
	goto L42
L42:
	;
	v236 = int32(2)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v229<<(uint(v236)%32))+uint32(_consts[1223])))
	*(*int32)(unsafe.Add(mBase, uint32(v224<<(uint(v236)%32))+uint32(_consts[1223]))) = v244
	v246 = int32(1)
	v247 = v229 - v246
	v249 = v227 + v246
	if v249 != v221 {
		__phi224 = v229
		__phi227 = v249
		__phi229 = v247
		v224 = __phi224
		v227 = __phi227
		v229 = __phi229
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v255 = v229
	v256 = v247
	goto L38
L44:
	;
	goto L43
L45:
	;
	v269 = v255
	v270 = v256
	goto L46
L46:
	;
	v277 = int32(2)
	v282 = v270 << (uint(v277) % 32)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1223])))
	*(*int32)(unsafe.Add(mBase, uint32(v269<<(uint(v277)%32))+uint32(_consts[1223]))) = v285
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1227])))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1223]))) = v289
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1228])))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1227]))) = v293
	v296 = v270 - int32(3)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v296<<(uint(v277)%32))+uint32(_consts[1223])))
	*(*int32)(unsafe.Add(mBase, uint32(v282)+uint32(_consts[1228]))) = v301
	if v195 < v296 {
		v269 = v296
		v270 = v270 - int32(4)
		goto L46
	} else {
		goto L48
	}
L47:
	;
	goto L36
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v195
	v338 = *(*int32)(unsafe.Add(mBase, _consts[843]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v338
	F_errmsg_internal(m, int32(466854), v15)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(492709), int32(120), int32(65450))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_timeout_active(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0*int32(40))+uint32(_consts[844]))))
	return v6
}
