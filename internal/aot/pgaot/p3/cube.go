package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_contains(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
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
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
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
	v13 = int32(0)
	if base.B2i32(v6 == v13)|base.B2i32(v11 == v13) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v184 != v6 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v183 = v165
	goto L4
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v31 = int32(2147483647)
	v32 = v30 & v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v35 = v33 & v31
	if base.Ui32(v32) < base.Ui32(v35) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v32
	goto L10
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(v32) < base.Ui32(v35) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v57 = v11 + int32(8) + v43<<(uint(int32(3))%32)
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
	if base.F64_ne(v58, float64(0)) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if base.B2i32(v33 < int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v57+v35<<(uint(int32(3))%32))))
	if base.F64_ne(v66, float64(0)) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v70 = v43 + int32(1)
	if v70 != v35 {
		v43 = v70
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L11
L18:
	;
	v87 = v32
	goto L20
L19:
	;
	v87 = v35
	goto L20
L20:
	;
	if v87 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v183 = int32(1)
	goto L4
L22:
	;
	goto L23
L23:
	;
	v91 = int32(8)
	v98 = int32(0)
	goto L24
L24:
	;
	v110 = int32(0)
	v112 = v98 << (uint(int32(3)) % 32)
	v113 = v6 + v91 + v112
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v116 = base.B2i32(v30 < v110)
	if v30 < v110 {
		v123 = v114
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v165 = v151
	goto L5
L26:
	;
	v124 = v112 + (v11 + v91)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	v127 = base.B2i32(v33 < int32(0))
	if v33 < int32(0) {
		v134 = v125
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v113+v30<<(uint(int32(3))%32))))
	if base.F64_lt(v114, v120) != 0 {
		v123 = v114
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v123 = v120
	goto L26
L29:
	;
	if base.F64_gt(v123, v134) != 0 {
		v165 = v110
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v124+v33<<(uint(int32(3))%32))))
	if base.F64_lt(v125, v131) != 0 {
		v134 = v125
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v134 = v131
	goto L29
L32:
	;
	if v30 < v110 {
		v141 = v114
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v33 < int32(0) {
		v148 = v125
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v113+v30<<(uint(int32(3))%32))))
	if base.F64_gt(v114, v139) != 0 {
		v141 = v114
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v141 = v139
	goto L33
L36:
	;
	if base.F64_gt(v148, v141) != 0 {
		v165 = v110
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v124+v33<<(uint(int32(3))%32))))
	if base.F64_gt(v125, v146) != 0 {
		v148 = v125
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v148 = v146
	goto L36
L39:
	;
	v151 = int32(1)
	v153 = v98 + v151
	if v153 != v87 {
		v98 = v153
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L25
L41:
	;
	F_pfree(m, v6)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v188 != v11 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v11)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return v183
L48:
	;
	goto L47
}
func F_cube_gt(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(int32(0) < v496)
L121:
	;
	goto L120
}
func F_cube_is_point(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v34 float64
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v47 != v9 {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v14 = int32(0)
	if base.B2i32(v13 < v14)|base.B2i32(v13 == v14) != 0 {
		v42 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(0)
	goto L5
L5:
	;
	v27 = int32(3)
	v29 = v9 + int32(8) + v25<<(uint(v27)%32)
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
	v34 = *(*float64)(unsafe.Add(mBase, uint32(v29+v13<<(uint(v27)%32))))
	if base.F64_eq(v30, v34) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v42 = int32(0)
	goto L1
L7:
	;
	v36 = int32(1)
	v38 = v25 + v36
	if v13 != v38 {
		v25 = v38
		goto L5
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	goto L6
L10:
	;
	v42 = v36
	goto L1
L11:
	;
	F_pfree(m, v9)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	return v42
L14:
	;
	goto L13
}
func F_cube_ne(m *base.Module, l0 int32) int32 {
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
	return base.B2i32(v496 != int32(0))
L121:
	;
	goto L120
}
func F_cube_overlap(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v72 int32
	_ = v72
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v83 int32
	_ = v83
	var v87 float64
	_ = v87
	var v90 float64
	_ = v90
	var v95 float64
	_ = v95
	var v97 float64
	_ = v97
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v151 float64
	_ = v151
	var v154 int32
	_ = v154
	var v159 float64
	_ = v159
	var v161 int32
	_ = v161
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v177 float64
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
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
	v16 = int32(0)
	if base.B2i32(v6 == v16)|base.B2i32(v11 == v16) != 0 {
		v196 = v16
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v219 != v6 {
		goto L52
	} else {
		goto L53
	}
L5:
	;
	v218 = v196
	goto L4
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v32 = int32(2147483647)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v37 = base.B2i32(base.Ui32(v31&v32) < base.Ui32(v34&v32))
	if base.Ui32(v31&v32) < base.Ui32(v34&v32) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v11
	goto L9
L8:
	;
	v38 = v6
	goto L9
L9:
	;
	if base.Ui32(v31&v32) < base.Ui32(v34&v32) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v126 = v110 & int32(2147483647)
	if base.Ui32(v126) <= base.Ui32(v42) {
		goto L34
	} else {
		goto L35
	}
L11:
	;
	v39 = v6
	goto L13
L12:
	;
	v39 = v11
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v42 = v40 & int32(2147483647)
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v110 = v45
	goto L10
L15:
	;
	goto L16
L16:
	;
	v46 = int32(8)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v53 = int32(0)
	goto L17
L17:
	;
	v68 = v53 << (uint(int32(3)) % 32)
	v69 = v38 + v46 + v68
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
	v72 = base.B2i32(v50 < int32(0))
	if v50 < int32(0) {
		v79 = v70
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v110 = v50
	goto L10
L19:
	;
	v80 = v68 + (v39 + v46)
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v80)))
	v83 = base.B2i32(v40 < int32(0))
	if v40 < int32(0) {
		v90 = v81
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v69+v50<<(uint(int32(3))%32))))
	if base.F64_lt(v70, v76) != 0 {
		v79 = v70
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = v76
	goto L19
L22:
	;
	if base.F64_gt(v79, v90) != 0 {
		v196 = v16
		goto L5
	} else {
		goto L25
	}
L23:
	;
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v80+v40<<(uint(int32(3))%32))))
	if base.F64_gt(v81, v87) != 0 {
		v90 = v81
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v90 = v87
	goto L22
L25:
	;
	if v50 < int32(0) {
		v97 = v70
		goto L26
	} else {
		goto L27
	}
L26:
	;
	if v40 < int32(0) {
		v104 = v81
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v69+v50<<(uint(int32(3))%32))))
	if base.F64_gt(v70, v95) != 0 {
		v97 = v70
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v97 = v95
	goto L26
L29:
	;
	if base.F64_gt(v104, v97) != 0 {
		v196 = v16
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v80+v40<<(uint(int32(3))%32))))
	if base.F64_lt(v81, v102) != 0 {
		v104 = v81
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v104 = v102
	goto L29
L32:
	;
	v108 = v53 + int32(1)
	if v108 != v42 {
		v53 = v108
		goto L17
	} else {
		goto L33
	}
L33:
	;
	goto L18
L34:
	;
	v218 = int32(1)
	goto L4
L35:
	;
	goto L36
L36:
	;
	v138 = v42
	goto L37
L37:
	;
	v150 = v38 + int32(8) + v138<<(uint(int32(3))%32)
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v150)))
	if base.B2i32(v110 < int32(0)) == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v196 = int32(0)
	goto L5
L39:
	;
	goto L38
L40:
	;
	if base.F64_lt(v177, float64(0)) != 0 {
		goto L39
	} else {
		goto L50
	}
L41:
	;
	v154 = int32(0)
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v150+v126<<(uint(int32(3))%32))))
	if base.F64_lt(v151, v159) != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	if base.F64_gt(v151, float64(0)) != 0 {
		goto L39
	} else {
		goto L49
	}
L44:
	;
	v161 = v154
	goto L46
L45:
	;
	v161 = v110
	goto L46
L46:
	;
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v150+v161<<(uint(int32(3))%32))))
	if base.F64_gt(v165, float64(0)) != 0 {
		v196 = v154
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v150+v110<<(uint(int32(3))%32))))
	if base.F64_gt(v151, v171) == int32(0) {
		v177 = v171
		goto L40
	} else {
		goto L48
	}
L48:
	;
	v177 = v151
	goto L40
L49:
	;
	v177 = v151
	goto L40
L50:
	;
	v181 = int32(1)
	v183 = v138 + v181
	if v126 != v183 {
		v138 = v183
		goto L37
	} else {
		goto L51
	}
L51:
	;
	v196 = v181
	goto L5
L52:
	;
	F_pfree(m, v6)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v223 != v11 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_pfree(m, v11)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	return v218
L59:
	;
	goto L58
}
func F_cube_yy_delete_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v4 == int32(0) {
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v10 = v4 + v7<<(uint(int32(2))%32)
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if l0 != v11 {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(0)
			}
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16 == int32(0) {
			F_pfree(m, l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				return
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v19 == int32(0) {
				F_pfree(m, l0)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					return
				}
			} else {
				F_pfree(m, v19)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
func F_cube_yy_scan_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v4 = int32(0)
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v77 = v4
		return v77
	} else {
		v9 = l1 - int32(2)
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v9))))
		if v11 != 0 {
			v77 = v4
			return v77
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+l1-int32(1)))))
			if v15 != 0 {
				v77 = v4
				return v77
			} else {
				v17 = F_palloc(m, int32(48))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v17 == int32(0) {
						F_yy_fatal_error_6(m, int32(_a_F_cube_yy_scan_buffer_0))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v23 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v23
						*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v9
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(4294967296)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v9
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v23
						F_cube_yyensure_buffer_stack(m, l2)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v37+v38<<(uint(int32(2))%32))))
							if v42 == v17 {
								v77 = v17
							} else {
								if v42 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)))
									*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v45)
									v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
									v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v47+v48<<(uint(int32(2))%32))))
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+36))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v53
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
									*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v55
									v57 = v48
									v58 = v47
								} else {
									v57 = v38
									v58 = v37
								}
								*(*int32)(unsafe.Add(mBase, uint32(v58+v57<<(uint(int32(2))%32)))) = v17
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v64
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v66
								*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v66
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v69
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v71)
								v77 = v17
							}
							return v77
						}
					}
				}
			}
		}
	}
}
func F_cube_yyget_column(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == v2 {
		v16 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4+v7<<(uint(int32(2))%32))))
		if v11 == int32(0) {
			v16 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
			v16 = v14
		}
	}
	return v16
}
func F_cube_yyget_leng(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	return v2
}
func F_cube_yyparse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
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
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
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
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
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
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v843 int32
	_ = v843
	var v850 int32
	_ = v850
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v909 int32
	_ = v909
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v971 int32
	_ = v971
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v1001 int32
	_ = v1001
	v5 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(1104)
	m.G0 = v25
	*(*int32)(unsafe.Add(mBase, uint32(v25)+1100)) = v5
	v33 = v25 + int32(96)
	v35 = v25 + int32(896)
	v41 = v5
	v43 = v35
	v44 = v33
	v48 = v35
	v49 = int32(-2)
	v51 = int32(200)
	v54 = v33
	goto L1
L1:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v43))) = uint8(v41)
	if base.Ui32(v48+v51-int32(1)) <= base.Ui32(v43) {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	if v25+int32(896) != v986 {
		goto L236
	} else {
		goto L237
	}
L3:
	;
	goto L2
L4:
	;
	v41 = base.I32_extend8_s(v971)
	v43 = v955 + int32(1)
	v44 = v956
	v48 = v112
	v49 = v961
	v51 = v113
	v54 = v114
	goto L1
L5:
	;
	v920 = v110 - v175<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v920)+4)) = v909
	v923 = v920 + int32(4)
	v924 = v109 - v175
	v925 = int32(*(*int8)(unsafe.Add(mBase, uint32(v924))))
	v928 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_cube_yyparse[0]))))
	v930 = v928 - int32(9)
	v933 = int32(*(*int8)(unsafe.Add(mBase, uint32(v930)+uint32(_c_F_cube_yyparse[1]))))
	v934 = v925 + v933
	if base.Ui32(v934) <= base.Ui32(int32(17)) {
		goto L232
	} else {
		goto L233
	}
L6:
	;
	v893 = F_write_point_as_box(m, v875, v589, l0, l2)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L17
	} else {
		goto L229
	}
L7:
	;
	v868 = F_write_point_as_box(m, v850, v665, l0, l2)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L17
	} else {
		goto L227
	}
L8:
	;
	v981 = int32(1)
	v986 = v76
	goto L3
L9:
	;
	F_cube_yyerror(m, l0, l1, l2, l3, int32(_a_F_cube_yyparse_0))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L17
	} else {
		goto L226
	}
L10:
	;
	if int32(_a_F_cube_yyparse_1) < v51 {
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v109 = v43
	v110 = v44
	v112 = v48
	v113 = v51
	v114 = v54
	goto L12
L12:
	;
	if v41 == int32(10) {
		goto L31
	} else {
		goto L32
	}
L13:
	;
	v66 = int32(_a_F_cube_yyparse_2)
	v68 = v51 << (uint(int32(1)) % 32)
	if v66 <= v68 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = v66
	goto L16
L15:
	;
	v71 = v68
	goto L16
L16:
	;
	v76 = F_palloc(m, v71*int32(5)+int32(3))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v76 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v82 = v43 - v48
	v84 = v82 + int32(1)
	if v84 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	base.MemoryCopy(m, v76, v48, v84)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v89 = base.I32_div_s(v71+int32(3), int32(4))
	v90 = int32(2)
	v92 = v76 + v89<<(uint(v90)%32)
	v94 = v84 << (uint(v90) % 32)
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	base.MemoryCopy(m, v92, v54, v94)
	goto L25
L24:
	;
	goto L25
L25:
	;
	if v25+int32(896) != v48 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_pfree(m, v48)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L17
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v71-int32(1) <= v82 {
		goto L8
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v109 = v76 + v82
	v110 = v94 + v92 - int32(4)
	v112 = v76
	v113 = v71
	v114 = v92
	goto L12
L31:
	;
	v981 = int32(0)
	v986 = v112
	goto L3
L32:
	;
	goto L33
L33:
	;
	v119 = int32(1) << (uint(v41) % 32)
	if v119&int32(_a_F_cube_yyparse_3) != 0 {
		v163 = v49
		goto L34
	} else {
		goto L35
	}
L34:
	;
	if v119&int32(_a_F_cube_yyparse_4) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L35:
	;
	v124 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_cube_yyparse[2]))))
	if v49 == int32(-2) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v147 = v124 + v146
	if base.Ui32(int32(17)) < base.Ui32(v147) {
		v163 = v145
		goto L34
	} else {
		goto L48
	}
L37:
	;
	v129 = F_cube_yylex(m, v25+int32(1100), l3)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L17
	} else {
		goto L40
	}
L38:
	;
	v131 = v49
	goto L39
L39:
	;
	if v131 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v131 = v129
	goto L39
L41:
	;
	v134 = int32(0)
	v145 = v134
	v146 = v134
	goto L36
L42:
	;
	goto L43
L43:
	;
	if v131 == int32(256) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L45:
	;
	goto L46
L46:
	;
	if base.Ui32(int32(263)) < base.Ui32(v131) {
		v145 = v131
		v146 = int32(2)
		goto L36
	} else {
		goto L47
	}
L47:
	;
	v144 = int32(*(*int8)(unsafe.Add(mBase, uint32(v131)+uint32(_c_F_cube_yyparse[3]))))
	v145 = v131
	v146 = v144
	goto L36
L48:
	;
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_cube_yyparse[4]))))
	if v146 != v152 {
		v163 = v145
		goto L34
	} else {
		goto L49
	}
L49:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v25)+1100))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v154
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_cube_yyparse[5]))))
	v955 = v109
	v956 = v110 + int32(4)
	v961 = int32(-2)
	v971 = v161
	goto L4
L50:
	;
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_cube_yyparse[6]))))
	v175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v172)+uint32(_c_F_cube_yyparse[7]))))
	v177 = int32(2)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v110+(int32(1)-v175)<<(uint(v177)%32))))
	switch v172&int32(255) - v177 {
	case 0:
		goto L60
	case 1:
		goto L59
	case 2:
		goto L58
	case 3:
		goto L57
	case 4:
		goto L56
	case 5:
		goto L55
	case 6:
		goto L54
	case 7:
		goto L53
	default:
		v909 = v180
		goto L5
	}
L51:
	;
	goto L52
L52:
	;
	F_cube_yyerror(m, l0, l1, l2, l3, int32(_a_F_cube_yyparse_5))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L17
	} else {
		goto L225
	}
L53:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v110-int32(8))))
	v827 = F_strlen(m, v826)
	mBase = m.M
	v829 = int32(44)
	*(*uint16)(unsafe.Add(mBase, uint32(v827+v826))) = uint16(v829)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v832 = F_strlen(m, v826)
	mBase = m.M
	v834 = F_strcpy(m, v832+v826, v831)
	mBase = m.M
	goto L224
L54:
	;
	v747 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L17
	} else {
		goto L202
	}
L55:
	;
	v745 = F_pstrdup(m, int32(_a_F_cube_yyparse_6))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L17
	} else {
		goto L201
	}
L56:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v110-int32(4))))
	v909 = v743
	goto L5
L57:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v666 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v666 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L58:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v590 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L59:
	;
	v387 = int32(0)
	v390 = v110 - int32(8)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v391))))
	if v392 == v387 {
		v444 = v387
		goto L109
	} else {
		goto L110
	}
L60:
	;
	v185 = int32(0)
	v188 = v110 - int32(12)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v190 == v185 {
		v242 = v185
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v261 = v110 - int32(4)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262))))
	if v263 == int32(0) {
		v319 = v185
		goto L75
	} else {
		goto L76
	}
L62:
	;
	v193 = int32(1)
	v194 = int32(44)
	v195 = F___strchrnul(m, v189, v194)
	mBase = m.M
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v197 == v194 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v201 == int32(0) {
		v242 = v193
		goto L61
	} else {
		goto L67
	}
L64:
	;
	v201 = v195
	goto L66
L65:
	;
	v201 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	v208 = v193
	v209 = v201
	goto L68
L68:
	;
	v226 = int32(1)
	v227 = v208 + v226
	v230 = int32(44)
	v231 = F___strchrnul(m, v209+v226, v230)
	mBase = m.M
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231))))
	if v233 == v230 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v242 = v227
	goto L61
L70:
	;
	if v237 != 0 {
		v208 = v227
		v209 = v237
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v237 = v231
	goto L73
L72:
	;
	v237 = int32(0)
	goto L73
L73:
	;
	goto L70
L74:
	;
	goto L69
L75:
	;
	if v242 != v319 {
		goto L89
	} else {
		goto L90
	}
L76:
	;
	v266 = int32(1)
	v267 = int32(44)
	v268 = F___strchrnul(m, v262, v267)
	mBase = m.M
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v268))))
	if v270 == v267 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	if v274 == int32(0) {
		v319 = v266
		goto L75
	} else {
		goto L81
	}
L78:
	;
	v274 = v268
	goto L80
L79:
	;
	v274 = int32(0)
	goto L80
L80:
	;
	goto L77
L81:
	;
	v282 = v274
	v285 = v266
	goto L82
L82:
	;
	v299 = int32(1)
	v300 = v285 + v299
	v303 = int32(44)
	v304 = F___strchrnul(m, v282+v299, v303)
	mBase = m.M
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v304))))
	if v306 == v303 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v319 = v300
	goto L75
L84:
	;
	if v310 != 0 {
		v282 = v310
		v285 = v300
		goto L82
	} else {
		goto L88
	}
L85:
	;
	v310 = v304
	goto L87
L86:
	;
	v310 = int32(0)
	goto L87
L87:
	;
	goto L84
L88:
	;
	goto L83
L89:
	;
	v334 = int32(1)
	v335 = F_errsave_start(m, l2)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if int32(101) <= v242 {
		goto L98
	} else {
		goto L99
	}
L92:
	;
	if v335 == int32(0) {
		v981 = v334
		v986 = v112
		goto L3
	} else {
		goto L93
	}
L93:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L17
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L17
	} else {
		goto L95
	}
L95:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v346
	F_errdetail(m, int32(_a_F_cube_yyparse_8), v25+int32(16))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L17
	} else {
		goto L96
	}
L96:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(58), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L17
	} else {
		goto L97
	}
L97:
	;
	v981 = v334
	v986 = v112
	goto L3
L98:
	;
	v362 = int32(1)
	v363 = F_errsave_start(m, l2)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L17
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v384 = F_write_box(m, v242, v189, v262, l0, l2)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L17
	} else {
		goto L107
	}
L101:
	;
	if v363 == int32(0) {
		v981 = v362
		v986 = v112
		goto L3
	} else {
		goto L102
	}
L102:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L17
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L17
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(100)
	F_errdetail(m, int32(_a_F_cube_yyparse_11), v25)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L17
	} else {
		goto L105
	}
L105:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(67), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L17
	} else {
		goto L106
	}
L106:
	;
	v981 = v362
	v986 = v112
	goto L3
L107:
	;
	if v384 != 0 {
		v909 = v180
		goto L5
	} else {
		goto L108
	}
L108:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L109:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462))))
	if v463 == int32(0) {
		v519 = v387
		goto L123
	} else {
		goto L124
	}
L110:
	;
	v395 = int32(1)
	v396 = int32(44)
	v397 = F___strchrnul(m, v391, v396)
	mBase = m.M
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397))))
	if v399 == v396 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v403 == int32(0) {
		v444 = v395
		goto L109
	} else {
		goto L115
	}
L112:
	;
	v403 = v397
	goto L114
L113:
	;
	v403 = int32(0)
	goto L114
L114:
	;
	goto L111
L115:
	;
	v410 = v395
	v411 = v403
	goto L116
L116:
	;
	v428 = int32(1)
	v429 = v410 + v428
	v432 = int32(44)
	v433 = F___strchrnul(m, v411+v428, v432)
	mBase = m.M
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433))))
	if v435 == v432 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	v444 = v429
	goto L109
L118:
	;
	if v439 != 0 {
		v410 = v429
		v411 = v439
		goto L116
	} else {
		goto L122
	}
L119:
	;
	v439 = v433
	goto L121
L120:
	;
	v439 = int32(0)
	goto L121
L121:
	;
	goto L118
L122:
	;
	goto L117
L123:
	;
	if v444 != v519 {
		goto L137
	} else {
		goto L138
	}
L124:
	;
	v466 = int32(1)
	v467 = int32(44)
	v468 = F___strchrnul(m, v462, v467)
	mBase = m.M
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v470 == v467 {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v474 == int32(0) {
		v519 = v466
		goto L123
	} else {
		goto L129
	}
L126:
	;
	v474 = v468
	goto L128
L127:
	;
	v474 = int32(0)
	goto L128
L128:
	;
	goto L125
L129:
	;
	v482 = v474
	v485 = v466
	goto L130
L130:
	;
	v499 = int32(1)
	v500 = v485 + v499
	v503 = int32(44)
	v504 = F___strchrnul(m, v482+v499, v503)
	mBase = m.M
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v504))))
	if v506 == v503 {
		goto L133
	} else {
		goto L134
	}
L131:
	;
	v519 = v500
	goto L123
L132:
	;
	if v510 != 0 {
		v482 = v510
		v485 = v500
		goto L130
	} else {
		goto L136
	}
L133:
	;
	v510 = v504
	goto L135
L134:
	;
	v510 = int32(0)
	goto L135
L135:
	;
	goto L132
L136:
	;
	goto L131
L137:
	;
	v534 = int32(1)
	v535 = F_errsave_start(m, l2)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L17
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	if int32(101) <= v444 {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	if v535 == int32(0) {
		v981 = v534
		v986 = v112
		goto L3
	} else {
		goto L141
	}
L141:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L17
	} else {
		goto L142
	}
L142:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L17
	} else {
		goto L143
	}
L143:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v547
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = v546
	F_errdetail(m, int32(_a_F_cube_yyparse_8), v25+int32(48))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L17
	} else {
		goto L144
	}
L144:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(88), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L17
	} else {
		goto L145
	}
L145:
	;
	v981 = v534
	v986 = v112
	goto L3
L146:
	;
	v562 = int32(1)
	v563 = F_errsave_start(m, l2)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L17
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v586 = F_write_box(m, v444, v391, v462, l0, l2)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L17
	} else {
		goto L155
	}
L149:
	;
	if v563 == int32(0) {
		v981 = v562
		v986 = v112
		goto L3
	} else {
		goto L150
	}
L150:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L17
	} else {
		goto L151
	}
L151:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L17
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(100)
	F_errdetail(m, int32(_a_F_cube_yyparse_11), v25+int32(32))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L17
	} else {
		goto L153
	}
L153:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(97), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L17
	} else {
		goto L154
	}
L154:
	;
	v981 = v562
	v986 = v112
	goto L3
L155:
	;
	if v586 != 0 {
		v909 = v180
		goto L5
	} else {
		goto L156
	}
L156:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L157:
	;
	v875 = int32(0)
	goto L6
L158:
	;
	goto L159
L159:
	;
	v594 = int32(1)
	v595 = int32(44)
	v596 = F___strchrnul(m, v589, v595)
	mBase = m.M
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v596))))
	if v598 == v595 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	if v602 == int32(0) {
		v875 = v594
		goto L6
	} else {
		goto L164
	}
L161:
	;
	v602 = v596
	goto L163
L162:
	;
	v602 = int32(0)
	goto L163
L163:
	;
	goto L160
L164:
	;
	v609 = v594
	v613 = v602
	goto L165
L165:
	;
	v627 = int32(1)
	v628 = v609 + v627
	v631 = int32(44)
	v632 = F___strchrnul(m, v613+v627, v631)
	mBase = m.M
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632))))
	if v634 == v631 {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if base.Ui32(v609) < base.Ui32(int32(100)) {
		v875 = v628
		goto L6
	} else {
		goto L172
	}
L167:
	;
	if v638 != 0 {
		v609 = v628
		v613 = v638
		goto L165
	} else {
		goto L171
	}
L168:
	;
	v638 = v632
	goto L170
L169:
	;
	v638 = int32(0)
	goto L170
L170:
	;
	goto L167
L171:
	;
	goto L166
L172:
	;
	v641 = int32(1)
	v642 = F_errsave_start(m, l2)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L17
	} else {
		goto L173
	}
L173:
	;
	if v642 == int32(0) {
		v981 = v641
		v986 = v112
		goto L3
	} else {
		goto L174
	}
L174:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L17
	} else {
		goto L175
	}
L175:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L17
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = int32(100)
	F_errdetail(m, int32(_a_F_cube_yyparse_11), v25-int32(-64))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L17
	} else {
		goto L177
	}
L177:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(116), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L17
	} else {
		goto L178
	}
L178:
	;
	v981 = v641
	v986 = v112
	goto L3
L179:
	;
	v850 = int32(0)
	goto L7
L180:
	;
	goto L181
L181:
	;
	v670 = int32(1)
	v671 = int32(44)
	v672 = F___strchrnul(m, v665, v671)
	mBase = m.M
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672))))
	if v674 == v671 {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	if v678 == int32(0) {
		v850 = v670
		goto L7
	} else {
		goto L186
	}
L183:
	;
	v678 = v672
	goto L185
L184:
	;
	v678 = int32(0)
	goto L185
L185:
	;
	goto L182
L186:
	;
	v685 = v670
	v689 = v678
	goto L187
L187:
	;
	v703 = int32(1)
	v704 = v685 + v703
	v707 = int32(44)
	v708 = F___strchrnul(m, v689+v703, v707)
	mBase = m.M
	v710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	if v710 == v707 {
		goto L190
	} else {
		goto L191
	}
L188:
	;
	if base.Ui32(v685) < base.Ui32(int32(100)) {
		v850 = v704
		goto L7
	} else {
		goto L194
	}
L189:
	;
	if v714 != 0 {
		v685 = v704
		v689 = v714
		goto L187
	} else {
		goto L193
	}
L190:
	;
	v714 = v708
	goto L192
L191:
	;
	v714 = int32(0)
	goto L192
L192:
	;
	goto L189
L193:
	;
	goto L188
L194:
	;
	v717 = int32(1)
	v718 = F_errsave_start(m, l2)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L17
	} else {
		goto L195
	}
L195:
	;
	if v718 == int32(0) {
		v981 = v717
		v986 = v112
		goto L3
	} else {
		goto L196
	}
L196:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L17
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_cube_yyparse_7), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L17
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = int32(100)
	F_errdetail(m, int32(_a_F_cube_yyparse_11), v25+int32(80))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L17
	} else {
		goto L199
	}
L199:
	;
	F_errsave_finish(m, l2, int32(_a_F_cube_yyparse_9), int32(135), int32(_a_F_cube_yyparse_10))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L17
	} else {
		goto L200
	}
L200:
	;
	v981 = v717
	v986 = v112
	goto L3
L201:
	;
	v909 = v745
	goto L5
L202:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if (v749^v747)&int32(3) != 0 {
		goto L206
	} else {
		goto L207
	}
L203:
	;
	v909 = v747
	goto L5
L204:
	;
	goto L203
L205:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v804))) = uint8(v803)
	if v803&int32(255) == int32(0) {
		goto L204
	} else {
		goto L220
	}
L206:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749))))
	v802 = v749
	v803 = v755
	v804 = v747
	goto L205
L207:
	;
	goto L208
L208:
	;
	if v749&int32(3) != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v759 = v749
	v761 = v747
	goto L212
L210:
	;
	v773 = v749
	v775 = v747
	goto L211
L211:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	v780 = int32(-2139062144)
	if (int32(16843008)-v777|v777)&v780 != v780 {
		v802 = v773
		v803 = v777
		v804 = v775
		goto L205
	} else {
		goto L216
	}
L212:
	;
	v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v759))))
	*(*uint8)(unsafe.Add(mBase, uint32(v761))) = uint8(v762)
	if v762 == int32(0) {
		goto L204
	} else {
		goto L214
	}
L213:
	;
	v773 = v769
	v775 = v767
	goto L211
L214:
	;
	v766 = int32(1)
	v767 = v761 + v766
	v769 = v759 + v766
	if v769&int32(3) != 0 {
		v759 = v769
		v761 = v767
		goto L212
	} else {
		goto L215
	}
L215:
	;
	goto L213
L216:
	;
	v785 = v773
	v786 = v777
	v787 = v775
	goto L217
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v787))) = v786
	v789 = int32(4)
	v790 = v787 + v789
	v792 = v785 + v789
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v785)+4))
	v797 = int32(-2139062144)
	if (int32(16843008)-v794|v794)&v797 == v797 {
		v785 = v792
		v786 = v794
		v787 = v790
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v802 = v792
	v803 = v794
	v804 = v790
	goto L205
L219:
	;
	goto L218
L220:
	;
	v811 = v802
	v813 = v804
	goto L221
L221:
	;
	v814 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v813)+1)) = uint8(v814)
	v816 = int32(1)
	if v814 != 0 {
		v811 = v811 + v816
		v813 = v813 + v816
		goto L221
	} else {
		goto L223
	}
L222:
	;
	goto L204
L223:
	;
	goto L222
L224:
	;
	v909 = v826
	goto L5
L225:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L226:
	;
	v981 = int32(2)
	v986 = v48
	goto L3
L227:
	;
	if v868 != 0 {
		v909 = v180
		goto L5
	} else {
		goto L228
	}
L228:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L229:
	;
	if v893 != 0 {
		v909 = v180
		goto L5
	} else {
		goto L230
	}
L230:
	;
	v981 = int32(1)
	v986 = v112
	goto L3
L231:
	;
	v948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+uint32(_c_F_cube_yyparse[5]))))
	v955 = v924
	v956 = v923
	v961 = v163
	v971 = v948
	goto L4
L232:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v934)+uint32(_c_F_cube_yyparse[4]))))
	if v939 == v925&int32(255) {
		goto L231
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930)+uint32(_c_F_cube_yyparse[8]))))
	v955 = v924
	v956 = v923
	v961 = v163
	v971 = v945
	goto L4
L235:
	;
	goto L234
L236:
	;
	F_pfree(m, v986)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L17
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	m.G0 = v25 + int32(1104)
	return v981
L239:
	;
	goto L238
}
func F_cube_yypush_buffer_state(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	if l0 != 0 {
		F_cube_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
			if v13 != 0 {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v15)
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v17+v18<<(uint(int32(2))%32))))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v25
				v28 = v18 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v28
				v30 = v17
				v31 = v28
			} else {
				v30 = v8
				v31 = v9
			}
			*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = l0
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			return
		}
	} else {
		return
	}
}
func F_cube_yyrestart(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v8+v9<<(uint(int32(2))%32))))
		if v13 != 0 {
			v27 = v13
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0]))
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
			v40 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
			if v50 == v27 {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
				v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				v70 = v40
				v72 = int32(40)
			} else {
				v63 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
				v70 = int32(0)
				v72 = int32(36)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0])) = v29
			v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
			v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
			*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
			v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
			return
		} else {
			F_cube_yyensure_buffer_stack(m, l1)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v19 = F_cube_yy_create_buffer(m, v17, int32(_a_F_cube_yyrestart_0), l1)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
					v27 = v19
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0]))
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
					if v50 == v27 {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						v70 = v40
						v72 = int32(40)
					} else {
						v63 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
						*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
						v70 = int32(0)
						v72 = int32(36)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
					*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0])) = v29
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
					v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
					return
				}
			}
		}
	} else {
		F_cube_yyensure_buffer_stack(m, l1)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = F_cube_yy_create_buffer(m, v17, int32(_a_F_cube_yyrestart_0), l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v21+v22<<(uint(int32(2))%32)))) = v19
				v27 = v19
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0]))
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v30
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v32))) = uint8(v30)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)) = uint8(v30)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v30
				v40 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v43
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v45+v46<<(uint(int32(2))%32))))
				if v50 == v27 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v52
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v57
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v59)
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					v70 = v40
					v72 = int32(40)
				} else {
					v63 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
					*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v63
					v70 = int32(0)
					v72 = int32(36)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v72+v27))) = v70
				*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_cube_yyrestart[0])) = v29
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v79
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v81
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v81
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v84
				v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v86)
				return
			}
		}
	}
}
func F_cube_yyset_column(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l0
			return
		} else {
			F_yy_fatal_error_6(m, int32(_a_F_cube_yyset_column_0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_yy_fatal_error_6(m, int32(_a_F_cube_yyset_column_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cube_yyset_out(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = l0
	return
}
