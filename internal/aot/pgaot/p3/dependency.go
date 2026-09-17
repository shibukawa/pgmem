package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_deleteDependencyRecordsForClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v15 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v11, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v11+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_systable_beginscan(m, v15, int32(2673), int32(1), int32(0), int32(2), v11)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v37 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = v37
	v48 = v5
	goto L10
L8:
	;
	v72 = v5
	goto L9
L9:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v51 = v49 + v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v52 != l2 {
		v62 = v48
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v72 = v62
	goto L9
L12:
	;
	v63 = F_systable_getnext(m, v35)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+24)))
	if v54 != l3&int32(255) {
		v62 = v48
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_simple_heap_delete(m, v15, v41+int32(4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v62 = v48 + int32(1)
	goto L12
L16:
	;
	if v63 != 0 {
		v41 = v63
		v48 = v62
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_relation_close(m, v15, int32(3))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v11 + int32(96)
	return v72
}
func F_deleteDependencyRecordsForSpecific(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v16 = F_table_open(m, int32(2608), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_ScanKeyInit(m, v12, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v12+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = F_systable_beginscan(m, v16, int32(2673), int32(1), int32(0), int32(2), v12)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v36 = F_systable_getnext(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v36 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = v36
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v51 = v49 + v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	if v52 != l3 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v65 = F_systable_getnext(m, v34)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	if v54 != l4 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+24)))
	if v56 != l2&int32(255) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_simple_heap_delete(m, v16, v40+int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	if v65 != 0 {
		v40 = v65
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	F_relation_close(m, v16, int32(3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v12 + int32(96)
	return
}
func F_recordDependencyOnSingleRelExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int64
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int64
	_ = v232
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
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	v19 = int32(16)
	v20 = v17 + v19
	base.MemoryFill(m, v20, int32(0), int32(136))
	v25 = F_palloc(m, v19)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(137438953472)
	v30 = F_palloc(m, int32(384))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v30
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v35
	v38 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)) = uint8(v38)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
	v51 = F_list_make1_impl(m, v35, v17+int32(4))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v51
	v56 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v56
	v61 = F_find_expr_references_walker(m, l1, v17+int32(152))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if int32(2) <= v64 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	F_pg_qsort(m, v67, v64, int32(12), int32(462))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v144 = v63
	goto L9
L9:
	;
	if base.B2i32(l4 == int32(0))&base.B2i32(l3 == int32(110)) != 0 {
		v323 = v144
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if int32(2) <= v72 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v78 = v75
	v85 = v35
	v86 = int32(1)
	goto L14
L12:
	;
	v131 = v35
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = v131
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v144 = v138
	goto L9
L14:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v95 = v92 + v86*int32(12)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v91 != v96 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v131 = v116
	goto L13
L16:
	;
	v120 = v86 + int32(1)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v120 < v121 {
		v78 = v115
		v85 = v116
		v86 = v120
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = v109
	v115 = v78 + int32(12)
	v116 = v85 + int32(1)
	goto L16
L18:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v98 != v99 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if v101 == v102 {
		v115 = v78
		v116 = v85
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v101 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v102
	v115 = v78
	v116 = v85
	goto L16
L22:
	;
	goto L15
L23:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v323)))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v323)+8))
	F_recordMultipleDependencies(m, l0, v332, v333, int32(110))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L58
	}
L24:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	if v158 <= int32(0) {
		v323 = v144
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v162 = F_palloc(m, int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v162)+8)) = int64(137438953472)
	v166 = int32(0)
	v168 = F_palloc(m, int32(384))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v168
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v170 < v174 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v180 = v177
	v187 = v173
	v188 = int32(0)
	v190 = v166
	goto L31
L29:
	;
	v256 = v173
	v259 = v166
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v256)+8)) = v259
	if l4 != 0 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	v196 = v193 + v188*int32(12)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	if v197 != int32(1259) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v256 = v239
	v259 = v241
	goto L30
L33:
	;
	v245 = v188 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	if v245 < v246 {
		v180 = v238
		v187 = v239
		v188 = v245
		v190 = v241
		goto L31
	} else {
		goto L41
	}
L34:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+8)) = v230
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v180))) = v232
	v238 = v180 + int32(12)
	v239 = v187
	v241 = v190 + int32(1)
	goto L33
L35:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v200 != l2 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	if v204 <= v203 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v204 << (uint(int32(1)) % 32)
	v211 = F_repalloc(m, v202, v204*int32(24))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v216 = v187
	v217 = v202
	v218 = v203
	goto L39
L39:
	;
	v221 = v217 + v218*int32(12)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v221)+8)) = v222
	v224 = *(*int64)(unsafe.Add(mBase, uint32(v196)))
	*(*int64)(unsafe.Add(mBase, uint32(v221))) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v226 + int32(1)
	v238 = v180
	v239 = v216
	v241 = v190
	goto L33
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v211
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v216 = v215
	v217 = v211
	v218 = v214
	goto L39
L41:
	;
	goto L32
L42:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	F_pfree(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L52
	}
L43:
	;
	v263 = int32(0)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v264 <= v263 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	F_recordMultipleDependencies(m, l0, v291, v292, l3)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L51
	}
L46:
	;
	v268 = v263
	goto L47
L47:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	F_recordDependencyOn(m, v281+v268*int32(12), l0, l3)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v288 = v268 + int32(1)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v288 < v289 {
		v268 = v288
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	goto L42
L52:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v312 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v312)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_pfree(m, v162)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v323 = v317
	goto L23
L58:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	F_pfree(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v341 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_pfree(m, v341)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v337)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	m.G0 = v17 + int32(160)
	return
}
