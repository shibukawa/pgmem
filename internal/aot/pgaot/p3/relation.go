package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RelationAssumeNewRelfilelocator(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_RelationAssumeNewRelfilelocator[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v7 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v5
	} else {
	}
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_RelationAssumeNewRelfilelocator[1]))
	if v12 <= int32(31) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, _c_F_RelationAssumeNewRelfilelocator[1])) = v12 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_c_F_RelationAssumeNewRelfilelocator[2]))) = v15
		return
	} else {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_RelationAssumeNewRelfilelocator[3])) = uint8(v26)
		return
	}
}
func F_RelationBuildRuleLock(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[0]))
	v24 = F_AllocSetContextCreateInternal(m, v19, int32(_a_F_RelationBuildRuleLock_0), v2, int32(1024), int32(_a_F_RelationBuildRuleLock_1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v24
	v27 = int32(4)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v31 = F_MemoryContextStrdup(m, v24, v28+v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = v31
	v35 = F_MemoryContextAlloc(m, v24, int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = v14 + int32(-56)
	v39 = int32(3)
	v42 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v38, v39, v39, int32(184), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = F_table_open(m, int32(2618), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v51 = int32(1)
	v54 = F_systable_beginscan(m, v47, int32(2693), v51, int32(0), v51, v38)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v56 = F_systable_getnext(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v59 = v56
	v64 = v27
	v65 = v2
	v66 = v35
	goto L12
L10:
	;
	v195 = v2
	v196 = v35
	goto L11
L11:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L47
	}
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v74 = F_MemoryContextAlloc(m, v24, int32(20))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v195 = v185
	v196 = v179
	goto L11
L14:
	;
	v76 = v71 + v72
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = v77
	v79 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v74)+4)) = v79 - int32(48)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+16)) = uint8(v83)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74)+17)) = uint8(v85)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+18)))
	if v88&int32(2040) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v105 = F_text_to_cstring(m, base.I32_wrap_i64(v103))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	v96 = F_getmissingattr(m, v49, int32(8), v14+int32(-57))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v101 = F_fastgetattr_3(m, v59, int32(8), v49, v14+int32(-57))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	v103 = v96
	goto L15
L20:
	;
	v103 = v101
	goto L15
L21:
	;
	v107 = int32(_a_F_RelationBuildRuleLock_2)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v24
	v111 = F_stringToNode(m, v105)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v111
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v108
	F_pfree(m, v105)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v118)+18)))
	if base.Ui32(v119&int32(2047)) <= base.Ui32(int32(6)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v136 = F_text_to_cstring(m, base.I32_wrap_i64(v134))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v127 = F_getmissingattr(m, v49, int32(7), v14+int32(-57))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v132 = F_fastgetattr_3(m, v59, int32(7), v49, v14+int32(-57))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v134 = v127
	goto L24
L29:
	;
	v134 = v132
	goto L24
L30:
	;
	v138 = int32(_a_F_RelationBuildRuleLock_2)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v24
	v142 = F_stringToNode(m, v136)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v142
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v139
	F_pfree(m, v136)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v151 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_setRuleCheckAsUser(m, v150, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L39
	}
L34:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v149)+80))
	v165 = v163
	goto L33
L35:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	if v154 != int32(118) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v157 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
	if v161 != 0 {
		v165 = int32(0)
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	F_setRuleCheckAsUser(m, v168, v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v64 <= v65 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = F_repalloc(m, v66, v64<<(uint(int32(3))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v178 = v64
	v179 = v66
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+v65<<(uint(int32(2))%32)))) = v74
	v185 = v65 + int32(1)
	v186 = F_systable_getnext(m, v54)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v178 = v64 << (uint(int32(1)) % 32)
	v179 = v174
	goto L43
L45:
	;
	if v186 != 0 {
		v59 = v186
		v64 = v178
		v65 = v185
		v66 = v179
		goto L12
	} else {
		goto L46
	}
L46:
	;
	goto L13
L47:
	;
	F_relation_close(m, v47, int32(1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v195 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	m.G0 = v16 - int32(-64)
	return
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = int64(0)
	F_MemoryContextDelete(m, v24)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v213 = F_MemoryContextAlloc(m, v24, int32(8))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v195
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v213
	goto L49
}
func F_RelationBuildTriggers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v270 int32
	_ = v270
	var v273 int64
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int64
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int64
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int64
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
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
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
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
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	v2 = int32(0)
	v31 = m.G0
	v33 = v31 + int32(-64)
	m.G0 = v33
	v36 = F_palloc(m, int32(960))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v39 = v31 + int32(-56)
	v43 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+56)))
	F_ScanKeyInit(m, v39, int32(2), int32(3), int32(184), v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L84
	}
L5:
	;
	v51 = int32(1)
	v54 = F_systable_beginscan(m, v48, int32(2701), v51, int32(0), v51, v39)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v56 = F_systable_getnext(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v56 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v64 = v56
	v65 = int32(16)
	v68 = v2
	v69 = v36
	goto L11
L9:
	;
	v325 = v2
	v326 = v36
	goto L10
L10:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L59
	}
L11:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	v91 = v89 + v90
	if v65 <= v68 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v325 = v313
	v326 = v100
	goto L10
L13:
	;
	v95 = F_repalloc(m, v69, v65*int32(120))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v99 = v65
	v100 = v69
	goto L15
L15:
	;
	v103 = v100 + v68*int32(60)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104
	v111 = F_DirectFunctionCall1Coll(m, int32(626), int32(0), base.I64_extend_i32_u(v91+int32(12)))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v99 = v65 << (uint(int32(1)) % 32)
	v100 = v95
	goto L15
L17:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v103)+4)) = uint32(v111)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v91)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+8)) = v114
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+12)) = uint16(v116)
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+14)) = uint8(v118)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+83)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+15)) = uint8(v120)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v91)+8))
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+16)) = uint8(base.B2i32(v122 != v123))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v91)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+24)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v91)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v103)+28)) = v130
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+32)) = uint8(v132)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+97)))
	*(*uint8)(unsafe.Add(mBase, uint32(v103)+33)) = uint8(v134)
	v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+98)))
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+34)) = uint16(v136)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v91)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+36)) = uint16(v138)
	v141 = v138 << (uint(int32(16)) % 32)
	if v123 < v141 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(0) < base.I32_extend16_s(v158) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v146 = F_palloc(m, int32(base.Ui32(v141)>>(uint(int32(15))%32)))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = int32(0)
	v158 = v136
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+40)) = v146
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+36)))
	v151 = v149 << (uint(int32(1)) % 32)
	if v151 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	base.MemoryCopy(m, v146, v91+int32(124), v151)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v103)+34)))
	v158 = v155
	goto L18
L26:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v273 = F_fastgetattr_2(m, v64, int32(18), v270, v31+int32(-57))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L42
	}
L27:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v167 = F_fastgetattr_2(m, v64, int32(16), v164, v31+int32(-57))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = int32(0)
	goto L26
L30:
	;
	v170 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v167))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+7)))
	if v172 == int32(1) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+34)))
	v179 = F_palloc(m, v176<<(uint(int32(2))%32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+44)) = v179
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+34)))
	if v182 <= int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v185 = int32(1)
	if v175&v185 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v189 = v185
	goto L37
L36:
	;
	v189 = int32(4)
	goto L37
L37:
	;
	v194 = v170 + v189
	v195 = int32(0)
	goto L38
L38:
	;
	v222 = F_pstrdup(m, v194)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L26
L40:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v103)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v224+v195<<(uint(int32(2))%32)))) = v222
	v229 = F_strlen(m, v194)
	mBase = m.M
	v231 = int32(1)
	v234 = v195 + v231
	v235 = int32(*(*int16)(unsafe.Add(mBase, uint32(v103)+34)))
	if v234 < v235 {
		v194 = v229 + v194 + v231
		v195 = v234
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+7)))
	if v275 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v282 = int32(0)
	goto L45
L44:
	;
	v279 = F_DirectFunctionCall1Coll(m, int32(626), int32(0), v273)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+52)) = v282
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v288 = F_fastgetattr_2(m, v64, int32(19), v285, v31+int32(-57))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v282 = base.I32_wrap_i64(v279)
	goto L45
L47:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+7)))
	if v290 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v297 = int32(0)
	goto L50
L49:
	;
	v294 = F_DirectFunctionCall1Coll(m, int32(626), int32(0), v288)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+56)) = v297
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v303 = F_fastgetattr_2(m, v64, int32(17), v300, v31+int32(-57))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v297 = base.I32_wrap_i64(v294)
	goto L50
L52:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+7)))
	if v305 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v310 = int32(0)
	goto L55
L54:
	;
	v308 = F_text_to_cstring(m, base.I32_wrap_i64(v303))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+48)) = v310
	v313 = v68 + int32(1)
	v314 = F_systable_getnext(m, v54)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v310 = v308
	goto L55
L57:
	;
	if v314 != 0 {
		v64 = v314
		v65 = v99
		v68 = v313
		v69 = v100
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	F_relation_close(m, v48, int32(1))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v325 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	m.G0 = v33 - int32(-64)
	return
L62:
	;
	F_pfree(m, v326)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v356 = F_palloc0(m, int32(32))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v356)+4)) = v325
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v326
	if int32(0) < v325 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+28)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+27)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+25)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+24)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+23)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+22)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+21)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+20)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+19)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+18)))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+17)))
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+16)))
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+15)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+14)))
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+13)))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+12)))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+11)))
	v379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+10)))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+9)))
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+8)))
	v389 = v362
	v390 = v364
	v391 = int32(0)
	v394 = v363
	v395 = v365
	v396 = v366
	v397 = v367
	v398 = v368
	v399 = v369
	v400 = v370
	v401 = v371
	v402 = v372
	v403 = v373
	v404 = v374
	v405 = v375
	v406 = v376
	v407 = v377
	v408 = v378
	v409 = v379
	v410 = v380
	v411 = v381
	goto L70
L68:
	;
	goto L69
L69:
	;
	v559 = int32(_a_F_RelationBuildTriggers_0)
	v560 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0]))
	v563 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0])) = v563
	v565 = F_CopyTriggerDesc(m, v356)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L82
	}
L70:
	;
	v415 = v326 + v391*int32(60)
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+12)))
	v418 = v416 & int32(99)
	v421 = v395 | base.B2i32(v418 == int32(32))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+24)) = uint8(v421)
	v425 = v396 | base.B2i32(v418 == int32(34))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+23)) = uint8(v425)
	v428 = v416 & int32(75)
	v431 = v397 | base.B2i32(v428 == int32(8))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+22)) = uint8(v431)
	v435 = v398 | base.B2i32(v428 == int32(10))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+21)) = uint8(v435)
	v439 = v399 | base.B2i32(v428 == int32(73))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+20)) = uint8(v439)
	v443 = v400 | base.B2i32(v428 == int32(9))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+19)) = uint8(v443)
	v447 = v401 | base.B2i32(v428 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+18)) = uint8(v447)
	v450 = v416 & int32(83)
	v453 = v402 | base.B2i32(v450 == int32(16))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+17)) = uint8(v453)
	v457 = v403 | base.B2i32(v450 == int32(18))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+16)) = uint8(v457)
	v461 = v404 | base.B2i32(v450 == int32(81))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+15)) = uint8(v461)
	v465 = v405 | base.B2i32(v450 == int32(17))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+14)) = uint8(v465)
	v469 = v406 | base.B2i32(v450 == int32(19))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+13)) = uint8(v469)
	v472 = v416 & int32(71)
	v473 = int32(4)
	v475 = v407 | base.B2i32(v472 == v473)
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+12)) = uint8(v475)
	v479 = v408 | base.B2i32(v472 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+11)) = uint8(v479)
	v483 = v409 | base.B2i32(v472 == int32(69))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+10)) = uint8(v483)
	v487 = v410 | base.B2i32(v472 == int32(5))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+9)) = uint8(v487)
	v491 = v411 | base.B2i32(v472 == int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+8)) = uint8(v491)
	if v416&v473 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v415)+56))
	v499 = base.B2i32(v495 != int32(0))
	goto L74
L73:
	;
	v499 = int32(0)
	goto L74
L74:
	;
	v500 = v499 | v390
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+25)) = uint8(v500)
	if v416&int32(16) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v356)+26)))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v415)+52))
	v506 = int32(0)
	v508 = v504 | base.B2i32(v505 != v506)
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+26)) = uint8(v508)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v415)+56))
	v514 = base.B2i32(v510 != v506)
	goto L77
L76:
	;
	v514 = int32(0)
	goto L77
L77:
	;
	v515 = v514 | v394
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+27)) = uint8(v515)
	if v416&int32(8) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v415)+52))
	v523 = base.B2i32(v519 != int32(0))
	goto L80
L79:
	;
	v523 = int32(0)
	goto L80
L80:
	;
	v524 = v523 | v389
	*(*uint8)(unsafe.Add(mBase, uint32(v356)+28)) = uint8(v524)
	v527 = v391 + int32(1)
	if v527 != v325 {
		v389 = v524
		v390 = v500
		v391 = v527
		v394 = v515
		v395 = v421
		v396 = v425
		v397 = v431
		v398 = v435
		v399 = v439
		v400 = v443
		v401 = v447
		v402 = v453
		v403 = v457
		v404 = v461
		v405 = v465
		v406 = v469
		v407 = v475
		v408 = v479
		v409 = v483
		v410 = v487
		v411 = v491
		goto L70
	} else {
		goto L81
	}
L81:
	;
	goto L71
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v565
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0])) = v560
	F_FreeTriggerDesc(m, v356)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L61
L84:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v609 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildTriggers_1), v33)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_RelationBuildTriggers_2), int32(1961), int32(_a_F_RelationBuildTriggers_3))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationClearRelation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v8 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+72))
		v13 = v11 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v13
		if v13 == int32(0) {
			v18 = v8 + int32(76)
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[0]))
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[1]))
				v27 = v22
			} else {
				v24 = int32(_a_F_RelationClearRelation_0)
				*(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[0])) = v24
				v27 = v24
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v27
			v29 = int32(_a_F_RelationClearRelation_0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v18
			*(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[1])) = v18
		} else {
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_smgrclose(m, v36)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
			if v41 != 0 {
				F_pfree(m, v41)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v44)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v44
					v49 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[2]))
					v51 = l0 + int32(56)
					v54 = F_hash_search(m, v49, v51, int32(2), v44)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						if v54 != 0 {
							F_RelationDestroyRelation(m, l0, int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v58 = F_errstart(m, int32(19), int32(0))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								if v58 == int32(0) {
									F_RelationDestroyRelation(m, l0, int32(0))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
										return
									}
								} else {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v62
									F_errmsg_internal(m, int32(_a_F_RelationClearRelation_1), v6)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2557), int32(_a_F_RelationClearRelation_3))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											F_RelationDestroyRelation(m, l0, int32(0))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												m.G0 = v6 + int32(16)
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
				v44 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v44)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v44
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[2]))
				v51 = l0 + int32(56)
				v54 = F_hash_search(m, v49, v51, int32(2), v44)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					if v54 != 0 {
						F_RelationDestroyRelation(m, l0, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v58 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							if v58 == int32(0) {
								F_RelationDestroyRelation(m, l0, int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v62
								F_errmsg_internal(m, int32(_a_F_RelationClearRelation_1), v6)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2557), int32(_a_F_RelationClearRelation_3))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_RelationDestroyRelation(m, l0, int32(0))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											m.G0 = v6 + int32(16)
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
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
		if v41 != 0 {
			F_pfree(m, v41)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				v44 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v44)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v44
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[2]))
				v51 = l0 + int32(56)
				v54 = F_hash_search(m, v49, v51, int32(2), v44)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					if v54 != 0 {
						F_RelationDestroyRelation(m, l0, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							m.G0 = v6 + int32(16)
							return
						}
					} else {
						v58 = F_errstart(m, int32(19), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							if v58 == int32(0) {
								F_RelationDestroyRelation(m, l0, int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
								*(*int32)(unsafe.Add(mBase, uint32(v6))) = v62
								F_errmsg_internal(m, int32(_a_F_RelationClearRelation_1), v6)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2557), int32(_a_F_RelationClearRelation_3))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return
									} else {
										F_RelationDestroyRelation(m, l0, int32(0))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											m.G0 = v6 + int32(16)
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
			v44 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)) = uint8(v44)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v44
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClearRelation[2]))
			v51 = l0 + int32(56)
			v54 = F_hash_search(m, v49, v51, int32(2), v44)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				if v54 != 0 {
					F_RelationDestroyRelation(m, l0, int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						m.G0 = v6 + int32(16)
						return
					}
				} else {
					v58 = F_errstart(m, int32(19), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						if v58 == int32(0) {
							F_RelationDestroyRelation(m, l0, int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								m.G0 = v6 + int32(16)
								return
							}
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v62
							F_errmsg_internal(m, int32(_a_F_RelationClearRelation_1), v6)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2557), int32(_a_F_RelationClearRelation_3))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_RelationDestroyRelation(m, l0, int32(0))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										m.G0 = v6 + int32(16)
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
func F_RelationCreateStorage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = int32(-1)
	switch l1 - int32(112) {
	case 0:
		v38 = v12
		v39 = int32(1)
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v40
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v42
		v46 = F_smgropen(m, v10+int32(16), v38)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			v48 = int32(0)
			F_smgrcreate(m, v46, v48, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v39 != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v52
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v67 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
									v72 = F_MemoryContextAlloc(m, v70, int32(28))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
										v78 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
										*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
										v85 = int32(_a_F_RelationCreateStorage_0)
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
											if int32(0) < v94 {
												m.G0 = v10 + int32(80)
												return v46
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
												if v98 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
														v115 = v112
														v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
															m.G0 = v10 + int32(80)
															return v46
														}
													}
												} else {
													v115 = v98
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v46
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
										if int32(0) < v94 {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
											if v98 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
													v115 = v112
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											} else {
												v115 = v98
												v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
													m.G0 = v10 + int32(80)
													return v46
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
						v72 = F_MemoryContextAlloc(m, v70, int32(28))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
							*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
							v85 = int32(_a_F_RelationCreateStorage_0)
							v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
							*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
								if int32(0) < v94 {
									m.G0 = v10 + int32(80)
									return v46
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
									if v98 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
											v115 = v112
											v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v121 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
												m.G0 = v10 + int32(80)
												return v46
											}
										}
									} else {
										v115 = v98
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v46
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
							if int32(0) < v94 {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
								if v98 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
										v115 = v112
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								} else {
									v115 = v98
									v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
										m.G0 = v10 + int32(80)
										return v46
									}
								}
							}
						}
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
			F_errmsg_internal(m, int32(_a_F_RelationCreateStorage_2), v10)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_RelationCreateStorage_3), int32(146), int32(_a_F_RelationCreateStorage_4))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 4:
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[6]))
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[7]))
		if v34 == int32(-1) {
			v37 = v32
		} else {
			v37 = v34
		}
		v38 = v37
		v39 = v4
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v40
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v42
		v46 = F_smgropen(m, v10+int32(16), v38)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			v48 = int32(0)
			F_smgrcreate(m, v46, v48, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v39 != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v52
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v67 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
									v72 = F_MemoryContextAlloc(m, v70, int32(28))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
										v78 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
										*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
										v85 = int32(_a_F_RelationCreateStorage_0)
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
											if int32(0) < v94 {
												m.G0 = v10 + int32(80)
												return v46
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
												if v98 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
														v115 = v112
														v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
															m.G0 = v10 + int32(80)
															return v46
														}
													}
												} else {
													v115 = v98
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v46
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
										if int32(0) < v94 {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
											if v98 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
													v115 = v112
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											} else {
												v115 = v98
												v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
													m.G0 = v10 + int32(80)
													return v46
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
						v72 = F_MemoryContextAlloc(m, v70, int32(28))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
							*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
							v85 = int32(_a_F_RelationCreateStorage_0)
							v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
							*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
								if int32(0) < v94 {
									m.G0 = v10 + int32(80)
									return v46
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
									if v98 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
											v115 = v112
											v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v121 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
												m.G0 = v10 + int32(80)
												return v46
											}
										}
									} else {
										v115 = v98
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v46
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
							if int32(0) < v94 {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
								if v98 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
										v115 = v112
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								} else {
									v115 = v98
									v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
										m.G0 = v10 + int32(80)
										return v46
									}
								}
							}
						}
					}
				}
			}
		}
	case 5:
		v38 = v12
		v39 = v4
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v40
		v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v42
		v46 = F_smgropen(m, v10+int32(16), v38)
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			v48 = int32(0)
			F_smgrcreate(m, v46, v48, v48)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				if v39 != 0 {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v52
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v46)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v54
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v67 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
									v72 = F_MemoryContextAlloc(m, v70, int32(28))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
										v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
										v78 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
										*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
										v85 = int32(_a_F_RelationCreateStorage_0)
										v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
										*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
											if int32(0) < v94 {
												m.G0 = v10 + int32(80)
												return v46
											} else {
												v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
												if v98 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
														v115 = v112
														v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return int32(0)
														} else {
															v121 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
															m.G0 = v10 + int32(80)
															return v46
														}
													}
												} else {
													v115 = v98
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v46
									} else {
										v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
										if int32(0) < v94 {
											m.G0 = v10 + int32(80)
											return v46
										} else {
											v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
											if v98 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
													v115 = v112
													v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return int32(0)
													} else {
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
														m.G0 = v10 + int32(80)
														return v46
													}
												}
											} else {
												v115 = v98
												v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return int32(0)
												} else {
													v121 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
													m.G0 = v10 + int32(80)
													return v46
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[0]))
						v72 = F_MemoryContextAlloc(m, v70, int32(28))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
							v76 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v72))) = v76
							v78 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v78)
							*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = v38
							v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[1]))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+20)) = v83
							v85 = int32(_a_F_RelationCreateStorage_0)
							v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2]))
							*(*int32)(unsafe.Add(mBase, uint32(v72)+24)) = v86
							*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[2])) = v72
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
								if int32(0) < v94 {
									m.G0 = v10 + int32(80)
									return v46
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
									if v98 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
											v115 = v112
											v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												v121 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
												m.G0 = v10 + int32(80)
												return v46
											}
										}
									} else {
										v115 = v98
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v46
						} else {
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[3]))
							if int32(0) < v94 {
								m.G0 = v10 + int32(80)
								return v46
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4]))
								if v98 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+68)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int64(16), v10+int32(32), int32(1064))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[4])) = v112
										v115 = v112
										v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											v121 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
											m.G0 = v10 + int32(80)
											return v46
										}
									}
								} else {
									v115 = v98
									v119 = F_hash_search(m, v115, l0, int32(1), v10+int32(32))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int32(0)
									} else {
										v121 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v119)+12)) = uint8(v121)
										m.G0 = v10 + int32(80)
										return v46
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
func F_RelationFindReplTupleByIndex(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	v15 = m.G0
	v17 = v15 - int32(1888)
	m.G0 = v17
	v20 = F_index_open(m, l1, int32(3))
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
	v24 = F_GetRelationIdentityOrPK(m, l0)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(4)
	v35 = F_build_replindex_scan_key(m, v17+int32(96), v20, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = int32(0)
	v39 = F_index_beginscan(m, l0, v20, v17+int32(24), int32(0), v35, v37, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = int32(0)
	goto L6
L6:
	;
	v58 = int32(0)
	F_index_rescan(m, v39, v17+int32(96), v35, v58, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_index_endscan(m, v39)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L40
	}
L8:
	;
	v63 = v47
	goto L10
L9:
	;
	goto L7
L10:
	;
	v77 = F_index_getnext_slot(m, v39, int32(1), l3)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+28))
	m.T0[v96].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	if v77 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	if l1 == v24 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L11
L15:
	;
	v94 = v47
	goto L14
L16:
	;
	goto L17
L17:
	;
	if v63 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v86 = F_palloc0_mul(m, int32(4), v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v88 = v63
	goto L20
L20:
	;
	v90 = F_tuples_equal(m, l3, l2, v88, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v88 = v86
	goto L20
L22:
	;
	if v90 == int32(0) {
		v63 = v88
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v94 = v88
	goto L14
L24:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	if v99 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v101 = v99
	goto L27
L26:
	;
	v101 = v100
	goto L27
L27:
	;
	if v101 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v102 = int32(0)
	F_XactLockTableWait(m, v101, v102, v102, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v107 = F_GetLatestSnapshot(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v47 = v94
	goto L6
L32:
	;
	F_PushActiveSnapshot(m, v107)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_RelationFindReplTupleByIndex[0]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	goto L34
L34:
	;
	v115 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v118 = int32(0)
	v121 = v17 + int32(4)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+104))
	v124 = m.T0[v123].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3+int32(32), v113, l3, v115, int32(3), v118, v118, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v128 = F_should_refetch_tuple(m, v124, v121)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v128 != 0 {
		v47 = v94
		goto L6
	} else {
		goto L39
	}
L39:
	;
	goto L9
L40:
	;
	F_relation_close(m, v20, int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	m.G0 = v17 + int32(1888)
	return v77
}
func F_RelationFindReplTupleSeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_palloc0_mul(m, int32(4), v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(4)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_RelationFindReplTupleSeq[0]))
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationFindReplTupleSeq[1])))
	if v28&int32(1) == int32(0) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v37 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v43 = m.T0[v42].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v14+int32(24), v37, v37, v37, int32(449))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L6
L8:
	;
	v46 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	goto L10
L10:
	;
	v59 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+188))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	m.T0[v66].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v43, v59, v59, v59, v59, v59)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+188))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	m.T0[v135].(func(*base.Module, int32))(m, v43)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L36
	}
L12:
	;
	goto L14
L13:
	;
	goto L11
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v81
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+188))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+20))
	v87 = m.T0[v86].(func(*base.Module, int32, int32, int32) int32)(m, v43, int32(1), v46)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	m.T0[v97].(func(*base.Module, int32, int32))(m, l2, v46)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L20
	}
L16:
	;
	if v87 == int32(0) {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v92 = F_tuples_equal(m, v46, l1, v19, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v92 == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	if v100 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v102 = v100
	goto L23
L22:
	;
	v102 = v101
	goto L23
L23:
	;
	if v102 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v103 = int32(0)
	F_XactLockTableWait(m, v102, v103, v103, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v108 = F_GetLatestSnapshot(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L10
L28:
	;
	F_PushActiveSnapshot(m, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_RelationFindReplTupleSeq[2]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	goto L30
L30:
	;
	v116 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v119 = int32(0)
	v122 = v14 + int32(4)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+104))
	v125 = m.T0[v124].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l2+int32(32), v114, l2, v116, int32(3), v119, v119, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v129 = F_should_refetch_tuple(m, v125, v122)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v129 != 0 {
		goto L10
	} else {
		goto L35
	}
L35:
	;
	goto L13
L36:
	;
	F_ExecDropSingleTupleTableSlot(m, v46)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	m.G0 = v14 + int32(96)
	return v87
L38:
	;
	F_errmsg_internal(m, int32(_a_F_RelationFindReplTupleSeq_0), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_RelationFindReplTupleSeq_1), int32(931), int32(_a_F_RelationFindReplTupleSeq_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetBufferForTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v170 int32
	_ = v170
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v432 int32
	_ = v432
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v493 int32
	_ = v493
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int64
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int64
	_ = v722
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int64
	_ = v739
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v927 int32
	_ = v927
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v984 int32
	_ = v984
	var v986 int32
	_ = v986
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1065 int32
	_ = v1065
	var v1074 int32
	_ = v1074
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1126 int32
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1140 int32
	_ = v1140
	var v1149 int32
	_ = v1149
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1170 int32
	_ = v1170
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int64
	_ = v1177
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1254 int32
	_ = v1254
	var v1257 int32
	_ = v1257
	var v1263 int32
	_ = v1263
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1279 int32
	_ = v1279
	var v1284 int32
	_ = v1284
	v25 = m.G0
	v27 = v25 - int32(368)
	m.G0 = v27
	v32 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(v32) < base.Ui32(int32(_a_F_RelationGetBufferForTuple_0)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L33
	} else {
		goto L376
	}
L2:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v37 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L33
	} else {
		goto L372
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v44 = base.I32_div_s(int32(_a_F_RelationGetBufferForTuple_1)-v39<<(uint(int32(13))%32), int32(100))
	v46 = v44
	goto L7
L6:
	;
	v46 = int32(0)
	goto L7
L7:
	;
	v47 = v46 + v32
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l2 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v68 = int32(-1)
	goto L10
L10:
	;
	if base.Ui32(int32(_a_F_RelationGetBufferForTuple_2)) < base.Ui32(v32) {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	v68 = v66
	goto L10
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[0]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l2^int32(-1))*int32(56))+16))
	v66 = v57
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v60 = int32(56)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l2*v60-v60)+16))
	v66 = v65
	goto L11
L15:
	;
	v70 = v32
	goto L17
L16:
	;
	v70 = int32(_a_F_RelationGetBufferForTuple_2)
	goto L17
L17:
	;
	v74 = l3 & int32(2)
	if l4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if base.Ui32(int32(_a_F_RelationGetBufferForTuple_2)) < base.Ui32(v47) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v101 == int32(0) {
		v106 = int32(-1)
		goto L18
	} else {
		goto L26
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if v77 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v106 = v98
	goto L18
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[0]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v77^int32(-1))*int32(56))+16))
	v98 = v89
	goto L22
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v92 = int32(56)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+v77*v92-v92)+16))
	v98 = v97
	goto L22
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+16))
	v106 = v104
	goto L18
L27:
	;
	v107 = v70
	goto L29
L28:
	;
	v107 = v47
	goto L29
L29:
	;
	if v74|base.B2i32(v106 != int32(-1)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v113 = F_GetPageWithFreeSpace(m, l0, v107)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v117 = v106
	goto L32
L32:
	;
	if v117 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	return int32(0)
L34:
	;
	v117 = v113
	goto L32
L35:
	;
	v121 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v125 = v117
	goto L37
L37:
	;
	v126 = int32(1)
	if l7 <= v126 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v125 = v121 - int32(1)
	goto L37
L39:
	;
	v129 = v126
	goto L41
L40:
	;
	v129 = l7
	goto L41
L41:
	;
	v131 = l3 & int32(4)
	v134 = int32(0)
	v135 = base.B2i32(v74 == v134)
	v144 = v125
	goto L45
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1224)+16)) = v1223
	m.G0 = v27 + int32(368)
	return v1229
L43:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+72))
	if v1209 != 0 {
		goto L369
	} else {
		goto L370
	}
L44:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1173 != 0 {
		v1223 = v974
		v1224 = v1173
		v1229 = v754
		goto L42
	} else {
		goto L366
	}
L45:
	;
	if v144 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v1107 = int32(4)
	v1108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+14)))
	v1109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+12)))
	v1110 = v1108 - v1109
	if v1110 <= v1107 {
		goto L347
	} else {
		goto L348
	}
L47:
	;
	v596 = int32(1)
	if v135|base.B2i32(l4 != v134) != 0 {
		goto L196
	} else {
		goto L197
	}
L48:
	;
	v170 = v144
	goto L49
L49:
	;
	if l2 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L47
L51:
	;
	F_LockBufferInternal(m, v377, int32(3))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L33
	} else {
		goto L124
	}
L52:
	;
	if l4 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	if v170 == v68 {
		goto L93
	} else {
		goto L94
	}
L55:
	;
	v243 = int32(0)
	v244 = base.B2i32(v243 <= v241)
	if v244 == v243 {
		goto L75
	} else {
		goto L76
	}
L56:
	;
	v197 = int32(0)
	v200 = F_ReadBufferExtended(m, l0, v197, v170, v197, v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L33
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v202 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v241 = v200
	goto L55
L60:
	;
	if v202 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v233 = int32(0)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v236 = F_ReadBufferExtended(m, l0, v233, v170, v233, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L33
	} else {
		goto L72
	}
L63:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v170 == v221 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[0]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206+(v202^int32(-1))*int32(56))+16))
	v221 = v212
	goto L63
L65:
	;
	goto L66
L66:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v215 = int32(56)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+v202*v215-v215)+16))
	v221 = v220
	goto L63
L67:
	;
	F_IncrBufferRefCount(m, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L33
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_ReleaseBuffer(m, v222)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L33
	} else {
		goto L71
	}
L70:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v241 = v226
	goto L55
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(0)
	goto L62
L72:
	;
	F_IncrBufferRefCount(m, v236)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L33
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v236
	v241 = v236
	goto L55
L74:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+10)))
	if v263&int32(4) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v248+(v241^int32(-1))<<(uint(int32(2))%32))))
	v262 = v254
	goto L74
L76:
	;
	goto L77
L77:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v262 = v256 + v241<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	F_visibilitymap_pin(m, l0, v170, l5)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L33
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v131 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	v377 = v241
	v378 = v241
	goto L51
L83:
	;
	goto L84
L84:
	;
	if v244 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v287)+12)))
	v295 = int32(0)
	if base.B2i32(base.Ui32(v288) < base.Ui32(int32(25)))|base.B2i32((v288+int32(_a_F_RelationGetBufferForTuple_3))&int32(_a_F_RelationGetBufferForTuple_4) == v295) == v295 {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v273+(v241^int32(-1))<<(uint(int32(2))%32))))
	v287 = v279
	goto L85
L87:
	;
	goto L88
L88:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v287 = v281 + v241<<(uint(int32(13))%32) + int32(-8192)
	goto L85
L89:
	;
	v377 = v241
	v378 = v241
	goto L51
L90:
	;
	goto L91
L91:
	;
	F_visibilitymap_pin(m, l0, v170, l5)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L33
	} else {
		goto L92
	}
L92:
	;
	v377 = v241
	v378 = v241
	goto L51
L93:
	;
	if l2 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	goto L95
L95:
	;
	v322 = F_ReadBuffer(m, l0, v170)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L33
	} else {
		goto L102
	}
L96:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+10)))
	if v315&int32(4) == int32(0) {
		v377 = l2
		v378 = l2
		goto L51
	} else {
		goto L100
	}
L97:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v306+(l2^int32(-1))<<(uint(int32(2))%32))))
	v314 = v308
	goto L96
L98:
	;
	goto L99
L99:
	;
	v310 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v314 = v310 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	F_visibilitymap_pin(m, l0, v68, l5)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L33
	} else {
		goto L101
	}
L101:
	;
	v377 = l2
	v378 = l2
	goto L51
L102:
	;
	if base.Ui32(v68) < base.Ui32(v170) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v322 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	goto L105
L105:
	;
	if v322 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L106:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+10)))
	if v343&int32(4) != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328+(v322^int32(-1))<<(uint(int32(2))%32))))
	v342 = v334
	goto L106
L108:
	;
	goto L109
L109:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v342 = v336 + v322<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L110:
	;
	F_visibilitymap_pin(m, l0, v170, l5)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L33
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_LockBufferInternal(m, l2, int32(3))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L33
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v377 = v322
	v378 = v322
	goto L51
L115:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v368)+10)))
	if v369&int32(4) != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v354 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v354+(v322^int32(-1))<<(uint(int32(2))%32))))
	v368 = v360
	goto L115
L117:
	;
	goto L118
L118:
	;
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v368 = v362 + v322<<(uint(int32(13))%32) + int32(-8192)
	goto L115
L119:
	;
	F_visibilitymap_pin(m, l0, v170, l5)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L33
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	F_LockBufferInternal(m, v322, int32(3))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L33
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v377 = l2
	v378 = v322
	goto L51
L124:
	;
	v382 = F_GetVisibilityMapPins(m, l0, v378, l2, v170, v68, l5, l6)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L33
	} else {
		goto L125
	}
L125:
	;
	if v378 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+14)))
	if v402 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v387+(v378^int32(-1))<<(uint(int32(2))%32))))
	v401 = v393
	goto L126
L128:
	;
	goto L129
L129:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v401 = v395 + v378<<(uint(int32(13))%32) + int32(-8192)
	goto L126
L130:
	;
	v405 = int32(_a_F_RelationGetBufferForTuple_5)
	v406 = int32(0)
	if v406|(v401&int32(3)|int32(1)) == v406 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	goto L132
L132:
	;
	v460 = int32(4)
	v461 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+14)))
	v462 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+12)))
	v463 = v461 - v462
	if v463 <= v460 {
		goto L146
	} else {
		goto L147
	}
L133:
	;
	F_MarkBufferDirty(m, v378)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L33
	} else {
		goto L144
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v401)+10)) = int32(_a_F_RelationGetBufferForTuple_6)
	v446 = int32(_a_F_RelationGetBufferForTuple_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v401)+18)) = uint16(v446)
	v452 = int32(_a_F_RelationGetBufferForTuple_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v401)+16)) = uint16(v452)
	*(*uint16)(unsafe.Add(mBase, uint32(v401)+14)) = uint16(v452)
	goto L133
L135:
	;
	goto L138
L136:
	;
	goto L137
L137:
	;
	goto L143
L138:
	;
	v423 = v401 + v405
	v425 = v401 + int32(4)
	if base.Ui32(v425) < base.Ui32(v423) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v427 = v423
	goto L141
L140:
	;
	v427 = v425
	goto L141
L141:
	;
	v432 = (v401^int32(-1)+v427)&int32(-4) + int32(4)
	if v432 == int32(0) {
		goto L134
	} else {
		goto L142
	}
L142:
	;
	base.MemoryFill(m, v401, int32(0), v432)
	goto L134
L143:
	;
	base.MemoryFill(m, v401, int32(0), v405)
	goto L134
L144:
	;
	goto L132
L145:
	;
	if base.Ui32(v107) <= base.Ui32(v523) {
		goto L164
	} else {
		goto L165
	}
L146:
	;
	v466 = v460
	goto L148
L147:
	;
	v466 = v463
	goto L148
L148:
	;
	v468 = v466 - int32(4)
	if v468 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v523 = int32(0)
	goto L145
L150:
	;
	goto L151
L151:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v462) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v523 = v468
	goto L145
L153:
	;
	v479 = int32(base.Ui32(v462+int32(_a_F_RelationGetBufferForTuple_3)) >> (uint(int32(2)) % 32))
	goto L155
L154:
	;
	v479 = int32(0)
	goto L155
L155:
	;
	if base.Ui32(v479&int32(_a_F_RelationGetBufferForTuple_8)) < base.Ui32(int32(291)) {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401)+10)))
	if v484&int32(1) == int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v523 = int32(0)
	goto L145
L158:
	;
	goto L159
L159:
	;
	v493 = int32(1)
	goto L160
L160:
	;
	v502 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401+int32(20)+v493&int32(_a_F_RelationGetBufferForTuple_8)<<(uint(int32(2))%32))+1)))
	if v502&int32(384) == int32(0) {
		goto L152
	} else {
		goto L162
	}
L161:
	;
	v523 = int32(0)
	goto L145
L162:
	;
	v508 = v493 + int32(1)
	v509 = int32(_a_F_RelationGetBufferForTuple_8)
	if base.Ui32(v508&v509) <= base.Ui32(v479&v509) {
		v493 = v508
		goto L160
	} else {
		goto L163
	}
L163:
	;
	goto L161
L164:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v525 != 0 {
		v1223 = v170
		v1224 = v525
		v1229 = v378
		goto L42
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	if l2 == int32(0) {
		goto L170
	} else {
		goto L171
	}
L167:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v527
	v529 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+80)) = v529
	v533 = F_smgropen(m, v27+int32(80), v526)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L33
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v533
	v1185 = v170
	v1186 = v533
	v1191 = v378
	goto L43
L169:
	;
	if l4 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L170:
	;
	F_UnlockReleaseBuffer(m, v378)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L33
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	if v170 != v68 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	goto L169
L174:
	;
	F_UnlockReleaseBuffer(m, v378)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L33
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	F_UnlockBuffer(m, v378)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L33
	} else {
		goto L179
	}
L177:
	;
	F_UnlockBuffer(m, l2)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L33
	} else {
		goto L178
	}
L178:
	;
	goto L169
L179:
	;
	goto L169
L180:
	;
	if v569 != int32(-1) {
		v170 = v569
		goto L49
	} else {
		goto L193
	}
L181:
	;
	if v74 != 0 {
		goto L47
	} else {
		goto L191
	}
L182:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v549 == int32(-1) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	if v74 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	F_RecordPageWithFreeSpace(m, l0, v170, v523)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L33
	} else {
		goto L187
	}
L185:
	;
	v557 = v549
	goto L186
L186:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(v558) <= base.Ui32(v557) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v557 = v556
	goto L186
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = int64(-1)
	v569 = v557
	goto L180
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v557 + int32(1)
	v569 = v557
	goto L180
L191:
	;
	v566 = F_RecordAndGetPageWithFreeSpace(m, l0, v170, v523, v107)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L33
	} else {
		goto L192
	}
L192:
	;
	v569 = v566
	goto L180
L193:
	;
	goto L50
L194:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v720
	v722 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v722
	v724 = m.G0
	v726 = v724 - int32(16)
	m.G0 = v726
	v729 = v27 - int32(-64)
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729)+8)))
	if v730 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L195:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(v696) < base.Ui32(v676) {
		goto L220
	} else {
		goto L221
	}
L196:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v598 != 0 {
		v676 = v129
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v686 = v596
	goto L198
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v686
	v714 = v686
	v716 = v596
	v719 = int32(0)
	goto L194
L199:
	;
	if l4 != 0 {
		goto L195
	} else {
		goto L216
	}
L200:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v599 != 0 {
		v676 = v129
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v600 = m.G0
	v601 = int32(16)
	v602 = v600 - v601
	m.G0 = v602
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v602))) = v604
	v606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v602)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v602)+4)) = v606
	v611 = m.G0
	v613 = v611 - v601
	m.G0 = v613
	v615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v615-int32(3))&int32(255)) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	m.G0 = v602 + int32(16)
	v676 = (v649 + int32(1)) * v129
	goto L199
L203:
	;
	v623 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[4]))
	v624 = F_get_hash_value(m, v623, v602)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L33
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L33
	} else {
		goto L213
	}
L206:
	;
	v627 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[5]))
	v634 = v627 + v624&int32(15)<<(uint(int32(7))%32) + int32(_a_F_RelationGetBufferForTuple_9)
	v636 = F_LWLockAcquire(m, v634, int32(0))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L33
	} else {
		goto L207
	}
L207:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[4]))
	v643 = F_hash_search_with_hash_value(m, v639, v602, v624, int32(0), v613+int32(15))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L33
	} else {
		goto L208
	}
L208:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613)+15)))
	if v645 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v643)+84))
	v649 = v648
	goto L211
L210:
	;
	v649 = int32(0)
	goto L211
L211:
	;
	F_LWLockRelease(m, v634)
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L33
	} else {
		goto L212
	}
L212:
	;
	m.G0 = v613 + int32(16)
	goto L202
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v613))) = v615
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_10), v613)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L33
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_11), int32(_a_F_RelationGetBufferForTuple_12), int32(_a_F_RelationGetBufferForTuple_13))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L33
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
	v680 = int32(64)
	if base.Ui32(v680) <= base.Ui32(v676) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v683 = v680
	goto L219
L218:
	;
	v683 = v676
	goto L219
L219:
	;
	v686 = v683
	goto L198
L220:
	;
	v698 = v676
	goto L222
L221:
	;
	v698 = v696
	goto L222
L222:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v698) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v701 = int32(64)
	goto L225
L224:
	;
	v701 = v698
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v701
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v703 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_ReleaseBuffer(m, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L33
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l0
	v711 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v714 = v701
	v716 = v129
	v719 = v711
	goto L194
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(0)
	goto L228
L230:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v729)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v733)+48))
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v734)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v729)+8)) = uint8(v735)
	goto L232
L231:
	;
	goto L232
L232:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v729)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v726)+8)) = v737
	v739 = *(*int64)(unsafe.Add(mBase, uint32(v729)))
	*(*int64)(unsafe.Add(mBase, uint32(v726))) = v739
	v748 = F_ExtendBufferedRelCommon(m, v726, int32(0), v719, int32(8), v714, int32(-1), v27+int32(112), v27+int32(108))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L33
	} else {
		goto L233
	}
L233:
	;
	m.G0 = v726 + int32(16)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	v755 = int32(0)
	v756 = base.B2i32(v755 <= v754)
	if v756 == v755 {
		goto L236
	} else {
		goto L237
	}
L234:
	;
	v927 = v748 + v753
	if v74|base.B2i32(base.Ui32(v906) <= base.Ui32(v716)) == int32(0) {
		goto L270
	} else {
		goto L271
	}
L235:
	;
	v775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v774)+14)))
	if v775 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	v760 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v760+(v754^int32(-1))<<(uint(int32(2))%32))))
	v774 = v766
	goto L235
L237:
	;
	goto L238
L238:
	;
	v768 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v774 = v768 + v754<<(uint(int32(13))%32) + int32(-8192)
	goto L235
L239:
	;
	v778 = int32(_a_F_RelationGetBufferForTuple_5)
	v779 = int32(0)
	if v779|(v774&int32(3)|int32(1)) == v779 {
		goto L244
	} else {
		goto L245
	}
L240:
	;
	goto L241
L241:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L33
	} else {
		goto L267
	}
L242:
	;
	F_MarkBufferDirty(m, v754)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L33
	} else {
		goto L253
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v774)+10)) = int32(_a_F_RelationGetBufferForTuple_6)
	v819 = int32(_a_F_RelationGetBufferForTuple_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v774)+18)) = uint16(v819)
	v825 = int32(_a_F_RelationGetBufferForTuple_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v774)+16)) = uint16(v825)
	*(*uint16)(unsafe.Add(mBase, uint32(v774)+14)) = uint16(v825)
	goto L242
L244:
	;
	goto L247
L245:
	;
	goto L246
L246:
	;
	goto L252
L247:
	;
	v796 = v774 + v778
	v798 = v774 + int32(4)
	if base.Ui32(v798) < base.Ui32(v796) {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v800 = v796
	goto L250
L249:
	;
	v800 = v798
	goto L250
L250:
	;
	v805 = (v774^int32(-1)+v800)&int32(-4) + int32(4)
	if v805 == int32(0) {
		goto L243
	} else {
		goto L251
	}
L251:
	;
	base.MemoryFill(m, v774, int32(0), v805)
	goto L243
L252:
	;
	base.MemoryFill(m, v774, int32(0), v778)
	goto L243
L253:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v832 = v135 & base.B2i32(base.Ui32(v716) < base.Ui32(v830))
	if v832 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	F_UnlockBuffer(m, v754)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L33
	} else {
		goto L257
	}
L255:
	;
	v836 = v830
	goto L256
L256:
	;
	v837 = int32(1)
	if base.Ui32(v836) <= base.Ui32(v837) {
		v906 = v836
		goto L234
	} else {
		goto L258
	}
L257:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v836 = v835
	goto L256
L258:
	;
	v841 = v837
	goto L259
L259:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(112)+v841<<(uint(int32(2))%32))))
	F_ReleaseBuffer(m, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L33
	} else {
		goto L261
	}
L260:
	;
	v906 = v882
	goto L234
L261:
	;
	if v74|base.B2i32(base.Ui32(v841) < base.Ui32(v716)) == int32(0) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	F_RecordPageWithFreeSpace(m, l0, v841+v748, int32(_a_F_RelationGetBufferForTuple_14))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L33
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v881 = v841 + int32(1)
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	if base.Ui32(v881) < base.Ui32(v882) {
		v841 = v881
		goto L259
	} else {
		goto L266
	}
L265:
	;
	goto L264
L266:
	;
	goto L260
L267:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v748
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v888 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_15), v27+int32(48))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L33
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(359), int32(_a_F_RelationGetBufferForTuple_17))
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L33
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v748+v716, v927)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L33
	} else {
		goto L273
	}
L271:
	;
	goto L272
L272:
	;
	if l4 != 0 {
		goto L274
	} else {
		goto L275
	}
L273:
	;
	goto L272
L274:
	;
	v935 = int32(1)
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v940 = base.B2i32(base.Ui32(v935) < base.Ui32(v938))
	if base.Ui32(v935) < base.Ui32(v938) {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	goto L276
L276:
	;
	if v754 < int32(0) {
		goto L285
	} else {
		goto L286
	}
L277:
	;
	v941 = v927 - v935
	goto L279
L278:
	;
	v941 = int32(-1)
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v941
	if base.Ui32(v935) < base.Ui32(v938) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v946 = v748 + int32(1)
	goto L282
L281:
	;
	v946 = int32(-1)
	goto L282
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v946
	F_IncrBufferRefCount(m, v754)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L33
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v754
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v951 + v952
	goto L276
L284:
	;
	if v756 == int32(0) {
		goto L289
	} else {
		goto L290
	}
L285:
	;
	v959 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[0]))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v959+(v754^int32(-1))*int32(56))+16))
	v974 = v965
	goto L284
L286:
	;
	goto L287
L287:
	;
	v967 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v968 = int32(56)
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v967+v754*v968-v968)+16))
	v974 = v973
	goto L284
L288:
	;
	if v131 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L289:
	;
	v978 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v978+(v754^int32(-1))<<(uint(int32(2))%32))))
	v992 = v984
	goto L288
L290:
	;
	goto L291
L291:
	;
	v986 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v992 = v986 + v754<<(uint(int32(13))%32) + int32(-8192)
	goto L288
L292:
	;
	goto L46
L293:
	;
	v1027 = F_GetVisibilityMapPins(m, l0, l2, v754, v68, v974, l6, l5)
	mBase = m.M
	v1028 = m.ExcPending
	if v1028 != 0 {
		goto L33
	} else {
		goto L319
	}
L294:
	;
	F_LockBufferInternal(m, v754, int32(3))
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L33
	} else {
		goto L318
	}
L295:
	;
	F_LockBufferInternal(m, l2, int32(3))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L33
	} else {
		goto L317
	}
L296:
	;
	if v832 != 0 {
		goto L309
	} else {
		goto L310
	}
L297:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v995 == int32(0) {
		goto L299
	} else {
		goto L300
	}
L298:
	;
	if v1003 != 0 {
		goto L296
	} else {
		goto L302
	}
L299:
	;
	v1003 = int32(0)
	goto L298
L300:
	;
	goto L301
L301:
	;
	v999 = F_BufferGetBlockNumber(m, v995)
	mBase = m.M
	v1001 = base.I32_div_u_s(v974, int32(_a_F_RelationGetBufferForTuple_18))
	v1003 = base.B2i32(v999 == v1001)
	goto L298
L302:
	;
	if v832 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L303:
	;
	F_UnlockBuffer(m, v754)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L33
	} else {
		goto L306
	}
L304:
	;
	goto L305
L305:
	;
	F_visibilitymap_pin(m, l0, v974, l5)
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L33
	} else {
		goto L307
	}
L306:
	;
	goto L305
L307:
	;
	if l2 != 0 {
		goto L295
	} else {
		goto L308
	}
L308:
	;
	goto L294
L309:
	;
	if l2 != 0 {
		goto L295
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	if l2 == int32(0) {
		goto L292
	} else {
		goto L313
	}
L312:
	;
	goto L294
L313:
	;
	v1013 = F_ConditionalLockBuffer(m, l2)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L33
	} else {
		goto L314
	}
L314:
	;
	if v1013 != 0 {
		v1026 = int32(0)
		goto L293
	} else {
		goto L315
	}
L315:
	;
	F_UnlockBuffer(m, v754)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L33
	} else {
		goto L316
	}
L316:
	;
	goto L295
L317:
	;
	goto L294
L318:
	;
	v1026 = int32(1)
	goto L293
L319:
	;
	v1032 = int32(4)
	v1033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+14)))
	v1034 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992)+12)))
	v1035 = v1033 - v1034
	if v1035 <= v1032 {
		goto L321
	} else {
		goto L322
	}
L320:
	;
	if base.Ui32(v32) <= base.Ui32(v1095) {
		goto L44
	} else {
		goto L339
	}
L321:
	;
	v1038 = v1032
	goto L323
L322:
	;
	v1038 = v1035
	goto L323
L323:
	;
	v1040 = v1038 - int32(4)
	if v1040 == int32(0) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	v1095 = int32(0)
	goto L320
L325:
	;
	goto L326
L326:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1034) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1095 = v1040
	goto L320
L328:
	;
	v1051 = int32(base.Ui32(v1034+int32(_a_F_RelationGetBufferForTuple_3)) >> (uint(int32(2)) % 32))
	goto L330
L329:
	;
	v1051 = int32(0)
	goto L330
L330:
	;
	if base.Ui32(v1051&int32(_a_F_RelationGetBufferForTuple_8)) < base.Ui32(int32(291)) {
		goto L327
	} else {
		goto L331
	}
L331:
	;
	v1056 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+10)))
	if v1056&int32(1) == int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v1095 = int32(0)
	goto L320
L333:
	;
	goto L334
L334:
	;
	v1065 = int32(1)
	goto L335
L335:
	;
	v1074 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992+int32(20)+v1065&int32(_a_F_RelationGetBufferForTuple_8)<<(uint(int32(2))%32))+1)))
	if v1074&int32(384) == int32(0) {
		goto L327
	} else {
		goto L337
	}
L336:
	;
	v1095 = int32(0)
	goto L320
L337:
	;
	v1080 = v1065 + int32(1)
	v1081 = int32(_a_F_RelationGetBufferForTuple_8)
	if base.Ui32(v1080&v1081) <= base.Ui32(v1051&v1081) {
		v1065 = v1080
		goto L335
	} else {
		goto L338
	}
L338:
	;
	goto L336
L339:
	;
	if v1027|v1026 != int32(1) {
		goto L1
	} else {
		goto L340
	}
L340:
	;
	if l2 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	F_UnlockBuffer(m, l2)
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L33
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	F_UnlockReleaseBuffer(m, v754)
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L33
	} else {
		goto L345
	}
L344:
	;
	goto L343
L345:
	;
	v144 = v974
	goto L45
L346:
	;
	if base.Ui32(v1170) < base.Ui32(v32) {
		goto L1
	} else {
		goto L365
	}
L347:
	;
	v1113 = v1107
	goto L349
L348:
	;
	v1113 = v1110
	goto L349
L349:
	;
	v1115 = v1113 - int32(4)
	if v1115 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1170 = int32(0)
	goto L346
L351:
	;
	goto L352
L352:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1109) {
		goto L354
	} else {
		goto L355
	}
L353:
	;
	v1170 = v1115
	goto L346
L354:
	;
	v1126 = int32(base.Ui32(v1109+int32(_a_F_RelationGetBufferForTuple_3)) >> (uint(int32(2)) % 32))
	goto L356
L355:
	;
	v1126 = int32(0)
	goto L356
L356:
	;
	if base.Ui32(v1126&int32(_a_F_RelationGetBufferForTuple_8)) < base.Ui32(int32(291)) {
		goto L353
	} else {
		goto L357
	}
L357:
	;
	v1131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v992)+10)))
	if v1131&int32(1) == int32(0) {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1170 = int32(0)
	goto L346
L359:
	;
	goto L360
L360:
	;
	v1140 = int32(1)
	goto L361
L361:
	;
	v1149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v992+int32(20)+v1140&int32(_a_F_RelationGetBufferForTuple_8)<<(uint(int32(2))%32))+1)))
	if v1149&int32(384) == int32(0) {
		goto L353
	} else {
		goto L363
	}
L362:
	;
	v1170 = int32(0)
	goto L346
L363:
	;
	v1155 = v1140 + int32(1)
	v1156 = int32(_a_F_RelationGetBufferForTuple_8)
	if base.Ui32(v1155&v1156) <= base.Ui32(v1126&v1156) {
		v1140 = v1155
		goto L361
	} else {
		goto L364
	}
L364:
	;
	goto L362
L365:
	;
	goto L44
L366:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1175
	v1177 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v1177
	v1181 = F_smgropen(m, v27+int32(32), v1174)
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L33
	} else {
		goto L367
	}
L367:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1181
	v1185 = v974
	v1186 = v1181
	v1191 = v754
	goto L43
L368:
	;
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1223 = v1185
	v1224 = v1221
	v1229 = v1191
	goto L42
L369:
	;
	v1217 = v1209
	goto L371
L370:
	;
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+76))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1210)+4)) = v1211
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1211))) = v1213
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+72))
	v1217 = v1215
	goto L371
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1186)+72)) = v1217 + int32(1)
	goto L368
L372:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L33
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(_a_F_RelationGetBufferForTuple_19)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v32
	F_errmsg(m, int32(_a_F_RelationGetBufferForTuple_20), v27)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L33
	} else {
		goto L374
	}
L374:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(534), int32(_a_F_RelationGetBufferForTuple_21))
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L33
	} else {
		goto L375
	}
L375:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v32
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_22), v27+int32(16))
	mBase = m.M
	v1279 = m.ExcPending
	if v1279 != 0 {
		goto L33
	} else {
		goto L377
	}
L377:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(869), int32(_a_F_RelationGetBufferForTuple_21))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L33
	} else {
		goto L378
	}
L378:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v207
L2:
	;
	v13 = F_copyObjectImpl(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v207 = v13
	goto L1
L7:
	;
	v207 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v21 = int32(0)
	v24 = F_heap_attisnull(m, v17, int32(20), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v24 != 0 {
		v207 = v21
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[0]))
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(_a_F_RelationGetIndexExpressions_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1]))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v35
	v38 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v160 = v28
	goto L14
L14:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+18)))
	if base.Ui32(v167&int32(2044)) <= base.Ui32(int32(19)) {
		goto L40
	} else {
		goto L41
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+4)) = int64(-4294965047)
	v44 = v21
	goto L16
L16:
	;
	v49 = int32(100)
	v50 = v44 * v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	base.MemoryCopy(m, v50+(v38+v51<<(uint(int32(3))%32))+int32(28), v50+int32(_a_F_RelationGetIndexExpressions_1), v49)
	F_populate_compact_attribute(m, v38, v44)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v68 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v68 < v77 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v65 = v44 + int32(1)
	if v65 != int32(21) {
		v44 = v65
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[0])) = v38
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v32
	v160 = v38
	goto L14
L21:
	;
	v81 = v38 + int32(28)
	v88 = v68
	v89 = v77
	v91 = v68
	goto L25
L22:
	;
	v145 = v68
	v152 = v77
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v145
	goto L20
L24:
	;
	v145 = v139
	v152 = v118
	goto L23
L25:
	;
	v97 = v81 + v77<<(uint(int32(3))%32) + v88*int32(100)
	v100 = v81 + v88<<(uint(int32(3))%32)
	if v77 != v89 {
		v118 = v89
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v139 = v77
	goto L24
L27:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v119 <= int32(0) {
		v139 = v88
		goto L24
	} else {
		goto L35
	}
L28:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
	if v102 != int32(118) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v88
	goto L27
L30:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	if v105 != int32(1) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	if v108&int32(6) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v111 <= int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v114 != int32(118) {
		v118 = v77
		goto L27
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v122 == int32(118) {
		v139 = v88
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
	v131 = (v91 + v125 - int32(1)) & (int32(0) - v125)
	if int32(_a_F_RelationGetIndexExpressions_2) < v131 {
		v139 = v88
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v100))) = uint16(v131)
	v137 = v88 + int32(1)
	if v137 != v77 {
		v88 = v137
		v89 = v118
		v91 = v131 + v119
		goto L25
	} else {
		goto L38
	}
L38:
	;
	goto L26
L39:
	;
	v184 = F_text_to_cstring(m, base.I32_wrap_i64(v182))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L45
	}
L40:
	;
	v175 = F_getmissingattr(m, v160, int32(20), v10+int32(15))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v180 = F_fastgetattr_3(m, v26, int32(20), v160, v10+int32(15))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v182 = v175
	goto L39
L44:
	;
	v182 = v180
	goto L39
L45:
	;
	v186 = F_stringToNode(m, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v191 = F_eval_const_expressions(m, int32(0), v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	F_fix_opfuncids(m, v191)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v195 = int32(_a_F_RelationGetIndexExpressions_0)
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1]))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v198
	v200 = F_copyObjectImpl(m, v191)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+228)) = v200
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v196
	v207 = v191
	goto L1
}
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int64
	_ = v175
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v212
L2:
	;
	v13 = F_copyObjectImpl(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v212 = v13
	goto L1
L7:
	;
	v212 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v21 = int32(0)
	v24 = F_heap_attisnull(m, v17, int32(21), v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v24 != 0 {
		v212 = v21
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[0]))
	if v28 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(_a_F_RelationGetIndexPredicate_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1]))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v35
	v38 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v160 = v28
	goto L14
L14:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+18)))
	if base.Ui32(v167&int32(2047)) <= base.Ui32(int32(20)) {
		goto L40
	} else {
		goto L41
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+4)) = int64(-4294965047)
	v44 = v21
	goto L16
L16:
	;
	v49 = int32(100)
	v50 = v44 * v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	base.MemoryCopy(m, v50+(v38+v51<<(uint(int32(3))%32))+int32(28), v50+int32(_a_F_RelationGetIndexPredicate_1), v49)
	F_populate_compact_attribute(m, v38, v44)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L18
	}
L17:
	;
	v68 = int32(0)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v68 < v77 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v65 = v44 + int32(1)
	if v65 != int32(21) {
		v44 = v65
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[0])) = v38
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v32
	v160 = v38
	goto L14
L21:
	;
	v81 = v38 + int32(28)
	v88 = v68
	v89 = v77
	v91 = v68
	goto L25
L22:
	;
	v145 = v68
	v152 = v77
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v145
	goto L20
L24:
	;
	v145 = v139
	v152 = v118
	goto L23
L25:
	;
	v97 = v81 + v77<<(uint(int32(3))%32) + v88*int32(100)
	v100 = v81 + v88<<(uint(int32(3))%32)
	if v77 != v89 {
		v118 = v89
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v139 = v77
	goto L24
L27:
	;
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v119 <= int32(0) {
		v139 = v88
		goto L24
	} else {
		goto L35
	}
L28:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+7)))
	if v102 != int32(118) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v118 = v88
	goto L27
L30:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+4)))
	if v105 != int32(1) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+6)))
	if v108&int32(6) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v111 = int32(*(*int16)(unsafe.Add(mBase, uint32(v100)+2)))
	if v111 <= int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v114 != int32(118) {
		v118 = v77
		goto L27
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+90)))
	if v122 == int32(118) {
		v139 = v88
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+5)))
	v131 = (v91 + v125 - int32(1)) & (int32(0) - v125)
	if int32(_a_F_RelationGetIndexPredicate_2) < v131 {
		v139 = v88
		goto L24
	} else {
		goto L37
	}
L37:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v100))) = uint16(v131)
	v137 = v88 + int32(1)
	if v137 != v77 {
		v88 = v137
		v89 = v118
		v91 = v131 + v119
		goto L25
	} else {
		goto L38
	}
L38:
	;
	goto L26
L39:
	;
	v184 = F_text_to_cstring(m, base.I32_wrap_i64(v182))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L45
	}
L40:
	;
	v175 = F_getmissingattr(m, v160, int32(21), v10+int32(15))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v180 = F_fastgetattr_3(m, v26, int32(21), v160, v10+int32(15))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v182 = v175
	goto L39
L44:
	;
	v182 = v180
	goto L39
L45:
	;
	v186 = F_stringToNode(m, v184)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_pfree(m, v184)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v191 = F_eval_const_expressions(m, int32(0), v186)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v194 = F_canonicalize_qual(m, v191, int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v196 = F_make_ands_implicit(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	F_fix_opfuncids(m, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v200 = int32(_a_F_RelationGetIndexPredicate_0)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1]))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v203
	v205 = F_copyObjectImpl(m, v196)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+232)) = v205
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v201
	v212 = v196
	goto L1
}
func F_RelationGetPartitionDesc(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v215 int64
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v238 int64
	_ = v238
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v241 int64
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int64
	_ = v266
	var v267 int32
	_ = v267
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v638 int32
	_ = v638
	var v652 int32
	_ = v652
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v840 int32
	_ = v840
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v906 int32
	_ = v906
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1014 int32
	_ = v1014
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1086 int32
	_ = v1086
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1176 int32
	_ = v1176
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1235 int32
	_ = v1235
	var v1237 int64
	_ = v1237
	var v1241 int32
	_ = v1241
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1273 int32
	_ = v1273
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1301 int32
	_ = v1301
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1354 int32
	_ = v1354
	var v1365 int32
	_ = v1365
	var v1369 int32
	_ = v1369
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1397 int32
	_ = v1397
	var v1406 int32
	_ = v1406
	var v1414 int32
	_ = v1414
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1434 int32
	_ = v1434
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int64
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1482 int64
	_ = v1482
	var v1483 int64
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1519 int32
	_ = v1519
	var v1535 int32
	_ = v1535
	var v1551 int32
	_ = v1551
	var v1565 int32
	_ = v1565
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1585 int32
	_ = v1585
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1597 int32
	_ = v1597
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1612 int32
	_ = v1612
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1646 int32
	_ = v1646
	var v1651 int32
	_ = v1651
	var v1659 int32
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int64
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1708 int32
	_ = v1708
	var v1709 int64
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1728 int32
	_ = v1728
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1783 int32
	_ = v1783
	var v1787 int32
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1818 int32
	_ = v1818
	var v1821 int32
	_ = v1821
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1835 int32
	_ = v1835
	var v1839 int32
	_ = v1839
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1862 int32
	_ = v1862
	var v1866 int32
	_ = v1866
	var v1870 int32
	_ = v1870
	var v1875 int32
	_ = v1875
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1894 int32
	_ = v1894
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1901 int32
	_ = v1901
	var v1919 int32
	_ = v1919
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1942 int32
	_ = v1942
	var v1953 int32
	_ = v1953
	var v1969 int32
	_ = v1969
	var v1970 int32
	_ = v1970
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1979 int64
	_ = v1979
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int64
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2024 int32
	_ = v2024
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2099 int32
	_ = v2099
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2149 int32
	_ = v2149
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2178 int32
	_ = v2178
	var v2228 int32
	_ = v2228
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2242 int32
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2247 int32
	_ = v2247
	var v2248 int32
	_ = v2248
	var v2262 int32
	_ = v2262
	var v2266 int32
	_ = v2266
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2290 int32
	_ = v2290
	var v2294 int32
	_ = v2294
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2307 int32
	_ = v2307
	var v2309 int32
	_ = v2309
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2326 int32
	_ = v2326
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2338 int32
	_ = v2338
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2347 int32
	_ = v2347
	var v2350 int32
	_ = v2350
	var v2351 int32
	_ = v2351
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2360 int32
	_ = v2360
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2386 int32
	_ = v2386
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2401 int32
	_ = v2401
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2409 int32
	_ = v2409
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2469 int32
	_ = v2469
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2496 int32
	_ = v2496
	var v2520 int32
	_ = v2520
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2532 int32
	_ = v2532
	var v2534 int32
	_ = v2534
	var v2535 int32
	_ = v2535
	var v2537 int32
	_ = v2537
	var v2539 int64
	_ = v2539
	var v2542 int32
	_ = v2542
	var v2546 int32
	_ = v2546
	var v2547 int32
	_ = v2547
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2554 int64
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2556 int32
	_ = v2556
	var v2558 int32
	_ = v2558
	var v2565 int32
	_ = v2565
	var v2595 int32
	_ = v2595
	var v2625 int32
	_ = v2625
	var v2626 int32
	_ = v2626
	var v2629 int32
	_ = v2629
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2634 int32
	_ = v2634
	var v2643 int32
	_ = v2643
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2667 int32
	_ = v2667
	var v2673 int32
	_ = v2673
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2691 int32
	_ = v2691
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2701 int32
	_ = v2701
	var v2703 int32
	_ = v2703
	var v2705 int32
	_ = v2705
	var v2707 int32
	_ = v2707
	var v2708 int32
	_ = v2708
	var v2710 int32
	_ = v2710
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2720 int32
	_ = v2720
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2743 int32
	_ = v2743
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2766 int32
	_ = v2766
	var v2768 int32
	_ = v2768
	var v2771 int32
	_ = v2771
	var v2802 int32
	_ = v2802
	var v2830 int32
	_ = v2830
	var v2837 int32
	_ = v2837
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2848 int32
	_ = v2848
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2858 int32
	_ = v2858
	var v2865 int32
	_ = v2865
	var v2879 int32
	_ = v2879
	var v2883 int32
	_ = v2883
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2893 int32
	_ = v2893
	var v2900 int32
	_ = v2900
	var v2914 int32
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2921 int32
	_ = v2921
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2931 int32
	_ = v2931
	var v2938 int32
	_ = v2938
	var v2950 int32
	_ = v2950
	var v2954 int32
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2964 int32
	_ = v2964
	var v2971 int32
	_ = v2971
	var v2988 int32
	_ = v2988
	var v3019 int32
	_ = v3019
	var v3023 int32
	_ = v3023
	var v3028 int32
	_ = v3028
	var v3032 int32
	_ = v3032
	var v3038 int32
	_ = v3038
	var v3043 int32
	_ = v3043
	var v3047 int32
	_ = v3047
	var v3054 int32
	_ = v3054
	var v3059 int32
	_ = v3059
	v3 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(128)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v32 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3047 = m.ExcPending
	if v3047 != 0 {
		goto L17
	} else {
		goto L469
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3032 = m.ExcPending
	if v3032 != 0 {
		goto L17
	} else {
		goto L466
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3019 = m.ExcPending
	if v3019 != 0 {
		goto L17
	} else {
		goto L463
	}
L4:
	;
	m.G0 = v30 + int32(128)
	return v2988
L5:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	if l1 == int32(0) {
		v2988 = v32
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+4)))
	if v37&int32(1) == int32(0) {
		v2988 = v32
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L9
L9:
	;
	if v43 != int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2988 = v46
	goto L4
L11:
	;
	v68 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L17
	} else {
		goto L20
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v49 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L14
L14:
	;
	if base.B2i32(v53 != int32(0)) == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	goto L16
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v62 = F_XidInMVCCSnapshot(m, v61, v60)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v62 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v2988 = v66
	goto L4
L20:
	;
	v70 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)) = uint8(v70)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v70
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v80 = F_find_inheritance_children_extended(m, v74, l1, v70, v30+int32(127), v30+int32(120))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L24
	}
L21:
	;
	v2309 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[1]))
	v2314 = F_AllocSetContextCreateInternal(m, v2309, int32(_a_F_RelationGetPartitionDesc_0), int32(0), int32(1024), int32(_a_F_RelationGetPartitionDesc_1))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L17
	} else {
		goto L321
	}
L22:
	;
	v2290 = v2262
	v2294 = v2266
	v2303 = v2275
	v2304 = v2276
	v2307 = int32(0)
	goto L21
L23:
	;
	v92 = v3
	v93 = v80 + int32(4)
	v98 = v82
	v102 = v3
	v104 = v80
	goto L29
L24:
	;
	if v80 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if int32(0) < v82 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	v85 = v3
	goto L27
L27:
	;
	v2262 = v85
	v2266 = v3
	v2275 = v3
	v2276 = v3
	goto L22
L28:
	;
	v85 = v82
	goto L27
L29:
	;
	v116 = v98 << (uint(int32(2)) % 32)
	v117 = F_palloc(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v2262 = int32(0)
	v2266 = v2231
	v2275 = v117
	v2276 = v119
	goto L22
L31:
	;
	v119 = F_palloc(m, v98)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v121 = F_palloc(m, v116)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if int32(0) < v123 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L17
	} else {
		goto L315
	}
L35:
	;
	v140 = v102
	goto L38
L36:
	;
	goto L37
L37:
	;
	v352 = int32(0)
	v357 = v30 - int32(-64)
	v359 = F_palloc_mul(m, int32(4), v98)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L17
	} else {
		goto L99
	}
L38:
	;
	v155 = v140 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155+v156)))
	v159 = base.I64_extend_i32_u(v158)
	v160 = F_SearchSysCache1(m, int32(57), v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L17
	} else {
		goto L42
	}
L39:
	;
	goto L37
L40:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v299 != int32(98) {
		goto L2
	} else {
		goto L90
	}
L41:
	;
	v185 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L17
	} else {
		goto L53
	}
L42:
	;
	if v160 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v168 = F_SysCacheGetAttr(m, int32(57), v160, int32(34), v30-int32(-64))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+64)))
	if v170 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ReleaseCatCache(m, v160)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v176 = F_text_to_cstring(m, base.I32_wrap_i64(v168))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L17
	} else {
		goto L49
	}
L48:
	;
	goto L41
L49:
	;
	v178 = F_stringToNode(m, v176)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	F_ReleaseCatCache(m, v160)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	if v178 != 0 {
		v294 = v178
		goto L40
	} else {
		goto L52
	}
L52:
	;
	goto L41
L53:
	;
	v188 = v30 - int32(-64)
	F_ScanKeyInit(m, v188, int32(1), int32(3), int32(184), v159)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v194 = int32(0)
	v196 = int32(1)
	v199 = F_systable_beginscan(m, v185, int32(2662), v196, v194, v196, v188)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	F_systable_endscan(m, v199)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L17
	} else {
		goto L86
	}
L56:
	;
	v201 = F_systable_getnext(m, v199)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	if v201 == int32(0) {
		v279 = v194
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v185)+52))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+18)))
	if base.Ui32(v207&int32(2046)) <= base.Ui32(int32(33)) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+63)))
	if v272 != 0 {
		v279 = int32(0)
		goto L55
	} else {
		goto L83
	}
L60:
	;
	v215 = F_getmissingattr(m, v205, int32(34), v30+int32(63))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L17
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+63)) = uint8(v217)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+20)))
	if v219&int32(1) == v217 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v270 = v215
	goto L59
L64:
	;
	v266 = F_nocachegetattr(m, v201, int32(34), v205)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L17
	} else {
		goto L82
	}
L65:
	;
	v224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v205)+292)))
	if v224 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+27)))
	if v258&int32(2) != 0 {
		goto L64
	} else {
		goto L81
	}
L68:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+22)))
	v229 = v206 + v227 + v224
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+296)))
	if v230 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v233 = int32(*(*int16)(unsafe.Add(mBase, uint32(v205)+294)))
	if base.I32_popcnt(v233) != int32(1) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v270 = base.I64_extend_i32_u(v229)
	goto L59
L72:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L17
	} else {
		goto L78
	}
L73:
	;
	switch base.I32_ctz(v233) {
	case 0:
		goto L77
	case 1:
		goto L76
	case 2:
		goto L75
	case 3:
		goto L74
	default:
		goto L72
	}
L74:
	;
	v241 = *(*int64)(unsafe.Add(mBase, uint32(v229)))
	v270 = v241
	goto L59
L75:
	;
	v240 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v229))))
	v270 = v240
	goto L59
L76:
	;
	v239 = int64(*(*int16)(unsafe.Add(mBase, uint32(v229))))
	v270 = v239
	goto L59
L77:
	;
	v238 = int64(*(*int8)(unsafe.Add(mBase, uint32(v229))))
	v270 = v238
	goto L59
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+48)) = v233
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_2), v30+int32(48))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L17
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_3), int32(123), int32(_a_F_RelationGetPartitionDesc_4))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L17
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
	v261 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+63)) = uint8(v261)
	v270 = int64(0)
	goto L59
L82:
	;
	v270 = v266
	goto L59
L83:
	;
	v274 = F_text_to_cstring(m, base.I32_wrap_i64(v270))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L84
	}
L84:
	;
	v276 = F_stringToNode(m, v274)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L17
	} else {
		goto L85
	}
L85:
	;
	v279 = v276
	goto L55
L86:
	;
	F_relation_close(m, v185, int32(1))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L17
	} else {
		goto L87
	}
L87:
	;
	if base.B2i32(v279 == int32(0))&(v92^int32(-1)) != 0 {
		goto L34
	} else {
		goto L88
	}
L88:
	;
	if v279 == int32(0) {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v294 = v279
	goto L40
L90:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+5)))
	if v302 == int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v306 = F_get_default_partition_oid(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L17
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155+v117))) = v158
	v313 = F_get_rel_relkind(m, v158)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L17
	} else {
		goto L96
	}
L94:
	;
	if v306 != v158 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v140+v119))) = uint8(base.B2i32(v313 != int32(112)))
	*(*int32)(unsafe.Add(mBase, uint32(v155+v121))) = v294
	v321 = v140 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v321 < v322 {
		v140 = v321
		goto L38
	} else {
		goto L97
	}
L97:
	;
	goto L39
L98:
	;
	v2290 = v98
	v2294 = int32(1)
	v2303 = v117
	v2304 = v119
	v2307 = v2228
	goto L21
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v357))) = v359
	if v98 <= int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v514 = int32(0)
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	switch v515 - int32(104) {
	case 0:
		goto L121
	default:
		v2178 = v514
		goto L112
	case 4:
		goto L120
	case 10:
		goto L119
	}
L101:
	;
	v365 = v98 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v98) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v372 = v352
	v374 = v352
	goto L105
L103:
	;
	v424 = v352
	goto L104
L104:
	;
	v451 = v424
	v452 = v352
	goto L109
L105:
	;
	v398 = v372 << (uint(int32(2)) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v401 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v398+v399))) = v401
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v403+v398)+4)) = v401
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v407+v398)+8)) = v401
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v411+v398)+12)) = v401
	v415 = int32(4)
	v416 = v372 + v415
	v418 = v374 + v415
	if v418 != v98&int32(2147483644) {
		v372 = v416
		v374 = v418
		goto L105
	} else {
		goto L107
	}
L106:
	;
	if v365 == int32(0) {
		goto L100
	} else {
		goto L108
	}
L107:
	;
	goto L106
L108:
	;
	v424 = v416
	goto L104
L109:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v476+v451<<(uint(int32(2))%32)))) = int32(-1)
	v482 = int32(1)
	v485 = v452 + v482
	if v485 != v365 {
		v451 = v451 + v482
		v452 = v485
		goto L109
	} else {
		goto L111
	}
L110:
	;
	goto L100
L111:
	;
	goto L110
L112:
	;
	v2228 = v2178
	goto L98
L113:
	;
	F_qsort_arg(m, v1901, v1897, int32(16), int32(959), v68)
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L17
	} else {
		goto L276
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1879 = m.ExcPending
	if v1879 != 0 {
		goto L17
	} else {
		goto L273
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L17
	} else {
		goto L270
	}
L116:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L17
	} else {
		goto L267
	}
L117:
	;
	v1847 = F_palloc(m, int32(0))
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L17
	} else {
		goto L266
	}
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1835 = m.ExcPending
	if v1835 != 0 {
		goto L17
	} else {
		goto L263
	}
L119:
	;
	v1284 = F_palloc0(m, int32(36))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L17
	} else {
		goto L198
	}
L120:
	;
	v938 = F_palloc0(m, int32(36))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L17
	} else {
		goto L160
	}
L121:
	;
	v519 = F_palloc0(m, int32(36))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L17
	} else {
		goto L122
	}
L122:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v519)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v519))) = v521
	v526 = F_palloc_mul(m, int32(12), v98)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L17
	} else {
		goto L123
	}
L123:
	;
	if int32(0) < v98 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v533 = int32(0)
	goto L127
L125:
	;
	goto L126
L126:
	;
	F_pg_qsort(m, v526, v98, int32(12), int32(957))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L17
	} else {
		goto L131
	}
L127:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v121+v533<<(uint(int32(2))%32))))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v561)+4)))
	if v562 != int32(104) {
		goto L118
	} else {
		goto L129
	}
L128:
	;
	goto L126
L129:
	;
	v567 = v526 + v533*int32(12)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v567))) = v568
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v561)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v567)+8)) = v533
	*(*int32)(unsafe.Add(mBase, uint32(v567)+4)) = v570
	v574 = v533 + int32(1)
	if v574 != v98 {
		v533 = v574
		goto L127
	} else {
		goto L130
	}
L130:
	;
	goto L128
L131:
	;
	v607 = int32(12)
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v526+v98*v607-v607)))
	*(*int32)(unsafe.Add(mBase, uint32(v519)+4)) = v98
	v615 = F_palloc0_mul(m, int32(4), v98)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L17
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+20)) = v612
	*(*int64)(unsafe.Add(mBase, uint32(v519)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v519)+8)) = v615
	v623 = F_palloc(m, v612<<(uint(int32(2))%32))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L17
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v519)+24)) = v623
	if v612 <= int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v782 = F_palloc(m, v98<<(uint(int32(4))%32))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L17
	} else {
		goto L146
	}
L135:
	;
	v629 = v612 & int32(3)
	v630 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v612) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v638 = v630
	v652 = int32(0)
	goto L139
L137:
	;
	v690 = v630
	goto L138
L138:
	;
	v717 = v690
	v719 = v514
	goto L143
L139:
	;
	v664 = v638 << (uint(int32(2)) % 32)
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	v667 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v664+v665))) = v667
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v669+v664)+4)) = v667
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v673+v664)+8)) = v667
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v677+v664)+12)) = v667
	v681 = int32(4)
	v682 = v638 + v681
	v684 = v652 + v681
	if v684 != v612&int32(2147483644) {
		v638 = v682
		v652 = v684
		goto L139
	} else {
		goto L141
	}
L140:
	;
	if v629 == int32(0) {
		goto L134
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v690 = v682
	goto L138
L143:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v742+v717<<(uint(int32(2))%32)))) = int32(-1)
	v748 = int32(1)
	v751 = v719 + v748
	if v751 != v629 {
		v717 = v717 + v748
		v719 = v751
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L134
L145:
	;
	goto L144
L146:
	;
	if int32(0) < v98 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v790 = int32(0)
	goto L150
L148:
	;
	goto L149
L149:
	;
	F_pfree(m, v526)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L17
	} else {
		goto L159
	}
L150:
	;
	v816 = v526 + v790*int32(12)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)+4))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	v820 = v790 << (uint(int32(2)) % 32)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v820+v821))) = v782 + v790<<(uint(int32(4))%32)
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v827+v820)))
	*(*int64)(unsafe.Add(mBase, uint32(v829))) = base.I64_extend_i32_s(v818)
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v519)+8))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v832+v820)))
	*(*int64)(unsafe.Add(mBase, uint32(v834)+8)) = base.I64_extend_i32_s(v817)
	if v817 < v612 {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	goto L149
L152:
	;
	v840 = v817
	goto L155
L153:
	;
	goto L154
L154:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v816)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v899+v900<<(uint(int32(2))%32)))) = v790
	v906 = v790 + int32(1)
	if v906 != v98 {
		v790 = v906
		goto L150
	} else {
		goto L158
	}
L155:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v519)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v865+v840<<(uint(int32(2))%32)))) = v790
	v870 = v840 + v818
	if v870 < v612 {
		v840 = v870
		goto L155
	} else {
		goto L157
	}
L156:
	;
	goto L154
L157:
	;
	goto L156
L158:
	;
	goto L151
L159:
	;
	v2228 = v519
	goto L98
L160:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v938)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v938))) = v940
	if v98 <= int32(0) {
		goto L117
	} else {
		goto L161
	}
L161:
	;
	v954 = v3
	v958 = v352
	goto L162
L162:
	;
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v121+v958<<(uint(int32(2))%32))))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)+16))
	if v977 == int32(0) {
		v1126 = v954
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1151 = F_palloc(m, v1126<<(uint(int32(4))%32))
	mBase = m.M
	v1152 = m.ExcPending
	if v1152 != 0 {
		goto L17
	} else {
		goto L179
	}
L164:
	;
	v1146 = v958 + int32(1)
	if v1146 != v98 {
		v954 = v1126
		v958 = v1146
		goto L162
	} else {
		goto L178
	}
L165:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977)+4))
	if v980 <= int32(0) {
		v1126 = v954
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v984 = v980 & int32(3)
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v977)+12))
	if base.Ui32(v980) < base.Ui32(int32(4)) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v1080 = v1053
	v1081 = int32(0)
	v1086 = v1059
	goto L175
L168:
	;
	v1053 = int32(0)
	v1059 = v954
	goto L167
L169:
	;
	goto L170
L170:
	;
	v992 = int32(0)
	v996 = v992
	v1002 = v954
	v1014 = v992
	goto L171
L171:
	;
	v1023 = v985 + v996<<(uint(int32(2))%32)
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1023)))
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024)+32)))
	v1026 = int32(1)
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+4))
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1029)+32)))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+8))
	v1035 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+32)))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1023)+12))
	v1040 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039)+32)))
	v1043 = v1002 + (v1025 ^ v1026) + (v1030 ^ v1026) + (v1035 ^ v1026) + (v1040 ^ v1026)
	v1044 = int32(4)
	v1045 = v996 + v1044
	v1047 = v1014 + v1044
	if v1047 != v980&int32(2147483644) {
		v996 = v1045
		v1002 = v1043
		v1014 = v1047
		goto L171
	} else {
		goto L173
	}
L172:
	;
	if v984 == int32(0) {
		v1126 = v1043
		goto L164
	} else {
		goto L174
	}
L173:
	;
	goto L172
L174:
	;
	v1053 = v1045
	v1059 = v1043
	goto L167
L175:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v985+v1080<<(uint(int32(2))%32))))
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1108)+32)))
	v1110 = int32(1)
	v1112 = v1086 + (v1109 ^ v1110)
	v1116 = v1081 + v1110
	if v1116 != v984 {
		v1080 = v1080 + v1110
		v1081 = v1116
		v1086 = v1112
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v1126 = v1112
	goto L164
L177:
	;
	goto L176
L178:
	;
	goto L163
L179:
	;
	v1154 = int32(-1)
	v1159 = int32(0)
	v1161 = v1154
	v1163 = v1154
	v1176 = int32(0)
	goto L180
L180:
	;
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v121+v1159<<(uint(int32(2))%32))))
	v1187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186)+4)))
	if v1187 != int32(108) {
		goto L116
	} else {
		goto L182
	}
L181:
	;
	v1894 = v1258
	v1896 = v1260
	v1897 = v1126
	v1901 = v1151
	goto L113
L182:
	;
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1186)+5)))
	if v1190 != 0 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	v1281 = v1159 + int32(1)
	if v1281 != v98 {
		v1159 = v1281
		v1161 = v1258
		v1163 = v1260
		v1176 = v1273
		goto L180
	} else {
		goto L197
	}
L184:
	;
	v1258 = v1159
	v1260 = v1163
	v1273 = v1176
	goto L183
L185:
	;
	goto L186
L186:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1186)+16))
	if v1191 == int32(0) {
		v1258 = v1161
		v1260 = v1163
		v1273 = v1176
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v1194 = int32(0)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	if v1195 <= v1194 {
		v1258 = v1161
		v1260 = v1163
		v1273 = v1176
		goto L183
	} else {
		goto L188
	}
L188:
	;
	v1200 = v1194
	v1205 = v1163
	v1213 = v1195
	v1218 = v1176
	goto L189
L189:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+12))
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v1225+v1200<<(uint(int32(2))%32))))
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229)+32)))
	if v1230 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L190:
	;
	v1258 = v1161
	v1260 = v1246
	v1273 = v1249
	goto L183
L191:
	;
	v1251 = v1200 + int32(1)
	if v1251 < v1248 {
		v1200 = v1251
		v1205 = v1246
		v1213 = v1248
		v1218 = v1249
		goto L189
	} else {
		goto L196
	}
L192:
	;
	v1235 = v1151 + v1218<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1235))) = v1159
	v1237 = *(*int64)(unsafe.Add(mBase, uint32(v1229)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v1235)+8)) = v1237
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1191)+4))
	v1246 = v1205
	v1248 = v1241
	v1249 = v1218 + int32(1)
	goto L191
L193:
	;
	goto L194
L194:
	;
	if base.B2i32(v1205 == int32(-1)) == int32(0) {
		goto L115
	} else {
		goto L195
	}
L195:
	;
	v1246 = v1159
	v1248 = v1213
	v1249 = v1218
	goto L191
L196:
	;
	goto L190
L197:
	;
	goto L181
L198:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int64)(unsafe.Add(mBase, uint32(v1284)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1284))) = v1286
	v1290 = int32(-1)
	v1294 = F_palloc0_mul(m, int32(4), v98<<(uint(int32(1))%32))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L17
	} else {
		goto L199
	}
L199:
	;
	if int32(0) < v98 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1301 = int32(0)
	v1308 = v1290
	v1312 = v3
	goto L203
L201:
	;
	v1365 = v1290
	v1369 = v3
	goto L202
L202:
	;
	F_qsort_arg(m, v1294, v1369, int32(4), int32(958), v68)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L17
	} else {
		goto L213
	}
L203:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v121+v1301<<(uint(int32(2))%32))))
	v1330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+4)))
	if v1330 != int32(114) {
		goto L114
	} else {
		goto L205
	}
L204:
	;
	v1365 = v1351
	v1369 = v1352
	goto L202
L205:
	;
	v1333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1329)+5)))
	if v1333 != 0 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v1354 = v1301 + int32(1)
	if v1354 != v98 {
		v1301 = v1354
		v1308 = v1351
		v1312 = v1352
		goto L203
	} else {
		goto L212
	}
L207:
	;
	v1351 = v1301
	v1352 = v1312
	goto L206
L208:
	;
	goto L209
L209:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+20))
	v1336 = F_make_one_partition_rbound(m, v68, v1301, v1334, int32(1))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L17
	} else {
		goto L210
	}
L210:
	;
	v1340 = v1294 + v1312<<(uint(int32(2))%32)
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v1329)+24))
	v1343 = F_make_one_partition_rbound(m, v68, v1301, v1341, int32(0))
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L17
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1340)+4)) = v1343
	*(*int32)(unsafe.Add(mBase, uint32(v1340))) = v1336
	v1351 = v1308
	v1352 = v1312 + int32(2)
	goto L206
L212:
	;
	goto L204
L213:
	;
	v1389 = F_palloc(m, v1369<<(uint(int32(2))%32))
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L17
	} else {
		goto L214
	}
L214:
	;
	if int32(0) < v1369 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1397 = int32(0)
	v1406 = v352
	v1414 = v3
	goto L218
L216:
	;
	v1565 = v352
	goto L217
L217:
	;
	F_pfree(m, v1294)
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L17
	} else {
		goto L234
	}
L218:
	;
	v1423 = v1294 + v1414<<(uint(int32(2))%32)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	v1425 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	if v1425 <= int32(0) {
		v1535 = v1406
		goto L220
	} else {
		goto L221
	}
L219:
	;
	v1565 = v1535
	goto L217
L220:
	;
	v1551 = v1414 + int32(1)
	if v1551 != v1369 {
		v1397 = v1424
		v1406 = v1535
		v1414 = v1551
		goto L218
	} else {
		goto L233
	}
L221:
	;
	if v1397 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v1434 = int32(0)
	goto L225
L223:
	;
	v1519 = v1424
	goto L224
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1389+v1406<<(uint(int32(2))%32)))) = v1519
	v1535 = v1406 + int32(1)
	goto L220
L225:
	;
	v1460 = v1434 << (uint(int32(2)) % 32)
	v1461 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+8))
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1460+v1461)))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+8))
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1464+v1460)))
	if v1463 != v1466 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1423)))
	v1519 = v1491
	goto L224
L227:
	;
	goto L226
L228:
	;
	if v1463 != 0 {
		v1535 = v1406
		goto L220
	} else {
		goto L229
	}
L229:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v68)+24))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v68)+28))
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1472+v1460)))
	v1476 = v1434 << (uint(int32(3)) % 32)
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1424)+4))
	v1479 = *(*int64)(unsafe.Add(mBase, uint32(v1476+v1477)))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1397)+4))
	v1482 = *(*int64)(unsafe.Add(mBase, uint32(v1480+v1476)))
	v1483 = F_FunctionCall2Coll(m, v1468+v1434*int32(28), v1474, v1479, v1482)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L17
	} else {
		goto L230
	}
L230:
	;
	if base.I32_wrap_i64(v1483) != 0 {
		goto L227
	} else {
		goto L231
	}
L231:
	;
	v1487 = v1434 + int32(1)
	v1488 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	if v1487 < v1488 {
		v1434 = v1487
		goto L225
	} else {
		goto L232
	}
L232:
	;
	v1535 = v1406
	goto L220
L233:
	;
	goto L219
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+4)) = v1565
	v1584 = F_palloc0_mul(m, int32(4), v1565)
	mBase = m.M
	v1585 = m.ExcPending
	if v1585 != 0 {
		goto L17
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+8)) = v1584
	v1588 = F_palloc0_mul(m, int32(4), v1565)
	mBase = m.M
	v1589 = m.ExcPending
	if v1589 != 0 {
		goto L17
	} else {
		goto L236
	}
L236:
	;
	v1591 = v1565 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+20)) = v1591
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+12)) = v1588
	v1597 = F_palloc_mul(m, int32(4), v1591)
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L17
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+24)) = v1597
	v1600 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	v1601 = v1565 * v1600
	v1604 = F_palloc(m, v1601<<(uint(int32(3))%32))
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L17
	} else {
		goto L238
	}
L238:
	;
	v1607 = F_palloc_mul(m, int32(4), v1601)
	mBase = m.M
	v1608 = m.ExcPending
	if v1608 != 0 {
		goto L17
	} else {
		goto L239
	}
L239:
	;
	v1609 = int32(0)
	if v1609 < v1565 {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1612 = int32(0)
	v1622 = v1612
	v1623 = v3
	goto L243
L241:
	;
	v1787 = v1609
	v1793 = v3
	goto L242
L242:
	;
	F_pfree(m, v1389)
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L17
	} else {
		goto L259
	}
L243:
	;
	v1642 = int32(2)
	v1643 = v1622 << (uint(v1642) % 32)
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+8))
	v1646 = v1622 * v1600
	*(*int32)(unsafe.Add(mBase, uint32(v1643+v1644))) = v1604 + v1646<<(uint(int32(3))%32)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1651+v1643))) = v1607 + v1646<<(uint(v1642)%32)
	if base.B2i32(v1600 <= v1612) == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	v1787 = v1565
	v1793 = v1778
	goto L242
L245:
	;
	v1659 = v1643 + v1389
	v1663 = int32(0)
	goto L248
L246:
	;
	goto L247
L247:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1643+v1389)))
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759)+12)))
	if v1760 != 0 {
		v1775 = int32(-1)
		v1778 = v1623
		goto L255
	} else {
		goto L256
	}
L248:
	;
	v1689 = v1663 << (uint(int32(2)) % 32)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1659)))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+8))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1689+v1691)))
	if v1693 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L249:
	;
	goto L247
L250:
	;
	v1697 = v1663 << (uint(int32(3)) % 32)
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+4))
	v1700 = *(*int64)(unsafe.Add(mBase, uint32(v1697+v1698)))
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v68)+44))
	v1703 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1701+v1663))))
	v1704 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
	v1708 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1704+v1663<<(uint(int32(1))%32)))))
	v1709 = F_datumCopy(m, v1700, v1703, v1708)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L17
	} else {
		goto L253
	}
L251:
	;
	v1720 = v1693
	goto L252
L252:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+12))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1722+v1643)))
	*(*int32)(unsafe.Add(mBase, uint32(v1724+v1689))) = v1720
	v1728 = v1663 + int32(1)
	if v1728 != v1600 {
		v1663 = v1728
		goto L248
	} else {
		goto L254
	}
L253:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+8))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1711+v1643)))
	*(*int64)(unsafe.Add(mBase, uint32(v1713+v1697))) = v1709
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1659)))
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1716)+8))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1717+v1689)))
	v1720 = v1719
	goto L252
L254:
	;
	goto L249
L255:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1779+v1643))) = v1775
	v1783 = v1622 + int32(1)
	if v1783 != v1565 {
		v1622 = v1783
		v1623 = v1778
		goto L243
	} else {
		goto L258
	}
L256:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1759)))
	v1763 = v1761 << (uint(int32(2)) % 32)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v1765 = v1763 + v1764
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1765)))
	if v1766 != int32(-1) {
		v1775 = v1766
		v1778 = v1623
		goto L255
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1765))) = v1623
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1763)))
	v1775 = v1774
	v1778 = v1623 + int32(1)
	goto L255
L258:
	;
	goto L244
L259:
	;
	if v1365 != int32(-1) {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	v1817 = v1365 << (uint(int32(2)) % 32)
	v1818 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v1817+v1818))) = v1793
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1821+v1817)))
	*(*int32)(unsafe.Add(mBase, uint32(v1284)+32)) = v1823
	goto L262
L261:
	;
	goto L262
L262:
	;
	v1826 = *(*int32)(unsafe.Add(mBase, uint32(v1284)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1826+v1787<<(uint(int32(2))%32)))) = int32(-1)
	v2178 = v1284
	goto L112
L263:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L17
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(370), int32(_a_F_RelationGetPartitionDesc_7))
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L17
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	v1894 = int32(-1)
	v1896 = int32(-1)
	v1897 = v3
	v1901 = v1847
	goto L113
L267:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1857 = m.ExcPending
	if v1857 != 0 {
		goto L17
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(490), int32(_a_F_RelationGetPartitionDesc_8))
	mBase = m.M
	v1862 = m.ExcPending
	if v1862 != 0 {
		goto L17
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_9), int32(0))
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L17
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(520), int32(_a_F_RelationGetPartitionDesc_8))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L17
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L17
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(708), int32(_a_F_RelationGetPartitionDesc_10))
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L17
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+4)) = v1897
	v1922 = F_palloc0_mul(m, int32(4), v1897)
	mBase = m.M
	v1923 = m.ExcPending
	if v1923 != 0 {
		goto L17
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+20)) = v1897
	*(*int64)(unsafe.Add(mBase, uint32(v938)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v938)+8)) = v1922
	v1930 = F_palloc(m, v1897<<(uint(int32(2))%32))
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L17
	} else {
		goto L278
	}
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+24)) = v1930
	v1935 = F_palloc(m, v1897<<(uint(int32(3))%32))
	mBase = m.M
	v1936 = m.ExcPending
	if v1936 != 0 {
		goto L17
	} else {
		goto L279
	}
L279:
	;
	if int32(0) < v1897 {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1942 = int32(0)
	v1953 = v3
	goto L283
L281:
	;
	v2024 = v3
	goto L282
L282:
	;
	F_pfree(m, v1901)
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L17
	} else {
		goto L290
	}
L283:
	;
	v1969 = v1901 + v1942<<(uint(int32(4))%32)
	v1970 = *(*int32)(unsafe.Add(mBase, uint32(v1969)))
	v1972 = v1942 << (uint(int32(2)) % 32)
	v1973 = *(*int32)(unsafe.Add(mBase, uint32(v938)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1972+v1973))) = v1935 + v1942<<(uint(int32(3))%32)
	v1979 = *(*int64)(unsafe.Add(mBase, uint32(v1969)+8))
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(v68)+44))
	v1981 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1980))))
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
	v1983 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1982))))
	v1984 = F_datumCopy(m, v1979, v1981, v1983)
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L17
	} else {
		goto L285
	}
L284:
	;
	v2024 = v2004
	goto L282
L285:
	;
	v1986 = *(*int32)(unsafe.Add(mBase, uint32(v938)+8))
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1986+v1972)))
	*(*int64)(unsafe.Add(mBase, uint32(v1988))) = v1984
	v1991 = v1970 << (uint(int32(2)) % 32)
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v1993 = v1991 + v1992
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1993)))
	if v1994 == int32(-1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1993))) = v1953
	v2000 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v2000+v1991)))
	v2003 = v2002
	v2004 = v1953 + int32(1)
	goto L288
L287:
	;
	v2003 = v1994
	v2004 = v1953
	goto L288
L288:
	;
	v2005 = *(*int32)(unsafe.Add(mBase, uint32(v938)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2005+v1972))) = v2003
	v2009 = v1942 + int32(1)
	if v2009 != v1897 {
		v1942 = v2009
		v1953 = v2004
		goto L283
	} else {
		goto L289
	}
L289:
	;
	goto L284
L290:
	;
	if v1896 != int32(-1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v2043 = v1896 << (uint(int32(2)) % 32)
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v2045 = v2043 + v2044
	v2046 = *(*int32)(unsafe.Add(mBase, uint32(v2045)))
	if v2046 == int32(-1) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	v2061 = v2024
	goto L293
L293:
	;
	if v1894 != int32(-1) {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2045))) = v2024
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(v2052+v2043)))
	v2055 = v2024 + int32(1)
	v2056 = v2054
	goto L296
L295:
	;
	v2055 = v2024
	v2056 = v2046
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+28)) = v2056
	v2061 = v2055
	goto L293
L297:
	;
	v2065 = v1894 << (uint(int32(2)) % 32)
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v2065+v2066))) = v2061
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v2069+v2065)))
	*(*int32)(unsafe.Add(mBase, uint32(v938)+32)) = v2071
	goto L299
L298:
	;
	goto L299
L299:
	;
	if v98 < int32(2) {
		v2178 = v938
		goto L112
	} else {
		goto L300
	}
L300:
	;
	v2076 = int32(-1)
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v938)+4))
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v938)+28))
	v2082 = *(*int32)(unsafe.Add(mBase, uint32(v938)+32))
	if v2077+base.B2i32(v2078 != v2076)+base.B2i32(v2082 != v2076) == v98 {
		v2149 = v2082
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if v2149 == int32(-1) {
		v2178 = v938
		goto L112
	} else {
		goto L313
	}
L302:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v938)+20))
	if v2087 <= int32(0) {
		v2149 = v2082
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v2093 = v2076
	v2094 = v2087
	v2099 = int32(0)
	goto L304
L304:
	;
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v938)+24))
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v2118+v2099<<(uint(int32(2))%32))))
	if v2093 <= v2122 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v938)+32))
	v2149 = v2140
	goto L301
L306:
	;
	v2138 = v2099 + int32(1)
	if v2138 < v2135 {
		v2093 = v2122
		v2094 = v2135
		v2099 = v2138
		goto L304
	} else {
		goto L312
	}
L307:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v938)+28))
	if base.B2i32(v2124 == int32(-1))|base.B2i32(v2122 != v2124) != 0 {
		v2135 = v2094
		goto L306
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v2130 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v2131 = F_bms_add_member(m, v2130, v2122)
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L17
	} else {
		goto L311
	}
L310:
	;
	goto L309
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+16)) = v2131
	v2134 = *(*int32)(unsafe.Add(mBase, uint32(v938)+20))
	v2135 = v2134
	goto L306
L312:
	;
	goto L305
L313:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v938)+16))
	v2171 = F_bms_add_member(m, v2170, v2149)
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L17
	} else {
		goto L314
	}
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v938)+16)) = v2171
	v2228 = v938
	goto L98
L315:
	;
	v2231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)) = uint8(v2231)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v2231
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2242 = F_find_inheritance_children_extended(m, v2236, l1, v2231, v30+int32(127), v30+int32(120))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L17
	} else {
		goto L316
	}
L316:
	;
	if v2242 != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v2247 = int32(0)
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(v2242)+4))
	if v2248 <= v2247 {
		v2290 = v2248
		v2294 = v2231
		v2303 = v117
		v2304 = v119
		v2307 = v2247
		goto L21
	} else {
		goto L320
	}
L318:
	;
	goto L319
L319:
	;
	goto L30
L320:
	;
	v92 = int32(1)
	v93 = v2242 + int32(4)
	v98 = v2248
	v102 = v2231
	v104 = v2242
	goto L29
L321:
	;
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2319 = F_MemoryContextStrdup(m, v2314, v2316+int32(4))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L17
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+36)) = v2319
	v2323 = F_MemoryContextAllocZero(m, v2314, int32(32))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L17
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323))) = v2290
	v2326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2323)+4)) = uint8(v2326)
	if v2294 != 0 {
		goto L326
	} else {
		goto L327
	}
L324:
	;
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2950 != 0 {
		goto L443
	} else {
		goto L444
	}
L325:
	;
	v2917 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[2]))
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+16))
	if v2921 != v2917 {
		goto L427
	} else {
		goto L428
	}
L326:
	;
	v2330 = int32(_a_F_RelationGetPartitionDesc_11)
	v2331 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3])) = v2314
	v2336 = F_palloc(m, int32(36))
	mBase = m.M
	v2337 = m.ExcPending
	if v2337 != 0 {
		goto L17
	} else {
		goto L329
	}
L327:
	;
	v2830 = v2326
	goto L328
L328:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v2830&int32(1) == int32(0)) != 0 {
		goto L325
	} else {
		goto L385
	}
L329:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2307)))
	*(*int32)(unsafe.Add(mBase, uint32(v2336))) = v2338
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+4)) = v2340
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+20)) = v2342
	v2344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	v2346 = F_palloc_mul(m, int32(4), v2340)
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L17
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+8)) = v2346
	v2350 = base.B2i32(v2340 <= int32(0))
	if v2340 <= int32(0) {
		goto L333
	} else {
		goto L334
	}
L331:
	;
	v2625 = F_palloc_mul(m, int32(4), v2342)
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L17
	} else {
		goto L370
	}
L332:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v2451 = base.B2i32(v2449 == int32(104))
	if v2449 == int32(104) {
		goto L347
	} else {
		goto L348
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+12)) = int32(0)
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+16))
	v2418 = F_bms_copy(m, v2417)
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L17
	} else {
		goto L345
	}
L334:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+12))
	if v2351 == int32(0) {
		goto L333
	} else {
		goto L335
	}
L335:
	;
	v2355 = v2340 << (uint(int32(2)) % 32)
	v2356 = F_palloc(m, v2355)
	mBase = m.M
	v2357 = m.ExcPending
	if v2357 != 0 {
		goto L17
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+12)) = v2356
	v2360 = v2344 << (uint(int32(2)) % 32)
	v2363 = F_palloc(m, v2355*v2344)
	mBase = m.M
	v2364 = m.ExcPending
	if v2364 != 0 {
		goto L17
	} else {
		goto L337
	}
L337:
	;
	v2386 = int32(0)
	goto L338
L338:
	;
	v2392 = int32(2)
	v2393 = v2386 << (uint(v2392) % 32)
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2393+v2394))) = v2363 + v2344*v2386<<(uint(v2392)%32)
	if v2360 != 0 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+16))
	v2412 = F_bms_copy(m, v2411)
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L17
	} else {
		goto L344
	}
L340:
	;
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+12))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2401+v2393)))
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+12))
	v2406 = *(*int32)(unsafe.Add(mBase, uint32(v2404+v2393)))
	base.MemoryCopy(m, v2403, v2406, v2360)
	goto L342
L341:
	;
	goto L342
L342:
	;
	v2409 = v2386 + int32(1)
	if v2409 != v2340 {
		v2386 = v2409
		goto L338
	} else {
		goto L343
	}
L343:
	;
	goto L339
L344:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+16)) = v2412
	goto L332
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+16)) = v2418
	if v2340 <= int32(0) {
		goto L331
	} else {
		goto L346
	}
L346:
	;
	goto L332
L347:
	;
	v2452 = int32(2)
	goto L349
L348:
	;
	v2452 = v2344
	goto L349
L349:
	;
	v2456 = F_palloc(m, v2340*v2452<<(uint(int32(3))%32))
	mBase = m.M
	v2457 = m.ExcPending
	if v2457 != 0 {
		goto L17
	} else {
		goto L350
	}
L350:
	;
	v2469 = int32(0)
	goto L351
L351:
	;
	v2488 = v2469 << (uint(int32(2)) % 32)
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2488+v2489))) = v2456 + v2469*v2452<<(uint(int32(3))%32)
	v2496 = int32(0)
	if base.B2i32(v2452 <= int32(0)) == v2496 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	goto L331
L353:
	;
	v2520 = v2496
	goto L356
L354:
	;
	goto L355
L355:
	;
	v2595 = v2469 + int32(1)
	if v2595 != v2340 {
		v2469 = v2595
		goto L351
	} else {
		goto L369
	}
L356:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+12))
	if v2526 != 0 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	goto L355
L358:
	;
	v2565 = v2520 + int32(1)
	if v2565 != v2452 {
		v2520 = v2565
		goto L356
	} else {
		goto L368
	}
L359:
	;
	v2528 = *(*int32)(unsafe.Add(mBase, uint32(v2526+v2488)))
	v2532 = *(*int32)(unsafe.Add(mBase, uint32(v2528+v2520<<(uint(int32(2))%32))))
	if v2532 != 0 {
		goto L358
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v2534 = v2520 << (uint(int32(3)) % 32)
	v2535 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+8))
	v2537 = *(*int32)(unsafe.Add(mBase, uint32(v2535+v2488)))
	v2539 = *(*int64)(unsafe.Add(mBase, uint32(v2534+v2537)))
	if v2449 == int32(104) {
		goto L364
	} else {
		goto L365
	}
L362:
	;
	goto L361
L363:
	;
	v2554 = F_datumCopy(m, v2539, v2551&int32(1), v2550)
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L17
	} else {
		goto L367
	}
L364:
	;
	v2550 = int32(4)
	v2551 = int32(1)
	goto L363
L365:
	;
	goto L366
L366:
	;
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v68)+40))
	v2546 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2542+v2520<<(uint(int32(1))%32)))))
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(v68)+44))
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2547+v2520))))
	v2550 = v2546
	v2551 = v2549
	goto L363
L367:
	;
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2336)+8))
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v2556+v2488)))
	*(*int64)(unsafe.Add(mBase, uint32(v2558+v2534))) = v2554
	goto L358
L368:
	;
	goto L357
L369:
	;
	goto L352
L370:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+24)) = v2625
	v2629 = v2342 << (uint(int32(2)) % 32)
	if v2629 != 0 {
		goto L371
	} else {
		goto L372
	}
L371:
	;
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+24))
	base.MemoryCopy(m, v2625, v2630, v2629)
	goto L373
L372:
	;
	goto L373
L373:
	;
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+28)) = v2632
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v2307)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2336)+32)) = v2634
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2323)+20)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+16)) = v2336
	v2643 = F_palloc(m, v2290<<(uint(int32(2))%32))
	mBase = m.M
	v2644 = m.ExcPending
	if v2644 != 0 {
		goto L17
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+8)) = v2643
	v2646 = F_palloc(m, v2290)
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L17
	} else {
		goto L375
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2323)+12)) = v2646
	if v2290 <= int32(0) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3])) = v2331
	v2802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+127)))
	v2830 = v2802
	goto L328
L377:
	;
	v2651 = int32(0)
	if v2290 != int32(1) {
		goto L378
	} else {
		goto L379
	}
L378:
	;
	v2667 = int32(0)
	v2673 = v2651
	goto L381
L379:
	;
	v2743 = v2651
	goto L380
L380:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+8))
	v2757 = int32(2)
	v2758 = v2743 << (uint(v2757) % 32)
	v2759 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	v2761 = *(*int32)(unsafe.Add(mBase, uint32(v2758+v2759)))
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(v2758+v2303)))
	*(*int32)(unsafe.Add(mBase, uint32(v2756+v2761<<(uint(v2757)%32)))) = v2766
	v2768 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+12))
	v2771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2743+v2304))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2761+v2768))) = uint8(v2771)
	goto L376
L381:
	;
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+8))
	v2687 = int32(2)
	v2688 = v2673 << (uint(v2687) % 32)
	v2689 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2688+v2689)))
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v2688+v2303)))
	*(*int32)(unsafe.Add(mBase, uint32(v2686+v2691<<(uint(v2687)%32)))) = v2696
	v2698 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+12))
	v2701 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2673+v2304))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2691+v2698))) = uint8(v2701)
	v2703 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+8))
	v2705 = v2673 | int32(1)
	v2707 = v2705 << (uint(v2687) % 32)
	v2708 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	v2710 = *(*int32)(unsafe.Add(mBase, uint32(v2707+v2708)))
	v2715 = *(*int32)(unsafe.Add(mBase, uint32(v2707+v2303)))
	*(*int32)(unsafe.Add(mBase, uint32(v2703+v2710<<(uint(v2687)%32)))) = v2715
	v2717 = *(*int32)(unsafe.Add(mBase, uint32(v2323)+12))
	v2720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2705+v2304))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2710+v2717))) = uint8(v2720)
	v2723 = v2673 + v2687
	v2725 = v2667 + v2687
	if v2725 != v2290&int32(2147483646) {
		v2667 = v2725
		v2673 = v2723
		goto L381
	} else {
		goto L383
	}
L382:
	;
	if v2290&int32(1) == int32(0) {
		goto L376
	} else {
		goto L384
	}
L383:
	;
	goto L382
L384:
	;
	v2743 = v2723
	goto L380
L385:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L386
L386:
	;
	if base.B2i32(v2837 != int32(0)) == int32(0) {
		goto L325
	} else {
		goto L387
	}
L387:
	;
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	v2844 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[2]))
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+16))
	if v2848 != v2844 {
		goto L389
	} else {
		goto L390
	}
L388:
	;
	if v2842 == int32(0) {
		goto L324
	} else {
		goto L405
	}
L389:
	;
	if v2848 == int32(0) {
		goto L392
	} else {
		goto L393
	}
L390:
	;
	goto L391
L391:
	;
	goto L388
L392:
	;
	if v2844 != 0 {
		goto L399
	} else {
		goto L400
	}
L393:
	;
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+28))
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+24))
	if v2853 != 0 {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	if v2852 == int32(0) {
		goto L392
	} else {
		goto L398
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2853)+28)) = v2852
	goto L394
L396:
	;
	goto L397
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2848)+20)) = v2852
	goto L394
L398:
	;
	v2858 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2852)+24)) = v2858
	goto L392
L399:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+16)) = v2844
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2844)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+28)) = v2865
	if v2865 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	goto L401
L401:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2314)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+16)) = int32(0)
	goto L391
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2865)+24)) = v2314
	goto L404
L403:
	;
	goto L404
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+20)) = v2314
	goto L388
L405:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2879 != 0 {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2879)+16))
	if v2883 != v2314 {
		goto L410
	} else {
		goto L411
	}
L407:
	;
	goto L408
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2314
	v2914 = *(*int32)(unsafe.Add(mBase, uint32(v30)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v2914
	v2988 = v2323
	goto L4
L409:
	;
	goto L408
L410:
	;
	if v2883 == int32(0) {
		goto L413
	} else {
		goto L414
	}
L411:
	;
	goto L412
L412:
	;
	goto L409
L413:
	;
	if v2314 != 0 {
		goto L420
	} else {
		goto L421
	}
L414:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(v2879)+28))
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(v2879)+24))
	if v2888 != 0 {
		goto L416
	} else {
		goto L417
	}
L415:
	;
	if v2887 == int32(0) {
		goto L413
	} else {
		goto L419
	}
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2888)+28)) = v2887
	goto L415
L417:
	;
	goto L418
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2883)+20)) = v2887
	goto L415
L419:
	;
	v2893 = *(*int32)(unsafe.Add(mBase, uint32(v2879)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2887)+24)) = v2893
	goto L413
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2879)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2879)+16)) = v2314
	v2900 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2879)+28)) = v2900
	if v2900 != 0 {
		goto L423
	} else {
		goto L424
	}
L421:
	;
	goto L422
L422:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2879)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2879)+16)) = int32(0)
	goto L412
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2900)+24)) = v2879
	goto L425
L424:
	;
	goto L425
L425:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+20)) = v2879
	goto L409
L426:
	;
	goto L324
L427:
	;
	if v2921 == int32(0) {
		goto L430
	} else {
		goto L431
	}
L428:
	;
	goto L429
L429:
	;
	goto L426
L430:
	;
	if v2917 != 0 {
		goto L437
	} else {
		goto L438
	}
L431:
	;
	v2925 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+28))
	v2926 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+24))
	if v2926 != 0 {
		goto L433
	} else {
		goto L434
	}
L432:
	;
	if v2925 == int32(0) {
		goto L430
	} else {
		goto L436
	}
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2926)+28)) = v2925
	goto L432
L434:
	;
	goto L435
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2921)+20)) = v2925
	goto L432
L436:
	;
	v2931 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2925)+24)) = v2931
	goto L430
L437:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+16)) = v2917
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2917)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+28)) = v2938
	if v2938 != 0 {
		goto L440
	} else {
		goto L441
	}
L438:
	;
	goto L439
L439:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2314)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+16)) = int32(0)
	goto L429
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2938)+24)) = v2314
	goto L442
L441:
	;
	goto L442
L442:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2917)+20)) = v2314
	goto L426
L443:
	;
	v2954 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+16))
	if v2954 != v2314 {
		goto L447
	} else {
		goto L448
	}
L444:
	;
	goto L445
L445:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2314
	v2988 = v2323
	goto L4
L446:
	;
	goto L445
L447:
	;
	if v2954 == int32(0) {
		goto L450
	} else {
		goto L451
	}
L448:
	;
	goto L449
L449:
	;
	goto L446
L450:
	;
	if v2314 != 0 {
		goto L457
	} else {
		goto L458
	}
L451:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+28))
	v2959 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+24))
	if v2959 != 0 {
		goto L453
	} else {
		goto L454
	}
L452:
	;
	if v2958 == int32(0) {
		goto L450
	} else {
		goto L456
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2959)+28)) = v2958
	goto L452
L454:
	;
	goto L455
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2954)+20)) = v2958
	goto L452
L456:
	;
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v2950)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2958)+24)) = v2964
	goto L450
L457:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2950)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2950)+16)) = v2314
	v2971 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2950)+28)) = v2971
	if v2971 != 0 {
		goto L460
	} else {
		goto L461
	}
L458:
	;
	goto L459
L459:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2950)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2950)+16)) = int32(0)
	goto L449
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2971)+24)) = v2950
	goto L462
L461:
	;
	goto L462
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2314)+20)) = v2950
	goto L446
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v158
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_12), v30)
	mBase = m.M
	v3023 = m.ExcPending
	if v3023 != 0 {
		goto L17
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(280), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v3028 = m.ExcPending
	if v3028 != 0 {
		goto L17
	} else {
		goto L465
	}
L465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v158
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_15), v30+int32(32))
	mBase = m.M
	v3038 = m.ExcPending
	if v3038 != 0 {
		goto L17
	} else {
		goto L467
	}
L467:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(282), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v3043 = m.ExcPending
	if v3043 != 0 {
		goto L17
	} else {
		goto L468
	}
L468:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v158
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_16), v30+int32(16))
	mBase = m.M
	v3054 = m.ExcPending
	if v3054 != 0 {
		goto L17
	} else {
		goto L470
	}
L470:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(296), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v3059 = m.ExcPending
	if v3059 != 0 {
		goto L17
	} else {
		goto L471
	}
L471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetPrimaryKeyIndex(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+27)))
	if v3 == int32(0) {
		v6 = F_RelationGetIndexList(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_list_free(m, v6)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				if l1 == int32(0) {
					v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
					if v15 != 0 {
						v17 = int32(0)
					} else {
						v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
						v17 = v16
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
					v17 = v16
				}
				return v17
			}
		}
	} else {
		if l1 == int32(0) {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
			if v15 != 0 {
				v17 = int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
				v17 = v16
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
			v17 = v16
		}
		return v17
	}
}
func F_RelationIdGetRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[0]))
	v16 = F_hash_search(m, v11, v6+int32(12), v2, v2)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if v16 == int32(0) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v48 = F_RelationBuildDesc(m, v46, int32(1))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				if v48 == int32(0) {
					v68 = v2
					m.G0 = v6 + int32(16)
					return v68
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
					F_ResourceOwnerEnlarge(m, v53)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v56 + int32(1)
						v61 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[2]))
						if v61 != 0 {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
							F_ResourceOwnerRemember(m, v63, base.I64_extend_i32_u(v48), int32(_a_F_RelationIdGetRelation_0))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v68 = v48
								m.G0 = v6 + int32(16)
								return v68
							}
						} else {
							v68 = v48
							m.G0 = v6 + int32(16)
							return v68
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			if v22 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v48 = F_RelationBuildDesc(m, v46, int32(1))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					if v48 == int32(0) {
						v68 = v2
						m.G0 = v6 + int32(16)
						return v68
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
						F_ResourceOwnerEnlarge(m, v53)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v56 + int32(1)
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[2]))
							if v61 != 0 {
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
								F_ResourceOwnerRemember(m, v63, base.I64_extend_i32_u(v48), int32(_a_F_RelationIdGetRelation_0))
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									v68 = v48
									m.G0 = v6 + int32(16)
									return v68
								}
							} else {
								v68 = v48
								m.G0 = v6 + int32(16)
								return v68
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
				if v25 != 0 {
					v68 = v2
					m.G0 = v6 + int32(16)
					return v68
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
					F_ResourceOwnerEnlarge(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v30 + int32(1)
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[2]))
						if v35 != 0 {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
							F_ResourceOwnerRemember(m, v37, base.I64_extend_i32_u(v22), int32(_a_F_RelationIdGetRelation_0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
								if v42 != 0 {
									v68 = v22
									m.G0 = v6 + int32(16)
									return v68
								} else {
									F_RelationRebuildRelation(m, v22)
									mBase = m.M
									v44 = m.ExcPending
									if v44 != 0 {
										return int32(0)
									} else {
										v68 = v22
										m.G0 = v6 + int32(16)
										return v68
									}
								}
							}
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
							if v42 != 0 {
								v68 = v22
								m.G0 = v6 + int32(16)
								return v68
							} else {
								F_RelationRebuildRelation(m, v22)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v68 = v22
									m.G0 = v6 + int32(16)
									return v68
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationInitTableAccessMethod(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)))
	if v10 == int32(83) {
		v44 = int32(3)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v44
		v49 = v44
		v50 = m.G0
		v52 = v50 - int32(16)
		m.G0 = v52
		v54 = F_OidFunctionCall0Coll(m, v49)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			v56 = base.I32_wrap_i64(v54)
			if v56 != 0 {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				if v57 == int32(445) {
					v73 = int32(16)
					m.G0 = v52 + v73
					*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v56
					m.G0 = v7 + v73
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
						F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
					F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v13) < base.Ui32(int32(_a_F_RelationInitTableAccessMethod_3)) {
			v44 = int32(3)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v44
			v49 = v44
			v50 = m.G0
			v52 = v50 - int32(16)
			m.G0 = v52
			v54 = F_OidFunctionCall0Coll(m, v49)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				v56 = base.I32_wrap_i64(v54)
				if v56 != 0 {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					if v57 == int32(445) {
						v73 = int32(16)
						m.G0 = v52 + v73
						*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v56
						m.G0 = v7 + v73
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
							F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
						F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
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
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v18 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+84)))
			v19 = F_SearchSysCache1(m, int32(2), v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				if v19 != 0 {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v21+v22)+68))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v24
					F_ReleaseCatCache(m, v19)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
						v49 = v28
						v50 = m.G0
						v52 = v50 - int32(16)
						m.G0 = v52
						v54 = F_OidFunctionCall0Coll(m, v49)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = base.I32_wrap_i64(v54)
							if v56 != 0 {
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								if v57 == int32(445) {
									v73 = int32(16)
									m.G0 = v52 + v73
									*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v56
									m.G0 = v7 + v73
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
										F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
									F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(37), int32(_a_F_RelationInitTableAccessMethod_2))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
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
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = v34
						F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_4), v7)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_5), int32(1858), int32(_a_F_RelationInitTableAccessMethod_6))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
}
func F_RelationInvalidatesSnapshotsOnly(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	v3 = int32(1)
	if l0 <= int32(2963) {
		if base.B2i32(l0 == int32(1214))|base.B2i32(base.Ui32(l0-int32(2608)) < base.Ui32(int32(2))) != 0 {
			v20 = v3
		} else {
			if l0 != int32(2396) {
				v20 = int32(0)
			} else {
				v20 = v3
			}
		}
	} else {
		switch l0 - int32(3592) {
		case 0, 4:
			v20 = v3
		case 1, 2, 3:
			v20 = int32(0)
		default:
			if l0 == int32(2964) {
				v20 = v3
			} else {
				v20 = int32(0)
			}
		}
	}
	return v20
}
func F_RelationPreserveStorage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPreserveStorage[0]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = v7
	v12 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v14 != v15 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	if v13 != 0 {
		v11 = v13
		v12 = v30
		goto L4
	} else {
		goto L25
	}
L7:
	;
	v30 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v17 != v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = v11
	goto L6
L11:
	;
	goto L12
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v20 != v21 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v30 = v11
	goto L6
L14:
	;
	goto L15
L15:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if l1 != v23 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v30 = v11
	goto L6
L17:
	;
	goto L18
L18:
	;
	if v12 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_pfree(m, v11)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v13
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationPreserveStorage[0])) = v13
	goto L19
L23:
	;
	return
L24:
	;
	v30 = v12
	goto L6
L25:
	;
	goto L5
}
func F_RelationPutHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	if l0 < int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPutHeapTuple[0]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l0^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPutHeapTuple[1]))
		v23 = v17 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v28 = F_PageAddItemExtended(m, v23, v24, v25, int32(0), int32(2))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		if v28 != 0 {
			if l0 < int32(0) {
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPutHeapTuple[2]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(l0^int32(-1))*int32(56))+16))
				v48 = v39
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPutHeapTuple[3]))
				v42 = int32(56)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+l0*v42-v42)+16))
				v48 = v47
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v28)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v48)
			v52 = int32(base.Ui32(v48) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
			if l2 == int32(0) {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v23+v28<<(uint(int32(2))%32))+20))
				v62 = v23 + v59&int32(_a_F_RelationPutHeapTuple_0)
				v64 = l1 + int32(4)
				v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)) = uint16(v65)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
				*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v67
			} else {
			}
			return
		} else {
			F_errstart_cold(m, int32(24), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_RelationPutHeapTuple_1), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_RelationPutHeapTuple_2), int32(63), int32(_a_F_RelationPutHeapTuple_3))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
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
func F_find_relation_notnullatts(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+116))
	if v10 == int32(0) {
		v28 = int32(0)
		m.G0 = v6 + int32(16)
		return v28
	} else {
		v14 = int32(0)
		v20 = F_hash_search(m, v10, v6+int32(12), v14, v6+int32(11))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)))
			if v24 != int32(1) {
				v28 = v14
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
				v28 = v27
			}
			m.G0 = v6 + int32(16)
			return v28
		}
	}
}
func F_get_relation_data_width(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = F_table_open(m, l0, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_get_rel_data_width(m, v4, l1)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_relation_close(m, v4, int32(0))
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v8
			}
		}
	}
}
func F_set_relation_column_names(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
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
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v19 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v163 < v154 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v24 = F_relation_open(m, v22, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v19 != int32(3) {
		goto L20
	} else {
		goto L21
	}
L5:
	;
	return
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v30 = F_palloc(m, v27<<(uint(int32(2))%32))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if int32(0) < v27 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v38 = v4
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v57 = v26 + v51<<(uint(int32(3))%32) + v38*int32(100)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+119)))
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v64 = int32(0)
	goto L15
L14:
	;
	v62 = F_pstrdup(m, v57+int32(32))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v38<<(uint(int32(2))%32)))) = v64
	v67 = v38 + int32(1)
	if v67 != v27 {
		v38 = v67
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v64 = v62
	goto L15
L17:
	;
	goto L12
L18:
	;
	v154 = v27
	v159 = v30
	goto L1
L19:
	;
	if v105 != 0 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v103
	v105 = v103
	goto L19
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v88 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v91 = int32(1)
	v92 = int32(0)
	F_expandRTE(m, l1, v91, v92, v92, int32(-1), v91, v17+int32(12), v92)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v105 = v101
	goto L19
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v108 = v106
	goto L26
L25:
	;
	v108 = int32(0)
	goto L26
L26:
	;
	v111 = F_palloc(m, v108<<(uint(int32(2))%32))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v113 == int32(0) {
		v154 = v108
		v159 = v111
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v116 <= int32(0) {
		v154 = v108
		v159 = v111
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v124 = int32(0)
	goto L30
L30:
	;
	v135 = v124 << (uint(int32(2)) % 32)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v137+v135)))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v142 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v154 = v108
	v159 = v111
	goto L1
L32:
	;
	v143 = v140
	goto L34
L33:
	;
	v143 = int32(0)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111+v135))) = v143
	v146 = v124 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v146 < v147 {
		v124 = v146
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v165 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	v186 = F_palloc(m, v154<<(uint(int32(2))%32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L47
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v154
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v179
	goto L38
L40:
	;
	v169 = F_palloc0_mul(m, int32(4), v154)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v172 = F_mul_size(m, int32(4), v163)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v179 = v169
	goto L39
L44:
	;
	v175 = F_mul_size(m, int32(4), v154)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v177 = F_repalloc0(m, v165, v172, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v179 = v177
	goto L39
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v186
	v189 = F_palloc(m, v154)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v189
	F_build_colinfo_names_hash(m, l2)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	v194 = int32(0)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	if v196 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v198 = v197
	goto L52
L51:
	;
	v198 = v4
	goto L52
L52:
	;
	if int32(0) < v154 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v206 = int32(0)
	v209 = v194
	v214 = v4
	goto L56
L54:
	;
	v309 = v194
	v314 = v4
	goto L55
L55:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v316 != 0 {
		goto L80
	} else {
		goto L81
	}
L56:
	;
	v217 = v206 << (uint(int32(2)) % 32)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v159+v217)))
	if v219 == int32(0) {
		v295 = v209
		v298 = v214
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v309 = v295
	v314 = v298
	goto L55
L58:
	;
	v300 = v206 + int32(1)
	if v300 != v154 {
		v206 = v300
		v209 = v295
		v214 = v298
		goto L56
	} else {
		goto L79
	}
L59:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222+v217)))
	if v224 != 0 {
		v251 = v224
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v254+v209<<(uint(int32(2))%32)))) = v251
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v259+v209))) = uint8(base.B2i32(v198 <= v206))
	v264 = v209 + int32(1)
	if v214 != 0 {
		goto L69
	} else {
		goto L70
	}
L61:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v225 == int32(0) {
		v237 = v219
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v239 = F_make_colname_unique(m, v237, l0, l2)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L66
	}
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+8))
	if v228 == int32(0) {
		v237 = v219
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v231 <= v206 {
		v237 = v219
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233+v217)))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	v237 = v236
	goto L62
L66:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v217))) = v239
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v244 == int32(0) {
		v251 = v239
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v249 = F_hash_search(m, v244, v239, int32(1), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	v251 = v239
	goto L60
L69:
	;
	v295 = v264
	v298 = int32(1)
	goto L58
L70:
	;
	goto L71
L71:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251))))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219))))
	if base.B2i32(v268 == int32(0))|base.B2i32(v268 != v271) != 0 {
		v289 = v268
		v290 = v271
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v295 = v264
	v298 = base.B2i32(v289-v290 != int32(0))
	goto L58
L73:
	;
	goto L72
L74:
	;
	v274 = v251
	v275 = v219
	goto L75
L75:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v279 == int32(0) {
		v289 = v279
		v290 = v278
		goto L73
	} else {
		goto L77
	}
L76:
	;
	v289 = v279
	v290 = v278
	goto L73
L77:
	;
	v282 = int32(1)
	if v279 == v278 {
		v274 = v274 + v282
		v275 = v275 + v282
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L57
L80:
	;
	F_hash_destroy(m, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v309
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	switch v322 {
	case 0:
		goto L88
	default:
		goto L85
	case 3:
		goto L87
	case 4:
		goto L86
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(0)
	goto L82
L84:
	;
	m.G0 = v17 + int32(16)
	return
L85:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v328 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v326)
	goto L84
L87:
	;
	v324 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v324)
	goto L84
L88:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v314)
	goto L84
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v314)
	goto L84
L90:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	if v331 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v334)
	goto L84
}
