package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_contains_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v47 float64
	_ = v47
	var v56 float64
	_ = v56
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	var v106 int32
	_ = v106
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v116 float64
	_ = v116
	var v118 int32
	_ = v118
	var v123 float64
	_ = v123
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v133 float64
	_ = v133
	var v135 float64
	_ = v135
	var v141 float64
	_ = v141
	var v143 float64
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	v3 = int32(0)
	if l0 == v3 {
		v154 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v154
L2:
	;
	if l1 == int32(0) {
		v154 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = int32(2147483647)
	v21 = v19 & v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v24 = v22 & v20
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = l1 + int32(8)
	v32 = v21
	goto L7
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v27+v32<<(uint(int32(3))%32))))
	if base.F64_ne(v47, float64(0)) != 0 {
		v154 = v3
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if base.B2i32(v22 < int32(0)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v27+(v32+v24)<<(uint(int32(3))%32))))
	if base.F64_ne(v56, float64(0)) != 0 {
		v154 = v3
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v60 = v32 + int32(1)
	if v60 != v24 {
		v32 = v60
		goto L7
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L8
L15:
	;
	v77 = v21
	goto L17
L16:
	;
	v77 = v24
	goto L17
L17:
	;
	if v77 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	v82 = int32(8)
	v83 = l1 + v82
	v85 = l0 + v82
	v89 = int32(0)
	goto L21
L21:
	;
	v102 = v89 << (uint(int32(3)) % 32)
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v85+v102)))
	v106 = base.B2i32(v19 < int32(0))
	if v19 < int32(0) {
		v114 = v104
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v154 = v146
	goto L1
L23:
	;
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v83+v102)))
	v118 = base.B2i32(v22 < int32(0))
	if v22 < int32(0) {
		v126 = v116
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v85+(v89+v19)<<(uint(int32(3))%32))))
	if base.F64_lt(v104, v111) != 0 {
		v114 = v104
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v114 = v111
	goto L23
L26:
	;
	v127 = int32(0)
	if base.F64_gt(v114, v126) != 0 {
		v154 = v127
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v123 = *(*float64)(unsafe.Add(mBase, uint32(v83+(v89+v22)<<(uint(int32(3))%32))))
	if base.F64_lt(v116, v123) != 0 {
		v126 = v116
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v126 = v123
	goto L26
L29:
	;
	if v19 < int32(0) {
		v135 = v104
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v22 < int32(0) {
		v143 = v116
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v85+(v89+v19)<<(uint(int32(3))%32))))
	if base.F64_gt(v104, v133) != 0 {
		v135 = v104
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v135 = v133
	goto L30
L33:
	;
	if base.F64_gt(v143, v135) != 0 {
		v154 = v127
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v83+(v89+v22)<<(uint(int32(3))%32))))
	if base.F64_gt(v116, v141) != 0 {
		v143 = v116
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v143 = v141
	goto L33
L36:
	;
	v146 = int32(1)
	v148 = v89 + v146
	if v148 != v77 {
		v89 = v148
		goto L21
	} else {
		goto L37
	}
L37:
	;
	goto L22
}
func F_cube_enlarge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 float64
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v129 float64
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v159 int32
	_ = v159
	var v160 float64
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v306 int32
	_ = v306
	var v309 float64
	_ = v309
	var v314 float64
	_ = v314
	var v317 int32
	_ = v317
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v27 = v25 & int32(2147483647)
	v28 = int32(100)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v28 <= v29 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = v28
	goto L5
L4:
	;
	v32 = v29
	goto L5
L5:
	;
	v33 = int32(0)
	if v33 < v32 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v36 = v32
	goto L8
L7:
	;
	v36 = v33
	goto L8
L8:
	;
	if base.Ui32(v36) < base.Ui32(v27) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = v27
	goto L11
L10:
	;
	v38 = v36
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v39)))
	if base.F64_gt(v40, float64(0)) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v38
	goto L14
L13:
	;
	v43 = v27
	goto L14
L14:
	;
	v47 = v43<<(uint(int32(4))%32) | int32(8)
	v48 = F_palloc0(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v47 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v56 = v53&int32(-2147483648) | v43
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v60 = v58 & int32(2147483647)
	if v60 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v61 = int32(8)
	v62 = v48 + v61
	v64 = v21 + v61
	v65 = int32(0)
	v66 = base.B2i32(v65 <= v58)
	v69 = v65
	v70 = v43
	goto L19
L17:
	;
	v140 = v43
	goto L18
L18:
	;
	if base.Ui32(v43) <= base.Ui32(v60) {
		goto L38
	} else {
		goto L39
	}
L19:
	;
	v88 = v69 << (uint(int32(3)) % 32)
	v89 = v64 + v88
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	if int32(0) <= v58 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v140 = v134
	goto L18
L21:
	;
	v120 = v62 + v70<<(uint(int32(3))%32)
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v117)))
	v122 = base.F64_add(v40, v121)
	*(*float64)(unsafe.Add(mBase, uint32(v120))) = v122
	v124 = v62 + v88
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	if base.F64_lt(v122, v125) != 0 {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v64+(v69+v60)<<(uint(int32(3))%32))))
	v98 = v97
	goto L24
L23:
	;
	v98 = v90
	goto L24
L24:
	;
	if base.F64_le(v98, v90) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if v65 <= v58 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v62+v88))) = base.F64_sub(v90, v40)
	if v65 <= v58 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v105 = v64 + (v69+v58)<<(uint(int32(3))%32)
	goto L30
L29:
	;
	v105 = v89
	goto L30
L30:
	;
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v105)))
	*(*float64)(unsafe.Add(mBase, uint32(v62+v88))) = base.F64_sub(v106, v40)
	v117 = v89
	goto L21
L31:
	;
	v116 = v64 + (v69+v58)<<(uint(int32(3))%32)
	goto L33
L32:
	;
	v116 = v89
	goto L33
L33:
	;
	v117 = v116
	goto L21
L34:
	;
	v129 = base.F64_mul(base.F64_add(v125, v122), float64(0.5))
	*(*float64)(unsafe.Add(mBase, uint32(v124))) = v129
	*(*float64)(unsafe.Add(mBase, uint32(v120))) = v129
	goto L36
L35:
	;
	goto L36
L36:
	;
	v133 = int32(1)
	v134 = v70 + v133
	v136 = v69 + v133
	if v136 != v60 {
		v69 = v136
		v70 = v134
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	v280 = int32(0)
	if v53 < v280 {
		goto L52
	} else {
		goto L53
	}
L39:
	;
	v159 = v48 + int32(8)
	v160 = base.F64_neg(v40)
	v163 = (v43 - v58) & int32(3)
	if v163 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v60-v43) {
		goto L38
	} else {
		goto L47
	}
L41:
	;
	v202 = v60
	v203 = v140
	goto L40
L42:
	;
	goto L43
L43:
	;
	v168 = v60
	v169 = v140
	v170 = int32(0)
	goto L44
L44:
	;
	v186 = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v159+v168<<(uint(v186)%32)))) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v159+v169<<(uint(v186)%32)))) = v40
	v194 = int32(1)
	v195 = v169 + v194
	v197 = v168 + v194
	v199 = v170 + v194
	if v199 != v163 {
		v168 = v197
		v169 = v195
		v170 = v199
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v202 = v197
	v203 = v195
	goto L40
L46:
	;
	goto L45
L47:
	;
	v224 = v202
	v225 = v203
	goto L48
L48:
	;
	v242 = int32(3)
	v244 = v159 + v224<<(uint(v242)%32)
	*(*float64)(unsafe.Add(mBase, uint32(v244))) = v160
	v248 = v159 + v225<<(uint(v242)%32)
	*(*float64)(unsafe.Add(mBase, uint32(v248))) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v244)+8)) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v248)+8)) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v244)+16)) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v248)+16)) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v244)+24)) = v160
	*(*float64)(unsafe.Add(mBase, uint32(v248)+24)) = v40
	v256 = int32(4)
	v259 = v224 + v256
	if v259 != v43 {
		v224 = v259
		v225 = v225 + v256
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L38
L50:
	;
	goto L49
L51:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v365 != v21 {
		goto L59
	} else {
		goto L60
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v43 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v43<<(uint(int32(5))%32) + int32(32)
	goto L51
L53:
	;
	if v56 == int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v286 = v48 + int32(8)
	v289 = v280
	goto L55
L55:
	;
	v306 = int32(3)
	v309 = *(*float64)(unsafe.Add(mBase, uint32(v286+v289<<(uint(v306)%32))))
	v314 = *(*float64)(unsafe.Add(mBase, uint32(v286+(v289+v56)<<(uint(v306)%32))))
	if base.F64_ne(v309, v314) != 0 {
		goto L51
	} else {
		goto L57
	}
L56:
	;
	goto L52
L57:
	;
	v317 = v289 + int32(1)
	if v317 != v56 {
		v289 = v317
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_pfree(m, v21)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	return v48
L62:
	;
	goto L61
}
func F_cube_lt(m *base.Module, l0 int32) int32 {
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 float64
	_ = v286
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 float64
	_ = v350
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 float64
	_ = v406
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v430 int32
	_ = v430
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
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
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 != v6 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	v487 = v465
	goto L4
L6:
	;
	v465 = int32(1)
	goto L5
L7:
	;
	v36 = v31
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(8)
	v38 = v11 + v37
	v40 = v6 + v37
	v48 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L47
	} else {
		goto L48
	}
L13:
	;
	v60 = v48 << (uint(int32(3)) % 32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v40+v60)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v72 = v62
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = int32(8)
	v108 = v11 + v107
	v110 = v6 + v107
	v119 = int32(0)
	goto L30
L15:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38+v60)))
	v76 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v83 = v74
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v69) != 0 {
		v72 = v62
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if base.F64_gt(v72, v83) != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v81) != 0 {
		v83 = v74
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v83 = v81
	goto L18
L21:
	;
	if v29 < int32(0) {
		v92 = v62
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 < int32(0) {
		v100 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v90) != 0 {
		v92 = v62
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v90
	goto L22
L25:
	;
	v102 = int32(-1)
	if base.F64_lt(v92, v100) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v98) != 0 {
		v100 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	goto L25
L28:
	;
	v105 = v48 + int32(1)
	if v105 != v36 {
		v48 = v105
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v131 = v119 << (uint(int32(3)) % 32)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v110+v131)))
	v135 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v143 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v108+v131)))
	v147 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v154 = v145
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v140) != 0 {
		v143 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v143 = v140
	goto L32
L35:
	;
	if base.F64_gt(v143, v154) != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v152) != 0 {
		v154 = v145
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = v152
	goto L35
L38:
	;
	if v29 < int32(0) {
		v163 = v133
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v32 < int32(0) {
		v171 = v145
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v161) != 0 {
		v163 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L39
L42:
	;
	if base.F64_lt(v163, v171) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v169) != 0 {
		v171 = v145
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v171 = v169
	goto L42
L45:
	;
	v175 = v119 + int32(1)
	if v175 != v36 {
		v119 = v175
		goto L30
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	v197 = v6 + int32(8)
	v207 = v36
	goto L50
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v220 = v197 + v207<<(uint(int32(3))%32)
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v266 = v36
	goto L68
L52:
	;
	if base.F64_lt(v243, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v225 = int32(3)
	v227 = v197 + (v207+v29)<<(uint(v225)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v207+v31)<<(uint(v225)%32))))
	if base.F64_lt(v221, v232) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_gt(v221, float64(0)) != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v234 = v220
	goto L58
L57:
	;
	v234 = v227
	goto L58
L58:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v234)))
	if base.F64_gt(v235, float64(0)) != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.F64_lt(v221, v238) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v240 = v221
	goto L62
L61:
	;
	v240 = v238
	goto L62
L62:
	;
	v243 = v240
	goto L52
L63:
	;
	v243 = v221
	goto L52
L64:
	;
	v487 = int32(-1)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v250 = v207 + int32(1)
	if v250 != v31 {
		v207 = v250
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	v274 = v197 + v266<<(uint(int32(3))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v487 = int32(-1)
	goto L4
L70:
	;
	if base.F64_lt(v297, float64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v279 = int32(3)
	v281 = v197 + (v29+v266)<<(uint(v279)%32)
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v266+v31)<<(uint(v279)%32))))
	if base.F64_gt(v275, v286) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if base.F64_gt(v275, float64(0)) != 0 {
		goto L6
	} else {
		goto L81
	}
L74:
	;
	v288 = v274
	goto L76
L75:
	;
	v288 = v281
	goto L76
L76:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.F64_gt(v289, float64(0)) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_gt(v275, v292) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v275
	goto L80
L79:
	;
	v294 = v292
	goto L80
L80:
	;
	v297 = v294
	goto L70
L81:
	;
	v297 = v275
	goto L70
L82:
	;
	v304 = int32(1)
	v306 = v266 + v304
	if v306 == v31 {
		v465 = v304
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L69
L85:
	;
	v266 = v306
	goto L68
L86:
	;
	v487 = int32(0)
	goto L4
L87:
	;
	goto L88
L88:
	;
	v312 = v11 + int32(8)
	v328 = v31
	goto L89
L89:
	;
	v335 = v312 + v328<<(uint(int32(3))%32)
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v386 = v36
	goto L106
L91:
	;
	if base.F64_lt(v364, float64(0)) != 0 {
		goto L6
	} else {
		goto L104
	}
L92:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_lt(v336, v361) != 0 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v340 = int32(3)
	v342 = v312 + (v32+v328)<<(uint(v340)%32)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v328+v34)<<(uint(v340)%32))))
	if base.F64_lt(v336, v347) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.F64_gt(v336, float64(0)) == int32(0) {
		v364 = v336
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v349 = v335
	goto L98
L97:
	;
	v349 = v342
	goto L98
L98:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v349)))
	if base.F64_gt(v350, float64(0)) == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v487 = int32(-1)
	goto L4
L100:
	;
	v487 = int32(-1)
	goto L4
L101:
	;
	v363 = v336
	goto L103
L102:
	;
	v363 = v361
	goto L103
L103:
	;
	v364 = v363
	goto L91
L104:
	;
	v370 = v328 + int32(1)
	if v370 != v34 {
		v328 = v370
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v394 = v312 + v386<<(uint(int32(3))%32)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v465 = int32(-1)
	goto L5
L108:
	;
	if base.F64_lt(v423, float64(0)) != 0 {
		goto L6
	} else {
		goto L121
	}
L109:
	;
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v401)))
	if base.F64_gt(v395, v420) != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v399 = int32(3)
	v401 = v312 + (v32+v386)<<(uint(v399)%32)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v386+v34)<<(uint(v399)%32))))
	if base.F64_gt(v395, v406) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.F64_gt(v395, float64(0)) == int32(0) {
		v423 = v395
		goto L108
	} else {
		goto L117
	}
L113:
	;
	v408 = v394
	goto L115
L114:
	;
	v408 = v401
	goto L115
L115:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)))
	if base.F64_gt(v409, float64(0)) == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v487 = int32(-1)
	goto L4
L117:
	;
	v487 = int32(-1)
	goto L4
L118:
	;
	v422 = v395
	goto L120
L119:
	;
	v422 = v420
	goto L120
L120:
	;
	v423 = v422
	goto L108
L121:
	;
	v430 = v386 + int32(1)
	if v34 != v430 {
		v386 = v430
		goto L106
	} else {
		goto L122
	}
L122:
	;
	goto L107
L123:
	;
	F_pfree(m, v6)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v492 != v11 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_pfree(m, v11)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	return int32(base.Ui32(v487) >> (uint(int32(31)) % 32))
L130:
	;
	goto L129
}
func F_cube_subset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 float64
	_ = v108
	var v113 int32
	_ = v113
	var v120 float64
	_ = v120
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v24 = F_array_contains_nulls(m, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v24 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v32 = F_ArrayGetNItems(m, v29, v22+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L37
	}
L10:
	;
	if int32(101) <= v32 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v38 = int32(8)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v44 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = v32<<(uint(int32(3))%32) + v38
	goto L14
L13:
	;
	v47 = v32<<(uint(int32(4))%32) | v38
	goto L14
L14:
	;
	v48 = F_palloc0(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v47 << (uint(int32(2)) % 32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v53&int32(-2147483648) | v32
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v58 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v32 | int32(-2147483648)
	goto L18
L17:
	;
	goto L18
L18:
	;
	if int32(0) < v32 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v28 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v136 != v17 {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	v72 = v28
	goto L24
L23:
	;
	v72 = (v29<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L24
L24:
	;
	v74 = int32(8)
	v75 = v48 + v74
	v77 = v17 + v74
	v80 = int32(0)
	goto L25
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v22+v72+v80<<(uint(int32(2))%32))))
	if v93 <= int32(0) {
		goto L3
	} else {
		goto L27
	}
L26:
	;
	goto L21
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if base.Ui32(v96&int32(2147483647)) < base.Ui32(v93) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v100 = int32(3)
	v104 = v93 - int32(1)
	v108 = *(*float64)(unsafe.Add(mBase, uint32(v77+v104<<(uint(v100)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v75+v80<<(uint(v100)%32)))) = v108
	if int32(0) <= v96 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v113 = int32(3)
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v77+(v104+v96)<<(uint(v113)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v75+(v80+v32)<<(uint(v113)%32)))) = v120
	goto L31
L30:
	;
	goto L31
L31:
	;
	v123 = v80 + int32(1)
	if v123 != v32 {
		v80 = v123
		goto L25
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	F_pfree(m, v17)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	m.G0 = v14 + int32(16)
	return v48
L36:
	;
	goto L35
L37:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(174084), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(496403), int32(260), int32(105495))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(325759), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(100)
	F_errdetail(m, int32(573688), v14)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(496403), int32(270), int32(105495))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(171522), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(496403), int32(285), int32(105495))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_yy_flush_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	if l0 == int32(0) {
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v6
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v6)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)) = uint8(v6)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v20 == v6 {
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v20+v23<<(uint(int32(2))%32))))
			if l0 != v27 {
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v36)
			}
		}
	}
	return
}
func F_cube_yylex_destroy(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v46 int64
	_ = v46
	var v59 int32
	_ = v59
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v9 = v5 + v6<<(uint(int32(2))%32)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 != 0 {
			v11 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if v13 == v11 {
				F_pfree(m, v10)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32)))) = int32(0)
					v33 = v26
					F_pfree(m, v33)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
						if v41 != 0 {
							F_pfree(m, v41)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
								v46 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
								F_pfree(m, l0)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						} else {
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
							v46 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
							*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
							*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
							*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
							*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
							F_pfree(m, l0)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						}
					}
				}
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 == int32(0) {
					F_pfree(m, v10)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32)))) = int32(0)
						v33 = v26
						F_pfree(m, v33)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							if v41 != 0 {
								F_pfree(m, v41)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
									v46 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
									F_pfree(m, l0)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
								v46 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
								*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
								F_pfree(m, l0)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						}
					}
				} else {
					F_pfree(m, v16)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_pfree(m, v10)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v26+v27<<(uint(int32(2))%32)))) = int32(0)
							v33 = v26
							F_pfree(m, v33)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
								v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
								if v41 != 0 {
									F_pfree(m, v41)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v44 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
										v46 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
										*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
										*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
										*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
										*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
										F_pfree(m, l0)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return int32(0)
										} else {
											return int32(0)
										}
									}
								} else {
									v44 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
									v46 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
									*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
									*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
									F_pfree(m, l0)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										return int32(0)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v33 = v5
			F_pfree(m, v33)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				if v41 != 0 {
					F_pfree(m, v41)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
						v46 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
						*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
						*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
						*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
						F_pfree(m, l0)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				} else {
					v44 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
					v46 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
					*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
					*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
					F_pfree(m, l0)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v41 != 0 {
			F_pfree(m, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
				v46 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
				*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
				F_pfree(m, l0)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			v44 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
			v46 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
			F_pfree(m, l0)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	}
}
func F_cube_yypop_buffer_state(m *base.Module, l0 int32) {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		return
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v10 = v4 + v7<<(uint(int32(2))%32)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		if v11 == int32(0) {
			return
		} else {
			v14 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			if v16 == v14 {
				F_pfree(m, v11)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v32
					if v28 == v32 {
					} else {
						v37 = v28 - int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v27+v37<<(uint(int32(2))%32))))
						if v42 == int32(0) {
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v45
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v47
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v47
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v52)
						}
					}
					return
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				if v19 == int32(0) {
					F_pfree(m, v11)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v32 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v32
						if v28 == v32 {
						} else {
							v37 = v28 - int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
							v42 = *(*int32)(unsafe.Add(mBase, uint32(v27+v37<<(uint(int32(2))%32))))
							if v42 == int32(0) {
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v45
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v47
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v47
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v52)
							}
						}
						return
					}
				} else {
					F_pfree(m, v19)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_pfree(m, v11)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v32 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v27+v28<<(uint(int32(2))%32)))) = v32
							if v28 == v32 {
							} else {
								v37 = v28 - int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v37
								v42 = *(*int32)(unsafe.Add(mBase, uint32(v27+v37<<(uint(int32(2))%32))))
								if v42 == int32(0) {
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v45
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v47
									*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v47
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v50
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)) = uint8(v52)
								}
							}
							return
						}
					}
				}
			}
		}
	}
}
func F_cube_yyrealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	if l0 != 0 {
		v4 = F_repalloc(m, l0, l1)
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			return v4
		}
	} else {
		v9 = F_palloc(m, l1)
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
