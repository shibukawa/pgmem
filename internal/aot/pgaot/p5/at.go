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
	v283 = F_palloc0(m, int32(144))
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
	var v27 int32
	_ = v27
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
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
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
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
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v405 int64
	_ = v405
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int64
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
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
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int64
	_ = v463
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v651 int32
	_ = v651
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v727 int32
	_ = v727
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v826 int32
	_ = v826
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v856 int32
	_ = v856
	var v861 int32
	_ = v861
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v992 int32
	_ = v992
	v27 = m.G0
	v29 = v27 - int32(128)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+28)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
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
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L216
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L212
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L208
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L205
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L201
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L197
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L1
	} else {
		goto L190
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L186
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L182
	}
L12:
	;
	v47 = l3 + int32(48)
	v50 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	F_ATSimplePermissions(m, v38, l3, int32(289))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+131)))
	if v43 == int32(1) {
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
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v34)+16)))
	if v52 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	m.G0 = v29 + int32(128)
	return
L20:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v130 = F_check_for_column_name_collision(m, l3, v127, v33&int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v56 = F_SearchSysCacheCopyAttName(m, v31, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v56 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	F_typenameTypeIdAndMod(m, int32(0), v63, v29+int32(116), v29+int32(104))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v71 = v60 + v61
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+68))
	if v70 != v72 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v29)+104))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	if v74 != v75 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v78 = F_GetColumnDefCollation(m, int32(0), v34, v70)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v71)+96))
	if v78 != v80 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v71)+94)))
	v84 = v82 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+94)) = uint16(v84)
	if base.I32_extend16_s(v84) != v84 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	F_CatalogTupleUpdate(m, v50, v56+int32(4), v56)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v56)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v96 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = v98 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_0), v29+int32(32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_relation_close(m, v50, int32(3))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_2), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v121
	v124 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddColumn[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v124
	goto L19
L40:
	;
	if v130 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_relation_close(m, v50, int32(3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v143 = int32(0)
	if l6|base.B2i32(l9 == v143) == v143 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v138
	v141 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecAddColumn[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v141
	goto L19
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v149 = F_ATParseTransformCmd(m, l2, l3, v148, l5, l8, l9)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v153 = v34
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
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v149
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	v153 = v152
	goto L47
L49:
	;
	v168 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L55
	}
L50:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+36)))
	if v156 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+119)))
	if v160 == int32(112) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v164 = F_find_inheritance_children(m, v31, int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v164 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	goto L49
L55:
	;
	v173 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v31), int64(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v173 == int32(0) {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+22)))
	v179 = v177 + v178
	v180 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179)+120)))
	if int32(1600) <= v180 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+119)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v153
	v189 = F_list_make1_impl(m, int32(1), v29+int32(28))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v191 = F_BuildDescForRelation(m, v189)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v196 = v191 + v193<<(uint(int32(3))%32)
	v198 = v180 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+102)) = uint16(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)+96))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196)+124))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+24)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v203
	v211 = F_list_make1_impl(m, int32(480), v29+int32(24))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+118)))
	if v215 == int32(118) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v218 = int32(8)
	goto L64
L63:
	;
	v218 = int32(0)
	goto L64
L64:
	;
	F_CheckAttributeType(m, v196+int32(32), v200, v201, v211, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v221 = int32(0)
	F_InsertPgAttributeTuples(m, v50, v191, v31, v221, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_relation_close(m, v50, int32(3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v179)+120)) = uint16(v198)
	F_CatalogTupleUpdate(m, v168, v173+int32(4), v173)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_pfree(m, v173)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAddColumn[2]))
	if v236 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v31, v198, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	F_relation_close(m, v168, int32(3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	if v246 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v248 = F_palloc(m, int32(12))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	switch v183 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L85
	default:
		goto L84
	}
L79:
	;
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+102)))
	*(*uint16)(unsafe.Add(mBase, uint32(v248))) = uint16(v250)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v153)+28))
	v253 = F_copyObjectImpl(m, v252)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v248)+4)) = v253
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v248)+8)) = uint8(v256)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v248
	v263 = F_list_make1_impl(m, int32(1), v29+int32(20))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v265 = int32(0)
	v270 = F_AddRelationNewConstraints(m, l3, v263, v265, v265, int32(1), v265, v265)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L78
L84:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v196)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+108)) = v524
	*(*int32)(unsafe.Add(mBase, uint32(v29)+104)) = int32(1247)
	v535 = v29 + int32(116)
	v537 = v29 + int32(104)
	F_recordDependencyOn(m, v535, v537, int32(110))
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L1
	} else {
		goto L143
	}
L85:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+36)))
	if v277 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)))
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+19)))
	v513 = v511 | v512
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)) = uint8(v513)
	goto L84
L87:
	;
	v350 = F_expression_planner(m, v345)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L108
	}
L88:
	;
	v279 = F_palloc0(m, int32(12))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v295 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196)+102)))
	v296 = F_build_column_default(m, l3, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = int32(59)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v153)+40))
	v284 = int32(0)
	v288 = F_RangeVarGetRelidExtended(m, v283, v284, v284, v284, v284)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+4)) = v288
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v196)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+8)) = v291
	v293 = F_DomainHasConstraints(m, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v345 = v279
	v348 = v293
	goto L87
L94:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v196)+96))
	v299 = F_DomainHasConstraints(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v301 = int32(0)
	if v296|base.B2i32(v299 == v301) == v301 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v307 = v196 + int32(104)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+116)) = v308
	v311 = v196 + int32(96)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v315 = F_getBaseTypeAndTypmod(m, v312, v29+int32(116))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v296 == int32(0) {
		goto L86
	} else {
		goto L107
	}
L99:
	;
	v317 = F_get_typcollation(m, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v29)+116))
	v321 = F_makeNullConst(m, v315, v320, v317)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v307)))
	v328 = F_coerce_to_target_type(m, int32(0), v321, v315, v323, v324, int32(1), int32(2), int32(-1))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	if v328 != 0 {
		v345 = v328
		v348 = v299
		goto L87
	} else {
		goto L103
	}
L103:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_4), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_5), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	v345 = v296
	v348 = v299
	goto L87
L108:
	;
	v353 = F_palloc0(m, int32(16))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+102)))
	*(*int32)(unsafe.Add(mBase, uint32(v353)+4)) = v350
	*(*uint16)(unsafe.Add(mBase, uint32(v353))) = uint16(v355)
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v353)+12)) = uint8(base.B2i32(v358 != int32(0)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v363 = F_lappend(m, v362, v353)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+68)) = v363
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+119)))
	if v367 != int32(114) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	F_FreeExecutorState(m, v378)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L142
	}
L112:
	;
	if v493&int32(255) == int32(118) {
		goto L86
	} else {
		goto L141
	}
L113:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+44)))
	v493 = v492
	goto L112
L114:
	;
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+44)))
	if (base.B2i32(v370 != int32(0))|v348)&int32(1) != 0 {
		v493 = v370
		goto L112
	} else {
		goto L115
	}
L115:
	;
	v376 = F_contain_volatile_functions(m, v350)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	if v376 != 0 {
		goto L113
	} else {
		goto L117
	}
L117:
	;
	v378 = F_CreateExecutorState(m)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v380 = F_ExecPrepareExpr(m, v350, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v378)+152))
	if v382 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v385 = v382
	goto L122
L121:
	;
	v383 = F_MakePerTupleExprContext(m, v378)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v380)+24))
	v389 = m.T0[v388].(func(*base.Module, int32, int32, int32) int64)(m, v380, v385, v29+int32(116))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L124
	}
L123:
	;
	v385 = v383
	goto L122
L124:
	;
	v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+116)))
	if v391 != 0 {
		goto L111
	} else {
		goto L125
	}
L125:
	;
	v392 = int32(*(*int16)(unsafe.Add(mBase, uint32(v196)+102)))
	v393 = m.G0
	v395 = v393 - int32(288)
	m.G0 = v395
	*(*int64)(unsafe.Add(mBase, uint32(v395)+280)) = v389
	v400 = int32(0)
	base.MemoryFill(m, v395+int32(80), v400, int32(192))
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+72)) = uint8(v400)
	v405 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v395)+64)) = v405
	*(*int64)(unsafe.Add(mBase, uint32(v395)+56)) = v405
	*(*int64)(unsafe.Add(mBase, uint32(v395)+48)) = v405
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+40)) = uint8(v400)
	*(*int64)(unsafe.Add(mBase, uint32(v395)+32)) = v405
	*(*int64)(unsafe.Add(mBase, uint32(v395)+24)) = v405
	*(*int64)(unsafe.Add(mBase, uint32(v395)+16)) = v405
	v421 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v424 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l3)+56)))
	v426 = F_SearchSysCache2(m, int32(7), v424, base.I64_extend_i32_s(v392))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v426 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v426)+16))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448)+22)))
	v450 = v448 + v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+68))
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v450)+72)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+82)))
	v454 = int32(*(*int8)(unsafe.Add(mBase, uint32(v450)+83)))
	v455 = F_construct_array(m, v395+int32(280), int32(1), v451, v452, v453, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L134
	}
L131:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v395)+4)) = v434
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v392
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_6), v395)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_7), int32(2069), int32(_a_F_ATExecAddColumn_8))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v395)+184)) = int64(1)
	v459 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+29)) = uint8(v459)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+40)) = uint8(v459)
	v463 = base.I64_extend_i32_u(v455)
	*(*int64)(unsafe.Add(mBase, uint32(v395)+280)) = v463
	*(*int64)(unsafe.Add(mBase, uint32(v395)+272)) = v463
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v421)+52))
	v473 = F_heap_modify_tuple(m, v426, v466, v395+int32(80), v395+int32(48), v395+int32(16))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_CatalogTupleUpdate(m, v421, v473+int32(4), v473)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_ReleaseCatCache(m, v426)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_relation_close(m, v421, int32(3))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	m.G0 = v395 + int32(288)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_FreeExecutorState(m, v378)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L84
L141:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v498 | int32(2)
	goto L86
L142:
	;
	goto L86
L143:
	;
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v196)+124))
	v542 = int32(0)
	if base.B2i32(v541 == v542)|base.B2i32(v541 == int32(100)) == v542 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+124)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v29)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+108)) = v541
	*(*int32)(unsafe.Add(mBase, uint32(v29)+104)) = int32(3456)
	F_recordDependencyOn(m, v535, v537, int32(110))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L147
	}
L145:
	;
	goto L146
L146:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v563 = F_find_inheritance_children(m, v562, l7)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L148
	}
L147:
	;
	goto L146
L148:
	;
	if v563 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v565 = l5
	goto L151
L150:
	;
	v565 = int32(1)
	goto L151
L151:
	;
	if v565 == int32(0) {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l6 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v571 = F_copyObjectImpl(m, v568)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L156
	}
L154:
	;
	v579 = v568
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+104)) = v579
	if v563 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v571)+20))
	v574 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v573)+18)) = uint8(v574)
	v576 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v573)+16)) = uint16(v576)
	v579 = v571
	goto L155
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L19
L158:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v583 <= int32(0) {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v589 = int32(0)
	goto L160
L160:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v563)+12))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v613+v589<<(uint(int32(2))%32))))
	v619 = F_table_open(m, v617, int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L162
	}
L161:
	;
	goto L157
L162:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v619)+48))
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621)+118)))
	if v622 == int32(116) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v619)+24)))
	if v625 == int32(0) {
		goto L3
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	F_CheckTableNotInUse(m, v619, int32(_a_F_ATExecAddColumn_9))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L167
	}
L166:
	;
	goto L165
L167:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v619)+56))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v632 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	F_ATExecAddColumn(m, v29+int32(116), l1, v727, v619, v29+int32(104), l5, int32(1), l7, l8, l9)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L179
	}
L169:
	;
	v702 = F_palloc0(m, int32(144))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L176
	}
L170:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v632)+4))
	if v635 <= int32(0) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v632)+12))
	v651 = int32(0)
	goto L172
L172:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v638+v651<<(uint(int32(2))%32))))
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v669)))
	if v670 == v631 {
		v727 = v669
		goto L168
	} else {
		goto L174
	}
L173:
	;
	goto L169
L174:
	;
	v673 = v651 + int32(1)
	if v635 != v673 {
		v651 = v673
		goto L172
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v702)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v702))) = v631
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v619)+48))
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v702)+4)) = uint8(v708)
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v619)+52))
	v711 = F_CreateTupleDescCopyConstr(m, v710)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v702)+8)) = v711
	*(*int64)(unsafe.Add(mBase, uint32(v702)+88)) = int64(0)
	v716 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v702)+84)) = uint8(v716)
	v718 = int32(_a_F_ATExecAddColumn_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v702)+96)) = uint16(v718)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v721 = F_lappend(m, v720, v702)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v721
	v727 = v702
	goto L168
L179:
	;
	F_relation_close(m, v619, int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v761 = v589 + int32(1)
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v563)+4))
	if v761 < v762 {
		v589 = v761
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
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_11), int32(0))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_12), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
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
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v847
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v846 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_13), v29+int32(80))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_14), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
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
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v870
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v869 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_15), v29-int32(-64))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v880 = F_get_collation_name(m, v78)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v71)+96))
	v883 = F_get_collation_name(m, v882)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v883
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v880
	v890 = F_errdetail(m, int32(_a_F_ATExecAddColumn_16), v29+int32(48))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_17), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
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
	v903 = m.ExcPending
	if v903 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_18), int32(0))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_19), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
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
	v919 = m.ExcPending
	if v919 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_20), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_21), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
	F_errmsg_internal(m, int32(_a_F_ATExecAddColumn_22), v29)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_23), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
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
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(1600)
	F_errmsg(m, int32(_a_F_ATExecAddColumn_24), v29+int32(16))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_25), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
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
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_26), int32(0))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_27), int32(_a_F_ATExecAddColumn_3))
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
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
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_ATExecAddColumn_28), int32(0))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_ATExecAddColumn_1), int32(_a_F_ATExecAddColumn_29), int32(_a_F_ATExecAddColumn_30))
	mBase = m.M
	v992 = m.ExcPending
	if v992 != 0 {
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
	var v41 int32
	_ = v41
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v187 int32
	_ = v187
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
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v219 int32
	_ = v219
	v18 = m.G0
	v20 = v18 + int32(-64)
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
	m.G0 = v20 - int32(-64)
	return v219
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+74)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v33 == v34 {
		v219 = int32(0)
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
	v41 = v18 + int32(-56)
	v45 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v28))))
	F_ScanKeyInit(m, v41, int32(11), int32(3), int32(184), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v48 = int32(1)
	v53 = F_systable_beginscan(m, l2, int32(2699), v48, int32(0), v48, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v55 = F_systable_getnext(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v55 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v65 = v55
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_systable_endscan(m, v53)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L33
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+22)))
	v76 = v74 + v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v77 != v78 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v81 = F_list_append_unique_oid(m, v80, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v76)+76))
	v86 = v84 - int32(1644)
	v93 = int32(0)
	if base.B2i32(base.Ui32(int32(11)) < base.Ui32(v86))|base.B2i32(int32(1)<<(uint(v86)%32)&int32(3075) == v93) == v93 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v81
	goto L19
L21:
	;
	v98 = F_heap_copytuple(m, v65)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v122 = F_systable_getnext(m, v53)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+16))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	v102 = v100 + v101
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+97)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v102)+96)) = uint8(v39)
	F_CatalogTupleUpdate(m, l2, v98+int32(4), v98)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecAlterConstrDeferrability[0]))
	if v110 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v113 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v112, v113, v113, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v98)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
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
	if v122 != 0 {
		v65 = v122
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
		v219 = v48
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+119)))
	if v146 != int32(112) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v149 = F_get_rel_relkind(m, v29)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v154 = v18 + int32(-56)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158)+22)))
	v161 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v158+v159))))
	F_ScanKeyInit(m, v154, int32(12), int32(3), int32(184), v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	if v149 != int32(112) {
		v219 = v48
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v165 = int32(1)
	v168 = F_systable_beginscan(m, l1, int32(2579), v165, int32(0), v165, v154)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	v187 = F_systable_getnext(m, v168)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	F_systable_endscan(m, v168)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L51
	}
L44:
	;
	if v187 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v187)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v190)+80))
	v193 = F_table_open(m, v192, l7)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
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
	v196 = F_ATExecAlterConstrDeferrability(m, l0, l1, l2, v193, v187, int32(1), l6, l7)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_relation_close(m, v193, int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v219 = v48
	goto L3
}
func F_ATExecChangeOwner(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int64
	_ = v216
	var v218 int32
	_ = v218
	var v236 int32
	_ = v236
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
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
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v462 int32
	_ = v462
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v503 int32
	_ = v503
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	v16 = m.G0
	v18 = v16 - int32(848)
	m.G0 = v18
	v20 = F_relation_open(m, l0, l3)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = base.I64_extend_i32_u(l0)
	v28 = F_SearchSysCache1(m, int32(57), v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L9
	}
L4:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[0]))
	if v633 != 0 {
		goto L166
	} else {
		goto L167
	}
L5:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v32)+72))
	if v435 != 0 {
		goto L121
	} else {
		goto L122
	}
L6:
	;
	if v385 == int32(73) {
		goto L5
	} else {
		goto L119
	}
L7:
	;
	if int32(1)<<(uint(v387)%32)&int32(_a_F_ATExecChangeOwner_0) != 0 {
		goto L5
	} else {
		goto L118
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L113
	}
L9:
	;
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+22)))
	v32 = v30 + v31
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
	switch v33 - int32(73) {
	case 0:
		goto L17
	default:
		goto L8
	case 10:
		goto L16
	case 26:
		goto L15
	case 29, 36, 39, 41, 45:
		v164 = l1
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
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L110
	}
L13:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if v167 == v164 {
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
		v164 = l1
		goto L13
	} else {
		goto L48
	}
L16:
	;
	if l2 != 0 {
		v164 = l1
		goto L13
	} else {
		goto L34
	}
L17:
	;
	if l2 != 0 {
		v164 = l1
		goto L13
	} else {
		goto L28
	}
L18:
	;
	if l2 != 0 {
		v164 = l1
		goto L13
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if v36 == l1 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	v164 = v64
	goto L13
L21:
	;
	v40 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v40 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v32 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_1), v18+int32(32))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errhint(m, int32(_a_F_ATExecChangeOwner_2), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_4), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
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
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v32 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_1), v18+int32(48))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(_a_F_ATExecChangeOwner_2), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_6), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	if v89 == l1 {
		v164 = l1
		goto L13
	} else {
		goto L35
	}
L35:
	;
	v93 = v18 + int32(224)
	v95 = v18 + int32(576)
	v96 = F_sequenceIsOwned(m, l0, int32(97), v93, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v96 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v101 = F_sequenceIsOwned(m, l0, int32(105), v93, v95)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
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
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	if v101 == int32(0) {
		v164 = l1
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
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v113 = v32 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v113
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_7), v18+int32(80))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v18)+224))
	v121 = F_get_rel_name(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v113
	v128 = F_errdetail(m, int32(_a_F_ATExecChangeOwner_8), v18-int32(-64))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_9), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
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
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v32 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_10), v18+int32(112))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = int32(_a_F_ATExecChangeOwner_11)
	F_errhint(m, int32(_a_F_ATExecChangeOwner_12), v18+int32(96))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_13), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v164 = l1
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
	v216 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v216
	v218 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+208)) = uint16(v218)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+200)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+192)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+184)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+176)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v216
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v216
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+160)) = uint16(v218)
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+133)) = uint8(v236)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+264)) = base.I64_extend_i32_u(v164)
	v244 = F_SysCacheGetAttr(m, int32(57), v28, int32(32), v18+int32(127))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	v169 = F_superuser(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v169 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+68))
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[1]))
	v175 = F_object_ownercheck(m, int32(1259), l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v175 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v180 = F_get_rel_relkind(m, l0)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecChangeOwner[1]))
	F_check_can_set_role(m, v200, v164)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L73
	}
L64:
	;
	switch v180 - int32(73) {
	case 0, 32:
		v191 = int32(20)
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
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	F_aclcheck_error(m, int32(2), v193, v194+int32(4))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	v193 = v191
	goto L65
L67:
	;
	v191 = int32(42)
	goto L66
L68:
	;
	v193 = int32(18)
	goto L65
L69:
	;
	v193 = int32(23)
	goto L65
L70:
	;
	v193 = int32(52)
	goto L65
L71:
	;
	v193 = int32(38)
	goto L65
L72:
	;
	goto L63
L73:
	;
	v205 = F_object_aclcheck(m, int32(2615), v171, v164, int64(512))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v205 == int32(0) {
		goto L56
	} else {
		goto L75
	}
L75:
	;
	v210 = F_get_namespace_name(m, v171)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_aclcheck_error(m, v205, int32(37), v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L56
L78:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+127)))
	if v246 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v250 = F_pg_detoast_datum(m, base.I32_wrap_i64(v244))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v267 = F_heap_modify_tuple(m, v28, v260, v18+int32(224), v18+int32(176), v18+int32(128))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	v253 = F_aclnewowner(m, v250, v252, v164)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+159)) = uint8(v255)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+472)) = base.I64_extend_i32_u(v253)
	goto L81
L84:
	;
	F_CatalogTupleUpdate(m, v24, v267+int32(4), v267)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_pfree(m, v267)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v32)+80))
	v278 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v281 = v18 + int32(784)
	F_ScanKeyInit(m, v281, int32(1), int32(3), int32(184), v27)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v288 = int32(1)
	v291 = F_systable_beginscan(m, v278, int32(2659), v288, int32(0), v288, v281)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v293 = F_systable_getnext(m, v291)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v293 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v297 = v293
	goto L94
L92:
	;
	goto L93
L93:
	;
	F_systable_endscan(m, v291)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L107
	}
L94:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v297)+16))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+22)))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v311)+91)))
	if v313 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	v363 = F_systable_getnext(m, v291)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L105
	}
L97:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v318 = F_heap_getattr_6(m, v297, int32(22), v315, v18+int32(511))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+511)))
	if v320 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+568)) = uint8(v321)
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+560)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v18)+552)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v18)+544)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v18)+512)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v18)+520)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v18)+528)) = v323
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+536)) = uint8(v321)
	v338 = F_pg_detoast_datum(m, base.I32_wrap_i64(v318))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v340 = F_aclnewowner(m, v338, v275, v164)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	v342 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+533)) = uint8(v342)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+744)) = base.I64_extend_i32_u(v340)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v278)+52))
	v353 = F_heap_modify_tuple(m, v297, v346, v18+int32(576), v18+int32(544), v18+int32(512))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_CatalogTupleUpdate(m, v278, v353+int32(4), v353)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_pfree(m, v353)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	if v363 != 0 {
		v297 = v363
		goto L94
	} else {
		goto L106
	}
L106:
	;
	goto L95
L107:
	;
	F_relation_close(m, v278, int32(3))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
	v387 = v385 - int32(99)
	if base.Ui32(v387) <= base.Ui32(int32(17)) {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	goto L6
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_ATExecChangeOwner_14), v18)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_15), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
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
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v32 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecChangeOwner_16), v18+int32(16))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v418 = int32(*(*int8)(unsafe.Add(mBase, uint32(v32)+119)))
	F_errdetail_relkind_not_supported(m, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ATExecChangeOwner_3), int32(_a_F_ATExecChangeOwner_17), int32(_a_F_ATExecChangeOwner_5))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
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
	F_changeDependencyOnOwner(m, int32(1259), l0, v164)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L5
L121:
	;
	F_AlterTypeOwnerInternal(m, v435, v164)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+119)))
	v440 = v438 - int32(109)
	v447 = int32(0)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v440))|base.B2i32(int32(1)<<(uint(v440)%32)&int32(169) == v447) == v447 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v452 = F_RelationGetIndexList(m, v20)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L129
	}
L126:
	;
	goto L127
L127:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	if v519 != 0 {
		goto L137
	} else {
		goto L138
	}
L128:
	;
	F_list_free(m, v452)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L136
	}
L129:
	;
	if v452 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v456 <= int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v462 = int32(0)
	goto L132
L132:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v452)+12))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475+v462<<(uint(int32(2))%32))))
	F_ATExecChangeOwner(m, v479, v164, int32(1), l3)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L128
L134:
	;
	v484 = v462 + int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	if v484 < v485 {
		v462 = v484
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
	F_ATExecChangeOwner(m, v519, v164, int32(1), l3)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v525 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	v528 = v18 + int32(576)
	F_ScanKeyInit(m, v528, int32(4), int32(3), int32(184), int64(1259))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_ScanKeyInit(m, v18+int32(632), int32(5), int32(3), int32(184), v27)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v546 = F_systable_beginscan(m, v525, int32(2674), int32(1), int32(0), int32(2), v528)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v548 = F_systable_getnext(m, v546)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	if v548 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v552 = v548
	goto L149
L147:
	;
	goto L148
L148:
	;
	F_systable_endscan(m, v546)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L164
	}
L149:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v552)+16))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+22)))
	v567 = v565 + v566
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+20))
	if v568 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	v595 = F_systable_getnext(m, v546)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L162
	}
L152:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v567)))
	if v571 != int32(1259) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v567)+8))
	if v574 != 0 {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567)+24)))
	switch v575 - int32(97) {
	case 0, 8:
		goto L155
	default:
		goto L151
	}
L155:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	v579 = F_relation_open(m, v578, l3)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v579)+48))
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581)+119)))
	if v582 == int32(83) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	F_ATExecChangeOwner(m, v585, v164, int32(1), l3)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	v590 = l3
	goto L159
L159:
	;
	F_relation_close(m, v579, v590)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	v590 = int32(0)
	goto L159
L161:
	;
	goto L151
L162:
	;
	if v595 != 0 {
		v552 = v595
		goto L149
	} else {
		goto L163
	}
L163:
	;
	goto L150
L164:
	;
	F_relation_close(m, v525, int32(1))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L4
L166:
	;
	v635 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v635, v635, v635)
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	F_relation_close(m, v24, int32(3))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_relation_close(m, v20, int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	m.G0 = v18 + int32(848)
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
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
		v19 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(v11), int64(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				*(*uint8)(unsafe.Add(mBase, uint32(v21+v22)+128)) = uint8(v2)
				F_CatalogTupleUpdate(m, v14, v19+int32(4), v19)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecForceNoForceRowSecurity[0]))
					if v30 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v33 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v32, v33, v33, v33)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_relation_close(m, v14, int32(3))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_pfree(m, v19)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
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
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_pfree(m, v19)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
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
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg_internal(m, int32(_a_F_ATExecForceNoForceRowSecurity_0), v9)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATExecForceNoForceRowSecurity_1), int32(_a_F_ATExecForceNoForceRowSecurity_2), int32(_a_F_ATExecForceNoForceRowSecurity_3))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
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
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v335 int32
	_ = v335
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
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
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
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
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v509 int32
	_ = v509
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int64
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v803 int32
	_ = v803
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v842 int64
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1049 int32
	_ = v1049
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1165 int32
	_ = v1165
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1244 int32
	_ = v1244
	var v1248 int32
	_ = v1248
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1317 int32
	_ = v1317
	var v1322 int32
	_ = v1322
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(144)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = F_table_open(m, v32, v3)
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v56 = F_CreateExecutorState(m)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L4:
	;
	v45 = int32(1)
	v47 = F_GetCurrentCommandId(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	v39 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v42 = int32(0)
	v51 = v42
	v52 = v3
	v53 = v3
	v54 = v3
	v55 = v42
	goto L3
L8:
	;
	if v39 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v49 = F_GetBulkInsertState(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v51 = v39
	v52 = v45
	v53 = int32(2)
	v54 = v47
	v55 = v49
	goto L3
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v58 == int32(0) {
		v140 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v156 != 0 {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v61 <= int32(0) {
		v140 = v3
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v66 = int32(0)
	v76 = v3
	goto L16
L16:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v66<<(uint(int32(2))%32))))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	switch v97 - int32(5) {
	case 0:
		goto L19
	default:
		goto L20
	case 4:
		v124 = v76
		goto L18
	}
L17:
	;
	v140 = v124
	goto L13
L18:
	;
	v126 = v66 + int32(1)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v126 < v127 {
		v66 = v126
		v76 = v124
		goto L16
	} else {
		goto L26
	}
L19:
	;
	v116 = int32(1)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	v119 = F_expand_generated_columns_in_expr(m, v117, v34, v116)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+128)) = v104
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_0), v30+int32(128))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_2), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	v121 = F_ExecPrepareExpr(m, v119, v56)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+28)) = v121
	v124 = v116
	goto L18
L26:
	;
	goto L17
L27:
	;
	v158 = F_ExecPrepareExpr(m, v156, v56)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v161 = v140
	v162 = int32(0)
	goto L29
L29:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v163 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v161 = int32(1)
	v162 = v158
	goto L29
L31:
	;
	if v52 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v166 <= int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v171 = int32(0)
	goto L34
L34:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v171<<(uint(int32(2))%32))))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	v204 = F_ExecInitExpr(m, v202, int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L31
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v204
	v208 = v171 + int32(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v208 < v209 {
		v171 = v208
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L247
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v1171
	F_errmsg(m, int32(_a_F_ATRewriteTable_4), v30)
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L1
	} else {
		goto L244
	}
L40:
	;
	F_FreeExecutorState(m, v56)
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L1
	} else {
		goto L233
	}
L41:
	;
	v393 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L67
	}
L42:
	;
	if v302 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L43:
	;
	v335 = int32(1)
	if (v161|v52)&v335 == int32(0) {
		goto L40
	} else {
		goto L61
	}
L44:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v240 != int32(1) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v243 <= int32(0) {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v250 = int32(0)
	v251 = v243
	v263 = v3
	v265 = v3
	goto L49
L49:
	;
	v278 = v36 + int32(28) + v250<<(uint(int32(3))%32)
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+7)))
	if v279 != int32(118) {
		v301 = v263
		v302 = v265
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v301|v302 != 0 {
		goto L42
	} else {
		goto L60
	}
L51:
	;
	v304 = v250 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v304 < v305 {
		v250 = v304
		v251 = v305
		v263 = v301
		v265 = v302
		goto L49
	} else {
		goto L59
	}
L52:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+6)))
	if v282&int32(4) != 0 {
		v301 = v263
		v302 = v265
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v290 = v36 + v251<<(uint(int32(3))%32) + v250*int32(100)
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v290)+102)))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+118)))
	if v292 != int32(118) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v295 = F_lappend_int(m, v263, v291)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v297 = F_lappend_int(m, v265, v291)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v301 = v295
	v302 = v265
	goto L51
L58:
	;
	v301 = v263
	v302 = v297
	goto L51
L59:
	;
	goto L50
L60:
	;
	goto L43
L61:
	;
	v341 = int32(0)
	v378 = v341
	v380 = v341
	v385 = v335
	v386 = v3
	goto L41
L62:
	;
	v378 = v301
	v380 = int32(0)
	v385 = int32(1)
	v386 = v3
	goto L41
L63:
	;
	goto L64
L64:
	;
	v347 = int32(_a_F_ATRewriteTable_5)
	v348 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0]))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v56)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v350
	v353 = F_palloc0(m, int32(216))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v353))) = int32(394)
	v357 = int32(0)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v56)+132))
	F_InitResultRelInfo(m, v353, v34, v357, v357, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v348
	v378 = v301
	v380 = v302
	v385 = v3
	v386 = v353
	goto L41
L67:
	;
	if v52 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v56)+152))
	if v427 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L69:
	;
	if v393 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if v393 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L72:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+96)) = v395 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_6), v30+int32(96))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_TransferPredicateLocksToHeapRelation(m, v34)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_7), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L68
L78:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+112)) = v413 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_8), v30+int32(112))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_9), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L68
L81:
	;
	v430 = F_MakePerTupleExprContext(m, v56)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v432 = v427
	goto L83
L83:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v434 = F_table_slot_callbacks(m, v34)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v432 = v430
	goto L83
L85:
	;
	if v433 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v449 = int32(0)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v449 < v450 {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v436 = F_MakeSingleTupleTableSlot(m, v37, v434)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v445 = F_MakeSingleTupleTableSlot(m, v36, v434)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L94
	}
L90:
	;
	v438 = F_table_slot_callbacks(m, v51)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v440 = F_MakeSingleTupleTableSlot(m, v36, v438)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_ExecStoreAllNullTuple(m, v440)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v447 = v440
	v448 = v436
	goto L86
L94:
	;
	v447 = int32(0)
	v448 = v445
	goto L86
L95:
	;
	v455 = int32(0)
	v457 = v450
	v465 = v449
	goto L98
L96:
	;
	v509 = v449
	goto L97
L97:
	;
	v525 = F_GetLatestSnapshot(m)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L1
	} else {
		goto L105
	}
L98:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36+v457<<(uint(int32(3))%32)+v455*int32(100))+119)))
	if v487 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v509 = v494
	goto L97
L100:
	;
	v490 = F_lappend_int(m, v465, v455)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v493 = v457
	v494 = v465
	goto L102
L102:
	;
	v496 = v455 + int32(1)
	if v496 < v493 {
		v455 = v496
		v457 = v493
		v465 = v494
		goto L98
	} else {
		goto L104
	}
L103:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v493 = v492
	v494 = v490
	goto L102
L104:
	;
	goto L99
L105:
	;
	v527 = F_RegisterSnapshot(m, v525)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[1]))
	if v530 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATRewriteTable[2])))
	if v532&int32(1) == int32(0) {
		goto L38
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v537 = int32(0)
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v34)+188))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	v543 = m.T0[v542].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v34, v527, v537, v537, v537, int32(449))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L111
	}
L110:
	;
	goto L109
L111:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v56)+152))
	if v545 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v548 = F_MakePerTupleExprContext(m, v56)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L1
	} else {
		goto L115
	}
L113:
	;
	v550 = v545
	goto L114
L114:
	;
	v551 = int32(_a_F_ATRewriteTable_5)
	v552 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0]))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v550)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v554
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+40)) = v557
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v560)+188))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+20))
	v563 = m.T0[v562].(func(*base.Module, int32, int32, int32) int32)(m, v543, int32(1), v448)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L1
	} else {
		goto L116
	}
L115:
	;
	v550 = v548
	goto L114
L116:
	;
	if v563 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	goto L120
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[0])) = v552
	v1236 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(v1236)+188))
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(v1237)+12))
	m.T0[v1238].(func(*base.Module, int32))(m, v543)
	mBase = m.M
	v1240 = m.ExcPending
	if v1240 != 0 {
		goto L1
	} else {
		goto L228
	}
L120:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v592 <= int32(0) {
		v859 = v448
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L119
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v859
	if v378 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L123:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v595)))
	v597 = int32(*(*int16)(unsafe.Add(mBase, uint32(v448)+6)))
	if v597 < v596 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v448)+8))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v599)+16))
	m.T0[v600].(func(*base.Module, int32, int32))(m, v448, v596)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+12))
	m.T0[v604].(func(*base.Module, int32))(m, v447)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L1
	} else {
		goto L128
	}
L127:
	;
	goto L126
L128:
	;
	v607 = int32(*(*int16)(unsafe.Add(mBase, uint32(v448)+6)))
	v609 = v607 << (uint(int32(3)) % 32)
	if v609 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v448)+16))
	base.MemoryCopy(m, v610, v611, v609)
	goto L131
L130:
	;
	goto L131
L131:
	;
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v448)+6)))
	if v613 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v448)+20))
	base.MemoryCopy(m, v614, v615, v613)
	goto L134
L133:
	;
	goto L134
L134:
	;
	if v509 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v447)+40)) = v690
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v448
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v693 == int32(0) {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v619 = int32(0)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v620 <= v619 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v624 = v619
	goto L138
L138:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v651+v624<<(uint(int32(2))%32))))
	v657 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v650+v655))) = uint8(v657)
	v660 = v624 + v657
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v660 < v661 {
		v624 = v660
		goto L138
	} else {
		goto L140
	}
L139:
	;
	goto L135
L140:
	;
	goto L139
L141:
	;
	v785 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v447)+4)))
	v787 = v785 & int32(_a_F_ATRewriteTable_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v447)+4)) = uint16(v787)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v447)+12))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v789)))
	*(*uint16)(unsafe.Add(mBase, uint32(v447)+6)) = uint16(v790)
	goto L151
L142:
	;
	v696 = int32(0)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	if v697 <= v696 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v701 = v696
	goto L144
L144:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v693)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v701<<(uint(int32(2))%32))))
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v731)+12)))
	if v732 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	goto L141
L146:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v737 = int32(*(*int16)(unsafe.Add(mBase, uint32(v731))))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v735)+24))
	v742 = m.T0[v741].(func(*base.Module, int32, int32, int32) int64)(m, v735, v432, v736+v737-int32(1))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v755 = v701 + int32(1)
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v693)+4))
	if v755 < v756 {
		v701 = v755
		goto L144
	} else {
		goto L150
	}
L149:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	v745 = int32(*(*int16)(unsafe.Add(mBase, uint32(v731))))
	*(*int64)(unsafe.Add(mBase, uint32(v744+v745<<(uint(int32(3))%32)-int32(8)))) = v742
	goto L148
L150:
	;
	goto L145
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v432)+4)) = v447
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v793 == int32(0) {
		v859 = v447
		goto L122
	} else {
		goto L152
	}
L152:
	;
	v796 = int32(0)
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v797 <= v796 {
		v859 = v447
		goto L122
	} else {
		goto L153
	}
L153:
	;
	v803 = v796
	goto L154
L154:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v793)+12))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v827+v803<<(uint(int32(2))%32))))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831)+12)))
	if v832 == int32(1) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v859 = v447
	goto L122
L156:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v831)+8))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v447)+20))
	v837 = int32(*(*int16)(unsafe.Add(mBase, uint32(v831))))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v835)+24))
	v842 = m.T0[v841].(func(*base.Module, int32, int32, int32) int64)(m, v835, v432, v836+v837-int32(1))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v855 = v803 + int32(1)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v793)+4))
	if v855 < v856 {
		v803 = v855
		goto L154
	} else {
		goto L160
	}
L159:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v447)+16))
	v845 = int32(*(*int16)(unsafe.Add(mBase, uint32(v831))))
	*(*int64)(unsafe.Add(mBase, uint32(v844+v845<<(uint(int32(3))%32)-int32(8)))) = v842
	goto L158
L160:
	;
	goto L155
L161:
	;
	if v385 != 0 {
		goto L179
	} else {
		goto L180
	}
L162:
	;
	v888 = int32(0)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v889 <= v888 {
		goto L161
	} else {
		goto L163
	}
L163:
	;
	v895 = v888
	goto L164
L164:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v378)+12))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v919+v895<<(uint(int32(2))%32))))
	v924 = int32(*(*int16)(unsafe.Add(mBase, uint32(v859)+6)))
	if v924 < v923 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L174
	}
L166:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v859)+8))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v926)+16))
	m.T0[v927].(func(*base.Module, int32, int32))(m, v859, v923)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v859)+20))
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v930+v923-int32(1)))))
	if v934 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	goto L168
L170:
	;
	v938 = v895 + int32(1)
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v378)+4))
	if v939 <= v938 {
		goto L161
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
	v895 = v938
	goto L164
L174:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+84)) = v949 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v36 + v941<<(uint(int32(3))%32) + v923*int32(100) - int32(68)
	F_errmsg(m, int32(_a_F_ATRewriteTable_11), v30+int32(80))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	F_errtablecol(m, v34, v923)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_12), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
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
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1039 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L180:
	;
	v1001 = F_ExecRelGenVirtualNotNull(m, v386, v859, v56, v380)
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	if v1001 == int32(0) {
		goto L179
	} else {
		goto L182
	}
L182:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+68)) = v1013 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+64)) = v36 + v1005<<(uint(int32(3))%32) + v1001*int32(100) - int32(68)
	F_errmsg(m, int32(_a_F_ATRewriteTable_11), v30-int32(-64))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	F_errtablecol(m, v34, v1001)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_13), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L188:
	;
	if v162 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L189:
	;
	v1042 = int32(0)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	if v1043 <= v1042 {
		goto L188
	} else {
		goto L190
	}
L190:
	;
	v1049 = v1042
	goto L191
L191:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+12))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1073+v1049<<(uint(int32(2))%32))))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	switch v1078 - int32(1) {
	case 0, 8:
		goto L193
	default:
		goto L194
	case 4:
		goto L195
	}
L192:
	;
	goto L188
L193:
	;
	v1127 = v1049 + int32(1)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1039)+4))
	if v1127 < v1128 {
		v1049 = v1127
		goto L191
	} else {
		goto L206
	}
L194:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L203
	}
L195:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+28))
	v1082 = F_ExecCheck(m, v1081, v432)
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	if v1082 != 0 {
		goto L193
	} else {
		goto L197
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v1092
	*(*int32)(unsafe.Add(mBase, uint32(v30)+52)) = v1091 + int32(4)
	F_errmsg(m, int32(_a_F_ATRewriteTable_14), v30+int32(48))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	F_errtableconstraint(m, v34, v1102)
	mBase = m.M
	v1104 = m.ExcPending
	if v1104 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_15), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L203:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v1114
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_0), v30+int32(32))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_16), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L206:
	;
	goto L192
L207:
	;
	if v52 != 0 {
		goto L217
	} else {
		goto L218
	}
L208:
	;
	v1159 = F_ExecCheck(m, v162, v432)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	if v1159 != 0 {
		goto L207
	} else {
		goto L210
	}
L210:
	;
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v1171 = v1169 + int32(4)
	if v1161 == int32(1) {
		goto L39
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v1171
	F_errmsg(m, int32(_a_F_ATRewriteTable_17), v30+int32(16))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errtable(m, v34)
	mBase = m.M
	v1181 = m.ExcPending
	if v1181 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_18), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1187)+80))
	m.T0[v1188].(func(*base.Module, int32, int32, int32, int32, int32))(m, v51, v859, v54, v53, v55)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v432)+20))
	F_MemoryContextReset(m, v1191)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, _c_F_ATRewriteTable[3]))
	if v1195 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v448)+40)) = v1199
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1202)+188))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1203)+20))
	v1205 = m.T0[v1204].(func(*base.Module, int32, int32, int32) int32)(m, v543, int32(1), v448)
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	if v1205 != 0 {
		goto L120
	} else {
		goto L227
	}
L227:
	;
	goto L121
L228:
	;
	F_UnregisterSnapshot(m, v527)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_ExecDropSingleTupleTableSlot(m, v448)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	if v447 == int32(0) {
		goto L40
	} else {
		goto L231
	}
L231:
	;
	F_ExecDropSingleTupleTableSlot(m, v447)
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	goto L40
L233:
	;
	F_relation_close(m, v34, int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	if v52 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	F_FreeBulkInsertState(m, v55)
	mBase = m.M
	v1282 = m.ExcPending
	if v1282 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	goto L237
L237:
	;
	m.G0 = v30 + int32(144)
	return
L238:
	;
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(v51)+188))
	if v1283 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	F_relation_close(m, v51, int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L1
	} else {
		goto L243
	}
L240:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v1283)+108))
	if v1286 == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	m.T0[v1286].(func(*base.Module, int32, int32))(m, v51, v53)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	goto L239
L243:
	;
	goto L237
L244:
	;
	F_errtable(m, v34)
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_1), int32(_a_F_ATRewriteTable_19), int32(_a_F_ATRewriteTable_3))
	mBase = m.M
	v1309 = m.ExcPending
	if v1309 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errmsg_internal(m, int32(_a_F_ATRewriteTable_20), int32(0))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	F_errfinish(m, int32(_a_F_ATRewriteTable_21), int32(931), int32(_a_F_ATRewriteTable_22))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L249
	}
L249:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
