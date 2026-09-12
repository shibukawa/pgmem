package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_hstoreUniquePairs(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var __phi35 int32
	_ = __phi35
	var v36 int32
	_ = v36
	var __phi36 int32
	_ = __phi36
	var v38 int32
	_ = v38
	var __phi38 int32
	_ = __phi38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	v4 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v4
	if l1 <= int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v161
L2:
	;
	if l1 != int32(1) {
		v161 = l1
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	F_pg_qsort(m, l0, l1, int32(20), int32(5862))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v18 = v4
	goto L8
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v18 = v17
	goto L8
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v18 + v19
	return int32(1)
L9:
	;
	return int32(0)
L10:
	;
	__phi35 = l0
	__phi36 = l0 + int32(20)
	__phi38 = l0
	v35 = __phi35
	v36 = __phi36
	v38 = __phi38
	goto L11
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if v41 != v42 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+16)))
	if v149 != 0 {
		goto L46
	} else {
		goto L47
	}
L13:
	;
	v140 = int32(20)
	v141 = v36 + v140
	v144 = base.I32_div_s(v141-l0, v140)
	if v144 < l1 {
		__phi35 = v136
		__phi36 = v141
		__phi38 = v36
		v35 = __phi35
		v36 = __phi36
		v38 = __phi38
		goto L11
	} else {
		goto L45
	}
L14:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	if v121 != 0 {
		goto L39
	} else {
		goto L40
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if base.Ui32(int32(4)) <= base.Ui32(v41) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v107 != 0 {
		goto L14
	} else {
		goto L34
	}
L17:
	;
	v107 = int32(0)
	goto L16
L18:
	;
	v81 = v76
	v82 = v77
	v83 = v78
	goto L28
L19:
	;
	if (v44|v45)&int32(3) != 0 {
		v76 = v44
		v77 = v45
		v78 = v41
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v69 = v44
	v70 = v45
	v71 = v41
	goto L21
L21:
	;
	if v71 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L22:
	;
	v53 = v44
	v54 = v45
	v55 = v41
	goto L23
L23:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v58 != v59 {
		v76 = v53
		v77 = v54
		v78 = v55
		goto L18
	} else {
		goto L25
	}
L24:
	;
	v69 = v64
	v70 = v62
	v71 = v66
	goto L21
L25:
	;
	v61 = int32(4)
	v62 = v54 + v61
	v64 = v53 + v61
	v66 = v55 - v61
	if base.Ui32(int32(3)) < base.Ui32(v66) {
		v53 = v64
		v54 = v62
		v55 = v66
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v76 = v69
	v77 = v70
	v78 = v71
	goto L18
L28:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == v87 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v107 = v86 - v87
	goto L16
L30:
	;
	v89 = int32(1)
	v94 = v83 - v89
	if v94 != 0 {
		v81 = v81 + v89
		v82 = v82 + v89
		v83 = v94
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
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+37)))
	if v108 != int32(1) {
		v136 = v35
		goto L13
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v44)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if v113 == int32(0) {
		v136 = v35
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v136 = v35
	goto L13
L39:
	;
	v123 = int32(0)
	goto L41
L40:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v123 = v122
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v120 + (v123 + v42)
	v128 = v35 + int32(20)
	if v35 != v38 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v128)+16)) = v132
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+8)) = v134
	goto L44
L43:
	;
	goto L44
L44:
	;
	v136 = v128
	goto L13
L45:
	;
	goto L12
L46:
	;
	v151 = int32(0)
	goto L48
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v151 = v150
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v148 + (v151 + v146)
	v156 = int32(20)
	v159 = base.I32_div_s(v136-l0+v156, v156)
	v161 = v159
	goto L1
}
func F_hstoreUpgrade(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
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
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v383 int32
	_ = v383
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v468 int32
	_ = v468
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v500 int32
	_ = v500
	v15 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if int32(0) <= v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if l0 == v15 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v500 = v15
	goto L5
L5:
	;
	return v500
L6:
	;
	v23 = F_pg_detoast_datum_copy(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v26 = v19
	v27 = v15
	goto L8
L8:
	;
	v29 = v27 + int32(4)
	if v26 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v26 = v25
	v27 = v23
	goto L8
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v483
	v500 = v27
	goto L5
L11:
	;
	v71 = int32(2)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v74 = v72 & int32(268435455)
	if v74 == int32(0) {
		v202 = v71
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v483 = (v57 + v56) << (uint(int32(2)) % 32)
	goto L10
L13:
	;
	v50 = v41 << (uint(int32(3)) % 32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v29+v50)))
	v56 = v50 + int32(8)
	v57 = v54
	goto L12
L14:
	;
	v56 = int32(8)
	v57 = int32(0)
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(-2147483648)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if base.Ui32(int32(131071)) < base.Ui32(v34) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
	if int32(0) <= v37 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v41 = v26 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v41 | int32(-2147483648)
	if v41 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	if base.Ui32(int32(268435455)) < base.Ui32(v26) {
		goto L57
	} else {
		goto L58
	}
L22:
	;
	v222 = v202
	goto L21
L23:
	;
	if v72 < int32(0) {
		v202 = v71
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v80 = v27 + int32(8)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	if int32(0) <= v81 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v222 = int32(0)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v85 = int32(0)
	v87 = v74 << (uint(int32(3)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v80-int32(4))))
	v96 = v87 + v91&int32(1073741823) + int32(8)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v99 = int32(base.Ui32(v97) >> (uint(int32(2)) % 32))
	if base.Ui32(v99) < base.Ui32(v96) {
		v202 = v85
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v101 = int32(1)
	v105 = v101
	goto L29
L29:
	;
	v117 = v80 + v105<<(uint(int32(2))%32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v118 < int32(0) {
		v202 = v85
		goto L22
	} else {
		goto L31
	}
L30:
	;
	if v74 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v121 = int32(1073741823)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117-int32(4))))
	if base.Ui32(v118&v121) < base.Ui32(v125&v121) {
		v202 = v85
		goto L22
	} else {
		goto L32
	}
L32:
	;
	v130 = v105 + int32(1)
	if v130 != v74<<(uint(v101)%32) {
		v105 = v130
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	v134 = int32(2)
	if base.Ui32(v74) <= base.Ui32(v134) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	if v96 == v99 {
		goto L51
	} else {
		goto L52
	}
L37:
	;
	v137 = v134
	goto L39
L38:
	;
	v137 = v74
	goto L39
L39:
	;
	v142 = int32(1)
	goto L40
L40:
	;
	v151 = v142 << (uint(int32(3)) % 32)
	v152 = v80 + v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v155 = v153 & int32(1073741823)
	if int32(0) <= v153 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L36
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v152-int32(4))))
	v164 = v155 - v160&int32(1073741823)
	goto L44
L43:
	;
	v164 = v155
	goto L44
L44:
	;
	v165 = v27 + v151
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v168 = v166 & int32(1073741823)
	if int32(0) <= v166 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v165-int32(4))))
	v177 = v168 - v173&int32(1073741823)
	goto L47
L46:
	;
	v177 = v168
	goto L47
L47:
	;
	v178 = int32(0)
	if v153&int32(1073741824) != 0 {
		v202 = v178
		goto L22
	} else {
		goto L48
	}
L48:
	;
	if base.Ui32(v164) < base.Ui32(v177) {
		v202 = v178
		goto L22
	} else {
		goto L49
	}
L49:
	;
	v183 = v142 + int32(1)
	if v183 != v137 {
		v142 = v183
		goto L40
	} else {
		goto L50
	}
L50:
	;
	goto L41
L51:
	;
	v199 = int32(2)
	goto L53
L52:
	;
	v199 = int32(1)
	goto L53
L53:
	;
	v202 = v199
	goto L22
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v468
	v483 = v480
	goto L10
L55:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	v450 = int32(-2147483648)
	*(*int32)(unsafe.Add(mBase, uint32(v237))) = v449 | v450
	v456 = v339 << (uint(int32(3)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v456+v237-int32(4))))
	v468 = v339 | v450
	v480 = (v456+v460)<<(uint(int32(2))%32) + int32(32)
	goto L54
L56:
	;
	if v339 != 0 {
		goto L55
	} else {
		goto L101
	}
L57:
	;
	if v222 != 0 {
		goto L92
	} else {
		goto L93
	}
L58:
	;
	v228 = v26<<(uint(int32(3))%32) + int32(8)
	v230 = int32(base.Ui32(v34) >> (uint(int32(2)) % 32))
	if base.Ui32(v230) < base.Ui32(v228) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v232 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if base.Ui32(v232) < base.Ui32(v233) {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v237 = v27 + int32(8)
	if v26 != int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v242 = v232
	goto L64
L62:
	;
	goto L63
L63:
	;
	v278 = int32(1)
	if v26 <= v278 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	v255 = v242 << (uint(int32(3)) % 32)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237+v255))))
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+v255))))
	if base.Ui32(v257) < base.Ui32(v259) {
		goto L57
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	v262 = v242 + int32(1)
	if v262 != v26 {
		v242 = v262
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v281 = v278
	goto L70
L69:
	;
	v281 = v26
	goto L70
L70:
	;
	v282 = int32(0)
	v286 = v282
	v288 = v282
	goto L71
L71:
	;
	v300 = v237 + v288<<(uint(int32(3))%32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if int32(base.Ui32(v301)>>(uint(int32(1))%32)) != v286 {
		goto L57
	} else {
		goto L73
	}
L72:
	;
	if base.Ui32(v230) < base.Ui32(v311+v228) {
		goto L57
	} else {
		goto L78
	}
L73:
	;
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300))))
	if v301&int32(1) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v309 = int32(0)
	goto L76
L75:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+2)))
	v309 = v308
	goto L76
L76:
	;
	v311 = v309 + (v286 + v305)
	v313 = v288 + int32(1)
	if v313 != v281 {
		v286 = v311
		v288 = v313
		goto L71
	} else {
		goto L77
	}
L77:
	;
	goto L72
L78:
	;
	if v222 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v339 <= int32(0) {
		goto L56
	} else {
		goto L85
	}
L80:
	;
	v321 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if v321 == int32(0) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	F_errmsg_internal(m, int32(414823), int32(0))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(476737), int32(312), int32(400106))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L79
L85:
	;
	v345 = int32(0)
	goto L86
L86:
	;
	v359 = v237 + v345<<(uint(int32(3))%32)
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359)+2)))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359))))
	v363 = v359 + int32(4)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)))
	v365 = int32(1)
	v367 = v361 + int32(base.Ui32(v364)>>(uint(v365)%32))
	*(*int32)(unsafe.Add(mBase, uint32(v359))) = v367 & int32(1073741823)
	v373 = v364 & v365
	if v373 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L55
L88:
	;
	v374 = int32(0)
	goto L90
L89:
	;
	v374 = v360
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v363))) = (v374+v367)&int32(1073741823) | v373<<(uint(int32(30))%32)
	v383 = v345 + int32(1)
	if v383 != v339 {
		v345 = v383
		goto L86
	} else {
		goto L91
	}
L91:
	;
	goto L87
L92:
	;
	v400 = v26 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v400 | int32(-2147483648)
	if v400 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v483 = int32(32)
	goto L10
L96:
	;
	goto L97
L97:
	;
	v408 = v400 << (uint(int32(3)) % 32)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v29+v408)))
	v483 = (v408+v410)<<(uint(int32(2))%32) + int32(32)
	goto L10
L98:
	;
	F_errmsg_internal(m, int32(409167), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(476737), int32(274), int32(400106))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	v468 = int32(-2147483648)
	v480 = int32(32)
	goto L54
}
func F_hstore_avals(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	return v119
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
	v38 = v33 + v21<<(uint(int32(3))%32)&int32(2147483640)
	v42 = F_palloc(m, v23<<(uint(int32(2))%32))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v119 = v30
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
	v66 = v33 + v46<<(uint(int32(3))%32) + int32(4)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v67&int32(1073741824) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v106 = F_construct_md_array(m, v42, v44, int32(1), v14+int32(12), v14+int32(8), int32(25), int32(-1), int32(0), int32(105))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L21
	}
L12:
	;
	v87 = int32(1)
	v90 = int32(0)
	goto L14
L13:
	;
	if v67 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42+v46<<(uint(int32(2))%32)))) = v90
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v44))) = uint8(v87)
	v95 = v46 + int32(1)
	if v95 != v23 {
		v46 = v95
		goto L10
	} else {
		goto L20
	}
L15:
	;
	v85 = F_cstring_to_text_with_len(m, v83, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	v82 = v67 & int32(1073741823)
	v83 = v38
	goto L15
L17:
	;
	goto L18
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66-int32(4))))
	v78 = v76 & int32(1073741823)
	v82 = v67 - v78
	v83 = v78 + v38
	goto L15
L19:
	;
	v87 = int32(0)
	v90 = v85
	goto L14
L20:
	;
	goto L11
L21:
	;
	v119 = v106
	goto L3
}
func F_hstore_contains(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
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
	v24 = F_hstoreUpgrade(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v28 = v26 & int32(268435455)
	if v28 == int32(0) {
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
	v33 = int32(8)
	v34 = v19 + v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v36 = int32(3)
	v42 = v24 + v33
	v45 = v42 + v28<<(uint(v36)%32)
	v47 = v35 & int32(268435455)
	v52 = int32(0)
	v58 = int32(0)
	goto L8
L7:
	;
	return int32(0)
L8:
	;
	v71 = v42 + v58<<(uint(int32(3))%32)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v72 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v350
L10:
	;
	if v47 <= v52 {
		goto L7
	} else {
		goto L14
	}
L11:
	;
	v87 = v72 & int32(1073741823)
	v88 = v45
	goto L10
L12:
	;
	goto L13
L13:
	;
	v77 = int32(1073741823)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v71-int32(4))))
	v83 = v81 & v77
	v87 = v72&v77 - v83
	v88 = v83 + v45
	goto L10
L14:
	;
	v92 = v52
	v94 = v47
	goto L16
L15:
	;
	goto L9
L16:
	;
	v111 = base.I32_div_s(v94-v92, int32(2))
	v112 = v111 + v92
	v115 = v34 + v112<<(uint(int32(3))%32)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v118 = v116 & int32(1073741823)
	if int32(0) <= v116 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v341 = int32(1)
	v345 = v58 + v341
	if base.Ui32(v345) < base.Ui32(v28) {
		v52 = v112 + v341
		v58 = v345
		goto L8
	} else {
		goto L91
	}
L18:
	;
	goto L17
L19:
	;
	v332 = int32(0)
	v336 = base.B2i32(v330 < v332)
	if v330 < v332 {
		goto L84
	} else {
		goto L85
	}
L20:
	;
	v138 = v137 + (v34 + v47<<(uint(v36)%32))
	if base.Ui32(int32(4)) <= base.Ui32(v87) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	if base.Ui32(v87) < base.Ui32(v130) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115-int32(4))))
	v125 = v123 & int32(1073741823)
	v126 = v118 - v125
	if v126 != v87 {
		v130 = v126
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v118 == v87 {
		v137 = int32(0)
		goto L20
	} else {
		goto L26
	}
L25:
	;
	v137 = v125
	goto L20
L26:
	;
	v130 = v118
	goto L21
L27:
	;
	v135 = int32(1)
	goto L29
L28:
	;
	v135 = int32(-1)
	goto L29
L29:
	;
	v330 = v135
	goto L19
L30:
	;
	if v200 != 0 {
		v330 = v200
		goto L19
	} else {
		goto L48
	}
L31:
	;
	v200 = int32(0)
	goto L30
L32:
	;
	v174 = v169
	v175 = v170
	v176 = v171
	goto L42
L33:
	;
	if (v138|v88)&int32(3) != 0 {
		v169 = v138
		v170 = v88
		v171 = v87
		goto L32
	} else {
		goto L36
	}
L34:
	;
	v162 = v138
	v163 = v88
	v164 = v87
	goto L35
L35:
	;
	if v164 == int32(0) {
		goto L31
	} else {
		goto L41
	}
L36:
	;
	v146 = v138
	v147 = v88
	v148 = v87
	goto L37
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v151 != v152 {
		v169 = v146
		v170 = v147
		v171 = v148
		goto L32
	} else {
		goto L39
	}
L38:
	;
	v162 = v157
	v163 = v155
	v164 = v159
	goto L35
L39:
	;
	v154 = int32(4)
	v155 = v147 + v154
	v157 = v146 + v154
	v159 = v148 - v154
	if base.Ui32(int32(3)) < base.Ui32(v159) {
		v146 = v157
		v147 = v155
		v148 = v159
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v169 = v162
	v170 = v163
	v171 = v164
	goto L32
L42:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	if v179 == v180 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v200 = v179 - v180
	goto L30
L44:
	;
	v182 = int32(1)
	v187 = v176 - v182
	if v187 != 0 {
		v174 = v174 + v182
		v175 = v175 + v182
		v176 = v187
		goto L42
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	goto L31
L48:
	;
	v201 = int32(0)
	if v112 < v201 {
		v350 = v201
		goto L15
	} else {
		goto L49
	}
L49:
	;
	v208 = v42 + v58<<(uint(int32(1))%32)<<(uint(int32(2))%32) + int32(4)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v211 = v209 & int32(1073741823)
	v213 = v209 & int32(1073741824)
	if int32(0) <= v209 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v208-int32(4))))
	v222 = v211 - v218&int32(1073741823)
	goto L52
L51:
	;
	v222 = v211
	goto L52
L52:
	;
	v224 = v115 + int32(4)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = int32(30)
	if int32(base.Ui32(v225)>>(uint(v226)%32))&int32(1) != int32(base.Ui32(v213)>>(uint(v226)%32)) {
		v350 = v201
		goto L15
	} else {
		goto L53
	}
L53:
	;
	if v213 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	v234 = v225 & int32(1073741823)
	if int32(0) <= v225 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v224-int32(4))))
	v243 = v234 - v239&int32(1073741823)
	goto L57
L56:
	;
	v243 = v234
	goto L57
L57:
	;
	if v243 != v222 {
		v350 = v201
		goto L15
	} else {
		goto L58
	}
L58:
	;
	v245 = int32(0)
	if v245 <= v209 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v208-int32(4))))
	v253 = v250 & int32(1073741823)
	goto L61
L60:
	;
	v253 = v201
	goto L61
L61:
	;
	v254 = v253 + v45
	if int32(0) <= v225 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v224-int32(4))))
	v262 = v259 & int32(1073741823)
	goto L64
L63:
	;
	v262 = v245
	goto L64
L64:
	;
	v263 = v262 + (v34 + v35<<(uint(v36)%32)&int32(2147483640))
	if base.Ui32(int32(4)) <= base.Ui32(v222) {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if v325 == int32(0) {
		goto L18
	} else {
		goto L83
	}
L66:
	;
	v325 = int32(0)
	goto L65
L67:
	;
	v299 = v294
	v300 = v295
	v301 = v296
	goto L77
L68:
	;
	if (v254|v263)&int32(3) != 0 {
		v294 = v254
		v295 = v263
		v296 = v222
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v287 = v254
	v288 = v263
	v289 = v222
	goto L70
L70:
	;
	if v289 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v271 = v254
	v272 = v263
	v273 = v222
	goto L72
L72:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v271)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	if v276 != v277 {
		v294 = v271
		v295 = v272
		v296 = v273
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v287 = v282
	v288 = v280
	v289 = v284
	goto L70
L74:
	;
	v279 = int32(4)
	v280 = v272 + v279
	v282 = v271 + v279
	v284 = v273 - v279
	if base.Ui32(int32(3)) < base.Ui32(v284) {
		v271 = v282
		v272 = v280
		v273 = v284
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v294 = v287
	v295 = v288
	v296 = v289
	goto L67
L77:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300))))
	if v304 == v305 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v325 = v304 - v305
	goto L65
L79:
	;
	v307 = int32(1)
	v312 = v301 - v307
	if v312 != 0 {
		v299 = v299 + v307
		v300 = v300 + v307
		v301 = v312
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
	return int32(0)
L84:
	;
	v337 = v112 + int32(1)
	goto L86
L85:
	;
	v337 = v92
	goto L86
L86:
	;
	if v330 < v332 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v338 = v94
	goto L89
L88:
	;
	v338 = v112
	goto L89
L89:
	;
	if v337 < v338 {
		v92 = v337
		v94 = v338
		goto L16
	} else {
		goto L90
	}
L90:
	;
	v350 = v332
	goto L15
L91:
	;
	v350 = v341
	goto L15
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v62 = F_palloc(m, int32(base.Ui32(v59)>>(uint(int32(2))%32)))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v33 = int32(4)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v35&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v48 = int32(1)
	if v30 != 0 {
		v58 = int32(base.Ui32(v28)>>(uint(v48)%32)) - v48
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v44 = v33
	goto L10
L9:
	;
	v44 = base.B2i32(v35 == int32(18)) << (uint(v33) % 32)
	goto L10
L10:
	;
	if v35 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v33
	goto L13
L12:
	;
	v47 = v44
	goto L13
L13:
	;
	v58 = v47
	goto L4
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v65 & int32(-4)
	v70 = v64 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v70 | int32(-2147483648)
	v75 = v62 + int32(8)
	v78 = v75 + v70<<(uint(int32(3))%32)
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v266&int32(268435455) != v252 {
		goto L63
	} else {
		goto L64
	}
L17:
	;
	v250 = int32(0)
	v252 = v2
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
	v84 = v27
	goto L22
L21:
	;
	v84 = v24 + int32(4)
	goto L22
L22:
	;
	v86 = v19 + int32(8)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v94 = int32(0)
	v97 = v2
	v98 = v75
	v104 = v78
	goto L23
L23:
	;
	v113 = v86 + v94<<(uint(int32(3))%32)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v116 = v114 & int32(1073741823)
	if v114 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v241 = v235 - v78
	if v233 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L25:
	;
	v128 = v127 + (v86 + v87<<(uint(int32(3))%32)&int32(2147483640))
	if v126 == v58 {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v126 = v116
	v127 = int32(0)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113-int32(4))))
	v124 = v122 & int32(1073741823)
	v126 = v116 - v124
	v127 = v124
	goto L25
L29:
	;
	v239 = v94 + int32(1)
	if v239 != v70 {
		v94 = v239
		v97 = v233
		v98 = v234
		v104 = v235
		goto L23
	} else {
		goto L59
	}
L30:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v58) {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	goto L32
L32:
	;
	v198 = v86 + v94<<(uint(int32(3))%32) + int32(4)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v201 = v199 & int32(1073741823)
	if int32(0) <= v199 {
		goto L52
	} else {
		goto L53
	}
L33:
	;
	if v191 == int32(0) {
		v233 = v97
		v234 = v98
		v235 = v104
		goto L29
	} else {
		goto L51
	}
L34:
	;
	v191 = int32(0)
	goto L33
L35:
	;
	v165 = v160
	v166 = v161
	v167 = v162
	goto L45
L36:
	;
	if (v128|v84)&int32(3) != 0 {
		v160 = v128
		v161 = v84
		v162 = v58
		goto L35
	} else {
		goto L39
	}
L37:
	;
	v153 = v128
	v154 = v84
	v155 = v58
	goto L38
L38:
	;
	if v155 == int32(0) {
		goto L34
	} else {
		goto L44
	}
L39:
	;
	v137 = v128
	v138 = v84
	v139 = v58
	goto L40
L40:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v142 != v143 {
		v160 = v137
		v161 = v138
		v162 = v139
		goto L35
	} else {
		goto L42
	}
L41:
	;
	v153 = v148
	v154 = v146
	v155 = v150
	goto L38
L42:
	;
	v145 = int32(4)
	v146 = v138 + v145
	v148 = v137 + v145
	v150 = v139 - v145
	if base.Ui32(int32(3)) < base.Ui32(v150) {
		v137 = v148
		v138 = v146
		v139 = v150
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v160 = v153
	v161 = v154
	v162 = v155
	goto L35
L45:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	if v170 == v171 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v191 = v170 - v171
	goto L33
L47:
	;
	v173 = int32(1)
	v178 = v167 - v173
	if v178 != 0 {
		v165 = v165 + v173
		v166 = v166 + v173
		v167 = v178
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
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v198-int32(4))))
	v210 = v201 - v206&int32(1073741823)
	goto L54
L53:
	;
	v210 = v201
	goto L54
L54:
	;
	v211 = v126 + v210
	if v211 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v214 = v213 + v211
	v215 = v214 - v78
	v217 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = (v215 - v210) & v217
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = v215&v217 | v222&int32(1073741824)
	v233 = v97 + int32(1)
	v234 = v98 + int32(8)
	v235 = v214
	goto L29
L56:
	;
	v212 = F__emscripten_memcpy_bulkmem(m, v104, v128, v211)
	mBase = m.M
	v213 = v212
	goto L58
L57:
	;
	v213 = v104
	goto L58
L58:
	;
	goto L55
L59:
	;
	goto L24
L60:
	;
	v250 = v241
	v252 = int32(0)
	goto L16
L61:
	;
	goto L62
L62:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v245 | int32(-2147483648)
	v250 = v241
	v252 = v233
	goto L16
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v252 | int32(-2147483648)
	v277 = v75 + v252<<(uint(int32(3))%32)&int32(2147483640)
	if v277 == v78 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = (v252<<(uint(int32(3))%32)+v250)<<(uint(int32(2))%32) + int32(32)
	return v62
L66:
	;
	goto L65
L67:
	;
	goto L66
L68:
	;
	v281 = v277 + v250
	if base.Ui32(v78-v281) <= base.Ui32(int32(0)-v250<<(uint(int32(1))%32)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v288 = F___memcpy(m, v277, v78, v250)
	mBase = m.M
	goto L66
L70:
	;
	goto L71
L71:
	;
	v291 = (v277 ^ v78) & int32(3)
	if base.Ui32(v277) < base.Ui32(v78) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	if v393 == int32(0) {
		goto L67
	} else {
		goto L108
	}
L73:
	;
	if base.Ui32(v371) <= base.Ui32(int32(3)) {
		v392 = v370
		v393 = v371
		v394 = v372
		goto L72
	} else {
		goto L104
	}
L74:
	;
	if v291 != 0 {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	if v291 != 0 {
		v353 = v250
		goto L87
	} else {
		goto L88
	}
L77:
	;
	v392 = v78
	v393 = v250
	v394 = v277
	goto L72
L78:
	;
	goto L79
L79:
	;
	if v277&int32(3) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v370 = v78
	v371 = v250
	v372 = v277
	goto L73
L81:
	;
	goto L82
L82:
	;
	v298 = v78
	v299 = v250
	v300 = v277
	goto L83
L83:
	;
	if v299 == int32(0) {
		goto L67
	} else {
		goto L85
	}
L84:
	;
	v370 = v307
	v371 = v309
	v372 = v311
	goto L73
L85:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298))))
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v304)
	v306 = int32(1)
	v307 = v298 + v306
	v309 = v299 - v306
	v311 = v300 + v306
	if v311&int32(3) != 0 {
		v298 = v307
		v299 = v309
		v300 = v311
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	if v353 == int32(0) {
		goto L67
	} else {
		goto L100
	}
L88:
	;
	if v281&int32(3) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v318 = v250
	goto L92
L90:
	;
	v333 = v250
	goto L91
L91:
	;
	if base.Ui32(v333) <= base.Ui32(int32(3)) {
		v353 = v333
		goto L87
	} else {
		goto L96
	}
L92:
	;
	if v318 == int32(0) {
		goto L67
	} else {
		goto L94
	}
L93:
	;
	v333 = v324
	goto L91
L94:
	;
	v324 = v318 - int32(1)
	v325 = v277 + v324
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v324))))
	*(*uint8)(unsafe.Add(mBase, uint32(v325))) = uint8(v327)
	if v325&int32(3) != 0 {
		v318 = v324
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v340 = v333
	goto L97
L97:
	;
	v344 = v340 - int32(4)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v78+v344)))
	*(*int32)(unsafe.Add(mBase, uint32(v277+v344))) = v347
	if base.Ui32(int32(3)) < base.Ui32(v344) {
		v340 = v344
		goto L97
	} else {
		goto L99
	}
L98:
	;
	v353 = v344
	goto L87
L99:
	;
	goto L98
L100:
	;
	v360 = v353
	goto L101
L101:
	;
	v364 = v360 - int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78+v364))))
	*(*uint8)(unsafe.Add(mBase, uint32(v277+v364))) = uint8(v367)
	if v364 != 0 {
		v360 = v364
		goto L101
	} else {
		goto L103
	}
L102:
	;
	goto L67
L103:
	;
	goto L102
L104:
	;
	v377 = v370
	v378 = v371
	v379 = v372
	goto L105
L105:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	*(*int32)(unsafe.Add(mBase, uint32(v379))) = v381
	v383 = int32(4)
	v384 = v377 + v383
	v386 = v379 + v383
	v388 = v378 - v383
	if base.Ui32(int32(3)) < base.Ui32(v388) {
		v377 = v384
		v378 = v388
		v379 = v386
		goto L105
	} else {
		goto L107
	}
L106:
	;
	v392 = v384
	v393 = v388
	v394 = v386
	goto L72
L107:
	;
	goto L106
L108:
	;
	v399 = v392
	v400 = v393
	v401 = v394
	goto L109
L109:
	;
	v403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v403)
	v405 = int32(1)
	v410 = v400 - v405
	if v410 != 0 {
		v399 = v399 + v405
		v400 = v410
		v401 = v401 + v405
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L67
L111:
	;
	goto L110
}
func F_hstore_each(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		v17 = int32(4455216)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v19 = F_hstoreUpgrade(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v31 = F_palloc(m, int32(base.Ui32(v28)>>(uint(int32(2))%32)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v35 = int32(base.Ui32(v33) >> (uint(int32(2)) % 32))
					if v35 != 0 {
						v36 = F__emscripten_memcpy_bulkmem(m, v31, v19, v35)
						mBase = m.M
						v37 = v36
					} else {
						v37 = v31
					}
					*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v37
					v42 = F_get_call_result_type(m, l0, int32(0), v11+int32(12))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 != int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(354556), int32(0))
								mBase = m.M
								v171 = m.ExcPending
								if v171 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(479033), int32(869), int32(293927))
									mBase = m.M
									v178 = m.ExcPending
									if v178 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v47 = F_BlessTupleDesc(m, v46)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v47
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
								v61 = v59 & int32(268435455)
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
								if base.Ui32(v62) < base.Ui32(v61) {
									v64 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v64)
									v67 = v58 + int32(8)
									v68 = int32(3)
									v70 = v67 + v61<<(uint(v68)%32)
									v75 = v67 + v62<<(uint(v68)%32)
									v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
									if v76 < v64 {
										v90 = v76 & int32(1073741823)
										v92 = v70
									} else {
										v81 = int32(1073741823)
										v85 = *(*int32)(unsafe.Add(mBase, uint32(v75-int32(4))))
										v87 = v85 & v81
										v90 = v76&v81 - v87
										v92 = v70 + v87
									}
									v93 = F_cstring_to_text_with_len(m, v92, v90)
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v93
										v100 = v67 + v62<<(uint(int32(1))%32)<<(uint(int32(2))%32) + int32(4)
										v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
										if v101&int32(1073741824) != 0 {
											v104 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v104)
											*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
											v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
											v133 = F_heap_form_tuple(m, v128, v11+int32(4), v11+int32(2))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int32(0)
											} else {
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
												v136 = F_HeapTupleHeaderGetDatum(m, v135)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													v138 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
													*(*int64)(unsafe.Add(mBase, uint32(v57))) = v138 + int64(1)
													v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1)
													v153 = v136
													m.G0 = v11 + int32(16)
													return v153
												}
											}
										} else {
											if v101 < int32(0) {
												v119 = v101 & int32(1073741823)
												v120 = v70
											} else {
												v114 = *(*int32)(unsafe.Add(mBase, uint32(v100-int32(4))))
												v116 = v114 & int32(1073741823)
												v119 = v101 - v116
												v120 = v70 + v116
											}
											v122 = F_cstring_to_text_with_len(m, v120, v119)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v122
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
												v133 = F_heap_form_tuple(m, v128, v11+int32(4), v11+int32(2))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
													v136 = F_HeapTupleHeaderGetDatum(m, v135)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														v138 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
														*(*int64)(unsafe.Add(mBase, uint32(v57))) = v138 + int64(1)
														v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1)
														v153 = v136
														m.G0 = v11 + int32(16)
														return v153
													}
												}
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return int32(0)
									} else {
										v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = int32(2)
										v150 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v150)
										v153 = int32(0)
										m.G0 = v11 + int32(16)
										return v153
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
			*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)) = uint16(v64)
			v67 = v58 + int32(8)
			v68 = int32(3)
			v70 = v67 + v61<<(uint(v68)%32)
			v75 = v67 + v62<<(uint(v68)%32)
			v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
			if v76 < v64 {
				v90 = v76 & int32(1073741823)
				v92 = v70
			} else {
				v81 = int32(1073741823)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v75-int32(4))))
				v87 = v85 & v81
				v90 = v76&v81 - v87
				v92 = v70 + v87
			}
			v93 = F_cstring_to_text_with_len(m, v92, v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v93
				v100 = v67 + v62<<(uint(int32(1))%32)<<(uint(int32(2))%32) + int32(4)
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
				if v101&int32(1073741824) != 0 {
					v104 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+3)) = uint8(v104)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(0)
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
					v133 = F_heap_form_tuple(m, v128, v11+int32(4), v11+int32(2))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return int32(0)
					} else {
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
						v136 = F_HeapTupleHeaderGetDatum(m, v135)
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int32(0)
						} else {
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = v138 + int64(1)
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1)
							v153 = v136
							m.G0 = v11 + int32(16)
							return v153
						}
					}
				} else {
					if v101 < int32(0) {
						v119 = v101 & int32(1073741823)
						v120 = v70
					} else {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v100-int32(4))))
						v116 = v114 & int32(1073741823)
						v119 = v101 - v116
						v120 = v70 + v116
					}
					v122 = F_cstring_to_text_with_len(m, v120, v119)
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v122
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v57)+28))
						v133 = F_heap_form_tuple(m, v128, v11+int32(4), v11+int32(2))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
							v136 = F_HeapTupleHeaderGetDatum(m, v135)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int32(0)
							} else {
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = v138 + int64(1)
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v142)+20)) = int32(1)
								v153 = v136
								m.G0 = v11 + int32(16)
								return v153
							}
						}
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return int32(0)
			} else {
				v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = int32(2)
				v150 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v150)
				v153 = int32(0)
				m.G0 = v11 + int32(16)
				return v153
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
	var v50 int32
	_ = v50
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
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
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v235 int32
	_ = v235
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
	return v235
L6:
	;
	v36 = v21 + int32(8)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v39 = v37 & int32(268435455)
	v44 = int32(0)
	v50 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v235 = int32(0)
	goto L5
L9:
	;
	if v39 <= v44 {
		v186 = v44
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v202 = v50 + int32(1)
	if v202 != v32 {
		v44 = v186
		v50 = v202
		goto L9
	} else {
		goto L53
	}
L12:
	;
	v62 = v30 + v50*int32(20)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v65 = v44
	v72 = v39
	goto L13
L13:
	;
	v82 = base.I32_div_s(v72-v65, int32(2))
	v83 = v82 + v65
	v86 = v36 + v83<<(uint(int32(3))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v89 = v87 & int32(1073741823)
	if int32(0) <= v87 {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v186 = v183
	goto L11
L15:
	;
	v182 = base.B2i32(v177 < int32(0))
	if v177 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	v109 = v108 + (v36 + v39<<(uint(int32(3))%32))
	if base.Ui32(int32(4)) <= base.Ui32(v63) {
		goto L29
	} else {
		goto L30
	}
L17:
	;
	if base.Ui32(v63) < base.Ui32(v101) {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v86-int32(4))))
	v96 = v94 & int32(1073741823)
	v97 = v89 - v96
	if v97 != v63 {
		v101 = v97
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v89 == v63 {
		v108 = int32(0)
		goto L16
	} else {
		goto L22
	}
L21:
	;
	v108 = v96
	goto L16
L22:
	;
	v101 = v89
	goto L17
L23:
	;
	v106 = int32(1)
	goto L25
L24:
	;
	v106 = int32(-1)
	goto L25
L25:
	;
	v177 = v106
	goto L15
L26:
	;
	if v171 != 0 {
		v177 = v171
		goto L15
	} else {
		goto L44
	}
L27:
	;
	v171 = int32(0)
	goto L26
L28:
	;
	v145 = v140
	v146 = v141
	v147 = v142
	goto L38
L29:
	;
	if (v109|v64)&int32(3) != 0 {
		v140 = v109
		v141 = v64
		v142 = v63
		goto L28
	} else {
		goto L32
	}
L30:
	;
	v133 = v109
	v134 = v64
	v135 = v63
	goto L31
L31:
	;
	if v135 == int32(0) {
		goto L27
	} else {
		goto L37
	}
L32:
	;
	v117 = v109
	v118 = v64
	v119 = v63
	goto L33
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	if v122 != v123 {
		v140 = v117
		v141 = v118
		v142 = v119
		goto L28
	} else {
		goto L35
	}
L34:
	;
	v133 = v128
	v134 = v126
	v135 = v130
	goto L31
L35:
	;
	v125 = int32(4)
	v126 = v118 + v125
	v128 = v117 + v125
	v130 = v119 - v125
	if base.Ui32(int32(3)) < base.Ui32(v130) {
		v117 = v128
		v118 = v126
		v119 = v130
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v140 = v133
	v141 = v134
	v142 = v135
	goto L28
L38:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146))))
	if v150 == v151 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v171 = v150 - v151
	goto L26
L40:
	;
	v153 = int32(1)
	v158 = v147 - v153
	if v158 != 0 {
		v145 = v145 + v153
		v146 = v146 + v153
		v147 = v158
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	goto L27
L44:
	;
	if int32(0) <= v83 {
		v235 = int32(1)
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v186 = v83 + int32(1)
	goto L11
L46:
	;
	v183 = v83 + int32(1)
	goto L48
L47:
	;
	v183 = v65
	goto L48
L48:
	;
	if v177 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v184 = v72
	goto L51
L50:
	;
	v184 = v83
	goto L51
L51:
	;
	if v183 < v184 {
		v65 = v183
		v72 = v184
		goto L13
	} else {
		goto L52
	}
L52:
	;
	goto L14
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5466), int32(0), v4, v5)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(5466), int32(0), v4, v5)
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
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
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v606 int32
	_ = v606
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v649 int32
	_ = v649
	v4 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(32)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v31 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = F_pg_detoast_datum_packed(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	v633 = m.ExcPending
	if v633 != 0 {
		goto L4
	} else {
		goto L160
	}
L4:
	;
	return
L5:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+29)) = uint8(v38)
	v40 = int32(1)
	v41 = v36 + v40
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v44&v40 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v47 = v41
	goto L8
L7:
	;
	v47 = v36 + int32(4)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v47
	if v44 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v79 = F_hstoreCheckKeyLen(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L20
	}
L10:
	;
	v51 = int32(4)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v53&int32(254) == int32(2) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v66 = int32(1)
	if v44&v66 != 0 {
		v78 = int32(base.Ui32(v44)>>(uint(v66)%32)) - v66
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v62 = v51
	goto L15
L14:
	;
	v62 = base.B2i32(v53 == int32(18)) << (uint(v51) % 32)
	goto L15
L15:
	;
	if v53 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v65 = v51
	goto L18
L17:
	;
	v65 = v62
	goto L18
L18:
	;
	v78 = v65
	goto L9
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v78 = int32(base.Ui32(v72)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+44)))
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+28)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v133
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v137 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L22:
	;
	v128 = int32(1)
	v132 = v4
	v133 = int32(0)
	goto L21
L23:
	;
	goto L24
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v86 = F_pg_detoast_datum_packed(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v88 = int32(1)
	v89 = v86 + v88
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v94 = v92 & v88
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v95 = v89
	goto L28
L27:
	;
	v95 = v86 + int32(4)
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v95
	if v92 == int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v126 = F_hstoreCheckValLen(m, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L40
	}
L30:
	;
	v99 = int32(4)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89))))
	if v101&int32(254) == int32(2) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v114 = int32(1)
	if v94 != 0 {
		v124 = int32(base.Ui32(v92)>>(uint(v114)%32)) - v114
		goto L29
	} else {
		goto L39
	}
L33:
	;
	v110 = v99
	goto L35
L34:
	;
	v110 = base.B2i32(v101 == int32(18)) << (uint(v99) % 32)
	goto L35
L35:
	;
	if v101 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v113 = v99
	goto L38
L37:
	;
	v113 = v110
	goto L38
L38:
	;
	v124 = v113
	goto L29
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v124 = int32(base.Ui32(v118)>>(uint(int32(2))%32)) - int32(4)
	goto L29
L40:
	;
	v128 = int32(0)
	v132 = v95
	v133 = v126
	goto L21
L41:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v622))) = v606
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v625 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v624))) = uint8(v625)
	m.G0 = v27 + int32(32)
	return
L42:
	;
	v144 = F_hstorePairs(m, v27+int32(12), int32(1), v79+v133)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = F_hstoreUpgrade(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	v606 = v144
	goto L41
L46:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v156 = v154 & int32(268435455)
	v159 = int32(8)
	v160 = v156<<(uint(int32(3))%32) + v159
	v164 = v79 + int32(base.Ui32(v150)>>(uint(int32(2))%32)) + v160 + v133 + v159
	v165 = F_palloc(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v156 - int32(2147483647)
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v164 << (uint(int32(2)) % 32)
	v173 = int32(8)
	v174 = v148 + v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v178 = int32(2147483640)
	v180 = v174 + v175<<(uint(int32(3))%32)&v178
	v181 = int32(0)
	v184 = v165 + v173
	v187 = v184 + v160&v178
	v190 = v184
	v193 = base.B2i32(v156 != v181)
	v194 = v187
	v195 = v181
	v200 = v181
	v202 = v4
	goto L48
L48:
	;
	v214 = int32(1)
	if v193&v214 == int32(0) {
		v321 = v214
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v429 | int32(-2147483648)
	v433 = v414 - v187
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v434&int32(268435455) != v422 {
		goto L111
	} else {
		goto L112
	}
L50:
	;
	v422 = v202 + int32(1)
	v424 = v190 + int32(8)
	v425 = v420 + v195
	v426 = base.B2i32(base.Ui32(v425) < base.Ui32(v156))
	if base.Ui32(v425) < base.Ui32(v156) {
		v190 = v424
		v193 = v426
		v194 = v414
		v195 = v425
		v200 = v417
		v202 = v422
		goto L48
	} else {
		goto L109
	}
L51:
	;
	v361 = v174 + v195<<(uint(int32(3))%32) + int32(4)
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v364 = v362 & int32(1073741823)
	if int32(0) <= v362 {
		goto L96
	} else {
		goto L97
	}
L52:
	;
	v345 = int32(1073741823)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v314-int32(4))))
	v351 = v349 & v345
	v355 = v315&v345 - v351
	v356 = v351 + v180
	goto L51
L53:
	;
	if v79 != 0 {
		goto L86
	} else {
		goto L87
	}
L54:
	;
	if int32(0) < v200 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v314 = v174 + v195<<(uint(int32(3))%32)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	if int32(0) <= v315 {
		goto L52
	} else {
		goto L84
	}
L56:
	;
	v223 = v174 + v195<<(uint(int32(3))%32)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v226 = v224 & int32(1073741823)
	if int32(0) <= v224 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v244 = v243 + v180
	if base.Ui32(int32(4)) <= base.Ui32(v79) {
		goto L68
	} else {
		goto L69
	}
L58:
	;
	if v238 <= v79 {
		goto L55
	} else {
		goto L64
	}
L59:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v223-int32(4))))
	v233 = v231 & int32(1073741823)
	v234 = v226 - v233
	if v234 != v79 {
		v238 = v234
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	if v226 == v79 {
		v243 = int32(0)
		goto L57
	} else {
		goto L63
	}
L62:
	;
	v243 = v233
	goto L57
L63:
	;
	v238 = v226
	goto L58
L64:
	;
	v321 = int32(1)
	goto L53
L65:
	;
	if int32(0) <= v306 {
		v321 = v306
		goto L53
	} else {
		goto L83
	}
L66:
	;
	v306 = int32(0)
	goto L65
L67:
	;
	v280 = v275
	v281 = v276
	v282 = v277
	goto L77
L68:
	;
	if (v244|v47)&int32(3) != 0 {
		v275 = v244
		v276 = v47
		v277 = v79
		goto L67
	} else {
		goto L71
	}
L69:
	;
	v268 = v244
	v269 = v47
	v270 = v79
	goto L70
L70:
	;
	if v270 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v252 = v244
	v253 = v47
	v254 = v79
	goto L72
L72:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v252)))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	if v257 != v258 {
		v275 = v252
		v276 = v253
		v277 = v254
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v268 = v263
	v269 = v261
	v270 = v265
	goto L70
L74:
	;
	v260 = int32(4)
	v261 = v253 + v260
	v263 = v252 + v260
	v265 = v254 - v260
	if base.Ui32(int32(3)) < base.Ui32(v265) {
		v252 = v263
		v253 = v261
		v254 = v265
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v275 = v268
	v276 = v269
	v277 = v270
	goto L67
L77:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v285 == v286 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v306 = v285 - v286
	goto L65
L79:
	;
	v288 = int32(1)
	v293 = v282 - v288
	if v293 != 0 {
		v280 = v280 + v288
		v281 = v281 + v288
		v282 = v293
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
	v355 = v315 & int32(1073741823)
	v356 = v180
	goto L51
L85:
	;
	v325 = v324 + v79
	v328 = (v325 - v187) & int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v328
	if v82 != 0 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v323 = F__emscripten_memcpy_bulkmem(m, v194, v47, v79)
	mBase = m.M
	v324 = v323
	goto L88
L87:
	;
	v324 = v194
	goto L88
L88:
	;
	goto L85
L89:
	;
	v338 = v325
	v339 = v328 | int32(1073741824)
	goto L91
L90:
	;
	if v133 != 0 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v339
	v414 = v338
	v417 = v200 + int32(1)
	v420 = base.B2i32(v321 == int32(0))
	goto L50
L92:
	;
	v334 = v333 + v133
	v338 = v334
	v339 = (v334 - v187) & int32(1073741823)
	goto L91
L93:
	;
	v332 = F__emscripten_memcpy_bulkmem(m, v325, v132, v133)
	mBase = m.M
	v333 = v332
	goto L95
L94:
	;
	v333 = v325
	goto L95
L95:
	;
	goto L92
L96:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v361-int32(4))))
	v373 = v364 - v369&int32(1073741823)
	goto L98
L97:
	;
	v373 = v364
	goto L98
L98:
	;
	v374 = v373 + v355
	if v374 != 0 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v379 = v377 & int32(1073741823)
	if int32(0) <= v377 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v375 = F__emscripten_memcpy_bulkmem(m, v194, v356, v374)
	mBase = m.M
	v376 = v375
	goto L102
L101:
	;
	v376 = v194
	goto L102
L102:
	;
	goto L99
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v314-int32(4))))
	v388 = v379 - v384&int32(1073741823)
	goto L105
L104:
	;
	v388 = v379
	goto L105
L105:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v392 = int32(0)
	if v392 <= v389 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v395 = v379
	goto L108
L107:
	;
	v395 = v392
	goto L108
L108:
	;
	v396 = v389&int32(1073741823) - v395
	v398 = v396 + (v388 + v376)
	v399 = v398 - v187
	v401 = int32(1073741823)
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = (v399 - v396) & v401
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v404&int32(1073741824) | v399&v401
	v414 = v398
	v417 = v200
	v420 = int32(1)
	goto L50
L109:
	;
	if v417 <= int32(0) {
		v190 = v424
		v193 = v426
		v194 = v414
		v195 = v425
		v200 = v417
		v202 = v422
		goto L48
	} else {
		goto L110
	}
L110:
	;
	goto L49
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v422 | int32(-2147483648)
	v445 = v184 + v422<<(uint(int32(3))%32)&int32(2147483640)
	if v445 == v187 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = (v422<<(uint(int32(3))%32)+v433)<<(uint(int32(2))%32) + int32(32)
	v606 = v165
	goto L41
L114:
	;
	goto L113
L115:
	;
	goto L114
L116:
	;
	v449 = v445 + v433
	if base.Ui32(v187-v449) <= base.Ui32(int32(0)-v433<<(uint(int32(1))%32)) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v456 = F___memcpy(m, v445, v187, v433)
	mBase = m.M
	goto L114
L118:
	;
	goto L119
L119:
	;
	v459 = (v445 ^ v187) & int32(3)
	if base.Ui32(v445) < base.Ui32(v187) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	if v561 == int32(0) {
		goto L115
	} else {
		goto L156
	}
L121:
	;
	if base.Ui32(v539) <= base.Ui32(int32(3)) {
		v560 = v538
		v561 = v539
		v562 = v540
		goto L120
	} else {
		goto L152
	}
L122:
	;
	if v459 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v459 != 0 {
		v521 = v433
		goto L135
	} else {
		goto L136
	}
L125:
	;
	v560 = v187
	v561 = v433
	v562 = v445
	goto L120
L126:
	;
	goto L127
L127:
	;
	if v445&int32(3) == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v538 = v187
	v539 = v433
	v540 = v445
	goto L121
L129:
	;
	goto L130
L130:
	;
	v466 = v187
	v467 = v433
	v468 = v445
	goto L131
L131:
	;
	if v467 == int32(0) {
		goto L115
	} else {
		goto L133
	}
L132:
	;
	v538 = v475
	v539 = v477
	v540 = v479
	goto L121
L133:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	*(*uint8)(unsafe.Add(mBase, uint32(v468))) = uint8(v472)
	v474 = int32(1)
	v475 = v466 + v474
	v477 = v467 - v474
	v479 = v468 + v474
	if v479&int32(3) != 0 {
		v466 = v475
		v467 = v477
		v468 = v479
		goto L131
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	if v521 == int32(0) {
		goto L115
	} else {
		goto L148
	}
L136:
	;
	if v449&int32(3) != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v486 = v433
	goto L140
L138:
	;
	v501 = v433
	goto L139
L139:
	;
	if base.Ui32(v501) <= base.Ui32(int32(3)) {
		v521 = v501
		goto L135
	} else {
		goto L144
	}
L140:
	;
	if v486 == int32(0) {
		goto L115
	} else {
		goto L142
	}
L141:
	;
	v501 = v492
	goto L139
L142:
	;
	v492 = v486 - int32(1)
	v493 = v445 + v492
	v495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v492))))
	*(*uint8)(unsafe.Add(mBase, uint32(v493))) = uint8(v495)
	if v493&int32(3) != 0 {
		v486 = v492
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v508 = v501
	goto L145
L145:
	;
	v512 = v508 - int32(4)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v187+v512)))
	*(*int32)(unsafe.Add(mBase, uint32(v445+v512))) = v515
	if base.Ui32(int32(3)) < base.Ui32(v512) {
		v508 = v512
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v521 = v512
	goto L135
L147:
	;
	goto L146
L148:
	;
	v528 = v521
	goto L149
L149:
	;
	v532 = v528 - int32(1)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v532))))
	*(*uint8)(unsafe.Add(mBase, uint32(v445+v532))) = uint8(v535)
	if v532 != 0 {
		v528 = v532
		goto L149
	} else {
		goto L151
	}
L150:
	;
	goto L115
L151:
	;
	goto L150
L152:
	;
	v545 = v538
	v546 = v539
	v547 = v540
	goto L153
L153:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	*(*int32)(unsafe.Add(mBase, uint32(v547))) = v549
	v551 = int32(4)
	v552 = v545 + v551
	v554 = v547 + v551
	v556 = v546 - v551
	if base.Ui32(int32(3)) < base.Ui32(v556) {
		v545 = v552
		v546 = v556
		v547 = v554
		goto L153
	} else {
		goto L155
	}
L154:
	;
	v560 = v552
	v561 = v556
	v562 = v554
	goto L120
L155:
	;
	goto L154
L156:
	;
	v567 = v560
	v568 = v561
	v569 = v562
	goto L157
L157:
	;
	v571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567))))
	*(*uint8)(unsafe.Add(mBase, uint32(v569))) = uint8(v571)
	v573 = int32(1)
	v578 = v568 - v573
	if v578 != 0 {
		v567 = v567 + v573
		v568 = v578
		v569 = v569 + v573
		goto L157
	} else {
		goto L159
	}
L158:
	;
	goto L115
L159:
	;
	goto L158
L160:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(292057), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(478244), int32(156), int32(269811))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	v63 = v34 + v47<<(uint(int32(3))%32)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v64 < int32(0) {
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
	v82 = v47 << (uint(int32(1)) % 32)
	v86 = F_cstring_to_text_with_len(m, v80, v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L15
	}
L12:
	;
	v79 = v64 & int32(1073741823)
	v80 = v37
	goto L11
L13:
	;
	goto L14
L14:
	;
	v69 = int32(1073741823)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63-int32(4))))
	v75 = v73 & v69
	v79 = v64&v69 - v75
	v80 = v75 + v37
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v82<<(uint(int32(2))%32)))) = v86
	v89 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v82+v45))) = uint8(v89)
	v93 = int32(1)
	v95 = v82 | v93
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
	if v99 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v121 = v89
	v122 = v93
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43+v97))) = v121
	*(*uint8)(unsafe.Add(mBase, uint32(v95+v45))) = uint8(v122)
	v129 = v47 + int32(1)
	if v129 != v25 {
		v47 = v129
		goto L9
	} else {
		goto L24
	}
L19:
	;
	v119 = F_cstring_to_text_with_len(m, v117, v115)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L23
	}
L20:
	;
	v115 = v99 & int32(1073741823)
	v117 = v37
	goto L19
L21:
	;
	goto L22
L22:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98-int32(4))))
	v112 = v110 & int32(1073741823)
	v115 = v99 - v112
	v117 = v112 + v37
	goto L19
L23:
	;
	v121 = v119
	v122 = int32(0)
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
