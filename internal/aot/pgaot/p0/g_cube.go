package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_cube_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v131 int32
	_ = v131
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v167 int32
	_ = v167
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 float64
	_ = v280
	var v286 float64
	_ = v286
	var v292 float64
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v335 float64
	_ = v335
	var v337 int32
	_ = v337
	var v341 float64
	_ = v341
	var v353 float64
	_ = v353
	var v357 float64
	_ = v357
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 float64
	_ = v386
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v399 float64
	_ = v399
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v420 int32
	_ = v420
	var v452 int32
	_ = v452
	var v496 int32
	_ = v496
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(base.B2i32(v496 == int32(0)))
	return v12
L5:
	;
	v496 = int32(-1)
	goto L4
L6:
	;
	v496 = v452
	goto L4
L7:
	;
	v452 = int32(1)
	goto L6
L8:
	;
	v36 = v31
	goto L10
L9:
	;
	v36 = v34
	goto L10
L10:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(8)
	v54 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v60 = v54 << (uint(int32(3)) % 32)
	v61 = v5 + v37 + v60
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v71 = v62
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v103 = int32(8)
	v121 = int32(0)
	goto L31
L16:
	;
	v72 = v60 + (v10 + v37)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	v75 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v82 = v73
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v68) != 0 {
		v71 = v62
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v71 = v68
	goto L16
L19:
	;
	if base.F64_gt(v71, v82) != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v79) != 0 {
		v82 = v73
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = v79
	goto L19
L22:
	;
	if v29 < int32(0) {
		v89 = v62
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v32 < int32(0) {
		v96 = v73
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v87) != 0 {
		v89 = v62
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v87
	goto L23
L26:
	;
	v98 = int32(-1)
	if base.F64_lt(v89, v96) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L29
	}
L27:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v94) != 0 {
		v96 = v73
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = v94
	goto L26
L29:
	;
	v101 = v54 + int32(1)
	if v101 != v36 {
		v54 = v101
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	v127 = v121 << (uint(int32(3)) % 32)
	v128 = v5 + v103 + v127
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
	v131 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v138 = v129
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v139 = v127 + (v10 + v103)
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v139)))
	v142 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v149 = v140
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v135) != 0 {
		v138 = v129
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = v135
	goto L33
L36:
	;
	if base.F64_gt(v138, v149) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v146) != 0 {
		v149 = v140
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v149 = v146
	goto L36
L39:
	;
	if v29 < int32(0) {
		v156 = v129
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v32 < int32(0) {
		v163 = v140
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v154) != 0 {
		v156 = v129
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v156 = v154
	goto L40
L43:
	;
	if base.F64_lt(v156, v163) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v161) != 0 {
		v163 = v140
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v163 = v161
	goto L43
L46:
	;
	v167 = v121 + int32(1)
	if v167 != v36 {
		v121 = v167
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v189 = v5 + int32(8)
	v192 = v36
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L81
	} else {
		goto L82
	}
L51:
	;
	v212 = v189 + v192<<(uint(int32(3))%32)
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v255 = v36
	goto L65
L53:
	;
	if base.F64_lt(v238, float64(0)) != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v212+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v220) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_gt(v213, float64(0)) != 0 {
		goto L7
	} else {
		goto L62
	}
L57:
	;
	v222 = int32(0)
	goto L59
L58:
	;
	v222 = v29
	goto L59
L59:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v212+v222<<(uint(int32(3))%32))))
	if base.F64_gt(v226, float64(0)) != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v212+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v232) == int32(0) {
		v238 = v232
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v238 = v213
	goto L53
L62:
	;
	v238 = v213
	goto L53
L63:
	;
	v242 = v192 + int32(1)
	if v242 != v31 {
		v192 = v242
		goto L51
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	v266 = v189 + v255<<(uint(int32(3))%32)
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L5
L67:
	;
	if base.F64_lt(v292, float64(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v266+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v274) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if base.F64_gt(v267, float64(0)) != 0 {
		goto L7
	} else {
		goto L76
	}
L71:
	;
	v276 = int32(0)
	goto L73
L72:
	;
	v276 = v29
	goto L73
L73:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v266+v276<<(uint(int32(3))%32))))
	if base.F64_gt(v280, float64(0)) != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v266+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v286) == int32(0) {
		v292 = v286
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v292 = v267
	goto L67
L76:
	;
	v292 = v267
	goto L67
L77:
	;
	v297 = int32(1)
	v299 = v255 + v297
	if v299 == v31 {
		v452 = v297
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L66
L80:
	;
	v255 = v299
	goto L65
L81:
	;
	v496 = int32(0)
	goto L4
L82:
	;
	goto L83
L83:
	;
	v304 = v10 + int32(8)
	v314 = v31
	goto L84
L84:
	;
	v327 = v304 + v314<<(uint(int32(3))%32)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v374 = v36
	goto L99
L86:
	;
	if base.F64_lt(v357, float64(0)) != 0 {
		goto L7
	} else {
		goto L97
	}
L87:
	;
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v327+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v353) == int32(0) {
		v357 = v353
		goto L86
	} else {
		goto L96
	}
L88:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v327+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v335) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if base.F64_gt(v328, float64(0)) == int32(0) {
		v357 = v328
		goto L86
	} else {
		goto L95
	}
L91:
	;
	v337 = int32(0)
	goto L93
L92:
	;
	v337 = v32
	goto L93
L93:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v327+v337<<(uint(int32(3))%32))))
	if base.F64_gt(v341, float64(0)) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	goto L5
L96:
	;
	v357 = v328
	goto L86
L97:
	;
	v361 = v314 + int32(1)
	if v361 != v34 {
		v314 = v361
		goto L84
	} else {
		goto L98
	}
L98:
	;
	goto L85
L99:
	;
	v385 = v304 + v374<<(uint(int32(3))%32)
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v385)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v452 = int32(-1)
	goto L6
L101:
	;
	if base.F64_lt(v415, float64(0)) != 0 {
		goto L7
	} else {
		goto L112
	}
L102:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v385+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v411) == int32(0) {
		v415 = v411
		goto L101
	} else {
		goto L111
	}
L103:
	;
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v385+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v393) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_gt(v386, float64(0)) == int32(0) {
		v415 = v386
		goto L101
	} else {
		goto L110
	}
L106:
	;
	v395 = int32(0)
	goto L108
L107:
	;
	v395 = v32
	goto L108
L108:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v385+v395<<(uint(int32(3))%32))))
	if base.F64_gt(v399, float64(0)) == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L5
L110:
	;
	goto L5
L111:
	;
	v415 = v386
	goto L101
L112:
	;
	v420 = v374 + int32(1)
	if v34 != v420 {
		v374 = v420
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
}
