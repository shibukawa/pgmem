package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var __phi36 int32
	_ = __phi36
	var v37 int32
	_ = v37
	var __phi37 int32
	_ = __phi37
	var v39 int32
	_ = v39
	var __phi39 int32
	_ = __phi39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
	if l1 <= int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v162
L2:
	;
	if l1 != int32(1) {
		v162 = l1
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_pg_qsort(m, l0, l1, int32(20), int32(_a_F_hstoreUniquePairs_0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = int32(0)
	goto L8
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v19 = v18
	goto L8
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19 + v20
	return int32(1)
L9:
	;
	return int32(0)
L10:
	;
	__phi36 = l0
	__phi37 = l0 + int32(20)
	__phi39 = l0
	v36 = __phi36
	v37 = __phi37
	v39 = __phi39
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v42 != v43 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+16)))
	if v149 != 0 {
		goto L46
	} else {
		goto L47
	}
L13:
	;
	v141 = int32(20)
	v142 = v37 + v141
	v145 = base.I32_div_s(v142-l0, v141)
	if v145 < l1 {
		__phi36 = v137
		__phi37 = v142
		__phi39 = v37
		v36 = __phi36
		v37 = __phi37
		v39 = __phi39
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+16)))
	if v121 != 0 {
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if base.Ui32(int32(4)) <= base.Ui32(v42) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v108 != 0 {
		goto L14
	} else {
		goto L34
	}
L17:
	;
	v108 = int32(0)
	goto L16
L18:
	;
	v82 = v77
	v83 = v78
	v84 = v79
	goto L28
L19:
	;
	if (v45|v46)&int32(3) != 0 {
		v77 = v45
		v78 = v46
		v79 = v42
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v70 = v45
	v71 = v46
	v72 = v42
	goto L21
L21:
	;
	if v72 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v54 = v45
	v55 = v46
	v56 = v42
	goto L23
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v59 != v60 {
		v77 = v54
		v78 = v55
		v79 = v56
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v70 = v65
	v71 = v63
	v72 = v67
	goto L21
L25:
	;
	v62 = int32(4)
	v63 = v55 + v62
	v65 = v54 + v62
	v67 = v56 - v62
	if base.Ui32(int32(3)) < base.Ui32(v67) {
		v54 = v65
		v55 = v63
		v56 = v67
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v77 = v70
	v78 = v71
	v79 = v72
	goto L18
L28:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == v88 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v108 = v87 - v88
	goto L16
L30:
	;
	v90 = int32(1)
	v95 = v84 - v90
	if v95 != 0 {
		v82 = v82 + v90
		v83 = v83 + v90
		v84 = v95
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L17
L34:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+37)))
	if v109 != int32(1) {
		v137 = v36
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v45)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v114 == int32(0) {
		v137 = v36
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v137 = v36
	goto L13
L39:
	;
	v124 = int32(0)
	goto L41
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v124 = v123
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v120 + (v124 + v43)
	v129 = v36 + int32(20)
	if v36 != v39 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+16)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v37)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = v133
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
	*(*int64)(unsafe.Add(mBase, uint32(v129))) = v135
	goto L44
L43:
	;
	goto L44
L44:
	;
	v137 = v129
	goto L13
L45:
	;
	goto L12
L46:
	;
	v152 = int32(0)
	goto L48
L47:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v137)+12))
	v152 = v151
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148 + (v152 + v147)
	v157 = int32(20)
	v160 = base.I32_div_s(v137-l0+v157, v157)
	v162 = v160
	goto L1
}
func F_hstoreUpgrade(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
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
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v493 int32
	_ = v493
	v15 = base.I32_wrap_i64(l0)
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if int32(0) <= v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if v15 == v16 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v493 = v16
	goto L5
L5:
	;
	return v493
L6:
	;
	v24 = F_pg_detoast_datum_copy(m, v15)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v27 = v20
	v28 = v16
	goto L8
L8:
	;
	v30 = v28 + int32(4)
	if v27 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v27 = v26
	v28 = v24
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v476
	v493 = v28
	goto L5
L11:
	;
	v61 = int32(0)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v74 = v72 & int32(268435455)
	if base.B2i32(v74 == v61)|base.B2i32(v72 < v61) != 0 {
		v205 = int32(2)
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v476 = (v57 + v55) << (uint(int32(2)) % 32)
	goto L10
L13:
	;
	v50 = v42 << (uint(int32(3)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50+v30)))
	v55 = v50 + int32(8)
	v57 = v54
	goto L12
L14:
	;
	v55 = int32(8)
	v57 = int32(0)
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(-2147483648)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if base.Ui32(int32(_a_F_hstoreUpgrade_0)) < base.Ui32(v35) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if int32(0) <= v38 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v42 = v27 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v42 | int32(-2147483648)
	if v42 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v27) {
		goto L55
	} else {
		goto L56
	}
L22:
	;
	v224 = v205
	goto L21
L23:
	;
	v81 = v28 + int32(8)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	if int32(0) <= v82 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v224 = int32(0)
	goto L21
L25:
	;
	goto L26
L26:
	;
	v86 = int32(0)
	v88 = v74 << (uint(int32(3)) % 32)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88+v81-int32(4))))
	v97 = v88 + v92&int32(1073741823) + int32(8)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v100 = int32(base.Ui32(v98) >> (uint(int32(2)) % 32))
	if base.Ui32(v100) < base.Ui32(v97) {
		v205 = v86
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v102 = int32(1)
	v106 = v102
	goto L28
L28:
	;
	v118 = v81 + v106<<(uint(int32(2))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v119 < int32(0) {
		v205 = v86
		goto L22
	} else {
		goto L30
	}
L29:
	;
	if v74 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v122 = int32(1073741823)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118-int32(4))))
	if base.Ui32(v119&v122) < base.Ui32(v126&v122) {
		v205 = v86
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v131 = v106 + int32(1)
	if v131 != v74<<(uint(v102)%32) {
		v106 = v131
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v135 = int32(2)
	if base.Ui32(v74) <= base.Ui32(v135) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v97 == v100 {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v138 = v135
	goto L38
L37:
	;
	v138 = v74
	goto L38
L38:
	;
	v142 = int32(1)
	goto L39
L39:
	;
	v152 = v142 << (uint(int32(3)) % 32)
	v153 = v81 + v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v156 = v154 & int32(1073741823)
	if int32(0) <= v154 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L35
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v153-int32(4))))
	v165 = v156 - v161&int32(1073741823)
	goto L43
L42:
	;
	v165 = v156
	goto L43
L43:
	;
	v166 = v28 + v152
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v169 = v167 & int32(1073741823)
	if int32(0) <= v167 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166-int32(4))))
	v178 = v169 - v174&int32(1073741823)
	goto L46
L45:
	;
	v178 = v169
	goto L46
L46:
	;
	if v154&int32(1073741824)|base.B2i32(base.Ui32(v165) < base.Ui32(v178)) != 0 {
		v205 = int32(0)
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v185 = v142 + int32(1)
	if v185 != v138 {
		v142 = v185
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
L49:
	;
	v201 = int32(2)
	goto L51
L50:
	;
	v201 = int32(1)
	goto L51
L51:
	;
	v205 = v201
	goto L22
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v459
	v476 = v472
	goto L10
L53:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	v442 = int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v441 | v442
	v448 = v337 << (uint(int32(3)) % 32)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v448+v239-int32(4))))
	v459 = v337 | v442
	v472 = (v448+v452)<<(uint(int32(2))%32) + int32(32)
	goto L52
L54:
	;
	if v337 != 0 {
		goto L53
	} else {
		goto L99
	}
L55:
	;
	if v224 != 0 {
		goto L90
	} else {
		goto L91
	}
L56:
	;
	v230 = v27<<(uint(int32(3))%32) + int32(8)
	v232 = int32(base.Ui32(v35) >> (uint(int32(2)) % 32))
	if base.Ui32(v232) < base.Ui32(v230) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v234 = int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if base.Ui32(v234) < base.Ui32(v235) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v239 = v28 + int32(8)
	if v27 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v243 = v234
	goto L62
L60:
	;
	goto L61
L61:
	;
	v280 = int32(1)
	if v27 <= v280 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v257 = v243 << (uint(int32(3)) % 32)
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239+v257))))
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257+v28))))
	if base.Ui32(v259) < base.Ui32(v261) {
		goto L55
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v264 = v243 + int32(1)
	if v264 != v27 {
		v243 = v264
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v283 = v280
	goto L68
L67:
	;
	v283 = v27
	goto L68
L68:
	;
	v284 = int32(0)
	v287 = v284
	v289 = v284
	goto L69
L69:
	;
	v302 = v239 + v289<<(uint(int32(3))%32)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	if int32(base.Ui32(v303)>>(uint(int32(1))%32)) != v287 {
		goto L55
	} else {
		goto L71
	}
L70:
	;
	if base.Ui32(v232) < base.Ui32(v314+v230) {
		goto L55
	} else {
		goto L76
	}
L71:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302))))
	if v303&int32(1) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v312 = int32(0)
	goto L74
L73:
	;
	v311 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v302)+2)))
	v312 = v311
	goto L74
L74:
	;
	v314 = v312 + (v287 + v307)
	v316 = v289 + int32(1)
	if v316 != v283 {
		v287 = v314
		v289 = v316
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L70
L76:
	;
	if v224 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if v337 <= int32(0) {
		goto L54
	} else {
		goto L83
	}
L78:
	;
	v324 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v324 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	F_errmsg_internal(m, int32(_a_F_hstoreUpgrade_1), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_hstoreUpgrade_2), int32(311), int32(_a_F_hstoreUpgrade_3))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L77
L83:
	;
	v342 = int32(0)
	goto L84
L84:
	;
	v357 = v239 + v342<<(uint(int32(3))%32)
	v358 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357)+2)))
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357))))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v357)+4))
	v361 = int32(1)
	v363 = v359 + int32(base.Ui32(v360)>>(uint(v361)%32))
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v363 & int32(1073741823)
	v369 = v360 & v361
	if v369 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L53
L86:
	;
	v370 = int32(0)
	goto L88
L87:
	;
	v370 = v358
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357)+4)) = (v370+v363)&int32(1073741823) | v369<<(uint(int32(30))%32)
	v379 = v342 + int32(1)
	if v379 != v337 {
		v342 = v379
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v396 = v27 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v396 | int32(-2147483648)
	if v396 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L96
	}
L93:
	;
	v476 = int32(32)
	goto L10
L94:
	;
	goto L95
L95:
	;
	v404 = v396 << (uint(int32(3)) % 32)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v404+v30)))
	v476 = (v404+v406)<<(uint(int32(2))%32) + int32(32)
	goto L10
L96:
	;
	F_errmsg_internal(m, int32(_a_F_hstoreUpgrade_4), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_hstoreUpgrade_2), int32(273), int32(_a_F_hstoreUpgrade_3))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	v459 = int32(-2147483648)
	v472 = int32(32)
	goto L52
}
func F_hstore_avals(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_hstoreUpgrade(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v23 = v21 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = int32(1)
	if v23 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v14 + int32(16)
	return base.I64_extend_i32_u(v114)
L4:
	;
	v30 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v33 = v17 + int32(8)
	v34 = int32(3)
	v38 = v33 + v21<<(uint(v34)%32)&int32(2147483640)
	v42 = F_palloc(m, v23<<(uint(v34)%32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v114 = v30
	goto L3
L8:
	;
	v44 = F_palloc(m, v23)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(0)
	goto L10
L10:
	;
	v58 = v46 << (uint(int32(3)) % 32)
	v59 = v33 + v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v60&int32(1073741824) != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v101 = F_construct_md_array(m, v42, v44, int32(1), v14+int32(12), v14+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42+v58))) = v83
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v44))) = uint8(v84)
	v90 = v46 + int32(1)
	if v90 != v23 {
		v46 = v90
		goto L10
	} else {
		goto L21
	}
L13:
	;
	v83 = int64(0)
	v84 = int32(1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	if v60 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v77 = F_cstring_to_text_with_len(m, v76, v74)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v74 = v60 & int32(1073741823)
	v76 = v38
	goto L16
L18:
	;
	goto L19
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v71 = v69 & int32(1073741823)
	v74 = v60 - v71
	v76 = v71 + v38
	goto L16
L20:
	;
	v83 = base.I64_extend_i32_u(v77)
	v84 = int32(0)
	goto L12
L21:
	;
	goto L11
L22:
	;
	v114 = v101
	goto L3
}
func F_hstore_contains(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v355 int64
	_ = v355
	v20 = int64(0)
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_hstoreUpgrade(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_hstoreUpgrade(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	return v355
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v31 = v29 & int32(268435455)
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = int32(8)
	v33 = v22 + v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v35 = int32(3)
	v41 = v27 + v32
	v44 = v41 + v31<<(uint(v35)%32)
	v46 = v34 & int32(268435455)
	v51 = int32(0)
	v62 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	v355 = int64(1)
	goto L3
L8:
	;
	v73 = v41 + v62<<(uint(int32(3))%32)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	if v46 <= v51 {
		v355 = v20
		goto L3
	} else {
		goto L14
	}
L11:
	;
	v89 = v74 & int32(1073741823)
	v90 = v44
	goto L10
L12:
	;
	goto L13
L13:
	;
	v79 = int32(1073741823)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73-int32(4))))
	v85 = v83 & v79
	v89 = v74&v79 - v85
	v90 = v85 + v44
	goto L10
L14:
	;
	v92 = v51
	v96 = v46
	goto L17
L15:
	;
	v310 = int32(1)
	v313 = v62 + v310
	if base.Ui32(v313) < base.Ui32(v31) {
		v51 = v115 + v310
		v62 = v313
		goto L8
	} else {
		goto L84
	}
L16:
	;
	v242 = v222 + v44
	v243 = v241 + (v33 + v34<<(uint(v35)%32)&int32(2147483640))
	if base.Ui32(int32(4)) <= base.Ui32(v223) {
		goto L68
	} else {
		goto L69
	}
L17:
	;
	v115 = int32(base.Ui32(v96-v92)>>(uint(int32(1))%32)) + v92
	v118 = v33 + v115<<(uint(int32(3))%32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v121 = v119 & int32(1073741823)
	if int32(0) <= v119 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	if v223 != v225 {
		v355 = v20
		goto L3
	} else {
		goto L64
	}
L19:
	;
	goto L18
L20:
	;
	v235 = base.B2i32(v231 < int32(0))
	if v231 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L21:
	;
	v141 = v140 + (v33 + v46<<(uint(v35)%32))
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L34
	} else {
		goto L35
	}
L22:
	;
	if base.Ui32(v89) < base.Ui32(v133) {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v118-int32(4))))
	v128 = v126 & int32(1073741823)
	v129 = v121 - v128
	if v129 != v89 {
		v133 = v129
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v121 == v89 {
		v140 = int32(0)
		goto L21
	} else {
		goto L27
	}
L26:
	;
	v140 = v128
	goto L21
L27:
	;
	v133 = v121
	goto L22
L28:
	;
	v138 = int32(1)
	goto L30
L29:
	;
	v138 = int32(-1)
	goto L30
L30:
	;
	v231 = v138
	goto L20
L31:
	;
	if v203 != 0 {
		v231 = v203
		goto L20
	} else {
		goto L49
	}
L32:
	;
	v203 = int32(0)
	goto L31
L33:
	;
	v177 = v172
	v178 = v173
	v179 = v174
	goto L43
L34:
	;
	if (v141|v90)&int32(3) != 0 {
		v172 = v141
		v173 = v90
		v174 = v89
		goto L33
	} else {
		goto L37
	}
L35:
	;
	v165 = v141
	v166 = v90
	v167 = v89
	goto L36
L36:
	;
	if v167 == int32(0) {
		goto L32
	} else {
		goto L42
	}
L37:
	;
	v149 = v141
	v150 = v90
	v151 = v89
	goto L38
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v154 != v155 {
		v172 = v149
		v173 = v150
		v174 = v151
		goto L33
	} else {
		goto L40
	}
L39:
	;
	v165 = v160
	v166 = v158
	v167 = v162
	goto L36
L40:
	;
	v157 = int32(4)
	v158 = v150 + v157
	v160 = v149 + v157
	v162 = v151 - v157
	if base.Ui32(int32(3)) < base.Ui32(v162) {
		v149 = v160
		v150 = v158
		v151 = v162
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v172 = v165
	v173 = v166
	v174 = v167
	goto L33
L43:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if v182 == v183 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v203 = v182 - v183
	goto L31
L45:
	;
	v185 = int32(1)
	v190 = v179 - v185
	if v190 != 0 {
		v177 = v177 + v185
		v178 = v178 + v185
		v179 = v190
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	goto L32
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	v205 = int32(30)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v211 = v209 & int32(1073741824)
	if int32(base.Ui32(v204)>>(uint(v205)%32))&int32(1) != int32(base.Ui32(v211)>>(uint(v205)%32)) {
		v355 = v20
		goto L3
	} else {
		goto L50
	}
L50:
	;
	if v211 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	v215 = int32(1073741823)
	v219 = int32(0)
	if v219 <= v209 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v222 = v74 & v215
	goto L54
L53:
	;
	v222 = v219
	goto L54
L54:
	;
	v223 = v209&v215 - v222
	v225 = v204 & int32(1073741823)
	if v204 < int32(0) {
		goto L19
	} else {
		goto L55
	}
L55:
	;
	if v223 == v225-v121 {
		v241 = v121
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v355 = v20
	goto L3
L57:
	;
	v236 = v115 + int32(1)
	goto L59
L58:
	;
	v236 = v92
	goto L59
L59:
	;
	if v231 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v237 = v96
	goto L62
L61:
	;
	v237 = v115
	goto L62
L62:
	;
	if v236 < v237 {
		v92 = v236
		v96 = v237
		goto L17
	} else {
		goto L63
	}
L63:
	;
	v355 = v20
	goto L3
L64:
	;
	v241 = int32(0)
	goto L16
L65:
	;
	if v305 != 0 {
		v355 = v20
		goto L3
	} else {
		goto L83
	}
L66:
	;
	v305 = int32(0)
	goto L65
L67:
	;
	v279 = v274
	v280 = v275
	v281 = v276
	goto L77
L68:
	;
	if (v242|v243)&int32(3) != 0 {
		v274 = v242
		v275 = v243
		v276 = v223
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v267 = v242
	v268 = v243
	v269 = v223
	goto L70
L70:
	;
	if v269 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v251 = v242
	v252 = v243
	v253 = v223
	goto L72
L72:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	if v256 != v257 {
		v274 = v251
		v275 = v252
		v276 = v253
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v267 = v262
	v268 = v260
	v269 = v264
	goto L70
L74:
	;
	v259 = int32(4)
	v260 = v252 + v259
	v262 = v251 + v259
	v264 = v253 - v259
	if base.Ui32(int32(3)) < base.Ui32(v264) {
		v251 = v262
		v252 = v260
		v253 = v264
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v274 = v267
	v275 = v268
	v276 = v269
	goto L67
L77:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	if v284 == v285 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v305 = v284 - v285
	goto L65
L79:
	;
	v287 = int32(1)
	v292 = v281 - v287
	if v292 != 0 {
		v279 = v279 + v287
		v280 = v280 + v287
		v281 = v292
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	goto L15
L84:
	;
	goto L9
}
func F_hstore_delete(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	v2 = int32(0)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = F_hstoreUpgrade(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v24 = F_pg_detoast_datum_packed(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v27 = int32(1)
	v28 = v26 & v27
	if v26 == v27 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v59 = F_palloc(m, int32(base.Ui32(v56)>>(uint(int32(2))%32)))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v34 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v45 = int32(1)
	if v28 != 0 {
		v55 = int32(base.Ui32(v26)>>(uint(v45)%32)) - v45
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = int32(16)
	goto L10
L9:
	;
	v37 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v44 = int32(4)
	goto L13
L12:
	;
	v44 = v37
	goto L13
L13:
	;
	v55 = v44
	goto L4
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v62 & int32(-4)
	v67 = v61 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v67 | int32(-2147483648)
	v72 = v59 + int32(8)
	v75 = v72 + v67<<(uint(int32(3))%32)
	if v67 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v242 == v254&int32(268435455) {
		goto L62
	} else {
		goto L63
	}
L17:
	;
	v238 = int32(0)
	v242 = v2
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v28 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v81 = int32(1)
	goto L22
L21:
	;
	v81 = int32(4)
	goto L22
L22:
	;
	v82 = v24 + v81
	v84 = v19 + int32(8)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v91 = v72
	v94 = v75
	v96 = v2
	v101 = v2
	goto L23
L23:
	;
	v110 = v84 + v101<<(uint(int32(3))%32)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v113 = v111 & int32(1073741823)
	if v111 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v229 = v222 - v75
	if v223 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L25:
	;
	v125 = v124 + (v84 + v85<<(uint(int32(3))%32)&int32(2147483640))
	if v123 == v55 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v123 = v113
	v124 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v110-int32(4))))
	v121 = v119 & int32(1073741823)
	v123 = v113 - v121
	v124 = v121
	goto L25
L29:
	;
	v227 = v101 + int32(1)
	if v227 != v67 {
		v91 = v220
		v94 = v222
		v96 = v223
		v101 = v227
		goto L23
	} else {
		goto L58
	}
L30:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v55) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	goto L32
L32:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v192 = int32(1073741823)
	v196 = int32(0)
	if v196 <= v191 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	if v188 == int32(0) {
		v220 = v91
		v222 = v94
		v223 = v96
		goto L29
	} else {
		goto L51
	}
L34:
	;
	v188 = int32(0)
	goto L33
L35:
	;
	v162 = v157
	v163 = v158
	v164 = v159
	goto L45
L36:
	;
	if (v125|v82)&int32(3) != 0 {
		v157 = v125
		v158 = v82
		v159 = v55
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v150 = v125
	v151 = v82
	v152 = v55
	goto L38
L38:
	;
	if v152 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v134 = v125
	v135 = v82
	v136 = v55
	goto L40
L40:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v139 != v140 {
		v157 = v134
		v158 = v135
		v159 = v136
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v150 = v145
	v151 = v143
	v152 = v147
	goto L38
L42:
	;
	v142 = int32(4)
	v143 = v135 + v142
	v145 = v134 + v142
	v147 = v136 - v142
	if base.Ui32(int32(3)) < base.Ui32(v147) {
		v134 = v145
		v135 = v143
		v136 = v147
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v157 = v150
	v158 = v151
	v159 = v152
	goto L35
L45:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v167 == v168 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v188 = v167 - v168
	goto L33
L47:
	;
	v170 = int32(1)
	v175 = v164 - v170
	if v175 != 0 {
		v162 = v162 + v170
		v163 = v163 + v170
		v164 = v175
		goto L45
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L46
L50:
	;
	goto L34
L51:
	;
	goto L32
L52:
	;
	v199 = v111 & v192
	goto L54
L53:
	;
	v199 = v196
	goto L54
L54:
	;
	v200 = v191&v192 - v199
	v201 = v123 + v200
	if v201 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v94, v125, v201)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v203 = v201 + v94
	v204 = v203 - v75
	v206 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = (v204 - v200) & v206
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v204&v206 | v211&int32(1073741824)
	v220 = v91 + int32(8)
	v222 = v203
	v223 = v96 + int32(1)
	goto L29
L58:
	;
	goto L24
L59:
	;
	v238 = v229
	v242 = int32(0)
	goto L16
L60:
	;
	goto L61
L61:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v233 | int32(-2147483648)
	v238 = v229
	v242 = v223
	goto L16
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = (v242<<(uint(int32(3))%32)+v238)<<(uint(int32(2))%32) + int32(32)
	return base.I64_extend_i32_u(v59)
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v242 | int32(-2147483648)
	if v238 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	base.MemoryCopy(m, v72+v242<<(uint(int32(3))%32)&int32(2147483640), v75, v238)
	goto L62
}
func F_hstore_each(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int64
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v18 = F_hstoreUpgrade(m, v17)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			v22 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = int32(_a_F_hstore_each_0)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0]))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0])) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v32 = F_palloc(m, int32(base.Ui32(v29)>>(uint(int32(2))%32)))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v36 = int32(base.Ui32(v34) >> (uint(int32(2)) % 32))
					if v36 != 0 {
						base.MemoryCopy(m, v32, v18, v36)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v32
					v42 = F_get_call_result_type(m, l0, int32(0), v11+int32(16))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						if v42 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_hstore_each_1), int32(0))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_hstore_each_2), int32(869), int32(_a_F_hstore_each_3))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
							v47 = F_BlessTupleDesc(m, v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v47
								*(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0])) = v25
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
								v62 = v60 & int32(268435455)
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
								if base.Ui32(v63) < base.Ui32(v62) {
									v65 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v65)
									v68 = v59 + int32(8)
									v69 = int32(3)
									v71 = v68 + v62<<(uint(v69)%32)
									v74 = v68 + v63<<(uint(v69)%32)
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
									if v75 < v65 {
										v89 = v75 & int32(1073741823)
										v91 = v71
									} else {
										v80 = int32(1073741823)
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v74-int32(4))))
										v86 = v84 & v80
										v89 = v75&v80 - v86
										v91 = v71 + v86
									}
									v92 = F_cstring_to_text_with_len(m, v91, v89)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = base.I64_extend_i32_u(v92)
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
										if v96&int32(1073741824) != 0 {
											v99 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v99)
											*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
											v127 = F_heap_form_tuple(m, v122, v11+int32(16), v11+int32(14))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int64(0)
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
												v130 = F_HeapTupleHeaderGetDatum(m, v129)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int64(0)
												} else {
													v132 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
													*(*int64)(unsafe.Add(mBase, uint32(v58))) = v132 + int64(1)
													v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(1)
													v150 = v130
													m.G0 = v11 + int32(32)
													return v150
												}
											}
										} else {
											if v96 < int32(0) {
												v113 = v71
												v114 = v96 & int32(1073741823)
											} else {
												v107 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
												v109 = v107 & int32(1073741823)
												v113 = v109 + v71
												v114 = v96 - v109
											}
											v115 = F_cstring_to_text_with_len(m, v113, v114)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = base.I64_extend_i32_u(v115)
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
												v127 = F_heap_form_tuple(m, v122, v11+int32(16), v11+int32(14))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int64(0)
												} else {
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
													v130 = F_HeapTupleHeaderGetDatum(m, v129)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int64(0)
													} else {
														v132 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
														*(*int64)(unsafe.Add(mBase, uint32(v58))) = v132 + int64(1)
														v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(1)
														v150 = v130
														m.G0 = v11 + int32(32)
														return v150
													}
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = int32(2)
										v144 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v144)
										v150 = int64(0)
										m.G0 = v11 + int32(32)
										return v150
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
		v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
		v62 = v60 & int32(268435455)
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
		if base.Ui32(v63) < base.Ui32(v62) {
			v65 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+14)) = uint16(v65)
			v68 = v59 + int32(8)
			v69 = int32(3)
			v71 = v68 + v62<<(uint(v69)%32)
			v74 = v68 + v63<<(uint(v69)%32)
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
			if v75 < v65 {
				v89 = v75 & int32(1073741823)
				v91 = v71
			} else {
				v80 = int32(1073741823)
				v84 = *(*int32)(unsafe.Add(mBase, uint32(v74-int32(4))))
				v86 = v84 & v80
				v89 = v75&v80 - v86
				v91 = v71 + v86
			}
			v92 = F_cstring_to_text_with_len(m, v91, v89)
			mBase = m.M
			v93 = m.ExcPending
			if v93 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = base.I64_extend_i32_u(v92)
				v96 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
				if v96&int32(1073741824) != 0 {
					v99 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v99)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
					v127 = F_heap_form_tuple(m, v122, v11+int32(16), v11+int32(14))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return int64(0)
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
						v130 = F_HeapTupleHeaderGetDatum(m, v129)
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return int64(0)
						} else {
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
							*(*int64)(unsafe.Add(mBase, uint32(v58))) = v132 + int64(1)
							v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(1)
							v150 = v130
							m.G0 = v11 + int32(32)
							return v150
						}
					}
				} else {
					if v96 < int32(0) {
						v113 = v71
						v114 = v96 & int32(1073741823)
					} else {
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						v109 = v107 & int32(1073741823)
						v113 = v109 + v71
						v114 = v96 - v109
					}
					v115 = F_cstring_to_text_with_len(m, v113, v114)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int64(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = base.I64_extend_i32_u(v115)
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v58)+28))
						v127 = F_heap_form_tuple(m, v122, v11+int32(16), v11+int32(14))
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int64(0)
						} else {
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
							v130 = F_HeapTupleHeaderGetDatum(m, v129)
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int64(0)
							} else {
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
								*(*int64)(unsafe.Add(mBase, uint32(v58))) = v132 + int64(1)
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(1)
								v150 = v130
								m.G0 = v11 + int32(32)
								return v150
							}
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int64(0)
			} else {
				v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = int32(2)
				v144 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v144)
				v150 = int64(0)
				m.G0 = v11 + int32(32)
				return v150
			}
		}
	}
}
func F_hstore_exists_any(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v219 int64
	_ = v219
	v16 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = F_hstoreUpgrade(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = F_hstoreArrayToPairs(m, v27, v19+int32(12))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v33 <= int32(0) {
		v219 = v16
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v19 + int32(16)
	return v219
L6:
	;
	v37 = v22 + int32(8)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v40 = v38 & int32(268435455)
	v45 = int32(0)
	v52 = int32(0)
	goto L7
L7:
	;
	if v45 < v40 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v219 = v16
	goto L5
L9:
	;
	v64 = v31 + v52*int32(20)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v67 = v45
	v69 = v40
	goto L12
L10:
	;
	v185 = v45
	goto L11
L11:
	;
	v202 = v52 + int32(1)
	if v202 != v33 {
		v45 = v185
		v52 = v202
		goto L7
	} else {
		goto L51
	}
L12:
	;
	v86 = int32(base.Ui32(v69-v67)>>(uint(int32(1))%32)) + v67
	v89 = v37 + v86<<(uint(int32(3))%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v92 = v90 & int32(1073741823)
	if int32(0) <= v90 {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v185 = v182
	goto L11
L14:
	;
	v181 = base.B2i32(v176 < int32(0))
	if v176 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L15:
	;
	v112 = v111 + (v37 + v40<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v65) {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	if base.Ui32(v65) < base.Ui32(v104) {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(4))))
	v99 = v97 & int32(1073741823)
	v100 = v92 - v99
	if v100 != v65 {
		v104 = v100
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v92 == v65 {
		v111 = int32(0)
		goto L15
	} else {
		goto L21
	}
L20:
	;
	v111 = v99
	goto L15
L21:
	;
	v104 = v92
	goto L16
L22:
	;
	v109 = int32(1)
	goto L24
L23:
	;
	v109 = int32(-1)
	goto L24
L24:
	;
	v176 = v109
	goto L14
L25:
	;
	if v174 != 0 {
		v176 = v174
		goto L14
	} else {
		goto L43
	}
L26:
	;
	v174 = int32(0)
	goto L25
L27:
	;
	v148 = v143
	v149 = v144
	v150 = v145
	goto L37
L28:
	;
	if (v112|v66)&int32(3) != 0 {
		v143 = v112
		v144 = v66
		v145 = v65
		goto L27
	} else {
		goto L31
	}
L29:
	;
	v136 = v112
	v137 = v66
	v138 = v65
	goto L30
L30:
	;
	if v138 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L31:
	;
	v120 = v112
	v121 = v66
	v122 = v65
	goto L32
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v125 != v126 {
		v143 = v120
		v144 = v121
		v145 = v122
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v136 = v131
	v137 = v129
	v138 = v133
	goto L30
L34:
	;
	v128 = int32(4)
	v129 = v121 + v128
	v131 = v120 + v128
	v133 = v122 - v128
	if base.Ui32(int32(3)) < base.Ui32(v133) {
		v120 = v131
		v121 = v129
		v122 = v133
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v143 = v136
	v144 = v137
	v145 = v138
	goto L27
L37:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v153 == v154 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v174 = v153 - v154
	goto L25
L39:
	;
	v156 = int32(1)
	v161 = v150 - v156
	if v161 != 0 {
		v148 = v148 + v156
		v149 = v149 + v156
		v150 = v161
		goto L37
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	goto L26
L43:
	;
	v219 = int64(1)
	goto L5
L44:
	;
	v182 = v86 + int32(1)
	goto L46
L45:
	;
	v182 = v67
	goto L46
L46:
	;
	if v176 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v183 = v69
	goto L49
L48:
	;
	v183 = v86
	goto L49
L49:
	;
	if v182 < v183 {
		v67 = v182
		v69 = v183
		goto L12
	} else {
		goto L50
	}
L50:
	;
	goto L13
L51:
	;
	goto L8
}
func F_hstore_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_gt_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int32(0) < base.I32_wrap_i64(v6)))
	}
}
func F_hstore_ne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_ne_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.I32_wrap_i64(v6) != int32(0)))
	}
}
func F_hstore_subscript_assign(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v445 int32
	_ = v445
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	v4 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(32)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v30 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v35 = F_pg_detoast_datum_packed(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L110
	}
L4:
	;
	return
L5:
	;
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+29)) = uint8(v37)
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v41&v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v44 = v39
	goto L8
L7:
	;
	v44 = int32(4)
	goto L8
L8:
	;
	v45 = v35 + v44
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v45
	if v41 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v76 = F_hstoreCheckKeyLen(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+1)))
	if v52 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v63 = int32(1)
	if v41&v63 != 0 {
		v75 = int32(base.Ui32(v41)>>(uint(v63)%32)) - v63
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v55 = int32(16)
	goto L15
L14:
	;
	v55 = int32(0)
	goto L15
L15:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v62 = int32(4)
	goto L18
L17:
	;
	v62 = v55
	goto L18
L18:
	;
	v75 = v62
	goto L9
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v76
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+48)))
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v132 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L22:
	;
	v123 = int32(1)
	v127 = v4
	v128 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v83 = F_pg_detoast_datum_packed(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v85 = int32(1)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = v87 & v85
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v90 = v85
	goto L28
L27:
	;
	v90 = int32(4)
	goto L28
L28:
	;
	v91 = v83 + v90
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v91
	if v87 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v121 = F_hstoreCheckValLen(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L40
	}
L30:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v99 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v110 = int32(1)
	if v89 != 0 {
		v120 = int32(base.Ui32(v87)>>(uint(v110)%32)) - v110
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v102 = int32(16)
	goto L35
L34:
	;
	v102 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v99-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = int32(4)
	goto L38
L37:
	;
	v109 = v102
	goto L38
L38:
	;
	v120 = v109
	goto L29
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v120 = int32(base.Ui32(v114)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	v123 = int32(0)
	v127 = v91
	v128 = v121
	goto L21
L41:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v460))) = base.I64_extend_i32_u(v445)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v464 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v463))) = uint8(v464)
	m.G0 = v26 + int32(32)
	return
L42:
	;
	v139 = F_hstorePairs(m, v26+int32(12), int32(1), v76+v128)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	v143 = F_hstoreUpgrade(m, v142)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	v445 = v139
	goto L41
L46:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v151 = v149 & int32(268435455)
	v154 = int32(8)
	v155 = v151<<(uint(int32(3))%32) + v154
	v159 = v76 + int32(base.Ui32(v145)>>(uint(int32(2))%32)) + v155 + v128 + v154
	v160 = F_palloc(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v151 - int32(2147483647)
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v159 << (uint(int32(2)) % 32)
	v168 = int32(8)
	v169 = v143 + v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v173 = int32(2147483640)
	v175 = v169 + v170<<(uint(int32(3))%32)&v173
	v179 = v160 + v168
	v182 = v179 + v155&v173
	v183 = v179
	v186 = base.B2i32(v151 != int32(0))
	v187 = v182
	v198 = v4
	v200 = v4
	v201 = v4
	goto L48
L48:
	;
	if v186 == int32(0) {
		v309 = int32(1)
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v409 | int32(-2147483648)
	v413 = v393 - v182
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v401 == v414&int32(268435455) {
		goto L107
	} else {
		goto L108
	}
L50:
	;
	v401 = v198 + int32(1)
	v404 = v399 + v200
	v405 = base.B2i32(base.Ui32(v404) < base.Ui32(v151))
	if v405|base.B2i32(v398 <= int32(0)) != 0 {
		v183 = v183 + int32(8)
		v186 = v405
		v187 = v393
		v198 = v401
		v200 = v404
		v201 = v398
		goto L48
	} else {
		goto L106
	}
L51:
	;
	v346 = v302 + int32(4)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v350 = int32(0)
	if v350 <= v347 {
		goto L94
	} else {
		goto L95
	}
L52:
	;
	v332 = int32(1073741823)
	v333 = v303 & v332
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v302-int32(4))))
	v338 = v336 & v332
	v341 = v333
	v342 = v333 - v338
	v344 = v175 + v338
	goto L51
L53:
	;
	if v76 != 0 {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	v210 = v200 << (uint(int32(3)) % 32)
	if int32(0) < v201 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v302 = v210 + v169
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	if int32(0) <= v303 {
		goto L52
	} else {
		goto L84
	}
L56:
	;
	v213 = v210 + v169
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v216 = v214 & int32(1073741823)
	if int32(0) <= v214 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v234 = v175 + v233
	if base.Ui32(int32(4)) <= base.Ui32(v76) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v228 <= v76 {
		goto L55
	} else {
		goto L64
	}
L59:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v213-int32(4))))
	v223 = v221 & int32(1073741823)
	v224 = v216 - v223
	if v224 != v76 {
		v228 = v224
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v216 == v76 {
		v233 = int32(0)
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v233 = v223
	goto L57
L63:
	;
	v228 = v216
	goto L58
L64:
	;
	v309 = int32(1)
	goto L53
L65:
	;
	if int32(0) <= v296 {
		v309 = v296
		goto L53
	} else {
		goto L83
	}
L66:
	;
	v296 = int32(0)
	goto L65
L67:
	;
	v270 = v265
	v271 = v266
	v272 = v267
	goto L77
L68:
	;
	if (v234|v45)&int32(3) != 0 {
		v265 = v234
		v266 = v45
		v267 = v76
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v258 = v234
	v259 = v45
	v260 = v76
	goto L70
L70:
	;
	if v260 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v242 = v234
	v243 = v45
	v244 = v76
	goto L72
L72:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	if v247 != v248 {
		v265 = v242
		v266 = v243
		v267 = v244
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v258 = v253
	v259 = v251
	v260 = v255
	goto L70
L74:
	;
	v250 = int32(4)
	v251 = v243 + v250
	v253 = v242 + v250
	v255 = v244 - v250
	if base.Ui32(int32(3)) < base.Ui32(v255) {
		v242 = v253
		v243 = v251
		v244 = v255
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v265 = v258
	v266 = v259
	v267 = v260
	goto L67
L77:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270))))
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v275 == v276 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v296 = v275 - v276
	goto L65
L79:
	;
	v278 = int32(1)
	v283 = v272 - v278
	if v283 != 0 {
		v270 = v270 + v278
		v271 = v271 + v278
		v272 = v283
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	goto L55
L84:
	;
	v307 = v303 & int32(1073741823)
	v341 = v307
	v342 = v307
	v344 = v175
	goto L51
L85:
	;
	base.MemoryCopy(m, v187, v45, v76)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v313 = v187 + v76
	v316 = (v313 - v182) & int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v316
	if v79 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v325 = v313
	v326 = v316 | int32(1073741824)
	goto L90
L89:
	;
	if v128 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v326
	v393 = v325
	v398 = v201 + int32(1)
	v399 = base.B2i32(v309 == int32(0))
	goto L50
L91:
	;
	base.MemoryCopy(m, v313, v127, v128)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v321 = v313 + v128
	v325 = v321
	v326 = (v321 - v182) & int32(1073741823)
	goto L90
L94:
	;
	v353 = v341
	goto L96
L95:
	;
	v353 = v350
	goto L96
L96:
	;
	v355 = v342 + (v347&int32(1073741823) - v353)
	if v355 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	base.MemoryCopy(m, v187, v344, v355)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v358 = int32(1073741823)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v362 = v360 & v358
	v363 = int32(0)
	if v363 <= v357 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v366 = v362
	goto L102
L101:
	;
	v366 = v363
	goto L102
L102:
	;
	v367 = v357&v358 - v366
	if int32(0) <= v360 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v302-int32(4))))
	v376 = v362 - v372&int32(1073741823)
	goto L105
L104:
	;
	v376 = v362
	goto L105
L105:
	;
	v378 = v367 + (v376 + v187)
	v379 = v378 - v182
	v381 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = (v379 - v367) & v381
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v384&int32(1073741824) | v379&v381
	v393 = v378
	v398 = v201
	v399 = int32(1)
	goto L50
L106:
	;
	goto L49
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = (v401<<(uint(int32(3))%32)+v413)<<(uint(int32(2))%32) + int32(32)
	v445 = v160
	goto L41
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v401 | int32(-2147483648)
	if v413 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	base.MemoryCopy(m, v179+v401<<(uint(int32(3))%32)&int32(2147483640), v182, v413)
	goto L107
L110:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_hstore_subscript_assign_0), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_hstore_subscript_assign_1), int32(157), int32(_a_F_hstore_subscript_assign_2))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hstore_to_array_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int64
	_ = v125
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(8589934592)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(4294967297)
	v24 = v18 & int32(268435455)
	if v24 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v156
L2:
	;
	v28 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v33 = l0 + int32(8)
	v36 = v33 + v24<<(uint(int32(3))%32)
	v38 = v24 << (uint(int32(1)) % 32)
	v39 = base.I32_div_u_s(v38, l1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v39
	v43 = F_palloc(m, v24<<(uint(int32(4))%32))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v156 = v28
	goto L1
L7:
	;
	v45 = F_palloc(m, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v48 = int32(0)
	goto L9
L9:
	;
	v62 = v48 << (uint(int32(1)) % 32)
	v63 = int32(3)
	v68 = v33 + v48<<(uint(v63)%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v141 = F_construct_md_array(m, v43, v45, l1, v16+int32(8), v16, int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L25
	}
L11:
	;
	v86 = F_cstring_to_text_with_len(m, v85, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L15
	}
L12:
	;
	v83 = v69 & int32(1073741823)
	v85 = v36
	goto L11
L13:
	;
	goto L14
L14:
	;
	v74 = int32(1073741823)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v68-int32(4))))
	v80 = v78 & v74
	v83 = v69&v74 - v80
	v85 = v80 + v36
	goto L11
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43+v62<<(uint(v63)%32)))) = base.I64_extend_i32_u(v86)
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v45))) = uint8(v91)
	v93 = int32(1)
	v96 = v62 | v93
	v99 = v33 + v96<<(uint(int32(2))%32)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100&int32(1073741824) == v91 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v100 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v123 = v93
	v125 = int64(0)
	goto L18
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v43+v96<<(uint(int32(3))%32)))) = v125
	*(*uint8)(unsafe.Add(mBase, uint32(v96+v45))) = uint8(v123)
	v133 = v48 + int32(1)
	if v133 != v24 {
		v48 = v133
		goto L9
	} else {
		goto L24
	}
L19:
	;
	v119 = F_cstring_to_text_with_len(m, v118, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L23
	}
L20:
	;
	v116 = v100 & int32(1073741823)
	v118 = v36
	goto L19
L21:
	;
	goto L22
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v99-int32(4))))
	v113 = v111 & int32(1073741823)
	v116 = v100 - v113
	v118 = v113 + v36
	goto L19
L23:
	;
	v123 = int32(0)
	v125 = base.I64_extend_i32_u(v119)
	goto L18
L24:
	;
	goto L10
L25:
	;
	v156 = v141
	goto L1
}
func F_hstore_to_matrix(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_hstoreUpgrade(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_hstore_to_array_internal(m, v3, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
