package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_datumTransfer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v6 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(511796), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476045), int32(1367), int32(215106))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v24 = int32(4443856)
		v25 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v27
		v30 = F_datumTransfer(m, l0, int32(0), l1)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v25
			return v30
		}
	}
}
func F_SPI_execute_plan_with_paramlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v3 = l2
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(-6)
	if l0 == int32(0) {
		v62 = v12
		m.G0 = v10 + int32(32)
		return v62
	} else {
		if l3 < int32(0) {
			v62 = v12
			m.G0 = v10 + int32(32)
			return v62
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v17 != int32(569278163) {
				v62 = v12
				m.G0 = v10 + int32(32)
				return v62
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, _consts[434]))
				if v21 == int32(0) {
					v62 = int32(-4)
					m.G0 = v10 + int32(32)
					return v62
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
					v29 = *(*int32)(unsafe.Add(mBase, _consts[434]))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = v27
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v32
					v34 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v34
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(l3)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v34
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v3)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = l1
					v44 = int32(0)
					v47 = F__SPI_execute_plan(m, l0, v10+int32(8), v44, v44, int32(1))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, _consts[434]))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+20))
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v54
						*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = int32(0)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+24))
						F_MemoryContextReset(m, v58)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v62 = v47
							m.G0 = v10 + int32(32)
							return v62
						}
					}
				}
			}
		}
	}
}
func F__SPI_commit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v18 = v2
	v19 = v2
	v20 = v2
	v21 = v2
	v22 = v2
	v23 = v2
	v24 = v13
	v25 = int32(-1)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v25 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v285 = int32(m.ExcTag)
	v286 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v285 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	v29 = v24 - int32(16)
	m.G0 = v29
	v32 = v29 - int32(160)
	m.G0 = v32
	v35 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v37 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+40)))
	if v38 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v147 = v18
	v148 = v19
	v149 = v20
	v150 = v21
	v151 = v22
	v152 = v23
	v153 = v24
	goto L9
L9:
	;
	if v152 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	v83 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+28))
	goto L17
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errcode(m, int32(1282))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errmsg(m, int32(249761), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errfinish(m, int32(476045), int32(241), int32(94882))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L16
	}
L16:
	;
	goto L3
L17:
	;
	if int32(1) < v84 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if l0 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errcode(m, int32(1282))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errmsg(m, int32(327347), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	F_errfinish(m, int32(476045), int32(256), int32(94882))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		v284 = v32
		goto L6
	} else {
		goto L24
	}
L24:
	;
	goto L3
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	v129 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v129
	v132 = int32(*(*uint8)(unsafe.Add(mBase, _consts[143])))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v132)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _consts[144])))
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)) = uint8(v135)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	v141 = *(*int32)(unsafe.Add(mBase, _consts[362]))
	goto L29
L28:
	;
	goto L27
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v13 + int32(8)
	goto L32
L30:
	;
	v147 = v32
	v148 = v29
	v149 = v35
	v150 = v141
	v151 = v139
	v152 = int32(0)
	v153 = v32
	goto L9
L32:
	;
	goto L30
L33:
	;
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v147
	v159 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v160 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v159)+41)) = uint8(v160)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_HoldPinnedPortals(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v150
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v228 = F_CopyErrorData(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L44
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_ForgetPortalSnapshots(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_CommitTransactionCommand(m)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_StartTransactionCommand(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L39
	}
L39:
	;
	if l0 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v196
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v199)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v202)
	goto L43
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v149
	v207 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v208 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v207)+41)) = uint8(v208)
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v151
	*(*int32)(unsafe.Add(mBase, _consts[362])) = v150
	m.G0 = v13 + int32(32)
	return
L43:
	;
	goto L42
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_FlushErrorState(m)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_AbortCurrentTransaction(m)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_StartTransactionCommand(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if l0 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	*(*int32)(unsafe.Add(mBase, _consts[67])) = v257
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+4)))
	*(*uint8)(unsafe.Add(mBase, _consts[143])) = uint8(v260)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+5)))
	*(*uint8)(unsafe.Add(mBase, _consts[144])) = uint8(v263)
	goto L51
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v149
	v268 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v268)+41)) = uint8(v269)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v148
	F_ReThrowError(m, v228)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		v284 = v153
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	goto L5
L53:
	;
	v290 = int32(v286)
	m.G0 = v284
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v290)+4))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v290)))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293)))
	if v13+int32(8) == v297 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	m.ExcPending = 1
	goto L62
L55:
	;
	if v300 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v293)+4))
	v300 = v299
	goto L58
L57:
	;
	v300 = int32(0)
	goto L58
L58:
	;
	goto L55
L59:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v18 = v302
	v19 = v301
	v20 = v303
	v21 = v304
	v22 = v305
	v23 = v292
	v24 = v284
	v25 = v300
	goto L1
L60:
	;
	goto L61
L61:
	;
	F___wasm_longjmp(m, v293, v292)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__SPI_make_plan_non_temp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	v10 = *(*int32)(unsafe.Add(mBase, _consts[434]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v16 = F_AllocSetContextCreateInternal(m, v11, int32(270697), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(4443856)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v16
	v25 = F_palloc0(m, int32(40))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = int32(569278163)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+24)) = v34
	if int32(0) < v34 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v56 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v40 = F_palloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = int32(0)
	goto L4
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+28)) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v46 = v44 << (uint(int32(2)) % 32)
	if v46 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L4
L10:
	;
	v47 = F__emscripten_memcpy_bulkmem(m, v40, v43, v46)
	mBase = m.M
	goto L12
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return v25
L14:
	;
	v59 = int32(0)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v60 <= v59 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v64 = v59
	goto L16
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v64<<(uint(int32(2))%32))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+94)))
	if v76 != int32(1) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L13
L18:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	v178 = F_lappend(m, v177, v75)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L67
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L64
	}
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+92)))
	if v79 == int32(1) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L61
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+56))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	if v86 != v11 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v75)+88))
	if v116 != 0 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	if v86 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	if v11 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+28))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v91 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v90 == int32(0) {
		goto L28
	} else {
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+28)) = v90
	goto L30
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v90
	goto L30
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v96
	goto L28
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = v11
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v82)+28)) = v103
	if v103 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v82)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = int32(0)
	goto L27
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+24)) = v82
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v82
	goto L24
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+32))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	if v121 != v11 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	goto L43
L43:
	;
	goto L18
L44:
	;
	goto L43
L45:
	;
	if v121 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	if v11 != 0 {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v117)+28))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
	if v126 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v125 == int32(0) {
		goto L48
	} else {
		goto L54
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v126)+28)) = v125
	goto L50
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+20)) = v125
	goto L50
L54:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v117)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+24)) = v131
	goto L48
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = v11
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+28)) = v138
	if v138 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v117)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+16)) = int32(0)
	goto L47
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v138)+24)) = v117
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v117
	goto L44
L61:
	;
	F_errmsg_internal(m, int32(58224), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(477493), int32(1619), int32(59597))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errmsg_internal(m, int32(58170), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(477493), int32(1621), int32(59597))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v178
	v182 = v64 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v182 < v183 {
		v64 = v182
		goto L16
	} else {
		goto L68
	}
L68:
	;
	goto L17
}
