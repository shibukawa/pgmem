package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FigureColnameInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v366 int32
	_ = v366
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	if l0 == v3 {
		v366 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v366
L2:
	;
	v21 = l0
	goto L25
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v348
	v366 = v352
	goto L1
L4:
	;
	v348 = int32(_a_F_FigureColnameInternal_0)
	v352 = v35
	goto L3
L5:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v312<<(uint(int32(2))%32))+uint32(_c_F_FigureColnameInternal[0])))
	v348 = v344
	v352 = v35
	goto L3
L6:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(int32(7)) <= base.Ui32(v336) {
		v366 = v3
		goto L1
	} else {
		goto L93
	}
L7:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(int32(15)) <= base.Ui32(v330) {
		v366 = v3
		goto L1
	} else {
		goto L92
	}
L8:
	;
	v348 = int32(_a_F_FigureColnameInternal_1)
	v352 = v35
	goto L3
L9:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if base.Ui32(v312) < base.Ui32(int32(3)) {
		goto L5
	} else {
		goto L88
	}
L10:
	;
	v348 = int32(_a_F_FigureColnameInternal_2)
	v352 = v35
	goto L3
L11:
	;
	v348 = int32(_a_F_FigureColnameInternal_3)
	v352 = v35
	goto L3
L12:
	;
	v348 = int32(_a_F_FigureColnameInternal_4)
	v352 = v35
	goto L3
L13:
	;
	v348 = int32(_a_F_FigureColnameInternal_5)
	v352 = v35
	goto L3
L14:
	;
	v348 = int32(_a_F_FigureColnameInternal_6)
	v352 = v35
	goto L3
L15:
	;
	v348 = int32(_a_F_FigureColnameInternal_7)
	v352 = v35
	goto L3
L16:
	;
	v348 = int32(_a_F_FigureColnameInternal_8)
	v352 = v35
	goto L3
L17:
	;
	v348 = int32(_a_F_FigureColnameInternal_9)
	v352 = v35
	goto L3
L18:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	switch v302 {
	case 0:
		v348 = int32(_a_F_FigureColnameInternal_10)
		v352 = v35
		goto L3
	case 1:
		goto L87
	default:
		v366 = v3
		goto L1
	}
L19:
	;
	v348 = int32(_a_F_FigureColnameInternal_11)
	v352 = v35
	goto L3
L20:
	;
	v348 = int32(_a_F_FigureColnameInternal_12)
	v352 = v35
	goto L3
L21:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v293 = F_FigureColnameInternal(m, v292, l1)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L77
	} else {
		goto L85
	}
L22:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	switch v281 {
	case 0:
		v348 = int32(_a_F_FigureColnameInternal_13)
		v352 = v35
		goto L3
	default:
		v366 = v3
		goto L1
	case 4:
		goto L81
	case 6:
		goto L82
	}
L23:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v260 = F_FigureColnameInternal(m, v259, l1)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L77
	} else {
		goto L78
	}
L24:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v254 != int32(5) {
		v366 = v3
		goto L1
	} else {
		goto L76
	}
L25:
	;
	v35 = int32(2)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	switch v37 - int32(10) {
	case 0:
		goto L8
	default:
		v366 = v3
		goto L1
	case 3:
		goto L4
	case 12:
		goto L22
	case 22:
		goto L21
	case 26:
		goto L20
	case 28:
		goto L19
	case 29:
		goto L18
	case 30:
		goto L7
	case 31:
		goto L6
	case 59:
		goto L30
	case 61:
		goto L24
	case 63:
		goto L23
	case 64:
		goto L28
	case 66:
		goto L27
	case 69:
		goto L29
	case 70:
		v348 = int32(_a_F_FigureColnameInternal_14)
		v352 = v35
		goto L3
	case 85:
		goto L17
	case 112:
		goto L9
	case 117:
		goto L16
	case 118:
		goto L15
	case 119:
		goto L14
	case 120:
		goto L13
	case 121, 122:
		goto L12
	case 124:
		goto L11
	case 125:
		goto L10
	}
L26:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v245+v246<<(uint(int32(2))%32)-int32(4))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	v348 = v253
	v352 = v35
	goto L3
L27:
	;
	goto L26
L28:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v243 != 0 {
		v21 = v243
		goto L25
	} else {
		goto L75
	}
L29:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v134 == int32(0) {
		goto L28
	} else {
		goto L52
	}
L30:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v40 == int32(0) {
		v366 = v3
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v44 <= v43 {
		v122 = v43
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v122 != 0 {
		v348 = v122
		v352 = v35
		goto L3
	} else {
		goto L51
	}
L33:
	;
	if v44 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v50 = int32(0)
	if v50 < v44 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v99 = v43
	v100 = int32(0)
	goto L36
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v100<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v116 != int32(468) {
		v122 = v99
		goto L32
	} else {
		goto L50
	}
L37:
	;
	v53 = v44
	goto L39
L38:
	;
	v53 = v50
	goto L39
L39:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v59 = int32(0)
	v63 = v43
	v64 = v59
	v65 = v59
	goto L40
L40:
	;
	v77 = v58 + v64<<(uint(int32(2))%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 == int32(468) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v53&int32(1) == int32(0) {
		v122 = v89
		goto L32
	} else {
		goto L49
	}
L42:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v83 = v82
	goto L44
L43:
	;
	v83 = v63
	goto L44
L44:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	if v85 == int32(468) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	v89 = v88
	goto L47
L46:
	;
	v89 = v83
	goto L47
L47:
	;
	v90 = int32(2)
	v91 = v64 + v90
	v93 = v65 + v90
	if v93 != v53&int32(2147483646) {
		v63 = v89
		v64 = v91
		v65 = v93
		goto L40
	} else {
		goto L48
	}
L48:
	;
	goto L41
L49:
	;
	v99 = v89
	v100 = v91
	goto L36
L50:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v122 = v119
	goto L32
L51:
	;
	v366 = v3
	goto L1
L52:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v137 <= int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v217 != 0 {
		v348 = v217
		v352 = v35
		goto L3
	} else {
		goto L74
	}
L54:
	;
	v217 = int32(0)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v141 = int32(0)
	if v137 != int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v145 = int32(0)
	if v145 < v137 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v194 = v141
	v195 = v141
	goto L59
L59:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206+v195<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v211 != int32(468) {
		v217 = v194
		goto L53
	} else {
		goto L73
	}
L60:
	;
	v148 = v137
	goto L62
L61:
	;
	v148 = v145
	goto L62
L62:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v154 = int32(0)
	v158 = v141
	v159 = v154
	v160 = v154
	goto L63
L63:
	;
	v172 = v153 + v159<<(uint(int32(2))%32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v174 == int32(468) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	if v148&int32(1) == int32(0) {
		v217 = v184
		goto L53
	} else {
		goto L72
	}
L65:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v178 = v177
	goto L67
L66:
	;
	v178 = v158
	goto L67
L67:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	if v180 == int32(468) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v184 = v183
	goto L70
L69:
	;
	v184 = v178
	goto L70
L70:
	;
	v185 = int32(2)
	v186 = v159 + v185
	v188 = v160 + v185
	if v188 != v148&int32(2147483646) {
		v158 = v184
		v159 = v186
		v160 = v188
		goto L63
	} else {
		goto L71
	}
L71:
	;
	goto L64
L72:
	;
	v194 = v184
	v195 = v186
	goto L59
L73:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v217 = v214
	goto L53
L74:
	;
	goto L28
L75:
	;
	v366 = v3
	goto L1
L76:
	;
	v348 = int32(_a_F_FigureColnameInternal_15)
	v352 = v35
	goto L3
L77:
	;
	return int32(0)
L78:
	;
	if base.Ui32(int32(1)) < base.Ui32(v260) {
		v366 = int32(2)
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v266 == int32(0) {
		v366 = v260
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v270+v271<<(uint(int32(2))%32)-int32(4))))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+4))
	v348 = v278
	v352 = int32(1)
	goto L3
L81:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 != int32(67) {
		v366 = v3
		goto L1
	} else {
		goto L83
	}
L82:
	;
	v348 = int32(_a_F_FigureColnameInternal_14)
	v352 = v35
	goto L3
L83:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+76))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v289)+12))
	if v290 != 0 {
		v348 = v290
		v352 = v35
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v366 = v3
	goto L1
L85:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v293) {
		v366 = int32(2)
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v348 = int32(_a_F_FigureColnameInternal_16)
	v352 = int32(1)
	goto L3
L87:
	;
	v348 = int32(_a_F_FigureColnameInternal_17)
	v352 = v35
	goto L3
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L77
	} else {
		goto L89
	}
L89:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v319
	F_errmsg_internal(m, int32(_a_F_FigureColnameInternal_18), v17)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L77
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_FigureColnameInternal_19), int32(2034), int32(_a_F_FigureColnameInternal_20))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L77
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v330<<(uint(int32(2))%32))+uint32(_c_F_FigureColnameInternal[1])))
	v348 = v335
	v352 = v35
	goto L3
L93:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v336<<(uint(int32(2))%32))+uint32(_c_F_FigureColnameInternal[2])))
	v348 = v341
	v352 = v35
	goto L3
}
func F_Float4ToHalf(m *base.Module, l0 float32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 float32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	v14 = base.I32_reinterpret_f32(l0)
	v16 = int32(base.Ui32(v14) >> (uint(v11) % 32))
	v17 = base.F32_abs(l0)
	if base.F32_ne(v17, math.Float32frombits(uint32(0x7f800000))) != 0 {
		v21 = v16 & int32(_a_F_Float4ToHalf_0)
		v23 = v14 & int32(_a_F_Float4ToHalf_1)
		if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v17)) {
			v97 = v21 | int32(base.Ui32(v23)>>(uint(int32(13))%32)) | int32(_a_F_Float4ToHalf_2)
		} else {
			v35 = int32(base.Ui32(v14)>>(uint(int32(23))%32)) & int32(255)
			if base.Ui32(v35) < base.Ui32(int32(99)) {
				v97 = v21
			} else {
				if base.Ui32(v35) <= base.Ui32(int32(112)) {
					v47 = int32(1)<<(uint(v35-int32(90))%32) + int32(base.Ui32(v23)>>(uint(int32(113)-v35)%32))
					v49 = v47 | v14
					v50 = v47
				} else {
					v49 = v14
					v50 = v23
				}
				v52 = int32(base.Ui32(v50) >> (uint(int32(13)) % 32))
				v57 = int32(1)
				v61 = int32(3)
				v62 = int32(base.Ui32(v50)>>(uint(int32(12))%32)) & v61
				if base.B2i32(v62 != v61)&(base.B2i32(v49&int32(4095) == int32(0))|base.B2i32(v62 != v57)) != 0 {
					v73 = v52
				} else {
					v73 = v52 + v57
				}
				v75 = base.B2i32(v73 == int32(1024))
				if v73 == int32(1024) {
					v76 = int32(-126)
				} else {
					v76 = int32(-127)
				}
				v77 = v76 + v35
				if int32(16) <= v77 {
					v97 = v21 | int32(_a_F_Float4ToHalf_3)
				} else {
					if int32(-15) < v77 {
						v89 = v77<<(uint(int32(10))%32) + int32(_a_F_Float4ToHalf_4) | v16&int32(_a_F_Float4ToHalf_0)
					} else {
						v89 = v21
					}
					if v73 == int32(1024) {
						v91 = int32(0)
					} else {
						v91 = v73
					}
					v97 = v89 | v91
				}
			}
		}
		if v97&int32(_a_F_Float4ToHalf_5) != int32(_a_F_Float4ToHalf_3) {
			v129 = v97
			m.G0 = v12 + int32(16)
			return v129 & int32(_a_F_Float4ToHalf_6)
		} else {
			v103 = F_palloc(m, int32(16))
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				v107 = F_float_to_shortest_decimal_bufn(m, l0, v103)
				mBase = m.M
				v109 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v107+v103))) = uint8(v109)
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v114 = m.ExcPending
				if v114 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v103
						F_errmsg(m, int32(_a_F_Float4ToHalf_7), v12)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_Float4ToHalf_8), int32(257), int32(_a_F_Float4ToHalf_9))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		v129 = v16 & int32(_a_F_Float4ToHalf_10)
		m.G0 = v12 + int32(16)
		return v129 & int32(_a_F_Float4ToHalf_6)
	}
}
func F_ForgetPortalSnapshots(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = v7 + int32(12)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ForgetPortalSnapshots[0]))
	F_hash_seq_init(m, v10, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_hash_seq_search(m, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v15
	v19 = v1
	goto L7
L5:
	;
	v34 = v1
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ForgetPortalSnapshots[1]))
	goto L14
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+100))
	if v22 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v34 = v27
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+100)) = int32(0)
	v27 = v19 + int32(1)
	goto L11
L10:
	;
	v27 = v19
	goto L11
L11:
	;
	v30 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v30 != 0 {
		v18 = v30
		v19 = v27
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	if v37 != int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = v1
	goto L18
L16:
	;
	v55 = v1
	goto L17
L17:
	;
	if v34 != v55 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v45 = v43 + int32(1)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v55 = v45
	goto L17
L20:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_ForgetPortalSnapshots[1]))
	goto L21
L21:
	;
	if v49 != int32(0) {
		v43 = v45
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	m.G0 = v7 + int32(32)
	return
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
	F_errmsg_internal(m, int32(_a_F_ForgetPortalSnapshots_0), v7)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_ForgetPortalSnapshots_1), int32(1292), int32(_a_F_ForgetPortalSnapshots_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FreeBulkInsertState(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_bms_free(m, v6)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		F_bms_free(m, v6)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_FreeSpaceMapPrepareTruncateRel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	v16 = m.G0
	v18 = v16 - int32(48)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v20 == int32(0) {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v26
		v30 = F_smgropen(m, v18+int32(24), v23)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v30
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
			if v36 != 0 {
				v44 = v36
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+80))
				*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v30)+76))
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v40
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v30)+72))
				v44 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(v30)+72)) = v44 + int32(1)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v49 = v48
			v50 = int32(-1)
			v52 = F_smgrexists(m, v49, int32(1))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v52 == int32(0) {
					v358 = v50
					m.G0 = v18 + int32(48)
					return v358
				} else {
					v56 = int32(4069)
					v57 = base.I32_div_u_s(l1, v56)
					*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = base.I64_extend_i32_u(v57) << (uint(int64(32)) % 64)
					v64 = l1 - v57*v56
					if v64 != 0 {
						v65 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v65
						v70 = F_fsm_readbuf(m, l0, v18+int32(16), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							if v70 == int32(0) {
								v358 = v50
								m.G0 = v18 + int32(48)
								return v358
							} else {
								F_LockBuffer(m, v70, int32(2))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v77 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
									v79 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
									*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v79 + int32(1)
									if v70 < int32(0) {
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[1]))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+(v70^int32(-1))<<(uint(int32(2))%32))))
										v100 = v92
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[2]))
										v100 = v94 + v70<<(uint(int32(13))%32) + int32(-8192)
									}
									v103 = v64 + v100 + int32(_a_F_FreeSpaceMapPrepareTruncateRel_1)
									if base.Ui32(v100-int32(-8192)) <= base.Ui32(v103) {
									} else {
										v108 = int32(4069) - v64
										v109 = int32(3)
										v110 = v108 & v109
										if base.Ui32(v64-int32(4066)) < base.Ui32(v109) {
											v159 = v103
											v161 = int32(0)
											v174 = int32(0)
											v175 = v159
											v177 = v161
											for {
												v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v188)
												v190 = int32(1)
												v194 = base.B2i32(v187 != v188) | v177
												v196 = v174 + v190
												if v196 != v110 {
													v174 = v196
													v175 = v175 + v190
													v177 = v194
													continue
												} else {
													break
												}
												break
											}
											v203 = v194
										} else {
											v118 = int32(0)
											v122 = v118
											v123 = v103
											v125 = v118
											for {
												v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
												v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
												v137 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v137)
												v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
												v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
												*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)) = uint16(v137)
												v143 = int32(4)
												v144 = v123 + v143
												v150 = base.B2i32(v136|(v139|v140)|v135 != v137) | v125
												v152 = v122 + v143
												if v152 != v108&int32(-4) {
													v122 = v152
													v123 = v144
													v125 = v150
													continue
												} else {
													break
												}
												break
											}
											if v110 == int32(0) {
												v203 = v150
											} else {
												v159 = v144
												v161 = v150
												v174 = int32(0)
												v175 = v159
												v177 = v161
												for {
													v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
													v188 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v188)
													v190 = int32(1)
													v194 = base.B2i32(v187 != v188) | v177
													v196 = v174 + v190
													if v196 != v110 {
														v174 = v196
														v175 = v175 + v190
														v177 = v194
														continue
													} else {
														break
													}
													break
												}
												v203 = v194
											}
										}
										if v203&int32(1) == int32(0) {
										} else {
											v220 = v100 + int32(28)
											v224 = int32(4094)
											for {
												if base.Ui32(int32(4081)) < base.Ui32(v224) {
													v252 = int32(0)
												} else {
													v241 = v224 << (uint(int32(1)) % 32)
													v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(29)+v241))))
													if v224 == int32(4081) {
														v252 = v243
													} else {
														v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v241)+2)))
														if base.Ui32(v247) < base.Ui32(v243) {
															v249 = v243
														} else {
															v249 = v247
														}
														v252 = v249
													}
												}
												v253 = v224 + v220
												v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
												if v254 != v252&int32(255) {
													*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v252)
												} else {
												}
												if v224 != 0 {
													v224 = v224 - int32(1)
													continue
												} else {
													break
												}
												break
											}
										}
									}
									F_MarkBufferDirty(m, v70)
									mBase = m.M
									v277 = m.ExcPending
									if v277 != 0 {
										return int32(0)
									} else {
										v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[3])))
										if v279 != 0 {
											v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
											v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
											*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
											F_UnlockReleaseBuffer(m, v70)
											mBase = m.M
											v313 = m.ExcPending
											if v313 != 0 {
												return int32(0)
											} else {
												v315 = base.I32_div_u_s(l1, int32(16556761))
												v358 = v57 + v315 + int32(3)
												m.G0 = v18 + int32(48)
												return v358
											}
										} else {
											v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
											v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+118)))
											if v281 != int32(112) {
												v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
												v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
												*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
												F_UnlockReleaseBuffer(m, v70)
												mBase = m.M
												v313 = m.ExcPending
												if v313 != 0 {
													return int32(0)
												} else {
													v315 = base.I32_div_u_s(l1, int32(16556761))
													v358 = v57 + v315 + int32(3)
													m.G0 = v18 + int32(48)
													return v358
												}
											} else {
												v285 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[4]))
												if v285 <= int32(0) {
													v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
													if v288 != 0 {
														v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
														v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
														*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
														F_UnlockReleaseBuffer(m, v70)
														mBase = m.M
														v313 = m.ExcPending
														if v313 != 0 {
															return int32(0)
														} else {
															v315 = base.I32_div_u_s(l1, int32(16556761))
															v358 = v57 + v315 + int32(3)
															m.G0 = v18 + int32(48)
															return v358
														}
													} else {
														v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
														if v289 != 0 {
															v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
															v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
															*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
															F_UnlockReleaseBuffer(m, v70)
															mBase = m.M
															v313 = m.ExcPending
															if v313 != 0 {
																return int32(0)
															} else {
																v315 = base.I32_div_u_s(l1, int32(16556761))
																v358 = v57 + v315 + int32(3)
																m.G0 = v18 + int32(48)
																return v358
															}
														} else {
															v291 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[5]))
															v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+252))
															if base.B2i32(v292 != int32(0)) == int32(0) {
																v298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[6])))
																if v298&int32(1) == int32(0) {
																	v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																	v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																	*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																	F_UnlockReleaseBuffer(m, v70)
																	mBase = m.M
																	v313 = m.ExcPending
																	if v313 != 0 {
																		return int32(0)
																	} else {
																		v315 = base.I32_div_u_s(l1, int32(16556761))
																		v358 = v57 + v315 + int32(3)
																		m.G0 = v18 + int32(48)
																		return v358
																	}
																} else {
																	F_log_newpage_buffer(m, v70, int32(0))
																	mBase = m.M
																	v305 = m.ExcPending
																	if v305 != 0 {
																		return int32(0)
																	} else {
																		v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																		v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																		*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																		F_UnlockReleaseBuffer(m, v70)
																		mBase = m.M
																		v313 = m.ExcPending
																		if v313 != 0 {
																			return int32(0)
																		} else {
																			v315 = base.I32_div_u_s(l1, int32(16556761))
																			v358 = v57 + v315 + int32(3)
																			m.G0 = v18 + int32(48)
																			return v358
																		}
																	}
																}
															} else {
																F_log_newpage_buffer(m, v70, int32(0))
																mBase = m.M
																v305 = m.ExcPending
																if v305 != 0 {
																	return int32(0)
																} else {
																	v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																	v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																	*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																	F_UnlockReleaseBuffer(m, v70)
																	mBase = m.M
																	v313 = m.ExcPending
																	if v313 != 0 {
																		return int32(0)
																	} else {
																		v315 = base.I32_div_u_s(l1, int32(16556761))
																		v358 = v57 + v315 + int32(3)
																		m.G0 = v18 + int32(48)
																		return v358
																	}
																}
															}
														}
													}
												} else {
													v291 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[5]))
													v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+252))
													if base.B2i32(v292 != int32(0)) == int32(0) {
														v298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[6])))
														if v298&int32(1) == int32(0) {
															v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
															v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
															*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
															F_UnlockReleaseBuffer(m, v70)
															mBase = m.M
															v313 = m.ExcPending
															if v313 != 0 {
																return int32(0)
															} else {
																v315 = base.I32_div_u_s(l1, int32(16556761))
																v358 = v57 + v315 + int32(3)
																m.G0 = v18 + int32(48)
																return v358
															}
														} else {
															F_log_newpage_buffer(m, v70, int32(0))
															mBase = m.M
															v305 = m.ExcPending
															if v305 != 0 {
																return int32(0)
															} else {
																v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																F_UnlockReleaseBuffer(m, v70)
																mBase = m.M
																v313 = m.ExcPending
																if v313 != 0 {
																	return int32(0)
																} else {
																	v315 = base.I32_div_u_s(l1, int32(16556761))
																	v358 = v57 + v315 + int32(3)
																	m.G0 = v18 + int32(48)
																	return v358
																}
															}
														}
													} else {
														F_log_newpage_buffer(m, v70, int32(0))
														mBase = m.M
														v305 = m.ExcPending
														if v305 != 0 {
															return int32(0)
														} else {
															v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
															v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
															*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
															F_UnlockReleaseBuffer(m, v70)
															mBase = m.M
															v313 = m.ExcPending
															if v313 != 0 {
																return int32(0)
															} else {
																v315 = base.I32_div_u_s(l1, int32(16556761))
																v358 = v57 + v315 + int32(3)
																m.G0 = v18 + int32(48)
																return v358
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v321 = base.I32_div_u_s(l1, int32(16556761))
						v324 = v57 + v321 + int32(2)
						v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v325 != 0 {
							v349 = v325
							v351 = F_smgrnblocks(m, v349, int32(1))
							mBase = m.M
							v352 = m.ExcPending
							if v352 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v351) <= base.Ui32(v324) {
									v354 = int32(-1)
								} else {
									v354 = v324
								}
								v358 = v354
								m.G0 = v18 + int32(48)
								return v358
							}
						} else {
							v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v327
							v329 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v18))) = v329
							v331 = F_smgropen(m, v18, v326)
							mBase = m.M
							v332 = m.ExcPending
							if v332 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v331
								v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+72))
								if v335 != 0 {
									v343 = v335
								} else {
									v336 = *(*int32)(unsafe.Add(mBase, uint32(v331)+76))
									v337 = *(*int32)(unsafe.Add(mBase, uint32(v331)+80))
									*(*int32)(unsafe.Add(mBase, uint32(v336)+4)) = v337
									v339 = *(*int32)(unsafe.Add(mBase, uint32(v331)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v337))) = v339
									v341 = *(*int32)(unsafe.Add(mBase, uint32(v331)+72))
									v343 = v341
								}
								*(*int32)(unsafe.Add(mBase, uint32(v331)+72)) = v343 + int32(1)
								v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v349 = v347
								v351 = F_smgrnblocks(m, v349, int32(1))
								mBase = m.M
								v352 = m.ExcPending
								if v352 != 0 {
									return int32(0)
								} else {
									if base.Ui32(v351) <= base.Ui32(v324) {
										v354 = int32(-1)
									} else {
										v354 = v324
									}
									v358 = v354
									m.G0 = v18 + int32(48)
									return v358
								}
							}
						}
					}
				}
			}
		}
	} else {
		v49 = v20
		v50 = int32(-1)
		v52 = F_smgrexists(m, v49, int32(1))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			if v52 == int32(0) {
				v358 = v50
				m.G0 = v18 + int32(48)
				return v358
			} else {
				v56 = int32(4069)
				v57 = base.I32_div_u_s(l1, v56)
				*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = base.I64_extend_i32_u(v57) << (uint(int64(32)) % 64)
				v64 = l1 - v57*v56
				if v64 != 0 {
					v65 = *(*int64)(unsafe.Add(mBase, uint32(v18)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v65
					v70 = F_fsm_readbuf(m, l0, v18+int32(16), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						if v70 == int32(0) {
							v358 = v50
							m.G0 = v18 + int32(48)
							return v358
						} else {
							F_LockBuffer(m, v70, int32(2))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								v77 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
								v79 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
								*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v79 + int32(1)
								if v70 < int32(0) {
									v86 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[1]))
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+(v70^int32(-1))<<(uint(int32(2))%32))))
									v100 = v92
								} else {
									v94 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[2]))
									v100 = v94 + v70<<(uint(int32(13))%32) + int32(-8192)
								}
								v103 = v64 + v100 + int32(_a_F_FreeSpaceMapPrepareTruncateRel_1)
								if base.Ui32(v100-int32(-8192)) <= base.Ui32(v103) {
								} else {
									v108 = int32(4069) - v64
									v109 = int32(3)
									v110 = v108 & v109
									if base.Ui32(v64-int32(4066)) < base.Ui32(v109) {
										v159 = v103
										v161 = int32(0)
										v174 = int32(0)
										v175 = v159
										v177 = v161
										for {
											v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
											v188 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v188)
											v190 = int32(1)
											v194 = base.B2i32(v187 != v188) | v177
											v196 = v174 + v190
											if v196 != v110 {
												v174 = v196
												v175 = v175 + v190
												v177 = v194
												continue
											} else {
												break
											}
											break
										}
										v203 = v194
									} else {
										v118 = int32(0)
										v122 = v118
										v123 = v103
										v125 = v118
										for {
											v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
											v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
											v137 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(v123))) = uint16(v137)
											v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
											v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
											*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)) = uint16(v137)
											v143 = int32(4)
											v144 = v123 + v143
											v150 = base.B2i32(v136|(v139|v140)|v135 != v137) | v125
											v152 = v122 + v143
											if v152 != v108&int32(-4) {
												v122 = v152
												v123 = v144
												v125 = v150
												continue
											} else {
												break
											}
											break
										}
										if v110 == int32(0) {
											v203 = v150
										} else {
											v159 = v144
											v161 = v150
											v174 = int32(0)
											v175 = v159
											v177 = v161
											for {
												v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v175))) = uint8(v188)
												v190 = int32(1)
												v194 = base.B2i32(v187 != v188) | v177
												v196 = v174 + v190
												if v196 != v110 {
													v174 = v196
													v175 = v175 + v190
													v177 = v194
													continue
												} else {
													break
												}
												break
											}
											v203 = v194
										}
									}
									if v203&int32(1) == int32(0) {
									} else {
										v220 = v100 + int32(28)
										v224 = int32(4094)
										for {
											if base.Ui32(int32(4081)) < base.Ui32(v224) {
												v252 = int32(0)
											} else {
												v241 = v224 << (uint(int32(1)) % 32)
												v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+int32(29)+v241))))
												if v224 == int32(4081) {
													v252 = v243
												} else {
													v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220+v241)+2)))
													if base.Ui32(v247) < base.Ui32(v243) {
														v249 = v243
													} else {
														v249 = v247
													}
													v252 = v249
												}
											}
											v253 = v224 + v220
											v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253))))
											if v254 != v252&int32(255) {
												*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v252)
											} else {
											}
											if v224 != 0 {
												v224 = v224 - int32(1)
												continue
											} else {
												break
											}
											break
										}
									}
								}
								F_MarkBufferDirty(m, v70)
								mBase = m.M
								v277 = m.ExcPending
								if v277 != 0 {
									return int32(0)
								} else {
									v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[3])))
									if v279 != 0 {
										v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
										v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
										*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
										F_UnlockReleaseBuffer(m, v70)
										mBase = m.M
										v313 = m.ExcPending
										if v313 != 0 {
											return int32(0)
										} else {
											v315 = base.I32_div_u_s(l1, int32(16556761))
											v358 = v57 + v315 + int32(3)
											m.G0 = v18 + int32(48)
											return v358
										}
									} else {
										v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
										v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+118)))
										if v281 != int32(112) {
											v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
											v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
											*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
											F_UnlockReleaseBuffer(m, v70)
											mBase = m.M
											v313 = m.ExcPending
											if v313 != 0 {
												return int32(0)
											} else {
												v315 = base.I32_div_u_s(l1, int32(16556761))
												v358 = v57 + v315 + int32(3)
												m.G0 = v18 + int32(48)
												return v358
											}
										} else {
											v285 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[4]))
											if v285 <= int32(0) {
												v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
												if v288 != 0 {
													v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
													v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
													*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
													F_UnlockReleaseBuffer(m, v70)
													mBase = m.M
													v313 = m.ExcPending
													if v313 != 0 {
														return int32(0)
													} else {
														v315 = base.I32_div_u_s(l1, int32(16556761))
														v358 = v57 + v315 + int32(3)
														m.G0 = v18 + int32(48)
														return v358
													}
												} else {
													v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													if v289 != 0 {
														v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
														v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
														*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
														F_UnlockReleaseBuffer(m, v70)
														mBase = m.M
														v313 = m.ExcPending
														if v313 != 0 {
															return int32(0)
														} else {
															v315 = base.I32_div_u_s(l1, int32(16556761))
															v358 = v57 + v315 + int32(3)
															m.G0 = v18 + int32(48)
															return v358
														}
													} else {
														v291 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[5]))
														v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+252))
														if base.B2i32(v292 != int32(0)) == int32(0) {
															v298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[6])))
															if v298&int32(1) == int32(0) {
																v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																F_UnlockReleaseBuffer(m, v70)
																mBase = m.M
																v313 = m.ExcPending
																if v313 != 0 {
																	return int32(0)
																} else {
																	v315 = base.I32_div_u_s(l1, int32(16556761))
																	v358 = v57 + v315 + int32(3)
																	m.G0 = v18 + int32(48)
																	return v358
																}
															} else {
																F_log_newpage_buffer(m, v70, int32(0))
																mBase = m.M
																v305 = m.ExcPending
																if v305 != 0 {
																	return int32(0)
																} else {
																	v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																	v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																	*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																	F_UnlockReleaseBuffer(m, v70)
																	mBase = m.M
																	v313 = m.ExcPending
																	if v313 != 0 {
																		return int32(0)
																	} else {
																		v315 = base.I32_div_u_s(l1, int32(16556761))
																		v358 = v57 + v315 + int32(3)
																		m.G0 = v18 + int32(48)
																		return v358
																	}
																}
															}
														} else {
															F_log_newpage_buffer(m, v70, int32(0))
															mBase = m.M
															v305 = m.ExcPending
															if v305 != 0 {
																return int32(0)
															} else {
																v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
																v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
																*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
																F_UnlockReleaseBuffer(m, v70)
																mBase = m.M
																v313 = m.ExcPending
																if v313 != 0 {
																	return int32(0)
																} else {
																	v315 = base.I32_div_u_s(l1, int32(16556761))
																	v358 = v57 + v315 + int32(3)
																	m.G0 = v18 + int32(48)
																	return v358
																}
															}
														}
													}
												}
											} else {
												v291 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[5]))
												v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+252))
												if base.B2i32(v292 != int32(0)) == int32(0) {
													v298 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[6])))
													if v298&int32(1) == int32(0) {
														v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
														v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
														*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
														F_UnlockReleaseBuffer(m, v70)
														mBase = m.M
														v313 = m.ExcPending
														if v313 != 0 {
															return int32(0)
														} else {
															v315 = base.I32_div_u_s(l1, int32(16556761))
															v358 = v57 + v315 + int32(3)
															m.G0 = v18 + int32(48)
															return v358
														}
													} else {
														F_log_newpage_buffer(m, v70, int32(0))
														mBase = m.M
														v305 = m.ExcPending
														if v305 != 0 {
															return int32(0)
														} else {
															v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
															v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
															*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
															F_UnlockReleaseBuffer(m, v70)
															mBase = m.M
															v313 = m.ExcPending
															if v313 != 0 {
																return int32(0)
															} else {
																v315 = base.I32_div_u_s(l1, int32(16556761))
																v358 = v57 + v315 + int32(3)
																m.G0 = v18 + int32(48)
																return v358
															}
														}
													}
												} else {
													F_log_newpage_buffer(m, v70, int32(0))
													mBase = m.M
													v305 = m.ExcPending
													if v305 != 0 {
														return int32(0)
													} else {
														v306 = int32(_a_F_FreeSpaceMapPrepareTruncateRel_0)
														v308 = *(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0]))
														*(*int32)(unsafe.Add(mBase, _c_F_FreeSpaceMapPrepareTruncateRel[0])) = v308 - int32(1)
														F_UnlockReleaseBuffer(m, v70)
														mBase = m.M
														v313 = m.ExcPending
														if v313 != 0 {
															return int32(0)
														} else {
															v315 = base.I32_div_u_s(l1, int32(16556761))
															v358 = v57 + v315 + int32(3)
															m.G0 = v18 + int32(48)
															return v358
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v321 = base.I32_div_u_s(l1, int32(16556761))
					v324 = v57 + v321 + int32(2)
					v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v325 != 0 {
						v349 = v325
						v351 = F_smgrnblocks(m, v349, int32(1))
						mBase = m.M
						v352 = m.ExcPending
						if v352 != 0 {
							return int32(0)
						} else {
							if base.Ui32(v351) <= base.Ui32(v324) {
								v354 = int32(-1)
							} else {
								v354 = v324
							}
							v358 = v354
							m.G0 = v18 + int32(48)
							return v358
						}
					} else {
						v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v327
						v329 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v18))) = v329
						v331 = F_smgropen(m, v18, v326)
						mBase = m.M
						v332 = m.ExcPending
						if v332 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v331
							v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+72))
							if v335 != 0 {
								v343 = v335
							} else {
								v336 = *(*int32)(unsafe.Add(mBase, uint32(v331)+76))
								v337 = *(*int32)(unsafe.Add(mBase, uint32(v331)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v336)+4)) = v337
								v339 = *(*int32)(unsafe.Add(mBase, uint32(v331)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v337))) = v339
								v341 = *(*int32)(unsafe.Add(mBase, uint32(v331)+72))
								v343 = v341
							}
							*(*int32)(unsafe.Add(mBase, uint32(v331)+72)) = v343 + int32(1)
							v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v349 = v347
							v351 = F_smgrnblocks(m, v349, int32(1))
							mBase = m.M
							v352 = m.ExcPending
							if v352 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v351) <= base.Ui32(v324) {
									v354 = int32(-1)
								} else {
									v354 = v324
								}
								v358 = v354
								m.G0 = v18 + int32(48)
								return v358
							}
						}
					}
				}
			}
		}
	}
}
func F_fastgetattr_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn13878(m, l0, l1, l2, l3, int32(_a_F_fastgetattr_2_0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_ferror(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return int32(base.Ui32(v2)>>(uint(int32(5))%32)) & int32(1)
}
func F_fetch_input_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	if v4 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_fetch_input_tuple[0]))
		if v6 != 0 {
			F_ProcessInterrupts(m)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
				v12 = v11
				v14 = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
				v17 = F_tuplesort_gettupleslot(m, v12, int32(1), v14, v15, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						v44 = int32(0)
						return v44
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
						v30 = v21
						if v30 == int32(0) {
							return int32(0)
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
							if v35&int32(2) != 0 {
								v44 = v30
								return v44
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
								if v38 == int32(0) {
									v44 = v30
									return v44
								} else {
									F_tuplesort_puttupleslot(m, v38, v30)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v44 = v30
										return v44
									}
								}
							}
						}
					}
				}
			}
		} else {
			v12 = v4
			v14 = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
			v17 = F_tuplesort_gettupleslot(m, v12, int32(1), v14, v15, v14)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				if v17 == int32(0) {
					v44 = int32(0)
					return v44
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
					v30 = v21
					if v30 == int32(0) {
						return int32(0)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
						if v35&int32(2) != 0 {
							v44 = v30
							return v44
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
							if v38 == int32(0) {
								v44 = v30
								return v44
							} else {
								F_tuplesort_puttupleslot(m, v38, v30)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = v30
									return v44
								}
							}
						}
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
		if v23 != 0 {
			F_ExecReScan(m, v22)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
				v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v22)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v30 = v27
					if v30 == int32(0) {
						return int32(0)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
						if v35&int32(2) != 0 {
							v44 = v30
							return v44
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
							if v38 == int32(0) {
								v44 = v30
								return v44
							} else {
								F_tuplesort_puttupleslot(m, v38, v30)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v44 = v30
									return v44
								}
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
			v27 = m.T0[v26].(func(*base.Module, int32) int32)(m, v22)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v30 = v27
				if v30 == int32(0) {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+4)))
					if v35&int32(2) != 0 {
						v44 = v30
						return v44
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
						if v38 == int32(0) {
							v44 = v30
							return v44
						} else {
							F_tuplesort_puttupleslot(m, v38, v30)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v44 = v30
								return v44
							}
						}
					}
				}
			}
		}
	}
}
func F_fillQT(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	F_check_stack_depth(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v8 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return
L4:
	;
	v12 = l1
	v13 = v7
	goto L7
L5:
	;
	v43 = l1
	v44 = v7
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v54 = v52 & int32(4095)
	if v54 != 0 {
		goto L13
	} else {
		goto L14
	}
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v15))) = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 + int32(12)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	F_fillQT(m, l0, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v43 = v35
	v44 = v38
	goto L6
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v26 != int32(2) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = base.I32_div_s(v29-v15, int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v39 != int32(1) {
		v12 = v35
		v13 = v38
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	base.MemoryCopy(m, v55, v56, v54)
	goto L15
L14:
	;
	goto L15
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v60 = int32(4095)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v65 = int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v59&v60 | (v62-v63)<<(uint(v65)%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69+v71&v60))) = uint8(v75)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v79 + v65
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v83 + v78&v60 + int32(1)
	goto L3
}
func F_finalize_plan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
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
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
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
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
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
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v558 int32
	_ = v558
	var v568 int32
	_ = v568
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v820 int32
	_ = v820
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v859 int32
	_ = v859
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	if l1 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L14
	} else {
		goto L243
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L14
	} else {
		goto L240
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L14
	} else {
		goto L237
	}
L4:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = l0
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v21 == v18 {
		v129 = l3
		v137 = v6
		v138 = v6
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v859 = int32(0)
	goto L6
L6:
	;
	m.G0 = v16 + int32(48)
	return v859
L7:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v141 = v16 + int32(40)
	v142 = F_finalize_primnode(m, v139, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L28
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if int32(0) < v24 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = v6
	v38 = v6
	v39 = v6
	goto L12
L10:
	;
	v119 = v6
	v120 = v6
	goto L11
L11:
	;
	if v119 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v44 = int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v43+v34<<(uint(v44)%32))))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v42+v48<<(uint(v44)%32)-int32(4))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+64))
	v56 = F_bms_add_members(m, v39, v55)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v119 = v102
	v120 = v56
	goto L11
L14:
	;
	return int32(0)
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v60 == int32(0) {
		v102 = v38
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v105 = v34 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v105 < v106 {
		v34 = v105
		v38 = v102
		v39 = v56
		goto L12
	} else {
		goto L23
	}
L17:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v64 <= v63 {
		v102 = v38
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v76 = v63
	v78 = v38
	goto L19
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v76<<(uint(int32(2))%32))))
	v85 = F_bms_add_member(m, v78, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L21
	}
L20:
	;
	v102 = v85
	goto L16
L21:
	;
	v88 = v76 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v88 < v89 {
		v76 = v88
		v78 = v85
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L13
L24:
	;
	v129 = l3
	v137 = int32(0)
	v138 = v120
	goto L7
L25:
	;
	goto L26
L26:
	;
	v124 = F_bms_union(m, l3, v119)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L14
	} else {
		goto L27
	}
L27:
	;
	v129 = v124
	v137 = v119
	v138 = v120
	goto L7
L28:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v145 = F_finalize_primnode(m, v144, v141)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v147 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if l2 < int32(0) {
		goto L3
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v156 = int32(-1)
	v157 = int32(0)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v158 - int32(331) {
	case 0:
		goto L36
	case 1, 29, 31, 32, 33, 36, 40:
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	case 2:
		goto L54
	case 3:
		goto L53
	case 4:
		goto L52
	case 5:
		goto L44
	case 6:
		goto L51
	case 7:
		goto L50
	case 8:
		goto L71
	case 9:
		goto L70
	case 10:
		goto L69
	case 11:
		goto L68
	case 12:
		goto L67
	case 13:
		goto L66
	case 14:
		goto L65
	case 15:
		goto L64
	case 16:
		goto L63
	case 17:
		goto L62
	case 18:
		goto L60
	case 19:
		goto L61
	case 20:
		goto L59
	case 21:
		goto L57
	case 22:
		goto L58
	case 23:
		goto L56
	case 24:
		goto L55
	case 25:
		goto L49
	default:
		goto L37
	case 27:
		goto L48
	case 28:
		goto L47
	case 30:
		goto L38
	case 34:
		goto L42
	case 35:
		goto L41
	case 37:
		goto L40
	case 38:
		goto L39
	case 39:
		goto L46
	case 41:
		goto L43
	case 42:
		goto L45
	}
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v153 = F_bms_add_member(m, v152, l2)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v153
	goto L32
L35:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v746 = F_finalize_plan(m, l0, v745, v734, v735, v736)
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L14
	} else {
		goto L202
	}
L36:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v730 = F_finalize_primnode(m, v727, v16+int32(40))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L14
	} else {
		goto L201
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L14
	} else {
		goto L198
	}
L38:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v711 = F_finalize_primnode(m, v708, v16+int32(40))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L14
	} else {
		goto L197
	}
L39:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v701 < int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v701
		v741 = v157
		goto L35
	} else {
		goto L194
	}
L40:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v694 < int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v694
		v741 = v157
		goto L35
	} else {
		goto L191
	}
L41:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	v688 = v16 + int32(40)
	v689 = F_finalize_primnode(m, v686, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L14
	} else {
		goto L189
	}
L42:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v670 != int32(2) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L186
	}
L43:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v662 = F_bms_copy(m, v129)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L14
	} else {
		goto L182
	}
L44:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v657 = F_bms_copy(m, v129)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L14
	} else {
		goto L180
	}
L45:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v650 = v16 + int32(40)
	v651 = F_finalize_primnode(m, v648, v650)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L14
	} else {
		goto L178
	}
L46:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v646 = F_finalize_primnode(m, v643, v16+int32(40))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L14
	} else {
		goto L177
	}
L47:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v637 = v16 + int32(40)
	v638 = F_finalize_primnode(m, v635, v637)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L14
	} else {
		goto L175
	}
L48:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v629 = v16 + int32(40)
	v630 = F_finalize_primnode(m, v627, v629)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L14
	} else {
		goto L173
	}
L49:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v593 = F_finalize_primnode(m, v590, v16+int32(40))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L14
	} else {
		goto L166
	}
L50:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v555 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L159
	}
L51:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	if v520 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L152
	}
L52:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v485 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L145
	}
L53:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v450 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L138
	}
L54:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	v431 = F_bms_copy(m, v129)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L14
	} else {
		goto L131
	}
L55:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v390 = F_finalize_primnode(m, v387, v16+int32(40))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L14
	} else {
		goto L122
	}
L56:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v377 = v16 + int32(40)
	v378 = F_finalize_primnode(m, v375, v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L14
	} else {
		goto L119
	}
L57:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v372 = F_bms_add_members(m, v371, l4)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L14
	} else {
		goto L118
	}
L58:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v365 = F_bms_add_member(m, v363, v364)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L14
	} else {
		goto L116
	}
L59:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v339 <= int32(0) {
		goto L2
	} else {
		goto L111
	}
L60:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v333 = F_finalize_primnode(m, v330, v16+int32(40))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L14
	} else {
		goto L109
	}
L61:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v324 = F_finalize_primnode(m, v321, v16+int32(40))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L14
	} else {
		goto L107
	}
L62:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v260 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v234 = F_find_base_rel(m, l0, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L89
	}
L64:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v227 = F_finalize_primnode(m, v224, v16+int32(40))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L14
	} else {
		goto L87
	}
L65:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v218 = F_finalize_primnode(m, v215, v16+int32(40))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L14
	} else {
		goto L85
	}
L66:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v209 = F_finalize_primnode(m, v206, v16+int32(40))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L83
	}
L67:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v204 = F_finalize_primnode(m, v201, v16+int32(40))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L82
	}
L68:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v188 = v16 + int32(40)
	v189 = F_finalize_primnode(m, v186, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L78
	}
L69:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	v176 = v16 + int32(40)
	v177 = F_finalize_primnode(m, v174, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L14
	} else {
		goto L75
	}
L70:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v168 = F_finalize_primnode(m, v165, v16+int32(40))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L14
	} else {
		goto L73
	}
L71:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v162 = F_bms_add_members(m, v161, l4)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L14
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v162
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L73:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v171 = F_bms_add_members(m, v170, l4)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v171
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L75:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v180 = F_finalize_primnode(m, v179, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L76
	}
L76:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v183 = F_bms_add_members(m, v182, l4)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L14
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v183
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L78:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v192 = F_finalize_primnode(m, v191, v188)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v195 = F_finalize_primnode(m, v194, v188)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L14
	} else {
		goto L80
	}
L80:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v198 = F_bms_add_members(m, v197, l4)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L14
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v198
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L82:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L83:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v212 = F_bms_add_members(m, v211, l4)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L14
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v212
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L85:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v221 = F_bms_add_members(m, v220, l4)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v221
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L87:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v230 = F_bms_add_members(m, v229, l4)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L14
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v230
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L89:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)+140))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+24))
	if int32(0) <= l2 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v240 = F_bms_copy(m, v237)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L93
	}
L91:
	;
	v245 = v237
	v246 = v236
	goto L92
L92:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v249 = F_finalize_plan(m, v246, v247, l2, v245, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L14
	} else {
		goto L95
	}
L93:
	;
	v242 = F_bms_add_member(m, v240, l2)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L14
	} else {
		goto L94
	}
L94:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v234)+140))
	v245 = v242
	v246 = v244
	goto L92
L95:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+64))
	v254 = F_bms_add_members(m, v251, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L14
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v254
	v257 = F_bms_add_members(m, v254, l4)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L14
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v257
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L98:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v318 = F_bms_add_members(m, v317, l4)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L14
	} else {
		goto L106
	}
L99:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v263 <= int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v273 = int32(0)
	goto L101
L101:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v273<<(uint(int32(2))%32))))
	v285 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v285
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(0)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v284)+4))
	v292 = F_finalize_primnode(m, v289, v16+int32(32))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L14
	} else {
		goto L103
	}
L102:
	;
	goto L98
L103:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+28)) = v294
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v297 = F_bms_add_members(m, v296, v294)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L14
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v297
	v301 = v273 + int32(1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+4))
	if v301 < v302 {
		v273 = v301
		goto L101
	} else {
		goto L105
	}
L105:
	;
	goto L102
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v318
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L107:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v327 = F_bms_add_members(m, v326, l4)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L14
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v327
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L109:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v336 = F_bms_add_members(m, v335, l4)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L14
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v336
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L111:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+8))
	if v343 == int32(0) {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	if v346 < v339 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v349+v339<<(uint(int32(2))%32)-int32(4))))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v355)+64))
	v357 = F_bms_add_members(m, v348, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L14
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v357
	v360 = F_bms_add_members(m, v357, l4)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L14
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v360
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v365
	v368 = F_bms_add_members(m, v365, l4)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L14
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v368
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v372
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L119:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	v381 = F_finalize_primnode(m, v380, v377)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L14
	} else {
		goto L120
	}
L120:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v384 = F_bms_add_members(m, v383, l4)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L14
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v384
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L122:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v393 = F_bms_add_members(m, v392, l4)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L14
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v393
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l1)+84))
	if v396 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L124
	}
L124:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v399 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L125
	}
L125:
	;
	v409 = int32(0)
	v413 = v393
	goto L126
L126:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v396)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416+v409<<(uint(int32(2))%32))))
	v421 = F_finalize_plan(m, l0, v420, l2, v129, l4)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L14
	} else {
		goto L128
	}
L127:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L128:
	;
	v423 = F_bms_add_members(m, v413, v421)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L14
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v423
	v427 = v409 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v396)+4))
	if v427 < v428 {
		v409 = v427
		v413 = v423
		goto L126
	} else {
		goto L130
	}
L130:
	;
	goto L127
L131:
	;
	v433 = F_bms_add_member(m, v431, v430)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	v435 = F_bms_copy(m, l4)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L14
	} else {
		goto L133
	}
L133:
	;
	v437 = F_bms_add_member(m, v435, v430)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L14
	} else {
		goto L134
	}
L134:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v441 = v16 + int32(40)
	v442 = F_finalize_primnode(m, v439, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L14
	} else {
		goto L135
	}
L135:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l1)+140))
	v445 = F_finalize_primnode(m, v444, v441)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L14
	} else {
		goto L136
	}
L136:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l1)+148))
	v448 = F_finalize_primnode(m, v447, v441)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	v734 = l2
	v735 = v433
	v736 = v437
	v740 = v430
	v741 = v157
	goto L35
L138:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v453 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L139
	}
L139:
	;
	v463 = int32(0)
	goto L140
L140:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v471+v463<<(uint(int32(2))%32))))
	v476 = F_finalize_plan(m, l0, v475, l2, v129, l4)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L14
	} else {
		goto L142
	}
L141:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L142:
	;
	v478 = F_bms_add_members(m, v470, v476)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v478
	v482 = v463 + int32(1)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	if v482 < v483 {
		v463 = v482
		goto L140
	} else {
		goto L144
	}
L144:
	;
	goto L141
L145:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	if v488 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L146
	}
L146:
	;
	v498 = int32(0)
	goto L147
L147:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v506+v498<<(uint(int32(2))%32))))
	v511 = F_finalize_plan(m, l0, v510, l2, v129, l4)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L14
	} else {
		goto L149
	}
L148:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L149:
	;
	v513 = F_bms_add_members(m, v505, v511)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L14
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v513
	v517 = v498 + int32(1)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	if v517 < v518 {
		v498 = v517
		goto L147
	} else {
		goto L151
	}
L151:
	;
	goto L148
L152:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v523 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L153
	}
L153:
	;
	v533 = int32(0)
	goto L154
L154:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v520)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v541+v533<<(uint(int32(2))%32))))
	v546 = F_finalize_plan(m, l0, v545, l2, v129, l4)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L14
	} else {
		goto L156
	}
L155:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L156:
	;
	v548 = F_bms_add_members(m, v540, v546)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L14
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v548
	v552 = v533 + int32(1)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v520)+4))
	if v552 < v553 {
		v533 = v552
		goto L154
	} else {
		goto L158
	}
L158:
	;
	goto L155
L159:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v558 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L160
	}
L160:
	;
	v568 = int32(0)
	goto L161
L161:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v576+v568<<(uint(int32(2))%32))))
	v581 = F_finalize_plan(m, l0, v580, l2, v129, l4)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L14
	} else {
		goto L163
	}
L162:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L163:
	;
	v583 = F_bms_add_members(m, v575, v581)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L14
	} else {
		goto L164
	}
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v583
	v587 = v568 + int32(1)
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v587 < v588 {
		v568 = v587
		goto L161
	} else {
		goto L165
	}
L165:
	;
	goto L162
L166:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v595 == int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L167
	}
L167:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v598 <= int32(0) {
		v734 = l2
		v735 = v129
		v736 = l4
		v740 = v156
		v741 = v157
		goto L35
	} else {
		goto L168
	}
L168:
	;
	v608 = int32(0)
	v611 = v157
	goto L169
L169:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v595)+12))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v615+v608<<(uint(int32(2))%32))))
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v619)+4))
	v621 = F_bms_add_member(m, v611, v620)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L14
	} else {
		goto L171
	}
L170:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v621
	goto L35
L171:
	;
	v624 = v608 + int32(1)
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v595)+4))
	if v624 < v625 {
		v608 = v624
		v611 = v621
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	v633 = F_finalize_primnode(m, v632, v629)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L14
	} else {
		goto L174
	}
L174:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L175:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v641 = F_finalize_primnode(m, v640, v637)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L14
	} else {
		goto L176
	}
L176:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L177:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L178:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v654 = F_finalize_primnode(m, v653, v650)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L14
	} else {
		goto L179
	}
L179:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L180:
	;
	v659 = F_bms_add_member(m, v657, v656)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L14
	} else {
		goto L181
	}
L181:
	;
	v734 = l2
	v735 = v659
	v736 = l4
	v740 = v656
	v741 = v157
	goto L35
L182:
	;
	v664 = F_bms_add_member(m, v662, v661)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L14
	} else {
		goto L183
	}
L183:
	;
	v666 = F_bms_copy(m, l4)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L14
	} else {
		goto L184
	}
L184:
	;
	v668 = F_bms_add_member(m, v666, v661)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L14
	} else {
		goto L185
	}
L185:
	;
	v734 = l2
	v735 = v664
	v736 = v668
	v740 = v661
	v741 = v157
	goto L35
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l0
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v678 = v16 + int32(32)
	v679 = F_finalize_agg_primnode(m, v676, v678)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L14
	} else {
		goto L187
	}
L187:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v682 = F_finalize_agg_primnode(m, v681, v678)
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L14
	} else {
		goto L188
	}
L188:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v684
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L189:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	v692 = F_finalize_primnode(m, v691, v688)
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L14
	} else {
		goto L190
	}
L190:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L191:
	;
	v697 = F_bms_copy(m, v129)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L14
	} else {
		goto L192
	}
L192:
	;
	v699 = F_bms_add_member(m, v697, v694)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L14
	} else {
		goto L193
	}
L193:
	;
	v734 = v694
	v735 = v699
	v736 = l4
	v740 = v694
	v741 = v157
	goto L35
L194:
	;
	v704 = F_bms_copy(m, v129)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L14
	} else {
		goto L195
	}
L195:
	;
	v706 = F_bms_add_member(m, v704, v701)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L14
	} else {
		goto L196
	}
L196:
	;
	v734 = v701
	v735 = v706
	v736 = l4
	v740 = v701
	v741 = v157
	goto L35
L197:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L198:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v717
	F_errmsg_internal(m, int32(_a_F_finalize_plan_0), v16)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L14
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_finalize_plan_1), int32(2929), int32(_a_F_finalize_plan_2))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L14
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	v734 = l2
	v735 = v129
	v736 = l4
	v740 = v156
	v741 = v157
	goto L35
L202:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v749 = F_bms_add_members(m, v748, v746)
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L14
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v749
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v741 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v765 = F_bms_add_members(m, v764, v763)
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L14
	} else {
		goto L213
	}
L205:
	;
	v753 = F_bms_union(m, v741, v735)
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L14
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v761 = F_finalize_plan(m, l0, v752, v734, v735, v736)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L14
	} else {
		goto L212
	}
L208:
	;
	v755 = F_finalize_plan(m, l0, v752, v734, v753, v736)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L14
	} else {
		goto L209
	}
L209:
	;
	v757 = F_bms_difference(m, v755, v741)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L14
	} else {
		goto L210
	}
L210:
	;
	F_bms_free(m, v741)
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L14
	} else {
		goto L211
	}
L211:
	;
	v763 = v757
	goto L204
L212:
	;
	v763 = v761
	goto L204
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v765
	if int32(0) <= v740 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v770 = F_bms_del_member(m, v765, v740)
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L14
	} else {
		goto L217
	}
L215:
	;
	v773 = v765
	goto L216
L216:
	;
	v774 = int32(0)
	if v773 == v774 {
		goto L219
	} else {
		goto L220
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v770
	v773 = v770
	goto L216
L218:
	;
	if v827 == int32(0) {
		goto L1
	} else {
		goto L232
	}
L219:
	;
	v827 = int32(1)
	goto L218
L220:
	;
	goto L221
L221:
	;
	if v735 == int32(0) {
		v820 = v774
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v827 = v820
	goto L218
L223:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v773)+4))
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	if v784 < v783 {
		v820 = v774
		goto L222
	} else {
		goto L224
	}
L224:
	;
	v786 = int32(1)
	if v783 <= v786 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v789 = v786
	goto L227
L226:
	;
	v789 = v783
	goto L227
L227:
	;
	v790 = int32(8)
	v795 = int32(0)
	goto L228
L228:
	;
	v802 = v795 << (uint(int32(2)) % 32)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v773+v790+v802)))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v735+v790+v802)))
	v809 = v804 & (v806 ^ int32(-1))
	v811 = base.B2i32(v809 == int32(0))
	if v809 != 0 {
		v820 = v811
		goto L222
	} else {
		goto L230
	}
L229:
	;
	v820 = v811
	goto L222
L230:
	;
	v813 = v795 + int32(1)
	if v813 != v789 {
		v795 = v813
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v831 = F_bms_union(m, v830, v138)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L14
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v831
	v834 = F_bms_add_members(m, v831, v137)
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L14
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v834
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v838 = F_bms_union(m, v837, v138)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L14
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v838
	v841 = F_bms_del_members(m, v838, v137)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L14
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v841
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v859 = v844
	goto L6
L237:
	;
	F_errmsg_internal(m, int32(_a_F_finalize_plan_3), int32(0))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L14
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_finalize_plan_1), int32(2468), int32(_a_F_finalize_plan_2))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L14
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v339
	F_errmsg_internal(m, int32(_a_F_finalize_plan_4), v16+int32(16))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L14
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_finalize_plan_1), int32(2636), int32(_a_F_finalize_plan_2))
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L14
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errmsg_internal(m, int32(_a_F_finalize_plan_5), int32(0))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L14
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_finalize_plan_1), int32(2978), int32(_a_F_finalize_plan_2))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L14
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_among(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v4 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = l2
	v26 = v4
	v27 = v4
	v29 = v4
	v31 = v4
	goto L1
L1:
	;
	v40 = v24
	v43 = v27
	v45 = v29
	v47 = v31
	goto L3
L2:
	;
	v145 = v117
	goto L28
L3:
	;
	if v45 < v47 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v134 = base.B2i32(v114 != v117) & base.B2i32(v117 <= int32(0))
	if v134 != 0 {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	if int32(1) < v114-v117 {
		v40 = v114
		v43 = v117
		v45 = v119
		v47 = v121
		goto L3
	} else {
		goto L19
	}
L6:
	;
	v114 = v40
	v117 = v59
	v119 = v100
	v121 = v47
	goto L5
L7:
	;
	v55 = v45
	goto L9
L8:
	;
	v55 = v47
	goto L9
L9:
	;
	v59 = (v40-v43)>>(uint(int32(1))%32) + v43
	v62 = l1 + v59*int32(20)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 <= v55 {
		v100 = v55
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v69 = v55
	goto L13
L11:
	;
	v114 = v59
	v117 = v43
	v119 = v45
	v121 = v95
	goto L5
L12:
	;
	if int32(0) <= v88 {
		v100 = v69
		goto L6
	} else {
		goto L18
	}
L13:
	;
	if v17 == v69+v18 {
		v95 = v17 - v18
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v114 = v40
	v117 = v59
	v119 = v63
	v121 = v47
	goto L5
L15:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69+(v20+v18)))))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v69))))
	v88 = v84 - v87
	if v88 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v90 = v69 + int32(1)
	if v90 != v63 {
		v69 = v90
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	v95 = v69
	goto L11
L19:
	;
	goto L4
L20:
	;
	if v134 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	goto L2
L23:
	;
	v136 = int32(1)
	goto L25
L24:
	;
	v136 = v26
	goto L25
L25:
	;
	if v26 == int32(0) {
		v24 = v114
		v26 = v136
		v27 = v117
		v29 = v119
		v31 = v121
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	return v179
L28:
	;
	v158 = l1 + v145*int32(20)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 <= v119 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v179 = v178
	goto L27
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v159 + v18
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	if v163 == int32(0) {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v174 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	if v174 <= v175 {
		v145 = v175
		goto L28
	} else {
		goto L38
	}
L34:
	;
	v166 = m.T0[v163].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	return int32(0)
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 + v18
	if v166 != 0 {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v179 = v174
	goto L27
}
func F_find_composite_type_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v99 int32
	_ = v99
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v12 = m.G0
	v14 = v12 - int32(160)
	m.G0 = v14
	F_check_stack_depth(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v14 - int32(-64)
	F_ScanKeyInit(m, v23, int32(4), int32(3), int32(184), int32(1247))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v14+int32(112), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = F_systable_beginscan(m, v20, int32(2674), int32(1), int32(0), int32(2), v23)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L60
	}
L7:
	;
	v43 = F_systable_getnext(m, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = v43
	goto L12
L10:
	;
	goto L11
L11:
	;
	F_systable_endscan(m, v41)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L58
	}
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	v58 = v56 + v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v61 = v59 - int32(1247)
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	goto L11
L14:
	;
	v221 = F_systable_getnext(m, v41)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L56
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v69 = F_relation_open(m, v67, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	if v61 == int32(12) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	F_find_composite_type_dependencies(m, v64, l1, l2)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	goto L15
L20:
	;
	goto L14
L22:
	;
	goto L14
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+52))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v74 = int32(0)
	if base.B2i32(v73 <= v74)|base.B2i32(v72 < v73) == v74 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_relation_close(m, v69, int32(1))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L55
	}
L25:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+119)))
	switch v132 - int32(73) {
	case 0, 10, 32, 36, 39, 41, 43:
		goto L40
	default:
		goto L39
	}
L26:
	;
	v125 = v71 + v72<<(uint(int32(4))%32) + v73*int32(100) - int32(80)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if v72 <= int32(0) {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	v99 = int32(1)
	goto L30
L30:
	;
	v109 = v71 + v72<<(uint(int32(4))%32) - int32(80) + v99*int32(100)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+68))
	if l0 == v110 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v109 == int32(0) {
		goto L24
	} else {
		goto L38
	}
L32:
	;
	goto L31
L33:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+91)))
	if v112 != int32(1) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v116 = v99 + int32(1)
	if v116 <= v72 {
		v99 = v116
		goto L30
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	goto L24
L38:
	;
	v125 = v109
	goto L25
L39:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v131)+72))
	if v191 == int32(0) {
		goto L24
	} else {
		goto L53
	}
L40:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+119)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v144 = int32(4)
	v145 = v125 + v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v148 = v146 + v144
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v151 = v149 + v144
	switch v136 - int32(99) {
	case 0:
		goto L46
	default:
		goto L44
	case 3:
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v151
	F_errmsg(m, int32(_a_F_find_composite_type_dependencies_0), v14)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L51
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v151
	F_errmsg(m, int32(_a_F_find_composite_type_dependencies_1), v14+int32(32))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L49
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v151
	F_errmsg(m, int32(_a_F_find_composite_type_dependencies_2), v14+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_find_composite_type_dependencies_3), int32(_a_F_find_composite_type_dependencies_4), int32(_a_F_find_composite_type_dependencies_5))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errfinish(m, int32(_a_F_find_composite_type_dependencies_3), int32(_a_F_find_composite_type_dependencies_6), int32(_a_F_find_composite_type_dependencies_5))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errfinish(m, int32(_a_F_find_composite_type_dependencies_3), int32(_a_F_find_composite_type_dependencies_7), int32(_a_F_find_composite_type_dependencies_5))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_find_composite_type_dependencies(m, v191, l1, l2)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L24
L55:
	;
	goto L14
L56:
	;
	if v221 != 0 {
		v48 = v221
		goto L12
	} else {
		goto L57
	}
L57:
	;
	goto L13
L58:
	;
	F_relation_close(m, v20, int32(1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	m.G0 = v14 + int32(160)
	return
L60:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v69)+48))
	v250 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+56)) = v125 + v250
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v249 + v250
	F_errmsg(m, int32(_a_F_find_composite_type_dependencies_2), v14+int32(48))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_find_composite_type_dependencies_3), int32(_a_F_find_composite_type_dependencies_8), int32(_a_F_find_composite_type_dependencies_5))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_computable_ec_member(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	v14 = F_pull_var_clause(m, l2, int32(85))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = l3
	goto L8
L7:
	;
	v25 = int32(0)
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v30 = int32(-1)
	v33 = v26
	v36 = v18
	goto L9
L9:
	;
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return v138
L11:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v138 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v127 = v30
	v130 = v33
	v133 = v36
	v137 = v40
	goto L11
L13:
	;
	goto L14
L14:
	;
	v43 = v30
	goto L15
L15:
	;
	if v25 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v121)+12))
	v127 = v108
	v130 = v124
	v133 = v121
	v137 = v124
	goto L11
L17:
	;
	if v108 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v108 = base.I32_ctz(v94) | v95<<(uint(int32(5))%32)
	goto L17
L19:
	;
	v108 = int32(-2)
	goto L17
L20:
	;
	v59 = v43 + int32(1)
	v61 = base.I32_div_s(v59, int32(32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v62 <= v61 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v65 = v25 + int32(8)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v61<<(uint(int32(2))%32))))
	v72 = v69 & (int32(-1) << (uint(v59) % 32))
	if v72 != 0 {
		v94 = v72
		v95 = v61
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v74 = v61 + int32(1)
	if v74 == v62 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v77 = v74
	goto L24
L24:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v65+v77<<(uint(int32(2))%32))))
	if v84 != 0 {
		v94 = v84
		v95 = v77
		goto L18
	} else {
		goto L26
	}
L25:
	;
	goto L19
L26:
	;
	v86 = v77 + int32(1)
	if v86 != v62 {
		v77 = v86
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(0)
L29:
	;
	goto L30
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v113 <= v108 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	goto L33
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v108<<(uint(int32(2))%32))))
	if v121 == int32(0) {
		v43 = v108
		goto L15
	} else {
		goto L34
	}
L34:
	;
	goto L16
L35:
	;
	return int32(0)
L36:
	;
	goto L37
L37:
	;
	v144 = v130 + int32(4)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if base.Ui32(v144) < base.Ui32(v137+v146<<(uint(int32(2))%32)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v151 = v144
	goto L40
L39:
	;
	v151 = int32(0)
	goto L40
L40:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+12)))
	if v152 != 0 {
		v30 = v127
		v33 = v151
		v36 = v133
		goto L9
	} else {
		goto L41
	}
L41:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+13)))
	if v153 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	v157 = int32(0)
	if v156 == v157 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L44
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v215 = F_pull_var_clause(m, v213, int32(21))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L61
	}
L45:
	;
	if v210 == int32(0) {
		v30 = v127
		v33 = v151
		v36 = v133
		goto L9
	} else {
		goto L59
	}
L46:
	;
	v210 = int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	if l3 == int32(0) {
		v203 = v157
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v210 = v203
	goto L45
L50:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v167 < v166 {
		v203 = v157
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v169 = int32(1)
	if v166 <= v169 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v169
	goto L54
L53:
	;
	v172 = v166
	goto L54
L54:
	;
	v173 = int32(8)
	v178 = int32(0)
	goto L55
L55:
	;
	v185 = v178 << (uint(int32(2)) % 32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v156+v173+v185)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3+v173+v185)))
	v192 = v187 & (v189 ^ int32(-1))
	v194 = base.B2i32(v192 == int32(0))
	if v192 != 0 {
		v203 = v194
		goto L49
	} else {
		goto L57
	}
L56:
	;
	v203 = v194
	goto L49
L57:
	;
	v196 = v178 + int32(1)
	if v196 != v172 {
		v178 = v196
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	goto L44
L60:
	;
	F_list_free(m, v215)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L72
	}
L61:
	;
	if v215 == int32(0) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v219 = int32(0)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v220 <= v219 {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v232 = v219
	goto L64
L64:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v232<<(uint(int32(2))%32))))
	v240 = F_list_member(m, v14, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	F_list_free(m, v215)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L71
	}
L66:
	;
	if v240 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v243 = v232 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v243 < v244 {
		v232 = v243
		goto L64
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	goto L65
L70:
	;
	goto L60
L71:
	;
	v30 = v127
	v33 = v151
	v36 = v133
	goto L9
L72:
	;
	if l4 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v263 = F_is_parallel_safe(m, l0, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	goto L10
L76:
	;
	if v263 == int32(0) {
		v30 = v127
		v33 = v151
		v36 = v133
		goto L9
	} else {
		goto L77
	}
L77:
	;
	goto L75
}
func F_find_dependent_phvs_in_jointree(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+68))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v261
L2:
	;
	v12 = F_bms_make_singleton(m, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v261 = int32(0)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	if l1 == v16 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v105 = int32(0)
	v108 = F_get_relids_in_jointree(m, l1, v105, v105)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L34
	}
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v21 == int32(319) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v90 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v90
	v97 = F_query_tree_walker_impl(m, l1, int32(853), v8+int32(8), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L32
	}
L10:
	;
	v85 = F_expression_tree_walker_impl(m, l1, int32(853), v8+int32(8))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L30
	}
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v24 != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v79 = v21
	goto L13
L13:
	;
	if v79 == int32(67) {
		goto L9
	} else {
		goto L29
	}
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v26 = int32(0)
	if base.B2i32(v12 == v26)|base.B2i32(v25 == v26) != 0 {
		v72 = base.B2i32(v12|v25 == v26)
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v72 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v40 != v41 {
		v72 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v43 = int32(1)
	if v40 <= v43 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v46 = v43
	goto L21
L20:
	;
	v46 = v40
	goto L21
L21:
	;
	v47 = int32(8)
	v52 = int32(0)
	goto L22
L22:
	;
	v60 = v52 << (uint(int32(2)) % 32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12+v47+v60)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v25+v47+v60)))
	v65 = base.B2i32(v62 == v64)
	if v62 != v64 {
		v72 = v65
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v72 = v65
	goto L16
L24:
	;
	v68 = v52 + int32(1)
	if v68 != v46 {
		v52 = v68
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v261 = int32(1)
	goto L1
L27:
	;
	goto L28
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v79 = v78
	goto L13
L29:
	;
	goto L10
L30:
	;
	if v85 == int32(0) {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	v261 = int32(1)
	goto L1
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v99 - int32(1)
	if v97 != 0 {
		v261 = v90
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L7
L34:
	;
	if v108 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v166 < int32(0) {
		v261 = v105
		goto L1
	} else {
		goto L46
	}
L36:
	;
	v166 = base.I32_ctz(v152) | v153<<(uint(int32(5))%32)
	goto L35
L37:
	;
	v166 = int32(-2)
	goto L35
L38:
	;
	v119 = base.I32_div_s(int32(0), int32(32))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v120 <= v119 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v123 = v108 + int32(8)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v119<<(uint(int32(2))%32))))
	v130 = v127 & int32(-1)
	if v130 != 0 {
		v152 = v130
		v153 = v119
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v132 = v119 + int32(1)
	if v132 == v120 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v135 = v132
	goto L42
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v123+v135<<(uint(int32(2))%32))))
	if v142 != 0 {
		v152 = v142
		v153 = v135
		goto L36
	} else {
		goto L44
	}
L43:
	;
	goto L37
L44:
	;
	v144 = v135 + int32(1)
	if v144 != v120 {
		v135 = v144
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v170 = v166
	goto L47
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+52))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v176+v170<<(uint(int32(2))%32)-int32(4))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+124)))
	if v183 != int32(1) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L4
L49:
	;
	if v108 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v190 = F_range_table_entry_walker_impl(m, v182, int32(853), v8+int32(8), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	if v190 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v261 = int32(1)
	goto L1
L53:
	;
	if int32(0) <= v250 {
		v170 = v250
		goto L47
	} else {
		goto L64
	}
L54:
	;
	v250 = base.I32_ctz(v236) | v237<<(uint(int32(5))%32)
	goto L53
L55:
	;
	v250 = int32(-2)
	goto L53
L56:
	;
	v201 = v170 + int32(1)
	v203 = base.I32_div_s(v201, int32(32))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v204 <= v203 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v207 = v108 + int32(8)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v203<<(uint(int32(2))%32))))
	v214 = v211 & (int32(-1) << (uint(v201) % 32))
	if v214 != 0 {
		v236 = v214
		v237 = v203
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v216 = v203 + int32(1)
	if v216 == v204 {
		goto L55
	} else {
		goto L59
	}
L59:
	;
	v219 = v216
	goto L60
L60:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v207+v219<<(uint(int32(2))%32))))
	if v226 != 0 {
		v236 = v226
		v237 = v219
		goto L54
	} else {
		goto L62
	}
L61:
	;
	goto L55
L62:
	;
	v228 = v219 + int32(1)
	if v228 != v204 {
		v219 = v228
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	goto L48
}
func F_find_dependent_phvs_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 == int32(319) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v89 = F_expression_tree_walker_impl(m, l0, int32(853), l1)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L24
	} else {
		goto L26
	}
L5:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 != v12 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	v70 = v8
	goto L7
L7:
	;
	if v70 != int32(67) {
		goto L4
	} else {
		goto L23
	}
L8:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v16 = int32(0)
	if base.B2i32(v14 == v16)|base.B2i32(v15 == v16) != 0 {
		v62 = base.B2i32(v14|v15 == v16)
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v62 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	goto L9
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v30 != v31 {
		v62 = int32(0)
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v33 = int32(1)
	if v30 <= v33 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v36 = v33
	goto L15
L14:
	;
	v36 = v30
	goto L15
L15:
	;
	v37 = int32(8)
	v42 = int32(0)
	goto L16
L16:
	;
	v50 = v42 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14+v37+v50)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v15+v37+v50)))
	v55 = base.B2i32(v52 == v54)
	if v52 != v54 {
		v62 = v55
		goto L10
	} else {
		goto L18
	}
L17:
	;
	v62 = v55
	goto L10
L18:
	;
	v58 = v42 + int32(1)
	if v58 != v36 {
		v42 = v58
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(1)
L21:
	;
	goto L22
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v70 = v69
	goto L7
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v73 + int32(1)
	v79 = F_query_tree_walker_impl(m, l0, int32(853), l1, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v83 - int32(1)
	return v79
L26:
	;
	return v89
}
func F_find_my_exec(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
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
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	goto L4
L1:
	;
	v130 = Fn13880(m, l1, int32(47))
	mBase = m.M
	goto L34
L2:
	;
	v126 = F_strlen(m, v115)
	mBase = m.M
	goto L1
L4:
	;
	goto L5
L5:
	;
	v16 = int32(1023)
	if (l1^l0)&int32(3) != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v119)
	goto L2
L7:
	;
	v100 = v95
	v101 = v96
	v102 = v97
	goto L28
L8:
	;
	if v90 == int32(0) {
		v115 = v88
		v116 = v89
		goto L6
	} else {
		goto L27
	}
L9:
	;
	v88 = l0
	v89 = l1
	v90 = v16
	goto L8
L10:
	;
	goto L11
L11:
	;
	v20 = int32(0)
	if base.B2i32(l0&int32(3) == v20)|int32(0) == v20 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v56 == int32(0) {
		v115 = v53
		v116 = v54
		goto L6
	} else {
		goto L21
	}
L13:
	;
	v32 = l0
	v33 = l1
	v34 = v16
	goto L16
L14:
	;
	goto L15
L15:
	;
	v53 = l0
	v54 = l1
	v55 = v16
	v56 = int32(1)
	goto L12
L16:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v36)
	if v36 == int32(0) {
		v95 = v32
		v96 = v33
		v97 = v34
		goto L7
	} else {
		goto L18
	}
L17:
	;
	v53 = v47
	v54 = v41
	v55 = v43
	v56 = v45
	goto L12
L18:
	;
	v40 = int32(1)
	v41 = v33 + v40
	v43 = v34 - v40
	v44 = int32(0)
	v45 = base.B2i32(v43 != v44)
	v47 = v32 + v40
	if v47&int32(3) == v44 {
		v53 = v47
		v54 = v41
		v55 = v43
		v56 = v45
		goto L12
	} else {
		goto L19
	}
L19:
	;
	if v43 != 0 {
		v32 = v47
		v33 = v41
		v34 = v43
		goto L16
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if base.B2i32(v59 == int32(0))|base.B2i32(base.Ui32(v55) < base.Ui32(int32(4))) != 0 {
		v88 = v53
		v89 = v54
		v90 = v55
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v66 = v53
	v67 = v54
	v68 = v55
	goto L23
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v74 = int32(-2139062144)
	if (int32(16843008)-v71|v71)&v74 != v74 {
		v95 = v66
		v96 = v67
		v97 = v68
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v88 = v82
	v89 = v80
	v90 = v84
	goto L8
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = v71
	v79 = int32(4)
	v80 = v67 + v79
	v82 = v66 + v79
	v84 = v68 - v79
	if base.Ui32(int32(3)) < base.Ui32(v84) {
		v66 = v82
		v67 = v80
		v68 = v84
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v95 = v88
	v96 = v89
	v97 = v90
	goto L7
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v101))) = uint8(v104)
	if v104 == int32(0) {
		v115 = v100
		v116 = v101
		goto L6
	} else {
		goto L30
	}
L29:
	;
	v115 = v111
	v116 = v109
	goto L6
L30:
	;
	v108 = int32(1)
	v109 = v101 + v108
	v111 = v100 + v108
	v113 = v102 - v108
	if v113 != 0 {
		v100 = v111
		v101 = v109
		v102 = v113
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	m.G0 = v8 + int32(160)
	return v643
L33:
	;
	v618 = int32(-1)
	v621 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L47
	} else {
		goto L189
	}
L34:
	;
	if v130 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = F___fstatat(m, int32(-100), l1, v8-int32(-64), int32(0))
	mBase = m.M
	goto L38
L36:
	;
	goto L37
L37:
	;
	v159 = int32(_a_F_find_my_exec_0)
	v160 = int32(0)
	v165 = F___strchrnul(m, v159, int32(61))
	mBase = m.M
	if v159 == v165 {
		goto L52
	} else {
		goto L53
	}
L38:
	;
	if v135 < int32(0) {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	v140 = v138 & int32(_a_F_find_my_exec_1)
	if v140 != int32(_a_F_find_my_exec_2) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v140 == int32(_a_F_find_my_exec_3) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v151 = F_access(m, l1, int32(4))
	mBase = m.M
	v153 = F_access(m, l1, int32(1))
	mBase = m.M
	if v153|v151 != 0 {
		goto L33
	} else {
		goto L46
	}
L43:
	;
	v148 = int32(31)
	goto L45
L44:
	;
	v148 = int32(63)
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_my_exec[0])) = v148
	goto L33
L46:
	;
	v155 = F_normalize_exec_path(m, l1)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return int32(0)
L48:
	;
	v643 = v155
	goto L32
L49:
	;
	v615 = F_normalize_exec_path(m, l1)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L47
	} else {
		goto L188
	}
L50:
	;
	v591 = int32(-1)
	v594 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L47
	} else {
		goto L183
	}
L51:
	;
	if v207 == int32(0) {
		goto L50
	} else {
		goto L67
	}
L52:
	;
	v207 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v168 = v165 - v159
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+uint32(_c_F_find_my_exec[1]))))
	if v170 != 0 {
		v201 = v160
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v207 = v201
	goto L51
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_find_my_exec[2]))
	if v172 == int32(0) {
		v201 = v160
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v175 == int32(0) {
		v201 = v160
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v179 = v172
	v180 = v175
	goto L59
L59:
	;
	v183 = F_strncmp(m, v159, v180, v168)
	mBase = m.M
	if v183 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v201 = v187 + int32(1)
	goto L55
L61:
	;
	goto L60
L62:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v187 = v186 + v168
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if v188 == int32(61) {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v192 != 0 {
		v179 = v179 + int32(4)
		v180 = v192
		goto L59
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v201 = v160
	goto L55
L67:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
	if v210 == int32(0) {
		goto L50
	} else {
		goto L68
	}
L68:
	;
	v214 = Fn13880(m, v207, int32(58))
	mBase = m.M
	goto L69
L69:
	;
	if v214 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v217 = F_strlen(m, v207)
	mBase = m.M
	v219 = v217 + v207
	goto L72
L71:
	;
	v219 = v214
	goto L72
L72:
	;
	v220 = int32(1024)
	v223 = v219 - v207 + int32(1)
	if v220 <= v223 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v226 = v220
	goto L75
L74:
	;
	v226 = v223
	goto L75
L75:
	;
	if v226 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	F_join_path_components(m, l1, l1, l0)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L47
	} else {
		goto L107
	}
L77:
	;
	v342 = F_strlen(m, v338)
	mBase = m.M
	goto L76
L78:
	;
	v338 = v207
	goto L77
L79:
	;
	goto L80
L80:
	;
	v232 = v226 - int32(1)
	if (l1^v207)&int32(3) != 0 {
		goto L84
	} else {
		goto L85
	}
L81:
	;
	v335 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v332))) = uint8(v335)
	v338 = v331
	goto L77
L82:
	;
	v316 = v311
	v317 = v312
	v318 = v313
	goto L103
L83:
	;
	if v306 == int32(0) {
		v331 = v304
		v332 = v305
		goto L81
	} else {
		goto L102
	}
L84:
	;
	v304 = v207
	v305 = l1
	v306 = v232
	goto L83
L85:
	;
	goto L86
L86:
	;
	v236 = int32(0)
	if base.B2i32(v207&int32(3) == v236)|base.B2i32(v232 == v236) == v236 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v272 == int32(0) {
		v331 = v269
		v332 = v270
		goto L81
	} else {
		goto L96
	}
L88:
	;
	v248 = v207
	v249 = l1
	v250 = v232
	goto L91
L89:
	;
	goto L90
L90:
	;
	v269 = v207
	v270 = l1
	v271 = v232
	v272 = base.B2i32(v232 != v236)
	goto L87
L91:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	*(*uint8)(unsafe.Add(mBase, uint32(v249))) = uint8(v252)
	if v252 == int32(0) {
		v311 = v248
		v312 = v249
		v313 = v250
		goto L82
	} else {
		goto L93
	}
L92:
	;
	v269 = v263
	v270 = v257
	v271 = v259
	v272 = v261
	goto L87
L93:
	;
	v256 = int32(1)
	v257 = v249 + v256
	v259 = v250 - v256
	v260 = int32(0)
	v261 = base.B2i32(v259 != v260)
	v263 = v248 + v256
	if v263&int32(3) == v260 {
		v269 = v263
		v270 = v257
		v271 = v259
		v272 = v261
		goto L87
	} else {
		goto L94
	}
L94:
	;
	if v259 != 0 {
		v248 = v263
		v249 = v257
		v250 = v259
		goto L91
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if base.B2i32(v275 == int32(0))|base.B2i32(base.Ui32(v271) < base.Ui32(int32(4))) != 0 {
		v304 = v269
		v305 = v270
		v306 = v271
		goto L83
	} else {
		goto L97
	}
L97:
	;
	v282 = v269
	v283 = v270
	v284 = v271
	goto L98
L98:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v290 = int32(-2139062144)
	if (int32(16843008)-v287|v287)&v290 != v290 {
		v311 = v282
		v312 = v283
		v313 = v284
		goto L82
	} else {
		goto L100
	}
L99:
	;
	v304 = v298
	v305 = v296
	v306 = v300
	goto L83
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v287
	v295 = int32(4)
	v296 = v283 + v295
	v298 = v282 + v295
	v300 = v284 - v295
	if base.Ui32(int32(3)) < base.Ui32(v300) {
		v282 = v298
		v283 = v296
		v284 = v300
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	v311 = v304
	v312 = v305
	v313 = v306
	goto L82
L103:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316))))
	*(*uint8)(unsafe.Add(mBase, uint32(v317))) = uint8(v320)
	if v320 == int32(0) {
		v331 = v316
		v332 = v317
		goto L81
	} else {
		goto L105
	}
L104:
	;
	v331 = v327
	v332 = v325
	goto L81
L105:
	;
	v324 = int32(1)
	v325 = v317 + v324
	v327 = v316 + v324
	v329 = v318 - v324
	if v329 != 0 {
		v316 = v327
		v317 = v325
		v318 = v329
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	F_canonicalize_path_enc(m, l1)
	mBase = m.M
	v352 = F___fstatat(m, int32(-100), l1, v8-int32(-64), int32(0))
	mBase = m.M
	goto L109
L108:
	;
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if v394 == int32(0) {
		goto L50
	} else {
		goto L124
	}
L109:
	;
	if v352 < int32(0) {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	v357 = v355 & int32(_a_F_find_my_exec_1)
	if v357 != int32(_a_F_find_my_exec_2) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	if v357 == int32(_a_F_find_my_exec_3) {
		goto L114
	} else {
		goto L115
	}
L112:
	;
	goto L113
L113:
	;
	v368 = F_access(m, l1, int32(4))
	mBase = m.M
	v370 = F_access(m, l1, int32(1))
	mBase = m.M
	if v370 != 0 {
		goto L108
	} else {
		goto L117
	}
L114:
	;
	v365 = int32(31)
	goto L116
L115:
	;
	v365 = int32(63)
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_my_exec[0])) = v365
	goto L108
L117:
	;
	if v368 == int32(0) {
		goto L49
	} else {
		goto L118
	}
L118:
	;
	v375 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L47
	} else {
		goto L119
	}
L119:
	;
	if v375 == int32(0) {
		goto L108
	} else {
		goto L120
	}
L120:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L47
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = l1
	F_errmsg_internal(m, int32(_a_F_find_my_exec_4), v8+int32(32))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L47
	} else {
		goto L122
	}
L122:
	;
	F_errfinish(m, int32(_a_F_find_my_exec_5), int32(218), int32(_a_F_find_my_exec_6))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L47
	} else {
		goto L123
	}
L123:
	;
	goto L108
L124:
	;
	v401 = v219
	goto L125
L125:
	;
	v403 = v401 + int32(1)
	v405 = Fn13880(m, v403, int32(58))
	mBase = m.M
	goto L127
L126:
	;
	goto L50
L127:
	;
	if v405 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v408 = F_strlen(m, v403)
	mBase = m.M
	v410 = v408 + v403
	goto L130
L129:
	;
	v410 = v405
	goto L130
L130:
	;
	v411 = int32(1024)
	v414 = v410 - v403 + int32(1)
	if v411 <= v414 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v417 = v411
	goto L133
L132:
	;
	v417 = v414
	goto L133
L133:
	;
	if v417 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L134:
	;
	F_join_path_components(m, l1, l1, l0)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L47
	} else {
		goto L165
	}
L135:
	;
	v533 = F_strlen(m, v529)
	mBase = m.M
	goto L134
L136:
	;
	v529 = v403
	goto L135
L137:
	;
	goto L138
L138:
	;
	v423 = v417 - int32(1)
	if (l1^v403)&int32(3) != 0 {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v526 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v523))) = uint8(v526)
	v529 = v522
	goto L135
L140:
	;
	v507 = v502
	v508 = v503
	v509 = v504
	goto L161
L141:
	;
	if v497 == int32(0) {
		v522 = v495
		v523 = v496
		goto L139
	} else {
		goto L160
	}
L142:
	;
	v495 = v403
	v496 = l1
	v497 = v423
	goto L141
L143:
	;
	goto L144
L144:
	;
	v427 = int32(0)
	if base.B2i32(v403&int32(3) == v427)|base.B2i32(v423 == v427) == v427 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v463 == int32(0) {
		v522 = v460
		v523 = v461
		goto L139
	} else {
		goto L154
	}
L146:
	;
	v439 = v403
	v440 = l1
	v441 = v423
	goto L149
L147:
	;
	goto L148
L148:
	;
	v460 = v403
	v461 = l1
	v462 = v423
	v463 = base.B2i32(v423 != v427)
	goto L145
L149:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439))))
	*(*uint8)(unsafe.Add(mBase, uint32(v440))) = uint8(v443)
	if v443 == int32(0) {
		v502 = v439
		v503 = v440
		v504 = v441
		goto L140
	} else {
		goto L151
	}
L150:
	;
	v460 = v454
	v461 = v448
	v462 = v450
	v463 = v452
	goto L145
L151:
	;
	v447 = int32(1)
	v448 = v440 + v447
	v450 = v441 - v447
	v451 = int32(0)
	v452 = base.B2i32(v450 != v451)
	v454 = v439 + v447
	if v454&int32(3) == v451 {
		v460 = v454
		v461 = v448
		v462 = v450
		v463 = v452
		goto L145
	} else {
		goto L152
	}
L152:
	;
	if v450 != 0 {
		v439 = v454
		v440 = v448
		v441 = v450
		goto L149
	} else {
		goto L153
	}
L153:
	;
	goto L150
L154:
	;
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if base.B2i32(v466 == int32(0))|base.B2i32(base.Ui32(v462) < base.Ui32(int32(4))) != 0 {
		v495 = v460
		v496 = v461
		v497 = v462
		goto L141
	} else {
		goto L155
	}
L155:
	;
	v473 = v460
	v474 = v461
	v475 = v462
	goto L156
L156:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	v481 = int32(-2139062144)
	if (int32(16843008)-v478|v478)&v481 != v481 {
		v502 = v473
		v503 = v474
		v504 = v475
		goto L140
	} else {
		goto L158
	}
L157:
	;
	v495 = v489
	v496 = v487
	v497 = v491
	goto L141
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = v478
	v486 = int32(4)
	v487 = v474 + v486
	v489 = v473 + v486
	v491 = v475 - v486
	if base.Ui32(int32(3)) < base.Ui32(v491) {
		v473 = v489
		v474 = v487
		v475 = v491
		goto L156
	} else {
		goto L159
	}
L159:
	;
	goto L157
L160:
	;
	v502 = v495
	v503 = v496
	v504 = v497
	goto L140
L161:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507))))
	*(*uint8)(unsafe.Add(mBase, uint32(v508))) = uint8(v511)
	if v511 == int32(0) {
		v522 = v507
		v523 = v508
		goto L139
	} else {
		goto L163
	}
L162:
	;
	v522 = v518
	v523 = v516
	goto L139
L163:
	;
	v515 = int32(1)
	v516 = v508 + v515
	v518 = v507 + v515
	v520 = v509 - v515
	if v520 != 0 {
		v507 = v518
		v508 = v516
		v509 = v520
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	F_canonicalize_path_enc(m, l1)
	mBase = m.M
	v543 = F___fstatat(m, int32(-100), l1, v8-int32(-64), int32(0))
	mBase = m.M
	goto L167
L166:
	;
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v585 != 0 {
		v401 = v410
		goto L125
	} else {
		goto L182
	}
L167:
	;
	if v543 < int32(0) {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v8)+68))
	v548 = v546 & int32(_a_F_find_my_exec_1)
	if v548 != int32(_a_F_find_my_exec_2) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	if v548 == int32(_a_F_find_my_exec_3) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v559 = F_access(m, l1, int32(4))
	mBase = m.M
	v561 = F_access(m, l1, int32(1))
	mBase = m.M
	if v561 != 0 {
		goto L166
	} else {
		goto L175
	}
L172:
	;
	v556 = int32(31)
	goto L174
L173:
	;
	v556 = int32(63)
	goto L174
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_my_exec[0])) = v556
	goto L166
L175:
	;
	if v559 == int32(0) {
		goto L49
	} else {
		goto L176
	}
L176:
	;
	v566 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L47
	} else {
		goto L177
	}
L177:
	;
	if v566 == int32(0) {
		goto L166
	} else {
		goto L178
	}
L178:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L47
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_find_my_exec_4), v8+int32(16))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L47
	} else {
		goto L180
	}
L180:
	;
	F_errfinish(m, int32(_a_F_find_my_exec_5), int32(218), int32(_a_F_find_my_exec_6))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L47
	} else {
		goto L181
	}
L181:
	;
	goto L166
L182:
	;
	goto L126
L183:
	;
	if v594 == int32(0) {
		v643 = v591
		goto L32
	} else {
		goto L184
	}
L184:
	;
	F_errcode(m, int32(16908805))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L47
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_find_my_exec_7), v8)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L47
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_find_my_exec_5), int32(225), int32(_a_F_find_my_exec_6))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L47
	} else {
		goto L187
	}
L187:
	;
	v643 = v591
	goto L32
L188:
	;
	v643 = v615
	goto L32
L189:
	;
	if v621 == int32(0) {
		v643 = v618
		goto L32
	} else {
		goto L190
	}
L190:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L47
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = l1
	F_errmsg_internal(m, int32(_a_F_find_my_exec_8), v8+int32(48))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L47
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_find_my_exec_5), int32(174), int32(_a_F_find_my_exec_6))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L47
	} else {
		goto L193
	}
L193:
	;
	v643 = v618
	goto L32
}
func F_findconstraintloop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var __phi32 int32
	_ = __phi32
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v35 int32
	_ = v35
	var __phi35 int32
	_ = __phi35
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v321 int32
	_ = v321
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v15 = m.T0[v14].(func(*base.Module) int32)(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(101)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v27 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v25 = v23
	goto L9
L8:
	;
	v25 = int32(19)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v25
	goto L1
L10:
	;
	return v321
L11:
	;
	if l1 == v27 {
		v321 = int32(0)
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v275 != 0 {
		goto L106
	} else {
		goto L107
	}
L14:
	;
	__phi32 = v27
	__phi33 = int32(0)
	__phi35 = l1
	v32 = __phi32
	v33 = __phi33
	v35 = __phi35
	goto L15
L15:
	;
	if v33 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	__phi32 = v274
	__phi33 = v267
	__phi35 = v32
	v32 = __phi32
	v33 = __phi33
	v35 = __phi35
	goto L15
L18:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v107 != 0 {
		goto L44
	} else {
		goto L45
	}
L19:
	;
	if v82 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L20:
	;
	v90 = int32(0)
	if l1 != v32 {
		v267 = v90
		goto L17
	} else {
		goto L40
	}
L21:
	;
	v42 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	if v44 == v42 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v82 = v33
	goto L23
L23:
	;
	if l1 == v32 {
		goto L19
	} else {
		goto L39
	}
L24:
	;
	v49 = v44
	v51 = v42
	v55 = v42
	goto L25
L25:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v32 != v57 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	if v72 <= int32(1) {
		goto L36
	} else {
		goto L37
	}
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v74 != 0 {
		v49 = v74
		v51 = v70
		v55 = v72
		goto L25
	} else {
		goto L35
	}
L28:
	;
	v69 = v51
	v70 = v51
	v72 = v55
	goto L27
L29:
	;
	goto L30
L30:
	;
	v59 = int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	switch v60 - int32(76) {
	case 0, 18, 21, 38:
		v66 = v49
		v67 = v59
		goto L31
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L32
	default:
		goto L33
	}
L31:
	;
	v69 = v66
	v70 = v66
	v72 = v67 + v55
	goto L27
L32:
	;
	v66 = v51
	v67 = int32(0)
	goto L31
L33:
	;
	if v60 == int32(36) {
		v66 = v49
		v67 = v59
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	goto L26
L36:
	;
	v78 = v69
	goto L38
L37:
	;
	v78 = int32(0)
	goto L38
L38:
	;
	v82 = v78
	goto L23
L39:
	;
	v267 = v82
	goto L17
L40:
	;
	v98 = l1
	v100 = v90
	v103 = v27
	goto L18
L41:
	;
	v98 = l1
	v100 = int32(0)
	v103 = v27
	goto L18
L42:
	;
	goto L43
L43:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	v98 = v96
	v100 = v82
	v103 = v95
	goto L18
L44:
	;
	v110 = v107
	goto L47
L45:
	;
	goto L46
L46:
	;
	v131 = F_newstate(m, l0)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L2
	} else {
		goto L50
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+24)) = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v110)+28))
	if v120 != 0 {
		v110 = v120
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	goto L48
L50:
	;
	if v131 == int32(0) {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v135 = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_clonesuccessorstates(m, l0, v103, v131, v98, v100, v135, v135, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	if v141 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v131)+12))
	if v142 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v162 = v131
	goto L56
L55:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v131)+4)) = uint8(v143)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(-1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)+32))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v131)+28))
	if v148 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v98)+20))
	if v163 == int32(0) {
		goto L1
	} else {
		goto L65
	}
L57:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v131)+28))
	if v147 != 0 {
		goto L62
	} else {
		goto L63
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v147
	goto L57
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v147
	goto L57
L61:
	;
	v154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v131)+32)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+28)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v131
	v162 = v154
	goto L56
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+28)) = v151
	goto L61
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v151
	goto L61
L65:
	;
	v168 = v163
	goto L66
L66:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	if v177 != v103 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L1
L68:
	;
	if v176 != 0 {
		v168 = v176
		goto L66
	} else {
		goto L105
	}
L69:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	switch v179 - int32(76) {
	case 0, 18, 21, 38:
		goto L70
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L68
	default:
		goto L71
	}
L70:
	;
	if v162 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	if v179 != int32(36) {
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	F_cparc(m, l0, v168, v98, v162)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v168)+4)))
	if v192 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L75
L77:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	if v262 != 0 {
		goto L1
	} else {
		goto L103
	}
L78:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v168)+20))
	if v227 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L79:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v197 = v195 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v197))|base.B2i32(int32(1)<<(uint(v197)%32)&int32(_a_F_findconstraintloop_0) == int32(0)) != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v207 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v168)+36))
	if v208 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	if v220 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v168)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v192*int32(24))+12)) = v216
	v220 = v216
	goto L82
L84:
	;
	goto L85
L85:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v168)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v208)+32)) = v218
	v220 = v218
	goto L82
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v220)+36)) = v208
	goto L88
L87:
	;
	goto L88
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v168)+32)) = int64(0)
	goto L78
L89:
	;
	if v226 != 0 {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v191)+20)) = v226
	goto L89
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = v226
	goto L89
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226)+20)) = v227
	goto L95
L94:
	;
	goto L95
L95:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v191)+12)) = v233 - int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v168)+24))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v168)+28))
	if v238 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v237 != 0 {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+16)) = v237
	goto L96
L98:
	;
	goto L99
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238)+24)) = v237
	goto L96
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v237)+28)) = v238
	goto L102
L101:
	;
	goto L102
L102:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v190)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v244 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = int32(0)
	v251 = v168 + int32(8)
	v252 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v251)+16)) = v252
	*(*int64)(unsafe.Add(mBase, uint32(v251)+8)) = v252
	*(*int64)(unsafe.Add(mBase, uint32(v251))) = v252
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v168)+16)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v168
	goto L77
L103:
	;
	if v176 != 0 {
		v168 = v176
		goto L66
	} else {
		goto L104
	}
L104:
	;
	goto L1
L105:
	;
	goto L67
L106:
	;
	v278 = v275
	goto L109
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v321 = int32(0)
	goto L10
L109:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	switch v286 - int32(76) {
	case 0, 18, 21, 38:
		goto L112
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37:
		goto L111
	default:
		goto L113
	}
L110:
	;
	goto L108
L111:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	if v298 != 0 {
		v278 = v298
		goto L109
	} else {
		goto L117
	}
L112:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v278)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v291
	v293 = F_findconstraintloop(m, l0, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L115
	}
L113:
	;
	if v286 != int32(36) {
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	if v293 == int32(0) {
		goto L111
	} else {
		goto L116
	}
L116:
	;
	goto L1
L117:
	;
	goto L110
}
func F_findoprnd_1(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	goto L1
L1:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = l0 + v13*int32(12)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16))))
	switch v17 - int32(2) {
	case 0, 4:
		goto L6
	default:
		goto L5
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13 + int32(1)
	if v25 == int32(33) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v20 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)) = uint16(v20)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13 + int32(1)
	return
L7:
	;
	v31 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)) = uint16(v31)
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_findoprnd_1(m, l0, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = v35 - v13
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+2)) = uint16(v36)
	goto L1
}
func F_findwrd(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v31 int32
	_ = v31
	var __phi31 int32
	_ = __phi31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v4 = int32(0)
	v8 = l0
	goto L1
L1:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.B2i32(base.Ui32(v15-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v15 == int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v64 = F_pg_mblen_cstr(m, v8)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L13
	} else {
		goto L23
	}
L6:
	;
	__phi29 = v8
	__phi30 = v15
	__phi31 = v8
	v29 = __phi29
	v30 = __phi30
	v31 = __phi31
	goto L9
L7:
	;
	v58 = v4
	v59 = v4
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v59
	return v58
L9:
	;
	switch v30 {
	case 0, 9, 10, 11, 12, 13, 32:
		goto L11
	default:
		goto L12
	}
L10:
	;
	if base.B2i32(l2 == int32(0))|base.B2i32(v31-v29 != int32(1)) != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	v32 = F_pg_mblen_cstr(m, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v36 = v32 + v31
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	__phi29 = v31
	__phi30 = v37
	__phi31 = v36
	v29 = __phi29
	v30 = __phi30
	v31 = __phi31
	goto L9
L15:
	;
	v58 = v8
	v59 = v54
	goto L8
L16:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v52)
	v54 = v51
	goto L15
L17:
	;
	if l2 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v44 != int32(42) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v51 = v29
	v52 = int32(2)
	goto L16
L20:
	;
	v54 = v31
	goto L15
L21:
	;
	goto L22
L22:
	;
	v51 = v31
	v52 = int32(0)
	goto L16
L23:
	;
	v8 = v64 + v8
	goto L1
}
func F_finite_interval_mi(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = v11 - v12
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v13
	if base.B2i32(v13 < v11)^base.B2i32(int32(0) < v12) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_finite_interval_mi_0), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_finite_interval_mi_1), int32(3573), int32(_a_F_finite_interval_mi_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v21 = v19 - v20
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
		if base.B2i32(v21 < v19)^base.B2i32(int32(0) < v20) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_finite_interval_mi_0), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_finite_interval_mi_1), int32(3573), int32(_a_F_finite_interval_mi_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v29 = v27 - v28
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v29
			if base.B2i32(v29 < v27)^base.B2i32(int64(0) < v28) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_finite_interval_mi_0), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_finite_interval_mi_1), int32(3573), int32(_a_F_finite_interval_mi_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v13 != int32(2147483647) {
					v37 = int32(-2147483648)
					if base.B2i32(v13 != v37)|base.B2i32(v21 != v37) != 0 {
						return
					} else {
						if v29 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_finite_interval_mi_0), int32(0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_finite_interval_mi_1), int32(3573), int32(_a_F_finite_interval_mi_2))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							return
						}
					}
				} else {
					if base.B2i32(v21 != int32(2147483647))|base.B2i32(v29 != int64(9223372036854775807)) != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_finite_interval_mi_0), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_finite_interval_mi_1), int32(3573), int32(_a_F_finite_interval_mi_2))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_finite_interval_pl(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v13 = v11 + v12
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v13
	if base.B2i32(v12 < int32(0))^base.B2i32(v13 < v11) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			F_errcode(m, int32(134217858))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_finite_interval_pl_0), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_finite_interval_pl_1), int32(3517), int32(_a_F_finite_interval_pl_2))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v21 = v19 + v20
		*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v21
		if base.B2i32(v20 < int32(0))^base.B2i32(v21 < v19) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_finite_interval_pl_0), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_finite_interval_pl_1), int32(3517), int32(_a_F_finite_interval_pl_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			v29 = v27 + v28
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v29
			if base.B2i32(v28 < int64(0))^base.B2i32(v29 < v27) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_finite_interval_pl_0), int32(0))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_finite_interval_pl_1), int32(3517), int32(_a_F_finite_interval_pl_2))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v13 != int32(2147483647) {
					v37 = int32(-2147483648)
					if base.B2i32(v13 != v37)|base.B2i32(v21 != v37) != 0 {
						return
					} else {
						if v29 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_finite_interval_pl_0), int32(0))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_finite_interval_pl_1), int32(3517), int32(_a_F_finite_interval_pl_2))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							return
						}
					}
				} else {
					if base.B2i32(v21 != int32(2147483647))|base.B2i32(v29 != int64(9223372036854775807)) != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_finite_interval_pl_0), int32(0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_finite_interval_pl_1), int32(3517), int32(_a_F_finite_interval_pl_2))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_first_dir_separator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v4 = l0
	goto L2
L1:
	;
	return v14
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L1
L4:
	;
	goto L3
L5:
	;
	v14 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v7 == int32(47) {
		v14 = v4
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v4 = v4 + int32(1)
	goto L2
}
func F_first_path_var_separator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	v4 = l0
	goto L2
L1:
	;
	return v14
L2:
	;
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v7 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L1
L4:
	;
	goto L3
L5:
	;
	v14 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	if v7 == int32(58) {
		v14 = v4
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v4 = v4 + int32(1)
	goto L2
}
func F_fix_windowagg_condition_expr_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	if l0 == int32(0) {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v8 == int32(11) {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = F_tlist_member(m, l0, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				if v14 != 0 {
					v18 = F_makeVarFromTargetEntry(m, v11, v14)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v18)+40)) = uint16(v20)
						*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v20
						return v18
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_fix_windowagg_condition_expr_mutator_0), int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_fix_windowagg_condition_expr_mutator_1), int32(3458), int32(_a_F_fix_windowagg_condition_expr_mutator_2))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		} else {
			v39 = F_expression_tree_mutator_impl(m, l0, int32(841), l1)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				return v39
			}
		}
	}
}
func F_flatten_reloptions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v19 = F_SysCacheGetAttr(m, int32(57), v10, int32(33), v7+int32(31))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+31)))
				if v21 == int32(0) {
					v25 = v7 + int32(12)
					F_initStringInfo(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						F_get_reloptions(m, v25, v19)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							v31 = v30
							F_ReleaseCatCache(m, v10)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(32)
								return v31
							}
						}
					}
				} else {
					v31 = int32(0)
					F_ReleaseCatCache(m, v10)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v31
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
				F_errmsg_internal(m, int32(_a_F_flatten_reloptions_0), v7)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_flatten_reloptions_1), int32(_a_F_flatten_reloptions_2), int32(_a_F_flatten_reloptions_3))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_float4_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 float32
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v51 float64
	_ = v51
	var v58 float64
	_ = v58
	var v65 int32
	_ = v65
	var v66 float64
	_ = v66
	var v77 float64
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		if v20 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v136 = m.ExcPending
			if v136 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float4_accum_0)
				F_errmsg_internal(m, int32(_a_F_float4_accum_1), v13)
				mBase = m.M
				v143 = m.ExcPending
				if v143 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_float4_accum_2), int32(2938), int32(_a_F_float4_accum_3))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v23 != int32(3) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v136 = m.ExcPending
				if v136 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float4_accum_0)
					F_errmsg_internal(m, int32(_a_F_float4_accum_1), v13)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_float4_accum_2), int32(2938), int32(_a_F_float4_accum_3))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
				if v26 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v136 = m.ExcPending
					if v136 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float4_accum_0)
						F_errmsg_internal(m, int32(_a_F_float4_accum_1), v13)
						mBase = m.M
						v143 = m.ExcPending
						if v143 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_float4_accum_2), int32(2938), int32(_a_F_float4_accum_3))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					if v27 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_float4_accum_0)
							F_errmsg_internal(m, int32(_a_F_float4_accum_1), v13)
							mBase = m.M
							v143 = m.ExcPending
							if v143 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_float4_accum_2), int32(2938), int32(_a_F_float4_accum_3))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
						v31 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
						v32 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v16)+40))
						*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v33
						v36 = base.F64_add(v32, float64(1))
						*(*float64)(unsafe.Add(mBase, uint32(v13)+40)) = v36
						v38 = base.F64_promote_f32(v30)
						v39 = base.F64_add(v31, v38)
						*(*float64)(unsafe.Add(mBase, uint32(v13)+32)) = v39
						if base.F64_gt(v32, float64(0)) != 0 {
							v44 = base.F64_sub(base.F64_mul(v38, v36), v39)
							v48 = base.F64_add(v33, base.F64_div(base.F64_mul(v44, v44), base.F64_mul(v32, v36)))
							*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v48
							v51 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v39), v51)&base.F64_ne(base.F64_abs(v48), v51) != 0 {
								v77 = v48
								v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v81 == int32(0) {
									v109 = int32(0)
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
									switch v84 - int32(429) {
									case 0:
										v109 = int32(1)
									case 1:
										v109 = int32(2)
									default:
										v109 = int32(0)
									}
								}
								if v109 != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v77
									*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
									*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
									v128 = v16
									m.G0 = v13 + int32(48)
									return v128
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
									v126 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
									mBase = m.M
									v127 = m.ExcPending
									if v127 != 0 {
										return int32(0)
									} else {
										v128 = v126
										m.G0 = v13 + int32(48)
										return v128
									}
								}
							} else {
								v58 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v31), v58)|base.F64_eq(base.F64_abs(v38), v58) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
									v77 = math.Float64frombits(uint64(0x7ff8000000000000))
									v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v81 == int32(0) {
										v109 = int32(0)
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
										switch v84 - int32(429) {
										case 0:
											v109 = int32(1)
										case 1:
											v109 = int32(2)
										default:
											v109 = int32(0)
										}
									}
									if v109 != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v77
										*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
										*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
										v128 = v16
										m.G0 = v13 + int32(48)
										return v128
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
										*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
										v126 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											v128 = v126
											m.G0 = v13 + int32(48)
											return v128
										}
									}
								} else {
									F_float_overflow_error(m)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v66 = base.F64_abs(v38)
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)) {
								*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
								v77 = math.Float64frombits(uint64(0x7ff8000000000000))
							} else {
								if base.F64_ne(v66, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v77 = v33
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(9221120237041090560)
									v77 = math.Float64frombits(uint64(0x7ff8000000000000))
								}
							}
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v81 == int32(0) {
								v109 = int32(0)
							} else {
								v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
								switch v84 - int32(429) {
								case 0:
									v109 = int32(1)
								case 1:
									v109 = int32(2)
								default:
									v109 = int32(0)
								}
							}
							if v109 != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v16)+40)) = v77
								*(*float64)(unsafe.Add(mBase, uint32(v16)+32)) = v39
								*(*float64)(unsafe.Add(mBase, uint32(v16)+24)) = v36
								v128 = v16
								m.G0 = v13 + int32(48)
								return v128
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v13 + int32(24)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v13 + int32(32)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v13 + int32(40)
								v126 = F_construct_array_builtin(m, v13+int32(12), int32(3), int32(701))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									v128 = v126
									m.G0 = v13 + int32(48)
									return v128
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_float4eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float32
	_ = v9
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(2147483647)
	v8 = base.I32_reinterpret_f32(v5) & v7
	v9 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2139095041)) <= base.Ui32(base.I32_reinterpret_f32(v9)&v7) {
		return base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(v8))
	} else {
		return base.B2i32(base.Ui32(v8) < base.Ui32(int32(2139095041))) & base.F32_eq(v5, v9)
	}
}
func F_float4recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgfloat4(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return base.I32_reinterpret_f32(v3)
	}
}
func F_float84ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = base.F64_promote_f32(v11)
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_float8abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_abs(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_float8out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v7 = F_palloc(m, int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_float8out[0]))
		if int32(0) < v12 {
			F_double_to_shortest_decimal_buf(m, v5, v7)
			mBase = m.M
			return v7
		} else {
			F_pg_strfromd(m, v7, v12+int32(15), v5)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v7
			}
		}
	}
}
func F_float8recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pq_getmsgfloat8(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_Float8GetDatum(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_fmodl(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64) {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int32
	_ = v135
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v193 int32
	_ = v193
	var v197 int64
	_ = v197
	var v198 int64
	_ = v198
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int64
	_ = v242
	var v243 int64
	_ = v243
	var v248 int32
	_ = v248
	var v254 int64
	_ = v254
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int64
	_ = v262
	var v263 int64
	_ = v263
	var v271 int64
	_ = v271
	var v277 int64
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v280 int64
	_ = v280
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v288 int64
	_ = v288
	var v294 int64
	_ = v294
	var v296 int32
	_ = v296
	var v300 int64
	_ = v300
	var v305 int64
	_ = v305
	var v308 int64
	_ = v308
	var v314 int64
	_ = v314
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v346 int64
	_ = v346
	var v351 int64
	_ = v351
	var v354 int64
	_ = v354
	var v360 int64
	_ = v360
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v373 int64
	_ = v373
	var v375 int32
	_ = v375
	var v378 int64
	_ = v378
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v389 int64
	_ = v389
	var v396 int64
	_ = v396
	var v398 int32
	_ = v398
	var v401 int64
	_ = v401
	var v405 int32
	_ = v405
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v433 int64
	_ = v433
	var v435 int64
	_ = v435
	v10 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v26 = l4 & int64(9223372036854775807)
	v27 = int64(9223090561878065152)
	if v26 == v27 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v435
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v433
	m.G0 = v15 + int32(128)
	return
L2:
	;
	v119 = int64(9223372036854775807)
	v120 = l2 & v119
	v122 = l4 & v119
	v126 = int32(1)
	v130 = v120 & v119
	v131 = int64(9223090561878065152)
	if v130 == v131 {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	F___multf3(m, v15+int32(16), l1, l2, l3, l4)
	mBase = m.M
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	F___divtf3(m, v15, v114, v115, v114, v115)
	mBase = m.M
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	v433 = v117
	v435 = v118
	goto L1
L4:
	;
	if v74 == int32(0) {
		goto L3
	} else {
		goto L32
	}
L5:
	;
	v74 = v70
	goto L4
L6:
	;
	v31 = base.B2i32(l3 != v10)
	goto L8
L7:
	;
	v31 = base.B2i32(base.Ui64(v27) < base.Ui64(v26))
	goto L8
L8:
	;
	if v31 != 0 {
		v70 = int32(1)
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L11
L11:
	;
	goto L12
L12:
	;
	goto L13
L13:
	;
	if l3|v10|(v26|int64(0)) == int64(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v74 = int32(0)
	goto L4
L15:
	;
	goto L16
L16:
	;
	if int64(0) <= l4&v10 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if l4 == v10 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if l4 == v10 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v53 = base.B2i32(base.Ui64(l3) < base.Ui64(v10))
	goto L22
L21:
	;
	v53 = base.B2i32(l4 < v10)
	goto L22
L22:
	;
	if v53 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = int32(-1)
	goto L4
L24:
	;
	goto L25
L25:
	;
	v74 = base.B2i32(l3^v10|(l4^v10) != int64(0))
	goto L4
L26:
	;
	v63 = base.B2i32(base.Ui64(v10) < base.Ui64(l3))
	goto L28
L27:
	;
	v63 = base.B2i32(v10 < l4)
	goto L28
L28:
	;
	if v63 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v74 = int32(-1)
	goto L4
L30:
	;
	goto L31
L31:
	;
	v70 = base.B2i32(l3^v10|(l4^v10) != int64(0))
	goto L5
L32:
	;
	v78 = l4 & int64(281474976710655)
	v82 = int32(_a_F_fmodl_0)
	v83 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & v82
	if v83 != v82 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	if v97 == int32(0) {
		goto L3
	} else {
		goto L42
	}
L34:
	;
	v97 = v96
	goto L33
L35:
	;
	if v83 != 0 {
		v96 = int32(4)
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v96 = base.B2i32(l3|v78 == int64(0))
	goto L34
L38:
	;
	if l3|v78 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v92 = int32(2)
	goto L41
L40:
	;
	v92 = int32(3)
	goto L41
L41:
	;
	v97 = v92
	goto L33
L42:
	;
	v102 = base.I32_wrap_i64(int64(base.Ui64(l2) >> (uint(int64(48)) % 64)))
	v103 = int32(_a_F_fmodl_0)
	v104 = v102 & v103
	if v104 != v103 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	goto L3
L44:
	;
	if v178 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L45:
	;
	v178 = v174
	goto L44
L46:
	;
	v135 = base.B2i32(l1 != int64(0))
	goto L48
L47:
	;
	v135 = base.B2i32(base.Ui64(v131) < base.Ui64(v130))
	goto L48
L48:
	;
	if v135 != 0 {
		v174 = v126
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v139 = v122 & int64(9223372036854775807)
	v140 = int64(9223090561878065152)
	if v139 == v140 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v144 = base.B2i32(l3 != int64(0))
	goto L52
L51:
	;
	v144 = base.B2i32(base.Ui64(v140) < base.Ui64(v139))
	goto L52
L52:
	;
	if v144 != 0 {
		v174 = v126
		goto L45
	} else {
		goto L53
	}
L53:
	;
	if l1|l3|(v130|v139) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v178 = int32(0)
	goto L44
L55:
	;
	goto L56
L56:
	;
	if int64(0) <= v120&v122 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v120 == v122 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	if v120 == v122 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v157 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L62
L61:
	;
	v157 = base.B2i32(v120 < v122)
	goto L62
L62:
	;
	if v157 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v178 = int32(-1)
	goto L44
L64:
	;
	goto L65
L65:
	;
	v178 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L44
L66:
	;
	v167 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L68
L67:
	;
	v167 = base.B2i32(v122 < v120)
	goto L68
L68:
	;
	if v167 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v178 = int32(-1)
	goto L44
L70:
	;
	goto L71
L71:
	;
	v174 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L45
L72:
	;
	v184 = int32(1)
	v188 = v120 & int64(9223372036854775807)
	v189 = int64(9223090561878065152)
	if v188 == v189 {
		goto L77
	} else {
		goto L78
	}
L73:
	;
	goto L74
L74:
	;
	v248 = base.I32_wrap_i64(int64(base.Ui64(l4)>>(uint(int64(48))%64))) & int32(_a_F_fmodl_0)
	if v104 != 0 {
		goto L106
	} else {
		goto L107
	}
L75:
	;
	if v236 != 0 {
		goto L103
	} else {
		goto L104
	}
L76:
	;
	v236 = v232
	goto L75
L77:
	;
	v193 = base.B2i32(l1 != int64(0))
	goto L79
L78:
	;
	v193 = base.B2i32(base.Ui64(v189) < base.Ui64(v188))
	goto L79
L79:
	;
	if v193 != 0 {
		v232 = v184
		goto L76
	} else {
		goto L80
	}
L80:
	;
	v197 = v122 & int64(9223372036854775807)
	v198 = int64(9223090561878065152)
	if v197 == v198 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v202 = base.B2i32(l3 != int64(0))
	goto L83
L82:
	;
	v202 = base.B2i32(base.Ui64(v198) < base.Ui64(v197))
	goto L83
L83:
	;
	if v202 != 0 {
		v232 = v184
		goto L76
	} else {
		goto L84
	}
L84:
	;
	if l1|l3|(v188|v197) == int64(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v236 = int32(0)
	goto L75
L86:
	;
	goto L87
L87:
	;
	if int64(0) <= v120&v122 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v120 == v122 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v120 == v122 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	v215 = base.B2i32(base.Ui64(l1) < base.Ui64(l3))
	goto L93
L92:
	;
	v215 = base.B2i32(v120 < v122)
	goto L93
L93:
	;
	if v215 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v236 = int32(-1)
	goto L75
L95:
	;
	goto L96
L96:
	;
	v236 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L75
L97:
	;
	v225 = base.B2i32(base.Ui64(l3) < base.Ui64(l1))
	goto L99
L98:
	;
	v225 = base.B2i32(v122 < v120)
	goto L99
L99:
	;
	if v225 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v236 = int32(-1)
	goto L75
L101:
	;
	goto L102
L102:
	;
	v232 = base.B2i32(l1^l3|(v120^v122) != int64(0))
	goto L76
L103:
	;
	v433 = l2
	v435 = l1
	goto L1
L104:
	;
	goto L105
L105:
	;
	v239 = int64(0)
	F___multf3(m, v15+int32(112), l1, l2, v239, v239)
	mBase = m.M
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v15)+120))
	v243 = *(*int64)(unsafe.Add(mBase, uint32(v15)+112))
	v433 = v242
	v435 = v243
	goto L1
L106:
	;
	v261 = v104
	v262 = v120
	v263 = l1
	goto L108
L107:
	;
	F___multf3(m, v15+int32(96), l1, v120, int64(0), int64(4645181540655955968))
	mBase = m.M
	v254 = *(*int64)(unsafe.Add(mBase, uint32(v15)+104))
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v15)+96))
	v261 = base.I32_wrap_i64(int64(base.Ui64(v254)>>(uint(int64(48))%64))) - int32(120)
	v262 = v254
	v263 = v260
	goto L108
L108:
	;
	if v248 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F___multf3(m, v15+int32(80), l3, v122, int64(0), int64(4645181540655955968))
	mBase = m.M
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v15)+88))
	v277 = *(*int64)(unsafe.Add(mBase, uint32(v15)+80))
	v278 = v277
	v279 = base.I32_wrap_i64(int64(base.Ui64(v271)>>(uint(int64(48))%64))) - int32(120)
	v280 = v271
	goto L111
L110:
	;
	v278 = l3
	v279 = v248
	v280 = v122
	goto L111
L111:
	;
	v281 = int64(281474976710655)
	v283 = int64(281474976710656)
	v284 = v280&v281 | v283
	v288 = v262&v281 | v283
	if v279 < v261 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v294 = v263
	v296 = v261
	v300 = v288
	goto L115
L113:
	;
	v340 = v263
	v342 = v261
	v346 = v288
	goto L114
L114:
	;
	v351 = v346 - v284 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v340) < base.Ui64(v278)))
	if v351 < int64(0) {
		goto L126
	} else {
		goto L127
	}
L115:
	;
	v305 = v300 - v284 - base.I64_extend_i32_u(base.B2i32(base.Ui64(v294) < base.Ui64(v278)))
	if int64(0) <= v305 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	v340 = v332
	v342 = v279
	v346 = v330
	goto L114
L117:
	;
	v332 = v329 << (uint(int64(1)) % 64)
	v334 = v296 - int32(1)
	if v279 < v334 {
		v294 = v332
		v296 = v334
		v300 = v330
		goto L115
	} else {
		goto L124
	}
L118:
	;
	v308 = v294 - v278
	if v305|v308 == int64(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L120
L120:
	;
	v329 = v294
	v330 = v300<<(uint(int64(1))%64) | int64(base.Ui64(v294)>>(uint(int64(63))%64))
	goto L117
L121:
	;
	v314 = int64(0)
	F___multf3(m, v15+int32(32), l1, l2, v314, v314)
	mBase = m.M
	v317 = *(*int64)(unsafe.Add(mBase, uint32(v15)+40))
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v15)+32))
	v433 = v317
	v435 = v318
	goto L1
L122:
	;
	goto L123
L123:
	;
	v329 = v308
	v330 = v305<<(uint(int64(1))%64) | int64(base.Ui64(v308)>>(uint(int64(63))%64))
	goto L117
L124:
	;
	goto L116
L125:
	;
	if base.Ui64(v366) <= base.Ui64(int64(281474976710655)) {
		goto L130
	} else {
		goto L131
	}
L126:
	;
	v365 = v340
	v366 = v346
	goto L125
L127:
	;
	goto L128
L128:
	;
	v354 = v340 - v278
	if v351|v354 != int64(0) {
		v365 = v354
		v366 = v351
		goto L125
	} else {
		goto L129
	}
L129:
	;
	v360 = int64(0)
	F___multf3(m, v15+int32(48), l1, l2, v360, v360)
	mBase = m.M
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v15)+56))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v15)+48))
	v433 = v363
	v435 = v364
	goto L1
L130:
	;
	v373 = v365
	v375 = v342
	v378 = v366
	goto L133
L131:
	;
	v396 = v365
	v398 = v342
	v401 = v366
	goto L132
L132:
	;
	v405 = v102 & int32(_a_F_fmodl_1)
	if v398 <= int32(0) {
		goto L136
	} else {
		goto L137
	}
L133:
	;
	v384 = v375 - int32(1)
	v385 = int64(1)
	v386 = v373 << (uint(v385) % 64)
	v389 = int64(base.Ui64(v373)>>(uint(int64(63))%64)) | v378<<(uint(v385)%64)
	if base.Ui64(v389) < base.Ui64(int64(281474976710656)) {
		v373 = v386
		v375 = v384
		v378 = v389
		goto L133
	} else {
		goto L135
	}
L134:
	;
	v396 = v386
	v398 = v384
	v401 = v389
	goto L132
L135:
	;
	goto L134
L136:
	;
	F___multf3(m, v15-int32(-64), v396, v401&int64(281474976710655)|base.I64_extend_i32_u(v398+int32(120)|v405)<<(uint(int64(48))%64), int64(0), int64(4577627546245398528))
	mBase = m.M
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v15)+72))
	v423 = *(*int64)(unsafe.Add(mBase, uint32(v15)+64))
	v433 = v422
	v435 = v423
	goto L1
L137:
	;
	goto L138
L138:
	;
	v433 = v401&int64(281474976710655) | base.I64_extend_i32_u(v398|v405)<<(uint(int64(48))%64)
	v435 = v396
	goto L1
}
func F_fmtint(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v69 int32
	_ = v69
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int64
	_ = v92
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	v10 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(80)
	m.G0 = v19
	v21 = int32(_a_F_fmtint_0)
	switch l1 - int32(88) {
	case 0:
		v39 = int32(_a_F_fmtint_1)
		v40 = int32(1)
		v41 = int32(0)
		v42 = v39
		v43 = v40
		v44 = v10
		v45 = v41
		if l7 != 0 {
			v50 = l6 | base.B2i32(l0 != int64(0))
		} else {
			v50 = int32(1)
		}
		if v50 == int32(0) {
			v159 = v10
			v163 = v44
		} else {
			if v45 == int32(0) {
				if v43 != 0 {
					v92 = l0
					v101 = v10
					for {
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v92)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v101)+79)) = uint8(v113)
						v116 = v101 + int32(1)
						v118 = int64(base.Ui64(v92) >> (uint(int64(4)) % 64))
						if v118 != int64(0) {
							v92 = v118
							v101 = v116
							continue
						} else {
							break
						}
						break
					}
					v159 = v116
					v163 = v44
				} else {
					v121 = l0
					v130 = v10
					for {
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v121)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v130)+79)) = uint8(v142)
						v145 = v130 + int32(1)
						v147 = int64(base.Ui64(v121) >> (uint(int64(3)) % 64))
						if v147 != int64(0) {
							v121 = v147
							v130 = v145
							continue
						} else {
							break
						}
						break
					}
					v159 = v145
					v163 = v44
				}
			} else {
				v55 = l0
				v56 = v42
				v58 = v44
				v60 = v55
				v69 = v10
				for {
					v77 = int64(10)
					v78 = base.I64_div_u_s(v60, v77)
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
					v87 = v69 + int32(1)
					if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
						v60 = v78
						v69 = v87
						continue
					} else {
						break
					}
					break
				}
				v159 = v87
				v163 = v58
			}
		}
		v166 = int32(0)
		v167 = l6 - v159
		v170 = base.B2i32(v166 < v167)
		if v166 < v167 {
			v171 = v167
		} else {
			v171 = v166
		}
		v173 = l4 - (v159 + v171)
		v174 = int32(0)
		if v174 < v173 {
			v177 = v173
		} else {
			v177 = v174
		}
		if l3 != 0 {
			v179 = v166 - v177
		} else {
			v179 = v177
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v179
		F_leading_pad(m, l5, v163, v19+int32(12), l8)
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return
		} else {
			if v166 < v167 {
				F_dopr_outchmulti(m, int32(48), v171, l8)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					F_dostr(m, v19-v159+int32(80), v159, l8)
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v193 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v19-v159+int32(80), v159, l8)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v193 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	default:
		m.G0 = v19 + int32(80)
		return
	case 12, 17:
		if int64(0) <= l0 {
			if l2 != 0 {
				v33 = int32(43)
			} else {
				v33 = int32(0)
			}
			v42 = v21
			v43 = v10
			v44 = v33
			v45 = int32(1)
			if l7 != 0 {
				v50 = l6 | base.B2i32(l0 != int64(0))
			} else {
				v50 = int32(1)
			}
			if v50 == int32(0) {
				v159 = v10
				v163 = v44
			} else {
				if v45 == int32(0) {
					if v43 != 0 {
						v92 = l0
						v101 = v10
						for {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v92)&int32(15)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v101)+79)) = uint8(v113)
							v116 = v101 + int32(1)
							v118 = int64(base.Ui64(v92) >> (uint(int64(4)) % 64))
							if v118 != int64(0) {
								v92 = v118
								v101 = v116
								continue
							} else {
								break
							}
							break
						}
						v159 = v116
						v163 = v44
					} else {
						v121 = l0
						v130 = v10
						for {
							v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v121)&int32(7)))))
							*(*uint8)(unsafe.Add(mBase, uint32(v19-v130)+79)) = uint8(v142)
							v145 = v130 + int32(1)
							v147 = int64(base.Ui64(v121) >> (uint(int64(3)) % 64))
							if v147 != int64(0) {
								v121 = v147
								v130 = v145
								continue
							} else {
								break
							}
							break
						}
						v159 = v145
						v163 = v44
					}
				} else {
					v55 = l0
					v56 = v42
					v58 = v44
					v60 = v55
					v69 = v10
					for {
						v77 = int64(10)
						v78 = base.I64_div_u_s(v60, v77)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
						v87 = v69 + int32(1)
						if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
							v60 = v78
							v69 = v87
							continue
						} else {
							break
						}
						break
					}
					v159 = v87
					v163 = v58
				}
			}
		} else {
			v55 = int64(0) - l0
			v56 = v21
			v58 = int32(45)
			v60 = v55
			v69 = v10
			for {
				v77 = int64(10)
				v78 = base.I64_div_u_s(v60, v77)
				v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
				*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
				v87 = v69 + int32(1)
				if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
					v60 = v78
					v69 = v87
					continue
				} else {
					break
				}
				break
			}
			v159 = v87
			v163 = v58
		}
		v166 = int32(0)
		v167 = l6 - v159
		v170 = base.B2i32(v166 < v167)
		if v166 < v167 {
			v171 = v167
		} else {
			v171 = v166
		}
		v173 = l4 - (v159 + v171)
		v174 = int32(0)
		if v174 < v173 {
			v177 = v173
		} else {
			v177 = v174
		}
		if l3 != 0 {
			v179 = v166 - v177
		} else {
			v179 = v177
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v179
		F_leading_pad(m, l5, v163, v19+int32(12), l8)
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return
		} else {
			if v166 < v167 {
				F_dopr_outchmulti(m, int32(48), v171, l8)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					F_dostr(m, v19-v159+int32(80), v159, l8)
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v193 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v19-v159+int32(80), v159, l8)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v193 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 23:
		v42 = v21
		v43 = v10
		v44 = v10
		v45 = v10
		if l7 != 0 {
			v50 = l6 | base.B2i32(l0 != int64(0))
		} else {
			v50 = int32(1)
		}
		if v50 == int32(0) {
			v159 = v10
			v163 = v44
		} else {
			if v45 == int32(0) {
				if v43 != 0 {
					v92 = l0
					v101 = v10
					for {
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v92)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v101)+79)) = uint8(v113)
						v116 = v101 + int32(1)
						v118 = int64(base.Ui64(v92) >> (uint(int64(4)) % 64))
						if v118 != int64(0) {
							v92 = v118
							v101 = v116
							continue
						} else {
							break
						}
						break
					}
					v159 = v116
					v163 = v44
				} else {
					v121 = l0
					v130 = v10
					for {
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v121)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v130)+79)) = uint8(v142)
						v145 = v130 + int32(1)
						v147 = int64(base.Ui64(v121) >> (uint(int64(3)) % 64))
						if v147 != int64(0) {
							v121 = v147
							v130 = v145
							continue
						} else {
							break
						}
						break
					}
					v159 = v145
					v163 = v44
				}
			} else {
				v55 = l0
				v56 = v42
				v58 = v44
				v60 = v55
				v69 = v10
				for {
					v77 = int64(10)
					v78 = base.I64_div_u_s(v60, v77)
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
					v87 = v69 + int32(1)
					if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
						v60 = v78
						v69 = v87
						continue
					} else {
						break
					}
					break
				}
				v159 = v87
				v163 = v58
			}
		}
		v166 = int32(0)
		v167 = l6 - v159
		v170 = base.B2i32(v166 < v167)
		if v166 < v167 {
			v171 = v167
		} else {
			v171 = v166
		}
		v173 = l4 - (v159 + v171)
		v174 = int32(0)
		if v174 < v173 {
			v177 = v173
		} else {
			v177 = v174
		}
		if l3 != 0 {
			v179 = v166 - v177
		} else {
			v179 = v177
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v179
		F_leading_pad(m, l5, v163, v19+int32(12), l8)
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return
		} else {
			if v166 < v167 {
				F_dopr_outchmulti(m, int32(48), v171, l8)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					F_dostr(m, v19-v159+int32(80), v159, l8)
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v193 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v19-v159+int32(80), v159, l8)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v193 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 29:
		v39 = v21
		v40 = v10
		v41 = int32(1)
		v42 = v39
		v43 = v40
		v44 = v10
		v45 = v41
		if l7 != 0 {
			v50 = l6 | base.B2i32(l0 != int64(0))
		} else {
			v50 = int32(1)
		}
		if v50 == int32(0) {
			v159 = v10
			v163 = v44
		} else {
			if v45 == int32(0) {
				if v43 != 0 {
					v92 = l0
					v101 = v10
					for {
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v92)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v101)+79)) = uint8(v113)
						v116 = v101 + int32(1)
						v118 = int64(base.Ui64(v92) >> (uint(int64(4)) % 64))
						if v118 != int64(0) {
							v92 = v118
							v101 = v116
							continue
						} else {
							break
						}
						break
					}
					v159 = v116
					v163 = v44
				} else {
					v121 = l0
					v130 = v10
					for {
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v121)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v130)+79)) = uint8(v142)
						v145 = v130 + int32(1)
						v147 = int64(base.Ui64(v121) >> (uint(int64(3)) % 64))
						if v147 != int64(0) {
							v121 = v147
							v130 = v145
							continue
						} else {
							break
						}
						break
					}
					v159 = v145
					v163 = v44
				}
			} else {
				v55 = l0
				v56 = v42
				v58 = v44
				v60 = v55
				v69 = v10
				for {
					v77 = int64(10)
					v78 = base.I64_div_u_s(v60, v77)
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
					v87 = v69 + int32(1)
					if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
						v60 = v78
						v69 = v87
						continue
					} else {
						break
					}
					break
				}
				v159 = v87
				v163 = v58
			}
		}
		v166 = int32(0)
		v167 = l6 - v159
		v170 = base.B2i32(v166 < v167)
		if v166 < v167 {
			v171 = v167
		} else {
			v171 = v166
		}
		v173 = l4 - (v159 + v171)
		v174 = int32(0)
		if v174 < v173 {
			v177 = v173
		} else {
			v177 = v174
		}
		if l3 != 0 {
			v179 = v166 - v177
		} else {
			v179 = v177
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v179
		F_leading_pad(m, l5, v163, v19+int32(12), l8)
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return
		} else {
			if v166 < v167 {
				F_dopr_outchmulti(m, int32(48), v171, l8)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					F_dostr(m, v19-v159+int32(80), v159, l8)
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v193 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v19-v159+int32(80), v159, l8)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v193 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	case 32:
		v39 = v21
		v40 = int32(1)
		v41 = int32(0)
		v42 = v39
		v43 = v40
		v44 = v10
		v45 = v41
		if l7 != 0 {
			v50 = l6 | base.B2i32(l0 != int64(0))
		} else {
			v50 = int32(1)
		}
		if v50 == int32(0) {
			v159 = v10
			v163 = v44
		} else {
			if v45 == int32(0) {
				if v43 != 0 {
					v92 = l0
					v101 = v10
					for {
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v92)&int32(15)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v101)+79)) = uint8(v113)
						v116 = v101 + int32(1)
						v118 = int64(base.Ui64(v92) >> (uint(int64(4)) % 64))
						if v118 != int64(0) {
							v92 = v118
							v101 = v116
							continue
						} else {
							break
						}
						break
					}
					v159 = v116
					v163 = v44
				} else {
					v121 = l0
					v130 = v10
					for {
						v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+base.I32_wrap_i64(v121)&int32(7)))))
						*(*uint8)(unsafe.Add(mBase, uint32(v19-v130)+79)) = uint8(v142)
						v145 = v130 + int32(1)
						v147 = int64(base.Ui64(v121) >> (uint(int64(3)) % 64))
						if v147 != int64(0) {
							v121 = v147
							v130 = v145
							continue
						} else {
							break
						}
						break
					}
					v159 = v145
					v163 = v44
				}
			} else {
				v55 = l0
				v56 = v42
				v58 = v44
				v60 = v55
				v69 = v10
				for {
					v77 = int64(10)
					v78 = base.I64_div_u_s(v60, v77)
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+base.I32_wrap_i64(v60-v78*v77)))))
					*(*uint8)(unsafe.Add(mBase, uint32(v19-v69)+79)) = uint8(v84)
					v87 = v69 + int32(1)
					if base.B2i32(base.Ui64(v60) < base.Ui64(v77)) == int32(0) {
						v60 = v78
						v69 = v87
						continue
					} else {
						break
					}
					break
				}
				v159 = v87
				v163 = v58
			}
		}
		v166 = int32(0)
		v167 = l6 - v159
		v170 = base.B2i32(v166 < v167)
		if v166 < v167 {
			v171 = v167
		} else {
			v171 = v166
		}
		v173 = l4 - (v159 + v171)
		v174 = int32(0)
		if v174 < v173 {
			v177 = v173
		} else {
			v177 = v174
		}
		if l3 != 0 {
			v179 = v166 - v177
		} else {
			v179 = v177
		}
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v179
		F_leading_pad(m, l5, v163, v19+int32(12), l8)
		mBase = m.M
		v184 = m.ExcPending
		if v184 != 0 {
			return
		} else {
			if v166 < v167 {
				F_dopr_outchmulti(m, int32(48), v171, l8)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return
				} else {
					F_dostr(m, v19-v159+int32(80), v159, l8)
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
						if int32(0) <= v193 {
							m.G0 = v19 + int32(80)
							return
						} else {
							F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return
							} else {
								m.G0 = v19 + int32(80)
								return
							}
						}
					}
				}
			} else {
				F_dostr(m, v19-v159+int32(80), v159, l8)
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
					if int32(0) <= v193 {
						m.G0 = v19 + int32(80)
						return
					} else {
						F_dopr_outchmulti(m, int32(32), int32(0)-v193, l8)
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return
						} else {
							m.G0 = v19 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_fputs(m *base.Module, l0 int32, l1 int32) int32 {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v6 = F_strlen(m, l0)
	v7 = F_fwrite(m, l0, int32(1), v6, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != v6 {
			v12 = int32(-1)
		} else {
			v12 = int32(0)
		}
		return v12
	}
}
func F_freearc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v75 int32
	_ = v75
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v9 < int32(0) {
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v14 = v12 - int32(97)
		if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v14))|base.B2i32(int32(1)<<(uint(v14)%32)&int32(_a_F_freearc_0) == int32(0)) != 0 {
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v24 != 0 {
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				if v25 == int32(0) {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v29+v9*int32(24))+12)) = v33
					v37 = v33
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v35
					v37 = v35
				}
				if v37 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v37)+36)) = v25
				} else {
				}
				*(*int64)(unsafe.Add(mBase, uint32(l1)+32)) = int64(0)
			}
		}
	}
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v44 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v43
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v44)+16)) = v43
	}
	if v43 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v43)+20)) = v44
	} else {
	}
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v50 - int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v55 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v54
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v55)+24)) = v54
	}
	if v54 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v54)+28)) = v55
	} else {
	}
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v61 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v68 = l1 + int32(8)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	return
}
func F_freestate_cluster(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+152))
		if v6 != 0 {
			v9 = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			F_ExecDropSingleTupleTableSlot(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
				F_FreeExecutorState(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v7 = F_MakePerTupleExprContext(m, v5)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				v9 = v7
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
				F_ExecDropSingleTupleTableSlot(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
					F_FreeExecutorState(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
func F_ftoi2(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13881(m, l0, int32(_a_F_ftoi2_0), int32(1328), int32(_a_F_ftoi2_1), float32(32768), float32(-32768))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ftoi4(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13881(m, l0, int32(_a_F_ftoi4_0), int32(1303), int32(_a_F_ftoi4_1), float32(2.1474836e+09), float32(-2.1474836e+09))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_ftoi8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float32
	_ = v3
	var v4 float32
	_ = v4
	var v12 int32
	_ = v12
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	v3 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = base.F32_nearest(v3)
	v12 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)))|base.B2i32(base.F32_ge(v4, float32(-9.223372e+18)) == v12) == v12)&base.F32_lt(v4, float32(9.223372e+18)) == v12 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_ftoi8_0), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_ftoi8_1), int32(1347), int32(_a_F_ftoi8_2))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
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
		v41 = F_Int64GetDatum(m, base.I64_trunc_sat_f32_s(v4))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			return v41
		}
	}
}
