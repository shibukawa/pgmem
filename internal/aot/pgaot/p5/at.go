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
	var v20 int64
	_ = v20
	var v23 int32
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
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	v20 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v20
	v23 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
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
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v179
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
	v178 = v50
	v179 = v51
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
	v178 = int32(0)
	v179 = int32(2606)
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
	F_CheckTableNotInUse(m, v224, int32(565633))
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
	v299 = int32(28672)
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
	F_sequence_close(m, v224, int32(0))
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
	F_errmsg(m, int32(252122), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(517363), int32(10023), int32(97262))
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
	F_errmsg(m, int32(152422), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(517363), int32(4460), int32(428883))
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
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
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
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
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
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
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
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
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int64
	_ = v402
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
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
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v517 int32
	_ = v517
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
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
	var v578 int32
	_ = v578
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
	var v643 int32
	_ = v643
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
		goto L214
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L210
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L1
	} else {
		goto L206
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L1
	} else {
		goto L203
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L199
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L195
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L188
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L184
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L180
	}
L12:
	;
	v52 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+131)))
	if v44 == int32(1) {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v49 = l3 + int32(48)
	goto L12
L17:
	;
	v49 = l3 + int32(48)
	goto L12
L18:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+16)))
	if v54 <= int32(0) {
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
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v132 = F_check_for_column_name_collision(m, l3, v129, v32&int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	v58 = F_SearchSysCacheCopyAttName(m, v30, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v58 == int32(0) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+22)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	F_typenameTypeIdAndMod(m, int32(0), v65, v28+int32(116), v28+int32(104))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v73 = v62 + v63
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+68))
	if v72 != v74 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v28)+104))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+76))
	if v76 != v77 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v80 = F_GetColumnDefCollation(m, int32(0), v33, v72)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	if v80 != v82 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+94)))
	v86 = v84 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v73)+94)) = uint16(v86)
	if base.I32_extend16_s(v86) != v86 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	F_CatalogTupleUpdate(m, v52, v58+int32(4), v58)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v58)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v98 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v98 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v100 + int32(4)
	F_errmsg(m, int32(755543), v28+int32(32))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_sequence_close(m, v52, int32(3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_errfinish(m, int32(517363), int32(7304), int32(287817))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v123
	v126 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v126
	goto L19
L40:
	;
	if v132 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_sequence_close(m, v52, int32(3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if l6 != 0 {
		v152 = v33
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _consts[304]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
	v143 = *(*int64)(unsafe.Add(mBase, _consts[305]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v143
	goto L19
L45:
	;
	if l5 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if l9 == int32(0) {
		v152 = v33
		goto L45
	} else {
		goto L47
	}
L47:
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
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v148
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	v152 = v151
	goto L45
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
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
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
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
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+110)))
	F_CheckAttributeType(m, v194+v204, v198, v199, v209, base.B2i32(v211 == int32(118))<<(uint(int32(3))%32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v218 = int32(0)
	F_InsertPgAttributeTuples(m, v52, v189, v30, v218, v218)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_sequence_close(m, v52, int32(3))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+120)) = uint16(v196)
	F_CatalogTupleUpdate(m, v167, v171+int32(4), v171)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_pfree(m, v171)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v233 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_RunObjectPostCreateHook(m, int32(1259), v30, v196, int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_sequence_close(m, v167, int32(3))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L69
L71:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	if v243 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v245 = F_palloc(m, int32(12))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	switch v181 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L82
	default:
		goto L81
	}
L76:
	;
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+94)))
	*(*uint16)(unsafe.Add(mBase, uint32(v245))) = uint16(v247)
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v152)+28))
	v250 = F_copyObjectImpl(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v245)+4)) = v250
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v245)+8)) = uint8(v253)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = v245
	v260 = F_list_make1_impl(m, int32(1), v28+int32(20))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v262 = int32(0)
	v267 = F_AddRelationNewConstraints(m, l3, v260, v262, v262, int32(1), v262, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L75
L81:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(1247)
	F_recordDependencyOn(m, v28+int32(116), v28+int32(104), int32(110))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L141
	}
L82:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)))
	if v274 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+19)))
	v507 = v505 | v506
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+76)) = uint8(v507)
	goto L81
L84:
	;
	v344 = F_expression_planner(m, v339)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L105
	}
L85:
	;
	v276 = F_palloc0(m, int32(12))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v292 = int32(*(*int16)(unsafe.Add(mBase, uint32(v194)+94)))
	v293 = F_build_column_default(m, l3, v292)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L91
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = int32(59)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v152)+40))
	v281 = int32(0)
	v285 = F_RangeVarGetRelidExtended(m, v280, v281, v281, v281, v281)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276)+4)) = v285
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v276)+8)) = v288
	v290 = F_DomainHasConstraints(m, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v339 = v276
	v342 = v290
	goto L84
L91:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v194)+88))
	v296 = F_DomainHasConstraints(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v293 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v293 == int32(0) {
		goto L83
	} else {
		goto L104
	}
L94:
	;
	if v296 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v301 = v194 + int32(96)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = v302
	v305 = v194 + int32(88)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v309 = F_getBaseTypeAndTypmod(m, v306, v28+int32(116))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v311 = F_get_typcollation(m, v309)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v28)+116))
	v315 = F_makeNullConst(m, v309, v314, v311)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v301)))
	v322 = F_coerce_to_target_type(m, int32(0), v315, v309, v317, v318, int32(1), int32(2), int32(-1))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v322 != 0 {
		v339 = v322
		v342 = v296
		goto L84
	} else {
		goto L100
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errmsg_internal(m, int32(290549), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(517363), int32(7507), int32(287817))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v339 = v293
	v342 = v296
	goto L84
L105:
	;
	v347 = F_palloc0(m, int32(16))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+94)))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v344
	*(*uint16)(unsafe.Add(mBase, uint32(v347))) = uint16(v349)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	*(*uint8)(unsafe.Add(mBase, uint32(v347)+12)) = uint8(base.B2i32(v352 != int32(0)))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v357 = F_lappend(m, v356, v347)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+68)) = v357
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360)+119)))
	if v361 != int32(114) {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	F_FreeExecutorState(m, v372)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L140
	}
L109:
	;
	if v488&int32(255) == int32(118) {
		goto L83
	} else {
		goto L139
	}
L110:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	v488 = v487
	goto L109
L111:
	;
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+44)))
	if (base.B2i32(v364 != int32(0))|v342)&int32(1) != 0 {
		v488 = v364
		goto L109
	} else {
		goto L112
	}
L112:
	;
	v370 = F_contain_volatile_functions(m, v344)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	if v370 != 0 {
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v372 = F_CreateExecutorState(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v374 = F_ExecPrepareExpr(m, v344, v372)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v372)+152))
	if v376 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v379 = v376
	goto L119
L118:
	;
	v377 = F_MakePerTupleExprContext(m, v372)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v374)+20))
	v383 = m.T0[v382].(func(*base.Module, int32, int32, int32) int32)(m, v374, v379, v28+int32(116))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L121
	}
L120:
	;
	v379 = v377
	goto L119
L121:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+116)))
	if v385 != 0 {
		goto L108
	} else {
		goto L122
	}
L122:
	;
	v386 = int32(*(*int16)(unsafe.Add(mBase, uint32(v194)+94)))
	v387 = m.G0
	v389 = v387 - int32(192)
	m.G0 = v389
	*(*int32)(unsafe.Add(mBase, uint32(v389)+188)) = v383
	v397 = F__emscripten_memset_bulkmem(m, v389+int32(80), base.I32_extend8_s(int32(0)), int32(96))
	mBase = m.M
	goto L123
L123:
	;
	v398 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+72)) = uint8(v398)
	v402 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v389-int32(-64)))) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v389)+56)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v389)+48)) = v402
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+40)) = uint8(v398)
	*(*int64)(unsafe.Add(mBase, uint32(v389)+32)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v389)+24)) = v402
	*(*int64)(unsafe.Add(mBase, uint32(v389)+16)) = v402
	v418 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v422 = F_SearchSysCache2(m, int32(7), v421, v386)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	if v422 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v422)+16))
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444)+22)))
	v446 = v444 + v445
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)+68))
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v446)+72)))
	v449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446)+82)))
	v450 = int32(*(*int8)(unsafe.Add(mBase, uint32(v446)+83)))
	v451 = F_construct_array(m, v389+int32(188), int32(1), v447, v448, v449, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L132
	}
L129:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v430
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v386
	F_errmsg_internal(m, int32(50140), v389)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(519147), int32(2051), int32(329032))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v453 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+132)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v389)+188)) = v451
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+29)) = uint8(v453)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+176)) = v451
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+40)) = uint8(v453)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v418)+52))
	v468 = F_heap_modify_tuple(m, v422, v461, v389+int32(80), v389+int32(48), v389+int32(16))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_CatalogTupleUpdate(m, v418, v468+int32(4), v468)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_ReleaseCatCache(m, v422)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	F_sequence_close(m, v418, int32(3))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	m.G0 = v389 + int32(192)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_FreeExecutorState(m, v372)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L81
L139:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l2)+80))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v493 | int32(2)
	goto L83
L140:
	;
	goto L83
L141:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v194)+116))
	if v534 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	v557 = F_find_inheritance_children(m, v556, l7)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	if v534 == int32(100) {
		goto L142
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+124)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v28)+120)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v28)+116)) = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+112)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v534
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = int32(3456)
	F_recordDependencyOn(m, v28+int32(116), v28+int32(104), int32(110))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	goto L142
L146:
	;
	if v557 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v559 = l5
	goto L149
L148:
	;
	v559 = int32(1)
	goto L149
L149:
	;
	if v559 == int32(0) {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if l6 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v565 = F_copyObjectImpl(m, v562)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L154
	}
L152:
	;
	v573 = v562
	goto L153
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v573
	if v557 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v565)+20))
	v568 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v567)+18)) = uint8(v568)
	v570 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v567)+16)) = uint16(v570)
	v573 = v565
	goto L153
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L19
L156:
	;
	v577 = int32(0)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v578 <= v577 {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v583 = v577
	goto L158
L158:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v606+v583<<(uint(int32(2))%32))))
	v612 = F_table_open(m, v610, int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	goto L155
L160:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v612)+48))
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614)+118)))
	if v615 == int32(116) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v618 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+24)))
	if v618 == int32(0) {
		goto L3
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	F_CheckTableNotInUse(m, v612, int32(565633))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L165
	}
L164:
	;
	goto L163
L165:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v612)+56))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v625 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	F_ATExecAddColumn(m, v28+int32(116), l1, v718, v612, v28+int32(104), l5, int32(1), l7, l8, l9)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L177
	}
L167:
	;
	v693 = F_palloc0(m, int32(140))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L174
	}
L168:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v625)+4))
	if v628 <= int32(0) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625)+12))
	v643 = int32(0)
	goto L170
L170:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v631+v643<<(uint(int32(2))%32))))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)))
	if v662 == v624 {
		v718 = v661
		goto L166
	} else {
		goto L172
	}
L171:
	;
	goto L167
L172:
	;
	v665 = v643 + int32(1)
	if v628 != v665 {
		v643 = v665
		goto L170
	} else {
		goto L173
	}
L173:
	;
	goto L171
L174:
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
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v693)+8)) = v702
	*(*int64)(unsafe.Add(mBase, uint32(v693)+88)) = int64(0)
	v707 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v693)+84)) = uint8(v707)
	v709 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v693)+96)) = uint16(v709)
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v712 = F_lappend(m, v711, v693)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v712
	v718 = v693
	goto L166
L177:
	;
	F_sequence_close(m, v612, int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v751 = v583 + int32(1)
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v751 < v752 {
		v583 = v751
		goto L158
	} else {
		goto L179
	}
L179:
	;
	goto L159
L180:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	F_errmsg(m, int32(261097), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(517363), int32(7250), int32(287817))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v835
	*(*int32)(unsafe.Add(mBase, uint32(v28)+80)) = v834 + int32(4)
	F_errmsg(m, int32(742924), v28+int32(80))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errfinish(m, int32(517363), int32(7280), int32(287817))
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
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
	F_errcode(m, int32(17432708))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v858
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v857 + int32(4)
	F_errmsg(m, int32(742867), v28-int32(-64))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v868 = F_get_collation_name(m, v80)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v73)+96))
	v871 = F_get_collation_name(m, v870)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+52)) = v871
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v868
	F_errdetail(m, int32(731904), v28+int32(48))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(517363), int32(7289), int32(287817))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errmsg(m, int32(128411), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(517363), int32(7296), int32(287817))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	F_errmsg(m, int32(176122), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	F_errfinish(m, int32(517363), int32(7356), int32(287817))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
	F_errmsg_internal(m, int32(49952), v28)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	F_errfinish(m, int32(517363), int32(7362), int32(287817))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
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
	F_errcode(m, int32(17039621))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = int32(1600)
	F_errmsg(m, int32(157483), v28+int32(16))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(517363), int32(7372), int32(287817))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errmsg(m, int32(252167), int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(517363), int32(7603), int32(287817))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	F_errmsg(m, int32(152422), int32(0))
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(517363), int32(4460), int32(428883))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
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
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
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
	v324 = m.ExcPending
	if v324 != 0 {
		goto L16
	} else {
		goto L79
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L16
	} else {
		goto L75
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L16
	} else {
		goto L70
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L16
	} else {
		goto L67
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L16
	} else {
		goto L63
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L59
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L16
	} else {
		goto L55
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L16
	} else {
		goto L51
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
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L46
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
	v72 = *(*int32)(unsafe.Add(mBase, _consts[231]))
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
	F_sequence_close(m, v34, int32(3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	if v21 != int32(112) {
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
	if l5 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v95 = F_find_inheritance_children(m, v94, l4)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	if v95 == int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v99 <= int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v104 = int32(0)
	goto L40
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v104<<(uint(int32(2))%32))))
	v124 = F_table_open(m, v122, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L42
	}
L41:
	;
	goto L34
L42:
	;
	v126 = int32(1)
	F_ATExecAddIdentity(m, v16+int32(132), v124, l2, l3, l4, v126, v126)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L43
	}
L43:
	;
	F_sequence_close(m, v124, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	v134 = v104 + int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v134 < v135 {
		v104 = v134
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(412833), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L16
	} else {
		goto L48
	}
L48:
	;
	F_errhint(m, int32(672007), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L16
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(517363), int32(8256), int32(10890))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(261266), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(517363), int32(8261), int32(10890))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v196 + int32(4)
	F_errmsg(m, int32(77509), v16)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(517363), int32(8270), int32(10890))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L16
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(743702), v16+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(517363), int32(8279), int32(10890))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v234 + int32(4)
	F_errmsg(m, int32(483373), v16+int32(112))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(517363), int32(8290), int32(10890))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L16
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v253 + int32(4)
	F_errmsg_internal(m, int32(741006), v16+int32(32))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(517363), int32(8305), int32(10890))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L16
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L16
	} else {
		goto L71
	}
L71:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v276 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v57 + v276
	*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v275 + v276
	F_errmsg(m, int32(738562), v16+int32(96))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L16
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = int32(542852)
	F_errhint(m, int32(631914), v16+int32(80))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L16
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(517363), int32(8314), int32(10890))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v306 + int32(4)
	F_errmsg(m, int32(286005), v16-int32(-64))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(517363), int32(8321), int32(10890))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L16
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v328 + int32(4)
	F_errmsg(m, int32(362337), v16+int32(48))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(517363), int32(8327), int32(10890))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v212 int32
	_ = v212
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
	return v212
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+74)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v33 == v34 {
		v212 = int32(0)
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
	v137 = m.ExcPending
	if v137 != 0 {
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
	if base.Ui32(int32(11)) < base.Ui32(v84) {
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
	v117 = F_systable_getnext(m, v51)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	if int32(1)<<(uint(v84)%32)&int32(3075) == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v93 = F_heap_copytuple(m, v63)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+22)))
	v97 = v95 + v96
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+97)) = uint8(v38)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+96)) = uint8(v39)
	F_CatalogTupleUpdate(m, l2, v93+int32(4), v93)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v105 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v108 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v107, v108, v108, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v93)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	goto L21
L31:
	;
	if v117 != 0 {
		v63 = v117
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
		v212 = v46
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l3)+48))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+119)))
	if v141 != int32(112) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v144 = F_get_rel_relkind(m, v29)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+22)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151+v152)))
	F_ScanKeyInit(m, v20, int32(12), int32(3), int32(184), v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	if v144 != int32(112) {
		v212 = v46
		goto L3
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v158 = int32(1)
	v161 = F_systable_beginscan(m, l1, int32(2579), v158, int32(0), v158, v20)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L42
L42:
	;
	v180 = F_systable_getnext(m, v161)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	F_systable_endscan(m, v161)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L51
	}
L44:
	;
	if v180 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+22)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v182+v183)+80))
	v186 = F_table_open(m, v185, l7)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
	v189 = F_ATExecAlterConstrDeferrability(m, l0, l1, l2, v186, v180, int32(1), l6, l7)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_sequence_close(m, v186, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v212 = v46
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
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
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v236 int32
	_ = v236
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
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
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
	var v318 int64
	_ = v318
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v487 int32
	_ = v487
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
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
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
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
	v611 = *(*int32)(unsafe.Add(mBase, _consts[231]))
	if v611 != 0 {
		goto L166
	} else {
		goto L167
	}
L5:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v29)+72))
	if v426 != 0 {
		goto L121
	} else {
		goto L122
	}
L6:
	;
	if v376 == int32(73) {
		goto L5
	} else {
		goto L119
	}
L7:
	;
	if int32(1)<<(uint(v378)%32)&int32(131137) != 0 {
		goto L5
	} else {
		goto L118
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
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
		v165 = l1
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
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L110
	}
L13:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	if v166 == v165 {
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
		v165 = l1
		goto L13
	} else {
		goto L48
	}
L16:
	;
	if l2 != 0 {
		v165 = l1
		goto L13
	} else {
		goto L34
	}
L17:
	;
	if l2 != 0 {
		v165 = l1
		goto L13
	} else {
		goto L28
	}
L18:
	;
	if l2 != 0 {
		v165 = l1
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
	v165 = v61
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
	F_errmsg(m, int32(728447), v16+int32(32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errhint(m, int32(679755), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(517363), int32(16118), int32(229439))
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
	F_errmsg(m, int32(728447), v16+int32(48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errhint(m, int32(679755), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(517363), int32(16130), int32(229439))
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
		v165 = l1
		goto L13
	} else {
		goto L35
	}
L35:
	;
	v93 = F_sequenceIsOwned(m, l0, int32(97), v16+int32(224), v16+int32(432))
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
	v102 = F_sequenceIsOwned(m, l0, int32(105), v16+int32(224), v16+int32(432))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
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
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	if v102 == int32(0) {
		v165 = l1
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
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v114 = v29 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v114
	F_errmsg(m, int32(754376), v16+int32(80))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v16)+224))
	v122 = F_get_rel_name(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v114
	F_errdetail(m, int32(696138), v16-int32(-64))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(517363), int32(16148), int32(229439))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v29 + int32(4)
	F_errmsg(m, int32(387403), v16+int32(112))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = int32(564513)
	F_errhint(m, int32(679364), v16+int32(96))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(517363), int32(16160), int32(229439))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v165 = l1
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
	v215 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+208)) = uint16(v215)
	v217 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+200)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+192)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+152)) = v217
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+160)) = uint16(v215)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+128)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+184)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+176)) = v217
	*(*int64)(unsafe.Add(mBase, uint32(v16)+136)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v16)+244)) = v165
	v236 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+133)) = uint8(v236)
	v242 = F_SysCacheGetAttr(m, int32(57), v25, int32(32), v16+int32(127))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L78
	}
L57:
	;
	v168 = F_superuser(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v168 != 0 {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	v173 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	v174 = F_object_ownercheck(m, int32(1259), l0, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v174 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v179 = F_get_rel_relkind(m, l0)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _consts[237]))
	F_check_can_set_role(m, v199, v165)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L73
	}
L64:
	;
	switch v179 - int32(73) {
	case 0, 32:
		goto L71
	default:
		v190 = int32(41)
		goto L66
	case 10:
		goto L70
	case 29:
		goto L67
	case 36:
		goto L68
	case 45:
		goto L69
	}
L65:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
	F_aclcheck_error(m, int32(2), v192, v193+int32(4))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L72
	}
L66:
	;
	v192 = v190
	goto L65
L67:
	;
	v190 = int32(18)
	goto L66
L68:
	;
	v192 = int32(23)
	goto L65
L69:
	;
	v192 = int32(51)
	goto L65
L70:
	;
	v192 = int32(37)
	goto L65
L71:
	;
	v192 = int32(20)
	goto L65
L72:
	;
	goto L63
L73:
	;
	v204 = F_object_aclcheck(m, int32(2615), v170, v165, int64(512))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	if v204 == int32(0) {
		goto L56
	} else {
		goto L75
	}
L75:
	;
	v209 = F_get_namespace_name(m, v170)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_aclcheck_error(m, v204, int32(36), v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L56
L78:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+127)))
	if v244 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v247 = F_pg_detoast_datum(m, v242)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v262 = F_heap_modify_tuple(m, v25, v255, v16+int32(224), v16+int32(176), v16+int32(128))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v250 = F_aclnewowner(m, v247, v249, v165)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+348)) = v250
	v253 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+159)) = uint8(v253)
	goto L81
L84:
	;
	F_CatalogTupleUpdate(m, v22, v262+int32(4), v262)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_pfree(m, v262)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v29)+80))
	v273 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_ScanKeyInit(m, v16+int32(544), int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v283 = int32(1)
	v288 = F_systable_beginscan(m, v273, int32(2659), v283, int32(0), v283, v16+int32(544))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v290 = F_systable_getnext(m, v288)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	if v290 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v294 = v290
	goto L94
L92:
	;
	goto L93
L93:
	;
	F_systable_endscan(m, v288)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L107
	}
L94:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+22)))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v306)+91)))
	if v308 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L93
L96:
	;
	v356 = F_systable_getnext(m, v288)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L105
	}
L97:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v273)+52))
	v313 = F_heap_getattr_6(m, v294, int32(22), v310, v16+int32(367))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+367)))
	if v315 != 0 {
		goto L96
	} else {
		goto L99
	}
L99:
	;
	v316 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+424)) = uint8(v316)
	v318 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+416)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v16)+384)) = v318
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+392)) = uint8(v316)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+408)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v16)+400)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v16)+368)) = v318
	*(*int64)(unsafe.Add(mBase, uint32(v16)+376)) = v318
	v332 = F_pg_detoast_datum(m, v313)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v334 = F_aclnewowner(m, v332, v270, v165)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+516)) = v334
	v337 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+389)) = uint8(v337)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v273)+52))
	v346 = F_heap_modify_tuple(m, v294, v339, v16+int32(432), v16+int32(400), v16+int32(368))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_CatalogTupleUpdate(m, v273, v346+int32(4), v346)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_pfree(m, v346)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L96
L105:
	;
	if v356 != 0 {
		v294 = v356
		goto L94
	} else {
		goto L106
	}
L106:
	;
	goto L95
L107:
	;
	F_sequence_close(m, v273, int32(3))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	v378 = v376 - int32(99)
	if base.Ui32(v378) <= base.Ui32(int32(17)) {
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
	F_errmsg_internal(m, int32(49952), v16)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(517363), int32(16090), int32(229439))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
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
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v29 + int32(4)
	F_errmsg(m, int32(740265), v16+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v409 = int32(*(*int8)(unsafe.Add(mBase, uint32(v29)+119)))
	F_errdetail_relkind_not_supported(m, v409)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(517363), int32(16171), int32(229439))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
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
	F_changeDependencyOnOwner(m, int32(1259), l0, v165)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L5
L121:
	;
	F_AlterTypeOwnerInternal(m, v426, v165)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+119)))
	v431 = v429 - int32(109)
	if base.Ui32(int32(7)) < base.Ui32(v431) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	goto L123
L125:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	if v501 != 0 {
		goto L137
	} else {
		goto L138
	}
L126:
	;
	if int32(1)<<(uint(v431)%32)&int32(169) == int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v440 = F_RelationGetIndexList(m, v18)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L129
	}
L128:
	;
	F_list_free(m, v440)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L136
	}
L129:
	;
	if v440 == int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v444 = int32(0)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v445 <= v444 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v450 = v444
	goto L132
L132:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v461+v450<<(uint(int32(2))%32))))
	F_ATExecChangeOwner(m, v465, v165, int32(1), l3)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L128
L134:
	;
	v470 = v450 + int32(1)
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v440)+4))
	if v470 < v471 {
		v450 = v470
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	goto L125
L137:
	;
	F_ATExecChangeOwner(m, v501, v165, int32(1), l3)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v507 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L141
	}
L140:
	;
	goto L139
L141:
	;
	F_ScanKeyInit(m, v16+int32(432), int32(4), int32(3), int32(184), int32(1259))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_ScanKeyInit(m, v16+int32(480), int32(5), int32(3), int32(184), l0)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v530 = F_systable_beginscan(m, v507, int32(2674), int32(1), int32(0), int32(2), v16+int32(432))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v532 = F_systable_getnext(m, v530)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	if v532 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v536 = v532
	goto L149
L147:
	;
	goto L148
L148:
	;
	F_systable_endscan(m, v530)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L164
	}
L149:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v536)+16))
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+22)))
	v549 = v547 + v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+20))
	if v550 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	goto L148
L151:
	;
	v577 = F_systable_getnext(m, v530)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L162
	}
L152:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	if v553 != int32(1259) {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v549)+8))
	if v556 != 0 {
		goto L151
	} else {
		goto L154
	}
L154:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+24)))
	switch v557 - int32(97) {
	case 0, 8:
		goto L155
	default:
		goto L151
	}
L155:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	v561 = F_relation_open(m, v560, l3)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v561)+48))
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563)+119)))
	if v564 == int32(83) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v549)+4))
	F_ATExecChangeOwner(m, v567, v165, int32(1), l3)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	v572 = l3
	goto L159
L159:
	;
	F_relation_close(m, v561, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L161
	}
L160:
	;
	v572 = int32(0)
	goto L159
L161:
	;
	goto L151
L162:
	;
	if v577 != 0 {
		v536 = v577
		goto L149
	} else {
		goto L163
	}
L163:
	;
	goto L150
L164:
	;
	F_relation_close(m, v507, int32(1))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	goto L4
L166:
	;
	v613 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v613, v613, v613)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
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
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	F_sequence_close(m, v22, int32(3))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_relation_close(m, v18, int32(0))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
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
					v29 = *(*int32)(unsafe.Add(mBase, _consts[231]))
					if v29 != 0 {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v32 = int32(0)
						F_RunObjectPostAlterHook(m, int32(1259), v31, v32, v32, v32)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_sequence_close(m, v14, int32(3))
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
						F_sequence_close(m, v14, int32(3))
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
					F_errmsg_internal(m, int32(49952), v9)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errfinish(m, int32(517363), int32(18647), int32(11104))
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
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
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
	var v135 int32
	_ = v135
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
	var v164 int32
	_ = v164
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
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
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
	var v295 int32
	_ = v295
	var v324 int32
	_ = v324
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
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
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
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
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
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v499 int32
	_ = v499
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
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
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v786 int32
	_ = v786
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v802 int32
	_ = v802
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
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v968 int32
	_ = v968
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1113 int32
	_ = v1113
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1252 int32
	_ = v1252
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
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
	v45 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
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
	v50 = v3
	v51 = int32(1)
	v52 = v3
	v53 = v3
	v54 = v3
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
	v47 = F_GetBulkInsertState(m)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v50 = v39
	v51 = int32(0)
	v52 = int32(2)
	v53 = v47
	v54 = v45
	goto L3
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v57 == int32(0) {
		v135 = v3
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v153 != 0 {
		goto L29
	} else {
		goto L30
	}
L14:
	;
	v60 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v61 <= v60 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v135 = v3
	goto L13
L16:
	;
	goto L17
L17:
	;
	v65 = v60
	v72 = v3
	goto L18
L18:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v90+v65<<(uint(int32(2))%32))))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	switch v95 - int32(5) {
	case 0:
		goto L21
	default:
		goto L22
	case 4:
		v122 = v72
		goto L20
	}
L19:
	;
	v135 = v122
	goto L13
L20:
	;
	v124 = v65 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	if v124 < v125 {
		v65 = v124
		v72 = v122
		goto L18
	} else {
		goto L28
	}
L21:
	;
	v114 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
	v117 = F_expand_generated_columns_in_expr(m, v115, v33, v114)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L26
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+128)) = v102
	F_errmsg_internal(m, int32(506121), v29+int32(128))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(517363), int32(6200), int32(415426))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v119 = F_ExecPrepareExpr(m, v117, v55)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+28)) = v119
	v122 = v114
	goto L20
L28:
	;
	goto L19
L29:
	;
	v155 = F_ExecPrepareExpr(m, v153, v55)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	v158 = v135
	v159 = int32(0)
	goto L31
L31:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v160 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v158 = int32(1)
	v159 = v155
	goto L31
L33:
	;
	if v51 != 0 {
		goto L44
	} else {
		goto L45
	}
L34:
	;
	v163 = int32(0)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v164 <= v163 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v168 = v163
	goto L36
L36:
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
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+8)) = v200
	v204 = v168 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v204 < v205 {
		v168 = v204
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_FreeExecutorState(m, v55)
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L247
	}
L41:
	;
	v383 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L67
	}
L42:
	;
	if v292 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L43:
	;
	v324 = int32(1)
	if (v51^int32(-1)|v158)&v324 == int32(0) {
		goto L40
	} else {
		goto L61
	}
L44:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+76)))
	if v233 != int32(1) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v236 <= int32(0) {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v240 = v35 + int32(20)
	v243 = int32(0)
	v245 = v236
	v255 = v3
	v257 = v3
	goto L49
L49:
	;
	v270 = v240 + v243<<(uint(int32(4))%32)
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+11)))
	if v271 != int32(118) {
		v291 = v255
		v292 = v257
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v291|v292 != 0 {
		goto L42
	} else {
		goto L60
	}
L51:
	;
	v294 = v243 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v294 < v295 {
		v243 = v294
		v245 = v295
		v255 = v291
		v257 = v292
		goto L49
	} else {
		goto L59
	}
L52:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+9)))
	if v274 != 0 {
		v291 = v255
		v292 = v257
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v280 = v240 + v245<<(uint(int32(4))%32) + v243*int32(100)
	v281 = int32(*(*int16)(unsafe.Add(mBase, uint32(v280)+74)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280)+90)))
	if v282 != int32(118) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v285 = F_lappend_int(m, v255, v281)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v287 = F_lappend_int(m, v257, v281)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v291 = v285
	v292 = v257
	goto L51
L58:
	;
	v291 = v255
	v292 = v287
	goto L51
L59:
	;
	goto L50
L60:
	;
	goto L43
L61:
	;
	v332 = int32(0)
	v368 = v332
	v370 = v332
	v373 = v3
	v376 = v324
	goto L41
L62:
	;
	v368 = v291
	v370 = int32(0)
	v373 = v3
	v376 = int32(1)
	goto L41
L63:
	;
	goto L64
L64:
	;
	v338 = int32(4554240)
	v339 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v55)+100))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v341
	v344 = F_palloc0(m, int32(216))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
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
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v339
	v368 = v291
	v370 = v292
	v373 = v344
	v376 = v3
	goto L41
L67:
	;
	if v51 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v419 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L69:
	;
	if v383 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	if v383 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L72:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+112)) = v387 + int32(4)
	F_errmsg_internal(m, int32(752259), v29+int32(112))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_TransferPredicateLocksToHeapRelation(m, v33)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	F_errfinish(m, int32(517363), int32(6289), int32(415426))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
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
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+96)) = v405 + int32(4)
	F_errmsg_internal(m, int32(752238), v29+int32(96))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(517363), int32(6293), int32(415426))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L68
L81:
	;
	v422 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v424 = v419
	goto L83
L83:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v426 = F_table_slot_callbacks(m, v33)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L85
	}
L84:
	;
	v424 = v422
	goto L83
L85:
	;
	if v425 != 0 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v441 = int32(0)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v441 < v442 {
		goto L95
	} else {
		goto L96
	}
L87:
	;
	v428 = F_MakeSingleTupleTableSlot(m, v36, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v437 = F_MakeSingleTupleTableSlot(m, v35, v426)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L94
	}
L90:
	;
	v430 = F_table_slot_callbacks(m, v50)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v432 = F_MakeSingleTupleTableSlot(m, v35, v430)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	F_ExecStoreAllNullTuple(m, v432)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v439 = v432
	v440 = v428
	goto L86
L94:
	;
	v439 = int32(0)
	v440 = v437
	goto L86
L95:
	;
	v449 = int32(0)
	v450 = v442
	v456 = v441
	goto L98
L96:
	;
	v499 = v441
	goto L97
L97:
	;
	v517 = F_GetLatestSnapshot(m)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L105
	}
L98:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+int32(111)+v450<<(uint(int32(4))%32)+v449*int32(100)))))
	if v480 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v499 = v487
	goto L97
L100:
	;
	v483 = F_lappend_int(m, v456, v449)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	v486 = v450
	v487 = v456
	goto L102
L102:
	;
	v489 = v449 + int32(1)
	if v489 < v486 {
		v449 = v489
		v450 = v486
		v456 = v487
		goto L98
	} else {
		goto L104
	}
L103:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v486 = v485
	v487 = v483
	goto L102
L104:
	;
	goto L99
L105:
	;
	v519 = F_RegisterSnapshot(m, v517)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v521 = int32(0)
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v33)+188))
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)+8))
	v527 = m.T0[v526].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v33, v519, v521, v521, v521, int32(449))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v55)+152))
	if v529 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v532 = F_MakePerTupleExprContext(m, v55)
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	v534 = v529
	goto L110
L110:
	;
	v535 = int32(4554240)
	v536 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v534)+20))
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v538
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v541
	v544 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	if v544 != 0 {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v534 = v532
	goto L110
L112:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+104)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L1
	} else {
		goto L236
	}
L113:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L233
	}
L114:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, _consts[245])))
	if v546&int32(1) == int32(0) {
		goto L113
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	goto L118
L117:
	;
	goto L116
L118:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+188))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+20))
	v581 = m.T0[v580].(func(*base.Module, int32, int32, int32) int32)(m, v527, int32(1), v440)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L123
	}
L119:
	;
	goto L113
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v859
	if v368 == int32(0) {
		goto L171
	} else {
		goto L172
	}
L121:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v33)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+36)) = v694
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v440
	v697 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v697 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L122:
	;
	v630 = v610
	goto L148
L123:
	;
	if v581 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v583 <= int32(0) {
		v859 = v440
		goto L120
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v536
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v616)+188))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+12))
	m.T0[v618].(func(*base.Module, int32))(m, v527)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L143
	}
L127:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v440)+12))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	v588 = int32(*(*int16)(unsafe.Add(mBase, uint32(v440)+6)))
	if v588 < v587 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_slot_getsomeattrs_int(m, v440, v587)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v439)+8))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+12))
	m.T0[v593].(func(*base.Module, int32))(m, v439)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L132
	}
L131:
	;
	goto L130
L132:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v440)+16))
	v598 = int32(*(*int16)(unsafe.Add(mBase, uint32(v440)+6)))
	v600 = v598 << (uint(int32(2)) % 32)
	if v600 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v440)+20))
	v605 = int32(*(*int16)(unsafe.Add(mBase, uint32(v440)+6)))
	if v605 != 0 {
		goto L138
	} else {
		goto L139
	}
L134:
	;
	v601 = F__emscripten_memcpy_bulkmem(m, v596, v597, v600)
	mBase = m.M
	goto L136
L135:
	;
	goto L136
L136:
	;
	goto L133
L137:
	;
	if v499 == int32(0) {
		goto L121
	} else {
		goto L141
	}
L138:
	;
	v606 = F__emscripten_memcpy_bulkmem(m, v603, v604, v605)
	mBase = m.M
	goto L140
L139:
	;
	goto L140
L140:
	;
	goto L137
L141:
	;
	v610 = int32(0)
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v610 < v611 {
		goto L122
	} else {
		goto L142
	}
L142:
	;
	goto L121
L143:
	;
	F_UnregisterSnapshot(m, v519)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	F_ExecDropSingleTupleTableSlot(m, v440)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	if v439 == int32(0) {
		goto L40
	} else {
		goto L146
	}
L146:
	;
	F_ExecDropSingleTupleTableSlot(m, v439)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L40
L148:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v656+v630<<(uint(int32(2))%32))))
	v662 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v655+v660))) = uint8(v662)
	v665 = v630 + v662
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v665 < v666 {
		v630 = v665
		goto L148
	} else {
		goto L150
	}
L149:
	;
	goto L121
L150:
	;
	goto L149
L151:
	;
	v786 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v439)+4)))
	v788 = v786 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+4)) = uint16(v788)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v439)+12))
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v790)))
	*(*uint16)(unsafe.Add(mBase, uint32(v439)+6)) = uint16(v791)
	goto L161
L152:
	;
	v700 = int32(0)
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	if v701 <= v700 {
		goto L151
	} else {
		goto L153
	}
L153:
	;
	v705 = v700
	goto L154
L154:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v697)+12))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v730+v705<<(uint(int32(2))%32))))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+12)))
	if v735 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	goto L151
L156:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v734)+8))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	v740 = int32(*(*int16)(unsafe.Add(mBase, uint32(v734))))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v738)+20))
	v745 = m.T0[v744].(func(*base.Module, int32, int32, int32) int32)(m, v738, v424, v739+v740-int32(1))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v757 = v705 + int32(1)
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v697)+4))
	if v757 < v758 {
		v705 = v757
		goto L154
	} else {
		goto L160
	}
L159:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v748 = int32(*(*int16)(unsafe.Add(mBase, uint32(v734))))
	*(*int32)(unsafe.Add(mBase, uint32(v747+v748<<(uint(int32(2))%32)-int32(4)))) = v745
	goto L158
L160:
	;
	goto L155
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = v439
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v794 == int32(0) {
		v859 = v439
		goto L120
	} else {
		goto L162
	}
L162:
	;
	v797 = int32(0)
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v798 <= v797 {
		v859 = v439
		goto L120
	} else {
		goto L163
	}
L163:
	;
	v802 = v797
	goto L164
L164:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v794)+12))
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v827+v802<<(uint(int32(2))%32))))
	v832 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v831)+12)))
	if v832 == int32(1) {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v859 = v439
	goto L120
L166:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v831)+8))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v439)+20))
	v837 = int32(*(*int16)(unsafe.Add(mBase, uint32(v831))))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v835)+20))
	v842 = m.T0[v841].(func(*base.Module, int32, int32, int32) int32)(m, v835, v424, v836+v837-int32(1))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v854 = v802 + int32(1)
	v855 = *(*int32)(unsafe.Add(mBase, uint32(v794)+4))
	if v854 < v855 {
		v802 = v854
		goto L164
	} else {
		goto L170
	}
L169:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v439)+16))
	v845 = int32(*(*int16)(unsafe.Add(mBase, uint32(v831))))
	*(*int32)(unsafe.Add(mBase, uint32(v844+v845<<(uint(int32(2))%32)-int32(4)))) = v842
	goto L168
L170:
	;
	goto L165
L171:
	;
	if v376 != 0 {
		goto L189
	} else {
		goto L190
	}
L172:
	;
	v886 = int32(0)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v887 <= v886 {
		goto L171
	} else {
		goto L173
	}
L173:
	;
	v891 = v886
	goto L174
L174:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v368)+12))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v916+v891<<(uint(int32(2))%32))))
	v921 = int32(*(*int16)(unsafe.Add(mBase, uint32(v859)+6)))
	if v921 < v920 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L184
	}
L176:
	;
	F_slot_getsomeattrs_int(m, v859, v920)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v859)+20))
	v929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v925+v920-int32(1)))))
	if v929 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L178
L180:
	;
	v933 = v891 + int32(1)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	if v934 <= v933 {
		goto L171
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	goto L175
L183:
	;
	v891 = v933
	goto L174
L184:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v945 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v944 + v945
	*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v35 + v936<<(uint(v945)%32) + v920*int32(100) - int32(76)
	F_errmsg(m, int32(167762), v29+int32(80))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_errtablecol(m, v33, v920)
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(517363), int32(6461), int32(415426))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v1033 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L190:
	;
	v995 = F_ExecRelGenVirtualNotNull(m, v373, v859, v55, v370)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	if v995 == int32(0) {
		goto L189
	} else {
		goto L192
	}
L192:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_errcode(m, int32(33575106))
	mBase = m.M
	v1006 = m.ExcPending
	if v1006 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1008 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v1007 + v1008
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v35 + v999<<(uint(v1008)%32) + v995*int32(100) - int32(76)
	F_errmsg(m, int32(167762), v29-int32(-64))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_errtablecol(m, v33, v995)
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(517363), int32(6481), int32(415426))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	if v159 != 0 {
		goto L217
	} else {
		goto L218
	}
L199:
	;
	v1036 = int32(0)
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+4))
	if v1037 <= v1036 {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1041 = v1036
	goto L201
L201:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+12))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1066+v1041<<(uint(int32(2))%32))))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	switch v1071 - int32(1) {
	case 0, 8:
		goto L203
	default:
		goto L204
	case 4:
		goto L205
	}
L202:
	;
	goto L198
L203:
	;
	v1120 = v1041 + int32(1)
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1033)+4))
	if v1120 < v1121 {
		v1041 = v1120
		goto L201
	} else {
		goto L216
	}
L204:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L1
	} else {
		goto L213
	}
L205:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+28))
	v1075 = F_ExecCheck(m, v1074, v424)
	mBase = m.M
	v1076 = m.ExcPending
	if v1076 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	if v1075 != 0 {
		goto L203
	} else {
		goto L207
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1085
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v1084 + int32(4)
	F_errmsg(m, int32(31990), v29+int32(48))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1070)))
	F_errtableconstraint(m, v33, v1095)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(517363), int32(6498), int32(415426))
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1070)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v1107
	F_errmsg_internal(m, int32(506121), v29+int32(32))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(517363), int32(6506), int32(415426))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
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
	goto L202
L217:
	;
	v1149 = F_ExecCheck(m, v159, v424)
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L1
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	if v51 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	if v1149 == int32(0) {
		goto L112
	} else {
		goto L221
	}
L221:
	;
	goto L219
L222:
	;
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1155)+80))
	m.T0[v1156].(func(*base.Module, int32, int32, int32, int32, int32))(m, v50, v859, v54, v52, v53)
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v424)+20))
	F_MemoryContextReset(m, v1159)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, _consts[45]))
	if v1163 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1165 = m.ExcPending
	if v1165 != 0 {
		goto L1
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v440)+36)) = v1167
	v1170 = *(*int32)(unsafe.Add(mBase, _consts[244]))
	if v1170 == int32(0) {
		goto L118
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	v1174 = int32(*(*uint8)(unsafe.Add(mBase, _consts[245])))
	if v1174&int32(1) != 0 {
		goto L118
	} else {
		goto L232
	}
L232:
	;
	goto L119
L233:
	;
	F_errmsg_internal(m, int32(352357), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(342438), int32(1034), int32(90875))
	mBase = m.M
	v1215 = m.ExcPending
	if v1215 != 0 {
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
	F_errcode(m, int32(67391682))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v1226 = v1224 + int32(4)
	if v1216 != int32(1) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v1226
	F_errmsg(m, int32(31928), v29+int32(16))
	mBase = m.M
	v1234 = m.ExcPending
	if v1234 != 0 {
		goto L1
	} else {
		goto L241
	}
L239:
	;
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v1226
	F_errmsg(m, int32(32053), v29)
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L1
	} else {
		goto L244
	}
L241:
	;
	F_errtable(m, v33)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(517363), int32(6523), int32(415426))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errtable(m, v33)
	mBase = m.M
	v1247 = m.ExcPending
	if v1247 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(517363), int32(6517), int32(415426))
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
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
	F_sequence_close(m, v33, int32(0))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L248
	}
L248:
	;
	if v51 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	F_FreeBulkInsertState(m, v53)
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	m.G0 = v29 + int32(144)
	return
L252:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v50)+188))
	if v1288 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	F_sequence_close(m, v50, int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L257
	}
L254:
	;
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v1288)+108))
	if v1291 == int32(0) {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	m.T0[v1291].(func(*base.Module, int32, int32))(m, v50, v52)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L256
	}
L256:
	;
	goto L253
L257:
	;
	goto L251
}
