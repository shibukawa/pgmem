package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_plpgsql_build_datatype(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(82), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			F_errstart_cold(m, int32(21), int32(555940))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(50314), v9)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495586), int32(1960), int32(365906))
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
		} else {
			v31 = F_build_datatype(m, v12, l1, l2, l3)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v12)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v31
				}
			}
		}
	}
}
func F_plpgsql_call_handler(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
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
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int64
	_ = v222
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v341 int32
	_ = v341
	var v361 int32
	_ = v361
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int64
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
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
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	v2 = int32(0)
	v18 = m.G0
	v20 = v18 + int32(-64)
	m.G0 = v20
	v25 = v2
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
	v35 = v2
	v36 = v2
	v37 = int32(-1)
	v38 = v2
	v39 = v20
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	m.G0 = v20 - int32(-64)
	return v410
L4:
	;
	goto L3
L5:
	;
	if v37 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v377 = int32(m.ExcTag)
	v378 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v377 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	v43 = v39 - int32(16)
	m.G0 = v43
	v46 = v43 - int32(160)
	m.G0 = v46
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v48
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v51 == v48 {
		v60 = v48
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v138 = v25
	v139 = v26
	v140 = v27
	v141 = v28
	v142 = v29
	v143 = v30
	v144 = v31
	v145 = v32
	v146 = v33
	v147 = v34
	v148 = v35
	v149 = v36
	v151 = v38
	v152 = v39
	goto L10
L10:
	;
	if v151 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v25
	v70 = l0 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v43
	v75 = v60 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v75)
	F_SPI_connect_ext(m, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		v376 = v46
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v54 != int32(214) {
		v60 = v48
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+4)))
	v60 = v57 ^ int32(1)
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v75)
	v92 = F_plpgsql_compile(m, l0, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		v376 = v46
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v92)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+24)) = v94 + int64(1)
	v99 = v92 + int32(24)
	v101 = v92 + int32(528)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+528))
	v103 = int32(0)
	if v75 == v103 {
		v127 = v35
		v128 = v103
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v132 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	goto L20
L17:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+524)))
	if v108 != int32(1) {
		v127 = v35
		v128 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v75)
	v125 = F_ResourceOwnerCreate(m, int32(0), int32(171249))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		v376 = v46
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v127 = v125
	v128 = v125
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v18 + int32(-52)
	goto L23
L21:
	;
	v138 = v92
	v139 = v43
	v140 = v46
	v141 = v128
	v142 = v70
	v143 = v132
	v144 = v130
	v145 = v99
	v146 = v102
	v147 = v101
	v148 = v127
	v149 = v60
	v151 = v103
	v152 = v46
	goto L10
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v144
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v143
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = v222 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v146
	if v141 != 0 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v140
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	if v155 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v205 = int32(1)
	v206 = v149 & v205
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v206)
	v208 = int32(0)
	v214 = F_plpgsql_exec_function(m, v138, l0, v208, v208, v141, (v149^int32(-1))&v205)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L32
	}
L27:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	switch v158 - int32(441) {
	case 0:
		goto L28
	case 1:
		goto L29
	default:
		goto L26
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v190 = v149 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v190)
	F_plpgsql_exec_event_trigger(m, v138, v155)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v173 = v149 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v173)
	v175 = F_plpgsql_exec_trigger(m, v138, v155)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v175
	goto L24
L31:
	;
	goto L24
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v214
	goto L24
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v239 = v149 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v239)
	F_ReleaseAllPlanCacheRefsInOwner(m, v141)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v151 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v239)
	F_ResourceOwnerDelete(m, v141)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v270 = v149 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v270)
	F_pg_re_throw(m)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v144
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	v290 = v149 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v290)
	v292 = F_SPI_finish(m)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	if v292 == int32(2) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v290)
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v290)
	v324 = F_SPI_result_code_string(m, v292)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v324
	F_errmsg_internal(m, int32(203951), v20)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v20)+44)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+52)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v20)+56)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+60)) = v139
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)) = uint8(v290)
	F_errfinish(m, int32(495089), int32(302), int32(219358))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		v376 = v152
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	v382 = int32(v378)
	m.G0 = v376
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	if v18+int32(-52) == v389 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	m.ExcPending = 1
	goto L57
L50:
	;
	if v392 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	v392 = v391
	goto L53
L52:
	;
	v392 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+51)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	v25 = v397
	v26 = v393
	v27 = v394
	v28 = v402
	v29 = v395
	v30 = v403
	v31 = v404
	v32 = v400
	v33 = v399
	v34 = v398
	v35 = v401
	v36 = v396
	v37 = v392
	v38 = v384
	v39 = v376
	goto L2
L55:
	;
	goto L56
L56:
	;
	F___wasm_longjmp(m, v385, v384)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int32(0)
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_compile_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
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
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v478 int32
	_ = v478
	var v489 int32
	_ = v489
	var v509 int32
	_ = v509
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v987 int32
	_ = v987
	var v994 int32
	_ = v994
	var v999 int32
	_ = v999
	var v1003 int32
	_ = v1003
	var v1010 int32
	_ = v1010
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1042 int32
	_ = v1042
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1106 int32
	_ = v1106
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1138 int32
	_ = v1138
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1280 int32
	_ = v1280
	var v1282 int32
	_ = v1282
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1298 int32
	_ = v1298
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1370 int32
	_ = v1370
	var v1377 int32
	_ = v1377
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1418 int32
	_ = v1418
	var v1439 int32
	_ = v1439
	var v1445 int32
	_ = v1445
	var v1457 int32
	_ = v1457
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1465 int32
	_ = v1465
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1509 int32
	_ = v1509
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1532 int32
	_ = v1532
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1562 int32
	_ = v1562
	var v1565 int32
	_ = v1565
	var v1582 int32
	_ = v1582
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1594 int32
	_ = v1594
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1645 int32
	_ = v1645
	var v1659 int32
	_ = v1659
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1680 int32
	_ = v1680
	var v1686 int32
	_ = v1686
	var v1691 int32
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1702 int32
	_ = v1702
	var v1707 int32
	_ = v1707
	var v1711 int32
	_ = v1711
	var v1718 int32
	_ = v1718
	var v1723 int32
	_ = v1723
	var v1727 int32
	_ = v1727
	var v1733 int32
	_ = v1733
	var v1738 int32
	_ = v1738
	v5 = l4
	v20 = m.G0
	v22 = v20 - int32(416)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v26 = v24 + v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v30 == int32(441) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v40 = int32(2)
	goto L3
L3:
	;
	v43 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(26))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v33 = int32(1)
	goto L6
L5:
	;
	v33 = int32(2)
	goto L6
L6:
	;
	if v30 != int32(442) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v37 = v33
	goto L9
L8:
	;
	v37 = int32(0)
	goto L9
L9:
	;
	v40 = v37
	goto L3
L10:
	;
	return
L11:
	;
	v45 = F_text_to_cstring(m, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v47 = F_plpgsql_scanner_init(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v51 = v26 + int32(4)
	v52 = F_pstrdup(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = v52
	*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, _consts[1248])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v22)+412)) = v47
	v60 = int32(0)
	if v5 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v62 = v45
	goto L17
L16:
	;
	v62 = v60
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+408)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v22)+400)) = int32(6786)
	v66 = int32(4508024)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v22 + int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+396)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+404)) = v22 + int32(408)
	v77 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v82 = F_AllocSetContextCreateInternal(m, v77, int32(253954), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v85 = int32(4515120)
	v86 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[1249])) = v86
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v92 = F_format_procedure(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v82)+36)) = v92
	goto L20
L20:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v97
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v99
	v105 = *(*int32)(unsafe.Add(mBase, _consts[1250]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+488)) = v105
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1251])))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+492)) = uint8(v108)
	v111 = *(*int32)(unsafe.Add(mBase, _consts[1252]))
	if v5 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v113 = v111
	goto L23
L22:
	;
	v113 = int32(0)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+496)) = v113
	v116 = *(*int32)(unsafe.Add(mBase, _consts[1253]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v40
	if v5 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = v116
	goto L26
L25:
	;
	v119 = int32(0)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+500)) = v119
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+96)))
	v122 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+524)) = uint16(v122)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+520)) = v122
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, _consts[1254])) = v122
	goto L27
L27:
	;
	F_plpgsql_ns_push(m, v51, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1255])) = int32(128)
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[1256])) = uint8(v137)
	*(*int32)(unsafe.Add(mBase, _consts[1257])) = v137
	v143 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
	v145 = F_MemoryContextAlloc(m, v143, int32(512))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1258])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1259])) = v145
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	switch v152 {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L59
	default:
		goto L38
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L10
	} else {
		goto L395
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L10
	} else {
		goto L392
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L10
	} else {
		goto L388
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1680 = m.ExcPending
	if v1680 != 0 {
		goto L10
	} else {
		goto L385
	}
L34:
	;
	v1313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+64)) = uint8(base.B2i32(v1313 != int32(118)))
	v1319 = F_SearchSysCache1(m, int32(82), int32(16))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		goto L10
	} else {
		goto L312
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v1208
	v1210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)) = uint8(v1210)
	v1213 = F_SearchSysCache1(m, int32(82), v1208)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L10
	} else {
		goto L282
	}
L36:
	;
	v1208 = int32(23)
	goto L35
L37:
	;
	v1208 = int32(3904)
	goto L35
L38:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L10
	} else {
		goto L279
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L10
	} else {
		goto L276
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L10
	} else {
		goto L273
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L10
	} else {
		goto L269
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L10
	} else {
		goto L266
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L10
	} else {
		goto L263
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L10
	} else {
		goto L260
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L10
	} else {
		goto L257
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L10
	} else {
		goto L254
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L10
	} else {
		goto L251
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L10
	} else {
		goto L248
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L10
	} else {
		goto L245
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1003 = m.ExcPending
	if v1003 != 0 {
		goto L10
	} else {
		goto L242
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L10
	} else {
		goto L239
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L10
	} else {
		goto L234
	}
L53:
	;
	switch v489 - int32(5078) {
	case 0:
		v1208 = v509
		goto L35
	default:
		goto L36
	case 2:
		goto L37
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L10
	} else {
		goto L229
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L10
	} else {
		goto L226
	}
L56:
	;
	v1298 = v923
	v1312 = int32(0)
	goto L34
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(2278)
	v876 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	if v876 != 0 {
		goto L41
	} else {
		goto L215
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	v546 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	if v546 != 0 {
		goto L52
	} else {
		goto L150
	}
L59:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v155
	v163 = F_get_func_arg_info(m, l1, v22+int32(392), v22+int32(388), v22+int32(384))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v22)+392))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v22)+384))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v170 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
	F_cfunc_resolve_polymorphic_argtypes(m, v163, v165, v166, v168, v5, v170)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v174 = v163 << (uint(int32(2)) % 32)
	v175 = F_palloc(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v177 = F_palloc(m, v174)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
	if v163 <= int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v26)+108))
	if v489 <= int32(3830) {
		goto L128
	} else {
		goto L129
	}
L65:
	;
	v478 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v184 = int32(0)
	v188 = v60
	v194 = v184
	v196 = v184
	goto L68
L68:
	;
	v206 = v188 << (uint(int32(2)) % 32)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v22)+392))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v206+v207)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v22)+384))
	if v210 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v289 <= int32(1) {
		goto L102
	} else {
		goto L103
	}
L70:
	;
	v212 = int32(*(*int8)(unsafe.Add(mBase, uint32(v210+v188))))
	v214 = v212
	goto L72
L71:
	;
	v214 = int32(105)
	goto L72
L72:
	;
	v216 = v188 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v216
	v224 = F_pg_snprintf(m, v22+int32(352), int32(32), int32(466817), v22+int32(144))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v228 = F_SearchSysCache1(m, int32(82), v209)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	if v228 == int32(0) {
		goto L55
	} else {
		goto L75
	}
L75:
	;
	v234 = F_build_datatype(m, v228, int32(-1), v226, int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L76
	}
L76:
	;
	F_ReleaseCatCache(m, v228)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L10
	} else {
		goto L77
	}
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+8))
	if v238 == int32(2) {
		goto L54
	} else {
		goto L78
	}
L78:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v22)+388))
	if v241 != 0 {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v249 = int32(0)
	v251 = F_plpgsql_build_variable(m, v248, v249, v234, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L10
	} else {
		goto L84
	}
L80:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v241+v206)))
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v244 != 0 {
		v248 = v243
		goto L79
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v248 = v22 + int32(352)
	goto L79
L83:
	;
	goto L82
L84:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	v255 = v214 - int32(98)
	if base.Ui32(int32(20)) < base.Ui32(v255) {
		v271 = v196
		goto L85
	} else {
		goto L86
	}
L85:
	;
	if v253 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if int32(1)<<(uint(v255)%32)&int32(1048705) == int32(0) {
		v271 = v196
		goto L85
	} else {
		goto L87
	}
L87:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v175+v196<<(uint(int32(2))%32)))) = v267
	v271 = v196 + int32(1)
	goto L85
L88:
	;
	v274 = int32(2)
	goto L90
L89:
	;
	v274 = int32(1)
	goto L90
L90:
	;
	if base.Ui32(int32(18)) < base.Ui32(v255) {
		v289 = v194
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	F_add_parameter_name(m, v274, v290, v22+int32(352))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L94
	}
L92:
	;
	if int32(1)<<(uint(v255)%32)&int32(270337) == int32(0) {
		v289 = v194
		goto L91
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v177+v194<<(uint(int32(2))%32)))) = v251
	v289 = v194 + int32(1)
	goto L91
L94:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v22)+388))
	if v295 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v216 != v163 {
		v188 = v216
		v194 = v289
		v196 = v271
		goto L68
	} else {
		goto L99
	}
L96:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295+v206)))
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v300 == int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	F_add_parameter_name(m, v274, v303, v299)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	goto L95
L99:
	;
	goto L69
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = v451
	v478 = v458
	goto L64
L101:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v177)))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	v451 = v449
	v458 = v312
	goto L100
L102:
	;
	if v289 != int32(1) {
		v478 = v289
		goto L64
	} else {
		goto L105
	}
L103:
	;
	v316 = v289
	goto L104
L104:
	;
	v318 = F_palloc0(m, int32(40))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L10
	} else {
		goto L107
	}
L105:
	;
	v312 = int32(1)
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)))
	if v313 != int32(112) {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v316 = v312
	goto L104
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v318)+8)) = int32(670336)
	*(*int32)(unsafe.Add(mBase, uint32(v318))) = int32(1)
	v326 = F_CreateTemplateTupleDesc(m, v316)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+28)) = v316
	*(*int32)(unsafe.Add(mBase, uint32(v318)+24)) = v326
	v331 = v316 << (uint(int32(2)) % 32)
	v332 = F_palloc(m, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L10
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+32)) = v332
	v335 = F_palloc(m, v331)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+36)) = v335
	v340 = int32(0)
	goto L111
L111:
	;
	v359 = v340 << (uint(int32(2)) % 32)
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v177+v359)))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	switch v362 {
	case 0, 4:
		goto L114
	default:
		goto L115
	case 2:
		goto L116
	}
L112:
	;
	v419 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	v421 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v423 = *(*int32)(unsafe.Add(mBase, _consts[1255]))
	if v421 == v423 {
		goto L123
	} else {
		goto L124
	}
L113:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v318)+32))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v361)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v393+v359))) = v395
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v318)+36))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v397+v359))) = v399
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v318)+24))
	v403 = v340 + int32(1)
	v404 = base.I32_extend16_s(v403)
	F_TupleDescInitEntry(m, v401, v404, v395, v392, v391, int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L10
	} else {
		goto L120
	}
L114:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v361)+24))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v383)+16))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v383)+24))
	v388 = v383 + int32(4)
	v389 = v386
	v391 = v387
	goto L113
L115:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L10
	} else {
		goto L117
	}
L116:
	;
	v388 = v361 + int32(28)
	v389 = int32(0)
	v391 = int32(-1)
	goto L113
L117:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v371
	F_errmsg_internal(m, int32(484012), v22+int32(48))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L10
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(495586), int32(1879), int32(135972))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v318)+24))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	*(*int32)(unsafe.Add(mBase, uint32(v408+v409<<(uint(int32(4))%32)+v404*int32(100))+16)) = v389
	goto L121
L121:
	;
	if v403 != v316 {
		v340 = v403
		goto L111
	} else {
		goto L122
	}
L122:
	;
	goto L112
L123:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1255])) = v421 << (uint(int32(1)) % 32)
	v432 = F_repalloc(m, v419, v421<<(uint(int32(3))%32))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L10
	} else {
		goto L126
	}
L124:
	;
	v437 = v421
	v438 = v419
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+4)) = v437
	*(*int32)(unsafe.Add(mBase, _consts[1257])) = v437 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v438+v437<<(uint(int32(2))%32)))) = v318
	v451 = v437
	v458 = v316
	goto L100
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1259])) = v432
	v436 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v437 = v436
	v438 = v432
	goto L125
L127:
	;
	if v5 != 0 {
		goto L137
	} else {
		goto L138
	}
L128:
	;
	switch v489 - int32(2277) {
	case 0, 6:
		goto L127
	case 1, 2, 3, 4, 5:
		v1208 = v489
		goto L35
	default:
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	if base.Ui32(v489-int32(5077)) < base.Ui32(int32(4)) {
		goto L127
	} else {
		goto L134
	}
L131:
	;
	if v489 == int32(2776) {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	if v489 == int32(3500) {
		goto L127
	} else {
		goto L133
	}
L133:
	;
	v1208 = v489
	goto L35
L134:
	;
	if base.Ui32(v489-int32(4537)) < base.Ui32(int32(2)) {
		goto L127
	} else {
		goto L135
	}
L135:
	;
	if v489 == int32(3831) {
		goto L127
	} else {
		goto L136
	}
L136:
	;
	v1208 = v489
	goto L35
L137:
	;
	v509 = int32(1007)
	if int32(5077) < v489 {
		goto L53
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v520 = F_get_fn_expr_rettype(m, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L10
	} else {
		goto L144
	}
L140:
	;
	if v489 == int32(2277) {
		v1208 = v509
		goto L35
	} else {
		goto L141
	}
L141:
	;
	if v489 == int32(3831) {
		goto L37
	} else {
		goto L142
	}
L142:
	;
	if v489 != int32(4537) {
		goto L36
	} else {
		goto L143
	}
L143:
	;
	v1208 = int32(4451)
	goto L35
L144:
	;
	if v520 != 0 {
		v1208 = v520
		goto L35
	} else {
		goto L145
	}
L145:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L10
	} else {
		goto L146
	}
L146:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	v530 = *(*int32)(unsafe.Add(mBase, _consts[1246]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v530
	F_errmsg(m, int32(703422), v22+int32(128))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(495586), int32(423), int32(318423))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	v548 = F_palloc0(m, int32(40))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548))) = int32(2)
	v553 = F_pstrdup(m, int32(32151))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v548)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v548)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v548)+8)) = v553
	v563 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	v565 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v567 = *(*int32)(unsafe.Add(mBase, _consts[1255]))
	if v565 == v567 {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1255])) = v565 << (uint(int32(1)) % 32)
	v576 = F_repalloc(m, v563, v565<<(uint(int32(3))%32))
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L10
	} else {
		goto L156
	}
L154:
	;
	v582 = v565
	v583 = v553
	v584 = v563
	goto L155
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v548)+4)) = v582
	*(*int32)(unsafe.Add(mBase, _consts[1257])) = v582 + int32(1)
	v590 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v584+v582<<(uint(v590)%32)))) = v548
	F_plpgsql_ns_additem(m, v590, v582, v583)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L10
	} else {
		goto L157
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1259])) = v576
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v581 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v582 = v581
	v583 = v579
	v584 = v576
	goto L155
L157:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+480)) = v597
	v600 = F_palloc0(m, int32(40))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L10
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v600))) = int32(2)
	v605 = F_pstrdup(m, int32(429598))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v600)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v600)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v600)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v600)+8)) = v605
	v615 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	v617 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v619 = *(*int32)(unsafe.Add(mBase, _consts[1255]))
	if v617 == v619 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1255])) = v617 << (uint(int32(1)) % 32)
	v628 = F_repalloc(m, v615, v617<<(uint(int32(3))%32))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L10
	} else {
		goto L163
	}
L161:
	;
	v634 = v617
	v635 = v605
	v636 = v615
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v600)+4)) = v634
	*(*int32)(unsafe.Add(mBase, _consts[1257])) = v634 + int32(1)
	v642 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v636+v634<<(uint(v642)%32)))) = v600
	F_plpgsql_ns_additem(m, v642, v634, v635)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L10
	} else {
		goto L164
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1259])) = v628
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v600)+8))
	v633 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	v634 = v633
	v635 = v631
	v636 = v628
	goto L162
L164:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v600)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+484)) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v654 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	if v654 == int32(0) {
		goto L51
	} else {
		goto L166
	}
L166:
	;
	v660 = F_build_datatype(m, v654, int32(-1), v651, int32(0))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	F_ReleaseCatCache(m, v654)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	v667 = F_plpgsql_build_variable(m, int32(378906), int32(0), v660, int32(1))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v667)+48)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v667))) = int32(4)
	v673 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v676 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	if v676 == int32(0) {
		goto L50
	} else {
		goto L171
	}
L171:
	;
	v682 = F_build_datatype(m, v676, int32(-1), v673, int32(0))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	F_ReleaseCatCache(m, v676)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	v689 = F_plpgsql_build_variable(m, int32(282250), int32(0), v682, int32(1))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v689)+48)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v689))) = int32(4)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v698 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L10
	} else {
		goto L175
	}
L175:
	;
	if v698 == int32(0) {
		goto L49
	} else {
		goto L176
	}
L176:
	;
	v704 = F_build_datatype(m, v698, int32(-1), v695, int32(0))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	F_ReleaseCatCache(m, v698)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	v711 = F_plpgsql_build_variable(m, int32(305839), int32(0), v704, int32(1))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L10
	} else {
		goto L179
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711)+48)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v711))) = int32(4)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v720 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L10
	} else {
		goto L180
	}
L180:
	;
	if v720 == int32(0) {
		goto L48
	} else {
		goto L181
	}
L181:
	;
	v726 = F_build_datatype(m, v720, int32(-1), v717, int32(0))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	F_ReleaseCatCache(m, v720)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	v733 = F_plpgsql_build_variable(m, int32(235148), int32(0), v726, int32(1))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	v735 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v733)+48)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v733))) = v735
	v741 = F_SearchSysCache1(m, int32(82), int32(26))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L10
	} else {
		goto L185
	}
L185:
	;
	if v741 == int32(0) {
		goto L47
	} else {
		goto L186
	}
L186:
	;
	v746 = int32(0)
	v748 = F_build_datatype(m, v741, int32(-1), v746, v746)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	F_ReleaseCatCache(m, v741)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	v755 = F_plpgsql_build_variable(m, int32(434943), int32(0), v748, int32(1))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v755)+48)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v755))) = int32(4)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v764 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	if v764 == int32(0) {
		goto L46
	} else {
		goto L191
	}
L191:
	;
	v770 = F_build_datatype(m, v764, int32(-1), v761, int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	F_ReleaseCatCache(m, v764)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	v777 = F_plpgsql_build_variable(m, int32(377532), int32(0), v770, int32(1))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v777)+48)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v777))) = int32(4)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v786 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L10
	} else {
		goto L195
	}
L195:
	;
	if v786 == int32(0) {
		goto L45
	} else {
		goto L196
	}
L196:
	;
	v792 = F_build_datatype(m, v786, int32(-1), v783, int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	F_ReleaseCatCache(m, v786)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	v799 = F_plpgsql_build_variable(m, int32(379746), int32(0), v792, int32(1))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v799)+48)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v799))) = int32(4)
	v805 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v808 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L200
	}
L200:
	;
	if v808 == int32(0) {
		goto L44
	} else {
		goto L201
	}
L201:
	;
	v814 = F_build_datatype(m, v808, int32(-1), v805, int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L10
	} else {
		goto L202
	}
L202:
	;
	F_ReleaseCatCache(m, v808)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	v821 = F_plpgsql_build_variable(m, int32(505917), int32(0), v814, int32(1))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v821)+48)) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v821))) = int32(4)
	v829 = F_SearchSysCache1(m, int32(82), int32(23))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L10
	} else {
		goto L205
	}
L205:
	;
	if v829 == int32(0) {
		goto L43
	} else {
		goto L206
	}
L206:
	;
	v833 = int32(0)
	v837 = F_build_datatype(m, v829, int32(-1), v833, v833)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	F_ReleaseCatCache(m, v829)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	v844 = F_plpgsql_build_variable(m, int32(155090), int32(0), v837, int32(1))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844)+48)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = int32(4)
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v853 = F_SearchSysCache1(m, int32(82), int32(1009))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L10
	} else {
		goto L210
	}
L210:
	;
	if v853 == int32(0) {
		goto L42
	} else {
		goto L211
	}
L211:
	;
	v859 = F_build_datatype(m, v853, int32(-1), v850, int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	F_ReleaseCatCache(m, v853)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L10
	} else {
		goto L213
	}
L213:
	;
	v866 = F_plpgsql_build_variable(m, int32(35662), int32(0), v859, int32(1))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L10
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v866)+48)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v866))) = int32(4)
	v923 = v833
	goto L56
L215:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v880 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L10
	} else {
		goto L216
	}
L216:
	;
	if v880 == int32(0) {
		goto L40
	} else {
		goto L217
	}
L217:
	;
	v886 = F_build_datatype(m, v880, int32(-1), v877, int32(0))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	F_ReleaseCatCache(m, v880)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	v893 = F_plpgsql_build_variable(m, int32(91081), int32(0), v886, int32(1))
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L10
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v893)+48)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v893))) = int32(4)
	v899 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v902 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L10
	} else {
		goto L221
	}
L221:
	;
	if v902 == int32(0) {
		goto L39
	} else {
		goto L222
	}
L222:
	;
	v908 = F_build_datatype(m, v902, int32(-1), v899, int32(0))
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	F_ReleaseCatCache(m, v902)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	v915 = F_plpgsql_build_variable(m, int32(337533), int32(0), v908, int32(1))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v915)+48)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v915))) = int32(4)
	v923 = int32(0)
	goto L56
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v209
	F_errmsg_internal(m, int32(50314), v22+int32(16))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L229:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	v949 = F_format_type_be(m, v209)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v949
	F_errmsg(m, int32(188487), v22+int32(32))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(495586), int32(330), int32(318423))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	F_errmsg(m, int32(121954), int32(0))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		goto L10
	} else {
		goto L236
	}
L236:
	;
	F_errhint(m, int32(651873), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L10
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(495586), int32(496), int32(318423))
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L10
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = int32(19)
	F_errmsg_internal(m, int32(50314), v22+int32(160))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L10
	} else {
		goto L240
	}
L240:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v999 = m.ExcPending
	if v999 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = int32(25)
	F_errmsg_internal(m, int32(50314), v22+int32(176))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L10
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = int32(25)
	F_errmsg_internal(m, int32(50314), v22+int32(192))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = int32(25)
	F_errmsg_internal(m, int32(50314), v22+int32(208))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = int32(26)
	F_errmsg_internal(m, int32(50314), v22+int32(224))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L10
	} else {
		goto L252
	}
L252:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = int32(19)
	F_errmsg_internal(m, int32(50314), v22+int32(240))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L10
	} else {
		goto L255
	}
L255:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L10
	} else {
		goto L256
	}
L256:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = int32(19)
	F_errmsg_internal(m, int32(50314), v22+int32(256))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L10
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(19)
	F_errmsg_internal(m, int32(50314), v22+int32(272))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L10
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L10
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = int32(23)
	F_errmsg_internal(m, int32(50314), v22+int32(288))
	mBase = m.M
	v1122 = m.ExcPending
	if v1122 != 0 {
		goto L10
	} else {
		goto L264
	}
L264:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L10
	} else {
		goto L265
	}
L265:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+304)) = int32(1009)
	F_errmsg_internal(m, int32(50314), v22+int32(304))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L10
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L10
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1150 = m.ExcPending
	if v1150 != 0 {
		goto L10
	} else {
		goto L270
	}
L270:
	;
	F_errmsg(m, int32(121948), int32(0))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L10
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(495586), int32(629), int32(318423))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L10
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = int32(25)
	F_errmsg_internal(m, int32(50314), v22+int32(320))
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L10
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L10
	} else {
		goto L275
	}
L275:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = int32(25)
	F_errmsg_internal(m, int32(50314), v22+int32(336))
	mBase = m.M
	v1186 = m.ExcPending
	if v1186 != 0 {
		goto L10
	} else {
		goto L277
	}
L277:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1191 = m.ExcPending
	if v1191 != 0 {
		goto L10
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L279:
	;
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1196
	F_errmsg_internal(m, int32(486420), v22)
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L10
	} else {
		goto L280
	}
L280:
	;
	F_errfinish(m, int32(495586), int32(657), int32(318423))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L10
	} else {
		goto L281
	}
L281:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L282:
	;
	if v1213 == int32(0) {
		goto L33
	} else {
		goto L283
	}
L283:
	;
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1213)+16))
	v1218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217)+22)))
	v1219 = v1217 + v1218
	v1220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+79)))
	if v1220 != int32(112) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1249 = F_type_is_rowtype(m, v1208)
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L10
	} else {
		goto L294
	}
L285:
	;
	if v1208 == int32(2249) {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	switch v1208 - int32(2278) {
	case 0:
		goto L284
	case 1:
		goto L32
	default:
		goto L287
	}
L287:
	;
	if v1208 == int32(3838) {
		goto L32
	} else {
		goto L288
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L10
	} else {
		goto L289
	}
L289:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L10
	} else {
		goto L290
	}
L290:
	;
	v1236 = F_format_type_be(m, v1208)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L10
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v1236
	F_errmsg(m, int32(191282), v22+int32(80))
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L10
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(495586), int32(456), int32(318423))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L10
	} else {
		goto L293
	}
L293:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L294:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+61)) = uint8(v1249)
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+62)) = uint8(base.B2i32(v1252 == int32(100)))
	v1256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1219)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+60)) = uint8(v1256)
	v1258 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1219)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v1258
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(v26)+108))
	if v1260 <= int32(3830) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	F_ReleaseCatCache(m, v1213)
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L10
	} else {
		goto L311
	}
L296:
	;
	if v478 != 0 {
		goto L295
	} else {
		goto L308
	}
L297:
	;
	if v1260 != int32(4538) {
		goto L295
	} else {
		goto L307
	}
L298:
	;
	switch v1260 - int32(2277) {
	case 0, 6:
		goto L296
	case 1, 2, 3, 4, 5:
		goto L297
	default:
		goto L301
	}
L299:
	;
	goto L300
L300:
	;
	if base.Ui32(v1260-int32(5077)) < base.Ui32(int32(4)) {
		goto L296
	} else {
		goto L304
	}
L301:
	;
	if v1260 == int32(2776) {
		goto L296
	} else {
		goto L302
	}
L302:
	;
	if v1260 != int32(3500) {
		goto L297
	} else {
		goto L303
	}
L303:
	;
	goto L296
L304:
	;
	if v1260 == int32(3831) {
		goto L296
	} else {
		goto L305
	}
L305:
	;
	if v1260 == int32(4537) {
		goto L296
	} else {
		goto L306
	}
L306:
	;
	goto L297
L307:
	;
	goto L296
L308:
	;
	v1280 = int32(0)
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v1284 = F_build_datatype(m, v1213, int32(-1), v1282, v1280)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L10
	} else {
		goto L309
	}
L309:
	;
	v1287 = F_plpgsql_build_variable(m, int32(569801), v1280, v1284, int32(1))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L10
	} else {
		goto L310
	}
L310:
	;
	goto L295
L311:
	;
	v1298 = v175
	v1312 = base.B2i32(int32(0) < v478)
	goto L34
L312:
	;
	if v1319 == int32(0) {
		goto L31
	} else {
		goto L313
	}
L313:
	;
	v1324 = int32(0)
	v1326 = F_build_datatype(m, v1319, int32(-1), v1324, v1324)
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L10
	} else {
		goto L314
	}
L314:
	;
	F_ReleaseCatCache(m, v1319)
	mBase = m.M
	v1329 = m.ExcPending
	if v1329 != 0 {
		goto L10
	} else {
		goto L315
	}
L315:
	;
	v1333 = F_plpgsql_build_variable(m, int32(424135), int32(0), v1326, int32(1))
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L10
	} else {
		goto L316
	}
L316:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1333)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+476)) = v1335
	v1339 = F_plpgsql_yyparse(m, l3+int32(516), v47)
	mBase = m.M
	v1340 = m.ExcPending
	if v1340 != 0 {
		goto L10
	} else {
		goto L317
	}
L317:
	;
	if v1339 != 0 {
		goto L30
	} else {
		goto L318
	}
L318:
	;
	F_scanner_finish(m, v47)
	mBase = m.M
	v1342 = m.ExcPending
	if v1342 != 0 {
		goto L10
	} else {
		goto L319
	}
L319:
	;
	F_pfree(m, v45)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L10
	} else {
		goto L320
	}
L320:
	;
	if v1312 != 0 {
		goto L322
	} else {
		goto L323
	}
L321:
	;
	v1353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v1353
	if v1353 <= int32(0) {
		goto L327
	} else {
		goto L328
	}
L322:
	;
	F_add_dummy_return(m, l3)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L10
	} else {
		goto L326
	}
L323:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	if v1345 == int32(2278) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	v1348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)))
	if v1348 != int32(1) {
		goto L321
	} else {
		goto L325
	}
L325:
	;
	goto L322
L326:
	;
	goto L321
L327:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+504)) = v1487
	v1491 = F_palloc(m, v1487<<(uint(int32(2))%32))
	mBase = m.M
	v1492 = m.ExcPending
	if v1492 != 0 {
		goto L10
	} else {
		goto L339
	}
L328:
	;
	v1358 = v1353 & int32(3)
	v1360 = l3 + int32(72)
	v1361 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1353) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v1370 = v1361
	v1377 = int32(0)
	goto L332
L330:
	;
	v1418 = v1361
	goto L331
L331:
	;
	if v1358 == int32(0) {
		goto L327
	} else {
		goto L335
	}
L332:
	;
	v1388 = v1370 << (uint(int32(2)) % 32)
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1388)))
	*(*int32)(unsafe.Add(mBase, uint32(v1360+v1388))) = v1391
	v1393 = int32(4)
	v1394 = v1388 | v1393
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1394)))
	*(*int32)(unsafe.Add(mBase, uint32(v1360+v1394))) = v1397
	v1400 = v1388 | int32(8)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1400)))
	*(*int32)(unsafe.Add(mBase, uint32(v1360+v1400))) = v1403
	v1406 = v1388 | int32(12)
	v1409 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1406)))
	*(*int32)(unsafe.Add(mBase, uint32(v1360+v1406))) = v1409
	v1412 = v1370 + v1393
	v1414 = v1377 + v1393
	if v1414 != v1353&int32(32764) {
		v1370 = v1412
		v1377 = v1414
		goto L332
	} else {
		goto L334
	}
L333:
	;
	v1418 = v1412
	goto L331
L334:
	;
	goto L333
L335:
	;
	v1439 = v1418
	v1445 = v1361
	goto L336
L336:
	;
	v1457 = v1439 << (uint(int32(2)) % 32)
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1298+v1457)))
	*(*int32)(unsafe.Add(mBase, uint32(v1360+v1457))) = v1460
	v1462 = int32(1)
	v1465 = v1445 + v1462
	if v1465 != v1358 {
		v1439 = v1439 + v1462
		v1445 = v1465
		goto L336
	} else {
		goto L338
	}
L337:
	;
	goto L327
L338:
	;
	goto L337
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+508)) = v1491
	v1495 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
	if v1495 <= int32(0) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+512)) = v1594
	v1612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+525)))
	if v1612 == int32(1) {
		goto L360
	} else {
		goto L361
	}
L341:
	;
	v1594 = int32(0)
	goto L340
L342:
	;
	goto L343
L343:
	;
	v1499 = int32(1)
	v1502 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
	if v1495 == v1499 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if v1495&v1499 == int32(0) {
		v1594 = v1562
		goto L340
	} else {
		goto L357
	}
L345:
	;
	v1505 = int32(0)
	v1562 = v1505
	v1565 = v1505
	goto L344
L346:
	;
	goto L347
L347:
	;
	v1509 = int32(0)
	v1514 = v1509
	v1517 = v1509
	v1521 = v1509
	goto L348
L348:
	;
	v1532 = v1517 << (uint(int32(2)) % 32)
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1532+v1502)))
	*(*int32)(unsafe.Add(mBase, uint32(v1491+v1532))) = v1535
	v1537 = *(*int32)(unsafe.Add(mBase, uint32(v1535)))
	switch v1537 {
	case 0, 4:
		goto L352
	default:
		v1542 = v1514
		goto L350
	case 2:
		goto L351
	}
L349:
	;
	v1562 = v1554
	v1565 = v1556
	goto L344
L350:
	;
	v1544 = v1532 | int32(4)
	v1547 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v1502)))
	*(*int32)(unsafe.Add(mBase, uint32(v1491+v1544))) = v1547
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(v1547)))
	switch v1549 {
	case 0, 4:
		goto L354
	default:
		v1554 = v1542
		goto L353
	case 2:
		goto L355
	}
L351:
	;
	v1542 = v1514 + int32(40)
	goto L350
L352:
	;
	v1542 = v1514 + int32(56)
	goto L350
L353:
	;
	v1555 = int32(2)
	v1556 = v1517 + v1555
	v1558 = v1521 + v1555
	if v1558 != v1495&int32(2147483646) {
		v1514 = v1554
		v1517 = v1556
		v1521 = v1558
		goto L348
	} else {
		goto L356
	}
L354:
	;
	v1554 = v1542 + int32(56)
	goto L353
L355:
	;
	v1554 = v1542 + int32(40)
	goto L353
L356:
	;
	goto L349
L357:
	;
	v1582 = v1565 << (uint(int32(2)) % 32)
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v1582+v1502)))
	*(*int32)(unsafe.Add(mBase, uint32(v1491+v1582))) = v1585
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(v1585)))
	switch v1587 {
	case 0, 4:
		goto L358
	default:
		v1594 = v1562
		goto L340
	case 2:
		goto L359
	}
L358:
	;
	v1594 = v1562 + int32(56)
	goto L340
L359:
	;
	v1594 = v1562 + int32(40)
	goto L340
L360:
	;
	F_plpgsql_mark_local_assignment_targets(m, l3)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L10
	} else {
		goto L363
	}
L361:
	;
	goto L362
L362:
	;
	v1618 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
	if v1618 == int32(1) {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	goto L362
L364:
	;
	F_plpgsql_dumptree(m, l3)
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L10
	} else {
		goto L367
	}
L365:
	;
	goto L366
L366:
	;
	v1624 = *(*int32)(unsafe.Add(mBase, _consts[508]))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	if v1628 != v1624 {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	goto L366
L368:
	;
	v1659 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1246])) = v1659
	*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v1659)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v22)+396))
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v1665
	v1667 = int32(4608316)
	v1668 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
	*(*int32)(unsafe.Add(mBase, _consts[1249])) = v1659
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v1668
	m.G0 = v22 + int32(416)
	return
L369:
	;
	if v1628 == int32(0) {
		goto L372
	} else {
		goto L373
	}
L370:
	;
	goto L371
L371:
	;
	goto L368
L372:
	;
	if v1624 != 0 {
		goto L379
	} else {
		goto L380
	}
L373:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v1633 != 0 {
		goto L375
	} else {
		goto L376
	}
L374:
	;
	if v1632 == int32(0) {
		goto L372
	} else {
		goto L378
	}
L375:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1633)+28)) = v1632
	goto L374
L376:
	;
	goto L377
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1628)+20)) = v1632
	goto L374
L378:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1632)+24)) = v1638
	goto L372
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v1624
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(v1624)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v1645
	if v1645 != 0 {
		goto L382
	} else {
		goto L383
	}
L380:
	;
	goto L381
L381:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = int32(0)
	goto L371
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1645)+24)) = v82
	goto L384
L383:
	;
	goto L384
L384:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1624)+20)) = v82
	goto L368
L385:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1208
	F_errmsg_internal(m, int32(50314), v22-int32(-64))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L10
	} else {
		goto L386
	}
L386:
	;
	F_errfinish(m, int32(495586), int32(438), int32(318423))
	mBase = m.M
	v1691 = m.ExcPending
	if v1691 != 0 {
		goto L10
	} else {
		goto L387
	}
L387:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L388:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L10
	} else {
		goto L389
	}
L389:
	;
	F_errmsg(m, int32(134439), int32(0))
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L10
	} else {
		goto L390
	}
L390:
	;
	F_errfinish(m, int32(495586), int32(451), int32(318423))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L10
	} else {
		goto L391
	}
L391:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L392:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(16)
	F_errmsg_internal(m, int32(50314), v22+int32(96))
	mBase = m.M
	v1718 = m.ExcPending
	if v1718 != 0 {
		goto L10
	} else {
		goto L393
	}
L393:
	;
	F_errfinish(m, int32(495586), int32(1960), int32(365906))
	mBase = m.M
	v1723 = m.ExcPending
	if v1723 != 0 {
		goto L10
	} else {
		goto L394
	}
L394:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L395:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v1339
	F_errmsg_internal(m, int32(478904), v22+int32(112))
	mBase = m.M
	v1733 = m.ExcPending
	if v1733 != 0 {
		goto L10
	} else {
		goto L396
	}
L396:
	;
	F_errfinish(m, int32(495586), int32(680), int32(318423))
	mBase = m.M
	v1738 = m.ExcPending
	if v1738 != 0 {
		goto L10
	} else {
		goto L397
	}
L397:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_plpgsql_compile_inline(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = F_plpgsql_scanner_init(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[1246])) = int32(316325)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, _consts[447])))
		*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(6786)
		v31 = int32(4508024)
		v32 = *(*int32)(unsafe.Add(mBase, _consts[49]))
		*(*int32)(unsafe.Add(mBase, _consts[49])) = v14 + int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v14 + int32(40)
		v43 = F_palloc0(m, int32(536))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[1248])) = v43
			v47 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v52 = F_AllocSetContextCreateInternal(m, v47, int32(61780), int32(0), int32(8192), int32(8388608))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v55 = int32(4515120)
				v56 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, _consts[1249])) = v56
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
				v61 = F_pstrdup(m, int32(316325))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+472)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v52
					*(*int64)(unsafe.Add(mBase, uint32(v43)+40)) = int64(2)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v61
					v70 = *(*int32)(unsafe.Add(mBase, _consts[1250]))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+488)) = v70
					v73 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1251])))
					v74 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v43)+524)) = uint16(v74)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+520)) = v74
					*(*int64)(unsafe.Add(mBase, uint32(v43)+496)) = int64(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v43)+492)) = uint8(v73)
					*(*int32)(unsafe.Add(mBase, _consts[1254])) = v74
					F_plpgsql_ns_push(m, int32(316325), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[1255])) = int32(128)
						v92 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _consts[1256])) = uint8(v92)
						*(*int32)(unsafe.Add(mBase, _consts[1257])) = v92
						v98 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
						v100 = F_MemoryContextAlloc(m, v98, int32(512))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[1258])) = int32(0)
							*(*int32)(unsafe.Add(mBase, _consts[1259])) = v100
							*(*int64)(unsafe.Add(mBase, uint32(v43)+52)) = int64(17179871462)
							*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = int32(1)
							v111 = int32(26112)
							*(*uint16)(unsafe.Add(mBase, uint32(v43)+64)) = uint16(v111)
							v115 = F_SearchSysCache1(m, int32(82), int32(16))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int32(0)
							} else {
								if v115 != 0 {
									v118 = int32(0)
									v120 = F_build_datatype(m, v115, int32(-1), v118, v118)
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
										return int32(0)
									} else {
										F_ReleaseCatCache(m, v115)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											v127 = F_plpgsql_build_variable(m, int32(424135), int32(0), v120, int32(1))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
												*(*int32)(unsafe.Add(mBase, uint32(v43)+476)) = v129
												v133 = F_plpgsql_yyparse(m, v43+int32(516), v16)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													if v133 != 0 {
														F_errstart_cold(m, int32(21), int32(555940))
														mBase = m.M
														v291 = m.ExcPending
														if v291 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v133
															F_errmsg_internal(m, int32(478904), v14+int32(16))
															mBase = m.M
															v297 = m.ExcPending
															if v297 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(495586), int32(840), int32(373139))
																mBase = m.M
																v302 = m.ExcPending
																if v302 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														F_scanner_finish(m, v16)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v43)+52))
															if v137 == int32(2278) {
																F_add_dummy_return(m, v43)
																mBase = m.M
																v141 = m.ExcPending
																if v141 != 0 {
																	return int32(0)
																} else {
																	v142 = int32(0)
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v142
																	v146 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+504)) = v146
																	v150 = F_palloc(m, v146<<(uint(int32(2))%32))
																	mBase = m.M
																	v151 = m.ExcPending
																	if v151 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v150
																		v154 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
																		if v154 <= int32(0) {
																			v234 = v142
																		} else {
																			v157 = int32(1)
																			v160 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
																			if v154 == v157 {
																				v209 = int32(0)
																				v210 = v142
																			} else {
																				v169 = int32(0)
																				v170 = v142
																				v174 = int32(0)
																				for {
																					v179 = v169 << (uint(int32(2)) % 32)
																					v182 = *(*int32)(unsafe.Add(mBase, uint32(v179+v160)))
																					*(*int32)(unsafe.Add(mBase, uint32(v150+v179))) = v182
																					v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
																					switch v184 {
																					case 0, 4:
																						v189 = v170 + int32(56)
																					default:
																						v189 = v170
																					case 2:
																						v189 = v170 + int32(40)
																					}
																					v191 = v179 | int32(4)
																					v194 = *(*int32)(unsafe.Add(mBase, uint32(v191+v160)))
																					*(*int32)(unsafe.Add(mBase, uint32(v150+v191))) = v194
																					v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
																					switch v196 {
																					case 0, 4:
																						v201 = v189 + int32(56)
																					default:
																						v201 = v189
																					case 2:
																						v201 = v189 + int32(40)
																					}
																					v202 = int32(2)
																					v203 = v169 + v202
																					v205 = v174 + v202
																					if v205 != v154&int32(2147483646) {
																						v169 = v203
																						v170 = v201
																						v174 = v205
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				v209 = v203
																				v210 = v201
																			}
																			if v154&v157 == int32(0) {
																				v234 = v210
																			} else {
																				v221 = v209 << (uint(int32(2)) % 32)
																				v224 = *(*int32)(unsafe.Add(mBase, uint32(v221+v160)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v221))) = v224
																				v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
																				switch v226 {
																				case 0, 4:
																					v234 = v210 + int32(56)
																				default:
																					v234 = v210
																				case 2:
																					v234 = v210 + int32(40)
																				}
																			}
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v234
																		v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																		if v243 == int32(1) {
																			F_plpgsql_mark_local_assignment_targets(m, v43)
																			mBase = m.M
																			v247 = m.ExcPending
																			if v247 != 0 {
																				return int32(0)
																			} else {
																				v249 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
																				if v249 == int32(1) {
																					F_plpgsql_dumptree(m, v43)
																					mBase = m.M
																					v253 = m.ExcPending
																					if v253 != 0 {
																						return int32(0)
																					} else {
																						v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																						v258 = int32(0)
																						*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																						*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																						v263 = int32(4608316)
																						v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																						*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																						*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																						m.G0 = v14 + int32(48)
																						return v43
																					}
																				} else {
																					v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																					v258 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																					*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																					v263 = int32(4608316)
																					v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																					*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																					*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			}
																		} else {
																			v249 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
																			if v249 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v253 = m.ExcPending
																				if v253 != 0 {
																					return int32(0)
																				} else {
																					v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																					v258 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																					*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																					v263 = int32(4608316)
																					v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																					*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																					*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																				v258 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																				*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																				v263 = int32(4608316)
																				v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																				*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																				*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		}
																	}
																}
															} else {
																v142 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v142
																v146 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
																*(*int32)(unsafe.Add(mBase, uint32(v43)+504)) = v146
																v150 = F_palloc(m, v146<<(uint(int32(2))%32))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v150
																	v154 = *(*int32)(unsafe.Add(mBase, _consts[1257]))
																	if v154 <= int32(0) {
																		v234 = v142
																	} else {
																		v157 = int32(1)
																		v160 = *(*int32)(unsafe.Add(mBase, _consts[1259]))
																		if v154 == v157 {
																			v209 = int32(0)
																			v210 = v142
																		} else {
																			v169 = int32(0)
																			v170 = v142
																			v174 = int32(0)
																			for {
																				v179 = v169 << (uint(int32(2)) % 32)
																				v182 = *(*int32)(unsafe.Add(mBase, uint32(v179+v160)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v179))) = v182
																				v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
																				switch v184 {
																				case 0, 4:
																					v189 = v170 + int32(56)
																				default:
																					v189 = v170
																				case 2:
																					v189 = v170 + int32(40)
																				}
																				v191 = v179 | int32(4)
																				v194 = *(*int32)(unsafe.Add(mBase, uint32(v191+v160)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v191))) = v194
																				v196 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
																				switch v196 {
																				case 0, 4:
																					v201 = v189 + int32(56)
																				default:
																					v201 = v189
																				case 2:
																					v201 = v189 + int32(40)
																				}
																				v202 = int32(2)
																				v203 = v169 + v202
																				v205 = v174 + v202
																				if v205 != v154&int32(2147483646) {
																					v169 = v203
																					v170 = v201
																					v174 = v205
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			v209 = v203
																			v210 = v201
																		}
																		if v154&v157 == int32(0) {
																			v234 = v210
																		} else {
																			v221 = v209 << (uint(int32(2)) % 32)
																			v224 = *(*int32)(unsafe.Add(mBase, uint32(v221+v160)))
																			*(*int32)(unsafe.Add(mBase, uint32(v150+v221))) = v224
																			v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
																			switch v226 {
																			case 0, 4:
																				v234 = v210 + int32(56)
																			default:
																				v234 = v210
																			case 2:
																				v234 = v210 + int32(40)
																			}
																		}
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v234
																	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																	if v243 == int32(1) {
																		F_plpgsql_mark_local_assignment_targets(m, v43)
																		mBase = m.M
																		v247 = m.ExcPending
																		if v247 != 0 {
																			return int32(0)
																		} else {
																			v249 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
																			if v249 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v253 = m.ExcPending
																				if v253 != 0 {
																					return int32(0)
																				} else {
																					v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																					v258 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																					*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																					v263 = int32(4608316)
																					v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																					*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																					*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																				v258 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																				*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																				v263 = int32(4608316)
																				v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																				*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																				*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		}
																	} else {
																		v249 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1256])))
																		if v249 == int32(1) {
																			F_plpgsql_dumptree(m, v43)
																			mBase = m.M
																			v253 = m.ExcPending
																			if v253 != 0 {
																				return int32(0)
																			} else {
																				v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																				v258 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																				*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																				v263 = int32(4608316)
																				v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																				*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																				*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		} else {
																			v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																			*(*int32)(unsafe.Add(mBase, _consts[49])) = v255
																			v258 = int32(0)
																			*(*int32)(unsafe.Add(mBase, _consts[1246])) = v258
																			*(*uint8)(unsafe.Add(mBase, _consts[1247])) = uint8(v258)
																			v263 = int32(4608316)
																			v264 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
																			*(*int32)(unsafe.Add(mBase, _consts[1249])) = v258
																			*(*int32)(unsafe.Add(mBase, _consts[0])) = v264
																			m.G0 = v14 + int32(48)
																			return v43
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
								} else {
									F_errstart_cold(m, int32(21), int32(555940))
									mBase = m.M
									v277 = m.ExcPending
									if v277 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(16)
										F_errmsg_internal(m, int32(50314), v14)
										mBase = m.M
										v282 = m.ExcPending
										if v282 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495586), int32(1960), int32(365906))
											mBase = m.M
											v287 = m.ExcPending
											if v287 != 0 {
												return int32(0)
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
			}
		}
	}
}
func F_plpgsql_delete_callback(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_plpgsql_free_function_memory(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_plpgsql_exec_get_datum_type(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v10 {
	case 0, 4:
		v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v74 = v68 + int32(4)
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
		m.G0 = v8 + int32(32)
		return v75
	default:
		F_errstart_cold(m, int32(21), int32(555940))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v55
			F_errmsg_internal(m, int32(484012), v8)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(499747), int32(5508), int32(366297))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v11 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v12 == int32(2249) {
				v74 = v11 + int32(36)
			} else {
				v74 = l1 + int32(28)
			}
		} else {
			v74 = l1 + int32(28)
		}
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
		m.G0 = v8 + int32(32)
		return v75
	case 3:
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v19+v20<<(uint(int32(2))%32))))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
		if v25 == int32(0) {
			F_instantiate_empty_record_variable(m, l0, v24)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
				v33 = v32
				v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
				v35 = *(*int64)(unsafe.Add(mBase, uint32(v33)+48))
				if v34 != v35 {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v40 = F_expanded_record_lookup_field(m, v33, v37, l1+int32(32))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 == int32(0) {
							F_errstart_cold(m, int32(21), int32(555940))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v89
									F_errmsg(m, int32(722557), v8+int32(16))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499747), int32(5499), int32(366297))
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
							v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+48))
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v45
							v74 = l1 + int32(36)
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
							m.G0 = v8 + int32(32)
							return v75
						}
					}
				} else {
					v74 = l1 + int32(36)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					m.G0 = v8 + int32(32)
					return v75
				}
			}
		} else {
			v33 = v25
			v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v33)+48))
			if v34 != v35 {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v40 = F_expanded_record_lookup_field(m, v33, v37, l1+int32(32))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if v40 == int32(0) {
						F_errstart_cold(m, int32(21), int32(555940))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
								v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v89
								F_errmsg(m, int32(722557), v8+int32(16))
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499747), int32(5499), int32(366297))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+48))
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = v45
						v74 = l1 + int32(36)
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
						m.G0 = v8 + int32(32)
						return v75
					}
				}
			} else {
				v74 = l1 + int32(36)
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
				m.G0 = v8 + int32(32)
				return v75
			}
		}
	}
}
func F_plpgsql_latest_lineno(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+196))
	return v3
}
func F_plpgsql_ns_pop(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	if v4 != 0 {
		v5 = v3
		for {
			v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v7 != 0 {
				v5 = v6
				continue
			} else {
				break
			}
			break
		}
		v8 = v6
	} else {
		v8 = v3
	}
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, _consts[1254])) = v10
	return
}
func F_plpgsql_param_compile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+68))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = v16 - int32(1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15+v18<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(54)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	switch v28 {
	case 0:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
		v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+12)))
		if v30 == int32(65535) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			if v18 != v33 {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6798)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
				if v35 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6798)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6797)
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6799)
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6802)
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6801)
	case 3:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6800)
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
		if v47 == int32(65535) {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6801)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(6802)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l1
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v61
	F_ExprEvalPushStep(m, l2, v12+int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		return
	} else {
		m.G0 = v12 + int32(48)
		return
	}
}
func F_plpgsql_param_ref(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v118 int32
	_ = v118
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
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
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	v21 = F_pg_snprintf(m, v12+int32(16), int32(32), int32(466817), v12)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if v26 == v25 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	if v160 != 0 {
		goto L40
	} else {
		goto L41
	}
L4:
	;
	goto L3
L5:
	;
	v160 = v54
	goto L4
L7:
	;
	v160 = int32(0)
	goto L4
L8:
	;
	v40 = v26
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v54 = v40
	v56 = v49
	goto L14
L12:
	;
	v76 = v40
	goto L13
L13:
	;
	goto L21
L14:
	;
	v61 = F_strcmp(m, v54+int32(12), v12+int32(16))
	mBase = m.M
	if v61 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v76 = v70
	goto L13
L16:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != 0 {
		v54 = v70
		v56 = v71
		goto L14
	} else {
		goto L20
	}
L17:
	;
	if base.B2i32(v56 == int32(1))&int32(0) != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	goto L5
L20:
	;
	goto L15
L21:
	;
	goto L37
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v118 != 0 {
		v40 = v118
		goto L9
	} else {
		goto L38
	}
L38:
	;
	goto L10
L40:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+528))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+68))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173+v174<<(uint(int32(2))%32))))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v180 = int32(4515120)
	v181 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v171)+48))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v185 = F_bms_add_member(m, v184, v174)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v206 = v25
	goto L42
L42:
	;
	m.G0 = v12 + int32(48)
	return v206
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v185
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v181
	v190 = F_palloc0(m, int32(28))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v174 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v190))) = int64(8)
	F_plpgsql_exec_get_datum_type_info(m, v172, v178, v190+int32(12), v190+int32(16), v190+int32(20))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v190)+24)) = v179
	v206 = v190
	goto L42
}
func F_plpgsql_parse_err_condition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(134238)
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1260])))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == v2 {
		v36 = v16
		v37 = v17
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v125
L2:
	;
	if v37-v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v16 != v17 {
		v36 = v16
		v37 = v17
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v21 = l0
	v22 = v12
	goto L6
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v25
		v37 = v26
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v36 = v25
	v37 = v26
	goto L3
L8:
	;
	v29 = int32(1)
	if v25 == v26 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v44 = int32(350418)
	v45 = v2
	v47 = int32(4174032)
	v48 = v2
	goto L13
L11:
	;
	goto L12
L12:
	;
	v116 = F_palloc(m, int32(12))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L26
	} else {
		goto L34
	}
L13:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v52 == int32(0) {
		v71 = v51
		v72 = v52
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v86 != 0 {
		v125 = v86
		goto L1
	} else {
		goto L29
	}
L15:
	;
	if v72-v71 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	goto L15
L17:
	;
	if v51 != v52 {
		v71 = v51
		v72 = v52
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v56 = l0
	v57 = v44
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v61 == int32(0) {
		v71 = v60
		v72 = v61
		goto L16
	} else {
		goto L21
	}
L20:
	;
	v71 = v60
	v72 = v61
	goto L16
L21:
	;
	v64 = int32(1)
	if v60 == v61 {
		v56 = v56 + v64
		v57 = v57 + v64
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	v77 = F_palloc(m, int32(12))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v86 = v45
	goto L25
L25:
	;
	v88 = v48 + int32(1)
	v90 = v88 << (uint(int32(3)) % 32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)+uint32(_consts[1261])))
	if v93 != 0 {
		v44 = v93
		v45 = v86
		v47 = v90 + int32(4174032)
		v48 = v88
		goto L13
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = l0
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v83
	v86 = v77
	goto L25
L28:
	;
	goto L14
L29:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(703164), v9)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(495586), int32(2190), int32(250502))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L26
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(-1)
	v125 = v116
	goto L1
}
func F_plpgsql_peek(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v12 = F_internal_yylex(m, v8+int32(8), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
		if int32(4) <= v17 {
			F_errstart_cold(m, int32(21), int32(555940))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(319264), int32(0))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(495059), int32(388), int32(282055))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))+76)) = v12
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
			v45 = v16 + v42*int32(24)
			v46 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v45)+108)) = v46
			v48 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v45)+100)) = v48
			v50 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v45)+92)) = v50
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+72))
			*(*int32)(unsafe.Add(mBase, uint32(v52)+72)) = v53 + int32(1)
			m.G0 = v8 + int32(32)
			return v12
		}
	}
}
func F_plpgsql_post_column_ref(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+488))
	if v14 == int32(1) {
		v63 = v4
		m.G0 = v10 + int32(16)
		return v63
	} else {
		if v14 == int32(2) {
			v20 = l2
		} else {
			v20 = int32(0)
		}
		if v20 != 0 {
			v63 = v4
			m.G0 = v10 + int32(16)
			return v63
		} else {
			v23 = F_resolve_column_ref(m, l0, v12, l1, base.B2i32(l2 == int32(0)))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if l2 == int32(0) {
					v63 = v23
					m.G0 = v10 + int32(16)
					return v63
				} else {
					if v23 == int32(0) {
						v63 = v23
						m.G0 = v10 + int32(16)
						return v63
					} else {
						F_errstart_cold(m, int32(21), int32(555940))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(33583236))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								v41 = F_NameListToString(m, v40)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = v41
									F_errmsg(m, int32(114801), v10)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_errdetail(m, int32(618720), int32(0))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											F_parser_errposition(m, l0, v53)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495586), int32(1046), int32(339098))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
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
				}
			}
		}
	}
}
func F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
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
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v926 int32
	_ = v926
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v1001 int32
	_ = v1001
	var v1008 int32
	_ = v1008
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 {
	case 0:
		goto L26
	case 1:
		goto L25
	case 2:
		goto L24
	case 3:
		goto L23
	case 4:
		goto L22
	case 5:
		goto L21
	case 6:
		goto L20
	case 7:
		goto L19
	case 8:
		goto L18
	case 9:
		goto L17
	case 10:
		goto L16
	case 11:
		goto L15
	case 12:
		goto L14
	case 13:
		goto L13
	case 14:
		goto L12
	case 15:
		goto L11
	case 16:
		goto L10
	case 17:
		goto L9
	case 18:
		goto L8
	case 19, 22, 25, 26:
		goto L2
	case 20:
		goto L7
	case 21:
		goto L6
	case 23:
		goto L5
	case 24:
		goto L4
	default:
		goto L1
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(555940))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L32
	} else {
		goto L284
	}
L2:
	;
	m.G0 = v11 + int32(16)
	return
L3:
	;
	v976 = F_bms_is_member(m, v971, l1)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L32
	} else {
		goto L283
	}
L4:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v962 == int32(0) {
		goto L2
	} else {
		goto L281
	}
L5:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v956 == int32(0) {
		goto L2
	} else {
		goto L279
	}
L6:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v950 == int32(0) {
		goto L2
	} else {
		goto L277
	}
L7:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v887 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L8:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v811 == int32(0) {
		goto L236
	} else {
		goto L237
	}
L9:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v768 == int32(0) {
		goto L223
	} else {
		goto L224
	}
L10:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v762 == int32(0) {
		goto L2
	} else {
		goto L221
	}
L11:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v746 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L12:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v672 == int32(0) {
		goto L196
	} else {
		goto L197
	}
L13:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v619 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L14:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v613 == int32(0) {
		goto L2
	} else {
		goto L177
	}
L15:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v607 == int32(0) {
		goto L2
	} else {
		goto L175
	}
L16:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v601 == int32(0) {
		goto L2
	} else {
		goto L173
	}
L17:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v565 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L18:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v526 == int32(0) {
		goto L154
	} else {
		goto L155
	}
L19:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v487 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L20:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v431 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L21:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v395 == int32(0) {
		goto L117
	} else {
		goto L118
	}
L22:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v370 == int32(0) {
		goto L2
	} else {
		goto L111
	}
L23:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v259 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v108 == int32(0) {
		goto L2
	} else {
		goto L48
	}
L26:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v47 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L28:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v22 = v3
	goto L30
L30:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22<<(uint(int32(2))%32))))
	F_mark_stmt(m, v32, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L27
L32:
	;
	return
L33:
	;
	v36 = v22 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v36 < v37 {
		v22 = v36
		goto L30
	} else {
		goto L34
	}
L34:
	;
	goto L31
L35:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v50 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 <= int32(0) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v57 = int32(0)
	goto L38
L38:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v57<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v70 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L2
L40:
	;
	v105 = v57 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v105 < v106 {
		v57 = v105
		goto L38
	} else {
		goto L47
	}
L41:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74 <= v73 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v79 = v73
	goto L43
L43:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v79<<(uint(int32(2))%32))))
	F_mark_stmt(m, v89, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L32
	} else {
		goto L45
	}
L44:
	;
	goto L40
L45:
	;
	v93 = v79 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v93 < v94 {
		v79 = v93
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L39
L48:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if int32(0) <= v111 {
		v970 = v108
		v971 = v111
		goto L3
	} else {
		goto L49
	}
L49:
	;
	goto L2
L50:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v124 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v117 < int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v120 = F_bms_is_member(m, v117, l1)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+20)) = uint8(v120)
	goto L50
L54:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v158 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L55:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v127 <= int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v133 = int32(0)
	goto L57
L57:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v133<<(uint(int32(2))%32))))
	F_mark_stmt(m, v143, l1)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L32
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	v147 = v133 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v147 < v148 {
		v133 = v147
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v233 == int32(0) {
		goto L2
	} else {
		goto L78
	}
L62:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v161 <= int32(0) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v168 = v3
	goto L64
L64:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v168<<(uint(int32(2))%32))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v177 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	goto L61
L66:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	if v187 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	if v180 < int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v183 = F_bms_is_member(m, v180, l1)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L32
	} else {
		goto L69
	}
L69:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+20)) = uint8(v183)
	goto L66
L70:
	;
	v222 = v168 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v222 < v223 {
		v168 = v222
		goto L64
	} else {
		goto L77
	}
L71:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v191 <= v190 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v196 = v190
	goto L73
L73:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v196<<(uint(int32(2))%32))))
	F_mark_stmt(m, v206, l1)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L32
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	v210 = v196 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v210 < v211 {
		v196 = v210
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	goto L65
L78:
	;
	v236 = int32(0)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v237 <= v236 {
		goto L2
	} else {
		goto L79
	}
L79:
	;
	v242 = v236
	goto L80
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248+v242<<(uint(int32(2))%32))))
	F_mark_stmt(m, v252, l1)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L32
	} else {
		goto L82
	}
L81:
	;
	goto L2
L82:
	;
	v256 = v242 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v256 < v257 {
		v242 = v256
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v269 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v262 < int32(0) {
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v265 = F_bms_is_member(m, v262, l1)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L32
	} else {
		goto L87
	}
L87:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+20)) = uint8(v265)
	goto L84
L88:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v344 == int32(0) {
		goto L2
	} else {
		goto L105
	}
L89:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v272 <= int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v279 = v3
	goto L91
L91:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v279<<(uint(int32(2))%32))))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v288 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	goto L88
L93:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v298 == int32(0) {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	if v291 < int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v294 = F_bms_is_member(m, v291, l1)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L32
	} else {
		goto L96
	}
L96:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+20)) = uint8(v294)
	goto L93
L97:
	;
	v333 = v279 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v333 < v334 {
		v279 = v333
		goto L91
	} else {
		goto L104
	}
L98:
	;
	v301 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v302 <= v301 {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v307 = v301
	goto L100
L100:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313+v307<<(uint(int32(2))%32))))
	F_mark_stmt(m, v317, l1)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L32
	} else {
		goto L102
	}
L101:
	;
	goto L97
L102:
	;
	v321 = v307 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v321 < v322 {
		v307 = v321
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L92
L105:
	;
	v347 = int32(0)
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v348 <= v347 {
		goto L2
	} else {
		goto L106
	}
L106:
	;
	v353 = v347
	goto L107
L107:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v353<<(uint(int32(2))%32))))
	F_mark_stmt(m, v363, l1)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L32
	} else {
		goto L109
	}
L108:
	;
	goto L2
L109:
	;
	v367 = v353 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v367 < v368 {
		v353 = v367
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v373 <= int32(0) {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	v378 = v3
	goto L113
L113:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v378<<(uint(int32(2))%32))))
	F_mark_stmt(m, v388, l1)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L32
	} else {
		goto L115
	}
L114:
	;
	goto L2
L115:
	;
	v392 = v378 + int32(1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v392 < v393 {
		v378 = v392
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v405 == int32(0) {
		goto L2
	} else {
		goto L121
	}
L118:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v398 < int32(0) {
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v401 = F_bms_is_member(m, v398, l1)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L32
	} else {
		goto L120
	}
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)) = uint8(v401)
	goto L117
L121:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v408 <= int32(0) {
		goto L2
	} else {
		goto L122
	}
L122:
	;
	v414 = int32(0)
	goto L123
L123:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v414<<(uint(int32(2))%32))))
	F_mark_stmt(m, v424, l1)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L32
	} else {
		goto L125
	}
L124:
	;
	goto L2
L125:
	;
	v428 = v414 + int32(1)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v428 < v429 {
		v414 = v428
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v441 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+16))
	if v434 < int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v437 = F_bms_is_member(m, v434, l1)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L32
	} else {
		goto L130
	}
L130:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+20)) = uint8(v437)
	goto L127
L131:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v451 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L132:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	if v444 < int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v447 = F_bms_is_member(m, v444, l1)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L32
	} else {
		goto L134
	}
L134:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v441)+20)) = uint8(v447)
	goto L131
L135:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v461 == int32(0) {
		goto L2
	} else {
		goto L139
	}
L136:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451)+16))
	if v454 < int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v457 = F_bms_is_member(m, v454, l1)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L32
	} else {
		goto L138
	}
L138:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+20)) = uint8(v457)
	goto L135
L139:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v464 <= int32(0) {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	v470 = int32(0)
	goto L141
L141:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476+v470<<(uint(int32(2))%32))))
	F_mark_stmt(m, v480, l1)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L32
	} else {
		goto L143
	}
L142:
	;
	goto L2
L143:
	;
	v484 = v470 + int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v484 < v485 {
		v470 = v484
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v520 == int32(0) {
		goto L2
	} else {
		goto L152
	}
L146:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v490 <= int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v495 = v3
	goto L148
L148:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v495<<(uint(int32(2))%32))))
	F_mark_stmt(m, v505, l1)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L32
	} else {
		goto L150
	}
L149:
	;
	goto L145
L150:
	;
	v509 = v495 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v509 < v510 {
		v495 = v509
		goto L148
	} else {
		goto L151
	}
L151:
	;
	goto L149
L152:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+16))
	if int32(0) <= v523 {
		v970 = v520
		v971 = v523
		goto L3
	} else {
		goto L153
	}
L153:
	;
	goto L2
L154:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v559 == int32(0) {
		goto L2
	} else {
		goto L161
	}
L155:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v529 <= int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v534 = v3
	goto L157
L157:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v534<<(uint(int32(2))%32))))
	F_mark_stmt(m, v544, l1)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L32
	} else {
		goto L159
	}
L158:
	;
	goto L154
L159:
	;
	v548 = v534 + int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v548 < v549 {
		v534 = v548
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v559)+16))
	if int32(0) <= v562 {
		v970 = v559
		v971 = v562
		goto L3
	} else {
		goto L162
	}
L162:
	;
	goto L2
L163:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v575 == int32(0) {
		goto L2
	} else {
		goto L167
	}
L164:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	if v568 < int32(0) {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v571 = F_bms_is_member(m, v568, l1)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L32
	} else {
		goto L166
	}
L166:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+20)) = uint8(v571)
	goto L163
L167:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v578 <= int32(0) {
		goto L2
	} else {
		goto L168
	}
L168:
	;
	v584 = int32(0)
	goto L169
L169:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v590+v584<<(uint(int32(2))%32))))
	F_mark_stmt(m, v594, l1)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L32
	} else {
		goto L171
	}
L170:
	;
	goto L2
L171:
	;
	v598 = v584 + int32(1)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v598 < v599 {
		v584 = v598
		goto L169
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	if int32(0) <= v604 {
		v970 = v601
		v971 = v604
		goto L3
	} else {
		goto L174
	}
L174:
	;
	goto L2
L175:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+16))
	if int32(0) <= v610 {
		v970 = v607
		v971 = v610
		goto L3
	} else {
		goto L176
	}
L176:
	;
	goto L2
L177:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	if int32(0) <= v616 {
		v970 = v613
		v971 = v616
		goto L3
	} else {
		goto L178
	}
L178:
	;
	goto L2
L179:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v629 == int32(0) {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	if v622 < int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v625 = F_bms_is_member(m, v622, l1)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L32
	} else {
		goto L182
	}
L182:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v619)+20)) = uint8(v625)
	goto L179
L183:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v639 == int32(0) {
		goto L2
	} else {
		goto L187
	}
L184:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v629)+16))
	if v632 < int32(0) {
		goto L183
	} else {
		goto L185
	}
L185:
	;
	v635 = F_bms_is_member(m, v632, l1)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L32
	} else {
		goto L186
	}
L186:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v629)+20)) = uint8(v635)
	goto L183
L187:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	if v642 <= int32(0) {
		goto L2
	} else {
		goto L188
	}
L188:
	;
	v648 = int32(0)
	goto L189
L189:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v639)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v648<<(uint(int32(2))%32))))
	if v658 == int32(0) {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	goto L2
L191:
	;
	v669 = v648 + int32(1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	if v669 < v670 {
		v648 = v669
		goto L189
	} else {
		goto L195
	}
L192:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v658)+16))
	if v661 < int32(0) {
		goto L191
	} else {
		goto L193
	}
L193:
	;
	v664 = F_bms_is_member(m, v661, l1)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L32
	} else {
		goto L194
	}
L194:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+20)) = uint8(v664)
	goto L191
L195:
	;
	goto L190
L196:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v712 == int32(0) {
		goto L2
	} else {
		goto L206
	}
L197:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v675 <= int32(0) {
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v680 = v3
	goto L199
L199:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v686+v680<<(uint(int32(2))%32))))
	if v690 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	goto L196
L201:
	;
	v701 = v680 + int32(1)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v701 < v702 {
		v680 = v701
		goto L199
	} else {
		goto L205
	}
L202:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v690)+16))
	if v693 < int32(0) {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v696 = F_bms_is_member(m, v693, l1)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L32
	} else {
		goto L204
	}
L204:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v690)+20)) = uint8(v696)
	goto L201
L205:
	;
	goto L200
L206:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v715 <= int32(0) {
		goto L2
	} else {
		goto L207
	}
L207:
	;
	v721 = int32(0)
	goto L208
L208:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v721<<(uint(int32(2))%32))))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	if v732 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	goto L2
L210:
	;
	v743 = v721 + int32(1)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v743 < v744 {
		v721 = v743
		goto L208
	} else {
		goto L214
	}
L211:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v732)+16))
	if v735 < int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v738 = F_bms_is_member(m, v735, l1)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L32
	} else {
		goto L213
	}
L213:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v732)+20)) = uint8(v738)
	goto L210
L214:
	;
	goto L209
L215:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v756 == int32(0) {
		goto L2
	} else {
		goto L219
	}
L216:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v746)+16))
	if v749 < int32(0) {
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v752 = F_bms_is_member(m, v749, l1)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L32
	} else {
		goto L218
	}
L218:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+20)) = uint8(v752)
	goto L215
L219:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v756)+16))
	if int32(0) <= v759 {
		v970 = v756
		v971 = v759
		goto L3
	} else {
		goto L220
	}
L220:
	;
	goto L2
L221:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v762)+16))
	if int32(0) <= v765 {
		v970 = v762
		v971 = v765
		goto L3
	} else {
		goto L222
	}
L222:
	;
	goto L2
L223:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v778 == int32(0) {
		goto L2
	} else {
		goto L227
	}
L224:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v768)+16))
	if v771 < int32(0) {
		goto L223
	} else {
		goto L225
	}
L225:
	;
	v774 = F_bms_is_member(m, v771, l1)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L32
	} else {
		goto L226
	}
L226:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v768)+20)) = uint8(v774)
	goto L223
L227:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v781 <= int32(0) {
		goto L2
	} else {
		goto L228
	}
L228:
	;
	v787 = int32(0)
	goto L229
L229:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793+v787<<(uint(int32(2))%32))))
	if v797 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L230:
	;
	goto L2
L231:
	;
	v808 = v787 + int32(1)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v808 < v809 {
		v787 = v808
		goto L229
	} else {
		goto L235
	}
L232:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v797)+16))
	if v800 < int32(0) {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v803 = F_bms_is_member(m, v800, l1)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L32
	} else {
		goto L234
	}
L234:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v797)+20)) = uint8(v803)
	goto L231
L235:
	;
	goto L230
L236:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v844 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L237:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v814 <= int32(0) {
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v819 = v3
	goto L239
L239:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v825+v819<<(uint(int32(2))%32))))
	F_mark_stmt(m, v829, l1)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L32
	} else {
		goto L241
	}
L240:
	;
	goto L236
L241:
	;
	v833 = v819 + int32(1)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v833 < v834 {
		v819 = v833
		goto L239
	} else {
		goto L242
	}
L242:
	;
	goto L240
L243:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v854 == int32(0) {
		goto L2
	} else {
		goto L247
	}
L244:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v844)+16))
	if v847 < int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v850 = F_bms_is_member(m, v847, l1)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L32
	} else {
		goto L246
	}
L246:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v844)+20)) = uint8(v850)
	goto L243
L247:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	if v857 <= int32(0) {
		goto L2
	} else {
		goto L248
	}
L248:
	;
	v863 = int32(0)
	goto L249
L249:
	;
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v854)+12))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v869+v863<<(uint(int32(2))%32))))
	if v873 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L250:
	;
	goto L2
L251:
	;
	v884 = v863 + int32(1)
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v854)+4))
	if v884 < v885 {
		v863 = v884
		goto L249
	} else {
		goto L255
	}
L252:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v873)+16))
	if v876 < int32(0) {
		goto L251
	} else {
		goto L253
	}
L253:
	;
	v879 = F_bms_is_member(m, v876, l1)
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L32
	} else {
		goto L254
	}
L254:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v873)+20)) = uint8(v879)
	goto L251
L255:
	;
	goto L250
L256:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v897 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v887)+16))
	if v890 < int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v893 = F_bms_is_member(m, v890, l1)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L32
	} else {
		goto L259
	}
L259:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v887)+20)) = uint8(v893)
	goto L256
L260:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v907 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L261:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v897)+16))
	if v900 < int32(0) {
		goto L260
	} else {
		goto L262
	}
L262:
	;
	v903 = F_bms_is_member(m, v900, l1)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L32
	} else {
		goto L263
	}
L263:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v897)+20)) = uint8(v903)
	goto L260
L264:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v917 == int32(0) {
		goto L2
	} else {
		goto L268
	}
L265:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v907)+16))
	if v910 < int32(0) {
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v913 = F_bms_is_member(m, v910, l1)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L32
	} else {
		goto L267
	}
L267:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v907)+20)) = uint8(v913)
	goto L264
L268:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v917)+4))
	if v920 <= int32(0) {
		goto L2
	} else {
		goto L269
	}
L269:
	;
	v926 = int32(0)
	goto L270
L270:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v917)+12))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v932+v926<<(uint(int32(2))%32))))
	if v936 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	goto L2
L272:
	;
	v947 = v926 + int32(1)
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v917)+4))
	if v947 < v948 {
		v926 = v947
		goto L270
	} else {
		goto L276
	}
L273:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v936)+16))
	if v939 < int32(0) {
		goto L272
	} else {
		goto L274
	}
L274:
	;
	v942 = F_bms_is_member(m, v939, l1)
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L32
	} else {
		goto L275
	}
L275:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v936)+20)) = uint8(v942)
	goto L272
L276:
	;
	goto L271
L277:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(v950)+16))
	if int32(0) <= v953 {
		v970 = v950
		v971 = v953
		goto L3
	} else {
		goto L278
	}
L278:
	;
	goto L2
L279:
	;
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v956)+16))
	if int32(0) <= v959 {
		v970 = v956
		v971 = v959
		goto L3
	} else {
		goto L280
	}
L280:
	;
	goto L2
L281:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v962)+16))
	if v965 < int32(0) {
		goto L2
	} else {
		goto L282
	}
L282:
	;
	v970 = v962
	v971 = v965
	goto L3
L283:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v970)+20)) = uint8(v976)
	goto L2
L284:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v996
	F_errmsg_internal(m, int32(484101), v11)
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L32
	} else {
		goto L285
	}
L285:
	;
	F_errfinish(m, int32(494661), int32(595), int32(300898))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L32
	} else {
		goto L286
	}
L286:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
