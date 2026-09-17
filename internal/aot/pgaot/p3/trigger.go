package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EnableDisableTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v4 = l3
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(128)
	m.G0 = v19
	v23 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v19+int32(32), int32(2), int32(3), int32(184), v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L64
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L60
	}
L6:
	;
	F_ScanKeyInit(m, v19+int32(80), int32(4), int32(3), int32(62), l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v45 = int32(1)
	goto L8
L8:
	;
	v48 = F_systable_beginscan(m, v23, int32(2701), int32(1), int32(0), v45, v19+int32(32))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v45 = int32(2)
	goto L8
L10:
	;
	v50 = F_systable_getnext(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v50 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v61 = v50
	v64 = v8
	v67 = v8
	goto L15
L13:
	;
	v209 = v8
	v212 = v8
	goto L14
L14:
	;
	F_systable_endscan(m, v48)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L50
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+22)))
	v72 = v70 + v71
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v209 = v191
	v212 = v194
	goto L14
L17:
	;
	v197 = F_systable_getnext(m, v48)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L48
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if l2 != v73 {
		v191 = v64
		v194 = v67
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+83)))
	if v75 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	if l4 != 0 {
		v191 = v64
		v194 = v67
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+82)))
	if v4&int32(255) != v82 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v78 = F_superuser(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v78 == int32(0) {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v84 = F_heap_copytuple(m, v61)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v99 = v67
	goto L30
L30:
	;
	if l5 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86+v87)+82)) = uint8(v4)
	F_CatalogTupleUpdate(m, v23, v84+int32(4), v84)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v84)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v99 = int32(1)
	goto L30
L34:
	;
	v169 = int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableTrigger[0]))
	if v171 == int32(0) {
		v191 = v169
		v194 = v99
		goto L17
	} else {
		goto L46
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+119)))
	if v103 != int32(112) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+80)))
	if v106&int32(1) == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v112 = F_RelationGetPartitionDesc(m, l0, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v114 <= int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v125 = int32(0)
	goto L40
L40:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v125<<(uint(int32(2))%32))))
	v139 = F_relation_open(m, v138, l6)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L34
L42:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	F_EnableDisableTrigger(m, v139, int32(0), v142, v4, l4, int32(1), l6)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_relation_close(m, v139, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v150 = v125 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v150 < v151 {
		v125 = v150
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v176 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v175, v176, v176, v176)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v191 = v169
	v194 = v99
	goto L17
L48:
	;
	if v197 != 0 {
		v61 = v197
		v64 = v191
		v67 = v194
		goto L15
	} else {
		goto L49
	}
L49:
	;
	goto L16
L50:
	;
	F_relation_close(m, v23, int32(3))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v209 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v221 = int32(0)
	goto L54
L53:
	;
	v221 = l1
	goto L54
L54:
	;
	if v221 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v212 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	m.G0 = v19 + int32(128)
	return
L59:
	;
	goto L58
L60:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v72 + int32(12)
	F_errmsg(m, int32(_a_F_EnableDisableTrigger_0), v19+int32(16))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_EnableDisableTrigger_1), int32(1778), int32(_a_F_EnableDisableTrigger_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v254 + int32(4)
	F_errmsg(m, int32(_a_F_EnableDisableTrigger_3), v19)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_EnableDisableTrigger_1), int32(1839), int32(_a_F_EnableDisableTrigger_2))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_TriggerEnabled(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	v8 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[0]))
	if v17 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return v157
L2:
	;
	if l3&int32(3) != int32(2) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	switch v15 - int32(68) {
	case 0, 11:
		v157 = v8
		goto L1
	default:
		goto L2
	}
L4:
	;
	goto L5
L5:
	;
	v23 = v15 - int32(68)
	if base.B2i32(v23 == int32(0))|base.B2i32(v23 == int32(14)) != 0 {
		v157 = v8
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v73 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+36)))
	if v34 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v44 = v8
	goto L10
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47+v44<<(uint(int32(1))%32)))))
	v54 = F_bms_is_member(m, v51+int32(7), l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v157 = int32(0)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	if v54 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v59 = v44 + int32(1)
	v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+36)))
	if v59 < v60 {
		v44 = v59
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v157 = int32(1)
	goto L1
L17:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v81 = base.I32_div_s(l2-v78, int32(15))
	v82 = v76 + v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v86 = int32(_a_F_TriggerEnabled_0)
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v89
	v91 = F_stringToNode(m, v73)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v118 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v95 = F_expand_generated_columns_in_expr(m, v91, v93, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v99 = F_expand_generated_columns_in_expr(m, v95, v97, int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	F_ChangeVarNodes(m, v99, int32(1), int32(-1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_ChangeVarNodes(m, v99, int32(2), int32(-2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v109 = F_make_ands_implicit(m, v99)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v111 = F_ExecPrepareQual(m, v109, l0)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v111
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v87
	goto L20
L28:
	;
	v121 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L12
	} else {
		goto L31
	}
L29:
	;
	v123 = v118
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = l5
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v126 == int32(0) {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	v123 = v121
	goto L30
L32:
	;
	v130 = int32(_a_F_TriggerEnabled_0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v133
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v138 = m.T0[v137].(func(*base.Module, int32, int32, int32) int32)(m, v126, v123, v13+int32(15))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v131
	if v138 == int32(0) {
		v157 = int32(0)
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L16
}
