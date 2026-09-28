package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FetchRelationStates(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	v4 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v4)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[0]))
	if v11 == int32(2) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 != 0 {
		goto L41
	} else {
		goto L42
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[0])) = int32(1)
	v18 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_FetchRelationStates[1])) = uint8(v18)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[2]))
	F_list_free_deep(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[2])) = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[3]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	goto L5
L5:
	;
	if base.B2i32(v29 == int32(2)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[4]))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v42 = int32(1)
	v45 = F_GetSubscriptionRelations(m, v41, v42, v42, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v36 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v36)
	goto L8
L10:
	;
	v47 = int32(_a_F_FetchRelationStates_0)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[5]))
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[5])) = v51
	if v45 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[5])) = v48
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[2]))
	if v109 != 0 {
		goto L24
	} else {
		goto L25
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v55 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v60 = int32(0)
	goto L14
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v60<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = F_get_rel_relkind(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L11
L16:
	;
	v95 = v60 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v95 < v96 {
		v60 = v95
		goto L14
	} else {
		goto L23
	}
L17:
	;
	if v71 == int32(83) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_FetchRelationStates[1])) = uint8(v76)
	goto L16
L19:
	;
	goto L20
L20:
	;
	v79 = F_palloc(m, int32(24))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v69)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+16)) = v81
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v83
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	*(*int64)(unsafe.Add(mBase, uint32(v79))) = v85
	v87 = int32(_a_F_FetchRelationStates_1)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[2]))
	v90 = F_lappend(m, v89, v79)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[2])) = v90
	goto L16
L23:
	;
	goto L15
L24:
	;
	v171 = int32(1)
	goto L26
L25:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[4]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v114 = m.G0
	v116 = v114 + int32(-64)
	m.G0 = v116
	v120 = F_table_open(m, int32(_a_F_FetchRelationStates_2), int32(1))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L27
	}
L26:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_FetchRelationStates[7])) = uint8(v171)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[0]))
	if v174 != int32(1) {
		goto L1
	} else {
		goto L40
	}
L27:
	;
	F_ScanKeyInit(m, v116, int32(1), int32(3), int32(184), base.I64_extend_i32_u(v113))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v128 = int32(0)
	v132 = F_systable_beginscan(m, v120, v128, v128, v128, int32(1), v116)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	goto L30
L30:
	;
	v141 = F_systable_getnext(m, v132)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L32
	}
L31:
	;
	F_systable_endscan(m, v132)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L38
	}
L32:
	;
	if v141 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+22)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143+v144)+4))
	v147 = F_get_rel_relkind(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L31
L36:
	;
	if v147&int32(-3) != int32(112) {
		goto L30
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	F_relation_close(m, v120, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v116 - int32(-64)
	v171 = base.B2i32(v141 != int32(0))
	goto L26
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FetchRelationStates[0])) = int32(2)
	goto L1
L41:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FetchRelationStates[7])))
	*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v188)
	goto L43
L42:
	;
	goto L43
L43:
	;
	if l1 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FetchRelationStates[1])))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v191)
	goto L46
L45:
	;
	goto L46
L46:
	;
	return
}
func F_RelationCacheInitializePhase3(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[2]))
	F_read_relmap_file(m, int32(_a_F_RelationCacheInitializePhase3_0), v17, int32(0), int32(22))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v22 = int32(_a_F_RelationCacheInitializePhase3_1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3]))
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3])) = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v29 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[3])) = v23
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v72 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L7:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_2), int32(83), int32(0), int32(34), int32(_a_F_RelationCacheInitializePhase3_3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L8:
	;
	v33 = F_load_relcache_init_file(m, int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v33 == int32(0) {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v68 = v12 ^ int32(1)
	goto L6
L11:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_4), int32(75), int32(0), int32(25), int32(_a_F_RelationCacheInitializePhase3_5))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_6), int32(81), int32(0), int32(30), int32(_a_F_RelationCacheInitializePhase3_7))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_formrdesc(m, int32(_a_F_RelationCacheInitializePhase3_8), int32(71), int32(0), int32(32), int32(_a_F_RelationCacheInitializePhase3_9))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v68 = int32(1)
	goto L6
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L106
	}
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L102
	}
L17:
	;
	m.G0 = v9 + int32(48)
	return
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[5])))
	if v76 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_load_critical_index(m, int32(2662), int32(1259))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])))
	if v111 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L22:
	;
	F_load_critical_index(m, int32(2659), int32(1249))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_load_critical_index(m, int32(2679), int32(2610))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	F_load_critical_index(m, int32(2687), int32(2616))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_load_critical_index(m, int32(2655), int32(2603))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_load_critical_index(m, int32(2693), int32(2618))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	F_load_critical_index(m, int32(2701), int32(2620))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[5])) = uint8(v108)
	goto L21
L29:
	;
	F_load_critical_index(m, int32(2671), int32(1262))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v150 = v9 + int32(28)
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[6]))
	F_hash_seq_init(m, v150, v152)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L40
	}
L32:
	;
	F_load_critical_index(m, int32(2672), int32(1262))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_load_critical_index(m, int32(2676), int32(1260))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_load_critical_index(m, int32(2677), int32(1260))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_load_critical_index(m, int32(2695), int32(1261))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_load_critical_index(m, int32(3593), int32(3592))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	F_load_critical_index(m, int32(_a_F_RelationCacheInitializePhase3_10), int32(_a_F_RelationCacheInitializePhase3_11))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	F_load_critical_index(m, int32(_a_F_RelationCacheInitializePhase3_12), int32(_a_F_RelationCacheInitializePhase3_11))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[0])) = uint8(v147)
	goto L31
L40:
	;
	v155 = F_hash_seq_search(m, v150)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if v155 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v157 = v155
	goto L45
L43:
	;
	goto L44
L44:
	;
	if v68&int32(1) == int32(0) {
		goto L17
	} else {
		goto L95
	}
L45:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerEnlarge(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v168 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v173 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerRemember(m, v175, base.I64_extend_i32_u(v163), int32(_a_F_RelationCacheInitializePhase3_13))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+80))
	if v181 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L50
L52:
	;
	v185 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v163)+56)))
	v186 = F_SearchSysCache1(m, int32(57), v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	v207 = v180
	goto L54
L54:
	;
	v210 = base.B2i32(v181 == int32(0))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+124)))
	if v211 != int32(1) {
		v223 = v207
		v224 = v210
		goto L64
	} else {
		goto L65
	}
L55:
	;
	if v186 == int32(0) {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+22)))
	base.MemoryCopy(m, v190, v191+v192, int32(144))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v163)+180))
	if v196 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	F_pfree(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	F_RelationParseRelOptions(m, v163, v186)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	F_ReleaseCatCache(m, v186)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+80))
	if v204 == int32(0) {
		goto L15
	} else {
		goto L63
	}
L63:
	;
	v207 = v203
	goto L54
L64:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+125)))
	if v225 != int32(1) {
		v237 = v223
		v238 = v224
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v163)+68))
	if v214 != 0 {
		v223 = v207
		v224 = v210
		goto L64
	} else {
		goto L66
	}
L66:
	;
	F_RelationBuildRuleLock(m, v163)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v218 = int32(1)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v163)+68))
	if v219 != 0 {
		v223 = v217
		v224 = v218
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v220 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v217)+124)) = uint8(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v223 = v222
	v224 = v218
	goto L64
L69:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+127)))
	if v239 != int32(1) {
		v246 = v238
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v163)+76))
	if v228 != 0 {
		v237 = v223
		v238 = v224
		goto L69
	} else {
		goto L71
	}
L71:
	;
	F_RelationBuildTriggers(m, v163)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v232 = int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v163)+76))
	if v233 != 0 {
		v237 = v231
		v238 = v232
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+125)) = uint8(v234)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v237 = v236
	v238 = v232
	goto L69
L74:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v163)+188))
	if v247 != 0 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v163)+80))
	if v242 != 0 {
		v246 = v238
		goto L74
	} else {
		goto L76
	}
L76:
	;
	F_RelationBuildRowSecurity(m, v163)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	v246 = int32(1)
	goto L74
L78:
	;
	v293 = F_hash_seq_search(m, v9+int32(28))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L93
	}
L79:
	;
	v283 = v9 + int32(28)
	F_hash_seq_term(m, v283)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L91
	}
L80:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v268 - int32(1)
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v273 != 0 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+119)))
	switch v249 - int32(83) {
	case 0, 26, 31, 33:
		goto L82
	default:
		goto L80
	}
L82:
	;
	F_RelationInitTableAccessMethod(m, v163)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+16)) = v254 - int32(1)
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[1]))
	if v259 == int32(0) {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerForget(m, v263, base.I64_extend_i32_u(v163), int32(_a_F_RelationCacheInitializePhase3_13))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	goto L79
L86:
	;
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[7]))
	F_ResourceOwnerForget(m, v275, base.I64_extend_i32_u(v163), int32(_a_F_RelationCacheInitializePhase3_13))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v246 == int32(0) {
		goto L78
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	goto L79
L91:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitializePhase3[6]))
	F_hash_seq_init(m, v283, v287)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	goto L78
L93:
	;
	if v293 != 0 {
		v157 = v293
		goto L45
	} else {
		goto L94
	}
L94:
	;
	goto L46
L95:
	;
	v306 = int32(0)
	goto L96
L96:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v306<<(uint(int32(2))%32))+uint32(_c_F_RelationCacheInitializePhase3[8])))
	F_InitCatCachePhase2(m, v314, int32(1))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L4
	} else {
		goto L98
	}
L97:
	;
	F_write_relcache_init_file(m, int32(1))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L100
	}
L98:
	;
	v319 = v306 + int32(1)
	if v319 != int32(85) {
		v306 = v319
		goto L96
	} else {
		goto L99
	}
L99:
	;
	goto L97
L100:
	;
	F_write_relcache_init_file(m, int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	goto L17
L102:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v163)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v344
	F_errmsg_internal(m, int32(_a_F_RelationCacheInitializePhase3_14), v9)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_RelationCacheInitializePhase3_15), int32(_a_F_RelationCacheInitializePhase3_16), int32(_a_F_RelationCacheInitializePhase3_17))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v163)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v358 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationCacheInitializePhase3_18), v9+int32(16))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_RelationCacheInitializePhase3_15), int32(_a_F_RelationCacheInitializePhase3_19), int32(_a_F_RelationCacheInitializePhase3_17))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCopyStorage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[0]))
	v26 = (base.B2i32(l2 == int32(3))&base.B2i32(l3 == int32(117)) | base.B2i32(l3 == int32(112))) & base.B2i32(v5 < v23)
	v28 = F_palloc(m, int32(424))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)) = uint8(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l1
	v35 = F_smgrnblocks(m, l1, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+400)) = v35
	v38 = F_GetRedoRecPtr(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v28)+408)) = v38
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+416)) = v42
	v44 = F_smgrnblocks(m, l0, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v108 = v12 + int32(20)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v108, v109, v110, v111, v112, l2)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L6:
	;
	if v44 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v51 = v5
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_smgr_bulk_finish(m, v28)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorage[2]))
	if v56 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v59 = F_smgr_bulk_get_buf(m, v28)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v59
	F_smgrreadv(m, l0, l2, v51, v12+int32(20))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationCopyStorage[3])))
	if v69 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = int32(5)
	goto L20
L19:
	;
	v70 = int32(1)
	goto L20
L20:
	;
	v73 = F_PageIsVerified(m, v59, v51, v70, v12+int32(95))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+95)))
	if v75 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_pgstat_prepare_report_checksum_failure(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if v73 == int32(0) {
		goto L5
	} else {
		goto L27
	}
L25:
	;
	F_pgstat_report_checksum_failures_in_db(m, v78, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	F_smgr_bulk_write(m, v28, v51, v59, int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v91 = v51 + int32(1)
	if v91 != v44 {
		v51 = v91
		goto L10
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	m.G0 = v12 + int32(96)
	return
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v108
	F_errmsg(m, int32(_a_F_RelationCopyStorage_0), v12)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_RelationCopyStorage_1), int32(550), int32(_a_F_RelationCopyStorage_2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationCopyStorageUsingBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v104 int32
	_ = v104
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
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	v5 = int32(0)
	v16 = m.G0
	v20 = (v16 - int32(_a_F_RelationCopyStorageUsingBuffer_0)) & int32(-4096)
	m.G0 = v20
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[0]))
	if v5 < v27 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = l3 | base.B2i32(l2 == int32(3))
	goto L3
L2:
	;
	v30 = v5
	goto L3
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4080)) = v31
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4072)) = v33
	v38 = F_smgropen(m, v20+int32(4072), int32(-1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v40 = F_smgrnblocks(m, v38, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v20 + int32(_a_F_RelationCopyStorageUsingBuffer_1)
	base.MemoryFill(m, v43, int32(0), int32(_a_F_RelationCopyStorageUsingBuffer_2))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4064)) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4056)) = v49
	v54 = F_smgropen(m, v20+int32(4056), int32(-1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	m.G0 = v16
	return
L10:
	;
	v56 = int32(1)
	F_smgrextend(m, v54, l2, v40-v56, v43, v56)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v62 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v65 = F_GetAccessStrategy(m, int32(2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4092)) = v40
	v68 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4088)) = v68
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4040)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4048)) = v72
	v79 = F_smgropen(m, v20+int32(4040), int32(-1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	if l3 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v83 = int32(112)
	goto L17
L16:
	;
	v83 = int32(117)
	goto L17
L17:
	;
	v88 = F_read_stream_begin_impl(m, int32(12), v62, v68, v79, v83, l2, int32(3), v20+int32(4088), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v104 = v5
	goto L19
L19:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[1]))
	if v106 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	F_read_stream_end(m, v88)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L44
	}
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v110 = F_read_stream_next_buffer(m, v88, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4032)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+4024)) = v155
	v160 = F_ReadBufferWithoutRelcache(m, v20+int32(4024), l2, v152, int32(1), v65, l3)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L32
	}
L26:
	;
	if v110 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[2]))
	v117 = v110 ^ int32(-1)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[3]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v117<<(uint(int32(2))%32))))
	v149 = v115 + v117*int32(56)
	v151 = v126
	goto L25
L28:
	;
	goto L29
L29:
	;
	v127 = int32(56)
	v128 = v110 * v127
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	F_BufferLockAcquire(m, v110, v128+v130-v127, int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[4]))
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[5]))
	v149 = v138 + v128 - int32(56)
	v151 = v143 + v110<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L31:
	;
	v180 = int32(_a_F_RelationCopyStorageUsingBuffer_3)
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6])) = v182 + int32(1)
	base.MemoryCopy(m, v179, v151, int32(_a_F_RelationCopyStorageUsingBuffer_2))
	F_MarkBufferDirty(m, v160)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L4
	} else {
		goto L36
	}
L32:
	;
	if v160 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[3]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165+(v160^int32(-1))<<(uint(int32(2))%32))))
	v179 = v171
	goto L31
L34:
	;
	goto L35
L35:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[5]))
	v179 = v173 + v160<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L36:
	;
	if v30 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_log_newpage_buffer(m, v160, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v193 = int32(_a_F_RelationCopyStorageUsingBuffer_3)
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationCopyStorageUsingBuffer[6])) = v195 - int32(1)
	F_UnlockReleaseBuffer(m, v160)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	F_UnlockReleaseBuffer(m, v110)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v204 = v104 + int32(1)
	if v204 != v40 {
		v104 = v204
		goto L19
	} else {
		goto L43
	}
L43:
	;
	goto L20
L44:
	;
	F_bms_free(m, v62)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	F_bms_free(m, v65)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	goto L9
}
func F_RelationDropStorage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[0]))
	v7 = F_MemoryContextAlloc(m, v5, int32(28))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v9
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v14)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v13
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[1]))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v19
		v21 = int32(_a_F_RelationDropStorage_0)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v22
		*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[2])) = v7
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v26 != 0 {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+72))
			v31 = v29 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v31
			if v31 == int32(0) {
				v36 = v26 + int32(76)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[3]))
				if v38 != 0 {
					v40 = *(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[4]))
					v45 = v40
				} else {
					v42 = int32(_a_F_RelationDropStorage_1)
					*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[3])) = v42
					v45 = v42
				}
				*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v45
				v47 = int32(_a_F_RelationDropStorage_1)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v47
				*(*int32)(unsafe.Add(mBase, uint32(v45)+4)) = v36
				*(*int32)(unsafe.Add(mBase, _c_F_RelationDropStorage[4])) = v36
			} else {
			}
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_smgrclose(m, v54)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
				return
			}
		} else {
			return
		}
	}
}
func F_RelationGetExclusionInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int64
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	v15 = m.G0
	v17 = v15 - int32(128)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(v20)+10)))
	v22 = F_palloc_mul(m, int32(4), v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
	v26 = F_palloc_mul(m, int32(4), v21)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26
	v30 = F_palloc_mul(m, int32(2), v21)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v30
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	if v33 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L85
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L82
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L79
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L76
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L10:
	;
	F_systable_endscan(m, v69)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L71
	}
L11:
	;
	m.G0 = v17 + int32(128)
	return
L12:
	;
	v35 = v21 << (uint(int32(2)) % 32)
	v36 = int32(0)
	v37 = base.B2i32(v35 == v36)
	if v37 == v36 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v53 = v17 - int32(-64)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v58 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v57)+4)))
	F_ScanKeyInit(m, v53, int32(9), int32(3), int32(184), v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L22
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	base.MemoryCopy(m, v22, v40, v35)
	goto L17
L16:
	;
	goto L17
L17:
	;
	if v37 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v26, v44, v35)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v47 = v21 << (uint(int32(1)) % 32)
	if v47 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	base.MemoryCopy(m, v30, v50, v47)
	goto L11
L22:
	;
	v63 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v66 = int32(1)
	v69 = F_systable_beginscan(m, v63, int32(2665), v66, int32(0), v66, v53)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v71 = F_systable_getnext(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v71 == int32(0) {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	v76 = v21 << (uint(int32(2)) % 32)
	v78 = v71
	v88 = int32(0)
	goto L27
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
	v93 = v91 + v92
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+72)))
	if v94 == int32(120) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	F_systable_endscan(m, v69)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L47
	}
L29:
	;
	v132 = F_systable_getnext(m, v69)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L45
	}
L30:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v93)+88))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v102 != v103 {
		v131 = v88
		goto L29
	} else {
		goto L33
	}
L31:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+107)))
	if v97 != int32(1) {
		v131 = v88
		goto L29
	} else {
		goto L32
	}
L32:
	;
	switch v94 - int32(112) {
	case 0, 5:
		goto L30
	default:
		v131 = v88
		goto L29
	}
L33:
	;
	if v88 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v63)+52))
	v109 = F_fastgetattr_3(m, v78, int32(27), v106, v17+int32(63))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+63)))
	if v111 == int32(1) {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	v115 = F_pg_detoast_datum(m, base.I32_wrap_i64(v109))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v117 != int32(1) {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	if v120 != v21 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	if v122 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	if v123 != int32(26) {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	if v76 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	base.MemoryCopy(m, v22, v115+int32(24), v76)
	goto L44
L43:
	;
	goto L44
L44:
	;
	v131 = int32(1)
	goto L29
L45:
	;
	if v132 != 0 {
		v78 = v132
		v88 = v131
		goto L27
	} else {
		goto L46
	}
L46:
	;
	goto L28
L47:
	;
	F_relation_close(m, v63, int32(1))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v131 == int32(0) {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	if int32(0) < v21 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v197 = int32(_a_F_RelationGetExclusionInfo_0)
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0])) = v200
	v203 = F_palloc_mul(m, int32(4), v21)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L59
	}
L53:
	;
	v159 = v145 << (uint(int32(2)) % 32)
	v161 = v159 + v22
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v163 = F_get_opcode(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26+v159))) = v163
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v170+v159)))
	v173 = F_get_op_opfamily_strategy(m, v169, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v145<<(uint(int32(1))%32)))) = uint16(v173)
	if v173&int32(_a_F_RelationGetExclusionInfo_1) == int32(0) {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v181 = v145 + int32(1)
	if v181 != v21 {
		v145 = v181
		goto L53
	} else {
		goto L58
	}
L58:
	;
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+236)) = v203
	v207 = F_palloc_mul(m, int32(4), v21)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+240)) = v207
	v211 = F_palloc_mul(m, int32(2), v21)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v211
	v214 = int32(0)
	v215 = base.B2i32(v76 == v214)
	if v215 == v214 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	base.MemoryCopy(m, v218, v22, v76)
	goto L64
L63:
	;
	goto L64
L64:
	;
	if v215 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	base.MemoryCopy(m, v222, v26, v76)
	goto L67
L66:
	;
	goto L67
L67:
	;
	v225 = v21 << (uint(int32(1)) % 32)
	if v225 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	base.MemoryCopy(m, v226, v30, v225)
	goto L70
L69:
	;
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetExclusionInfo[0])) = v198
	goto L11
L71:
	;
	F_relation_close(m, v63, int32(1))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L9
L73:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v270 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_2), v17)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_4), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v286 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_6), v17+int32(16))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_7), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
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
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v304 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_8), v17+int32(32))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_9), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_10), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_11), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v336+v145<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v335
	F_errmsg_internal(m, int32(_a_F_RelationGetExclusionInfo_12), v17+int32(48))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_RelationGetExclusionInfo_3), int32(_a_F_RelationGetExclusionInfo_13), int32(_a_F_RelationGetExclusionInfo_5))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetReplicaIndex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v2 == int32(0) {
		v5 = F_RelationGetIndexList(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			F_list_free(m, v5)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
				return v11
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
		return v11
	}
}
func F_RelationInitPhysicalAddr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
	switch v11 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L4
	default:
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L37
	} else {
		goto L85
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L37
	} else {
		goto L82
	}
L3:
	;
	m.G0 = v8 + int32(48)
	return
L4:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v10)+92))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[0]))
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v18 = v15
	goto L7
L6:
	;
	v18 = v17
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[1]))
	if v18 != int32(1664) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v25 = v21
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
	if v27 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[2]))
	if base.B2i32(v207 == v14)|base.B2i32(v211 < int32(0)) != 0 {
		goto L3
	} else {
		goto L80
	}
L12:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[3]))
	goto L16
L13:
	;
	goto L14
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+117)))
	v101 = int32(0)
	if v100 == v101 {
		goto L44
	} else {
		goto L45
	}
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+88))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v97
	v207 = v97
	goto L11
L16:
	;
	if base.B2i32(v29 != int32(0)) == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[4]))
	if v35 <= int32(1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[5])))
	if v39&int32(1) == int32(0) {
		goto L15
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	if v45 != int32(112) {
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	if v35 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v50 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L28
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v51 != 0 {
		goto L15
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	if base.B2i32(base.Ui32(v52) < base.Ui32(int32(_a_F_RelationInitPhysicalAddr_0))) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v57 == int32(0) {
		goto L15
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[6]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	goto L35
L32:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+119)))
	switch v61 - int32(109) {
	case 0, 5:
		goto L33
	default:
		goto L15
	}
L33:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+112)))
	if v64 != int32(1) {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	if base.B2i32(v70 == int32(2)) == int32(0) {
		goto L15
	} else {
		goto L36
	}
L36:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v79 = F_ScanPgRelation(m, v75, base.B2i32(v75 != int32(2662)), int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	return
L38:
	;
	if v79 == int32(0) {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	v86 = v84 + v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+92)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+88)) = v90
	F_pfree(m, v79)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L15
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v203
	if v203 == int32(0) {
		goto L1
	} else {
		goto L79
	}
L42:
	;
	goto L41
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v203 = v198
	goto L42
L44:
	;
	v106 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[7]))
	if v106 < v108 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v149 = int32(0)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[8]))
	if v149 < v151 {
		goto L65
	} else {
		goto L66
	}
L47:
	;
	v112 = v106
	goto L50
L48:
	;
	goto L49
L49:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[9]))
	if v131 <= int32(0) {
		v203 = v101
		goto L42
	} else {
		goto L56
	}
L50:
	;
	v117 = v112 << (uint(int32(3)) % 32)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_RelationInitPhysicalAddr[10])))
	if v118 == v99 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v197 = v117 + int32(_a_F_RelationInitPhysicalAddr_1)
	goto L43
L53:
	;
	goto L54
L54:
	;
	v123 = v112 + int32(1)
	if v123 != v108 {
		v112 = v123
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v136 = int32(0)
	goto L57
L57:
	;
	v141 = v136 << (uint(int32(3)) % 32)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+uint32(_c_F_RelationInitPhysicalAddr[11])))
	if v142 != v99 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v197 = v141 + int32(_a_F_RelationInitPhysicalAddr_2)
	goto L43
L59:
	;
	v145 = v136 + int32(1)
	if v131 != v145 {
		v136 = v145
		goto L57
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
	v203 = v101
	goto L42
L63:
	;
	v179 = int32(0)
	goto L73
L64:
	;
	v197 = v160 + int32(_a_F_RelationInitPhysicalAddr_3)
	goto L43
L65:
	;
	v155 = v149
	goto L68
L66:
	;
	goto L67
L67:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitPhysicalAddr[12]))
	if v172 <= int32(0) {
		v203 = v101
		goto L42
	} else {
		goto L72
	}
L68:
	;
	v160 = v155 << (uint(int32(3)) % 32)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+uint32(_c_F_RelationInitPhysicalAddr[13])))
	if v99 == v161 {
		goto L64
	} else {
		goto L70
	}
L69:
	;
	goto L67
L70:
	;
	v164 = v155 + int32(1)
	if v164 != v151 {
		v155 = v164
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L63
L73:
	;
	v184 = v179 << (uint(int32(3)) % 32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+uint32(_c_F_RelationInitPhysicalAddr[14])))
	if v185 != v99 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v197 = v184 + int32(_a_F_RelationInitPhysicalAddr_4)
	goto L43
L75:
	;
	v188 = v179 + int32(1)
	if v172 != v188 {
		v179 = v188
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	v203 = v101
	goto L42
L79:
	;
	v207 = v203
	goto L11
L80:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v215
	v217 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v217
	v221 = F_RelFileLocatorSkippingWAL(m, v8+int32(16))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L37
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v221
	goto L3
L82:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v234
	F_errmsg_internal(m, int32(_a_F_RelationInitPhysicalAddr_5), v8+int32(32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_RelationInitPhysicalAddr_6), int32(1380), int32(_a_F_RelationInitPhysicalAddr_7))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L37
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v250 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationInitPhysicalAddr_8), v8)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L37
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_RelationInitPhysicalAddr_6), int32(1398), int32(_a_F_RelationInitPhysicalAddr_7))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L37
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationTruncate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int64
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(144)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v15 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v19
	v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v21
	v25 = F_smgropen(m, v13+int32(72), v18)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v42 = v15
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = int32(-1)
	v45 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = v45
	*(*int64)(unsafe.Add(mBase, uint32(v42)+16)) = v45
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v49
	v52 = F_smgrnblocks(m, v42, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L4
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v25
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v42 = v41
	goto L3
L7:
	;
	v37 = v29
	goto L9
L8:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v25)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	v37 = v35
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = v37 + int32(1)
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v52
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v56 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v62
	v66 = F_smgropen(m, v13+int32(56), v59)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v83 = v56
	goto L13
L13:
	;
	v86 = v13 + int32(136)
	v88 = v13 + int32(124)
	v90 = v13 + int32(112)
	v91 = int32(1)
	v93 = F_smgrexists(m, v83, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L20
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	if v70 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v83 = v82
	goto L13
L16:
	;
	v78 = v70
	goto L18
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v66)+76))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v72
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+72))
	v78 = v76
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+72)) = v78 + int32(1)
	goto L15
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v122 != 0 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	if v93 == int32(0) {
		v117 = v91
		v118 = v86
		v119 = v88
		v120 = v90
		v121 = v3
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v97 = F_FreeSpaceMapPrepareTruncateRel(m, l0, l1)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v97
	if v97 == int32(-1) {
		v117 = v91
		v118 = v86
		v119 = v88
		v120 = v90
		v121 = v3
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v108 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v108
	v112 = F_smgrnblocks(m, v42, v108)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+124)) = v112
	v117 = int32(2)
	v118 = v13 + int32(140)
	v119 = v13 + int32(128)
	v120 = v13 + int32(116)
	v121 = v108
	goto L19
L25:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[0]))
	if v170 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v148 = v122
	goto L28
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v124
	v126 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v126
	v130 = F_smgropen(m, v13+int32(40), v123)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L29
	}
L28:
	;
	v150 = F_smgrexists(m, v148, int32(2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L34
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v130
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
	if v134 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = v146
	goto L28
L31:
	;
	v142 = v134
	goto L33
L32:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+76))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v130)+72))
	v142 = v140
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130)+72)) = v142 + int32(1)
	goto L30
L34:
	;
	if v150 == int32(0) {
		v168 = v117
		goto L25
	} else {
		goto L35
	}
L35:
	;
	v154 = F_visibilitymap_prepare_truncate(m, l0, l1)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v154
	if v154 == int32(-1) {
		v168 = v117
		goto L25
	} else {
		goto L37
	}
L37:
	;
	v159 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v159
	v162 = F_smgrnblocks(m, v42, v159)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = v162
	v168 = v117 + int32(1)
	goto L25
L39:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[1]))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v210)+336)) = v211 | int32(3)
	v215 = int32(_a_F_RelationTruncate_0)
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2])) = v217 + int32(1)
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+118)))
	if v222 != int32(112) {
		goto L51
	} else {
		goto L52
	}
L40:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v173 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = v173
	goto L43
L42:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v175
	v177 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v177
	v181 = F_smgropen(m, v13+int32(24), v174)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L44
	}
L43:
	;
	v200 = int32(0)
	v202 = F_hash_search(m, v170, v199, v200, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L49
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+72))
	if v185 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v199 = v197
	goto L43
L46:
	;
	v193 = v185
	goto L48
L47:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+76))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v181)+72))
	v193 = v191
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181)+72)) = v193 + int32(1)
	goto L45
L49:
	;
	if v202 == int32(0) {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v202)+12)) = uint8(v206)
	goto L39
L51:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v251 != 0 {
		goto L62
	} else {
		goto L63
	}
L52:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[3]))
	if v226 <= int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v229 != 0 {
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = l1
	v232 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+92)) = v232
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v234
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = int32(7)
	F_XLogBeginInsert(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v230 != 0 {
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	F_XLogRegisterData(m, v13+int32(88), int32(20))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v247 = F_XLogInsert(m, int32(2), int32(33))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_XLogFlush(m, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	goto L51
L62:
	;
	v277 = v251
	goto L64
L63:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v253
	v255 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v255
	v259 = F_smgropen(m, v13+int32(8), v252)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L65
	}
L64:
	;
	F_smgrtruncate(m, v277, v13+int32(132), v168, v13+int32(120), v13+int32(108))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L4
	} else {
		goto L70
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v259
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v259)+72))
	if v263 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v277 = v275
	goto L64
L67:
	;
	v271 = v263
	goto L69
L68:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v259)+76))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v259)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v264)+4)) = v265
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v259)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v267
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v259)+72))
	v271 = v269
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+72)) = v271 + int32(1)
	goto L66
L70:
	;
	v286 = int32(_a_F_RelationTruncate_0)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[2])) = v288 - int32(1)
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncate[1]))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+336))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+336)) = v294 & int32(-4)
	if v121 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	F_FreeSpaceMapVacuumRange(m, l0, l1, int32(-1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	m.G0 = v13 + int32(144)
	return
L74:
	;
	goto L73
}
func F_SetRelationHasSubclass(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v17 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(l0), int64(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			if v17 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
				v21 = v19 + v20
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+126)))
				if v2 != v22 {
					*(*uint8)(unsafe.Add(mBase, uint32(v21)+126)) = uint8(v2)
					F_CatalogTupleUpdate(m, v12, v17+int32(4), v17)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_pfree(m, v17)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_relation_close(m, v12, int32(3))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				} else {
					F_CacheInvalidateRelcacheByTuple(m, v17)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_pfree(m, v17)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							F_relation_close(m, v12, int32(3))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(_a_F_SetRelationHasSubclass_0), v8)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetRelationHasSubclass_1), int32(3697), int32(_a_F_SetRelationHasSubclass_2))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
func F_UnlockRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v9
	v15 = F_LockRelease(m, v5, int32(8), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_UnlockRelationForExtension(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v10
	v15 = F_LockRelease(m, v6, l1, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		m.G0 = v6 + int32(16)
		return
	}
}
func F_UnlockRelationIdForSession(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	Fn14224(m, l0, l1, int32(1))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_UnlockRelationOid(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v53 int32
	_ = v53
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = int32(1)
	if l0 <= int32(3591) {
		if l0 <= int32(2670) {
			switch l0 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v82 = v11
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v82 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(l0-int32(2396)) {
					v82 = int32(0)
				} else {
					v82 = v11
				}
			}
		} else {
			v23 = l0 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v23))|base.B2i32(int32(1)<<(uint(v23)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(l0-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l0-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v82 = v11
				} else {
					v82 = int32(0)
				}
			} else {
				v82 = v11
			}
		}
	} else {
		if l0 <= int32(_a_F_UnlockRelationOid_0) {
			v36 = l0 - int32(_a_F_UnlockRelationOid_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v36))|base.B2i32(int32(1)<<(uint(v36)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(l0-int32(3592)) < base.Ui32(int32(2)) {
					v82 = v11
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(l0-int32(4060)) {
						v82 = int32(0)
					} else {
						v82 = v11
					}
				}
			} else {
				v82 = v11
			}
		} else {
			switch l0 - int32(_a_F_UnlockRelationOid_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v82 = v11
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v82 = int32(0)
			default:
				if base.Ui32(l0-int32(_a_F_UnlockRelationOid_3)) < base.Ui32(int32(3)) {
					v82 = v11
				} else {
					v53 = l0 - int32(_a_F_UnlockRelationOid_4)
					if base.Ui32(int32(15)) < base.Ui32(v53) {
						v82 = int32(0)
					} else {
						if int32(1)<<(uint(v53)%32)&int32(_a_F_UnlockRelationOid_5) != 0 {
							v82 = v11
						} else {
							v82 = int32(0)
						}
					}
				}
			}
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l0
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockRelationOid[0]))
	if v82 != 0 {
		v89 = int32(0)
	} else {
		v89 = v88
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v89
	v92 = F_LockRelease(m, v7, l1, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_get_relation_notnullatts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
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
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return
L2:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+16)))
	if v19 != int32(1) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+116))
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v41 = v23
	goto L6
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(34359738372)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_get_relation_notnullatts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v27
	v34 = F_hash_create(m, int32(_a_F_get_relation_notnullatts_0), int64(64), v9+int32(-56), int32(1064))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v47 = F_hash_search(m, v41, v9+int32(-4), int32(1), v9+int32(-56))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+116)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+116))
	v41 = v39
	goto L6
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+8)))
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if int32(0) < v51 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = int32(0)
	v59 = v50
	v60 = v51
	v61 = v3
	goto L14
L12:
	;
	v85 = v3
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v85
	goto L1
L14:
	;
	v64 = v55 + int32(1)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59+v55<<(uint(int32(3))%32))+35)))
	if v68 == int32(118) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v85 = v77
	goto L13
L16:
	;
	v71 = F_bms_add_member(m, v61, v64)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L19
	}
L17:
	;
	v75 = v59
	v76 = v60
	v77 = v61
	goto L18
L18:
	;
	if v64 < v76 {
		v55 = v64
		v59 = v75
		v60 = v76
		v61 = v77
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v75 = v73
	v76 = v74
	v77 = v71
	goto L18
L20:
	;
	goto L15
}
func F_get_relation_statistics_worker(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	v13 = F_SearchSysCache2(m, int32(62), base.I64_extend_i32_u(l2), base.I64_extend_i32_u(l3))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v19 = F_statext_is_kind_built(m, v13, int32(100))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					v22 = F_palloc0(m, int32(28))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(274)
						v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
						v28 = int32(100)
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+16)) = uint8(v28)
						*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = l1
						*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)) = uint8(v27)
						v32 = F_bms_copy(m, l4)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l5
							*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v32
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = F_lappend(m, v36, v22)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = v37
								v43 = F_statext_is_kind_built(m, v13, int32(102))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									if v43 != 0 {
										v46 = F_palloc0(m, int32(28))
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(274)
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
											v52 = int32(102)
											*(*uint8)(unsafe.Add(mBase, uint32(v46)+16)) = uint8(v52)
											*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = l1
											*(*uint8)(unsafe.Add(mBase, uint32(v46)+8)) = uint8(v51)
											v56 = F_bms_copy(m, l4)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = l5
												*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v56
												v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v61 = F_lappend(m, v60, v46)
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0))) = v61
													v67 = F_statext_is_kind_built(m, v13, int32(109))
													mBase = m.M
													v68 = m.ExcPending
													if v68 != 0 {
														return
													} else {
														if v67 != 0 {
															v70 = F_palloc0(m, int32(28))
															mBase = m.M
															v71 = m.ExcPending
															if v71 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(274)
																v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
																v76 = int32(109)
																*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v76)
																*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l1
																*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
																v80 = F_bms_copy(m, l4)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l5
																	*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v80
																	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																	v85 = F_lappend(m, v84, v70)
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
																		v91 = F_statext_is_kind_built(m, v13, int32(101))
																		mBase = m.M
																		v92 = m.ExcPending
																		if v92 != 0 {
																			return
																		} else {
																			if v91 != 0 {
																				v94 = F_palloc0(m, int32(28))
																				mBase = m.M
																				v95 = m.ExcPending
																				if v95 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
																					*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
																					v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
																					v100 = int32(101)
																					*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
																					*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
																					*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
																					v104 = F_bms_copy(m, l4)
																					mBase = m.M
																					v105 = m.ExcPending
																					if v105 != 0 {
																						return
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																						*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																						v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																						v109 = F_lappend(m, v108, v94)
																						mBase = m.M
																						v110 = m.ExcPending
																						if v110 != 0 {
																							return
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																							F_ReleaseCatCache(m, v13)
																							mBase = m.M
																							v116 = m.ExcPending
																							if v116 != 0 {
																								return
																							} else {
																								return
																							}
																						}
																					}
																				}
																			} else {
																				F_ReleaseCatCache(m, v13)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v91 = F_statext_is_kind_built(m, v13, int32(101))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																if v91 != 0 {
																	v94 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v95 = m.ExcPending
																	if v95 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
																		v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
																		v100 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
																		v104 = F_bms_copy(m, l4)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v109 = F_lappend(m, v108, v94)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																				F_ReleaseCatCache(m, v13)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v67 = F_statext_is_kind_built(m, v13, int32(109))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											if v67 != 0 {
												v70 = F_palloc0(m, int32(28))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(274)
													v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
													v76 = int32(109)
													*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v76)
													*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l1
													*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
													v80 = F_bms_copy(m, l4)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v80
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v85 = F_lappend(m, v84, v70)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
															v91 = F_statext_is_kind_built(m, v13, int32(101))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																if v91 != 0 {
																	v94 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v95 = m.ExcPending
																	if v95 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
																		v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
																		v100 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
																		v104 = F_bms_copy(m, l4)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v109 = F_lappend(m, v108, v94)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																				F_ReleaseCatCache(m, v13)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v91 = F_statext_is_kind_built(m, v13, int32(101))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v91 != 0 {
														v94 = F_palloc0(m, int32(28))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
															v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
															v100 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
															v104 = F_bms_copy(m, l4)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v109 = F_lappend(m, v108, v94)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v43 = F_statext_is_kind_built(m, v13, int32(102))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						if v43 != 0 {
							v46 = F_palloc0(m, int32(28))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(274)
								v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
								v52 = int32(102)
								*(*uint8)(unsafe.Add(mBase, uint32(v46)+16)) = uint8(v52)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = l1
								*(*uint8)(unsafe.Add(mBase, uint32(v46)+8)) = uint8(v51)
								v56 = F_bms_copy(m, l4)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = l5
									*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = v56
									v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v61 = F_lappend(m, v60, v46)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v61
										v67 = F_statext_is_kind_built(m, v13, int32(109))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											if v67 != 0 {
												v70 = F_palloc0(m, int32(28))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(274)
													v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
													v76 = int32(109)
													*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v76)
													*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l1
													*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
													v80 = F_bms_copy(m, l4)
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l5
														*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v80
														v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v85 = F_lappend(m, v84, v70)
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
															v91 = F_statext_is_kind_built(m, v13, int32(101))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																if v91 != 0 {
																	v94 = F_palloc0(m, int32(28))
																	mBase = m.M
																	v95 = m.ExcPending
																	if v95 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
																		v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
																		v100 = int32(101)
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
																		*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
																		*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
																		v104 = F_bms_copy(m, l4)
																		mBase = m.M
																		v105 = m.ExcPending
																		if v105 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																			*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																			v109 = F_lappend(m, v108, v94)
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																				F_ReleaseCatCache(m, v13)
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					return
																				}
																			}
																		}
																	}
																} else {
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												v91 = F_statext_is_kind_built(m, v13, int32(101))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v91 != 0 {
														v94 = F_palloc0(m, int32(28))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
															v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
															v100 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
															v104 = F_bms_copy(m, l4)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v109 = F_lappend(m, v108, v94)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							v67 = F_statext_is_kind_built(m, v13, int32(109))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return
							} else {
								if v67 != 0 {
									v70 = F_palloc0(m, int32(28))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v70)+4)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v70))) = int32(274)
										v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
										v76 = int32(109)
										*(*uint8)(unsafe.Add(mBase, uint32(v70)+16)) = uint8(v76)
										*(*int32)(unsafe.Add(mBase, uint32(v70)+12)) = l1
										*(*uint8)(unsafe.Add(mBase, uint32(v70)+8)) = uint8(v75)
										v80 = F_bms_copy(m, l4)
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = l5
											*(*int32)(unsafe.Add(mBase, uint32(v70)+20)) = v80
											v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v85 = F_lappend(m, v84, v70)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0))) = v85
												v91 = F_statext_is_kind_built(m, v13, int32(101))
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													if v91 != 0 {
														v94 = F_palloc0(m, int32(28))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
															v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
															v100 = int32(101)
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
															*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
															*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
															v104 = F_bms_copy(m, l4)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
																*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
																v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
																v109 = F_lappend(m, v108, v94)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
																	F_ReleaseCatCache(m, v13)
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		return
																	}
																}
															}
														}
													} else {
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										}
									}
								} else {
									v91 = F_statext_is_kind_built(m, v13, int32(101))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										if v91 != 0 {
											v94 = F_palloc0(m, int32(28))
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v94))) = int32(274)
												v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)))
												v100 = int32(101)
												*(*uint8)(unsafe.Add(mBase, uint32(v94)+16)) = uint8(v100)
												*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = l1
												*(*uint8)(unsafe.Add(mBase, uint32(v94)+8)) = uint8(v99)
												v104 = F_bms_copy(m, l4)
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = l5
													*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v104
													v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v109 = F_lappend(m, v108, v94)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
														F_ReleaseCatCache(m, v13)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return
														} else {
															return
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v13)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			return
		}
	}
}
func F_relation_mark_replica_identity(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v2 = l1
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v18 = F_table_open(m, int32(1259), int32(3))
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
	v21 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	v23 = F_SearchSysCacheCopy(m, int32(57), v21, int64(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L42
	}
L4:
	;
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+130)))
	if v28 != v2&int32(255) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L39
	}
L8:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+130)) = uint8(v2)
	F_CatalogTupleUpdate(m, v18, v23+int32(4), v23)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v18, int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	F_pfree(m, v23)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v44 = F_table_open(m, int32(2610), int32(3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v46 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	F_relation_close(m, v44, int32(3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L38
	}
L16:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v50 <= int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v55 = int32(0)
	goto L19
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66+v55<<(uint(int32(2))%32))))
	v73 = F_SearchSysCacheCopy(m, int32(34), base.I64_extend_i32_u(v70), int64(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	if v73 == int32(0) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+22)))
	v79 = v77 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)))
	if l2 == v70 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	F_pfree(m, v73)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L24:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+22)) = uint8(v92)
	F_CatalogTupleUpdate(m, v44, v73+int32(4), v73)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v82 = int32(1)
	if v80&v82 == int32(0) {
		v92 = v82
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v87 = int32(0)
	if v80&int32(1) == v87 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L23
L29:
	;
	v92 = v87
	goto L24
L30:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_relation_mark_replica_identity[0]))
	if v99 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2610), v70, v101, v101, int32(1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_CacheInvalidateRelcache(m, l0)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L23
L36:
	;
	v112 = v55 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v112 < v113 {
		v55 = v112
		goto L19
	} else {
		goto L37
	}
L37:
	;
	goto L20
L38:
	;
	m.G0 = v14 + int32(32)
	return
L39:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v136 + int32(4)
	F_errmsg_internal(m, int32(_a_F_relation_mark_replica_identity_0), v14)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_relation_mark_replica_identity_1), int32(_a_F_relation_mark_replica_identity_2), int32(_a_F_relation_mark_replica_identity_3))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v70
	F_errmsg_internal(m, int32(_a_F_relation_mark_replica_identity_4), v14+int32(16))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_relation_mark_replica_identity_1), int32(_a_F_relation_mark_replica_identity_5), int32(_a_F_relation_mark_replica_identity_3))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_swap_relation_files(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 float32
	_ = v365
	var v366 float32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	v20 = m.G0
	v22 = v20 - int32(224)
	m.G0 = v22
	v26 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(l0), int64(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L1
	} else {
		goto L243
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L240
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L1
	} else {
		goto L237
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L1
	} else {
		goto L234
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L228
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L222
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L219
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L216
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L213
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L210
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L1
	} else {
		goto L207
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L204
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L201
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L198
	}
L17:
	;
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v38 = F_SearchSysCacheCopy(m, int32(57), base.I64_extend_i32_u(l1), int64(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L1
	} else {
		goto L195
	}
L21:
	;
	if v38 == int32(0) {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+22)))
	v44 = v42 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+84))
	v46 = v33 + v34
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)+88))
	v49 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+88))
	if base.B2i32(v48 == v49)|base.B2i32(v51 == v49) == v49 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v314 = F_relation_open(m, l0, int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L117
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+88)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v44)+88)) = v48
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+92)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v44)+92)) = v59
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v44)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+84)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v44)+84)) = v63
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)) = uint8(v68)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)) = uint8(v67)
	if l3 != 0 {
		v310 = l7
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v48|v51 != 0 {
		goto L15
	} else {
		goto L28
	}
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+112)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v44)+112)) = v71
	v310 = l7
	goto L23
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v46)+92))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v44)+92))
	if v76 != v77 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+118)))
	if v79 != v80 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	if v45 != v47 {
		goto L12
	} else {
		goto L31
	}
L31:
	;
	if l3 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	if v85 != 0 {
		goto L11
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+117)))
	v88 = int32(0)
	if v87 == v88 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v86 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	if v190 == int32(0) {
		goto L10
	} else {
		goto L75
	}
L38:
	;
	goto L37
L39:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v190 = v185
	goto L38
L40:
	;
	v93 = int32(0)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[0]))
	if v93 < v95 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v136 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[1]))
	if v136 < v138 {
		goto L61
	} else {
		goto L62
	}
L43:
	;
	v99 = v93
	goto L46
L44:
	;
	goto L45
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[2]))
	if v118 <= int32(0) {
		v190 = v88
		goto L38
	} else {
		goto L52
	}
L46:
	;
	v104 = v99 << (uint(int32(3)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+uint32(_c_F_swap_relation_files[3])))
	if v105 == l0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v184 = v104 + int32(_a_F_swap_relation_files_0)
	goto L39
L49:
	;
	goto L50
L50:
	;
	v110 = v99 + int32(1)
	if v110 != v95 {
		v99 = v110
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v123 = int32(0)
	goto L53
L53:
	;
	v128 = v123 << (uint(int32(3)) % 32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_c_F_swap_relation_files[4])))
	if v129 != l0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v184 = v128 + int32(_a_F_swap_relation_files_1)
	goto L39
L55:
	;
	v132 = v123 + int32(1)
	if v118 != v132 {
		v123 = v132
		goto L53
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v190 = v88
	goto L38
L59:
	;
	v166 = int32(0)
	goto L69
L60:
	;
	v184 = v147 + int32(_a_F_swap_relation_files_2)
	goto L39
L61:
	;
	v142 = v136
	goto L64
L62:
	;
	goto L63
L63:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[5]))
	if v159 <= int32(0) {
		v190 = v88
		goto L38
	} else {
		goto L68
	}
L64:
	;
	v147 = v142 << (uint(int32(3)) % 32)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+uint32(_c_F_swap_relation_files[6])))
	if l0 == v148 {
		goto L60
	} else {
		goto L66
	}
L65:
	;
	goto L63
L66:
	;
	v151 = v142 + int32(1)
	if v151 != v138 {
		v142 = v151
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	goto L59
L69:
	;
	v171 = v166 << (uint(int32(3)) % 32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+uint32(_c_F_swap_relation_files[7])))
	if v172 != l0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v184 = v171 + int32(_a_F_swap_relation_files_3)
	goto L39
L71:
	;
	v175 = v166 + int32(1)
	if v159 != v175 {
		v166 = v175
		goto L69
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	goto L70
L74:
	;
	v190 = v88
	goto L38
L75:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+117)))
	v194 = int32(0)
	if v193 == v194 {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	if v296 == int32(0) {
		goto L9
	} else {
		goto L114
	}
L77:
	;
	goto L76
L78:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v296 = v291
	goto L77
L79:
	;
	v199 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[0]))
	if v199 < v201 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v242 = int32(0)
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[1]))
	if v242 < v244 {
		goto L100
	} else {
		goto L101
	}
L82:
	;
	v205 = v199
	goto L85
L83:
	;
	goto L84
L84:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[2]))
	if v224 <= int32(0) {
		v296 = v194
		goto L77
	} else {
		goto L91
	}
L85:
	;
	v210 = v205 << (uint(int32(3)) % 32)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+uint32(_c_F_swap_relation_files[3])))
	if v211 == l1 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L84
L87:
	;
	v290 = v210 + int32(_a_F_swap_relation_files_0)
	goto L78
L88:
	;
	goto L89
L89:
	;
	v216 = v205 + int32(1)
	if v216 != v201 {
		v205 = v216
		goto L85
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v229 = int32(0)
	goto L92
L92:
	;
	v234 = v229 << (uint(int32(3)) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+uint32(_c_F_swap_relation_files[4])))
	if v235 != l1 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v290 = v234 + int32(_a_F_swap_relation_files_1)
	goto L78
L94:
	;
	v238 = v229 + int32(1)
	if v224 != v238 {
		v229 = v238
		goto L92
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L93
L97:
	;
	v296 = v194
	goto L77
L98:
	;
	v272 = int32(0)
	goto L108
L99:
	;
	v290 = v253 + int32(_a_F_swap_relation_files_2)
	goto L78
L100:
	;
	v248 = v242
	goto L103
L101:
	;
	goto L102
L102:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[5]))
	if v265 <= int32(0) {
		v296 = v194
		goto L77
	} else {
		goto L107
	}
L103:
	;
	v253 = v248 << (uint(int32(3)) % 32)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+uint32(_c_F_swap_relation_files[6])))
	if l1 == v254 {
		goto L99
	} else {
		goto L105
	}
L104:
	;
	goto L102
L105:
	;
	v257 = v248 + int32(1)
	if v257 != v244 {
		v248 = v257
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	goto L98
L108:
	;
	v277 = v272 << (uint(int32(3)) % 32)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+uint32(_c_F_swap_relation_files[7])))
	if v278 != l1 {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v290 = v277 + int32(_a_F_swap_relation_files_3)
	goto L78
L110:
	;
	v281 = v272 + int32(1)
	if v265 != v281 {
		v272 = v281
		goto L108
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	goto L109
L113:
	;
	v296 = v194
	goto L77
L114:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+117)))
	F_RelationMapUpdateMap(m, l0, v296, v299, int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+117)))
	F_RelationMapUpdateMap(m, l1, v190, v303, int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = l1
	v310 = l7 + int32(4)
	goto L23
L117:
	;
	v317 = F_relation_open(m, l1, int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+32)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v314)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+36)) = v321
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v314)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+40)) = v323
	v326 = F_GetCurrentSubTransactionId(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v314)+36)) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v314)+40))
	if v328 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	F_relation_close(m, v314, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L126
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v314)+40)) = v326
	goto L122
L121:
	;
	goto L122
L122:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[8]))
	if v333 <= int32(31) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v314)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[8])) = v333 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v333<<(uint(int32(2))%32))+uint32(_c_F_swap_relation_files[9]))) = v336
	goto L119
L124:
	;
	goto L125
L125:
	;
	v347 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_swap_relation_files[10])) = uint8(v347)
	goto L119
L126:
	;
	F_relation_close(m, v317, int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
	if v356 != int32(105) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+140)) = l6
	*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = l5
	goto L130
L129:
	;
	goto L130
L130:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v46)+96))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v44)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+96)) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v44)+96)) = v361
	v365 = *(*float32)(unsafe.Add(mBase, uint32(v46)+100))
	v366 = *(*float32)(unsafe.Add(mBase, uint32(v44)+100))
	*(*float32)(unsafe.Add(mBase, uint32(v46)+100)) = v366
	*(*float32)(unsafe.Add(mBase, uint32(v44)+100)) = v365
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v46)+104))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v44)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+104)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v44)+104)) = v369
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v46)+108))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v44)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+108)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v44)+108)) = v373
	if l2 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v45 != v47 {
		goto L141
	} else {
		goto L142
	}
L132:
	;
	v381 = F_CatalogOpenIndexes(m, v26)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_CacheInvalidateRelcacheByTuple(m, v31)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L139
	}
L135:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v31+int32(4), v31, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_CatalogTupleUpdateWithInfo(m, v26, v38+int32(4), v38, v381)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_CatalogCloseIndexes(m, v381)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	goto L131
L139:
	;
	F_CacheInvalidateRelcacheByTuple(m, v38)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L131
L141:
	;
	v399 = F_changeDependencyFor(m, int32(1259), l0, int32(2601), v47, v45)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	goto L143
L143:
	;
	v410 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[11]))
	if v410 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L144:
	;
	if v399 != int32(1) {
		goto L8
	} else {
		goto L145
	}
L145:
	;
	v405 = F_changeDependencyFor(m, int32(1259), l1, int32(2601), v45, v47)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v405 != int32(1) {
		goto L7
	} else {
		goto L147
	}
L147:
	;
	goto L143
L148:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	if v428 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L149:
	;
	v414 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l0, v414, v414, l4)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_swap_relation_files[11]))
	if v419 == int32(0) {
		goto L148
	} else {
		goto L151
	}
L151:
	;
	v423 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), l1, v423, v423, int32(1))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	goto L148
L153:
	;
	F_pfree(m, v31)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L192
	}
L154:
	;
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
	if v503 != int32(116) {
		goto L153
	} else {
		goto L187
	}
L155:
	;
	if l3 == int32(0) {
		goto L153
	} else {
		goto L186
	}
L156:
	;
	v445 = int32(1)
	if base.Ui32(l0) < base.Ui32(int32(_a_F_swap_relation_files_4)) {
		v453 = v445
		goto L166
	} else {
		goto L167
	}
L157:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v431 == int32(0) {
		goto L155
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	if l3 == int32(0) {
		goto L156
	} else {
		goto L162
	}
L160:
	;
	if l3 == int32(0) {
		goto L156
	} else {
		goto L161
	}
L161:
	;
	goto L3
L162:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v438 == int32(0) {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	F_swap_relation_files(m, v428, v438, l2, int32(1), l4, l5, l6, v310)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	goto L154
L165:
	;
	if v453 != 0 {
		goto L6
	} else {
		goto L169
	}
L166:
	;
	goto L165
L167:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v448 == int32(99) {
		v453 = v445
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v451 = F_isTempToastNamespace(m, v448)
	mBase = m.M
	v453 = v451
	goto L166
L169:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	if v454 != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v457 = F_deleteDependencyRecordsFor(m, int32(1259), v454, int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v462 != 0 {
		goto L175
	} else {
		goto L176
	}
L173:
	;
	if v457 != int32(1) {
		goto L5
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v465 = F_deleteDependencyRecordsFor(m, int32(1259), v462, int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v470 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = v470
	v472 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v472
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v470
	*(*int32)(unsafe.Add(mBase, uint32(v22)+200)) = v472
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v46)+112))
	if v478 != 0 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v465 != int32(1) {
		goto L4
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v478
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l0
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v44)+112))
	if v488 == int32(0) {
		goto L153
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+204)) = v488
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = l1
	F_recordDependencyOn(m, v22+int32(200), v22+int32(212), int32(105))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	goto L153
L186:
	;
	goto L154
L187:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
	if v506 != int32(116) {
		goto L153
	} else {
		goto L188
	}
L188:
	;
	v509 = F_toast_get_valid_index(m, l0)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	v511 = F_toast_get_valid_index(m, l1)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v514 = int32(0)
	F_swap_relation_files(m, v509, v511, l2, int32(1), l4, v514, v514, v310)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	goto L153
L192:
	;
	F_pfree(m, v38)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	F_relation_close(m, v26, int32(3))
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	m.G0 = v22 + int32(224)
	return
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_5), v22)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1650), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_5), v22+int32(16))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1655), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_8), v22+int32(96))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1705), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_9), v22+int32(192))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1716), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_10), v22+int32(176))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1719), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_11), v22+int32(160))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1722), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_12), v22+int32(144))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1726), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+116)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v46 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_13), v22+int32(112))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1734), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v44 + int32(4)
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_13), v22+int32(128))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1738), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
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
	v684 = F_get_rel_namespace(m, l0)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v686 = F_get_namespace_name(m, v684)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	v688 = F_get_rel_name(m, l0)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+84)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v686
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_14), v22+int32(80))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1852), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	v706 = F_get_rel_namespace(m, l1)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	v708 = F_get_namespace_name(m, v706)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v710 = F_get_rel_name(m, l1)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v710
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v708
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_14), v22-int32(-64))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1860), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_15), int32(0))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1922), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v457
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_16), v22+int32(48))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1932), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v465
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_16), v22+int32(32))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1941), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v766 = m.ExcPending
	if v766 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L243:
	;
	F_errmsg_internal(m, int32(_a_F_swap_relation_files_17), int32(0))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	F_errfinish(m, int32(_a_F_swap_relation_files_6), int32(1895), int32(_a_F_swap_relation_files_7))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
