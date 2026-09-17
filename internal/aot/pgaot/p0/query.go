package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeQueryCompletion(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return
}
func F_ScanQueryForLocks(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
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
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v2 = l1
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)) = uint8(v2)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v14 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v67 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v3
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v23<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	switch v33 {
	case 0:
		goto L8
	case 1:
		goto L7
	default:
		goto L6
	}
L5:
	;
	goto L1
L6:
	;
	v56 = v23 + int32(1)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v56 < v57 {
		v23 = v56
		goto L4
	} else {
		goto L23
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v40 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_LockRelationOid(m, v35, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	F_UnlockRelationOid(m, v35, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L14
	}
L12:
	;
	return
L13:
	;
	goto L6
L14:
	;
	goto L6
L15:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	F_ScanQueryForLocks(m, v49, v2)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L22
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	if v2 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_LockRelationOid(m, v40, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_UnlockRelationOid(m, v40, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	goto L15
L22:
	;
	goto L6
L23:
	;
	goto L5
L24:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v102 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L25:
	;
	v70 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v71 <= v70 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v77 = v70
	goto L27
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82+v77<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	F_ScanQueryForLocks(m, v87, v2)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L29
	}
L28:
	;
	goto L24
L29:
	;
	v91 = v77 + int32(1)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v91 < v92 {
		v77 = v91
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v109 = F_query_tree_walker_impl(m, l0, int32(1588), v11+int32(15), int32(3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	m.G0 = v11 + int32(16)
	return
L34:
	;
	goto L33
}
func F_assign_query_collations_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(142) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v17 <= int32(0) {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v53 = F_assign_collations_walker(m, l0, v8+int32(8))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v24<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v38 = F_assign_collations_walker(m, v30, v8+int32(8))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L1
L10:
	;
	return int32(0)
L11:
	;
	v43 = v24 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v43 < v44 {
		v24 = v43
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	goto L1
}
func F_extract_query_dependencies(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v5 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(480)
	m.G0 = v8
	base.MemoryFill(m, v8+int32(392), v5, int32(88))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+448)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v8)+388)) = int32(266)
	base.MemoryFill(m, v8, v5, int32(384))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(267)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v8 + int32(388)
	v27 = F_extract_query_dependencies_walker(m, l0, v8)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+444))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v29
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+448))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v31
		v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+469)))
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v33)
		m.G0 = v8 + int32(480)
		return
	}
}
func F_query_is_distinct_for(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
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
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
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
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v13 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v344
L2:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v108 != 0 {
		goto L24
	} else {
		goto L25
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= int32(0) {
		v344 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v24 = v4
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v24<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = F_get_sortgroupclause_tle(m, v36, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v344 = v91
	goto L1
L7:
	;
	return int32(0)
L8:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+8)))
	v47 = int32(0)
	goto L9
L9:
	;
	v56 = int32(0)
	if l1 == v56 {
		v65 = v56
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v73+v47<<(uint(int32(2))%32))))
	if v83 == int32(0) {
		goto L2
	} else {
		goto L20
	}
L11:
	;
	if l2 == int32(0) {
		goto L2
	} else {
		goto L14
	}
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v59 <= v47 {
		v65 = v56
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v65 = v61 + v47<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.B2i32(v65 == int32(0))|base.B2i32(v70 <= v47) != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v73 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v42 != v76 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = v47 + int32(1)
	goto L9
L18:
	;
	goto L19
L19:
	;
	goto L10
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v87 = F_equality_ops_are_compatible(m, v83, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	if v87 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v91 = int32(1)
	v93 = v24 + v91
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v93 < v94 {
		v24 = v93
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	return int32(0)
L25:
	;
	goto L26
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v112 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v220 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L28:
	;
	if v111 != 0 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	if v111 != 0 {
		goto L53
	} else {
		goto L54
	}
L31:
	;
	return int32(0)
L32:
	;
	goto L33
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v116 <= int32(0) {
		v344 = int32(1)
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v124 = int32(0)
	goto L35
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v124<<(uint(int32(2))%32))))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v138 = F_get_sortgroupclause_tle(m, v136, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L37
	}
L36:
	;
	v344 = v189
	goto L1
L37:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v138)+8)))
	v145 = int32(0)
	goto L38
L38:
	;
	v154 = int32(0)
	if l1 == v154 {
		v163 = v154
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v171+v145<<(uint(int32(2))%32))))
	if v181 == int32(0) {
		goto L27
	} else {
		goto L49
	}
L40:
	;
	if l2 == int32(0) {
		goto L27
	} else {
		goto L43
	}
L41:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v157 <= v145 {
		v163 = v154
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v163 = v159 + v145<<(uint(int32(2))%32)
	goto L40
L43:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.B2i32(v163 == int32(0))|base.B2i32(v168 <= v145) != 0 {
		goto L27
	} else {
		goto L44
	}
L44:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v171 == int32(0) {
		goto L27
	} else {
		goto L45
	}
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v140 != v174 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v145 = v145 + int32(1)
	goto L38
L47:
	;
	goto L48
L48:
	;
	goto L39
L49:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v185 = F_equality_ops_are_compatible(m, v181, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L50
	}
L50:
	;
	if v185 == int32(0) {
		goto L27
	} else {
		goto L51
	}
L51:
	;
	v189 = int32(1)
	v191 = v124 + v189
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v191 < v192 {
		v124 = v191
		goto L35
	} else {
		goto L52
	}
L52:
	;
	goto L36
L53:
	;
	v194 = int32(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v195 == v194 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v205 = int32(1)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v206 != 0 {
		v344 = v205
		goto L1
	} else {
		goto L60
	}
L56:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+4))
	if v200 == int32(0) {
		v344 = v194
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	return int32(0)
L59:
	;
	goto L58
L60:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v207 != 0 {
		v344 = v205
		goto L1
	} else {
		goto L61
	}
L61:
	;
	goto L27
L62:
	;
	return int32(0)
L63:
	;
	goto L64
L64:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+8)))
	if v226 != 0 {
		v344 = int32(0)
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v220)+32))
	if v227 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v230 = v228
	goto L68
L67:
	;
	v230 = int32(0)
	goto L68
L68:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v231 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	return int32(1)
L70:
	;
	goto L71
L71:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v237 <= int32(0) {
		v344 = int32(1)
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v245 = v230
	v249 = int32(0)
	goto L73
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v231)+12))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v253+v249<<(uint(int32(2))%32))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+26)))
	if v258 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v344 = v336
	goto L1
L75:
	;
	v261 = int32(0)
	v263 = v245 + int32(4)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v220)+32))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	if base.Ui32(v263) < base.Ui32(v266+v267<<(uint(int32(2))%32)) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v328 = v245
	goto L77
L77:
	;
	v336 = int32(1)
	v338 = v249 + v336
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	if v338 < v339 {
		v245 = v328
		v249 = v338
		goto L73
	} else {
		goto L95
	}
L78:
	;
	v272 = v263
	goto L80
L79:
	;
	v272 = v261
	goto L80
L80:
	;
	v273 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257)+8)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v280 = v261
	goto L81
L81:
	;
	v287 = int32(0)
	if l1 == v287 {
		v297 = v287
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v306+v280<<(uint(int32(2))%32))))
	if v316 == int32(0) {
		v344 = v298
		goto L1
	} else {
		goto L92
	}
L83:
	;
	v298 = int32(0)
	if l2 == v298 {
		v344 = v298
		goto L1
	} else {
		goto L86
	}
L84:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v291 <= v280 {
		v297 = int32(0)
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v297 = v293 + v280<<(uint(int32(2))%32)
	goto L83
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.B2i32(v297 == int32(0))|base.B2i32(v303 <= v280) != 0 {
		v344 = v298
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v306 == int32(0) {
		v344 = v298
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	if v273 != v309 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v280 = v280 + int32(1)
	goto L81
L90:
	;
	goto L91
L91:
	;
	goto L82
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v274)+8))
	v320 = F_equality_ops_are_compatible(m, v316, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	if v320 == int32(0) {
		v344 = v298
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v328 = v272
	goto L77
L95:
	;
	goto L74
}
func F_query_or_expression_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(67) {
			v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v10 = F_query_tree_walker_impl(m, l0, l1, l2, l3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
