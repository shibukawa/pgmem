package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpOffsetNumbers(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	return v3 - v4
}
func F_cmpTheLexeme(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v73
L2:
	;
	if v6 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v40 = base.B2i32(v6 != int32(0))
	goto L4
L4:
	;
	if v40 != 0 {
		v73 = v40
		goto L1
	} else {
		goto L15
	}
L5:
	;
	return int32(-1)
L6:
	;
	goto L7
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v14 == int32(0))|base.B2i32(v14 != v17) != 0 {
		v35 = v14
		v36 = v17
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v40 = v35 - v36
	goto L4
L9:
	;
	goto L8
L10:
	;
	v20 = v7
	v21 = v6
	goto L11
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v25
		v36 = v24
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v35 = v25
	v36 = v24
	goto L9
L13:
	;
	v28 = int32(1)
	if v25 == v24 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v41 = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 == v41 {
		v73 = v41
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v45 == int32(0) {
		v73 = v41
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v48 == v49 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+4)))
	if v51 == v52 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if base.Ui32(v49) < base.Ui32(v48) {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+6)))
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v54 == v55 {
		v73 = v41
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if base.Ui32(v52) < base.Ui32(v51) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if base.Ui32(v55) < base.Ui32(v54) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v60 = int32(-1)
	goto L27
L26:
	;
	v60 = int32(1)
	goto L27
L27:
	;
	return v60
L28:
	;
	v65 = int32(-1)
	goto L30
L29:
	;
	v65 = int32(1)
	goto L30
L30:
	;
	return v65
L31:
	;
	v70 = int32(-1)
	goto L33
L32:
	;
	v70 = int32(1)
	goto L33
L33:
	;
	v73 = v70
	goto L1
}
func F_cmp_abs_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v172 int32
	_ = v172
	v7 = int32(0)
	if base.B2i32(l5 < l2)&base.B2i32(v7 < l1) == v7 {
		v38 = l2
		v42 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v172
L2:
	;
	if base.B2i32(l4 <= int32(0))|base.B2i32(l5 <= v38) != 0 {
		v74 = l5
		v76 = v7
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v19 = l2
	v23 = v7
	goto L4
L4:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v23<<(uint(int32(1))%32)))))
	if v29 != 0 {
		v172 = int32(1)
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v38 = v33
	v42 = v31
	goto L2
L6:
	;
	v30 = int32(1)
	v31 = v23 + v30
	v33 = v19 - v30
	if v33 <= l5 {
		v38 = v33
		v42 = v31
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v31 < l1 {
		v19 = v33
		v23 = v31
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	if v38 != v74 {
		v117 = v42
		v118 = v76
		goto L16
	} else {
		goto L17
	}
L10:
	;
	v55 = l5
	v57 = v7
	goto L11
L11:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v57<<(uint(int32(1))%32)))))
	if v62 != 0 {
		v172 = int32(-1)
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v74 = v66
	v76 = v64
	goto L9
L13:
	;
	v63 = int32(1)
	v64 = v57 + v63
	v66 = v55 - v63
	if v66 <= v38 {
		v74 = v66
		v76 = v64
		goto L9
	} else {
		goto L14
	}
L14:
	;
	if v64 < l4 {
		v55 = v66
		v57 = v64
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	if l1 < v117 {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	v85 = v42
	v86 = v76
	goto L18
L18:
	;
	if base.B2i32(l1 <= v85)|base.B2i32(l4 <= v86) != 0 {
		v117 = v85
		v118 = v86
		goto L16
	} else {
		goto L20
	}
L19:
	;
	if base.I32_extend16_s(v102) < base.I32_extend16_s(v100) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v91 = int32(1)
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v85<<(uint(v91)%32)))))
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86<<(uint(v91)%32)+l3))))
	if v100 == v102 {
		v85 = v85 + v91
		v86 = v86 + v91
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
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
	return v109
L25:
	;
	v121 = v117
	goto L27
L26:
	;
	v121 = l1
	goto L27
L27:
	;
	v128 = v117
	goto L28
L28:
	;
	if v121 == v128 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v172 = v155
	goto L1
L30:
	;
	if l4 < v118 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v155 = int32(1)
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v128<<(uint(v155)%32)))))
	if v161 == int32(0) {
		v128 = v128 + v155
		goto L28
	} else {
		goto L42
	}
L33:
	;
	v133 = v118
	goto L35
L34:
	;
	v133 = l4
	goto L35
L35:
	;
	v141 = v118
	goto L36
L36:
	;
	if v133 == v141 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v172 = int32(-1)
	goto L1
L38:
	;
	return int32(0)
L39:
	;
	goto L40
L40:
	;
	v146 = int32(1)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141<<(uint(v146)%32)+l3))))
	if v151 == int32(0) {
		v141 = v141 + v146
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	goto L29
}
func F_cmp_numerics(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = base.I32_extend16_s(v14)
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(int32(_a_F_cmp_numerics_0)) <= base.Ui32(v16) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v16 != int32(_a_F_cmp_numerics_1) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(int32(-16384)) <= base.Ui32(v15) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	if v15 != int32(-4096) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	if v16 != int32(_a_F_cmp_numerics_0) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v15 == int32(-16384) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return base.B2i32(v15 != int32(-16384))
L9:
	;
	v31 = int32(-1)
	goto L11
L10:
	;
	v31 = base.B2i32(v15 != int32(-12288))
	goto L11
L11:
	;
	return v31
L12:
	;
	v37 = int32(-1)
	goto L14
L13:
	;
	v37 = int32(0)
	goto L14
L14:
	;
	return v37
L15:
	;
	if v15 == int32(-4096) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v48 = l0 + int32(6)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v54 = base.I32_extend16_s(v16)
	v56 = base.B2i32(int32(0) <= v54)
	if int32(0) <= v54 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v45 = int32(1)
	goto L20
L19:
	;
	v45 = int32(-1)
	goto L20
L20:
	;
	return v45
L21:
	;
	v57 = int32(-8)
	goto L23
L22:
	;
	v57 = int32(-6)
	goto L23
L23:
	;
	if int32(0) <= v54 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48))))
	v69 = v59
	goto L26
L25:
	;
	v69 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
	goto L26
L26:
	;
	v71 = int32(base.Ui32(int32(base.Ui32(v49)>>(uint(int32(2))%32))+v57) >> (uint(int32(1)) % 32))
	v73 = l1 + int32(6)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v80 = base.B2i32(int32(0) <= v15)
	if int32(0) <= v15 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = int32(-8)
	goto L29
L28:
	;
	v81 = int32(-6)
	goto L29
L29:
	;
	if int32(0) <= v15 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73))))
	v93 = v83
	goto L32
L31:
	;
	v93 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
	goto L32
L32:
	;
	v94 = int32(1)
	v95 = int32(base.Ui32(int32(base.Ui32(v74)>>(uint(int32(2))%32))+v81) >> (uint(v94) % 32))
	v101 = v14 & int32(_a_F_cmp_numerics_0)
	if v101 == int32(_a_F_cmp_numerics_2) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v104 = v14 << (uint(v94) % 32) & int32(_a_F_cmp_numerics_3)
	goto L35
L34:
	;
	v104 = v101
	goto L35
L35:
	;
	if v71 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v95 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v122 = v16 & int32(_a_F_cmp_numerics_0)
	if v122 == int32(_a_F_cmp_numerics_2) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	return int32(0)
L40:
	;
	goto L41
L41:
	;
	if v104 == int32(_a_F_cmp_numerics_3) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v115 = int32(1)
	goto L44
L43:
	;
	v115 = int32(-1)
	goto L44
L44:
	;
	return v115
L45:
	;
	v125 = v16 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_3)
	goto L47
L46:
	;
	v125 = v122
	goto L47
L47:
	;
	if v95 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v125 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if v54 < int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v130 = int32(-1)
	goto L53
L52:
	;
	v130 = int32(1)
	goto L53
L53:
	;
	return v130
L54:
	;
	v136 = v48
	goto L56
L55:
	;
	v136 = l0 + int32(8)
	goto L56
L56:
	;
	if v15 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v141 = v73
	goto L59
L58:
	;
	v141 = l1 + int32(8)
	goto L59
L59:
	;
	if v125 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v104 == int32(_a_F_cmp_numerics_3) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v104 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L63:
	;
	return int32(1)
L64:
	;
	goto L65
L65:
	;
	v148 = int32(0)
	if base.B2i32(v93 < v69)&base.B2i32(v148 < v71) == v148 {
		v179 = v69
		v183 = v148
		goto L68
	} else {
		goto L69
	}
L66:
	;
	return v321
L67:
	;
	v321 = v311
	goto L66
L68:
	;
	if base.B2i32(v95 <= int32(0))|base.B2i32(v93 <= v179) != 0 {
		v215 = v93
		v217 = v148
		goto L75
	} else {
		goto L76
	}
L69:
	;
	v160 = v69
	v164 = v148
	goto L70
L70:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v164<<(uint(int32(1))%32)))))
	if v170 != 0 {
		v311 = int32(1)
		goto L67
	} else {
		goto L72
	}
L71:
	;
	v179 = v174
	v183 = v172
	goto L68
L72:
	;
	v171 = int32(1)
	v172 = v164 + v171
	v174 = v160 - v171
	if v174 <= v93 {
		v179 = v174
		v183 = v172
		goto L68
	} else {
		goto L73
	}
L73:
	;
	if v172 < v71 {
		v160 = v174
		v164 = v172
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	if v179 != v215 {
		v257 = v183
		v258 = v217
		goto L82
	} else {
		goto L83
	}
L76:
	;
	v196 = v93
	v198 = v148
	goto L77
L77:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+v198<<(uint(int32(1))%32)))))
	if v203 != 0 {
		v311 = int32(-1)
		goto L67
	} else {
		goto L79
	}
L78:
	;
	v215 = v207
	v217 = v205
	goto L75
L79:
	;
	v204 = int32(1)
	v205 = v198 + v204
	v207 = v196 - v204
	if v207 <= v179 {
		v215 = v207
		v217 = v205
		goto L75
	} else {
		goto L80
	}
L80:
	;
	if v205 < v95 {
		v196 = v207
		v198 = v205
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	if v71 < v257 {
		goto L91
	} else {
		goto L92
	}
L83:
	;
	v226 = v183
	v227 = v217
	goto L84
L84:
	;
	if base.B2i32(v71 <= v226)|base.B2i32(v95 <= v227) != 0 {
		v257 = v226
		v258 = v227
		goto L82
	} else {
		goto L86
	}
L85:
	;
	if base.I32_extend16_s(v243) < base.I32_extend16_s(v241) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v232 = int32(1)
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v226<<(uint(v232)%32)))))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227<<(uint(v232)%32)+v141))))
	if v241 == v243 {
		v226 = v226 + v232
		v227 = v227 + v232
		goto L84
	} else {
		goto L87
	}
L87:
	;
	goto L85
L88:
	;
	v250 = int32(1)
	goto L90
L89:
	;
	v250 = int32(-1)
	goto L90
L90:
	;
	v321 = v250
	goto L66
L91:
	;
	v261 = v257
	goto L93
L92:
	;
	v261 = v71
	goto L93
L93:
	;
	v268 = v257
	goto L94
L94:
	;
	if v261 == v268 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v311 = v294
	goto L67
L96:
	;
	if v95 < v258 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v294 = int32(1)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v268<<(uint(v294)%32)))))
	if v300 == int32(0) {
		v268 = v268 + v294
		goto L94
	} else {
		goto L108
	}
L99:
	;
	v273 = v258
	goto L101
L100:
	;
	v273 = v95
	goto L101
L101:
	;
	v281 = v258
	goto L102
L102:
	;
	if v273 == v281 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v311 = int32(-1)
	goto L67
L104:
	;
	v321 = int32(0)
	goto L66
L105:
	;
	goto L106
L106:
	;
	v285 = int32(1)
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281<<(uint(v285)%32)+v141))))
	if v290 == int32(0) {
		v281 = v281 + v285
		goto L102
	} else {
		goto L107
	}
L107:
	;
	goto L103
L108:
	;
	goto L95
L109:
	;
	return int32(-1)
L110:
	;
	goto L111
L111:
	;
	v327 = int32(0)
	if base.B2i32(v69 < v93)&base.B2i32(v327 < v95) == v327 {
		v358 = v93
		v362 = v327
		goto L114
	} else {
		goto L115
	}
L112:
	;
	return v500
L113:
	;
	v500 = v490
	goto L112
L114:
	;
	if base.B2i32(v71 <= int32(0))|base.B2i32(v69 <= v358) != 0 {
		v394 = v69
		v396 = v327
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v339 = v93
	v343 = v327
	goto L116
L116:
	;
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+v343<<(uint(int32(1))%32)))))
	if v349 != 0 {
		v490 = int32(1)
		goto L113
	} else {
		goto L118
	}
L117:
	;
	v358 = v353
	v362 = v351
	goto L114
L118:
	;
	v350 = int32(1)
	v351 = v343 + v350
	v353 = v339 - v350
	if v353 <= v69 {
		v358 = v353
		v362 = v351
		goto L114
	} else {
		goto L119
	}
L119:
	;
	if v351 < v95 {
		v339 = v353
		v343 = v351
		goto L116
	} else {
		goto L120
	}
L120:
	;
	goto L117
L121:
	;
	if v358 != v394 {
		v436 = v362
		v437 = v396
		goto L128
	} else {
		goto L129
	}
L122:
	;
	v375 = v69
	v377 = v327
	goto L123
L123:
	;
	v382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v377<<(uint(int32(1))%32)))))
	if v382 != 0 {
		v490 = int32(-1)
		goto L113
	} else {
		goto L125
	}
L124:
	;
	v394 = v386
	v396 = v384
	goto L121
L125:
	;
	v383 = int32(1)
	v384 = v377 + v383
	v386 = v375 - v383
	if v386 <= v358 {
		v394 = v386
		v396 = v384
		goto L121
	} else {
		goto L126
	}
L126:
	;
	if v384 < v71 {
		v375 = v386
		v377 = v384
		goto L123
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	if v95 < v436 {
		goto L137
	} else {
		goto L138
	}
L129:
	;
	v405 = v362
	v406 = v396
	goto L130
L130:
	;
	if base.B2i32(v95 <= v405)|base.B2i32(v71 <= v406) != 0 {
		v436 = v405
		v437 = v406
		goto L128
	} else {
		goto L132
	}
L131:
	;
	if base.I32_extend16_s(v422) < base.I32_extend16_s(v420) {
		goto L134
	} else {
		goto L135
	}
L132:
	;
	v411 = int32(1)
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+v405<<(uint(v411)%32)))))
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v406<<(uint(v411)%32)+v136))))
	if v420 == v422 {
		v405 = v405 + v411
		v406 = v406 + v411
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v429 = int32(1)
	goto L136
L135:
	;
	v429 = int32(-1)
	goto L136
L136:
	;
	v500 = v429
	goto L112
L137:
	;
	v440 = v436
	goto L139
L138:
	;
	v440 = v95
	goto L139
L139:
	;
	v447 = v436
	goto L140
L140:
	;
	if v440 == v447 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v490 = v473
	goto L113
L142:
	;
	if v71 < v437 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v473 = int32(1)
	v479 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141+v447<<(uint(v473)%32)))))
	if v479 == int32(0) {
		v447 = v447 + v473
		goto L140
	} else {
		goto L154
	}
L145:
	;
	v452 = v437
	goto L147
L146:
	;
	v452 = v71
	goto L147
L147:
	;
	v460 = v437
	goto L148
L148:
	;
	if v452 == v460 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v490 = int32(-1)
	goto L113
L150:
	;
	v500 = int32(0)
	goto L112
L151:
	;
	goto L152
L152:
	;
	v464 = int32(1)
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460<<(uint(v464)%32)+v136))))
	if v469 == int32(0) {
		v460 = v460 + v464
		goto L148
	} else {
		goto L153
	}
L153:
	;
	goto L149
L154:
	;
	goto L141
}
