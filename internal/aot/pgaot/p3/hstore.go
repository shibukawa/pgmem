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
func F_hstoreUpgrade(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
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
	var v104 int32
	_ = v104
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v373 int32
	_ = v373
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v482 int32
	_ = v482
	v14 = F_pg_detoast_datum(m, l0)
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
	if int32(0) <= v18 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l0 == v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v482 = v14
	goto L5
L5:
	;
	return v482
L6:
	;
	v22 = F_pg_detoast_datum_copy(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v25 = v18
	v26 = v14
	goto L8
L8:
	;
	v28 = v26 + int32(4)
	if v25 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v25 = v24
	v26 = v22
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v466
	v482 = v26
	goto L5
L11:
	;
	v59 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v72 = v70 & int32(268435455)
	if base.B2i32(v72 == v59)|base.B2i32(v70 < v59) != 0 {
		v203 = int32(2)
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v466 = (v55 + v53) << (uint(int32(2)) % 32)
	goto L10
L13:
	;
	v48 = v40 << (uint(int32(3)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v28)))
	v53 = v48 + int32(8)
	v55 = v52
	goto L12
L14:
	;
	v53 = int32(8)
	v55 = int32(0)
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(-2147483648)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if base.Ui32(int32(_a_F_hstoreUpgrade_0)) < base.Ui32(v33) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if int32(0) <= v36 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v40 = v25 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v40 | int32(-2147483648)
	if v40 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v25) {
		goto L55
	} else {
		goto L56
	}
L22:
	;
	v222 = v203
	goto L21
L23:
	;
	v79 = v26 + int32(8)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if int32(0) <= v80 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v222 = int32(0)
	goto L21
L25:
	;
	goto L26
L26:
	;
	v84 = int32(0)
	v86 = v72 << (uint(int32(3)) % 32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v79-int32(4))))
	v95 = v86 + v90&int32(1073741823) + int32(8)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v98 = int32(base.Ui32(v96) >> (uint(int32(2)) % 32))
	if base.Ui32(v98) < base.Ui32(v95) {
		v203 = v84
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v100 = int32(1)
	v104 = v100
	goto L28
L28:
	;
	v116 = v79 + v104<<(uint(int32(2))%32)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 < int32(0) {
		v203 = v84
		goto L22
	} else {
		goto L30
	}
L29:
	;
	if v72 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v120 = int32(1073741823)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(4))))
	if base.Ui32(v117&v120) < base.Ui32(v124&v120) {
		v203 = v84
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v129 = v104 + int32(1)
	if v129 != v72<<(uint(v100)%32) {
		v104 = v129
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v133 = int32(2)
	if base.Ui32(v72) <= base.Ui32(v133) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v95 == v98 {
		goto L49
	} else {
		goto L50
	}
L36:
	;
	v136 = v133
	goto L38
L37:
	;
	v136 = v72
	goto L38
L38:
	;
	v140 = int32(1)
	goto L39
L39:
	;
	v150 = v140 << (uint(int32(3)) % 32)
	v151 = v79 + v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v154 = v152 & int32(1073741823)
	if int32(0) <= v152 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L35
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v151-int32(4))))
	v163 = v154 - v159&int32(1073741823)
	goto L43
L42:
	;
	v163 = v154
	goto L43
L43:
	;
	v164 = v26 + v150
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	v167 = v165 & int32(1073741823)
	if int32(0) <= v165 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v164-int32(4))))
	v176 = v167 - v172&int32(1073741823)
	goto L46
L45:
	;
	v176 = v167
	goto L46
L46:
	;
	if v152&int32(1073741824)|base.B2i32(base.Ui32(v163) < base.Ui32(v176)) != 0 {
		v203 = int32(0)
		goto L22
	} else {
		goto L47
	}
L47:
	;
	v183 = v140 + int32(1)
	if v183 != v136 {
		v140 = v183
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
L49:
	;
	v199 = int32(2)
	goto L51
L50:
	;
	v199 = int32(1)
	goto L51
L51:
	;
	v203 = v199
	goto L22
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v450
	v466 = v463
	goto L10
L53:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v434 = int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v433 | v434
	v440 = v332 << (uint(int32(3)) % 32)
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v440+v237-int32(4))))
	v450 = v332 | v434
	v463 = (v440+v444)<<(uint(int32(2))%32) + int32(32)
	goto L52
L54:
	;
	if v332 != 0 {
		goto L53
	} else {
		goto L99
	}
L55:
	;
	if v222 != 0 {
		goto L90
	} else {
		goto L91
	}
L56:
	;
	v228 = v25<<(uint(int32(3))%32) + int32(8)
	v230 = int32(base.Ui32(v33) >> (uint(int32(2)) % 32))
	if base.Ui32(v230) < base.Ui32(v228) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v232 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if base.Ui32(v232) < base.Ui32(v233) {
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v237 = v26 + int32(8)
	if v25 != int32(1) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v240 = v232
	goto L62
L60:
	;
	goto L61
L61:
	;
	v276 = int32(1)
	if v25 <= v276 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v254 = v240 << (uint(int32(3)) % 32)
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237+v254))))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v254+v26))))
	if base.Ui32(v256) < base.Ui32(v258) {
		goto L55
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	v261 = v240 + int32(1)
	if v261 != v25 {
		v240 = v261
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v279 = v276
	goto L68
L67:
	;
	v279 = v25
	goto L68
L68:
	;
	v280 = int32(0)
	v282 = v280
	v284 = v280
	goto L69
L69:
	;
	v297 = v237 + v284<<(uint(int32(3))%32)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	if int32(base.Ui32(v298)>>(uint(int32(1))%32)) != v282 {
		goto L55
	} else {
		goto L71
	}
L70:
	;
	if base.Ui32(v230) < base.Ui32(v309+v228) {
		goto L55
	} else {
		goto L76
	}
L71:
	;
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297))))
	if v298&int32(1) != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v307 = int32(0)
	goto L74
L73:
	;
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+2)))
	v307 = v306
	goto L74
L74:
	;
	v309 = v307 + (v282 + v302)
	v311 = v284 + int32(1)
	if v311 != v279 {
		v282 = v309
		v284 = v311
		goto L69
	} else {
		goto L75
	}
L75:
	;
	goto L70
L76:
	;
	if v222 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v332 <= int32(0) {
		goto L54
	} else {
		goto L83
	}
L78:
	;
	v319 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	if v319 == int32(0) {
		goto L77
	} else {
		goto L80
	}
L80:
	;
	F_errmsg_internal(m, int32(_a_F_hstoreUpgrade_1), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_hstoreUpgrade_2), int32(312), int32(_a_F_hstoreUpgrade_3))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L77
L83:
	;
	v336 = int32(0)
	goto L84
L84:
	;
	v351 = v237 + v336<<(uint(int32(3))%32)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351)+2)))
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v351)+4))
	v355 = int32(1)
	v357 = v353 + int32(base.Ui32(v354)>>(uint(v355)%32))
	*(*int32)(unsafe.Add(mBase, uint32(v351))) = v357 & int32(1073741823)
	v363 = v354 & v355
	if v363 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L53
L86:
	;
	v364 = int32(0)
	goto L88
L87:
	;
	v364 = v352
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v351)+4)) = (v364+v357)&int32(1073741823) | v363<<(uint(int32(30))%32)
	v373 = v336 + int32(1)
	if v373 != v332 {
		v336 = v373
		goto L84
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v389 = v25 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v389 | int32(-2147483648)
	if v389 == int32(0) {
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
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L96
	}
L93:
	;
	v466 = int32(32)
	goto L10
L94:
	;
	goto L95
L95:
	;
	v397 = v389 << (uint(int32(3)) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v397+v28)))
	v466 = (v397+v399)<<(uint(int32(2))%32) + int32(32)
	goto L10
L96:
	;
	F_errmsg_internal(m, int32(_a_F_hstoreUpgrade_4), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_hstoreUpgrade_2), int32(274), int32(_a_F_hstoreUpgrade_3))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
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
	v450 = int32(-2147483648)
	v463 = int32(32)
	goto L52
}
func F_hstore_avals(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_hstoreUpgrade(m, v15)
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
	v22 = v20 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(1)
	if v22 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return v112
L4:
	;
	v29 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v32 = v16 + int32(8)
	v37 = v32 + v20<<(uint(int32(3))%32)&int32(2147483640)
	v41 = F_palloc(m, v22<<(uint(int32(2))%32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v112 = v29
	goto L3
L8:
	;
	v43 = F_palloc(m, v22)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v45 = int32(0)
	goto L10
L10:
	;
	v61 = v32 + v45<<(uint(int32(3))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v62&int32(1073741824) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v100 = F_construct_md_array(m, v41, v43, int32(1), v13+int32(12), v13+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L12:
	;
	v83 = int32(1)
	v84 = int32(0)
	goto L14
L13:
	;
	v66 = int32(0)
	if v62 < v66 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41+v45<<(uint(int32(2))%32)))) = v84
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v43))) = uint8(v83)
	v89 = v45 + int32(1)
	if v89 != v22 {
		v45 = v89
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v79 = F_cstring_to_text_with_len(m, v78, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v76 = v62 & int32(1073741823)
	v78 = v37
	goto L15
L17:
	;
	goto L18
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v73 = v71 & int32(1073741823)
	v76 = v62 - v73
	v78 = v73 + v37
	goto L15
L19:
	;
	v83 = v66
	v84 = v79
	goto L14
L20:
	;
	goto L11
L21:
	;
	v112 = v100
	goto L3
}
func F_hstore_contains(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_hstoreUpgrade(m, v20)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = F_hstoreUpgrade(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v30 = v28 & int32(268435455)
	if v30 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(1)
L5:
	;
	goto L6
L6:
	;
	v35 = int32(8)
	v36 = v21 + v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v38 = int32(3)
	v44 = v26 + v35
	v47 = v44 + v30<<(uint(v38)%32)
	v49 = v37 & int32(268435455)
	v54 = int32(0)
	v65 = int32(0)
	goto L7
L7:
	;
	v75 = v44 + v65<<(uint(int32(3))%32)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v76 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v327
L9:
	;
	if v49 <= v54 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v91 = v76 & int32(1073741823)
	v92 = v47
	goto L9
L11:
	;
	goto L12
L12:
	;
	v81 = int32(1073741823)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75-int32(4))))
	v87 = v85 & v81
	v91 = v76&v81 - v87
	v92 = v87 + v47
	goto L9
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v96 = v54
	v101 = v49
	goto L19
L16:
	;
	goto L8
L17:
	;
	v318 = int32(1)
	v322 = v65 + v318
	if base.Ui32(v322) < base.Ui32(v30) {
		v54 = v118 + v318
		v65 = v322
		goto L7
	} else {
		goto L89
	}
L18:
	;
	if v223 <= v213 {
		goto L67
	} else {
		goto L68
	}
L19:
	;
	v118 = int32(base.Ui32(v101-v96)>>(uint(int32(1))%32)) + v96
	v121 = v36 + v118<<(uint(int32(3))%32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v124 = v122 & int32(1073741823)
	if int32(0) <= v122 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	if v227 != v229 {
		v327 = v207
		goto L16
	} else {
		goto L66
	}
L21:
	;
	goto L20
L22:
	;
	v236 = int32(0)
	v240 = base.B2i32(v235 < v236)
	if v235 < v236 {
		goto L59
	} else {
		goto L60
	}
L23:
	;
	v144 = v143 + (v36 + v49<<(uint(v38)%32))
	if base.Ui32(int32(4)) <= base.Ui32(v91) {
		goto L36
	} else {
		goto L37
	}
L24:
	;
	if base.Ui32(v91) < base.Ui32(v136) {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121-int32(4))))
	v131 = v129 & int32(1073741823)
	v132 = v124 - v131
	if v132 != v91 {
		v136 = v132
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v124 == v91 {
		v143 = int32(0)
		goto L23
	} else {
		goto L29
	}
L28:
	;
	v143 = v131
	goto L23
L29:
	;
	v136 = v124
	goto L24
L30:
	;
	v141 = int32(1)
	goto L32
L31:
	;
	v141 = int32(-1)
	goto L32
L32:
	;
	v235 = v141
	goto L22
L33:
	;
	if v206 != 0 {
		v235 = v206
		goto L22
	} else {
		goto L51
	}
L34:
	;
	v206 = int32(0)
	goto L33
L35:
	;
	v180 = v175
	v181 = v176
	v182 = v177
	goto L45
L36:
	;
	if (v144|v92)&int32(3) != 0 {
		v175 = v144
		v176 = v92
		v177 = v91
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v168 = v144
	v169 = v92
	v170 = v91
	goto L38
L38:
	;
	if v170 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v152 = v144
	v153 = v92
	v154 = v91
	goto L40
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v157 != v158 {
		v175 = v152
		v176 = v153
		v177 = v154
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v168 = v163
	v169 = v161
	v170 = v165
	goto L38
L42:
	;
	v160 = int32(4)
	v161 = v153 + v160
	v163 = v152 + v160
	v165 = v154 - v160
	if base.Ui32(int32(3)) < base.Ui32(v165) {
		v152 = v163
		v153 = v161
		v154 = v165
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v175 = v168
	v176 = v169
	v177 = v170
	goto L35
L45:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v185 == v186 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v206 = v185 - v186
	goto L33
L47:
	;
	v188 = int32(1)
	v193 = v182 - v188
	if v193 != 0 {
		v180 = v180 + v188
		v181 = v181 + v188
		v182 = v193
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
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v209 = int32(30)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v215 = v213 & int32(1073741824)
	if int32(base.Ui32(v208)>>(uint(v209)%32))&int32(1) != int32(base.Ui32(v215)>>(uint(v209)%32)) {
		v327 = v207
		goto L16
	} else {
		goto L52
	}
L52:
	;
	if v215 != 0 {
		goto L17
	} else {
		goto L53
	}
L53:
	;
	v219 = int32(1073741823)
	v222 = v76 & v219
	v223 = int32(0)
	v225 = base.B2i32(v223 <= v213)
	if v223 <= v213 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v226 = v222
	goto L56
L55:
	;
	v226 = v223
	goto L56
L56:
	;
	v227 = v213&v219 - v226
	v229 = v208 & int32(1073741823)
	if v208 < int32(0) {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	if v227 == v229-v124 {
		v246 = v124
		goto L18
	} else {
		goto L58
	}
L58:
	;
	v327 = v207
	goto L16
L59:
	;
	v241 = v118 + int32(1)
	goto L61
L60:
	;
	v241 = v96
	goto L61
L61:
	;
	if v235 < v236 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v242 = v101
	goto L64
L63:
	;
	v242 = v118
	goto L64
L64:
	;
	if v241 < v242 {
		v96 = v241
		v101 = v242
		goto L19
	} else {
		goto L65
	}
L65:
	;
	v327 = v236
	goto L16
L66:
	;
	v246 = int32(0)
	goto L18
L67:
	;
	v248 = v222
	goto L69
L68:
	;
	v248 = int32(0)
	goto L69
L69:
	;
	v249 = v47 + v248
	v250 = v246 + (v36 + v37<<(uint(v38)%32)&int32(2147483640))
	if base.Ui32(int32(4)) <= base.Ui32(v227) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v312 != 0 {
		v327 = v207
		goto L16
	} else {
		goto L88
	}
L71:
	;
	v312 = int32(0)
	goto L70
L72:
	;
	v286 = v281
	v287 = v282
	v288 = v283
	goto L82
L73:
	;
	if (v249|v250)&int32(3) != 0 {
		v281 = v249
		v282 = v250
		v283 = v227
		goto L72
	} else {
		goto L76
	}
L74:
	;
	v274 = v249
	v275 = v250
	v276 = v227
	goto L75
L75:
	;
	if v276 == int32(0) {
		goto L71
	} else {
		goto L81
	}
L76:
	;
	v258 = v249
	v259 = v250
	v260 = v227
	goto L77
L77:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	if v263 != v264 {
		v281 = v258
		v282 = v259
		v283 = v260
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v274 = v269
	v275 = v267
	v276 = v271
	goto L75
L79:
	;
	v266 = int32(4)
	v267 = v259 + v266
	v269 = v258 + v266
	v271 = v260 - v266
	if base.Ui32(int32(3)) < base.Ui32(v271) {
		v258 = v269
		v259 = v267
		v260 = v271
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v281 = v274
	v282 = v275
	v283 = v276
	goto L72
L82:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286))))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287))))
	if v291 == v292 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v312 = v291 - v292
	goto L70
L84:
	;
	v294 = int32(1)
	v299 = v288 - v294
	if v299 != 0 {
		v286 = v286 + v294
		v287 = v287 + v294
		v288 = v299
		goto L82
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	goto L71
L88:
	;
	goto L17
L89:
	;
	v327 = v318
	goto L16
}
func F_hstore_delete(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
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
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	v2 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v26 = int32(1)
	v27 = v24 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	v30 = v28 & v26
	if v28 == v26 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v61 = F_palloc(m, int32(base.Ui32(v58)>>(uint(int32(2))%32)))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v36 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v47 = int32(1)
	if v30 != 0 {
		v57 = int32(base.Ui32(v28)>>(uint(v47)%32)) - v47
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v39 = int32(16)
	goto L10
L9:
	;
	v39 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = int32(4)
	goto L13
L12:
	;
	v46 = v39
	goto L13
L13:
	;
	v57 = v46
	goto L4
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v64 & int32(-4)
	v69 = v63 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v69 | int32(-2147483648)
	v74 = v61 + int32(8)
	v77 = v74 + v69<<(uint(int32(3))%32)
	if v69 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v244 == v256&int32(268435455) {
		goto L62
	} else {
		goto L63
	}
L17:
	;
	v240 = int32(0)
	v244 = v2
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v30 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v83 = v27
	goto L22
L21:
	;
	v83 = v24 + int32(4)
	goto L22
L22:
	;
	v85 = v19 + int32(8)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v93 = v74
	v96 = v77
	v98 = v2
	v99 = int32(0)
	goto L23
L23:
	;
	v112 = v85 + v99<<(uint(int32(3))%32)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v115 = v113 & int32(1073741823)
	if v113 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v231 = v224 - v77
	if v225 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L25:
	;
	v127 = v126 + (v85 + v86<<(uint(int32(3))%32)&int32(2147483640))
	if v125 == v57 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v125 = v115
	v126 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v112-int32(4))))
	v123 = v121 & int32(1073741823)
	v125 = v115 - v123
	v126 = v123
	goto L25
L29:
	;
	v229 = v99 + int32(1)
	if v229 != v69 {
		v93 = v222
		v96 = v224
		v98 = v225
		v99 = v229
		goto L23
	} else {
		goto L58
	}
L30:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v57) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	goto L32
L32:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v194 = int32(1073741823)
	v198 = int32(0)
	if v198 <= v193 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	if v190 == int32(0) {
		v222 = v93
		v224 = v96
		v225 = v98
		goto L29
	} else {
		goto L51
	}
L34:
	;
	v190 = int32(0)
	goto L33
L35:
	;
	v164 = v159
	v165 = v160
	v166 = v161
	goto L45
L36:
	;
	if (v127|v83)&int32(3) != 0 {
		v159 = v127
		v160 = v83
		v161 = v57
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v152 = v127
	v153 = v83
	v154 = v57
	goto L38
L38:
	;
	if v154 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v136 = v127
	v137 = v83
	v138 = v57
	goto L40
L40:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	if v141 != v142 {
		v159 = v136
		v160 = v137
		v161 = v138
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v152 = v147
	v153 = v145
	v154 = v149
	goto L38
L42:
	;
	v144 = int32(4)
	v145 = v137 + v144
	v147 = v136 + v144
	v149 = v138 - v144
	if base.Ui32(int32(3)) < base.Ui32(v149) {
		v136 = v147
		v137 = v145
		v138 = v149
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v159 = v152
	v160 = v153
	v161 = v154
	goto L35
L45:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v169 == v170 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v190 = v169 - v170
	goto L33
L47:
	;
	v172 = int32(1)
	v177 = v166 - v172
	if v177 != 0 {
		v164 = v164 + v172
		v165 = v165 + v172
		v166 = v177
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
	v201 = v113 & v194
	goto L54
L53:
	;
	v201 = v198
	goto L54
L54:
	;
	v202 = v193&v194 - v201
	v203 = v125 + v202
	if v203 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	base.MemoryCopy(m, v96, v127, v203)
	goto L57
L56:
	;
	goto L57
L57:
	;
	v205 = v203 + v96
	v206 = v205 - v77
	v208 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = (v206 - v202) & v208
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v206&v208 | v213&int32(1073741824)
	v222 = v93 + int32(8)
	v224 = v205
	v225 = v98 + int32(1)
	goto L29
L58:
	;
	goto L24
L59:
	;
	v240 = v231
	v244 = int32(0)
	goto L16
L60:
	;
	goto L61
L61:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v235 | int32(-2147483648)
	v240 = v231
	v244 = v225
	goto L16
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = (v244<<(uint(int32(3))%32)+v240)<<(uint(int32(2))%32) + int32(32)
	return v61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v244 | int32(-2147483648)
	if v240 == int32(0) {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	base.MemoryCopy(m, v74+v244<<(uint(int32(3))%32)&int32(2147483640), v77, v240)
	goto L62
}
func F_hstore_each(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v17 = F_hstoreUpgrade(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(_a_F_hstore_each_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0])) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v31 = F_palloc(m, int32(base.Ui32(v28)>>(uint(int32(2))%32)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v35 = int32(base.Ui32(v33) >> (uint(int32(2)) % 32))
					if v35 != 0 {
						base.MemoryCopy(m, v31, v17, v35)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v31
					v41 = F_get_call_result_type(m, l0, int32(0), v10+int32(12))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						if v41 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v155 = m.ExcPending
							if v155 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_hstore_each_1), int32(0))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_hstore_each_2), int32(869), int32(_a_F_hstore_each_3))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v46 = F_BlessTupleDesc(m, v45)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v46
								*(*int32)(unsafe.Add(mBase, _c_F_hstore_each[0])) = v24
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = v59 & int32(268435455)
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								if base.Ui32(v62) < base.Ui32(v61) {
									v64 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)) = uint16(v64)
									v67 = v58 + int32(8)
									v68 = int32(3)
									v70 = v67 + v61<<(uint(v68)%32)
									v73 = v67 + v62<<(uint(v68)%32)
									v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
									if v74 < v64 {
										v88 = v74 & int32(1073741823)
										v90 = v70
									} else {
										v79 = int32(1073741823)
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v73-int32(4))))
										v85 = v83 & v79
										v88 = v74&v79 - v85
										v90 = v70 + v85
									}
									v91 = F_cstring_to_text_with_len(m, v90, v88)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v91
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
										if v94&int32(1073741824) != 0 {
											v97 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)) = uint8(v97)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
											v124 = F_heap_form_tuple(m, v119, v10+int32(4), v10+int32(2))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
												v127 = F_HeapTupleHeaderGetDatum(m, v126)
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
													return int32(0)
												} else {
													v129 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
													*(*int64)(unsafe.Add(mBase, uint32(v57))) = v129 + int64(1)
													v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = int32(1)
													v145 = v127
													m.G0 = v10 + int32(16)
													return v145
												}
											}
										} else {
											if v94 < int32(0) {
												v111 = v70
												v112 = v94 & int32(1073741823)
											} else {
												v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
												v107 = v105 & int32(1073741823)
												v111 = v107 + v70
												v112 = v94 - v107
											}
											v113 = F_cstring_to_text_with_len(m, v111, v112)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v113
												v119 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
												v124 = F_heap_form_tuple(m, v119, v10+int32(4), v10+int32(2))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
													v127 = F_HeapTupleHeaderGetDatum(m, v126)
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int32(0)
													} else {
														v129 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
														*(*int64)(unsafe.Add(mBase, uint32(v57))) = v129 + int64(1)
														v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = int32(1)
														v145 = v127
														m.G0 = v10 + int32(16)
														return v145
													}
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return int32(0)
									} else {
										v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v138)+20)) = int32(2)
										v141 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v141)
										v145 = int32(0)
										m.G0 = v10 + int32(16)
										return v145
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
		v61 = v59 & int32(268435455)
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
		if base.Ui32(v62) < base.Ui32(v61) {
			v64 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v10)+2)) = uint16(v64)
			v67 = v58 + int32(8)
			v68 = int32(3)
			v70 = v67 + v61<<(uint(v68)%32)
			v73 = v67 + v62<<(uint(v68)%32)
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
			if v74 < v64 {
				v88 = v74 & int32(1073741823)
				v90 = v70
			} else {
				v79 = int32(1073741823)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(v73-int32(4))))
				v85 = v83 & v79
				v88 = v74&v79 - v85
				v90 = v70 + v85
			}
			v91 = F_cstring_to_text_with_len(m, v90, v88)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v91
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
				if v94&int32(1073741824) != 0 {
					v97 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)) = uint8(v97)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
					v119 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
					v124 = F_heap_form_tuple(m, v119, v10+int32(4), v10+int32(2))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
						v127 = F_HeapTupleHeaderGetDatum(m, v126)
						mBase = m.M
						v128 = m.ExcPending
						if v128 != 0 {
							return int32(0)
						} else {
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = v129 + int64(1)
							v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = int32(1)
							v145 = v127
							m.G0 = v10 + int32(16)
							return v145
						}
					}
				} else {
					if v94 < int32(0) {
						v111 = v70
						v112 = v94 & int32(1073741823)
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
						v107 = v105 & int32(1073741823)
						v111 = v107 + v70
						v112 = v94 - v107
					}
					v113 = F_cstring_to_text_with_len(m, v111, v112)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v113
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
						v124 = F_heap_form_tuple(m, v119, v10+int32(4), v10+int32(2))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
							v127 = F_HeapTupleHeaderGetDatum(m, v126)
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								v129 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = v129 + int64(1)
								v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v133)+20)) = int32(1)
								v145 = v127
								m.G0 = v10 + int32(16)
								return v145
							}
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return int32(0)
			} else {
				v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v138)+20)) = int32(2)
				v141 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v141)
				v145 = int32(0)
				m.G0 = v10 + int32(16)
				return v145
			}
		}
	}
}
func F_hstore_exists_any(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v198 int32
	_ = v198
	var v231 int32
	_ = v231
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_hstoreUpgrade(m, v20)
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = F_hstoreArrayToPairs(m, v26, v18+int32(12))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if int32(0) < v32 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	m.G0 = v18 + int32(16)
	return v231
L6:
	;
	v36 = v21 + int32(8)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v39 = v37 & int32(268435455)
	v44 = int32(0)
	v51 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v231 = int32(0)
	goto L5
L9:
	;
	if v44 < v39 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v62 = v30 + v51*int32(20)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v65 = v44
	v67 = v39
	goto L14
L12:
	;
	v182 = v44
	goto L13
L13:
	;
	v198 = v51 + int32(1)
	if v198 != v32 {
		v44 = v182
		v51 = v198
		goto L9
	} else {
		goto L53
	}
L14:
	;
	v83 = int32(base.Ui32(v67-v65)>>(uint(int32(1))%32)) + v65
	v86 = v36 + v83<<(uint(int32(3))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v89 = v87 & int32(1073741823)
	if int32(0) <= v87 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v182 = v179
	goto L13
L16:
	;
	v178 = base.B2i32(v173 < int32(0))
	if v173 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L17:
	;
	v109 = v108 + (v36 + v39<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v63) {
		goto L30
	} else {
		goto L31
	}
L18:
	;
	if base.Ui32(v63) < base.Ui32(v101) {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86-int32(4))))
	v96 = v94 & int32(1073741823)
	v97 = v89 - v96
	if v97 != v63 {
		v101 = v97
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v89 == v63 {
		v108 = int32(0)
		goto L17
	} else {
		goto L23
	}
L22:
	;
	v108 = v96
	goto L17
L23:
	;
	v101 = v89
	goto L18
L24:
	;
	v106 = int32(1)
	goto L26
L25:
	;
	v106 = int32(-1)
	goto L26
L26:
	;
	v173 = v106
	goto L16
L27:
	;
	if v171 != 0 {
		v173 = v171
		goto L16
	} else {
		goto L45
	}
L28:
	;
	v171 = int32(0)
	goto L27
L29:
	;
	v145 = v140
	v146 = v141
	v147 = v142
	goto L39
L30:
	;
	if (v109|v64)&int32(3) != 0 {
		v140 = v109
		v141 = v64
		v142 = v63
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v133 = v109
	v134 = v64
	v135 = v63
	goto L32
L32:
	;
	if v135 == int32(0) {
		goto L28
	} else {
		goto L38
	}
L33:
	;
	v117 = v109
	v118 = v64
	v119 = v63
	goto L34
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v122 != v123 {
		v140 = v117
		v141 = v118
		v142 = v119
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v133 = v128
	v134 = v126
	v135 = v130
	goto L32
L36:
	;
	v125 = int32(4)
	v126 = v118 + v125
	v128 = v117 + v125
	v130 = v119 - v125
	if base.Ui32(int32(3)) < base.Ui32(v130) {
		v117 = v128
		v118 = v126
		v119 = v130
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L29
L39:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 == v151 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v171 = v150 - v151
	goto L27
L41:
	;
	v153 = int32(1)
	v158 = v147 - v153
	if v158 != 0 {
		v145 = v145 + v153
		v146 = v146 + v153
		v147 = v158
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L40
L44:
	;
	goto L28
L45:
	;
	v231 = int32(1)
	goto L5
L46:
	;
	v179 = v83 + int32(1)
	goto L48
L47:
	;
	v179 = v65
	goto L48
L48:
	;
	if v173 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v180 = v67
	goto L51
L50:
	;
	v180 = v83
	goto L51
L51:
	;
	if v179 < v180 {
		v65 = v179
		v67 = v180
		goto L14
	} else {
		goto L52
	}
L52:
	;
	goto L15
L53:
	;
	goto L10
}
func F_hstore_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_gt_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v6)
	}
}
func F_hstore_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(_a_F_hstore_ne_0), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 != int32(0))
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
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
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v446 int32
	_ = v446
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
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
	v474 = m.ExcPending
	if v474 != 0 {
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
	v40 = v35 + v39
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v43&v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = v40
	goto L8
L7:
	;
	v46 = v35 + int32(4)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = v46
	if v43 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v77 = F_hstoreCheckKeyLen(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v53 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v64 = int32(1)
	if v43&v64 != 0 {
		v76 = int32(base.Ui32(v43)>>(uint(v64)%32)) - v64
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v56 = int32(16)
	goto L15
L14:
	;
	v56 = int32(0)
	goto L15
L15:
	;
	if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v63 = int32(4)
	goto L18
L17:
	;
	v63 = v56
	goto L18
L18:
	;
	v76 = v63
	goto L9
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v77
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+44)))
	if v80 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+28)) = uint8(v125)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v131
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	if v135 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L22:
	;
	v125 = int32(1)
	v130 = v4
	v131 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v28)+40))
	v84 = F_pg_detoast_datum_packed(m, v83)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v86 = int32(1)
	v87 = v84 + v86
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v92 = v90 & v86
	if v92 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = v87
	goto L28
L27:
	;
	v93 = v84 + int32(4)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v93
	if v90 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v123 = F_hstoreCheckValLen(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L40
	}
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v101 == int32(18) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v112 = int32(1)
	if v92 != 0 {
		v122 = int32(base.Ui32(v90)>>(uint(v112)%32)) - v112
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v104 = int32(16)
	goto L35
L34:
	;
	v104 = int32(0)
	goto L35
L35:
	;
	if base.Ui32((v101-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v111 = int32(4)
	goto L38
L37:
	;
	v111 = v104
	goto L38
L38:
	;
	v122 = v111
	goto L29
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v122 = int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	v125 = int32(0)
	v130 = v93
	v131 = v123
	goto L21
L41:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v446
	v465 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v466 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v465))) = uint8(v466)
	m.G0 = v26 + int32(32)
	return
L42:
	;
	v142 = F_hstorePairs(m, v26+int32(12), int32(1), v77+v131)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v146 = F_hstoreUpgrade(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	v446 = v142
	goto L41
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v154 = v152 & int32(268435455)
	v157 = int32(8)
	v158 = v154<<(uint(int32(3))%32) + v157
	v162 = v77 + int32(base.Ui32(v148)>>(uint(int32(2))%32)) + v158 + v131 + v157
	v163 = F_palloc(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v154 - int32(2147483647)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v162 << (uint(int32(2)) % 32)
	v171 = int32(8)
	v172 = v146 + v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v176 = int32(2147483640)
	v178 = v172 + v173<<(uint(int32(3))%32)&v176
	v182 = v163 + v171
	v185 = v182 + v158&v176
	v186 = v182
	v189 = base.B2i32(v154 != int32(0))
	v190 = v185
	v201 = v4
	v202 = v4
	v204 = v4
	goto L48
L48:
	;
	if v189 == int32(0) {
		v312 = int32(1)
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v412 | int32(-2147483648)
	v416 = v396 - v185
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v404 == v417&int32(268435455) {
		goto L107
	} else {
		goto L108
	}
L50:
	;
	v404 = v201 + int32(1)
	v407 = v402 + v204
	v408 = base.B2i32(base.Ui32(v407) < base.Ui32(v154))
	if v408|base.B2i32(v401 <= int32(0)) != 0 {
		v186 = v186 + int32(8)
		v189 = v408
		v190 = v396
		v201 = v404
		v202 = v401
		v204 = v407
		goto L48
	} else {
		goto L106
	}
L51:
	;
	v349 = v305 + int32(4)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v353 = int32(0)
	if v353 <= v350 {
		goto L94
	} else {
		goto L95
	}
L52:
	;
	v335 = int32(1073741823)
	v336 = v306 & v335
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v305-int32(4))))
	v341 = v339 & v335
	v344 = v336
	v345 = v336 - v341
	v347 = v178 + v341
	goto L51
L53:
	;
	if v77 != 0 {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	v213 = v204 << (uint(int32(3)) % 32)
	if int32(0) < v202 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v305 = v213 + v172
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	if int32(0) <= v306 {
		goto L52
	} else {
		goto L84
	}
L56:
	;
	v216 = v213 + v172
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v219 = v217 & int32(1073741823)
	if int32(0) <= v217 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v237 = v178 + v236
	if base.Ui32(int32(4)) <= base.Ui32(v77) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v231 <= v77 {
		goto L55
	} else {
		goto L64
	}
L59:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v216-int32(4))))
	v226 = v224 & int32(1073741823)
	v227 = v219 - v226
	if v227 != v77 {
		v231 = v227
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v219 == v77 {
		v236 = int32(0)
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v236 = v226
	goto L57
L63:
	;
	v231 = v219
	goto L58
L64:
	;
	v312 = int32(1)
	goto L53
L65:
	;
	if int32(0) <= v299 {
		v312 = v299
		goto L53
	} else {
		goto L83
	}
L66:
	;
	v299 = int32(0)
	goto L65
L67:
	;
	v273 = v268
	v274 = v269
	v275 = v270
	goto L77
L68:
	;
	if (v237|v46)&int32(3) != 0 {
		v268 = v237
		v269 = v46
		v270 = v77
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v261 = v237
	v262 = v46
	v263 = v77
	goto L70
L70:
	;
	if v263 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v245 = v237
	v246 = v46
	v247 = v77
	goto L72
L72:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	if v250 != v251 {
		v268 = v245
		v269 = v246
		v270 = v247
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v261 = v256
	v262 = v254
	v263 = v258
	goto L70
L74:
	;
	v253 = int32(4)
	v254 = v246 + v253
	v256 = v245 + v253
	v258 = v247 - v253
	if base.Ui32(int32(3)) < base.Ui32(v258) {
		v245 = v256
		v246 = v254
		v247 = v258
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v268 = v261
	v269 = v262
	v270 = v263
	goto L67
L77:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v278 == v279 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v299 = v278 - v279
	goto L65
L79:
	;
	v281 = int32(1)
	v286 = v275 - v281
	if v286 != 0 {
		v273 = v273 + v281
		v274 = v274 + v281
		v275 = v286
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
	v310 = v306 & int32(1073741823)
	v344 = v310
	v345 = v310
	v347 = v178
	goto L51
L85:
	;
	base.MemoryCopy(m, v190, v46, v77)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v316 = v190 + v77
	v319 = (v316 - v185) & int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = v319
	if v80 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v328 = v316
	v329 = v319 | int32(1073741824)
	goto L90
L89:
	;
	if v131 != 0 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v329
	v396 = v328
	v401 = v202 + int32(1)
	v402 = base.B2i32(v312 == int32(0))
	goto L50
L91:
	;
	base.MemoryCopy(m, v316, v130, v131)
	goto L93
L92:
	;
	goto L93
L93:
	;
	v324 = v316 + v131
	v328 = v324
	v329 = (v324 - v185) & int32(1073741823)
	goto L90
L94:
	;
	v356 = v344
	goto L96
L95:
	;
	v356 = v353
	goto L96
L96:
	;
	v358 = v345 + (v350&int32(1073741823) - v356)
	if v358 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	base.MemoryCopy(m, v190, v347, v358)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	v361 = int32(1073741823)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v365 = v363 & v361
	v366 = int32(0)
	if v366 <= v360 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v369 = v365
	goto L102
L101:
	;
	v369 = v366
	goto L102
L102:
	;
	v370 = v360&v361 - v369
	if int32(0) <= v363 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v305-int32(4))))
	v379 = v365 - v375&int32(1073741823)
	goto L105
L104:
	;
	v379 = v365
	goto L105
L105:
	;
	v381 = v370 + (v379 + v190)
	v382 = v381 - v185
	v384 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = (v382 - v370) & v384
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v387&int32(1073741824) | v382&v384
	v396 = v381
	v401 = v202
	v402 = int32(1)
	goto L50
L106:
	;
	goto L49
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = (v404<<(uint(int32(3))%32)+v416)<<(uint(int32(2))%32) + int32(32)
	v446 = v163
	goto L41
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v404 | int32(-2147483648)
	if v416 == int32(0) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	base.MemoryCopy(m, v182+v404<<(uint(int32(3))%32)&int32(2147483640), v185, v416)
	goto L107
L110:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_hstore_subscript_assign_0), int32(0))
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_hstore_subscript_assign_1), int32(156), int32(_a_F_hstore_subscript_assign_2))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v47 int32
	_ = v47
	var v62 int32
	_ = v62
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
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
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
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v153 int32
	_ = v153
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(8589934592)
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(4294967297)
	v25 = v19 & int32(268435455)
	if v25 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v153
L2:
	;
	v29 = F_construct_empty_array(m, int32(25))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v34 = l0 + int32(8)
	v36 = v25 << (uint(int32(3)) % 32)
	v37 = v34 + v36
	v39 = v25 << (uint(int32(1)) % 32)
	v40 = base.I32_div_u_s(v39, l1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v40
	v43 = F_palloc(m, v36)
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
	v153 = v29
	goto L1
L7:
	;
	v45 = F_palloc(m, v39)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v47 = int32(0)
	goto L9
L9:
	;
	v62 = v47 << (uint(int32(1)) % 32)
	v68 = v34 + v47<<(uint(int32(3))%32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v69 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v137 = F_construct_md_array(m, v43, v45, l1, v17+int32(8), v17, int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
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
	v85 = v37
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
	v85 = v80 + v37
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v62<<(uint(int32(2))%32)))) = v86
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v45))) = uint8(v89)
	v93 = int32(1)
	v95 = v62 | v93
	v97 = v95 << (uint(int32(2)) % 32)
	v98 = v34 + v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99&int32(1073741824) == v89 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v104 = int32(0)
	if v99 < v104 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v121 = v89
	v123 = v93
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v97))) = v121
	*(*uint8)(unsafe.Add(mBase, uint32(v45+v95))) = uint8(v123)
	v129 = v47 + int32(1)
	if v129 != v25 {
		v47 = v129
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
	v116 = v99 & int32(1073741823)
	v118 = v37
	goto L19
L21:
	;
	goto L22
L22:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v98-int32(4))))
	v113 = v111 & int32(1073741823)
	v116 = v99 - v113
	v118 = v113 + v37
	goto L19
L23:
	;
	v121 = v119
	v123 = v104
	goto L18
L24:
	;
	goto L10
L25:
	;
	v153 = v137
	goto L1
}
func F_hstore_to_matrix(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_hstoreUpgrade(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_hstore_to_array_internal(m, v3, int32(2))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
