package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecSetNotNull(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
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
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	v9 = int32(0)
	v13 = m.G0
	v15 = v13 + int32(-64)
	m.G0 = v15
	F_check_stack_depth(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
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
	F_ATSimplePermissions(m, int32(16), l2, int32(289))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v24 = F_get_attnum(m, v23, l4)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L84
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L80
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L76
	}
L10:
	;
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v24 <= int32(0) {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L72
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v29 = F_findNotNullConstraintAttnum(m, v28, v24)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	m.G0 = v15 - int32(-64)
	return
L16:
	;
	if v29 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	if l5 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	if l5 != 0 {
		v105 = v9
		goto L42
	} else {
		goto L43
	}
L20:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+106)))
	if v34&int32(1) != 0 {
		goto L8
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if l6 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+76)))
	if v80 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L25:
	;
	v65 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L35
	}
L26:
	;
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+104)))
	v39 = v37 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+104)) = uint16(v39)
	if base.I32_extend16_s(v39) == v39 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+103)))
	if v59 != 0 {
		goto L24
	} else {
		goto L34
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_ATExecSetNotNull_0), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ATExecSetNotNull_1), int32(_a_F_ATExecSetNotNull_2), int32(_a_F_ATExecSetNotNull_3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v60 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+103)) = uint8(v60)
	goto L25
L35:
	;
	F_CatalogTupleUpdate(m, v65, v29+int32(4), v29)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	F_relation_close(m, v65, int32(3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	goto L15
L38:
	;
	F_ATExecValidateConstraint(m, l0, l1, l2, v33+int32(4), l5, int32(0), l7)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetNotNull[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecSetNotNull[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v92
	goto L15
L41:
	;
	goto L15
L42:
	;
	if l6 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v96 = F_find_inheritance_children(m, v94, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v96 == int32(0) {
		v105 = v9
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+119)))
	if v101 == int32(112) {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v105 = int32(1)
	goto L42
L47:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	v114 = F_ChooseConstraintName(m, v108+int32(4), l4, int32(_a_F_ATExecSetNotNull_4), v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	v116 = l3
	goto L49
L49:
	;
	v117 = F_makeString(m, l4)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v116 = v114
	goto L49
L51:
	;
	v119 = F_makeNotNullConstraint(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v116
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+17)) = uint8(v105)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v119
	v129 = F_list_make1_impl(m, int32(1), v13+int32(-36))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v131 = int32(0)
	v136 = F_AddRelationNewConstraints(m, l2, int32(0), v129, v131, l6^int32(1), v131, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetNotNull[2]))
	if v142 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v145 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v144, v24, v145, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	F_set_attnotnull(m, l1, l2, v24, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	if l5 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	goto L15
L61:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v155 = F_find_inheritance_children(m, v154, l7)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v155 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v159 <= int32(0) {
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v165 = int32(0)
	goto L65
L65:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v165<<(uint(int32(2))%32))))
	v181 = F_table_open(m, v179, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	goto L60
L67:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v187 = int32(1)
	F_ATExecSetNotNull(m, v13+int32(-16), l1, v181, v116, l4, v187, v187, l7)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_relation_close(m, v181, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v195 = v165 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v195 < v196 {
		v165 = v195
		goto L65
	} else {
		goto L71
	}
L71:
	;
	goto L66
L72:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v237 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecSetNotNull_5), v15)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ATExecSetNotNull_1), int32(_a_F_ATExecSetNotNull_6), int32(_a_F_ATExecSetNotNull_3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
	F_errmsg(m, int32(_a_F_ATExecSetNotNull_7), v13+int32(-48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_ATExecSetNotNull_1), int32(_a_F_ATExecSetNotNull_8), int32(_a_F_ATExecSetNotNull_3))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v276 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v33 + v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v275 + v276
	F_errmsg(m, int32(_a_F_ATExecSetNotNull_9), v13+int32(-32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_ATExecSetNotNull_1), int32(_a_F_ATExecSetNotNull_10), int32(_a_F_ATExecSetNotNull_3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_ATExecSetNotNull_11), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errhint(m, int32(_a_F_ATExecSetNotNull_12), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_ATExecSetNotNull_1), int32(_a_F_ATExecSetNotNull_13), int32(_a_F_ATExecSetNotNull_3))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecSetTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v48 int64
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int64
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int64
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v422 int32
	_ = v422
	v10 = m.G0
	v12 = v10 - int32(160)
	m.G0 = v12
	v14 = F_relation_open(m, l0, l2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(160)
	return
L2:
	;
	return
L3:
	;
	v16 = F_CheckRelationTableSpaceMove(m, v14, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v16 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetTableSpace[0]))
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+112))
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v24 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v23, v24, v24, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v14, int32(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L2
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	goto L1
L13:
	;
	v35 = F_relation_open(m, v34, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v43 = int32(0)
	v44 = v33
	goto L15
L15:
	;
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v44)+118)))
	v46 = F_GetNewRelFileNumber(m, l1, int32(0), v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v37 = F_RelationGetIndexList(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	F_relation_close(m, v35, l2)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v43 = v37
	v44 = v41
	goto L15
L19:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+128)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = l1
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
	if v53 == int32(105) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_SetRelationTableSpace(m, v14, l1, v46)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L2
	} else {
		goto L115
	}
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+152)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+144)) = v58
	F_FlushRelationBuffers(m, v14)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v14)+188))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)+120))
	m.T0[v338].(func(*base.Module, int32, int32))(m, v14, v12+int32(128))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L2
	} else {
		goto L114
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62)+118)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+136))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v12)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+112)) = v66
	v71 = F_RelationCreateStorage(m, v12+int32(112), v63, int32(1))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v73 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v99 = v73
	goto L28
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+96)) = v77
	v81 = F_smgropen(m, v12+int32(96), v74)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L29
	}
L28:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101)+118)))
	F_RelationCopyStorage(m, v99, v71, int32(0), v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v81
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+72))
	if v85 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v99 = v97
	goto L28
L31:
	;
	v93 = v85
	goto L33
L32:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+76))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+72))
	v93 = v91
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+72)) = v93 + int32(1)
	goto L30
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v105 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v131 = v105
	goto L37
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v109
	v113 = F_smgropen(m, v12+int32(80), v106)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L2
	} else {
		goto L38
	}
L37:
	;
	v133 = F_smgrexists(m, v131, int32(1))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L43
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v113
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+72))
	if v117 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v131 = v129
	goto L37
L40:
	;
	v125 = v117
	goto L42
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v113)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v113)+72))
	v125 = v123
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+72)) = v125 + int32(1)
	goto L39
L43:
	;
	if v133 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_smgrcreate(m, v71, int32(1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v181 != 0 {
		goto L61
	} else {
		goto L62
	}
L47:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+118)))
	if v140 == int32(112) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_log_smgrcreate(m, v12+int32(144), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v148 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v174 = v148
	goto L54
L53:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v150
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v152
	v156 = F_smgropen(m, v12-int32(-64), v149)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L55
	}
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v177 = int32(*(*int8)(unsafe.Add(mBase, uint32(v176)+118)))
	F_RelationCopyStorage(m, v174, v71, int32(1), v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L2
	} else {
		goto L60
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v156
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+72))
	if v160 != 0 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v174 = v172
	goto L54
L57:
	;
	v168 = v160
	goto L59
L58:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156)+76))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v156)+72))
	v168 = v166
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156)+72)) = v168 + int32(1)
	goto L56
L60:
	;
	goto L46
L61:
	;
	v207 = v181
	goto L63
L62:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v183
	v185 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v185
	v189 = F_smgropen(m, v12+int32(48), v182)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L64
	}
L63:
	;
	v209 = F_smgrexists(m, v207, int32(2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L69
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v189
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v189)+72))
	if v193 != 0 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v207 = v205
	goto L63
L66:
	;
	v201 = v193
	goto L68
L67:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+76))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v194)+4)) = v195
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v189)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v197
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v189)+72))
	v201 = v199
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189)+72)) = v201 + int32(1)
	goto L65
L69:
	;
	if v209 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_smgrcreate(m, v71, int32(2), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L2
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v257 != 0 {
		goto L87
	} else {
		goto L88
	}
L73:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+118)))
	if v216 == int32(112) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_log_smgrcreate(m, v12+int32(144), int32(2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v224 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L76
L78:
	;
	v250 = v224
	goto L80
L79:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v226
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v228
	v232 = F_smgropen(m, v12+int32(32), v225)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L2
	} else {
		goto L81
	}
L80:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v253 = int32(*(*int8)(unsafe.Add(mBase, uint32(v252)+118)))
	F_RelationCopyStorage(m, v250, v71, int32(2), v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L86
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v232
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)+72))
	if v236 != 0 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v250 = v248
	goto L80
L83:
	;
	v244 = v236
	goto L85
L84:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+4)) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)+72))
	v244 = v242
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+72)) = v244 + int32(1)
	goto L82
L86:
	;
	goto L72
L87:
	;
	v283 = v257
	goto L89
L88:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v259
	v261 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v261
	v265 = F_smgropen(m, v12+int32(16), v258)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L2
	} else {
		goto L90
	}
L89:
	;
	v285 = F_smgrexists(m, v283, int32(3))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L95
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v265
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v265)+72))
	if v269 != 0 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v283 = v281
	goto L89
L92:
	;
	v277 = v269
	goto L94
L93:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v265)+76))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v265)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v270)+4)) = v271
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v265)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v271))) = v273
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v265)+72))
	v277 = v275
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v265)+72)) = v277 + int32(1)
	goto L91
L95:
	;
	if v285 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_smgrcreate(m, v71, int32(3), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L2
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_RelationDropStorage(m, v14)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L2
	} else {
		goto L112
	}
L99:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+118)))
	switch v292 - int32(112) {
	case 0, 5:
		goto L101
	default:
		goto L100
	}
L100:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v300 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	F_log_smgrcreate(m, v12+int32(144), int32(3))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v324 = v300
	goto L105
L104:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v302
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v304
	v306 = F_smgropen(m, v12, v301)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L2
	} else {
		goto L106
	}
L105:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v327 = int32(*(*int8)(unsafe.Add(mBase, uint32(v326)+118)))
	F_RelationCopyStorage(m, v324, v71, int32(3), v327)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L2
	} else {
		goto L111
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v306
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v306)+72))
	if v310 != 0 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v324 = v322
	goto L105
L108:
	;
	v318 = v310
	goto L110
L109:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v306)+76))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v311)+4)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v306)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v306)+72))
	v318 = v316
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+72)) = v318 + int32(1)
	goto L107
L111:
	;
	goto L98
L112:
	;
	F_smgrclose(m, v71)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L2
	} else {
		goto L113
	}
L113:
	;
	goto L20
L114:
	;
	goto L20
L115:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetTableSpace[0]))
	if v346 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v349 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v348, v349, v349, v349)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L2
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v355 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if v357 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	goto L118
L120:
	;
	F_relation_close(m, v14, int32(0))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L2
	} else {
		goto L127
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v355
	goto L123
L122:
	;
	goto L123
L123:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecSetTableSpace[1]))
	if v362 <= int32(31) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_ATExecSetTableSpace[1])) = v362 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v362<<(uint(int32(2))%32))+uint32(_c_F_ATExecSetTableSpace[2]))) = v365
	goto L120
L125:
	;
	goto L126
L126:
	;
	v376 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_ATExecSetTableSpace[3])) = uint8(v376)
	goto L120
L127:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	if v34 != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_ATExecSetTableSpace(m, v34, l1, l2)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	if v43 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	goto L131
L133:
	;
	F_list_free(m, v43)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L2
	} else {
		goto L140
	}
L134:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v388 <= int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v392 = int32(0)
	goto L136
L136:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v401+v392<<(uint(int32(2))%32))))
	F_ATExecSetTableSpace(m, v405, l1, l2)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L138
	}
L137:
	;
	goto L133
L138:
	;
	v409 = v392 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if v409 < v410 {
		v392 = v409
		goto L136
	} else {
		goto L139
	}
L139:
	;
	goto L137
L140:
	;
	goto L1
}
func F_ATParseTransformCmd(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	v4 = l3
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v20 = F_palloc0(m, int32(20))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(146)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v28 = F_get_namespace_name(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v33 = F_pstrdup(m, v30+int32(4))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_makeRangeVar(m, v28, v33, int32(-1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v36
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+16)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = l2
	v46 = F_list_make1_impl(m, int32(1), v17+int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+16)) = uint8(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v46
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v59 = F_transformAlterTableStmt(m, v53, v20, v54, v17+int32(44), v17+int32(40))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
	if v61 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	if v108 == int32(0) {
		v207 = v7
		goto L17
	} else {
		goto L18
	}
L9:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v64 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v75 = int32(0)
	goto L11
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v75<<(uint(int32(2))%32))))
	F_ProcessUtilityForAlterTable(m, v85, l5)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v91 = v75 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v91 < v92 {
		v75 = v91
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L47
	}
L17:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v212 = F_list_concat(m, v210, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L46
	}
L18:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v111 <= int32(0) {
		v207 = v7
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v126 = v7
	v127 = v7
	goto L20
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v126<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	switch v137 - int32(14) {
	case 0:
		v154 = int32(9)
		goto L26
	default:
		v160 = l4
		goto L24
	case 2:
		goto L27
	case 7:
		goto L28
	case 11:
		v158 = int32(11)
		goto L25
	}
L21:
	;
	v207 = v191
	goto L17
L22:
	;
	v193 = v126 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v193 < v194 {
		v126 = v193
		v127 = v191
		goto L20
	} else {
		goto L45
	}
L23:
	;
	v183 = l0 + int32(16) + v158<<(uint(int32(2))%32)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v185 = F_lappend(m, v184, v136)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L44
	}
L24:
	;
	if v127 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L25:
	;
	if l4 < v158 {
		goto L23
	} else {
		goto L36
	}
L26:
	;
	if v154 < l4 {
		goto L16
	} else {
		goto L35
	}
L27:
	;
	if v4 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v154 = int32(8)
	goto L26
L29:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+29)) = uint8(v141)
	goto L31
L30:
	;
	goto L31
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v146 = v144 - int32(1)
	if base.Ui32(int32(7)) < base.Ui32(v146) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v154 = int32(10)
	goto L26
L33:
	;
	goto L34
L34:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v146<<(uint(int32(2))%32))+uint32(_c_F_ATParseTransformCmd[0])))
	v154 = v152
	goto L26
L35:
	;
	v158 = v154
	goto L25
L36:
	;
	v160 = v158
	goto L24
L37:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v165 == v137 {
		v191 = v136
		goto L22
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
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v160
	F_errmsg_internal(m, int32(_a_F_ATParseTransformCmd_0), v17)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_ATParseTransformCmd_1), int32(_a_F_ATParseTransformCmd_2), int32(_a_F_ATParseTransformCmd_3))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v185
	v191 = v127
	goto L22
L45:
	;
	goto L21
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v212
	m.G0 = v17 + int32(48)
	return v207
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v154
	F_errmsg_internal(m, int32(_a_F_ATParseTransformCmd_4), v17+int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_ATParseTransformCmd_1), int32(_a_F_ATParseTransformCmd_5), int32(_a_F_ATParseTransformCmd_3))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
