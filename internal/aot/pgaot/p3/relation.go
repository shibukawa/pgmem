package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetRelationPublications(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13850(m, l0, int32(53))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
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
	var v42 int32
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	v38 = v14 + int32(-48)
	v39 = int32(3)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
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
	v193 = v2
	v194 = v35
	goto L11
L11:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
	v193 = v183
	v194 = v177
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
	v104 = F_text_to_cstring(m, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	v96 = F_getmissingattr(m, v49, int32(8), v14+int32(-49))
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
	v101 = F_fastgetattr_3(m, v59, int32(8), v49, v14+int32(-49))
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
	v106 = int32(_a_F_RelationBuildRuleLock_2)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v24
	v110 = F_stringToNode(m, v104)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v110
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v107
	F_pfree(m, v104)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+18)))
	if base.Ui32(v118&int32(2047)) <= base.Ui32(int32(6)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v134 = F_text_to_cstring(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L30
	}
L25:
	;
	v126 = F_getmissingattr(m, v49, int32(7), v14+int32(-49))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v131 = F_fastgetattr_3(m, v59, int32(7), v49, v14+int32(-49))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	v133 = v126
	goto L24
L29:
	;
	v133 = v131
	goto L24
L30:
	;
	v136 = int32(_a_F_RelationBuildRuleLock_2)
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v24
	v140 = F_stringToNode(m, v134)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v140
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildRuleLock[1])) = v137
	F_pfree(m, v134)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v149 != int32(1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	F_setRuleCheckAsUser(m, v148, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L39
	}
L34:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v147)+80))
	v163 = v161
	goto L33
L35:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+119)))
	if v152 != int32(118) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v155 == int32(0) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+5)))
	if v159 != 0 {
		v163 = int32(0)
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	F_setRuleCheckAsUser(m, v166, v163)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
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
	v172 = F_repalloc(m, v66, v64<<(uint(int32(3))%32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	v176 = v64
	v177 = v66
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177+v65<<(uint(int32(2))%32)))) = v74
	v183 = v65 + int32(1)
	v184 = F_systable_getnext(m, v54)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	v176 = v64 << (uint(int32(1)) % 32)
	v177 = v172
	goto L43
L45:
	;
	if v184 != 0 {
		v59 = v184
		v64 = v176
		v65 = v183
		v66 = v177
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
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v193 == int32(0) {
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
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v211 = F_MemoryContextAlloc(m, v24, int32(8))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L54
	}
L53:
	;
	goto L49
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v193
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v211
	goto L49
}
func F_RelationBuildTriggers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
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
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
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
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
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
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	v2 = int32(0)
	v30 = m.G0
	v32 = v30 + int32(-64)
	m.G0 = v32
	v35 = F_palloc(m, int32(960))
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
	v38 = v30 + int32(-48)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v38, int32(2), int32(3), int32(184), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v47 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L84
	}
L5:
	;
	v50 = int32(1)
	v53 = F_systable_beginscan(m, v47, int32(2701), v50, int32(0), v50, v38)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v55 = F_systable_getnext(m, v53)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v55 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v63 = v55
	v64 = int32(16)
	v67 = v2
	v68 = v35
	goto L11
L9:
	;
	v316 = v2
	v317 = v35
	goto L10
L10:
	;
	F_systable_endscan(m, v53)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L59
	}
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+22)))
	v89 = v87 + v88
	if v64 <= v67 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v316 = v304
	v317 = v98
	goto L10
L13:
	;
	v93 = F_repalloc(m, v68, v64*int32(120))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v97 = v64
	v98 = v68
	goto L15
L15:
	;
	v101 = v98 + v67*int32(60)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102
	v108 = F_DirectFunctionCall1Coll(m, int32(581), int32(0), v89+int32(12))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v97 = v64 << (uint(int32(1)) % 32)
	v98 = v93
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v108
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v89)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v111
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+12)) = uint16(v113)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+14)) = uint8(v115)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+83)))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+15)) = uint8(v117)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+16)) = uint8(base.B2i32(v119 != v120))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+24)) = v125
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v89)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+28)) = v127
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+32)) = uint8(v129)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+97)))
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+33)) = uint8(v131)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+98)))
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+34)) = uint16(v133)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v89)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+36)) = uint16(v135)
	v138 = v135 << (uint(int32(16)) % 32)
	if v120 < v138 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(0) < base.I32_extend16_s(v155) {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v143 = F_palloc(m, int32(base.Ui32(v138)>>(uint(int32(15))%32)))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+40)) = int32(0)
	v155 = v133
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+40)) = v143
	v146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+36)))
	v148 = v146 << (uint(int32(1)) % 32)
	if v148 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	base.MemoryCopy(m, v143, v89+int32(124), v148)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+34)))
	v155 = v152
	goto L18
L26:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v267 = F_fastgetattr_2(m, v63, int32(18), v264, v30+int32(-49))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L42
	}
L27:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v164 = F_fastgetattr_2(m, v63, int32(16), v161, v30+int32(-49))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+44)) = int32(0)
	goto L26
L30:
	;
	v166 = F_pg_detoast_datum_packed(m, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+15)))
	if v168 == int32(1) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
	v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+34)))
	v175 = F_palloc(m, v172<<(uint(int32(2))%32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+44)) = v175
	v178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+34)))
	if v178 <= int32(0) {
		goto L26
	} else {
		goto L34
	}
L34:
	;
	v181 = int32(1)
	if v171&v181 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v185 = v181
	goto L37
L36:
	;
	v185 = int32(4)
	goto L37
L37:
	;
	v190 = v166 + v185
	v191 = int32(0)
	goto L38
L38:
	;
	v217 = F_pstrdup(m, v190)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L26
L40:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v101)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v219+v191<<(uint(int32(2))%32)))) = v217
	v224 = F_strlen(m, v190)
	mBase = m.M
	v226 = int32(1)
	v229 = v191 + v226
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+34)))
	if v229 < v230 {
		v190 = v224 + v190 + v226
		v191 = v229
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+15)))
	if v269 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v275 = int32(0)
	goto L45
L44:
	;
	v273 = F_DirectFunctionCall1Coll(m, int32(581), int32(0), v267)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+52)) = v275
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v281 = F_fastgetattr_2(m, v63, int32(19), v278, v30+int32(-49))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	v275 = v273
	goto L45
L47:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+15)))
	if v283 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v289 = int32(0)
	goto L50
L49:
	;
	v287 = F_DirectFunctionCall1Coll(m, int32(581), int32(0), v281)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+56)) = v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v295 = F_fastgetattr_2(m, v63, int32(17), v292, v30+int32(-49))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v289 = v287
	goto L50
L52:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+15)))
	if v297 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v301 = int32(0)
	goto L55
L54:
	;
	v299 = F_text_to_cstring(m, v295)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+48)) = v301
	v304 = v67 + int32(1)
	v305 = F_systable_getnext(m, v53)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	v301 = v299
	goto L55
L57:
	;
	if v305 != 0 {
		v63 = v305
		v64 = v97
		v67 = v304
		v68 = v98
		goto L11
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	F_relation_close(m, v47, int32(1))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v316 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	m.G0 = v32 - int32(-64)
	return
L62:
	;
	F_pfree(m, v317)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v346 = F_palloc0(m, int32(32))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L66
	}
L65:
	;
	goto L61
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346)+4)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v317
	if int32(0) < v316 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+28)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+27)))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+25)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+24)))
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+23)))
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+22)))
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+21)))
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+20)))
	v360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+19)))
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+18)))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+17)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+16)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+15)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+14)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+13)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+11)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+10)))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+9)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+8)))
	v379 = v352
	v380 = v354
	v381 = int32(0)
	v384 = v353
	v385 = v355
	v386 = v356
	v387 = v357
	v388 = v358
	v389 = v359
	v390 = v360
	v391 = v361
	v392 = v362
	v393 = v363
	v394 = v364
	v395 = v365
	v396 = v366
	v397 = v367
	v398 = v368
	v399 = v369
	v400 = v370
	v401 = v371
	goto L70
L68:
	;
	goto L69
L69:
	;
	v547 = int32(_a_F_RelationBuildTriggers_0)
	v548 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0]))
	v551 = *(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0])) = v551
	v553 = F_CopyTriggerDesc(m, v346)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L82
	}
L70:
	;
	v404 = v317 + v381*int32(60)
	v405 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v404)+12)))
	v407 = v405 & int32(99)
	v410 = v385 | base.B2i32(v407 == int32(32))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+24)) = uint8(v410)
	v414 = v386 | base.B2i32(v407 == int32(34))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+23)) = uint8(v414)
	v417 = v405 & int32(75)
	v420 = v387 | base.B2i32(v417 == int32(8))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+22)) = uint8(v420)
	v424 = v388 | base.B2i32(v417 == int32(10))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+21)) = uint8(v424)
	v428 = v389 | base.B2i32(v417 == int32(73))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+20)) = uint8(v428)
	v432 = v390 | base.B2i32(v417 == int32(9))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+19)) = uint8(v432)
	v436 = v391 | base.B2i32(v417 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+18)) = uint8(v436)
	v439 = v405 & int32(83)
	v442 = v392 | base.B2i32(v439 == int32(16))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+17)) = uint8(v442)
	v446 = v393 | base.B2i32(v439 == int32(18))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+16)) = uint8(v446)
	v450 = v394 | base.B2i32(v439 == int32(81))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+15)) = uint8(v450)
	v454 = v395 | base.B2i32(v439 == int32(17))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+14)) = uint8(v454)
	v458 = v396 | base.B2i32(v439 == int32(19))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+13)) = uint8(v458)
	v461 = v405 & int32(71)
	v462 = int32(4)
	v464 = v397 | base.B2i32(v461 == v462)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+12)) = uint8(v464)
	v468 = v398 | base.B2i32(v461 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+11)) = uint8(v468)
	v472 = v399 | base.B2i32(v461 == int32(69))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+10)) = uint8(v472)
	v476 = v400 | base.B2i32(v461 == int32(5))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+9)) = uint8(v476)
	v480 = v401 | base.B2i32(v461 == int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+8)) = uint8(v480)
	if v405&v462 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L69
L72:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v404)+56))
	v488 = base.B2i32(v484 != int32(0))
	goto L74
L73:
	;
	v488 = int32(0)
	goto L74
L74:
	;
	v489 = v488 | v380
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+25)) = uint8(v489)
	if v405&int32(16) != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+26)))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v404)+52))
	v495 = int32(0)
	v497 = v493 | base.B2i32(v494 != v495)
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+26)) = uint8(v497)
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v404)+56))
	v503 = base.B2i32(v499 != v495)
	goto L77
L76:
	;
	v503 = int32(0)
	goto L77
L77:
	;
	v504 = v503 | v384
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+27)) = uint8(v504)
	if v405&int32(8) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v404)+52))
	v512 = base.B2i32(v508 != int32(0))
	goto L80
L79:
	;
	v512 = int32(0)
	goto L80
L80:
	;
	v513 = v512 | v379
	*(*uint8)(unsafe.Add(mBase, uint32(v346)+28)) = uint8(v513)
	v516 = v381 + int32(1)
	if v516 != v316 {
		v379 = v513
		v380 = v489
		v381 = v516
		v384 = v504
		v385 = v410
		v386 = v414
		v387 = v420
		v388 = v424
		v389 = v428
		v390 = v432
		v391 = v436
		v392 = v442
		v393 = v446
		v394 = v450
		v395 = v454
		v396 = v458
		v397 = v464
		v398 = v468
		v399 = v472
		v400 = v476
		v401 = v480
		goto L70
	} else {
		goto L81
	}
L81:
	;
	goto L71
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v553
	*(*int32)(unsafe.Add(mBase, _c_F_RelationBuildTriggers[0])) = v548
	F_FreeTriggerDesc(m, v346)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	goto L61
L84:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v596 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationBuildTriggers_1), v32)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_RelationBuildTriggers_2), int32(1946), int32(_a_F_RelationBuildTriggers_3))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
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
										F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2563), int32(_a_F_RelationClearRelation_3))
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
									F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2563), int32(_a_F_RelationClearRelation_3))
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
									F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2563), int32(_a_F_RelationClearRelation_3))
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
								F_errfinish(m, int32(_a_F_RelationClearRelation_2), int32(2563), int32(_a_F_RelationClearRelation_3))
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
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
													v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
												v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
										v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v104 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCreateStorage[5]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v104
									v112 = F_hash_create(m, int32(_a_F_RelationCreateStorage_1), int32(16), v10+int32(32), int32(1064))
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
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
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
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
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
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(1648)
	m.G0 = v20
	v23 = F_index_open(m, l1, int32(3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = F_GetRelationIdentityOrPK(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(4)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
	v35 = F_SysCacheGetAttrNotNull(m, int32(34), v33, int32(18))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v38 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+10)))
	if int32(0) < v38 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L62
	}
L6:
	;
	v50 = v37
	v51 = int32(0)
	v56 = v5
	goto L9
L7:
	;
	v142 = v5
	goto L8
L8:
	;
	v153 = int32(0)
	v155 = F_index_beginscan(m, l0, v23, v20+int32(40), v153, v142, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+int32(48)+v51<<(uint(int32(1))%32)))))
	if v66 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v142 = v125
	goto L8
L11:
	;
	v68 = v51 << (uint(int32(2)) % 32)
	v69 = v35 + int32(24) + v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = F_get_opclass_input_type(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v122 = v50
	v125 = v56
	goto L13
L13:
	;
	v129 = v51 + int32(1)
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v122)+10)))
	if v129 < v130 {
		v50 = v122
		v51 = v129
		v56 = v125
		goto L9
	} else {
		goto L24
	}
L14:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v74 = F_get_opclass_family(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+84))
	v80 = F_IndexAmTranslateCompareType(m, int32(3), v78, v74, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v83 = F_get_opfamily_member(m, v74, v71, v71, base.I32_extend16_s(v80))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v83 == int32(0) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v91 = v20 + int32(112) + v56*int32(48)
	v95 = F_get_opcode(m, v83)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v99 = v66 - int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+v99<<(uint(int32(2))%32))))
	F_ScanKeyInit(m, v91, base.I32_extend16_s(v51+int32(1)), v80, v95, v103)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v23)+248))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106+v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+12)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v99))))
	if v112 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v115 | int32(65)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v23)+192))
	v122 = v121
	v125 = v56 + int32(1)
	goto L13
L24:
	;
	goto L10
L25:
	;
	v160 = int32(0)
	goto L26
L26:
	;
	v178 = int32(0)
	F_index_rescan(m, v155, v20+int32(112), v142, v178, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	F_index_endscan(m, v155)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L60
	}
L28:
	;
	v187 = v160
	goto L30
L29:
	;
	goto L27
L30:
	;
	v200 = F_index_getnext_slot(m, v155, int32(1), l3)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+28))
	m.T0[v219].(func(*base.Module, int32))(m, l3)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L44
	}
L32:
	;
	if v200 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if l1 == v27 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L31
L35:
	;
	v217 = v160
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v187 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v210 = F_palloc0(m, v207<<(uint(int32(2))%32))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v212 = v187
	goto L40
L40:
	;
	v213 = F_tuples_equal(m, l3, l2, v212)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v212 = v210
	goto L40
L42:
	;
	if v213 == int32(0) {
		v187 = v212
		goto L30
	} else {
		goto L43
	}
L43:
	;
	v217 = v212
	goto L34
L44:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
	if v222 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v224 = v222
	goto L47
L46:
	;
	v224 = v223
	goto L47
L47:
	;
	if v224 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v225 = int32(0)
	F_XactLockTableWait(m, v224, v225, v225, v225)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v230 = F_GetLatestSnapshot(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L52
	}
L51:
	;
	v160 = v217
	goto L26
L52:
	;
	F_PushActiveSnapshot(m, v230)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_RelationFindReplTupleByIndex[0]))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	goto L54
L54:
	;
	v238 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v241 = int32(0)
	v244 = v20 + int32(20)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+104))
	v247 = m.T0[v246].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3+int32(28), v236, l3, v238, int32(3), v241, v241, v244)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v251 = F_should_refetch_tuple(m, v247, v244)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v251 != 0 {
		v160 = v217
		goto L26
	} else {
		goto L59
	}
L59:
	;
	goto L29
L60:
	;
	F_relation_close(m, v23, int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	m.G0 = v20 + int32(1648)
	return v200
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v80
	F_errmsg_internal(m, int32(_a_F_RelationFindReplTupleByIndex_0), v20)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_RelationFindReplTupleByIndex_1), int32(101), int32(_a_F_RelationFindReplTupleByIndex_2))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationFindReplTupleSeq(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v20 = F_palloc0(m, v17<<(uint(int32(2))%32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(4)
	v30 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v36 = m.T0[v35].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v14+int32(24), v30, v30, v30, int32(449))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v39 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L6
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L38
	}
L6:
	;
	v52 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+188))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	m.T0[v59].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v36, v52, v52, v52, v52, v52)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+188))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	m.T0[v135].(func(*base.Module, int32))(m, v36)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L36
	}
L8:
	;
	goto L10
L9:
	;
	goto L7
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v74
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_RelationFindReplTupleSeq[0]))
	if v77 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+32))
	m.T0[v97].(func(*base.Module, int32, int32))(m, l2, v39)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationFindReplTupleSeq[1])))
	if v79&int32(1) == int32(0) {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+188))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+20))
	v88 = m.T0[v87].(func(*base.Module, int32, int32, int32) int32)(m, v36, int32(1), v39)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v88 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v92 = F_tuples_equal(m, v39, l1, v20)
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
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L11
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
	goto L6
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
	v125 = m.T0[v124].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l2+int32(28), v114, l2, v116, int32(3), v119, v119, v122)
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
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L9
L36:
	;
	F_ExecDropSingleTupleTableSlot(m, v39)
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
	return v88
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
	F_errfinish(m, int32(_a_F_RelationFindReplTupleSeq_1), int32(1034), int32(_a_F_RelationFindReplTupleSeq_2))
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
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int64
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int64
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int64
	_ = v767
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v824 int32
	_ = v824
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v847 int32
	_ = v847
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v871 int32
	_ = v871
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1126 int32
	_ = v1126
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1172 int32
	_ = v1172
	var v1181 int32
	_ = v1181
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int64
	_ = v1209
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1247 int32
	_ = v1247
	var v1249 int32
	_ = v1249
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1316 int32
	_ = v1316
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
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L33
	} else {
		goto L381
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
	v1286 = m.ExcPending
	if v1286 != 0 {
		goto L33
	} else {
		goto L377
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
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l2^int32(-1))<<(uint(int32(6))%32))+16))
	v66 = v57
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l2<<(uint(int32(6))%32)+int32(-64))+16))
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v77^int32(-1))<<(uint(int32(6))%32))+16))
	v98 = v89
	goto L22
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91+v77<<(uint(int32(6))%32)+int32(-64))+16))
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
	*(*int32)(unsafe.Add(mBase, uint32(v1256)+16)) = v1255
	m.G0 = v27 + int32(368)
	return v1261
L43:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+72))
	if v1241 != 0 {
		goto L374
	} else {
		goto L375
	}
L44:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1205 != 0 {
		v1255 = v1003
		v1256 = v1205
		v1261 = v782
		goto L42
	} else {
		goto L371
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
	v1139 = int32(4)
	v1140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021)+14)))
	v1141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021)+12)))
	v1142 = v1140 - v1141
	if v1142 <= v1139 {
		goto L352
	} else {
		goto L353
	}
L47:
	;
	v594 = int32(1)
	if v135|base.B2i32(l4 != v134) != 0 {
		goto L193
	} else {
		goto L194
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
	F_LockBuffer(m, v377, int32(2))
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
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206+(v202^int32(-1))<<(uint(int32(6))%32))+16))
	v221 = v212
	goto L63
L65:
	;
	goto L66
L66:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+v202<<(uint(int32(6))%32)+int32(-64))+16))
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
	F_LockBuffer(m, l2, int32(2))
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
	F_LockBuffer(m, v322, int32(2))
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
		v1255 = v170
		v1256 = v525
		v1261 = v378
		goto L42
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	F_LockBuffer(m, v378, int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L33
	} else {
		goto L169
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
	v1217 = v170
	v1218 = v533
	v1223 = v378
	goto L43
L169:
	;
	if l2 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if l4 == int32(0) {
		goto L178
	} else {
		goto L179
	}
L171:
	;
	if v170 == v68 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	F_ReleaseBuffer(m, v378)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L33
	} else {
		goto L176
	}
L174:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L33
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	goto L170
L177:
	;
	if v567 != int32(-1) {
		v170 = v567
		goto L49
	} else {
		goto L190
	}
L178:
	;
	if v74 != 0 {
		goto L47
	} else {
		goto L188
	}
L179:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v547 == int32(-1) {
		goto L178
	} else {
		goto L180
	}
L180:
	;
	if v74 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_RecordPageWithFreeSpace(m, l0, v170, v523)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L33
	} else {
		goto L184
	}
L182:
	;
	v555 = v547
	goto L183
L183:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(v556) <= base.Ui32(v555) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v555 = v554
	goto L183
L185:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = int64(-1)
	v567 = v555
	goto L177
L186:
	;
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v555 + int32(1)
	v567 = v555
	goto L177
L188:
	;
	v564 = F_RecordAndGetPageWithFreeSpace(m, l0, v170, v523, v107)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L33
	} else {
		goto L189
	}
L189:
	;
	v567 = v564
	goto L177
L190:
	;
	goto L50
L191:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v718
	v720 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v720
	v722 = m.G0
	v724 = v722 - int32(32)
	m.G0 = v724
	v727 = v27 - int32(-64)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	if v728 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L192:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(v694) < base.Ui32(v674) {
		goto L217
	} else {
		goto L218
	}
L193:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v596 != 0 {
		v674 = v129
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v684 = v594
	goto L195
L195:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v684
	v712 = v684
	v715 = v594
	v717 = int32(0)
	goto L191
L196:
	;
	if l4 != 0 {
		goto L192
	} else {
		goto L213
	}
L197:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v597 != 0 {
		v674 = v129
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v598 = m.G0
	v599 = int32(16)
	v600 = v598 - v599
	m.G0 = v600
	v602 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v600))) = v602
	v604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v600)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v600)+4)) = v604
	v609 = m.G0
	v611 = v609 - v599
	m.G0 = v611
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v613-int32(3))&int32(255)) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	m.G0 = v600 + int32(16)
	v674 = (v647 + int32(1)) * v129
	goto L196
L200:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[4]))
	v622 = F_get_hash_value(m, v621, v600)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L33
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L33
	} else {
		goto L210
	}
L203:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[5]))
	v632 = v625 + v622&int32(15)<<(uint(int32(7))%32) + int32(_a_F_RelationGetBufferForTuple_9)
	v634 = F_LWLockAcquire(m, v632, int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L33
	} else {
		goto L204
	}
L204:
	;
	v637 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[4]))
	v641 = F_hash_search_with_hash_value(m, v637, v600, v622, int32(0), v611+int32(15))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L33
	} else {
		goto L205
	}
L205:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v611)+15)))
	if v643 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v641)+84))
	v647 = v646
	goto L208
L207:
	;
	v647 = int32(0)
	goto L208
L208:
	;
	F_LWLockRelease(m, v632)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L33
	} else {
		goto L209
	}
L209:
	;
	m.G0 = v611 + int32(16)
	goto L199
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v611))) = v613
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_10), v611)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L33
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_11), int32(_a_F_RelationGetBufferForTuple_12), int32(_a_F_RelationGetBufferForTuple_13))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L33
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
	v678 = int32(64)
	if base.Ui32(v678) <= base.Ui32(v674) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v681 = v678
	goto L216
L215:
	;
	v681 = v674
	goto L216
L216:
	;
	v684 = v681
	goto L195
L217:
	;
	v696 = v674
	goto L219
L218:
	;
	v696 = v694
	goto L219
L219:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v696) {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	v699 = int32(64)
	goto L222
L221:
	;
	v699 = v696
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v699
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v701 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	F_ReleaseBuffer(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L33
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l0
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v712 = v699
	v715 = v129
	v717 = v709
	goto L191
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(0)
	goto L225
L227:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	if v732 != 0 {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	goto L229
L229:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+8)) = v765
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v767
	v776 = F_ExtendBufferedRelCommon(m, v724, int32(0), v717, int32(8), v712, int32(-1), v27+int32(112), v27+int32(108))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L33
	} else {
		goto L238
	}
L230:
	;
	v758 = v732
	goto L232
L231:
	;
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v731)+20))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v731)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+24)) = v734
	v736 = *(*int64)(unsafe.Add(mBase, uint32(v731)))
	*(*int64)(unsafe.Add(mBase, uint32(v724)+16)) = v736
	v740 = F_smgropen(m, v724+int32(16), v733)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L33
	} else {
		goto L233
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v758
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v731)+48))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+8)) = uint8(v761)
	goto L229
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v740
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v740)+72))
	if v744 != 0 {
		goto L235
	} else {
		goto L236
	}
L234:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	v758 = v756
	goto L232
L235:
	;
	v752 = v744
	goto L237
L236:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v740)+76))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v740)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+4)) = v746
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v740)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v740)+72))
	v752 = v750
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v740)+72)) = v752 + int32(1)
	goto L234
L238:
	;
	m.G0 = v724 + int32(32)
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	v783 = int32(0)
	v784 = base.B2i32(v783 <= v782)
	if v784 == v783 {
		goto L241
	} else {
		goto L242
	}
L239:
	;
	v958 = v781 + v776 - int32(1)
	if v74|base.B2i32(base.Ui32(v936) <= base.Ui32(v715)) == int32(0) {
		goto L275
	} else {
		goto L276
	}
L240:
	;
	v803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v802)+14)))
	if v803 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L241:
	;
	v788 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v788+(v782^int32(-1))<<(uint(int32(2))%32))))
	v802 = v794
	goto L240
L242:
	;
	goto L243
L243:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v802 = v796 + v782<<(uint(int32(13))%32) + int32(-8192)
	goto L240
L244:
	;
	v806 = int32(_a_F_RelationGetBufferForTuple_5)
	v807 = int32(0)
	if v807|(v802&int32(3)|int32(1)) == v807 {
		goto L249
	} else {
		goto L250
	}
L245:
	;
	goto L246
L246:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L33
	} else {
		goto L272
	}
L247:
	;
	F_MarkBufferDirty(m, v782)
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L33
	} else {
		goto L258
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+10)) = int32(_a_F_RelationGetBufferForTuple_6)
	v847 = int32(_a_F_RelationGetBufferForTuple_7)
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+18)) = uint16(v847)
	v853 = int32(_a_F_RelationGetBufferForTuple_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)) = uint16(v853)
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+14)) = uint16(v853)
	goto L247
L249:
	;
	goto L252
L250:
	;
	goto L251
L251:
	;
	goto L257
L252:
	;
	v824 = v802 + v806
	v826 = v802 + int32(4)
	if base.Ui32(v826) < base.Ui32(v824) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v828 = v824
	goto L255
L254:
	;
	v828 = v826
	goto L255
L255:
	;
	v833 = (v802^int32(-1)+v828)&int32(-4) + int32(4)
	if v833 == int32(0) {
		goto L248
	} else {
		goto L256
	}
L256:
	;
	base.MemoryFill(m, v802, int32(0), v833)
	goto L248
L257:
	;
	base.MemoryFill(m, v802, int32(0), v806)
	goto L248
L258:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v860 = v135 & base.B2i32(base.Ui32(v715) < base.Ui32(v858))
	if v860 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	F_LockBuffer(m, v782, int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L33
	} else {
		goto L262
	}
L260:
	;
	v865 = v858
	goto L261
L261:
	;
	v867 = int32(1)
	if base.Ui32(v865) <= base.Ui32(v867) {
		v936 = v865
		goto L239
	} else {
		goto L263
	}
L262:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v865 = v864
	goto L261
L263:
	;
	v871 = v867
	goto L264
L264:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(112)+v871<<(uint(int32(2))%32))))
	F_ReleaseBuffer(m, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L33
	} else {
		goto L266
	}
L265:
	;
	v936 = v912
	goto L239
L266:
	;
	if v74|base.B2i32(base.Ui32(v871) < base.Ui32(v715)) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_RecordPageWithFreeSpace(m, l0, v871+v776, int32(_a_F_RelationGetBufferForTuple_14))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L33
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	v911 = v871 + int32(1)
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	if base.Ui32(v911) < base.Ui32(v912) {
		v871 = v911
		goto L264
	} else {
		goto L271
	}
L270:
	;
	goto L269
L271:
	;
	goto L265
L272:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v776
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v918 + int32(4)
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_15), v27+int32(48))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L33
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(361), int32(_a_F_RelationGetBufferForTuple_17))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L33
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v776+v715, v958)
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L33
	} else {
		goto L278
	}
L276:
	;
	goto L277
L277:
	;
	if l4 != 0 {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	goto L277
L279:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v969 = base.B2i32(base.Ui32(int32(1)) < base.Ui32(v967))
	if base.Ui32(int32(1)) < base.Ui32(v967) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	if v782 < int32(0) {
		goto L290
	} else {
		goto L291
	}
L282:
	;
	v970 = v958
	goto L284
L283:
	;
	v970 = int32(-1)
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v970
	if base.Ui32(int32(1)) < base.Ui32(v967) {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v975 = v776 + int32(1)
	goto L287
L286:
	;
	v975 = int32(-1)
	goto L287
L287:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v975
	F_IncrBufferRefCount(m, v782)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L33
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v782
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v980 + v981
	goto L281
L289:
	;
	if v784 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L290:
	;
	v988 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[0]))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v988+(v782^int32(-1))<<(uint(int32(6))%32))+16))
	v1003 = v994
	goto L289
L291:
	;
	goto L292
L292:
	;
	v996 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[1]))
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v996+v782<<(uint(int32(6))%32)+int32(-64))+16))
	v1003 = v1002
	goto L289
L293:
	;
	if v131 == int32(0) {
		goto L301
	} else {
		goto L302
	}
L294:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[2]))
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v1007+(v782^int32(-1))<<(uint(int32(2))%32))))
	v1021 = v1013
	goto L293
L295:
	;
	goto L296
L296:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetBufferForTuple[3]))
	v1021 = v1015 + v782<<(uint(int32(13))%32) + int32(-8192)
	goto L293
L297:
	;
	goto L46
L298:
	;
	v1058 = F_GetVisibilityMapPins(m, l0, l2, v782, v68, v1003, l6, l5)
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L33
	} else {
		goto L324
	}
L299:
	;
	F_LockBuffer(m, v782, int32(2))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L33
	} else {
		goto L323
	}
L300:
	;
	F_LockBuffer(m, l2, int32(2))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L33
	} else {
		goto L322
	}
L301:
	;
	if v860 != 0 {
		goto L314
	} else {
		goto L315
	}
L302:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v1024 == int32(0) {
		goto L304
	} else {
		goto L305
	}
L303:
	;
	if v1032 != 0 {
		goto L301
	} else {
		goto L307
	}
L304:
	;
	v1032 = int32(0)
	goto L303
L305:
	;
	goto L306
L306:
	;
	v1028 = F_BufferGetBlockNumber(m, v1024)
	mBase = m.M
	v1030 = base.I32_div_u_s(v1003, int32(_a_F_RelationGetBufferForTuple_18))
	v1032 = base.B2i32(v1028 == v1030)
	goto L303
L307:
	;
	if v860 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	F_LockBuffer(m, v782, int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L33
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	F_visibilitymap_pin(m, l0, v1003, l5)
	mBase = m.M
	v1039 = m.ExcPending
	if v1039 != 0 {
		goto L33
	} else {
		goto L312
	}
L311:
	;
	goto L310
L312:
	;
	if l2 != 0 {
		goto L300
	} else {
		goto L313
	}
L313:
	;
	goto L299
L314:
	;
	if l2 != 0 {
		goto L300
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	if l2 == int32(0) {
		goto L297
	} else {
		goto L318
	}
L317:
	;
	goto L299
L318:
	;
	v1043 = F_ConditionalLockBuffer(m, l2)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L33
	} else {
		goto L319
	}
L319:
	;
	if v1043 != 0 {
		v1057 = int32(0)
		goto L298
	} else {
		goto L320
	}
L320:
	;
	F_LockBuffer(m, v782, int32(0))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L33
	} else {
		goto L321
	}
L321:
	;
	goto L300
L322:
	;
	goto L299
L323:
	;
	v1057 = int32(1)
	goto L298
L324:
	;
	v1063 = int32(4)
	v1064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021)+14)))
	v1065 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021)+12)))
	v1066 = v1064 - v1065
	if v1066 <= v1063 {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	if base.Ui32(v32) <= base.Ui32(v1126) {
		goto L44
	} else {
		goto L344
	}
L326:
	;
	v1069 = v1063
	goto L328
L327:
	;
	v1069 = v1066
	goto L328
L328:
	;
	v1071 = v1069 - int32(4)
	if v1071 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1126 = int32(0)
	goto L325
L330:
	;
	goto L331
L331:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1065) {
		goto L333
	} else {
		goto L334
	}
L332:
	;
	v1126 = v1071
	goto L325
L333:
	;
	v1082 = int32(base.Ui32(v1065+int32(_a_F_RelationGetBufferForTuple_3)) >> (uint(int32(2)) % 32))
	goto L335
L334:
	;
	v1082 = int32(0)
	goto L335
L335:
	;
	if base.Ui32(v1082&int32(_a_F_RelationGetBufferForTuple_8)) < base.Ui32(int32(291)) {
		goto L332
	} else {
		goto L336
	}
L336:
	;
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+10)))
	if v1087&int32(1) == int32(0) {
		goto L337
	} else {
		goto L338
	}
L337:
	;
	v1126 = int32(0)
	goto L325
L338:
	;
	goto L339
L339:
	;
	v1096 = int32(1)
	goto L340
L340:
	;
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021+int32(20)+v1096&int32(_a_F_RelationGetBufferForTuple_8)<<(uint(int32(2))%32))+1)))
	if v1105&int32(384) == int32(0) {
		goto L332
	} else {
		goto L342
	}
L341:
	;
	v1126 = int32(0)
	goto L325
L342:
	;
	v1111 = v1096 + int32(1)
	v1112 = int32(_a_F_RelationGetBufferForTuple_8)
	if base.Ui32(v1111&v1112) <= base.Ui32(v1082&v1112) {
		v1096 = v1111
		goto L340
	} else {
		goto L343
	}
L343:
	;
	goto L341
L344:
	;
	if v1058|v1057 != int32(1) {
		goto L1
	} else {
		goto L345
	}
L345:
	;
	if l2 != 0 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L33
	} else {
		goto L349
	}
L347:
	;
	goto L348
L348:
	;
	F_UnlockReleaseBuffer(m, v782)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L33
	} else {
		goto L350
	}
L349:
	;
	goto L348
L350:
	;
	v144 = v1003
	goto L45
L351:
	;
	if base.Ui32(v1202) < base.Ui32(v32) {
		goto L1
	} else {
		goto L370
	}
L352:
	;
	v1145 = v1139
	goto L354
L353:
	;
	v1145 = v1142
	goto L354
L354:
	;
	v1147 = v1145 - int32(4)
	if v1147 == int32(0) {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1202 = int32(0)
	goto L351
L356:
	;
	goto L357
L357:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1141) {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	v1202 = v1147
	goto L351
L359:
	;
	v1158 = int32(base.Ui32(v1141+int32(_a_F_RelationGetBufferForTuple_3)) >> (uint(int32(2)) % 32))
	goto L361
L360:
	;
	v1158 = int32(0)
	goto L361
L361:
	;
	if base.Ui32(v1158&int32(_a_F_RelationGetBufferForTuple_8)) < base.Ui32(int32(291)) {
		goto L358
	} else {
		goto L362
	}
L362:
	;
	v1163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1021)+10)))
	if v1163&int32(1) == int32(0) {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1202 = int32(0)
	goto L351
L364:
	;
	goto L365
L365:
	;
	v1172 = int32(1)
	goto L366
L366:
	;
	v1181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1021+int32(20)+v1172&int32(_a_F_RelationGetBufferForTuple_8)<<(uint(int32(2))%32))+1)))
	if v1181&int32(384) == int32(0) {
		goto L358
	} else {
		goto L368
	}
L367:
	;
	v1202 = int32(0)
	goto L351
L368:
	;
	v1187 = v1172 + int32(1)
	v1188 = int32(_a_F_RelationGetBufferForTuple_8)
	if base.Ui32(v1187&v1188) <= base.Ui32(v1158&v1188) {
		v1172 = v1187
		goto L366
	} else {
		goto L369
	}
L369:
	;
	goto L367
L370:
	;
	goto L44
L371:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1207
	v1209 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v1209
	v1213 = F_smgropen(m, v27+int32(32), v1206)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L33
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1213
	v1217 = v1003
	v1218 = v1213
	v1223 = v782
	goto L43
L373:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1255 = v1217
	v1256 = v1253
	v1261 = v1223
	goto L42
L374:
	;
	v1249 = v1241
	goto L376
L375:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+76))
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1242)+4)) = v1243
	v1245 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1243))) = v1245
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v1218)+72))
	v1249 = v1247
	goto L376
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+72)) = v1249 + int32(1)
	goto L373
L377:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L33
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(_a_F_RelationGetBufferForTuple_19)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v32
	F_errmsg(m, int32(_a_F_RelationGetBufferForTuple_20), v27)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L33
	} else {
		goto L379
	}
L379:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(536), int32(_a_F_RelationGetBufferForTuple_21))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L33
	} else {
		goto L380
	}
L380:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v32
	F_errmsg_internal(m, int32(_a_F_RelationGetBufferForTuple_22), v27+int32(16))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L33
	} else {
		goto L382
	}
L382:
	;
	F_errfinish(m, int32(_a_F_RelationGetBufferForTuple_16), int32(870), int32(_a_F_RelationGetBufferForTuple_21))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L33
	} else {
		goto L383
	}
L383:
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
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
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
	return v121
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
	v121 = v13
	goto L1
L7:
	;
	v121 = int32(0)
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
		v121 = v21
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
	v75 = v28
	goto L14
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+18)))
	if base.Ui32(v82&int32(2044)) <= base.Ui32(int32(19)) {
		goto L21
	} else {
		goto L22
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
	base.MemoryCopy(m, v50+(v38+v51<<(uint(int32(4))%32))+int32(20), v50+int32(_a_F_RelationGetIndexExpressions_1), v49)
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
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[0])) = v38
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v32
	v75 = v38
	goto L14
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
	v98 = F_text_to_cstring(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L26
	}
L21:
	;
	v90 = F_getmissingattr(m, v75, int32(20), v10+int32(15))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v95 = F_fastgetattr_3(m, v26, int32(20), v75, v10+int32(15))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	v97 = v90
	goto L20
L25:
	;
	v97 = v95
	goto L20
L26:
	;
	v100 = F_stringToNode(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v105 = F_eval_const_expressions(m, int32(0), v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	F_fix_opfuncids(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v109 = int32(_a_F_RelationGetIndexExpressions_0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v112
	v114 = F_copyObjectImpl(m, v105)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+228)) = v114
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexExpressions[1])) = v110
	v121 = v105
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
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
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
	return v126
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
	v126 = v13
	goto L1
L7:
	;
	v126 = int32(0)
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
		v126 = v21
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
	v75 = v28
	goto L14
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+18)))
	if base.Ui32(v82&int32(2047)) <= base.Ui32(int32(20)) {
		goto L21
	} else {
		goto L22
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
	base.MemoryCopy(m, v50+(v38+v51<<(uint(int32(4))%32))+int32(20), v50+int32(_a_F_RelationGetIndexPredicate_1), v49)
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
	*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[0])) = v38
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v32
	v75 = v38
	goto L14
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
	v98 = F_text_to_cstring(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L26
	}
L21:
	;
	v90 = F_getmissingattr(m, v75, int32(21), v10+int32(15))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v95 = F_fastgetattr_3(m, v26, int32(21), v75, v10+int32(15))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L25
	}
L24:
	;
	v97 = v90
	goto L20
L25:
	;
	v97 = v95
	goto L20
L26:
	;
	v100 = F_stringToNode(m, v98)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v98)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v105 = F_eval_const_expressions(m, int32(0), v100)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v108 = F_canonicalize_qual(m, v105, int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v110 = F_make_ands_implicit(m, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_fix_opfuncids(m, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v114 = int32(_a_F_RelationGetIndexPredicate_0)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1]))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v117
	v119 = F_copyObjectImpl(m, v110)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+232)) = v119
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexPredicate[1])) = v115
	v126 = v110
	goto L1
}
func F_RelationGetPartitionDesc(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v142 int32
	_ = v142
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
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
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v637 int32
	_ = v637
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v770 int32
	_ = v770
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v817 int32
	_ = v817
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v965 int32
	_ = v965
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v986 int32
	_ = v986
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1030 int32
	_ = v1030
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1056 int32
	_ = v1056
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1095 int32
	_ = v1095
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1144 int32
	_ = v1144
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1202 int32
	_ = v1202
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1213 int32
	_ = v1213
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1239 int32
	_ = v1239
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1272 int32
	_ = v1272
	var v1279 int32
	_ = v1279
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1320 int32
	_ = v1320
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1367 int32
	_ = v1367
	var v1378 int32
	_ = v1378
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1403 int32
	_ = v1403
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1429 int32
	_ = v1429
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1450 int32
	_ = v1450
	var v1477 int32
	_ = v1477
	var v1489 int32
	_ = v1489
	var v1508 int32
	_ = v1508
	var v1518 int32
	_ = v1518
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1562 int32
	_ = v1562
	var v1563 int32
	_ = v1563
	var v1566 int32
	_ = v1566
	var v1574 int32
	_ = v1574
	var v1592 int32
	_ = v1592
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1610 int32
	_ = v1610
	var v1614 int32
	_ = v1614
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1643 int32
	_ = v1643
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1658 int32
	_ = v1658
	var v1659 int32
	_ = v1659
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1675 int32
	_ = v1675
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
	var v1733 int32
	_ = v1733
	var v1754 int32
	_ = v1754
	var v1758 int32
	_ = v1758
	var v1762 int32
	_ = v1762
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1771 int32
	_ = v1771
	var v1780 int32
	_ = v1780
	var v1784 int32
	_ = v1784
	var v1789 int32
	_ = v1789
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1802 int32
	_ = v1802
	var v1807 int32
	_ = v1807
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1828 int32
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1863 int32
	_ = v1863
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1884 int32
	_ = v1884
	var v1895 int32
	_ = v1895
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1919 int32
	_ = v1919
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1963 int32
	_ = v1963
	var v1977 int32
	_ = v1977
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1990 int32
	_ = v1990
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2007 int32
	_ = v2007
	var v2009 int32
	_ = v2009
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2020 int32
	_ = v2020
	var v2025 int32
	_ = v2025
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2036 int32
	_ = v2036
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2085 int32
	_ = v2085
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2114 int32
	_ = v2114
	var v2162 int32
	_ = v2162
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2169 int32
	_ = v2169
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2181 int32
	_ = v2181
	var v2196 int32
	_ = v2196
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2236 int32
	_ = v2236
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2240 int32
	_ = v2240
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2271 int32
	_ = v2271
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2291 int32
	_ = v2291
	var v2309 int32
	_ = v2309
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2320 int32
	_ = v2320
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2365 int32
	_ = v2365
	var v2366 int32
	_ = v2366
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2379 int32
	_ = v2379
	var v2387 int32
	_ = v2387
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2410 int32
	_ = v2410
	var v2417 int32
	_ = v2417
	var v2437 int32
	_ = v2437
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2464 int32
	_ = v2464
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2469 int32
	_ = v2469
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2478 int32
	_ = v2478
	var v2484 int32
	_ = v2484
	var v2513 int32
	_ = v2513
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2546 int32
	_ = v2546
	var v2548 int32
	_ = v2548
	var v2550 int32
	_ = v2550
	var v2559 int32
	_ = v2559
	var v2560 int32
	_ = v2560
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2590 int32
	_ = v2590
	var v2592 int32
	_ = v2592
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2603 int32
	_ = v2603
	var v2604 int32
	_ = v2604
	var v2606 int32
	_ = v2606
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2616 int32
	_ = v2616
	var v2618 int32
	_ = v2618
	var v2620 int32
	_ = v2620
	var v2622 int32
	_ = v2622
	var v2623 int32
	_ = v2623
	var v2625 int32
	_ = v2625
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2635 int32
	_ = v2635
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2659 int32
	_ = v2659
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2675 int32
	_ = v2675
	var v2680 int32
	_ = v2680
	var v2682 int32
	_ = v2682
	var v2685 int32
	_ = v2685
	var v2715 int32
	_ = v2715
	var v2742 int32
	_ = v2742
	var v2749 int32
	_ = v2749
	var v2754 int32
	_ = v2754
	var v2756 int32
	_ = v2756
	var v2760 int32
	_ = v2760
	var v2764 int32
	_ = v2764
	var v2765 int32
	_ = v2765
	var v2770 int32
	_ = v2770
	var v2777 int32
	_ = v2777
	var v2791 int32
	_ = v2791
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2800 int32
	_ = v2800
	var v2805 int32
	_ = v2805
	var v2812 int32
	_ = v2812
	var v2826 int32
	_ = v2826
	var v2829 int32
	_ = v2829
	var v2833 int32
	_ = v2833
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2843 int32
	_ = v2843
	var v2850 int32
	_ = v2850
	var v2862 int32
	_ = v2862
	var v2866 int32
	_ = v2866
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2876 int32
	_ = v2876
	var v2883 int32
	_ = v2883
	var v2900 int32
	_ = v2900
	var v2930 int32
	_ = v2930
	var v2934 int32
	_ = v2934
	var v2939 int32
	_ = v2939
	var v2943 int32
	_ = v2943
	var v2949 int32
	_ = v2949
	var v2954 int32
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2965 int32
	_ = v2965
	var v2970 int32
	_ = v2970
	v3 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(128)
	m.G0 = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v31 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2958 = m.ExcPending
	if v2958 != 0 {
		goto L17
	} else {
		goto L467
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L17
	} else {
		goto L464
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2930 = m.ExcPending
	if v2930 != 0 {
		goto L17
	} else {
		goto L461
	}
L4:
	;
	m.G0 = v29 + int32(128)
	return v2900
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
		v2900 = v31
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	if v36&int32(1) == int32(0) {
		v2900 = v31
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L9
L9:
	;
	if v42 != int32(0) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	v2900 = v45
	goto L4
L11:
	;
	v67 = F_RelationGetPartitionKey(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L17
	} else {
		goto L20
	}
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v48 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L14
L14:
	;
	if base.B2i32(v52 != int32(0)) == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	goto L16
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v61 = F_XidInMVCCSnapshot(m, v60, v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v61 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v2900 = v65
	goto L4
L20:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v69
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v79 = F_find_inheritance_children_extended(m, v73, l1, v69, v29+int32(127), v29+int32(120))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L17
	} else {
		goto L24
	}
L21:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[1]))
	v2245 = F_AllocSetContextCreateInternal(m, v2240, int32(_a_F_RelationGetPartitionDesc_0), int32(0), int32(1024), int32(_a_F_RelationGetPartitionDesc_1))
	mBase = m.M
	v2246 = m.ExcPending
	if v2246 != 0 {
		goto L17
	} else {
		goto L318
	}
L22:
	;
	v2223 = v2196
	v2227 = v3
	v2236 = v2209
	v2237 = v2210
	v2238 = int32(0)
	goto L21
L23:
	;
	v91 = v3
	v93 = v79 + int32(4)
	v95 = v79
	v98 = v81
	goto L29
L24:
	;
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if int32(0) < v81 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	v84 = v3
	goto L27
L27:
	;
	v2196 = v84
	v2209 = v3
	v2210 = v3
	goto L22
L28:
	;
	v84 = v81
	goto L27
L29:
	;
	v114 = v98 << (uint(int32(2)) % 32)
	v115 = F_palloc(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v2196 = int32(0)
	v2209 = v115
	v2210 = v117
	goto L22
L31:
	;
	v117 = F_palloc(m, v98)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v119 = F_palloc(m, v114)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v121 = int32(0)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v121 < v122 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L17
	} else {
		goto L312
	}
L35:
	;
	v142 = v121
	goto L38
L36:
	;
	goto L37
L37:
	;
	v343 = int32(0)
	v348 = v29 - int32(-64)
	v350 = v98 << (uint(int32(2)) % 32)
	v351 = F_palloc(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L17
	} else {
		goto L95
	}
L38:
	;
	v153 = v142 << (uint(int32(2)) % 32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v95)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153+v154)))
	v157 = F_SearchSysCache1(m, int32(57), v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L17
	} else {
		goto L42
	}
L39:
	;
	goto L37
L40:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	if v291 != int32(98) {
		goto L2
	} else {
		goto L86
	}
L41:
	;
	v181 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L17
	} else {
		goto L53
	}
L42:
	;
	if v157 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v165 = F_SysCacheGetAttr(m, int32(57), v157, int32(34), v29-int32(-64))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+64)))
	if v167 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ReleaseCatCache(m, v157)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L17
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v172 = F_text_to_cstring(m, v165)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L17
	} else {
		goto L49
	}
L48:
	;
	goto L41
L49:
	;
	v174 = F_stringToNode(m, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	F_ReleaseCatCache(m, v157)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	if v174 != 0 {
		v286 = v174
		goto L40
	} else {
		goto L52
	}
L52:
	;
	goto L41
L53:
	;
	v184 = v29 - int32(-64)
	F_ScanKeyInit(m, v184, int32(1), int32(3), int32(184), v156)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v190 = int32(0)
	v192 = int32(1)
	v195 = F_systable_beginscan(m, v181, int32(2662), v192, v190, v192, v184)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	F_systable_endscan(m, v195)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L17
	} else {
		goto L82
	}
L56:
	;
	v197 = F_systable_getnext(m, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	if v197 == int32(0) {
		v271 = v190
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v181)+52))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+18)))
	if base.Ui32(v203&int32(2046)) <= base.Ui32(int32(33)) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)))
	if v265 != 0 {
		v271 = int32(0)
		goto L55
	} else {
		goto L79
	}
L60:
	;
	v211 = F_getmissingattr(m, v201, int32(34), v29+int32(63))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L17
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v213 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v213)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+20)))
	if v215&int32(1) == v213 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v263 = v211
	goto L59
L64:
	;
	v259 = F_nocachegetattr(m, v197, int32(34), v201)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L17
	} else {
		goto L78
	}
L65:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v201)+548))
	if v220 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+27)))
	if v251&int32(2) != 0 {
		goto L64
	} else {
		goto L77
	}
L68:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+22)))
	v225 = v202 + v223 + v220
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+554)))
	if v226 != int32(1) {
		v263 = v225
		goto L59
	} else {
		goto L69
	}
L69:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+552)))
	switch v229 - int32(1) {
	case 0:
		goto L73
	case 1:
		goto L72
	default:
		goto L70
	case 3:
		goto L71
	}
L70:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L17
	} else {
		goto L74
	}
L71:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	v263 = v234
	goto L59
L72:
	;
	v233 = int32(*(*int16)(unsafe.Add(mBase, uint32(v225))))
	v263 = v233
	goto L59
L73:
	;
	v232 = int32(*(*int8)(unsafe.Add(mBase, uint32(v225))))
	v263 = v232
	goto L59
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = base.I32_extend16_s(v229)
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_2), v29+int32(48))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L17
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_3), int32(70), int32(_a_F_RelationGetPartitionDesc_4))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v254)
	v263 = int32(0)
	goto L59
L78:
	;
	v263 = v259
	goto L59
L79:
	;
	v266 = F_text_to_cstring(m, v263)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v268 = F_stringToNode(m, v266)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	v271 = v268
	goto L55
L82:
	;
	F_relation_close(m, v181, int32(1))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	if base.B2i32(v271 == int32(0))&(v91^int32(-1)) != 0 {
		goto L34
	} else {
		goto L84
	}
L84:
	;
	if v271 == int32(0) {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v286 = v271
	goto L40
L86:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+5)))
	if v294 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v298 = F_get_default_partition_oid(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L17
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153+v115))) = v156
	v305 = F_get_rel_relkind(m, v156)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	if v298 != v156 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v142+v117))) = uint8(base.B2i32(v305 != int32(112)))
	*(*int32)(unsafe.Add(mBase, uint32(v153+v119))) = v286
	v313 = v142 + int32(1)
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v313 < v314 {
		v142 = v313
		goto L38
	} else {
		goto L93
	}
L93:
	;
	goto L39
L94:
	;
	v2223 = v98
	v2227 = int32(1)
	v2236 = v115
	v2237 = v117
	v2238 = v2162
	goto L21
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v348))) = v351
	if v98 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v502 = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	switch v503 - int32(104) {
	case 0:
		goto L117
	default:
		v2114 = v502
		goto L108
	case 4:
		goto L116
	case 10:
		goto L115
	}
L97:
	;
	v357 = v98 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v98) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v364 = v343
	v366 = v343
	goto L101
L99:
	;
	v415 = v343
	goto L100
L100:
	;
	v441 = v415
	v442 = v343
	goto L105
L101:
	;
	v389 = v364 << (uint(int32(2)) % 32)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v392 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v389+v390))) = v392
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v394+v389)+4)) = v392
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v398+v389)+8)) = v392
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v402+v389)+12)) = v392
	v406 = int32(4)
	v407 = v364 + v406
	v409 = v366 + v406
	if v409 != v98&int32(2147483644) {
		v364 = v407
		v366 = v409
		goto L101
	} else {
		goto L103
	}
L102:
	;
	if v357 == int32(0) {
		goto L96
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	v415 = v407
	goto L100
L105:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v465+v441<<(uint(int32(2))%32)))) = int32(-1)
	v471 = int32(1)
	v474 = v442 + v471
	if v474 != v357 {
		v441 = v441 + v471
		v442 = v474
		goto L105
	} else {
		goto L107
	}
L106:
	;
	goto L96
L107:
	;
	goto L106
L108:
	;
	v2162 = v2114
	goto L94
L109:
	;
	F_qsort_arg(m, v1842, v1841, int32(8), int32(908), v67)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L17
	} else {
		goto L273
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1824 = m.ExcPending
	if v1824 != 0 {
		goto L17
	} else {
		goto L270
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L17
	} else {
		goto L267
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L17
	} else {
		goto L264
	}
L113:
	;
	v1792 = F_palloc(m, int32(0))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L17
	} else {
		goto L263
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L17
	} else {
		goto L260
	}
L115:
	;
	v1250 = F_palloc0(m, int32(36))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L17
	} else {
		goto L194
	}
L116:
	;
	v912 = F_palloc0(m, int32(36))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L17
	} else {
		goto L156
	}
L117:
	;
	v507 = F_palloc0(m, int32(36))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L17
	} else {
		goto L118
	}
L118:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v507)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v507))) = v509
	v514 = v98 * int32(12)
	v515 = F_palloc(m, v514)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L17
	} else {
		goto L119
	}
L119:
	;
	if int32(0) < v98 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v522 = int32(0)
	goto L123
L121:
	;
	goto L122
L122:
	;
	F_pg_qsort(m, v515, v98, int32(12), int32(906))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L17
	} else {
		goto L127
	}
L123:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v119+v522<<(uint(int32(2))%32))))
	v550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+4)))
	if v550 != int32(104) {
		goto L114
	} else {
		goto L125
	}
L124:
	;
	goto L122
L125:
	;
	v555 = v515 + v522*int32(12)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v549)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = v556
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v549)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v555)+8)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v558
	v562 = v522 + int32(1)
	if v562 != v98 {
		v522 = v562
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v514+v515-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v507)+4)) = v98
	v599 = F_palloc0(m, v350)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L17
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507)+20)) = v597
	*(*int64)(unsafe.Add(mBase, uint32(v507)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v507)+8)) = v599
	v607 = F_palloc(m, v597<<(uint(int32(2))%32))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L17
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v507)+24)) = v607
	if v597 <= int32(0) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v762 = F_palloc(m, v98<<(uint(int32(3))%32))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L17
	} else {
		goto L142
	}
L131:
	;
	v613 = v597 & int32(3)
	v614 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v597) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v622 = v614
	v637 = int32(0)
	goto L135
L133:
	;
	v673 = v614
	goto L134
L134:
	;
	v699 = v673
	v701 = v502
	goto L139
L135:
	;
	v647 = v622 << (uint(int32(2)) % 32)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	v650 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v647+v648))) = v650
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v652+v647)+4)) = v650
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v656+v647)+8)) = v650
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v660+v647)+12)) = v650
	v664 = int32(4)
	v665 = v622 + v664
	v667 = v637 + v664
	if v667 != v597&int32(2147483644) {
		v622 = v665
		v637 = v667
		goto L135
	} else {
		goto L137
	}
L136:
	;
	if v613 == int32(0) {
		goto L130
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v673 = v665
	goto L134
L139:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v723+v699<<(uint(int32(2))%32)))) = int32(-1)
	v729 = int32(1)
	v732 = v701 + v729
	if v732 != v613 {
		v699 = v699 + v729
		v701 = v732
		goto L139
	} else {
		goto L141
	}
L140:
	;
	goto L130
L141:
	;
	goto L140
L142:
	;
	if int32(0) < v98 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v770 = int32(0)
	goto L146
L144:
	;
	goto L145
L145:
	;
	F_pfree(m, v515)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L17
	} else {
		goto L155
	}
L146:
	;
	v795 = v515 + v770*int32(12)
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v795)))
	v799 = v770 << (uint(int32(2)) % 32)
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v799+v800))) = v762 + v770<<(uint(int32(3))%32)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v806+v799)))
	*(*int32)(unsafe.Add(mBase, uint32(v808))) = v797
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v507)+8))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v810+v799)))
	*(*int32)(unsafe.Add(mBase, uint32(v812)+4)) = v796
	if v796 < v597 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	goto L145
L148:
	;
	v817 = v796
	goto L151
L149:
	;
	goto L150
L150:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v795)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v874+v875<<(uint(int32(2))%32)))) = v770
	v881 = v770 + int32(1)
	if v881 != v98 {
		v770 = v881
		goto L146
	} else {
		goto L154
	}
L151:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v507)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v841+v817<<(uint(int32(2))%32)))) = v770
	v846 = v817 + v797
	if v846 < v597 {
		v817 = v846
		goto L151
	} else {
		goto L153
	}
L152:
	;
	goto L150
L153:
	;
	goto L152
L154:
	;
	goto L147
L155:
	;
	v2162 = v507
	goto L94
L156:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v912)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v912))) = v914
	if v98 <= int32(0) {
		goto L113
	} else {
		goto L157
	}
L157:
	;
	v927 = v3
	v928 = v343
	goto L158
L158:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v119+v928<<(uint(int32(2))%32))))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)+16))
	if v950 == int32(0) {
		v1095 = v927
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v1120 = F_palloc(m, v1095<<(uint(int32(3))%32))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L17
	} else {
		goto L175
	}
L160:
	;
	v1115 = v928 + int32(1)
	if v1115 != v98 {
		v927 = v1095
		v928 = v1115
		goto L158
	} else {
		goto L174
	}
L161:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v950)+4))
	if v953 <= int32(0) {
		v1095 = v927
		goto L160
	} else {
		goto L162
	}
L162:
	;
	v957 = v953 & int32(3)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v950)+12))
	if base.Ui32(v953) < base.Ui32(int32(4)) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v1051 = v1025
	v1052 = int32(0)
	v1056 = v1030
	goto L171
L164:
	;
	v1025 = int32(0)
	v1030 = v927
	goto L163
L165:
	;
	goto L166
L166:
	;
	v965 = int32(0)
	v969 = v965
	v974 = v927
	v986 = v965
	goto L167
L167:
	;
	v995 = v958 + v969<<(uint(int32(2))%32)
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v995)))
	v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v996)+24)))
	v998 = int32(1)
	v1001 = *(*int32)(unsafe.Add(mBase, uint32(v995)+4))
	v1002 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1001)+24)))
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v995)+8))
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+24)))
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v995)+12))
	v1012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1011)+24)))
	v1015 = v974 + (v997 ^ v998) + (v1002 ^ v998) + (v1007 ^ v998) + (v1012 ^ v998)
	v1016 = int32(4)
	v1017 = v969 + v1016
	v1019 = v986 + v1016
	if v1019 != v953&int32(2147483644) {
		v969 = v1017
		v974 = v1015
		v986 = v1019
		goto L167
	} else {
		goto L169
	}
L168:
	;
	if v957 == int32(0) {
		v1095 = v1015
		goto L160
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v1025 = v1017
	v1030 = v1015
	goto L163
L171:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v958+v1051<<(uint(int32(2))%32))))
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1078)+24)))
	v1080 = int32(1)
	v1082 = v1056 + (v1079 ^ v1080)
	v1086 = v1052 + v1080
	if v1086 != v957 {
		v1051 = v1051 + v1080
		v1052 = v1086
		v1056 = v1082
		goto L171
	} else {
		goto L173
	}
L172:
	;
	v1095 = v1082
	goto L160
L173:
	;
	goto L172
L174:
	;
	goto L159
L175:
	;
	v1123 = int32(-1)
	v1128 = int32(0)
	v1130 = v1123
	v1131 = v1123
	v1144 = int32(0)
	goto L176
L176:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v119+v1128<<(uint(int32(2))%32))))
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154)+4)))
	if v1155 != int32(108) {
		goto L112
	} else {
		goto L178
	}
L177:
	;
	v1839 = v1225
	v1840 = v1226
	v1841 = v1095
	v1842 = v1120
	goto L109
L178:
	;
	v1158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1154)+5)))
	if v1158 != 0 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	v1247 = v1128 + int32(1)
	if v1247 != v98 {
		v1128 = v1247
		v1130 = v1225
		v1131 = v1226
		v1144 = v1239
		goto L176
	} else {
		goto L193
	}
L180:
	;
	v1225 = v1130
	v1226 = v1128
	v1239 = v1144
	goto L179
L181:
	;
	goto L182
L182:
	;
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1154)+16))
	if v1159 == int32(0) {
		v1225 = v1130
		v1226 = v1131
		v1239 = v1144
		goto L179
	} else {
		goto L183
	}
L183:
	;
	v1162 = int32(0)
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+4))
	if v1163 <= v1162 {
		v1225 = v1130
		v1226 = v1131
		v1239 = v1144
		goto L179
	} else {
		goto L184
	}
L184:
	;
	v1168 = v1162
	v1171 = v1130
	v1182 = v1163
	v1185 = v1144
	goto L185
L185:
	;
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+12))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1192+v1168<<(uint(int32(2))%32))))
	v1197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196)+24)))
	if v1197 == int32(0) {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1225 = v1213
	v1226 = v1131
	v1239 = v1216
	goto L179
L187:
	;
	v1218 = v1168 + int32(1)
	if v1218 < v1215 {
		v1168 = v1218
		v1171 = v1213
		v1182 = v1215
		v1185 = v1216
		goto L185
	} else {
		goto L192
	}
L188:
	;
	v1202 = v1120 + v1185<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1202))) = v1128
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1196)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1202)+4)) = v1204
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+4))
	v1213 = v1171
	v1215 = v1208
	v1216 = v1185 + int32(1)
	goto L187
L189:
	;
	goto L190
L190:
	;
	if base.B2i32(v1171 == int32(-1)) == int32(0) {
		goto L111
	} else {
		goto L191
	}
L191:
	;
	v1213 = v1128
	v1215 = v1182
	v1216 = v1185
	goto L187
L192:
	;
	goto L186
L193:
	;
	goto L177
L194:
	;
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v1250)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1250))) = v1252
	v1256 = int32(-1)
	v1259 = F_palloc0(m, v98<<(uint(int32(3))%32))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L17
	} else {
		goto L195
	}
L195:
	;
	if v98 <= int32(0) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	F_qsort_arg(m, v1259, v1335, int32(4), int32(907), v67)
	mBase = m.M
	v1351 = m.ExcPending
	if v1351 != 0 {
		goto L17
	} else {
		goto L210
	}
L197:
	;
	v1328 = v1256
	v1335 = int32(0)
	goto L196
L198:
	;
	goto L199
L199:
	;
	v1264 = int32(0)
	v1268 = v1264
	v1272 = v1256
	v1279 = v1264
	goto L200
L200:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v119+v1268<<(uint(int32(2))%32))))
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295)+4)))
	if v1296 != int32(114) {
		goto L110
	} else {
		goto L202
	}
L201:
	;
	v1328 = v1317
	v1335 = v1318
	goto L196
L202:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1295)+5)))
	if v1299 != 0 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1320 = v1268 + int32(1)
	if v1320 != v98 {
		v1268 = v1320
		v1272 = v1317
		v1279 = v1318
		goto L200
	} else {
		goto L209
	}
L204:
	;
	v1317 = v1268
	v1318 = v1279
	goto L203
L205:
	;
	goto L206
L206:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+20))
	v1302 = F_make_one_partition_rbound(m, v67, v1268, v1300, int32(1))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L17
	} else {
		goto L207
	}
L207:
	;
	v1306 = v1259 + v1279<<(uint(int32(2))%32)
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1295)+24))
	v1309 = F_make_one_partition_rbound(m, v67, v1268, v1307, int32(0))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L17
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1306)+4)) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v1306))) = v1302
	v1317 = v1272
	v1318 = v1279 + int32(2)
	goto L203
L209:
	;
	goto L201
L210:
	;
	v1354 = F_palloc(m, v1335<<(uint(int32(2))%32))
	mBase = m.M
	v1355 = m.ExcPending
	if v1355 != 0 {
		goto L17
	} else {
		goto L211
	}
L211:
	;
	if int32(0) < v1335 {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1362 = int32(0)
	v1367 = v343
	v1378 = v3
	goto L215
L213:
	;
	v1518 = v343
	goto L214
L214:
	;
	F_pfree(m, v1259)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L17
	} else {
		goto L231
	}
L215:
	;
	v1387 = v1259 + v1378<<(uint(int32(2))%32)
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1389 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v1389 <= int32(0) {
		v1489 = v1367
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v1518 = v1489
	goto L214
L217:
	;
	v1508 = v1378 + int32(1)
	if v1508 != v1335 {
		v1362 = v1388
		v1367 = v1489
		v1378 = v1508
		goto L215
	} else {
		goto L230
	}
L218:
	;
	if v1362 != 0 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v1403 = int32(0)
	goto L222
L220:
	;
	v1477 = v1388
	goto L221
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1354+v1367<<(uint(int32(2))%32)))) = v1477
	v1489 = v1367 + int32(1)
	goto L217
L222:
	;
	v1423 = v1403 << (uint(int32(2)) % 32)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+8))
	v1426 = *(*int32)(unsafe.Add(mBase, uint32(v1423+v1424)))
	v1427 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+8))
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1427+v1423)))
	if v1426 != v1429 {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	v1450 = *(*int32)(unsafe.Add(mBase, uint32(v1387)))
	v1477 = v1450
	goto L221
L224:
	;
	goto L223
L225:
	;
	if v1426 != 0 {
		v1489 = v1367
		goto L217
	} else {
		goto L226
	}
L226:
	;
	v1431 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1435+v1423)))
	v1438 = *(*int32)(unsafe.Add(mBase, uint32(v1388)+4))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1438+v1423)))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1362)+4))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1441+v1423)))
	v1444 = F_FunctionCall2Coll(m, v1431+v1403*int32(28), v1437, v1440, v1443)
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L17
	} else {
		goto L227
	}
L227:
	;
	if v1444 != 0 {
		goto L224
	} else {
		goto L228
	}
L228:
	;
	v1447 = v1403 + int32(1)
	v1448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v1447 < v1448 {
		v1403 = v1447
		goto L222
	} else {
		goto L229
	}
L229:
	;
	v1489 = v1367
	goto L217
L230:
	;
	goto L216
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+4)) = v1518
	v1540 = v1518 << (uint(int32(2)) % 32)
	v1541 = F_palloc0(m, v1540)
	mBase = m.M
	v1542 = m.ExcPending
	if v1542 != 0 {
		goto L17
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+8)) = v1541
	v1544 = F_palloc(m, v1540)
	mBase = m.M
	v1545 = m.ExcPending
	if v1545 != 0 {
		goto L17
	} else {
		goto L233
	}
L233:
	;
	v1547 = v1518 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+20)) = v1547
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+12)) = v1544
	v1554 = F_palloc(m, v1547<<(uint(int32(2))%32))
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L17
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+24)) = v1554
	v1557 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	v1558 = v1540 * v1557
	v1559 = F_palloc(m, v1558)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L17
	} else {
		goto L235
	}
L235:
	;
	v1561 = F_palloc(m, v1558)
	mBase = m.M
	v1562 = m.ExcPending
	if v1562 != 0 {
		goto L17
	} else {
		goto L236
	}
L236:
	;
	v1563 = int32(0)
	if v1563 < v1518 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1566 = int32(0)
	v1574 = v1566
	v1592 = v3
	goto L240
L238:
	;
	v1733 = v1563
	v1754 = v3
	goto L239
L239:
	;
	F_pfree(m, v1354)
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L17
	} else {
		goto L256
	}
L240:
	;
	v1595 = int32(2)
	v1596 = v1574 << (uint(v1595) % 32)
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	v1601 = v1574 * v1557 << (uint(v1595) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1596+v1597))) = v1559 + v1601
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1604+v1596))) = v1601 + v1561
	if base.B2i32(v1557 <= v1566) == int32(0) {
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1733 = v1518
	v1754 = v1724
	goto L239
L242:
	;
	v1610 = v1596 + v1354
	v1614 = int32(0)
	goto L245
L243:
	;
	goto L244
L244:
	;
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v1596+v1354)))
	v1706 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1705)+12)))
	if v1706 != 0 {
		v1721 = int32(-1)
		v1724 = v1592
		goto L252
	} else {
		goto L253
	}
L245:
	;
	v1639 = v1614 << (uint(int32(2)) % 32)
	v1640 = *(*int32)(unsafe.Add(mBase, uint32(v1610)))
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+8))
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1639+v1641)))
	if v1643 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	goto L244
L247:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1640)+4))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1646+v1639)))
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v1651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1649+v1614))))
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v1656 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1652+v1614<<(uint(int32(1))%32)))))
	v1657 = F_datumCopy(m, v1648, v1651, v1656)
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L17
	} else {
		goto L250
	}
L248:
	;
	v1668 = v1643
	goto L249
L249:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+12))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1669+v1596)))
	*(*int32)(unsafe.Add(mBase, uint32(v1671+v1639))) = v1668
	v1675 = v1614 + int32(1)
	if v1675 != v1557 {
		v1614 = v1675
		goto L245
	} else {
		goto L251
	}
L250:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+8))
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1659+v1596)))
	*(*int32)(unsafe.Add(mBase, uint32(v1661+v1639))) = v1657
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1610)))
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v1664)+8))
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v1665+v1639)))
	v1668 = v1667
	goto L249
L251:
	;
	goto L246
L252:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1725+v1596))) = v1721
	v1729 = v1574 + int32(1)
	if v1729 != v1518 {
		v1574 = v1729
		v1592 = v1724
		goto L240
	} else {
		goto L255
	}
L253:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v1705)))
	v1709 = v1707 << (uint(int32(2)) % 32)
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1711 = v1709 + v1710
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1711)))
	if v1712 != int32(-1) {
		v1721 = v1712
		v1724 = v1592
		goto L252
	} else {
		goto L254
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1711))) = v1592
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1718+v1709)))
	v1721 = v1720
	v1724 = v1592 + int32(1)
	goto L252
L255:
	;
	goto L241
L256:
	;
	if v1328 != int32(-1) {
		goto L257
	} else {
		goto L258
	}
L257:
	;
	v1762 = v1328 << (uint(int32(2)) % 32)
	v1763 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v1762+v1763))) = v1754
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1768 = *(*int32)(unsafe.Add(mBase, uint32(v1766+v1762)))
	*(*int32)(unsafe.Add(mBase, uint32(v1250)+32)) = v1768
	goto L259
L258:
	;
	goto L259
L259:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1771+v1733<<(uint(int32(2))%32)))) = int32(-1)
	v2114 = v1250
	goto L108
L260:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L17
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(372), int32(_a_F_RelationGetPartitionDesc_7))
	mBase = m.M
	v1789 = m.ExcPending
	if v1789 != 0 {
		goto L17
	} else {
		goto L262
	}
L262:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L263:
	;
	v1839 = int32(-1)
	v1840 = int32(-1)
	v1841 = v3
	v1842 = v1792
	goto L109
L264:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L17
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(493), int32(_a_F_RelationGetPartitionDesc_8))
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L17
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L267:
	;
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_9), int32(0))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L17
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(523), int32(_a_F_RelationGetPartitionDesc_8))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_5), int32(0))
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L17
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_6), int32(713), int32(_a_F_RelationGetPartitionDesc_10))
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v912)+4)) = v1841
	v1866 = v1841 << (uint(int32(2)) % 32)
	v1867 = F_palloc0(m, v1866)
	mBase = m.M
	v1868 = m.ExcPending
	if v1868 != 0 {
		goto L17
	} else {
		goto L274
	}
L274:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+20)) = v1841
	*(*int64)(unsafe.Add(mBase, uint32(v912)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v912)+8)) = v1867
	v1873 = F_palloc(m, v1866)
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L17
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+24)) = v1873
	v1876 = F_palloc(m, v1866)
	mBase = m.M
	v1877 = m.ExcPending
	if v1877 != 0 {
		goto L17
	} else {
		goto L276
	}
L276:
	;
	v1878 = int32(0)
	if v1878 < v1841 {
		goto L277
	} else {
		goto L278
	}
L277:
	;
	v1884 = int32(0)
	v1895 = v1878
	goto L280
L278:
	;
	v1963 = v1878
	goto L279
L279:
	;
	F_pfree(m, v1842)
	mBase = m.M
	v1977 = m.ExcPending
	if v1977 != 0 {
		goto L17
	} else {
		goto L287
	}
L280:
	;
	v1910 = v1842 + v1884<<(uint(int32(3))%32)
	v1911 = *(*int32)(unsafe.Add(mBase, uint32(v1910)))
	v1913 = v1884 << (uint(int32(2)) % 32)
	v1914 = *(*int32)(unsafe.Add(mBase, uint32(v912)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1913+v1914))) = v1913 + v1876
	v1918 = *(*int32)(unsafe.Add(mBase, uint32(v1910)+4))
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v1920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1919))))
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v1922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1921))))
	v1923 = F_datumCopy(m, v1918, v1920, v1922)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L17
	} else {
		goto L282
	}
L281:
	;
	v1963 = v1943
	goto L279
L282:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v912)+8))
	v1927 = *(*int32)(unsafe.Add(mBase, uint32(v1925+v1913)))
	*(*int32)(unsafe.Add(mBase, uint32(v1927))) = v1923
	v1930 = v1911 << (uint(int32(2)) % 32)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1932 = v1930 + v1931
	v1933 = *(*int32)(unsafe.Add(mBase, uint32(v1932)))
	if v1933 == int32(-1) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1932))) = v1895
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1941 = *(*int32)(unsafe.Add(mBase, uint32(v1939+v1930)))
	v1942 = v1941
	v1943 = v1895 + int32(1)
	goto L285
L284:
	;
	v1942 = v1933
	v1943 = v1895
	goto L285
L285:
	;
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v912)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1944+v1913))) = v1942
	v1948 = v1884 + int32(1)
	if v1948 != v1841 {
		v1884 = v1948
		v1895 = v1943
		goto L280
	} else {
		goto L286
	}
L286:
	;
	goto L281
L287:
	;
	if v1839 != int32(-1) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1981 = v1839 << (uint(int32(2)) % 32)
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1983 = v1981 + v1982
	v1984 = *(*int32)(unsafe.Add(mBase, uint32(v1983)))
	if v1984 == int32(-1) {
		goto L291
	} else {
		goto L292
	}
L289:
	;
	v1999 = v1963
	goto L290
L290:
	;
	if v1840 != int32(-1) {
		goto L294
	} else {
		goto L295
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1983))) = v1963
	v1990 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v1992 = *(*int32)(unsafe.Add(mBase, uint32(v1990+v1981)))
	v1993 = v1963 + int32(1)
	v1994 = v1992
	goto L293
L292:
	;
	v1993 = v1963
	v1994 = v1984
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+28)) = v1994
	v1999 = v1993
	goto L290
L294:
	;
	v2003 = v1840 << (uint(int32(2)) % 32)
	v2004 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v2003+v2004))) = v1999
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v2007+v2003)))
	*(*int32)(unsafe.Add(mBase, uint32(v912)+32)) = v2009
	goto L296
L295:
	;
	goto L296
L296:
	;
	if v98 < int32(2) {
		v2114 = v912
		goto L108
	} else {
		goto L297
	}
L297:
	;
	v2014 = int32(-1)
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v912)+4))
	v2016 = *(*int32)(unsafe.Add(mBase, uint32(v912)+28))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v912)+32))
	if v2015+base.B2i32(v2016 != v2014)+base.B2i32(v2020 != v2014) == v98 {
		v2085 = v2020
		goto L298
	} else {
		goto L299
	}
L298:
	;
	if v2085 == int32(-1) {
		v2114 = v912
		goto L108
	} else {
		goto L310
	}
L299:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v912)+20))
	if v2025 <= int32(0) {
		v2085 = v2020
		goto L298
	} else {
		goto L300
	}
L300:
	;
	v2031 = v2014
	v2032 = v2025
	v2036 = int32(0)
	goto L301
L301:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v912)+24))
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2055+v2036<<(uint(int32(2))%32))))
	if v2031 <= v2059 {
		goto L304
	} else {
		goto L305
	}
L302:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v912)+32))
	v2085 = v2077
	goto L298
L303:
	;
	v2075 = v2036 + int32(1)
	if v2075 < v2072 {
		v2031 = v2059
		v2032 = v2072
		v2036 = v2075
		goto L301
	} else {
		goto L309
	}
L304:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(v912)+28))
	if base.B2i32(v2061 == int32(-1))|base.B2i32(v2059 != v2061) != 0 {
		v2072 = v2032
		goto L303
	} else {
		goto L307
	}
L305:
	;
	goto L306
L306:
	;
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v912)+16))
	v2068 = F_bms_add_member(m, v2067, v2059)
	mBase = m.M
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L17
	} else {
		goto L308
	}
L307:
	;
	goto L306
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+16)) = v2068
	v2071 = *(*int32)(unsafe.Add(mBase, uint32(v912)+20))
	v2072 = v2071
	goto L303
L309:
	;
	goto L302
L310:
	;
	v2106 = *(*int32)(unsafe.Add(mBase, uint32(v912)+16))
	v2107 = F_bms_add_member(m, v2106, v2085)
	mBase = m.M
	v2108 = m.ExcPending
	if v2108 != 0 {
		goto L17
	} else {
		goto L311
	}
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v912)+16)) = v2107
	v2162 = v912
	goto L94
L312:
	;
	v2165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)) = uint8(v2165)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v2165
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2175 = F_find_inheritance_children_extended(m, v2169, l1, v2165, v29+int32(127), v29+int32(120))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L17
	} else {
		goto L313
	}
L313:
	;
	if v2175 != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v2180 = int32(0)
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v2175)+4))
	if v2181 <= v2180 {
		v2223 = v2181
		v2227 = v3
		v2236 = v115
		v2237 = v117
		v2238 = v2180
		goto L21
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	goto L30
L317:
	;
	v91 = int32(1)
	v93 = v2175 + int32(4)
	v95 = v2175
	v98 = v2181
	goto L29
L318:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2250 = F_MemoryContextStrdup(m, v2245, v2247+int32(4))
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L17
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+36)) = v2250
	v2254 = F_MemoryContextAllocZero(m, v2245, int32(32))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L17
	} else {
		goto L320
	}
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2254))) = v2223
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2254)+4)) = uint8(v2257)
	if v2227 != 0 {
		goto L323
	} else {
		goto L324
	}
L321:
	;
	v2862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2862 != 0 {
		goto L441
	} else {
		goto L442
	}
L322:
	;
	v2829 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[2]))
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+16))
	if v2833 != v2829 {
		goto L425
	} else {
		goto L426
	}
L323:
	;
	v2261 = int32(_a_F_RelationGetPartitionDesc_11)
	v2262 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3])) = v2245
	v2267 = F_palloc(m, int32(36))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L17
	} else {
		goto L326
	}
L324:
	;
	v2742 = v2257
	goto L325
L325:
	;
	if base.B2i32(l1 == int32(0))|base.B2i32(v2742&int32(1) == int32(0)) != 0 {
		goto L322
	} else {
		goto L383
	}
L326:
	;
	v2269 = *(*int32)(unsafe.Add(mBase, uint32(v2238)))
	*(*int32)(unsafe.Add(mBase, uint32(v2267))) = v2269
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+4)) = v2271
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+20)) = v2273
	v2275 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	v2277 = v2271 << (uint(int32(2)) % 32)
	v2278 = F_palloc(m, v2277)
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L17
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+8)) = v2278
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+12))
	if v2281 != 0 {
		goto L329
	} else {
		goto L330
	}
L328:
	;
	v2365 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+16))
	v2366 = F_bms_copy(m, v2365)
	mBase = m.M
	v2367 = m.ExcPending
	if v2367 != 0 {
		goto L17
	} else {
		goto L341
	}
L329:
	;
	v2282 = F_palloc(m, v2277)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L17
	} else {
		goto L332
	}
L330:
	;
	goto L331
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+12)) = int32(0)
	goto L328
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+12)) = v2282
	v2286 = F_palloc(m, v2277*v2275)
	mBase = m.M
	v2287 = m.ExcPending
	if v2287 != 0 {
		goto L17
	} else {
		goto L333
	}
L333:
	;
	if v2271 <= int32(0) {
		goto L328
	} else {
		goto L334
	}
L334:
	;
	v2291 = v2275 << (uint(int32(2)) % 32)
	v2309 = int32(0)
	goto L335
L335:
	;
	v2318 = int32(2)
	v2319 = v2309 << (uint(v2318) % 32)
	v2320 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2319+v2320))) = v2286 + v2275*v2309<<(uint(v2318)%32)
	if v2291 != 0 {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	goto L328
L337:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+12))
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2327+v2319)))
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+12))
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2330+v2319)))
	base.MemoryCopy(m, v2329, v2332, v2291)
	goto L339
L338:
	;
	goto L339
L339:
	;
	v2335 = v2309 + int32(1)
	if v2335 != v2271 {
		v2309 = v2335
		goto L335
	} else {
		goto L340
	}
L340:
	;
	goto L336
L341:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+16)) = v2366
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v2372 = base.B2i32(v2370 == int32(104))
	if v2370 == int32(104) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v2373 = int32(2)
	goto L344
L343:
	;
	v2373 = v2275
	goto L344
L344:
	;
	v2375 = F_palloc(m, v2277*v2373)
	mBase = m.M
	v2376 = m.ExcPending
	if v2376 != 0 {
		goto L17
	} else {
		goto L345
	}
L345:
	;
	if int32(0) < v2271 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v2379 = int32(0)
	v2387 = v2379
	goto L349
L347:
	;
	goto L348
L348:
	;
	v2542 = v2273 << (uint(int32(2)) % 32)
	v2543 = F_palloc(m, v2542)
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L17
	} else {
		goto L368
	}
L349:
	;
	v2408 = int32(2)
	v2409 = v2387 << (uint(v2408) % 32)
	v2410 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2409+v2410))) = v2375 + v2387*v2373<<(uint(v2408)%32)
	v2417 = int32(0)
	if base.B2i32(v2373 <= v2379) == v2417 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	goto L348
L351:
	;
	v2437 = v2417
	goto L354
L352:
	;
	goto L353
L353:
	;
	v2513 = v2387 + int32(1)
	if v2513 != v2271 {
		v2387 = v2513
		goto L349
	} else {
		goto L367
	}
L354:
	;
	if v2370 == int32(104) {
		goto L357
	} else {
		goto L358
	}
L355:
	;
	goto L353
L356:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+12))
	if v2458 != 0 {
		goto L361
	} else {
		goto L362
	}
L357:
	;
	v2456 = int32(1)
	v2457 = int32(4)
	goto L356
L358:
	;
	goto L359
L359:
	;
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v2450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2448+v2437))))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v2455 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2451+v2437<<(uint(int32(1))%32)))))
	v2456 = v2450
	v2457 = v2455
	goto L356
L360:
	;
	v2484 = v2437 + int32(1)
	if v2484 != v2373 {
		v2437 = v2484
		goto L354
	} else {
		goto L366
	}
L361:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2458+v2409)))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2460+v2437<<(uint(int32(2))%32))))
	if v2464 != 0 {
		goto L360
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v2466 = v2437 << (uint(int32(2)) % 32)
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+8))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2467+v2409)))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2466+v2469)))
	v2474 = F_datumCopy(m, v2471, v2456&int32(1), v2457)
	mBase = m.M
	v2475 = m.ExcPending
	if v2475 != 0 {
		goto L17
	} else {
		goto L365
	}
L364:
	;
	goto L363
L365:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v2267)+8))
	v2478 = *(*int32)(unsafe.Add(mBase, uint32(v2476+v2409)))
	*(*int32)(unsafe.Add(mBase, uint32(v2478+v2466))) = v2474
	goto L360
L366:
	;
	goto L355
L367:
	;
	goto L350
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+24)) = v2543
	if v2542 != 0 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+24))
	base.MemoryCopy(m, v2543, v2546, v2542)
	goto L371
L370:
	;
	goto L371
L371:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+28)) = v2548
	v2550 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2267)+32)) = v2550
	*(*int32)(unsafe.Add(mBase, uint32(v2254)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2254)+20)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2254)+16)) = v2267
	v2559 = F_palloc(m, v2223<<(uint(int32(2))%32))
	mBase = m.M
	v2560 = m.ExcPending
	if v2560 != 0 {
		goto L17
	} else {
		goto L372
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2254)+8)) = v2559
	v2562 = F_palloc(m, v2223)
	mBase = m.M
	v2563 = m.ExcPending
	if v2563 != 0 {
		goto L17
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2254)+12)) = v2562
	if v2223 <= int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[3])) = v2262
	v2715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)))
	v2742 = v2715
	goto L325
L375:
	;
	v2567 = int32(0)
	if v2223 != int32(1) {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2590 = v2567
	v2592 = int32(0)
	goto L379
L377:
	;
	v2659 = v2567
	goto L378
L378:
	;
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+8))
	v2671 = int32(2)
	v2672 = v2659 << (uint(v2671) % 32)
	v2673 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2675 = *(*int32)(unsafe.Add(mBase, uint32(v2672+v2673)))
	v2680 = *(*int32)(unsafe.Add(mBase, uint32(v2672+v2236)))
	*(*int32)(unsafe.Add(mBase, uint32(v2670+v2675<<(uint(v2671)%32)))) = v2680
	v2682 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+12))
	v2685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2659+v2237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2675+v2682))) = uint8(v2685)
	goto L374
L379:
	;
	v2601 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+8))
	v2602 = int32(2)
	v2603 = v2590 << (uint(v2602) % 32)
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2606 = *(*int32)(unsafe.Add(mBase, uint32(v2603+v2604)))
	v2611 = *(*int32)(unsafe.Add(mBase, uint32(v2603+v2236)))
	*(*int32)(unsafe.Add(mBase, uint32(v2601+v2606<<(uint(v2602)%32)))) = v2611
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+12))
	v2616 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2590+v2237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2606+v2613))) = uint8(v2616)
	v2618 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+8))
	v2620 = v2590 | int32(1)
	v2622 = v2620 << (uint(v2602) % 32)
	v2623 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2625 = *(*int32)(unsafe.Add(mBase, uint32(v2622+v2623)))
	v2630 = *(*int32)(unsafe.Add(mBase, uint32(v2622+v2236)))
	*(*int32)(unsafe.Add(mBase, uint32(v2618+v2625<<(uint(v2602)%32)))) = v2630
	v2632 = *(*int32)(unsafe.Add(mBase, uint32(v2254)+12))
	v2635 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2620+v2237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2625+v2632))) = uint8(v2635)
	v2638 = v2590 + v2602
	v2640 = v2592 + v2602
	if v2640 != v2223&int32(2147483646) {
		v2590 = v2638
		v2592 = v2640
		goto L379
	} else {
		goto L381
	}
L380:
	;
	if v2223&int32(1) == int32(0) {
		goto L374
	} else {
		goto L382
	}
L381:
	;
	goto L380
L382:
	;
	v2659 = v2638
	goto L378
L383:
	;
	v2749 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[0]))
	goto L384
L384:
	;
	if base.B2i32(v2749 != int32(0)) == int32(0) {
		goto L322
	} else {
		goto L385
	}
L385:
	;
	v2754 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v2756 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetPartitionDesc[2]))
	v2760 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+16))
	if v2760 != v2756 {
		goto L387
	} else {
		goto L388
	}
L386:
	;
	if v2754 == int32(0) {
		goto L321
	} else {
		goto L403
	}
L387:
	;
	if v2760 == int32(0) {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	goto L389
L389:
	;
	goto L386
L390:
	;
	if v2756 != 0 {
		goto L397
	} else {
		goto L398
	}
L391:
	;
	v2764 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+28))
	v2765 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+24))
	if v2765 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	if v2764 == int32(0) {
		goto L390
	} else {
		goto L396
	}
L393:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2765)+28)) = v2764
	goto L392
L394:
	;
	goto L395
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2760)+20)) = v2764
	goto L392
L396:
	;
	v2770 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2764)+24)) = v2770
	goto L390
L397:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+16)) = v2756
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2756)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+28)) = v2777
	if v2777 != 0 {
		goto L400
	} else {
		goto L401
	}
L398:
	;
	goto L399
L399:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2245)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+16)) = int32(0)
	goto L389
L400:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2777)+24)) = v2245
	goto L402
L401:
	;
	goto L402
L402:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2756)+20)) = v2245
	goto L386
L403:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2791 != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v2791)+16))
	if v2795 != v2245 {
		goto L408
	} else {
		goto L409
	}
L405:
	;
	goto L406
L406:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v2254
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2245
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v2826
	v2900 = v2254
	goto L4
L407:
	;
	goto L406
L408:
	;
	if v2795 == int32(0) {
		goto L411
	} else {
		goto L412
	}
L409:
	;
	goto L410
L410:
	;
	goto L407
L411:
	;
	if v2245 != 0 {
		goto L418
	} else {
		goto L419
	}
L412:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2791)+28))
	v2800 = *(*int32)(unsafe.Add(mBase, uint32(v2791)+24))
	if v2800 != 0 {
		goto L414
	} else {
		goto L415
	}
L413:
	;
	if v2799 == int32(0) {
		goto L411
	} else {
		goto L417
	}
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2800)+28)) = v2799
	goto L413
L415:
	;
	goto L416
L416:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+20)) = v2799
	goto L413
L417:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v2791)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2799)+24)) = v2805
	goto L411
L418:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2791)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2791)+16)) = v2245
	v2812 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2791)+28)) = v2812
	if v2812 != 0 {
		goto L421
	} else {
		goto L422
	}
L419:
	;
	goto L420
L420:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2791)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2791)+16)) = int32(0)
	goto L410
L421:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2812)+24)) = v2791
	goto L423
L422:
	;
	goto L423
L423:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+20)) = v2791
	goto L407
L424:
	;
	goto L321
L425:
	;
	if v2833 == int32(0) {
		goto L428
	} else {
		goto L429
	}
L426:
	;
	goto L427
L427:
	;
	goto L424
L428:
	;
	if v2829 != 0 {
		goto L435
	} else {
		goto L436
	}
L429:
	;
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+28))
	v2838 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+24))
	if v2838 != 0 {
		goto L431
	} else {
		goto L432
	}
L430:
	;
	if v2837 == int32(0) {
		goto L428
	} else {
		goto L434
	}
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2838)+28)) = v2837
	goto L430
L432:
	;
	goto L433
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2833)+20)) = v2837
	goto L430
L434:
	;
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2837)+24)) = v2843
	goto L428
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+16)) = v2829
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v2829)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+28)) = v2850
	if v2850 != 0 {
		goto L438
	} else {
		goto L439
	}
L436:
	;
	goto L437
L437:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2245)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+16)) = int32(0)
	goto L427
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2850)+24)) = v2245
	goto L440
L439:
	;
	goto L440
L440:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2829)+20)) = v2245
	goto L424
L441:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, uint32(v2862)+16))
	if v2866 != v2245 {
		goto L445
	} else {
		goto L446
	}
L442:
	;
	goto L443
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2254
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2245
	v2900 = v2254
	goto L4
L444:
	;
	goto L443
L445:
	;
	if v2866 == int32(0) {
		goto L448
	} else {
		goto L449
	}
L446:
	;
	goto L447
L447:
	;
	goto L444
L448:
	;
	if v2245 != 0 {
		goto L455
	} else {
		goto L456
	}
L449:
	;
	v2870 = *(*int32)(unsafe.Add(mBase, uint32(v2862)+28))
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(v2862)+24))
	if v2871 != 0 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	if v2870 == int32(0) {
		goto L448
	} else {
		goto L454
	}
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2871)+28)) = v2870
	goto L450
L452:
	;
	goto L453
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2866)+20)) = v2870
	goto L450
L454:
	;
	v2876 = *(*int32)(unsafe.Add(mBase, uint32(v2862)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2870)+24)) = v2876
	goto L448
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2862)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2862)+16)) = v2245
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2245)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2862)+28)) = v2883
	if v2883 != 0 {
		goto L458
	} else {
		goto L459
	}
L456:
	;
	goto L457
L457:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2862)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2862)+16)) = int32(0)
	goto L447
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2883)+24)) = v2862
	goto L460
L459:
	;
	goto L460
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2245)+20)) = v2862
	goto L444
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v156
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_12), v29)
	mBase = m.M
	v2934 = m.ExcPending
	if v2934 != 0 {
		goto L17
	} else {
		goto L462
	}
L462:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(280), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L17
	} else {
		goto L463
	}
L463:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v156
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_15), v29+int32(32))
	mBase = m.M
	v2949 = m.ExcPending
	if v2949 != 0 {
		goto L17
	} else {
		goto L465
	}
L465:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(282), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		goto L17
	} else {
		goto L466
	}
L466:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v156
	F_errmsg_internal(m, int32(_a_F_RelationGetPartitionDesc_16), v29+int32(16))
	mBase = m.M
	v2965 = m.ExcPending
	if v2965 != 0 {
		goto L17
	} else {
		goto L468
	}
L468:
	;
	F_errfinish(m, int32(_a_F_RelationGetPartitionDesc_13), int32(296), int32(_a_F_RelationGetPartitionDesc_14))
	mBase = m.M
	v2970 = m.ExcPending
	if v2970 != 0 {
		goto L17
	} else {
		goto L469
	}
L469:
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v47 = F_RelationBuildDesc(m, v45, int32(1))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				if v47 == int32(0) {
					v66 = v2
					m.G0 = v6 + int32(16)
					return v66
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
					F_ResourceOwnerEnlarge(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v55 + int32(1)
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[2]))
						if v60 != 0 {
							v62 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
							F_ResourceOwnerRemember(m, v62, v47, int32(_a_F_RelationIdGetRelation_0))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								v66 = v47
								m.G0 = v6 + int32(16)
								return v66
							}
						} else {
							v66 = v47
							m.G0 = v6 + int32(16)
							return v66
						}
					}
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			if v22 == int32(0) {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v47 = F_RelationBuildDesc(m, v45, int32(1))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v47 == int32(0) {
						v66 = v2
						m.G0 = v6 + int32(16)
						return v66
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
						F_ResourceOwnerEnlarge(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v55 + int32(1)
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[2]))
							if v60 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, _c_F_RelationIdGetRelation[1]))
								F_ResourceOwnerRemember(m, v62, v47, int32(_a_F_RelationIdGetRelation_0))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									v66 = v47
									m.G0 = v6 + int32(16)
									return v66
								}
							} else {
								v66 = v47
								m.G0 = v6 + int32(16)
								return v66
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
				if v25 != 0 {
					v66 = v2
					m.G0 = v6 + int32(16)
					return v66
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
							F_ResourceOwnerRemember(m, v37, v22, int32(_a_F_RelationIdGetRelation_0))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
								if v41 != 0 {
									v66 = v22
									m.G0 = v6 + int32(16)
									return v66
								} else {
									F_RelationRebuildRelation(m, v22)
									mBase = m.M
									v43 = m.ExcPending
									if v43 != 0 {
										return int32(0)
									} else {
										v66 = v22
										m.G0 = v6 + int32(16)
										return v66
									}
								}
							}
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+26)))
							if v41 != 0 {
								v66 = v22
								m.G0 = v6 + int32(16)
								return v66
							} else {
								F_RelationRebuildRelation(m, v22)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v66 = v22
									m.G0 = v6 + int32(16)
									return v66
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
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
			if v54 != 0 {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
				if v56 == int32(439) {
					v72 = int32(16)
					m.G0 = v52 + v72
					*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v54
					m.G0 = v7 + v72
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
						F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
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
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
					F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
				if v54 != 0 {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
					if v56 == int32(439) {
						v72 = int32(16)
						m.G0 = v52 + v72
						*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v54
						m.G0 = v7 + v72
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
							F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
						F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
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
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+84))
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
							if v54 != 0 {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
								if v56 == int32(439) {
									v72 = int32(16)
									m.G0 = v52 + v72
									*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v54
									m.G0 = v7 + v72
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
										F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
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
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v52))) = v49
									F_errmsg_internal(m, int32(_a_F_RelationInitTableAccessMethod_0), v52)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_1), int32(38), int32(_a_F_RelationInitTableAccessMethod_2))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
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
							F_errfinish(m, int32(_a_F_RelationInitTableAccessMethod_5), int32(1863), int32(_a_F_RelationInitTableAccessMethod_6))
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
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(l0^int32(-1))<<(uint(int32(6))%32))+16))
				v48 = v39
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _c_F_RelationPutHeapTuple[3]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+l0<<(uint(int32(6))%32)+int32(-64))+16))
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
			F_errstart_cold(m, int32(23), int32(0))
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
					F_errfinish(m, int32(_a_F_RelationPutHeapTuple_2), int32(65), int32(_a_F_RelationPutHeapTuple_3))
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
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
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
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
	v57 = v26 + v51<<(uint(int32(4))%32) + v38*int32(100)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+111)))
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
	v62 = F_pstrdup(m, v57+int32(24))
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
	v185 = F_palloc(m, v154<<(uint(int32(2))%32))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L5
	} else {
		goto L45
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v154
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v178
	goto L38
L40:
	;
	v170 = F_palloc0(m, v154<<(uint(int32(2))%32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v172 = int32(2)
	v176 = F_repalloc0(m, v165, v163<<(uint(v172)%32), v154<<(uint(v172)%32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v178 = v170
	goto L39
L44:
	;
	v178 = v176
	goto L39
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v185
	v188 = F_palloc(m, v154)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v188
	F_build_colinfo_names_hash(m, l2)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v193 = int32(0)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+8))
	if v195 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v197 = v196
	goto L50
L49:
	;
	v197 = v4
	goto L50
L50:
	;
	if int32(0) < v154 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v205 = int32(0)
	v208 = v193
	v213 = v4
	goto L54
L52:
	;
	v308 = v193
	v313 = v4
	goto L53
L53:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v315 != 0 {
		goto L78
	} else {
		goto L79
	}
L54:
	;
	v216 = v205 << (uint(int32(2)) % 32)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v159+v216)))
	if v218 == int32(0) {
		v294 = v208
		v297 = v213
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v308 = v294
	v313 = v297
	goto L53
L56:
	;
	v299 = v205 + int32(1)
	if v299 != v154 {
		v205 = v299
		v208 = v294
		v213 = v297
		goto L54
	} else {
		goto L77
	}
L57:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221+v216)))
	if v223 != 0 {
		v250 = v223
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v253+v208<<(uint(int32(2))%32)))) = v250
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v258+v208))) = uint8(base.B2i32(v197 <= v205))
	v263 = v208 + int32(1)
	if v213 != 0 {
		goto L67
	} else {
		goto L68
	}
L59:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v224 == int32(0) {
		v236 = v218
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v238 = F_make_colname_unique(m, v236, l0, l2)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L64
	}
L61:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+8))
	if v227 == int32(0) {
		v236 = v218
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v230 <= v205 {
		v236 = v218
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v232+v216)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v236 = v235
	goto L60
L64:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v240+v216))) = v238
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v243 == int32(0) {
		v250 = v238
		goto L58
	} else {
		goto L65
	}
L65:
	;
	v248 = F_hash_search(m, v243, v238, int32(1), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v250 = v238
	goto L58
L67:
	;
	v294 = v263
	v297 = int32(1)
	goto L56
L68:
	;
	goto L69
L69:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218))))
	if base.B2i32(v267 == int32(0))|base.B2i32(v267 != v270) != 0 {
		v288 = v267
		v289 = v270
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v294 = v263
	v297 = base.B2i32(v288-v289 != int32(0))
	goto L56
L71:
	;
	goto L70
L72:
	;
	v273 = v250
	v274 = v218
	goto L73
L73:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+1)))
	if v278 == int32(0) {
		v288 = v278
		v289 = v277
		goto L71
	} else {
		goto L75
	}
L74:
	;
	v288 = v278
	v289 = v277
	goto L71
L75:
	;
	v281 = int32(1)
	if v278 == v277 {
		v273 = v273 + v281
		v274 = v274 + v281
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L55
L78:
	;
	F_hash_destroy(m, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v308
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	switch v321 {
	case 0:
		goto L86
	default:
		goto L83
	case 3:
		goto L85
	case 4:
		goto L84
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(0)
	goto L80
L82:
	;
	m.G0 = v17 + int32(16)
	return
L83:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v327 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L84:
	;
	v325 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v325)
	goto L82
L85:
	;
	v323 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v323)
	goto L82
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v313)
	goto L82
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v313)
	goto L82
L88:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)+8))
	if v330 == int32(0) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v333 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v333)
	goto L82
}
