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
	v5 = int32(4470560)
	v6 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v8
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
		v25 = F_AllocSetContextCreateInternal(m, v19, int32(61305), int32(0), int32(8192), int32(8388608))
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
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v6
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
	v6 = int32(4470560)
	v7 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v9
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v7
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
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
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
	F_errstart_cold(m, int32(21), int32(540539))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L107
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(540539))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v31 = *(*int32)(unsafe.Add(mBase, _consts[185]))
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
	v240 = *(*int32)(unsafe.Add(mBase, _consts[258]))
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
	v183 = *(*int32)(unsafe.Add(mBase, _consts[0]))
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
	F_ResourceOwnerRemember(m, v46, v41, int32(1714604))
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
	v93 = int32(4470560)
	v94 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v97
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v94
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
	v167 = *(*int32)(unsafe.Add(mBase, _consts[258]))
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v188
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v200
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
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v183
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
	F_errstart_cold(m, int32(21), int32(540539))
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
	F_errmsg(m, int32(29920), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_set_errcontext_domain(m, int32(540539))
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
	F_errcontext_msg(m, int32(195964), v16)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(491440), int32(5734), int32(203907))
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
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(495961), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_set_errcontext_domain(m, int32(540539))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v328
	F_errcontext_msg(m, int32(195964), v14+int32(-16))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(491440), int32(5697), int32(203907))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
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
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v349
	F_errmsg_plural(m, int32(270429), int32(146530), v349, v14+int32(-32))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_set_errcontext_domain(m, int32(540539))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v360
	F_errcontext_msg(m, int32(195964), v14+int32(-48))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(491440), int32(5709), int32(203907))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
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
	F_errmsg_internal(m, int32(477745), v6)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(486502), int32(1062), int32(258777))
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
	F_errmsg_internal(m, int32(310410), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(486502), int32(889), int32(258777))
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
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
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
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
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
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
	goto L4
L4:
	;
	v13 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v15 - int32(1) {
	case 0:
		goto L70
	default:
		v352 = v13
		goto L6
	case 2:
		goto L69
	case 3:
		goto L68
	case 4, 24, 25, 30, 43, 59, 61, 73, 81, 82, 125, 133, 134, 318:
		v347 = int32(4)
		goto L7
	case 5:
		goto L67
	case 6:
		goto L66
	case 7:
		goto L65
	case 8:
		goto L64
	case 9:
		goto L63
	case 10:
		goto L62
	case 12:
		goto L61
	case 13:
		goto L60
	case 14:
		goto L59
	case 15:
		goto L58
	case 16, 17, 18:
		goto L57
	case 19:
		goto L56
	case 20:
		goto L55
	case 21:
		goto L54
	case 26:
		goto L53
	case 27:
		goto L52
	case 28:
		goto L51
	case 29:
		goto L50
	case 31:
		goto L49
	case 32:
		goto L48
	case 34:
		goto L47
	case 35:
		goto L46
	case 36:
		goto L45
	case 37:
		goto L44
	case 38:
		goto L43
	case 39:
		goto L42
	case 40:
		goto L41
	case 41:
		goto L40
	case 44:
		goto L39
	case 45:
		goto L38
	case 46:
		goto L36
	case 47:
		goto L37
	case 51:
		goto L35
	case 52:
		goto L34
	case 54:
		goto L33
	case 55, 56:
		goto L32
	case 60:
		goto L31
	case 67:
		goto L22
	case 68, 69:
		goto L29
	case 70:
		goto L30
	case 71:
		goto L28
	case 72:
		goto L25
	case 75:
		goto L27
	case 79, 80:
		goto L26
	case 83:
		goto L24
	case 88:
		goto L23
	case 89:
		goto L21
	case 94, 208:
		goto L19
	case 95, 131, 132:
		goto L11
	case 96:
		goto L10
	case 97:
		goto L9
	case 98:
		goto L8
	case 106, 109:
		goto L18
	case 110:
		goto L17
	case 111:
		goto L16
	case 112:
		goto L15
	case 113:
		goto L14
	case 114:
		goto L13
	case 129, 130:
		goto L12
	case 160:
		goto L20
	}
L5:
	;
	return v352
L6:
	;
	goto L5
L7:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v9+v347)))
	if v349 != 0 {
		v9 = v349
		goto L4
	} else {
		goto L247
	}
L8:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v345
L9:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v343
L10:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v341
L11:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v339
L12:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v337
L13:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v335
L14:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v333
L15:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v331
L16:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v329
L17:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v327
L18:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v325
L19:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v323
L20:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
	return v321
L21:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	return v319
L22:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v317
L23:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v315
L24:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v313
L25:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+28))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v295 = F_exprLocation(m, v294)
	mBase = m.M
	if base.Ui32(v295) < base.Ui32(v293) {
		goto L229
	} else {
		goto L230
	}
L26:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v289
L27:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v278 = F_exprLocation(m, v277)
	mBase = m.M
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v279) < base.Ui32(v278) {
		goto L220
	} else {
		goto L221
	}
L28:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v275
L29:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	return v273
L30:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v262 = F_exprLocation(m, v261)
	mBase = m.M
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v263) < base.Ui32(v262) {
		goto L211
	} else {
		goto L212
	}
L31:
	;
	v347 = int32(12)
	goto L7
L32:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v258
L33:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v247 = F_exprLocation(m, v246)
	mBase = m.M
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v248) < base.Ui32(v247) {
		goto L202
	} else {
		goto L203
	}
L34:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v235 = F_exprLocation(m, v234)
	mBase = m.M
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v236) < base.Ui32(v235) {
		goto L193
	} else {
		goto L194
	}
L35:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v223 = F_exprLocation(m, v222)
	mBase = m.M
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v224) < base.Ui32(v223) {
		goto L184
	} else {
		goto L185
	}
L36:
	;
	v347 = int32(8)
	goto L7
L37:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v210 = F_exprLocation(m, v209)
	mBase = m.M
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if base.Ui32(v211) < base.Ui32(v210) {
		goto L175
	} else {
		goto L176
	}
L38:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v207
L39:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v205
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v203
L41:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v192 = F_exprLocation(m, v191)
	mBase = m.M
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if base.Ui32(v193) < base.Ui32(v192) {
		goto L166
	} else {
		goto L167
	}
L42:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v189
L43:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v187
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v185
L45:
	;
	v347 = int32(20)
	goto L7
L46:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v182
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v180
L48:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v178
L49:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v176
L50:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v165 = F_exprLocation(m, v164)
	mBase = m.M
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v166) < base.Ui32(v165) {
		goto L157
	} else {
		goto L158
	}
L51:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v153 = F_exprLocation(m, v152)
	mBase = m.M
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v154) < base.Ui32(v153) {
		goto L148
	} else {
		goto L149
	}
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v141 = F_exprLocation(m, v140)
	mBase = m.M
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if base.Ui32(v142) < base.Ui32(v141) {
		goto L139
	} else {
		goto L140
	}
L53:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v129 = F_exprLocation(m, v128)
	mBase = m.M
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v130) < base.Ui32(v129) {
		goto L130
	} else {
		goto L131
	}
L54:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v118 = F_exprLocation(m, v117)
	mBase = m.M
	if base.Ui32(v118) < base.Ui32(v116) {
		goto L121
	} else {
		goto L122
	}
L55:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v105 = F_exprLocation(m, v104)
	mBase = m.M
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v106) < base.Ui32(v105) {
		goto L112
	} else {
		goto L113
	}
L56:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v93 = F_exprLocation(m, v92)
	mBase = m.M
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v94) < base.Ui32(v93) {
		goto L103
	} else {
		goto L104
	}
L57:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v81 = F_exprLocation(m, v80)
	mBase = m.M
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v82) < base.Ui32(v81) {
		goto L94
	} else {
		goto L95
	}
L58:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v69 = F_exprLocation(m, v68)
	mBase = m.M
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v70) < base.Ui32(v69) {
		goto L85
	} else {
		goto L86
	}
L59:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v57 = F_exprLocation(m, v56)
	mBase = m.M
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v58) < base.Ui32(v57) {
		goto L76
	} else {
		goto L77
	}
L60:
	;
	v347 = int32(32)
	goto L7
L61:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v53
L62:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	return v51
L63:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v49
L64:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v47
L65:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v45
L66:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v43
L67:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	return v41
L68:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v39
L69:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v37
L70:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v18 <= int32(0) {
		v352 = v13
		goto L6
	} else {
		goto L71
	}
L71:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v23 = int32(0)
	goto L72
L72:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23<<(uint(int32(2))%32))))
	v31 = F_exprLocation(m, v30)
	mBase = m.M
	if int32(0) <= v31 {
		v352 = v31
		goto L6
	} else {
		goto L74
	}
L73:
	;
	v352 = v31
	goto L6
L74:
	;
	v35 = v23 + int32(1)
	if v35 != v18 {
		v23 = v35
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v60 = v58
	goto L78
L77:
	;
	v60 = v57
	goto L78
L78:
	;
	if v57 < int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v63 = v58
	goto L81
L80:
	;
	v63 = v60
	goto L81
L81:
	;
	if v58 < int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v66 = v57
	goto L84
L83:
	;
	v66 = v63
	goto L84
L84:
	;
	return v66
L85:
	;
	v72 = v70
	goto L87
L86:
	;
	v72 = v69
	goto L87
L87:
	;
	if v69 < int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v75 = v70
	goto L90
L89:
	;
	v75 = v72
	goto L90
L90:
	;
	if v70 < int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v78 = v69
	goto L93
L92:
	;
	v78 = v75
	goto L93
L93:
	;
	return v78
L94:
	;
	v84 = v82
	goto L96
L95:
	;
	v84 = v81
	goto L96
L96:
	;
	if v81 < int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v87 = v82
	goto L99
L98:
	;
	v87 = v84
	goto L99
L99:
	;
	if v82 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v90 = v81
	goto L102
L101:
	;
	v90 = v87
	goto L102
L102:
	;
	return v90
L103:
	;
	v96 = v94
	goto L105
L104:
	;
	v96 = v93
	goto L105
L105:
	;
	if v93 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v99 = v94
	goto L108
L107:
	;
	v99 = v96
	goto L108
L108:
	;
	if v94 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v102 = v93
	goto L111
L110:
	;
	v102 = v99
	goto L111
L111:
	;
	return v102
L112:
	;
	v108 = v106
	goto L114
L113:
	;
	v108 = v105
	goto L114
L114:
	;
	if v105 < int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v111 = v106
	goto L117
L116:
	;
	v111 = v108
	goto L117
L117:
	;
	if v106 < int32(0) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v114 = v105
	goto L120
L119:
	;
	v114 = v111
	goto L120
L120:
	;
	return v114
L121:
	;
	v120 = v118
	goto L123
L122:
	;
	v120 = v116
	goto L123
L123:
	;
	if v116 < int32(0) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v123 = v118
	goto L126
L125:
	;
	v123 = v120
	goto L126
L126:
	;
	if v118 < int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v126 = v116
	goto L129
L128:
	;
	v126 = v123
	goto L129
L129:
	;
	return v126
L130:
	;
	v132 = v130
	goto L132
L131:
	;
	v132 = v129
	goto L132
L132:
	;
	if v129 < int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v135 = v130
	goto L135
L134:
	;
	v135 = v132
	goto L135
L135:
	;
	if v130 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v138 = v129
	goto L138
L137:
	;
	v138 = v135
	goto L138
L138:
	;
	return v138
L139:
	;
	v144 = v142
	goto L141
L140:
	;
	v144 = v141
	goto L141
L141:
	;
	if v141 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v147 = v142
	goto L144
L143:
	;
	v147 = v144
	goto L144
L144:
	;
	if v142 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v150 = v141
	goto L147
L146:
	;
	v150 = v147
	goto L147
L147:
	;
	return v150
L148:
	;
	v156 = v154
	goto L150
L149:
	;
	v156 = v153
	goto L150
L150:
	;
	if v153 < int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v159 = v154
	goto L153
L152:
	;
	v159 = v156
	goto L153
L153:
	;
	if v154 < int32(0) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v162 = v153
	goto L156
L155:
	;
	v162 = v159
	goto L156
L156:
	;
	return v162
L157:
	;
	v168 = v166
	goto L159
L158:
	;
	v168 = v165
	goto L159
L159:
	;
	if v165 < int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v171 = v166
	goto L162
L161:
	;
	v171 = v168
	goto L162
L162:
	;
	if v166 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v174 = v165
	goto L165
L164:
	;
	v174 = v171
	goto L165
L165:
	;
	return v174
L166:
	;
	v195 = v193
	goto L168
L167:
	;
	v195 = v192
	goto L168
L168:
	;
	if v192 < int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v198 = v193
	goto L171
L170:
	;
	v198 = v195
	goto L171
L171:
	;
	if v193 < int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v201 = v192
	goto L174
L173:
	;
	v201 = v198
	goto L174
L174:
	;
	return v201
L175:
	;
	v213 = v211
	goto L177
L176:
	;
	v213 = v210
	goto L177
L177:
	;
	if v210 < int32(0) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v216 = v211
	goto L180
L179:
	;
	v216 = v213
	goto L180
L180:
	;
	if v211 < int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v219 = v210
	goto L183
L182:
	;
	v219 = v216
	goto L183
L183:
	;
	return v219
L184:
	;
	v226 = v224
	goto L186
L185:
	;
	v226 = v223
	goto L186
L186:
	;
	if v223 < int32(0) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v229 = v224
	goto L189
L188:
	;
	v229 = v226
	goto L189
L189:
	;
	if v224 < int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v232 = v223
	goto L192
L191:
	;
	v232 = v229
	goto L192
L192:
	;
	return v232
L193:
	;
	v238 = v236
	goto L195
L194:
	;
	v238 = v235
	goto L195
L195:
	;
	if v235 < int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v241 = v236
	goto L198
L197:
	;
	v241 = v238
	goto L198
L198:
	;
	if v236 < int32(0) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v244 = v235
	goto L201
L200:
	;
	v244 = v241
	goto L201
L201:
	;
	return v244
L202:
	;
	v250 = v248
	goto L204
L203:
	;
	v250 = v247
	goto L204
L204:
	;
	if v247 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v253 = v248
	goto L207
L206:
	;
	v253 = v250
	goto L207
L207:
	;
	if v248 < int32(0) {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v256 = v247
	goto L210
L209:
	;
	v256 = v253
	goto L210
L210:
	;
	return v256
L211:
	;
	v265 = v263
	goto L213
L212:
	;
	v265 = v262
	goto L213
L213:
	;
	if v262 < int32(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v268 = v263
	goto L216
L215:
	;
	v268 = v265
	goto L216
L216:
	;
	if v263 < int32(0) {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v271 = v262
	goto L219
L218:
	;
	v271 = v268
	goto L219
L219:
	;
	return v271
L220:
	;
	v281 = v279
	goto L222
L221:
	;
	v281 = v278
	goto L222
L222:
	;
	if v278 < int32(0) {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v284 = v279
	goto L225
L224:
	;
	v284 = v281
	goto L225
L225:
	;
	if v279 < int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v287 = v278
	goto L228
L227:
	;
	v287 = v284
	goto L228
L228:
	;
	return v287
L229:
	;
	v297 = v295
	goto L231
L230:
	;
	v297 = v293
	goto L231
L231:
	;
	if v293 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v300 = v295
	goto L234
L233:
	;
	v300 = v297
	goto L234
L234:
	;
	if v295 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v303 = v293
	goto L237
L236:
	;
	v303 = v300
	goto L237
L237:
	;
	if base.Ui32(v303) < base.Ui32(v291) {
		goto L238
	} else {
		goto L239
	}
L238:
	;
	v305 = v303
	goto L240
L239:
	;
	v305 = v291
	goto L240
L240:
	;
	if v291 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v308 = v303
	goto L243
L242:
	;
	v308 = v305
	goto L243
L243:
	;
	if v303 < int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v311 = v291
	goto L246
L245:
	;
	v311 = v308
	goto L246
L246:
	;
	return v311
L247:
	;
	v352 = v13
	goto L6
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
	F_errmsg_internal(m, int32(477745), v7)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(486502), int32(1306), int32(258760))
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
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = v4 - int32(9)
	if base.Ui32(int32(30)) < base.Ui32(v6) {
	} else {
		v10 = int32(1) << (uint(v6) % 32)
		if v10&int32(3904) == int32(0) {
			if v10&int32(5) != 0 {
				v22 = int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(l0+v22))) = l1
			} else {
				if v6 != int32(30) {
				} else {
					v22 = int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(l0+v22))) = l1
				}
			}
		} else {
			v22 = int32(24)
			*(*int32)(unsafe.Add(mBase, uint32(l0+v22))) = l1
		}
	}
	return
}
func F_exprTypmod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
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
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
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
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
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
	var v302 int32
	_ = v302
	v7 = int32(-1)
	if l0 == int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v302
L2:
	;
	v10 = l0
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v16 - int32(6) {
	case 0, 2, 8, 19, 23:
		goto L7
	case 1:
		goto L29
	case 3, 4, 5, 6, 7, 11, 12, 14, 15, 20, 22, 24, 27, 30, 31, 35, 36, 37, 40, 43, 44, 45, 46, 47, 48, 52, 53, 54:
		v302 = v7
		goto L1
	case 9:
		goto L28
	case 10:
		goto L27
	case 13:
		goto L26
	case 16:
		goto L25
	case 17:
		goto L24
	case 18:
		goto L23
	case 21:
		goto L22
	case 25:
		goto L21
	case 26:
		goto L20
	case 28:
		goto L19
	case 29:
		goto L18
	case 32:
		goto L17
	case 33:
		goto L16
	case 34:
		goto L15
	case 38:
		goto L14
	case 39:
		goto L13
	case 41:
		goto L11
	case 42:
		goto L12
	case 49:
		goto L10
	case 50, 51:
		goto L9
	case 55:
		goto L8
	default:
		goto L30
	}
L4:
	;
	v302 = v7
	goto L1
L5:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v296 != 0 {
		v10 = v296
		goto L3
	} else {
		goto L117
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L43
	} else {
		goto L114
	}
L7:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v302 = v281
	goto L1
L8:
	;
	v295 = v10 + int32(12)
	goto L5
L9:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v277
L10:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v275
L11:
	;
	v295 = v10 + int32(8)
	goto L5
L12:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
	return v271
L13:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	return v268
L14:
	;
	v295 = v10 + int32(8)
	goto L5
L15:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v263
L16:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+12))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v222 = F_exprType(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L43
	} else {
		goto L97
	}
L17:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v177 = F_exprType(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L43
	} else {
		goto L80
	}
L18:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v127 == int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L61
	}
L19:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v125
L20:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v78 == int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L42
	}
L21:
	;
	v295 = v10 + int32(4)
	goto L5
L22:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	return v74
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v295 = v73
	goto L5
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v67 - int32(4) {
	case 0, 2:
		goto L41
	default:
		v302 = v7
		goto L1
	}
L25:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v53 - int32(4) {
	case 0, 2:
		goto L38
	default:
		v302 = v7
		goto L1
	}
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v295 = v52
	goto L5
L27:
	;
	v295 = v10 + int32(4)
	goto L5
L28:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v26 = int32(1)
	if base.Ui32(v26) < base.Ui32(v25-v26) {
		v302 = v7
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	return v23
L30:
	;
	if v16 != int32(319) {
		v302 = v7
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v295 = v10 + int32(4)
	goto L5
L32:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v30 == int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if base.Ui32(v33-int32(4)) < base.Ui32(int32(-2)) {
		v302 = v7
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 != int32(7) {
		v302 = v7
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v43 != int32(23) {
		v302 = v7
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+24)))
	if v46 != 0 {
		v302 = v7
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	return v47
L38:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v56 == int32(0) {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v59 != int32(67) {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+76))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v295 = v64 + int32(4)
	goto L5
L41:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	return v70
L42:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v82 = F_exprType(m, v78)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	if v81 != v82 {
		v302 = v7
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v88 = F_exprTypmod(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v88 < int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v92 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	return v88
L49:
	;
	goto L50
L50:
	;
	v96 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v97 <= v96 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	return v88
L52:
	;
	goto L53
L53:
	;
	v101 = v96
	goto L54
L54:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v107+v101<<(uint(int32(2))%32))))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v113 = F_exprType(m, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L43
	} else {
		goto L56
	}
L55:
	;
	return v88
L56:
	;
	if v113 != v81 {
		v302 = v7
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v117 = F_exprTypmod(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L43
	} else {
		goto L58
	}
L58:
	;
	if v117 != v88 {
		v302 = v7
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v121 = v101 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if v121 < v122 {
		v101 = v121
		goto L54
	} else {
		goto L60
	}
L60:
	;
	goto L55
L61:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = F_exprTypmod(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L43
	} else {
		goto L62
	}
L62:
	;
	if v132 < int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v136 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return v132
L65:
	;
	goto L66
L66:
	;
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v141 <= v140 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	return v132
L68:
	;
	goto L69
L69:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v147 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v148 = int32(4)
	goto L72
L71:
	;
	v148 = int32(12)
	goto L72
L72:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v10+v148)))
	v152 = v140
	goto L73
L73:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157+v152<<(uint(int32(2))%32))))
	v162 = F_exprType(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L43
	} else {
		goto L75
	}
L74:
	;
	return v132
L75:
	;
	if v162 != v150 {
		v302 = v7
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v165 = F_exprTypmod(m, v161)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	if v165 != v132 {
		v302 = v7
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v169 = v152 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v169 < v170 {
		v152 = v169
		goto L73
	} else {
		goto L79
	}
L79:
	;
	goto L74
L80:
	;
	if v173 != v177 {
		v302 = v7
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+12))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v183 = F_exprTypmod(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L43
	} else {
		goto L82
	}
L82:
	;
	if v183 < int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v187 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	return v183
L85:
	;
	goto L86
L86:
	;
	v191 = int32(1)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v192 <= v191 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	return v183
L88:
	;
	goto L89
L89:
	;
	v196 = v191
	goto L90
L90:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v196<<(uint(int32(2))%32))))
	v207 = F_exprType(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L43
	} else {
		goto L92
	}
L91:
	;
	return v183
L92:
	;
	if v207 != v173 {
		v302 = v7
		goto L1
	} else {
		goto L93
	}
L93:
	;
	v210 = F_exprTypmod(m, v206)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L43
	} else {
		goto L94
	}
L94:
	;
	if v210 != v183 {
		v302 = v7
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v214 = v196 + int32(1)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v214 < v215 {
		v196 = v214
		goto L90
	} else {
		goto L96
	}
L96:
	;
	goto L91
L97:
	;
	if v218 != v222 {
		v302 = v7
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v225)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v228 = F_exprTypmod(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L43
	} else {
		goto L99
	}
L99:
	;
	if v228 < int32(0) {
		v302 = v7
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	if v232 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	return v228
L102:
	;
	goto L103
L103:
	;
	v236 = int32(1)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v237 <= v236 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	return v228
L105:
	;
	goto L106
L106:
	;
	v241 = v236
	goto L107
L107:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v247+v241<<(uint(int32(2))%32))))
	v252 = F_exprType(m, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L43
	} else {
		goto L109
	}
L108:
	;
	return v228
L109:
	;
	if v252 != v218 {
		v302 = v7
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v255 = F_exprTypmod(m, v251)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L43
	} else {
		goto L111
	}
L111:
	;
	if v255 != v228 {
		v302 = v7
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v259 = v241 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
	if v259 < v260 {
		v241 = v259
		goto L107
	} else {
		goto L113
	}
L113:
	;
	goto L108
L114:
	;
	F_errmsg_internal(m, int32(310457), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L43
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(486502), int32(350), int32(415206))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L43
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	goto L4
}
func F_expr_setup_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	v3 = int32(0)
	if l0 == v3 {
		v69 = v3
		return v69
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(6) {
		case 0:
			v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v11 + int32(2) {
			case 0:
				v21 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+2)))
				v22 = base.I32_extend16_s(v10)
				if v22 < v21 {
					v24 = v21
				} else {
					v24 = v22
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)) = uint16(v24)
				return int32(0)
			case 1:
				v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
				v15 = base.I32_extend16_s(v10)
				if v15 < v14 {
					v17 = v14
				} else {
					v17 = v15
				}
				*(*uint16)(unsafe.Add(mBase, uint32(l1))) = uint16(v17)
				return int32(0)
			default:
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				switch v28 {
				case 0:
					v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
					v30 = base.I32_extend16_s(v10)
					if v30 < v29 {
						v32 = v29
					} else {
						v32 = v30
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v32)
					return int32(0)
				case 1:
					v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
					v37 = base.I32_extend16_s(v10)
					if v37 < v36 {
						v39 = v36
					} else {
						v39 = v37
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v39)
					return int32(0)
				case 2:
					v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
					v44 = base.I32_extend16_s(v10)
					if v44 < v43 {
						v46 = v43
					} else {
						v46 = v44
					}
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v46)
					return int32(0)
				default:
					v69 = v3
					return v69
				}
			}
		default:
			v60 = v7
			if base.Ui32(v60-int32(9)) < base.Ui32(int32(3)) {
				v69 = v3
				return v69
			} else {
				v67 = F_expression_tree_walker_impl(m, l0, int32(586), l1)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v69 = v67
					return v69
				}
			}
		case 17:
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v50 != int32(5) {
				v67 = F_expression_tree_walker_impl(m, l0, int32(586), l1)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v69 = v67
					return v69
				}
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v54 = F_lappend(m, v53, l0)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v54
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v60 = v59
					if base.Ui32(v60-int32(9)) < base.Ui32(int32(3)) {
						v69 = v3
						return v69
					} else {
						v67 = F_expression_tree_walker_impl(m, l0, int32(586), l1)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							v69 = v67
							return v69
						}
					}
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
			v13 = F_expression_tree_walker_impl(m, l0, int32(839), l1)
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
