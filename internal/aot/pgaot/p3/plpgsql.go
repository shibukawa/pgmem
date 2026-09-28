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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_build_datatype_0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_plpgsql_build_datatype_1), v9)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_plpgsql_build_datatype_2), int32(1983), int32(_a_F_plpgsql_build_datatype_3))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v32 = F_build_datatype(m, v13, l1, l2, l3)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v13)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v32
				}
			}
		}
	}
}
func F_plpgsql_call_handler(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
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
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v194 int64
	_ = v194
	var v195 int32
	_ = v195
	var v202 int64
	_ = v202
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v307 int32
	_ = v307
	var v322 int32
	_ = v322
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
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
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(224)
	m.G0 = v18
	v23 = v2
	v24 = v2
	v25 = v2
	v26 = int32(-1)
	v27 = v2
	v28 = v2
	v29 = v2
	v30 = v2
	v31 = v2
	v32 = v2
	v33 = v2
	v34 = v2
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
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v18)+176))
	m.G0 = v18 + int32(224)
	return v366
L4:
	;
	goto L3
L5:
	;
	if v26 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v335 = int32(m.ExcTag)
	v336 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v335 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+176)) = int64(0)
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v41 == v40 {
		v50 = v40
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v124 = v23
	v125 = v24
	v126 = v25
	v128 = v27
	v129 = v28
	v130 = v29
	v131 = v30
	v132 = v31
	v133 = v32
	v134 = v33
	v135 = v34
	goto L10
L10:
	;
	if v135 != 0 {
		goto L24
	} else {
		goto L25
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v23
	v60 = l0 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v60
	v63 = v50 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v63)
	F_SPI_connect_ext(m, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v44 != int32(214) {
		v50 = v40
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	v50 = v47 ^ int32(1)
	goto L11
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v63)
	v78 = F_plpgsql_compile(m, l0, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v78)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v78)+24)) = v80 + int64(1)
	v85 = v78 + int32(24)
	v87 = v78 + int32(528)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+528))
	v89 = int32(0)
	if v63 == v89 {
		v111 = v33
		v112 = v89
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0]))
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1]))
	goto L20
L17:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+524)))
	if v94 != int32(1) {
		v111 = v33
		v112 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v60
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v63)
	v109 = F_ResourceOwnerCreate(m, int32(0), int32(_a_F_plpgsql_call_handler_0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	v111 = v109
	v112 = v109
	goto L16
L20:
	;
	v118 = v18 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v118))) = v18 + int32(12)
	goto L23
L21:
	;
	v124 = v78
	v125 = v112
	v126 = v60
	v128 = v116
	v129 = v114
	v130 = v85
	v131 = v88
	v132 = v87
	v133 = v50
	v134 = v111
	v135 = v89
	goto L10
L23:
	;
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0])) = v129
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v128
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = v202 - int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v131
	if v125 != 0 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v18 + int32(16)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	if v140 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v185 = int32(1)
	v186 = v133 & v185
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v186)
	v188 = int32(0)
	v194 = F_plpgsql_exec_function(m, v124, l0, v188, v188, v125, (v133^int32(-1))&v185)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L7
	} else {
		goto L32
	}
L27:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	switch v143 - int32(447) {
	case 0:
		goto L28
	case 1:
		goto L29
	default:
		goto L26
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v172 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v172)
	F_plpgsql_exec_event_trigger(m, v124, v140)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v156 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v156)
	v158 = F_plpgsql_exec_trigger(m, v124, v140)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+176)) = base.I64_extend_i32_u(v158)
	goto L24
L31:
	;
	goto L24
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+176)) = v194
	goto L24
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v217 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v217)
	F_ReleaseAllPlanCacheRefsInOwner(m, v125)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v135 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v217)
	F_ResourceOwnerDelete(m, v125)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v244 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v244)
	F_pg_re_throw(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[0])) = v129
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_call_handler[1])) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	v262 = v133 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v262)
	v264 = F_SPI_finish(m)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L7
	} else {
		goto L42
	}
L41:
	;
	goto L1
L42:
	;
	if v264 == int32(2) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v262)
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_call_handler_1))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v262)
	v292 = F_SPI_result_code_string(m, v264)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v262)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v292
	F_errmsg_internal(m, int32(_a_F_plpgsql_call_handler_2), v18)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+188)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v18)+184)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v18)+192)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v18)+196)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v18)+200)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v18)+204)) = v131
	*(*int32)(unsafe.Add(mBase, uint32(v18)+208)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v18)+220)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)) = uint8(v262)
	F_errfinish(m, int32(_a_F_plpgsql_call_handler_3), int32(302), int32(_a_F_plpgsql_call_handler_4))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	goto L1
L48:
	;
	v340 = int32(v336)
	m.G0 = v18
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v340)+4))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	if v18+int32(12) == v346 {
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
	if v350 != 0 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v343)+4))
	v350 = v348
	goto L53
L52:
	;
	v350 = int32(0)
	goto L53
L53:
	;
	goto L50
L54:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v18)+220))
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+219)))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v18)+212))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v18)+208))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v18)+204))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v18)+200))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v18)+196))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v18)+188))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v18)+184))
	v23 = v353
	v24 = v358
	v25 = v351
	v26 = v350
	v27 = v359
	v28 = v360
	v29 = v356
	v30 = v355
	v31 = v354
	v32 = v352
	v33 = v357
	v34 = v342
	goto L2
L55:
	;
	goto L56
L56:
	;
	F___wasm_longjmp(m, v343, v342)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	return int64(0)
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
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
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
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
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v578 int32
	_ = v578
	var v590 int32
	_ = v590
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1005 int32
	_ = v1005
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1031 int32
	_ = v1031
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1127 int32
	_ = v1127
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1145 int32
	_ = v1145
	var v1149 int32
	_ = v1149
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1200 int32
	_ = v1200
	var v1205 int32
	_ = v1205
	var v1209 int32
	_ = v1209
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1408 int32
	_ = v1408
	var v1415 int32
	_ = v1415
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1445 int32
	_ = v1445
	var v1448 int32
	_ = v1448
	var v1451 int32
	_ = v1451
	var v1453 int32
	_ = v1453
	var v1461 int32
	_ = v1461
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1502 int32
	_ = v1502
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1540 int32
	_ = v1540
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1557 int32
	_ = v1557
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1586 int32
	_ = v1586
	var v1588 int32
	_ = v1588
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1602 int32
	_ = v1602
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1639 int32
	_ = v1639
	var v1658 int32
	_ = v1658
	var v1662 int32
	_ = v1662
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1704 int32
	_ = v1704
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1725 int32
	_ = v1725
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1747 int32
	_ = v1747
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1763 int32
	_ = v1763
	var v1767 int32
	_ = v1767
	var v1773 int32
	_ = v1773
	var v1778 int32
	_ = v1778
	var v1803 int32
	_ = v1803
	v5 = l4
	v6 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(416)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v31 == int32(447) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v41 = int32(2)
	goto L3
L3:
	;
	v44 = F_SysCacheGetAttrNotNull(m, int32(47), l1, int32(26))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v34 = int32(1)
	goto L6
L5:
	;
	v34 = int32(2)
	goto L6
L6:
	;
	if v31 != int32(448) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v38 = v34
	goto L9
L8:
	;
	v38 = int32(0)
	goto L9
L9:
	;
	v41 = v38
	goto L3
L10:
	;
	return
L11:
	;
	v48 = F_text_to_cstring(m, base.I32_wrap_i64(v44))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = F_plpgsql_scanner_init(m, v48)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v54 = v27 + int32(4)
	v55 = F_pstrdup(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0])) = v55
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[1])) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[2])) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v23)+412)) = v50
	if v5 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = v48
	goto L17
L16:
	;
	v64 = int32(0)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+408)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v23)+400)) = int32(_a_F_plpgsql_compile_callback_0)
	v68 = int32(_a_F_plpgsql_compile_callback_1)
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3])) = v23 + int32(396)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+396)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v23)+404)) = v23 + int32(408)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	v80 = F_format_procedure(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4]))
	v88 = F_AllocSetContextCreateInternal(m, v83, int32(_a_F_plpgsql_compile_callback_2), int32(0), int32(_a_F_plpgsql_compile_callback_3), int32(_a_F_plpgsql_compile_callback_4))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	v91 = int32(_a_F_plpgsql_compile_callback_5)
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v88
	v96 = F_pstrdup(m, v80)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+32)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = v96
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+36)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+48)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(l3)+44)) = v103
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[6]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+488)) = v109
	v112 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[7])))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+492)) = uint8(v112)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[8]))
	if v5 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v117 = v115
	goto L23
L22:
	;
	v117 = int32(0)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+496)) = v117
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[9]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v41
	if v5 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v123 = v120
	goto L26
L25:
	;
	v123 = int32(0)
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+500)) = v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+96)))
	v126 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l3)+524)) = uint16(v126)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+520)) = v126
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)) = uint8(v125)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[10])) = v126
	goto L27
L27:
	;
	F_plpgsql_ns_push(m, v54, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = int32(128)
	v141 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[12])) = uint8(v141)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v141
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	v149 = F_MemoryContextAlloc(m, v147, int32(512))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[14])) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v149
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	switch v156 {
	case 0:
		goto L58
	case 1:
		goto L57
	case 2:
		goto L59
	default:
		goto L39
	}
L30:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(1983), int32(_a_F_plpgsql_compile_callback_7))
	mBase = m.M
	v1803 = m.ExcPending
	if v1803 != 0 {
		goto L10
	} else {
		goto L394
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1767 = m.ExcPending
	if v1767 != 0 {
		goto L10
	} else {
		goto L391
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1756 = m.ExcPending
	if v1756 != 0 {
		goto L10
	} else {
		goto L389
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L10
	} else {
		goto L385
	}
L34:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L10
	} else {
		goto L382
	}
L35:
	;
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+101)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+64)) = uint8(base.B2i32(v1353 != int32(118)))
	v1359 = F_SearchSysCache1(m, int32(82), int64(16))
	mBase = m.M
	v1360 = m.ExcPending
	if v1360 != 0 {
		goto L10
	} else {
		goto L310
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v1244
	v1246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)) = uint8(v1246)
	v1250 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v1244))
	mBase = m.M
	v1251 = m.ExcPending
	if v1251 != 0 {
		goto L10
	} else {
		goto L283
	}
L37:
	;
	v1244 = int32(23)
	goto L36
L38:
	;
	v1244 = int32(3904)
	goto L36
L39:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L10
	} else {
		goto L280
	}
L40:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L10
	} else {
		goto L278
	}
L41:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L10
	} else {
		goto L276
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L10
	} else {
		goto L272
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1182 = m.ExcPending
	if v1182 != 0 {
		goto L10
	} else {
		goto L270
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L10
	} else {
		goto L268
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L10
	} else {
		goto L266
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L10
	} else {
		goto L264
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L10
	} else {
		goto L262
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L10
	} else {
		goto L260
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L10
	} else {
		goto L258
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L10
	} else {
		goto L256
	}
L51:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L10
	} else {
		goto L254
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		goto L10
	} else {
		goto L252
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L10
	} else {
		goto L247
	}
L54:
	;
	switch v590 - int32(_a_F_plpgsql_compile_callback_9) {
	case 0:
		v1244 = v613
		goto L36
	default:
		goto L37
	case 2:
		goto L38
	}
L55:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L10
	} else {
		goto L242
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L10
	} else {
		goto L240
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(2278)
	v981 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)))
	if v981 != 0 {
		goto L42
	} else {
		goto L229
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+104)))
	if v650 != 0 {
		goto L53
	} else {
		goto L164
	}
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v159
	v167 = F_get_func_arg_info(m, l1, v23+int32(392), v23+int32(388), v23+int32(384))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)+392))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v23)+384))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+24))
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0]))
	F_cfunc_resolve_polymorphic_argtypes(m, v167, v169, v170, v172, v5, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	v178 = F_palloc_mul(m, int32(4), v167)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	v181 = F_palloc_mul(m, int32(4), v167)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v88
	if v167 <= int32(0) {
		v578 = v6
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	if v590 <= int32(3830) {
		goto L145
	} else {
		goto L146
	}
L65:
	;
	v190 = int32(0)
	v196 = v6
	v197 = int32(0)
	goto L66
L66:
	;
	v209 = v190 << (uint(int32(2)) % 32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v23)+392))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v209+v210)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v23)+384))
	if v213 != 0 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v296 <= int32(1) {
		goto L100
	} else {
		goto L101
	}
L68:
	;
	v215 = int32(*(*int8)(unsafe.Add(mBase, uint32(v190+v213))))
	v217 = v215
	goto L70
L69:
	;
	v217 = int32(105)
	goto L70
L70:
	;
	v219 = v190 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v219
	v227 = F_pg_snprintf(m, v23+int32(352), int32(32), int32(_a_F_plpgsql_compile_callback_10), v23+int32(144))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v232 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v212))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	if v232 == int32(0) {
		goto L56
	} else {
		goto L73
	}
L73:
	;
	v238 = F_build_datatype(m, v232, int32(-1), v229, int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	F_ReleaseCatCache(m, v232)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	if v242 == int32(2) {
		goto L55
	} else {
		goto L76
	}
L76:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v23)+388))
	if v245 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v253 = int32(0)
	v255 = F_plpgsql_build_variable(m, v252, v253, v238, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L10
	} else {
		goto L82
	}
L78:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v245+v209)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	if v248 != 0 {
		v252 = v247
		goto L77
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v252 = v23 + int32(352)
	goto L77
L81:
	;
	goto L80
L82:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v259 = v217 - int32(98)
	v266 = int32(0)
	if base.B2i32(base.Ui32(int32(20)) < base.Ui32(v259))|base.B2i32(int32(1)<<(uint(v259)%32)&int32(_a_F_plpgsql_compile_callback_11) == v266) == v266 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v178+v197<<(uint(int32(2))%32)))) = v274
	v278 = v197 + int32(1)
	goto L85
L84:
	;
	v278 = v197
	goto L85
L85:
	;
	v283 = int32(0)
	if base.B2i32(int32(1)<<(uint(v259)%32)&int32(_a_F_plpgsql_compile_callback_12) == v283)|base.B2i32(base.Ui32(int32(18)) < base.Ui32(v259)) == v283 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181+v196<<(uint(int32(2))%32)))) = v255
	v296 = v196 + int32(1)
	goto L88
L87:
	;
	v296 = v196
	goto L88
L88:
	;
	if v257 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v299 = int32(2)
	goto L91
L90:
	;
	v299 = int32(1)
	goto L91
L91:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	F_add_parameter_name(m, v299, v300, v23+int32(352))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v23)+388))
	if v305 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v219 != v167 {
		v190 = v219
		v196 = v296
		v197 = v278
		goto L66
	} else {
		goto L97
	}
L94:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v305+v209)))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v310 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	F_add_parameter_name(m, v299, v313, v309)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	goto L93
L97:
	;
	goto L67
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+472)) = v550
	v578 = v557
	goto L64
L99:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+4))
	v550 = v548
	v557 = v322
	goto L98
L100:
	;
	if v296 != int32(1) {
		v578 = v296
		goto L64
	} else {
		goto L103
	}
L101:
	;
	v326 = v296
	goto L102
L102:
	;
	v328 = F_palloc0(m, int32(40))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L10
	} else {
		goto L105
	}
L103:
	;
	v322 = int32(1)
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+65)))
	if v323 != int32(112) {
		goto L99
	} else {
		goto L104
	}
L104:
	;
	v326 = v322
	goto L102
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v328)+8)) = int32(_a_F_plpgsql_compile_callback_13)
	*(*int32)(unsafe.Add(mBase, uint32(v328))) = int32(1)
	v336 = F_CreateTemplateTupleDesc(m, v326)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L10
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+28)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v328)+24)) = v336
	v341 = F_palloc_mul(m, int32(4), v326)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+32)) = v341
	v345 = F_palloc_mul(m, int32(4), v326)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+36)) = v345
	v350 = int32(0)
	goto L109
L109:
	;
	v370 = v350 << (uint(int32(2)) % 32)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v181+v370)))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	switch v373 {
	case 0, 4:
		goto L112
	default:
		goto L113
	case 2:
		goto L114
	}
L110:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v430 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v429)))
	if v430 < v439 {
		goto L122
	} else {
		goto L123
	}
L111:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v399)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v328)+32))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v404+v370))) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v328)+36))
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v408+v370))) = v410
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v414 = v350 + int32(1)
	v415 = base.I32_extend16_s(v414)
	F_TupleDescInitEntry(m, v412, v415, v406, v403, v402, int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L10
	} else {
		goto L118
	}
L112:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v372)+24))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v394)+24))
	v399 = v394 + int32(4)
	v401 = v397
	v402 = v398
	goto L111
L113:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L115
	}
L114:
	;
	v399 = v372 + int32(28)
	v401 = int32(0)
	v402 = int32(-1)
	goto L111
L115:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v382
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_14), v23+int32(48))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L10
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(1898), int32(_a_F_plpgsql_compile_callback_15))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L10
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v328)+24))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	*(*int32)(unsafe.Add(mBase, uint32(v419+v420<<(uint(int32(3))%32)+v415*int32(100))+24)) = v401
	goto L119
L119:
	;
	if v414 != v326 {
		v350 = v414
		goto L109
	} else {
		goto L120
	}
L120:
	;
	goto L110
L121:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v520 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v520 == v522 {
		goto L140
	} else {
		goto L141
	}
L122:
	;
	v443 = v429 + int32(28)
	v450 = v430
	v451 = v439
	v453 = v430
	goto L126
L123:
	;
	v507 = v430
	v514 = v439
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v429)+20)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v429)+16)) = v507
	goto L121
L125:
	;
	v507 = v501
	v514 = v480
	goto L124
L126:
	;
	v459 = v443 + v439<<(uint(int32(3))%32) + v450*int32(100)
	v462 = v443 + v450<<(uint(int32(3))%32)
	if v439 != v451 {
		v480 = v451
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v501 = v439
	goto L125
L128:
	;
	v481 = int32(*(*int16)(unsafe.Add(mBase, uint32(v462)+2)))
	if v481 <= int32(0) {
		v501 = v450
		goto L125
	} else {
		goto L136
	}
L129:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+7)))
	if v464 != int32(118) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v480 = v450
	goto L128
L131:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+4)))
	if v467 != int32(1) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+6)))
	if v470&int32(6) != 0 {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v473 = int32(*(*int16)(unsafe.Add(mBase, uint32(v462)+2)))
	if v473 <= int32(0) {
		goto L130
	} else {
		goto L134
	}
L134:
	;
	v476 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+90)))
	if v476 != int32(118) {
		v480 = v439
		goto L128
	} else {
		goto L135
	}
L135:
	;
	goto L130
L136:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v459)+90)))
	if v484 == int32(118) {
		v501 = v450
		goto L125
	} else {
		goto L137
	}
L137:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v462)+5)))
	v493 = (v453 + v487 - int32(1)) & (int32(0) - v487)
	if int32(_a_F_plpgsql_compile_callback_16) < v493 {
		v501 = v450
		goto L125
	} else {
		goto L138
	}
L138:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v462))) = uint16(v493)
	v499 = v450 + int32(1)
	if v499 != v439 {
		v450 = v499
		v451 = v480
		v453 = v493 + v481
		goto L126
	} else {
		goto L139
	}
L139:
	;
	goto L127
L140:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v520 << (uint(int32(1)) % 32)
	v531 = F_repalloc(m, v518, v520<<(uint(int32(3))%32))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L10
	} else {
		goto L143
	}
L141:
	;
	v536 = v520
	v537 = v518
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v328)+4)) = v536
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v536 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v537+v536<<(uint(int32(2))%32)))) = v328
	v550 = v536
	v557 = v326
	goto L98
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v531
	v535 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v536 = v535
	v537 = v531
	goto L142
L144:
	;
	if v5 != 0 {
		goto L151
	} else {
		goto L152
	}
L145:
	;
	switch v590 - int32(2277) {
	case 0, 6:
		goto L144
	case 1, 2, 3, 4, 5:
		v1244 = v590
		goto L36
	default:
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	if base.B2i32(base.Ui32(v590-int32(_a_F_plpgsql_compile_callback_17)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v590-int32(_a_F_plpgsql_compile_callback_18)) < base.Ui32(int32(2)))|base.B2i32(v590 == int32(3831)) != 0 {
		goto L144
	} else {
		goto L150
	}
L148:
	;
	if base.B2i32(v590 == int32(2776))|base.B2i32(v590 == int32(3500)) != 0 {
		goto L144
	} else {
		goto L149
	}
L149:
	;
	v1244 = v590
	goto L36
L150:
	;
	v1244 = v590
	goto L36
L151:
	;
	v613 = int32(1007)
	if int32(_a_F_plpgsql_compile_callback_17) < v590 {
		goto L54
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v624 = F_get_fn_expr_rettype(m, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L10
	} else {
		goto L158
	}
L154:
	;
	if v590 == int32(2277) {
		v1244 = v613
		goto L36
	} else {
		goto L155
	}
L155:
	;
	if v590 == int32(3831) {
		goto L38
	} else {
		goto L156
	}
L156:
	;
	if v590 != int32(_a_F_plpgsql_compile_callback_18) {
		goto L37
	} else {
		goto L157
	}
L157:
	;
	v1244 = int32(_a_F_plpgsql_compile_callback_19)
	goto L36
L158:
	;
	if v624 != 0 {
		v1244 = v624
		goto L36
	} else {
		goto L159
	}
L159:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L10
	} else {
		goto L160
	}
L160:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L10
	} else {
		goto L161
	}
L161:
	;
	v634 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v634
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_20), v23+int32(128))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L10
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(427), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L164:
	;
	v652 = F_palloc0(m, int32(40))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L10
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652))) = int32(2)
	v657 = F_pstrdup(m, int32(_a_F_plpgsql_compile_callback_22))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v652)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v652)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v652)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v652)+8)) = v657
	v667 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v669 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v671 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v669 == v671 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v669 << (uint(int32(1)) % 32)
	v680 = F_repalloc(m, v667, v669<<(uint(int32(3))%32))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L10
	} else {
		goto L170
	}
L168:
	;
	v686 = v657
	v687 = v669
	v688 = v667
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v652)+4)) = v687
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v687 + int32(1)
	v694 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v688+v687<<(uint(v694)%32)))) = v652
	F_plpgsql_ns_additem(m, v694, v687, v686)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L10
	} else {
		goto L171
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v680
	v684 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v652)+8))
	v686 = v685
	v687 = v684
	v688 = v680
	goto L169
L171:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+480)) = v701
	v704 = F_palloc0(m, int32(40))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L10
	} else {
		goto L172
	}
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v704))) = int32(2)
	v709 = F_pstrdup(m, int32(_a_F_plpgsql_compile_callback_23))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v704)+32)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v704)+24)) = int64(9659381448704)
	*(*int32)(unsafe.Add(mBase, uint32(v704)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v704)+8)) = v709
	v719 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v721 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v723 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11]))
	if v721 == v723 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[11])) = v721 << (uint(int32(1)) % 32)
	v732 = F_repalloc(m, v719, v721<<(uint(int32(3))%32))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L10
	} else {
		goto L177
	}
L175:
	;
	v738 = v709
	v739 = v721
	v740 = v719
	goto L176
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v704)+4)) = v739
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13])) = v739 + int32(1)
	v746 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v740+v739<<(uint(v746)%32)))) = v704
	F_plpgsql_ns_additem(m, v746, v739, v738)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L10
	} else {
		goto L178
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15])) = v732
	v736 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v704)+8))
	v738 = v737
	v739 = v736
	v740 = v732
	goto L176
L178:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v704)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+484)) = v753
	v755 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v758 = F_SearchSysCache1(m, int32(82), int64(19))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L10
	} else {
		goto L179
	}
L179:
	;
	if v758 == int32(0) {
		goto L52
	} else {
		goto L180
	}
L180:
	;
	v764 = F_build_datatype(m, v758, int32(-1), v755, int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L10
	} else {
		goto L181
	}
L181:
	;
	F_ReleaseCatCache(m, v758)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L10
	} else {
		goto L182
	}
L182:
	;
	v771 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_24), int32(0), v764, int32(1))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L10
	} else {
		goto L183
	}
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v771)+52)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v771))) = int32(4)
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v780 = F_SearchSysCache1(m, int32(82), int64(25))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L10
	} else {
		goto L184
	}
L184:
	;
	if v780 == int32(0) {
		goto L51
	} else {
		goto L185
	}
L185:
	;
	v786 = F_build_datatype(m, v780, int32(-1), v777, int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L10
	} else {
		goto L186
	}
L186:
	;
	F_ReleaseCatCache(m, v780)
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L10
	} else {
		goto L187
	}
L187:
	;
	v793 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_25), int32(0), v786, int32(1))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L10
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v793)+52)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v793))) = int32(4)
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v802 = F_SearchSysCache1(m, int32(82), int64(25))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L10
	} else {
		goto L189
	}
L189:
	;
	if v802 == int32(0) {
		goto L50
	} else {
		goto L190
	}
L190:
	;
	v808 = F_build_datatype(m, v802, int32(-1), v799, int32(0))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L10
	} else {
		goto L191
	}
L191:
	;
	F_ReleaseCatCache(m, v802)
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	v815 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_26), int32(0), v808, int32(1))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L10
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v815)+52)) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v815))) = int32(4)
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v824 = F_SearchSysCache1(m, int32(82), int64(25))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L10
	} else {
		goto L194
	}
L194:
	;
	if v824 == int32(0) {
		goto L49
	} else {
		goto L195
	}
L195:
	;
	v830 = F_build_datatype(m, v824, int32(-1), v821, int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	F_ReleaseCatCache(m, v824)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L10
	} else {
		goto L197
	}
L197:
	;
	v837 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_27), int32(0), v830, int32(1))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	v839 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v837)+52)) = v839
	*(*int32)(unsafe.Add(mBase, uint32(v837))) = v839
	v845 = F_SearchSysCache1(m, int32(82), int64(26))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	if v845 == int32(0) {
		goto L48
	} else {
		goto L200
	}
L200:
	;
	v850 = int32(0)
	v852 = F_build_datatype(m, v845, int32(-1), v850, v850)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L10
	} else {
		goto L201
	}
L201:
	;
	F_ReleaseCatCache(m, v845)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L202
	}
L202:
	;
	v859 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_28), int32(0), v852, int32(1))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L10
	} else {
		goto L203
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v859)+52)) = int32(5)
	*(*int32)(unsafe.Add(mBase, uint32(v859))) = int32(4)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v868 = F_SearchSysCache1(m, int32(82), int64(19))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L10
	} else {
		goto L204
	}
L204:
	;
	if v868 == int32(0) {
		goto L47
	} else {
		goto L205
	}
L205:
	;
	v874 = F_build_datatype(m, v868, int32(-1), v865, int32(0))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L10
	} else {
		goto L206
	}
L206:
	;
	F_ReleaseCatCache(m, v868)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L10
	} else {
		goto L207
	}
L207:
	;
	v881 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_29), int32(0), v874, int32(1))
	mBase = m.M
	v882 = m.ExcPending
	if v882 != 0 {
		goto L10
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v881)+52)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v881))) = int32(4)
	v887 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v890 = F_SearchSysCache1(m, int32(82), int64(19))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L10
	} else {
		goto L209
	}
L209:
	;
	if v890 == int32(0) {
		goto L46
	} else {
		goto L210
	}
L210:
	;
	v896 = F_build_datatype(m, v890, int32(-1), v887, int32(0))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L10
	} else {
		goto L211
	}
L211:
	;
	F_ReleaseCatCache(m, v890)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L10
	} else {
		goto L212
	}
L212:
	;
	v903 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_30), int32(0), v896, int32(1))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L10
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v903)+52)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v903))) = int32(4)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v912 = F_SearchSysCache1(m, int32(82), int64(19))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L10
	} else {
		goto L214
	}
L214:
	;
	if v912 == int32(0) {
		goto L45
	} else {
		goto L215
	}
L215:
	;
	v918 = F_build_datatype(m, v912, int32(-1), v909, int32(0))
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L10
	} else {
		goto L216
	}
L216:
	;
	F_ReleaseCatCache(m, v912)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L10
	} else {
		goto L217
	}
L217:
	;
	v925 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_31), int32(0), v918, int32(1))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L10
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v925)+52)) = int32(7)
	*(*int32)(unsafe.Add(mBase, uint32(v925))) = int32(4)
	v933 = F_SearchSysCache1(m, int32(82), int64(23))
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L10
	} else {
		goto L219
	}
L219:
	;
	if v933 == int32(0) {
		goto L44
	} else {
		goto L220
	}
L220:
	;
	v937 = int32(0)
	v941 = F_build_datatype(m, v933, int32(-1), v937, v937)
	mBase = m.M
	v942 = m.ExcPending
	if v942 != 0 {
		goto L10
	} else {
		goto L221
	}
L221:
	;
	F_ReleaseCatCache(m, v933)
	mBase = m.M
	v944 = m.ExcPending
	if v944 != 0 {
		goto L10
	} else {
		goto L222
	}
L222:
	;
	v948 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_32), int32(0), v941, int32(1))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L10
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v948)+52)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v948))) = int32(4)
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v957 = F_SearchSysCache1(m, int32(82), int64(1009))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L10
	} else {
		goto L224
	}
L224:
	;
	if v957 == int32(0) {
		goto L43
	} else {
		goto L225
	}
L225:
	;
	v963 = F_build_datatype(m, v957, int32(-1), v954, int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L10
	} else {
		goto L226
	}
L226:
	;
	F_ReleaseCatCache(m, v957)
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L10
	} else {
		goto L227
	}
L227:
	;
	v970 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_33), int32(0), v963, int32(1))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L10
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v970)+52)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v970))) = int32(4)
	v1334 = int32(0)
	v1339 = v937
	goto L35
L229:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v985 = F_SearchSysCache1(m, int32(82), int64(25))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L10
	} else {
		goto L230
	}
L230:
	;
	if v985 == int32(0) {
		goto L41
	} else {
		goto L231
	}
L231:
	;
	v989 = int32(0)
	v992 = F_build_datatype(m, v985, int32(-1), v982, v989)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L10
	} else {
		goto L232
	}
L232:
	;
	F_ReleaseCatCache(m, v985)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L10
	} else {
		goto L233
	}
L233:
	;
	v999 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_34), int32(0), v992, int32(1))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v999)+52)) = int32(10)
	*(*int32)(unsafe.Add(mBase, uint32(v999))) = int32(4)
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v1008 = F_SearchSysCache1(m, int32(82), int64(25))
	mBase = m.M
	v1009 = m.ExcPending
	if v1009 != 0 {
		goto L10
	} else {
		goto L235
	}
L235:
	;
	if v1008 == int32(0) {
		goto L40
	} else {
		goto L236
	}
L236:
	;
	v1014 = F_build_datatype(m, v1008, int32(-1), v1005, int32(0))
	mBase = m.M
	v1015 = m.ExcPending
	if v1015 != 0 {
		goto L10
	} else {
		goto L237
	}
L237:
	;
	F_ReleaseCatCache(m, v1008)
	mBase = m.M
	v1017 = m.ExcPending
	if v1017 != 0 {
		goto L10
	} else {
		goto L238
	}
L238:
	;
	v1021 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_35), int32(0), v1014, int32(1))
	mBase = m.M
	v1022 = m.ExcPending
	if v1022 != 0 {
		goto L10
	} else {
		goto L239
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1021)+52)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v1021))) = int32(4)
	v1334 = v989
	v1339 = int32(0)
	goto L35
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v212
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(16))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L10
	} else {
		goto L241
	}
L241:
	;
	goto L30
L242:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L10
	} else {
		goto L243
	}
L243:
	;
	v1045 = F_format_type_be(m, v212)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L10
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v1045
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_37), v23+int32(32))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L10
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(334), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L10
	} else {
		goto L246
	}
L246:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L247:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_38), int32(0))
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L10
	} else {
		goto L249
	}
L249:
	;
	F_errhint(m, int32(_a_F_plpgsql_compile_callback_39), int32(0))
	mBase = m.M
	v1074 = m.ExcPending
	if v1074 != 0 {
		goto L10
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(500), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1079 = m.ExcPending
	if v1079 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+160)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(160))
	mBase = m.M
	v1090 = m.ExcPending
	if v1090 != 0 {
		goto L10
	} else {
		goto L253
	}
L253:
	;
	goto L30
L254:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(176))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L10
	} else {
		goto L255
	}
L255:
	;
	goto L30
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+192)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(192))
	mBase = m.M
	v1112 = m.ExcPending
	if v1112 != 0 {
		goto L10
	} else {
		goto L257
	}
L257:
	;
	goto L30
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+208)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(208))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L10
	} else {
		goto L259
	}
L259:
	;
	goto L30
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+224)) = int32(26)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(224))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L10
	} else {
		goto L261
	}
L261:
	;
	goto L30
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+240)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(240))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L10
	} else {
		goto L263
	}
L263:
	;
	goto L30
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+256)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(256))
	mBase = m.M
	v1156 = m.ExcPending
	if v1156 != 0 {
		goto L10
	} else {
		goto L265
	}
L265:
	;
	goto L30
L266:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+272)) = int32(19)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(272))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L10
	} else {
		goto L267
	}
L267:
	;
	goto L30
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+288)) = int32(23)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(288))
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L10
	} else {
		goto L269
	}
L269:
	;
	goto L30
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+304)) = int32(1009)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(304))
	mBase = m.M
	v1189 = m.ExcPending
	if v1189 != 0 {
		goto L10
	} else {
		goto L271
	}
L271:
	;
	goto L30
L272:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L10
	} else {
		goto L273
	}
L273:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_40), int32(0))
	mBase = m.M
	v1200 = m.ExcPending
	if v1200 != 0 {
		goto L10
	} else {
		goto L274
	}
L274:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(633), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v23)+320)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(320))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L10
	} else {
		goto L277
	}
L277:
	;
	goto L30
L278:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+336)) = int32(25)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(336))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L10
	} else {
		goto L279
	}
L279:
	;
	goto L30
L280:
	;
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v1232
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_41), v23)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L10
	} else {
		goto L281
	}
L281:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(661), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L10
	} else {
		goto L282
	}
L282:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L283:
	;
	if v1250 == int32(0) {
		goto L34
	} else {
		goto L284
	}
L284:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1250)+16))
	v1257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1256)+22)))
	v1258 = v1256 + v1257
	v1259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+79)))
	if base.B2i32(v1244 == int32(2249))|base.B2i32(v1259 != int32(112)) != 0 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1287 = F_type_is_rowtype(m, v1244)
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L10
	} else {
		goto L294
	}
L286:
	;
	switch v1244 - int32(2278) {
	case 0:
		goto L285
	case 1:
		goto L33
	default:
		goto L287
	}
L287:
	;
	if v1244 == int32(3838) {
		goto L33
	} else {
		goto L288
	}
L288:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_plpgsql_compile_callback_8))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L10
	} else {
		goto L289
	}
L289:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L10
	} else {
		goto L290
	}
L290:
	;
	v1274 = F_format_type_be(m, v1244)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L10
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1274
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_42), v23+int32(80))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L10
	} else {
		goto L292
	}
L292:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(460), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
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
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+61)) = uint8(v1287)
	v1290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+62)) = uint8(base.B2i32(v1290 == int32(100)))
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1258)+78)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+60)) = uint8(v1294)
	v1296 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1258)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+56)) = v1296
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v27)+108))
	if v1298 <= int32(3830) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	F_ReleaseCatCache(m, v1250)
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L10
	} else {
		goto L309
	}
L296:
	;
	if v578 != 0 {
		goto L295
	} else {
		goto L306
	}
L297:
	;
	if v1298 != int32(_a_F_plpgsql_compile_callback_43) {
		goto L295
	} else {
		goto L305
	}
L298:
	;
	switch v1298 - int32(2277) {
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
	if base.B2i32(v1298 == int32(3831))|base.B2i32(base.Ui32(v1298-int32(_a_F_plpgsql_compile_callback_17)) < base.Ui32(int32(4)))|base.B2i32(v1298 == int32(_a_F_plpgsql_compile_callback_18)) != 0 {
		goto L296
	} else {
		goto L304
	}
L301:
	;
	if v1298 == int32(2776) {
		goto L296
	} else {
		goto L302
	}
L302:
	;
	if v1298 != int32(3500) {
		goto L297
	} else {
		goto L303
	}
L303:
	;
	goto L296
L304:
	;
	goto L297
L305:
	;
	goto L296
L306:
	;
	v1320 = int32(0)
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
	v1324 = F_build_datatype(m, v1250, int32(-1), v1322, v1320)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L10
	} else {
		goto L307
	}
L307:
	;
	v1327 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_44), v1320, v1324, int32(1))
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L10
	} else {
		goto L308
	}
L308:
	;
	goto L295
L309:
	;
	v1334 = base.B2i32(int32(0) < v578)
	v1339 = v178
	goto L35
L310:
	;
	if v1359 == int32(0) {
		goto L32
	} else {
		goto L311
	}
L311:
	;
	v1364 = int32(0)
	v1366 = F_build_datatype(m, v1359, int32(-1), v1364, v1364)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L10
	} else {
		goto L312
	}
L312:
	;
	F_ReleaseCatCache(m, v1359)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L10
	} else {
		goto L313
	}
L313:
	;
	v1373 = F_plpgsql_build_variable(m, int32(_a_F_plpgsql_compile_callback_45), int32(0), v1366, int32(1))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L10
	} else {
		goto L314
	}
L314:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1373)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+476)) = v1375
	v1379 = F_plpgsql_yyparse(m, l3+int32(516), v50)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L10
	} else {
		goto L315
	}
L315:
	;
	if v1379 != 0 {
		goto L31
	} else {
		goto L316
	}
L316:
	;
	F_scanner_finish(m, v50)
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L10
	} else {
		goto L317
	}
L317:
	;
	F_pfree(m, v48)
	mBase = m.M
	v1384 = m.ExcPending
	if v1384 != 0 {
		goto L10
	} else {
		goto L318
	}
L318:
	;
	if v1334 != 0 {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1393 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+68)) = v1393
	if v1393 <= int32(0) {
		goto L325
	} else {
		goto L326
	}
L320:
	;
	F_add_dummy_return(m, l3)
	mBase = m.M
	v1392 = m.ExcPending
	if v1392 != 0 {
		goto L10
	} else {
		goto L324
	}
L321:
	;
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
	if v1385 == int32(2278) {
		goto L320
	} else {
		goto L322
	}
L322:
	;
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+63)))
	if v1388 != int32(1) {
		goto L319
	} else {
		goto L323
	}
L323:
	;
	goto L320
L324:
	;
	goto L319
L325:
	;
	v1533 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+504)) = v1533
	v1536 = F_palloc_mul(m, int32(4), v1533)
	mBase = m.M
	v1537 = m.ExcPending
	if v1537 != 0 {
		goto L10
	} else {
		goto L337
	}
L326:
	;
	v1398 = l3 + int32(72)
	v1399 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1393) {
		goto L327
	} else {
		goto L328
	}
L327:
	;
	v1408 = v1399
	v1415 = int32(0)
	goto L330
L328:
	;
	v1461 = v1399
	goto L329
L329:
	;
	v1483 = v1461
	v1489 = v1399
	goto L334
L330:
	;
	v1427 = v1408 << (uint(int32(2)) % 32)
	v1430 = *(*int32)(unsafe.Add(mBase, uint32(v1427+v1339)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1427))) = v1430
	v1432 = int32(4)
	v1433 = v1427 | v1432
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1339+v1433)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1433))) = v1436
	v1439 = v1427 | int32(8)
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1339+v1439)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1439))) = v1442
	v1445 = v1427 | int32(12)
	v1448 = *(*int32)(unsafe.Add(mBase, uint32(v1445+v1339)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1445))) = v1448
	v1451 = v1408 + v1432
	v1453 = v1415 + v1432
	if v1453 != v1393&int32(_a_F_plpgsql_compile_callback_46) {
		v1408 = v1451
		v1415 = v1453
		goto L330
	} else {
		goto L332
	}
L331:
	;
	if v1393&int32(3) == int32(0) {
		goto L325
	} else {
		goto L333
	}
L332:
	;
	goto L331
L333:
	;
	v1461 = v1451
	goto L329
L334:
	;
	v1502 = v1483 << (uint(int32(2)) % 32)
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1502+v1339)))
	*(*int32)(unsafe.Add(mBase, uint32(v1398+v1502))) = v1505
	v1507 = int32(1)
	v1510 = v1489 + v1507
	if v1510 != v1393&int32(3) {
		v1483 = v1483 + v1507
		v1489 = v1510
		goto L334
	} else {
		goto L336
	}
L335:
	;
	goto L325
L336:
	;
	goto L335
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+508)) = v1536
	v1540 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[13]))
	if v1540 <= int32(0) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+512)) = v1639
	v1658 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+525)))
	if v1658 == int32(1) {
		goto L357
	} else {
		goto L358
	}
L339:
	;
	v1639 = int32(0)
	goto L338
L340:
	;
	goto L341
L341:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[15]))
	v1546 = int32(0)
	if v1540 != int32(1) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1557 = v1546
	v1561 = v1546
	v1564 = int32(0)
	goto L345
L343:
	;
	v1608 = v1546
	v1612 = v1546
	goto L344
L344:
	;
	v1627 = v1612 << (uint(int32(2)) % 32)
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1627+v1545)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536+v1627))) = v1630
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v1630)))
	switch v1632 {
	case 0, 4:
		goto L355
	default:
		v1639 = v1608
		goto L338
	case 2:
		goto L356
	}
L345:
	;
	v1576 = v1561 << (uint(int32(2)) % 32)
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1576+v1545)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536+v1576))) = v1579
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1579)))
	switch v1581 {
	case 0, 4:
		goto L349
	default:
		v1586 = v1557
		goto L347
	case 2:
		goto L348
	}
L346:
	;
	if v1540&int32(1) == int32(0) {
		v1639 = v1598
		goto L338
	} else {
		goto L354
	}
L347:
	;
	v1588 = v1576 | int32(4)
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v1588+v1545)))
	*(*int32)(unsafe.Add(mBase, uint32(v1536+v1588))) = v1591
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1591)))
	switch v1593 {
	case 0, 4:
		goto L351
	default:
		v1598 = v1586
		goto L350
	case 2:
		goto L352
	}
L348:
	;
	v1586 = v1557 + int32(40)
	goto L347
L349:
	;
	v1586 = v1557 + int32(56)
	goto L347
L350:
	;
	v1599 = int32(2)
	v1600 = v1561 + v1599
	v1602 = v1564 + v1599
	if v1602 != v1540&int32(2147483646) {
		v1557 = v1598
		v1561 = v1600
		v1564 = v1602
		goto L345
	} else {
		goto L353
	}
L351:
	;
	v1598 = v1586 + int32(56)
	goto L350
L352:
	;
	v1598 = v1586 + int32(40)
	goto L350
L353:
	;
	goto L346
L354:
	;
	v1608 = v1598
	v1612 = v1600
	goto L344
L355:
	;
	v1639 = v1608 + int32(56)
	goto L338
L356:
	;
	v1639 = v1608 + int32(40)
	goto L338
L357:
	;
	F_plpgsql_mark_local_assignment_targets(m, l3)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L10
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v1664 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[12])))
	if v1664 == int32(1) {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	goto L359
L361:
	;
	F_plpgsql_dumptree(m, l3)
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L10
	} else {
		goto L364
	}
L362:
	;
	goto L363
L363:
	;
	v1670 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[16]))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	if v1674 != v1670 {
		goto L366
	} else {
		goto L367
	}
L364:
	;
	goto L363
L365:
	;
	v1704 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[0])) = v1704
	*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[1])) = uint8(v1704)
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v23)+396))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[3])) = v1710
	v1712 = int32(_a_F_plpgsql_compile_callback_47)
	v1713 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[5])) = v1704
	*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_callback[4])) = v1713
	m.G0 = v23 + int32(416)
	return
L366:
	;
	if v1674 == int32(0) {
		goto L369
	} else {
		goto L370
	}
L367:
	;
	goto L368
L368:
	;
	goto L365
L369:
	;
	if v1670 != 0 {
		goto L376
	} else {
		goto L377
	}
L370:
	;
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v88)+28))
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v88)+24))
	if v1679 != 0 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	if v1678 == int32(0) {
		goto L369
	} else {
		goto L375
	}
L372:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1679)+28)) = v1678
	goto L371
L373:
	;
	goto L374
L374:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1674)+20)) = v1678
	goto L371
L375:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v88)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1678)+24)) = v1684
	goto L369
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v1670
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1670)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v1691
	if v1691 != 0 {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	goto L378
L378:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = int32(0)
	goto L368
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1691)+24)) = v88
	goto L381
L380:
	;
	goto L381
L381:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1670)+20)) = v88
	goto L365
L382:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v1244
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23-int32(-64))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L10
	} else {
		goto L383
	}
L383:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(442), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L10
	} else {
		goto L384
	}
L384:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L385:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L10
	} else {
		goto L386
	}
L386:
	;
	F_errmsg(m, int32(_a_F_plpgsql_compile_callback_48), int32(0))
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L10
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(455), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1752 = m.ExcPending
	if v1752 != 0 {
		goto L10
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+96)) = int32(16)
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_36), v23+int32(96))
	mBase = m.M
	v1763 = m.ExcPending
	if v1763 != 0 {
		goto L10
	} else {
		goto L390
	}
L390:
	;
	goto L30
L391:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v1379
	F_errmsg_internal(m, int32(_a_F_plpgsql_compile_callback_49), v23+int32(112))
	mBase = m.M
	v1773 = m.ExcPending
	if v1773 != 0 {
		goto L10
	} else {
		goto L392
	}
L392:
	;
	F_errfinish(m, int32(_a_F_plpgsql_compile_callback_6), int32(684), int32(_a_F_plpgsql_compile_callback_21))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L10
	} else {
		goto L393
	}
L393:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L394:
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
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
							v115 = F_SearchSysCache1(m, int32(82), int64(16))
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
														v289 = m.ExcPending
														if v289 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v133
															F_errmsg_internal(m, int32(_a_F_plpgsql_compile_inline_10), v14+int32(16))
															mBase = m.M
															v295 = m.ExcPending
															if v295 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_plpgsql_compile_inline_11), int32(844), int32(_a_F_plpgsql_compile_inline_12))
																mBase = m.M
																v300 = m.ExcPending
																if v300 != 0 {
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
																	v149 = F_palloc_mul(m, int32(4), v146)
																	mBase = m.M
																	v150 = m.ExcPending
																	if v150 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v149
																		v153 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																		if v153 <= int32(0) {
																			v231 = v142
																		} else {
																			v157 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[14]))
																			v158 = int32(0)
																			if v153 != int32(1) {
																				v166 = v158
																				v167 = v142
																				v175 = int32(0)
																				for {
																					v177 = v166 << (uint(int32(2)) % 32)
																					v180 = *(*int32)(unsafe.Add(mBase, uint32(v177+v157)))
																					*(*int32)(unsafe.Add(mBase, uint32(v149+v177))) = v180
																					v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
																					switch v182 {
																					case 0, 4:
																						v187 = v167 + int32(56)
																					default:
																						v187 = v167
																					case 2:
																						v187 = v167 + int32(40)
																					}
																					v189 = v177 | int32(4)
																					v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v157)))
																					*(*int32)(unsafe.Add(mBase, uint32(v149+v189))) = v192
																					v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
																					switch v194 {
																					case 0, 4:
																						v199 = v187 + int32(56)
																					default:
																						v199 = v187
																					case 2:
																						v199 = v187 + int32(40)
																					}
																					v200 = int32(2)
																					v201 = v166 + v200
																					v203 = v175 + v200
																					if v203 != v153&int32(2147483646) {
																						v166 = v201
																						v167 = v199
																						v175 = v203
																						continue
																					} else {
																						break
																					}
																					break
																				}
																				if v153&int32(1) == int32(0) {
																					v231 = v199
																				} else {
																					v208 = v201
																					v209 = v199
																					v219 = v208 << (uint(int32(2)) % 32)
																					v222 = *(*int32)(unsafe.Add(mBase, uint32(v219+v157)))
																					*(*int32)(unsafe.Add(mBase, uint32(v149+v219))) = v222
																					v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
																					switch v224 {
																					case 0, 4:
																						v231 = v209 + int32(56)
																					default:
																						v231 = v209
																					case 2:
																						v231 = v209 + int32(40)
																					}
																				}
																			} else {
																				v208 = v158
																				v209 = v142
																				v219 = v208 << (uint(int32(2)) % 32)
																				v222 = *(*int32)(unsafe.Add(mBase, uint32(v219+v157)))
																				*(*int32)(unsafe.Add(mBase, uint32(v149+v219))) = v222
																				v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
																				switch v224 {
																				case 0, 4:
																					v231 = v209 + int32(56)
																				default:
																					v231 = v209
																				case 2:
																					v231 = v209 + int32(40)
																				}
																			}
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v231
																		v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																		if v241 == int32(1) {
																			F_plpgsql_mark_local_assignment_targets(m, v43)
																			mBase = m.M
																			v245 = m.ExcPending
																			if v245 != 0 {
																				return int32(0)
																			} else {
																				v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																				if v247 == int32(1) {
																					F_plpgsql_dumptree(m, v43)
																					mBase = m.M
																					v251 = m.ExcPending
																					if v251 != 0 {
																						return int32(0)
																					} else {
																						v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																						v256 = int32(0)
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																						*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																						v261 = int32(_a_F_plpgsql_compile_inline_13)
																						v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																						*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																						m.G0 = v14 + int32(48)
																						return v43
																					}
																				} else {
																					v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																					v256 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																					v261 = int32(_a_F_plpgsql_compile_inline_13)
																					v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			}
																		} else {
																			v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																			if v247 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v251 = m.ExcPending
																				if v251 != 0 {
																					return int32(0)
																				} else {
																					v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																					v256 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																					v261 = int32(_a_F_plpgsql_compile_inline_13)
																					v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																				v256 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																				v261 = int32(_a_F_plpgsql_compile_inline_13)
																				v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
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
																v149 = F_palloc_mul(m, int32(4), v146)
																mBase = m.M
																v150 = m.ExcPending
																if v150 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+508)) = v149
																	v153 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[12]))
																	if v153 <= int32(0) {
																		v231 = v142
																	} else {
																		v157 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[14]))
																		v158 = int32(0)
																		if v153 != int32(1) {
																			v166 = v158
																			v167 = v142
																			v175 = int32(0)
																			for {
																				v177 = v166 << (uint(int32(2)) % 32)
																				v180 = *(*int32)(unsafe.Add(mBase, uint32(v177+v157)))
																				*(*int32)(unsafe.Add(mBase, uint32(v149+v177))) = v180
																				v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
																				switch v182 {
																				case 0, 4:
																					v187 = v167 + int32(56)
																				default:
																					v187 = v167
																				case 2:
																					v187 = v167 + int32(40)
																				}
																				v189 = v177 | int32(4)
																				v192 = *(*int32)(unsafe.Add(mBase, uint32(v189+v157)))
																				*(*int32)(unsafe.Add(mBase, uint32(v149+v189))) = v192
																				v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
																				switch v194 {
																				case 0, 4:
																					v199 = v187 + int32(56)
																				default:
																					v199 = v187
																				case 2:
																					v199 = v187 + int32(40)
																				}
																				v200 = int32(2)
																				v201 = v166 + v200
																				v203 = v175 + v200
																				if v203 != v153&int32(2147483646) {
																					v166 = v201
																					v167 = v199
																					v175 = v203
																					continue
																				} else {
																					break
																				}
																				break
																			}
																			if v153&int32(1) == int32(0) {
																				v231 = v199
																			} else {
																				v208 = v201
																				v209 = v199
																				v219 = v208 << (uint(int32(2)) % 32)
																				v222 = *(*int32)(unsafe.Add(mBase, uint32(v219+v157)))
																				*(*int32)(unsafe.Add(mBase, uint32(v149+v219))) = v222
																				v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
																				switch v224 {
																				case 0, 4:
																					v231 = v209 + int32(56)
																				default:
																					v231 = v209
																				case 2:
																					v231 = v209 + int32(40)
																				}
																			}
																		} else {
																			v208 = v158
																			v209 = v142
																			v219 = v208 << (uint(int32(2)) % 32)
																			v222 = *(*int32)(unsafe.Add(mBase, uint32(v219+v157)))
																			*(*int32)(unsafe.Add(mBase, uint32(v149+v219))) = v222
																			v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
																			switch v224 {
																			case 0, 4:
																				v231 = v209 + int32(56)
																			default:
																				v231 = v209
																			case 2:
																				v231 = v209 + int32(40)
																			}
																		}
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v43)+512)) = v231
																	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+525)))
																	if v241 == int32(1) {
																		F_plpgsql_mark_local_assignment_targets(m, v43)
																		mBase = m.M
																		v245 = m.ExcPending
																		if v245 != 0 {
																			return int32(0)
																		} else {
																			v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																			if v247 == int32(1) {
																				F_plpgsql_dumptree(m, v43)
																				mBase = m.M
																				v251 = m.ExcPending
																				if v251 != 0 {
																					return int32(0)
																				} else {
																					v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																					v256 = int32(0)
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																					*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																					v261 = int32(_a_F_plpgsql_compile_inline_13)
																					v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																					*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																					m.G0 = v14 + int32(48)
																					return v43
																				}
																			} else {
																				v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																				v256 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																				v261 = int32(_a_F_plpgsql_compile_inline_13)
																				v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		}
																	} else {
																		v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[11])))
																		if v247 == int32(1) {
																			F_plpgsql_dumptree(m, v43)
																			mBase = m.M
																			v251 = m.ExcPending
																			if v251 != 0 {
																				return int32(0)
																			} else {
																				v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																				v256 = int32(0)
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																				*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																				v261 = int32(_a_F_plpgsql_compile_inline_13)
																				v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																				*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
																				m.G0 = v14 + int32(48)
																				return v43
																			}
																		} else {
																			v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[3])) = v253
																			v256 = int32(0)
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[0])) = v256
																			*(*uint8)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[2])) = uint8(v256)
																			v261 = int32(_a_F_plpgsql_compile_inline_13)
																			v262 = *(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6]))
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[6])) = v256
																			*(*int32)(unsafe.Add(mBase, _c_F_plpgsql_compile_inline[5])) = v262
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
									v275 = m.ExcPending
									if v275 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(16)
										F_errmsg_internal(m, int32(_a_F_plpgsql_compile_inline_14), v14)
										mBase = m.M
										v280 = m.ExcPending
										if v280 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_plpgsql_compile_inline_11), int32(1983), int32(_a_F_plpgsql_compile_inline_15))
											mBase = m.M
											v285 = m.ExcPending
											if v285 != 0 {
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
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+192))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
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
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
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
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+76))
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
	F_errfinish(m, int32(_a_F_plpgsql_parse_err_condition_3), int32(2213), int32(_a_F_plpgsql_parse_err_condition_4))
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
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
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
			*(*int32)(unsafe.Add(mBase, uint32(v16+v17<<(uint(int32(2))%32))+72)) = v12
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
			v40 = v16 + v37*int32(24)
			v41 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+104)) = v41
			v43 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+96)) = v43
			v45 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v40)+88)) = v45
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+68))
			*(*int32)(unsafe.Add(mBase, uint32(v47)+68)) = v48 + int32(1)
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
	var v49 int32
	_ = v49
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
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
								v49 = F_errdetail(m, int32(_a_F_plpgsql_post_column_ref_2), int32(0))
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
										F_errfinish(m, int32(_a_F_plpgsql_post_column_ref_3), int32(1050), int32(_a_F_plpgsql_post_column_ref_4))
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
