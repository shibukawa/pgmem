package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CreateExprContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v5 = int32(_a_F_CreateExprContext_0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CreateExprContext[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_CreateExprContext[0])) = v8
	v11 = F_palloc0(m, int32(72))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(382)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v19
		v25 = F_AllocSetContextCreateInternal(m, v19, int32(_a_F_CreateExprContext_1), int32(0), int32(_a_F_CreateExprContext_2), int32(_a_F_CreateExprContext_3))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
			*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l0
			v34 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v34)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v30
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
			v46 = F_lcons(m, v11, v45)
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+140)) = v46
				*(*int32)(unsafe.Add(mBase, _c_F_CreateExprContext[0])) = v6
				return v11
			}
		}
	}
}
func F_FreeExprContext(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v5 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = int32(_a_F_FreeExprContext_0)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_FreeExprContext[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_FreeExprContext[0])) = v9
	v13 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_MemoryContextDelete(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L9
	} else {
		goto L13
	}
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v15
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_FreeExprContext[0])) = v7
	goto L3
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	m.T0[v18].(func(*base.Module, int32))(m, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_pfree(m, v13)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L9
	} else {
		goto L11
	}
L9:
	;
	return
L10:
	;
	goto L8
L11:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v23 != 0 {
		v13 = v23
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v33 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+140))
	v35 = F_list_delete_ptr(m, v34, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_pfree(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+140)) = v35
	goto L16
L18:
	;
	return
}
func F_exec_eval_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
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
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int64
	_ = v262
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v18 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_exec_prepare_plan(m, l0, l1, int32(2048))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L4
	} else {
		goto L107
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L101
	}
L8:
	;
	m.G0 = v16 - int32(-64)
	return v300
L9:
	;
	v249 = F_exec_run_select(m, l0, l1, int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L86
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+56))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)))
	if v33 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v36 == v32 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_EnsurePortalSnapshotExists(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v32 != v42 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[2]))
	F_ReleaseCachedPlan(m, v100, v240)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L85
	}
L17:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v178)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v178
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v184 != v32 {
		goto L71
	} else {
		goto L72
	}
L18:
	;
	if v79 != 0 {
		goto L32
	} else {
		goto L33
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v46 = v44
	goto L21
L20:
	;
	v46 = int32(0)
	goto L21
L21:
	;
	if v41 == int32(0) {
		v77 = v6
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = v77
	goto L18
L23:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+95)))
	if v49&int32(1) == int32(0) {
		v77 = v6
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v41 != v54 {
		v77 = v6
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+10)))
	if v56 != int32(1) {
		v77 = v6
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	v60 = F_SearchPathMatchesCurrentEnvironment(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if v60 == int32(0) {
		v77 = v6
		goto L22
	} else {
		goto L28
	}
L28:
	;
	if v46 == int32(0) {
		v79 = int32(1)
		goto L18
	} else {
		goto L29
	}
L29:
	;
	F_ResourceOwnerEnlarge(m, v46)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v69 = int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v41)+28)) = v70 + v69
	F_ResourceOwnerRemember(m, v46, v41, int32(_a_F_exec_eval_expr_12))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v77 = v69
	goto L22
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v32
	goto L17
L33:
	;
	goto L34
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v32 == v81 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	F_ReleaseCachedPlan(m, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v87 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+60)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(0)
	v93 = int32(_a_F_exec_eval_expr_11)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v100 = F_SPI_plan_get_cached_plan(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v94
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	if v107 == v104 {
		v155 = v104
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v155 == int32(0) {
		goto L16
	} else {
		goto L66
	}
L41:
	;
	goto L40
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v110 != int32(1) {
		v155 = v104
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+60))
	if v115 == int32(0) {
		v155 = v104
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v118 != int32(1) {
		v155 = v104
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	if v123 != int32(67) {
		v155 = v104
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v126 != int32(1) {
		v155 = v104
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+52))
	if v129 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+36)))
	if v130 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+37)))
	if v131 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+38)))
	if v132 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L51
	}
L51:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+39)))
	if v133 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v122)+48))
	if v134 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v122)+60))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+4))
	if v136 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L54
	}
L54:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135)+8))
	if v137 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v122)+100))
	if v138 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L56
	}
L56:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v122)+108))
	if v139 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v122)+112))
	if v140 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122)+116))
	if v141 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v122)+120))
	if v142 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L60
	}
L60:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v122)+124))
	if v143 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L61
	}
L61:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v122)+128))
	if v144 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v122)+132))
	if v145 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L63
	}
L63:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v122)+144))
	if v146 != 0 {
		v155 = v104
		goto L41
	} else {
		goto L64
	}
L64:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v122)+76))
	if v147 == int32(0) {
		v155 = v104
		goto L41
	} else {
		goto L65
	}
L65:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v155 = base.B2i32(v150 == int32(1))
	goto L41
L66:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v160 = F_CachedPlanAllowsSimpleValidityCheck(m, v158, v100, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	if v160 == int32(0) {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v100
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[2]))
	F_ReleaseCachedPlan(m, v100, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_exec_save_simple_expr(m, l1, v100)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	goto L17
L71:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v191 = F_ExecInitExprWithParams(m, v190, v178)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	if v202 != int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v32
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v194)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v191
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v183
	v300 = v236
	goto L8
L76:
	;
	v225 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v225)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+20))
	v229 = m.T0[v228].(func(*base.Module, int32, int32, int32) int32)(m, v227, v26, l2)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L84
	}
L77:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v205 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v208 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_PushActiveSnapshot(m, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v212 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v212)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	v216 = m.T0[v215].(func(*base.Module, int32, int32, int32) int32)(m, v214, v26, l2)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v218)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v178)+20)) = v179
	F_PopActiveSnapshot(m)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v236 = v216
	goto L75
L84:
	;
	v231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v231)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v231
	*(*int32)(unsafe.Add(mBase, uint32(v178)+20)) = v179
	v236 = v229
	goto L75
L85:
	;
	goto L9
L86:
	;
	if v249 != int32(5) {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	if v255 != int32(1) {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v254)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v254)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v260
	v262 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	if base.Ui64(v262) <= base.Ui64(int64(1)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v253)+4))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	v298 = F_SPI_getbinval(m, v296, v254, int32(1), l2)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L100
	}
L90:
	;
	if base.I32_wrap_i64(v262) == int32(1) {
		goto L89
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L4
	} else {
		goto L94
	}
L93:
	;
	v268 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v268)
	v300 = int32(0)
	goto L8
L94:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_exec_eval_expr_9), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v285
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v16)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_10), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	v300 = v298
	goto L8
L101:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_exec_eval_expr_1), int32(0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v327
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v14+int32(-16))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_4), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v348
	F_errmsg_plural(m, int32(_a_F_exec_eval_expr_6), int32(_a_F_exec_eval_expr_7), v348, v14+int32(-32))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v359
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v14+int32(-48))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_8), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_exprCollation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v95 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v6 + int32(16)
	return v95
L2:
	;
	v10 = l0
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v13 - int32(6) {
	case 0, 2, 8, 9, 11, 12, 13, 19, 23:
		goto L7
	case 1, 3, 5:
		goto L28
	case 4, 14, 15, 20, 24, 30, 31, 40, 46, 47, 52, 53:
		v95 = v2
		goto L1
	default:
		goto L8
	case 7:
		goto L27
	case 10, 54, 313:
		goto L26
	case 16:
		goto L25
	case 17:
		goto L24
	case 18:
		goto L23
	case 21:
		goto L22
	case 22:
		goto L21
	case 25, 26:
		goto L20
	case 28:
		goto L19
	case 29, 32, 33:
		goto L18
	case 34:
		goto L17
	case 35:
		goto L16
	case 38:
		goto L15
	case 39:
		goto L14
	case 41:
		goto L12
	case 42:
		goto L13
	case 49:
		goto L11
	case 50, 51:
		goto L10
	case 55:
		goto L9
	}
L4:
	;
	v95 = v2
	goto L1
L5:
	;
	if v93 != 0 {
		v10 = v93
		goto L3
	} else {
		goto L52
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L45
	} else {
		goto L49
	}
L7:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v95 = v79
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L45
	} else {
		goto L46
	}
L9:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v62 != 0 {
		v10 = v62
		goto L3
	} else {
		goto L44
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v95 = v61
	goto L1
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v95 = v60
	goto L1
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v59 != 0 {
		v93 = v59
		goto L5
	} else {
		goto L43
	}
L13:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v95 = v58
	goto L1
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v57 != 0 {
		v93 = v57
		goto L5
	} else {
		goto L42
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v56 != 0 {
		v10 = v56
		goto L3
	} else {
		goto L41
	}
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v52 == int32(6) {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v46 == int32(19) {
		goto L35
	} else {
		goto L36
	}
L18:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v95 = v43
	goto L1
L19:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v95 = v42
	goto L1
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v95 = v41
	goto L1
L21:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v95 = v40
	goto L1
L22:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v95 = v39
	goto L1
L23:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != 0 {
		v10 = v38
		goto L3
	} else {
		goto L34
	}
L24:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v32 - int32(4) {
	case 0, 2:
		goto L33
	default:
		v95 = v2
		goto L1
	}
L25:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v19 - int32(4) {
	case 0, 2:
		goto L29
	default:
		v95 = v2
		goto L1
	}
L26:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v93 = v18
	goto L5
L27:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v95 = v17
	goto L1
L28:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v95 = v16
	goto L1
L29:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v25 != int32(67) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v31 != 0 {
		v10 = v31
		goto L3
	} else {
		goto L32
	}
L32:
	;
	v95 = v2
	goto L1
L33:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	v95 = v35
	goto L1
L34:
	;
	v95 = v2
	goto L1
L35:
	;
	v49 = int32(950)
	goto L37
L36:
	;
	v49 = int32(0)
	goto L37
L37:
	;
	v95 = v49
	goto L1
L38:
	;
	v55 = int32(100)
	goto L40
L39:
	;
	v55 = int32(0)
	goto L40
L40:
	;
	v95 = v55
	goto L1
L41:
	;
	v95 = v2
	goto L1
L42:
	;
	v95 = v2
	goto L1
L43:
	;
	v95 = v2
	goto L1
L44:
	;
	v95 = v2
	goto L1
L45:
	;
	return int32(0)
L46:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v69
	F_errmsg_internal(m, int32(_a_F_exprCollation_0), v6)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_exprCollation_1), int32(1062), int32(_a_F_exprCollation_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L45
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errmsg_internal(m, int32(_a_F_exprCollation_3), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_exprCollation_1), int32(889), int32(_a_F_exprCollation_2))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L45
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	goto L4
}
func F_exprLocation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
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
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	v9 = l0
	goto L12
L4:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v241
L5:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v230 = F_exprLocation(m, v229)
	mBase = m.M
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v231) < base.Ui32(v230) {
		goto L163
	} else {
		goto L164
	}
L6:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v218 = F_exprLocation(m, v217)
	mBase = m.M
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v219) < base.Ui32(v218) {
		goto L154
	} else {
		goto L155
	}
L7:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v206 = F_exprLocation(m, v205)
	mBase = m.M
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v207) < base.Ui32(v206) {
		goto L145
	} else {
		goto L146
	}
L8:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v203
L9:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v201
L10:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v199
L11:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v197
L12:
	;
	v13 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v15 - int32(1) {
	case 0:
		goto L39
	default:
		v194 = v13
		goto L14
	case 2, 7, 31, 38, 113:
		goto L11
	case 3:
		goto L38
	case 4, 24, 25, 30, 43, 59, 61, 73, 81, 82, 125, 133, 134, 318:
		v189 = int32(4)
		goto L15
	case 5:
		goto L37
	case 6, 44, 67, 97, 114:
		goto L10
	case 8:
		goto L36
	case 9, 35, 45, 88, 94, 95, 111, 131, 132, 208:
		goto L9
	case 10:
		goto L35
	case 12, 32, 41, 96, 98, 106, 109:
		goto L8
	case 13:
		goto L34
	case 14, 16, 17, 18, 19:
		goto L7
	case 15, 29, 51:
		goto L6
	case 20:
		goto L33
	case 21:
		goto L32
	case 26, 54:
		goto L5
	case 27:
		goto L31
	case 28:
		goto L30
	case 34:
		goto L29
	case 36:
		goto L28
	case 37, 39, 55, 56, 71, 79, 80, 110, 112, 129, 130:
		goto L4
	case 40:
		goto L27
	case 46:
		goto L25
	case 47:
		goto L26
	case 52:
		goto L24
	case 60:
		goto L23
	case 68, 69:
		goto L21
	case 70:
		goto L22
	case 72:
		goto L19
	case 75:
		goto L20
	case 83:
		goto L18
	case 89:
		goto L17
	case 160:
		goto L16
	}
L13:
	;
	return v194
L14:
	;
	goto L13
L15:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9+v189)))
	if v191 != 0 {
		v9 = v191
		goto L12
	} else {
		goto L144
	}
L16:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
	return v187
L17:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	return v185
L18:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v183
L19:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+28))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v165 = F_exprLocation(m, v164)
	mBase = m.M
	if base.Ui32(v165) < base.Ui32(v163) {
		goto L126
	} else {
		goto L127
	}
L20:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v150 = F_exprLocation(m, v149)
	mBase = m.M
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v151) < base.Ui32(v150) {
		goto L117
	} else {
		goto L118
	}
L21:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	return v147
L22:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v136 = F_exprLocation(m, v135)
	mBase = m.M
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v137) < base.Ui32(v136) {
		goto L108
	} else {
		goto L109
	}
L23:
	;
	v189 = int32(12)
	goto L15
L24:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v123 = F_exprLocation(m, v122)
	mBase = m.M
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v124) < base.Ui32(v123) {
		goto L99
	} else {
		goto L100
	}
L25:
	;
	v189 = int32(8)
	goto L15
L26:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v110 = F_exprLocation(m, v109)
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if base.Ui32(v111) < base.Ui32(v110) {
		goto L90
	} else {
		goto L91
	}
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v98 = F_exprLocation(m, v97)
	mBase = m.M
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if base.Ui32(v99) < base.Ui32(v98) {
		goto L81
	} else {
		goto L82
	}
L28:
	;
	v189 = int32(20)
	goto L15
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v94
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v83 = F_exprLocation(m, v82)
	mBase = m.M
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v84) < base.Ui32(v83) {
		goto L72
	} else {
		goto L73
	}
L31:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v71 = F_exprLocation(m, v70)
	mBase = m.M
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if base.Ui32(v72) < base.Ui32(v71) {
		goto L63
	} else {
		goto L64
	}
L32:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v60 = F_exprLocation(m, v59)
	mBase = m.M
	if base.Ui32(v60) < base.Ui32(v58) {
		goto L54
	} else {
		goto L55
	}
L33:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v47 = F_exprLocation(m, v46)
	mBase = m.M
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v48) < base.Ui32(v47) {
		goto L45
	} else {
		goto L46
	}
L34:
	;
	v189 = int32(32)
	goto L15
L35:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	return v43
L36:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v41
L37:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	return v39
L38:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v37
L39:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v18 <= int32(0) {
		v194 = v13
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v23 = int32(0)
	goto L41
L41:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23<<(uint(int32(2))%32))))
	v31 = F_exprLocation(m, v30)
	mBase = m.M
	if int32(0) <= v31 {
		v194 = v31
		goto L14
	} else {
		goto L43
	}
L42:
	;
	v194 = v31
	goto L14
L43:
	;
	v35 = v23 + int32(1)
	if v35 != v18 {
		v23 = v35
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v50 = v48
	goto L47
L46:
	;
	v50 = v47
	goto L47
L47:
	;
	if v47 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v53 = v48
	goto L50
L49:
	;
	v53 = v50
	goto L50
L50:
	;
	if v48 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v56 = v47
	goto L53
L52:
	;
	v56 = v53
	goto L53
L53:
	;
	return v56
L54:
	;
	v62 = v60
	goto L56
L55:
	;
	v62 = v58
	goto L56
L56:
	;
	if v58 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v65 = v60
	goto L59
L58:
	;
	v65 = v62
	goto L59
L59:
	;
	if v60 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v68 = v58
	goto L62
L61:
	;
	v68 = v65
	goto L62
L62:
	;
	return v68
L63:
	;
	v74 = v72
	goto L65
L64:
	;
	v74 = v71
	goto L65
L65:
	;
	if v71 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v77 = v72
	goto L68
L67:
	;
	v77 = v74
	goto L68
L68:
	;
	if v72 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v80 = v71
	goto L71
L70:
	;
	v80 = v77
	goto L71
L71:
	;
	return v80
L72:
	;
	v86 = v84
	goto L74
L73:
	;
	v86 = v83
	goto L74
L74:
	;
	if v83 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v89 = v84
	goto L77
L76:
	;
	v89 = v86
	goto L77
L77:
	;
	if v84 < int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v92 = v83
	goto L80
L79:
	;
	v92 = v89
	goto L80
L80:
	;
	return v92
L81:
	;
	v101 = v99
	goto L83
L82:
	;
	v101 = v98
	goto L83
L83:
	;
	if v98 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v104 = v99
	goto L86
L85:
	;
	v104 = v101
	goto L86
L86:
	;
	if v99 < int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v107 = v98
	goto L89
L88:
	;
	v107 = v104
	goto L89
L89:
	;
	return v107
L90:
	;
	v113 = v111
	goto L92
L91:
	;
	v113 = v110
	goto L92
L92:
	;
	if v110 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v116 = v111
	goto L95
L94:
	;
	v116 = v113
	goto L95
L95:
	;
	if v111 < int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v119 = v110
	goto L98
L97:
	;
	v119 = v116
	goto L98
L98:
	;
	return v119
L99:
	;
	v126 = v124
	goto L101
L100:
	;
	v126 = v123
	goto L101
L101:
	;
	if v123 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v129 = v124
	goto L104
L103:
	;
	v129 = v126
	goto L104
L104:
	;
	if v124 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v132 = v123
	goto L107
L106:
	;
	v132 = v129
	goto L107
L107:
	;
	return v132
L108:
	;
	v139 = v137
	goto L110
L109:
	;
	v139 = v136
	goto L110
L110:
	;
	if v136 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v142 = v137
	goto L113
L112:
	;
	v142 = v139
	goto L113
L113:
	;
	if v137 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v145 = v136
	goto L116
L115:
	;
	v145 = v142
	goto L116
L116:
	;
	return v145
L117:
	;
	v153 = v151
	goto L119
L118:
	;
	v153 = v150
	goto L119
L119:
	;
	if v150 < int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v156 = v151
	goto L122
L121:
	;
	v156 = v153
	goto L122
L122:
	;
	if v151 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v159 = v150
	goto L125
L124:
	;
	v159 = v156
	goto L125
L125:
	;
	return v159
L126:
	;
	v167 = v165
	goto L128
L127:
	;
	v167 = v163
	goto L128
L128:
	;
	if v163 < int32(0) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v170 = v165
	goto L131
L130:
	;
	v170 = v167
	goto L131
L131:
	;
	if v165 < int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v173 = v163
	goto L134
L133:
	;
	v173 = v170
	goto L134
L134:
	;
	if base.Ui32(v173) < base.Ui32(v161) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v175 = v173
	goto L137
L136:
	;
	v175 = v161
	goto L137
L137:
	;
	if v161 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v178 = v173
	goto L140
L139:
	;
	v178 = v175
	goto L140
L140:
	;
	if v173 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v181 = v161
	goto L143
L142:
	;
	v181 = v178
	goto L143
L143:
	;
	return v181
L144:
	;
	v194 = v13
	goto L14
L145:
	;
	v209 = v207
	goto L147
L146:
	;
	v209 = v206
	goto L147
L147:
	;
	if v206 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v212 = v207
	goto L150
L149:
	;
	v212 = v209
	goto L150
L150:
	;
	if v207 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v215 = v206
	goto L153
L152:
	;
	v215 = v212
	goto L153
L153:
	;
	return v215
L154:
	;
	v221 = v219
	goto L156
L155:
	;
	v221 = v218
	goto L156
L156:
	;
	if v218 < int32(0) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v224 = v219
	goto L159
L158:
	;
	v224 = v221
	goto L159
L159:
	;
	if v219 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v227 = v218
	goto L162
L161:
	;
	v227 = v224
	goto L162
L162:
	;
	return v227
L163:
	;
	v233 = v231
	goto L165
L164:
	;
	v233 = v230
	goto L165
L165:
	;
	if v230 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v236 = v231
	goto L168
L167:
	;
	v236 = v233
	goto L168
L168:
	;
	if v231 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v239 = v230
	goto L171
L170:
	;
	v239 = v236
	goto L171
L171:
	;
	return v239
}
func F_exprSetCollation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = l0
	goto L5
L1:
	;
	m.G0 = v7 + int32(16)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9+v37))) = l1
	goto L1
L3:
	;
	v37 = int32(20)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v14 - int32(6) {
	case 0, 2, 8, 9, 11, 12, 13, 19, 23:
		goto L3
	case 1, 3, 5, 22, 50, 51:
		v37 = int32(12)
		goto L2
	case 4, 10, 14, 15, 16, 20, 24, 30, 31, 34, 35, 40, 41, 46, 47, 52, 53:
		goto L1
	default:
		goto L4
	case 7, 26, 29, 32, 33:
		goto L11
	case 21, 49:
		goto L10
	case 38:
		goto L9
	case 39:
		goto L8
	case 42:
		goto L7
	}
L6:
	;
	v37 = int32(56)
	goto L2
L7:
	;
	goto L6
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v20 != 0 {
		v9 = v20
		goto L5
	} else {
		goto L12
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v9 = v19
	goto L5
L10:
	;
	v37 = int32(16)
	goto L2
L11:
	;
	v37 = int32(8)
	goto L2
L12:
	;
	goto L1
L13:
	;
	return
L14:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v26
	F_errmsg_internal(m, int32(_a_F_exprSetCollation_0), v7)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_exprSetCollation_1), int32(1306), int32(_a_F_exprSetCollation_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_exprSetInputCollation(m *base.Module, l0 int32, l1 int32) {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = v5 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v7) {
	} else {
		v11 = int32(1) << (uint(v7) % 32)
		if v11&int32(3904) == int32(0) {
			if v11&int32(5) != 0 {
				v23 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v23+l0))) = l1
			} else {
				if v7 != int32(30) {
				} else {
					v23 = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(v23+l0))) = l1
				}
			}
		} else {
			v23 = int32(24)
			*(*int32)(unsafe.Add(mBase, uint32(v23+l0))) = l1
		}
	}
	return
}
func F_exprTypmod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	v2 = int32(0)
	v7 = int32(-1)
	if l0 == v2 {
		v289 = v7
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v292
L2:
	;
	return v289
L3:
	;
	v10 = l0
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v16 - int32(6) {
	case 0, 2, 8, 19, 23:
		goto L8
	case 1:
		goto L30
	case 3, 4, 5, 6, 7, 11, 12, 14, 15, 20, 22, 24, 27, 30, 31, 35, 36, 37, 40, 43, 44, 45, 46, 47, 48, 52, 53, 54:
		v289 = v7
		goto L2
	case 9:
		goto L29
	case 10:
		goto L28
	case 13:
		goto L27
	case 16:
		goto L26
	case 17:
		goto L25
	case 18:
		goto L24
	case 21:
		goto L23
	case 25:
		goto L22
	case 26:
		goto L21
	case 28:
		goto L20
	case 29:
		goto L19
	case 32:
		goto L18
	case 33:
		goto L17
	case 34:
		goto L16
	case 38:
		goto L15
	case 39:
		goto L14
	case 41:
		goto L12
	case 42:
		goto L13
	case 49:
		goto L11
	case 50, 51:
		goto L10
	case 55:
		goto L9
	default:
		goto L31
	}
L5:
	;
	v289 = v7
	goto L2
L6:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v283 != 0 {
		v10 = v283
		goto L4
	} else {
		goto L102
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L44
	} else {
		goto L99
	}
L8:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v289 = v268
	goto L2
L9:
	;
	v282 = v10 + int32(12)
	goto L6
L10:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v264
L11:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v262
L12:
	;
	v282 = v10 + int32(8)
	goto L6
L13:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)+12))
	return v258
L14:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v254)+12))
	return v255
L15:
	;
	v282 = v10 + int32(8)
	goto L6
L16:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v250
L17:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	v212 = F_exprType(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L44
	} else {
		goto L86
	}
L18:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v170 = F_exprType(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L44
	} else {
		goto L73
	}
L19:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v124 == int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L58
	}
L20:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v122
L21:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v78 == int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L43
	}
L22:
	;
	v282 = v10 + int32(4)
	goto L6
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v74
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v282 = v73
	goto L6
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v67 - int32(4) {
	case 0, 2:
		goto L42
	default:
		v289 = v7
		goto L2
	}
L26:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v53 - int32(4) {
	case 0, 2:
		goto L39
	default:
		v289 = v7
		goto L2
	}
L27:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v282 = v52
	goto L6
L28:
	;
	v282 = v10 + int32(4)
	goto L6
L29:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v26 = int32(1)
	if base.Ui32(v26) < base.Ui32(v25-v26) {
		v289 = v7
		goto L2
	} else {
		goto L33
	}
L30:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v23
L31:
	;
	if v16 != int32(319) {
		v289 = v7
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v282 = v10 + int32(4)
	goto L6
L33:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v30 == int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.Ui32(v33-int32(4)) < base.Ui32(int32(-2)) {
		v289 = v7
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != int32(7) {
		v289 = v7
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 != int32(23) {
		v289 = v7
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)))
	if v46 != 0 {
		v289 = v7
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	return v47
L39:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v56 == int32(0) {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v59 != int32(67) {
		goto L7
	} else {
		goto L41
	}
L41:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v282 = v64 + int32(4)
	goto L6
L42:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	return v70
L43:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v82 = F_exprType(m, v78)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	return int32(0)
L45:
	;
	if v81 != v82 {
		v289 = v7
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v88 = F_exprTypmod(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	if v88 < int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v92 == int32(0) {
		v292 = v88
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v95 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v96 <= v95 {
		v292 = v88
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v99 = v95
	goto L51
L51:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v99<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v111 = F_exprType(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L44
	} else {
		goto L53
	}
L52:
	;
	v292 = v88
	goto L1
L53:
	;
	if v111 != v81 {
		v289 = v7
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v115 = F_exprTypmod(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	if v115 != v88 {
		v289 = v7
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v119 = v99 + int32(1)
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v119 < v120 {
		v99 = v119
		goto L51
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = F_exprTypmod(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	if v129 < int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v133 == int32(0) {
		v292 = v129
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v136 <= int32(0) {
		v292 = v129
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v141 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v142 = int32(4)
	goto L65
L64:
	;
	v142 = int32(12)
	goto L65
L65:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v10+v142)))
	v148 = v2
	goto L66
L66:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v133)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151+v148<<(uint(int32(2))%32))))
	v156 = F_exprType(m, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L44
	} else {
		goto L68
	}
L67:
	;
	v292 = v129
	goto L1
L68:
	;
	if v156 != v144 {
		v289 = v7
		goto L2
	} else {
		goto L69
	}
L69:
	;
	v159 = F_exprTypmod(m, v155)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L44
	} else {
		goto L70
	}
L70:
	;
	if v159 != v129 {
		v289 = v7
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v163 = v148 + int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v163 < v164 {
		v148 = v163
		goto L66
	} else {
		goto L72
	}
L72:
	;
	goto L67
L73:
	;
	if v166 != v170 {
		v289 = v7
		goto L2
	} else {
		goto L74
	}
L74:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = F_exprTypmod(m, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	if v176 < int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v180 == int32(0) {
		v292 = v176
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v183 = int32(1)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v184 <= v183 {
		v292 = v176
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v187 = v183
	goto L79
L79:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v193+v187<<(uint(int32(2))%32))))
	v198 = F_exprType(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L44
	} else {
		goto L81
	}
L80:
	;
	v292 = v176
	goto L1
L81:
	;
	if v198 != v166 {
		v289 = v7
		goto L2
	} else {
		goto L82
	}
L82:
	;
	v201 = F_exprTypmod(m, v197)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L44
	} else {
		goto L83
	}
L83:
	;
	if v201 != v176 {
		v289 = v7
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v205 = v187 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v205 < v206 {
		v187 = v205
		goto L79
	} else {
		goto L85
	}
L85:
	;
	goto L80
L86:
	;
	if v208 != v212 {
		v289 = v7
		goto L2
	} else {
		goto L87
	}
L87:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v218 = F_exprTypmod(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L44
	} else {
		goto L88
	}
L88:
	;
	if v218 < int32(0) {
		v289 = v7
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v222 == int32(0) {
		v292 = v218
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v225 = int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v226 <= v225 {
		v292 = v218
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v229 = v225
	goto L92
L92:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v222)+12))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v235+v229<<(uint(int32(2))%32))))
	v240 = F_exprType(m, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L44
	} else {
		goto L94
	}
L93:
	;
	v292 = v218
	goto L1
L94:
	;
	if v240 != v208 {
		v289 = v7
		goto L2
	} else {
		goto L95
	}
L95:
	;
	v243 = F_exprTypmod(m, v239)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L44
	} else {
		goto L96
	}
L96:
	;
	if v243 != v218 {
		v289 = v7
		goto L2
	} else {
		goto L97
	}
L97:
	;
	v247 = v229 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v247 < v248 {
		v229 = v247
		goto L92
	} else {
		goto L98
	}
L98:
	;
	goto L93
L99:
	;
	F_errmsg_internal(m, int32(_a_F_exprTypmod_0), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L44
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_exprTypmod_1), int32(350), int32(_a_F_exprTypmod_2))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L44
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	goto L5
}
func F_expr_setup_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v3 = int32(0)
	if l0 == v3 {
		v74 = v3
		return v74
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v10 = v8 - int32(6)
		if v10 != 0 {
			if v10 == int32(17) {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v53 != int32(5) {
					v70 = F_expression_tree_walker_impl(m, l0, int32(587), l1)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v74 = v70
						return v74
					}
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v57 = F_lappend(m, v56, l0)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v57
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v63 = v62
						if base.Ui32(v63-int32(9)) < base.Ui32(int32(3)) {
							v74 = v3
							return v74
						} else {
							v70 = F_expression_tree_walker_impl(m, l0, int32(587), l1)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v74 = v70
								return v74
							}
						}
					}
				}
			} else {
				v63 = v8
				if base.Ui32(v63-int32(9)) < base.Ui32(int32(3)) {
					v74 = v3
					return v74
				} else {
					v70 = F_expression_tree_walker_impl(m, l0, int32(587), l1)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						v74 = v70
						return v74
					}
				}
			}
		} else {
			v13 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v14 + int32(2) {
			case 0:
				v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+2)))
				v25 = base.I32_extend16_s(v13)
				if v25 < v24 {
					v27 = v24
				} else {
					v27 = v25
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)) = uint16(v27)
				return int32(0)
			case 1:
				v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
				v18 = base.I32_extend16_s(v13)
				if v18 < v17 {
					v20 = v17
				} else {
					v20 = v18
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v20)
				return int32(0)
			default:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				switch v31 {
				case 0:
					v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
					v33 = base.I32_extend16_s(v13)
					if v33 < v32 {
						v35 = v32
					} else {
						v35 = v33
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v35)
					return int32(0)
				case 1:
					v39 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
					v40 = base.I32_extend16_s(v13)
					if v40 < v39 {
						v42 = v39
					} else {
						v42 = v40
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v42)
					return int32(0)
				case 2:
					v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
					v47 = base.I32_extend16_s(v13)
					if v47 < v46 {
						v49 = v46
					} else {
						v49 = v47
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v49)
					return int32(0)
				default:
					v74 = v3
					return v74
				}
			}
		}
	}
}
func F_fix_scan_expr_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_fix_expr_common(m, v7, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v13 = F_expression_tree_walker_impl(m, l0, int32(840), l1)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
