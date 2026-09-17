package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_dim(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v10 != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v9 & int32(2147483647)
			}
		} else {
			return v9 & int32(2147483647)
		}
	}
}
func F_cube_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v497 != v6 {
		goto L114
	} else {
		goto L115
	}
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
	v61 = v6 + v37 + v60
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
	v72 = v60 + (v11 + v37)
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
	v128 = v6 + v103 + v127
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
	v139 = v127 + (v11 + v103)
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
	v189 = v6 + int32(8)
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
	v304 = v11 + int32(8)
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
L114:
	;
	F_pfree(m, v6)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v501 != v11 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, v11)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	return base.B2i32(v496 == int32(0))
L121:
	;
	goto L120
}
func F_cube_inter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v120 float64
	_ = v120
	var v122 float64
	_ = v122
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 float64
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 float64
	_ = v140
	var v149 int32
	_ = v149
	var v151 float64
	_ = v151
	var v152 int32
	_ = v152
	var v157 float64
	_ = v157
	var v159 int32
	_ = v159
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	var v169 int32
	_ = v169
	var v171 float64
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v179 float64
	_ = v179
	var v183 int32
	_ = v183
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v187 float64
	_ = v187
	var v193 int32
	_ = v193
	var v194 float64
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 float64
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v219 int32
	_ = v219
	var v226 float64
	_ = v226
	var v229 int32
	_ = v229
	var v243 int32
	_ = v243
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v281 int32
	_ = v281
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 float64
	_ = v302
	var v308 float64
	_ = v308
	var v310 int32
	_ = v310
	var v314 float64
	_ = v314
	var v321 float64
	_ = v321
	var v323 float64
	_ = v323
	var v325 float64
	_ = v325
	var v327 float64
	_ = v327
	var v330 int32
	_ = v330
	var v331 float64
	_ = v331
	var v333 float64
	_ = v333
	var v335 int32
	_ = v335
	var v337 float64
	_ = v337
	var v338 float64
	_ = v338
	var v341 float64
	_ = v341
	var v344 float64
	_ = v344
	var v346 float64
	_ = v346
	var v347 float64
	_ = v347
	var v350 float64
	_ = v350
	var v353 float64
	_ = v353
	var v358 int32
	_ = v358
	var v389 int32
	_ = v389
	var v400 int32
	_ = v400
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v431 float64
	_ = v431
	var v435 float64
	_ = v435
	var v438 int32
	_ = v438
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v39 = int32(2147483647)
	v40 = v38 & v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v43 = v41 & v39
	v44 = base.B2i32(base.Ui32(v40) < base.Ui32(v43))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v45 = v36
	goto L6
L5:
	;
	v45 = v31
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v50 = v46<<(uint(int32(4))%32) | int32(8)
	v51 = F_palloc0(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v61 = v56&int32(-2147483648) | v46&int32(2147483647)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v61
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if base.Ui32(v67) < base.Ui32(v243) {
		goto L59
	} else {
		goto L60
	}
L9:
	;
	v64 = v31
	goto L11
L10:
	;
	v64 = v36
	goto L11
L11:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v67 = v65 & int32(2147483647)
	if v67 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v243 = v63 & int32(2147483647)
	goto L8
L13:
	;
	goto L14
L14:
	;
	v73 = v63 & int32(2147483647)
	v74 = int32(8)
	v75 = v51 + v74
	v100 = int32(0)
	goto L15
L15:
	;
	v112 = v100 << (uint(int32(3)) % 32)
	v113 = v45 + v74 + v112
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v116 = base.B2i32(v63 < int32(0))
	if v63 < int32(0) {
		v122 = v114
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v243 = v73
	goto L8
L17:
	;
	v124 = v112 + (v64 + v74)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	v126 = int32(0)
	v127 = base.B2i32(v65 < v126)
	if v127 == v126 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v113+v63<<(uint(int32(3))%32))))
	if base.F64_gt(v120, v114) != 0 {
		v122 = v114
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v122 = v120
	goto L17
L20:
	;
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	if v127 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L21:
	;
	v179 = *(*float64)(unsafe.Add(mBase, uint32(v113+v63<<(uint(int32(3))%32))))
	if base.F64_gt(v173, v179) == int32(0) {
		v183 = v172
		v184 = v179
		v185 = v173
		goto L20
	} else {
		goto L40
	}
L22:
	;
	v169 = v112 + v75
	*(*float64)(unsafe.Add(mBase, uint32(v169))) = v167
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	if v63 < int32(0) {
		v183 = v169
		v184 = v171
		v185 = v171
		goto L20
	} else {
		goto L39
	}
L23:
	;
	if base.B2i32(int32(0) <= v63) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v134 = *(*float64)(unsafe.Add(mBase, uint32(v124+v65<<(uint(int32(3))%32))))
	v135 = base.F64_lt(v125, v134)
	if v135 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if base.F64_gt(v122, v125) == int32(0) {
		v167 = v125
		goto L22
	} else {
		goto L32
	}
L27:
	;
	v136 = int32(0)
	goto L29
L28:
	;
	v136 = v65
	goto L29
L29:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v124+v136<<(uint(int32(3))%32))))
	if base.F64_gt(v122, v140) != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v135 != 0 {
		v167 = v125
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v167 = v134
	goto L22
L32:
	;
	goto L23
L33:
	;
	v149 = v112 + v75
	*(*float64)(unsafe.Add(mBase, uint32(v149))) = v114
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v183 = v149
	v184 = v151
	v185 = v151
	goto L20
L34:
	;
	goto L35
L35:
	;
	v152 = v112 + v75
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v113+v63<<(uint(int32(3))%32))))
	if base.F64_lt(v114, v157) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v159 = int32(0)
	goto L38
L37:
	;
	v159 = v63
	goto L38
L38:
	;
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v113+v159<<(uint(int32(3))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v152))) = v163
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v172 = v152
	v173 = v165
	goto L21
L39:
	;
	v172 = v169
	v173 = v171
	goto L21
L40:
	;
	v183 = v172
	v184 = v173
	v185 = v173
	goto L20
L41:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v219)))
	*(*float64)(unsafe.Add(mBase, uint32(v183+v73<<(uint(int32(3))%32)))) = v226
	v229 = v100 + int32(1)
	if v229 != v67 {
		v100 = v229
		goto L15
	} else {
		goto L58
	}
L42:
	;
	if v195 != 0 {
		v219 = v124
		goto L41
	} else {
		goto L57
	}
L43:
	;
	v216 = v113 + v63<<(uint(int32(3))%32)
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v216)))
	if base.F64_gt(v185, v217) != 0 {
		v219 = v113
		goto L41
	} else {
		goto L56
	}
L44:
	;
	v193 = v124 + v65<<(uint(int32(3))%32)
	v194 = *(*float64)(unsafe.Add(mBase, uint32(v193)))
	v195 = base.F64_gt(v187, v194)
	if v195 != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v206 = base.F64_lt(v184, v187)
	if v206 != 0 {
		goto L52
	} else {
		goto L53
	}
L47:
	;
	v196 = int32(0)
	goto L49
L48:
	;
	v196 = v65
	goto L49
L49:
	;
	v200 = *(*float64)(unsafe.Add(mBase, uint32(v124+v196<<(uint(int32(3))%32))))
	if base.F64_lt(v184, v200) == int32(0) {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	if v116 == int32(0) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v219 = v113
	goto L41
L52:
	;
	v207 = v113
	goto L54
L53:
	;
	v207 = v124
	goto L54
L54:
	;
	if v116|base.B2i32(v206 == int32(0)) != 0 {
		v219 = v207
		goto L41
	} else {
		goto L55
	}
L55:
	;
	goto L43
L56:
	;
	v219 = v216
	goto L41
L57:
	;
	v219 = v193
	goto L41
L58:
	;
	goto L16
L59:
	;
	v261 = int32(8)
	v262 = v51 + v261
	v264 = v243 << (uint(int32(3)) % 32)
	v281 = v67
	goto L62
L60:
	;
	goto L61
L61:
	;
	v389 = int32(0)
	if base.B2i32(v61 == v389)|base.B2i32(v56 < v389) == v389 {
		goto L95
	} else {
		goto L96
	}
L62:
	;
	v300 = v281 << (uint(int32(3)) % 32)
	v301 = v45 + v261 + v300
	v302 = *(*float64)(unsafe.Add(mBase, uint32(v301)))
	if base.B2i32(v63 < int32(0)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	goto L61
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v300+(v262+v264)))) = v353
	v358 = v281 + int32(1)
	if v358 != v243 {
		v281 = v358
		goto L62
	} else {
		goto L93
	}
L65:
	;
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v301+v264)))
	if base.F64_lt(v302, v308) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v341 = float64(0)
	if base.F64_lt(v302, v341) != 0 {
		goto L87
	} else {
		goto L88
	}
L68:
	;
	v310 = int32(0)
	goto L70
L69:
	;
	v310 = v63
	goto L70
L70:
	;
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v301+v310<<(uint(int32(3))%32))))
	if base.F64_lt(v314, float64(0)) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v325 = float64(0)
	goto L73
L72:
	;
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v301+v63<<(uint(int32(3))%32))))
	if base.F64_lt(v302, v321) != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v300+v262))) = v325
	v327 = *(*float64)(unsafe.Add(mBase, uint32(v301)))
	v330 = v301 + v63<<(uint(int32(3))%32)
	v331 = *(*float64)(unsafe.Add(mBase, uint32(v330)))
	if base.F64_gt(v327, v331) != 0 {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v323 = v302
	goto L76
L75:
	;
	v323 = v321
	goto L76
L76:
	;
	v325 = v323
	goto L73
L77:
	;
	v333 = v327
	goto L79
L78:
	;
	v333 = v331
	goto L79
L79:
	;
	v335 = base.F64_gt(v333, float64(0))
	if v335 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if v335 != 0 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v338 = *(*float64)(unsafe.Add(mBase, uint32(v330)))
	if base.F64_gt(v327, v338) != 0 {
		v353 = v327
		goto L64
	} else {
		goto L86
	}
L83:
	;
	v337 = float64(0)
	goto L85
L84:
	;
	v337 = v327
	goto L85
L85:
	;
	v353 = v337
	goto L64
L86:
	;
	v353 = v338
	goto L64
L87:
	;
	v344 = v341
	goto L89
L88:
	;
	v344 = v302
	goto L89
L89:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v300+v262))) = v344
	v346 = float64(0)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v301)))
	if base.F64_gt(v347, v346) != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v350 = v346
	goto L92
L91:
	;
	v350 = v347
	goto L92
L92:
	;
	v353 = v350
	goto L64
L93:
	;
	goto L63
L94:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(v40) < base.Ui32(v43) {
		goto L105
	} else {
		goto L106
	}
L95:
	;
	v400 = v389
	goto L98
L96:
	;
	goto L97
L97:
	;
	v472 = v46<<(uint(int32(3))%32) + int32(8)
	v473 = F_repalloc(m, v51, v472)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L102
	}
L98:
	;
	v428 = int32(3)
	v430 = v51 + int32(8) + v400<<(uint(v428)%32)
	v431 = *(*float64)(unsafe.Add(mBase, uint32(v430)))
	v435 = *(*float64)(unsafe.Add(mBase, uint32(v430+v61<<(uint(v428)%32))))
	if base.F64_ne(v431, v435) != 0 {
		v487 = v51
		goto L94
	} else {
		goto L100
	}
L99:
	;
	goto L97
L100:
	;
	v438 = v400 + int32(1)
	if v438 != v61 {
		v400 = v438
		goto L98
	} else {
		goto L101
	}
L101:
	;
	goto L99
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v473))) = v472 << (uint(int32(2)) % 32)
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v473)+4)) = v478 | int32(-2147483648)
	v487 = v473
	goto L94
L103:
	;
	return v487
L104:
	;
	F_pfree(m, v36)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L118
	}
L105:
	;
	if v511 != v31 {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	goto L107
L107:
	;
	if v511 != v31 {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	F_pfree(m, v31)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v36 != v516 {
		goto L104
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	goto L103
L113:
	;
	F_pfree(m, v31)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v36 == v521 {
		goto L103
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	goto L104
L118:
	;
	goto L103
}
func F_cube_ll_coord(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = float64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v12 <= int32(0) {
			v38 = v6
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v17 = v15 & int32(2147483647)
			if base.Ui32(v17) < base.Ui32(v12) {
				v38 = v6
			} else {
				v21 = v8 + v12<<(uint(int32(3))%32)
				if v15 < int32(0) {
					v33 = v21
				} else {
					v24 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(v21+v17<<(uint(int32(3))%32))))
					if base.F64_lt(v24, v28) != 0 {
						v33 = v21
					} else {
						v33 = v21 + v15<<(uint(int32(3))%32)
					}
				}
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
				v38 = v34
			}
		}
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v39 != v8 {
			F_pfree(m, v8)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = F_Float8GetDatum(m, v38)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					return v43
				}
			}
		} else {
			v43 = F_Float8GetDatum(m, v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				return v43
			}
		}
	}
}
func F_cube_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v52 float64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v95 float64
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 float64
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	F_initStringInfo(m, v11)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_appendStringInfoChar(m, v11, int32(40))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = v18 & int32(2147483647)
	if v25 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_appendStringInfoChar(m, v11, int32(41))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v29 = F_float8out_internal(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_appendStringInfoString(m, v11, v29)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v33 = int32(1)
	if v25 == v33 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v39 = v33
	goto L10
L10:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_cube_out_0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L5
L12:
	;
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v14+int32(8)+v39<<(uint(int32(3))%32))))
	v53 = F_float8out_internal(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	F_appendStringInfoString(m, v11, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v58 = v39 + int32(1)
	if v58 != v25 {
		v39 = v58
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v71 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.B2i32(v72 < v71)|base.B2i32(v72 == v71) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v170 != v14 {
		goto L44
	} else {
		goto L45
	}
L18:
	;
	v79 = v14 + int32(8)
	v81 = v71
	goto L19
L19:
	;
	v88 = int32(3)
	v90 = v79 + v81<<(uint(v88)%32)
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v90+v72<<(uint(v88)%32))))
	if base.F64_eq(v91, v95) != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_cube_out_1))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L25
	}
L21:
	;
	v98 = v81 + int32(1)
	if v72 != v98 {
		v81 = v98
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L17
L25:
	;
	if v25 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_appendStringInfoChar(m, v11, int32(41))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L43
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v106 = int32(0)
	if v106 < v105 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v109 = v105
	goto L30
L29:
	;
	v109 = v106
	goto L30
L30:
	;
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v79+v109<<(uint(int32(3))%32))))
	v114 = F_float8out_internal(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_appendStringInfoString(m, v11, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v118 = int32(1)
	if v25 == v118 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	v122 = v118
	goto L34
L34:
	;
	F_appendStringInfoString(m, v11, int32(_a_F_cube_out_0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L26
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v136 = int32(0)
	if v136 < v135 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v139 = v135
	goto L39
L38:
	;
	v139 = v136
	goto L39
L39:
	;
	v143 = *(*float64)(unsafe.Add(mBase, uint32(v79+v122<<(uint(int32(3))%32)+v139<<(uint(int32(3))%32))))
	v144 = F_float8out_internal(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_appendStringInfoString(m, v11, v144)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v149 = v122 + int32(1)
	if v149 != v25 {
		v122 = v149
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	goto L17
L44:
	;
	F_pfree(m, v14)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	m.G0 = v11 + int32(16)
	return v174
L47:
	;
	goto L46
}
func F_cube_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pq_getmsgint(m, v12, int32(4))
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
	v19 = v14 & int32(2147483647)
	if base.Ui32(v19) < base.Ui32(int32(101)) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v22 = int32(0)
	if v14 < v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v27 = v19
	goto L8
L7:
	;
	v27 = v14 << (uint(int32(1)) % 32)
	goto L8
L8:
	;
	v31 = v27<<(uint(int32(3))%32) + int32(8)
	v32 = F_palloc(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v31 << (uint(int32(2)) % 32)
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v22
	goto L13
L11:
	;
	goto L12
L12:
	;
	m.G0 = v10 + int32(16)
	return v32
L13:
	;
	v50 = F_pq_getmsgfloat8(m, v12)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v32+int32(8)+v40<<(uint(int32(3))%32)))) = v50
	v54 = v40 + int32(1)
	if v54 != v27 {
		v40 = v54
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errmsg(m, int32(_a_F_cube_recv_0), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(100)
	F_errdetail(m, int32(_a_F_cube_recv_1), v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_cube_recv_2), int32(371), int32(_a_F_cube_recv_3))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_scanner_finish(m *base.Module, l0 int32) {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = F_cube_yylex_destroy(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_cube_send(m *base.Module, l0 int32) int32 {
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
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v57 float64
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_pq_begintypsend(m, v8)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	F_enlargeStringInfo(m, v8, int32(4))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v27 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v22+v23))) = base.I32_rotr(v18, int32(24))&v27 | base.I32_rotr(v18&v27, int32(8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v22 + int32(4)
	v39 = v15 & int32(2147483647)
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v49 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v70 << (uint(int32(2)) % 32)
	goto L12
L8:
	;
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v11+int32(8)+v49<<(uint(int32(3))%32))))
	F_pq_sendfloat8(m, v8, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v61 = v49 + int32(1)
	if v61 != v39<<(uint(int32(base.Ui32(v40^int32(-1))>>(uint(int32(31))%32)))%32) {
		v49 = v61
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	m.G0 = v8 + int32(16)
	return v69
}
func F_cube_union_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v104 int32
	_ = v104
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v112 int32
	_ = v112
	var v113 float64
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 float64
	_ = v128
	var v137 int32
	_ = v137
	var v139 float64
	_ = v139
	var v140 int32
	_ = v140
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v151 float64
	_ = v151
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v157 int32
	_ = v157
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v163 int32
	_ = v163
	var v167 float64
	_ = v167
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v174 int32
	_ = v174
	var v175 float64
	_ = v175
	var v181 int32
	_ = v181
	var v182 float64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 float64
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v205 float64
	_ = v205
	var v207 int32
	_ = v207
	var v214 float64
	_ = v214
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v265 int32
	_ = v265
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 float64
	_ = v282
	var v290 float64
	_ = v290
	var v292 int32
	_ = v292
	var v296 float64
	_ = v296
	var v303 float64
	_ = v303
	var v305 float64
	_ = v305
	var v307 float64
	_ = v307
	var v309 float64
	_ = v309
	var v312 int32
	_ = v312
	var v313 float64
	_ = v313
	var v315 float64
	_ = v315
	var v317 int32
	_ = v317
	var v319 float64
	_ = v319
	var v320 float64
	_ = v320
	var v323 float64
	_ = v323
	var v326 float64
	_ = v326
	var v328 float64
	_ = v328
	var v329 float64
	_ = v329
	var v332 float64
	_ = v332
	var v334 float64
	_ = v334
	var v340 int32
	_ = v340
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 float64
	_ = v405
	var v409 float64
	_ = v409
	var v412 int32
	_ = v412
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return l0
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = int32(2147483647)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = base.B2i32(base.Ui32(v28&v29) < base.Ui32(v31&v29))
	if base.Ui32(v28&v29) < base.Ui32(v31&v29) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = l1
	goto L6
L5:
	;
	v35 = l0
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v40 = v36<<(uint(int32(4))%32) | int32(8)
	v41 = F_palloc0(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
	v49 = v36 & int32(2147483647)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v53 = v49 | v50&int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if base.Ui32(v28&v29) < base.Ui32(v31&v29) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if base.Ui32(v59) < base.Ui32(v231) {
		goto L60
	} else {
		goto L61
	}
L10:
	;
	v56 = l0
	goto L12
L11:
	;
	v56 = l1
	goto L12
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v59 = v57 & int32(2147483647)
	if v59 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v231 = v55 & int32(2147483647)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v65 = v55 & int32(2147483647)
	v66 = int32(8)
	v67 = v41 + v66
	v93 = int32(0)
	goto L16
L16:
	;
	v100 = v93 << (uint(int32(3)) % 32)
	v101 = v35 + v66 + v100
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v101)))
	v104 = base.B2i32(v55 < int32(0))
	if v55 < int32(0) {
		v110 = v102
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v231 = v65
	goto L9
L18:
	;
	v112 = v100 + (v56 + v66)
	v113 = *(*float64)(unsafe.Add(mBase, uint32(v112)))
	v114 = int32(0)
	v115 = base.B2i32(v57 < v114)
	if v115 == v114 {
		goto L25
	} else {
		goto L26
	}
L19:
	;
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v101+v55<<(uint(int32(3))%32))))
	if base.F64_gt(v108, v102) != 0 {
		v110 = v102
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v110 = v108
	goto L18
L21:
	;
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v112)))
	if v115 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L22:
	;
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v101+v55<<(uint(int32(3))%32))))
	if base.F64_gt(v160, v167) == int32(0) {
		v171 = v167
		v172 = v160
		v174 = v163
		goto L21
	} else {
		goto L41
	}
L23:
	;
	v157 = v100 + v67
	*(*float64)(unsafe.Add(mBase, uint32(v157))) = v154
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v101)))
	if v55 < int32(0) {
		v171 = v159
		v172 = v159
		v174 = v157
		goto L21
	} else {
		goto L40
	}
L24:
	;
	if base.B2i32(int32(0) <= v55) == int32(0) {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v112+v57<<(uint(int32(3))%32))))
	v123 = base.F64_lt(v113, v122)
	if v123 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if base.F64_lt(v110, v113) == int32(0) {
		v154 = v113
		goto L23
	} else {
		goto L33
	}
L28:
	;
	v124 = int32(0)
	goto L30
L29:
	;
	v124 = v57
	goto L30
L30:
	;
	v128 = *(*float64)(unsafe.Add(mBase, uint32(v112+v124<<(uint(int32(3))%32))))
	if base.F64_lt(v110, v128) != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if v123 != 0 {
		v154 = v113
		goto L23
	} else {
		goto L32
	}
L32:
	;
	v154 = v122
	goto L23
L33:
	;
	goto L24
L34:
	;
	v137 = v100 + v67
	*(*float64)(unsafe.Add(mBase, uint32(v137))) = v102
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v101)))
	v171 = v139
	v172 = v139
	v174 = v137
	goto L21
L35:
	;
	goto L36
L36:
	;
	v140 = v100 + v67
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v101+v55<<(uint(int32(3))%32))))
	if base.F64_lt(v102, v145) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v147 = int32(0)
	goto L39
L38:
	;
	v147 = v55
	goto L39
L39:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v101+v147<<(uint(int32(3))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v140))) = v151
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v101)))
	v160 = v153
	v163 = v140
	goto L22
L40:
	;
	v160 = v159
	v163 = v157
	goto L22
L41:
	;
	v171 = v160
	v172 = v160
	v174 = v163
	goto L21
L42:
	;
	v214 = *(*float64)(unsafe.Add(mBase, uint32(v207)))
	*(*float64)(unsafe.Add(mBase, uint32(v174+v65<<(uint(int32(3))%32)))) = v214
	v217 = v93 + int32(1)
	if v217 != v59 {
		v93 = v217
		goto L16
	} else {
		goto L59
	}
L43:
	;
	if v183 != 0 {
		v207 = v112
		goto L42
	} else {
		goto L58
	}
L44:
	;
	v204 = v101 + v55<<(uint(int32(3))%32)
	v205 = *(*float64)(unsafe.Add(mBase, uint32(v204)))
	if base.F64_gt(v172, v205) != 0 {
		v207 = v101
		goto L42
	} else {
		goto L57
	}
L45:
	;
	v181 = v112 + v57<<(uint(int32(3))%32)
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v181)))
	v183 = base.F64_gt(v175, v182)
	if v183 != 0 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v194 = base.F64_gt(v171, v175)
	if v194 != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v184 = int32(0)
	goto L50
L49:
	;
	v184 = v57
	goto L50
L50:
	;
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v112+v184<<(uint(int32(3))%32))))
	if base.F64_gt(v171, v188) == int32(0) {
		goto L43
	} else {
		goto L51
	}
L51:
	;
	if v104 == int32(0) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v207 = v101
	goto L42
L53:
	;
	v195 = v101
	goto L55
L54:
	;
	v195 = v112
	goto L55
L55:
	;
	if v104|base.B2i32(v194 == int32(0)) != 0 {
		v207 = v195
		goto L42
	} else {
		goto L56
	}
L56:
	;
	goto L44
L57:
	;
	v207 = v204
	goto L42
L58:
	;
	v207 = v181
	goto L42
L59:
	;
	goto L17
L60:
	;
	v245 = int32(8)
	v246 = v41 + v245
	v265 = v59
	goto L63
L61:
	;
	goto L62
L62:
	;
	v367 = int32(0)
	if base.B2i32(v53 == v367)|base.B2i32(v50 < v367) == v367 {
		goto L96
	} else {
		goto L97
	}
L63:
	;
	v280 = v265 << (uint(int32(3)) % 32)
	v281 = v35 + v245 + v280
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.B2i32(v55 < int32(0)) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L62
L65:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v280+(v246+v49<<(uint(int32(3))%32))))) = v334
	v340 = v265 + int32(1)
	if v340 != v231 {
		v265 = v340
		goto L63
	} else {
		goto L94
	}
L66:
	;
	v290 = *(*float64)(unsafe.Add(mBase, uint32(v281+v231<<(uint(int32(3))%32))))
	if base.F64_lt(v282, v290) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	v323 = float64(0)
	if base.F64_gt(v282, v323) != 0 {
		goto L88
	} else {
		goto L89
	}
L69:
	;
	v292 = int32(0)
	goto L71
L70:
	;
	v292 = v55
	goto L71
L71:
	;
	v296 = *(*float64)(unsafe.Add(mBase, uint32(v281+v292<<(uint(int32(3))%32))))
	if base.F64_gt(v296, float64(0)) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v307 = float64(0)
	goto L74
L73:
	;
	v303 = *(*float64)(unsafe.Add(mBase, uint32(v281+v55<<(uint(int32(3))%32))))
	if base.F64_lt(v282, v303) != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v280+v246))) = v307
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	v312 = v281 + v55<<(uint(int32(3))%32)
	v313 = *(*float64)(unsafe.Add(mBase, uint32(v312)))
	if base.F64_gt(v309, v313) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v305 = v282
	goto L77
L76:
	;
	v305 = v303
	goto L77
L77:
	;
	v307 = v305
	goto L74
L78:
	;
	v315 = v309
	goto L80
L79:
	;
	v315 = v313
	goto L80
L80:
	;
	v317 = base.F64_lt(v315, float64(0))
	if v317 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v317 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v320 = *(*float64)(unsafe.Add(mBase, uint32(v312)))
	if base.F64_gt(v309, v320) != 0 {
		v334 = v309
		goto L65
	} else {
		goto L87
	}
L84:
	;
	v319 = float64(0)
	goto L86
L85:
	;
	v319 = v309
	goto L86
L86:
	;
	v334 = v319
	goto L65
L87:
	;
	v334 = v320
	goto L65
L88:
	;
	v326 = v323
	goto L90
L89:
	;
	v326 = v282
	goto L90
L90:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v280+v246))) = v326
	v328 = float64(0)
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_lt(v329, v328) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v332 = v328
	goto L93
L92:
	;
	v332 = v329
	goto L93
L93:
	;
	v334 = v332
	goto L65
L94:
	;
	goto L64
L95:
	;
	return v41
L96:
	;
	v377 = v367
	goto L99
L97:
	;
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v36 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v36<<(uint(int32(5))%32) + int32(32)
	goto L95
L99:
	;
	v402 = int32(3)
	v404 = v41 + int32(8) + v377<<(uint(v402)%32)
	v405 = *(*float64)(unsafe.Add(mBase, uint32(v404)))
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v404+v53<<(uint(v402)%32))))
	if base.F64_ne(v405, v409) != 0 {
		goto L95
	} else {
		goto L101
	}
L100:
	;
	goto L98
L101:
	;
	v412 = v377 + int32(1)
	if v412 != v53 {
		v377 = v412
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
}
func F_cube_yy_scan_bytes(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	v4 = int32(0)
	v13 = l1 + int32(2)
	v14 = F_palloc(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			if l1 <= int32(0) {
			} else {
				v21 = l1 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(l1) {
					v29 = v4
					v35 = v4
					for {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v29))))
						*(*uint8)(unsafe.Add(mBase, uint32(v29+v14))) = uint8(v39)
						v42 = v29 | int32(1)
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v42))))
						*(*uint8)(unsafe.Add(mBase, uint32(v14+v42))) = uint8(v45)
						v48 = v29 | int32(2)
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v48))))
						*(*uint8)(unsafe.Add(mBase, uint32(v14+v48))) = uint8(v51)
						v54 = v29 | int32(3)
						v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v54))))
						*(*uint8)(unsafe.Add(mBase, uint32(v14+v54))) = uint8(v57)
						v59 = int32(4)
						v60 = v29 + v59
						v62 = v35 + v59
						if v62 != l1&int32(2147483644) {
							v29 = v60
							v35 = v62
							continue
						} else {
							break
						}
						break
					}
					if v21 == int32(0) {
					} else {
						v69 = v60
						v80 = v69
						v87 = v4
						for {
							v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
							*(*uint8)(unsafe.Add(mBase, uint32(v80+v14))) = uint8(v90)
							v92 = int32(1)
							v95 = v87 + v92
							if v95 != v21 {
								v80 = v80 + v92
								v87 = v95
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v69 = v4
					v80 = v69
					v87 = v4
					for {
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
						*(*uint8)(unsafe.Add(mBase, uint32(v80+v14))) = uint8(v90)
						v92 = int32(1)
						v95 = v87 + v92
						if v95 != v21 {
							v80 = v80 + v92
							v87 = v95
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v109 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v14))) = uint16(v109)
			v111 = F_cube_yy_scan_buffer(m, v14, v13, l2)
			mBase = m.M
			v112 = m.ExcPending
			if v112 != 0 {
				return int32(0)
			} else {
				if v111 == int32(0) {
					F_yy_fatal_error_6(m, int32(_a_F_cube_yy_scan_bytes_0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v111)+20)) = int32(1)
					return v111
				}
			}
		} else {
			F_yy_fatal_error_6(m, int32(_a_F_cube_yy_scan_bytes_1))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_cube_yy_scan_string(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	v3 = int32(0)
	v12 = F_strlen(m, l0)
	mBase = m.M
	v14 = v12 + int32(2)
	v15 = F_palloc(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 != 0 {
			if v12 <= int32(0) {
			} else {
				v22 = v12 & int32(3)
				if base.Ui32(int32(4)) <= base.Ui32(v12) {
					v29 = v3
					v36 = v3
					for {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v29))))
						*(*uint8)(unsafe.Add(mBase, uint32(v29+v15))) = uint8(v40)
						v43 = v29 | int32(1)
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v43))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v43))) = uint8(v46)
						v49 = v29 | int32(2)
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v49))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v49))) = uint8(v52)
						v55 = v29 | int32(3)
						v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v55))))
						*(*uint8)(unsafe.Add(mBase, uint32(v15+v55))) = uint8(v58)
						v60 = int32(4)
						v61 = v29 + v60
						v63 = v36 + v60
						if v63 != v12&int32(2147483644) {
							v29 = v61
							v36 = v63
							continue
						} else {
							break
						}
						break
					}
					if v22 == int32(0) {
					} else {
						v69 = v61
						v80 = v69
						v88 = v3
						for {
							v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
							*(*uint8)(unsafe.Add(mBase, uint32(v80+v15))) = uint8(v91)
							v93 = int32(1)
							v96 = v88 + v93
							if v96 != v22 {
								v80 = v80 + v93
								v88 = v96
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v69 = v3
					v80 = v69
					v88 = v3
					for {
						v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v80))))
						*(*uint8)(unsafe.Add(mBase, uint32(v80+v15))) = uint8(v91)
						v93 = int32(1)
						v96 = v88 + v93
						if v96 != v22 {
							v80 = v80 + v93
							v88 = v96
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v110 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v15+v12))) = uint16(v110)
			v112 = F_cube_yy_scan_buffer(m, v15, v14, l1)
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				if v112 == int32(0) {
					F_yy_fatal_error_6(m, int32(_a_F_cube_yy_scan_string_0))
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = int32(1)
					return v112
				}
			}
		} else {
			F_yy_fatal_error_6(m, int32(_a_F_cube_yy_scan_string_1))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_cube_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_6(m, int32(_a_F_cube_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_6(m, int32(_a_F_cube_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_cube_yyget_lval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	return v2
}
func F_cube_yyget_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	return v2
}
func F_cube_yyset_debug(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = l0
	return
}
func F_cube_yyset_in(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	return
}
