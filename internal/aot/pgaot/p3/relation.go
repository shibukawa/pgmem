package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetRelationPublications(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	v2 = int32(0)
	v9 = F_SearchSysCacheList(m, int32(53), int32(1), l0, v2, v2)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if int32(0) < v13 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = int32(0)
	v20 = v2
	goto L6
L4:
	;
	v38 = v2
	goto L5
L5:
	;
	F_ReleaseCatCacheList(m, v9)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(48)+v19<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27+v28)+4))
	v31 = F_lappend_oid(m, v20, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v31
	goto L5
L8:
	;
	v34 = v19 + int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if v34 < v35 {
		v19 = v34
		v20 = v31
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	return v38
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v7 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v5
	} else {
	}
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1073]))
	if v12 <= int32(31) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, _consts[1073])) = v12 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32))+uint32(_consts[1074]))) = v15
		return
	} else {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[1075])) = uint8(v26)
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
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
	v19 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v24 = F_AllocSetContextCreateInternal(m, v19, int32(156464), v2, int32(1024), int32(8192))
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
	goto L4
L4:
	;
	v35 = F_MemoryContextAlloc(m, v24, int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = int32(3)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v14+int32(-48), v39, v39, int32(184), v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v47 = F_table_open(m, int32(2618), int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+52))
	v51 = int32(1)
	v56 = F_systable_beginscan(m, v47, int32(2693), v51, int32(0), v51, v14+int32(-48))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v58 = F_systable_getnext(m, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v58 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = v58
	v66 = v2
	v67 = v27
	v69 = v35
	goto L13
L11:
	;
	v194 = v2
	v197 = v35
	goto L12
L12:
	;
	F_systable_endscan(m, v56)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L48
	}
L13:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+22)))
	v76 = F_MemoryContextAlloc(m, v24, int32(20))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	v194 = v185
	v197 = v179
	goto L12
L15:
	;
	v78 = v73 + v74
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v79
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v81 - int32(48)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+73)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+16)) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+74)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+17)) = uint8(v87)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+18)))
	if v90&int32(2040) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v106 = F_text_to_cstring(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L22
	}
L17:
	;
	v98 = F_getmissingattr(m, v49, int32(8), v14+int32(-49))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v103 = F_fastgetattr_3(m, v63, int32(8), v49, v14+int32(-49))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v105 = v98
	goto L16
L21:
	;
	v105 = v103
	goto L16
L22:
	;
	v108 = int32(4455216)
	v109 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	v112 = F_stringToNode(m, v106)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v112
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v109
	F_pfree(m, v106)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+18)))
	if base.Ui32(v120&int32(2047)) <= base.Ui32(int32(6)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v136 = F_text_to_cstring(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L31
	}
L26:
	;
	v128 = F_getmissingattr(m, v49, int32(7), v14+int32(-49))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v133 = F_fastgetattr_3(m, v63, int32(7), v49, v14+int32(-49))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v135 = v128
	goto L25
L30:
	;
	v135 = v133
	goto L25
L31:
	;
	v138 = int32(4455216)
	v139 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
	v142 = F_stringToNode(m, v136)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v142
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v139
	F_pfree(m, v136)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v151 != int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_setRuleCheckAsUser(m, v150, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L40
	}
L35:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v149)+80))
	v165 = v163
	goto L34
L36:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	if v154 != int32(118) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v157 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
	if v161 != 0 {
		v165 = int32(0)
		goto L34
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	F_setRuleCheckAsUser(m, v168, v165)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v67 <= v66 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v174 = F_repalloc(m, v69, v67<<(uint(int32(3))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v178 = v67
	v179 = v69
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+v66<<(uint(int32(2))%32)))) = v76
	v185 = v66 + int32(1)
	v186 = F_systable_getnext(m, v56)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	v178 = v67 << (uint(int32(1)) % 32)
	v179 = v174
	goto L44
L46:
	;
	if v186 != 0 {
		v63 = v186
		v66 = v185
		v67 = v178
		v69 = v179
		goto L13
	} else {
		goto L47
	}
L47:
	;
	goto L14
L48:
	;
	F_sequence_close(m, v47, int32(1))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v194 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	m.G0 = v16 - int32(-64)
	return
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = int64(0)
	F_MemoryContextDelete(m, v24)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v213 = F_MemoryContextAlloc(m, v24, int32(8))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = v194
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v213
	goto L50
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
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
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
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
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
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
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
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
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v31+int32(-48), int32(2), int32(3), int32(184), v43)
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
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L102
	}
L5:
	;
	v51 = int32(1)
	v56 = F_systable_beginscan(m, v48, int32(2701), v51, int32(0), v51, v31+int32(-48))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v58 = F_systable_getnext(m, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v58 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v66 = v58
	v68 = v2
	v69 = int32(16)
	v71 = v36
	goto L11
L9:
	;
	v375 = v2
	v378 = v36
	goto L10
L10:
	;
	F_systable_endscan(m, v56)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L77
	}
L11:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v66)+16))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
	v93 = v91 + v92
	if v69 <= v68 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v375 = v365
	v378 = v102
	goto L10
L13:
	;
	v97 = F_repalloc(m, v71, v69*int32(120))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v101 = v69
	v102 = v71
	goto L15
L15:
	;
	v105 = v102 + v68*int32(60)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v106
	v112 = F_DirectFunctionCall1Coll(m, int32(580), int32(0), v93+int32(12))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v101 = v69 << (uint(int32(1)) % 32)
	v102 = v97
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v93)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v115
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+12)) = uint16(v117)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+82)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+14)) = uint8(v119)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+83)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+15)) = uint8(v121)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v124 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+16)) = uint8(base.B2i32(v123 != v124))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v93)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+20)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v93)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+24)) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v93)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v105)+28)) = v131
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+32)) = uint8(v133)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+97)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105)+33)) = uint8(v135)
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+98)))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+34)) = uint16(v137)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v93)+116))
	*(*uint16)(unsafe.Add(mBase, uint32(v105)+36)) = uint16(v139)
	v142 = v139 << (uint(int32(16)) % 32)
	if v124 < v142 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(0) < base.I32_extend16_s(v160) {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v147 = F_palloc(m, int32(base.Ui32(v142)>>(uint(int32(15))%32)))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+40)) = int32(0)
	v160 = v137
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+40)) = v147
	v152 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+36)))
	v154 = v152 << (uint(int32(1)) % 32)
	if v154 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+34)))
	v160 = v157
	goto L18
L24:
	;
	v155 = F__emscripten_memcpy_bulkmem(m, v147, v93+int32(124), v154)
	mBase = m.M
	goto L26
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	v325 = int32(0)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v330 = F_fastgetattr_2(m, v66, int32(18), v327, v31+int32(-49))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L60
	}
L28:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v168 = F_fastgetattr_2(m, v66, int32(16), v165, v31+int32(-49))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+44)) = int32(0)
	goto L27
L31:
	;
	v170 = F_pg_detoast_datum_packed(m, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+15)))
	if v172 == int32(1) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	v176 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+34)))
	v179 = F_palloc(m, v176<<(uint(int32(2))%32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+44)) = v179
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+34)))
	if v182 <= int32(0) {
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v185 = int32(1)
	if v175&v185 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v189 = v185
	goto L38
L37:
	;
	v189 = int32(4)
	goto L38
L38:
	;
	v193 = v170 + v189
	v194 = int32(0)
	goto L39
L39:
	;
	v222 = F_pstrdup(m, v193)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	goto L27
L41:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v105)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v224+v194<<(uint(int32(2))%32)))) = v222
	if v193&int32(3) == int32(0) {
		v252 = v193
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v287 = int32(1)
	v290 = v194 + v287
	v291 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+34)))
	if v290 < v291 {
		v193 = v285 + v193 + v287
		v194 = v290
		goto L39
	} else {
		goto L59
	}
L43:
	;
	v285 = v277 - v193
	goto L42
L44:
	;
	v256 = v252
	goto L53
L45:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	if v236 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v285 = int32(0)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v241 = v193
	goto L49
L49:
	;
	v245 = v241 + int32(1)
	if v245&int32(3) == int32(0) {
		v252 = v245
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v277 = v245
	goto L43
L51:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245))))
	if v250 != 0 {
		v241 = v245
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v265 = int32(-2139062144)
	if (int32(16843008)-v262|v262)&v265 == v265 {
		v256 = v256 + int32(4)
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v271 = v256
	goto L56
L55:
	;
	goto L54
L56:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	if v275 != 0 {
		v271 = v271 + int32(1)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v277 = v271
	goto L43
L58:
	;
	goto L57
L59:
	;
	goto L40
L60:
	;
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+15)))
	if v333 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v338 = int32(0)
	goto L63
L62:
	;
	v336 = F_DirectFunctionCall1Coll(m, int32(580), int32(0), v330)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+52)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v344 = F_fastgetattr_2(m, v66, int32(19), v341, v31+int32(-49))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L65
	}
L64:
	;
	v338 = v336
	goto L63
L65:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+15)))
	if v346 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v351 = v325
	goto L68
L67:
	;
	v349 = F_DirectFunctionCall1Coll(m, int32(580), int32(0), v344)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+56)) = v351
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v357 = F_fastgetattr_2(m, v66, int32(17), v354, v31+int32(-49))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	v351 = v349
	goto L68
L70:
	;
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+15)))
	if v359 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v362 = v325
	goto L73
L72:
	;
	v360 = F_text_to_cstring(m, v357)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+48)) = v362
	v365 = v68 + int32(1)
	v366 = F_systable_getnext(m, v56)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	v362 = v360
	goto L73
L75:
	;
	if v366 != 0 {
		v66 = v366
		v68 = v365
		v69 = v101
		v71 = v102
		goto L11
	} else {
		goto L76
	}
L76:
	;
	goto L12
L77:
	;
	F_sequence_close(m, v48, int32(1))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v375 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	m.G0 = v33 - int32(-64)
	return
L80:
	;
	F_pfree(m, v378)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v408 = F_palloc0(m, int32(32))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L79
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v408)+4)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = v378
	if int32(0) < v375 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+28)))
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+27)))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+25)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+24)))
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+23)))
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+22)))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+21)))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+20)))
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+19)))
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+18)))
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+17)))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+16)))
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+15)))
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+14)))
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+13)))
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+11)))
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+10)))
	v432 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+9)))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+8)))
	v441 = int32(0)
	v443 = v414
	v446 = v415
	v447 = v416
	v448 = v417
	v449 = v418
	v450 = v419
	v451 = v420
	v452 = v421
	v453 = v422
	v454 = v423
	v455 = v424
	v456 = v425
	v457 = v426
	v458 = v427
	v459 = v428
	v460 = v429
	v461 = v430
	v462 = v431
	v463 = v432
	v464 = v433
	goto L88
L86:
	;
	goto L87
L87:
	;
	v610 = int32(4455216)
	v611 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v614 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v614
	v616 = F_CopyTriggerDesc(m, v408)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L100
	}
L88:
	;
	v467 = v378 + v441*int32(60)
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v467)+12)))
	v470 = v468 & int32(99)
	v473 = v448 | base.B2i32(v470 == int32(32))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+24)) = uint8(v473)
	v477 = v449 | base.B2i32(v470 == int32(34))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+23)) = uint8(v477)
	v480 = v468 & int32(75)
	v483 = v450 | base.B2i32(v480 == int32(8))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+22)) = uint8(v483)
	v487 = v451 | base.B2i32(v480 == int32(10))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+21)) = uint8(v487)
	v491 = v452 | base.B2i32(v480 == int32(73))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+20)) = uint8(v491)
	v495 = v453 | base.B2i32(v480 == int32(9))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+19)) = uint8(v495)
	v499 = v454 | base.B2i32(v480 == int32(11))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+18)) = uint8(v499)
	v502 = v468 & int32(83)
	v505 = v455 | base.B2i32(v502 == int32(16))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+17)) = uint8(v505)
	v509 = v456 | base.B2i32(v502 == int32(18))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+16)) = uint8(v509)
	v513 = v457 | base.B2i32(v502 == int32(81))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+15)) = uint8(v513)
	v517 = v458 | base.B2i32(v502 == int32(17))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+14)) = uint8(v517)
	v521 = v459 | base.B2i32(v502 == int32(19))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+13)) = uint8(v521)
	v524 = v468 & int32(71)
	v525 = int32(4)
	v527 = v460 | base.B2i32(v524 == v525)
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+12)) = uint8(v527)
	v531 = v461 | base.B2i32(v524 == int32(6))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+11)) = uint8(v531)
	v535 = v462 | base.B2i32(v524 == int32(69))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+10)) = uint8(v535)
	v539 = v463 | base.B2i32(v524 == int32(5))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+9)) = uint8(v539)
	v543 = v464 | base.B2i32(v524 == int32(7))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+8)) = uint8(v543)
	v545 = int32(0)
	if v468&v525 != 0 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L87
L90:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	v551 = base.B2i32(v548 != int32(0))
	goto L92
L91:
	;
	v551 = v545
	goto L92
L92:
	;
	v552 = v551 | v447
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+25)) = uint8(v552)
	if v468&int32(16) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408)+26)))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v467)+52))
	v559 = int32(0)
	v561 = v557 | base.B2i32(v558 != v559)
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+26)) = uint8(v561)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v467)+56))
	v566 = base.B2i32(v563 != v559)
	goto L95
L94:
	;
	v566 = int32(0)
	goto L95
L95:
	;
	v567 = v566 | v446
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+27)) = uint8(v567)
	if v468&int32(8) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v467)+52))
	v574 = base.B2i32(v571 != int32(0))
	goto L98
L97:
	;
	v574 = v545
	goto L98
L98:
	;
	v575 = v574 | v443
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+28)) = uint8(v575)
	v578 = v441 + int32(1)
	if v578 != v375 {
		v441 = v578
		v443 = v575
		v446 = v567
		v447 = v552
		v448 = v473
		v449 = v477
		v450 = v483
		v451 = v487
		v452 = v491
		v453 = v495
		v454 = v499
		v455 = v505
		v456 = v509
		v457 = v513
		v458 = v517
		v459 = v521
		v460 = v527
		v461 = v531
		v462 = v535
		v463 = v539
		v464 = v543
		goto L88
	} else {
		goto L99
	}
L99:
	;
	goto L89
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v616
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v611
	F_FreeTriggerDesc(m, v408)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L79
L102:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v660 + int32(4)
	F_errmsg_internal(m, int32(670066), v33)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(478723), int32(1946), int32(128025))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
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
			v20 = *(*int32)(unsafe.Add(mBase, _consts[869]))
			if v20 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[870]))
				v27 = v22
			} else {
				v24 = int32(4378696)
				*(*int32)(unsafe.Add(mBase, _consts[869])) = v24
				v27 = v24
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v27
			v29 = int32(4378696)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v29
			*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v18
			*(*int32)(unsafe.Add(mBase, _consts[870])) = v18
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
					v49 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
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
									F_errmsg_internal(m, int32(54240), v6)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errfinish(m, int32(482315), int32(2563), int32(254958))
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
				v49 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
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
								F_errmsg_internal(m, int32(54240), v6)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(482315), int32(2563), int32(254958))
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
				v49 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
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
								F_errmsg_internal(m, int32(54240), v6)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return
								} else {
									F_errfinish(m, int32(482315), int32(2563), int32(254958))
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
			v49 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
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
							F_errmsg_internal(m, int32(54240), v6)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(482315), int32(2563), int32(254958))
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int64
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
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
	var v122 int32
	_ = v122
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	v12 = int32(-1)
	switch l1 - int32(112) {
	case 0:
		v38 = int32(1)
		v40 = v12
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v43
		v47 = F_smgropen(m, v10+int32(16), v40)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = int32(0)
			F_smgrcreate(m, v47, v49, v49)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v38 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v53
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v55
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v68 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
									v73 = F_MemoryContextAlloc(m, v71, int32(28))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
										v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
										v79 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
										*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
										v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
										v86 = int32(4352048)
										v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
										*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(0) < v95 {
												m.G0 = v10 + int32(80)
												return v47
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
												if v99 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
													v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
														v116 = v113
														v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															v122 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
															m.G0 = v10 + int32(80)
															return v47
														}
													}
												} else {
													v116 = v99
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v47
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(0) < v95 {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
											if v99 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
												v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
													v116 = v113
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											} else {
												v116 = v99
												v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													v122 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
													m.G0 = v10 + int32(80)
													return v47
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
						v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
						v73 = F_MemoryContextAlloc(m, v71, int32(28))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
							v79 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
							v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
							v86 = int32(4352048)
							v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
							*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(0) < v95 {
									m.G0 = v10 + int32(80)
									return v47
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
									if v99 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
										v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
											v116 = v113
											v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												v122 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
												m.G0 = v10 + int32(80)
												return v47
											}
										}
									} else {
										v116 = v99
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v47
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(0) < v95 {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
								if v99 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
									v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
										v116 = v113
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								} else {
									v116 = v99
									v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
										m.G0 = v10 + int32(80)
										return v47
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
			F_errmsg_internal(m, int32(485506), v10)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(482384), int32(146), int32(392677))
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
		v32 = *(*int32)(unsafe.Add(mBase, _consts[100]))
		v34 = *(*int32)(unsafe.Add(mBase, _consts[442]))
		if v34 == int32(-1) {
			v37 = v32
		} else {
			v37 = v34
		}
		v38 = v4
		v40 = v37
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v43
		v47 = F_smgropen(m, v10+int32(16), v40)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = int32(0)
			F_smgrcreate(m, v47, v49, v49)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v38 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v53
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v55
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v68 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
									v73 = F_MemoryContextAlloc(m, v71, int32(28))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
										v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
										v79 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
										*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
										v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
										v86 = int32(4352048)
										v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
										*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(0) < v95 {
												m.G0 = v10 + int32(80)
												return v47
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
												if v99 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
													v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
														v116 = v113
														v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															v122 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
															m.G0 = v10 + int32(80)
															return v47
														}
													}
												} else {
													v116 = v99
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v47
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(0) < v95 {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
											if v99 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
												v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
													v116 = v113
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											} else {
												v116 = v99
												v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													v122 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
													m.G0 = v10 + int32(80)
													return v47
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
						v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
						v73 = F_MemoryContextAlloc(m, v71, int32(28))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
							v79 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
							v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
							v86 = int32(4352048)
							v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
							*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(0) < v95 {
									m.G0 = v10 + int32(80)
									return v47
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
									if v99 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
										v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
											v116 = v113
											v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												v122 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
												m.G0 = v10 + int32(80)
												return v47
											}
										}
									} else {
										v116 = v99
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v47
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(0) < v95 {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
								if v99 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
									v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
										v116 = v113
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								} else {
									v116 = v99
									v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
										m.G0 = v10 + int32(80)
										return v47
									}
								}
							}
						}
					}
				}
			}
		}
	case 5:
		v38 = v4
		v40 = v12
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v41
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v43
		v47 = F_smgropen(m, v10+int32(16), v40)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int32(0)
		} else {
			v49 = int32(0)
			F_smgrcreate(m, v47, v49, v49)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				if v38 != 0 {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v53
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v55
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_XLogRegisterData(m, v10+int32(32), int32(16))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v68 = F_XLogInsert(m, int32(2), int32(17))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
									v73 = F_MemoryContextAlloc(m, v71, int32(28))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return int32(0)
									} else {
										v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
										v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
										v79 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
										*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
										v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
										v86 = int32(4352048)
										v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
										*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
										*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
										if l1 != int32(112) {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
											if int32(0) < v95 {
												m.G0 = v10 + int32(80)
												return v47
											} else {
												v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
												if v99 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
													v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
													*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
													v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
													mBase = m.M
													v114 = m.ExcPending
													if v114 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
														v116 = v113
														v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															v122 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
															m.G0 = v10 + int32(80)
															return v47
														}
													}
												} else {
													v116 = v99
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											}
										}
									}
								} else {
									if l1 != int32(112) {
										m.G0 = v10 + int32(80)
										return v47
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
										if int32(0) < v95 {
											m.G0 = v10 + int32(80)
											return v47
										} else {
											v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
											if v99 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
												v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
												*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
												v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
													v116 = v113
													v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
														m.G0 = v10 + int32(80)
														return v47
													}
												}
											} else {
												v116 = v99
												v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return int32(0)
												} else {
													v122 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
													m.G0 = v10 + int32(80)
													return v47
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
						v71 = *(*int32)(unsafe.Add(mBase, _consts[12]))
						v73 = F_MemoryContextAlloc(m, v71, int32(28))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+8)) = v75
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
							v79 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v73)+16)) = uint8(v79)
							*(*int32)(unsafe.Add(mBase, uint32(v73)+12)) = v40
							v83 = *(*int32)(unsafe.Add(mBase, _consts[72]))
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+20)) = v84
							v86 = int32(4352048)
							v87 = *(*int32)(unsafe.Add(mBase, _consts[138]))
							*(*int32)(unsafe.Add(mBase, uint32(v73)+24)) = v87
							*(*int32)(unsafe.Add(mBase, _consts[138])) = v73
							if l1 != int32(112) {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
								if int32(0) < v95 {
									m.G0 = v10 + int32(80)
									return v47
								} else {
									v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
									if v99 == int32(0) {
										*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
										v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
										*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
										v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
											v116 = v113
											v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												v122 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
												m.G0 = v10 + int32(80)
												return v47
											}
										}
									} else {
										v116 = v99
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								}
							}
						}
					} else {
						if l1 != int32(112) {
							m.G0 = v10 + int32(80)
							return v47
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, _consts[15]))
							if int32(0) < v95 {
								m.G0 = v10 + int32(80)
								return v47
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, _consts[115]))
								if v99 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = int64(68719476748)
									v105 = *(*int32)(unsafe.Add(mBase, _consts[105]))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+72)) = v105
									v113 = F_hash_create(m, int32(311841), int32(16), v10+int32(32), int32(1064))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[115])) = v113
										v116 = v113
										v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											v122 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
											m.G0 = v10 + int32(80)
											return v47
										}
									}
								} else {
									v116 = v99
									v120 = F_hash_search(m, v116, l0, int32(1), v10+int32(32))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										v122 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v120)+12)) = uint8(v122)
										m.G0 = v10 + int32(80)
										return v47
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
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v188 int32
	_ = v188
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
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
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
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L62
	}
L6:
	;
	v50 = v37
	v52 = int32(0)
	v57 = v5
	goto L9
L7:
	;
	v143 = v5
	goto L8
L8:
	;
	v153 = int32(0)
	v155 = F_index_beginscan(m, l0, v23, v20+int32(40), v153, v143, v153)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L25
	}
L9:
	;
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31+int32(48)+v52<<(uint(int32(1))%32)))))
	if v66 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v143 = v126
	goto L8
L11:
	;
	v68 = v52 << (uint(int32(2)) % 32)
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
	v126 = v57
	goto L13
L13:
	;
	v129 = v52 + int32(1)
	v130 = int32(*(*int16)(unsafe.Add(mBase, uint32(v122)+10)))
	if v129 < v130 {
		v50 = v122
		v52 = v129
		v57 = v126
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
	v91 = v20 + int32(112) + v57*int32(48)
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
	F_ScanKeyInit(m, v91, base.I32_extend16_s(v52+int32(1)), v80, v95, v103)
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
	v126 = v57 + int32(1)
	goto L13
L24:
	;
	goto L10
L25:
	;
	v167 = int32(0)
	goto L26
L26:
	;
	v178 = int32(0)
	F_index_rescan(m, v155, v20+int32(112), v143, v178, v178)
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
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L60
	}
L28:
	;
	v188 = v167
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
	v217 = v167
	goto L34
L36:
	;
	goto L37
L37:
	;
	if v188 == int32(0) {
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
	v212 = v188
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
		v188 = v212
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
	v167 = v217
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
	v235 = *(*int32)(unsafe.Add(mBase, _consts[104]))
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
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+104))
	v247 = m.T0[v246].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l3+int32(28), v236, l3, v238, int32(3), v241, v241, v20+int32(20))
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
	v253 = F_should_refetch_tuple(m, v247, v20+int32(20))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v253 != 0 {
		v167 = v217
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
	v262 = m.ExcPending
	if v262 != 0 {
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
	F_errmsg_internal(m, int32(38124), v20)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(479554), int32(101), int32(20055))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
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
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
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
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
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
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v19 = F_palloc0(m, v16<<(uint(int32(2))%32))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(4)
	v29 = int32(0)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v35 = m.T0[v34].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v13+int32(24), v29, v29, v29, int32(449))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v38 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L38
	}
L6:
	;
	v50 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+188))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	m.T0[v57].(func(*base.Module, int32, int32, int32, int32, int32, int32))(m, v35, v50, v50, v50, v50, v50)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+188))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	m.T0[v133].(func(*base.Module, int32))(m, v35)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, _consts[452]))
	if v74 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+32))
	m.T0[v94].(func(*base.Module, int32, int32))(m, l2, v38)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[453])))
	if v76&int32(1) == int32(0) {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+188))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+20))
	v85 = m.T0[v84].(func(*base.Module, int32, int32, int32) int32)(m, v35, int32(1), v38)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	if v85 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	v89 = F_tuples_equal(m, v38, l1, v19)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v89 == int32(0) {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L11
L20:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v97 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v99 = v97
	goto L23
L22:
	;
	v99 = v98
	goto L23
L23:
	;
	if v99 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v100 = int32(0)
	F_XactLockTableWait(m, v99, v100, v100, v100)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v105 = F_GetLatestSnapshot(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L6
L28:
	;
	F_PushActiveSnapshot(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	goto L30
L30:
	;
	v113 = F_GetCurrentCommandId(m, int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v116 = int32(0)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+104))
	v122 = m.T0[v121].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, l0, l2+int32(28), v111, l2, v113, int32(3), v116, v116, v13+int32(4))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v128 = F_should_refetch_tuple(m, v122, v13+int32(4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v128 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	goto L9
L36:
	;
	F_ExecDropSingleTupleTableSlot(m, v38)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	m.G0 = v13 + int32(96)
	return v85
L38:
	;
	F_errmsg_internal(m, int32(323972), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(314654), int32(1034), int32(81516))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v167 int32
	_ = v167
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int64
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v666 int32
	_ = v666
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int64
	_ = v718
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
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v786 int32
	_ = v786
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v830 int32
	_ = v830
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v888 int32
	_ = v888
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v979 int32
	_ = v979
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1065 int32
	_ = v1065
	var v1070 int32
	_ = v1070
	var v1079 int32
	_ = v1079
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1111 int32
	_ = v1111
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1143 int32
	_ = v1143
	var v1148 int32
	_ = v1148
	var v1157 int32
	_ = v1157
	var v1168 int32
	_ = v1168
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int64
	_ = v1197
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1224 int32
	_ = v1224
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1274 int32
	_ = v1274
	var v1277 int32
	_ = v1277
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	v25 = m.G0
	v27 = v25 - int32(368)
	m.G0 = v27
	v32 = (l1 + int32(7)) & int32(-8)
	if base.Ui32(v32) < base.Ui32(int32(8161)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v1294 = m.ExcPending
	if v1294 != 0 {
		goto L33
	} else {
		goto L383
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
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L33
	} else {
		goto L379
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v44 = base.I32_div_s(int32(819200)-v39<<(uint(int32(13))%32), int32(100))
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
	if base.Ui32(int32(8016)) < base.Ui32(v32) {
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
	v51 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(l2^int32(-1))<<(uint(int32(6))%32))+16))
	v66 = v57
	goto L11
L13:
	;
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+l2<<(uint(int32(6))%32)+int32(-64))+16))
	v66 = v65
	goto L11
L15:
	;
	v70 = v32
	goto L17
L16:
	;
	v70 = int32(8016)
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
	if base.Ui32(int32(8016)) < base.Ui32(v47) {
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
	v83 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v83+(v77^int32(-1))<<(uint(int32(6))%32))+16))
	v98 = v89
	goto L22
L24:
	;
	goto L25
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _consts[7]))
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
	if v74 != 0 {
		v114 = v106
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v114 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	if v106 != int32(-1) {
		v114 = v106
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v110 = F_GetPageWithFreeSpace(m, l0, v107)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v114 = v110
	goto L30
L35:
	;
	v120 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v124 = v114
	goto L37
L37:
	;
	if int32(1) < l7 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v124 = v120 - int32(1)
	goto L37
L39:
	;
	v126 = l7
	goto L41
L40:
	;
	v126 = int32(1)
	goto L41
L41:
	;
	v128 = l3 & int32(4)
	v131 = int32(0)
	v132 = base.B2i32(v74 == v131)
	v141 = v124
	goto L45
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1244)+16)) = v1243
	m.G0 = v27 + int32(368)
	return v1249
L43:
	;
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1243 = v1218
	v1244 = v1241
	v1249 = v1224
	goto L42
L44:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1193 != 0 {
		v1243 = v986
		v1244 = v1193
		v1249 = v780
		goto L42
	} else {
		goto L373
	}
L45:
	;
	if v141 == int32(-1) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v1124 = int32(4)
	v1125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+14)))
	v1126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+12)))
	v1127 = v1125 - v1126
	if v1127 <= v1124 {
		goto L354
	} else {
		goto L355
	}
L47:
	;
	v595 = int32(1)
	if v132|base.B2i32(l4 != v131) != 0 {
		goto L196
	} else {
		goto L197
	}
L48:
	;
	v167 = v141
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
	F_LockBuffer(m, v371, int32(2))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
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
	if v167 == v68 {
		goto L93
	} else {
		goto L94
	}
L55:
	;
	v240 = int32(0)
	v241 = base.B2i32(v240 <= v238)
	if v241 == v240 {
		goto L75
	} else {
		goto L76
	}
L56:
	;
	v194 = int32(0)
	v197 = F_ReadBufferExtended(m, l0, v194, v167, v194, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L33
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v199 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v238 = v197
	goto L55
L60:
	;
	if v199 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	goto L62
L62:
	;
	v230 = int32(0)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v233 = F_ReadBufferExtended(m, l0, v230, v167, v230, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L33
	} else {
		goto L72
	}
L63:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v167 == v218 {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203+(v199^int32(-1))<<(uint(int32(6))%32))+16))
	v218 = v209
	goto L63
L65:
	;
	goto L66
L66:
	;
	v211 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211+v199<<(uint(int32(6))%32)+int32(-64))+16))
	v218 = v217
	goto L63
L67:
	;
	F_IncrBufferRefCount(m, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L33
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	F_ReleaseBuffer(m, v219)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L33
	} else {
		goto L71
	}
L70:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v238 = v223
	goto L55
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(0)
	goto L62
L72:
	;
	F_IncrBufferRefCount(m, v233)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L33
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v233
	v238 = v233
	goto L55
L74:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+10)))
	if v260&int32(4) != 0 {
		goto L78
	} else {
		goto L79
	}
L75:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v245+(v238^int32(-1))<<(uint(int32(2))%32))))
	v259 = v251
	goto L74
L76:
	;
	goto L77
L77:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v259 = v253 + v238<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	F_visibilitymap_pin(m, l0, v167, l5)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L33
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	if v128 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L80
L82:
	;
	v371 = v238
	v372 = v238
	goto L51
L83:
	;
	goto L84
L84:
	;
	if v241 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	F_visibilitymap_pin(m, l0, v167, l5)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L33
	} else {
		goto L92
	}
L86:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+12)))
	if base.Ui32(v285) < base.Ui32(int32(25)) {
		goto L85
	} else {
		goto L90
	}
L87:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v270+(v238^int32(-1))<<(uint(int32(2))%32))))
	v284 = v276
	goto L86
L88:
	;
	goto L89
L89:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v284 = v278 + v238<<(uint(int32(13))%32) + int32(-8192)
	goto L86
L90:
	;
	if (v285+int32(262120))&int32(262140) == int32(0) {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v371 = v238
	v372 = v238
	goto L51
L92:
	;
	v371 = v238
	v372 = v238
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
	v316 = F_ReadBuffer(m, l0, v167)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L33
	} else {
		goto L102
	}
L96:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+10)))
	if v309&int32(4) == int32(0) {
		v371 = l2
		v372 = l2
		goto L51
	} else {
		goto L100
	}
L97:
	;
	v300 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v300+(l2^int32(-1))<<(uint(int32(2))%32))))
	v308 = v302
	goto L96
L98:
	;
	goto L99
L99:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v308 = v304 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L96
L100:
	;
	F_visibilitymap_pin(m, l0, v68, l5)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L33
	} else {
		goto L101
	}
L101:
	;
	v371 = l2
	v372 = l2
	goto L51
L102:
	;
	if base.Ui32(v68) < base.Ui32(v167) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if v316 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	goto L105
L105:
	;
	if v316 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L106:
	;
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+10)))
	if v337&int32(4) != 0 {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322+(v316^int32(-1))<<(uint(int32(2))%32))))
	v336 = v328
	goto L106
L108:
	;
	goto L109
L109:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v336 = v330 + v316<<(uint(int32(13))%32) + int32(-8192)
	goto L106
L110:
	;
	F_visibilitymap_pin(m, l0, v167, l5)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
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
	v344 = m.ExcPending
	if v344 != 0 {
		goto L33
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v371 = v316
	v372 = v316
	goto L51
L115:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362)+10)))
	if v363&int32(4) != 0 {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	v348 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348+(v316^int32(-1))<<(uint(int32(2))%32))))
	v362 = v354
	goto L115
L117:
	;
	goto L118
L118:
	;
	v356 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v362 = v356 + v316<<(uint(int32(13))%32) + int32(-8192)
	goto L115
L119:
	;
	F_visibilitymap_pin(m, l0, v167, l5)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L33
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	F_LockBuffer(m, v316, int32(2))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L33
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v371 = l2
	v372 = v316
	goto L51
L124:
	;
	v377 = F_GetVisibilityMapPins(m, l0, v372, l2, v167, v68, l5, l6)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L33
	} else {
		goto L125
	}
L125:
	;
	if v372 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396)+14)))
	if v397 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L127:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v382+(v372^int32(-1))<<(uint(int32(2))%32))))
	v396 = v388
	goto L126
L128:
	;
	goto L129
L129:
	;
	v390 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v396 = v390 + v372<<(uint(int32(13))%32) + int32(-8192)
	goto L126
L130:
	;
	if v396&int32(3) != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	goto L132
L132:
	;
	v446 = int32(4)
	v447 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396)+14)))
	v448 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v396)+12)))
	v449 = v447 - v448
	if v449 <= v446 {
		goto L145
	} else {
		goto L146
	}
L133:
	;
	F_MarkBufferDirty(m, v372)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L33
	} else {
		goto L143
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v396)+10)) = int32(1572864)
	v432 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+18)) = uint16(v432)
	v438 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+16)) = uint16(v438)
	*(*uint16)(unsafe.Add(mBase, uint32(v396)+14)) = uint16(v438)
	goto L133
L135:
	;
	v426 = F___memset(m, v396, int32(0), int32(8192))
	mBase = m.M
	goto L134
L136:
	;
	goto L135
L143:
	;
	goto L132
L144:
	;
	if base.Ui32(v107) <= base.Ui32(v511) {
		goto L163
	} else {
		goto L164
	}
L145:
	;
	v452 = v446
	goto L147
L146:
	;
	v452 = v449
	goto L147
L147:
	;
	v454 = v452 - int32(4)
	if v454 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v511 = int32(0)
	goto L144
L149:
	;
	goto L150
L150:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v448) {
		goto L152
	} else {
		goto L153
	}
L151:
	;
	v511 = v454
	goto L144
L152:
	;
	v465 = int32(base.Ui32(v448+int32(262120)) >> (uint(int32(2)) % 32))
	goto L154
L153:
	;
	v465 = int32(0)
	goto L154
L154:
	;
	if base.Ui32(v465&int32(65535)) < base.Ui32(int32(291)) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v396)+10)))
	if v470&int32(1) == int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v511 = int32(0)
	goto L144
L157:
	;
	goto L158
L158:
	;
	v479 = int32(1)
	goto L159
L159:
	;
	v490 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v479&int32(65535)<<(uint(int32(2))%32)+(v396+int32(24))-int32(3)))))
	if v490&int32(384) == int32(0) {
		goto L151
	} else {
		goto L161
	}
L160:
	;
	v511 = int32(0)
	goto L144
L161:
	;
	v496 = v479 + int32(1)
	v497 = int32(65535)
	if base.Ui32(v496&v497) <= base.Ui32(v465&v497) {
		v479 = v496
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v513 != 0 {
		v1243 = v167
		v1244 = v513
		v1249 = v372
		goto L42
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	F_LockBuffer(m, v372, int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L33
	} else {
		goto L172
	}
L166:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v515
	v517 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+80)) = v517
	v521 = F_smgropen(m, v27+int32(80), v514)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L33
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v521
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v521)+72))
	if v525 != 0 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v1218 = v167
	v1224 = v372
	goto L43
L169:
	;
	v533 = v525
	goto L171
L170:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v521)+76))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v521)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v526)+4)) = v527
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v521)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v527))) = v529
	v531 = *(*int32)(unsafe.Add(mBase, uint32(v521)+72))
	v533 = v531
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v521)+72)) = v533 + int32(1)
	goto L168
L172:
	;
	if l2 != 0 {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	if l4 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L174:
	;
	if v167 == v68 {
		goto L173
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	F_ReleaseBuffer(m, v372)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L33
	} else {
		goto L179
	}
L177:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L33
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	goto L173
L180:
	;
	if v568 != int32(-1) {
		v167 = v568
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
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v548 == int32(-1) {
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
	F_RecordPageWithFreeSpace(m, l0, v167, v511)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L33
	} else {
		goto L187
	}
L185:
	;
	v556 = v548
	goto L186
L186:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if base.Ui32(v557) <= base.Ui32(v556) {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v556 = v555
	goto L186
L188:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l4)+8)) = int64(-1)
	v568 = v556
	goto L180
L189:
	;
	goto L190
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v556 + int32(1)
	v568 = v556
	goto L180
L191:
	;
	v565 = F_RecordAndGetPageWithFreeSpace(m, l0, v167, v511, v107)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L33
	} else {
		goto L192
	}
L192:
	;
	v568 = v565
	goto L180
L193:
	;
	goto L50
L194:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v27)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v716
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v27)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+64)) = v718
	v722 = m.G0
	v724 = v722 - int32(32)
	m.G0 = v724
	v727 = v27 - int32(-64)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v727)+4))
	if v728 == int32(0) {
		goto L230
	} else {
		goto L231
	}
L195:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	if base.Ui32(v693) < base.Ui32(v677) {
		goto L220
	} else {
		goto L221
	}
L196:
	;
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v597 != 0 {
		v674 = v595
		goto L199
	} else {
		goto L200
	}
L197:
	;
	v684 = v595
	goto L198
L198:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+100)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v684
	v711 = v684
	v714 = v595
	v715 = int32(0)
	goto L194
L199:
	;
	v677 = v674 * v126
	if l4 != 0 {
		goto L195
	} else {
		goto L216
	}
L200:
	;
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v598 != 0 {
		v674 = v595
		goto L199
	} else {
		goto L201
	}
L201:
	;
	v599 = m.G0
	v600 = int32(16)
	v601 = v599 - v600
	m.G0 = v601
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v601))) = v603
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int64)(unsafe.Add(mBase, uint32(v601)+8)) = int64(72339069014638592)
	*(*int32)(unsafe.Add(mBase, uint32(v601)+4)) = v605
	v610 = m.G0
	v612 = v610 - v600
	m.G0 = v612
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+15)))
	if base.Ui32(int32(253)) < base.Ui32((v614-int32(3))&int32(255)) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	m.G0 = v601 + int32(16)
	v674 = v648 + int32(1)
	goto L199
L203:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v623 = F_get_hash_value(m, v622, v601)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
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
	v657 = m.ExcPending
	if v657 != 0 {
		goto L33
	} else {
		goto L213
	}
L206:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v633 = v626 + v623&int32(15)<<(uint(int32(7))%32) + int32(23296)
	v635 = F_LWLockAcquire(m, v633, int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L33
	} else {
		goto L207
	}
L207:
	;
	v638 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v642 = F_hash_search_with_hash_value(m, v638, v601, v623, int32(0), v612+int32(15))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L33
	} else {
		goto L208
	}
L208:
	;
	v644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v612)+15)))
	if v644 == int32(1) {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v642)+84))
	v648 = v647
	goto L211
L210:
	;
	v648 = int32(0)
	goto L211
L211:
	;
	F_LWLockRelease(m, v633)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L33
	} else {
		goto L212
	}
L212:
	;
	m.G0 = v612 + int32(16)
	goto L202
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v612))) = v614
	F_errmsg_internal(m, int32(470690), v612)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L33
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(480813), int32(4834), int32(84704))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
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
	v678 = int32(64)
	if base.Ui32(v678) <= base.Ui32(v677) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v681 = v678
	goto L219
L218:
	;
	v681 = v677
	goto L219
L219:
	;
	v684 = v681
	goto L198
L220:
	;
	v695 = v677
	goto L222
L221:
	;
	v695 = v693
	goto L222
L222:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v695) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v698 = int32(64)
	goto L225
L224:
	;
	v698 = v695
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+108)) = v698
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v700 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_ReleaseBuffer(m, v700)
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
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
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v711 = v698
	v714 = v126
	v715 = v708
	goto L194
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = int32(0)
	goto L228
L230:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727)))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	if v732 != 0 {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	goto L232
L232:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v727)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v724)+8)) = v765
	v767 = *(*int64)(unsafe.Add(mBase, uint32(v727)))
	*(*int64)(unsafe.Add(mBase, uint32(v724))) = v767
	v774 = F_ExtendBufferedRelCommon(m, v724, int32(0), v715, int32(8), v711, int32(-1), v27+int32(112), v27+int32(108))
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L33
	} else {
		goto L241
	}
L233:
	;
	v758 = v732
	goto L235
L234:
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
		goto L236
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v727)+4)) = v758
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v731)+48))
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(v727)+8)) = uint8(v761)
	goto L232
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v731)+12)) = v740
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v740)+72))
	if v744 != 0 {
		goto L238
	} else {
		goto L239
	}
L237:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v731)+12))
	v758 = v756
	goto L235
L238:
	;
	v752 = v744
	goto L240
L239:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v740)+76))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v740)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v745)+4)) = v746
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v740)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v748
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v740)+72))
	v752 = v750
	goto L240
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v740)+72)) = v752 + int32(1)
	goto L237
L241:
	;
	m.G0 = v724 + int32(32)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	v781 = int32(0)
	v782 = base.B2i32(v781 <= v780)
	if v782 == v781 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	v944 = v774 + v779 - int32(1)
	if v74 != 0 {
		goto L277
	} else {
		goto L278
	}
L243:
	;
	v801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v800)+14)))
	if v801 == int32(0) {
		goto L247
	} else {
		goto L248
	}
L244:
	;
	v786 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v786+(v780^int32(-1))<<(uint(int32(2))%32))))
	v800 = v792
	goto L243
L245:
	;
	goto L246
L246:
	;
	v794 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v800 = v794 + v780<<(uint(int32(13))%32) + int32(-8192)
	goto L243
L247:
	;
	if v800&int32(3) != 0 {
		goto L252
	} else {
		goto L253
	}
L248:
	;
	goto L249
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L33
	} else {
		goto L274
	}
L250:
	;
	F_MarkBufferDirty(m, v780)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L33
	} else {
		goto L260
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v800)+10)) = int32(1572864)
	v836 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v800)+18)) = uint16(v836)
	v842 = int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v800)+16)) = uint16(v842)
	*(*uint16)(unsafe.Add(mBase, uint32(v800)+14)) = uint16(v842)
	goto L250
L252:
	;
	v830 = F___memset(m, v800, int32(0), int32(8192))
	mBase = m.M
	goto L251
L253:
	;
	goto L252
L260:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v849 = v132 & base.B2i32(base.Ui32(v714) < base.Ui32(v847))
	if v849 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	F_LockBuffer(m, v780, int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L33
	} else {
		goto L264
	}
L262:
	;
	v854 = v847
	goto L263
L263:
	;
	v856 = int32(1)
	if base.Ui32(v854) <= base.Ui32(v856) {
		v922 = v854
		goto L242
	} else {
		goto L265
	}
L264:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v854 = v853
	goto L263
L265:
	;
	v860 = v856
	goto L266
L266:
	;
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(112)+v860<<(uint(int32(2))%32))))
	F_ReleaseBuffer(m, v888)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L33
	} else {
		goto L268
	}
L267:
	;
	v922 = v898
	goto L242
L268:
	;
	if v74 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v897 = v860 + int32(1)
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	if base.Ui32(v897) < base.Ui32(v898) {
		v860 = v897
		goto L266
	} else {
		goto L273
	}
L270:
	;
	if base.Ui32(v860) < base.Ui32(v714) {
		goto L269
	} else {
		goto L271
	}
L271:
	;
	F_RecordPageWithFreeSpace(m, l0, v860+v774, int32(8168))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L33
	} else {
		goto L272
	}
L272:
	;
	goto L269
L273:
	;
	goto L267
L274:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v774
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = v904 + int32(4)
	F_errmsg_internal(m, int32(81464), v27+int32(48))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L33
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(479301), int32(361), int32(146913))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L33
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	if l4 != 0 {
		goto L281
	} else {
		goto L282
	}
L278:
	;
	if base.Ui32(v922) <= base.Ui32(v714) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	F_FreeSpaceMapVacuumRange(m, l0, v774+v714, v944)
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L33
	} else {
		goto L280
	}
L280:
	;
	goto L277
L281:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	v952 = base.B2i32(base.Ui32(int32(1)) < base.Ui32(v950))
	if base.Ui32(int32(1)) < base.Ui32(v950) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	if v780 < int32(0) {
		goto L292
	} else {
		goto L293
	}
L284:
	;
	v953 = v944
	goto L286
L285:
	;
	v953 = int32(-1)
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v953
	if base.Ui32(int32(1)) < base.Ui32(v950) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v958 = v774 + int32(1)
	goto L289
L288:
	;
	v958 = int32(-1)
	goto L289
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v958
	F_IncrBufferRefCount(m, v780)
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L33
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v780
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+16)) = v963 + v964
	goto L283
L291:
	;
	if v782 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L292:
	;
	v971 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v971+(v780^int32(-1))<<(uint(int32(6))%32))+16))
	v986 = v977
	goto L291
L293:
	;
	goto L294
L294:
	;
	v979 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v979+v780<<(uint(int32(6))%32)+int32(-64))+16))
	v986 = v985
	goto L291
L295:
	;
	if v128 == int32(0) {
		goto L303
	} else {
		goto L304
	}
L296:
	;
	v990 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v990+(v780^int32(-1))<<(uint(int32(2))%32))))
	v1004 = v996
	goto L295
L297:
	;
	goto L298
L298:
	;
	v998 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v1004 = v998 + v780<<(uint(int32(13))%32) + int32(-8192)
	goto L295
L299:
	;
	goto L46
L300:
	;
	v1041 = F_GetVisibilityMapPins(m, l0, l2, v780, v68, v986, l6, l5)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L33
	} else {
		goto L326
	}
L301:
	;
	F_LockBuffer(m, v780, int32(2))
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L33
	} else {
		goto L325
	}
L302:
	;
	F_LockBuffer(m, l2, int32(2))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L33
	} else {
		goto L324
	}
L303:
	;
	if v849 != 0 {
		goto L316
	} else {
		goto L317
	}
L304:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v1007 == int32(0) {
		goto L306
	} else {
		goto L307
	}
L305:
	;
	if v1015 != 0 {
		goto L303
	} else {
		goto L309
	}
L306:
	;
	v1015 = int32(0)
	goto L305
L307:
	;
	goto L308
L308:
	;
	v1011 = F_BufferGetBlockNumber(m, v1007)
	mBase = m.M
	v1013 = base.I32_div_u_s(v986, int32(32672))
	v1015 = base.B2i32(v1011 == v1013)
	goto L305
L309:
	;
	if v849 == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	F_LockBuffer(m, v780, int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L33
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	F_visibilitymap_pin(m, l0, v986, l5)
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L33
	} else {
		goto L314
	}
L313:
	;
	goto L312
L314:
	;
	if l2 != 0 {
		goto L302
	} else {
		goto L315
	}
L315:
	;
	goto L301
L316:
	;
	if l2 != 0 {
		goto L302
	} else {
		goto L319
	}
L317:
	;
	goto L318
L318:
	;
	if l2 == int32(0) {
		goto L299
	} else {
		goto L320
	}
L319:
	;
	goto L301
L320:
	;
	v1026 = F_ConditionalLockBuffer(m, l2)
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L33
	} else {
		goto L321
	}
L321:
	;
	if v1026 != 0 {
		v1040 = int32(0)
		goto L300
	} else {
		goto L322
	}
L322:
	;
	F_LockBuffer(m, v780, int32(0))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L33
	} else {
		goto L323
	}
L323:
	;
	goto L302
L324:
	;
	goto L301
L325:
	;
	v1040 = int32(1)
	goto L300
L326:
	;
	v1046 = int32(4)
	v1047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+14)))
	v1048 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1004)+12)))
	v1049 = v1047 - v1048
	if v1049 <= v1046 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if base.Ui32(v32) <= base.Ui32(v1111) {
		goto L44
	} else {
		goto L346
	}
L328:
	;
	v1052 = v1046
	goto L330
L329:
	;
	v1052 = v1049
	goto L330
L330:
	;
	v1054 = v1052 - int32(4)
	if v1054 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1111 = int32(0)
	goto L327
L332:
	;
	goto L333
L333:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1048) {
		goto L335
	} else {
		goto L336
	}
L334:
	;
	v1111 = v1054
	goto L327
L335:
	;
	v1065 = int32(base.Ui32(v1048+int32(262120)) >> (uint(int32(2)) % 32))
	goto L337
L336:
	;
	v1065 = int32(0)
	goto L337
L337:
	;
	if base.Ui32(v1065&int32(65535)) < base.Ui32(int32(291)) {
		goto L334
	} else {
		goto L338
	}
L338:
	;
	v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004)+10)))
	if v1070&int32(1) == int32(0) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1111 = int32(0)
	goto L327
L340:
	;
	goto L341
L341:
	;
	v1079 = int32(1)
	goto L342
L342:
	;
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1079&int32(65535)<<(uint(int32(2))%32)+(v1004+int32(24))-int32(3)))))
	if v1090&int32(384) == int32(0) {
		goto L334
	} else {
		goto L344
	}
L343:
	;
	v1111 = int32(0)
	goto L327
L344:
	;
	v1096 = v1079 + int32(1)
	v1097 = int32(65535)
	if base.Ui32(v1096&v1097) <= base.Ui32(v1065&v1097) {
		v1079 = v1096
		goto L342
	} else {
		goto L345
	}
L345:
	;
	goto L343
L346:
	;
	if v1041|v1040 != int32(1) {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	if l2 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	F_LockBuffer(m, l2, int32(0))
	mBase = m.M
	v1118 = m.ExcPending
	if v1118 != 0 {
		goto L33
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	F_UnlockReleaseBuffer(m, v780)
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L33
	} else {
		goto L352
	}
L351:
	;
	goto L350
L352:
	;
	v141 = v986
	goto L45
L353:
	;
	if base.Ui32(v1189) < base.Ui32(v32) {
		goto L1
	} else {
		goto L372
	}
L354:
	;
	v1130 = v1124
	goto L356
L355:
	;
	v1130 = v1127
	goto L356
L356:
	;
	v1132 = v1130 - int32(4)
	if v1132 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1189 = int32(0)
	goto L353
L358:
	;
	goto L359
L359:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1126) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v1189 = v1132
	goto L353
L361:
	;
	v1143 = int32(base.Ui32(v1126+int32(262120)) >> (uint(int32(2)) % 32))
	goto L363
L362:
	;
	v1143 = int32(0)
	goto L363
L363:
	;
	if base.Ui32(v1143&int32(65535)) < base.Ui32(int32(291)) {
		goto L360
	} else {
		goto L364
	}
L364:
	;
	v1148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1004)+10)))
	if v1148&int32(1) == int32(0) {
		goto L365
	} else {
		goto L366
	}
L365:
	;
	v1189 = int32(0)
	goto L353
L366:
	;
	goto L367
L367:
	;
	v1157 = int32(1)
	goto L368
L368:
	;
	v1168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1157&int32(65535)<<(uint(int32(2))%32)+(v1004+int32(24))-int32(3)))))
	if v1168&int32(384) == int32(0) {
		goto L360
	} else {
		goto L370
	}
L369:
	;
	v1189 = int32(0)
	goto L353
L370:
	;
	v1174 = v1157 + int32(1)
	v1175 = int32(65535)
	if base.Ui32(v1174&v1175) <= base.Ui32(v1143&v1175) {
		v1157 = v1174
		goto L368
	} else {
		goto L371
	}
L371:
	;
	goto L369
L372:
	;
	goto L44
L373:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1195
	v1197 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v1197
	v1201 = F_smgropen(m, v27+int32(32), v1194)
	mBase = m.M
	v1202 = m.ExcPending
	if v1202 != 0 {
		goto L33
	} else {
		goto L374
	}
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1201
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+72))
	if v1205 != 0 {
		goto L376
	} else {
		goto L377
	}
L375:
	;
	v1218 = v986
	v1224 = v780
	goto L43
L376:
	;
	v1213 = v1205
	goto L378
L377:
	;
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+76))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v1206)+4)) = v1207
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v1207))) = v1209
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1201)+72))
	v1213 = v1211
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1201)+72)) = v1213 + int32(1)
	goto L375
L379:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L33
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(8160)
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v32
	F_errmsg(m, int32(35870), v27)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L33
	} else {
		goto L381
	}
L381:
	;
	F_errfinish(m, int32(479301), int32(536), int32(370828))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L33
	} else {
		goto L382
	}
L382:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v32
	F_errmsg_internal(m, int32(35981), v27+int32(16))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L33
	} else {
		goto L384
	}
L384:
	;
	F_errfinish(m, int32(479301), int32(870), int32(370828))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L33
	} else {
		goto L385
	}
L385:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationGetIndexExpressions(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+228))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v124
L2:
	;
	v14 = F_copyObjectImpl(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v124 = v14
	goto L1
L7:
	;
	v124 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v22 = int32(0)
	v25 = F_heap_attisnull(m, v18, int32(20), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v25 != 0 {
		v124 = v22
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1080]))
	if v29 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = int32(4455216)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
	v39 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v79 = v29
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+18)))
	if base.Ui32(v86&int32(2044)) <= base.Ui32(int32(19)) {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(-4294965047)
	v46 = v22
	goto L16
L16:
	;
	v53 = int32(100)
	v54 = v46 * v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	goto L19
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1080])) = v39
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
	v79 = v39
	goto L14
L18:
	;
	F_populate_compact_attribute(m, v39, v46)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	v63 = F__emscripten_memcpy_bulkmem(m, v54+(v39+int32(20)+v55<<(uint(int32(4))%32)), v54+int32(1724512), v53)
	mBase = m.M
	goto L21
L21:
	;
	goto L18
L22:
	;
	v68 = v46 + int32(1)
	if v68 != int32(21) {
		v46 = v68
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v102 = F_text_to_cstring(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L30
	}
L25:
	;
	v94 = F_getmissingattr(m, v79, int32(20), v11+int32(15))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v99 = F_fastgetattr_3(m, v27, int32(20), v79, v11+int32(15))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	v101 = v94
	goto L24
L29:
	;
	v101 = v99
	goto L24
L30:
	;
	v104 = F_stringToNode(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_pfree(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v109 = F_eval_const_expressions(m, int32(0), v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	F_fix_opfuncids(m, v109)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v113 = int32(4455216)
	v114 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v116
	v118 = F_copyObjectImpl(m, v109)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+228)) = v118
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v114
	v124 = v109
	goto L1
}
func F_RelationGetIndexPredicate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
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
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v129
L2:
	;
	v14 = F_copyObjectImpl(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v129 = v14
	goto L1
L7:
	;
	v129 = int32(0)
	goto L1
L8:
	;
	goto L9
L9:
	;
	v22 = int32(0)
	v25 = F_heap_attisnull(m, v18, int32(21), v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	if v25 != 0 {
		v129 = v22
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1080]))
	if v29 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v32 = int32(4455216)
	v33 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v36 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v36
	v39 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L15
	}
L13:
	;
	v79 = v29
	goto L14
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+18)))
	if base.Ui32(v86&int32(2047)) <= base.Ui32(int32(20)) {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(-4294965047)
	v46 = v22
	goto L16
L16:
	;
	v53 = int32(100)
	v54 = v46 * v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	goto L19
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1080])) = v39
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
	v79 = v39
	goto L14
L18:
	;
	F_populate_compact_attribute(m, v39, v46)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	v63 = F__emscripten_memcpy_bulkmem(m, v54+(v39+int32(20)+v55<<(uint(int32(4))%32)), v54+int32(1724512), v53)
	mBase = m.M
	goto L21
L21:
	;
	goto L18
L22:
	;
	v68 = v46 + int32(1)
	if v68 != int32(21) {
		v46 = v68
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L17
L24:
	;
	v102 = F_text_to_cstring(m, v101)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L30
	}
L25:
	;
	v94 = F_getmissingattr(m, v79, int32(21), v11+int32(15))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v99 = F_fastgetattr_3(m, v27, int32(21), v79, v11+int32(15))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	v101 = v94
	goto L24
L29:
	;
	v101 = v99
	goto L24
L30:
	;
	v104 = F_stringToNode(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_pfree(m, v102)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v109 = F_eval_const_expressions(m, int32(0), v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v112 = F_canonicalize_qual(m, v109, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v114 = F_make_ands_implicit(m, v112)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_fix_opfuncids(m, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v118 = int32(4455216)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v121
	v123 = F_copyObjectImpl(m, v114)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L5
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+232)) = v123
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v119
	v129 = v114
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
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
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v630 int32
	_ = v630
	var v636 int32
	_ = v636
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v778 int32
	_ = v778
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v825 int32
	_ = v825
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v991 int32
	_ = v991
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1101 int32
	_ = v1101
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
	var v1139 int32
	_ = v1139
	var v1146 int32
	_ = v1146
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int32
	_ = v1217
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1293 int32
	_ = v1293
	var v1297 int32
	_ = v1297
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1318 int32
	_ = v1318
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1340 int32
	_ = v1340
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1385 int32
	_ = v1385
	var v1396 int32
	_ = v1396
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1416 int32
	_ = v1416
	var v1423 int32
	_ = v1423
	var v1428 int32
	_ = v1428
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1443 int32
	_ = v1443
	var v1454 int32
	_ = v1454
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1531 int32
	_ = v1531
	var v1545 int32
	_ = v1545
	var v1562 int32
	_ = v1562
	var v1574 int32
	_ = v1574
	var v1591 int32
	_ = v1591
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1635 int32
	_ = v1635
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1655 int32
	_ = v1655
	var v1658 int32
	_ = v1658
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1729 int32
	_ = v1729
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
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1820 int32
	_ = v1820
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1854 int32
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1860 int32
	_ = v1860
	var v1873 int32
	_ = v1873
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1895 int32
	_ = v1895
	var v1902 int32
	_ = v1902
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1950 int32
	_ = v1950
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1959 int32
	_ = v1959
	var v1970 int32
	_ = v1970
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2031 int32
	_ = v2031
	var v2036 int32
	_ = v2036
	var v2042 int32
	_ = v2042
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2087 int32
	_ = v2087
	var v2092 int32
	_ = v2092
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2126 int32
	_ = v2126
	var v2172 int32
	_ = v2172
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2180 int32
	_ = v2180
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2209 int32
	_ = v2209
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2236 int32
	_ = v2236
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2268 int32
	_ = v2268
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2309 int32
	_ = v2309
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2329 int32
	_ = v2329
	var v2336 int32
	_ = v2336
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2342 int32
	_ = v2342
	var v2345 int32
	_ = v2345
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2404 int32
	_ = v2404
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2427 int32
	_ = v2427
	var v2438 int32
	_ = v2438
	var v2458 int32
	_ = v2458
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2479 int32
	_ = v2479
	var v2481 int32
	_ = v2481
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2488 int32
	_ = v2488
	var v2494 int32
	_ = v2494
	var v2523 int32
	_ = v2523
	var v2552 int32
	_ = v2552
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2557 int32
	_ = v2557
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2578 int32
	_ = v2578
	var v2580 int32
	_ = v2580
	var v2588 int32
	_ = v2588
	var v2599 int32
	_ = v2599
	var v2612 int32
	_ = v2612
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2629 int32
	_ = v2629
	var v2631 int32
	_ = v2631
	var v2633 int32
	_ = v2633
	var v2634 int32
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2641 int32
	_ = v2641
	var v2643 int32
	_ = v2643
	var v2646 int32
	_ = v2646
	var v2649 int32
	_ = v2649
	var v2651 int32
	_ = v2651
	var v2666 int32
	_ = v2666
	var v2681 int32
	_ = v2681
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2686 int32
	_ = v2686
	var v2691 int32
	_ = v2691
	var v2693 int32
	_ = v2693
	var v2696 int32
	_ = v2696
	var v2754 int32
	_ = v2754
	var v2758 int32
	_ = v2758
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2773 int32
	_ = v2773
	var v2774 int32
	_ = v2774
	var v2779 int32
	_ = v2779
	var v2786 int32
	_ = v2786
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2815 int32
	_ = v2815
	var v2822 int32
	_ = v2822
	var v2837 int32
	_ = v2837
	var v2840 int32
	_ = v2840
	var v2844 int32
	_ = v2844
	var v2848 int32
	_ = v2848
	var v2849 int32
	_ = v2849
	var v2854 int32
	_ = v2854
	var v2861 int32
	_ = v2861
	var v2875 int32
	_ = v2875
	var v2879 int32
	_ = v2879
	var v2883 int32
	_ = v2883
	var v2884 int32
	_ = v2884
	var v2889 int32
	_ = v2889
	var v2896 int32
	_ = v2896
	var v2914 int32
	_ = v2914
	var v2944 int32
	_ = v2944
	var v2948 int32
	_ = v2948
	var v2953 int32
	_ = v2953
	var v2957 int32
	_ = v2957
	var v2963 int32
	_ = v2963
	var v2968 int32
	_ = v2968
	var v2972 int32
	_ = v2972
	var v2979 int32
	_ = v2979
	var v2984 int32
	_ = v2984
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
	v2972 = m.ExcPending
	if v2972 != 0 {
		goto L17
	} else {
		goto L475
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2957 = m.ExcPending
	if v2957 != 0 {
		goto L17
	} else {
		goto L472
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2944 = m.ExcPending
	if v2944 != 0 {
		goto L17
	} else {
		goto L469
	}
L4:
	;
	m.G0 = v29 + int32(128)
	return v2914
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
		v2914 = v31
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+4)))
	if v36&int32(1) == int32(0) {
		v2914 = v31
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[104]))
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
	v2914 = v45
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
	v52 = *(*int32)(unsafe.Add(mBase, _consts[104]))
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[104]))
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
	v2914 = v65
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
	v2251 = *(*int32)(unsafe.Add(mBase, _consts[174]))
	v2256 = F_AllocSetContextCreateInternal(m, v2251, int32(200116), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L17
	} else {
		goto L322
	}
L22:
	;
	v2236 = v2209
	v2242 = v2215
	v2245 = v2218
	v2246 = v2219
	v2249 = int32(0)
	goto L21
L23:
	;
	v96 = v3
	v98 = v79
	v109 = v81
	v110 = v79 + int32(4)
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
	v86 = v3
	v87 = int32(0)
	goto L27
L27:
	;
	v2209 = v87
	v2215 = v86
	v2218 = v3
	v2219 = v3
	goto L22
L28:
	;
	v86 = v81
	v87 = int32(0)
	goto L27
L29:
	;
	v117 = v109 << (uint(int32(2)) % 32)
	v118 = F_palloc(m, v117)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L17
	} else {
		goto L31
	}
L30:
	;
	v2209 = v2175
	v2215 = int32(0)
	v2218 = v118
	v2219 = v120
	goto L22
L31:
	;
	v120 = F_palloc(m, v109)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v122 = F_palloc(m, v117)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v124 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	if v124 < v125 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	F_ReceiveSharedInvalidMessages(m)
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L17
	} else {
		goto L316
	}
L35:
	;
	v130 = v124
	goto L38
L36:
	;
	goto L37
L37:
	;
	v350 = int32(0)
	v355 = v29 - int32(-64)
	v357 = v109 << (uint(int32(2)) % 32)
	v358 = F_palloc(m, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L17
	} else {
		goto L95
	}
L38:
	;
	v156 = v130 << (uint(int32(2)) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156+v157)))
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
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	if v298 != int32(98) {
		goto L2
	} else {
		goto L86
	}
L41:
	;
	v184 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	v168 = F_SysCacheGetAttr(m, int32(57), v160, int32(34), v29-int32(-64))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+64)))
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
	v175 = F_text_to_cstring(m, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L17
	} else {
		goto L49
	}
L48:
	;
	goto L41
L49:
	;
	v177 = F_stringToNode(m, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L17
	} else {
		goto L50
	}
L50:
	;
	F_ReleaseCatCache(m, v160)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L17
	} else {
		goto L51
	}
L51:
	;
	if v177 != 0 {
		v292 = v177
		goto L40
	} else {
		goto L52
	}
L52:
	;
	goto L41
L53:
	;
	F_ScanKeyInit(m, v29-int32(-64), int32(1), int32(3), int32(184), v159)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L17
	} else {
		goto L54
	}
L54:
	;
	v193 = int32(0)
	v195 = int32(1)
	v200 = F_systable_beginscan(m, v184, int32(2662), v195, v193, v195, v29-int32(-64))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L17
	} else {
		goto L56
	}
L55:
	;
	F_systable_endscan(m, v200)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L17
	} else {
		goto L82
	}
L56:
	;
	v202 = F_systable_getnext(m, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L17
	} else {
		goto L57
	}
L57:
	;
	if v202 == int32(0) {
		v276 = v193
		goto L55
	} else {
		goto L58
	}
L58:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v184)+52))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v202)+16))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207)+18)))
	if base.Ui32(v208&int32(2046)) <= base.Ui32(int32(33)) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)))
	if v271 != 0 {
		v276 = int32(0)
		goto L55
	} else {
		goto L79
	}
L60:
	;
	v216 = F_getmissingattr(m, v206, int32(34), v29+int32(63))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L17
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v218)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+20)))
	if v220&int32(1) == v218 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v269 = v216
	goto L59
L64:
	;
	v264 = F_nocachegetattr(m, v202, int32(34), v206)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L17
	} else {
		goto L78
	}
L65:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v206)+548))
	if v225 < int32(0) {
		goto L64
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+27)))
	if v256&int32(2) != 0 {
		goto L64
	} else {
		goto L77
	}
L68:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+22)))
	v230 = v207 + v228 + v225
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+554)))
	if v231 != int32(1) {
		v269 = v230
		goto L59
	} else {
		goto L69
	}
L69:
	;
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206)+552)))
	switch v234 - int32(1) {
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
	v243 = m.ExcPending
	if v243 != 0 {
		goto L17
	} else {
		goto L74
	}
L71:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v269 = v239
	goto L59
L72:
	;
	v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(v230))))
	v269 = v238
	goto L59
L73:
	;
	v237 = int32(*(*int8)(unsafe.Add(mBase, uint32(v230))))
	v269 = v237
	goto L59
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = base.I32_extend16_s(v234)
	F_errmsg_internal(m, int32(466883), v29+int32(48))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L17
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(314613), int32(70), int32(65727))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
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
	v259 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v259)
	v269 = int32(0)
	goto L59
L78:
	;
	v269 = v264
	goto L59
L79:
	;
	v272 = F_text_to_cstring(m, v269)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L17
	} else {
		goto L80
	}
L80:
	;
	v274 = F_stringToNode(m, v272)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L17
	} else {
		goto L81
	}
L81:
	;
	v276 = v274
	goto L55
L82:
	;
	F_sequence_close(m, v184, int32(1))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L17
	} else {
		goto L83
	}
L83:
	;
	if base.B2i32(v276 == int32(0))&(v96^int32(-1)) != 0 {
		goto L34
	} else {
		goto L84
	}
L84:
	;
	if v276 == int32(0) {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	v292 = v276
	goto L40
L86:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292)+5)))
	if v301 == int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v305 = F_get_default_partition_oid(m, v304)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L17
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v156+v118))) = v159
	v312 = F_get_rel_relkind(m, v159)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L17
	} else {
		goto L92
	}
L90:
	;
	if v305 != v159 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v130+v120))) = uint8(base.B2i32(v312 != int32(112)))
	*(*int32)(unsafe.Add(mBase, uint32(v156+v122))) = v292
	v320 = v130 + int32(1)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	if v320 < v321 {
		v130 = v320
		goto L38
	} else {
		goto L93
	}
L93:
	;
	goto L39
L94:
	;
	v2236 = int32(1)
	v2242 = v109
	v2245 = v118
	v2246 = v120
	v2249 = v2172
	goto L21
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v355))) = v358
	if v109 <= int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	switch v510 - int32(104) {
	case 0:
		goto L113
	default:
		v2126 = int32(0)
		goto L108
	case 4:
		goto L112
	case 10:
		goto L110
	}
L97:
	;
	v364 = v109 & int32(3)
	if base.Ui32(int32(4)) <= base.Ui32(v109) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v371 = v350
	v375 = v350
	goto L101
L99:
	;
	v420 = v350
	goto L100
L100:
	;
	if v364 == int32(0) {
		goto L96
	} else {
		goto L104
	}
L101:
	;
	v396 = v371 << (uint(int32(2)) % 32)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v399 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v396+v397))) = v399
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v401+v396)+4)) = v399
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v405+v396)+8)) = v399
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v409+v396)+12)) = v399
	v413 = int32(4)
	v414 = v371 + v413
	v416 = v375 + v413
	if v416 != v109&int32(2147483644) {
		v371 = v414
		v375 = v416
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v420 = v414
	goto L100
L103:
	;
	goto L102
L104:
	;
	v448 = v420
	v449 = v350
	goto L105
L105:
	;
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v472+v448<<(uint(int32(2))%32)))) = int32(-1)
	v478 = int32(1)
	v481 = v449 + v478
	if v481 != v364 {
		v448 = v448 + v478
		v449 = v481
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
	v2172 = v2126
	goto L94
L109:
	;
	F_qsort_arg(m, v1854, v1848, int32(8), int32(907), v67)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L17
	} else {
		goto L275
	}
L110:
	;
	v1304 = F_palloc0(m, int32(36))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L17
	} else {
		goto L205
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L17
	} else {
		goto L202
	}
L112:
	;
	v920 = F_palloc0(m, int32(36))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L17
	} else {
		goto L152
	}
L113:
	;
	v514 = F_palloc0(m, int32(36))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L17
	} else {
		goto L114
	}
L114:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v514)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v516
	v521 = v109 * int32(12)
	v522 = F_palloc(m, v521)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L17
	} else {
		goto L115
	}
L115:
	;
	if int32(0) < v109 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v529 = int32(0)
	goto L119
L117:
	;
	goto L118
L118:
	;
	F_pg_qsort(m, v522, v109, int32(12), int32(905))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L17
	} else {
		goto L123
	}
L119:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v122+v529<<(uint(int32(2))%32))))
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556)+4)))
	if v557 != int32(104) {
		goto L111
	} else {
		goto L121
	}
L120:
	;
	goto L118
L121:
	;
	v562 = v522 + v529*int32(12)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v556)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v562))) = v563
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v556)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v562)+8)) = v529
	*(*int32)(unsafe.Add(mBase, uint32(v562)+4)) = v565
	v569 = v529 + int32(1)
	if v569 != v109 {
		v529 = v569
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v521+v522-int32(12))))
	*(*int32)(unsafe.Add(mBase, uint32(v514)+4)) = v109
	v606 = F_palloc0(m, v357)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L17
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+20)) = v604
	*(*int64)(unsafe.Add(mBase, uint32(v514)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v514)+8)) = v606
	v614 = F_palloc(m, v604<<(uint(int32(2))%32))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L17
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+24)) = v614
	if v604 <= int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v770 = F_palloc(m, v109<<(uint(int32(3))%32))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L17
	} else {
		goto L138
	}
L127:
	;
	v620 = v604 & int32(3)
	v621 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v604) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v630 = v621
	v636 = int32(0)
	goto L131
L129:
	;
	v679 = v621
	goto L130
L130:
	;
	if v620 == int32(0) {
		goto L126
	} else {
		goto L134
	}
L131:
	;
	v655 = v630 << (uint(int32(2)) % 32)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	v658 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v655+v656))) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v660+v655)+4)) = v658
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v664+v655)+8)) = v658
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v668+v655)+12)) = v658
	v672 = int32(4)
	v673 = v630 + v672
	v675 = v636 + v672
	if v675 != v604&int32(2147483644) {
		v630 = v673
		v636 = v675
		goto L131
	} else {
		goto L133
	}
L132:
	;
	v679 = v673
	goto L130
L133:
	;
	goto L132
L134:
	;
	v707 = v679
	v711 = v621
	goto L135
L135:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v731+v707<<(uint(int32(2))%32)))) = int32(-1)
	v737 = int32(1)
	v740 = v711 + v737
	if v740 != v620 {
		v707 = v707 + v737
		v711 = v740
		goto L135
	} else {
		goto L137
	}
L136:
	;
	goto L126
L137:
	;
	goto L136
L138:
	;
	if int32(0) < v109 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v778 = int32(0)
	goto L142
L140:
	;
	goto L141
L141:
	;
	F_pfree(m, v522)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L17
	} else {
		goto L151
	}
L142:
	;
	v803 = v522 + v778*int32(12)
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v803)))
	v807 = v778 << (uint(int32(2)) % 32)
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v807+v808))) = v770 + v778<<(uint(int32(3))%32)
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v814+v807)))
	*(*int32)(unsafe.Add(mBase, uint32(v816))) = v805
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v818+v807)))
	*(*int32)(unsafe.Add(mBase, uint32(v820)+4)) = v804
	if v804 < v604 {
		goto L144
	} else {
		goto L145
	}
L143:
	;
	goto L141
L144:
	;
	v825 = v804
	goto L147
L145:
	;
	goto L146
L146:
	;
	v882 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v803)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v882+v883<<(uint(int32(2))%32)))) = v778
	v889 = v778 + int32(1)
	if v889 != v109 {
		v778 = v889
		goto L142
	} else {
		goto L150
	}
L147:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v514)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v849+v825<<(uint(int32(2))%32)))) = v778
	v854 = v825 + v805
	if v854 < v604 {
		v825 = v854
		goto L147
	} else {
		goto L149
	}
L148:
	;
	goto L146
L149:
	;
	goto L148
L150:
	;
	goto L143
L151:
	;
	v2172 = v514
	goto L94
L152:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v920)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v920))) = v922
	if int32(0) < v109 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v933 = int32(0)
	v939 = v350
	goto L156
L154:
	;
	goto L155
L155:
	;
	v1285 = int32(0)
	v1287 = F_palloc(m, v1285)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L17
	} else {
		goto L201
	}
L156:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v122+v939<<(uint(int32(2))%32))))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)+16))
	if v959 == int32(0) {
		v1101 = v933
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v1129 = F_palloc(m, v1101<<(uint(int32(3))%32))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L17
	} else {
		goto L173
	}
L158:
	;
	v1124 = v939 + int32(1)
	if v1124 != v109 {
		v933 = v1101
		v939 = v1124
		goto L156
	} else {
		goto L172
	}
L159:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v959)+4))
	if v962 <= int32(0) {
		v1101 = v933
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v966 = v962 & int32(3)
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v959)+12))
	if base.Ui32(v962) < base.Ui32(int32(4)) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	if v966 == int32(0) {
		v1101 = v1034
		goto L158
	} else {
		goto L168
	}
L162:
	;
	v1032 = int32(0)
	v1034 = v933
	goto L161
L163:
	;
	goto L164
L164:
	;
	v974 = int32(0)
	v978 = v974
	v980 = v933
	v991 = v974
	goto L165
L165:
	;
	v1004 = v967 + v978<<(uint(int32(2))%32)
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v1004)))
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1005)+24)))
	v1007 = int32(1)
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+4))
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1010)+24)))
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+8))
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1015)+24)))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1004)+12))
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1020)+24)))
	v1024 = v980 + (v1006 ^ v1007) + (v1011 ^ v1007) + (v1016 ^ v1007) + (v1021 ^ v1007)
	v1025 = int32(4)
	v1026 = v978 + v1025
	v1028 = v991 + v1025
	if v1028 != v962&int32(2147483644) {
		v978 = v1026
		v980 = v1024
		v991 = v1028
		goto L165
	} else {
		goto L167
	}
L166:
	;
	v1032 = v1026
	v1034 = v1024
	goto L161
L167:
	;
	goto L166
L168:
	;
	v1060 = v1032
	v1062 = v1034
	v1063 = int32(0)
	goto L169
L169:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v967+v1060<<(uint(int32(2))%32))))
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087)+24)))
	v1089 = int32(1)
	v1091 = v1062 + (v1088 ^ v1089)
	v1095 = v1063 + v1089
	if v1095 != v966 {
		v1060 = v1060 + v1089
		v1062 = v1091
		v1063 = v1095
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v1101 = v1091
	goto L158
L171:
	;
	goto L170
L172:
	;
	goto L157
L173:
	;
	v1132 = int32(-1)
	v1139 = int32(0)
	v1146 = v1132
	v1149 = int32(0)
	v1150 = v1132
	goto L175
L174:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1274 = m.ExcPending
	if v1274 != 0 {
		goto L17
	} else {
		goto L198
	}
L175:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v122+v1139<<(uint(int32(2))%32))))
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+4)))
	if v1164 == int32(108) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L17
	} else {
		goto L195
	}
L177:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+5)))
	if v1167 != 0 {
		goto L181
	} else {
		goto L182
	}
L178:
	;
	goto L179
L179:
	;
	goto L176
L180:
	;
	v1256 = v1139 + int32(1)
	if v1256 != v109 {
		v1139 = v1256
		v1146 = v1241
		v1149 = v1244
		v1150 = v1245
		goto L175
	} else {
		goto L194
	}
L181:
	;
	v1241 = v1146
	v1244 = v1149
	v1245 = v1139
	goto L180
L182:
	;
	goto L183
L183:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+16))
	if v1168 == int32(0) {
		v1241 = v1146
		v1244 = v1149
		v1245 = v1150
		goto L180
	} else {
		goto L184
	}
L184:
	;
	v1171 = int32(0)
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+4))
	if v1172 <= v1171 {
		v1241 = v1146
		v1244 = v1149
		v1245 = v1150
		goto L180
	} else {
		goto L185
	}
L185:
	;
	v1177 = v1171
	v1187 = v1146
	v1189 = v1172
	v1190 = v1149
	goto L186
L186:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+12))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1201+v1177<<(uint(int32(2))%32))))
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1205)+24)))
	if v1206 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v1241 = v1223
	v1244 = v1225
	v1245 = v1150
	goto L180
L188:
	;
	v1227 = v1177 + int32(1)
	if v1227 < v1224 {
		v1177 = v1227
		v1187 = v1223
		v1189 = v1224
		v1190 = v1225
		goto L186
	} else {
		goto L193
	}
L189:
	;
	v1211 = v1129 + v1190<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v1211))) = v1139
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1211)+4)) = v1213
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1168)+4))
	v1223 = v1187
	v1224 = v1217
	v1225 = v1190 + int32(1)
	goto L188
L190:
	;
	goto L191
L191:
	;
	if base.B2i32(v1187 == int32(-1)) == int32(0) {
		goto L174
	} else {
		goto L192
	}
L192:
	;
	v1223 = v1139
	v1224 = v1189
	v1225 = v1190
	goto L188
L193:
	;
	goto L187
L194:
	;
	v1848 = v1101
	v1854 = v1129
	v1856 = v1241
	v1860 = v1245
	goto L109
L195:
	;
	F_errmsg_internal(m, int32(475006), int32(0))
	mBase = m.M
	v1265 = m.ExcPending
	if v1265 != 0 {
		goto L17
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(477464), int32(493), int32(164948))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L17
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
	F_errmsg_internal(m, int32(400705), int32(0))
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L17
	} else {
		goto L199
	}
L199:
	;
	F_errfinish(m, int32(477464), int32(523), int32(164948))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L17
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
	v1848 = v1285
	v1854 = v1287
	v1856 = int32(-1)
	v1860 = int32(-1)
	goto L109
L202:
	;
	F_errmsg_internal(m, int32(475006), int32(0))
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L17
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(477464), int32(372), int32(165013))
	mBase = m.M
	v1302 = m.ExcPending
	if v1302 != 0 {
		goto L17
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v1304)+28)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1304))) = v1306
	v1310 = int32(-1)
	v1313 = F_palloc0(m, v109<<(uint(int32(3))%32))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L17
	} else {
		goto L206
	}
L206:
	;
	if v109 <= int32(0) {
		goto L209
	} else {
		goto L210
	}
L207:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L17
	} else {
		goto L272
	}
L208:
	;
	F_qsort_arg(m, v1313, v1385, int32(4), int32(906), v67)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L17
	} else {
		goto L222
	}
L209:
	;
	v1385 = int32(0)
	v1396 = v1310
	goto L208
L210:
	;
	goto L211
L211:
	;
	v1318 = int32(0)
	v1322 = v1318
	v1329 = v1318
	v1340 = v1310
	goto L212
L212:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(v122+v1322<<(uint(int32(2))%32))))
	v1350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349)+4)))
	if v1350 != int32(114) {
		goto L207
	} else {
		goto L214
	}
L213:
	;
	v1385 = v1371
	v1396 = v1372
	goto L208
L214:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1349)+5)))
	if v1353 != 0 {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v1374 = v1322 + int32(1)
	if v1374 != v109 {
		v1322 = v1374
		v1329 = v1371
		v1340 = v1372
		goto L212
	} else {
		goto L221
	}
L216:
	;
	v1371 = v1329
	v1372 = v1322
	goto L215
L217:
	;
	goto L218
L218:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+20))
	v1356 = F_make_one_partition_rbound(m, v67, v1322, v1354, int32(1))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L17
	} else {
		goto L219
	}
L219:
	;
	v1360 = v1313 + v1329<<(uint(int32(2))%32)
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v1349)+24))
	v1363 = F_make_one_partition_rbound(m, v67, v1322, v1361, int32(0))
	mBase = m.M
	v1364 = m.ExcPending
	if v1364 != 0 {
		goto L17
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1360)+4)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v1360))) = v1356
	v1371 = v1329 + int32(2)
	v1372 = v1340
	goto L215
L221:
	;
	goto L213
L222:
	;
	v1408 = F_palloc(m, v1385<<(uint(int32(2))%32))
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L17
	} else {
		goto L223
	}
L223:
	;
	if int32(0) < v1385 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1416 = int32(0)
	v1423 = v350
	v1428 = v3
	goto L227
L225:
	;
	v1574 = v350
	goto L226
L226:
	;
	F_pfree(m, v1313)
	mBase = m.M
	v1591 = m.ExcPending
	if v1591 != 0 {
		goto L17
	} else {
		goto L243
	}
L227:
	;
	v1441 = v1313 + v1428<<(uint(int32(2))%32)
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1441)))
	v1443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v1443 <= int32(0) {
		v1545 = v1423
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v1574 = v1545
	goto L226
L229:
	;
	v1562 = v1428 + int32(1)
	if v1562 != v1385 {
		v1416 = v1442
		v1423 = v1545
		v1428 = v1562
		goto L227
	} else {
		goto L242
	}
L230:
	;
	if v1416 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1454 = int32(0)
	goto L234
L232:
	;
	v1531 = v1442
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1408+v1423<<(uint(int32(2))%32)))) = v1531
	v1545 = v1423 + int32(1)
	goto L229
L234:
	;
	v1477 = v1454 << (uint(int32(2)) % 32)
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+8))
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1477+v1478)))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+8))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1481+v1477)))
	if v1480 != v1483 {
		goto L236
	} else {
		goto L237
	}
L235:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1441)))
	v1531 = v1504
	goto L233
L236:
	;
	goto L235
L237:
	;
	if v1480 != 0 {
		v1545 = v1423
		goto L229
	} else {
		goto L238
	}
L238:
	;
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v67)+28))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1489+v1477)))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v1442)+4))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1492+v1477)))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+4))
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(v1495+v1477)))
	v1498 = F_FunctionCall2Coll(m, v1485+v1454*int32(28), v1491, v1494, v1497)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L17
	} else {
		goto L239
	}
L239:
	;
	if v1498 != 0 {
		goto L236
	} else {
		goto L240
	}
L240:
	;
	v1501 = v1454 + int32(1)
	v1502 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	if v1501 < v1502 {
		v1454 = v1501
		goto L234
	} else {
		goto L241
	}
L241:
	;
	v1545 = v1423
	goto L229
L242:
	;
	goto L228
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+4)) = v1574
	v1594 = v1574 << (uint(int32(2)) % 32)
	v1595 = F_palloc0(m, v1594)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L17
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+8)) = v1595
	v1598 = F_palloc(m, v1594)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L17
	} else {
		goto L245
	}
L245:
	;
	v1601 = v1574 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+20)) = v1601
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+12)) = v1598
	v1608 = F_palloc(m, v1601<<(uint(int32(2))%32))
	mBase = m.M
	v1609 = m.ExcPending
	if v1609 != 0 {
		goto L17
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+24)) = v1608
	v1611 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	v1612 = v1594 * v1611
	v1613 = F_palloc(m, v1612)
	mBase = m.M
	v1614 = m.ExcPending
	if v1614 != 0 {
		goto L17
	} else {
		goto L247
	}
L247:
	;
	v1615 = F_palloc(m, v1612)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L17
	} else {
		goto L248
	}
L248:
	;
	v1617 = int32(0)
	if v1617 < v1574 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1620 = int32(0)
	v1635 = v1620
	v1647 = v3
	goto L252
L250:
	;
	v1787 = v1617
	v1809 = v3
	goto L251
L251:
	;
	F_pfree(m, v1408)
	mBase = m.M
	v1812 = m.ExcPending
	if v1812 != 0 {
		goto L17
	} else {
		goto L268
	}
L252:
	;
	v1649 = int32(2)
	v1650 = v1635 << (uint(v1649) % 32)
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+8))
	v1655 = v1611 * v1635 << (uint(v1649) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1650+v1651))) = v1613 + v1655
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1658+v1650))) = v1655 + v1615
	if base.B2i32(v1611 <= v1620) == int32(0) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1787 = v1574
	v1809 = v1778
	goto L251
L254:
	;
	v1664 = v1650 + v1408
	v1668 = int32(0)
	goto L257
L255:
	;
	goto L256
L256:
	;
	v1759 = *(*int32)(unsafe.Add(mBase, uint32(v1650+v1408)))
	v1760 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1759)+12)))
	if v1760 != 0 {
		v1775 = int32(-1)
		v1778 = v1647
		goto L264
	} else {
		goto L265
	}
L257:
	;
	v1693 = v1668 << (uint(int32(2)) % 32)
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1664)))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+8))
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1693+v1695)))
	if v1697 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	goto L256
L259:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1694)+4))
	v1702 = *(*int32)(unsafe.Add(mBase, uint32(v1700+v1693)))
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v1705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1703+v1668))))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v1710 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1706+v1668<<(uint(int32(1))%32)))))
	v1711 = F_datumCopy(m, v1702, v1705, v1710)
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L17
	} else {
		goto L262
	}
L260:
	;
	v1722 = v1697
	goto L261
L261:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+12))
	v1725 = *(*int32)(unsafe.Add(mBase, uint32(v1723+v1650)))
	*(*int32)(unsafe.Add(mBase, uint32(v1725+v1693))) = v1722
	v1729 = v1668 + int32(1)
	if v1729 != v1611 {
		v1668 = v1729
		goto L257
	} else {
		goto L263
	}
L262:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+8))
	v1715 = *(*int32)(unsafe.Add(mBase, uint32(v1713+v1650)))
	*(*int32)(unsafe.Add(mBase, uint32(v1715+v1693))) = v1711
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v1664)))
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1718)+8))
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1719+v1693)))
	v1722 = v1721
	goto L261
L263:
	;
	goto L258
L264:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1779+v1650))) = v1775
	v1783 = v1635 + int32(1)
	if v1783 != v1574 {
		v1635 = v1783
		v1647 = v1778
		goto L252
	} else {
		goto L267
	}
L265:
	;
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(v1759)))
	v1763 = v1761 << (uint(int32(2)) % 32)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1765 = v1763 + v1764
	v1766 = *(*int32)(unsafe.Add(mBase, uint32(v1765)))
	if v1766 != int32(-1) {
		v1775 = v1766
		v1778 = v1647
		goto L264
	} else {
		goto L266
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1765))) = v1647
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1772+v1763)))
	v1775 = v1774
	v1778 = v1647 + int32(1)
	goto L264
L267:
	;
	goto L253
L268:
	;
	if v1396 != int32(-1) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1816 = v1396 << (uint(int32(2)) % 32)
	v1817 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v1816+v1817))) = v1809
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1820+v1816)))
	*(*int32)(unsafe.Add(mBase, uint32(v1304)+32)) = v1822
	goto L271
L270:
	;
	goto L271
L271:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(v1304)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1825+v1787<<(uint(int32(2))%32)))) = int32(-1)
	v2126 = v1304
	goto L108
L272:
	;
	F_errmsg_internal(m, int32(475006), int32(0))
	mBase = m.M
	v1838 = m.ExcPending
	if v1838 != 0 {
		goto L17
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(477464), int32(713), int32(165076))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L17
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
	*(*int32)(unsafe.Add(mBase, uint32(v920)+4)) = v1848
	v1876 = v1848 << (uint(int32(2)) % 32)
	v1877 = F_palloc0(m, v1876)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L17
	} else {
		goto L276
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+20)) = v1848
	*(*int64)(unsafe.Add(mBase, uint32(v920)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v920)+8)) = v1877
	v1883 = F_palloc(m, v1876)
	mBase = m.M
	v1884 = m.ExcPending
	if v1884 != 0 {
		goto L17
	} else {
		goto L277
	}
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+24)) = v1883
	v1886 = F_palloc(m, v1876)
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L17
	} else {
		goto L278
	}
L278:
	;
	if v1848 <= int32(0) {
		goto L280
	} else {
		goto L281
	}
L279:
	;
	F_pfree(m, v1854)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L17
	} else {
		goto L290
	}
L280:
	;
	v1970 = int32(0)
	goto L279
L281:
	;
	goto L282
L282:
	;
	v1891 = int32(0)
	v1895 = v1891
	v1902 = v1891
	goto L283
L283:
	;
	v1921 = v1854 + v1895<<(uint(int32(3))%32)
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v1921)))
	v1924 = v1895 << (uint(int32(2)) % 32)
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v920)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1924+v1925))) = v1924 + v1886
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1921)+4))
	v1930 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v1931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1930))))
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v1933 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1932))))
	v1934 = F_datumCopy(m, v1929, v1931, v1933)
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L17
	} else {
		goto L285
	}
L284:
	;
	v1970 = v1954
	goto L279
L285:
	;
	v1936 = *(*int32)(unsafe.Add(mBase, uint32(v920)+8))
	v1938 = *(*int32)(unsafe.Add(mBase, uint32(v1936+v1924)))
	*(*int32)(unsafe.Add(mBase, uint32(v1938))) = v1934
	v1941 = v1922 << (uint(int32(2)) % 32)
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1943 = v1941 + v1942
	v1944 = *(*int32)(unsafe.Add(mBase, uint32(v1943)))
	if v1944 == int32(-1) {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1943))) = v1902
	v1950 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1950+v1941)))
	v1953 = v1952
	v1954 = v1902 + int32(1)
	goto L288
L287:
	;
	v1953 = v1944
	v1954 = v1902
	goto L288
L288:
	;
	v1955 = *(*int32)(unsafe.Add(mBase, uint32(v920)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1955+v1924))) = v1953
	v1959 = v1895 + int32(1)
	if v1959 != v1848 {
		v1895 = v1959
		v1902 = v1954
		goto L283
	} else {
		goto L289
	}
L289:
	;
	goto L284
L290:
	;
	if v1856 != int32(-1) {
		goto L291
	} else {
		goto L292
	}
L291:
	;
	v1992 = v1856 << (uint(int32(2)) % 32)
	v1993 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v1994 = v1992 + v1993
	v1995 = *(*int32)(unsafe.Add(mBase, uint32(v1994)))
	if v1995 == int32(-1) {
		goto L294
	} else {
		goto L295
	}
L292:
	;
	v2010 = v1970
	goto L293
L293:
	;
	if v1860 != int32(-1) {
		goto L297
	} else {
		goto L298
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1994))) = v1970
	v2001 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v2001+v1992)))
	v2004 = v1970 + int32(1)
	v2005 = v2003
	goto L296
L295:
	;
	v2004 = v1970
	v2005 = v1995
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+28)) = v2005
	v2010 = v2004
	goto L293
L297:
	;
	v2014 = v1860 << (uint(int32(2)) % 32)
	v2015 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	*(*int32)(unsafe.Add(mBase, uint32(v2014+v2015))) = v2010
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v355)))
	v2020 = *(*int32)(unsafe.Add(mBase, uint32(v2018+v2014)))
	*(*int32)(unsafe.Add(mBase, uint32(v920)+32)) = v2020
	goto L299
L298:
	;
	goto L299
L299:
	;
	if v109 < int32(2) {
		v2126 = v920
		goto L108
	} else {
		goto L300
	}
L300:
	;
	v2025 = int32(-1)
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(v920)+4))
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(v920)+28))
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(v920)+32))
	if v2026+base.B2i32(v2027 != v2025)+base.B2i32(v2031 != v2025) == v109 {
		v2092 = v2031
		goto L301
	} else {
		goto L302
	}
L301:
	;
	if v2092 == int32(-1) {
		v2126 = v920
		goto L108
	} else {
		goto L314
	}
L302:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v920)+20))
	if v2036 <= int32(0) {
		v2092 = v2031
		goto L301
	} else {
		goto L303
	}
L303:
	;
	v2042 = v2025
	v2044 = int32(0)
	v2045 = v2036
	goto L304
L304:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v920)+24))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v2066+v2044<<(uint(int32(2))%32))))
	if v2042 <= v2070 {
		goto L307
	} else {
		goto L308
	}
L305:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(v920)+32))
	v2092 = v2087
	goto L301
L306:
	;
	v2085 = v2044 + int32(1)
	if v2085 < v2083 {
		v2042 = v2070
		v2044 = v2085
		v2045 = v2083
		goto L304
	} else {
		goto L313
	}
L307:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v920)+28))
	if v2072 == int32(-1) {
		v2083 = v2045
		goto L306
	} else {
		goto L310
	}
L308:
	;
	goto L309
L309:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(v920)+16))
	v2078 = F_bms_add_member(m, v2077, v2070)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L17
	} else {
		goto L312
	}
L310:
	;
	if v2070 != v2072 {
		v2083 = v2045
		goto L306
	} else {
		goto L311
	}
L311:
	;
	goto L309
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+16)) = v2078
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v920)+20))
	v2083 = v2081
	goto L306
L313:
	;
	goto L305
L314:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v920)+16))
	v2117 = F_bms_add_member(m, v2116, v2092)
	mBase = m.M
	v2118 = m.ExcPending
	if v2118 != 0 {
		goto L17
	} else {
		goto L315
	}
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v920)+16)) = v2117
	v2172 = v920
	goto L94
L316:
	;
	v2175 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)) = uint8(v2175)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v2175
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2186 = F_find_inheritance_children_extended(m, v2180, l1, v2175, v29+int32(127), v29+int32(120))
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L17
	} else {
		goto L317
	}
L317:
	;
	if v2186 != 0 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2191 = int32(0)
	v2192 = *(*int32)(unsafe.Add(mBase, uint32(v2186)+4))
	if v2192 <= v2191 {
		v2236 = v2175
		v2242 = v2192
		v2245 = v118
		v2246 = v120
		v2249 = v2191
		goto L21
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	goto L30
L321:
	;
	v96 = int32(1)
	v98 = v2186
	v109 = v2192
	v110 = v2186 + int32(4)
	goto L29
L322:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2261 = F_MemoryContextStrdup(m, v2256, v2258+int32(4))
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L17
	} else {
		goto L323
	}
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+36)) = v2261
	goto L324
L324:
	;
	v2265 = F_MemoryContextAllocZero(m, v2256, int32(32))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L17
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2265))) = v2242
	v2268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2265)+4)) = uint8(v2268)
	if v2236 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2270 = int32(4455216)
	v2271 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v2256
	v2276 = F_palloc(m, int32(36))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L17
	} else {
		goto L329
	}
L327:
	;
	goto L328
L328:
	;
	if l1 == int32(0) {
		goto L389
	} else {
		goto L390
	}
L329:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2249)))
	*(*int32)(unsafe.Add(mBase, uint32(v2276))) = v2278
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+4)) = v2280
	v2282 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+20)) = v2282
	v2284 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+4)))
	v2286 = v2280 << (uint(int32(2)) % 32)
	v2287 = F_palloc(m, v2286)
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L17
	} else {
		goto L330
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+8)) = v2287
	v2290 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+12))
	if v2290 != 0 {
		goto L332
	} else {
		goto L333
	}
L331:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+16))
	v2376 = F_bms_copy(m, v2375)
	mBase = m.M
	v2377 = m.ExcPending
	if v2377 != 0 {
		goto L17
	} else {
		goto L345
	}
L332:
	;
	v2291 = F_palloc(m, v2286)
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L17
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = int32(0)
	goto L331
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+12)) = v2291
	v2295 = F_palloc(m, v2284*v2286)
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L17
	} else {
		goto L336
	}
L336:
	;
	if v2280 <= int32(0) {
		goto L331
	} else {
		goto L337
	}
L337:
	;
	v2300 = v2284 << (uint(int32(2)) % 32)
	v2309 = int32(0)
	goto L338
L338:
	;
	v2327 = int32(2)
	v2328 = v2309 << (uint(v2327) % 32)
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2328+v2329))) = v2295 + v2309*v2284<<(uint(v2327)%32)
	v2336 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v2336+v2328)))
	v2339 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+12))
	v2341 = *(*int32)(unsafe.Add(mBase, uint32(v2339+v2328)))
	if v2300 != 0 {
		goto L341
	} else {
		goto L342
	}
L339:
	;
	goto L331
L340:
	;
	v2345 = v2309 + int32(1)
	if v2345 != v2280 {
		v2309 = v2345
		goto L338
	} else {
		goto L344
	}
L341:
	;
	v2342 = F__emscripten_memcpy_bulkmem(m, v2338, v2341, v2300)
	mBase = m.M
	goto L343
L342:
	;
	goto L343
L343:
	;
	goto L340
L344:
	;
	goto L339
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+16)) = v2376
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v2382 = base.B2i32(v2380 == int32(104))
	if v2380 == int32(104) {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v2383 = int32(2)
	goto L348
L347:
	;
	v2383 = v2284
	goto L348
L348:
	;
	v2385 = F_palloc(m, v2286*v2383)
	mBase = m.M
	v2386 = m.ExcPending
	if v2386 != 0 {
		goto L17
	} else {
		goto L349
	}
L349:
	;
	if int32(0) < v2280 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v2389 = int32(0)
	v2404 = v2389
	goto L353
L351:
	;
	goto L352
L352:
	;
	v2552 = v2282 << (uint(int32(2)) % 32)
	v2553 = F_palloc(m, v2552)
	mBase = m.M
	v2554 = m.ExcPending
	if v2554 != 0 {
		goto L17
	} else {
		goto L372
	}
L353:
	;
	v2418 = int32(2)
	v2419 = v2404 << (uint(v2418) % 32)
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2419+v2420))) = v2385 + v2383*v2404<<(uint(v2418)%32)
	v2427 = int32(0)
	if base.B2i32(v2383 <= v2389) == v2427 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	goto L352
L355:
	;
	v2438 = v2427
	goto L358
L356:
	;
	goto L357
L357:
	;
	v2523 = v2404 + int32(1)
	if v2523 != v2280 {
		v2404 = v2523
		goto L353
	} else {
		goto L371
	}
L358:
	;
	if v2380 == int32(104) {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	goto L357
L360:
	;
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+12))
	if v2468 != 0 {
		goto L365
	} else {
		goto L366
	}
L361:
	;
	v2466 = int32(1)
	v2467 = int32(4)
	goto L360
L362:
	;
	goto L363
L363:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
	v2460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2458+v2438))))
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v67)+40))
	v2465 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2461+v2438<<(uint(int32(1))%32)))))
	v2466 = v2460
	v2467 = v2465
	goto L360
L364:
	;
	v2494 = v2438 + int32(1)
	if v2494 != v2383 {
		v2438 = v2494
		goto L358
	} else {
		goto L370
	}
L365:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2468+v2419)))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2470+v2438<<(uint(int32(2))%32))))
	if v2474 != 0 {
		goto L364
	} else {
		goto L368
	}
L366:
	;
	goto L367
L367:
	;
	v2476 = v2438 << (uint(int32(2)) % 32)
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+8))
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(v2477+v2419)))
	v2481 = *(*int32)(unsafe.Add(mBase, uint32(v2476+v2479)))
	v2484 = F_datumCopy(m, v2481, v2466&int32(1), v2467)
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L17
	} else {
		goto L369
	}
L368:
	;
	goto L367
L369:
	;
	v2486 = *(*int32)(unsafe.Add(mBase, uint32(v2276)+8))
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2486+v2419)))
	*(*int32)(unsafe.Add(mBase, uint32(v2488+v2476))) = v2484
	goto L364
L370:
	;
	goto L359
L371:
	;
	goto L354
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+24)) = v2553
	v2556 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+24))
	if v2552 != 0 {
		goto L374
	} else {
		goto L375
	}
L373:
	;
	v2559 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+28)) = v2559
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(v2249)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v2276)+32)) = v2561
	*(*int32)(unsafe.Add(mBase, uint32(v2265)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2265)+20)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v2265)+16)) = v2276
	v2570 = F_palloc(m, v2242<<(uint(int32(2))%32))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		goto L17
	} else {
		goto L377
	}
L374:
	;
	v2557 = F__emscripten_memcpy_bulkmem(m, v2553, v2556, v2552)
	mBase = m.M
	goto L376
L375:
	;
	goto L376
L376:
	;
	goto L373
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2265)+8)) = v2570
	v2573 = F_palloc(m, v2242)
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L17
	} else {
		goto L378
	}
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2265)+12)) = v2573
	if v2242 <= int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v2271
	goto L328
L380:
	;
	v2578 = int32(1)
	v2580 = int32(0)
	if v2242 != v2578 {
		goto L381
	} else {
		goto L382
	}
L381:
	;
	v2588 = int32(0)
	v2599 = v2580
	goto L384
L382:
	;
	v2666 = v2580
	goto L383
L383:
	;
	if v2242&v2578 == int32(0) {
		goto L379
	} else {
		goto L387
	}
L384:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+8))
	v2613 = int32(2)
	v2614 = v2599 << (uint(v2613) % 32)
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(v2614+v2615)))
	v2622 = *(*int32)(unsafe.Add(mBase, uint32(v2614+v2245)))
	*(*int32)(unsafe.Add(mBase, uint32(v2612+v2617<<(uint(v2613)%32)))) = v2622
	v2624 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+12))
	v2627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2599+v2246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2617+v2624))) = uint8(v2627)
	v2629 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+8))
	v2631 = v2599 | int32(1)
	v2633 = v2631 << (uint(v2613) % 32)
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2633+v2634)))
	v2641 = *(*int32)(unsafe.Add(mBase, uint32(v2633+v2245)))
	*(*int32)(unsafe.Add(mBase, uint32(v2629+v2636<<(uint(v2613)%32)))) = v2641
	v2643 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+12))
	v2646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2631+v2246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2636+v2643))) = uint8(v2646)
	v2649 = v2599 + v2613
	v2651 = v2588 + v2613
	if v2651 != v2242&int32(2147483646) {
		v2588 = v2651
		v2599 = v2649
		goto L384
	} else {
		goto L386
	}
L385:
	;
	v2666 = v2649
	goto L383
L386:
	;
	goto L385
L387:
	;
	v2681 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+8))
	v2682 = int32(2)
	v2683 = v2666 << (uint(v2682) % 32)
	v2684 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v2686 = *(*int32)(unsafe.Add(mBase, uint32(v2683+v2684)))
	v2691 = *(*int32)(unsafe.Add(mBase, uint32(v2683+v2245)))
	*(*int32)(unsafe.Add(mBase, uint32(v2681+v2686<<(uint(v2682)%32)))) = v2691
	v2693 = *(*int32)(unsafe.Add(mBase, uint32(v2265)+12))
	v2696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2666+v2246))))
	*(*uint8)(unsafe.Add(mBase, uint32(v2686+v2693))) = uint8(v2696)
	goto L379
L388:
	;
	v2875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v2875 != 0 {
		goto L449
	} else {
		goto L450
	}
L389:
	;
	v2840 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+16))
	if v2844 != v2840 {
		goto L433
	} else {
		goto L434
	}
L390:
	;
	v2754 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+127)))
	if v2754 != int32(1) {
		goto L389
	} else {
		goto L391
	}
L391:
	;
	v2758 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	goto L392
L392:
	;
	if base.B2i32(v2758 != int32(0)) == int32(0) {
		goto L389
	} else {
		goto L393
	}
L393:
	;
	v2763 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	v2765 = *(*int32)(unsafe.Add(mBase, _consts[499]))
	v2769 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+16))
	if v2769 != v2765 {
		goto L395
	} else {
		goto L396
	}
L394:
	;
	if v2763 == int32(0) {
		goto L388
	} else {
		goto L411
	}
L395:
	;
	if v2769 == int32(0) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	goto L397
L397:
	;
	goto L394
L398:
	;
	if v2765 != 0 {
		goto L405
	} else {
		goto L406
	}
L399:
	;
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+28))
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+24))
	if v2774 != 0 {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	if v2773 == int32(0) {
		goto L398
	} else {
		goto L404
	}
L401:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2774)+28)) = v2773
	goto L400
L402:
	;
	goto L403
L403:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+20)) = v2773
	goto L400
L404:
	;
	v2779 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2773)+24)) = v2779
	goto L398
L405:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+16)) = v2765
	v2786 = *(*int32)(unsafe.Add(mBase, uint32(v2765)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+28)) = v2786
	if v2786 != 0 {
		goto L408
	} else {
		goto L409
	}
L406:
	;
	goto L407
L407:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2256)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+16)) = int32(0)
	goto L397
L408:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2786)+24)) = v2256
	goto L410
L409:
	;
	goto L410
L410:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2765)+20)) = v2256
	goto L394
L411:
	;
	v2801 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v2801 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v2801)+16))
	if v2805 != v2256 {
		goto L416
	} else {
		goto L417
	}
L413:
	;
	goto L414
L414:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v2265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v2256
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v29)+120))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+116)) = v2837
	v2914 = v2265
	goto L4
L415:
	;
	goto L414
L416:
	;
	if v2805 == int32(0) {
		goto L419
	} else {
		goto L420
	}
L417:
	;
	goto L418
L418:
	;
	goto L415
L419:
	;
	if v2256 != 0 {
		goto L426
	} else {
		goto L427
	}
L420:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(v2801)+28))
	v2810 = *(*int32)(unsafe.Add(mBase, uint32(v2801)+24))
	if v2810 != 0 {
		goto L422
	} else {
		goto L423
	}
L421:
	;
	if v2809 == int32(0) {
		goto L419
	} else {
		goto L425
	}
L422:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2810)+28)) = v2809
	goto L421
L423:
	;
	goto L424
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2805)+20)) = v2809
	goto L421
L425:
	;
	v2815 = *(*int32)(unsafe.Add(mBase, uint32(v2801)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2809)+24)) = v2815
	goto L419
L426:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2801)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2801)+16)) = v2256
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2801)+28)) = v2822
	if v2822 != 0 {
		goto L429
	} else {
		goto L430
	}
L427:
	;
	goto L428
L428:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2801)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2801)+16)) = int32(0)
	goto L418
L429:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2822)+24)) = v2801
	goto L431
L430:
	;
	goto L431
L431:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+20)) = v2801
	goto L415
L432:
	;
	goto L388
L433:
	;
	if v2844 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L434:
	;
	goto L435
L435:
	;
	goto L432
L436:
	;
	if v2840 != 0 {
		goto L443
	} else {
		goto L444
	}
L437:
	;
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+28))
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+24))
	if v2849 != 0 {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	if v2848 == int32(0) {
		goto L436
	} else {
		goto L442
	}
L439:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2849)+28)) = v2848
	goto L438
L440:
	;
	goto L441
L441:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2844)+20)) = v2848
	goto L438
L442:
	;
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2848)+24)) = v2854
	goto L436
L443:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+16)) = v2840
	v2861 = *(*int32)(unsafe.Add(mBase, uint32(v2840)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+28)) = v2861
	if v2861 != 0 {
		goto L446
	} else {
		goto L447
	}
L444:
	;
	goto L445
L445:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2256)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+16)) = int32(0)
	goto L435
L446:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2861)+24)) = v2256
	goto L448
L447:
	;
	goto L448
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2840)+20)) = v2256
	goto L432
L449:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+16))
	if v2879 != v2256 {
		goto L453
	} else {
		goto L454
	}
L450:
	;
	goto L451
L451:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+100)) = v2265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v2256
	v2914 = v2265
	goto L4
L452:
	;
	goto L451
L453:
	;
	if v2879 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L454:
	;
	goto L455
L455:
	;
	goto L452
L456:
	;
	if v2256 != 0 {
		goto L463
	} else {
		goto L464
	}
L457:
	;
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+28))
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+24))
	if v2884 != 0 {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	if v2883 == int32(0) {
		goto L456
	} else {
		goto L462
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2884)+28)) = v2883
	goto L458
L460:
	;
	goto L461
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2879)+20)) = v2883
	goto L458
L462:
	;
	v2889 = *(*int32)(unsafe.Add(mBase, uint32(v2875)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2883)+24)) = v2889
	goto L456
L463:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2875)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2875)+16)) = v2256
	v2896 = *(*int32)(unsafe.Add(mBase, uint32(v2256)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2875)+28)) = v2896
	if v2896 != 0 {
		goto L466
	} else {
		goto L467
	}
L464:
	;
	goto L465
L465:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2875)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2875)+16)) = int32(0)
	goto L455
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2896)+24)) = v2875
	goto L468
L467:
	;
	goto L468
L468:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2256)+20)) = v2875
	goto L452
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v159
	F_errmsg_internal(m, int32(44638), v29)
	mBase = m.M
	v2948 = m.ExcPending
	if v2948 != 0 {
		goto L17
	} else {
		goto L470
	}
L470:
	;
	F_errfinish(m, int32(482861), int32(280), int32(472110))
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L17
	} else {
		goto L471
	}
L471:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L472:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = v159
	F_errmsg_internal(m, int32(44675), v29+int32(32))
	mBase = m.M
	v2963 = m.ExcPending
	if v2963 != 0 {
		goto L17
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(482861), int32(282), int32(472110))
	mBase = m.M
	v2968 = m.ExcPending
	if v2968 != 0 {
		goto L17
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v305
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v159
	F_errmsg_internal(m, int32(38929), v29+int32(16))
	mBase = m.M
	v2979 = m.ExcPending
	if v2979 != 0 {
		goto L17
	} else {
		goto L476
	}
L476:
	;
	F_errfinish(m, int32(482861), int32(296), int32(472110))
	mBase = m.M
	v2984 = m.ExcPending
	if v2984 != 0 {
		goto L17
	} else {
		goto L477
	}
L477:
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
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1072]))
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
					v52 = *(*int32)(unsafe.Add(mBase, _consts[170]))
					F_ResourceOwnerEnlarge(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v55 + int32(1)
						v60 = *(*int32)(unsafe.Add(mBase, _consts[84]))
						if v60 != 0 {
							v62 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							F_ResourceOwnerRemember(m, v62, v47, int32(1706432))
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
						v52 = *(*int32)(unsafe.Add(mBase, _consts[170]))
						F_ResourceOwnerEnlarge(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
							*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v55 + int32(1)
							v60 = *(*int32)(unsafe.Add(mBase, _consts[84]))
							if v60 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, _consts[170]))
								F_ResourceOwnerRemember(m, v62, v47, int32(1706432))
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
					v27 = *(*int32)(unsafe.Add(mBase, _consts[170]))
					F_ResourceOwnerEnlarge(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v30 + int32(1)
						v35 = *(*int32)(unsafe.Add(mBase, _consts[84]))
						if v35 != 0 {
							v37 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							F_ResourceOwnerRemember(m, v37, v22, int32(1706432))
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
						F_errmsg_internal(m, int32(103383), v52)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(480861), int32(38), int32(360282))
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
					F_errmsg_internal(m, int32(103383), v52)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(480861), int32(38), int32(360282))
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
		if base.Ui32(v13) < base.Ui32(int32(12000)) {
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
							F_errmsg_internal(m, int32(103383), v52)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(480861), int32(38), int32(360282))
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
						F_errmsg_internal(m, int32(103383), v52)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return
						} else {
							F_errfinish(m, int32(480861), int32(38), int32(360282))
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
										F_errmsg_internal(m, int32(103383), v52)
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_errfinish(m, int32(480861), int32(38), int32(360282))
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
									F_errmsg_internal(m, int32(103383), v52)
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_errfinish(m, int32(480861), int32(38), int32(360282))
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
						F_errmsg_internal(m, int32(51623), v7)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_errfinish(m, int32(482315), int32(1863), int32(408086))
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
	var v19 int32
	_ = v19
	v3 = int32(1)
	if l0 <= int32(2963) {
		if base.Ui32(l0-int32(2608)) < base.Ui32(int32(2)) {
			v19 = v3
		} else {
			if l0 == int32(1214) {
				v19 = v3
			} else {
				if l0 != int32(2396) {
					v19 = int32(0)
				} else {
					v19 = v3
				}
			}
		}
	} else {
		switch l0 - int32(3592) {
		case 0, 4:
			v19 = v3
		case 1, 2, 3:
			v19 = int32(0)
		default:
			if l0 == int32(2964) {
				v19 = v3
			} else {
				v19 = int32(0)
			}
		}
	}
	return v19
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = int32(0)
	v12 = v7
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
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
		v11 = v30
		v12 = v13
		goto L4
	} else {
		goto L25
	}
L7:
	;
	v30 = v12
	goto L6
L8:
	;
	goto L9
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v17 != v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v30 = v12
	goto L6
L11:
	;
	goto L12
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v20 != v21 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v30 = v12
	goto L6
L14:
	;
	goto L15
L15:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if l1 != v23 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v30 = v12
	goto L6
L17:
	;
	goto L18
L18:
	;
	if v11 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	F_pfree(m, v12)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v13
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[138])) = v13
	goto L19
L23:
	;
	return
L24:
	;
	v30 = v11
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
		v9 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9+(l0^int32(-1))<<(uint(int32(2))%32))))
		v23 = v15
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
				v33 = *(*int32)(unsafe.Add(mBase, _consts[6]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v33+(l0^int32(-1))<<(uint(int32(6))%32))+16))
				v48 = v39
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, _consts[7]))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v41+l0<<(uint(int32(6))%32)+int32(-64))+16))
				v48 = v47
			}
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v28)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v48)
			v52 = int32(base.Ui32(v48) >> (uint(int32(16)) % 32))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
			if l2 == int32(0) {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v28<<(uint(int32(2))%32)+v23)+20))
				v62 = v23 + v59&int32(32767)
				v64 = l1 + int32(4)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
				*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v65
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)) = uint16(v67)
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
				F_errmsg_internal(m, int32(393819), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errfinish(m, int32(479301), int32(65), int32(370872))
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
			F_sequence_close(m, v4, int32(0))
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
	var v40 int32
	_ = v40
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
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
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
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
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v165 < v157 {
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
	v40 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L18
	}
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v60 = v26 + int32(20) + v54<<(uint(int32(4))%32) + v40*int32(100)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+91)))
	if v61 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v66 = v4
	goto L15
L14:
	;
	v64 = F_pstrdup(m, v60+int32(4))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30+v40<<(uint(int32(2))%32)))) = v66
	v69 = v40 + int32(1)
	if v69 != v27 {
		v40 = v69
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v66 = v64
	goto L15
L17:
	;
	goto L12
L18:
	;
	v157 = v27
	v161 = v30
	goto L1
L19:
	;
	if v107 != 0 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v105
	v107 = v105
	goto L19
L21:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v90 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v93 = int32(1)
	v94 = int32(0)
	F_expandRTE(m, l1, v93, v94, v94, int32(-1), v93, v17+int32(12), v94)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v107 = v103
	goto L19
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v110 = v108
	goto L26
L25:
	;
	v110 = int32(0)
	goto L26
L26:
	;
	v113 = F_palloc(m, v110<<(uint(int32(2))%32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v115 == int32(0) {
		v157 = v110
		v161 = v113
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v118 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v119 <= v118 {
		v157 = v110
		v161 = v113
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v125 = v118
	goto L30
L30:
	;
	v137 = v125 << (uint(int32(2)) % 32)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v139+v137)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+4))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v144 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v157 = v110
	v161 = v113
	goto L1
L32:
	;
	v145 = v142
	goto L34
L33:
	;
	v145 = int32(0)
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113+v137))) = v145
	v148 = v125 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v148 < v149 {
		v125 = v148
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v167 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	v187 = F_palloc(m, v157<<(uint(int32(2))%32))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L5
	} else {
		goto L45
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v157
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v180
	goto L38
L40:
	;
	v172 = F_palloc0(m, v157<<(uint(int32(2))%32))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v174 = int32(2)
	v178 = F_repalloc0(m, v167, v165<<(uint(v174)%32), v157<<(uint(v174)%32))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L5
	} else {
		goto L44
	}
L43:
	;
	v180 = v172
	goto L39
L44:
	;
	v180 = v178
	goto L39
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v187
	v190 = F_palloc(m, v157)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v190
	F_build_colinfo_names_hash(m, l2)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	v195 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	if v197 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	v199 = v198
	goto L50
L49:
	;
	v199 = v4
	goto L50
L50:
	;
	if v157 <= int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v320 != 0 {
		goto L82
	} else {
		goto L83
	}
L52:
	;
	v310 = v195
	v313 = int32(0)
	goto L51
L53:
	;
	goto L54
L54:
	;
	v203 = int32(0)
	v208 = v203
	v209 = v195
	v212 = v203
	goto L55
L55:
	;
	v220 = v208 << (uint(int32(2)) % 32)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v161+v220)))
	if v222 == int32(0) {
		v298 = v209
		v300 = v212
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v310 = v298
	v313 = v300
	goto L51
L57:
	;
	v304 = v208 + int32(1)
	if v304 != v157 {
		v208 = v304
		v209 = v298
		v212 = v300
		goto L55
	} else {
		goto L81
	}
L58:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225+v220)))
	if v227 != 0 {
		v254 = v227
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v257+v209<<(uint(int32(2))%32)))) = v254
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v262+v209))) = uint8(base.B2i32(v199 <= v208))
	v266 = int32(1)
	v267 = v209 + v266
	if v212&v266 != 0 {
		goto L70
	} else {
		goto L71
	}
L60:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v228 == int32(0) {
		v240 = v222
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v242 = F_make_colname_unique(m, v240, l0, l2)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L67
	}
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	if v231 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v234 = v232
	goto L65
L64:
	;
	v234 = int32(0)
	goto L65
L65:
	;
	if v234 <= v208 {
		v240 = v222
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236+v220)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v240 = v239
	goto L61
L67:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v244+v220))) = v242
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	if v247 == int32(0) {
		v254 = v242
		goto L59
	} else {
		goto L68
	}
L68:
	;
	v252 = F_hash_search(m, v247, v242, int32(1), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	v254 = v242
	goto L59
L70:
	;
	v298 = v267
	v300 = int32(1)
	goto L57
L71:
	;
	goto L72
L72:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if v274 == int32(0) {
		v293 = v273
		v294 = v274
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v298 = v267
	v300 = base.B2i32(v294-v293 != int32(0))
	goto L57
L74:
	;
	goto L73
L75:
	;
	if v273 != v274 {
		v293 = v273
		v294 = v274
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v278 = v254
	v279 = v222
	goto L77
L77:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	if v283 == int32(0) {
		v293 = v282
		v294 = v283
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v293 = v282
	v294 = v283
	goto L74
L79:
	;
	v286 = int32(1)
	if v282 == v283 {
		v278 = v278 + v286
		v279 = v279 + v286
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	goto L56
L82:
	;
	F_hash_destroy(m, v320)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v310
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	switch v326 {
	case 0:
		goto L90
	default:
		goto L87
	case 3:
		goto L89
	case 4:
		goto L88
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+48)) = int32(0)
	goto L84
L86:
	;
	m.G0 = v17 + int32(16)
	return
L87:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v334 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	v332 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v332)
	goto L86
L89:
	;
	v330 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v330)
	goto L86
L90:
	;
	v328 = v313 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v328)
	goto L86
L91:
	;
	v343 = v313 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v343)
	goto L86
L92:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	if v337 == int32(0) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v340 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)) = uint8(v340)
	goto L86
}
