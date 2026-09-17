package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATAddCheckNNConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
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
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_ATAddCheckNNConstraint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v20
	v23 = *(*int64)(unsafe.Add(mBase, _c_F_ATAddCheckNNConstraint[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
	F_check_stack_depth(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l6 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ATSimplePermissions(m, int32(16), l3, int32(289))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = F_copyObjectImpl(m, l4)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v31
	v39 = F_list_make1_impl(m, int32(1), v17+int32(12))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	m.G0 = v17 + int32(32)
	return
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L86
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L82
	}
L11:
	;
	v45 = F_AddRelationNewConstraints(m, l3, int32(0), v39, l6|l7, l6^int32(1), l7, int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v47 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L81
	}
L16:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v171
	F_CommandCounterIncrement(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L50
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v171 = v51
	v178 = v50
	v183 = l0 + int32(4)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+21)))
	if v56 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v74 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v57 == int32(1) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v61 = F_palloc0(m, int32(32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v70 = F_lappend(m, v69, v61)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+64)) = v70
	goto L20
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v77
	goto L27
L26:
	;
	goto L27
L27:
	;
	v79 = int32(1)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v80 == v79 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+12)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
	F_set_attnotnull(m, l1, l3, v83, (v84^int32(-1))&int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if int32(1) < v91 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L30
L32:
	;
	v103 = v79
	goto L35
L33:
	;
	v157 = v55
	goto L34
L34:
	;
	v171 = int32(2606)
	v178 = int32(0)
	v183 = v157 + int32(4)
	goto L16
L35:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108+v103<<(uint(int32(2))%32))))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+21)))
	if v113 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v157 = v112
	goto L34
L37:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v131 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v114 == int32(1) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v118 = F_palloc0(m, int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+24)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v127 = F_lappend(m, v126, v118)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+64)) = v127
	goto L37
L42:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v134
	goto L44
L43:
	;
	goto L44
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v136 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+12)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+15)))
	F_set_attnotnull(m, l1, l3, v139, (v140^int32(-1))&int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v148 = v103 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v148 < v149 {
		v103 = v148
		goto L35
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	goto L36
L50:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+17)))
	if v190 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v193 = F_find_inheritance_children(m, v192, l8)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v193 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v195 = l5
	goto L55
L54:
	;
	v195 = int32(1)
	goto L55
L55:
	;
	if v195 == int32(0) {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	if v193 == int32(0) {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v201 <= v200 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v206 = v200
	goto L59
L59:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218+v206<<(uint(int32(2))%32))))
	v224 = F_table_open(m, v222, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L8
L61:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+48))
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+118)))
	if v227 == int32(116) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+24)))
	if v230 == int32(0) {
		goto L9
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_CheckTableNotInUse(m, v224, int32(_a_F_ATAddCheckNNConstraint_0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v224)+56))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v237 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_ATAddCheckNNConstraint(m, v17+int32(16), l1, v314, v224, l4, l5, int32(1), l7, l8)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L78
	}
L68:
	;
	v283 = F_palloc0(m, int32(140))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L75
	}
L69:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v240 <= int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v237)+12))
	v251 = int32(0)
	goto L71
L71:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v243+v251<<(uint(int32(2))%32))))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)))
	if v263 == v236 {
		v314 = v262
		goto L67
	} else {
		goto L73
	}
L72:
	;
	goto L68
L73:
	;
	v266 = v251 + int32(1)
	if v240 != v266 {
		v251 = v266
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = v236
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v224)+48))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v288)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+4)) = uint8(v289)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v224)+52))
	v292 = F_CreateTupleDescCopyConstr(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+8)) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v283)+88)) = int64(0)
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+84)) = uint8(v297)
	v299 = int32(_a_F_ATAddCheckNNConstraint_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v283)+96)) = uint16(v299)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v302 = F_lappend(m, v301, v283)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v302
	v314 = v283
	goto L67
L78:
	;
	F_relation_close(m, v224, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v328 = v206 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	if v328 < v329 {
		v206 = v328
		goto L59
	} else {
		goto L80
	}
L80:
	;
	goto L60
L81:
	;
	goto L8
L82:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_ATAddCheckNNConstraint_2), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ATAddCheckNNConstraint_3), int32(_a_F_ATAddCheckNNConstraint_4), int32(_a_F_ATAddCheckNNConstraint_5))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_ATAddCheckNNConstraint_6), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ATAddCheckNNConstraint_3), int32(_a_F_ATAddCheckNNConstraint_7), int32(_a_F_ATAddCheckNNConstraint_8))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecAddColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v142 int32
	_ = v142
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
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
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
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v403 int64
	_ = v403
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v518 int32
	_ = v518
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v583 int32
	_ = v583
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v644 int32
	_ = v644
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v971 int32
	_ = v971
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	v26 = m.G0
	v28 = v26 - int32(128)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+28)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l6 != 0 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L216
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L212
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L208
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L205
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L201
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L197
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L190
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L186
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L182
	}
L12:
	;
	v46 = l3 + int32(48)
	v49 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	F_ATSimplePermissions(m, v37, l3, int32(289))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+131)))
	if v42 == int32(1) {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	goto L12
L18:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+16)))
	if v51 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	m.G0 = v28 + int32(128)
	return
L20:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v129 = F_check_for_column_name_collision(m, l3, v126, v32&int32(1))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v55 = F_SearchSysCacheCopyAttName(m, v30, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v55 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_typenameTypeIdAndMod(m, int32(0), v62, v28+int32(116), v28+int32(104))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v70 = v59 + v60
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
	if v69 != v71 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v28)+104))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	if v73 != v74 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v77 = F_GetColumnDefCollation(m, int32(0), v33, v69)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+96))
	if v77 != v79 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v70)+94)))
	v83 = v81 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+94)) = uint16(v83)
	if base.I32_extend16_s(v83) != v83 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	F_CatalogTupleUpdate(m, v49, v55+int32(4), v55)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v55)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v95 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v95 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v97 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_0), v28+int32(32))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_relation_close(m, v49, int32(3))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_2), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v120
	v123 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddColumn[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v123
	goto L19
L40:
	;
	if v129 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_relation_close(m, v49, int32(3))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v142 = int32(0)
	if l6|base.B2i32(l9 == v142) == v142 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v137
	v140 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddColumn[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v140
	goto L19
L45:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v148 = F_ATParseTransformCmd(m, l2, l3, v147, l5, l8, l9)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v152 = v33
	goto L47
L47:
	;
	if l5 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	v152 = v151
	goto L47
L49:
	;
	v167 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L55
	}
L50:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)))
	if v155 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+119)))
	if v159 == int32(112) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v163 = F_find_inheritance_children(m, v30, int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v163 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	v171 = F_SearchSysCacheCopy(m, int32(57), v30, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v171 == int32(0) {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+22)))
	v177 = v175 + v176
	v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v177)+120)))
	if int32(1600) <= v178 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v28)+100)) = v152
	v187 = F_list_make1_impl(m, int32(1), v28+int32(28))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v189 = F_BuildDescForRelation(m, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v194 = v189 + v191<<(uint(int32(4))%32)
	v196 = v178 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v194)+94)) = uint16(v196)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+116))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v201
	v204 = int32(24)
	v209 = F_list_make1_impl(m, int32(472), v28+v204)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+110)))
	if v213 == int32(118) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v216 = int32(8)
	goto L64
L63:
	;
	v216 = int32(0)
	goto L64
L64:
	;
	F_CheckAttributeType(m, v194+v204, v198, v199, v209, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v219 = int32(0)
	F_InsertPgAttributeTuples(m, v49, v189, v30, v219, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_relation_close(m, v49, int32(3))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+120)) = uint16(v196)
	F_CatalogTupleUpdate(m, v167, v171+int32(4), v171)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v171)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[2]))
	if v234 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v30, v196, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_relation_close(m, v167, int32(3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	if v244 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v246 = F_palloc(m, int32(12))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	switch v181 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L85
	default:
		goto L84
	}
L79:
	;
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+94)))
	*(*uint16)(unsafe.Add(mBase, uint32(v246))) = uint16(v248)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	v251 = F_copyObjectImpl(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+4)) = v251
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v246)+8)) = uint8(v254)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v246
	v261 = F_list_make1_impl(m, int32(1), v28+int32(20))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v263 = int32(0)
	v268 = F_AddRelationNewConstraints(m, l3, v261, v263, v263, int32(1), v263, v263)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(1247)
	v529 = v28 + int32(116)
	v531 = v28 + int32(104)
	F_recordDependencyOn(m, v529, v531, int32(110))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L143
	}
L85:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)))
	if v275 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+19)))
	v508 = v506 | v507
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)) = uint8(v508)
	goto L84
L87:
	;
	v348 = F_expression_planner(m, v343)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L108
	}
L88:
	;
	v277 = F_palloc0(m, int32(12))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v293 = int32(*(*int16)(unsafe.Add(mBase, uint32(v194)+94)))
	v294 = F_build_column_default(m, l3, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = int32(59)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v152)+40))
	v282 = int32(0)
	v286 = F_RangeVarGetRelidExtended(m, v281, v282, v282, v282, v282)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+4)) = v286
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v277)+8)) = v289
	v291 = F_DomainHasConstraints(m, v289)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v343 = v277
	v346 = v291
	goto L87
L94:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	v297 = F_DomainHasConstraints(m, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v299 = int32(0)
	if v294|base.B2i32(v297 == v299) == v299 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v305 = v194 + int32(96)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v306
	v309 = v194 + int32(88)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v313 = F_getBaseTypeAndTypmod(m, v310, v28+int32(116))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v294 == int32(0) {
		goto L86
	} else {
		goto L107
	}
L99:
	;
	v315 = F_get_typcollation(m, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v319 = F_makeNullConst(m, v313, v318, v315)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v309)))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v326 = F_coerce_to_target_type(m, int32(0), v319, v313, v321, v322, int32(1), int32(2), int32(-1))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v326 != 0 {
		v343 = v326
		v346 = v297
		goto L87
	} else {
		goto L103
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_4), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_5), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	v343 = v294
	v346 = v297
	goto L87
L108:
	;
	v351 = F_palloc0(m, int32(16))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v353 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+94)))
	*(*int32)(unsafe.Add(mBase, uint32(v351)+4)) = v348
	*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v353)
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v351)+12)) = uint8(base.B2i32(v356 != int32(0)))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v361 = F_lappend(m, v360, v351)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+68)) = v361
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v364)+119)))
	if v365 != int32(114) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	F_FreeExecutorState(m, v376)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L142
	}
L112:
	;
	if v489&int32(255) == int32(118) {
		goto L86
	} else {
		goto L141
	}
L113:
	;
	v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	v489 = v488
	goto L112
L114:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	if (base.B2i32(v368 != int32(0))|v346)&int32(1) != 0 {
		v489 = v368
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v374 = F_contain_volatile_functions(m, v348)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v374 != 0 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v376 = F_CreateExecutorState(m)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v378 = F_ExecPrepareExpr(m, v348, v376)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v376)+152))
	if v380 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v383 = v380
	goto L122
L121:
	;
	v381 = F_MakePerTupleExprContext(m, v376)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v378)+20))
	v387 = m.T0[v386].(func(*base.Module, int32, int32, int32) int32)(m, v378, v383, v28+int32(116))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	v383 = v381
	goto L122
L124:
	;
	v389 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+116)))
	if v389 != 0 {
		goto L111
	} else {
		goto L125
	}
L125:
	;
	v390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v194)+94)))
	v391 = m.G0
	v393 = v391 - int32(192)
	m.G0 = v393
	*(*int32)(unsafe.Add(mBase, uint32(v393)+188)) = v387
	v398 = int32(0)
	base.MemoryFill(m, v393+int32(80), v398, int32(96))
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+72)) = uint8(v398)
	v403 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v393)+64)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v393)+56)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v393)+48)) = v403
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+40)) = uint8(v398)
	*(*int64)(unsafe.Add(mBase, uint32(v393)+32)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v393)+24)) = v403
	*(*int64)(unsafe.Add(mBase, uint32(v393)+16)) = v403
	v419 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v423 = F_SearchSysCache2(m, int32(7), v422, v390)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v423 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v423)+16))
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v445)+22)))
	v447 = v445 + v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+68))
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447)+72)))
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+82)))
	v451 = int32(*(*int8)(unsafe.Add(mBase, uint32(v447)+83)))
	v452 = F_construct_array(m, v393+int32(188), int32(1), v448, v449, v450, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L134
	}
L131:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v393)+4)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v393))) = v390
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_6), v393)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_7), int32(2051), int32(_a_F_ATExecAddColumn_8))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L134:
	;
	v454 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v393)+132)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v393)+188)) = v452
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+29)) = uint8(v454)
	*(*int32)(unsafe.Add(mBase, uint32(v393)+176)) = v452
	*(*uint8)(unsafe.Add(mBase, uint32(v393)+40)) = uint8(v454)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v419)+52))
	v469 = F_heap_modify_tuple(m, v423, v462, v393+int32(80), v393+int32(48), v393+int32(16))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_CatalogTupleUpdate(m, v419, v469+int32(4), v469)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_ReleaseCatCache(m, v423)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_relation_close(m, v419, int32(3))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	m.G0 = v393 + int32(192)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_FreeExecutorState(m, v376)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L84
L141:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v494 | int32(2)
	goto L86
L142:
	;
	goto L86
L143:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v194)+116))
	v536 = int32(0)
	if base.B2i32(v535 == v536)|base.B2i32(v535 == int32(100)) == v536 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v535
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(3456)
	F_recordDependencyOn(m, v529, v531, int32(110))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v557 = F_find_inheritance_children(m, v556, l7)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	if v557 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v559 = l5
	goto L151
L150:
	;
	v559 = int32(1)
	goto L151
L151:
	;
	if v559 == int32(0) {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l6 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v565 = F_copyObjectImpl(m, v562)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	v573 = v562
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v573
	if v557 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v565)+20))
	v568 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+18)) = uint8(v568)
	v570 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v567)+16)) = uint16(v570)
	v573 = v565
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L19
L158:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v577 <= int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v583 = int32(0)
	goto L160
L160:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606+v583<<(uint(int32(2))%32))))
	v612 = F_table_open(m, v610, int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L162
	}
L161:
	;
	goto L157
L162:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v612)+48))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+118)))
	if v615 == int32(116) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+24)))
	if v618 == int32(0) {
		goto L3
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	F_CheckTableNotInUse(m, v612, int32(_a_F_ATExecAddColumn_9))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v612)+56))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v625 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	F_ATExecAddColumn(m, v28+int32(116), l1, v718, v612, v28+int32(104), l5, int32(1), l7, l8, l9)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L179
	}
L169:
	;
	v693 = F_palloc0(m, int32(140))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L176
	}
L170:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v628 <= int32(0) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	v644 = int32(0)
	goto L172
L172:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v631+v644<<(uint(int32(2))%32))))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	if v662 == v624 {
		v718 = v661
		goto L168
	} else {
		goto L174
	}
L173:
	;
	goto L169
L174:
	;
	v665 = v644 + int32(1)
	if v628 != v665 {
		v644 = v665
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v693))) = v624
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v612)+48))
	v699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v698)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v693)+4)) = uint8(v699)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v612)+52))
	v702 = F_CreateTupleDescCopyConstr(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+8)) = v702
	*(*int64)(unsafe.Add(mBase, uint32(v693)+88)) = int64(0)
	v707 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v693)+84)) = uint8(v707)
	v709 = int32(_a_F_ATExecAddColumn_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v693)+96)) = uint16(v709)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v712 = F_lappend(m, v711, v693)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v712
	v718 = v693
	goto L168
L179:
	;
	F_relation_close(m, v612, int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v751 = v583 + int32(1)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v751 < v752 {
		v583 = v751
		goto L160
	} else {
		goto L181
	}
L181:
	;
	goto L161
L182:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_11), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_12), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L186:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v834 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_13), v28+int32(80))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_14), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L190:
	;
	F_errcode(m, int32(17432708))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v858
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v857 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_15), v28-int32(-64))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v868 = F_get_collation_name(m, v77)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v70)+96))
	v871 = F_get_collation_name(m, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v871
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v868
	F_errdetail(m, int32(_a_F_ATExecAddColumn_16), v28+int32(48))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_17), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_18), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_19), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L201:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_20), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_21), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_22), v28)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_23), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(1600)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_24), v28+int32(16))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_25), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_26), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_27), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_28), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_29), int32(_a_F_ATExecAddColumn_30))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecAddIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(144)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	if base.B2i32(l5 == v8)&base.B2i32(v21 == int32(112)) == v8 {
		goto L9
	} else {
		goto L10
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L16
	} else {
		goto L78
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L16
	} else {
		goto L74
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L16
	} else {
		goto L69
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L16
	} else {
		goto L66
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L62
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L16
	} else {
		goto L58
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L16
	} else {
		goto L54
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L16
	} else {
		goto L50
	}
L9:
	;
	if l6 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L16
	} else {
		goto L45
	}
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
	if v29&int32(1) != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v34 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	return
L17:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v37 = F_SearchSysCacheCopyAttName(m, v36, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v37 == int32(0) {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+74)))
	if v44 <= int32(0) {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+86)))
	if v47 == int32(0) {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v51 = F_findNotNullConstraintAttnum(m, v50, v44)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	if v51 == int32(0) {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+22)))
	v57 = v55 + v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+76)))
	if v58 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+89)))
	if v61 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+87)))
	if v62 == int32(1) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+89)) = uint8(v65)
	F_CatalogTupleUpdate(m, v34, v37+int32(4), v37)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddIdentity[0]))
	if v72 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+74)))
	v76 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v74, v75, v76, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L16
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v82
	F_pfree(m, v37)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v21 != int32(112)) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	m.G0 = v16 + int32(144)
	return
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v96 = F_find_inheritance_children(m, v95, l4)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	if v96 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v100 <= int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v105 = int32(0)
	goto L39
L39:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v105<<(uint(int32(2))%32))))
	v125 = F_table_open(m, v123, int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L16
	} else {
		goto L41
	}
L40:
	;
	goto L34
L41:
	;
	v127 = int32(1)
	F_ATExecAddIdentity(m, v16+int32(132), v125, l2, l3, l4, v127, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L16
	} else {
		goto L42
	}
L42:
	;
	F_relation_close(m, v125, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	v135 = v105 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v135 < v136 {
		v105 = v135
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_0), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	F_errhint(m, int32(_a_F_ATExecAddIdentity_1), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_3), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_5), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_6), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v197 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_7), v16)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_8), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_9), v16+int32(16))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_10), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v235 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_11), v16+int32(112))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_12), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v254 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ATExecAddIdentity_13), v16+int32(32))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_14), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L16
	} else {
		goto L70
	}
L70:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v277 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v57 + v277
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v276 + v277
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_15), v16+int32(96))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = int32(_a_F_ATExecAddIdentity_16)
	F_errhint(m, int32(_a_F_ATExecAddIdentity_17), v16+int32(80))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_18), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v307 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_19), v16-int32(-64))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_20), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v329 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddIdentity_21), v16+int32(48))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_ATExecAddIdentity_2), int32(_a_F_ATExecAddIdentity_22), int32(_a_F_ATExecAddIdentity_4))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecAlterConstrDeferrability(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v215 int32
	_ = v215
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	F_check_stack_depth(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
	v28 = v26 + v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+96))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+73)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	if v30 == v31 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v20 + int32(48)
	return v215
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+74)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v33 == v34 {
		v215 = int32(0)
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_AlterConstrUpdateConstraintEntry(m, l0, l1, l4)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	F_ScanKeyInit(m, v20, int32(11), int32(3), int32(184), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(1)
	v51 = F_systable_beginscan(m, l2, int32(2699), v46, int32(0), v46, v20)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v53 = F_systable_getnext(m, v51)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v53 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v63 = v53
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_systable_endscan(m, v51)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L33
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v74 = v72 + v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v75 != v76 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v79 = F_list_append_unique_oid(m, v78, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+76))
	v84 = v82 - int32(1644)
	v91 = int32(0)
	if base.B2i32(base.Ui32(int32(11)) < base.Ui32(v84))|base.B2i32(int32(1)<<(uint(v84)%32)&int32(3075) == v91) == v91 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v79
	goto L19
L21:
	;
	v96 = F_heap_copytuple(m, v63)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v120 = F_systable_getnext(m, v51)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+16))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+22)))
	v100 = v98 + v99
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+97)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+96)) = uint8(v39)
	F_CatalogTupleUpdate(m, l2, v96+int32(4), v96)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAlterConstrDeferrability[0]))
	if v108 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v111 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v110, v111, v111, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v96)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L23
L31:
	;
	if v120 != 0 {
		v63 = v120
		goto L15
	} else {
		goto L32
	}
L32:
	;
	goto L16
L33:
	;
	if l5 == int32(0) {
		v215 = v46
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+119)))
	if v144 != int32(112) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v147 = F_get_rel_relkind(m, v29)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+22)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154+v155)))
	F_ScanKeyInit(m, v20, int32(12), int32(3), int32(184), v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	if v147 != int32(112) {
		v215 = v46
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v161 = int32(1)
	v164 = F_systable_beginscan(m, l1, int32(2579), v161, int32(0), v161, v20)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	v183 = F_systable_getnext(m, v164)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	F_systable_endscan(m, v164)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L51
	}
L44:
	;
	if v183 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v183)+16))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+22)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185+v186)+80))
	v189 = F_table_open(m, v188, l7)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L43
L48:
	;
	v192 = F_ATExecAlterConstrDeferrability(m, l0, l1, l2, v189, v183, int32(1), l6, l7)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_relation_close(m, v189, int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v215 = v46
	goto L3
}
func F_ATExecChangeOwner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
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
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int64
	_ = v314
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v486 int32
	_ = v486
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	v14 = m.G0
	v16 = v14 - int32(592)
	m.G0 = v16
	v18 = F_relation_open(m, l0, l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v608 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[0]))
	if v608 != 0 {
		goto L166
	} else {
		goto L167
	}
L5:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if v422 != 0 {
		goto L121
	} else {
		goto L122
	}
L6:
	;
	if v372 == int32(73) {
		goto L5
	} else {
		goto L119
	}
L7:
	;
	if int32(1)<<(uint(v374)%32)&int32(_a_F_ATExecChangeOwner_0) != 0 {
		goto L5
	} else {
		goto L118
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L113
	}
L9:
	;
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v29 = v27 + v28
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	switch v30 - int32(73) {
	case 0:
		goto L17
	default:
		goto L8
	case 10:
		goto L16
	case 26:
		goto L15
	case 29, 36, 39, 41, 45:
		v161 = l1
		goto L13
	case 32:
		goto L18
	case 43:
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L110
	}
L13:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v164 == v161 {
		goto L4
	} else {
		goto L55
	}
L14:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L54
	}
L15:
	;
	if l2 != 0 {
		v161 = l1
		goto L13
	} else {
		goto L48
	}
L16:
	;
	if l2 != 0 {
		v161 = l1
		goto L13
	} else {
		goto L34
	}
L17:
	;
	if l2 != 0 {
		v161 = l1
		goto L13
	} else {
		goto L28
	}
L18:
	;
	if l2 != 0 {
		v161 = l1
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v33 == l1 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v161 = v61
	goto L13
L21:
	;
	v37 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v37 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v29 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_1), v16+int32(32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errhint(m, int32(_a_F_ATExecChangeOwner_2), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_4), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L20
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v29 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_1), v16+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(_a_F_ATExecChangeOwner_2), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_6), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v86 == l1 {
		v161 = l1
		goto L13
	} else {
		goto L35
	}
L35:
	;
	v90 = v16 + int32(224)
	v92 = v16 + int32(432)
	v93 = F_sequenceIsOwned(m, l0, int32(97), v90, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v93 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v98 = F_sequenceIsOwned(m, l0, int32(105), v90, v92)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	if v98 == int32(0) {
		v161 = l1
		goto L13
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v110 = v29 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v110
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_7), v16+int32(80))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v16)+224))
	v118 = F_get_rel_name(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v110
	F_errdetail(m, int32(_a_F_ATExecChangeOwner_8), v16-int32(-64))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_9), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v29 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_10), v16+int32(112))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = int32(_a_F_ATExecChangeOwner_11)
	F_errhint(m, int32(_a_F_ATExecChangeOwner_12), v16+int32(96))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_13), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v161 = l1
	goto L13
L55:
	;
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v213 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v213
	v215 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+208)) = uint16(v215)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+200)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+192)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+184)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+176)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v213
	*(*int64)(unsafe.Add(mBase, uint32(v16)+152)) = v213
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+160)) = uint16(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v161
	v234 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+133)) = uint8(v234)
	v240 = F_SysCacheGetAttr(m, int32(57), v25, int32(32), v16+int32(127))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	v166 = F_superuser(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v166 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[1]))
	v172 = F_object_ownercheck(m, int32(1259), l0, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v172 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v177 = F_get_rel_relkind(m, l0)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[1]))
	F_check_can_set_role(m, v197, v161)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L73
	}
L64:
	;
	switch v177 - int32(73) {
	case 0, 32:
		v188 = int32(20)
		goto L66
	default:
		goto L67
	case 10:
		goto L71
	case 29:
		goto L68
	case 36:
		goto L69
	case 45:
		goto L70
	}
L65:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_aclcheck_error(m, int32(2), v190, v191+int32(4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	v190 = v188
	goto L65
L67:
	;
	v188 = int32(41)
	goto L66
L68:
	;
	v190 = int32(18)
	goto L65
L69:
	;
	v190 = int32(23)
	goto L65
L70:
	;
	v190 = int32(51)
	goto L65
L71:
	;
	v190 = int32(37)
	goto L65
L72:
	;
	goto L63
L73:
	;
	v202 = F_object_aclcheck(m, int32(2615), v168, v161, int64(512))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v202 == int32(0) {
		goto L56
	} else {
		goto L75
	}
L75:
	;
	v207 = F_get_namespace_name(m, v168)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_aclcheck_error(m, v202, int32(36), v207)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L56
L78:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+127)))
	if v242 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v245 = F_pg_detoast_datum(m, v240)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v260 = F_heap_modify_tuple(m, v25, v253, v16+int32(224), v16+int32(176), v16+int32(128))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v248 = F_aclnewowner(m, v245, v247, v161)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+348)) = v248
	v251 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)) = uint8(v251)
	goto L81
L84:
	;
	F_CatalogTupleUpdate(m, v22, v260+int32(4), v260)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_pfree(m, v260)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v271 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v274 = v16 + int32(544)
	F_ScanKeyInit(m, v274, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v281 = int32(1)
	v284 = F_systable_beginscan(m, v271, int32(2659), v281, int32(0), v281, v274)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v286 = F_systable_getnext(m, v284)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v286 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v290 = v286
	goto L94
L92:
	;
	goto L93
L93:
	;
	F_systable_endscan(m, v284)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L107
	}
L94:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v290)+16))
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301)+22)))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v301+v302)+91)))
	if v304 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	v352 = F_systable_getnext(m, v284)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L105
	}
L97:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v271)+52))
	v309 = F_heap_getattr_6(m, v290, int32(22), v306, v16+int32(367))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+367)))
	if v311 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+424)) = uint8(v312)
	v314 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+416)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v16)+408)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v16)+400)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v16)+368)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v16)+376)) = v314
	*(*int64)(unsafe.Add(mBase, uint32(v16)+384)) = v314
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+392)) = uint8(v312)
	v328 = F_pg_detoast_datum(m, v309)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v330 = F_aclnewowner(m, v328, v268, v161)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+516)) = v330
	v333 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+389)) = uint8(v333)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v271)+52))
	v342 = F_heap_modify_tuple(m, v290, v335, v16+int32(432), v16+int32(400), v16+int32(368))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_CatalogTupleUpdate(m, v271, v342+int32(4), v342)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_pfree(m, v342)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	if v352 != 0 {
		v290 = v352
		goto L94
	} else {
		goto L106
	}
L106:
	;
	goto L95
L107:
	;
	F_relation_close(m, v271, int32(3))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	v374 = v372 - int32(99)
	if base.Ui32(v374) <= base.Ui32(int32(17)) {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	goto L6
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(_a_F_ATExecChangeOwner_14), v16)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_15), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v29 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_16), v16+int32(16))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v405 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29)+119)))
	F_errdetail_relkind_not_supported(m, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_17), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	goto L6
L119:
	;
	F_changeDependencyOnOwner(m, int32(1259), l0, v161)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L5
L121:
	;
	F_AlterTypeOwnerInternal(m, v422, v161)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	v427 = v425 - int32(109)
	v434 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v427))|base.B2i32(int32(1)<<(uint(v427)%32)&int32(169) == v434) == v434 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v439 = F_RelationGetIndexList(m, v18)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v500 != 0 {
		goto L137
	} else {
		goto L138
	}
L128:
	;
	F_list_free(m, v439)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L136
	}
L129:
	;
	if v439 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v443 <= int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v449 = int32(0)
	goto L132
L132:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v460+v449<<(uint(int32(2))%32))))
	F_ATExecChangeOwner(m, v464, v161, int32(1), l3)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L128
L134:
	;
	v469 = v449 + int32(1)
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	if v469 < v470 {
		v449 = v469
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	goto L127
L137:
	;
	F_ATExecChangeOwner(m, v500, v161, int32(1), l3)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v506 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v509 = v16 + int32(432)
	F_ScanKeyInit(m, v509, int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_ScanKeyInit(m, v16+int32(480), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v527 = F_systable_beginscan(m, v506, int32(2674), int32(1), int32(0), int32(2), v509)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v529 = F_systable_getnext(m, v527)
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	if v529 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v533 = v529
	goto L149
L147:
	;
	goto L148
L148:
	;
	F_systable_endscan(m, v527)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L1
	} else {
		goto L164
	}
L149:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v533)+16))
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v544)+22)))
	v546 = v544 + v545
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v546)+20))
	if v547 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	v574 = F_systable_getnext(m, v527)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L162
	}
L152:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v550 != int32(1259) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v546)+8))
	if v553 != 0 {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+24)))
	switch v554 - int32(97) {
	case 0, 8:
		goto L155
	default:
		goto L151
	}
L155:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	v558 = F_relation_open(m, v557, l3)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v558)+48))
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v560)+119)))
	if v561 == int32(83) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	F_ATExecChangeOwner(m, v564, v161, int32(1), l3)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	v569 = l3
	goto L159
L159:
	;
	F_relation_close(m, v558, v569)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	v569 = int32(0)
	goto L159
L161:
	;
	goto L151
L162:
	;
	if v574 != 0 {
		v533 = v574
		goto L149
	} else {
		goto L163
	}
L163:
	;
	goto L150
L164:
	;
	F_relation_close(m, v506, int32(1))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L4
L166:
	;
	v610 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v610, v610, v610)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_ReleaseCatCache(m, v25)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	F_relation_close(m, v22, int32(3))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	m.G0 = v16 + int32(592)
	return
}
func F_ATExecForceNoForceRowSecurity(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v14 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v18 = F_SearchSysCacheCopy(m, int32(57), v11, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v20+v21)+128)) = uint8(v2)
				F_CatalogTupleUpdate(m, v14, v18+int32(4), v18)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecForceNoForceRowSecurity[0]))
					if v29 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v32 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v31, v32, v32, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_relation_close(m, v14, int32(3))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					} else {
						F_relation_close(m, v14, int32(3))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							F_pfree(m, v18)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg_internal(m, int32(_a_F_ATExecForceNoForceRowSecurity_0), v9)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATExecForceNoForceRowSecurity_1), int32(_a_F_ATExecForceNoForceRowSecurity_2), int32(_a_F_ATExecForceNoForceRowSecurity_3))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_ATRewriteTable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v499 int32
	_ = v499
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v684 int32
	_ = v684
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v895 int32
	_ = v895
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v947 int32
	_ = v947
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1107 int32
	_ = v1107
	var v1112 int32
	_ = v1112
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1233 int32
	_ = v1233
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1244 int32
	_ = v1244
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1281 int32
	_ = v1281
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	v3 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(144)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = F_table_open(m, v31, v3)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v55 = F_CreateExecutorState(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v44 = int32(1)
	v46 = F_GetCurrentCommandId(m, v44)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v38 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(0)
	v50 = v41
	v51 = v3
	v52 = v3
	v53 = v3
	v54 = v41
	goto L3
L8:
	;
	if v38 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v48 = F_GetBulkInsertState(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = v38
	v51 = v44
	v52 = int32(2)
	v53 = v46
	v54 = v48
	goto L3
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v57 == int32(0) {
		v139 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v153 != 0 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v60 <= int32(0) {
		v139 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v65 = int32(0)
	v76 = v3
	goto L16
L16:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v65<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	switch v95 - int32(5) {
	case 0:
		goto L19
	default:
		goto L20
	case 4:
		v122 = v76
		goto L18
	}
L17:
	;
	v139 = v122
	goto L13
L18:
	;
	v124 = v65 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v124 < v125 {
		v65 = v124
		v76 = v122
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v114 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
	v117 = F_expand_generated_columns_in_expr(m, v115, v33, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v102
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_0), v29+int32(128))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_2), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v119 = F_ExecPrepareExpr(m, v117, v55)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+28)) = v119
	v122 = v114
	goto L18
L26:
	;
	goto L17
L27:
	;
	v155 = F_ExecPrepareExpr(m, v153, v55)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v158 = v139
	v159 = int32(0)
	goto L29
L29:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v160 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v158 = int32(1)
	v159 = v155
	goto L29
L31:
	;
	if v51 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v163 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v168 = int32(0)
	goto L34
L34:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v168<<(uint(int32(2))%32))))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v200 = F_ExecInitExpr(m, v198, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v200
	v204 = v168 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v204 < v205 {
		v168 = v204
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_FreeExecutorState(m, v55)
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L239
	}
L39:
	;
	v383 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L65
	}
L40:
	;
	if v294 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L41:
	;
	v326 = int32(1)
	if (v158|v51)&v326 == int32(0) {
		goto L38
	} else {
		goto L59
	}
L42:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v235 != int32(1) {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v238 <= int32(0) {
		goto L41
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	v245 = int32(0)
	v246 = v238
	v258 = v3
	v259 = v3
	goto L47
L47:
	;
	v272 = v35 + int32(20) + v245<<(uint(int32(4))%32)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+11)))
	if v273 != int32(118) {
		v293 = v258
		v294 = v259
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v293|v294 != 0 {
		goto L40
	} else {
		goto L58
	}
L49:
	;
	v296 = v245 + int32(1)
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v296 < v297 {
		v245 = v296
		v246 = v297
		v258 = v293
		v259 = v294
		goto L47
	} else {
		goto L57
	}
L50:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+9)))
	if v276 != 0 {
		v293 = v258
		v294 = v259
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v282 = v35 + v246<<(uint(int32(4))%32) + v245*int32(100)
	v283 = int32(*(*int16)(unsafe.Add(mBase, uint32(v282)+94)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+110)))
	if v284 != int32(118) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v287 = F_lappend_int(m, v258, v283)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v289 = F_lappend_int(m, v259, v283)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v293 = v287
	v294 = v259
	goto L49
L56:
	;
	v293 = v258
	v294 = v289
	goto L49
L57:
	;
	goto L48
L58:
	;
	goto L41
L59:
	;
	v332 = int32(0)
	v369 = v332
	v370 = v332
	v376 = v326
	v377 = v3
	goto L39
L60:
	;
	v369 = v293
	v370 = int32(0)
	v376 = int32(1)
	v377 = v3
	goto L39
L61:
	;
	goto L62
L62:
	;
	v338 = int32(_a_F_ATRewriteTable_4)
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v55)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v341
	v344 = F_palloc0(m, int32(216))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = int32(388)
	v348 = int32(0)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v55)+132))
	F_InitResultRelInfo(m, v344, v33, v348, v348, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v339
	v369 = v293
	v370 = v294
	v376 = v3
	v377 = v344
	goto L39
L65:
	;
	if v51 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v417 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L67:
	;
	if v383 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	if v383 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L70:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v385 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_5), v29+int32(96))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_TransferPredicateLocksToHeapRelation(m, v33)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_6), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	goto L66
L76:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v403 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_7), v29+int32(112))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_8), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L66
L79:
	;
	v420 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	v422 = v417
	goto L81
L81:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v424 = F_table_slot_callbacks(m, v33)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	v422 = v420
	goto L81
L83:
	;
	if v423 != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v439 = int32(0)
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v439 < v440 {
		goto L93
	} else {
		goto L94
	}
L85:
	;
	v426 = F_MakeTupleTableSlot(m, v36, v424)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v435 = F_MakeTupleTableSlot(m, v35, v424)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L92
	}
L88:
	;
	v428 = F_table_slot_callbacks(m, v50)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v430 = F_MakeTupleTableSlot(m, v35, v428)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_ExecStoreAllNullTuple(m, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v437 = v430
	v438 = v426
	goto L84
L92:
	;
	v437 = int32(0)
	v438 = v435
	goto L84
L93:
	;
	v445 = int32(0)
	v446 = v440
	v456 = v439
	goto L96
L94:
	;
	v499 = v439
	goto L95
L95:
	;
	v513 = F_GetLatestSnapshot(m)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L103
	}
L96:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v446<<(uint(int32(4))%32)+v445*int32(100))+111)))
	if v476 == int32(1) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v499 = v483
	goto L95
L98:
	;
	v479 = F_lappend_int(m, v456, v445)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	v482 = v446
	v483 = v456
	goto L100
L100:
	;
	v485 = v445 + int32(1)
	if v485 < v482 {
		v445 = v485
		v446 = v482
		v456 = v483
		goto L96
	} else {
		goto L102
	}
L101:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v482 = v481
	v483 = v479
	goto L100
L102:
	;
	goto L97
L103:
	;
	v515 = F_RegisterSnapshot(m, v513)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v517 = int32(0)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+8))
	v523 = m.T0[v522].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v33, v515, v517, v517, v517, int32(449))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v525 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v528 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v530 = v525
	goto L108
L108:
	;
	v531 = int32(_a_F_ATRewriteTable_4)
	v532 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0]))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v530)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v534
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+36)) = v537
	v540 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[1]))
	if v540 != 0 {
		goto L112
	} else {
		goto L113
	}
L109:
	;
	v530 = v528
	goto L108
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v1157
	F_errmsg(m, int32(_a_F_ATRewriteTable_9), v29)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L236
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L233
	}
L112:
	;
	v542 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATRewriteTable[2])))
	if v542&int32(1) == int32(0) {
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	goto L116
L115:
	;
	goto L114
L116:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v574)+188))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v575)+20))
	v577 = m.T0[v576].(func(*base.Module, int32, int32, int32) int32)(m, v523, int32(1), v438)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	goto L111
L118:
	;
	if v376 != 0 {
		goto L184
	} else {
		goto L185
	}
L119:
	;
	if v577 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v579 <= int32(0) {
		v838 = v438
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v532
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v950)+188))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v951)+12))
	m.T0[v952].(func(*base.Module, int32))(m, v523)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L179
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v838
	if v369 == int32(0) {
		goto L118
	} else {
		goto L162
	}
L124:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v582)))
	v584 = int32(*(*int16)(unsafe.Add(mBase, uint32(v438)+6)))
	if v584 < v583 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	F_slot_getsomeattrs_int(m, v438, v583)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v437)+8))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v588)+12))
	m.T0[v589].(func(*base.Module, int32))(m, v437)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	goto L127
L129:
	;
	v592 = int32(*(*int16)(unsafe.Add(mBase, uint32(v438)+6)))
	v594 = v592 << (uint(int32(2)) % 32)
	if v594 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v438)+16))
	base.MemoryCopy(m, v595, v596, v594)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v598 = int32(*(*int16)(unsafe.Add(mBase, uint32(v438)+6)))
	if v598 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v437)+20))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v438)+20))
	base.MemoryCopy(m, v599, v600, v598)
	goto L135
L134:
	;
	goto L135
L135:
	;
	if v499 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v437)+36)) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v438
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v676 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L137:
	;
	v604 = int32(0)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v605 <= v604 {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v609 = v604
	goto L139
L139:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v437)+20))
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v635+v609<<(uint(int32(2))%32))))
	v641 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v634+v639))) = uint8(v641)
	v644 = v609 + v641
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v644 < v645 {
		v609 = v644
		goto L139
	} else {
		goto L141
	}
L140:
	;
	goto L136
L141:
	;
	goto L140
L142:
	;
	v765 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)))
	v767 = v765 & int32(_a_F_ATRewriteTable_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+4)) = uint16(v767)
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v769)))
	*(*uint16)(unsafe.Add(mBase, uint32(v437)+6)) = uint16(v770)
	goto L152
L143:
	;
	v679 = int32(0)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v680 <= v679 {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v684 = v679
	goto L145
L145:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v676)+12))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709+v684<<(uint(int32(2))%32))))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v713)+12)))
	if v714 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	goto L142
L147:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v713)+8))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v437)+20))
	v719 = int32(*(*int16)(unsafe.Add(mBase, uint32(v713))))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v717)+20))
	v724 = m.T0[v723].(func(*base.Module, int32, int32, int32) int32)(m, v717, v422, v718+v719-int32(1))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v736 = v684 + int32(1)
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	if v736 < v737 {
		v684 = v736
		goto L145
	} else {
		goto L151
	}
L150:
	;
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v727 = int32(*(*int16)(unsafe.Add(mBase, uint32(v713))))
	*(*int32)(unsafe.Add(mBase, uint32(v726+v727<<(uint(int32(2))%32)-int32(4)))) = v724
	goto L149
L151:
	;
	goto L146
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v422)+4)) = v437
	v773 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v773 == int32(0) {
		v838 = v437
		goto L123
	} else {
		goto L153
	}
L153:
	;
	v776 = int32(0)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v773)+4))
	if v777 <= v776 {
		v838 = v437
		goto L123
	} else {
		goto L154
	}
L154:
	;
	v781 = v776
	goto L155
L155:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v773)+12))
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v806+v781<<(uint(int32(2))%32))))
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v810)+12)))
	if v811 == int32(1) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v838 = v437
	goto L123
L157:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v810)+8))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v437)+20))
	v816 = int32(*(*int16)(unsafe.Add(mBase, uint32(v810))))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v814)+20))
	v821 = m.T0[v820].(func(*base.Module, int32, int32, int32) int32)(m, v814, v422, v815+v816-int32(1))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v833 = v781 + int32(1)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v773)+4))
	if v833 < v834 {
		v781 = v833
		goto L155
	} else {
		goto L161
	}
L160:
	;
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v437)+16))
	v824 = int32(*(*int16)(unsafe.Add(mBase, uint32(v810))))
	*(*int32)(unsafe.Add(mBase, uint32(v823+v824<<(uint(int32(2))%32)-int32(4)))) = v821
	goto L159
L161:
	;
	goto L156
L162:
	;
	v865 = int32(0)
	v866 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v866 <= v865 {
		goto L118
	} else {
		goto L163
	}
L163:
	;
	v870 = v865
	goto L164
L164:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v369)+12))
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v895+v870<<(uint(int32(2))%32))))
	v900 = int32(*(*int16)(unsafe.Add(mBase, uint32(v838)+6)))
	if v900 < v899 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L174
	}
L166:
	;
	F_slot_getsomeattrs_int(m, v838, v899)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v838)+20))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v904+v899-int32(1)))))
	if v908 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L168
L170:
	;
	v912 = v870 + int32(1)
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v369)+4))
	if v913 <= v912 {
		goto L118
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	goto L165
L173:
	;
	v870 = v912
	goto L164
L174:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v924 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v923 + v924
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v35 + v915<<(uint(v924)%32) + v899*int32(100) - int32(76)
	F_errmsg(m, int32(_a_F_ATRewriteTable_11), v29+int32(80))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errtablecol(m, v33, v899)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_12), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_UnregisterSnapshot(m, v515)
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	F_ExecDropSingleTupleTableSlot(m, v438)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	if v437 == int32(0) {
		goto L38
	} else {
		goto L182
	}
L182:
	;
	F_ExecDropSingleTupleTableSlot(m, v437)
	mBase = m.M
	v962 = m.ExcPending
	if v962 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	goto L38
L184:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1027 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L185:
	;
	v989 = F_ExecRelGenVirtualNotNull(m, v377, v838, v55, v370)
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	if v989 == int32(0) {
		goto L184
	} else {
		goto L187
	}
L187:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1002 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1001 + v1002
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v35 + v993<<(uint(v1002)%32) + v989*int32(100) - int32(76)
	F_errmsg(m, int32(_a_F_ATRewriteTable_11), v29-int32(-64))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	F_errtablecol(m, v33, v989)
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_13), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L193:
	;
	if v159 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L194:
	;
	v1030 = int32(0)
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	if v1031 <= v1030 {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v1035 = v1030
	goto L196
L196:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+12))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1060+v1035<<(uint(int32(2))%32))))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	switch v1065 - int32(1) {
	case 0, 8:
		goto L198
	default:
		goto L199
	case 4:
		goto L200
	}
L197:
	;
	goto L193
L198:
	;
	v1114 = v1035 + int32(1)
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1027)+4))
	if v1114 < v1115 {
		v1035 = v1114
		goto L196
	} else {
		goto L211
	}
L199:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1100 = m.ExcPending
	if v1100 != 0 {
		goto L1
	} else {
		goto L208
	}
L200:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+28))
	v1069 = F_ExecCheck(m, v1068, v422)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	if v1069 != 0 {
		goto L198
	} else {
		goto L202
	}
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1064)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1079
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1078 + int32(4)
	F_errmsg(m, int32(_a_F_ATRewriteTable_14), v29+int32(48))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1064)))
	F_errtableconstraint(m, v33, v1089)
	mBase = m.M
	v1091 = m.ExcPending
	if v1091 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_15), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1101
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_0), v29+int32(32))
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_16), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L211:
	;
	goto L197
L212:
	;
	if v51 != 0 {
		goto L222
	} else {
		goto L223
	}
L213:
	;
	v1145 = F_ExecCheck(m, v159, v422)
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	if v1145 != 0 {
		goto L212
	} else {
		goto L215
	}
L215:
	;
	v1147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1157 = v1155 + int32(4)
	if v1147 == int32(1) {
		goto L110
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1157
	F_errmsg(m, int32(_a_F_ATRewriteTable_17), v29+int32(16))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	F_errtable(m, v33)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_18), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L222:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1173)+80))
	m.T0[v1174].(func(*base.Module, int32, int32, int32, int32, int32))(m, v50, v838, v53, v52, v54)
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v422)+20))
	F_MemoryContextReset(m, v1177)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[3]))
	if v1181 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1184)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+36)) = v1185
	v1188 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[1]))
	if v1188 == int32(0) {
		goto L116
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATRewriteTable[2])))
	if v1192&int32(1) != 0 {
		goto L116
	} else {
		goto L232
	}
L232:
	;
	goto L117
L233:
	;
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_19), int32(0))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_20), int32(1034), int32(_a_F_ATRewriteTable_21))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errtable(m, v33)
	mBase = m.M
	v1239 = m.ExcPending
	if v1239 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_22), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	F_relation_close(m, v33, int32(0))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	if v51 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	F_FreeBulkInsertState(m, v54)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	m.G0 = v29 + int32(144)
	return
L244:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	if v1278 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	F_relation_close(m, v50, int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L1
	} else {
		goto L249
	}
L246:
	;
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1278)+108))
	if v1281 == int32(0) {
		goto L245
	} else {
		goto L247
	}
L247:
	;
	m.T0[v1281].(func(*base.Module, int32, int32))(m, v50, v52)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	goto L245
L249:
	;
	goto L243
}
