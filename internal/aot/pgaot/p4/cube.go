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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v56 float64
	_ = v56
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v107 int32
	_ = v107
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v118 int32
	_ = v118
	var v122 float64
	_ = v122
	var v125 float64
	_ = v125
	var v130 float64
	_ = v130
	var v132 float64
	_ = v132
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v156 int32
	_ = v156
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 == v3) != 0 {
		v156 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v156
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = int32(2147483647)
	v22 = v20 & v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = v23 & v21
	if base.Ui32(v22) < base.Ui32(v25) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v22
	goto L6
L4:
	;
	goto L5
L5:
	;
	if base.Ui32(v22) < base.Ui32(v25) {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v47 = l1 + int32(8) + v33<<(uint(int32(3))%32)
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v47)))
	if base.F64_ne(v48, float64(0)) != 0 {
		v156 = v3
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	if base.B2i32(v23 < int32(0)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = *(*float64)(unsafe.Add(mBase, uint32(v47+v25<<(uint(int32(3))%32))))
	if base.F64_ne(v56, float64(0)) != 0 {
		v156 = v3
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v60 = v33 + int32(1)
	if v60 != v25 {
		v33 = v60
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L7
L14:
	;
	v77 = v22
	goto L16
L15:
	;
	v77 = v25
	goto L16
L16:
	;
	if v77 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(1)
L18:
	;
	goto L19
L19:
	;
	v82 = int32(8)
	v89 = int32(0)
	goto L20
L20:
	;
	v101 = int32(0)
	v103 = v89 << (uint(int32(3)) % 32)
	v104 = l0 + v82 + v103
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v104)))
	v107 = base.B2i32(v20 < v101)
	if v20 < v101 {
		v114 = v105
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v156 = v142
	goto L1
L22:
	;
	v115 = v103 + (l1 + v82)
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v115)))
	v118 = base.B2i32(v23 < int32(0))
	if v23 < int32(0) {
		v125 = v116
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v111 = *(*float64)(unsafe.Add(mBase, uint32(v104+v20<<(uint(int32(3))%32))))
	if base.F64_lt(v105, v111) != 0 {
		v114 = v105
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v114 = v111
	goto L22
L25:
	;
	if base.F64_gt(v114, v125) != 0 {
		v156 = v101
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v122 = *(*float64)(unsafe.Add(mBase, uint32(v115+v23<<(uint(int32(3))%32))))
	if base.F64_lt(v116, v122) != 0 {
		v125 = v116
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v125 = v122
	goto L25
L28:
	;
	if v20 < v101 {
		v132 = v105
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v23 < int32(0) {
		v139 = v116
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v104+v20<<(uint(int32(3))%32))))
	if base.F64_gt(v105, v130) != 0 {
		v132 = v105
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v132 = v130
	goto L29
L32:
	;
	if base.F64_gt(v139, v132) != 0 {
		v156 = v101
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v137 = *(*float64)(unsafe.Add(mBase, uint32(v115+v23<<(uint(int32(3))%32))))
	if base.F64_gt(v116, v137) != 0 {
		v139 = v116
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v139 = v137
	goto L32
L35:
	;
	v142 = int32(1)
	v144 = v89 + v142
	if v144 != v77 {
		v89 = v144
		goto L20
	} else {
		goto L36
	}
L36:
	;
	goto L21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v104 float64
	_ = v104
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v118 float64
	_ = v118
	var v120 int32
	_ = v120
	var v121 float64
	_ = v121
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v156 float64
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v276 int32
	_ = v276
	var v289 int32
	_ = v289
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 float64
	_ = v308
	var v312 float64
	_ = v312
	var v315 int32
	_ = v315
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v30 = int32(0)
	if v30 < v29 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = v29
	goto L5
L4:
	;
	v33 = v30
	goto L5
L5:
	;
	if int32(100) <= v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v36 = int32(100)
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
	v61 = int32(0)
	if v61 <= v58 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v137 = v43
	goto L18
L18:
	;
	if base.Ui32(v43) <= base.Ui32(v60) {
		goto L35
	} else {
		goto L36
	}
L19:
	;
	v64 = v58
	goto L21
L20:
	;
	v64 = v61
	goto L21
L21:
	;
	v65 = int32(8)
	v66 = v48 + v65
	v72 = v43
	v74 = int32(0)
	goto L22
L22:
	;
	v89 = v74 << (uint(int32(3)) % 32)
	v90 = v21 + v65 + v89
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v90)))
	if int32(0) <= v58 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v137 = v130
	goto L18
L24:
	;
	v116 = v66 + v72<<(uint(int32(3))%32)
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v118 = base.F64_add(v40, v117)
	*(*float64)(unsafe.Add(mBase, uint32(v116))) = v118
	v120 = v66 + v89
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v120)))
	if base.F64_lt(v118, v121) != 0 {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v97 = *(*float64)(unsafe.Add(mBase, uint32(v90+v60<<(uint(int32(3))%32))))
	v98 = v97
	goto L27
L26:
	;
	v98 = v91
	goto L27
L27:
	;
	if base.F64_le(v98, v91) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v90+v64<<(uint(int32(3))%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v66+v89))) = base.F64_sub(v104, v40)
	v113 = v90
	goto L24
L29:
	;
	goto L30
L30:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v66+v89))) = base.F64_sub(v91, v40)
	v113 = v90 + v64<<(uint(int32(3))%32)
	goto L24
L31:
	;
	v125 = base.F64_mul(base.F64_add(v121, v118), float64(0.5))
	*(*float64)(unsafe.Add(mBase, uint32(v120))) = v125
	*(*float64)(unsafe.Add(mBase, uint32(v116))) = v125
	goto L33
L32:
	;
	goto L33
L33:
	;
	v129 = int32(1)
	v130 = v72 + v129
	v132 = v74 + v129
	if v132 != v60 {
		v72 = v130
		v74 = v132
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L23
L35:
	;
	v276 = int32(0)
	if base.B2i32(v56 == v276)|base.B2i32(v53 < v276) == v276 {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v155 = v48 + int32(8)
	v156 = base.F64_neg(v40)
	v159 = (v43 - v58) & int32(3)
	if v159 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(int32(-4)) < base.Ui32(v60-v43) {
		goto L35
	} else {
		goto L44
	}
L38:
	;
	v198 = v60
	v200 = v137
	goto L37
L39:
	;
	goto L40
L40:
	;
	v164 = v60
	v166 = v137
	v168 = int32(0)
	goto L41
L41:
	;
	v182 = int32(3)
	*(*float64)(unsafe.Add(mBase, uint32(v155+v164<<(uint(v182)%32)))) = v156
	*(*float64)(unsafe.Add(mBase, uint32(v155+v166<<(uint(v182)%32)))) = v40
	v190 = int32(1)
	v191 = v166 + v190
	v193 = v164 + v190
	v195 = v168 + v190
	if v195 != v159 {
		v164 = v193
		v166 = v191
		v168 = v195
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v198 = v193
	v200 = v191
	goto L37
L43:
	;
	goto L42
L44:
	;
	v220 = v198
	v222 = v200
	goto L45
L45:
	;
	v238 = int32(3)
	v240 = v155 + v220<<(uint(v238)%32)
	*(*float64)(unsafe.Add(mBase, uint32(v240))) = v156
	v244 = v155 + v222<<(uint(v238)%32)
	*(*float64)(unsafe.Add(mBase, uint32(v244))) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v240)+8)) = v156
	*(*float64)(unsafe.Add(mBase, uint32(v244)+8)) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v240)+16)) = v156
	*(*float64)(unsafe.Add(mBase, uint32(v244)+16)) = v40
	*(*float64)(unsafe.Add(mBase, uint32(v240)+24)) = v156
	*(*float64)(unsafe.Add(mBase, uint32(v244)+24)) = v40
	v252 = int32(4)
	v255 = v220 + v252
	if v255 != v43 {
		v220 = v255
		v222 = v222 + v252
		goto L45
	} else {
		goto L47
	}
L46:
	;
	goto L35
L47:
	;
	goto L46
L48:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v363 != v21 {
		goto L56
	} else {
		goto L57
	}
L49:
	;
	v289 = v276
	goto L52
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v43 | int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v43<<(uint(int32(5))%32) + int32(32)
	goto L48
L52:
	;
	v305 = int32(3)
	v307 = v48 + int32(8) + v289<<(uint(v305)%32)
	v308 = *(*float64)(unsafe.Add(mBase, uint32(v307)))
	v312 = *(*float64)(unsafe.Add(mBase, uint32(v307+v56<<(uint(v305)%32))))
	if base.F64_ne(v308, v312) != 0 {
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L51
L54:
	;
	v315 = v289 + int32(1)
	if v315 != v56 {
		v289 = v315
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_pfree(m, v21)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	return v48
L59:
	;
	goto L58
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
	return int32(base.Ui32(v496) >> (uint(int32(31)) % 32))
L121:
	;
	goto L120
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
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v108 int32
	_ = v108
	var v114 float64
	_ = v114
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
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
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v32 = F_ArrayGetNItemsSafe(m, v29, v22+int32(16))
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
	v141 = m.ExcPending
	if v141 != 0 {
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
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v130 != v17 {
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
	v78 = int32(0)
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v22+v72+v78<<(uint(int32(2))%32))))
	if v91 <= int32(0) {
		goto L3
	} else {
		goto L27
	}
L26:
	;
	goto L21
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if base.Ui32(v94&int32(2147483647)) < base.Ui32(v91) {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v98 = int32(3)
	v100 = v48 + int32(8) + v78<<(uint(v98)%32)
	v103 = v17 + v91<<(uint(v98)%32)
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v103)))
	*(*float64)(unsafe.Add(mBase, uint32(v100))) = v104
	if int32(0) <= v94 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = int32(3)
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v103+v94<<(uint(v108)%32))))
	*(*float64)(unsafe.Add(mBase, uint32(v100+v32<<(uint(v108)%32)))) = v114
	goto L31
L30:
	;
	goto L31
L31:
	;
	v117 = v78 + int32(1)
	if v117 != v32 {
		v78 = v117
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
	v133 = m.ExcPending
	if v133 != 0 {
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
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_cube_subset_0), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_cube_subset_1), int32(260), int32(_a_F_cube_subset_2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_cube_subset_3), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(100)
	F_errdetail(m, int32(_a_F_cube_subset_4), v14)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_cube_subset_1), int32(270), int32(_a_F_cube_subset_2))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_cube_subset_5), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_cube_subset_1), int32(285), int32(_a_F_cube_subset_2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
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
								*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
								*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
							*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
							*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
									*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
									*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
								*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
								*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
										*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
										*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
									*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
									*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
						*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
						*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
					*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
					*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
			*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v46
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v44
			*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = v46
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
