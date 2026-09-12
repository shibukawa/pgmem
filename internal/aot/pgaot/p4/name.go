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
	F_errmsg(m, int32(406502), v15)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(485194), int32(474), int32(122013))
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
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v402 int32
	_ = v402
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v450 int32
	_ = v450
	var v460 int32
	_ = v460
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
		v460 = v27
		goto L5
	case 6:
		goto L9
	default:
		goto L8
	}
L5:
	;
	v21 = v21 + int32(1)
	v27 = v460
	goto L3
L6:
	;
	m.G0 = v15 + int32(16)
	return v450
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
		v460 = v27
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
	v460 = v40
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
		v450 = v45
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
	F_errmsg(m, int32(151309), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errsave_finish(m, l5, int32(488619), int32(1924), int32(159088))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v450 = v45
	goto L6
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
	if v65 == int32(0) {
		v450 = v62
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v71 = v21 + int32(1)
	if v71&int32(3) == int32(0) {
		v95 = v71
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v130 = v128 + v71
	goto L37
L21:
	;
	v128 = v120 - v71
	goto L20
L22:
	;
	v99 = v95
	goto L31
L23:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v79 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v128 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v84 = v71
	goto L27
L27:
	;
	v88 = v84 + int32(1)
	if v88&int32(3) == int32(0) {
		v95 = v88
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v120 = v88
	goto L21
L29:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v93 != 0 {
		v84 = v88
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 == v108 {
		v99 = v99 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v114 = v99
	goto L34
L33:
	;
	goto L32
L34:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v118 != 0 {
		v114 = v114 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v120 = v114
	goto L21
L36:
	;
	goto L35
L37:
	;
	v143 = v130 - int32(1)
	if base.Ui32(v71) < base.Ui32(v143) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v155 != int32(41) {
		goto L44
	} else {
		goto L45
	}
L39:
	;
	v145 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143))))
	goto L42
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	if base.B2i32(v145 == int32(32))|base.B2i32(base.Ui32((v145-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v130 = v143
		goto L37
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v158 = F_errsave_start(m, l5)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v143))) = uint8(v174)
	v177 = v174
	v184 = v71
	v187 = v7
	goto L52
L47:
	;
	if v158 == int32(0) {
		v450 = v62
		goto L6
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(151280), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errsave_finish(m, l5, int32(488619), int32(1942), int32(159088))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v450 = v62
	goto L6
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v177
	v197 = v184
	goto L54
L53:
	;
	v450 = v419
	goto L6
L54:
	;
	v204 = int32(*(*int8)(unsafe.Add(mBase, uint32(v197))))
	goto L56
L55:
	;
	v214 = int32(0)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v216 == v214 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	if base.B2i32(v204 == int32(32))|base.B2i32(base.Ui32((v204-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v197 = v197 + int32(1)
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(100) <= v416 {
		goto L119
	} else {
		goto L120
	}
L59:
	;
	F_pfree(m, v17)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L118
	}
L60:
	;
	if v187 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v239 = v214
	v245 = v197
	v246 = v216
	v248 = v214
	goto L71
L63:
	;
	v222 = int32(0)
	v223 = F_errsave_start(m, l5)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v223 == int32(0) {
		v450 = v222
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(373114), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_errsave_finish(m, l5, int32(488619), int32(1961), int32(159088))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v450 = v222
	goto L6
L69:
	;
	v316 = v245 - int32(1)
	if base.Ui32(v316) < base.Ui32(v197) {
		goto L93
	} else {
		goto L94
	}
L70:
	;
	v308 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v308)
	v310 = int32(1)
	v313 = v310
	v314 = v245 + v310
	goto L69
L71:
	;
	v252 = v246 & int32(255)
	switch v252 - int32(34) {
	case 0:
		goto L78
	case 1, 2, 3, 4, 5, 6, 7, 8, 9:
		goto L75
	case 10:
		goto L77
	default:
		goto L76
	}
L72:
	;
	v283 = int32(0)
	if (v239|base.B2i32(v248 != v283))&int32(1) == v283 {
		v313 = v283
		v314 = v245
		goto L69
	} else {
		goto L87
	}
L73:
	;
	goto L72
L74:
	;
	v281 = v245 + int32(1)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v239 = v278
	v245 = v281
	v246 = v282
	v248 = v279
	goto L71
L75:
	;
	if v239&int32(1) != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	if v252 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L77:
	;
	v257 = int32(0)
	if (v239|base.B2i32(v248 != v257))&int32(1) == v257 {
		goto L70
	} else {
		goto L79
	}
L78:
	;
	v278 = v239 ^ int32(1)
	v279 = v248
	goto L74
L79:
	;
	goto L75
L80:
	;
	goto L75
L81:
	;
	v278 = int32(1)
	v279 = v248
	goto L74
L82:
	;
	goto L83
L83:
	;
	v269 = int32(0)
	switch v252 - int32(40) {
	case 0:
		goto L85
	case 1:
		goto L84
	default:
		goto L86
	}
L84:
	;
	v278 = v269
	v279 = v248 - int32(1)
	goto L74
L85:
	;
	v278 = v269
	v279 = v248 + int32(1)
	goto L74
L86:
	;
	switch v252 - int32(91) {
	case 0:
		goto L85
	default:
		v278 = v269
		v279 = v248
		goto L74
	case 2:
		goto L84
	}
L87:
	;
	v291 = int32(0)
	v292 = F_errsave_start(m, l5)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v292 == int32(0) {
		v450 = v291
		goto L6
	} else {
		goto L89
	}
L89:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_errmsg(m, int32(373060), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_errsave_finish(m, l5, int32(488619), int32(1993), int32(159088))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	v450 = v291
	goto L6
L93:
	;
	if l1 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L94:
	;
	v318 = v316
	goto L95
L95:
	;
	v330 = int32(*(*int8)(unsafe.Add(mBase, uint32(v318))))
	goto L97
L96:
	;
	goto L93
L97:
	;
	if base.B2i32(v330 == int32(32))|base.B2i32(base.Ui32((v330-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v342 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v342)
	v345 = v318 - int32(1)
	if base.Ui32(v197) <= base.Ui32(v345) {
		v318 = v345
		goto L95
	} else {
		goto L99
	}
L99:
	;
	goto L96
L100:
	;
	v411 = F_parseTypeString(m, v197, v15+int32(12), v15+int32(8), l5)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L116
	}
L101:
	;
	v364 = v197
	v365 = int32(364359)
	goto L103
L102:
	;
	if v402 != 0 {
		goto L100
	} else {
		goto L115
	}
L103:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364))))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v365))))
	if v368 == v369 {
		v391 = v368
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v402 = int32(0)
	goto L102
L105:
	;
	v393 = int32(1)
	if v391 != 0 {
		v364 = v364 + v393
		v365 = v365 + v393
		goto L103
	} else {
		goto L114
	}
L106:
	;
	if base.Ui32((v368-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v379 = v368 | int32(32)
	goto L109
L108:
	;
	v379 = v368
	goto L109
L109:
	;
	if base.Ui32((v369-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v388 = v369 | int32(32)
	goto L112
L111:
	;
	v388 = v369
	goto L112
L112:
	;
	if v379 == v388 {
		v391 = v379
		goto L105
	} else {
		goto L113
	}
L113:
	;
	v402 = v379 - v388
	goto L102
L114:
	;
	goto L104
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	goto L58
L116:
	;
	if v411 != 0 {
		goto L58
	} else {
		goto L117
	}
L117:
	;
	v450 = int32(0)
	goto L6
L118:
	;
	v450 = int32(1)
	goto L6
L119:
	;
	v419 = int32(0)
	v420 = F_errsave_start(m, l5)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v416<<(uint(int32(2))%32)))) = v439
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v177 = v441 + int32(1)
	v184 = v314
	v187 = v313
	goto L52
L121:
	;
	goto L53
L122:
	;
	if v420 == int32(0) {
		v450 = v419
		goto L6
	} else {
		goto L123
	}
L123:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errmsg(m, int32(117922), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errsave_finish(m, l5, int32(488619), int32(2029), int32(159088))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L121
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
