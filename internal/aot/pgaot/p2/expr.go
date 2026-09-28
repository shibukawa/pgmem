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
	var v31 int64
	_ = v31
	var v39 int32
	_ = v39
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
	v11 = F_palloc0(m, int32(88))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(388)
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
			v31 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v31
			*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = l0
			v39 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+64)) = uint8(v39)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v31
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+48)) = uint8(v39)
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
	var v17 int64
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+80)) = v15
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
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	m.T0[v18].(func(*base.Module, int64))(m, v17)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
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
func F_exec_eval_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int64 {
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
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
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int64
	_ = v238
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int64
	_ = v264
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v310 int64
	_ = v310
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int64(0)
L5:
	;
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L107
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L101
	}
L8:
	;
	m.G0 = v16 - int32(-64)
	return v310
L9:
	;
	v251 = F_exec_run_select(m, l0, l1, int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L86
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
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
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[2]))
	F_ReleaseCachedPlan(m, v101, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L4
	} else {
		goto L85
	}
L17:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v179
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v185 != v32 {
		goto L71
	} else {
		goto L72
	}
L18:
	;
	if v80 != 0 {
		goto L32
	} else {
		goto L33
	}
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v46 = v44
	goto L21
L20:
	;
	v46 = int32(0)
	goto L21
L21:
	;
	if v41 == int32(0) {
		v78 = v6
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v80 = v78
	goto L18
L23:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+95)))
	if v49&int32(1) == int32(0) {
		v78 = v6
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v40)+88))
	if v41 != v54 {
		v78 = v6
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+10)))
	if v56 != int32(1) {
		v78 = v6
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
		v78 = v6
		goto L22
	} else {
		goto L28
	}
L28:
	;
	if v46 == int32(0) {
		v80 = int32(1)
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
	F_ResourceOwnerRemember(m, v46, base.I64_extend_i32_u(v41), int32(_a_F_exec_eval_expr_12))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v78 = v69
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
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	if v32 == v82 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_ReleaseCachedPlan(m, v84, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+60)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(l1)+48)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(0)
	v94 = int32(_a_F_exec_eval_expr_11)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v101 = F_SPI_plan_get_cached_plan(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v95
	v105 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	if v108 == v105 {
		v156 = v105
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v156 == int32(0) {
		goto L16
	} else {
		goto L66
	}
L41:
	;
	goto L40
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v111 != int32(1) {
		v156 = v105
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+60))
	if v116 == int32(0) {
		v156 = v105
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v119 != int32(1) {
		v156 = v105
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	if v124 != int32(67) {
		v156 = v105
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v127 != int32(1) {
		v156 = v105
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
	if v130 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L48
	}
L48:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+36)))
	if v131 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+37)))
	if v132 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L50
	}
L50:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+38)))
	if v133 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L51
	}
L51:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+39)))
	if v134 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v123)+48))
	if v135 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v123)+60))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v137 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L54
	}
L54:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	if v138 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v123)+100))
	if v139 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L56
	}
L56:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v123)+108))
	if v140 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v123)+112))
	if v141 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L58
	}
L58:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v123)+116))
	if v142 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L59
	}
L59:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v123)+120))
	if v143 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L60
	}
L60:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v123)+124))
	if v144 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L61
	}
L61:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v123)+128))
	if v145 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L62
	}
L62:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v123)+132))
	if v146 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L63
	}
L63:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v123)+144))
	if v147 != 0 {
		v156 = v105
		goto L41
	} else {
		goto L64
	}
L64:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v123)+76))
	if v148 == int32(0) {
		v156 = v105
		goto L41
	} else {
		goto L65
	}
L65:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v156 = base.B2i32(v151 == int32(1))
	goto L41
L66:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v161 = F_CachedPlanAllowsSimpleValidityCheck(m, v159, v101, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	if v161 == int32(0) {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v101
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[2]))
	F_ReleaseCachedPlan(m, v101, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_exec_save_simple_expr(m, l1, v101)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	goto L17
L71:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+100))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v192 = F_ExecInitExprWithParams(m, v191, v179)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
	if v203 != int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v32
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v195)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = v192
	goto L73
L75:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_eval_expr[1])) = v184
	v310 = v238
	goto L8
L76:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v228)+24))
	v230 = m.T0[v229].(func(*base.Module, int32, int32, int32) int64)(m, v228, v26, l2)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L84
	}
L77:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v206 != 0 {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v209 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_PushActiveSnapshot(m, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+24))
	v217 = m.T0[v216].(func(*base.Module, int32, int32, int32) int64)(m, v215, v26, l2)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v219)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v179)+20)) = v180
	F_PopActiveSnapshot(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v238 = v217
	goto L75
L84:
	;
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+72)) = uint8(v232)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v232
	*(*int32)(unsafe.Add(mBase, uint32(v179)+20)) = v180
	v238 = v230
	goto L75
L85:
	;
	goto L9
L86:
	;
	if v251 != int32(5) {
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	if v257 != int32(1) {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v256)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v256)+112))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v262
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	if base.Ui64(v264) <= base.Ui64(int64(1)) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v300 = F_SPI_getbinval(m, v298, v256, int32(1), l2)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L100
	}
L90:
	;
	if base.I32_wrap_i64(v264) == int32(1) {
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
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L94
	}
L93:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v270)
	v310 = int64(0)
	goto L8
L94:
	;
	F_errcode(m, int32(66))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_exec_eval_expr_9), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v287
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v16)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_10), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
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
	v310 = v300
	goto L8
L101:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_exec_eval_expr_1), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v329
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v14+int32(-16))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_4), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
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
	v347 = m.ExcPending
	if v347 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v350
	F_errmsg_plural(m, int32(_a_F_exec_eval_expr_6), int32(_a_F_exec_eval_expr_7), v350, v14+int32(-32))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	F_set_errcontext_domain(m, int32(_a_F_exec_eval_expr_0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v361
	F_errcontext_msg(m, int32(_a_F_exec_eval_expr_2), v14+int32(-48))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_exec_eval_expr_3), int32(_a_F_exec_eval_expr_8), int32(_a_F_exec_eval_expr_5))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 == v2 {
		v99 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v99
L2:
	;
	v11 = l0
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	switch v15 - int32(6) {
	case 0, 2, 8, 9, 11, 12, 13, 19, 23:
		goto L7
	case 1, 3, 5:
		goto L28
	case 4, 14, 15, 20, 24, 30, 31, 40, 46, 47, 52, 53:
		v99 = v2
		goto L1
	default:
		goto L8
	case 7:
		goto L27
	case 10, 54, 315:
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
	v99 = v2
	goto L1
L5:
	;
	if v97 != 0 {
		v11 = v97
		goto L3
	} else {
		goto L49
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L42
	} else {
		goto L46
	}
L7:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v99 = v82
	goto L1
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L42
	} else {
		goto L43
	}
L9:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v97 = v65
	goto L5
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v99 = v64
	goto L1
L11:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v99 = v63
	goto L1
L12:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v62 != 0 {
		v97 = v62
		goto L5
	} else {
		goto L41
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
	v99 = v61
	goto L1
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v59 != 0 {
		v97 = v59
		goto L5
	} else {
		goto L39
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v97 = v58
	goto L5
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v54 == int32(6) {
		goto L36
	} else {
		goto L37
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v48 == int32(19) {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v99 = v45
	goto L1
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v99 = v44
	goto L1
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v99 = v43
	goto L1
L21:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v99 = v42
	goto L1
L22:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v99 = v41
	goto L1
L23:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v97 = v40
	goto L5
L24:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v34 - int32(4) {
	case 0, 2:
		goto L32
	default:
		v99 = v2
		goto L1
	}
L25:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	switch v21 - int32(4) {
	case 0, 2:
		goto L29
	default:
		v99 = v2
		goto L1
	}
L26:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v97 = v20
	goto L5
L27:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v99 = v19
	goto L1
L28:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v99 = v18
	goto L1
L29:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v24 == int32(0) {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v27 != int32(67) {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v97 = v33
	goto L5
L32:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v99 = v37
	goto L1
L33:
	;
	v51 = int32(950)
	goto L35
L34:
	;
	v51 = int32(0)
	goto L35
L35:
	;
	v99 = v51
	goto L1
L36:
	;
	v57 = int32(100)
	goto L38
L37:
	;
	v57 = int32(0)
	goto L38
L38:
	;
	v99 = v57
	goto L1
L39:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	if v60 != 0 {
		v97 = v60
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v99 = v2
	goto L1
L41:
	;
	v99 = v2
	goto L1
L42:
	;
	return int32(0)
L43:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v72
	F_errmsg_internal(m, int32(_a_F_exprCollation_0), v7)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_exprCollation_1), int32(1070), int32(_a_F_exprCollation_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L42
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errmsg_internal(m, int32(_a_F_exprCollation_3), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_exprCollation_1), int32(889), int32(_a_F_exprCollation_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L42
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
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
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
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
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
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
	goto L11
L4:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	return v251
L5:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v240 = F_exprLocation(m, v239)
	mBase = m.M
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	if base.Ui32(v241) < base.Ui32(v240) {
		goto L168
	} else {
		goto L169
	}
L6:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v228 = F_exprLocation(m, v227)
	mBase = m.M
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if base.Ui32(v229) < base.Ui32(v228) {
		goto L159
	} else {
		goto L160
	}
L7:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v216 = F_exprLocation(m, v215)
	mBase = m.M
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	if base.Ui32(v217) < base.Ui32(v216) {
		goto L150
	} else {
		goto L151
	}
L8:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	return v213
L9:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	return v211
L10:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	return v209
L11:
	;
	v13 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	switch v15 - int32(1) {
	case 0:
		goto L44
	default:
		v206 = v13
		goto L13
	case 2, 7, 31, 38, 45, 111, 113:
		goto L10
	case 3:
		goto L43
	case 4, 24, 25, 30, 43, 59, 61, 73, 81, 82, 125, 133, 134, 320:
		v201 = int32(4)
		goto L14
	case 5:
		goto L42
	case 6:
		goto L41
	case 8:
		goto L40
	case 9, 35, 88, 94, 95, 131, 132, 208:
		goto L9
	case 10:
		goto L39
	case 12, 32, 41, 96, 98, 106, 109:
		goto L8
	case 13:
		goto L38
	case 14, 16, 17, 18, 19:
		goto L7
	case 15, 29, 51:
		goto L6
	case 20:
		goto L37
	case 21:
		goto L36
	case 26, 54:
		goto L5
	case 27:
		goto L35
	case 28:
		goto L34
	case 34:
		goto L33
	case 36:
		goto L32
	case 37, 39, 55, 56, 71, 79, 80, 110, 112, 129, 130:
		goto L4
	case 40:
		goto L31
	case 44:
		goto L30
	case 46:
		goto L28
	case 47:
		goto L29
	case 52:
		goto L27
	case 60:
		goto L26
	case 67:
		goto L20
	case 68, 69:
		goto L24
	case 70:
		goto L25
	case 72:
		goto L22
	case 75:
		goto L23
	case 83:
		goto L21
	case 89:
		goto L19
	case 91:
		goto L18
	case 97:
		goto L15
	case 114:
		goto L16
	case 160:
		goto L17
	}
L12:
	;
	return v206
L13:
	;
	goto L12
L14:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v9+v201)))
	if v203 != 0 {
		v9 = v203
		goto L11
	} else {
		goto L149
	}
L15:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v199
L16:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v197
L17:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v9)+104))
	return v195
L18:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	return v193
L19:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v9)+64))
	return v191
L20:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	return v189
L21:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v187
L22:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+28))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v169 = F_exprLocation(m, v168)
	mBase = m.M
	if base.Ui32(v169) < base.Ui32(v167) {
		goto L131
	} else {
		goto L132
	}
L23:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v154 = F_exprLocation(m, v153)
	mBase = m.M
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if base.Ui32(v155) < base.Ui32(v154) {
		goto L122
	} else {
		goto L123
	}
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	return v151
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v140 = F_exprLocation(m, v139)
	mBase = m.M
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v141) < base.Ui32(v140) {
		goto L113
	} else {
		goto L114
	}
L26:
	;
	v201 = int32(12)
	goto L14
L27:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v127 = F_exprLocation(m, v126)
	mBase = m.M
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v128) < base.Ui32(v127) {
		goto L104
	} else {
		goto L105
	}
L28:
	;
	v201 = int32(8)
	goto L14
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v114 = F_exprLocation(m, v113)
	mBase = m.M
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
	if base.Ui32(v115) < base.Ui32(v114) {
		goto L95
	} else {
		goto L96
	}
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	return v111
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	v100 = F_exprLocation(m, v99)
	mBase = m.M
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	if base.Ui32(v101) < base.Ui32(v100) {
		goto L86
	} else {
		goto L87
	}
L32:
	;
	v201 = int32(20)
	goto L14
L33:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	return v96
L34:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v85 = F_exprLocation(m, v84)
	mBase = m.M
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	if base.Ui32(v86) < base.Ui32(v85) {
		goto L77
	} else {
		goto L78
	}
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v73 = F_exprLocation(m, v72)
	mBase = m.M
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if base.Ui32(v74) < base.Ui32(v73) {
		goto L68
	} else {
		goto L69
	}
L36:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v62 = F_exprLocation(m, v61)
	mBase = m.M
	if base.Ui32(v62) < base.Ui32(v60) {
		goto L59
	} else {
		goto L60
	}
L37:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v49 = F_exprLocation(m, v48)
	mBase = m.M
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	if base.Ui32(v50) < base.Ui32(v49) {
		goto L50
	} else {
		goto L51
	}
L38:
	;
	v201 = int32(32)
	goto L14
L39:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	return v45
L40:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v43
L41:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	return v41
L42:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	return v39
L43:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+68))
	return v37
L44:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v18 <= int32(0) {
		v206 = v13
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v23 = int32(0)
	goto L46
L46:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21+v23<<(uint(int32(2))%32))))
	v31 = F_exprLocation(m, v30)
	mBase = m.M
	if int32(0) <= v31 {
		v206 = v31
		goto L13
	} else {
		goto L48
	}
L47:
	;
	v206 = v31
	goto L13
L48:
	;
	v35 = v23 + int32(1)
	if v35 != v18 {
		v23 = v35
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v52 = v50
	goto L52
L51:
	;
	v52 = v49
	goto L52
L52:
	;
	if v49 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v55 = v50
	goto L55
L54:
	;
	v55 = v52
	goto L55
L55:
	;
	if v50 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v58 = v49
	goto L58
L57:
	;
	v58 = v55
	goto L58
L58:
	;
	return v58
L59:
	;
	v64 = v62
	goto L61
L60:
	;
	v64 = v60
	goto L61
L61:
	;
	if v60 < int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v67 = v62
	goto L64
L63:
	;
	v67 = v64
	goto L64
L64:
	;
	if v62 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v70 = v60
	goto L67
L66:
	;
	v70 = v67
	goto L67
L67:
	;
	return v70
L68:
	;
	v76 = v74
	goto L70
L69:
	;
	v76 = v73
	goto L70
L70:
	;
	if v73 < int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v79 = v74
	goto L73
L72:
	;
	v79 = v76
	goto L73
L73:
	;
	if v74 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v82 = v73
	goto L76
L75:
	;
	v82 = v79
	goto L76
L76:
	;
	return v82
L77:
	;
	v88 = v86
	goto L79
L78:
	;
	v88 = v85
	goto L79
L79:
	;
	if v85 < int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v91 = v86
	goto L82
L81:
	;
	v91 = v88
	goto L82
L82:
	;
	if v86 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v94 = v85
	goto L85
L84:
	;
	v94 = v91
	goto L85
L85:
	;
	return v94
L86:
	;
	v103 = v101
	goto L88
L87:
	;
	v103 = v100
	goto L88
L88:
	;
	if v100 < int32(0) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v106 = v101
	goto L91
L90:
	;
	v106 = v103
	goto L91
L91:
	;
	if v101 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v109 = v100
	goto L94
L93:
	;
	v109 = v106
	goto L94
L94:
	;
	return v109
L95:
	;
	v117 = v115
	goto L97
L96:
	;
	v117 = v114
	goto L97
L97:
	;
	if v114 < int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v120 = v115
	goto L100
L99:
	;
	v120 = v117
	goto L100
L100:
	;
	if v115 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v123 = v114
	goto L103
L102:
	;
	v123 = v120
	goto L103
L103:
	;
	return v123
L104:
	;
	v130 = v128
	goto L106
L105:
	;
	v130 = v127
	goto L106
L106:
	;
	if v127 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v133 = v128
	goto L109
L108:
	;
	v133 = v130
	goto L109
L109:
	;
	if v128 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v136 = v127
	goto L112
L111:
	;
	v136 = v133
	goto L112
L112:
	;
	return v136
L113:
	;
	v143 = v141
	goto L115
L114:
	;
	v143 = v140
	goto L115
L115:
	;
	if v140 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v146 = v141
	goto L118
L117:
	;
	v146 = v143
	goto L118
L118:
	;
	if v141 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v149 = v140
	goto L121
L120:
	;
	v149 = v146
	goto L121
L121:
	;
	return v149
L122:
	;
	v157 = v155
	goto L124
L123:
	;
	v157 = v154
	goto L124
L124:
	;
	if v154 < int32(0) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v160 = v155
	goto L127
L126:
	;
	v160 = v157
	goto L127
L127:
	;
	if v155 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v163 = v154
	goto L130
L129:
	;
	v163 = v160
	goto L130
L130:
	;
	return v163
L131:
	;
	v171 = v169
	goto L133
L132:
	;
	v171 = v167
	goto L133
L133:
	;
	if v167 < int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v174 = v169
	goto L136
L135:
	;
	v174 = v171
	goto L136
L136:
	;
	if v169 < int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v177 = v167
	goto L139
L138:
	;
	v177 = v174
	goto L139
L139:
	;
	if base.Ui32(v177) < base.Ui32(v165) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v179 = v177
	goto L142
L141:
	;
	v179 = v165
	goto L142
L142:
	;
	if v165 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v182 = v177
	goto L145
L144:
	;
	v182 = v179
	goto L145
L145:
	;
	if v177 < int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v185 = v165
	goto L148
L147:
	;
	v185 = v182
	goto L148
L148:
	;
	return v185
L149:
	;
	v206 = v13
	goto L13
L150:
	;
	v219 = v217
	goto L152
L151:
	;
	v219 = v216
	goto L152
L152:
	;
	if v216 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v222 = v217
	goto L155
L154:
	;
	v222 = v219
	goto L155
L155:
	;
	if v217 < int32(0) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v225 = v216
	goto L158
L157:
	;
	v225 = v222
	goto L158
L158:
	;
	return v225
L159:
	;
	v231 = v229
	goto L161
L160:
	;
	v231 = v228
	goto L161
L161:
	;
	if v228 < int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v234 = v229
	goto L164
L163:
	;
	v234 = v231
	goto L164
L164:
	;
	if v229 < int32(0) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v237 = v228
	goto L167
L166:
	;
	v237 = v234
	goto L167
L167:
	;
	return v237
L168:
	;
	v243 = v241
	goto L170
L169:
	;
	v243 = v240
	goto L170
L170:
	;
	if v240 < int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v246 = v241
	goto L173
L172:
	;
	v246 = v243
	goto L173
L173:
	;
	if v241 < int32(0) {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v249 = v240
	goto L176
L175:
	;
	v249 = v246
	goto L176
L176:
	;
	return v249
}
func F_exprSetCollation(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = l0
	goto L5
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v40+v10))) = l1
	goto L1
L3:
	;
	v40 = int32(20)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L15
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v16 - int32(6) {
	case 0, 2, 8, 9, 11, 12, 13, 19, 23:
		goto L3
	case 1, 3, 5, 22, 50, 51:
		v40 = int32(12)
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
	v40 = int32(56)
	goto L2
L7:
	;
	goto L6
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v22 != 0 {
		v10 = v22
		goto L5
	} else {
		goto L12
	}
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v10 = v21
	goto L5
L10:
	;
	v40 = int32(16)
	goto L2
L11:
	;
	v40 = int32(8)
	goto L2
L12:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v23 != 0 {
		v10 = v23
		goto L5
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	return
L15:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
	F_errmsg_internal(m, int32(_a_F_exprSetCollation_0), v8)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_exprSetCollation_1), int32(1317), int32(_a_F_exprSetCollation_2))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
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
	if v16 != int32(321) {
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
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+32)))
	if v46 != 0 {
		v289 = v7
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
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
					v70 = F_expression_tree_walker_impl(m, l0, int32(633), l1)
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
							v70 = F_expression_tree_walker_impl(m, l0, int32(633), l1)
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
					v70 = F_expression_tree_walker_impl(m, l0, int32(633), l1)
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
			v13 = F_expression_tree_walker_impl(m, l0, int32(888), l1)
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
