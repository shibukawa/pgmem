package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_checkNameSpaceConflicts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v133 int32
	_ = v133
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(0)
	if v22 < v19 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = v19
	goto L7
L6:
	;
	v25 = v22
	goto L7
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = v3
	goto L8
L8:
	;
	if l1 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v133 = v33 + int32(1)
	if v133 != v25 {
		v33 = v133
		goto L8
	} else {
		goto L36
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26+v33<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+20)))
	if v45&int32(1) == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v50 <= int32(0) {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	v55 = int32(0)
	if v55 < v50 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = v50
	goto L16
L15:
	;
	v58 = v55
	goto L16
L16:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = int32(0)
	goto L17
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v60+v62<<(uint(int32(2))%32))))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+20)))
	if v78 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L10
L19:
	;
	v118 = v62 + int32(1)
	if v118 != v58 {
		v62 = v118
		goto L17
	} else {
		goto L35
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v87 == int32(0) {
		v106 = v86
		v107 = v87
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v107-v106 != 0 {
		goto L19
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	if v86 != v87 {
		v106 = v86
		v107 = v87
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v91 = v83
	v92 = v54
	goto L25
L25:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+1)))
	if v96 == int32(0) {
		v106 = v95
		v107 = v96
		goto L22
	} else {
		goto L27
	}
L26:
	;
	v106 = v95
	v107 = v96
	goto L22
L27:
	;
	v99 = int32(1)
	if v95 == v96 {
		v91 = v91 + v99
		v92 = v92 + v99
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if v109 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v110 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v111 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v113 == v114 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L19
L35:
	;
	goto L18
L36:
	;
	goto L9
L37:
	;
	return
L38:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v54
	F_errmsg(m, int32(416575), v15)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(497507), int32(474), int32(125413))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parseNameAndArgTypes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v213 int32
	_ = v213
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
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v346 int32
	_ = v346
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	v7 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_pstrdup(m, l0)
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
	v21 = v17
	v27 = v7
	goto L3
L3:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	switch v33 - int32(34) {
	case 0:
		goto L10
	case 1, 2, 3, 4, 5:
		v404 = v27
		goto L5
	case 6:
		goto L9
	default:
		goto L8
	}
L5:
	;
	v21 = v21 + int32(1)
	v27 = v404
	goto L3
L6:
	;
	m.G0 = v15 + int32(16)
	return v394
L7:
	;
	v62 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v62)
	v65 = F_stringToQualifiedNameList(m, v17, l5)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L8:
	;
	if v33 != 0 {
		v404 = v27
		goto L5
	} else {
		goto L12
	}
L9:
	;
	v40 = int32(1)
	if v27&v40 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L10:
	;
	v36 = int32(1)
	v21 = v21 + v36
	v27 = v27 ^ v36
	goto L3
L11:
	;
	v404 = v40
	goto L5
L12:
	;
	v45 = int32(0)
	v46 = F_errsave_start(m, l5)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v46 == int32(0) {
		v394 = v45
		goto L6
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errmsg(m, int32(155169), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l5, int32(501061), int32(1924), int32(163105))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v394 = v45
	goto L6
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
	if v65 == int32(0) {
		v394 = v62
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v71 = v21 + int32(1)
	v72 = F_strlen(m, v71)
	mBase = m.M
	v74 = v72 + v71
	goto L20
L20:
	;
	v87 = v74 - int32(1)
	if base.Ui32(v71) < base.Ui32(v87) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v99 != int32(41) {
		goto L27
	} else {
		goto L28
	}
L22:
	;
	v89 = int32(*(*int8)(unsafe.Add(mBase, uint32(v87))))
	goto L25
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	if base.B2i32(v89 == int32(32))|base.B2i32(base.Ui32((v89-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v74 = v87
		goto L20
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v102 = F_errsave_start(m, l5)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v118 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v118)
	v121 = v118
	v128 = v71
	v131 = v7
	goto L35
L30:
	;
	if v102 == int32(0) {
		v394 = v62
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(155140), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errsave_finish(m, l5, int32(501061), int32(1942), int32(163105))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v394 = v62
	goto L6
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v121
	v141 = v128
	goto L37
L36:
	;
	v394 = v363
	goto L6
L37:
	;
	v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v141))))
	goto L39
L38:
	;
	v158 = int32(0)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v160 == v158 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	if base.B2i32(v148 == int32(32))|base.B2i32(base.Ui32((v148-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v141 = v141 + int32(1)
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(100) <= v360 {
		goto L102
	} else {
		goto L103
	}
L42:
	;
	F_pfree(m, v17)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L101
	}
L43:
	;
	if v131 == int32(0) {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v183 = v158
	v189 = v141
	v190 = v160
	v192 = v158
	goto L54
L46:
	;
	v166 = int32(0)
	v167 = F_errsave_start(m, l5)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v167 == int32(0) {
		v394 = v166
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(382732), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, l5, int32(501061), int32(1961), int32(163105))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v394 = v166
	goto L6
L52:
	;
	v260 = v189 - int32(1)
	if base.Ui32(v260) < base.Ui32(v141) {
		goto L76
	} else {
		goto L77
	}
L53:
	;
	v252 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v252)
	v254 = int32(1)
	v257 = v254
	v258 = v189 + v254
	goto L52
L54:
	;
	v196 = v190 & int32(255)
	switch v196 - int32(34) {
	case 0:
		goto L61
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L58
	case 10:
		goto L60
	default:
		goto L59
	}
L55:
	;
	v227 = int32(0)
	if (v183|base.B2i32(v192 != v227))&int32(1) == v227 {
		v257 = v227
		v258 = v189
		goto L52
	} else {
		goto L70
	}
L56:
	;
	goto L55
L57:
	;
	v225 = v189 + int32(1)
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	v183 = v222
	v189 = v225
	v190 = v226
	v192 = v223
	goto L54
L58:
	;
	if v183&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L59:
	;
	if v196 == int32(0) {
		goto L56
	} else {
		goto L63
	}
L60:
	;
	v201 = int32(0)
	if (v183|base.B2i32(v192 != v201))&int32(1) == v201 {
		goto L53
	} else {
		goto L62
	}
L61:
	;
	v222 = v183 ^ int32(1)
	v223 = v192
	goto L57
L62:
	;
	goto L58
L63:
	;
	goto L58
L64:
	;
	v222 = int32(1)
	v223 = v192
	goto L57
L65:
	;
	goto L66
L66:
	;
	v213 = int32(0)
	switch v196 - int32(40) {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L69
	}
L67:
	;
	v222 = v213
	v223 = v192 - int32(1)
	goto L57
L68:
	;
	v222 = v213
	v223 = v192 + int32(1)
	goto L57
L69:
	;
	switch v196 - int32(91) {
	case 0:
		goto L68
	default:
		v222 = v213
		v223 = v192
		goto L57
	case 2:
		goto L67
	}
L70:
	;
	v235 = int32(0)
	v236 = F_errsave_start(m, l5)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v236 == int32(0) {
		v394 = v235
		goto L6
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(382678), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, l5, int32(501061), int32(1993), int32(163105))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v394 = v235
	goto L6
L76:
	;
	if l1 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v262 = v260
	goto L78
L78:
	;
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v262))))
	goto L80
L79:
	;
	goto L76
L80:
	;
	if base.B2i32(v274 == int32(32))|base.B2i32(base.Ui32((v274-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v286 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v262))) = uint8(v286)
	v289 = v262 - int32(1)
	if base.Ui32(v141) <= base.Ui32(v289) {
		v262 = v289
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v355 = F_parseTypeString(m, v141, v15+int32(12), v15+int32(8), l5)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L99
	}
L84:
	;
	v308 = v141
	v309 = int32(373553)
	goto L86
L85:
	;
	if v346 != 0 {
		goto L83
	} else {
		goto L98
	}
L86:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v312 == v313 {
		v335 = v312
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v346 = int32(0)
	goto L85
L88:
	;
	v337 = int32(1)
	if v335 != 0 {
		v308 = v308 + v337
		v309 = v309 + v337
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v312-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v323 = v312 | int32(32)
	goto L92
L91:
	;
	v323 = v312
	goto L92
L92:
	;
	if base.Ui32((v313-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v332 = v313 | int32(32)
	goto L95
L94:
	;
	v332 = v313
	goto L95
L95:
	;
	if v323 == v332 {
		v335 = v323
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v346 = v323 - v332
	goto L85
L97:
	;
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	goto L41
L99:
	;
	if v355 != 0 {
		goto L41
	} else {
		goto L100
	}
L100:
	;
	v394 = int32(0)
	goto L6
L101:
	;
	v394 = int32(1)
	goto L6
L102:
	;
	v363 = int32(0)
	v364 = F_errsave_start(m, l5)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v360<<(uint(int32(2))%32)))) = v383
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v121 = v385 + int32(1)
	v128 = v258
	v131 = v257
	goto L35
L104:
	;
	goto L36
L105:
	;
	if v364 == int32(0) {
		v394 = v363
		goto L6
	} else {
		goto L106
	}
L106:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(121238), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errsave_finish(m, l5, int32(501061), int32(2029), int32(163105))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	goto L104
}
func F_scanNameSpaceForENR(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = F_get_visible_ENR_metadata(m, v3, l1)
	mBase = m.M
	return base.B2i32(v4 != int32(0))
}
