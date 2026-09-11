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
	var v47 int32
	_ = v47
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
	var v71 int32
	_ = v71
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
	v47 = v5
	goto L10
L8:
	;
	v71 = v5
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
		v62 = v47
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v71 = v62
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
		v62 = v47
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_CatalogTupleDelete(m, v15, v41+int32(4))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v62 = v47 + int32(1)
	goto L12
L16:
	;
	if v63 != 0 {
		v41 = v63
		v47 = v62
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_sequence_close(m, v15, int32(3))
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
	return v71
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
	F_CatalogTupleDelete(m, v16, v40+int32(4))
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
	F_sequence_close(m, v16, int32(3))
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
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int64
	_ = v235
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	v24 = F__emscripten_memset_bulkmem(m, v17+int32(16), base.I32_extend8_s(int32(0)), int32(136))
	mBase = m.M
	goto L1
L1:
	;
	v26 = F_palloc(m, int32(16))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	return
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+8)) = int64(137438953472)
	v31 = F_palloc(m, int32(384))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v31
	v36 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v36
	v39 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v26
	v48 = v17 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v48
	v56 = F_list_make1_impl(m, v36, v17+int32(4))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v56
	v61 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v61
	v66 = F_find_expr_references_walker(m, l1, v17+int32(152))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if int32(2) <= v69 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	F_pg_qsort(m, v72, v69, int32(12), int32(462))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	v150 = v68
	goto L10
L10:
	;
	if base.B2i32(l4 == int32(0))&base.B2i32(l3 == int32(110)) != 0 {
		v329 = v150
		goto L24
	} else {
		goto L25
	}
L11:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if int32(2) <= v77 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v87 = v80
	v90 = int32(1)
	v91 = v36
	goto L15
L13:
	;
	v137 = v36
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v137
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v150 = v143
	goto L10
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v100 = v97 + v90*int32(12)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	if v96 != v101 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v137 = v121
	goto L14
L17:
	;
	v125 = v90 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	if v125 < v126 {
		v87 = v120
		v90 = v125
		v91 = v121
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v113 = v87 + int32(12)
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v100)))
	*(*int64)(unsafe.Add(mBase, uint32(v113))) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v116
	v120 = v113
	v121 = v91 + int32(1)
	goto L17
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v103 != v104 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	if v106 == v107 {
		v120 = v87
		v121 = v91
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v106 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v107
	v120 = v87
	v121 = v91
	goto L17
L23:
	;
	goto L16
L24:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v329)+8))
	F_recordMultipleDependencies(m, l0, v337, v338, int32(110))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L2
	} else {
		goto L59
	}
L25:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	if v163 <= int32(0) {
		v329 = v150
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v167 = F_palloc(m, int32(16))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167)+8)) = int64(137438953472)
	v171 = int32(0)
	v173 = F_palloc(m, int32(384))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v173
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+8))
	if v175 < v179 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v185 = v182
	v192 = int32(0)
	v193 = v178
	v194 = v171
	goto L32
L30:
	;
	v262 = v178
	v263 = v171
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v262)+8)) = v263
	if l4 != 0 {
		goto L44
	} else {
		goto L45
	}
L32:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v201 = v198 + v192*int32(12)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v202 != int32(1259) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v262 = v244
	v263 = v245
	goto L31
L34:
	;
	v250 = v192 + int32(1)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	if v250 < v251 {
		v185 = v243
		v192 = v250
		v193 = v244
		v194 = v245
		goto L32
	} else {
		goto L42
	}
L35:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	*(*int64)(unsafe.Add(mBase, uint32(v185))) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+8)) = v237
	v243 = v185 + int32(12)
	v244 = v193
	v245 = v194 + int32(1)
	goto L34
L36:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
	if v205 != l2 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	if v209 <= v208 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = v209 << (uint(int32(1)) % 32)
	v216 = F_repalloc(m, v207, v209*int32(24))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L41
	}
L39:
	;
	v221 = v193
	v222 = v207
	v223 = v208
	goto L40
L40:
	;
	v226 = v223*int32(12) + v222
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v201)))
	*(*int64)(unsafe.Add(mBase, uint32(v226))) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v226)+8)) = v229
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v231 + int32(1)
	v243 = v185
	v244 = v221
	v245 = v194
	goto L34
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v216
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	v221 = v219
	v222 = v216
	v223 = v220
	goto L40
L42:
	;
	goto L33
L43:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	F_pfree(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L2
	} else {
		goto L53
	}
L44:
	;
	v268 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	if v269 <= v268 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	F_recordMultipleDependencies(m, l0, v296, v297, l3)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L2
	} else {
		goto L52
	}
L47:
	;
	v277 = v268
	goto L48
L48:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
	F_recordDependencyOn(m, v286+v277*int32(12), l0, l3)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L50
	}
L49:
	;
	goto L43
L50:
	;
	v293 = v277 + int32(1)
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	if v293 < v294 {
		v277 = v293
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L43
L53:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
	if v317 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	F_pfree(m, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L2
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	F_pfree(m, v167)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L2
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v329 = v322
	goto L24
L59:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	F_pfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v342)+4))
	if v346 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_pfree(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L2
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	F_pfree(m, v342)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L2
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	m.G0 = v17 + int32(160)
	return
}
