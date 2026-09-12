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
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v72
L2:
	;
	if v6 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v39 = base.B2i32(v6 != int32(0))
	goto L4
L4:
	;
	if v39 != 0 {
		v72 = v39
		goto L1
	} else {
		goto L16
	}
L5:
	;
	return int32(-1)
L6:
	;
	goto L7
L7:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
	if v15 == int32(0) {
		v34 = v14
		v35 = v15
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v39 = v35 - v34
	goto L4
L9:
	;
	goto L8
L10:
	;
	if v14 != v15 {
		v34 = v14
		v35 = v15
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v19 = v7
	v20 = v6
	goto L12
L12:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v24 == int32(0) {
		v34 = v23
		v35 = v24
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v34 = v23
	v35 = v24
	goto L9
L14:
	;
	v27 = int32(1)
	if v23 == v24 {
		v19 = v19 + v27
		v20 = v20 + v27
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 == v40 {
		v72 = v40
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v44 == int32(0) {
		v72 = v40
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v47 == v48 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	if v50 == v51 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(v48) < base.Ui32(v47) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)))
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+6)))
	if v53 == v54 {
		v72 = v40
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(v51) < base.Ui32(v50) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	if base.Ui32(v54) < base.Ui32(v53) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v59 = int32(-1)
	goto L28
L27:
	;
	v59 = int32(1)
	goto L28
L28:
	;
	return v59
L29:
	;
	v64 = int32(-1)
	goto L31
L30:
	;
	v64 = int32(1)
	goto L31
L31:
	;
	return v64
L32:
	;
	v69 = int32(-1)
	goto L34
L33:
	;
	v69 = int32(1)
	goto L34
L34:
	;
	v72 = v69
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v170 int32
	_ = v170
	v7 = int32(0)
	if base.B2i32(l5 < l2)&base.B2i32(v7 < l1) == v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v170
L2:
	;
	if l5 <= v38 {
		v73 = l5
		v75 = v7
		goto L11
	} else {
		goto L12
	}
L3:
	;
	v38 = l2
	v42 = v7
	goto L2
L4:
	;
	goto L5
L5:
	;
	v19 = l2
	v23 = v7
	goto L6
L6:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v23<<(uint(int32(1))%32)))))
	if v29 != 0 {
		v170 = int32(1)
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v33
	v42 = v31
	goto L2
L8:
	;
	v30 = int32(1)
	v31 = v23 + v30
	v33 = v19 - v30
	if v33 <= l5 {
		v38 = v33
		v42 = v31
		goto L2
	} else {
		goto L9
	}
L9:
	;
	if v31 < l1 {
		v19 = v33
		v23 = v31
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	if v38 != v73 {
		v115 = v42
		v116 = v75
		goto L19
	} else {
		goto L20
	}
L12:
	;
	if l4 <= int32(0) {
		v73 = l5
		v75 = v7
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = l5
	v56 = v7
	goto L14
L14:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v56<<(uint(int32(1))%32)))))
	if v61 != 0 {
		v170 = int32(-1)
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v73 = v65
	v75 = v63
	goto L11
L16:
	;
	v62 = int32(1)
	v63 = v56 + v62
	v65 = v54 - v62
	if v65 <= v38 {
		v73 = v65
		v75 = v63
		goto L11
	} else {
		goto L17
	}
L17:
	;
	if v63 < l4 {
		v54 = v65
		v56 = v63
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	if l1 < v115 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v84 = v42
	v85 = v75
	goto L21
L21:
	;
	if l1 <= v84 {
		v115 = v84
		v116 = v85
		goto L19
	} else {
		goto L23
	}
L22:
	;
	if base.I32_extend16_s(v100) < base.I32_extend16_s(v98) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if l4 <= v85 {
		v115 = v84
		v116 = v85
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v89 = int32(1)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v84<<(uint(v89)%32)))))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85<<(uint(v89)%32)+l3))))
	if v98 == v100 {
		v84 = v84 + v89
		v85 = v85 + v89
		goto L21
	} else {
		goto L25
	}
L25:
	;
	goto L22
L26:
	;
	v107 = int32(1)
	goto L28
L27:
	;
	v107 = int32(-1)
	goto L28
L28:
	;
	return v107
L29:
	;
	v119 = v115
	goto L31
L30:
	;
	v119 = l1
	goto L31
L31:
	;
	v126 = v115
	goto L32
L32:
	;
	if v119 == v126 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v170 = v153
	goto L1
L34:
	;
	if l4 < v116 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v153 = int32(1)
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+v126<<(uint(v153)%32)))))
	if v159 == int32(0) {
		v126 = v126 + v153
		goto L32
	} else {
		goto L46
	}
L37:
	;
	v131 = v116
	goto L39
L38:
	;
	v131 = l4
	goto L39
L39:
	;
	v139 = v116
	goto L40
L40:
	;
	if v131 == v139 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v170 = int32(-1)
	goto L1
L42:
	;
	return int32(0)
L43:
	;
	goto L44
L44:
	;
	v144 = int32(1)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v139<<(uint(v144)%32)))))
	if v149 == int32(0) {
		v139 = v139 + v144
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	goto L33
}
func F_cmp_numerics(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v483 int32
	_ = v483
	var v493 int32
	_ = v493
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = int32(_a_F_cmp_numerics_0)
	v18 = v16 & v17
	if v18 == v17 {
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
	if base.Ui32(int32(_a_F_cmp_numerics_0)) <= base.Ui32(v15) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	if v15 != int32(_a_F_cmp_numerics_2) {
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
	if v15 == int32(_a_F_cmp_numerics_0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return base.B2i32(v15 != int32(_a_F_cmp_numerics_0))
L9:
	;
	v33 = int32(-1)
	goto L11
L10:
	;
	v33 = base.B2i32(v15 != int32(_a_F_cmp_numerics_1))
	goto L11
L11:
	;
	return v33
L12:
	;
	v39 = int32(-1)
	goto L14
L13:
	;
	v39 = int32(0)
	goto L14
L14:
	;
	return v39
L15:
	;
	if v15 == int32(_a_F_cmp_numerics_2) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v50 = l0 + int32(6)
	v55 = base.B2i32(int32(0) <= base.I32_extend16_s(v16))
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v47 = int32(1)
	goto L20
L19:
	;
	v47 = int32(-1)
	goto L20
L20:
	;
	return v47
L21:
	;
	v56 = int32(-8)
	goto L23
L22:
	;
	v56 = int32(-6)
	goto L23
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50))))
	v70 = v60
	goto L26
L25:
	;
	v70 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
	goto L26
L26:
	;
	v71 = v56 + int32(base.Ui32(v57)>>(uint(int32(2))%32))
	v73 = l1 + int32(6)
	v78 = base.B2i32(int32(0) <= base.I32_extend16_s(v15))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v79 = int32(-8)
	goto L29
L28:
	;
	v79 = int32(-6)
	goto L29
L29:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) <= base.I32_extend16_s(v15) {
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
	v93 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
	goto L32
L32:
	;
	v94 = v79 + int32(base.Ui32(v80)>>(uint(int32(2))%32))
	v100 = v15 & int32(_a_F_cmp_numerics_0)
	if v100 == int32(_a_F_cmp_numerics_3) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v103 = v15 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_4)
	goto L35
L34:
	;
	v103 = v100
	goto L35
L35:
	;
	if base.Ui32(v71) <= base.Ui32(int32(1)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.Ui32(v94) < base.Ui32(int32(2)) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v18 == int32(_a_F_cmp_numerics_3) {
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
	if v103 == int32(_a_F_cmp_numerics_4) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(1)
	goto L44
L43:
	;
	v114 = int32(-1)
	goto L44
L44:
	;
	return v114
L45:
	;
	v122 = v16 << (uint(int32(1)) % 32) & int32(_a_F_cmp_numerics_4)
	goto L47
L46:
	;
	v122 = v18
	goto L47
L47:
	;
	if base.Ui32(v94) <= base.Ui32(int32(1)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	if v122 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	if int32(0) <= base.I32_extend16_s(v16) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v127 = int32(-1)
	goto L53
L52:
	;
	v127 = int32(1)
	goto L53
L53:
	;
	return v127
L54:
	;
	v131 = l0 + int32(8)
	goto L56
L55:
	;
	v131 = v50
	goto L56
L56:
	;
	v133 = int32(base.Ui32(v71) >> (uint(int32(1)) % 32))
	if int32(0) <= base.I32_extend16_s(v15) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v136 = l1 + int32(8)
	goto L59
L58:
	;
	v136 = v73
	goto L59
L59:
	;
	v138 = int32(base.Ui32(v94) >> (uint(int32(1)) % 32))
	if v122 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v103 == int32(_a_F_cmp_numerics_4) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	if v103 == int32(0) {
		goto L113
	} else {
		goto L114
	}
L63:
	;
	return int32(1)
L64:
	;
	goto L65
L65:
	;
	v145 = int32(0)
	if base.B2i32(v93 < v70)&base.B2i32(v145 < v133) == v145 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	return v316
L67:
	;
	v316 = v306
	goto L66
L68:
	;
	if v93 <= v176 {
		v211 = v93
		v213 = v145
		goto L77
	} else {
		goto L78
	}
L69:
	;
	v176 = v70
	v180 = v145
	goto L68
L70:
	;
	goto L71
L71:
	;
	v157 = v70
	v161 = v145
	goto L72
L72:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v161<<(uint(int32(1))%32)))))
	if v167 != 0 {
		v306 = int32(1)
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v176 = v171
	v180 = v169
	goto L68
L74:
	;
	v168 = int32(1)
	v169 = v161 + v168
	v171 = v157 - v168
	if v171 <= v93 {
		v176 = v171
		v180 = v169
		goto L68
	} else {
		goto L75
	}
L75:
	;
	if v169 < v133 {
		v157 = v171
		v161 = v169
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	if v176 != v211 {
		v252 = v180
		v253 = v213
		goto L85
	} else {
		goto L86
	}
L78:
	;
	if v138 <= int32(0) {
		v211 = v93
		v213 = v145
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v192 = v93
	v194 = v145
	goto L80
L80:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v194<<(uint(int32(1))%32)))))
	if v199 != 0 {
		v306 = int32(-1)
		goto L67
	} else {
		goto L82
	}
L81:
	;
	v211 = v203
	v213 = v201
	goto L77
L82:
	;
	v200 = int32(1)
	v201 = v194 + v200
	v203 = v192 - v200
	if v203 <= v176 {
		v211 = v203
		v213 = v201
		goto L77
	} else {
		goto L83
	}
L83:
	;
	if v201 < v138 {
		v192 = v203
		v194 = v201
		goto L80
	} else {
		goto L84
	}
L84:
	;
	goto L81
L85:
	;
	if v133 < v252 {
		goto L95
	} else {
		goto L96
	}
L86:
	;
	v222 = v180
	v223 = v213
	goto L87
L87:
	;
	if v133 <= v222 {
		v252 = v222
		v253 = v223
		goto L85
	} else {
		goto L89
	}
L88:
	;
	if base.I32_extend16_s(v238) < base.I32_extend16_s(v236) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v138 <= v223 {
		v252 = v222
		v253 = v223
		goto L85
	} else {
		goto L90
	}
L90:
	;
	v227 = int32(1)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v222<<(uint(v227)%32)))))
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v223<<(uint(v227)%32)+v136))))
	if v236 == v238 {
		v222 = v222 + v227
		v223 = v223 + v227
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v245 = int32(1)
	goto L94
L93:
	;
	v245 = int32(-1)
	goto L94
L94:
	;
	v316 = v245
	goto L66
L95:
	;
	v256 = v252
	goto L97
L96:
	;
	v256 = v133
	goto L97
L97:
	;
	v263 = v252
	goto L98
L98:
	;
	if v256 == v263 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v306 = v289
	goto L67
L100:
	;
	if v138 < v253 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v289 = int32(1)
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v263<<(uint(v289)%32)))))
	if v295 == int32(0) {
		v263 = v263 + v289
		goto L98
	} else {
		goto L112
	}
L103:
	;
	v268 = v253
	goto L105
L104:
	;
	v268 = v138
	goto L105
L105:
	;
	v276 = v253
	goto L106
L106:
	;
	if v268 == v276 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v306 = int32(-1)
	goto L67
L108:
	;
	v316 = int32(0)
	goto L66
L109:
	;
	goto L110
L110:
	;
	v280 = int32(1)
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v276<<(uint(v280)%32)))))
	if v285 == int32(0) {
		v276 = v276 + v280
		goto L106
	} else {
		goto L111
	}
L111:
	;
	goto L107
L112:
	;
	goto L99
L113:
	;
	return int32(-1)
L114:
	;
	goto L115
L115:
	;
	v322 = int32(0)
	if base.B2i32(v70 < v93)&base.B2i32(v322 < v138) == v322 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	return v493
L117:
	;
	v493 = v483
	goto L116
L118:
	;
	if v70 <= v353 {
		v388 = v70
		v390 = v322
		goto L127
	} else {
		goto L128
	}
L119:
	;
	v353 = v93
	v357 = v322
	goto L118
L120:
	;
	goto L121
L121:
	;
	v334 = v93
	v338 = v322
	goto L122
L122:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v338<<(uint(int32(1))%32)))))
	if v344 != 0 {
		v483 = int32(1)
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v353 = v348
	v357 = v346
	goto L118
L124:
	;
	v345 = int32(1)
	v346 = v338 + v345
	v348 = v334 - v345
	if v348 <= v70 {
		v353 = v348
		v357 = v346
		goto L118
	} else {
		goto L125
	}
L125:
	;
	if v346 < v138 {
		v334 = v348
		v338 = v346
		goto L122
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	if v353 != v388 {
		v429 = v357
		v430 = v390
		goto L135
	} else {
		goto L136
	}
L128:
	;
	if v133 <= int32(0) {
		v388 = v70
		v390 = v322
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v369 = v70
	v371 = v322
	goto L130
L130:
	;
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v371<<(uint(int32(1))%32)))))
	if v376 != 0 {
		v483 = int32(-1)
		goto L117
	} else {
		goto L132
	}
L131:
	;
	v388 = v380
	v390 = v378
	goto L127
L132:
	;
	v377 = int32(1)
	v378 = v371 + v377
	v380 = v369 - v377
	if v380 <= v353 {
		v388 = v380
		v390 = v378
		goto L127
	} else {
		goto L133
	}
L133:
	;
	if v378 < v133 {
		v369 = v380
		v371 = v378
		goto L130
	} else {
		goto L134
	}
L134:
	;
	goto L131
L135:
	;
	if v138 < v429 {
		goto L145
	} else {
		goto L146
	}
L136:
	;
	v399 = v357
	v400 = v390
	goto L137
L137:
	;
	if v138 <= v399 {
		v429 = v399
		v430 = v400
		goto L135
	} else {
		goto L139
	}
L138:
	;
	if base.I32_extend16_s(v415) < base.I32_extend16_s(v413) {
		goto L142
	} else {
		goto L143
	}
L139:
	;
	if v133 <= v400 {
		v429 = v399
		v430 = v400
		goto L135
	} else {
		goto L140
	}
L140:
	;
	v404 = int32(1)
	v413 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v399<<(uint(v404)%32)))))
	v415 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v400<<(uint(v404)%32)+v131))))
	if v413 == v415 {
		v399 = v399 + v404
		v400 = v400 + v404
		goto L137
	} else {
		goto L141
	}
L141:
	;
	goto L138
L142:
	;
	v422 = int32(1)
	goto L144
L143:
	;
	v422 = int32(-1)
	goto L144
L144:
	;
	v493 = v422
	goto L116
L145:
	;
	v433 = v429
	goto L147
L146:
	;
	v433 = v138
	goto L147
L147:
	;
	v440 = v429
	goto L148
L148:
	;
	if v433 == v440 {
		goto L150
	} else {
		goto L151
	}
L149:
	;
	v483 = v466
	goto L117
L150:
	;
	if v133 < v430 {
		goto L153
	} else {
		goto L154
	}
L151:
	;
	goto L152
L152:
	;
	v466 = int32(1)
	v472 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136+v440<<(uint(v466)%32)))))
	if v472 == int32(0) {
		v440 = v440 + v466
		goto L148
	} else {
		goto L162
	}
L153:
	;
	v445 = v430
	goto L155
L154:
	;
	v445 = v133
	goto L155
L155:
	;
	v453 = v430
	goto L156
L156:
	;
	if v445 == v453 {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v483 = int32(-1)
	goto L117
L158:
	;
	v493 = int32(0)
	goto L116
L159:
	;
	goto L160
L160:
	;
	v457 = int32(1)
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131+v453<<(uint(v457)%32)))))
	if v462 == int32(0) {
		v453 = v453 + v457
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L157
L162:
	;
	goto L149
}
