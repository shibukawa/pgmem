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
			F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_datatype_0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_plpgsql_build_datatype_1), v9)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_plpgsql_build_datatype_2), int32(1960), int32(_a_F_plpgsql_build_datatype_3))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v199 int64
	_ = v199
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v304 int32
	_ = v304
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v333 int64
	_ = v333
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(224)
	m.G0 = v17
	v22 = v2
	v23 = v2
	v24 = v2
	v25 = int32(-1)
	v26 = v2
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
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
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v17)+180))
	m.G0 = v17 + int32(224)
	return v363
L4:
	;
	goto L3
L5:
	;
	if v25 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v332 = int32(m.ExcTag)
	v333 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v332 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v36
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 == v36 {
		v48 = v36
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v122 = v22
	v123 = v23
	v124 = v24
	v126 = v26
	v127 = v27
	v128 = v28
	v129 = v29
	v130 = v30
	v131 = v31
	v132 = v32
	v133 = v33
	goto L10
L10:
	;
	if v133 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v22
	v58 = l0 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v58
	v61 = v48 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v61)
	F_SPI_connect_ext(m, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v42 != int32(214) {
		v48 = v36
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	v48 = v45 ^ int32(1)
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v58
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v61)
	v76 = F_plpgsql_compile(m, l0, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v78 + int64(1)
	v83 = v76 + int32(24)
	v85 = v76 + int32(528)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v76)+528))
	v87 = int32(0)
	if v61 == v87 {
		v109 = v32
		v110 = v87
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0]))
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1]))
	goto L20
L17:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+524)))
	if v92 != int32(1) {
		v109 = v32
		v110 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v58
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v61)
	v107 = F_ResourceOwnerCreate(m, int32(0), int32(_a_F_plpgsql_call_handler_0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v109 = v107
	v110 = v107
	goto L16
L20:
	;
	v116 = v17 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v17 + int32(12)
	goto L23
L21:
	;
	v122 = v76
	v123 = v110
	v124 = v58
	v126 = v114
	v127 = v112
	v128 = v83
	v129 = v86
	v130 = v85
	v131 = v48
	v132 = v109
	v133 = v87
	goto L10
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0])) = v127
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v126
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v128)))
	*(*int64)(unsafe.Add(mBase, uint32(v128))) = v199 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v129
	if v123 != 0 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v17 + int32(16)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v138 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v182 = int32(1)
	v183 = v131 & v182
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v183)
	v185 = int32(0)
	v191 = F_plpgsql_exec_function(m, v122, l0, v185, v185, v123, (v131^int32(-1))&v182)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L32
	}
L27:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	switch v141 - int32(441) {
	case 0:
		goto L28
	case 1:
		goto L29
	default:
		goto L26
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v169 = v131 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v169)
	F_plpgsql_exec_event_trigger(m, v122, v138)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v154 = v131 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v154)
	v156 = F_plpgsql_exec_trigger(m, v122, v138)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v156
	goto L24
L31:
	;
	goto L24
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+180)) = v191
	goto L24
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v214 = v131 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v214)
	F_ReleaseAllPlanCacheRefsInOwner(m, v123)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v133 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v214)
	F_ResourceOwnerDelete(m, v123)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v241 = v131 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v241)
	F_pg_re_throw(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0])) = v127
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	v259 = v131 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v259)
	v261 = F_SPI_finish(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L7
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	if v261 == int32(2) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v259)
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_call_handler_1))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v259)
	v289 = F_SPI_result_code_string(m, v261)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v259)
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v289
	F_errmsg_internal(m, int32(_a_F_plpgsql_call_handler_2), v17)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+188)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v17)+184)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v17)+196)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v17)+200)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v17)+204)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v17)+212)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v17)+220)) = v124
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)) = uint8(v259)
	F_errfinish(m, int32(_a_F_plpgsql_call_handler_3), int32(302), int32(_a_F_plpgsql_call_handler_4))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	v337 = int32(v333)
	m.G0 = v17
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	if v17+int32(12) == v343 {
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
	if v347 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v347 = v345
	goto L53
L52:
	;
	v347 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v17)+220))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+219)))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v17)+212))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v17)+208))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v17)+204))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v17)+200))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v17)+196))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v17)+192))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v17)+188))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v17)+184))
	v22 = v350
	v23 = v355
	v24 = v348
	v25 = v347
	v26 = v356
	v27 = v357
	v28 = v353
	v29 = v352
	v30 = v351
	v31 = v349
	v32 = v354
	v33 = v339
	goto L2
L55:
	;
	goto L56
L56:
	;
	F___wasm_longjmp(m, v340, v339)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
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
	var v6 int32
	_ = v6
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
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
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
	var v169 int32
	_ = v169
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
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v482 int32
	_ = v482
	var v493 int32
	_ = v493
	var v516 int32
	_ = v516
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v934 int32
	_ = v934
	var v940 int32
	_ = v940
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v993 int32
	_ = v993
	var v997 int32
	_ = v997
	var v1004 int32
	_ = v1004
	var v1008 int32
	_ = v1008
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1172 int32
	_ = v1172
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1177 int32
	_ = v1177
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1254 int32
	_ = v1254
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1309 int32
	_ = v1309
	var v1314 int32
	_ = v1314
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1361 int32
	_ = v1361
	var v1382 int32
	_ = v1382
	var v1388 int32
	_ = v1388
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1495 int32
	_ = v1495
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1535 int32
	_ = v1535
	var v1553 int32
	_ = v1553
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1579 int32
	_ = v1579
	var v1586 int32
	_ = v1586
	var v1599 int32
	_ = v1599
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1620 int32
	_ = v1620
	var v1626 int32
	_ = v1626
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1642 int32
	_ = v1642
	var v1647 int32
	_ = v1647
	var v1651 int32
	_ = v1651
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1668 int32
	_ = v1668
	var v1673 int32
	_ = v1673
	var v1697 int32
	_ = v1697
	v5 = l4
	v6 = int32(0)
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0])) = v52
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[1])) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v22)+412)) = v47
	if v5 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v45
	goto L17
L16:
	;
	v61 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+408)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v22)+400)) = int32(_a_F_plpgsql_compile_callback_0)
	v65 = int32(_a_F_plpgsql_compile_callback_1)
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3])) = v22 + int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+396)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v22)+404)) = v22 + int32(408)
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4]))
	v81 = F_AllocSetContextCreateInternal(m, v76, int32(_a_F_plpgsql_compile_callback_2), int32(0), int32(_a_F_plpgsql_compile_callback_3), int32(_a_F_plpgsql_compile_callback_4))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v84 = int32(_a_F_plpgsql_compile_callback_5)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5])) = v85
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v81
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = F_format_procedure(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v81)+36)) = v91
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v98
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[6]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+488)) = v104
	v107 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[7])))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+492)) = uint8(v107)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[8]))
	if v5 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v112 = v110
	goto L22
L21:
	;
	v112 = int32(0)
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+496)) = v112
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[9]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v40
	if v5 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v118 = v115
	goto L25
L24:
	;
	v118 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+500)) = v118
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+96)))
	v121 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+524)) = uint16(v121)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+520)) = v121
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)) = uint8(v120)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[10])) = v121
	goto L26
L26:
	;
	F_plpgsql_ns_push(m, v51, int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = int32(128)
	v136 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[12])) = uint8(v136)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v136
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	v144 = F_MemoryContextAlloc(m, v142, int32(512))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[14])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v144
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	switch v151 {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L58
	default:
		goto L38
	}
L29:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(1960), int32(_a_F_plpgsql_compile_callback_7))
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L10
	} else {
		goto L374
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L10
	} else {
		goto L371
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L10
	} else {
		goto L369
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1635 = m.ExcPending
	if v1635 != 0 {
		goto L10
	} else {
		goto L365
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L10
	} else {
		goto L362
	}
L34:
	;
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+64)) = uint8(base.B2i32(v1254 != int32(118)))
	v1260 = F_SearchSysCache1(m, int32(82), int32(16))
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L10
	} else {
		goto L290
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v1147
	v1149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)) = uint8(v1149)
	v1152 = F_SearchSysCache1(m, int32(82), v1147)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L10
	} else {
		goto L263
	}
L36:
	;
	v1147 = int32(23)
	goto L35
L37:
	;
	v1147 = int32(3904)
	goto L35
L38:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L10
	} else {
		goto L260
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L10
	} else {
		goto L258
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L10
	} else {
		goto L256
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1096 = m.ExcPending
	if v1096 != 0 {
		goto L10
	} else {
		goto L252
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1085 = m.ExcPending
	if v1085 != 0 {
		goto L10
	} else {
		goto L250
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L10
	} else {
		goto L248
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L10
	} else {
		goto L246
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L10
	} else {
		goto L244
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L10
	} else {
		goto L242
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L10
	} else {
		goto L240
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1019 = m.ExcPending
	if v1019 != 0 {
		goto L10
	} else {
		goto L238
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1008 = m.ExcPending
	if v1008 != 0 {
		goto L10
	} else {
		goto L236
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L10
	} else {
		goto L234
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L10
	} else {
		goto L232
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L10
	} else {
		goto L227
	}
L53:
	;
	switch v493 - int32(_a_F_plpgsql_compile_callback_9) {
	case 0:
		v1147 = v516
		goto L35
	default:
		goto L36
	case 2:
		goto L37
	}
L54:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L10
	} else {
		goto L222
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L10
	} else {
		goto L220
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(2278)
	v884 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	if v884 != 0 {
		goto L41
	} else {
		goto L209
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	v553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+104)))
	if v553 != 0 {
		goto L52
	} else {
		goto L144
	}
L58:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v154
	v162 = F_get_func_arg_info(m, l1, v22+int32(392), v22+int32(388), v22+int32(384))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v22)+392))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v22)+384))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+24))
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0]))
	F_cfunc_resolve_polymorphic_argtypes(m, v162, v164, v165, v167, v5, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v173 = v162 << (uint(int32(2)) % 32)
	v174 = F_palloc(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v176 = F_palloc(m, v173)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v81
	if v162 <= int32(0) {
		v482 = v6
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v26)+108))
	if v493 <= int32(3830) {
		goto L125
	} else {
		goto L126
	}
L64:
	;
	v182 = int32(0)
	v186 = v182
	v191 = v182
	v192 = v6
	goto L65
L65:
	;
	v204 = v186 << (uint(int32(2)) % 32)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v22)+392))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v204+v205)))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v22)+384))
	if v208 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v290 <= int32(1) {
		goto L99
	} else {
		goto L100
	}
L67:
	;
	v210 = int32(*(*int8)(unsafe.Add(mBase, uint32(v186+v208))))
	v212 = v210
	goto L69
L68:
	;
	v212 = int32(105)
	goto L69
L69:
	;
	v214 = v186 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v214
	v222 = F_pg_snprintf(m, v22+int32(352), int32(32), int32(_a_F_plpgsql_compile_callback_10), v22+int32(144))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L10
	} else {
		goto L70
	}
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v226 = F_SearchSysCache1(m, int32(82), v207)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	if v226 == int32(0) {
		goto L55
	} else {
		goto L72
	}
L72:
	;
	v232 = F_build_datatype(m, v226, int32(-1), v224, int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	F_ReleaseCatCache(m, v226)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	if v236 == int32(2) {
		goto L54
	} else {
		goto L75
	}
L75:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v22)+388))
	if v239 != 0 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v247 = int32(0)
	v249 = F_plpgsql_build_variable(m, v246, v247, v232, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L81
	}
L77:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v239+v204)))
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241))))
	if v242 != 0 {
		v246 = v241
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v246 = v22 + int32(352)
	goto L76
L80:
	;
	goto L79
L81:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)))
	v253 = v212 - int32(98)
	v260 = int32(0)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v253))|base.B2i32(int32(1)<<(uint(v253)%32)&int32(_a_F_plpgsql_compile_callback_11) == v260) == v260 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v174+v191<<(uint(int32(2))%32)))) = v268
	v272 = v191 + int32(1)
	goto L84
L83:
	;
	v272 = v191
	goto L84
L84:
	;
	v277 = int32(0)
	if base.B2i32(int32(1)<<(uint(v253)%32)&int32(_a_F_plpgsql_compile_callback_12) == v277)|base.B2i32(base.Ui32(int32(18)) < base.Ui32(v253)) == v277 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176+v192<<(uint(int32(2))%32)))) = v249
	v290 = v192 + int32(1)
	goto L87
L86:
	;
	v290 = v192
	goto L87
L87:
	;
	if v251 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v293 = int32(2)
	goto L90
L89:
	;
	v293 = int32(1)
	goto L90
L90:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	F_add_parameter_name(m, v293, v294, v22+int32(352))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v22)+388))
	if v299 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v214 != v162 {
		v186 = v214
		v191 = v272
		v192 = v290
		goto L65
	} else {
		goto L96
	}
L93:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v299+v204)))
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v304 == int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v249)+4))
	F_add_parameter_name(m, v293, v307, v303)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L66
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = v460
	v482 = v462
	goto L63
L98:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+4))
	v460 = v453
	v462 = v316
	goto L97
L99:
	;
	if v290 != int32(1) {
		v482 = v290
		goto L63
	} else {
		goto L102
	}
L100:
	;
	v320 = v290
	goto L101
L101:
	;
	v322 = F_palloc0(m, int32(40))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L10
	} else {
		goto L104
	}
L102:
	;
	v316 = int32(1)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)))
	if v317 != int32(112) {
		goto L98
	} else {
		goto L103
	}
L103:
	;
	v320 = v316
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v322)+8)) = int32(_a_F_plpgsql_compile_callback_13)
	*(*int32)(unsafe.Add(mBase, uint32(v322))) = int32(1)
	v330 = F_CreateTemplateTupleDesc(m, v320)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+28)) = v320
	*(*int32)(unsafe.Add(mBase, uint32(v322)+24)) = v330
	v335 = v320 << (uint(int32(2)) % 32)
	v336 = F_palloc(m, v335)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+32)) = v336
	v339 = F_palloc(m, v335)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+36)) = v339
	v349 = int32(0)
	goto L108
L108:
	;
	v363 = v349 << (uint(int32(2)) % 32)
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v176+v363)))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	switch v366 {
	case 0, 4:
		goto L111
	default:
		goto L112
	case 2:
		goto L113
	}
L109:
	;
	v423 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v425 == v427 {
		goto L120
	} else {
		goto L121
	}
L110:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v322)+32))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v365)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v397+v363))) = v399
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v322)+36))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v365)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v401+v363))) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v322)+24))
	v407 = v349 + int32(1)
	v408 = base.I32_extend16_s(v407)
	F_TupleDescInitEntry(m, v405, v408, v399, v396, v395, int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L10
	} else {
		goto L117
	}
L111:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v365)+24))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v387)+16))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v387)+24))
	v392 = v387 + int32(4)
	v394 = v390
	v395 = v391
	goto L110
L112:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L10
	} else {
		goto L114
	}
L113:
	;
	v392 = v365 + int32(28)
	v394 = int32(0)
	v395 = int32(-1)
	goto L110
L114:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v365)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v375
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_14), v22+int32(48))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(1879), int32(_a_F_plpgsql_compile_callback_15))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L10
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
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v322)+24))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v412)))
	*(*int32)(unsafe.Add(mBase, uint32(v412+v413<<(uint(int32(4))%32)+v408*int32(100))+16)) = v394
	goto L118
L118:
	;
	if v407 != v320 {
		v349 = v407
		goto L108
	} else {
		goto L119
	}
L119:
	;
	goto L109
L120:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v425 << (uint(int32(1)) % 32)
	v436 = F_repalloc(m, v423, v425<<(uint(int32(3))%32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L10
	} else {
		goto L123
	}
L121:
	;
	v441 = v425
	v442 = v423
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v322)+4)) = v441
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v441 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v442+v441<<(uint(int32(2))%32)))) = v322
	v460 = v441
	v462 = v320
	goto L97
L123:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v436
	v440 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v441 = v440
	v442 = v436
	goto L122
L124:
	;
	if v5 != 0 {
		goto L131
	} else {
		goto L132
	}
L125:
	;
	switch v493 - int32(2277) {
	case 0, 6:
		goto L124
	case 1, 2, 3, 4, 5:
		v1147 = v493
		goto L35
	default:
		goto L128
	}
L126:
	;
	goto L127
L127:
	;
	if base.B2i32(base.Ui32(v493-int32(_a_F_plpgsql_compile_callback_16)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v493-int32(_a_F_plpgsql_compile_callback_17)) < base.Ui32(int32(2)))|base.B2i32(v493 == int32(3831)) != 0 {
		goto L124
	} else {
		goto L130
	}
L128:
	;
	if base.B2i32(v493 == int32(2776))|base.B2i32(v493 == int32(3500)) != 0 {
		goto L124
	} else {
		goto L129
	}
L129:
	;
	v1147 = v493
	goto L35
L130:
	;
	v1147 = v493
	goto L35
L131:
	;
	v516 = int32(1007)
	if int32(_a_F_plpgsql_compile_callback_16) < v493 {
		goto L53
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v527 = F_get_fn_expr_rettype(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L10
	} else {
		goto L138
	}
L134:
	;
	if v493 == int32(2277) {
		v1147 = v516
		goto L35
	} else {
		goto L135
	}
L135:
	;
	if v493 == int32(3831) {
		goto L37
	} else {
		goto L136
	}
L136:
	;
	if v493 != int32(_a_F_plpgsql_compile_callback_17) {
		goto L36
	} else {
		goto L137
	}
L137:
	;
	v1147 = int32(_a_F_plpgsql_compile_callback_18)
	goto L35
L138:
	;
	if v527 != 0 {
		v1147 = v527
		goto L35
	} else {
		goto L139
	}
L139:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L140
	}
L140:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L10
	} else {
		goto L141
	}
L141:
	;
	v537 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v537
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_19), v22+int32(128))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L10
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(423), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L10
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	v555 = F_palloc0(m, int32(40))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L10
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(2)
	v560 = F_pstrdup(m, int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L10
	} else {
		goto L146
	}
L146:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v555)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v555)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v555)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v555)+8)) = v560
	v570 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v572 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v574 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v572 == v574 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v572 << (uint(int32(1)) % 32)
	v583 = F_repalloc(m, v570, v572<<(uint(int32(3))%32))
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L10
	} else {
		goto L150
	}
L148:
	;
	v589 = v572
	v590 = v560
	v591 = v570
	goto L149
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v555)+4)) = v589
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v589 + int32(1)
	v597 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v591+v589<<(uint(v597)%32)))) = v555
	F_plpgsql_ns_additem(m, v597, v589, v590)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L10
	} else {
		goto L151
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v583
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	v588 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v589 = v588
	v590 = v586
	v591 = v583
	goto L149
L151:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+480)) = v604
	v607 = F_palloc0(m, int32(40))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v607))) = int32(2)
	v612 = F_pstrdup(m, int32(_a_F_plpgsql_compile_callback_22))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L10
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v607)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v607)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v607)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v607)+8)) = v612
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v624 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v624 == v626 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v624 << (uint(int32(1)) % 32)
	v635 = F_repalloc(m, v622, v624<<(uint(int32(3))%32))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L10
	} else {
		goto L157
	}
L155:
	;
	v641 = v624
	v642 = v612
	v643 = v622
	goto L156
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v607)+4)) = v641
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v641 + int32(1)
	v649 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v643+v641<<(uint(v649)%32)))) = v607
	F_plpgsql_ns_additem(m, v649, v641, v642)
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L158
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v635
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v607)+8))
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v641 = v640
	v642 = v638
	v643 = v635
	goto L156
L158:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v607)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+484)) = v656
	v658 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v661 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L10
	} else {
		goto L159
	}
L159:
	;
	if v661 == int32(0) {
		goto L51
	} else {
		goto L160
	}
L160:
	;
	v667 = F_build_datatype(m, v661, int32(-1), v658, int32(0))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L10
	} else {
		goto L161
	}
L161:
	;
	F_ReleaseCatCache(m, v661)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L10
	} else {
		goto L162
	}
L162:
	;
	v674 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_23), int32(0), v667, int32(1))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+48)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = int32(4)
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v683 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L10
	} else {
		goto L164
	}
L164:
	;
	if v683 == int32(0) {
		goto L50
	} else {
		goto L165
	}
L165:
	;
	v689 = F_build_datatype(m, v683, int32(-1), v680, int32(0))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	F_ReleaseCatCache(m, v683)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L10
	} else {
		goto L167
	}
L167:
	;
	v696 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_24), int32(0), v689, int32(1))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L10
	} else {
		goto L168
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v696)+48)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v696))) = int32(4)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v705 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L10
	} else {
		goto L169
	}
L169:
	;
	if v705 == int32(0) {
		goto L49
	} else {
		goto L170
	}
L170:
	;
	v711 = F_build_datatype(m, v705, int32(-1), v702, int32(0))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	F_ReleaseCatCache(m, v705)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	v718 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_25), int32(0), v711, int32(1))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v718)+48)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v718))) = int32(4)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v727 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	if v727 == int32(0) {
		goto L48
	} else {
		goto L175
	}
L175:
	;
	v733 = F_build_datatype(m, v727, int32(-1), v724, int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L10
	} else {
		goto L176
	}
L176:
	;
	F_ReleaseCatCache(m, v727)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L10
	} else {
		goto L177
	}
L177:
	;
	v740 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_26), int32(0), v733, int32(1))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L10
	} else {
		goto L178
	}
L178:
	;
	v742 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v740)+48)) = v742
	*(*int32)(unsafe.Add(mBase, uint32(v740))) = v742
	v748 = F_SearchSysCache1(m, int32(82), int32(26))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L10
	} else {
		goto L179
	}
L179:
	;
	if v748 == int32(0) {
		goto L47
	} else {
		goto L180
	}
L180:
	;
	v753 = int32(0)
	v755 = F_build_datatype(m, v748, int32(-1), v753, v753)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L10
	} else {
		goto L181
	}
L181:
	;
	F_ReleaseCatCache(m, v748)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	v762 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_27), int32(0), v755, int32(1))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v762)+48)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v762))) = int32(4)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v771 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	if v771 == int32(0) {
		goto L46
	} else {
		goto L185
	}
L185:
	;
	v777 = F_build_datatype(m, v771, int32(-1), v768, int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_ReleaseCatCache(m, v771)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	v784 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_28), int32(0), v777, int32(1))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v784)+48)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v784))) = int32(4)
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v793 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	if v793 == int32(0) {
		goto L45
	} else {
		goto L190
	}
L190:
	;
	v799 = F_build_datatype(m, v793, int32(-1), v790, int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	F_ReleaseCatCache(m, v793)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	v806 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_29), int32(0), v799, int32(1))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v806)+48)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v806))) = int32(4)
	v812 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v815 = F_SearchSysCache1(m, int32(82), int32(19))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	if v815 == int32(0) {
		goto L44
	} else {
		goto L195
	}
L195:
	;
	v821 = F_build_datatype(m, v815, int32(-1), v812, int32(0))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	F_ReleaseCatCache(m, v815)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v828 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_30), int32(0), v821, int32(1))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+48)) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v828))) = int32(4)
	v836 = F_SearchSysCache1(m, int32(82), int32(23))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	if v836 == int32(0) {
		goto L43
	} else {
		goto L200
	}
L200:
	;
	v840 = int32(0)
	v844 = F_build_datatype(m, v836, int32(-1), v840, v840)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	F_ReleaseCatCache(m, v836)
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L10
	} else {
		goto L202
	}
L202:
	;
	v851 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_31), int32(0), v844, int32(1))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v851)+48)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v851))) = int32(4)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v860 = F_SearchSysCache1(m, int32(82), int32(1009))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	if v860 == int32(0) {
		goto L42
	} else {
		goto L205
	}
L205:
	;
	v866 = F_build_datatype(m, v860, int32(-1), v857, int32(0))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L10
	} else {
		goto L206
	}
L206:
	;
	F_ReleaseCatCache(m, v860)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	v873 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_32), int32(0), v866, int32(1))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v873)+48)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v873))) = int32(4)
	v1236 = v840
	v1241 = int32(0)
	goto L34
L209:
	;
	v885 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v888 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L10
	} else {
		goto L210
	}
L210:
	;
	if v888 == int32(0) {
		goto L40
	} else {
		goto L211
	}
L211:
	;
	v892 = int32(0)
	v895 = F_build_datatype(m, v888, int32(-1), v885, v892)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	F_ReleaseCatCache(m, v888)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L10
	} else {
		goto L213
	}
L213:
	;
	v902 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_33), int32(0), v895, int32(1))
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L10
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v902)+48)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v902))) = int32(4)
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v911 = F_SearchSysCache1(m, int32(82), int32(25))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L10
	} else {
		goto L215
	}
L215:
	;
	if v911 == int32(0) {
		goto L39
	} else {
		goto L216
	}
L216:
	;
	v917 = F_build_datatype(m, v911, int32(-1), v908, int32(0))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L10
	} else {
		goto L217
	}
L217:
	;
	F_ReleaseCatCache(m, v911)
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	v924 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_34), int32(0), v917, int32(1))
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v924)+48)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v924))) = int32(4)
	v1236 = int32(0)
	v1241 = v892
	goto L34
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v207
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(16))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L10
	} else {
		goto L221
	}
L221:
	;
	goto L29
L222:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	v948 = F_format_type_be(m, v207)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v948
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_36), v22+int32(32))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L10
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(330), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_37), int32(0))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L10
	} else {
		goto L229
	}
L229:
	;
	F_errhint(m, int32(_a_F_plpgsql_compile_callback_38), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(496), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L10
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(160))
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	goto L29
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(176))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	goto L29
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(192))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L10
	} else {
		goto L237
	}
L237:
	;
	goto L29
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(208))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L10
	} else {
		goto L239
	}
L239:
	;
	goto L29
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = int32(26)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(224))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	goto L29
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(240))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L10
	} else {
		goto L243
	}
L243:
	;
	goto L29
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(256))
	mBase = m.M
	v1059 = m.ExcPending
	if v1059 != 0 {
		goto L10
	} else {
		goto L245
	}
L245:
	;
	goto L29
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(272))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L10
	} else {
		goto L247
	}
L247:
	;
	goto L29
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = int32(23)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(288))
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	goto L29
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+304)) = int32(1009)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(304))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	goto L29
L252:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_39), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L10
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(629), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L10
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(320))
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	goto L29
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(336))
	mBase = m.M
	v1130 = m.ExcPending
	if v1130 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	goto L29
L260:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v1135
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_40), v22)
	mBase = m.M
	v1139 = m.ExcPending
	if v1139 != 0 {
		goto L10
	} else {
		goto L261
	}
L261:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(657), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1144 = m.ExcPending
	if v1144 != 0 {
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
	if v1152 == int32(0) {
		goto L33
	} else {
		goto L264
	}
L264:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1152)+16))
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1158)+22)))
	v1160 = v1158 + v1159
	v1161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160)+79)))
	if base.B2i32(v1147 == int32(2249))|base.B2i32(v1161 != int32(112)) != 0 {
		goto L265
	} else {
		goto L266
	}
L265:
	;
	v1189 = F_type_is_rowtype(m, v1147)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L10
	} else {
		goto L274
	}
L266:
	;
	switch v1147 - int32(2278) {
	case 0:
		goto L265
	case 1:
		goto L32
	default:
		goto L267
	}
L267:
	;
	if v1147 == int32(3838) {
		goto L32
	} else {
		goto L268
	}
L268:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L10
	} else {
		goto L269
	}
L269:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L10
	} else {
		goto L270
	}
L270:
	;
	v1176 = F_format_type_be(m, v1147)
	mBase = m.M
	v1177 = m.ExcPending
	if v1177 != 0 {
		goto L10
	} else {
		goto L271
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v1176
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_41), v22+int32(80))
	mBase = m.M
	v1183 = m.ExcPending
	if v1183 != 0 {
		goto L10
	} else {
		goto L272
	}
L272:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(456), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L10
	} else {
		goto L273
	}
L273:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L274:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+61)) = uint8(v1189)
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+62)) = uint8(base.B2i32(v1192 == int32(100)))
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1160)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+60)) = uint8(v1196)
	v1198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1160)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v1198
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v26)+108))
	if v1200 <= int32(3830) {
		goto L278
	} else {
		goto L279
	}
L275:
	;
	F_ReleaseCatCache(m, v1152)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L10
	} else {
		goto L289
	}
L276:
	;
	if v482 != 0 {
		goto L275
	} else {
		goto L286
	}
L277:
	;
	if v1200 != int32(_a_F_plpgsql_compile_callback_42) {
		goto L275
	} else {
		goto L285
	}
L278:
	;
	switch v1200 - int32(2277) {
	case 0, 6:
		goto L276
	case 1, 2, 3, 4, 5:
		goto L277
	default:
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	if base.B2i32(v1200 == int32(3831))|base.B2i32(base.Ui32(v1200-int32(_a_F_plpgsql_compile_callback_16)) < base.Ui32(int32(4)))|base.B2i32(v1200 == int32(_a_F_plpgsql_compile_callback_17)) != 0 {
		goto L276
	} else {
		goto L284
	}
L281:
	;
	if v1200 == int32(2776) {
		goto L276
	} else {
		goto L282
	}
L282:
	;
	if v1200 != int32(3500) {
		goto L277
	} else {
		goto L283
	}
L283:
	;
	goto L276
L284:
	;
	goto L277
L285:
	;
	goto L276
L286:
	;
	v1222 = int32(0)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v1226 = F_build_datatype(m, v1152, int32(-1), v1224, v1222)
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L10
	} else {
		goto L287
	}
L287:
	;
	v1229 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_43), v1222, v1226, int32(1))
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L10
	} else {
		goto L288
	}
L288:
	;
	goto L275
L289:
	;
	v1236 = v174
	v1241 = base.B2i32(int32(0) < v482)
	goto L34
L290:
	;
	if v1260 == int32(0) {
		goto L31
	} else {
		goto L291
	}
L291:
	;
	v1265 = int32(0)
	v1267 = F_build_datatype(m, v1260, int32(-1), v1265, v1265)
	mBase = m.M
	v1268 = m.ExcPending
	if v1268 != 0 {
		goto L10
	} else {
		goto L292
	}
L292:
	;
	F_ReleaseCatCache(m, v1260)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L10
	} else {
		goto L293
	}
L293:
	;
	v1274 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_44), int32(0), v1267, int32(1))
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L10
	} else {
		goto L294
	}
L294:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v1274)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+476)) = v1276
	v1280 = F_plpgsql_yyparse(m, l3+int32(516), v47)
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L10
	} else {
		goto L295
	}
L295:
	;
	if v1280 != 0 {
		goto L30
	} else {
		goto L296
	}
L296:
	;
	F_scanner_finish(m, v47)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L10
	} else {
		goto L297
	}
L297:
	;
	F_pfree(m, v45)
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L10
	} else {
		goto L298
	}
L298:
	;
	if v1241 != 0 {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1294 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v1294
	if v1294 <= int32(0) {
		goto L305
	} else {
		goto L306
	}
L300:
	;
	F_add_dummy_return(m, l3)
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L10
	} else {
		goto L304
	}
L301:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	if v1286 == int32(2278) {
		goto L300
	} else {
		goto L302
	}
L302:
	;
	v1289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)))
	if v1289 != int32(1) {
		goto L299
	} else {
		goto L303
	}
L303:
	;
	goto L300
L304:
	;
	goto L299
L305:
	;
	v1430 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+504)) = v1430
	v1434 = F_palloc(m, v1430<<(uint(int32(2))%32))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L10
	} else {
		goto L317
	}
L306:
	;
	v1299 = l3 + int32(72)
	v1300 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1294) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1309 = v1300
	v1314 = int32(0)
	goto L310
L308:
	;
	v1361 = v1300
	goto L309
L309:
	;
	v1382 = v1361
	v1388 = v1300
	goto L314
L310:
	;
	v1327 = v1309 << (uint(int32(2)) % 32)
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1327)))
	*(*int32)(unsafe.Add(mBase, uint32(v1299+v1327))) = v1330
	v1332 = int32(4)
	v1333 = v1327 | v1332
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1333)))
	*(*int32)(unsafe.Add(mBase, uint32(v1299+v1333))) = v1336
	v1339 = v1327 | int32(8)
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1339)))
	*(*int32)(unsafe.Add(mBase, uint32(v1299+v1339))) = v1342
	v1345 = v1327 | int32(12)
	v1348 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1345)))
	*(*int32)(unsafe.Add(mBase, uint32(v1299+v1345))) = v1348
	v1351 = v1309 + v1332
	v1353 = v1314 + v1332
	if v1353 != v1294&int32(_a_F_plpgsql_compile_callback_45) {
		v1309 = v1351
		v1314 = v1353
		goto L310
	} else {
		goto L312
	}
L311:
	;
	if v1294&int32(3) == int32(0) {
		goto L305
	} else {
		goto L313
	}
L312:
	;
	goto L311
L313:
	;
	v1361 = v1351
	goto L309
L314:
	;
	v1400 = v1382 << (uint(int32(2)) % 32)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1236+v1400)))
	*(*int32)(unsafe.Add(mBase, uint32(v1299+v1400))) = v1403
	v1405 = int32(1)
	v1408 = v1388 + v1405
	if v1408 != v1294&int32(3) {
		v1382 = v1382 + v1405
		v1388 = v1408
		goto L314
	} else {
		goto L316
	}
L315:
	;
	goto L305
L316:
	;
	goto L315
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+508)) = v1434
	v1438 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	if v1438 <= int32(0) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+512)) = v1535
	v1553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+525)))
	if v1553 == int32(1) {
		goto L337
	} else {
		goto L338
	}
L319:
	;
	v1535 = int32(0)
	goto L318
L320:
	;
	goto L321
L321:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v1444 = int32(0)
	if v1438 != int32(1) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1454 = v1444
	v1455 = v1444
	v1460 = int32(0)
	goto L325
L323:
	;
	v1504 = v1444
	v1505 = v1444
	goto L324
L324:
	;
	v1523 = v1504 << (uint(int32(2)) % 32)
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1523+v1443)))
	*(*int32)(unsafe.Add(mBase, uint32(v1434+v1523))) = v1526
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1526)))
	switch v1528 {
	case 0, 4:
		goto L335
	default:
		v1535 = v1505
		goto L318
	case 2:
		goto L336
	}
L325:
	;
	v1473 = v1454 << (uint(int32(2)) % 32)
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1443)))
	*(*int32)(unsafe.Add(mBase, uint32(v1434+v1473))) = v1476
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1476)))
	switch v1478 {
	case 0, 4:
		goto L329
	default:
		v1483 = v1455
		goto L327
	case 2:
		goto L328
	}
L326:
	;
	if v1438&int32(1) == int32(0) {
		v1535 = v1495
		goto L318
	} else {
		goto L334
	}
L327:
	;
	v1485 = v1473 | int32(4)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1485+v1443)))
	*(*int32)(unsafe.Add(mBase, uint32(v1434+v1485))) = v1488
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1488)))
	switch v1490 {
	case 0, 4:
		goto L331
	default:
		v1495 = v1483
		goto L330
	case 2:
		goto L332
	}
L328:
	;
	v1483 = v1455 + int32(40)
	goto L327
L329:
	;
	v1483 = v1455 + int32(56)
	goto L327
L330:
	;
	v1496 = int32(2)
	v1497 = v1454 + v1496
	v1499 = v1460 + v1496
	if v1499 != v1438&int32(2147483646) {
		v1454 = v1497
		v1455 = v1495
		v1460 = v1499
		goto L325
	} else {
		goto L333
	}
L331:
	;
	v1495 = v1483 + int32(56)
	goto L330
L332:
	;
	v1495 = v1483 + int32(40)
	goto L330
L333:
	;
	goto L326
L334:
	;
	v1504 = v1497
	v1505 = v1495
	goto L324
L335:
	;
	v1535 = v1505 + int32(56)
	goto L318
L336:
	;
	v1535 = v1505 + int32(40)
	goto L318
L337:
	;
	F_plpgsql_mark_local_assignment_targets(m, l3)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L10
	} else {
		goto L340
	}
L338:
	;
	goto L339
L339:
	;
	v1559 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[12])))
	if v1559 == int32(1) {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	goto L339
L341:
	;
	F_plpgsql_dumptree(m, l3)
	mBase = m.M
	v1563 = m.ExcPending
	if v1563 != 0 {
		goto L10
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1565 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[16]))
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	if v1569 != v1565 {
		goto L346
	} else {
		goto L347
	}
L344:
	;
	goto L343
L345:
	;
	v1599 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0])) = v1599
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[1])) = uint8(v1599)
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(v22)+396))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3])) = v1605
	v1607 = int32(_a_F_plpgsql_compile_callback_46)
	v1608 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5])) = v1599
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v1608
	m.G0 = v22 + int32(416)
	return
L346:
	;
	if v1569 == int32(0) {
		goto L349
	} else {
		goto L350
	}
L347:
	;
	goto L348
L348:
	;
	goto L345
L349:
	;
	if v1565 != 0 {
		goto L356
	} else {
		goto L357
	}
L350:
	;
	v1573 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	if v1574 != 0 {
		goto L352
	} else {
		goto L353
	}
L351:
	;
	if v1573 == int32(0) {
		goto L349
	} else {
		goto L355
	}
L352:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1574)+28)) = v1573
	goto L351
L353:
	;
	goto L354
L354:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1569)+20)) = v1573
	goto L351
L355:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v81)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1573)+24)) = v1579
	goto L349
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v1565
	v1586 = *(*int32)(unsafe.Add(mBase, uint32(v1565)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+28)) = v1586
	if v1586 != 0 {
		goto L359
	} else {
		goto L360
	}
L357:
	;
	goto L358
L358:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v81)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = int32(0)
	goto L348
L359:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1586)+24)) = v81
	goto L361
L360:
	;
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1565)+20)) = v81
	goto L345
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v1147
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22-int32(-64))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L10
	} else {
		goto L363
	}
L363:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(438), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L10
	} else {
		goto L364
	}
L364:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L365:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1638 = m.ExcPending
	if v1638 != 0 {
		goto L10
	} else {
		goto L366
	}
L366:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_47), int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L10
	} else {
		goto L367
	}
L367:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(451), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1647 = m.ExcPending
	if v1647 != 0 {
		goto L10
	} else {
		goto L368
	}
L368:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = int32(16)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_35), v22+int32(96))
	mBase = m.M
	v1658 = m.ExcPending
	if v1658 != 0 {
		goto L10
	} else {
		goto L370
	}
L370:
	;
	goto L29
L371:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v1280
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_48), v22+int32(112))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L10
	} else {
		goto L372
	}
L372:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(680), int32(_a_F_plpgsql_compile_callback_20))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L10
	} else {
		goto L373
	}
L373:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L374:
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
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
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = F_plpgsql_scanner_init(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = int32(_a_F_plpgsql_compile_inline_0)
		v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[1])))
		*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v25)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(_a_F_plpgsql_compile_inline_1)
		v31 = int32(_a_F_plpgsql_compile_inline_2)
		v32 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3]))
		*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v14 + int32(28)
		*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v32
		*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v14 + int32(40)
		v43 = F_palloc0(m, int32(536))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[4])) = v43
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5]))
			v52 = F_AllocSetContextCreateInternal(m, v47, int32(_a_F_plpgsql_compile_inline_3), int32(0), int32(_a_F_plpgsql_compile_inline_4), int32(_a_F_plpgsql_compile_inline_5))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v55 = int32(_a_F_plpgsql_compile_inline_6)
				v56 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5]))
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v56
				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v52
				v61 = F_pstrdup(m, int32(_a_F_plpgsql_compile_inline_0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v43)+472)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+48)) = v52
					*(*int64)(unsafe.Add(mBase, uint32(v43)+40)) = int64(2)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+32)) = v61
					v70 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[7]))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+488)) = v70
					v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[8])))
					v74 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v43)+524)) = uint16(v74)
					*(*int32)(unsafe.Add(mBase, uint32(v43)+520)) = v74
					*(*int64)(unsafe.Add(mBase, uint32(v43)+496)) = int64(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v43)+492)) = uint8(v73)
					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[9])) = v74
					F_plpgsql_ns_push(m, int32(_a_F_plpgsql_compile_inline_0), int32(0))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[10])) = int32(128)
						v92 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])) = uint8(v92)
						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12])) = v92
						v98 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
						v100 = F_MemoryContextAlloc(m, v98, int32(512))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[13])) = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[14])) = v100
							*(*int64)(unsafe.Add(mBase, uint32(v43)+52)) = int64(17179871462)
							*(*int32)(unsafe.Add(mBase, uint32(v43)+60)) = int32(1)
							v111 = int32(_a_F_plpgsql_compile_inline_7)
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
											v127 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_inline_8), int32(0), v120, int32(1))
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
														F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_inline_9))
														mBase = m.M
														v290 = m.ExcPending
														if v290 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v133
															F_errmsg_internal(m, int32(_a_F_plpgsql_compile_inline_10), v14+int32(16))
															mBase = m.M
															v296 = m.ExcPending
															if v296 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_plpgsql_compile_inline_11), int32(840), int32(_a_F_plpgsql_compile_inline_12))
																mBase = m.M
																v301 = m.ExcPending
																if v301 != 0 {
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
																	v146 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+504)) = v146
																	v150 = F_palloc(m, v146<<(uint(int32(2))%32))
																	mBase = m.M
																	v151 = m.ExcPending
																	if v151 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v150
																		v154 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																		if v154 <= int32(0) {
																			v232 = v142
																		} else {
																			v158 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[14]))
																			v159 = int32(0)
																			if v154 != int32(1) {
																				v167 = v159
																				v168 = v142
																				v176 = int32(0)
																				for {
																					v178 = v167 << (uint(int32(2)) % 32)
																					v181 = *(*int32)(unsafe.Add(mBase, uint32(v178+v158)))
																					*(*int32)(unsafe.Add(mBase, uint32(v150+v178))) = v181
																					v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
																					switch v183 {
																					case 0, 4:
																						v188 = v168 + int32(56)
																					default:
																						v188 = v168
																					case 2:
																						v188 = v168 + int32(40)
																					}
																					v190 = v178 | int32(4)
																					v193 = *(*int32)(unsafe.Add(mBase, uint32(v190+v158)))
																					*(*int32)(unsafe.Add(mBase, uint32(v150+v190))) = v193
																					v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
																					switch v195 {
																					case 0, 4:
																						v200 = v188 + int32(56)
																					default:
																						v200 = v188
																					case 2:
																						v200 = v188 + int32(40)
																					}
																					v201 = int32(2)
																					v202 = v167 + v201
																					v204 = v176 + v201
																					if v204 != v154&int32(2147483646) {
																						v167 = v202
																						v168 = v200
																						v176 = v204
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v154&int32(1) == int32(0) {
																					v232 = v200
																				} else {
																					v209 = v202
																					v210 = v200
																					v220 = v209 << (uint(int32(2)) % 32)
																					v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v158)))
																					*(*int32)(unsafe.Add(mBase, uint32(v150+v220))) = v223
																					v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
																					switch v225 {
																					case 0, 4:
																						v232 = v210 + int32(56)
																					default:
																						v232 = v210
																					case 2:
																						v232 = v210 + int32(40)
																					}
																				}
																			} else {
																				v209 = v159
																				v210 = v142
																				v220 = v209 << (uint(int32(2)) % 32)
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v158)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v220))) = v223
																				v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
																				switch v225 {
																				case 0, 4:
																					v232 = v210 + int32(56)
																				default:
																					v232 = v210
																				case 2:
																					v232 = v210 + int32(40)
																				}
																			}
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v232
																		v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																		if v242 == int32(1) {
																			F_plpgsql_mark_local_assignment_targets(m, v43)
																			mBase = m.M
																			v246 = m.ExcPending
																			if v246 != 0 {
																				return int32(0)
																			} else {
																				v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																				if v248 == int32(1) {
																					F_plpgsql_dumptree(m, v43)
																					mBase = m.M
																					v252 = m.ExcPending
																					if v252 != 0 {
																						return int32(0)
																					} else {
																						v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																						v257 = int32(0)
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																						*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																						v262 = int32(_a_F_plpgsql_compile_inline_13)
																						v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																						m.G0 = v14 + int32(48)
																						return v43
																					}
																				} else {
																					v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																					v257 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																					v262 = int32(_a_F_plpgsql_compile_inline_13)
																					v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			}
																		} else {
																			v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																			if v248 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v252 = m.ExcPending
																				if v252 != 0 {
																					return int32(0)
																				} else {
																					v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																					v257 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																					v262 = int32(_a_F_plpgsql_compile_inline_13)
																					v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																				v257 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																				v262 = int32(_a_F_plpgsql_compile_inline_13)
																				v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		}
																	}
																}
															} else {
																v142 = int32(0)
																*(*int32)(unsafe.Add(mBase, uint32(v43)+68)) = v142
																v146 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																*(*int32)(unsafe.Add(mBase, uint32(v43)+504)) = v146
																v150 = F_palloc(m, v146<<(uint(int32(2))%32))
																mBase = m.M
																v151 = m.ExcPending
																if v151 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v150
																	v154 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																	if v154 <= int32(0) {
																		v232 = v142
																	} else {
																		v158 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[14]))
																		v159 = int32(0)
																		if v154 != int32(1) {
																			v167 = v159
																			v168 = v142
																			v176 = int32(0)
																			for {
																				v178 = v167 << (uint(int32(2)) % 32)
																				v181 = *(*int32)(unsafe.Add(mBase, uint32(v178+v158)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v178))) = v181
																				v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
																				switch v183 {
																				case 0, 4:
																					v188 = v168 + int32(56)
																				default:
																					v188 = v168
																				case 2:
																					v188 = v168 + int32(40)
																				}
																				v190 = v178 | int32(4)
																				v193 = *(*int32)(unsafe.Add(mBase, uint32(v190+v158)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v190))) = v193
																				v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
																				switch v195 {
																				case 0, 4:
																					v200 = v188 + int32(56)
																				default:
																					v200 = v188
																				case 2:
																					v200 = v188 + int32(40)
																				}
																				v201 = int32(2)
																				v202 = v167 + v201
																				v204 = v176 + v201
																				if v204 != v154&int32(2147483646) {
																					v167 = v202
																					v168 = v200
																					v176 = v204
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v154&int32(1) == int32(0) {
																				v232 = v200
																			} else {
																				v209 = v202
																				v210 = v200
																				v220 = v209 << (uint(int32(2)) % 32)
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v158)))
																				*(*int32)(unsafe.Add(mBase, uint32(v150+v220))) = v223
																				v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
																				switch v225 {
																				case 0, 4:
																					v232 = v210 + int32(56)
																				default:
																					v232 = v210
																				case 2:
																					v232 = v210 + int32(40)
																				}
																			}
																		} else {
																			v209 = v159
																			v210 = v142
																			v220 = v209 << (uint(int32(2)) % 32)
																			v223 = *(*int32)(unsafe.Add(mBase, uint32(v220+v158)))
																			*(*int32)(unsafe.Add(mBase, uint32(v150+v220))) = v223
																			v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
																			switch v225 {
																			case 0, 4:
																				v232 = v210 + int32(56)
																			default:
																				v232 = v210
																			case 2:
																				v232 = v210 + int32(40)
																			}
																		}
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v232
																	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																	if v242 == int32(1) {
																		F_plpgsql_mark_local_assignment_targets(m, v43)
																		mBase = m.M
																		v246 = m.ExcPending
																		if v246 != 0 {
																			return int32(0)
																		} else {
																			v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																			if v248 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v252 = m.ExcPending
																				if v252 != 0 {
																					return int32(0)
																				} else {
																					v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																					v257 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																					v262 = int32(_a_F_plpgsql_compile_inline_13)
																					v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																				v257 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																				v262 = int32(_a_F_plpgsql_compile_inline_13)
																				v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		}
																	} else {
																		v248 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																		if v248 == int32(1) {
																			F_plpgsql_dumptree(m, v43)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int32(0)
																			} else {
																				v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																				v257 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																				v262 = int32(_a_F_plpgsql_compile_inline_13)
																				v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		} else {
																			v254 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v254
																			v257 = int32(0)
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v257
																			*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v257)
																			v262 = int32(_a_F_plpgsql_compile_inline_13)
																			v263 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v257
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v263
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
									F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_inline_9))
									mBase = m.M
									v276 = m.ExcPending
									if v276 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(16)
										F_errmsg_internal(m, int32(_a_F_plpgsql_compile_inline_14), v14)
										mBase = m.M
										v281 = m.ExcPending
										if v281 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_plpgsql_compile_inline_11), int32(1960), int32(_a_F_plpgsql_compile_inline_15))
											mBase = m.M
											v286 = m.ExcPending
											if v286 != 0 {
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v10 {
	case 0, 4:
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v69 = v63 + int32(4)
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
		m.G0 = v8 + int32(32)
		return v70
	default:
		F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_0))
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v53
			F_errmsg_internal(m, int32(_a_F_plpgsql_exec_get_datum_type_1), v8)
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_2), int32(_a_F_plpgsql_exec_get_datum_type_3), int32(_a_F_plpgsql_exec_get_datum_type_4))
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
	case 2:
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
		if v11 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
			if v12 == int32(2249) {
				v69 = v11 + int32(36)
			} else {
				v69 = l1 + int32(28)
			}
		} else {
			v69 = l1 + int32(28)
		}
		v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
		m.G0 = v8 + int32(32)
		return v70
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
							F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50360452))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
									v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v83
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v82
									F_errmsg(m, int32(_a_F_plpgsql_exec_get_datum_type_5), v8+int32(16))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_2), int32(_a_F_plpgsql_exec_get_datum_type_6), int32(_a_F_plpgsql_exec_get_datum_type_4))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
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
							v69 = l1 + int32(36)
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
							m.G0 = v8 + int32(32)
							return v70
						}
					}
				} else {
					v69 = l1 + int32(36)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
					m.G0 = v8 + int32(32)
					return v70
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
						F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_exec_get_datum_type_0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50360452))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v83
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v82
								F_errmsg(m, int32(_a_F_plpgsql_exec_get_datum_type_5), v8+int32(16))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_plpgsql_exec_get_datum_type_2), int32(_a_F_plpgsql_exec_get_datum_type_6), int32(_a_F_plpgsql_exec_get_datum_type_4))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
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
						v69 = l1 + int32(36)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
						m.G0 = v8 + int32(32)
						return v70
					}
				}
			} else {
				v69 = l1 + int32(36)
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
				m.G0 = v8 + int32(32)
				return v70
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_pop[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_ns_pop[0])) = v10
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
		if v30 == int32(_a_F_plpgsql_param_compile_0) {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			if v18 != v33 {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_1)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
				if v35 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_1)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_2)
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_3)
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_4)
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_5)
	case 3:
		*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_6)
	case 4:
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
		v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+12)))
		if v47 == int32(_a_F_plpgsql_param_compile_0) {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_5)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(_a_F_plpgsql_param_compile_4)
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v122 int32
	_ = v122
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	v17 = v11 + int32(16)
	v20 = F_pg_snprintf(m, v17, int32(32), int32(_a_F_plpgsql_param_ref_0), v11)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v25 == v24 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	if v164 != 0 {
		goto L39
	} else {
		goto L40
	}
L4:
	;
	goto L3
L5:
	;
	v164 = v51
	goto L4
L7:
	;
	v164 = int32(0)
	goto L4
L8:
	;
	v37 = v25
	goto L9
L9:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L7
L11:
	;
	v51 = v37
	v53 = v46
	goto L14
L12:
	;
	v76 = v37
	goto L13
L13:
	;
	goto L21
L14:
	;
	v58 = F_strcmp(m, v51+int32(12), v17)
	mBase = m.M
	if v58|base.B2i32(v53 == int32(1))&int32(0) == int32(0) {
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
	goto L5
L17:
	;
	goto L18
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	if v71 != 0 {
		v51 = v70
		v53 = v71
		goto L14
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v122 != 0 {
		v37 = v122
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L10
L39:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+528))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+68))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v177+v178<<(uint(int32(2))%32))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v184 = int32(_a_F_plpgsql_param_ref_1)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_param_ref[0]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v175)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_param_ref[0])) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v190 = F_bms_add_member(m, v189, v178)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	v212 = v24
	goto L41
L41:
	;
	m.G0 = v11 + int32(48)
	return v212
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v190
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_param_ref[0])) = v185
	v196 = F_palloc0(m, int32(28))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196)+8)) = v178 + int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v196))) = int64(8)
	F_plpgsql_exec_get_datum_type_info(m, v176, v182, v196+int32(12), v196+int32(16), v196+int32(20))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196)+24)) = v183
	v212 = v196
	goto L41
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a_F_plpgsql_parse_err_condition_0)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_parse_err_condition[0])))
	if base.B2i32(v14 == v2)|base.B2i32(v14 != v17) != 0 {
		v35 = v14
		v36 = v17
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v117
L2:
	;
	if v35-v36 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	goto L2
L4:
	;
	v20 = l0
	v21 = v11
	goto L5
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v25 == int32(0) {
		v35 = v25
		v36 = v24
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v35 = v25
	v36 = v24
	goto L3
L7:
	;
	v28 = int32(1)
	if v25 == v24 {
		v20 = v20 + v28
		v21 = v21 + v28
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v39 = v2
	v42 = v2
	goto L12
L10:
	;
	goto L11
L11:
	;
	v109 = F_palloc(m, int32(12))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L24
	} else {
		goto L32
	}
L12:
	;
	v45 = v42 << (uint(int32(3)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_plpgsql_parse_err_condition[1])))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	if base.B2i32(v51 == int32(0))|base.B2i32(v51 != v54) != 0 {
		v72 = v51
		v73 = v54
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if v86 != 0 {
		v117 = v86
		goto L1
	} else {
		goto L27
	}
L14:
	;
	if v72-v73 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	goto L14
L16:
	;
	v57 = l0
	v58 = v48
	goto L17
L17:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+1)))
	if v62 == int32(0) {
		v72 = v62
		v73 = v61
		goto L15
	} else {
		goto L19
	}
L18:
	;
	v72 = v62
	v73 = v61
	goto L15
L19:
	;
	v65 = int32(1)
	if v62 == v61 {
		v57 = v57 + v65
		v58 = v58 + v65
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v78 = F_palloc(m, int32(12))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v86 = v39
	goto L23
L23:
	;
	v89 = v42 + int32(1)
	if v89 != int32(251) {
		v39 = v86
		v42 = v89
		goto L12
	} else {
		goto L26
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = l0
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_c_F_plpgsql_parse_err_condition[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v84
	v86 = v78
	goto L23
L26:
	;
	goto L13
L27:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_parse_err_condition_1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L24
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
	F_errmsg(m, int32(_a_F_plpgsql_parse_err_condition_2), v9)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_plpgsql_parse_err_condition_3), int32(2190), int32(_a_F_plpgsql_parse_err_condition_4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(-1)
	v117 = v109
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
			F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_peek_0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_plpgsql_peek_1), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_plpgsql_peek_2), int32(388), int32(_a_F_plpgsql_peek_3))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
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
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+72))
			v40 = v16 + v37*int32(24)
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+108)) = v41
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+100)) = v43
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+92)) = v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+72))
			*(*int32)(unsafe.Add(mBase, uint32(v47)+72)) = v48 + int32(1)
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
	var v29 int32
	_ = v29
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
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+488))
	if v14 == int32(2) {
		v20 = l2
	} else {
		v20 = v4
	}
	if base.B2i32(v14 == int32(1))|v20 != 0 {
		v59 = v4
		m.G0 = v10 + int32(16)
		return v59
	} else {
		v23 = base.B2i32(l2 == int32(0))
		v26 = F_resolve_column_ref(m, l0, v12, l1, v23)
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			if v23|base.B2i32(v26 == int32(0)) != 0 {
				v59 = v26
				m.G0 = v10 + int32(16)
				return v59
			} else {
				F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_post_column_ref_0))
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
							F_errmsg(m, int32(_a_F_plpgsql_post_column_ref_1), v10)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errdetail(m, int32(_a_F_plpgsql_post_column_ref_2), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
									F_parser_errposition(m, l0, v51)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_plpgsql_post_column_ref_3), int32(1046), int32(_a_F_plpgsql_post_column_ref_4))
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
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
	var v167 int32
	_ = v167
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
	var v278 int32
	_ = v278
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
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v884 int32
	_ = v884
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v13 {
	case 0:
		goto L28
	case 1:
		goto L27
	case 2:
		goto L26
	case 3:
		goto L25
	case 4:
		goto L24
	case 5:
		goto L23
	case 6:
		goto L22
	case 7:
		goto L21
	case 8:
		goto L20
	case 9:
		goto L19
	case 10:
		goto L18
	case 11:
		goto L17
	case 12:
		goto L16
	case 13:
		goto L15
	case 14:
		goto L14
	case 15:
		goto L13
	case 16:
		goto L12
	case 17:
		goto L11
	case 18:
		goto L10
	case 19, 22, 25, 26:
		goto L1
	case 20:
		goto L9
	case 21:
		goto L8
	case 23:
		goto L7
	case 24:
		goto L6
	default:
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return
L2:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v958 == int32(0) {
		goto L1
	} else {
		goto L279
	}
L3:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v844)+16))
	if v951 < int32(0) {
		goto L2
	} else {
		goto L277
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2_0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L34
	} else {
		goto L274
	}
L5:
	;
	v934 = F_bms_is_member(m, v926, l1)
	mBase = m.M
	v935 = m.ExcPending
	if v935 != 0 {
		goto L34
	} else {
		goto L273
	}
L6:
	;
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v920 == int32(0) {
		goto L1
	} else {
		goto L271
	}
L7:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v914 == int32(0) {
		goto L1
	} else {
		goto L269
	}
L8:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v908 == int32(0) {
		goto L1
	} else {
		goto L267
	}
L9:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v845 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L10:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v811 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L11:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v768 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L12:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v762 == int32(0) {
		goto L1
	} else {
		goto L223
	}
L13:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v746 == int32(0) {
		goto L217
	} else {
		goto L218
	}
L14:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v672 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L15:
	;
	v619 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v619 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L16:
	;
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v613 == int32(0) {
		goto L1
	} else {
		goto L179
	}
L17:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v607 == int32(0) {
		goto L1
	} else {
		goto L177
	}
L18:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v601 == int32(0) {
		goto L1
	} else {
		goto L175
	}
L19:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v565 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L20:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v526 == int32(0) {
		goto L156
	} else {
		goto L157
	}
L21:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v487 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L22:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v431 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L23:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v395 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L24:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v370 == int32(0) {
		goto L1
	} else {
		goto L113
	}
L25:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v259 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L26:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v114 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v108 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L28:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v14 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v47 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L30:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v22 = v3
	goto L32
L32:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v22<<(uint(int32(2))%32))))
	F_mark_stmt(m, v32, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L29
L34:
	;
	return
L35:
	;
	v36 = v22 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v36 < v37 {
		v22 = v36
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	if v50 == int32(0) {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 <= int32(0) {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v57 = int32(0)
	goto L40
L40:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v57<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v70 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L1
L42:
	;
	v105 = v57 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v105 < v106 {
		v57 = v105
		goto L40
	} else {
		goto L49
	}
L43:
	;
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v74 <= v73 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v79 = v73
	goto L45
L45:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v85+v79<<(uint(int32(2))%32))))
	F_mark_stmt(m, v89, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L34
	} else {
		goto L47
	}
L46:
	;
	goto L42
L47:
	;
	v93 = v79 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v93 < v94 {
		v79 = v93
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L41
L50:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if int32(0) <= v111 {
		v926 = v111
		v928 = v108
		goto L5
	} else {
		goto L51
	}
L51:
	;
	goto L1
L52:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v124 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v117 < int32(0) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v120 = F_bms_is_member(m, v117, l1)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L34
	} else {
		goto L55
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v114)+20)) = uint8(v120)
	goto L52
L56:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v158 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v127 <= int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v133 = int32(0)
	goto L59
L59:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139+v133<<(uint(int32(2))%32))))
	F_mark_stmt(m, v143, l1)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L34
	} else {
		goto L61
	}
L60:
	;
	goto L56
L61:
	;
	v147 = v133 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v147 < v148 {
		v133 = v147
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v233 == int32(0) {
		goto L1
	} else {
		goto L80
	}
L64:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v161 <= int32(0) {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v167 = v3
	goto L66
L66:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v167<<(uint(int32(2))%32))))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v177 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	goto L63
L68:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v176)+8))
	if v187 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	if v180 < int32(0) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v183 = F_bms_is_member(m, v180, l1)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L34
	} else {
		goto L71
	}
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+20)) = uint8(v183)
	goto L68
L72:
	;
	v222 = v167 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v222 < v223 {
		v167 = v222
		goto L66
	} else {
		goto L79
	}
L73:
	;
	v190 = int32(0)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v191 <= v190 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v196 = v190
	goto L75
L75:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v196<<(uint(int32(2))%32))))
	F_mark_stmt(m, v206, l1)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L34
	} else {
		goto L77
	}
L76:
	;
	goto L72
L77:
	;
	v210 = v196 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	if v210 < v211 {
		v196 = v210
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	goto L67
L80:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v236 <= int32(0) {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v242 = int32(0)
	goto L82
L82:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v248+v242<<(uint(int32(2))%32))))
	F_mark_stmt(m, v252, l1)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L34
	} else {
		goto L84
	}
L83:
	;
	goto L1
L84:
	;
	v256 = v242 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v256 < v257 {
		v242 = v256
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v269 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259)+16))
	if v262 < int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v265 = F_bms_is_member(m, v262, l1)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L34
	} else {
		goto L89
	}
L89:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+20)) = uint8(v265)
	goto L86
L90:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v344 == int32(0) {
		goto L1
	} else {
		goto L107
	}
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v272 <= int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v278 = v3
	goto L93
L93:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283+v278<<(uint(int32(2))%32))))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v288 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	goto L90
L95:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v287)+8))
	if v298 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)+16))
	if v291 < int32(0) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v294 = F_bms_is_member(m, v291, l1)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L34
	} else {
		goto L98
	}
L98:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v288)+20)) = uint8(v294)
	goto L95
L99:
	;
	v333 = v278 + int32(1)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	if v333 < v334 {
		v278 = v333
		goto L93
	} else {
		goto L106
	}
L100:
	;
	v301 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v302 <= v301 {
		goto L99
	} else {
		goto L101
	}
L101:
	;
	v307 = v301
	goto L102
L102:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v298)+12))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313+v307<<(uint(int32(2))%32))))
	F_mark_stmt(m, v317, l1)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L34
	} else {
		goto L104
	}
L103:
	;
	goto L99
L104:
	;
	v321 = v307 + int32(1)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v298)+4))
	if v321 < v322 {
		v307 = v321
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	goto L94
L107:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v347 <= int32(0) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v353 = int32(0)
	goto L109
L109:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v344)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v359+v353<<(uint(int32(2))%32))))
	F_mark_stmt(m, v363, l1)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L34
	} else {
		goto L111
	}
L110:
	;
	goto L1
L111:
	;
	v367 = v353 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v367 < v368 {
		v353 = v367
		goto L109
	} else {
		goto L112
	}
L112:
	;
	goto L110
L113:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v373 <= int32(0) {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v378 = v3
	goto L115
L115:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v378<<(uint(int32(2))%32))))
	F_mark_stmt(m, v388, l1)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L34
	} else {
		goto L117
	}
L116:
	;
	goto L1
L117:
	;
	v392 = v378 + int32(1)
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v392 < v393 {
		v378 = v392
		goto L115
	} else {
		goto L118
	}
L118:
	;
	goto L116
L119:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v405 == int32(0) {
		goto L1
	} else {
		goto L123
	}
L120:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	if v398 < int32(0) {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v401 = F_bms_is_member(m, v398, l1)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L34
	} else {
		goto L122
	}
L122:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+20)) = uint8(v401)
	goto L119
L123:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v408 <= int32(0) {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v414 = int32(0)
	goto L125
L125:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v414<<(uint(int32(2))%32))))
	F_mark_stmt(m, v424, l1)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L34
	} else {
		goto L127
	}
L126:
	;
	goto L1
L127:
	;
	v428 = v414 + int32(1)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	if v428 < v429 {
		v414 = v428
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v441 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+16))
	if v434 < int32(0) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v437 = F_bms_is_member(m, v434, l1)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L34
	} else {
		goto L132
	}
L132:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v431)+20)) = uint8(v437)
	goto L129
L133:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v451 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L134:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v441)+16))
	if v444 < int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v447 = F_bms_is_member(m, v444, l1)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L34
	} else {
		goto L136
	}
L136:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v441)+20)) = uint8(v447)
	goto L133
L137:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v461 == int32(0) {
		goto L1
	} else {
		goto L141
	}
L138:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v451)+16))
	if v454 < int32(0) {
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v457 = F_bms_is_member(m, v454, l1)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L34
	} else {
		goto L140
	}
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v451)+20)) = uint8(v457)
	goto L137
L141:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v464 <= int32(0) {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v470 = int32(0)
	goto L143
L143:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v476+v470<<(uint(int32(2))%32))))
	F_mark_stmt(m, v480, l1)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L34
	} else {
		goto L145
	}
L144:
	;
	goto L1
L145:
	;
	v484 = v470 + int32(1)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v484 < v485 {
		v470 = v484
		goto L143
	} else {
		goto L146
	}
L146:
	;
	goto L144
L147:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v520 == int32(0) {
		goto L1
	} else {
		goto L154
	}
L148:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v490 <= int32(0) {
		goto L147
	} else {
		goto L149
	}
L149:
	;
	v495 = v3
	goto L150
L150:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v501+v495<<(uint(int32(2))%32))))
	F_mark_stmt(m, v505, l1)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L34
	} else {
		goto L152
	}
L151:
	;
	goto L147
L152:
	;
	v509 = v495 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v509 < v510 {
		v495 = v509
		goto L150
	} else {
		goto L153
	}
L153:
	;
	goto L151
L154:
	;
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v520)+16))
	if int32(0) <= v523 {
		v926 = v523
		v928 = v520
		goto L5
	} else {
		goto L155
	}
L155:
	;
	goto L1
L156:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v559 == int32(0) {
		goto L1
	} else {
		goto L163
	}
L157:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v529 <= int32(0) {
		goto L156
	} else {
		goto L158
	}
L158:
	;
	v534 = v3
	goto L159
L159:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v540+v534<<(uint(int32(2))%32))))
	F_mark_stmt(m, v544, l1)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L34
	} else {
		goto L161
	}
L160:
	;
	goto L156
L161:
	;
	v548 = v534 + int32(1)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v548 < v549 {
		v534 = v548
		goto L159
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v559)+16))
	if int32(0) <= v562 {
		v926 = v562
		v928 = v559
		goto L5
	} else {
		goto L164
	}
L164:
	;
	goto L1
L165:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v575 == int32(0) {
		goto L1
	} else {
		goto L169
	}
L166:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v565)+16))
	if v568 < int32(0) {
		goto L165
	} else {
		goto L167
	}
L167:
	;
	v571 = F_bms_is_member(m, v568, l1)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L34
	} else {
		goto L168
	}
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v565)+20)) = uint8(v571)
	goto L165
L169:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v578 <= int32(0) {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	v584 = int32(0)
	goto L171
L171:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v575)+12))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v590+v584<<(uint(int32(2))%32))))
	F_mark_stmt(m, v594, l1)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L34
	} else {
		goto L173
	}
L172:
	;
	goto L1
L173:
	;
	v598 = v584 + int32(1)
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v575)+4))
	if v598 < v599 {
		v584 = v598
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	if int32(0) <= v604 {
		v926 = v604
		v928 = v601
		goto L5
	} else {
		goto L176
	}
L176:
	;
	goto L1
L177:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v607)+16))
	if int32(0) <= v610 {
		v926 = v610
		v928 = v607
		goto L5
	} else {
		goto L178
	}
L178:
	;
	goto L1
L179:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v613)+16))
	if int32(0) <= v616 {
		v926 = v616
		v928 = v613
		goto L5
	} else {
		goto L180
	}
L180:
	;
	goto L1
L181:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v629 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(v619)+16))
	if v622 < int32(0) {
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v625 = F_bms_is_member(m, v622, l1)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L34
	} else {
		goto L184
	}
L184:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v619)+20)) = uint8(v625)
	goto L181
L185:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v639 == int32(0) {
		goto L1
	} else {
		goto L189
	}
L186:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v629)+16))
	if v632 < int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v635 = F_bms_is_member(m, v632, l1)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L34
	} else {
		goto L188
	}
L188:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v629)+20)) = uint8(v635)
	goto L185
L189:
	;
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	if v642 <= int32(0) {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	v648 = int32(0)
	goto L191
L191:
	;
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v639)+12))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v654+v648<<(uint(int32(2))%32))))
	if v658 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	goto L1
L193:
	;
	v669 = v648 + int32(1)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	if v669 < v670 {
		v648 = v669
		goto L191
	} else {
		goto L197
	}
L194:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v658)+16))
	if v661 < int32(0) {
		goto L193
	} else {
		goto L195
	}
L195:
	;
	v664 = F_bms_is_member(m, v661, l1)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L34
	} else {
		goto L196
	}
L196:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v658)+20)) = uint8(v664)
	goto L193
L197:
	;
	goto L192
L198:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v712 == int32(0) {
		goto L1
	} else {
		goto L208
	}
L199:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v675 <= int32(0) {
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v680 = v3
	goto L201
L201:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v672)+12))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v686+v680<<(uint(int32(2))%32))))
	if v690 == int32(0) {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	goto L198
L203:
	;
	v701 = v680 + int32(1)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v672)+4))
	if v701 < v702 {
		v680 = v701
		goto L201
	} else {
		goto L207
	}
L204:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(v690)+16))
	if v693 < int32(0) {
		goto L203
	} else {
		goto L205
	}
L205:
	;
	v696 = F_bms_is_member(m, v693, l1)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L34
	} else {
		goto L206
	}
L206:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v690)+20)) = uint8(v696)
	goto L203
L207:
	;
	goto L202
L208:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v715 <= int32(0) {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v721 = int32(0)
	goto L210
L210:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v712)+12))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v727+v721<<(uint(int32(2))%32))))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	if v732 == int32(0) {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	goto L1
L212:
	;
	v743 = v721 + int32(1)
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v712)+4))
	if v743 < v744 {
		v721 = v743
		goto L210
	} else {
		goto L216
	}
L213:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v732)+16))
	if v735 < int32(0) {
		goto L212
	} else {
		goto L214
	}
L214:
	;
	v738 = F_bms_is_member(m, v735, l1)
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L34
	} else {
		goto L215
	}
L215:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v732)+20)) = uint8(v738)
	goto L212
L216:
	;
	goto L211
L217:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v756 == int32(0) {
		goto L1
	} else {
		goto L221
	}
L218:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v746)+16))
	if v749 < int32(0) {
		goto L217
	} else {
		goto L219
	}
L219:
	;
	v752 = F_bms_is_member(m, v749, l1)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L34
	} else {
		goto L220
	}
L220:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v746)+20)) = uint8(v752)
	goto L217
L221:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v756)+16))
	if int32(0) <= v759 {
		v926 = v759
		v928 = v756
		goto L5
	} else {
		goto L222
	}
L222:
	;
	goto L1
L223:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v762)+16))
	if int32(0) <= v765 {
		v926 = v765
		v928 = v762
		goto L5
	} else {
		goto L224
	}
L224:
	;
	goto L1
L225:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v778 == int32(0) {
		goto L1
	} else {
		goto L229
	}
L226:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v768)+16))
	if v771 < int32(0) {
		goto L225
	} else {
		goto L227
	}
L227:
	;
	v774 = F_bms_is_member(m, v771, l1)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L34
	} else {
		goto L228
	}
L228:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v768)+20)) = uint8(v774)
	goto L225
L229:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v781 <= int32(0) {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v787 = int32(0)
	goto L231
L231:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v778)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v793+v787<<(uint(int32(2))%32))))
	if v797 == int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	goto L1
L233:
	;
	v808 = v787 + int32(1)
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v778)+4))
	if v808 < v809 {
		v787 = v808
		goto L231
	} else {
		goto L237
	}
L234:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v797)+16))
	if v800 < int32(0) {
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v803 = F_bms_is_member(m, v800, l1)
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L34
	} else {
		goto L236
	}
L236:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v797)+20)) = uint8(v803)
	goto L233
L237:
	;
	goto L232
L238:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v844 != 0 {
		goto L3
	} else {
		goto L245
	}
L239:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v814 <= int32(0) {
		goto L238
	} else {
		goto L240
	}
L240:
	;
	v819 = v3
	goto L241
L241:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v811)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v825+v819<<(uint(int32(2))%32))))
	F_mark_stmt(m, v829, l1)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L34
	} else {
		goto L243
	}
L242:
	;
	goto L238
L243:
	;
	v833 = v819 + int32(1)
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v811)+4))
	if v833 < v834 {
		v819 = v833
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	goto L2
L246:
	;
	v855 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v855 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v845)+16))
	if v848 < int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v851 = F_bms_is_member(m, v848, l1)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L34
	} else {
		goto L249
	}
L249:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v845)+20)) = uint8(v851)
	goto L246
L250:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v865 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L251:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v855)+16))
	if v858 < int32(0) {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v861 = F_bms_is_member(m, v858, l1)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L34
	} else {
		goto L253
	}
L253:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v855)+20)) = uint8(v861)
	goto L250
L254:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v875 == int32(0) {
		goto L1
	} else {
		goto L258
	}
L255:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v865)+16))
	if v868 < int32(0) {
		goto L254
	} else {
		goto L256
	}
L256:
	;
	v871 = F_bms_is_member(m, v868, l1)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L34
	} else {
		goto L257
	}
L257:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v865)+20)) = uint8(v871)
	goto L254
L258:
	;
	v878 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v878 <= int32(0) {
		goto L1
	} else {
		goto L259
	}
L259:
	;
	v884 = int32(0)
	goto L260
L260:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v875)+12))
	v894 = *(*int32)(unsafe.Add(mBase, uint32(v890+v884<<(uint(int32(2))%32))))
	if v894 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L261:
	;
	goto L1
L262:
	;
	v905 = v884 + int32(1)
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v875)+4))
	if v905 < v906 {
		v884 = v905
		goto L260
	} else {
		goto L266
	}
L263:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v894)+16))
	if v897 < int32(0) {
		goto L262
	} else {
		goto L264
	}
L264:
	;
	v900 = F_bms_is_member(m, v897, l1)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L34
	} else {
		goto L265
	}
L265:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v894)+20)) = uint8(v900)
	goto L262
L266:
	;
	goto L261
L267:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v908)+16))
	if int32(0) <= v911 {
		v926 = v911
		v928 = v908
		goto L5
	} else {
		goto L268
	}
L268:
	;
	goto L1
L269:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v914)+16))
	if int32(0) <= v917 {
		v926 = v917
		v928 = v914
		goto L5
	} else {
		goto L270
	}
L270:
	;
	goto L1
L271:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v920)+16))
	if v923 < int32(0) {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	v926 = v923
	v928 = v920
	goto L5
L273:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v928)+20)) = uint8(v934)
	goto L1
L274:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v941
	F_errmsg_internal(m, int32(_a_F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2_1), v11)
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L34
	} else {
		goto L275
	}
L275:
	;
	F_errfinish(m, int32(_a_F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2_2), int32(595), int32(_a_F_plpgsql_statement_tree_walker_impl_x2especialized_x2e2_3))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L34
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
	v954 = F_bms_is_member(m, v951, l1)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L34
	} else {
		goto L278
	}
L278:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v844)+20)) = uint8(v954)
	goto L2
L279:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v961 <= int32(0) {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	v967 = int32(0)
	goto L281
L281:
	;
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v958)+12))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v973+v967<<(uint(int32(2))%32))))
	if v977 == int32(0) {
		goto L283
	} else {
		goto L284
	}
L282:
	;
	goto L1
L283:
	;
	v988 = v967 + int32(1)
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v988 < v989 {
		v967 = v988
		goto L281
	} else {
		goto L287
	}
L284:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v977)+16))
	if v980 < int32(0) {
		goto L283
	} else {
		goto L285
	}
L285:
	;
	v983 = F_bms_is_member(m, v980, l1)
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L34
	} else {
		goto L286
	}
L286:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v977)+20)) = uint8(v983)
	goto L283
L287:
	;
	goto L282
}
