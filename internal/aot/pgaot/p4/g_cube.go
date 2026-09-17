package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_internal_consistent(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	var v157 float64
	_ = v157
	var v160 int32
	_ = v160
	var v165 float64
	_ = v165
	var v167 int32
	_ = v167
	var v171 float64
	_ = v171
	var v177 float64
	_ = v177
	var v183 float64
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v202 int32
	_ = v202
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v272 int32
	_ = v272
	var v273 float64
	_ = v273
	var v281 float64
	_ = v281
	var v285 int32
	_ = v285
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 float64
	_ = v329
	var v331 int32
	_ = v331
	var v335 float64
	_ = v335
	var v338 float64
	_ = v338
	var v339 int32
	_ = v339
	var v340 float64
	_ = v340
	var v342 int32
	_ = v342
	var v346 float64
	_ = v346
	var v349 float64
	_ = v349
	var v354 float64
	_ = v354
	var v356 float64
	_ = v356
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v380 int32
	_ = v380
	var v398 int32
	_ = v398
	v4 = int32(0)
	if base.Ui32(int32(14)) < base.Ui32(l2) {
		v226 = v4
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v228 = int32(0)
	if base.B2i32(l0 == v228)|base.B2i32(l1 == v228) != 0 {
		v380 = v228
		goto L55
	} else {
		goto L56
	}
L2:
	;
	return v226
L3:
	;
	v9 = int32(1) << (uint(l2) % 32)
	if v9&int32(_a_F_g_cube_internal_consistent_0) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if base.B2i32(v9&int32(_a_F_g_cube_internal_consistent_1) == int32(0))&base.B2i32(l2 != int32(3)) != 0 {
		v226 = v4
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = int32(0)
	if base.B2i32(l0 == v22)|base.B2i32(l1 == v22) != 0 {
		v202 = v22
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v226 = v224
	goto L2
L7:
	;
	v224 = v202
	goto L6
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = int32(2147483647)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v43 = base.B2i32(base.Ui32(v37&v38) < base.Ui32(v40&v38))
	if base.Ui32(v37&v38) < base.Ui32(v40&v38) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = l1
	goto L11
L10:
	;
	v44 = l0
	goto L11
L11:
	;
	if base.Ui32(v37&v38) < base.Ui32(v40&v38) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v132 = v116 & int32(2147483647)
	if base.Ui32(v132) <= base.Ui32(v48) {
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v45 = l0
	goto L15
L14:
	;
	v45 = l1
	goto L15
L15:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = v46 & int32(2147483647)
	if v48 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v116 = v51
	goto L12
L17:
	;
	goto L18
L18:
	;
	v52 = int32(8)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v59 = int32(0)
	goto L19
L19:
	;
	v74 = v59 << (uint(int32(3)) % 32)
	v75 = v44 + v52 + v74
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v75)))
	v78 = base.B2i32(v56 < int32(0))
	if v56 < int32(0) {
		v85 = v76
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v116 = v56
	goto L12
L21:
	;
	v86 = v74 + (v45 + v52)
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	v89 = base.B2i32(v46 < int32(0))
	if v46 < int32(0) {
		v96 = v87
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v75+v56<<(uint(int32(3))%32))))
	if base.F64_lt(v76, v82) != 0 {
		v85 = v76
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v85 = v82
	goto L21
L24:
	;
	if base.F64_gt(v85, v96) != 0 {
		v202 = v22
		goto L7
	} else {
		goto L27
	}
L25:
	;
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v86+v46<<(uint(int32(3))%32))))
	if base.F64_gt(v87, v93) != 0 {
		v96 = v87
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v96 = v93
	goto L24
L27:
	;
	if v56 < int32(0) {
		v103 = v76
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v46 < int32(0) {
		v110 = v87
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v75+v56<<(uint(int32(3))%32))))
	if base.F64_gt(v76, v101) != 0 {
		v103 = v76
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v103 = v101
	goto L28
L31:
	;
	if base.F64_gt(v110, v103) != 0 {
		v202 = v22
		goto L7
	} else {
		goto L34
	}
L32:
	;
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v86+v46<<(uint(int32(3))%32))))
	if base.F64_lt(v87, v108) != 0 {
		v110 = v87
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v110 = v108
	goto L31
L34:
	;
	v114 = v59 + int32(1)
	if v114 != v48 {
		v59 = v114
		goto L19
	} else {
		goto L35
	}
L35:
	;
	goto L20
L36:
	;
	v224 = int32(1)
	goto L6
L37:
	;
	goto L38
L38:
	;
	v144 = v48
	goto L39
L39:
	;
	v156 = v44 + int32(8) + v144<<(uint(int32(3))%32)
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v156)))
	if base.B2i32(v116 < int32(0)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v202 = int32(0)
	goto L7
L41:
	;
	goto L40
L42:
	;
	if base.F64_lt(v183, float64(0)) != 0 {
		goto L41
	} else {
		goto L52
	}
L43:
	;
	v160 = int32(0)
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v156+v132<<(uint(int32(3))%32))))
	if base.F64_lt(v157, v165) != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if base.F64_gt(v157, float64(0)) != 0 {
		goto L41
	} else {
		goto L51
	}
L46:
	;
	v167 = v160
	goto L48
L47:
	;
	v167 = v116
	goto L48
L48:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v156+v167<<(uint(int32(3))%32))))
	if base.F64_gt(v171, float64(0)) != 0 {
		v202 = v160
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v177 = *(*float64)(unsafe.Add(mBase, uint32(v156+v116<<(uint(int32(3))%32))))
	if base.F64_gt(v157, v177) == int32(0) {
		v183 = v177
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v183 = v157
	goto L42
L51:
	;
	v183 = v157
	goto L42
L52:
	;
	v187 = int32(1)
	v189 = v144 + v187
	if v132 != v189 {
		v144 = v189
		goto L39
	} else {
		goto L53
	}
L53:
	;
	v202 = v187
	goto L7
L54:
	;
	return v398
L55:
	;
	v398 = v380
	goto L54
L56:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v246 = int32(2147483647)
	v247 = v245 & v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v250 = v248 & v246
	if base.Ui32(v247) < base.Ui32(v250) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v258 = v247
	goto L60
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(v247) < base.Ui32(v250) {
		goto L68
	} else {
		goto L69
	}
L60:
	;
	v272 = l1 + int32(8) + v258<<(uint(int32(3))%32)
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v272)))
	if base.F64_ne(v273, float64(0)) != 0 {
		v380 = v228
		goto L55
	} else {
		goto L62
	}
L61:
	;
	goto L59
L62:
	;
	if base.B2i32(v248 < int32(0)) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v272+v250<<(uint(int32(3))%32))))
	if base.F64_ne(v281, float64(0)) != 0 {
		v380 = v228
		goto L55
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v285 = v258 + int32(1)
	if v285 != v250 {
		v258 = v285
		goto L60
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	goto L61
L68:
	;
	v302 = v247
	goto L70
L69:
	;
	v302 = v250
	goto L70
L70:
	;
	if v302 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v398 = int32(1)
	goto L54
L72:
	;
	goto L73
L73:
	;
	v306 = int32(8)
	v313 = int32(0)
	goto L74
L74:
	;
	v325 = int32(0)
	v327 = v313 << (uint(int32(3)) % 32)
	v328 = l0 + v306 + v327
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
	v331 = base.B2i32(v245 < v325)
	if v245 < v325 {
		v338 = v329
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v380 = v366
	goto L55
L76:
	;
	v339 = v327 + (l1 + v306)
	v340 = *(*float64)(unsafe.Add(mBase, uint32(v339)))
	v342 = base.B2i32(v248 < int32(0))
	if v248 < int32(0) {
		v349 = v340
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v328+v245<<(uint(int32(3))%32))))
	if base.F64_lt(v329, v335) != 0 {
		v338 = v329
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v338 = v335
	goto L76
L79:
	;
	if base.F64_gt(v338, v349) != 0 {
		v380 = v325
		goto L55
	} else {
		goto L82
	}
L80:
	;
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v339+v248<<(uint(int32(3))%32))))
	if base.F64_lt(v340, v346) != 0 {
		v349 = v340
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v349 = v346
	goto L79
L82:
	;
	if v245 < v325 {
		v356 = v329
		goto L83
	} else {
		goto L84
	}
L83:
	;
	if v248 < int32(0) {
		v363 = v340
		goto L86
	} else {
		goto L87
	}
L84:
	;
	v354 = *(*float64)(unsafe.Add(mBase, uint32(v328+v245<<(uint(int32(3))%32))))
	if base.F64_gt(v329, v354) != 0 {
		v356 = v329
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v356 = v354
	goto L83
L86:
	;
	if base.F64_gt(v363, v356) != 0 {
		v380 = v325
		goto L55
	} else {
		goto L89
	}
L87:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v339+v248<<(uint(int32(3))%32))))
	if base.F64_gt(v340, v361) != 0 {
		v363 = v340
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v363 = v361
	goto L86
L89:
	;
	v366 = int32(1)
	v368 = v313 + v366
	if v368 != v302 {
		v313 = v368
		goto L74
	} else {
		goto L90
	}
L90:
	;
	goto L75
}
