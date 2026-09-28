package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_HnswCheckNorm(m *base.Module, l0 int32, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = F_FunctionCall1Coll(m, v3, v4, l1)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.F64_gt(base.F64_reinterpret_i64(v5), float64(0))
	}
}
func F_HnswFindElementNeighbors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 float64
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int64
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+65)))
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = base.I64_extend_i32_u(v33)
	if l3 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	v33 = v24
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v26 == v25 {
		v33 = v25
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = l0 + v26 - int32(1)
	goto L1
L6:
	;
	v38 = l1 - l0
	v39 = int32(16)
	v43 = (int32(base.Ui32(v38)>>(uint(v39)%32)) ^ v38) * int32(-2048144789)
	v48 = (int32(base.Ui32(v43)>>(uint(int32(13))%32)) ^ v43) * int32(-1028477387)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = int32(base.Ui32(v48)>>(uint(v39)%32)) ^ v48
	goto L8
L7:
	;
	goto L8
L8:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l7 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	m.G0 = v19 + int32(16)
	return
L12:
	;
	v55 = l1
	goto L14
L13:
	;
	v55 = int32(0)
	goto L14
L14:
	;
	v59 = F_HnswEntryCandidate(m, l0, l2, v19+int32(8), l3, l4, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v59
	v64 = F_list_make1_impl(m, int32(1), v19)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	if base.Ui32(v21) < base.Ui32(v66) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v70 = v66
	v80 = v64
	goto L21
L19:
	;
	v109 = v64
	goto L20
L20:
	;
	if base.Ui32(v21) < base.Ui32(v66) {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v86 = int32(1)
	v88 = int32(0)
	v92 = F_HnswSearchLayer(m, l0, v19+int32(8), v80, v86, v70, l3, l4, l5, v86, v55, v88, v88, v86, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L15
	} else {
		goto L23
	}
L22:
	;
	v109 = v92
	goto L20
L23:
	;
	v95 = v70 - int32(1)
	if base.Ui32(v21) < base.Ui32(v95) {
		v70 = v95
		v80 = v92
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v114 = v21
	goto L27
L26:
	;
	v114 = v66
	goto L27
L27:
	;
	v118 = v114
	v128 = v109
	goto L28
L28:
	;
	v134 = int32(1)
	v135 = int32(0)
	v139 = F_HnswSearchLayer(m, l0, v19+int32(8), v128, l6+l7, v118, l3, l4, l5, v134, v55, v135, v135, v134, v135)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L15
	} else {
		goto L32
	}
L29:
	;
	goto L11
L30:
	;
	v285 = l5 << (uint(base.B2i32(v118 == int32(0))) % 32)
	if l0 != 0 {
		goto L64
	} else {
		goto L65
	}
L31:
	;
	v204 = int32(0)
	v208 = base.AtomicRmwOr32(m, v204, int32(_a_F_HnswFindElementNeighbors_0), v204)
	if v188 == v204 {
		v273 = v204
		goto L30
	} else {
		goto L46
	}
L32:
	;
	if v139 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v141 = int32(0)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v141 < v143 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v197 = int32(0)
	if l3 == v197 {
		v273 = v197
		goto L30
	} else {
		goto L45
	}
L36:
	;
	v148 = v141
	v153 = v141
	goto L39
L37:
	;
	v188 = v141
	goto L38
L38:
	;
	if l3 != 0 {
		goto L31
	} else {
		goto L44
	}
L39:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v139)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v148<<(uint(int32(2))%32))))
	v168 = F_palloc(m, int32(12))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L15
	} else {
		goto L41
	}
L40:
	;
	v188 = v175
	goto L38
L41:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v170
	v172 = *(*float64)(unsafe.Add(mBase, uint32(v166)+32))
	*(*float32)(unsafe.Add(mBase, uint32(v168)+4)) = base.F32_demote_f64(v172)
	v175 = F_lappend(m, v153, v168)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	v178 = v148 + int32(1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v178 < v179 {
		v148 = v178
		v153 = v175
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v273 = v188
	goto L30
L45:
	;
	v200 = int32(0)
	v203 = base.AtomicRmwOr32(m, v200, int32(_a_F_HnswFindElementNeighbors_0), v200)
	v273 = v197
	goto L30
L46:
	;
	v211 = int32(0)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v212 <= v211 {
		v273 = v204
		goto L30
	} else {
		goto L47
	}
L47:
	;
	v217 = v211
	v221 = v204
	v228 = v212
	goto L48
L48:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v217<<(uint(int32(2))%32))))
	if l0 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v273 = v262
	goto L30
L50:
	;
	if v55 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	v247 = v238
	goto L50
L52:
	;
	goto L53
L53:
	;
	v239 = int32(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	if v240 == v239 {
		v247 = v239
		goto L50
	} else {
		goto L54
	}
L54:
	;
	v247 = l0 + v240 - int32(1)
	goto L50
L55:
	;
	v265 = v217 + int32(1)
	if v265 < v263 {
		v217 = v265
		v221 = v262
		v228 = v263
		goto L48
	} else {
		goto L62
	}
L56:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+64)))
	if v256 == int32(0) {
		v262 = v221
		v263 = v228
		goto L55
	} else {
		goto L60
	}
L57:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v247)+76))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v250 != v251 {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+80)))
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+80)))
	if v253 == v254 {
		v262 = v221
		v263 = v228
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	v259 = F_lappend(m, v221, v235)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L61
	}
L61:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	v262 = v259
	v263 = v261
	goto L55
L62:
	;
	goto L49
L63:
	;
	if v331 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	v286 = int32(0)
	v288 = v118 << (uint(int32(2)) % 32)
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288+(l0+v289)-int32(1))))
	v301 = F_SelectNeighbors(m, l0, v273, v285, l4, l0+v294+int32(3), v286, v286, v286)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L15
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v314 = int32(0)
	v316 = v118 << (uint(int32(2)) % 32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316+v317)))
	v325 = F_SelectNeighbors(m, v314, v273, v285, l4, v319+int32(4), v314, v314, v314)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L15
	} else {
		goto L69
	}
L67:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0+v303+v288-int32(1))))
	if v308 == int32(0) {
		v331 = v301
		v332 = v286
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v331 = v301
	v332 = l0 + v308 - int32(1)
	goto L63
L69:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v327+v316)))
	v331 = v325
	v332 = v329
	goto L63
L70:
	;
	if int32(0) < v118 {
		v118 = v118 - int32(1)
		v128 = v139
		goto L28
	} else {
		goto L76
	}
L71:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v335 <= int32(0) {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v343 = int32(0)
	goto L73
L73:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v359 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v358 + v359
	v364 = v332 + int32(8) + v358*int32(12)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v357+v343<<(uint(int32(2))%32))))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+8)) = v369
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v368)))
	*(*int64)(unsafe.Add(mBase, uint32(v364))) = v371
	v374 = v343 + v359
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v374 < v375 {
		v343 = v374
		goto L73
	} else {
		goto L75
	}
L74:
	;
	goto L70
L75:
	;
	goto L74
L76:
	;
	goto L29
}
func F_HnswGetMetaPageInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v8 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_LockBufferInternal(m, v8, int32(1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v8 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, _c_F_HnswGetMetaPageInfo[0]))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+(v8^int32(-1))<<(uint(int32(2))%32))))
				v30 = v22
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_HnswGetMetaPageInfo[1]))
				v30 = v24 + v8<<(uint(int32(13))%32) + int32(-8192)
			}
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
			if v31 == int32(-1454134957) {
				if l1 != 0 {
					v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+36)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v34
				} else {
				}
				if l2 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v30)+40))
					if v36 != int32(-1) {
						v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+44)))
						v41 = F_palloc(m, int32(108))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v41)+80)) = uint16(v39)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+76)) = v36
							v45 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v41)+88)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(v41)+72)) = v45
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v41
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+46)))
							*(*uint8)(unsafe.Add(mBase, uint32(v41)+65)) = uint8(v50)
							F_UnlockReleaseBuffer(m, v8)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
						F_UnlockReleaseBuffer(m, v8)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_UnlockReleaseBuffer(m, v8)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswGetMetaPageInfo_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_HnswGetMetaPageInfo_1), int32(311), int32(_a_F_HnswGetMetaPageInfo_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
	}
}
func F_HnswInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 float64
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_HnswInit[0])))
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
		v16 = F_LWLockAcquire(m, v12+int32(2688), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v20 = F_ShmemInitStruct(m, v5+int32(15))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
				if v22 == int32(1) {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					v30 = v25
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v30
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
					F_LWLockRelease(m, v34+int32(2688))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						v41 = F_add_reloption_kind(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v41
							F_add_int_reloption(m, v41, int32(_a_F_HnswInit_0), int32(_a_F_HnswInit_1), int32(16), int32(2), int32(100))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
								F_add_int_reloption(m, v52, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(64), int32(4), int32(1000))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(40), int32(1), int32(1000), int32(6), int32(0))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										v73 = int32(0)
										F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), v73, int32(_a_F_HnswInit_10), v73, int32(_a_F_HnswInit_11), int32(6))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v82 = int32(0)
											F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_12), int32(_a_F_HnswInit_13), v82, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), int32(1), int32(2147483647), int32(6), v82)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												v95 = float64(1)
												F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(0), int32(_a_F_HnswInit_18), v95, v95, float64(1000), int32(6))
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_19))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														m.G0 = v5 + int32(16)
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
				} else {
					v27 = F_LWLockNewTrancheId(m, int32(_a_F_HnswInit_20))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v27
						v30 = v27
						*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[2])) = v30
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[1]))
						F_LWLockRelease(m, v34+int32(2688))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v41 = F_add_reloption_kind(m)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v41
								F_add_int_reloption(m, v41, int32(_a_F_HnswInit_0), int32(_a_F_HnswInit_1), int32(16), int32(2), int32(100))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									v52 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
									F_add_int_reloption(m, v52, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(64), int32(4), int32(1000))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(40), int32(1), int32(1000), int32(6), int32(0))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v73 = int32(0)
											F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), v73, int32(_a_F_HnswInit_10), v73, int32(_a_F_HnswInit_11), int32(6))
											mBase = m.M
											v79 = m.ExcPending
											if v79 != 0 {
												return
											} else {
												v82 = int32(0)
												F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_12), int32(_a_F_HnswInit_13), v82, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), int32(1), int32(2147483647), int32(6), v82)
												mBase = m.M
												v90 = m.ExcPending
												if v90 != 0 {
													return
												} else {
													v95 = float64(1)
													F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(0), int32(_a_F_HnswInit_18), v95, v95, float64(1000), int32(6))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_19))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															m.G0 = v5 + int32(16)
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
	} else {
		v41 = F_add_reloption_kind(m)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3])) = v41
			F_add_int_reloption(m, v41, int32(_a_F_HnswInit_0), int32(_a_F_HnswInit_1), int32(16), int32(2), int32(100))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInit[3]))
				F_add_int_reloption(m, v52, int32(_a_F_HnswInit_2), int32(_a_F_HnswInit_3), int32(64), int32(4), int32(1000))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_4), int32(_a_F_HnswInit_5), int32(_a_F_HnswInit_6), int32(_a_F_HnswInit_7), int32(40), int32(1), int32(1000), int32(6), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v73 = int32(0)
						F_DefineCustomEnumVariable(m, int32(_a_F_HnswInit_8), int32(_a_F_HnswInit_9), v73, int32(_a_F_HnswInit_10), v73, int32(_a_F_HnswInit_11), int32(6))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							v82 = int32(0)
							F_DefineCustomIntVariable(m, int32(_a_F_HnswInit_12), int32(_a_F_HnswInit_13), v82, int32(_a_F_HnswInit_14), int32(_a_F_HnswInit_15), int32(1), int32(2147483647), int32(6), v82)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return
							} else {
								v95 = float64(1)
								F_DefineCustomRealVariable(m, int32(_a_F_HnswInit_16), int32(_a_F_HnswInit_17), int32(0), int32(_a_F_HnswInit_18), v95, v95, float64(1000), int32(6))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_MarkGUCPrefixReserved(m, int32(_a_F_HnswInit_19))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return
									} else {
										m.G0 = v5 + int32(16)
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
func F_HnswInitLockTranche(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
	v12 = F_LWLockAcquire(m, v8+int32(2688), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v16 = F_ShmemInitStruct(m, v5+int32(15))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+15)))
			if v18 == int32(1) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v26 = v21
				*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v26
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
				F_LWLockRelease(m, v30+int32(2688))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			} else {
				v23 = F_LWLockNewTrancheId(m, int32(_a_F_HnswInitLockTranche_0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v23
					v26 = v23
					*(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[1])) = v26
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInitLockTranche[0]))
					F_LWLockRelease(m, v30+int32(2688))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_HnswLoadElementImpl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int64
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v61 float64
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v2 = l1
	v12 = F_ReadBuffer(m, l4, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		F_LockBufferInternal(m, v12, int32(1))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v12 < int32(0) {
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[0]))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20+(v12^int32(-1))<<(uint(int32(2))%32))))
				v34 = v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_HnswLoadElementImpl[1]))
				v34 = v28 + v12<<(uint(int32(13))%32) + int32(-8192)
			}
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v2<<(uint(int32(2))%32))+20))
			v41 = v38&int32(_a_F_HnswLoadElementImpl_0) + v34
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
			if v42 == int32(0) {
				if l2 == int32(0) {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
					if v67 == int32(0) {
						v71 = F_palloc(m, int32(108))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v2)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = l0
							v75 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v71)+88)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(l8))) = v71
							v80 = v71
							F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						v80 = v67
						F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return
						} else {
							F_UnlockReleaseBuffer(m, v12)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v47 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
					if base.I32_wrap_i64(v47) != 0 {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
						v54 = F_FunctionCall2Coll(m, v49, v50, v47, base.I64_extend_i32_u(v41+int32(72)))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = v54
							*(*int64)(unsafe.Add(mBase, uint32(l2))) = v57
							if l7 == int32(0) {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
								if v67 == int32(0) {
									v71 = F_palloc(m, int32(108))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v2)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = l0
										v75 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+88)) = v75
										*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v75
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v71
										v80 = v71
										F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v80 = v67
									F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v61 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
								if base.F64_gt(v61, base.F64_reinterpret_i64(v57)) == int32(0) {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										return
									}
								} else {
									v67 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
									if v67 == int32(0) {
										v71 = F_palloc(m, int32(108))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v2)
											*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = l0
											v75 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v71)+88)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v75
											*(*int32)(unsafe.Add(mBase, uint32(l8))) = v71
											v80 = v71
											F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												F_UnlockReleaseBuffer(m, v12)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return
												} else {
													return
												}
											}
										}
									} else {
										v80 = v67
										F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												return
											}
										}
									}
								}
							}
						}
					} else {
						v57 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l2))) = v57
						if l7 == int32(0) {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
							if v67 == int32(0) {
								v71 = F_palloc(m, int32(108))
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v2)
									*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = l0
									v75 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v71)+88)) = v75
									*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v75
									*(*int32)(unsafe.Add(mBase, uint32(l8))) = v71
									v80 = v71
									F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
											return
										}
									}
								}
							} else {
								v80 = v67
								F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									F_UnlockReleaseBuffer(m, v12)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return
									} else {
										return
									}
								}
							}
						} else {
							v61 = *(*float64)(unsafe.Add(mBase, uint32(l7)))
							if base.F64_gt(v61, base.F64_reinterpret_i64(v57)) == int32(0) {
								F_UnlockReleaseBuffer(m, v12)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									return
								}
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
								if v67 == int32(0) {
									v71 = F_palloc(m, int32(108))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										*(*uint16)(unsafe.Add(mBase, uint32(v71)+80)) = uint16(v2)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = l0
										v75 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v71)+88)) = v75
										*(*int32)(unsafe.Add(mBase, uint32(v71)+72)) = v75
										*(*int32)(unsafe.Add(mBase, uint32(l8))) = v71
										v80 = v71
										F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											F_UnlockReleaseBuffer(m, v12)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return
											} else {
												return
											}
										}
									}
								} else {
									v80 = v67
									F_HnswLoadElementFromTuple(m, v80, v41, int32(1), l6)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										F_UnlockReleaseBuffer(m, v12)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return
										} else {
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
				v91 = m.ExcPending
				if v91 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_HnswLoadElementImpl_1), int32(0))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_HnswLoadElementImpl_2), int32(550), int32(_a_F_HnswLoadElementImpl_3))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
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
	}
}
func F_HnswParallelBuildMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v10 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_HnswParallelBuildMain[0])) = v10
		F_pgstat_report_activity(m, int32(3), v10)
		mBase = m.M
		v17 = F_shm_toc_lookup(m, l1, int64(-6917529027641081855), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+8)))
			if v22 != 0 {
				v23 = int32(4)
			} else {
				v23 = int32(5)
			}
			v24 = F_table_open(m, v19, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				if v22 != 0 {
					v29 = int32(3)
				} else {
					v29 = int32(8)
				}
				v30 = F_index_open(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					v34 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						F_HnswParallelScanAndInsert(m, v24, v30, v17, v34, int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_relation_close(m, v30, v29)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_relation_close(m, v24, v23)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
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
func F_HnswParallelScanAndInsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 float64
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v10 = m.G0
	v12 = v10 - int32(224)
	m.G0 = v12
	v14 = F_BuildIndexInfo(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+121)) = uint8(v16)
		v19 = v12 + int32(16)
		F_InitBuildState_1(m, v19, l0, l1, v14, int32(0))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = l2 + int32(40)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+204)) = int32(_a_F_HnswParallelScanAndInsert_0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v19
			v31 = int32(0)
			v38 = F_table_beginscan_parallel(m, l0, l2+int32(160), v31)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+140))
				v42 = m.T0[v41].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, v14, int32(1), v31, l4, v31, int32(-1), int32(_a_F_HnswParallelScanAndInsert_1), v19, v38)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v46 = base.AtomicRmwXchg32(m, l2, int32(24), int32(1))
					if v46 != 0 {
						F_s_lock(m, l2+int32(24), int32(_a_F_HnswParallelScanAndInsert_2))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v52 + int32(1)
							v56 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v56, v42)
							v59 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+24)), uint32(v59))
							v64 = F_errstart(m, int32(14), v59)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								if v64 != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_trunc_sat_f64_s(v42)
									if l4 != 0 {
										v70 = int32(_a_F_HnswParallelScanAndInsert_3)
									} else {
										v70 = int32(_a_F_HnswParallelScanAndInsert_4)
									}
									F_errmsg(m, v70, v12)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										if l4 != 0 {
											v76 = int32(825)
										} else {
											v76 = int32(827)
										}
										F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_5), v76, int32(_a_F_HnswParallelScanAndInsert_6))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											F_ConditionVariableSignal(m, l2+int32(12))
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return
											} else {
												v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
												F_MemoryContextDelete(m, v84)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return
												} else {
													v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
													F_MemoryContextDelete(m, v87)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														m.G0 = v12 + int32(224)
														return
													}
												}
											}
										}
									}
								} else {
									F_ConditionVariableSignal(m, l2+int32(12))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return
									} else {
										v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
										F_MemoryContextDelete(m, v84)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return
										} else {
											v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
											F_MemoryContextDelete(m, v87)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												m.G0 = v12 + int32(224)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v52 + int32(1)
						v56 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
						*(*float64)(unsafe.Add(mBase, uint32(l2)+32)) = base.F64_add(v56, v42)
						v59 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l2)+24)), uint32(v59))
						v64 = F_errstart(m, int32(14), v59)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return
						} else {
							if v64 != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(v12))) = base.I64_trunc_sat_f64_s(v42)
								if l4 != 0 {
									v70 = int32(_a_F_HnswParallelScanAndInsert_3)
								} else {
									v70 = int32(_a_F_HnswParallelScanAndInsert_4)
								}
								F_errmsg(m, v70, v12)
								mBase = m.M
								v72 = m.ExcPending
								if v72 != 0 {
									return
								} else {
									if l4 != 0 {
										v76 = int32(825)
									} else {
										v76 = int32(827)
									}
									F_errfinish(m, int32(_a_F_HnswParallelScanAndInsert_5), v76, int32(_a_F_HnswParallelScanAndInsert_6))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										F_ConditionVariableSignal(m, l2+int32(12))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return
										} else {
											v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
											F_MemoryContextDelete(m, v84)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return
											} else {
												v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
												F_MemoryContextDelete(m, v87)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													m.G0 = v12 + int32(224)
													return
												}
											}
										}
									}
								}
							} else {
								F_ConditionVariableSignal(m, l2+int32(12))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return
								} else {
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
									F_MemoryContextDelete(m, v84)
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return
									} else {
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
										F_MemoryContextDelete(m, v87)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											m.G0 = v12 + int32(224)
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
func F_hnsw_bit_support(m *base.Module, l0 int32) int64 {
	return int64(4171848)
}
