package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ReindexIsProcessingIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexIsProcessingIndex[0]))
	if v3 != l0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ReindexIsProcessingIndex[1]))
	v7 = int32(0)
	if v6 == v7 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v47 = int32(1)
	goto L3
L3:
	;
	return v47
L4:
	;
	v47 = v45
	goto L3
L5:
	;
	v45 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v13 <= int32(0) {
		v39 = v7
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v45 = v39
	goto L4
L9:
	;
	v16 = int32(0)
	if v16 < v13 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v19 = v13
	goto L12
L11:
	;
	v19 = v16
	goto L12
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v22 = int32(0)
	goto L13
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v20+v22<<(uint(int32(2))%32))))
	v31 = base.B2i32(v30 == l0)
	if v30 == l0 {
		v39 = v31
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v39 = v31
	goto L8
L15:
	;
	v33 = v22 + int32(1)
	if v33 != v19 {
		v22 = v33
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
}
func F_reindex_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v18&int32(4) != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L9
	} else {
		goto L87
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L9
	} else {
		goto L84
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L9
	} else {
		goto L81
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L9
	} else {
		goto L77
	}
L5:
	;
	if v29 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v22 = F_try_table_open(m, l1, int32(5))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v27 = F_table_open(m, l1, int32(5))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	v29 = v22
	goto L5
L11:
	;
	v29 = v27
	goto L5
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
	if v31 == int32(112) {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	v256 = int32(0)
	goto L14
L14:
	;
	m.G0 = v16 + int32(32)
	return v256
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+112))
	v35 = F_RelationGetIndexList(m, v29)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v38 = l2 & int32(2)
	if v38 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[0]))
	if v40 != 0 {
		goto L3
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v63 = int32(0)
	if base.B2i32(l2&int32(1) == v63)|base.B2i32(v34 == v63) == v63 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[1]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+72))
	if v44 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v47&int32(1) != 0 {
		goto L2
	} else {
		goto L25
	}
L22:
	;
	v47 = int32(1)
	goto L24
L23:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+76)))
	v47 = v46
	goto L24
L24:
	;
	goto L21
L25:
	;
	v51 = F_list_copy(m, v35)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[0])) = v51
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[1]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+28))
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[2])) = v57
	F_CommandCounterIncrement(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L19
L29:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v70 & int32(-5)
	v78 = F_reindex_relation(m, l0, v34, l2, v16+int32(24))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v81 = int32(0)
	goto L31
L31:
	;
	if l2&int32(8) != 0 {
		v90 = int32(117)
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v81 = v78
	goto L31
L33:
	;
	v91 = int32(0)
	if v35 == v91 {
		v230 = v91
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if l2&int32(16) != 0 {
		v90 = int32(112)
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+118)))
	v90 = v89
	goto L33
L36:
	;
	F_relation_close(m, v29, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L76
	}
L37:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v95 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v230 = int32(1)
	goto L36
L39:
	;
	goto L40
L40:
	;
	v99 = int32(0)
	v106 = v99
	v114 = int32(1)
	goto L41
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v106<<(uint(int32(2))%32))))
	v123 = F_get_rel_namespace(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L9
	} else {
		goto L45
	}
L42:
	;
	v230 = v220
	goto L36
L43:
	;
	v220 = int32(1)
	v222 = v106 + v220
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v222 < v223 {
		v106 = v222
		v114 = v219
		goto L41
	} else {
		goto L75
	}
L44:
	;
	F_reindex_index(m, l0, v122, base.B2i32(l2&int32(4) == v99), base.I32_extend8_s(v90), l3)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L69
	}
L45:
	;
	if v123 != int32(99) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v129 == int32(0) {
		goto L44
	} else {
		goto L50
	}
L47:
	;
	v127 = F_isTempToastNamespace(m, v123)
	mBase = m.M
	v129 = v127
	goto L49
L48:
	;
	v129 = int32(1)
	goto L49
L49:
	;
	goto L46
L50:
	;
	v132 = F_get_index_isvalid(m, v122)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	if v132 != 0 {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v136 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	if v136 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v38 == int32(0) {
		v219 = v114
		goto L43
	} else {
		goto L62
	}
L57:
	;
	v141 = F_get_namespace_name(m, v123)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v143 = F_get_rel_name(m, v122)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v141
	F_errmsg(m, int32(_a_F_reindex_relation_0), v16+int32(16))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_reindex_relation_1), int32(4061), int32(_a_F_reindex_relation_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	goto L56
L62:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[1]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+72))
	if v163 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v166&int32(1) != 0 {
		goto L1
	} else {
		goto L67
	}
L64:
	;
	v166 = int32(1)
	goto L66
L65:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+76)))
	v166 = v165
	goto L66
L66:
	;
	goto L63
L67:
	;
	v169 = int32(_a_F_reindex_relation_3)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[0]))
	v172 = F_list_delete_ptr(m, v171, v122)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[0])) = v172
	v219 = v114
	goto L43
L69:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[3]))
	if v183 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v219 = v114 + int32(1)
	goto L43
L72:
	;
	goto L71
L73:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_reindex_relation[4])))
	if v187&int32(1) == int32(0) {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v192 = int32(_a_F_reindex_relation_4)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[5]))
	v195 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[5])) = v194 + v195
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v198 + v195
	*(*int64)(unsafe.Add(mBase, uint32(v183+int32(56))+232)) = base.I64_extend_i32_s(v114)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v206 + v195
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_reindex_relation[5])) = v212 - v195
	goto L72
L75:
	;
	goto L42
L76:
	;
	v256 = v230 | v81
	goto L14
L77:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+68))
	v267 = F_get_namespace_name(m, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v269 + int32(4)
	F_errmsg_internal(m, int32(_a_F_reindex_relation_5), v16)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_reindex_relation_1), int32(3980), int32(_a_F_reindex_relation_2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	F_errmsg_internal(m, int32(_a_F_reindex_relation_6), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L9
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_reindex_relation_1), int32(_a_F_reindex_relation_7), int32(_a_F_reindex_relation_8))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L9
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
	F_errmsg_internal(m, int32(_a_F_reindex_relation_9), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_reindex_relation_1), int32(_a_F_reindex_relation_10), int32(_a_F_reindex_relation_8))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errmsg_internal(m, int32(_a_F_reindex_relation_9), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_reindex_relation_1), int32(_a_F_reindex_relation_11), int32(_a_F_reindex_relation_12))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L9
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
