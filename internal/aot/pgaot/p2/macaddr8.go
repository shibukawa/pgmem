package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v361 int64
	_ = v361
	v2 = int32(0)
	v18 = int64(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v25 = v24
	goto L2
L1:
	;
	m.G0 = v21 + int32(16)
	return v361
L2:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if base.B2i32(base.Ui32(v43-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v43 == int32(32)) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v326 = F_errsave_start(m, v23)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L67
	} else {
		goto L69
	}
L4:
	;
	v25 = v25 + int32(1)
	goto L2
L5:
	;
	if v43 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L3
L7:
	;
	goto L6
L8:
	;
	v55 = v25
	v57 = v43
	v58 = v2
	v59 = v2
	v60 = v2
	v61 = v2
	v65 = v2
	v66 = v2
	v67 = v2
	v68 = v2
	v69 = v2
	v71 = v2
	goto L9
L9:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v73 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	switch v268 - int32(6) {
	case 0:
		v289 = int32(254)
		v290 = int32(255)
		v291 = v271
		v292 = v272
		v293 = v273
		goto L65
	default:
		goto L7
	case 2:
		goto L66
	}
L11:
	;
	goto L10
L12:
	;
	v268 = v58
	v271 = v59
	v272 = v60
	v273 = v61
	v277 = v65
	v278 = v66
	v279 = v67
	v280 = v68
	v281 = v69
	goto L11
L13:
	;
	goto L14
L14:
	;
	switch v58 {
	case 0:
		goto L23
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	case 4:
		goto L19
	case 5:
		goto L18
	case 6:
		goto L17
	case 7:
		goto L16
	default:
		goto L7
	}
L15:
	;
	v207 = v55 + int32(2)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+2)))
	v210 = v208 - int32(45)
	if base.Ui32(int32(13)) < base.Ui32(v210) {
		v228 = v71
		v229 = v207
		goto L48
	} else {
		goto L49
	}
L16:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L45
	}
L17:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L42
	}
L18:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L39
	}
L19:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L36
	}
L20:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L33
	}
L21:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L30
	}
L22:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L27
	}
L23:
	;
	if base.I32_extend8_s(v57) < int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v81 < int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v85 < int32(0) {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v198 = v59
	v199 = v60
	v200 = v61
	v201 = v85 + v81<<(uint(int32(4))%32)
	v202 = v66
	v203 = v67
	v204 = v68
	v205 = v69
	goto L15
L27:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v96 < int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v100 < int32(0) {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v198 = v59
	v199 = v60
	v200 = v61
	v201 = v65
	v202 = v100 + v96<<(uint(int32(4))%32)
	v203 = v67
	v204 = v68
	v205 = v69
	goto L15
L30:
	;
	v111 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v111 < int32(0) {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v115 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v115 < int32(0) {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v198 = v59
	v199 = v60
	v200 = v61
	v201 = v65
	v202 = v66
	v203 = v115 + v111<<(uint(int32(4))%32)
	v204 = v68
	v205 = v69
	goto L15
L33:
	;
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v126 < int32(0) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v130 < int32(0) {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v198 = v130 + v126<<(uint(int32(4))%32)
	v199 = v60
	v200 = v61
	v201 = v65
	v202 = v66
	v203 = v67
	v204 = v68
	v205 = v69
	goto L15
L36:
	;
	v141 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v141 < int32(0) {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v145 < int32(0) {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v198 = v59
	v199 = v145 + v141<<(uint(int32(4))%32)
	v200 = v61
	v201 = v65
	v202 = v66
	v203 = v67
	v204 = v68
	v205 = v69
	goto L15
L39:
	;
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v156 < int32(0) {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v160 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v160 < int32(0) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v198 = v59
	v199 = v60
	v200 = v160 + v156<<(uint(int32(4))%32)
	v201 = v65
	v202 = v66
	v203 = v67
	v204 = v68
	v205 = v69
	goto L15
L42:
	;
	v171 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v171 < int32(0) {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v175 < int32(0) {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v198 = v59
	v199 = v60
	v200 = v61
	v201 = v65
	v202 = v66
	v203 = v67
	v204 = v175 + v171<<(uint(int32(4))%32)
	v205 = v69
	goto L15
L45:
	;
	v186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v57&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v73|v186 < int32(0) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v190 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+uint32(_c_F_macaddr8_in[0]))))
	if v190 < int32(0) {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v198 = v59
	v199 = v60
	v200 = v61
	v201 = v65
	v202 = v66
	v203 = v67
	v204 = v68
	v205 = v190 + v186<<(uint(int32(4))%32)
	goto L15
L48:
	;
	v231 = v58 + int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	if v58&int32(-3) == int32(5) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	if int32(1)<<(uint(v210)%32)&int32(_a_F_macaddr8_in_0) == int32(0) {
		v228 = v71
		v229 = v207
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v220 = v71 & int32(255)
	if v220 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v228 = v224
	v229 = v55 + int32(3)
	goto L48
L52:
	;
	v224 = v208
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v208 != v220 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v224 = v71
	goto L51
L56:
	;
	switch v232 {
	case 0:
		v268 = v231
		v271 = v198
		v272 = v199
		v273 = v200
		v277 = v201
		v278 = v202
		v279 = v203
		v280 = v204
		v281 = v205
		goto L11
	default:
		v55 = v229
		v57 = v232
		v58 = v231
		v59 = v198
		v60 = v199
		v61 = v200
		v65 = v201
		v66 = v202
		v67 = v203
		v68 = v204
		v69 = v205
		v71 = v228
		goto L9
	case 9, 10, 11, 12, 13, 32:
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v232 != 0 {
		v55 = v229
		v57 = v232
		v58 = v231
		v59 = v198
		v60 = v199
		v61 = v200
		v65 = v201
		v66 = v202
		v67 = v203
		v68 = v204
		v69 = v205
		v71 = v228
		goto L9
	} else {
		goto L64
	}
L59:
	;
	v237 = v229
	goto L60
L60:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)))
	if base.B2i32(base.Ui32(v257-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v257 == int32(32)) != 0 {
		v237 = v237 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if v257 == int32(0) {
		v268 = v231
		v271 = v198
		v272 = v199
		v273 = v200
		v277 = v201
		v278 = v202
		v279 = v203
		v280 = v204
		v281 = v205
		goto L11
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	goto L7
L64:
	;
	v268 = v231
	v271 = v198
	v272 = v199
	v273 = v200
	v277 = v201
	v278 = v202
	v279 = v203
	v280 = v204
	v281 = v205
	goto L11
L65:
	;
	v295 = F_palloc0(m, int32(8))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v289 = v272
	v290 = v271
	v291 = v273
	v292 = v280
	v293 = v281
	goto L65
L67:
	;
	return int64(0)
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+7)) = uint8(v293)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+6)) = uint8(v292)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+5)) = uint8(v291)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+4)) = uint8(v289)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+3)) = uint8(v290)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+2)) = uint8(v279)
	*(*uint8)(unsafe.Add(mBase, uint32(v295)+1)) = uint8(v278)
	*(*uint8)(unsafe.Add(mBase, uint32(v295))) = uint8(v277)
	v361 = base.I64_extend_i32_u(v295)
	goto L1
L69:
	;
	if v326 == int32(0) {
		v361 = v18
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(_a_F_macaddr8_in_1)
	F_errmsg(m, int32(_a_F_macaddr8_in_2), v21)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	F_errsave_finish(m, v23, int32(_a_F_macaddr8_in_3), int32(227), int32(_a_F_macaddr8_in_4))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v361 = v18
	goto L1
}
