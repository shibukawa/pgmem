package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bool_anytrue(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int64
	_ = v7
	var v11 int32
	_ = v11
	var v15 int64
	_ = v15
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 != 0 {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
		return int32(0)
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v4 == int32(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v7 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			if v7 != int64(0) {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
				return base.B2i32(int64(0) < v15)
			} else {
				v11 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
				return int32(0)
			}
		}
	}
}
func F_makeBoolConst(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v2 = l1
	v5 = F_palloc0(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(-1)
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+25)) = uint8(v11)
		*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)) = uint8(v2)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(4294967295)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(68719476743)
		return v5
	}
}
func F_makeBoolExpr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(21)
		return v6
	}
}
func F_parse_bool_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v5 - int32(48) {
	case 0:
		goto L7
	case 1:
		goto L8
	default:
		goto L5
	case 22, 54:
		goto L12
	case 30, 62:
		goto L10
	case 31, 63:
		goto L9
	case 36, 68:
		goto L13
	case 41, 73:
		goto L11
	}
L1:
	;
	return v358
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v354)
	v358 = v356
	goto L1
L3:
	;
	if l2 == int32(0) {
		v358 = v349
		goto L1
	} else {
		goto L126
	}
L4:
	;
	v349 = int32(1)
	goto L3
L5:
	;
	v345 = int32(0)
	if l2 != 0 {
		v354 = v345
		v356 = v345
		goto L2
	} else {
		goto L125
	}
L6:
	;
	v354 = int32(0)
	v356 = v340
	goto L2
L7:
	;
	v334 = int32(1)
	if l1 != v334 {
		goto L5
	} else {
		goto L123
	}
L8:
	;
	v331 = int32(1)
	if l1 != v331 {
		goto L5
	} else {
		goto L122
	}
L9:
	;
	v221 = int32(2)
	if base.Ui32(l1) <= base.Ui32(v221) {
		goto L84
	} else {
		goto L85
	}
L10:
	;
	v169 = l0
	v170 = int32(241550)
	v171 = l1
	goto L67
L11:
	;
	v116 = l0
	v117 = int32(157710)
	v118 = l1
	goto L50
L12:
	;
	v62 = l0
	v63 = int32(362226)
	v64 = l1
	goto L32
L13:
	;
	v11 = l0
	v12 = int32(345210)
	v13 = l1
	goto L15
L14:
	;
	if v58 != 0 {
		goto L5
	} else {
		goto L30
	}
L15:
	;
	if v13 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v58 = int32(0)
	goto L14
L17:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 == v17 {
		v39 = v16
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v41 = int32(1)
	if v39 != 0 {
		v11 = v11 + v41
		v12 = v12 + v41
		v13 = v13 - v41
		goto L15
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v16-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v27 = v16 | int32(32)
	goto L24
L23:
	;
	v27 = v16
	goto L24
L24:
	;
	if base.Ui32((v17-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v36 = v17 | int32(32)
	goto L27
L26:
	;
	v36 = v17
	goto L27
L27:
	;
	if v27 == v36 {
		v39 = v27
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v58 = v27 - v36
	goto L14
L29:
	;
	goto L19
L30:
	;
	goto L4
L31:
	;
	if v109 != 0 {
		goto L5
	} else {
		goto L47
	}
L32:
	;
	if v64 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v109 = int32(0)
	goto L31
L34:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	if v67 == v68 {
		v90 = v67
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v92 = int32(1)
	if v90 != 0 {
		v62 = v62 + v92
		v63 = v63 + v92
		v64 = v64 - v92
		goto L32
	} else {
		goto L46
	}
L38:
	;
	if base.Ui32((v67-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v78 = v67 | int32(32)
	goto L41
L40:
	;
	v78 = v67
	goto L41
L41:
	;
	if base.Ui32((v68-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v87 = v68 | int32(32)
	goto L44
L43:
	;
	v87 = v68
	goto L44
L44:
	;
	if v78 == v87 {
		v90 = v78
		goto L37
	} else {
		goto L45
	}
L45:
	;
	v109 = v78 - v87
	goto L31
L46:
	;
	goto L36
L47:
	;
	if l2 != 0 {
		v340 = int32(1)
		goto L6
	} else {
		goto L48
	}
L48:
	;
	return int32(1)
L49:
	;
	if v163 == int32(0) {
		goto L4
	} else {
		goto L65
	}
L50:
	;
	if v118 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v163 = int32(0)
	goto L49
L52:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v121 == v122 {
		v144 = v121
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	goto L51
L55:
	;
	v146 = int32(1)
	if v144 != 0 {
		v116 = v116 + v146
		v117 = v117 + v146
		v118 = v118 - v146
		goto L50
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v121-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v132 = v121 | int32(32)
	goto L59
L58:
	;
	v132 = v121
	goto L59
L59:
	;
	if base.Ui32((v122-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v141 = v122 | int32(32)
	goto L62
L61:
	;
	v141 = v122
	goto L62
L62:
	;
	if v132 == v141 {
		v144 = v132
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v163 = v132 - v141
	goto L49
L64:
	;
	goto L54
L65:
	;
	goto L5
L66:
	;
	if v216 != 0 {
		goto L5
	} else {
		goto L82
	}
L67:
	;
	if v171 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v216 = int32(0)
	goto L66
L69:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	if v174 == v175 {
		v197 = v174
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	v199 = int32(1)
	if v197 != 0 {
		v169 = v169 + v199
		v170 = v170 + v199
		v171 = v171 - v199
		goto L67
	} else {
		goto L81
	}
L73:
	;
	if base.Ui32((v174-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v185 = v174 | int32(32)
	goto L76
L75:
	;
	v185 = v174
	goto L76
L76:
	;
	if base.Ui32((v175-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v194 = v175 | int32(32)
	goto L79
L78:
	;
	v194 = v175
	goto L79
L79:
	;
	if v185 == v194 {
		v197 = v185
		goto L72
	} else {
		goto L80
	}
L80:
	;
	v216 = v185 - v194
	goto L66
L81:
	;
	goto L71
L82:
	;
	if l2 != 0 {
		v340 = int32(1)
		goto L6
	} else {
		goto L83
	}
L83:
	;
	return int32(1)
L84:
	;
	v224 = v221
	goto L86
L85:
	;
	v224 = l1
	goto L86
L86:
	;
	v227 = l0
	v228 = int32(273738)
	v229 = v224
	goto L88
L87:
	;
	if v274 == int32(0) {
		goto L4
	} else {
		goto L103
	}
L88:
	;
	if v229 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v274 = int32(0)
	goto L87
L90:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
	if v232 == v233 {
		v255 = v232
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	goto L89
L93:
	;
	v257 = int32(1)
	if v255 != 0 {
		v227 = v227 + v257
		v228 = v228 + v257
		v229 = v229 - v257
		goto L88
	} else {
		goto L102
	}
L94:
	;
	if base.Ui32((v232-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v243 = v232 | int32(32)
	goto L97
L96:
	;
	v243 = v232
	goto L97
L97:
	;
	if base.Ui32((v233-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v252 = v233 | int32(32)
	goto L100
L99:
	;
	v252 = v233
	goto L100
L100:
	;
	if v243 == v252 {
		v255 = v243
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v274 = v243 - v252
	goto L87
L102:
	;
	goto L92
L103:
	;
	v280 = l0
	v281 = int32(339752)
	v282 = v224
	goto L105
L104:
	;
	if v327 != 0 {
		goto L5
	} else {
		goto L120
	}
L105:
	;
	if v282 != 0 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v327 = int32(0)
	goto L104
L107:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v285 == v286 {
		v308 = v285
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	goto L106
L110:
	;
	v310 = int32(1)
	if v308 != 0 {
		v280 = v280 + v310
		v281 = v281 + v310
		v282 = v282 - v310
		goto L105
	} else {
		goto L119
	}
L111:
	;
	if base.Ui32((v285-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v296 = v285 | int32(32)
	goto L114
L113:
	;
	v296 = v285
	goto L114
L114:
	;
	if base.Ui32((v286-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v305 = v286 | int32(32)
	goto L117
L116:
	;
	v305 = v286
	goto L117
L117:
	;
	if v296 == v305 {
		v308 = v296
		goto L110
	} else {
		goto L118
	}
L118:
	;
	v327 = v296 - v305
	goto L104
L119:
	;
	goto L109
L120:
	;
	if l2 != 0 {
		v340 = int32(1)
		goto L6
	} else {
		goto L121
	}
L121:
	;
	return int32(1)
L122:
	;
	v349 = v331
	goto L3
L123:
	;
	if l2 != 0 {
		v340 = v334
		goto L6
	} else {
		goto L124
	}
L124:
	;
	return int32(1)
L125:
	;
	v358 = v345
	goto L1
L126:
	;
	v354 = v349
	v356 = int32(1)
	goto L2
}
