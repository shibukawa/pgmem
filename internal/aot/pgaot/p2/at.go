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
	var v89 int64
	_ = v89
	var v92 int32
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
		goto L86
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L82
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L78
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
		goto L74
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
		goto L43
	} else {
		goto L44
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
	F_errmsg(m, int32(120627), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(494434), int32(7977), int32(303725))
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
	F_sequence_close(m, v65, int32(3))
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
	v89 = *(*int64)(unsafe.Add(mBase, _consts[284]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v89
	v92 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v92
	goto L15
L41:
	;
	goto L15
L42:
	;
	if l6 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v105 = v9
	goto L42
L44:
	;
	goto L45
L45:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v96 = F_find_inheritance_children(m, v94, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v96 == int32(0) {
		v105 = v9
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+119)))
	if v101 == int32(112) {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	v105 = int32(1)
	goto L42
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	v114 = F_ChooseConstraintName(m, v108+int32(4), l4, int32(301999), v112, int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	v116 = l3
	goto L51
L51:
	;
	v117 = F_makeString(m, l4)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v116 = v114
	goto L51
L53:
	;
	v119 = F_makeNotNullConstraint(m, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
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
		goto L55
	}
L55:
	;
	v131 = int32(0)
	v136 = F_AddRelationNewConstraints(m, l2, int32(0), v129, v131, l6^int32(1), v131, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v142 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v142 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v145 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v144, v24, v145, v145)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_set_attnotnull(m, l1, l2, v24, int32(1))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	if l5 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	goto L15
L63:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v155 = F_find_inheritance_children(m, v154, l7)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v155 == int32(0) {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v159 <= int32(0) {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v165 = int32(0)
	goto L67
L67:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v155)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v165<<(uint(int32(2))%32))))
	v181 = F_table_open(m, v179, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v187 = int32(1)
	F_ATExecSetNotNull(m, v13+int32(-16), l1, v181, v116, l4, v187, v187, l7)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_sequence_close(m, v181, int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v195 = v165 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if v195 < v196 {
		v165 = v195
		goto L67
	} else {
		goto L73
	}
L73:
	;
	goto L68
L74:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v237 + int32(4)
	F_errmsg(m, int32(71602), v15)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(494434), int32(7940), int32(303725))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l4
	F_errmsg(m, int32(711024), v13+int32(-48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(494434), int32(7947), int32(303725))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v276 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v33 + v276
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v275 + v276
	F_errmsg(m, int32(705912), v13+int32(-32))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(494434), int32(7964), int32(303725))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(240398), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errhint(m, int32(644316), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(494434), int32(8025), int32(303725))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecSetTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
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
	var v80 int64
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
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
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
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
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = F_relation_open(m, l0, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return
L2:
	;
	return
L3:
	;
	v17 = F_CheckRelationTableSpaceMove(m, v15, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+112))
	if v35 != 0 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v25 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v24, v25, v25, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v15, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v36 = F_relation_open(m, v35, l2)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	v44 = int32(0)
	v45 = v34
	goto L15
L15:
	;
	v46 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45)+118)))
	v47 = F_GetNewRelFileNumber(m, l1, int32(0), v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	v38 = F_RelationGetIndexList(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	F_relation_close(m, v36, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v44 = v38
	v45 = v42
	goto L15
L19:
	;
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = l1
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+119)))
	if v54 == int32(105) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_SetRelationTableSpace(m, v15, l1, v47)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L2
	} else {
		goto L70
	}
L21:
	;
	v58 = v13 + int32(72)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v61
	F_FlushRelationBuffers(m, v15)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+120))
	m.T0[v204].(func(*base.Module, int32, int32))(m, v15, v13-int32(-64))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L69
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+118)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v69
	v74 = F_RelationCreateStorage(m, v13+int32(48), v66, int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v102 = v76
	goto L28
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v80
	v84 = F_smgropen(m, v13+int32(32), v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L2
	} else {
		goto L29
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v105 = int32(*(*int8)(unsafe.Add(mBase, uint32(v104)+118)))
	F_RelationCopyStorage(m, v102, v74, int32(0), v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v84
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+72))
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v102 = v100
	goto L28
L31:
	;
	v96 = v88
	goto L33
L32:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)+76))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v84)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v84)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v90))) = v92
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v84)+72))
	v96 = v94
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+72)) = v96 + int32(1)
	goto L30
L34:
	;
	v109 = int32(1)
	goto L35
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v119 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F_RelationDropStorage(m, v15)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L2
	} else {
		goto L67
	}
L37:
	;
	v145 = v119
	goto L39
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v121
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v123
	v127 = F_smgropen(m, v13+int32(16), v120)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L2
	} else {
		goto L40
	}
L39:
	;
	v146 = F_smgrexists(m, v145, v109)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L2
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	if v131 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v145 = v143
	goto L39
L42:
	;
	v139 = v131
	goto L44
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v127)+76))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v133))) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)+72))
	v139 = v137
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+72)) = v139 + int32(1)
	goto L41
L45:
	;
	if v146 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	F_smgrcreate(m, v74, v109, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v194 = v109 + int32(1)
	if v194 != int32(4) {
		v109 = v194
		goto L35
	} else {
		goto L66
	}
L49:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+118)))
	if v152 != int32(112) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v163 != 0 {
		goto L57
	} else {
		goto L58
	}
L51:
	;
	if v109 != int32(3) {
		goto L50
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	F_log_smgrcreate(m, v13+int32(80), v109)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L2
	} else {
		goto L56
	}
L54:
	;
	if v152 != int32(117) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	goto L50
L57:
	;
	v187 = v163
	goto L59
L58:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v165
	v167 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v167
	v169 = F_smgropen(m, v13, v164)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L2
	} else {
		goto L60
	}
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v189 = int32(*(*int8)(unsafe.Add(mBase, uint32(v188)+118)))
	F_RelationCopyStorage(m, v187, v74, v109, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L65
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v169
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)+72))
	if v173 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v187 = v185
	goto L59
L62:
	;
	v181 = v173
	goto L64
L63:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+76))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v169)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v169)+72))
	v181 = v179
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v169)+72)) = v181 + int32(1)
	goto L61
L65:
	;
	goto L48
L66:
	;
	goto L36
L67:
	;
	F_smgrclose(m, v74)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	goto L20
L69:
	;
	goto L20
L70:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	if v220 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v223 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v222, v223, v223, v223)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L2
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v229 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v15)+40))
	if v231 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L73
L75:
	;
	F_relation_close(m, v15, int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L2
	} else {
		goto L82
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v229
	goto L78
L77:
	;
	goto L78
L78:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _consts[287]))
	if v236 <= int32(31) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	*(*int32)(unsafe.Add(mBase, _consts[287])) = v236 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v236<<(uint(int32(2))%32))+uint32(_consts[288]))) = v239
	goto L75
L80:
	;
	goto L81
L81:
	;
	v250 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[289])) = uint8(v250)
	goto L75
L82:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L2
	} else {
		goto L83
	}
L83:
	;
	if v35 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_ATExecSetTableSpace(m, v35, l1, l2)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L2
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if v44 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	F_list_free(m, v44)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L2
	} else {
		goto L95
	}
L89:
	;
	v262 = int32(0)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v263 <= v262 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v266 = v262
	goto L91
L91:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v266<<(uint(int32(2))%32))))
	F_ATExecSetTableSpace(m, v280, l1, l2)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L2
	} else {
		goto L93
	}
L92:
	;
	goto L88
L93:
	;
	v284 = v266 + int32(1)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v284 < v285 {
		v266 = v284
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
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
	var v73 int32
	_ = v73
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
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
	if v108 != 0 {
		goto L18
	} else {
		goto L19
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
	v73 = int32(0)
	goto L11
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v73<<(uint(int32(2))%32))))
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
	v91 = v73 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	if v91 < v92 {
		v73 = v91
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v17)+40))
	v228 = F_list_concat(m, v226, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L52
	}
L17:
	;
	v123 = v109
	v126 = v7
	goto L22
L18:
	;
	v109 = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v109 < v110 {
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v222 = v7
	goto L16
L21:
	;
	goto L20
L22:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v123<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	switch v137 - int32(14) {
	case 0:
		v155 = int32(9)
		goto L29
	default:
		v163 = l4
		goto L27
	case 2:
		goto L30
	case 7:
		goto L31
	case 11:
		v159 = int32(11)
		goto L28
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L49
	}
L24:
	;
	goto L23
L25:
	;
	v194 = v123 + int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v194 < v195 {
		v123 = v194
		v126 = v191
		goto L22
	} else {
		goto L48
	}
L26:
	;
	v184 = l0 + int32(16) + v159<<(uint(int32(2))%32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = F_lappend(m, v185, v136)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L47
	}
L27:
	;
	if v126 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	if l4 < v159 {
		goto L26
	} else {
		goto L39
	}
L29:
	;
	if v155 < l4 {
		goto L24
	} else {
		goto L38
	}
L30:
	;
	if v4 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v155 = int32(8)
	goto L29
L32:
	;
	v141 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v136)+29)) = uint8(v141)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v146 = v144 - int32(1)
	if base.Ui32(int32(7)) < base.Ui32(v146) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v155 = int32(10)
	goto L29
L36:
	;
	goto L37
L37:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v146<<(uint(int32(2))%32))+uint32(_consts[408])))
	v155 = v154
	goto L29
L38:
	;
	v159 = v155
	goto L28
L39:
	;
	v163 = v159
	goto L27
L40:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v166 == v137 {
		v191 = v136
		goto L25
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v163
	F_errmsg_internal(m, int32(469267), v17)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(494434), int32(5824), int32(429409))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v191 = v126
	goto L25
L48:
	;
	v222 = v191
	goto L16
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v155
	F_errmsg_internal(m, int32(469322), v17+int32(16))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(494434), int32(5804), int32(429409))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v228
	m.G0 = v17 + int32(48)
	return v222
}
