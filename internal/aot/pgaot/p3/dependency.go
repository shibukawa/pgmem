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
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v5 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(112)
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
	F_ScanKeyInit(m, v11, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v11+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v37 = F_systable_beginscan(m, v15, int32(2673), int32(1), int32(0), int32(2), v11)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = F_systable_getnext(m, v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v39
	v50 = v5
	goto L10
L8:
	;
	v74 = v5
	goto L9
L9:
	;
	F_systable_endscan(m, v37)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
	v53 = v51 + v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v54 != l2 {
		v64 = v50
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v74 = v64
	goto L9
L12:
	;
	v65 = F_systable_getnext(m, v37)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+24)))
	if v56 != l3&int32(255) {
		v64 = v50
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_simple_heap_delete(m, v15, v43+int32(4))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v64 = v50 + int32(1)
	goto L12
L16:
	;
	if v65 != 0 {
		v43 = v65
		v50 = v64
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
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v11 + int32(112)
	return v74
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
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	v10 = m.G0
	v12 = v10 - int32(112)
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
	F_ScanKeyInit(m, v12, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v12+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v36 = F_systable_beginscan(m, v16, int32(2673), int32(1), int32(0), int32(2), v12)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v38 = F_systable_getnext(m, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v42 = v38
	goto L10
L8:
	;
	goto L9
L9:
	;
	F_systable_endscan(m, v36)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
	v53 = v51 + v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	if v54 != l3 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v67 = F_systable_getnext(m, v36)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	if v56 != l4 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+24)))
	if v58 != l2&int32(255) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	F_simple_heap_delete(m, v16, v42+int32(4))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	if v67 != 0 {
		v42 = v67
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
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v12 + int32(112)
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int64
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
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
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
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
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int64
	_ = v226
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int64
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
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
	v31 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v31
	v36 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v36
	v39 = int32(114)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+37)) = uint8(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+152)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
	v52 = F_list_make1_impl(m, v36, v17+int32(4))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v52
	v57 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+156)) = v57
	v62 = F_find_expr_references_walker(m, l1, v17+int32(152))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if int32(2) <= v65 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	F_pg_qsort(m, v68, v65, int32(12), int32(497))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	v145 = v64
	goto L9
L9:
	;
	if base.B2i32(l4 == int32(0))&base.B2i32(l3 == int32(110)) != 0 {
		v325 = v145
		goto L23
	} else {
		goto L24
	}
L10:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if int32(2) <= v73 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v79 = v76
	v86 = v36
	v87 = int32(1)
	goto L14
L12:
	;
	v132 = v36
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v132
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v145 = v139
	goto L9
L14:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v96 = v93 + v87*int32(12)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v92 != v97 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v132 = v117
	goto L13
L16:
	;
	v121 = v87 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v121 < v122 {
		v79 = v116
		v86 = v117
		v87 = v121
		goto L14
	} else {
		goto L22
	}
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+12)) = v110
	v116 = v79 + int32(12)
	v117 = v86 + int32(1)
	goto L16
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 != v100 {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	if v102 == v103 {
		v116 = v79
		v117 = v86
		goto L16
	} else {
		goto L20
	}
L20:
	;
	if v102 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = v103
	v116 = v79
	v117 = v86
	goto L16
L22:
	;
	goto L15
L23:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	F_recordMultipleDependencies(m, l0, v334, v335, int32(110))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L58
	}
L24:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	if v159 <= int32(0) {
		v325 = v145
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v163 = F_palloc(m, int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v163)+8)) = int64(137438953472)
	v167 = int32(0)
	v170 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v172 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v163)+4)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v170
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+8))
	if v172 < v176 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v182 = v179
	v189 = v175
	v190 = int32(0)
	v192 = v167
	goto L31
L29:
	;
	v258 = v175
	v261 = v167
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v258)+8)) = v261
	if l4 != 0 {
		goto L43
	} else {
		goto L44
	}
L31:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v198 = v195 + v190*int32(12)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	if v199 != int32(1259) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	v258 = v241
	v261 = v243
	goto L30
L33:
	;
	v247 = v190 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	if v247 < v248 {
		v182 = v240
		v189 = v241
		v190 = v247
		v192 = v243
		goto L31
	} else {
		goto L41
	}
L34:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+8)) = v232
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	*(*int64)(unsafe.Add(mBase, uint32(v182))) = v234
	v240 = v182 + int32(12)
	v241 = v189
	v243 = v192 + int32(1)
	goto L33
L35:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v202 != l2 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	if v206 <= v205 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163)+12)) = v206 << (uint(int32(1)) % 32)
	v213 = F_repalloc(m, v204, v206*int32(24))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v218 = v189
	v219 = v204
	v220 = v205
	goto L39
L39:
	;
	v223 = v219 + v220*int32(12)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v198)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+8)) = v224
	v226 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	*(*int64)(unsafe.Add(mBase, uint32(v223))) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+8)) = v228 + int32(1)
	v240 = v182
	v241 = v218
	v243 = v192
	goto L33
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v213
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v218 = v217
	v219 = v213
	v220 = v216
	goto L39
L41:
	;
	goto L32
L42:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	F_pfree(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L52
	}
L43:
	;
	v265 = int32(0)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v266 <= v265 {
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	F_recordMultipleDependencies(m, l0, v293, v294, l3)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L51
	}
L46:
	;
	v270 = v265
	goto L47
L47:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	F_recordDependencyOn(m, v283+v270*int32(12), l0, l3)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v290 = v270 + int32(1)
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	if v290 < v291 {
		v270 = v290
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
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v163)+4))
	if v314 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v314)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_pfree(m, v163)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v325 = v319
	goto L23
L58:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v17)+152))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)))
	F_pfree(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v343 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_pfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v339)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
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
