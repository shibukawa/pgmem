package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_partial_path(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 float64
	_ = v128
	var v129 float64
	_ = v129
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v170 int32
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v192 float64
	_ = v192
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	v3 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_add_partial_path[0]))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v20 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v26 = v3
	v27 = v20
	v31 = v3
	goto L9
L7:
	;
	v307 = v3
	v312 = int32(0)
	goto L8
L8:
	;
	v313 = F_list_insert_nth(m, v312, v307, l1)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L99
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v26 < v36 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v307 = v290
	v312 = v295
	goto L8
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v26<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+64))
	if v38 == v44 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v290 = v31
	goto L13
L13:
	;
	goto L10
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+40))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v107 = int32(1)
	if v104 == int32(3) {
		v223 = v107
		goto L43
	} else {
		goto L44
	}
L15:
	;
	v104 = int32(0)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v53 = int32(0)
	goto L20
L18:
	;
	if v93 != 0 {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v88 = int32(0)
	if v75 != 0 {
		goto L32
	} else {
		goto L33
	}
L20:
	;
	v57 = int32(0)
	if v38 == v57 {
		v67 = v57
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v104 = int32(3)
	goto L14
L22:
	;
	if v44 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v61 <= v53 {
		v67 = int32(0)
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v67 = v63 + v53<<(uint(int32(2))%32)
	goto L22
L25:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if base.B2i32(v67 == v73)|base.B2i32(v75 == v73) != 0 {
		goto L19
	} else {
		goto L30
	}
L26:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v53 < v68 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v70 = int32(0)
	v93 = base.B2i32(v67 == v70)
	v95 = v70
	goto L18
L29:
	;
	goto L28
L30:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75+v53<<(uint(int32(2))%32))))
	if v83 == v85 {
		v53 = v53 + int32(1)
		goto L20
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	v92 = int32(2)
	goto L34
L33:
	;
	v92 = v88
	goto L34
L34:
	;
	v93 = base.B2i32(v67 == v88)
	v95 = v92
	goto L18
L35:
	;
	v97 = v95
	goto L37
L36:
	;
	v97 = int32(1)
	goto L37
L37:
	;
	v104 = v97
	goto L14
L38:
	;
	if v273 != 0 {
		v26 = v272 + int32(1)
		v27 = v273
		v31 = v275
		goto L9
	} else {
		goto L98
	}
L39:
	;
	F_pfree(m, l1)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L97
	}
L40:
	;
	if v255 != 0 {
		v272 = v26
		v273 = v27
		v275 = v31
		goto L38
	} else {
		goto L96
	}
L41:
	;
	if v244 == int32(0) {
		goto L39
	} else {
		goto L95
	}
L42:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v43)+56))
	if base.F64_ge(v238, v239) == int32(0) {
		v255 = v232
		goto L40
	} else {
		goto L94
	}
L43:
	;
	if v105 < v106 {
		v244 = v223
		goto L41
	} else {
		goto L92
	}
L44:
	;
	if v105 != v106 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v223 = base.B2i32(v104 == int32(1))
	goto L43
L46:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v209 = F_list_delete_nth_cell(m, v208, v26)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L90
	}
L47:
	;
	v170 = int32(0)
	switch v104 - int32(1) {
	case 0:
		goto L46
	case 1:
		v232 = v170
		goto L42
	default:
		goto L77
	}
L48:
	;
	if v104 != int32(2) {
		goto L46
	} else {
		goto L76
	}
L49:
	;
	if v106 < v105 {
		goto L48
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v112 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v43)+56))
	if base.F64_gt(v112, base.F64_mul(v113, float64(1.01))) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L45
L53:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v118 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_lt(base.F64_mul(v112, float64(1.01)), v113) != 0 {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v43)+48))
	v129 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	if base.F64_gt(v128, base.F64_mul(v129, float64(1.01))) == int32(0) {
		goto L45
	} else {
		goto L62
	}
L57:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+24)))
	if v121 == int32(0) {
		goto L45
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+25)))
	if v124 != int32(1) {
		goto L45
	} else {
		goto L61
	}
L60:
	;
	goto L56
L61:
	;
	goto L56
L62:
	;
	v232 = int32(1)
	goto L42
L63:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v139 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	v153 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v43)+48))
	if base.F64_gt(v153, base.F64_mul(v154, float64(1.01))) != 0 {
		goto L45
	} else {
		goto L74
	}
L66:
	;
	if v104 == int32(2) {
		v232 = v107
		goto L42
	} else {
		goto L72
	}
L67:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+24)))
	if v142 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+25)))
	if v143 != int32(1) {
		goto L48
	} else {
		goto L71
	}
L70:
	;
	goto L48
L71:
	;
	goto L66
L72:
	;
	v148 = *(*float64)(unsafe.Add(mBase, uint32(l1)+48))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v43)+48))
	if base.F64_gt(v148, base.F64_mul(v149, float64(1.01))) != 0 {
		v232 = v107
		goto L42
	} else {
		goto L73
	}
L73:
	;
	goto L46
L74:
	;
	if base.F64_gt(v154, base.F64_mul(v153, float64(1.01))) == int32(0) {
		goto L47
	} else {
		goto L75
	}
L75:
	;
	goto L48
L76:
	;
	v223 = v107
	goto L43
L77:
	;
	if base.F64_gt(v112, base.F64_mul(v113, float64(1.0000000001))) != 0 {
		v232 = v170
		goto L42
	} else {
		goto L78
	}
L78:
	;
	if base.F64_lt(base.F64_mul(v112, float64(1.0000000001)), v113) != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	if v180 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L81
L81:
	;
	v192 = float64(1.0000000001)
	if base.B2i32(base.F64_gt(v154, base.F64_mul(v153, v192)) == int32(0))|base.F64_gt(v153, base.F64_mul(v154, v192)) != 0 {
		v232 = v170
		goto L42
	} else {
		goto L89
	}
L82:
	;
	if base.F64_gt(v153, base.F64_mul(v154, float64(1.0000000001))) == int32(0) {
		goto L46
	} else {
		goto L88
	}
L83:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+24)))
	if v183 != 0 {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+25)))
	if v184 != int32(1) {
		goto L46
	} else {
		goto L87
	}
L86:
	;
	goto L46
L87:
	;
	goto L82
L88:
	;
	v232 = v170
	goto L42
L89:
	;
	goto L46
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v209
	F_pfree(m, v43)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	v272 = v26 - int32(1)
	v273 = v209
	v275 = v31
	goto L38
L92:
	;
	if v105 != v106 {
		v255 = v223
		goto L40
	} else {
		goto L93
	}
L93:
	;
	v232 = v223
	goto L42
L94:
	;
	v244 = v232
	goto L41
L95:
	;
	v272 = v26
	v273 = v27
	v275 = v26 + int32(1)
	goto L38
L96:
	;
	goto L39
L97:
	;
	return
L98:
	;
	v290 = v275
	goto L13
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v313
	return
}
func F_add_partial_path_precheck(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v64 int32
	_ = v64
	var v71 float64
	_ = v71
	var v77 int32
	_ = v77
	var v79 float64
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 float64
	_ = v235
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
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
	var v321 int32
	_ = v321
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v384
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v192 = v16
	goto L4
L4:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v199 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L5:
	;
	v21 = float64(1.01)
	v22 = base.F64_mul(l2, v21)
	v34 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v192 = v183
	goto L4
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v34<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+40))
	if l1 != v47 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L7
L10:
	;
	v165 = v34 + int32(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v165 < v166 {
		v34 = v165
		goto L8
	} else {
		goto L52
	}
L11:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v46)+64))
	if l4 == v88 {
		goto L27
	} else {
		goto L28
	}
L12:
	;
	v49 = base.B2i32(l1 < v47)
	v85 = v49
	v87 = v49
	goto L11
L13:
	;
	goto L14
L14:
	;
	v50 = *(*float64)(unsafe.Add(mBase, uint32(v46)+56))
	if base.F64_lt(base.F64_mul(v50, float64(1.01)), l3) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v54 = int32(0)
	if v16&int32(1) == v54 {
		v85 = v54
		v87 = v54
		goto L11
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if base.F64_gt(v50, base.F64_mul(l3, v21)) != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v58 = int32(0)
	v59 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	if base.F64_gt(v59, v22) == v58 {
		v85 = v54
		v87 = v58
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L10
L20:
	;
	v64 = int32(1)
	if v16&v64 == int32(0) {
		v85 = v64
		v87 = v64
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v77 = int32(0)
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	if base.F64_gt(l2, base.F64_mul(v79, float64(1.01))) != 0 {
		v85 = v77
		v87 = v77
		goto L11
	} else {
		goto L25
	}
L23:
	;
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v46)+48))
	if base.F64_gt(l2, base.F64_mul(v71, float64(1.01))) == int32(0) {
		v85 = v64
		v87 = int32(1)
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L10
L25:
	;
	v85 = int32(1)
	v87 = base.F64_gt(v79, v22)
	goto L11
L26:
	;
	if v148 == int32(3) {
		goto L10
	} else {
		goto L50
	}
L27:
	;
	v148 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v97 = int32(0)
	goto L32
L30:
	;
	if v137 != 0 {
		goto L47
	} else {
		goto L48
	}
L31:
	;
	v132 = int32(0)
	if v119 != 0 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v101 = int32(0)
	if l4 == v101 {
		v111 = v101
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v148 = int32(3)
	goto L26
L34:
	;
	if v88 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v105 <= v97 {
		v111 = int32(0)
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v111 = v107 + v97<<(uint(int32(2))%32)
	goto L34
L37:
	;
	v117 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if base.B2i32(v111 == v117)|base.B2i32(v119 == v117) != 0 {
		goto L31
	} else {
		goto L42
	}
L38:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v97 < v112 {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v114 = int32(0)
	v137 = base.B2i32(v111 == v114)
	v139 = v114
	goto L30
L41:
	;
	goto L40
L42:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v119+v97<<(uint(int32(2))%32))))
	if v127 == v129 {
		v97 = v97 + int32(1)
		goto L32
	} else {
		goto L43
	}
L43:
	;
	goto L33
L44:
	;
	v136 = int32(2)
	goto L46
L45:
	;
	v136 = v132
	goto L46
L46:
	;
	v137 = base.B2i32(v111 == v132)
	v139 = v136
	goto L30
L47:
	;
	v141 = v139
	goto L49
L48:
	;
	v141 = int32(1)
	goto L49
L49:
	;
	v148 = v141
	goto L26
L50:
	;
	v153 = v85 | base.B2i32(v148 == int32(1))
	if base.B2i32(v153 == int32(0))|v87&base.B2i32(v148 != int32(2)) != 0 {
		v384 = v153
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L10
L52:
	;
	goto L9
L53:
	;
	return int32(1)
L54:
	;
	goto L55
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v204 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(1)
L57:
	;
	goto L58
L58:
	;
	v217 = int32(0)
	goto L59
L59:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v217<<(uint(int32(2))%32))))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)+40))
	if l1 != v230 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v384 = v374
	goto L1
L61:
	;
	if v192 != 0 {
		goto L68
	} else {
		goto L69
	}
L62:
	;
	if v230 <= l1 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v229)+56))
	if base.F64_le(l3, base.F64_mul(v235, float64(1.01))) == int32(0) {
		goto L61
	} else {
		goto L66
	}
L65:
	;
	return int32(1)
L66:
	;
	return int32(1)
L67:
	;
	v374 = int32(1)
	v376 = v217 + v374
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v376 < v377 {
		v217 = v376
		goto L59
	} else {
		goto L115
	}
L68:
	;
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v229)+48))
	if base.F64_gt(l2, base.F64_mul(v243, float64(1.01))) == int32(0) {
		goto L67
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	if v250 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v253 = int32(0)
	goto L74
L73:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v229)+64))
	v253 = v252
	goto L74
L74:
	;
	if l4 == v253 {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v313&int32(-3) != 0 {
		goto L67
	} else {
		goto L99
	}
L76:
	;
	v313 = int32(0)
	goto L75
L77:
	;
	goto L78
L78:
	;
	v262 = int32(0)
	goto L81
L79:
	;
	if v302 != 0 {
		goto L96
	} else {
		goto L97
	}
L80:
	;
	v297 = int32(0)
	if v284 != 0 {
		goto L93
	} else {
		goto L94
	}
L81:
	;
	v266 = int32(0)
	if l4 == v266 {
		v276 = v266
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v313 = int32(3)
	goto L75
L83:
	;
	if v253 != 0 {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v270 <= v262 {
		v276 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v276 = v272 + v262<<(uint(int32(2))%32)
	goto L83
L86:
	;
	v282 = int32(0)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	if base.B2i32(v276 == v282)|base.B2i32(v284 == v282) != 0 {
		goto L80
	} else {
		goto L91
	}
L87:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	if v262 < v277 {
		goto L86
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v279 = int32(0)
	v302 = base.B2i32(v276 == v279)
	v304 = v279
	goto L79
L90:
	;
	goto L89
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v284+v262<<(uint(int32(2))%32))))
	if v292 == v294 {
		v262 = v262 + int32(1)
		goto L81
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	v301 = int32(2)
	goto L95
L94:
	;
	v301 = v297
	goto L95
L95:
	;
	v302 = base.B2i32(v276 == v297)
	v304 = v301
	goto L79
L96:
	;
	v306 = v304
	goto L98
L97:
	;
	v306 = int32(1)
	goto L98
L98:
	;
	v313 = v306
	goto L75
L99:
	;
	v316 = int32(0)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v229)+16))
	if v317 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v320 = v318
	goto L102
L101:
	;
	v320 = int32(0)
	goto L102
L102:
	;
	v321 = int32(0)
	if int32(1)|base.B2i32(v320 == v321) != 0 {
		v367 = base.B2i32(v316|v320 == v321)
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v367 != 0 {
		v384 = int32(0)
		goto L1
	} else {
		goto L114
	}
L104:
	;
	goto L103
L105:
	;
	v335 = *(*int32)(unsafe.Add(mBase, 4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v335 != v336 {
		v367 = int32(0)
		goto L104
	} else {
		goto L106
	}
L106:
	;
	v338 = int32(1)
	if v335 <= v338 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v341 = v338
	goto L109
L108:
	;
	v341 = v335
	goto L109
L109:
	;
	v342 = int32(8)
	v347 = int32(0)
	goto L110
L110:
	;
	v355 = v347 << (uint(int32(2)) % 32)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v342+v355)))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v320+v342+v355)))
	v360 = base.B2i32(v357 == v359)
	if v357 != v359 {
		v367 = v360
		goto L104
	} else {
		goto L112
	}
L111:
	;
	v367 = v360
	goto L104
L112:
	;
	v363 = v347 + int32(1)
	if v363 != v341 {
		v347 = v363
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	goto L67
L115:
	;
	goto L60
}
func F_findPartialMatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
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
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v17 + int32(4)
	v27 = int32(-1)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v28 == int64(0) {
		v50 = v27
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)))
	v64 = v61
	goto L12
L2:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v50
	goto L1
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v33 = int32(0)
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33*int32(12))+4))
	if v41 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v50 = v27
	goto L2
L6:
	;
	v50 = v33
	goto L2
L7:
	;
	goto L8
L8:
	;
	v45 = v33 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v45)) < base.Ui64(v28) {
		v33 = v45
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	m.G0 = v17 + int32(16)
	return v267
L11:
	;
	if v96 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v64&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v96 = v79
	goto L11
L14:
	;
	v96 = int32(0)
	goto L11
L15:
	;
	goto L16
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v75 = v71 & (v72 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v75
	v79 = v70 + v72*int32(12)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v83 = v80 & (v81 ^ v75)
	if v83 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v86)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v90 != int32(1) {
		v64 = base.B2i32(v83 == int32(0))
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v267 = int32(0)
	goto L10
L22:
	;
	goto L23
L23:
	;
	v101 = v20 - int32(1)
	v105 = v96
	goto L24
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[0]))
	if v117 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v267 = int32(0)
	goto L10
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v125 = F_ExecStoreMinimalTuple(m, v122, v123, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L29
	} else {
		goto L31
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L28
L31:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_MemoryContextReset(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(_a_F_findPartialMatch_0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v129
	if int32(0) <= v101 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v133
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v223 = v17 + int32(4)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+8)))
	v230 = v227
	goto L54
L34:
	;
	v141 = v101
	goto L37
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_findPartialMatch[1])) = v133
	v267 = int32(1)
	goto L10
L37:
	;
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v141<<(uint(int32(1))%32)))))
	v156 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v156 < v155 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	m.T0[v159].(func(*base.Module, int32, int32))(m, l1, v155)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v163 = v155 - int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v164))))
	if v166 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	if int32(0) < v141 {
		v141 = v141 - int32(1)
		goto L37
	} else {
		goto L52
	}
L44:
	;
	v168 = v163 << (uint(int32(3)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v168+v169)))
	v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128)+6)))
	if v172 < v155 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	m.T0[v175].(func(*base.Module, int32, int32))(m, v128, v155)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L29
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v128)+20))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v163))))
	if v180 != 0 {
		goto L43
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v127+v141<<(uint(int32(2))%32))))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v188+v168)))
	v191 = F_FunctionCall2Coll(m, l2+v141*int32(28), v187, v171, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L29
	} else {
		goto L50
	}
L50:
	;
	if v191 == int64(0) {
		goto L33
	} else {
		goto L51
	}
L51:
	;
	goto L43
L52:
	;
	goto L38
L53:
	;
	if v262 != 0 {
		v105 = v262
		goto L24
	} else {
		goto L63
	}
L54:
	;
	if v230&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v262 = v245
	goto L53
L56:
	;
	v262 = int32(0)
	goto L53
L57:
	;
	goto L58
L58:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v221)+20))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v241 = v237 & (v238 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v241
	v245 = v236 + v238*int32(12)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v249 = v246 & (v247 ^ v241)
	if v249 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v252 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+8)) = uint8(v252)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
	if v256 != int32(1) {
		v230 = base.B2i32(v249 == int32(0))
		goto L54
	} else {
		goto L62
	}
L62:
	;
	goto L55
L63:
	;
	goto L25
}
