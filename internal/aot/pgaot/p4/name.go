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
	var v32 int32
	_ = v32
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
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v134 int32
	_ = v134
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
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
	v154 = m.ExcPending
	if v154 != 0 {
		goto L36
	} else {
		goto L37
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
	v32 = v3
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
	v134 = v32 + int32(1)
	if v134 != v25 {
		v32 = v134
		goto L8
	} else {
		goto L35
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v26+v32<<(uint(int32(2))%32))))
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
	v119 = v62 + int32(1)
	if v119 != v58 {
		v62 = v119
		goto L17
	} else {
		goto L34
	}
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if base.B2i32(v86 == int32(0))|base.B2i32(v86 != v89) != 0 {
		v107 = v86
		v108 = v89
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v107-v108 != 0 {
		goto L19
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v92 = v83
	v93 = v54
	goto L24
L24:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+1)))
	if v97 == int32(0) {
		v107 = v97
		v108 = v96
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v107 = v97
	v108 = v96
	goto L22
L26:
	;
	v100 = int32(1)
	if v97 == v96 {
		v92 = v92 + v100
		v93 = v93 + v100
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	if v110 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v114 == v115 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L19
L34:
	;
	goto L18
L35:
	;
	goto L9
L36:
	;
	return
L37:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v54
	F_errmsg(m, int32(_a_F_checkNameSpaceConflicts_0), v15)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_checkNameSpaceConflicts_1), int32(470), int32(_a_F_checkNameSpaceConflicts_2))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v406 int32
	_ = v406
	v7 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = F_pstrdup(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = v18
	v30 = v7
	goto L4
L3:
	;
	m.G0 = v16 + int32(16)
	return v406
L4:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v35 != int32(34) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v18)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L109
	}
L6:
	;
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v394 = int32(1)
	v22 = v22 + v394
	v30 = v30 ^ v394
	goto L4
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v63)
	v65 = F_stringToQualifiedNameList(m, v18, l5)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	if (base.B2i32(v35 != int32(40))|v30)&int32(1) == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v47 = F_errsave_start(m, l5)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v22 = v22 + int32(1)
	goto L4
L15:
	;
	if v47 == int32(0) {
		v406 = v7
		goto L3
	} else {
		goto L16
	}
L16:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_0), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2049), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v406 = v7
	goto L3
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
	if v65 == int32(0) {
		v406 = v7
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v70 = int32(1)
	v71 = v22 + v70
	v72 = F_strlen(m, v71)
	mBase = m.M
	v81 = v72 + v70
	goto L22
L22:
	;
	v89 = v81 - int32(1)
	v90 = v22 + v89
	if v81 < int32(3) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	if v104 != int32(41) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	goto L23
L25:
	;
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90))))
	goto L26
L26:
	;
	if base.B2i32(v93 == int32(32))|base.B2i32(base.Ui32((v93-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v81 = v89
		goto L22
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v107 = int32(0)
	v108 = F_errsave_start(m, l5)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v124)
	v127 = v124
	v129 = v71
	v139 = v7
	goto L36
L31:
	;
	if v108 == int32(0) {
		v406 = v107
		goto L3
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_3), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2067), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v406 = v107
	goto L3
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v127
	v143 = v129
	goto L38
L38:
	;
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v143))))
	goto L40
L39:
	;
	v166 = int32(0)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	if v168 == v166 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if base.B2i32(v156 == int32(32))|base.B2i32(base.Ui32((v156-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v143 = v143 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	if v139 == int32(0) {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v191 = v168
	v193 = v143
	v197 = v166
	v201 = v166
	goto L51
L45:
	;
	v174 = int32(0)
	v175 = F_errsave_start(m, l5)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v175 == int32(0) {
		v406 = v174
		goto L3
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_4), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2086), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v406 = v174
	goto L3
L51:
	;
	if v191 != int32(34) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	v191 = v393
	v193 = v193 + int32(1)
	v197 = v389
	v201 = v390
	goto L51
L54:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if int32(100) <= v361 {
		goto L101
	} else {
		goto L102
	}
L55:
	;
	if v191 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	goto L57
L57:
	;
	v389 = v197 ^ int32(1)
	v390 = v201
	goto L53
L58:
	;
	v259 = v193 - int32(1)
	if base.Ui32(v259) < base.Ui32(v143) {
		goto L76
	} else {
		goto L77
	}
L59:
	;
	v251 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v193))) = uint8(v251)
	v253 = int32(1)
	v256 = v253
	v257 = v193 + v253
	goto L58
L60:
	;
	if (base.B2i32(v191 != int32(44))|v197)&int32(1)|v201 == int32(0) {
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v226 = int32(0)
	if (v197|base.B2i32(v201 != v226))&int32(1) == v226 {
		v256 = v226
		v257 = v193
		goto L58
	} else {
		goto L70
	}
L63:
	;
	if v197&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v389 = int32(1)
	v390 = v201
	goto L53
L65:
	;
	goto L66
L66:
	;
	v217 = int32(0)
	switch v191 - int32(40) {
	case 0:
		goto L68
	case 1:
		goto L67
	default:
		goto L69
	}
L67:
	;
	v389 = v217
	v390 = v201 - int32(1)
	goto L53
L68:
	;
	v389 = v217
	v390 = v201 + int32(1)
	goto L53
L69:
	;
	switch v191 - int32(91) {
	case 0:
		goto L68
	default:
		v389 = v217
		v390 = v201
		goto L53
	case 2:
		goto L67
	}
L70:
	;
	v234 = int32(0)
	v235 = F_errsave_start(m, l5)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v235 == int32(0) {
		v406 = v234
		goto L3
	} else {
		goto L72
	}
L72:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_5), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2118), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v406 = v234
	goto L3
L76:
	;
	if l1 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	v261 = v259
	goto L78
L78:
	;
	v274 = int32(*(*int8)(unsafe.Add(mBase, uint32(v261))))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v261))) = uint8(v286)
	v289 = v261 - int32(1)
	if base.Ui32(v143) <= base.Ui32(v289) {
		v261 = v289
		goto L78
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	v356 = F_parseTypeString(m, v143, v16+int32(12), v16+int32(8), l5)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L99
	}
L84:
	;
	v309 = v143
	v310 = int32(_a_F_parseNameAndArgTypes_6)
	goto L86
L85:
	;
	if v347 != 0 {
		goto L83
	} else {
		goto L98
	}
L86:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310))))
	if v313 == v314 {
		v336 = v313
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v347 = int32(0)
	goto L85
L88:
	;
	v338 = int32(1)
	if v336 != 0 {
		v309 = v309 + v338
		v310 = v310 + v338
		goto L86
	} else {
		goto L97
	}
L89:
	;
	if base.Ui32((v313-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v324 = v313 | int32(32)
	goto L92
L91:
	;
	v324 = v313
	goto L92
L92:
	;
	if base.Ui32((v314-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v333 = v314 | int32(32)
	goto L95
L94:
	;
	v333 = v314
	goto L95
L95:
	;
	if v324 == v333 {
		v336 = v324
		goto L88
	} else {
		goto L96
	}
L96:
	;
	v347 = v324 - v333
	goto L85
L97:
	;
	goto L87
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	goto L54
L99:
	;
	if v356 != 0 {
		goto L54
	} else {
		goto L100
	}
L100:
	;
	v406 = int32(0)
	goto L3
L101:
	;
	v364 = int32(0)
	v365 = F_errsave_start(m, l5)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l4+v361<<(uint(int32(2))%32)))) = v384
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v127 = v386 + int32(1)
	v129 = v257
	v139 = v256
	goto L36
L104:
	;
	if v365 == int32(0) {
		v406 = v364
		goto L3
	} else {
		goto L105
	}
L105:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(_a_F_parseNameAndArgTypes_7), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errsave_finish(m, l5, int32(_a_F_parseNameAndArgTypes_1), int32(2154), int32(_a_F_parseNameAndArgTypes_2))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v406 = v364
	goto L3
L109:
	;
	v406 = int32(1)
	goto L3
}
