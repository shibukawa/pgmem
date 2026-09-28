package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetAllPublicationRelations(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l1 == int32(114) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = int32(1)
	v19 = F_get_publication_relations(m, l0, l2^v16, v16)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v23 = v4
	goto L3
L3:
	;
	v26 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	v23 = v19
	goto L3
L6:
	;
	F_ScanKeyInit(m, v12, int32(18), int32(3), int32(61), base.I64_extend_i32_s(l1))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v35 = F_table_beginscan_catalog(m, v26, int32(1), v12)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v37 = F_heap_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v37 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = v37
	v44 = v4
	goto L13
L11:
	;
	v117 = v4
	goto L12
L12:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+188))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	m.T0[v123].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L43
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
	v50 = v48 + v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+119)))
	switch v52 - int32(112) {
	case 0, 2:
		goto L16
	case 1:
		v109 = v44
		goto L15
	default:
		goto L17
	}
L14:
	;
	v117 = v109
	goto L12
L15:
	;
	v110 = F_heap_getnext(m, v35)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L41
	}
L16:
	;
	goto L19
L17:
	;
	if v52 != int32(83) {
		v109 = v44
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if base.B2i32(base.Ui32(v51) < base.Ui32(int32(_a_F_GetAllPublicationRelations_0)))|base.B2i32(base.Ui32(v51) < base.Ui32(int32(_a_F_GetAllPublicationRelations_1))) != 0 {
		v109 = v44
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+118)))
	if v62 != int32(112) {
		v109 = v44
		goto L15
	} else {
		goto L21
	}
L21:
	;
	if l2 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+131)))
	if v65&int32(1) != 0 {
		v109 = v44
		goto L15
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v68 = int32(0)
	if v23 == v68 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L24
L26:
	;
	if v106 != 0 {
		v109 = v44
		goto L15
	} else {
		goto L39
	}
L27:
	;
	v106 = int32(0)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v74 <= int32(0) {
		v100 = v68
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v106 = v100
	goto L26
L31:
	;
	v77 = int32(0)
	if v77 < v74 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v80 = v74
	goto L34
L33:
	;
	v80 = v77
	goto L34
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v83 = int32(0)
	goto L35
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81+v83<<(uint(int32(2))%32))))
	v92 = base.B2i32(v91 == v51)
	if v91 == v51 {
		v100 = v92
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v100 = v92
	goto L30
L37:
	;
	v94 = v83 + int32(1)
	if v94 != v80 {
		v83 = v94
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v107 = F_lappend_oid(m, v44, v51)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v109 = v107
	goto L15
L41:
	;
	if v110 != 0 {
		v40 = v110
		v44 = v109
		goto L13
	} else {
		goto L42
	}
L42:
	;
	goto L14
L43:
	;
	if l2 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_ScanKeyInit(m, v12, int32(18), int32(3), int32(61), int64(112))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L47
	}
L45:
	;
	v227 = v117
	goto L46
L46:
	;
	F_relation_close(m, v26, int32(1))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L4
	} else {
		goto L81
	}
L47:
	;
	v133 = F_table_beginscan_catalog(m, v26, int32(1), v12)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v135 = F_heap_getnext(m, v133)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v135 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v138 = v135
	v142 = v117
	goto L53
L51:
	;
	v213 = v117
	goto L52
L52:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+188))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	m.T0[v219].(func(*base.Module, int32))(m, v133)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L80
	}
L53:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+22)))
	v148 = v146 + v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+119)))
	switch v150 - int32(112) {
	case 0, 2:
		goto L56
	case 1:
		v205 = v142
		goto L55
	default:
		goto L57
	}
L54:
	;
	v213 = v205
	goto L52
L55:
	;
	v206 = F_heap_getnext(m, v133)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L78
	}
L56:
	;
	goto L59
L57:
	;
	if v150 != int32(83) {
		v205 = v142
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	if base.B2i32(base.Ui32(v149) < base.Ui32(int32(_a_F_GetAllPublicationRelations_0)))|base.B2i32(base.Ui32(v149) < base.Ui32(int32(_a_F_GetAllPublicationRelations_1))) != 0 {
		v205 = v142
		goto L55
	} else {
		goto L60
	}
L60:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+118)))
	if v160 != int32(112) {
		v205 = v142
		goto L55
	} else {
		goto L61
	}
L61:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+131)))
	if v163 != 0 {
		v205 = v142
		goto L55
	} else {
		goto L62
	}
L62:
	;
	v164 = int32(0)
	if v23 == v164 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if v202 != 0 {
		v205 = v142
		goto L55
	} else {
		goto L76
	}
L64:
	;
	v202 = int32(0)
	goto L63
L65:
	;
	goto L66
L66:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v170 <= int32(0) {
		v196 = v164
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v202 = v196
	goto L63
L68:
	;
	v173 = int32(0)
	if v173 < v170 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v176 = v170
	goto L71
L70:
	;
	v176 = v173
	goto L71
L71:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v179 = int32(0)
	goto L72
L72:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v177+v179<<(uint(int32(2))%32))))
	v188 = base.B2i32(v187 == v149)
	if v187 == v149 {
		v196 = v188
		goto L67
	} else {
		goto L74
	}
L73:
	;
	v196 = v188
	goto L67
L74:
	;
	v190 = v179 + int32(1)
	if v190 != v176 {
		v179 = v190
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v203 = F_lappend_oid(m, v142, v149)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v205 = v203
	goto L55
L78:
	;
	if v206 != 0 {
		v138 = v206
		v142 = v205
		goto L53
	} else {
		goto L79
	}
L79:
	;
	goto L54
L80:
	;
	v227 = v213
	goto L46
L81:
	;
	m.G0 = v12 - int32(-64)
	return v227
}
func F_ResetAllOptions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	v1 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0]))
	if base.B2i32(v9 == v1)|base.B2i32(v9 == int32(_a_F_ResetAllOptions_0)) == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = v9
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(-64))))
	if base.Ui32(int32(1)) < base.Ui32(v27-int32(5)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if v24 != int32(_a_F_ResetAllOptions_0) {
		v21 = v24
		goto L4
	} else {
		goto L94
	}
L7:
	;
	v33 = v21 - int32(48)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v34&int32(16) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v38 = v21 - int32(36)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if base.Ui32(v39) < base.Ui32(int32(11)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_push_old_value(m, v21-int32(68), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(44))))
	switch v49 {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	default:
		goto L12
	}
L12:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(32))))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v287 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L13:
	;
	F_pfree(m, v270)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L10
	} else {
		goto L81
	}
L14:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	if v231 != 0 {
		goto L71
	} else {
		goto L72
	}
L15:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v161 != 0 {
		goto L49
	} else {
		goto L50
	}
L16:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	if v124 != 0 {
		goto L39
	} else {
		goto L40
	}
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v87 != 0 {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
	if v50 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+48)))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	m.T0[v50].(func(*base.Module, int32, int32))(m, v51, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+48)))
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v58)
	v61 = v21 - int32(8)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v65
	if base.B2i32(v62 == int32(0))|base.B2i32(v65 == v62) != 0 {
		goto L12
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v73 = v21 - int32(12)
	goto L24
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v80 == int32(0) {
		v270 = v62
		goto L13
	} else {
		goto L26
	}
L25:
	;
	goto L12
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
	if v62 == v83 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
	if v62 != v85 {
		v73 = v80
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	m.T0[v87].(func(*base.Module, int32, int32))(m, v88, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v95
	v98 = v21 - int32(8)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v102
	if base.B2i32(v99 == int32(0))|base.B2i32(v102 == v99) != 0 {
		goto L12
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v110 = v21 - int32(12)
	goto L34
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v117 == int32(0) {
		v270 = v99
		goto L13
	} else {
		goto L36
	}
L35:
	;
	goto L12
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v117)+40))
	if v99 == v120 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)+56))
	if v99 != v122 {
		v110 = v117
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v21)+76))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	m.T0[v124].(func(*base.Module, float64, int32))(m, v125, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v132 = *(*float64)(unsafe.Add(mBase, uint32(v21)+76))
	*(*float64)(unsafe.Add(mBase, uint32(v131))) = v132
	v135 = v21 - int32(8)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v139
	if base.B2i32(v136 == int32(0))|base.B2i32(v139 == v136) != 0 {
		goto L12
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v147 = v21 - int32(12)
	goto L44
L44:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v154 == int32(0) {
		v270 = v136
		goto L13
	} else {
		goto L46
	}
L45:
	;
	goto L12
L46:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154)+40))
	if v136 == v157 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v154)+56))
	if v136 != v159 {
		v147 = v154
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	m.T0[v161].(func(*base.Module, int32, int32))(m, v162, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v170
	if v169 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v205 = v21 - int32(8)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v209
	if base.B2i32(v206 == int32(0))|base.B2i32(v209 == v206) != 0 {
		goto L12
	} else {
		goto L65
	}
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if base.B2i32(v169 == v175)|base.B2i32(v170 == v169) != 0 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	if v169 == v179 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v185 = v21 - int32(12)
	goto L57
L57:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v190 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	F_pfree(m, v169)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L64
	}
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+32))
	if v169 == v191 {
		goto L53
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v190)+48))
	if v169 != v193 {
		v185 = v190
		goto L57
	} else {
		goto L63
	}
L63:
	;
	goto L53
L64:
	;
	goto L53
L65:
	;
	v217 = v21 - int32(12)
	goto L66
L66:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	if v224 == int32(0) {
		v270 = v206
		goto L13
	} else {
		goto L68
	}
L67:
	;
	goto L12
L68:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+40))
	if v206 == v227 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+56))
	if v206 != v229 {
		v217 = v224
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	m.T0[v231].(func(*base.Module, int32, int32))(m, v232, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v239
	v242 = v21 - int32(8)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = v246
	if base.B2i32(v243 == int32(0))|base.B2i32(v246 == v243) != 0 {
		goto L12
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	v254 = v21 - int32(12)
	goto L76
L76:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v261 == int32(0) {
		v270 = v243
		goto L13
	} else {
		goto L78
	}
L77:
	;
	goto L12
L78:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+40))
	if v243 == v264 {
		goto L12
	} else {
		goto L79
	}
L79:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261)+56))
	if v243 != v266 {
		v254 = v261
		goto L76
	} else {
		goto L80
	}
L80:
	;
	goto L77
L81:
	;
	goto L12
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v286
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v21-int32(28)))) = v318
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v21-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v21-int32(20)))) = v324
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v326&int32(64) == int32(0) {
		goto L6
	} else {
		goto L92
	}
L83:
	;
	if v286 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	if v286 != 0 {
		goto L82
	} else {
		goto L91
	}
L86:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0]))
	if v293 != 0 {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v300
	v302 = int32(_a_F_ResetAllOptions_0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v300)+4)) = v21
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[1])) = v21
	goto L82
L88:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[1]))
	v300 = v295
	goto L87
L89:
	;
	goto L90
L90:
	;
	v297 = int32(_a_F_ResetAllOptions_0)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[0])) = v297
	v300 = v297
	goto L87
L91:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+4)) = v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v308))) = v310
	goto L82
L92:
	;
	v332 = v21 - int32(40)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v333&int32(4) != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v333 | int32(4)
	v339 = int32(_a_F_ResetAllOptions_1)
	v340 = *(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v340
	*(*int32)(unsafe.Add(mBase, _c_F_ResetAllOptions[2])) = v21 + int32(12)
	goto L6
L94:
	;
	goto L5
}
