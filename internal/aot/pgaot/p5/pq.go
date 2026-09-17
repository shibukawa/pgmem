package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pq_copymsgbytes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	if l2 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_pq_copymsgbytes_0), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_pq_copymsgbytes_1), int32(533), int32(_a_F_pq_copymsgbytes_2))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v7-v8 < l2 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				F_errcode(m, int32(16908800))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_pq_copymsgbytes_0), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pq_copymsgbytes_1), int32(533), int32(_a_F_pq_copymsgbytes_2))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if l2 != 0 {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, l1, v11+v8, l2)
			} else {
			}
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v14 + l2
			return
		}
	}
}
func F_pq_endmessage_reuse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+12)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pq_endmessage_reuse[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = m.T0[v7].(func(*base.Module, int32, int32, int32) int32)(m, v2, v3, v4)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_pq_getkeepalivesidle(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l0 == v2 {
		v21 = v2
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
		if v10 == int32(1) {
			v21 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+400))
			if v13 != 0 {
				v21 = v13
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+384))
				if v14 != 0 {
					v21 = v14
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(4)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(384))))
					v21 = v20
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return v21
}
func F_pq_getmessage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v281 int32
	_ = v281
	var v292 int32
	_ = v292
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v384 int32
	_ = v384
	var v385 int64
	_ = v385
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v413 int32
	_ = v413
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(192)
	m.G0 = v13
	v19 = v3
	v20 = v3
	v21 = int32(-1)
	v22 = v3
	v23 = v3
	goto L2
L1:
	;
	m.G0 = v13 + int32(192)
	return v413
L2:
	;
	goto L4
L3:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L4:
	;
	if v21 != int32(1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v384 = int32(m.ExcTag)
	v385 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v384 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[0])) = v174
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[1])) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	v292 = int32(0)
	if v171 == v292 {
		v333 = v292
		goto L70
	} else {
		goto L71
	}
L8:
	;
	v281 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pq_getmessage[2])) = uint8(v281)
	v413 = v281
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v32)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32
	goto L12
L10:
	;
	v171 = v19
	v172 = v20
	v174 = v22
	v175 = v23
	goto L11
L11:
	;
	if v172 != 0 {
		goto L7
	} else {
		goto L45
	}
L12:
	;
	v45 = v13 + int32(176)
	v46 = int32(4)
	goto L13
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[3]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4]))
	if v52 <= v54 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v13)+176))
	v106 = int32(16711935)
	v114 = base.I32_rotr(v105&v106, int32(8)) | base.I32_rotr(v105, int32(24))&v106
	if base.B2i32(int32(4) <= v114)&base.B2i32(v114 <= l1) == int32(0) {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	v59 = F_pq_recvbuf(m)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v94 = v52 - v54
	if base.Ui32(v94) < base.Ui32(v46) {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	if v59 == int32(0) {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	v68 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v70 = int32(-1)
	if v68 == int32(0) {
		v413 = v70
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errcode(m, int32(16908800))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errmsg(m, int32(_a_F_pq_getmessage_0), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errfinish(m, int32(_a_F_pq_getmessage_1), int32(1217), int32(_a_F_pq_getmessage_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v413 = v70
	goto L1
L25:
	;
	v96 = v94
	goto L27
L26:
	;
	v96 = v46
	goto L27
L27:
	;
	if v96 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	base.MemoryCopy(m, v45, v54+int32(_a_F_pq_getmessage_3), v96)
	goto L30
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4])) = v96 + v54
	v104 = v46 - v96
	if v104 != 0 {
		v45 = v45 + v96
		v46 = v104
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L14
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	v126 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	if v105 == int32(67108864) {
		goto L8
	} else {
		goto L40
	}
L35:
	;
	v128 = int32(-1)
	if v126 == int32(0) {
		v413 = v128
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errcode(m, int32(16908800))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errmsg(m, int32(_a_F_pq_getmessage_4), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v19
	F_errfinish(m, int32(_a_F_pq_getmessage_1), int32(1227), int32(_a_F_pq_getmessage_2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v413 = v128
	goto L1
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[0]))
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[1]))
	goto L41
L41:
	;
	v162 = v13 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = v13 + int32(12)
	goto L44
L42:
	;
	v171 = v114 - int32(4)
	v172 = int32(0)
	v174 = v158
	v175 = v160
	goto L11
L44:
	;
	goto L42
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[1])) = v13 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_enlargeStringInfo(m, l0, v171)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[0])) = v174
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[1])) = v175
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v171 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v196 = v191
	v197 = v171
	goto L50
L48:
	;
	v261 = v191
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v171
	v269 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171+v261))) = uint8(v269)
	goto L8
L50:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[3]))
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4]))
	if v203 <= v205 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v261 = v256
	goto L49
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	v210 = F_pq_recvbuf(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v245 = v203 - v205
	if base.Ui32(v245) < base.Ui32(v197) {
		goto L62
	} else {
		goto L63
	}
L55:
	;
	if v210 == int32(0) {
		goto L50
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	v219 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L6
	} else {
		goto L57
	}
L57:
	;
	v221 = int32(-1)
	if v219 == int32(0) {
		v413 = v221
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errcode(m, int32(16908800))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errmsg(m, int32(_a_F_pq_getmessage_5), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errfinish(m, int32(_a_F_pq_getmessage_1), int32(1262), int32(_a_F_pq_getmessage_2))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	v413 = v221
	goto L1
L62:
	;
	v247 = v245
	goto L64
L63:
	;
	v247 = v197
	goto L64
L64:
	;
	if v247 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	base.MemoryCopy(m, v196, v205+int32(_a_F_pq_getmessage_3), v247)
	goto L67
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4])) = v247 + v205
	v255 = v197 - v247
	if v255 != 0 {
		v196 = v196 + v247
		v197 = v255
		goto L50
	} else {
		goto L68
	}
L68:
	;
	goto L51
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_pq_getmessage[2])) = uint8(v368)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_pg_re_throw(m)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L89
	}
L70:
	;
	if v333 != int32(-1) {
		goto L69
	} else {
		goto L83
	}
L71:
	;
	v300 = v171
	goto L72
L72:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[3]))
	v308 = *(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4]))
	if v306 <= v308 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v333 = int32(0)
	goto L70
L74:
	;
	v310 = F_pq_recvbuf(m)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v316 = v306 - v308
	if base.Ui32(v316) < base.Ui32(v300) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v310 == int32(0) {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	v333 = int32(-1)
	goto L70
L79:
	;
	v318 = v316
	goto L81
L80:
	;
	v318 = v300
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pq_getmessage[4])) = v308 + v318
	v321 = v300 - v318
	if v321 != 0 {
		v300 = v321
		goto L72
	} else {
		goto L82
	}
L82:
	;
	goto L73
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	v341 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v341 == int32(0) {
		goto L69
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errcode(m, int32(16908800))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errmsg(m, int32(_a_F_pq_getmessage_5), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+184)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v13)+180)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v171
	F_errfinish(m, int32(_a_F_pq_getmessage_1), int32(1249), int32(_a_F_pq_getmessage_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L88
	}
L88:
	;
	goto L69
L89:
	;
	goto L5
L90:
	;
	v389 = int32(v385)
	m.G0 = v13
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v392)))
	if v13+int32(12) == v395 {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	m.ExcPending = 1
	goto L99
L92:
	;
	if v399 != 0 {
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	v399 = v397
	goto L95
L94:
	;
	v399 = int32(0)
	goto L95
L95:
	;
	goto L92
L96:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v13)+184))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v13)+180))
	v19 = v400
	v20 = v391
	v21 = v399
	v22 = v402
	v23 = v401
	goto L2
L97:
	;
	goto L98
L98:
	;
	F___wasm_longjmp(m, v392, v391)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	return int32(0)
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pq_putemptymessage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pq_putemptymessage[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = m.T0[v6].(func(*base.Module, int32, int32, int32) int32)(m, l0, v2, v2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_pq_puttextmessage(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = int32(_a_F_pq_puttextmessage_0)
	v6 = F_strlen(m, v4)
	mBase = m.M
	v7 = F_pg_server_to_client(m, v4, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_pq_puttextmessage[0]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		if v7 != int32(_a_F_pq_puttextmessage_0) {
			v15 = F_strlen(m, v7)
			mBase = m.M
			v18 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), v7, v15+int32(1))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_pfree(m, v7)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v26 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, int32(67), int32(_a_F_pq_puttextmessage_0), v6+int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_sendstring(m *base.Module, l0 int32, l1 int32) {
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
	var v19 int32
	_ = v19
	v5 = F_strlen(m, l1)
	v6 = F_pg_server_to_client(m, l1, v5)
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if l1 != v6 {
			v9 = F_strlen(m, v6)
			F_appendBinaryStringInfoNT(m, l0, v6, v9+int32(1))
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_pfree(m, v6)
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_appendBinaryStringInfoNT(m, l0, l1, v5+int32(1))
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_pq_settcpusertimeout(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l1 == int32(0) {
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if v11 == int32(1) {
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+412))
			if l0 == v14 {
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
				if int32(0) < v16 {
					if l0 == int32(0) {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+412)) = v49
				} else {
					v19 = int32(0)
					v21 = m.G0
					v23 = v21 - int32(16)
					m.G0 = v23
					if l1 == v19 {
						v38 = v19
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
						if v27 == int32(1) {
							v38 = v19
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+412))
							if v30 != 0 {
								v38 = v30
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
								if v31 != 0 {
									v38 = v31
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(4)
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(396))))
									v38 = v37
								}
							}
						}
					}
					m.G0 = v23 + int32(16)
					if int32(0) <= v38 {
						if l0 == int32(0) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+396))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v46
						} else {
						}
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+412)) = v49
					} else {
					}
				}
			}
		}
	}
	m.G0 = v6 + int32(16)
	return
}
func F_pq_startmsgread(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pq_startmsgread[0])))
	if v2 != 0 {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			F_errcode(m, int32(16908800))
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_pq_startmsgread_0), int32(0))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_pq_startmsgread_1), int32(1151), int32(_a_F_pq_startmsgread_2))
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v20 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_pq_startmsgread[0])) = uint8(v20)
		return
	}
}
