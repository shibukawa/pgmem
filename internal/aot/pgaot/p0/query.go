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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
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
	if v102 != 0 {
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
	v107 = F_query_tree_walker_impl(m, l0, int32(1604), v11+int32(15), int32(3))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
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
	v6 = m.G0
	v8 = v6 - int32(480)
	m.G0 = v8
	v15 = F__emscripten_memset_bulkmem(m, v8+int32(392), base.I32_extend8_s(int32(0)), int32(88))
	mBase = m.M
	v16 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+448)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+388)) = int32(266)
	v23 = F__emscripten_memset_bulkmem(m, v8, base.I32_extend8_s(v16), int32(384))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(267)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v23 + int32(388)
	v29 = F_extract_query_dependencies_walker(m, l0, v23)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		return
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v23)+444))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+448))
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+469)))
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v35)
		m.G0 = v23 + int32(480)
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
	var v26 int32
	_ = v26
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
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
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
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v13 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v341
L2:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v107 != 0 {
		goto L23
	} else {
		goto L24
	}
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v17 <= int32(0) {
		v341 = int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = v4
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v26<<(uint(int32(2))%32))))
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
	v341 = v90
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v82 == int32(0) {
		goto L2
	} else {
		goto L19
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
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v68 <= v47 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v65 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = v72 + v47<<(uint(int32(2))%32)
	if v75 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v80 != v42 {
		v47 = v47 + int32(1)
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v86 = F_equality_ops_are_compatible(m, v82, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	if v86 == int32(0) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v90 = int32(1)
	v92 = v26 + v90
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v92 < v93 {
		v26 = v92
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	return int32(0)
L24:
	;
	goto L25
L25:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v111 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v218 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L27:
	;
	if v110 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v110 != 0 {
		goto L51
	} else {
		goto L52
	}
L30:
	;
	return int32(0)
L31:
	;
	goto L32
L32:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v115 <= int32(0) {
		v341 = int32(1)
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v125 = int32(0)
	goto L34
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v111)+12))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v125<<(uint(int32(2))%32))))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v137 = F_get_sortgroupclause_tle(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L7
	} else {
		goto L36
	}
L35:
	;
	v341 = v187
	goto L1
L36:
	;
	v139 = int32(*(*int16)(unsafe.Add(mBase, uint32(v137)+8)))
	v144 = int32(0)
	goto L37
L37:
	;
	v153 = int32(0)
	if l1 == v153 {
		v162 = v153
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	if v179 == int32(0) {
		goto L26
	} else {
		goto L47
	}
L39:
	;
	if l2 == int32(0) {
		goto L26
	} else {
		goto L42
	}
L40:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v156 <= v144 {
		v162 = v153
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v162 = v158 + v144<<(uint(int32(2))%32)
	goto L39
L42:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v165 <= v144 {
		goto L26
	} else {
		goto L43
	}
L43:
	;
	if v162 == int32(0) {
		goto L26
	} else {
		goto L44
	}
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v172 = v169 + v144<<(uint(int32(2))%32)
	if v172 == int32(0) {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v177 != v139 {
		v144 = v144 + int32(1)
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	v183 = F_equality_ops_are_compatible(m, v179, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L48
	}
L48:
	;
	if v183 == int32(0) {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v187 = int32(1)
	v189 = v125 + v187
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	if v189 < v190 {
		v125 = v189
		goto L34
	} else {
		goto L50
	}
L50:
	;
	goto L35
L51:
	;
	v192 = int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v193 == v192 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v203 = int32(1)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v204 != 0 {
		v341 = v203
		goto L1
	} else {
		goto L58
	}
L54:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v198 == int32(0) {
		v341 = v192
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	return int32(0)
L57:
	;
	goto L56
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v205 != 0 {
		v341 = v203
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L26
L60:
	;
	return int32(0)
L61:
	;
	goto L62
L62:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+8)))
	if v224 != 0 {
		v341 = int32(0)
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	if v225 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v228 = v226
	goto L66
L65:
	;
	v228 = int32(0)
	goto L66
L66:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v229 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	return int32(1)
L68:
	;
	goto L69
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v235 <= int32(0) {
		v341 = int32(1)
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v246 = v228
	v247 = int32(0)
	goto L71
L71:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v229)+12))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v251+v247<<(uint(int32(2))%32))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+26)))
	if v256 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v341 = v333
	goto L1
L73:
	;
	v259 = int32(0)
	v261 = v246 + int32(4)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v218)+32))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+12))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	if base.Ui32(v261) < base.Ui32(v264+v265<<(uint(int32(2))%32)) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v328 = v246
	goto L75
L75:
	;
	v333 = int32(1)
	v335 = v247 + v333
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v335 < v336 {
		v246 = v328
		v247 = v335
		goto L71
	} else {
		goto L92
	}
L76:
	;
	v270 = v261
	goto L78
L77:
	;
	v270 = v259
	goto L78
L78:
	;
	v271 = int32(*(*int16)(unsafe.Add(mBase, uint32(v255)+8)))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v277 = v259
	goto L79
L79:
	;
	v285 = int32(0)
	if l1 == v285 {
		v295 = v285
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v306)))
	if v313 == int32(0) {
		v341 = v296
		goto L1
	} else {
		goto L89
	}
L81:
	;
	v296 = int32(0)
	if l2 == v296 {
		v341 = v296
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v289 <= v277 {
		v295 = int32(0)
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v295 = v291 + v277<<(uint(int32(2))%32)
	goto L81
L84:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v299 <= v277 {
		v341 = v296
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v295 == int32(0) {
		v341 = v296
		goto L1
	} else {
		goto L86
	}
L86:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v306 = v303 + v277<<(uint(int32(2))%32)
	if v306 == int32(0) {
		v341 = v296
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v311 != v271 {
		v277 = v277 + int32(1)
		goto L79
	} else {
		goto L88
	}
L88:
	;
	goto L80
L89:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
	v317 = F_equality_ops_are_compatible(m, v313, v316)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L90
	}
L90:
	;
	if v317 == int32(0) {
		v341 = v296
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v328 = v270
	goto L75
L92:
	;
	goto L72
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
