package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecIndexEvalRuntimeKeys(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(_a_F_ExecIndexEvalRuntimeKeys_0)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0])) = v18
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v4
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecIndexEvalRuntimeKeys[0])) = v16
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v34 = l1 + v27*int32(12)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v40 = m.T0[v39].(func(*base.Module, int32, int32, int32) int64)(m, v36, l0, v13+int32(15))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v42 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v61
	v64 = v27 + int32(1)
	if v64 != l2 {
		v27 = v64
		goto L4
	} else {
		goto L16
	}
L9:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v40
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v61 = v46 | int32(1)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	if v49 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v53 = F_pg_detoast_datum(m, base.I32_wrap_i64(v40))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v56 = v40
	goto L14
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v61 = v58 & int32(-2)
	goto L8
L15:
	;
	v56 = base.I64_extend_i32_u(v53)
	goto L14
L16:
	;
	goto L5
}
func F_IndexAmTranslateCompareType(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	if base.B2i32(l1 != int32(403))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(l0-int32(1))) == v5 {
		v45 = l0
		m.G0 = v9 + int32(16)
		return v45 & int32(_a_F_IndexAmTranslateCompareType_0)
	} else {
		v21 = F_GetIndexAmRoutineByAmId(m, l1, int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+136))
			if v25 != 0 {
				v26 = m.T0[v25].(func(*base.Module, int32, int32) int32)(m, l0, l2)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v26
					if l3|v28 != 0 {
						v45 = v28
						m.G0 = v9 + int32(16)
						return v45 & int32(_a_F_IndexAmTranslateCompareType_0)
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
							F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v9)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(178), int32(_a_F_IndexAmTranslateCompareType_3))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
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
			} else {
				v28 = v5
				if l3|v28 != 0 {
					v45 = v28
					m.G0 = v9 + int32(16)
					return v45 & int32(_a_F_IndexAmTranslateCompareType_0)
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg_internal(m, int32(_a_F_IndexAmTranslateCompareType_1), v9)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_IndexAmTranslateCompareType_2), int32(178), int32(_a_F_IndexAmTranslateCompareType_3))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
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
func F_IndexOnlyNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 float64
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 float64
	_ = v249
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = v21 * v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v27 != 0 {
		v58 = v27
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v59 = F_index_getnext_tid(m, v58, v24)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L3
	} else {
		goto L18
	}
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v36 = F_ScanRelIsReadOnly(m, l0)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v40 = int32(1024)
	goto L7
L6:
	;
	v40 = int32(0)
	goto L7
L7:
	;
	v41 = F_index_beginscan(m, v28, v29, v30, v31, v32, v33, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v41
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+28)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+144)))
	if v49 != int32(1) {
		v58 = v41
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_index_rescan(m, v41, v52, v53, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L3
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	v58 = v41
	goto L1
L14:
	;
	m.G0 = v18 + int32(16)
	return v25
L15:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v317 <= int32(0) {
		goto L80
	} else {
		goto L81
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L3
	} else {
		goto L77
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L3
	} else {
		goto L74
	}
L18:
	;
	if v59 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = v59
	goto L22
L20:
	;
	goto L21
L21:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v285)+12))
	m.T0[v286].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L3
	} else {
		goto L73
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[0]))
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L21
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L3
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+2)))
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
	v88 = F_visibilitymap_get_status(m, v82, v83|v84<<(uint(int32(16))%32), l0+int32(172))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L3
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v268 = F_index_getnext_tid(m, v58, v24)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L71
	}
L29:
	;
	v91 = v88 & int32(1)
	if v91 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v94 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+12))
	m.T0[v114].(func(*base.Module, int32))(m, v25)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L3
	} else {
		goto L40
	}
L33:
	;
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v94)+408))
	*(*float64)(unsafe.Add(mBase, uint32(v94)+408)) = base.F64_add(v95, float64(1))
	goto L35
L34:
	;
	goto L35
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v100 = F_index_fetch_heap(m, v58, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	if v100 == int32(0) {
		goto L28
	} else {
		goto L37
	}
L37:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+12))
	m.T0[v106].(func(*base.Module, int32))(m, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+66)))
	if v109 == int32(1) {
		goto L17
	} else {
		goto L39
	}
L39:
	;
	goto L32
L40:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	if v117 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)))
	v214 = v212 & int32(_a_F_IndexOnlyNext_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v214)
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+6)) = uint16(v217)
	goto L61
L42:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	F_heap_deform_tuple(m, v117, v118, v119, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L3
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v58)+44))
	if v123 == int32(0) {
		goto L16
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v58)+48))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+6)))
	if int32(0) <= base.I32_extend16_s(v131) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v142 == int32(0) {
		goto L41
	} else {
		goto L51
	}
L48:
	;
	v135 = int32(8)
	goto L50
L49:
	;
	v135 = int32(16)
	goto L50
L50:
	;
	F_index_deform_tuple_internal(m, v126, v127, v128, v123+v135, v123+int32(8), int32(base.Ui32(v131)>>(uint(int32(15))%32)))
	mBase = m.M
	goto L47
L51:
	;
	v145 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	if v146 <= v145 {
		goto L41
	} else {
		goto L52
	}
L52:
	;
	v151 = v145
	goto L53
L53:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v164+v151<<(uint(int32(1))%32)))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168+v169))))
	if v171 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L41
L55:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
	v177 = F_MemoryContextAlloc(m, v175, int32(64))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L3
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v195 = v151 + int32(1)
	if v195 != v146 {
		v151 = v195
		goto L53
	} else {
		goto L60
	}
L58:
	;
	v180 = v168 << (uint(int32(3)) % 32)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180+v181)))
	v185 = F_strncpy(m, v177, v183, int32(64))
	mBase = m.M
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v185)+63)) = uint8(v186)
	goto L59
L59:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v188+v180))) = base.I64_extend_i32_u(v177)
	goto L57
L60:
	;
	goto L54
L61:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+72)))
	if v219 != int32(1) {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v25
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	if v223 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	F_MemoryContextReset(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L3
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v229 = int32(_a_F_IndexOnlyNext_1)
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1])) = v232
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v223)+24))
	v237 = m.T0[v236].(func(*base.Module, int32, int32, int32) int64)(m, v223, v26, v18+int32(15))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L3
	} else {
		goto L67
	}
L66:
	;
	goto L15
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IndexOnlyNext[1])) = v230
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	F_MemoryContextReset(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L68
	}
L68:
	;
	if v237 != int64(0) {
		goto L15
	} else {
		goto L69
	}
L69:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v246 == int32(0) {
		goto L28
	} else {
		goto L70
	}
L70:
	;
	v249 = *(*float64)(unsafe.Add(mBase, uint32(v246)+432))
	*(*float64)(unsafe.Add(mBase, uint32(v246)+432)) = base.F64_add(v249, float64(1))
	goto L28
L71:
	;
	if v268 != 0 {
		v67 = v268
		goto L22
	} else {
		goto L72
	}
L72:
	;
	goto L23
L73:
	;
	goto L14
L74:
	;
	F_errmsg_internal(m, int32(_a_F_IndexOnlyNext_2), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L3
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(185), int32(_a_F_IndexOnlyNext_4))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errmsg_internal(m, int32(_a_F_IndexOnlyNext_5), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(320), int32(_a_F_IndexOnlyNext_6))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	if v91 == int32(0) {
		goto L14
	} else {
		goto L87
	}
L81:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+84)))
	if v320 != int32(1) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L83
	}
L83:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_IndexOnlyNext_7), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L3
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_IndexOnlyNext_3), int32(225), int32(_a_F_IndexOnlyNext_4))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L3
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+2)))
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67))))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	F_PredicateLockPage(m, v341, v342|v343<<(uint(int32(16))%32), v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	goto L14
}
func F_IndexScanOK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v3 - int32(1) {
	case 0, 1:
		v13 = v2
	default:
		v13 = int32(1)
	case 7, 9, 10, 20, 42, 43:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_IndexScanOK[0])))
		if v9 != int32(1) {
			v13 = v2
		} else {
			v13 = int32(1)
		}
	case 33:
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_IndexScanOK[1])))
		if v7 != 0 {
			v13 = int32(1)
		} else {
			v13 = v2
		}
	}
	return v13
}
func F_IsIndexCompatibleAsArbiter(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+12)))
	if v10 != v12 {
		v104 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v104
L2:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v14 != 0 {
		v104 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v15 != 0 {
		v104 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
	if v16 != v17 {
		v104 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+13)))
	if v19 != v20 {
		v104 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+10)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+10)))
	if v22 != v23 {
		v104 = v3
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v25 = base.I32_extend16_s(v22)
	if v25 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v81 = F_RelationGetIndexExpressions(m, l0)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v28 = int32(48)
	v36 = int32(0)
	goto L10
L10:
	;
	v42 = v36 << (uint(int32(1)) % 32)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+v28+v42))))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42+(v11+v28)))))
	if v44 != v46 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	return int32(0)
L12:
	;
	return int32(0)
L13:
	;
	goto L14
L14:
	;
	v51 = v36 << (uint(int32(2)) % 32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55+v51)))
	if v54 != v57 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	goto L17
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61+v51)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+208))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v64+v51)))
	if v63 == v66 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v69 = v36 + int32(1)
	if v69 == v25 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L11
L21:
	;
	v36 = v69
	goto L10
L22:
	;
	return int32(0)
L23:
	;
	v85 = F_RelationGetIndexExpressions(m, l1)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v87 = F_equal(m, v81, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	if v87 == int32(0) {
		v104 = v3
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v91 = F_RelationGetIndexPredicate(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v93 = F_RelationGetIndexPredicate(m, l1)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v95 = F_equal(m, v91, v93)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v104 = v95
	goto L1
}
func F_build_index_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	v4 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v12 == v4 {
		v222 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v222
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v15 == int32(0) {
		v222 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v18 <= int32(0) {
		v222 = v4
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = int32(0)
	v30 = v4
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v35 <= v26 {
		v222 = v30
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v222 = v207
	goto L1
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v26))))
	v41 = v26 << (uint(int32(2)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41+v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if base.B2i32(l2 != int32(-1)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60+v41)))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v63+v41)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v66+v41)))
	v69 = int32(1)
	v73 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v77 = F_make_pathkey_from_sortinfo(m, l0, v45, v62, v65, v68, v58&v69, v59&v69, v73, v75, v73)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v48 = int32(1)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50+v26))))
	v58 = v39 ^ v48
	v59 = v52 ^ v48
	goto L8
L10:
	;
	goto L11
L11:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55+v26))))
	v58 = v39
	v59 = v57
	goto L8
L12:
	;
	v213 = v26 + int32(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v213 < v214 {
		v26 = v213
		v30 = v207
		goto L5
	} else {
		goto L47
	}
L13:
	;
	return int32(0)
L14:
	;
	if v77 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+40)))
	if v82 != 0 {
		v207 = v30
		goto L12
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v26<<(uint(int32(2))%32))))
	if base.Ui32(v127) <= base.Ui32(int32(_a_F_build_index_pathkeys_0)) {
		goto L29
	} else {
		goto L30
	}
L18:
	;
	if v30 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v121 = F_lappend(m, v30, v77)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L26
	}
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	if v85 <= int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	v93 = int32(0)
	goto L22
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v88+v93<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v81 == v105 {
		v207 = v30
		goto L12
	} else {
		goto L24
	}
L23:
	;
	goto L19
L24:
	;
	v108 = v93 + int32(1)
	if v85 != v108 {
		v93 = v108
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v207 = v121
	goto L12
L27:
	;
	if v198 == int32(0) {
		v222 = v30
		goto L1
	} else {
		goto L46
	}
L28:
	;
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+204))
	if v142 == v140 {
		v198 = v140
		goto L27
	} else {
		goto L35
	}
L29:
	;
	if base.B2i32(v127 == int32(424))|base.B2i32(v127 == int32(2222)) != 0 {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v137 = F_op_in_opfamily(m, int32(91), v127)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v198 = int32(0)
	goto L27
L33:
	;
	if v137 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v198 = int32(0)
	goto L27
L35:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if int32(0) < v145 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v152 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	v198 = int32(0)
	goto L27
L39:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v142)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160+v152<<(uint(int32(2))%32))))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+10)))
	if v165 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v172 = v152 + int32(1)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v172 < v173 {
		v152 = v172
		goto L39
	} else {
		goto L45
	}
L42:
	;
	v166 = F_match_boolean_index_clause(m, l0, v164, v26, l1)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	if v166 == int32(0) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v198 = int32(1)
	goto L27
L45:
	;
	goto L40
L46:
	;
	v207 = v30
	goto L12
L47:
	;
	goto L6
}
func F_get_index_constraint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(176)
	m.G0 = v8
	v12 = F_table_open(m, int32(2608), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_ScanKeyInit(m, v8, int32(1), int32(3), int32(184), int64(1259))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v8+int32(56), int32(2), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v32 = int32(3)
	F_ScanKeyInit(m, v8+int32(112), v32, v32, int32(65), int64(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v42 = F_systable_beginscan(m, v12, int32(2673), int32(1), int32(0), int32(3), v8)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v44 == int32(0) {
		v70 = v2
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v48 = v44
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v56 != int32(2606) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v70 = v2
	goto L6
L12:
	;
	v64 = F_systable_getnext(m, v42)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L16
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v59 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+24)))
	if v60 != int32(105) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v70 = v63
	goto L6
L16:
	;
	if v64 != 0 {
		v48 = v64
		goto L10
	} else {
		goto L17
	}
L17:
	;
	goto L11
L18:
	;
	F_relation_close(m, v12, int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	m.G0 = v8 + int32(176)
	return v70
}
func F_get_index_isclustered(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_SearchSysCache1(m, int32(34), base.I64_extend_i32_u(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_get_index_isclustered_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_get_index_isclustered_1), int32(3943), int32(_a_F_get_index_isclustered_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29+v30)+17)))
			F_ReleaseCatCache(m, v10)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_index_am_handler_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_index_am_handler_out_0), int32(371), int32(_a_F_index_am_handler_out_1), int32(_a_F_index_am_handler_out_2), int32(_a_F_index_am_handler_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_index_beginscan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v14 != int32(5) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L20
	} else {
		goto L31
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L20
	} else {
		goto L27
	}
L3:
	;
	v49 = int32(0)
	v51 = F_index_beginscan_internal(m, l1, l4, l5, l2, v49, v49)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_index_beginscan[2]))
	if v18 <= int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_beginscan[3])))
	if v22&int32(1) == int32(0) {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+118)))
	if v28 != int32(112) {
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L7
L9:
	;
	if v18 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v33 != 0 {
		goto L2
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L15
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v34 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	if base.Ui32(v35) < base.Ui32(int32(_a_F_index_beginscan_6)) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v38 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+119)))
	switch v42 - int32(109) {
	case 0, 5:
		goto L18
	default:
		goto L2
	}
L18:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+112)))
	if v45 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	goto L3
L20:
	;
	return int32(0)
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+40)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v51))) = l0
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_index_beginscan[0]))
	if v59 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_index_beginscan[1])))
	if v61&int32(1) == int32(0) {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+44))
	v68 = m.T0[v67].(func(*base.Module, int32, int32) int32)(m, l0, l6)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L20
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+68)) = v68
	m.G0 = v12 + int32(16)
	return v51
L27:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v83 + int32(4)
	F_errmsg(m, int32(_a_F_index_beginscan_3), v12)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_index_beginscan_4), int32(275), int32(_a_F_index_beginscan_5))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L20
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errmsg_internal(m, int32(_a_F_index_beginscan_0), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L20
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_index_beginscan_1), int32(1256), int32(_a_F_index_beginscan_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L20
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_create_copy(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v238 int32
	_ = v238
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v254 int32
	_ = v254
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v293 int64
	_ = v293
	var v301 int64
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v322 int64
	_ = v322
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	v25 = m.G0
	v27 = v25 - int32(48)
	m.G0 = v27
	v30 = F_index_open(m, l2, int32(3))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = F_BuildIndexInfo(m, v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = l1 & int32(8)
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L81
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L78
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L75
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L71
	}
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	if v38 != 0 {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v40 = base.I64_extend_i32_u(l2)
	v41 = F_SearchSysCache1(m, int32(34), v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L10
L12:
	;
	if v41 == int32(0) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v46)+16)))
	v51 = F_SysCacheGetAttrNotNull(m, int32(34), v41, int32(18))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v55 = F_SysCacheGetAttrNotNull(m, int32(34), v41, int32(19))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v58 = F_SearchSysCache1(m, int32(57), v40)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v58 == int32(0) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v66 = F_SysCacheGetAttr(m, int32(57), v58, int32(33), v27+int32(47))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	if v69 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v72 = F_SysCacheGetAttrNotNull(m, int32(34), v41, int32(20))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v81 = int32(0)
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v34)+84))
	if v83 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v75 = F_text_to_cstring(m, base.I32_wrap_i64(v72))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v77 = F_stringToNode(m, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_pfree(m, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v81 = v77
	goto L21
L26:
	;
	v86 = F_SysCacheGetAttrNotNull(m, int32(34), v41, int32(21))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v97 = int32(0)
	goto L28
L28:
	;
	v99 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v34)+132))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+116)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+117)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v30)+204))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+28)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+124)))
	v112 = F_makeIndexInfo(m, v100, v101, v102, v81, v97, v103, v104, base.B2i32(v37 == v99), base.B2i32(v37 != v99), v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L34
	}
L29:
	;
	v89 = F_text_to_cstring(m, base.I32_wrap_i64(v86))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v91 = F_stringToNode(m, v89)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v93 = F_make_ands_implicit(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v89)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v97 = v93
	goto L28
L34:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v30)+192))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+15)))
	if v115 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v34)+92))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+92)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v34)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+96)) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v34)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+100)) = v122
	goto L37
L36:
	;
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if int32(0) < v124 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v127 = int32(12)
	v138 = int32(0)
	v142 = v99
	goto L41
L39:
	;
	v188 = v99
	goto L40
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v204 = F_palloc0_mul(m, int32(8), v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L45
	}
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v166 = F_lappend(m, v142, v156+v157<<(uint(int32(3))%32)+v138*int32(100)+int32(32))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v188 = v166
	goto L40
L43:
	;
	v168 = int32(1)
	v169 = v138 << (uint(v168) % 32)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v34+v127+v169))))
	*(*uint16)(unsafe.Add(mBase, uint32(v112+v127+v169))) = uint16(v172)
	v175 = v138 + v168
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v175 < v176 {
		v138 = v175
		v142 = v166
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v206 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v206 < v207 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v216 = v206
	goto L49
L47:
	;
	v254 = v207
	goto L48
L48:
	;
	v270 = F_palloc0_mul(m, int32(16), v254)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L53
	}
L49:
	;
	v238 = v216 + int32(1)
	v240 = F_get_attoptions(m, l2, base.I32_extend16_s(v238))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v254 = v243
	goto L48
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v204+v216<<(uint(int32(3))%32)))) = v240
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if v238 < v243 {
		v216 = v238
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	if int32(0) < v272 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v293 = int64(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	v348 = int32(0)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v30)+48))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+84))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v30)+248))
	v356 = int32(24)
	if v48&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	v301 = v293 + int64(1)
	v303 = F_SearchSysCache2(m, int32(7), v40, base.I64_extend16_s(v301))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	goto L56
L59:
	;
	if v303 == int32(0) {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v311 = F_SysCacheGetAttr(m, int32(7), v303, int32(21), v27+int32(47))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_ReleaseCatCache(m, v303)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v318 = v270 + base.I32_wrap_i64(v293)<<(uint(int32(4))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v318))) = v311
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+47)))
	*(*uint8)(unsafe.Add(mBase, uint32(v318)+8)) = uint8(v320)
	v322 = int64(*(*int32)(unsafe.Add(mBase, uint32(v112)+4)))
	if v301 < v322 {
		v293 = v301
		goto L57
	} else {
		goto L63
	}
L63:
	;
	goto L58
L64:
	;
	v365 = l1
	goto L66
L65:
	;
	v365 = l1 | int32(256)
	goto L66
L66:
	;
	v366 = int32(0)
	v370 = F_index_create(m, l0, l4, v348, v348, v348, v348, v112, v188, v353, l3, v354, base.I32_wrap_i64(v51)+v356, v204, base.I32_wrap_i64(v55)+v356, v270, v66, v365, v366, int32(1), v366, v366)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	F_relation_close(m, v30, int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_ReleaseCatCache(m, v41)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_ReleaseCatCache(m, v58)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	m.G0 = v27 + int32(48)
	return v370
L71:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(_a_F_index_create_copy_0), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_index_create_copy_1), int32(1348), int32(_a_F_index_create_copy_2))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l2
	F_errmsg_internal(m, int32(_a_F_index_create_copy_3), v27)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_index_create_copy_1), int32(1353), int32(_a_F_index_create_copy_2))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = l2
	F_errmsg_internal(m, int32(_a_F_index_create_copy_4), v27+int32(16))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errfinish(m, int32(_a_F_index_create_copy_1), int32(1372), int32(_a_F_index_create_copy_2))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = l2
	*(*uint32)(unsafe.Add(mBase, uint32(v27)+32)) = uint32(v301)
	F_errmsg_internal(m, int32(_a_F_index_create_copy_5), v27+int32(32))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_index_create_copy_1), int32(1464), int32(_a_F_index_create_copy_2))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_index_deform_tuple_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v273 int64
	_ = v273
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v301 int32
	_ = v301
	var v312 int32
	_ = v312
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v426 int64
	_ = v426
	var v430 int32
	_ = v430
	v7 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v7
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v22 < v19 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v22
	goto L3
L2:
	;
	v24 = v19
	goto L3
L3:
	;
	if l5 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v84 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L5:
	;
	v83 = v19
	v84 = v24
	goto L4
L6:
	;
	goto L7
L7:
	;
	v28 = v19 >> (uint(int32(3)) % 32)
	if v28 <= int32(0) {
		v59 = v7
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v59))))
	v73 = base.I32_ctz(v67^int32(-1)) + v59<<(uint(int32(3))%32)
	if v73 < v19 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v38 = v7
	goto L10
L10:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v38))))
	if v46 != int32(255) {
		v59 = v38
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v59 = v28
	goto L8
L12:
	;
	v50 = v38 + int32(1)
	if v50 != v28 {
		v38 = v50
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v75 = v73
	goto L16
L15:
	;
	v75 = v19
	goto L16
L16:
	;
	if v24 < v75 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v77 = v24
	goto L19
L18:
	;
	v77 = v75
	goto L19
L19:
	;
	v83 = v75
	v84 = v77
	goto L4
L20:
	;
	if v150 < v83 {
		goto L35
	} else {
		goto L36
	}
L21:
	;
	v150 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v105 = int32(0)
	goto L24
L24:
	;
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v105))) = uint8(v113)
	v116 = v105 << (uint(int32(3)) % 32)
	v117 = l0 + int32(28) + v116
	v118 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117))))
	v119 = l3 + v118
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	if v121 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0+v84<<(uint(int32(3))%32))+22)))
	v141 = int32(*(*int16)(unsafe.Add(mBase, uint32(v117))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v140 + v141
	v150 = v84
	goto L20
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1+v116))) = v132
	v135 = v105 + int32(1)
	if v135 != v84 {
		v105 = v135
		goto L24
	} else {
		goto L34
	}
L27:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+2)))
	switch v124 - int32(1) {
	case 0:
		goto L31
	case 1:
		goto L32
	default:
		goto L30
	case 3:
		goto L33
	}
L28:
	;
	goto L29
L29:
	;
	v132 = base.I64_extend_i32_u(v119)
	goto L26
L30:
	;
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	v132 = v130
	goto L26
L31:
	;
	v129 = int64(*(*int8)(unsafe.Add(mBase, uint32(v119))))
	v132 = v129
	goto L26
L32:
	;
	v128 = int64(*(*int16)(unsafe.Add(mBase, uint32(v119))))
	v132 = v128
	goto L26
L33:
	;
	v127 = int64(*(*int32)(unsafe.Add(mBase, uint32(v119))))
	v132 = v127
	goto L26
L34:
	;
	goto L25
L35:
	;
	v167 = v150
	goto L38
L36:
	;
	v284 = v150
	goto L37
L37:
	;
	if v284 < v19 {
		goto L69
	} else {
		goto L70
	}
L38:
	;
	v176 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v167))) = uint8(v176)
	v179 = v167 << (uint(int32(3)) % 32)
	v182 = v17 + int32(12)
	v183 = v179 + (l0 + int32(28))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
	v185 = int32(*(*int16)(unsafe.Add(mBase, uint32(v183)+2)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+5)))
	if v176 < v185 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v284 = v83
	goto L37
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1+v179))) = v273
	v276 = v167 + int32(1)
	if v276 != v83 {
		v167 = v276
		goto L38
	} else {
		goto L68
	}
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v195 = (v186 + v189 - int32(1)) & (int32(0) - v186)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v195 + v185
	v198 = l3 + v195
	if v184 != 0 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	if v185 == int32(-1) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	switch v185 - int32(1) {
	case 0:
		goto L50
	case 1:
		goto L49
	default:
		goto L47
	case 3:
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v273 = base.I64_extend_i32_u(v198)
	goto L40
L47:
	;
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v198)))
	v273 = v204
	goto L40
L48:
	;
	v203 = int64(*(*int32)(unsafe.Add(mBase, uint32(v198))))
	v273 = v203
	goto L40
L49:
	;
	v202 = int64(*(*int16)(unsafe.Add(mBase, uint32(v198))))
	v273 = v202
	goto L40
L50:
	;
	v201 = int64(*(*int8)(unsafe.Add(mBase, uint32(v198))))
	v273 = v201
	goto L40
L51:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v206))))
	if v210&int32(1) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v256 = int32(1)
	v260 = (v206 + v186 - v256) & (int32(0) - v186)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v260
	v262 = l3 + v260
	v263 = F_strlen(m, v262)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v263 + v260 + v256
	v273 = base.I64_extend_i32_u(v262)
	goto L40
L54:
	;
	v220 = (v206 + v186 - int32(1)) & (int32(0) - v186)
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v220
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v220))))
	v224 = v220
	v225 = v223
	goto L56
L55:
	;
	v224 = v206
	v225 = v210
	goto L56
L56:
	;
	v226 = l3 + v224
	if v225 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182))) = v251 + v224
	v273 = base.I64_extend_i32_u(v226)
	goto L40
L58:
	;
	v230 = int32(18)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v226)+1)))
	if v232 == v230 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v243 = int32(1)
	if v225&v243 != 0 {
		v251 = int32(base.Ui32(v225) >> (uint(v243) % 32))
		goto L57
	} else {
		goto L67
	}
L61:
	;
	v235 = v230
	goto L63
L62:
	;
	v235 = int32(2)
	goto L63
L63:
	;
	if base.Ui32((v232-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v242 = int32(6)
	goto L66
L65:
	;
	v242 = v235
	goto L66
L66:
	;
	v251 = v242
	goto L57
L67:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	v251 = int32(base.Ui32(v247) >> (uint(int32(2)) % 32))
	goto L57
L68:
	;
	goto L39
L69:
	;
	v301 = v284
	goto L72
L70:
	;
	goto L71
L71:
	;
	m.G0 = v17 + int32(16)
	return
L72:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+int32(base.Ui32(v301)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v312)>>(uint(v301&int32(7))%32))&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L71
L74:
	;
	v430 = v301 + int32(1)
	if v430 != v19 {
		v301 = v430
		goto L72
	} else {
		goto L106
	}
L75:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1+v301<<(uint(int32(3))%32)))) = int64(0)
	v326 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v301))) = uint8(v326)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v329 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2+v301))) = uint8(v329)
	v332 = v301 << (uint(int32(3)) % 32)
	v335 = v17 + int32(12)
	v336 = v332 + (l0 + int32(28))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+4)))
	v338 = int32(*(*int16)(unsafe.Add(mBase, uint32(v336)+2)))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+5)))
	if v329 < v338 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1+v332))) = v426
	goto L74
L79:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v348 = (v339 + v342 - int32(1)) & (int32(0) - v339)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v348 + v338
	v351 = l3 + v348
	if v337 != 0 {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	if v338 == int32(-1) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	switch v338 - int32(1) {
	case 0:
		goto L88
	case 1:
		goto L87
	default:
		goto L85
	case 3:
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	v426 = base.I64_extend_i32_u(v351)
	goto L78
L85:
	;
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v351)))
	v426 = v357
	goto L78
L86:
	;
	v356 = int64(*(*int32)(unsafe.Add(mBase, uint32(v351))))
	v426 = v356
	goto L78
L87:
	;
	v355 = int64(*(*int16)(unsafe.Add(mBase, uint32(v351))))
	v426 = v355
	goto L78
L88:
	;
	v354 = int64(*(*int8)(unsafe.Add(mBase, uint32(v351))))
	v426 = v354
	goto L78
L89:
	;
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v359))))
	if v363&int32(1) == int32(0) {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	v409 = int32(1)
	v413 = (v359 + v339 - v409) & (int32(0) - v339)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v413
	v415 = l3 + v413
	v416 = F_strlen(m, v415)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v416 + v413 + v409
	v426 = base.I64_extend_i32_u(v415)
	goto L78
L92:
	;
	v373 = (v359 + v339 - int32(1)) & (int32(0) - v339)
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v373
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3+v373))))
	v377 = v373
	v378 = v376
	goto L94
L93:
	;
	v377 = v359
	v378 = v363
	goto L94
L94:
	;
	v379 = l3 + v377
	if v378 == int32(1) {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v335))) = v404 + v377
	v426 = base.I64_extend_i32_u(v379)
	goto L78
L96:
	;
	v383 = int32(18)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v379)+1)))
	if v385 == v383 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v396 = int32(1)
	if v378&v396 != 0 {
		v404 = int32(base.Ui32(v378) >> (uint(v396) % 32))
		goto L95
	} else {
		goto L105
	}
L99:
	;
	v388 = v383
	goto L101
L100:
	;
	v388 = int32(2)
	goto L101
L101:
	;
	if base.Ui32((v385-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v395 = int32(6)
	goto L104
L103:
	;
	v395 = v388
	goto L104
L104:
	;
	v404 = v395
	goto L95
L105:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	v404 = int32(base.Ui32(v400) >> (uint(int32(2)) % 32))
	goto L95
L106:
	;
	goto L73
}
func F_index_endscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+108))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
			m.T0[v15].(func(*base.Module, int32))(m, v12)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+204))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+108))
				v23 = v22
				m.T0[v23].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					F_RelationDecrementReferenceCount(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
						if v29 == int32(1) {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							F_UnregisterSnapshot(m, v32)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v35 != 0 {
									F_pfree(m, v35)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v38 != 0 {
											F_pfree(m, v38)
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return
											} else {
												F_pfree(m, l0)
												mBase = m.M
												v42 = m.ExcPending
												if v42 != 0 {
													return
												} else {
													m.G0 = v7 + int32(16)
													return
												}
											}
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									}
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v35 != 0 {
								F_pfree(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v23 = v11
			m.T0[v23].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				F_RelationDecrementReferenceCount(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
					if v29 == int32(1) {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						F_UnregisterSnapshot(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v35 != 0 {
								F_pfree(m, v35)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v38 != 0 {
										F_pfree(m, v38)
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return
										} else {
											F_pfree(m, l0)
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
											} else {
												m.G0 = v7 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								}
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v35 != 0 {
							F_pfree(m, v35)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v38 != 0 {
									F_pfree(m, v38)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										F_pfree(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											m.G0 = v7 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							}
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v38 != 0 {
								F_pfree(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									F_pfree(m, l0)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								F_pfree(m, l0)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_index_endscan_0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v51 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_endscan_1), v7)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_endscan_2), int32(397), int32(_a_F_index_endscan_3))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_index_getprocinfo(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+220))
	v14 = int32(1)
	v15 = l1 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)))
	v21 = l2 + v15*v17 - v14
	v24 = v13 + v21*int32(28)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v25 != 0 {
		m.G0 = v11 + int32(16)
		return v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v21<<(uint(int32(2))%32))))
		if v30 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v67 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_getprocinfo_0), v11)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_getprocinfo_1), int32(925), int32(_a_F_index_getprocinfo_2))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
			F_fmgr_info_cxt(m, v30, v24, v34)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v33 == l2 {
					m.G0 = v11 + int32(16)
					return v24
				} else {
					v41 = F_RelationGetIndexAttOptions(m, l0, int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = int32(_a_F_index_getprocinfo_3)
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0]))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
						*(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0])) = v46
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v41+v15<<(uint(int32(2))%32))))
						F_set_fn_opclass_options(m, v24, v51)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_index_getprocinfo[0])) = v44
							m.G0 = v11 + int32(16)
							return v24
						}
					}
				}
			}
		}
	}
}
func F_index_opclass_options(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
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
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v286 int32
	_ = v286
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
	var v307 int32
	_ = v307
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+8)))
	if v19 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v359
L2:
	;
	v82 = v16 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v82))) = int64(0)
	goto L16
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+6)))
	v25 = int32(2)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21*(l1-int32(1))<<(uint(v25)%32)+v19<<(uint(v25)%32)-int32(4))))
	if v33 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v34 = int32(0)
	if base.I32_wrap_i64(l2) == v34 {
		v359 = v34
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v41 = F_SysCacheGetAttrNotNull(m, int32(34), v39, int32(18))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v41)+l1<<(uint(int32(2))%32))+20))
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v57 = m.G0
	v59 = v57 - int32(16)
	m.G0 = v59
	F_initStringInfo(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_get_opclass_name(m, v49, int32(0), v59)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	m.G0 = v59 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v66 + int32(1)
	F_errmsg(m, int32(_a_F_index_opclass_options_0), v16)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_index_opclass_options_1), int32(1049), int32(_a_F_index_opclass_options_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L16:
	;
	v87 = F_index_getprocinfo(m, l0, l1, v19)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v91 = F_FunctionCall1Coll(m, v87, int32(0), base.I64_extend_i32_u(v82))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v93 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v95 = v94
	goto L21
L20:
	;
	v95 = v5
	goto L21
L21:
	;
	v97 = F_palloc_mul(m, int32(12), v95)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v99 == int32(0) {
		v165 = v5
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v170 = F_palloc_mul(m, int32(16), v165)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L32
	}
L24:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if int32(0) < v102 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v106 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v152 == int32(0) {
		v165 = v5
		goto L23
	} else {
		goto L31
	}
L28:
	;
	v121 = v97 + v106*int32(12)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v106<<(uint(int32(2))%32))))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v133
	v136 = v106 + int32(1)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v136 < v137 {
		v106 = v136
		goto L28
	} else {
		goto L30
	}
L29:
	;
	goto L27
L30:
	;
	goto L29
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v165 = v155
	goto L23
L32:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	if v172 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if l2 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v175 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v179 = int32(0)
	goto L36
L36:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v179<<(uint(int32(2))%32))))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v200 = v170 + v179<<(uint(int32(4))%32)
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v200)+4)) = uint8(v201)
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v197
	v205 = v179 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v205 < v206 {
		v179 = v205
		goto L36
	} else {
		goto L38
	}
L37:
	;
	goto L33
L38:
	;
	goto L37
L39:
	;
	F_parseRelOptionsInternal(m, l2, l3, v170, v165)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if int32(0) < v95 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v229 = int32(0)
	v235 = v225
	goto L46
L44:
	;
	v286 = v225
	goto L45
L45:
	;
	v293 = F_palloc0(m, v286)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L8
	} else {
		goto L68
	}
L46:
	;
	v244 = v170 + v229<<(uint(int32(4))%32)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)+20))
	if v246 == int32(5) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v286 = v274
	goto L45
L48:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244)+4)))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v245)+36))
	if v250 != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v274 = v235
	goto L50
L50:
	;
	v278 = v229 + int32(1)
	if v278 != v95 {
		v229 = v278
		v235 = v274
		goto L46
	} else {
		goto L67
	}
L51:
	;
	v274 = v272 + v235
	goto L50
L52:
	;
	if v249&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v249&int32(1) != 0 {
		goto L64
	} else {
		goto L65
	}
L55:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v255 = m.T0[v250].(func(*base.Module, int32, int32) int32)(m, v253, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+28)))
	if v257 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v272 = v255
	goto L51
L59:
	;
	v260 = int32(0)
	goto L61
L60:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v245)+40))
	v260 = v259
	goto L61
L61:
	;
	v262 = m.T0[v250].(func(*base.Module, int32, int32) int32)(m, v260, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	v272 = v262
	goto L51
L63:
	;
	v272 = v269 + int32(1)
	goto L51
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v244)+8))
	v267 = F_strlen(m, v266)
	mBase = m.M
	v269 = v267
	goto L63
L65:
	;
	goto L66
L66:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v245)+24))
	v269 = v268
	goto L63
L67:
	;
	goto L47
L68:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	F_fillRelOptions(m, v293, v295, v170, v95, l3, v97, v95)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	if l3 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	if v97 != 0 {
		goto L78
	} else {
		goto L79
	}
L71:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v300 == int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v303 <= int32(0) {
		goto L70
	} else {
		goto L73
	}
L73:
	;
	v307 = int32(0)
	goto L74
L74:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v300)+12))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320+v307<<(uint(int32(2))%32))))
	m.T0[v324].(func(*base.Module, int32, int32, int32))(m, v293, v170, v95)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L8
	} else {
		goto L76
	}
L75:
	;
	goto L70
L76:
	;
	v328 = v307 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v300)+4))
	if v328 < v329 {
		v307 = v328
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_pfree(m, v97)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L8
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v359 = v293
	goto L1
L81:
	;
	goto L80
}
func F_index_open(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = F_relation_open(m, l0, l1)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+119)))
		if v13|int32(32) != int32(105) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = v25 + int32(4)
					F_errmsg(m, int32(_a_F_index_open_0), v6)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_index_open_1), int32(205), int32(_a_F_index_open_2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
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
			m.G0 = v6 + int32(16)
			return v8
		}
	}
}
func F_index_other_operands_eval_cost(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 float64
	_ = v25
	var v27 int32
	_ = v27
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v84 float64
	_ = v84
	v3 = int32(0)
	v8 = float64(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1 == v3 {
		v84 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v84
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 <= int32(0) {
		v84 = v8
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = v3
	v25 = v8
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v32 == int32(320) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v84 = v72
	goto L1
L6:
	;
	F_cost_qual_eval_node(m, v11+int32(16), v64, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L18
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	v64 = v63
	goto L6
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+28))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v64 = v44
	goto L6
L10:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v64 = v41
	goto L6
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v37 = v35
	v38 = v36
	goto L13
L12:
	;
	v37 = v31
	v38 = v32
	goto L13
L13:
	;
	switch v38 - int32(17) {
	case 0:
		goto L7
	default:
		goto L8
	case 3:
		goto L9
	case 20:
		goto L10
	case 35:
		v64 = int32(0)
		goto L6
	}
L14:
	;
	return float64(0)
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v51
	F_errmsg_internal(m, int32(_a_F_index_other_operands_eval_cost_0), v11)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_index_other_operands_eval_cost_1), int32(_a_F_index_other_operands_eval_cost_2), int32(_a_F_index_other_operands_eval_cost_3))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v72 = base.F64_add(v25, base.F64_add(v69, v70))
	v74 = v23 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v74 < v75 {
		v23 = v74
		v25 = v72
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L5
}
func F_index_pages_fetched(m *base.Module, l0 float64, l1 int32, l2 float64, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	v8 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v8) {
		v11 = v8
	} else {
		v11 = l1
	}
	v12 = base.F64_convert_i32_u(v11)
	v13 = base.F64_add(v12, v12)
	v14 = float64(1)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_index_pages_fetched[0]))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l3)+304))
	v20 = base.F64_add(l2, v19)
	if base.F64_gt(v20, v14) != 0 {
		v24 = v20
	} else {
		v24 = v14
	}
	v25 = base.F64_div(base.F64_mul(v12, base.F64_convert_i32_s(v16)), v24)
	if base.F64_le(v25, float64(1)) != 0 {
		v29 = v14
	} else {
		v29 = base.F64_ceil(v25)
	}
	if base.F64_le(v12, v29) != 0 {
		v33 = base.F64_div(base.F64_mul(l0, v13), base.F64_add(v13, l0))
		if base.F64_ge(v33, v12) != 0 {
			v52 = v12
			return v52
		} else {
			return base.F64_ceil(v33)
		}
	} else {
		v39 = base.F64_div(base.F64_mul(v13, v29), base.F64_sub(v13, v29))
		if base.F64_ge(v39, l0) != 0 {
			v49 = base.F64_div(base.F64_mul(l0, v13), base.F64_add(v13, l0))
		} else {
			v49 = base.F64_add(v29, base.F64_div(base.F64_mul(base.F64_sub(v12, v29), base.F64_sub(l0, v39)), v12))
		}
		v52 = base.F64_ceil(v49)
		return v52
	}
}
func F_index_restrpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+204))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+116))
	if v11 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
		if v12 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
			m.T0[v15].(func(*base.Module, int32))(m, v12)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v19 = v18
				v20 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v20)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v20)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+204))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+116))
				m.T0[v25].(func(*base.Module, int32))(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			}
		} else {
			v19 = v9
			v20 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+66)) = uint8(v20)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v20)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+204))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+116))
			m.T0[v25].(func(*base.Module, int32))(m, l0)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+48))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_index_restrpos_0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v36 + int32(4)
			F_errmsg_internal(m, int32(_a_F_index_restrpos_1), v7)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_index_restrpos_2), int32(453), int32(_a_F_index_restrpos_3))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_index_strategy_get_limit(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	switch l0 - int32(1) {
	case 0:
		v29 = int32(_a_F_index_strategy_get_limit_0)
		v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		m.G0 = v7 + int32(16)
		return v30
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return float64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
			F_errmsg_internal(m, int32(_a_F_index_strategy_get_limit_1), v7)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(_a_F_index_strategy_get_limit_2), int32(316), int32(_a_F_index_strategy_get_limit_3))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 6:
		v29 = int32(_a_F_index_strategy_get_limit_4)
		v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		m.G0 = v7 + int32(16)
		return v30
	case 8:
		v29 = int32(_a_F_index_strategy_get_limit_5)
		v30 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
		m.G0 = v7 + int32(16)
		return v30
	}
}
func F_index_vacuum_cleanup(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_index_vacuum_cleanup[0]))
	if v13 != v11 {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_index_vacuum_cleanup[1]))
		v17 = F_list_member_ptr(m, v16, v11)
		mBase = m.M
		v19 = v17
	} else {
		v19 = int32(1)
	}
	if v19 == int32(0) {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+204))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+56))
		if v23 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_index_vacuum_cleanup_0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v57 + int32(4)
				F_errmsg_internal(m, int32(_a_F_index_vacuum_cleanup_1), v8+int32(16))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_vacuum_cleanup_2), int32(800), int32(_a_F_index_vacuum_cleanup_3))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v26 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, l0, l1)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(32)
				return v26
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41 + int32(4)
				F_errmsg(m, int32(_a_F_index_vacuum_cleanup_4), v8)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_index_vacuum_cleanup_2), int32(799), int32(_a_F_index_vacuum_cleanup_3))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
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
func F_plan_create_index_workers(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 float64
	_ = v131
	var v132 float64
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_plan_create_index_workers[0])))
	if v13 != int32(1) {
		v272 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v272
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	if v17 == int32(0) {
		v272 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = F_palloc0(m, int32(168))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = int64(4294967363)
	v28 = F_palloc0(m, int32(128))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(268)
	v33 = F_palloc0(m, int32(400))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = int32(269)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+360)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+300)) = v42
	v47 = F_palloc0(m, int32(8))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(275)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v47
	v56 = F_list_make1_impl(m, int32(1), v10+int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+92)) = v56
	v60 = F_palloc0(m, int32(136))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v62 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(101)
	v69 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+124)) = uint16(v69)
	v71 = int32(_a_F_plan_create_index_workers_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+20)) = uint16(v71)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v60
	v78 = F_list_make1_impl(m, v62, v10+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v78
	v83 = F_addRTEPermissionInfo(m, v21+int32(56), v60)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_setup_simple_rel_arrays(m, v33)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v89 = F_build_simple_rel(m, v33, int32(1), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v92 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v95 = F_index_open(m, l1, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v97 = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)+48))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+118)))
	if v99 == int32(116) {
		v259 = v97
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_relation_close(m, v95, int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L4
	} else {
		goto L81
	}
L18:
	;
	v102 = F_RelationGetIndexExpressions(m, v95)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v104 = F_is_parallel_safe(m, v33, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	if v104 == int32(0) {
		v259 = v97
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v108 = F_RelationGetIndexPredicate(m, v95)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v110 = F_is_parallel_safe(m, v33, v108)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	if v110 == int32(0) {
		v259 = v97
		goto L17
	} else {
		goto L24
	}
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+156))
	if v114 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	if v114 < v118 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_estimate_rel_size(m, v92, int32(0), v10+int32(44), v10+int32(32), v10+int32(24))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L31
	}
L28:
	;
	v120 = v114
	goto L30
L29:
	;
	v120 = v118
	goto L30
L30:
	;
	v259 = v120
	goto L17
L31:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v131 = base.F64_convert_i32_u(v130)
	v132 = float64(-1)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[1]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v89)+156))
	if v137 != int32(-1) {
		v226 = v137
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if v236 <= int32(0) {
		goto L72
	} else {
		goto L73
	}
L33:
	;
	goto L32
L34:
	;
	if v226 < v134 {
		goto L69
	} else {
		goto L70
	}
L35:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v140 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v157 = int32(0)
	if base.F64_ge(v131, float64(0)) == v157 {
		v189 = v157
		goto L44
	} else {
		goto L45
	}
L37:
	;
	if base.F64_ge(v131, float64(0)) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[3]))
	if base.F64_lt(v131, base.F64_convert_i32_s(v145)) != 0 {
		v236 = int32(0)
		goto L33
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if base.F64_ge(v132, float64(0)) == int32(0) {
		goto L36
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[4]))
	if base.F64_lt(v132, base.F64_convert_i32_s(v154)) != 0 {
		v236 = int32(0)
		goto L33
	} else {
		goto L43
	}
L43:
	;
	goto L36
L44:
	;
	if base.F64_ge(v132, float64(0)) == int32(0) {
		v226 = v189
		goto L34
	} else {
		goto L53
	}
L45:
	;
	v162 = int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[3]))
	if v164 <= v162 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v167 = v162
	goto L48
L47:
	;
	v167 = v164
	goto L48
L48:
	;
	v169 = v167
	v173 = int32(1)
	goto L49
L49:
	;
	v176 = v169 * int32(3)
	if base.F64_ge(v131, base.F64_convert_i32_u(v176)) == int32(0) {
		v189 = v173
		goto L44
	} else {
		goto L51
	}
L50:
	;
	v189 = v182
	goto L44
L51:
	;
	v182 = v173 + int32(1)
	if v176 < int32(715827883) {
		v169 = v176
		v173 = v182
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v195 = int32(1)
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[4]))
	if v197 <= v195 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v200 = v195
	goto L56
L55:
	;
	v200 = v197
	goto L56
L56:
	;
	v202 = v200
	v207 = int32(1)
	goto L57
L57:
	;
	v209 = v202 * int32(3)
	if base.F64_le(base.F64_convert_i32_u(v209), v132) != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v189 < v216 {
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v213 = v207 + int32(1)
	if v209 < int32(715827883) {
		v202 = v209
		v207 = v213
		goto L57
	} else {
		goto L62
	}
L60:
	;
	v216 = v207
	goto L61
L61:
	;
	goto L58
L62:
	;
	v216 = v213
	goto L61
L63:
	;
	v218 = v189
	goto L65
L64:
	;
	v218 = v216
	goto L65
L65:
	;
	if int32(0) < v189 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v221 = v218
	goto L68
L67:
	;
	v221 = v216
	goto L68
L68:
	;
	v226 = v221
	goto L34
L69:
	;
	v229 = v226
	goto L71
L70:
	;
	v229 = v134
	goto L71
L71:
	;
	v236 = v229
	goto L33
L72:
	;
	v259 = v236
	goto L17
L73:
	;
	goto L74
L74:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_plan_create_index_workers[5]))
	v241 = v236
	goto L75
L75:
	;
	v250 = base.I32_div_s(v240, v241+int32(1))
	if int32(_a_F_plan_create_index_workers_1) < v250 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v259 = v97
	goto L17
L77:
	;
	v259 = v241
	goto L17
L78:
	;
	goto L79
L79:
	;
	v253 = int32(1)
	if v253 < v241 {
		v241 = v241 - v253
		goto L75
	} else {
		goto L80
	}
L80:
	;
	goto L76
L81:
	;
	F_relation_close(m, v92, int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v272 = v259
	goto L1
}
func F_validate_index(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int64
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v267 float32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int64
	_ = v281
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int64
	_ = v296
	var v299 int64
	_ = v299
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v454 int64
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 float64
	_ = v554
	var v556 float64
	_ = v556
	var v558 float64
	_ = v558
	var v562 int32
	_ = v562
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	v11 = m.G0
	v13 = v11 - int32(160)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+144)) = v16
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+136)) = v19
	v22 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+128)) = v22
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+112)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v13)+104)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = int64(4)
	v36 = v13 + int32(128)
	v38 = v13 + int32(80)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
	if v48 == int32(0) {
	} else {
		v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
		if v52&int32(1) == int32(0) {
		} else {
			v57 = int32(_a_F_validate_index_0)
			v59 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
			v60 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v59 + v60
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			*(*int32)(unsafe.Add(mBase, uint32(v48))) = v63 + v60
			v67 = int32(0)
			v70 = base.AtomicRmwOr32(m, v67, int32(_a_F_validate_index_1), v67)
			v76 = v48 + int32(232)
			v82 = int32(0)
			v85 = int32(0)
			for {
				v91 = int32(2)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v36+v85<<(uint(v91)%32))))
				v95 = int32(3)
				v101 = *(*int64)(unsafe.Add(mBase, uint32(v38+v85<<(uint(v95)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v76+v94<<(uint(v95)%32)))) = v101
				v104 = v85 | int32(1)
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v36+v104<<(uint(v91)%32))))
				v115 = *(*int64)(unsafe.Add(mBase, uint32(v38+v104<<(uint(v95)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v76+v108<<(uint(v95)%32)))) = v115
				v118 = v85 | v91
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v36+v118<<(uint(v91)%32))))
				v129 = *(*int64)(unsafe.Add(mBase, uint32(v38+v118<<(uint(v95)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v76+v122<<(uint(v95)%32)))) = v129
				v132 = v85 | v95
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v36+v132<<(uint(v91)%32))))
				v143 = *(*int64)(unsafe.Add(mBase, uint32(v38+v132<<(uint(v95)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v76+v136<<(uint(v95)%32)))) = v143
				v145 = int32(4)
				v146 = v85 + v145
				v148 = v82 + v145
				if v148 != int32(4) {
					v82 = v148
					v85 = v146
					continue
				} else {
					break
				}
				break
			}
			v162 = int32(0)
			v165 = v146
			for {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v36+v165<<(uint(int32(2))%32))))
				v175 = int32(3)
				v181 = *(*int64)(unsafe.Add(mBase, uint32(v38+v165<<(uint(v175)%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v76+v174<<(uint(v175)%32)))) = v181
				v183 = int32(1)
				v186 = v162 + v183
				if v186 != int32(1) {
					v162 = v186
					v165 = v165 + v183
					continue
				} else {
					break
				}
				break
			}
			v197 = int32(0)
			v200 = base.AtomicRmwOr32(m, v197, int32(_a_F_validate_index_1), v197)
			v201 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			v202 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v48))) = v201 + v202
			v205 = int32(_a_F_validate_index_0)
			v207 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
			*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v207 - v202
		}
	}
	v221 = F_table_open(m, l0, int32(4))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		return
	} else {
		v228 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[6]))
		*(*int32)(unsafe.Add(mBase, uint32(v13+int32(124)))) = v228
		v231 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[7]))
		*(*int32)(unsafe.Add(mBase, uint32(v13+int32(120)))) = v231
		v233 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
		v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+80))
		v235 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v235 | int32(2)
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v234
		v243 = int32(_a_F_validate_index_2)
		v245 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[8]))
		v247 = v245 + int32(1)
		*(*int32)(unsafe.Add(mBase, _c_F_validate_index[8])) = v247
		F_RestrictSearchPath(m)
		mBase = m.M
		v250 = m.ExcPending
		if v250 != 0 {
			return
		} else {
			v252 = F_index_open(m, l1, int32(3))
			mBase = m.M
			v253 = m.ExcPending
			if v253 != 0 {
				return
			} else {
				v254 = F_BuildIndexInfo(m, v252)
				mBase = m.M
				v255 = m.ExcPending
				if v255 != 0 {
					return
				} else {
					v256 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v254)+121)) = uint8(v256)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+92)) = int32(13)
					*(*uint8)(unsafe.Add(mBase, uint32(v13)+90)) = uint8(v256)
					v262 = int32(256)
					*(*uint16)(unsafe.Add(mBase, uint32(v13)+88)) = uint16(v262)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v221
					*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v252
					v266 = *(*int32)(unsafe.Add(mBase, uint32(v221)+48))
					v267 = *(*float32)(unsafe.Add(mBase, uint32(v266)+100))
					v268 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v268
					*(*float64)(unsafe.Add(mBase, uint32(v13)+96)) = base.F64_promote_f32(v267)
					v277 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[9]))
					v279 = F_tuplesort_begin_datum(m, int32(20), int32(412), v268, v268, v277, v268)
					mBase = m.M
					v280 = m.ExcPending
					if v280 != 0 {
						return
					} else {
						v281 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+136)) = v281
						*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v279
						*(*int64)(unsafe.Add(mBase, uint32(v13)+144)) = v281
						*(*int64)(unsafe.Add(mBase, uint32(v13)+152)) = v281
						v290 = F_index_bulk_delete(m, v38, int32(0), int32(503), v36)
						mBase = m.M
						v291 = m.ExcPending
						if v291 != 0 {
							return
						} else {
							v293 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[10]))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v293
							v296 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[11]))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v296
							v299 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[12]))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v299
							v302 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[13]))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v302
							v305 = *(*int64)(unsafe.Add(mBase, _c_F_validate_index[14]))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v305
							v321 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
							if v321 == int32(0) {
							} else {
								v325 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
								if v325&int32(1) == int32(0) {
								} else {
									v330 = int32(_a_F_validate_index_0)
									v332 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
									v333 = int32(1)
									*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v332 + v333
									v336 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
									*(*int32)(unsafe.Add(mBase, uint32(v321))) = v336 + v333
									v340 = int32(0)
									v343 = base.AtomicRmwOr32(m, v340, int32(_a_F_validate_index_1), v340)
									v435 = int32(0)
									v438 = int32(0)
									for {
										v447 = *(*int32)(unsafe.Add(mBase, uint32(v13-int32(-64)+v438<<(uint(int32(2))%32))))
										v448 = int32(3)
										v454 = *(*int64)(unsafe.Add(mBase, uint32(v13+int32(32)+v438<<(uint(v448)%32))))
										*(*int64)(unsafe.Add(mBase, uint32(v321+int32(232)+v447<<(uint(v448)%32)))) = v454
										v456 = int32(1)
										v459 = v435 + v456
										if v459 != int32(3) {
											v435 = v459
											v438 = v438 + v456
											continue
										} else {
											break
										}
										break
									}
									v470 = int32(0)
									v473 = base.AtomicRmwOr32(m, v470, int32(_a_F_validate_index_1), v470)
									v474 = *(*int32)(unsafe.Add(mBase, uint32(v321)))
									v475 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v321))) = v474 + v475
									v478 = int32(_a_F_validate_index_0)
									v480 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
									*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v480 - v475
								}
							}
							v493 = *(*int32)(unsafe.Add(mBase, uint32(v13)+128))
							F_tuplesort_performsort(m, v493)
							mBase = m.M
							v495 = m.ExcPending
							if v495 != 0 {
								return
							} else {
								v500 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[3]))
								if v500 == int32(0) {
								} else {
									v504 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_index[4])))
									if v504&int32(1) == int32(0) {
									} else {
										v509 = int32(_a_F_validate_index_0)
										v511 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
										v512 = int32(1)
										*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v511 + v512
										v515 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
										*(*int32)(unsafe.Add(mBase, uint32(v500))) = v515 + v512
										v519 = int32(0)
										v521 = int32(_a_F_validate_index_1)
										v522 = base.AtomicRmwOr32(m, v519, v521, v519)
										*(*int64)(unsafe.Add(mBase, uint32(v500+int32(72))+232)) = int64(6)
										v530 = base.AtomicRmwOr32(m, v519, v521, v519)
										v531 = *(*int32)(unsafe.Add(mBase, uint32(v500)))
										*(*int32)(unsafe.Add(mBase, uint32(v500))) = v531 + v512
										v537 = *(*int32)(unsafe.Add(mBase, _c_F_validate_index[5]))
										*(*int32)(unsafe.Add(mBase, _c_F_validate_index[5])) = v537 - v512
									}
								}
								v541 = *(*int32)(unsafe.Add(mBase, uint32(v221)+188))
								v542 = *(*int32)(unsafe.Add(mBase, uint32(v541)+144))
								m.T0[v542].(func(*base.Module, int32, int32, int32, int32, int32))(m, v221, v252, v254, l2, v36)
								mBase = m.M
								v544 = m.ExcPending
								if v544 != 0 {
									return
								} else {
									v545 = *(*int32)(unsafe.Add(mBase, uint32(v13)+128))
									F_tuplesort_end(m, v545)
									mBase = m.M
									v547 = m.ExcPending
									if v547 != 0 {
										return
									} else {
										F_index_insert_cleanup(m, v252, v254)
										mBase = m.M
										v549 = m.ExcPending
										if v549 != 0 {
											return
										} else {
											v552 = F_errstart(m, int32(13), int32(0))
											mBase = m.M
											v553 = m.ExcPending
											if v553 != 0 {
												return
											} else {
												if v552 != 0 {
													v554 = *(*float64)(unsafe.Add(mBase, uint32(v13)+152))
													*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = v554
													v556 = *(*float64)(unsafe.Add(mBase, uint32(v13)+136))
													*(*float64)(unsafe.Add(mBase, uint32(v13))) = v556
													v558 = *(*float64)(unsafe.Add(mBase, uint32(v13)+144))
													*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v558
													F_errmsg_internal(m, int32(_a_F_validate_index_3), v13)
													mBase = m.M
													v562 = m.ExcPending
													if v562 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_validate_index_4), int32(3588), int32(_a_F_validate_index_5))
														mBase = m.M
														v567 = m.ExcPending
														if v567 != 0 {
															return
														} else {
															F_AtEOXact_GUC(m, int32(0), v247)
															mBase = m.M
															v570 = m.ExcPending
															if v570 != 0 {
																return
															} else {
																v571 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
																v572 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
																*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v572
																*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v571
																F_relation_close(m, v252, int32(0))
																mBase = m.M
																v579 = m.ExcPending
																if v579 != 0 {
																	return
																} else {
																	F_relation_close(m, v221, int32(0))
																	mBase = m.M
																	v582 = m.ExcPending
																	if v582 != 0 {
																		return
																	} else {
																		m.G0 = v13 + int32(160)
																		return
																	}
																}
															}
														}
													}
												} else {
													F_AtEOXact_GUC(m, int32(0), v247)
													mBase = m.M
													v570 = m.ExcPending
													if v570 != 0 {
														return
													} else {
														v571 = *(*int32)(unsafe.Add(mBase, uint32(v13)+124))
														v572 = *(*int32)(unsafe.Add(mBase, uint32(v13)+120))
														*(*int32)(unsafe.Add(mBase, _c_F_validate_index[7])) = v572
														*(*int32)(unsafe.Add(mBase, _c_F_validate_index[6])) = v571
														F_relation_close(m, v252, int32(0))
														mBase = m.M
														v579 = m.ExcPending
														if v579 != 0 {
															return
														} else {
															F_relation_close(m, v221, int32(0))
															mBase = m.M
															v582 = m.ExcPending
															if v582 != 0 {
																return
															} else {
																m.G0 = v13 + int32(160)
																return
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
	}
}
