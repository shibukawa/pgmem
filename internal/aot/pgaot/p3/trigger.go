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
	var v30 int64
	_ = v30
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v4 = l3
	v8 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(144)
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
	v30 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
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
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L64
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L60
	}
L6:
	;
	F_ScanKeyInit(m, v19+int32(88), int32(4), int32(3), int32(62), base.I64_extend_i32_u(l1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v46 = int32(1)
	goto L8
L8:
	;
	v49 = F_systable_beginscan(m, v23, int32(2701), int32(1), int32(0), v46, v19+int32(32))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v46 = int32(2)
	goto L8
L10:
	;
	v51 = F_systable_getnext(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v51 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v62 = v51
	v65 = v8
	v68 = v8
	goto L15
L13:
	;
	v210 = v8
	v213 = v8
	goto L14
L14:
	;
	F_systable_endscan(m, v49)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L50
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v73 = v71 + v72
	if l2 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v210 = v192
	v213 = v195
	goto L14
L17:
	;
	v198 = F_systable_getnext(m, v49)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L48
	}
L18:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	if l2 != v74 {
		v192 = v65
		v195 = v68
		goto L17
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+83)))
	if v76 == int32(1) {
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
		v192 = v65
		v195 = v68
		goto L17
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+82)))
	if v4&int32(255) != v83 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v79 = F_superuser(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v79 == int32(0) {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v85 = F_heap_copytuple(m, v62)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v100 = v68
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
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v87+v88)+82)) = uint8(v4)
	F_CatalogTupleUpdate(m, v23, v85+int32(4), v85)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v85)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v100 = int32(1)
	goto L30
L34:
	;
	v170 = int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableTrigger[0]))
	if v172 == int32(0) {
		v192 = v170
		v195 = v100
		goto L17
	} else {
		goto L46
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+119)))
	if v104 != int32(112) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+80)))
	if v107&int32(1) == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v113 = F_RelationGetPartitionDesc(m, l0, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v115 <= int32(0) {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	v126 = int32(0)
	goto L40
L40:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135+v126<<(uint(int32(2))%32))))
	v140 = F_relation_open(m, v139, l6)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L34
L42:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	F_EnableDisableTrigger(m, v140, int32(0), v143, v4, l4, int32(1), l6)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_relation_close(m, v140, int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v151 = v126 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v151 < v152 {
		v126 = v151
		goto L40
	} else {
		goto L45
	}
L45:
	;
	goto L41
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v177 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2620), v176, v177, v177, v177)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v192 = v170
	v195 = v100
	goto L17
L48:
	;
	if v198 != 0 {
		v62 = v198
		v65 = v192
		v68 = v195
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v210 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v222 = int32(0)
	goto L54
L53:
	;
	v222 = l1
	goto L54
L54:
	;
	if v222 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v213 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	m.G0 = v19 + int32(144)
	return
L59:
	;
	goto L58
L60:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v73 + int32(12)
	F_errmsg(m, int32(_a_F_EnableDisableTrigger_0), v19+int32(16))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_EnableDisableTrigger_1), int32(1793), int32(_a_F_EnableDisableTrigger_2))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v255 + int32(4)
	F_errmsg(m, int32(_a_F_EnableDisableTrigger_3), v19)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_EnableDisableTrigger_1), int32(1854), int32(_a_F_EnableDisableTrigger_2))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
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
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v160 int32
	_ = v160
	v8 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+14)))
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[0]))
	if v18 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v160
L2:
	;
	if l3&int32(3) != int32(2) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	switch v16 - int32(68) {
	case 0, 11:
		v160 = v8
		goto L1
	default:
		goto L2
	}
L4:
	;
	goto L5
L5:
	;
	v24 = v16 - int32(68)
	if base.B2i32(v24 == int32(0))|base.B2i32(v24 == int32(14)) != 0 {
		v160 = v8
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v76 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+36)))
	if v35 <= int32(0) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v45 = v8
	goto L10
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49+v45<<(uint(int32(1))%32)))))
	v56 = F_bms_is_member(m, v53+int32(7), l4)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v160 = int32(0)
	goto L1
L12:
	;
	return int32(0)
L13:
	;
	if v56 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v61 = v45 + int32(1)
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+36)))
	if v61 < v62 {
		v45 = v61
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v160 = int32(1)
	goto L1
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v84 = base.I32_div_s(l2-v81, int32(15))
	v85 = v79 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v89 = int32(_a_F_TriggerEnabled_0)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v92
	v94 = F_stringToNode(m, v76)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v121 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v98 = F_expand_generated_columns_in_expr(m, v94, v96, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v102 = F_expand_generated_columns_in_expr(m, v98, v100, int32(2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	F_ChangeVarNodes(m, v102, int32(1), int32(-1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	F_ChangeVarNodes(m, v102, int32(2), int32(-2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v112 = F_make_ands_implicit(m, v102)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	v114 = F_ExecPrepareQual(m, v112, l0)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v114
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v90
	goto L20
L28:
	;
	v124 = F_MakePerTupleExprContext(m, l0)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L31
	}
L29:
	;
	v126 = v121
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+12)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v126)+8)) = l5
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v129 == int32(0) {
		goto L16
	} else {
		goto L32
	}
L31:
	;
	v126 = v124
	goto L30
L32:
	;
	v133 = int32(_a_F_TriggerEnabled_0)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v136
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v129)+24))
	v141 = m.T0[v140].(func(*base.Module, int32, int32, int32) int64)(m, v129, v126, v14+int32(15))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_TriggerEnabled[1])) = v134
	if v141 == int64(0) {
		v160 = int32(0)
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L16
}
