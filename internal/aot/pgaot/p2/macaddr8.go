package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_macaddr8_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v253 int32
	_ = v253
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = v23
	goto L2
L1:
	;
	m.G0 = v20 + int32(16)
	return v338
L2:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(base.Ui32(v41-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v41 == int32(32)) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v319 = int32(0)
	v320 = F_errsave_start(m, v22)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L67
	} else {
		goto L69
	}
L4:
	;
	v24 = v24 + int32(1)
	goto L2
L5:
	;
	if v41 == int32(0) {
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
	v53 = v24
	v55 = v41
	v56 = v2
	v57 = v2
	v58 = v2
	v59 = v2
	v63 = v2
	v64 = v2
	v65 = v2
	v66 = v2
	v67 = v2
	v69 = v2
	goto L9
L9:
	;
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v70 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	switch v264 - int32(6) {
	case 0:
		v284 = int32(254)
		v285 = int32(255)
		v286 = v267
		v287 = v268
		v288 = v269
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
	v264 = v56
	v267 = v57
	v268 = v58
	v269 = v59
	v273 = v63
	v274 = v64
	v275 = v65
	v276 = v66
	v277 = v67
	goto L11
L13:
	;
	goto L14
L14:
	;
	switch v56 {
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
	v204 = v53 + int32(2)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)))
	v207 = v205 - int32(45)
	if base.Ui32(int32(13)) < base.Ui32(v207) {
		v225 = v69
		v226 = v204
		goto L48
	} else {
		goto L49
	}
L16:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L45
	}
L17:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L42
	}
L18:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L39
	}
L19:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L36
	}
L20:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L33
	}
L21:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L30
	}
L22:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L27
	}
L23:
	;
	if base.I32_extend8_s(v55) < int32(0) {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v78 < int32(0) {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v82 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v82 < int32(0) {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	v195 = v57
	v196 = v58
	v197 = v59
	v198 = v82 + v78<<(uint(int32(4))%32)
	v199 = v64
	v200 = v65
	v201 = v66
	v202 = v67
	goto L15
L27:
	;
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v93 < int32(0) {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v97 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v97 < int32(0) {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	v195 = v57
	v196 = v58
	v197 = v59
	v198 = v63
	v199 = v97 + v93<<(uint(int32(4))%32)
	v200 = v65
	v201 = v66
	v202 = v67
	goto L15
L30:
	;
	v108 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v108 < int32(0) {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v112 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v112 < int32(0) {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	v195 = v57
	v196 = v58
	v197 = v59
	v198 = v63
	v199 = v64
	v200 = v112 + v108<<(uint(int32(4))%32)
	v201 = v66
	v202 = v67
	goto L15
L33:
	;
	v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v123 < int32(0) {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v127 < int32(0) {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v195 = v127 + v123<<(uint(int32(4))%32)
	v196 = v58
	v197 = v59
	v198 = v63
	v199 = v64
	v200 = v65
	v201 = v66
	v202 = v67
	goto L15
L36:
	;
	v138 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v138 < int32(0) {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v142 < int32(0) {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	v195 = v57
	v196 = v142 + v138<<(uint(int32(4))%32)
	v197 = v59
	v198 = v63
	v199 = v64
	v200 = v65
	v201 = v66
	v202 = v67
	goto L15
L39:
	;
	v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v153 < int32(0) {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v157 < int32(0) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v195 = v57
	v196 = v58
	v197 = v157 + v153<<(uint(int32(4))%32)
	v198 = v63
	v199 = v64
	v200 = v65
	v201 = v66
	v202 = v67
	goto L15
L42:
	;
	v168 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v168 < int32(0) {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v172 < int32(0) {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v195 = v57
	v196 = v58
	v197 = v59
	v198 = v63
	v199 = v64
	v200 = v65
	v201 = v172 + v168<<(uint(int32(4))%32)
	v202 = v67
	goto L15
L45:
	;
	v183 = int32(*(*int8)(unsafe.Add(mBase, uint32(v55&int32(255))+uint32(_c_F_macaddr8_in[0]))))
	if v70|v183 < int32(0) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v187 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_macaddr8_in[0]))))
	if v187 < int32(0) {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	v195 = v57
	v196 = v58
	v197 = v59
	v198 = v63
	v199 = v64
	v200 = v65
	v201 = v66
	v202 = v187 + v183<<(uint(int32(4))%32)
	goto L15
L48:
	;
	v228 = v56 + int32(1)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226))))
	if v56&int32(-3) == int32(5) {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	if int32(1)<<(uint(v207)%32)&int32(_a_F_macaddr8_in_0) == int32(0) {
		v225 = v69
		v226 = v204
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v217 = v69 & int32(255)
	if v217 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v225 = v221
	v226 = v53 + int32(3)
	goto L48
L52:
	;
	v221 = v205
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v205 != v217 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	v221 = v69
	goto L51
L56:
	;
	switch v229 {
	case 0:
		v264 = v228
		v267 = v195
		v268 = v196
		v269 = v197
		v273 = v198
		v274 = v199
		v275 = v200
		v276 = v201
		v277 = v202
		goto L11
	default:
		v53 = v226
		v55 = v229
		v56 = v228
		v57 = v195
		v58 = v196
		v59 = v197
		v63 = v198
		v64 = v199
		v65 = v200
		v66 = v201
		v67 = v202
		v69 = v225
		goto L9
	case 9, 10, 11, 12, 13, 32:
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v229 != 0 {
		v53 = v226
		v55 = v229
		v56 = v228
		v57 = v195
		v58 = v196
		v59 = v197
		v63 = v198
		v64 = v199
		v65 = v200
		v66 = v201
		v67 = v202
		v69 = v225
		goto L9
	} else {
		goto L64
	}
L59:
	;
	v234 = v226
	goto L60
L60:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+1)))
	if base.B2i32(base.Ui32(v253-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v253 == int32(32)) != 0 {
		v234 = v234 + int32(1)
		goto L60
	} else {
		goto L62
	}
L61:
	;
	if v253 == int32(0) {
		v264 = v228
		v267 = v195
		v268 = v196
		v269 = v197
		v273 = v198
		v274 = v199
		v275 = v200
		v276 = v201
		v277 = v202
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
	v264 = v228
	v267 = v195
	v268 = v196
	v269 = v197
	v273 = v198
	v274 = v199
	v275 = v200
	v276 = v201
	v277 = v202
	goto L11
L65:
	;
	v290 = F_palloc0(m, int32(8))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v284 = v268
	v285 = v267
	v286 = v269
	v287 = v276
	v288 = v277
	goto L65
L67:
	;
	return int32(0)
L68:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+7)) = uint8(v288)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+6)) = uint8(v287)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+5)) = uint8(v286)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+4)) = uint8(v284)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v285)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v275)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v274)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v273)
	v338 = v290
	goto L1
L69:
	;
	if v320 == int32(0) {
		v338 = v319
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L67
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(_a_F_macaddr8_in_1)
	F_errmsg(m, int32(_a_F_macaddr8_in_2), v20)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L67
	} else {
		goto L72
	}
L72:
	;
	F_errsave_finish(m, v22, int32(_a_F_macaddr8_in_3), int32(227), int32(_a_F_macaddr8_in_4))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L67
	} else {
		goto L73
	}
L73:
	;
	v338 = v319
	goto L1
}
